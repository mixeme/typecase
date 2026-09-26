package widget

import (
	"strconv"
	"strings"
)

// MsgPart is what a message's text holds.
type MsgPart interface {
	part
	msgPart()
}

func (Text) msgPart()   {}
func (Strong) msgPart() {}
func (Bold) msgPart()   {}
func (Mono) msgPart()   {}
func (Span) msgPart()   {}
func (Link) msgPart()   {}
func (Break) msgPart()  {}
func (Button) msgPart() {}

// Bold is bold running text that holds more than words: a name with an
// address in it.
type Bold struct {
	Parts []Inline
}

func (b Bold) html(mix) string { return `<b>` + inlines(b.Parts) + `</b>` }

// Emphasis is a word in bold in the accent: what to press in a sentence.
type Emphasis string

func (e Emphasis) html(mix) string { return `<b class="tc-dim--accent">` + esc(string(e)) + `</b>` }

// joined prints pieces of running text one after another.
func joined[T part](parts []T) string {
	var s strings.Builder
	for _, p := range parts {
		s.WriteString(p.html(nil))
	}
	return s.String()
}

// MsgTone is what a message says about what it tells.
type MsgTone int

const (
	MsgPlain  MsgTone = iota
	MsgAlert          // something is wrong or cannot be undone
	MsgAccent         // something went well
)

func (t MsgTone) class() string { return [...]string{"", "is-alert", "is-accent"}[t] }

// MsgSpace is the space a message keeps around it.
type MsgSpace int

const (
	MsgSpacePage     MsgSpace = iota // the page's gutters
	MsgFlush                         // none: inside a sheet or a card
	MsgFlushTop2                     // none, 2 px above
	MsgFlushTop4                     // none, 4 px above
	MsgFlushTop11                    // none, 11 px above
	MsgFlushBottom12                 // none, 12 px below
	MsgFlushSides                    // no gutters at the sides: in a dialog
	MsgFlushTop9                     // none, 9 px above: under a card's rubric
	MsgFlushEnd                      // no margin at the end: beside a meter
	MsgInset20                       // 20 px at the sides: in a side panel
	MsgSheet                         // 0 14 12: in a sheet's picker
)

func (s MsgSpace) class() string {
	return [...]string{"", "tc-msg--flush", "tc-msg--flush-top2", "tc-msg--flush-top4", "tc-msg--flush-top11", "tc-msg--flush-bot12", "tc-msg--mr0-ml0",
		"tc-msg--flush-top9", "tc-msg--flush-r0", "tc-msg--mx20", "tc-msg--m0-14-12"}[s]
}

// Msg — tc-msg
//
// Роль. Сообщение: что произошло, чего нельзя отменить, что получилось.
// Вход. Msg{Parts, Tone, Icon, End, Space, Place}.
// Состояния. plain | alert | accent | icon | flush | rich.
// Инварианты. Тон — перечисление, и тревожный тон не единственный носитель смысла: сказано словами; сообщение печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `components-22`, фрагмент `.tc-msg`.
// Содержимое. Значок в строку с текстом, текст, жирное, моноширинное, приглушённое, бейдж.
// Чего нет. Кнопки закрытия: сообщение живёт, пока верно то, что оно говорит.
type Msg struct {
	Parts []MsgPart
	Tone  MsgTone
	Icon  Icon  // an icon in a row with the text
	End   *Push // buttons at the end of a message in a row: install, later
	Space MsgSpace
	Place Place
}

// Render prints the message.
func (g Msg) Render() string { return g.html(nil) }

func (g Msg) html(m mix) string {
	s := `<div` + class(m, "tc-msg", g.Tone.class(), is("row", g.Icon != ""), g.Space.class(), g.Place.class()) + `>`
	if g.Icon != "" {
		s += `<svg width="15" height="15" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.4" fill="none">` +
			g.Icon.use() + `</svg><span class="tc-msg__aside">` + joined(g.Parts) + `</span>`
		if g.End != nil {
			s += g.End.html(nil)
		}
		return s + `</div>`
	}
	return s + joined(g.Parts) + `</div>`
}

