package widget

import (
	"strings"

	"carrel.dev/typecase/layout"
)

// List — tc-list
//
// Роль. Ряд строк, разделённых линовкой, с рубрикаторами между группами.
// Вход. List{ Groups, Density, Empty }.
// Состояния. normal | empty.
// Инварианты. Пустой список показывает действие; плотность от оболочки.
// Поведение. Ходьба стрелками по строкам, Enter открывает выделенную.
// Геометрия. Высота строки 38px / 52px; без рамки, линовка --tc-hair.
// Содержимое. tc-rubric, строки tc-row.
// Чего нет. Не сортирует и не фильтрует.
type List struct {
	Groups  []ListGroup
	Density Density
	Empty   Empty
}

func (l List) Render() string {
	if len(l.Groups) == 0 {
		return `<div class="tc-list tc-list--empty">` + l.Empty.Render() + `</div>`
	}
	dens := l.Density
	if dens == "" {
		dens = DensityNormal
	}
	var groups string
	for _, g := range l.Groups {
		var rows string
		for _, r := range g.Rows {
			rows += r.Render()
		}
		groups += g.Rubric.Render() + rows
	}
	return `<div` + classAttr("tc-list", "tc-list--"+string(dens)) + ` role="list" data-tc-list>` + groups + `</div>`
}

// Row — tc-row
//
// Роль. Одна запись в списке: ведущая метка, содержимое, хвост.
// Вход. Row{ Href, Mark, Cells, Badges, Menu, Selected, State, Variant }.
// Состояния. normal | selected | muted | alert.
// Инварианты. Метка — полоса слева; меню у flex-строки; сеточная — только в List.
// Корень строки не ссылка: при Href ссылкой становится главный текст (title,
// у meta и tickcols — name), растянутый на всю строку; галка, теги и меню лежат
// поверх и нажимаются отдельно.
// Поведение. Наведение, выделение, меню, Enter.
// Геометрия. Сеточные варианты — таблица layout; дети = слоты.
// Содержимое. Ячейки, tc-badge; tc-menu у flex-вариантов.
// Чего нет. Нет перетаскивания и вложенных строк.
type Row struct {
	Href     string
	Mark     Mark
	Cells    RowCells
	Badges   []Badge
	Menu     []MenuItem
	Selected bool
	State    RowState
	Variant  RowVariant
}

func (r Row) Render() string {
	if g, ok := rowGrid(r.Variant); ok {
		return renderGridRow(r, g)
	}
	cls := classes("tc-row", rowVariantClass(r.Variant))
	if r.Selected || r.State == RowSelected {
		cls = classes(cls, "is-selected")
	}
	switch r.State {
	case RowMuted:
		cls = classes(cls, "is-muted")
	case RowAlert:
		cls = classes(cls, "is-alert")
	}
	mark := r.Mark.Render()
	var badges string
	for _, b := range r.Badges {
		badges += b.Render()
	}
	cells := renderRowCells(r.Variant, r.Cells, r.Href)
	menu := ""
	if len(r.Menu) > 0 {
		menu = Menu{Label: labels.Actions, Items: r.Menu}.Render()
	}
	inner := mark + `<div class="tc-row__cells">` + cells + `</div><div class="tc-row__badges">` + badges + `</div><div class="tc-row__tail">` + menu + `</div>`
	return `<div` + classAttr(cls) + ` role="listitem" data-tc-row>` + inner + `</div>`
}

// rowMain prints the row's main text node. With href it is the row link.
func rowMain(slotClass, inner, href string) string {
	if href == "" {
		return `<span class="` + slotClass + `">` + inner + `</span>`
	}
	return `<a href="` + esc(href) + `" class="` + slotClass + ` tc-row__link">` + inner + `</a>`
}

// rowMainSlot is the slot that carries the link of a grid row.
func rowMainSlot(g layout.Grid) string {
	for _, slot := range []string{"title", "name"} {
		if layout.HasSlot(g, slot) {
			return slot
		}
	}
	return ""
}

func renderGridRow(r Row, g layout.Grid) string {
	cls := classes("tc-row", rowVariantClass(r.Variant))
	if r.Selected || r.State == RowSelected {
		cls = classes(cls, "is-selected")
	}
	switch r.State {
	case RowMuted:
		cls = classes(cls, "is-muted")
	case RowAlert:
		cls = classes(cls, "is-alert")
	}
	main := rowMainSlot(g)
	var inner string
	for _, tr := range g.Tracks {
		href := ""
		if tr.Slot == main {
			href = r.Href
		}
		inner += rowSlot(tr.Slot, r, href)
	}
	return `<div` + classAttr(cls) + ` role="listitem" data-tc-row>` + inner + `</div>`
}

