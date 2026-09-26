package browser_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/chromedp/chromedp"

	"carrel.dev/typecase"
	"carrel.dev/typecase/internal/browser"
	"carrel.dev/typecase/internal/gate"
	"carrel.dev/typecase/internal/screens"
)

// framesPage is the page both trees are laid out on: the mockup's CSS and
// sprite, the frame in #host and the screen in #screenhost.
func framesPage(t *testing.T) (h *browser.Harness, run func(...chromedp.Action) error, done func(), frames map[string]string) {
	t.Helper()
	root := browser.ModuleRoot(t)
	raw, err := browser.LoadMockupHTML(root, "")
	if err != nil {
		t.Fatal(err)
	}
	parts, err := gate.ParseMockupParts(string(raw))
	if err != nil {
		t.Fatal(err)
	}
	frames = map[string]string{}
	for _, f := range parts.Frames {
		frames[f.ID] = f.HTML
	}
	h = browser.NewHarness(t)
	run, done = h.Tab(90 * time.Minute)
	url := h.RegisterPage("frames", browser.MockupShotShell(parts.Styles, parts.Sprite))
	if err := run(chromedp.EmulateViewport(1351, 900), chromedp.Navigate(url), chromedp.WaitReady("body"), chromedp.Evaluate(fontsReadyJS, nil)); err != nil {
		done()
		t.Fatal(err)
	}
	return h, run, done, frames
}

// compareFrame lays a frame and its screen out at width w and returns what
// differs and which invariants of the screen fail.
func compareFrame(run func(...chromedp.Action) error, frameHTML, screenHTML string, w, max int) (diffs, broken []string, err error) {
	return compareFrameAfter(run, frameHTML, screenHTML, w, max, "0")
}

// compareFrameAfter is compareFrame with a script run on the page before the
// walk: the bites spoil the screen with it.
func compareFrameAfter(run func(...chromedp.Action) error, frameHTML, screenHTML string, w, max int, spoil string) (diffs, broken []string, err error) {
	fj, _ := json.Marshal(frameHTML)
	sj, _ := json.Marshal(screenHTML)
	var ok bool
	err = run(
		chromedp.EmulateViewport(int64(w), 900),
		chromedp.Evaluate(fmt.Sprintf("(%s)(%s, %s, %d)", browser.FramesPairJS, fj, sj, w), &ok),
		chromedp.Evaluate(spoil, nil),
		chromedp.Evaluate(fmt.Sprintf("(%s)(%d)", browser.FramesWalkJS, max), &diffs),
		chromedp.Evaluate(fmt.Sprintf("(%s)(%d)", browser.FramesInvariantsJS, max), &broken),
	)
	if err == nil && !ok {
		err = fmt.Errorf("the frame or the screen has no .tc-root")
	}
	return diffs, broken, err
}

// framesAwaiting are the breakages of the mockup itself the ladder found
// on the narrow steps: the screen is its frame there too, so the fix is in
// the mockup (Р-3, Р-4, Р-21). Each is keyed by frame, step and the start of
// the finding; the run fails on a finding not listed and on a listed one
// that is gone.
var framesAwaiting = map[string]string{
	"s-find-2@390: text out of its box: button.tc-menu__item":            "the search's suggestions are squeezed under the phone's search button",
	"s-find-3@390: text out of its box: button.tc-menu__item":            "the search's suggestions are squeezed under the phone's search button",
	"s-conflict-2@640: text out of its box: button.tc-btn.is-primary":    "«New note» is cut to its box on a phone's top bar",
	"s-files-4@390: under something: div.tc-sheet":                       "the page's buttons are drawn over the move sheet",
	"s-publish-5@390: under something: div.tc-sheet":                     "the page is drawn over the sheet",
	"account-6@390: under something: div.tc-sheet":                       "the page's primary button is drawn over the sheet",
	"s-admin-5@390: text out of its box: div.tc-diag":                    "a diagnostic line does not wrap on a phone",
	"s-profile-1@390: text out of its box: div.tc-diag":                  "a diagnostic line does not wrap on a phone",
	"s-profile-10@390: text out of its box: div.tc-diag":                 "a diagnostic line does not wrap on a phone",
	"s-backup-2@640: text out of its box: div.tc-diag.tc-diag--top9-ink": "the line to paste does not wrap on a phone",
}

