package layout_test

import (
	"strings"
	"testing"

	"carrel.dev/typecase/layout"
)

func TestCSS_threeTracks(t *testing.T) {
	spec := layout.Spec{
		Grids: []layout.Grid{{
			Selector: ".tc-test",
			Gap:      "8px",
			Tracks: []layout.Track{
				{Slot: "mark", Size: "3px"},
				{Slot: "title", Size: "minmax(140px,1.2fr)"},
				{Slot: "date", Size: "80px"},
			},
		}},
	}
	css := layout.Format(spec)
	if !strings.Contains(css, "minmax(140px,1.2fr)") {
		t.Fatalf("minmax comma spacing: %s", css)
	}
	if strings.Contains(css, "minmax(140px, 1.2fr)") {
		t.Fatalf("space after minmax comma: %s", css)
	}
	got := columnsOf(css, ".tc-test")
	want := []string{"3px", "minmax(140px,1.2fr)", "80px"}
	if len(got) != len(want) {
		t.Fatalf("tracks %v want %v in %s", got, want, css)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("track %d %q want %q", i, got[i], want[i])
		}
	}
}

func TestCSS_containerOmitsBase(t *testing.T) {
	css := layout.CSS()
	at := strings.Index(css, "@container")
	if at < 0 {
		t.Fatal("expected @container in excerpt CSS")
	}
	if strings.Contains(css[at:], "minmax(140px,1.2fr)") {
		t.Fatal("@container must not repeat base track sizes")
	}
}

func TestCSS_namedShellContainer(t *testing.T) {
	css := layout.CSS()
	if !strings.Contains(css, "@container tcbox (max-width:1190px)") {
		t.Fatal("shell hide-side rule must use named tcbox")
	}
	if strings.Contains(css, "@container (max-width:1190px)") {
		t.Fatal("unnamed @container is forbidden")
	}
	if !strings.Contains(css, ".tc-shell:not(.tc-shell--narrow) .tc-shell__side{display:none}") {
		t.Fatal("1190 rule must hide .tc-shell__side")
	}
}

func TestCSS_shellTracks(t *testing.T) {
	css := layout.CSS()
	got := columnsOf(css, ".tc-shell__body")
	want := []string{"264px", "minmax(0,1fr)"}
	if len(got) != len(want) {
		t.Fatalf("shell tracks %v want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("shell track %d %q want %q", i, got[i], want[i])
		}
	}
}

func TestCSS_stable(t *testing.T) {
	a := layout.CSS()
	b := layout.CSS()
	if a != b {
		t.Fatal("CSS() is not byte-stable")
	}
}

func TestCSS_formSpacer(t *testing.T) {
	css := layout.CSS()
	want := `*:has(> .tc-field > .tc-field__label) > .tc-field:has(> .tc-check:only-child)::before,
*:has(> .tc-field > .tc-field__label) > .tc-field:has(> .tc-btn:only-child)::before,
*:has(> .tc-field > .tc-field__label) > .tc-field:has(> .tc-seg:only-child)::before{content:"";height:16px;flex:none}
.tc-field.is-own-line::before{display:none}`
	if !strings.Contains(css, want) {
		t.Fatal("form spacer CSS missing or rewritten")
	}
}

func TestCSS_fieldWidths(t *testing.T) {
	css := layout.CSS()
	for _, frag := range []string{
		".tc-field--xs, .tc-pfield--xs{width:96px}",
		".tc-field--sm, .tc-pfield--sm{width:140px}",
		".tc-field--md, .tc-pfield--md{width:210px}",
		".tc-field--lg, .tc-pfield--lg{width:300px}",
		".tc-field--xl, .tc-pfield--xl{width:420px}",
	} {
		if !strings.Contains(css, frag) {
			t.Fatalf("missing field width %s in %s", frag, css)
		}
	}
	if !strings.Contains(css, ".tc-field--full, .tc-pfield--full{") {
		t.Fatal("missing full field width")
	}
	if !strings.Contains(css, "max-width:760px") {
		t.Fatal("full width must cap at 760px")
	}
}

