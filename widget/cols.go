package widget

// Cols — tc-cols
//
// Роль. Запасной выход: экран собирает форму, которой нет в каталоге.
// Вход. Cols{ Shape, Items }.
// Состояния. halves | thirds | sidebar.
// Инварианты. Число прямых детей равно числу дорожек формы.
// Поведение. нет.
// Геометрия. Три формы; зазор 14; на 620 одна колонка по ближайшему tcbox над сеткой; пустой Items — пустая строка.
// Содержимое. Любые виджеты, по одному на дорожку.
// Чего нет. Не заменяет Shell, Doc и Fields. Нет Gap, ширин и выравнивания снаружи. Не плитка с переносом.
type Cols struct {
	Shape ColsShape
	Items []Node
}

func (c Cols) Render() string {
	if len(c.Items) == 0 {
		return ""
	}
	var inner string
	for _, item := range c.Items {
		inner += `<div class="tc-cols__col">` + item.Render() + `</div>`
	}
	return `<div` + classAttr("tc-cols", "tc-cols--"+string(c.Shape)) + `>` + inner + `</div>`
}
