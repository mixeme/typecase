package widget

import (
	"strings"
	"testing"
)

// TestMenuRadioOf: a radio item outside its form names the form (Carrel's
// task 203); without Of it belongs to the form it stands in.
func TestMenuRadioOf(t *testing.T) {
	m := Menu{Items: []MenuPart{
		MenuItem{Label: "nextcloud · Дела", On: true, Name: "list", Value: "nextcloud/dela", Of: "task-edit"},
		MenuItem{Label: "UTC", Name: "tz", Value: "UTC"},
	}}.Render()
	if !strings.Contains(m, `<input class="tc-seg__box" type="radio" form="task-edit" role="menuitemradio" name="list" value="nextcloud/dela" checked>`) {
		t.Errorf("the radio does not name its form:\n%s", m)
	}
	if !strings.Contains(m, `<input class="tc-seg__box" type="radio" role="menuitemradio" name="tz" value="UTC">`) {
		t.Errorf("a radio without Of names a form:\n%s", m)
	}
}
