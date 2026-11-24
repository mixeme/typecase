package browser_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/chromedp/cdproto/input"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"

	"carrel.dev/typecase/internal/browser"
	"carrel.dev/typecase/internal/refs"
	"carrel.dev/typecase/page"
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

// TestRadioDisables (Carrel's task 194): an option of a switch switches its
// fields off while it is chosen — «The whole calendar» and the days of a
// range —, and choosing another option of the same switch switches them on
// again, though the change comes to that other option only.
func TestRadioDisables(t *testing.T) {
	needChrome(t)
	form := func(whole bool) string {
		return w.Form{Action: "/x", Items: []w.FormPart{
			w.Field{Label: "How much of it", Parts: []w.FieldPart{w.Seg{Name: "how", Options: []w.SegOption{
				{Label: "A date range", Value: "range", On: !whole},
				{Label: "The whole calendar", Value: "whole", On: whole, Disables: "from to"}}}}},
			w.Field{Label: "From", Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "from", Label: "From"}}},
			w.Field{Label: "To", Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "to", Label: "To"}}},
		}}.Render()
	}
	days := `[...document.querySelectorAll('[name=from], [name=to]')].map(f => f.disabled).join(" ")`
	h := browser.NewHarness(t)
	rng := h.RegisterPage("range", browser.PageHTML(`<div class="tc-root"><div class="tc-main">`+form(false)+`</div></div>`, ""))
	whole := h.RegisterPage("whole", browser.PageHTML(`<div class="tc-root"><div class="tc-main">`+form(true)+`</div></div>`, ""))
	var start, off, on, startWhole string
	err := h.Run(rng, 900, 600,
		chromedp.Evaluate(days, &start),
		chromedp.Click(`.tc-seg__opt:nth-child(2)`, chromedp.ByQuery),
		chromedp.Evaluate(days, &off),
		chromedp.Click(`.tc-seg__opt:nth-child(1)`, chromedp.ByQuery),
		chromedp.Evaluate(days, &on),
	)
	if err != nil {
		t.Fatal(err)
	}
	if err := h.Run(whole, 900, 600, chromedp.Evaluate(days, &startWhole)); err != nil {
		t.Fatal(err)
	}
	if start != "false false" || off != "true true" || on != "false false" {
		t.Errorf("days off: range at start %q, whole chosen %q, range again %q; want false, true, false", start, off, on)
	}
	if startWhole != "true true" {
		t.Errorf("days off on a form that opens with the whole calendar: %q; want true", startWhole)
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
			document.addEventListener("submit", e => { e.preventDefault(); window.__sent = e.target.id === "post-import" && e.target.querySelector("input[type=file]").files.length === 1; });
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

// TestDropRefusal (Carrel's task 265, frames s-io-1 and s-io-24): the drop
// sheet refuses on its own what the field of its form does not take —
// several files at once while they are dragged, a name its accept does not
// take once dropped — in the words of that case (else of any other), sends
// nothing and stays; a file it takes goes. A dialog with an ID takes a
// dropped file into its field as one picked there, and does not send.
func TestDropRefusal(t *testing.T) {
	needChrome(t)
	post := w.PostForm{ID: "post-import", Action: "/import", File: "file", Accept: ".vcf"}
	sheet := w.Empty{Lead: "Drop to import", Variant: w.EmptyDrop, Refusals: []w.Refusal{
		{Of: ".ics", Parts: []w.Inline{w.Text("ics elsewhere")}},
		{Of: w.RefuseOther, Parts: []w.Inline{w.Text("only vcf")}},
		{Of: w.RefuseSeveral, Parts: []w.Inline{w.Text("one at a time")}}}}
	drop := w.Scrim{ID: "drop", Closed: true, Drop: "post-import", Content: sheet}
	step := w.Dialog{ID: "step-1", Action: "/restore", Items: []w.DialogPart{w.Field{Label: "Archive", Parts: []w.FieldPart{w.Line{Items: []w.LinePart{w.Button{Label: "Choose a file", File: &w.FileField{Name: "archive", Accept: ".zip", Label: "Archive"}}}}}}}}
	step2 := w.Scrim{ID: "drop2", Closed: true, Drop: "step-1", Content: w.Empty{Lead: "Drop the archive", Variant: w.EmptyDrop}}
	h := browser.NewHarness(t)
	url := h.RegisterPage("drop-refusal", browser.PageHTML(`<div class="tc-root">`+post.Render()+drop.Render()+`</div>`, ""))
	url2 := h.RegisterPage("drop-dialog", browser.PageHTML(`<div class="tc-root">`+step.Render()+step2.Render()+`</div>`, ""))
	const js = `(() => {
		const said = [];
		document.addEventListener("submit", e => { e.preventDefault(); said.push("sent " + e.target.id); });
		const sheet = document.getElementById("drop");
		const files = (...names) => { const dt = new DataTransfer(); names.forEach(n => dt.items.add(new File(["x"], n))); return dt; };
		const fire = (type, dt) => document.dispatchEvent(new DragEvent(type, {dataTransfer: dt, bubbles: true, cancelable: true}));
		const state = what => {
			const on = sheet.querySelector(".tc-empty__refusal.is-on");
			const lead = sheet.querySelector(".tc-empty__lead");
			if (!sheet.getClientRects().length) return said.push(what + ": hidden, " + (sheet.querySelector(".tc-empty.is-refused") ? "refused" : "not refused"));
			said.push(what + ": shown, " + (on ? "says " + on.textContent : "no refusal") + ", lead " + (lead.getClientRects().length > 0 ? "seen" : "unseen"));
		};
		fire("dragenter", files("a.vcf", "b.vcf")); state("two dragged");
		fire("dragover", files("a.vcf")); state("one dragged");
		fire("drop", files("a.ics")); state("an .ics dropped");
		fire("dragenter", files("b.txt")); fire("drop", files("b.txt")); state("a .txt dropped");
		fire("dragenter", files("c.VCF")); fire("drop", files("c.VCF")); state("a .VCF dropped");
		return said.join("; ");
	})()`
	const js2 = `(() => {
		let sent = false;
		document.addEventListener("submit", e => { e.preventDefault(); sent = true; });
		const dt = new DataTransfer(); dt.items.add(new File(["PK"], "backup.zip"));
		document.dispatchEvent(new DragEvent("dragenter", {dataTransfer: dt, bubbles: true, cancelable: true}));
		document.dispatchEvent(new DragEvent("drop", {dataTransfer: dt, bubbles: true, cancelable: true}));
		const f = document.querySelector("#step-1 input[type=file]");
		return (f.files.length ? f.files[0].name : "none") + (sent ? ", sent" : ", not sent");
	})()`
	var got, got2 string
	if err := h.Run(url, 900, 600, chromedp.Evaluate(js, &got)); err != nil {
		t.Fatal(err)
	}
	if err := h.Run(url2, 900, 600, chromedp.Evaluate(js2, &got2)); err != nil {
		t.Fatal(err)
	}
	want := "two dragged: shown, says one at a time, lead unseen; one dragged: shown, no refusal, lead seen; " +
		"an .ics dropped: shown, says ics elsewhere, lead unseen; a .txt dropped: shown, says only vcf, lead unseen; " +
		"sent post-import; a .VCF dropped: hidden, not refused"
	if got != want {
		t.Errorf("a drop the form does not take:\n got %s\nwant %s", got, want)
	}
	if got2 != "backup.zip, not sent" {
		t.Errorf("a file dropped for a dialog step: %s, want it in the field, not sent", got2)
	}
}

// TestSideDrop (a tail of Carrel's task 88, frame s-io-2): a file over the
// open card's panel marks the panel and keeps the page's drop sheet away;
// dropped there, it goes to the panel's form, not the sheet's. Off the panel
// the sheet is the target, and over the sheet the panel under it still wins.
func TestSideDrop(t *testing.T) {
	needChrome(t)
	side := w.Side{Drop: "post-attach", Items: []w.SidePart{w.Head{Title: &w.HeadTitle{Text: "Синк-ап по релизу"}}}}
	body := `<div class="tc-root"><div style="display:flex"><div id="away" style="flex:1;height:400px"></div>` + side.Render() + `</div>` +
		w.PostForm{ID: "post-attach", Action: "/attach", File: "file"}.Render() + w.PostForm{ID: "post-import", Action: "/import", File: "file"}.Render() +
		w.Scrim{ID: "drop", Closed: true, Drop: "post-import", Content: w.Sheet{Title: "Drop to import"}}.Render() + `</div>`
	h := browser.NewHarness(t)
	url := h.RegisterPage("side-drop", browser.PageHTML(body, ""))
	var got string
	err := h.Run(url, 900, 600,
		chromedp.Evaluate(`(() => {
			const said = [];
			document.addEventListener("submit", e => { e.preventDefault(); said.push("sent " + e.target.id); });
			const at = el => { const r = el.getBoundingClientRect(); return [r.left + r.width / 2, r.top + r.height / 2]; };
			const panel = document.querySelector(".tc-side"), sheet = document.getElementById("drop");
			const [px, py] = at(panel), [ax, ay] = at(document.getElementById("away"));
			const drag = (type, x, y) => {
				const dt = new DataTransfer();
				dt.items.add(new File(["x"], "release-checklist.md"));
				document.elementFromPoint(x, y).dispatchEvent(new DragEvent(type, {dataTransfer: dt, clientX: x, clientY: y, bubbles: true, cancelable: true}));
			};
			const state = what => said.push(what + ": panel " + panel.classList.contains("tc-side--drop") + ", sheet " + (sheet.getClientRects().length > 0));
			drag("dragenter", px, py); state("over the panel");
			drag("dragover", ax, ay); state("off it");
			drag("dragover", px, py); state("back over it");
			drag("drop", px, py); state("dropped on it");
			drag("dragenter", ax, ay); drag("drop", ax, ay); state("dropped off it");
			return said.join("; ");
		})()`, &got),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := "over the panel: panel true, sheet false; off it: panel false, sheet true; back over it: panel true, sheet false; " +
		"sent post-attach; dropped on it: panel false, sheet false; sent post-import; dropped off it: panel false, sheet false"
	if got != want {
		t.Errorf("a file over the card's panel:\n got %s\nwant %s", got, want)
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

// TestBarFormSubmits (Carrel's task 195): the journal's action and day are a
// GET form — without a script its Show sends them with the hidden segment,
// an empty action too; with the library's script choosing an action in the
// menu does, and Show is not shown.
func TestBarFormSubmits(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	form := func(action string) string {
		f := w.BarForm{Items: []w.BarFormPart{
			w.Drop{Trigger: w.Pick{Value: "All actions"}, Menu: w.Menu{Items: []w.MenuPart{
				w.MenuItem{Label: "All actions", On: true, Name: "action", Value: ""},
				w.MenuItem{Label: "login.fail", Name: "action", Value: "login.fail"}}, Pop: w.MenuUnder}},
			w.Dim{Parts: []w.Inline{w.Text("since")}},
			w.Input{Kind: w.InputDate, Name: "since", Value: "2026-08-01", Label: "Since"}},
			Action: action, Hidden: []w.Hidden{{Name: "show", Value: "failures"}}, Show: w.Button{Label: "Show"}}
		return `<div class="tc-root"><div class="tc-main"><div class="tc-bar">` + f.Render() + `</div></div></div>`
	}
	bareURL := h.RegisterPage("barform-bare", "")
	h.RegisterPage("barform-bare", strings.Replace(browser.PageHTML(form(bareURL), ""), `<script src="/tc.js"></script>`, "", 1))
	var showShown bool
	var search string
	if err := h.Run(bareURL, 900, 600,
		chromedp.Evaluate(`document.querySelector(".tc-barform__show button").checkVisibility()`, &showShown),
		chromedp.Click(`.tc-barform__show button`, chromedp.ByQuery),
		chromedp.Sleep(300*time.Millisecond),
		chromedp.Evaluate(`location.search`, &search),
	); err != nil {
		t.Fatal(err)
	}
	if !showShown || !strings.Contains(search, "action=&") || !strings.Contains(search, "since=2026-08-01") || !strings.Contains(search, "show=failures") {
		t.Errorf("without a script: Show shown %v, sent %q; want true and action=, since=2026-08-01, show=failures", showShown, search)
	}
	pageURL := h.RegisterPage("barform", "")
	h.RegisterPage("barform", browser.PageHTML(form(pageURL), ""))
	var hidden bool
	search = ""
	if err := h.Run(pageURL, 900, 600,
		chromedp.Evaluate(`!document.querySelector(".tc-barform__show button").checkVisibility()`, &hidden),
		chromedp.Click(`.tc-barform details.tc-drop > summary`, chromedp.ByQuery),
		chromedp.Click(`input[value="login.fail"]`, chromedp.ByQuery),
		chromedp.Sleep(300*time.Millisecond),
		chromedp.Evaluate(`location.search`, &search),
	); err != nil {
		t.Fatal(err)
	}
	if !hidden || !strings.Contains(search, "action=login.fail") || !strings.Contains(search, "show=failures") {
		t.Errorf("with the script: no Show %v, sent %q; want true and action=login.fail, show=failures", hidden, search)
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
		// a switch that does not fit at 380: it goes under «⋯» by the CSS's
		// width without a script and by measure with one
		seg := `<span class="tc-seg is-2nd">`
		for _, o := range []string{"A–Z", "Z–A", "Newest first", "Oldest first", "Largest", "Smallest"} {
			seg += `<button class="tc-seg__opt" type="button">` + o + `</button>`
		}
		return `<div class="tc-root"><div class="tc-main"><div class="tc-bar">` + seg + `</span>` + more.Render() + `</div></div></div>`
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
			Row: &w.Line{Place: w.Grown, Gap: w.LineGapNormal, Items: []w.LinePart{w.Text("{name} "), w.Dim{Parts: []w.Inline{w.Text("{size}")}},
				w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Place: w.Pushed, Removes: true}}}},
		w.Line{Block: true, Gap: w.LineGapNormal, Place: w.Grown, ID: "picked"}}}

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
	row := w.Line{Gap: w.LineGapNormal, Items: []w.LinePart{w.Faced{Face: w.FaceLarge, Items: []w.LinePart{w.Text("{name} "), w.Dim{Parts: []w.Inline{w.Text("{size}")}}}},
		w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Removes: true}}}
	attach := w.Drop{Trigger: w.Button{Label: "Attach", Menu: true}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "up", Label: "From this computer", File: &w.FileField{Name: "file", Label: "From this computer", List: "files"}, Row: row}}}}
	list := w.Line{ID: "files", Gap: w.LineGapNormal, Items: []w.LinePart{attach, w.Dim{Parts: []w.Inline{w.Text("Nothing attached")}, Empty: true}}}
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
		Foot: []w.SheetFootPart{w.Button{Label: "Move", Kind: w.ButtonPrimary, Submit: true}}}
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

