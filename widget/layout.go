package widget

import "strings"

// LineGap is the gap between the items of a line: a step, not a number (the
// owner, 09.11). The steps are the root's variables --tc-gap-tight (4 px),
// --tc-gap (8), --tc-gap-wide (12) and --tc-gap-loose (16: fields side by
// side as in a form, a paragraph after a section); a container may set
// them otherwise.
type LineGap int

const (
	LineGap0      LineGap = iota
	LineGapTight          // marks, tags
	LineGapNormal         // buttons, fields, words
	LineGapWide           // checks, a logo with its badge
	LineGapFields         // fields side by side
)

func (g LineGap) class() string {
	return [...]string{"", "tc-line--tight", "tc-line--normal", "tc-line--wide", "tc-line--fields"}[g]
}

// LinePart is what a line holds.
type LinePart interface {
	part
	linePart()
}

func (Button) linePart() {}
func (Input) linePart()  {}
func (Pick) linePart()   {}
func (Kbd) linePart()    {}
func (Dim) linePart()    {}
func (Check) linePart()  {}
func (Mark) linePart()   {}
func (Mono) linePart()   {}
func (Tag) linePart()    {}
func (Badge) linePart()  {}
func (Field) linePart()  {}
func (Text) linePart()   {}
func (Strong) linePart() {}
func (Faced) linePart()  {}
func (Push) linePart()   {}
func (Seg) linePart()    {}
func (Line) linePart()   {}
func (Drop) linePart()   {}

// Faced is a piece of a line in a face of its own that holds blocks: a file
// name with its type and a badge after it.
type Faced struct {
	Face  Face
	Items []LinePart
}

func (f Faced) html(mix) string {
	var s strings.Builder
	s.WriteString(`<span class="` + f.Face.class() + `">`)
	for _, p := range f.Items {
		s.WriteString(p.html(nil))
	}
	s.WriteString(`</span>`)
	return s.String()
}

// Line — tc-line
//
// Роль. Строка: несколько вещей в ряд с одним зазором — кнопки и клавиша, поле и кнопка, метки, теги.
// Вход. Line{Items, Gap, Block, Place, ID}.
// Состояния. span | block | gap.
// Инварианты. Зазор — перечисление, а не число; строка не переносит чужие отступы; строка печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `s-contacts-3`, фрагмент `.tc-line`.
// Содержимое. Кнопки, кнопки с меню (Drop), поля, клавиши, метки, теги, бейджи, приглушённое, моноширинное, поля формы, жирное и простой текст.
// Чего нет. Своего фона и рамки: это раскладка, а не блок с видом.
type Line struct {
	Items []LinePart
	Gap   LineGap
	Block bool // a line of its own (div), not a piece of a row (span)
	Place Place
	ID    string // the list a picked file joins on an edit (FileField.List)
}

// Render prints the line.
func (l Line) Render() string { return l.html(nil) }

func (l Line) html(m mix) string {
	tag := "span"
	if l.Block {
		tag = "div"
	}
	var s strings.Builder
	s.WriteString(`<` + tag + class(m, "tc-line", l.Gap.class(), l.Place.class()) + attr("id", l.ID) + `>`)
	for _, p := range l.Items {
		s.WriteString(p.html(nil))
	}
	s.WriteString(`</` + tag + `>`)
	return s.String()
}

// StackGap is the gap between the items of a column: the same steps as a
// line's (the owner, 09.11).
type StackGap int

const (
	StackGap0      StackGap = iota
	StackGapTight           // the items of a list, checks
	StackGapNormal          // tasks, side blocks
	StackGapWide            // cards and sections one under another
)

func (g StackGap) class() string {
	return [...]string{"", "tc-stack--tight", "tc-stack--normal", "tc-stack--wide"}[g]
}

// StackPart is what a column holds.
type StackPart interface {
	part
	stackPart()
}

func (Button) stackPart() {}

// Stack — tc-stack
//
// Роль. Столбец: несколько вещей друг под другом с одним зазором.
// Вход. Stack{Items, Gap}.
// Состояния. default.
// Инварианты. Зазор — перечисление; столбец печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `components-21`, фрагмент `.tc-stack`.
// Содержимое. Кнопки, карточки.
// Чего нет. Своего фона и рамки.
type Stack struct {
	Items []StackPart
	Gap   StackGap
}

// Render prints the column.
func (k Stack) Render() string { return k.html(nil) }

func (k Stack) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-stack", k.Gap.class()) + `>`)
	for _, p := range k.Items {
		s.WriteString(p.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// Grow — tc-grow
//
// Роль. Перенос строки в форме: всё, что после него, начинается с новой строки.
// Вход. Grow{}.
// Состояния. нет.
// Инварианты. Пуст и невидим; печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `s-contacts-3`, фрагмент `.tc-grow`.
// Содержимое. нет.
// Чего нет. Высоты: перенос не отступ.
type Grow struct{}

// Render prints the break.
func (g Grow) Render() string { return g.html(nil) }

func (g Grow) html(m mix) string { return `<div` + class(m, "tc-grow") + `></div>` }

// PushPart is what a pushed group holds.
type PushPart interface {
	part
	pushPart()
}

func (Button) pushPart() {}
func (Text) pushPart()   {}

// Push — tc-push
//
// Роль. Группа, прижатая к концу ряда: кнопка остановки у очереди, «Default» в строке меню.
// Вход. Push{Items, Face}.
// Состояния. default.
// Инварианты. Прижата к концу, сколько бы ни было перед ней; печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `components-22`, фрагмент `.tc-push`.
// Содержимое. Кнопка или текст.
// Чего нет. Своего вида: только место.
type Push struct {
	Items []PushPart
	Face  *Face // pushed words in a face of their own
}

// Render prints the pushed group.
func (p Push) Render() string { return p.html(nil) }

func (p Push) html(m mix) string {
	var s strings.Builder
	face := ""
	if p.Face != nil {
		face = p.Face.class()
	}
	s.WriteString(`<span` + class(m, "tc-push", face) + `>`)
	for _, it := range p.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</span>`)
	return s.String()
}
