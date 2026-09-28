package widget

import "strconv"

// MarkKind is what the mark stands for.
type MarkKind int

const (
	MarkOne      MarkKind = iota // one source, in its colour
	MarkMulti                    // several sources at once: no single colour
	MarkExternal                 // a source read from outside, in its colour
)

func (k MarkKind) class() string {
	return [...]string{"", "is-multi", "is-ext"}[k]
}

// MarkSize is the drawn height of the mark where it is not the row's.
type MarkSize int

const (
	MarkDefault MarkSize = iota
	MarkH11
	MarkH11Inline // 11 px, set in a line of text
	MarkH12
	MarkH14
	MarkH18
	MarkH19
	MarkH20
	MarkH32
	MarkH34
	MarkOnHair // the empty mark on a hairline
)

func (s MarkSize) class() string {
	return [...]string{"", "tc-mark--h11", "tc-mark--h11 tc-mark--inline", "tc-mark--h12", "tc-mark--h14",
		"tc-mark--h18", "tc-mark--h19", "tc-mark--h20", "tc-mark--h32", "tc-mark--h34", "tc-mark--on-hair"}[s]
}

// Mark — tc-mark
//
// Роль. Цветная метка источника в строке, в рейле, в заголовке.
// Вход. Mark{Colour, Kind, Size, Focus}.
// Состояния. one | multi | external | sized | focus.
// Инварианты. Цвет приходит данными и становится свойством сценарием, а не стилем в разметке; род метки — перечисление; метка печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `components-11`, фрагмент `.tc-mark`.
// Содержимое. нет.
// Чего нет. Цвет не единственный носитель смысла: рядом с меткой всегда стоит название; без сценария метка без цвета.
type Mark struct {
	Colour string // "#4A6B52"; empty for a multi mark or one on a hairline
	Kind   MarkKind
	Size   MarkSize
	Focus  bool // the drawn keyboard focus of the mockup's states
}

// Render prints the mark.
func (k Mark) Render() string { return k.html(nil) }

func (k Mark) html(m mix) string {
	return `<span` + class(m, "tc-mark", k.Kind.class(), k.Size.class(), is("focus", k.Focus)) + attr("data-tc-colour", k.Colour) + `></span>`
}

// Tag — tc-tag
//
// Роль. Метка категории: тег записи, фильтр по тегу, приложенная вещь, связанная запись.
// Вход. Tag{Label, On, Icon, Mark, Removable, Place}.
// Состояния. plain | on | icon | mark | removable.
// Инварианты. Тег не правится на месте: он или выбран, или снимается крестиком; тег печатает фрагмент своего кадра.
// Поведение. Выбранный фильтр (On) снимается нажатием; тег с крестиком убирается крестиком.
// Эталон. кадр `s-contacts-1`, фрагмент `.tc-tag`.
// Содержимое. Значок или метка перед подписью, крестик после.
// Чего нет. Своего цвета: цвет несёт метка внутри.
type Tag struct {
	Label     string
	On        bool
	Icon      Icon
	Mark      *Mark
	Removable bool // a pill with a cross after the label
	Place     Place
}

// Render prints the tag.
func (t Tag) Render() string { return t.html(nil) }

func (t Tag) html(m mix) string {
	s := `<span` + class(m, "tc-tag", is("on", t.On), pick(t.Removable, "tc-tag--pill"), t.Place.class()) + `>`
	if t.Icon != "" {
		s += `<svg width="10" height="10" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.6" fill="none">` + t.Icon.use() + `</svg> `
	}
	if t.Mark != nil {
		s += t.Mark.html(nil)
	}
	s += esc(t.Label)
	if t.Removable {
		s += ` <svg width="9" height="9" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.2" fill="none">` + Icon("close").use() + `</svg>`
	}
	return s + `</span>`
}

// BadgeTone is what the badge says about its row.
type BadgeTone int

const (
	BadgeNeutral BadgeTone = iota
	BadgeLinked            // joined with other records
	BadgeAlert             // needs attention
	BadgeLocal             // reachable only here
)

func (t BadgeTone) class() string {
	return [...]string{"", "is-linked", "is-alert", "is-local"}[t]
}

// Badge — tc-badge
//
// Роль. Короткая пометка рядом с именем: «read-only», «linked ×2», «tentative».
// Вход. Badge{Label, Tone, Icon, Accent, Title, Place}.
// Состояния. neutral | linked | alert | local | accent | icon.
// Инварианты. В бейдже есть текст; бейдж не ссылка и не кнопка; бейдж печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `components-11`, фрагмент `.tc-badge`.
// Содержимое. Значок перед текстом.
// Чего нет. Действия: что с помеченным делать, говорит строка, не бейдж.
type Badge struct {
	Label  string
	Tone   BadgeTone
	Icon   Icon
	Accent bool   // accent edge and text: an unlocked state that runs out
	Title  string // what the badge means, as a hint
	Place  Place
}

// Render prints the badge.
func (b Badge) Render() string { return b.html(nil) }

func (b Badge) html(m mix) string {
	s := `<span` + class(m, "tc-badge", b.Tone.class(), pick(b.Accent, "tc-badge--edge-accent-accent"), b.Place.class()) + attr("title", b.Title) + `>`
	if b.Icon != "" {
		s += `<svg width="9" height="9" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2" fill="none" class="tc-ico--inline">` + b.Icon.use() + `</svg> `
	}
	return s + esc(b.Label) + `</span>`
}

