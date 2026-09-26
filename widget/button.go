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
// Вход. Button{Label, Icon, IconBox, Bold, Folds, Aside, After, Tick, Menu, Kind, Small, Variant, State, Reason, Focus, Title, Submit, Opens, Place, Secondary, Tertiary}.
// Состояния. normal | icon | primary | danger | danger-solid | small | menu | folds | variant | off | disabled | focus | with-content | tertiary.
// Инварианты. Род, состояние и вариант — перечисления: главную тревожную кнопку, недоступную и выключенную сразу, не собрать; у кнопки без подписи есть доступное имя; кнопка печатает фрагмент своего кадра.
// Поведение. Нажатие — действие потребителя; кнопка со стрелкой (Menu) открывает свой список через узел меню, а не сама.
// Эталон. кадр `components-9`, фрагмент `.tc-btn`; формы — в таблице эталонов.
// Содержимое. Значок, подпись (жирная или уходящая при сужении), моноширинный хвост с текстом после, отметка перед подписью, стрелка.
// Чего нет. Своего цвета, своих отступов и ссылки вместо действия; причину недоступности кнопка показывает подсказкой, а строкой под рядом — её контейнер.
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
	Opens   string // the ID of the sheet (Scrim) the button opens in place
	Place   Place
	// Secondary is the first to go under «⋯» when its bar narrows.
	Secondary bool
	// Tertiary goes under «⋯» on a phone, after the secondary ones.
	Tertiary bool
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
	s.WriteString(`<button` + class(m, "tc-btn", b.Kind.class(), is("sm", b.Small), state, is("focus", b.Focus), b.Variant.class(), b.Place.class(), rank(b.Secondary), rank3(b.Tertiary)) +
		` type="` + typ + `"` + attr("title", title) + attr("aria-label", pick(b.Label == "", b.Title)) + attr("data-tc-open", b.Opens) + flag("disabled", b.State == ButtonDisabled) + `>`)
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
	s.WriteString(`</button>`)
	return s.String()
}

// Close — tc-close
//
// Роль. Крестик, который закрывает панель, лист или сообщение.
// Вход. Close{Label, Place}.
// Состояния. нет.
// Инварианты. У крестика всегда есть доступное имя — что он закрывает; крестик печатает фрагмент своего кадра.
// Поведение. Закрывает свой слой; Esc делает то же, фокус возвращается на вызвавший контрол.
// Эталон. кадр `components-33`, фрагмент `.tc-close`.
// Содержимое. Значок крестика, больше ничего.
// Чего нет. Подписи: крестик без слов, имя — в подсказке.
type Close struct {
	Label string // what it closes: "Close the panel"
	Place Place
}

// Render prints the close button.
func (c Close) Render() string { return c.html(nil) }

func (c Close) html(m mix) string {
	return `<button` + class(m, "tc-btn", "tc-close", c.Place.class()) + attr("title", c.Label) + attr("aria-label", c.Label) + ` type="button"><svg viewBox="0 0 24 24">` +
		Icon("close").use() + `</svg></button>`
}
