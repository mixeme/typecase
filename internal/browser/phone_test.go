package browser_test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"

	"carrel.dev/typecase/internal/browser"
)

// touch makes the tab a phone's: a finger, no hover (pointer:coarse,
// hover:none).
func touch() chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		if err := emulation.SetTouchEmulationEnabled(true).WithMaxTouchPoints(5).Do(ctx); err != nil {
			return err
		}
		return emulation.SetEmulatedMedia().WithFeatures([]*emulation.MediaFeature{
			{Name: "pointer", Value: "coarse"}, {Name: "hover", Value: "none"},
			{Name: "any-pointer", Value: "coarse"}, {Name: "any-hover", Value: "none"},
		}).Do(ctx)
	})
}

// TestTouchFields: under a finger every field is 16 px — the monospaced
// ones too, or the phone zooms the page when one is touched — and the key
// hints are gone: there is no keyboard (Carrel's task 297).
func TestTouchFields(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	body := `<div class="tc-root"><input class="tc-in" id="plain" value="Anna"><input class="tc-in is-mono" id="mono" value="+7 912">` +
		`<label class="tc-in"><input class="tc-q" id="q" placeholder="Filter"></label><kbd class="tc-kbd" id="key">Esc</kbd></div>`
	const probe = `(() => { const f = id => getComputedStyle(document.getElementById(id));
		return [matchMedia("(pointer:coarse) and (hover:none)").matches, f("plain").fontSize, f("mono").fontSize, f("q").fontSize, f("key").display].join(" "); })()`
	var mouse, finger string
	if err := h.Run(h.RegisterPage("touch-fields", browser.PageHTML(body, "")), 900, 400,
		chromedp.Evaluate(probe, &mouse),
		touch(),
		chromedp.Evaluate(probe, &finger),
	); err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(finger, "true ") {
		t.Fatalf("the tab is not a phone's under emulation: %q", finger)
	}
	if want := "true 16px 16px 16px none"; finger != want {
		t.Errorf("under a finger: got %q, want %q (with a mouse: %q)", finger, want, mouse)
	}
	if strings.HasSuffix(mouse, " none") {
		t.Errorf("with a mouse the key hint is hidden: %q", mouse)
	}
}

// TestFocusRing: a field, a field with a field inside (the filter) and a
// button reached by the keyboard are ringed in the accent, as is-focus
// draws them, not in the browser's blue; the inner field has no ring of
// its own (Carrel's task 297).
func TestFocusRing(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	body := `<div class="tc-root" style="padding:20px"><input class="tc-in" id="a" value="x"> <label class="tc-in" id="box"><input class="tc-q" id="q"></label> <button class="tc-btn" id="b" type="button">Save</button></div>`
	ring := func(id string) string {
		return `(() => { const c = getComputedStyle(document.getElementById("` + id + `")); return c.outlineStyle + " " + c.outlineColor; })()`
	}
	var a, box, q, b, accent string
	if err := h.Run(h.RegisterPage("focus-ring", browser.PageHTML(body, "")), 900, 300,
		chromedp.Evaluate(`(() => { const s = document.createElement("span"); s.style.color = "var(--tc-accent)"; document.querySelector(".tc-root").append(s); return getComputedStyle(s).color; })()`, &accent),
		chromedp.KeyEvent(kb.Tab),
		chromedp.Evaluate(ring("a"), &a),
		chromedp.KeyEvent(kb.Tab),
		chromedp.Evaluate(ring("box"), &box),
		chromedp.Evaluate(ring("q"), &q),
		chromedp.KeyEvent(kb.Tab),
		chromedp.Evaluate(ring("b"), &b),
	); err != nil {
		t.Fatal(err)
	}
	want := "solid " + accent
	for _, c := range []struct{ name, got, want string }{
		{"field", a, want}, {"filter's box", box, want}, {"filter's own field", q, "none " + accent}, {"button", b, want},
	} {
		if c.name == "filter's own field" {
			if !strings.HasPrefix(c.got, "none") {
				t.Errorf("%s: ring %q, want none", c.name, c.got)
			}
			continue
		}
		if c.got != c.want {
			t.Errorf("%s: ring %q, want %q", c.name, c.got, c.want)
		}
	}
}

