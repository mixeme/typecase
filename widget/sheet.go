package widget

import "strings"

// SheetIconTone is the colour of a sheet's icon: what kind of question the
// sheet asks.
type SheetIconTone int

const (
	SheetIconAccent SheetIconTone = iota // an ordinary task: move, attach, create
	SheetIconDanger                      // a failure or a deletion
	SheetIconInk                         // a grave step that is not a failure: handing over a key
)

func (t SheetIconTone) svg() string {
	return [...]string{
		`<svg width="15" height="15" viewBox="0 0 24 24" stroke-width="1.4" fill="none" stroke-linecap="round" stroke-linejoin="round" class="tc-ico--accent">`,
		`<svg width="15" height="15" viewBox="0 0 24 24" stroke-width="1.6" fill="none" class="tc-ico--danger">`,
		`<svg width="15" height="15" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.4" fill="none">`}[t]
}

// SheetHeadPart is what stands after a sheet's title.
type SheetHeadPart interface {
	part
	sheetHeadPart()
}

func (Dim) sheetHeadPart()          {}
func (Pick) sheetHeadPart()         {}
func (Filter) sheetHeadPart()       {}
func (Button) sheetHeadPart()       {}
func (Close) sheetHeadPart()        {}
func (SheetSubject) sheetHeadPart() {}
func (Tag) sheetHeadPart()          {}
func (Crumbs) sheetHeadPart()       {}

// SheetSubject is what a sheet acts on, named in its head: a source's mark
// and the thing's name.
type SheetSubject struct {
	Mark Mark
	Text string
}

func (t SheetSubject) html(mix) string {
	return `<span class="tc-table__file tc-push">` + t.Mark.html(mix{"tc-table__filename"}) +
		`<span class="tc-table__filename tc-table__filename--t12">` + esc(t.Text) + `</span></span>`
}

// SheetPart is what a sheet's body holds.
type SheetPart interface {
	part
	sheetPart()
}

func (Rubric) sheetPart()      {}
func (Tree) sheetPart()        {}
func (Field) sheetPart()       {}
func (Area) sheetPart()        {}
func (Line) sheetPart()        {}
func (Dim) sheetPart()         {}
func (RailSec) sheetPart()     {}
func (Form) sheetPart()        {}
func (SheetFields) sheetPart() {}
func (SheetKeys) sheetPart()   {}
func (SheetHint) sheetPart()   {}
func (SheetNote) sheetPart()   {}
func (SheetBlock) sheetPart()  {}
func (SheetChecks) sheetPart() {}
func (SheetFrom) sheetPart()   {}
func (SheetPick) sheetPart()   {}

// SheetFields is a row of fields in a sheet: the first takes the rest of it.
type SheetFields struct {
	Items []Field
}

func (f SheetFields) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-line tc-sheet__fields">`)
	for _, it := range f.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// SheetKey is a line of a table of keys: the keys and what they do.
type SheetKey struct {
	Keys string
	Does string
}

// SheetKeys is a table of keys in a sheet.
type SheetKeys struct {
	Items []SheetKey
}

func (k SheetKeys) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-sheet__keys">`)
	for _, it := range k.Items {
		s.WriteString(Kbd{Keys: it.Keys}.html(nil) + `<span class="tc-stack__item">` + esc(it.Does) + `</span>`)
	}
	s.WriteString(`</div>`)
	return s.String()
}

// SheetHint is a line of explanation under a sheet's table of keys.
type SheetHint struct {
	Text string
}

func (h SheetHint) html(mix) string {
	return `<span class="tc-field__hint tc-field__hint--top6">` + esc(h.Text) + `</span>`
}

// SheetNote is a muted paragraph of a sheet: what a failure means.
type SheetNote struct {
	Text string
}

func (n SheetNote) html(mix) string {
	return `<span class="tc-dim--muted tc-dim--t11-5 tc-dim--lh1-55 tc-dim--wrap">` + esc(n.Text) + `</span>`
}

// SheetLine is a line of a sheet's list: a name in bold and what happens to
// it.
type SheetLine struct {
	Name string
	Note Dim
}

// SheetBlock is a list under a rubric in a sheet: what points at a thing
// that is about to go.
type SheetBlock struct {
	Title Rubric
	Items []SheetLine
}

func (b SheetBlock) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-sheet__block">` + b.Title.html(nil) + `<div class="tc-stack tc-stack--gap-6 tc-line__text">`)
	for _, it := range b.Items {
		s.WriteString(`<span class="tc-stack__item"><b>` + esc(it.Name) + `</b> ` + it.Note.html(nil) + `</span>`)
	}
	s.WriteString(`</div></div>`)
	return s.String()
}