func awaited(frame string, w int, finding string) (string, bool) {
	for k := range framesAwaiting {
		at, prefix, _ := strings.Cut(k, ": ")
		if at == fmt.Sprintf("%s@%d", frame, w) && strings.HasPrefix(finding, prefix) {
			return k, true
		}
	}
	return "", false
}

// TestFramesLadder (plan 3, step 3): every route and its frame, laid out on
// every step of the ladder, draw the same boxes with the same computed style;
// and the screen keeps the invariants of a screen on each step.
// TC_FRAMES=<frame>,… limits the run to those frames.
func TestFramesLadder(t *testing.T) {
	needChrome(t)
	_, run, done, frames := framesPage(t)
	defer done()
	only := map[string]bool{}
	for _, id := range strings.Split(os.Getenv("TC_FRAMES"), ",") {
		if id = strings.TrimSpace(id); id != "" {
			only[id] = true
		}
	}
	seen := map[string]bool{}
	for _, r := range screens.Routes {
		if len(only) > 0 && !only[r.Frame] {
			continue
		}
		fh, ok := frames[r.Frame]
		if !ok {
			t.Errorf("%s: no such frame", r.Frame)
			continue
		}
		screen := r.Screen().Render()
		for _, w := range browser.MockupLadderWidths {
			diffs, broken, err := compareFrame(run, fh, screen, w, 8)
			if err != nil {
				t.Fatalf("%s @%d: %v", r.Frame, w, err)
			}
			for _, d := range diffs {
				t.Errorf("%s @%d: %s", r.Frame, w, d)
			}
			for _, b := range broken {
				if k, ok := awaited(r.Frame, w, b); ok {
					seen[k] = true
					continue
				}
				t.Errorf("%s @%d: %s", r.Frame, w, b)
			}
		}
	}
	if len(only) == 0 {
		for k, why := range framesAwaiting {
			if !seen[k] {
				t.Errorf("%s (%s) is fixed: take it off framesAwaiting", k, why)
			}
		}
	}
}

// TestFramesBites (L-9.3 for step 3): the ladder comparison is red, with the
// path to the node, when one token, one rule or the order of two widgets
// moves on the screen side only.
func TestFramesBites(t *testing.T) {
	needChrome(t)
	_, run, done, frames := framesPage(t)
	defer done()
	r, ok := screens.Find("s-contacts-1")
	if !ok {
		t.Fatal("no route s-contacts-1")
	}
	screen := r.Screen().Render()
	if diffs, _, err := compareFrame(run, frames[r.Frame], screen, 1351, 8); err != nil || len(diffs) > 0 {
		t.Fatalf("the unspoilt pair differs: %v %v", err, diffs)
	}
	for name, c := range map[string]struct{ spoil, says string }{
		"token": {`document.querySelector("#screenhost > .frame > .tc-root").style.setProperty("--tc-accent", "#000")`, "color"},
		"rule":  {`(() => { const st = document.createElement("style"); st.textContent = "#screenhost .tc-bar .tc-btn{padding-left:12px}"; document.head.appendChild(st); })()`, "tc-bar"},
		"order": {`(() => { const a = document.querySelector("#screenhost .tc-top__actions"); a.insertBefore(a.children[1], a.children[0]); })()`, "tc-top__actions"},
	} {
		diffs, _, err := compareFrameAfter(run, frames[r.Frame], screen, 1351, 8, c.spoil)
		if err != nil {
			t.Fatal(err)
		}
		// Each bite leaves its sheet behind: the next one starts from a clean page.
		_ = run(chromedp.Evaluate(`document.querySelectorAll("head > style:not(:first-of-type)").forEach((s) => { if (s.textContent.includes("#screenhost")) s.remove(); })`, nil))
		if len(diffs) == 0 || !strings.Contains(strings.Join(diffs, "\n"), c.says) {
			t.Errorf("%s: the comparison did not bite (%v)", name, diffs)
		}
	}
}

