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
				if len(probe.Issues) > 0 {
					t.Errorf("%s @%d: %s", fr.ID, w, strings.Join(probe.Issues, "; "))
				}
				if probe.Mode == "drawer" && len(probe.Taps) > 0 {
					t.Errorf("%s @%d tap <44px: %s", fr.ID, w, strings.Join(probe.Taps, "; "))
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