// SheetChecks is a column of checks at the end of a sheet's form.
type SheetChecks struct {
	Items []Check
}

func (c SheetChecks) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-stack tc-stack--gap-5 tc-sheet__checks">`)
	for _, it := range c.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// SheetFrom says in whose name a letter goes: a label and lines of text.
type SheetFrom struct {
	Label string
	Lines [][]Inline
}

func (f SheetFrom) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-line tc-line--gap-8 tc-sheet__from">` + Dim{Parts: []Inline{Text(f.Label)}, Variant: DimFixed}.html(nil) + `<span class="tc-line__item">`)
	for i, l := range f.Lines {
		if i > 0 {
			s.WriteString(`<br>`)
		}
		s.WriteString(inlines(l))
	}
	s.WriteString(`</span></div>`)
	return s.String()
}

// SheetPickRail is what the rail of a sheet's picker holds.
type SheetPickRail interface {
	part
	sheetPickRail()
}

func (RailSec) sheetPickRail() {}
func (Tree) sheetPickRail()    {}

// SheetPickList is what the list of a sheet's picker holds.
type SheetPickList interface {
	part
	sheetPickList()
}

func (ListGroup) sheetPickList() {}
func (List) sheetPickList()      {}
func (Crumbs) sheetPickList()    {}
func (Form) sheetPickList()      {}
func (Table) sheetPickList()     {}

// SheetPick is a picker in a sheet: sources on the left, what is in them on
// the right.
type SheetPick struct {
	Rail  []SheetPickRail
	List  []SheetPickList
	Short bool // 196 px high: a picker of folders
}

func (p SheetPick) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(nil, "tc-line", "tc-sheet__pick", pick(p.Short, "tc-sheet__pick--196")) + `><div class="tc-sheet__pick-rail">`)
	for _, it := range p.Rail {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div><div class="tc-sheet__pick-list">`)
	for _, it := range p.List {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div></div>`)
	return s.String()
}

// SheetFootPart is what a sheet's foot holds.
type SheetFootPart interface {
	part
	sheetFootPart()
}

func (Button) sheetFootPart() {}
func (Kbd) sheetFootPart()    {}
func (Dim) sheetFootPart()    {}
func (Line) sheetFootPart()   {}
func (Push) sheetFootPart()   {}

// SheetWidth is the width of a sheet.
type SheetWidth int

const (
	SheetWidth560 SheetWidth = iota
	SheetWidth400
	SheetWidth480
	SheetWidth520
	SheetWidth680
	SheetWidthFull // the width of the page: a phone, a picker
)

func (w SheetWidth) class() string {
	return [...]string{"", "tc-sheet--w400", "tc-sheet--w480", "tc-sheet--w520", "tc-sheet--w680", "tc-sheet--max100pc-full"}[w]
}

// Sheet — tc-sheet
//
// Роль. Лист поверх страницы: быстрая запись, перенос, выбор, подтверждение — одно дело и выход из него.
// Вход. Sheet{Icon, Tone, Title, Head, Body, Flush, Foot, Width}.
// Состояния. plain | icon | close | danger | pick.
// Инварианты. Заголовок — один, жирным; главная кнопка — первой в подвале, отмена — в конце; ширина — перечисление; лист печатает фрагмент своего кадра.
// Поведение. Esc закрывает лист и возвращает фокус тому, что его открыло; Enter делает главное.
// Эталон. кадр `s-contacts-2`, фрагмент `.tc-sheet`.
// Содержимое. Значок и заголовок, прижатое к концу шапки (пояснение, выбор, фильтр, кнопка, предмет), крестик; тело — поля, дерево, выбор, таблица клавиш, списки, сообщения; подвал — кнопки, клавиши, пояснение.
// Чего нет. Второго листа поверх: следующий шаг открывается вместо него.
type Sheet struct {
	Icon  Icon // none: a sheet that only asks
	Tone  SheetIconTone
	Title string
	Head  []SheetHeadPart
	Body  []SheetPart
	Flush bool // the body has no padding: a picker fills it
	Foot  []SheetFootPart
	Width SheetWidth
}

// Render prints the sheet.
func (h Sheet) Render() string { return h.html(nil) }

