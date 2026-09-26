package catalog

import (
	"fmt"
	"html"
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

// PageHTML renders a full catalog page: the widgets' cells family by family,
// and after the last cell of a widget the mockup's component and fragment
// frames built around it (Р-16).
func PageHTML(title string, entries []Entry, withJS bool, frames ...Frame) string {
	byBlock := map[string][]Frame{}
	for _, f := range frames {
		byBlock[f.Block] = append(byBlock[f.Block], f)
	}
	shown := map[string]bool{}
	var sections strings.Builder
	family := ""
	for i, e := range entries {
		if e.Family != family {
			if family != "" {
				sections.WriteString("</section>")
			}
			family = e.Family
			sections.WriteString(`<section class="catalog-family"><h2>` + family + `</h2>`)
		}
		sections.WriteString(`<article class="catalog-entry" id="` + slug(e.Widget) + `-` + slug(e.State) + `"><h3><code>` + e.Widget + `</code> — ` + e.State + `</h3><div class="catalog-demo tc-root">` + anchored(e.HTML) + `</div></article>`)
		last := i+1 == len(entries) || entries[i+1].Widget != e.Widget || entries[i+1].Family != e.Family
		if last && !shown[e.Widget] {
			shown[e.Widget] = true
			for _, f := range byBlock[e.Widget] {
				sections.WriteString(`<article class="catalog-entry catalog-frame" id="frame-` + f.ID + `"><h3>` + html.EscapeString(f.Caption) + ` <code>` + f.ID + `</code> — a frame of the mockup</h3><div class="catalog-demo">` + f.HTML + `</div></article>`)
			}
		}
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
%s
%s
</body>
</html>`, title, defs, title, ladderHTML(), sections.String(), roomJS, js)
}

// anchored gives a menu shown without its trigger (an open menu's record)
// the anchor its trigger would be: the menu hangs below it, as in its frame.
func anchored(h string) string {
	if i := strings.Index(h, ">"); i > 0 && strings.Contains(h[:i], `class="tc-menu `) && strings.Contains(h[:i], " is-pop") {
		return `<div class="catalog-anchor">` + h + `</div>`
	}
	return h
}

func slug(s string) string {
	return strings.NewReplacer(" ", "-", "/", "-", ".", "-").Replace(s)
}
