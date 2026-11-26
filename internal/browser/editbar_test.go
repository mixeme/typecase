package browser_test

import (
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"

	"carrel.dev/typecase/internal/browser"
	"carrel.dev/typecase/internal/screens"
	"github.com/chromedp/chromedp"
)

// editPage is the note's edit (s-notes-20) as Carrel prints it, with the
// library's script.
func editPage(t *testing.T, h *browser.Harness, id string) string {
	t.Helper()
	r, ok := screens.Find("s-notes-20")
	if !ok {
		t.Fatal("no route s-notes-20")
	}
	return h.RegisterPage(id, browser.PageHTML(r.Screen().Render(), ""))
}

// TestEditBarMarks (Carrel's task 286): B, I, H2, list, tick box and link
// put Markdown on the edit's text — around the selection or at the start of
// its lines — and come off when pressed again.
func TestEditBarMarks(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	url := editPage(t, h, "edit-marks")
	area := `document.querySelector("form#note-edit textarea.tc-in--body")`
	set := func(v string, a, b int) chromedp.Action {
		return chromedp.Evaluate(`(() => { const t = `+area+`; t.value = `+jsq(v)+`; t.focus(); t.setSelectionRange(`+itoa(a)+`, `+itoa(b)+`); return true; })()`, nil)
	}
	press := func(mark string) chromedp.Action {
		return chromedp.Click(`.tc-docbar [data-tc-marks="`+mark+`"]`, chromedp.ByQuery)
	}
	type step struct {
		value      string
		a, b       int
		mark, want string
		caret      int // where the caret stands after; -1: not checked
	}
	for _, s := range []step{
		{"hello world", 6, 11, "bold", "hello **world**", 15},
		{"hello **world**", 8, 13, "bold", "hello world", -1},
		{"hello world", 6, 11, "italic", "hello *world*", 13},
		{"hello **world**", 8, 13, "italic", "hello ***world***", -1},
		{"", 0, 0, "bold", "****", 2},
		{"hello world", 3, 3, "heading", "## hello world", 6},
		{"## hello world", 5, 5, "list", "- hello world", -1},
		{"- hello world", 0, 0, "tick", "- [ ] hello world", -1},
		{"- [ ] hello world", 0, 0, "tick", "hello world", -1},
		{"- [ ] hello", 0, 0, "list", "- hello", -1},
		{"one\ntwo\n\nthree", 0, 14, "list", "- one\n- two\n\n- three", 20},
		{"- one\n- two", 0, 11, "list", "one\ntwo", -1},
		{"see the docs", 8, 12, "link", "see the [docs]()", 15},
		{"go https://example.org now", 3, 22, "link", "go [](https://example.org) now", 4},
	} {
		var got string
		var caret int
		if err := h.Run(url, 1351, 900, set(s.value, s.a, s.b), press(s.mark),
			chromedp.Evaluate(area+`.value`, &got), chromedp.Evaluate(area+`.selectionEnd`, &caret)); err != nil {
			t.Fatal(err)
		}
		if got != s.want {
			t.Errorf("%s on %q [%d,%d]: %q, want %q", s.mark, s.value, s.a, s.b, got, s.want)
		}
		if s.caret >= 0 && caret != s.caret {
			t.Errorf("%s on %q: the caret at %d, want %d", s.mark, s.value, caret, s.caret)
		}
	}
}

