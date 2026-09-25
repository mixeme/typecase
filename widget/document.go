package widget

// Doc — tc-doc
//
// Роль. Сплошной текст, который читают и правят.
// Вход. Doc{ Mode, Width, Body, Edit, Bar, Meta }.
// Состояния. reading | editing.
// Инварианты. Переключатель меры — часть виджета; тихого сохранения нет.
// Поведение. Переключение меры, переход в правку, предупреждение при уходе.
// Геометрия. Сетка .tc-doc__layout: minmax(0,1fr) 400px; в рабочей области уже 880 — 296px.
// Содержимое. tc-docbar, tc-prose, tc-docedit, tc-docmeta.
// Чего нет. Нет панели форматирования с разметкой.
type Doc struct {
	Mode  DocMode
	Width DocWidth
	Body  Prose
	Edit  *DocEdit
	Bar   DocBar
	Meta  DocMeta
}

func (d Doc) Render() string {
	mode := d.Mode
	if mode == "" {
		mode = DocReading
	}
	width := d.Width
	if width == "" {
		width = DocWidthMeasure
	}
	body := d.Body.Render()
	if d.Edit != nil {
		body = d.Edit.Render()
	}
	return `<article` + classAttr("tc-doc", "tc-doc--"+string(mode), "tc-doc--"+string(width)) + `>` + d.Bar.Render() + `<div class="tc-doc__layout">` + body + d.Meta.Render() + `</div></article>`
}

// DocEdit — tc-docedit
//
// Роль. Кадр сосредоточенной правки: заголовок, свойства, текст.
// Вход. DocEdit{ Title, Rows []PFieldRow, Text }.
// Состояния. нет.
// Инварианты. Заголовок, свойства и текст на всю ширину колонки.
// Поведение. нет.
// Геометрия. Столбец на всю колонку содержимого.
// Содержимое. tc-pfield, tc-pfieldrow, tc-in.
// Чего нет. Нет рейла и боковой колонки.
type DocEdit struct {
	Title string
	Rows  []PFieldRow
	Text  string
}

func (e DocEdit) Render() string {
	title := PField{Name: "title", Label: "Title", Kind: InputText, Value: e.Title, Title: true, Width: PFieldWidthFull}.Render()
	var rows string
	for _, r := range e.Rows {
		rows += r.Render()
	}
	text := PField{Name: "body", Kind: InputTextarea, Value: e.Text, Width: PFieldWidthFull}.Render()
	return `<div class="tc-docedit">` + title + rows + text + `</div>`
}

// Prose — tc-prose
//
// Роль. Показ размеченного текста, пришедшего деревом блоков.
// Вход. Prose{ Blocks }.
// Состояния. нет.
// Инварианты. Готовую разметку принять невозможно.
// Поведение. нет.
// Геометрия. Мера строки от tc-doc.
// Содержимое. Блоки объявленных родов.
// Чего нет. Нет разбора разметки.
type Prose struct {
	Blocks []ProseBlock
}

func (p Prose) Render() string {
	var blocks string
	for _, b := range p.Blocks {
		blocks += renderProseBlock(b)
	}
	return `<div class="tc-prose">` + blocks + `</div>`
}

func renderProseBlock(b ProseBlock) string {
	switch b.Kind {
	case ProseH1, ProseH2, ProseH3:
		tag := string(b.Kind)
		return `<` + tag + ` class="tc-prose__` + tag + `">` + esc(b.Text) + `</` + tag + `>`
	case ProseP:
		return `<p class="tc-prose__p">` + esc(b.Text) + `</p>`
	case ProseUL:
		var items string
		for _, it := range b.Items {
			items += `<li>` + esc(it) + `</li>`
		}
		return `<ul class="tc-prose__ul">` + items + `</ul>`
	case ProseCode:
		return `<pre class="tc-prose__code"><code>` + esc(b.Code) + `</code></pre>`
	case ProseLink:
		return `<p class="tc-prose__p"><a href="` + esc(b.Href) + `" class="tc-prose__link">` + esc(b.Text) + `</a></p>`
	default:
		return `<p class="tc-prose__p">` + esc(b.Text) + `</p>`
	}
}