// longLabelsJS doubles the words of every button in the screen's bars, edit
// bars and heads — labels longer than the frames draw — and lets tc.js lay
// the bars out; it returns how many bars its safety caught.
const longLabelsJS = `(() => {
  const root = document.querySelector("#screenhost > .frame > .tc-root");
  root.querySelectorAll(".tc-bar .tc-btn, .tc-docbar .tc-btn, .tc-head .tc-btn").forEach((b) => {
    b.childNodes.forEach((n) => { if (n.nodeType === 3 && n.textContent.trim()) n.textContent = n.textContent + " " + n.textContent.trim(); });
    b.querySelectorAll(".tc-btn__label").forEach((l) => { l.textContent = l.textContent + " " + l.textContent; });
  });
  window.__tc.layoutUpdates();
  return root.querySelectorAll("[data-tc-overflow]").length;
})()`

// TestFramesLongLabels (plan 3, step 3): with labels longer than the frames
// draw, the CSS cannot know their length; the safety of tc.js catches the
// bars they overflow (at least somewhere), and nothing else scrolls sideways.
func TestFramesLongLabels(t *testing.T) {
	needChrome(t)
	_, run, done, frames := framesPage(t)
	defer done()
	js, err := typecase.JS()
	if err != nil {
		t.Fatal(err)
	}
	if err := run(chromedp.Evaluate(string(js), nil)); err != nil {
		t.Fatal(err)
	}
	caught := 0
	for _, r := range screens.Routes {
		fh := frames[r.Frame]
		screen := r.Screen().Render()
		for _, w := range browser.MockupLadderWidths {
			var n int
			var bad []string
			fj, _ := json.Marshal(fh)
			sj, _ := json.Marshal(screen)
			err := run(
				chromedp.EmulateViewport(int64(w), 900),
				chromedp.Evaluate(fmt.Sprintf("(%s)(%s, %s, %d)", browser.FramesPairJS, fj, sj, w), nil),
				chromedp.Evaluate(longLabelsJS, &n),
				chromedp.Evaluate(fmt.Sprintf("(%s)(%d)", browser.FramesInvariantsJS, 8), &bad),
			)
			if err != nil {
				t.Fatalf("%s @%d: %v", r.Frame, w, err)
			}
			caught += n
			for _, b := range bad {
				if strings.HasPrefix(b, "scrolls sideways: ") && !strings.Contains(b, "tc-bar") && !strings.Contains(b, "tc-docbar") && !strings.Contains(b, "tc-head") {
					t.Errorf("%s @%d with long labels: %s", r.Frame, w, b)
				}
			}
		}
	}
	if caught == 0 {
		t.Error("with long labels the safety never caught a bar")
	}
}

// TestFramesTablet (plan 3, step 3, the owner's observation of 2026-09-26):
// the ladder is measured in a desktop Chrome; a tablet also has a mobile
// viewport, a pixel ratio of 2 and touch. Samsung Galaxy Tab S5e in Vivaldi
// (mobile Chromium): about 1280×800 landscape and 800×1280 portrait. Every
// route and its frame are laid out so, and the screen keeps its invariants.
func TestFramesTablet(t *testing.T) {
	needChrome(t)
	_, run, done, frames := framesPage(t)
	defer done()
	for _, r := range screens.Routes {
		fh, sj := frames[r.Frame], r.Screen().Render()
		for _, v := range [][2]int{{1280, 800}, {800, 1280}} {
			fjs, _ := json.Marshal(fh)
			sjs, _ := json.Marshal(sj)
			var diffs, bad []string
			err := run(
				chromedp.EmulateViewport(int64(v[0]), int64(v[1]), chromedp.EmulateScale(2), chromedp.EmulateMobile, chromedp.EmulateTouch),
				chromedp.Evaluate(fmt.Sprintf("(%s)(%s, %s, %d)", browser.FramesPairJS, fjs, sjs, v[0]), nil),
				chromedp.Evaluate(fmt.Sprintf("(%s)(%d)", browser.FramesWalkJS, 8), &diffs),
				chromedp.Evaluate(fmt.Sprintf("(%s)(%d)", browser.FramesInvariantsJS, 8), &bad),
			)
			if err != nil {
				t.Fatalf("%s @%dx%d: %v", r.Frame, v[0], v[1], err)
			}
			for _, d := range append(diffs, bad...) {
				t.Errorf("%s @tablet %dx%d: %s", r.Frame, v[0], v[1], d)
			}
		}
	}
}
