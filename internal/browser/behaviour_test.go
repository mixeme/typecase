package browser_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/chromedp/cdproto/runtime"
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
		w.Field{Label: "Start time", Parts: []w.FieldPart{w.Input{Name: "start.time", Label: "Start time"}}},
		w.Field{Label: "Due time", Parts: []w.FieldPart{w.Input{Name: "due.time", Label: "Due time"}}},
		w.Field{Parts: []w.FieldPart{w.Check{Name: "allday", Value: "1", Disables: "start.time due.time", Parts: []w.CheckPart{w.Text("All day")}}}},
		w.Field{Label: "Theme", Parts: []w.FieldPart{w.Seg{Name: "theme", Theme: true, Options: []w.SegOption{{Label: "Auto", Value: "auto", On: true}, {Label: "Dark", Value: "dark"}}}}},
	}}
	row := w.ExcerptRow{Mark: w.Mark{Colour: "#4A6B52"}, Title: []w.Inline{w.Text("Заметка")}, Date: "10 Aug", Source: "nextcloud"}
	list := w.List{Compact: true, Rows: []w.ListRow{row, row}}
	page := h.RegisterPage("checks", browser.PageHTML(`<div class="tc-root"><div class="tc-main">`+form.Render()+list.Render()+`</div></div>`, flush))
	var off0, off1, back bool
	var times string
	var theme, stored string
	var height float64
	err := h.Run(page, 1200, 700,
		chromedp.Evaluate(`document.querySelector('[name=password]').disabled`, &off0),
		chromedp.Click(`.tc-check`, chromedp.ByQuery),
		chromedp.Evaluate(`document.querySelector('[name=password]').disabled`, &off1),
		chromedp.Click(`.tc-check`, chromedp.ByQuery),
		chromedp.Evaluate(`document.querySelector('[name=password]').disabled`, &back),
		chromedp.Click(`label.tc-check:has([name=allday])`, chromedp.ByQuery),
		chromedp.Evaluate(`[...document.querySelectorAll('[name="start.time"], [name="due.time"]')].map(f => f.disabled).join(" ")`, &times),
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
	if times != "true true" {
		t.Errorf("All day ticked: Start time and Due time off %q; want both (Carrel's task 111)", times)
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

// TestCheckFormInMenu (Carrel's task 82): a bar's Compact rows has its twin
// under «⋯» — the same GET form, the check a menu item: without a script its
// Apply sends it, with the library's script a change does.
func TestCheckFormInMenu(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	page := func(action string) string {
		more := w.Drop{More: true, Trigger: w.Button{Icon: "dots", Title: "More"}, Menu: w.Menu{Items: []w.MenuPart{
			w.CheckForm{Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}}, Action: action, Hidden: []w.Hidden{{Name: "sort", Value: "az"}}, Apply: w.Button{Label: "Apply", Small: true, Submit: true}}},
			Pop: w.MenuUnderEnd, Inline: true}}
		return `<div class="tc-root"><div class="tc-main"><div class="tc-bar"><span class="tc-seg is-2nd"><button class="tc-seg__opt" type="button">A–Z</button></span>` + more.Render() + `</div></div></div>`
	}
	bareURL := h.RegisterPage("menu-checkform-bare", "")
	h.RegisterPage("menu-checkform-bare", strings.Replace(browser.PageHTML(page(bareURL), ""), `<script src="/tc.js"></script>`, "", 1))
	var search string
	if err := h.Run(bareURL, 380, 600,
		chromedp.Click(`details.tc-drop > summary`, chromedp.ByQuery),
		chromedp.Click(`.tc-menu label.tc-menu__item`, chromedp.ByQuery),
		chromedp.Click(`.tc-checkform__apply button`, chromedp.ByQuery),
		chromedp.Sleep(300*time.Millisecond),
		chromedp.Evaluate(`location.search`, &search),
	); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(search, "compact=1") || !strings.Contains(search, "sort=az") {
		t.Errorf("without a script: sent %q; want compact=1 and sort=az", search)
	}
	liveURL := h.RegisterPage("menu-checkform", "")
	h.RegisterPage("menu-checkform", browser.PageHTML(page(liveURL), ""))
	var applyShown bool
	search = ""
	if err := h.Run(liveURL, 380, 600,
		chromedp.Click(`details.tc-drop > summary`, chromedp.ByQuery),
		chromedp.Evaluate(`document.querySelector(".tc-checkform__apply button").checkVisibility()`, &applyShown),
		chromedp.Click(`.tc-menu label.tc-menu__item`, chromedp.ByQuery),
		chromedp.Sleep(300*time.Millisecond),
		chromedp.Evaluate(`location.search`, &search),
	); err != nil {
		t.Fatal(err)
	}
	if applyShown || !strings.Contains(search, "compact=1") {
		t.Errorf("with the script: Apply shown %v, sent %q; want false, compact=1", applyShown, search)
	}
}

