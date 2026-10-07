package widget

import "strings"

// InputKind is what the field takes; it is the input's type.
type InputKind int

const (
	InputText InputKind = iota
	InputDate
	InputTime
	InputSearch
	InputPassword
)

func (k InputKind) attr() string {
	return [...]string{"text", "date", "time", "search", "password"}[k]
}

// InputVariant is the drawn size of a field where it is not its column's.
type InputVariant int

const (
	InputAuto      InputVariant = iota
	InputChip                   // a number inside a sentence or a menu row
	InputFill                   // takes the rest of its row
	Input11_5                   // 11.5 px: an address
	Input11_5Fill               // 11.5 px, fills the row
	Input11                     // 11 px
	Input11Fill                 // 11 px, fills the row
	InputLineFill               // one line of a long value, fills the row
	InputKeyFill                // a key or a password, fills the row
	InputKeySmFill              // a smaller key, fills the row
	InputDusty                  // a value that stays as it is
	InputW52                    // 52 px: a threshold
	InputW190                   // 190 px filter
	InputH26W120                // 26 px high, 120 wide: a date in a bar
	InputH26W200                // 26 px high, 200 wide: a filter in a bar
)

func (v InputVariant) class() string {
	return [...]string{"", "tc-in--chip", "tc-in--fill", "tc-in--t11-5", "tc-in--t11-5-fill", "tc-in--t11",
		"tc-in--t11-fill", "tc-in--line tc-in--fill", "tc-in--key tc-in--fill", "tc-in--key-sm tc-in--fill",
		"tc-in--dusty", "tc-in--w52", "tc-in--w190", "tc-in--h26-w120", "tc-in--h26-w200"}[v]
}

// Input — tc-in
//
// Роль. Поле ввода одной строки: текст, дата, время, поиск, пароль.
// Вход. Input{Kind, Name, Value, Placeholder, Label, Mono, Focus, Error, Numeric, Disabled, Variant, Place, ReadOnly, Required, ID, Picked}.
// Состояния. value | placeholder | mono | focus | error | disabled | variant.
// Инварианты. Плейсхолдер выглядит плейсхолдером всегда: is-ph ставит виджет, когда значения нет; род поля — перечисление; у поля есть доступное имя; поле печатает фрагмент своего кадра.
// Поведение. Ввод — браузера; нативный выбор даты и времени не заменяется.
// Эталон. кадр `components-10`, фрагмент `.tc-in`.
// Содержимое. нет.
// Чего нет. Своей подписи: подпись — у поля формы (tc-field), вокруг.
type Input struct {
	Kind        InputKind
	Name        string
	Value       string
	Placeholder string
	Label       string // the accessible name when no field label wraps it
	Mono        bool
	Focus       bool // the drawn keyboard focus of the mockup's states
	Error       bool
	Numeric     bool // a phone keyboard shows digits
	Disabled    bool // switched off by a choice next to it (a password the server does not ask)
	ReadOnly    bool // shown and sent, not edited
	Required    bool // the form is not sent without it
	Variant     InputVariant
	Place       Place
	// ID is what a file field beside it names (FileField.Shows): the line
	// shows the picked file's name, read-only; without a script it is not
	// shown — the file input shows the name itself (Carrel's task 133).
	ID string
	// Picked marks such a line: what was picked beside it.
	Picked bool
}

// Render prints the input.
func (i Input) Render() string { return i.html(nil) }

func (i Input) html(m mix) string {
	return `<input` + class(m, "tc-in", is("mono", i.Mono), is("ph", i.Value == "" && i.Placeholder != ""),
		is("focus", i.Focus), is("error", i.Error), i.Variant.class(), i.Place.class()) +
		` type="` + i.Kind.attr() + `"` + attr("name", i.Name) + attr("value", i.Value) +
		attr("placeholder", i.Placeholder) + attr("inputmode", pick(i.Numeric, "numeric")) + attr("aria-label", i.Label) + pick(i.Disabled, " disabled") + flag("readonly", i.ReadOnly || i.Picked) + flag("required", i.Required) +
		attr("id", i.ID) + flag("data-tc-filename", i.Picked) + `>`
}

// PickVariant is the drawn size of a pick where it is not its column's.
type PickVariant int

