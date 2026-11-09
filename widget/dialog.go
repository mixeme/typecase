package widget

import "strings"

// DialogTitleGap is the space under a dialog's own title.
type DialogTitleGap int

const (
	DialogTitleGap0 DialogTitleGap = iota
	DialogTitleGapTight
	DialogTitleGapNormal
)

// DialogTitle is the title of a dialog that has no head: a sign-in page.
type DialogTitle struct {
	Text string
	Gap  DialogTitleGap
}

func (t DialogTitle) html(mix) string {
	return `<div` + class(nil, "tc-head__title", [...]string{"", "tc-head__title--below-tight", "tc-head__title--below-normal"}[t.Gap]) + `>` + esc(t.Text) + `</div>`
}

// DialogSub is the line under a dialog's own title.
type DialogSub struct {
	Parts []Inline
}

func (b DialogSub) html(mix) string {
	return `<div class="tc-head__sub tc-head__sub--below-wide">` + inlines(b.Parts) + `</div>`
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
	return `<div` + class(nil, "tc-auth__tagline", pick(t.Centered, "tc-auth__tagline--mid"), pick(!t.Centered, "tc-auth__tagline--below-wide")) +
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
	Links []Link // under the button: where else to go (Forgotten your password?)
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
			if l.Href != "" {
				s.WriteString(`<a class="tc-auth__link"` + attr("href", l.Href) + `>` + esc(l.Text) + `</a>`)
			} else {
				s.WriteString(`<span class="tc-auth__link">` + esc(l.Text) + `</span>`)
			}
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
	Src    string // the picture being cropped
	Note   string
	Fields []Field
}

func (p PhotoCrop) html(mix) string {
	var s strings.Builder
	pic := ""
	if p.Src != "" {
		pic = `<img class="tc-dialog__photo-img"` + attr("src", p.Src) + ` alt="">`
	}
	s.WriteString(`<div class="tc-dialog__split"><div` + attr("data-tc-colour", p.Colour) + ` class="tc-dialog__photo">` + pic + `<div class="tc-dialog__photo-frame"></div>` +
		`<div class="tc-dialog__photo-rule"></div><div class="tc-dialog__photo-rule-v"></div></div><div class="tc-stack tc-stack--wide tc-dialog__side">` +
		`<span class="tc-dim--muted tc-dim--t12-5">` + esc(p.Note) + `</span>`)
	for _, f := range p.Fields {
		s.WriteString(f.html(nil))
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
func (Field) dialogPart()       {}
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
func (PostForm) dialogPart()    {}

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
	// A step of a wizard whose widest step holds a table: the dialog is as
	// wide as that step, so the steps and the head stand where they stand
	// there, and what is under them keeps its own width (the owner's
	// decision of 13.10).
	DialogStep520 // the content 520 px
	DialogStep560 // the content 560 px
	DialogStep760 // the content 760 px
)

func (v DialogVariant) class() string {
	return [...]string{"", "is-wide", "is-narrow tc-dialog--padt0-padb0", "tc-dialog--max520", "tc-dialog--max540", "tc-dialog--max560",
		"tc-dialog--max440-padt0-padb0", "tc-dialog--max520-padt0-padb0", "tc-dialog--max420-flush",
		"is-wide tc-dialog--step520", "is-wide tc-dialog--step560", "is-wide tc-dialog--step760"}[v]
}

// Dialog — tc-dialog
//
// Роль. Центрированное полотно страницы: формы, мастера, подтверждения, вход в приложение.
// Вход. Dialog{Items, Variant, Action, Method, Fields, Ranges, RangesOf}.
// Состояния. narrow | wide | width | page | canvas.
// Инварианты. Ширина и поля — перечисление; главная кнопка — в подвале формы или в действиях шапки; диалог печатает фрагмент своего кадра.
// Поведение. Esc закрывает диалог, если он открыт поверх страницы, и возвращает фокус.
// Эталон. кадр `s-contacts-5`, фрагмент `.tc-dialog`.
// Содержимое. Шапка или свой заголовок со строкой, логотип, слоган, форма и её подвал, бар, шаги, сообщения, карточки, таблица, обрезка фото, скрытые формы кнопок подвала.
// Чего нет. Своего затемнения: диалог стоит на странице; поверх страницы — лист.
type Dialog struct {
	Items   []DialogPart
	Variant DialogVariant
	// Action makes the dialog a form that sends its fields there, as Sheet
	// does: a step of a wizard whose fields stand in several cards (Carrel's
	// task 149). Method is "post" when empty, Fields its hidden fields. A
	// form holds no form: when a part of the dialog is a form of its own,
	// the dialog is printed plain and that part sends.
	Action string
	Method string
	Fields []Hidden
	// Ranges and RangesOf make the form send, with the library's live
	// layer, parts of a file a step before it kept in the page — RangesOf
	// names that step's field (FileField.Ranges) — instead of the file:
	// Ranges is asked which parts, with this form's fields (Carrel's task
	// 220: the records of the chosen collections, at Restore). Without the
	// layer, or with no file kept, the form goes as it is.
	Ranges   string
	RangesOf string
}

// Render prints the dialog.
func (d Dialog) Render() string { return d.html(nil) }

func (d Dialog) html(m mix) string {
	var s strings.Builder
	var body, posts strings.Builder
	for _, it := range d.Items {
		// a footer button's hidden form stands after the dialog's own form,
		// not in it: a form holds no form
		if f, ok := it.(PostForm); ok && d.Action != "" {
			posts.WriteString(f.html(nil))
			continue
		}
		body.WriteString(it.html(nil))
	}
	// A form holds no form: a dialog one of whose parts is a form of its
	// own is printed plain, and that part sends.
	tag := "div"
	if d.Action != "" && !strings.Contains(body.String(), "<form") {
		tag = "form"
		s.WriteString(`<form` + class(m, "tc-dialog", d.Variant.class()) + attr("action", d.Action) + ` method="` + method(d.Method) + `"` + attr("data-tc-ranges", d.Ranges) + attr("data-tc-ranges-of", d.RangesOf) + `>`)
		for _, h := range d.Fields {
			s.WriteString(h.html(nil))
		}
	} else {
		s.WriteString(`<div` + class(m, "tc-dialog", d.Variant.class()) + `>`)
	}
	s.WriteString(body.String())
	s.WriteString(`</` + tag + `>`)
	// a file field among its parts: the form sends multipart (Carrel's
	// task 218: the archive of a restore from this computer)
	return multipart(s.String()) + posts.String()
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
