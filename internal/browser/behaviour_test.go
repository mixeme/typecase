package browser_test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/chromedp/chromedp"

	"carrel.dev/typecase/internal/browser"
	"carrel.dev/typecase/internal/refs"
	w "carrel.dev/typecase/widget"
)

func needChrome(t *testing.T) {
	t.Helper()
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if os.Getenv("CI") == "" && testing.Short() {
		t.Skip("needs Chrome; skip with -short outside CI")
	}
}

// flush lets a fixture's root take the whole width of the window, as a
// screen's root does: the container queries of the mockup then see the ladder
// width itself.
const flush = `body { padding: 0 !important; }`

// wantMode is the shell's state for a root width, as docs/visual/mockup.md
// (Р-3) states it.
func wantMode(width int) string {
	switch {
	case width <= 640:
		return "drawer"
	case width <= 1007:
		return "rail"
	}
	return "full"
}

// TestShellModeFromCSS: tc.js has no threshold of its own; the state it
// writes on the shell's root is the one the mockup's CSS sets, on every step
// of the ladder.
func TestShellModeFromCSS(t *testing.T) {
	needChrome(t)
	var shell string
	for _, e := range refs.Entries {
		if s, ok := e.Fixture.(w.Shell); ok && e.State == "side" {
			shell = s.Render()
			break
		}
	}
	if shell == "" {
		t.Fatal("no shell record with a side panel")
	}
	h := browser.NewHarness(t)
	url := h.RegisterPage("shell-mode", browser.PageHTML(shell, flush))
	for _, width := range browser.MockupLadderWidths {
		var mode string
		err := h.Run(url, width, 800, chromedp.Evaluate(`(() => { window.__tc.layoutUpdates(); return document.querySelector(".tc-shell").getAttribute("data-tc-mode"); })()`, &mode))
		if err != nil {
			t.Fatal(err)
		}
		if mode != wantMode(width) {
			t.Errorf("%d px: data-tc-mode %q, want %q", width, mode, wantMode(width))
		}
	}
}

// TestBarSafetyQuiet (Р-4): what goes under «⋯» and when is the mockup's
// CSS; tc.js only catches a bar the CSS did not fit. On the reference bars,
// edit bars and heads it never has to, on any step of the ladder: the
// narrowest steps give the tertiary items (is-3rd) to «⋯» (the owner's
// decision of 2026-09-26: no sideways scroll).
func TestBarSafetyQuiet(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	for i, e := range refs.Entries {
		switch e.Fixture.(type) {
		case w.Bar, w.Docbar, w.Head:
		default:
			continue
		}
		body := `<div class="tc-root">` + e.Fixture.Render() + `</div>`
		url := h.RegisterPage(fmt.Sprintf("bar-%d", i), browser.PageHTML(body, flush))
		for _, width := range browser.MockupLadderWidths {
			var n int
			// The frames draw drops open; a screen shows them closed.
			js := `(() => { document.querySelectorAll(".tc-drop > .tc-menu, .tc-drop__anchor > :last-child").forEach((m) => { m.hidden = true; m.style.display = "none"; }); return window.__tc.layoutUpdates(); })()`
			if err := h.Run(url, width, 600, chromedp.Evaluate(js, &n)); err != nil {
				t.Fatal(err)
			}
			if n != 0 {
				t.Errorf("%s at %d px: the CSS did not fit it and the safety tripped", e.Ref, width)
			}
		}
	}
}

// TestOpensInPlace: a button with Opens shows the closed sheet it names; Esc
// closes it and gives the focus back to the button.
func TestOpensInPlace(t *testing.T) {
	needChrome(t)
	body := `<div class="tc-root">` + w.Button{Label: "New note", Opens: "quick"}.Render() +
		w.Scrim{ID: "quick", Closed: true, Content: w.Sheet{Title: "New note",
			Body: []w.SheetPart{w.Area{Text: "", Variant: w.AreaBodySm, Label: "Text"}},
			Foot: []w.SheetFootPart{w.Button{Label: "Save", Kind: w.ButtonPrimary}}}}.Render() + `</div>`
	h := browser.NewHarness(t)
	url := h.RegisterPage("opens", browser.PageHTML(body, ""))
	var before, opened, closed, back bool
	shown := `(() => { const s = document.getElementById("quick"); return !s.hidden && s.getClientRects().length > 0; })()`
	err := h.Run(url, 800, 600,
		chromedp.Evaluate(shown, &before),
		chromedp.Click(`[data-tc-open="quick"]`, chromedp.ByQuery),
		chromedp.Evaluate(shown, &opened),
		chromedp.Evaluate(`document.dispatchEvent(new KeyboardEvent("keydown", { key: "Escape", bubbles: true }))`, nil),
		chromedp.Evaluate(`(() => { const s = document.getElementById("quick"); return s.hidden && s.getClientRects().length === 0; })()`, &closed),
		chromedp.Evaluate(`document.activeElement === document.querySelector('[data-tc-open="quick"]')`, &back),
	)
	if err != nil {
		t.Fatal(err)
	}
	if before || !opened || !closed || !back {
		t.Errorf("closed at start %v, opened %v, closed by Esc %v, focus back %v", !before, opened, closed, back)
	}
}