// TestSaveAs (Carrel's task 94): Save as makes a file of the text shown —
// with the library's script, in the page; without a script the button is not
// shown, and Copy and selecting the text remain.
func TestSaveAs(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	body := `<div class="tc-root"><div class="tc-main">` + w.Diag{ID: "script", Text: "#!/bin/sh\necho hello"}.Render() +
		w.Button{Label: "Save as", SaveAs: "carrel.sh", SaveFrom: "script"}.Render() + `</div></div>`
	bare := h.RegisterPage("save-bare", strings.Replace(browser.PageHTML(body, ""), `<script src="/tc.js"></script>`, "", 1))
	var shown bool
	if err := h.Run(bare, 900, 600,
		chromedp.Evaluate(`document.querySelector("[data-tc-save]").checkVisibility()`, &shown),
	); err != nil {
		t.Fatal(err)
	}
	if shown {
		t.Error("without a script Save as is shown")
	}
	page := h.RegisterPage("save", browser.PageHTML(body, ""))
	var got string
	if err := h.Run(page, 900, 600,
		chromedp.Evaluate(`window.__saved = []; HTMLAnchorElement.prototype.click = function () { const a = this; window.__saved.push(fetch(a.href).then(r => r.text()).then(x => a.download + "|" + x)); }; true`, nil),
		chromedp.Click(`[data-tc-save]`, chromedp.ByQuery),
		chromedp.Evaluate(`Promise.all(window.__saved).then(x => x.join(";"))`, &got, func(p *runtime.EvaluateParams) *runtime.EvaluateParams { return p.WithAwaitPromise(true) }),
	); err != nil {
		t.Fatal(err)
	}
	if got != "carrel.sh|#!/bin/sh\necho hello" {
		t.Errorf("with the script: saved %q; want carrel.sh with the text", got)
	}
}

// TestFilePick: a button or a menu item that picks a file is the label of its
// file input (Carrel's task 109). With the script a picked file sends its form
// at once and «Upload» is not shown; without one the input shows what was
// picked and «Upload» beside it sends it.
func TestFilePick(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	photo := &w.FileField{Form: "post-photo", Name: "photo", Accept: "image/*", Label: "Photo"}
	attach := &w.FileField{Form: "post-attach", Name: "file", Label: "From this computer"}
	body := `<form id="post-photo" method="post" action="/photo" enctype="multipart/form-data"></form>` +
		`<form id="post-attach" method="post" action="/attach" enctype="multipart/form-data"></form>` +
		`<div class="tc-root"><div class="tc-main">` +
		w.Line{Items: []w.LinePart{w.Button{Label: "Replace", File: photo}, w.Button{Label: "Upload", Form: "post-photo", NoScript: true}}}.Render() +
		w.Menu{Items: []w.MenuPart{w.MenuItem{Icon: "up", Label: "From this computer", File: attach}, w.MenuItem{Label: "Upload", Form: "post-attach", NoScript: true}}}.Render() +
		`</div></div>`
	f := filepath.Join(t.TempDir(), "anna.png")
	if err := os.WriteFile(f, []byte("png"), 0o644); err != nil {
		t.Fatal(err)
	}
	const shown = `[...document.querySelectorAll(".tc-file, .tc-nojs")].map(e => e.checkVisibility({opacityProperty: true})).join(" ")`
	bare := h.RegisterPage("file-bare", strings.Replace(browser.PageHTML(body, ""), `<script src="/tc.js"></script>`, "", 1))
	var got string
	if err := h.Run(bare, 900, 600, chromedp.Evaluate(shown, &got)); err != nil {
		t.Fatal(err)
	}
	if got != "true true true true" {
		t.Errorf("without a script: inputs and Upload shown %q; want all", got)
	}
	page := h.RegisterPage("file", browser.PageHTML(body, ""))
	var sent string
	if err := h.Run(page, 900, 600,
		chromedp.Evaluate(shown, &got),
		chromedp.Evaluate(`window.__sent = []; document.addEventListener("submit", e => { e.preventDefault(); window.__sent.push(e.target.id + ":" + new FormData(e.target).getAll(e.target.id == "post-photo" ? "photo" : "file").map(x => x.name).join()); }); true`, nil),
		chromedp.SetUploadFiles(`label.tc-btn .tc-file`, []string{f}, chromedp.ByQuery),
		chromedp.SetUploadFiles(`label.tc-menu__item .tc-file`, []string{f}, chromedp.ByQuery),
		chromedp.Evaluate(`window.__sent.join(" ")`, &sent),
	); err != nil {
		t.Fatal(err)
	}
	if got != "false false false false" {
		t.Errorf("with the script: inputs and Upload shown %q; want none", got)
	}
	if sent != "post-photo:anna.png post-attach:anna.png" {
		t.Errorf("with the script: sent %q; want each form with its file", sent)
	}
}

