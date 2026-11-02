package browser_test

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"

	"carrel.dev/typecase/internal/browser"
	"carrel.dev/typecase/internal/gate"
)

// Two tools for looking at the mockup, run on demand. Neither is a gate: the
// ladder and the style check guard the file; these answer "what does it look
// like" and "what moved" — the questions a person reviewing a change asks.
// A list built of rail nav items folded to icons at 641–1007 passed every
// check and was caught only on a picture (docs/visual/mockup.md).

// injectFrameJS puts one frame into the shell page. Width 0 leaves the frame
// as drawn; any other width is a step of the ladder, as TestMockupLadder sets it.
const injectFrameJS = `(function(){
  var host = document.getElementById("host");
  host.innerHTML = %s;
  var f = host.querySelector(".frame");
  if (!f) return false;
  var w = %d;
  if (w > 0) {
    f.style.width = "max-content";
    f.style.maxWidth = "none";
    f.style.margin = "0";
    var root = f.querySelector(".tc-root");
    if (root) {
      root.style.width = w + "px";
      root.style.maxWidth = w + "px";
      root.style.flex = "none";
      // A label longer than the screen wraps, so a shot is the screen's width.
      f.style.width = "min-content";
      var label = f.querySelector(".frame-label");
      if (label) label.style.flexWrap = "wrap";
    }
  }
  return true;
})()`

// frameRectsJS lists every element box of the frame, relative to the frame.
const frameRectsJS = `(function(){
  var f = document.querySelector(".frame");
  if (!f) return [];
  var fr = f.getBoundingClientRect();
  var out = [];
  f.querySelectorAll("*").forEach(function(el){
    var b = el.getBoundingClientRect();
    var c = "";
    try { c = el.getAttribute("class") || ""; } catch (e) {}
    out.push({box: [Math.round(b.left-fr.left), Math.round(b.top-fr.top), Math.round(b.width), Math.round(b.height)].join(","),
              name: (el.tagName || "").toLowerCase() + (c ? "." + c.split(/\s+/).join(".") : "")});
  });
  return out;
})()`

type frameRect struct {
	Box  string `json:"box"`
	Name string `json:"name"`
}

// frameShotTargets parses "s-dup-1@390,s-plan-2" into frame and width (0: as drawn).
func frameShotTargets(spec string) (map[string][]int, error) {
	out := map[string][]int{}
	for _, item := range strings.Split(spec, ",") {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		id, ws, found := strings.Cut(item, "@")
		w := 0
		if found {
			n, err := strconv.Atoi(ws)
			if err != nil {
				return nil, fmt.Errorf("%q: width %q is not a number", item, ws)
			}
			w = n
		}
		out[id] = append(out[id], w)
	}
	return out, nil
}

func TestFrameShotTargets(t *testing.T) {
	got, err := frameShotTargets("s-dup-1@390, s-plan-2,s-dup-1@1351")
	if err != nil {
		t.Fatal(err)
	}
	if fmt.Sprint(got["s-dup-1"]) != "[390 1351]" || fmt.Sprint(got["s-plan-2"]) != "[0]" {
		t.Fatalf("got %v", got)
	}
	if _, err := frameShotTargets("s-dup-1@wide"); err == nil {
		t.Fatal("a width that is not a number must fail")
	}
}

