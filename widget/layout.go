package widget

import "strings"

// LineGap is the gap between the items of a line.
type LineGap int

const (
	LineGap0 LineGap = iota
	LineGap3
	LineGap4
	LineGap6
	LineGap7
	LineGap8
	LineGap9
	LineGap10
	LineGap12
)

func (g LineGap) class() string {
	return [...]string{"", "tc-line--gap-3", "tc-line--gap-4", "tc-line--gap-6", "tc-line--gap-7", "tc-line--gap-8",
		"tc-line--gap-9", "tc-line--gap-10", "tc-line--gap-12"}[g]
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
// Вход. Line{Items, Gap, Block, Place}.
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
}

// Render prints the line.
func (l Line) Render() string { return l.html(nil) }

func (l Line) html(m mix) string {
	tag := "span"
	if l.Block {
		tag = "div"
	}
	var s strings.Builder
	s.WriteString(`<` + tag + class(m, "tc-line", l.Gap.class(), l.Place.class()) + `>`)
	for _, p := range l.Items {
		s.WriteString(p.html(nil))
	}
	s.WriteString(`</` + tag + `>`)
	return s.String()
}

// StackGap is the gap between the items of a column.
type StackGap int

const (
	StackGap9 StackGap = iota
	StackGap10
)

func (g StackGap) class() string {
	return [...]string{"tc-stack--gap-9", "tc-stack--gap-10"}[g]
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
