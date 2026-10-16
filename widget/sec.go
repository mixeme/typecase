package widget

import "strings"

// ValuePart is what a value of a panel's fields holds.
type ValuePart interface {
	part
	valuePart()
}

func (Text) valuePart()  {}
func (Mono) valuePart()  {}
func (Span) valuePart()  {}
func (Dim) valuePart()   {}
func (Badge) valuePart() {}
func (Tag) valuePart()   {}
func (Break) valuePart() {}

// A mark in a value names the collection the value came from: a card of a
// linked record shows the fields of all its copies, and a value that is not
// in the copy opened carries the mark and the name of its own (Р-58).
func (Mark) valuePart() {}

// ValueFace is how a value of a panel's fields is set.
type ValueFace int

const (
	ValuePlain      ValueFace = iota
	ValueMono                 // monospaced: a phone, a date
	ValueMonoSmall            // monospaced, 11 px: a long identifier
	ValueMonoAccent           // monospaced in the accent: a phone to call
	ValueAccent               // in the accent: an address to write to
)

func (f ValueFace) class() string {
	return [...]string{"", "tc-mono", "tc-mono tc-mono--t11", "tc-mono tc-mono--accent", "tc-dim--accent"}[f]
}

// FieldsItem is a name and its value in a panel.
type FieldsItem struct {
	Term  string
	Value []ValuePart
	Face  ValueFace
	Href  string // the value goes somewhere: mailto:, tel:, the linked record
	// After stands after a value that is a link, outside the link: the mark
	// and the name of the collection a value came from (Р-58).
	After []ValuePart
}

// FieldsVariant is the drawn form of a panel's fields.
type FieldsVariant int

const (
	FieldsPlain    FieldsVariant = iota
	FieldsBleed                  // out to the card's edges
	FieldsBleed130               // out to the edges, names 130 px
	Fields130Flush               // names 130 px, flush, 14 px above
	Fields88                     // names 88 px: a phone
	Fields88Pad                  // names 88 px, 4/14/12 padding: a phone's card
)

func (v FieldsVariant) class() string {
	return [...]string{"", "tc-fields--bleed", "tc-fields--bleed tc-fields--cols130", "tc-fields--cols130 tc-fields--flush-padt14", "tc-fields--cols88",
		"tc-fields--cols88 tc-fields--pad4-14-12"}[v]
}

// Fields — tc-fields
//
// Роль. Поля панели для чтения: имя слева, значение справа — телефон, когда, повторы, теги.
// Вход. Fields{Items, Variant, Place}.
// Состояния. plain | bleed | narrow | phone.
// Инварианты. Имя и значение — пара, одно под другим не переносится; вид значения — перечисление; поля печатают фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `s-contacts-6`, фрагмент `.tc-fields`.
// Содержимое. Имена, значения: текст, моноширинное, приглушённое, бейдж, тег, метка, перенос строки.
// Чего нет. Правки на месте: правят в форме.
type Fields struct {
	Items   []FieldsItem
	Variant FieldsVariant
	Place   Place
}

// Render prints the fields.
func (f Fields) Render() string { return f.html(nil) }

