package main

import (
	"strings"
	"testing"
)

func TestReplaceSelBoundary(t *testing.T) {
	css := `.m-row{x}.m-row--note{y}.m-row .t{z}`
	got := replaceSel(css, ".m-row", ".tc-row")
	want := `.tc-row{x}.m-row--note{y}.tc-row .t{z}`
	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestApplyRenamesClassAndCSS(t *testing.T) {
	tsv := ".m-btn\t.tc-btn\n.mock\t.tc-root\n.m-seg\t.tc-seg\n.m-seg span\t.tc-seg__opt\n"
	rules, err := parseTSV(tsv)
	if err != nil {
		t.Fatal(err)
	}
	in := `<style data-tc>.mock .m-btn{height:26px}.m-seg span{padding:0}</style>` +
		`<style>body{margin:0}</style>` +
		`<div class="mock"><span class="m-btn">Go</span><span class="m-seg"><span class="is-on">A</span><span>B</span></span></div>`
	got, err := apply(in, rules)
	if err != nil {
		t.Fatal(err)
	}
	if got == in {
		t.Fatal("expected changes")
	}
	if !containsAll(got, `.tc-root .tc-btn{height:26px}`, `.tc-seg__opt{padding:0}`, `class="tc-root"`, `class="tc-btn"`, `tc-seg__opt`) {
		t.Fatalf("got %s", got)
	}
	again, err := apply(got, rules)
	if err != nil {
		t.Fatal(err)
	}
	if again != got {
		t.Fatalf("second apply changed file:\n%s\nvs\n%s", again, got)
	}
}

func TestCrumbsLinkKeepsParentInCSS(t *testing.T) {
	tsv := ".m-crumbs\t.tc-crumbs\n.m-crumbs a\t.tc-crumbs .tc-crumbs__link\n"
	rules, err := parseTSV(tsv)
	if err != nil {
		t.Fatal(err)
	}
	in := `<style data-tc>.m-crumbs a{color:red}</style><style>body{}</style>` +
		`<nav class="m-crumbs"><a>Docs</a></nav>`
	got, err := apply(in, rules)
	if err != nil {
		t.Fatal(err)
	}
	if !containsAll(got, `.tc-crumbs .tc-crumbs__link{color:red}`, `class="tc-crumbs__link"`) {
		t.Fatalf("got %s", got)
	}
	if containsAll(got, `class="tc-crumbs tc-crumbs__link"`) {
		t.Fatalf("anchor must not also carry tc-crumbs: %s", got)
	}
	again, err := apply(got, rules)
	if err != nil {
		t.Fatal(err)
	}
	if again != got {
		t.Fatalf("second apply changed file:\n%s\nvs\n%s", again, got)
	}
}

func TestNavItemKeepsParentInCSS(t *testing.T) {
	tsv := ".m-nav\t.tc-nav\n.m-nav a\t.tc-nav .tc-nav__item\n.m-tabbar\t.tc-nav--bottom\n.m-tabbar a\t.tc-nav__item\n"
	rules, err := parseTSV(tsv)
	if err != nil {
		t.Fatal(err)
	}
	in := `<style data-tc>.m-nav a{height:29px}.m-nav a.is-on{font-weight:600}.m-tabbar a{font-size:9.5px}</style><style>body{}</style>` +
		`<nav class="m-nav"><a>Notes</a></nav><div class="m-tabbar"><a>Notes</a></div>`
	got, err := apply(in, rules)
	if err != nil {
		t.Fatal(err)
	}
	if !containsAll(got, `.tc-nav .tc-nav__item{height:29px}`, `.tc-nav .tc-nav__item.is-on{font-weight:600}`, `class="tc-nav--bottom"`, `class="tc-nav__item"`) {
		t.Fatalf("got %s", got)
	}
	stripped := strings.ReplaceAll(got, `.tc-nav .tc-nav__item{height:29px}`, "")
	if strings.Contains(stripped, `.tc-nav__item{height:29px}`) {
		t.Fatalf("rail height must not leak onto bare tc-nav__item: %s", got)
	}
}

func TestNoteMonoDateIsNotExcerpt(t *testing.T) {
	// Child 3 of a note row is .s in the notes list and .mono in the mail
	// attach sheet. Mapping nth-child(3) → __excerpt paints the date 12px.
	tsv := ".m-row\t.tc-row\n.m-row--note\t.tc-row--excerpt\n.m-row .s\t.tc-row .tc-row__excerpt\n.mono\t.tc-mono\n"
	rules, err := parseTSV(tsv)
	if err != nil {
		t.Fatal(err)
	}
	in := `<div class="m-row m-row--note"><span class="m-bar3"></span><span class="t">Title</span><span class="mono">18 Aug</span></div>`
	got, err := apply(in, rules)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(got, "tc-row__excerpt") {
		t.Fatalf("mono date must not get excerpt: %s", got)
	}
	if !strings.Contains(got, "tc-mono") {
		t.Fatalf("got %s", got)
	}
}

func TestNthChildAddsSlot(t *testing.T) {
	tsv := ".m-row\t.tc-row\n.m-row--contact\t.tc-row--avatarcols\n.m-row .s\t.tc-row__excerpt\n.m-row--contact > :nth-child(6)\t.tc-row--avatarcols .tc-row__caption\n"
	rules, err := parseTSV(tsv)
	if err != nil {
		t.Fatal(err)
	}
	in := `<div class="m-row m-row--contact"><i></i><i></i><i></i><i></i><i></i><span class="s">Org</span><span class="d">src</span></div>`
	got, err := apply(in, rules)
	if err != nil {
		t.Fatal(err)
	}
	if !containsAll(got, `tc-row--avatarcols`, `tc-row__excerpt`, `tc-row__caption`) {
		t.Fatalf("got %s", got)
	}
}

func TestMoveMkRules(t *testing.T) {
	tsv := ".m-win\t.mk-win\n.m-btn\t.tc-btn\n"
	rules, err := parseTSV(tsv)
	if err != nil {
		t.Fatal(err)
	}
	in := `<style data-tc>.m-win{border:1px}.m-btn{height:26px}</style><style>body{margin:0}</style>`
	got, err := apply(in, rules)
	if err != nil {
		t.Fatal(err)
	}
	_, pageOpen, pageClose, _ := findStyle(got, false)
	_, prodOpen, prodClose, _ := findStyle(got, true)
	page, prod := got[pageOpen:pageClose], got[prodOpen:prodClose]
	if !containsAll(page, `.mk-win{border:1px}`) || containsAll(prod, `.mk-win`) {
		t.Fatalf("mk rule not moved: %s", got)
	}
	if !containsAll(prod, `.tc-btn{height:26px}`) {
		t.Fatalf("product rule lost: %s", got)
	}
}

func containsAll(s string, subs ...string) bool {
	for _, sub := range subs {
		if !stringContains(s, sub) {
			return false
		}
	}
	return true
}

func stringContains(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || len(sub) == 0 || indexOfStr(s, sub) >= 0)
}

func indexOfStr(s, sub string) int {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}

func indexOfStyle(html string, dataTC bool) int {
	_, open, _, ok := findStyle(html, dataTC)
	if !ok {
		return 0
	}
	return open
}
