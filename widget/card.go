package widget

import "strings"

// Key is a word in bold that names a setting inside a card's text.
type Key string

func (k Key) html(mix) string { return `<b class="tc-card__key">` + esc(string(k)) + `</b>` }

// CardTextPart is what a card's text holds.
type CardTextPart interface {
	part
	cardTextPart()
}

func (Text) cardTextPart()   {}
func (Strong) cardTextPart() {}
func (Bold) cardTextPart()   {}
func (Mono) cardTextPart()   {}
func (Span) cardTextPart()   {}
func (Link) cardTextPart()   {}
func (Dim) cardTextPart()    {}
func (Break) cardTextPart()  {}
func (Key) cardTextPart()    {}
func (Button) cardTextPart() {}

// CardTextKind is which line of text a card holds.
type CardTextKind int

const (
	CardLine       CardTextKind = iota // a record and where it lives
	CardNum                            // a big number: how many were created
	CardTerm                           // the name of a way out, in bold
	CardUsage                          // how to use what the card gives
	CardSmallPrint                     // a dusty note under the rows
	CardWhy                            // why the card is not offered
	CardWarn                           // what goes wrong if ignored
	CardAlertText                      // what an alert card means
	CardLede                           // the first sentence of a setting
	CardCaption                        // what follows a letter's attachments
	CardAbove                          // a line over the card's content
)

func (k CardTextKind) class() string {
	return [...]string{"tc-card__line", "tc-card__num", "tc-card__term", "tc-card__note", "tc-card__note-dusty", "tc-card__why", "tc-card__warn",
		"tc-card__alert-note", "tc-card__lede", "tc-card__caption", "tc-card__above"}[k]
}

// CardText is a line of text in a card.
type CardText struct {
	Kind  CardTextKind
	Parts []CardTextPart
}

func (t CardText) html(mix) string {
	return `<div class="` + t.Kind.class() + `">` + joined(t.Parts) + `</div>`
}

// CardHint is an explanation in a card: under a field (a line) or among
// the card's pieces (a span).
type CardHint struct {
	Parts []CardTextPart
	Line  bool // a line of its own, 8 px above
}

func (h CardHint) html(mix) string {
	if h.Line {
		return `<div class="tc-field__hint tc-field__hint--top8">` + joined(h.Parts) + `</div>`
	}
	return `<span class="tc-field__hint">` + joined(h.Parts) + `</span>`
}

// CardRowKind is which row of a card it is: where it stands and what it is
// for.
type CardRowKind int

const (
	CardActs     CardRowKind = iota // the card's buttons
	CardTop2                        // 2 px above: a key and its button
	CardTop8                        // 8 px above
	CardTop9                        // 9 px above
	CardTop10                       // 10 px above
	CardTop11                       // 11 px above: fields side by side
	CardAttached                    // a letter's attachments
	CardFoot                        // a note at the card's foot
)

func (k CardRowKind) class() string {
	return [...]string{"tc-card__acts", "tc-card__top2", "tc-card__top8", "tc-card__top9", "tc-card__top10", "tc-card__top11", "tc-card__files", "tc-card__foot"}[k]
}

// CardRowPart is what a card's row holds.
type CardRowPart interface {
	part
	cardRowPart()
}

func (Button) cardRowPart()   {}
func (Input) cardRowPart()    {}
func (Field) cardRowPart()    {}
func (Tag) cardRowPart()      {}
func (Dim) cardRowPart()      {}
func (Span) cardRowPart()     {}
func (FillHint) cardRowPart() {}
func (Turn) cardRowPart()     {}

// FillHint is an explanation that takes the rest of a card's row.
type FillHint struct {
	Parts []CardTextPart
}

func (h FillHint) html(mix) string {
	return `<span class="tc-field__hint tc-field__hint--flush-fill">` + joined(h.Parts) + `</span>`
}

// Turn is the arrow at a card's foot that points back up: this follows from
// the above.
type Turn struct{}

