package browser_test

import (
	"strings"
	"testing"
	"time"

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
	top := w.Top{Brand: w.Brand{Title: "Контакты"}, Actions: []w.TopAction{w.TopDrop{Trigger: w.UserMenu{Name: "mix"},
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

// TestFailedMenuStandsWhole (Carrel's task 276 p. 3): the run bar's «N
// failed» opens its menu of refusals at its own width (420) where the
// window has room, and no wider than the window where it has not — not
// squeezed to its button's width with the server's answer broken letter by
// letter.
func TestFailedMenuStandsWhole(t *testing.T) {
	needChrome(t)
	bar := w.Runbar{Result: w.RunFailed, Parts: []w.RunbarPart{w.Say{Full: "Polled · 5 of 6 answered · 7 events ·", Short: "5 of 6 · 7 events ·"},
		w.Drop{Trigger: w.Button{Label: "1 failed", Menu: true, Kind: w.ButtonDanger, Small: true}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuRefusal{Mark: w.Mark{Colour: "#9A9280", Size: w.MarkH11}, Source: "baikal · Общий", Answer: "PROPFIND /dav.php/calendars/mix/ → 401 Unauthorized: the saved password was refused. Open Connections to update it.", Retry: w.Button{Label: "Retry", Small: true}}},
			Width: w.MenuWidth420, Pop: w.MenuUnderEnd}},
		w.Mono{Text: "2 min ago", Place: w.Pushed}}}
	shell := w.Shell{Top: &w.Top{Brand: w.Brand{Title: "Calendar"}}, Body: w.ShellBody{
		Rail: []w.RailPart{w.Nav{Items: []w.NavItem{{Icon: "cal", Label: "Calendar", Href: "#"}}, On: 1}},
		Main: w.Main{Items: []w.MainPart{w.Head{Title: &w.HeadTitle{Text: "Combined calendar"}}, bar}}}}
	h := browser.NewHarness(t)
	url := h.RegisterPage("failed-menu", browser.PageHTML(shell.Render(), ""))
	const measure = `(() => { const m = document.querySelector(".tc-runbar .tc-menu"), r = m.getBoundingClientRect(), root = document.querySelector(".tc-root").getBoundingClientRect();
		return [Math.round(r.width), Math.round(r.left - root.left), Math.round(r.right - root.left), Math.round(root.width)]; })()`
	for _, width := range []int{860, 1351, 390} {
		var got []int
		if err := h.Run(url, width, 800, chromedp.Click(`.tc-runbar .tc-btn.is-danger`, chromedp.ByQuery), chromedp.Evaluate(measure, &got)); err != nil {
			t.Fatal(err)
		}
		want := min(420, got[3]-24)
		if got[0] < want || got[1] < 0 || got[2] > got[3] {
			t.Errorf("at %d px the menu is %d wide at %d…%d of a %d px page (from its left edge); want at least %d and inside the page", width, got[0], got[1], got[2], got[3], want)
		}
	}
}

// TestTopMenuOverTheList (Carrel's task 276 p. 4): a menu of the top bar
// opened over the list stands over everything the list draws — the bar's
// «Columns» did not stay on top of «New»'s items.
func TestTopMenuOverTheList(t *testing.T) {
	needChrome(t)
	newDrop := w.TopDrop{Trigger: w.Button{Label: "New", Icon: "plus", Kind: w.ButtonPrimary, Menu: true}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "note", Label: "New note"}, w.MenuItem{Icon: "cal", Label: "New event"}, w.MenuItem{Icon: "task", Label: "New task"},
		w.MenuItem{Icon: "con", Label: "New contact"}, w.MenuItem{Icon: "upload", Label: "Import"}}}}
	bar := w.Bar{Items: []w.BarPart{w.Filter{Query: w.Query{Placeholder: "Filter by word", Label: "Filter by word"}, Variant: w.FilterH26W200},
		w.Button{Label: "Columns", Icon: "cols"}, w.Button{Label: "Print", Icon: "print"}, w.Button{Label: "Export", Icon: "export"}}}
	// The head's «New note» with its menu, and the bar under it with
	// «Columns» as a drop of its own, as Carrel's list pages print them.
	headNew := w.Drop{Trigger: w.Button{Label: "New note", Icon: "plus", Kind: w.ButtonPrimary, Menu: true}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "note", Label: "New note"}, w.MenuItem{Icon: "upload", Label: "Import notes"}, w.MenuItem{Icon: "export", Label: "Export"}, w.MenuItem{Icon: "print", Label: "Print"}},
		Pop: w.MenuUnderEnd}}
	bar.Items = append(bar.Items, w.Drop{Trigger: w.Button{Label: "Columns", Icon: "cols", Menu: true}, Menu: w.Menu{Items: []w.MenuPart{w.MenuItem{Label: "Title"}}, Pop: w.MenuUnder}})
	shell := w.Shell{Top: &w.Top{Brand: w.Brand{Title: "Notes"}, Actions: []w.TopAction{newDrop, w.UserMenu{Name: "mix"}}},
		Body: w.ShellBody{Rail: []w.RailPart{w.Nav{Items: []w.NavItem{{Icon: "note", Label: "Notes", Href: "#"}}, On: 1}},
			Main: w.Main{Items: []w.MainPart{w.Head{Title: &w.HeadTitle{Text: "Combined notebook"}, Acts: &w.Acts{Items: []w.ActsPart{headNew}}}, bar}}}}
	h := browser.NewHarness(t)
	url := h.RegisterPage("top-menu", browser.PageHTML(shell.Render(), ""))
	const covered = `(() => { const out = [];
		document.querySelectorAll(".tc-top .tc-menu [role=menuitem]").forEach(it => { const r = it.getBoundingClientRect();
			for (const x of [r.left + 4, r.left + r.width / 2, r.right - 4]) { const e = document.elementFromPoint(x, r.top + r.height / 2);
				if (!e || !e.closest(".tc-top .tc-menu")) out.push(it.textContent.trim() + " under " + (e ? (e.closest(".tc-btn") || e).textContent.trim() : "nothing")); } });
		return out; })()`
	headCovered := strings.ReplaceAll(covered, ".tc-top .tc-menu", ".tc-head .tc-menu")
	for _, width := range []int{390, 630, 900, 1351} {
		var hidden []string
		if err := h.Run(url, width, 700, chromedp.Click(`.tc-top .tc-btn.is-primary`, chromedp.ByQuery), chromedp.Evaluate(covered, &hidden)); err != nil {
			t.Fatal(err)
		}
		for _, x := range hidden {
			t.Errorf("at %d px, the top bar's New: %s", width, x)
		}
		hidden = nil
		if err := h.Run(url, width, 700, chromedp.Click(`.tc-head .tc-btn.is-primary`, chromedp.ByQuery), chromedp.Evaluate(headCovered, &hidden)); err != nil {
			t.Fatal(err)
		}
		for _, x := range hidden {
			t.Errorf("at %d px, the head's New note: %s", width, x)
		}
	}
}

