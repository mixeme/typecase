package widget

import "strings"

// MenuTone is the colour of a menu item's words.
type MenuTone int

const (
	MenuToneInk    MenuTone = iota
	MenuToneDanger          // destroys something
	MenuToneDusty           // a way out: cancel, clear
	MenuToneMuted           // a quieter item
	MenuToneAccent          // resets a list to its defaults
	MenuToneAside           // a small note row: what an item above means
)

func (t MenuTone) class() string {
	return [...]string{"", "tc-menu__item--danger", "tc-menu__item--dusty", "tc-menu__item--muted", "tc-menu__item--accent", "tc-menu__item--note"}[t]
}

// MenuTrail is what stands at the end of a menu item.
type MenuTrail interface {
	part
	menuTrail()
}

func (Mono) menuTrail()  {}
func (Dim) menuTrail()   {}
func (Badge) menuTrail() {}
func (Kbd) menuTrail()   {}
func (Tag) menuTrail()   {}

// MenuWhen is when an item of a search's suggestions (Top.Suggest) is
// shown, with the library's script, by what is typed in the search line
// (Carrel's task 92). Without a script all are shown as printed.
type MenuWhen int

const (
	WhenAlways MenuWhen = iota // a recent search
	WhenEmpty                  // only while the line is empty: «Type to search everything»
	WhenTyped                  // only while something is typed: «Show all results for «{q}»», the words in braces being what is typed
	WhenFound                  // only while its words hold what is typed, set bold: a jump to what the page already holds
)

func (w MenuWhen) attr() string {
	return attr("data-tc-when", [...]string{"", "empty", "typed", "found"}[w])
}

// MenuItem is a row of a menu that does one thing.
type MenuItem struct {
	Icon   Icon
	Faint  bool // the icon is a handle, drawn faint
	Mark   *Mark
	Tick   *TickBox // a tick box: the item shows or hides something
	Label  string
	Found  []Inline // the label with what was found in bold, instead of Label
	Note   *Span    // a warning after the label
	Detail *Mono    // a monospaced detail after the label: an address
	Trail  MenuTrail
	On     bool // the chosen or current item
	Tone   MenuTone
	// Reason, when set, switches the item off and says why: someone the
	// list names but who cannot be chosen now.
	Reason string
	// Href makes the item a link: a sheet or a page opens at the address.
	Href string
	// Swap is Button.Swap for an item that is a link: with the live layer
	// the elements it names are put in place from the answer — the
	// connection of a sheet's tree chosen in its head (Carrel's task 205).
	Swap string
	// Name makes the item a checkbox of the menu's form (Menu.Action): ticked
	// (Tick), it sends Value; the items' order is the order they are sent in.
	// Fixed keeps it ticked and off: a column that is always shown. Without a
	// Tick, a Form or an Href, Name makes it a radio button of the form the
	// menu stands in, or of the form Of names: one of the items, the chosen
	// one (On) sends its Value, empty too («All actions» of a filter) — a
	// time zone (Carrel's task 73).
	Name  string
	Value string
	// Of is the ID of the form a radio item (Name) belongs to without
	// standing in it: the collection chosen in an edit's head, outside the
	// edit's form (Carrel's task 203).
	Of    string
	Fixed bool
	Drag  bool // the drawn state of an item being dragged (the mockup's states)
	// Tertiary is the twin of a bar's tertiary item in the menu of «⋯»:
	// shown on a phone only, where the item itself is under «⋯».
	Tertiary bool
	// Form is the ID of a PostForm the item submits (Sign out): the item is a
	// submit button then, sending Name and Value when they are set.
	Form string
	// To makes the item a submit button of the form it stands in, sending
	// that form to its own address (formaction), as Button.To does: Upload
	// on a new record keeps what is typed (Carrel's task 122).
	To string
	// File makes the item a label that picks a file for a PostForm (Carrel's
	// task 109): «From this computer» opens the choice without a script.
	File *FileField
	// Row is the row a file picked by the item becomes in its edit's list
	// (File.List) — the list's own row, «{name}» and «{size}» in its words; a
	// button of the row with Removes takes it away before Save.
	Row PickedRow
	// NoScript shows the item only on a page without a script: «Upload».
	NoScript bool
	// ScriptOnly shows the item only on a page with a script (Carrel's task
	// 90): what it does needs one.
	ScriptOnly bool
	// Adds is the tag the item adds to its TagInput: a suggestion of a person
	// (Carrel's task 77). It waits in a template; the library's script puts it
	// before the typing line.
	Adds *Tag
	// Key is the key that presses the item, with the library's script, the
	// menu open or closed (Carrel's task 87): «?» on «Keyboard shortcuts». It
	// does nothing while the cursor is in a field. An item that destroys
	// (MenuToneDanger) takes no key, nor does an item that is a tick, a
	// radio button or a file's label.
	Key string
	// When is when the item of a search's suggestions is shown (Carrel's
	// task 92).
	When MenuWhen
}

