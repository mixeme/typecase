package proof

import (
	"strconv"
	"strings"
)

// Overlay is .m-scrim with one sheet, over the shell. Top26 is the scrim offset
// the confirmation frames use.
type Overlay struct {
	Top26 bool
	Sheet Sheet
}

func (o Overlay) Render() string {
	c := "tcx-scrim"
	if o.Top26 {
		c += " is-top-26"
	}
	return `<div class="` + c + `">` + o.Sheet.Render() + `</div>`
}

// SheetWidth is the sheet width a frame sets.
type SheetWidth string

const (
	SheetDefault SheetWidth = ""
	Sheet560     SheetWidth = "tcx-sheet-560"
	Sheet480     SheetWidth = "tcx-sheet-480"
)

// Sheet is .m-sheet: head (.h), body (.b), foot (.f).
type Sheet struct {
	Width SheetWidth
	Icon  Icon
	Title string
	Pick  *Pick
	Meta  string
	Close bool
	Body  Nodes
	Foot  Nodes
}

func (s Sheet) Render() string {
	c := "tcx-sheet"
	if s.Width != "" {
		c += " " + string(s.Width)
	}
	var b strings.Builder
	b.WriteString(`<div class="` + c + `" role="dialog" aria-modal="true" aria-label="` + esc(s.Title) + `"><div class="h">`)
	b.WriteString(s.Icon.Render() + `<h2 class="tcx-sheet-title">` + esc(s.Title) + `</h2>`)
	if s.Pick != nil {
		b.WriteString(s.Pick.Render())
	}
	if s.Meta != "" {
		b.WriteString(`<span class="tcx-sheet-meta">` + esc(s.Meta) + `</span>`)
	}
	if s.Close {
		b.WriteString(Btn{Icon: "close", Close: true, Title: "Close"}.Render())
	}
	b.WriteString(`</div><div class="b">` + s.Body.Render() + `</div><div class="f">` + s.Foot.Render() + `</div></div>`)
	return b.String()
}

// Pick is .m-in.is-pick in a sheet head: where the new record goes.
type Pick struct{ Label string }

func (p Pick) Render() string {
	return `<button type="button" class="tcx-in is-pick tcx-pick-head" aria-haspopup="listbox">` + esc(p.Label) + `</button>`
}

// Kbd is .m-kbd.
type Kbd string

func (k Kbd) Render() string { return `<kbd class="tcx-kbd">` + esc(string(k)) + `</kbd>` }

// FootEnd is the group a sheet foot pushes to its end.
type FootEnd Nodes

func (f FootEnd) Render() string {
	return `<span class="tcx-foot-end">` + Nodes(f).Render() + `</span>`
}

// Impact is one thing that points at what is being deleted.
type Impact struct {
	What   string
	Effect string
}

// Impacts is the list of what points at it, in a confirmation.
type Impacts struct {
	Title string
	Items []Impact
}

func (im Impacts) Render() string {
	var b strings.Builder
	b.WriteString(`<div><div class="tcx-rubric tcx-rubric-6">` + esc(im.Title) + `</div><div class="tcx-impacts">`)
	for _, it := range im.Items {
		b.WriteString(`<span><b>` + esc(it.What) + `</b> ` + Inline{Dim, it.Effect}.Render() + `</span>`)
	}
	b.WriteString(`</div></div>`)
	return b.String()
}

// Line is a row of actions and a note inside a sheet body.
type Line Nodes

func (l Line) Render() string { return `<div class="tcx-line">` + Nodes(l).Render() + `</div>` }

// Menu is .m-menu.is-pop.
type Menu struct {
	Right bool
	Wide  bool // the frame sets 222 px
	Items []MenuItem
}

// MenuItem is one entry; Sep is the rule between groups.
type MenuItem struct {
	Icon   string
	Label  Nodes
	Danger bool
	Sep    bool
}