// TestTopSplitOverTheHeadSplit (Carrel's task 283): the menu of the top
// bar's split «New note» stands over the list head's split «New contact» —
// both splits with a menu stand a layer up, and the head's came after.
func TestTopSplitOverTheHeadSplit(t *testing.T) {
	needChrome(t)
	split := func(label, href string) w.Split {
		return w.Split{Main: w.Button{Label: label, Icon: "plus", Kind: w.ButtonPrimary, Href: href},
			More: w.SplitDrop{Arrow: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				Menu: w.Menu{Items: []w.MenuPart{w.MenuItem{Icon: "con", Label: "Contact"}, w.MenuItem{Icon: "cal", Label: "Event"},
					w.MenuItem{Icon: "task", Label: "Task"}, w.MenuItem{Icon: "upload", Label: "Import"}, w.MenuItem{Icon: "export", Label: "Export"}}}}}
	}
	shell := w.Shell{Top: &w.Top{Brand: w.Brand{Title: "Contacts"}, Actions: []w.TopAction{split("New note", "/app/notes/quick"), w.UserMenu{Name: "mix"}}},
		Body: w.ShellBody{Rail: []w.RailPart{w.Nav{Items: []w.NavItem{{Icon: "con", Label: "Contacts", Href: "#"}}, On: 1}},
			Main: w.Main{Items: []w.MainPart{w.Head{Title: &w.HeadTitle{Text: "All contacts"}, Acts: &w.Acts{Items: []w.ActsPart{split("New contact", "/app/contacts/new")}}}}}}}
	h := browser.NewHarness(t)
	url := h.RegisterPage("top-split", browser.PageHTML(shell.Render(), ""))
	const covered = `(() => { const out = [];
		document.querySelectorAll(".tc-top .tc-menu [role=menuitem]").forEach(it => { const r = it.getBoundingClientRect();
			for (const x of [r.left + 4, r.left + r.width / 2, r.right - 4]) { const e = document.elementFromPoint(x, r.top + r.height / 2);
				if (!e || !e.closest(".tc-top .tc-menu")) out.push(it.textContent.trim() + " under " + (e ? (e.closest(".tc-btn") || e).textContent.trim() : "nothing")); } });
		return out; })()`
	for _, width := range []int{630, 900, 1351} {
		var hidden []string
		if err := h.Run(url, width, 700, chromedp.Click(`.tc-top .tc-split .tc-drop__anchor .tc-btn`, chromedp.ByQuery), chromedp.Evaluate(covered, &hidden)); err != nil {
			t.Fatal(err)
		}
		for _, x := range hidden {
			t.Errorf("at %d px, the top bar's split: %s", width, x)
		}
	}
}

