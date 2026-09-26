package catalog

import (
	"fmt"
	"strings"

	"carrel.dev/typecase"
	"carrel.dev/typecase/internal/refs"
)

// Entry is one widget demonstration in the catalog (L-9.1).
type Entry struct {
	Family string
	Widget string
	State  string
	HTML   string
}

// Entries returns one catalog cell per reference record: the widget printed
// from the data of its frame, in the state that record shows (L-9.1). The
// catalog is built from the table the fragments are checked against, so a
// widget in the catalog is exactly a form the mockup draws.
func Entries() []Entry {
	var out []Entry
	for _, e := range refs.Entries {
		out = append(out, Entry{
			Family: e.Family,
			Widget: e.Ref.Class,
			State:  e.State,
			HTML:   e.Fixture.Render(),
		})
	}
	return out
}

// PageHTML renders a full catalog page.
func PageHTML(title string, entries []Entry, withJS bool) string {
	var sections strings.Builder
	family := ""
	for _, e := range entries {
		if e.Family != family {
			if family != "" {
				sections.WriteString("</section>")
			}
			family = e.Family
			sections.WriteString(`<section class="catalog-family"><h2>` + family + `</h2>`)
		}
		sections.WriteString(`<article class="catalog-entry" id="` + slug(e.Widget) + `-` + slug(e.State) + `"><h3><code>` + e.Widget + `</code> — ` + e.State + `</h3><div class="catalog-demo tc-root">` + e.HTML + `</div></article>`)
	}
	if family != "" {
		sections.WriteString("</section>")
	}
	// The theme boot script runs on both pages: it must precede first paint
	// (L-5.3) and is not widget behavior. Only tc.js is withheld from the
	// second page.
	js := ""
	if withJS {
		js = `<script src="/tc.js"></script>`
	}
	sprite, err := typecase.Sprite()
	if err != nil {
		sprite = nil
	}
	defs := string(sprite)
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en" data-tc-root>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>%s</title>
<script src="/theme.js"></script>
<link rel="stylesheet" href="/typecase.css">
<link rel="stylesheet" href="/catalog.css">
</head>
<body data-tc-root>
<div hidden aria-hidden="true">%s</div>
<header class="catalog-nav"><a href="/">With behavior</a><a href="/static">Without behavior</a></header>
<h1>%s</h1>
%s
%s
</body>
</html>`, title, defs, title, sections.String(), js)
}

func slug(s string) string {
	return strings.NewReplacer(" ", "-", "/", "-", ".", "-").Replace(s)
}
