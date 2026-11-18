package widget

import (
	"strconv"
	"strings"
	"time"
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
func (Mark) msgPart()   {}

// Bold is bold running text that holds more than words: a name with an
// address in it.
type Bold struct {
	Parts []Inline
}

func (b Bold) html(mix) string { return `<b>` + inlines(b.Parts) + `</b>` }

// Emphasis is a word in bold in the accent: what to press in a sentence.
type Emphasis string

func (e Emphasis) html(mix) string { return `<b class="tc-dim--accent">` + esc(string(e)) + `</b>` }

// FilePick is the words of an Empty that pick a file: «choose on the
// computer» of a place files are dropped into (Carrel's task 119). They are
// the label of their file input, which the browser opens without a script;
// there the input shows what was picked. The files wait for the button of
// the form the words stand in (File.Form and File.To stay empty).
type FilePick struct {
	Label string
	File  FileField
	// Row is the row each picked file becomes in the list File.List names
	// (the owner's decision of 09.10: what was picked is shown by name) —
	// «{name}» and «{size}» in its words, the name being the file's path
	// inside a picked folder; a button of the row with Removes takes the
	// file back.
	Row *Line
}

func (p FilePick) html(mix) string {
	row := ""
	if p.Row != nil {
		row = `<template data-tc-row>` + p.Row.html(nil) + `</template>`
	}
	return `<label class="tc-empty__pick">` + Emphasis(p.Label).html(nil) + p.File.input() + row + `</label>`
}

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
	MsgSpacePage        MsgSpace = iota // the page's gutters
	MsgFlush                            // none: inside a sheet or a card
	MsgFlushTop2                        // none, 2 px above
	MsgFlushAboveTight                  // none, the tight step above
	MsgFlushAboveWide                   // none, the wide step above
	MsgFlushBelowWide                   // none, the wide step below
	MsgFlushSides                       // no gutters at the sides: in a dialog
	MsgFlushAboveNormal                 // none, the normal step above: under a card's rubric
	MsgFlushEnd                         // no margin at the end: beside a meter
	MsgInset20                          // 20 px at the sides: in a side panel
	MsgSheet                            // 0 14 12: in a sheet's picker
	MsgAroundNormal                     // the normal step above and below, none at the sides: a refusal between a field and what follows
)

func (s MsgSpace) class() string {
	return [...]string{"", "tc-msg--flush", "tc-msg--flush-top2", "tc-msg--flush-above-tight", "tc-msg--flush-above-wide", "tc-msg--flush-below-wide", "tc-msg--mr0-ml0",
		"tc-msg--flush-above-normal", "tc-msg--flush-r0", "tc-msg--mx20", "tc-msg--m0-14-12", "tc-msg--around-normal"}[s]
}

// NoticesPart is what the zone of messages holds: a message, and a message's
// technical block under it (Carrel's task 235).
type NoticesPart interface {
	part
	noticesPart()
}

func (Msg) noticesPart()      {}
func (MainDiag) noticesPart() {}

// Notices — tc-notices
//
// Роль. Полоса сообщений над содержимым страницы, которую живой слой заменяет целиком.
// Вход. Notices{Items, Zone}.
// Состояния. plain | empty.
// Инварианты. Своих отступов нет — они у сообщений, как без полосы; пустая полоса скрыта; полоса печатает фрагмент своего кадра.
// Поведение. Пустой ответ живого слоя убирает сообщения и скрывает полосу; ответ с сообщениями её показывает.
// Эталон. кадр `s-dup-6`, фрагмент `.tc-notices`.
// Содержимое. Сообщения и технический блок под сообщением.
// Чего нет. Своего вида: вид у сообщений.
type Notices struct {
	Items []NoticesPart
	// Zone is the name the consumer gives this area of the page: the live
	// layer puts an answer for that name here (Carrel's task 226). The
	// library has no names of its own.
	Zone string
}

// Render prints the messages.
func (n Notices) Render() string { return n.html(nil) }

func (n Notices) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-notices") + attr("data-tc-zone", n.Zone) + flag("hidden", len(n.Items) == 0) + `>`)
	for _, it := range n.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// Msg — tc-msg