// TestRailReopensSources (Р-3): in the rail of icons the current section,
// pressed again, opens the sheet with its sources; where the rail shows the
// sources itself, the current section stays a link.
func TestRailReopensSources(t *testing.T) {
	needChrome(t)
	body := w.Shell{Body: w.ShellBody{
		Rail: []w.RailPart{w.Nav{Items: []w.NavItem{{Icon: "con", Label: "Contacts", Href: "#"}, {Icon: "note", Label: "Notes", Href: "#"}}, On: 2, Opens: "sources"}},
		Main: w.Main{Items: []w.MainPart{
			w.Head{Title: &w.HeadTitle{Text: "Combined notebook"}},
			w.Scrim{ID: "sources", Closed: true, Content: w.Sheet{Title: "Sources",
				Body: []w.SheetPart{w.Dim{Parts: []w.Inline{w.Text("Записи")}}},
				Foot: []w.SheetFootPart{w.Button{Label: "Apply", Kind: w.ButtonPrimary}}}},
		}}}}.Render()
	h := browser.NewHarness(t)
	url := h.RegisterPage("rail-sources", browser.PageHTML(body, flush))
	shown := `(() => { const s = document.getElementById("sources"); return !s.hidden && s.getClientRects().length > 0; })()`
	for _, c := range []struct {
		width int
		want  bool
	}{{820, true}, {1351, false}} {
		var opened bool
		err := h.Run(url, c.width, 700,
			chromedp.Click(`.tc-nav__item[aria-current="page"]`, chromedp.ByQuery),
			chromedp.Evaluate(shown, &opened),
		)
		if err != nil {
			t.Fatal(err)
		}
		if opened != c.want {
			t.Errorf("%d px: the sources sheet opened %v, want %v", c.width, opened, c.want)
		}
	}
}

// TestDocTableLine: on a narrow screen a table in a text has Wrap / One line
// over it; One line keeps each row to one line and the table scrolls
// sideways, Wrap brings the words back. On a wide screen the switch is not
// there.
func TestDocTableLine(t *testing.T) {
	needChrome(t)
	row := [][]w.Inline{{w.Text("Билеты до Казани, туда и обратно")}, {w.Text("Виктор")}, {w.Text("до 10.08")}, {w.Text("11 200")}, {w.Text("поезд, не самолёт")}}
	doc := w.Doc{Items: []w.DocPart{w.DocTable{
		Head: [][]w.Inline{{w.Text("Что")}, {w.Text("Кто")}, {w.Text("Когда")}, {w.Text("Сколько")}, {w.Text("Заметка")}},
		Rows: [][][]w.Inline{row, row}, Switch: w.DocTableSwitch(false)}}}
	body := `<div class="tc-root"><div class="tc-main">` + doc.Render() + `</div></div>`
	h := browser.NewHarness(t)
	url := h.RegisterPage("doctable", browser.PageHTML(body, flush))
	state := `(() => { const t = document.querySelector(".tc-doc__table"), s = t.querySelector(".tc-doc__scroll");
		return [getComputedStyle(t.querySelector(".tc-doc__tablebar")).display !== "none", t.classList.contains("is-line"), s.scrollWidth > s.clientWidth]; })()`
	var wide, narrow, line, back []bool
	if err := h.Run(url, 1200, 700, chromedp.Evaluate(state, &wide)); err != nil {
		t.Fatal(err)
	}
	err := h.Run(url, 390, 700,
		chromedp.Evaluate(state, &narrow),
		chromedp.Click(`.tc-doc__tablebar .tc-seg__opt:nth-child(2)`, chromedp.ByQuery),
		chromedp.Evaluate(state, &line),
		chromedp.Click(`.tc-doc__tablebar .tc-seg__opt:nth-child(1)`, chromedp.ByQuery),
		chromedp.Evaluate(state, &back),
	)
	if err != nil {
		t.Fatal(err)
	}
	if wide[0] {
		t.Errorf("wide: the switch shows")
	}
	if !narrow[0] || narrow[1] || narrow[2] {
		t.Errorf("narrow: switch %v, one line %v, scrolls %v; want the switch, wrapped, no scroll", narrow[0], narrow[1], narrow[2])
	}
	if !line[1] || !line[2] {
		t.Errorf("One line: one line %v, scrolls %v", line[1], line[2])
	}
	if back[1] || back[2] {
		t.Errorf("Wrap again: one line %v, scrolls %v", back[1], back[2])
	}
}