// MenuItems prints menu items without their menu: what a server answers a
// TagInput's suggestions with (TagAdd.Suggest).
func MenuItems(items ...MenuPart) string { return joined(items) }

func (it MenuItem) html(mix) string {
	var s strings.Builder
	tag := "button"
	cls := class(nil, is("on", it.On), "tc-menu__item", it.Tone.class(), rank3(it.Tertiary), pick(it.ScriptOnly, "tc-js"), pick(it.NoScript, "tc-nojs"))
	if it.File != nil {
		s.WriteString(`<label` + cls + ` role="menuitem">`)
		if it.Icon != "" {
			s.WriteString(`<svg` + class(nil, pick(it.Faint, "tc-ico--op-35")) + `>` + it.Icon.use() + `</svg>`)
		}
		s.WriteString(esc(it.Label) + it.File.input())
		if it.Row != nil {
			s.WriteString(`<template data-tc-row>` + it.Row.html(nil) + `</template>`)
		}
		s.WriteString(`</label>`)
		return s.String()
	}
	checked := attr("aria-checked", pick(it.Tick != nil && it.Tick.Ticked != TickedOff, "true"))
	key := attr("data-tc-key", pick(it.Tone != MenuToneDanger, it.Key)) + it.When.attr()
	if it.Name != "" && it.Tick != nil {
		cls = class(nil, is("on", it.On), is("drag", it.Drag), "tc-menu__item", it.Tone.class(), rank3(it.Tertiary))
		s.WriteString(`<label` + cls + ` role="menuitem"><input class="tc-check__box" type="checkbox"` + attr("name", it.Name) + attr("value", it.Value) +
			flag("checked", it.Tick != nil && it.Tick.Ticked != TickedOff) + flag("disabled", it.Fixed) + `>`)
		if it.Icon != "" {
			s.WriteString(`<svg` + class(nil, pick(it.Faint, "tc-ico--op-35")) + `>` + it.Icon.use() + `</svg>`)
		}
		s.WriteString(`<span` + class(nil, "tc-tickrow__tick", [...]string{"", "is-part", "is-off"}[it.Tick.Ticked]) + `></span>` + esc(it.Label))
		if it.Trail != nil {
			s.WriteString(it.Trail.html(nil))
		}
		s.WriteString(`</label>`)
		return s.String()
	}
	if it.Name != "" && it.Form == "" && it.Href == "" {
		// a collection that cannot be chosen is shown and switched off, its
		// reason in the hint: read-only (Carrel's task 103)
		s.WriteString(`<label` + cls + attr("title", it.Reason) + `><input class="tc-seg__box" type="radio"` + attr("form", it.Of) + ` role="menuitemradio"` + attr("name", it.Name) + ` value="` + esc(it.Value) + `"` +
			flag("checked", it.On) + flag("disabled", it.Reason != "") + `>`)
		if it.Icon != "" {
			s.WriteString(`<svg` + class(nil, pick(it.Faint, "tc-ico--op-35")) + `>` + it.Icon.use() + `</svg>`)
		}
		if it.Mark != nil {
			s.WriteString(it.Mark.html(nil))
		}
		s.WriteString(esc(it.Label))
		if it.Trail != nil {
			s.WriteString(it.Trail.html(nil))
		}
		s.WriteString(`</label>`)
		return s.String()
	}
	if it.Href != "" {
		tag = "a"
		if it.Reason != "" {
			s.WriteString(`<a` + cls + ` role="menuitem"` + checked + attr("title", it.Reason) + ` aria-disabled="true">`)
		} else {
			s.WriteString(`<a` + cls + attr("href", it.Href) + attr("data-tc-swap", it.Swap) + ` role="menuitem"` + checked + key + `>`)
		}
	} else if it.To != "" {
		s.WriteString(`<button` + cls + ` type="submit"` + attr("formaction", it.To) + attr("name", it.Name) + attr("value", it.Value) + ` role="menuitem"` + key + attr("title", it.Reason) + flag("disabled", it.Reason != "") + `>`)
	} else if it.Form != "" {
		s.WriteString(`<button` + cls + ` type="submit"` + attr("form", it.Form) + attr("name", it.Name) + attr("value", it.Value) + ` role="menuitem"` + key + attr("title", it.Reason) + flag("disabled", it.Reason != "") + `>`)
	} else {
		s.WriteString(`<button` + cls + ` type="button" role="menuitem"` + checked + key + attr("title", it.Reason) + flag("disabled", it.Reason != "") + `>`)
	}
	if it.Icon != "" {
		s.WriteString(`<svg` + class(nil, pick(it.Faint, "tc-ico--op-35")) + `>` + it.Icon.use() + `</svg>`)
	}
	if it.Mark != nil {
		s.WriteString(it.Mark.html(nil))
	}
	if it.Tick != nil {
		s.WriteString(`<span` + class(nil, "tc-tickrow__tick", [...]string{"", "is-part", "is-off"}[it.Tick.Ticked]) + `></span>`)
	}
	if it.Found != nil {
		s.WriteString(`<span class="tc-menu__label">` + inlines(it.Found) + `</span>`)
	} else if it.When == WhenFound {
		// the words the script sets what is typed bold in
		s.WriteString(`<span class="tc-menu__label">` + esc(it.Label) + `</span>`)
	} else {
		s.WriteString(esc(it.Label))
	}
	if it.Note != nil {
		s.WriteString(" " + it.Note.html(nil))
	}
	if it.Detail != nil {
		s.WriteString(it.Detail.html(nil))
	}
	if it.Trail != nil {
		s.WriteString(it.Trail.html(nil))
	}
	if it.Adds != nil {
		s.WriteString(`<template>` + it.Adds.html(nil) + `</template>`)
	}
	s.WriteString(`</` + tag + `>`)
	return s.String()
}

