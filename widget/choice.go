package widget

import "strings"

// SegOption is one option of a segmented switch.
type SegOption struct {
	Label string
	Icon  Icon
	On    bool
	Title string // the accessible name of an option without a label
}

// SegVariant is the drawn size and alignment of a switch.
type SegVariant int

const (
	SegDefault    SegVariant = iota
	SegStart                 // aligned to the start of a form row
	SegCompact               // 22 px high: inside a row
	SegStartSmall            // aligned to the start, 22 px, 11.5 px text
)

func (v SegVariant) class() string {
	return [...]string{"", "tc-seg--start", "tc-seg--h22", "tc-seg--start tc-seg--h22 tc-seg--t11-5"}[v]
}

// Seg — tc-seg
//
// Роль. Сомкнутый переключатель: один из нескольких вариантов, все на виду.
// Вход. Seg{Options, Variant, Place, Secondary}.
// Состояния. default | start | compact | icons.
// Инварианты. Выбран ровно один вариант; у варианта без подписи есть доступное имя; переключатель печатает фрагмент своего кадра.
// Поведение. Нажатие выбирает вариант; стрелки ходят по вариантам внутри переключателя.
// Эталон. кадр `components-13`, фрагмент `.tc-seg`.
// Содержимое. Варианты: подпись, значок или оба.
// Чего нет. Выпадающего списка: вариантов столько, сколько помещается; больше — это выбор (Pick).
type Seg struct {
	Options   []SegOption
	Variant   SegVariant
	Place     Place
	Secondary bool
}

// Render prints the switch.
func (g Seg) Render() string { return g.html(nil) }

func (g Seg) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<span` + class(m, "tc-seg", g.Variant.class(), g.Place.class(), rank(g.Secondary)) + ` role="radiogroup">`)
	for _, o := range g.Options {
		s.WriteString(`<button` + class(nil, is("on", o.On), "tc-seg__opt") + ` type="button" role="radio"` +
			attr("aria-checked", pick(o.On, "true")) + attr("aria-label", pick(o.Label == "", o.Title)) + `>`)
		if o.Icon != "" {
			s.WriteString(`<svg width="13" height="13" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.4" fill="none">` + o.Icon.use() + `</svg>`)
		}
		s.WriteString(esc(o.Label) + `</button>`)
	}
	s.WriteString(`</span>`)
	return s.String()
}

// CheckPart is what a check's label holds: running text, a badge, a number
// field inside the sentence, a choice of a conflict.
type CheckPart interface {
	part
	checkPart()
}

func (Text) checkPart()   {}
func (Dim) checkPart()    {}
func (Mono) checkPart()   {}
func (Badge) checkPart()  {}
func (Choice) checkPart() {}

// Choice is one side of a conflict in a check's label: where the value comes
// from, then the value.
type Choice struct {
	Source string
	Value  string
}

func (c Choice) html(mix) string {
	return `<span class="tc-diff__value"><span class="tc-diff__source">` + esc(c.Source) + `</span>` + esc(c.Value) + `</span>`
}

// Check — tc-check
//
// Роль. Флажок с подписью: включено или нет.
// Вход. Check{On, Parts, Label, Tertiary}.
// Состояния. on | off | rich | tertiary.
// Инварианты. Включено или нет — одно поле; флажок — кнопка с подписью, у неё есть доступное имя; флажок печатает фрагмент своего кадра.
// Поведение. Нажатие переключает; пробел — тоже.
// Эталон. кадр `components-10`, фрагмент `.tc-check`.
// Содержимое. Подпись: текст, приглушённое пояснение, моноширинный кусок, бейдж, число внутри фразы, сторона конфликта.
// Чего нет. Третьего состояния: частичный выбор — у строки источника (tc-tickrow), не у флажка.
type Check struct {
	On    bool
	Parts []CheckPart
	Label string // the accessible name when the parts are not plain words
	// Tertiary goes under «⋯» on a phone when the check stands in a bar.
	Tertiary bool
}

// Render prints the check.
func (c Check) Render() string { return c.html(nil) }

func (c Check) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<button` + class(m, "tc-check", rank3(c.Tertiary)) + ` type="button" role="checkbox"` + attr("aria-checked", boolText(c.On)) +
		attr("aria-label", c.Label) + `><span` + class(nil, "tc-tickrow__tick", pick(!c.On, "is-off")) + `></span>`)
	for i, p := range c.Parts {
		if i > 0 {
			s.WriteString(" ")
		}
		s.WriteString(p.html(nil))
	}
	s.WriteString(`</button>`)
	return s.String()
}

// Range — tc-range
//
// Роль. Диапазон дат: два поля, стрелка между ними и кнопка, которая показывает.
// Вход. Range{From, To, Show, Wrap, Secondary}.
// Состояния. default | wrap.
// Инварианты. Начало — слева, конец — справа, между ними стрелка; диапазон печатает фрагмент своего кадра.
// Поведение. Show применяет диапазон; Enter в поле — то же.
// Эталон. кадр `components-18`, фрагмент `.tc-range`.
// Содержимое. Два поля и кнопка.
// Чего нет. Календаря в одном поле: две даты — два поля.
type Range struct {
	From      Input
	To        Input
	Show      Button
	Wrap      bool // the range wraps inside a narrow menu
	Secondary bool
}

// Render prints the range.
func (r Range) Render() string { return r.html(nil) }

func (r Range) html(m mix) string {
	return `<span` + class(m, "tc-range", pick(r.Wrap, "tc-range--wrap"), rank(r.Secondary)) + `>` + r.From.html(nil) +
		`<i class="tc-range__dash">→</i>` + r.To.html(nil) + r.Show.html(nil) + `</span>`
}

func boolText(on bool) string {
	if on {
		return "true"
	}
	return "false"
}
