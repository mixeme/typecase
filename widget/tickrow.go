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

// box prints the tick box; label is what is ticked.
func (t Ticked) box(label string, more ...string) string {
	own := append([]string{"tc-tickrow__tick", [...]string{"", "is-part", "is-off"}[t]}, more...)
	return `<button` + class(nil, own...) + ` type="button"` + attr("aria-label", label) +
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
	TickRowTouch                     // a row of a phone: 44 px high
)

func (f TickRowForm) class() string {
	return [...]string{"", "tc-tickrow--padl0", "tc-tickrow--padl14", "tc-tickrow--h44"}[f]
}

// TickRowTrail is what stands at the end of a row of the rail.
type TickRowTrail interface {
	part
	tickRowTrail()
}

func (Badge) tickRowTrail()  {}
func (Button) tickRowTrail() {}
func (Skel) tickRowTrail()   {}

func (TickRow) fieldGridPart() {}

// TickRow — tc-tickrow
//
// Роль. Строка источника с галкой: показать его или скрыть, выбрать для копии, для выгрузки.
// Вход. TickRow{Ticked, Kind, Form, Mark, Icon, Label, Meta, MetaTitle, Current, Hidden, Error, Trail}.
// Состояния. one | all | root | external | current | hidden | error | group | nested | touch | loading.
// Инварианты. Галка — кнопка с доступным именем; её состояния, род строки и место строки — перечисления; строка печатает фрагмент своего кадра.
// Поведение. Нажатие на галку показывает или скрывает источник; галка группы ставит или снимает все галки под ней.
// Эталон. кадр `components-11`, фрагмент `.tc-tickrow`.
// Содержимое. Галка, метка источника (у корня — значок), название, число или причина ошибки, бейдж, кнопка повтора или заготовка.
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
}

// Render prints the row.
func (r TickRow) Render() string { return r.html(nil) }

func (r TickRow) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-tickrow", r.Kind.class(), is("on", r.Current), is("off", r.Hidden), is("error", r.Error), r.Form.class()) + `>`)
	s.WriteString(r.Ticked.box(r.Label))
	if r.Icon != "" {
		s.WriteString(`<svg viewBox="0 0 24 24">` + r.Icon.use() + `</svg>`)
	} else {
		s.WriteString(r.Mark.html(nil))
	}
	s.WriteString(`<span` + class(nil, "tc-tickrow__label", pick(r.Form == TickRowGroup, "tc-tickrow__label--fw600")) + `>` + esc(r.Label) + `</span>`)
	if r.Meta != "" {
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
// Эталон. кадр `m-states-3`, фрагмент `.tc-skel`.
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
