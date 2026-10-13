package widget

import "strings"

// ListGroup is the heading of a group of rows in a page: a letter, a day, a
// kind of record, with how many there are. It stands before the group's list,
// table or empty state.
type ListGroup struct {
	Title Rubric
	Count *Num
}

func (g ListGroup) html(mix) string {
	s := `<div class="tc-list__group">` + g.Title.html(nil)
	if g.Count != nil {
		s += g.Count.html(nil)
	}
	return s + `</div>`
}

// MainTree is a tree set in a page with the page's gutters.
type MainTree struct {
	Tree Tree
}

func (t MainTree) html(mix) string {
	return `<div class="tc-main__tree">` + t.Tree.html(nil) + `</div>`
}

// MainHint is a line of explanation in a page, set with the page's gutters.
type MainHint struct {
	Text string
}

func (h MainHint) html(mix) string {
	return `<div class="tc-field__hint tc-field__hint--gutter">` + esc(h.Text) + `</div>`
}

// BottomItem is a section of a phone's bottom navigation: an icon over a
// short name.
type BottomItem struct {
	Icon  Icon
	Label string
	Href  string
}

// BottomNav is a phone's navigation at the bottom of the page.
type BottomNav struct {
	Items []BottomItem
	On    int // the current section, counted from 1
}