const (
	PickAuto    PickVariant = iota
	PickT12W200             // 12 px, 200 wide: where to write, in a sheet's head
	PickT12W240             // 12 px, 240 wide
	PickT12Fill             // 12 px, fills its row
	PickH26W190             // 26 px high: a filter of a list
)

func (v PickVariant) class() string {
	return [...]string{"", "tc-in--t12-w200", "tc-in--t12-w240", "tc-in--t12-fill", "tc-in--h26-w190"}[v]
}

// Pick — tc-in
//
// Роль. Поле со списком: значение выбирают из списка, а не набирают.
// Вход. Pick{Value, Mark, Aside, Unset, Label, Focus, Variant, Place, Name, Options}.
// Состояния. value | unset | aside | focus | variant | select.
// Инварианты. Стрелку рисует система (is-pick), иначе список неотличим от строки ввода; пустой выбор выглядит плейсхолдером; выбор печатает фрагмент своего кадра.
// Поведение. Нажатие открывает список под полем; выбор закрывает его, фокус возвращается на поле.
// Эталон. кадр `components-10`, фрагмент `.tc-in.is-pick`.
// Содержимое. Приглушённое пояснение после значения.
// Чего нет. Набора с клавиатуры: чтобы искать, есть фильтр списка.
type Pick struct {
	// Name and Options make the pick a list of the browser (select): it
	// sends the chosen option's value and works without scripts. Value and
	// Aside are then not drawn — the chosen option is.
	Name    string
	Options []PickOption
	Value   string
	Mark    *Mark // the colour of the chosen collection: a notebook (Carrel's task 103)
	Aside   *Dim  // "as sent" after a time zone
	Unset   bool  // nothing is chosen: the value is a placeholder
	Label   string
	Focus   bool
	Variant PickVariant
	Place   Place
}

// Render prints the pick.
func (p Pick) Render() string { return p.html(nil) }

// PickOption is one option of a pick's list.
type PickOption struct {
	Value       string
	Label       string
	On          bool // chosen
	Unavailable bool // shown, not chosen: a kind the server does not take
}

func (p Pick) html(m mix) string {
	if p.Name != "" {
		var s strings.Builder
		s.WriteString(`<select` + class(m, "tc-in", is("ph", p.Unset), is("focus", p.Focus), "is-pick", p.Variant.class(), p.Place.class()) + attr("name", p.Name) + attr("aria-label", p.Label) + `>`)
		for _, o := range p.Options {
			s.WriteString(`<option` + attr("value", o.Value) + flag("selected", o.On) + flag("disabled", o.Unavailable) + `>` + esc(o.Label) + `</option>`)
		}
		s.WriteString(`</select>`)
		return s.String()
	}
	s := `<button` + class(m, "tc-in", is("ph", p.Unset), is("focus", p.Focus), "is-pick", p.Variant.class(), p.Place.class()) +
		` type="button"` + attr("aria-label", p.Label) + `>`
	if p.Mark != nil {
		s += p.Mark.html(nil)
	}
	s += esc(p.Value)
	if p.Aside != nil {
		s += " " + p.Aside.html(nil)
	}
	return s + `</button>`
}

// Query — tc-q
//
// Роль. Строка запроса внутри поля поиска или фильтра.
// Вход. Query{Value, Placeholder, Label, Name}.
// Состояния. нет.
// Инварианты. Тип поля — search; у строки есть доступное имя; строка печатает фрагмент своего кадра.
// Поведение. Ввод — браузера; Esc очищает, если поле не пусто.
// Эталон. кадр `shell-1`, фрагмент `.tc-q`.
// Содержимое. нет.
// Чего нет. Рамки и значка: их рисует поле вокруг.
type Query struct {
	Value       string
	Placeholder string // what the empty line says: «Search everything…»
	Label       string
	Name        string // what the search sends it as: q
}

// Render prints the query line.
func (q Query) Render() string { return q.html(nil) }

func (q Query) html(m mix) string {
	return `<input` + class(m, "tc-q") + ` type="search"` + attr("name", q.Name) + attr("value", q.Value) + attr("placeholder", q.Placeholder) + attr("aria-label", q.Label) + `>`
}

// FilterVariant is the drawn size of a filter.
type FilterVariant int

