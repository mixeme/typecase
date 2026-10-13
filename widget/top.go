package widget

import "strings"

// TopAction is what stands at the end of the top bar.
type TopAction interface {
	part
	topAction()
}

func (Button) topAction()   {}
func (UserMenu) topAction() {}
func (TopDrop) topAction()  {}

// TopDropTrigger is what opens a menu in the top bar.
type TopDropTrigger interface {
	part
	topDropTrigger()
}

func (Button) topDropTrigger()   {}
func (UserMenu) topDropTrigger() {}

// TopDrop is a button of the top bar with the menu it opens under it.
type TopDrop struct {
	Trigger TopDropTrigger
	Menu    Menu
	End     bool // the menu is aligned to the button's end: the last button of the bar
	Open    bool // drawn open; closed, the menu waits hidden for its button
}

func (d TopDrop) html(mix) string {
	panel := "tc-drop__panel"
	if d.End {
		panel = "tc-drop__panel--right"
	}
	if !d.Open {
		if sum, ok := summaryOf(d.Trigger.html(nil)); ok {
			return `<details class="tc-drop__anchor">` + sum + `<span class="` + panel + `">` + d.Menu.html(nil) + `</span></details>`
		}
	}
	return `<span class="tc-drop__anchor">` + d.Trigger.html(nil) + `<span class="` + panel + `"` + flag("hidden", !d.Open) + `>` + d.Menu.html(nil) + `</span></span>`
}

// Top — tc-top
//
// Роль. Верхняя полоса приложения: знак, поиск по всему, главные действия и учётная запись.
// Вход. Top{Lead, Brand, Search, Suggest, Actions, Tight, Zone}.
// Состояния. plain | phone | suggest | drop.
// Инварианты. Знак — один и первый, после кнопки рейла на телефоне; действия — в конце полосы; полоса печатает фрагмент своего кадра.
// Поведение. нет: поведение у поиска, кнопок и меню учётной записи.
// Эталон. кадр `shell-1`, фрагмент `.tc-top`.
// Содержимое. Кнопка рейла (телефон), знак или название раздела, поиск с подсказками под ним, кнопки, кнопки с меню, меню учётной записи.
// Чего нет. Навигации по разделам: она в рейле.
type Top struct {
	Lead    *Button // on a phone: the button that opens the rail
	Brand   Brand
	Search  *Search
	Suggest *Menu // what the search suggests, under it
	Actions []TopAction
	Tight   bool // 2 px between the actions: a phone's icon buttons
	// Zone is the name the consumer gives this area of the page: the live
	// layer puts an answer for that name here (Carrel's task 84). The library
	// has no names of its own.
	Zone string
}

// Render prints the bar.
func (t Top) Render() string { return t.html(nil) }

func (t Top) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-top") + attr("data-tc-zone", t.Zone) + `>`)
	if t.Lead != nil {
		s.WriteString(t.Lead.html(nil))
	}
	s.WriteString(t.Brand.html(nil))
	switch {
	case t.Search != nil && t.Suggest != nil:
		s.WriteString(`<span class="tc-line tc-top__search">` + t.Search.html(nil) + `<span class="tc-top__suggest">` + t.Suggest.html(nil) + `</span></span>`)
	case t.Search != nil:
		s.WriteString(t.Search.html(nil))
	}
	s.WriteString(`<div` + class(nil, "tc-top__actions", pick(t.Tight, "tc-top__actions--gap2")) + `>`)
	for _, a := range t.Actions {
		s.WriteString(a.html(nil))
	}
	s.WriteString(`</div></div>`)
	return s.String()
}
