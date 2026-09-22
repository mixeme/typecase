package browser_test

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/chromedp/chromedp"

	"carrel.dev/typecase/internal/browser"
	"carrel.dev/typecase/internal/gate"
)

// TestMockupComputedStyles compares what the mockup's CSS computes to, node by
// node, between HEAD and the working tree.
//
// This is the check for a change that is not supposed to move anything — a
// rename, a regrouping of rules. The frames are then the same tree, so if
// every node computes to the same values it renders the same, and a difference
// is named as the node and the property instead of as a count of pixels. It
// takes minutes, not half an hour, because nothing is screenshotted.
//
// It is not a replacement for TestMockupRefactor: the trees have to match, so
// it says nothing about a change that adds or moves a node, and it does not
// see the sprite, the fonts or the page's own CSS. Point it at another commit
// with TYPECASE_MOCKUP_BASE.
func TestMockupComputedStyles(t *testing.T) {
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if os.Getenv("CI") == "" && testing.Short() {
		t.Skip("needs Chrome; skip with -short outside CI")
	}
	root := browser.ModuleRoot(t)
	base := strings.TrimSpace(os.Getenv("TYPECASE_MOCKUP_BASE"))
	if base == "" {
		base = "HEAD"
	}
	afterHTML, err := browser.LoadMockupHTML(root, "")
	if err != nil {
		t.Fatal(err)
	}
	beforeHTML, err := browser.LoadMockupHTML(root, base)
	if err != nil {
		t.Fatal(err)
	}
	if string(afterHTML) == string(beforeHTML) {
		t.Skipf("the mockup matches %s", base)
	}
	after, err := gate.ParseMockupParts(string(afterHTML))
	if err != nil {
		t.Fatal(err)
	}
	before, err := gate.ParseMockupParts(string(beforeHTML))
	if err != nil {
		t.Fatal(err)
	}
	beforeByID := map[string]gate.MockupFrame{}
	for _, f := range before.Frames {
		beforeByID[f.ID] = f
	}

	side := func(parts gate.MockupParts, pick func(id string) (string, bool)) map[string][]string {
		t.Helper()
		h := browser.NewHarness(t)
		run, cancel := h.Tab(20 * time.Minute)
		defer func() {
			cancel()
			h.CloseBrowser()
		}()
		url := h.RegisterPage("styles", browser.MockupShotShell(parts.Styles, parts.Sprite))
		if err := run(
			chromedp.EmulateViewport(int64(browser.MockupShotViewport), int64(browser.MockupShotMaxHeight)),
			chromedp.Navigate(url),
			chromedp.WaitReady("body"),
		); err != nil {
			t.Fatal(err)
		}
		out := map[string][]string{}
		for _, fr := range parts.Frames {
			html, ok := pick(fr.ID)
			if !ok {
				continue
			}
			raw, err := json.Marshal(html)
			if err != nil {
				t.Fatal(err)
			}
			for _, w := range browser.MockupRefactorWidths {
				var got []string
				js := fmt.Sprintf(computedStylesJS, raw, w, w)
				if err := run(chromedp.Evaluate(js, &got)); err != nil {
					t.Fatalf("%s @%d: %v", fr.ID, w, err)
				}
				out[fmt.Sprintf("%s@%d", fr.ID, w)] = got
			}
		}
		return out
	}

	afterStyles := side(after, func(id string) (string, bool) {
		fr, ok := beforeByID[id]
		_ = fr
		if !ok {
			return "", false // not in the base: the compare skips it too
		}
		for _, f := range after.Frames {
			if f.ID == id {
				return f.HTML, true
			}
		}
		return "", false
	})
	beforeStyles := side(before, func(id string) (string, bool) {
		fr, ok := beforeByID[id]
		return fr.HTML, ok
	})

	keys := make([]string, 0, len(afterStyles))
	for k := range afterStyles {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	bad := 0
	for _, k := range keys {
		a, b := afterStyles[k], beforeStyles[k]
		if len(a) != len(b) {
			t.Errorf("%s: %d nodes in the working tree, %d in %s — this check "+
				"needs the same tree on both sides", k, len(a), len(b), base)
			bad++
			continue
		}
		for i := range a {
			if a[i] == b[i] {
				continue
			}
			bad++
			if bad <= 24 {
				t.Errorf("%s node %d <%s>: %s", k, i, nodeTag(a[i]),
					strings.Join(propDiff(b[i], a[i]), ", "))
			}
		}
	}
	if bad > 20 {
		t.Errorf("… %d differing nodes in all", bad)
	}
	t.Logf("compared %d frame×width pairs against %s", len(keys), base)
}

func nodeTag(line string) string {
	if i := strings.IndexByte(line, ';'); i > 0 {
		return line[:i]
	}
	return line
}

// propDiff names the properties whose value moved, so a difference reads as
// "padding-top: 9px -> 12px" instead of three hundred properties that agree.
func propDiff(before, after string) []string {
	split := func(s string) map[string]string {
		out := map[string]string{}
		for _, part := range strings.Split(s, ";") {
			if i := strings.IndexByte(part, ':'); i > 0 {
				out[part[:i]] = part[i+1:]
			}
		}
		return out
	}
	b, a := split(before), split(after)
	seen := map[string]bool{}
	var names []string
	for p, v := range a {
		if b[p] != v {
			names = append(names, p)
			seen[p] = true
		}
	}
	for p := range b {
		if _, ok := a[p]; !ok && !seen[p] {
			names = append(names, p)
		}
	}
	sort.Strings(names)
	var out []string
	for _, p := range names {
		out = append(out, fmt.Sprintf("%s: %s -> %s", p, b[p], a[p]))
	}
	return out
}

// computedStylesJS injects one frame, sets its width, and returns one line per
// node: its tag and every computed property that the product CSS can set.
// Class names are left out on purpose — the rename changes them by design.
const computedStylesJS = `(function(){
  var host = document.getElementById("host");
  host.innerHTML = %s;
  var f = host.querySelector(".frame");
  if (!f) return [];
  f.style.width = %d + "px";
  f.style.maxWidth = %d + "px";
  f.style.margin = "0";
  void f.offsetHeight;
  var out = [];
  var nodes = f.querySelectorAll("*");
  for (var i = 0; i < nodes.length; i++) {
    var el = nodes[i];
    var cs = getComputedStyle(el);
    var parts = [el.tagName.toLowerCase()];
    for (var j = 0; j < cs.length; j++) {
      var p = cs[j];
      parts.push(p + ":" + cs.getPropertyValue(p));
    }
    var r = el.getBoundingClientRect();
    parts.push("@" + Math.round(r.width) + "x" + Math.round(r.height));
    out.push(parts.join(";"));
  }
  return out;
})()`

// TestMockupDarkRules checks that the components with a hard-coded colour are
// actually repainted in a dark frame.
//
// Tokens carry the dark theme almost everywhere, but the avatar and the
// preview tile are filled with cardboard and the scrim is ink over a light
// page, so each has a rule of its own under [data-tc-theme="dark"]. dead-rules
// only says the selector matches a node somewhere; this says the node ends up
// the right colour, which is what the frames were drawn for.
func TestMockupDarkRules(t *testing.T) {
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if os.Getenv("CI") == "" && testing.Short() {
		t.Skip("needs Chrome; skip with -short outside CI")
	}
	root := browser.ModuleRoot(t)
	html, err := browser.LoadMockupHTML(root, "")
	if err != nil {
		t.Fatal(err)
	}
	parts, err := gate.ParseMockupParts(string(html))
	if err != nil {
		t.Fatal(err)
	}
	byID := map[string]string{}
	for _, f := range parts.Frames {
		byID[f.ID] = f.HTML
	}

	h := browser.NewHarness(t)
	run, cancel := h.Tab(5 * time.Minute)
	defer cancel()
	url := h.RegisterPage("dark", browser.MockupShotShell(parts.Styles, parts.Sprite))
	if err := run(
		chromedp.EmulateViewport(int64(browser.MockupShotViewport), int64(browser.MockupShotMaxHeight)),
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
	); err != nil {
		t.Fatal(err)
	}

	for _, c := range []struct{ frame, sel, prop, want string }{
		{"m-dark-2", ".tc-avatar", "background-color", "rgb(58, 70, 61)"},
		{"m-dark-2", ".tc-avatar", "color", "rgb(201, 191, 166)"},
		{"m-dark-2", ".tc-scrim", "background-color", "rgba(0, 0, 0, 0.42)"},
		{"m-dark-3", ".tc-thumb", "background-color", "rgb(51, 64, 58)"},
	} {
		frame, ok := byID[c.frame]
		if !ok {
			t.Fatalf("no frame %s", c.frame)
		}
		raw, err := json.Marshal(frame)
		if err != nil {
			t.Fatal(err)
		}
		var got string
		js := fmt.Sprintf(`(function(){
  var host = document.getElementById("host");
  host.innerHTML = %s;
  var el = host.querySelector(%q);
  if (!el) return "no node";
  return getComputedStyle(el).getPropertyValue(%q);
})()`, raw, c.sel, c.prop)
		if err := run(chromedp.Evaluate(js, &got)); err != nil {
			t.Fatalf("%s %s: %v", c.frame, c.sel, err)
		}
		if got != c.want {
			t.Errorf("%s %s %s: %s, want %s", c.frame, c.sel, c.prop, got, c.want)
		}
	}
}