// DiagVariant is the drawn form of a diagnostics block.
type DiagVariant int

const (
	DiagPlain  DiagVariant = iota
	DiagInk                // ink, 9 px above: a line to paste
	DiagDanger             // what failed
)

func (v DiagVariant) class() string {
	return [...]string{"", "tc-diag--top9-ink", "tc-diag--danger"}[v]
}

// Diag — tc-diag
//
// Роль. Диагностика: что ответил сервер, строка для вставки, идентификаторы — моноширинным, как есть.
// Вход. Diag{Text, Variant}.
// Состояния. plain | ink | danger.
// Инварианты. Текст — как есть, с переносами строк; диагностика печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `components-30`, фрагмент `.tc-diag`.
// Содержимое. Только текст.
// Чего нет. Разметки внутри: ответ сервера показывается как пришёл.
type Diag struct {
	Text    string
	Variant DiagVariant
}

// Render prints the block.
func (d Diag) Render() string { return d.html(nil) }

func (d Diag) html(m mix) string {
	return `<div` + class(m, "tc-diag", d.Variant.class()) + `>` + esc(d.Text) + `</div>`
}

// DangerSpace is the space a danger zone keeps around it.
type DangerSpace int

const (
	DangerSpacePage DangerSpace = iota
	DangerAtEnd                 // pushed to the end of a column, flush, 11 px above
	DangerTight                 // no space above, 11/14 px inside: drawn alone
)

func (s DangerSpace) class() string {
	return [...]string{"", "tc-danger--topauto-flush-padt11", "tc-danger--top0-pad11-14"}[s]
}

// Danger — tc-danger
//
// Роль. Опасная зона: кнопка, которая удаляет насовсем, и что это значит.
// Вход. Danger{Button, Note, Space}.
// Состояния. plain | aside | end.
// Инварианты. Кнопка — одна и красная; опасная зона — последней в своей колонке; зона печатает фрагмент своего кадра.
// Поведение. Нажатие спрашивает подтверждения листом.
// Эталон. кадр `s-contacts-6`, фрагмент `.tc-danger`.
// Содержимое. Кнопка, пояснение.
// Чего нет. Второй кнопки.
type Danger struct {
	Button Button
	Note   string
	Space  DangerSpace
}

// Render prints the zone.
func (d Danger) Render() string { return d.html(nil) }

func (d Danger) html(m mix) string {
	return `<div` + class(m, "tc-danger", d.Space.class()) + `>` + d.Button.html(nil) + esc(d.Note) + `</div>`
}

// MeterWidth is the drawn width of a meter.
type MeterWidth int

const (
	MeterAuto MeterWidth = iota
	Meter60              // 60 px: a phone's bar
	MeterFill            // grows up to 260 px
)

func (w MeterWidth) class() string {
	return [...]string{"", "tc-meter--w60", "tc-meter--max260-fill"}[w]
}

// Meter — tc-meter
//
// Роль. Шкала: сколько сделано — источников опрошено, файла загружено.
// Вход. Meter{Percent, Track, Alert, Width}.
// Состояния. plain | track | alert | width.
// Инварианты. Доля — число от 0 до 100, и её ставит сценарий, а не стиль в разметке; шкала печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `components-26`, фрагмент `.tc-meter`.
// Содержимое. Полоса.
// Чего нет. Подписи: что и сколько, говорит строка рядом.
type Meter struct {
	Percent int
	Track   int // the track's own width in per cent, where the frame draws it
	Alert   bool
	Width   MeterWidth
}

// Render prints the meter.
func (t Meter) Render() string { return t.html(nil) }

func (t Meter) html(m mix) string {
	track := ""
	if t.Track > 0 {
		track = strconv.Itoa(t.Track) + "%"
	}
	return `<span` + class(m, "tc-meter", t.Width.class()) + attr("data-tc-width", track) + `><i` + class(nil, "tc-meter__bar", pick(t.Alert, "tc-meter__bar--alert")) +
		attr("data-tc-width", strconv.Itoa(t.Percent)+"%") + ` role="progressbar"` + attr("aria-valuenow", strconv.Itoa(t.Percent)) + `></i></span>`
}

