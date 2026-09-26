package widget

import "strings"

// DialogTitleGap is the space under a dialog's own title.
type DialogTitleGap int

const (
	DialogTitleGap0 DialogTitleGap = iota
	DialogTitleGap4
	DialogTitleGap6
	DialogTitleGap8
)

// DialogTitle is the title of a dialog that has no head: a sign-in page.
type DialogTitle struct {
	Text string
	Gap  DialogTitleGap
}

func (t DialogTitle) html(mix) string {
	return `<div` + class(nil, "tc-head__title", [...]string{"", "tc-head__title--bot4", "tc-head__title--bot6", "tc-head__title--bot8"}[t.Gap]) + `>` + esc(t.Text) + `</div>`
}

// DialogSub is the line under a dialog's own title.
type DialogSub struct {
	Parts []Inline
}

func (b DialogSub) html(mix) string {
	return `<div class="tc-head__sub tc-head__sub--bot12">` + inlines(b.Parts) + `</div>`
}

// DialogLogo is the logo over a sign-in dialog.
type DialogLogo struct {
	Logo Logo
	Wide bool // less space under it: a page with a title under the logo
}

func (l DialogLogo) html(mix) string {
	return `<div` + class(nil, "tc-line", pick(!l.Wide, "tc-dialog__logo"), pick(l.Wide, "tc-dialog__logo--wide")) + `>` + l.Logo.html(nil) + `</div>`
}

// Tagline is what the application is, under its logo: a lead in bold and a
// sentence.
type Tagline struct {
	Lead     string
	Text     string
	Centered bool // under a centred logo; otherwise 14 px above what follows
}

func (t Tagline) html(mix) string {
	return `<div` + class(nil, "tc-auth__tagline", pick(t.Centered, "tc-auth__tagline--mid"), pick(!t.Centered, "tc-auth__tagline--bot14")) +
		`><strong class="tc-auth__tagline-lead">` + esc(t.Lead) + `</strong> ` + esc(t.Text) + `</div>`
}

// AuthStackPart is what the sign-in column holds.
type AuthStackPart interface {
	part
	authStackPart()
}

func (Form) authStackPart()   {}
func (Button) authStackPart() {}

// AuthStack is the sign-in column: the form, the button, the links under
// them.
type AuthStack struct {
	Items []AuthStackPart
	Links []string
}

func (a AuthStack) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-auth__stack">`)
	for _, it := range a.Items {
		s.WriteString(it.html(nil))
	}
	if len(a.Links) > 0 {
		s.WriteString(`<div class="tc-auth__links">`)
		for _, l := range a.Links {
			s.WriteString(`<span class="tc-auth__link">` + esc(l) + `</span>`)
		}
		s.WriteString(`</div>`)
	}
	s.WriteString(`</div>`)
	return s.String()
}

// Lede is the one sentence of a dialog drawn as a sample canvas.
type Lede struct {
	Text string
}

func (l Lede) html(mix) string { return `<p class="tc-dialog__lede">` + esc(l.Text) + `</p>` }

// PhotoCrop is a photo being cropped beside the fields that set it: the
// picture with its frame and rules, a note, the fields.
type PhotoCrop struct {
	Colour string // the picture's place, in a colour
	Note   string
	Fields []Field
	After  *Dim // what is done to the file on its way in
}

func (p PhotoCrop) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-dialog__split"><div` + attr("data-tc-colour", p.Colour) + ` class="tc-dialog__photo"><div class="tc-dialog__photo-frame"></div>` +
		`<div class="tc-dialog__photo-rule"></div><div class="tc-dialog__photo-rule-v"></div></div><div class="tc-stack tc-stack--gap-10 tc-dialog__side">` +
		`<span class="tc-dim--muted tc-dim--t12-5">` + esc(p.Note) + `</span>`)
	for _, f := range p.Fields {
		s.WriteString(f.html(nil))
	}
	if p.After != nil {
		s.WriteString(p.After.html(nil))
	}
	s.WriteString(`</div></div>`)
	return s.String()
}

// DialogPart is what a dialog holds.
type DialogPart interface {
	part
	dialogPart()
}

func (Head) dialogPart()        {}
func (Form) dialogPart()        {}
func (FormFoot) dialogPart()    {}
func (Bar) dialogPart()         {}
func (DialogTitle) dialogPart() {}
func (DialogSub) dialogPart()   {}
func (DialogLogo) dialogPart()  {}
func (Tagline) dialogPart()     {}
func (AuthStack) dialogPart()   {}
func (Lede) dialogPart()        {}
func (PhotoCrop) dialogPart()   {}

// DialogVariant is the drawn width and padding of a dialog.
type DialogVariant int

const (
	DialogPlain     DialogVariant = iota // 760 px
	DialogWide                           // a table inside
	DialogNarrow                         // a sign-in page: narrow, no padding above and below
	Dialog520                            // 520 px
	Dialog540                            // 540 px
	Dialog560                            // 560 px
	DialogPage440                        // 440 px, no padding above and below: a page of the sign-in flow
	DialogPage520                        // 520 px, no padding above and below
	DialogCanvas420                      // 420 px, flush: the canvas drawn alone
)

func (v DialogVariant) class() string {
	return [...]string{"", "is-wide", "is-narrow tc-dialog--padt0-padb0", "tc-dialog--max520", "tc-dialog--max540", "tc-dialog--max560",
		"tc-dialog--max440-padt0-padb0", "tc-dialog--max520-padt0-padb0", "tc-dialog--max420-flush"}[v]
}

// Dialog — tc-dialog
//
// Роль. Центрированное полотно страницы: формы, мастера, подтверждения, вход в приложение.
// Вход. Dialog{Items, Variant}.
// Состояния. narrow | wide | width | page | canvas.
// Инварианты. Ширина и поля — перечисление; главная кнопка — в подвале формы или в действиях шапки; диалог печатает фрагмент своего кадра.
// Поведение. Esc закрывает диалог, если он открыт поверх страницы, и возвращает фокус.
// Эталон. кадр `s-contacts-5`, фрагмент `.tc-dialog`.
// Содержимое. Шапка или свой заголовок со строкой, логотип, слоган, форма и её подвал, бар, шаги, сообщения, карточки, таблица, обрезка фото.
// Чего нет. Своего затемнения: диалог стоит на странице; поверх страницы — лист.
type Dialog struct {
	Items   []DialogPart
	Variant DialogVariant
}

// Render prints the dialog.
func (d Dialog) Render() string { return d.html(nil) }

func (d Dialog) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-dialog", d.Variant.class()) + `>`)
	for _, it := range d.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

func (Dialog) mainPart() {}

// Auth — tc-auth
//
// Роль. Страница входа: один диалог посередине пустой страницы.
// Вход. Auth{Dialog}.
// Состояния. plain.
// Инварианты. На странице один диалог; страница печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `s-auth-1`, фрагмент `.tc-auth`.
// Содержимое. Диалог.
// Чего нет. Рейла, верхней полосы и навигации: до входа их нет.
type Auth struct {
	Dialog Dialog
}

// Render prints the page.
func (a Auth) Render() string { return a.html(nil) }

func (a Auth) html(m mix) string {
	return `<div` + class(m, "tc-auth") + `>` + a.Dialog.html(nil) + `</div>`
}
