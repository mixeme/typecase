package shape_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"carrel.dev/typecase/internal/screens"
	"carrel.dev/typecase/shape"
	w "carrel.dev/typecase/widget"
)

func mockup(t *testing.T) *shape.Mockup {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("..", "docs", "visual", "carrel-ui-mockups.html"))
	if err != nil {
		t.Fatal(err)
	}
	m, err := shape.Parse(string(b))
	if err != nil {
		t.Fatal(err)
	}
	return m
}

// TestScreensHaveTheirShape: every demo screen, built of widgets, has the
// shape of its frame.
func TestScreensHaveTheirShape(t *testing.T) {
	m := mockup(t)
	for _, r := range screens.Routes {
		if err := m.Check(r.Frame, r.Screen().Render()); err != nil {
			t.Error(err)
		}
	}
}

// TestShapeNotWords: the same shape with other words, more rows and the
// meaning of HTML passes; a missing or extra element, another class, does
// not.
func TestShapeNotWords(t *testing.T) {
	m := mockup(t)
	r, ok := screens.Find("s-notes-1")
	if !ok {
		t.Fatal("no s-notes-1")
	}
	page := r.Screen().Render()
	i := strings.Index(page, `<div class="tc-row tc-row--excerpt">`)
	if i < 0 {
		t.Fatal("no plain excerpt row in s-notes-1")
	}
	// A row of the frame's with other words, twice, as links.
	one := page[i : i+strings.Index(page[i:], "</div>")+len("</div>")]
	row := strings.Replace(strings.Replace(one, `<div class="tc-row tc-row--excerpt">`, `<a class="tc-row tc-row--excerpt" href="/app/notes/x/y/z">`, 1), "</div>", "</a>", 1)
	for _, from := range []string{"Что почитать про CRDT", "Кляйн"} {
		row = strings.Replace(row, from, "Другие слова", 1)
	}
	more := page[:i] + row + row + page[i:]
	if err := m.Check("s-notes-1", more); err != nil {
		t.Errorf("more rows, other words, rows as links: %v", err)
	}
	link := strings.Replace(page, `<button class="tc-btn" type="button"><svg viewBox="0 0 24 24"><use href="#i-up"/></svg>Import</button>`,
		w.Button{Label: "Import", Icon: "up", Href: "/app/notes/import"}.Render(), 1)
	if link == page {
		t.Fatal("no Import button in s-notes-1")
	}
	if err := m.Check("s-notes-1", link); err != nil {
		t.Errorf("a button as a link: %v", err)
	}
	primary := `<button class="tc-btn is-primary" type="button"><svg viewBox="0 0 24 24"><use href="#i-plus"/></svg>New note<svg viewBox="0 0 24 24"><use href="#i-chev"/></svg></button>`
	gone := strings.Replace(page, primary, "", 1)
	if gone == page {
		t.Fatal("no New note button in s-notes-1's head")
	}
	if err := m.Check("s-notes-1", gone); err == nil {
		t.Error("a missing button passes")
	}
	other := strings.Replace(page, `class="tc-btn is-primary"`, `class="tc-btn"`, 1)
	if err := m.Check("s-notes-1", other); err == nil {
		t.Error("another class passes")
	}
}
