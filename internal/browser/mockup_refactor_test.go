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
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"

	"carrel.dev/typecase/internal/browser"
	"carrel.dev/typecase/internal/gate"
)

const mockupStageSize = 1

func TestMockupShotShellFlattensStyles(t *testing.T) {
	got := browser.MockupShotShell(`<style>a{color:red}</style><style data-tc>b{color:blue}</style>`, "")
	if n := strings.Count(strings.ToLower(got), "<style"); n != 1 {
		t.Fatalf("want one <style>, got %d in %s", n, got)
	}
	if !strings.Contains(got, "a{color:red}") || !strings.Contains(got, "b{color:blue}") {
		t.Fatalf("missing concatenated rules: %s", got)
	}
}

func TestStripFontFaceCSS(t *testing.T) {
	in := "@font-face{font-family:X;src:url(a.ttf)}.a{color:red}@font-face {font-family:Y;src:url(b.ttf)}.b{color:blue}"
	got := browser.StripFontFaceCSS(in)
	if strings.Contains(strings.ToLower(got), "@font-face") {
		t.Fatalf("still has @font-face: %s", got)
	}
	if !strings.Contains(got, ".a{color:red}") || !strings.Contains(got, ".b{color:blue}") {
		t.Fatalf("stripped rules: %s", got)
	}
}

