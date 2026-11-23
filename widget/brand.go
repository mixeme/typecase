package widget

import "strings"

// LogoSize is the drawn size of the logo where it is not the brand's.
type LogoSize int

const (
	LogoBrand LogoSize = iota // the top bar's
	Logo17                    // 17 px high: a window's title bar
	Logo42                    // 42 px high: the first screen of an installer
	Logo60                    // 60 px high: the sign-in page
	Logo48                    // 48 px high: a page of the sign-in flow
)

func (s LogoSize) class() string {
	return [...]string{"", "tc-brand__logo--h17-w44", "tc-brand__logo--h42-w104", "tc-brand__logo--h60-w150", "tc-brand__logo--h48-w118"}[s]
}

// Logo is the consumer's logo: two symbols of the sprite, #logo-word in ink
// and #logo-shade in the accent, in one picture with a name.
type Logo struct {
	Label   string // the name the picture stands for
	ViewBox string // the box of the two symbols
	Size    LogoSize
}

func (l Logo) html(mix) string {
	return `<svg` + class(nil, "tc-brand__logo", l.Size.class()) + attr("viewBox", l.ViewBox) + ` role="img"` + attr("aria-label", l.Label) +
		`><use href="#logo-word" class="tc-logo__word"/><use href="#logo-shade" class="tc-logo__shade"/></svg>`
}

// Brand — tc-brand
//
// Роль. Знак приложения в начале верхней полосы; на телефоне на его месте — название раздела.
// Вход. Brand{Logo, Title}.
// Состояния. logo | title.
// Инварианты. Знак — одна картинка с именем, без повторения имени словом рядом; бренд печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `shell-1`, фрагмент `.tc-brand`.
// Содержимое. Логотип или название раздела.
// Чего нет. Ссылки: на главную ведёт навигация рейла.
type Brand struct {
	Logo  *Logo
	Title string // on a phone: the section's name instead of the logo
}

// Render prints the brand.
func (b Brand) Render() string { return b.html(nil) }

func (b Brand) html(m mix) string {
	if b.Logo != nil {
		return `<div` + class(m, "tc-brand") + `>` + b.Logo.html(nil) + `</div>`
	}
	return `<div` + class(m, "tc-brand", "tc-brand--t16") + `>` + esc(b.Title) + `</div>`
}

// SearchVariant is the drawn width of the search.
type SearchVariant int

const (
	SearchPlain SearchVariant = iota
	SearchFull                // the whole middle of the top bar: a search in progress
	SearchFill                // grows from 180 px: a top bar with an icon rail
)

func (v SearchVariant) class() string {
	return [...]string{"", "tc-search--full", "tc-search--fill180-wauto"}[v]
}

// Search — tc-search
//
// Роль. Поиск по всему в верхней полосе: значок и строка запроса.
// Вход. Search{Query, Focus, Clear, Variant, Action}.
// Состояния. plain | focus | clear | variant.
// Инварианты. Строка запроса — одна, с доступным именем; ширина — перечисление; поиск печатает фрагмент своего кадра.
// Поведение. Ввод ищет; Esc очищает запрос и закрывает подсказки.
// Эталон. кадр `shell-1`, фрагмент `.tc-search`.
// Содержимое. Значок поиска, крестик очистки, строка запроса.
// Чего нет. Своих подсказок: список недавнего — меню под поиском в верхней полосе.
type Search struct {
	Query   Query
	Focus   bool
	Clear   bool // a cross that empties the query
	Variant SearchVariant
	Action  string // the search is a form that sends its query there (GET)
}

// Render prints the search.
func (s Search) Render() string { return s.html(nil) }

func (s Search) html(m mix) string {
	tag := "div"
	form := ""
	if s.Action != "" {
		tag, form = "form", attr("action", s.Action)+` method="get"`
	}
	out := `<` + tag + class(m, "tc-search", is("focus", s.Focus), s.Variant.class()) + form + ` role="search">` +
		`<svg width="13" height="13" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.5" fill="none">` + Icon("search").use() + `</svg>`
	if s.Clear {
		out += `<svg width="12" height="12" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.8" fill="none">` + Icon("close").use() + `</svg>`
	}
	return out + s.Query.html(nil) + `</` + tag + `>`
}