// TestFormsWithoutScripts (task 29): a form of widgets with names works with
// no script at all — the check and the segment switch by the browser's own
// checkbox and radio button and draw their state, the form sends every named
// field and its hidden CSRF — and a button that is a link looks like the
// button next to it.
func TestFormsWithoutScripts(t *testing.T) {
	needChrome(t)
	form := w.Form{Action: "/app/settings/account", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}, Items: []w.FormPart{
		w.Field{Label: "Week starts on", Parts: []w.FieldPart{w.Pick{Name: "week", Label: "Week starts on", Options: []w.PickOption{{Value: "mon", Label: "Monday", On: true}, {Value: "sun", Label: "Sunday"}}}}},
		w.Field{Label: "Theme", Parts: []w.FieldPart{w.Seg{Name: "theme", Options: []w.SegOption{{Label: "Auto", Value: "auto", On: true}, {Label: "Dark", Value: "dark"}}}}},
		w.Field{Parts: []w.FieldPart{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}}},
	}}
	body := `<div class="tc-root">` + form.Render() +
		`<div id="pair">` + w.Button{Label: "Import"}.Render() + w.Button{Label: "Import", Href: "/import"}.Render() + `</div></div>`
	page := strings.Replace(browser.PageHTML(body, ""), `<script src="/tc.js"></script>`, "", 1)
	h := browser.NewHarness(t)
	url := h.RegisterPage("forms", page)
	state := `(() => {
		const tick = getComputedStyle(document.querySelector(".tc-check .tc-tickrow__tick")).backgroundColor;
		const on = [...document.querySelectorAll(".tc-seg__opt")].map(o => getComputedStyle(o).fontWeight);
		const sent = [...new FormData(document.querySelector("form.tc-form")).entries()].map(e => e.join("=")).join("&");
		return [tick, on.join(","), sent];
	})()`
	look := `(() => {
		const [b, a] = document.querySelectorAll("#pair .tc-btn"), cb = getComputedStyle(b), ca = getComputedStyle(a);
		const rb = b.getBoundingClientRect(), ra = a.getBoundingClientRect();
		return ["height", "fontSize", "color", "backgroundColor", "borderTopWidth", "borderTopColor", "paddingLeft", "display"].filter(k => cb[k] !== ca[k]).concat(Math.abs(rb.width - ra.width) > 0.5 || Math.abs(rb.height - ra.height) > 0.5 ? ["box"] : []);
	})()`
	var before, after []string
	var differ []string
	err := h.Run(url, 900, 600,
		chromedp.Evaluate(state, &before),
		chromedp.Click(`.tc-check`, chromedp.ByQuery),
		chromedp.Click(`.tc-seg__opt:nth-child(2)`, chromedp.ByQuery),
		chromedp.Evaluate(state, &after),
		chromedp.Evaluate(look, &differ),
	)
	if err != nil {
		t.Fatal(err)
	}
	if before[0] == after[0] {
		t.Errorf("the check's tick draws the same before and after a click: %s", before[0])
	}
	if before[1] != "600,400" || after[1] != "400,600" {
		t.Errorf("segment weights before %s, after %s; want 600,400 then 400,600", before[1], after[1])
	}
	if want := "csrf=7f3a91de&week=mon&theme=dark&compact=1"; after[2] != want {
		t.Errorf("the form sends %q, want %q", after[2], want)
	}
	if len(differ) > 0 {
		t.Errorf("the link button differs from the button in %v", differ)
	}
}