// TestEditBarDraftAndAutoSave (Carrel's task 286): typed words are kept in
// the window as a draft, and a reload brings them back; with Auto-save
// ticked — kept in the browser — the edit goes after a pause by itself, as
// Save sends it, without leaving the page, and the form goes on to the
// address the answer came from (a new record's).
func TestEditBarDraftAndAutoSave(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	url := editPage(t, h, "edit-draft")
	var mu sync.Mutex
	var posts []string
	save := h.RegisterStream("edit-save", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			body, _ := io.ReadAll(r.Body)
			mu.Lock()
			posts = append(posts, r.Header.Get("Tc-Autosave")+" "+string(body))
			mu.Unlock()
			http.Redirect(rw, r, "/stream/edit-record", http.StatusSeeOther)
			return
		}
		rw.WriteHeader(http.StatusOK)
	})
	h.RegisterStream("edit-record", func(rw http.ResponseWriter, r *http.Request) { rw.WriteHeader(http.StatusOK) })
	area := `form#note-edit textarea.tc-in--body`
	note := `document.querySelector(".tc-docbar [data-tc-draft]")`
	var hiddenAtStart, shown, back, kept bool
	var noteText, restored string
	err := h.Run(url, 1351, 900,
		chromedp.Evaluate(`localStorage.clear(); sessionStorage.clear(); location.reload(); true`, nil),
		chromedp.Sleep(500*time.Millisecond),
		chromedp.Evaluate(note+`.hidden`, &hiddenAtStart),
		chromedp.Evaluate(`(() => { const t = document.querySelector("`+area+`"); t.value = ""; t.focus(); return true; })()`, nil),
		chromedp.SendKeys(area, "kept in the window", chromedp.ByQuery),
		chromedp.Sleep(2600*time.Millisecond),
		chromedp.Evaluate(`!`+note+`.hidden && /^Draft in this window/.test(`+note+`.textContent)`, &shown),
		chromedp.Reload(),
		chromedp.Sleep(500*time.Millisecond),
		chromedp.Evaluate(`document.querySelector("`+area+`").value`, &restored),
		chromedp.Evaluate(`!`+note+`.hidden`, &back),
		// Auto-save ticked, then the next pause writes the edit
		chromedp.Evaluate(`document.getElementById("note-edit").setAttribute("action", `+jsq(save)+`); true`, nil),
		chromedp.Click(`.tc-docbar [data-tc-autosave]`, chromedp.ByQuery),
		chromedp.Evaluate(`localStorage.getItem("tc-autosave:notes") === "on" && document.querySelector(".tc-docbar [data-tc-autosave]").getAttribute("aria-checked") === "true"`, &kept),
		chromedp.SendKeys(area, " and saved", chromedp.ByQuery),
		chromedp.Sleep(2800*time.Millisecond),
		chromedp.Evaluate(note+`.textContent`, &noteText),
	)
	if err != nil {
		t.Fatal(err)
	}
	if !hiddenAtStart {
		t.Error("the draft's note stands before there is a draft")
	}
	if !shown {
		t.Error("after a pause in typing the note does not say the draft is kept")
	}
	if restored != "kept in the window" || !back {
		t.Errorf("after a reload: the text %q, the note shown %v; want the draft back", restored, back)
	}
	if !kept {
		t.Error("Auto-save is not ticked and kept in the browser")
	}
	mu.Lock()
	got := append([]string(nil), posts...)
	mu.Unlock()
	if len(got) != 1 || !strings.HasPrefix(got[0], "1 ") || !strings.Contains(got[0], "kept in the window and saved") || !strings.Contains(got[0], `name="summary"`) {
		t.Errorf("auto-save sent %d edits: %q; want one, with the header and the form's fields", len(got), got)
	}
	if !strings.HasPrefix(noteText, "Saved · ") {
		t.Errorf("after the auto-save the note says %q, want «Saved · …»", noteText)
	}
	var still bool
	if err := h.Run(url, 1351, 900, chromedp.Evaluate(`document.querySelector(".tc-docbar [data-tc-autosave]").getAttribute("aria-checked") === "true"`, &still)); err != nil {
		t.Fatal(err)
	}
	if !still {
		t.Error("a new page of the edit forgot the Auto-save tick")
	}
}

// TestEditBarFocus (Carrel's task 286): Focus leaves the edit alone — no
// rail, no record's column; pressed again, or Esc from the text, brings
// them back.
func TestEditBarFocus(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	url := editPage(t, h, "edit-focus")
	state := `(() => { const shown = s => { const e = document.querySelector(s); return !!e && e.getClientRects().length > 0; };
		return [shown(".tc-shell__rail"), shown(".tc-docmeta"), document.querySelector("[data-tc-focus-mode]").getAttribute("aria-pressed") || ""].join(" "); })()`
	var before, focused, again, escaped string
	err := h.Run(url, 1351, 900,
		chromedp.Evaluate(state, &before),
		chromedp.Click(`[data-tc-focus-mode]`, chromedp.ByQuery),
		chromedp.Evaluate(state, &focused),
		chromedp.Click(`[data-tc-focus-mode]`, chromedp.ByQuery),
		chromedp.Evaluate(state, &again),
		chromedp.Click(`[data-tc-focus-mode]`, chromedp.ByQuery),
		chromedp.Focus(`form#note-edit textarea.tc-in--body`, chromedp.ByQuery),
		chromedp.KeyEvent("\x1b"),
		chromedp.Evaluate(state, &escaped),
	)
	if err != nil {
		t.Fatal(err)
	}
	if before != "true true " || focused != "false false true" || again != "true true false" || escaped != "true true false" {
		t.Errorf("rail, column, pressed: before %q, Focus %q, again %q, Esc %q", before, focused, again, escaped)
	}
}

