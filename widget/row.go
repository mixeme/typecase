package widget

import "strings"

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
	s := `<span` + class(nil, append([]string{"tc-row__title", is("strong", strong)}, own...)...) + `>` + inlines(parts)
	if badge != nil {
		s += " " + badge.html(nil)
	}
	return s + `</span>`
}

// source prints the source cell: where the record lives.
func source(text, hint string, badge *Badge) string {
	s := `<span class="tc-row__cell tc-row__source"` + attr("title", hint) + `>` + esc(text)
	if badge != nil {
		s += " " + badge.html(nil)
	}
	return s + `</span>`
}

func tags(items []RowTag) string {
	var s strings.Builder
	s.WriteString(`<span class="tc-row__cell tc-row__tags">`)
	for _, t := range items {
		s.WriteString(t.html(nil))
	}
	s.WriteString(`</span>`)
	return s.String()
}

// inset is the variant of a row set in a card: 14 px at the sides.
func inset(kind string, on bool) string {
	return pick(on, "tc-row--"+kind+"--padl14-padr14")
}

// TimeRow — tc-row
//
// Роль. Строка со временем: событие дня, приглашение, запись журнала.
// Вход. TimeRow{Mark, Time, Title, Strong, Detail, Caption, Source, On, Past, Inset, Short, Href}.
// Состояния. plain | on | caption | past | inset | found.
// Инварианты. Колонки строки — её тип: у строки времени нет аватара и срока; метка источника — первой; строка печатает фрагмент своего кадра.
// Поведение. Нажатие открывает запись; стрелки ходят по строкам списка.
// Эталон. кадр `s-agenda-1`, фрагмент `.tc-row`.
// Содержимое. Метка, время моноширинным, название (с найденным жирным), место, бейдж, источник.
// Чего нет. Своих кнопок: действия с записью — в панели и в меню «⋯» шапки.
type TimeRow struct {
	Mark    Mark
	Time    string
	Title   []Inline
	Strong  bool // not seen yet
	Detail  string
	Caption *Badge
	Source  string
	On      bool
	Past    bool // gone by: faded
	Inset   bool
	Short   bool   // the lead is a time of day only: its column as narrow as the time (search results)
	Href    string // the row is a link to its record
}

// Render prints the row.
func (r TimeRow) Render() string { return r.html(nil) }

