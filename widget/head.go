package widget

import "strings"

// HeadTitleSize is the size of a head's title where it is not the page's.
type HeadTitleSize int

const (
	HeadTitle22 HeadTitleSize = iota // a page
	HeadTitle19                      // a panel, a card
	HeadTitle18                      // a folder, a preview
	HeadTitle17                      // a file in a narrow panel
)

func (s HeadTitleSize) class() string {
	return [...]string{"", "tc-head__title--t19", "tc-head__title--t18", "tc-head__title--t17"}[s]
}

// HeadTitle is the title of a head: the source's mark, the words, a badge.
type HeadTitle struct {
	Mark  *Mark
	Text  string
	Badge *Badge
	Size  HeadTitleSize
}

func (t HeadTitle) html() string {
	s := `<div` + class(nil, "tc-head__title", t.Size.class()) + `>`
	if t.Mark != nil {
		s += t.Mark.html(nil)
	}
	s += esc(t.Text)
	if t.Badge != nil {
		s += " " + t.Badge.html(nil)
	}
	return s + `</div>`
}

// HeadSubPart is what a line under the title holds.
type HeadSubPart interface {
	part
	headSubPart()
}

func (Text) headSubPart()   {}
func (Strong) headSubPart() {}
func (Span) headSubPart()   {}
func (Mono) headSubPart()   {}
func (Mark) headSubPart()   {}
func (Badge) headSubPart()  {}
func (Button) headSubPart() {}

// HeadSub is a line under the title: where the record lives, what is in the
// page, what is wrong.
type HeadSub struct {
	Parts []HeadSubPart
}

func (b HeadSub) html(own ...string) string {
	var s strings.Builder
	s.WriteString(`<div` + class(nil, append([]string{"tc-head__sub"}, own...)...) + `>`)
	for _, p := range b.Parts {
		s.WriteString(p.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// HeadPhoto is the photo in front of a person's head: the picture's place,
// or initials while there is none.
type HeadPhoto struct {
	Initials string
	Large    bool // 76 px: a card
}

func (p HeadPhoto) html() string {
	return `<span` + class(nil, "tc-avatar--photo", is("ph", p.Initials != ""), pick(p.Large, "tc-avatar--photo--sq76-t20")) + `>` + esc(p.Initials) + `</span>`
}

// HeadPad is the head's padding where it is not the page's.
type HeadPad int

const (
	HeadPadPage     HeadPad = iota
	HeadPadTop11            // 11 px above: a head inside a component
	HeadPadBottom6          // 6 px below
	HeadPadBottom8          // 8 px below: a snug head over a list
	HeadPadBottom10         // 10 px below
	HeadRule                // a rule below: a head over a document
)

func (p HeadPad) class() string {
	return [...]string{"", "tc-head--padt11", "tc-head--padb6", "tc-head--snug", "tc-head--padb10", "tc-head--rule"}[p]
}

// HeadTool is a small button under a head's lines, or one with its menu.
type HeadTool interface {
	part
	headTool()
}

func (Button) headTool() {}
func (Drop) headTool()   {}

// Head — tc-head
//
// Роль. Шапка страницы или панели: заголовок, строки под ним, действия справа.
// Вход. Head{Photo, Avatar, Kicker, Title, Crumbs, Subs, Tools, Meta, Acts, Grouped, Middle, Pad}.
// Состояния. plain | acts | photo | avatar | crumbs | meta | middle | padded.
// Инварианты. Заголовок — один; строки под ним — после него, действия — в конце; отступы — перечисление; шапка печатает фрагмент своего кадра.
// Поведение. нет: поведение у действий.
// Эталон. кадр `s-contacts-1`, фрагмент `.tc-head`.
// Содержимое. Фотография или аватар, рубрика над заголовком, заголовок с меткой и бейджем, путь, строки под заголовком, строка кнопок, действия.
// Чего нет. Вкладок: разделы страницы — переключатель в действиях или бар под шапкой.
type Head struct {
	Photo   *HeadPhoto
	Avatar  *Avatar
	Kicker  *Rubric // a rubric over the title: what kind of page it is
	Title   *HeadTitle
	Crumbs  *Crumbs
	Subs    []HeadSub
	Tools   []HeadTool // small buttons under the lines: a panel's own actions, and its «⋯» with the menu
	Meta    *HeadSub   // a single line instead of the titles: a document's head
	Acts    *Acts
	Grouped bool // the photo and the titles side by side as one group
	Middle  bool // the actions centred on the titles, not on their top
	Pad     HeadPad
}

// Render prints the head.
func (h Head) Render() string { return h.html(nil) }

func (h Head) titles() string {
	var s strings.Builder
	if h.Kicker != nil {
		s.WriteString(h.Kicker.html(nil))
	}
	if h.Title != nil {
		s.WriteString(h.Title.html())
	}
	if h.Crumbs != nil {
		s.WriteString(h.Crumbs.html(nil))
	}
	for _, b := range h.Subs {
		s.WriteString(b.html())
	}
	if len(h.Tools) > 0 {
		s.WriteString(`<div class="tc-line tc-line--gap-5 tc-sec__acts">`)
		for _, b := range h.Tools {
			s.WriteString(b.html(nil))
		}
		s.WriteString(`</div>`)
	}
	return s.String()
}

func (h Head) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-head", pick(h.Middle, "tc-head--mid"), h.Pad.class()) + `>`)
	switch {
	case h.Meta != nil:
		s.WriteString(h.Meta.html("tc-head__sub--top0"))
	case h.Grouped:
		s.WriteString(`<div class="tc-head__group tc-head__group--ai-center">` + h.Photo.html() + `<div class="tc-head__titles">` + h.titles() + `</div></div>`)
	case h.Photo != nil:
		s.WriteString(`<div class="tc-head__titles tc-line tc-line--gap-14">` + h.Photo.html() + `<div class="tc-head__text">` + h.titles() + `</div></div>`)
	case h.Avatar != nil:
		s.WriteString(`<div class="tc-head__titles tc-head__titles--row">` + h.Avatar.html(nil) + `<div class="tc-head__text">` + h.titles() + `</div></div>`)
	default:
		s.WriteString(`<div class="tc-head__titles">` + h.titles() + `</div>`)
	}
	if h.Acts != nil {
		s.WriteString(h.Acts.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}
