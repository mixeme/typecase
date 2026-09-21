package browser_test

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/chromedp/chromedp"

	"carrel.dev/typecase/internal/browser"
)

type numberedEffect struct {
	Class  string   `json:"class"`
	Uses   int      `json:"uses"`
	Props  []string `json:"props"`
	Moved  []string `json:"moved"`
	Skip   string   `json:"skip"`
	Sample string   `json:"sample"`
}

// TestNumberedClassEffect asks each --rN class whether it does anything: it
// reads the properties the rule declares, takes the class off the node, and
// reads them again. A class whose properties do not move is redundant — the
// node was already getting those values — and folding it away costs no pixels.
//
// Run on demand: go test ./internal/browser -run TestNumberedClassEffect -v
func TestNumberedClassEffect(t *testing.T) {
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if os.Getenv("TC_NUMBERED_EFFECT") == "" {
		t.Skip("set TC_NUMBERED_EFFECT=1 to measure")
	}
	root := browser.ModuleRoot(t)
	html, err := browser.LoadMockupHTML(root, "")
	if err != nil {
		t.Fatal(err)
	}
	h := browser.NewHarness(t)
	run, cancel := h.Tab(20 * time.Minute)
	defer cancel()
	url := h.RegisterPage("mockup-effect", string(html))
	if err := run(
		chromedp.EmulateViewport(1500, 900),
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
	); err != nil {
		t.Fatal(err)
	}

	var got []numberedEffect
	if err := run(chromedp.Evaluate(numberedEffectJS, &got)); err != nil {
		t.Fatal(err)
	}

	var redundant, effective, skipped []numberedEffect
	for _, e := range got {
		switch {
		case e.Skip != "":
			skipped = append(skipped, e)
		case len(e.Moved) == 0:
			redundant = append(redundant, e)
		default:
			effective = append(effective, e)
		}
	}
	sort.Slice(redundant, func(i, j int) bool { return redundant[i].Class < redundant[j].Class })
	t.Logf("measured %d classes: %d do nothing, %d change the drawing, %d skipped",
		len(got), len(redundant), len(effective), len(skipped))
	for _, e := range redundant {
		t.Logf("  redundant %-28s uses=%d  %v", e.Class, e.Uses, e.Props)
	}
	reasons := map[string]int{}
	for _, e := range skipped {
		reasons[e.Skip]++
	}
	for r, n := range reasons {
		t.Logf("  skipped: %s ×%d", r, n)
	}
	out := map[string]any{"redundant": redundant, "effective": effective, "skipped": skipped}
	blob, _ := json.MarshalIndent(out, "", " ")
	path := fmt.Sprintf("%s/numbered-effect.json", os.TempDir())
	_ = os.WriteFile(path, blob, 0o644)
	t.Logf("full result: %s", path)
}

const numberedEffectJS = `(function(){
  var numbered = /--r[0-9]+$/;
  // Every rule that names a numbered class, plain or inside @container/@media.
  var rules = {};   // class -> [{plain:bool, props:[...]}]
  function walk(list){
    for (var i = 0; i < list.length; i++) {
      var r = list[i];
      // A CSSStyleRule also has .cssRules in Chrome (CSS nesting), so the
      // style rule is read first and the recursion happens afterwards.
      if (r.cssRules && r.cssRules.length) walk(r.cssRules);
      if (!r.selectorText || !r.style || !r.style.length) continue;
      var names = (r.selectorText.match(/\.[A-Za-z_][\w-]*/g) || [])
        .map(function(s){ return s.slice(1); })
        .filter(function(c){ return numbered.test(c); });
      if (!names.length) continue;
      var props = [];
      for (var k = 0; k < r.style.length; k++) props.push(r.style[k]);
      var plain = r.parentRule === null &&
        r.selectorText.replace(/\s/g, "") === "." + names[0] + "." + names[0];
      names.forEach(function(c){
        (rules[c] = rules[c] || []).push({plain: plain, props: props});
      });
    }
  }
  for (var s = 0; s < document.styleSheets.length; s++) {
    try { walk(document.styleSheets[s].cssRules); } catch (e) {}
  }

  var out = [];
  Object.keys(rules).forEach(function(cls){
    var entry = {"class": cls, uses: 0, props: [], moved: [], skip: "", sample: ""};
    var nodes = document.querySelectorAll("." + CSS.escape(cls));
    entry.uses = nodes.length;
    if (!nodes.length) { entry.skip = "no node"; out.push(entry); return; }
    var defs = rules[cls];
    if (defs.some(function(d){ return !d.plain; })) {
      entry.skip = "context rule (@container/@media or a compound selector)";
      out.push(entry); return;
    }
    var props = defs[0].props;
    entry.props = props;
    var frame = nodes[0].closest("[data-frame]");
    entry.sample = frame ? frame.getAttribute("data-frame") : "demo";
    var moved = {};
    for (var n = 0; n < nodes.length; n++) {
      var el = nodes[n];
      var before = {};
      props.forEach(function(p){ before[p] = getComputedStyle(el).getPropertyValue(p); });
      el.classList.remove(cls);
      void el.offsetHeight;
      props.forEach(function(p){
        if (getComputedStyle(el).getPropertyValue(p) !== before[p]) moved[p] = true;
      });
      el.classList.add(cls);
      void el.offsetHeight;
    }
    entry.moved = Object.keys(moved);
    out.push(entry);
  });
  return out;
})()`