func (Turn) html(mix) string {
	return `<svg width="12" height="12" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2" fill="none" class="tc-ico--turn">` + Icon("chev").use() + `</svg>`
}

// CardRow is a row of a card.
type CardRow struct {
	Kind  CardRowKind
	Gap   int // the gap of the row's line, in px; 0: not a line
	Items []CardRowPart
}

func (r CardRow) html(mix) string {
	var s strings.Builder
	line := ""
	if r.Gap > 0 {
		line = "tc-line tc-line--gap-" + itoa(r.Gap)
	}
	s.WriteString(`<div` + class(nil, r.Kind.class(), line) + `>`)
	for _, it := range r.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// CardList is a short list in a card: what will happen.
type CardList struct {
	Items []string
}

func (l CardList) html(mix) string {
	var s strings.Builder
	s.WriteString(`<ul class="tc-stack tc-stack--gap-4 tc-card__list">`)
	for _, it := range l.Items {
		s.WriteString(`<li>` + esc(it) + `</li>`)
	}
	s.WriteString(`</ul>`)
	return s.String()
}

// CardFieldsPart is what a card's row of fields holds.
type CardFieldsPart interface {
	part
	cardFieldsPart()
}

func (Field) cardFieldsPart()  {}
func (Button) cardFieldsPart() {}
func (Span) cardFieldsPart()   {}

// CardFields is a row of fields in a card that wraps.
type CardFields struct {
	Items []CardFieldsPart
	Wide  bool // 10/12 px gaps
	Rule  bool // a rule above: another way to do it
}

func (f CardFields) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(nil, "tc-card__fields", pick(f.Wide, "tc-card__fields--gap10-12"), pick(f.Rule, "tc-card__fields--rule")) + `>`)
	for _, it := range f.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// CardQRPart is what stands beside a code.
type CardQRPart interface {
	part
	cardQRPart()
}

func (Dim) cardQRPart()    {}
func (Button) cardQRPart() {}

// CardQR is a code to scan and what to do with it.
type CardQR struct {
	Colour string // the code's paper
	Symbol string // the code in the sprite, drawn in a 27 box
	Side   []CardQRPart
	Src    string // the code as a picture from the server, instead of the sprite
}

func (q CardQR) html(mix) string {
	var s strings.Builder
	if q.Src != "" {
		s.WriteString(`<div class="tc-card__split"><div` + attr("data-tc-colour", q.Colour) + ` class="tc-card__qr"><img class="tc-card__qr-img"` + attr("src", q.Src) +
			` width="104" height="104" alt="QR"></div><div class="tc-stack tc-stack--gap-8 tc-card__qr-side">`)
	} else {
		s.WriteString(`<div class="tc-card__split"><div` + attr("data-tc-colour", q.Colour) + ` class="tc-card__qr"><svg width="104" height="104" viewBox="-1 -1 27 27" role="img" aria-label="QR">` +
			`<rect x="-1" y="-1" width="27" height="27"` + attr("fill", q.Colour) + `/><use href="#` + esc(q.Symbol) + `"/></svg></div><div class="tc-stack tc-stack--gap-8 tc-card__qr-side">`)
	}
	for _, it := range q.Side {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div></div>`)
	return s.String()
}

// CardPick is a choice a card offers: a tick box, a title, a note, an aside
// or a field.
type CardPick struct {
	Ticked Ticked
	Send   Send // what pressing the tick sends
	Title  string
	Note   []CardTextPart
	Aside  string
	Field  *Field
}

func (p CardPick) html(mix) string {
	s := `<div class="tc-card__pick">` + p.Ticked.box(p.Title, p.Send, "tc-tickrow__tick--top3") + `<div class="tc-stack__item"><div class="tc-stack__item-title">` + esc(p.Title) +
		`</div><div class="tc-stack__item-note">` + joined(p.Note) + `</div>`
	if p.Aside != "" {
		s += `<div class="tc-stack__item-aside">` + esc(p.Aside) + `</div>`
	}
	if p.Field != nil {
		s += p.Field.html(nil)
	}
	return s + `</div></div>`
}

// CardGridPart is what a card's grid holds.
type CardGridPart interface {
	part
	cardGridPart()
}

func (Dim) cardGridPart()    {}
func (Input) cardGridPart()  {}
func (Area) cardGridPart()   {}
func (Button) cardGridPart() {}

// CardGrid lays out names, values and their buttons in columns.
type CardGrid struct {
	Items []CardGridPart
}

func (g CardGrid) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-card__grid">`)
	for _, it := range g.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// CardRowsItem is a line of a card's rows: a monospaced head and what it
// means.
type CardRowsItem struct {
	Head Mono
	Note Dim
}

// CardRows is what a key opens, one request a line.
type CardRows struct {
	Items []CardRowsItem
}

func (r CardRows) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-card__rows">`)
	for _, it := range r.Items {
		s.WriteString(`<span class="tc-stack__item">` + it.Head.html(nil) + `<br>` + it.Note.html(nil) + `</span>`)
	}
	s.WriteString(`</div>`)
	return s.String()
}

