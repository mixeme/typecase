package widget

import (
	"strconv"
	"strings"
)

// FieldPart is what a form field holds: its controls and its own elements —
// the hint, the swatches, the zoom, a row, a group, the photo.
type FieldPart interface {
	part
	fieldPart()
}

func (Input) fieldPart()     {}
func (Pick) fieldPart()      {}
func (Area) fieldPart()      {}
func (TagInput) fieldPart()  {}
func (Seg) fieldPart()       {}
func (Check) fieldPart()     {}
func (Button) fieldPart()    {}
func (Line) fieldPart()      {}
func (FieldHint) fieldPart() {}
func (Swatches) fieldPart()  {}
func (Zoom) fieldPart()      {}
func (FieldRow) fieldPart()  {}
func (Group) fieldPart()     {}
func (Photo) fieldPart()     {}
func (Mark) fieldPart()      {}
func (Dim) fieldPart()       {}
func (Drop) fieldPart()      {}
func (Menu) fieldPart()      {}

// HintVariant is the drawn form of a field's hint.
type HintVariant int

const (
	HintPlain  HintVariant = iota
	HintNarrow             // at most 150 px: beside a photo
	HintSmall              // 11 px, dusty
	HintDanger             // what is wrong with the value
)

func (v HintVariant) class() string {
	return [...]string{"", "tc-field__hint--max150", "tc-field__hint--t11-dusty", "tc-dim--danger"}[v]
}

// FieldHint is the line under a field that says what the value means.
type FieldHint struct {
	Parts   []Inline
	Variant HintVariant
}

func (h FieldHint) html(mix) string {
	var s strings.Builder
	s.WriteString(`<span` + class(nil, "tc-field__hint", h.Variant.class()) + `>`)
	for _, p := range h.Parts {
		s.WriteString(p.html(nil))
	}
	s.WriteString(`</span>`)
	return s.String()
}

// Swatches is a choice of colour: one swatch per colour, one of them chosen.
type Swatches struct {
	Colours []string
	On      int // the index of the chosen colour
}

func (w Swatches) html(mix) string {
	var s strings.Builder
	s.WriteString(`<span class="tc-line tc-line--gap-9 tc-field__swatches" role="radiogroup">`)
	for i, c := range w.Colours {
		s.WriteString(`<span` + attr("data-tc-colour", c) + class(nil, "tc-field__swatch", pick(i == w.On, "tc-field__swatch--on")) +
			` role="radio"` + attr("aria-checked", boolText(i == w.On)) + attr("aria-label", c) + `></span>`)
	}
	s.WriteString(`</span>`)
	return s.String()
}

// Zoom is a zoom slider: the rail, its fill and knob, the value after it.
type Zoom struct {
	Value string // "1.6×"
}

func (z Zoom) html(mix) string {
	return `<span class="tc-line tc-line--gap-9 tc-field__zoom"><span class="tc-field__rail"><span class="tc-field__fill"></span>` +
		`<span class="tc-field__knob"></span></span>` + Mono{Text: z.Value, Variant: Mono11Dusty}.html(nil) + `</span>`
}

// FieldRow is a row of small things set in a field: marks of a colour with a
// note after them.
type FieldRow struct {
	Items []FieldPart
}

func (r FieldRow) html(mix) string {
	var s strings.Builder
	s.WriteString(`<span class="tc-line tc-line--gap-6 tc-field__row">`)
	for _, it := range r.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</span>`)
	return s.String()
}

// GroupForm is how a group of checks sits under a label.
type GroupForm int

const (
	GroupRow         GroupForm = iota // side by side, 14 px apart
	GroupColumn                       // one under another, 2 px apart
	GroupColumnLoose                  // one under another, 5 px apart
)

// Group is several controls under one label.
type Group struct {
	Form  GroupForm
	Items []FieldPart
}

func (g Group) html(mix) string {
	var s strings.Builder
	s.WriteString([...]string{
		`<div class="tc-line tc-line--gap-14 tc-card__top2">`,
		`<div class="tc-stack tc-stack--gap-2 tc-card__top2">`,
		`<div class="tc-stack tc-stack--gap-5 tc-card__top2">`}[g.Form])
	for _, it := range g.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// Photo is a field's photo: the picture's place and, beside it, what can be
// done with it.
type Photo struct {
	Initials string // shown while there is no photo
	Side     []FieldPart
	Src      string // the photo's address: the picture fills the place
}

func (p Photo) html(mix) string {
	var s strings.Builder
	if p.Src != "" {
		s.WriteString(`<div class="tc-field__photo-row"><span class="tc-avatar--photo"><img class="tc-avatar__img"` + attr("src", p.Src) +
			` alt=""></span><div class="tc-stack tc-stack--gap-6 tc-field__photo-side">`)
	} else {
		s.WriteString(`<div class="tc-field__photo-row"><span` + class(nil, "tc-avatar--photo", is("ph", p.Initials != "")) + `>` +
			esc(p.Initials) + `</span><div class="tc-stack tc-stack--gap-6 tc-field__photo-side">`)
	}
	for _, it := range p.Side {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div></div>`)
	return s.String()
}

