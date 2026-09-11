package proof

import (
	"strconv"
	"strings"
)

// Node is anything a widget can hold.
type Node interface{ Render() string }

// Nodes renders in order.
type Nodes []Node

func (ns Nodes) Render() string {
	var b strings.Builder
	for _, n := range ns {
		if n != nil {
			b.WriteString(n.Render())
		}
	}
	return b.String()
}

// Text is escaped plain text.
type Text string

func (t Text) Render() string { return esc(string(t)) }

// Inline is a run of text with one of the frame's text treatments.
type Inline struct {
	Kind InlineKind
	Text string
}

// InlineKind names a text treatment the frames use.
type InlineKind string

const (
	Mono      InlineKind = "mono"
	Dim       InlineKind = "tcx-dim"
	Ref       InlineKind = "tcx-ref"
	Strong    InlineKind = "tcx-strong"
	MonoMuted InlineKind = "mono tcx-muted"
	Alert     InlineKind = "tcx-alert-text"
	SmallDim  InlineKind = "tcx-small-dim"
	Mono11    InlineKind = "mono tcx-mono-11"
	MonoFlat  InlineKind = "mono tcx-ls0"
	Bold      InlineKind = "b"
)

func (i Inline) Render() string {
	if i.Kind == Bold {
		return `<b>` + esc(i.Text) + `</b>`
	}
	return `<span class="` + string(i.Kind) + `">` + esc(i.Text) + `</span>`
}

// Icon is one sprite icon. Box prints the 24 viewBox; Size and Stroke print the
// presentational attributes some frames put on the svg.
type Icon struct {
	Name   string
	Box    bool
	Size   int
	Stroke string
	Caps   bool
	Class  string
	Tint   Tint // stroke colour the frame gives this icon
}

// Tint is an icon stroke colour set by a frame.
type Tint string

const (
	TintAccent Tint = "tcx-tint-accent"
	TintAlert  Tint = "tcx-tint-alert"
)

const sprite = "/mock-sprite.svg#"

func (i Icon) Render() string {
	var b strings.Builder
	b.WriteString(`<svg`)
	cls := strings.TrimSpace(i.Class + " " + string(i.Tint))
	if cls != "" {
		b.WriteString(` class="` + cls + `"`)
	}
	if i.Size > 0 {
		s := strconv.Itoa(i.Size)
		b.WriteString(` width="` + s + `" height="` + s + `"`)
	}
	if i.Box || i.Size > 0 {
		b.WriteString(` viewBox="0 0 24 24"`)
	}
	if i.Stroke != "" {
		if i.Tint == "" {
			b.WriteString(` stroke="currentColor"`)
		}
		b.WriteString(` stroke-width="` + i.Stroke + `" fill="none"`)
	}
	if i.Caps {
		b.WriteString(` stroke-linecap="round" stroke-linejoin="round"`)
	}
	b.WriteString(` aria-hidden="true"><use href="` + sprite + "i-" + i.Name + `"/></svg>`)
	return b.String()
}

// Btn is .m-btn.
type Btn struct {
	Label       string
	Icon        string
	Chev        bool
	Primary     bool
	Small       bool
	DangerSolid bool
	Close       bool
	More        bool
	Pressed     bool
	Title       string
	Lbl         bool  // label hides at 800 in a selection bar
	Danger      bool  // outline danger
	Focus       bool  // the focus ring a frame draws on the button that opened a sheet
	Narrow      bool  // icon button in a table row: the frame trims its padding
	Rich        Nodes // label with inline parts, instead of Label
	Menu        *Menu // open menu; printed as .tcx-drop: a menu is not content of a button
	opens       bool
}

