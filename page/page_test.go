package page_test

import (
	"strings"
	"testing"

	"carrel.dev/typecase/page"
	w "carrel.dev/typecase/widget"
)

// TestDocument: the page holds its head, the sprite, the root and the forms,
// and nothing inline a strict content policy would refuse.
func TestDocument(t *testing.T) {
	d := page.Document{Title: "Contacts", Lang: "ru", Assets: page.At("/static/tc/"),
		Root:  w.Page{Items: []w.PagePart{w.Msg{Parts: []w.MsgPart{w.Text("Hello")}}}},
		Forms: []w.PostForm{{ID: "post-x", Action: "/x", Fields: []w.Hidden{{Name: "csrf", Value: "1"}}}}}
	s, err := d.Render()
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`<!doctype html><html lang="ru"`, `<title>Contacts</title>`,
		`<script src="/static/tc/js/theme.js"></script><link rel="stylesheet" href="/static/tc/css/typecase.css">`,
		`<symbol`, `class="tc-root"`, `id="post-x"`, `<script src="/static/tc/js/tc.js"></script></body></html>`} {
		if !strings.Contains(s, want) {
			t.Errorf("the page has no %s", want)
		}
	}
	if strings.Contains(s, "<script>") || strings.Contains(s, "<style") || strings.Contains(s, " style=") || strings.Contains(s, " onclick=") {
		t.Error("the page has something inline")
	}
	if strings.Index(s, "theme.js") > strings.Index(s, "typecase.css") {
		t.Error("the theme's script comes after the stylesheet: the first paint would miss the theme")
	}
}
