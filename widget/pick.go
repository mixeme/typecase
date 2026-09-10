package widget

import "strings"

// PickItem is one option in a tc-pick menu.
type PickItem struct {
	Label string
	Hint  string
	Icon  string
}

// Pick — tc-pick
//
// Роль. Поиск и выбор связей: выбранные значки, ввод и список кандидатов.
// Вход. Pick{ Name, ID, Label, Chosen []string, Query, Options []PickItem, Manual, Open }.
// Состояния. closed | open.
// Инварианты. Выбранное снимается крестиком; ввод UID вручную — последний пункт.
// Поведение. Ввод фильтрует список; выбор добавляет значок.
// Геометрия. Многострочная зона ввода с переносом значков.
// Содержимое. Значки выбранного, список кандидатов.
// Чего нет. Нет запросов к сети: кандидаты приходят в Options.
type Pick struct {
	Name    string
	ID      string
	Label   string
	Chosen  []string
	Query   string
	Options []PickItem
	Manual  string
	Open    bool
}

func (p Pick) Render() string {
	cls := "tc-pick"
	if p.Open {
		cls = classes(cls, "is-open")
	}
	id := p.ID
	if id == "" {
		id = uniqueID("tc-pick")
	}
	var tags string
	for _, c := range p.Chosen {
		tags += `<span class="tc-pick__tag"><span class="tc-pick__tag-label">` + esc(c) + `</span><button type="button" class="tc-pick__tag-remove"` + attr("aria-label", labels.Remove+" "+c) + `><span aria-hidden="true">×</span></button></span>`
	}
	ph := ""
	if len(p.Chosen) == 0 && p.Query == "" {
		ph = attr("placeholder", p.Label)
	}
	al := attr("aria-label", p.Label)
	menuID := id + "-menu"
	controls := ""
	var items string
	for _, o := range p.Options {
		ic := renderIcon(o.Icon)
		hint := ""
		if o.Hint != "" {
			hint = `<span class="tc-pick__hint">` + esc(o.Hint) + `</span>`
		}
		items += `<li class="tc-pick__item" role="option"` + attr("data-tc-pick-value", o.Label) + `>` + ic + esc(o.Label) + hint + `</li>`
	}
	if p.Manual != "" {
		items += `<li class="tc-pick__sep" role="separator"></li><li class="tc-pick__item tc-pick__item--manual" role="option">` + esc(p.Manual) + `</li>`
	}
	menu := ""
	if items != "" {
		hidden := ""
		if !p.Open {
			hidden = " hidden"
		}
		menu = `<ul class="tc-pick__menu" role="listbox"` + attr("id", menuID) + hidden + `>` + items + `</ul>`
		controls = attr("aria-controls", menuID)
	}
	query := `<input type="text" class="tc-pick__query" autocomplete="off"` + attr("id", id) + al + ph + attr("value", p.Query) + ` role="combobox" aria-autocomplete="list"` + ariaBool("aria-expanded", p.Open) + controls + `>`
	return `<div` + classAttr(cls) + ` data-tc-pick` + attr("data-tc-remove", labels.Remove) + `><div class="tc-pick__box">` + tags + query + `</div>` + menu + `<input type="hidden" name="` + esc(p.Name) + `" value="` + esc(strings.Join(p.Chosen, ",")) + `"></div>`
}