func (b Btn) Render() string {
	if b.Menu != nil {
		m := *b.Menu
		b.Menu, b.opens = nil, true
		return `<span class="tcx-drop">` + b.Render() + m.Render() + `</span>`
	}
	c := []string{"tcx-btn"}
	add := func(on bool, cls string) {
		if on {
			c = append(c, cls)
		}
	}
	add(b.Small, "is-sm")
	add(b.Primary, "is-primary")
	add(b.DangerSolid, "is-danger-solid")
	add(b.Danger, "is-danger")
	add(b.Close, "tcx-close")
	add(b.More, "tcx-more")
	add(b.Pressed, "is-pressed")
	add(b.Focus, "is-focus-ring")
	add(b.Narrow, "is-narrow-pad")
	var s strings.Builder
	s.WriteString(`<button type="button" class="` + strings.Join(c, " ") + `"`)
	name := b.Title
	if name == "" {
		name = b.Label
	}
	if b.Title != "" {
		s.WriteString(` title="` + esc(b.Title) + `"`)
	}
	if (b.Label == "" && b.Rich == nil) || b.Lbl {
		s.WriteString(` aria-label="` + esc(name) + `"`)
	}
	if b.Pressed {
		s.WriteString(` aria-pressed="true"`)
	}
	if b.opens {
		s.WriteString(` aria-haspopup="menu" aria-expanded="true"`)
	}
	s.WriteString(`>`)
	if b.Icon != "" {
		s.WriteString(Icon{Name: b.Icon, Box: true}.Render())
	}
	switch {
	case b.Rich != nil:
		s.WriteString(b.Rich.Render())
	case b.Lbl:
		s.WriteString(`<span class="lbl">` + esc(b.Label) + `</span>`)
	default:
		s.WriteString(esc(b.Label))
	}
	if b.Chev {
		s.WriteString(Icon{Name: "chev", Box: true}.Render())
	}
	s.WriteString(`</button>`)
	return s.String()
}

// Tag is .m-tag.
type Tag struct {
	Label string
	On    bool
}

func (t Tag) Render() string {
	return `<span class="tcx-tag` + on(t.On) + `">` + esc(t.Label) + `</span>`
}

// Badge is .m-badge. Linked prints the link glyph the note list uses.
type Badge struct {
	Label  string
	Linked bool
}

func (b Badge) Render() string {
	if b.Linked {
		return `<span class="tcx-badge is-linked">` +
			Icon{Name: "link", Size: 9, Stroke: "2", Class: "tcx-badge-ico"}.Render() + ` ` + esc(b.Label) + `</span>`
	}
	return `<span class="tcx-badge">` + esc(b.Label) + `</span>`
}

// Top is .m-top: brand, search, the right wing.
type Top struct {
	Search  string
	Actions Nodes
	User    string
}

func (t Top) Render() string {
	return `<header class="tcx-top"><div class="tcx-brand">` +
		`<svg class="tcx-logo" viewBox="4.79 -31.36 291.08 116.63" role="img" aria-label="Carrel">` +
		`<use href="` + sprite + `logo-word" class="wd"/><use href="` + sprite + `logo-shade" class="sh"/></svg></div>` +
		`<div class="tcx-search" role="search">` + Icon{Name: "search", Size: 13, Stroke: "1.5"}.Render() +
		`<input class="tcx-q" type="search" placeholder="` + esc(t.Search) + `" aria-label="Search"></div>` +
		`<div class="tcx-top-right">` + t.Actions.Render() +
		`<button type="button" class="tcx-user">` + esc(t.User) + Icon{Name: "chev", Size: 12, Stroke: "1.5"}.Render() + `</button>` +
		`</div></header>`
}

// SideSize is the side column width the frame sets.
type SideSize string

const (
	SideNormal SideSize = ""
	SideWide   SideSize = "is-side-440"
	SideNarrow SideSize = "is-side-340"
)

// Screen is .mock: top bar and the shell, with an optional sheet over it.
type Screen struct {
	Top     Top
	Rail    Rail
	Main    Nodes
	Side    *Side
	Size    SideSize
	Tall    bool // the frame holds the shell at 520
	Overlay *Overlay
	Dark    bool
}

