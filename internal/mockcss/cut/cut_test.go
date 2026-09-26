package cut_test

import (
	"strings"
	"testing"

	"carrel.dev/typecase/internal/mockcss/cut"
)

func TestCut(t *testing.T) {
	html := "<style>.page{}</style>\n<style data-tc>:root{--tc-ink:#000}\n.tc-x{}\n</style><div></div>"
	got, err := cut.Cut(html)
	if err != nil {
		t.Fatal(err)
	}
	if got != ":root{--tc-ink:#000}\n.tc-x{}\n" {
		t.Fatalf("cut changed the text: %q", got)
	}
}

func TestCutRefuses(t *testing.T) {
	for name, html := range map[string]string{
		"none":     "<style>.page{}</style>",
		"two":      "<style data-tc>a{}</style><style data-tc>b{}</style>",
		"unclosed": "<style data-tc>a{}",
	} {
		if _, err := cut.Cut(html); err == nil {
			t.Errorf("%s: no error", name)
		}
	}
}

func TestCutLiveMockup(t *testing.T) {
	css, err := cut.FromRoot("../../..")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(css, ":root{") || !strings.Contains(css, ".tc-root{") {
		t.Fatalf("live cut does not look like the product CSS: %.80q", css)
	}
}
