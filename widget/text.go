package widget

import "strings"

// Text is plain running text inside a text block.
type Text string

func (t Text) html(mix) string { return esc(string(t)) }

// Strong is a bold piece of running text: a name in a sentence.
type Strong string

func (s Strong) html(mix) string { return `<b>` + esc(string(s)) + `</b>` }

// Em is italic words in a document's text: Markdown's *emphasis* (Carrel's
// task 102).
type Em string

func (e Em) html(mix) string { return `<em>` + esc(string(e)) + `</em>` }

// Hint is a word with an explanation in a hint: «verified», with what was
// verified.
type Hint struct {
	Text  string
	Title string
}

func (h Hint) html(mix) string {
	return `<span class="tc-dim__hint"` + attr("title", h.Title) + `>` + esc(h.Text) + `</span>`
}

// DotTone is the state a dot shows.
type DotTone int

const (
	DotAccent DotTone = iota // running as it should
	DotDanger                // failing
	DotDusty                 // off
)

// Dot is the state dot before a line of status text.
type Dot struct {
	Tone DotTone
}

func (d Dot) html(mix) string {
	return `<span class="` + [...]string{"tc-dot--accent", "tc-dot--danger", "tc-dot--dusty"}[d.Tone] + `"></span>`
}

// Inline is running text: plain text, bold, a monospaced piece, a hint, a
// state dot, a piece in a face of its own. A dim line, a field's hint and a
// check's label are made of it.
type Inline interface {
	part
	inline()
}

func (Text) inline()    {}
func (Strong) inline()  {}
func (Em) inline()      {}
func (Hint) inline()    {}
func (Dot) inline()     {}
func (Mono) inline()    {}
func (Span) inline()    {}
func (Mention) inline() {}

// Face is a face without a block of its own: a tone or a size of dim text
// set by its variant alone («another tone is the variant instead of the
// base», docs/visual/mockup.md).
type Face int

const (
	FaceMuted      Face = iota // tc-dim--muted: typed text, still being written
	FaceAccent                 // tc-dim--accent: a good outcome
	FaceDanger                 // tc-dim--danger: a failure
	FaceAccent12               // accent, 12 px
	FaceAccent11_5             // accent, 11.5 px
	FaceDanger11_5             // danger, 11.5 px
	FaceLarge                  // tc-dim--t12-5: a sentence a little larger
	FaceSmall                  // tc-dim--t12: a condition under a drop zone
	FaceMuted12                // muted, 12 px: what a field in a card is for
)

func (f Face) class() string {
	return [...]string{"tc-dim--muted", "tc-dim--accent", "tc-dim--danger", "tc-dim--accent tc-dim--t12",
		"tc-dim--accent tc-dim--t11-5", "tc-dim--danger tc-dim--t11-5", "tc-dim--t12-5", "tc-dim--t12", "tc-dim--muted tc-dim--t12"}[f]
}

// Mention is a mention of a card whose card is not in the books a person
// looks in (Mentions, approved and deferred): the name the text keeps, not a
// link, with a hint that says why. A mention whose card is found is a Link.
type Mention struct {
	Name string // as the text writes it, with its @
}

func (x Mention) html(mix) string {
	return `<span class="tc-mention is-lost" title="This card is not in the books you look in">` + esc(x.Name) + `</span>`
}

// Span is words in a face of their own.
type Span struct {
	Face Face
	Text string
	Href string // words in the accent that go somewhere: a link
	// Title is what the words say in full when they are short: «1 not
	// copied» names the collection and why in its hint.
	Title string
}

func (s Span) html(mix) string {
	if s.Href != "" {
		return `<a class="` + s.Face.class() + `"` + attr("href", s.Href) + `>` + esc(s.Text) + `</a>`
	}
	return `<span class="` + s.Face.class() + `"` + attr("title", s.Title) + `>` + esc(s.Text) + `</span>`
}

// DimSize is the size of a dim line where it is not the running text's.
type DimSize int

const (
	DimText DimSize = iota
	Dim10_5
	Dim11
	Dim11_5
)

func (s DimSize) class() string {
	return [...]string{"", "tc-dim--t10-5", "tc-dim--t11", "tc-dim--t11-5"}[s]
}

// DimVariant is how a dim line sits in its row.
type DimVariant int

const (
	DimInline DimVariant = iota
	DimWrap              // wraps instead of cutting: a long context line
	DimTight             // no letter spacing: a source name in a list
	DimIndent            // 6 px from what it follows
	DimTop               // pinned to the top of its row
	DimSlot              // the name of a slot in a letter: From, To
	DimStruck            // struck through: a done item
	DimFixed             // a fixed-width label column
)

func (v DimVariant) class() string {
	return [...]string{"", "tc-dim--wrap", "tc-dim--ls0", "tc-dim--ml6", "tc-dim--start", "tc-dim--slot",
		"tc-dim--struck", "tc-dim--fix"}[v]
}

// Dim — tc-dim
//
// Роль. Приглушённый текст рядом с главным: «· in 7 months», «412 KB», имя источника.
// Вход. Dim{Parts, Size, Variant, Place, Empty}.
// Состояния. plain | sized | wrap | tight | indent | top | slot | struck | fixed | mixed.
// Инварианты. Тон один — приглушённый, другой тон — не этот блок; размер и посадка — перечисления; строка печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `s-contacts-6`, фрагмент `.tc-dim`.
// Содержимое. Текст, жирное, моноширинный кусок, подсказка, точка состояния.
// Чего нет. Ссылок и кнопок: приглушённое не нажимают.
type Dim struct {
	Parts   []Inline
	Size    DimSize
	Variant DimVariant
	Place   Place
	// Empty is the words of an empty list: «Nothing attached»; the library's
	// script hides them while a picked file stands in the list.
	Empty bool
}

