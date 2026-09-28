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
}

// At is the addresses of Static mounted at base: base/css/typecase.css,
// base/fonts/…, base/js/theme.js, base/js/tc.js.
func At(base string) Assets {
	base = strings.TrimRight(base, "/")
	return Assets{CSS: base + "/css/typecase.css", Theme: base + "/js/theme.js", JS: base + "/js/tc.js"}
}

// Document is a whole page: its title and language, the library's files and
// the root it shows, with the forms its buttons submit by name.
type Document struct {
	Title  string
	Lang   string // "en" when empty
	Assets Assets
	Root   widget.Root
	Forms  []widget.PostForm // forms buttons outside a form submit (Button.Form)
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
	b.WriteString(`<!doctype html><html lang="` + html.EscapeString(lang) + `" data-tc-root><head><meta charset="utf-8">` +
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
	b.WriteString(`<script src="` + html.EscapeString(d.Assets.JS) + `"></script></body></html>`)
	return b.String(), nil
}

// Static serves the library's files from the binary in the layout At names:
// css/typecase.css, fonts/*, js/theme.js, js/tc.js. Mount it under a prefix
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
	mux.Handle("/fonts/", http.FileServer(http.FS(typecase.Fonts())))
	return mux
}