//
// Роль. Сообщение: что произошло, чего нельзя отменить, что получилось.
// Вход. Msg{Parts, Tone, Icon, End, Space, Place, NoScript, ScriptOnly, Column}.
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
	// NoScript shows the message only on a page without a script: it says
	// what needs one and is not there — a drop, a paste, the live queue
	// (Carrel's task 90). ScriptOnly shows it only with one: the words about
	// what a script does.
	NoScript   bool
	ScriptOnly bool
	// Column keeps the message in the reading column of a settings page,
	// 580 px at most, as its fields' hints (the owner, 27.10).
	Column bool
}

// Render prints the message.
func (g Msg) Render() string { return g.html(nil) }

func (g Msg) html(m mix) string {
	s := `<div` + class(m, "tc-msg", g.Tone.class(), is("row", g.Icon != ""), g.Space.class(), g.Place.class(), pick(g.Column, "tc-msg--max580"), pick(g.ScriptOnly, "tc-js"), pick(g.NoScript, "tc-nojs")) + `>`
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
	DiagInk                // ink, the normal step above: a line to paste
	DiagDanger             // what failed
)

func (v DiagVariant) class() string {
	return [...]string{"", "tc-diag--above-normal-ink", "tc-diag--danger"}[v]
}

// Diag — tc-diag
//
// Роль. Диагностика: что ответил сервер, строка для вставки, идентификаторы — моноширинным, как есть.
// Вход. Diag{Text, Variant, Place, ID}.
// Состояния. plain | ink | danger.
// Инварианты. Текст — как есть, с переносами строк; диагностика печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `components-30`, фрагмент `.tc-diag`.
// Содержимое. Только текст.
// Чего нет. Разметки внутри: ответ сервера показывается как пришёл.
type Diag struct {
	Text    string
	Variant DiagVariant
	Place   Place
	ID      string // what a button saves the text from (Button.SaveFrom)
}

// Render prints the block.
func (d Diag) Render() string { return d.html(nil) }

func (d Diag) html(m mix) string {
	return `<div` + class(m, "tc-diag", d.Variant.class(), d.Place.class()) + attr("id", d.ID) + `>` + esc(d.Text) + `</div>`
}

// DangerSpace is the space a danger zone keeps around it.
type DangerSpace int

const (
	DangerSpacePage DangerSpace = iota
	DangerAtEnd                 // pushed to the end of a column, flush, 11 px above
	DangerTight                 // no space above, 11/14 px inside: drawn alone
)

func (s DangerSpace) class() string {
	return [...]string{"", "tc-danger--topauto-flush-padt", "tc-danger--top0-pad"}[s]
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
	MeterFill            // grows up to 260 px
)

func (w MeterWidth) class() string {
	return [...]string{"", "tc-meter--max260-fill"}[w]
}

// Meter — tc-meter
//
// Роль. Шкала: сколько сделано — источников опрошено, файла загружено.
// Вход. Meter{Percent, Track, Width}.
// Состояния. plain | track | width.
// Инварианты. Доля — число от 0 до 100, и её ставит сценарий, а не стиль в разметке; шкала печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `components-26`, фрагмент `.tc-meter`.
// Содержимое. Полоса.
// Чего нет. Подписи: что и сколько, говорит строка рядом.
type Meter struct {
	Percent int
	Track   int // the track's own width in per cent, where the frame draws it
	Width   MeterWidth
}

// Render prints the meter.
func (t Meter) Render() string { return t.html(nil) }

func (t Meter) html(m mix) string {
	track := ""
	if t.Track > 0 {
		track = strconv.Itoa(t.Track) + "%"
	}
	return `<span` + class(m, "tc-meter", t.Width.class()) + attr("data-tc-width", track) + `><i` + class(nil, "tc-meter__bar") +
		attr("data-tc-width", strconv.Itoa(t.Percent)+"%") + ` role="progressbar"` + attr("aria-valuenow", strconv.Itoa(t.Percent)) + `></i></span>`
}

// RunResult is how a run that is over ended, shown by a dot in place of the meter.
type RunResult int

// RunResult values.
const (
	RunGoing    RunResult = iota // still running: the meter
	RunAnswered                  // over, every source answered: a green dot
	RunFailed                    // over, something did not answer: a red dot
	RunStopped                   // over, stopped before every source answered: a grey dot
)

// RunbarPart is what follows a run's meter.
type RunbarPart interface {
	part
	runbarPart()
}

func (Text) runbarPart() {}
func (Span) runbarPart() {}
func (Mono) runbarPart() {}
func (Say) runbarPart()  {}
func (Push) runbarPart() {}
func (Drop) runbarPart() {}

