package browser_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/chromedp/chromedp"

	"carrel.dev/typecase/internal/browser"
	"carrel.dev/typecase/internal/gate"
)

// barFitDiffers are the frames and widths where tc.js folds a bar of the
// frame otherwise than the CSS's widths draw it, each with what the owner
// decided: «frame@width».
var barFitDiffers = map[string]string{}

// barStateJS lists every element of every bar in the painted frame with
// whether it shows and how wide: «bar.element class width», «-» for hidden.
const barStateJS = `(() => {
  const out = [];
  document.querySelectorAll("#host .tc-bar, #host .tc-docbar").forEach((bar, b) => {
    bar.querySelectorAll("*").forEach((el, i) => {
      const shown = el.checkVisibility ? el.checkVisibility() : el.getClientRects().length > 0;
      let c = el.getAttribute("class") || el.tagName.toLowerCase();
      c = c.split(" ").filter((x) => x && !x.startsWith("is-on")).slice(0, 2).join(".");
      out.push(b + "." + i + " " + c + " " + (shown ? Math.round(el.getBoundingClientRect().width) : "-"));
    });
  });
  return out;
})()`

// TestBarFitMatchesFrames (Р-4, the owner's decision of 10.11): on the
// mockup's own frames the script folds every bar as the CSS's widths draw it
// — those widths were measured for these very sets — at every width of the
// ladder; where it does not, the frame waits in barFitDiffers for the owner.
// TC_BARFIT_REPORT=path writes the differences instead of failing.
func TestBarFitMatchesFrames(t *testing.T) {
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
	fit, err := browser.FitBarsJS()
	if err != nil {
		t.Fatal(err)
	}
	h := browser.NewHarness(t)
	run, cancel := h.Tab(40 * time.Minute)
	defer cancel()
	url := h.RegisterPage("barfit", browser.MockupShotShell(parts.Styles, parts.Sprite))
	if err := run(
		chromedp.EmulateViewport(int64(browser.MockupShotViewport), int64(browser.MockupShotMaxHeight)),
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
	); err != nil {
		t.Fatal(err)
	}
	var lines, broken []string
	frames := 0
	for _, f := range parts.Frames {
		if !strings.Contains(f.HTML, "tc-bar") && !strings.Contains(f.HTML, "tc-docbar") {
			continue
		}
		if _, ok := ladderPending[f.ID]; ok {
			continue
		}
		raw, err := json.Marshal(f.HTML)
		if err != nil {
			t.Fatal(err)
		}
		counted := false
		for _, w := range browser.MockupLadderWidths {
			var before, after []string
			var n int
			var ladderB, ladderA ladderProbe
			var geoB, geoA geometryProbe
			if err := run(
				chromedp.Evaluate(paintFrameJS(string(raw), w), nil),
				chromedp.Evaluate(barStateJS, &before),
				chromedp.Evaluate(browser.MockupLadderProbeJS, &ladderB),
				chromedp.Evaluate(browser.MockupGeometryProbeJS, &geoB),
				chromedp.Evaluate(fit, &n),
				chromedp.Evaluate(barStateJS, &after),
				chromedp.Evaluate(browser.MockupLadderProbeJS, &ladderA),
				chromedp.Evaluate(browser.MockupGeometryProbeJS, &geoA),
			); err != nil {
				t.Fatalf("%s@%d: %v", f.ID, w, err)
			}
			if n == 0 {
				break
			}
			// what the folded bar shows that the frame does not: the script's
			// own fault, whatever the owner decides on the look
			seen := map[string]bool{}
			for _, l := range probeLines(ladderB, geoB) {
				seen[l] = true
			}
			for _, l := range probeLines(ladderA, geoA) {
				if !seen[l] {
					broken = append(broken, fmt.Sprintf("%s@%d\t%s", f.ID, w, l))
				}
			}
			if !counted {
				frames++
				counted = true
			}
			if os.Getenv("TC_BARFIT_DEBUG") == fmt.Sprintf("%s@%d", f.ID, w) {
				var steps []string
				if err := run(chromedp.Evaluate(`(() => { const out = []; document.querySelectorAll("#host .tc-bar, #host .tc-docbar").forEach((bar) => {
  const keep = bar.getAttribute("data-tc-fold"); bar.setAttribute("data-tc-measure", "");
  for (const s of ["", "end", "end 2nd", "end 2nd labels", "end 2nd labels 3rd"]) {
    bar.setAttribute("data-tc-fold", s);
    const box = bar.getBoundingClientRect();
    const kids = [...bar.querySelectorAll(":scope > *, :scope > .tc-barform > *")].filter((k) => k.getClientRects().length).map((k) => (k.className || k.tagName).toString().split(" ")[0] + "[" + Math.round(k.getBoundingClientRect().left - box.left) + "," + Math.round(k.getBoundingClientRect().right - box.left) + "/" + Math.round(k.getBoundingClientRect().top - box.top) + "]");
    out.push((bar.className.split(" ")[0]) + " w" + Math.round(box.width) + " «" + s + "» " + kids.join(" "));
  }
  bar.removeAttribute("data-tc-measure"); bar.setAttribute("data-tc-fold", keep); }); return out; })()`, &steps)); err != nil {
					t.Fatal(err)
				}
				t.Logf("%s@%d steps:\n%s", f.ID, w, strings.Join(steps, "\n"))
			}
			if d := barStateDiff(before, after); d != "" {
				lines = append(lines, fmt.Sprintf("%s@%d\t%s", f.ID, w, d))
			}
		}
	}
	t.Logf("%d frames with bars", frames)
	for _, l := range broken {
		t.Errorf("the folded bar breaks the frame: %s", l)
	}
	if report := os.Getenv("TC_BARFIT_REPORT"); report != "" {
		if err := os.WriteFile(report, []byte(strings.Join(lines, "\n")+"\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		t.Logf("%d differences written to %s", len(lines), report)
		return
	}
	for _, l := range lines {
		key, _, _ := strings.Cut(l, "\t")
		if why, ok := barFitDiffers[key]; ok {
			t.Logf("decided: %s (%s)", l, why)
			continue
		}
		t.Errorf("the script folds otherwise than the frame: %s", l)
	}
}

// barStateDiff names what the script shows that the frame hides (+), hides
// that the frame shows (−), and draws wider or narrower by more than a pixel.
func barStateDiff(before, after []string) string {
	if len(before) != len(after) {
		return fmt.Sprintf("%d elements, then %d", len(before), len(after))
	}
	var out []string
	for i := range before {
		if before[i] == after[i] {
			continue
		}
		b := strings.Fields(before[i])
		a := strings.Fields(after[i])
		if len(a) < 3 || len(b) < 3 {
			continue
		}
		name := a[1]
		switch {
		case b[2] == "-":
			out = append(out, "+"+name)
		case a[2] == "-":
			out = append(out, "−"+name)
		default:
			var x, y int
			fmt.Sscan(b[2], &x)
			fmt.Sscan(a[2], &y)
			if x-y > 1 || y-x > 1 {
				out = append(out, fmt.Sprintf("%s %d→%d", name, x, y))
			}
		}
	}
	return strings.Join(out, " ")
}

// probeLines are the probes' findings as measureComposed prints them, less
// the width.
func probeLines(l ladderProbe, g geometryProbe) []string {
	var out []string
	for _, is := range l.Issues {
		out = append(out, "ladder	"+is)
	}
	if l.Mode == "drawer" {
		for _, tp := range l.Taps {
			out = append(out, "tap	"+tp)
		}
	}
	for _, r := range g.Rows {
		out = append(out, "row	"+r)
	}
	for _, o := range g.Out {
		out = append(out, "out	"+o)
	}
	for _, q := range g.Squeezed {
		out = append(out, "squeezed	"+q)
	}
	return out
}
