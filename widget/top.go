package widget

import "strings"

// TopAction is what stands at the end of the top bar.
type TopAction interface {
	part
	topAction()
}

func (Button) topAction()   {}
func (UserMenu) topAction() {}

// Top — tc-top
//
// Роль. Верхняя полоса приложения: знак, поиск по всему, главные действия и учётная запись.
// Вход. Top{Lead, Brand, Search, Actions, Tight}.
// Состояния. plain | phone.
// Инварианты. Знак — один и первый, после кнопки рейла на телефоне; действия — в конце полосы; полоса печатает фрагмент своего кадра.
// Поведение. нет: поведение у поиска, кнопок и меню учётной записи.
// Эталон. кадр `shell-1`, фрагмент `.tc-top`.
// Содержимое. Кнопка рейла (телефон), знак или название раздела, поиск, кнопки, меню учётной записи.
// Чего нет. Навигации по разделам: она в рейле.
type Top struct {
	Lead    *Button // on a phone: the button that opens the rail
	Brand   Brand
	Search  *Search
	Actions []TopAction
	Tight   bool // 2 px between the actions: a phone's icon buttons
}

// Render prints the bar.
func (t Top) Render() string { return t.html(nil) }

func (t Top) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-top") + `>`)
	if t.Lead != nil {
		s.WriteString(t.Lead.html(nil))
	}
	s.WriteString(t.Brand.html(nil))
	if t.Search != nil {
		s.WriteString(t.Search.html(nil))
	}
	s.WriteString(`<div` + class(nil, "tc-top__actions", pick(t.Tight, "tc-top__actions--gap2")) + `>`)
	for _, a := range t.Actions {
		s.WriteString(a.html(nil))
	}
	s.WriteString(`</div></div>`)
	return s.String()
}