func (r TimeRow) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<` + rowTag(r.Href) + class(m, "tc-row", "tc-row--timecols", is("on", r.On), inset("timecols", r.Inset), pick(r.Past, "tc-row--timecols--op-7"), pick(r.Short, "tc-row--timecols--time")) + attr("href", r.Href) + `>`)
	s.WriteString(r.Mark.html(nil))
	s.WriteString(cell(r.Time, "tc-mono", "tc-row__lead"))
	s.WriteString(title(r.Title, r.Strong, nil))
	s.WriteString(cell(r.Detail, "tc-row__excerpt", "tc-row__detail"))
	s.WriteString(`<span class="tc-row__cell tc-row__caption">`)
	if r.Caption != nil {
		s.WriteString(r.Caption.html(nil))
	}
	s.WriteString(`</span>`)
	s.WriteString(source(r.Source, "", nil))
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
// Эталон. кадр `m-states-3`, фрагмент `.tc-row`.
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
// Вход. FaceRow{Mark, Avatar, Title, Strong, TitleBadge, Subtitle, Detail, Caption, Source, SourceBadge, On, Inset, Href}.
// Состояния. plain | on | strong | badge | inset | found.
// Инварианты. Колонки строки — её тип: у строки человека есть аватар и нет времени; метка источника — первой; строка печатает фрагмент своего кадра.
// Поведение. Нажатие открывает карточку; стрелки ходят по строкам списка.
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
	SourceBadge *Badge
	On          bool
	Inset       bool
	Href        string // the row is a link to its record
}

// Render prints the row.
func (r FaceRow) Render() string { return r.html(nil) }

func (r FaceRow) html(m mix) string {
	return `<` + rowTag(r.Href) + class(m, "tc-row", "tc-row--avatarcols", is("on", r.On), inset("avatarcols", r.Inset)) + attr("href", r.Href) + `>` +
		r.Mark.html(nil) + r.Avatar.html(nil) + title(r.Title, r.Strong, r.TitleBadge) +
		cell(r.Subtitle, "tc-row__excerpt", "tc-mono", "tc-row__subtitle") + cell(r.Detail, "tc-row__excerpt", "tc-row__detail") +
		cell(r.Caption, "tc-row__excerpt", "tc-row__caption") + source(r.Source, "", r.SourceBadge) + `</` + rowTag(r.Href) + `>`
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
// Роль. Строка текста: название, начало текста, теги, дата, источник.
// Вход. ExcerptRow{Mark, Title, Strong, TitleBadge, Excerpt, Tags, Date, Source, SourceHint, On, Inset, Href}.
// Состояния. plain | on | badge | inset | found | link.
// Инварианты. Колонки строки — её тип; метка источника — первой; строка печатает фрагмент своего кадра.
// Поведение. Нажатие открывает текст; стрелки ходят по строкам списка.
// Эталон. кадр `s-notes-1`, фрагмент `.tc-row`.
// Содержимое. Метка, название (с найденным жирным и бейджем), начало текста, теги и бейджи, дата моноширинным, источник (с подсказкой, когда их несколько).
// Чего нет. Всего текста: строка показывает начало, текст открывается.
type ExcerptRow struct {
	Mark       Mark
	Title      []Inline
	Strong     bool
	TitleBadge *Badge
	Excerpt    string
	Tags       []RowTag
	Date       string
	Source     string
	SourceHint string // the sources in full when there are several
	On         bool
	Inset      bool
	Href       string // the row is a link to its record
}

// Render prints the row.
func (r ExcerptRow) Render() string { return r.html(nil) }

func (r ExcerptRow) html(m mix) string {
	return `<` + rowTag(r.Href) + class(m, "tc-row", "tc-row--excerpt", is("on", r.On), inset("excerpt", r.Inset)) + attr("href", r.Href) + `>` + r.Mark.html(nil) +
		title(r.Title, r.Strong, r.TitleBadge) + cell(r.Excerpt, "tc-row__excerpt") + tags(r.Tags) +
		cell(r.Date, "tc-mono", "tc-row__date") + source(r.Source, r.SourceHint, nil) + `</` + rowTag(r.Href) + `>`
}

// BriefRow — tc-row
//
// Роль. Короткая строка текста в выборе: название, дата, где лежит.
// Вход. BriefRow{Mark, Title, Date, Where, On, Href}.
// Состояния. plain | on.
// Инварианты. Колонок три, и это её тип; строка печатает фрагмент своего кадра.
// Поведение. Нажатие выбирает строку.
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
}

// Render prints the row.
func (r BriefRow) Render() string { return r.html(nil) }

func (r BriefRow) html(m mix) string {
	return `<` + rowTag(r.Href) + class(m, "tc-row", "tc-row--excerpt", is("on", r.On), "tc-row--excerpt--cols3-62-148") + attr("href", r.Href) + `>` + r.Mark.html(nil) +
		cell(r.Title, "tc-row__title", "is-strong") + r.Date.html(nil) + cell(r.Where, "tc-row__cell", "tc-row__tags") + `</` + rowTag(r.Href) + `>`
}

// DueRow — tc-row
//
// Роль. Строка дела: галка, название, теги, срок, сколько сделано, источник.
// Вход. DueRow{Tick, Mark, Title, Strong, TitleBadge, Tags, Due, Progress, Source, SourceHint, On, Overdue, Inset}.
// Состояния. plain | on | overdue | done | inset | found.
// Инварианты. Сделано ли дело, говорит одна галка: строка берёт это из неё; колонки строки — её тип; строка печатает фрагмент своего кадра.
// Поведение. Галка отмечает дело сделанным; нажатие на строку открывает его.
// Эталон. кадр `s-tasks-1`, фрагмент `.tc-row`.
// Содержимое. Галка, метка, название (с найденным жирным и бейджем), теги, срок и доля моноширинным, источник.
// Чего нет. Подзадач: они — в панели дела.
type DueRow struct {
	Tick       Tick
	Mark       Mark
	Title      []Inline
	Strong     bool
	TitleBadge *Badge
	Tags       []RowTag
	Due        string
	Progress   string
	Source     string
	SourceHint string
	On         bool
	Overdue    bool
	Inset      bool
}

// Render prints the row.
func (r DueRow) Render() string { return r.html(nil) }

func (r DueRow) html(m mix) string {
	return `<div` + class(m, "tc-row", "tc-row--tickmarkcols", is("on", r.On), is("overdue", r.Overdue), is("done", r.Tick.Done), inset("tickmarkcols", r.Inset)) + `>` +
		r.Tick.html(nil) + r.Mark.html(nil) + title(r.Title, r.Strong, r.TitleBadge) + tags(r.Tags) +
		cell(r.Due, "tc-mono", "tc-row__due") + cell(r.Progress, "tc-row__cell", "tc-mono", "tc-row__progress") +
		source(r.Source, r.SourceHint, nil) + `</div>`
}

// SizeRow — tc-row
//
// Роль. Строка вещи с размером: галка, значок и имя, размер, когда менялась, права, источник.
// Вход. SizeRow{Ticked, Icon, Name, Bytes, Time, Access, Source}.
// Состояния. found.
// Инварианты. Галка — кнопка с доступным именем; колонки строки — её тип; строка печатает фрагмент своего кадра.
// Поведение. Галка выбирает вещь для действий шапки; нажатие на имя открывает её.
// Эталон. кадр `s-find-1`, фрагмент `.tc-row`.
// Содержимое. Галка, значок рода, имя (с найденным жирным), размер, время и права моноширинным, источник.
// Чего нет. Метки источника: источник назван словами.
type SizeRow struct {
	Ticked Ticked
	Icon   Icon
	Name   []Inline
	Bytes  string // 240 KB
	Time   string
	Access string // rw, ro
	Source string
}

// Render prints the row.
func (r SizeRow) Render() string { return r.html(nil) }

func (r SizeRow) html(m mix) string {
	return `<div` + class(m, "tc-row", "tc-row--tickcols") + `>` + r.Ticked.box(plain(r.Name)) +
		`<span class="tc-row__title tc-row__name"><svg>` + r.Icon.use() + `</svg><span class="tc-row__file">` + inlines(r.Name) + `</span></span>` +
		cell(r.Bytes, "tc-row__excerpt", "tc-mono", "tc-row__size") + cell(r.Time, "tc-row__excerpt", "tc-mono", "tc-row__time") +
		cell(r.Access, "tc-row__cell", "tc-mono", "tc-row__detail") + source(r.Source, "", nil) + `</div>`
}

// SlimRow — tc-row
//
// Роль. Узкая строка выбора: галка, метка, что выбрать и что в нём, род.
// Вход. SlimRow{Ticked, Mark, Title, Inside, Aside}.
// Состояния. plain | off.
// Инварианты. Галка — кнопка с доступным именем; строка печатает фрагмент своего кадра.
// Поведение. Галка выбирает строку.
// Эталон. кадр `s-plan-1`, фрагмент `.tc-row`.
// Содержимое. Галка, метка, название с тем, что в нём, род справа.
// Чего нет. Источника и даты.
type SlimRow struct {
	Ticked Ticked
	Mark   Mark
	Title  string
	Inside string // what is in it, after the title
	Aside  string
}

// Render prints the row.
func (r SlimRow) Render() string { return r.html(nil) }

func (r SlimRow) html(m mix) string {
	return `<div` + class(m, "tc-row", "tc-row--slim") + `>` + r.Ticked.box(r.Title) + r.Mark.html(nil) +
		`<span class="tc-row__title">` + esc(r.Title) + ` ` + cell(r.Inside, "tc-row__cell") + `</span>` + cell(r.Aside, "tc-row__cell") + `</div>`
}

// TickBox is a tick box in front of a phone's row.
type TickBox struct {
	Ticked Ticked
}

// PhoneRow — tc-row
//
// Роль. Строка списка на телефоне: две строки текста и одна пометка справа.
// Вход. PhoneRow{Mark, Pick, Avatar, Icon, Title, TitleBadge, Sub, SubMono, Aside, More, On}.
// Состояния. plain | on | avatar | badge | pick.
// Инварианты. Спереди — метка или галка, не обе; галка — кнопка с доступным именем; строка печатает фрагмент своего кадра.
// Поведение. Нажатие открывает запись; галка выбирает её.
// Эталон. кадр `m-phone-4`, фрагмент `.tc-row`.
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
		s.WriteString(r.Pick.Ticked.box(r.Title, "tc-tickrow__tick--ml2"))
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
// Вход. List{Rows, Inset, Compact}.
// Состояния. plain | inset | compact.
// Инварианты. В списке только строки; список печатает фрагмент своего кадра.
// Поведение. Стрелки ходят по строкам; Enter открывает строку.
// Эталон. кадр `components-12`, фрагмент `.tc-list`.
// Содержимое. Строки любого из типов колонок.
// Чего нет. Заголовков групп и пустого состояния: они стоят рядом со списком.
type List struct {
	Rows    []ListRow
	Inset   bool // set in a card: out to its edges
	Compact bool // rows of 30 px instead of 38 (Compact rows); a phone keeps its own
}

// Render prints the list.
func (l List) Render() string { return l.html(nil) }

func (l List) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-list", is("compact", l.Compact), pick(l.Inset, "tc-list--mx-14")) + `>`)
	for _, r := range l.Rows {
		s.WriteString(r.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
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
)

func (c Column) class() string {
	return [...]string{"tc-row__title tc-col--name", "tc-row__excerpt tc-mono tc-col--phone", "tc-row__excerpt tc-col--email",
		"tc-row__excerpt tc-col--org", "tc-row__excerpt tc-mono tc-col--birthday", "tc-row__cell tc-row__tags tc-col--tags",
		"tc-row__cell tc-mono tc-col--changed", "tc-row__cell tc-col--source"}[c]
}

// ColumnCell is a cell of a row of chosen columns: its column and what it
// shows — words, or tags in the tags column.
type ColumnCell struct {
	Column Column
	Text   string
	Tags   []RowTag
}

// ColumnRow — tc-row
//
// Роль. Строка списка со столбцами, которые выбрал человек (Columns): набор и порядок — его, ширина — рода столбца.
// Вход. ColumnRow{Mark, Avatar, Cells, Strong, On, Href}.
// Состояния. plain | on | link.
// Инварианты. Столбцы ровны: у всех строк списка один набор в одном порядке; ширина столбца и ступень, на которой он уходит, — его рода; строка печатает фрагмент своего кадра.
// Поведение. Нажатие выбирает строку; строка со ссылкой ведёт на запись.
// Эталон. кадр `s-contacts-23`, фрагмент `.tc-row--cols`.
// Содержимое. Метка источника, аватар, ячейки столбцов: слова, моноширинное, теги.
// Чего нет. Своих ширин: столбец задаёт род, а не данные строки.
type ColumnRow struct {
	Mark   Mark
	Avatar *Avatar
	Cells  []ColumnCell
	Strong bool // the name in bold: the chosen row
	On     bool
	Href   string
}

// Render prints the row.
func (r ColumnRow) Render() string { return r.html(nil) }

func (r ColumnRow) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<` + rowTag(r.Href) + class(m, "tc-row", "tc-row--cols", is("on", r.On)) + attr("href", r.Href) + `>` + r.Mark.html(nil))
	if r.Avatar != nil {
		s.WriteString(r.Avatar.html(nil))
	}
	for _, c := range r.Cells {
		cls := c.Column.class()
		if c.Column == ColumnName && r.Strong {
			cls = "tc-row__title is-strong tc-col--name"
		}
		s.WriteString(`<span class="` + cls + `">`)
		if c.Column == ColumnTags {
			for _, t := range c.Tags {
				s.WriteString(t.html(nil))
			}
		} else {
			s.WriteString(esc(c.Text))
		}
		s.WriteString(`</span>`)
	}
	s.WriteString(`</` + rowTag(r.Href) + `>`)
	return s.String()
}
