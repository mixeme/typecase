package refs

import w "carrel.dev/typecase/widget"

var _ = add(
	// tc-head
	shell("group", w.Head{Avatar: &w.Avatar{Initials: "РС", Size: w.Avatar44}, Title: &w.HeadTitle{Text: "Редакция «Слово»"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Group · 4 members")}}, w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Edit", Icon: "pencil", Kind: w.ButtonPrimary}, w.Button{Label: "Print", Icon: "print"}, w.Button{Icon: "dots", Title: "More"}}}}, "mentions-9", "tc-head", 1),
	shell("acts", w.Head{Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti}, Text: "Combined address book"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 312 cards")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "New contact", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}, w.Button{Label: "Import", Icon: "up"}, w.Button{Label: "Export", Icon: "down"}, w.Button{Label: "Print", Icon: "print"}}}}, "s-contacts-1", "tc-head", 1),
	shell("plain", w.Head{Title: &w.HeadTitle{Text: "Photo"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Анна Ковалёва · metadata will be stripped")}}}}, "s-contacts-5", "tc-head", 1),
	shell("acts", w.Head{Title: &w.HeadTitle{Text: "Анна Ковалёва"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Save", Kind: w.ButtonPrimary}}}}, "s-contacts-3", "tc-head", 1),
	shell("acts", w.Head{Title: &w.HeadTitle{Text: "Set up a new device"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Everything discovery found — as one file or one code "), w.Badge{Label: "credentials unlocked · 4:12"}, w.Button{Label: "Lock now", Small: true}}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Seg{Options: []w.SegOption{{Label: "Apple", On: true}, {Label: "Android"}, {Label: "Thunderbird"}, {Label: "Anything else"}}}}}}, "s-plan-1", "tc-head", 1),
	shell("acts", w.Head{Title: &w.HeadTitle{Text: "Import contacts"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mono{Text: "Разное (mix).vcf"}, w.Text(" · 84 KB · vCard 3.0")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Import 122 contacts", Kind: w.ButtonPrimary}}}}, "s-io-5", "tc-head", 1),
	shell("padded", w.Head{Title: &w.HeadTitle{Text: "Connections"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 servers · 12 collections · 2 subscriptions")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Connect account", Icon: "plus", Kind: w.ButtonPrimary}}}, Pad: w.HeadPadBottom8}, "s-profile-7", "tc-head", 1),
	shell("padded", w.Head{Title: &w.HeadTitle{Text: "2026-08-10-sync-board.png", Size: w.HeadTitle17}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("sftpgo · vps · Documents")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Icon: "close", Title: "Close the panel"}}}, Pad: w.HeadPadBottom8}, "s-files-1", "tc-head", 2),
	shell("plain", w.Head{Title: &w.HeadTitle{Text: "Import contacts"}}, "s-io-3", "tc-head", 1),
	shell("acts", w.Head{Title: &w.HeadTitle{Text: "Connections"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Signed in as "), w.Strong("mix"), w.Text(" · 3 servers · 12 collections · 2 subscriptions")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Connect account", Icon: "plus", Kind: w.ButtonPrimary}, w.Button{Label: "Subscribe to a feed", Icon: "link"}}}}, "s-profile-1", "tc-head", 1),
	shell("padded", w.Head{Title: &w.HeadTitle{Text: "attachments", Size: w.HeadTitle18}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("sftpgo · vps")}}}, Pad: w.HeadPadTop11}, "components-16", "tc-head", 1),
	shell("photo", w.Head{Photo: &w.HeadPhoto{}, Title: &w.HeadTitle{Text: "Анна Ковалёва"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Издательство «Слово» · Выпускающий редактор")}}, w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Edit", Icon: "pencil", Kind: w.ButtonPrimary}, w.Button{Label: "Send", Icon: "mail"}, w.Button{Label: "Print", Icon: "print"}, w.Button{Icon: "dots", Title: "More"}}}}, "s-contacts-6", "tc-head", 1),
	shell("acts", w.Head{Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti}, Text: "Combined calendar"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("5 of 7 calendars · 3 accounts · 10 – 16 August 2026 · "), w.Mono{Text: "Europe/Moscow"}}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Icon: "plus", Kind: w.ButtonPrimary}, w.Button{Label: "Import", Icon: "up"}, w.Button{Label: "Export", Icon: "down"}, w.Button{Label: "Print", Icon: "print"}}}}, "s-agenda-1", "tc-head", 1),
	shell("acts", w.Head{Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#4A6B52"}, Text: "Синк-ап по релизу", Size: w.HeadTitle19}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Личный")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Edit", Small: true}, w.Button{Icon: "dots", Small: true, Title: "More"}, w.Button{Icon: "close", Title: "Close the panel"}}}}, "s-agenda-1", "tc-head", 2),
	shell("acts", w.Head{Title: &w.HeadTitle{Text: "Собрать чемодан", Size: w.HeadTitle19}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Дела")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Edit", Small: true}, w.Button{Icon: "dots", Small: true, Title: "More"}, w.Button{Icon: "close", Title: "Close the panel"}}}}, "s-tasks-1", "tc-head", 2),
	shell("meta", w.Head{Meta: &w.HeadSub{Parts: []w.HeadSubPart{w.Mark{Colour: "#4A6B52"}, w.Text("nextcloud · Записи · "), w.Mono{Text: "10 Aug 2026 11:05"}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Edit", Icon: "pencil", Kind: w.ButtonPrimary}, w.Button{Label: "Export Markdown"}, w.Button{Label: "Print", Icon: "print"}, w.Button{Label: "Source"}, w.Seg{Options: []w.SegOption{{Label: "Full width", On: true}, {Label: "Reading width"}}}, w.Button{Icon: "panel", Title: "Hide the side column"}, w.Button{Icon: "dots", Title: "More"}}}, Pad: w.HeadRule}, "s-notes-2", "tc-head", 1),
	shell("crumbs", w.Head{Title: &w.HeadTitle{Text: "attachments"}, Crumbs: &w.Crumbs{Path: []w.Crumb{{Label: "Documents", Href: "#"}, {Label: "carrel", Href: "#"}}, Current: "attachments"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("sftpgo · vps · "), w.Badge{Label: "attachment folder"}}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Upload", Menu: true, Icon: "up", Kind: w.ButtonPrimary}, w.Button{Label: "New folder", Icon: "plus"}, w.Button{Icon: "refresh", Title: "More"}, w.Button{Icon: "panel", Variant: w.ButtonOnCard, Title: "Details"}, w.Button{Icon: "dots", Title: "More"}}}}, "s-files-1", "tc-head", 1),
	shell("crumbs", w.Head{Title: &w.HeadTitle{Text: "attachments"}, Crumbs: &w.Crumbs{Path: []w.Crumb{{Label: "Documents", Href: "#"}, {Label: "carrel", Href: "#"}}, Current: "attachments"}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Upload", Menu: true, Icon: "up", Kind: w.ButtonPrimary}, w.Button{Label: "New folder", Icon: "plus"}}}}, "s-files-2", "tc-head", 1),
	shell("crumbs", w.Head{Title: &w.HeadTitle{Text: "2026-07", Size: w.HeadTitle18}, Crumbs: &w.Crumbs{Path: []w.Crumb{{Label: "Documents", Href: "#"}, {Label: "carrel", Href: "#"}}, Current: "2026-07"}, Acts: &w.Acts{Items: []w.ActsPart{w.Seg{Options: []w.SegOption{{Label: "List"}, {Label: "Tiles", On: true}}}}}, Pad: w.HeadPadBottom6}, "s-files-5", "tc-head", 1),
	shell("padded", w.Head{Title: &w.HeadTitle{Text: "2026-08-10-sync-board.png", Size: w.HeadTitle19}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mono{Text: "image/png · 412 KB · changed 2026-08-10"}}}}, Pad: w.HeadPadBottom6}, "s-publish-9", "tc-head", 1),
	shell("photo", w.Head{Grouped: true, Photo: &w.HeadPhoto{Initials: "АК", Large: true}, Title: &w.HeadTitle{Text: "Анна Ковалёва", Size: w.HeadTitle19}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Edit", Small: true}, w.Button{Label: "All records", Small: true}}}, Middle: true}, "components-14", "tc-head", 1),
	shell("photo", w.Head{Photo: &w.HeadPhoto{Initials: "АК"}, Title: &w.HeadTitle{Text: "Анна Ковалёва", Size: w.HeadTitle19}, Tools: []w.Button{w.Button{Label: "Edit", Small: true}, w.Button{Label: "All records", Small: true}, w.Button{Icon: "dots", Small: true, Title: "More"}}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Icon: "close", Title: "Close the panel"}}}, Middle: true}, "s-contacts-1", "tc-head", 2),
	shell("avatar", w.Head{Avatar: &w.Avatar{Initials: "АК", Size: w.Avatar44}, Kicker: &w.Rubric{Text: "Timeline", Variant: w.RubricBottom2}, Title: &w.HeadTitle{Text: "Анна Ковалёва"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Издательство «Слово» · Выпускающий редактор "), w.Badge{Label: "in 2 books", Tone: w.BadgeLinked}}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Write a note", Icon: "quill", Kind: w.ButtonPrimary}, w.Button{Label: "New event", Icon: "cal"}, w.Button{Label: "New task", Icon: "task"}, w.Button{Label: "Print", Icon: "print"}}}, Pad: w.HeadPadBottom10}, "s-contacts-7", "tc-head", 1),
	shell("photo", w.Head{Photo: &w.HeadPhoto{Initials: "АК", Large: true}, Kicker: &w.Rubric{Text: "Card"}, Tools: []w.Button{w.Button{Label: "Open the card", Small: true}, w.Button{Label: "Edit", Small: true}, w.Button{Icon: "dots", Small: true, Title: "More"}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Icon: "close", Title: "Close the panel"}}}, Middle: true, Pad: w.HeadPadBottom10}, "s-contacts-7", "tc-head", 2),
	shell("acts", w.Head{Title: &w.HeadTitle{Text: "Собрать чемодан"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Дела · "), w.Span{Face: w.FaceDanger, Text: "overdue by 3 days"}}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Edit", Icon: "pencil", Kind: w.ButtonPrimary}, w.Button{Label: "Mark done"}, w.Button{Label: "Print", Icon: "print"}, w.Button{Icon: "dots", Title: "More"}}}}, "s-tasks-2", "tc-head", 1),
	shell("middle", w.Head{Title: &w.HeadTitle{Text: "Заметка со встречи 10.08", Size: w.HeadTitle19}, Tools: []w.Button{w.Button{Label: "Edit", Small: true}, w.Button{Icon: "expand", Small: true, Title: "Open the note page"}, w.Button{Icon: "dots", Small: true, Title: "More"}}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Записи · "), w.Mono{Text: "10 Aug 2026 11:05"}}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Icon: "close", Title: "Close the panel"}}}, Middle: true}, "s-notes-1", "tc-head", 2),
	shell("padded", w.Head{Title: &w.HeadTitle{Text: "Selected collections", Size: w.HeadTitle18}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 of 3 storages on 2 servers")}}}, Pad: w.HeadPadBottom8}, "s-files-6", "tc-head", 1),
	shell("padded", w.Head{Title: &w.HeadTitle{Text: "Connections", Size: w.HeadTitle19}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 servers · 12 collections · 2 subscriptions")}}}, Pad: w.HeadPadBottom8}, "s-profile-3", "tc-head", 1),
	shell("padded", w.Head{Title: &w.HeadTitle{Text: "Connections", Size: w.HeadTitle19}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 servers · 12 collections · 2 subscriptions")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Badge{Label: "credentials unlocked · 4:12"}, w.Button{Label: "Lock now", Small: true}}}, Pad: w.HeadPadBottom8}, "s-profile-4", "tc-head", 1),
	shell("acts", w.Head{Title: &w.HeadTitle{Text: "davloom", Badge: &w.Badge{Label: "local", Tone: w.BadgeLocal, Title: "Local instance: may be unreachable from outside"}}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("One address for every device instead of an account for every server · "), w.Span{Face: w.FaceDanger, Text: "a phone on mobile data will not reach a local address"}}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Add device", Icon: "plus", Kind: w.ButtonPrimary}}}}, "s-davloom-1", "tc-head", 1),
	shell("plain", w.Head{Title: &w.HeadTitle{Text: "Published links", Badge: &w.Badge{Label: "local", Tone: w.BadgeLocal, Title: "Local instance: may be unreachable from outside"}}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Read-only, one object or one collection per link · "), w.Span{Face: w.FaceDanger, Text: "this instance is local: the link opens only from your network"}}}}}, "s-publish-1", "tc-head", 1),
	shell("padded", w.Head{Title: &w.HeadTitle{Text: "Личный"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("10 – 31 August 2026 · the page is read from the server on every visit")}}}, Pad: w.HeadPadBottom8}, "s-publish-6", "tc-head", 1),
	shell("crumbs", w.Head{Title: &w.HeadTitle{Text: "Материалы к релизу", Size: w.HeadTitle19}, Crumbs: &w.Crumbs{Current: "Материалы к релизу", Aside: &w.Dim{Parts: []w.Inline{w.Text("· 4 files · 6.6 MB")}}}, Pad: w.HeadPadBottom8}, "s-publish-12", "tc-head", 1),
	shell("avatar", w.Head{Avatar: &w.Avatar{Initials: "АК", Size: w.Avatar44}, Kicker: &w.Rubric{Text: "Mentions", Variant: w.RubricBottom2}, Title: &w.HeadTitle{Text: "Анна Ковалёва"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Where she is named, not where she takes part · the timeline has that")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Write a note", Icon: "quill", Kind: w.ButtonPrimary}, w.Button{Label: "Timeline", Icon: "all"}}}, Pad: w.HeadPadBottom10}, "mentions-2", "tc-head", 1),
	shell("acts", w.Head{Title: &w.HeadTitle{Text: "Команда", Badge: &w.Badge{Label: "read-only"}}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("radicale · vps · the server granted no write access")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Export", Icon: "down"}, w.Button{Label: "Print", Icon: "print"}}}}, "states-4", "tc-head", 1),
	shell("plain", w.Head{Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti}, Text: "Combined calendar"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("shown from this session’s memory")}}}}, "states-5", "tc-head", 1),
	shell("padded", w.Head{Title: &w.HeadTitle{Text: "Синк-ап по релизу", Size: w.HeadTitle19}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, w.Text("nextcloud · Личный")}}}, Pad: w.HeadPadBottom8}, "phone-3", "tc-head", 1),
	shell("padded", w.Head{Title: &w.HeadTitle{Text: "Application", Size: w.HeadTitle19}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Account and collection settings are inside, in Settings")}}}, Pad: w.HeadPadBottom6}, "desktop-13", "tc-head", 1),

	// tc-bar
	shell("selection", w.Bar{Items: []w.BarPart{w.Mono{Text: "412 KB", Variant: w.MonoMuted}, w.BarSep{}, w.Button{Folds: true, Label: "Download", Icon: "down", Title: "Download"}, w.Button{Folds: true, Label: "Rename", Icon: "pencil", Title: "Rename"}, w.Button{Folds: true, Label: "Delete", Icon: "trash", Kind: w.ButtonDangerSolid, Title: "Delete"}}, Count: &w.BarCount{Text: "1 selected"}, Right: &w.BarRight{Items: []w.BarPart{w.Button{Label: "Clear selection", Small: true}}}, Variant: w.BarSelection}, "components-16", "tc-bar", 1),
	shell("more", w.Bar{Items: []w.BarPart{w.Input{Kind: w.InputSearch, Placeholder: "Filter this folder", Variant: w.InputW190, Label: "Filter this folder"}, w.Seg{Options: []w.SegOption{{Label: "Name", On: true}, {Label: "Size"}, {Label: "Changed"}}, Secondary: true}, w.Button{Label: "Columns", Icon: "cols"}}, More: &w.Button{Icon: "dots", Title: "More"}, Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "14 items · 38.2 MB"}}, Secondary: true}}, "components-16", "tc-bar", 2),
	shell("right", w.Bar{Items: []w.BarPart{w.Input{Kind: w.InputSearch, Placeholder: "Filter by word", Variant: w.InputH26W200, Label: "Filter by word"}, w.Seg{Options: []w.SegOption{{Label: "Newest", On: true}, {Label: "Oldest"}, {Label: "Title"}}}, w.Button{Label: "Columns", Icon: "cols"}}, Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "86"}}, Secondary: true}}, "s-notes-1", "tc-bar", 1),
	shell("plain", w.Bar{Items: []w.BarPart{w.Input{Kind: w.InputSearch, Placeholder: "Filter by login", Variant: w.InputH26W200, Label: "Filter by login"}, w.Seg{Options: []w.SegOption{{Label: "All", On: true}, {Label: "Administrators"}, {Label: "Disabled"}}}, w.Button{Folds: true, Label: "Columns", Icon: "cols", Title: "Columns"}}}, "s-admin-2", "tc-bar", 1),
	shell("selection", w.Bar{Items: []w.BarPart{w.Mono{Text: "2.8 MB", Variant: w.MonoMuted}, w.BarSep{}, w.Button{Folds: true, Label: "Download", Icon: "down", Title: "Download"}, w.Button{Folds: true, Label: "Rename", Icon: "pencil", Title: "Rename"}, w.Button{Folds: true, Label: "Move", Icon: "move", Title: "Move"}, w.Button{Folds: true, Label: "Delete", Icon: "trash", Kind: w.ButtonDangerSolid, Title: "Delete"}}, Count: &w.BarCount{Text: "1 selected"}, Variant: w.BarSelection}, "s-files-5", "tc-bar", 1),
	shell("right", w.Bar{Items: []w.BarPart{w.Filter{Query: w.Query{Value: "Filter the loaded cards", Label: "Filter the loaded cards"}, Waiting: true}, w.Seg{Options: []w.SegOption{{Label: "A–Z", On: true}, {Label: "Recently changed"}}, Secondary: true}, w.Button{Label: "Columns", Icon: "cols"}}, Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "248"}}, Secondary: true}}, "s-io-1", "tc-bar", 1),
	shell("wide", w.Bar{Items: []w.BarPart{w.Range{From: w.Input{Kind: w.InputDate, Value: "2026-01-01", Mono: true, Label: "Date"}, To: w.Input{Value: "today", Mono: true, Label: "Date"}, Show: w.Button{Label: "Show"}}, w.Seg{Options: []w.SegOption{{Label: "All 24", On: true}, {Label: "Events 14"}, {Label: "Tasks 2"}, {Label: "Notes 4"}, {Label: "Files 4"}}, Secondary: true}, w.Button{Label: "Columns", Icon: "cols", Secondary: true}}, More: &w.Button{Icon: "dots", Title: "More"}, Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "newest first"}}, Secondary: true}, Variant: w.BarWide}, "s-contacts-7", "tc-bar", 1),
	shell("selection", w.Bar{Items: []w.BarPart{w.BarSep{}, w.Button{Label: "Restore into", Icon: "back"}, w.Button{Label: "Skip"}}, Count: &w.BarCount{Text: "3 selected"}, Right: &w.BarRight{Items: []w.BarPart{w.Button{Label: "Select all", Small: true}}}, Variant: w.BarSelection}, "s-backup-8", "tc-bar", 1),
	shell("selection", w.Bar{Items: []w.BarPart{w.Button{Icon: "down", Title: "More"}, w.Button{Icon: "move", Title: "More"}, w.Button{Icon: "trash", Kind: w.ButtonDangerSolid, Title: "Delete"}, w.Button{Icon: "close", Title: "More", Place: w.Pushed}}, Count: &w.BarCount{Text: "2 selected", Small: true}, Variant: w.BarSelection}, "phone-6", "tc-bar", 1),

	// tc-main
	shell("cards", w.Main{Items: []w.MainPart{
		w.MainCards{Cards: []w.Card{w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Temporary password", Items: []w.RubricPart{w.Mono{Text: "visible only now", Variant: w.MonoTightDanger}}},
			w.Msg{Parts: []w.MsgPart{w.Text("The user "), w.Strong("viktor"), w.Text(" is created. Hand this over; until they change it, you can sign in as them.")}, Space: w.MsgFlushBottom12},
			w.Diag{Text: "login     viktor\npassword  9f2-brick-lamp", Variant: w.DiagInk},
			w.CardRow{Kind: w.CardTop8, Gap: 8, Items: []w.CardRowPart{w.Button{Label: "Copy", Icon: "dup", Kind: w.ButtonPrimary}, w.Button{Label: "Done"}}}}, Variant: w.CardChosen}}}}}, "s-admin-17", "tc-main", 1),
	shell("plain", w.Main{Items: []w.MainPart{
		w.Banner{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}, Badge: &w.Badge{Label: "read-only"}, End: w.Line{Gap: w.LineGap6, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Download .ics", Icon: "down"}}}, Wide: true},
		w.Head{Title: &w.HeadTitle{Text: "Синк-ап по релизу"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Monday 10 August 2026, 10:00 – 11:00 · Переговорная 2")}}}, Pad: w.HeadPadBottom8}}}, "s-publish-16", "tc-main", 1),
	shell("plain", w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Text: "Анна Ковалёва"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Save", Kind: w.ButtonPrimary}}}},
		w.Form{},
		w.FormFoot{}}}, "s-contacts-3", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti}, Text: "Combined notebook"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 of 3 notebooks · 86 notes · VJOURNAL")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "New note", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}, w.Button{Label: "Import", Icon: "up"}, w.Button{Label: "Export", Icon: "down"}}}},
		w.Bar{Items: []w.BarPart{w.Input{Kind: w.InputSearch, Placeholder: "Filter by word", Variant: w.InputH26W200, Label: "Filter by word"}, w.Seg{Options: []w.SegOption{{Label: "Newest", On: true}, {Label: "Oldest"}, {Label: "Title"}}}, w.Button{Label: "Columns", Icon: "cols"}}, Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "86"}}, Secondary: true}},
		w.ListGroup{Title: w.Rubric{Text: "August 2026"}, Count: &w.Num{Value: 9}},
		w.List{Rows: []w.ListRow{
			w.ExcerptRow{Mark: w.Mark{Colour: "#4A6B52"}, Title: []w.Inline{w.Text("Заметка со встречи 10.08")}, Strong: true, Excerpt: "Договорились заморозить ветку в среду; письмо пользователям пишет Анна…", Tags: []w.RowTag{w.Tag{Label: "встречи"}}, Date: "10 Aug", Source: "nextcloud · Записи", On: true},
			w.ExcerptRow{Mark: w.Mark{Colour: "#4A6B52"}, Title: []w.Inline{w.Text("Что почитать про CRDT")}, TitleBadge: &w.Badge{Label: "linked ×2", Tone: w.BadgeLinked}, Excerpt: "Кляйн, глава 5; статья Шапиро 2011; демка automerge…", Tags: []w.RowTag{w.Tag{Label: "книги"}}, Date: "8 Aug", Source: "2 sources", SourceHint: "nextcloud · Записи, radicale · vps · Дневник"},
			w.ExcerptRow{Mark: w.Mark{Colour: "#6B4E7A"}, Title: []w.Inline{w.Text("Про дорогу до дачи")}, Excerpt: "Съезд после заправки, потом грунтовка два километра…", Tags: []w.RowTag{w.Badge{Label: "linked", Tone: w.BadgeLinked, Icon: "link"}}, Date: "3 Aug", Source: "nextcloud · Дневник"}}}}}, "s-notes-1", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
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
		w.Empty{Parts: []w.EmptyPart{w.Text("Loading more…")}, Variant: w.EmptyDusty}}}, "s-contacts-1", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti}, Text: "Combined calendar"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("5 of 7 calendars · 3 accounts · 10 – 16 August 2026 · "), w.Mono{Text: "Europe/Moscow"}}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Icon: "plus", Kind: w.ButtonPrimary}, w.Button{Label: "Import", Icon: "up"}, w.Button{Label: "Export", Icon: "down"}, w.Button{Label: "Print", Icon: "print"}}}},
		w.Bar{Items: []w.BarPart{w.Seg{Options: []w.SegOption{{Label: "Day"}, {Label: "Week", On: true}, {Label: "Month"}, {Label: "Range"}}}, w.Range{From: w.Input{Kind: w.InputDate, Value: "2026-08-10", Mono: true, Label: "Date"}, To: w.Input{Kind: w.InputDate, Value: "2026-08-16", Mono: true, Label: "Date"}, Show: w.Button{Label: "Show"}, Secondary: true}, w.BarSep{Secondary: true}, w.Check{On: true, Parts: []w.CheckPart{w.Text("Events")}}, w.Check{Parts: []w.CheckPart{w.Text("Tasks due")}}, w.Check{Parts: []w.CheckPart{w.Text("Notes")}}, w.Drop{Open: true, Trigger: w.Button{Icon: "dots", Title: "More"}, More: true, Menu: w.Menu{Items: []w.MenuPart{
			w.Range{From: w.Input{Kind: w.InputDate, Value: "2026-08-10", Mono: true, Label: "Date"}, To: w.Input{Kind: w.InputDate, Value: "2026-08-16", Mono: true, Label: "Date"}, Show: w.Button{Label: "Show"}, Wrap: true},
			w.MenuSep{},
			w.MenuItem{Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Compact rows"}}, Pop: w.MenuUnderEnd, Inline: true}}}, Right: &w.BarRight{Items: []w.BarPart{w.Check{Parts: []w.CheckPart{w.Text("Compact rows")}}}, Secondary: true}, Variant: w.BarWide},
		w.Runbar{Meter: w.Meter{Percent: 83}, Parts: []w.RunbarPart{w.Text("Polled 5 of 6 sources · 14 events · "), w.Span{Face: w.FaceDanger, Text: "baikal unavailable"}, w.Mono{Text: "from cache · 2 min ago", Place: w.Pushed}}},
		w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August"}, Count: &w.Num{Value: 4}},
		w.List{Rows: []w.ListRow{
			w.TimeRow{Mark: w.Mark{Colour: "#4A6B52"}, Time: "10:00", Title: []w.Inline{w.Text("Синк-ап по релизу")}, Strong: true, Detail: "Переговорная 2", Caption: &w.Badge{Label: "4 attendees"}, Source: "nextcloud · Личный", On: true},
			w.TimeRow{Mark: w.Mark{Colour: "#3F6E8C"}, Time: "all day", Title: []w.Inline{w.Text("Дежурство: первая линия")}, Detail: "—", Caption: &w.Badge{Label: "repeats weekly"}, Source: "nextcloud · Работа"},
			w.TimeRow{Mark: w.Mark{Colour: "#8C6239"}, Time: "all day", Title: []w.Inline{w.Text("Дежурство: первая линия")}, Detail: "—", Caption: &w.Badge{Label: "repeats weekly"}, Source: "radicale · vps · Дежурства"},
			w.TimeRow{Mark: w.Mark{Colour: "#3F6E8C"}, Time: "15:00", Title: []w.Inline{w.Text("Созвон с подрядчиком")}, Detail: "Jitsi", Caption: &w.Badge{Label: "tentative", Tone: w.BadgeAlert}, Source: "nextcloud · Работа"}}},
		w.ListGroup{Title: w.Rubric{Text: "Tuesday · 11 August"}, Count: &w.Num{Value: 2}},
		w.List{Rows: []w.ListRow{
			w.TimeRow{Mark: w.Mark{Colour: "#6B4E7A"}, Time: "09:30", Title: []w.Inline{w.Text("Ревью архитектуры")}, Detail: "Переговорная 1", Caption: &w.Badge{Label: "2 attendees"}, Source: "radicale · vps · Проекты"},
			w.TimeRow{Mark: w.Mark{Colour: "#4A6B52"}, Time: "18:00", Title: []w.Inline{w.Text("Забрать документы")}, Detail: "МФЦ, Мясницкая 13", Source: "nextcloud · Личный"}}},
		w.ListGroup{Title: w.Rubric{Text: "Wednesday · 12 August"}, Count: &w.Num{Value: 0}},
		w.Empty{Parts: []w.EmptyPart{w.Text("Nothing this day.")}, Variant: w.EmptySmall}}}, "s-agenda-1", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Head{Meta: &w.HeadSub{Parts: []w.HeadSubPart{w.Mark{Colour: "#4A6B52"}, w.Text("nextcloud · Записи · "), w.Mono{Text: "10 Aug 2026 11:05"}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Edit", Icon: "pencil", Kind: w.ButtonPrimary}, w.Button{Label: "Export Markdown"}, w.Button{Label: "Print", Icon: "print"}, w.Button{Label: "Source"}, w.Seg{Options: []w.SegOption{{Label: "Full width", On: true}, {Label: "Reading width"}}}, w.Button{Icon: "panel", Title: "Hide the side column"}, w.Button{Icon: "dots", Title: "More"}}}, Pad: w.HeadRule},
		w.DocLayout{Items: []w.DocLayoutPart{w.Doc{Items: []w.DocPart{
			w.DocTitle{Text: "Заметка со встречи 10.08"},
			w.Prose{Parts: []w.Inline{w.Text("Договорились заморозить ветку в среду. Письмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.")}, Variant: w.ProseLead},
			w.DocTasks{Title: w.Rubric{Text: "Что делаем", Variant: w.RubricBottom8}, Items: []w.DocTask{{Tick: w.Tick{Variant: w.TickNudge, Label: "проверить changelog"}, Text: "проверить changelog"}, {Tick: w.Tick{Done: true, Variant: w.TickNudge, Label: "попросить у Артёма скриншоты"}, Text: "попросить у Артёма скриншоты"}}},
			w.Prose{Parts: []w.Inline{w.Text("Перенос старых вложений в новую папку — Виктор посмотрит, сколько их вообще, и скажет в четверг.")}, Variant: w.ProseLeadAfter}}}, w.Docmeta{Sections: []w.DocmetaSec{
			{Title: w.Rubric{Text: "Tags"}, Items: []w.DocmetaPart{w.Line{Block: true, Gap: w.LineGap4, Items: []w.LinePart{w.Tag{Label: "встречи"}, w.Tag{Label: "релиз"}}}}},
			{Title: w.Rubric{Text: "Related"}, Items: []w.DocmetaPart{w.DocRefs{Items: []w.DocRef{{Name: "Синк-ап по релизу", What: w.Dim{Parts: []w.Inline{w.Text("· event · 10 Aug")}}}, {Name: "Дописать раздел про кэш", What: w.Dim{Parts: []w.Inline{w.Text("· task")}}}}}}},
			{Title: w.Rubric{Text: "Attachments", Items: []w.RubricPart{w.Button{Label: "Attach", Menu: true, Small: true}}}, Items: []w.DocmetaPart{w.DocAttach{Name: "2026-08-10-sync-board.png", Size: w.Dim{Parts: []w.Inline{w.Text("· 412 KB")}}, Where: w.Badge{Label: "in the note"}}, w.DocAttach{Name: "запись-встречи.m4a", Size: w.Dim{Parts: []w.Inline{w.Text("· 24 MB")}}, Where: w.Badge{Label: "Documents / carrel"}}}},
			{Title: w.Rubric{Text: "The same note in other notebooks", Items: []w.RubricPart{w.Badge{Label: "1", Tone: w.BadgeLinked}}}, Items: []w.DocmetaPart{w.DocRefs{Items: []w.DocRef{{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, Name: "nextcloud · Дневник", Href: "#", What: w.Dim{Parts: []w.Inline{w.Text("· Text differ")}}}}}}},
			{Title: w.Rubric{Text: "Source"}, Items: []w.DocmetaPart{w.Diag{Text: "UID 9e31-40b7-c012 Read 2 min ago"}}}}, Danger: &w.Danger{Button: w.Button{Label: "Delete note", Kind: w.ButtonDangerSolid}, Space: w.DangerAtEnd}}}}}}, "s-notes-2", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Text: "Duplicates"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("361 records compared · 4 to look at")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Seg{Options: []w.SegOption{{Label: "All 4", On: true}, {Label: "Contacts 1"}, {Label: "Events 1"}, {Label: "Tasks 1"}, {Label: "Notes 1"}}}, w.Field{Label: "Match threshold", Width: w.FieldInline, Parts: []w.FieldPart{w.Input{Value: "60", Mono: true, Variant: w.InputW52, Label: "Match threshold"}, w.Button{Label: "Weights", Small: true}}}}, Stack: true}},
		w.Msg{Parts: []w.MsgPart{w.Text("Linking changes nothing on any server: the records stay where they are and are shown as one. Merging on the server is a separate, destructive step you are asked to confirm.")}, Tone: w.MsgAccent},
		w.MainResults{Items: []w.ResultsPart{w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Артём Белов"},
			w.List{Rows: []w.ListRow{
				w.FaceRow{Mark: w.Mark{Colour: "#4A6B52"}, Avatar: w.Avatar{Initials: "АБ"}, Title: []w.Inline{w.Text("Артём Белов")}, Subtitle: "+7 903 118-22-04", Detail: "a.belov@example.org", Caption: "Тензор", Source: "nextcloud · Контакты", Inset: true},
				w.FaceRow{Mark: w.Mark{Colour: "#6B4E7A"}, Avatar: w.Avatar{Photo: true}, Title: []w.Inline{w.Text("Белов Артём")}, Subtitle: "+7 903 118-22-04", Detail: "belov@team.example", Caption: "—", Source: "radicale · vps · Команда", Inset: true}}, Inset: true},
			w.CardRow{Kind: w.CardTop11, Gap: 22, Items: []w.CardRowPart{w.Field{Label: "Display name", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Organisation", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Photo", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}},
			w.CardRow{Kind: w.CardActs, Gap: 8, Items: []w.CardRowPart{w.Button{Label: "Link them", Kind: w.ButtonPrimary}, w.Button{Label: "Not duplicates"}, w.Button{Label: "Merge on the server", Kind: w.ButtonDanger, Place: w.Pushed}}}}}, w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Дежурство: первая линия"},
			w.List{Rows: []w.ListRow{
				w.TimeRow{Mark: w.Mark{Colour: "#3F6E8C"}, Time: "all day", Title: []w.Inline{w.Text("Дежурство: первая линия")}, Detail: "—", Caption: &w.Badge{Label: "repeats weekly"}, Source: "nextcloud · Работа", Inset: true},
				w.TimeRow{Mark: w.Mark{Colour: "#8C6239"}, Time: "all day", Title: []w.Inline{w.Text("Дежурство: первая линия")}, Detail: "—", Caption: &w.Badge{Label: "repeats weekly"}, Source: "radicale · vps · Дежурства", Inset: true}}, Inset: true},
			w.CardHint{Parts: []w.CardTextPart{w.Text("The same duty in two calendars, different UIDs. Agenda shows it twice; here you decide whether to keep both or write one and delete the other.")}, Line: true},
			w.CardRow{Kind: w.CardActs, Gap: 8, Items: []w.CardRowPart{w.Button{Label: "Link them", Kind: w.ButtonPrimary}, w.Button{Label: "Not duplicates"}, w.Button{Label: "Merge on the server", Kind: w.ButtonDanger, Place: w.Pushed}}}}}, w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Собрать чемодан"},
			w.List{Rows: []w.ListRow{
				w.DueRow{Tick: w.Tick{Label: "Done"}, Mark: w.Mark{Colour: "#4A6B52"}, Title: []w.Inline{w.Text("Собрать чемодан")}, Tags: []w.RowTag{w.Tag{Label: "дом"}}, Due: "12 Aug 18:00", Progress: "—", Source: "nextcloud · Дела", Inset: true},
				w.DueRow{Tick: w.Tick{Label: "Done"}, Mark: w.Mark{Colour: "#6B4E7A"}, Title: []w.Inline{w.Text("Собрать чемодан")}, Due: "12 Aug", Progress: "—", Source: "radicale · vps · Личное", Inset: true}}, Inset: true},
			w.CardRow{Kind: w.CardActs, Gap: 8, Items: []w.CardRowPart{w.Button{Label: "Link them", Kind: w.ButtonPrimary}, w.Button{Label: "Not duplicates"}, w.Button{Label: "Merge on the server", Kind: w.ButtonDanger, Place: w.Pushed}}}}}, w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Что почитать про CRDT"},
			w.List{Rows: []w.ListRow{
				w.ExcerptRow{Mark: w.Mark{Colour: "#4A6B52"}, Title: []w.Inline{w.Text("Что почитать про CRDT")}, Excerpt: "Кляйн, глава 5; статья Шапиро 2011…", Tags: []w.RowTag{w.Tag{Label: "книги"}}, Date: "8 Aug", Source: "nextcloud · Записи", Inset: true},
				w.ExcerptRow{Mark: w.Mark{Colour: "#6B4E7A"}, Title: []w.Inline{w.Text("Что почитать про CRDT")}, Excerpt: "Кляйн, глава 5; статья Шапиро 2011…", Date: "8 Aug", Source: "radicale · vps · Дневник", Inset: true}}, Inset: true},
			w.CardHint{Parts: []w.CardTextPart{w.Text("Same title and the same date: a note copied into a second notebook. Linking keeps both; merging writes one and deletes the other.")}, Line: true},
			w.CardRow{Kind: w.CardActs, Gap: 8, Items: []w.CardRowPart{w.Button{Label: "Link them", Kind: w.ButtonPrimary}, w.Button{Label: "Not duplicates"}, w.Button{Label: "Merge on the server", Kind: w.ButtonDanger, Place: w.Pushed}}}}}, w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Анна Ковалёва"},
			w.List{Rows: []w.ListRow{
				w.FaceRow{Mark: w.Mark{Colour: "#4A6B52"}, Avatar: w.Avatar{Initials: "АК"}, Title: []w.Inline{w.Text("Анна Ковалёва")}, Subtitle: "+7 912 445-10-88", Detail: "a.kovaleva@example.org", Caption: "Издательство «Слово»", Source: "nextcloud · Контакты", Inset: true},
				w.FaceRow{Mark: w.Mark{Colour: "#6B4E7A"}, Avatar: w.Avatar{Photo: true}, Title: []w.Inline{w.Text("Анна Ковалёва")}, Subtitle: "+7 912 445-10-88", Detail: "a.kovaleva@example.org", Caption: "—", Source: "radicale · vps · Команда", Inset: true},
				w.FaceRow{Mark: w.Mark{Colour: "#9A9280"}, Avatar: w.Avatar{Initials: "АК"}, Title: []w.Inline{w.Text("А. Ковалёва")}, Subtitle: "—", Detail: "a.kovaleva@example.org", Caption: "Слово", Source: "baikal · Общий", SourceBadge: &w.Badge{Label: "ro"}, Inset: true}}, Inset: true},
			w.CardRow{Kind: w.CardTop11, Gap: 22, Items: []w.CardRowPart{w.Field{Label: "Photo", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}},
			w.CardRow{Kind: w.CardActs, Gap: 8, Items: []w.CardRowPart{w.Button{Label: "Keep these values", Kind: w.ButtonPrimary}, w.Button{Label: "Unlink"}, w.Button{Label: "Merge on the server", Kind: w.ButtonDanger, Place: w.Pushed}}},
			w.CardText{Kind: w.CardWhy, Parts: []w.CardTextPart{w.Text("Общий is read-only: its record could be neither rewritten nor deleted.")}}}}, w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Сергей Орлов"},
			w.List{Rows: []w.ListRow{
				w.FaceRow{Mark: w.Mark{Colour: "#4A6B52"}, Avatar: w.Avatar{Initials: "СО"}, Title: []w.Inline{w.Text("Сергей Орлов")}, Subtitle: "+7 495 120-45-00", Detail: "s.orlov@slovo.example", Caption: "Издательство «Слово»", Source: "nextcloud · Контакты", Inset: true},
				w.FaceRow{Mark: w.Mark{Colour: "#6B4E7A"}, Avatar: w.Avatar{Initials: "МО"}, Title: []w.Inline{w.Text("Мария Орлова")}, Subtitle: "+7 495 120-45-00", Detail: "m.orlova@slovo.example", Caption: "Издательство «Слово»", Source: "radicale · vps · Команда", Inset: true}}, Inset: true},
			w.CardRow{Kind: w.CardActs, Gap: 8, Items: []w.CardRowPart{w.Button{Label: "Offer again"}}}}}}}}}, "s-dup-1", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Banner{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}, Badge: &w.Badge{Label: "read-only"}, End: w.Line{Gap: w.LineGap6, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Download .ics", Icon: "down"}}}, Wide: true},
		w.Strip{Label: "Subscribe in your own client", Items: []w.StripPart{w.Input{Value: "webcal://carrel.example/p/7f3ac9/personal.ics", Mono: true, Variant: w.InputLineFill, Label: "Address"}, w.Button{Label: "Copy", Small: true}, w.Button{Label: "Open in calendar app", Small: true}, w.Dim{Parts: []w.Inline{w.Text("read-only, updates on its own")}, Size: w.Dim11}}},
		w.Head{Title: &w.HeadTitle{Text: "Личный"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("10 – 31 August 2026 · the page is read from the server on every visit")}}}, Pad: w.HeadPadBottom8},
		w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August"}, Count: &w.Num{Value: 2}},
		w.List{Rows: []w.ListRow{
			w.TimeRow{Mark: w.Mark{Colour: "#4A6B52"}, Time: "10:00", Title: []w.Inline{w.Text("Синк-ап по релизу")}, Detail: "Переговорная 2"},
			w.TimeRow{Mark: w.Mark{Colour: "#4A6B52"}, Time: "18:00", Title: []w.Inline{w.Text("Забрать документы")}, Detail: "МФЦ, Мясницкая 13"}}},
		w.ListGroup{Title: w.Rubric{Text: "Tuesday · 11 August"}, Count: &w.Num{Value: 1}},
		w.List{Rows: []w.ListRow{
			w.TimeRow{Mark: w.Mark{Colour: "#4A6B52"}, Time: "09:30", Title: []w.Inline{w.Text("Ревью архитектуры")}, Detail: "Переговорная 1"}}},
		w.MainFoot{Text: "This link works until 8 September 2026. The owner can revoke it at any moment.", Roomy: true}}}, "s-publish-6", "tc-main", 1),
	shell("zone", w.Main{Items: []w.MainPart{
		w.Zone{Title: w.Rubric{Text: "Content"}, Note: w.Mono{Text: "fluid · min 0", Variant: w.Mono11Top}}}, Variant: w.MainZone}, "shell-1", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Head{Meta: &w.HeadSub{Parts: []w.HeadSubPart{w.Mark{Colour: "#4A6B52"}, w.Text("nextcloud · Записи · "), w.Mono{Text: "10 Aug 2026 11:05"}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Edit", Icon: "pencil", Kind: w.ButtonPrimary}, w.Button{Label: "Export Markdown"}, w.Button{Label: "Print", Icon: "print"}, w.Button{Label: "Source"}, w.Seg{Options: []w.SegOption{{Label: "Full width"}, {Label: "Reading width", On: true}}}, w.Button{Icon: "panel", Title: "Hide the side column"}, w.Button{Icon: "dots", Title: "More"}}}, Pad: w.HeadRule},
		w.DocLayout{Items: []w.DocLayoutPart{w.Doc{Items: []w.DocPart{
			w.DocTitle{Text: "Заметка со встречи 10.08"},
			w.Prose{Parts: []w.Inline{w.Text("Договорились заморозить ветку в среду. Письмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.")}, Variant: w.ProseLead},
			w.DocTasks{Title: w.Rubric{Text: "Что делаем", Variant: w.RubricBottom8}, Items: []w.DocTask{{Tick: w.Tick{Variant: w.TickNudge, Label: "проверить changelog"}, Text: "проверить changelog"}, {Tick: w.Tick{Done: true, Variant: w.TickNudge, Label: "попросить у Артёма скриншоты"}, Text: "попросить у Артёма скриншоты"}}},
			w.Prose{Parts: []w.Inline{w.Text("Перенос старых вложений в новую папку — Виктор посмотрит, сколько их вообще, и скажет в четверг.")}, Variant: w.ProseLeadAfter}}}, w.Docmeta{Sections: []w.DocmetaSec{
			{Title: w.Rubric{Text: "Tags"}, Items: []w.DocmetaPart{w.Line{Block: true, Gap: w.LineGap4, Items: []w.LinePart{w.Tag{Label: "встречи"}, w.Tag{Label: "релиз"}}}}},
			{Title: w.Rubric{Text: "Related"}, Items: []w.DocmetaPart{w.DocRefs{Items: []w.DocRef{{Name: "Синк-ап по релизу", What: w.Dim{Parts: []w.Inline{w.Text("· event · 10 Aug")}}}, {Name: "Дописать раздел про кэш", What: w.Dim{Parts: []w.Inline{w.Text("· task")}}}}}}},
			{Title: w.Rubric{Text: "Attachments", Items: []w.RubricPart{w.Button{Label: "Attach", Menu: true, Small: true}}}, Items: []w.DocmetaPart{w.DocAttach{Name: "2026-08-10-sync-board.png", Size: w.Dim{Parts: []w.Inline{w.Text("· 412 KB")}}, Where: w.Badge{Label: "in the note"}}, w.DocAttach{Name: "запись-встречи.m4a", Size: w.Dim{Parts: []w.Inline{w.Text("· 24 MB")}}, Where: w.Badge{Label: "Documents / carrel"}}}},
			{Title: w.Rubric{Text: "Source"}, Items: []w.DocmetaPart{w.Diag{Text: "UID 9e31-40b7-c012 Read 2 min ago"}}}}, Danger: &w.Danger{Button: w.Button{Label: "Delete note", Kind: w.ButtonDangerSolid}, Space: w.DangerAtEnd}}}, View: w.DocViewMeasure}}}, "s-notes-3", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Docbar{Items: []w.DocbarPart{w.Button{Label: "Save", Kind: w.ButtonPrimary}, w.Button{Label: "Cancel"}, w.BarSep{}, w.Button{Title: "More"}, w.Button{Label: "I", Variant: w.ButtonItalic}, w.Button{Label: "H2"}, w.Button{Label: "• list"}, w.Button{Folds: true, Label: "task", Icon: "task", Title: "Task"}, w.Button{Icon: "link", Title: "More"}, w.Button{Folds: true, Label: "Attach", Menu: true, Icon: "up", Title: "Attach"}}, More: &w.Button{Icon: "dots", Title: "More"}, End: []w.DocbarEnd{w.Check{Parts: []w.CheckPart{w.Text("Auto-save")}}, w.DraftNote{Text: "Draft in this window ·", Time: w.Mono{Text: "11:42"}, Secondary: true}, w.Kbd{Keys: "Ctrl+S", Secondary: true}, w.Seg{Options: []w.SegOption{{Label: "Full width", On: true}, {Label: "Reading width"}}, Secondary: true}, w.Button{Folds: true, Label: "Focus", Icon: "focus", Variant: w.ButtonOnCard, Title: "Focus"}}},
		w.DocLayout{Items: []w.DocLayoutPart{w.Doc{Items: []w.DocPart{
			w.Area{Text: "Заметка со встречи 10.08", Variant: w.AreaTitle, Label: "Title"},
			w.DocFields{Items: []w.Field{w.Field{Label: "Notebook", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Date", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Time", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Tags", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Related to", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}},
			w.Area{Text: noteBody, Variant: w.AreaBody, Focus: true, Label: "Text"},
			w.DocNote{Parts: []w.Inline{w.Text("Properties Carrel does not know — "), w.Mono{Text: "X-JTX-COLOR"}, w.Text(" and two more — are kept unchanged and not shown here.")}}}}}, View: w.DocViewFocus}}}, "s-notes-4", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Text: "attachments"}, Crumbs: &w.Crumbs{Path: []w.Crumb{{Label: "Documents", Href: "#"}, {Label: "carrel", Href: "#"}}, Current: "attachments"}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Upload", Menu: true, Icon: "up", Kind: w.ButtonPrimary}, w.Button{Label: "New folder", Icon: "plus"}}}},
		w.Bar{Items: []w.BarPart{w.Mono{Text: "4.4 MB", Variant: w.MonoMuted}, w.BarSep{}, w.Button{Folds: true, Label: "Download", Icon: "down", Title: "Download"}, w.Button{Folds: true, Label: "Rename", Icon: "pencil", Title: "Rename"}, w.Button{Folds: true, Label: "Move", Icon: "move", Title: "Move"}, w.Button{Folds: true, Label: "Copy", Icon: "dup", Title: "Copy"}, w.Button{Folds: true, Label: "Delete", Icon: "trash", Kind: w.ButtonDangerSolid, Title: "Delete"}}, Count: &w.BarCount{Text: "3 selected"}, Right: &w.BarRight{Items: []w.BarPart{w.Button{Label: "Clear selection", Small: true}}}, Variant: w.BarSelection},
		w.Table{Cols: []w.TableCol{{Width: w.ColTick, Tick: ptr(w.TickedPart)}, {Label: "Name"}, {Label: "Size", Width: w.ColSize, Rank: w.Rank3}, {Label: "Type", Width: w.ColType, Rank: w.Rank4}, {Label: "Changed", Width: w.ColTime}}, Rows: []w.TableRow{
			{Cells: []w.TableCell{{Pick: ptr(w.TickedOn)}, {Parts: []w.CellPart{w.TableFile{Icon: "img", Name: "2026-08-10-sync-board.png"}}}, {Parts: []w.CellPart{w.Text("412 KB")}, Mono: true}, {Parts: []w.CellPart{w.Text("image/png")}, Mono: true}, {Parts: []w.CellPart{w.Text("2026-08-10 11:07")}, Mono: true}}, State: w.RowPicked},
			{Cells: []w.TableCell{{Pick: ptr(w.TickedOn)}, {Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "release-checklist.md"}}}, {Parts: []w.CellPart{w.Text("3 KB")}, Mono: true}, {Parts: []w.CellPart{w.Text("text/markdown")}, Mono: true}, {Parts: []w.CellPart{w.Text("2026-08-09 18:44")}, Mono: true}}, State: w.RowPicked},
			{Cells: []w.TableCell{{Pick: ptr(w.TickedOn)}, {Parts: []w.CellPart{w.TableFile{Icon: "img", Name: "2026-08-04-whiteboard.jpg"}}}, {Parts: []w.CellPart{w.Text("2.8 MB")}, Mono: true}, {Parts: []w.CellPart{w.Text("image/jpeg")}, Mono: true}, {Parts: []w.CellPart{w.Text("2026-08-04 17:03")}, Mono: true}}, State: w.RowPicked},
			{Cells: []w.TableCell{{Pick: ptr(w.TickedOff)}, {Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "смета-подрядчик.pdf"}}}, {Parts: []w.CellPart{w.Text("1.2 MB")}, Mono: true}, {Parts: []w.CellPart{w.Text("application/pdf")}, Mono: true}, {Parts: []w.CellPart{w.Text("2026-08-05 14:20")}, Mono: true}}}}},
		w.Queue{Title: "Uploading 2 of 3", Total: w.Mono{Text: "12.4 MB", Variant: w.MonoDusty, Place: w.Pushed}, Toggle: w.Button{Label: "Hide", Small: true}, Items: []w.QueueItem{{Name: "отчёт-июль.pdf", Meta: "4.1 MB · done", Meter: &w.Meter{Percent: 100, Track: 100}}, {Name: "плакат.png", Meta: "8.3 MB · 62%", Meter: &w.Meter{Percent: 62, Track: 100}}, {Name: "release-checklist.md", Meta: "name taken", Failed: true, Acts: []w.Button{w.Button{Label: "Keep both", Small: true}, w.Button{Label: "Replace", Small: true}, w.Button{Label: "Skip", Small: true}}}}}}, Variant: w.MainRoomy}, "s-files-2", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Queue{Title: "Uploading 2 of 3", Total: w.Mono{Text: "62%", Variant: w.MonoDusty, Place: w.Pushed}, Toggle: w.Button{Label: "Show", Small: true}, Min: true}}, Variant: w.MainShort}, "s-files-3", "tc-main", 1),
	shell("groups", w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Text: "«ковал»"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("6 of 7 collections · 7 records · 1 source timed out")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Seg{Options: []w.SegOption{{Label: "All", On: true}, {Label: "Contacts"}, {Label: "Events"}, {Label: "Tasks"}, {Label: "Notes"}, {Label: "Files"}}}}}},
		w.Runbar{Meter: w.Meter{Percent: 100}, Parts: []w.RunbarPart{w.Text("Done · 5 of 6 sources answered"), w.Mono{Text: "1.8 s"}, w.Line{Gap: w.LineGap6, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "answered 11:42"}, w.Button{Label: "Search again", Icon: "refresh", Small: true}}}}},
		w.ListGroup{Title: w.Rubric{Text: "Contacts"}, Count: &w.Num{Value: 2}},
		w.List{Rows: []w.ListRow{
			w.FaceRow{Mark: w.Mark{Colour: "#4A6B52"}, Avatar: w.Avatar{Initials: "АК"}, Title: []w.Inline{w.Strong("Ковал"), w.Text("ёва Анна")}, Subtitle: "+7 912 445-10-88", Detail: "a.kovaleva@example.org", Caption: "Издательство «Слово»", Source: "nextcloud · Контакты"},
			w.FaceRow{Mark: w.Mark{Colour: "#6B4E7A"}, Avatar: w.Avatar{Initials: "ПК"}, Title: []w.Inline{w.Strong("Ковал"), w.Text("ёв Пётр")}, Subtitle: "—", Detail: "p.kovalev@example.org", Caption: "—", Source: "radicale · vps · Команда"}}},
		w.ListGroup{Title: w.Rubric{Text: "Events"}, Count: &w.Num{Value: 1}},
		w.List{Rows: []w.ListRow{
			w.TimeRow{Mark: w.Mark{Colour: "#4A6B52"}, Time: "15:00", Title: []w.Inline{w.Text("Звонок с "), w.Strong("Ковал"), w.Text("ёвой")}, Detail: "Zoom", Source: "nextcloud · Личный"}}},
		w.ListGroup{Title: w.Rubric{Text: "Tasks"}, Count: &w.Num{Value: 1}},
		w.List{Rows: []w.ListRow{
			w.DueRow{Tick: w.Tick{Label: "Done"}, Mark: w.Mark{Colour: "#3F6E8C"}, Title: []w.Inline{w.Text("Ответить "), w.Strong("Ковал"), w.Text("ёвой по смете")}, Tags: []w.RowTag{w.Tag{Label: "релиз"}}, Due: "12 Aug", Progress: "40%", Source: "nextcloud · Работа"}}},
		w.ListGroup{Title: w.Rubric{Text: "Notes"}, Count: &w.Num{Value: 1}},
		w.List{Rows: []w.ListRow{
			w.ExcerptRow{Mark: w.Mark{Colour: "#4A6B52"}, Title: []w.Inline{w.Text("Звонок с "), w.Strong("Ковал"), w.Text("ёвой — тезисы")}, Excerpt: "Договорились заморозить ветку в среду; письмо пишет Анна…", Tags: []w.RowTag{w.Tag{Label: "встречи"}}, Date: "10 Aug", Source: "nextcloud · Записи"}}},
		w.ListGroup{Title: w.Rubric{Text: "Files"}, Count: &w.Num{Value: 2}},
		w.List{Rows: []w.ListRow{
			w.SizeRow{Ticked: w.TickedOff, Icon: "doc", Name: []w.Inline{w.Text("договор-"), w.Strong("ковал"), w.Text("ёва.pdf")}, Bytes: "240 KB", Time: "2026-08-04 11:20", Access: "rw", Source: "sftpgo · vps · Documents"},
			w.SizeRow{Ticked: w.TickedOff, Icon: "doc", Name: []w.Inline{w.Strong("Ковал"), w.Text("ёва — правки по обложке.md")}, Bytes: "6 KB", Time: "2026-07-21 16:02", Access: "rw", Source: "sftpgo · vps · Documents"}}},
		w.Msg{Parts: []w.MsgPart{w.Text("Files are searched by name: "), w.Mono{Text: "PROPFIND"}, w.Text(" returns names anyway, and walking a folder costs one request per folder. Carrel does not look inside files — and says so here, so that “nothing found” is not read as “not in the files”.")}}}}, "s-find-1", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Text: "Connections"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Signed in as "), w.Strong("mix"), w.Text(" · 3 servers · 12 collections · 2 subscriptions")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Connect account", Icon: "plus", Kind: w.ButtonPrimary}, w.Button{Label: "Subscribe to a feed", Icon: "link"}}}},
		w.MainResults{Items: []w.ResultsPart{w.Rubric{Text: "DAV servers", Items: []w.RubricPart{w.Mono{Text: "3", Variant: w.MonoTightPlain}}}, w.Card{Items: []w.CardPart{
			w.Rubric{Text: "nextcloud"},
			w.Mono{Text: "https://cloud.example.org/remote.php/dav/ · mix · last answer 2 min ago", Form: w.MonoLine, Variant: w.Mono11_5DustyGap},
			w.Table{Cols: []w.TableCol{{Width: w.ColTickSmall}, {Label: "Collection"}, {Label: "What is in it", Rank: w.Rank3}, {Label: "Objects", Width: w.ColNum, Rank: w.Rank4, Secondary: true}, {Label: "Changed on the server", Width: w.ColChanged}, {Width: w.ColMenu}}, Rows: []w.TableRow{
				{Cells: []w.TableCell{{Pick: ptr(w.TickedOn)}, {Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Личный"}}}, {Parts: []w.CellPart{w.Text("calendar")}, Mono: true}, {Parts: []w.CellPart{w.Text("128")}, Mono: true}, {Parts: []w.CellPart{w.Text("3 h ago")}, Mono: true}, {Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}}, Acts: true}}},
				{Cells: []w.TableCell{{Pick: ptr(w.TickedOn)}, {Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Работа"}}}, {Parts: []w.CellPart{w.Text("calendar")}, Mono: true}, {Parts: []w.CellPart{w.Text("74")}, Mono: true}, {Parts: []w.CellPart{w.Text("yesterday")}, Mono: true}, {Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}}, Acts: true}}},
				{Cells: []w.TableCell{{Pick: ptr(w.TickedOff)}, {Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "Дни рождения"}}}, {Parts: []w.CellPart{w.Text("calendar")}, Mono: true}, {Parts: []w.CellPart{w.Text("31")}, Mono: true}, {Parts: []w.CellPart{w.Text("12 August")}, Mono: true}, {Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}}, Acts: true}}, State: w.RowOff},
				{Cells: []w.TableCell{{Pick: ptr(w.TickedOn)}, {Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Контакты"}}}, {Parts: []w.CellPart{w.Text("address book")}, Mono: true}, {Parts: []w.CellPart{w.Text("248")}, Mono: true}, {Parts: []w.CellPart{w.Text("4 d ago")}, Mono: true}, {Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}}, Acts: true}}},
				{Cells: []w.TableCell{{Pick: ptr(w.TickedOn)}, {Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Записи"}}}, {Parts: []w.CellPart{w.Text("journal")}, Mono: true}, {Parts: []w.CellPart{w.Text("86")}, Mono: true}, {Parts: []w.CellPart{w.Text("2 h ago")}, Mono: true}, {Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}}, Acts: true}}},
				{Cells: []w.TableCell{{Pick: ptr(w.TickedOn)}, {Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Дела"}}}, {Parts: []w.CellPart{w.Text("tasks")}, Mono: true}, {Parts: []w.CellPart{w.Text("30")}, Mono: true}, {Parts: []w.CellPart{w.Text("20 min ago")}, Mono: true}, {Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}}, Acts: true}}}}}}}, w.Card{Items: []w.CardPart{
			w.Rubric{Text: "sftpgo · vps"},
			w.Mono{Text: "https://files.example.org/dav/ · mix · plain WebDAV, no principal · last answer 5 min ago", Form: w.MonoLine, Variant: w.Mono11_5DustyGap},
			w.Table{Cols: []w.TableCol{{Width: w.ColTickSmall}, {Label: "Folder"}, {Label: "What is in it", Rank: w.Rank3}, {Label: "Objects", Width: w.ColNum, Rank: w.Rank4, Secondary: true}, {Label: "Changed on the server", Width: w.ColChanged}, {Width: w.ColMenu}}, Rows: []w.TableRow{
				{Cells: []w.TableCell{{Pick: ptr(w.TickedOn)}, {Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Documents"}}}, {Parts: []w.CellPart{w.Text("files")}, Mono: true}, {Parts: []w.CellPart{w.Text("412")}, Mono: true}, {Parts: []w.CellPart{w.Text("5 min ago")}, Mono: true}, {Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}}, Acts: true}}},
				{Cells: []w.TableCell{{Pick: ptr(w.TickedOff)}, {Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#9A9280"}, Name: "Public"}}}, {Parts: []w.CellPart{w.Text("files "), w.Badge{Label: "ro"}}, Mono: true}, {Parts: []w.CellPart{w.Text("96")}, Mono: true}, {Parts: []w.CellPart{w.Text("3 d ago")}, Mono: true}, {Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}}, Acts: true}}, State: w.RowOff}}}}}, w.Card{Items: []w.CardPart{
			w.Rubric{Text: "baikal"},
			w.Mono{Text: "https://dav.office.example/dav.php/ · m.yenuchenko · last answer refused", Form: w.MonoLine, Variant: w.Mono11_5DustyGap},
			w.Diag{Text: "[well-known] GET /.well-known/caldav → 404 [base] PROPFIND /dav.php/ → 401 Unauthorized WWW-Authenticate: Basic realm=\"BaikalDAV\" [principal] не запрашивался: предыдущий шаг не прошёл"},
			w.CardText{Kind: w.CardAlertText, Parts: []w.CardTextPart{w.Text("The server refused the password. It may have been changed on the server: the one Carrel saved no longer fits.")}}}, Variant: w.CardAlert}, w.Rubric{Text: "Subscriptions · external, read-only", Items: []w.RubricPart{w.Mono{Text: "2", Variant: w.MonoTightPlain}}}, w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Праздники РФ"},
			w.Mono{Text: "https://feeds.example/ru-holidays.ics · no password · read 2 min ago · 34 KB · 0.4 s", Form: w.MonoLine, Variant: w.Mono11_5DustyGap},
			w.Table{Cols: []w.TableCol{{Width: w.ColTickSmall}, {Label: "Lands in"}, {Label: "Objects", Width: w.ColNum, Rank: w.Rank3, Secondary: true}, {Label: "Changed on the server", Width: w.ColChanged, Rank: w.Rank4}}, Rows: []w.TableRow{
				{Cells: []w.TableCell{{Pick: ptr(w.TickedOn)}, {Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#8C6239", Kind: w.MarkExternal}, Name: "Calendar"}}}, {Parts: []w.CellPart{w.Text("118")}, Mono: true}, {Parts: []w.CellPart{w.Text("yesterday")}, Mono: true}}},
				{Cells: []w.TableCell{{Pick: ptr(w.TickedOff)}, {Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#8C6239", Kind: w.MarkExternal}, Name: "Tasks"}}}, {Parts: []w.CellPart{w.Text("6")}, Mono: true}}, State: w.RowOff},
				{Cells: []w.TableCell{{Pick: ptr(w.TickedOn)}, {Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#8C6239", Kind: w.MarkExternal}, Name: "Notes"}}}, {Parts: []w.CellPart{w.Text("2")}, Mono: true}}}}}}}, w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Расписание лиги"},
			w.Mono{Text: "webcal://league.example/season.ics · password saved · last read yesterday", Form: w.MonoLine, Variant: w.Mono11_5DustyGap},
			w.Diag{Text: "[fetch] GET /season.ics → 401 Unauthorized WWW-Authenticate: Basic realm=\"league\" [cache] 46 events from yesterday are still shown, marked stale"},
			w.CardText{Kind: w.CardAlertText, Parts: []w.CardTextPart{w.Text("The feed asks for a password and the saved one no longer fits. A feed that stopped answering must not look like a feed with nothing in it.")}}}, Variant: w.CardAlert}}},
		w.Msg{Parts: []w.MsgPart{w.Text("A feed is read again on every visit to the section, so there is nothing here to synchronise and nothing to refresh. "), w.Strong("Check"), w.Text(" asks whether the address still replies.")}}}}, "s-profile-10", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Text: "davloom", Badge: &w.Badge{Label: "local", Tone: w.BadgeLocal, Title: "Local instance: may be unreachable from outside"}}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("One address for every device instead of an account for every server · "), w.Span{Face: w.FaceDanger, Text: "a phone on mobile data will not reach a local address"}}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Add device", Icon: "plus", Kind: w.ButtonPrimary}}}},
		w.Form{},
		w.Msg{Parts: []w.MsgPart{w.Text("Changing your Carrel password does "), w.Strong("not"), w.Text(" disconnect devices: each has its own key wrapper. To disconnect one, revoke it here — that is the only thing that stops it.")}, Tone: w.MsgAccent},
		w.ListGroup{Title: w.Rubric{Text: "Devices"}, Count: &w.Num{Value: 2}},
		w.Table{Cols: []w.TableCol{{Label: "Device"}, {Label: "What it can see and do"}, {Label: "Created", Rank: w.Rank4, Secondary: true}, {Label: "Last request", Secondary: true}, {Width: w.Col90}}, Rows: []w.TableRow{
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.TableFile{Icon: "device", Name: "Телефон · DAVx5"}}}, {Parts: []w.CellPart{w.CellWrap{Items: []w.CellPart{w.Line{Gap: w.LineGap3, Items: []w.LinePart{w.Mark{Colour: "#4A6B52"}, w.Mark{Colour: "#3F6E8C"}, w.Mark{Colour: "#6B4E7A"}, w.Mark{Colour: "#A05A2C"}}}, w.Mono{Text: "Calendars, Address books", Variant: w.Mono11}, w.Button{Label: "5 collections · Change", Small: true}, w.Seg{Options: []w.SegOption{{Label: "read only"}, {Label: "read and write", On: true}}, Variant: w.SegCompact, Place: w.Pushed}}}}}, {Parts: []w.CellPart{w.Text("2026-06-02")}, Mono: true}, {Parts: []w.CellPart{w.Text("2026-08-15 09:31")}, Mono: true}, {Parts: []w.CellPart{w.Button{Label: "Revoke", Kind: w.ButtonDanger, Small: true}}, Acts: true}}},
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.TableFile{Icon: "device", Name: "Ноутбук · Thunderbird"}}}, {Parts: []w.CellPart{w.CellWrap{Items: []w.CellPart{w.Line{Gap: w.LineGap3, Items: []w.LinePart{w.Mark{Kind: w.MarkMulti}}}, w.Mono{Text: "all kinds", Variant: w.Mono11}, w.Button{Label: "12 collections · Change", Small: true}, w.Seg{Options: []w.SegOption{{Label: "read only", On: true}, {Label: "read and write"}}, Variant: w.SegCompact, Place: w.Pushed}}}}}, {Parts: []w.CellPart{w.Text("2026-06-02")}, Mono: true}, {Parts: []w.CellPart{w.Text("2026-08-15 08:02")}, Mono: true}, {Parts: []w.CellPart{w.Button{Label: "Revoke", Kind: w.ButtonDanger, Small: true}}, Acts: true}}}}, Variant: w.TableFit},
		w.ListGroup{Title: w.Rubric{Text: "What the client sees"}},
		w.MainTree{Tree: w.Tree{Rows: []w.TreeRow{
			{Icon: "file", Label: "/dav/", Code: true, Trail: w.Dim{Parts: []w.Inline{w.Text("root")}, Size: w.Dim11}},
			{Icon: "con", Label: "principal/", Code: true, Indent: w.TreeLevel22},
			{Icon: "cal", Label: "cal/", Code: true, Indent: w.TreeLevel22, Trail: w.Dim{Parts: []w.Inline{w.Text("Calendars · events, tasks, notes")}, Size: w.Dim11}},
			{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "c/k4f2/%2F…%2Fpersonal%2F/", Code: true, Indent: w.TreeLevel44, Trail: w.Dim{Parts: []w.Inline{w.Text("Личный")}, Size: w.Dim11}},
			{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "c/k4f2/%2F…%2Fwork%2F/", Code: true, Indent: w.TreeLevel44, Trail: w.Dim{Parts: []w.Inline{w.Text("Работа")}, Size: w.Dim11}},
			{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "c/r9a1/%2Fmix%2Fprojects%2F/", Code: true, Indent: w.TreeLevel44, Trail: w.Dim{Parts: []w.Inline{w.Text("Проекты")}, Size: w.Dim11}},
			{Icon: "con", Label: "card/", Code: true, Indent: w.TreeLevel22, Trail: w.Dim{Parts: []w.Inline{w.Text("Address books")}, Size: w.Dim11}},
			{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "c/k4f2/%2F…%2Fcontacts%2F/", Code: true, Indent: w.TreeLevel44, Trail: w.Dim{Parts: []w.Inline{w.Text("Контакты")}, Size: w.Dim11}},
			{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "c/r9a1/%2Fmix%2Fteam%2F/", Code: true, Indent: w.TreeLevel44, Trail: w.Dim{Parts: []w.Inline{w.Text("Команда")}, Size: w.Dim11}},
			{Icon: "file", Label: "files/", Code: true, Indent: w.TreeLevel22, Trail: w.Dim{Parts: []w.Inline{w.Text("Files · plain WebDAV")}, Size: w.Dim11}},
			{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "c/s3c8/%2FDocuments%2F/", Code: true, Indent: w.TreeLevel44, Trail: w.Dim{Parts: []w.Inline{w.Text("Documents")}, Size: w.Dim11}}}, Variant: w.TreeNarrow}}}}, "s-davloom-1", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Hero{Symbol: "davloom-tile", Label: "davloom", Title: "davloom", Sub: "DAV proxy for Carrel · one address instead of every configured account", Text: w.Prose{Parts: []w.Inline{w.Text("This is a connection point for clients, not a page to read. Enter the address in anything that speaks CalDAV, CardDAV or WebDAV — DAVx5, Thunderbird, Calendar and Contacts on Apple, a file manager mounting a network drive. The user name is your Carrel one; the password is the device password issued in Carrel.")}, Variant: w.ProseTop}, Fields: &w.Fields{Items: []w.FieldsItem{
			{Term: "Address", Value: []w.ValuePart{w.Text("https://carrel.example/dav/")}, Face: w.ValueMono},
			{Term: "Serves", Value: []w.ValuePart{w.Text("CalDAV · CardDAV · WebDAV")}, Face: w.ValueMono},
			{Term: "Version", Value: []w.ValuePart{w.Text("0.0.1")}, Face: w.ValueMono},
			{Term: "Build", Value: []w.ValuePart{w.Text("786949a")}, Face: w.ValueMono},
			{Term: "Licence", Value: []w.ValuePart{w.Text("AGPL-3.0-or-later")}}}, Variant: w.Fields130Flush}, Acts: []w.Button{w.Button{Label: "Source code", Kind: w.ButtonPrimary}, w.Button{Label: "Open Carrel"}}}}}, "s-davloom-4", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Text: "Backup"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Three jobs · their own sources, destinations and triggers")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "New job", Icon: "plus", Kind: w.ButtonPrimary}, w.Button{Label: "Restore", Icon: "back"}}}},
		w.Table{Cols: []w.TableCol{{Width: w.ColTick}, {Label: "Job"}, {Label: "What it copies", Rank: w.Rank3, Secondary: true}, {Label: "Where to", Rank: w.Rank4, Secondary: true}, {Label: "When"}, {Label: "Keep", Secondary: true}, {Label: "Last backup"}, {Width: w.ColTime}}, Rows: []w.TableRow{
			{Cells: []w.TableCell{{Pick: ptr(w.TickedOn)}, {Parts: []w.CellPart{w.TableFile{Icon: "archive", Name: "Рабочее, ежедневно"}}}, {Parts: []w.CellPart{w.Line{Gap: w.LineGap3, Items: []w.LinePart{w.Mark{Colour: "#3F6E8C"}, w.Mark{Colour: "#6B4E7A"}, w.Mark{Colour: "#4A6B52"}, w.Mono{Text: "3 collections", Variant: w.Mono11Indent}}}}}, {Parts: []w.CellPart{w.Text("sftpgo · vps · Documents"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("/backup/work")}}}, Mono: true}, {Parts: []w.CellPart{w.Text("cron · 03:00"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("external trigger")}}}, Mono: true}, {Parts: []w.CellPart{w.Text("14")}, Mono: true}, {Parts: []w.CellPart{w.Text("2026-08-16 03:10"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("18.4 MB · verified")}}}, Mono: true}, {Parts: []w.CellPart{w.Button{Label: "Run", Small: true}, w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}}, Acts: true}}},
			{Cells: []w.TableCell{{Pick: ptr(w.TickedOn)}, {Parts: []w.CellPart{w.TableFile{Icon: "archive", Name: "Личное на домашний диск"}}}, {Parts: []w.CellPart{w.Line{Gap: w.LineGap3, Items: []w.LinePart{w.Mark{Colour: "#4A6B52"}, w.Mark{Colour: "#A05A2C"}, w.Mono{Text: "2 collections + files", Variant: w.Mono11Indent}}}}}, {Parts: []w.CellPart{w.Text("nextcloud · Files"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("/carrel-backup")}}}, Mono: true}, {Parts: []w.CellPart{w.Text("on sign-in"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("if older than 7 days")}}}, Mono: true}, {Parts: []w.CellPart{w.Text("8")}, Mono: true}, {Parts: []w.CellPart{w.Text("2026-08-14 09:12"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("2.1 GB · verified")}}}, Mono: true}, {Parts: []w.CellPart{w.Button{Label: "Run", Small: true}, w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}}, Acts: true}}},
			{Cells: []w.TableCell{{Pick: ptr(w.TickedOff)}, {Parts: []w.CellPart{w.TableFile{Icon: "archive", Name: "Contacts archive, by hand"}}}, {Parts: []w.CellPart{w.Line{Gap: w.LineGap3, Items: []w.LinePart{w.Mark{Colour: "#A05A2C"}, w.Mono{Text: "1 book", Variant: w.Mono11Indent}}}}}, {Parts: []w.CellPart{w.Text("download in browser"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("nothing is written")}}}, Mono: true}, {Parts: []w.CellPart{w.Text("on demand")}, Mono: true}, {Parts: []w.CellPart{w.Text("—")}, Mono: true}, {Parts: []w.CellPart{w.Text("never")}, Mono: true, Tone: w.CellDusty}, {Parts: []w.CellPart{w.Button{Label: "Run", Small: true}, w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}}, Acts: true}}, State: w.RowFaded}}, Variant: w.TableWide},
		w.MainHint{Text: "verified — the archive was read back from where it was written and the checksum matched. Written and not read back is not verified, and is marked red."},
		w.ListGroup{Title: w.Rubric{Text: "History, all jobs"}, Count: &w.Num{Value: 12}},
		w.Table{Cols: []w.TableCol{{Label: "When"}, {Label: "Job"}, {Label: "Size", Rank: w.Rank3, Secondary: true}, {Label: "Collections", Rank: w.Rank4}, {Label: "Verified"}, {Label: "Trigger", Secondary: true}, {Width: w.ColTime}}, Rows: []w.TableRow{
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.Text("2026-08-16 03:10")}, Mono: true}, {Parts: []w.CellPart{w.Text("Рабочее, ежедневно")}}, {Parts: []w.CellPart{w.Text("18.4 MB")}, Mono: true}, {Parts: []w.CellPart{w.Text("3 / 3")}, Mono: true}, {Parts: []w.CellPart{w.Text("read back")}}, {Parts: []w.CellPart{w.Text("cron")}, Mono: true}, {Parts: []w.CellPart{w.Button{Label: "Restore", Small: true}, w.Button{Icon: "down", Small: true, Variant: w.ButtonPadX3, Title: "Download"}}, Acts: true}}},
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.Text("2026-08-15 03:10")}, Mono: true}, {Parts: []w.CellPart{w.Text("Рабочее, ежедневно")}}, {Parts: []w.CellPart{w.Text("18.2 MB")}, Mono: true}, {Parts: []w.CellPart{w.Text("2 / 3")}, Mono: true, Tone: w.CellDanger}, {Parts: []w.CellPart{w.Text("baikal · Общий: 401")}, Tone: w.CellDanger}, {Parts: []w.CellPart{w.Text("cron")}, Mono: true}, {Parts: []w.CellPart{w.Button{Label: "Restore", Small: true}, w.Button{Icon: "down", Small: true, Variant: w.ButtonPadX3, Title: "Download"}}, Acts: true}}},
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.Text("2026-08-14 09:12")}, Mono: true}, {Parts: []w.CellPart{w.Text("Личное на домашний диск")}}, {Parts: []w.CellPart{w.Text("2.1 GB")}, Mono: true}, {Parts: []w.CellPart{w.Text("2 / 2 + 1 240 files")}, Mono: true}, {Parts: []w.CellPart{w.Text("read back")}}, {Parts: []w.CellPart{w.Text("on sign-in")}, Mono: true}, {Parts: []w.CellPart{w.Button{Label: "Restore", Small: true}, w.Button{Icon: "down", Small: true, Variant: w.ButtonPadX3, Title: "Download"}}, Acts: true}}}}}}}, "s-backup-1", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Banner{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}, Badge: &w.Badge{Label: "read-only"}, End: w.Line{Gap: w.LineGap6, Place: w.Pushed, Items: []w.LinePart{w.Seg{Options: []w.SegOption{{Label: "Full width", On: true}, {Label: "Reading width"}}}, w.Button{Label: "Download .md", Icon: "down"}}}},
		w.DocLayout{Items: []w.DocLayoutPart{w.Doc{Items: []w.DocPart{
			w.DocTitle{Text: "Заметка со встречи 10.08"},
			w.DocSub{Parts: []w.DocSubPart{w.Mono{Text: "10 August 2026, 11:05"}}},
			w.Prose{Parts: []w.Inline{w.Text("Договорились заморозить ветку в среду. Письмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.")}, Variant: w.ProseLead},
			w.DocTasks{Title: w.Rubric{Text: "Что делаем", Variant: w.RubricBottom8}, Items: []w.DocTask{{Tick: w.Tick{Variant: w.TickNudge, Label: "проверить changelog"}, Text: "проверить changelog"}, {Tick: w.Tick{Done: true, Variant: w.TickNudge, Label: "попросить у Артёма скриншоты"}, Text: "попросить у Артёма скриншоты"}}},
			w.DocFiles{Title: w.Rubric{Text: "Attachments", Variant: w.RubricBottom7}, Items: []w.DocFile{{Name: "2026-08-10-sync-board.png", Size: w.Dim{Parts: []w.Inline{w.Text("· 412 KB")}}}, {Name: "release-checklist.md", Size: w.Dim{Parts: []w.Inline{w.Text("· 3 KB")}}}}}}}}, View: w.DocViewFocus},
		w.MainFoot{Text: "This link works until 19 August 2026."}}}, "s-publish-8", "tc-main", 1),
	shell("plain", w.Main{Items: []w.MainPart{
		w.Banner{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}, Badge: &w.Badge{Label: "read-only"}, End: w.Push{Items: []w.PushPart{w.Button{Label: "Download", Icon: "down", Kind: w.ButtonPrimary}}}},
		w.Head{Title: &w.HeadTitle{Text: "2026-08-10-sync-board.png", Size: w.HeadTitle19}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mono{Text: "image/png · 412 KB · changed 2026-08-10"}}}}, Pad: w.HeadPadBottom6},
		w.MainThumbs{Thumb: w.Thumb{Icon: "img", Form: w.ThumbWide}},
		w.MainFoot{Text: "This link works until 19 August 2026."}}}, "s-publish-9", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Banner{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}, Badge: &w.Badge{Label: "read-only"}, End: w.Line{Gap: w.LineGap6, Place: w.Pushed, Items: []w.LinePart{w.Seg{Options: []w.SegOption{{Label: "Full width", On: true}, {Label: "Reading width"}}}, w.Button{Label: "Download", Icon: "down"}}}},
		w.Head{Title: &w.HeadTitle{Text: "release-checklist.md", Size: w.HeadTitle19}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mono{Text: "text/markdown · 3.1 KB · changed 2026-08-09"}}}}, Pad: w.HeadPadBottom6},
		w.Doc{Items: []w.DocPart{
			w.DocHeading{Text: "Перед заморозкой"},
			w.DocBullets{Items: [][]w.Inline{[]w.Inline{w.Text("прогнать "), w.Mono{Text: "go test -race ./...", Variant: w.MonoChip}}, []w.Inline{w.Text("проверить changelog")}, []w.Inline{w.Text("письмо пользователям — черновик у Анны")}}},
			w.DocHeading{Text: "После"},
			w.DocBody{Text: "Тег ставит Виктор. Образ собирается вручную до появления CI."}}, Text: true},
		w.MainFoot{Text: "This link works until 19 August 2026 · the file is read from the server on every visit"}}}, "s-publish-10", "tc-main", 1),
	shell("plain", w.Main{Items: []w.MainPart{
		w.Banner{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}, Badge: &w.Badge{Label: "read-only"}},
		w.Stage{Thumb: w.Thumb{Icon: "doc", Form: w.ThumbSquare}, Title: "макет-обложки.psd", Meta: w.Mono{Text: "image/vnd.adobe.photoshop · 84.2 MB · changed 2026-08-04"}, Download: w.Button{Label: "Download 84.2 MB", Icon: "down", Kind: w.ButtonPrimary}, Note: "This file cannot be shown in a browser, and Carrel does not try. It is streamed with its own type, transcoded nowhere and stored nowhere."},
		w.MainFoot{Text: "This link works until 19 August 2026."}}}, "s-publish-11", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Banner{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}, Badge: &w.Badge{Label: "read-only"}, End: w.Line{Gap: w.LineGap6, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Download folder", Icon: "down"}}}},
		w.Head{Title: &w.HeadTitle{Text: "Материалы к релизу", Size: w.HeadTitle19}, Crumbs: &w.Crumbs{Current: "Материалы к релизу", Aside: &w.Dim{Parts: []w.Inline{w.Text("· 4 files · 6.6 MB")}}}, Pad: w.HeadPadBottom8},
		w.Strip{Label: "Mount as a drive", Items: []w.StripPart{w.Input{Value: "https://carrel.example/p/5c40e2/dav/", Mono: true, Variant: w.InputLineFill, Label: "Address"}, w.Button{Label: "Copy", Small: true}, w.Button{Label: "How to mount", Small: true}, w.Dim{Parts: []w.Inline{w.Text("read-only WebDAV, no password")}, Size: w.Dim11}}},
		w.Table{Cols: []w.TableCol{{Label: "Name"}, {Label: "Size", Width: w.ColSize}, {Label: "Type", Width: w.ColChanged, Rank: w.Rank3}, {Label: "Changed", Width: w.ColTime, Rank: w.Rank4}, {Width: w.Col90}}, Rows: []w.TableRow{
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.TableFile{Icon: "file", Name: "скриншоты"}}}, {Parts: []w.CellPart{w.Text("—")}, Mono: true}, {Parts: []w.CellPart{w.Text("folder · 6 files")}, Mono: true}, {Parts: []w.CellPart{w.Text("2026-08-10 11:07")}, Mono: true}, {Parts: []w.CellPart{w.Button{Label: "Open", Small: true}}, Acts: true}}, State: w.RowFolder},
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "release-checklist.md"}}}, {Parts: []w.CellPart{w.Text("3 KB")}, Mono: true}, {Parts: []w.CellPart{w.Text("text/markdown")}, Mono: true}, {Parts: []w.CellPart{w.Text("2026-08-09 18:44")}, Mono: true}, {Parts: []w.CellPart{w.Button{Label: "Download", Small: true}}, Acts: true}}},
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.TableFile{Icon: "img", Name: "2026-08-04-whiteboard.jpg"}}}, {Parts: []w.CellPart{w.Text("2.8 MB")}, Mono: true}, {Parts: []w.CellPart{w.Text("image/jpeg")}, Mono: true}, {Parts: []w.CellPart{w.Text("2026-08-04 17:03")}, Mono: true}, {Parts: []w.CellPart{w.Button{Label: "Download", Small: true}}, Acts: true}}},
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "смета-подрядчик.pdf"}}}, {Parts: []w.CellPart{w.Text("1.2 MB")}, Mono: true}, {Parts: []w.CellPart{w.Text("application/pdf")}, Mono: true}, {Parts: []w.CellPart{w.Text("2026-08-05 14:20")}, Mono: true}, {Parts: []w.CellPart{w.Button{Label: "Download", Small: true}}, Acts: true}}}}},
		w.MainFoot{Text: "This link works until 1 September 2026 · the contents are read from the server on every visit"}}}, "s-publish-12", "tc-main", 1),
	shell("groups", w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Text: "Контакты"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("312 записей · 2 отмечено")}}}},
		w.ListGroup{Title: w.Rubric{Text: "А"}, Count: &w.Num{Value: 4}},
		w.List{Rows: []w.ListRow{
			w.FaceRow{Mark: w.Mark{Colour: "#4A6B52"}, Avatar: w.Avatar{Initials: "АК"}, Title: []w.Inline{w.Text("Анна Ковалёва")}, Strong: true, Subtitle: "+7 912 445-10-88", Detail: "a.kovaleva@example.org", Caption: "Издательство «Слово»", Source: "nextcloud · Личный", On: true},
			w.FaceRow{Mark: w.Mark{Colour: "#6B4E7A"}, Avatar: w.Avatar{Photo: true}, Title: []w.Inline{w.Text("Артём Белов")}, Strong: true, Subtitle: "+7 903 118-22-04", Detail: "a.belov@example.org", Caption: "—", Source: "radicale · vps · Команда", On: true},
			w.FaceRow{Mark: w.Mark{Colour: "#4A6B52"}, Avatar: w.Avatar{Initials: "АГ"}, Title: []w.Inline{w.Text("Алексей Гринёв")}, Strong: true, Subtitle: "+7 921 330-07-19", Detail: "a.grinev@example.org", Caption: "—", Source: "nextcloud · Личный"},
			w.FaceRow{Mark: w.Mark{Colour: "#3F6E8C"}, Avatar: w.Avatar{Initials: "АЛ"}, Title: []w.Inline{w.Text("Антонина Лебедева")}, Strong: true, Subtitle: "+7 916 204-55-31", Detail: "a.lebedeva@example.org", Caption: "Проектное бюро", Source: "nextcloud · Работа"}}},
		w.ListGroup{Title: w.Rubric{Text: "Б"}, Count: &w.Num{Value: 3}},
		w.List{Rows: []w.ListRow{
			w.FaceRow{Mark: w.Mark{Colour: "#4A6B52"}, Avatar: w.Avatar{Initials: "БК"}, Title: []w.Inline{w.Text("Борис Кузьмин")}, Strong: true, Subtitle: "+7 905 771-42-08", Detail: "b.kuzmin@example.org", Caption: "—", Source: "nextcloud · Личный"},
			w.FaceRow{Mark: w.Mark{Colour: "#3F6E8C"}, Avatar: w.Avatar{Photo: true}, Title: []w.Inline{w.Text("Богдан Ремизов")}, Strong: true, Subtitle: "+7 999 018-63-77", Detail: "b.remizov@example.org", Caption: "СМУ-14", Source: "nextcloud · Работа"},
			w.FaceRow{Mark: w.Mark{Colour: "#6B4E7A"}, Avatar: w.Avatar{Initials: "БЯ"}, Title: []w.Inline{w.Text("Белла Янковская")}, Strong: true, Subtitle: "+7 926 887-11-40", Detail: "b.yankovskaya@example.org", Caption: "—", Source: "radicale · vps · Команда"}}},
		w.ListGroup{Title: w.Rubric{Text: "В"}, Count: &w.Num{Value: 2}},
		w.List{Rows: []w.ListRow{
			w.FaceRow{Mark: w.Mark{Colour: "#4A6B52"}, Avatar: w.Avatar{Initials: "ВД"}, Title: []w.Inline{w.Text("Вера Дорошенко")}, Strong: true, Subtitle: "+7 903 552-19-06", Detail: "v.doroshenko@example.org", Caption: "—", Source: "nextcloud · Личный"},
			w.FaceRow{Mark: w.Mark{Colour: "#8C6239"}, Avatar: w.Avatar{Initials: "ВП"}, Title: []w.Inline{w.Text("Виктор Панов")}, Strong: true, Subtitle: "+7 910 443-28-15", Detail: "v.panov@example.org", Caption: "Смета-Сервис", Source: "baikal · Офис"}}},
		w.ListGroup{Title: w.Rubric{Text: "Г"}, Count: &w.Num{Value: 4}},
		w.List{Rows: []w.ListRow{
			w.FaceRow{Mark: w.Mark{Colour: "#4A6B52"}, Avatar: w.Avatar{Initials: "ГМ"}, Title: []w.Inline{w.Text("Галина Мещерякова")}, Strong: true, Subtitle: "+7 962 114-90-23", Detail: "g.mescher@example.org", Caption: "—", Source: "nextcloud · Личный"},
			w.FaceRow{Mark: w.Mark{Colour: "#3F6E8C"}, Avatar: w.Avatar{Initials: "ГС"}, Title: []w.Inline{w.Text("Геннадий Соловьёв")}, Strong: true, Subtitle: "+7 911 605-37-84", Detail: "g.solovyov@example.org", Caption: "Проектное бюро", Source: "nextcloud · Работа"},
			w.FaceRow{Mark: w.Mark{Colour: "#6B4E7A"}, Avatar: w.Avatar{Photo: true}, Title: []w.Inline{w.Text("Глеб Тарасов")}, Strong: true, Subtitle: "+7 985 229-46-12", Detail: "g.tarasov@example.org", Caption: "—", Source: "radicale · vps · Команда"},
			w.FaceRow{Mark: w.Mark{Colour: "#8C6239"}, Avatar: w.Avatar{Initials: "ГФ"}, Title: []w.Inline{w.Text("Григорий Фомин")}, Strong: true, Subtitle: "+7 917 783-51-60", Detail: "g.fomin@example.org", Caption: "СМУ-14", Source: "baikal · Офис"}}}}}, "s-mail-1", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Text: "Add users"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Three ways, for three situations. There is no global mode: any of them works.")}}}},
		w.MainCards{Cards: []w.Card{w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Invitation"},
			w.CardFields{Items: []w.CardFieldsPart{w.Field{Label: "Role", Width: w.FieldSM, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Button{Label: "Create link", Kind: w.ButtonPrimary}}, Wide: true},
			w.CardRow{Kind: w.CardTop10, Gap: 6, Items: []w.CardRowPart{w.Input{Value: "https://carrel.example/invite/7f3a…", Mono: true, Variant: w.InputKeyFill, Label: "Key"}, w.Button{Label: "Copy", Small: true}}},
			w.CardFields{Items: []w.CardFieldsPart{w.Field{Label: "Or send to", Width: w.FieldMD, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Button{Label: "Send invitation"}}, Wide: true, Rule: true}}}, w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Temporary password"},
			w.Msg{Parts: []w.MsgPart{w.Text("Until the user changes this password, you can sign in as them.")}, Tone: w.MsgAlert, Space: w.MsgFlushBottom12},
			w.CardFields{Items: []w.CardFieldsPart{w.Field{Label: "Role", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Login", Width: w.FieldSM, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Email (optional)", Width: w.FieldMD, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Temporary password", Width: w.FieldMD, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Button{Label: "Create user", Kind: w.ButtonPrimary}}, Wide: true}}}, w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Self-registration"},
			w.CardText{Kind: w.CardLede, Parts: []w.CardTextPart{w.Text("A visitor creates an account and confirms the address by email. Cannot be turned on until mail works.")}},
			w.Check{Parts: []w.CheckPart{w.Text("Publish the registration form")}},
			w.CardRow{Kind: w.CardTop10, Gap: 6, Items: []w.CardRowPart{w.Input{Value: "https://carrel.example/register", Mono: true, Variant: w.InputKeyFill, Label: "Key"}, w.Button{Label: "Copy", Small: true}}},
			w.CardRow{Kind: w.CardTop10, Items: []w.CardRowPart{w.Button{Label: "Save"}}}}}}},
		w.ListGroup{Title: w.Rubric{Text: "Pending invitations"}, Count: &w.Num{Value: 2}},
		w.Table{Cols: []w.TableCol{{Label: "Login"}, {Label: "Delivery"}, {Label: "Status", Rank: w.Rank3}, {Label: "Mail", Rank: w.Rank4}, {Label: "Expires"}, {}}, Rows: []w.TableRow{
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.Text("—")}}, {Parts: []w.CellPart{w.Text("link")}, Mono: true}, {Parts: []w.CellPart{w.Text("pending")}}, {Parts: []w.CellPart{w.Text("—")}}, {Parts: []w.CellPart{w.Text("2026-08-18 12:00")}, Mono: true}, {Parts: []w.CellPart{w.Button{Label: "Copy link", Small: true}, w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}}, Acts: true}}},
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.Text("anna")}}, {Parts: []w.CellPart{w.Text("email")}, Mono: true}, {Parts: []w.CellPart{w.Text("pending")}}, {Parts: []w.CellPart{w.Text("failed — 550 mailbox unavailable")}, Mono: true, Tone: w.CellDanger}, {Parts: []w.CellPart{w.Text("2026-08-17 09:30")}, Mono: true}, {Parts: []w.CellPart{w.Button{Label: "Resend", Small: true}, w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}}, Acts: true}}}}}}}, "s-admin-1", "tc-main", 1),
	shell("plain", w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Text: "DAV validator"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Discovery without saving credentials")}}}},
		w.Form{},
		w.MainDiag{Diag: w.Diag{Text: "[well-known] GET /.well-known/caldav → 301 → /dav.php/ [principal] PROPFIND /dav.php/ → 207, principal /dav.php/principals/m.yenuchenko/ [home-set] calendar-home-set → /dav.php/calendars/m.yenuchenko/ [home-set] addressbook-home-set → /dav.php/addressbooks/m.yenuchenko/ [collections] 4 calendars, 1 address book, 0 plain folders"}},
		w.Msg{Parts: []w.MsgPart{w.Text("The second, mutating stage — create test objects, exercise the operations, read them back, delete them — never runs silently: it is agreed to explicitly and written to the audit log.")}}}}, "s-admin-5", "tc-main", 1),
	shell("phone", w.Main{Items: []w.MainPart{
		w.Runbar{Meter: w.Meter{Percent: 83, Width: w.Meter60}, Parts: []w.RunbarPart{w.Text("5 of 6 · 14 events · "), w.Span{Face: w.FaceDanger, Text: "1 unavailable"}}},
		w.ListGroup{Title: w.Rubric{Text: "Mon · 10 Aug"}, Count: &w.Num{Value: 3}},
		w.List{Rows: []w.ListRow{
			w.PhoneRow{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH32}, Title: "Синк-ап по релизу", Sub: []w.Inline{w.Mono{Text: "10:00 – 11:00"}, w.Text(" · Переговорная 2")}, Aside: "Личный"},
			w.PhoneRow{Mark: &w.Mark{Colour: "#8C6239", Size: w.MarkH32}, Title: "Дежурство: первая линия", Sub: []w.Inline{w.Mono{Text: "all day"}, w.Text(" · repeats")}, Aside: "Дежурства"},
			w.PhoneRow{Mark: &w.Mark{Colour: "#3F6E8C", Size: w.MarkH32}, Title: "Созвон с подрядчиком", Sub: []w.Inline{w.Mono{Text: "15:00"}, w.Text(" · Jitsi")}, Aside: "Работа"}}},
		w.ListGroup{Title: w.Rubric{Text: "Tue · 11 Aug"}, Count: &w.Num{Value: 2}},
		w.List{Rows: []w.ListRow{
			w.PhoneRow{Mark: &w.Mark{Colour: "#6B4E7A", Size: w.MarkH32}, Title: "Ревью архитектуры", Sub: []w.Inline{w.Mono{Text: "09:30"}, w.Text(" · Переговорная 1")}, Aside: "Проекты"}}},
		w.BottomNav{Items: []w.BottomItem{{Icon: "con", Label: "Contacts", Href: "#"}, {Icon: "cal", Label: "Calendar", Href: "#"}, {Icon: "task", Label: "Tasks", Href: "#"}, {Icon: "note", Label: "Notes", Href: "#"}, {Icon: "file", Label: "Files", Href: "#"}}, On: 2}}}, "phone-1", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Text: "Синк-ап по релизу", Size: w.HeadTitle19}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, w.Text("nextcloud · Личный")}}}, Pad: w.HeadPadBottom8},
		w.Fields{Items: []w.FieldsItem{
			{Term: "When", Value: []w.ValuePart{w.Mono{Text: "10:00 – 11:00"}, w.Break{}, w.Text("Monday, 10 August")}},
			{Term: "Repeats", Value: []w.ValuePart{w.Text("Weekly on Mon")}},
			{Term: "Where", Value: []w.ValuePart{w.Text("Переговорная 2")}},
			{Term: "Status", Value: []w.ValuePart{w.Text("Confirmed")}}}, Variant: w.Fields88},
		w.Sec{Items: []w.SecPart{
			w.Rubric{Text: "Description"},
			w.Prose{Parts: []w.Inline{w.Text("Обсудить оставшиеся блокеры, дату заморозки и кто пишет письмо пользователям.")}, Variant: w.Prose13}}},
		w.Sec{Items: []w.SecPart{
			w.Rubric{Text: "Attendees 4"},
			w.SecList{Items: [][]w.SecItemPart{[]w.SecItemPart{w.Text("Анна Ковалёва "), w.Dim{Parts: []w.Inline{w.Text("· accepted")}}}, []w.SecItemPart{w.Text("Артём Белов "), w.Dim{Parts: []w.Inline{w.Text("· tentative")}}}}, Gap: 5}}},
		w.FootActs{Items: []w.Button{w.Button{Label: "Write a note"}, w.Button{Label: "Delete", Kind: w.ButtonDangerSolid}}}}}, "phone-3", "tc-main", 1),
	shell("phone", w.Main{Items: []w.MainPart{
		w.ListGroup{Title: w.Rubric{Text: "А"}, Count: &w.Num{Value: 3}},
		w.List{Rows: []w.ListRow{
			w.PhoneRow{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH34}, Avatar: &w.Avatar{Initials: "АК", Size: w.Avatar34}, Title: "Анна Ковалёва", Sub: []w.Inline{w.Text("+7 912 445-10-88")}, SubMono: true, Aside: "Контакты"},
			w.PhoneRow{Mark: &w.Mark{Colour: "#6B4E7A", Size: w.MarkH34}, Avatar: &w.Avatar{Photo: true, Size: w.Avatar34}, Title: "Артём Белов", TitleBadge: &w.Badge{Label: "×2", Tone: w.BadgeLinked}, Sub: []w.Inline{w.Text("+7 903 118-22-04")}, SubMono: true, Aside: "Команда"},
			w.PhoneRow{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH34}, Avatar: &w.Avatar{Initials: "АС", Size: w.Avatar34}, Title: "Алексей Смирнов", Sub: []w.Inline{w.Text("+7 916 200-71-15")}, SubMono: true, Aside: "Контакты"}}},
		w.ListGroup{Title: w.Rubric{Text: "В"}, Count: &w.Num{Value: 2}},
		w.List{Rows: []w.ListRow{
			w.PhoneRow{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH34}, Avatar: &w.Avatar{Initials: "ВГ", Size: w.Avatar34}, Title: "Виктор Гончаров", Sub: []w.Inline{w.Text("v.goncharov@example.org")}, Aside: "Контакты"}}},
		w.BottomNav{Items: []w.BottomItem{{Icon: "con", Label: "Contacts", Href: "#"}, {Icon: "cal", Label: "Calendar", Href: "#"}, {Icon: "task", Label: "Tasks", Href: "#"}, {Icon: "note", Label: "Notes", Href: "#"}, {Icon: "file", Label: "Files", Href: "#"}}, On: 1}}}, "phone-4", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Person{Initials: "АК", Name: "Анна Ковалёва", Sub: "Издательство «Слово»", Badge: &w.Badge{Label: "in 2 books", Tone: w.BadgeLinked}, Acts: []w.Button{w.Button{Label: "Call"}, w.Button{Label: "Email"}, w.Button{Icon: "quill", Title: "More"}}},
		w.Bar{Items: []w.BarPart{w.Seg{Options: []w.SegOption{{Label: "Card", On: true}, {Label: "Connections 27"}}}}},
		w.Fields{Items: []w.FieldsItem{
			{Term: "Phone", Value: []w.ValuePart{w.Text("+7 912 445-10-88")}, Face: w.ValueMonoAccent},
			{Term: "Email", Value: []w.ValuePart{w.Text("a.kovaleva@example.org")}, Face: w.ValueAccent},
			{Term: "Birthday", Value: []w.ValuePart{w.Text("1987-03-14")}, Face: w.ValueMono}}, Variant: w.Fields88Pad},
		w.BottomNav{Items: []w.BottomItem{{Icon: "con", Label: "Contacts", Href: "#"}, {Icon: "cal", Label: "Calendar", Href: "#"}, {Icon: "task", Label: "Tasks", Href: "#"}, {Icon: "note", Label: "Notes", Href: "#"}, {Icon: "file", Label: "Files", Href: "#"}}, On: 1}}}, "phone-5", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Bar{Items: []w.BarPart{w.Button{Icon: "down", Title: "More"}, w.Button{Icon: "move", Title: "More"}, w.Button{Icon: "trash", Kind: w.ButtonDangerSolid, Title: "Delete"}, w.Button{Icon: "close", Title: "More", Place: w.Pushed}}, Count: &w.BarCount{Text: "2 selected"}, Variant: w.BarSelection},
		w.List{Rows: []w.ListRow{
			w.PhoneRow{Pick: &w.TickBox{}, Icon: "img", Title: "2026-08-10-sync-board.png", Sub: []w.Inline{w.Text("412 KB · image/png")}, SubMono: true, More: &w.Button{Icon: "dots", IconBox: w.IconBox16, Variant: w.ButtonPadX4, Title: "More"}, On: true},
			w.PhoneRow{Pick: &w.TickBox{}, Icon: "doc", Title: "release-checklist.md", Sub: []w.Inline{w.Text("3 KB · text/markdown")}, SubMono: true, More: &w.Button{Icon: "dots", IconBox: w.IconBox16, Variant: w.ButtonPadX4, Title: "More"}, On: true},
			w.PhoneRow{Pick: &w.TickBox{Ticked: w.TickedOff}, Icon: "file", Title: "2026-07", Sub: []w.Inline{w.Text("34 items")}, SubMono: true, More: &w.Button{Icon: "dots", IconBox: w.IconBox16, Variant: w.ButtonPadX4, Title: "More"}}}},
		w.BottomNav{Items: []w.BottomItem{{Icon: "con", Label: "Contacts", Href: "#"}, {Icon: "cal", Label: "Calendar", Href: "#"}, {Icon: "task", Label: "Tasks", Href: "#"}, {Icon: "note", Label: "Notes", Href: "#"}, {Icon: "file", Label: "Files", Href: "#"}}, On: 5}}}, "phone-6", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Doc{Items: []w.DocPart{
			w.DocTitle{Text: "Заметка со встречи 10.08"},
			w.DocSub{Parts: []w.DocSubPart{w.Mono{Text: "10 August, 11:05"}, w.Text(" · "), w.Tag{Label: "встречи"}}},
			w.Prose{Parts: []w.Inline{w.Text("Договорились заморозить ветку в среду. Письмо пользователям пишет Анна, черновик к вечеру вторника.")}, Variant: w.ProseReadTop},
			w.DocTasks{Title: w.Rubric{Text: "Что делаем", Variant: w.RubricBottom7}, Items: []w.DocTask{{Tick: w.Tick{Variant: w.TickNudge, Label: "проверить changelog"}, Text: "проверить changelog"}, {Tick: w.Tick{Done: true, Variant: w.TickNudge, Label: "попросить скриншоты"}, Text: "попросить скриншоты"}}, Phone: true},
			w.DocRelated{Title: w.Rubric{Text: "Related to", Variant: w.RubricBottom7}, Name: "Синк-ап по релизу", What: w.Dim{Parts: []w.Inline{w.Text("· event")}}}}},
		w.BottomNav{Items: []w.BottomItem{{Icon: "con", Label: "Contacts", Href: "#"}, {Icon: "cal", Label: "Calendar", Href: "#"}, {Icon: "task", Label: "Tasks", Href: "#"}, {Icon: "note", Label: "Notes", Href: "#"}, {Icon: "file", Label: "Files", Href: "#"}}, On: 4}}}, "phone-7", "tc-main", 1),

	// tc-side
	shell("plain", w.Side{Items: []w.SidePart{w.Head{Photo: &w.HeadPhoto{Initials: "АК"}, Title: &w.HeadTitle{Text: "Анна Ковалёва", Size: w.HeadTitle19}, Tools: []w.Button{w.Button{Label: "Edit", Small: true}, w.Button{Label: "All records", Small: true}, w.Button{Icon: "dots", Small: true, Title: "More"}}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Icon: "close", Title: "Close the panel"}}}, Middle: true}, w.Fields{Items: []w.FieldsItem{
		{Term: "Organisation", Value: []w.ValuePart{w.Text("Издательство «Слово»")}},
		{Term: "Title", Value: []w.ValuePart{w.Text("Выпускающий редактор")}},
		{Term: "Phone · mobile", Value: []w.ValuePart{w.Text("+7 912 445-10-88")}, Face: w.ValueMono},
		{Term: "Phone · work", Value: []w.ValuePart{w.Text("+7 495 120-45-00")}, Face: w.ValueMono},
		{Term: "Email · work", Value: []w.ValuePart{w.Text("a.kovaleva@example.org")}},
		{Term: "Birthday", Value: []w.ValuePart{w.Text("1987-03-14")}, Face: w.ValueMono},
		{Term: "Categories", Value: []w.ValuePart{w.Tag{Label: "работа"}, w.Tag{Label: "издательство"}}},
		{Term: "Note", Value: []w.ValuePart{w.Text("Пишет по вторникам, отвечает в тот же день.")}}}}, w.Sec{Items: []w.SecPart{
		w.Rubric{Text: "Kept as the server sent them"}}}, w.Sec{Items: []w.SecPart{
		w.Rubric{Text: "Source"},
		w.Diag{Text: "UID 4f2b-9c31-a0e7 Read 2 min ago"}}}, w.Danger{Button: w.Button{Label: "Delete contact", Kind: w.ButtonDangerSolid}}}}, "s-contacts-1", "tc-side", 1),
	shell("zone", w.Side{Items: []w.SidePart{w.Zone{Title: w.Rubric{Text: "Details"}, Note: w.Mono{Text: "400 px · closes by ✕", Variant: w.Mono11Top}}}, Variant: w.SideZone}, "shell-1", "tc-side", 1),
	shell("acts", w.Side{Items: []w.SidePart{w.Head{Title: &w.HeadTitle{Text: "2026-08-10-sync-board.png", Size: w.HeadTitle17}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("sftpgo · vps · Documents")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Icon: "close", Title: "Close the panel"}}}, Pad: w.HeadPadBottom8}, w.Pad{Items: []w.PadPart{w.Thumb{Icon: "img", Form: w.ThumbPreview}}}, w.SideActs{Items: []w.Button{w.Button{Label: "Download", Icon: "down", Kind: w.ButtonPrimary}, w.Button{Label: "Open"}, w.Button{Label: "Rename", Icon: "pencil"}, w.Button{Label: "Move", Icon: "move"}}}, w.Fields{Items: []w.FieldsItem{
		{Term: "Size", Value: []w.ValuePart{w.Text("412 KB")}, Face: w.ValueMono},
		{Term: "Type", Value: []w.ValuePart{w.Text("image/png")}, Face: w.ValueMono},
		{Term: "Changed", Value: []w.ValuePart{w.Text("2026-08-10 11:07")}, Face: w.ValueMono},
		{Term: "Path", Value: []w.ValuePart{w.Text("/dav/Documents/carrel/attachments/2026-08-10-sync-board.png")}, Face: w.ValueMonoSmall}}}, w.Sec{Items: []w.SecPart{
		w.Rubric{Text: "Links"},
		w.SecNote{Text: "Carrel keeps no index of what points at this file. It may be attached to a note or an event — check before deleting."}}}, w.Danger{Button: w.Button{Label: "Delete", Icon: "trash", Kind: w.ButtonDangerSolid}}}}, "s-files-1", "tc-side", 1),
	shell("drop", w.Side{Items: []w.SidePart{w.Head{Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#4A6B52"}, Text: "Синк-ап по релизу", Size: w.HeadTitle19}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Личный")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Edit", Small: true}, w.Button{Icon: "close", Title: "Close the panel"}}}}, w.Fields{Items: []w.FieldsItem{
		{Term: "When", Value: []w.ValuePart{w.Mono{Text: "10:00 – 11:00"}, w.Text(", Monday 10 August")}},
		{Term: "Where", Value: []w.ValuePart{w.Text("Переговорная 2")}}}}, w.Sec{Items: []w.SecPart{
		w.Rubric{Text: "Attachments"},
		w.SecList{Items: [][]w.SecItemPart{[]w.SecItemPart{w.Text("release-checklist.md "), w.Dim{Parts: []w.Inline{w.Text("text/markdown · 3 KB")}}, w.Text(" "), w.Badge{Label: "in the event"}}}, Gap: 6},
		w.Empty{Lead: "Attach to this event", Parts: []w.EmptyPart{w.Dim{Parts: []w.Inline{w.Text("Over 2 MB it goes to"), w.Mono{Text: "Documents / carrel"}, w.Text("and the event keeps a link")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Variant: w.EmptyDropSide}}}, w.Sec{Items: []w.SecPart{
		w.Rubric{Text: "Related"},
		w.SecList{Items: [][]w.SecItemPart{[]w.SecItemPart{w.Span{Face: w.FaceAccent, Text: "Заметка со встречи 10.08"}, w.Text(" "), w.Dim{Parts: []w.Inline{w.Text("· note")}}}}}}}}, Variant: w.SideDrop}, "s-io-2", "tc-side", 1),

	// tc-shell
	shell("plain", w.Shell{Top: &w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}}, Actions: []w.TopAction{w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary}, w.Button{Label: "2 min ago", Icon: "refresh"}, w.UserMenu{Name: "mix"}}}, Body: w.ShellBody{Rail: []w.RailPart{w.Nav{Items: []w.NavItem{
		{Icon: "back", Label: "Back to the card", Href: "#"}}}, w.RailSec{Items: []w.RailSecPart{
		w.Rubric{Text: "nextcloud"},
		w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "248", Current: true},
		w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "51"}}}}, Main: w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Text: "Анна Ковалёва"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Save", Kind: w.ButtonPrimary}}}},
		w.Form{},
		w.FormFoot{}}}}}, "s-contacts-3", "tc-shell", 1),
	shell("overlay", w.Shell{Top: &w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}}, Actions: []w.TopAction{w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary}, w.Button{Label: "2 min ago", Icon: "refresh"}, w.UserMenu{Name: "mix"}}}, Body: w.Overlay{Body: w.OverlayBody{Rail: []w.OverlayRailPart{w.Nav{Items: []w.NavItem{
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
		w.Msg{Parts: []w.MsgPart{w.Text("Inside one server this is "), w.Mono{Text: "MOVE", Form: w.MonoCode}, w.Text(": instant, no traffic. To another address book the card is copied and then deleted — Carrel says so before it starts, not after. Copy is this same sheet with "), w.Strong("Copy here"), w.Text(".")}, Space: w.MsgFlushTop2}}, Foot: []w.SheetFootPart{w.Button{Label: "Move here", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Enter"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth520}, Pad: w.ScrimPad26}}}, "s-contacts-2", "tc-shell", 1),
	shell("no-top", w.Shell{Body: w.ShellBody{Main: w.Main{Items: []w.MainPart{
		w.Form{}}}}}, "account-8", "tc-shell", 1),
	shell("side", w.Shell{Top: &w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}}, Actions: []w.TopAction{w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary}, w.Button{Label: "2 min ago", Icon: "refresh"}, w.UserMenu{Name: "mix"}}}, Body: w.ShellBody{Rail: []w.RailPart{w.Nav{Items: []w.NavItem{
		{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
		{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
		{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
		{Icon: "file", Label: "Files", Href: "#"},
		{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"}}, On: 1}, w.RailSec{Items: []w.RailSecPart{
		w.TickRow{Ticked: w.TickedPart, Kind: w.TickRowAll, Mark: w.Mark{Kind: w.MarkMulti}, Label: "Combined address book", Meta: "312", Current: true}}}, w.RailSec{Items: []w.RailSecPart{
		w.Rubric{Text: "nextcloud"},
		w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "248"},
		w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "51"}}}, w.RailSec{Items: []w.RailSecPart{
		w.Rubric{Text: "radicale · vps"},
		w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "13", Trail: w.Badge{Label: "ro"}},
		w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#8C6239"}, Label: "Рассылки", Meta: "—", Hidden: true}}}, w.RailSec{Items: []w.RailSecPart{
		w.Rubric{Text: "Tags"}}, Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "работа"}, w.Tag{Label: "издательство"}}}, w.RailFoot{Items: []w.RailFootPart{w.Button{Icon: "plus", Title: "New address book"}, w.Button{Icon: "server", Title: "Connect account"}, w.Button{Icon: "link", Title: "Subscribe to a feed"}, w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}}, About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"}}}, Main: w.Main{Items: []w.MainPart{
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
		w.Empty{Parts: []w.EmptyPart{w.Text("Loading more…")}, Variant: w.EmptyDusty}}}, Side: &w.Side{Items: []w.SidePart{w.Head{Photo: &w.HeadPhoto{Initials: "АК"}, Title: &w.HeadTitle{Text: "Анна Ковалёва", Size: w.HeadTitle19}, Tools: []w.Button{w.Button{Label: "Edit", Small: true}, w.Button{Label: "All records", Small: true}, w.Button{Icon: "dots", Small: true, Title: "More"}}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Icon: "close", Title: "Close the panel"}}}, Middle: true}, w.Fields{Items: []w.FieldsItem{
		{Term: "Organisation", Value: []w.ValuePart{w.Text("Издательство «Слово»")}},
		{Term: "Title", Value: []w.ValuePart{w.Text("Выпускающий редактор")}},
		{Term: "Phone · mobile", Value: []w.ValuePart{w.Text("+7 912 445-10-88")}, Face: w.ValueMono},
		{Term: "Phone · work", Value: []w.ValuePart{w.Text("+7 495 120-45-00")}, Face: w.ValueMono},
		{Term: "Email · work", Value: []w.ValuePart{w.Text("a.kovaleva@example.org")}},
		{Term: "Birthday", Value: []w.ValuePart{w.Text("1987-03-14")}, Face: w.ValueMono},
		{Term: "Categories", Value: []w.ValuePart{w.Tag{Label: "работа"}, w.Tag{Label: "издательство"}}},
		{Term: "Note", Value: []w.ValuePart{w.Text("Пишет по вторникам, отвечает в тот же день.")}}}}, w.Sec{Items: []w.SecPart{
		w.Rubric{Text: "Kept as the server sent them"}}}, w.Sec{Items: []w.SecPart{
		w.Rubric{Text: "Source"},
		w.Diag{Text: "UID 4f2b-9c31-a0e7 Read 2 min ago"}}}, w.Danger{Button: w.Button{Label: "Delete contact", Kind: w.ButtonDangerSolid}}}}}}, "s-contacts-1", "tc-shell", 1),
	shell("no-rail", w.Shell{Top: &w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}}, Actions: []w.TopAction{w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary}, w.Button{Label: "2 min ago", Icon: "refresh"}, w.UserMenu{Name: "mix"}}}, Body: w.ShellBody{Main: w.Main{Items: []w.MainPart{
		w.Docbar{Items: []w.DocbarPart{w.Button{Label: "Save", Kind: w.ButtonPrimary}, w.Button{Label: "Cancel"}, w.BarSep{}, w.Button{Title: "More"}, w.Button{Label: "I", Variant: w.ButtonItalic}, w.Button{Label: "H2"}, w.Button{Label: "• list"}, w.Button{Folds: true, Label: "task", Icon: "task", Title: "Task"}, w.Button{Icon: "link", Title: "More"}, w.Button{Folds: true, Label: "Attach", Menu: true, Icon: "up", Title: "Attach"}}, More: &w.Button{Icon: "dots", Title: "More"}, End: []w.DocbarEnd{w.Check{Parts: []w.CheckPart{w.Text("Auto-save")}}, w.DraftNote{Text: "Draft in this window ·", Time: w.Mono{Text: "11:42"}, Secondary: true}, w.Kbd{Keys: "Ctrl+S", Secondary: true}, w.Seg{Options: []w.SegOption{{Label: "Full width", On: true}, {Label: "Reading width"}}, Secondary: true}, w.Button{Folds: true, Label: "Focus", Icon: "focus", Variant: w.ButtonOnCard, Title: "Focus"}}},
		w.DocLayout{Items: []w.DocLayoutPart{w.Doc{Items: []w.DocPart{
			w.Area{Text: "Заметка со встречи 10.08", Variant: w.AreaTitle, Label: "Title"},
			w.DocFields{Items: []w.Field{w.Field{Label: "Notebook", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Date", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Time", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Tags", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Related to", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}},
			w.Area{Text: noteBody, Variant: w.AreaBody, Focus: true, Label: "Text"},
			w.DocNote{Parts: []w.Inline{w.Text("Properties Carrel does not know — "), w.Mono{Text: "X-JTX-COLOR"}, w.Text(" and two more — are kept unchanged and not shown here.")}}}}}, View: w.DocViewFocus}}}}}, "s-notes-4", "tc-shell", 1),
	shell("plain", w.Shell{Top: &w.Top{Brand: w.Brand{Title: "Carrel"}}, Body: w.ShellBody{Rail: []w.RailPart{w.Nav{Items: []w.NavItem{
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
		w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true}}}, w.RailFoot{Items: []w.RailFootPart{w.Button{Icon: "plus", Title: "New calendar"}, w.Button{Icon: "server", Title: "Connect account"}, w.Button{Icon: "link", Title: "Subscribe to a feed"}, w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}}, About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"}}}, Main: w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti}, Text: "Combined calendar"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("5 of 7 calendars · 10 – 16 August 2026")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}}}},
		w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August"}, Count: &w.Num{Value: 3}},
		w.List{Rows: []w.ListRow{
			w.TimeRow{Mark: w.Mark{Colour: "#4A6B52"}, Time: "10:00", Title: []w.Inline{w.Text("Синк-ап по релизу")}, Detail: "Переговорная 2", Source: "nextcloud · Личный"},
			w.TimeRow{Mark: w.Mark{Colour: "#3F6E8C"}, Time: "13:30", Title: []w.Inline{w.Text("Звонок с Ковалёвой")}, Detail: "Zoom", Source: "nextcloud · Работа"},
			w.TimeRow{Mark: w.Mark{Colour: "#6B4E7A"}, Time: "18:00", Title: []w.Inline{w.Text("Забрать документы")}, Detail: "МФЦ на Ленина", Source: "radicale · vps · Проекты"}}}}}, Height: w.ShellHeight340}}, "s-notes-7", "tc-shell", 1),
	shell("zone", w.Shell{Top: &w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}}, Actions: []w.TopAction{w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary}, w.Button{Label: "2 min ago", Icon: "refresh"}, w.UserMenu{Name: "mix"}}}, Body: w.ShellBody{Rail: []w.RailPart{w.Zone{Title: w.Rubric{Text: "Rail"}, Note: w.Mono{Text: "264 px · sections + sources", Variant: w.Mono11Top}}}, RailForm: w.RailZone, Main: w.Main{Items: []w.MainPart{
		w.Zone{Title: w.Rubric{Text: "Content"}, Note: w.Mono{Text: "fluid · min 0", Variant: w.Mono11Top}}}, Variant: w.MainZone}, Side: &w.Side{Items: []w.SidePart{w.Zone{Title: w.Rubric{Text: "Details"}, Note: w.Mono{Text: "400 px · closes by ✕", Variant: w.Mono11Top}}}, Variant: w.SideZone}, Height: w.ShellHeight230}}, "shell-1", "tc-shell", 1),
	shell("icons", w.Shell{Top: &w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}, Variant: w.SearchFill}, Actions: []w.TopAction{w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary}, w.UserMenu{Name: "mix", NoArrow: true}}}, Body: w.ShellBody{Rail: []w.RailPart{w.Nav{Items: []w.NavItem{
		{Icon: "back", Href: "#", Title: "Back"}}}}, RailForm: w.RailIcons, Main: w.Main{}}}, "s-conflict-2", "tc-shell", 1),
	shell("plain", w.Shell{Top: &w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}}, Actions: []w.TopAction{w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary}, w.Button{Label: "2 min ago", Icon: "refresh"}, w.UserMenu{Name: "mix"}}}, Body: w.ShellBody{Rail: []w.RailPart{w.Nav{Items: []w.NavItem{
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
		w.Head{Title: &w.HeadTitle{Text: "Connections", Size: w.HeadTitle19}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 servers · 12 collections · 2 subscriptions")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Badge{Label: "credentials unlocked · 4:12"}, w.Button{Label: "Lock now", Small: true}}}, Pad: w.HeadPadBottom8},
		w.Pad{Items: []w.PadPart{w.Card{Items: []w.CardPart{
			w.Rubric{Text: "nextcloud"},
			w.Mono{Text: "https://cloud.example.org/remote.php/dav/", Form: w.MonoLine, Variant: w.Mono11_5DustyGap},
			w.Line{Block: true, Gap: w.LineGap8, Items: []w.LinePart{w.Field{Label: "User name", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}, w.Field{Label: "Password", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}}}, Variant: w.CardChosen}, w.Card{Items: []w.CardPart{
			w.Rubric{Text: "sftpgo · vps"},
			w.Mono{Text: "https://files.example.org/dav/ · mix · last answer 5 min ago", Form: w.MonoLine, Variant: w.Mono11_5DustyGap}}}}},
		w.Msg{Parts: []w.MsgPart{w.Text("While credentials are unlocked, every account opens without a question — here, on the New device screen and when a profile is built. "), w.Strong("Lock now"), w.Text(" ends it at once; so does signing out.")}}}}, Height: w.ShellHeight280}}, "s-profile-4", "tc-shell", 1),
	shell("side", w.Shell{Top: &w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}}, Actions: []w.TopAction{w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary}, w.Button{Label: "2 min ago", Icon: "refresh"}, w.UserMenu{Name: "mix"}}}, Body: w.ShellBody{Rail: []w.RailPart{w.Nav{Items: []w.NavItem{
		{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
		{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
		{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
		{Icon: "file", Label: "Files", Href: "#"}}, On: 2}, w.RailSec{Items: []w.RailSecPart{
		w.TickRow{Ticked: w.TickedPart, Kind: w.TickRowAll, Mark: w.Mark{Kind: w.MarkMulti}, Label: "Combined calendar", Meta: "14", Current: true}}}, w.RailSec{Items: []w.RailSecPart{
		w.Rubric{Text: "nextcloud"},
		w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "7"},
		w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "4"}}}, w.RailFoot{Items: []w.RailFootPart{w.Button{Icon: "plus", Title: "New calendar"}, w.Button{Icon: "server", Title: "Connect account"}, w.Button{Icon: "link", Title: "Subscribe to a feed"}, w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}}, About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"}}}, Main: w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti}, Text: "Combined calendar"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("5 of 7 calendars · 10 – 16 August 2026")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}}}},
		w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August"}, Count: &w.Num{Value: 3}},
		w.List{Rows: []w.ListRow{
			w.TimeRow{Mark: w.Mark{Colour: "#4A6B52"}, Time: "10:00", Title: []w.Inline{w.Text("Синк-ап по релизу")}, Strong: true, Detail: "Переговорная 2", Caption: &w.Badge{Label: "4 attendees"}, Source: "nextcloud · Личный", On: true},
			w.TimeRow{Mark: w.Mark{Colour: "#8C6239"}, Time: "all day", Title: []w.Inline{w.Text("Дежурство: первая линия")}, Detail: "—", Source: "2 sources"},
			w.TimeRow{Mark: w.Mark{Colour: "#3F6E8C"}, Time: "15:00", Title: []w.Inline{w.Text("Созвон с подрядчиком")}, Detail: "Jitsi", Source: "nextcloud · Работа"}}}}}, Side: &w.Side{Items: []w.SidePart{w.Head{Title: &w.HeadTitle{Text: "Синк-ап по релизу", Size: w.HeadTitle17}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Личный")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Edit", Small: true}, w.Button{Icon: "close", Title: "Close the panel"}}}, Pad: w.HeadPadBottom8}, w.Fields{Items: []w.FieldsItem{
		{Term: "When", Value: []w.ValuePart{w.Mono{Text: "10:00 – 11:00"}}},
		{Term: "Where", Value: []w.ValuePart{w.Text("Переговорная 2")}}}}, w.Sec{Items: []w.SecPart{
		w.Rubric{Text: "Attachments"},
		w.SecOpen{Items: []w.SecOpenPart{w.SecOpened{Name: "release-checklist.md", Size: w.Dim{Parts: []w.Inline{w.Text("· 3 KB")}}}, w.Line{Gap: w.LineGap6, Items: []w.LinePart{w.Button{Label: "Open in editor", Small: true}, w.Button{Label: "Show in folder", Small: true}}}, w.Dim{Parts: []w.Inline{w.Text("Edits go to the server when the program closes the file. The temporary copy lives until then and is deleted.")}, Size: w.Dim11}}}}}}}, Height: w.ShellHeight390}}, "desktop-1", "tc-shell", 1),
)

func init() {
	Widgets = append(Widgets, w.Head{}, w.Bar{}, w.Main{}, w.Side{}, w.Shell{})
	Parts = append(Parts, w.HeadTitle{}, w.HeadSub{}, w.HeadPhoto{}, w.BarSep{}, w.BarCount{}, w.BarRight{}, w.ListGroup{}, w.MainTree{},
		w.MainHint{}, w.BottomNav{}, w.BottomItem{}, w.FootActs{}, w.SideActs{}, w.ShellBody{}, w.Banner{}, w.Strip{}, w.MainFoot{},
		w.Stage{}, w.Person{})
}