// TestFileSizes: an upload picks several files or a folder, and the lengths
// of the picked files go before them as «size.<name>» (Carrel's task 119); the
// words of a drop place pick files too, and those wait for their form's
// button, each a row of the place's list with a Remove of its own (the
// owner's decision of 09.10). The pick of a folder itself is the system's dialog and is not made
// here: the input is checked to ask for one. A form with a file field of its
// own says it sends files.
func TestFileSizes(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	menu := w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "doc", Label: "Files", File: &w.FileField{Form: "post-upload", Name: "file", Label: "Files", Size: "size", Many: true}},
		w.MenuItem{Icon: "file", Label: "Folder", File: &w.FileField{Form: "post-upload", Name: "file", Label: "Folder", Size: "size", Folder: true}},
		w.MenuItem{Label: "Upload", Form: "post-upload", NoScript: true}}}
	place := w.Empty{Variant: w.EmptyDashed, Parts: []w.EmptyPart{w.Text("Drop files here, or "),
		w.FilePick{Label: "choose on the computer", File: w.FileField{Name: "file", Label: "Files", List: "picked", Size: "size", Many: true},
			Row: &w.Line{Place: w.Grown, Gap: w.LineGap8, Items: []w.LinePart{w.Text("{name} "), w.Dim{Parts: []w.Inline{w.Text("{size}")}},
				w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Place: w.Pushed, Removes: true}}}},
		w.Line{Block: true, Gap: w.LineGap8, Place: w.Grown, ID: "picked"}}}

	body := `<form id="post-upload" method="post" action="/upload" enctype="multipart/form-data"></form>` +
		`<div class="tc-root"><div class="tc-main">` + menu.Render() +
		w.Sheet{Title: "Upload files", ID: "sheet", Action: "/sheet", Body: []w.SheetPart{place}, Foot: []w.SheetFootPart{w.Button{Label: "Upload", Kind: w.ButtonPrimary, Submit: true}}}.Render() +
		`</div></div>`
	dir := t.TempDir()
	a, b := filepath.Join(dir, "plan.md"), filepath.Join(dir, "kazan.png")
	if err := os.WriteFile(a, []byte("abc"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(b, []byte("abcde"), 0o644); err != nil {
		t.Fatal(err)
	}
	const catch = `window.__sent = []; document.addEventListener("submit", e => { e.preventDefault(); window.__sent.push(e.target.id + ": " + [...new FormData(e.target)].map(([k, v]) => k + "=" + (v instanceof File ? v.name : v)).join(" ")); }); true`
	page := h.RegisterPage("file-sizes", browser.PageHTML(body, ""))
	var attrs, sent, waits, rows, button string
	if err := h.Run(page, 900, 600,
		chromedp.Evaluate(catch, nil),
		chromedp.Evaluate(`[...document.querySelectorAll(".tc-file")].map(f => (f.multiple ? "many" : "one") + (f.webkitdirectory ? "+folder" : "") + (f.labels.length ? "" : " no label")).join(" ") + " · " + document.getElementById("sheet").enctype`, &attrs),
		chromedp.SetUploadFiles(`label.tc-menu__item .tc-file`, []string{a, b}, chromedp.ByQuery),
		chromedp.Evaluate(`window.__sent.join(" | ")`, &sent),
		chromedp.SetUploadFiles(`.tc-empty__pick .tc-file`, []string{a, b}, chromedp.ByQuery),
		chromedp.Evaluate(`String(window.__sent.length)`, &waits),
		chromedp.Evaluate(`[...document.querySelectorAll("#picked > [data-tc-picked]")].map(r => r.textContent.replace("Remove", "").trim()).join(" | ")`, &rows),
		chromedp.Click(`#picked > [data-tc-picked] button`, chromedp.ByQuery),
		chromedp.Click(`.tc-sheet__foot button`, chromedp.ByQuery),
		chromedp.Evaluate(`window.__sent.slice(1).join(" | ")`, &button),
	); err != nil {
		t.Fatal(err)
	}
	if attrs != "many many+folder many · multipart/form-data" {
		t.Errorf("the inputs: %q; want Files and the words many, Folder a folder, each in its label, and the sheet a form that sends files", attrs)
	}
	// Folder's input, empty, goes as a part without a name: a browser sends it so
	if sent != "post-upload: size.plan.md=3 size.kazan.png=5 file=plan.md file=kazan.png file=" {
		t.Errorf("Files: sent %q; want the lengths before the files", sent)
	}
	if waits != "1" {
		t.Errorf("the words of the place: %s forms sent after the pick; want the files to wait for the button", waits)
	}
	if rows != "plan.md 3 B | kazan.png 5 B" {
		t.Errorf("the picked files: rows %q; want a row for each, its name and size", rows)
	}
	// the first file's Remove was pressed: it does not go, nor does its length;
	// the input waiting for the next pick goes empty
	if button != "sheet: file= size.kazan.png=5 file=kazan.png" {
		t.Errorf("the button: sent %q; want the file left in the list, with its length", button)
	}
}

// TestTagList: a tag field's line has the browser's own list of choices on a
// page without a script (Carrel's task 96); where the server suggests, the
// live layer takes the list away, and keeps it where it does not.
func TestTagList(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	opts := []w.ListOption{{Value: "7c1e-40aa", Label: "Поездка в Казань · event"}, {Value: "91de-4402", Label: "Собрать чемодан · task"}}
	field := func(name, suggest string) string {
		return w.TagInput{Name: name, Add: &w.TagAdd{Name: name + ".add", Label: "Related to", Suggest: suggest, Options: opts}}.Render()
	}
	body := `<div class="tc-root"><div class="tc-main"><form action="/x">` + field("related", "/suggest") + field("plain", "") + `</form></div></div>`
	const lists = `[...document.querySelectorAll(".tc-in__add")].map(e => (e.list ? [...e.list.options].map(o => o.value + "=" + o.label).join(",") : "-")).join(" ")`
	bare := h.RegisterPage("list-bare", strings.Replace(strings.Replace(browser.PageHTML(body, ""), `<script src="/tc.js"></script>`, "", 1), `<script src="/tc-live.js"></script>`, "", 1))
	var got string
	if err := h.Run(bare, 900, 600, chromedp.Evaluate(lists, &got)); err != nil {
		t.Fatal(err)
	}
	both := "7c1e-40aa=Поездка в Казань · event,91de-4402=Собрать чемодан · task"
	if got != both+" "+both {
		t.Errorf("without a script: lists %q; want both choices on each line", got)
	}
	page := h.RegisterPage("list", browser.PageHTML(body, ""))
	if err := h.Run(page, 900, 600,
		chromedp.Click(`[name="related.add"]`, chromedp.ByQuery),
		chromedp.Click(`[name="plain.add"]`, chromedp.ByQuery),
		chromedp.Evaluate(lists, &got),
	); err != nil {
		t.Fatal(err)
	}
	if got != "- "+both {
		t.Errorf("with the live layer: lists %q; want none where the server suggests, the choices where it does not", got)
	}
}

// TestNotebookPick: a pick of collections carries the chosen one's mark
// (Carrel's task 103); a read-only collection is shown and is not chosen.
func TestNotebookPick(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	book := func(label, value, colour string, on bool, reason string) w.MenuItem {
		it := w.MenuItem{Mark: &w.Mark{Colour: colour}, Label: label, Name: "notebook", Value: value, On: on, Reason: reason}
		if reason != "" {
			it.Trail = w.Badge{Label: "ro", Place: w.Pushed}
		}
		return it
	}
	drop := w.Drop{Trigger: w.Pick{Value: "nextcloud · Записи", Mark: &w.Mark{Colour: "#4A6B52"}, Label: "Notebook"},
		Menu: w.Menu{Items: []w.MenuPart{
			book("nextcloud · Записи", "zapisi", "#4A6B52", true, ""),
			book("nextcloud · Дневник", "dnevnik", "#6B4E7A", false, ""),
			book("nextcloud · Общие с семьёй", "semya", "#3F6E8C", false, "Read-only: this account cannot write here")},
			Width: w.MenuWidth290, Pop: w.MenuUnder}}
	page := h.RegisterPage("notebook", browser.PageHTML(`<div class="tc-root"><div class="tc-main"><form action="/x">`+drop.Render()+`</form></div></div>`, ""))
	const pick = `(() => { const p = document.querySelector("summary.is-pick"); const m = p.querySelector(".tc-mark"); return (m ? m.getAttribute("data-tc-colour") : "-") + " " + p.textContent.trim(); })()`
	var chosen, after string
	var off bool
	if err := h.Run(page, 900, 600,
		chromedp.Click(`summary.is-pick`, chromedp.ByQuery),
		chromedp.Click(`label:has(> input[value="dnevnik"])`, chromedp.ByQuery),
		chromedp.Evaluate(pick, &chosen),
		chromedp.Evaluate(`document.querySelector('input[value="semya"]').disabled`, &off),
		chromedp.Click(`summary.is-pick`, chromedp.ByQuery),
		chromedp.Click(`label:has(> input[value="semya"])`, chromedp.ByQuery),
		chromedp.Evaluate(pick, &after),
	); err != nil {
		t.Fatal(err)
	}
	if chosen != "#6B4E7A nextcloud · Дневник" {
		t.Errorf("chosen: pick %q; want Дневник with its mark", chosen)
	}
	if !off || after != chosen {
		t.Errorf("read-only: off %v, pick after choosing it %q; want off and Дневник kept", off, after)
	}
}

// TestFileInForm: on a new record a picked file goes with the form it stands
// in to the attach address, and what is typed goes with it (Carrel's task
// 122); without a script Upload sends the same form there.
func TestFileInForm(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	const to = "/app/calendar/new/attach"
	menu := w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "up", Label: "From this computer", File: &w.FileField{Name: "file", Label: "From this computer", To: to}},
		w.MenuItem{Label: "Upload", To: to, NoScript: true}}}
	body := `<div class="tc-root"><div class="tc-main"><form action="/app/calendar/new" method="post" enctype="multipart/form-data"><input name="title" value="Поездка">` +
		menu.Render() + `</form></div></div>`
	f := filepath.Join(t.TempDir(), "plan.pdf")
	if err := os.WriteFile(f, []byte("pdf"), 0o644); err != nil {
		t.Fatal(err)
	}
	const catch = `window.__sent = []; document.addEventListener("submit", e => { e.preventDefault(); const fd = new FormData(e.target, e.submitter); window.__sent.push(new URL((e.submitter && e.submitter.formAction) || e.target.action).pathname + " " + fd.get("title") + " " + (fd.get("file") && fd.get("file").name)); }); true`
	var sent, action string
	page := h.RegisterPage("file-in-form", browser.PageHTML(body, ""))
	if err := h.Run(page, 900, 600,
		chromedp.Evaluate(catch, nil),
		chromedp.SetUploadFiles(`.tc-file`, []string{f}, chromedp.ByQuery),
		chromedp.Evaluate(`window.__sent.join("|")`, &sent),
		chromedp.Evaluate(`new URL(document.querySelector("form").action).pathname`, &action),
	); err != nil {
		t.Fatal(err)
	}
	if sent != to+" Поездка plan.pdf" || action != "/app/calendar/new" {
		t.Errorf("with the script: sent %q, the form's action after %q; want the whole form to %s, its action kept", sent, action, to)
	}
	bare := h.RegisterPage("file-in-form-bare", strings.Replace(browser.PageHTML(body, ""), `<script src="/tc.js"></script>`, "", 1))
	if err := h.Run(bare, 900, 600,
		chromedp.Evaluate(catch, nil),
		chromedp.SetUploadFiles(`.tc-file`, []string{f}, chromedp.ByQuery),
		chromedp.Click(`.tc-nojs`, chromedp.ByQuery),
		chromedp.Evaluate(`window.__sent.join("|")`, &sent),
	); err != nil {
		t.Fatal(err)
	}
	if sent != to+" Поездка plan.pdf" {
		t.Errorf("without a script: Upload sent %q; want the whole form to %s", sent, to)
	}
}

