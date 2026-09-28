package browser_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chromedp/chromedp"

	"carrel.dev/typecase/internal/browser"
	"carrel.dev/typecase/page"
	w "carrel.dev/typecase/widget"
)

// TestPageUnderStrictPolicy (task 38): a page of package page, its files
// served by page.Static, works under script-src 'self'; style-src 'self' —
// the theme is set, the stylesheet applies, the library's script runs, and
// nothing is refused.
func TestPageUnderStrictPolicy(t *testing.T) {
	needChrome(t)
	doc := page.Document{Title: "Page", Assets: page.At("/static/tc"),
		Root: w.Page{Items: []w.PagePart{w.Msg{Parts: []w.MsgPart{w.Text("Hello")}}}}}
	html, err := doc.Render()
	if err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.Handle("/static/tc/", http.StripPrefix("/static/tc", page.Static()))
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self'; font-src 'self'")
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = rw.Write([]byte(html))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()
	h := browser.NewHarness(t)
	var got []any
	err = h.Run(srv.URL+"/", 900, 600, chromedp.Evaluate(`(() => [
		document.documentElement.getAttribute("data-tc-theme"),
		getComputedStyle(document.body).marginTop,
		getComputedStyle(document.querySelector(".tc-root")).boxSizing,
		typeof window.__tc,
		document.fonts.check("12px 'Public Sans'"),
	])()`, &got))
	if err != nil {
		t.Fatal(err)
	}
	if got[0] != "light" && got[0] != "dark" {
		t.Errorf("theme %v: theme.js did not run", got[0])
	}
	if got[1] != "0px" {
		t.Errorf("body margin %v: the stylesheet did not apply", got[1])
	}
	if got[3] != "object" {
		t.Errorf("window.__tc is %v: tc.js did not run", got[3])
	}
}