// TestEndFollowsStart (Carrel's task 152): an end named as following its
// start (Input.Follows) moves by as much as the start was moved — the length
// stays. A time and its day are one pair (Input.Day): an end carried over
// midnight takes its date along. An empty end stays empty, a time switched
// off by «All day» is not touched, an end set by hand follows with its new
// length. Without a script nothing moves: what is on the screen is sent.
func TestEndFollowsStart(t *testing.T) {
	needChrome(t)
	in := func(kind w.InputKind, name, value, follows, day string) w.Field {
		return w.Field{Label: name, Parts: []w.FieldPart{w.Input{Kind: kind, Name: name, Value: value, Mono: true, Follows: follows, Day: day}}}
	}
	form := func(sd, st, ed, et string) string {
		f := w.Form{Action: "/x", Items: []w.FormPart{w.FieldSet{Rubric: w.Rubric{Text: "When"}, Items: []w.FieldSetPart{
			w.Field{Parts: []w.FieldPart{w.Check{Name: "start.allday", Value: "1", Parts: []w.CheckPart{w.Text("All day")}, Disables: "start.time end.time"}}},
			in(w.InputDate, "start.date", sd, "", ""), in(w.InputTime, "start.time", st, "", ""),
			in(w.InputDate, "end.date", ed, "start.date", ""), in(w.InputTime, "end.time", et, "start.time", "end.date"),
		}}}}
		return `<div class="tc-root">` + f.Render() + `</div>`
	}
	// set types a value into a field as a person does: the value and a change
	const set = `(n, v) => { const f = document.querySelector('[name="' + n + '"]'); f.value = v; f.dispatchEvent(new Event("change", {bubbles: true})); }`
	const state = `["start.date", "start.time", "end.date", "end.time"].map(n => document.querySelector('[name="' + n + '"]').value).join(" ")`
	h := browser.NewHarness(t)
	for _, c := range []struct {
		name           string
		sd, st, ed, et string
		steps          string
		want           string
		bare           bool
	}{
		{"the start half an hour later", "2026-08-10", "10:00", "2026-08-10", "11:00",
			`set("start.time", "10:30")`, "2026-08-10 10:30 2026-08-10 11:30", false},
		{"the start a day later", "2026-08-10", "10:00", "2026-08-10", "11:00",
			`set("start.date", "2026-08-11")`, "2026-08-11 10:00 2026-08-11 11:00", false},
		{"an end past midnight comes back to the start's day", "2026-08-10", "23:30", "2026-08-11", "00:30",
			`set("start.time", "22:00")`, "2026-08-10 22:00 2026-08-10 23:00", false},
		{"an end goes over midnight with its date", "2026-08-31", "22:00", "2026-08-31", "23:00",
			`set("start.time", "23:30")`, "2026-08-31 23:30 2026-09-01 00:30", false},
		{"an empty end stays empty", "2026-08-10", "10:00", "", "",
			`set("start.time", "10:30"); set("start.date", "2026-08-12")`, "2026-08-12 10:30  ", false},
		{"an end set by hand follows with its new length", "2026-08-10", "10:00", "2026-08-10", "11:00",
			`set("end.time", "12:00"); set("start.time", "11:00")`, "2026-08-10 11:00 2026-08-10 13:00", false},
		{"two moves in a row", "2026-08-10", "10:00", "2026-08-10", "11:00",
			`set("start.time", "10:30"); set("start.time", "09:00")`, "2026-08-10 09:00 2026-08-10 10:00", false},
		{"all day: the dates go, the times are not touched", "2026-08-10", "10:00", "2026-08-11", "11:00",
			`document.querySelector('[name="start.allday"]').click(); set("start.date", "2026-08-12")`, "2026-08-12 10:00 2026-08-13 11:00", false},
		{"without a script nothing moves", "2026-08-10", "10:00", "2026-08-10", "11:00",
			`set("start.time", "10:30")`, "2026-08-10 10:30 2026-08-10 11:00", true},
	} {
		page := browser.PageHTML(form(c.sd, c.st, c.ed, c.et), "")
		if c.bare {
			page = strings.Replace(page, `<script src="/tc.js"></script>`, "", 1)
		}
		var got string
		if err := h.Run(h.RegisterPage("follows", page), 900, 600,
			chromedp.Evaluate(`(() => { const set = `+set+`; `+c.steps+`; return true })()`, nil),
			chromedp.Evaluate(state, &got),
		); err != nil {
			t.Fatal(err)
		}
		if got != c.want {
			t.Errorf("%s: %s; want %s", c.name, got, c.want)
		}
	}
}

// TestFieldChanged: a field brought by a change (Field.Changed; the owner,
// 16.10) shows while a named field differs from what the page brought, goes
// again when the edit is undone, sends nothing while hidden, and is not
// touched when the page brought it shown (the screen asked after Save); a
// check its page brought off with a Reason stays off when it shows (review
// finding 6).
func TestFieldChanged(t *testing.T) {
	needChrome(t)
	form := func(hidden bool) string {
		f := w.Form{Action: "/x", Items: []w.FormPart{
			w.Field{Label: "Server URL", Parts: []w.FieldPart{w.Input{Kind: w.InputText, Name: "server.url", Value: "https://dav.home.example/"}}},
			w.Field{Label: "Username", Parts: []w.FieldPart{w.Input{Kind: w.InputText, Name: "username", Value: "mix"}}},
			w.Field{Label: "Label", Parts: []w.FieldPart{w.Input{Kind: w.InputText, Name: "label", Value: ""}}},
			w.Field{Label: "Published links", Changed: []string{"server.url", "username"}, Hidden: hidden, Parts: []w.FieldPart{
				w.Check{Name: "published.links", Value: "keep", Radio: true, Parts: []w.CheckPart{w.Text("Keep them")}},
				w.Check{Name: "published.links", Value: "revoke", Radio: true, Parts: []w.CheckPart{w.Text("Revoke them")}},
				w.Check{Name: "published.notify", Value: "yes", Reason: "No mail relay", Parts: []w.CheckPart{w.Text("Tell their readers")}},
			}},
		}}
		return `<div class="tc-root">` + f.Render() + `</div>`
	}
	const typ = `(n, v) => { const f = document.querySelector('[name="' + n + '"]'); f.value = v; f.dispatchEvent(new Event("input", {bubbles: true})); }`
	// state: the field shown or hidden, its question on or off, the check
	// with a Reason off whatever the field is
	const state = `(() => { const f = document.querySelector("[data-tc-changed]"); const b = f.querySelector("input"); const r = f.querySelector('[name="published.notify"]');
		return (f.hidden ? "hidden" : "shown") + (b.disabled ? " off" : " on") + (r.disabled ? "" : " reason-on"); })()`
	h := browser.NewHarness(t)
	for _, c := range []struct {
		name   string
		hidden bool
		steps  string
		want   string
	}{
		{"untouched: hidden, not sent", true, ``, "hidden off"},
		{"another label asks nothing", true, `typ("label", "Home")`, "hidden off"},
		{"another address asks", true, `typ("server.url", "https://cloud.home.example/")`, "shown on"},
		{"another login asks", true, `typ("username", "anna")`, "shown on"},
		{"the address back: the question goes", true, `typ("server.url", "https://cloud.home.example/"); typ("server.url", "https://dav.home.example/")`, "hidden off"},
		{"brought shown after Save: left alone", false, `typ("server.url", "https://dav.home.example/x")`, "shown on"},
	} {
		var got string
		if err := h.Run(h.RegisterPage("changed", browser.PageHTML(form(c.hidden), "")), 900, 600,
			chromedp.Evaluate(`(() => { const typ = `+typ+`; `+c.steps+`; return true })()`, nil),
			chromedp.Evaluate(state, &got),
		); err != nil {
			t.Fatal(err)
		}
		if got != c.want {
			t.Errorf("%s: %s; want %s", c.name, got, c.want)
		}
	}
}

// TestThumbPictureGone: a thumbnail's picture that did not arrive goes, and
// the icon under it is what the tile shows; one that arrived hides the icon.
// Without a script the icon stays under the picture either way. The bite: a
// page without tc.js keeps the broken picture.
func TestThumbPictureGone(t *testing.T) {
	needChrome(t)
	const ok = `data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='40' height='30'%3E%3Crect width='40' height='30' fill='%239FB4C7'/%3E%3C/svg%3E`
	body := `<div class="tc-root"><div id="a">` + w.Thumb{Icon: "img", Src: "/no-such-thumbnail.png", Alt: "a.png"}.Render() +
		`</div><div id="b">` + w.Thumb{Icon: "img", Src: ok, Alt: "b.png"}.Render() + `</div></div>`
	// state says, for each thumbnail, whether its picture and its icon are drawn.
	const state = `(async () => {
		await Promise.all([...document.images].map(i => i.complete ? 0 : new Promise(r => { i.addEventListener("load", r); i.addEventListener("error", r) })));
		const drawn = el => el.getClientRects().length > 0 ? "shown" : "gone";
		return ["a", "b"].map(id => id + ": picture " + drawn(document.querySelector("#" + id + " img")) + ", icon " + drawn(document.querySelector("#" + id + " svg"))).join("; ");
	})()`
	h := browser.NewHarness(t)
	for _, c := range []struct {
		name string
		bare bool
		want string
	}{
		{"with the script", false, "a: picture gone, icon shown; b: picture shown, icon gone"},
		{"without a script", true, "a: picture shown, icon shown; b: picture shown, icon shown"},
	} {
		page := browser.PageHTML(body, "")
		if c.bare {
			page = strings.Replace(page, `<script src="/tc.js"></script>`, "", 1)
		}
		var got string
		if err := h.Run(h.RegisterPage("thumb", page), 900, 600,
			chromedp.Evaluate(state, &got, func(p *runtime.EvaluateParams) *runtime.EvaluateParams { return p.WithAwaitPromise(true) }),
		); err != nil {
			t.Fatal(err)
		}
		if got != c.want {
			t.Errorf("%s: %s; want %s", c.name, got, c.want)
		}
	}
}

// TestSelectionForm (Carrel's task 148): the ticks of a table and of tiles
// belong to a form they do not stand in (Send.Of), and the selection bar's
// buttons send that form — the ticked names and which action. The bar of an
// empty selection (Bar.NoScript) stands on a page without a script and is not
// drawn on one with it. The bite: the same page with tc.js has no bar.
func TestSelectionForm(t *testing.T) {
	needChrome(t)
	off := w.TickedOff
	bar := w.Bar{Variant: w.BarSelection, NoScript: true, Count: &w.BarCount{Text: "With the ticked"},
		Items: []w.BarPart{w.Button{Label: "Delete", Form: "sel", Name: "op", Value: "delete"}}}
	table := w.Table{Rows: []w.TableRow{{Cells: []w.TableCell{{Pick: &off, Send: w.Send{Of: "sel", Name: "item", Value: "t.pdf"}}, {Parts: []w.CellPart{w.Text("t.pdf")}}}}}}
	grid := w.Grid{Tiles: []w.Tile{
		{Thumb: w.Thumb{Icon: "img"}, Name: "a.jpg", Meta: "2 MB", Send: w.Send{Of: "sel", Name: "item", Value: "a.jpg"}},
		{Thumb: w.Thumb{Icon: "img"}, Name: "b.png", Meta: "1 MB", Send: w.Send{Of: "sel", Name: "item", Value: "b.png"}}}}
	body := `<div class="tc-root"><div class="tc-main">` + bar.Render() + table.Render() + grid.Render() +
		w.PostForm{ID: "sel", Action: "/x", Fields: []w.Hidden{{Name: "csrf", Value: "k"}}}.Render() + `</div></div>`
	h := browser.NewHarness(t)
	const shown = `document.querySelector(".tc-bar").getClientRects().length > 0 ? "bar" : "no bar"`
	var with, without, sent string
	if err := h.Run(h.RegisterPage("sel", browser.PageHTML(body, "")), 900, 600, chromedp.Evaluate(shown, &with)); err != nil {
		t.Fatal(err)
	}
	bare := h.RegisterPage("sel-bare", strings.Replace(browser.PageHTML(body, ""), `<script src="/tc.js"></script>`, "", 1))
	if err := h.Run(bare, 900, 600,
		chromedp.Evaluate(shown, &without),
		chromedp.Click(`.tc-table input`, chromedp.ByQuery),
		chromedp.Click(`.tc-tile:nth-child(2) .tc-thumb`, chromedp.ByQuery),
		chromedp.Evaluate(`document.addEventListener("submit", e => { e.preventDefault(); window.__sent = new URLSearchParams(new FormData(e.target, e.submitter)).toString(); }); true`, nil),
		chromedp.Click(`.tc-bar .tc-btn`, chromedp.ByQuery),
		chromedp.Evaluate(`window.__sent || ""`, &sent),
	); err != nil {
		t.Fatal(err)
	}
	if with != "no bar" || without != "bar" {
		t.Errorf("the bar of an empty selection: with a script %q, without %q; want no bar, bar", with, without)
	}
	if sent != "op=delete&item=t.pdf&item=b.png&csrf=k" {
		t.Errorf("sent %q; want op=delete&item=t.pdf&item=b.png&csrf=k", sent)
	}
}

// TestTreeNodeSwap (Carrel's task 158): a tree's node that is a radio button
// and has an address — choosing it asks the page drawn with that node chosen
// and puts the sheet it names in place: the button that was off is on, the
// page is not loaded again. The bite: without the live layer the choice
// moves and the button stays as the server drew it.
func TestTreeNodeSwap(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	page := func(a, b string, on int) string {
		move := w.Button{Label: "Move", Kind: w.ButtonPrimary, Submit: true}
		if on == 1 {
			move.State = w.ButtonDisabled
		}
		sheet := w.Sheet{Title: "Move 3 items", Action: "/x", ID: "move-sheet", Body: []w.SheetPart{w.Tree{Rows: []w.TreeRow{
			{Icon: "file", Label: "here", Send: w.Send{Name: "to", Value: "here"}, Href: a, Swap: "move-sheet"},
			{Icon: "file", Label: "there", Send: w.Send{Name: "to", Value: "there"}, Href: b, Swap: "move-sheet"}}, On: on}},
			Foot: []w.SheetFootPart{move}}
		return `<div class="tc-root">` + sheet.Render() + `</div>`
	}
	here := h.RegisterPage("node-here", "")
	there := h.RegisterPage("node-there", "")
	h.RegisterPage("node-here", browser.PageHTML(page(here, there, 1), ""))
	h.RegisterPage("node-there", browser.PageHTML(page(here, there, 2), ""))
	bare := h.RegisterPage("node-bare", strings.Replace(browser.PageHTML(page(here, there, 1), ""), `<script src="/tc-live.js"></script>`, "", 1))
	const state = `(document.querySelector(".tc-btn.is-primary").disabled ? "off" : "on") + " " + document.querySelector("input:checked").value + " " + (window.__kept === true)`
	for _, c := range []struct{ name, url, want string }{
		{"with the live layer", here, "on there true"},
		{"without it", bare, "off there true"},
	} {
		var got string
		if err := h.Run(c.url, 900, 600,
			chromedp.Evaluate(`window.__kept = true; true`, nil),
			chromedp.Click(`.tc-tree__row:nth-child(2)`, chromedp.ByQuery),
			chromedp.Sleep(1500*time.Millisecond),
			chromedp.Evaluate(state, &got),
		); err != nil {
			t.Fatal(err)
		}
		if got != c.want {
			t.Errorf("%s: %s; want %s", c.name, got, c.want)
		}
	}
}

// TestFormRecount (Carrel's task 194): a form asked again as it changes —
// a day of the calendar to attach changes, the live layer asks the form's
// address with its fields and puts the line at Attach in place; a day that
// is off is not asked with. The bite: without the live layer the line stays
// as the server drew it.
func TestFormRecount(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	asked := h.RegisterStream("recount", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		q := r.URL.Query()
		_, _ = io.WriteString(rw, `<span id="attach-line">`+q.Get("what")+` `+q.Get("span")+` `+q.Get("from")+` `+q.Get("to")+`</span>`)
	})
	page := func(live bool) string {
		form := w.Form{Action: asked + "?what=calendar", ID: "attach-calendar", Recount: "attach-line", Items: []w.FormPart{
			w.Field{Label: "What goes in", Parts: []w.FieldPart{w.Seg{Name: "span", Options: []w.SegOption{
				{Label: "Whole calendar", Value: "whole", Disables: "from to"}, {Label: "Dates", Value: "dates", On: true}}}}},
			w.Field{Label: "From", Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From"}}},
			w.Field{Label: "To", Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-31", Label: "To"}}},
		}}
		line := w.Dim{Parts: []w.Inline{w.Text("as drawn")}, ID: "attach-line"}
		doc := browser.PageHTML(`<div class="tc-root">`+form.Render()+line.Render()+`</div>`, "")
		if !live {
			doc = strings.Replace(doc, `<script src="/tc-live.js"></script>`, "", 1)
		}
		return doc
	}
	const set = `(() => { const f = document.querySelector('[name="to"]'); f.value = "2026-09-30"; f.dispatchEvent(new Event("change", {bubbles: true})); return true; })()`
	const whole = `(() => { document.querySelector('[value="whole"]').click(); return true; })()`
	const line = `document.getElementById("attach-line").textContent`
	var days, wholeLine, bare string
	if err := h.Run(h.RegisterPage("recount-live", page(true)), 900, 600,
		chromedp.Evaluate(set, nil),
		chromedp.Sleep(1500*time.Millisecond),
		chromedp.Evaluate(line, &days),
		chromedp.Evaluate(whole, nil),
		chromedp.Sleep(1500*time.Millisecond),
		chromedp.Evaluate(line, &wholeLine),
	); err != nil {
		t.Fatal(err)
	}
	if err := h.Run(h.RegisterPage("recount-bare", page(false)), 900, 600,
		chromedp.Evaluate(set, nil),
		chromedp.Sleep(1500*time.Millisecond),
		chromedp.Evaluate(line, &bare),
	); err != nil {
		t.Fatal(err)
	}
	if days != "calendar dates 2026-08-10 2026-09-30" {
		t.Errorf("the line after To changed: %q; want calendar dates 2026-08-10 2026-09-30", days)
	}
	if wholeLine != "calendar whole  " {
		t.Errorf("the line after the whole calendar: %q; want the days not asked with", wholeLine)
	}
	if bare != "as drawn" {
		t.Errorf("without the live layer: %q; want as drawn", bare)
	}
}

// TestFilterNarrows (the owner, 07.11): a filter narrows what the page
// shows as one types — each pause asks the filter's address with the typed
// word, the live layer puts the main in place, the line keeps the word, the
// caret and the focus, and the page's address carries the filter; an empty
// line asks the bare address and the page shows everything again. The
// bite: without the live layer typing changes nothing until Enter.
func TestFilterNarrows(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	titles := []string{"Синк-ап по релизу", "Забрать документы", "Поездка в Казань"}
	var asked string
	main := func(q string) string {
		f := w.Filter{Query: w.Query{Name: "filter", Value: q, Label: "Filter by word"}, Waiting: q == "", Variant: w.FilterH26W200, Action: asked, Swap: "list-main"}
		if q != "" {
			f.Clear = asked
		}
		s := `<div class="tc-main" id="list-main">` + f.Render()
		for _, r := range titles {
			if q == "" || strings.Contains(strings.ToLower(r), strings.ToLower(q)) {
				s += `<div class="tc-row"><span class="tc-row__title">` + r + `</span></div>`
			}
		}
		return s + `</div>`
	}
	asked = h.RegisterStream("narrow", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = io.WriteString(rw, main(r.URL.Query().Get("filter")))
	})
	page := func(live bool) string {
		doc := browser.PageHTML(`<div class="tc-root">`+main("")+`</div>`, "")
		if !live {
			doc = strings.Replace(doc, `<script src="/tc-live.js"></script>`, "", 1)
		}
		return doc
	}
	const typeIn = `(v => { const q = document.querySelector('[name="filter"]'); q.focus(); q.value = v; q.setSelectionRange(v.length, v.length); q.dispatchEvent(new Event("input", {bubbles: true})); return true; })`
	const state = `(() => { const q = document.querySelector('[name="filter"]'); return Array.from(document.querySelectorAll(".tc-row__title"), x => x.textContent).join(", ") + " | " + q.value + " | " + (document.activeElement === q) + " " + q.selectionStart + " | " + decodeURIComponent(location.search) + " | " + !!document.querySelector(".tc-in__clear"); })()`
	var narrowed, widened, bare string
	if err := h.Run(h.RegisterPage("narrow-live", page(true)), 900, 600,
		chromedp.Evaluate(typeIn+`("син")`, nil),
		chromedp.Sleep(1500*time.Millisecond),
		chromedp.Evaluate(state, &narrowed),
		chromedp.Evaluate(typeIn+`("")`, nil),
		chromedp.Sleep(1500*time.Millisecond),
		chromedp.Evaluate(state, &widened),
	); err != nil {
		t.Fatal(err)
	}
	if err := h.Run(h.RegisterPage("narrow-bare", page(false)), 900, 600,
		chromedp.Evaluate(typeIn+`("син")`, nil),
		chromedp.Sleep(1500*time.Millisecond),
		chromedp.Evaluate(state, &bare),
	); err != nil {
		t.Fatal(err)
	}
	if want := "Синк-ап по релизу | син | true 3 | ?filter=син | true"; narrowed != want {
		t.Errorf("after typing: %q; want %q", narrowed, want)
	}
	if want := "Синк-ап по релизу, Забрать документы, Поездка в Казань |  | true 0 |  | false"; widened != want {
		t.Errorf("after emptying the line: %q; want %q", widened, want)
	}
	if want := "Синк-ап по релизу, Забрать документы, Поездка в Казань | син | true 3 |  | false"; bare != want {
		t.Errorf("without the live layer: %q; want %q", bare, want)
	}
}

