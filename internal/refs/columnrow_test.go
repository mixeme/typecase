package refs_test

import (
	"strings"
	"testing"

	w "carrel.dev/typecase/widget"
)

// TestColumnRowLink: a row with a tick is not a link, its name is (Carrel's
// task 123); a row without one is the link itself.
func TestColumnRowLink(t *testing.T) {
	task := w.ColumnRow{Tick: &w.Tick{Label: "Done"}, Cells: []w.ColumnCell{{Column: w.ColumnName, Text: "Купить билеты"}}, Href: "/app/tasks/x"}.Render()
	if !strings.HasPrefix(task, `<div class="tc-row`) || !strings.Contains(task, `<a class="tc-row__title tc-col--name" href="/app/tasks/x">Купить билеты</a>`) {
		t.Errorf("a row with a tick: %s", task)
	}
	card := w.ColumnRow{Cells: []w.ColumnCell{{Column: w.ColumnName, Text: "Анна"}}, Href: "/app/contacts/x"}.Render()
	if !strings.HasPrefix(card, `<a class="tc-row`) || strings.Count(card, "<a") != 1 {
		t.Errorf("a row without one: %s", card)
	}
}