func TestMockupShotShellPair(t *testing.T) {
	got := browser.MockupShotShellPair(`<style>a{color:red}@font-face{font-family:X;src:url(a.ttf)}</style>`, `<style>b{color:blue}@font-face{font-family:Y;src:url(b.ttf)}</style>`, "")
	if !strings.Contains(got, `id="tc-css-after"`) || !strings.Contains(got, `id="tc-css-before"`) {
		t.Fatalf("missing pair ids: %s", got)
	}
	if !strings.Contains(got, `media="not all"`) {
		t.Fatal("baseline sheet must start disabled")
	}
	if strings.Count(strings.ToLower(got), "@font-face") != 1 {
		t.Fatalf("want @font-face only on after sheet, got %s", got)
	}
}

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
	envBase := strings.TrimSpace(os.Getenv("TYPECASE_MOCKUP_BASE"))
	if envBase != "" {
		spec.Base = envBase
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
				kept := st.frames[:0]
				for _, fr := range st.frames {
					if _, ok := beforeByID[fr.ID]; !ok {
						continue
					}
					kept = append(kept, fr)
				}
				if len(kept) == 0 {
					return
				}
				st.frames = kept
			}
			withTab := func(fn func(open func(html, id string) error, shot func(html, id string, width int) ([]byte, error)) error) error {
				h := browser.NewHarness(t)
				run, cancel := h.Tab(90 * time.Second)
				defer func() {
					cancel()
					h.CloseBrowser()
				}()

				open := func(html, id string) error {
					url := h.RegisterPage(id, html)
					return run(
						chromedp.EmulateViewport(int64(browser.MockupShotViewport), 900, chromedp.EmulateScale(browser.MockupShotScale)),
						chromedp.Navigate(url),
						chromedp.WaitReady("body"),
						chromedp.Evaluate(`(async function(){
  if (document.fonts && document.fonts.ready) {
    await Promise.race([
      document.fonts.ready,
      new Promise(function(r){ setTimeout(r, 2000); })
    ]);
  }
  return true;
})()`, nil),
						chromedp.ActionFunc(func(ctx context.Context) error {
							_, err := page.CaptureScreenshot().WithFormat(page.CaptureScreenshotFormatPng).Do(ctx)
							return err
						}),
					)
				}
				shot := func(html, id string, width int) ([]byte, error) {
					raw, err := json.Marshal(html)
					if err != nil {
						return nil, err
					}
					js := fmt.Sprintf(`(async function(){
  var host = document.getElementById("host");
  host.innerHTML = %s;
  var f = host.querySelector(".frame");
  if (!f) return false;
  f.style.width = %d + "px";
  f.style.maxWidth = %d + "px";
  f.style.maxHeight = %d + "px";
  f.style.overflow = "hidden";
  f.style.margin = "0";
  void f.offsetHeight;
  await new Promise(function(r){ requestAnimationFrame(function(){ requestAnimationFrame(r); }); });
  var ae = document.activeElement;
  if (ae && ae.blur) ae.blur();
  host.querySelectorAll("input,textarea,button,select").forEach(function(el){ el.blur(); });
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
						return nil, fmt.Errorf("%s @%d: %w", id, width, err)
					}
					if len(png) == 0 {
						return nil, fmt.Errorf("%s @%d: empty screenshot", id, width)
					}
					return png, nil
				}
				return fn(open, shot)
			}

			// One side of one stage: the shell of that side, then every frame
			// at every width. The browser process is the harness's own, so a
			// retry of this function starts a clean one.
			side := func(styles, sprite string, id string, html func(fr gate.MockupFrame) string) (map[string][]byte, error) {
				out := map[string][]byte{}
				err := withTab(func(open func(html, id string) error, shot func(html, id string, width int) ([]byte, error)) error {
					if err := open(browser.MockupShotShell(styles, sprite), id); err != nil {
						return err
					}
					for _, fr := range st.frames {
						for _, w := range browser.MockupRefactorWidths {
							if envBase == "" && spec.Skip(fr.ID, w) {
								continue
							}
							png, err := shot(html(fr), fr.ID, w)
							if err != nil {
								return err
							}
							out[fmt.Sprintf("%s@%d", fr.ID, w)] = png
						}
					}
					return nil
				})
				return out, err
			}

			// A hung snapshot is not a расхождение: say so, and give the stage
			// one more go in a new browser process.
			capture := func(styles, sprite, id string, html func(fr gate.MockupFrame) string) (map[string][]byte, bool) {
				png, err := side(styles, sprite, id, html)
				if err != nil && browser.IsHang(err) {
					t.Logf("%s: снимок завис (%v) — повтор в новом процессе браузера", id, err)
					png, err = side(styles, sprite, id, html)
				}
				if err != nil {
					if browser.IsHang(err) {
						t.Errorf("%s: снимок завис и после повтора (%v) — это таймаут, не расхождение", id, err)
					} else {
						t.Errorf("%s: %v", id, err)
					}
					return nil, false
				}
				return png, true
			}

			afterPNG, ok := capture(after.Styles, after.Sprite, "after", func(fr gate.MockupFrame) string { return fr.HTML })
			if !ok || sameDoc {
				return
			}
			beforePNG, ok := capture(before.Styles, before.Sprite, "before", func(fr gate.MockupFrame) string { return beforeByID[fr.ID].HTML })
			if !ok {
				return
			}
			for _, fr := range st.frames {
				for _, w := range browser.MockupRefactorWidths {
					if envBase == "" && spec.Skip(fr.ID, w) {
						continue
					}
					key := fmt.Sprintf("%s@%d", fr.ID, w)
					if err := browser.PNGsMatch(beforePNG[key], afterPNG[key], browser.MockupChannelTol); err != nil {
						if n := mockupDiffPixels(err); n > 0 && n <= browser.MockupAAPixelCap {
							continue
						}
						dumpMockupPNG(fr.ID, w, "before", beforePNG[key])
						dumpMockupPNG(fr.ID, w, "after", afterPNG[key])
						t.Errorf("%s @%d: %v", fr.ID, w, err)
					}
				}
			}
		})
	}
}

func mockupDiffPixels(err error) int {
	var n int
	if _, scanErr := fmt.Sscanf(err.Error(), "%d pixels", &n); scanErr != nil {
		return -1
	}
	return n
}

func dumpMockupPNG(id string, w int, side string, png []byte) {
	dir := filepath.Join(os.TempDir(), "tc-mockup-diff")
	_ = os.MkdirAll(dir, 0755)
	_ = os.WriteFile(filepath.Join(dir, fmt.Sprintf("%s-%d-%s.png", id, w, side)), png, 0644)
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