func rowGrid(v RowVariant) (layout.Grid, bool) {
	switch v {
	case RowExcerpt:
		return layout.RowExcerpt, true
	case RowAvatarCols:
		return layout.RowAvatarCols, true
	case RowTickCols:
		return layout.RowTickCols, true
	case RowTimeCols:
		return layout.RowTimeCols, true
	case RowTickMarkCols:
		return layout.RowTickMarkCols, true
	case RowFind:
		return layout.RowFind, true
	default:
		return layout.Grid{}, false
	}
}

func rowSlot(slot string, r Row, href string) string {
	switch slot {
	case "mark":
		html := r.Mark.Render()
		if html == "" {
			return `<span class="tc-mark"></span>`
		}
		return html
	case "tick":
		label := r.Cells.Name
		if label == "" {
			label = r.Cells.Title
		}
		if label == "" {
			label = labels.Toggle
		}
		return Tick{Done: r.Cells.Done, Label: label}.Render()
	case "avatar":
		av := r.Cells.Avatar
		if av.Initials == "" && av.Image == "" {
			av = Avatar{Initials: initials(r.Cells.Title)}
		}
		return av.Render()
	case "title":
		return rowMain("tc-row__title", esc(r.Cells.Title), href)
	case "excerpt":
		return `<span class="tc-row__excerpt">` + esc(r.Cells.Excerpt) + `</span>`
	case "tags":
		var tags string
		for _, t := range r.Cells.Tags {
			tags += t.Render()
		}
		for _, b := range r.Badges {
			tags += b.Render()
		}
		return `<span class="tc-row__tags">` + tags + `</span>`
	case "date":
		return `<span class="tc-row__date">` + esc(r.Cells.Date) + `</span>`
	case "source":
		return `<span class="tc-row__source">` + esc(r.Cells.Source) + `</span>`
	case "subtitle":
		return `<span class="tc-row__subtitle">` + esc(r.Cells.Subtitle) + `</span>`
	case "detail":
		return `<span class="tc-row__detail">` + esc(r.Cells.Detail) + `</span>`
	case "caption":
		return `<span class="tc-row__caption">` + esc(r.Cells.Caption) + `</span>`
	case "name":
		return rowMain("tc-row__name", rowCellIcon(r.Cells.Icon)+esc(r.Cells.Name), href)
	case "size":
		return `<span class="tc-row__size">` + esc(r.Cells.Size) + `</span>`
	case "time":
		return `<span class="tc-row__time">` + esc(r.Cells.Time) + `</span>`
	case "lead":
		return `<span class="tc-row__lead">` + esc(r.Cells.Lead) + `</span>`
	case "due":
		return `<span class="tc-row__due">` + esc(r.Cells.Due) + `</span>`
	case "progress":
		if r.Cells.Progress != nil {
			return `<progress class="tc-row__progress" max="100" value="` + itoa(*r.Cells.Progress) + `"` + attr("aria-label", labels.Progress) + `></progress>`
		}
		return `<span class="tc-row__progress"></span>`
	default:
		return ""
	}
}

func rowCellIcon(icon string) string {
	if icon == "" {
		return ""
	}
	if strings.HasPrefix(icon, "tc-i-") {
		return renderNavIcon(icon)
	}
	return `<span class="tc-row__icon" aria-hidden="true">` + esc(icon) + `</span>`
}

func rowVariantClass(v RowVariant) string {
	if v == "" || v == RowDefault {
		return ""
	}
	return "tc-row--" + string(v)
}

func renderRowCells(v RowVariant, c RowCells, href string) string {
	title := rowMain("tc-row__title", esc(c.Title), href)
	switch v {
	case RowAvatar:
		av := c.Avatar
		if av.Initials == "" && av.Image == "" {
			av = Avatar{Initials: initials(c.Title)}
		}
		return av.Render() + `<div class="tc-row__text">` + title + `<span class="tc-row__subtitle">` + esc(c.Subtitle) + `</span></div>`
	case RowLead:
		return `<span class="tc-row__lead">` + esc(c.Lead) + `</span><div class="tc-row__text">` + title + `<span class="tc-row__subtitle">` + esc(c.Subtitle) + `</span></div>`
	case RowTick:
		tick := Tick{Done: c.Done, Label: c.Title}.Render()
		return tick + title + `<span class="tc-row__due">` + esc(c.Due) + `</span>`
	case RowMeta:
		icon := esc(c.Icon)
		if strings.HasPrefix(c.Icon, "tc-i-") {
			icon = renderNavIcon(c.Icon)
		}
		return `<span class="tc-row__icon" aria-hidden="true">` + icon + `</span>` + rowMain("tc-row__name", esc(c.Name), href) + `<span class="tc-row__size">` + esc(c.Size) + `</span><span class="tc-row__time">` + esc(c.Time) + `</span>`
	default:
		return title + `<span class="tc-row__subtitle">` + esc(c.Subtitle) + `</span>`
	}
}