// MenuSep is a rule between groups of a menu.
type MenuSep struct {
	Tertiary bool // it parts the twins of a bar's tertiary items: a phone only
}

func (p MenuSep) html(mix) string {
	return `<span` + class(nil, "tc-menu__sep", rank3(p.Tertiary)) + `></span>`
}

// MenuHeadingForm is how a menu's heading row is drawn.
type MenuHeadingForm int

const (
	MenuHeadingPlain MenuHeadingForm = iota
	MenuHeadingRoomy                 // 5/7 px around: over a list of columns
	MenuHeadingSplit                 // words and a value at the two edges: the account over the user menu
)

// MenuHeading is a menu's heading row: a rubric the menu sets as an item.
type MenuHeading struct {
	Rubric Rubric
	Form   MenuHeadingForm
}

func (h MenuHeading) html(mix) string {
	return h.Rubric.html(mix{"tc-menu__item", [...]string{"", "tc-menu__item--pad5-8-7", "tc-menu__item--split"}[h.Form]})
}

// MenuField is a row of a menu that sets a value: a signal's weight.
type MenuField struct {
	Label   string
	Badge   *Badge
	Default Dim // the value it has by default, pushed to the end
	Input   Input
}

func (f MenuField) html(mix) string {
	s := `<div class="tc-menu__item">` + esc(f.Label)
	if f.Badge != nil {
		s += " " + f.Badge.html(nil)
	}
	return s + f.Default.html(nil) + f.Input.html(nil) + `</div>`
}