// TestEditFile: a file of an edit waits for Save (the owner's decision of
// 02.10): picked, it stands as its list's row, «Nothing attached» gives way,
// Remove takes it back; a photo shows at once; Save with files shows how far
// the upload went.
func TestEditFile(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	row := w.Line{Gap: w.LineGap8, Items: []w.LinePart{w.Faced{Face: w.FaceLarge, Items: []w.LinePart{w.Text("{name} "), w.Dim{Parts: []w.Inline{w.Text("{size}")}}}},
		w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Removes: true}}}
	attach := w.Drop{Trigger: w.Button{Label: "Attach", Menu: true}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "up", Label: "From this computer", File: &w.FileField{Name: "file", Label: "From this computer", List: "files"}, Row: row}}}}
	list := w.Line{ID: "files", Gap: w.LineGap8, Items: []w.LinePart{attach, w.Dim{Parts: []w.Inline{w.Text("Nothing attached")}, Empty: true}}}
	photo := w.Photo{ID: "photo", Initials: "АК", Side: []w.FieldPart{w.Button{Label: "Add", File: &w.FileField{Name: "photo", Accept: "image/*", Label: "Photo", Preview: "photo"}}}}
	save := w.Runbar{Meter: w.Meter{}, Parts: []w.RunbarPart{w.Text("Saving · "), w.Mono{Text: "0 of 0 MB"}}, Progress: "edit"}
	body := `<div class="tc-root"><div class="tc-main">` + save.Render() + `<form id="edit" action="/save" method="post" enctype="multipart/form-data"><input name="title" value="Поездка">` +
		list.Render() + w.Field{Parts: []w.FieldPart{photo}}.Render() + `<button class="send" type="submit">Save</button></form></div></div>`
	dir := t.TempDir()
	f := filepath.Join(dir, "plan.pdf")
	if err := os.WriteFile(f, make([]byte, 3000), 0o644); err != nil {
		t.Fatal(err)
	}
	img := filepath.Join(dir, "anna.png")
	if err := os.WriteFile(img, []byte("\x89PNG\r\n\x1a\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	page := h.RegisterPage("edit-file", browser.PageHTML(body, ""))
	const rows = `[...document.querySelectorAll("#files > [data-tc-picked]")].map(r => r.textContent.trim().replace(/\s+/g, " ")).join("|") + " · empty " + document.querySelector("[data-tc-empty]").hidden`
	var picked, removed, preview, line, sent string
	if err := h.Run(page, 900, 600,
		chromedp.Evaluate(`window.__submits = 0; document.addEventListener("submit", () => window.__submits++); true`, nil),
		chromedp.SetUploadFiles(`#files .tc-file`, []string{f}, chromedp.ByQuery),
		chromedp.Evaluate(rows+` + " · in the row " + !!document.querySelector("#files > [data-tc-picked] .tc-file") + " · submits " + window.__submits`, &picked),
		chromedp.Click(`#files > [data-tc-picked] [data-tc-remove]`, chromedp.ByQuery),
		chromedp.Evaluate(rows, &removed),
		chromedp.SetUploadFiles(`#files .tc-file`, []string{f}, chromedp.ByQuery),
		chromedp.SetUploadFiles(`.tc-field__photo-row .tc-file`, []string{img}, chromedp.ByQuery),
		chromedp.Evaluate(`(() => { const i = document.querySelector("#photo img"); return i ? i.src.slice(0, 5) : "-"; })()`, &preview),
		chromedp.Evaluate(`window.__xhr = ""; XMLHttpRequest.prototype.send = function (d) { window.__xhr = [...d.keys()].join(","); this.upload.onprogress({lengthComputable: true, loaded: 1572864, total: 3145728}); }; true`, nil),
		chromedp.Click(`.send`, chromedp.ByQuery),
		chromedp.Evaluate(`(() => { const l = document.querySelector("[data-tc-progress]"); return l.hidden + " " + l.querySelector(".tc-meter__bar").style.width + " " + l.querySelector(".tc-mono").textContent; })()`, &line),
		chromedp.Evaluate(`window.__xhr`, &sent),
	); err != nil {
		t.Fatal(err)
	}
	if picked != "plan.pdf 3 KBRemove · empty true · in the row true · submits 0" {
		t.Errorf("picked: %q; want the row with the name and the size, the input in it, nothing sent", picked)
	}
	if removed != " · empty false" {
		t.Errorf("removed: %q; want no row and «Nothing attached» back", removed)
	}
	if preview != "blob:" {
		t.Errorf("the photo: %q; want the picked picture at once", preview)
	}
	if line != "false 50% 1.5 of 3.0 MB" {
		t.Errorf("Save: the line %q; want it shown, half the way", line)
	}
	if !strings.Contains(sent, "title") || !strings.Contains(sent, "file") || !strings.Contains(sent, "photo") {
		t.Errorf("Save sent %q; want the title, the file and the photo", sent)
	}
}

// TestSwap: a link with Swap changes the page in place with the live layer —
// the column of the page at its address, no reload, the address in the
// history (Carrel's task 104); an edit's bar without a script keeps only its
// Save and Cancel.
func TestSwap(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	page := func(text, other string, on bool) string {
		acts := w.Seg{Options: []w.SegOption{{Label: "Full width", Href: other + "?w=full", On: on}, {Label: "Reading width", Href: other + "?w=read", On: !on}}, Swap: "note-main"}
		return `<div class="tc-root"><div class="tc-main" id="note-main">` + acts.Render() + `<p class="text">` + text + `</p>` +
			w.Button{Label: "Source", Href: other + "?view=source", Swap: "note-main"}.Render() + `</div></div>`
	}
	read := h.RegisterPage("swap-read", "")
	full := h.RegisterPage("swap-full", "")
	h.RegisterPage("swap-read", browser.PageHTML(page("narrow", full, false), ""))
	h.RegisterPage("swap-full", browser.PageHTML(page("wide", read, true), ""))
	var text, path string
	var kept bool
	if err := h.Run(full, 900, 600,
		chromedp.Evaluate(`window.__kept = true; true`, nil),
		chromedp.Click(`a.tc-seg__opt:not(.is-on)`, chromedp.ByQuery),
		chromedp.Sleep(1500*time.Millisecond),
		chromedp.Evaluate(`document.querySelector(".text").textContent`, &text),
		chromedp.Evaluate(`location.pathname + location.search`, &path),
		chromedp.Evaluate(`window.__kept === true`, &kept),
	); err != nil {
		t.Fatal(err)
	}
	if text != "narrow" || !kept || !strings.HasSuffix(path, "?w=read") {
		t.Errorf("with the live layer: text %q, page kept %v, address %q; want the other column in place, no reload, the address", text, kept, path)
	}
	bar := w.Docbar{Edit: true, Items: []w.DocbarPart{w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "note-edit"}, w.Button{Label: "Cancel", Href: "/x"},
		w.Button{Label: "B"}, w.Button{Label: "H2"}}, End: []w.DocbarEnd{w.Check{Parts: []w.CheckPart{w.Text("Auto-save")}}}}
	bare := h.RegisterPage("edit-bar-bare", strings.Replace(browser.PageHTML(`<div class="tc-root">`+bar.Render()+`</div>`, ""), `<script src="/tc.js"></script>`, "", 1))
	var shown string
	if err := h.Run(bare, 900, 600,
		chromedp.Evaluate(`[...document.querySelectorAll(".tc-docbar *")].filter(e => (e.matches(".tc-btn, .tc-check")) && e.checkVisibility()).map(e => e.textContent.trim()).join(" ")`, &shown),
	); err != nil {
		t.Fatal(err)
	}
	if shown != "Save Cancel" {
		t.Errorf("an edit's bar without a script shows %q; want Save Cancel", shown)
	}
}