// LabelPart is what a field's label holds after its words.
type LabelPart interface {
	part
	labelPart()
}

func (Dim) labelPart()   {}
func (Plain) labelPart() {}

// FieldWidth is the width of a field in its form row.
type FieldWidth int

const (
	FieldAuto FieldWidth = iota
	FieldXS
	FieldSM
	FieldMD
	FieldLG
	FieldXL
	FieldFull
	FieldW100
	FieldW280
	FieldFix96
	FieldFix130
	FieldFix150
	FieldFixAuto
	FieldFill120
	FieldFill200
	FieldFill300
	FieldFlex2_280
	FieldFlex3_320
	FieldFlex3_340
	FieldMax640Fill460
	FieldInline // label and control in one line
)

func (w FieldWidth) class() string {
	return [...]string{"", "tc-field--xs", "tc-field--sm", "tc-field--md", "tc-field--lg", "tc-field--xl", "tc-field--full",
		"tc-field--w100", "tc-field--w280", "tc-field--fix96", "tc-field--fix130", "tc-field--fix150", "tc-field--fixauto",
		"tc-field--fill120", "tc-field--fill200", "tc-field--fill300", "tc-field--flex2-280", "tc-field--flex3-320",
		"tc-field--flex3-340", "tc-field--max640-fill460", "tc-field--inline"}[w]
}

// Field — tc-field
//
// Роль. Поле формы: подпись и контролы под ней, подсказка о том, что значит значение.
// Вход. Field{Label, LabelAside, LabelSpan, Parts, Width, Gap, OwnLine, Place}.
// Состояния. labelled | bare | hint | group | swatches | zoom | photo | row | own-line.
// Инварианты. Ширина — перечисление: поле даты не растягивается, растягивается поле тегов; подпись связана с контролом; поле печатает фрагмент своего кадра.
// Поведение. нет: поведение у контролов.
// Эталон. кадр `components-10`, фрагмент `.tc-field`.
// Содержимое. Контролы (ввод, выбор, текст, теги, сегмент, флажки, кнопки, строка), подсказка, образцы цвета, масштаб, ряд меток, группа флажков, фотография, выпадашка и меню под полем.
// Чего нет. Проверки значения: что неверно, говорит подсказка, которую ставит экран.
type Field struct {
	Label      string // " " keeps an empty label: the row lines up with its neighbours
	LabelAside LabelPart
	LabelSpan  bool // the label is not for one control: the photo field
	Parts      []FieldPart
	Width      FieldWidth
	Gap        int // 7 or 8 between several checks; 0 is the form's
	OwnLine    bool
	Place      Place
}

// Render prints the field.
func (f Field) Render() string { return f.html(nil) }

