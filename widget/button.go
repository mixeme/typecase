package widget

import "strings"

// ButtonKind is what pressing the button does to the data (the check on a
// button is one: what is left after pressing it).
type ButtonKind int

const (
	ButtonPlain       ButtonKind = iota // an ordinary action, a paper frame
	ButtonPrimary                       // the one primary action of the screen
	ButtonDanger                        // breaks a link, keeps the data
	ButtonDangerSolid                   // destroys a record, a collection, a key
)

func (k ButtonKind) class() string {
	switch k {
	case ButtonPrimary:
		return "is-primary"
	case ButtonDanger:
		return "is-danger"
	case ButtonDangerSolid:
		return "is-danger-solid"
	}
	return ""
}

// ButtonState says whether the button can be pressed.
type ButtonState int

const (
	ButtonReady    ButtonState = iota
	ButtonOff                  // shown in a selection bar but not for this selection
	ButtonDisabled             // cannot be pressed now; Reason says why
)

// ButtonVariant is a drawn variant of the button's box.
type ButtonVariant int

const (
	ButtonDefault     ButtonVariant = iota
	ButtonOnCard                    // pressed-looking, on card colour
	ButtonPadX3                     // icon button in a table row
	ButtonPadX4                     // icon button of a phone header
	ButtonPadX6                     // search button of a phone header
	ButtonPadX8                     // primary icon button of a phone header
	ButtonItalic                    // the italic toggle of an edit bar
	ButtonSelfStart                 // pinned to the top of its row
	ButtonCentred                   // label centred in a wide button
	ButtonCentredFill               // label centred, button fills the row
)

func (v ButtonVariant) class() string {
	return [...]string{"", "tc-btn--on-card", "tc-btn--padx3", "tc-btn--padx4", "tc-btn--padx6",
		"tc-btn--padx8", "tc-btn--fontstyle-italic", "tc-btn--self-start", "tc-btn--jc-center",
		"tc-btn--jc-center-fill"}[v]
}

// IconBox is the box a leading icon is drawn in when it is not the default.
type IconBox int

const (
	IconBoxDefault IconBox = iota
	IconBox16
	IconBox17
	IconBox19
)

func (b IconBox) class() string {
	return [...]string{"", "tc-ico--sq16", "tc-ico--sq17", "tc-ico--sq19"}[b]
}

