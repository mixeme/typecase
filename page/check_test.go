package page_test

import (
	"strings"
	"testing"

	"carrel.dev/typecase/page"
	w "carrel.dev/typecase/widget"
)

// TestCheck: a page whose parts each render right but name each other
// wrong — a button submits a form the page has not, a panel takes a file
// into a form with no file field, «Add» copies a template that is not
// there, a link asks a zone no part carries, an option shows a field its
// form has not, a field is derived from a name its form has not, an id is given twice — is told so, part by part; the same
// page put together right has nothing to say.
func TestCheck(t *testing.T) {
	// the parts as a consumer puts them on a page of its own
	pageOf := func(right bool) string {
		forms := []w.PostForm{{ID: "post-import", Action: "/import", File: "file"}, {ID: "post-attach", Action: "/attach"}}
		pick := w.Pick{Name: "remind", Label: "Remind", Options: []w.PickOption{{Value: "10m", Label: "10 minutes before"}, {Value: "custom", Label: "Custom…", Shows: "remind.count"}}}
		count := w.Input{Name: "remind.cnt", Label: "How long"}
		tpl, zone, from := "tpl-tels", "", "nam"
		if right {
			forms[1].File = "file"
			forms = append(forms, w.PostForm{ID: "post-signout", Action: "/signout"})
			count.Name = "remind.count"
			tpl, zone, from = "tpl-tel", "side", "name"
		}
		form := w.Form{Action: "/event", Items: []w.FormPart{
			w.Field{Label: "Remind", Parts: []w.FieldPart{pick}},
			w.Field{Label: "How long", Parts: []w.FieldPart{count}},
			w.Field{Label: "Name", Parts: []w.FieldPart{w.Input{Name: "name", Label: "Name"}}},
			w.Field{Label: "Address", Parts: []w.FieldPart{w.Input{Name: "address", Label: "Address", From: from}}},
			w.FieldSet{Items: []w.FieldSetPart{w.Template{ID: "tpl-tel"}, w.Button{Label: "Add phone", Submit: true, Name: "add", Value: "tel", Adds: tpl}}},
		}}
		s := `<!doctype html><html><body><div class="tc-root">` +
			w.Button{Label: "Sign out", Form: "post-signout"}.Render() +
			w.Scrim{ID: "drop", Closed: true, Drop: "post-import", Content: w.Sheet{Title: "Drop to import"}}.Render() +
			w.Side{Drop: "post-attach", Zone: zone, Items: []w.SidePart{w.Msg{Parts: []w.MsgPart{w.Text("Синк-ап")}}}}.Render() +
			w.Form{Action: "/x", Target: "side"}.Render() + form.Render()
		for _, f := range forms {
			s += f.Render()
		}
		return s + `</div></body></html>`
	}
	got, err := page.Check(pageOf(false))
	if err != nil {
		t.Fatal(err)
	}
	var said []string
	for _, p := range got {
		said = append(said, p.String())
	}
	want := []string{
		`button.tc-btn form="post-signout": names nothing on the page`,
		`div.tc-side data-tc-drop="post-attach": names a form with no file field: the file goes nowhere`,
		`form.tc-form data-tc-target="side": names no zone on the page: what is pressed goes the plain way`,
		`option data-tc-shows="remind.count": names no field of its form`,
		`input.tc-in data-tc-from="nam": names no field of its form`,
		`button.tc-btn data-tc-add="tpl-tels": names nothing on the page`,
	}
	if strings.Join(said, "\n") != strings.Join(want, "\n") {
		t.Errorf("a page put together wrong:\n%s\nwant\n%s", strings.Join(said, "\n"), strings.Join(want, "\n"))
	}
	got, err = page.Check(pageOf(true))
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range got {
		t.Errorf("a page put together right: %s", p)
	}
	d := page.Document{Title: "Event", Assets: page.At("/static/tc"), Forms: []w.PostForm{{ID: "post-signout", Action: "/signout"}},
		Root: w.Page{Items: []w.PagePart{w.Msg{Parts: []w.MsgPart{w.Text("Saved")}}}}}
	if got, err := d.Check(); err != nil || len(got) > 0 {
		t.Errorf("a document's own check: %v %v; want nothing", got, err)
	}
	twice, err := page.Check(`<!doctype html><body><p id="a"></p><p id="a"></p><label for="a">A</label></body>`)
	if err != nil {
		t.Fatal(err)
	}
	if len(twice) != 2 || twice[0].Why != "the id is given twice: the script finds the first only" || twice[1].Attr != "for" {
		t.Errorf("an id given twice and a label of a paragraph: %v", twice)
	}
}

