package browser

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/chromedp/chromedp"

	"carrel.dev/typecase"
	"carrel.dev/typecase/internal/gate"
)

// Rect is a DOM bounding box in CSS pixels.
type Rect struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Top    float64 `json:"top"`
	Left   float64 `json:"left"`
	Bottom float64 `json:"bottom"`
	Right  float64 `json:"right"`
}

// Harness serves typecase assets and runs chromedp actions.
type Harness struct {
	t       *testing.T
	server  *httptest.Server
	pages   map[string]string
	streams map[string]http.HandlerFunc
	browser *instance
}

// NewHarness starts a test server with bundled CSS and JS.
func NewHarness(t *testing.T) *Harness {
	t.Helper()
	css, err := typecase.CSS()
	if err != nil {
		t.Fatalf("css: %v", err)
	}
	theme, err := typecase.ThemeJS()
	if err != nil {
		t.Fatalf("theme js: %v", err)
	}
	js, err := typecase.JS()
	if err != nil {
		t.Fatalf("js: %v", err)
	}
	live, err := typecase.LiveJS()
	if err != nil {
		t.Fatalf("js: %v", err)
	}

	h := &Harness{t: t, pages: make(map[string]string), streams: make(map[string]http.HandlerFunc)}
	mux := http.NewServeMux()
	root, err := gate.FindModuleRoot()
	if err != nil {
		t.Fatalf("module root: %v", err)
	}
	mockupFonts := filepath.Join(root, "docs", "visual", "fonts")
	mux.Handle("/page/fonts/", http.StripPrefix("/page/fonts/", http.FileServer(http.Dir(mockupFonts))))
	mux.HandleFunc("/typecase.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		_, _ = w.Write([]byte(css))
	})
	mux.HandleFunc("/theme.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		_, _ = w.Write(theme)
	})
	mux.HandleFunc("/tc.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		_, _ = w.Write(js)
	})
	mux.HandleFunc("/tc-live.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		_, _ = w.Write(live)
	})
	mux.HandleFunc("/page/", func(w http.ResponseWriter, r *http.Request) {
		id, err := url.PathUnescape(r.URL.Path[len("/page/"):])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		html, ok := h.pages[id]
		if !ok {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(html))
	})
	mux.HandleFunc("/stream/", func(w http.ResponseWriter, r *http.Request) {
		if fn, ok := h.streams[r.URL.Path[len("/stream/"):]]; ok {
			fn(w, r)
			return
		}
		http.NotFound(w, r)
	})
	mux.HandleFunc("/css/", func(w http.ResponseWriter, r *http.Request) {
		id, err := url.PathUnescape(r.URL.Path[len("/css/"):])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		css, ok := h.pages[id]
		if !ok {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		_, _ = w.Write([]byte(css))
	})

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	h.server = httptest.NewUnstartedServer(mux)
	h.server.Listener = ln
	h.server.Start()
	t.Cleanup(h.server.Close)

	in, err := launchChrome()
	if err != nil {
		t.Fatalf("launch %s: %v", ChromePath(), err)
	}
	h.browser = in
	t.Cleanup(in.stop)
	return h
}

// CloseBrowser kills the harness Chrome process. Safe to call more than once.
func (h *Harness) CloseBrowser() {
	if h.browser != nil {
		h.browser.stop()
	}
}

// RegisterPage stores HTML and returns its URL.
func (h *Harness) RegisterPage(id, html string) string {
	h.pages[id] = html
	return h.server.URL + "/page/" + url.PathEscape(id)
}

// RegisterStream serves a handler of the test's own — a stream of events —
// and returns its URL.
func (h *Harness) RegisterStream(id string, fn http.HandlerFunc) string {
	h.streams[id] = fn
	return h.server.URL + "/stream/" + url.PathEscape(id)
}

// RegisterCSS stores a stylesheet and returns its URL.
func (h *Harness) RegisterCSS(id, css string) string {
	h.pages[id] = css
	return h.server.URL + "/css/" + url.PathEscape(id)
}

