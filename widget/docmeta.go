package widget

import "strings"

// DocRef is a record a document points at: the source's mark, the name (a
// link when it opens elsewhere) and what it is.
type DocRef struct {
	Mark *Mark
	Name string
	Href string // a link to the record
	What Dim
	// From is the copy it comes from, on a linked record (Р-78; Carrel's
	// task 230): the collection's mark and name.
	From *Dim
}

// DocmetaPart is what a section of the side column holds under its rubric.
type DocmetaPart interface {
	part
	docmetaPart()
}

func (Line) docmetaPart()      {}
func (DocRefs) docmetaPart()   {}
func (DocAttach) docmetaPart() {}
func (Diag) docmetaPart()      {}
func (PostForm) docmetaPart()  {}
func (SecHint) docmetaPart()   {}
func (SecActs) docmetaPart()   {}

// DocRefs is a column of records a document points at.
type DocRefs struct {
	Items []DocRef
}

func (r DocRefs) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-stack tc-stack--tight tc-line__text">`)
	for _, it := range r.Items {
		s.WriteString(`<span class="tc-stack__item">`)
		if it.Mark != nil {
			s.WriteString(it.Mark.html(nil) + " ")
		}
		if it.Href != "" {
			s.WriteString(Link{Text: it.Name, Href: it.Href}.html(nil))
		} else {
			s.WriteString(`<span class="tc-dim--accent">` + esc(it.Name) + `</span>`)
		}
		if len(it.What.Parts) > 0 { // a copy that holds the same says nothing
			s.WriteString(" " + it.What.html(nil))
		}
		if it.From != nil {
			s.WriteString(" " + it.From.html(nil))
		}
		s.WriteString(`</span>`)
	}
	s.WriteString(`</div>`)
	return s.String()
}

// DocAttach is an attachment line in the side column: the name, the size,
// where it lives.
type DocAttach struct {
	Name  string
	Size  Dim
	Where Badge
	// Href makes the name a link (the owner's decision of 02.10): an
	// attachment in the record downloads, one in Files opens there.
	Href string
	// Remove takes the attachment off: on an edit only (the owner's decision
	// of 02.10).
	Remove *Button
	// From is the copy it comes from, on a linked record (Р-78).
	From *Dim
}

func (a DocAttach) html(mix) string {
	name := esc(a.Name)
	if a.Href != "" {
		name = Link{Text: a.Name, Href: a.Href}.html(nil)
	}
	s := `<div class="tc-line__text">` + name + ` ` + a.Size.html(nil)
	if a.Where.Label != "" {
		s += ` ` + a.Where.html(nil)
	}
	if a.From != nil {
		s += ` ` + a.From.html(nil)
	}
	if a.Remove != nil {
		s += ` ` + a.Remove.html(nil)
	}
	return s + `</div>`
}

// DocmetaSec is a section of a document's side column.
type DocmetaSec struct {
	Title Rubric
	Items []DocmetaPart
	ID    string // the list a picked file joins on an edit (FileField.List)
}

// Docmeta — tc-docmeta
//
// Роль. Колонка сведений документа: теги, связанное, вложения, копии в других блокнотах, источник, удаление.
// Вход. Docmeta{Sections, Danger}.
// Состояния. read.
// Инварианты. Каждая секция — рубрика и то, что под ней; удаление — последним; колонка печатает фрагмент своего кадра.
// Поведение. нет: поведение у ссылок и кнопок.
// Эталон. кадр `s-notes-2`, фрагмент `.tc-docmeta`.
// Содержимое. Секции: теги, связанные записи, вложения с формой их Attach, копии, сохранённое как есть (Show N и пояснение), источник; опасная зона.
// Чего нет. Полей правки: правят в документе.
type Docmeta struct {
	Sections []DocmetaSec
	Danger   *Danger
}

// Render prints the column.
func (d Docmeta) Render() string { return d.html(nil) }

func (d Docmeta) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-docmeta") + `>`)
	for _, sec := range d.Sections {
		s.WriteString(`<div class="tc-docmeta__sec"` + attr("id", sec.ID) + `>` + sec.Title.html(nil))
		for _, it := range sec.Items {
			s.WriteString(it.html(nil))
		}
		s.WriteString(`</div>`)
	}
	if d.Danger != nil {
		s.WriteString(d.Danger.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}
