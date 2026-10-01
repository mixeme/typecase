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
		fill(v)
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

func fill(v reflect.Value) {
	switch v.Kind() {
	case reflect.Pointer:
		if !v.IsNil() {
			fill(v.Elem())
		}
	case reflect.Interface:
		if !v.IsNil() {
			c := reflect.New(v.Elem().Type()).Elem()
			c.Set(v.Elem())
			fill(c)
			v.Set(c)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			if !v.Type().Field(i).IsExported() || !f.CanSet() {
				continue
			}
			if meaningFields[v.Type().Field(i).Name] && f.Kind() == reflect.String && f.String() == "" {
				f.SetString("m")
				continue
			}
			fill(f)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			fill(v.Index(i))
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

// Carrel's task 49: a head's «⋯» with its menu, closed, is the frame's «⋯».
func TestClosedMenuIsItsButton(t *testing.T) {
	m := mockup(t)
	r, ok := screens.Find("s-contacts-6")
	if !ok {
		t.Fatal("no s-contacts-6")
	}
	page := r.Screen().Render()
	dots := w.Button{Icon: "dots", Title: "More"}.Render()
	if !strings.Contains(page, dots) {
		t.Fatal("no «⋯» in s-contacts-6")
	}
	menu := w.Drop{Trigger: w.Button{Icon: "dots", Title: "More"}, Menu: w.Menu{Items: []w.MenuPart{w.MenuItem{Label: "Export .vcf", Href: "/app/contacts/x/y/z.vcf"}}}}.Render()
	if err := m.Check("s-contacts-6", strings.Replace(page, dots, menu, 1)); err != nil {
		t.Errorf("«⋯» with its menu: %v", err)
	}
}