func (s Screen) Render() string {
	c := "tcx-shell"
	if s.Side != nil {
		c += " has-side"
	}
	if s.Size != "" {
		c += " " + string(s.Size)
	}
	if s.Tall {
		c += " is-tall-520"
	}
	side := ""
	if s.Side != nil {
		side = s.Side.Render()
	}
	shell := `<div class="` + c + `">` + s.Rail.Render() + `<main class="tcx-main">` + s.Main.Render() + `</main>` + side + `</div>`
	if s.Overlay != nil {
		shell = `<div class="tcx-overlay">` + shell + s.Overlay.Render() + `</div>`
	}
	return `<div class="tcx">` + s.Top.Render() + shell + `</div>`
}

// Mark is the colour bar before a title. Ext marks a subscribed collection.
type Mark struct {
	Colour string
	Multi  bool
	Ext    bool
}

func (m Mark) Render() string {
	if m.Multi {
		return `<span class="tcx-bar3 is-multi"></span>`
	}
	c := "tcx-bar3"
	if m.Ext {
		c += " is-ext"
	}
	return `<span class="` + c + `" data-tc-colour="` + esc(m.Colour) + `"></span>`
}

// Crumb is one step of .m-crumbs; the last one is the current folder.
type Crumb struct {
	Label string
	Href  string
}

// TitleSize is the heading size a frame gives a panel title.
type TitleSize string

const (
	TitleScreen TitleSize = ""
	TitlePanel  TitleSize = "is-panel"
	TitleItem   TitleSize = "is-item"
)

// Head is .m-head.
type Head struct {
	Mark   *Mark
	Title  string
	Size   TitleSize
	Crumbs []Crumb
	Sub    Nodes
	Acts   Nodes
	Tight  bool
	Mark19 bool // the dark calendar frame draws the title mark at 19, not 20
}

func (h Head) Render() string {
	c := "tcx-head"
	if h.Tight {
		c += " is-snug"
	}
	h1 := "tcx-h1"
	if h.Size != "" {
		h1 += " " + string(h.Size)
	}
	if h.Mark19 {
		h1 += " is-mark-19"
	}
	var b strings.Builder
	b.WriteString(`<div class="` + c + `"><div><h2 class="` + h1 + `">`)
	if h.Mark != nil {
		b.WriteString(h.Mark.Render())
	}
	b.WriteString(esc(h.Title) + `</h2>`)
	if len(h.Crumbs) > 0 {
		b.WriteString(`<nav class="tcx-crumbs" aria-label="Path">`)
		for i, cr := range h.Crumbs {
			if i > 0 {
				b.WriteString("/")
			}
			if i == len(h.Crumbs)-1 {
				b.WriteString(`<b aria-current="page">` + esc(cr.Label) + `</b>`)
			} else {
				b.WriteString(`<a href="` + esc(cr.Href) + `">` + esc(cr.Label) + `</a>`)
			}
		}
		b.WriteString(`</nav>`)
	}
	b.WriteString(`<div class="tcx-sub">` + h.Sub.Render() + `</div></div>`)
	b.WriteString(`<div class="tcx-acts">` + h.Acts.Render() + `</div></div>`)
	return b.String()
}

// Bar is .m-bar; Sel is the selection bar.
type Bar struct {
	Sel   bool
	Items Nodes
}

func (b Bar) Render() string {
	c := "tcx-bar"
	if b.Sel {
		c += " tcx-sel"
	}
	return `<div class="` + c + `">` + b.Items.Render() + `</div>`
}

// FilterWidth is the width a frame gives a bar filter.
type FilterWidth string

const (
	Filter200 FilterWidth = "is-w200"
	Filter190 FilterWidth = "is-w190"
)

// Filter is the filter field of a bar.
type Filter struct {
	Placeholder string
	Width       FilterWidth
	Icon        bool
}

