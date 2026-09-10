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

func TestSplitTracks_outsideParens(t *testing.T) {
	got := layout.SplitTracks("3px minmax(140px,1.2fr) 80px")
	if len(got) != 3 || got[1] != "minmax(140px,1.2fr)" {
		t.Fatalf("got %v", got)
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