func TestCSS_restBoxes(t *testing.T) {
	css := layout.CSS()
	for _, frag := range []string{
		".tc-grid{display:grid;gap:14px;grid-template-columns:repeat(auto-fill,minmax(132px,1fr))}",
		".tc-doc__layout{display:grid;grid-template-columns:minmax(0,1fr) 400px;min-height:0}",
		"@container tcbox (max-width:880px){\n.tc-doc__layout{grid-template-columns:minmax(0,1fr) 296px}\n.tc-docmeta{width:auto}\n}",
		".tc-skel__line{",
		"height:9px",
		".tc-meter{",
		"width:64px",
		"height:3px",
		".tc-chip, .tc-tag{",
		"font-size:10.5px",
		"min-height:16px",
		".tc-queue{",
		"width:330px",
		"position:absolute",
		".tc-prose__h1{font-size:27px;line-height:1.15}",
		".tc-fields__row{display:grid;grid-template-columns:104px minmax(0,1fr);gap:7px 14px}",
		".tc-tick{height:13px;width:13px}",
		".tc-btn{border-color:var(--tc-edge)}",
		".tc-table__td{border-color:var(--tc-hair)}",
		".tc-top{border-color:var(--tc-rule)}",
		".tc-tree__row{",
		"gap:7px",
		"height:26px",
		".tc-empty{",
		"text-align:start",
		".tc-list > .tc-rubric{",
		"padding:15px var(--tc-page-gutter) 5px",
	} {
		if !strings.Contains(css, frag) {
			t.Fatalf("missing %q in generated CSS", frag)
		}
	}
	if strings.Contains(css, "--tc-row-height") {
		t.Fatal("generated CSS must not set --tc-row-height on rest boxes")
	}
}

func TestCSS_cols(t *testing.T) {
	css := layout.CSS()
	halves := columnsOf(css, ".tc-cols--halves")
	if len(halves) != 2 || halves[0] != "minmax(0,1fr)" || halves[1] != "minmax(0,1fr)" {
		t.Fatalf("halves tracks %v", halves)
	}
	thirds := columnsOf(css, ".tc-cols--thirds")
	if len(thirds) != 3 {
		t.Fatalf("thirds tracks %v", thirds)
	}
	side := columnsOf(css, ".tc-cols--sidebar")
	if len(side) != 2 || side[1] != "296px" {
		t.Fatalf("sidebar tracks %v", side)
	}
	doc := columnsOf(css, ".tc-doc__layout")
	if len(doc) < 2 || doc[1] != "400px" {
		t.Fatalf("doc meta column is the details panel, 400px: %v", doc)
	}
	if !strings.Contains(css, ".tc-cols--halves{display:grid;grid-template-columns:minmax(0,1fr) minmax(0,1fr);gap:14px}") {
		t.Fatal("halves base rule missing")
	}
	if !strings.Contains(css, "@container tcbox (max-width:620px){\n.tc-cols--halves{grid-template-columns:minmax(0,1fr)}\n}") {
		t.Fatal("halves 620 rule missing")
	}
	if !strings.Contains(css, ".tc-cols__col{container:tcbox / inline-size;min-width:0}") {
		t.Fatal("col box missing")
	}
}

func TestCSS_sheetAndDialog(t *testing.T) {
	css := layout.CSS()
	if !strings.Contains(css, ".tc-sheet{") || !strings.Contains(css, "width:560px") {
		t.Fatal("sheet width 560 missing")
	}
	if !strings.Contains(css, ".tc-dialog{") || !strings.Contains(css, "max-width:420px") {
		t.Fatal("dialog max-width 420 missing")
	}
	if !strings.Contains(css, ".tc-sheet__target::after") {
		t.Fatal("sheet target pick chevron missing")
	}
}

func TestSplitTracks_outsideParens(t *testing.T) {
	got := layout.SplitTracks("3px minmax(140px,1.2fr) 80px")
	if len(got) != 3 || got[1] != "minmax(140px,1.2fr)" {
		t.Fatalf("got %v", got)
	}
}

func TestTracksAt(t *testing.T) {
	got := layout.TracksAt(layout.RowTimeCols, 616)
	if len(got) < 2 {
		t.Fatalf("tracks %v", got)
	}
	if got[1].Size != "54px" {
		t.Fatalf("lead at 616 %q want 54px (620 block)", got[1].Size)
	}
	wide := layout.TracksAt(layout.RowTimeCols, 900)
	if wide[1].Size != "62px" {
		t.Fatalf("lead at 900 %q want 62px", wide[1].Size)
	}
}

func TestFormat_unnamedContainerPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic")
		}
	}()
	layout.Format(layout.Spec{Grids: []layout.Grid{{
		Selector: ".tc-x",
		Tracks:   []layout.Track{{Size: "1fr"}},
		Narrow:   []layout.Breakpoint{{MaxPx: 700, Tracks: []layout.Track{{Size: "1fr"}}}},
	}}})
}

func columnsOf(css, sel string) []string {
	needle := sel + "{"
	i := strings.Index(css, needle)
	if i < 0 {
		return nil
	}
	body := css[i+len(needle):]
	end := strings.Index(body, "}")
	if end < 0 {
		return nil
	}
	for _, part := range strings.Split(body[:end], ";") {
		k, v, ok := strings.Cut(part, ":")
		if ok && strings.TrimSpace(k) == "grid-template-columns" {
			return layout.SplitTracks(v)
		}
	}
	return nil
}