func (f Filter) Render() string {
	c := "tcx-in is-ph " + string(f.Width)
	if !f.Icon {
		return `<input class="` + c + `" type="search" placeholder="` + esc(f.Placeholder) + `" aria-label="` + esc(f.Placeholder) + `">`
	}
	return `<label class="` + c + `">` + Icon{Name: "search", Size: 13, Stroke: "1.5", Class: "tcx-in-ico"}.Render() +
		`<input class="tcx-q" type="search" placeholder="` + esc(f.Placeholder) + `" aria-label="` + esc(f.Placeholder) + `"></label>`
}

// Seg is .m-seg: a pressed-one-of set.
type Seg struct {
	Label string
	Items []SegItem
	Is2nd bool
}

// SegItem is one option; Icon replaces the label on screen.
type SegItem struct {
	Label string
	Icon  string
	On    bool
}

func (s Seg) Render() string {
	c := "tcx-seg"
	if s.Is2nd {
		c += " is-2nd"
	}
	var b strings.Builder
	b.WriteString(`<div class="` + c + `" role="group" aria-label="` + esc(s.Label) + `">`)
	for _, it := range s.Items {
		b.WriteString(`<button type="button"` + onAttr(it.On) + ` aria-pressed="` + strconv.FormatBool(it.On) + `"`)
		if it.Icon != "" {
			b.WriteString(` aria-label="` + esc(it.Label) + `">` + Icon{Name: it.Icon, Size: 13, Stroke: "1.4"}.Render())
		} else {
			b.WriteString(`>` + esc(it.Label))
		}
		b.WriteString(`</button>`)
	}
	b.WriteString(`</div>`)
	return b.String()
}

// Right is .m-right: what a bar pushes to its end.
type Right struct {
	Items Nodes
	Is2nd bool
}

func (r Right) Render() string {
	c := "tcx-right"
	if r.Is2nd {
		c += " is-2nd"
	}
	return `<span class="` + c + `">` + r.Items.Render() + `</span>`
}

// Sep is .m-sep.
type Sep struct{}

func (Sep) Render() string { return `<span class="tcx-sep"></span>` }

// Group is .m-group: the rubric above a run of rows.
type Group struct {
	Rubric string
	Count  int
}

func (g Group) Render() string {
	return `<div class="tcx-group"><span class="tcx-rubric">` + esc(g.Rubric) + `</span><span class="tcx-num">` + strconv.Itoa(g.Count) + `</span></div>`
}

// List is .m-list.
type List struct{ Rows Nodes }

func (l List) Render() string {
	return `<div class="tcx-list" role="list">` + l.Rows.Render() + `</div>`
}

// NoteRow is .m-row--note.
type NoteRow struct {
	Mark    Mark
	Title   string
	Excerpt string
	Tags    []Tag
	Badge   *Badge
	Date    string
	Source  string
	On      bool
}

func (r NoteRow) Render() string {
	t := "t"
	if r.On {
		t += " is-strong"
	}
	var tags strings.Builder
	for _, tg := range r.Tags {
		tags.WriteString(tg.Render())
	}
	if r.Badge != nil {
		tags.WriteString(r.Badge.Render())
	}
	return `<div class="tcx-row tcx-row--note` + on(r.On) + `" role="listitem">` + r.Mark.Render() +
		`<span class="` + t + `">` + esc(r.Title) + `</span>` +
		`<span class="s tcx-hide-3">` + esc(r.Excerpt) + `</span>` +
		`<span class="d">` + tags.String() + `</span>` +
		`<span class="mono">` + esc(r.Date) + `</span>` +
		`<span class="d">` + esc(r.Source) + `</span></div>`
}

// FileRow is .m-row--file.
type FileRow struct {
	Icon   string
	Name   string
	Size   string
	Time   string
	Mode   string
	Source string
	Sel    bool
}

