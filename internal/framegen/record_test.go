package framegen

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"carrel.dev/typecase/internal/fragment"
)

// TestRecordCountsAsFragment: the n-th element of a block the record tool
// drafts from is the n-th the comparison means (fragment.Ref), in every
// frame, for every block class — counted otherwise, a redraft took a line
// laid out by another block for the line it should have (20.11.2026).
func TestRecordCountsAsFragment(t *testing.T) {
	m, err := fragment.Load("../..")
	if err != nil {
		t.Fatal(err)
	}
	mock, err := os.ReadFile(filepath.Join("../..", fragment.Source))
	if err != nil {
		t.Fatal(err)
	}
	classes := map[string]bool{}
	for _, c := range regexp.MustCompile(`\btc-[a-z0-9]+(?:-[a-z0-9]+)*\b`).FindAllString(string(mock), -1) {
		// tc-root is a block only as a page's root, where both count one;
		// elsewhere it is the frame's box, and no record stands for it.
		if !strings.Contains(c, "--") && c != "tc-root" {
			classes[c] = true
		}
	}
	roots := map[string]*El{}
	checked := 0
	bad := 0
	for cls := range classes {
		want := map[string]int{}
		for _, s := range m.Shapes(cls) {
			for _, r := range s.Refs {
				want[r.Frame]++
			}
		}
		for _, f := range m.Frames {
			root, ok := roots[f.ID]
			if !ok {
				if root, err = frameRoot(string(mock), f.ID); err != nil {
					t.Fatal(err)
				}
				roots[f.ID] = root
			}
			checked++
			if got := len(occurrences(root, cls)); got != want[f.ID] && bad < 20 {
				bad++
				t.Errorf("%s .%s: the record tool counts %d, the comparison %d", f.ID, cls, got, want[f.ID])
			}
		}
	}
	if len(classes) < 100 || checked < 10000 {
		t.Errorf("%d classes, %d frame×class pairs: the test lost its mockup", len(classes), checked)
	}
	t.Logf("%d classes, %d frame×class pairs", len(classes), checked)
}

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

// TestInputKeepsItsID: an input's id is what a Copy beside it names
// (data-tc-copy-from) as much as a file field's line; the generator dropped
// it but for the line, and the screen's Copy named nothing (s-profile-4,
// s-backup-13).
func TestInputKeepsItsID(t *testing.T) {
	en, err := loadEnums(filepath.Join("..", "..", "widget"))
	if err != nil {
		t.Fatal(err)
	}
	g := &gen{en: en, faces: en.faces()}
	for src, want := range map[string]string{
		`<input class="tc-in is-mono" type="text" value="z3" aria-label="Password" id="pw-nextcloud" readonly>`: `ID: "pw-nextcloud"`,
		`<input class="tc-in" type="text" id="pick-name" data-tc-filename readonly>`:                            `ID: "pick-name"`,
	} {
		got, why := g.element(Parse(src))
		if why != "" || !strings.Contains(got, want) {
			t.Errorf("%s: want %s in %s %s", src, want, got, why)
		}
	}
}
