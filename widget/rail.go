package widget

import "strings"

// NavItem is a section of the navigation: an icon or an emblem, the name, a
// count.
type NavItem struct {
	Icon   Icon
	Emblem *Emblem // a service's emblem instead of the icon
	Label  string
	Count  *Num
	Href   string
	Title  string // the name of an item drawn without a label
}

func (it NavItem) html(on bool, opens string) string {
	var s strings.Builder
	s.WriteString(`<a` + class(nil, is("on", on), "tc-nav__item") + attr("href", it.Href) + attr("title", it.Title) +
		attr("aria-label", pick(it.Label == "", it.Title)) + attr("aria-current", pick(on, "page")) + attr("data-tc-open", pick(on, opens)) + `>`)
	if it.Emblem != nil {
		s.WriteString(it.Emblem.html(nil))
	} else {
		s.WriteString(`<svg>` + it.Icon.use() + `</svg>`)
	}
	s.WriteString(esc(it.Label))
	if it.Count != nil {
		s.WriteString(it.Count.html(nil))
	}
	s.WriteString(`</a>`)
	return s.String()
}

// Nav — tc-nav
//
// Роль. Навигация по разделам: в рейле приложения, в рейле настроек, строка «назад».
// Вход. Nav{Items, On, Apart, Opens}.
// Состояния. plain | back | apart | icons.
// Инварианты. Текущий раздел — один, его номер, а не флаг у каждого пункта; пункт — ссылка; навигация печатает фрагмент своего кадра.
// Поведение. Нажатие открывает раздел; текущий помечен для чтения (aria-current). В рейле из значков (641–1007) повторное нажатие на текущий раздел открывает лист его источников (Opens).
// Эталон. кадр `s-contacts-1`, фрагмент `.tc-nav`.
// Содержимое. Пункты: значок или знак службы, название, число.
// Чего нет. Вложенных пунктов: подразделы — секции рейла под навигацией.
type Nav struct {
	Items []NavItem
	On    int  // the current section, counted from 1; 0 for none
	Apart bool // 2 px above: a single «back» under the top bar
	// Opens is the ID of the sheet (Scrim) with the section's sources: in the
	// rail of icons (Р-3, rail) pressing the current section again opens it.
	Opens string
}

// Render prints the navigation.
func (n Nav) Render() string { return n.html(nil) }

func (n Nav) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-nav", pick(n.Apart, "tc-nav--top2")) + `>`)
	for i, it := range n.Items {
		s.WriteString(it.html(i+1 == n.On, n.Opens))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// RailSecPart is what a section of the rail holds.
type RailSecPart interface {
	part
	railSecPart()
}

func (Rubric) railSecPart()  {}
func (TickRow) railSecPart() {}
func (Tree) railSecPart()    {}
func (Nav) railSecPart()     {}

// RailSec — tc-railsec
//
// Роль. Секция рейла: рубрика и то, что под ней, — источники с галками, дерево, навигация настроек, теги.
// Вход. RailSec{Items, Tags, Note, Apart}.
// Состояния. plain | apart | tags | empty.
// Инварианты. Рубрика — первой, если есть; теги — строкой после строк; секция печатает фрагмент своего кадра.
// Поведение. нет: поведение у строк и тегов.
// Эталон. кадр `components-11`, фрагмент `.tc-railsec`.
// Содержимое. Рубрика, строки источников, дерево, навигация, строка тегов, примечание «ничего нет».
// Чего нет. Сворачивания: секция всегда открыта; на узком экране рейл прячет секции целиком.
type RailSec struct {
	Items []RailSecPart
	Tags  []Tag
	Note  string // what the section says when it has nothing
	Apart bool   // 14 px above: a section set apart from the one before
}

// Render prints the section.
func (r RailSec) Render() string { return r.html(nil) }

func (r RailSec) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-railsec", pick(r.Apart, "tc-railsec--top14")) + `>`)
	for _, it := range r.Items {
		s.WriteString(it.html(nil))
	}
	if len(r.Tags) > 0 {
		s.WriteString(`<div class="tc-line tc-line--gap-4 tc-railsec__tags">`)
		for _, t := range r.Tags {
			s.WriteString(t.html(nil))
		}
		s.WriteString(`</div>`)
	}
	if r.Note != "" {
		s.WriteString(`<div class="tc-railsec__note">` + esc(r.Note) + `</div>`)
	}
	s.WriteString(`</div>`)
	return s.String()
}

// RailStatusTone is the state a line of the rail's foot reports.
type RailStatusTone int

const (
	RailStatusRunning RailStatusTone = iota
	RailStatusAlert                  // not answering
	RailStatusOff                    // stopped
)

// RailStatus — tc-railstatus
//
// Роль. Состояние службы в подвале рейла: точка, что с ней, вторая строка — когда она отвечала.
// Вход. RailStatus{Tone, Text, Detail}.
// Состояния. running | alert | off.
// Инварианты. Цвет точки следует из состояния, а не задаётся отдельно; обе строки строятся одинаково во всех состояниях; строка состояния печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `components-35`, фрагмент `.tc-railstatus`.
// Содержимое. Точка, строка состояния, приглушённая вторая строка.
// Чего нет. Кнопок: что делать со службой, решают на её странице настроек.
type RailStatus struct {
	Tone   RailStatusTone
	Text   string
	Detail Dim
}

// Render prints the note.
func (n RailStatus) Render() string { return n.html(nil) }

func (n RailStatus) html(m mix) string {
	dot := [...]DotTone{DotAccent, DotDanger, DotDusty}[n.Tone]
	return `<div` + class(m, "tc-railstatus", is("alert", n.Tone == RailStatusAlert), is("off", n.Tone == RailStatusOff)) + `>` +
		Dot{Tone: dot}.html(nil) + esc(n.Text) + `<br>` + n.Detail.html(nil) + `</div>`
}

// RailFootPart is what the rail's foot holds before the about line.
type RailFootPart interface {
	part
	railFootPart()
}

func (Button) railFootPart()     {}
func (Emblem) railFootPart()     {}
func (RailStatus) railFootPart() {}

// About is the line at the very end of the rail: the application and its
// build.
type About struct {
	Name  string
	Build Mono
	Href  string
}

// RailFoot — tc-railfoot
//
// Роль. Подвал рейла: кнопки «создать» и «подключить», состояние службы, сборка приложения.
// Вход. RailFoot{Items, About, Strip}.
// Состояния. plain | strip | about.
// Инварианты. Знак службы прижат вправо и не кнопка; сборка — последней строкой; подвал печатает фрагмент своего кадра.
// Поведение. нет: поведение у кнопок.
// Эталон. кадр `components-36`, фрагмент `.tc-railfoot`.
// Содержимое. Кнопки со значками, знак службы, строки о службах, строка сборки.
// Чего нет. На узком экране подвала нет: он скрыт вместе с секциями рейла.
type RailFoot struct {
	Items []RailFootPart
	About *About
	Strip bool // the rail's width with its rule on the right: the foot drawn without a rail
}

// Render prints the foot.
func (f RailFoot) Render() string { return f.html(nil) }

func (f RailFoot) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-railfoot", pick(f.Strip, "tc-railfoot--strip")) + `>`)
	for _, it := range f.Items {
		s.WriteString(it.html(nil))
	}
	if f.About != nil {
		s.WriteString(`<a class="tc-railfoot__about"` + attr("href", f.About.Href) + `>` + esc(f.About.Name) + ` ` + f.About.Build.html(nil) + `</a>`)
	}
	s.WriteString(`</div>`)
	return s.String()
}
