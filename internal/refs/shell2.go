package refs

import w "carrel.dev/typecase/widget"

var _ = add(
	// tc-head
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
	shell("acts", w.Head{Title: &w.HeadTitle{Text: "Команда", Badge: &w.Badge{Label: "read-only"}}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("radicale · vps · the server granted no write access")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Export", Icon: "down"}, w.Button{Label: "Print", Icon: "print"}}}}, "m-states-4", "tc-head", 1),
	shell("plain", w.Head{Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti}, Text: "Combined calendar"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("shown from this session’s memory")}}}}, "m-states-5", "tc-head", 1),
	shell("padded", w.Head{Title: &w.HeadTitle{Text: "Синк-ап по релизу", Size: w.HeadTitle19}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, w.Text("nextcloud · Личный")}}}, Pad: w.HeadPadBottom8}, "m-phone-3", "tc-head", 1),
	shell("padded", w.Head{Title: &w.HeadTitle{Text: "Application", Size: w.HeadTitle19}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Account and collection settings are inside, in Settings")}}}, Pad: w.HeadPadBottom6}, "m-desktop-13", "tc-head", 1),

	// tc-bar
	shell("selection", w.Bar{Items: []w.BarPart{w.Mono{Text: "412 KB", Variant: w.MonoMuted}, w.BarSep{}, w.Button{Folds: true, Label: "Download", Icon: "down", Title: "Download"}, w.Button{Folds: true, Label: "Rename", Icon: "pencil", Title: "Rename"}, w.Button{Folds: true, Label: "Delete", Icon: "trash", Kind: w.ButtonDangerSolid, Title: "Delete"}}, Count: &w.BarCount{Text: "1 selected"}, Right: &w.BarRight{Items: []w.BarPart{w.Button{Label: "Clear selection", Small: true}}}, Variant: w.BarSelection}, "components-16", "tc-bar", 1),
	shell("more", w.Bar{Items: []w.BarPart{w.Input{Kind: w.InputSearch, Placeholder: "Filter this folder", Variant: w.InputW190, Label: "Filter this folder"}, w.Seg{Options: []w.SegOption{{Label: "Name", On: true}, {Label: "Size"}, {Label: "Changed"}}, Secondary: true}, w.Button{Label: "Columns", Icon: "cols"}}, More: &w.Button{Icon: "dots", Title: "More"}, Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "14 items · 38.2 MB"}}, Secondary: true}}, "components-16", "tc-bar", 2),
	shell("right", w.Bar{Items: []w.BarPart{w.Input{Kind: w.InputSearch, Placeholder: "Filter by word", Variant: w.InputH26W200, Label: "Filter by word"}, w.Seg{Options: []w.SegOption{{Label: "Newest", On: true}, {Label: "Oldest"}, {Label: "Title"}}}, w.Button{Label: "Columns", Icon: "cols"}}, Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "86"}}, Secondary: true}}, "s-notes-1", "tc-bar", 1),
	shell("plain", w.Bar{Items: []w.BarPart{w.Input{Kind: w.InputSearch, Placeholder: "Filter by login", Variant: w.InputH26W200, Label: "Filter by login"}, w.Seg{Options: []w.SegOption{{Label: "All", On: true}, {Label: "Administrators"}, {Label: "Disabled"}}}, w.Button{Label: "Columns", Icon: "cols"}}}, "s-admin-2", "tc-bar", 1),
	shell("selection", w.Bar{Items: []w.BarPart{w.Mono{Text: "2.8 MB", Variant: w.MonoMuted}, w.BarSep{}, w.Button{Folds: true, Label: "Download", Icon: "down", Title: "Download"}, w.Button{Folds: true, Label: "Rename", Icon: "pencil", Title: "Rename"}, w.Button{Folds: true, Label: "Move", Icon: "move", Title: "Move"}, w.Button{Folds: true, Label: "Delete", Icon: "trash", Kind: w.ButtonDangerSolid, Title: "Delete"}}, Count: &w.BarCount{Text: "1 selected"}, Variant: w.BarSelection}, "s-files-5", "tc-bar", 1),
	shell("right", w.Bar{Items: []w.BarPart{w.Filter{Query: w.Query{Value: "Filter the loaded cards", Label: "Filter the loaded cards"}, Waiting: true}, w.Seg{Options: []w.SegOption{{Label: "A–Z", On: true}, {Label: "Recently changed"}}, Secondary: true}, w.Button{Label: "Columns", Icon: "cols"}}, Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "248"}}, Secondary: true}}, "s-io-1", "tc-bar", 1),
	shell("wide", w.Bar{Items: []w.BarPart{w.Range{From: w.Input{Kind: w.InputDate, Value: "2026-01-01", Mono: true, Label: "Date"}, To: w.Input{Value: "today", Mono: true, Label: "Date"}, Show: w.Button{Label: "Show"}}, w.Seg{Options: []w.SegOption{{Label: "All 24", On: true}, {Label: "Events 14"}, {Label: "Tasks 2"}, {Label: "Notes 4"}, {Label: "Files 4"}}, Secondary: true}, w.Button{Label: "Columns", Icon: "cols", Secondary: true}}, More: &w.Button{Icon: "dots", Title: "More"}, Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "newest first"}}, Secondary: true}, Variant: w.BarWide}, "s-contacts-7", "tc-bar", 1),
	shell("selection", w.Bar{Items: []w.BarPart{w.BarSep{}, w.Button{Label: "Restore into", Icon: "back"}, w.Button{Label: "Skip"}}, Count: &w.BarCount{Text: "3 selected"}, Right: &w.BarRight{Items: []w.BarPart{w.Button{Label: "Select all", Small: true}}}, Variant: w.BarSelection}, "s-backup-8", "tc-bar", 1),
	shell("selection", w.Bar{Items: []w.BarPart{w.Button{Icon: "down", Title: "More"}, w.Button{Icon: "move", Title: "More"}, w.Button{Icon: "trash", Kind: w.ButtonDangerSolid, Title: "Delete"}, w.Button{Icon: "close", Title: "More", Place: w.Pushed}}, Count: &w.BarCount{Text: "2 selected", Small: true}, Variant: w.BarSelection}, "m-phone-6", "tc-bar", 1),

	// tc-main
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
		w.MainFoot{Text: "This link works until 9 September 2026. The owner can revoke it at any moment.", Roomy: true}}}, "s-publish-6", "tc-main", 1),
	shell("zone", w.Main{Items: []w.MainPart{
		w.Zone{Title: w.Rubric{Text: "Content"}, Note: w.Mono{Text: "fluid · min 0", Variant: w.Mono11Top}}}, Variant: w.MainZone}, "shell-1", "tc-main", 1),
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
	shell("phone", w.Main{Items: []w.MainPart{
		w.ListGroup{Title: w.Rubric{Text: "А"}, Count: &w.Num{Value: 3}},
		w.List{Rows: []w.ListRow{
			w.PhoneRow{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH34}, Avatar: &w.Avatar{Initials: "АК", Size: w.Avatar34}, Title: "Анна Ковалёва", Sub: []w.Inline{w.Text("+7 912 445-10-88")}, SubMono: true, Aside: "Контакты"},
			w.PhoneRow{Mark: &w.Mark{Colour: "#6B4E7A", Size: w.MarkH34}, Avatar: &w.Avatar{Photo: true, Size: w.Avatar34}, Title: "Артём Белов", TitleBadge: &w.Badge{Label: "×2", Tone: w.BadgeLinked}, Sub: []w.Inline{w.Text("+7 903 118-22-04")}, SubMono: true, Aside: "Команда"},
			w.PhoneRow{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH34}, Avatar: &w.Avatar{Initials: "АС", Size: w.Avatar34}, Title: "Алексей Смирнов", Sub: []w.Inline{w.Text("+7 916 200-71-15")}, SubMono: true, Aside: "Контакты"}}},
		w.ListGroup{Title: w.Rubric{Text: "В"}, Count: &w.Num{Value: 2}},
		w.List{Rows: []w.ListRow{
			w.PhoneRow{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH34}, Avatar: &w.Avatar{Initials: "ВГ", Size: w.Avatar34}, Title: "Виктор Гончаров", Sub: []w.Inline{w.Text("v.goncharov@example.org")}, Aside: "Контакты"}}},
		w.BottomNav{Items: []w.BottomItem{{Icon: "con", Label: "Contacts", Href: "#"}, {Icon: "cal", Label: "Calendar", Href: "#"}, {Icon: "task", Label: "Tasks", Href: "#"}, {Icon: "note", Label: "Notes", Href: "#"}, {Icon: "file", Label: "Files", Href: "#"}}, On: 1}}}, "m-phone-4", "tc-main", 1),
	shell("variant", w.Main{Items: []w.MainPart{
		w.Bar{Items: []w.BarPart{w.Button{Icon: "down", Title: "More"}, w.Button{Icon: "move", Title: "More"}, w.Button{Icon: "trash", Kind: w.ButtonDangerSolid, Title: "Delete"}, w.Button{Icon: "close", Title: "More", Place: w.Pushed}}, Count: &w.BarCount{Text: "2 selected", Small: true}, Variant: w.BarSelection},
		w.List{Rows: []w.ListRow{
			w.PhoneRow{Pick: &w.TickBox{}, Icon: "img", Title: "2026-08-10-sync-board.png", Sub: []w.Inline{w.Text("412 KB · image/png")}, SubMono: true, More: &w.Button{Icon: "dots", IconBox: w.IconBox16, Variant: w.ButtonPadX4, Title: "More"}, On: true},
			w.PhoneRow{Pick: &w.TickBox{}, Icon: "doc", Title: "release-checklist.md", Sub: []w.Inline{w.Text("3 KB · text/markdown")}, SubMono: true, More: &w.Button{Icon: "dots", IconBox: w.IconBox16, Variant: w.ButtonPadX4, Title: "More"}, On: true},
			w.PhoneRow{Pick: &w.TickBox{Ticked: w.TickedOff}, Icon: "file", Title: "2026-07", Sub: []w.Inline{w.Text("34 items")}, SubMono: true, More: &w.Button{Icon: "dots", IconBox: w.IconBox16, Variant: w.ButtonPadX4, Title: "More"}}}},
		w.BottomNav{Items: []w.BottomItem{{Icon: "con", Label: "Contacts", Href: "#"}, {Icon: "cal", Label: "Calendar", Href: "#"}, {Icon: "task", Label: "Tasks", Href: "#"}, {Icon: "note", Label: "Notes", Href: "#"}, {Icon: "file", Label: "Files", Href: "#"}}, On: 5}}}, "m-phone-6", "tc-main", 1),

	// tc-side
	shell("zone", w.Side{Items: []w.SidePart{w.Zone{Title: w.Rubric{Text: "Details"}, Note: w.Mono{Text: "400 px · closes by ✕", Variant: w.Mono11Top}}}, Variant: w.SideZone}, "shell-1", "tc-side", 1),

	// tc-shell
	shell("plain", w.Shell{Top: &w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}}, Actions: []w.TopAction{w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary}, w.Button{Label: "2 min ago", Icon: "refresh"}, w.UserMenu{Name: "mix"}}}, Body: w.ShellBody{Rail: []w.RailPart{w.Nav{Items: []w.NavItem{
		{Icon: "back", Label: "Back to the card", Href: "#"}}}, w.RailSec{Items: []w.RailSecPart{
		w.Rubric{Text: "nextcloud"},
		w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "248", Current: true},
		w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "51"}}}}, Main: w.Main{Items: []w.MainPart{
		w.Head{Title: &w.HeadTitle{Text: "Анна Ковалёва"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Save", Kind: w.ButtonPrimary}}}},
		w.Form{},
		w.FormFoot{}}}}}, "s-contacts-3", "tc-shell", 1),
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
		w.Rubric{Text: "Tags"}}, Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "работа"}, w.Tag{Label: "издательство"}}}, w.RailFoot{Items: []w.RailFootPart{w.Button{Icon: "plus", Title: "New address book"}, w.Button{Icon: "server", Title: "Connect account"}, w.Button{Icon: "link", Title: "Subscribe to a feed"}, w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}}, About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"}}}, Main: w.Main{}, Side: &w.Side{}}}, "s-contacts-1", "tc-shell", 1),
	shell("no-rail", w.Shell{Top: &w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}}, Actions: []w.TopAction{w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary}, w.Button{Label: "2 min ago", Icon: "refresh"}, w.UserMenu{Name: "mix"}}}, Body: w.ShellBody{Main: w.Main{}}}, "s-notes-4", "tc-shell", 1),
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
			{Icon: "theme", Label: "Appearance", Href: "#"}}, On: 2}}}, w.RailFoot{Items: []w.RailFootPart{w.RailStatus{Text: "davloom · running · 3 devices", Detail: w.Dim{Parts: []w.Inline{w.Text("Last response "), w.Mono{Text: "2026-08-16 09:41"}}, Size: w.Dim11}}, w.RailStatus{Text: "Backups are running", Detail: w.Dim{Parts: []w.Inline{w.Text("Last backup "), w.Mono{Text: "2026-08-16 03:10"}, w.Text("· ")}, Size: w.Dim11}}}, About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"}}}, Main: w.Main{}, Height: w.ShellHeight280}}, "s-profile-4", "tc-shell", 1),
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
			w.TimeRow{Mark: w.Mark{Colour: "#3F6E8C"}, Time: "15:00", Title: []w.Inline{w.Text("Созвон с подрядчиком")}, Detail: "Jitsi", Source: "nextcloud · Работа"}}}}}, Side: &w.Side{}, Height: w.ShellHeight390}}, "m-desktop-1", "tc-shell", 1),
)

