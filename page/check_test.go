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