// TestCheckNamesOfFields (review findings 7 and 10): the other names a part
// gives — the fields a change brings a question for, the day a time carries
// over, the field a hint says, the elements a form recounts — hold; a submit
// button of no form is told so. A stream line's data-tc-changed is the name
// of its event, not of a field (Carrel's task 247).
func TestCheckNamesOfFields(t *testing.T) {
	const wrong = `<!doctype html><body>
<form action="/c"><input name="url"><div class="tc-field" data-tc-changed="login" hidden></div>
<input type="time" name="end" data-tc-day="end.dat"><span data-tc-slug-of="adress"></span></form>
<form action="/a" data-tc-recount="attach-lin"><select name="x"></select></form>
<div class="tc-runbar" data-tc-stream="/poll/events" data-tc-changed="poll" data-tc-ended="done"></div>
<div class="tc-sheet"><button class="tc-btn" type="submit">Send</button><button class="tc-btn" type="button">Cancel</button></div>
<template><button>Add</button></template></body>`
	got, err := page.Check(wrong)
	if err != nil {
		t.Fatal(err)
	}
	var said []string
	for _, p := range got {
		said = append(said, p.String())
	}
	want := []string{
		`div.tc-field data-tc-changed="login": names no field of its form`,
		`input data-tc-day="end.dat": names no field of its form`,
		`span data-tc-slug-of="adress": names no field of its form`,
		`form data-tc-recount="attach-lin": names nothing on the page`,
		`button.tc-btn type="submit": submits no form: pressed, it does nothing`,
	}
	if strings.Join(said, "\n") != strings.Join(want, "\n") {
		t.Errorf("names that do not hold:\n%s\nwant\n%s", strings.Join(said, "\n"), strings.Join(want, "\n"))
	}
	right := strings.NewReplacer(`"login"`, `"url"`, `"end.dat"`, `"end"`, `"adress"`, `"url"`, `"attach-lin"`, `"x"`,
		`<button class="tc-btn" type="submit">Send</button>`, `<button class="tc-btn" type="submit" form="send">Send</button><form id="send" action="/s"></form>`).Replace(wrong)
	right = strings.Replace(right, `<select name="x">`, `<select name="x" id="x">`, 1)
	if got, err := page.Check(right); err != nil || len(got) > 0 {
		t.Errorf("the same page put together right: %v %v", got, err)
	}
}

// TestCheckMenuToOf: Carrel's task 247 — an item of a menu outside a form,
// sending that form to its own address, submits it.
func TestCheckMenuToOf(t *testing.T) {
	menu := w.Menu{Items: []w.MenuPart{w.MenuItem{Label: "From Files", To: "/choose", Form: "note-new"}}}.Render()
	got, err := page.Check(`<!doctype html><body>` + menu + `<form id="note-new" action="/save" method="post"></form></body>`)
	if err != nil || len(got) > 0 {
		t.Errorf("an item naming its form: %v %v", got, err)
	}
}

// TestCheckDialogPostForm (review 31.10): a dialog that is a form keeps its
// form when a footer button's hidden form is one of its parts — that form
// stands after the dialog's, not in it, and both buttons send.
func TestCheckDialogPostForm(t *testing.T) {
	d := w.Dialog{Action: "/login", Items: []w.DialogPart{
		w.Field{Label: "Name", Parts: []w.FieldPart{w.Input{Name: "name", Label: "Name"}}},
		w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Sign in", Kind: w.ButtonPrimary, Submit: true}, w.Button{Label: "Sign out", Form: "post-signout"}}},
		w.PostForm{ID: "post-signout", Action: "/logout"},
	}}.Render()
	if !strings.HasPrefix(d, `<form class="tc-dialog" action="/login"`) {
		t.Errorf("the dialog lost its form:\n%s", d)
	}
	got, err := page.Check(`<!doctype html><body>` + d + `</body>`)
	if err != nil || len(got) > 0 {
		t.Errorf("a dialog with a hidden form: %v %v", got, err)
	}
}