// TestNoScriptSheetAndChecks (task 29, step 4): a sheet opens by its address
// with no script; with the library's script a check switches its field off,
// the theme's switch keeps the choice, and a compact list has 30 px rows.
func TestNoScriptSheetAndChecks(t *testing.T) {
	needChrome(t)
	sheet := `<div class="tc-root"><div class="tc-bar">` + w.Button{Label: "Sources", Href: "#src"}.Render() + `</div>` +
		w.Scrim{ID: "src", Closed: true, Content: w.Sheet{Title: "Sources", Body: []w.SheetPart{w.Area{Label: "Text"}}}}.Render() + `</div>`
	h := browser.NewHarness(t)
	bare := h.RegisterPage("sheet", strings.Replace(browser.PageHTML(sheet, ""), `<script src="/tc.js"></script>`, "", 1))
	shown := `(() => { const s = document.getElementById("src"); return s.getClientRects().length > 0; })()`
	var closed, open bool
	if err := h.Run(bare, 900, 600, chromedp.Evaluate(shown, &closed)); err != nil {
		t.Fatal(err)
	}
	if err := h.Run(bare+"#src", 900, 600, chromedp.Evaluate(shown, &open)); err != nil {
		t.Fatal(err)
	}
	if closed || !open {
		t.Errorf("without a script: shown at start %v, shown at #src %v", closed, open)
	}

	form := w.Form{Action: "/x", Items: []w.FormPart{
		w.Field{Label: "Password", Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "password", Label: "Password"}}},
		w.Field{Parts: []w.FieldPart{w.Check{Name: "nopass", Value: "1", Disables: "password", Parts: []w.CheckPart{w.Text("The server asks no password")}}}},
		w.Field{Label: "Theme", Parts: []w.FieldPart{w.Seg{Name: "theme", Theme: true, Options: []w.SegOption{{Label: "Auto", Value: "auto", On: true}, {Label: "Dark", Value: "dark"}}}}},
	}}
	row := w.ExcerptRow{Mark: w.Mark{Colour: "#4A6B52"}, Title: []w.Inline{w.Text("Заметка")}, Date: "10 Aug", Source: "nextcloud"}
	list := w.List{Compact: true, Rows: []w.ListRow{row, row}}
	page := h.RegisterPage("checks", browser.PageHTML(`<div class="tc-root"><div class="tc-main">`+form.Render()+list.Render()+`</div></div>`, flush))
	var off0, off1, back bool
	var theme, stored string
	var height float64
	err := h.Run(page, 1200, 700,
		chromedp.Evaluate(`document.querySelector('[name=password]').disabled`, &off0),
		chromedp.Click(`.tc-check`, chromedp.ByQuery),
		chromedp.Evaluate(`document.querySelector('[name=password]').disabled`, &off1),
		chromedp.Click(`.tc-check`, chromedp.ByQuery),
		chromedp.Evaluate(`document.querySelector('[name=password]').disabled`, &back),
		chromedp.Click(`.tc-seg__opt:nth-child(2)`, chromedp.ByQuery),
		chromedp.Evaluate(`document.documentElement.getAttribute("data-tc-theme")`, &theme),
		chromedp.Evaluate(`localStorage.getItem("tc-theme") || ""`, &stored),
		chromedp.Evaluate(`document.querySelector(".tc-list.is-compact .tc-row").getBoundingClientRect().height`, &height),
	)
	if err != nil {
		t.Fatal(err)
	}
	if off0 || !off1 || back {
		t.Errorf("password off: at start %v, ticked %v, unticked %v; want false, true, false", off0, off1, back)
	}
	if theme != "dark" || stored != "dark" {
		t.Errorf("theme %q, stored %q; want dark, dark", theme, stored)
	}
	if height < 29.5 || height > 31.5 {
		t.Errorf("a compact row is %.1f px high, want 30", height)
	}
}

// TestAddRowAndDrop (task 29, step 4): «Add phone» adds its template's row
// with the next free number in its name; a file dropped on the page shows the
// drop sheet and sends the named form with the file.
func TestAddRowAndDrop(t *testing.T) {
	needChrome(t)
	num := func(name string) w.Field {
		return w.Field{Label: "Number", Parts: []w.FieldPart{w.Input{Name: name, Mono: true, Label: "Number"}}}
	}
	set := w.FieldSet{Rubric: w.Rubric{Text: "Phone numbers"}, Items: []w.FieldSetPart{
		num("tel.0"), w.Template{ID: "tpl-tel", Items: []w.FieldSetPart{num("tel.{n}")}},
		w.Button{Label: "Add phone", Small: true, Submit: true, Name: "add", Value: "tel", Adds: "tpl-tel"}}}
	form := w.Form{Action: "/x", Items: []w.FormPart{set}}
	post := w.PostForm{ID: "post-import", Action: "/import", File: "file"}
	drop := w.Scrim{ID: "drop", Closed: true, Drop: "post-import", Content: w.Sheet{Title: "Drop to import", Body: []w.SheetPart{w.Area{Label: "Text"}}}}
	body := `<div class="tc-root">` + form.Render() + post.Render() + drop.Render() + `</div>`
	h := browser.NewHarness(t)
	url := h.RegisterPage("addrow", browser.PageHTML(body, ""))
	var names string
	var shown, sent bool
	err := h.Run(url, 900, 600,
		chromedp.Click(`[data-tc-add]`, chromedp.ByQuery),
		chromedp.Click(`[data-tc-add]`, chromedp.ByQuery),
		chromedp.Evaluate(`[...document.querySelectorAll("form.tc-form input")].map(i => i.name).join(",")`, &names),
		chromedp.Evaluate(`(() => {
			window.__sent = false;
			HTMLFormElement.prototype.submit = function () { window.__sent = this.id === "post-import" && this.querySelector("input[type=file]").files.length === 1; };
			const dt = new DataTransfer();
			dt.items.add(new File(["BEGIN:VCARD"], "a.vcf"));
			document.dispatchEvent(new DragEvent("dragenter", {dataTransfer: dt, bubbles: true, cancelable: true}));
			return document.getElementById("drop").getClientRects().length > 0;
		})()`, &shown),
		chromedp.Evaluate(`(() => {
			const dt = new DataTransfer();
			dt.items.add(new File(["BEGIN:VCARD"], "a.vcf"));
			document.dispatchEvent(new DragEvent("drop", {dataTransfer: dt, bubbles: true, cancelable: true}));
			return window.__sent;
		})()`, &sent),
	)
	if err != nil {
		t.Fatal(err)
	}
	if names != "tel.0,tel.1,tel.2" {
		t.Errorf("rows after two «Add»: %s, want tel.0,tel.1,tel.2", names)
	}
	if !shown || !sent {
		t.Errorf("drop sheet shown while dragging %v, form sent with the file %v", shown, sent)
	}
}

