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
	run := w.Runbar{Zone: "run", Stream: "/run/events", Changed: "changed", Ended: "ended", Poll: "/run", Every: 5}.Render()
	if !strings.Contains(run, `data-tc-zone="run" data-tc-stream="/run/events" data-tc-changed="changed" data-tc-ended="ended" data-tc-poll="/run" data-tc-every="5"`) {
		t.Errorf("the run's line does not name its stream and poll: %s", run)
	}
	nav := w.Nav{Target: "main", Items: []w.NavItem{{Label: "Tasks", Href: "/tasks"}}}.Render()
	if !strings.Contains(nav, `class="tc-nav" data-tc-target="main"`) {
		t.Errorf("the nav does not name its target: %s", nav)
	}
}
