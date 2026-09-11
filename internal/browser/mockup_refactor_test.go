package browser_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"

	"carrel.dev/typecase/internal/browser"
	"carrel.dev/typecase/internal/gate"
)

const mockupStageSize = 1

func TestParseAllowedDiffs(t *testing.T) {
	md := "База: abc\n\n| Кадр | Ширина | Решение | Что меняется |\n|------|-------:|---------|----------------|\n| s-notes-1 | 820 | Р-7 | меню |\n"
	got, err := browser.ParseAllowedDiffs(md)
	if err != nil {
		t.Fatal(err)
	}
	if got.Base != "abc" {
		t.Fatalf("base: %q", got.Base)
	}
	if len(got.Rows) != 1 || got.Rows[0].Frame != "s-notes-1" || got.Rows[0].Width != 820 {
		t.Fatalf("rows: %+v", got.Rows)
	}
	if !got.Skip("s-notes-1", 820) || got.Skip("s-notes-1", 1351) {
		t.Fatal("skip mismatch")
	}
}

func TestPNGsMatch(t *testing.T) {
	img := image.NewNRGBA(image.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.NRGBA{R: 10, G: 20, B: 30, A: 255})
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		t.Fatal(err)
	}
	pngBytes := append([]byte(nil), buf.Bytes()...)
	if err := browser.PNGsMatch(pngBytes, pngBytes, browser.MockupChannelTol); err != nil {
		t.Fatal(err)
	}
	img.Set(0, 0, color.NRGBA{R: 40, G: 20, B: 30, A: 255})
	buf.Reset()
	if err := png.Encode(&buf, img); err != nil {
		t.Fatal(err)
	}
	if err := browser.PNGsMatch(pngBytes, buf.Bytes(), browser.MockupChannelTol); err == nil {
		t.Fatal("expected mismatch")
	}
}

type mockupStage struct {
	name   string
	frames []gate.MockupFrame
}

func frameSection(id string) string {
	i := strings.LastIndex(id, "-")
	if i <= 0 {
		return id
	}
	return id[:i]
}

func mockupStages(frames []gate.MockupFrame, chunk int) []mockupStage {
	if chunk < 1 {
		chunk = 1
	}
	var groups [][]gate.MockupFrame
	var cur []gate.MockupFrame
	var sec string
	for _, f := range frames {
		s := frameSection(f.ID)
		if len(cur) > 0 && (s != sec || len(cur) >= chunk) {
			groups = append(groups, cur)
			cur = nil
		}
		sec = s
		cur = append(cur, f)
	}
	if len(cur) > 0 {
		groups = append(groups, cur)
	}
	out := make([]mockupStage, 0, len(groups))
	for _, g := range groups {
		sec := frameSection(g[0].ID)
		a := strings.TrimPrefix(g[0].ID, sec+"-")
		name := sec + "/" + a
		if len(g) > 1 {
			name += "-" + strings.TrimPrefix(g[len(g)-1].ID, sec+"-")
		}
		out = append(out, mockupStage{name: name, frames: g})
	}
	return out
}

func TestMockupStages(t *testing.T) {
	frames := []gate.MockupFrame{
		{ID: "s-notes-1"}, {ID: "s-notes-2"}, {ID: "s-notes-3"},
		{ID: "s-files-1"},
	}
	got := mockupStages(frames, 2)
	if len(got) != 3 || got[0].name != "s-notes/1-2" || got[1].name != "s-notes/3" || got[2].name != "s-files/1" {
		names := make([]string, len(got))
		for i, g := range got {
			names[i] = g.name
		}
		t.Fatalf("%v", names)
	}
}

