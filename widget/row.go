package widget

import (
	"strconv"
	"strings"
)

// A row of a list is laid out in columns, and the frames draw nine sets of
// them. Each set is a type of its own, so a row gets exactly the cells its
// columns have: a row of times has no avatar, a row of people has no due
// date. All of them are tc-row.

// ListRow is a row a list holds.
type ListRow interface {
	part
	listRow()
}

func (TimeRow) listRow()    {}
func (ColumnRow) listRow()  {}
func (SkelRow) listRow()    {}
func (FaceRow) listRow()    {}
func (FindRow) listRow()    {}
func (LinkedRow) listRow()  {}
func (ExcerptRow) listRow() {}
func (BriefRow) listRow()   {}
func (DueRow) listRow()     {}
func (SizeRow) listRow()    {}
func (SlimRow) listRow()    {}
func (PhoneRow) listRow()   {}
func (MailRow) listRow()    {}

// RowTag is what the tags cell of a row holds.
type RowTag interface {
	part
	rowTag()
}

func (Tag) rowTag()   {}
func (Badge) rowTag() {}

// inlines prints running text without separators: the pieces carry their
// own spaces.
func inlines(parts []Inline) string {
	var s strings.Builder
	for _, p := range parts {
		s.WriteString(p.html(nil))
	}
	return s.String()
}

// plain is running text without its markup: an accessible name.
func plain(parts []Inline) string {
	var s strings.Builder
	for _, p := range parts {
		switch v := p.(type) {
		case Text:
			s.WriteString(string(v))
		case Strong:
			s.WriteString(string(v))
		case Mono:
			s.WriteString(v.Text)
		case Span:
			s.WriteString(v.Text)
		case Hint:
			s.WriteString(v.Text)
		}
	}
	return s.String()
}

// cell prints a cell of a row with its element classes.
func cell(text string, own ...string) string {
	return `<span` + class(nil, own...) + `>` + esc(text) + `</span>`
}

// title prints the title cell: the words, the bold of what was found, a
// badge after them.
func title(parts []Inline, strong bool, badge *Badge, own ...string) string {
	return `<span` + class(nil, append([]string{"tc-row__title", is("strong", strong)}, own...)...) + `>` + titled(parts, badge) + `</span>`
}

// titled is a title's words and its badge: with a badge the words are an
// element of their own, which shortens with «…» while the badge stays whole
// («Анна Ковалёва linked ×3» at 640, Carrel's task 190).
func titled(parts []Inline, badge *Badge) string {
	if badge == nil {
		return inlines(parts)
	}
	return `<span class="tc-row__words">` + inlines(parts) + `</span> ` + badge.html(nil)
}

// linkTitle is the row's title, a link to the record when href is set: in a
// row with a control of its own the title, not the row, goes there.
func linkTitle(href, page string, parts []Inline, strong bool, badge *Badge) string {
	if href == "" {
		return title(parts, strong, badge)
	}
	return `<a` + class(nil, "tc-row__title", is("strong", strong)) + attr("href", href) + attr("data-tc-page", page) + `>` + titled(parts, badge) + `</a>`
}

// source prints the source cell: where the record lives.
func source(text, hint string, badge *Badge, single ...bool) string {
	if len(single) > 0 && single[0] {
		return ""
	}
	s := `<span class="tc-row__cell tc-row__source"` + attr("title", hint) + `>` + esc(text)
	if badge != nil {
		s += " " + badge.html(nil)
	}
	return s + `</span>`
}

// rowTagHTML prints a tag of a row; in a row that is a link itself a tag
// is words, not a link inside a link (HTML has none).
func rowTagHTML(t RowTag, inLink bool) string {
	if tg, ok := t.(Tag); ok && inLink {
		tg.Href = ""
		return tg.html(nil)
	}
	return t.html(nil)
}

// rowTagLabel is the words of a tag of a row.
func rowTagLabel(t RowTag) string {
	switch v := t.(type) {
	case Tag:
		return v.Label
	case Badge:
		return v.Label
	}
	return ""
}

// inset is the variant of a row set in a card: the narrow gutter at the
// sides.
func inset(kind string, on bool) string {
	return pick(on, "tc-row--"+kind+"--inset")
}