// Button — tc-btn
//
// Роль. Действие одним нажатием: подпись, значок или оба, со стрелкой, если нажатие раскрывает список.
// Вход. Button{Label, Icon, IconBox, Bold, Folds, Aside, After, Tick, Menu, Kind, Small, Variant, State, Reason, Focus, Title, Submit, Href, Form, To, Name, Value, Adds, Opens, SaveAs, SaveFrom, Copy, CopyFrom, Copied, Swap, File, NoScript, Removes, Place, Secondary, Tertiary, Target}.
// Состояния. normal | icon | primary | danger | danger-solid | small | menu | folds | variant | off | disabled | focus | with-content | tertiary | link | link-off | post.
// Инварианты. Род, состояние и вариант — перечисления: главную тревожную кнопку, недоступную и выключенную сразу, не собрать; у кнопки без подписи есть доступное имя; кнопка печатает фрагмент своего кадра.
// Поведение. Нажатие — действие потребителя; кнопка со стрелкой (Menu) открывает свой список через узел меню, а не сама.
// Эталон. кадр `components-9`, фрагмент `.tc-btn`; формы — в таблице эталонов.
// Содержимое. Значок, подпись (жирная или уходящая при сужении), моноширинный хвост с текстом после, отметка перед подписью, стрелка.
// Чего нет. Своего цвета и своих отступов: кнопка-ссылка и кнопка отправки выглядят как любая другая; причину недоступности кнопка показывает подсказкой, а строкой под рядом — её контейнер.
type Button struct {
	Label   string
	Icon    Icon
	IconBox IconBox
	Bold    bool   // the label is set bold (the B toggle of an edit bar)
	Folds   bool   // the label goes when the bar is narrow; the icon stays
	Aside   *Mono  // a monospaced tail after the label: ".ics", a count
	After   string // text after the tail
	Tick    *Tick  // a tick before the label
	Menu    bool   // the button opens a list: the arrow is drawn
	Kind    ButtonKind
	Small   bool
	Variant ButtonVariant
	State   ButtonState
	Reason  string // why the button is disabled or off
	Focus   bool   // the drawn keyboard focus of the mockup's states
	Title   string // the accessible name when there is no label
	Submit  bool   // submits its form
	// Href makes the button a link: pressing it goes to the address (GET).
	// Disabled, the link has no address.
	Href string
	// Form is the ID of a PostForm the button submits from wherever it
	// stands; the button is a submit button then.
	Form string
	// To is the address this submit button sends its form to instead of the
	// form's own (formaction): «Send invitation» next to «Create link».
	To string
	// Name and Value are the pair a submit button adds to what is sent: which
	// action was pressed.
	Name  string
	Value string
	// Adds is the ID of the Template this button adds a row of, in place,
	// with the library's script; without it the button submits (Name/Value)
	// and the server answers with one more row.
	Adds  string
	Opens string // the ID of the sheet (Scrim) the button opens in place
	// SaveAs is the name of the file the button saves the text of the
	// element SaveFrom names into, with the library's script (Carrel's task
	// 94: a secret shown once has no address to download from). Without a
	// script the button is not shown: Copy and selecting the text remain.
	SaveAs   string
	SaveFrom string
	// Copy is the text the button puts on the clipboard, CopyFrom the ID of
	// the element whose text it puts there, with the library's script
	// (Carrel's task 118). Without a script a button is not shown — the text
	// is selected by hand; one with Href stays a link to the page that shows
	// the text.
	Copy     string
	CopyFrom string
	// Copied is what the button says for a moment after it has copied
	// («Copied»). Its right edge stays where it is and the button grows or
	// shrinks to the left, so the buttons after it do not move.
	Copied string
	// Swap, with Href, names the elements (IDs, separated by spaces) the
	// library's live layer takes from the page at Href and puts in place of
	// the same ones here — no reload; without the layer the link goes there
	// (Source, the width of a note's text; Carrel's task 104).
	Swap string
	// Removes takes away the row it stands in, with the library's script:
	// the row of a picked file, before Save (FileField.Row) — the button
	// sends nothing; or the row of a repeated field in a FieldSet, the fields
	// before it — they are emptied and hidden in place, under their names
	// (Carrel's task 86). Without a script the second one submits (Submit,
	// Name, Value) and the server answers without the row.
	Removes bool
	// File makes the button a label that picks a file for a PostForm
	// (Carrel's task 109): the browser opens the choice without a script.
	File *FileField
	// NoScript shows the button only on a page without a script: the
	// «Upload» that sends a picked file there (Carrel's tasks 109, 90).
	NoScript bool
	Place    Place
	// Secondary is the first to go under «⋯» when its bar narrows.
	Secondary bool
	// Tertiary goes under «⋯» on a phone, after the secondary ones.
	Tertiary bool
	// Target is the name of the zone the links and forms inside ask their
	// address for with the live layer, instead of loading the page (Carrel's
	// task 84); without the layer they are plain links and forms.
	Target string
}

// Render prints the button.
func (b Button) Render() string { return b.html(nil) }

