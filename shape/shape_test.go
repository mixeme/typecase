package shape_test

import (
	"os"
	"path/filepath"
	"reflect"
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
	// Task 39: words and parts that are links stand for the frame's spans.
	r2, _ := screens.Find("s-auth-1")
	auth := r2.Screen().Render()
	linked := strings.Replace(auth, `<span class="tc-auth__link">Forgotten your password?</span>`,
		`<a class="tc-auth__link" href="/forgot">Forgotten your password?</a>`, 1)
	if linked == auth {
		t.Fatal("no auth link in s-auth-1")
	}
	if err := m.Check("s-auth-1", linked); err != nil {
		t.Errorf("an auth link with its address: %v", err)
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

// TestMeaningIsNoDifference: every demo screen with every meaning field set
// (links, forms, names — Р-25, Р-27) still has its frame's shape, so a
// client's screen with its addresses passes where the frame draws none.
func TestMeaningIsNoDifference(t *testing.T) {
	m := mockup(t)
	for _, r := range screens.Routes {
		v := clone(reflect.ValueOf(r.Screen()))
		fill(v, false)
		if err := m.Check(r.Frame, v.Interface().(interface{ Render() string }).Render()); err != nil {
			t.Error(err)
		}
	}
}

var meaningFields = map[string]bool{"Href": true, "Form": true, "Name": true, "Action": true}

func clone(v reflect.Value) reflect.Value {
	switch v.Kind() {
	case reflect.Pointer:
		if v.IsNil() {
			return v
		}
		c := reflect.New(v.Elem().Type())
		c.Elem().Set(clone(v.Elem()))
		return c
	case reflect.Interface:
		if v.IsNil() {
			return v
		}
		c := reflect.New(v.Type()).Elem()
		c.Set(clone(v.Elem()))
		return c
	case reflect.Slice:
		if v.IsNil() {
			return v
		}
		c := reflect.MakeSlice(v.Type(), v.Len(), v.Len())
		for i := 0; i < v.Len(); i++ {
			c.Index(i).Set(clone(v.Index(i)))
		}
		return c
	case reflect.Struct:
		c := reflect.New(v.Type()).Elem()
		c.Set(v)
		for i := 0; i < v.NumField(); i++ {
			if c.Field(i).CanSet() {
				c.Field(i).Set(clone(v.Field(i)))
			}
		}
		return c
	}
	return v
}

// fill sets the meaning fields of v. A struct whose Action it sets is a
// form, and a form within a form is not HTML (the inner end closes the outer):
// inside one, Action is left empty.
func fill(v reflect.Value, inForm bool) {
	switch v.Kind() {
	case reflect.Pointer:
		if !v.IsNil() {
			fill(v.Elem(), inForm)
		}
	case reflect.Interface:
		if !v.IsNil() {
			c := reflect.New(v.Elem().Type()).Elem()
			c.Set(v.Elem())
			fill(c, inForm)
			v.Set(c)
		}
	case reflect.Struct:
		if f := v.FieldByName("Action"); f.IsValid() && f.Kind() == reflect.String && f.CanSet() {
			if f.String() == "" && !inForm {
				f.SetString("m")
			}
			inForm = inForm || f.String() != ""
		}
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			name := v.Type().Field(i).Name
			if !v.Type().Field(i).IsExported() || !f.CanSet() || name == "Action" {
				continue
			}
			if meaningFields[name] && f.Kind() == reflect.String && f.String() == "" {
				f.SetString("m")
				continue
			}
			fill(f, inForm)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			fill(v.Index(i), inForm)
		}
	}
}

// Carrel's task 48: a row's avatar with a photo is the frame's avatar with
// initials — the picture is data.
func TestAvatarPhotoIsData(t *testing.T) {
	m := mockup(t)
	r, ok := screens.Find("s-contacts-1")
	if !ok {
		t.Fatal("no s-contacts-1")
	}
	page := r.Screen().Render()
	initials := w.Avatar{Initials: "АК"}.Render()
	if !strings.Contains(page, initials) {
		t.Fatal("no avatar АК in s-contacts-1")
	}
	photo := strings.Replace(page, initials, w.Avatar{Initials: "АК", Src: "/app/contacts/x/y/z/photo"}.Render(), 1)
	if err := m.Check("s-contacts-1", photo); err != nil {
		t.Errorf("an avatar with a photo: %v", err)
	}
}

// Carrel's task 49: a head's «⋯» with its menu is its button: the frame
// draws the closed menu (details), a screen with the bare button passes too.
func TestClosedMenuIsItsButton(t *testing.T) {
	m := mockup(t)
	r, ok := screens.Find("s-contacts-6")
	if !ok {
		t.Fatal("no s-contacts-6")
	}
	page := r.Screen().Render()
	i := strings.Index(page, `<details class="tc-drop">`)
	if i < 0 {
		t.Fatal("no closed menu in s-contacts-6")
	}
	j := strings.Index(page[i:], "</details>") + i + len("</details>")
	bare := page[:i] + w.Button{Icon: "dots", Title: "More"}.Render() + page[j:]
	if err := m.Check("s-contacts-6", bare); err != nil {
		t.Errorf("«⋯» without its menu: %v", err)
	}
	if err := m.Check("s-contacts-6", page); err != nil {
		t.Errorf("«⋯» with its menu: %v", err)
	}
}

// Carrel's task 61: a repeating row of rows is compared at any length — one
// month of a timeline where the frame draws two, six fields of a card where
// it draws eight.
func TestRepeatsAnyLength(t *testing.T) {
	m := mockup(t)
	r, ok := screens.Find("s-contacts-7")
	if !ok {
		t.Fatal("no s-contacts-7")
	}
	page := r.Screen().Render()
	g := `<div class="tc-list__group">`
	first := strings.Index(page, g)
	second := strings.Index(page[first+len(g):], g) + first + len(g)
	if first < 0 || second <= first {
		t.Fatal("no two months in s-contacts-7")
	}
	end := strings.Index(page[second:], `<div class="tc-empty`) + second
	one := page[:second] + page[end:]
	if err := m.Check("s-contacts-7", one); err != nil {
		t.Errorf("one month: %v", err)
	}
	r, _ = screens.Find("s-contacts-6")
	card := r.Screen().Render()
	i := strings.Index(card, "<dt>Address · work</dt>")
	j := strings.Index(card, "<dt>Tags</dt>")
	if i < 0 || j < i {
		t.Fatal("no address and birthday in s-contacts-6")
	}
	if err := m.Check("s-contacts-6", card[:i]+card[j:]); err != nil {
		t.Errorf("six fields: %v", err)
	}
}

// Carrel's task 75: a template is not drawn — the reminder a script copies in
// is no field of the row. And a day without events is not printed (the
// owner's decision of 02.10): a week that adds one is red.
func TestTemplatesAndEmptyDays(t *testing.T) {
	m := mockup(t)
	r, _ := screens.Find("s-agenda-3")
	edit := r.Screen().Render()
	add := "Add reminder</button></div>"
	if !strings.Contains(edit, add) {
		t.Fatal("no Add reminder in s-agenda-3")
	}
	tpl := strings.Replace(edit, add, add+`<template id="tpl-reminder"><div class="tc-field tc-field--md"><label>Reminder</label><select class="tc-in is-pick"></select></div></template>`, 1)
	if err := m.Check("s-agenda-3", tpl); err != nil {
		t.Errorf("a template after Add reminder: %v", err)
	}

	r, _ = screens.Find("s-agenda-1")
	week := r.Screen().Render()
	g := `<div class="tc-list__group">`
	first := strings.Index(week, g)
	second := strings.Index(week[first+len(g):], g) + first + len(g)
	if first < 0 || second <= first {
		t.Fatal("no two days in s-agenda-1")
	}
	empty := `<div class="tc-list__group"><span class="tc-rubric">Wednesday · 12 August</span><span class="tc-num">0</span></div>` +
		`<div class="tc-empty tc-empty--t12-pad14-20">Nothing this day.</div>`
	if err := m.Check("s-agenda-1", week[:second]+empty+week[second:]); err == nil {
		t.Error("an empty day passes")
	}
	if err := m.Check("s-agenda-1", week); err != nil {
		t.Errorf("the week: %v", err)
	}
}

// Carrel's task 146: a datalist is not drawn — the choices of a tag field's
// line on a page without a script (TagAdd.Options) are no part of the field.
func TestDatalistIsNotDrawn(t *testing.T) {
	m := mockup(t)
	for _, id := range []string{"s-tasks-3", "s-tasks-14", "s-notes-20", "s-notes-30"} {
		r, ok := screens.Find(id)
		if !ok {
			t.Fatalf("no screen %s", id)
		}
		page := r.Screen().Render()
		add := ` autocomplete="off">`
		i := strings.Index(page, `name="related.add"`)
		if i < 0 {
			t.Fatalf("no related.add in %s", id)
		}
		j := strings.Index(page[i:], add)
		if j < 0 {
			t.Fatalf("no end of the line in %s", id)
		}
		j += i + len(add)
		list := `<datalist id="tc-list-related.add"><option value="c41a" label="Собрать чемодан"></option></datalist>`
		if err := m.Check(id, page[:j]+list+page[j:]); err != nil {
			t.Errorf("a datalist beside the line: %v", err)
		}
	}
}

// Carrel's task 76: what the frame draws, the screen has — a button of the
// head, a section of the card, a piece of a field's value missing is red.
func TestMissingIsRed(t *testing.T) {
	m := mockup(t)
	r, _ := screens.Find("s-agenda-2")
	card := r.Screen().Render()
	cut := func(what, from, to string) {
		t.Helper()
		i := strings.Index(card, from)
		if i < 0 {
			t.Fatalf("no %s in s-agenda-2", what)
		}
		j := strings.Index(card[i:], to) + i + len(to)
		if err := m.Check("s-agenda-2", card[:i]+card[j:]); err == nil {
			t.Errorf("a card without %s passes", what)
		}
	}
	cut("Print", `<button class="tc-btn" type="button"><svg viewBox="0 0 24 24"><use href="#i-print"/></svg>`, "</button>")
	cut("Related", `<div class="tc-sec"><div class="tc-rubric">Related</div>`, "</ul></div>")
	cut("the time of When", `<span class="tc-mono">10:00 – 11:00</span>`, "</span>")
	cut("the badge of When", `<span class="tc-badge" title="Created in another time zone">`, "</span>")
}
