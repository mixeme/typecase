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

// TestMenuRadioRequired: a radio item of a choice the form cannot be sent
// without carries required (Carrel's task 238); the others do not.
func TestMenuRadioRequired(t *testing.T) {
	m := Menu{Items: []MenuPart{
		MenuItem{Label: "nextcloud · Записи", Name: "into", Value: "nextcloud/zapisi", Required: true},
		MenuItem{Label: "UTC", Name: "tz", Value: "UTC"},
	}}.Render()
	if !strings.Contains(m, `name="into" value="nextcloud/zapisi" required>`) {
		t.Errorf("the radio is not required:\n%s", m)
	}
	if strings.Contains(m, `value="UTC" required`) {
		t.Errorf("a radio without Required is required:\n%s", m)
	}
}

// TestMenuToOf: an item sending a form to its own address names that form
// when it stands outside it (Carrel's task 247); without Form it sends the
// form it stands in.
func TestMenuToOf(t *testing.T) {
	m := Menu{Items: []MenuPart{
		MenuItem{Label: "From Files", To: "/choose", Form: "note-new"},
		MenuItem{Label: "Upload", To: "/upload"},
	}}.Render()
	if !strings.Contains(m, `type="submit" form="note-new" formaction="/choose" role="menuitem">`) {
		t.Errorf("the item does not name its form:\n%s", m)
	}
	if !strings.Contains(m, `type="submit" formaction="/upload" role="menuitem">`) {
		t.Errorf("an item without Form names a form:\n%s", m)
	}
}
