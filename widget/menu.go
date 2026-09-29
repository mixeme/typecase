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
	// Name makes the item a checkbox of the menu's form (Menu.Action): ticked
	// (Tick), it sends Value; the items' order is the order they are sent in.
	// Fixed keeps it ticked and off: a column that is always shown.
	Name  string
	Value string
	Fixed bool
	Drag  bool // the drawn state of an item being dragged (the mockup's states)
	// Tertiary is the twin of a bar's tertiary item in the menu of «⋯»:
	// shown on a phone only, where the item itself is under «⋯».
	Tertiary bool
}

func (it MenuItem) html(mix) string {
	var s strings.Builder
	tag := "button"
	cls := class(nil, is("on", it.On), "tc-menu__item", it.Tone.class(), rank3(it.Tertiary))
	checked := attr("aria-checked", pick(it.Tick != nil && it.Tick.Ticked != TickedOff, "true"))
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
	if it.Href != "" {
		tag = "a"
		if it.Reason != "" {
			s.WriteString(`<a` + cls + ` role="menuitem"` + checked + attr("title", it.Reason) + ` aria-disabled="true">`)
		} else {
			s.WriteString(`<a` + cls + attr("href", it.Href) + ` role="menuitem"` + checked + `>`)
		}
	} else {
		s.WriteString(`<button` + cls + ` type="button" role="menuitem"` + checked + attr("title", it.Reason) + flag("disabled", it.Reason != "") + `>`)
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
func (Rubric) menuPart()      {}
func (Input) menuPart()       {}
func (Range) menuPart()       {}

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
// Вход. Menu{Items, Width, Pop, Fields, Inline, Action, Hidden}.
// Состояния. on | icons | marks | danger | heading | pop | fields | found | form.
// Инварианты. Пункт — кнопка с ролью пункта меню; ширина и место — перечисления; меню печатает фрагмент своего кадра.
// Поведение. Стрелки ходят по пунктам, Enter выбирает, Esc закрывает меню и возвращает фокус кнопке, которая его открыла.
// Эталон. кадр `components-1`, фрагмент `.tc-menu`.
// Содержимое. Пункты (значок, метка, галка, название, приглушённый хвост, клавиша), разделители, заголовки, поле, строка с переключателем, диапазон, подвал.
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
	s.WriteString(`<` + tag + class(m, "tc-menu", is("pop", u.Pop != MenuInPlace), is("right", u.Pop == MenuUnderEnd), u.Width.class(), pick(u.Fields, "tc-menu--fields")) + form + ` role="menu"` + flag("hidden", !open) + `>`)
	for _, h := range u.Hidden {
		s.WriteString(h.html(nil))
	}
	for _, it := range u.Items {
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
// Вход. Drop{Trigger, Menu, More, Open}.
// Состояния. more | pick.
// Инварианты. Меню стоит сразу за тем, что его открывает; «⋯» бара — кнопка бара и тогда, когда она в выпадашке; выпадашка печатает фрагмент своего кадра.
// Поведение. Нажатие открывает меню под кнопкой; Esc закрывает и возвращает фокус.
// Эталон. кадр `components-19`, фрагмент `.tc-drop`.
// Содержимое. Кнопка или поле выбора, меню.
// Чего нет. Своего вида: вид у кнопки и меню. Без сценария закрытое меню не открыть.
type Drop struct {
	Trigger DropTrigger
	Menu    Menu
	More    bool // the trigger is its bar's «⋯»
	Open    bool // drawn open, as the frames draw it; closed, the menu waits hidden for its button
}

// Render prints the drop.
func (d Drop) Render() string { return d.html(nil) }

func (d Drop) html(m mix) string {
	t := mix(nil)
	if d.More {
		t = mix{"tc-bar__more"}
	}
	return `<span` + class(m, "tc-drop") + `>` + d.Trigger.html(t) + d.Menu.shown(nil, d.Open) + `</span>`
}