func (b Button) html(m mix) string {
	var s strings.Builder
	typ := "button"
	if b.Submit {
		typ = "submit"
	}
	title := b.Title
	if title == "" && b.State != ButtonReady {
		title = b.Reason
	}
	state := ""
	if b.State == ButtonOff {
		state = "is-off"
	}
	tag := "button"
	cls := class(m, "tc-btn", b.Kind.class(), is("sm", b.Small), state, is("focus", b.Focus), b.Variant.class(), b.Place.class(), rank(b.Secondary), rank3(b.Tertiary), pick(b.NoScript, "tc-nojs"))
	common := attr("title", title) + attr("aria-label", pick(b.Label == "", b.Title)) + attr("data-tc-open", b.Opens)
	copies := attr("data-tc-copy", b.Copy) + attr("data-tc-copy-from", b.CopyFrom) + attr("data-tc-copied", b.Copied)
	if b.File != nil {
		tag = "label"
		s.WriteString(`<label` + cls + common + `>`)
	} else if b.Href != "" {
		tag = "a"
		if b.State == ButtonDisabled {
			s.WriteString(`<a` + cls + common + ` aria-disabled="true">`)
		} else {
			s.WriteString(`<a` + cls + attr("href", b.Href) + attr("data-tc-swap", b.Swap) + attr("data-tc-target", b.Target) + copies + common + `>`)
		}
	} else {
		if b.Form != "" {
			typ = "submit"
		}
		s.WriteString(`<button` + cls + ` type="` + typ + `"` + attr("data-tc-target", b.Target) + attr("form", b.Form) + attr("formaction", b.To) + attr("name", b.Name) + attr("value", b.Value) + attr("data-tc-add", b.Adds) + flag("data-tc-remove", b.Removes) +
			attr("data-tc-save", b.SaveAs) + attr("data-tc-from", b.SaveFrom) + copies + common + flag("disabled", b.State == ButtonDisabled) + `>`)
	}
	if b.Tick != nil {
		s.WriteString(b.Tick.html(nil))
	}
	if b.Icon != "" {
		s.WriteString(`<svg viewBox="0 0 24 24"` + class(nil, b.IconBox.class()) + `>` + b.Icon.use() + `</svg>`)
	}
	switch {
	case b.Label == "":
	case b.Folds:
		s.WriteString(`<span class="tc-btn__label">` + esc(b.Label) + `</span>`)
	case b.Bold:
		s.WriteString(`<b>` + esc(b.Label) + `</b>`)
	default:
		s.WriteString(esc(b.Label))
	}
	if b.Aside != nil {
		s.WriteString(b.Aside.html(nil))
	}
	if b.After != "" {
		s.WriteString(esc(b.After))
	}
	if b.Menu {
		s.WriteString(`<svg viewBox="0 0 24 24">` + Icon("chev").use() + `</svg>`)
	}
	if b.File != nil {
		s.WriteString(b.File.input())
	}
	s.WriteString(`</` + tag + `>`)
	return s.String()
}

// Close — tc-close
//
// Роль. Крестик, который закрывает панель, лист или сообщение.
// Вход. Close{Label, Place, Href}.
// Состояния. нет.
// Инварианты. У крестика всегда есть доступное имя — что он закрывает; крестик печатает фрагмент своего кадра.
// Поведение. Закрывает свой слой; Esc делает то же, фокус возвращается на вызвавший контрол.
// Эталон. кадр `components-33`, фрагмент `.tc-close`.
// Содержимое. Значок крестика, больше ничего.
// Чего нет. Подписи: крестик без слов, имя — в подсказке.
type Close struct {
	Label string // what it closes: "Close the panel"
	Place Place
	// Href makes the cross a link: "#" closes a sheet opened by its address,
	// without scripts.
	Href string
}

// Render prints the close button.
func (c Close) Render() string { return c.html(nil) }

func (c Close) html(m mix) string {
	if c.Href != "" {
		return `<a` + class(m, "tc-btn", "tc-close", c.Place.class()) + attr("href", c.Href) + attr("title", c.Label) + attr("aria-label", c.Label) + `><svg viewBox="0 0 24 24">` +
			Icon("close").use() + `</svg></a>`
	}
	return `<button` + class(m, "tc-btn", "tc-close", c.Place.class()) + attr("title", c.Label) + attr("aria-label", c.Label) + ` type="button"><svg viewBox="0 0 24 24">` +
		Icon("close").use() + `</svg></button>`
}
