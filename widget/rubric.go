package widget

import "strings"

// RubricForm is the element a rubric is.
type RubricForm int

const (
	RubricBlock  RubricForm = iota // a heading line
	RubricInline                   // a letter heading a list group
	RubricItem                     // a heading row of a menu: the menu prints it as its item
)

// RubricVariant is the drawn spacing and tone of a rubric.
type RubricVariant int

const (
	RubricPlain    RubricVariant = iota
	RubricBottom2                // 2 px below: over a title
	RubricBelow                  // the normal step below: over a list
	RubricRail                   // the padding of a rail section
	RubricSplit                  // words left, a count at the right edge, nudged
	RubricDanger                 // a server that does not answer
	RubricRow                    // a line: words, then items, 8 px apart
	RubricItemLine               // a menu's heading row: words, then a value at the edge
)

func (v RubricVariant) class() string {
	return [...]string{"", "tc-rubric--bot2", "tc-rubric--below-normal", "tc-rubric--pad0-14-6",
		"tc-rubric--split tc-rubric--nudge", "tc-rubric--danger", "tc-line tc-line--normal", "tc-line"}[v]
}

// RubricPart is what a rubric holds after its words.
type RubricPart interface {
	part
	rubricPart()
}

func (Mono) rubricPart()        {}
func (Badge) rubricPart()       {}
func (Num) rubricPart()         {}
func (Line) rubricPart()        {}
func (Push) rubricPart()        {}
func (RubricAside) rubricPart() {}
func (RubricCol) rubricPart()   {}
func (Button) rubricPart()      {}
func (Plain) rubricPart()       {}
func (Drop) rubricPart()        {}

// RubricAside is the quieter side of a rubric: what the heading's group is.
type RubricAside struct {
	Text  string
	Items []RubricAsidePart
}

// RubricAsidePart is what a rubric's aside holds.
type RubricAsidePart interface {
	part
	rubricAsidePart()
}

func (Badge) rubricAsidePart() {}
func (Mono) rubricAsidePart()  {}

func (a RubricAside) html(mix) string {
	var s strings.Builder
	s.WriteString(`<span class="tc-rubric__aside">` + esc(a.Text))
	for _, it := range a.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</span>`)
	return s.String()
}

// RubricCol is the heading of a column under a menu's heading row.
type RubricCol struct {
	Text string
}

func (c RubricCol) html(mix) string { return `<span class="tc-rubric__col">` + esc(c.Text) + `</span>` }

// Rubric — tc-rubric
//
// Роль. Рубрика: имя группы капителью — над полями, в рейле, над группой списка, в строке меню.
// Вход. Rubric{Text, Lead, Items, Form, Variant}.
// Состояния. block | inline | item | aside | variant | row.
// Инварианты. Рубрика — не заголовок документа: она называет группу, а не страницу; кнопка в рубрике — обычным регистром (Р-12); рубрика печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `shell-1`, фрагмент `.tc-rubric`.
// Содержимое. Слова, моноширинное перед словами (Lead) или после, число, бейдж, приглушённая сторона, строка кнопок, кнопка с меню (Drop), прижатая группа.
// Чего нет. Уровней: рубрика одна, без h1–h6.
type Rubric struct {
	Text    string
	Lead    *Mono // a monospaced piece before the words: a property of a conflict
	Items   []RubricPart
	Form    RubricForm
	Variant RubricVariant
}

// Render prints the rubric.
func (r Rubric) Render() string { return r.html(nil) }

func (r Rubric) html(m mix) string {
	tag := [...]string{"div", "span", "div"}[r.Form]
	extra := ""
	if r.Form == RubricItem {
		extra = ` role="presentation"`
	}
	var s strings.Builder
	s.WriteString(`<` + tag + class(m, "tc-rubric", r.Variant.class()) + extra + `>`)
	if r.Lead != nil {
		s.WriteString(r.Lead.html(nil) + " ")
	}
	s.WriteString(esc(r.Text))
	for _, it := range r.Items {
		s.WriteString(" " + it.html(nil))
	}
	s.WriteString(`</` + tag + `>`)
	return s.String()
}

func (Rubric) fieldSetRubric() {}