// RunbarPart is what follows a run's meter.
type RunbarPart interface {
	part
	runbarPart()
}

func (Text) runbarPart() {}
func (Span) runbarPart() {}
func (Mono) runbarPart() {}
func (Push) runbarPart() {}
func (Line) runbarPart() {}

// Runbar — tc-runbar
//
// Роль. Ход опроса над списком: шкала, что опрошено, что не ответило, когда, остановить.
// Вход. Runbar{Meter, Parts, Failed}.
// Состояния. plain | failed | stop.
// Инварианты. Шкала — первой; неответившее — словами, красным; строка печатает фрагмент своего кадра.
// Поведение. «Stop» прерывает опрос.
// Эталон. кадр `components-16`, фрагмент `.tc-runbar`.
// Содержимое. Шкала, текст, красные куски, моноширинное время, прижатая кнопка.
// Чего нет. Своей кнопки закрытия: строка уходит, когда опрос кончился.
type Runbar struct {
	Meter  Meter
	Parts  []RunbarPart
	Failed bool // the run failed as a whole
}

// Render prints the line.
func (r Runbar) Render() string { return r.html(nil) }

func (r Runbar) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-runbar", pick(r.Failed, "tc-runbar--danger")) + `>` + r.Meter.html(nil))
	for _, p := range r.Parts {
		s.WriteString(p.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// QueueItem is a file in the upload queue: its name, how far, a meter.
type QueueItem struct {
	Name   string
	Meta   string
	Meter  *Meter
	Failed bool     // the upload stopped: the name and the reason in red
	Acts   []Button // what to do with a stopped upload
}

// Queue — tc-queue
//
// Роль. Очередь загрузки: что грузится, сколько, что не вышло и что с ним сделать.
// Вход. Queue{Title, Total, Toggle, Items, Min, Variant}.
// Состояния. min | flush.
// Инварианты. Заголовок — сколько из скольких; ошибка — словами у файла, а не общей плашкой; очередь печатает фрагмент своего кадра.
// Поведение. «Hide» сворачивает очередь в строку, «Show» разворачивает.
// Эталон. кадр `components-26`, фрагмент `.tc-queue`.
// Содержимое. Заголовок, общий объём, кнопка сворачивания, строки файлов со шкалой, кнопки у ошибки.
// Чего нет. Отмены всех загрузок одной кнопкой.
type Queue struct {
	Title   string
	Total   Mono
	Toggle  Button
	Items   []QueueItem
	Min     bool // folded to its head
	Variant QueueVariant
}

// QueueVariant is where a queue stands.
type QueueVariant int

const (
	QueueFloat     QueueVariant = iota // over the page's corner
	QueueFlushFull                     // in its place, the width of it
	QueueFlushAuto                     // in its place, its own width
)

func (v QueueVariant) class() string {
	return [...]string{"", "tc-queue--flush-full", "tc-queue--flush-wauto"}[v]
}

// Render prints the queue.
func (q Queue) Render() string { return q.html(nil) }

func (q Queue) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-queue", is("min", q.Min), q.Variant.class()) + `><div class="tc-queue__head"><b>` + esc(q.Title) + `</b>` +
		q.Total.html(nil) + q.Toggle.html(nil) + `</div>`)
	if len(q.Items) > 0 {
		s.WriteString(`<ul class="tc-queue__list">`)
		for _, it := range q.Items {
			s.WriteString(`<li><div class="tc-queue__row">`)
			if it.Failed {
				s.WriteString(`<span class="tc-dim--danger">` + esc(it.Name) + `</span><span class="tc-dim--danger">` + esc(it.Meta) + `</span>`)
			} else {
				s.WriteString(`<span class="tc-queue__name">` + esc(it.Name) + `</span><span class="tc-queue__meta">` + esc(it.Meta) + `</span>`)
			}
			s.WriteString(`</div>`)
			if it.Meter != nil {
				s.WriteString(it.Meter.html(nil))
			}
			if len(it.Acts) > 0 {
				s.WriteString(`<div class="tc-line tc-line--gap-6 tc-queue__acts">`)
				for _, b := range it.Acts {
					s.WriteString(b.html(nil))
				}
				s.WriteString(`</div>`)
			}
			s.WriteString(`</li>`)
		}
		s.WriteString(`</ul>`)
	}
	s.WriteString(`</div>`)
	return s.String()
}