// TestPickDisables (Carrel's task 218): an option of a pick switches fields
// of its form off while it is chosen — «None — up to 16 MB» and the folder
// for a large archive; another option switches it on, and an off field is
// not sent. The bite: without the library's script the field stays as the
// server printed it.
func TestPickDisables(t *testing.T) {
	needChrome(t)
	form := w.Form{Action: "/x", Items: []w.FormPart{
		w.Field{Label: "Storage", Parts: []w.FieldPart{w.Pick{Name: "store", Label: "Storage", Options: []w.PickOption{
			{Value: "", Label: "None — up to 16 MB", On: true, Disables: "folder"}, {Value: "nextcloud|files", Label: "nextcloud · Files"}}}}},
		w.Field{Label: "Folder", Parts: []w.FieldPart{w.Input{Name: "folder", Value: "/carrel/restore", Label: "Folder"}}},
	}}
	body := `<div class="tc-root">` + form.Render() + `</div>`
	const choose = `v => { const s = document.querySelector('[name="store"]'); s.value = v; s.dispatchEvent(new Event("change", {bubbles: true})); }`
	const state = `(document.querySelector('[name="folder"]').disabled ? "off" : "on") + " | " + new URLSearchParams(new FormData(document.querySelector("form"))).toString()`
	h := browser.NewHarness(t)
	var start, chosen, back, bare string
	if err := h.Run(h.RegisterPage("pick-disables", browser.PageHTML(body, "")), 900, 600,
		chromedp.Evaluate(state, &start),
		chromedp.Evaluate(`(`+choose+`)("nextcloud|files"); true`, nil),
		chromedp.Evaluate(state, &chosen),
		chromedp.Evaluate(`(`+choose+`)(""); true`, nil),
		chromedp.Evaluate(state, &back),
	); err != nil {
		t.Fatal(err)
	}
	if err := h.Run(h.RegisterPage("pick-disables-bare", strings.Replace(browser.PageHTML(body, ""), `<script src="/tc.js"></script>`, "", 1)), 900, 600,
		chromedp.Evaluate(`(`+choose+`)("nextcloud|files"); true`, nil),
		chromedp.Evaluate(state, &bare),
	); err != nil {
		t.Fatal(err)
	}
	if start != "off | store=" {
		t.Errorf("as printed: %q; want the folder off and not sent", start)
	}
	if chosen != "on | store=nextcloud%7Cfiles&folder=%2Fcarrel%2Frestore" {
		t.Errorf("a storage chosen: %q; want the folder on and sent", chosen)
	}
	if back != "off | store=" {
		t.Errorf("None again: %q; want the folder off", back)
	}
	if bare != "on | store=nextcloud%7Cfiles&folder=%2Fcarrel%2Frestore" {
		t.Errorf("without the script: %q; want the field as printed (on)", bare)
	}
}

// TestPickShows (Carrel's task 156): an option of a pick shows fields of its
// form while it is chosen — «Custom…» of a reminder opens the number and the
// unit; another option hides them again, and a hidden field is not sent.
// Without a script the fields stay as the server printed them. The bite: the
// page without tc.js keeps them hidden after the choice.
func TestPickShows(t *testing.T) {
	needChrome(t)
	form := w.Form{Action: "/x", Items: []w.FormPart{w.FieldSet{Rubric: w.Rubric{Text: "Reminders"}, Items: []w.FieldSetPart{
		w.Field{Label: "First", Parts: []w.FieldPart{w.Pick{Name: "reminder", Label: "First", Options: []w.PickOption{
			{Value: "10", Label: "10 minutes before", On: true}, {Value: "custom", Label: "Custom…", Shows: "count unit"}}}}},
		w.Field{Label: "How long", Hidden: true, Parts: []w.FieldPart{w.Input{Name: "count", Value: "2", Mono: true}}},
		w.Field{Label: "Unit", Hidden: true, Parts: []w.FieldPart{w.Pick{Name: "unit", Label: "Unit", Options: []w.PickOption{{Value: "days", Label: "days before", On: true}}}}},
	}}}}
	body := `<div class="tc-root">` + form.Render() + `</div>`
	const choose = `v => { const s = document.querySelector('[name="reminder"]'); s.value = v; s.dispatchEvent(new Event("change", {bubbles: true})); }`
	const state = `(() => { const shown = n => document.querySelector('[name="' + n + '"]').getClientRects().length > 0 ? "shown" : "hidden";
		return shown("count") + " " + shown("unit") + " | " + new URLSearchParams(new FormData(document.querySelector("form"))).toString(); })()`
	h := browser.NewHarness(t)
	for _, c := range []struct {
		name, steps, want string
		bare              bool
	}{
		{"as printed", ``, "hidden hidden | reminder=10", false},
		{"Custom… chosen", `choose("custom")`, "shown shown | reminder=custom&count=2&unit=days", false},
		{"and a step again", `choose("custom"); choose("10")`, "hidden hidden | reminder=10", false},
		{"without a script", `choose("custom")`, "hidden hidden | reminder=custom&count=2&unit=days", true},
	} {
		page := browser.PageHTML(body, "")
		if c.bare {
			page = strings.Replace(page, `<script src="/tc.js"></script>`, "", 1)
		}
		var got string
		if err := h.Run(h.RegisterPage("shows", page), 900, 600,
			chromedp.Evaluate(`(() => { const choose = `+choose+`; `+c.steps+`; return true })()`, nil),
			chromedp.Evaluate(state, &got),
		); err != nil {
			t.Fatal(err)
		}
		if got != c.want {
			t.Errorf("%s: %s; want %s", c.name, got, c.want)
		}
	}
}

// TestTemplateShows (Carrel's task 161): a row added from a template
// numbers what its options show as it numbers its fields — «Custom…» of an
// added reminder opens that reminder's number, the number is sent under the
// row's name, and as added the hidden field is not sent. The bite: with the
// names of data-tc-shows left as «{n}», the choice opens nothing.
func TestTemplateShows(t *testing.T) {
	needChrome(t)
	row := []w.FieldSetPart{
		w.Field{Label: "Reminder", Parts: []w.FieldPart{w.Pick{Name: "reminder.{n}", Label: "Reminder", Options: []w.PickOption{
			{Value: "10", Label: "10 minutes before", On: true}, {Value: "custom", Label: "Custom…", Shows: "reminder.{n}.count"}}}}},
		w.Field{Label: "How long", Hidden: true, Parts: []w.FieldPart{w.Input{Name: "reminder.{n}.count", Value: "2", Mono: true}}},
	}
	form := w.Form{Action: "/x", Items: []w.FormPart{w.FieldSet{Rubric: w.Rubric{Text: "Reminders"}, Items: []w.FieldSetPart{
		w.Template{ID: "tpl-rem", Items: row},
		w.Button{Label: "Add reminder", Small: true, Submit: true, Name: "add", Value: "reminder", Adds: "tpl-rem"},
	}}}}
	body := `<div class="tc-root">` + form.Render() + `</div>`
	const add = `document.querySelector('[data-tc-adds]') ? document.querySelector('[data-tc-adds]').click() : document.querySelector('button[name="add"]').click()`
	const choose = `v => { const s = document.querySelector('[name="reminder.0"]'); s.value = v; s.dispatchEvent(new Event("change", {bubbles: true})); }`
	const state = `(() => { const f = document.querySelector('[name="reminder.0.count"]');
		return (f && f.getClientRects().length > 0 ? "shown" : "hidden") + " | " + new URLSearchParams(new FormData(document.querySelector("form"))).toString(); })()`
	h := browser.NewHarness(t)
	for _, c := range []struct{ name, steps, want string }{
		{"added", add, "hidden | reminder.0=10"},
		{"Custom… chosen", add + `; choose("custom")`, "shown | reminder.0=custom&reminder.0.count=2"},
		{"and a step again", add + `; choose("custom"); choose("10")`, "hidden | reminder.0=10"},
	} {
		var got string
		if err := h.Run(h.RegisterPage("tplshows", browser.PageHTML(body, "")), 900, 600,
			chromedp.Evaluate(`(() => { const choose = `+choose+`; `+c.steps+`; return true })()`, nil),
			chromedp.Evaluate(state, &got),
		); err != nil {
			t.Fatal(err)
		}
		if got != c.want {
			t.Errorf("%s: %s; want %s", c.name, got, c.want)
		}
	}
}

// TestTallyCount (the owner's decision of 14.10): the button of the sheet
// «Rewrite the addresses» says how many records its choices rewrite, at each
// tick and each choice of an address — a record of two groups only while
// both are ticked, an address left as it is takes its records away, none
// turns the button off, one is said in the singular. Without a script the
// label stays as printed. The bite: without the tally the label never moves.
func TestTallyCount(t *testing.T) {
	needChrome(t)
	var sheet w.Sheet
	for _, e := range refs.Entries {
		if s, ok := e.Fixture.(w.Sheet); ok && e.Ref.Frame == "s-contacts-29" {
			sheet = s
		}
	}
	if sheet.Title == "" {
		t.Fatal("no sheet of s-contacts-29 in the references")
	}
	one := w.Sheet{Action: "/x", Title: "One", Body: []w.SheetPart{
		w.Check{On: true, Name: "group", Value: "a", Parts: []w.CheckPart{w.Text("A")}},
		w.Tally{Records: []string{"group=a", "group=a group=b"}}},
		Foot: []w.SheetFootPart{w.Button{Label: "Rewrite 1 record", Kind: w.ButtonPrimary, Submit: true, Count: "Rewrite {n} records", CountOne: "Rewrite 1 record"}}}
	const set = `(n, v, on) => { const f = [...document.querySelectorAll('[name="' + n + '"]')].find(x => x.tagName === "SELECT" || x.value === v);
		if (f.tagName === "SELECT") f.value = v; else f.checked = on; f.dispatchEvent(new Event("change", {bubbles: true})); }`
	const state = `(() => { const b = document.querySelector("[data-tc-count]"); return b.textContent.trim() + (b.disabled ? " (off)" : ""); })()`
	h := browser.NewHarness(t)
	for _, c := range []struct {
		name  string
		body  w.Sheet
		steps string
		want  string
		bare  bool
	}{
		{"as printed", sheet, ``, "Rewrite 17 records", false},
		{"past meetings left", sheet, `set("group", "past", false)`, "Rewrite 12 records", false},
		{"organised by others ticked", sheet, `set("group", "others", true)`, "Rewrite 20 records", false},
		{"the first address kept", sheet, `set("to.1", "")`, "Rewrite 0 records (off)", false},
		{"only the second rewritten", sheet, `set("to.1", ""); set("to.2", "anna@home.example")`, "Rewrite 3 records", false},
		{"one", one, `set("group", "a", true)`, "Rewrite 1 record", false},
		{"none", one, `set("group", "a", false)`, "Rewrite 0 records (off)", false},
		{"without a script", sheet, `set("group", "past", false)`, "Rewrite 17 records", true},
	} {
		page := browser.PageHTML(`<div class="tc-root">`+c.body.Render()+`</div>`, "")
		if c.bare {
			page = strings.Replace(page, `<script src="/tc.js"></script>`, "", 1)
		}
		var got string
		if err := h.Run(h.RegisterPage("tally", page), 900, 900,
			chromedp.Evaluate(`(() => { const set = `+set+`; `+c.steps+`; return true })()`, nil),
			chromedp.Evaluate(state, &got),
		); err != nil {
			t.Fatal(err)
		}
		if got != c.want {
			t.Errorf("%s: %s; want %s", c.name, got, c.want)
		}
	}
}

// zonePart is fixed markup as a part of a fragment.
type zonePart string

func (z zonePart) Render() string { return string(z) }

// TestZones (Carrel's task 84): a link or a form whose container names a
// target asks its own address from the page with the headers the document
// names, and the answer — the part and the other zones in templates — is put
// into the zones; the zone that owns the address puts it into the history;
// an empty answer empties and hides the zone, as does one of the page's
// forms only; a widget carrying the zone's
// name takes its place whole; an answer that is no part sends the page there
// the plain way. The bite: without the live layer the link loads the page.
func TestZones(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	part := h.RegisterPage("zone-part", page.Fragment{Body: []page.Part{zonePart(`<p class="got">new main</p>`)},
		Extra: []page.ZonePart{{Zone: "notices", Body: []page.Part{zonePart(`<p>saved</p>`)}}}}.Render())
	empty := h.RegisterPage("zone-empty", page.Fragment{}.Render())
	formsOnly := h.RegisterPage("zone-forms", page.Fragment{Forms: []w.PostForm{{ID: "post-only", Action: "/only"}}}.Render())
	whole := h.RegisterPage("zone-whole", page.Fragment{Body: []page.Part{zonePart(`<div data-tc-zone="main" class="whole">a new zone</div>`)}}.Render())
	missing := h.RegisterPage("zone-missing", "")
	missing = strings.Replace(missing, "zone-missing", "zone-not-there", 1)
	nav := w.Nav{Target: "main", Items: []w.NavItem{{Label: "Part", Href: part}, {Label: "Empty", Href: empty}, {Label: "Whole", Href: whole}, {Label: "Gone", Href: missing}, {Label: "Forms", Href: formsOnly}}}
	form := w.Form{Action: part, Method: "post", Target: "main", Fields: []w.Hidden{{Name: "csrf", Value: "tok"}},
		Items: []w.FormPart{w.Field{Label: "Name", Parts: []w.FieldPart{w.Input{Name: "name", Value: "x"}, w.Button{Label: "Send", Submit: true}}}}}
	body := `<div class="tc-root">` + nav.Render() + form.Render() + `<div data-tc-zone="main"><p>old main</p></div><div data-tc-zone="notices"><p>old notice</p></div></div>`
	doc := strings.Replace(browser.PageHTML(body, ""), `data-tc-noscript>`,
		`data-tc-noscript data-tc-view-header="X-View" data-tc-view-value="part" data-tc-target-header="X-Target" data-tc-token-field="csrf" data-tc-token-header="X-Token" data-tc-address-zone="main">`, 1)
	with := h.RegisterPage("zones", doc)
	bare := h.RegisterPage("zones-bare", strings.Replace(doc, `<script src="/tc-live.js"></script>`, "", 1))
	const spy = `(() => { window.__kept = true; window.__asked = []; const f = window.fetch;
		window.fetch = (u, o) => { const h = (o && o.headers) || {}; window.__asked.push((o && o.method) + " " + ["X-View", "X-Target", "X-Token"].map(k => h[k] || "-").join(" ")); return f(u, o); }; return true })()`
	const state = `(() => { const z = n => { const e = document.querySelector('[data-tc-zone="' + n + '"]'); return e ? (e.getClientRects().length ? e.textContent.trim() : "hidden") + (e.className ? "." + e.className : "") : "no zone"; };
		return z("main") + " | " + z("notices") + " | " + (window.__kept === true ? "kept" : "loaded") + " | " + location.pathname.split("/").pop() + " | " + (window.__asked || []).join(",") })()`
	for _, c := range []struct{ name, url, click, want string }{
		{"a link's part and another zone", with, `.tc-nav__item:nth-child(1)`, "new main | saved | kept | zone-part | GET part main -"},
		{"an empty answer", with, `.tc-nav__item:nth-child(2)`, "hidden | old notice | kept | zone-empty | GET part main -"},
		{"the zone's own widget", with, `.tc-nav__item:nth-child(3)`, "a new zone.whole | old notice | kept | zone-whole | GET part main -"},
		{"a form with its token", with, `.tc-form .tc-btn`, "new main | saved | kept | zones | POST part main tok"},
		{"no part at the address", with, `.tc-nav__item:nth-child(4)`, "no zone | no zone | loaded | zone-not-there | "},
		// review finding 12: the form went to the page, the zone is emptied
		{"an answer of forms only", with, `.tc-nav__item:nth-child(5)`, "hidden | old notice | kept | zone-forms | GET part main -"},
		// the page went to the address: what is shown is the answer itself, its other zone a template
		{"without the layer", bare, `.tc-nav__item:nth-child(1)`, "no zone | hidden | loaded | zone-part | "},
	} {
		var got string
		if err := h.Run(c.url, 900, 600,
			chromedp.Evaluate(spy, nil),
			chromedp.Click(c.click, chromedp.ByQuery),
			chromedp.Sleep(1500*time.Millisecond),
			chromedp.Evaluate(state, &got),
		); err != nil {
			t.Fatal(err)
		}
		if got != c.want {
			t.Errorf("%s: %s; want %s", c.name, got, c.want)
		}
	}
}