func (n BottomNav) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-nav--bottom">`)
	for i, it := range n.Items {
		s.WriteString(NavItem{Icon: it.Icon, Label: it.Label, Href: it.Href}.html(i+1 == n.On, ""))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// FootActs is a row of buttons at the foot of a phone's page.
type FootActs struct {
	Items []Button
}

func (f FootActs) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-line tc-line--gap-8 tc-main__foot-acts">`)
	for _, b := range f.Items {
		s.WriteString(b.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// MainPart is what a page holds.
type MainPart interface {
	part
	mainPart()
}

func (Head) mainPart()      {}
func (Bar) mainPart()       {}
func (PostForm) mainPart()  {}
func (List) mainPart()      {}
func (ListGroup) mainPart() {}
func (Form) mainPart()      {}
func (FormFoot) mainPart()  {}
func (MainTree) mainPart()  {}
func (MainHint) mainPart()  {}
func (RailSec) mainPart()   {}
func (Zone) mainPart()      {}
func (BottomNav) mainPart() {}
func (FootActs) mainPart()  {}
func (Banner) mainPart()    {}
func (Strip) mainPart()     {}
func (MainFoot) mainPart()  {}
func (Stage) mainPart()     {}
func (Person) mainPart()    {}

// MainVariant is the drawn form of a page.
type MainVariant int

const (
	MainPlain MainVariant = iota
	MainZone              // the shell's scheme: a zone's label in the middle
	MainRoomy             // 180 px below: room for a queue over the page
	MainShort             // at least 88 px: a page cut short in a frame
)

func (v MainVariant) class() string {
	return [...]string{"", "tc-main--zone", "tc-main--padb180", "tc-main--minh88"}[v]
}

// Main — tc-main
//
// Роль. Страница в середине оболочки: шапка, бар, списки по группам, формы, документ.
// Вход. Main{Items, Variant, ID, Zone}.
// Состояния. plain | groups | zone | variant | phone.
// Инварианты. Страница прокручивается сама, оболочка — нет; порядок — шапка, бар, содержимое; страница печатает фрагмент своего кадра.
// Поведение. нет: поведение у того, что в ней.
// Эталон. кадр `s-contacts-1`, фрагмент `.tc-main`.
// Содержимое. Шапка, бар, заголовки групп, списки, формы, дерево, пояснение, секция рейла (телефон), нижняя навигация и кнопки телефона, человек (телефон); у публичной страницы — баннер, полоса адреса, сцена файла, подвал.
// Чего нет. Своих полей по краям: поля задают шапка, бар и строки.
type Main struct {
	Items   []MainPart
	Variant MainVariant
	ID      string // what a link with Swap names: the column it puts in place
	// Zone is the name the consumer gives this area of the page: the live
	// layer puts an answer for that name here (Carrel's task 84). The library
	// has no names of its own.
	Zone string
}

// Render prints the page.
func (p Main) Render() string { return p.html(nil) }

func (p Main) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-main", p.Variant.class()) + attr("id", p.ID) + attr("data-tc-zone", p.Zone) + `>`)
	for _, it := range p.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// SideActs is a row of buttons in a panel, set with the panel's gutters.
type SideActs struct {
	Items []Button
}

func (a SideActs) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-line tc-line--gap-6 tc-side__acts">`)
	for _, b := range a.Items {
		s.WriteString(b.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// SidePart is what a side panel holds.
type SidePart interface {
	part
	sidePart()
}

func (Head) sidePart()     {}
func (Form) sidePart()     {}
func (FormFoot) sidePart() {}
func (Zone) sidePart()     {}
func (SideActs) sidePart() {}

// SideVariant is the drawn form of a side panel.
type SideVariant int

const (
	SidePlain SideVariant = iota
	SideZone              // the shell's scheme: a zone's label in the middle
	SideDrop              // a panel a file is dropped on
)

func (v SideVariant) class() string {
	return [...]string{"", "tc-side--zone", "tc-side--drop"}[v]
}

// Side — tc-side
//
// Роль. Панель справа: подробности выбранной строки, её поля и действия.
// Вход. Side{Items, Variant, Zone}.
// Состояния. plain | acts | drop | zone.
// Инварианты. Ширина — токен оболочки, а не свойство панели; панель закрывается крестиком шапки; панель печатает фрагмент своего кадра.
// Поведение. Esc закрывает панель и возвращает фокус строке, которая её открыла.
// Эталон. кадр `s-contacts-1`, фрагмент `.tc-side`.
// Содержимое. Шапка, секции, поля, опасная зона, форма, сообщение, строка кнопок.
// Чего нет. Второй панели: подробности подробностей открываются страницей.
type Side struct {
	Items   []SidePart
	Variant SideVariant
	// Zone is the name the consumer gives this area of the page: the live
	// layer puts an answer for that name here (Carrel's task 84). The library
	// has no names of its own.
	Zone string
}

// Render prints the panel.
func (p Side) Render() string { return p.html(nil) }

func (p Side) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-side", p.Variant.class()) + attr("data-tc-zone", p.Zone) + `>`)
	for _, it := range p.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// RailPart is what the shell's rail holds.
type RailPart interface {
	part
	railPart()
}

func (Nav) railPart()      {}
func (RailSec) railPart()  {}
func (RailFoot) railPart() {}
func (Zone) railPart()     {}

// RailForm is the drawn form of the shell's rail.
type RailForm int

const (
	RailFull  RailForm = iota
	RailIcons          // 56 px: icons only
	RailZone           // the shell's scheme: a zone's label in the middle
)

// ShellHeight is the least height of the shell's body in a frame.
type ShellHeight int

const (
	ShellHeightAuto ShellHeight = iota
	ShellHeight230
	ShellHeight280
	ShellHeight340
	ShellHeight390
	ShellHeight560
)

func (h ShellHeight) class() string {
	return [...]string{"", "tc-shell__body--minh230", "tc-shell__body--minh280", "tc-shell__body--minh340",
		"tc-shell__body--minh390", "tc-shell__body--minh560"}[h]
}

// ShellBody is the shell under the top bar: the rail, the page, the side
// panel. It stands in the shell, or under an overlay that dims it.
type ShellBody struct {
	Rail     []RailPart // none: a page without a rail
	RailForm RailForm
	Main     Main
	Side     *Side
	Height   ShellHeight
}

func (b ShellBody) html(mix) string {
	rail := make([]part, len(b.Rail))
	for i, r := range b.Rail {
		rail[i] = r
	}
	return shellBody(rail, b.RailForm, b.Main, b.Side, b.Height)
}

// shellBody prints the shell's body, in the shell or under an overlay.
func shellBody(rail []part, form RailForm, main Main, side *Side, height ShellHeight) string {
	var s strings.Builder
	s.WriteString(`<div` + class(nil, "tc-shell__body", pick(len(rail) == 0, "tc-shell__body--no-rail"), pick(side != nil, "tc-shell__body--side"),
		height.class(), pick(form == RailIcons, "tc-shell__body--tcrail-56")) + `>`)
	if len(rail) > 0 {
		s.WriteString(`<div` + class(nil, "tc-shell__rail", [...]string{"", "tc-shell__rail--icons", "tc-shell__rail--zone"}[form]) + `>`)
		for _, it := range rail {
			s.WriteString(it.html(nil))
		}
		s.WriteString(`</div>`)
	}
	s.WriteString(main.html(nil))
	if side != nil {
		s.WriteString(side.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// ShellContent is what stands in the shell under the top bar.
type ShellContent interface {
	part
	shellContent()
}

func (ShellBody) shellContent() {}

// Shell — tc-shell
//
// Роль. Оболочка приложения: верхняя полоса, рейл, страница, панель — корень экрана.
// Вход. Shell{Top, Body}.
// Состояния. plain | side | no-rail | no-top | icons | zone | overlay.
// Инварианты. Зон четыре, и каждая на своём месте: полоса сверху, рейл слева, страница, панель справа; как оболочка ведёт себя на узком экране, решает CSS (--tc-shell-mode); оболочка печатает фрагмент своего кадра.
// Поведение. На узком экране рейл становится ящиком, панель — листом поверх страницы.
// Эталон. кадр `shell-1`, фрагмент `.tc-shell`.
// Содержимое. Верхняя полоса и тело: рейл, страница, панель.
// Чего нет. Порогов в коде: ширины и пороги — токены и контейнерные запросы CSS.
type Shell struct {
	Top  *Top
	Body ShellContent
}

// Render prints the shell.
func (s Shell) Render() string { return s.html(nil) }

func (s Shell) html(m mix) string {
	out := `<div` + class(m, "tc-root", "tc-shell") + `>`
	if s.Top != nil {
		out += s.Top.html(nil)
	}
	if s.Body != nil {
		out += s.Body.html(nil)
	}
	return out + `</div>`
}