// Stat — tc-stat
//
// Роль. Счёт итога: сколько готово, сколько с ошибками.
// Вход. Stat{Value, Label, Error}.
// Состояния. plain | error.
// Инварианты. Значение — строка, как её показать, решает экран; счёт печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `components-30`, фрагмент `.tc-stat`.
// Содержимое. Жирное число и подпись.
// Чего нет. Единиц отдельно от подписи.
type Stat struct {
	Value string
	Label string
	Error bool
}

// Render prints the count.
func (t Stat) Render() string { return t.html(nil) }

func (t Stat) html(m mix) string {
	return `<span` + class(m, "tc-stat", is("error", t.Error)) + `><b>` + esc(t.Value) + `</b> ` + esc(t.Label) + `</span>`
}

// Stats — tc-stats
//
// Роль. Итоги в одну строку: готово, с ошибками, совпадения.
// Вход. Stats{Items}.
// Состояния. plain.
// Инварианты. Итоги — одна строка, без переноса; итоги печатают фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `components-30`, фрагмент `.tc-stats`.
// Содержимое. Счёты.
// Чего нет. Графиков.
type Stats struct {
	Items []Stat
}

// Render prints the counts.
func (t Stats) Render() string { return t.html(nil) }

func (t Stats) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-stats") + `>`)
	for _, it := range t.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// Steps — tc-steps
//
// Роль. Шаги мастера: где мы из трёх.
// Вход. Steps{Items, On}.
// Состояния. first | middle | last.
// Инварианты. Текущий шаг — один, его номер; между шагами — черта; шаги печатают фрагмент своего кадра.
// Поведение. нет: шаги не нажимаются, переход — кнопками мастера.
// Эталон. кадр `components-30`, фрагмент `.tc-steps`.
// Содержимое. Подписи шагов, черты.
// Чего нет. Перехода нажатием на шаг.
type Steps struct {
	Items []string
	On    int // the current step, counted from 1
}

// Render prints the steps.
func (t Steps) Render() string { return t.html(nil) }

func (t Steps) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-steps") + `>`)
	for i, it := range t.Items {
		if i > 0 {
			s.WriteString(`<i class="tc-steps__rule"></i>`)
		}
		s.WriteString(`<span` + class(nil, is("on", i+1 == t.On), "tc-steps__item") + attr("aria-current", pick(i+1 == t.On, "step")) + `>` + esc(it) + `</span>`)
	}
	s.WriteString(`</div>`)
	return s.String()
}

// EmptyPart is what an empty state holds after its words.
type EmptyPart interface {
	part
	emptyPart()
}

func (Button) emptyPart()   {}
func (Span) emptyPart()     {}
func (Mono) emptyPart()     {}
func (Strong) emptyPart()   {}
func (Text) emptyPart()     {}
func (Emphasis) emptyPart() {}
func (Dim) emptyPart()      {}
func (Aside) emptyPart()    {}

// Aside is running text in a face of its own that holds more than words: a
// sentence with a monospaced name in it.
type Aside struct {
	Face  Face
	Parts []Inline
}

func (a Aside) html(mix) string {
	return `<span class="` + a.Face.class() + `">` + inlines(a.Parts) + `</span>`
}

// EmptyVariant is the drawn form of an empty state.
type EmptyVariant int

const (
	EmptyPlain    EmptyVariant = iota
	EmptySmall                 // 12 px, a list's padding: a day with nothing
	EmptySmallTop              // 12 px, 12/20 padding
	EmptyDusty                 // dusty: a list filtered to nothing
	EmptyDashed                // a dashed frame: a place to drop a thing
	EmptyDrop                  // a drop zone over the page
	EmptyDropSide              // a drop zone in a side panel
)

