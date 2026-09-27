package refs

import w "carrel.dev/typecase/widget"

func overlay(state string, fixture Renderer, frame, class string, n int) Entry {
	return e("overlays", state, fixture, ref(frame, class, n))
}

var _ = add(
	// tc-menu
	overlay("icons", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "move", Label: "Move"},
		w.MenuItem{Icon: "dup", Label: "Copy"},
		w.MenuItem{Icon: "down", Detail: &w.Mono{Text: ".vcf", Variant: w.Mono11Dusty}, Label: "Export"},
		w.MenuItem{Icon: "link", Label: "Publish a link"}}, Width: w.MenuWidth212}, "s-contacts-8", "tc-menu", 1),
	overlay("danger", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "con", Label: "Open"},
		w.MenuItem{Icon: "pencil", Label: "Edit"},
		w.MenuItem{Icon: "move", Label: "Move"},
		w.MenuItem{Icon: "dup", Label: "Copy"},
		w.MenuItem{Icon: "down", Detail: &w.Mono{Text: ".vcf", Variant: w.Mono11Dusty}, Label: "Export"},
		w.MenuItem{Icon: "link", Label: "Publish a link"},
		w.MenuSep{},
		w.MenuItem{Icon: "trash", Label: "Delete contact", Tone: w.MenuToneDanger}}, Width: w.MenuWidth212}, "components-1", "tc-menu", 1),
	overlay("icons", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "doc", Label: "Files"},
		w.MenuItem{Icon: "file", Label: "Folder"}}, Width: w.MenuWidth212}, "components-7", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "nextcloud", On: true},
		w.MenuItem{Label: "radicale · vps"}}, Width: w.MenuWidth212}, "s-profile-13", "tc-menu", 1),
	overlay("danger", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "doc", Label: "Open"},
		w.MenuItem{Icon: "down", Label: "Download"},
		w.MenuItem{Icon: "pencil", Label: "Rename"},
		w.MenuItem{Icon: "move", Label: "Move"},
		w.MenuItem{Icon: "dup", Label: "Copy"},
		w.MenuItem{Icon: "link", Label: "Publish a link"},
		w.MenuSep{},
		w.MenuItem{Icon: "trash", Label: "Delete", Tone: w.MenuToneDanger}}, Width: w.MenuWidth212}, "components-5", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "work", On: true},
		w.MenuItem{Label: "home"},
		w.MenuItem{Label: "other"}}, Width: w.MenuWidth212}, "s-contacts-13", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "At the time"},
		w.MenuItem{Label: "5 minutes before"},
		w.MenuItem{Label: "10 minutes before", On: true},
		w.MenuItem{Label: "30 minutes before"},
		w.MenuItem{Label: "1 hour before"},
		w.MenuItem{Label: "1 day before"},
		w.MenuSep{},
		w.MenuItem{Label: "No reminder"}}, Width: w.MenuWidth212}, "s-agenda-9", "tc-menu", 1),
	overlay("marks", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Записи", On: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "nextcloud · Дневник"}}, Width: w.MenuWidth290}, "s-notes-15", "tc-menu", 1),
	overlay("marks", w.Menu{Items: []w.MenuPart{
		w.Rubric{Text: "Write into", Form: w.RubricInline},
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Trail: w.Mono{Text: "last used", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Дела", On: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "Работа"}}, Width: w.MenuWidth290}, "s-tasks-5", "tc-menu", 1),
	overlay("icons", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "move", Label: "Move"},
		w.MenuItem{Icon: "dup", Label: "Copy"},
		w.MenuItem{Icon: "link", Label: "Publish a link"}}, Width: w.MenuWidth212}, "s-notes-14", "tc-menu", 1),
	overlay("marks", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Контакты", On: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "nextcloud · Архив"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "radicale · vps · Команда"}}, Width: w.MenuWidth290}, "s-notes-16", "tc-menu", 1),
	overlay("marks", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Личный", On: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа"},
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "nextcloud · Дни рождения"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "radicale · vps · Проекты"},
		w.MenuItem{Mark: &w.Mark{Colour: "#8C6239"}, Label: "radicale · vps · Дежурства"},
		w.MenuItem{Mark: &w.Mark{Colour: "#9A9280"}, Label: "baikal · Общий"}}, Width: w.MenuWidth290}, "s-notes-17", "tc-menu", 1),
	overlay("marks", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Записи", On: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "nextcloud · Дневник"},
		w.MenuSep{},
		w.MenuItem{Label: "New notebook"}}, Width: w.MenuWidth290}, "s-io-11", "tc-menu", 1),
	overlay("pop", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "pencil", Label: "Rename"},
		w.MenuItem{Icon: "down", Detail: &w.Mono{Text: ".vcf", Variant: w.Mono11Dusty}, Label: "Export"}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}, "components-19", "tc-menu", 1),
	overlay("pop", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "All", On: true},
		w.MenuItem{Label: "Overdue"}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}, "components-19", "tc-menu", 2),
	overlay("heading", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Write into", Form: w.RubricItem}},
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Trail: w.Mono{Text: "last used", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Контакты", On: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "Архив"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "Команда"}}}, "components-6", "tc-menu", 1),
	overlay("icons", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "up", Label: "Import"},
		w.MenuItem{Icon: "down", Label: "Export"},
		w.MenuItem{Icon: "print", Label: "Print"}}}, "components-8", "tc-menu", 1),
	overlay("heading", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Columns", Form: w.RubricItem}, Form: w.MenuHeadingRoomy},
		w.MenuItem{Icon: "drag", Faint: true, Tick: &w.TickBox{}, Trail: w.Dim{Parts: []w.Inline{w.Text("always")}, Size: w.Dim10_5, Place: w.Pushed}, Label: "Name"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Phone"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Email"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Organisation"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Birthday"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Categories"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Changed"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Source"},
		w.MenuSep{},
		w.MenuItem{Label: "Reset to defaults", Tone: w.MenuToneAccent}}}, "components-21", "tc-menu", 1),
	overlay("marks", w.Menu{Items: []w.MenuPart{
		w.Rubric{Text: "Write into", Form: w.RubricInline},
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Trail: w.Mono{Text: "last used", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Контакты", On: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "Архив"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "Команда"}}, Width: w.MenuWidth290}, "s-contacts-9", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "mobile", On: true},
		w.MenuItem{Label: "work"},
		w.MenuItem{Label: "home"},
		w.MenuItem{Label: "main"},
		w.MenuItem{Label: "fax"},
		w.MenuItem{Label: "other"}}, Width: w.MenuWidth212}, "s-contacts-12", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "All 24", On: true},
		w.MenuItem{Label: "Events 14"},
		w.MenuItem{Label: "Tasks 2"},
		w.MenuItem{Label: "Notes 4"},
		w.MenuItem{Label: "Files 4"},
		w.MenuSep{},
		w.MenuItem{Icon: "cols", Label: "Columns"}}, Width: w.MenuWidth212}, "s-contacts-14", "tc-menu", 1),
	overlay("pop", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Tick: &w.TickBox{Ticked: w.TickedOn}, Label: "Events", Tertiary: true},
		w.MenuItem{Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Tasks due", Tertiary: true},
		w.MenuItem{Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Notes", Tertiary: true},
		w.MenuSep{Tertiary: true},
		w.Range{From: w.Input{Kind: w.InputDate, Value: "2026-08-10", Mono: true, Label: "Date"}, To: w.Input{Kind: w.InputDate, Value: "2026-08-16", Mono: true, Label: "Date"}, Show: w.Button{Label: "Show"}, Wrap: true},
		w.MenuSep{},
		w.MenuItem{Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Compact rows"}}, Pop: w.MenuUnderEnd, Inline: true}, "s-agenda-1", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "No status"},
		w.MenuSep{},
		w.MenuItem{Label: "Tentative"},
		w.MenuItem{Label: "Confirmed", On: true},
		w.MenuItem{Label: "Cancelled"}}, Width: w.MenuWidth212}, "s-agenda-7", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.Input{Kind: w.InputSearch, Placeholder: "Find a zone", Label: "Find a zone"},
		w.Rubric{Text: "As sent", Form: w.RubricInline},
		w.MenuItem{Label: "Europe/Moscow", On: true},
		w.MenuSep{},
		w.MenuItem{Label: "Asia/Yekaterinburg"},
		w.MenuItem{Label: "UTC"},
		w.MenuSep{},
		w.MenuItem{Label: "Floating — the same clock time anywhere"}}, Width: w.MenuWidth290}, "s-agenda-8", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "Does not repeat"},
		w.MenuSep{},
		w.MenuItem{Label: "Daily"},
		w.MenuItem{Label: "Weekly", On: true},
		w.MenuItem{Label: "Monthly"},
		w.MenuItem{Label: "Yearly"}}, Width: w.MenuWidth212}, "s-agenda-10", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "Needs action", On: true},
		w.MenuItem{Label: "In process"},
		w.MenuItem{Label: "Completed"},
		w.MenuItem{Label: "Cancelled"}}, Width: w.MenuWidth212}, "s-tasks-8", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "No priority"},
		w.MenuSep{},
		w.MenuItem{Label: "High · 1", On: true},
		w.MenuItem{Label: "Medium · 5"},
		w.MenuItem{Label: "Low · 9"}}, Width: w.MenuWidth212}, "s-tasks-9", "tc-menu", 1),
	overlay("danger", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "move", Label: "Move"},
		w.MenuItem{Icon: "dup", Label: "Copy"},
		w.MenuItem{Icon: "down", Detail: &w.Mono{Text: ".md", Variant: w.Mono11Dusty}, Label: "Export"},
		w.MenuItem{Icon: "link", Label: "Publish a link"},
		w.MenuSep{},
		w.MenuItem{Icon: "trash", Label: "Delete note", Tone: w.MenuToneDanger}}, Width: w.MenuWidth212}, "s-notes-12", "tc-menu", 1),
	overlay("pop", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "cal", Trail: w.Mono{Text: "10 Aug · Личный", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Звонок с Ковалёвой", On: true},
		w.MenuItem{Icon: "task", Trail: w.Mono{Text: "task · Работа", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Ответить Анне по смете"},
		w.MenuItem{Icon: "con", Trail: w.Mono{Text: "contact", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Анна Ковалёва"},
		w.MenuSep{},
		w.MenuItem{Label: "Enter a UID by hand", Tone: w.MenuToneDusty}}, Width: w.MenuWidthFull, Pop: w.MenuUnder}, "s-notes-4", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "quill", Trail: w.Kbd{Keys: "N"}, Label: "Note", On: true},
		w.MenuSep{},
		w.MenuItem{Icon: "con", Trail: w.Kbd{Keys: "C"}, Label: "Contact"},
		w.MenuItem{Icon: "cal", Trail: w.Kbd{Keys: "E"}, Label: "Event"},
		w.MenuItem{Icon: "task", Trail: w.Kbd{Keys: "T"}, Label: "Task"},
		w.MenuSep{},
		w.MenuItem{Icon: "up", Label: "Upload file"},
		w.MenuSep{},
		w.MenuItem{Icon: "mail", Label: "Letter"}}}, "s-notes-7", "tc-menu", 1),
	overlay("heading", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Recent searches", Form: w.RubricItem}},
		w.MenuItem{Icon: "clock", Trail: w.Mono{Text: "8 found · 2 min ago", Variant: w.Mono11DustyTight, Place: w.Pushed}, Label: "ковал"},
		w.MenuItem{Icon: "clock", Trail: w.Mono{Text: "3 found · yesterday", Variant: w.Mono11DustyTight, Place: w.Pushed}, Label: "смета"},
		w.MenuItem{Icon: "clock", Trail: w.Mono{Text: "nothing found · Monday", Variant: w.Mono11DustyTight, Place: w.Pushed}, Label: "дача"},
		w.MenuSep{},
		w.MenuItem{Trail: w.Kbd{Keys: "Enter"}, Label: "Type to search everything", Tone: w.MenuToneDusty}}, Width: w.MenuWidthFull}, "s-find-2", "tc-menu", 1),
	overlay("found", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Jump to", Form: w.RubricItem}},
		w.MenuItem{Icon: "con", Found: []w.Inline{w.Text("Анна "), w.Strong("Ковал"), w.Text("ёва")}, Trail: w.Dim{Parts: []w.Inline{w.Text("nextcloud · Контакты")}, Size: w.Dim11, Variant: w.DimTight, Place: w.Pushed}},
		w.MenuItem{Icon: "cal", Found: []w.Inline{w.Text("Звонок с "), w.Strong("Ковал"), w.Text("ёвой")}, Trail: w.Dim{Parts: []w.Inline{w.Text("10 Aug · Личный")}, Size: w.Dim11, Variant: w.DimTight, Place: w.Pushed}},
		w.MenuItem{Icon: "file", Found: []w.Inline{w.Text("договор-"), w.Strong("ковал"), w.Text("ёва.pdf")}, Trail: w.Dim{Parts: []w.Inline{w.Text("sftpgo · vps · Documents")}, Size: w.Dim11, Variant: w.DimTight, Place: w.Pushed}},
		w.MenuSep{},
		w.MenuItem{Icon: "search", Trail: w.Kbd{Keys: "Enter"}, Label: "Show all results for «ковал»", On: true},
		w.MenuItem{Icon: "all", Trail: w.Dim{Parts: []w.Inline{w.Text("asks every server")}, Size: w.Dim11, Variant: w.DimTight, Place: w.Pushed}, Label: "Timeline · Анна Ковалёва"},
		w.MenuSep{},
		w.MenuHeading{Rubric: w.Rubric{Text: "Recent searches", Form: w.RubricItem}},
		w.MenuItem{Icon: "clock", Trail: w.Mono{Text: "3 found · yesterday", Variant: w.Mono11DustyTight, Place: w.Pushed}, Label: "смета"}}, Width: w.MenuWidthFull}, "s-find-3", "tc-menu", 1),
	overlay("fields", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Points per signal", Items: []w.RubricPart{w.Push{Items: []w.PushPart{w.Text("Default")}}, w.RubricCol{Text: "Yours"}}, Form: w.RubricItem}, Form: w.MenuHeadingRoomy},
		w.MenuField{Label: "Same UID", Default: w.Dim{Parts: []w.Inline{w.Text("80")}, Size: w.Dim10_5, Place: w.Pushed}, Input: w.Input{Value: "80", Mono: true, Variant: w.InputChip, Label: "Same UID"}},
		w.MenuField{Label: "Same email", Default: w.Dim{Parts: []w.Inline{w.Text("60")}, Size: w.Dim10_5, Place: w.Pushed}, Input: w.Input{Value: "60", Mono: true, Variant: w.InputChip, Label: "Same email"}},
		w.MenuField{Label: "Same phone", Default: w.Dim{Parts: []w.Inline{w.Text("60")}, Size: w.Dim10_5, Place: w.Pushed}, Input: w.Input{Value: "60", Mono: true, Variant: w.InputChip, Label: "Same phone"}},
		w.MenuField{Label: "Same name", Default: w.Dim{Parts: []w.Inline{w.Text("20")}, Size: w.Dim10_5, Place: w.Pushed}, Input: w.Input{Value: "20", Mono: true, Variant: w.InputChip, Label: "Same name"}},
		w.MenuField{Label: "Same birthday", Badge: &w.Badge{Label: "changed"}, Default: w.Dim{Parts: []w.Inline{w.Text("10")}, Size: w.Dim10_5, Place: w.Pushed}, Input: w.Input{Value: "20", Mono: true, Variant: w.InputChip, Label: "Same birthday"}},
		w.MenuField{Label: "Same title and date (events, tasks, notes)", Default: w.Dim{Parts: []w.Inline{w.Text("70")}, Size: w.Dim10_5, Place: w.Pushed}, Input: w.Input{Value: "70", Mono: true, Variant: w.InputChip, Label: "Same title and date (events, tasks, notes)"}},
		w.MenuFoot{Items: []w.Button{w.Button{Label: "Save", Kind: w.ButtonPrimary, Small: true}, w.Button{Label: "Back to defaults", Small: true}}}}, Width: w.MenuWidth420, Fields: true}, "s-dup-3", "tc-menu", 1),
	overlay("marks", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Контакты", On: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "nextcloud · Архив"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "radicale · vps · Команда"},
		w.MenuSep{},
		w.MenuItem{Label: "New address book"}}, Width: w.MenuWidth290}, "s-io-10", "tc-menu", 1),
	overlay("marks", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Личный", On: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа"},
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "nextcloud · Дни рождения"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "radicale · vps · Проекты"},
		w.MenuItem{Mark: &w.Mark{Colour: "#8C6239"}, Label: "radicale · vps · Дежурства"},
		w.MenuItem{Mark: &w.Mark{Colour: "#9A9280"}, Label: "baikal · Общий"},
		w.MenuSep{},
		w.MenuItem{Label: "New calendar"}}, Width: w.MenuWidth290}, "s-io-12", "tc-menu", 1),
	overlay("heading", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "mix", Items: []w.RubricPart{w.Mono{Text: "carrel.example", Variant: w.MonoTightPlain}}, Form: w.RubricItem, Variant: w.RubricItemLine}, Form: w.MenuHeadingSplit},
		w.MenuItem{Icon: "gear", Label: "Settings"},
		w.MenuItem{Icon: "adm", Label: "Administration"},
		w.MenuSep{},
		w.MenuColumn{Label: "Theme", Control: w.Seg{Options: []w.SegOption{{Label: "Auto", On: true}, {Label: "Light"}, {Label: "Dark"}}, Variant: w.SegStartSmall}},
		w.MenuItem{Icon: "search", Trail: w.Kbd{Keys: "?"}, Label: "Keyboard shortcuts"},
		w.MenuSep{},
		w.MenuItem{Label: "Sign out"}}, Width: w.MenuWidth290}, "account-5", "tc-menu", 1),
	overlay("pop", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "pencil", Label: "Rename and colour"},
		w.MenuItem{Icon: "move", Label: "Move to another server"},
		w.MenuItem{Icon: "dup", Label: "Copy to another server"},
		w.MenuItem{Icon: "down", Detail: &w.Mono{Text: ".ics", Variant: w.Mono11Dusty}, Label: "Export"},
		w.MenuItem{Icon: "link", Label: "Publish a link"},
		w.MenuSep{},
		w.MenuItem{Icon: "trash", Label: "Delete on the server…", Tone: w.MenuToneDanger}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}, "s-profile-1", "tc-menu", 1),
	overlay("danger", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "pencil", Label: "Rename"},
		w.MenuItem{Icon: "move", Label: "Move to another server"},
		w.MenuItem{Icon: "dup", Label: "Copy to another server"},
		w.MenuItem{Icon: "link", Label: "Publish a link"},
		w.MenuSep{},
		w.MenuItem{Icon: "trash", Label: "Delete on the server…", Tone: w.MenuToneDanger}}, Width: w.MenuWidth212}, "s-profile-12", "tc-menu", 1),
	overlay("marks", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Files"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "sftpgo · vps · Documents", On: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#9A9280"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "sftpgo · vps · Public files"}}, Width: w.MenuWidth290}, "attachments-2", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "Date, then the original name", On: true},
		w.MenuItem{Label: "The original name"},
		w.MenuItem{Label: "Record title, then the original name"}}, Width: w.MenuWidth290}, "attachments-3", "tc-menu", 1),
	overlay("pop", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "radicale · vps · Проекты", On: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Личный"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа"},
		w.MenuSep{},
		w.MenuItem{Icon: "plus", Label: "New collection"},
		w.MenuItem{Label: "Skip this one", Tone: w.MenuToneDusty}}, Pop: w.MenuUnder}, "s-backup-8", "tc-menu", 1),
	overlay("danger", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "mail", Label: "Send"},
		w.MenuSep{},
		w.MenuItem{Label: "Revoke", Tone: w.MenuToneDanger}}, Width: w.MenuWidth212}, "s-publish-15", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "7 days"},
		w.MenuItem{Label: "14 days"},
		w.MenuItem{Label: "30 days", On: true},
		w.MenuSep{},
		w.MenuItem{Label: "No expiry"}}, Width: w.MenuWidth212}, "s-publish-14", "tc-menu", 1),
	overlay("icons", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "con", Label: "Contact"},
		w.MenuItem{Icon: "con", Label: "Address book"},
		w.MenuItem{Icon: "cal", Label: "Event"},
		w.MenuItem{Icon: "cal", Label: "Calendar"},
		w.MenuItem{Icon: "task", Label: "Task"},
		w.MenuItem{Icon: "note", Label: "Note"},
		w.MenuSep{},
		w.MenuItem{Icon: "doc", Label: "File"},
		w.MenuItem{Icon: "file", Trail: w.Tag{Label: "link", Place: w.Pushed}, Label: "Folder"}}}, "s-mail-2", "tc-menu", 1),
	overlay("marks", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Detail: &w.Mono{Text: "a.belov@example.org", Variant: w.Mono11Dusty}, Trail: w.Tag{Label: "work", Place: w.Pushed}, Label: "Артём Белов", On: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Detail: &w.Mono{Text: "belov@personal.example", Variant: w.Mono11Dusty}, Trail: w.Tag{Label: "home", Place: w.Pushed}, Label: "Артём Белов"},
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Detail: &w.Mono{Text: "b.yankovskaya@example.org", Variant: w.Mono11Dusty}, Label: "Белла Янковская"},
		w.MenuSep{},
		w.MenuItem{Icon: "mail", Label: "Use «бел» as typed", Tone: w.MenuToneMuted},
		w.MenuSep{},
		w.MenuItem{Note: &w.Span{Face: w.FaceDanger, Text: "baikal · Офис did not answer"}, Label: "Searched 4 of 5 address books ·", Tone: w.MenuToneAside}}, Width: w.MenuWidthFull}, "s-mail-3", "tc-menu", 1),
	overlay("danger", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "Extend"},
		w.MenuSep{},
		w.MenuItem{Label: "Revoke", Tone: w.MenuToneDanger}}, Width: w.MenuWidth212}, "s-admin-11", "tc-menu", 1),
	overlay("danger", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "Set role"},
		w.MenuItem{Label: "Kill sessions"},
		w.MenuItem{Label: "Disable"},
		w.MenuSep{},
		w.MenuItem{Icon: "trash", Label: "Delete", Tone: w.MenuToneDanger}}, Width: w.MenuWidth212}, "s-admin-12", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "STARTTLS", On: true},
		w.MenuItem{Label: "TLS"},
		w.MenuSep{},
		w.MenuItem{Label: "None"}}, Width: w.MenuWidth212}, "s-admin-9", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "All actions", On: true},
		w.MenuSep{},
		w.MenuItem{Label: "login"},
		w.MenuItem{Label: "login_failed"},
		w.MenuItem{Label: "escrow_recover"},
		w.MenuItem{Label: "settings_update"}}, Width: w.MenuWidth290}, "s-admin-10", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "con", Trail: w.Mono{Text: "nextcloud · Контакты", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Анна Ковалёва", On: true},
		w.MenuItem{Icon: "con", Trail: w.Mono{Text: "nextcloud · Контакты", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Пётр Ковалёв"},
		w.MenuItem{Icon: "con", Trail: w.Mono{Text: "radicale · vps · Команда", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Ковалёв Игорь"}}, Width: w.MenuWidthFull}, "mentions-3", "tc-menu", 1),
	overlay("heading", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "mix", Items: []w.RubricPart{w.Mono{Text: "carrel.example", Variant: w.MonoTightPlain}}, Form: w.RubricItem, Variant: w.RubricItemLine}, Form: w.MenuHeadingSplit},
		w.MenuItem{Icon: "gear", Label: "Settings"},
		w.MenuItem{Icon: "adm", Label: "Administration"},
		w.MenuSep{},
		w.MenuItem{Icon: "device", Label: "Application settings"},
		w.MenuItem{Icon: "server", Label: "Switch instance"},
		w.MenuSep{},
		w.MenuColumn{Label: "Theme", Control: w.Seg{Options: []w.SegOption{{Label: "Auto", On: true}, {Label: "Light"}, {Label: "Dark"}}, Variant: w.SegStartSmall}},
		w.MenuItem{Icon: "search", Trail: w.Kbd{Keys: "?"}, Label: "Keyboard shortcuts"},
		w.MenuSep{},
		w.MenuItem{Label: "Sign out…"}}, Width: w.MenuWidth290}, "m-desktop-2", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "Only as set in the event"},
		w.MenuSep{},
		w.MenuItem{Label: "At the time · and as set in the event"},
		w.MenuItem{Label: "10 minutes · and as set in the event", On: true},
		w.MenuItem{Label: "30 minutes · and as set in the event"},
		w.MenuItem{Label: "1 hour · and as set in the event"},
		w.MenuSep{},
		w.MenuItem{Label: "Never"}}, Width: w.MenuWidth290}, "m-desktop-20", "tc-menu", 1),

	// tc-drop
	overlay("more", w.Drop{Open: true, Trigger: w.Button{Icon: "dots", Title: "More"}, More: true, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "pencil", Label: "Rename"},
		w.MenuItem{Icon: "down", Detail: &w.Mono{Text: ".vcf", Variant: w.Mono11Dusty}, Label: "Export"}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}}, "components-19", "tc-drop", 1),
	overlay("pick", w.Drop{Open: true, Trigger: w.Pick{Value: "Синк-ап по релизу ковал", Label: "Where"}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "cal", Trail: w.Mono{Text: "10 Aug · Личный", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Звонок с Ковалёвой", On: true},
		w.MenuItem{Icon: "task", Trail: w.Mono{Text: "task · Работа", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Ответить Анне по смете"},
		w.MenuItem{Icon: "con", Trail: w.Mono{Text: "contact", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Анна Ковалёва"},
		w.MenuSep{},
		w.MenuItem{Label: "Enter a UID by hand", Tone: w.MenuToneDusty}}, Width: w.MenuWidthFull, Pop: w.MenuUnder}}, "s-notes-4", "tc-drop", 1),

	// the forms that waited for menus
	overlay("suggest", w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}, Focus: true, Variant: w.SearchFull}, Suggest: &w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Recent searches", Form: w.RubricItem}},
		w.MenuItem{Icon: "clock", Trail: w.Mono{Text: "8 found · 2 min ago", Variant: w.Mono11DustyTight, Place: w.Pushed}, Label: "ковал"},
		w.MenuItem{Icon: "clock", Trail: w.Mono{Text: "3 found · yesterday", Variant: w.Mono11DustyTight, Place: w.Pushed}, Label: "смета"},
		w.MenuItem{Icon: "clock", Trail: w.Mono{Text: "nothing found · Monday", Variant: w.Mono11DustyTight, Place: w.Pushed}, Label: "дача"},
		w.MenuSep{},
		w.MenuItem{Trail: w.Kbd{Keys: "Enter"}, Label: "Type to search everything", Tone: w.MenuToneDusty}}, Width: w.MenuWidthFull}, Actions: []w.TopAction{w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary}, w.Button{Label: "2 min ago", Icon: "refresh"}, w.UserMenu{Name: "mix"}}}, "s-find-2", "tc-top", 1),
	overlay("drop", w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}}, Actions: []w.TopAction{w.TopDrop{Open: true, Trigger: w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary, Focus: true}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "quill", Trail: w.Kbd{Keys: "N"}, Label: "Note", On: true},
		w.MenuSep{},
		w.MenuItem{Icon: "con", Trail: w.Kbd{Keys: "C"}, Label: "Contact"},
		w.MenuItem{Icon: "cal", Trail: w.Kbd{Keys: "E"}, Label: "Event"},
		w.MenuItem{Icon: "task", Trail: w.Kbd{Keys: "T"}, Label: "Task"},
		w.MenuSep{},
		w.MenuItem{Icon: "up", Label: "Upload file"},
		w.MenuSep{},
		w.MenuItem{Icon: "mail", Label: "Letter"}}}}, w.Button{Label: "2 min ago", Icon: "refresh"}, w.UserMenu{Name: "mix"}}}, "s-notes-7", "tc-top", 1),
	overlay("drop", w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}}, Actions: []w.TopAction{w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary}, w.Button{Label: "2 min ago", Icon: "refresh"}, w.TopDrop{Open: true, Trigger: w.UserMenu{Name: "mix", Focus: true}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "mix", Items: []w.RubricPart{w.Mono{Text: "carrel.example", Variant: w.MonoTightPlain}}, Form: w.RubricItem, Variant: w.RubricItemLine}, Form: w.MenuHeadingSplit},
		w.MenuItem{Icon: "gear", Label: "Settings"},
		w.MenuItem{Icon: "adm", Label: "Administration"},
		w.MenuSep{},
		w.MenuColumn{Label: "Theme", Control: w.Seg{Options: []w.SegOption{{Label: "Auto", On: true}, {Label: "Light"}, {Label: "Dark"}}, Variant: w.SegStartSmall}},
		w.MenuItem{Icon: "search", Trail: w.Kbd{Keys: "?"}, Label: "Keyboard shortcuts"},
		w.MenuSep{},
		w.MenuItem{Label: "Sign out"}}, Width: w.MenuWidth290}, End: true}}}, "account-5", "tc-top", 1),
	overlay("more", w.Bar{Items: []w.BarPart{w.Seg{Options: []w.SegOption{{Label: "A–Z", On: true}, {Label: "Recently changed"}}}, w.Seg{Options: []w.SegOption{{Label: "All", On: true}, {Label: "Overdue"}}, Secondary: true}, w.Button{Label: "Columns", Icon: "cols"}, w.Drop{Open: true, Trigger: w.Button{Icon: "dots", Title: "More"}, More: true, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "All", On: true},
		w.MenuItem{Label: "Overdue"}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}}}}, "components-19", "tc-bar", 1),
	overlay("more", w.Bar{Items: []w.BarPart{w.Seg{Options: []w.SegOption{{Label: "Day"}, {Label: "Week", On: true}, {Label: "Month"}, {Label: "Range"}}}, w.Range{From: w.Input{Kind: w.InputDate, Value: "2026-08-10", Mono: true, Label: "Date"}, To: w.Input{Kind: w.InputDate, Value: "2026-08-16", Mono: true, Label: "Date"}, Show: w.Button{Label: "Show"}, Secondary: true}, w.BarSep{Secondary: true}, w.Check{On: true, Parts: []w.CheckPart{w.Text("Events")}, Tertiary: true}, w.Check{Parts: []w.CheckPart{w.Text("Tasks due")}, Tertiary: true}, w.Check{Parts: []w.CheckPart{w.Text("Notes")}, Tertiary: true}, w.Drop{Open: true, Trigger: w.Button{Icon: "dots", Title: "More"}, More: true, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Tick: &w.TickBox{Ticked: w.TickedOn}, Label: "Events", Tertiary: true},
		w.MenuItem{Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Tasks due", Tertiary: true},
		w.MenuItem{Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Notes", Tertiary: true},
		w.MenuSep{Tertiary: true},
		w.Range{From: w.Input{Kind: w.InputDate, Value: "2026-08-10", Mono: true, Label: "Date"}, To: w.Input{Kind: w.InputDate, Value: "2026-08-16", Mono: true, Label: "Date"}, Show: w.Button{Label: "Show"}, Wrap: true},
		w.MenuSep{},
		w.MenuItem{Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Compact rows"}}, Pop: w.MenuUnderEnd, Inline: true}}}, Right: &w.BarRight{Items: []w.BarPart{w.Check{Parts: []w.CheckPart{w.Text("Compact rows")}}}, Secondary: true}, Variant: w.BarWide}, "s-agenda-1", "tc-bar", 1),

	// tc-sheet
	overlay("icon", w.Sheet{Icon: "move", Title: "Move contact", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Анна Ковалёва · from nextcloud · Контакты")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Rubric{Text: "Write into"},
		w.Tree{Rows: []w.TreeRow{
			{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Unavailable: true, Trail: w.Mono{Text: "already there", Variant: w.Mono10Dusty}},
			{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "Архив", Trail: w.Mono{Text: "nextcloud", Variant: w.Mono10Dusty}},
			{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Trail: w.Mono{Text: "radicale", Variant: w.Mono10Dusty}},
			{Mark: &w.Mark{Colour: "#8C6239"}, Label: "Shared", Unavailable: true, Trail: w.Badge{Label: "read-only"}}}, On: 2},
		w.Msg{Parts: []w.MsgPart{w.Text("Inside one server this is "), w.Mono{Text: "MOVE", Form: w.MonoCode}, w.Text(": instant, no traffic. To another address book the card is copied and then deleted — Carrel says so before it starts, not after. Copy is this same sheet with "), w.Strong("Copy here"), w.Text(".")}, Space: w.MsgFlushTop2}}, Foot: []w.SheetFootPart{w.Button{Label: "Move here", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Enter"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth520}, "s-contacts-2", "tc-sheet", 1),
	overlay("plain", w.Sheet{Title: "Publish a calendar", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("nextcloud · Личный")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Field{Label: "Link", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}, Foot: []w.SheetFootPart{w.Button{Label: "Send by letter", Kind: w.ButtonPrimary}, w.Button{Label: "Open the link"}, w.Button{Label: "All published links"}}}, "s-publish-13", "tc-sheet", 1),
	overlay("close", w.Sheet{Icon: "quill", Title: "New note", Head: []w.SheetHeadPart{w.Pick{Value: "nextcloud · Записи", Variant: w.PickT12W200, Label: "Where", Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.Area{Text: "Позвонить в МФЦ во вторник до 12:00 — уточнить, нужен ли оригинал доверенности.", Variant: w.AreaBodySm, Focus: true, Label: "Text"}}, Foot: []w.SheetFootPart{w.Button{Label: "Save note", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Ctrl+Enter"}, w.Button{Label: "Open the card", Icon: "expand"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}}, "s-notes-5", "tc-sheet", 1),
	overlay("close", w.Sheet{Icon: "con", Title: "New contact", Head: []w.SheetHeadPart{w.Pick{Value: "nextcloud · Контакты", Variant: w.PickT12W200, Label: "Where", Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.SheetFields{Items: []w.Field{w.Field{Label: "Display name", Place: w.Grown, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Phone", Width: w.FieldFill200, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Email", Width: w.FieldFill200, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}},
		w.Dim{Parts: []w.Inline{w.Text("The first word is Given, the rest Family — «Пётр Ковалёв», not «Ковалёв Пётр». Everything else is filled in on the card.")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Foot: []w.SheetFootPart{w.Button{Label: "Save contact", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Enter"}, w.Button{Label: "Open the card", Icon: "expand"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}}, "s-notes-8", "tc-sheet", 1),
	overlay("close", w.Sheet{Icon: "quill", Title: "New note", Head: []w.SheetHeadPart{w.Close{Label: "Close", Place: w.Pushed}}, Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("There is nowhere to write. Connect an account that has a journal, or create a notebook after connecting — the note is not kept in this window.")}, Space: w.MsgFlush}}, Foot: []w.SheetFootPart{w.Button{Label: "Connect account", Icon: "server", Kind: w.ButtonPrimary}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth480}, "s-notes-6", "tc-sheet", 1),
	overlay("close", w.Sheet{Icon: "task", Title: "New task", Head: []w.SheetHeadPart{w.Pick{Value: "nextcloud · Дела", Variant: w.PickT12W200, Label: "Where", Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.SheetFields{Items: []w.Field{w.Field{Label: "Summary", Place: w.Grown, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Due date", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Due time", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}},
		w.Dim{Parts: []w.Inline{w.Text("A task with no due date is a task all the same — the field may stay empty.")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Foot: []w.SheetFootPart{w.Button{Label: "Save task", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Enter"}, w.Button{Label: "Open the card", Icon: "expand"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth520}, "s-notes-10", "tc-sheet", 1),
	overlay("icon", w.Sheet{Icon: "shield", Tone: w.SheetIconInk, Title: "Deposit my key", Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("An administrator who holds the escrow master password will be able to decrypt the DAV credentials you save here, and through them your contacts, calendars, tasks, notes and files. Ending sessions does not undo this. You can withdraw the copy later from Account.")}, Tone: w.MsgAlert, Space: w.MsgFlush},
		w.Field{Label: "Your Carrel password", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}, Foot: []w.SheetFootPart{w.Button{Label: "Deposit my key", Kind: w.ButtonPrimary}, w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc", Place: w.Pushed}}, Width: w.SheetWidth480}, "account-2", "tc-sheet", 1),
	overlay("close", w.Sheet{Icon: "search", Title: "Keyboard shortcuts", Head: []w.SheetHeadPart{w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.SheetKeys{Items: []w.SheetKey{{Keys: "N", Does: "Quick note over this screen"}, {Keys: "/", Does: "Focus the search field in the header"}, {Keys: "Enter", Does: "Open the selected record full screen"}, {Keys: "Esc", Does: "Close a sheet, the details panel or focus mode"}, {Keys: "Ctrl+S", Does: "Save in the editor"}, {Keys: "Ctrl+Enter", Does: "Save a quick note"}, {Keys: "Ctrl+V", Does: "Paste an image as an attachment"}}},
		w.SheetHint{Text: "Shortcuts do nothing while the cursor is in a field. None of them destroy anything."}}, Foot: []w.SheetFootPart{w.Button{Label: "Close"}, w.Kbd{Keys: "Esc"}, w.Kbd{Keys: "?", Place: w.Pushed}}, Width: w.SheetWidth480}, "account-6", "tc-sheet", 1),
	overlay("icon", w.Sheet{Icon: "shield", Title: "Show password", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("nextcloud")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Field{Label: "Your Carrel password", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Dim{Parts: []w.Inline{w.Text("Unlocks credentials everywhere for five minutes: passwords, device profiles, key escrow — anything handed out or handed over.")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Foot: []w.SheetFootPart{w.Button{Label: "Show", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Enter"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth400}, "s-profile-3", "tc-sheet", 1),
	overlay("icon", w.Sheet{Icon: "plus", Title: "New calendar", Body: []w.SheetPart{
		w.Field{Label: "On which server", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Field{Label: "Name", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Field{Label: "Colour", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Field{Label: "What goes in it", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Field{Label: "Address", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}, Foot: []w.SheetFootPart{w.Button{Label: "Create", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Enter"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}}, "s-profile-5", "tc-sheet", 1),
	overlay("danger", w.Sheet{Icon: "plus", Tone: w.SheetIconDanger, Title: "New calendar", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("baikal")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("The server would not make it. Nothing was created, and the form is as you left it.")}, Tone: w.MsgAlert, Space: w.MsgFlush},
		w.Diag{Text: "[create] MKCALENDAR /dav.php/calendars/m.yenuchenko/poezdki/ → 403 <D:error><D:need-privileges/></D:error> no DAV:bind on the calendar home-set"},
		w.SheetNote{Text: "This account may sign in and read, but not create. Either the user has no such right on the server, or the account was connected with a password issued read-only. Neither is something Carrel can fix from here — but it can say which of the two it is."}}, Foot: []w.SheetFootPart{w.Button{Label: "Back to the form", Kind: w.ButtonPrimary}, w.Button{Label: "Try again"}}}, "s-profile-6", "tc-sheet", 1),
	overlay("danger", w.Sheet{Icon: "trash", Tone: w.SheetIconDanger, Title: "Delete on the server", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("nextcloud · Дни рождения")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("The calendar and all "), w.Strong("51"), w.Text(" events in it are deleted on "), w.Mono{Text: "cloud.example.org"}, w.Text(". Carrel holds no copy — it reads and writes, it does not store — so there is nothing here to undo it with.")}, Tone: w.MsgAlert, Space: w.MsgFlush},
		w.SheetBlock{Title: w.Rubric{Text: "What is pointing at it now"}, Items: []w.SheetLine{{Name: "1 published link", Note: w.Dim{Parts: []w.Inline{w.Text("— answers “not found” from that moment")}}}, {Name: "Backup job «Личное на домашний диск»", Note: w.Dim{Parts: []w.Inline{w.Text("— copies this collection; it will run one short and say so")}}}, {Name: "3 devices through davloom", Note: w.Dim{Parts: []w.Inline{w.Text("— it disappears for them at the next sync")}}}}},
		w.Msg{Parts: []w.MsgPart{w.Text("Last time it was in an archive: "), w.Mono{Text: "2026-08-14"}, w.Text(". Restoring brings the records back into a "), w.Strong("new"), w.Text(" collection — not into this one, and not to this address.")}, Space: w.MsgFlush},
		w.Line{Block: true, Gap: w.LineGap10, Items: []w.LinePart{w.Field{Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}},
		w.Field{Label: "Type the name to confirm", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}, Foot: []w.SheetFootPart{w.Button{Label: "Delete on the server", Kind: w.ButtonDangerSolid}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}}, "s-profile-9", "tc-sheet", 1),
	overlay("icon", w.Sheet{Icon: "refresh", Title: "Rewrite the address in her records?", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Анна Ковалёва · saved")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.SheetBlock{Title: w.Rubric{Text: "The address", Variant: w.RubricBottom7}, Items: []w.SheetLine{{Name: "a.kovaleva@example.org → a.kovaleva@slovo.example", Note: w.Dim{Parts: []w.Inline{w.Text("— 14 events and tasks name the old one; she drops out of their timeline")}}}}},
		w.SheetBlock{Title: w.Rubric{Text: "What rewriting can set off", Variant: w.RubricBottom7}, Items: []w.SheetLine{
			{Name: "9 records on servers that schedule", Note: w.Dim{Parts: []w.Inline{w.Text("— the server may send a cancellation to the old address, an invitation to the new one and an update to the other attendees")}}},
			{Name: "2 records on baikal", Note: w.Dim{Parts: []w.Inline{w.Text("— the server did not say whether it schedules")}}},
			{Name: "3 meetings organised by others", Note: w.Dim{Parts: []w.Inline{w.Text("— the organiser’s server may refuse the change, and its next update may undo it")}}},
			{Name: "Past meetings", Note: w.Dim{Parts: []w.Inline{w.Text("— will show the new address, as if she had been invited at it")}}}}},
		w.SheetBlock{Title: w.Rubric{Text: "What will not change", Variant: w.RubricBottom7}, Items: []w.SheetLine{{Name: "2 records in Shared", Note: w.Dim{Parts: []w.Inline{w.Text("— read-only; they keep the old address")}}}}}},
		Foot:  []w.SheetFootPart{w.Button{Label: "Rewrite 12 records", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Enter"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Leave as is"}, w.Kbd{Keys: "Esc"}}}},
		Width: w.SheetWidth520}, "s-contacts-15", "tc-sheet", 1),
	overlay("plain", w.Sheet{Title: "Show password", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("nextcloud")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Field{Label: "Your Carrel password", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Dim{Parts: []w.Inline{w.Text("Unlocks credentials everywhere for five minutes: passwords, device profiles, anything else handed out.")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Foot: []w.SheetFootPart{w.Button{Label: "Show", Kind: w.ButtonPrimary}, w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc", Place: w.Pushed}}, Width: w.SheetWidth400}, "s-plan-3", "tc-sheet", 1),
	overlay("plain", w.Sheet{Title: "Publish a calendar", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("nextcloud · Личный")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Field{Label: "How much of it", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Line{Block: true, Gap: w.LineGap10, Items: []w.LinePart{w.Field{Label: "From", Width: w.FieldSM, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "To", Width: w.FieldSM, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}},
		w.Field{Label: "Link lives for", Width: w.FieldSM, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.SheetChecks{Items: []w.Check{w.Check{On: true, Parts: []w.CheckPart{w.Text("Allow"), w.Mono{Text: ".ics"}, w.Text("download and subscription")}}, w.Check{Parts: []w.CheckPart{w.Text("Show event descriptions")}}}}}, Foot: []w.SheetFootPart{w.Button{Label: "Create link", Kind: w.ButtonPrimary}, w.Button{Label: "Cancel"}}}, "s-publish-2", "tc-sheet", 1),
	overlay("close", w.Sheet{Icon: "mail", Title: "Letter", Head: []w.SheetHeadPart{w.Tag{Label: "from mix@carrel.example", On: true, Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.Field{Label: "To", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Field{Label: "Subject", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Area{Text: "Анна, Артём — вот карточка подрядчика и смета, которую он прислал. Договор в папке, ссылка ниже.", Variant: w.AreaBodySm, Focus: true, Label: "Text"},
		w.Sec{Items: []w.SecPart{
			w.Rubric{Text: "Attached", Items: []w.RubricPart{w.Push{Items: []w.PushPart{w.Text("Attach")}}}, Variant: w.RubricItemLine},
			w.SecFiles{Items: []w.SecFile{{Icon: "con", Name: "Пётр Смирнов · подрядчик", Size: w.Mono{Text: ".vcf · 3 KB"}, Tag: w.Tag{Label: "attached"}, Remove: w.Button{Icon: "close", Small: true, Title: "More"}}, {Icon: "cal", Name: "Встреча на объекте · 18 August 10:00", Size: w.Mono{Text: ".ics · 2 KB"}, Tag: w.Tag{Label: "attached"}, Remove: w.Button{Icon: "close", Small: true, Title: "More"}}, {Icon: "doc", Name: "смета-подрядчик.pdf", Size: w.Mono{Text: "4.2 MB"}, Tag: w.Tag{Label: "as a link", On: true}, Remove: w.Button{Icon: "close", Small: true, Title: "More"}}, {Icon: "file", Name: "Договор подряда · 9 файлов", Size: w.Mono{Text: "folder"}, Tag: w.Tag{Label: "as a link", On: true}, Remove: w.Button{Icon: "close", Small: true, Title: "More"}}}},
			w.Msg{Parts: []w.MsgPart{w.Text("Two of these go as links, not as files: the PDF is over the 2 MB ceiling, and a folder always goes as a link. Both links live 30 days and can be revoked in "), w.Strong("Published links"), w.Text(" — including for the people in this letter.")}, Tone: w.MsgAccent, Space: w.MsgFlushTop11}}, Flush: true}}, Foot: []w.SheetFootPart{w.Button{Label: "Send", Icon: "mail", Kind: w.ButtonPrimary}, w.Button{Label: "Preview"}, w.Dim{Parts: []w.Inline{w.Text("Signed Михаил Енученко · replies go to mix@example.org")}, Size: w.Dim11_5, Variant: w.DimWrap}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth680}, "s-mail-1", "tc-sheet", 1),
	overlay("close", w.Sheet{Icon: "mail", Title: "Send this link", Head: []w.SheetHeadPart{w.SheetSubject{Mark: w.Mark{Colour: "#4A6B52"}, Text: "Личный · 10–31 August"}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.Field{Label: "To", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Field{Label: "A line from you", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.SheetFrom{Label: "Sent as", Lines: [][]w.Inline{[]w.Inline{w.Strong("Михаил Енученко"), w.Text(" "), w.Mono{Text: "<mix@carrel.example>"}}, []w.Inline{w.Text("Replies go to "), w.Mono{Text: "mix@example.org"}, w.Text(" · "), w.Span{Face: w.FaceAccent, Text: "change in Account"}}}},
		w.Sec{Items: []w.SecPart{
			w.Rubric{Text: "The letter carries"},
			w.SecList{Items: [][]w.SecItemPart{[]w.SecItemPart{w.Text("The link and your line — nothing of the calendar itself")}, []w.SecItemPart{w.Text("Revoking the link closes it for this recipient too")}, []w.SecItemPart{w.Text("It expires on 8 September, as the link does")}}, Gap: 5, Muted: true}}, Flush: true}}, Foot: []w.SheetFootPart{w.Button{Label: "Send", Icon: "mail", Kind: w.ButtonPrimary}, w.Button{Label: "Copy link instead"}, w.Push{Items: []w.PushPart{w.Button{Label: "Cancel"}}}}, Width: w.SheetWidthFull}, "s-mail-4", "tc-sheet", 1),
	overlay("pick", w.Sheet{Icon: "note", Title: "Attach a note", Head: []w.SheetHeadPart{w.Filter{Query: w.Query{Value: "подряд", Label: "подряд"}, Quiet: true, Variant: w.FilterT12W200, Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.SheetPick{Rail: []w.SheetPickRail{w.RailSec{Items: []w.RailSecPart{
			w.TickRow{Ticked: w.TickedPart, Kind: w.TickRowAll, Mark: w.Mark{Kind: w.MarkMulti}, Label: "Combined notebook", Meta: "86", Current: true}}}, w.RailSec{Items: []w.RailSecPart{
			w.Rubric{Text: "nextcloud"},
			w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "61"},
			w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Рабочее", Meta: "25"}}}}, List: []w.SheetPickList{w.ListGroup{Title: w.Rubric{Text: "Найдено"}, Count: &w.Num{Value: 3}}, w.List{Rows: []w.ListRow{
			w.BriefRow{Mark: w.Mark{Colour: "#3F6E8C"}, Title: "Встреча с подрядчиком, итоги", Date: w.Mono{Text: "18 Aug"}, Where: "nextcloud · Рабочее", On: true},
			w.BriefRow{Mark: w.Mark{Colour: "#3F6E8C"}, Title: "Подрядчики: кого смотрели", Date: w.Mono{Text: "02 Aug"}, Where: "nextcloud · Рабочее"},
			w.BriefRow{Mark: w.Mark{Colour: "#4A6B52"}, Title: "Ремонт: что осталось", Date: w.Mono{Text: "27 Jul"}, Where: "nextcloud · Записи"}}}}},
		w.Sec{Items: []w.SecPart{
			w.Field{Label: "Send it as", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}}}, Flush: true, Foot: []w.SheetFootPart{w.Button{Label: "Attach", Kind: w.ButtonPrimary}, w.Dim{Parts: []w.Inline{w.Text("1 note · .md · 4 KB · goes as a file")}, Size: w.Dim11_5, Variant: w.DimWrap}, w.Push{Items: []w.PushPart{w.Button{Label: "Cancel"}}}}, Width: w.SheetWidthFull}, "s-mail-5", "tc-sheet", 1),
	overlay("pick", w.Sheet{Icon: "cal", Title: "Attach a calendar", Head: []w.SheetHeadPart{w.Close{Label: "Close", Place: w.Pushed}}, Body: []w.SheetPart{
		w.SheetPick{Rail: []w.SheetPickRail{w.Tree{Rows: []w.TreeRow{
			{Icon: "server", Label: "nextcloud"},
			{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "Личный", Indent: w.TreeLevel22},
			{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Indent: w.TreeLevel22},
			{Icon: "server", Label: "radicale · vps"},
			{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Indent: w.TreeLevel22}}, On: 2, Variant: w.TreePadded}}, List: []w.SheetPickList{w.Form{}}, Short: true}}, Flush: true, Foot: []w.SheetFootPart{w.Button{Label: "Attach", Kind: w.ButtonPrimary}, w.Dim{Parts: []w.Inline{w.Text("Calendar «Личный» · 10 August 2026 – 31 August 2026 · .ics · 5 KB · goes as a file")}, Size: w.Dim11_5, Variant: w.DimWrap}, w.Push{Items: []w.PushPart{w.Button{Label: "Cancel"}}}}, Width: w.SheetWidthFull}, "s-mail-10", "tc-sheet", 1),
	overlay("pick", w.Sheet{Icon: "doc", Title: "Attach a file", Head: []w.SheetHeadPart{w.Crumbs{Path: []w.Crumb{{Label: "Documents", Href: "#"}, {Label: "Стройка", Href: "#"}}, Current: "Подряд"}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.SheetPick{Rail: []w.SheetPickRail{w.Tree{Rows: []w.TreeRow{
			{Icon: "server", Label: "sftpgo · vps"},
			{Icon: "file", Label: "Documents", Indent: w.TreeLevel22},
			{Icon: "file", Label: "Стройка", Indent: w.TreeLevel36},
			{Icon: "file", Label: "Подряд", Indent: w.TreeLevel44},
			{Icon: "file", Label: "Личное", Indent: w.TreeLevel36}}, On: 4, Variant: w.TreePadded}}, List: []w.SheetPickList{w.Table{Cols: []w.TableCol{{Label: "Name"}, {Label: "Size", Width: w.Col90}, {Label: "Changed", Width: w.Col110, Rank: w.Rank3}, {Width: w.ColType, Rank: w.Rank4}}, Rows: []w.TableRow{
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "смета-подрядчик.pdf"}}}, {Parts: []w.CellPart{w.Text("4.2 MB")}, Mono: true}, {Parts: []w.CellPart{w.Text("12 Aug")}, Mono: true}, {Parts: []w.CellPart{w.Tag{Label: "over 2 MB · as a link", On: true}}}}, State: w.RowOn},
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "договор-подряда.pdf"}}}, {Parts: []w.CellPart{w.Text("840 KB")}, Mono: true}, {Parts: []w.CellPart{w.Text("09 Aug")}, Mono: true}, {Parts: []w.CellPart{w.Tag{Label: "as a file"}}}}},
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.TableFile{Icon: "img", Name: "объект-фасад.jpg"}}}, {Parts: []w.CellPart{w.Text("1.6 MB")}, Mono: true}, {Parts: []w.CellPart{w.Text("09 Aug")}, Mono: true}, {Parts: []w.CellPart{w.Tag{Label: "as a file"}}}}}}}}, Short: true},
		w.Msg{Parts: []w.MsgPart{w.Text("This file is over the limit for an attachment, so the letter will carry a secret link to it instead. The link is listed in "), w.Strong("Published links"), w.Text(", works for 30 days, and revoking it there closes it for whoever got the letter.")}, Tone: w.MsgAccent, Space: w.MsgSheet}}, Flush: true, Foot: []w.SheetFootPart{w.Button{Label: "Attach", Kind: w.ButtonPrimary}, w.Dim{Parts: []w.Inline{w.Text("смета-подрядчик.pdf · 4.2 MB · goes as a link")}, Size: w.Dim11_5, Variant: w.DimWrap}, w.Push{Items: []w.PushPart{w.Button{Label: "Cancel"}}}}, Width: w.SheetWidthFull}, "s-mail-6", "tc-sheet", 1),
	overlay("close", w.Sheet{Title: "Sources", Head: []w.SheetHeadPart{w.Button{Label: "Apply", Small: true, Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.RailSec{Items: []w.RailSecPart{
			w.TickRow{Ticked: w.TickedPart, Kind: w.TickRowAll, Mark: w.Mark{Kind: w.MarkMulti}, Label: "Combined notebook", Meta: "86", Current: true}}},
		w.RailSec{Items: []w.RailSecPart{
			w.Rubric{Text: "nextcloud"},
			w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "64"},
			w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "22"},
			w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "—", Hidden: true}}}}}, "m-phone-9", "tc-sheet", 1),

	// tc-scrim
	overlay("pad", w.Scrim{Content: w.Sheet{Icon: "move", Title: "Move contact", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Анна Ковалёва · from nextcloud · Контакты")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Rubric{Text: "Write into"},
		w.Tree{Rows: []w.TreeRow{
			{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Unavailable: true, Trail: w.Mono{Text: "already there", Variant: w.Mono10Dusty}},
			{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "Архив", Trail: w.Mono{Text: "nextcloud", Variant: w.Mono10Dusty}},
			{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Trail: w.Mono{Text: "radicale", Variant: w.Mono10Dusty}},
			{Mark: &w.Mark{Colour: "#8C6239"}, Label: "Shared", Unavailable: true, Trail: w.Badge{Label: "read-only"}}}, On: 2},
		w.Msg{Parts: []w.MsgPart{w.Text("Inside one server this is "), w.Mono{Text: "MOVE", Form: w.MonoCode}, w.Text(": instant, no traffic. To another address book the card is copied and then deleted — Carrel says so before it starts, not after. Copy is this same sheet with "), w.Strong("Copy here"), w.Text(".")}, Space: w.MsgFlushTop2}}, Foot: []w.SheetFootPart{w.Button{Label: "Move here", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Enter"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth520}, Pad: w.ScrimPad26}, "s-contacts-2", "tc-scrim", 1),
	overlay("plain", w.Scrim{Content: w.Sheet{Icon: "quill", Title: "New note", Head: []w.SheetHeadPart{w.Pick{Value: "nextcloud · Записи", Variant: w.PickT12W200, Label: "Where", Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.Area{Text: "Позвонить в МФЦ во вторник до 12:00 — уточнить, нужен ли оригинал доверенности.", Variant: w.AreaBodySm, Focus: true, Label: "Text"}}, Foot: []w.SheetFootPart{w.Button{Label: "Save note", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Ctrl+Enter"}, w.Button{Label: "Open the card", Icon: "expand"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}}}, "s-notes-5", "tc-scrim", 1),
	overlay("pad", w.Scrim{Content: w.Empty{Icon: "up", Lead: "Drop to import", Parts: []w.EmptyPart{w.Aside{Face: w.FaceLarge, Parts: []w.Inline{w.Text("One"), w.Mono{Text: ".vcf"}, w.Text("file into"), w.Strong("nextcloud · Контакты")}}, w.Dim{Parts: []w.Inline{w.Text("The preview opens first: nothing is written until it has been read")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Variant: w.EmptyDrop}, Pad: w.ScrimPad38}, "s-io-1", "tc-scrim", 1),

	// tc-overlay
	overlay("plain", w.Overlay{Body: w.OverlayBody{Rail: []w.OverlayRailPart{w.Nav{Items: []w.NavItem{
		{Icon: "con", Label: "Contacts", Href: "#"},
		{Icon: "cal", Label: "Calendar", Href: "#"},
		{Icon: "task", Label: "Tasks", Href: "#"},
		{Icon: "note", Label: "Notes", Href: "#"},
		{Icon: "file", Label: "Files", Href: "#"},
		{Icon: "dup", Label: "Duplicates", Href: "#"}}, On: 2}, w.RailSec{Items: []w.RailSecPart{
		w.TickRow{Ticked: w.TickedOff, Kind: w.TickRowAll, Mark: w.Mark{Kind: w.MarkMulti}, Label: "Combined calendar", Meta: "—"}}}, w.Empty{Parts: []w.EmptyPart{w.Text("No account connected.")}}, w.RailFoot{Items: []w.RailFootPart{w.Button{Icon: "plus", Title: "New calendar"}, w.Button{Icon: "server", Title: "Connect account"}, w.Button{Icon: "link", Title: "Subscribe to a feed"}, w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}}, About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"}}}, Main: w.Main{Items: []w.MainPart{
		w.Empty{Parts: []w.EmptyPart{w.Text("Connect an account to see calendars, address books and notebooks.")}}}}, Height: w.ShellHeight340}, Scrim: w.Scrim{Content: w.Sheet{Icon: "quill", Title: "New note", Head: []w.SheetHeadPart{w.Close{Label: "Close", Place: w.Pushed}}, Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("There is nowhere to write. Connect an account that has a journal, or create a notebook after connecting — the note is not kept in this window.")}, Space: w.MsgFlush}}, Foot: []w.SheetFootPart{w.Button{Label: "Connect account", Icon: "server", Kind: w.ButtonPrimary}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth480}}}, "s-notes-6", "tc-overlay", 1),
	overlay("plain", w.Overlay{Body: w.OverlayBody{Rail: []w.OverlayRailPart{w.Nav{Items: []w.NavItem{
		{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
		{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
		{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
		{Icon: "file", Label: "Files", Href: "#"}}, On: 2}, w.RailSec{Items: []w.RailSecPart{
		w.TickRow{Ticked: w.TickedPart, Kind: w.TickRowAll, Mark: w.Mark{Kind: w.MarkMulti}, Label: "Combined calendar", Meta: "14", Current: true}}}, w.RailSec{Items: []w.RailSecPart{
		w.Rubric{Text: "nextcloud"},
		w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "7"},
		w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "4"}}}}, Main: w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti}, Text: "Combined calendar"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("5 of 7 calendars · 10 – 16 August 2026")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}}}},
		w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August"}, Count: &w.Num{Value: 3}},
		w.List{Rows: []w.ListRow{
			w.TimeRow{Mark: w.Mark{Colour: "#4A6B52"}, Time: "10:00", Title: []w.Inline{w.Text("Синк-ап по релизу")}, Detail: "Переговорная 2", Source: "nextcloud · Личный"},
			w.TimeRow{Mark: w.Mark{Colour: "#8C6239"}, Time: "all day", Title: []w.Inline{w.Text("Дежурство: первая линия")}, Detail: "—", Source: "2 sources"},
			w.TimeRow{Mark: w.Mark{Colour: "#3F6E8C"}, Time: "15:00", Title: []w.Inline{w.Text("Созвон с подрядчиком")}, Detail: "Jitsi", Source: "nextcloud · Работа"}}},
		w.ListGroup{Title: w.Rubric{Text: "Tuesday · 11 August"}, Count: &w.Num{Value: 2}},
		w.List{Rows: []w.ListRow{
			w.TimeRow{Mark: w.Mark{Colour: "#6B4E7A"}, Time: "09:30", Title: []w.Inline{w.Text("Ревью архитектуры")}, Detail: "Переговорная 1", Source: "radicale · vps · Проекты"}}}}}}, Scrim: w.Scrim{Content: w.Sheet{Icon: "quill", Title: "New note", Head: []w.SheetHeadPart{w.Pick{Value: "nextcloud · Записи", Variant: w.PickT12W200, Label: "Where", Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.Area{Text: "Позвонить в МФЦ во вторник до 12:00 — уточнить, нужен ли оригинал доверенности.", Variant: w.AreaBodySm, Focus: true, Label: "Text"}}, Foot: []w.SheetFootPart{w.Button{Label: "Save note", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Ctrl+Enter"}, w.Button{Label: "Open the card", Icon: "expand"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}}}}, "s-notes-5", "tc-overlay", 1),
	overlay("plain", w.Overlay{Body: w.OverlayBody{Rail: []w.OverlayRailPart{w.Nav{Items: []w.NavItem{
		{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
		{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
		{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
		{Icon: "file", Label: "Files", Href: "#"},
		{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"}}, On: 1}, w.RailSec{Items: []w.RailSecPart{
		w.TickRow{Ticked: w.TickedPart, Kind: w.TickRowAll, Mark: w.Mark{Kind: w.MarkMulti}, Label: "Combined address book", Meta: "312", Current: true}}}, w.RailSec{Items: []w.RailSecPart{
		w.Rubric{Text: "nextcloud"},
		w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "248"},
		w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "51"}}}}, Main: w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti}, Text: "Combined address book"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 312 cards")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "New contact", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}, w.Button{Label: "Import", Icon: "up"}, w.Button{Label: "Export", Icon: "down"}, w.Button{Label: "Print", Icon: "print"}}}},
		w.Bar{Items: []w.BarPart{w.Filter{Query: w.Query{Value: "Filter the loaded cards", Label: "Filter the loaded cards"}, Waiting: true}, w.Seg{Options: []w.SegOption{{Label: "A–Z", On: true}, {Label: "Recently changed"}}, Secondary: true}, w.Button{Label: "Columns", Icon: "cols"}}, More: &w.Button{Icon: "dots", Title: "More"}, Right: &w.BarRight{Items: []w.BarPart{w.Check{Parts: []w.CheckPart{w.Text("Compact rows")}}, w.Mono{Text: "312"}}, Secondary: true}},
		w.ListGroup{Title: w.Rubric{Text: "А"}, Count: &w.Num{Value: 3}},
		w.List{Rows: []w.ListRow{
			w.FaceRow{Mark: w.Mark{Colour: "#4A6B52"}, Avatar: w.Avatar{Initials: "АК"}, Title: []w.Inline{w.Text("Анна Ковалёва")}, Strong: true, Subtitle: "+7 912 445-10-88", Detail: "a.kovaleva@example.org", Caption: "Издательство «Слово»", Source: "nextcloud · Контакты", On: true},
			w.FaceRow{Mark: w.Mark{Colour: "#6B4E7A"}, Avatar: w.Avatar{Photo: true}, Title: []w.Inline{w.Text("Артём Белов")}, TitleBadge: &w.Badge{Label: "linked ×2", Tone: w.BadgeLinked}, Subtitle: "+7 903 118-22-04", Detail: "a.belov@example.org", Caption: "Тензор", Source: "radicale · vps · Команда"},
			w.FaceRow{Mark: w.Mark{Colour: "#4A6B52"}, Avatar: w.Avatar{Initials: "АС"}, Title: []w.Inline{w.Text("Алексей Смирнов")}, Subtitle: "+7 916 200-71-15", Detail: "—", Caption: "—", Source: "nextcloud · Контакты"}}},
		w.ListGroup{Title: w.Rubric{Text: "В"}, Count: &w.Num{Value: 2}},
		w.List{Rows: []w.ListRow{
			w.FaceRow{Mark: w.Mark{Colour: "#4A6B52"}, Avatar: w.Avatar{Initials: "ВГ"}, Title: []w.Inline{w.Text("Виктор Гончаров")}, Subtitle: "—", Detail: "v.goncharov@example.org", Caption: "Городская библиотека №4", Source: "nextcloud · Контакты"},
			w.FaceRow{Mark: w.Mark{Colour: "#A05A2C"}, Avatar: w.Avatar{Initials: "ВЛ"}, Title: []w.Inline{w.Text("Вера Лебедева")}, Subtitle: "+7 905 331-09-42", Detail: "vera@example.net", Caption: "—", Source: "nextcloud · Архив"}}},
		w.Empty{Parts: []w.EmptyPart{w.Text("Loading more…")}, Variant: w.EmptyDusty}}}, Height: w.ShellHeight560}, Scrim: w.Scrim{Content: w.Sheet{Icon: "move", Title: "Move contact", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Анна Ковалёва · from nextcloud · Контакты")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Rubric{Text: "Write into"},
		w.Tree{Rows: []w.TreeRow{
			{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Unavailable: true, Trail: w.Mono{Text: "already there", Variant: w.Mono10Dusty}},
			{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "Архив", Trail: w.Mono{Text: "nextcloud", Variant: w.Mono10Dusty}},
			{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Trail: w.Mono{Text: "radicale", Variant: w.Mono10Dusty}},
			{Mark: &w.Mark{Colour: "#8C6239"}, Label: "Shared", Unavailable: true, Trail: w.Badge{Label: "read-only"}}}, On: 2},
		w.Msg{Parts: []w.MsgPart{w.Text("Inside one server this is "), w.Mono{Text: "MOVE", Form: w.MonoCode}, w.Text(": instant, no traffic. To another address book the card is copied and then deleted — Carrel says so before it starts, not after. Copy is this same sheet with "), w.Strong("Copy here"), w.Text(".")}, Space: w.MsgFlushTop2}}, Foot: []w.SheetFootPart{w.Button{Label: "Move here", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Enter"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth520}, Pad: w.ScrimPad26}}, "s-contacts-2", "tc-overlay", 1),
	overlay("top", w.Overlay{Top: &w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}}, Actions: []w.TopAction{w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary}, w.Button{Label: "2 min ago", Icon: "refresh"}, w.UserMenu{Name: "mix"}}}, Body: w.OverlayBody{Rail: []w.OverlayRailPart{w.Nav{Items: []w.NavItem{
		{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
		{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
		{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
		{Icon: "file", Label: "Files", Href: "#"},
		{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"}}, On: 2}, w.RailSec{Items: []w.RailSecPart{
		w.TickRow{Ticked: w.TickedPart, Kind: w.TickRowAll, Mark: w.Mark{Kind: w.MarkMulti}, Label: "Combined calendar", Meta: "14", Current: true}}}, w.RailSec{Items: []w.RailSecPart{
		w.Rubric{Text: "nextcloud"},
		w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "7"},
		w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "4"},
		w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true}}}}, Main: w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti}, Text: "Combined calendar"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("5 of 7 calendars · 10 – 16 August 2026")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}}}},
		w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August"}, Count: &w.Num{Value: 3}},
		w.List{Rows: []w.ListRow{
			w.TimeRow{Mark: w.Mark{Colour: "#4A6B52"}, Time: "10:00", Title: []w.Inline{w.Text("Синк-ап по релизу")}, Detail: "Переговорная 2", Source: "nextcloud · Личный"},
			w.TimeRow{Mark: w.Mark{Colour: "#3F6E8C"}, Time: "13:30", Title: []w.Inline{w.Text("Звонок с Ковалёвой")}, Detail: "Zoom", Source: "nextcloud · Работа"}}}}}, Height: w.ShellHeight340}, Scrim: w.Scrim{Content: w.Sheet{Icon: "search", Title: "Keyboard shortcuts", Head: []w.SheetHeadPart{w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.SheetKeys{Items: []w.SheetKey{{Keys: "N", Does: "Quick note over this screen"}, {Keys: "/", Does: "Focus the search field in the header"}, {Keys: "Enter", Does: "Open the selected record full screen"}, {Keys: "Esc", Does: "Close a sheet, the details panel or focus mode"}, {Keys: "Ctrl+S", Does: "Save in the editor"}, {Keys: "Ctrl+Enter", Does: "Save a quick note"}, {Keys: "Ctrl+V", Does: "Paste an image as an attachment"}}},
		w.SheetHint{Text: "Shortcuts do nothing while the cursor is in a field. None of them destroy anything."}}, Foot: []w.SheetFootPart{w.Button{Label: "Close"}, w.Kbd{Keys: "Esc"}, w.Kbd{Keys: "?", Place: w.Pushed}}, Width: w.SheetWidth480}}}, "account-6", "tc-overlay", 1),
	overlay("plain", w.Overlay{Body: w.OverlayBody{Rail: []w.OverlayRailPart{w.Nav{Items: []w.NavItem{
		{Icon: "back", Label: "Back to app", Href: "#"}}}, w.RailSec{Items: []w.RailSecPart{
		w.Rubric{Text: "Settings"},
		w.Nav{Items: []w.NavItem{
			{Icon: "user", Label: "Account", Href: "#"},
			{Icon: "server", Label: "Connections", Count: &w.Num{Value: 3}, Href: "#"},
			{Icon: "file", Label: "Attachments", Href: "#"},
			{Icon: "device", Label: "New device", Href: "#"},
			{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
			{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
			{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
			{Icon: "theme", Label: "Appearance", Href: "#"}}, On: 2}}}, w.RailFoot{Items: []w.RailFootPart{w.RailStatus{Text: "davloom · running · 3 devices", Detail: w.Dim{Parts: []w.Inline{w.Text("Last response "), w.Mono{Text: "2026-08-16 09:41"}}, Size: w.Dim11}}, w.RailStatus{Text: "Backups are running", Detail: w.Dim{Parts: []w.Inline{w.Text("Last backup "), w.Mono{Text: "2026-08-16 03:10"}, w.Text("· ")}, Size: w.Dim11}}}, About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"}}}, Main: w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Text: "Connections", Size: w.HeadTitle19}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 servers · 12 collections · 2 subscriptions")}}}, Pad: w.HeadPadBottom8},
		w.Pad{Items: []w.PadPart{w.Card{Items: []w.CardPart{
			w.Rubric{Text: "nextcloud"},
			w.Mono{Text: "https://cloud.example.org/remote.php/dav/ · mix · last answer 2 min ago", Form: w.MonoLine, Variant: w.Mono11_5DustyGap}}}, w.Card{Items: []w.CardPart{
			w.Rubric{Text: "sftpgo · vps"},
			w.Mono{Text: "https://files.example.org/dav/ · mix · last answer 5 min ago", Form: w.MonoLine, Variant: w.Mono11_5DustyGap}}}}}}}, Height: w.ShellHeight280}, Scrim: w.Scrim{Content: w.Sheet{Icon: "shield", Title: "Show password", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("nextcloud")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Field{Label: "Your Carrel password", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Dim{Parts: []w.Inline{w.Text("Unlocks credentials everywhere for five minutes: passwords, device profiles, key escrow — anything handed out or handed over.")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Foot: []w.SheetFootPart{w.Button{Label: "Show", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Enter"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth400}, Pad: w.ScrimPad26}}, "s-profile-3", "tc-overlay", 1),
	overlay("plain", w.Overlay{Body: w.OverlayBody{Rail: []w.OverlayRailPart{w.Nav{Items: []w.NavItem{
		{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
		{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
		{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
		{Icon: "file", Label: "Files", Href: "#"},
		{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"}}, On: 2}, w.RailSec{Items: []w.RailSecPart{
		w.TickRow{Ticked: w.TickedPart, Kind: w.TickRowAll, Mark: w.Mark{Kind: w.MarkMulti}, Label: "Combined calendar", Meta: "14"}}}, w.RailSec{Items: []w.RailSecPart{
		w.Rubric{Text: "nextcloud"},
		w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "7", Current: true},
		w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "4"}}}}, Main: w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#4A6B52"}, Text: "Личный"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · 10 – 16 August 2026")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Icon: "plus", Kind: w.ButtonPrimary}}}},
		w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August"}, Count: &w.Num{Value: 2}},
		w.List{Rows: []w.ListRow{
			w.TimeRow{Mark: w.Mark{Colour: "#4A6B52"}, Time: "10:00", Title: []w.Inline{w.Text("Синк-ап по релизу")}, Detail: "Переговорная 2", Source: "nextcloud · Личный"},
			w.TimeRow{Mark: w.Mark{Colour: "#4A6B52"}, Time: "18:00", Title: []w.Inline{w.Text("Забрать документы")}, Detail: "МФЦ на Ленина", Source: "nextcloud · Личный"}}}}}, Height: w.ShellHeight390}, Scrim: w.Scrim{Content: w.Sheet{Title: "Publish a calendar", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("nextcloud · Личный")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Field{Label: "How much of it", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Line{Block: true, Gap: w.LineGap10, Items: []w.LinePart{w.Field{Label: "From", Width: w.FieldSM, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "To", Width: w.FieldSM, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}},
		w.Field{Label: "Link lives for", Width: w.FieldSM, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.SheetChecks{Items: []w.Check{w.Check{On: true, Parts: []w.CheckPart{w.Text("Allow"), w.Mono{Text: ".ics"}, w.Text("download and subscription")}}, w.Check{Parts: []w.CheckPart{w.Text("Show event descriptions")}}}}}, Foot: []w.SheetFootPart{w.Button{Label: "Create link", Kind: w.ButtonPrimary}, w.Button{Label: "Cancel"}}}, Pad: w.ScrimPad26}}, "s-publish-2", "tc-overlay", 1),

	// tc-dialog
	overlay("width", w.Dialog{Items: []w.DialogPart{
		w.Head{Title: &w.HeadTitle{Text: "Merge on the server"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Артём Белов · 2 records")}}}},
		w.Msg{Parts: []w.MsgPart{w.Strong("One record is rewritten and 1 other deleted."), w.Text(" The kept record is rewritten with the values you chose, and only once the server has kept it are the others deleted. This cannot be undone from here.")}, Space: w.MsgFlushSides},
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Kept and rewritten"},
			w.CardText{Kind: w.CardLine, Parts: []w.CardTextPart{w.Text("Артём Белов "), w.Dim{Parts: []w.Inline{w.Text("· nextcloud · Контакты")}}}}}, Variant: w.CardApart},
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Deleted afterwards"},
			w.CardText{Kind: w.CardLine, Parts: []w.CardTextPart{w.Text("Белов Артём "), w.Dim{Parts: []w.Inline{w.Text("· radicale · vps · Команда")}}}}}},
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "The kept record gains"},
			w.CardList{Items: []string{"EMAIL;TYPE=work: belov@team.example", "PHOTO (from radicale · vps · Команда)"}}}},
		w.FormFoot{Variant: w.FormFootFlush, Items: []w.FormFootPart{w.Button{Label: "Write the merge and delete the rest", Kind: w.ButtonDangerSolid}, w.Button{Label: "Cancel"}, w.Button{Label: "Keep Белов Артём instead", Place: w.Pushed}}}}}, "s-dup-2", "tc-dialog", 1),
	overlay("wide", w.Dialog{Items: []w.DialogPart{
		w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 2},
		w.Head{Title: &w.HeadTitle{Text: "Import contacts"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mono{Text: "Разное (mix).vcf"}, w.Text(" · 84 KB · vCard 3.0")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Import 122 contacts", Kind: w.ButtonPrimary}}}},
		w.Form{},
		w.Stats{Items: []w.Stat{w.Stat{Value: "122", Label: "ready"}, w.Stat{Value: "3", Label: "with errors", Error: true}, w.Stat{Value: "5", Label: "UID collisions"}}},
		w.Table{Cols: []w.TableCol{{Label: "Source"}, {Label: "Name"}, {Label: "UID", Rank: w.Rank3}, {Label: "Result", Rank: w.Rank4}}, Rows: []w.TableRow{
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.Text("Разное (mix).vcf")}, Mono: true}, {Parts: []w.CellPart{w.Text("Анна Ковалёва")}}, {Parts: []w.CellPart{w.Text("4f2b-9c31-a0e7")}, Mono: true}, {Parts: []w.CellPart{w.Text("Ready")}}}},
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.Text("Разное (mix).vcf")}, Mono: true}, {Parts: []w.CellPart{w.Text("Артём Белов")}}, {Parts: []w.CellPart{w.Text("1a77-40de-bb02")}, Mono: true}, {Parts: []w.CellPart{w.Text("New UID will be assigned")}}}},
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.Text("Разное (mix).vcf")}, Mono: true}, {Parts: []w.CellPart{w.Text("— (no FN)")}, Dim: true}, {Parts: []w.CellPart{w.Text("—")}, Mono: true}, {Parts: []w.CellPart{w.Text("vCard 2.1 quoted-printable: cannot decode line 12")}, Tone: w.CellDanger}}}}},
		w.Msg{Parts: []w.MsgPart{w.Text("Import always creates. A UID that already exists gets a new one and is listed in the report.")}}}, Variant: w.DialogWide}, "s-io-5", "tc-dialog", 1),
	overlay("page", w.Dialog{Items: []w.DialogPart{
		w.DialogTitle{Text: "What we can do", Gap: w.DialogTitleGap8},
		w.Prose{Parts: []w.Inline{w.Text("Account "), w.Strong("mix"), w.Text(" has a confirmed email. A silent password reset does not exist: a new password would be a new key, and everything stored under the old one would be lost anyway.")}, Variant: w.ProseMuted},
		w.Form{},
		w.FormFoot{Variant: w.FormFootFlush, Items: []w.FormFootPart{w.Button{Label: "Back to sign in"}}}}, Variant: w.DialogPage520}, "s-auth-6", "tc-dialog", 1),
	overlay("width", w.Dialog{Items: []w.DialogPart{
		w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 3},
		w.Head{Title: &w.HeadTitle{Text: "Import report"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}}}},
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Created"},
			w.CardText{Kind: w.CardNum, Parts: []w.CardTextPart{w.Text("122")}}}},
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "UID collisions", Items: []w.RubricPart{w.Mono{Text: "5", Variant: w.MonoTightPlain}}},
			w.Diag{Text: "1a77-40de-bb02 → 9c02-4411-7ee1 33b0-9ffa-1120 → 55da-8e10-02cc"}}},
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Failed", Items: []w.RubricPart{w.Mono{Text: "3", Variant: w.MonoTightPlain}}},
			w.Diag{Text: "line 12: cannot decode quoted-printable", Variant: w.DiagDanger}}},
		w.FormFoot{Variant: w.FormFootFlush, Items: []w.FormFootPart{w.Button{Label: "Back to list", Kind: w.ButtonPrimary}}}}, Variant: w.Dialog520}, "s-io-9", "tc-dialog", 1),
	overlay("width", w.Dialog{Items: []w.DialogPart{
		w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 1},
		w.Head{Title: &w.HeadTitle{Text: "Import contacts"}},
		w.Form{},
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "From a file"},
			w.Field{Label: "vCard file", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
			w.CardHint{Parts: []w.CardTextPart{w.Text("vCard 2.1, 3.0 and 4.0. One file holds any number of cards; the calendar and the list of tasks take "), w.Mono{Text: ".ics"}, w.Text(" here instead.")}, Line: true}}},
		w.Msg{Parts: []w.MsgPart{w.Text("Nothing is written until the preview has been read: the next step counts what was parsed and names what failed.")}}}, Variant: w.Dialog560}, "s-io-3", "tc-dialog", 1),
	overlay("narrow", w.Dialog{Items: []w.DialogPart{
		w.DialogLogo{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo48}, Wide: true},
		w.DialogTitle{Text: "Check your email", Gap: w.DialogTitleGap6},
		w.Prose{Parts: []w.Inline{w.Text("The letter went to "), w.Mono{Text: "v@example.org"}, w.Text(". The account cannot be used until you open the link in it. Sending again with the same login and address, while the account is unconfirmed, simply replaces the letter and the password.")}, Variant: w.ProseMuted},
		w.FormFoot{Variant: w.FormFootFlush, Items: []w.FormFootPart{w.Button{Label: "Back to sign in"}}}}, Variant: w.DialogNarrow}, "s-auth-3", "tc-dialog", 1),
	overlay("width", w.Dialog{Items: []w.DialogPart{
		w.Head{Title: &w.HeadTitle{Text: "Conflict"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Another client changed this contact while you were editing it. Nothing has been overwritten.")}}}},
		w.DiffAgree{Property: w.Mono{Text: "FN"}, Text: "agrees — kept as it is."},
		w.Diff{Items: []w.DiffItem{
			{Title: w.Rubric{Text: "", Lead: &w.Mono{Text: "TEL;work"}, Items: []w.RubricPart{w.RubricAside{Text: "only in your edit"}}}, Options: []w.Check{w.Check{On: true, Parts: []w.CheckPart{w.Choice{Source: "Your edit · include", Value: "+7 495 120-45-00"}}}, w.Check{Parts: []w.CheckPart{w.Choice{Source: "Leave it out", Value: "—"}}}}},
			{Title: w.Rubric{Text: "", Lead: &w.Mono{Text: "EMAIL;work"}, Items: []w.RubricPart{w.RubricAside{Text: "both changed"}}}, Options: []w.Check{w.Check{On: true, Parts: []w.CheckPart{w.Choice{Source: "Your edit", Value: "a.kovaleva@example.org"}}}, w.Check{Parts: []w.CheckPart{w.Choice{Source: "Server", Value: "anna.kovaleva@slovo.example"}}}}},
			{Title: w.Rubric{Text: "", Lead: &w.Mono{Text: "NOTE"}, Items: []w.RubricPart{w.RubricAside{Text: "only in your edit"}}}, Options: []w.Check{w.Check{On: true, Parts: []w.CheckPart{w.Choice{Source: "Your edit · include", Value: "Пишет по вторникам, отвечает в тот же день."}}}, w.Check{Parts: []w.CheckPart{w.Choice{Source: "Leave it out", Value: "—"}}}}},
			{Title: w.Rubric{Text: "", Lead: &w.Mono{Text: "X-EVOLUTION-FILE-AS"}, Items: []w.RubricPart{w.RubricAside{Text: "only on the server"}}}, Options: []w.Check{w.Check{Parts: []w.CheckPart{w.Choice{Source: "Your edit", Value: "—"}}}, w.Check{On: true, Parts: []w.CheckPart{w.Choice{Source: "Server · include", Value: "Ковалёва, Анна"}}}}}}},
		w.FormFoot{Variant: w.FormFootFlush, Items: []w.FormFootPart{w.Button{Label: "Write this combination", Kind: w.ButtonPrimary}, w.Button{Label: "Use only my edit"}, w.Button{Label: "Use only the server version"}, w.Button{Label: "Cancel"}}}}}, "s-conflict-1", "tc-dialog", 1),
	overlay("width", w.Dialog{Items: []w.DialogPart{
		w.Head{Title: &w.HeadTitle{Text: "Connect account"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("One connection gives calendars and address books alike.")}}}},
		w.Form{},
		w.FormFoot{Variant: w.FormFootFlush, Items: []w.FormFootPart{w.Button{Label: "Connect", Kind: w.ButtonPrimary}, w.Button{Label: "Test only"}, w.Button{Label: "Cancel"}}}}, Variant: w.Dialog540}, "s-profile-2", "tc-dialog", 1),
	overlay("page", w.Dialog{Items: []w.DialogPart{
		w.DialogTitle{Text: "Create an account", Gap: w.DialogTitleGap4},
		w.DialogSub{Parts: []w.Inline{w.Text("A confirmation link will be sent to the address you give; the account cannot be used until you open it.")}},
		w.Msg{Parts: []w.MsgPart{w.Text("Your password is the only key to your own data. Nobody — including this server — can recover it for you.")}, Tone: w.MsgAlert, Space: w.MsgFlushBottom12},
		w.Form{},
		w.FormFoot{Variant: w.FormFootFlush, Items: []w.FormFootPart{w.Button{Label: "Create account", Kind: w.ButtonPrimary}, w.Button{Label: "Already have an account? Sign in"}}}}, Variant: w.DialogPage440}, "s-auth-2", "tc-dialog", 1),
	overlay("canvas", w.Dialog{Items: []w.DialogPart{
		w.Lede{Text: "Центрированное полотно для форм, мастеров и диалогов выбора."}}, Variant: w.DialogCanvas420}, "components-31", "tc-dialog", 1),
	overlay("width", w.Dialog{Items: []w.DialogPart{
		w.Head{Title: &w.HeadTitle{Text: "Photo"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Анна Ковалёва · metadata will be stripped")}}}},
		w.PhotoCrop{Colour: "#8C7B6A", Note: "Drag the photo to choose the frame. The box is what other clients will see.", Fields: []w.Field{w.Field{Label: "Zoom", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Rotate", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Saved as", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}, After: &w.Dim{Parts: []w.Inline{w.Text("EXIF orientation is already applied, metadata is stripped.")}, Size: w.Dim11}},
		w.FormFoot{Variant: w.FormFootFlush, Items: []w.FormFootPart{w.Button{Label: "Save photo", Kind: w.ButtonPrimary}, w.Button{Label: "Cancel"}, w.Button{Title: "More", Place: w.Pushed}}}}, Variant: w.Dialog520}, "s-contacts-5", "tc-dialog", 1),
	overlay("narrow", w.Dialog{Items: []w.DialogPart{
		w.DialogLogo{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo60}},
		w.Tagline{Lead: "One desk. Every server.", Text: "Calendars, contacts, tasks, notes and files — one place.", Centered: true},
		w.AuthStack{Items: []w.AuthStackPart{w.Form{}, w.Button{Label: "Sign in", Kind: w.ButtonPrimary}}, Links: []string{"Create an account", "Forgotten your password?"}}}, Variant: w.DialogNarrow}, "s-auth-1", "tc-dialog", 1),
	overlay("page", w.Dialog{Items: []w.DialogPart{
		w.DialogTitle{Text: "Forgotten your password?", Gap: w.DialogTitleGap8},
		w.Prose{Parts: []w.Inline{w.Text("Enter the login or the confirmed email of the account. There is no way to set a new password and keep the data stored here: the encryption key is derived from the password, and the server never has that key.")}, Variant: w.ProseMutedClose},
		w.Form{},
		w.FormFoot{Variant: w.FormFootFlush, Items: []w.FormFootPart{w.Button{Label: "Continue", Kind: w.ButtonPrimary}, w.Button{Label: "Back to sign in"}}}}, Variant: w.DialogPage440}, "s-auth-5", "tc-dialog", 1),
	overlay("page", w.Dialog{Items: []w.DialogPart{
		w.DialogTitle{Text: "Set a new password", Gap: w.DialogTitleGap4},
		w.DialogSub{Parts: []w.Inline{w.Text("The account "), w.Mono{Text: "mix"}, w.Text(" is empty. Nothing was kept.")}},
		w.Msg{Parts: []w.MsgPart{w.Text("Your password is the only key to your own data. Nobody — including this server — can recover it for you.")}, Tone: w.MsgAlert, Space: w.MsgFlushBottom12},
		w.Form{},
		w.FormFoot{Variant: w.FormFootFlush, Items: []w.FormFootPart{w.Button{Label: "Set the password and sign in", Kind: w.ButtonPrimary}}}}, Variant: w.DialogPage440}, "s-auth-11", "tc-dialog", 1),
	overlay("narrow", w.Dialog{Items: []w.DialogPart{
		w.DialogLogo{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo42}, Wide: true},
		w.DialogTitle{Text: "Carrel"},
		w.DialogSub{Parts: []w.Inline{w.Mono{Text: "0.0.1 · 786949a"}}},
		w.Tagline{Lead: "One desk. Every server.", Text: "Calendars, contacts, tasks, notes and files — one place."},
		w.Prose{Parts: []w.Inline{w.Text("AGPL-3.0-or-later: if you run a modified version for others, you must make the source available to them.")}, Variant: w.ProseMuted},
		w.FormFoot{Variant: w.FormFootFlush, Items: []w.FormFootPart{w.Button{Label: "Source code", Small: true}, w.Button{Label: "Third-party licences", Small: true}, w.Button{Title: "More", Place: w.Pushed}}}}, Variant: w.DialogNarrow}, "rest-3", "tc-dialog", 1),

	// tc-auth
	overlay("plain", w.Auth{Dialog: w.Dialog{Items: []w.DialogPart{
		w.DialogLogo{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo60}},
		w.Tagline{Lead: "One desk. Every server.", Text: "Calendars, contacts, tasks, notes and files — one place.", Centered: true},
		w.AuthStack{Items: []w.AuthStackPart{w.Form{}, w.Button{Label: "Sign in", Kind: w.ButtonPrimary}}, Links: []string{"Create an account", "Forgotten your password?"}}}, Variant: w.DialogNarrow}}, "s-auth-1", "tc-auth", 1),
)

func init() {
	Widgets = append(Widgets, w.Menu{}, w.Drop{Open: true}, w.Sheet{}, w.Scrim{}, w.Overlay{}, w.Dialog{}, w.Auth{})
	Parts = append(Parts, w.MenuItem{}, w.MenuSep{}, w.MenuHeading{}, w.MenuField{}, w.MenuColumn{}, w.MenuFoot{}, w.TopDrop{Open: true},
		w.SheetSubject{}, w.SheetFields{}, w.SheetKey{}, w.SheetKeys{}, w.SheetHint{}, w.SheetNote{}, w.SheetLine{}, w.SheetBlock{},
		w.SheetChecks{}, w.SheetFrom{}, w.SheetPick{}, w.DialogTitle{}, w.DialogSub{}, w.DialogLogo{}, w.Tagline{}, w.AuthStack{},
		w.Lede{}, w.PhotoCrop{}, w.OverlayBody{}, w.DiffAgree{})
}
