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
	overlay("danger", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "pencil", Label: "Edit"},
		w.MenuSep{},
		w.MenuItem{Icon: "trash", Label: "Delete job", Tone: w.MenuToneDanger}}, Width: w.MenuWidth212}, "s-backup-10", "tc-menu", 1),
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
	overlay("more", w.Drop{Trigger: w.Button{Icon: "dots", Title: "More"}, More: true, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "pencil", Label: "Rename"},
		w.MenuItem{Icon: "down", Detail: &w.Mono{Text: ".vcf", Variant: w.Mono11Dusty}, Label: "Export"}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}}, "components-19", "tc-drop", 1),
	overlay("pick", w.Drop{Trigger: w.Pick{Value: "Синк-ап по релизу ковал", Label: "Where"}, Menu: w.Menu{Items: []w.MenuPart{
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
	overlay("drop", w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}}, Actions: []w.TopAction{w.TopDrop{Trigger: w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary, Focus: true}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "quill", Trail: w.Kbd{Keys: "N"}, Label: "Note", On: true},
		w.MenuSep{},
		w.MenuItem{Icon: "con", Trail: w.Kbd{Keys: "C"}, Label: "Contact"},
		w.MenuItem{Icon: "cal", Trail: w.Kbd{Keys: "E"}, Label: "Event"},
		w.MenuItem{Icon: "task", Trail: w.Kbd{Keys: "T"}, Label: "Task"},
		w.MenuSep{},
		w.MenuItem{Icon: "up", Label: "Upload file"},
		w.MenuSep{},
		w.MenuItem{Icon: "mail", Label: "Letter"}}}}, w.Button{Label: "2 min ago", Icon: "refresh"}, w.UserMenu{Name: "mix"}}}, "s-notes-7", "tc-top", 1),
	overlay("drop", w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}}, Actions: []w.TopAction{w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary}, w.Button{Label: "2 min ago", Icon: "refresh"}, w.TopDrop{Trigger: w.UserMenu{Name: "mix", Focus: true}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "mix", Items: []w.RubricPart{w.Mono{Text: "carrel.example", Variant: w.MonoTightPlain}}, Form: w.RubricItem, Variant: w.RubricItemLine}, Form: w.MenuHeadingSplit},
		w.MenuItem{Icon: "gear", Label: "Settings"},
		w.MenuItem{Icon: "adm", Label: "Administration"},
		w.MenuSep{},
		w.MenuColumn{Label: "Theme", Control: w.Seg{Options: []w.SegOption{{Label: "Auto", On: true}, {Label: "Light"}, {Label: "Dark"}}, Variant: w.SegStartSmall}},
		w.MenuItem{Icon: "search", Trail: w.Kbd{Keys: "?"}, Label: "Keyboard shortcuts"},
		w.MenuSep{},
		w.MenuItem{Label: "Sign out"}}, Width: w.MenuWidth290}, End: true}}}, "account-5", "tc-top", 1),
	overlay("more", w.Bar{Items: []w.BarPart{w.Seg{Options: []w.SegOption{{Label: "A–Z", On: true}, {Label: "Recently changed"}}}, w.Seg{Options: []w.SegOption{{Label: "All", On: true}, {Label: "Overdue"}}, Secondary: true}, w.Button{Label: "Columns", Icon: "cols"}, w.Drop{Trigger: w.Button{Icon: "dots", Title: "More"}, More: true, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "All", On: true},
		w.MenuItem{Label: "Overdue"}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}}}}, "components-19", "tc-bar", 1),
	overlay("more", w.Bar{Items: []w.BarPart{w.Seg{Options: []w.SegOption{{Label: "Day"}, {Label: "Week", On: true}, {Label: "Month"}, {Label: "Range"}}}, w.Range{From: w.Input{Kind: w.InputDate, Value: "2026-08-10", Mono: true, Label: "Date"}, To: w.Input{Kind: w.InputDate, Value: "2026-08-16", Mono: true, Label: "Date"}, Show: w.Button{Label: "Show"}, Secondary: true}, w.BarSep{Secondary: true}, w.Check{On: true, Parts: []w.CheckPart{w.Text("Events")}}, w.Check{Parts: []w.CheckPart{w.Text("Tasks due")}}, w.Check{Parts: []w.CheckPart{w.Text("Notes")}}, w.Drop{Trigger: w.Button{Icon: "dots", Title: "More"}, More: true, Menu: w.Menu{Items: []w.MenuPart{
		w.Range{From: w.Input{Kind: w.InputDate, Value: "2026-08-10", Mono: true, Label: "Date"}, To: w.Input{Kind: w.InputDate, Value: "2026-08-16", Mono: true, Label: "Date"}, Show: w.Button{Label: "Show"}, Wrap: true},
		w.MenuSep{},
		w.MenuItem{Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Compact rows"}}, Pop: w.MenuUnderEnd, Inline: true}}}, Right: &w.BarRight{Items: []w.BarPart{w.Check{Parts: []w.CheckPart{w.Text("Compact rows")}}}, Secondary: true}, Variant: w.BarWide}, "s-agenda-1", "tc-bar", 1),

	// tc-sheet
	overlay("plain", w.Sheet{Title: "Publish a calendar", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("nextcloud · Личный")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Field{Label: "Link", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}, Foot: []w.SheetFootPart{w.Button{Label: "Send by letter", Kind: w.ButtonPrimary}, w.Button{Label: "Open the link"}, w.Button{Label: "All published links"}}}, "s-publish-13", "tc-sheet", 1),
	overlay("close", w.Sheet{Icon: "quill", Title: "New note", Head: []w.SheetHeadPart{w.Pick{Value: "nextcloud · Записи", Variant: w.PickT12W200, Label: "Where", Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.Area{Text: "Позвонить в МФЦ во вторник до 12:00 — уточнить, нужен ли оригинал доверенности.", Variant: w.AreaBodySm, Focus: true, Label: "Text"}}, Foot: []w.SheetFootPart{w.Button{Label: "Save note", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Ctrl+Enter"}, w.Button{Label: "Open the card", Icon: "expand"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}}, "s-notes-5", "tc-sheet", 1),
	overlay("close", w.Sheet{Icon: "con", Title: "New contact", Head: []w.SheetHeadPart{w.Pick{Value: "nextcloud · Контакты", Variant: w.PickT12W200, Label: "Where", Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.SheetFields{Items: []w.Field{w.Field{Label: "Display name", Place: w.Grown, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Phone", Width: w.FieldFill200, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Email", Width: w.FieldFill200, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}},
		w.Dim{Parts: []w.Inline{w.Text("The first word is Given, the rest Family — «Пётр Ковалёв», not «Ковалёв Пётр». Everything else is filled in on the card.")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Foot: []w.SheetFootPart{w.Button{Label: "Save contact", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Enter"}, w.Button{Label: "Open the card", Icon: "expand"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}}, "s-notes-8", "tc-sheet", 1),
	overlay("close", w.Sheet{Icon: "task", Title: "New task", Head: []w.SheetHeadPart{w.Pick{Value: "nextcloud · Дела", Variant: w.PickT12W200, Label: "Where", Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.SheetFields{Items: []w.Field{w.Field{Label: "Summary", Place: w.Grown, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Due date", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Due time", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}},
		w.Dim{Parts: []w.Inline{w.Text("A task with no due date is a task all the same — the field may stay empty.")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Foot: []w.SheetFootPart{w.Button{Label: "Save task", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Enter"}, w.Button{Label: "Open the card", Icon: "expand"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth520}, "s-notes-10", "tc-sheet", 1),
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
	overlay("plain", w.Sheet{Title: "Show password", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("nextcloud")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Field{Label: "Your Carrel password", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Dim{Parts: []w.Inline{w.Text("Unlocks credentials everywhere for five minutes: passwords, device profiles, anything else handed out.")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Foot: []w.SheetFootPart{w.Button{Label: "Show", Kind: w.ButtonPrimary}, w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc", Place: w.Pushed}}, Width: w.SheetWidth400}, "s-plan-3", "tc-sheet", 1),
	overlay("plain", w.Sheet{Title: "Publish a calendar", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("nextcloud · Личный")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Field{Label: "How much of it", Width: w.FieldW100, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Line{Block: true, Gap: w.LineGap10, Items: []w.LinePart{w.Field{Label: "From", Width: w.FieldSM, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "To", Width: w.FieldSM, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}},
		w.Field{Label: "Link lives for", Width: w.FieldSM, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.SheetChecks{Items: []w.Check{w.Check{On: true, Parts: []w.CheckPart{w.Text("Allow download and subscription")}}, w.Check{Parts: []w.CheckPart{w.Text("Show event descriptions")}}}}}, Foot: []w.SheetFootPart{w.Button{Label: "Create link", Kind: w.ButtonPrimary}, w.Button{Label: "Cancel"}}}, "s-publish-2", "tc-sheet", 1),
	overlay("close", w.Sheet{Title: "Sources", Head: []w.SheetHeadPart{w.Button{Label: "Apply", Small: true, Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.RailSec{Items: []w.RailSecPart{
			w.TickRow{Ticked: w.TickedPart, Kind: w.TickRowAll, Mark: w.Mark{Kind: w.MarkMulti}, Label: "Combined notebook", Meta: "86", Current: true}}},
		w.RailSec{Items: []w.RailSecPart{
			w.Rubric{Text: "nextcloud"},
			w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "64"},
			w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "22"},
			w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "—", Hidden: true}}}}}, "m-phone-9", "tc-sheet", 1),

	// tc-scrim
	overlay("pad", w.Scrim{Content: w.Sheet{Title: "Sheet"}, Pad: w.ScrimPad26}, "s-contacts-2", "tc-scrim", 1),
	overlay("plain", w.Scrim{Content: w.Sheet{Icon: "quill", Title: "New note", Head: []w.SheetHeadPart{w.Pick{Value: "nextcloud · Записи", Variant: w.PickT12W200, Label: "Where", Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.Area{Text: "Позвонить в МФЦ во вторник до 12:00 — уточнить, нужен ли оригинал доверенности.", Variant: w.AreaBodySm, Focus: true, Label: "Text"}}, Foot: []w.SheetFootPart{w.Button{Label: "Save note", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Ctrl+Enter"}, w.Button{Label: "Open the card", Icon: "expand"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}}}, "s-notes-5", "tc-scrim", 1),

	// tc-overlay
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
		w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "51"}}}}, Main: w.Main{}, Height: w.ShellHeight560}, Scrim: w.Scrim{Content: w.Sheet{Title: "Sheet"}, Pad: w.ScrimPad26}}, "s-contacts-2", "tc-overlay", 1),
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
			{Icon: "theme", Label: "Appearance", Href: "#"}}, On: 2}}}, w.RailFoot{Items: []w.RailFootPart{w.RailStatus{Text: "davloom · running · 3 devices", Detail: w.Dim{Parts: []w.Inline{w.Text("Last response "), w.Mono{Text: "2026-08-16 09:41"}}, Size: w.Dim11}}, w.RailStatus{Text: "Backups are running", Detail: w.Dim{Parts: []w.Inline{w.Text("Last backup "), w.Mono{Text: "2026-08-16 03:10"}, w.Text("· ")}, Size: w.Dim11}}}, About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"}}}, Main: w.Main{}, Height: w.ShellHeight280}, Scrim: w.Scrim{Content: w.Sheet{Icon: "shield", Title: "Show password", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("nextcloud")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
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
		w.SheetChecks{Items: []w.Check{w.Check{On: true, Parts: []w.CheckPart{w.Text("Allow download and subscription")}}, w.Check{Parts: []w.CheckPart{w.Text("Show event descriptions")}}}}}, Foot: []w.SheetFootPart{w.Button{Label: "Create link", Kind: w.ButtonPrimary}, w.Button{Label: "Cancel"}}}, Pad: w.ScrimPad26}}, "s-publish-2", "tc-overlay", 1),

	// tc-dialog
	overlay("width", w.Dialog{Items: []w.DialogPart{
		w.Head{Title: &w.HeadTitle{Text: "Connect account"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("One connection gives calendars and address books alike.")}}}},
		w.Form{},
		w.FormFoot{Variant: w.FormFootFlush, Items: []w.FormFootPart{w.Button{Label: "Connect", Kind: w.ButtonPrimary}, w.Button{Label: "Test only"}, w.Button{Label: "Cancel"}}}}, Variant: w.Dialog540}, "s-profile-2", "tc-dialog", 1),
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

	// tc-auth
	overlay("plain", w.Auth{Dialog: w.Dialog{Items: []w.DialogPart{
		w.DialogLogo{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo60}},
		w.Tagline{Lead: "One desk. Every server.", Text: "Calendars, contacts, tasks, notes and files — one place.", Centered: true},
		w.AuthStack{Items: []w.AuthStackPart{w.Form{}, w.Button{Label: "Sign in", Kind: w.ButtonPrimary}}, Links: []string{"Create an account", "Forgotten your password?"}}}, Variant: w.DialogNarrow}}, "s-auth-1", "tc-auth", 1),
)

func init() {
	Widgets = append(Widgets, w.Menu{}, w.Drop{}, w.Sheet{}, w.Scrim{}, w.Overlay{}, w.Dialog{}, w.Auth{})
	Parts = append(Parts, w.MenuItem{}, w.MenuSep{}, w.MenuHeading{}, w.MenuField{}, w.MenuColumn{}, w.MenuFoot{}, w.TopDrop{},
		w.SheetSubject{}, w.SheetFields{}, w.SheetKey{}, w.SheetKeys{}, w.SheetHint{}, w.SheetNote{}, w.SheetLine{}, w.SheetBlock{},
		w.SheetChecks{}, w.SheetFrom{}, w.SheetPick{}, w.DialogTitle{}, w.DialogSub{}, w.DialogLogo{}, w.Tagline{}, w.AuthStack{},
		w.Lede{}, w.PhotoCrop{}, w.OverlayBody{})
}

var _ = wait("children from later families (messages, cards, tables, prose, steps, diagnostics, states)",
	ref("s-contacts-2", "tc-sheet", 1), ref("s-notes-6", "tc-sheet", 1), ref("account-2", "tc-sheet", 1), ref("s-profile-6", "tc-sheet", 1), ref("s-profile-9", "tc-sheet", 1), ref("s-mail-1", "tc-sheet", 1), ref("s-mail-4", "tc-sheet", 1), ref("s-mail-5", "tc-sheet", 1), ref("s-mail-10", "tc-sheet", 1), ref("s-mail-6", "tc-sheet", 1), ref("s-io-1", "tc-scrim", 1), ref("s-notes-6", "tc-overlay", 1), ref("s-dup-2", "tc-dialog", 1), ref("s-io-5", "tc-dialog", 1), ref("s-auth-6", "tc-dialog", 1), ref("s-io-9", "tc-dialog", 1), ref("s-io-3", "tc-dialog", 1), ref("s-auth-3", "tc-dialog", 1), ref("s-conflict-1", "tc-dialog", 1), ref("s-auth-2", "tc-dialog", 1), ref("s-auth-5", "tc-dialog", 1), ref("s-auth-11", "tc-dialog", 1), ref("rest-3", "tc-dialog", 1))
