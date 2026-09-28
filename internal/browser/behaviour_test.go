package browser_test

import (
	"fmt"
	"os"
	"testing"

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