// DocBar — tc-docbar
//
// Роль. Действия над документом: сохранить, переключить меру, сосредоточиться.
// Вход. DocBar{ Sub, Mark, Left, Right, RightNodes, Rank2, Dirty, Collapsed }.
// Состояния. full | collapsed.
// Инварианты. Свой порог отступления; Sub, действия и отступление — одна строка кадра.
// Поведение. Как tc-bar, плюс признак несохранённого.
// Геометрия. Высота панели инструментов.
// Содержимое. tc-btn, tc-menu, tc-kbd, tc-mark.
// Чего нет. Нет кнопок, не относящихся к документу.
type DocBar struct {
	Sub        string
	Mark       Mark
	Left       []Action
	Right      []Action
	RightNodes []Node
	Rank2      []Action
	Dirty      bool
	Collapsed  bool
}

func (d DocBar) Render() string {
	cls := "tc-docbar"
	if d.Dirty {
		cls = classes(cls, "is-dirty")
	}
	if d.Collapsed {
		cls = classes(cls, "is-collapsed")
	}
	if d.Sub != "" {
		cls = classes(cls, "tc-docbar--frame")
	}
	dirty := ""
	if d.Dirty {
		dirty = `<span class="tc-docbar__dirty" aria-label="` + esc(labels.UnsavedChanges) + `">●</span>`
	}
	right := append(ActionsAsNodes(d.Right), d.RightNodes...)
	bar := Bar{Left: ActionsAsNodes(d.Left), Right: right, Rank2: d.Rank2}.Render()
	if d.Sub == "" {
		return `<div` + classAttr(cls) + `>` + dirty + bar + `</div>`
	}
	return `<div` + classAttr(cls) + `><div class="tc-docbar__frame"><div class="tc-docbar__sub">` + d.Mark.Render() + esc(d.Sub) + `</div><div class="tc-docbar__tools">` + dirty + bar + `</div></div></div>`
}

// DocMeta — tc-docmeta
//
// Роль. Дата, метки, связи, вложения сбоку от текста.
// Вход. DocMeta{ Sections, Danger }.
// Состояния. нет.
// Инварианты. Не превращается в форму.
// Поведение. нет.
// Геометрия. Колонка 400, как панель подробностей; в рабочей области уже 880 — 296; в узкой раскладке под текстом.
// Содержимое. tc-tag, tc-fields, tc-row, tc-btn.
// Чего нет. Нет прокрутки отдельно от документа.
type DocMeta struct {
	Sections []MetaSection
	Danger   *Action
}

func (m DocMeta) Render() string {
	var secs string
	for _, s := range m.Sections {
		var tags string
		for _, t := range s.Tags {
			tags += t.Render()
		}
		var rows string
		for _, r := range s.Rows {
			rows += FieldRowHTML(r)
		}
		for _, r := range s.Rows2 {
			rows += r.Render()
		}
		action := ""
		if s.Action != nil {
			a := *s.Action
			if a.Size == "" {
				a.Size = ButtonSizeTable
			}
			action = `<span class="tc-docmeta__action">` + a.Render() + `</span>`
		}
		secs += `<section class="tc-docmeta__section"><h3 class="tc-docmeta__title">` + esc(s.Title) + action + `</h3><div class="tc-docmeta__tags">` + tags + `</div><div class="tc-docmeta__rows">` + rows + `</div></section>`
	}
	danger := ""
	if m.Danger != nil {
		a := *m.Danger
		if a.Kind == "" {
			a.Kind = ButtonDangerSolid
		}
		if a.Size == "" {
			a.Size = ButtonSizeForm
		}
		danger = `<div class="tc-docmeta__danger">` + a.Render() + `</div>`
	}
	if secs == "" && danger == "" {
		return ""
	}
	return `<aside class="tc-docmeta">` + secs + danger + `</aside>`
}
