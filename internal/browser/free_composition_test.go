package browser_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/chromedp/chromedp"

	"carrel.dev/typecase/internal/browser"
	"carrel.dev/typecase/internal/fragment"
	"carrel.dev/typecase/internal/gate"
	"carrel.dev/typecase/internal/refs"
)

// compositionPending are the compositions whose findings wait for work named
// here, the list to go to nothing as ladderPending: what they show is logged,
// not failed. The key is the composition as Composition.String prints it,
// without the host: «Owner.Slot+Part».
var compositionPending = map[string]string{
	// the top bar: a drop at its end opens past the edge, a button more is cut on the phone
	"Top.TopAction+TopDrop":           topBar,
	"TopDrop.TopDropTrigger+UserMenu": topBar,
	"Top.TopAction+Button":            topBar,
	"TopDrop.TopDropTrigger+Button":   topBar,
	// rows of another kind, or a rubric with another gutter at 640
	"List.ListRow+LinkedRow": rowKinds,
	"Main.MainPart+List":     rowKinds,
	"Main.MainPart+Table":    rowKinds,
	// a table's cards in a sheet's pick list beside its rubric
	"SheetPick.SheetPickList+Table": rowKinds,
	// a second target inside a row on the phone: its hit zone takes the
	// row's words, or the zone of the name above it in a cell
	"TableCell.CellPart+CellWrap": linkInRow,
	"PhoneRow.Inline+Link":        linkInRow,
	"SizeRow.Inline+Link":         linkInRow,
	"Drop.DropTrigger+TagInput":   linkInRow,
}

const (
	topBar    = "free composition 10.11: the top bar's end and the phone"
	rowKinds  = "free composition 10.11: rows of different kinds side by side, gutters at 640"
	linkInRow = "free composition 10.11: a second target inside a row on the phone"
)