// CardHeadLine is a line of a letter's head: its slot and what is in it.
type CardHeadLine struct {
	Slot  string
	Parts []CardTextPart
}

// CardHead is the head of a letter shown in a card: from, reply-to,
// subject.
type CardHead struct {
	Lines []CardHeadLine
}

func (h CardHead) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-stack tc-stack--gap-3 tc-card__head">`)
	for _, l := range h.Lines {
		s.WriteString(`<div class="tc-stack__item">` + Dim{Parts: []Inline{Text(l.Slot)}, Variant: DimSlot}.html(nil) + joined(l.Parts) + `</div>`)
	}
	s.WriteString(`</div>`)
	return s.String()
}

// SpacedField is a field set 9 px under what is above it in a card.
type SpacedField struct {
	Field Field
}

func (f SpacedField) html(mix) string { return f.Field.html(mix{"tc-card__top9"}) }

// CardPart is what a card holds.
type CardPart interface {
	part
	cardPart()
}

func (Rubric) cardPart()      {}
func (List) cardPart()        {}
func (Tree) cardPart()        {}
func (Seg) cardPart()         {}
func (Diag) cardPart()        {}
func (Mono) cardPart()        {}
func (Line) cardPart()        {}
func (Msg) cardPart()         {}
func (Prose) cardPart()       {}
func (Table) cardPart()       {}
func (Check) cardPart()       {}
func (CardText) cardPart()    {}
func (CardHint) cardPart()    {}
func (CardRow) cardPart()     {}
func (CardList) cardPart()    {}
func (CardFields) cardPart()  {}
func (CardQR) cardPart()      {}
func (CardPick) cardPart()    {}
func (CardGrid) cardPart()    {}
func (CardRows) cardPart()    {}
func (CardHead) cardPart()    {}
func (SpacedField) cardPart() {}
func (Field) cardPart()       {}
func (Button) cardPart()      {}

// CardVariant is the drawn form of a card.
type CardVariant int

const (
	CardPlain           CardVariant = iota
	CardAlert                       // a server that refused
	CardApart                       // 14 px above
	CardSnug                        // 12/14 px inside: a way out among others
	CardFlush                       // no margin: set in a column of its own
	CardChosen                      // an accent edge: the chosen one
	CardChosenFlush                 // an accent edge, no margin
	CardChosenFlushGrow             // an accent edge, flush, both classes: a card that grows in a row
)

func (v CardVariant) class() string {
	return [...]string{"", "tc-card--alert", "tc-card--top14", "tc-card--pad12-14", "tc-card--flush", "tc-card--edge-accent",
		"tc-card--edge-accent-flush", "tc-card--edge-accent tc-card--flush"}[v]
}