// TimeRow — tc-row
//
// Роль. Строка со временем: событие дня, приглашение, запись журнала.
// Вход. TimeRow{Mark, Time, Title, Strong, Detail, Badges, Source, On, Past, Inset, Single, Href, Page}.
// Состояния. plain | on | caption | past | inset | found.
// Инварианты. Колонки строки — её тип: у строки времени нет аватара и срока; метка источника — первой; строка печатает фрагмент своего кадра.
// Поведение. Нажатие открывает запись; стрелки ходят по строкам списка; с Page, где ширина не рисует панель, нажатие ведёт на страницу записи (задача Carrel 290).
// Эталон. кадр `s-agenda-1`, фрагмент `.tc-row`.
// Содержимое. Метка, время моноширинным, название (с найденным жирным), место, плашки по важности (видны те, что умещаются целыми; подсказка ячейки называет все, когда их больше одной), источник.
// Чего нет. Своих кнопок: действия с записью — в панели и в меню «⋯» шапки.
type TimeRow struct {
	Mark   Mark
	Time   string
	Title  []Inline
	Strong bool // not seen yet
	Detail string
	// Badges are the event's badges by importance — its status, attendees,
	// repeat; the row shows those that fit whole, two side by side.
	Badges []Badge
	Source string
	On     bool
	Past   bool // gone by: faded
	Inset  bool
	Single bool   // the list of one collection: the row has no source column
	Href   string // the row is a link to its record
	// Page is the record\'s own page: where the row goes when the page has no
	// panel to open it in (Carrel\'s task 290).
	Page string
}

// Render prints the row.
func (r TimeRow) Render() string { return r.html(nil) }