func init() {
	Widgets = append(Widgets, w.Head{}, w.Bar{}, w.Main{}, w.Side{}, w.Shell{})
	Parts = append(Parts, w.HeadTitle{}, w.HeadSub{}, w.HeadPhoto{}, w.BarSep{}, w.BarCount{}, w.BarRight{}, w.ListGroup{}, w.MainTree{},
		w.MainHint{}, w.BottomNav{}, w.BottomItem{}, w.FootActs{}, w.SideActs{}, w.ShellBody{}, w.Banner{}, w.Strip{}, w.MainFoot{},
		w.Stage{}, w.Person{})
}

var _ = wait("children from the family of overlays (a drop with a menu in a bar)", ref("components-19", "tc-bar", 1), ref("s-agenda-1", "tc-bar", 1))

var _ = wait("children from later families (cards, tables, documents, overlays, states)",
	ref("s-contacts-1", "tc-main", 1), ref("s-agenda-1", "tc-main", 1), ref("s-notes-2", "tc-main", 1), ref("s-dup-1", "tc-main", 1), ref("s-notes-3", "tc-main", 1), ref("s-notes-4", "tc-main", 1), ref("s-files-2", "tc-main", 1), ref("s-files-3", "tc-main", 1), ref("s-find-1", "tc-main", 1), ref("s-profile-10", "tc-main", 1), ref("s-davloom-1", "tc-main", 1), ref("s-davloom-4", "tc-main", 1), ref("s-backup-1", "tc-main", 1), ref("s-publish-8", "tc-main", 1), ref("s-publish-9", "tc-main", 1), ref("s-publish-10", "tc-main", 1), ref("s-publish-11", "tc-main", 1), ref("s-publish-12", "tc-main", 1), ref("s-admin-1", "tc-main", 1), ref("s-admin-5", "tc-main", 1), ref("m-phone-1", "tc-main", 1), ref("m-phone-3", "tc-main", 1), ref("m-phone-5", "tc-main", 1), ref("m-phone-7", "tc-main", 1), ref("s-contacts-1", "tc-side", 1), ref("s-files-1", "tc-side", 1), ref("s-io-2", "tc-side", 1), ref("s-contacts-2", "tc-shell", 1))