func (r FileRow) Render() string {
	c := "tcx-row tcx-row--file"
	if r.Sel {
		c += " is-sel"
	}
	return `<div class="` + c + `" role="listitem">` + Box{Off: !r.Sel, Label: "Select " + r.Name}.Render() +
		`<span class="t">` + Icon{Name: r.Icon}.Render() + `<span>` + esc(r.Name) + `</span></span>` +
		`<span class="s mono">` + esc(r.Size) + `</span>` +
		`<span class="s mono">` + esc(r.Time) + `</span>` +
		`<span class="d mono">` + esc(r.Mode) + `</span>` +
		`<span class="d">` + esc(r.Source) + `</span></div>`
}

// Box is .m-box: a tick.
type Box struct {
	Off   bool
	Part  bool
	Label string
}

func (b Box) Render() string {
	c := "tcx-box"
	switch {
	case b.Part:
		c += " is-part"
	case b.Off:
		c += " is-off"
	}
	return `<button type="button" class="` + c + `" aria-label="` + esc(b.Label) + `"></button>`
}

// Col is a table column; Kind gives its width.
type Col struct {
	Label string
	Kind  ColKind
	Tick  bool
}

// ColKind is a column width role.
type ColKind string

const (
	ColFlex    ColKind = ""
	ColTick    ColKind = "tcx-col-tick"
	ColSize    ColKind = "tcx-col-size"
	ColType    ColKind = "tcx-col-type"
	ColTime    ColKind = "tcx-col-time"
	ColActs    ColKind = "tcx-col-acts"
	ColTickSm  ColKind = "tcx-col-tick-sm"
	ColCount   ColKind = "tcx-col-count"
	ColChanged ColKind = "tcx-col-changed"
	ColMenu    ColKind = "tcx-col-menu"
)

// Table is .m-table.
type Table struct {
	Cols []Col
	Rows []Tr
}

// Tr is one body row.
type Tr struct {
	Folder bool
	Sel    bool
	Off    bool
	Cells  []Td
}

// Td is one cell. Mono and Acts are the cell treatments of the frame.
type Td struct {
	Body Nodes
	Span int
	Rows int
	Mono bool
	Acts bool
}

func (t Table) Render() string {
	var b strings.Builder
	b.WriteString(`<table class="tcx-table"><thead><tr>`)
	for _, c := range t.Cols {
		b.WriteString(`<th`)
		if c.Kind != "" {
			b.WriteString(` class="` + string(c.Kind) + `"`)
		}
		b.WriteString(`>`)
		if c.Tick {
			b.WriteString(Box{Off: true, Label: "Select all"}.Render())
		} else {
			b.WriteString(esc(c.Label))
		}
		b.WriteString(`</th>`)
	}
	b.WriteString(`</tr></thead><tbody>`)
	for _, r := range t.Rows {
		var c []string
		if r.Folder {
			c = append(c, "is-folder")
		}
		if r.Sel {
			c = append(c, "is-sel")
		}
		if r.Off {
			c = append(c, "is-off")
		}
		b.WriteString(`<tr`)
		if len(c) > 0 {
			b.WriteString(` class="` + strings.Join(c, " ") + `"`)
		}
		b.WriteString(`>`)
		for _, d := range r.Cells {
			b.WriteString(`<td`)
			switch {
			case d.Mono:
				b.WriteString(` class="mono"`)
			case d.Acts:
				b.WriteString(` class="tcx-acts"`)
			}
			if d.Span > 1 {
				b.WriteString(` colspan="` + strconv.Itoa(d.Span) + `"`)
			}
			if d.Rows > 1 {
				b.WriteString(` rowspan="` + strconv.Itoa(d.Rows) + `"`)
			}
			b.WriteString(`>` + d.Body.Render() + `</td>`)
		}
		b.WriteString(`</tr>`)
	}
	b.WriteString(`</tbody></table>`)
	return b.String()
}

