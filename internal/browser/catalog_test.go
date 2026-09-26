package browser_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"

	"carrel.dev/typecase/internal/browser"
	"carrel.dev/typecase/internal/catalog"
)

// catalogServer serves the whole catalog site: the page plus every asset it
// links. Both this helper and cmd/catalog build the same mux, so a route the
// catalog forgets is missing here too (В-12).
func catalogServer(t *testing.T, withJS bool) *httptest.Server {
	return catalogServerCSS(t, withJS, "")
}

func catalogServerCSS(t *testing.T, withJS bool, cssOverride string) *httptest.Server {
	t.Helper()
	mux, err := catalog.Mux(cssOverride)
	if err != nil {
		t.Fatalf("catalog mux: %v", err)
	}
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return srv
}

// catalogURL is the page to open: with behaviour, or the static one.
func catalogURL(srv *httptest.Server, withJS bool) string {
	if withJS {
		return srv.URL
	}
	return srv.URL + "/static"
}

// escapedJS reports catalog cells whose positioned children spill outside
// the cell, plus the theme the boot script settled on.
const escapedJS = `(() => {
  const bad = [];
  document.querySelectorAll('.catalog-demo').forEach(cell => {
    const c = cell.getBoundingClientRect();
    for (const el of cell.querySelectorAll('*')) {
      const pos = getComputedStyle(el).position;
      if (pos !== 'fixed' && pos !== 'absolute') continue;
      const r = el.getBoundingClientRect();
      if (r.width === 0 && r.height === 0) continue;
      if (r.top < c.top - 2 || r.left < c.left - 2 || r.bottom > c.bottom + 2 || r.right > c.right + 2) {
        bad.push(cell.closest('.catalog-entry').id + ' / ' + el.className);
      }
    }
  });
  return {
    entries: document.querySelectorAll('.catalog-entry').length,
    theme: document.documentElement.getAttribute('data-tc-theme'),
    escaped: bad
  };
})()`

// inkOverflowJS flags text nodes whose font bounding box exceeds the line box
// inside overflow-hidden ancestors. scrollHeight misses ink overflow.
const inkOverflowJS = `(() => {
  const bad = [];
  const canvas = document.createElement('canvas');
  const ctx = canvas.getContext('2d');
  for (const el of document.querySelectorAll('*')) {
    const cs = getComputedStyle(el);
    if (cs.overflow === 'visible') continue;
    const text = el.textContent.trim();
    if (!text) continue;
    const lh = parseFloat(cs.lineHeight);
    if (!lh) continue;
    ctx.font = cs.font;
    const m = ctx.measureText(text);
    const fontBox = (m.fontBoundingBoxAscent || m.actualBoundingBoxAscent || 0) +
      (m.fontBoundingBoxDescent || m.actualBoundingBoxDescent || 0);
    const pad = (parseFloat(cs.paddingTop) || 0) + (parseFloat(cs.paddingBottom) || 0);
    if (fontBox > lh + pad + 0.5) {
      bad.push(el.className || el.tagName);
    }
  }
  return bad;
})()`

type catalogProbe struct {
	Entries int      `json:"entries"`
	Theme   string   `json:"theme"`
	Escaped []string `json:"escaped"`
}

// TestCatalogCellsContainOverlays keeps every state readable on one page.
// Scrims, overlays and the rail drawer are fixed to the viewport by design, so
// without a per-cell containing block a single open scrim hides the whole
// catalog.
func TestCatalogCellsContainOverlays(t *testing.T) {
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if testing.Short() {
		t.Skip("catalog layout check needs Chrome")
	}
	srv := catalogServer(t, true)
	h := browser.NewHarness(t)

	var probe catalogProbe
	if err := h.Run(srv.URL, 1200, 900, chromedp.Evaluate(escapedJS, &probe)); err != nil {
		t.Fatal(err)
	}
	if probe.Entries == 0 {
		t.Fatal("catalog rendered no entries")
	}
	if probe.Theme != "light" && probe.Theme != "dark" {
		t.Errorf("theme boot script did not run before paint: data-tc-theme=%q", probe.Theme)
	}
	for _, e := range probe.Escaped {
		t.Errorf("escapes its cell: %s", e)
	}
}

// TestCatalogFocusOnLoad ensures pre-open overlays do not steal focus at load.
func TestCatalogFocusOnLoad(t *testing.T) {
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if testing.Short() {
		t.Skip("catalog focus check needs Chrome")
	}
	srv := catalogServer(t, true)
	h := browser.NewHarness(t)

	var onBody bool
	if err := h.Run(srv.URL, 1200, 900, chromedp.Evaluate(
		`document.activeElement === document.body || document.activeElement === document.documentElement`,
		&onBody,
	)); err != nil {
		t.Fatal(err)
	}
	if !onBody {
		t.Fatalf("focus should stay on document body after catalog load, got %v", onBody)
	}
}