// TestConflictForm: a conflict is a form (Carrel's task 132): a property's
// versions are radios, the chosen one is marked, «Write this combination»
// sends the choices — without a script too.
func TestConflictForm(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	opt := func(n, v, src, val string, on bool) w.Check {
		return w.Check{Name: "choice." + n, Value: v, Radio: true, On: on, Parts: []w.CheckPart{w.Choice{Source: src, Value: val}}}
	}
	diff := w.Diff{ID: "conflict", Action: "/x", Items: []w.DiffItem{
		{Title: w.Rubric{Text: "EMAIL"}, Options: []w.Check{opt("1", "mine", "Your edit", "a@example.org", true), opt("1", "server", "Server", "b@example.org", false)}}}}
	body := `<div class="tc-root"><div class="tc-main">` + diff.Render() + w.Button{Label: "Write this combination", Kind: w.ButtonPrimary, Form: "conflict"}.Render() + `</div></div>`
	bare := h.RegisterPage("conflict-bare", strings.Replace(browser.PageHTML(body, ""), `<script src="/tc.js"></script>`, "", 1))
	const marked = `[...document.querySelectorAll(".tc-diff__options .tc-check")].map(c => getComputedStyle(c).borderTopColor === getComputedStyle(document.querySelector(".tc-btn.is-primary")).backgroundColor ? "on" : "off").join(" ")`
	var before, after, sent string
	if err := h.Run(bare, 900, 600,
		chromedp.Evaluate(marked, &before),
		chromedp.Click(`label.tc-check:nth-child(2)`, chromedp.ByQuery),
		chromedp.Evaluate(marked, &after),
		chromedp.Evaluate(`document.addEventListener("submit", e => { e.preventDefault(); window.__sent = new URLSearchParams(new FormData(e.target)).toString(); }); true`, nil),
		chromedp.Click(`.tc-btn.is-primary`, chromedp.ByQuery),
		chromedp.Evaluate(`window.__sent || ""`, &sent),
	); err != nil {
		t.Fatal(err)
	}
	if before != "on off" || after != "off on" {
		t.Errorf("the chosen version marked: before %q, after %q; want on off, off on", before, after)
	}
	if sent != "choice.1=server" {
		t.Errorf("sent %q; want choice.1=server", sent)
	}
}

