package browser_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/chromedp/cdproto/runtime"
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

// TestAppWorker (Carrel's task 91): a document that names an application
// prints its manifest, its colour and its icon; tc.js registers the
// library's worker, which keeps the library's files and the offline page and
// nothing else; with the server gone, a page that loads is the offline page.
func TestAppWorker(t *testing.T) {
	needChrome(t)
	doc := page.Document{Title: "App", Assets: page.At("/static/tc"),
		App:  &page.App{Manifest: "/manifest.json", ThemeColor: "#2f4a36", Icon: "/icon.png", Offline: "/offline"},
		Root: w.Page{Items: []w.PagePart{w.Msg{Parts: []w.MsgPart{w.Text("Online")}}}}}
	html, err := doc.Render()
	if err != nil {
		t.Fatal(err)
	}
	var down atomic.Bool
	mux := http.NewServeMux()
	mux.Handle("/static/tc/", http.StripPrefix("/static/tc", page.Static()))
	mux.HandleFunc("/offline", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = rw.Write([]byte(`<!doctype html><title>Offline</title><p>You are offline</p>`))
	})
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		if down.Load() {
			// the network is gone: the connection breaks
			if c, _, err := rw.(http.Hijacker).Hijack(); err == nil {
				c.Close()
			}
			return
		}
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = rw.Write([]byte(html))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()
	h := browser.NewHarness(t)
	ready := chromedp.ActionFunc(func(ctx context.Context) error {
		_, exc, err := runtime.Evaluate(`navigator.serviceWorker.ready.then(() => new Promise(r => setTimeout(r, 300)))`).WithAwaitPromise(true).Do(ctx)
		if exc != nil {
			return fmt.Errorf("%s", exc.Text)
		}
		return err
	})
	var head, kept, gone string
	err = h.Run(srv.URL+"/one", 900, 600,
		chromedp.Evaluate(`[...document.head.querySelectorAll('meta[name="theme-color"], link[rel="manifest"], link[rel="icon"]')].map(e => e.getAttribute("content") || e.getAttribute("href")).join(" ")`, &head),
		ready,
		chromedp.ActionFunc(func(ctx context.Context) error {
			v, exc, err := runtime.Evaluate(`caches.open("tc-offline").then(c => c.keys()).then(ks => ks.map(k => new URL(k.url).pathname).sort().join(" "))`).WithAwaitPromise(true).Do(ctx)
			if err != nil || exc != nil {
				return fmt.Errorf("caches: %v %v", err, exc)
			}
			kept = string(v.Value)
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error { down.Store(true); return nil }),
		chromedp.Navigate(srv.URL+"/two"),
		chromedp.Evaluate(`document.body.textContent.trim()`, &gone),
	)
	if err != nil {
		t.Fatal(err)
	}
	if head != "#2f4a36 /manifest.json /icon.png" {
		t.Errorf("the head names %q; want the colour, the manifest and the icon", head)
	}
	if kept != `"/offline /static/tc/css/typecase.css /static/tc/js/tc.js /static/tc/js/theme.js"` {
		t.Errorf("the worker keeps %s; want the offline page and the library's files only", kept)
	}
	if gone != "You are offline" {
		t.Errorf("with the server gone a page shows %q; want the offline page", gone)
	}
}
