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
	// Worker is the library's worker for a page installed as an application
	// (Document.App); Static serves it allowed to serve the whole site.
	Worker string
}

// At is the addresses of Static mounted at base: base/css/typecase.css,
// base/fonts/…, base/js/theme.js, base/js/tc.js, base/js/tc-live.js,
// base/js/tc-sw.js.
func At(base string) Assets {
	base = strings.TrimRight(base, "/")
	return Assets{CSS: base + "/css/typecase.css", Theme: base + "/js/theme.js", JS: base + "/js/tc.js", Live: base + "/js/tc-live.js",
		Worker: base + "/js/tc-sw.js"}
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
	// App makes the page one that can be installed as an application
	// (Carrel's task 91); nil, it is a page.
	App *App
	// Theme is the theme the server chose for this browser (Carrel's task
	// 186: the choice is kept in a cookie): "light" or "dark" is printed on
	// the page, so a page without scripts shows it too; "auto" leaves it to
	// the system. Either way theme.js and the theme's switch keep it and do
	// not read or write what this browser remembers. Empty, the library
	// decides as before: the address's ?theme=, what this browser
	// remembers, the system.
	Theme string
	// Suggest is the address of the search line's suggestions (Carrel's
	// task 227): with the live layer the line asks it ?q=… as it gets the
	// focus and as something is typed, and the answer — a Menu whose items
	// say When — stands under the line. The suggestions come from what the
	// server holds, so the page prints none; without the live layer the
	// line has no menu and Enter sends the search.
	Suggest string
}

func (d Document) suggestAttr() string {
	if d.Suggest == "" {
		return ""
	}
	return ` data-tc-search-suggest="` + html.EscapeString(d.Suggest) + `"`
}

func (d Document) themeAttrs() string {
	switch d.Theme {
	case "light", "dark":
		return ` data-tc-theme="` + d.Theme + `" data-tc-theme-set="` + d.Theme + `"`
	case "auto":
		return ` data-tc-theme-set="auto"`
	}
	return ""
}

// App is what a page installed as an application names: the manifest and
// the icons are the application's own files, served by it. With Offline the
// library's worker (Assets.Worker) is registered by tc.js: it keeps the
// library's files and that page, shown when the network is gone, and
// nothing else — the application's pages and data always come from the
// network.
type App struct {
	Manifest   string // the address of the web app manifest
	ThemeColor string // the colour of the browser's bar around the page: "#2f4a36"
	Icon       string // the address of the page's icon
	Offline    string // the address of the page shown offline; empty, no worker
	Scope      string // the addresses the worker serves; "/" when empty
}

func (a *App) head() string {
	if a == nil {
		return ""
	}
	s := ""
	if a.ThemeColor != "" {
		s += `<meta name="theme-color" content="` + html.EscapeString(a.ThemeColor) + `">`
	}
	if a.Manifest != "" {
		s += `<link rel="manifest" href="` + html.EscapeString(a.Manifest) + `">`
	}
	if a.Icon != "" {
		s += `<link rel="icon" href="` + html.EscapeString(a.Icon) + `">`
	}
	return s
}

func (a *App) attrs(worker string) string {
	if a == nil || a.Offline == "" || worker == "" {
		return ""
	}
	scope := a.Scope
	if scope == "" {
		scope = "/"
	}
	return ` data-tc-worker="` + html.EscapeString(worker) + `" data-tc-offline="` + html.EscapeString(a.Offline) + `" data-tc-scope="` + html.EscapeString(scope) + `"`
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

// Printed is what the library's widgets printed, as a part of an answer:
// widget.PrintMain's result — parts of a page that are no widgets by
// themselves, a message's technical block among them. It is not a door for
// markup of the application's own.
type Printed string

// Render gives the printed parts back.
func (p Printed) Render() string { return string(p) }

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
	b.WriteString(`<!doctype html><html lang="` + html.EscapeString(lang) + `" data-tc-root data-tc-noscript` + d.themeAttrs() + d.Zones.attrs() + d.suggestAttr() + d.App.attrs(d.Assets.Worker) + `><head><meta charset="utf-8">` +
		`<meta name="viewport" content="width=device-width, initial-scale=1"><title>` + html.EscapeString(d.Title) + `</title>` + d.App.head() +
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
	worker := script(typecase.WorkerJS)
	mux.HandleFunc("/js/tc-sw.js", func(w http.ResponseWriter, r *http.Request) {
		// the worker lies with the library's files and serves the whole site
		w.Header().Set("Service-Worker-Allowed", "/")
		worker(w, r)
	})
	mux.Handle("/fonts/", http.FileServer(http.FS(typecase.Fonts())))
	return mux
}