func (f Fields) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<dl` + class(m, "tc-fields", f.Place.class(), f.Variant.class()) + `>`)
	for _, it := range f.Items {
		s.WriteString(`<dt>` + esc(it.Term) + `</dt><dd` + class(nil, it.Face.class()) + `>`)
		if it.Href != "" {
			s.WriteString(`<a class="tc-fields__link"` + attr("href", it.Href) + `>`)
		}
		for _, p := range it.Value {
			s.WriteString(p.html(nil))
		}
		if it.Href != "" {
			s.WriteString(`</a>`)
		}
		for _, p := range it.After {
			s.WriteString(" " + p.html(nil))
		}
		s.WriteString(`</dd>`)
	}
	s.WriteString(`</dl>`)
	return s.String()
}

// SecItemPart is what a line of a section's list holds.
type SecItemPart interface {
	part
	secItemPart()
}

func (Text) secItemPart()   {}
func (Strong) secItemPart() {}
func (Mono) secItemPart()   {}
func (Span) secItemPart()   {}
func (Link) secItemPart()   {}
func (Dim) secItemPart()    {}
func (Mark) secItemPart()   {}
func (Badge) secItemPart()  {}
func (Tag) secItemPart()    {}
func (Button) secItemPart() {}
func (Break) secItemPart()  {}

// SecList is a list in a section: related records, what the letter carries.
type SecList struct {
	Items [][]SecItemPart
	Gap   int  // 5 or 6 px between the items; 0: no column
	Muted bool // quieter: what a letter carries
	// Columns sets the items in columns across the width: the attendees of
	// an event (Carrel's task 77).
	Columns bool
}

func (l SecList) html(mix) string {
	var s strings.Builder
	col := ""
	if l.Gap > 0 {
		col = "tc-stack tc-stack--gap-" + itoa(l.Gap)
	}
	s.WriteString(`<ul` + class(nil, col, "tc-sec__list", pick(l.Muted, "tc-sec__list--muted"), pick(l.Columns, "tc-sec__list--cols")) + `>`)
	for _, it := range l.Items {
		s.WriteString(`<li>` + joined(it) + `</li>`)
	}
	s.WriteString(`</ul>`)
	return s.String()
}

// SecHint is a line of explanation in a section.
type SecHint struct {
	Text  string
	Apart bool // 6 px above: after a list
}

func (h SecHint) html(mix) string {
	return `<div` + class(nil, "tc-field__hint", pick(h.Apart, "tc-field__hint--top6")) + `>` + esc(h.Text) + `</div>`
}

// SecNote is what a section says when it cannot list anything.
type SecNote struct {
	Text string
}

func (n SecNote) html(mix) string { return `<div class="tc-sec__note">` + esc(n.Text) + `</div>` }

// SecActs is the buttons of a section.
type SecActs struct {
	Items []Button
}

func (a SecActs) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-sec__acts">`)
	for _, b := range a.Items {
		s.WriteString(b.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// SecTasks is a note's tasks in a section of the side panel.
type SecTasks struct {
	Items []DocTask
}

func (t SecTasks) html(mix) string {
	var s strings.Builder
	s.WriteString(`<ul class="tc-stack tc-stack--gap-7 tc-doc__list">`)
	for _, it := range t.Items {
		s.WriteString(`<li class="tc-doc__task">` + it.Tick.html(nil))
		if it.Tick.Done {
			// a done task is struck, as on the note's page
			s.WriteString(Dim{Parts: []Inline{Text(it.Text)}, Variant: DimStruck}.html(nil))
		} else {
			s.WriteString(esc(it.Text))
		}
		s.WriteString(`</li>`)
	}
	s.WriteString(`</ul>`)
	return s.String()
}

// SecFile is a thing attached to a letter: its icon and name, its kind and
// size, what it is, the button that takes it off.
type SecFile struct {
	Icon Icon
	Name string
	Size Mono
	Tag  Tag
	// Terms and Change are for a file that goes as a link (the owner's
	// decision of 12.10, Б38): the link's terms in words, as the list of
	// links says them («30 days»), and the button that opens the publish
	// sheet over the letter.
	Terms  *Mono
	Change *Button
	Remove Button
}

// SecFiles is what a letter carries.
type SecFiles struct {
	Items []SecFile
}

func (f SecFiles) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-stack tc-stack--gap-6 tc-card__top8">`)
	for _, it := range f.Items {
		s.WriteString(`<div class="tc-line tc-line--gap-9 tc-line__text"><span class="tc-table__file tc-table__file--fill"><svg>` + it.Icon.use() +
			`</svg><span class="tc-table__filename">` + esc(it.Name) + `</span></span>` + it.Size.html(nil) + it.Tag.html(nil))
		if it.Terms != nil {
			s.WriteString(it.Terms.html(nil))
		}
		if it.Change != nil {
			s.WriteString(it.Change.html(nil))
		}
		s.WriteString(it.Remove.html(nil) + `</div>`)
	}
	s.WriteString(`</div>`)
	return s.String()
}

// SecOpenPart is what the lines of an opened file hold.
type SecOpenPart interface {
	part
	secOpenPart()
}

func (Line) secOpenPart()      {}
func (Dim) secOpenPart()       {}
func (SecOpened) secOpenPart() {}

// SecOpened is the name of a file opened in a program, with its size.
type SecOpened struct {
	Name string
	Size Dim
}

func (o SecOpened) html(mix) string {
	return `<span class="tc-stack__item">` + esc(o.Name) + ` ` + o.Size.html(nil) + `</span>`
}

// SecOpen is a file opened in a program, its buttons and what happens to it.
type SecOpen struct {
	Items []SecOpenPart
}

func (o SecOpen) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-stack tc-stack--gap-7 tc-line__text">`)
	for _, it := range o.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// SecPart is what a section holds.
type SecPart interface {
	part
	secPart()
}

func (Rubric) secPart()   {}
func (List) secPart()     {}
func (Prose) secPart()    {}
func (Msg) secPart()      {}
func (Line) secPart()     {}
func (SecList) secPart()  {}
func (SecHint) secPart()  {}
func (SecNote) secPart()  {}
func (SecActs) secPart()  {}
func (SecTasks) secPart() {}
func (SecFiles) secPart() {}
func (SecOpen) secPart()  {}
func (Empty) secPart()    {}

// A note's text in its panel has the parts it has on its page (Carrel's task
// 125): headings, lists, quotes, code, tables.
func (DocHeading) secPart() {}
func (DocBullets) secPart() {}
func (DocQuote) secPart()   {}
func (DocCode) secPart()    {}
func (DocTable) secPart()   {}

// Sec — tc-sec
//
// Роль. Секция панели или листа: рубрика и то, что под ней, — поля, связанное, дела, вложения.
// Вход. Sec{Items, Flush}.
// Состояния. plain | list | hint | acts | flush.
// Инварианты. Рубрика — первой; секция выглядит как карточка без рамки; секция печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `s-contacts-6`, фрагмент `.tc-sec`.
// Содержимое. Рубрика, список строк, поля, абзац, таблица, дерево, сообщение, строки связанного, подсказка, примечание, кнопки, дела, вложения письма, открытый файл.
// Чего нет. Сворачивания.
type Sec struct {
	Items []SecPart
	Flush bool // no side padding, 11 px above: in a sheet
}

// Render prints the section.
func (c Sec) Render() string { return c.html(nil) }

func (c Sec) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-sec", pick(c.Flush, "tc-sec--flush-flush-padt11")) + `>`)
	for _, it := range c.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}
