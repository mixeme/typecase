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

// TestMockupStrays finds the frames that changed though nobody changed them
// (docs/development.md, Р-52): the working mockup against HEAD.
//
// A frame whose own HTML changed was meant to change — a person looks at its
// picture; it is named here and not measured. A frame whose HTML is the same
// can only have moved through the CSS or the sprite. When neither changed
// there is nothing to measure and no browser is started; when a rule changed,
// the frames its selectors can reach are measured on every step of the
// ladder — each node's box and everything its CSS computes to — and a frame
// that differs is a stray: red, named with the first node and property.
//
// Red is not a table to fill in. A stray is either mended or shown to the
// owner with the frames that were meant; the commit says which frames
// changed. After the commit HEAD is the new base.
//
// Another base: TYPECASE_MOCKUP_BASE=<commit>. Some frames only:
// TC_STRAYS_ONLY=s-dup-1,s-plan-2.
func TestMockupStrays(t *testing.T) {
	if os.Getenv("CI") == "" && testing.Short() {
		t.Skip("skip with -short outside CI")
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
	plan := gate.PlanStrays(before, after)
	t.Logf("against %s: %d frames meant to change %v, %d new %v, %d gone %v, %d the same",
		base, len(plan.Meant), plan.Meant, len(plan.New), plan.New, len(plan.Gone), plan.Gone, plan.Same)
	check := plan.Check
	if only := strings.TrimSpace(os.Getenv("TC_STRAYS_ONLY")); only != "" {
		want := map[string]bool{}
		for _, id := range strings.Split(only, ",") {
			want[strings.TrimSpace(id)] = true
		}
		check = nil
		for _, id := range plan.Check {
			if want[id] {
				check = append(check, id)
			}
		}
	}
	if len(check) == 0 {
		t.Logf("nothing can stray: %s", orNothing(plan.Why))
		return
	}
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	t.Logf("%s: %d frames to measure", plan.Why, len(check))
	strays := measureStrays(t, before, after, check)
	ids := make([]string, 0, len(strays))
	for id := range strays {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		t.Errorf("%s strayed: %s", id, strays[id])
	}
	t.Logf("%d of %d frames strayed against %s", len(ids), len(check), base)
}

func orNothing(why string) string {
	if why == "" {
		return "the CSS and the sprite are the same"
	}
	return why + ", and no unchanged frame is in reach"
}

// measureStrays measures the named frames with the CSS of both sides and says
// for each frame that differs where it does: the steps of the ladder and the
// first node and properties at the widest of them.
func measureStrays(t *testing.T, before, after gate.MockupParts, ids []string) map[string]string {
	t.Helper()
	want := map[string]bool{}
	for _, id := range ids {
		want[id] = true
	}
	side := func(parts gate.MockupParts) map[string][]string {
		t.Helper()
		h := browser.NewHarness(t)
		run, cancel := h.Tab(50 * time.Minute)
		defer func() {
			cancel()
			h.CloseBrowser()
		}()
		url := h.RegisterPage("strays", browser.MockupShotShell(parts.Styles, parts.Sprite))
		if err := run(
			chromedp.EmulateViewport(int64(browser.MockupShotViewport), int64(browser.MockupShotMaxHeight)),
			chromedp.Navigate(url),
			chromedp.WaitReady("body"),
			// without this the boxes are measured against the fallback font on
			// one side and the real one on the other, and every text node moves
			chromedp.Evaluate(fontsReadyJS, nil),
		); err != nil {
			t.Fatal(err)
		}
		out := map[string][]string{}
		for _, fr := range parts.Frames {
			if !want[fr.ID] {
				continue
			}
			raw, err := json.Marshal(fr.HTML)
			if err != nil {
				t.Fatal(err)
			}
			for _, w := range browser.MockupLadderWidths {
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
	now, was := side(after), side(before)

	out := map[string]string{}
	for _, id := range ids {
		var widths []string
		first := ""
		for _, w := range browser.MockupLadderWidths {
			k := fmt.Sprintf("%s@%d", id, w)
			a, b := now[k], was[k]
			d := ""
			if len(a) != len(b) {
				d = fmt.Sprintf("%d nodes, were %d", len(a), len(b))
			}
			for i := 0; d == "" && i < len(a); i++ {
				if a[i] != b[i] {
					d = fmt.Sprintf("node %d <%s>: %s", i, nodeTag(a[i]), strings.Join(firstFew(propDiff(b[i], a[i]), 4), ", "))
				}
			}
			if d == "" {
				continue
			}
			widths = append(widths, fmt.Sprint(w))
			if first == "" {
				first = d
			}
		}
		if len(widths) == 0 {
			continue
		}
		at := strings.Join(widths, ", ")
		if len(widths) == len(browser.MockupLadderWidths) {
			at = "all widths"
		}
		out[id] = fmt.Sprintf("at %s — %s", at, first)
	}
	return out
}

func firstFew(s []string, n int) []string {
	if len(s) <= n {
		return s
	}
	return append(s[:n:n], fmt.Sprintf("… %d more", len(s)-n))
}

// TestMockupStraysBite: a rule changed for one class moves a frame that has
// the class and nobody edited — the plan sends it to be measured, the measure
// names it; the frame the rule cannot reach is not measured at all.
func TestMockupStraysBite(t *testing.T) {
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
	all, err := gate.ParseMockupParts(string(html))
	if err != nil {
		t.Fatal(err)
	}
	before := gate.MockupParts{Styles: all.Styles, Sprite: all.Sprite}
	for _, f := range all.Frames {
		// a sheet with a key hint, a row of the queue without one, a frame far away
		if f.ID == "s-files-11" || f.ID == "s-files-24" || f.ID == "s-auth-1" {
			before.Frames = append(before.Frames, f)
		}
	}
	if len(before.Frames) != 3 {
		t.Fatalf("%d of the three frames found", len(before.Frames))
	}
	for _, c := range []struct {
		name, css string
		check     []string
		stray     []string
	}{
		// the colour moves no box: a compare of boxes alone would pass it
		{"a colour", ".tc-kbd{color:rgb(1, 2, 3)}", []string{"s-files-11"}, []string{"s-files-11"}},
		{"a width", ".tc-kbd{padding-left:40px}", []string{"s-files-11"}, []string{"s-files-11"}},
		{"a rule that changes nothing", ".tc-kbd{}", []string{"s-files-11"}, nil},
	} {
		after := before
		after.Styles = before.Styles + "<style>" + c.css + "</style>"
		plan := gate.PlanStrays(before, after)
		if fmt.Sprint(plan.Check) != fmt.Sprint(c.check) {
			t.Errorf("%s: to measure %v, want %v (%s)", c.name, plan.Check, c.check, plan.Why)
			continue
		}
		got := measureStrays(t, before, after, plan.Check)
		var ids []string
		for id := range got {
			ids = append(ids, id)
		}
		sort.Strings(ids)
		if fmt.Sprint(ids) != fmt.Sprint(c.stray) {
			t.Errorf("%s: strayed %v, want %v", c.name, got, c.stray)
		}
		for id, d := range got {
			t.Logf("%s: %s %s", c.name, id, d)
		}
	}
}

// fontsReadyJS waits for the webfonts, or gives up after two seconds.
const fontsReadyJS = `(async function(){
  if (document.fonts && document.fonts.ready) {
    await Promise.race([
      document.fonts.ready,
      new Promise(function(r){ setTimeout(r, 2000); })
    ]);
  }
  return true;
})()`

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
// node: its tag, its box in the frame and every computed property that the
// product CSS can set. Class names are left out on purpose — a rename changes
// them by design.
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
  var origin = f.getBoundingClientRect();
  var nodes = f.querySelectorAll("*");
  for (var i = 0; i < nodes.length; i++) {
    var el = nodes[i];
    var cs = getComputedStyle(el);
    var parts = [];
    for (var j = 0; j < cs.length; j++) {
      var p = cs[j];
      parts.push(p + ":" + cs.getPropertyValue(p));
    }
    var r = el.getBoundingClientRect();
    parts.push("rect:" + Math.round(r.left - origin.left) + "," + Math.round(r.top - origin.top) + " " + Math.round(r.width) + "x" + Math.round(r.height));
    // sorted: custom properties are enumerated in an order that depends on
    // which rules declare them, so adding a rule reshuffles the list
    parts.sort();
    out.push(el.tagName.toLowerCase() + ";" + parts.join(";"));
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
		{"dark-2", ".tc-avatar", "background-color", "rgb(58, 70, 61)"},
		{"dark-2", ".tc-avatar", "color", "rgb(201, 191, 166)"},
		{"dark-2", ".tc-scrim", "background-color", "rgba(0, 0, 0, 0.42)"},
		// a file's thumbnail: a folder's has no ground of its own (tc-thumb--folder)
		{"dark-3", ".tc-thumb:not(.tc-thumb--folder)", "background-color", "rgb(51, 64, 58)"},
		{"dark-3", ".tc-thumb--folder", "background-color", "rgba(0, 0, 0, 0)"},
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
