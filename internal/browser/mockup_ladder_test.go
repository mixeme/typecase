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

func TestMockupLadderWidths(t *testing.T) {
	want := []int{1351, 1191, 1190, 1008, 1007, 820, 641, 640, 460, 390}
	got := browser.MockupLadderWidths
	if len(got) != len(want) {
		t.Fatalf("len %d want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("%d: %d want %d", i, got[i], want[i])
		}
	}
}

type ladderProbe struct {
	Issues []string `json:"issues"`
	Taps   []string `json:"taps"`
	Mode   string   `json:"mode"`
}

// ladderPending are the frames whose widgets are still cut or squeezed
// (Р-65) and why: their «cut» and «holds» are logged, not failed, until the
// work named here is done — the list must go to nothing. Anything else in
// them still fails.
var ladderPending = map[string]string{}

// tapPending are the frames whose taps miss for a reason named here, the
// list to go to nothing as ladderPending.
var tapPending = map[string]string{}

func TestMockupLadder(t *testing.T) {
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if os.Getenv("CI") == "" && testing.Short() {
		t.Skip("mockup ladder needs Chrome; skip with -short outside CI")
	}

	root := browser.ModuleRoot(t)
	html, err := browser.LoadMockupHTML(root, "")
	if err != nil {
		t.Fatal(err)
	}
	if err := gate.CheckMockupLint(string(html)); err != nil {
		t.Fatalf("current mockup: %v", err)
	}
	parts, err := gate.ParseMockupParts(string(html))
	if err != nil {
		t.Fatal(err)
	}

	h := browser.NewHarness(t)
	run, cancel := h.Tab(40 * time.Minute)
	defer cancel()

	url := h.RegisterPage("ladder", browser.MockupShotShell(parts.Styles, parts.Sprite))
	if err := run(
		chromedp.EmulateViewport(int64(browser.MockupShotViewport), int64(browser.MockupShotMaxHeight)),
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
	); err != nil {
		t.Fatal(err)
	}

	for _, fr := range parts.Frames {
		rawHTML, err := json.Marshal(fr.HTML)
		if err != nil {
			t.Fatal(err)
		}
		t.Run(fr.ID, func(t *testing.T) {
			for _, w := range browser.MockupLadderWidths {
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
})()`, rawHTML, w, w)
				if err := run(chromedp.Evaluate(js, nil), chromedp.Sleep(20*time.Millisecond)); err != nil {
					t.Fatalf("%s @%d inject: %v", fr.ID, w, err)
				}
				var probe ladderProbe
				if err := run(chromedp.Evaluate(browser.MockupLadderProbeJS, &probe)); err != nil {
					t.Fatalf("%s @%d probe: %v", fr.ID, w, err)
				}
				var issues, pending []string
				for _, is := range probe.Issues {
					if why, ok := ladderPending[fr.ID]; ok && (strings.HasPrefix(is, "cut ") || strings.HasPrefix(is, "holds ")) {
						pending = append(pending, is+" ("+why+")")
						continue
					}
					issues = append(issues, is)
				}
				if len(issues) > 0 {
					t.Errorf("%s @%d: %s", fr.ID, w, strings.Join(issues, "; "))
				}
				if len(pending) > 0 {
					t.Logf("%s @%d pending: %s", fr.ID, w, strings.Join(pending, "; "))
				}
				if probe.Mode == "drawer" && len(probe.Taps) > 0 {
					if why, ok := tapPending[fr.ID]; ok {
						t.Logf("%s @%d tap pending: %s (%s)", fr.ID, w, strings.Join(probe.Taps, "; "), why)
					} else {
						t.Errorf("%s @%d tap <44px: %s", fr.ID, w, strings.Join(probe.Taps, "; "))
					}
				}
				if w <= 640 && strings.Contains(fr.HTML, "tc-shell__body") && probe.Mode != "drawer" {
					t.Errorf("%s @%d: --tc-shell-mode=%q want drawer", fr.ID, w, probe.Mode)
				}
				if w >= 641 && w <= 1007 && strings.Contains(fr.HTML, "tc-shell__rail") && probe.Mode != "rail" {
					t.Errorf("%s @%d: --tc-shell-mode=%q want rail", fr.ID, w, probe.Mode)
				}
				if w >= 1008 && strings.Contains(fr.HTML, "tc-shell__body") && probe.Mode != "full" {
					t.Errorf("%s @%d: --tc-shell-mode=%q want full", fr.ID, w, probe.Mode)
				}
			}
		})
	}
}