// MenuColumn is a row of a menu with a control under its name: the theme.
type MenuColumn struct {
	Label   string
	Control Seg
}

func (c MenuColumn) html(mix) string {
	return `<div class="tc-menu__item tc-menu__item--col"><span class="tc-line tc-line--gap-9 tc-menu__item-indent">` + esc(c.Label) + `</span>` +
		c.Control.html(nil) + `</div>`
}

// MenuRefusal is a row of a menu of refusals (Carrel's decision of 02.10): a
// source that did not answer, what was asked and what it answered, in full,
// and its Retry. The rail only marks the source; the reason is here.
type MenuRefusal struct {
	Mark   Mark
	Source string // «radicale · vps · Дежурства»
	Answer string // «REPORT /dav/duty/ → 503 Service Unavailable»
	Retry  Button
}

func (r MenuRefusal) html(mix) string {
	m := r.Mark
	m.Size = MarkH11
	return `<div class="tc-menu__refusal">` + m.html(nil) + `<span class="tc-menu__what"><span class="tc-menu__source">` + esc(r.Source) +
		`</span><span class="tc-menu__answer">` + esc(r.Answer) + `</span></span>` + r.Retry.html(nil) + `</div>`
}

// MenuFoot is the last row of a menu that edits: save, go back.
type MenuFoot struct {
	Items []Button
}