// TestCopy (Carrel's task 118): Copy puts the text it names, or the text of
// the element it names, on the clipboard — with the library's script. Without
// a script a button is not shown and the text is selected by hand; a link
// stays a link to the page that shows the text.
func TestCopy(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	body := `<div class="tc-root"><div class="tc-main">` + w.Diag{ID: "script", Text: "#!/bin/sh\necho hello"}.Render() +
		w.Button{Label: "Copy the script", CopyFrom: "script"}.Render() +
		w.Button{Label: "Copy", Copy: "https://carrel.example/p/9c1f4e", Copied: "Copied"}.Render() + w.Button{Label: "Send"}.Render() +
		w.Button{Label: "Copy link", Href: "/app/files/published/9c1f4e", Copy: "https://carrel.example/p/9c1f4e"}.Render() + `</div></div>`
	bare := h.RegisterPage("copy-bare", strings.Replace(browser.PageHTML(body, ""), `<script src="/tc.js"></script>`, "", 1))
	var shown string
	if err := h.Run(bare, 900, 600,
		chromedp.Evaluate(`[...document.querySelectorAll("[data-tc-copy], [data-tc-copy-from]")].map(b => b.tagName + ":" + b.checkVisibility()).join(" ")`, &shown),
	); err != nil {
		t.Fatal(err)
	}
	if shown != "BUTTON:false BUTTON:false A:true" {
		t.Errorf("without a script: %q; want the buttons not shown and the link shown", shown)
	}
	page := h.RegisterPage("copy", browser.PageHTML(body, ""))
	// what the button says, its right edge, its padding and where its neighbour stands
	const place = `(() => { const b = document.querySelector("button[data-tc-copy]"), r = b.getBoundingClientRect(), c = getComputedStyle(b); return b.textContent + " " + r.right + " " + c.paddingLeft + " " + c.paddingRight + " " + b.nextElementSibling.getBoundingClientRect().x; })()`
	var got, at, was, now string
	if err := h.Run(page, 900, 600,
		chromedp.Evaluate(`window.__copied = []; Object.defineProperty(navigator, "clipboard", {value: {writeText: x => { window.__copied.push(x); return Promise.resolve(); }}, configurable: true}); true`, nil),
		chromedp.Click(`[data-tc-copy-from]`, chromedp.ByQuery),
		chromedp.Evaluate(place, &was),
		chromedp.Click(`button[data-tc-copy]`, chromedp.ByQuery),
		chromedp.Evaluate(place, &now),
		chromedp.Click(`a[data-tc-copy]`, chromedp.ByQuery),
		chromedp.Evaluate(`window.__copied.join("|")`, &got),
		chromedp.Evaluate(`location.pathname`, &at),
	); err != nil {
		t.Fatal(err)
	}
	if want := "#!/bin/sh\necho hello|https://carrel.example/p/9c1f4e|https://carrel.example/p/9c1f4e"; got != want {
		t.Errorf("with the script: copied %q; want %q", got, want)
	}
	if !strings.HasPrefix(was, "Copy ") || !strings.HasPrefix(now, "Copied ") || strings.TrimPrefix(was, "Copy ") != strings.TrimPrefix(now, "Copied ") {
		t.Errorf("the pressed button: %q, then %q; want Copy, then Copied, the same right edge and padding, the neighbour in its place", was, now)
	}
	if strings.Contains(at, "published") {
		t.Errorf("with the script the Copy link went to %s; want it to stay", at)
	}
}