func (v EmptyVariant) class() string {
	return [...]string{"", "tc-empty--t12-pad14-20", "tc-empty--t12-pad12-20", "tc-empty--dusty tc-empty--t12-pad12-20", "tc-empty--dashed",
		"tc-empty--drop", "tc-empty--drop-side"}[v]
}

// Empty — tc-empty
//
// Роль. Пусто: чего здесь нет и что сделать; место, куда бросить файл.
// Вход. Empty{Icon, Lead, Parts, Variant}.
// Состояния. plain | small | dusty | action | drop | dashed.
// Инварианты. Пустое состояние говорит словами, что пусто; вид — перечисление; оно печатает фрагмент своего кадра.
// Поведение. Бросок файла на зону загружает его.
// Эталон. кадр `components-22`, фрагмент `.tc-empty`.
// Содержимое. Значок, крупная первая строка, слова, жирное, моноширинное, клавиша, кнопка.
// Чего нет. Картинок-иллюстраций.
type Empty struct {
	Icon    Icon
	Lead    string // a first line in bold
	Parts   []EmptyPart
	Variant EmptyVariant
}

// Render prints the empty state.
func (e Empty) Render() string { return e.html(nil) }

func (e Empty) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-empty", e.Variant.class()) + `>`)
	if e.Icon != "" {
		s.WriteString(`<svg width="20" height="20" viewBox="0 0 24 24" stroke-width="1.4" fill="none" stroke-linecap="round" stroke-linejoin="round" class="tc-ico--accent">` + e.Icon.use() + `</svg>`)
	}
	if e.Lead != "" {
		lead := "tc-empty__lead"
		if e.Variant == EmptyDropSide {
			lead = "tc-empty__lead-sm"
		}
		s.WriteString(`<b class="` + lead + `">` + esc(e.Lead) + `</b>`)
	}
	for _, p := range e.Parts {
		s.WriteString(p.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// ThumbForm is the drawn form of a thumbnail.
type ThumbForm int

const (
	ThumbTile    ThumbForm = iota // in a tile of a grid
	ThumbOnCard                   // on card colour: a file that is not a picture
	ThumbPreview                  // a preview in a side panel
	ThumbWide                     // a wide preview on a public page
	ThumbSquare                   // a square: a file that cannot be shown
)

// Thumb — tc-thumb
//
// Роль. Миниатюра файла: место картинки со значком рода.
// Вход. Thumb{Icon, Form}.
// Состояния. tile | card | preview | wide | square.
// Инварианты. Вид — перечисление; миниатюра печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `s-files-5`, фрагмент `.tc-thumb`.
// Содержимое. Значок.
// Чего нет. Самой картинки: макет рисует её место.
type Thumb struct {
	Icon Icon
	Form ThumbForm
}

// Render prints the thumbnail.
func (t Thumb) Render() string { return t.html(nil) }

func (t Thumb) html(m mix) string {
	round := `viewBox="0 0 24 24" stroke-linecap="round" stroke-linejoin="round"`
	switch t.Form {
	case ThumbOnCard:
		return `<div` + class(m, "tc-thumb", "tc-thumb--on-card") + `><svg>` + t.Icon.use() + `</svg></div>`
	case ThumbPreview:
		return `<div` + class(m, "tc-thumb", "tc-thumb--preview") + `><svg ` + round + `>` + t.Icon.use() + `</svg></div>`
	case ThumbWide:
		return `<div` + class(m, "tc-thumb", "tc-thumb--wide") + `><svg ` + round + ` class="tc-ico--sq34">` + t.Icon.use() + `</svg></div>`
	case ThumbSquare:
		return `<span` + class(m, "tc-thumb", "tc-thumb--tile") + `><svg ` + round + ` class="tc-ico--sq30">` + t.Icon.use() + `</svg></span>`
	}
	return `<div` + class(m, "tc-thumb") + `><svg>` + t.Icon.use() + `</svg></div>`
}

// Tile — tc-tile
//
// Роль. Плитка файла в сетке: миниатюра, имя, размер.
// Вход. Tile{Thumb, Name, Meta, Selected}.
// Состояния. plain | selected.
// Инварианты. Имя — одной строкой с многоточием; плитка печатает фрагмент своего кадра.
// Поведение. Нажатие выбирает плитку; двойное — открывает файл.
// Эталон. кадр `components-25`, фрагмент `.tc-tile`.
// Содержимое. Миниатюра, имя, размер.
// Чего нет. Действий на плитке: они в баре выделения.
type Tile struct {
	Thumb    Thumb
	Name     string
	Meta     string
	Selected bool
}

// Render prints the tile.
func (t Tile) Render() string { return t.html(nil) }

func (t Tile) html(m mix) string {
	return `<div` + class(m, "tc-tile", is("sel", t.Selected)) + `>` + t.Thumb.html(nil) + `<span class="tc-tile__name">` + esc(t.Name) +
		`</span><span class="tc-tile__meta">` + esc(t.Meta) + `</span></div>`
}

// Grid — tc-grid
//
// Роль. Сетка плиток: файлы папки плитками.
// Вход. Grid{Tiles, Flush}.
// Состояния. plain | flush.
// Инварианты. В сетке только плитки; сетка печатает фрагмент своего кадра.
// Поведение. Стрелки ходят по плиткам.
// Эталон. кадр `s-files-5`, фрагмент `.tc-grid`.
// Содержимое. Плитки.
// Чего нет. Заголовков групп.
type Grid struct {
	Tiles []Tile
	Flush bool // no gutters, fills its place: drawn alone
}

// Render prints the grid.
func (g Grid) Render() string { return g.html(nil) }

func (g Grid) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-grid", pick(g.Flush, "tc-grid--flush-fill-pad8")) + `>`)
	for _, t := range g.Tiles {
		s.WriteString(t.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// PadPart is what a padded block holds.
type PadPart interface {
	part
	padPart()
}

func (Thumb) padPart() {}
func (Card) padPart()  {}

// Pad — tc-pad
//
// Роль. Отступ страницы вокруг того, у чего своих полей нет: карточки, превью.
// Вход. Pad{Items}.
// Состояния. plain.
// Инварианты. Поля — токен страницы; блок печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `s-files-1`, фрагмент `.tc-pad`.
// Содержимое. Карточки, превью.
// Чего нет. Своего вида.
type Pad struct {
	Items []PadPart
}

// Render prints the block.
func (p Pad) Render() string { return p.html(nil) }

func (p Pad) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-pad") + `>`)
	for _, it := range p.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// DiffItem is a property two versions disagree on: its rubric (the
// property and why they differ) and the checks of its versions.
type DiffItem struct {
	Title   Rubric
	Options []Check
}

// Diff — tc-diff
//
// Роль. Сравнение версий при конфликте: по каждому свойству — чья версия остаётся.
// Вход. Diff{Items}.
// Состояния. plain.
// Инварианты. У свойства выбран ровно один вариант; варианты — столбцом; сравнение печатает фрагмент своего кадра.
// Поведение. Нажатие выбирает вариант свойства.
// Эталон. кадр `s-conflict-1`, фрагмент `.tc-diff`.
// Содержимое. Свойства: имя моноширинным, почему расходятся, варианты.
// Чего нет. Слияния значений: выбирается одно.
type Diff struct {
	Items []DiffItem
}

// Render prints the comparison.
func (d Diff) Render() string { return d.html(nil) }

func (d Diff) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-diff") + `>`)
	for _, it := range d.Items {
		s.WriteString(`<div class="tc-diff__item">` + it.Title.html(nil) + `<div class="tc-diff__options">`)
		for _, o := range it.Options {
			s.WriteString(o.html(nil))
		}
		s.WriteString(`</div></div>`)
	}
	s.WriteString(`</div>`)
	return s.String()
}

// DiffAgree is what two versions agree on, above their differences.
type DiffAgree struct {
	Property Mono
	Text     string
}

func (a DiffAgree) html(mix) string {
	return `<div class="tc-diff__agree">` + a.Property.html(nil) + ` ` + esc(a.Text) + `</div>`
}

func (DiffAgree) dialogPart() {}
