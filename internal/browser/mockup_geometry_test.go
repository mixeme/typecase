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

// geometryPending are the frames whose geometry waits for work named here,
// the list to go to nothing as ladderPending: what they show is logged, not
// failed.
var geometryPending = map[string]string{
	"s-mailbox-14": "the mailbox is approved and put off (30.09): its frames wait",
}

type geometryProbe struct {
	Rows     []string `json:"rows"`
	Out      []string `json:"out"`
	Squeezed []string `json:"squeezed"`
}

// TestMockupGeometry measures every frame at every width of the ladder:
// rows under their heading, boxes inside their parents (the owner, 07.11:
// the geometry check first). TC_GEO_REPORT=path writes what it finds there
// instead of failing — the survey before the rules are set; TC_GEO_REV=HEAD
// measures a commit's mockup instead of the working tree's.
func TestMockupGeometry(t *testing.T) {
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if os.Getenv("CI") == "" && testing.Short() {
		t.Skip("the geometry check needs Chrome; skip with -short outside CI")
	}
	report := os.Getenv("TC_GEO_REPORT")
	root := browser.ModuleRoot(t)
	html, err := browser.LoadMockupHTML(root, os.Getenv("TC_GEO_REV"))
	if err != nil {
		t.Fatal(err)
	}
	parts, err := gate.ParseMockupParts(string(html))
	if err != nil {
		t.Fatal(err)
	}
	only := map[string]bool{}
	for _, id := range strings.Split(os.Getenv("TC_FRAMES"), ",") {
		if id != "" {
			only[id] = true
		}
	}
	var frames []gate.MockupFrame
	for _, fr := range parts.Frames {
		if len(only) == 0 || only[fr.ID] {
			frames = append(frames, fr)
		}
	}
	lines := measureGeometry(t, parts.Styles, parts.Sprite, frames, browser.MockupLadderWidths)
	if report != "" {
		if err := os.WriteFile(report, []byte(strings.Join(lines, "\n")+"\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		t.Logf("%d findings written to %s", len(lines), report)
		return
	}
	for _, l := range lines {
		id, _, _ := strings.Cut(l, "\t")
		if why, ok := geometryPending[id]; ok {
			t.Logf("pending: %s (%s)", l, why)
			continue
		}
		t.Error(l)
	}
}

// measureGeometry paints each frame at each width on one page and runs the
// probe: one line a finding, «id<TAB>width<TAB>row|out<TAB>what».
func measureGeometry(t *testing.T, styles, sprite string, frames []gate.MockupFrame, widths []int) []string {
	t.Helper()
	h := browser.NewHarness(t)
	run, cancel := h.Tab(50 * time.Minute)
	defer cancel()
	url := h.RegisterPage("geometry", browser.MockupShotShell(styles, sprite))
	if err := run(
		chromedp.EmulateViewport(int64(browser.MockupShotViewport), int64(browser.MockupShotMaxHeight)),
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
	); err != nil {
		t.Fatal(err)
	}
	var lines []string
	for _, fr := range frames {
		raw, err := json.Marshal(fr.HTML)
		if err != nil {
			t.Fatal(err)
		}
		for _, w := range widths {
			js := fmt.Sprintf(`(async function(){
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
			var probe geometryProbe
			if err := run(chromedp.Evaluate(js, nil), chromedp.Sleep(20*time.Millisecond), chromedp.Evaluate(browser.MockupGeometryProbeJS, &probe)); err != nil {
				t.Fatalf("%s @%d: %v", fr.ID, w, err)
			}
			for _, r := range probe.Rows {
				lines = append(lines, fmt.Sprintf("%s\t%d\trow\t%s", fr.ID, w, r))
			}
			for _, o := range probe.Out {
				lines = append(lines, fmt.Sprintf("%s\t%d\tout\t%s", fr.ID, w, o))
			}
			for _, q := range probe.Squeezed {
				lines = append(lines, fmt.Sprintf("%s\t%d\tsqueezed\t%s", fr.ID, w, q))
			}
		}
	}
	return lines
}

// TestMockupGeometryBite: faults of each kind the check was made for, put
// back into frames that are clean — the section list's second gutter (row),
// the card's header that does not wrap (out), the document bar that keeps its
// third rank and squeezes the notebook pick (squeezed) — are found, and the
// clean frames are not.
func TestMockupGeometryBite(t *testing.T) {
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
	byID := map[string]gate.MockupFrame{}
	for _, f := range parts.Frames {
		byID[f.ID] = f
	}
	for _, c := range []struct {
		name, frame, css, want string
		width                  int
	}{
		{"clean rows", "s-contacts-6", "", "", 1351},
		{"a second gutter", "s-contacts-6", ".tc-sec > .tc-list > .tc-row{padding-left:20px}", "\trow\t20 ", 1351},
		{"a clean card", "s-profile-1", "", "", 641},
		{"a header that does not wrap", "s-profile-1", ".tc-card > .tc-table th{white-space:nowrap}", "\tout\t", 641},
		{"a clean bar", "s-notes-20", "", "", 641},
		{"a bar that keeps its third rank", "s-notes-20", ".tc-docbar > .tc-btn.is-3rd,.tc-docbar > .tc-bar__right > .tc-check.is-3rd{display:inline-flex}", "\tsqueezed\t", 641},
	} {
		fr, ok := byID[c.frame]
		if !ok {
			t.Fatalf("%s: no frame %s", c.name, c.frame)
		}
		styles := parts.Styles
		if c.css != "" {
			styles += "<style>" + c.css + "</style>"
		}
		lines := measureGeometry(t, styles, parts.Sprite, []gate.MockupFrame{fr}, []int{c.width})
		got := strings.Join(lines, "\n")
		switch {
		case c.want == "" && len(lines) > 0:
			t.Errorf("%s: found %s", c.name, got)
		case c.want != "" && !strings.Contains(got, c.want):
			t.Errorf("%s: want a finding with %q, got %q", c.name, c.want, got)
		default:
			t.Logf("%s: %q", c.name, got)
		}
	}
}