func jsq(s string) string {
	r := strings.NewReplacer(`\`, `\\`, `"`, `\"`, "\n", `\n`)
	return `"` + r.Replace(s) + `"`
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	return string(b)
}

// TestReadingWidthOneMeasure (Carrel's task 284): Reading width is one
// measure for the whole column — the title, the fields and the text of an
// edit end at one edge on a wide window, where 78ch in each child's own size
// went apart (1920: the title 1089, the fields 620, the text 692).
func TestReadingWidthOneMeasure(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	url := editPage(t, h, "edit-measure")
	const widths = `(() => { document.querySelector(".tc-doc__layout").classList.add("is-measure");
		const d = document.querySelector("form#note-edit.tc-doc");
		const w = s => Math.round(d.querySelector(s).getBoundingClientRect().right);
		return [w(".tc-in--title"), w(".tc-doc__fields"), w(".tc-in--body")]; })()`
	for _, width := range []int{1920, 1351} {
		var got []int
		if err := h.Run(url, width, 1000, chromedp.Evaluate(widths, &got)); err != nil {
			t.Fatal(err)
		}
		if got[0] != got[1] || got[1] != got[2] {
			t.Errorf("at %d px with Reading width the title ends at %d, the fields at %d, the text at %d; want one edge", width, got[0], got[1], got[2])
		}
	}
}

// TestTestOnlyAnswersInItsZone (Carrel's task 282): Test only sends the
// connect form into the zone under the dialog's head, where Connect's refusal
// stands too — the outcome and the trace
// stand there, the form and the typed password stay as they were, the page
// does not go.
func TestTestOnlyAnswersInItsZone(t *testing.T) {
	needChrome(t)
	r, ok := screens.Find("s-profile-2")
	if !ok {
		t.Fatal("no route s-profile-2")
	}
	h := browser.NewHarness(t)
	var mu sync.Mutex
	var got []string
	test := h.RegisterStream("connect-answer", func(rw http.ResponseWriter, r *http.Request) {
		_ = r.ParseMultipartForm(1 << 20)
		mu.Lock()
		got = append(got, r.Header.Get("X-Target")+" "+r.FormValue("password"))
		mu.Unlock()
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = rw.Write([]byte(`<div class="tc-msg"><b>Test passed:</b> Carrel can sign in and found 8 collections.</div><div class="tc-diag">[collections] 8 collection(s)</div>`))
	})
	page := strings.Replace(browser.PageHTML(r.Screen().Render(), ""), `data-tc-noscript>`, `data-tc-noscript data-tc-view-header="X-View" data-tc-target-header="X-Target">`, 1)
	page = strings.Replace(page, `formaction="/app/settings/connections/new/test"`, `formaction="`+test+`"`, 1)
	var before, after string
	var shown bool
	err := h.Run(h.RegisterPage("connect", page), 1351, 900,
		chromedp.Evaluate(`[document.querySelector("[data-tc-zone=connect-answer]").hidden, document.querySelector("input[name=password]").value, location.pathname].join(" | ")`, &before),
		chromedp.Click(`button[data-tc-target="connect-answer"]`, chromedp.ByQuery),
		chromedp.Sleep(1500*time.Millisecond),
		chromedp.Evaluate(`(() => { const z = document.querySelector("[data-tc-zone=connect-answer]"); return !z.hidden && /Test passed/.test(z.textContent) && /8 collection/.test(z.textContent); })()`, &shown),
		chromedp.Evaluate(`[document.querySelector("[data-tc-zone=connect-answer]").hidden, document.querySelector("input[name=password]").value, location.pathname].join(" | ")`, &after),
	)
	if err != nil {
		t.Fatal(err)
	}
	mu.Lock()
	sent := append([]string(nil), got...)
	mu.Unlock()
	if len(sent) != 1 || sent[0] != "connect-answer xxxxxxxxxx" {
		t.Errorf("Test only sent %q; want one request for the zone connect-answer with the typed password", sent)
	}
	if !shown {
		t.Error("the outcome does not stand in its zone")
	}
	wantAfter := strings.Replace(before, "true | ", "false | ", 1)
	if !strings.HasPrefix(before, "true | xxxxxxxxxx") || after != wantAfter {
		t.Errorf("before %q, after %q; want the zone shown, the password and the page kept", before, after)
	}
}

// TestNeighboursShowTheOpenNote (Carrel's task 289): the rail's Neighbours
// holds the whole notebook and scrolls in its own block; the script puts the
// open note in the block's middle, or as near it as the block goes, and the
// rest of the rail does not move.
func TestNeighboursShowTheOpenNote(t *testing.T) {
	needChrome(t)
	r, ok := screens.Find("s-notes-2")
	if !ok {
		t.Fatal("no route s-notes-2")
	}
	page := browser.PageHTML(r.Screen().Render(), "")
	h := browser.NewHarness(t)
	const look = `(() => { const tree = document.querySelector(".tc-tree--scroll"), on = tree.querySelector(".is-on");
		const t = tree.getBoundingClientRect(), o = on.getBoundingClientRect();
		const rail = document.querySelector(".tc-shell__rail");
		return [Math.round(tree.scrollTop), Math.round(o.top - t.top), Math.round(t.bottom - o.bottom),
			Math.round((o.top + o.bottom) / 2 - (t.top + t.bottom) / 2), Math.round(rail.scrollTop), Math.round(window.scrollY),
			Math.round(tree.scrollHeight - tree.clientHeight), Math.round(o.height)]; })()`
	for _, c := range []struct {
		row    string
		centre bool // the row can stand in the middle: not near the end
	}{
		{"Заметка со встречи", false},
		{"Как устроить резервные копии", true},
		{"Пароли Wi-Fi у родителей", false},
	} {
		html := strings.Replace(page, `<div class="is-on tc-tree__row">`, `<div class="tc-tree__row">`, 1)
		mark := `<div class="tc-tree__row"><svg><use href="#i-note"/></svg>` + c.row + `</div>`
		if !strings.Contains(html, mark) {
			t.Fatalf("no row %q in s-notes-2", c.row)
		}
		html = strings.Replace(html, mark, strings.Replace(mark, `class="tc-tree__row"`, `class="is-on tc-tree__row"`, 1), 1)
		url := h.RegisterPage("neighbours-"+itoa(len(c.row)), html)
		var got []int
		if err := h.Run(url, 1351, 1000, chromedp.Evaluate(look, &got)); err != nil {
			t.Fatal(err)
		}
		scroll, above, below, off, rail, win, most := got[0], got[1], got[2], got[3], got[4], got[5], got[6]
		if got[7] != 26 {
			t.Errorf("%s: a row of Neighbours is %d px high; want 26, as in any tree — the block scrolls, the rows do not shrink", c.row, got[7])
		}
		if most <= 0 {
			t.Fatalf("Neighbours of 16 notes does not scroll (%v); want a block shorter than the notebook", got)
		}
		if above < 0 || below < 0 {
			t.Errorf("%s: the open note stands %d px from the block's top and %d from its bottom; want it inside", c.row, above, below)
		}
		if c.centre && (off < -2 || off > 2) {
			t.Errorf("%s: the open note is %d px off the block's middle; want it in the middle", c.row, off)
		}
		if c.row == "Заметка со встречи" && scroll != 0 {
			t.Errorf("the first note open: the block scrolled by %d; want it at its top", scroll)
		}
		if rail != 0 || win != 0 {
			t.Errorf("%s: the rail scrolled by %d and the page by %d; want only the block to scroll", c.row, rail, win)
		}
	}
}
