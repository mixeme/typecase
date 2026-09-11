package proof

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"carrel.dev/typecase"
	"carrel.dev/typecase/internal/mockcss"
)

// Case is one frame of the mockup and the screen that transliterates it.
type Case struct {
	Name   string `json:"name"`
	Label  string `json:"label"`
	Path   string `json:"path"`
	screen func() Screen
}

// Cases are the frames of the proof.
var Cases = []Case{
	{Name: "notes", Label: "Заметки · список и редактор", Path: "/p/notes", screen: Notes},
	{Name: "files", Label: "Файлы · обзор папки, свойства выбранного", Path: "/p/files", screen: Files},
	{Name: "dark", Label: "Календарь · тёмная тема", Path: "/p/dark", screen: Dark},
	{Name: "quick", Label: "Быстрая заметка поверх текущего экрана", Path: "/p/quick", screen: Quick},
	{Name: "delete", Label: "Удаление коллекции", Path: "/p/delete", screen: Delete},
	{Name: "conn", Label: "Настройки · подключения и подписки", Path: "/p/conn", screen: Connections},
}

// tokens maps mockup custom properties to library tokens (same values, one
// exception: --dusty is #6A6351 in the mockup, #6f6856 in Typecase).
var tokens = map[string]string{
	"ink": "tc-ink", "card": "tc-card", "card-border": "tc-edge", "accent": "tc-accent",
	"bg": "tc-bg", "paper": "tc-surface", "rule": "tc-rule", "hair": "tc-hair",
	"dusty": "tc-dusty", "muted": "tc-muted", "alert": "tc-danger",
	"f-serif": "tc-font-serif", "f-sans": "tc-font-sans", "f-mono": "tc-font-mono",
	"pad": "tc-page-gutter",
}

// elements: where a widget prints another element than the frame.
var elements = map[string]string{
	".m-seg span":   ".m-seg :is(span,button)",   // options are buttons
	".m-menu > div": ".m-menu > :is(div,button)", // menu items are buttons
}

// selectors: mockup devices that mean something else in the library.
var selectors = map[string]string{
	".is-dark ": "[data-tc-theme=dark] ", // the frame class is the theme attribute
}

// dropAt: mockup rules Typecase does not carry, with the reason.
var dropAt = []string{
	"@container (max-width:560px)", // rail 56 px; Typecase has the drawer (spec L-4 tc-shell)
}

// dropRules: mockup selectors not carried over, with the reason.
var dropRules = []string{
	".mock.is-dark", // dark token values come from the library's own tokens
}

// extraClasses: mockup classes a widget prints although its frame does not
// use them, with the reason.
var extraClasses = []string{
	"m-drop", // Btn with a menu prints the mockup's .m-drop; the frame nests the menu in .m-btn
}

// Options builds the generator input from the frames themselves: the classes
// kept are exactly the classes the frames use.
func Options(mockup string) (mockcss.Options, error) {
	keep := map[string]bool{"mock": true}
	for _, c := range extraClasses {
		keep[c] = true
	}
	for _, c := range Cases {
		frame := mockcss.Frame(mockup, c.Label)
		if frame == "" {
			return mockcss.Options{}, fmt.Errorf("proof: frame %q not found", c.Label)
		}
		for cls := range mockcss.ClassesIn(frame) {
			keep[cls] = true
		}
	}
	names := map[string]string{}
	for c := range keep {
		switch {
		case c == "mock":
			names[c] = "tcx"
		case strings.HasPrefix(c, "m-"):
			names[c] = "tcx-" + strings.TrimPrefix(c, "m-")
		}
	}
	return mockcss.Options{Keep: keep, Classes: names, Vars: tokens, Prefix: "tcx-",
		Elements: elements, Selectors: selectors, DropAt: dropAt, DropRules: dropRules}, nil
}

