package widget

import "strings"

// Ticked is the state of a tick box in front of a row: the row is chosen,
// part of what is under it is, or nothing is.
type Ticked int

const (
	TickedOn Ticked = iota
	TickedPart
	TickedOff
)

// Send is what pressing a tick box sends. With Form it is a submit button of
// that form, a toggle that acts at once — a source of the rail, a card's
// choice; with a Name only, a checkbox of the form it stands in, a choice
// acted on later — rows before a bulk action; Of names that form when the
// checkbox stands outside it (a PostForm the selection bar's buttons send:
// Carrel's task 148). Empty, the box is the script's.
type Send struct {
	Form  string
	Name  string
	Value string
	Of    string // the ID of the form a checkbox belongs to without standing in it
	One   bool   // one of several under the Name: a radio button, not a checkbox
}

// box prints the tick box; label is what is ticked.
func (t Ticked) box(label string, send Send, more ...string) string {
	if send.Name != "" && send.Form == "" {
		own := append([]string{"tc-tickrow__tick", [...]string{"", "is-part", "is-off"}[t]}, more...)
		kind := "checkbox"
		if send.One {
			kind = "radio"
		}
		return `<input` + class(nil, own...) + ` type="` + kind + `"` + attr("form", send.Of) + attr("name", send.Name) + attr("value", send.Value) + attr("aria-label", label) +
			flag("checked", t == TickedOn) + `>`
	}
	own := append([]string{"tc-tickrow__tick", [...]string{"", "is-part", "is-off"}[t]}, more...)
	typ := "button"
	if send.Form != "" {
		typ = "submit"
	}
	return `<button` + class(nil, own...) + ` type="` + typ + `"` + attr("form", send.Form) + attr("name", send.Name) + attr("value", send.Value) + attr("aria-label", label) +
		attr("aria-pressed", [...]string{"true", "mixed", "false"}[t]) + `></button>`
}

// TickRowKind is what a row of the rail stands for.
type TickRowKind int

const (
	TickRowOne      TickRowKind = iota // one source
	TickRowAll                         // all the sources at once: the combined view
	TickRowRoot                        // the root of a tree of sources: an icon, not a mark
	TickRowExternal                    // a source read from outside
)

func (k TickRowKind) class() string {
	return [...]string{"", "is-all", "is-root", "is-ext"}[k]
}

// TickRowForm is where the row stands.
type TickRowForm int

const (
	TickRowRail   TickRowForm = iota // a row of the rail
	TickRowGroup                     // the head of a group in a grid of choices: flush, bold
	TickRowNested                    // a row under such a head: indented
)

func (f TickRowForm) class() string {
	return [...]string{"", "tc-tickrow--padl0", "tc-tickrow--inset"}[f]
}

// TickRowTrail is what stands at the end of a row of the rail.
type TickRowTrail interface {
	part
	tickRowTrail()
}

func (Badge) tickRowTrail() {}
func (Skel) tickRowTrail()  {}

func (TickRow) fieldGridPart() {}

// TickRow — tc-tickrow
//
// Роль. Строка источника с галкой: показать его или скрыть, выбрать для копии, для выгрузки.
// Вход. TickRow{Ticked, Kind, Form, Mark, Icon, Label, Meta, MetaTitle, Current, Hidden, Error, Trail, Send}.
// Состояния. one | all | root | external | current | hidden | error | group | nested | loading.
// Инварианты. Галка — кнопка с доступным именем; её состояния, род строки и место строки — перечисления; строка печатает фрагмент своего кадра.
// Поведение. Галка с Send.Form — кнопка формы рейла: нажатие отправляет выбор, показать источник или скрыть. Галка группы (TickRowAll, TickRowRoot) — такая же кнопка: отправляет «все» или «ни одного»; галки под ней ставит ответ сервера, не сценарий (владелец 07.11). С живым слоем и формой рейла, которая спрашивает зону (PostForm.Target), ответ ставит на место рейл и список без загрузки страницы; без слоя приходит страница.
// Эталон. кадр `components-11`, фрагмент `.tc-tickrow`.
// Содержимое. Галка, метка источника (у корня — значок), название, число, значок ошибки с причиной в подсказке, бейдж, кнопка повтора или заготовка.
// Чего нет. Своего меню: действия с источником — в меню раздела рейла.
type TickRow struct {
	Ticked    Ticked
	Kind      TickRowKind
	Form      TickRowForm
	Mark      Mark
	Icon      Icon // the root's icon instead of the mark
	Label     string
	Meta      string // a count, a kind, the reason of an error; empty for none
	MetaTitle string // the reason in full
	Current   bool   // the source open now
	Hidden    bool   // the source is not shown
	Error     bool   // the source did not answer
	Trail     TickRowTrail
	Send      Send // what pressing the tick sends: the source's choice without a script
}

// Render prints the row.
func (r TickRow) Render() string { return r.html(nil) }

func (r TickRow) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-tickrow", r.Kind.class(), is("on", r.Current), is("off", r.Hidden), is("error", r.Error), r.Form.class()) + `>`)
	s.WriteString(r.Ticked.box(r.Label, r.Send))
	if r.Icon != "" {
		s.WriteString(`<svg viewBox="0 0 24 24">` + r.Icon.use() + `</svg>`)
	} else {
		s.WriteString(r.Mark.html(nil))
	}
	s.WriteString(`<span` + class(nil, "tc-tickrow__label", pick(r.Form == TickRowGroup, "tc-tickrow__label--fw600")) + `>` + esc(r.Label) + `</span>`)
	if r.Error && r.Meta == "" {
		// A source that did not answer: a mark of the error, the reason in
		// its hint; in full — in the runbar's menu of refusals (02.10).
		s.WriteString(`<span class="tc-tickrow__err" role="img"` + attr("aria-label", r.MetaTitle) + attr("title", r.MetaTitle) + `><svg viewBox="0 0 24 24">` + Icon("alert").use() + `</svg></span>`)
	} else if r.Meta != "" {
		s.WriteString(`<span class="tc-tickrow__meta"` + attr("title", r.MetaTitle) + `>` + esc(r.Meta) + `</span>`)
	}
	if r.Trail != nil {
		s.WriteString(r.Trail.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// SkelWidth is the width of a placeholder.
type SkelWidth int

const (
	SkelFill SkelWidth = iota // the width its place gives it
	Skel16                    // 16 px: a count
	Skel45                    // 45 %
	Skel52                    // 52 %
	Skel60                    // 60 %
)

func (w SkelWidth) class() string {
	return [...]string{"", "tc-skel--w16", "tc-skel--w45pc", "tc-skel--w52pc", "tc-skel--w60pc"}[w]
}

// Skel — tc-skel
//
// Роль. Заготовка на месте того, что ещё грузится: число в рейле, ячейка строки.
// Вход. Skel{Width}.
// Состояния. fill | sized.
// Инварианты. Ширина — перечисление; заготовка печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `states-3`, фрагмент `.tc-skel`.
// Содержимое. нет.
// Чего нет. Текста и движения: заготовка стоит, пока не придут данные.
type Skel struct {
	Width SkelWidth
}

// Render prints the placeholder.
func (k Skel) Render() string { return k.html(nil) }

func (k Skel) html(m mix) string {
	return `<span` + class(m, "tc-skel", k.Width.class()) + `></span>`
}