func TestCatalogTextInkFitsLineBox(t *testing.T) {
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if testing.Short() {
		t.Skip("catalog ink check needs Chrome")
	}
	srv := catalogServer(t, true)
	h := browser.NewHarness(t)

	var bad []string
	if err := h.Run(srv.URL, 1280, 900, chromedp.Evaluate(inkOverflowJS, &bad)); err != nil {
		t.Fatal(err)
	}
	for _, cls := range bad {
		t.Errorf("font box exceeds line box: %s", cls)
	}
}

// assetsJS reports faces the page failed to load and images that did not
// decode. Both fail silently in a browser: a font that answers with HTML is
// dropped for the fallback, and a broken picture keeps its box (В-12).
const assetsJS = `(async () => {
  await document.fonts.ready;
  const faces = [];
  document.fonts.forEach(f => { if (f.status !== "loaded") faces.push(f.family + ":" + f.status); });
  const imgs = [];
  for (const i of document.images) { if (i.naturalWidth === 0) imgs.push(i.getAttribute("src")); }
  return { faces: faces, images: imgs, declared: document.fonts.size };
})()`

type assetProbe struct {
	Faces    []string `json:"faces"`
	Images   []string `json:"images"`
	Declared int      `json:"declared"`
}

// assetProbeAction evaluates assetsJS and waits for the promise it returns.
func assetProbeAction(dest *assetProbe) chromedp.Action {
	return chromedp.Evaluate(assetsJS, dest, func(p *runtime.EvaluateParams) *runtime.EvaluateParams {
		return p.WithAwaitPromise(true)
	})
}

func TestCatalogAssetsLoad(t *testing.T) {
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if testing.Short() {
		t.Skip("catalog asset check needs Chrome")
	}
	srv := catalogServer(t, true)
	h := browser.NewHarness(t)

	var probe assetProbe
	if err := h.Run(srv.URL, 1280, 900, assetProbeAction(&probe)); err != nil {
		t.Fatal(err)
	}
	if probe.Declared == 0 {
		t.Fatal("catalog declares no font faces")
	}
	if len(probe.Faces) > 0 {
		t.Errorf("font faces not loaded: %v", probe.Faces)
	}
	if len(probe.Images) > 0 {
		t.Errorf("images did not decode: %v", probe.Images)
	}
}

// TestBite_catalogAssets: a catch-all that answers every path with the page
// is exactly what the catalog did before В-12, and the check must bite it.
func TestBite_catalogAssets(t *testing.T) {
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if testing.Short() {
		t.Skip("catalog asset bite needs Chrome")
	}
	mux, err := catalog.Mux("")
	if err != nil {
		t.Fatal(err)
	}
	page := catalog.PageHTML("Typecase catalog", catalog.Entries(), true)
	catchAll := http.NewServeMux()
	catchAll.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/", "/typecase.css", "/catalog.css", "/theme.js", "/tc.js":
			mux.ServeHTTP(w, r)
		default: // fonts and pictures get the page, as they used to
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			_, _ = w.Write([]byte(page))
		}
	})
	srv := httptest.NewServer(catchAll)
	t.Cleanup(srv.Close)
	h := browser.NewHarness(t)

	var probe assetProbe
	if err := h.Run(srv.URL, 1280, 900, assetProbeAction(&probe)); err != nil {
		t.Fatal(err)
	}
	if len(probe.Faces) == 0 && len(probe.Images) == 0 {
		t.Fatal("catch-all server should fail the asset check")
	}
}

// typographyJS reports text drawn with a family the library never declared:
// a control that does not inherit the font falls back to the system face
// (В-15). Sizes are the frames' business: the ladder compares them.
const typographyJS = `(() => {
  const declared = new Set(["Literata", "Public Sans", "Source Code Pro"]);
  const bad = new Set();
  for (const el of document.querySelectorAll(".catalog-demo *")) {
    let text = false;
    for (const n of el.childNodes) { if (n.nodeType === 3 && n.textContent.trim()) text = true; }
    if (!text) continue;
    const fam = getComputedStyle(el).fontFamily.split(",")[0].replace(/["']/g, "").trim();
    if (!declared.has(fam)) bad.add((el.className || el.tagName) + " → " + fam);
  }
  return [...bad];
})()`

func TestCatalogTypography(t *testing.T) {
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if testing.Short() {
		t.Skip("catalog typography check needs Chrome")
	}
	srv := catalogServer(t, true)
	h := browser.NewHarness(t)

	var families []string
	if err := h.Run(srv.URL, 1280, 900, chromedp.Evaluate(typographyJS, &families)); err != nil {
		t.Fatal(err)
	}
	for _, f := range families {
		t.Errorf("text drawn with an undeclared family: %s", f)
	}
}

// cellTypeJS reads how the library sets type in a cell: the cell's own face
// and a head title's, and whether tc.js is on the page. While the base hung on
// a class the script added, the page without it fell back to the browser's
// defaults (В-24).
const cellTypeJS = `(() => {
  const cell = document.querySelector(".catalog-demo");
  const title = document.querySelector(".catalog-demo .tc-head__title");
  const face = el => el ? getComputedStyle(el).fontFamily.split(",")[0].replace(/["']/g, "").trim() + " " + getComputedStyle(el).fontSize : "";
  return {
    cell: face(cell),
    title: face(title),
    scripted: document.querySelectorAll("script[src$='/tc.js']").length
  };
})()`

