package widget

import "strings"

// Link is a link inside running text, in the accent.
type Link struct {
	Text string
	Href string
}

func (l Link) html(mix) string {
	return `<a class="tc-dim--accent"` + attr("href", l.Href) + `>` + esc(l.Text) + `</a>`
}

func (Link) inline() {}

// ProseVariant is the drawn size, tone and spacing of a paragraph.
type ProseVariant int

const (
	ProsePlain      ProseVariant = iota
	ProseMuted                   // a quieter paragraph: what a dialog means
	ProseMutedClose              // muted, flush, 10 px below: over a button in a card
	Prose13                      // 13 px: an event's description
	Prose13Card                  // 13 px, flush, 11 px above: a card's text
	ProseLead                    // the first paragraph of a document, 14 px above
	ProseLeadAfter               // a paragraph after a section, 18 px above
	ProseRead                    // the reading measure
	ProseReadTop                 // the reading measure, 14 px above
	ProseTop                     // 14 px above
)

func (v ProseVariant) class() string {
	return [...]string{"", "tc-prose--muted", "tc-prose--muted-flush-bot10", "tc-prose--t13", "tc-prose--t13-flush-padt11",
		"tc-prose--lead tc-prose--top14", "tc-prose--lead tc-prose--top18", "tc-prose--read", "tc-prose--read tc-prose--top14", "tc-prose--top14"}[v]
}

// Prose — tc-prose
//
// Роль. Абзац текста: описание записи, начало документа, пояснение диалога.
// Вход. Prose{Parts, Variant}.
// Состояния. plain | muted | size | lead | read | top | rich.
// Инварианты. Текст переносится по словам и держит свою меру; вид — перечисление; абзац печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `s-notes-2`, фрагмент `.tc-prose`.
// Содержимое. Текст, жирное, моноширинное, ссылка.
// Чего нет. Разметки заголовков и списков: у документа для них свои части.
type Prose struct {
	Parts   []Inline
	Variant ProseVariant
}

// Render prints the paragraph.
func (p Prose) Render() string { return p.html(nil) }

func (p Prose) html(m mix) string {
	return `<div` + class(m, "tc-prose", p.Variant.class()) + `>` + inlines(p.Parts) + `</div>`
}

// DocTitle is a document's title as read.
type DocTitle struct {
	Text string
}

func (t DocTitle) html(mix) string { return `<div class="tc-doc__title">` + esc(t.Text) + `</div>` }

// DocSubPart is what the line under a document's title holds.
type DocSubPart interface {
	part
	docSubPart()
}

func (Text) docSubPart() {}
func (Mono) docSubPart() {}
func (Tag) docSubPart()  {}

// DocSub is the line under a document's title: when it was written, its tags.
type DocSub struct {
	Parts []DocSubPart
}