// UserMenu — tc-usermenu
//
// Роль. Кнопка учётной записи в конце верхней полосы: имя и стрелка меню.
// Вход. UserMenu{Name, Focus, NoArrow, Href}.
// Состояния. plain | focus | no-arrow.
// Инварианты. Имя — текст кнопки; кнопка печатает фрагмент своего кадра.
// Поведение. Нажатие открывает меню учётной записи; Esc закрывает его и возвращает фокус кнопке.
// Эталон. кадр `shell-1`, фрагмент `.tc-usermenu`.
// Содержимое. Имя, стрелка.
// Чего нет. Аватара: учётную запись называет имя.
type UserMenu struct {
	Name    string
	Focus   bool
	NoArrow bool   // nothing to open: a screen that stands alone
	Href    string // the account's page: where the button goes without a script
}

// Render prints the button.
func (u UserMenu) Render() string { return u.html(nil) }

func (u UserMenu) html(m mix) string {
	tag := "button"
	open := ` type="button"`
	if u.Href != "" {
		tag, open = "a", attr("href", u.Href)
	}
	s := `<` + tag + class(m, "tc-usermenu", is("focus", u.Focus)) + open + `>` + esc(u.Name)
	if !u.NoArrow {
		s += `<svg width="12" height="12" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.5" fill="none">` + Icon("chev").use() + `</svg>`
	}
	return s + `</` + tag + `>`
}

// Emblem — tc-emblem
//
// Роль. Знак службы, собранный из трёх частей спрайта: штрих и заливка чернилами, одна часть акцентом.
// Вход. Emblem{Symbol, Title, Alert}.
// Состояния. plain | titled | alert.
// Инварианты. Части — символы спрайта `#Symbol-stroke`, `-fill`, `-accent`, цвет — из токенов; знак без подписи скрыт от чтения, с подписью — картинка с именем; знак печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `components-36`, фрагмент `.tc-emblem`.
// Содержимое. нет.
// Чего нет. Своего цвета тревоги: тревогу несёт плашка вокруг, а не перекрашенный знак.
type Emblem struct {
	Symbol string // the parts' prefix in the sprite
	Title  string // what the emblem says; empty beside a label that says it
	Alert  bool
}

// Render prints the emblem.
func (e Emblem) Render() string { return e.html(nil) }

