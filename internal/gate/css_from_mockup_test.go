package gate

import (
	"strings"
	"testing"
)

// Bite: one changed byte in the library file is red, with where it is.
func TestBite_cssFromMockup(t *testing.T) {
	want := ":root{--tc-ink:#17211C}\n.tc-btn{height:26px}\n"
	if err := compareCut("x.css", want, want); err != nil {
		t.Fatalf("same text is red: %v", err)
	}
	got := strings.Replace(want, "26px", "28px", 1)
	err := compareCut("x.css", want, got)
	if err == nil || !strings.Contains(err.Error(), "line 2") {
		t.Fatalf("changed byte not caught with its line: %v", err)
	}
	if err := compareCut("x.css", want, want+"\n"); err == nil {
		t.Fatal("extra byte at the end not caught")
	}
}

func TestForeignUICSS(t *testing.T) {
	if got := foreignUICSS(`.page:has(.tc-menu, .tc-scrim) { position: relative }
.page > .box { padding: 0 }`); len(got) != 0 {
		t.Errorf("a page's own box is refused: %v", got)
	}
	if got := foreignUICSS(`/* c */ .page .tc-shell { min-height: 0 }`); len(got) != 1 {
		t.Errorf("a rule that styles a widget passes: %v", got)
	}
}

func TestBite_everyWidgetHasRecord(t *testing.T) {
	root, err := FindModuleRoot()
	if err != nil {
		t.Fatal(err)
	}
	dir := Paths(root).WidgetDir
	src, err := renderTypes(dir)
	if err != nil || len(src) == 0 {
		t.Fatalf("no widgets read from the source: %v", err)
	}
	all := map[string]int{}
	for _, n := range src {
		all[n] = 1
	}
	if err := CheckEveryWidgetHasRecord(dir, src, all); err != nil {
		t.Fatalf("complete list is red: %v", err)
	}
	delete(all, src[0])
	if err := CheckEveryWidgetHasRecord(dir, src, all); err == nil {
		t.Fatal("a widget without a record passes")
	}
	if err := CheckEveryWidgetHasRecord(dir, src[1:], map[string]int{}); err == nil {
		t.Fatal("a widget missing from refs.Widgets passes")
	}
}