// TestColumnsMenu (the choice of columns): the menu is a form whose ticked
// items are sent in their order; Alt+↓ moves an item down, the handle drags
// one, and a click on an item ticks it without closing anything.
func TestColumnsMenu(t *testing.T) {
	needChrome(t)
	item := func(label, value string, on bool) w.MenuItem {
		tick := w.TickBox{}
		if !on {
			tick.Ticked = w.TickedOff
		}
		return w.MenuItem{Icon: "drag", Tick: &tick, Label: label, Name: "col", Value: value}
	}
	menu := w.Menu{Action: "/columns", Hidden: []w.Hidden{{Name: "csrf", Value: "1"}}, Items: []w.MenuPart{
		item("Email", "email", true), item("Birthday", "birthday", true), item("Phone", "phone", false),
		w.MenuFoot{Items: []w.Button{{Label: "Apply", Small: true, Kind: w.ButtonPrimary, Submit: true}}}}}
	h := browser.NewHarness(t)
	url := h.RegisterPage("columns", browser.PageHTML(`<div class="tc-root">`+menu.Render()+`</div>`, ""))
	sent := `[...new FormData(document.querySelector("form.tc-menu")).getAll("col")].join(",")`
	var before, ticked, moved, dragged string
	err := h.Run(url, 900, 600,
		chromedp.Evaluate(sent, &before),
		chromedp.Click(`label.tc-menu__item:nth-of-type(3)`, chromedp.ByQuery),
		chromedp.Evaluate(sent, &ticked),
		chromedp.Evaluate(`(() => {
			const box = document.querySelector("input[value=email]");
			box.focus();
			box.dispatchEvent(new KeyboardEvent("keydown", {key: "ArrowDown", altKey: true, bubbles: true, cancelable: true}));
			return [...new FormData(document.querySelector("form.tc-menu")).getAll("col")].join(",");
		})()`, &moved),
		chromedp.Evaluate(`(() => {
			const handle = document.querySelector("label.tc-menu__item:last-of-type > svg"), first = document.querySelector("label.tc-menu__item");
			const hr = handle.getBoundingClientRect(), fr = first.getBoundingClientRect();
			handle.dispatchEvent(new PointerEvent("pointerdown", {bubbles: true, clientX: hr.x + 2, clientY: hr.y + 2}));
			document.dispatchEvent(new PointerEvent("pointermove", {bubbles: true, clientX: fr.x + 40, clientY: fr.y + 2}));
			document.dispatchEvent(new PointerEvent("pointerup", {bubbles: true}));
			return [...new FormData(document.querySelector("form.tc-menu")).getAll("col")].join(",");
		})()`, &dragged),
	)
	if err != nil {
		t.Fatal(err)
	}
	if before != "email,birthday" || ticked != "email,birthday,phone" {
		t.Errorf("sent %q, after ticking Phone %q", before, ticked)
	}
	if moved != "birthday,email,phone" {
		t.Errorf("after Alt+↓ on Email: %q, want birthday,email,phone", moved)
	}
	if dragged != "phone,birthday,email" {
		t.Errorf("after dragging Phone to the top: %q, want phone,birthday,email", dragged)
	}
}

// TestSignInOutsideItsForm (task 39): a button under the fields, outside the
// form, submits it by its ID, and Enter in a field does the same — with no
// script; the links under the button go to their addresses.
func TestSignInOutsideItsForm(t *testing.T) {
	needChrome(t)
	form := w.Form{ID: "signin", Action: "/login", Fields: []w.Hidden{{Name: "csrf", Value: "1"}}, Items: []w.FormPart{
		w.Field{Label: "Login", Parts: []w.FieldPart{w.Input{Name: "login", Value: "mix", Label: "Login"}}},
		w.Field{Label: "Password", Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "password", Value: "x", Label: "Password"}}}}}
	stack := w.AuthStack{Items: []w.AuthStackPart{form, w.Button{Label: "Sign in", Kind: w.ButtonPrimary, Form: "signin"}},
		Links: []w.Link{{Text: "Forgotten your password?", Href: "/forgot"}}}
	body := `<div class="tc-root">` + w.Dialog{Items: []w.DialogPart{stack}}.Render() + `</div>`
	h := browser.NewHarness(t)
	url := h.RegisterPage("signin", strings.Replace(browser.PageHTML(body, ""), `<script src="/tc.js"></script>`, "", 1))
	listen := `(() => { window.__sent = []; document.addEventListener("submit", e => { e.preventDefault();
		window.__sent.push((e.submitter ? e.submitter.textContent : "enter") + ":" + [...new FormData(e.target, e.submitter)].map(p => p.join("=")).join("&")); }, true); return true; })()`
	var ok bool
	var sent []string
	var href string
	err := h.Run(url, 900, 600,
		chromedp.Evaluate(listen, &ok),
		chromedp.Click(`button[form="signin"]`, chromedp.ByQuery),
		chromedp.Evaluate(`(() => { const i = document.querySelector("input[name=login]"); i.focus();
			i.dispatchEvent(new KeyboardEvent("keydown", {key: "Enter", bubbles: true})); document.querySelector("form#signin").requestSubmit(); return true; })()`, &ok),
		chromedp.Evaluate(`window.__sent`, &sent),
		chromedp.Evaluate(`document.querySelector(".tc-auth__link").getAttribute("href")`, &href),
	)
	if err != nil {
		t.Fatal(err)
	}
	if len(sent) < 1 || sent[0] != "Sign in:csrf=1&login=mix&password=x" {
		t.Errorf("the button outside the form sent %v", sent)
	}
	if href != "/forgot" {
		t.Errorf("the link under the button goes to %q", href)
	}
}