func (h Sheet) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-sheet", h.Width.class()) + ` role="dialog" aria-modal="true"` + attr("aria-label", h.Title) + `><div class="tc-sheet__head">`)
	if h.Icon != "" {
		s.WriteString(h.Tone.svg() + h.Icon.use() + `</svg>`)
	}
	s.WriteString(`<b class="tc-sheet__title">` + esc(h.Title) + `</b>`)
	for _, it := range h.Head {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div><div` + class(nil, "tc-sheet__body", pick(h.Flush, "tc-sheet__body--pad0")) + `>`)
	for _, it := range h.Body {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	if len(h.Foot) > 0 {
		s.WriteString(`<div class="tc-sheet__foot">`)
		for _, it := range h.Foot {
			s.WriteString(it.html(nil))
		}
		s.WriteString(`</div>`)
	}
	s.WriteString(`</div>`)
	return s.String()
}

// ScrimPart is what stands on a scrim.
type ScrimPart interface {
	part
	scrimPart()
}

func (Sheet) scrimPart() {}

// ScrimPad is how far from the top a scrim sets its sheet.
type ScrimPad int

const (
	ScrimPadDefault ScrimPad = iota
	ScrimPad26
	ScrimPad38
)

func (p ScrimPad) class() string {
	return [...]string{"", "tc-scrim--padt26", "tc-scrim--padt38"}[p]
}

// Scrim — tc-scrim
//
// Роль. Затемнение страницы под листом: страница видна, но недоступна, пока лист открыт.
// Вход. Scrim{Content, Pad, ID, Closed}.
// Состояния. plain | pad.
// Инварианты. На затемнении одно: лист или пустое состояние; отступ сверху — перечисление; затемнение печатает фрагмент своего кадра.
// Поведение. Нажатие на затемнение закрывает лист, как Esc; кнопка с Opens открывает закрытое затемнение на месте.
// Эталон. кадр `s-contacts-2`, фрагмент `.tc-scrim`.
// Содержимое. Лист.
// Чего нет. Своего содержимого: только то, что на нём стоит. Без сценария закрытое затемнение не открыть: экран тогда открывает лист своим адресом.
type Scrim struct {
	Content ScrimPart
	Pad     ScrimPad
	ID      string // what a button's Opens names
	Closed  bool   // printed closed, for a button to open in place
}

// Render prints the scrim.
func (c Scrim) Render() string { return c.html(nil) }

func (c Scrim) html(m mix) string {
	s := `<div` + class(m, "tc-scrim", c.Pad.class()) + attr("id", c.ID) + flag("hidden", c.Closed) + `>`
	if c.Content != nil {
		s += c.Content.html(nil)
	}
	return s + `</div>`
}

// OverlayRailPart is what the rail under an overlay holds.
type OverlayRailPart interface {
	part
	overlayRailPart()
}

func (Nav) overlayRailPart()      {}
func (RailSec) overlayRailPart()  {}
func (RailFoot) overlayRailPart() {}

// OverlayBody is the shell's body under an overlay: the rail and the page,
// dimmed; the side panel is closed while a sheet is open.
type OverlayBody struct {
	Rail   []OverlayRailPart
	Main   Main
	Height ShellHeight
}

func (b OverlayBody) html() string {
	rail := make([]part, len(b.Rail))
	for i, r := range b.Rail {
		rail[i] = r
	}
	return shellBody(rail, RailFull, b.Main, nil, b.Height)
}

// Overlay — tc-overlay
//
// Роль. Оболочка, над которой открыт лист: тело оболочки под затемнением и лист на нём.
// Вход. Overlay{Top, Body, Scrim}.
// Состояния. plain | top.
// Инварианты. Под затемнением — то же тело оболочки, что и без него; наложение печатает фрагмент своего кадра.
// Поведение. Пока лист открыт, фокус не уходит под затемнение.
// Эталон. кадр `s-contacts-2`, фрагмент `.tc-overlay`.
// Содержимое. Верхняя полоса (когда она тоже под затемнением), тело оболочки, затемнение с листом.
// Чего нет. Своего вида: вид у затемнения и листа.
type Overlay struct {
	Top   *Top
	Body  OverlayBody
	Scrim Scrim
}

// Render prints the overlay.
func (o Overlay) Render() string { return o.html(nil) }

func (o Overlay) html(m mix) string {
	s := `<div` + class(m, "tc-overlay") + `>`
	if o.Top != nil {
		s += o.Top.html(nil)
	}
	return s + o.Body.html() + o.Scrim.html(nil) + `</div>`
}

func (Overlay) shellContent() {}
func (Scrim) mainPart()       {}
