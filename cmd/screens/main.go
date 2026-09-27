// Command screens serves the demo screens: every frame of the mockup that is
// a route (decision Р-16, docs/development.md) built from widgets, at /f/<data-frame>.
//
//	go run ./cmd/screens [-addr :8085]
package main

import (
	"flag"
	"html"
	"log"
	"net/http"
	"strings"

	"carrel.dev/typecase"
	"carrel.dev/typecase/internal/screens"
)

func main() {
	addr := flag.String("addr", ":8085", "listen address")
	flag.Parse()
	h, err := handler()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("screens on http://localhost%s/", *addr)
	log.Fatal(http.ListenAndServe(*addr, h))
}

func handler() (http.Handler, error) {
	cssText, err := typecase.CSS()
	if err != nil {
		return nil, err
	}
	js, err := typecase.JS()
	if err != nil {
		return nil, err
	}
	theme, err := typecase.ThemeJS()
	if err != nil {
		return nil, err
	}
	sprite, err := typecase.Sprite()
	if err != nil {
		return nil, err
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/css/typecase.css", text("text/css; charset=utf-8", cssText))
	mux.HandleFunc("/screens.css", text("text/css; charset=utf-8", screens.PageCSS))
	mux.HandleFunc("/tc.js", text("application/javascript; charset=utf-8", string(js)))
	mux.HandleFunc("/theme.js", text("application/javascript; charset=utf-8", string(theme)))
	mux.Handle("/fonts/", http.FileServerFS(typecase.Fonts()))
	mux.HandleFunc("/f/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/f/")
		i := index(id)
		if i < 0 {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(page(i, string(sprite))))
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(list()))
	})
	return mux, nil
}

func text(ctype, body string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", ctype)
		_, _ = w.Write([]byte(body))
	}
}

func index(id string) int {
	for i, r := range screens.Routes {
		if r.Frame == id {
			return i
		}
	}
	return -1
}

const head = `<!doctype html><html lang="en"><head><meta charset="utf-8">` +
	`<meta name="viewport" content="width=device-width, initial-scale=1">` +
	`<script src="/theme.js"></script>` +
	`<link rel="stylesheet" href="/css/typecase.css"><link rel="stylesheet" href="/screens.css">`

// page is one screen under the demo's own bar: back to the list, the
// neighbours, the frame it is.
func page(i int, sprite string) string {
	r := screens.Routes[i]
	var b strings.Builder
	b.WriteString(head + `<title>` + html.EscapeString(r.Caption) + `</title></head><body>`)
	if r.Theme != "" {
		// The frame is drawn in a theme: so is the page, whatever the viewer chose.
		b.WriteString(`<script>document.documentElement.setAttribute("data-tc-theme", "` + r.Theme + `")</script>`)
	}
	b.WriteString(`<nav class="screens-bar"><a href="/">All screens</a>`)
	if i > 0 {
		p := screens.Routes[i-1]
		b.WriteString(`<a href="/f/` + p.Frame + `" rel="prev">← ` + html.EscapeString(p.Caption) + `</a>`)
	}
	if i+1 < len(screens.Routes) {
		n := screens.Routes[i+1]
		b.WriteString(`<a href="/f/` + n.Frame + `" rel="next">` + html.EscapeString(n.Caption) + ` →</a>`)
	}
	b.WriteString(`<span class="screens-bar__frame">` + html.EscapeString(r.Caption) + ` <code>` + r.Frame + `</code></span></nav>`)
	b.WriteString(`<div class="screens-sprite" hidden>` + sprite + `</div>`)
	b.WriteString(`<main class="screens-stage">` + r.Screen().Render() + `</main>`)
	b.WriteString(`<script src="/tc.js"></script></body></html>`)
	return b.String()
}

// list is every route in the mockup's order.
func list() string {
	var b strings.Builder
	b.WriteString(head + `<title>Screens</title></head><body><main class="screens-list"><h1>Screens</h1><ol>`)
	for _, r := range screens.Routes {
		b.WriteString(`<li><a href="/f/` + r.Frame + `">` + html.EscapeString(r.Caption) + `</a> <code>` + r.Frame + `</code></li>`)
	}
	b.WriteString(`</ol></main></body></html>`)
	return b.String()
}
