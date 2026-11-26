package page_test

import (
	"strings"
	"testing"
	"time"

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

// TestTheme (Carrel's task 186): a theme the server chose is printed on the
// page — light or dark as the theme itself, auto only as the server's mark —
// and nothing is printed when the server did not choose.
func TestTheme(t *testing.T) {
	for theme, want := range map[string]string{
		"dark":  ` data-tc-theme="dark" data-tc-theme-set="dark"`,
		"light": ` data-tc-theme="light" data-tc-theme-set="light"`,
		"auto":  ` data-tc-theme-set="auto"`,
		"":      ``,
	} {
		s, err := page.Document{Assets: page.At("/s/"), Theme: theme}.Render()
		if err != nil {
			t.Fatal(err)
		}
		head := s[:strings.Index(s, "<head>")]
		if !strings.Contains(head, `data-tc-noscript`+want+`>`) && !strings.Contains(head, `data-tc-noscript`+want+` `) {
			t.Errorf("Theme %q: the page opens %s; want %q after data-tc-noscript", theme, head, want)
		}
		if theme == "" && strings.Contains(head, "data-tc-theme") {
			t.Errorf("no theme chosen, yet the page names one: %s", head)
		}
	}
}

// TestZones (Carrel's task 84): the document names the headers of a request
// for a part of the page and the zone that owns the address; a fragment is
// the part asked for and, after it, the other zones in templates; widgets
// carry the names the application gave their zones and targets.
func TestZones(t *testing.T) {
	d := page.Document{Assets: page.At("/s"), Zones: page.Zones{ViewHeader: "X-View", ViewValue: "part", TargetHeader: "X-Target",
		TokenField: "csrf", TokenHeader: "X-Token", Address: "main"}}
	s, err := d.Render()
	if err != nil {
		t.Fatal(err)
	}
	want := `data-tc-noscript data-tc-view-header="X-View" data-tc-view-value="part" data-tc-target-header="X-Target" data-tc-token-field="csrf" data-tc-token-header="X-Token" data-tc-address-zone="main">`
	if !strings.Contains(s, want) {
		t.Errorf("the document does not name its zones' headers:\n%s", s[:300])
	}
	plain, err := page.Document{Assets: page.At("/s")}.Render()
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(plain, "data-tc-view-header") || strings.Contains(plain, "data-tc-address-zone") {
		t.Error("a document without zones names some")
	}
	// the search's suggestions: an address on the root, none without one
	// (Carrel's task 227)
	sug, err := page.Document{Assets: page.At("/s"), Suggest: "/app/search/suggest"}.Render()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(sug, ` data-tc-search-suggest="/app/search/suggest"`) || strings.Contains(plain, "data-tc-search-suggest") {
		t.Errorf("the suggestions' address: %s", sug[:200])
	}
	f := page.Fragment{Body: []page.Part{w.Main{Zone: "main"}},
		Extra: []page.ZonePart{{Zone: "notices", Body: []page.Part{w.Msg{Parts: []w.MsgPart{w.Text("Saved")}}}}}}.Render()
	if !strings.HasPrefix(f, `<div class="tc-main" data-tc-zone="main">`) || !strings.Contains(f, `</div><template data-tc-zone="notices"><div class="tc-msg">Saved</div></template>`) {
		t.Errorf("the fragment: %s", f)
	}
	if got := (page.Fragment{}).Render(); got != "" {
		t.Errorf("an empty fragment prints %q", got)
	}
	shell := w.Shell{Body: w.ShellBody{Rail: []w.RailPart{w.Nav{Items: []w.NavItem{{Label: "Tasks", Href: "/tasks"}}}}, RailZoneName: "rail", RailTarget: "main",
		Main: w.Main{Zone: "main", Target: "panel"}, Side: &w.Side{Zone: "panel", Closed: true}}}.Render()
	for _, want := range []string{`class="tc-shell__rail" data-tc-zone="rail" data-tc-target="main"`, `class="tc-main" data-tc-zone="main" data-tc-target="panel"`,
		`class="tc-side" data-tc-zone="panel" hidden`, `<div class="tc-shell__body">`} {
		if !strings.Contains(shell, want) {
			t.Errorf("the shell has no %s: %s", want, shell)
		}
	}
	// a message with its technical block, printed for a zone (task 83)
	notice := page.Printed(w.PrintMain(w.Msg{Parts: []w.MsgPart{w.Text("Read from the copy")}}, w.MainDiag{Diag: w.Diag{Text: "GET /dav → 503"}})).Render()
	if !strings.HasPrefix(notice, `<div class="tc-msg">Read from the copy</div><div class="tc-main__diag"><div class="tc-diag">`) {
		t.Errorf("the notice: %s", notice)
	}
	run := w.Runbar{Zone: "run", Stream: "/run/events", Changed: "changed", Ended: "ended", Poll: "/run", Every: 5 * time.Second}.Render()
	if !strings.Contains(run, `data-tc-zone="run" data-tc-stream="/run/events" data-tc-changed="changed" data-tc-ended="ended" data-tc-poll="/run" data-tc-every="5000"`) {
		t.Errorf("the run's line does not name its stream and poll: %s", run)
	}
	nav := w.Nav{Target: "main", Items: []w.NavItem{{Label: "Tasks", Href: "/tasks"}}}.Render()
	if !strings.Contains(nav, `class="tc-nav" data-tc-target="main"`) {
		t.Errorf("the nav does not name its target: %s", nav)
	}
}

// TestFragmentForms (Carrel's task 226): an answer prints the zone's widget,
// then the page's forms its buttons name, and the same in each other zone.
func TestFragmentForms(t *testing.T) {
	got := page.Fragment{Body: []page.Part{w.Side{Zone: "panel"}}, Forms: []w.PostForm{{ID: "post-del", Action: "/del"}},
		Extra: []page.ZonePart{{Zone: "queue", Body: []page.Part{w.Msg{Parts: []w.MsgPart{w.Text("q")}}}, Forms: []w.PostForm{{ID: "post-ans", Action: "/ans"}}}}}.Render()
	side, del, tpl, ans := strings.Index(got, `class="tc-side"`), strings.Index(got, `id="post-del"`), strings.Index(got, `<template data-tc-zone="queue">`), strings.Index(got, `id="post-ans"`)
	if side < 0 || del < side || tpl < del || ans < tpl || ans > strings.LastIndex(got, `</template>`) {
		t.Errorf("the forms do not follow their zone's widget: %s", got)
	}
	// in a template of their own: an answer of forms only empties its zone
	// (review finding 12)
	if n := strings.Count(got, `<template data-tc-forms><form class="tc-post" id="post-`); n != 2 {
		t.Errorf("the forms stand in %d templates of their own, want 2: %s", n, got)
	}
	if got := (page.Fragment{}).Render(); got != "" {
		t.Errorf("an answer without forms prints %q", got)
	}
}

// TestFragmentTitle (Carrel's task 294): an answer's title comes first,
// escaped, before the zone's content; an answer without one prints none.
func TestFragmentTitle(t *testing.T) {
	got := page.Fragment{Title: "Дневник & <co>", Body: []page.Part{w.Msg{Parts: []w.MsgPart{w.Text("q")}}}}.Render()
	if !strings.HasPrefix(got, `<title>Дневник &amp; &lt;co&gt;</title><`) {
		t.Errorf("the title does not open the answer: %s", got)
	}
	if strings.Contains(page.Fragment{Body: []page.Part{w.Msg{Parts: []w.MsgPart{w.Text("q")}}}}.Render(), "<title>") {
		t.Error("an answer without a title prints one")
	}
}