func (f MenuFoot) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-menu__foot">`)
	for _, b := range f.Items {
		s.WriteString(b.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// MenuPart is what a menu holds.
type MenuPart interface {
	part
	menuPart()
}

func (MenuItem) menuPart()    {}
func (MenuSep) menuPart()     {}
func (MenuHeading) menuPart() {}
func (MenuField) menuPart()   {}
func (MenuColumn) menuPart()  {}
func (MenuFoot) menuPart()    {}
func (MenuRefusal) menuPart() {}
func (Rubric) menuPart()      {}
func (Input) menuPart()       {}
func (Range) menuPart()       {}
func (CheckForm) menuPart()   {}

// MenuWidth is the width of a menu.
type MenuWidth int

const (
	MenuWidthAuto MenuWidth = iota // 236 px
	MenuWidth212
	MenuWidth290
	MenuWidth420
	MenuWidthFull // its place's width: under a search, in a field
)

func (w MenuWidth) class() string {
	return [...]string{"", "tc-menu--w212", "tc-menu--w290", "tc-menu--w420", "tc-menu--full"}[w]
}

// MenuPop is where a menu stands.
type MenuPop int

const (
	MenuInPlace  MenuPop = iota // drawn where it is, as in the frames of menus
	MenuUnder                   // pops up under its button
	MenuUnderEnd                // pops up under its button, aligned to its end
)

// Menu — tc-menu
//
// Роль. Меню: действия с записью, выбор значения, куда записать, подсказки поиска, настройка колонок.
// Вход. Menu{Items, Width, Pop, Fields, Inline, Action, Hidden, Target}.
// Состояния. on | icons | marks | danger | heading | pop | fields | found | form | radio.
// Инварианты. Пункт — кнопка с ролью пункта меню; ширина и место — перечисления; меню печатает фрагмент своего кадра.
// Поведение. Стрелки ходят по пунктам, Enter выбирает, Esc закрывает меню и возвращает фокус кнопке, которая его открыла; строка поиска оставляет пункты, где есть набранное (без сценария её нет).
// Эталон. кадр `components-1`, фрагмент `.tc-menu`.
// Содержимое. Пункты (значок, метка, галка, название, приглушённый хвост, клавиша; радио формы вокруг), строка поиска, разделители, заголовки, поле, строка с переключателем, диапазон, отметки-форма (CheckForm), отказы источников (MenuRefusal), подвал.
// Чего нет. Вложенных меню: второй уровень — лист или страница.
type Menu struct {
	Items  []MenuPart
	Width  MenuWidth
	Pop    MenuPop
	Fields bool // rows that set values: taller rows
	Inline bool // a span: a menu inside a line of a bar
	// Action makes the menu a form that sends its checkbox items there, with
	// its hidden Fields: the choice of columns.
	Action string
	Hidden []Hidden
	// Target is the name of the zone the links and forms inside ask their
	// address for with the live layer, instead of loading the page (Carrel's
	// task 84); without the layer they are plain links and forms.
	Target string
}

// Render prints the menu.
func (u Menu) Render() string { return u.html(nil) }

func (u Menu) html(m mix) string { return u.shown(m, true) }

// shown prints the menu open, or closed until its button opens it.
func (u Menu) shown(m mix, open bool) string {
	tag := "div"
	if u.Inline {
		tag = "span"
	}
	form := ""
	if u.Action != "" {
		tag = "form"
		form = attr("action", u.Action) + ` method="post"`
	}
	var s strings.Builder
	s.WriteString(`<` + tag + class(m, "tc-menu", is("pop", u.Pop != MenuInPlace), is("right", u.Pop == MenuUnderEnd), u.Width.class(), pick(u.Fields, "tc-menu--fields")) + form + attr("data-tc-target", u.Target) + ` role="menu"` + flag("hidden", !open) + `>`)
	for _, h := range u.Hidden {
		s.WriteString(h.html(nil))
	}
	for _, it := range u.Items {
		if f, ok := it.(CheckForm); ok {
			s.WriteString(f.inMenu())
			continue
		}
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</` + tag + `>`)
	return s.String()
}

// DropTrigger is what opens a drop.
type DropTrigger interface {
	part
	dropTrigger()
}

func (Button) dropTrigger()   {}
func (Pick) dropTrigger()     {}
func (TagInput) dropTrigger() {}

// Drop — tc-drop
//
// Роль. Кнопка или выбор вместе с меню, которое они открывают.
// Вход. Drop{Trigger, Menu, More, Row, Open, Place}.
// Состояния. more | pick.
// Инварианты. Меню стоит сразу за тем, что его открывает; «⋯» бара — кнопка бара и тогда, когда она в выпадашке; выпадашка печатает фрагмент своего кадра.
// Поведение. Нажатие открывает меню под кнопкой; Esc закрывает и возвращает фокус; у выбора с меню из радио (пояс) выбранный пункт становится значением выбора.
// Эталон. кадр `components-19`, фрагмент `.tc-drop`.
// Содержимое. Кнопка или поле выбора, меню.
// Чего нет. Своего вида: вид у кнопки и меню; закрытая открывается и без сценария (details).
type Drop struct {
	Trigger DropTrigger
	Menu    Menu
	More    bool // the trigger is its bar's «⋯»
	Row     bool // the trigger is a table row's «⋯» (TableCell.Drop, Carrel's task 117)
	Open    bool // drawn open, as the frames draw it; closed, the menu waits hidden for its button
	// Place is where the drop stands in its row: pushed to the end of a
	// sheet's head (Carrel's task 205).
	Place Place
}

// Render prints the drop.
func (d Drop) Render() string { return d.html(nil) }

func (d Drop) html(m mix) string {
	t := mix(nil)
	if d.More {
		t = mix{"tc-bar__more"}
	}
	if d.Row {
		t = mix{"tc-acts__more"}
	}
	if !d.Open {
		if sum, ok := summaryOf(d.Trigger.html(t)); ok {
			// Closed, the menu opens without a script (U-8): details and its summary.
			return `<details` + class(m, "tc-drop", d.Place.class()) + `>` + sum + d.Menu.shown(nil, true) + `</details>`
		}
	}
	return `<span` + class(m, "tc-drop", d.Place.class()) + `>` + d.Trigger.html(t) + d.Menu.shown(nil, d.Open) + `</span>`
}

// summaryOf turns a trigger button into the summary of a details: the same
// look, opening its menu without a script. A link cannot be one.
func summaryOf(button string) (string, bool) {
	if !strings.HasPrefix(button, "<button") || !strings.HasSuffix(button, "</button>") {
		return "", false
	}
	s := "<summary" + strings.TrimSuffix(strings.TrimPrefix(button, "<button"), "</button>") + "</summary>"
	s = strings.Replace(s, ` type="button"`, "", 1)
	return s, true
}