// TestSplitArrowOpensTheRest (Carrel's task 276 p. 6): «New note» is a
// split — its main half makes a note and opens no menu; the arrow opens
// the menu of the rest, without a script too, and Esc closes it.
func TestSplitArrowOpensTheRest(t *testing.T) {
	needChrome(t)
	split := w.Split{Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary, Href: "/app/notes/quick"},
		More: w.SplitDrop{Arrow: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
			Menu: w.Menu{Items: []w.MenuPart{w.MenuItem{Icon: "con", Label: "Contact"}, w.MenuItem{Icon: "cal", Label: "Event"}, w.MenuItem{Icon: "task", Label: "Task"}}}}}
	top := w.Top{Brand: w.Brand{Title: "Notes"}, Actions: []w.TopAction{split, w.UserMenu{Name: "mix"}}}
	body := `<div class="tc-root tc-shell">` + top.Render() + `</div>`
	h := browser.NewHarness(t)
	shown := `(() => { const m = document.querySelector(".tc-split .tc-menu"); return !!m && m.checkVisibility({visibilityProperty: true}); })()`
	bare := h.RegisterPage("split-bare", strings.Replace(browser.PageHTML(body, ""), `<script src="/tc.js"></script>`, "", 1))
	var before, after bool
	if err := h.Run(bare, 1200, 600, chromedp.Evaluate(shown, &before), chromedp.Click(`.tc-split .tc-drop__anchor .tc-btn`, chromedp.ByQuery), chromedp.Evaluate(shown, &after)); err != nil {
		t.Fatal(err)
	}
	if before || !after {
		t.Errorf("without a script: shown at start %v, after the arrow %v; want false, true", before, after)
	}
	url := h.RegisterPage("split", browser.PageHTML(body, ""))
	var mainOpens bool
	var href string
	if err := h.Run(url, 1200, 600,
		chromedp.Evaluate(`document.querySelector(".tc-split > a.tc-btn").getAttribute("href")`, &href),
		chromedp.Evaluate(`(() => { document.querySelector(".tc-split > a.tc-btn").addEventListener("click", e => e.preventDefault()); return true; })()`, nil),
		chromedp.Click(`.tc-split > a.tc-btn`, chromedp.ByQuery), chromedp.Evaluate(shown, &mainOpens)); err != nil {
		t.Fatal(err)
	}
	if href != "/app/notes/quick" || mainOpens {
		t.Errorf("the main half: goes to %q, opens the menu %v; want /app/notes/quick, false", href, mainOpens)
	}
	var opened, closed, back bool
	if err := h.Run(url, 1200, 600,
		chromedp.Click(`.tc-split .tc-drop__anchor .tc-btn`, chromedp.ByQuery), chromedp.Evaluate(shown, &opened),
		chromedp.KeyEvent("\x1b"), chromedp.Evaluate(`!(`+shown+`)`, &closed),
		chromedp.Evaluate(`document.activeElement === document.querySelector(".tc-split .tc-drop__anchor .tc-btn")`, &back)); err != nil {
		t.Fatal(err)
	}
	if !opened || !closed || !back {
		t.Errorf("with the script: the arrow opened %v, Esc closed %v, focus back on the arrow %v", opened, closed, back)
	}
}