func (r TimeRow) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<` + rowTag(r.Href) + class(m, "tc-row", "tc-row--timecols", is("on", r.On), inset("timecols", r.Inset), pick(r.Single, "tc-row--single"), pick(r.Past, "tc-row--timecols--op-7")) + attr("href", r.Href) + attr("data-tc-page", r.Page) + `>`)
	s.WriteString(r.Mark.html(nil))
	s.WriteString(cell(r.Time, "tc-mono", "tc-row__lead"))
	s.WriteString(title(r.Title, r.Strong, nil))
	s.WriteString(cell(r.Detail, "tc-row__excerpt", "tc-row__detail"))
	s.WriteString(badgesCell("tc-row__cell tc-row__caption", r.Badges))
	s.WriteString(source(r.Source, "", nil, r.Single))
	s.WriteString(`</` + rowTag(r.Href) + `>`)
	return s.String()
}

// SkelRow — tc-row
//
// Роль. Строка списка, пока он грузится: заготовки на местах колонок.
// Вход. SkelRow{Title}.
// Состояния. loading.
// Инварианты. Колонки — те же, что у строки времени; ширина заготовки названия задаёт ширины остальных, чтобы список не выглядел разлинованным; строка печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `states-3`, фрагмент `.tc-row`.
// Содержимое. Пустая метка на линии, заготовки времени, названия, места и источника.
// Чего нет. Текста: что грузится, говорит шапка.
type SkelRow struct {
	Title Skel // Skel60, Skel45 or Skel52: one of the three drawn lines
}

// Render prints the row.
func (r SkelRow) Render() string { return r.html(nil) }

func (r SkelRow) html(m mix) string {
	w := map[SkelWidth][2]string{Skel45: {"30pc", "55pc"}, Skel52: {"25pc", "62pc"}}[r.Title.Width]
	if w[0] == "" {
		w = [2]string{"40pc", "70pc"}
	}
	return `<div` + class(m, "tc-row", "tc-row--timecols") + `>` + Mark{Size: MarkOnHair}.html(nil) +
		Skel{}.html(mix{"tc-row__lead", "tc-row__lead--w36"}) + r.Title.html(nil) +
		Skel{}.html(mix{"tc-row__detail", "tc-row__detail--w" + w[0]}) + `<span class="tc-row__caption"></span>` +
		Skel{}.html(mix{"tc-row__source", "tc-row__source--w" + w[1]}) + `</div>`
}

// FaceRow — tc-row
//
// Роль. Строка человека: аватар, имя, телефон, почта, организация.
// Вход. FaceRow{Mark, Avatar, Title, Strong, TitleBadge, Subtitle, Detail, Caption, Source, SourceHint, SourceBadge, On, Inset, Single, Href, Page}.
// Состояния. plain | on | strong | badge | inset | found.
// Инварианты. Колонки строки — её тип: у строки человека есть аватар и нет времени; метка источника — первой; строка печатает фрагмент своего кадра.
// Поведение. Нажатие открывает карточку; стрелки ходят по строкам списка; с Page, где ширина не рисует панель, нажатие ведёт на страницу записи (задача Carrel 290).
// Эталон. кадр `components-12`, фрагмент `.tc-row`.
// Содержимое. Метка, аватар, имя (с найденным жирным и бейджем), телефон моноширинным, почта, организация, источник (с бейджем).
// Чего нет. Своих кнопок: действия с карточкой — в панели.
type FaceRow struct {
	Mark        Mark
	Avatar      Avatar
	Title       []Inline
	Strong      bool
	TitleBadge  *Badge
	Subtitle    string // monospaced: a phone
	Detail      string
	Caption     string
	Source      string
	SourceHint  string // the sources in full when there are several: a linked person
	SourceBadge *Badge
	On          bool
	Inset       bool
	Single      bool   // the list of one collection: the row has no source column
	Href        string // the row is a link to its record
	// Page is the record\'s own page: where the row goes when the page has no
	// panel to open it in (Carrel\'s task 290).
	Page string
}

// Render prints the row.
func (r FaceRow) Render() string { return r.html(nil) }

func (r FaceRow) html(m mix) string {
	return `<` + rowTag(r.Href) + class(m, "tc-row", "tc-row--avatarcols", is("on", r.On), inset("avatarcols", r.Inset), pick(r.Single, "tc-row--single")) + attr("href", r.Href) + attr("data-tc-page", r.Page) + `>` +
		r.Mark.html(nil) + r.Avatar.html(nil) + title(r.Title, r.Strong, r.TitleBadge) +
		cell(r.Subtitle, "tc-row__excerpt", "tc-mono", "tc-row__subtitle") + cell(r.Detail, "tc-row__excerpt", "tc-row__detail") +
		cell(r.Caption, "tc-row__excerpt", "tc-row__caption") + source(r.Source, r.SourceHint, r.SourceBadge, r.Single) + `</` + rowTag(r.Href) + `>`
}

// FindRow — tc-row
//
// Роль. Найденная запись: дата, название, строка с найденным, что это и где лежит.
// Вход. FindRow{Mark, Date, Title, Detail, What, Source, Href}.
// Состояния. plain | found.
// Инварианты. Колонки строки — её тип; найденное выделено акцентом внутри строки, а не отдельной колонкой; строка печатает фрагмент своего кадра.
// Поведение. Нажатие открывает запись.
// Эталон. кадр `mentions-2`, фрагмент `.tc-row`.
// Содержимое. Метка, дата моноширинным, название, строка с найденным, род записи, источник.
// Чего нет. Своих кнопок.
type FindRow struct {
	Mark   Mark
	Date   string
	Title  []Inline
	Detail []Inline
	What   string // what kind of record it is
	Source string
	Href   string // the row is a link to its record
}

// Render prints the row.
func (r FindRow) Render() string { return r.html(nil) }

func (r FindRow) html(m mix) string {
	return `<` + rowTag(r.Href) + class(m, "tc-row", "tc-row--find") + attr("href", r.Href) + `>` + r.Mark.html(nil) + cell(r.Date, "tc-mono", "tc-row__lead") +
		title(r.Title, false, nil) + `<span class="tc-row__excerpt tc-row__detail">` + inlines(r.Detail) + `</span>` +
		cell(r.What, "tc-row__cell", "tc-row__caption") + source(r.Source, "", nil) + `</` + rowTag(r.Href) + `>`
}

// LinkedRow — tc-row
//
// Роль. Связанная запись в карточке: запись другого рода, где встречается тот же человек.
// Вход. LinkedRow{Mark, Icon, Title, Detail, Href}.
// Состояния. plain.
// Инварианты. Колонки строки — её тип: значок рода записи на месте даты; строка печатает фрагмент своего кадра.
// Поведение. Нажатие открывает запись.
// Эталон. кадр `s-contacts-6`, фрагмент `.tc-row`.
// Содержимое. Метка, значок рода записи, название, когда и в какой роли.
// Чего нет. Источника: запись лежит там же, где карточка, или это сказано в строке.
type LinkedRow struct {
	Mark   Mark
	Icon   Icon
	Title  string
	Detail string
	Href   string // the row is a link to its record
}

// Render prints the row.
func (r LinkedRow) Render() string { return r.html(nil) }

func (r LinkedRow) html(m mix) string {
	return `<` + rowTag(r.Href) + class(m, "tc-row", "tc-row--find") + attr("href", r.Href) + `>` + r.Mark.html(nil) +
		`<svg width="12" height="12" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.5" fill="none" class="tc-row__lead">` + r.Icon.use() + `</svg>` +
		cell(r.Title, "tc-row__title") + cell(r.Detail, "tc-row__cell", "tc-row__detail") + `</` + rowTag(r.Href) + `>`
}

// ExcerptRow — tc-row
//
// Роль. Строка текста: название, начало текста, дата, источник.
// Вход. ExcerptRow{Mark, Title, Strong, TitleBadge, Excerpt, Date, Source, SourceHint, On, Inset, Single, Href, Page}.
// Состояния. plain | on | badge | inset | found | link.
// Инварианты. Колонки строки — её тип; метка источника — первой; строка печатает фрагмент своего кадра.
// Поведение. Нажатие открывает текст; стрелки ходят по строкам списка; с Page, где ширина не рисует панель, нажатие ведёт на страницу записи (задача Carrel 290).
// Эталон. кадр `s-notes-1`, фрагмент `.tc-row`.
// Содержимое. Метка, название (с найденным жирным и бейджем), начало текста, дата моноширинным, источник (с подсказкой, когда их несколько).
// Чего нет. Всего текста: строка показывает начало, текст открывается. Тегов: они в панели записи и в рейле, в строке — столбцом, который включил человек (ColumnRow, Р-53).
type ExcerptRow struct {
	Mark       Mark
	Title      []Inline
	Strong     bool
	TitleBadge *Badge
	Excerpt    string
	Date       string
	Source     string
	SourceHint string // the sources in full when there are several
	On         bool
	Inset      bool
	Single     bool   // the list of one collection: the row has no source column
	Href       string // the row is a link to its record
	// Page is the record\'s own page: where the row goes when the page has no
	// panel to open it in (Carrel\'s task 290).
	Page string
}

