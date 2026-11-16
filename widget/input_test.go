package widget

import (
	"strings"
	"testing"
)

// TestInputOf: a field outside its form names the form (Carrel's task 195);
// without Of it belongs to the form it stands in.
func TestInputOf(t *testing.T) {
	if got := (Input{Name: "url", Of: "install"}).Render(); !strings.Contains(got, ` form="install" name="url"`) {
		t.Errorf("the field does not name its form: %s", got)
	}
	if got := (Input{Name: "url"}).Render(); strings.Contains(got, "form=") {
		t.Errorf("a field without Of names a form: %s", got)
	}
}

// TestCheckOf: a check outside its form names the form (Carrel's task 195).
func TestCheckOf(t *testing.T) {
	if got := (Check{Name: "agree", Value: "on", Of: "dav", Parts: []CheckPart{Text("Agree")}}).Render(); !strings.Contains(got, ` form="dav" name="agree"`) {
		t.Errorf("the check does not name its form: %s", got)
	}
}

func TestCheckReason(t *testing.T) {
	got := (Check{Name: "people", Value: "on", Reason: "Mail is not configured.", Parts: []CheckPart{Text("Let users send")}}).Render()
	if !strings.Contains(got, `<label class="tc-check" title="Mail is not configured.">`) || !strings.Contains(got, ` disabled>`) {
		t.Errorf("the check is not off with its reason: %s", got)
	}
}

func TestSheetIDWithoutAction(t *testing.T) {
	got := (Sheet{Title: "Attach from Files", ID: "picker-sheet"}).Render()
	if !strings.HasPrefix(got, `<div class="tc-sheet" id="picker-sheet"`) {
		t.Errorf("a sheet that is no form does not carry its id: %s", got)
	}
}

func TestSegOptionForm(t *testing.T) {
	got := (Seg{Options: []SegOption{{Label: "read only", Form: "acts", To: "/a?op=read"}, {Label: "read and write", On: true, Form: "acts", To: "/a?op=write"}}}).Render()
	for _, want := range []string{`<button class="tc-seg__opt" type="submit" form="acts" formaction="/a?op=read" aria-pressed="false">read only</button>`,
		`<button class="is-on tc-seg__opt" type="submit" form="acts" formaction="/a?op=write" aria-pressed="true">read and write</button>`} {
		if !strings.Contains(got, want) {
			t.Errorf("an option that sends a form: want %s in %s", want, got)
		}
	}
}

// TestSegRole: buttons that send a form and links are no radios, so their
// seg is a group; plain options and radios of a form keep the radiogroup
// (review finding 11).
func TestSegRole(t *testing.T) {
	for _, c := range []struct {
		seg  Seg
		role string
	}{
		{Seg{Options: []SegOption{{Label: "read only", Form: "acts", To: "/a?op=read"}, {Label: "write", On: true, Form: "acts", To: "/a?op=write"}}}, "group"},
		{Seg{Options: []SegOption{{Label: "A–Z", On: true, Href: "/c?sort=az"}, {Label: "Recent", Href: "/c?sort=new"}}}, "group"},
		{Seg{Options: []SegOption{{Label: "Day", On: true}, {Label: "Week"}}}, "radiogroup"},
		{Seg{Name: "theme", Options: []SegOption{{Label: "Auto", Value: "auto", On: true}, {Label: "Light", Value: "light"}}}, "radiogroup"},
	} {
		if got := c.seg.Render(); !strings.Contains(got, `role="`+c.role+`"`) {
			t.Errorf("want role %q: %s", c.role, got)
		}
	}
}

// A query line with a short hint prints the short one and the whole one over
// it, which a narrow field wraps out of sight; with a value typed, or a
// short that is not the hint's start, there is nothing of the kind.
func TestQueryShort(t *testing.T) {
	got := Query{Placeholder: "Filter by word", Short: "Filter", Label: "Filter by word"}.Render()
	want := `<input class="tc-q" type="search" placeholder="Filter by word" aria-label="Filter by word"><span class="tc-q__ph" aria-hidden="true"><span class="tc-q__hint"><span class="tc-q__short">Filter</span><span class="tc-q__full">Filter by word</span></span></span>`
	if got != want {
		t.Errorf("short hint:\n got %s\nwant %s", got, want)
	}
	for _, q := range []Query{
		{Value: "синк", Placeholder: "Filter by word", Short: "Filter", Label: "Filter by word"},
		{Placeholder: "Filter by word", Short: "Find", Label: "Filter by word"},
		{Placeholder: "Filter by word", Label: "Filter by word"},
	} {
		if s := q.Render(); strings.Contains(s, "tc-q__ph") {
			t.Errorf("%+v: no short hint expected, got %s", q, s)
		}
	}
}