// TestRailTickZones (the owner, 07.11): the group tick of the rail is a
// button of the rail's form — it sends «all», and the answer, asked for the
// rail's zone, puts the rail with every tick under the group set and the
// list in place, the page kept. The bite: without the live layer the page
// goes to the answer.
func TestRailTickZones(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	rail := func(all w.Ticked, one w.Ticked) string {
		return w.RailSec{Items: []w.RailSecPart{
			w.TickRow{Ticked: all, Kind: w.TickRowAll, Mark: w.Mark{Kind: w.MarkMulti}, Label: "Combined calendar", Meta: "7", Send: w.Send{Form: "post-rail", Name: "src", Value: "all"}},
			w.TickRow{Ticked: one, Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "3", Hidden: one == w.TickedOff, Send: w.Send{Form: "post-rail", Name: "src", Value: "nextcloud/lichnyj"}},
		}}.Render()
	}
	answer := h.RegisterPage("rail-answer", page.Fragment{Body: []page.Part{zonePart(rail(w.TickedOn, w.TickedOn))},
		Extra: []page.ZonePart{{Zone: "main", Body: []page.Part{zonePart(`<p>7 events</p>`)}}}}.Render())
	form := w.PostForm{ID: "post-rail", Action: answer, Method: "post", Target: "rail"}
	body := `<div class="tc-root"><div data-tc-zone="rail">` + rail(w.TickedPart, w.TickedOff) + `</div><div data-tc-zone="main"><p>4 events</p></div>` + form.Render() + `</div>`
	doc := strings.Replace(browser.PageHTML(body, ""), `data-tc-noscript>`,
		`data-tc-noscript data-tc-view-header="X-View" data-tc-view-value="part" data-tc-target-header="X-Target">`, 1)
	with := h.RegisterPage("rail-ticks", doc)
	bare := h.RegisterPage("rail-ticks-bare", strings.Replace(doc, `<script src="/tc-live.js"></script>`, "", 1))
	const state = `(() => { window.__kept = window.__kept || false;
		const ticks = Array.from(document.querySelectorAll(".tc-tickrow__tick"), b => b.getAttribute("aria-pressed")).join(" ");
		const main = document.querySelector('[data-tc-zone="main"]');
		return ticks + " | " + (main ? main.textContent.trim() : "no zone") + " | " + (window.__kept ? "kept" : "loaded"); })()`
	for _, c := range []struct{ name, url, want string }{
		{"with the live layer", with, "true true | 7 events | kept"},
		{"without the layer", bare, "true true |  | loaded"}, // the answer itself: the list a template
	} {
		var got string
		if err := h.Run(c.url, 900, 600,
			chromedp.Evaluate(`window.__kept = true`, nil),
			chromedp.Click(`.tc-tickrow.is-all .tc-tickrow__tick`, chromedp.ByQuery),
			chromedp.Sleep(1500*time.Millisecond),
			chromedp.Evaluate(state, &got),
		); err != nil {
			t.Fatal(err)
		}
		if got != c.want {
			t.Errorf("%s: %s; want %s", c.name, got, c.want)
		}
	}
}

// TestListOneStop (the owner, 07.11): a list is one stop of Tab — Tab
// comes to one row and goes past the list; the arrows walk the rows, and
// the list remembers the row they left; Space on a row presses its tick,
// which is no stop of its own. The bite: without the library's script every
// tick and row is a stop.
func TestListOneStop(t *testing.T) {
	needChrome(t)
	row := func(id string) string {
		return `<div class="tc-row"><button class="tc-tick" type="button" aria-label="Done" aria-pressed="false"></button><a class="tc-row__title" href="#` + id + `">` + id + `</a></div>`
	}
	body := `<div class="tc-root"><button id="before" type="button">before</button><div class="tc-list">` + row("a") + row("b") + row("c") + `</div><button id="after" type="button">after</button></div>`
	const at = `(() => { const e = document.activeElement; return e.id || e.textContent || e.getAttribute("aria-label"); })()`
	const spy = `(() => { window.__pressed = []; document.querySelectorAll(".tc-tick").forEach((b, i) => b.addEventListener("click", () => window.__pressed.push(i))); document.getElementById("before").focus(); return true; })()`
	shiftTab := chromedp.KeyEvent(kb.Tab, chromedp.KeyModifiers(input.ModifierShift))
	h := browser.NewHarness(t)
	run := func(url string) string {
		var a, b, c, d, e, pressed string
		if err := h.Run(url, 900, 600,
			chromedp.Evaluate(spy, nil),
			chromedp.KeyEvent(kb.Tab),
			chromedp.Evaluate(at, &a),
			chromedp.KeyEvent(kb.Tab),
			chromedp.Evaluate(at, &b),
			shiftTab,
			chromedp.KeyEvent(kb.ArrowDown),
			chromedp.Evaluate(at, &c),
			chromedp.KeyEvent(kb.Tab),
			shiftTab,
			chromedp.Evaluate(at, &d),
			chromedp.KeyEvent(" "),
			chromedp.Evaluate(at, &e),
			chromedp.Evaluate(`window.__pressed.join(" ")`, &pressed),
		); err != nil {
			t.Fatal(err)
		}
		return strings.Join([]string{a, b, c, d, e, pressed}, " | ")
	}
	doc := browser.PageHTML(body, "")
	with := run(h.RegisterPage("one-stop", doc))
	bare := run(h.RegisterPage("one-stop-bare", strings.Replace(doc, `<script src="/tc.js"></script>`, "", 1)))
	if want := "a | after | b | b | b | 1"; with != want {
		t.Errorf("with the script: %s; want %s", with, want)
	}
	if want := "Done | a | Done | Done | Done | 0"; bare != want {
		t.Errorf("without the script: %s; want %s", bare, want)
	}
}

// TestEnterKey (the keys sheet: «Enter — open the chosen record full»): a
// panel's button that carries the key Enter is pressed by Enter while the
// focus is on nothing that presses; Enter on a row opens the row, and in a
// field it is the field's. The bite: without the library's script Enter on
// the page does nothing.
func TestEnterKey(t *testing.T) {
	needChrome(t)
	row := func(s string) w.ListRow {
		return w.ExcerptRow{Title: []w.Inline{w.Text(s)}, Href: "#" + s}
	}
	body := `<div class="tc-root">` + w.List{Rows: []w.ListRow{row("One"), row("Two")}}.Render() +
		`<form action="/x" method="get" onsubmit="return false">` + w.Input{Name: "t", Label: "Text"}.Render() + `</form>` +
		w.Button{Label: "Open full", Href: "#full", Key: "Enter"}.Render() + `</div>`
	const where = `location.hash`
	h := browser.NewHarness(t)
	doc := browser.PageHTML(body, "")
	run := func(url string) string {
		var onPage, onRow, inField string
		if err := h.Run(url, 900, 600,
			chromedp.Evaluate(`document.activeElement.blur(); true`, nil),
			chromedp.KeyEvent(""),
			chromedp.Evaluate(where, &onPage),
			chromedp.Evaluate(`history.replaceState(null, "", location.pathname); document.querySelector(".tc-row").focus(); true`, nil),
			chromedp.KeyEvent(""),
			chromedp.Evaluate(where, &onRow),
			chromedp.Evaluate(`history.replaceState(null, "", location.pathname); document.querySelector("input[name=t]").focus(); true`, nil),
			chromedp.KeyEvent(""),
			chromedp.Evaluate(where, &inField),
		); err != nil {
			t.Fatal(err)
		}
		return onPage + " | " + onRow + " | " + inField
	}
	with := run(h.RegisterPage("enter-key", doc))
	bare := run(h.RegisterPage("enter-key-bare", strings.Replace(doc, `<script src="/tc.js"></script>`, "", 1)))
	if want := "#full | #One | "; with != want {
		t.Errorf("with the script: %q; want %q", with, want)
	}
	if want := " | #One | "; bare != want {
		t.Errorf("without the script: %q; want %q", bare, want)
	}
}

// TestZoneFormOnce (Carrel's note 58): a form that is not GET, sent by the
// live layer, is sent once — an answer sent on to another address goes
// there, an answer that is not a part is shown whole, and neither is sent
// again the plain way. A GET form whose answer is not a part still goes the
// plain way: asking again changes nothing. The bite: the old layer sent
// each of them twice.
func TestZoneFormOnce(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	var sent atomic.Int32
	done := h.RegisterPage("once-done", browser.PageHTML(`<p class="done">written</p>`, ""))
	at := map[string]string{}
	at["once-moved"] = h.RegisterStream("once-moved", func(w http.ResponseWriter, r *http.Request) {
		sent.Add(1)
		http.Redirect(w, r, done, http.StatusSeeOther)
	})
	at["once-refused"] = h.RegisterStream("once-refused", func(w http.ResponseWriter, r *http.Request) {
		sent.Add(1)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`<!doctype html><p class="refused">the server refused</p>`))
	})
	at["once-asked"] = h.RegisterStream("once-asked", func(w http.ResponseWriter, r *http.Request) {
		sent.Add(1)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte("not a part"))
	})
	for _, c := range []struct {
		name, stream, method, want string
		times                      int32
	}{
		{"sent on", "once-moved", "post", "written | once-done", 1},
		{"refused", "once-refused", "post", "the server refused", 1},
		{"asked by GET", "once-asked", "get", "not a part", 2},
	} {
		sent.Store(0)
		form := w.Form{Action: at[c.stream], Method: c.method, Target: "main",
			Items: []w.FormPart{w.Field{Label: "Name", Parts: []w.FieldPart{w.Input{Name: "name", Value: "x"}, w.Button{Label: "Send", Submit: true}}}}}
		body := `<div class="tc-root">` + form.Render() + `<div data-tc-zone="main"><p>old main</p></div></div>`
		doc := strings.Replace(browser.PageHTML(body, ""), `data-tc-noscript>`, `data-tc-noscript data-tc-view-header="X-View" data-tc-target-header="X-Target">`, 1)
		var got string
		if err := h.Run(h.RegisterPage("once-"+c.stream, doc), 900, 600,
			chromedp.Click(`.tc-form .tc-btn`, chromedp.ByQuery),
			chromedp.Sleep(1500*time.Millisecond),
			chromedp.Evaluate(`(() => { const d = document.querySelector(".done"); return d ? d.textContent + " | " + location.pathname.split("/").pop() : document.body.textContent.trim(); })()`, &got),
		); err != nil {
			t.Fatal(err)
		}
		if got != c.want || sent.Load() != c.times {
			t.Errorf("%s: %q, sent %d times; want %q, %d", c.name, got, sent.Load(), c.want, c.times)
		}
	}
}

// TestQueueFolds (a tail of task 88): «Hide» folds the queue in place — the
// list hides, the toggle says Show, the total says the percent once the
// upload's words are known — and «Show» unfolds it, with no request.
// Without a script the toggle is the server's link. The bite: before, the
// toggle only went to the server.
func TestQueueFolds(t *testing.T) {
	needChrome(t)
	q := w.Queue{Title: "Uploading 2 of 3", Total: w.Mono{Text: "{sent} of {total}", Variant: w.MonoDusty, Place: w.Pushed},
		Toggle: w.Button{Label: "Hide", Small: true, Href: "/elsewhere"}, Variant: w.QueueFlushAuto,
		Items: []w.QueueItem{{Name: "a.pdf", Meta: "done"}, {Name: "b.png", Meta: "40%"}},
		Fold:  &w.QueueFold{Hide: "Hide", Show: "Show", Total: "{percent}%"}}
	body := `<div class="tc-root">` + q.Render() + `</div>`
	const fill = `(() => { const h = document.querySelector(".tc-queue__head"); window.__tc.fillWords(h, {"{sent}": "3 MB", "{total}": "5 MB", "{percent}": 62}); return true })()`
	const press = `document.querySelector(".tc-queue__head .tc-btn").click(); true`
	const state = `(() => { const q = document.querySelector(".tc-queue"), l = q.querySelector(".tc-queue__list");
		return (q.classList.contains("is-min") ? "folded" : "open") + " | " + (l.getClientRects().length ? "list" : "no list") + " | " +
			q.querySelector(".tc-queue__head .tc-mono").textContent + " | " + q.querySelector(".tc-queue__head .tc-btn").textContent.trim() + " | " + location.pathname.split("/").pop() })()`
	h := browser.NewHarness(t)
	for _, c := range []struct {
		name, steps, want string
		bare              bool
	}{
		{"hidden", press, "folded | no list | {sent} of {total} | Show | queue", false},
		{"shown again", press + ";" + press, "open | list | {sent} of {total} | Hide | queue", false},
		{"hidden while uploading", fill + ";" + press, "folded | no list | 62% | Show | queue", false},
		{"shown while uploading", fill + ";" + press + ";" + press, "open | list | 3 MB of 5 MB | Hide | queue", false},
	} {
		page := browser.PageHTML(body, "")
		if c.bare {
			page = strings.Replace(page, `<script src="/tc.js"></script>`, "", 1)
		}
		var got string
		if err := h.Run(h.RegisterPage("queue", page), 900, 600,
			chromedp.Evaluate(`(() => { `+c.steps+`; return true })()`, nil),
			chromedp.Sleep(300*time.Millisecond),
			chromedp.Evaluate(state, &got),
		); err != nil {
			t.Fatal(err)
		}
		if got != c.want {
			t.Errorf("%s: %s; want %s", c.name, got, c.want)
		}
	}
}