// Render prints the row.
func (r ExcerptRow) Render() string { return r.html(nil) }

func (r ExcerptRow) html(m mix) string {
	return `<` + rowTag(r.Href) + class(m, "tc-row", "tc-row--excerpt", is("on", r.On), inset("excerpt", r.Inset), pick(r.Single, "tc-row--single")) + attr("href", r.Href) + attr("data-tc-page", r.Page) + `>` + r.Mark.html(nil) +
		title(r.Title, r.Strong, r.TitleBadge) + cell(r.Excerpt, "tc-row__excerpt") +
		cell(r.Date, "tc-mono", "tc-row__date") + source(r.Source, r.SourceHint, nil, r.Single) + `</` + rowTag(r.Href) + `>`
}

// BriefRow — tc-row
//
// Роль. Короткая строка текста в выборе: название, дата, где лежит.
// Вход. BriefRow{Mark, Title, Date, Where, On, Href, Page}.
// Состояния. plain | on.
// Инварианты. Колонок три, и это её тип; строка печатает фрагмент своего кадра.
// Поведение. Нажатие выбирает строку; с Page, где ширина не рисует панель, нажатие ведёт на страницу записи (задача Carrel 290).
// Эталон. кадр `s-mail-5`, фрагмент `.tc-row`.
// Содержимое. Метка, название, дата моноширинным, источник.
// Чего нет. Начала текста и тегов: выбор короче списка.
type BriefRow struct {
	Mark  Mark
	Title string
	Date  Mono
	Where string
	On    bool
	Href  string // the row is a link to its record
	// Page is the record\'s own page: where the row goes when the page has no
	// panel to open it in (Carrel\'s task 290).
	Page string
}

// Render prints the row.
func (r BriefRow) Render() string { return r.html(nil) }

func (r BriefRow) html(m mix) string {
	return `<` + rowTag(r.Href) + class(m, "tc-row", "tc-row--excerpt", is("on", r.On), "tc-row--excerpt--cols3-62-148") + attr("href", r.Href) + attr("data-tc-page", r.Page) + `>` + r.Mark.html(nil) +
		cell(r.Title, "tc-row__title", "is-strong") + r.Date.html(nil) + cell(r.Where, "tc-row__cell", "tc-row__tags") + `</` + rowTag(r.Href) + `>`
}

// DueRow — tc-row
//
// Роль. Строка дела: галка, название, срок, сколько сделано, источник.
// Вход. DueRow{Tick, Mark, Title, Strong, TitleBadge, Due, Progress, Source, SourceHint, On, Overdue, Inset, Single, Href, Page}.
// Состояния. plain | on | overdue | done | inset | found | sign.
// Инварианты. Сделано ли дело, говорит одна галка: строка берёт это из неё; колонки строки — её тип; строка печатает фрагмент своего кадра.
// Поведение. Галка отмечает дело сделанным; нажатие на строку открывает его; в списке только для чтения галка — знак (Tick.Sign): дело не отметить; с Page, где ширина не рисует панель, нажатие ведёт на страницу записи (задача Carrel 290).
// Эталон. кадр `s-tasks-1`, фрагмент `.tc-row`.
// Содержимое. Галка, метка, название (с найденным жирным и бейджем), срок и доля моноширинным, источник.
// Чего нет. Подзадач: они — в панели дела. Тегов: они в панели дела и в рейле, в строке — столбцом, который включил человек (ColumnRow, Р-53).
type DueRow struct {
	Tick       Tick
	Mark       Mark
	Title      []Inline
	Strong     bool
	TitleBadge *Badge
	Due        string
	Progress   string
	Source     string
	SourceHint string
	On         bool
	Overdue    bool
	Inset      bool
	Single     bool // the list of one collection: the row has no source column
	// Href makes the title a link to the record (Carrel's task 95): the row
	// itself cannot be one, its tick is a control of its own.
	Href string
	// Page is the record's own page: where the row goes when the page has no
	// panel to open it in (Carrel's task 290).
	Page string
}

// Render prints the row.
func (r DueRow) Render() string { return r.html(nil) }