// TickRow — tc-tickrow
//
// Роль. Отмечаемая строка: галка, метка, подпись, счётчик, только чтение.
// Вход. TickRow{ Href, Label, Mark, Count, Tick, ReadOnly, State, Current, Menu }.
// Состояния. ok | read-only | unavailable | gone; галка on | off | partial.
// Инварианты. Три состояния галки — одно поле; счётчик пуст при отсутствии;
// заглушка метки всегда при Hold; ссылка коробки не заводит: при Href ссылкой
// становится подпись, растянутая на строку, галка и меню лежат поверх.
// Поведение. Галка и подпись — разные зоны.
// Геометрия. Таблица layout.TickRow.
// Содержимое. tc-mark, tc-num, tc-menu.
// Чего нет. Не хранит состояние.
type TickRow struct {
	Href     string
	Label    string
	Mark     Mark
	Count    *int
	Tick     TickState
	ReadOnly bool
	State    TickRowState
	Current  bool
	Menu     []MenuItem
}

func (t TickRow) Render() string {
	cls := "tc-tickrow"
	if t.Current {
		cls = classes(cls, "is-current")
	}
	switch t.State {
	case TickRowReadOnly:
		cls = classes(cls, "is-read-only")
	case TickRowUnavailable:
		cls = classes(cls, "is-unavailable")
	case TickRowGone:
		cls = classes(cls, "is-gone")
	}
	tickCls := "tc-tickrow__tick"
	switch t.Tick {
	case TickOn:
		tickCls = classes(tickCls, "is-on")
	case TickPartial:
		tickCls = classes(tickCls, "is-partial")
	case TickDisabled:
		tickCls = classes(tickCls, "is-disabled")
	}
	menu := ""
	if len(t.Menu) > 0 {
		menu = Menu{Label: labels.Actions, Items: t.Menu}.Render()
	}
	stateWord := ""
	if t.State == TickRowUnavailable {
		stateWord = `<span class="tc-tickrow__state">` + esc(labels.Unavailable) + `</span>`
	}
	var slots string
	for _, tr := range layout.TickRow.Tracks {
		slots += tickRowSlot(tr.Slot, t, tickCls)
	}
	return `<div` + classAttr(cls) + ` data-tc-tickrow>` + slots + stateWord + Num{Value: t.Count}.Render() + menu + `</div>`
}

func tickRowSlot(slot string, t TickRow, tickCls string) string {
	switch slot {
	case "tick":
		return `<button type="button" class="` + tickCls + `" aria-label="` + esc(labels.Toggle) + ` ` + esc(t.Label) + `"></button>`
	case "mark":
		html := t.Mark.Render()
		if html == "" {
			return `<span class="tc-mark"></span>`
		}
		return html
	case "label":
		if t.Href != "" {
			return `<a href="` + esc(t.Href) + `" class="tc-tickrow__label tc-tickrow__link">` + esc(t.Label) + `</a>`
		}
		return `<span class="tc-tickrow__label">` + esc(t.Label) + `</span>`
	default:
		return ""
	}
}

// Rubric — tc-rubric
//
// Роль. Мелкая надпись с разрядкой, открывающая группу.
// Вход. Rubric{ Text, Count, Action }.
// Состояния. нет.
// Инварианты. Не заголовок раздела и не кликается.
// Поведение. нет.
// Геометрия. В списке — как .m-group: padding 15 / gutter / 5, низ --tc-rule.
// Содержимое. tc-num, tc-btn.
// Чего нет. Нет уровней.
type Rubric struct {
	Text   string
	Count  *int
	Action *Action
}

func (r Rubric) Render() string {
	action := ""
	if r.Action != nil {
		a := *r.Action
		if a.Size == "" {
			a.Size = ButtonSizeTable
		}
		action = `<span class="tc-rubric__action">` + a.Render() + `</span>`
	}
	return `<div class="tc-rubric">` + `<span class="tc-rubric__text">` + esc(r.Text) + `</span>` + Num{Value: r.Count}.Render() + action + `</div>`
}