// TestFreeComposition puts every part a slot takes into a place of that slot
// — the owner's thought of 14.10 on the constructor: the library should fold
// into any screen and not break — and measures the host's frame at every
// width of the ladder: the ladder's probe (scroll, overlap, cut text, a widget
// holding its own, taps on the phone) and the geometry's (rows under their
// heading, nothing out of its box or squeezed in a bar). What the host frame
// shows with its own record in is not the composition's and is left out.
//
// TC_COMPOSE=a,b keeps the compositions whose name holds one of the words;
// TC_COMPOSE_CHUNK=i/n runs the i-th of n equal parts; TC_COMPOSE_REPORT=path
// writes the findings there instead of failing.
func TestFreeComposition(t *testing.T) {
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if os.Getenv("CI") == "" && testing.Short() {
		t.Skip("free composition needs Chrome; skip with -short outside CI")
	}
	html, err := browser.LoadMockupHTML(browser.ModuleRoot(t), "")
	if err != nil {
		t.Fatal(err)
	}
	parts, err := gate.ParseMockupParts(string(html))
	if err != nil {
		t.Fatal(err)
	}
	frames := map[string]string{}
	for _, f := range parts.Frames {
		frames[f.ID] = f.HTML
	}
	comps := refs.Compositions(func(e refs.Entry) bool {
		_, a := ladderPending[e.Ref.Frame]
		_, b := geometryPending[e.Ref.Frame]
		return a || b
	})
	comps = pickCompositions(t, comps)

	h := browser.NewHarness(t)
	lines, url := composeFindings(t, h, parts.Styles, parts.Sprite, frames, comps)
	if shots := os.Getenv("TC_COMPOSE_SHOTS"); shots != "" {
		shootCompositions(t, h, url, frames, comps, shots)
	}
	if report := os.Getenv("TC_COMPOSE_REPORT"); report != "" {
		if err := os.WriteFile(report, []byte(strings.Join(lines, "\n")+"\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		t.Logf("%d findings written to %s", len(lines), report)
		return
	}
	for _, l := range lines {
		name, _, _ := strings.Cut(l, " in ")
		if why, ok := compositionPending[name]; ok {
			t.Logf("pending: %s (%s)", l, why)
			continue
		}
		t.Error(l)
	}
}

// composeFindings measures each composition in its host's frame, less what
// the frame shows with the host's own record: one line a finding,
// «composition<TAB>width<TAB>kind<TAB>what». It returns the shell's url too.
func composeFindings(t *testing.T, h *browser.Harness, styles, sprite string, frames map[string]string, comps []refs.Composition) ([]string, string) {
	t.Helper()
	run, cancel := h.Tab(50 * time.Minute)
	defer cancel()
	url := h.RegisterPage("compose", browser.MockupShotShell(styles, sprite))
	if err := run(
		chromedp.EmulateViewport(int64(browser.MockupShotViewport), int64(browser.MockupShotMaxHeight)),
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
	); err != nil {
		t.Fatal(err)
	}
	fit, err := browser.FitJS()
	if err != nil {
		t.Fatal(err)
	}
	base := map[fragment.Ref]map[string]bool{}
	var lines []string
	for _, c := range comps {
		fr, ok := frames[c.Host.Ref.Frame]
		if !ok {
			t.Fatalf("%s: no frame", c)
		}
		if _, ok := base[c.Host.Ref]; !ok {
			own, err := fragment.Splice(fr, c.Host.Ref, c.Host.Fixture.Render())
			if err != nil {
				t.Fatalf("%s: %v", c.Host.Ref, err)
			}
			seen := map[string]bool{}
			for _, l := range measureComposed(t, run, own, fit) {
				seen[l] = true
			}
			base[c.Host.Ref] = seen
		}
		rendered, err := renderSafe(c.Fixture)
		if err != nil {
			lines = append(lines, fmt.Sprintf("%s\t-\tpanic\t%v", c, err))
			continue
		}
		composed, err := fragment.Splice(fr, c.Host.Ref, rendered)
		if err != nil {
			lines = append(lines, fmt.Sprintf("%s\t-\tsplice\t%v", c, err))
			continue
		}
		for _, l := range measureComposed(t, run, composed, fit) {
			if !base[c.Host.Ref][l] {
				lines = append(lines, c.String()+"\t"+l)
			}
		}
	}
	t.Logf("%d compositions in %d hosts", len(comps), len(base))
	return lines, url
}

// TestFreeCompositionBite: the photo's side column that did not give way
// (flex:0 0 auto, before 10.11) is found again with a message put beside the
// photo, and the same composition is clean with the rule as it is.
func TestFreeCompositionBite(t *testing.T) {
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if os.Getenv("CI") == "" && testing.Short() {
		t.Skip("needs Chrome; skip with -short outside CI")
	}
	html, err := browser.LoadMockupHTML(browser.ModuleRoot(t), "")
	if err != nil {
		t.Fatal(err)
	}
	parts, err := gate.ParseMockupParts(string(html))
	if err != nil {
		t.Fatal(err)
	}
	frames := map[string]string{}
	for _, f := range parts.Frames {
		frames[f.ID] = f.HTML
	}
	var comps []refs.Composition
	for _, c := range refs.Compositions(nil) {
		if c.Owner == "Photo" && c.Slot == "FieldPart" && c.Part == "Msg" {
			comps = append(comps, c)
		}
	}
	if len(comps) != 1 {
		t.Fatalf("want one Photo.FieldPart+Msg, got %d", len(comps))
	}
	h := browser.NewHarness(t)
	clean, _ := composeFindings(t, h, parts.Styles, parts.Sprite, frames, comps)
	if len(clean) > 0 {
		t.Errorf("clean: found %s", strings.Join(clean, "\n"))
	}
	old := parts.Styles + "<style>.tc-field--fixauto.tc-field--fixauto{flex:0 0 auto;min-width:auto;max-width:none}</style>"
	bitten, _ := composeFindings(t, h, old, parts.Sprite, frames, comps)
	if !strings.Contains(strings.Join(bitten, "\n"), "\tout\t") {
		t.Errorf("the old rule: want a finding out of the box, got %q", bitten)
	}
}

// shootCompositions takes pictures of each composition, and of its host
// with its own record for «before», at the widths spec names
// («1351,460»), into TC_COMPOSE_SHOTS_DIR.
func shootCompositions(t *testing.T, h *browser.Harness, url string, frames map[string]string, comps []refs.Composition, spec string) {
	t.Helper()
	dir := os.Getenv("TC_COMPOSE_SHOTS_DIR")
	if dir == "" {
		dir = filepath.Join(os.TempDir(), "tc-compose-shots")
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	fit, err := browser.FitJS()
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range comps {
		fr := frames[c.Host.Ref.Frame]
		rendered, err := renderSafe(c.Fixture)
		if err != nil {
			continue
		}
		before, err1 := fragment.Splice(fr, c.Host.Ref, c.Host.Fixture.Render())
		after, err2 := fragment.Splice(fr, c.Host.Ref, rendered)
		if err1 != nil || err2 != nil {
			continue
		}
		name := strings.NewReplacer(".", "-", "+", "-").Replace(c.Owner + "." + c.Slot + "+" + c.Part)
		for _, ws := range strings.Split(spec, ",") {
			w, err := strconv.Atoi(strings.TrimSpace(ws))
			if err != nil {
				t.Fatalf("TC_COMPOSE_SHOTS=%q: %v", spec, err)
			}
			for _, side := range []string{"b", "a"} {
				html := map[string]string{"b": before, "a": after}[side]
				raw, _ := json.Marshal(html)
				// a tab that took one picture may hang on the next: each picture
				// is taken in a tab of its own
				run, cancel := h.Tab(2 * time.Minute)
				err := run(
					chromedp.EmulateViewport(int64(browser.MockupShotViewport), int64(browser.MockupShotMaxHeight)),
					chromedp.Navigate(url),
					chromedp.WaitReady("body"),
					chromedp.Evaluate(fontsReadyJS, nil),
				)
				var buf []byte
				if err == nil {
					buf, err = shootFrame(run, string(raw), w, fit)
				}
				cancel()
				if err != nil {
					t.Fatalf("%s @%d: %v", c, w, err)
				}
				path := filepath.Join(dir, fmt.Sprintf("%s@%d-%s.png", name, w, side))
				if err := os.WriteFile(path, buf, 0o644); err != nil {
					t.Fatal(err)
				}
			}
		}
		t.Logf("shot %s", c)
	}
}

// pickCompositions keeps what TC_COMPOSE and TC_COMPOSE_CHUNK ask for.
func pickCompositions(t *testing.T, comps []refs.Composition) []refs.Composition {
	t.Helper()
	if words := os.Getenv("TC_COMPOSE"); words != "" {
		var kept []refs.Composition
		for _, c := range comps {
			for _, w := range strings.Split(words, ",") {
				if w != "" && strings.Contains(c.String(), w) {
					kept = append(kept, c)
					break
				}
			}
		}
		comps = kept
	}
	if chunk := os.Getenv("TC_COMPOSE_CHUNK"); chunk != "" {
		a, b, _ := strings.Cut(chunk, "/")
		i, err1 := strconv.Atoi(a)
		n, err2 := strconv.Atoi(b)
		if err1 != nil || err2 != nil || n < 1 || i < 1 || i > n {
			t.Fatalf("TC_COMPOSE_CHUNK=%q: want i/n", chunk)
		}
		from, to := len(comps)*(i-1)/n, len(comps)*i/n
		comps = comps[from:to]
	}
	return comps
}

func renderSafe(r refs.Renderer) (out string, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("%v", p)
		}
	}()
	return r.Render(), nil
}

// measureComposed paints a frame at every width of the ladder, folds its bars
// as tc.js does (fit, FitJS: a screen runs the script) and runs both
// probes: one line a finding, «width<TAB>kind<TAB>what».
func measureComposed(t *testing.T, run func(...chromedp.Action) error, frame, fit string) []string {
	t.Helper()
	raw, err := json.Marshal(frame)
	if err != nil {
		t.Fatal(err)
	}
	var out []string
	for _, w := range browser.MockupLadderWidths {
		var ladder ladderProbe
		var geo geometryProbe
		if err := run(
			chromedp.Evaluate(paintFrameJS(string(raw), w), nil),
			chromedp.Evaluate(fit, nil),
			chromedp.Sleep(20*time.Millisecond),
			chromedp.Evaluate(browser.MockupLadderProbeJS, &ladder),
			chromedp.Evaluate(browser.MockupGeometryProbeJS, &geo),
		); err != nil {
			t.Fatalf("@%d: %v", w, err)
		}
		for _, is := range ladder.Issues {
			out = append(out, fmt.Sprintf("%d\tladder\t%s", w, is))
		}
		if ladder.Mode == "drawer" {
			for _, tp := range ladder.Taps {
				out = append(out, fmt.Sprintf("%d\ttap\t%s", w, tp))
			}
		}
		for _, r := range geo.Rows {
			out = append(out, fmt.Sprintf("%d\trow\t%s", w, r))
		}
		for _, o := range geo.Out {
			out = append(out, fmt.Sprintf("%d\tout\t%s", w, o))
		}
		for _, q := range geo.Squeezed {
			out = append(out, fmt.Sprintf("%d\tsqueezed\t%s", w, q))
		}
	}
	return out
}

// paintFrameJS puts a frame (its markup as a JSON string) into the shell's
// host and sets its drawing to width w, as the ladder does.
func paintFrameJS(raw string, w int) string {
	return fmt.Sprintf(`(async function(){
  var host = document.getElementById("host");
  host.innerHTML = %s;
  var f = host.querySelector(".frame");
  if (!f) return false;
  f.style.width = "max-content";
  f.style.maxWidth = "none";
  f.style.margin = "0";
  var root = f.querySelector(".tc-root");
  if (root) {
    root.style.width = %d + "px";
    root.style.maxWidth = %d + "px";
    root.style.flex = "none";
  }
  return true;
})()`, raw, w, w)
}