// Runbar — tc-runbar
//
// Роль. Ход опроса над списком: шкала, что опрошено, что не ответило, когда, остановить.
// Вход. Runbar{Meter, Result, Parts, Failed, Progress, Saving, Zone, Stream, Changed, Ended, Poll, Every, Closed}.
// Состояния. plain | failed | stop | answered | refused | stopped.
// Инварианты. Шкала — первой, пока опрос идёт; кончившийся опрос вместо шкалы несёт точку итога (Result): зелёную, если ответили все, красную, если кто-то нет, серую, если опрос остановили (владелец 17.11); слова — по шаблону Р-84: «Polled · N of M answered · K events · F failed», узко (Say) — «N of M · K events · F failed»; неответившее — кнопкой «F failed» с меню отказов (Drop, MenuRefusal), где каждый расписан полностью; строка печатает фрагмент своего кадра.
// Поведение. «Stop» прерывает опрос; строка-область с адресом потока или опроса следит за ходом сама; скрытая строка (пустая, ждущая хода или опустевшая по концу) не следит.
// Эталон. кадр `components-16`, фрагмент `.tc-runbar`.
// Содержимое. Шкала, текст, красные куски, моноширинное время, прижатая кнопка.
// Чего нет. Своей кнопки закрытия: строка уходит, когда опрос кончился.
type Runbar struct {
	Meter Meter
	// Result is how a run that is over ended: the line then begins with a
	// dot in its colour instead of the meter (the owner, 17.11) — green,
	// every source answered; red, something did not; grey, stopped. The
	// zero value is a run still going, with the meter.
	Result RunResult
	Parts  []RunbarPart
	Failed bool // the run failed as a whole
	// Progress makes the line the Save of the form it names, which sends
	// files: hidden until then; the live layer moves the meter and writes
	// «3.2 of 24 MB» into the line's Mono. Saving draws it as during Save.
	Progress string
	Saving   bool
	// Zone is the name the consumer gives this area of the page: the live
	// layer puts an answer for that name here (Carrel's task 84). The library
	// has no names of its own.
	Zone string
	// A line that is a zone follows its run with the live layer (Carrel's
	// task 85). Stream is the address of a stream of events: Changed and
	// Ended are the names of its events, each carrying a fragment — the line
	// as it is now and whatever other zones changed; on Ended the stream is
	// closed. Poll is the address asked instead, every Every, when there is
	// no stream or it broke, for as long as the line that comes back names
	// it. Without the layer the run shows after a reload. Every is sent in
	// milliseconds (Carrel's task 226: 1500 ms is a step its contract has).
	Stream  string
	Changed string
	Ended   string
	Poll    string
	Every   time.Duration
	// Closed prints the line empty and hidden: a zone that waits for a run,
	// where Retry or «Search again» put their answer (Carrel's task 226). A
	// hidden line is not followed: no stream, no poll.
	Closed bool
}

// Render prints the line.
func (r Runbar) Render() string { return r.html(nil) }

