package widget

// Form — tc-form
//
// Роль. Набор полей с общими кнопками и общими ошибками.
// Вход. Form{ Action, Method, Sets, Errors, Foot, Busy }.
// Состояния. normal | busy.
// Инварианты. Ошибка поля у своего поля; кнопки в подвале; повторная отправка невозможна при busy.
// Поведение. Отправка, блокировка, фокус на первое поле с ошибкой.
// Геометрия. Выравнивание по строке подписи.
// Содержимое. tc-fset, tc-field, tc-formfoot.
// Чего нет. Нет проверки значений.
type Form struct {
	Action string
	Method string
	Sets   []FieldSet
	Errors []FieldError
	Foot   FormFoot
	Busy   bool
}

func (f Form) Render() string {
	method := f.Method
	if method == "" {
		method = "POST"
	}
	cls := "tc-form"
	if f.Busy {
		cls = classes(cls, "is-busy")
	}
	var sets string
	for _, s := range f.Sets {
		sets += s.Render()
	}
	var errs string
	for _, e := range f.Errors {
		errs += `<div class="tc-form__error" data-tc-field="` + esc(e.Name) + `">` + esc(e.Error) + `</div>`
	}
	dis := ""
	if f.Busy {
		dis = " aria-busy=\"true\""
	}
	return `<form action="` + esc(f.Action) + `" method="` + esc(method) + `"` + classAttr(cls) + dis + `>` + errs + sets + f.Foot.Render() + `</form>`
}

// FieldSet — tc-fset
//
// Роль. Именованный блок формы.
// Вход. FieldSet{ Legend, Rows, Note }.
// Состояния. нет.
// Инварианты. Заголовок группы — рубрикатор.
// Поведение. нет.
// Геометрия. Отступ между группами больше, чем между полями.
// Содержимое. tc-field.
// Чего нет. Не сворачивается.
type FieldSet struct {
	Legend string
	Rows   []Field
	Note   string
}

func (fs FieldSet) Render() string {
	var rows string
	for _, r := range fs.Rows {
		rows += r.Render()
	}
	note := ""
	if fs.Note != "" {
		note = `<p class="tc-fset__note">` + esc(fs.Note) + `</p>`
	}
	legend := ""
	if fs.Legend != "" {
		legend = `<legend class="tc-fset__legend">` + esc(fs.Legend) + `</legend>`
	}
	return `<fieldset class="tc-fset">` + legend + rows + note + `</fieldset>`
}

// FieldWidth is the width step for tc-field, matching .w-*.
type FieldWidth string

const (
	FieldWidthFull FieldWidth = "full"
	FieldWidthXL   FieldWidth = "xl"
	FieldWidthLG   FieldWidth = "lg"
	FieldWidthMD   FieldWidth = "md"
	FieldWidthSM   FieldWidth = "sm"
	FieldWidthXS   FieldWidth = "xs"
)

// Field — tc-field
//
// Роль. Подпись, ввод, подсказка, ошибка и кнопка, относящаяся к полю.
// Вход. Field{ Name, Label, Kind, Value, Options, Hint, Error, Extra, Width, OwnLine, ReadOnly, Required, Disabled, Check }.
// Состояния. normal | error | read-only | disabled.
// Инварианты. Подпись есть или поле явно bare; ширина через FieldWidth.
// Поведение. Показ ошибки, связь подсказки с полем.
// Геометрия. Подпись сверху; spacer 16 px у check/btn/seg, если у соседа есть label.
// Содержимое. tc-in, tc-check, tc-btn.
// Чего нет. Нет плавающей подписи внутри поля. Нет сетки «подпись слева».
type Field struct {
	Name     string
	Label    string
	Kind     InputKind
	Value    string
	Options  []SelectOption
	Hint     string
	Error    string
	Extra    Action
	Width    FieldWidth
	OwnLine  bool
	ReadOnly bool
	Required bool
	Disabled bool
	Check    *Check
}