// Run executes chromedp actions against a registered page URL.
//
// Each call gets a fresh tab in the harness browser; the viewport is set with
// device metrics rather than a window-size flag, so one process serves every
// check.
func (h *Harness) Run(pageURL string, width, height int, actions ...chromedp.Action) error {
	h.t.Helper()
	if width <= 0 {
		width = 800
	}
	if height <= 0 {
		height = 600
	}

	for try := 1; ; try++ {
		err := h.run(pageURL, width, height, actions)
		if err == nil {
			return nil
		}
		if !errors.Is(err, errNotOpened) || try == 2 {
			return fmt.Errorf("chromedp: %w", err)
		}
		// On a long run the browser now and then loses a new tab before the
		// page is in it (14.10.2026: TestBarSafetyQuiet, TestBite_catalogSeg);
		// the page is opened again in another tab. What the test does in the
		// tab is not tried again.
		h.t.Logf("browser: %s did not open in %s; opened again in a new tab", pageURL, openTimeout)
	}
}

// openTimeout is how long a new tab may take to have its page: tens of
// milliseconds as a rule, seconds for the catalog of every record, and more
// than 30 s for it late in a full run (14.10.2026: the catalog tests).
const openTimeout = 60 * time.Second

// errNotOpened is a tab that did not have its page in openTimeout.
var errNotOpened = errors.New("the tab did not open the page")

// run opens the page in a new tab and does the actions there.
func (h *Harness) run(pageURL string, width, height int, actions []chromedp.Action) error {
	allocCtx, cancelAlloc := chromedp.NewRemoteAllocator(context.Background(), h.browser.endpoint)
	defer cancelAlloc()
	ctx, cancelCtx := chromedp.NewContext(allocCtx)
	defer cancelCtx()

	opened := make(chan error, 1)
	go func() {
		opened <- chromedp.Run(ctx,
			chromedp.EmulateViewport(int64(width), int64(height)),
			chromedp.Navigate(pageURL),
			chromedp.WaitReady("body"),
		)
	}()
	select {
	case err := <-opened:
		if err != nil {
			return err
		}
	case <-time.After(openTimeout):
		cancelCtx()
		select {
		case <-opened:
		case <-time.After(5 * time.Second): // a tab that ignores the cancel is left behind
		}
		return errNotOpened
	}
	// what the test does has its own 90 s, however long the page took
	ctx, cancelTimeout := context.WithTimeout(ctx, 90*time.Second)
	defer cancelTimeout()
	return chromedp.Run(ctx, actions...)
}

// Tab runs chromedp actions in one tab until cancel is called.
func (h *Harness) Tab(timeout time.Duration) (run func(actions ...chromedp.Action) error, cancel func()) {
	h.t.Helper()
	if timeout <= 0 {
		timeout = 90 * time.Second
	}
	allocCtx, cancelAlloc := chromedp.NewRemoteAllocator(context.Background(), h.browser.endpoint)
	ctx, cancelCtx := chromedp.NewContext(allocCtx)
	ctx, cancelTimeout := context.WithTimeout(ctx, timeout)
	var once sync.Once
	cancel = func() {
		once.Do(func() {
			cancelTimeout()
			cancelCtx()
			cancelAlloc()
		})
	}
	// CaptureScreenshot on Edge can ignore context cancel. If the tab
	// deadline fires and chromedp still does not return, kill the OS
	// process tree so the test can fail instead of hanging forever.
	go func() {
		<-ctx.Done()
		if ctx.Err() != context.DeadlineExceeded {
			return
		}
		time.Sleep(2 * time.Second)
		h.browser.stop()
	}()
	h.t.Cleanup(cancel)
	run = func(actions ...chromedp.Action) error {
		if err := chromedp.Run(ctx, actions...); err != nil {
			return fmt.Errorf("chromedp: %w", err)
		}
		return nil
	}
	return run, cancel
}

// PageHTML builds a minimal document around a demo fragment.
func PageHTML(body string, extraCSS string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en" data-tc-root data-tc-noscript>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<script src="/theme.js"></script>
<link rel="stylesheet" href="/typecase.css">
<style>
body { margin: 0; padding: 16px; background: var(--tc-bg); color: var(--tc-ink); font-family: var(--tc-font-sans); }
.demo { max-width: 100%%; }
%s
</style>
</head>
<body data-tc-root>
<div class="demo">%s</div>
<script src="/tc.js"></script>
<script src="/tc-live.js"></script>
</body>
</html>`, extraCSS, body)
}

// ModuleRoot returns the typecase module directory.
func ModuleRoot(t *testing.T) string {
	t.Helper()
	root, err := gate.FindModuleRoot()
	if err != nil {
		t.Fatal(err)
	}
	return root
}
