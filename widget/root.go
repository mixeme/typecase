package widget

import "strings"

// PageVariant is the drawn box of a page without a shell.
type PageVariant int

const (
	PagePlain              PageVariant = iota // a page of the web app: sign-in, «about», a desktop window's inside
	PagePrint                                 // a page on paper: white, black type, 26/30 px margins
	PageOnBackground                          // a small window's page: the background colour, 18 px around
	PageOnBackgroundColumn                    // the same, its parts one under another, 14 px apart
)

func (v PageVariant) class() string {
	return [...]string{"", "tc-root--print", "tc-root--on-bg tc-root--pad18",
		"tc-stack tc-stack--gap-14 tc-root--on-bg tc-root--pad18"}[v]
}

// PagePart is what a page without a shell holds.
type PagePart interface {
	part
	pagePart()
}

func (Auth) pagePart()         {}
func (Window) pagePart()       {}
func (Msg) pagePart()          {}
func (Form) pagePart()         {}
func (PrintTitle) pagePart()   {}
func (PrintSources) pagePart() {}
func (PrintRubric) pagePart()  {}
func (PrintRow) pagePart()     {}
func (PrintFoot) pagePart()    {}

// Page — tc-root
//
// Роль. Страница без оболочки: вход и пароль, «О программе», внутренность окна настольного приложения, страница на бумаге.
// Вход. Page{Items, Variant}.
// Состояния. plain | print | on-background | on-background-column.
// Инварианты. Корень страницы — один, как у оболочки (tc-root, контейнер tcbox); вид корня — перечисление; страница печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `s-auth-1`, фрагмент `.tc-root`.
// Содержимое. Страница входа, тело окна, сообщение, форма; на бумаге — заголовок, источники, рубрика, строки, подвал.
// Чего нет. Верхней полосы и рейла: страница с ними — оболочка (tc-shell).
type Page struct {
	Items   []PagePart
	Variant PageVariant
}

// Render prints the page.
func (p Page) Render() string { return p.html(nil) }

func (p Page) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-root", p.Variant.class()) + `>`)
	for _, it := range p.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// WindowGap is the space between the parts of a window's inside.
type WindowGap int

const (
	WindowGap9 WindowGap = iota
	WindowGap10
	WindowGap12
	WindowGap13
	WindowGap14
)

func (g WindowGap) class() string {
	return [...]string{"tc-stack--gap-9", "tc-stack--gap-10", "tc-stack--gap-12", "tc-stack--gap-13", "tc-stack--gap-14"}[g]
}

// WindowPart is what a desktop window's inside holds.
type WindowPart interface {
	part
	windowPart()
}

func (Form) windowPart()        {}
func (Msg) windowPart()         {}
func (Line) windowPart()        {}
func (Stack) windowPart()       {}
func (Logo) windowPart()        {}
func (WindowNote) windowPart()  {}
func (Diag) windowPart()        {}
func (FieldSet) windowPart()    {}
func (SheetFields) windowPart() {}
func (Area) windowPart()        {}
func (Dim) windowPart()         {}

// Window is the inside of a desktop app's window: the window itself — its
// frame, title bar and status line — is the system's, not the page's.
type Window struct {
	Items []WindowPart
	Gap   WindowGap
	Small bool // a small window over everything: 12 px around, the surface colour
}

func (w Window) html(mix) string {
	var s strings.Builder
	pad := "tc-win__body--pad16-20"
	if w.Small {
		pad = "tc-win__sheet--pad12"
	}
	s.WriteString(`<div class="tc-stack ` + w.Gap.class() + ` ` + pad + `">`)
	for _, it := range w.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// WindowNote is a window's sentence of what it is for.
type WindowNote struct {
	Text string
}

func (n WindowNote) html(mix) string { return `<div class="tc-win__note">` + esc(n.Text) + `</div>` }

// PrintTitle is the title a page on paper starts with.
type PrintTitle struct {
	Text string
}

func (t PrintTitle) html(mix) string { return `<div class="tc-print__h1">` + esc(t.Text) + `</div>` }

// PrintSources says what the paper was printed from.
type PrintSources struct {
	Text string
}

func (t PrintSources) html(mix) string {
	return `<div class="tc-print__sources">` + esc(t.Text) + `</div>`
}

// PrintRubric heads a group of rows on paper: a day, «Folders».
type PrintRubric struct {
	Text string
}

func (t PrintRubric) html(mix) string { return `<div class="tc-print__rub">` + esc(t.Text) + `</div>` }

// PrintFoot is the line at the foot of the paper: what, when, how much of it.
type PrintFoot struct {
	Text string
}

func (t PrintFoot) html(mix) string { return `<div class="tc-print__foot">` + esc(t.Text) + `</div>` }

// PrintRowPart is a cell of a row on paper.
type PrintRowPart interface {
	part
	printRowPart()
}

func (Mono) printRowPart()      {}
func (PrintCell) printRowPart() {}
func (PrintNum) printRowPart()  {}

// PrintRow is a row on paper: a time and what is at it, or a name and its
// numbers in three columns.
type PrintRow struct {
	Items   []PrintRowPart
	Columns bool // name, then two numbers: a folder's list
}

func (r PrintRow) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(nil, "tc-print__row", pick(r.Columns, "tc-print__row--3col")) + `>`)
	for _, it := range r.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// PrintCell is the main cell of a row on paper: a name, or a title in bold
// with its detail under it.
type PrintCell struct {
	Name   string
	Detail string // under the name, dim: where and in which collection
}

func (c PrintCell) html(mix) string {
	if c.Detail == "" {
		return `<span class="tc-print__cell">` + esc(c.Name) + `</span>`
	}
	return `<span class="tc-print__cell"><b>` + esc(c.Name) + `</b><br><span class="tc-dim--print tc-dim--t11">` + esc(c.Detail) + `</span></span>`
}

// PrintNum is a number on paper: a size, a date.
type PrintNum struct {
	Text string
}

func (n PrintNum) html(mix) string { return `<span class="tc-print__num">` + esc(n.Text) + `</span>` }
