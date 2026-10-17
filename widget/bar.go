package widget

import "strings"

// BarPart is what a bar holds.
type BarPart interface {
	part
	barPart()
}

func (Button) barPart()    {}
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
func (BarForm) barPart()   {}

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
// Вход. Bar{Count, Items, More, Right, Variant, NoScript, Target}.
// Состояния. plain | selection | wide | more | right.
// Инварианты. Бар — одна строка, пока помещается: второстепенное (is-2nd) уходит под «⋯» первым, третьестепенное (is-3rd) — на телефоне, порядок и пороги — в CSS, а что не поместилось и после этого — переносится на следующую строку, краем не режется (Р-65); «⋯» — одна кнопка перед концом бара; бар печатает фрагмент своего кадра.
// Поведение. «⋯» открывает меню с тем, что не поместилось.
// Эталон. кадр `components-16`, фрагмент `.tc-bar`.
// Содержимое. Счёт выделенного, фильтр, поле, переключатели, диапазон, флажки, кнопки, выпадашка «⋯» с меню, разделители, моноширинное, конец бара; «⋯» с меню — Drop{More: true} среди Items: меню держит двойники того, что ушло под «⋯» (отметки — CheckForm, Carrel's task 82), а More — голая кнопка без меню.
// Чего нет. Второй строки: то, что не влезло, уходит под «⋯», а не переносится.
type Bar struct {
	Count   *BarCount
	Items   []BarPart
	More    *Button
	Right   *BarRight
	Variant BarVariant
	// NoScript is a bar only a page without a script shows: the selection
	// bar while nothing is ticked — with a script it comes with the first
	// tick, without one it stands always (the owner's decision of 09.10,
	// Carrel's task 148).
	NoScript bool
	// Target is the name of the zone the links and forms inside ask their
	// address for with the live layer, instead of loading the page (Carrel's
	// task 84); without the layer they are plain links and forms.
	Target string
}

// Render prints the bar.
func (b Bar) Render() string { return b.html(nil) }

func (b Bar) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-bar", b.Variant.class(), pick(b.NoScript, "tc-nojs")) + attr("data-tc-target", b.Target) + `>`)
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

// BarForm — tc-barform
//
// Роль. Фильтр полосы, который сам — форма: действие и день журнала уходят GET-запросом вместе (задача Carrel 195).
// Вход. BarForm{Items, Action, Hidden, Show}.
// Состояния. default.
// Инварианты. Поля внутри — поля формы (Name); форма отправляет их и скрытые поля; со сценарием — при смене, без сценария — кнопкой Show; виджет печатает фрагмент своего кадра.
// Поведение. Смена поля или выбор пункта меню отправляет форму (сценарий); без сценария — Show в конце.
// Эталон. кадр `s-admin-7`, фрагмент `.tc-barform`.
// Содержимое. Поля полосы: выбор с меню из радио, поле даты, приглушённая подпись; скрытые поля; Show для работы без сценария.
// Чего нет. Своего вида: форма не рисуется, её поля стоят в полосе как свои; Show виден, только пока страница без сценария (data-tc-noscript на html, его снимает tc.js).
type BarForm struct {
	Items  []BarFormPart
	Action string   // where the form goes (GET)
	Hidden []Hidden // the list's other choices that go with it: the segment chosen
	Show   Button   // what sends it on a page without a script: Show, a submit
}

// BarFormPart is what a bar's filter form holds.
type BarFormPart interface {
	part
	barFormPart()
}

func (Drop) barFormPart()  {}
func (Input) barFormPart() {}
func (Pick) barFormPart()  {}
func (Dim) barFormPart()   {}

// Render prints the form.
func (f BarForm) Render() string { return f.html(nil) }

func (f BarForm) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<form` + class(m, "tc-barform") + attr("action", f.Action) + ` method="get">`)
	for _, it := range f.Items {
		s.WriteString(it.html(nil))
	}
	for _, h := range f.Hidden {
		s.WriteString(h.html(nil))
	}
	show := f.Show
	show.Submit = true
	s.WriteString(`<span class="tc-barform__show">` + show.html(nil) + `</span></form>`)
	return s.String()
}