// TestClosedMenuWithoutScripts (Carrel's task 60): a closed menu is a details,
// so its button opens it without a script (U-8); with the library's script
// it opens, closes on Esc and on a click elsewhere.
func TestClosedMenuWithoutScripts(t *testing.T) {
	needChrome(t)
	drop := w.Drop{Trigger: w.Button{Icon: "dots", Title: "More"}, Menu: w.Menu{Items: []w.MenuPart{w.MenuItem{Label: "Export", Href: "/x.vcf"}}}}
	body := `<div class="tc-root"><div class="tc-main"><div class="tc-bar">` + drop.Render() + `</div><p class="away">away</p></div></div>`
	h := browser.NewHarness(t)
	shown := `(() => { const m = document.querySelector("details.tc-drop .tc-menu"); return !!m && m.checkVisibility({contentVisibilityAuto: true, visibilityProperty: true}); })()`
	bare := h.RegisterPage("drop-bare", strings.Replace(browser.PageHTML(body, ""), `<script src="/tc.js"></script>`, "", 1))
	var before, after bool
	if err := h.Run(bare, 900, 600,
		chromedp.Evaluate(shown, &before),
		chromedp.Click(`details.tc-drop > summary`, chromedp.ByQuery),
		chromedp.Evaluate(shown, &after),
	); err != nil {
		t.Fatal(err)
	}
	if before || !after {
		t.Errorf("without a script: shown at start %v, after a click %v; want false, true", before, after)
	}
	page := h.RegisterPage("drop", browser.PageHTML(body, ""))
	var opened, escaped, again, away bool
	if err := h.Run(page, 900, 600,
		chromedp.Click(`details.tc-drop > summary`, chromedp.ByQuery),
		chromedp.Evaluate(shown, &opened),
		chromedp.KeyEvent("\x1b"),
		chromedp.Evaluate(shown, &escaped),
		chromedp.Click(`details.tc-drop > summary`, chromedp.ByQuery),
		chromedp.Evaluate(shown, &again),
		chromedp.Click(`.away`, chromedp.ByQuery),
		chromedp.Evaluate(shown, &away),
	); err != nil {
		t.Fatal(err)
	}
	if !opened || escaped || !again || away {
		t.Errorf("with the script: opened %v, after Esc %v, again %v, after a click away %v; want true, false, true, false", opened, escaped, again, away)
	}
}

// TestCheckFormSubmits (Carrel's task 67): a bar's «Compact rows» is a GET
// form — without a script its Apply sends it, with the library's script a
// change of the check does, and Apply is not shown.
func TestCheckFormSubmits(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	form := func(action string) string {
		f := w.CheckForm{Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}}, Action: action, Hidden: []w.Hidden{{Name: "sort", Value: "az"}}, Apply: w.Button{Label: "Apply", Small: true, Submit: true}}
		return `<div class="tc-root"><div class="tc-main"><div class="tc-bar">` + f.Render() + `</div></div></div>`
	}
	bareURL := h.RegisterPage("checkform-bare", "")
	h.RegisterPage("checkform-bare", strings.Replace(browser.PageHTML(form(bareURL), ""), `<script src="/tc.js"></script>`, "", 1))
	var applyShown bool
	var search string
	if err := h.Run(bareURL, 900, 600,
		chromedp.Evaluate(`document.querySelector(".tc-checkform button").checkVisibility()`, &applyShown),
		chromedp.Click(`.tc-checkform .tc-check`, chromedp.ByQuery),
		chromedp.Click(`.tc-checkform button`, chromedp.ByQuery),
		chromedp.Sleep(300*time.Millisecond),
		chromedp.Evaluate(`location.search`, &search),
	); err != nil {
		t.Fatal(err)
	}
	if !applyShown || !strings.Contains(search, "compact=1") || !strings.Contains(search, "sort=az") {
		t.Errorf("without a script: Apply shown %v, sent %q; want true and compact=1, sort=az", applyShown, search)
	}
	pageURL := h.RegisterPage("checkform", "")
	h.RegisterPage("checkform", browser.PageHTML(form(pageURL), ""))
	var hidden bool
	search = ""
	if err := h.Run(pageURL, 900, 600,
		chromedp.Evaluate(`!document.querySelector(".tc-checkform button").checkVisibility()`, &hidden),
		chromedp.Click(`.tc-checkform .tc-check`, chromedp.ByQuery),
		chromedp.Sleep(300*time.Millisecond),
		chromedp.Evaluate(`location.search`, &search),
	); err != nil {
		t.Fatal(err)
	}
	if !hidden || !strings.Contains(search, "compact=1") {
		t.Errorf("with the script: no Apply %v, sent %q; want true and compact=1", hidden, search)
	}
}