// FName is .m-fname: type icon (or a collection mark) and name. Up is the
// "one folder up" link.
type FName struct {
	Icon string
	Mark *Mark
	Name string
	Up   bool
}

func (f FName) Render() string {
	name := `<span>` + esc(f.Name) + `</span>`
	if f.Up {
		name = `<span class="tcx-ref">` + esc(f.Name) + `</span>`
	}
	lead := Icon{Name: f.Icon}.Render()
	if f.Mark != nil {
		lead = f.Mark.Render()
	}
	return `<span class="tcx-fname">` + lead + name + `</span>`
}

// Msg is .m-msg. Flush drops its margin inside a sheet body.
type Msg struct {
	Text  string
	Body  Nodes
	Alert bool
	Flush bool
}

func (m Msg) Render() string {
	c := "tcx-msg"
	if m.Alert {
		c += " is-alert"
	}
	if m.Flush {
		c += " is-flush"
	}
	body := esc(m.Text)
	if m.Body != nil {
		body = m.Body.Render()
	}
	return `<div class="` + c + `">` + body + `</div>`
}

// Side is .m-side: the details panel.
type Side struct {
	Head Head
	Body Nodes
}

func (s Side) Render() string {
	return `<aside class="tcx-side" aria-label="Details">` + s.Head.Render() + s.Body.Render() + `</aside>`
}

// Form is .m-form.
type Form struct {
	Tight  bool
	Fields []Field
}

// Width is a field width, as .w-*.
type Width string

const (
	WXS   Width = "w-xs"
	WSM   Width = "w-sm"
	WMD   Width = "w-md"
	WFull Width = "w-full"
	W100  Width = "tcx-w100" // a field across a sheet body
)

// Field is .m-f: label on top, one control.
type Field struct {
	Label string
	Width Width
	In    In
}

// InKind is the control a field prints.
type InKind string

const (
	InText InKind = "text"
	InDate InKind = "date"
	InTime InKind = "time"
	InArea InKind = "area"
)

// In is .m-in as a real control. Tall is the note text box in a panel, Note
// the text box of the quick note sheet. Focus is the ring a frame draws.
type In struct {
	ID          string
	Kind        InKind
	Value       string
	Placeholder string
	Aria        string // name when there is no <label>
	Mono        bool
	Tall        bool
	Note        bool
	Focus       bool
}

func (f Form) Render() string {
	c := "tcx-form"
	if f.Tight {
		c += " is-tight"
	}
	var b strings.Builder
	b.WriteString(`<form class="` + c + `">`)
	for _, fl := range f.Fields {
		b.WriteString(fl.Render())
	}
	b.WriteString(`</form>`)
	return b.String()
}

func (i In) Render() string {
	c := "tcx-in"
	if i.Mono {
		c += " is-mono"
	}
	if i.Placeholder != "" {
		c += " is-ph"
	}
	if i.Focus {
		c += " is-focus"
	}
	attrs := ` id="` + i.ID + `"`
	if i.Aria != "" {
		attrs += ` aria-label="` + esc(i.Aria) + `"`
	}
	if i.Placeholder != "" {
		attrs += ` placeholder="` + esc(i.Placeholder) + `"`
	}
	if i.Kind == InArea {
		c += " is-area"
		if i.Tall {
			c += " is-tall"
		}
		if i.Note {
			c += " is-note"
		}
		return `<textarea` + attrs + ` class="` + c + `">` + esc(i.Value) + `</textarea>`
	}
	return `<input` + attrs + ` class="` + c + `" type="` + string(i.Kind) + `" value="` + esc(i.Value) + `">`
}

// Field is also usable alone, outside a Form.
func (f Field) Render() string {
	return `<div class="tcx-f ` + string(f.Width) + `"><label for="` + f.In.ID + `">` + esc(f.Label) + `</label>` + f.In.Render() + `</div>`
}

// Sec is .m-sec: a rubric, an optional action, a body.
type Sec struct {
	Rubric string
	Action *Btn
	Body   Nodes
}

