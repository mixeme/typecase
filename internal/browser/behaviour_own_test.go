package browser_test

import (
	"strings"
	"testing"

	"carrel.dev/typecase/internal/browser"
	w "carrel.dev/typecase/widget"
	"github.com/chromedp/cdproto/input"
	"github.com/chromedp/chromedp"
)

// The three widgets whose behaviour had been held only by a rule checked on
// another widget (docs/behaviour.md, 20.11.2026): each has its own test now.

// TestDangerAsks: the danger zone's button asks before it deletes — it opens
// the confirming sheet it names, the focus goes into the sheet, and Esc
// closes it and gives the focus back to the button.
func TestDangerAsks(t *testing.T) {
	needChrome(t)
	zone := w.Danger{Button: w.Button{Label: "Delete contact", Icon: "trash", Kind: w.ButtonDangerSolid, Opens: "confirm"}}
	sheet := w.Scrim{ID: "confirm", Closed: true, Content: w.Sheet{Title: "Delete the contact?",
		Body: []w.SheetPart{w.Dim{Parts: []w.Inline{w.Text("Анна Ковалёва goes from Контакты on the server.")}}},
		Foot: []w.SheetFootPart{w.Button{Label: "Delete", Kind: w.ButtonDangerSolid, Submit: true}, w.Button{Label: "Cancel", Closes: true}}}}
	body := `<div class="tc-root"><div class="tc-main">` + zone.Render() + sheet.Render() + `</div></div>`
	h := browser.NewHarness(t)
	url := h.RegisterPage("danger", browser.PageHTML(body, ""))
	shown := `(() => { const s = document.getElementById("confirm"); return !s.hidden && s.getClientRects().length > 0; })()`
	var before, opened, inside, closed, back bool
	err := h.Run(url, 800, 600,
		chromedp.Evaluate(shown, &before),
		chromedp.Click(`.tc-danger .tc-btn`, chromedp.ByQuery),
		chromedp.Evaluate(shown, &opened),
		chromedp.Evaluate(`document.getElementById("confirm").contains(document.activeElement)`, &inside),
		chromedp.KeyEvent("\x1b"),
		chromedp.Evaluate(`!(`+shown+`)`, &closed),
		chromedp.Evaluate(`document.activeElement === document.querySelector(".tc-danger .tc-btn")`, &back),
	)
	if err != nil {
		t.Fatal(err)
	}
	if before || !opened || !inside || !closed || !back {
		t.Errorf("closed at start %v, asked %v, focus in the sheet %v, closed by Esc %v, focus back on the button %v", !before, opened, inside, closed, back)
	}
}

