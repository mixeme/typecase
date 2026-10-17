package widget

import "testing"

// TestSlug: the rule of Carrel's architecture.md, as its server had it.
func TestSlug(t *testing.T) {
	for name, want := range map[string]string{
		"Поездки":               "poezdki",
		"Мой  Календарь — 2026": "moi-kalendar-2026",
		"Ёлка и йод":            "elka-i-iod",
		"Café Élan":             "cafe-elan",
		"Щука-Ъ":                "shchuka",
		"  Дела  ":              "dela",
		"---":                   SlugEmpty,
		"":                      SlugEmpty,
		"日本":                    SlugEmpty,
	} {
		if got := Slug(name); got != want {
			t.Errorf("Slug(%q) = %q; want %q", name, got, want)
		}
	}
}