// TestHeadFolds: a page head's buttons stand in one line; where they would
// wrap the secondary ones go to their icons and the primary keeps its
// word; where there is room every label stays (Carrel's task 297).
func TestHeadFolds(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	head := `<div class="tc-head"><div class="tc-head__titles"><div class="tc-head__title">Combined address book</div></div><div class="tc-acts">` +
		`<span class="tc-split"><button class="tc-btn is-primary" type="button"><svg viewBox="0 0 24 24"><use href="#i-plus"/></svg>New contact</button><button class="tc-btn is-primary" type="button" aria-label="More to create"><svg class="tc-ico--alone" viewBox="0 0 24 24"><use href="#i-chev"/></svg></button></span>` +
		`<button class="tc-btn" type="button"><svg viewBox="0 0 24 24"><use href="#i-import"/></svg>Import</button>` +
		`<button class="tc-btn" type="button"><svg viewBox="0 0 24 24"><use href="#i-export"/></svg>Export</button>` +
		`<button class="tc-btn" type="button"><svg viewBox="0 0 24 24"><use href="#i-print"/></svg>Print</button></div></div>`
	page := h.RegisterPage("head-folds", browser.PageHTML(`<div class="tc-root tc-shell"><div class="tc-shell__body tc-shell__body--no-rail"><div class="tc-main">`+head+`</div></div></div>`, ""))
	const probe = `(() => { const a = document.querySelector(".tc-head > .tc-acts"), tops = new Set([...a.children].map(k => Math.round(k.getBoundingClientRect().top)));
		const p = document.querySelector(".tc-split .is-primary"), s = a.querySelector(":scope > .tc-btn");
		return (a.getAttribute("data-tc-fold") || "-") + " lines " + tops.size + " primary " + getComputedStyle(p).fontSize + " import " + getComputedStyle(s).fontSize; })()`
	for _, c := range []struct {
		width int
		fold  string
	}{{360, "labels"}, {1200, "-"}} {
		var got string
		if err := h.Run(page, c.width, 600, chromedp.Evaluate(probe, &got)); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(got, " lines 1 ") {
			t.Errorf("at %d the head's buttons wrap: %s", c.width, got)
		}
		if f := strings.Fields(got)[0]; f != c.fold && !(c.fold == "-" && f == "") {
			t.Errorf("at %d the head folds %q, want %q: %s", c.width, f, c.fold, got)
		}
		if strings.Contains(got, "primary 0px") {
			t.Errorf("at %d the primary button lost its word: %s", c.width, got)
		}
	}
}

// TestDocHeadingAfterTasks: a heading of a note's text that has no tasks of
// its own keeps the space above it after a list of tasks, as a heading
// with tasks does after the text (Carrel's task 297).
func TestDocHeadingAfterTasks(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	// what DocTasks, DocHeading and Prose print
	parts := `<div class="tc-doc__sec"><div class="tc-doc__heading">Что делаем</div><ul class="tc-stack tc-stack--normal tc-doc__list">` +
		`<li class="tc-doc__task"><span class="tc-tick tc-tick--nudge" role="img" aria-label="Not done"></span>проверить changelog</li></ul></div>` +
		`<div class="tc-doc__heading">Отложено до следующего спринта</div>` +
		`<div class="tc-prose tc-prose--lead tc-prose--above-loose">Перенос старых вложений в новую папку.</div>`
	body := `<div class="tc-root"><div class="tc-doc"><div class="tc-doc__title">Заметка</div>` + parts + `</div></div>`
	const probe = `(() => { const l = document.querySelector(".tc-doc__list").getBoundingClientRect(), hs = document.querySelectorAll(".tc-doc__heading"), h = hs[hs.length - 1].getBoundingClientRect();
		return Math.round(h.top - l.bottom); })()`
	var gap int
	if err := h.Run(h.RegisterPage("doc-heading", browser.PageHTML(body, "")), 900, 600, chromedp.Evaluate(probe, &gap)); err != nil {
		t.Fatal(err)
	}
	if gap < 16 {
		t.Errorf("the heading after the tasks stands %d px under them, want at least 16", gap)
	}
}

// TestDrawerWithoutStyleQueries: a browser that does not know
// @container style() (Firefox 132 on a phone) skips the drawer's block; its
// copy under the window's width still gives a phone the drawer's look — the
// note's side column under the text, rows 52 px — and a wide window keeps
// the desktop's (Carrel's task 299).
func TestDrawerWithoutStyleQueries(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	data, err := os.ReadFile(filepath.Join(browser.ModuleRoot(t), "css", "typecase.css"))
	if err != nil {
		t.Fatal(err)
	}
	css := string(data)
	const q = "@container style(--tc-shell-mode: drawer){"
	i := strings.Index(css, q)
	if i < 0 {
		t.Fatal("no drawer block in typecase.css")
	}
	depth, end := 0, -1
	for j := i + len(q) - 1; j < len(css); j++ {
		if css[j] == '{' {
			depth++
		} else if css[j] == '}' {
			if depth--; depth == 0 {
				end = j
				break
			}
		}
	}
	old := css[:i] + css[end+1:] // what an old browser reads
	body := `<div class="tc-root tc-shell"><div class="tc-shell__body tc-shell__body--no-rail"><div class="tc-main">` +
		`<div class="tc-doc__layout"><div class="tc-doc">Text of the note</div><aside class="tc-docmeta">Side</aside></div>` +
		`<div class="tc-list"><a class="tc-row" href="#r">Row</a></div></div></div></div>`
	page := func(id, sheet string) string {
		return h.RegisterPage(id, `<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1">`+
			`<link rel="stylesheet" href="`+h.RegisterCSS(id+".css", sheet)+`"></head><body data-tc-root style="margin:0">`+body+`</body></html>`)
	}
	const probe = `(() => { const l = document.querySelector(".tc-doc__layout"), r = document.querySelector(".tc-row");
		return getComputedStyle(l).gridTemplateColumns.split(" ").length + " cols, row " + getComputedStyle(r).minHeight; })()`
	for _, c := range []struct {
		name, sheet string
		width       int
		want        string
	}{
		{"with style()", css, 360, "1 cols, row 52px"},
		{"without style()", old, 360, "1 cols, row 52px"},
		{"without style(), wide", old, 1200, "2 cols"},
	} {
		var got string
		if err := h.Run(page(strings.ReplaceAll(c.name, " ", "-"), c.sheet), c.width, 700, chromedp.Evaluate(probe, &got)); err != nil {
			t.Fatal(err)
		}
		if !strings.HasPrefix(got, c.want) {
			t.Errorf("%s at %d: %s, want %s", c.name, c.width, got, c.want)
		}
	}
}