const (
	FilterH26W220 FilterVariant = iota // a list's filter in its bar
	FilterH26W190                      // the same, narrower
	FilterT12W200                      // 12 px, in a sheet's head
	FilterH26W200                      // a bar's filter beside a wider segment (notes, users)
)

func (v FilterVariant) class() string {
	return [...]string{"tc-in--h26-w220", "tc-in--h26-w190", "tc-in--t12-w200", "tc-in--h26-w200"}[v]
}

// Filter — tc-in
//
// Роль. Фильтр списка: значок поиска и строка запроса в одной рамке.
// Вход. Filter{Query, Waiting, Quiet, Variant, Place, Action, Hidden}.
// Состояния. waiting | quiet.
// Инварианты. Значок — первый узел поля, строка занимает остаток; фильтр печатает фрагмент своего кадра.
// Поведение. Сужает список по мере набора; Esc очищает.
// Эталон. кадр `s-contacts-1`, фрагмент `label.tc-in`.
// Содержимое. Строка запроса.
// Чего нет. Подсказок по мере набора: это поиск в шапке, а не фильтр.
type Filter struct {
	Query   Query
	Waiting bool // nothing typed: the line shows what it filters, in the placeholder's colour
	Quiet   bool // a smaller, dusty icon: a filter inside a sheet
	Variant FilterVariant
	Place   Place
	Action  string   // the filter is a form that sends its query there (GET)
	Hidden  []Hidden // what the form sends with the query: the list's other choices (sort, tag)
}

// Render prints the filter.
func (f Filter) Render() string { return f.html(nil) }

func (f Filter) html(m mix) string {
	icon := `<svg width="13" height="13" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.5" fill="none" class="tc-ico--mr6">`
	if f.Quiet {
		icon = `<svg width="12" height="12" viewBox="0 0 24 24" stroke-width="1.5" fill="none" class="tc-ico--dusty">`
	}
	inner := icon + Icon("search").use() + `</svg>` + f.Query.html(nil)
	if f.Action != "" {
		for _, h := range f.Hidden {
			inner += h.html(nil)
		}
		return `<form` + class(m, "tc-in", is("ph", f.Waiting), f.Variant.class(), f.Place.class()) + attr("action", f.Action) + ` method="get">` + inner + `</form>`
	}
	return `<label` + class(m, "tc-in", is("ph", f.Waiting), f.Variant.class(), f.Place.class()) + `>` + inner + `</label>`
}

// TagInputPart is what a tag field holds: the chosen tags and the text being typed.
type TagInputPart interface {
	part
	tagInputPart()
}

func (Tag) tagInputPart()  {}
func (Dim) tagInputPart()  {}
func (Span) tagInputPart() {}

// TagInput — tc-in
//
// Роль. Поле, в котором выбранное стоит тегами, а новое набирают тут же: связанные записи, получатели письма.
// Вход. TagInput{Parts, Wrap, Focus, Columns, Name, Add}.
// Состояния. tags | wrap | focus | columns | add.
// Инварианты. Выбранное — теги, их снимают крестиком, а не правят; поле печатает фрагмент своего кадра.
// Поведение. С Name каждый тег отправляет своё значение, снятый крестиком — нет, и без сценария тоже. С живым слоем (tc-live.js) набор открывает подсказки под полем (Add.Suggest: сервер отвечает пунктами меню); выбор добавляет тег; крестик и Backspace в пустой строке снимают тег; Enter не отправляет форму; Esc закрывает подсказки.
// Эталон. кадр `s-tasks-3`, фрагмент `span.tc-in.tc-in--tags`.
// Содержимое. Теги и набранный текст.
// Чего нет. Ввода UID руками вне подсказок: его пункт стоит в конце списка подсказок.
type TagInput struct {
	Parts   []TagInputPart
	Wrap    bool // tags wrap onto more lines: the recipients of a letter
	Focus   bool
	Columns bool // tags stand in columns across the width: the attendees of an event
	// Name makes the field a part of its form: each tag with a Value sends
	// it as Name; a tag the cross unticks is not sent.
	Name string
	Add  *TagAdd // the line the next tag is typed in
}