const fontFaces = `@font-face{font-family:"Literata";font-weight:200 900;src:url("/fonts/Literata.ttf") format("truetype")}
@font-face{font-family:"Public Sans";font-weight:100 900;src:url("/fonts/PublicSans.ttf") format("truetype")}
@font-face{font-family:"Source Code Pro";font-weight:200 900;src:url("/fonts/SourceCodePro.ttf") format("truetype")}
`

// Mux serves the proof screens, the mockup with the same fonts, its sprite,
// and the comparison page.
func Mux(root string) (*http.ServeMux, error) {
	raw, err := os.ReadFile(filepath.Join(root, "docs", "visual", "carrel-ui-mockups.html"))
	if err != nil {
		return nil, err
	}
	mockup := string(raw)
	o, err := Options(mockup)
	if err != nil {
		return nil, err
	}
	generated := mockcss.Generate(mockcss.Style(mockup), o)
	tokensCSS, err := os.ReadFile(filepath.Join(root, "css", "00-tokens.css"))
	if err != nil {
		return nil, err
	}
	js, err := typecase.JS()
	if err != nil {
		return nil, err
	}
	spriteSVG, err := mockupSprite(mockup)
	if err != nil {
		return nil, err
	}
	proofCSS := string(tokensCSS) + fontFaces +
		"\n/* generated from the mockup <style> */\n" + generated +
		"\n/* controls the frames draw as text */\n" + controlCSS +
		"\n/* frame style=\"\" that became widget variants */\n" + inlineCSS(o)
	cfg, err := json.Marshal(map[string]any{"cases": Cases, "ignore": inlineClasses()})
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	text := func(mime, body string) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", mime)
			_, _ = w.Write([]byte(body))
		}
	}
	mux.HandleFunc("/proof.css", text("text/css; charset=utf-8", proofCSS))
	mux.HandleFunc("/generated.css", text("text/css; charset=utf-8", generated))
	mux.HandleFunc("/tc.js", text("application/javascript; charset=utf-8", string(js)))
	mux.HandleFunc("/mock-sprite.svg", text("image/svg+xml; charset=utf-8", spriteSVG))
	mux.Handle("/fonts/", http.FileServerFS(typecase.Fonts()))
	for _, c := range Cases {
		mux.HandleFunc(c.Path, text("text/html; charset=utf-8", page(c)))
	}
	mux.HandleFunc("/mockup", text("text/html; charset=utf-8", mockupWithFonts(mockup)))
	mux.HandleFunc("/compare", text("text/html; charset=utf-8", strings.Replace(comparePage, "/*CONFIG*/", string(cfg), 1)))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, "/compare", http.StatusFound)
	})
	return mux, nil
}

func page(c Case) string {
	s := c.screen()
	theme := ""
	if s.Dark {
		theme = ` data-tc-theme="dark"`
	}
	return `<!doctype html><html lang="ru"` + theme + `><head><meta charset="utf-8"><title>Proof · ` + c.Name + `</title>
<link rel="stylesheet" href="/proof.css">
<style>body{margin:0;background:var(--tc-bg)}</style>
</head><body>` + s.Render() + `<script src="/tc.js"></script></body></html>`
}

// mockupWithFonts gives the mockup the library font files and mono stack, so
// text measures the same on both sides. Layout rules are untouched.
func mockupWithFonts(html string) string {
	inject := "<style>" + fontFaces + `:root{--f-mono:"Source Code Pro",ui-monospace,"Cascadia Mono",monospace}</style>`
	return strings.Replace(html, "</head>", inject+"</head>", 1)
}

// mockupSprite lifts the mockup's inline <defs> into a standalone sprite.
func mockupSprite(html string) (string, error) {
	start := strings.Index(html, "<defs>")
	end := strings.Index(html, "</defs>")
	if start < 0 || end < start {
		return "", fmt.Errorf("proof: no <defs> in mockup")
	}
	return `<svg xmlns="http://www.w3.org/2000/svg">` + html[start:end+len("</defs>")] + `</svg>`, nil
}