func (b DocSub) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-head__sub tc-head__sub--top6">`)
	for _, p := range b.Parts {
		s.WriteString(p.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// DocTask is a task line of a document: a tick and what to do.
type DocTask struct {
	Tick Tick // Done says whether it is done
	Text string
}

// DocTasks is a list of tasks in a document under its heading: the
// author's Markdown heading, drawn as text, not as the interface's rubric
// (Carrel's task 196).
type DocTasks struct {
	Title string
	Items []DocTask
	Phone bool // a phone's list: wider gaps
}

func (t DocTasks) html(mix) string {
	var s strings.Builder
	if t.Phone {
		s.WriteString(`<div class="tc-doc__acts">` + DocHeading{Text: t.Title}.html(nil) + `<ul class="tc-stack tc-stack--gap-9 tc-doc__list">`)
	} else {
		s.WriteString(`<div class="tc-doc__sec">` + DocHeading{Text: t.Title}.html(nil) + `<ul class="tc-stack tc-stack--gap-7 tc-doc__list">`)
	}
	for _, it := range t.Items {
		s.WriteString(`<li class="tc-doc__task">` + it.Tick.html(nil))
		if it.Tick.Done {
			s.WriteString(Dim{Parts: []Inline{Text(it.Text)}, Variant: DimStruck}.html(nil))
		} else {
			s.WriteString(esc(it.Text))
		}
		s.WriteString(`</li>`)
	}
	s.WriteString(`</ul></div>`)
	return s.String()
}

// DocFile is an attachment named in a document: the name and its size.
type DocFile struct {
	Name string
	Size Dim
	Href string // where the file is fetched
}

// DocFiles is the attachments of a document under a rule.
type DocFiles struct {
	Title Rubric
	Items []DocFile
}

func (f DocFiles) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-doc__sec--rule">` + f.Title.html(nil) + `<div class="tc-stack tc-stack--gap-5 tc-doc__files">`)
	for _, it := range f.Items {
		name := `<span class="tc-dim--accent">` + esc(it.Name) + `</span>`
		if it.Href != "" {
			name = Link{Text: it.Name, Href: it.Href}.html(nil)
		}
		s.WriteString(`<span class="tc-stack__item">` + name + ` ` + it.Size.html(nil) + `</span>`)
	}
	s.WriteString(`</div></div>`)
	return s.String()
}

// DocRelated is what a document is related to, under a rule: a phone's note.
type DocRelated struct {
	Title Rubric
	Name  string
	What  Dim
	Href  string // the related record's address
}

func (r DocRelated) html(mix) string {
	tag := "div"
	if r.Href != "" {
		tag = "a"
	}
	return `<div class="tc-doc__sec--rule16">` + r.Title.html(nil) + `<` + tag + ` class="tc-doc__link"` + attr("href", r.Href) + `>` + esc(r.Name) + ` ` + r.What.html(nil) + `</` + tag + `></div>`
}

// DocFields is the row of a document's fields in the editor.
type DocFields struct {
	Items []Field
}

func (f DocFields) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-doc__fields">`)
	for _, it := range f.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// DocNote is the small line at the end of the editor: what is kept unseen.
type DocNote struct {
	Parts []Inline
}

func (n DocNote) html(mix) string { return `<div class="tc-doc__note">` + inlines(n.Parts) + `</div>` }

// DocHeading is a heading inside a document's text.
type DocHeading struct {
	Text string
}

func (h DocHeading) html(mix) string { return `<div class="tc-doc__heading">` + esc(h.Text) + `</div>` }

// DocBullets is a bulleted list inside a document's text. An item's parts
// may end with a DocBullets: the list nested in it (Carrel's task 102).
type DocBullets struct {
	Items [][]Inline
}

func (b DocBullets) html(mix) string {
	var s strings.Builder
	s.WriteString(`<ul class="tc-stack tc-stack--gap-6 tc-doc__bullets">`)
	for _, it := range b.Items {
		s.WriteString(`<li>` + inlines(it) + `</li>`)
	}
	s.WriteString(`</ul>`)
	return s.String()
}

func (DocBullets) inline() {}

// DocQuote is a quotation in a document's text: Markdown's «>» (Carrel's
// task 102).
type DocQuote struct {
	Parts []Inline
}

func (q DocQuote) html(mix) string {
	return `<blockquote class="tc-doc__quote">` + inlines(q.Parts) + `</blockquote>`
}

// DocCode is a block of code in a document's text, as written: Markdown's
// fenced or indented code (Carrel's task 102). A line too long for the column
// scrolls inside the block.
type DocCode struct {
	Text string
	// Wrap wraps the lines at the column: a note's source, written as prose
	// (Carrel's task 104), not code.
	Wrap bool
}

func (c DocCode) html(mix) string {
	return `<pre` + class(nil, "tc-doc__code", is("wrap", c.Wrap)) + `>` + esc(c.Text) + `</pre>`
}

// DocBody is a plain paragraph inside a document's text.
type DocBody struct {
	Text string
}

func (b DocBody) html(mix) string { return `<div class="tc-doc__body">` + esc(b.Text) + `</div>` }

// DocAlign is how a column of a table in a document's text is aligned.
type DocAlign int

const (
	DocAlignStart  DocAlign = iota // the column as written
	DocAlignCenter                 // «:---:»
	DocAlignEnd                    // «---:», numbers
)

func (a DocAlign) class() string { return [...]string{"", "is-center", "is-end"}[a] }

// DocTable is a table in a document's text: a head, rows, how each column is
// aligned. Words wrap between words, never inside one, and links and code do
// not wrap at all; what does not fit scrolls inside the table. In a container
// of 640 px and less Switch stands over it — Wrap / One line — and One line
// keeps each row to one line.
type DocTable struct {
	Head   [][]Inline
	Rows   [][][]Inline
	Align  []DocAlign // per column; a missing one is DocAlignStart
	Switch Seg        // Wrap / One line; its second option on is one line
}

// DocTableSwitch is the switch over a table: Wrap, or One line when line.
func DocTableSwitch(line bool) Seg {
	return Seg{Options: []SegOption{{Label: "Wrap", On: !line}, {Label: "One line", On: line}}, Variant: SegCompactSmall}
}

func (t DocTable) align(i int) string {
	if i < len(t.Align) {
		return t.Align[i].class()
	}
	return ""
}

func (t DocTable) html(mix) string {
	line := len(t.Switch.Options) > 1 && t.Switch.Options[1].On
	var s strings.Builder
	s.WriteString(`<div` + class(nil, "tc-doc__table", is("line", line)) + `>`)
	s.WriteString(`<div class="tc-doc__tablebar">` + t.Switch.html(nil) + `</div>`)
	s.WriteString(`<div class="tc-doc__scroll"><table class="tc-doc__grid"><thead><tr>`)
	for i, c := range t.Head {
		s.WriteString(`<th` + class(nil, t.align(i)) + `>` + inlines(c) + `</th>`)
	}
	s.WriteString(`</tr></thead><tbody>`)
	for _, r := range t.Rows {
		s.WriteString(`<tr>`)
		for i, c := range r {
			s.WriteString(`<td` + class(nil, t.align(i)) + `>` + inlines(c) + `</td>`)
		}
		s.WriteString(`</tr>`)
	}
	s.WriteString(`</tbody></table></div></div>`)
	return s.String()
}

// DocPart is what a document holds.
type DocPart interface {
	part
	docPart()
}

func (DocTitle) docPart()   {}
func (DocSub) docPart()     {}
func (Prose) docPart()      {}
func (DocTasks) docPart()   {}
func (DocFiles) docPart()   {}
func (DocRelated) docPart() {}
func (Area) docPart()       {}
func (DocFields) docPart()  {}
func (DocNote) docPart()    {}
func (DocHeading) docPart() {}
func (DocBullets) docPart() {}
func (DocBody) docPart()    {}
func (DocTable) docPart()   {}
func (DocQuote) docPart()   {}
func (DocCode) docPart()    {}

// Doc — tc-doc
//
// Роль. Документ: текст для чтения или в правке — заголовок, текст, дела, вложения.
// Вход. Doc{Items, Text, Action, Method, Fields, ID}.
// Состояния. read | edit | public | text | phone | table.
// Инварианты. Заголовок — первым; дела — галки, а не строки текста с квадратными скобками; документ печатает фрагмент своего кадра.
// Поведение. Галка дела в тексте — знак (Tick.Sign): дело отмечают в правке; над таблицей в узком контейнере переключатель «Wrap / One line» держит строки таблицы в одну линию.
// Эталон. кадр `s-notes-2`, фрагмент `.tc-doc`.
// Содержимое. Заголовок и строка под ним, абзацы, дела, вложения, связанное; в правке — поле заголовка, поля документа, поле текста, примечание; в тексте — подзаголовки, списки и вложенные списки, абзацы с курсивом, ссылками и кодом, цитаты, блоки кода, таблицы.
// Чего нет. Своего меню: действия — в шапке и в полосе правки.
type Doc struct {
	Items []DocPart
	Text  bool // the rendered text of a note alone: a column with 14 px gaps
	// Action makes the document in its edit a form that sends its fields
	// there, as Form does (Carrel's task 98): Method is "post" when empty,
	// Fields its hidden fields (CSRF, the record's version), ID what the
	// edit bar's Save names in its Form.
	Action string
	Method string
	Fields []Hidden
	ID     string
}

// Render prints the document.
func (d Doc) Render() string { return d.html(nil) }

func (d Doc) html(m mix) string {
	var s strings.Builder
	tag := "div"
	if d.Action != "" {
		tag = "form"
		s.WriteString(`<form` + class(m, "tc-doc", pick(d.Text, "tc-stack tc-stack--gap-14")) + attr("id", d.ID) + attr("action", d.Action) + ` method="` + method(d.Method) + `">`)
		for _, h := range d.Fields {
			s.WriteString(h.html(nil))
		}
	} else {
		s.WriteString(`<div` + class(m, "tc-doc", pick(d.Text, "tc-stack tc-stack--gap-14")) + `>`)
	}
	for _, it := range d.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</` + tag + `>`)
	return multipart(s.String())
}

// DocbarPart is what an edit bar holds before its end.
type DocbarPart interface {
	part
	docbarPart()
}

func (Button) docbarPart() {}
func (BarSep) docbarPart() {}
func (Drop) docbarPart()   {}

// Field in an edit's bar is where the record is saved: a note's notebook
// after Cancel (the owner's decision of 16.10).
func (Field) docbarPart() {}

// DraftNote says where the draft is kept and when it was saved.
type DraftNote struct {
	Text      string
	Time      Mono
	Secondary bool
}

func (n DraftNote) html(mix) string {
	return `<span` + class(nil, rank(n.Secondary)) + `>` + esc(n.Text) + ` ` + n.Time.html(nil) + `</span>`
}

// DocbarEnd is what the end of an edit bar holds.
type DocbarEnd interface {
	part
	docbarEnd()
}

func (Check) docbarEnd()     {}
func (DraftNote) docbarEnd() {}
func (Kbd) docbarEnd()       {}
func (Seg) docbarEnd()       {}
func (Button) docbarEnd()    {}

// Docbar — tc-docbar
//
// Роль. Полоса правки над документом: сохранить, отменить, разметка, вложения; справа — черновик и вид.
// Вход. Docbar{Edit, Items, More, End, Rule}.
// Состояния. plain | rule.
// Инварианты. Полоса — одна строка, пока помещается: второстепенное уходит под «⋯» первым, не поместившееся после этого — переносится, краем не режется (Р-65); сохранить — первой кнопкой; полоса печатает фрагмент своего кадра.
// Поведение. Ctrl+S сохраняет; «⋯» открывает то, что не поместилось.
// Эталон. кадр `components-29`, фрагмент `.tc-docbar`.
// Содержимое. Кнопки, разделители, «⋯», поле — куда сохраняется запись; в конце — автосохранение, черновик, клавиша, ширина текста, фокус.
// Чего нет. Второй строки: то, что не влезло, уходит под «⋯».
type Docbar struct {
	// Edit is an edit's bar: without a script only its Save and Cancel stand
	// — an edit is not reloaded midway, and the rest needs the page's script
	// (the owner's decisions of 02.10, Carrel's task 104).
	Edit  bool
	Items []DocbarPart
	More  *Button
	End   []DocbarEnd
	Rule  bool // a rule below: the bar drawn alone
}

// Render prints the bar.
func (b Docbar) Render() string { return b.html(nil) }

func (b Docbar) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-docbar", pick(b.Rule, "tc-docbar--rule"), pick(b.Edit, "tc-docbar--edit")) + `>`)
	for _, it := range b.Items {
		s.WriteString(it.html(nil))
	}
	if b.More != nil {
		s.WriteString(b.More.html(mix{"tc-bar__more"}))
	}
	if len(b.End) > 0 {
		s.WriteString(`<span class="tc-bar__right">`)
		for _, it := range b.End {
			s.WriteString(it.html(nil))
		}
		s.WriteString(`</span>`)
	}
	s.WriteString(`</div>`)
	return s.String()
}

// DocLayoutPart is what a document's layout holds.
type DocLayoutPart interface {
	part
	docLayoutPart()
}

func (Doc) docLayoutPart()     {}
func (Docbar) docLayoutPart()  {}
func (Docmeta) docLayoutPart() {}

// DocView is how a document is laid out in its page.
type DocView int

const (
	DocViewFull    DocView = iota // the page's width, the side column beside it
	DocViewMeasure                // the reading measure
	DocViewFocus                  // focus: the document alone
)

// DocLayout is a document in its page: the document and its side column.
type DocLayout struct {
	Items []DocLayoutPart
	View  DocView
}

func (l DocLayout) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(nil, is("measure", l.View == DocViewMeasure), is("focus", l.View == DocViewFocus), "tc-doc__layout") + `>`)
	for _, it := range l.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

func (DocLayout) mainPart() {}
func (Docbar) mainPart()    {}
func (Prose) dialogPart()   {}
func (Hero) mainPart()      {}
func (Doc) mainPart()       {}

// Hero is the head of a service's own page: its picture from the sprite, its
// name, a line, a paragraph and the buttons.
type Hero struct {
	Symbol string // the picture's symbol in the sprite, drawn in a 96 box
	Label  string // what the picture shows
	Title  string
	Sub    string
	Text   Prose
	Fields *Fields // what the service is: its address, its build
	Acts   []Button
}

func (h Hero) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-main__hero"><svg width="72" height="72" viewBox="0 0 96 96" role="img"` + attr("aria-label", h.Label) +
		` class="tc-ico--fix"><use href="#` + esc(h.Symbol) + `"/></svg><div class="tc-main__hero-side"><div class="tc-main__hero-title">` + esc(h.Title) +
		`</div><div class="tc-head__sub">` + esc(h.Sub) + `</div>` + h.Text.html(nil))
	if h.Fields != nil {
		s.WriteString(h.Fields.html(nil))
	}
	s.WriteString(`<div class="tc-line tc-line--gap-8 tc-doc__acts">`)
	for _, b := range h.Acts {
		s.WriteString(b.html(nil))
	}
	s.WriteString(`</div></div></div>`)
	return s.String()
}
