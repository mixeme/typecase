package catalog

import (
	"slices"
	"testing"

	"carrel.dev/typecase/internal/browser"
)

// The catalog's width switch has the steps of the ladder the checks use.
func TestLadderIsTheLadder(t *testing.T) {
	if !slices.Equal(Ladder, browser.MockupLadderWidths) {
		t.Fatalf("catalog ladder %v, browser ladder %v", Ladder, browser.MockupLadderWidths)
	}
}

// Every component and fragment frame finds a widget of the catalog to stand
// next to.
func TestFramesBesideWidgets(t *testing.T) {
	root, err := moduleRoot()
	if err != nil {
		t.Fatal(err)
	}
	frames, err := Frames(root)
	if err != nil {
		t.Fatal(err)
	}
	widgets := map[string]bool{}
	for _, e := range Entries() {
		widgets[e.Widget] = true
	}
	n := 0
	for _, f := range frames {
		n++
		if !widgets[f.Block] {
			t.Errorf("%s (%s): its block %q has no widget in the catalog", f.ID, f.Caption, f.Block)
		}
	}
	if n != 153 {
		t.Errorf("%d component and fragment frames, want 153 (Р-16; the 5 candidates are skipped)", n)
	}
}
