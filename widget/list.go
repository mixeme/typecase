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
	if r.Variant == RowExcerpt {
		return renderGridRow(r, layout.RowExcerpt)
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
	cells := renderRowCells(r.Variant, r.Cells)
	menu := ""
	if len(r.Menu) > 0 {
		menu = Menu{Label: labels.Actions, Items: r.Menu}.Render()
	}
	inner := mark + `<div class="tc-row__cells">` + cells + `</div><div class="tc-row__badges">` + badges + `</div><div class="tc-row__tail">` + menu + `</div>`
	if r.Href != "" {
		return `<a href="` + esc(r.Href) + `"` + classAttr(cls) + ` role="listitem" data-tc-row>` + inner + `</a>`
	}
	return `<div` + classAttr(cls) + ` role="listitem" data-tc-row>` + inner + `</div>`
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
	var inner string
	for _, tr := range g.Tracks {
		inner += rowSlot(tr.Slot, r)
	}
	if r.Href != "" {
		return `<a href="` + esc(r.Href) + `"` + classAttr(cls) + ` role="listitem" data-tc-row>` + inner + `</a>`
	}
	return `<div` + classAttr(cls) + ` role="listitem" data-tc-row>` + inner + `</div>`
}

func rowSlot(slot string, r Row) string {
	switch slot {
	case "mark":
		html := r.Mark.Render()
		if html == "" {
			return `<span class="tc-mark"></span>`
		}
		return html
	case "title":
		return `<span class="tc-row__title">` + esc(r.Cells.Title) + `</span>`
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
	default:
		return ""
	}
}

func rowVariantClass(v RowVariant) string {
	if v == "" || v == RowDefault {
		return ""
	}
	return "tc-row--" + string(v)
}

func renderRowCells(v RowVariant, c RowCells) string {
	switch v {
	case RowAvatar:
		av := c.Avatar
		if av.Initials == "" && av.Image == "" {
			av = Avatar{Initials: initials(c.Title)}
		}
		return av.Render() + `<div class="tc-row__text"><span class="tc-row__title">` + esc(c.Title) + `</span><span class="tc-row__subtitle">` + esc(c.Subtitle) + `</span></div>`
	case RowLead:
		return `<span class="tc-row__lead">` + esc(c.Lead) + `</span><div class="tc-row__text"><span class="tc-row__title">` + esc(c.Title) + `</span><span class="tc-row__subtitle">` + esc(c.Subtitle) + `</span></div>`
	case RowTick:
		tick := Tick{Done: c.Done, Label: c.Title}.Render()
		return tick + `<span class="tc-row__title">` + esc(c.Title) + `</span><span class="tc-row__due">` + esc(c.Due) + `</span>`
	case RowMeta:
		icon := esc(c.Icon)
		if strings.HasPrefix(c.Icon, "tc-i-") {
			icon = renderNavIcon(c.Icon)
		}
		return `<span class="tc-row__icon" aria-hidden="true">` + icon + `</span><span class="tc-row__name">` + esc(c.Name) + `</span><span class="tc-row__size">` + esc(c.Size) + `</span><span class="tc-row__time">` + esc(c.Time) + `</span>`
	case RowFind:
		return `<span class="tc-row__icon" aria-hidden="true">` + esc(c.Icon) + `</span><div class="tc-row__find"><span class="tc-row__title">` + esc(c.Title) + `</span><span class="tc-row__detail">` + esc(c.Detail) + `</span><span class="tc-row__caption">` + esc(c.Caption) + `</span></div>`
	default:
		return `<span class="tc-row__title">` + esc(c.Title) + `</span><span class="tc-row__subtitle">` + esc(c.Subtitle) + `</span>`
	}
}

// TickRow — tc-tickrow
//
// Роль. Отмечаемая строка: галка, метка, подпись, счётчик, только чтение.
// Вход. TickRow{ Href, Label, Mark, Count, Tick, ReadOnly, State, Current, Menu }.
// Состояния. ok | read-only | unavailable | gone; галка on | off | partial.
// Инварианты. Три состояния галки — одно поле; счётчик пуст при отсутствии;
// заглушка метки всегда при Hold; ссылка коробки не заводит.
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
	inner := slots + stateWord + Num{Value: t.Count}.Render() + menu
	if t.Href != "" {
		return `<div` + classAttr(cls) + ` data-tc-tickrow><a href="` + esc(t.Href) + `" class="tc-tickrow__link">` + inner + `</a></div>`
	}
	return `<div` + classAttr(cls) + ` data-tc-tickrow>` + inner + `</div>`
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
// Геометрия. Прописные, разрядка, приглушённый цвет.
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