func TestMockupRefactor(t *testing.T) {
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if os.Getenv("CI") == "" && testing.Short() {
		t.Skip("mockup refactor screenshots need Chrome; skip with -short outside CI")
	}

	root := browser.ModuleRoot(t)
	spec, err := browser.LoadAllowedDiffs(root)
	if err != nil {
		t.Fatal(err)
	}
	afterHTML, err := browser.LoadMockupHTML(root, "")
	if err != nil {
		t.Fatal(err)
	}
	beforeHTML, err := browser.LoadMockupHTML(root, spec.Base)
	if err != nil {
		t.Fatal(err)
	}
	if err := gate.CheckMockupLint(string(afterHTML)); err != nil {
		t.Fatalf("current mockup: %v", err)
	}
	after, err := gate.ParseMockupParts(string(afterHTML))
	if err != nil {
		t.Fatal(err)
	}
	before, err := gate.ParseMockupParts(string(beforeHTML))
	if err != nil {
		t.Fatal(err)
	}
	beforeByID := make(map[string]gate.MockupFrame, len(before.Frames))
	for _, f := range before.Frames {
		beforeByID[f.ID] = f
	}
	sameDoc := spec.Base == "" || string(beforeHTML) == string(afterHTML)

	for _, st := range mockupStages(after.Frames, mockupStageSize) {
		st := st
		t.Run(st.name, func(t *testing.T) {
			if spec.Base != "" {
				for _, fr := range st.frames {
					if _, ok := beforeByID[fr.ID]; !ok {
						t.Fatalf("frame %s missing from base %s", fr.ID, spec.Base)
					}
				}
			}
			h := browser.NewHarness(t)
			run, cancel := h.Tab(3 * time.Minute)
			defer cancel()

			open := func(styles, sprite, id string) {
				t.Helper()
				url := h.RegisterPage(id, browser.MockupShotShell(styles, sprite))
				if err := run(
					chromedp.EmulateViewport(int64(browser.MockupShotViewport), 900, chromedp.EmulateScale(browser.MockupShotScale)),
					chromedp.Navigate(url),
					chromedp.WaitReady("body"),
				); err != nil {
					t.Fatal(err)
				}
			}
			shot := func(html, id string, width int) []byte {
				t.Helper()
				raw, err := json.Marshal(html)
				if err != nil {
					t.Fatal(err)
				}
				js := fmt.Sprintf(`(function(){
  var host = document.getElementById("host");
  host.innerHTML = %s;
  var f = host.querySelector(".frame");
  if (!f) return false;
  f.style.width = %d + "px";
  f.style.maxWidth = %d + "px";
  f.style.maxHeight = %d + "px";
  f.style.overflow = "hidden";
  f.style.margin = "0";
  return true;
})()`, raw, width, width, browser.MockupShotMaxHeight)
				var png []byte
				if err := run(
					chromedp.Evaluate(js, nil),
					chromedp.Sleep(50*time.Millisecond),
					chromedp.ActionFunc(func(ctx context.Context) error {
						buf, err := captureFramePNG(ctx)
						if err != nil {
							return err
						}
						png = buf
						return nil
					}),
				); err != nil {
					t.Fatalf("%s @%d: %v", id, width, err)
				}
				if len(png) == 0 {
					t.Fatalf("%s @%d: empty screenshot", id, width)
				}
				return png
			}

			open(after.Styles, after.Sprite, "after")
			afterPNG := map[string][]byte{}
			for _, fr := range st.frames {
				for _, w := range browser.MockupRefactorWidths {
					if spec.Skip(fr.ID, w) {
						continue
					}
					afterPNG[fmt.Sprintf("%s@%d", fr.ID, w)] = shot(fr.HTML, fr.ID, w)
				}
			}
			if sameDoc {
				return
			}
			open(before.Styles, before.Sprite, "before")
			for _, fr := range st.frames {
				src := beforeByID[fr.ID].HTML
				for _, w := range browser.MockupRefactorWidths {
					if spec.Skip(fr.ID, w) {
						continue
					}
					key := fmt.Sprintf("%s@%d", fr.ID, w)
					got := shot(src, fr.ID, w)
					if err := browser.PNGsMatch(got, afterPNG[key], browser.MockupChannelTol); err != nil {
						t.Errorf("%s @%d: %v", fr.ID, w, err)
					}
				}
			}
		})
	}
}

const frameBoxJS = `(function(){
  var f = document.querySelector(".frame");
  if (!f) return null;
  var r = f.getBoundingClientRect();
  return {x:r.x,y:r.y,width:r.width,height:r.height,top:r.top,left:r.left,bottom:r.bottom,right:r.right};
})()`

func captureFramePNG(ctx context.Context) ([]byte, error) {
	var box browser.Rect
	if err := chromedp.Evaluate(frameBoxJS, &box).Do(ctx); err != nil {
		return nil, err
	}
	if box.Width < 1 || box.Height < 1 {
		return nil, fmt.Errorf("empty frame box")
	}
	vw := int64(browser.MockupShotViewport)
	vh := int64(browser.MockupShotMaxHeight)
	if err := chromedp.EmulateViewport(vw, vh, chromedp.EmulateScale(browser.MockupShotScale)).Do(ctx); err != nil {
		return nil, err
	}
	if err := chromedp.Evaluate(`window.scrollTo(0,0)`, nil).Do(ctx); err != nil {
		return nil, err
	}
	if err := chromedp.Evaluate(frameBoxJS, &box).Do(ctx); err != nil {
		return nil, err
	}
	x := math.Max(0, math.Floor(box.X))
	y := math.Max(0, math.Floor(box.Y))
	w := math.Ceil(box.Width)
	h := math.Ceil(box.Height)
	if x+w > float64(vw) {
		w = float64(vw) - x
	}
	if y+h > float64(vh) {
		h = float64(vh) - y
	}
	if w < 1 || h < 1 {
		return nil, fmt.Errorf("clip empty")
	}
	buf, err := page.CaptureScreenshot().
		WithFormat(page.CaptureScreenshotFormatPng).
		WithCaptureBeyondViewport(false).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	return cropPNG(buf, x, y, w, h, browser.MockupShotScale)
}

func cropPNG(src []byte, x, y, w, h, scale float64) ([]byte, error) {
	im, err := png.Decode(bytes.NewReader(src))
	if err != nil {
		return nil, err
	}
	b := im.Bounds()
	r := image.Rect(
		int(math.Round(x*scale)),
		int(math.Round(y*scale)),
		int(math.Round((x+w)*scale)),
		int(math.Round((y+h)*scale)),
	).Intersect(b)
	if r.Empty() {
		return nil, fmt.Errorf("crop empty")
	}
	dst := image.NewNRGBA(image.Rect(0, 0, r.Dx(), r.Dy()))
	draw.Draw(dst, dst.Bounds(), im, r.Min, draw.Src)
	var buf bytes.Buffer
	if err := png.Encode(&buf, dst); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
