package widget

import "strings"

// SegOption is one option of a segmented switch.
type SegOption struct {
	Label string
	Icon  Icon
	On    bool
	Title string // the accessible name of an option without a label
	Value string // what the option sends when the switch has a Name
	Href  string // the option is a link: a bar's order or filter by kind
	Count string // how many the option holds, after its label ("All 24")
}

// SegVariant is the drawn size and alignment of a switch.
type SegVariant int

const (
	SegDefault      SegVariant = iota
	SegStart                   // aligned to the start of a form row
	SegCompact                 // 22 px high: inside a row
	SegStartSmall              // aligned to the start, 22 px, 11.5 px text
	SegCompactSmall            // 22 px, 11.5 px text: over a table in a text
)

func (v SegVariant) class() string {
	return [...]string{"", "tc-seg--start", "tc-seg--h22", "tc-seg--start tc-seg--h22 tc-seg--t11-5", "tc-seg--h22 tc-seg--t11-5"}[v]
}

// Seg — tc-seg
//
// Роль. Сомкнутый переключатель: один из нескольких вариантов, все на виду.
// Вход. Seg{Options, Variant, Place, Secondary, Name, Theme, Multiple}.
// Состояния. default | start | compact | icons | radio | links | several.
// Инварианты. Выбран ровно один вариант, у ряда флажков (Multiple) — любые; у варианта без подписи есть доступное имя; переключатель печатает фрагмент своего кадра.
// Поведение. Нажатие выбирает вариант; стрелки ходят по вариантам внутри переключателя.
// Эталон. кадр `components-13`, фрагмент `.tc-seg`.
// Содержимое. Варианты: подпись, значок или оба, число; вариант — кнопка, радио формы или ссылка.
// Чего нет. Выпадающего списка: вариантов столько, сколько помещается; тесно — SegFold, больше — выбор (Pick).
type Seg struct {
	Options   []SegOption
	Variant   SegVariant
	Place     Place
	Secondary bool
	// Name makes the options radio buttons of a form: the chosen one sends
	// its Value, and the switch works without scripts.
	Name string
	// Theme makes the switch the theme's: options auto, light, dark; the
	// library's script keeps the choice in this browser (tc-theme).
	Theme bool
	// Multiple makes a switch with a Name a row of checkboxes: several
	// options at once, each ticked one sent (the days of a repeat).
	Multiple bool
}

// Render prints the switch.
func (g Seg) Render() string { return g.html(nil) }

func (g Seg) html(m mix) string { return g.seg(m) }

// SegFold — tc-segfold
//
// Роль. Сегмент, которому тесно: на широкой полосе — сомкнутый переключатель ссылок, на узкой — выпадающий выбор с текущим вариантом (решение 30.09.2026).
// Вход. SegFold{Seg, Phone}.
// Состояния. default.
// Инварианты. Выбран ровно один вариант; у каждого варианта адрес; складывается там же, где полоса прячет второстепенное (is-2nd); виджет печатает фрагмент своего кадра.
// Поведение. На узкой полосе нажатие на текущий вариант открывает меню остальных (details, без сценария); выбор — переход по ссылке. Короткий переключатель (`Phone`) складывается только на полосе телефона, 640&nbsp;px и уже: Day / Week / Month / Range агенды.
// Эталон. кадр `s-contacts-7`, фрагмент `.tc-segfold`.
// Содержимое. Варианты: подпись и число; сегмент и меню из одних и тех же вариантов.
// Чего нет. Кнопок и радио: вариант сложенного сегмента — всегда ссылка.
type SegFold struct {
	Seg Seg // the switch of links; its options are the fold's menu too
	// Phone folds it on a phone's bar only (640 px and less): a short switch
	// that a wider bar holds — the agenda's Day / Week / Month / Range.
	Phone bool
}

// Render prints the switch and its fold.
func (f SegFold) Render() string { return f.html(nil) }