// Render prints the dim line.
func (d Dim) Render() string { return d.html(nil) }

func (d Dim) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<span` + class(m, "tc-dim", d.Size.class(), d.Variant.class(), d.Place.class()) + flag("data-tc-empty", d.Empty) + `>`)
	for _, p := range d.Parts {
		s.WriteString(p.html(nil))
	}
	s.WriteString(`</span>`)
	return s.String()
}

// MonoForm is the element a monospaced piece is.
type MonoForm int

const (
	MonoSpan MonoForm = iota // inside a line
	MonoLine                 // a line of its own
	MonoCode                 // a word of a protocol or a command
)

// MonoVariant is the drawn variant of a monospaced piece.
type MonoVariant int

const (
	MonoDefault MonoVariant = iota
	Mono11
	MonoTight        // no letter spacing: a count
	Mono10Dusty      // 10 px, dusty: "last used"
	Mono11_5Dusty    // 11.5 px, dusty: a server address line
	Mono11_5DustyGap // the same with 9 px below
	MonoMuted        // muted: a size
	MonoTightDusty   // no spacing, dusty: points of a pair
	MonoDusty11      // dusty, 11 px: a count on a button
	Mono11Dusty      // 11 px, dusty: a zoom
	MonoDusty        // dusty: a size in a queue
	Mono11DustyTight // 11 px, dusty, no spacing: found and when
	MonoTightDanger  // no spacing, red: "visible only now"
	MonoFace         // the monospaced face only: a time in print
	Mono11Top        // 11 px, 4 px above: a measure under a frame
	Mono11Indent     // 11 px, 4 px before: a count after marks
	MonoTightPlain   // no spacing, no capitals: a host name in a rubric
	Mono11Muted      // 11 px, muted: an address in angle brackets
	MonoChip         // on a chip: a command
)

func (v MonoVariant) class() string {
	return [...]string{"", "tc-mono--t11", "tc-mono--ls0", "tc-mono--t10-dusty", "tc-mono--t11-5-dusty",
		"tc-mono--t11-5-dusty tc-mono--bot9", "tc-mono--muted", "tc-mono--ls0-dusty", "tc-mono--dusty-t11",
		"tc-mono--t11-dusty", "tc-mono--dusty", "tc-mono--t11-dusty-ls0", "tc-mono--ls0-danger", "tc-mono--f-mono",
		"tc-mono--t11-top4", "tc-mono--t11-ml4", "tc-mono--ls0-tt-none", "tc-mono--t11-muted", "tc-mono--chip"}[v]
}

// Say — tc-say
//
// Роль. Слова в двух длинах: полные и короткие для узкого экрана — сказать короче, а не сжать (Р-65); смысл и числа в обеих те же.
// Вход. Say{Full, Short, Mono}.
// Состояния. plain | mono.
// Инварианты. Видна одна длина: короткая — где колонка не шире 640 px; кусок печатает фрагмент своего кадра.
// Поведение. нет: длину выбирает CSS по ширине.
// Эталон. кадр `s-find-1`, фрагмент `.tc-say`.
// Содержимое. Два текста.
// Чего нет. Третьей длины и сценария.
type Say struct {
	Full  string
	Short string
	Mono  bool // monospaced: a time, a count
}

// Render prints the words.
func (s Say) Render() string { return s.html(nil) }

func (s Say) html(m mix) string {
	return `<span` + class(m, "tc-say", pick(s.Mono, "tc-say--mono")) + `><span class="tc-say__full">` + esc(s.Full) +
		`</span><span class="tc-say__short">` + esc(s.Short) + `</span></span>`
}

// Mono — tc-mono
//
// Роль. Моноширинный кусок: время, размер, адрес, идентификатор, команда.
// Вход. Mono{Text, Form, Variant, Place}.
// Состояния. span | line | code | variant.
// Инварианты. Элемент и вариант — перечисления; кусок печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `components-16`, фрагмент `.tc-mono`.
// Содержимое. Только текст.
// Чего нет. Разметки внутри: адрес и команда переносятся как есть.
type Mono struct {
	Text    string
	Form    MonoForm
	Variant MonoVariant
	Place   Place
}

// Render prints the monospaced piece.
func (o Mono) Render() string { return o.html(nil) }

func (o Mono) html(m mix) string {
	tag := [...]string{"span", "div", "code"}[o.Form]
	return `<` + tag + class(m, "tc-mono", o.Variant.class(), o.Place.class()) + `>` + esc(o.Text) + `</` + tag + `>`
}

// Plain — tc-plain
//
// Роль. Пояснение обычным регистром внутри заголовка, который кричит: «2 active» в рубрике.
// Вход. Plain{Text}.
// Состояния. нет.
// Инварианты. Строчные и без разрядки, какой бы ни была рубрика вокруг; пояснение печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `account-1`, фрагмент `.tc-plain`.
// Содержимое. Только текст.
// Чего нет. Своего размера и цвета: их даёт заголовок.
type Plain struct {
	Text string
}

// Render prints the plain aside.
func (p Plain) Render() string { return p.html(nil) }

func (p Plain) html(m mix) string {
	return `<span` + class(m, "tc-plain") + `>` + esc(p.Text) + `</span>`
}