// TickVariant is a drawn variant of the tick's box.
type TickVariant int

const (
	TickDefault    TickVariant = iota
	TickNudge                  // lowered to sit on a line of prose
	TickEdgeDanger             // red edge: a check that failed
	TickEdgeDusty              // dusty edge: a check not run
)

func (v TickVariant) class() string {
	return [...]string{"", "tc-tick--nudge", "tc-tick--edge-danger", "tc-tick--edge-dusty"}[v]
}

// Tick — tc-tick
//
// Роль. Галка отметки: задача сделана, пункт списка выполнен.
// Вход. Tick{Done, Variant, Glyph, Label, Send}.
// Состояния. open | done | variant | glyph.
// Инварианты. Отметка — кнопка с доступным именем, кроме знака внутри другой кнопки; сделано или нет — одно поле; отметка печатает фрагмент своего кадра.
// Поведение. Нажатие переключает отметку; знак (Glyph) не нажимается: нажимается кнопка, в которой он стоит.
// Эталон. кадр `s-tasks-1`, фрагмент `.tc-tick`.
// Содержимое. нет.
// Чего нет. Подписи: подпись — у строки или кнопки рядом.
type Tick struct {
	Done    bool
	Variant TickVariant
	Glyph   bool   // only the drawn box, inside another button
	Label   string // the accessible name: what is ticked
	Send    Send   // what pressing sends: the task done or not, without a script
}

// Render prints the tick.
func (t Tick) Render() string { return t.html(nil) }

func (t Tick) html(m mix) string {
	c := class(m, "tc-tick", is("done", t.Done), t.Variant.class())
	if t.Glyph {
		return `<span` + c + `></span>`
	}
	typ := "button"
	if t.Send.Form != "" {
		typ = "submit"
	}
	return `<button` + c + ` type="` + typ + `"` + attr("form", t.Send.Form) + attr("name", t.Send.Name) + attr("value", t.Send.Value) +
		attr("aria-label", t.Label) + attr("aria-pressed", strconv.FormatBool(t.Done)) + `></button>`
}

// Kbd — tc-kbd
//
// Роль. Клавиша или сочетание в подсказке и в таблице клавиш.
// Вход. Kbd{Keys, Place, Secondary}.
// Состояния. нет.
// Инварианты. Не нажимается: это надпись, а не кнопка; клавиша печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `components-29`, фрагмент `.tc-kbd`.
// Содержимое. нет.
// Чего нет. Привязки к действию: сочетание называет клавиши, а работает поведение библиотеки.
type Kbd struct {
	Keys      string // "Ctrl+S"
	Place     Place
	Secondary bool
}

// Render prints the key.
func (k Kbd) Render() string { return k.html(nil) }

func (k Kbd) html(m mix) string {
	return `<kbd` + class(m, "tc-kbd", k.Place.class(), rank(k.Secondary)) + `>` + esc(k.Keys) + `</kbd>`
}

// Num — tc-num
//
// Роль. Число рядом с подписью: сколько записей в источнике, в разделе, в группе.
// Вход. Num{Value, Alert}.
// Состояния. normal | alert.
// Инварианты. Значение — целое, а не строка; число печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `components-12`, фрагмент `.tc-num`.
// Содержимое. нет.
// Чего нет. Единиц и слов: число стоит рядом с подписью, которая их называет.
type Num struct {
	Value int
	Alert bool // the count is of things that need attention
}

// Render prints the number.
func (n Num) Render() string { return n.html(nil) }

func (n Num) html(m mix) string {
	return `<span` + class(m, "tc-num", is("alert", n.Alert)) + `>` + strconv.Itoa(n.Value) + `</span>`
}

// AvatarSize is the drawn size of an avatar where it is not the row's.
type AvatarSize int

const (
	AvatarDefault AvatarSize = iota
	Avatar34                 // 34 px square, 12 px letters: a phone row
	Avatar44                 // 44 px square, 15 px letters, radius 4: a card head
)

func (s AvatarSize) class() string {
	return [...]string{"", "tc-avatar--sq34-t12", "tc-avatar--sq44-t15-r4"}[s]
}

// Avatar — tc-avatar
//
// Роль. Лицо записи: инициалы или фотография.
// Вход. Avatar{Initials, Photo, Size, Src}.
// Состояния. initials | photo | sized | image.
// Инварианты. Аватар не пуст: инициалы или фотография; квадрат, размер — перечисление; аватар печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `components-12`, фрагмент `.tc-avatar`.
// Содержимое. нет.
// Чего нет. Подписи для чтеца: имя стоит рядом; без картинки (Src) — место фотографии, как в кадрах.
type Avatar struct {
	Initials string
	Photo    bool
	Size     AvatarSize
	Src      string // the photo's address: the picture fills the place
}

// Render prints the avatar.
func (a Avatar) Render() string { return a.html(nil) }

func (a Avatar) html(m mix) string {
	if a.Src != "" {
		return `<span` + class(m, "tc-avatar", "is-photo", a.Size.class()) + `><img class="tc-avatar__img"` + attr("src", a.Src) + ` alt=""></span>`
	}
	text := a.Initials
	if a.Photo {
		text = ""
	}
	return `<span` + class(m, "tc-avatar", is("photo", a.Photo), a.Size.class()) + `>` + esc(text) + `</span>`
}

// pick returns s when on.
func pick(on bool, s string) string {
	if on {
		return s
	}
	return ""
}
