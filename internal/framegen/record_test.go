package framegen

import (
	"path/filepath"
	"strings"
	"testing"
)

// TestRecordLeavesMixinsOut: a record is the block as it prints by itself —
// the comparison takes tc-grow, tc-push off the frame's root, so the draft
// has no Place for them (review finding 1).
func TestRecordLeavesMixinsOut(t *testing.T) {
	for _, c := range []struct {
		frame, class string
		n            int
		place        string
	}{
		{"s-backup-2", "tc-msg", 1, "Grown"},
		{"s-davloom-1", "tc-btn", 6, "Push"},
	} {
		got, err := Record("../..", c.frame, c.class, c.n)
		if err != nil {
			t.Fatal(err)
		}
		if strings.Contains(got, c.place) {
			t.Errorf("%s .%s #%d keeps the root's %s: %s", c.frame, c.class, c.n, c.place, got)
		}
	}
}

// TestSheetWithoutPartsIsAGap: a sheet drawn without the head, title or body
// widget.Sheet prints says which is missing instead of a nil pointer (review
// finding 2).
func TestSheetWithoutPartsIsAGap(t *testing.T) {
	en, err := loadEnums(filepath.Join("..", "..", "widget"))
	if err != nil {
		t.Fatal(err)
	}
	g := &gen{en: en, faces: en.faces()}
	for part, src := range map[string]string{
		"tc-sheet__head":  `<div class="tc-sheet"><div class="tc-sheet__body"></div></div>`,
		"tc-sheet__title": `<div class="tc-sheet"><div class="tc-sheet__head"></div><div class="tc-sheet__body"></div></div>`,
		"tc-sheet__body":  `<div class="tc-sheet"><div class="tc-sheet__head"><b class="tc-sheet__title">Send</b></div><div class="tc-sheet__foot"></div></div>`,
	} {
		_, why := g.element(Parse(src))
		if !strings.Contains(why, part) {
			t.Errorf("without %s: %q", part, why)
		}
	}
}