func (s Sec) Render() string {
	act := ""
	if s.Action != nil {
		act = " " + s.Action.Render()
	}
	return `<section class="tcx-sec"><div class="tcx-rubric">` + esc(s.Rubric) + act + `</div>` + s.Body.Render() + `</section>`
}

// Attach is one attachment line.
type Attach struct {
	Name  string
	Meta  string
	Where string
}

// Attachments is the attachment list of a note panel.
type Attachments []Attach

func (a Attachments) Render() string {
	var b strings.Builder
	b.WriteString(`<ul class="tcx-attach">`)
	for _, it := range a {
		b.WriteString(`<li>` + esc(it.Name) + ` ` + Inline{Dim, it.Meta}.Render() + ` ` + Badge{Label: it.Where}.Render() + `</li>`)
	}
	b.WriteString(`</ul>`)
	return b.String()
}

// Hint is .m-hint; Spaced is the hint under a list.
type Hint struct {
	Text   string
	Spaced bool
}

func (h Hint) Render() string {
	c := "tcx-hint"
	if h.Spaced {
		c += " is-spaced"
	}
	return `<div class="` + c + `">` + esc(h.Text) + `</div>`
}

// Related is one linked record.
type Related struct {
	Title string
	Meta  string
}

// RelatedList is the related records of a panel.
type RelatedList []Related

func (r RelatedList) Render() string {
	var b strings.Builder
	b.WriteString(`<ul class="tcx-related">`)
	for _, it := range r {
		b.WriteString(`<li>` + Inline{Ref, it.Title}.Render() + ` ` + Inline{Dim, it.Meta}.Render() + `</li>`)
	}
	b.WriteString(`</ul>`)
	return b.String()
}

// FieldRow is one dt/dd pair. Mono sets the whole value in mono, Rich holds
// inline parts; Path is the long path value.
type FieldRow struct {
	Label string
	Value string
	Rich  Nodes
	Mono  bool
	Path  bool
}

// Fields is .m-fields.
type Fields []FieldRow

func (f Fields) Render() string {
	var b strings.Builder
	b.WriteString(`<dl class="tcx-fields">`)
	for _, r := range f {
		var c []string
		if r.Mono {
			c = append(c, "mono")
		}
		if r.Path {
			c = append(c, "tcx-path")
		}
		dd := `<dd>`
		if len(c) > 0 {
			dd = `<dd class="` + strings.Join(c, " ") + `">`
		}
		val := esc(r.Value)
		if r.Rich != nil {
			val = r.Rich.Render()
		}
		b.WriteString(`<dt>` + esc(r.Label) + `</dt>` + dd + val + `</dd>`)
	}
	b.WriteString(`</dl>`)
	return b.String()
}

// Preview is the file preview of a panel.
type Preview struct{ Icon string }

func (p Preview) Render() string {
	return `<div class="tcx-preview"><div class="tcx-thumb is-wide">` + Icon{Name: p.Icon, Box: true, Caps: true}.Render() + `</div></div>`
}

// PanelActions is the row of actions under a preview.
type PanelActions Nodes

func (p PanelActions) Render() string {
	return `<div class="tcx-side-acts">` + Nodes(p).Render() + `</div>`
}

// Note is a short explanatory paragraph in a panel section.
type Note string

func (n Note) Render() string { return `<div class="tcx-note">` + esc(string(n)) + `</div>` }

// DangerZone is .m-danger-zone.
type DangerZone Nodes

func (d DangerZone) Render() string {
	return `<div class="tcx-danger-zone">` + Nodes(d).Render() + `</div>`
}

func on(v bool) string {
	if v {
		return " is-on"
	}
	return ""
}

func onAttr(v bool) string {
	if v {
		return ` class="is-on"`
	}
	return ""
}

func esc(s string) string {
	return strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", `"`, "&quot;").Replace(s)
}