// TestOverlayKeepsFocus: a page drawn with its sheet open over the shell (a
// sheet's own address, Р-16) keeps the focus in the sheet — Tab from the
// page's start goes into the sheet, not under the scrim, and Tab and
// Shift+Tab go round inside it; a click under the scrim reaches nothing.
func TestOverlayKeepsFocus(t *testing.T) {
	needChrome(t)
	o := w.Overlay{
		Top: &w.Top{Brand: w.Brand{Title: "Контакты"}, Actions: []w.TopAction{w.Button{Label: "New", Icon: "plus"}}},
		Body: w.OverlayBody{Main: w.Main{Items: []w.MainPart{
			w.Head{Title: &w.HeadTitle{Text: "Контакты"}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Import", Href: "/import"}}}},
		}}},
		Scrim: w.Scrim{Content: w.Sheet{Title: "Rename", Action: "/rename",
			Body: []w.SheetPart{w.Field{Label: "Name", Parts: []w.FieldPart{w.Input{Name: "name", Value: "Контакты"}}}},
			Foot: []w.SheetFootPart{w.Button{Label: "Save", Kind: w.ButtonPrimary, Submit: true}, w.Button{Label: "Cancel", Href: "/back"}}}},
	}
	body := `<div class="tc-root tc-shell">` + o.Render() + `</div>`
	h := browser.NewHarness(t)
	url := h.RegisterPage("overlay", browser.PageHTML(body, ""))
	where := `(() => { const a = document.activeElement; return a && a.closest(".tc-sheet") ? "sheet" : a && a.closest(".tc-overlay") ? "under" : "page"; })()`
	var first string
	steps := make([]string, 6)
	actions := []chromedp.Action{chromedp.KeyEvent("\t"), chromedp.Evaluate(where, &first)}
	for i := range steps {
		actions = append(actions, chromedp.KeyEvent("\t"), chromedp.Evaluate(where, &steps[i]))
	}
	var back string
	actions = append(actions,
		chromedp.Evaluate(`document.querySelector(".tc-sheet input").focus(); true`, nil),
		chromedp.KeyEvent("\t", chromedp.KeyModifiers(input.ModifierShift)),
		chromedp.Evaluate(where, &back),
	)
	var hit string
	actions = append(actions, chromedp.Evaluate(`(() => { const b = document.querySelector(".tc-head .tc-btn").getBoundingClientRect();
		const e = document.elementFromPoint(b.left + b.width / 2, b.top + b.height / 2); return e && e.closest(".tc-head") ? "under" : "scrim"; })()`, &hit))
	if err := h.Run(url, 900, 600, actions...); err != nil {
		t.Fatal(err)
	}
	if first != "sheet" {
		t.Errorf("the first Tab took the focus to %s; want the sheet", first)
	}
	for i, s := range steps {
		if s != "sheet" {
			t.Errorf("Tab %d took the focus to %s; want it kept in the sheet", i+2, s)
		}
	}
	if back != "sheet" {
		t.Errorf("Shift+Tab from the sheet's first field went to %s; want the sheet", back)
	}
	if hit != "scrim" {
		t.Errorf("a click on a button under the scrim reaches %s; want the scrim", hit)
	}
}

// TestUserMenuOpens: the account button opens the account menu under it —
// without a script too, as a details; with the library's script Esc closes
// it and gives the focus back to the button.
func TestUserMenuOpens(t *testing.T) {
	needChrome(t)
	top := w.Top{Brand: w.Brand{Title: "Контакты"}, Actions: []w.TopAction{w.TopDrop{Trigger: w.UserMenu{Name: "mix"}, End: true,
		Menu: w.Menu{Items: []w.MenuPart{w.MenuItem{Label: "Account", Href: "/account"}, w.MenuItem{Label: "Sign out", Href: "/sign-out"}}}}}}
	body := `<div class="tc-root tc-shell">` + top.Render() + `</div>`
	h := browser.NewHarness(t)
	shown := `(() => { const m = document.querySelector(".tc-top .tc-menu"); return !!m && m.checkVisibility({visibilityProperty: true}); })()`
	bare := h.RegisterPage("usermenu-bare", strings.Replace(browser.PageHTML(body, ""), `<script src="/tc.js"></script>`, "", 1))
	var bareBefore, bareAfter bool
	if err := h.Run(bare, 1200, 600,
		chromedp.Evaluate(shown, &bareBefore),
		chromedp.Click(`.tc-usermenu`, chromedp.ByQuery),
		chromedp.Evaluate(shown, &bareAfter),
	); err != nil {
		t.Fatal(err)
	}
	if bareBefore || !bareAfter {
		t.Errorf("without a script: shown at start %v, after a press %v; want false, true", bareBefore, bareAfter)
	}
	url := h.RegisterPage("usermenu", browser.PageHTML(body, ""))
	var opened, closed, back bool
	if err := h.Run(url, 1200, 600,
		chromedp.Click(`.tc-usermenu`, chromedp.ByQuery),
		chromedp.Evaluate(shown, &opened),
		chromedp.KeyEvent("\x1b"),
		chromedp.Evaluate(`!(`+shown+`)`, &closed),
		chromedp.Evaluate(`document.activeElement === document.querySelector(".tc-usermenu")`, &back),
	); err != nil {
		t.Fatal(err)
	}
	if !opened || !closed || !back {
		t.Errorf("with the script: opened %v, closed by Esc %v, focus back on the button %v", opened, closed, back)
	}
}

// TestOverlayCloseFreesThePage (Carrel's task 276 p. 1): on a sheet's own
// page the cross and Esc close the sheet in place, and what lay under it is
// no longer inert — the head's button takes a press again — and the
// address becomes the page under the sheet, the cross's link.
func TestOverlayCloseFreesThePage(t *testing.T) {
	needChrome(t)
	o := w.Overlay{
		Top: &w.Top{Brand: w.Brand{Title: "Notes"}},
		Body: w.OverlayBody{Main: w.Main{Items: []w.MainPart{
			w.Head{Title: &w.HeadTitle{Text: "Notes"}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Import", Href: "/import"}}}},
		}}},
		Scrim: w.Scrim{Content: w.Sheet{Title: "New note", Head: []w.SheetHeadPart{w.Close{Label: "Close", Href: "/app/notes"}},
			Body: []w.SheetPart{w.Area{Label: "Text"}},
			Foot: []w.SheetFootPart{w.Button{Label: "Save", Kind: w.ButtonPrimary}, w.Button{Label: "Cancel", Href: "/app/notes", Closes: true}}}},
	}
	body := `<div class="tc-root tc-shell">` + o.Render() + `</div>`
	h := browser.NewHarness(t)
	url := h.RegisterPage("sheet-page", browser.PageHTML(body, ""))
	const state = `(() => { const s = document.querySelector(".tc-scrim"), b = document.querySelector(".tc-head .tc-btn").getBoundingClientRect();
		const e = document.elementFromPoint(b.left + b.width / 2, b.top + b.height / 2);
		return [s.getClientRects().length ? "open" : "closed", document.querySelector("[inert]") ? "inert" : "free", e && e.closest(".tc-head") ? "reaches" : "blocked", location.pathname].join(" "); })()`
	for _, how := range []struct {
		name   string
		action chromedp.Action
	}{
		{"the cross", chromedp.Click(`.tc-sheet .tc-close`, chromedp.ByQuery)},
		{"Esc", chromedp.KeyEvent("\x1b")},
	} {
		var before, after string
		if err := h.Run(url, 1200, 700, chromedp.Evaluate(state, &before), how.action, chromedp.Evaluate(state, &after)); err != nil {
			t.Fatal(err)
		}
		if !strings.HasPrefix(before, "open inert blocked") {
			t.Errorf("before %s: %s; want the sheet open over an inert page", how.name, before)
		}
		if after != "closed free reaches /app/notes" {
			t.Errorf("after %s: %s; want closed free reaches /app/notes", how.name, after)
		}
	}
}