// TestZonePick (Carrel's task 73): a time zone is a pick with a menu of radio
// items — without a script the chosen one goes with the form (the search line
// is not shown: there is nothing to search with); with the library's script
// the search keeps what matches, and the choice becomes the pick's value.
func TestZonePick(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	zone := func(label, value string, on bool) w.MenuItem {
		return w.MenuItem{Label: label, Name: "tz", Value: value, On: on}
	}
	drop := w.Drop{Trigger: w.Pick{Value: "Europe/Moscow", Aside: &w.Dim{Parts: []w.Inline{w.Text("as sent")}, Variant: w.DimIndent}, Label: "Time zone"},
		Menu: w.Menu{Items: []w.MenuPart{
			w.Input{Kind: w.InputSearch, Placeholder: "Find a zone", Label: "Find a zone"},
			w.Rubric{Text: "As sent", Form: w.RubricInline},
			zone("Europe/Moscow", "Europe/Moscow", true),
			w.MenuSep{},
			zone("Asia/Yekaterinburg", "Asia/Yekaterinburg", false),
			zone("UTC", "UTC", false),
			w.MenuSep{},
			zone("Floating — the same clock time anywhere", "floating", false)}, Width: w.MenuWidth290, Pop: w.MenuUnder}}
	form := func(action string) string {
		return `<div class="tc-root"><div class="tc-main"><form action="` + action + `" method="get">` + drop.Render() +
			`<button class="send" type="submit">Save</button></form></div></div>`
	}
	bareURL := h.RegisterPage("zone-bare", "")
	h.RegisterPage("zone-bare", strings.Replace(browser.PageHTML(form(bareURL), ""), `<script src="/tc.js"></script>`, "", 1))
	var searchShown bool
	var search string
	if err := h.Run(bareURL, 900, 600,
		chromedp.Click(`details.tc-drop > summary`, chromedp.ByQuery),
		chromedp.Evaluate(`document.querySelector('.tc-menu > input[type="search"]').checkVisibility()`, &searchShown),
		chromedp.Click(`label:has(> input[value="UTC"])`, chromedp.ByQuery),
		chromedp.Evaluate(`document.querySelector(".tc-drop").open = false`, nil),
		chromedp.Click(`.send`, chromedp.ByQuery),
		chromedp.Sleep(300*time.Millisecond),
		chromedp.Evaluate(`location.search`, &search),
	); err != nil {
		t.Fatal(err)
	}
	if searchShown || !strings.Contains(search, "tz=UTC") {
		t.Errorf("without a script: search shown %v, sent %q; want false, tz=UTC", searchShown, search)
	}
	page := h.RegisterPage("zone", browser.PageHTML(form("/x"), ""))
	var focused, left, value, back string
	var open bool
	visibleItems := `[...document.querySelectorAll(".tc-menu > :not(input)")].filter(e => e.checkVisibility()).map(e => e.textContent.trim() || e.className).join("|")`
	if err := h.Run(page, 900, 600,
		chromedp.Click(`details.tc-drop > summary`, chromedp.ByQuery),
		chromedp.Evaluate(`document.activeElement.type`, &focused),
		chromedp.SendKeys(`.tc-menu > input[type="search"]`, "yek", chromedp.ByQuery),
		chromedp.Evaluate(visibleItems, &left),
		chromedp.Click(`label:has(> input[value="Asia/Yekaterinburg"])`, chromedp.ByQuery),
		chromedp.Evaluate(`document.querySelector("details.tc-drop > summary").textContent`, &value),
		chromedp.Evaluate(`document.querySelector("details.tc-drop").open`, &open),
		chromedp.Click(`details.tc-drop > summary`, chromedp.ByQuery),
		chromedp.Click(`label:has(> input[value="Europe/Moscow"])`, chromedp.ByQuery),
		chromedp.Evaluate(`document.querySelector("details.tc-drop > summary").textContent`, &back),
	); err != nil {
		t.Fatal(err)
	}
	if focused != "search" || left != "Asia/Yekaterinburg" || value != "Asia/Yekaterinburg" || open || back != "Europe/Moscow as sent" {
		t.Errorf("with the script: focus %q, left %q, value %q, open %v, back %q; want search, Asia/Yekaterinburg, Asia/Yekaterinburg, false, Europe/Moscow as sent",
			focused, left, value, open, back)
	}
}