// TestZoneSheet (Carrel's task 84): a sheet that comes as an answer lies in
// the layer of sheets — the scrim the application named as a zone — and is
// closed by its Cancel (Button.Closes), by a click on the scrim or by Esc
// without a request and without leaving the page;
// the focus is kept inside it while it is open.
func TestZoneSheet(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	sheet := w.Scrim{Zone: "sheet", Content: w.Sheet{Title: "Rename", Action: "/x", Body: []w.SheetPart{w.Field{Label: "Name", Parts: []w.FieldPart{w.Input{Name: "name", Value: "a.txt"}}}},
		Foot: []w.SheetFootPart{w.Button{Label: "Save", Kind: w.ButtonPrimary, Submit: true}, w.Button{Label: "Cancel", Href: "/back", Closes: true}}}}
	part := h.RegisterPage("sheet-part", page.Fragment{Body: []page.Part{sheet}}.Render())
	body := `<div class="tc-root">` + w.Button{Label: "Rename", Href: part, Target: "sheet"}.Render() + w.Scrim{Zone: "sheet", Closed: true}.Render() + `</div>`
	url := h.RegisterPage("sheet-page", browser.PageHTML(body, ""))
	const state = `(() => { const s = document.querySelector(".tc-scrim"); return (s.getClientRects().length ? "open" : "closed") + " " + (s.querySelector(".tc-sheet__title") ? s.querySelector(".tc-sheet__title").textContent : "empty") + " " + (window.__asked || 0) + " " + (s.contains(document.activeElement) ? "inside" : "outside") })()`
	const spy = `(() => { window.__asked = 0; const f = window.fetch; window.fetch = (u, o) => { window.__asked++; return f(u, o); }; return true })()`
	var opened, closed, again, esc, cancelled string
	if err := h.Run(url, 900, 600,
		chromedp.Evaluate(spy, nil),
		chromedp.Click(`a.tc-btn`, chromedp.ByQuery),
		chromedp.Sleep(1200*time.Millisecond),
		chromedp.Evaluate(`document.querySelector(".tc-sheet input").focus(); true`, nil),
		chromedp.Evaluate(state, &opened),
		chromedp.Evaluate(`document.querySelector(".tc-scrim").click(); true`, nil),
		chromedp.Evaluate(state, &closed),
		chromedp.Click(`a.tc-btn`, chromedp.ByQuery),
		chromedp.Sleep(1200*time.Millisecond),
		chromedp.Evaluate(state, &again),
		chromedp.KeyEvent("\x1b"),
		chromedp.Evaluate(state, &esc),
		chromedp.Click(`a.tc-btn`, chromedp.ByQuery),
		chromedp.Sleep(1200*time.Millisecond),
		chromedp.Click(`.tc-sheet__foot a.tc-btn`, chromedp.ByQuery),
		chromedp.Evaluate(state+` + " " + location.pathname.split("/").pop()`, &cancelled),
	); err != nil {
		t.Fatal(err)
	}
	if opened != "open Rename 1 inside" {
		t.Errorf("after the link: %s; want open Rename 1 inside", opened)
	}
	if !strings.HasPrefix(closed, "closed Rename 1") {
		t.Errorf("after a click on the scrim: %s; want closed, and no request more", closed)
	}
	if !strings.HasPrefix(again, "open Rename 2") || !strings.HasPrefix(esc, "closed Rename 2") {
		t.Errorf("opened again: %s; after Esc: %s; want open with a second request, then closed without a third", again, esc)
	}
	if !strings.HasPrefix(cancelled, "closed Rename 3") || !strings.HasSuffix(cancelled, "sheet-page") {
		t.Errorf("after Cancel: %s; want closed on the same page, the link not followed", cancelled)
	}
}

// TestZonePanel (Carrel's tasks 83 and 84): a row of a list whose target is
// the side panel fills the panel that stood closed, and the shell's body
// makes room for it; the answer brings a message with its technical block
// into the zone of messages; the panel's «×» closes it without a request and
// the room is taken back; an empty answer does the same.
func TestZonePanel(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	panel := w.Side{Zone: "panel", Items: []w.SidePart{w.Head{Title: &w.HeadTitle{Text: "Анна Ковалёва"}, Acts: &w.Acts{Items: []w.ActsPart{w.Close{Label: "Close"}}}}}}
	notice := page.Printed(w.PrintMain(w.Msg{Parts: []w.MsgPart{w.Text("Read from the copy")}}, w.MainDiag{Diag: w.Diag{Text: "GET /dav → 503"}}))
	part := h.RegisterPage("panel-part", page.Fragment{Body: []page.Part{panel}, Extra: []page.ZonePart{{Zone: "notices", Body: []page.Part{notice}}}}.Render())
	empty := h.RegisterPage("panel-empty", "")
	list := w.List{Target: "panel", Rows: []w.ListRow{w.FaceRow{Title: []w.Inline{w.Text("Анна Ковалёва")}, Href: part}, w.FaceRow{Title: []w.Inline{w.Text("Вера Лебедева")}, Href: empty}}}
	shell := w.Shell{Body: w.ShellBody{Main: w.Main{Zone: "main", Items: []w.MainPart{list}}, Side: &w.Side{Zone: "panel", Closed: true}}}
	body := shell.Render() + `<div data-tc-zone="notices"></div>`
	url := h.RegisterPage("panel-page", browser.PageHTML(body, ""))
	const state = `(() => { const p = document.querySelector(".tc-side"), b = document.querySelector(".tc-shell__body"), n = document.querySelector('[data-tc-zone="notices"]');
		return (p.getClientRects().length ? "open" : "closed") + " " + (b.classList.contains("tc-shell__body--side") ? "room" : "no room") + " | " + (n.querySelector(".tc-msg") ? "msg" : "-") + " " + (n.querySelector(".tc-main__diag .tc-diag") ? "diag" : "-") + " | " + (window.__asked || 0) })()`
	const spy = `(() => { window.__asked = 0; const f = window.fetch; window.fetch = (u, o) => { window.__asked++; return f(u, o); }; return true })()`
	var before, opened, closed, again, emptied string
	if err := h.Run(url, 1500, 700,
		chromedp.Evaluate(spy, nil),
		chromedp.Evaluate(state, &before),
		chromedp.Click(`.tc-row:nth-child(1)`, chromedp.ByQuery),
		chromedp.Sleep(1200*time.Millisecond),
		chromedp.Evaluate(state, &opened),
		chromedp.Evaluate(`document.querySelector(".tc-side .tc-close").click(); true`, nil),
		chromedp.Evaluate(state, &closed),
		chromedp.Click(`.tc-row:nth-child(1)`, chromedp.ByQuery),
		chromedp.Sleep(1200*time.Millisecond),
		chromedp.Evaluate(state, &again),
		chromedp.Click(`.tc-row:nth-child(2)`, chromedp.ByQuery),
		chromedp.Sleep(1200*time.Millisecond),
		chromedp.Evaluate(state, &emptied),
	); err != nil {
		t.Fatal(err)
	}
	// On a narrow screen the mockup does not draw the panel: the row's link
	// loads its address, it does not fill a panel nobody sees.
	var narrow string
	if err := h.Run(url, 900, 700,
		chromedp.Evaluate(spy, nil),
		chromedp.Click(`.tc-row:nth-child(1)`, chromedp.ByQuery),
		chromedp.Sleep(1200*time.Millisecond),
		chromedp.Evaluate(`location.pathname.split("/").pop() + " " + (window.__asked === undefined ? "loaded" : "kept")`, &narrow),
	); err != nil {
		t.Fatal(err)
	}
	if narrow != "panel-part loaded" {
		t.Errorf("on a narrow screen: %s; want the page at the row's address", narrow)
	}
	for _, c := range [][3]string{{"as printed", before, "closed no room | - - | 0"}, {"after a row", opened, "open room | msg diag | 1"},
		{"after «×»", closed, "closed no room | msg diag | 1"}, {"after the row again", again, "open room | msg diag | 2"},
		{"after an empty answer", emptied, "closed no room | msg diag | 3"}} {
		if c[1] != c[2] {
			t.Errorf("%s: %s; want %s", c[0], c[1], c[2])
		}
	}
}

// TestRunFollows (Carrel's task 85): a run's line that is a zone follows its
// run. With a stream it takes each event's fragment and closes the stream on
// the event that ends the run; with no stream, or one that refused, it asks
// its poll address every so many seconds for as long as the line that comes
// back names one. The bite: without the live layer the line stays as printed.
func TestRunFollows(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	bar := func(text, stream, poll string) w.Runbar {
		return w.Runbar{Zone: "run", Meter: w.Meter{Percent: 50, Track: 100}, Parts: []w.RunbarPart{w.Text(text)},
			Stream: stream, Changed: "changed", Ended: "ended", Poll: poll, Every: time.Second}
	}
	frag := func(b w.Runbar) string {
		return strings.ReplaceAll(page.Fragment{Body: []page.Part{b}}.Render(), "\n", " ")
	}
	const text = `(document.querySelector(".tc-runbar") || {textContent: "no line"}).textContent.trim()`
	wait := chromedp.Sleep(2600 * time.Millisecond)

	// a stream: two events, the second ends the run
	var opened int32
	stream := h.RegisterStream("run", func(rw http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&opened, 1)
		rw.Header().Set("Content-Type", "text/event-stream")
		fl, _ := rw.(http.Flusher)
		fmt.Fprintf(rw, "event: changed\ndata: %s\n\n", frag(bar("2 of 6", r.URL.Path, "")))
		fl.Flush()
		time.Sleep(400 * time.Millisecond)
		fmt.Fprintf(rw, "event: ended\ndata: %s\n\n", frag(w.Runbar{Zone: "run", Meter: w.Meter{Percent: 100, Track: 100}, Parts: []w.RunbarPart{w.Text("6 of 6")}}))
		fl.Flush()
	})
	var got string
	if err := h.Run(h.RegisterPage("run-stream", browser.PageHTML(`<div class="tc-root">`+bar("0 of 6", stream, "").Render()+`</div>`, "")), 900, 400,
		wait, chromedp.Evaluate(text, &got)); err != nil {
		t.Fatal(err)
	}
	if got != "6 of 6" || atomic.LoadInt32(&opened) != 1 {
		t.Errorf("a stream: the line says %q after %d openings; want 6 of 6 after one — the stream closed on the event that ends the run", got, opened)
	}

	// a poll: the line that comes back names the poll, then does not
	pollURL := h.RegisterPage("run-poll", "")
	h.RegisterPage("run-poll", frag(bar("3 of 6", "", pollURL)))
	var first, last string
	url := h.RegisterPage("run-polled", browser.PageHTML(`<div class="tc-root">`+bar("0 of 6", "", pollURL).Render()+`</div>`, ""))
	if err := h.Run(url, 900, 400,
		// the first answer of the poll: in a second, or later on a loaded
		// machine (1.6 s were too few in a whole run of 14.10)
		await(`(async () => {
			const said = () => `+text+`;
			for (let i = 0; i < 60 && said() === "0 of 6"; i++) await new Promise(r => setTimeout(r, 100));
			return said();
		})()`, &first),
		chromedp.ActionFunc(func(context.Context) error {
			h.RegisterPage("run-poll", frag(w.Runbar{Zone: "run", Meter: w.Meter{Percent: 100, Track: 100}, Parts: []w.RunbarPart{w.Text("done")}}))
			return nil
		}),
		wait, chromedp.Evaluate(text, &last)); err != nil {
		t.Fatal(err)
	}
	first = strings.Trim(first, `"`)
	if first != "3 of 6" || last != "done" {
		t.Errorf("a poll: %q, then %q; want 3 of 6, then done", first, last)
	}

	// a stream that refuses: the poll takes over
	h.RegisterPage("run-poll", frag(w.Runbar{Zone: "run", Meter: w.Meter{Percent: 100, Track: 100}, Parts: []w.RunbarPart{w.Text("polled")}}))
	refused := strings.Replace(stream, "/stream/run", "/stream/none", 1)
	if err := h.Run(h.RegisterPage("run-refused", browser.PageHTML(`<div class="tc-root">`+bar("0 of 6", refused, pollURL).Render()+`</div>`, "")), 900, 400,
		wait, chromedp.Evaluate(text, &got)); err != nil {
		t.Fatal(err)
	}
	if got != "polled" {
		t.Errorf("a stream that refused: the line says %q; want polled", got)
	}

	// a stream that breaks midway (a tail of 14.10): one event, then the
	// connection ends with no event that ends the run — the poll takes over,
	// and the stream is not opened again
	var broke int32
	breaking := h.RegisterStream("run-breaks", func(rw http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&broke, 1)
		rw.Header().Set("Content-Type", "text/event-stream")
		fmt.Fprintf(rw, "event: changed\ndata: %s\n\n", frag(bar("2 of 6", r.URL.Path, pollURL)))
		rw.(http.Flusher).Flush()
	})
	var mid string
	if err := h.Run(h.RegisterPage("run-breaking", browser.PageHTML(`<div class="tc-root">`+bar("0 of 6", breaking, pollURL).Render()+`</div>`, "")), 900, 400,
		wait, chromedp.Evaluate(text, &mid)); err != nil {
		t.Fatal(err)
	}
	if mid != "polled" || atomic.LoadInt32(&broke) != 1 {
		t.Errorf("a stream that broke: the line says %q after %d openings; want polled after one", mid, broke)
	}

	// without the layer the line stays as printed
	bare := strings.Replace(browser.PageHTML(`<div class="tc-root">`+bar("0 of 6", stream, pollURL).Render()+`</div>`, ""), `<script src="/tc-live.js"></script>`, "", 1)
	if err := h.Run(h.RegisterPage("run-bare", bare), 900, 400, wait, chromedp.Evaluate(text, &got)); err != nil {
		t.Fatal(err)
	}
	if got != "0 of 6" {
		t.Errorf("without the layer: the line says %q; want 0 of 6", got)
	}
}

// TestRunHiddenRests (Carrel's task 226): a run's line emptied by the event
// that ends the run is hidden and keeps its stream's address; it is not
// followed again — the stream is opened once, not after every end. A line
// that asks its poll every 300 ms asks it in 300 ms: the interval is in
// milliseconds. The bites: the old layer opened the stream again on each end,
// and read 300 as seconds.
func TestRunHiddenRests(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	frag := func(b w.Runbar) string {
		return strings.ReplaceAll(page.Fragment{Body: []page.Part{b}}.Render(), "\n", " ")
	}
	var opened int32
	stream := h.RegisterStream("run-ends-empty", func(rw http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&opened, 1)
		rw.Header().Set("Content-Type", "text/event-stream")
		fmt.Fprintf(rw, "event: ended\ndata: \n\n")
		rw.(http.Flusher).Flush()
	})
	bar := w.Runbar{Zone: "run", Meter: w.Meter{Percent: 50, Track: 100}, Parts: []w.RunbarPart{w.Text("0 of 6")}, Stream: stream, Changed: "changed", Ended: "ended"}
	const state = `(() => { const b = document.querySelector(".tc-runbar"); return (b.hidden ? "hidden" : "shown") + " " + JSON.stringify(b.textContent.trim()) })()`
	var got string
	if err := h.Run(h.RegisterPage("run-ends-empty-page", browser.PageHTML(`<div class="tc-root">`+bar.Render()+`</div>`, "")), 900, 400,
		chromedp.Sleep(2600*time.Millisecond), chromedp.Evaluate(state, &got)); err != nil {
		t.Fatal(err)
	}
	if got != `hidden ""` || atomic.LoadInt32(&opened) != 1 {
		t.Errorf("a line emptied by the end: %s after %d openings; want hidden and empty after one", got, opened)
	}

	poll := h.RegisterPage("run-fast-poll", frag(w.Runbar{Zone: "run", Meter: w.Meter{Percent: 100, Track: 100}, Parts: []w.RunbarPart{w.Text("fast")}}))
	fast := w.Runbar{Zone: "run", Meter: w.Meter{Percent: 10, Track: 100}, Parts: []w.RunbarPart{w.Text("0 of 6")}, Poll: poll, Every: 300 * time.Millisecond}
	if !strings.Contains(fast.Render(), `data-tc-every="300"`) {
		t.Fatalf("the interval is not printed in milliseconds: %s", fast.Render())
	}
	if err := h.Run(h.RegisterPage("run-fast-page", browser.PageHTML(`<div class="tc-root">`+fast.Render()+`</div>`, "")), 900, 400,
		chromedp.Sleep(1800*time.Millisecond), chromedp.Evaluate(state, &got)); err != nil {
		t.Fatal(err)
	}
	if got != `shown "fast"` {
		t.Errorf("a poll every 300 ms: %s after 1.8 s; want the answer of the poll", got)
	}
}

