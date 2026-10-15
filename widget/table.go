package widget

import (
	"strconv"
	"strings"
)

// ColWidth is the drawn width of a table's column.
type ColWidth int

const (
	ColAuto      ColWidth = iota
	ColTick               // a tick box
	ColTickSmall          // a small tick box
	ColSize               // a size
	ColType               // a type
	ColTime               // a date and time
	ColNum                // a count
	ColMenu               // a «⋯»
	ColChanged            // when it changed
	Col26                 // 26 %
	Col80                 // 80 px
	Col90                 // 90 px
	Col110                // 110 px
	Col120                // 120 px
)

func (w ColWidth) class() string {
	return [...]string{"", "tc-table__col--tick", "tc-table__col--tick-sm", "tc-table__col--size", "tc-table__col--type", "tc-table__col--time",
		"tc-table__col--num", "tc-table__col--menu", "tc-table__col--changed", "tc-table__col--w26", "tc-table__col--w80", "tc-table__col--w90",
		"tc-table__col--w110", "tc-table__col--w120"}[w]
}

// ColRank is which of the columns that narrow tables drop a column is: the
// third or the fourth. The rank goes on the column's head and every cell.
// Rank1 is the column of a packed table (TablePacked) that leaves first,
// where it does not fit.
type ColRank int

const (
	RankKept ColRank = iota
	Rank3
	Rank4
	Rank1
)

func (r ColRank) class() string {
	return [...]string{"", "tc-table__n3", "tc-table__n4", "tc-table__n1"}[r]
}

// TableCol is a column of a table: its head and what every cell of it
// shares.
type TableCol struct {
	Label     string
	Width     ColWidth
	Rank      ColRank
	Secondary bool    // the first to go when the table narrows
	Tick      *Ticked // a tick box in the head: all rows at once
	Send      Send    // what the head's tick sends
	// Main is the column that matters (the owner's decision of 15.10,
	// Р-66): on a phone a row is a card of its name, this column under it,
	// named by its Label, and «⋯»; the other columns go and the row's
	// buttons are in «⋯». Which column matters is the table's sense: the
	// frame names it.
	Main bool
}

// CellTone is the colour of a cell's words.
type CellTone int

const (
	CellInk CellTone = iota
	CellDusty
	CellDanger
	CellMuted
)

// CellPart is what a table's cell holds.
type CellPart interface {
	part
	cellPart()
}

func (Text) cellPart()      {}
func (Break) cellPart()     {}
func (TableFile) cellPart() {}
func (Line) cellPart()      {}
func (Dim) cellPart()       {}
func (Button) cellPart()    {}
func (Badge) cellPart()     {}
func (Tag) cellPart()       {}
func (Tick) cellPart()      {}
func (Drop) cellPart()      {}
func (CellWrap) cellPart()  {}
func (Chevron) cellPart()   {}
func (Mono) cellPart()      {}
func (Seg) cellPart()       {}
func (Span) cellPart()      {}

// Break ends a line inside a cell: a second line under the first.
type Break struct{}

func (Break) html(mix) string { return `<br>` }

// Chevron is the small arrow after a cell's value that can be changed.
type Chevron struct{}

func (Chevron) html(mix) string {
	return `<svg width="11" height="11" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.6" fill="none" class="tc-ico--inline">` + Icon("chev").use() + `</svg>`
}

// TableFile is the name cell of a table: an icon or a source's mark, then
// the name.
type TableFile struct {
	Icon   Icon
	Mark   *Mark
	Name   string
	Accent bool // a link to go up: «..»
	// Href makes the name a link: a folder opens, a file is chosen (Carrel's
	// task 114). The row cannot be one, its tick and buttons are controls of
	// their own.
	Href string
}

