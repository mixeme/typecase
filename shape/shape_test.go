package shape_test

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"testing"

	"carrel.dev/typecase/internal/screens"
	"carrel.dev/typecase/shape"
	w "carrel.dev/typecase/widget"
)

var (
	once    sync.Once
	parsed  *shape.Mockup
	readErr error
)

// mockup is the mockup read once for the package's tests: a Check does not
// change it, and reading 2.7 MB in each test cost two seconds (the review of
// 14.10).
func mockup(t *testing.T) *shape.Mockup {
	t.Helper()
	once.Do(func() {
		var b []byte
		b, readErr = os.ReadFile(filepath.Join("..", "docs", "visual", "carrel-ui-mockups.html"))
		if readErr == nil {
			parsed, readErr = shape.Parse(string(b))
		}
	})
	if readErr != nil {
		t.Fatal(readErr)
	}
	return parsed
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
	link := strings.Replace(page, `<button class="tc-btn" type="button"><svg viewBox="0 0 24 24"><use href="#i-import"/></svg>Import</button>`,
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
	cut("Print", `<button class="tc-btn is-3rd" type="button"><svg viewBox="0 0 24 24"><use href="#i-print"/></svg>`, "</button>")
	cut("Related", `<div class="tc-sec"><div class="tc-rubric">Related</div>`, "</ul></div>")
	cut("the time of When", `<span class="tc-mono">10:00 – 11:00</span>`, "</span>")
	cut("the badge of When", `<span class="tc-badge" title="Created in another time zone">`, "</span>")
}

// Carrel's task 161: an element with hidden is not drawn — a field
// «Custom…» hidden at a reminder of a ready step is no part of the frame,
// shown it is an extra field. A place a script opens, drawn hidden by the
// frame (the progress bar), is still the frame's.
func TestHiddenIsNotDrawn(t *testing.T) {
	m := mockup(t)
	r, ok := screens.Find("s-tasks-3")
	if !ok {
		t.Fatal("no s-tasks-3")
	}
	page := r.Screen().Render()
	first := `<button class="tc-in is-pick" type="button">1 day before</button></div>`
	i := strings.Index(page, first)
	if i < 0 {
		t.Fatal("no first reminder in s-tasks-3")
	}
	i += len(first)
	field := func(attr string) string {
		return `<div class="tc-field tc-field--xs"` + attr + `><label>How long</label><input class="tc-in is-mono" type="text" name="reminder.1.count" value="" aria-label="How long"></div>`
	}
	if err := m.Check("s-tasks-3", page[:i]+field(" hidden")+page[i:]); err != nil {
		t.Errorf("a hidden field «Custom…»: %v", err)
	}
	if err := m.Check("s-tasks-3", page[:i]+field("")+page[i:]); err == nil {
		t.Error("a field «Custom…» shown passes; want an extra field")
	}

	r, ok = screens.Find("states-24")
	if !ok {
		t.Fatal("no states-24")
	}
	page = r.Screen().Render()
	bar := `<div class="tc-runbar"`
	j := strings.Index(page, bar)
	if j < 0 {
		t.Fatal("no progress bar in states-24")
	}
	k := j + strings.Index(page[j:], "</div>") + len("</div>")
	if err := m.Check("states-24", page[:j]+page[k:]); err == nil {
		t.Error("states-24 without its hidden progress bar passes; want it missing")
	}
}

// The review of 14.10: the rules no test held. A hidden field of a form, a
// tick's box, the options of a select are no drawing — one more or one less
// passes; a select is opaque, whatever stands in it; a screen is one element.
func TestDroppedAndOpaque(t *testing.T) {
	m := mockup(t)
	find := func(needle string) (string, string) {
		t.Helper()
		for _, r := range screens.Routes {
			if p := r.Screen().Render(); strings.Contains(p, needle) {
				return r.Frame, p
			}
		}
		t.Fatalf("no screen with %s", needle)
		return "", ""
	}
	after := func(page, open string) int {
		i := strings.Index(page, open)
		return i + strings.Index(page[i:], ">") + 1
	}
	pass := func(id, page, what string) {
		t.Helper()
		if err := m.Check(id, page); err != nil {
			t.Errorf("%s (%s): %v", what, id, err)
		}
	}

	id, page := find("<form")
	j := after(page, "<form")
	pass(id, page[:j]+`<input type="hidden" name="more" value="1">`+page[j:], "a hidden field more")

	id, page = find(`<input class="tc-check__box"`)
	i := strings.Index(page, `<input class="tc-check__box"`)
	j = i + strings.Index(page[i:], ">") + 1
	pass(id, page[:i]+page[j:], "a tick without its box")

	id, page = find("<select")
	j = after(page, "<select")
	pass(id, page[:j]+`<option value="z">another</option>`+page[j:], "another option")
	pass(id, page[:j]+`<optgroup label="g"><option value="y">in a group</option></optgroup>`+page[j:], "a select's other insides")

	if err := m.Check(id, page+page); err == nil {
		t.Error("two top elements pass")
	}
}