func (f SegFold) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<span` + class(m, "tc-segfold", pick(f.Phone, "tc-segfold--phone")) + `>` + f.Seg.seg(nil))
	cur := ""
	for _, o := range f.Seg.Options {
		if o.On {
			cur = o.label()
		}
	}
	s.WriteString(`<details class="tc-segfold__pick"><summary class="tc-segfold__current">` + esc(cur) + `</summary><div class="tc-segfold__menu">`)
	for _, o := range f.Seg.Options {
		s.WriteString(`<a` + class(nil, is("on", o.On), "tc-segfold__item") + attr("href", o.Href) + attr("aria-current", pick(o.On, "true")) + `>` + esc(o.Label))
		if o.Count != "" {
			s.WriteString(`<span class="tc-segfold__count">` + esc(o.Count) + `</span>`)
		}
		s.WriteString(`</a>`)
	}
	s.WriteString(`</div></details></span>`)
	return s.String()
}

// label is the option's words with its count.
func (o SegOption) label() string {
	if o.Count == "" {
		return o.Label
	}
	return o.Label + " " + o.Count
}

func (g Seg) seg(m mix) string {
	var s strings.Builder
	s.WriteString(`<span` + class(m, "tc-seg", g.Variant.class(), g.Place.class(), rank(g.Secondary)) + ` role="radiogroup"` + flag("data-tc-theme-switch", g.Theme) + `>`)
	for _, o := range g.Options {
		if g.Name != "" {
			kind := "radio"
			if g.Multiple {
				kind = "checkbox"
			}
			s.WriteString(`<label class="tc-seg__opt"><input class="tc-seg__box" type="` + kind + `"` + attr("name", g.Name) + attr("value", o.Value) +
				attr("aria-label", pick(o.Label == "", o.Title)) + flag("checked", o.On) + `>`)
			if o.Icon != "" {
				s.WriteString(`<svg width="13" height="13" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.4" fill="none">` + o.Icon.use() + `</svg>`)
			}
			s.WriteString(esc(o.Label) + `</label>`)
			continue
		}
		if o.Href != "" {
			s.WriteString(`<a` + class(nil, is("on", o.On), "tc-seg__opt") + attr("href", o.Href) + attr("aria-current", pick(o.On, "true")) + `>` + esc(o.label()) + `</a>`)
			continue
		}
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
// Вход. Check{On, Parts, Label, Tertiary, Name, Value, Disables}.
// Состояния. on | off | rich | tertiary | box.
// Инварианты. Включено или нет — одно поле; флажок — кнопка с подписью или флажок формы (Name), у него есть доступное имя; флажок печатает фрагмент своего кадра.
// Поведение. Нажатие переключает; пробел — тоже.
// Эталон. кадр `components-10`, фрагмент `.tc-check`.
// Содержимое. Подпись: текст, приглушённое пояснение, моноширинный кусок, бейдж, число внутри фразы, сторона конфликта.
// Чего нет. Третьего состояния: частичный выбор — у строки источника (tc-tickrow), не у флажка.
type Check struct {
	On    bool
	Parts []CheckPart
	Label string // the accessible name when the parts are not plain words
	// Name makes the check a checkbox of a form: ticked, it sends Value, and
	// it switches without scripts.
	Name  string
	Value string
	// Disables names the form's fields the check switches off while it is
	// ticked, separated by spaces as classes are: «The server asks no
	// password» and Password; All day and «start.time due.time» (Carrel's
	// task 111).
	Disables string
	// Tertiary goes under «⋯» on a phone when the check stands in a bar.
	Tertiary bool
}

// Render prints the check.
func (c Check) Render() string { return c.html(nil) }

func (c Check) html(m mix) string {
	var s strings.Builder
	if c.Name != "" {
		s.WriteString(`<label` + class(m, "tc-check", rank3(c.Tertiary)) + `><input class="tc-check__box" type="checkbox"` + attr("name", c.Name) + attr("value", c.Value) +
			attr("aria-label", c.Label) + flag("checked", c.On) + attr("data-tc-disables", c.Disables) + `><span` + class(nil, "tc-tickrow__tick", pick(!c.On, "is-off")) + `></span>`)
		for _, p := range c.Parts {
			s.WriteString(p.html(nil))
		}
		s.WriteString(`</label>`)
		return s.String()
	}
	s.WriteString(`<button` + class(m, "tc-check", rank3(c.Tertiary)) + ` type="button" role="checkbox"` + attr("aria-checked", boolText(c.On)) +
		attr("aria-label", c.Label) + `><span` + class(nil, "tc-tickrow__tick", pick(!c.On, "is-off")) + `></span>`)
	for _, p := range c.Parts {
		s.WriteString(p.html(nil))
	}
	s.WriteString(`</button>`)
	return s.String()
}

