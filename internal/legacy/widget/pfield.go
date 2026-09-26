package widget

// PFieldWidth is the flex width hint for tc-pfield in a row.
type PFieldWidth string

const (
	PFieldWidthFull PFieldWidth = "full"
	PFieldWidthXL   PFieldWidth = "xl"
	PFieldWidthLG   PFieldWidth = "lg"
	PFieldWidthMD   PFieldWidth = "md"
	PFieldWidthSM   PFieldWidth = "sm"
	PFieldWidthXS   PFieldWidth = "xs"
)

// PField — tc-pfield
//
// Роль. Поле панели и кадра правки: подпись сверху, ввод на всю ширину ячейки.
// Вход. PField{ Name, Label, Kind, Value, Options, Pick, Width, Hint, Error, Title, ReadOnly, Disabled }.
// Состояния. normal | error | read-only | disabled.
// Инварианты. Подпись над полем, не слева; ширина задаётся Width внутри ряда.
// Поведение. Как tc-field.
// Геометрия. Столбец; Title — кегль заголовка документа.
// Содержимое. tc-in, tc-pick.
// Чего нет. Нет сетки формы 10em + 1fr.
type PField struct {
	Name     string
	Label    string
	Kind     InputKind
	Value    string
	Options  []SelectOption
	Pick     *Pick
	Width    PFieldWidth
	Hint     string
	Error    string
	Title    bool
	ReadOnly bool
	Disabled bool
}

func (f PField) Render() string {
	cls := "tc-pfield"
	if f.Width != "" {
		cls = classes(cls, "tc-pfield--"+string(f.Width))
	}
	if f.Title {
		cls = classes(cls, "tc-pfield--title")
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
	fieldID := uniqueID("tc-pfield")
	hint, err, describedBy := fieldNotes("tc-pfield", fieldID, f.Hint, f.Error)
	label := ""
	if f.Label != "" && !f.Title {
		label = `<label class="tc-pfield__label" for="` + esc(fieldID) + `">` + esc(f.Label) + `</label>`
	}
	var control string
	if f.Kind == InputPick && f.Pick != nil {
		p := *f.Pick
		if p.Name == "" {
			p.Name = f.Name
		}
		p.ID = fieldID
		p.describedBy = describedBy
		control = `<div class="tc-pfield__control">` + p.Render() + `</div>`
	} else {
		in := Input{ID: fieldID, Name: f.Name, Label: f.Label, Kind: f.Kind, Value: f.Value, Options: f.Options, ReadOnly: f.ReadOnly || f.Disabled, describedBy: describedBy}
		control = `<div class="tc-pfield__control">` + in.Render() + `</div>`
	}
	return `<div` + classAttr(cls) + `>` + label + control + hint + err + `</div>`
}

// PFieldRow — tc-pfieldrow
//
// Роль. Ряд полей панели: соседние ячейки с переносом.
// Вход. PFieldRow{ Fields []PField }.
// Состояния. нет.
// Инварианты. Каждое поле — tc-pfield, пустой ряд не рисуется.
// Поведение. нет.
// Геометрия. Flex-wrap, выравнивание по нижнему краю подписи.
// Содержимое. tc-pfield.
// Чего нет. Нет общей подписи на весь ряд.
type PFieldRow struct {
	Fields []PField
}

func (r PFieldRow) Render() string {
	if len(r.Fields) == 0 {
		return ""
	}
	var fields string
	for _, f := range r.Fields {
		fields += f.Render()
	}
	return `<div class="tc-pfieldrow">` + fields + `</div>`
}