// TestTilePick (Carrel's task 114): a tile with Send is a checkbox of its
// form — a press anywhere but the name picks it without a script, the picked
// one is the outlined one, a press again takes it back, and the name is a
// link.
func TestTilePick(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	grid := w.Grid{Tiles: []w.Tile{
		{Thumb: w.Thumb{Icon: "img"}, Name: "a.jpg", Meta: "2 MB", Selected: true, Href: "/files/a.jpg", Send: w.Send{Name: "item", Value: "a.jpg"}},
		{Thumb: w.Thumb{Icon: "img"}, Name: "b.png", Meta: "1 MB", Href: "/files/b.png", Send: w.Send{Name: "item", Value: "b.png"}}}}
	body := `<div class="tc-root"><div class="tc-main"><form id="sel" action="/x" method="post">` + grid.Render() + `</form>` +
		w.Button{Label: "Delete", Form: "sel"}.Render() + `</div></div>`
	bare := h.RegisterPage("tile-bare", strings.Replace(browser.PageHTML(body, ""), `<script src="/tc.js"></script>`, "", 1))
	const marked = `[...document.querySelectorAll(".tc-tile")].map(c => c.querySelector(":checked") ? "on" : "off").join(" ")`
	const edges = `[...document.querySelectorAll(".tc-tile")].map(c => getComputedStyle(c).borderTopColor).join("|")`
	var before, after, sent, e0, e1, href string
	if err := h.Run(bare, 900, 600,
		chromedp.Evaluate(marked, &before),
		chromedp.Evaluate(edges, &e0),
		chromedp.Click(`.tc-tile:nth-child(1) .tc-thumb`, chromedp.ByQuery),
		chromedp.Click(`.tc-tile:nth-child(2) .tc-thumb`, chromedp.ByQuery),
		chromedp.Evaluate(marked, &after),
		chromedp.Evaluate(edges, &e1),
		chromedp.Evaluate(`(() => { const r = document.querySelector(".tc-tile:nth-child(2) .tc-tile__name").getBoundingClientRect(); const a = document.elementFromPoint(r.x + 3, r.y + r.height / 2); return a.getAttribute("href") || a.tagName; })()`, &href),
		chromedp.Evaluate(`document.addEventListener("submit", e => { e.preventDefault(); window.__sent = new URLSearchParams(new FormData(e.target)).toString(); }); true`, nil),
		chromedp.Click(`.tc-btn`, chromedp.ByQuery),
		chromedp.Evaluate(`window.__sent || ""`, &sent),
	); err != nil {
		t.Fatal(err)
	}
	if before != "on off" || after != "off on" {
		t.Errorf("picked: before %q, after %q; want on off, off on", before, after)
	}
	a, b := strings.Split(e0, "|"), strings.Split(e1, "|")
	if len(a) != 2 || len(b) != 2 || a[0] == a[1] || a[0] != b[1] || a[1] != b[0] {
		t.Errorf("the outline: before %q, after %q; want it to move from the first tile to the second", e0, e1)
	}
	if href != "/files/b.png" {
		t.Errorf("under the name stands %q; want the link /files/b.png", href)
	}
	if sent != "item=b.png" {
		t.Errorf("sent %q; want item=b.png", sent)
	}
}