// CheckForm — tc-checkform
//
// Роль. Отметки полосы, которые сами — форма: «Compact rows», «Events · Tasks due · Notes» уходят GET-запросом с остальными выборами списка.
// Вход. CheckForm{Checks, Action, Hidden, Apply, Tertiary}.
// Состояния. default.
// Инварианты. Отметки — флажки формы (Check.Name); форма отправляет её и скрытые поля; без сценария — кнопкой Apply, со сценарием — при смене; виджет печатает фрагмент своего кадра.
// Поведение. Смена отметки отправляет форму (сценарий); без сценария — Apply рядом.
// Эталон. кадр `s-contacts-1`, фрагмент `.tc-checkform`.
// Содержимое. Флажок с подписью, скрытые поля, Apply для работы без сценария; в меню «⋯» — те же отметки пунктами-флажками (двойник отметок полосы, Carrel's task 82).
// Чего нет. Своего вида: вид у флажка; Apply виден, только пока страница без сценария (data-tc-noscript на html, его снимает tc.js).
type CheckForm struct {
	Checks []Check  // the checkboxes of the form: Name and Value set
	Action string   // where the form goes (GET)
	Hidden []Hidden // the list's other choices that go with it
	Apply  Button   // what sends it on a page without a script: Apply, small, a submit
	// Tertiary goes under «⋯» on a phone, with the checks.
	Tertiary bool
}

// Render prints the form.
func (f CheckForm) Render() string { return f.html(nil) }

func (f CheckForm) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<form` + class(m, "tc-checkform", rank3(f.Tertiary)) + attr("action", f.Action) + ` method="get">`)
	for _, c := range f.Checks {
		s.WriteString(c.html(nil))
	}
	for _, h := range f.Hidden {
		s.WriteString(h.html(nil))
	}
	s.WriteString(`<span class="tc-checkform__apply">` + f.Apply.html(nil) + `</span></form>`)
	return s.String()
}

// inMenu prints the form inside a menu — the twin of a bar's checks under
// «⋯»: the same GET form, its checks menu items with tick boxes.
func (f CheckForm) inMenu() string {
	var s strings.Builder
	s.WriteString(`<form` + class(nil, "tc-checkform", rank3(f.Tertiary)) + attr("action", f.Action) + ` method="get">`)
	for _, c := range f.Checks {
		t := TickedOff
		if c.On {
			t = TickedOn
		}
		label := c.Label
		if label == "" {
			for _, p := range c.Parts {
				if x, ok := p.(Text); ok {
					label += string(x)
				}
			}
		}
		s.WriteString(MenuItem{Name: c.Name, Value: c.Value, Tick: &TickBox{Ticked: t}, Label: label}.html(nil))
	}
	for _, h := range f.Hidden {
		s.WriteString(h.html(nil))
	}
	s.WriteString(`<span class="tc-checkform__apply">` + f.Apply.html(nil) + `</span></form>`)
	return s.String()
}

// Range — tc-range
//
// Роль. Диапазон дат: два поля, стрелка между ними и кнопка, которая показывает.
// Вход. Range{From, To, Show, Wrap, Secondary, Action, Hidden}.
// Состояния. default | wrap | form.
// Инварианты. Начало — слева, конец — справа, между ними стрелка; диапазон печатает фрагмент своего кадра.
// Поведение. Show применяет диапазон; Enter в поле — то же. С адресом (Action) диапазон — форма GET: From и To уходят полями, Show её отправляет, без сценария.
// Эталон. кадр `components-18`, фрагмент `.tc-range`.
// Содержимое. Два поля и кнопка.
// Чего нет. Выбора диапазона в одном поле: две даты — два поля.
type Range struct {
	From      Input
	To        Input
	Show      Button
	Wrap      bool // the range wraps inside a narrow menu
	Secondary bool
	Action    string // the range is a form that sends From and To there (GET)
	// Hidden are the list's other choices that go with it: view=range
	// (Carrel's task 107).
	Hidden []Hidden
}

// Render prints the range.
func (r Range) Render() string { return r.html(nil) }

func (r Range) html(m mix) string {
	tag, form := "span", ""
	if r.Action != "" {
		tag, form = "form", attr("action", r.Action)+` method="get"`
		r.Show.Submit = true
	}
	hidden := ""
	if r.Action != "" {
		for _, h := range r.Hidden {
			hidden += h.html(nil)
		}
	}
	return `<` + tag + class(m, "tc-range", pick(r.Wrap, "tc-range--wrap"), rank(r.Secondary)) + form + `>` + hidden + r.From.html(nil) +
		`<i class="tc-range__dash">→</i>` + r.To.html(nil) + r.Show.html(nil) + `</` + tag + `>`
}

func boolText(on bool) string {
	if on {
		return "true"
	}
	return "false"
}
