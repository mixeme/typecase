package browser_test

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	cdppage "github.com/chromedp/cdproto/page"
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

// await evaluates a script whose value is a promise and keeps what it gives.
func await(script string, got *string) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		v, exc, err := runtime.Evaluate(script).WithAwaitPromise(true).Do(ctx)
		if err != nil {
			return err
		}
		if exc != nil {
			return fmt.Errorf("%s: %s", script, exc.Text)
		}
		if got != nil {
			*got = string(v.Value)
		}
		return nil
	})
}

// TestAppWorkerUpdates (task 91, not checked before): the application
// serves the library at a new base — a new version — and the page that
// loads registers the worker anew; the new worker takes over and keeps the
// new files and the offline page, the files of the old base are gone.
func TestAppWorkerUpdates(t *testing.T) {
	needChrome(t)
	render := func(base string) string {
		doc := page.Document{Title: "App", Assets: page.At(base),
			App:  &page.App{Manifest: "/manifest.json", Offline: "/offline"},
			Root: w.Page{Items: []w.PagePart{w.Msg{Parts: []w.MsgPart{w.Text("Online")}}}}}
		html, err := doc.Render()
		if err != nil {
			t.Fatal(err)
		}
		return html
	}
	one, two := render("/static/1/tc"), render("/static/2/tc")
	var second atomic.Bool
	mux := http.NewServeMux()
	mux.Handle("/static/1/tc/", http.StripPrefix("/static/1/tc", page.Static()))
	mux.Handle("/static/2/tc/", http.StripPrefix("/static/2/tc", page.Static()))
	mux.HandleFunc("/offline", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = rw.Write([]byte(`<!doctype html><title>Offline</title><p>You are offline</p>`))
	})
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		if second.Load() {
			_, _ = rw.Write([]byte(two))
			return
		}
		_, _ = rw.Write([]byte(one))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()
	h := browser.NewHarness(t)
	// the worker at work and its copies made: the active worker's address
	// names the base, and the cache holds the files under it
	settled := func(base string) chromedp.Action {
		return await(fmt.Sprintf(`(async () => {
			for (let i = 0; i < 100; i++) {
				const r = await navigator.serviceWorker.getRegistration();
				const a = r && r.active, c = navigator.serviceWorker.controller;
				if (a && a.state === "activated" && a.scriptURL.includes(%[1]q) && c && c.scriptURL === a.scriptURL) return;
				await new Promise(r => setTimeout(r, 100));
			}
			throw new Error("no worker of " + %[1]q + " took over");
		})()`, base), nil)
	}
	keys := `caches.open("tc-offline").then(c => c.keys()).then(ks => ks.map(k => new URL(k.url).pathname).sort().join(" "))`
	var before, after string
	err := h.Run(srv.URL+"/one", 900, 600,
		settled("/static/1/tc/"),
		await(keys, &before),
		chromedp.ActionFunc(func(ctx context.Context) error { second.Store(true); return nil }),
		chromedp.Navigate(srv.URL+"/two"),
		settled("/static/2/tc/"),
		await(keys, &after),
	)
	if err != nil {
		t.Fatal(err)
	}
	if before != `"/offline /static/1/tc/css/typecase.css /static/1/tc/js/tc.js /static/1/tc/js/theme.js"` {
		t.Errorf("the first worker keeps %s; want the offline page and the files of the first base", before)
	}
	if after != `"/offline /static/2/tc/css/typecase.css /static/2/tc/js/tc.js /static/2/tc/js/theme.js"` {
		t.Errorf("the new worker keeps %s; want the offline page and the files of the new base only", after)
	}
}

// TestAppInstalls (task 91, not checked before): with the application's
// manifest and icon, a page whose document names them is one the browser
// offers to install — it finds nothing missing.
func TestAppInstalls(t *testing.T) {
	needChrome(t)
	doc := page.Document{Title: "App", Assets: page.At("/static/tc"),
		App:  &page.App{Manifest: "/manifest.json", ThemeColor: "#2f4a36", Icon: "/icon.png", Offline: "/offline"},
		Root: w.Page{Items: []w.PagePart{w.Msg{Parts: []w.MsgPart{w.Text("Online")}}}}}
	html, err := doc.Render()
	if err != nil {
		t.Fatal(err)
	}
	icon := func(size int) []byte {
		img := image.NewRGBA(image.Rect(0, 0, size, size))
		draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{0x2f, 0x4a, 0x36, 0xff}}, image.Point{}, draw.Src)
		var b bytes.Buffer
		if err := png.Encode(&b, img); err != nil {
			t.Fatal(err)
		}
		return b.Bytes()
	}
	big := icon(512)
	mux := http.NewServeMux()
	mux.Handle("/static/tc/", http.StripPrefix("/static/tc", page.Static()))
	mux.HandleFunc("/manifest.json", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/manifest+json")
		_, _ = rw.Write([]byte(`{"name":"App","short_name":"App","start_url":"/","scope":"/","display":"standalone",
			"background_color":"#2f4a36","theme_color":"#2f4a36",
			"icons":[{"src":"/icon.png","sizes":"512x512","type":"image/png"}]}`))
	})
	mux.HandleFunc("/icon.png", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "image/png")
		_, _ = rw.Write(big)
	})
	mux.HandleFunc("/offline", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = rw.Write([]byte(`<!doctype html><title>Offline</title><p>You are offline</p>`))
	})
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = rw.Write([]byte(html))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()
	h := browser.NewHarness(t)
	var errs []*cdppage.InstallabilityError
	err = h.Run(srv.URL+"/", 900, 600,
		await(`navigator.serviceWorker.ready.then(() => new Promise(r => setTimeout(r, 300)))`, nil),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			errs, err = cdppage.GetInstallabilityErrors().Do(ctx)
			return err
		}),
	)
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range errs {
		args := ""
		for _, a := range e.ErrorArguments {
			args += " " + a.Name + "=" + a.Value
		}
		t.Errorf("the browser would not install the page: %s%s", e.ErrorID, args)
	}
}