// TagAdd is the line of a TagInput where the next tag is typed.
type TagAdd struct {
	// Name is what the typed words are sent as: without a script the server
	// finds them itself (a name in the contacts, or an address).
	Name        string
	Placeholder string // «Name or address»
	Label       string
	Value       string // what is typed: the frames of suggestions
	// Suggest is where the suggestions come from: the library's live layer
	// (tc-live.js, L-1) asks Suggest?q=… as one types and shows the menu items
	// the server answers;
	// an item's Adds is the tag it adds, with the field's Name (MenuItems
	// prints the answer).
	Suggest string
	// Options are the choices of a page without a script (Carrel's task 96):
	// a datalist of the line, the browser's own list as one types — Value is
	// what is sent (a record's UID, an address), Label what is shown beside
	// it. The live layer takes the list away where it asks Suggest.
	Options []ListOption
}

// ListOption is a choice of a tag field's line on a page without a script.
type ListOption struct {
	Value string // what is sent
	Label string // what is shown beside it
}

// Render prints the tag field.
func (t TagInput) Render() string { return t.html(nil) }

func (t TagInput) html(m mix) string {
	var s strings.Builder
	variant := "tc-in--tags"
	if t.Wrap {
		variant = "tc-in--wrap-gap5"
	}
	sug := ""
	if t.Add != nil {
		sug = t.Add.Suggest
	}
	s.WriteString(`<span` + class(m, "tc-in", is("focus", t.Focus), variant, pick(t.Columns, "tc-in--cols")) + attr("data-tc-suggest", sug) + `>`)
	for _, p := range t.Parts {
		if g, ok := p.(Tag); ok {
			s.WriteString(" " + g.in(nil, t.Name))
			continue
		}
		s.WriteString(" " + p.html(nil))
	}
	if a := t.Add; a != nil {
		list := ""
		if len(a.Options) > 0 {
			list = "tc-list-" + a.Name
		}
		s.WriteString(` <input class="tc-in__add" type="text"` + attr("name", a.Name) + attr("value", a.Value) + attr("placeholder", a.Placeholder) + attr("aria-label", a.Label) + attr("list", list) + ` autocomplete="off">`)
		if list != "" {
			s.WriteString(`<datalist` + attr("id", list) + `>`)
			for _, o := range a.Options {
				s.WriteString(`<option` + attr("value", o.Value) + attr("label", o.Label) + `></option>`)
			}
			s.WriteString(`</datalist>`)
		}
	}
	s.WriteString(` </span>`)
	return s.String()
}

// AreaVariant is the drawn form of a text area.
type AreaVariant int

const (
	AreaPlain  AreaVariant = iota // a note field of a form
	AreaTitle                     // the title of a document, one line that grows
	AreaBody                      // the body of a document
	AreaBodySm                    // the body in a sheet
	AreaLetter                    // the words of a letter
	AreaPaste                     // text to copy as it is
	AreaTall                      // a taller body, 13.5 px
)

func (v AreaVariant) class() string {
	return [...]string{"is-area", "tc-in--title", "is-area tc-in--body", "is-area tc-in--body-sm", "is-area tc-in--note",
		"is-area tc-in--paste", "is-area tc-in--minh104-t13-5"}[v]
}

// Area — tc-in
//
// Роль. Многострочное поле: примечание к записи, текст документа, заголовок, который растёт.
// Вход. Area{Name, Text, Variant, Mono, Focus, Label}.
// Состояния. plain | title | body | letter | paste | tall | focus.
// Инварианты. Текст — как есть, с переводами строк; поле растёт по содержимому, а не прокручивается внутри; поле печатает фрагмент своего кадра.
// Поведение. Ввод — браузера.
// Эталон. кадр `s-contacts-3`, фрагмент `textarea.tc-in`.
// Содержимое. нет.
// Чего нет. Разметки: кнопки полосы правки вставляют символы Markdown в текст.
type Area struct {
	Name    string
	Text    string
	Variant AreaVariant
	Mono    bool
	Focus   bool
	Label   string
}

// Render prints the text area.
func (a Area) Render() string { return a.html(nil) }

func (a Area) html(m mix) string {
	return `<textarea` + class(m, "tc-in", a.Variant.class(), is("mono", a.Mono), is("focus", a.Focus)) +
		attr("name", a.Name) + attr("aria-label", a.Label) + `>` + esc(a.Text) + `</textarea>`
}
