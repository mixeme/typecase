// Package page prints the document around a screen: the head a page of the
// library needs and the body that holds one of the library's roots (task 38
// of Carrel's server). The application writes no markup of its own: the
// title, the language, where it serves the library's files and the root.
//
// The head follows a strict content policy (script-src 'self'; style-src
// 'self'): no inline script and no inline style — the theme's script and the
// stylesheet are loaded by address, the theme's script first so the theme is
// set before the first paint (L-5.3).
package page

import (
	"html"
	"net/http"
	"strings"

	"carrel.dev/typecase"
	"carrel.dev/typecase/widget"
)

// Assets are the addresses the application serves the library's files at.
type Assets struct {
	CSS   string // the stylesheet; the fonts lie at ../fonts/ from it
	Theme string // the theme's script, run before the first paint
	JS    string // the library's script, run after the body
	Live  string // the library's live layer, run after JS when the document asks for it
}

// At is the addresses of Static mounted at base: base/css/typecase.css,
// base/fonts/…, base/js/theme.js, base/js/tc.js, base/js/tc-live.js.
func At(base string) Assets {
	base = strings.TrimRight(base, "/")
	return Assets{CSS: base + "/css/typecase.css", Theme: base + "/js/theme.js", JS: base + "/js/tc.js", Live: base + "/js/tc-live.js"}
}

// Document is a whole page: its title and language, the library's files and
// the root it shows, with the forms its buttons submit by name.
type Document struct {
	Title  string
	Lang   string // "en" when empty
	Assets Assets
	Root   widget.Root
	Forms  []widget.PostForm // forms buttons outside a form submit (Button.Form)
	// Live loads the library's live layer after its script (L-1): a tag
	// field's suggestions; without it they are plain forms.
	Live bool
	// Zones is what the live layer says when it asks the server for a part
	// of the page (Carrel's task 84); the names are the application's.
	Zones Zones
}

// Zones names the headers of a request for a part of a page and the zone
// whose content owns the address. The live layer sends ViewHeader with
// ViewValue to say a part is wanted and TargetHeader with the name of the
// zone asked for; a form's token — its field TokenField — goes into
// TokenHeader as well. Address is the zone whose answer to a link puts the
// link's address into the history.
type Zones struct {
	ViewHeader   string
	ViewValue    string
	TargetHeader string
	TokenField   string
	TokenHeader  string
	Address      string
}

func (z Zones) attrs() string {
	var b strings.Builder
	for _, a := range [][2]string{{"data-tc-view-header", z.ViewHeader}, {"data-tc-view-value", z.ViewValue}, {"data-tc-target-header", z.TargetHeader},
		{"data-tc-token-field", z.TokenField}, {"data-tc-token-header", z.TokenHeader}, {"data-tc-address-zone", z.Address}} {
		if a[1] != "" {
			b.WriteString(` ` + a[0] + `="` + html.EscapeString(a[1]) + `"`)
		}
	}
	return b.String()
}

// Part is what an answer puts into a zone: a widget of the library.
type Part interface{ Render() string }

// Fragment is the answer to a request for a part of a page: Body is what
// goes into the zone that was asked for, Extra the other zones the answer
// changes with it — the messages, the count in the rail. A widget that
// carries its zone's name (Zone) takes the place of the zone whole; anything
// else becomes the zone's content. An empty Body empties the zone: a panel
// or a sheet closes.
type Fragment struct {
	Body  []Part
	Extra []ZonePart
}

// ZonePart is one more zone of an answer and what goes into it.
type ZonePart struct {
	Zone string
	Body []Part
}

// Render prints the answer: the body, then each other zone in a template
// that names it.
func (f Fragment) Render() string {
	var b strings.Builder
	for _, p := range f.Body {
		b.WriteString(p.Render())
	}
	for _, z := range f.Extra {
		b.WriteString(`<template data-tc-zone="` + html.EscapeString(z.Zone) + `">`)
		for _, p := range z.Body {
			b.WriteString(p.Render())
		}
		b.WriteString(`</template>`)
	}
	return b.String()
}

// Render prints the document.
func (d Document) Render() (string, error) {
	sprite, err := typecase.Sprite()
	if err != nil {
		return "", err
	}
	lang := d.Lang
	if lang == "" {
		lang = "en"
	}
	var b strings.Builder
	b.WriteString(`<!doctype html><html lang="` + html.EscapeString(lang) + `" data-tc-root data-tc-noscript` + d.Zones.attrs() + `><head><meta charset="utf-8">` +
		`<meta name="viewport" content="width=device-width, initial-scale=1"><title>` + html.EscapeString(d.Title) + `</title>` +
		`<script src="` + html.EscapeString(d.Assets.Theme) + `"></script>` +
		`<link rel="stylesheet" href="` + html.EscapeString(d.Assets.CSS) + `"></head><body data-tc-root>`)
	b.WriteString(`<div hidden>` + string(sprite) + `</div>`)
	if d.Root != nil {
		b.WriteString(d.Root.Render())
	}
	for _, f := range d.Forms {
		b.WriteString(f.Render())
	}
	b.WriteString(`<script src="` + html.EscapeString(d.Assets.JS) + `"></script>`)
	if d.Live && d.Assets.Live != "" {
		b.WriteString(`<script src="` + html.EscapeString(d.Assets.Live) + `"></script>`)
	}
	b.WriteString(`</body></html>`)
	return b.String(), nil
}

// Static serves the library's files from the binary in the layout At names:
// css/typecase.css, fonts/*, js/theme.js, js/tc.js, js/tc-live.js. Mount it under a prefix
// with http.StripPrefix.
func Static() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/css/typecase.css", func(w http.ResponseWriter, r *http.Request) {
		css, err := typecase.CSS()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		_, _ = w.Write([]byte(css))
	})
	script := func(read func() ([]byte, error)) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			b, err := read()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
			_, _ = w.Write(b)
		}
	}
	mux.HandleFunc("/js/theme.js", script(typecase.ThemeJS))
	mux.HandleFunc("/js/tc.js", script(typecase.JS))
	mux.HandleFunc("/js/tc-live.js", script(typecase.LiveJS))
	mux.Handle("/fonts/", http.FileServer(http.FS(typecase.Fonts())))
	return mux
}