type cellTypeProbe struct {
	Cell     string `json:"cell"`
	Title    string `json:"title"`
	Scripted int    `json:"scripted"`
}

// TestCatalogBaseWithoutScript: the page without behaviour sets type exactly
// as the page with it does.
func TestCatalogBaseWithoutScript(t *testing.T) {
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if testing.Short() {
		t.Skip("no-script base check needs Chrome")
	}
	srv := catalogServer(t, false)
	h := browser.NewHarness(t)

	var with, without cellTypeProbe
	if err := h.Run(catalogURL(srv, true), 1280, 900, chromedp.Evaluate(cellTypeJS, &with)); err != nil {
		t.Fatal(err)
	}
	if err := h.Run(catalogURL(srv, false), 1280, 900, chromedp.Evaluate(cellTypeJS, &without)); err != nil {
		t.Fatal(err)
	}
	if without.Scripted != 0 || with.Scripted == 0 {
		t.Fatalf("tc.js on the page with behaviour %d times, without it %d times", with.Scripted, without.Scripted)
	}
	if with.Title == "" {
		t.Fatal("no head title in the catalog")
	}
	if !strings.HasPrefix(with.Cell, "Public Sans ") {
		t.Errorf("a cell sets %q, want Public Sans", with.Cell)
	}
	if without != (cellTypeProbe{Cell: with.Cell, Title: with.Title}) {
		t.Errorf("without behaviour the type is %+v, with it %+v", without, with)
	}
}

// segClickJS presses the second option of the first switch and reports where
// the mark and the announced state ended up; segKeyJS then moves the choice
// with the arrow key from there.
const segClickJS = `(() => {
  const seg = [...document.querySelectorAll(".catalog-demo .tc-seg")].find(s => s.querySelectorAll(".tc-seg__opt").length > 2);
  if (!seg) return { missing: true };
  seg.id = "seg-under-test";
  const opts = [...seg.querySelectorAll(".tc-seg__opt")];
  opts[1].click();
  return segState(opts, 1);
  function segState(opts, want) {
    const on = opts.filter(b => b.classList.contains("is-on"));
    const checked = opts.filter(b => b.getAttribute("aria-checked") === "true");
    return { on: on.length, checked: checked.length, at: on.length === 1 && checked.length === 1 && on[0] === checked[0] && on[0] === opts[want] };
  }
})()`

const segKeyJS = `(() => {
  const seg = document.getElementById("seg-under-test");
  const opts = [...seg.querySelectorAll(".tc-seg__opt")];
  opts[1].focus();
  opts[1].dispatchEvent(new KeyboardEvent("keydown", { key: "ArrowRight", bubbles: true }));
  const on = opts.filter(b => b.classList.contains("is-on"));
  const checked = opts.filter(b => b.getAttribute("aria-checked") === "true");
  return { on: on.length, checked: checked.length, at: on.length === 1 && checked.length === 1 && on[0] === checked[0] && on[0] === opts[2] && document.activeElement === opts[2] };
})()`

type segProbe struct {
	Missing bool `json:"missing"`
	On      int  `json:"on"`
	Checked int  `json:"checked"`
	At      bool `json:"at"`
}

func segProbes(t *testing.T, withJS bool) (click, key segProbe) {
	t.Helper()
	srv := catalogServer(t, withJS)
	h := browser.NewHarness(t)
	if err := h.Run(catalogURL(srv, withJS), 1280, 900, chromedp.Evaluate(segClickJS, &click), chromedp.Evaluate(segKeyJS, &key)); err != nil {
		t.Fatal(err)
	}
	if click.Missing {
		t.Fatal("no tc-seg of three options in the catalog")
	}
	return click, key
}

// TestCatalogSegSelectionAgrees: pressing an option of a switch chooses it
// and the arrows move the choice (L-4 tc-seg); the mark and the announced
// state stay on the same option.
func TestCatalogSegSelectionAgrees(t *testing.T) {
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if testing.Short() {
		t.Skip("seg selection check needs Chrome")
	}
	click, key := segProbes(t, true)
	if !click.At {
		t.Errorf("after a click: %d marked, %d announced, not both on the option pressed", click.On, click.Checked)
	}
	if !key.At {
		t.Errorf("after the arrow: %d marked, %d announced, not both on the next option with the focus", key.On, key.Checked)
	}
}

// TestBite_catalogSeg: without tc.js nothing chooses, and the check says so.
func TestBite_catalogSeg(t *testing.T) {
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if testing.Short() {
		t.Skip("seg bite needs Chrome")
	}
	click, key := segProbes(t, false)
	if click.At || key.At {
		t.Fatal("a switch without behaviour passed the selection check")
	}
}