func (e Emblem) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<svg` + class(m, "tc-emblem", is("alert", e.Alert)) + ` viewBox="0 0 16 16"`)
	if e.Title == "" {
		s.WriteString(` aria-hidden="true">`)
	} else {
		s.WriteString(` role="img"><title>` + esc(e.Title) + `</title>`)
	}
	for _, p := range []string{"stroke", "fill", "accent"} {
		s.WriteString(`<use href="#` + esc(e.Symbol) + `-` + p + `" class="tc-emblem__` + p + `"/>`)
	}
	s.WriteString(`</svg>`)
	return s.String()
}

// Split — tc-split
//
// Роль. Кнопка с меню: главное действие и стрелка к остальным, сомкнутые в одну.
// Вход. Split{Main, More}.
// Состояния. plain | menu.
// Инварианты. Половин две, и обе — кнопки; сплит печатает фрагмент своего кадра.
// Поведение. Главная половина делает главное; стрелка открывает меню остального (More — SplitDrop): «New note» делает заметку, меню — контакт, событие, задачу… (владелец, задача Carrel 276). У первичного сплита между половинами светлая черта: видно, что их две.
// Эталон. кадр `components-9`, фрагмент `.tc-split`.
// Содержимое. Две кнопки; вторая — стрелка, одна или с меню.
// Чего нет. Третьей половины.
type Split struct {
	Main Button
	More SplitMore
}

// SplitMore is a split's second half: an arrow button, or the arrow with
// the menu of the rest (SplitDrop).
type SplitMore interface {
	part
	splitMore()
}

func (Button) splitMore()    {}
func (SplitDrop) splitMore() {}

// SplitDrop is a split's arrow with the menu it opens: the rest of what the
// main half makes (Carrel's task 276: «New note» makes a note, its arrow
// the contact, the event, the task…). Closed, it waits in a details, as a
// top bar's drop does; Open draws it open, as the frames do.
type SplitDrop struct {
	Arrow Button
	Menu  Menu
	Open  bool
}

func (d SplitDrop) html(m mix) string {
	return TopDrop{Trigger: d.Arrow, Menu: d.Menu, Open: d.Open}.html(m)
}

func (Split) topAction() {}
func (Split) actsPart()  {}

// Render prints the split button.
func (p Split) Render() string { return p.html(nil) }

func (p Split) html(m mix) string {
	more := ""
	if p.More != nil {
		more = p.More.html(nil)
	}
	return `<span` + class(m, "tc-split") + `>` + p.Main.html(nil) + more + `</span>`
}

// Zone — tc-zone
//
// Роль. Подпись зоны оболочки на схеме: как называется место и какой оно ширины.
// Вход. Zone{Title, Note}.
// Состояния. plain.
// Инварианты. Рубрика и моноширинная строка под ней; зона печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `shell-1`, фрагмент `.tc-zone`.
// Содержимое. Рубрика, моноширинная строка.
// Чего нет. Содержимого зоны: это подпись на схеме, а не место для виджетов.
type Zone struct {
	Title Rubric
	Note  Mono
}

// Render prints the zone.
func (z Zone) Render() string { return z.html(nil) }

func (z Zone) html(m mix) string {
	return `<div` + class(m, "tc-zone") + `>` + z.Title.html(nil) + z.Note.html(nil) + `</div>`
}

// Crumb is a step of a path that leads back.
type Crumb struct {
	Label string
	Href  string
}

// Crumbs — tc-crumbs
//
// Роль. Путь до открытой папки: шаги назад ссылками, где мы — жирным.
// Вход. Crumbs{Path, Current, Aside, Target}.
// Состояния. path | aside.
// Инварианты. Последний шаг — не ссылка; путь печатает фрагмент своего кадра.
// Поведение. Нажатие на шаг открывает ту папку.
// Эталон. кадр `components-15`, фрагмент `.tc-crumbs`.
// Содержимое. Ссылки, разделители, текущий шаг, приглушённая сторона.
// Чего нет. Сокращения длинного пути: путь переносится.
type Crumbs struct {
	Path    []Crumb
	Current string
	Aside   *Dim // what is in it: a count and a size
	// Target is the name of the zone the links and forms inside ask their
	// address for with the live layer, instead of loading the page (Carrel's
	// task 84); without the layer they are plain links and forms.
	Target string
}

// Render prints the path.
func (c Crumbs) Render() string { return c.html(nil) }

func (c Crumbs) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<nav` + class(m, "tc-crumbs") + attr("data-tc-target", c.Target) + `>`)
	for _, p := range c.Path {
		s.WriteString(`<a class="tc-crumbs__link"` + attr("href", p.Href) + `>` + esc(p.Label) + `</a>/`)
	}
	s.WriteString(`<b>` + esc(c.Current) + `</b>`)
	if c.Aside != nil {
		s.WriteString(" " + c.Aside.html(nil))
	}
	s.WriteString(`</nav>`)
	return s.String()
}

// ActsPart is what a group of actions holds.
type ActsPart interface {
	part
	actsPart()
}

func (Button) actsPart() {}
func (Close) actsPart()  {}
func (Field) actsPart()  {}
func (Filter) actsPart() {}
func (Seg) actsPart()    {}
func (Drop) actsPart()   {}

// Acts — tc-acts
//
// Роль. Действия в конце шапки: кнопки, переключатель вида, крестик панели.
// Вход. Acts{Items, Stack}.
// Состояния. row | stack.
// Инварианты. Порядок — от частого к редкому, крестик — последним; группа печатает фрагмент своего кадра.
// Поведение. нет: поведение у кнопок.
// Эталон. кадр `components-14`, фрагмент `.tc-acts`.
// Содержимое. Кнопки, крестик, бейдж, поле, фильтр, переключатель.
// Чего нет. Меню «⋯» сама по себе не заводит: оно — кнопка среди прочих.
type Acts struct {
	Items []ActsPart
	Stack bool // one under another: a switch over a field
}

// Render prints the group.
func (a Acts) Render() string { return a.html(nil) }

func (a Acts) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-acts", pick(a.Stack, "tc-acts--stack")) + `>`)
	for _, it := range a.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}
