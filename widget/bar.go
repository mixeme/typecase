package widget

import "strings"

// BarPart is what a bar holds.
type BarPart interface {
	part
	barPart()
}

func (Button) barPart()    {}
func (Check) barPart()     {}
func (CheckForm) barPart() {}
func (Input) barPart()     {}
func (Filter) barPart()    {}
func (Pick) barPart()      {}
func (Mono) barPart()      {}
func (Range) barPart()     {}
func (Seg) barPart()       {}
func (SegFold) barPart()   {}
func (BarSep) barPart()    {}
func (Drop) barPart()      {}

// BarSep is a thin rule between groups of a bar.
type BarSep struct {
	Secondary bool
	Tertiary  bool
}

func (p BarSep) html(mix) string {
	return `<span` + class(nil, "tc-bar__sep", rank(p.Secondary), rank3(p.Tertiary)) + `></span>`
}

// BarCount says how many rows are selected, at the start of a selection bar.
type BarCount struct {
	Text  string
	Small bool // 12.5 px: a phone
}

// BarRight is the end of a bar: a count of what is shown, a way out of the
// selection.
type BarRight struct {
	Items     []BarPart
	Secondary bool
}

// BarVariant is the drawn form of a bar.
type BarVariant int

const (
	BarPlain     BarVariant = iota
	BarSelection            // what to do with the selected rows
	BarWide                 // a range and switches: more room between groups
)

func (v BarVariant) class() string {
	return [...]string{"", "tc-bar--sel", "tc-bar--wide"}[v]
}

// Bar — tc-bar
//
// Роль. Бар над списком: фильтр, сортировка, вид, «⋯»; при выделении — что сделать с выделенным.
// Вход. Bar{Count, Items, More, Right, Variant}.
// Состояния. plain | selection | wide | more | right.
// Инварианты. Бар — одна строка: второстепенное (is-2nd) уходит под «⋯» первым, третьестепенное (is-3rd) — на телефоне, порядок и пороги — в CSS; «⋯» — одна кнопка перед концом бара; бар печатает фрагмент своего кадра.
// Поведение. «⋯» открывает меню с тем, что не поместилось.
// Эталон. кадр `components-16`, фрагмент `.tc-bar`.
// Содержимое. Счёт выделенного, фильтр, поле, переключатели, диапазон, флажки, кнопки, выпадашка «⋯» с меню, разделители, моноширинное, конец бара.
// Чего нет. Второй строки: то, что не влезло, уходит под «⋯», а не переносится.
type Bar struct {
	Count   *BarCount
	Items   []BarPart
	More    *Button
	Right   *BarRight
	Variant BarVariant
}

// Render prints the bar.
func (b Bar) Render() string { return b.html(nil) }

func (b Bar) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-bar", b.Variant.class()) + `>`)
	if b.Count != nil {
		s.WriteString(`<span` + class(nil, "tc-bar__count", pick(b.Count.Small, "tc-bar__count--t12-5")) + `>` + esc(b.Count.Text) + `</span>`)
	}
	for _, it := range b.Items {
		s.WriteString(it.html(nil))
	}
	if b.More != nil {
		s.WriteString(b.More.html(mix{"tc-bar__more"}))
	}
	if b.Right != nil {
		s.WriteString(`<span` + class(nil, "tc-bar__right", rank(b.Right.Secondary)) + `>`)
		for _, it := range b.Right.Items {
			s.WriteString(it.html(nil))
		}
		s.WriteString(`</span>`)
	}
	s.WriteString(`</div>`)
	return s.String()
}