// TestAttendeesSuggest (Carrel's task 77): the attendees of an event are a
// tag field — with the library's script typing asks the server and an
// item's template is the tag it adds, a cross takes a tag away, Enter does
// not send the edit; without a script a cross unticks the tag and the typed
// words go with the form for the server to find.
func TestAttendeesSuggest(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	person := func(name, mail, colour string) *w.Tag {
		return &w.Tag{Label: name, Mark: &w.Mark{Colour: colour, Size: w.MarkH11}, Removable: true, Title: mail, Name: "attendee", Value: "mailto:" + mail}
	}
	items := w.MenuItems(w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "Вера Лебедева", Adds: person("Вера Лебедева", "vera@example.net", "#A05A2C")},
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "Вениамин Орлов", Adds: person("Вениамин Орлов", "orlov@example.com", "#4A6B52")})
	sug := h.RegisterPage("attendees-suggest", items)
	form := func(action string) string {
		field := w.Drop{Trigger: w.TagInput{Columns: true, Name: "attendee", Parts: []w.TagInputPart{
			w.Tag{Label: "Мария Лебедева", Fixed: true, Value: "mailto:m.lebedeva@example.org"},
			*person("Анна Ковалёва", "a.kovaleva@example.org", "#4A6B52")},
			Add: &w.TagAdd{Name: "attendee.add", Placeholder: "Name or address", Label: "Add attendee", Suggest: sug}},
			Menu: w.Menu{Width: w.MenuWidth420, Pop: w.MenuUnder}}
		return `<div class="tc-root"><div class="tc-main"><form action="` + action + `" method="get"><div class="tc-field tc-field--full">` + field.Render() +
			`</div><button class="send" type="submit">Save</button></form></div></div>`
	}
	bareURL := h.RegisterPage("attendees-bare", "")
	h.RegisterPage("attendees-bare", strings.Replace(browser.PageHTML(form(bareURL), ""), `<script src="/tc.js"></script>`, "", 1))
	var search string
	if err := h.Run(bareURL, 900, 600,
		chromedp.Click(`label.tc-tag--pill svg`, chromedp.ByQuery),
		chromedp.SendKeys(`.tc-in__add`, "ivan@mail.example", chromedp.ByQuery),
		chromedp.Click(`.send`, chromedp.ByQuery),
		chromedp.Sleep(300*time.Millisecond),
		chromedp.Evaluate(`decodeURIComponent(location.search)`, &search),
	); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(search, "attendee=mailto:m.lebedeva") || strings.Contains(search, "kovaleva") || !strings.Contains(search, "attendee.add=ivan@mail.example") {
		t.Errorf("without a script: sent %q; want the organiser and the typed address, not the untick one", search)
	}
	page := h.RegisterPage("attendees", browser.PageHTML(form("/x"), ""))
	var shown, after, added string
	var open, gone bool
	var where string
	visible := `[...document.querySelectorAll(".tc-drop .tc-menu .tc-menu__item")].filter(e => e.checkVisibility()).map(e => e.textContent.trim()).join("|")`
	if err := h.Run(page, 900, 600,
		chromedp.SendKeys(`.tc-in__add`, "ве", chromedp.ByQuery),
		chromedp.WaitVisible(`.tc-drop .tc-menu .tc-menu__item`, chromedp.ByQuery),
		chromedp.Evaluate(visible, &shown),
		chromedp.Click(`.tc-drop .tc-menu .tc-menu__item`, chromedp.ByQuery),
		chromedp.Evaluate(`[...document.querySelectorAll(".tc-in--tags .tc-check__box")].map(b => b.name + "=" + b.value).join("|")`, &added),
		chromedp.Evaluate(`!document.querySelector(".tc-drop .tc-menu").checkVisibility()`, &gone),
		chromedp.Click(`label.tc-tag--pill[title="a.kovaleva@example.org"] svg`, chromedp.ByQuery),
		chromedp.Evaluate(`[...document.querySelectorAll(".tc-in--tags .tc-tag")].map(e => e.textContent.trim()).join("|")`, &after),
		chromedp.Focus(`.tc-in__add`, chromedp.ByQuery),
		chromedp.KeyEvent("\r"),
		chromedp.Sleep(200*time.Millisecond),
		chromedp.Evaluate(`location.pathname`, &where),
		chromedp.Evaluate(`!!document.querySelector(".tc-drop .tc-menu") && document.querySelector(".tc-drop .tc-menu").checkVisibility()`, &open),
	); err != nil {
		t.Fatal(err)
	}
	if shown != "Вера Лебедева|Вениамин Орлов" || !gone || !strings.Contains(added, "attendee=mailto:vera@example.net") ||
		after != "Мария Лебедева|Вера Лебедева" || strings.HasSuffix(where, "/x") || open {
		t.Errorf("with the script: shown %q, closed %v, ticks %q, tags after a cross %q, Enter went to %q (menu open %v)", shown, gone, added, after, where, open)
	}
}