func (m Menu) Render() string {
	c := "tcx-menu is-pop"
	if m.Right {
		c += " is-right"
	}
	if m.Wide {
		c += " tcx-menu-222"
	}
	var b strings.Builder
	b.WriteString(`<div class="` + c + `" role="menu">`)
	for _, it := range m.Items {
		if it.Sep {
			b.WriteString(`<span class="sep" role="separator"></span>`)
			continue
		}
		cls := ""
		if it.Danger {
			cls = ` class="tcx-danger-item"`
		}
		b.WriteString(`<button type="button" role="menuitem"` + cls + `>` + Icon{Name: it.Icon}.Render() + it.Label.Render() + `</button>`)
	}
	b.WriteString(`</div>`)
	return b.String()
}

// Card is .m-card: one server or feed.
type Card struct {
	Name    string
	Badge   string
	Actions []Btn
	URL     string
	Alert   bool
	Table   *Table
	Diag    string
	Note    string
}

func (c Card) Render() string {
	cls, rub, url := "tcx-card", "tcx-rubric", "mono tcx-url"
	if c.Alert {
		cls += " tcx-card-alert"
		rub += " tcx-alert-text"
	}
	if c.Diag != "" {
		url += " tcx-url-gap"
	}
	var b strings.Builder
	b.WriteString(`<section class="` + cls + `"><div class="` + rub + `">` + esc(c.Name))
	if c.Badge != "" {
		b.WriteString(` ` + Badge{Label: c.Badge}.Render())
	}
	b.WriteString(`<span class="tcx-card-acts">` + btns(c.Actions) + `</span></div>`)
	b.WriteString(`<div class="` + url + `">` + esc(c.URL) + `</div>`)
	if c.Table != nil {
		b.WriteString(c.Table.Render())
	}
	if c.Diag != "" {
		b.WriteString(Diag(c.Diag).Render())
	}
	if c.Note != "" {
		b.WriteString(`<div class="tcx-card-note">` + esc(c.Note) + `</div>`)
	}
	b.WriteString(`</section>`)
	return b.String()
}

// Diag is .m-diag: what the server answered, verbatim.
type Diag string

func (d Diag) Render() string { return `<div class="tcx-diag">` + esc(string(d)) + `</div>` }

// ListRubric is a rubric with a count heading a run of cards.
type ListRubric struct {
	Text  string
	Count int
}

func (l ListRubric) Render() string {
	return `<div class="tcx-rubric tcx-list-rubric">` + esc(l.Text) + Inline{MonoFlat, strconv.Itoa(l.Count)}.Render() + `</div>`
}

// Stack is the padded column the settings body sits in.
type Stack Nodes

func (s Stack) Render() string { return `<div class="tcx-stack">` + Nodes(s).Render() + `</div>` }

// Pad is the gutter-only wrapper of a settings body.
type Pad Nodes

func (p Pad) Render() string { return `<div class="tcx-pad">` + Nodes(p).Render() + `</div>` }

// Poll is .m-poll: how the sources answered.
type Poll struct {
	Percent int
	Body    Nodes
}

func (p Poll) Render() string {
	return `<div class="tcx-poll"><span class="tcx-prog" role="progressbar" aria-valuenow="` + strconv.Itoa(p.Percent) +
		`"><i data-tc-width="` + strconv.Itoa(p.Percent) + `%"></i></span>` + p.Body.Render() + `</div>`
}

// AgendaRow is .m-row--agenda.
type AgendaRow struct {
	Mark   Mark
	Time   string
	Title  string
	Place  string
	Badge  string
	Source string
	On     bool
}

func (r AgendaRow) Render() string {
	t := "t"
	if r.On {
		t += " is-strong"
	}
	badge := ""
	if r.Badge != "" {
		badge = Badge{Label: r.Badge}.Render()
	}
	return `<div class="tcx-row tcx-row--agenda` + on(r.On) + `" role="listitem">` + r.Mark.Render() +
		`<span class="mono">` + esc(r.Time) + `</span>` +
		`<span class="` + t + `">` + esc(r.Title) + `</span>` +
		`<span class="s tcx-hide-3">` + esc(r.Place) + `</span>` +
		`<span class="d">` + badge + `</span>` +
		`<span class="d">` + esc(r.Source) + `</span></div>`
}
