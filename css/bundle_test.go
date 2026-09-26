package css_test

import (
	"io/fs"
	"regexp"
	"sort"
	"strings"
	"testing"

	"carrel.dev/typecase/css"
	"carrel.dev/typecase/token"
)

func TestBundle(t *testing.T) {
	out, err := css.Bundle()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "@font-face") || !strings.Contains(out, ".tc-root{") {
		t.Fatal("bundle missing fonts or product rules")
	}
	a, err := css.Bundle()
	if err != nil {
		t.Fatal(err)
	}
	if a != out {
		t.Fatal("bundle is not byte-stable")
	}
}

var darkBlock = regexp.MustCompile(`\[data-tc-theme="dark"\]\{([^}]*)\}`)

// The dark theme changes palette values and nothing else (L-5.3).
func TestDarkThemeOnlyChangesPalette(t *testing.T) {
	data, err := fs.ReadFile(css.Files(), "typecase.css")
	if err != nil {
		t.Fatal(err)
	}
	m := darkBlock.FindStringSubmatch(string(data))
	if m == nil {
		t.Fatal(`[data-tc-theme="dark"]{…} missing`)
	}
	got := make([]string, 0)
	for name := range token.Defined(m[1]) {
		got = append(got, name)
	}
	sort.Strings(got)
	want := append([]string(nil), token.Palette...)
	sort.Strings(want)
	if strings.Join(got, " ") != strings.Join(want, " ") {
		t.Fatalf("dark theme defines %v, want the palette %v", got, want)
	}
}