func (f TableFile) html(mix) string {
	s := `<span class="tc-table__file">`
	if f.Mark != nil {
		s += f.Mark.html(mix{"tc-table__filename"})
	} else {
		s += `<svg>` + f.Icon.use() + `</svg>`
	}
	name := "span"
	if f.Href != "" {
		name = "a"
	}
	return s + `<` + name + class(nil, "tc-table__filename", pick(f.Accent, "tc-table__filename--accent")) + attr("href", f.Href) + `>` + esc(f.Name) + `</` + name + `></span>`
}

// CellWrap lets the pieces of a cell wrap: a device's rights.
type CellWrap struct {
	Items []CellPart
}

func (w CellWrap) html(mix) string {
	var s strings.Builder
	s.WriteString(`<span class="tc-table__wrap">`)
	for _, it := range w.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</span>`)
	return s.String()
}

// TableCell is a cell of a table.
type TableCell struct {
	Parts []CellPart
	Mono  bool
	Dim   bool // the whole cell dim
	Tone  CellTone
	Small bool    // 11 px
	Pick  *Ticked // a tick box that picks the row
	Send  Send    // what the pick sends: the row's name in a bulk action's form
	Acts  bool    // the row's buttons
	More  *Button // the row's «⋯»: shown only when the row's is-2nd buttons have no room
	// Drop is the row's «⋯» with its menu, in place of More: the twins of
	// the row's is-2nd buttons; it opens without a script (Carrel's task 117).
	Drop  *Drop
	Unset bool     // a placeholder value
	Name  bool     // the name column of a narrow list
	Span  int      // the cell spans that many columns
	Rows  int      // the cell spans that many rows: one value for several
	Rank  ColRank  // a table without a head: the cell carries its column's rank
	Width ColWidth // a table without a head: the cell carries its column's width
}

// RowState is the state of a table's row.
type RowState int

const (
	RowPlain  RowState = iota
	RowPicked          // picked by its tick box
	RowFolder          // a folder among files
	RowOff             // switched off
	RowOn              // the current row
	RowFaded           // a job that does not run
	RowGone            // a row that no longer applies
)

func (s RowState) class() string {
	return [...]string{"", "is-sel", "is-folder", "is-off", "is-on", "tc-table__row--dim", "tc-table__row--off"}[s]
}

// TableRow is a row of a table.
type TableRow struct {
	Cells []TableCell
	State RowState
}

// TableVariant is the drawn form of a table.
type TableVariant int

const (
	TablePlain TableVariant = iota
	TableFit                // its content's width
	TableWide               // wider than its page: scrolls
	TableInset              // out to a card's edges
	// TablePacked is a table of many narrow columns whose words are never
	// broken (the owner's decisions of 09.10): a column that does not fit
	// leaves — Rank1 first, Rank3 after it — and does not come back as its
	// panel narrows.
	TablePacked
)

func (v TableVariant) class() string {
	return [...]string{"", "tc-table--fit", "tc-table--wide", "tc-table--mx-20", "tc-table--packed"}[v]
}

// Table — tc-table
//
// Роль. Таблица: файлы папки, строки мастера, задания копирования, устройства.
// Вход. Table{Cols, Rows, Variant, Place, Target}.
// Состояния. plain | head | picked | folder | off | acts | wide | fit.
// Инварианты. Ширина и ранг колонки — перечисления, и ячейки колонки берут их из её заголовка; третья и четвёртая колонки уходят первыми на узком экране; таблица печатает фрагмент своего кадра.
// Поведение. Галка выбирает строку для действий бара; кнопки строки с `is-2nd` на узкой ширине уходят под «⋯» строки (`TableCell.More`), и «⋯» открывает её меню; имя со ссылкой (`TableFile.Href`) открывает строку. На телефоне таблица с главной колонкой (`TableCol.Main`) — карточки: имя и «⋯», под ними главное с подписью колонки (Р-66).
// Эталон. кадр `s-files-1`, фрагмент `.tc-table`.
// Содержимое. Заголовки колонок, строки: имя со значком или меткой, значения, моноширинное, две строки, метки, теги, бейджи, галка, кнопки, выпадашка.
// Чего нет. Сортировки по заголовку: сортирует переключатель бара.
type Table struct {
	Cols    []TableCol // none: a table without a head
	Rows    []TableRow
	Variant TableVariant
	Place   Place
	// Target is the name of the zone the links and forms inside ask their
	// address for with the live layer, instead of loading the page (Carrel's
	// task 84); without the layer they are plain links and forms.
	Target string
}

// Render prints the table.
func (t Table) Render() string { return t.html(nil) }

func (t Table) col(i int, c TableCell) TableCol {
	if i < len(t.Cols) {
		return t.Cols[i]
	}
	return TableCol{Rank: c.Rank, Width: c.Width}
}

func (t Table) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<table` + class(m, "tc-table", t.Place.class(), t.Variant.class()) + attr("data-tc-target", t.Target) + `>`)
	if len(t.Cols) > 0 {
		s.WriteString(`<thead><tr>`)
		for _, c := range t.Cols {
			s.WriteString(`<th` + class(nil, c.Width.class(), c.Rank.class(), rank(c.Secondary), pick(c.Main, "tc-table__main")) + `>`)
			if c.Tick != nil {
				s.WriteString(c.Tick.box("All", c.Send))
			}
			s.WriteString(esc(c.Label) + `</th>`)
		}
		s.WriteString(`</tr></thead>`)
	}
	s.WriteString(`<tbody>`)
	for _, r := range t.Rows {
		s.WriteString(`<tr` + class(nil, r.State.class()) + `>`)
		for i, c := range r.Cells {
			col := t.col(i, c)
			s.WriteString(`<td` + class(nil, c.classes(col)...) + attr("data-tc-label", pick(col.Main, col.Label)) + attr("colspan", pick(c.Span > 1, strconv.Itoa(c.Span))) + attr("rowspan", pick(c.Rows > 1, strconv.Itoa(c.Rows))) + `>`)
			if c.Pick != nil {
				s.WriteString(c.Pick.box("Pick", c.Send))
			}
			for _, p := range c.Parts {
				s.WriteString(p.html(nil))
			}
			if c.More != nil {
				s.WriteString(c.More.html(mix{"tc-acts__more"}))
			}
			if c.Drop != nil {
				d := *c.Drop
				d.Row = true
				s.WriteString(d.html(nil))
			}
			s.WriteString(`</td>`)
		}
		s.WriteString(`</tr>`)
	}
	s.WriteString(`</tbody></table>`)
	return s.String()
}

func (c TableCell) classes(col TableCol) []string {
	out := []string{pick(c.Mono, "tc-mono"), pick(c.Dim, "tc-dim"), pick(c.Acts, "tc-acts"), pick(c.Name, "tc-table__name"),
		pick(c.Unset, "is-ph"), col.Rank.class(), rank(col.Secondary), pick(col.Main, "tc-table__main")}
	if c.Width != ColAuto {
		out = append(out, c.Width.class())
	}
	r := col.Rank.class()
	switch {
	case r != "" && c.Tone == CellDusty:
		out = append(out, r+"--dusty")
	case r != "" && c.Tone == CellDanger:
		out = append(out, r+"--danger")
	case c.Tone == CellDusty:
		out = append(out, "tc-mono--dusty")
	case c.Tone == CellDanger && c.Mono:
		out = append(out, "tc-mono--danger")
	case c.Tone == CellDanger:
		out = append(out, "tc-dim--danger")
	case c.Tone == CellMuted:
		out = append(out, "tc-dim--muted")
	}
	if r != "" && c.Small {
		out = append(out, r+"--t11")
	}
	for _, p := range c.Parts {
		if _, ok := p.(Drop); ok && r != "" {
			out = append(out, r+"--pos-relative")
		}
	}
	return out
}
