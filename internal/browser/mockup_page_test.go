package browser_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/chromedp/chromedp"

	"carrel.dev/typecase/internal/browser"
)

// TestMockupLadderSwitch checks that the width switch above the contents
// reaches every frame.
//
// Frame widths are page classes (.frame.w392, .frame.cap900). The switch rule
// body:has(#wN:checked) .frame outranks them, so picking a step stretches the
// frames that are drawn narrower. An inline style= on .frame did not lose that
// contest, which is why frame-no-style now covers .frame itself.
func TestMockupLadderSwitch(t *testing.T) {
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if testing.Short() {
		t.Skip("needs Chrome")
	}
	root := browser.ModuleRoot(t)
	html, err := browser.LoadMockupHTML(root, "")
	if err != nil {
		t.Fatal(err)
	}

	h := browser.NewHarness(t)
	run, cancel := h.Tab(3 * time.Minute)
	defer cancel()

	url := h.RegisterPage("mockup-page", string(html))
	if err := run(
		chromedp.EmulateViewport(1500, 900),
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
	); err != nil {
		t.Fatal(err)
	}

	// What Х-5 fixed: a frame capped by its own rule. With a step picked, the
	// switch rule (body:has(#wN:checked) .frame) must win, so no frame may be
	// left with a fixed max-width. Frames that merely share a flex row with a
	// neighbour are narrow for a reason the switch has nothing to do with, so
	// width alone does not answer the question — the cap does.
	const capsJS = `(function(){
  var out = [];
  document.querySelectorAll(".frame[data-frame]").forEach(function(f){
    var mw = getComputedStyle(f).maxWidth;
    if (/^[0-9.]+px$/.test(mw)) out.push(f.getAttribute("data-frame") + " max-width:" + mw);
  });
  return out;
})()`

	pick := func(id string) {
		t.Helper()
		if err := run(chromedp.Evaluate(fmt.Sprintf(`(function(){
  var el = document.getElementById(%q);
  el.checked = true;
  el.dispatchEvent(new Event("change", {bubbles:true}));
  return true;
})()`, id), nil)); err != nil {
			t.Fatal(err)
		}
	}
	caps := func() []string {
		t.Helper()
		var got []string
		if err := run(chromedp.Evaluate(capsJS, &got)); err != nil {
			t.Fatal(err)
		}
		return got
	}

	pick("wauto")
	drawn := caps()
	if len(drawn) == 0 {
		t.Fatal("no frame is drawn with a max-width cap; the check proves nothing")
	}
	for _, step := range []string{"w1351", "w820", "w390"} {
		pick(step)
		if left := caps(); len(left) > 0 {
			t.Errorf("%s: the switch does not reach %d frames: %s", step, len(left), strings.Join(left, "; "))
		}
	}
	t.Logf("%d frames capped as drawn, none capped on a step", len(drawn))
}