// TestZoneAnswerForms (Carrel's task 226): an answer is the zone's widget and,
// after it, the page's forms its buttons name by ID. The widget takes the
// zone's place whole — a panel is not put inside the panel; a form takes the
// place of the page's form of the same ID, a new one joins the end of the
// page. The zone of messages (Notices) waits hidden and shows the messages
// that come. The bite: the old layer put all of it into the zone, the panel
// in the panel and a second form of the same ID.
func TestZoneAnswerForms(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	side := func(title, form string) w.Side {
		return w.Side{Zone: "panel", Items: []w.SidePart{w.Head{Title: &w.HeadTitle{Text: title}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Delete", Form: form}}}}}}
	}
	part := h.RegisterPage("forms-part", page.Fragment{Body: []page.Part{side("Новая", "post-del")},
		Forms: []w.PostForm{{ID: "post-del", Action: "/new-delete"}, {ID: "post-move", Action: "/move"}},
		Extra: []page.ZonePart{{Zone: "notices", Body: []page.Part{w.Msg{Parts: []w.MsgPart{w.Text("Read from the copy")}}}}}}.Render())
	body := `<div class="tc-root">` + w.Notices{Zone: "notices"}.Render() + w.Button{Label: "Open", Href: part, Target: "panel"}.Render() + side("Старая", "post-del").Render() + `</div>` +
		w.PostForm{ID: "post-del", Action: "/old-delete"}.Render()
	doc := strings.Replace(browser.PageHTML(body, ""), `data-tc-noscript>`, `data-tc-noscript data-tc-target-header="X-Target">`, 1)
	var got string
	if err := h.Run(h.RegisterPage("forms-page", doc), 900, 500,
		chromedp.Click(`a.tc-btn`, chromedp.ByQuery),
		chromedp.Sleep(1200*time.Millisecond),
		chromedp.Evaluate(`[document.querySelectorAll(".tc-side").length, document.querySelector(".tc-side .tc-head__title").textContent,
			document.querySelectorAll("#post-del").length, new URL(document.getElementById("post-del").action).pathname,
			document.getElementById("post-move") ? "move" : "no move", document.querySelector(".tc-side form") ? "a form in the panel" : "no form in the panel",
			(document.querySelector(".tc-notices").hidden ? "notices hidden" : "notices: " + document.querySelector(".tc-notices").textContent.trim())].join(" | ")`, &got),
	); err != nil {
		t.Fatal(err)
	}
	const want = "1 | Новая | 1 | /new-delete | move | no form in the panel | notices: Read from the copy"
	if got != want {
		t.Errorf("after the answer: %s; want %s", got, want)
	}
}

// TestShortcuts (Carrel's task 87): what is pressed carries its key as data.
// A letter presses its button and a sign puts the focus into the search —
// not while the cursor is in a field, and not from under a sheet; «?» opens
// the sheet of keys through an item of a menu that is closed; a chord with
// Ctrl works in a field, for the button of that field's form, on a layout
// without Latin letters too; what destroys carries no key. The arrows walk
// the rows of a list.
func TestShortcuts(t *testing.T) {
	needChrome(t)
	form := func(id string) string {
		return `<form id="` + id + `" action="/x" method="post">` + w.Input{Name: "t", Label: "Text"}.Render() +
			w.Button{Label: "Save", Submit: true, Name: "from", Value: id, Key: "Ctrl+S"}.Render() + `</form>`
	}
	row := func(s string) w.ListRow {
		return w.ExcerptRow{Title: []w.Inline{w.Text(s)}, Href: "#" + s}
	}
	body := `<div class="tc-root">` +
		w.Search{Query: w.Query{Label: "Search", Name: "q", Key: "/"}}.Render() +
		w.Button{Label: "New note", Opens: "quick", Key: "n"}.Render() +
		w.Button{Label: "Delete", Kind: w.ButtonDangerSolid, Key: "d"}.Render() +
		w.Drop{Trigger: w.Button{Label: "mix", Menu: true}, Menu: w.Menu{Pop: w.MenuUnder, Items: []w.MenuPart{
			w.MenuItem{Label: "Keyboard shortcuts", Href: "#keys", Key: "?"},
			w.MenuItem{Label: "Delete account", Tone: w.MenuToneDanger, Key: "x"}}}}.Render() +
		form("fa") + form("fb") +
		w.List{Rows: []w.ListRow{row("One"), row("Two"), row("Three")}}.Render() +
		w.Scrim{ID: "quick", Closed: true, Content: w.Sheet{Title: "Quick note", Action: "/q",
			Body: []w.SheetPart{w.Field{Label: "Text", Parts: []w.FieldPart{w.Input{Name: "body"}}}},
			Foot: []w.SheetFootPart{w.Button{Label: "Save", Submit: true, Name: "from", Value: "quick", Key: "Ctrl+Enter"}}}}.Render() +
		w.Scrim{ID: "keys", Closed: true, Content: w.Sheet{Title: "Keyboard shortcuts", Head: []w.SheetHeadPart{w.Close{Label: "Close"}}}}.Render() +
		`</div>`
	h := browser.NewHarness(t)
	const prep = `(() => { window.__sent = []; document.addEventListener("submit", e => { e.preventDefault(); window.__sent.push(e.submitter ? e.submitter.value : "?"); });
		window.__open = id => document.getElementById(id).getClientRects().length ? "open" : "closed";
		window.__key = (o) => { (document.activeElement || document.body).dispatchEvent(new KeyboardEvent("keydown", Object.assign({bubbles: true, cancelable: true}, o))); return true };
		window.__at = () => { const a = document.activeElement; return a.type === "search" ? "search" : (a.textContent || a.tagName).trim() }; return true })()`
	var inField, pressed, under, chord, other, slash, keys, keysEsc, own, first, danger, walked string
	if err := h.Run(h.RegisterPage("shortcuts", browser.PageHTML(body, "")), 900, 700,
		chromedp.Evaluate(prep, nil),
		chromedp.Focus(`#fa input`, chromedp.ByQuery),
		chromedp.KeyEvent("n"),
		chromedp.Evaluate(`__open("quick") + " " + document.querySelector("#fa input").value`, &inField),
		chromedp.Evaluate(`document.activeElement.blur(); true`, nil),
		chromedp.KeyEvent("n"),
		chromedp.Evaluate(`__open("quick") + " " + (document.getElementById("quick").contains(document.activeElement) ? "inside" : "outside")`, &pressed),
		chromedp.Evaluate(`document.querySelector("#quick button[type=submit]").focus(); true`, nil),
		chromedp.KeyEvent("/"),
		chromedp.Evaluate(`__at()`, &under),
		chromedp.Evaluate(`document.querySelector("#quick input").focus(); __key({key: "s", code: "KeyS", ctrlKey: true}); __sent.join(",")`, &other),
		chromedp.Evaluate(`__key({key: "Enter", code: "Enter", ctrlKey: true}); __sent.join(",")`, &chord),
		chromedp.Evaluate(`document.getElementById("quick").click(); document.activeElement.blur(); true`, nil),
		chromedp.KeyEvent("/"),
		chromedp.Evaluate(`__at()`, &slash),
		chromedp.Evaluate(`document.activeElement.blur(); true`, nil),
		chromedp.KeyEvent("?"),
		chromedp.Evaluate(`__open("keys")`, &keys),
		chromedp.KeyEvent("\x1b"),
		chromedp.Evaluate(`__open("keys")`, &keysEsc),
		chromedp.Evaluate(`document.querySelector("#fb input").focus(); __key({key: "ы", code: "KeyS", ctrlKey: true}); __sent.join(",")`, &own),
		chromedp.Evaluate(`document.activeElement.blur(); __key({key: "s", code: "KeyS", ctrlKey: true}); __sent.join(",")`, &first),
		chromedp.Evaluate(`String(document.querySelectorAll(".is-danger-solid[data-tc-key], .tc-menu__item--danger[data-tc-key]").length)`, &danger),
		chromedp.Evaluate(`(() => { const got = []; document.querySelector(".tc-list a").focus();
			for (const k of ["ArrowDown", "End", "ArrowUp", "Home", "ArrowUp"]) { __key({key: k}); got.push(__at()); } return got.join(",") })()`, &walked),
	); err != nil {
		t.Fatal(err)
	}
	if inField != "closed n" {
		t.Errorf("«n» typed in a field: %s; want the sheet closed and the letter in the field", inField)
	}
	if pressed != "open inside" {
		t.Errorf("«n» outside a field: %s; want the sheet open, the focus in it", pressed)
	}
	if under == "search" {
		t.Errorf("«/» over an open sheet put the focus into the search under it")
	}
	if other != "" || chord != "quick" {
		t.Errorf("in the sheet's field: after Ctrl+S sent %q, after Ctrl+Enter %q; want nothing, then quick", other, chord)
	}
	if slash != "search" {
		t.Errorf("«/»: the focus is in %q; want the search", slash)
	}
	if keys != "open" || keysEsc != "closed" {
		t.Errorf("«?»: the sheet of keys is %s, after Esc %s; want open, closed", keys, keysEsc)
	}
	if own != "quick,fb" || first != "quick,fb,fa" {
		t.Errorf("Ctrl+S in the second form's field sent %q, outside a field %q; want quick,fb and quick,fb,fa", own, first)
	}
	if danger != "0" {
		t.Errorf("%s of what destroys carry a key; want none", danger)
	}
	if walked != "Two,Three,Two,One,One" {
		t.Errorf("the arrows in a list: %s; want Two,Three,Two,One,One", walked)
	}
}

// TestButtonBusy (Carrel's task 242): a submit button with Busy says its
// words and is off once its form goes — after the event, so that the form
// still sends the button's name and value; its icon stays. A button
// without Busy is left as it is.
func TestButtonBusy(t *testing.T) {
	needChrome(t)
	body := `<div class="tc-root"><form id="letter" action="/send" method="post">` +
		w.Input{Name: "subject", Label: "Subject"}.Render() +
		w.Button{Label: "Send", Icon: "mail", Kind: w.ButtonPrimary, Submit: true, Name: "op", Value: "send", Busy: "Sending…"}.Render() +
		w.Button{Label: "Preview", Submit: true, Name: "op", Value: "preview"}.Render() +
		`</form></div>`
	h := browser.NewHarness(t)
	const prep = `(() => { window.__sent = []; document.addEventListener("submit", e => { e.preventDefault();
		window.__sent.push(new FormData(e.target, e.submitter).get("op")); }); return true })()`
	var before, after, preview, sent string
	if err := h.Run(h.RegisterPage("busy", browser.PageHTML(body, "")), 900, 500,
		chromedp.Evaluate(prep, nil),
		chromedp.Evaluate(`(() => { const b = document.querySelector("[data-tc-busy]"); return b.textContent.trim() + "|" + b.disabled })()`, &before),
		chromedp.Evaluate(`document.querySelector("[data-tc-busy]").click(); true`, nil),
		chromedp.Sleep(100*time.Millisecond),
		chromedp.Evaluate(`(() => { const b = document.querySelector("[data-tc-busy]"); return b.textContent.trim() + "|" + b.disabled + "|" + b.getAttribute("aria-busy") + "|" + !!b.querySelector("svg") })()`, &after),
		chromedp.Evaluate(`(() => { const b = document.querySelector("button[value=preview]"); b.click(); return b.textContent.trim() + "|" + b.disabled })()`, &preview),
		chromedp.Evaluate(`__sent.join(",")`, &sent),
	); err != nil {
		t.Fatal(err)
	}
	if before != "Send|false" {
		t.Errorf("before the press: %s; want Send|false", before)
	}
	if after != "Sending…|true|true|true" {
		t.Errorf("after the press: %s; want Sending…|true|true|true (words, off, busy, the icon kept)", after)
	}
	if preview != "Preview|false" {
		t.Errorf("a button without Busy after its press: %s; want it as it was", preview)
	}
	if sent != "send,preview" {
		t.Errorf("the form sent %q; want send,preview — the busy button's value goes too", sent)
	}
}

// TestButtonBusyBack (review findings 3–5): a busy button comes back when
// the live layer puts the answer in a zone and its form stays on the page —
// its words, on, not busy; sent again the plain way it keeps its own words,
// not the busy ones; a label in its own element (Folds) says the words too.
func TestButtonBusyBack(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	sent := h.RegisterStream("send", func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(300 * time.Millisecond)
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(rw, page.Fragment{Body: []page.Part{zonePart(`<p class="got">sent</p>`)}}.Render())
	})
	body := `<div class="tc-root"><div data-tc-zone="main"><p>old</p></div>` +
		`<form id="letter" action="` + sent + `" method="post" data-tc-target="main">` +
		w.Input{Name: "subject", Label: "Subject"}.Render() +
		w.Button{Label: "Send", Icon: "mail", Kind: w.ButtonPrimary, Submit: true, Busy: "Sending…"}.Render() +
		w.Button{Label: "Save", Icon: "save", Folds: true, Submit: true, Name: "op", Value: "save", Busy: "Saving…"}.Render() +
		`</form><form id="plain" action="/nowhere" method="post">` +
		w.Button{Label: "Go", Submit: true, Busy: "Going…"}.Render() + `</form></div>`
	const state = `(sel) => { const b = document.querySelector(sel); return b.textContent.trim() + "|" + b.disabled + "|" + b.getAttribute("aria-busy") }`
	var sending, back, folds, again string
	if err := h.Run(h.RegisterPage("busy-back", browser.PageHTML(body, "")), 900, 500,
		chromedp.Evaluate(`window.__state = `+state+`; document.addEventListener("submit", e => { if (e.target.id === "plain") e.preventDefault() }); true`, nil),
		chromedp.Evaluate(`document.querySelector("#letter .is-primary").click(); true`, nil),
		chromedp.Sleep(100*time.Millisecond),
		chromedp.Evaluate(`__state("#letter .is-primary")`, &sending),
		chromedp.Sleep(700*time.Millisecond),
		chromedp.Evaluate(`__state("#letter .is-primary") + "|" + document.querySelector(".got").textContent`, &back),
		chromedp.Evaluate(`document.querySelector("button[value=save]").click(); true`, nil),
		chromedp.Sleep(100*time.Millisecond),
		chromedp.Evaluate(`__state("button[value=save] .tc-btn__label")`, &folds),
		// the plain way again, as the live layer does when its asking fails
		chromedp.Evaluate(`document.querySelector("#plain button").click(); true`, nil),
		chromedp.Sleep(100*time.Millisecond),
		chromedp.Evaluate(`(() => { const f = document.getElementById("plain"); f.requestSubmit(f.querySelector("button")); return true })()`, nil),
		chromedp.Sleep(100*time.Millisecond),
		chromedp.Evaluate(`(() => { window.dispatchEvent(new PageTransitionEvent("pageshow", {persisted: true})); return __state("#plain button") })()`, &again),
	); err != nil {
		t.Fatal(err)
	}
	if sending != "Sending…|true|true" {
		t.Errorf("while the answer is on its way: %s; want Sending…|true|true", sending)
	}
	if back != "Send|false|null|sent" {
		t.Errorf("the answer in the zone, the form kept: %s; want Send|false|null|sent", back)
	}
	if !strings.HasPrefix(folds, "Saving…|") {
		t.Errorf("a label in its own element: %s; want Saving…", folds)
	}
	if again != "Go|false|null" {
		t.Errorf("sent a second time, then back: %s; want Go|false|null — its own words", again)
	}
}

// TestPasteAndQueue (Carrel's task 88): a picture pasted into a text that
// names a form becomes a file of that form's field — words are pasted as
// words; files pasted outside a field go where dropped ones would, to the
// form of the drop sheet. That form, whose target is a zone, is sent from
// the page: the queue naming it shows a row for every file, and the answer
// is put into the zone. Without the layer the form loads its answer.
func TestPasteAndQueue(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	up := h.RegisterStream("upload", func(rw http.ResponseWriter, r *http.Request) {
		var names []string
		if err := r.ParseMultipartForm(1 << 20); err == nil {
			for _, f := range r.MultipartForm.File["file"] {
				names = append(names, f.Filename)
			}
		}
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(rw, page.Fragment{Body: []page.Part{zonePart(`<p class="got">` + strings.Join(names, ",") + " " + r.Header.Get("X-Target") + `</p>`)}}.Render())
	})
	queue := w.Queue{Zone: "queue", Progress: "up", Title: "Uploading {count}", Total: w.Mono{Text: "{total}"}, Toggle: w.Button{Label: "Hide", Small: true},
		Row: &w.QueueItem{Name: "{name}", Meta: "{size} · {percent}%", Meter: &w.Meter{Track: 100}}}
	body := `<div class="tc-root">` + w.Area{Name: "text", Label: "Text", Attach: "edit"}.Render() +
		`<form id="edit" action="/x" method="post"></form>` + w.Button{Label: "Attach", File: &w.FileField{Name: "att", Form: "edit"}}.Render() +
		w.PostForm{ID: "up", Action: up, File: "file", Target: "main"}.Render() +
		w.Scrim{ID: "drop", Closed: true, Drop: "up", Content: w.Sheet{Title: "Drop to upload"}}.Render() + queue.Render() +
		`<div data-tc-zone="main"><p>old</p></div></div>`
	doc := strings.Replace(browser.PageHTML(body, ""), `data-tc-noscript>`, `data-tc-noscript data-tc-target-header="X-Target">`, 1)
	const prep = `(() => { window.__kept = true; document.addEventListener("submit", e => { if (e.target.id === "edit") e.preventDefault(); }, true);
		window.__paste = (at, names, words) => { const dt = new DataTransfer(); names.forEach(n => dt.items.add(new File(["abc"], n))); if (words) dt.setData("text/plain", words);
			const ev = new ClipboardEvent("paste", {clipboardData: dt, bubbles: true, cancelable: true}); at.dispatchEvent(ev);
			return (ev.defaultPrevented ? "taken" : "left") + " " + [...document.querySelector('input[form="edit"]').files].map(f => f.name).join(",") }; return true })()`
	const state = `(() => { const q = document.querySelector(".tc-queue"); return document.querySelector('[data-tc-zone="main"]').textContent.trim() + " | " + (q.getClientRects().length ? "shown" : "hidden") + " " +
		q.querySelector(".tc-queue__head").textContent.trim() + " | " + [...q.querySelectorAll("li")].map(l => l.textContent.trim() + " " + l.querySelector(".tc-meter__bar").getAttribute("aria-valuenow")).join(",") + " | " + (window.__kept ? "kept" : "loaded") })()`
	var words, picture, before, after, bare string
	if err := h.Run(h.RegisterPage("paste", doc), 900, 600,
		chromedp.Evaluate(prep, nil),
		chromedp.Evaluate(`__paste(document.querySelector("textarea"), ["p.png"], "hello")`, &words),
		chromedp.Evaluate(`__paste(document.querySelector("textarea"), ["p.png"])`, &picture),
		chromedp.Evaluate(state, &before),
		chromedp.Evaluate(`__paste(document.body, ["a.txt", "b.txt"]); true`, nil),
		chromedp.Sleep(1200*time.Millisecond),
		chromedp.Evaluate(state, &after),
	); err != nil {
		t.Fatal(err)
	}
	if words != "left " || picture != "taken p.png" {
		t.Errorf("pasted into the text: with words %q, a picture alone %q; want «left », «taken p.png»", words, picture)
	}
	if before != "old | hidden Uploading {count}{total}Hide |  | kept" {
		t.Errorf("before an upload: %s; want the queue hidden and empty", before)
	}
	if after != "a.txt,b.txt main | shown Uploading 26 BHide | a.txt3 B · 100% 100,b.txt3 B · 100% 100 | kept" {
		t.Errorf("after files pasted outside a field: %s; want the answer in the zone, the queue shown with two rows done, the page kept", after)
	}
	if err := h.Run(h.RegisterPage("paste-bare", strings.Replace(doc, `<script src="/tc-live.js"></script>`, "", 1)), 900, 600,
		chromedp.Evaluate(prep, nil),
		chromedp.Evaluate(`__paste(document.body, ["a.txt"]); true`, nil),
		chromedp.Sleep(1200*time.Millisecond),
		chromedp.Evaluate(`document.body.textContent.trim() + " | " + (window.__kept ? "kept" : "loaded")`, &bare),
	); err != nil {
		t.Fatal(err)
	}
	if bare != "a.txt | loaded" {
		t.Errorf("without the layer: %q; want the answer loaded as a page", bare)
	}
}

// TestThemeShowsKept (Carrel's task 89): the theme's switch shows the choice
// kept in this browser, not the one the server printed — as radio buttons
// and as buttons; with nothing kept it shows «Auto».
func TestThemeShowsKept(t *testing.T) {
	needChrome(t)
	opts := []w.SegOption{{Label: "Auto", Value: "auto"}, {Label: "Light", Value: "light", On: true}, {Label: "Dark", Value: "dark"}}
	body := `<div class="tc-root">` + w.Seg{Name: "theme", Theme: true, Options: opts}.Render() + w.Seg{Theme: true, Options: opts}.Render() + `</div>`
	h := browser.NewHarness(t)
	const state = `document.querySelector(".tc-seg input:checked").value + " " + document.querySelector(".tc-seg button.is-on").textContent + " " + document.querySelectorAll(".tc-seg button[aria-checked]").length`
	var none, kept string
	if err := h.Run(h.RegisterPage("theme-kept", browser.PageHTML(body, "")), 900, 300,
		chromedp.Evaluate(state, &none),
		chromedp.Evaluate(`localStorage.setItem("tc-theme", "dark"); true`, nil),
		chromedp.Reload(),
		chromedp.Evaluate(state+` + " " + document.documentElement.getAttribute("data-tc-theme")`, &kept),
		chromedp.Evaluate(`localStorage.removeItem("tc-theme"); true`, nil),
	); err != nil {
		t.Fatal(err)
	}
	if none != "auto Auto 1" || kept != "dark Dark 1 dark" {
		t.Errorf("the theme's switch: with nothing kept %q, with dark kept %q; want «auto Auto 1», «dark Dark 1 dark»", none, kept)
	}
}

// TestThemeFromServer (Carrel's task 186): a theme the server chose wins over
// what this browser remembers, the switch shows it, and choosing on the
// switch changes the page without writing to this browser's memory.
func TestThemeFromServer(t *testing.T) {
	needChrome(t)
	opts := []w.SegOption{{Label: "Auto", Value: "auto"}, {Label: "Light", Value: "light"}, {Label: "Dark", Value: "dark", On: true}}
	body := `<div class="tc-root">` + w.Seg{Theme: true, Options: opts}.Render() + `</div>`
	doc := strings.Replace(browser.PageHTML(body, ""), `data-tc-noscript>`, `data-tc-noscript data-tc-theme="dark" data-tc-theme-set="dark">`, 1)
	h := browser.NewHarness(t)
	const state = `document.documentElement.getAttribute("data-tc-theme") + " " + document.querySelector(".tc-seg button.is-on").textContent + " " + localStorage.getItem("tc-theme")`
	var first, after string
	if err := h.Run(h.RegisterPage("theme-server", doc), 900, 300,
		chromedp.Evaluate(`localStorage.setItem("tc-theme", "light"); true`, nil),
		chromedp.Reload(),
		chromedp.Evaluate(state, &first),
		chromedp.Evaluate(`[...document.querySelectorAll(".tc-seg__opt")].find(b => b.textContent === "Auto").click(); true`, nil),
		chromedp.Evaluate(state, &after),
		chromedp.Evaluate(`localStorage.removeItem("tc-theme"); true`, nil),
	); err != nil {
		t.Fatal(err)
	}
	// Auto is the system's: light in the harness; the memory keeps «light»
	// (written, Auto would have removed it)
	if first != "dark Dark light" || after != "light Auto light" {
		t.Errorf("the server's dark with light remembered: %q, after choosing Auto %q; want «dark Dark light», «light Auto light» (memory untouched)", first, after)
	}
}

// TestSlugAgrees (Carrel's task 200): a field derived from another shows, as
// its placeholder, what widget.Slug derives on the server — the script's rule
// and Go's agree on every name here — and the hint's piece follows it; what
// is typed in the field wins.
func TestSlugAgrees(t *testing.T) {
	needChrome(t)
	names := []string{"Поездки", "Мой  Календарь — 2026", "Ёлка и йод", "Café Élan", "---", "  ", "Щука-Ъ", "Їжак ґ є", "日本", "Work / Home", "  Дела  "}
	form := w.Form{Action: "/x", Items: []w.FormPart{
		w.Field{Label: "Name", Parts: []w.FieldPart{w.Input{Kind: w.InputText, Name: "name"}}},
		w.Field{Label: "Address", Parts: []w.FieldPart{w.Input{Kind: w.InputText, Name: "address", Mono: true, From: "name"},
			w.FieldHint{Parts: []w.Inline{w.Mono{Text: "…/calendars/mix/"}, w.SlugOf{Of: "address", Text: "…"}, w.Mono{Text: "/"}}}}},
	}}
	page := browser.PageHTML(`<div class="tc-root">`+form.Render()+`</div>`, "")
	var want []string
	for _, n := range names {
		want = append(want, w.Slug(n))
	}
	js, _ := json.Marshal(names)
	const typ = `(n, v) => { const f = document.querySelector('[name="' + n + '"]'); f.value = v; f.dispatchEvent(new Event("input", {bubbles: true})); }`
	var got []string
	var typed string
	h := browser.NewHarness(t)
	if err := h.Run(h.RegisterPage("slug", page), 900, 400,
		chromedp.Evaluate(`(() => { const typ = `+typ+`; return `+string(js)+`.map(v => { typ("name", v); const a = document.querySelector('[name="address"]'); const p = document.querySelector("[data-tc-slug-of]").textContent; return (a.placeholder || "…") === p ? p : a.placeholder + " ≠ " + p; }); })()`, &got),
		chromedp.Evaluate(`(() => { const typ = `+typ+`; typ("name", "Поездки"); typ("address", "trips"); return document.querySelector("[data-tc-slug-of]").textContent + " " + document.querySelector('[name="address"]').placeholder; })()`, &typed),
	); err != nil {
		t.Fatal(err)
	}
	for i := range names {
		// an empty name leaves no placeholder: the piece says «…»
		w := want[i]
		if strings.TrimSpace(names[i]) == "" {
			w = "…"
		}
		if i >= len(got) || got[i] != w {
			t.Errorf("%q: the script says %q; Slug says %q", names[i], got[i], w)
		}
	}
	if typed != "trips poezdki" {
		t.Errorf("typed «trips» over the name «Поездки»: %q; want «trips poezdki»", typed)
	}
}

// TestScriptOnly (Carrel's task 90): what needs a script is shown only on a
// page that runs one (ScriptOnly), and the words that say so only on a page
// that runs none (NoScript).
func TestScriptOnly(t *testing.T) {
	needChrome(t)
	body := `<div class="tc-root">` + w.Msg{Parts: []w.MsgPart{w.Text("Drop files")}, ScriptOnly: true}.Render() +
		w.Msg{Parts: []w.MsgPart{w.Text("Needs scripts")}, NoScript: true}.Render() +
		w.Button{Label: "Paste", ScriptOnly: true}.Render() + w.Button{Label: "Upload", NoScript: true}.Render() + `</div>`
	h := browser.NewHarness(t)
	const shown = `[...document.querySelectorAll(".tc-msg, .tc-btn")].filter(e => e.getClientRects().length).map(e => e.textContent).join(",")`
	var with, bare string
	if err := h.Run(h.RegisterPage("scriptonly", browser.PageHTML(body, "")), 900, 300, chromedp.Evaluate(shown, &with)); err != nil {
		t.Fatal(err)
	}
	doc := strings.Replace(strings.Replace(browser.PageHTML(body, ""), `<script src="/tc.js"></script>`, "", 1), `<script src="/tc-live.js"></script>`, "", 1)
	if err := h.Run(h.RegisterPage("scriptonly-bare", doc), 900, 300, chromedp.Evaluate(shown, &bare)); err != nil {
		t.Fatal(err)
	}
	if with != "Drop files,Paste" || bare != "Needs scripts,Upload" {
		t.Errorf("shown with a script %q, without %q; want «Drop files,Paste», «Needs scripts,Upload»", with, bare)
	}
}

// TestSearchSuggest (Carrel's task 92): the suggestions printed under the
// search line show while it has the focus and follow what is typed — with
// nothing typed, the history and its row; with something typed, the jumps
// whose words hold it, set bold, and the row that sends it with the typed
// words in it. ↓ goes into the menu, Esc hides it, leaving the line hides it.
func TestSearchSuggest(t *testing.T) {
	needChrome(t)
	top := w.Top{Search: &w.Search{Query: w.Query{Name: "q", Label: "Search everything…"}, Action: "/find", Variant: w.SearchFull},
		Suggest: &w.Menu{Width: w.MenuWidthFull, Items: []w.MenuPart{
			w.MenuHeading{Rubric: w.Rubric{Text: "Recent searches", Form: w.RubricItem}},
			w.MenuItem{Icon: "clock", Label: "смета", Href: "#h"},
			w.MenuSep{},
			w.MenuItem{Label: "Type to search everything", Tone: w.MenuToneDusty, When: w.WhenEmpty},
			w.MenuSep{},
			w.MenuHeading{Rubric: w.Rubric{Text: "Jump to", Form: w.RubricItem}},
			w.MenuItem{Icon: "con", Label: "Анна Ковалёва", Href: "#a", When: w.WhenFound},
			w.MenuItem{Icon: "file", Label: "отчёт.pdf", Href: "#f", When: w.WhenFound},
			w.MenuSep{},
			w.MenuItem{Icon: "search", Label: "Show all results for «{q}»", When: w.WhenTyped},
		}}}
	body := `<div class="tc-root">` + top.Render() + `<button id="away">Away</button></div>`
	h := browser.NewHarness(t)
	const state = `(() => { const box = document.querySelector(".tc-top__suggest"); if (!box.getClientRects().length) return "hidden";
		const b = box.querySelector(".tc-menu__label b");
		return [...box.querySelectorAll(".tc-menu > *")].filter(e => e.getClientRects().length && !e.classList.contains("tc-menu__sep")).map(e => e.textContent.trim()).join(",") + (b ? " bold:" + b.textContent : "") })()`
	var before, empty, typed, down, esc, away string
	if err := h.Run(h.RegisterPage("suggest", browser.PageHTML(body, "")), 900, 600,
		chromedp.Evaluate(state, &before),
		chromedp.Click(`.tc-top__search input`, chromedp.ByQuery),
		chromedp.Evaluate(state, &empty),
		chromedp.SendKeys(`.tc-top__search input`, "ковал", chromedp.ByQuery),
		chromedp.Evaluate(state, &typed),
		chromedp.KeyEvent(kb.ArrowDown),
		chromedp.Evaluate(`document.activeElement.textContent.trim()`, &down),
		chromedp.KeyEvent("\x1b"),
		chromedp.Evaluate(state+` + " " + (document.activeElement.matches(".tc-q") ? "line" : "elsewhere")`, &esc),
		chromedp.Focus(`#away`, chromedp.ByQuery),
		chromedp.Evaluate(state, &away),
	); err != nil {
		t.Fatal(err)
	}
	if before != "hidden" || empty != "Recent searches,смета,Type to search everything" {
		t.Errorf("before the focus %q, with the line empty %q; want hidden, then the history and its row", before, empty)
	}
	if typed != "Recent searches,смета,Jump to,Анна Ковалёва,Show all results for «ковал» bold:Ковал" {
		t.Errorf("with «ковал» typed: %q; want the history, the jump that holds it set bold, the row that sends it", typed)
	}
	if down != "смета" || esc != "hidden line" || away != "hidden" {
		t.Errorf("↓: %q; Esc: %q; leaving: %q; want смета, the menu hidden with the focus in the line, hidden", down, esc, away)
	}
}

// TestSearchSuggestAsks: a page that names its suggestions' address
// (Document.Suggest, Carrel's task 227) prints no menu; the search line asks
// the address as it gets the focus — the history comes — and again as
// something is typed — the jumps the server found, set bold — and the
// menu stands under the line, walked by ↓ and hidden by Esc, as a printed
// one. A page without the address asks nothing.
func TestSearchSuggestAsks(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	var asked []string
	url := h.RegisterStream("suggest", func(rw http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		asked = append(asked, q)
		items := []w.MenuPart{
			w.MenuHeading{Rubric: w.Rubric{Text: "Recent searches", Form: w.RubricItem}},
			w.MenuItem{Icon: "clock", Label: "смета", Href: "#h"},
			w.MenuSep{},
			w.MenuItem{Label: "Type to search everything", Tone: w.MenuToneDusty, When: w.WhenEmpty},
		}
		if q != "" {
			items = append(items, w.MenuSep{}, w.MenuHeading{Rubric: w.Rubric{Text: "Jump to", Form: w.RubricItem}},
				w.MenuItem{Icon: "con", Label: "Анна Ковалёва", Href: "#a", When: w.WhenFound},
				w.MenuItem{Icon: "search", Label: "Show all results for «{q}»", When: w.WhenTyped})
		}
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = rw.Write([]byte(w.Menu{Width: w.MenuWidthFull, Items: items}.Render()))
	})
	top := w.Top{Search: &w.Search{Query: w.Query{Name: "q", Label: "Search everything…"}, Action: "/find"}}
	body := `<div class="tc-root">` + top.Render() + `<button id="away">Away</button></div>`
	doc := strings.Replace(browser.PageHTML(body, ""), "<html", `<html data-tc-search-suggest="`+url+`"`, 1)
	const state = `(() => { const box = document.querySelector(".tc-top__suggest"); if (!box || !box.getClientRects().length) return "hidden";
		const b = box.querySelector(".tc-menu__label b");
		return [...box.querySelectorAll(".tc-menu > *")].filter(e => e.getClientRects().length && !e.classList.contains("tc-menu__sep")).map(e => e.textContent.trim()).join(",") + (b ? " bold:" + b.textContent : "") })()`
	wait := chromedp.Sleep(600 * time.Millisecond)
	var before, empty, typed, down, esc, away string
	if err := h.Run(h.RegisterPage("search-suggest", doc), 900, 600,
		chromedp.Evaluate(state, &before),
		chromedp.Click(`.tc-top .tc-q`, chromedp.ByQuery),
		wait,
		chromedp.Evaluate(state, &empty),
		chromedp.SendKeys(`.tc-top .tc-q`, "ковал", chromedp.ByQuery),
		wait,
		chromedp.Evaluate(state, &typed),
		chromedp.KeyEvent(kb.ArrowDown),
		chromedp.Evaluate(`document.activeElement.textContent.trim()`, &down),
		chromedp.KeyEvent("\x1b"),
		chromedp.Evaluate(state+` + " " + (document.activeElement.matches(".tc-q") ? "line" : "elsewhere")`, &esc),
		chromedp.Focus(`#away`, chromedp.ByQuery),
		chromedp.Evaluate(state, &away),
	); err != nil {
		t.Fatal(err)
	}
	if before != "hidden" || empty != "Recent searches,смета,Type to search everything" {
		t.Errorf("before the focus %q, with the line empty %q; want no menu, then the history the server sent", before, empty)
	}
	if typed != "Recent searches,смета,Jump to,Анна Ковалёва,Show all results for «ковал» bold:Ковал" {
		t.Errorf("with «ковал» typed: %q; want the server's jumps, the typed set bold, the row that sends it", typed)
	}
	if down != "смета" || esc != "hidden line" || away != "hidden" {
		t.Errorf("↓: %q; Esc: %q; leaving: %q; want смета, hidden with the focus in the line, hidden", down, esc, away)
	}
	if len(asked) == 0 || asked[0] != "" || asked[len(asked)-1] != "ковал" {
		t.Errorf("asked %q; want the focus to ask with nothing typed and the pause after typing to ask «ковал»", asked)
	}
	n := len(asked)
	var none string
	if err := h.Run(h.RegisterPage("search-suggest-none", browser.PageHTML(body, "")), 900, 600,
		chromedp.Click(`.tc-top .tc-q`, chromedp.ByQuery),
		chromedp.SendKeys(`.tc-top .tc-q`, "ковал", chromedp.ByQuery),
		wait,
		chromedp.Evaluate(state, &none),
	); err != nil {
		t.Fatal(err)
	}
	if none != "hidden" || len(asked) != n {
		t.Errorf("a page without the address: %q, asked %d more; want no menu and nothing asked", none, len(asked)-n)
	}
}

// TestFileRanges: a file field with data-tc-ranges sends only the parts of
// its file the server names (Carrel's task 213): the other fields with the
// file's name and size go to the ranges address first, its lines «offset
// length» come back, and the form goes to its action with those parts,
// each called «offset-length», and no whole file. Without the live layer
// the whole file goes.
func TestFileRanges(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	var asked, got []string
	ranges := h.RegisterStream("ranges", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(1 << 20); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		_, whole := r.MultipartForm.File["file"]
		asked = append(asked, r.FormValue("csrf")+" "+r.FormValue("file.name")+" "+r.FormValue("file.size")+" "+strconv.FormatBool(whole))
		_, _ = w.Write([]byte("0 4\n10 6\n"))
	})
	action := h.RegisterStream("open", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(1 << 20); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		for _, fh := range r.MultipartForm.File["file"] {
			f, _ := fh.Open()
			b, _ := io.ReadAll(f)
			f.Close()
			got = append(got, fh.Filename+"="+string(b))
		}
		got = append(got, "size="+r.FormValue("file.size"))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(`<!doctype html><title>opened</title><p id="done">opened</p>`))
	})
	pick := w.Button{Label: "Choose a file", File: &w.FileField{Name: "file", Label: "Archive", Ranges: ranges}}
	body := `<div class="tc-root"><form action="` + action + `" method="post" enctype="multipart/form-data"><input type="hidden" name="csrf" value="x1">` +
		pick.Render() + `<button class="tc-btn is-primary" type="submit">Open the manifest</button></form></div>`
	f := filepath.Join(t.TempDir(), "carrel.zip")
	if err := os.WriteFile(f, []byte("PK0123456789abcdefghij"), 0o644); err != nil {
		t.Fatal(err)
	}
	var title string
	if err := h.Run(h.RegisterPage("ranges", browser.PageHTML(body, "")), 900, 600,
		chromedp.SetUploadFiles(`.tc-file`, []string{f}, chromedp.ByQuery),
		chromedp.Click(`button[type="submit"]`, chromedp.ByQuery),
		chromedp.WaitVisible(`#done`, chromedp.ByQuery),
		chromedp.Title(&title),
	); err != nil {
		t.Fatal(err)
	}
	if strings.Join(asked, "|") != "x1 carrel.zip 22 false" {
		t.Errorf("the ranges address was asked with %q; want the other fields, the name and the size, no file", asked)
	}
	if strings.Join(got, "|") != "0-4=PK01|10-6=89abcd|size=22" {
		t.Errorf("the action got %q; want the two parts named offset-length and the size, no whole file", got)
	}
	got = nil
	bare := strings.Replace(browser.PageHTML(body, ""), `<script src="/tc-live.js"></script>`, "", 1)
	if err := h.Run(h.RegisterPage("ranges-bare", bare), 900, 600,
		chromedp.SetUploadFiles(`.tc-file`, []string{f}, chromedp.ByQuery),
		chromedp.Click(`button[type="submit"]`, chromedp.ByQuery),
		chromedp.WaitVisible(`#done`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if strings.Join(got, "|") != "carrel.zip=PK0123456789abcdefghij|size=" {
		t.Errorf("without the layer the action got %q; want the whole file", got)
	}
}

// TestTagFieldEmptyPlace (the owner, 14.11): a press on what in a tag field is
// not a control of its own — its empty place, a tag that only shows — puts
// the cursor in its typing line, with suggestions or without: on a phone the
// line stands under the tags and a finger a little high lands beside it. A
// tag that is sent with the form stays its own control: a press unticks it.
func TestTagFieldEmptyPlace(t *testing.T) {
	needChrome(t)
	field := w.TagInput{Name: "to", Parts: []w.TagInputPart{
		w.Tag{Label: "Синк-ап по релизу", Removable: true},
		w.Tag{Label: "Анна Ковалёва", Removable: true, Value: "a.kovaleva@example.org"}},
		Add: &w.TagAdd{Name: "to.add", Placeholder: "Name or address", Label: "Add a recipient"}}
	body := `<div class="tc-root"><form action="/x" method="get"><div class="tc-field tc-field--full">` + field.Render() + `</div></form></div>`
	h := browser.NewHarness(t)
	const active = `document.activeElement && document.activeElement.classList.contains("tc-in__add")`
	const blur = `document.activeElement && document.activeElement.blur(); true`
	var x, y float64
	var onShown, onEmpty, onSent, unticked bool
	if err := h.Run(h.RegisterPage("tag-empty-place", browser.PageHTML(body, "")), 900, 600,
		chromedp.Click(`span.tc-tag--pill`, chromedp.ByQuery),
		chromedp.Evaluate(active, &onShown),
		chromedp.Evaluate(blur, nil),
		chromedp.Evaluate(`document.querySelector(".tc-in--tags").getBoundingClientRect().right - 8`, &x),
		chromedp.Evaluate(`document.querySelector(".tc-in--tags").getBoundingClientRect().top + 8`, &y),
		chromedp.ActionFunc(func(ctx context.Context) error { return chromedp.MouseClickXY(x, y).Do(ctx) }),
		chromedp.Evaluate(active, &onEmpty),
		chromedp.Evaluate(blur, nil),
		chromedp.Click(`label.tc-tag--pill`, chromedp.ByQuery),
		chromedp.Evaluate(active, &onSent),
		chromedp.Evaluate(`!document.querySelector("label.tc-tag--pill .tc-check__box").checked`, &unticked),
	); err != nil {
		t.Fatal(err)
	}
	if !onShown || !onEmpty || onSent || !unticked {
		t.Errorf("typing line after a press on a shown tag %v, on the empty place %v, on a sent tag %v (unticked %v); want true, true, false (true)", onShown, onEmpty, onSent, unticked)
	}
}

// TestTagEnter (Carrel's task 221): Enter in a tag field's line sends its
// form with the line's unseen default button — the letter's «add the
// recipient», op=add — and not with the form's first visible submit (Attach's
// «Contact»). With the library's script Enter sends nothing: the line makes
// a tag. The bite: without TagAdd.Enter the first submit goes.
func TestTagEnter(t *testing.T) {
	needChrome(t)
	field := w.TagInput{Name: "to", Add: &w.TagAdd{Name: "to.add", Placeholder: "Name or address", Label: "Add a recipient",
		Enter: &w.TagEnter{Name: "op", Value: "add", Label: "Add", Of: "letter"}}}
	body := `<div class="tc-root"><form id="letter" action="/x" method="post">` + field.Render() +
		`<button class="tc-btn" type="submit" formaction="/x?op=attach&amp;what=contact">Contact</button></form></div>`
	const catch = `document.addEventListener("submit", e => { e.preventDefault(); window.__sent = (e.submitter ? (e.submitter.name + "=" + e.submitter.value + " " + (e.submitter.getAttribute("formaction") || "")) : "none"); }); true`
	h := browser.NewHarness(t)
	for _, c := range []struct {
		name string
		live bool
		want string
	}{{"without a script", false, "op=add "}, {"with the script", true, ""}} {
		doc := browser.PageHTML(body, "")
		if !c.live {
			doc = strings.Replace(strings.Replace(doc, `<script src="/tc.js"></script>`, "", 1), `<script src="/tc-live.js"></script>`, "", 1)
		}
		var sent string
		if err := h.Run(h.RegisterPage("tag-enter-"+strconv.FormatBool(c.live), doc), 900, 600,
			chromedp.Evaluate(catch, nil),
			chromedp.SendKeys(`.tc-in__add`, "anna@example.org\r", chromedp.ByQuery),
			chromedp.Sleep(500*time.Millisecond),
			chromedp.Evaluate(`window.__sent || ""`, &sent),
		); err != nil {
			t.Fatal(err)
		}
		if sent != c.want {
			t.Errorf("%s: Enter sent %q; want %q", c.name, sent, c.want)
		}
	}
}

// TestRangesNextStep (Carrel's task 220, Р-71): the restore from this
// computer keeps its file for the next step. Step 1 sends the parts the
// server names; its form names a zone, so step 2 comes in place and the page
// stays. Step 2's form (Dialog.Ranges, RangesOf) asks the same address with
// its own fields — the chosen collections — and sends the parts named then,
// of the file step 1 kept, never the whole file. The bite: without the
// layer's keeping, step 2 sends no parts.
func TestRangesNextStep(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	var asked, got []string
	ranges := h.RegisterStream("ranges2", func(rw http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(1 << 20); err != nil {
			http.Error(rw, err.Error(), 400)
			return
		}
		asked = append(asked, r.FormValue("file.name")+" "+strings.Join(r.MultipartForm.Value["collection"], ","))
		if len(r.MultipartForm.Value["collection"]) > 0 {
			_, _ = rw.Write([]byte("10 6\n"))
			return
		}
		_, _ = rw.Write([]byte("0 4\n"))
	})
	parts := func(r *http.Request) string {
		var ps []string
		for _, fh := range r.MultipartForm.File["file"] {
			f, _ := fh.Open()
			b, _ := io.ReadAll(f)
			f.Close()
			ps = append(ps, fh.Filename+"="+string(b))
		}
		return strings.Join(ps, ",")
	}
	var match string
	open := h.RegisterStream("open2", func(rw http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(1 << 20); err != nil {
			http.Error(rw, err.Error(), 400)
			return
		}
		got = append(got, "open "+parts(r))
		step2 := w.Dialog{Action: match, Ranges: ranges, RangesOf: "file", Items: []w.DialogPart{
			w.Form{Items: []w.FormPart{w.Field{Parts: []w.FieldPart{w.Check{Name: "collection", Value: "nextcloud/personal", On: true, Parts: []w.CheckPart{w.Text("personal")}}}}}},
			w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Restore", Kind: w.ButtonPrimary, Submit: true}}}}}
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = io.WriteString(rw, step2.Render())
	})
	match = h.RegisterStream("match2", func(rw http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(1 << 20); err != nil {
			http.Error(rw, err.Error(), 400)
			return
		}
		got = append(got, "match "+strings.Join(r.MultipartForm.Value["collection"], ",")+" "+parts(r))
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = io.WriteString(rw, `<p id="done">restored</p>`)
	})
	page := func(live bool) string {
		pick := w.Button{Label: "Choose a file", File: &w.FileField{Name: "file", Label: "Archive", Ranges: ranges}}
		step1 := w.Dialog{Action: open, Items: []w.DialogPart{
			w.Form{Items: []w.FormPart{w.Field{Parts: []w.FieldPart{pick}}}},
			w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Open the manifest", Kind: w.ButtonPrimary, Submit: true}}}}}
		main := w.Main{Zone: "restore", Target: "restore", Items: []w.MainPart{step1}}
		doc := browser.PageHTML(`<div class="tc-root">`+main.Render()+`</div>`, "")
		if !live {
			doc = strings.Replace(doc, `<script src="/tc-live.js"></script>`, "", 1)
		}
		return doc
	}
	f := filepath.Join(t.TempDir(), "carrel.zip")
	if err := os.WriteFile(f, []byte("PK0123456789abcdefghij"), 0o644); err != nil {
		t.Fatal(err)
	}
	var kept bool
	if err := h.Run(h.RegisterPage("ranges-next", page(true)), 900, 600,
		chromedp.Evaluate(`window.__stayed = true; true`, nil),
		chromedp.SetUploadFiles(`.tc-file`, []string{f}, chromedp.ByQuery),
		chromedp.Click(`button[type="submit"]`, chromedp.ByQuery),
		chromedp.WaitVisible(`input[name="collection"]`, chromedp.ByQuery),
		chromedp.Evaluate(`window.__stayed === true`, &kept),
		chromedp.Click(`button[type="submit"]`, chromedp.ByQuery),
		chromedp.WaitVisible(`#done`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if !kept {
		t.Error("step 2 loaded a page; want it in place, the page kept")
	}
	if strings.Join(asked, "|") != "carrel.zip |carrel.zip nextcloud/personal" {
		t.Errorf("the ranges address was asked %q; want step 1 alone, then with the chosen collection", asked)
	}
	if strings.Join(got, "|") != "open 0-4=PK01|match nextcloud/personal 10-6=89abcd" {
		t.Errorf("the steps got %q; want the parts named each time, never the whole file", got)
	}
}

// TestFileTaken: a file field with data-tc-taken asks which picked names
// are taken before the files go (Carrel's task 214): the taken one gets the
// queue's question row, nothing is sent until it is answered; Replace sends
// both with «taken.<name>=replace» before them (Carrel's task 219), Skip
// sends the other alone.
func TestFileTaken(t *testing.T) {
	needChrome(t)
	h := browser.NewHarness(t)
	var asked, got []string
	taken := h.RegisterStream("taken", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(1 << 20); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		asked = append(asked, r.FormValue("csrf")+" "+strings.Join(r.MultipartForm.Value["name"], ","))
		_, _ = w.Write([]byte("b.txt\n"))
	})
	upload := h.RegisterStream("upload", func(w http.ResponseWriter, r *http.Request) {
		// the body read as a stream, as the server reads it: an answer that
		// comes after its file is too late (Carrel's task 219)
		mr, err := r.MultipartReader()
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		var names []string
		answer := ""
		for {
			part, err := mr.NextPart()
			if err != nil {
				break
			}
			switch part.FormName() {
			case "file":
				names = append(names, part.FileName())
			case "taken.b.txt":
				v, _ := io.ReadAll(part)
				answer = string(v)
				if len(names) > 0 {
					answer += " (after a file)"
				}
			}
		}
		got = append(got, strings.Join(names, ",")+" "+answer)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(`<!doctype html><p id="done">uploaded</p>`))
	})
	q := w.Queue{Title: "Uploading", Total: w.Mono{Text: "", Variant: w.MonoDusty, Place: w.Pushed}, Toggle: w.Button{Label: "Hide", Small: true}, Progress: "up",
		Row: &w.QueueItem{Name: "{name}", Meta: "{size}"},
		Taken: &w.QueueItem{Name: "{name}", Meta: "name taken", Failed: true, Acts: []w.Button{
			{Label: "Keep both", Small: true, Name: "answer", Value: "keep"}, {Label: "Replace", Small: true, Name: "answer", Value: "replace"}, {Label: "Skip", Small: true, Name: "answer", Value: "skip"}}}}
	body := `<div class="tc-root"><form id="up" action="` + upload + `" method="post" enctype="multipart/form-data"><input type="hidden" name="csrf" value="x1">` +
		`<input class="tc-file" type="file" name="file" multiple aria-label="Files" data-tc-taken="` + taken + `"><button class="tc-btn" type="submit">Upload</button></form>` + q.Render() + `</div>`
	dir := t.TempDir()
	a, b := filepath.Join(dir, "a.txt"), filepath.Join(dir, "b.txt")
	for _, f := range []string{a, b} {
		if err := os.WriteFile(f, []byte("x"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	for _, c := range []struct{ answer, want string }{{"replace", "a.txt,b.txt replace"}, {"skip", "a.txt "}} {
		asked, got = nil, nil
		var rows int
		if err := h.Run(h.RegisterPage("taken-"+c.answer, browser.PageHTML(body, "")), 900, 600,
			chromedp.SetUploadFiles(`.tc-file`, []string{a, b}, chromedp.ByQuery),
			chromedp.Click(`button[type="submit"]`, chromedp.ByQuery),
			chromedp.WaitVisible(`.tc-queue__list button[value="`+c.answer+`"]`, chromedp.ByQuery),
			chromedp.Evaluate(`document.querySelectorAll(".tc-queue__list li").length`, &rows),
			chromedp.ActionFunc(func(ctx context.Context) error {
				if len(got) != 0 || rows != 1 {
					t.Errorf("%s: before the answer %d rows and %q sent; want one question and nothing sent", c.answer, rows, got)
				}
				return nil
			}),
			chromedp.Click(`.tc-queue__list button[value="`+c.answer+`"]`, chromedp.ByQuery),
			chromedp.WaitVisible(`#done`, chromedp.ByQuery),
		); err != nil {
			t.Fatal(err)
		}
		if strings.Join(asked, "|") != "x1 a.txt,b.txt" {
			t.Errorf("%s: the taken address was asked %q; want the hidden fields and both names", c.answer, asked)
		}
		if strings.Join(got, "|") != c.want {
			t.Errorf("%s: the upload got %q; want %q", c.answer, got, c.want)
		}
	}
}