// TestTreeSend (Carrel's task 116): a tree's nodes with Send are radio buttons
// of the sheet's form — a press chooses a folder without a script, the chosen
// one is the marked one, an unavailable one cannot be chosen, and the sheet's
// button sends the choice.
func TestTreeSend(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	node := func(label, v string, indent w.TreeIndent, off bool) w.TreeRow {
		return w.TreeRow{Icon: "file", Label: label, Indent: indent, Unavailable: off, Send: w.Send{Name: "to", Value: v}}
	}
	sheet := w.Sheet{Title: "Move 3 items", Action: "/x", Body: []w.SheetPart{w.Tree{Rows: []w.TreeRow{
		node("Documents", "documents", w.TreeLevel0, false), node("carrel", "documents/carrel", w.TreeLevel22, false),
		node("2026-07", "documents/carrel/2026-07", w.TreeLevel40, false), node("Public", "public", w.TreeLevel0, true)}, On: 3}},
		Foot: []w.SheetFootPart{w.Button{Label: "Move here", Kind: w.ButtonPrimary, Submit: true}}}
	body := `<div class="tc-root">` + sheet.Render() + `</div>`
	bare := h.RegisterPage("tree-bare", strings.Replace(browser.PageHTML(body, ""), `<script src="/tc.js"></script>`, "", 1))
	const marked = `[...document.querySelectorAll(".tc-tree__row")].map(c => getComputedStyle(c).fontWeight === "600" ? "on" : "off").join(" ")`
	var before, after, sent string
	if err := h.Run(bare, 900, 600,
		chromedp.Evaluate(marked, &before),
		chromedp.Click(`.tc-tree__row:nth-child(4)`, chromedp.ByQuery),
		chromedp.Click(`.tc-tree__row:nth-child(2)`, chromedp.ByQuery),
		chromedp.Evaluate(marked, &after),
		chromedp.Evaluate(`document.addEventListener("submit", e => { e.preventDefault(); window.__sent = new URLSearchParams(new FormData(e.target)).toString(); }); true`, nil),
		chromedp.Click(`.tc-btn.is-primary`, chromedp.ByQuery),
		chromedp.Evaluate(`window.__sent || ""`, &sent),
	); err != nil {
		t.Fatal(err)
	}
	if before != "off off on off" || after != "off on off off" {
		t.Errorf("the chosen folder marked: before %q, after %q; want the third, then the second", before, after)
	}
	if sent != "to=documents%2Fcarrel" {
		t.Errorf("sent %q; want to=documents%%2Fcarrel", sent)
	}
}

// TestRemoveRow (Carrel's task 86): Remove of a repeated field's row empties
// the row's fields and hides them in place — they stay in the form under
// their names, the other rows keep theirs, and the labels of the first row go
// to the row that is first now. Without a script the button submits.
func TestRemoveRow(t *testing.T) {
	needChrome(t)
	row := func(n, v string, labels bool) []w.FieldSetPart {
		lab := ""
		if labels {
			lab = "Number"
		}
		return []w.FieldSetPart{w.Field{Label: lab, Parts: []w.FieldPart{w.Input{Name: "tel." + n, Value: v, Mono: true, Label: "Number"}}},
			w.Button{Label: "Remove", Kind: w.ButtonDanger, Submit: true, Name: "remove", Value: "tel." + n, Removes: true}, w.Grow{}}
	}
	items := append(row("0", "+7 912 445-10-88", true), row("1", "+7 495 120-45-00", false)...)
	set := w.FieldSet{Rubric: w.Rubric{Text: "Phone numbers"}, Items: items}
	form := w.Form{Action: "/x", Items: []w.FormPart{set}}
	body := `<div class="tc-root">` + form.Render() + `</div>`
	h := browser.NewHarness(t)
	const state = `[...document.querySelectorAll("form.tc-form input")].map(i => i.name + "=" + i.value + (i.checkVisibility() ? "" : " hidden")).join(",")`
	const sends = `document.addEventListener("submit", e => { e.preventDefault(); window.__sent = (e.submitter && e.submitter.value) || "?"; }); true`
	var got, lab, sent string
	var buttons int
	if err := h.Run(h.RegisterPage("removerow", browser.PageHTML(body, "")), 900, 600,
		chromedp.Evaluate(sends, nil),
		chromedp.Click(`[data-tc-remove]`, chromedp.ByQuery),
		chromedp.Evaluate(state, &got),
		chromedp.Evaluate(`[...document.querySelectorAll(".tc-field")].filter(f => f.checkVisibility()).map(f => (f.querySelector("label") || {textContent: "-"}).textContent).join(",")`, &lab),
		chromedp.Evaluate(`[...document.querySelectorAll("[data-tc-remove]")].filter(b => b.checkVisibility()).length`, &buttons),
		chromedp.Evaluate(`window.__sent || ""`, &sent),
	); err != nil {
		t.Fatal(err)
	}
	if got != "tel.0= hidden,tel.1=+7 495 120-45-00" {
		t.Errorf("after Remove of the first row: %s; want tel.0 empty and hidden, tel.1 as it was", got)
	}
	if lab != "Number" || buttons != 1 || sent != "" {
		t.Errorf("the row left: label %q, %d Remove, sent %q; want Number, 1, nothing sent", lab, buttons, sent)
	}
	bare := h.RegisterPage("removerow-bare", strings.Replace(browser.PageHTML(body, ""), `<script src="/tc.js"></script>`, "", 1))
	if err := h.Run(bare, 900, 600,
		chromedp.Evaluate(sends, nil),
		chromedp.Click(`[data-tc-remove]`, chromedp.ByQuery),
		chromedp.Evaluate(state, &got),
		chromedp.Evaluate(`window.__sent || ""`, &sent),
	); err != nil {
		t.Fatal(err)
	}
	if sent != "tel.0" || strings.Contains(got, "hidden") {
		t.Errorf("without a script: sent %q, fields %s; want remove=tel.0 sent and nothing hidden", sent, got)
	}
}