func (f Field) Render() string {
	cls := "tc-field"
	if f.Width != "" {
		cls = classes(cls, "tc-field--"+string(f.Width))
	}
	if f.OwnLine {
		cls = classes(cls, "is-own-line")
	}
	if f.Error != "" {
		cls = classes(cls, "is-error")
	}
	if f.ReadOnly {
		cls = classes(cls, "is-read-only")
	}
	if f.Disabled {
		cls = classes(cls, "is-disabled")
	}
	req := ""
	if f.Required {
		req = `<span class="tc-field__req" aria-hidden="true">*</span>`
	}
	hint := ""
	if f.Hint != "" {
		hint = `<span class="tc-field__hint" id="` + esc(f.Name) + `-hint">` + esc(f.Hint) + `</span>`
	}
	err := ""
	if f.Error != "" {
		err = `<span class="tc-field__error" id="` + esc(f.Name) + `-error" role="alert">` + esc(f.Error) + `</span>`
	}
	extra := ""
	if f.Extra.Label != "" {
		extra = `<div class="tc-field__extra">` + f.Extra.Render() + `</div>`
	}
	if f.Label == "" {
		cls = classes(cls, "tc-field--bare")
	}
	fieldID := uniqueID("tc-field")
	if f.Check != nil {
		chk := *f.Check
		chk.ID = fieldID
		if chk.Name == "" {
			chk.Name = f.Name
		}
		if chk.Label == "" {
			chk.Label = f.Label
		}
		control := chk.Render()
		if extra != "" {
			control = `<div class="tc-field__control">` + control + extra + `</div>`
		}
		return `<div` + classAttr(cls) + `>` + control + hint + err + `</div>`
	}
	input := Input{ID: fieldID, Name: f.Name, Kind: f.Kind, Value: f.Value, Options: f.Options, ReadOnly: f.ReadOnly || f.Disabled}.Render()
	label := ""
	if f.Label != "" {
		label = `<label class="tc-field__label" for="` + esc(fieldID) + `">` + esc(f.Label) + req + `</label>`
	}
	return `<div` + classAttr(cls) + `>` + label + `<div class="tc-field__control">` + input + extra + `</div>` + hint + err + `</div>`
}

// FormFoot — tc-formfoot
//
// Роль. Кнопки формы в объявленном порядке.
// Вход. FormFoot{ Primary, Rest, Danger }.
// Состояния. нет.
// Инварианты. Первичное действие ровно одно; опасное отделено.
// Поведение. нет.
// Геометрия. Одна строка; в узкой раскладке — столбец.
// Содержимое. tc-btn.
// Чего нет. Нет кнопок, ведущих на другие экраны.
type FormFoot struct {
	Primary Action
	Rest    []Action
	Danger  []Action
}

func (ff FormFoot) Render() string {
	if ff.Primary.Label == "" && ff.Primary.Icon == "" && ff.Primary.Href == "" && len(ff.Rest) == 0 && len(ff.Danger) == 0 {
		return ""
	}
	var rest, danger string
	for _, a := range ff.Rest {
		aa := a
		if aa.Size == "" {
			aa.Size = ButtonSizeForm
		}
		rest += aa.Render()
	}
	for _, a := range ff.Danger {
		da := a
		if da.Size == "" {
			da.Size = ButtonSizeForm
		}
		if da.Kind == "" {
			da.Kind = ButtonDangerSolid
		}
		danger += da.Render()
	}
	primary := ""
	if ff.Primary.Label != "" || ff.Primary.Icon != "" || ff.Primary.Href != "" {
		p := ff.Primary
		if p.Size == "" {
			p.Size = ButtonSizeForm
		}
		if p.Kind == "" {
			p.Kind = ButtonPrimary
		}
		primary = p.Render()
	}
	return `<div class="tc-formfoot"><div class="tc-formfoot__primary">` + primary + `</div><div class="tc-formfoot__rest">` + rest + `</div><div class="tc-formfoot__danger">` + danger + `</div></div>`
}

// Danger — tc-danger
//
// Роль. Блок разрушающих действий внизу страницы настроек.
// Вход. Danger{ Title, Rows }.
// Состояния. нет.
// Инварианты. Каждое действие называет последствие; блок всегда последний.
// Поведение. нет; подтверждение — tc-dialog.
// Геометрия. Отбит от остальной страницы.
// Содержимое. tc-btn с состоянием danger.
// Чего нет. Нет действий без подтверждения.
type Danger struct {
	Title string
	Rows  []DangerRow
}

func (d Danger) Render() string {
	var rows string
	for _, r := range d.Rows {
		a := r.Action
		if a.Size == "" {
			a.Size = ButtonSizeForm
		}
		if a.Kind == "" {
			a.Kind = ButtonDangerSolid
		}
		rows += `<div class="tc-danger__row"><div class="tc-danger__text"><strong>` + esc(r.Title) + `</strong><p>` + esc(r.Text) + `</p></div><div class="tc-danger__action">` + a.Render() + `</div></div>`
	}
	return `<section class="tc-danger"><h2 class="tc-danger__title">` + esc(d.Title) + `</h2>` + rows + `</section>`
}