// TestRailOpensAsPanel (Carrel's tasks 274, 275): on a narrow width the rail
// opens as a panel. Settings show their sections as icons in the rail of
// icons, and on a phone «← App», the current section and «Sections» at the
// bottom; «Sections» or the current section opens the rail with all of them,
// Esc closes it. A list's current section opens its sources so, in both
// narrow modes; on a wide page it stays a link.
func TestRailOpensAsPanel(t *testing.T) {
	needChrome(t)
	sections := []w.NavItem{{Icon: "user", Label: "Account", Href: "#a"}, {Icon: "server", Label: "Connections", Href: "#c"}, {Icon: "file", Label: "Attachments", Href: "#t"},
		{Icon: "device", Label: "New device", Href: "#d"}, {Icon: "archive", Label: "Backup", Href: "#b"}, {Icon: "link", Label: "Published links", Href: "#p"}, {Icon: "theme", Label: "Appearance", Href: "#e"}}
	settings := w.Shell{Top: &w.Top{Brand: w.Brand{Title: "Settings"}}, Body: w.ShellBody{
		Rail: []w.RailPart{w.Nav{Items: []w.NavItem{{Icon: "back", Label: "Back to app", Href: "#app"}}, Sections: true},
			w.RailSec{Items: []w.RailSecPart{w.Rubric{Text: "Settings"}, w.Nav{Items: sections, On: 2}}}},
		Main: w.Main{Items: []w.MainPart{w.Head{Title: &w.HeadTitle{Text: "Connections"}}}}}}
	list := w.Shell{Top: &w.Top{Brand: w.Brand{Title: "Calendar"}}, Body: w.ShellBody{
		Rail: []w.RailPart{w.Nav{Items: []w.NavItem{{Icon: "con", Label: "Contacts", Href: "#con"}, {Icon: "cal", Label: "Calendar", Href: "#cal"}, {Icon: "task", Label: "Tasks", Href: "#task"}}, On: 2},
			w.RailSec{Items: []w.RailSecPart{w.Rubric{Text: "nextcloud"}, w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "5"}, w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "2"}}}},
		Main: w.Main{Items: []w.MainPart{w.Head{Title: &w.HeadTitle{Text: "Combined calendar"}}}}}}
	h := browser.NewHarness(t)
	surl := h.RegisterPage("rail-settings", browser.PageHTML(settings.Render(), ""))
	lurl := h.RegisterPage("rail-list", browser.PageHTML(list.Render(), ""))
	// what the rail shows, in the order it stands on the screen: top to bottom, then left to right
	const seen = `(() => [...document.querySelectorAll(".tc-shell__rail .tc-nav__item, .tc-shell__rail .tc-tickrow")].filter(e => e.getClientRects().length && e.checkVisibility())
		.map(e => ({r: e.getBoundingClientRect(), t: (e.querySelector(".tc-tickrow__label") || e).textContent.trim()})).sort((a, b) => Math.round(a.r.top) - Math.round(b.r.top) || a.r.left - b.r.left).map(e => e.t).join("|"))()`
	const state = `(() => [document.querySelector(".tc-shell__body").hasAttribute("data-tc-rail-open"), !!document.querySelector(".tc-main").getClientRects().length].join(" "))()`
	var icons, bottom, open, closed string
	var stOpen, stClosed string
	if err := h.Run(surl, 820, 700, chromedp.Evaluate(seen, &icons)); err != nil {
		t.Fatal(err)
	}
	if icons != "Back to app|Account|Connections|Attachments|New device|Backup|Published links|Appearance" {
		t.Errorf("settings in the rail of icons show %q; want back and the seven sections", icons)
	}
	if err := h.Run(surl, 390, 700, chromedp.Evaluate(seen, &bottom),
		chromedp.Click(`.tc-nav__all`, chromedp.ByQuery), chromedp.Evaluate(seen, &open), chromedp.Evaluate(state, &stOpen),
		chromedp.KeyEvent("\x1b"), chromedp.Evaluate(seen, &closed), chromedp.Evaluate(state, &stClosed)); err != nil {
		t.Fatal(err)
	}
	if bottom != "Back to app|Connections|Sections" {
		t.Errorf("settings on a phone show %q at the bottom; want back, the current section and Sections", bottom)
	}
	if !strings.Contains(open, "Account|Connections|Attachments|New device|Backup|Published links|Appearance") || stOpen != "true false" {
		t.Errorf("Sections pressed: %q, state %s; want every section, the rail open over the page", open, stOpen)
	}
	if closed != bottom || stClosed != "false true" {
		t.Errorf("after Esc: %q, state %s; want the bottom bar again and the page", closed, stClosed)
	}
	for _, width := range []int{820, 390} {
		var before, after, again string
		if err := h.Run(lurl, width, 700, chromedp.Evaluate(seen, &before),
			chromedp.Click(`.tc-shell__rail .tc-nav__item[aria-current="page"]`, chromedp.ByQuery), chromedp.Evaluate(seen, &after),
			chromedp.Click(`.tc-shell__rail .tc-nav__item[aria-current="page"]`, chromedp.ByQuery), chromedp.Evaluate(seen, &again)); err != nil {
			t.Fatal(err)
		}
		if strings.Contains(before, "Личный") || !strings.Contains(after, "Личный|Работа") || again != before {
			t.Errorf("at %d px a list's current section: before %q, pressed %q, pressed again %q; want the sources only while open", width, before, after, again)
		}
	}
	var wide, hash string
	if err := h.Run(lurl, 1351, 700, chromedp.Click(`.tc-shell__rail .tc-nav__item[aria-current="page"]`, chromedp.ByQuery),
		chromedp.Evaluate(state, &wide), chromedp.Evaluate(`location.hash`, &hash)); err != nil {
		t.Fatal(err)
	}
	if wide != "false true" || hash != "#cal" {
		t.Errorf("on a wide page the current section: state %s, address %q; want a plain link", wide, hash)
	}
}

