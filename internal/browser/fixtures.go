package browser

import (
	"strings"

	w "carrel.dev/typecase/widget"
)

// root wraps a fixture in the box every frame of the mockup has: the widgets'
// CSS is written for a tc-root, its fonts, its container.
func root(body string) string { return `<div class="tc-root">` + body + `</div>` }

// dropMenu is a button with its menu, closed until pressed.
func dropMenu(label string) string {
	return w.Drop{
		Trigger: w.Button{Label: label, Menu: true},
		Menu: w.Menu{Pop: w.MenuUnder, Width: w.MenuWidth212, Items: []w.MenuPart{
			w.MenuItem{Label: "Rename", Icon: "pencil"}, w.MenuItem{Label: "Export", Icon: "down"},
			w.MenuSep{}, w.MenuItem{Label: "Delete", Icon: "trash", Tone: w.MenuToneDanger}}},
	}.Render()
}

// GoodFixtures returns passing L-9.2 page bodies keyed by check name.
func GoodFixtures() map[Name]Fixture {
	longLabel := strings.Repeat("Export everything now ", 8)

	bar := w.Bar{Items: []w.BarPart{
		w.Filter{Query: w.Query{Value: "Filter this folder", Label: "Filter this folder"}, Waiting: true, Variant: w.FilterH26W190},
		w.Seg{Options: []w.SegOption{{Label: "Name", On: true}, {Label: "Size"}, {Label: "Changed"}}, Secondary: true},
		w.Button{Label: "Columns", Icon: "cols", Secondary: true},
		w.Button{Label: "Refresh", Icon: "refresh", Secondary: true},
	}, More: &w.Button{Icon: "dots", Title: "More"},
		Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "14 items · 38.2 MB"}}, Secondary: true}}.Render()

	list := w.List{Rows: []w.ListRow{w.TimeRow{Mark: w.Mark{Colour: "#4A6B52"}, Time: "10:00",
		Title:  []w.Inline{w.Text("A very long title of an event that must truncate in its column without overflowing the row")},
		Detail: "Room 2", Source: "nextcloud · Personal"}}}.Render()

	// Every small control in a phone's shell: a finger's zone of 44 px is the
	// mockup's rule for the drawer state (@container style(--tc-shell-mode:
	// drawer)), so they stand where that state applies.
	taps := w.Shell{Body: w.ShellBody{Main: w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Text: "Tasks"}, Acts: &w.Acts{Items: []w.ActsPart{w.Close{Label: "Close"}}}},
		w.List{Rows: []w.ListRow{w.DueRow{Tick: w.Tick{Label: "Done"}, Mark: w.Mark{Colour: "#4A6B52"}, Title: []w.Inline{w.Text("Pack")}, Due: "12 Aug", Progress: "—", Source: "nextcloud"}}},
		w.RailSec{Items: []w.RailSecPart{w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Personal", Meta: "7"}}},
		w.Form{Items: []w.FormPart{w.Check{On: true, Parts: []w.CheckPart{w.Text("I agree")}}}},
		w.BottomNav{Items: []w.BottomItem{{Icon: "con", Label: "Contacts", Href: "#"}, {Icon: "cal", Label: "Calendar", Href: "#"}}, On: 1},
	}}, Rail: []w.RailPart{w.Nav{Items: []w.NavItem{{Icon: "con", Label: "Contacts", Href: "#"}}, On: 1}}}}.Render()

	// A sheet and, over it, a second sheet: Esc closes the upper one only.
	sheet := func(id, title string) string {
		return `<div id="` + id + `">` + w.Scrim{Content: w.Sheet{Title: title, Body: []w.SheetPart{w.Dim{Parts: []w.Inline{w.Text("Under")}}},
			Foot: []w.SheetFootPart{w.Button{Label: "OK", Kind: w.ButtonPrimary}}}}.Render() + `</div>`
	}
	layers := sheet("bottom-sheet", "Sheet") + sheet("top-dialog", "Confirm")

	sample := `<p class="tc-sample">Sample text</p>`
	sampleCSS := `.tc-sample { display: inline-block; width: 180px; padding: 8px; background: var(--tc-surface); color: var(--tc-ink); }`

	return map[Name]Fixture{
		MenuLabelArrow: {Body: root(w.Drop{Trigger: w.Button{Label: longLabel, Menu: true}, Menu: w.Menu{Pop: w.MenuUnder,
			Items: []w.MenuPart{w.MenuItem{Label: "One"}}}}.Render()), Width: 220, Extra: `.demo { width: 200px; }`},
		PopoverAnchor: {Body: root(dropMenu("Actions")), Width: 800, Height: 600},
		PopoverFlip:   {Body: root(dropMenu("Actions")), Width: 400, Height: 300, Extra: `.demo { margin-top: 240px; }`},
		Ellipsis:      {Body: root(list), Width: 320, Extra: `.demo { width: 280px; }`},
		BarOneRow:     {Body: root(bar), Width: 300, Extra: `.demo { width: 280px; }`},
		TapTarget: {Body: taps, Width: 375, Height: 700,
			// Room around the controls: the probe asks each one about its own zone.
			Extra: `.tc-main > * { margin: 30px; }`},
		FocusReturn:   {Body: root(dropMenu("Actions"))},
		EscLayers:     {Body: root(layers), Width: 800, Height: 600},
		Contrast:      {Body: sample, Extra: sampleCSS},
		ThemeGeometry: {Body: sample, Extra: sampleCSS},
	}
}