func (r DueRow) html(m mix) string {
	return `<div` + class(m, "tc-row", "tc-row--tickmarkcols", is("on", r.On), is("overdue", r.Overdue), is("done", r.Tick.Done), inset("tickmarkcols", r.Inset), pick(r.Single, "tc-row--single")) + `>` +
		r.Tick.html(nil) + r.Mark.html(nil) + linkTitle(r.Href, r.Page, r.Title, r.Strong, r.TitleBadge) +
		cell(r.Due, "tc-mono", "tc-row__due") + cell(r.Progress, "tc-row__cell", "tc-mono", "tc-row__progress") +
		source(r.Source, r.SourceHint, nil, r.Single) + `</div>`
}

// SizeRow — tc-row
//
// Роль. Строка вещи с размером: галка, значок и имя, размер, когда менялась, права, источник.
// Вход. SizeRow{Ticked, Mark, Icon, Name, Bytes, Time, Access, Source, Send, Href}.
// Состояния. found.
// Инварианты. Галка — кнопка с доступным именем; колонки строки — её тип; строка печатает фрагмент своего кадра.
// Поведение. Галка выбирает вещь для действий шапки; нажатие на имя открывает её.
// Эталон. кадр `s-find-1`, фрагмент `.tc-row`.
// Содержимое. Галка, значок рода, имя (с найденным жирным), размер, время и права моноширинным, источник.
// Чего нет. Метки источника: источник назван словами.
type SizeRow struct {
	Ticked Ticked
	// Mark stands in the tick's place where nothing is done with the file
	// but open it: a search's results (Carrel's task 189).
	Mark   *Mark
	Send   Send // what pressing the tick sends
	Icon   Icon
	Name   []Inline
	Bytes  string // 240 KB
	Time   string
	Access string // rw, ro
	Source string
	// Href makes the name a link to the file (Carrel's task 129): the row
	// cannot be one, its tick is a control of its own.
	Href string
}

// Render prints the row.
func (r SizeRow) Render() string { return r.html(nil) }

func (r SizeRow) html(m mix) string {
	name := "span"
	if r.Href != "" {
		name = "a"
	}
	lead := r.Ticked.box(plain(r.Name), r.Send)
	if r.Mark != nil {
		lead = r.Mark.html(nil)
	}
	return `<div` + class(m, "tc-row", "tc-row--tickcols") + `>` + lead +
		`<` + name + ` class="tc-row__title tc-row__name"` + attr("href", r.Href) + `><svg>` + r.Icon.use() + `</svg><span class="tc-row__file">` + inlines(r.Name) + `</span></` + name + `>` +
		cell(r.Bytes, "tc-row__excerpt", "tc-mono", "tc-row__size") + cell(r.Time, "tc-row__excerpt", "tc-mono", "tc-row__time") +
		cell(r.Access, "tc-row__cell", "tc-mono", "tc-row__detail") + source(r.Source, "", nil) + `</div>`
}

// SlimRow — tc-row
//
// Роль. Узкая строка выбора: галка, метка, что выбрать и что в нём, род.
// Вход. SlimRow{Ticked, Mark, Title, Inside, Aside, Send, Reason}.
// Состояния. plain | off | disabled.
// Инварианты. Галка — кнопка с доступным именем; строка печатает фрагмент своего кадра.
// Поведение. Галка выбирает строку; с Reason галка выключена, причина — её подсказка, род справа — тревогой.
// Эталон. кадр `s-plan-1`, фрагмент `.tc-row`.
// Содержимое. Галка, метка, название с тем, что в нём, род справа.
// Чего нет. Источника и даты.
type SlimRow struct {
	Ticked Ticked
	Send   Send // what pressing the tick sends
	Mark   Mark
	Title  string
	Inside string // what is in it, after the title
	Aside  string
	// Reason turns the tick off and says why; the aside says it as an
	// alarm: an account whose saved password the server refused cannot go
	// into a device's profile (Carrel's task 211).
	Reason string
}

// Render prints the row.
func (r SlimRow) Render() string { return r.html(nil) }

func (r SlimRow) html(m mix) string {
	tick, aside := r.Ticked.box(r.Title, r.Send), cell(r.Aside, "tc-row__cell")
	if r.Reason != "" {
		tick = `<button class="tc-tickrow__tick is-off"` + attr("title", r.Reason) + ` type="button"` + attr("aria-label", r.Title) + ` disabled></button>`
		aside = cell(r.Aside, "tc-row__cell", "tc-dim--danger")
	}
	return `<div` + class(m, "tc-row", "tc-row--slim") + `>` + tick + r.Mark.html(nil) +
		`<span class="tc-row__title">` + esc(r.Title) + ` ` + cell(r.Inside, "tc-row__cell") + `</span>` + aside + `</div>`
}

// TickBox is a tick box in front of a phone's row.
type TickBox struct {
	Ticked Ticked
	Send   Send
}

