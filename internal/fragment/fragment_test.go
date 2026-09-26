package fragment

import (
	"strings"
	"testing"
)

const page = `<!doctype html><html><head>
<style>.page{white-space:pre-wrap}</style>
<style data-tc>.tc-prose{max-width:72ch;white-space:pre-wrap}
.tc-row .tc-x{white-space:pre}</style></head><body>
<div class="frame" data-frame="a-1"><div class="frame-label"><span class="tc-btn">label</span></div>
<div class="tc-root">
  <button class="tc-btn is-primary" type="button" title="Save it" aria-label="Save">
     Save   now </button>
  <div class="tc-prose">line one
  line two</div>
  <input class="tc-in" type="date" value="2026-08-10" aria-label="Date" id="d">
  <span class="tc-mark" data-tc-colour="#4A6B52"></span>
  <button class="tc-btn" type="button">Cancel</button>
</div></div>
<div class="frame" data-frame="a-2"><div class="tc-root"><button class="tc-btn" type="button">Other</button></div></div>
</body></html>`

func parsed(t *testing.T) *Mockup {
	t.Helper()
	m, err := Parse(page)
	if err != nil {
		t.Fatal(err)
	}
	return m
}

func TestBlockSkipsLabelAndCountsFromOne(t *testing.T) {
	m := parsed(t)
	n, err := m.Block(Ref{Frame: "a-1", Class: "tc-btn", N: 1})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(Skeleton(m.Normalize(n)), "is-primary") {
		t.Fatal("first .tc-btn should be Save, not the label's span")
	}
	if _, err := m.Block(Ref{Frame: "a-1", Class: "tc-btn", N: 3}); err == nil {
		t.Fatal("a third .tc-btn does not exist")
	}
	if _, err := m.Block(Ref{Frame: "a-1", Class: "tc-btn", N: 0}); err == nil {
		t.Fatal("N counts from 1")
	}
}

func TestNormalize(t *testing.T) {
	m := parsed(t)
	if !m.pre["tc-prose"] || !m.pre["tc-x"] || m.pre["page"] {
		t.Fatalf("pre classes: the subject of a pre rule in data-tc, not the page's: %v", m.pre)
	}
	if err := m.Compare(Ref{"a-1", "tc-btn", 1}, `<button class="is-primary tc-btn" type="submit">Save now</button>`); err != nil {
		t.Fatalf("white space collapses, class order and type of a button do not count: %v", err)
	}
	if err := m.Compare(Ref{"a-1", "tc-prose", 1}, `<div class="tc-prose">line one line two</div>`); err == nil {
		t.Fatal("pre-wrap text keeps its line breaks")
	}
	if err := m.Compare(Ref{"a-1", "tc-in", 1}, `<input class="tc-in" type="date" value="2026-08-11">`); err == nil {
		t.Fatal("input value is compared")
	}
	if err := m.Compare(Ref{"a-1", "tc-mark", 1}, `<span class="tc-mark" data-tc-colour="#000"></span>`); err == nil {
		t.Fatal("data-tc-colour is compared")
	}
	if err := m.Compare(Ref{"a-1", "tc-btn", 2}, `<button class="tc-btn">Cancel</button><button class="tc-btn">x</button>`); err == nil {
		t.Fatal("two top elements for one block")
	}
}

func TestShapes(t *testing.T) {
	m := parsed(t)
	shapes := m.Shapes("tc-btn")
	if len(shapes) != 2 {
		t.Fatalf("want 2 shapes (plain, primary), got %d: %+v", len(shapes), shapes)
	}
	if shapes[0].Count != 2 || shapes[0].First != (Ref{"a-1", "tc-btn", 2}) {
		t.Fatalf("plain button: %+v", shapes[0])
	}
}

func TestBlockAndMixins(t *testing.T) {
	m, err := Parse(`<div class="frame" data-frame="b-1"><div class="tc-root tc-shell">
<div class="tc-bar"><button class="tc-btn tc-push is-2nd tc-bar__more" type="button">x</button>
<span class="tc-push tc-dim">note</span><button class="tc-btn tc-close" type="button">×</button>
<h3 class="tc-rubric tc-line">Title</h3></div></div></div>`)
	if err != nil {
		t.Fatal(err)
	}
	f, _ := m.Frame("b-1")
	got := map[string]int{}
	for _, p := range m.Pairs() {
		got[p.Parent+">"+p.Child]++
	}
	for _, want := range []string{"tc-shell>tc-bar", "tc-bar>tc-btn", "tc-bar>tc-dim", "tc-bar>tc-close", "tc-bar>tc-rubric"} {
		if got[want] == 0 {
			t.Errorf("pair %s missing: %v", want, got)
		}
	}
	_ = f
	// The bar's own form keeps the mixins it puts on its children.
	bar, _ := m.Block(Ref{"b-1", "tc-bar", 1})
	key := Skeleton(m.Own(bar, false))
	for _, want := range []string{"tc-btn.is-2nd.tc-bar__more.tc-push", "tc-dim.tc-push", "tc-close", "tc-rubric.tc-line"} {
		if !strings.Contains(key, want) {
			t.Errorf("bar's own form lacks %s: %s", want, key)
		}
	}
	// The button printed alone has no mixins; the frame's root mixins are stripped.
	if err := m.Compare(Ref{"b-1", "tc-btn", 1}, `<button class="tc-btn" type="button">x</button>`); err != nil {
		t.Fatalf("root mixins are the container's: %v", err)
	}
	if err := m.Compare(Ref{"b-1", "tc-btn", 1}, `<button class="tc-btn tc-push" type="button">x</button>`); err == nil {
		t.Fatal("a widget printing a mixin by itself passed")
	}
	if err := m.Compare(Ref{"b-1", "tc-close", 1}, `<button class="tc-btn tc-close" type="button">×</button>`); err != nil {
		t.Fatalf("tc-close is built on tc-btn: %v", err)
	}
}