func (r Runbar) html(m mix) string {
	if r.Closed {
		return `<div` + class(m, "tc-runbar") + attr("data-tc-zone", r.Zone) + ` hidden></div>`
	}
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-runbar", pick(r.Failed, "tc-runbar--danger")) + attr("data-tc-zone", r.Zone) + attr("data-tc-stream", r.Stream) + attr("data-tc-changed", r.Changed) + attr("data-tc-ended", r.Ended) +
		attr("data-tc-poll", r.Poll) + attr("data-tc-every", pick(r.Every > 0, strconv.FormatInt(r.Every.Milliseconds(), 10))) + attr("data-tc-progress", r.Progress) + flag("hidden", r.Progress != "" && !r.Saving) + `>` + r.lead())
	for _, p := range r.Parts {
		s.WriteString(p.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// lead is what begins the line: the meter while the run goes, the dot of its result once it is over.
func (r Runbar) lead() string {
	switch r.Result {
	case RunAnswered:
		return Dot{Tone: DotAccent}.html(nil)
	case RunFailed:
		return Dot{Tone: DotDanger}.html(nil)
	case RunStopped:
		return Dot{Tone: DotDusty}.html(nil)
	}
	return r.Meter.html(nil)
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
// Вход. Queue{Title, Total, Toggle, Items, Min, Variant, Zone, Progress, Row, Fold, Taken}.
// Состояния. min | flush.
// Инварианты. Заголовок — сколько из скольких; ошибка — словами у файла, а не общей плашкой; очередь печатает фрагмент своего кадра.
// Поведение. «Hide» сворачивает очередь в строку, «Show» разворачивает: со сценарием — на месте (Fold), без него — ответом сервера.
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
	// Zone is the name the consumer gives this area of the page: the live
	// layer puts an answer for that name here (Carrel's task 84). The library
	// has no names of its own.
	Zone string
	// Progress is the ID of the form whose upload the queue shows, with the
	// library's live layer (Carrel's task 88): sent with files, the form
	// goes in place, the queue shows, every file gets a row printed from
	// Row, and the rows' meters move. The words in braces are the queue's:
	// «{name}», «{size}» and «{percent}» in a row; «{count}», «{done}»,
	// «{sent}», «{total}» and «{percent}» in the title and the total. A queue that has a
	// Progress and no Items waits hidden. Without the layer the form loads
	// its answer, and the queue is what the server printed.
	Progress string
	Row      *QueueItem
	// Fold, with the library's script, folds the queue in place when its
	// Toggle is pressed: the list hides, Toggle says Show, and the total says
	// Total — the words of the live layer's upload («{percent}%») once it
	// knows them; pressed again, it unfolds. Without a script Toggle is
	// what the server answers (Hide and Show as links).
	Fold *QueueFold
	// Taken is the row that asks about a file whose name is taken, before
	// it goes (FileField.Taken, Carrel's task 214): «{name}» is the file's,
	// its buttons' values say what to do — keep, replace or skip.
	Taken *QueueItem
}

// QueueFold is what a queue says folded and unfolded (Queue.Fold).
type QueueFold struct {
	Hide, Show, Total string
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
	s.WriteString(`<div` + class(m, "tc-queue", is("min", q.Min), q.Variant.class()) + attr("data-tc-zone", q.Zone) + attr("data-tc-progress", q.Progress) + q.Fold.attrs() +
		flag("hidden", q.Progress != "" && len(q.Items) == 0) + `><div class="tc-queue__head"><b>` + esc(q.Title) + `</b>` +
		q.Total.html(nil) + q.Toggle.html(nil) + `</div>`)
	if q.Row != nil {
		s.WriteString(`<template data-tc-row>` + q.Row.html() + `</template>`)
	}
	if q.Taken != nil {
		s.WriteString(`<template data-tc-taken>` + q.Taken.html() + `</template>`)
	}
	if len(q.Items) > 0 || q.Row != nil || q.Taken != nil {
		s.WriteString(`<ul class="tc-queue__list">`)
		for _, it := range q.Items {
			s.WriteString(it.html())
		}
		s.WriteString(`</ul>`)
	}
	s.WriteString(`</div>`)
	return s.String()
}

func (f *QueueFold) attrs() string {
	if f == nil {
		return ""
	}
	return attr("data-tc-fold-hide", f.Hide) + attr("data-tc-fold-show", f.Show) + attr("data-tc-fold-total", f.Total)
}

func (it QueueItem) html() string {
	var s strings.Builder
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
		s.WriteString(`<div class="tc-line tc-line--normal tc-queue__acts">`)
		for _, b := range it.Acts {
			s.WriteString(b.html(nil))
		}
		s.WriteString(`</div>`)
	}
	s.WriteString(`</li>`)
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
func (FilePick) emptyPart() {}
func (Line) emptyPart()     {}
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
	return [...]string{"", "tc-empty--t12-pad", "tc-empty--t12-pad", "tc-empty--dusty tc-empty--t12-pad", "tc-empty--dashed",
		"tc-empty--drop", "tc-empty--drop-side"}[v]
}

// Refusal is what a drop sheet says, in red and instead of all its words,
// of a file the form it names does not take (Carrel's task 265): Of is the
// case — an extension («.ics»: a file of it, which another kind takes —
// where to carry it), RefuseOther for any other file the field does not
// accept, RefuseSeveral for several files where the field takes one. What
// the field takes is its own accept and multiple (PostForm.Accept, Many;
// FileField.Accept, Many).
type Refusal struct {
	Of    string
	Parts []Inline
}

// The cases of a refusal that are not an extension.
const (
	RefuseOther   = "*"
	RefuseSeveral = "several"
)

func (r Refusal) html(on bool) string {
	return `<b` + class(nil, is("on", on), "tc-empty__refusal") + ` data-tc-refusal="` + esc(r.Of) + `">` + inlines(r.Parts) + `</b>`
}

// Empty — tc-empty
//
// Роль. Пусто: чего здесь нет и что сделать; место, куда бросить файл.
// Вход. Empty{Icon, Lead, Parts, Refusals, Refused, Variant}.
// Состояния. plain | small | dusty | action | drop | dashed | refused.
// Инварианты. Пустое состояние говорит словами, что пусто; вид — перечисление; оно печатает фрагмент своего кадра.
// Поведение. Бросок файла на зону загружает его; файл, которого форма листа не берёт (не то расширение, несколько вместо одного), лист отказывает сам — одной красной строкой отказа вместо всех слов, и ничего не уходит.
// Эталон. кадр `components-22`, фрагмент `.tc-empty`.
// Содержимое. Значок, крупная первая строка, слова, жирное, моноширинное, клавиша, кнопка, строки отказа броска.
// Чего нет. Картинок-иллюстраций.
type Empty struct {
	Icon  Icon
	Lead  string // a first line in bold
	Parts []EmptyPart
	// Refusals are the drop sheet's words for a file its form does not
	// take, unseen until the library's script refuses one: then the sheet
	// says only the refusal, in red, and sends nothing (Carrel's task 265).
	Refusals []Refusal
	// Refused prints the sheet refusing, with the refusal of this case: the
	// state the script sets, for a frame to draw.
	Refused string
	Variant EmptyVariant
}

// Render prints the empty state.
func (e Empty) Render() string { return e.html(nil) }

func (e Empty) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-empty", e.Variant.class(), is("refused", e.Refused != "")) + `>`)
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
	for _, r := range e.Refusals {
		s.WriteString(r.html(e.Refused != "" && r.Of == e.Refused))
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
	ThumbFolder                   // a folder in a tile: no card, a large folder in the accent colour
)

// Thumb — tc-thumb
//
// Роль. Миниатюра файла: место картинки со значком рода.
// Вход. Thumb{Icon, Form, Src, Alt}.
// Состояния. tile | card | folder | preview | wide | square | image.
// Инварианты. Вид — перечисление; миниатюра печатает фрагмент своего кадра.
// Поведение. Картинку, которая не пришла, сценарий убирает: остаётся значок.
// Эталон. кадр `s-files-5`, фрагмент `.tc-thumb`.
// Содержимое. Значок; картинка (Src) стоит поверх него.
// Чего нет. Своего размера: картинка (Src) заполняет место миниатюры, обрезаясь по нему.
type Thumb struct {
	Icon Icon
	Form ThumbForm
	Src  string // the picture's address: it stands over the icon, which shows when it does not arrive
	Alt  string // what the picture is: the file's name
}

// Render prints the thumbnail.
func (t Thumb) Render() string { return t.html(nil) }

func (t Thumb) html(m mix) string {
	round := `viewBox="0 0 24 24" stroke-linecap="round" stroke-linejoin="round"`
	pic := ""
	if t.Src != "" {
		pic = `<img class="tc-thumb__img"` + attr("src", t.Src) + ` alt="` + esc(t.Alt) + `">`
	}
	switch t.Form {
	case ThumbOnCard:
		return `<div` + class(m, "tc-thumb", "tc-thumb--on-card") + `><svg>` + t.Icon.use() + `</svg>` + pic + `</div>`
	case ThumbPreview:
		return `<div` + class(m, "tc-thumb", "tc-thumb--preview") + `><svg ` + round + `>` + t.Icon.use() + `</svg>` + pic + `</div>`
	case ThumbWide:
		return `<div` + class(m, "tc-thumb", "tc-thumb--wide") + `><svg ` + round + ` class="tc-ico--sq34">` + t.Icon.use() + `</svg>` + pic + `</div>`
	case ThumbFolder:
		return `<div` + class(m, "tc-thumb", "tc-thumb--folder") + `><svg>` + t.Icon.use() + `</svg>` + pic + `</div>`
	case ThumbSquare:
		return `<span` + class(m, "tc-thumb", "tc-thumb--tile") + `><svg ` + round + ` class="tc-ico--sq30">` + t.Icon.use() + `</svg>` + pic + `</span>`
	}
	return `<div` + class(m, "tc-thumb") + `><svg>` + t.Icon.use() + `</svg>` + pic + `</div>`
}

// Tile — tc-tile
//
// Роль. Плитка файла в сетке: миниатюра, имя, размер.
// Вход. Tile{Thumb, Name, Meta, Selected, Href, Send}.
// Состояния. plain | selected | pick.
// Инварианты. Имя — одной строкой с многоточием; плитка печатает фрагмент своего кадра.
// Поведение. Нажатие выбирает плитку: с `Send` она — флажок формы во всю плитку, без своего вида, выбранную рисует рамка; нажатие на имя со ссылкой открывает файл.
// Эталон. кадр `components-25`, фрагмент `.tc-tile`.
// Содержимое. Миниатюра, имя, размер.
// Чего нет. Действий на плитке: они в баре выделения.
type Tile struct {
	Thumb    Thumb
	Name     string
	Meta     string
	Selected bool
	// Href makes the name a link to the file (Carrel's task 114).
	Href string
	// Send makes the tile a checkbox of the form it stands in: its Name and
	// Value go with a bulk action, and the tile is picked without a script.
	Send Send
}

// Render prints the tile.
func (t Tile) Render() string { return t.html(nil) }

func (t Tile) html(m mix) string {
	name, box := "span", ""
	if t.Href != "" {
		name = "a"
	}
	if t.Send.Name != "" {
		box = `<input class="tc-tile__box" type="checkbox"` + attr("form", t.Send.Of) + attr("name", t.Send.Name) + attr("value", t.Send.Value) + attr("aria-label", "Pick "+t.Name) + flag("checked", t.Selected) + `>`
	}
	return `<div` + class(m, "tc-tile", is("sel", t.Selected)) + `>` + box + t.Thumb.html(nil) + `<` + name + ` class="tc-tile__name"` + attr("href", t.Href) + `>` + esc(t.Name) +
		`</` + name + `><span class="tc-tile__meta">` + esc(t.Meta) + `</span></div>`
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
	s.WriteString(`<div` + class(m, "tc-grid", pick(g.Flush, "tc-grid--flush-fill-pad")) + `>`)
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
func (Menu) padPart()  {}

// Pad — tc-pad
//
// Роль. Отступ страницы вокруг того, у чего своих полей нет: карточки, превью, меню на странице.
// Вход. Pad{Items}.
// Состояния. plain.
// Инварианты. Поля — токен страницы; блок печатает фрагмент своего кадра.
// Поведение. нет.
// Эталон. кадр `s-files-1`, фрагмент `.tc-pad`.
// Содержимое. Карточки, превью, меню, стоящее на странице (история запросов без запроса).
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
// Вход. Diff{Items, Action, Method, Fields, ID}.
// Состояния. plain.
// Инварианты. У свойства выбран ровно один вариант; варианты — столбцом; сравнение печатает фрагмент своего кадра.
// Поведение. Нажатие выбирает вариант свойства.
// Эталон. кадр `s-conflict-1`, фрагмент `.tc-diff`.
// Содержимое. Свойства: имя моноширинным, почему расходятся, варианты.
// Чего нет. Слияния значений: выбирается одно.
type Diff struct {
	Items []DiffItem
	// Action makes the comparison a form (Carrel's task 132): each property's
	// versions are radios of its own name (Check.Radio), «Write this
	// combination» sends them; Method, Fields and ID as Form's.
	Action string
	Method string
	Fields []Hidden
	ID     string
}

// Render prints the comparison.
func (d Diff) Render() string { return d.html(nil) }

func (d Diff) html(m mix) string {
	var s strings.Builder
	tag := "div"
	if d.Action != "" {
		tag = "form"
		s.WriteString(`<form` + class(m, "tc-diff") + attr("id", d.ID) + attr("action", d.Action) + ` method="` + method(d.Method) + `">`)
		for _, f := range d.Fields {
			s.WriteString(f.html(nil))
		}
	} else {
		s.WriteString(`<div` + class(m, "tc-diff") + `>`)
	}
	for _, it := range d.Items {
		s.WriteString(`<div class="tc-diff__item">` + it.Title.html(nil) + `<div class="tc-diff__options">`)
		for _, o := range it.Options {
			s.WriteString(o.html(nil))
		}
		s.WriteString(`</div></div>`)
	}
	s.WriteString(`</` + tag + `>`)
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