// PhoneRow — tc-row
//
// Роль. Строка списка на телефоне: две строки текста и одна пометка справа.
// Вход. PhoneRow{Mark, Pick, Avatar, Icon, Title, TitleBadge, Sub, SubMono, Aside, More, On}.
// Состояния. plain | on | avatar | badge | pick.
// Инварианты. Спереди — метка или галка, не обе; галка — кнопка с доступным именем; строка печатает фрагмент своего кадра.
// Поведение. Нажатие открывает запись; галка выбирает её.
// Эталон. кадр `phone-4`, фрагмент `.tc-row`.
// Содержимое. Метка или галка, аватар или значок, название (с бейджем), вторая строка, пометка или кнопка «⋯».
// Чего нет. Колонок: телефон ставит всё в две строки.
type PhoneRow struct {
	Mark       *Mark
	Pick       *TickBox // instead of the mark
	Avatar     *Avatar
	Icon       Icon // instead of the avatar
	Title      string
	TitleBadge *Badge
	Sub        []Inline
	SubMono    bool // the second line is monospaced
	Aside      string
	More       *Button // instead of the aside
	On         bool
}

// Render prints the row.
func (r PhoneRow) Render() string { return r.html(nil) }

func (r PhoneRow) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-row", "tc-row--phone", is("on", r.On)) + `>`)
	if r.Pick != nil {
		s.WriteString(r.Pick.Ticked.box(r.Title, r.Pick.Send, "tc-tickrow__tick--ml2"))
	} else if r.Mark != nil {
		s.WriteString(r.Mark.html(nil))
	}
	text := `<span class="tc-row__text">` + title([]Inline{Text(r.Title)}, false, r.TitleBadge) +
		`<span` + class(nil, "tc-row__cell", pick(r.SubMono, "tc-mono")) + `>` + inlines(r.Sub) + `</span></span>`
	switch {
	case r.Avatar != nil:
		s.WriteString(`<span class="tc-row__text tc-row__text--row">` + r.Avatar.html(nil) + text + `</span>`)
	case r.Icon != "":
		s.WriteString(`<span class="tc-row__text tc-row__text--row"><svg width="17" height="17" viewBox="0 0 24 24" stroke-width="1.3" fill="none" class="tc-ico--muted">` +
			r.Icon.use() + `</svg>` + text + `</span>`)
	default:
		s.WriteString(text)
	}
	if r.More != nil {
		s.WriteString(r.More.html(nil))
	} else {
		s.WriteString(cell(r.Aside, "tc-row__cell"))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// List — tc-list
//
// Роль. Список строк: записи раздела, найденное, выбор в листе.
// Вход. List{Rows, Inset, Compact, Target}.
// Состояния. plain | inset | compact.
// Инварианты. В списке только строки; список печатает фрагмент своего кадра.
// Поведение. Список — одна остановка Tab (владелец 07.11): Tab приходит на одну строку — ту, где фокус был последним, или открытую, или первую — и следующим нажатием уходит за список; галки строк своих остановок не имеют. Стрелки, Home и End ходят по строкам; Enter открывает строку; пробел нажимает её галку. Без сценария — порядок браузера: каждая строка и галка.
// Эталон. кадр `components-12`, фрагмент `.tc-list`.
// Содержимое. Строки любого из типов колонок.
// Чего нет. Заголовков групп и пустого состояния: они стоят рядом со списком.
type List struct {
	Rows    []ListRow
	Inset   bool // set in a card: out to its edges
	Compact bool // rows of 30 px instead of 38 (Compact rows); a phone keeps its own
	// Target is the name of the zone the links and forms inside ask their
	// address for with the live layer, instead of loading the page (Carrel's
	// task 84); without the layer they are plain links and forms.
	Target string
}

// Render prints the list.
func (l List) Render() string { return l.html(nil) }

func (l List) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-list", is("compact", l.Compact), pick(l.Inset, "tc-list--mx-14")) + attr("data-tc-target", l.Target) + `>`)
	for _, r := range l.Rows {
		s.WriteString(r.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// badgesCell prints a cell of a row's badges: those that do not fit whole
// go to the cell's hidden second line, and the hint names them all when
// there is more than one.
func badgesCell(cls string, badges []Badge) string {
	hint := ""
	if len(badges) > 1 {
		var all []string
		for _, b := range badges {
			all = append(all, b.Label)
		}
		hint = strings.Join(all, ", ")
	}
	var parts []string
	for _, b := range badges {
		parts = append(parts, b.html(nil))
	}
	return `<span class="` + cls + `"` + attr("title", hint) + `>` + strings.Join(parts, " ") + `</span>`
}

// rowTag is a row's element: a link when it has an address.
func rowTag(href string) string {
	if href != "" {
		return "a"
	}
	return "div"
}

// Column is one column a person may choose for a list (Columns): what the
// cell holds decides its width and when it goes on a narrow screen.
type Column int

const (
	ColumnName     Column = iota // the record's name: stretches, stays always
	ColumnPhone                  // monospaced, 142 px
	ColumnEmail                  // stretches; goes at 620 px
	ColumnOrg                    // stretches; goes at 800 px
	ColumnBirthday               // monospaced, 84 px; goes at 800 px
	ColumnTags                   // tags; goes at 800 px
	ColumnChanged                // monospaced, 96 px; goes at 800 px
	ColumnSource                 // 140 px; goes at 460 px
	ColumnDue                    // a task's due: monospaced, 96 px
	ColumnComplete               // how much is done: monospaced, 48 px; goes at 800 px
	ColumnPriority               // monospaced, 48 px; goes at 800 px
	ColumnStatus                 // a task's status, an event's attendees or status, in words, 112 px; goes at 620 px
	ColumnExcerpt                // the start of a text: stretches; goes at 620 px
	ColumnDate                   // a record's date, an event's time in the agenda (first, before the name): monospaced, 62 px
	ColumnDetail                 // why a record is in a person's timeline, where an event is: stretches, stays with the name
	ColumnKind                   // event, task, note or file: 56 px; goes at 460 px
	ColumnBadges                 // an event's badges by importance, as TimeRow's: 168 px; goes at 620 px
)

func (c Column) class() string {
	return [...]string{"tc-row__title tc-col--name", "tc-row__excerpt tc-mono tc-col--phone", "tc-row__excerpt tc-col--email",
		"tc-row__excerpt tc-col--org", "tc-row__excerpt tc-mono tc-col--birthday", "tc-row__cell tc-row__tags tc-col--tags",
		"tc-row__cell tc-mono tc-col--changed", "tc-row__cell tc-col--source", "tc-mono tc-col--due", "tc-row__cell tc-mono tc-col--complete",
		"tc-row__cell tc-mono tc-col--priority", "tc-row__cell tc-col--status", "tc-row__excerpt tc-col--excerpt", "tc-row__cell tc-mono tc-col--date",
		"tc-row__excerpt tc-col--detail", "tc-row__cell tc-col--kind", "tc-row__cell tc-col--badges"}[c]
}

// ColumnCell is a cell of a row of chosen columns: its column and what it
// shows — words, tags in the tags column, badges in the badges column (all
// the event's, by importance: the cell shows those that fit whole). Tags are all the record's
// tags: the cell shows the first and counts the rest («работа +1»), its
// hint names them all (Р-53). The one a person filters by goes first — the
// consumer puts it there.
type ColumnCell struct {
	Column Column
	Text   string
	Tags   []RowTag
	Badges []Badge
}

// ColumnRow — tc-row
//
// Роль. Строка списка со столбцами, которые выбрал человек (Columns): набор и порядок — его, ширина — рода столбца.
// Вход. ColumnRow{Tick, Mark, Avatar, Cells, Strong, On, Overdue, Href, Page}.
// Состояния. plain | on | overdue | done | link.
// Инварианты. Столбцы ровны: у всех строк списка один набор в одном порядке; ширина столбца и ступень, на которой он уходит, — его рода; строка печатает фрагмент своего кадра.
// Поведение. Нажатие выбирает строку; строка со ссылкой ведёт на запись; галка дела отмечает его сделанным; с Page, где ширина не рисует панель, нажатие ведёт на страницу записи (задача Carrel 290).
// Эталон. кадр `s-contacts-23`, фрагмент `.tc-row--cols`.
// Содержимое. Галка дела, метка источника, аватар, ячейки столбцов: слова, моноширинное, первый тег со счётчиком остальных, плашки события.
// Чего нет. Своих ширин: столбец задаёт род, а не данные строки.
type ColumnRow struct {
	Tick    *Tick // a task's tick, before the mark
	Mark    Mark
	Avatar  *Avatar
	Cells   []ColumnCell
	Strong  bool // the name in bold: the chosen row
	On      bool
	Overdue bool // a task past its due: the due in the alert colour
	// Href makes the row a link to the record; a row with a Tick cannot be
	// one, its name is the link then, as DueRow's title (Carrel's task 123).
	Href string
	// Page is the record's own page: where the row goes when the page has no
	// panel to open it in (Carrel's task 290).
	Page string
}

// Render prints the row.
func (r ColumnRow) Render() string { return r.html(nil) }

func (r ColumnRow) html(m mix) string {
	var s strings.Builder
	done := r.Tick != nil && r.Tick.Done
	row := pick(r.Tick == nil, r.Href) // a row with its own control is not a link
	s.WriteString(`<` + rowTag(row) + class(m, "tc-row", "tc-row--cols", is("on", r.On), is("overdue", r.Overdue), is("done", done)) + attr("href", row) + `>`)
	if r.Tick != nil {
		s.WriteString(r.Tick.html(nil))
	}
	s.WriteString(r.Mark.html(nil))
	if r.Avatar != nil {
		s.WriteString(r.Avatar.html(nil))
	}
	for _, c := range r.Cells {
		if c.Column == ColumnBadges {
			s.WriteString(badgesCell(c.Column.class(), c.Badges))
			continue
		}
		cls := c.Column.class()
		if c.Column == ColumnName && r.Strong {
			cls = "tc-row__title is-strong tc-col--name"
		}
		tag := "span"
		if c.Column == ColumnName && row == "" && r.Href != "" {
			tag = "a"
		}
		hint := ""
		if c.Column == ColumnTags && len(c.Tags) > 1 {
			var all []string
			for _, t := range c.Tags {
				all = append(all, rowTagLabel(t))
			}
			hint = strings.Join(all, ", ")
		}
		s.WriteString(`<` + tag + ` class="` + cls + `"` + attr("href", pick(tag == "a", r.Href)) + attr("title", hint) + `>`)
		if c.Column == ColumnTags {
			if len(c.Tags) > 0 {
				s.WriteString(rowTagHTML(c.Tags[0], row != ""))
			}
			if n := len(c.Tags) - 1; n > 0 {
				s.WriteString(`<span class="tc-row__more">+` + strconv.Itoa(n) + `</span>`)
			}
		} else {
			s.WriteString(esc(c.Text))
		}
		s.WriteString(`</` + tag + `>`)
	}
	s.WriteString(`</` + rowTag(row) + `>`)
	return s.String()
}

// MailRow — tc-row
//
// Роль. Строка письма в папке ящика: галка, флажок важного, отправитель, тема с началом текста, скрепка, размер, дата.
// Вход. MailRow{Ticked, Send, Important, Flag, Unread, On, Done, From, Topic, Start, Attached, Bytes, Date, Href}.
// Состояния. plain | unread | important | on | done.
// Инварианты. Галка и флажок — кнопки с доступным именем; строка печатает фрагмент своего кадра.
// Поведение. Галка выбирает письмо для действий полосы выбора; флажок отмечает важное сразу; нажатие на отправителя открывает письмо.
// Эталон. кадр `s-mailbox-1`, фрагмент `.tc-row`.
// Содержимое. Галка, флажок, отправитель, тема жирным и начало текста, скрепка со счётом вложений, размер и дата моноширинным.
// Чего нет. Метки коллекции: письмо — не запись коллекции (C-2.1); цвет ящика — в рейле.
type MailRow struct {
	Ticked Ticked
	Send   Send // what pressing the tick sends
	// Important is the letter's flag; Flag is what pressing it sends, a
	// toggle that acts at once.
	Important bool
	Flag      Send
	Unread    bool
	On        bool
	// Done strikes the letter through: moved away on a server that cannot
	// remove only it, it stays listed until the mailbox is cleaned.
	Done     bool
	From     string
	Topic    string
	Start    string // the beginning of the text, after the topic
	Attached int    // attachments: the clip and its count
	Bytes    string // 2.6 MB
	Date     string
	Href     string // the letter
}

// Render prints the row.
func (r MailRow) Render() string { return r.html(nil) }

func (r MailRow) html(m mix) string {
	flag, typ := "Mark as important", "button"
	if r.Important {
		flag = "Important"
	}
	if r.Flag.Form != "" {
		typ = "submit"
	}
	clip := `<span class="tc-row__clip"></span>`
	switch {
	case r.Attached == 1:
		clip = `<span class="tc-row__clip" title="1 attachment"><svg viewBox="0 0 24 24"><use href="#i-clip"/></svg></span>`
	case r.Attached > 1:
		clip = `<span class="tc-row__clip" title="` + strconv.Itoa(r.Attached) + ` attachments"><svg viewBox="0 0 24 24"><use href="#i-clip"/></svg></span>`
	}
	from := "span"
	if r.Href != "" {
		from = "a"
	}
	letter := `<span class="tc-row__letter"><span class="tc-row__topic">` + esc(r.Topic) + `</span>`
	if r.Start != "" {
		letter += ` — ` + esc(r.Start)
	}
	return `<div` + class(m, "tc-row", "tc-row--mailcols", is("unread", r.Unread), is("on", r.On), is("done", r.Done)) + `>` + r.Ticked.box(r.Topic, r.Send) +
		`<button` + class(nil, "tc-row__flag", is("on", r.Important)) + attr("title", flag) + attr("aria-label", flag) + ` type="` + typ + `"` + attr("form", r.Flag.Form) + attr("name", r.Flag.Name) + attr("value", r.Flag.Value) +
		attr("aria-pressed", map[bool]string{true: "true", false: "false"}[r.Important]) + `><svg viewBox="0 0 24 24"><use href="#i-flag"/></svg></button>` +
		`<` + from + ` class="tc-row__title"` + attr("href", r.Href) + `>` + esc(r.From) + `</` + from + `>` + letter + `</span>` + clip +
		cell(r.Bytes, "tc-row__cell", "tc-mono", "tc-row__size") + cell(r.Date, "tc-mono", "tc-row__date") + `</div>`
}