func (f Field) html(m mix) string {
	var s strings.Builder
	gap := ""
	if f.Gap != 0 {
		gap = "tc-field--gap" + strconv.Itoa(f.Gap)
	}
	s.WriteString(`<div` + class(m, "tc-field", is("own-line", f.OwnLine), f.Width.class(), gap, f.Place.class()) + `>`)
	switch {
	case f.LabelSpan:
		s.WriteString(`<span class="tc-field__label">` + esc(f.Label) + `</span>`)
	case f.Label != "":
		s.WriteString(`<label>` + esc(f.Label))
		if f.LabelAside != nil {
			s.WriteString(" " + f.LabelAside.html(nil))
		}
		s.WriteString(`</label>`)
	}
	for _, p := range f.Parts {
		s.WriteString(p.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// FieldSetPart is what a group of fields holds after its rubric.
type FieldSetPart interface {
	part
	fieldSetPart()
}

func (Field) fieldSetPart()    {}
func (Button) fieldSetPart()   {}
func (Template) fieldSetPart() {}

// Template is the row a button «Add …» adds (Button.Adds): its fields, with
// «{n}» in their names for the library's script to replace by the smallest
// number the form does not use yet. It is not drawn.
type Template struct {
	ID    string
	Items []FieldSetPart
}

func (t Template) html(mix) string {
	var s strings.Builder
	s.WriteString(`<template` + attr("id", t.ID) + `>`)
	for _, it := range t.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</template>`)
	return s.String()
}
func (Line) fieldSetPart()      {}
func (Grow) fieldSetPart()      {}
func (FieldGrid) fieldSetPart() {}

// FieldGrid is a grid of choices in a group of fields: sources to copy.
type FieldGrid struct {
	Items []FieldGridPart
}

// FieldGridPart is what a group's grid holds.
type FieldGridPart interface {
	part
	fieldGridPart()
}

func (g FieldGrid) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-grow tc-fset__grid">`)
	for _, it := range g.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// FieldSetVariant is the drawn form of a group of fields.
type FieldSetVariant int

const (
	FieldSetPlain FieldSetVariant = iota // a rule above: one group after another
	FieldSetLead                         // the first group: no rule above
	FieldSetRule                         // a rule the full width
)

func (v FieldSetVariant) class() string {
	return [...]string{"", "tc-fset--lead", "tc-fset--rule"}[v]
}

// FieldSet — tc-fset
//
// Роль. Группа полей формы под рубрикой: «Phone numbers», «Email», «Password».
// Вход. FieldSet{Rubric, Items, Variant}.
// Состояния. plain | lead.
// Инварианты. Рубрика — первой, поля — после неё; первая группа формы без линейки сверху; группа печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `s-contacts-3`, фрагмент `.tc-fset`.
// Содержимое. Поля, кнопки, строки, переносы, сетка выбора, сообщение, таблица.
// Чего нет. Сворачивания: группа всегда открыта.
type FieldSet struct {
	Rubric  FieldSetRubric
	Items   []FieldSetPart
	Variant FieldSetVariant
}

// FieldSetRubric is the heading of a group of fields.
type FieldSetRubric interface {
	part
	fieldSetRubric()
}

// Render prints the group.
func (g FieldSet) Render() string { return g.html(nil) }

func (g FieldSet) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-fset", g.Variant.class()) + `>`)
	if g.Rubric != nil {
		s.WriteString(g.Rubric.html(nil))
	}
	for _, it := range g.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// FormPart is what a form holds.
type FormPart interface {
	part
	formPart()
}

func (Field) formPart()      {}
func (FieldSet) formPart()   {}
func (Check) formPart()      {}
func (Rubric) formPart()     {}
func (FormCols) formPart()   {}
func (FormColumn) formPart() {}

// FormCols is a row of fields that stays one row: a name split into parts.
type FormCols struct {
	Items []Field
}

func (c FormCols) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-line tc-form__cols">`)
	for _, f := range c.Items {
		s.WriteString(f.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// FormColumn is a column of a form laid out in columns.
type FormColumn struct {
	Wide  bool // 460 px; otherwise 300
	Items []FormColumnPart
}

// FormColumnPart is what a form's column holds.
type FormColumnPart interface {
	part
	formColumnPart()
}

func (c FormColumn) html(mix) string {
	var s strings.Builder
	w := "tc-form__col--300"
	if c.Wide {
		w = "tc-form__col--460"
	}
	s.WriteString(`<div class="tc-stack tc-stack--gap-12 ` + w + `">`)
	for _, it := range c.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// FormVariant is the drawn padding and flow of a form.
type FormVariant int

const (
	FormPlain          FormVariant = iota
	FormFlush                      // no side padding: inside a dialog
	FormFlushBottom                // no bottom padding
	FormFlushAll                   // no padding: inside a card
	FormGap12                      // 12/16 gaps, no padding
	FormBottom12                   // 12 px below
	FormColumnFlush                // one column, no padding
	FormColumnGap10                // one column, 10 px gaps, no padding
	FormColumnGap10Mid             // the same, centred at 340 px
)

func (v FormVariant) class() string {
	return [...]string{"", "tc-form--padl0-padr0", "tc-form--padb0", "tc-form--pad0", "tc-form--gap12-16-pad0",
		"tc-form--padb12", "tc-form--col tc-form--pad0", "tc-form--col tc-form--gap10 tc-form--pad0",
		"tc-form--col tc-form--gap10 tc-form--mid340 tc-form--pad0"}[v]
}

// Form — tc-form
//
// Роль. Форма: поля и группы полей, которые переносятся по ширине.
// Вход. Form{Items, Variant, Tight, Action, Method, Fields}.
// Состояния. plain | tight | flush | column | columns | action.
// Инварианты. Поля переносятся, а не сжимаются; отступы формы — перечисление; форма печатает фрагмент своего кадра.
// Поведение. Enter в поле формы отправляет форму.
// Эталон. кадр `s-agenda-3`, фрагмент `.tc-form`.
// Содержимое. Поля, группы полей, флажки, ряды полей, колонки.
// Чего нет. Кнопок сохранения: они в шапке экрана или в подвале формы (tc-formfoot).
type Form struct {
	Items   []FormPart
	Variant FormVariant
	Tight   bool // the form of a side panel: smaller gaps
	// Action makes the form a form element that sends its fields there;
	// Method is "post" (when empty) or "get"; Fields are its hidden fields:
	// CSRF, the record's version.
	Action string
	Method string
	Fields []Hidden
}

// Render prints the form.
func (f Form) Render() string { return f.html(nil) }

func (f Form) html(m mix) string {
	var s strings.Builder
	tag := "div"
	if f.Action != "" {
		tag = "form"
		s.WriteString(`<form` + class(m, "tc-form", is("tight", f.Tight), f.Variant.class()) + attr("action", f.Action) + ` method="` + method(f.Method) + `">`)
		for _, h := range f.Fields {
			s.WriteString(h.html(nil))
		}
	} else {
		s.WriteString(`<div` + class(m, "tc-form", is("tight", f.Tight), f.Variant.class()) + `>`)
	}
	for _, it := range f.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</` + tag + `>`)
	return s.String()
}

func method(m string) string {
	if m == "" {
		return "post"
	}
	return strings.ToLower(m)
}

// Hidden is a hidden field of a form: its name and value go with the form and
// nothing is drawn — the CSRF token, a record's version.
type Hidden struct {
	Name  string
	Value string
}

func (h Hidden) html(mix) string {
	return `<input type="hidden"` + attr("name", h.Name) + attr("value", h.Value) + `>`
}

// PostForm — tc-post
//
// Роль. Форма POST без полей на виду: её отправляет кнопка, стоящая где угодно на странице, по её имени (Button.Form).
// Вход. PostForm{ID, Action, Fields, File}.
// Состояния. default.
// Инварианты. Форму не видно: в ней только скрытые поля; форма печатает фрагмент своего кадра.
// Поведение. Отправку делает браузер — сценарий не нужен.
// Эталон. кадр `components-37`, фрагмент `.tc-post`.
// Содержимое. Скрытые поля.
// Чего нет. Видимых полей и кнопок: кнопка стоит там, где её место, — в полосе, в строке, в меню.
type PostForm struct {
	ID     string
	Action string
	Fields []Hidden
	File   string // the name of a file field: the form a dropped file is sent with
}

// Render prints the form.
func (f PostForm) Render() string { return f.html(nil) }

func (f PostForm) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<form` + class(m, "tc-post") + attr("id", f.ID) + attr("action", f.Action) + ` method="post"` + attr("enctype", pick(f.File != "", "multipart/form-data")) + `>`)
	for _, h := range f.Fields {
		s.WriteString(h.html(nil))
	}
	if f.File != "" {
		s.WriteString(`<input type="file"` + attr("name", f.File) + ` aria-label="File">`)
	}
	s.WriteString(`</form>`)
	return s.String()
}

// FormFootPart is what a form's foot holds.
type FormFootPart interface {
	part
	formFootPart()
}

func (Button) formFootPart() {}
func (Dim) formFootPart()    {}
func (Push) formFootPart()   {}
func (Span) formFootPart()   {}

// FormFootVariant is the drawn form of a form's foot.
type FormFootVariant int

const (
	FormFootPlain  FormFootVariant = iota
	FormFootFlush                  // no side padding: inside a dialog
	FormFootNoRule                 // no rule above
	FormFootWrap                   // the buttons wrap
)

func (v FormFootVariant) class() string {
	return [...]string{"", "tc-formfoot--padl0-padr0", "tc-formfoot--bordertop-0", "tc-formfoot--wrap"}[v]
}

// FormFoot — tc-formfoot
//
// Роль. Подвал формы: сохранить, отменить и что ещё сделать с формой целиком.
// Вход. FormFoot{Items, Variant}.
// Состояния. plain | flush | no-rule | wrap.
// Инварианты. Главная кнопка — первой; подвал печатает фрагмент своего кадра.
// Поведение. нет: поведение у кнопок.
// Эталон. кадр `s-contacts-5`, фрагмент `.tc-formfoot`.
// Содержимое. Кнопки, прижатая группа, приглушённое пояснение.
// Чего нет. Полей: подвал решает, а не спрашивает.
type FormFoot struct {
	Items   []FormFootPart
	Variant FormFootVariant
}

// Render prints the foot.
func (f FormFoot) Render() string { return f.html(nil) }

func (f FormFoot) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-formfoot", f.Variant.class()) + `>`)
	for _, it := range f.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}