// TestShellFillsTheWindow (Carrel's task 272, the first deployment): on a
// page the shell is no lower than the window — the top bar and the body
// fill it, the rail stretches and its foot stands at the window's bottom;
// content taller than the window takes the foot down with the page, as the
// frames' shell does in its box. Frames are no page: the rule is the page's.
func TestShellFillsTheWindow(t *testing.T) {
	needChrome(t)
	shell := func(rows int) string {
		var items []w.MainPart
		items = append(items, w.Head{Title: &w.HeadTitle{Text: "Contacts"}})
		for i := 0; i < rows; i++ {
			items = append(items, w.Msg{Parts: []w.MsgPart{w.Text(fmt.Sprintf("Row %d", i+1))}})
		}
		doc := page.Document{Title: "Shell", Assets: page.At("/static/tc"),
			Root: w.Shell{Top: &w.Top{Brand: w.Brand{Title: "Contacts"}}, Body: w.ShellBody{
				Rail: []w.RailPart{w.Nav{Items: []w.NavItem{{Icon: "con", Label: "Contacts", Href: "#"}}, On: 1},
					w.RailFoot{About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.3.12 · 580f8fc9"}, Href: "#"}}},
				Main: w.Main{Items: items}}}}
		html, err := doc.Render()
		if err != nil {
			t.Fatal(err)
		}
		return html
	}
	short, long := shell(1), shell(60)
	sheet, err := page.Document{Title: "Sheet", Assets: page.At("/static/tc"),
		Root: w.Shell{Top: &w.Top{Brand: w.Brand{Title: "Contacts"}}, Body: w.Overlay{
			Body: w.OverlayBody{Rail: []w.OverlayRailPart{w.Nav{Items: []w.NavItem{{Icon: "con", Label: "Contacts", Href: "#"}}, On: 1},
				w.RailFoot{About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.3.12 · 580f8fc9"}, Href: "#"}}},
				Main: w.Main{Items: []w.MainPart{w.Head{Title: &w.HeadTitle{Text: "Contacts"}}}}},
			Scrim: w.Scrim{Content: w.Sheet{Title: "Rename", Foot: []w.SheetFootPart{w.Button{Label: "Save", Kind: w.ButtonPrimary}}}}}}}.Render()
	if err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/sheet", func(rw http.ResponseWriter, r *http.Request) { _, _ = rw.Write([]byte(sheet)) })
	mux.Handle("/static/tc/", http.StripPrefix("/static/tc", page.Static()))
	mux.HandleFunc("/short", func(rw http.ResponseWriter, r *http.Request) { _, _ = rw.Write([]byte(short)) })
	mux.HandleFunc("/long", func(rw http.ResponseWriter, r *http.Request) { _, _ = rw.Write([]byte(long)) })
	srv := httptest.NewServer(mux)
	defer srv.Close()
	h := browser.NewHarness(t)
	const measure = `(() => { const s = document.querySelector(".tc-shell").getBoundingClientRect(), f = document.querySelector(".tc-railfoot").getBoundingClientRect(),
		m = document.querySelector(".tc-main").getBoundingClientRect();
		return [Math.round(s.bottom), Math.round(f.bottom), Math.round(m.bottom), window.innerHeight, document.documentElement.scrollHeight]; })()`
	for _, c := range []struct{ w, h int }{{1920, 1040}, {1351, 800}} {
		var got []int
		if err := h.Run(srv.URL+"/short", c.w, c.h, chromedp.Evaluate(measure, &got)); err != nil {
			t.Fatal(err)
		}
		if got[0] != got[3] || got[1] != got[3] || got[4] > got[3] {
			t.Errorf("%d×%d, a short page: the shell ends at %d, the rail's foot at %d, the page is %d high; want all at the window's %d", c.w, c.h, got[0], got[1], got[4], got[3])
		}
		if err := h.Run(srv.URL+"/sheet", c.w, c.h, chromedp.Evaluate(measure, &got)); err != nil {
			t.Fatal(err)
		}
		if got[0] != got[3] || got[1] != got[3] {
			t.Errorf("%d×%d, a sheet's page: the shell ends at %d, the rail's foot under the scrim at %d; want both at the window's %d", c.w, c.h, got[0], got[1], got[3])
		}
		if err := h.Run(srv.URL+"/long", c.w, c.h, chromedp.Evaluate(measure, &got)); err != nil {
			t.Fatal(err)
		}
		if got[0] <= got[3] || got[1] != got[0] || got[2] > got[0] {
			t.Errorf("%d×%d, a long page: the shell ends at %d, the rail's foot at %d, the list at %d, the window %d; want the foot at the shell's end below the window", c.w, c.h, got[0], got[1], got[2], got[3])
		}
	}
}