// TestMockupFrameShots takes a full-size picture of frames of the working
// mockup, at the widths asked, into TC_FRAME_SHOTS_DIR (a temp dir by default).
//
//	TC_FRAME_SHOTS=s-dup-1@390,s-plan-2 go test ./internal/browser -run TestMockupFrameShots -v
//
// A frame without @width is shot as drawn; the widths are the ladder's
// (browser.MockupLadderWidths), though any number works. TC_FRAME_SHOTS_REV
// shoots the mockup of a git revision instead of the working one, for a
// picture of «before».
func TestMockupFrameShots(t *testing.T) {
	spec := os.Getenv("TC_FRAME_SHOTS")
	if spec == "" {
		t.Skip("set TC_FRAME_SHOTS=<frame>[@width],… to take pictures")
	}
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	targets, err := frameShotTargets(spec)
	if err != nil {
		t.Fatal(err)
	}
	dir := os.Getenv("TC_FRAME_SHOTS_DIR")
	if dir == "" {
		dir = filepath.Join(os.TempDir(), "tc-frame-shots")
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	root := browser.ModuleRoot(t)
	html, err := browser.LoadMockupHTML(root, os.Getenv("TC_FRAME_SHOTS_REV"))
	if err != nil {
		t.Fatal(err)
	}
	parts, err := gate.ParseMockupParts(string(html))
	if err != nil {
		t.Fatal(err)
	}
	h := browser.NewHarness(t)
	run, cancel := h.Tab(10 * time.Minute)
	defer cancel()
	url := h.RegisterPage("shots", browser.MockupShotShell(parts.Styles, parts.Sprite))
	if err := run(
		chromedp.EmulateViewport(int64(browser.MockupShotViewport), int64(browser.MockupShotMaxHeight)),
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.Evaluate(fontsReadyJS, nil),
	); err != nil {
		t.Fatal(err)
	}
	seen := map[string]bool{}
	for _, fr := range parts.Frames {
		ws, ok := targets[fr.ID]
		if !ok {
			continue
		}
		seen[fr.ID] = true
		raw, err := json.Marshal(fr.HTML)
		if err != nil {
			t.Fatal(err)
		}
		for _, w := range ws {
			var buf []byte
			if err := run(
				chromedp.EmulateViewport(int64(browser.MockupShotViewport), int64(browser.MockupShotMaxHeight)),
				chromedp.Evaluate(fmt.Sprintf(injectFrameJS, raw, w), nil),
				chromedp.Sleep(50*time.Millisecond),
				chromedp.ActionFunc(func(ctx context.Context) error {
					var box browser.Rect
					if err := chromedp.Evaluate(frameBoxJS, &box).Do(ctx); err != nil {
						return err
					}
					vw := int64(math.Max(float64(browser.MockupShotViewport), math.Ceil(box.Right)+10))
					vh := int64(math.Min(6000, math.Ceil(box.Bottom)+10))
					if err := chromedp.EmulateViewport(vw, vh).Do(ctx); err != nil {
						return err
					}
					if err := chromedp.Evaluate(frameBoxJS, &box).Do(ctx); err != nil {
						return err
					}
					png, err := page.CaptureScreenshot().WithFormat(page.CaptureScreenshotFormatPng).Do(ctx)
					if err != nil {
						return err
					}
					buf, err = cropPNG(png, math.Floor(box.X), math.Floor(box.Y), math.Ceil(box.Width), math.Min(math.Ceil(box.Height), float64(vh)), 1)
					return err
				}),
			); err != nil {
				t.Fatalf("%s @%d: %v", fr.ID, w, err)
			}
			name := fr.ID + ".png"
			if w > 0 {
				name = fmt.Sprintf("%s@%d.png", fr.ID, w)
			}
			path := filepath.Join(dir, name)
			if err := os.WriteFile(path, buf, 0644); err != nil {
				t.Fatal(err)
			}
			t.Logf("%s", path)
		}
	}
	for id := range targets {
		if !seen[id] {
			t.Errorf("no frame %q in the mockup", id)
		}
	}
}

// TestMockupFrameMoves says which frames moved against a commit, and at which
// steps of the ladder: every element box of every frame, before and after, on
// all ten widths. The style check compares two widths and stops printing at 24
// nodes; this one is for the list a person reviews by eye.
//
//	TC_FRAME_MOVES=HEAD~1 go test ./internal/browser -run TestMockupFrameMoves -v -timeout 30m
//
// TC_FRAME_MOVES_ONLY=s-dup-1,s-plan-2 limits it to some frames.
func TestMockupFrameMoves(t *testing.T) {
	base := strings.TrimSpace(os.Getenv("TC_FRAME_MOVES"))
	if base == "" {
		t.Skip("set TC_FRAME_MOVES=<commit> to compare the working mockup with it")
	}
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	only := map[string]bool{}
	for _, id := range strings.Split(os.Getenv("TC_FRAME_MOVES_ONLY"), ",") {
		if id = strings.TrimSpace(id); id != "" {
			only[id] = true
		}
	}
	root := browser.ModuleRoot(t)
	afterHTML, err := browser.LoadMockupHTML(root, "")
	if err != nil {
		t.Fatal(err)
	}
	beforeHTML, err := browser.LoadMockupHTML(root, base)
	if err != nil {
		t.Fatal(err)
	}
	after, err := gate.ParseMockupParts(string(afterHTML))
	if err != nil {
		t.Fatal(err)
	}
	before, err := gate.ParseMockupParts(string(beforeHTML))
	if err != nil {
		t.Fatal(err)
	}
	beforeByID := map[string]string{}
	for _, f := range before.Frames {
		beforeByID[f.ID] = f.HTML
	}

	measure := func(parts gate.MockupParts, htmlOf func(id string) (string, bool)) map[string][]frameRect {
		t.Helper()
		h := browser.NewHarness(t)
		run, cancel := h.Tab(30 * time.Minute)
		defer func() {
			cancel()
			h.CloseBrowser()
		}()
		url := h.RegisterPage("moves", browser.MockupShotShell(parts.Styles, parts.Sprite))
		if err := run(
			chromedp.EmulateViewport(int64(browser.MockupShotViewport), int64(browser.MockupShotMaxHeight)),
			chromedp.Navigate(url),
			chromedp.WaitReady("body"),
			chromedp.Evaluate(fontsReadyJS, nil),
		); err != nil {
			t.Fatal(err)
		}
		out := map[string][]frameRect{}
		for _, fr := range after.Frames {
			if len(only) > 0 && !only[fr.ID] {
				continue
			}
			html, ok := htmlOf(fr.ID)
			if !ok {
				continue
			}
			raw, err := json.Marshal(html)
			if err != nil {
				t.Fatal(err)
			}
			for _, w := range browser.MockupLadderWidths {
				var got []frameRect
				if err := run(
					chromedp.Evaluate(fmt.Sprintf(injectFrameJS, raw, w), nil),
					chromedp.Sleep(20*time.Millisecond),
					chromedp.Evaluate(frameRectsJS, &got),
				); err != nil {
					t.Fatalf("%s @%d: %v", fr.ID, w, err)
				}
				out[fmt.Sprintf("%s@%d", fr.ID, w)] = got
			}
		}
		return out
	}
	now := measure(after, func(id string) (string, bool) {
		if _, ok := beforeByID[id]; !ok {
			return "", false // a new frame has nothing to move against
		}
		for _, f := range after.Frames {
			if f.ID == id {
				return f.HTML, true
			}
		}
		return "", false
	})
	was := measure(before, func(id string) (string, bool) {
		html, ok := beforeByID[id]
		return html, ok
	})

	type verdict struct {
		widths []int
		first  string
		tree   bool
	}
	res := map[string]*verdict{}
	// widest first, so "first" names the move at the widest step it shows at
	keys := make([]string, 0, len(now))
	for k := range now {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		ii, wi, _ := strings.Cut(keys[i], "@")
		ij, wj, _ := strings.Cut(keys[j], "@")
		if ii != ij {
			return ii < ij
		}
		ni, _ := strconv.Atoi(wi)
		nj, _ := strconv.Atoi(wj)
		return ni > nj
	})
	for _, k := range keys {
		a, b := now[k], was[k]
		id, ws, _ := strings.Cut(k, "@")
		w, _ := strconv.Atoi(ws)
		v := res[id]
		if v == nil {
			v = &verdict{}
			res[id] = v
		}
		if len(a) != len(b) {
			v.tree = true
			v.widths = append(v.widths, w)
			continue
		}
		for i := range a {
			if a[i].Box != b[i].Box {
				v.widths = append(v.widths, w)
				if v.first == "" {
					v.first = fmt.Sprintf("%s %s → %s", a[i].Name, b[i].Box, a[i].Box)
				}
				break
			}
		}
	}
	ids := make([]string, 0, len(res))
	for id, v := range res {
		if len(v.widths) > 0 {
			ids = append(ids, id)
		}
	}
	sort.Strings(ids)
	for _, id := range ids {
		v := res[id]
		sort.Sort(sort.Reverse(sort.IntSlice(v.widths)))
		ws := fmt.Sprint(v.widths)
		if len(v.widths) == len(browser.MockupLadderWidths) {
			ws = "all widths"
		}
		switch {
		case v.tree:
			t.Logf("%-16s %s · the tree changed, boxes not compared", id, ws)
		default:
			t.Logf("%-16s %s · first: %s", id, ws, v.first)
		}
	}
	t.Logf("%d of %d frames moved against %s", len(ids), len(res), base)
}
