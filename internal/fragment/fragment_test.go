package fragment

import (
	"errors"
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

// A record stands for the occurrence its widget prints as (the owner,
// 20.11.2026): something put in before it does not move it; one no
// occurrence prints as is told the nearest. The state of a record and of a
// frame's element are its root's own classes (the review of 14.10).
func TestLocateAndState(t *testing.T) {
	m := parsed(t)
	at, err := m.Locate(Ref{Frame: "a-1", Class: "tc-btn"}, `<button class="tc-btn" type="button">Cancel</button>`)
	if err != nil || at.N != 2 {
		t.Fatalf("Cancel: %v %v; want #2", at, err)
	}
	if err := m.Compare(Ref{"a-1", "tc-btn", 1}, `<button class="tc-btn" type="button">Cancel</button>`); err == nil {
		t.Fatal("Cancel against the place #1: want a difference")
	}
	_, err = m.Locate(Ref{Frame: "a-1", Class: "tc-btn"}, `<button class="tc-btn" type="button">Close</button>`)
	var nf *NotFound
	if !errors.As(err, &nf) || nf.At.N != 2 || !strings.Contains(nf.Diff, "Close") {
		t.Fatalf("Close: %v; want the nearest #2 and the text that parts", err)
	}
	_, err = m.Locate(Ref{Frame: "a-1", Class: "tc-btn"}, `<button class="tc-btn is-primary" type="button" title="Save it" aria-label="Save">Keep</button>`)
	if !errors.As(err, &nf) || nf.At.N != 1 {
		t.Fatalf("a primary button: %v; want the nearest #1, the same tree", err)
	}
	// Two forms that differ only in the number of children tie on Diffs:
	// the one that shares the record's words is the nearest.
	two, err := Parse(`<div class="frame" data-frame="f-1"><div class="tc-root">
<div class="tc-field"><label>Server</label><span class="tc-mono">nextcloud</span><span class="tc-dim">baikal is not answering</span></div>
<div class="tc-field"><label>Colour</label><span class="tc-mono">green</span><span class="tc-dim">written to the server</span></div>
</div></div>`)
	if err != nil {
		t.Fatal(err)
	}
	_, err = two.Locate(Ref{Frame: "f-1", Class: "tc-field"}, `<div class="tc-field"><label>Colour</label><span class="tc-mono">green</span></div>`)
	if !errors.As(err, &nf) || nf.At.N != 2 {
		t.Fatalf("the colour field without its hint: %v; want the nearest #2, the field that shares its words", err)
	}
	if _, err := m.Block(Ref{Frame: "a-1", Class: "tc-btn"}); err == nil {
		t.Fatal("a record has no place of its own: want Block to refuse it")
	}
	want, err := m.RootState(Ref{"a-1", "tc-btn", 1})
	if err != nil || strings.Join(want, " ") != "is-primary" {
		t.Fatalf("the state of #1: %v %v; want is-primary", want, err)
	}
	got, err := RenderedState("tc-btn", `<button class="tc-btn tc-push" type="button">Cancel</button>`)
	if err != nil || len(got) != 0 {
		t.Fatalf("a mixin is no state: %v %v", got, err)
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
	for _, want := range []string{"tc-btn.is-2nd.tc-bar__more.tc-push", "tc-dim.tc-push", "tc-close", "tc-rubric)"} {
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

func TestWindowChrome(t *testing.T) {
	m, err := Parse(`<div class="frame" data-frame="m-1"><div class="tc-root tc-shell"><div class="mk-win">
<div class="mk-titlebar"><span class="mk-titlebar__title">App</span><button class="tc-btn" type="button">x</button></div>
<div class="tc-top"><button class="tc-btn" type="button">New</button></div>
<div class="mk-status"><span class="tc-mono">ok</span></div></div></div></div>`)
	if err != nil {
		t.Fatal(err)
	}
	n, err := m.Block(Ref{Frame: "m-1", Class: "tc-shell", N: 1})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := Skeleton(m.Own(n, false)), "div.tc-root.tc-shell(tc-top)"; got != want {
		t.Errorf("own = %s, want %s: the window is transparent and its frame unseen", got, want)
	}
	b, err := m.Block(Ref{Frame: "m-1", Class: "tc-btn", N: 1})
	if err != nil {
		t.Fatal(err)
	}
	if text(b) != "New" {
		t.Errorf("the first button is %q: a button in the title bar was counted", text(b))
	}
	for _, p := range m.Pairs() {
		if p.Child == "tc-mono" {
			t.Errorf("the status line's %v was paired", p)
		}
	}
}