// TestRowOpensItsPage (Carrel's task 290): a row that opens its record in the
// panel fills the panel where the width draws one, and goes to the record's
// own page where it does not — on a tablet the note opened nothing.
func TestRowOpensItsPage(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	panel := w.Side{Zone: "panel", Items: []w.SidePart{w.Head{Title: &w.HeadTitle{Text: "Заметка"}}}}
	part := h.RegisterPage("paged-part", panel.Render())
	own := h.RegisterPage("paged-own", browser.PageHTML(`<div class="tc-root"><p>the note's page</p></div>`, ""))
	list := w.List{Target: "panel", Rows: []w.ListRow{w.ExcerptRow{Title: []w.Inline{w.Text("Заметка")}, Href: part, Page: own}}}
	shell := w.Shell{Body: w.ShellBody{Main: w.Main{Zone: "main", Items: []w.MainPart{list}}, Side: &w.Side{Zone: "panel", Closed: true}}}
	url := h.RegisterPage("paged", browser.PageHTML(shell.Render(), ""))
	const where = `location.pathname.split("/").pop() + " " + (document.querySelector(".tc-side") && document.querySelector(".tc-side").getClientRects().length ? "panel" : "-")`
	var wide, narrow string
	if err := h.Run(url, 1500, 700, chromedp.Click(`.tc-row`, chromedp.ByQuery), chromedp.Sleep(1200*time.Millisecond), chromedp.Evaluate(where, &wide)); err != nil {
		t.Fatal(err)
	}
	if err := h.Run(url, 900, 700, chromedp.Click(`.tc-row`, chromedp.ByQuery), chromedp.Sleep(1200*time.Millisecond), chromedp.Evaluate(where, &narrow)); err != nil {
		t.Fatal(err)
	}
	if wide != "paged panel" {
		t.Errorf("on a wide page the row: %q; want the panel filled on the same page", wide)
	}
	if narrow != "paged-own -" {
		t.Errorf("without a panel the row: %q; want the record's own page", narrow)
	}
}