// Card — tc-card
//
// Роль. Карточка: подключение, кандидат в дубли, путь из мастера, ключ для вставки, письмо, выбор режима.
// Вход. Card{Items, Variant, Place}.
// Состояния. plain | alert | apart | snug | flush | chosen | pick.
// Инварианты. Вид — перечисление; части карточки — именованные, а не произвольная разметка; карточка печатает фрагмент своего кадра.
// Поведение. нет: поведение у того, что в ней.
// Эталон. кадр `s-dup-1`, фрагмент `.tc-card`.
// Содержимое. Рубрика, список, дерево, переключатель, диагностика, строки текста, подсказки, ряды кнопок и полей, поля, код, выбор, сетка, строки запросов, заголовок письма, сообщение, абзац, таблица, флажок.
// Чего нет. Вложенных карточек.
type Card struct {
	Items   []CardPart
	Variant CardVariant
	Place   Place
}

// Render prints the card.
func (c Card) Render() string { return c.html(nil) }

func (c Card) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-card", c.Place.class(), c.Variant.class()) + `>`)
	for _, it := range c.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// ResultsPart is what a page's column of results holds.
type ResultsPart interface {
	part
	resultsPart()
}

func (Card) resultsPart()   {}
func (Rubric) resultsPart() {}

// MainResults is a column of cards in a page: candidates, servers.
type MainResults struct {
	Items []ResultsPart
}

func (r MainResults) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-stack tc-stack--gap-10 tc-main__results">`)
	for _, it := range r.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// MainCards is a row of cards in a page that wraps: ways to add a user.
type MainCards struct {
	Cards []Card
}

func (c MainCards) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-main__cards">`)
	for _, it := range c.Cards {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// MainDiag is diagnostics set in a page with its gutters.
type MainDiag struct {
	Diag Diag
}

func (d MainDiag) html(mix) string {
	return `<div class="tc-main__diag">` + d.Diag.html(nil) + `</div>`
}

// MainThumbs is a public page's picture.
type MainThumbs struct {
	Thumb Thumb
}

func (t MainThumbs) html(mix) string {
	return `<div class="tc-main__thumbs">` + t.Thumb.html(nil) + `</div>`
}

func (MainResults) mainPart() {}
func (MainCards) mainPart()   {}
func (MainDiag) mainPart()    {}
func (MainThumbs) mainPart()  {}
func (Danger) mainPart()      {}
func (Empty) mainPart()       {}
func (Fields) mainPart()      {}
func (Grid) mainPart()        {}
func (Msg) mainPart()         {}
func (Pad) mainPart()         {}
func (Queue) mainPart()       {}
func (Runbar) mainPart()      {}
func (Sec) mainPart()         {}
func (Table) mainPart()       {}

func (Danger) sidePart() {}
func (Fields) sidePart() {}
func (Msg) sidePart()    {}
func (Pad) sidePart()    {}
func (Sec) sidePart()    {}

func (Card) dialogPart()   {}
func (Diff) dialogPart()   {}
func (Fields) dialogPart() {}
func (Msg) dialogPart()    {}
func (Stats) dialogPart()  {}
func (Steps) dialogPart()  {}
func (Table) dialogPart()  {}

func (Card) sheetPart()  {}
func (Diag) sheetPart()  {}
func (Empty) sheetPart() {}
func (Msg) sheetPart()   {}
func (Sec) sheetPart()   {}
func (Table) sheetPart() {}

func (Card) formPart() {}
func (Msg) formPart()  {}

func (Card) fieldSetPart()   {}
func (Diag) fieldSetPart()   {}
func (Fields) fieldSetPart() {}
func (Msg) fieldSetPart()    {}
func (Table) fieldSetPart()  {}

func (Diag) fieldPart() {}
func (Msg) fieldPart()  {}

func (Card) stackPart()        {}
func (Empty) scrimPart()       {}
func (Empty) overlayRailPart() {}
func (Field) secPart()         {}
func (Diag) secPart()          {}
func (Meter) linePart()        {}

func (Card) formColumnPart() {}
func (Msg) formColumnPart()  {}
