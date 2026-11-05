package refs

import w "carrel.dev/typecase/widget"

func overlay(state string, fixture Renderer, frame, class string, n int) Entry {
	return e("overlays", state, fixture, ref(frame, class, n))
}

var _ = add(
	// tc-menu
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuText{Title: "Windows", Text: "File Explorer › This PC › Map network drive, then paste the address."},
		w.MenuText{Title: "macOS", Text: "Finder › Go › Connect to Server, then paste the address."},
		w.MenuText{Title: "Linux", Text: "Files › Other Locations › Connect to Server, then type the address with davs:// in place of https://."}}, Width: w.MenuWidth290, Pop: w.MenuUnderEnd}, "s-publish-32", "tc-menu", 1),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "upload", File: &w.FileField{Form: "note-edit", Name: "file", Label: "From this computer"}, Label: "From this computer"},
		w.MenuItem{Icon: "file", Label: "From Files"}}, Width: w.MenuWidth212, Pop: w.MenuUnder}, "s-notes-4", "tc-menu", 2),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Записи", Name: "into", Value: "nextcloud/zapisi", Required: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "nextcloud · Дневник", Name: "into", Value: "nextcloud/dnevnik", Required: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "nextcloud · Общие с семьёй", Reason: "Read-only: nothing can be written here."}}, Width: w.MenuWidth290, Pop: w.MenuUnderEnd}, "s-notes-34", "tc-menu", 1),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuText{Title: "Windows", Text: "File Explorer › This PC › Map network drive, then paste the address."},
		w.MenuText{Title: "macOS", Text: "Finder › Go › Connect to Server, then paste the address."},
		w.MenuText{Title: "Linux", Text: "Files › Other Locations › Connect to Server, then type the address with davs:// in place of https://."}}, Width: w.MenuWidth290, Pop: w.MenuUnder}, "s-publish-12", "tc-menu", 1),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "10 Aug", Variant: w.Mono11}, Label: "Синк-ап по релизу", Href: "/app/calendar/nextcloud/rabota/aa01"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}, "s-contacts-19", "tc-menu", 1),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Trail: w.Mono{Text: "last used", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "nextcloud · Записи", On: true, Name: "into", Value: "nextcloud/zapisi"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "nextcloud · Дневник", Name: "into", Value: "nextcloud/dnevnik"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "nextcloud · Общие с семьёй", Reason: "Read-only: nothing can be written here."}}, Width: w.MenuWidth290, Pop: w.MenuUnderEnd}, "s-notes-5", "tc-menu", 1),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Trail: w.Mono{Text: "last used", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "nextcloud · Личный", On: true, Name: "into", Value: "nextcloud/lichnyj"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа", Name: "into", Value: "nextcloud/rabota"},
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "nextcloud · Дни рождения", Name: "into", Value: "nextcloud/dni-rozhdeniya"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "radicale · vps · Проекты", Name: "into", Value: "radicale/proekty"},
		w.MenuItem{Mark: &w.Mark{Colour: "#8C6239"}, Label: "radicale · vps · Дежурства", Name: "into", Value: "radicale/dezhurstva"},
		w.MenuItem{Mark: &w.Mark{Colour: "#9A9280"}, Label: "baikal · Общий", Name: "into", Value: "baikal/obshchij"}}, Width: w.MenuWidth290, Pop: w.MenuUnderEnd}, "s-notes-9", "tc-menu", 1),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Trail: w.Mono{Text: "last used", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "nextcloud · Дела", On: true, Name: "into", Value: "nextcloud/dela"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа", Name: "into", Value: "nextcloud/rabota"}}, Width: w.MenuWidth290, Pop: w.MenuUnderEnd}, "s-notes-10", "tc-menu", 1),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Files", On: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "sftpgo · vps · Documents"},
		w.MenuItem{Mark: &w.Mark{Colour: "#9A9280"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "sftpgo · vps · Public files"}}, Width: w.MenuWidth290}, "s-backup-9", "tc-menu", 1),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Личный", On: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа"},
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "nextcloud · Дни рождения"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "radicale · vps · Проекты"},
		w.MenuItem{Mark: &w.Mark{Colour: "#8C6239"}, Label: "radicale · vps · Дежурства"},
		w.MenuItem{Mark: &w.Mark{Colour: "#9A9280"}, Label: "baikal · Общий"}}, Width: w.MenuWidth290}, "desktop-19", "tc-menu", 1),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "print", Label: "Print"},
		w.MenuItem{Icon: "dup", Label: "Open in Duplicates", Href: "/app/duplicates"},
		w.MenuSep{},
		w.MenuItem{Icon: "link", Label: "Unlink", Form: "post-unlink", Name: "group", Value: "g3"}}, Width: w.MenuWidth212}, "s-contacts-36", "tc-menu", 1),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Контакты", On: true, Href: "/app/contacts/nextcloud/kontakty/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "nextcloud · Архив", Href: "/app/contacts/nextcloud/arhiv/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "radicale · vps · Команда", Reason: "Read-only: nothing can be written here."},
		w.MenuSep{},
		w.MenuItem{Label: "New address book", Href: "/app/collections/new?kind=contact&back=/app/contacts/nextcloud/kontakty/import"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}, "s-io-3", "tc-menu", 1),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Дела", On: true, Href: "/app/tasks/nextcloud/dela/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа", Href: "/app/tasks/nextcloud/rabota/import"},
		w.MenuSep{},
		w.MenuItem{Label: "New task list", Href: "/app/collections/new?kind=task&back=/app/tasks/nextcloud/dela/import"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}, "s-io-16", "tc-menu", 1),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Личный", On: true, Href: "/app/calendar/nextcloud/lichnyj/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа", Href: "/app/calendar/nextcloud/rabota/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "nextcloud · Дни рождения", Href: "/app/calendar/nextcloud/dni-rozhdeniya/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "radicale · vps · Проекты", Href: "/app/calendar/radicale/proekty/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#8C6239"}, Label: "radicale · vps · Дежурства", Href: "/app/calendar/radicale/dezhurstva/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#9A9280"}, Label: "baikal · Общий", Href: "/app/calendar/baikal/obshchij/import"},
		w.MenuSep{},
		w.MenuItem{Label: "New calendar", Href: "/app/collections/new?kind=event&back=/app/calendar/nextcloud/lichnyj/import"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}, "s-io-15", "tc-menu", 1),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Контакты", Href: "/app/contacts/nextcloud/kontakty/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "nextcloud · Архив", Href: "/app/contacts/nextcloud/arhiv/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "radicale · vps · Команда", Reason: "Read-only: nothing can be written here."},
		w.MenuSep{},
		w.MenuItem{Label: "New address book", Href: "/app/collections/new?kind=contact&back=/app/contacts/import"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}, "s-io-17", "tc-menu", 1),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Recent searches", Form: w.RubricItem}},
		w.MenuItem{Icon: "clock", Trail: w.Mono{Text: "7 found · 2 min ago", Variant: w.Mono11DustyTight, Place: w.Pushed}, Label: "ковал", Href: "/app/search?q=%D0%BA%D0%BE%D0%B2%D0%B0%D0%BB"},
		w.MenuItem{Icon: "clock", Trail: w.Mono{Text: "nothing found · Monday", Variant: w.Mono11DustyTight, Place: w.Pushed}, Label: "дача", Href: "/app/search?q=%D0%B4%D0%B0%D1%87%D0%B0"}}, Width: w.MenuWidthFull}, "s-find-6", "tc-menu", 1),
	overlay("marks", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Дела", On: true, Href: "/app/tasks/nextcloud/dela/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа", Href: "/app/tasks/nextcloud/rabota/import"},
		w.MenuSep{},
		w.MenuItem{Label: "New task list", Href: "/app/collections/new?kind=task&back=/app/tasks/nextcloud/dela/import"}}, Width: w.MenuWidth290}, "s-io-13", "tc-menu", 1),
	overlay("submits", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "con", Label: "Contact", To: "/app/mail/send?op=attach&what=contact"},
		w.MenuItem{Icon: "con", Label: "Address book", To: "/app/mail/send?op=attach&what=book"},
		w.MenuItem{Icon: "cal", Label: "Event", To: "/app/mail/send?op=attach&what=event"},
		w.MenuItem{Icon: "cal", Label: "Calendar", To: "/app/mail/send?op=attach&what=calendar"},
		w.MenuItem{Icon: "task", Label: "Task", To: "/app/mail/send?op=attach&what=task"},
		w.MenuItem{Icon: "note", Label: "Note", To: "/app/mail/send?op=attach&what=note"},
		w.MenuSep{},
		w.MenuItem{Icon: "doc", Label: "File", To: "/app/mail/send?op=attach&what=file"},
		w.MenuItem{Icon: "file", Trail: w.Tag{Label: "link", Place: w.Pushed}, Label: "Folder", To: "/app/mail/send?op=attach&what=folder"}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}, "s-mail-1", "tc-menu", 2),
	overlay("danger", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "mail", Label: "Send", Reason: "It is no longer there"},
		w.MenuSep{},
		w.MenuItem{Label: "Revoke", Tone: w.MenuToneDanger, Form: "post-revoke", Name: "id", Value: "5d02b7"}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}, "s-publish-21", "tc-menu", 1),
	overlay("radio", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "7 days", Name: "lives", Value: "7"},
		w.MenuItem{Label: "14 days", On: true, Name: "lives", Value: "14"},
		w.MenuItem{Label: "30 days", Name: "lives", Value: "30"},
		w.MenuSep{},
		w.MenuItem{Label: "No expiry", Name: "lives", Value: "never"}}, Width: w.MenuWidth212, Pop: w.MenuUnder}, "s-publish-5", "tc-menu", 1),
	overlay("radio", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "7 days", On: true, Name: "lives", Value: "7"},
		w.MenuItem{Label: "14 days", Name: "lives", Value: "14"},
		w.MenuItem{Label: "30 days", Name: "lives", Value: "30"},
		w.MenuSep{},
		w.MenuItem{Label: "No expiry", Name: "lives", Value: "never"}}, Width: w.MenuWidth212, Pop: w.MenuUnder}, "s-publish-3", "tc-menu", 2),
	overlay("radio", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "7 days", Name: "lives", Value: "7"},
		w.MenuItem{Label: "14 days", Name: "lives", Value: "14"},
		w.MenuItem{Label: "30 days", On: true, Name: "lives", Value: "30"},
		w.MenuSep{},
		w.MenuItem{Label: "No expiry", Name: "lives", Value: "never"}}, Width: w.MenuWidth212, Pop: w.MenuUnder}, "s-publish-2", "tc-menu", 1),
	overlay("radio-skip", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "radicale · vps · Проекты", On: true, Name: "into.2", Value: "radicale|proekty"},
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Личный", Name: "into.2", Value: "nextcloud|lichnyj"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа", Name: "into.2", Value: "nextcloud|rabota"},
		w.MenuSep{},
		w.MenuItem{Label: "New collection", Name: "into.2", Value: "new"},
		w.MenuItem{Label: "Skip this one", Tone: w.MenuToneDusty, Name: "into.2", Value: "skip"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}, "s-backup-8", "tc-menu", 3),
	overlay("pop", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "sftpgo · vps", On: true, Href: "/app/files/picker?server=sftpgo", Swap: "upload-sheet"},
		w.MenuItem{Label: "nextcloud", Href: "/app/files/picker?server=nextcloud", Swap: "upload-sheet"},
		w.MenuItem{Label: "webdav · nas", Reason: "webdav · nas did not answer.", Href: "#"}}, Width: w.MenuWidth290, Pop: w.MenuUnderEnd}, "s-notes-11", "tc-menu", 1),
	overlay("danger", w.Menu{Items: []w.MenuPart{
		w.MenuSep{},
		w.MenuItem{Label: "Revoke", Tone: w.MenuToneDanger}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}, "s-publish-20", "tc-menu", 1),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "sftpgo · vps", On: true, Href: "/app/files/sftpgo/documents/op/carrel/attachments?op=move&server=sftpgo", Swap: "move-sheet"},
		w.MenuItem{Label: "nextcloud", Href: "/app/files/sftpgo/documents/op/carrel/attachments?op=move&server=nextcloud", Swap: "move-sheet"},
		w.MenuItem{Label: "webdav · nas", Reason: "webdav · nas did not answer.", Href: "#"}}, Width: w.MenuWidth290}, "s-files-26", "tc-menu", 1),
	overlay("radio", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "All actions", On: true, Name: "action", Value: ""},
		w.MenuSep{},
		w.MenuItem{Label: "login.ok", Name: "action", Value: "login.ok"},
		w.MenuItem{Label: "login.fail", Name: "action", Value: "login.fail"},
		w.MenuItem{Label: "escrow.recover", Name: "action", Value: "escrow.recover"},
		w.MenuItem{Label: "settings.instance", Name: "action", Value: "settings.instance"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}, "s-admin-7", "tc-menu", 1),
	overlay("pop", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Личный", On: true, Name: "calendar", Value: "nextcloud/lichnyj", Of: "event-edit"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа", Name: "calendar", Value: "nextcloud/rabota", Of: "event-edit"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}, "s-agenda-3", "tc-menu", 1),
	overlay("pop", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "Run"},
		w.MenuItem{Label: "Edit"},
		w.MenuItem{Label: "External trigger", Href: "/app/settings/backup/trigger?job=work"}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}, "s-backup-1", "tc-menu", 1),
	overlay("icons", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "move", Label: "Move"},
		w.MenuItem{Icon: "dup", Label: "Copy"},
		w.MenuItem{Icon: "export", Detail: &w.Mono{Text: ".vcf", Variant: w.Mono11}, Label: "Export"},
		w.MenuItem{Icon: "link", Label: "Publish a link"}}, Width: w.MenuWidth212}, "s-contacts-11", "tc-menu", 1),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "10 Aug", Variant: w.Mono11}, Label: "Синк-ап по релизу", Href: "/app/calendar/nextcloud/rabota/aa01"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "14 Aug", Variant: w.Mono11}, Label: "Ретро спринта", Href: "/app/calendar/nextcloud/rabota/aa02"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "17 Aug", Variant: w.Mono11}, Label: "Синк-ап по релизу", Href: "/app/calendar/nextcloud/rabota/aa04"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Detail: &w.Mono{Text: "19 Aug", Variant: w.Mono11}, Label: "Демо для команды", Href: "/app/calendar/radicale/proekty/aa05"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "24 Aug", Variant: w.Mono11}, Label: "Синк-ап по релизу", Href: "/app/calendar/nextcloud/rabota/aa06"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "26 Aug", Variant: w.Mono11}, Label: "Разбор инцидента", Href: "/app/calendar/nextcloud/rabota/aa07"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "31 Aug", Variant: w.Mono11}, Label: "Синк-ап по релизу", Href: "/app/calendar/nextcloud/rabota/aa08"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "2 Sep", Variant: w.Mono11}, Label: "Планирование Q4", Href: "/app/calendar/nextcloud/rabota/aa03"},
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Detail: &w.Mono{Text: "4 Sep", Variant: w.Mono11}, Label: "Подготовить отчёт · task", Href: "/app/tasks/nextcloud/dela/aa09"}}, Width: w.MenuWidth290}, "s-contacts-30", "tc-menu", 1),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#9A9280"}, Detail: &w.Mono{Text: "12 Aug", Variant: w.Mono11}, Label: "Общий созвон", Href: "/app/calendar/baikal/obshchij/bb01"},
		w.MenuItem{Mark: &w.Mark{Colour: "#9A9280"}, Detail: &w.Mono{Text: "21 Aug", Variant: w.Mono11}, Label: "Демо для заказчика", Href: "/app/calendar/baikal/obshchij/bb02"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}, "s-contacts-29", "tc-menu", 2),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "10 Aug", Variant: w.Mono11}, Label: "Синк-ап по релизу", Href: "/app/calendar/nextcloud/rabota/aa01"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "14 Aug", Variant: w.Mono11}, Label: "Ретро спринта", Href: "/app/calendar/nextcloud/rabota/aa02"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "2 Sep", Variant: w.Mono11}, Label: "Планирование Q4", Href: "/app/calendar/nextcloud/rabota/aa03"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}, "s-contacts-29", "tc-menu", 1),
	overlay("columns", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Columns", Form: w.RubricItem}, Form: w.MenuHeadingRoomy},
		w.MenuItem{Icon: "drag", Faint: true, Tick: &w.TickBox{}, Trail: w.Dim{Parts: []w.Inline{w.Text("always")}, Size: w.Dim10_5, Place: w.Pushed}, Label: "Time", Name: "col", Value: "time", Fixed: true},
		w.MenuItem{Icon: "drag", Faint: true, Tick: &w.TickBox{}, Trail: w.Dim{Parts: []w.Inline{w.Text("always")}, Size: w.Dim10_5, Place: w.Pushed}, Label: "Title", Name: "col", Value: "title", Fixed: true},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Where", Name: "col", Value: "where"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Attendees and status", Name: "col", Value: "status"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Source", Name: "col", Value: "source"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Tags", Name: "col", Value: "tags"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Changed", Name: "col", Value: "changed"},
		w.MenuFoot{Items: []w.Button{w.Button{Label: "Apply", Kind: w.ButtonPrimary, Small: true, Submit: true}, w.Button{Label: "Reset to defaults", Small: true, Name: "col", Value: "default", Submit: true}}}}, Action: "/app/calendar/columns", Hidden: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}}, "s-agenda-18", "tc-menu", 1),
	overlay("columns", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Columns", Form: w.RubricItem}, Form: w.MenuHeadingRoomy},
		w.MenuItem{Icon: "drag", Faint: true, Tick: &w.TickBox{}, Trail: w.Dim{Parts: []w.Inline{w.Text("always")}, Size: w.Dim10_5, Place: w.Pushed}, Label: "Name", Name: "col", Value: "name", Fixed: true},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Phone", Name: "col", Value: "phone"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Email", Name: "col", Value: "email"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Organisation", Name: "col", Value: "org"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Source", Name: "col", Value: "source"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Birthday", Name: "col", Value: "birthday"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Tags", Name: "col", Value: "tags"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Changed", Name: "col", Value: "changed"},
		w.MenuFoot{Items: []w.Button{w.Button{Label: "Apply", Kind: w.ButtonPrimary, Small: true, Submit: true}, w.Button{Label: "Reset to defaults", Small: true, Name: "col", Value: "default", Submit: true}}}}, Action: "/app/contacts/columns", Hidden: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}}, "s-contacts-28", "tc-menu", 1),
	overlay("picker", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "doc", File: &w.FileField{Form: "post-upload", Name: "file", Label: "Files", Size: "size", Many: true}, Label: "Files"},
		w.MenuItem{Icon: "file", File: &w.FileField{Form: "post-upload", Name: "file", Label: "Folder", Size: "size", Folder: true}, Label: "Folder"},
		w.MenuItem{Label: "Upload", NoScript: true, Form: "post-upload"}}, Width: w.MenuWidth212}, "s-files-8", "tc-menu", 1),
	overlay("danger", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "Revoke", Tone: w.MenuToneDanger}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}, "s-publish-17", "tc-menu", 1),
	overlay("radio", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "sftpgo · vps", On: true, Name: "server", Value: "sftpgo"},
		w.MenuItem{Label: "nextcloud", Name: "server", Value: "nextcloud"}}, Width: w.MenuWidth290}, "s-notes-29", "tc-menu", 1),
	overlay("more", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "move", Label: "Move", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012/move"},
		w.MenuItem{Icon: "dup", Label: "Copy", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012/copy"},
		w.MenuItem{Icon: "link", Label: "Publish a link", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012/publish"}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}, "s-notes-2", "tc-menu", 1),
	overlay("edit", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "upload", File: &w.FileField{Form: "", Name: "file", Label: "From this computer", List: "event-attachments"}, Row: w.Line{Items: []w.LinePart{w.Faced{Face: w.FaceLarge, Items: []w.LinePart{w.Text("{name} "), w.Dim{Parts: []w.Inline{w.Text("{size}")}}}}, w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Removes: true}}, Gap: w.LineGap8}, Label: "From this computer"},
		w.MenuItem{Icon: "file", Label: "From Files"}}, Width: w.MenuWidth212, Pop: w.MenuUnder}, "s-agenda-3", "tc-menu", 4),
	overlay("edit", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "upload", File: &w.FileField{Form: "note-edit", Name: "file", Label: "From this computer", List: "note-attachments"}, Row: w.DocAttach{Name: "{name}", Size: w.Dim{Parts: []w.Inline{w.Text("· {size}")}}, Remove: &w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Removes: true}}, Label: "From this computer"},
		w.MenuItem{Icon: "file", Label: "From Files"}}, Width: w.MenuWidth212, Pop: w.MenuUnder}, "s-notes-20", "tc-menu", 2),
	overlay("row", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "mail", Label: "Send", Href: "/app/mail/link?id=9c1f4e"},
		w.MenuSep{},
		w.MenuItem{Label: "Revoke", Tone: w.MenuToneDanger, Form: "post-revoke", Name: "id", Value: "9c1f4e"}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}, "s-publish-1", "tc-menu", 1),
	overlay("sort", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "Name", On: true, Href: "/app/files/sftpgo/documents?sort=name"},
		w.MenuItem{Label: "Size", Href: "/app/files/sftpgo/documents?sort=size"},
		w.MenuItem{Label: "Changed", Href: "/app/files/sftpgo/documents?sort=changed"}}, Pop: w.MenuUnderEnd, Inline: true}, "s-files-1", "tc-menu", 1),
	overlay("more", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "cols", Label: "Columns"}}, Pop: w.MenuUnderEnd, Inline: true}, "s-contacts-7", "tc-menu", 1),
	overlay("radio", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Записи", On: true, Name: "notebook", Value: "nextcloud/zapisi", Of: "note-edit"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "nextcloud · Дневник", Name: "notebook", Value: "nextcloud/dnevnik", Of: "note-edit"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "nextcloud · Общие с семьёй", Reason: "Read-only: this account cannot write here", Name: "notebook", Value: "nextcloud/semya", Of: "note-edit"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}, "s-notes-20", "tc-menu", 1),
	overlay("marks", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Trail: w.Mono{Text: "last used", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "nextcloud · Дела", On: true, Name: "into", Value: "nextcloud/dela"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа", Name: "into", Value: "nextcloud/rabota"}}, Width: w.MenuWidth290}, "s-notes-19", "tc-menu", 1),
	overlay("picker", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "upload", File: &w.FileField{Form: "post-attach", Name: "file", Label: "From this computer"}, Label: "From this computer"},
		w.MenuItem{Icon: "file", Label: "From Files"},
		w.MenuItem{Label: "Upload", NoScript: true, Form: "post-attach"}}, Width: w.MenuWidth212}, "s-agenda-6", "tc-menu", 1),
	overlay("closed", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "move", Label: "Move"},
		w.MenuItem{Icon: "dup", Label: "Copy"},
		w.MenuItem{Icon: "export", Detail: &w.Mono{Text: ".vcf", Variant: w.Mono11}, Label: "Export"},
		w.MenuItem{Icon: "link", Label: "Publish a link"}}, Width: w.MenuWidth212, Pop: w.MenuUnder}, "s-contacts-7", "tc-menu", 3),
	overlay("closed-end", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "print", Label: "Print", Tertiary: true},
		w.MenuItem{Icon: "move", Label: "Move"},
		w.MenuItem{Icon: "dup", Label: "Copy"},
		w.MenuItem{Icon: "export", Detail: &w.Mono{Text: ".vcf", Variant: w.Mono11}, Label: "Export"},
		w.MenuItem{Icon: "link", Label: "Publish a link"}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}, "s-contacts-6", "tc-menu", 1),
	overlay("long", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "contact"},
		w.MenuItem{Label: "acquaintance"},
		w.MenuItem{Label: "friend"},
		w.MenuItem{Label: "met"},
		w.MenuItem{Label: "co-worker"},
		w.MenuItem{Label: "colleague", On: true},
		w.MenuItem{Label: "co-resident"},
		w.MenuItem{Label: "neighbor"},
		w.MenuItem{Label: "child"},
		w.MenuItem{Label: "parent"},
		w.MenuItem{Label: "sibling"},
		w.MenuItem{Label: "spouse"},
		w.MenuItem{Label: "kin"},
		w.MenuItem{Label: "muse"},
		w.MenuItem{Label: "crush"},
		w.MenuItem{Label: "date"},
		w.MenuItem{Label: "sweetheart"},
		w.MenuItem{Label: "me"},
		w.MenuItem{Label: "agent"},
		w.MenuItem{Label: "emergency"}}}, "mentions-8", "tc-menu", 1),
	overlay("timeline-cols", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Columns", Form: w.RubricItem}, Form: w.MenuHeadingRoomy},
		w.MenuItem{Icon: "drag", Faint: true, Tick: &w.TickBox{}, Trail: w.Dim{Parts: []w.Inline{w.Text("always")}, Size: w.Dim10_5, Place: w.Pushed}, Label: "Title", Name: "col", Value: "title", Fixed: true},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Date", Name: "col", Value: "date"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Kind", Name: "col", Value: "kind"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Source", Name: "col", Value: "source"},
		w.MenuFoot{Items: []w.Button{w.Button{Label: "Apply", Kind: w.ButtonPrimary, Small: true}, w.Button{Label: "Reset to defaults", Small: true, Name: "col", Value: "default"}}}}, Action: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/timeline/columns", Hidden: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}}, "s-contacts-25", "tc-menu", 1),
	overlay("cols-off3", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Columns", Form: w.RubricItem}, Form: w.MenuHeadingRoomy},
		w.MenuItem{Icon: "drag", Faint: true, Tick: &w.TickBox{}, Trail: w.Dim{Parts: []w.Inline{w.Text("always")}, Size: w.Dim10_5, Place: w.Pushed}, Label: "Title", Name: "col", Value: "title", Fixed: true},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Due", Name: "col", Value: "due"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Complete", Name: "col", Value: "complete"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Source", Name: "col", Value: "source"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Tags", Name: "col", Value: "tags"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Priority", Name: "col", Value: "priority"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Status", Name: "col", Value: "status"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Changed", Name: "col", Value: "changed"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Where", Name: "col", Value: "where"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Attendees", Name: "col", Value: "attendees"},
		w.MenuFoot{Items: []w.Button{w.Button{Label: "Apply", Kind: w.ButtonPrimary, Small: true, Submit: true}, w.Button{Label: "Reset to defaults", Small: true, Name: "col", Value: "default", Submit: true}}}}, Action: "/app/tasks/columns", Hidden: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}}, "s-tasks-12", "tc-menu", 1),
	overlay("cols-off1", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Columns", Form: w.RubricItem}, Form: w.MenuHeadingRoomy},
		w.MenuItem{Icon: "drag", Faint: true, Tick: &w.TickBox{}, Trail: w.Dim{Parts: []w.Inline{w.Text("always")}, Size: w.Dim10_5, Place: w.Pushed}, Label: "Title", Name: "col", Value: "title", Fixed: true},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Excerpt", Name: "col", Value: "excerpt"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Date", Name: "col", Value: "date"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Source", Name: "col", Value: "source"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Tags", Name: "col", Value: "tags"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Changed", Name: "col", Value: "changed"},
		w.MenuFoot{Items: []w.Button{w.Button{Label: "Apply", Kind: w.ButtonPrimary, Small: true, Submit: true}, w.Button{Label: "Reset to defaults", Small: true, Name: "col", Value: "default", Submit: true}}}}, Action: "/app/notes/columns", Hidden: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}}, "s-notes-23", "tc-menu", 1),
	overlay("log-cols", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Columns", Form: w.RubricItem}, Form: w.MenuHeadingRoomy},
		w.MenuItem{Icon: "drag", Faint: true, Tick: &w.TickBox{}, Trail: w.Dim{Parts: []w.Inline{w.Text("always")}, Size: w.Dim10_5, Place: w.Pushed}, Label: "When", Name: "col", Value: "when", Fixed: true},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Action", Name: "col", Value: "action"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Actor", Name: "col", Value: "actor"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Target", Name: "col", Value: "target"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Detail", Name: "col", Value: "detail"},
		w.MenuFoot{Items: []w.Button{w.Button{Label: "Apply", Kind: w.ButtonPrimary, Small: true}, w.Button{Label: "Reset to defaults", Small: true, Name: "col", Value: "default"}}}}, Action: "/admin/audit/columns", Hidden: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}}, "s-admin-23", "tc-menu", 1),
	overlay("users-cols", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Columns", Form: w.RubricItem}, Form: w.MenuHeadingRoomy},
		w.MenuItem{Icon: "drag", Faint: true, Tick: &w.TickBox{}, Trail: w.Dim{Parts: []w.Inline{w.Text("always")}, Size: w.Dim10_5, Place: w.Pushed}, Label: "Login", Name: "col", Value: "login", Fixed: true},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Role", Name: "col", Value: "role"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Created", Name: "col", Value: "created"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Last login", Name: "col", Value: "last-login"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "DAV", Name: "col", Value: "dav"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Sessions", Name: "col", Value: "sessions"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Escrow", Name: "col", Value: "escrow"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Address", Name: "col", Value: "address"},
		w.MenuFoot{Items: []w.Button{w.Button{Label: "Apply", Kind: w.ButtonPrimary, Small: true}, w.Button{Label: "Reset to defaults", Small: true, Name: "col", Value: "default"}}}}, Action: "/admin/columns", Hidden: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}}, "s-admin-21", "tc-menu", 1),
	overlay("form", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Columns", Form: w.RubricItem}, Form: w.MenuHeadingRoomy},
		w.MenuItem{Icon: "drag", Faint: true, Tick: &w.TickBox{}, Trail: w.Dim{Parts: []w.Inline{w.Text("always")}, Size: w.Dim10_5, Place: w.Pushed}, Label: "Name", Name: "col", Value: "name", Fixed: true},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Email", Name: "col", Value: "email"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Birthday", Name: "col", Value: "birthday"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Tags", Name: "col", Value: "tags", Drag: true},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Source", Name: "col", Value: "source"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Phone", Name: "col", Value: "phone"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Organisation", Name: "col", Value: "org"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Changed", Name: "col", Value: "changed"},
		w.MenuFoot{Items: []w.Button{w.Button{Label: "Apply", Kind: w.ButtonPrimary, Small: true}, w.Button{Label: "Reset to defaults", Small: true, Name: "col", Value: "default"}}}}, Action: "/app/contacts/nextcloud/kontakty/columns", Hidden: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}}, "components-42", "tc-menu", 1),
	overlay("three", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "Monday", On: true},
		w.MenuItem{Label: "Sunday"},
		w.MenuItem{Label: "Saturday"}}, Width: w.MenuWidth212}, "appearance-2", "tc-menu", 1),
	overlay("pop", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "upload", File: &w.FileField{Form: "post-attach", Name: "file", Label: "From this computer"}, Label: "From this computer"},
		w.MenuItem{Icon: "file", Label: "From Files"},
		w.MenuItem{Label: "Upload", NoScript: true, Form: "post-attach"}}, Width: w.MenuWidth212, Pop: w.MenuUnder}, "s-agenda-1", "tc-menu", 3),
	overlay("links", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "pencil", Label: "Rename and colour", Href: "/app/collections/rename?c=lichnyj"},
		w.MenuItem{Icon: "link", Label: "Publish a link", Href: "/app/publish?c=lichnyj"}}, Width: w.MenuWidth212}, "components-39", "tc-menu", 1),
	overlay("after-rule", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "home"},
		w.MenuItem{Label: "work"},
		w.MenuItem{Label: "other"},
		w.MenuSep{},
		w.MenuItem{Label: "дача", On: true}}, Width: w.MenuWidth212}, "s-contacts-22", "tc-menu", 1),
	overlay("disabled-icons", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "pencil", Label: "Rename", Reason: "This is the whole storage, not a folder in it."},
		w.MenuItem{Icon: "move", Label: "Move to another server", Href: "/app/collections/move?a=nextcloud&c=files"},
		w.MenuItem{Icon: "dup", Label: "Copy to another server", Href: "/app/collections/move?a=nextcloud&c=files&copy=1"},
		w.MenuItem{Icon: "link", Label: "Publish a link"},
		w.MenuSep{},
		w.MenuItem{Icon: "trash", Label: "Delete on the server…", Tone: w.MenuToneDanger, Reason: "This is the whole storage, not a folder in it."}}, Width: w.MenuWidth212}, "s-profile-17", "tc-menu", 1),
	overlay("disabled", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "anna · User", On: true},
		w.MenuItem{Label: "olga · User"},
		w.MenuItem{Label: "viktor · User", Reason: "Enable the account first: a disabled account gets no reset."},
		w.MenuItem{Label: "mix · Administrator", Reason: "Your own password is changed in your account settings."}}, Width: w.MenuWidth290}, "s-admin-15", "tc-menu", 1),
	overlay("pop", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Trail: w.Dim{Parts: []w.Inline{w.Text("event · 12 Aug")}, Size: w.Dim11, Place: w.Pushed}, Label: "Поездка в Казань"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Trail: w.Dim{Parts: []w.Inline{w.Text("task · 12 Aug")}, Size: w.Dim11, Place: w.Pushed}, Label: "Собрать чемодан"},
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Trail: w.Dim{Parts: []w.Inline{w.Text("note · 10 Aug")}, Size: w.Dim11, Place: w.Pushed}, Label: "Заметка со встречи 10.08"},
		w.MenuSep{},
		w.MenuItem{Label: "By UID…", Tone: w.MenuToneDusty}}, Width: w.MenuWidth290, Pop: w.MenuUnder}, "s-tasks-10", "tc-menu", 1),
	overlay("fields", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Points per signal", Items: []w.RubricPart{w.Push{Items: []w.PushPart{w.Text("Default")}}, w.RubricCol{Text: "Yours"}}, Form: w.RubricItem}, Form: w.MenuHeadingRoomy},
		w.MenuField{Label: "Same UID", Default: w.Dim{Parts: []w.Inline{w.Text("80")}, Size: w.Dim10_5, Place: w.Pushed}, Input: w.Input{Value: "80", Mono: true, Variant: w.InputChip, Label: "Same UID"}},
		w.MenuField{Label: "Same email", Default: w.Dim{Parts: []w.Inline{w.Text("60")}, Size: w.Dim10_5, Place: w.Pushed}, Input: w.Input{Value: "60", Mono: true, Variant: w.InputChip, Label: "Same email"}},
		w.MenuField{Label: "Same phone", Default: w.Dim{Parts: []w.Inline{w.Text("60")}, Size: w.Dim10_5, Place: w.Pushed}, Input: w.Input{Value: "60", Mono: true, Variant: w.InputChip, Label: "Same phone"}},
		w.MenuField{Label: "Same name", Default: w.Dim{Parts: []w.Inline{w.Text("20")}, Size: w.Dim10_5, Place: w.Pushed}, Input: w.Input{Value: "20", Mono: true, Variant: w.InputChip, Label: "Same name"}},
		w.MenuField{Label: "Same birthday", Default: w.Dim{Parts: []w.Inline{w.Text("10")}, Size: w.Dim10_5, Place: w.Pushed}, Input: w.Input{Value: "10", Mono: true, Variant: w.InputChip, Label: "Same birthday"}},
		w.MenuField{Label: "Same title and date (events, tasks, notes)", Default: w.Dim{Parts: []w.Inline{w.Text("70")}, Size: w.Dim10_5, Place: w.Pushed}, Input: w.Input{Value: "70", Mono: true, Variant: w.InputChip, Label: "Same title and date (events, tasks, notes)"}},
		w.MenuFoot{Items: []w.Button{w.Button{Label: "Save", Kind: w.ButtonPrimary, Small: true, Name: "do", Value: "weights", Submit: true}, w.Button{Label: "Back to defaults", Small: true, Title: "Every weight is the default already.", Name: "do", Value: "defaults", Submit: true, State: w.ButtonDisabled}}}}, Width: w.MenuWidth420, Fields: true, Action: "/app/duplicates/decide", Hidden: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}}, "s-dup-7", "tc-menu", 1),
	overlay("icons", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "print", Label: "Print", Tertiary: true},
		w.MenuItem{Icon: "move", Label: "Move"},
		w.MenuItem{Icon: "dup", Label: "Copy"},
		w.MenuItem{Icon: "export", Detail: &w.Mono{Text: ".vcf", Variant: w.Mono11}, Label: "Export"},
		w.MenuItem{Icon: "link", Label: "Publish a link"}}, Width: w.MenuWidth212}, "s-contacts-8", "tc-menu", 1),
	overlay("danger", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "con", Label: "Open"},
		w.MenuItem{Icon: "pencil", Label: "Edit"},
		w.MenuItem{Icon: "move", Label: "Move"},
		w.MenuItem{Icon: "dup", Label: "Copy"},
		w.MenuItem{Icon: "export", Detail: &w.Mono{Text: ".vcf", Variant: w.Mono11}, Label: "Export"},
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
		w.MenuItem{Icon: "download", Label: "Download"},
		w.MenuItem{Icon: "pencil", Label: "Rename"},
		w.MenuItem{Icon: "move", Label: "Move"},
		w.MenuItem{Icon: "dup", Label: "Copy"},
		w.MenuItem{Icon: "link", Label: "Publish a link"},
		w.MenuSep{},
		w.MenuItem{Icon: "trash", Label: "Delete", Tone: w.MenuToneDanger}}, Width: w.MenuWidth212}, "components-5", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "work"},
		w.MenuItem{Label: "home"},
		w.MenuItem{Label: "other"},
		w.MenuSep{},
		w.MenuItem{Label: "Личная", On: true}}, Width: w.MenuWidth212}, "s-contacts-13", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "At the time"},
		w.MenuItem{Label: "5 minutes before"},
		w.MenuItem{Label: "10 minutes before", On: true},
		w.MenuItem{Label: "30 minutes before"},
		w.MenuItem{Label: "1 hour before"},
		w.MenuItem{Label: "1 day before"},
		w.MenuItem{Label: "Custom…"},
		w.MenuSep{},
		w.MenuItem{Label: "No reminder"}}, Width: w.MenuWidth212}, "s-agenda-9", "tc-menu", 1),
	overlay("marks", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Записи", On: true, Name: "notebook", Value: "nextcloud/zapisi"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "nextcloud · Дневник", Name: "notebook", Value: "nextcloud/dnevnik"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "nextcloud · Общие с семьёй", Reason: "Read-only: this account cannot write here", Name: "notebook", Value: "nextcloud/semya"}}, Width: w.MenuWidth290}, "s-notes-15", "tc-menu", 1),
	overlay("marks", w.Menu{Items: []w.MenuPart{
		w.Rubric{Text: "Write into", Form: w.RubricInline},
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Trail: w.Mono{Text: "last used", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Дела", On: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "Работа"}}, Width: w.MenuWidth290}, "s-tasks-5", "tc-menu", 1),
	overlay("icons", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "print", Label: "Print", Tertiary: true},
		w.MenuItem{Icon: "move", Label: "Move"},
		w.MenuItem{Icon: "dup", Label: "Copy"},
		w.MenuItem{Icon: "export", Detail: &w.Mono{Text: ".ics", Variant: w.Mono11}, Label: "Export"},
		w.MenuItem{Icon: "link", Label: "Publish a link"}}, Width: w.MenuWidth212}, "s-notes-14", "tc-menu", 1),
	overlay("marks", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Trail: w.Mono{Text: "last used", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "nextcloud · Контакты", On: true, Name: "into", Value: "nextcloud/kontakty"},
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "nextcloud · Архив", Name: "into", Value: "nextcloud/arhiv"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "radicale · vps · Команда", Reason: "Read-only: nothing can be written here."}}, Width: w.MenuWidth290}, "s-notes-16", "tc-menu", 1),
	overlay("marks", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Trail: w.Mono{Text: "last used", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "nextcloud · Личный", On: true, Name: "into", Value: "nextcloud/lichnyj"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа", Name: "into", Value: "nextcloud/rabota"},
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "nextcloud · Дни рождения", Name: "into", Value: "nextcloud/dni-rozhdeniya"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "radicale · vps · Проекты", Name: "into", Value: "radicale/proekty"},
		w.MenuItem{Mark: &w.Mark{Colour: "#8C6239"}, Label: "radicale · vps · Дежурства", Name: "into", Value: "radicale/dezhurstva"},
		w.MenuItem{Mark: &w.Mark{Colour: "#9A9280"}, Label: "baikal · Общий", Name: "into", Value: "baikal/obshchij"}}, Width: w.MenuWidth290}, "s-notes-17", "tc-menu", 1),
	overlay("marks", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Записи", On: true, Href: "/app/notes/nextcloud/zapisi/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "nextcloud · Дневник", Href: "/app/notes/nextcloud/dnevnik/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "nextcloud · Общие с семьёй", Reason: "Read-only: nothing can be written here."},
		w.MenuSep{},
		w.MenuItem{Label: "New notebook", Href: "/app/collections/new?kind=note&back=/app/notes/nextcloud/zapisi/import"}}, Width: w.MenuWidth290}, "s-io-11", "tc-menu", 1),
	overlay("pop", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "pencil", Label: "Rename"},
		w.MenuItem{Icon: "export", Detail: &w.Mono{Text: ".vcf", Variant: w.Mono11}, Label: "Export"}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}, "components-19", "tc-menu", 1),
	overlay("pop", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "All", On: true},
		w.MenuItem{Label: "Overdue"}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}, "components-19", "tc-menu", 2),
	overlay("heading", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Write into", Form: w.RubricItem}},
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Trail: w.Mono{Text: "last used", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Контакты", On: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "Архив"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "Команда"}}}, "components-6", "tc-menu", 1),
	overlay("icons", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "import", Label: "Import"},
		w.MenuItem{Icon: "export", Label: "Export"},
		w.MenuItem{Icon: "print", Label: "Print"}}}, "components-8", "tc-menu", 1),
	overlay("heading", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Columns", Form: w.RubricItem}, Form: w.MenuHeadingRoomy},
		w.MenuItem{Icon: "drag", Faint: true, Tick: &w.TickBox{}, Trail: w.Dim{Parts: []w.Inline{w.Text("always")}, Size: w.Dim10_5, Place: w.Pushed}, Label: "Name"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Phone"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Email"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{}, Label: "Organisation"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Birthday"},
		w.MenuItem{Icon: "drag", Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Tags"},
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
		w.MenuItem{Label: "mobile"},
		w.MenuItem{Label: "work"},
		w.MenuItem{Label: "home"},
		w.MenuItem{Label: "main"},
		w.MenuItem{Label: "fax"},
		w.MenuItem{Label: "other"},
		w.MenuSep{},
		w.MenuItem{Label: "iPhone", On: true}}, Width: w.MenuWidth212}, "s-contacts-12", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "cols", Label: "Columns"}}, Width: w.MenuWidth212}, "s-contacts-14", "tc-menu", 1),
	overlay("pop", w.Menu{Items: []w.MenuPart{
		w.Range{From: w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true}, To: w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-16", Label: "To", Mono: true}, Show: w.Button{Label: "Show", Submit: true}, Hidden: []w.Hidden{{Name: "view", Value: "range"}}, Action: "/app/calendar", Wrap: true},
		w.MenuSep{},
		w.CheckForm{Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}}, Action: "/app/calendar", Apply: w.Button{Label: "Apply", Small: true, Submit: true}}}, Pop: w.MenuUnderEnd, Inline: true}, "s-agenda-1", "tc-menu", 1),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.Range{From: w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true}, To: w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-16", Label: "To", Mono: true}, Show: w.Button{Label: "Show", Submit: true}, Hidden: []w.Hidden{{Name: "view", Value: "range"}}, Action: "/app/calendar", Wrap: true},
		w.MenuSep{},
		w.CheckForm{Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}}, Action: "/app/calendar", Apply: w.Button{Label: "Apply", Small: true, Submit: true}}}}, "s-agenda-16", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "No status"},
		w.MenuSep{},
		w.MenuItem{Label: "Tentative"},
		w.MenuItem{Label: "Confirmed", On: true},
		w.MenuItem{Label: "Cancelled"}}, Width: w.MenuWidth212}, "s-agenda-7", "tc-menu", 1),
	overlay("radio", w.Menu{Items: []w.MenuPart{
		w.Input{Kind: w.InputSearch, Placeholder: "Find a zone", Label: "Find a zone"},
		w.Rubric{Text: "As sent", Form: w.RubricInline},
		w.MenuItem{Label: "Asia/Yekaterinburg", On: true, Name: "tz", Value: "Asia/Yekaterinburg"},
		w.MenuSep{},
		w.MenuItem{Label: "Europe/Moscow", Name: "tz", Value: "Europe/Moscow"},
		w.MenuItem{Label: "UTC", Name: "tz", Value: "UTC"},
		w.MenuSep{},
		w.MenuItem{Label: "Floating — the same clock time anywhere", Name: "tz", Value: "floating"}}, Width: w.MenuWidth290}, "s-agenda-8", "tc-menu", 1),
	// The owner's decision of 02.10: the sources that did not answer, in full, under the runbar's button.
	overlay("refusals", w.Menu{Items: []w.MenuPart{
		w.MenuRefusal{Mark: w.Mark{Colour: "#9A9280"}, Source: "baikal · Общий", Answer: "PROPFIND /dav.php/calendars/mix/ → 401 Unauthorized: the saved password was refused. Open Connections to update it.", Retry: w.Button{Label: "Retry", Small: true}}}, Width: w.MenuWidth420, Pop: w.MenuUnderEnd}, "s-contacts-7", "tc-menu", 2),
	overlay("refusals", w.Menu{Items: []w.MenuPart{
		w.MenuRefusal{Mark: w.Mark{Colour: "#8C6239"}, Source: "radicale · vps · Дежурства", Answer: "REPORT /dav/calendars/duty/ → 503 Service Unavailable", Retry: w.Button{Label: "Retry", Small: true}},
		w.MenuRefusal{Mark: w.Mark{Colour: "#6B4E7A"}, Source: "radicale · vps · Проекты", Answer: "REPORT /dav/calendars/projects/ → 503 Service Unavailable", Retry: w.Button{Label: "Retry", Small: true}},
		w.MenuRefusal{Mark: w.Mark{Colour: "#9A9280"}, Source: "baikal · Общий", Answer: "PROPFIND /dav.php/calendars/mix/ → 401 Unauthorized: the saved password was refused. Open Connections to update it.", Retry: w.Button{Label: "Retry", Small: true}},
		w.MenuFoot{Items: []w.Button{w.Button{Label: "Retry all", Small: true}}}}, Width: w.MenuWidth420}, "states-16", "tc-menu", 1),
	// Carrel's task 82: the list's «⋯» — the sort and Compact rows, twins of what goes under it.
	overlay("pop", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "A–Z", On: true, Href: "/app/contacts?sort=az"},
		w.MenuItem{Label: "Recently changed", Href: "/app/contacts?sort=changed"},
		w.MenuSep{},
		w.CheckForm{Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}}, Action: "/app/contacts", Apply: w.Button{Label: "Apply", Small: true}}}, Pop: w.MenuUnderEnd, Inline: true}, "s-contacts-1", "tc-menu", 1),
	// Carrel's task 77: the suggestions of attendees — people of the ticked books, or the address typed.
	overlay("found", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Found: []w.Inline{w.Strong("Ве"), w.Text("ра Лебедева")}, Detail: &w.Mono{Text: "vera@example.net", Variant: w.Mono11}, Trail: w.Mono{Text: "Архив", Variant: w.Mono10Dusty, Place: w.Pushed}, Adds: &w.Tag{Label: "Вера Лебедева", Mark: &w.Mark{Colour: "#A05A2C", Size: w.MarkH11}, Removable: true, Title: "vera@example.net", Name: "attendee", Value: "mailto:vera@example.net"}},
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Found: []w.Inline{w.Strong("Ве"), w.Text("ра Лебедева")}, Detail: &w.Mono{Text: "vera.l@work.example", Variant: w.Mono11}, Trail: w.Mono{Text: "Архив", Variant: w.Mono10Dusty, Place: w.Pushed}, Adds: &w.Tag{Label: "Вера Лебедева", Mark: &w.Mark{Colour: "#A05A2C", Size: w.MarkH11}, Removable: true, Title: "vera.l@work.example", Name: "attendee", Value: "mailto:vera.l@work.example"}},
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Found: []w.Inline{w.Text("Лебедева "), w.Strong("Ве"), w.Text("ра")}, Detail: &w.Mono{Text: "v.lebedeva@example.org", Variant: w.Mono11}, Trail: w.Mono{Text: "Контакты", Variant: w.Mono10Dusty, Place: w.Pushed}, Adds: &w.Tag{Label: "Лебедева Вера", Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, Removable: true, Title: "v.lebedeva@example.org", Name: "attendee", Value: "mailto:v.lebedeva@example.org"}}}, Width: w.MenuWidth420}, "s-agenda-12", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "plus", Label: "Add", Detail: &w.Mono{Text: "ivan.petrov@mail.example", Variant: w.Mono11}, On: true,
			Adds: &w.Tag{Address: &w.Mono{Text: "ivan.petrov@mail.example"}, Removable: true, Title: "ivan.petrov@mail.example", Name: "attendee", Value: "mailto:ivan.petrov@mail.example"}},
		w.MenuSep{},
		w.MenuItem{Label: "Not in the ticked address books.", Tone: w.MenuToneAside}}, Width: w.MenuWidth420}, "s-agenda-13", "tc-menu", 1),
	overlay("pop", w.Menu{Width: w.MenuWidth420, Pop: w.MenuUnder}, "s-agenda-3", "tc-menu", 3),
	overlay("pop", w.Menu{Items: []w.MenuPart{
		w.Input{Kind: w.InputSearch, Placeholder: "Find a zone", Label: "Find a zone"},
		w.Rubric{Text: "As sent", Form: w.RubricInline},
		w.MenuItem{Label: "Asia/Yekaterinburg", On: true, Name: "tz", Value: "Asia/Yekaterinburg"},
		w.MenuSep{},
		w.MenuItem{Label: "Europe/Moscow", Name: "tz", Value: "Europe/Moscow"},
		w.MenuItem{Label: "UTC", Name: "tz", Value: "UTC"},
		w.MenuSep{},
		w.MenuItem{Label: "Floating — the same clock time anywhere", Name: "tz", Value: "floating"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}, "s-agenda-3", "tc-menu", 2),
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
		w.MenuItem{Icon: "export", Detail: &w.Mono{Text: ".md", Variant: w.Mono11}, Label: "Export"},
		w.MenuItem{Icon: "link", Label: "Publish a link"},
		w.MenuSep{},
		w.MenuItem{Icon: "trash", Label: "Delete note", Tone: w.MenuToneDanger}}, Width: w.MenuWidth212}, "s-notes-12", "tc-menu", 1),
	overlay("pop", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Trail: w.Dim{Parts: []w.Inline{w.Text("event · 10 Aug")}, Size: w.Dim11, Place: w.Pushed}, Label: "Звонок с Ковалёвой", On: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Trail: w.Dim{Parts: []w.Inline{w.Text("task · 14 Aug")}, Size: w.Dim11, Place: w.Pushed}, Label: "Ответить Анне по смете"},
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Trail: w.Dim{Parts: []w.Inline{w.Text("contact")}, Size: w.Dim11, Place: w.Pushed}, Label: "Анна Ковалёва"},
		w.MenuSep{},
		w.MenuItem{Label: "By UID…", Tone: w.MenuToneDusty}}, Width: w.MenuWidthFull, Pop: w.MenuUnder}, "s-notes-4", "tc-menu", 3),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "quill", Trail: w.Kbd{Keys: "N"}, Label: "Note", On: true},
		w.MenuSep{},
		w.MenuItem{Icon: "con", Trail: w.Kbd{Keys: "C"}, Label: "Contact"},
		w.MenuItem{Icon: "cal", Trail: w.Kbd{Keys: "E"}, Label: "Event"},
		w.MenuItem{Icon: "task", Trail: w.Kbd{Keys: "T"}, Label: "Task"},
		w.MenuSep{},
		w.MenuItem{Icon: "upload", Label: "Upload file"},
		w.MenuSep{},
		w.MenuItem{Icon: "mail", Label: "Letter"}}}, "s-notes-7", "tc-menu", 1),
	overlay("heading", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Recent searches", Form: w.RubricItem}},
		w.MenuItem{Icon: "clock", Trail: w.Mono{Text: "7 found · 2 min ago", Variant: w.Mono11DustyTight, Place: w.Pushed}, Label: "ковал", Href: "/app/search?q=%D0%BA%D0%BE%D0%B2%D0%B0%D0%BB"},
		w.MenuItem{Icon: "clock", Trail: w.Mono{Text: "3 found · yesterday", Variant: w.Mono11DustyTight, Place: w.Pushed}, Label: "смета ковалёвой", Href: "/app/search?q=%D1%81%D0%BC%D0%B5%D1%82%D0%B0+%D0%BA%D0%BE%D0%B2%D0%B0%D0%BB%D1%91%D0%B2%D0%BE%D0%B9"},
		w.MenuItem{Icon: "clock", Trail: w.Mono{Text: "nothing found · Monday", Variant: w.Mono11DustyTight, Place: w.Pushed}, Label: "дача", Href: "/app/search?q=%D0%B4%D0%B0%D1%87%D0%B0"},
		w.MenuItem{Icon: "clock", Trail: w.Mono{Text: "no source answered · Sunday", Variant: w.Mono11DustyTight, Place: w.Pushed}, Label: "договор", Href: "/app/search?q=%D0%B4%D0%BE%D0%B3%D0%BE%D0%B2%D0%BE%D1%80"},
		w.MenuSep{},
		w.MenuItem{Trail: w.Kbd{Keys: "Enter"}, Label: "Type to search everything", Tone: w.MenuToneDusty}}, Width: w.MenuWidthFull}, "s-find-2", "tc-menu", 1),
	overlay("found", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Jump to", Form: w.RubricItem}},
		w.MenuItem{Icon: "con", Mark: &w.Mark{Colour: "#4A6B52"}, Found: []w.Inline{w.Text("Анна "), w.Strong("Ковал"), w.Text("ёва")}, Trail: w.Dim{Parts: []w.Inline{w.Text("nextcloud · Контакты")}, Size: w.Dim11, Variant: w.DimTight, Place: w.Pushed}},
		w.MenuItem{Icon: "cal", Mark: &w.Mark{Colour: "#4A6B52"}, Found: []w.Inline{w.Text("Звонок с "), w.Strong("Ковал"), w.Text("ёвой")}, Trail: w.Dim{Parts: []w.Inline{w.Text("10 Aug · nextcloud · Личный")}, Size: w.Dim11, Variant: w.DimTight, Place: w.Pushed}},
		w.MenuItem{Icon: "file", Mark: &w.Mark{Colour: "#4A6B52"}, Found: []w.Inline{w.Text("договор-"), w.Strong("ковал"), w.Text("ёва.pdf")}, Trail: w.Dim{Parts: []w.Inline{w.Text("sftpgo · vps · Documents")}, Size: w.Dim11, Variant: w.DimTight, Place: w.Pushed}},
		w.MenuSep{},
		w.MenuItem{Icon: "search", Trail: w.Kbd{Keys: "Enter"}, Label: "Show all results for «ковал»", On: true},
		w.MenuItem{Icon: "all", Found: []w.Inline{w.Text("Timeline · Анна "), w.Strong("Ковал"), w.Text("ёва")}, Trail: w.Dim{Parts: []w.Inline{w.Text("asks every server")}, Size: w.Dim11, Variant: w.DimTight, Place: w.Pushed}},
		w.MenuSep{},
		w.MenuHeading{Rubric: w.Rubric{Text: "Recent searches", Form: w.RubricItem}},
		w.MenuItem{Icon: "clock", Found: []w.Inline{w.Text("смета "), w.Strong("ковал"), w.Text("ёвой")}, Trail: w.Mono{Text: "3 found · yesterday", Variant: w.Mono11DustyTight, Place: w.Pushed}}}, Width: w.MenuWidthFull}, "s-find-3", "tc-menu", 1),
	overlay("fields", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Points per signal", Items: []w.RubricPart{w.Push{Items: []w.PushPart{w.Text("Default")}}, w.RubricCol{Text: "Yours"}}, Form: w.RubricItem}, Form: w.MenuHeadingRoomy},
		w.MenuField{Label: "Same UID", Default: w.Dim{Parts: []w.Inline{w.Text("80")}, Size: w.Dim10_5, Place: w.Pushed}, Input: w.Input{Value: "80", Mono: true, Variant: w.InputChip, Label: "Same UID"}},
		w.MenuField{Label: "Same email", Default: w.Dim{Parts: []w.Inline{w.Text("60")}, Size: w.Dim10_5, Place: w.Pushed}, Input: w.Input{Value: "60", Mono: true, Variant: w.InputChip, Label: "Same email"}},
		w.MenuField{Label: "Same phone", Default: w.Dim{Parts: []w.Inline{w.Text("60")}, Size: w.Dim10_5, Place: w.Pushed}, Input: w.Input{Value: "60", Mono: true, Variant: w.InputChip, Label: "Same phone"}},
		w.MenuField{Label: "Same name", Default: w.Dim{Parts: []w.Inline{w.Text("20")}, Size: w.Dim10_5, Place: w.Pushed}, Input: w.Input{Value: "20", Mono: true, Variant: w.InputChip, Label: "Same name"}},
		w.MenuField{Label: "Same birthday", Badge: &w.Badge{Label: "changed"}, Default: w.Dim{Parts: []w.Inline{w.Text("10")}, Size: w.Dim10_5, Place: w.Pushed}, Input: w.Input{Value: "20", Mono: true, Variant: w.InputChip, Label: "Same birthday"}},
		w.MenuField{Label: "Same title and date (events, tasks, notes)", Default: w.Dim{Parts: []w.Inline{w.Text("70")}, Size: w.Dim10_5, Place: w.Pushed}, Input: w.Input{Value: "70", Mono: true, Variant: w.InputChip, Label: "Same title and date (events, tasks, notes)"}},
		w.MenuFoot{Items: []w.Button{w.Button{Label: "Save", Kind: w.ButtonPrimary, Small: true, Name: "do", Value: "weights", Submit: true}, w.Button{Label: "Back to defaults", Small: true, Name: "do", Value: "defaults", Submit: true}}}}, Width: w.MenuWidth420, Fields: true, Action: "/app/duplicates/decide", Hidden: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}}, "s-dup-3", "tc-menu", 1),
	overlay("marks", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Контакты", On: true, Href: "/app/contacts/nextcloud/kontakty/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "nextcloud · Архив", Href: "/app/contacts/nextcloud/arhiv/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "radicale · vps · Команда", Reason: "Read-only: nothing can be written here."},
		w.MenuSep{},
		w.MenuItem{Label: "New address book", Href: "/app/collections/new?kind=contact&back=/app/contacts/nextcloud/kontakty/import"}}, Width: w.MenuWidth290}, "s-io-10", "tc-menu", 1),
	overlay("marks", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Личный", On: true, Href: "/app/calendar/nextcloud/lichnyj/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа", Href: "/app/calendar/nextcloud/rabota/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "nextcloud · Дни рождения", Href: "/app/calendar/nextcloud/dni-rozhdeniya/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "radicale · vps · Проекты", Href: "/app/calendar/radicale/proekty/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#8C6239"}, Label: "radicale · vps · Дежурства", Href: "/app/calendar/radicale/dezhurstva/import"},
		w.MenuItem{Mark: &w.Mark{Colour: "#9A9280"}, Label: "baikal · Общий", Href: "/app/calendar/baikal/obshchij/import"},
		w.MenuSep{},
		w.MenuItem{Label: "New calendar", Href: "/app/collections/new?kind=event&back=/app/calendar/nextcloud/lichnyj/import"}}, Width: w.MenuWidth290}, "s-io-12", "tc-menu", 1),
	overlay("heading", w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "mix", Items: []w.RubricPart{w.Mono{Text: "carrel.example", Variant: w.MonoTightPlain}}, Form: w.RubricItem, Variant: w.RubricItemLine}, Form: w.MenuHeadingSplit},
		w.MenuItem{Icon: "gear", Label: "Settings"},
		w.MenuItem{Icon: "adm", Label: "Administration"},
		w.MenuSep{},
		w.MenuColumn{Label: "Theme", Control: w.Seg{Options: []w.SegOption{{Label: "Auto", On: true}, {Label: "Light"}, {Label: "Dark"}}, Variant: w.SegStartSmall}},
		w.MenuItem{Icon: "search", Trail: w.Kbd{Keys: "?"}, Label: "Keyboard shortcuts", ScriptOnly: true, Href: "#keys", Key: "?"},
		w.MenuSep{},
		w.MenuItem{Label: "Sign out", Form: "post-signout"}}, Width: w.MenuWidth290}, "account-5", "tc-menu", 1),
	overlay("pop", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "pencil", Label: "Rename and colour"},
		w.MenuItem{Icon: "move", Label: "Move to another server"},
		w.MenuItem{Icon: "dup", Label: "Copy to another server"},
		w.MenuItem{Icon: "export", Detail: &w.Mono{Text: ".ics", Variant: w.Mono11}, Label: "Export"},
		w.MenuItem{Icon: "link", Label: "Publish a link"},
		w.MenuSep{},
		w.MenuItem{Icon: "trash", Label: "Delete on the server…", Tone: w.MenuToneDanger}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}, "s-profile-1", "tc-menu", 1),
	overlay("danger", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "pencil", Label: "Rename"},
		w.MenuItem{Icon: "move", Label: "Move to another server", Href: "/app/collections/move?a=sftpgo&c=documents"},
		w.MenuItem{Icon: "dup", Label: "Copy to another server", Href: "/app/collections/move?a=sftpgo&c=documents&copy=1"},
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
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "radicale · vps · Проекты", Name: "into", Value: "radicale|proekty"},
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Личный", On: true, Name: "into", Value: "nextcloud|lichnyj"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа", Name: "into", Value: "nextcloud|rabota"},
		w.MenuSep{},
		w.MenuItem{Label: "New collection", Name: "into", Value: "new"},
		w.MenuItem{Label: "Skip this one", Tone: w.MenuToneDusty, Name: "into", Value: "skip"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}, "s-backup-8", "tc-menu", 1),
	overlay("danger", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "mail", Label: "Send", Href: "/app/mail/link?id=9c1f4e"},
		w.MenuSep{},
		w.MenuItem{Label: "Revoke", Tone: w.MenuToneDanger, Form: "post-revoke", Name: "id", Value: "9c1f4e"}}, Width: w.MenuWidth212}, "s-publish-15", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "7 days", Name: "lives", Value: "7"},
		w.MenuItem{Label: "14 days", Name: "lives", Value: "14"},
		w.MenuItem{Label: "30 days", On: true, Name: "lives", Value: "30"},
		w.MenuSep{},
		w.MenuItem{Label: "No expiry", Name: "lives", Value: "never"}}, Width: w.MenuWidth212}, "s-publish-14", "tc-menu", 1),
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
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Detail: &w.Mono{Text: "a.belov@example.org", Variant: w.Mono11Dusty}, Trail: w.Tag{Label: "work", Place: w.Pushed}, Adds: &w.Tag{Label: "Артём Белов", Mark: &w.Mark{Colour: "#6B4E7A", Size: w.MarkH11}, Removable: true, Aside: &w.Dim{Parts: []w.Inline{w.Text("· a.belov@example.org")}}, Title: "a.belov@example.org", Name: "to", Value: "a.belov@example.org"}, Label: "Артём Белов", On: true},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Detail: &w.Mono{Text: "belov@personal.example", Variant: w.Mono11Dusty}, Trail: w.Tag{Label: "home", Place: w.Pushed}, Adds: &w.Tag{Label: "Артём Белов", Mark: &w.Mark{Colour: "#6B4E7A", Size: w.MarkH11}, Removable: true, Aside: &w.Dim{Parts: []w.Inline{w.Text("· belov@personal.example")}}, Title: "belov@personal.example", Name: "to", Value: "belov@personal.example"}, Label: "Артём Белов"},
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Detail: &w.Mono{Text: "b.yankovskaya@example.org", Variant: w.Mono11Dusty}, Adds: &w.Tag{Label: "Белла Янковская", Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, Removable: true, Aside: &w.Dim{Parts: []w.Inline{w.Text("· b.yankovskaya@example.org")}}, Title: "b.yankovskaya@example.org", Name: "to", Value: "b.yankovskaya@example.org"}, Label: "Белла Янковская"},
		w.MenuSep{},
		w.MenuItem{Icon: "mail", Adds: &w.Tag{Address: &w.Mono{Text: "бел"}, Removable: true, Title: "бел", Name: "to", Value: "бел"}, Label: "Use «бел» as typed", Tone: w.MenuToneMuted},
		w.MenuSep{},
		w.MenuItem{Note: &w.Span{Face: w.FaceDanger, Text: "baikal · Офис did not answer"}, Label: "Searched 4 of 5 address books ·", Tone: w.MenuToneAside}}, Width: w.MenuWidthFull}, "s-mail-3", "tc-menu", 1),
	overlay("danger", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "Extend"},
		w.MenuSep{},
		w.MenuItem{Label: "Revoke", Tone: w.MenuToneDanger}}, Width: w.MenuWidth212}, "s-admin-11", "tc-menu", 1),
	overlay("danger", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "Make administrator"},
		w.MenuItem{Label: "Kill sessions"},
		w.MenuItem{Label: "Disable"},
		w.MenuSep{},
		w.MenuItem{Icon: "trash", Label: "Delete", Tone: w.MenuToneDanger, Href: "/admin/delete?target=u-anna"}}, Width: w.MenuWidth212}, "s-admin-12", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "STARTTLS", On: true},
		w.MenuItem{Label: "TLS"},
		w.MenuSep{},
		w.MenuItem{Label: "None"}}, Width: w.MenuWidth212}, "s-admin-9", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "All actions", On: true, Name: "action", Value: ""},
		w.MenuSep{},
		w.MenuItem{Label: "login.ok", Name: "action", Value: "login.ok"},
		w.MenuItem{Label: "login.fail", Name: "action", Value: "login.fail"},
		w.MenuItem{Label: "escrow.recover", Name: "action", Value: "escrow.recover"},
		w.MenuItem{Label: "settings.instance", Name: "action", Value: "settings.instance"}}, Width: w.MenuWidth290}, "s-admin-10", "tc-menu", 1),
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
		w.MenuItem{Icon: "search", Trail: w.Kbd{Keys: "?"}, Label: "Keyboard shortcuts", ScriptOnly: true, Href: "#keys", Key: "?"},
		w.MenuSep{},
		w.MenuItem{Label: "Sign out…"}}, Width: w.MenuWidth290}, "desktop-2", "tc-menu", 1),
	overlay("on", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "Only as set in the event"},
		w.MenuSep{},
		w.MenuItem{Label: "At the time · and as set in the event"},
		w.MenuItem{Label: "10 minutes · and as set in the event", On: true},
		w.MenuItem{Label: "30 minutes · and as set in the event"},
		w.MenuItem{Label: "1 hour · and as set in the event"},
		w.MenuSep{},
		w.MenuItem{Label: "Never"}}, Width: w.MenuWidth290}, "desktop-20", "tc-menu", 1),

	// tc-drop
	overlay("row", w.Drop{Trigger: w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}, Row: true, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "mail", Label: "Send"},
		w.MenuSep{},
		w.MenuItem{Label: "Revoke", Tone: w.MenuToneDanger}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}}, "s-publish-1", "tc-drop", 1),
	overlay("closed", w.Drop{Trigger: w.Button{Icon: "dots", Title: "More"}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "move", Label: "Move"},
		w.MenuItem{Icon: "dup", Label: "Copy"},
		w.MenuItem{Icon: "export", Detail: &w.Mono{Text: ".vcf", Variant: w.Mono11}, Label: "Export"},
		w.MenuItem{Icon: "link", Label: "Publish a link"}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}}, "s-contacts-6", "tc-drop", 1),
	overlay("closed", w.Drop{Trigger: w.Button{Label: "Attach", Menu: true, Small: true}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "upload", Label: "From this computer"},
		w.MenuItem{Icon: "file", Label: "From Files"}}, Width: w.MenuWidth212, Pop: w.MenuUnder}}, "s-agenda-1", "tc-drop", 3),
	overlay("pick", w.Drop{Trigger: w.Pick{Value: "Asia/Yekaterinburg", Aside: &w.Dim{Parts: []w.Inline{w.Text("as sent")}, Variant: w.DimIndent}, Label: "Time zone"}, Menu: w.Menu{Items: []w.MenuPart{
		w.Input{Kind: w.InputSearch, Placeholder: "Find a zone", Label: "Find a zone"},
		w.Rubric{Text: "As sent", Form: w.RubricInline},
		w.MenuItem{Label: "Asia/Yekaterinburg", On: true, Name: "tz", Value: "Asia/Yekaterinburg"},
		w.MenuSep{},
		w.MenuItem{Label: "Europe/Moscow", Name: "tz", Value: "Europe/Moscow"},
		w.MenuItem{Label: "UTC", Name: "tz", Value: "UTC"},
		w.MenuSep{},
		w.MenuItem{Label: "Floating — the same clock time anywhere", Name: "tz", Value: "floating"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}, "s-agenda-3", "tc-drop", 2),
	overlay("more", w.Drop{Open: true, Trigger: w.Button{Icon: "dots", Title: "More"}, More: true, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "pencil", Label: "Rename"},
		w.MenuItem{Icon: "export", Detail: &w.Mono{Text: ".vcf", Variant: w.Mono11Dusty}, Label: "Export"}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}}, "components-19", "tc-drop", 1),
	overlay("pick", w.Drop{Open: true, Trigger: w.Pick{Value: "Синк-ап по релизу ковал", Label: "Where"}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "cal", Trail: w.Mono{Text: "10 Aug · Личный", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Звонок с Ковалёвой", On: true},
		w.MenuItem{Icon: "task", Trail: w.Mono{Text: "task · Работа", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Ответить Анне по смете"},
		w.MenuItem{Icon: "con", Trail: w.Mono{Text: "contact", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Анна Ковалёва"},
		w.MenuSep{},
		w.MenuItem{Label: "Enter a UID by hand", Tone: w.MenuToneDusty}}, Width: w.MenuWidthFull, Pop: w.MenuUnder}}, "s-notes-4", "tc-drop", 3),

	// the forms that waited for menus
	overlay("plain", w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}}, "offline-1", "tc-top", 1),
	overlay("suggest", w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}, Focus: true, Variant: w.SearchFull}, Suggest: &w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "Recent searches", Form: w.RubricItem}},
		w.MenuItem{Icon: "clock", Trail: w.Mono{Text: "8 found · 2 min ago", Variant: w.Mono11DustyTight, Place: w.Pushed}, Label: "ковал"},
		w.MenuItem{Icon: "clock", Trail: w.Mono{Text: "3 found · yesterday", Variant: w.Mono11DustyTight, Place: w.Pushed}, Label: "смета"},
		w.MenuItem{Icon: "clock", Trail: w.Mono{Text: "nothing found · Monday", Variant: w.Mono11DustyTight, Place: w.Pushed}, Label: "дача"},
		w.MenuSep{},
		w.MenuItem{Trail: w.Kbd{Keys: "Enter"}, Label: "Type to search everything", Tone: w.MenuToneDusty}}, Width: w.MenuWidthFull}, Actions: []w.TopAction{w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary}, w.Button{Label: "2 min ago", Icon: "refresh"}, w.UserMenu{Name: "mix"}}}, "s-find-2", "tc-top", 1),
	overlay("drop", w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}}, Actions: []w.TopAction{w.TopDrop{Open: true, Trigger: w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary, Focus: true}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "quill", Trail: w.Kbd{Keys: "N"}, Label: "Note", On: true},
		w.MenuSep{},
		w.MenuItem{Icon: "con", Trail: w.Kbd{Keys: "C"}, Label: "Contact"},
		w.MenuItem{Icon: "cal", Trail: w.Kbd{Keys: "E"}, Label: "Event"},
		w.MenuItem{Icon: "task", Trail: w.Kbd{Keys: "T"}, Label: "Task"},
		w.MenuSep{},
		w.MenuItem{Icon: "upload", Label: "Upload file"},
		w.MenuSep{},
		w.MenuItem{Icon: "mail", Label: "Letter"}}}}, w.Button{Label: "2 min ago", Icon: "refresh"}, w.UserMenu{Name: "mix"}}}, "s-notes-7", "tc-top", 1),
	overlay("drop", w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}}, Actions: []w.TopAction{w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary}, w.Button{Label: "2 min ago", Icon: "refresh"}, w.TopDrop{Open: true, Trigger: w.UserMenu{Name: "mix", Focus: true}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuHeading{Rubric: w.Rubric{Text: "mix", Items: []w.RubricPart{w.Mono{Text: "carrel.example", Variant: w.MonoTightPlain}}, Form: w.RubricItem, Variant: w.RubricItemLine}, Form: w.MenuHeadingSplit},
		w.MenuItem{Icon: "gear", Label: "Settings"},
		w.MenuItem{Icon: "adm", Label: "Administration"},
		w.MenuSep{},
		w.MenuColumn{Label: "Theme", Control: w.Seg{Options: []w.SegOption{{Label: "Auto", On: true}, {Label: "Light"}, {Label: "Dark"}}, Variant: w.SegStartSmall}},
		w.MenuItem{Icon: "search", Trail: w.Kbd{Keys: "?"}, Label: "Keyboard shortcuts"},
		w.MenuSep{},
		w.MenuItem{Label: "Sign out"}}, Width: w.MenuWidth290}, End: true}}}, "account-5", "tc-top", 1),
	overlay("plain", w.Bar{Items: []w.BarPart{w.Filter{Query: w.Query{Value: "синк", Label: "Filter by word", Name: "q"}, Action: "/app/calendar/nextcloud/lichnyj", Variant: w.FilterH26W200}, w.Button{Label: "Columns", Icon: "cols"}}, Variant: w.BarWide}, "s-agenda-31", "tc-bar", 1),
	overlay("selection", w.Bar{Items: []w.BarPart{w.Mono{Text: "2.8 MB", Variant: w.MonoMuted}, w.BarSep{}, w.Button{Folds: true, Label: "Download", Icon: "download", Title: "Download", Form: "files-sel", Name: "op", Value: "download"}, w.Button{Folds: true, Label: "Rename", Icon: "pencil", Title: "Rename", Form: "files-sel", Name: "op", Value: "rename"}, w.Button{Folds: true, Label: "Move", Icon: "move", Title: "Move", Form: "files-sel", Name: "op", Value: "move"}, w.Button{Folds: true, Label: "Copy", Icon: "dup", Title: "Copy", Form: "files-sel", Name: "op", Value: "copy"}, w.Button{Folds: true, Label: "Delete", Icon: "trash", Kind: w.ButtonDangerSolid, Title: "Delete", Form: "files-sel", Name: "op", Value: "delete"}}, Count: &w.BarCount{Text: "1 selected"}, Right: &w.BarRight{Items: []w.BarPart{w.Button{Folds: true, Label: "Select all", Icon: "selectall", Small: true, Title: "Select all", ScriptOnly: true}, w.Button{Folds: true, Label: "Clear selection", Icon: "close", Small: true, Title: "Clear selection"}}}, Variant: w.BarSelection}, "dark-3", "tc-bar", 1),
	overlay("selection", w.Bar{Items: []w.BarPart{w.BarSep{}, w.Button{Folds: true, Label: "Download", Icon: "download", Title: "Download", Form: "files-sel", Name: "op", Value: "download"}, w.Button{Folds: true, Label: "Rename", Icon: "pencil", Title: "Public is read-only.", Form: "files-sel", Name: "op", Value: "rename", State: w.ButtonDisabled}, w.Button{Folds: true, Label: "Move", Icon: "move", Title: "Public is read-only.", Form: "files-sel", Name: "op", Value: "move", State: w.ButtonDisabled}, w.Button{Folds: true, Label: "Copy", Icon: "dup", Title: "Copy", Form: "files-sel", Name: "op", Value: "copy"}, w.Button{Folds: true, Label: "Delete", Icon: "trash", Kind: w.ButtonDangerSolid, Title: "Public is read-only.", Form: "files-sel", Name: "op", Value: "delete", State: w.ButtonDisabled}}, Count: &w.BarCount{Text: "With the ticked"}, Variant: w.BarSelection}, "states-34", "tc-bar", 1),
	overlay("more", w.Bar{Items: []w.BarPart{w.Seg{Options: []w.SegOption{{Label: "A–Z", On: true}, {Label: "Recently changed"}}}, w.Seg{Options: []w.SegOption{{Label: "All", On: true}, {Label: "Overdue"}}, Secondary: true}, w.Button{Label: "Columns", Icon: "cols"}, w.Drop{Open: true, Trigger: w.Button{Icon: "dots", Title: "More"}, More: true, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "All", On: true},
		w.MenuItem{Label: "Overdue"}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}}}}, "components-19", "tc-bar", 1),
	overlay("closed", w.Drop{Trigger: w.Button{Icon: "dots", Title: "More"}, More: true, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Tick: &w.TickBox{}, Label: "Events", Tertiary: true},
		w.MenuItem{Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Tasks due", Tertiary: true},
		w.MenuItem{Tick: &w.TickBox{Ticked: w.TickedOff}, Label: "Notes", Tertiary: true},
		w.MenuSep{Tertiary: true},
		w.Range{From: w.Input{Kind: w.InputDate, Value: "2026-08-10", Label: "From", Mono: true}, To: w.Input{Kind: w.InputDate, Value: "2026-08-16", Label: "To", Mono: true}, Show: w.Button{Label: "Show"}, Wrap: true},
		w.MenuSep{},
		w.CheckForm{Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}}, Action: "/app/calendar", Apply: w.Button{Label: "Apply", Small: true}}}, Pop: w.MenuUnderEnd, Inline: true}}, "states-3", "tc-drop", 1),
	overlay("more", w.Bar{Items: []w.BarPart{w.Filter{Query: w.Query{Placeholder: "Filter by word", Label: "Filter by word", Name: "q"}, Waiting: true, Action: "/app/calendar", Variant: w.FilterH26W200}, w.Button{Icon: "back", Title: "Previous week", Href: "/app/calendar?from=2026-08-03"}, w.Button{Label: "Today", Href: "/app/calendar"}, w.Button{Icon: "next", Title: "Next week", Href: "/app/calendar?from=2026-08-17"}, w.SegFold{Seg: w.Seg{Options: []w.SegOption{{Label: "Day", Href: "/app/calendar?view=day"}, {Label: "Week", Href: "/app/calendar?view=week", On: true}, {Label: "Month", Href: "/app/calendar?view=month"}, {Label: "Range", Href: "/app/calendar?view=range"}}}, Phone: true}, w.Range{From: w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true}, To: w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-16", Label: "To", Mono: true}, Show: w.Button{Label: "Show", Submit: true}, Hidden: []w.Hidden{{Name: "view", Value: "range"}}, Secondary: true, Action: "/app/calendar"}, w.Button{Label: "Columns", Icon: "cols"}, w.Drop{Trigger: w.Button{Icon: "dots", Title: "More"}, More: true, Menu: w.Menu{Items: []w.MenuPart{
		w.Range{From: w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true}, To: w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-16", Label: "To", Mono: true}, Show: w.Button{Label: "Show", Submit: true}, Hidden: []w.Hidden{{Name: "view", Value: "range"}}, Action: "/app/calendar", Wrap: true},
		w.MenuSep{},
		w.CheckForm{Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}}, Action: "/app/calendar", Apply: w.Button{Label: "Apply", Small: true, Submit: true}}}, Pop: w.MenuUnderEnd, Inline: true}}}, Right: &w.BarRight{Items: []w.BarPart{w.CheckForm{Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}}, Action: "/app/calendar", Apply: w.Button{Label: "Apply", Small: true, Submit: true}}}, Secondary: true}, Variant: w.BarWide}, "s-agenda-1", "tc-bar", 1),

	// tc-sheet
	overlay("plain", w.Sheet{Icon: "mail", Title: "Letter", Head: []w.SheetHeadPart{w.Tag{Label: "from mix@carrel.example", On: true, Place: w.Pushed}, w.Close{Label: "Close", Href: "/app/contacts"}}, Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("The letter went, but not to everyone. Sending again goes only to the refused.")}, Tone: w.MsgAlert, Space: w.MsgFlushSides},
		w.Line{Items: []w.LinePart{w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, w.Text("Анна Ковалёва "), w.Dim{Parts: []w.Inline{w.Text("· a.kovaleva@example.org")}}, w.Text(" — "), w.Strong("sent")}, Gap: w.LineGap6, Block: true},
		w.Line{Items: []w.LinePart{w.Mark{Colour: "#6B4E7A", Size: w.MarkH11}, w.Text("Артём Белов "), w.Dim{Parts: []w.Inline{w.Text("· a.belov@example.org")}}, w.Text(" — "), w.Faced{Face: w.FaceDanger, Items: []w.LinePart{w.Text("refused: «550 5.1.1 Mailbox unavailable»")}}}, Gap: w.LineGap6, Block: true}}, Foot: []w.SheetFootPart{w.Button{Label: "Send again to the refused", Kind: w.ButtonPrimary, Href: "/app/mail/compose?to=refused"}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Close", Href: "/app/contacts"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth680}, "s-mail-18", "tc-sheet", 1),
	overlay("plain", w.Sheet{Action: "/app/mail/send?op=discard", Icon: "mail", Title: "Close the letter?", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("from mix@carrel.example")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("Carrel keeps no drafts: "), w.Strong("closing loses the text"), w.Text(".")}, Space: w.MsgFlush}}, Foot: []w.SheetFootPart{w.Button{Label: "Close and lose the text", Kind: w.ButtonDangerSolid, Submit: true}, w.Line{Items: []w.LinePart{w.Button{Label: "Keep writing", Kind: w.ButtonPrimary, Href: "/app/mail/compose"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth480}, "s-mail-15", "tc-sheet", 1),
	overlay("plain", w.Sheet{Action: "/app/mail/send", Icon: "eye", Title: "Preview", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("as the recipients get it")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Card{Items: []w.CardPart{
			w.CardHead{Lines: []w.CardHeadLine{{Slot: "From", Parts: []w.CardTextPart{w.Strong("Михаил Енученко"), w.Text(" "), w.Mono{Text: "<mix@carrel.example>", Variant: w.Mono11Muted}}}, {Slot: "Reply-To", Parts: []w.CardTextPart{w.Text("Михаил Енученко "), w.Mono{Text: "<mix@example.org>", Variant: w.Mono11Muted}}}, {Slot: "To", Parts: []w.CardTextPart{w.Text("Анна Ковалёва "), w.Mono{Text: "<a.kovaleva@example.org>", Variant: w.Mono11Muted}, w.Text(", Артём Белов "), w.Mono{Text: "<a.belov@example.org>", Variant: w.Mono11Muted}}}, {Slot: "Subject", Parts: []w.CardTextPart{w.Text("Контакты подрядчика и смета")}}}},
			w.Prose{Parts: []w.Inline{w.Text("Анна, Артём — вот карточка подрядчика и смета, которую он прислал. Договор в папке, ссылка ниже.\n\n-- \nМихаил Енученко")}, Variant: w.Prose13Card},
			w.CardRow{Kind: w.CardAttached, Gap: 6, Items: []w.CardRowPart{w.Tag{Label: "smirnov.vcf", Icon: "con"}, w.Tag{Label: "vstrecha.ics", Icon: "cal"}}},
			w.CardText{Kind: w.CardCaption, Parts: []w.CardTextPart{w.Text("смета-подрядчик.pdf — "), w.Span{Face: w.FaceAccent, Text: "carrel.example/p/9c1f4e"}, w.Break{}, w.Text("Договор подряда, 1 folder, 6 files — "), w.Span{Face: w.FaceAccent, Text: "carrel.example/p/2ab77d"}}}}, Variant: w.CardFlush}}, Foot: []w.SheetFootPart{w.Button{Label: "Send", Icon: "mail", Kind: w.ButtonPrimary, Busy: "Sending…", Submit: true}, w.Kbd{Keys: "Enter"}, w.Button{Label: "Edit", Href: "/app/mail/compose"}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", To: "/app/mail/send?op=close", Submit: true}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth680}, "s-mail-14", "tc-sheet", 1),
	overlay("plain", w.Sheet{Action: "/app/contacts/nextcloud/kontakty/7c2e-51a0-b3d9/copy", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}, Icon: "dup", Title: "Copy contact", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Виктор Гончаров · from nextcloud · Контакты")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Rubric{Text: "Write into"},
		w.Tree{Rows: []w.TreeRow{
			{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Unavailable: true, Send: w.Send{Name: "to", Value: "nextcloud/kontakty", One: true}, Trail: w.Mono{Text: "already there", Variant: w.Mono10Dusty}},
			{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "Архив", Send: w.Send{Name: "to", Value: "nextcloud/arhiv", One: true}, Trail: w.Mono{Text: "nextcloud", Variant: w.Mono10Dusty}},
			{Mark: &w.Mark{Colour: "#8C6239"}, Label: "Рассылки", Send: w.Send{Name: "to", Value: "radicale/rassylki", One: true}, Trail: w.Mono{Text: "radicale · vps", Variant: w.Mono10Dusty}}}, On: 2},
		w.Msg{Parts: []w.MsgPart{w.Text("A copy is written into the chosen address book whole, with a UID of its own: this server holds one record per UID across the account, so for other apps the copy is another card. The card here stays as it is.")}, Space: w.MsgFlushTop2}}, Foot: []w.SheetFootPart{w.Button{Label: "Copy", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/contacts/nextcloud/kontakty/7c2e-51a0-b3d9"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth520}, "s-contacts-37", "tc-sheet", 1),
	overlay("plain", w.Sheet{Icon: "plus", Title: "New calendar", Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("No server that holds calendars is answering now: baikal.")}, Tone: w.MsgAlert, Space: w.MsgFlush}}, Foot: []w.SheetFootPart{w.Button{Label: "Try again", Kind: w.ButtonPrimary, Href: "/app/collections/new?kind=event"}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/calendar"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth480}, "s-profile-21", "tc-sheet", 1),
	overlay("pick", w.Sheet{Action: "/app/mail/attach?what=contact&uid=p7c2", Icon: "con", Title: "Attach a contact", Head: []w.SheetHeadPart{w.Filter{Query: w.Query{Value: "смир", Label: "Search", Name: "q", Of: "attach-search"}, Quiet: true, Variant: w.FilterT12W200, Place: w.Pushed}, w.Close{Label: "Close", Href: "/app/mail/compose"}}, Body: []w.SheetPart{
		w.SheetPick{Rail: []w.SheetPickRail{w.Tree{Rows: []w.TreeRow{
			{Mark: &w.Mark{Kind: w.MarkMulti}, Label: "All address books", Href: "/app/mail/attach?what=contact&q=смир"},
			{Icon: "server", Label: "nextcloud"},
			{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "Личный", Indent: w.TreeLevel22, Href: "/app/mail/attach?what=contact&q=смир&c=lichnyj"},
			{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Indent: w.TreeLevel22, Href: "/app/mail/attach?what=contact&q=смир&c=rabota"},
			{Icon: "server", Label: "radicale · vps"},
			{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Indent: w.TreeLevel22, Href: "/app/mail/attach?what=contact&q=смир&c=komanda"}}, On: 1, Variant: w.TreePadded}}, List: []w.SheetPickList{w.ListGroup{Title: w.Rubric{Text: "Found", Form: w.RubricInline}, Count: &w.Num{Value: 2}}, w.List{Rows: []w.ListRow{
			w.FaceRow{Mark: w.Mark{Colour: "#3F6E8C"}, Avatar: w.Avatar{Initials: "ПС"}, Title: []w.Inline{w.Text("Пётр Смирнов")}, Strong: true, Subtitle: "+7 916 330-41-27", Detail: "p.smirnov@example.org", Caption: "подрядчик", Source: "nextcloud · Работа", On: true, Href: "/app/mail/attach?what=contact&q=смир&uid=p7c2"},
			w.FaceRow{Mark: w.Mark{Colour: "#4A6B52"}, Avatar: w.Avatar{Initials: "ИС"}, Title: []w.Inline{w.Text("Ирина Смирнова")}, Strong: true, Subtitle: "+7 903 221-08-64", Detail: "i.smirnova@example.org", Caption: "—", Source: "nextcloud · Личный", Href: "/app/mail/attach?what=contact&q=смир&uid=s1e8"}}}}, Short: true}}, Flush: true, Foot: []w.SheetFootPart{w.Button{Label: "Attach", Icon: "clip", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Dim{Parts: []w.Inline{w.Text("Пётр Смирнов · "), w.Mono{Text: ".vcf"}, w.Text(" · 3 KB · goes as a file")}, Size: w.Dim11_5, Variant: w.DimWrap}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/mail/compose"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth900}, "s-mail-12", "tc-sheet", 1),
	overlay("close", w.Sheet{Title: "Publish a calendar", Head: []w.SheetHeadPart{w.Badge{Label: "local", Tone: w.BadgeLocal, Title: "Local instance: the link will not open from outside your network"}, w.Dim{Parts: []w.Inline{w.Text("nextcloud · Личный")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.Field{Label: "Link", Parts: []w.FieldPart{w.Line{Items: []w.LinePart{w.Input{Value: "https://carrel.example/p/7Qm2c8XvLk3w", Label: "Link", Mono: true, ReadOnly: true, Variant: w.InputFill}, w.Button{Label: "Copy", Icon: "dup", Small: true, Copy: "https://carrel.example/p/7Qm2c8XvLk3w", Copied: "Copied"}}, Gap: w.LineGap6}, w.FieldHint{Parts: []w.Inline{w.Text("No expiry. Revoke it when it is no longer needed.")}}}, Width: w.FieldW100}}, Foot: []w.SheetFootPart{w.Button{Label: "Open the link", Href: "/p/7Qm2c8XvLk3w"}, w.Button{Label: "All published links", Href: "/app/files/published"}}}, "s-publish-22", "tc-sheet", 1),
	overlay("form-dim", w.Sheet{Action: "/app/publish?c=kontakty&uid=4f2b-9c31-a0e7", Title: "Publish a contact", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Анна Ковалёва")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Field{Label: "What the visitor sees", Parts: []w.FieldPart{w.Group{Items: []w.FieldPart{w.Check{On: true, Name: "name", Value: "1", Parts: []w.CheckPart{w.Text("Name and organisation")}}, w.Check{Name: "phones", Value: "1", Parts: []w.CheckPart{w.Text("Phone numbers")}}, w.Check{Name: "emails", Value: "1", Parts: []w.CheckPart{w.Text("Email addresses")}}, w.Check{Name: "photo", Value: "1", Parts: []w.CheckPart{w.Text("Photo")}}, w.Check{Name: "note", Value: "1", Parts: []w.CheckPart{w.Text("Note")}}}, Form: w.GroupColumnLoose}}, Width: w.FieldW100},
		w.Check{Name: "download", Value: "1", Parts: []w.CheckPart{w.Text("Allow "), w.Mono{Text: ".vcf"}, w.Text(" download and subscription")}},
		w.Field{Label: "Link lives for", Parts: []w.FieldPart{w.Drop{Trigger: w.Pick{Value: "7 days"}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuItem{Label: "7 days", On: true, Name: "lives", Value: "7"},
			w.MenuItem{Label: "14 days", Name: "lives", Value: "14"},
			w.MenuItem{Label: "30 days", Name: "lives", Value: "30"},
			w.MenuSep{},
			w.MenuItem{Label: "No expiry", Name: "lives", Value: "never"}}, Width: w.MenuWidth212, Pop: w.MenuUnder}}}, Width: w.FieldSM}}, Foot: []w.SheetFootPart{w.Button{Label: "Create link", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/contacts"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}}, "s-publish-3", "tc-sheet", 1),
	overlay("plain", w.Sheet{Title: "What the client sees", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("https://carrel.example/dav/")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Tree{Rows: []w.TreeRow{
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
			{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "c/s3c8/%2FDocuments%2F/", Code: true, Indent: w.TreeLevel44, Trail: w.Dim{Parts: []w.Inline{w.Text("Documents")}, Size: w.Dim11}}}, Variant: w.TreeNarrow}}, Foot: []w.SheetFootPart{w.Line{Items: []w.LinePart{w.Button{Label: "Close", Href: "/app/settings/davloom"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth680}, "s-davloom-9", "tc-sheet", 1),
	overlay("plain", w.Sheet{Action: "/app/settings/backup/trigger?job=work", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}, Title: "External trigger", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Рабочее, ежедневно")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Password"},
			w.Line{Items: []w.LinePart{w.Mono{Text: "cbk_…lamp"}, w.Dim{Parts: []w.Inline{w.Text("issued 2 August · last run 16 August 03:10 from "), w.Mono{Text: "10.0.0.5"}, w.Text(" · 14 runs")}, Size: w.Dim11_5}, w.Line{Items: []w.LinePart{w.Button{Label: "New password", To: "/app/settings/backup/trigger?job=work&op=new", Submit: true}, w.Button{Label: "Revoke", Kind: w.ButtonDanger, To: "/app/settings/backup/trigger?job=work&op=revoke", Submit: true}}, Gap: w.LineGap7, Place: w.Pushed}}, Gap: w.LineGap8, Block: true}}, Variant: w.CardFlush},
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "The line to paste", Items: []w.RubricPart{w.Badge{Label: "local", Tone: w.BadgeLocal}}},
			w.Seg{Options: []w.SegOption{{Label: "crontab", Href: "/app/settings/backup/trigger?job=work&tab=crontab", On: true}, {Label: "GoSentry", Href: "/app/settings/backup/trigger?job=work&tab=gosentry"}, {Label: "Gitea", Href: "/app/settings/backup/trigger?job=work&tab=gitea"}, {Label: "curl", Href: "/app/settings/backup/trigger?job=work&tab=curl"}}, Variant: w.SegStart},
			w.Diag{Text: "0 3 * * * curl -fsS -m 30 -X POST -H 'Authorization: Bearer cbk_••••' 'https://carrel.example/api/backup/run?job=work'", Variant: w.DiagInk},
			w.CardText{Kind: w.CardUsage, Parts: []w.CardTextPart{w.Text("The password was shown once and is hidden here. "), w.Strong("New password"), w.Text(" shows the line whole; the old one stops working.")}}}, Variant: w.CardFlush}}, Foot: []w.SheetFootPart{w.Line{Items: []w.LinePart{w.Button{Label: "Close", Href: "/app/settings/backup"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth680}, "s-backup-12", "tc-sheet", 1),
	overlay("plain", w.Sheet{Action: "/app/settings/backup/trigger?job=personal", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}, Title: "External trigger", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Личное на домашний диск")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("Turning it on stores this job’s data key under the server’s key — less protection, for convenience: the job can then read without you signed in.")}, Tone: w.MsgAlert, Space: w.MsgFlush}}, Foot: []w.SheetFootPart{w.Button{Label: "Turn on", Kind: w.ButtonPrimary, To: "/app/settings/backup/trigger?job=personal&op=on", Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/backup"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth680}, "s-backup-14", "tc-sheet", 1),
	overlay("danger", w.Sheet{Action: "/app/settings/connections/disconnect?account=nextcloud", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}, Icon: "server", Tone: w.SheetIconDanger, Title: "Disconnect", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("nextcloud · cloud.example.org")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("Nothing is deleted on cloud.example.org. Carrel removes:")}, Tone: w.MsgAlert, Space: w.MsgFlush},
		w.Msg{Parts: []w.MsgPart{w.Text("• 10 collections"), w.Break{}, w.Text("• 1 published link"), w.Break{}, w.Text("• 1 source of a backup job"), w.Break{}, w.Text("• 1 davloom device’s sync of them")}, Space: w.MsgFlush}}, Foot: []w.SheetFootPart{w.Button{Label: "Disconnect", Kind: w.ButtonDangerSolid, Submit: true}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/connections"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth480}, "s-profile-26", "tc-sheet", 1),
	overlay("plain", w.Sheet{Action: "/app/settings/unlock", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "back", Value: "/app/settings/davloom/devices/new"}}, Title: "Add device", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Телефон · DAVx5")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Field{Label: "Your Carrel password", Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "password", Label: "Your Carrel password", Focus: true, Required: true}}, Width: w.FieldW100},
		w.Dim{Parts: []w.Inline{w.Text("Unlocks stored credentials for five minutes.")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Foot: []w.SheetFootPart{w.Button{Label: "Create password", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/davloom/devices/new"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth400}, "s-davloom-5", "tc-sheet", 1),
	overlay("danger", w.Sheet{Action: "/app/settings/unlock", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "back", Value: "/app/settings/account"}}, Icon: "trash", Tone: w.SheetIconDanger, Title: "Delete account", Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("Deletes the account and the saved connections. Data on the DAV servers stays.")}, Tone: w.MsgAlert, Space: w.MsgFlush},
		w.Field{Label: "Your Carrel password", Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "password", Label: "Your Carrel password", Focus: true, Required: true}}, Width: w.FieldW100}}, Foot: []w.SheetFootPart{w.Button{Label: "Delete account", Kind: w.ButtonDangerSolid, Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/account"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth400}, "account-15", "tc-sheet", 1),
	overlay("danger", w.Sheet{Action: "/admin/", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "action", Value: "delete"}, {Name: "target", Value: "u-anna"}}, Icon: "trash", Tone: w.SheetIconDanger, Title: "Delete anna", Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("Carrel forgets this account: its connections and saved passwords, published links, devices, backups and the history of its password. Its records on the servers stay as they are.")}, Tone: w.MsgAlert, Space: w.MsgFlush},
		w.Field{Label: "Type the login to confirm", Parts: []w.FieldPart{w.Input{Name: "login", Placeholder: "anna", Label: "Type the login to confirm", Focus: true, Required: true}}, Width: w.FieldW100}}, Foot: []w.SheetFootPart{w.Button{Label: "Delete account", Kind: w.ButtonDangerSolid, Submit: true}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/admin/"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth480}, "s-admin-34", "tc-sheet", 1),
	overlay("plain", w.Sheet{Action: "/app/publish?c=rabota&uid=aa01", Title: "Publish an event", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Синк-ап по релизу")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Field{Label: "What the visitor sees", Parts: []w.FieldPart{w.Group{Items: []w.FieldPart{w.Check{On: true, Name: "time", Value: "1", Reason: "Always shown", Parts: []w.CheckPart{w.Text("Time"), w.Dim{Parts: []w.Inline{w.Text("always shown")}, Size: w.Dim11_5}}}, w.Check{On: true, Name: "title", Value: "1", Parts: []w.CheckPart{w.Text("Title")}}, w.Check{Name: "location", Value: "1", Parts: []w.CheckPart{w.Text("Location")}}, w.Check{Name: "description", Value: "1", Parts: []w.CheckPart{w.Text("Description")}}, w.Check{Name: "attendees", Value: "1", Parts: []w.CheckPart{w.Text("Attendees")}}, w.Check{Name: "attachments", Value: "1", Parts: []w.CheckPart{w.Text("Attachments")}}, w.Check{Name: "tags", Value: "1", Parts: []w.CheckPart{w.Text("Tags")}}}, Form: w.GroupColumnLoose}, w.FieldHint{Parts: []w.Inline{w.Text("Without the title the event shows as “Busy”.")}}}, Width: w.FieldW100},
		w.SheetChecks{Items: []w.Check{w.Check{Name: "download", Value: "1", Parts: []w.CheckPart{w.Text("Allow "), w.Mono{Text: ".ics"}, w.Text(" download and subscription")}}}},
		w.Field{Label: "Link lives for", Parts: []w.FieldPart{w.Drop{Trigger: w.Pick{Value: "30 days"}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuItem{Label: "7 days", Name: "lives", Value: "7"},
			w.MenuItem{Label: "14 days", Name: "lives", Value: "14"},
			w.MenuItem{Label: "30 days", On: true, Name: "lives", Value: "30"},
			w.MenuSep{},
			w.MenuItem{Label: "No expiry", Name: "lives", Value: "never"}}, Width: w.MenuWidth212, Pop: w.MenuUnder}}}, Width: w.FieldSM}}, Foot: []w.SheetFootPart{w.Button{Label: "Create link", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/calendar"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}}, "s-publish-19", "tc-sheet", 1),
	overlay("danger", w.Sheet{Action: "/app/tasks/nextcloud/dela/aa10/delete", Icon: "trash", Tone: w.SheetIconDanger, Title: "Delete task", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Отправить документы в МФЦ")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("The task is deleted on nextcloud · Дела. Carrel holds no copy — it reads and writes, it does not store — so there is nothing here to undo it with.")}, Tone: w.MsgAlert, Space: w.MsgFlush},
		w.SheetBlock{Title: w.Rubric{Text: "What is pointing at it now", Variant: w.RubricBottom7}, Items: []w.SheetLine{{Name: "1 published link", Note: w.Dim{Parts: []w.Inline{w.Text("— answers “not found” from that moment")}}}}}}, Foot: []w.SheetFootPart{w.Button{Label: "Delete", Kind: w.ButtonDangerSolid, Submit: true}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth480}, "s-tasks-18", "tc-sheet", 1),
	overlay("danger", w.Sheet{Action: "/app/calendar/nextcloud/lichnyj/8f1c/delete", Icon: "trash", Tone: w.SheetIconDanger, Title: "Delete event", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Синк-ап по релизу")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("Carrel keeps no copy: this cannot be undone.")}, Tone: w.MsgAlert, Space: w.MsgFlush},
		w.SheetBlock{Title: w.Rubric{Text: "What is pointing at it now", Variant: w.RubricBottom7}, Items: []w.SheetLine{{Name: "1 published link", Note: w.Dim{Parts: []w.Inline{w.Text("— answers “not found” from that moment")}}}, {Name: "2 records related to it", Note: w.Dim{Parts: []w.Inline{w.Text("— among the loaded ones; their link is left pointing at nothing")}}}}},
		w.Msg{Parts: []w.MsgPart{w.Text("The server may send its attendees a cancellation.")}, Space: w.MsgFlush},
		w.SheetBlock{Title: w.Rubric{Text: "Attachments", Variant: w.RubricBottom7}, Items: []w.SheetLine{{Name: "release-checklist.md", Note: w.Dim{Parts: []w.Inline{w.Text("— written into the event, goes with it")}}}, {Name: "agenda-10-08.pdf", Note: w.Dim{Parts: []w.Inline{w.Text("— a link to Documents / carrel / attachments, the file stays")}}}}},
		w.Check{Name: "attachment", Value: "documents/carrel/attachments/agenda-10-08.pdf", Parts: []w.CheckPart{w.Text("Delete agenda-10-08.pdf too")}},
		w.Rubric{Text: "The same event in other calendars"},
		w.Line{Items: []w.LinePart{w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, w.Check{On: true, Name: "copy", Value: "nextcloud/lichnyj", Parts: []w.CheckPart{w.Text("nextcloud · Личный")}}, w.Dim{Parts: []w.Inline{w.Text("the copy you opened")}, Size: w.Dim11_5}}, Gap: w.LineGap8, Block: true},
		w.Line{Items: []w.LinePart{w.Mark{Colour: "#6B4E7A", Size: w.MarkH11}, w.Check{Name: "copy", Value: "radicale/proekty", Parts: []w.CheckPart{w.Text("radicale · vps · Проекты")}}, w.Dim{Parts: []w.Inline{w.Text("Holds the same")}, Size: w.Dim11_5}}, Gap: w.LineGap8, Block: true},
		w.Line{Items: []w.LinePart{w.Mark{Colour: "#9A9280", Size: w.MarkH11}, w.Check{Name: "copy", Value: "baikal/obshchij", Parts: []w.CheckPart{w.Text("baikal · Общий")}}, w.Dim{Parts: []w.Inline{w.Text("Where differ")}, Size: w.Dim11_5}}, Gap: w.LineGap8, Block: true}}, Foot: []w.SheetFootPart{w.Button{Label: "Delete", Kind: w.ButtonDangerSolid, Submit: true}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}}, "s-agenda-26", "tc-sheet", 1),
	overlay("plain", w.Sheet{Action: "/app/tasks/done", Icon: "task", Title: "Which copy to mark done?", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Собрать чемодан · 2 copies")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.SheetNote{Text: "A copy that is not ticked stays as it is."},
		w.Line{Items: []w.LinePart{w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, w.Check{On: true, Name: "copy", Value: "nextcloud/dela", Parts: []w.CheckPart{w.Text("nextcloud · Дела")}}, w.Dim{Parts: []w.Inline{w.Text("due 12 Aug 18:00")}, Size: w.Dim11_5}}, Gap: w.LineGap8, Block: true},
		w.Line{Items: []w.LinePart{w.Mark{Colour: "#3F6E8C", Size: w.MarkH11}, w.Check{Name: "copy", Value: "nextcloud/rabota", Parts: []w.CheckPart{w.Text("nextcloud · Работа")}}, w.Dim{Parts: []w.Inline{w.Text("Due differ · due 13 Aug")}, Size: w.Dim11_5}}, Gap: w.LineGap8, Block: true}}, Foot: []w.SheetFootPart{w.Button{Label: "Mark done", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth520}, "s-tasks-16", "tc-sheet", 1),
	overlay("plain", w.Sheet{Action: "/app/contacts/linked/8d21-55ce-0b17/edit", Method: "get", Icon: "pencil", Title: "Which copy to edit?", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Вера Лебедева · 3 copies")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.SheetNote{Text: "Saving changes one copy only; the others stay as they are."},
		w.Line{Items: []w.LinePart{w.Mark{Colour: "#A05A2C", Size: w.MarkH11}, w.Check{On: true, Name: "copy", Value: "nextcloud/arhiv", Radio: true, Parts: []w.CheckPart{w.Text("nextcloud · Архив")}}, w.Dim{Parts: []w.Inline{w.Text("the copy you opened · vera@arhiv.example")}, Size: w.Dim11_5}}, Gap: w.LineGap8, Block: true},
		w.Line{Items: []w.LinePart{w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, w.Check{Name: "copy", Value: "nextcloud/kontakty", Radio: true, Parts: []w.CheckPart{w.Text("nextcloud · Контакты")}}, w.Dim{Parts: []w.Inline{w.Text("Email differ · v.lebedeva@example.org")}, Size: w.Dim11_5}}, Gap: w.LineGap8, Block: true},
		w.Line{Items: []w.LinePart{w.Mark{Colour: "#9A9280", Size: w.MarkH11}, w.Dim{Parts: []w.Inline{w.Text("baikal · Общий")}}, w.Dim{Parts: []w.Inline{w.Text("Phone differ · read-only: it cannot be edited")}, Size: w.Dim11_5}}, Gap: w.LineGap8, Block: true}}, Foot: []w.SheetFootPart{w.Button{Label: "Edit this copy", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth520}, "s-contacts-32", "tc-sheet", 1),
	overlay("plain", w.Sheet{Action: "/app/calendar/nextcloud/rabota/7c1e-22ab-90f4/edit", Icon: "refresh", Title: "Keep the meetings that were moved by hand?", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Дежурство: первая линия · saved in part")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.SheetBlock{Title: w.Rubric{Text: "Written", Variant: w.RubricBottom7}, Items: []w.SheetLine{{Name: "The series ends on 19 January", Note: w.Dim{Parts: []w.Inline{w.Text("— saved")}}}, {Name: "2 separate events", Note: w.Dim{Parts: []w.Inline{w.Text("— 26 January, 12:00 and 16 February, 14:30")}}}}},
		w.SheetBlock{Title: w.Rubric{Text: "Not written · 1", Variant: w.RubricBottom7}, Items: []w.SheetLine{{Name: "2 February, 09:00 — Дежурство: вторая линия", Note: w.Dim{Parts: []w.Inline{w.Text("— the server answered 507 Insufficient Storage")}}}}},
		w.Msg{Parts: []w.MsgPart{w.Text("The meeting of 2 February was not made a separate event: it stays in the series as it was before you saved.")}, Tone: w.MsgAlert, Space: w.MsgFlush},
		w.SheetNote{Text: "Nothing is lost: what could not be written apart is still in the series."}}, Foot: []w.SheetFootPart{w.Button{Label: "Try again", Kind: w.ButtonPrimary, To: "/app/calendar/nextcloud/rabota/7c1e-22ab-90f4/edit?op=apart&day=2027-02-02", Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Leave it in the series", Href: "/app/calendar/nextcloud/rabota/7c1e-22ab-90f4"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth520}, "s-agenda-24", "tc-sheet", 1),
	overlay("plain", w.Sheet{Action: "/app/calendar/nextcloud/rabota/7c1e-22ab-90f4/edit", Icon: "refresh", Title: "Keep the meeting that was moved by hand?", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Дежурство: первая линия · not saved yet")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.SheetBlock{Title: w.Rubric{Text: "What you changed", Variant: w.RubricBottom7}, Items: []w.SheetLine{{Name: "The series repeats every two weeks", Note: w.Dim{Parts: []w.Inline{w.Text("— it repeated weekly")}}}}},
		w.SheetBlock{Title: w.Rubric{Text: "Left without its day · 1", Variant: w.RubricBottom7}, Items: []w.SheetLine{{Name: "26 January, 12:00 — Duty moved", Note: w.Dim{Parts: []w.Inline{w.Text("— it was the meeting of 26 January, 09:00")}}}}},
		w.SheetNote{Text: "Kept, it becomes an event of its own, with what was changed in it. Removed, it goes: the series no longer has that day."}}, Foot: []w.SheetFootPart{w.Button{Label: "Keep as an event", Kind: w.ButtonPrimary, Name: "left", Value: "keep", Submit: true}, w.Kbd{Keys: "Enter"}, w.Button{Label: "Remove", Name: "left", Value: "remove", Submit: true}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Name: "left", Value: "cancel", Submit: true}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth520}, "s-agenda-23", "tc-sheet", 1),
	overlay("plain", w.Sheet{Action: "/app/calendar/nextcloud/rabota/7c1e-22ab-90f4/edit", Icon: "refresh", Title: "Keep the meetings that were moved by hand?", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Дежурство: первая линия · not saved yet")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.SheetBlock{Title: w.Rubric{Text: "What you changed", Variant: w.RubricBottom7}, Items: []w.SheetLine{{Name: "The series ends on 19 January", Note: w.Dim{Parts: []w.Inline{w.Text("— it ran weekly until 28 December")}}}}},
		w.SheetBlock{Title: w.Rubric{Text: "Left without their day · 3", Variant: w.RubricBottom7}, Items: []w.SheetLine{{Name: "26 January, 12:00 — Duty moved", Note: w.Dim{Parts: []w.Inline{w.Text("— it was the meeting of 26 January, 09:00")}}}, {Name: "2 February, 09:00 — Дежурство: вторая линия", Note: w.Dim{Parts: []w.Inline{w.Text("— its title was changed")}}}, {Name: "16 February, 14:30 — Duty moved", Note: w.Dim{Parts: []w.Inline{w.Text("— it was the meeting of 16 February, 09:00")}}}}},
		w.SheetNote{Text: "Kept, each becomes an event of its own, with what was changed in it. Removed, they go with the rest of the series after 19 January."}}, Foot: []w.SheetFootPart{w.Button{Label: "Keep as events", Kind: w.ButtonPrimary, Name: "left", Value: "keep", Submit: true}, w.Kbd{Keys: "Enter"}, w.Button{Label: "Remove", Name: "left", Value: "remove", Submit: true}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Name: "left", Value: "cancel", Submit: true}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth520}, "s-agenda-22", "tc-sheet", 1),
	overlay("plain", w.Sheet{Action: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/rewrite", Icon: "refresh", Title: "Rewrite the addresses in their records?", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Анна Ковалёва · saved")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Rubric{Text: "The addresses"},
		w.Line{Items: []w.LinePart{w.Mono{Text: "a.kovaleva@example.org"}, w.Dim{Parts: []w.Inline{w.Text("22 records name it")}, Size: w.Dim11_5}}, Gap: w.LineGap8, Block: true},
		w.Line{Items: []w.LinePart{w.Mono{Text: "→"}, w.Pick{Name: "to.1", Options: []w.PickOption{{Value: "a.kovaleva@slovo.example", Label: "a.kovaleva@slovo.example", On: true}, {Value: "anna@home.example", Label: "anna@home.example"}, {Value: "", Label: "Do not rewrite"}}, Label: "Rewrite a.kovaleva@example.org to", Variant: w.PickT12W240}}, Gap: w.LineGap8, Block: true},
		w.Line{Items: []w.LinePart{w.Mono{Text: "anna@oldfirm.example"}, w.Dim{Parts: []w.Inline{w.Text("3 records name it; they keep it")}, Size: w.Dim11_5}}, Gap: w.LineGap8, Block: true},
		w.Line{Items: []w.LinePart{w.Mono{Text: "→"}, w.Pick{Name: "to.2", Options: []w.PickOption{{Value: "a.kovaleva@slovo.example", Label: "a.kovaleva@slovo.example"}, {Value: "anna@home.example", Label: "anna@home.example"}, {Value: "", Label: "Do not rewrite", On: true}}, Label: "Rewrite anna@oldfirm.example to", Variant: w.PickT12W240}}, Gap: w.LineGap8, Block: true},
		w.Rubric{Text: "What rewriting can set off"},
		w.Line{Items: []w.LinePart{w.Check{On: true, Name: "group", Value: "scheduling", Parts: []w.CheckPart{w.Text("On servers that schedule")}}, w.Drop{Trigger: w.Button{Label: "9 records", Menu: true, Small: true}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "10 Aug", Variant: w.Mono11}, Label: "Синк-ап по релизу", Href: "/app/calendar/nextcloud/rabota/aa01"},
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "14 Aug", Variant: w.Mono11}, Label: "Ретро спринта", Href: "/app/calendar/nextcloud/rabota/aa02"},
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "2 Sep", Variant: w.Mono11}, Label: "Планирование Q4", Href: "/app/calendar/nextcloud/rabota/aa03"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}}, Gap: w.LineGap8, Block: true},
		w.Dim{Parts: []w.Inline{w.Text("The server may send a cancellation to the old address, an invitation to the new one and an update to the other attendees.")}, Size: w.Dim11_5, Variant: w.DimWrap},
		w.Line{Items: []w.LinePart{w.Check{On: true, Name: "group", Value: "unknown", Parts: []w.CheckPart{w.Text("On baikal")}}, w.Drop{Trigger: w.Button{Label: "2 records", Menu: true, Small: true}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuItem{Mark: &w.Mark{Colour: "#9A9280"}, Detail: &w.Mono{Text: "12 Aug", Variant: w.Mono11}, Label: "Общий созвон", Href: "/app/calendar/baikal/obshchij/bb01"},
			w.MenuItem{Mark: &w.Mark{Colour: "#9A9280"}, Detail: &w.Mono{Text: "21 Aug", Variant: w.Mono11}, Label: "Демо для заказчика", Href: "/app/calendar/baikal/obshchij/bb02"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}}, Gap: w.LineGap8, Block: true},
		w.Dim{Parts: []w.Inline{w.Text("The server did not say whether it schedules.")}, Size: w.Dim11_5, Variant: w.DimWrap},
		w.Line{Items: []w.LinePart{w.Check{Name: "group", Value: "others", Parts: []w.CheckPart{w.Text("Meetings organised by others")}}, w.Drop{Trigger: w.Button{Label: "3 records", Menu: true, Small: true}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "14 Aug", Variant: w.Mono11}, Label: "Ретро спринта", Href: "/app/calendar/nextcloud/rabota/aa02"},
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "18 Aug", Variant: w.Mono11}, Label: "Встреча с подрядчиком", Href: "/app/calendar/nextcloud/rabota/cc01"},
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "5 Sep", Variant: w.Mono11}, Label: "Юбилей отдела", Href: "/app/calendar/nextcloud/rabota/cc02"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}}, Gap: w.LineGap8, Block: true},
		w.Dim{Parts: []w.Inline{w.Text("Not ticked: these stay as they are. The organiser’s server may refuse the change, and its next update may undo it.")}, Size: w.Dim11_5, Variant: w.DimWrap},
		w.Line{Items: []w.LinePart{w.Check{On: true, Name: "group", Value: "past", Parts: []w.CheckPart{w.Text("Past meetings")}}, w.Drop{Trigger: w.Button{Label: "5 records", Menu: true, Small: true}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "3 Aug", Variant: w.Mono11}, Label: "Синк-ап по релизу", Href: "/app/calendar/nextcloud/rabota/dd01"},
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "27 Jul", Variant: w.Mono11}, Label: "Синк-ап по релизу", Href: "/app/calendar/nextcloud/rabota/dd02"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}}, Gap: w.LineGap8, Block: true},
		w.Dim{Parts: []w.Inline{w.Text("They will show the new address, as if the person had been invited at it.")}, Size: w.Dim11_5, Variant: w.DimWrap},
		w.Msg{Parts: []w.MsgPart{w.Text("A record that stands in two groups is rewritten only when both are ticked: «Ретро спринта» is on a server that schedules and is organised by another, so it stays.")}, Space: w.MsgFlush},
		w.SheetNote{Text: "2 records in Shared are read-only and keep the old address. Collections that are not ticked were not looked in."},
		w.Tally{Records: []string{"to.1 group=scheduling", "to.1 group=scheduling", "to.1 group=scheduling", "to.1 group=scheduling", "to.1 group=scheduling", "to.1 group=scheduling", "to.1 group=scheduling", "to.1 group=scheduling", "to.1 group=scheduling group=others", "to.1 group=others", "to.1 group=others", "to.1 group=past", "to.1 group=past", "to.1 group=past", "to.1 group=past", "to.1 group=past", "to.1 group=unknown", "to.1 group=unknown", "to.1", "to.1", "to.2", "to.2", "to.2"}}}, Foot: []w.SheetFootPart{w.Button{Label: "Rewrite 17 records", Kind: w.ButtonPrimary, Count: "Rewrite {n} records", CountOne: "Rewrite 1 record", Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Leave as is"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth520}, "s-contacts-29", "tc-sheet", 1),
	overlay("plain-w520", w.Sheet{Action: "/app/collections/move?a=nextcloud&c=lichnyj", Icon: "move", Title: "Move calendar", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("nextcloud · Личный")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Rubric{Text: "Write onto"},
		w.Tree{Rows: []w.TreeRow{
			{Icon: "server", Label: "nextcloud", Unavailable: true, Send: w.Send{Name: "to", Value: "nextcloud", One: true}, Trail: w.Mono{Text: "already there", Variant: w.Mono10Dusty}},
			{Icon: "server", Label: "radicale · vps", Send: w.Send{Name: "to", Value: "radicale", One: true}, Trail: w.Mono{Text: "calendars", Variant: w.Mono10Dusty}},
			{Icon: "server", Label: "baikal", Unavailable: true, Send: w.Send{Name: "to", Value: "baikal", One: true}, Trail: w.Mono{Text: "password · 401", Variant: w.Mono10Dusty}},
			{Icon: "server", Label: "sftpgo · vps", Unavailable: true, Send: w.Send{Name: "to", Value: "sftpgo", One: true}, Trail: w.Mono{Text: "files only", Variant: w.Mono10Dusty}}}, On: 2},
		w.Msg{Parts: []w.MsgPart{w.Text("Carrel creates a calendar there, copies all "), w.Strong("128"), w.Text(" events, then deletes this one on nextcloud. Devices, published links and backup jobs pointing at the old address will see it vanish — same as Delete, except the records live on. This is not renaming the path. nextcloud may tell the attendees of its meetings that they are cancelled.")}, Space: w.MsgFlushTop2}}, Foot: []w.SheetFootPart{w.Button{Label: "Move", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/connections"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth520}, "s-profile-8", "tc-sheet", 1),
	overlay("form-w480", w.Sheet{Action: "/app/collections/rename?a=sftpgo&c=documents", Icon: "pencil", Title: "Rename", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("sftpgo · vps · Documents")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Field{Label: "Name", Parts: []w.FieldPart{w.Input{Name: "name", Value: "Documents", Label: "Name", Focus: true}}, Width: w.FieldW100},
		w.Msg{Parts: []w.MsgPart{w.Text("Links, devices and backups lose the old address.")}, Space: w.MsgFlush}}, Foot: []w.SheetFootPart{w.Button{Label: "Rename", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/connections"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth480}, "s-profile-18", "tc-sheet", 1),
	overlay("plain", w.Sheet{Action: "/app/collections/nextcloud/dni-rozhdeniya/rename", Icon: "pencil", Title: "Rename and colour", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("nextcloud")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Field{Label: "Name", Parts: []w.FieldPart{w.Input{Name: "name", Value: "Дни рождения", Label: "Name", Focus: true}}, Width: w.FieldW100},
		w.Field{Label: "Colour", Parts: []w.FieldPart{w.Swatches{Colours: []string{"#4A6B52", "#3F6E8C", "#6B4E7A", "#A05A2C", "#8C6239", "#9A9280"}, On: 3, Name: "colour"}}, Width: w.FieldW100},
		w.Field{Label: "Address", LabelAside: w.Dim{Parts: []w.Inline{w.Text("— stays as it is")}}, Parts: []w.FieldPart{w.Input{Name: "address.stays.as.it.is", Value: "…/remote.php/dav/calendars/mix/birthdays/", Label: "Address", Mono: true, Variant: w.InputDusty}}, Width: w.FieldW100},
		w.Msg{Parts: []w.MsgPart{w.Text("Renaming is not moving: the address does not change, so devices, published links and backup jobs pointing here go on working. The new name and colour are written to the server — other clients will show them at their next sync.")}, Space: w.MsgFlushTop2}}, Foot: []w.SheetFootPart{w.Button{Label: "Save", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth520}, "s-profile-7", "tc-sheet", 1),
	overlay("quick-form", w.Sheet{Action: "/app/calendar/nextcloud/lichnyj/quick", Icon: "cal", Title: "New event", Head: []w.SheetHeadPart{w.Drop{Trigger: w.Pick{Value: "nextcloud · Личный", Mark: &w.Mark{Colour: "#4A6B52"}, Variant: w.PickT12W200}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Trail: w.Mono{Text: "last used", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "nextcloud · Личный", On: true, Name: "into", Value: "nextcloud/lichnyj"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа", Name: "into", Value: "nextcloud/rabota"},
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "nextcloud · Дни рождения", Name: "into", Value: "nextcloud/dni-rozhdeniya"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "radicale · vps · Проекты", Name: "into", Value: "radicale/proekty"},
		w.MenuItem{Mark: &w.Mark{Colour: "#8C6239"}, Label: "radicale · vps · Дежурства", Name: "into", Value: "radicale/dezhurstva"},
		w.MenuItem{Mark: &w.Mark{Colour: "#9A9280"}, Label: "baikal · Общий", Name: "into", Value: "baikal/obshchij"}}, Width: w.MenuWidth290, Pop: w.MenuUnderEnd}, Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.SheetFields{Items: []w.Field{w.Field{Label: "Title", Parts: []w.FieldPart{w.Input{Name: "title", Value: "Забрать документы", Label: "Title", Focus: true}}, Place: w.Grown}, w.Field{Label: "Date", Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "date", Value: "2026-08-11", Label: "Date", Mono: true}}, Width: w.FieldFix130}, w.Field{Label: "Start", Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Name: "start", Value: "18:00", Label: "Start", Mono: true}}, Width: w.FieldFix96}, w.Field{Label: "End", Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Name: "end", Value: "18:40", Label: "End", Mono: true, Follows: "start"}}, Width: w.FieldFix96}, w.Field{Label: " ", Parts: []w.FieldPart{w.Check{Name: "allday", Value: "1", Disables: "start end", Parts: []w.CheckPart{w.Text("All day")}}}, Width: w.FieldFill120}, w.Field{Label: "Remind", Parts: []w.FieldPart{w.Pick{Name: "remind", Options: []w.PickOption{{Value: "at.the.time", Label: "At the time"}, {Value: "5.minutes.before", Label: "5 minutes before"}, {Value: "10.minutes.before", Label: "10 minutes before", On: true}, {Value: "30.minutes.before", Label: "30 minutes before"}, {Value: "1.hour.before", Label: "1 hour before"}, {Value: "1.day.before", Label: "1 day before"}, {Value: "custom", Label: "Custom…", Shows: "remind.count remind.unit"}, {Value: "none", Label: "No reminder"}}, Label: "Remind"}}, Width: w.FieldFill200}, w.Field{Hidden: true, Label: "How long", Parts: []w.FieldPart{w.Input{Name: "remind.count", Value: "1", Label: "How long", Mono: true, Numeric: true}}, Width: w.FieldXS}, w.Field{Hidden: true, Label: "Unit", Parts: []w.FieldPart{w.Pick{Name: "remind.unit", Options: []w.PickOption{{Value: "minutes", Label: "minutes before"}, {Value: "hours", Label: "hours before"}, {Value: "days", Label: "days before", On: true}, {Value: "weeks", Label: "weeks before"}}, Label: "Unit"}}, Width: w.FieldSM}}},
		w.Dim{Parts: []w.Inline{w.Text("Repeats, guests and more live on the card. A reminder is set here because that is why most events are created.")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Foot: []w.SheetFootPart{w.Button{Label: "Save event", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Button{Label: "Open the card", Icon: "expand", To: "/app/calendar/nextcloud/lichnyj/quick?op=open", Submit: true}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}}, "s-notes-9", "tc-sheet", 1),
	overlay("danger", w.Sheet{Icon: "trash", Tone: w.SheetIconDanger, Title: "Delete 3 items", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Documents / carrel / attachments")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("Carrel keeps no copy: this cannot be undone.")}, Tone: w.MsgAlert, Space: w.MsgFlush},
		w.Msg{Parts: []w.MsgPart{w.Text("Notes or events may hold them as attachments: Carrel does not look through the records.")}, Space: w.MsgFlush}}, Foot: []w.SheetFootPart{w.Button{Label: "Delete", Kind: w.ButtonDangerSolid}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/files/sftpgo/documents"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth480}, "s-files-22", "tc-sheet", 1),
	overlay("server", w.Sheet{ID: "copy-sheet", Action: "/app/files/sftpgo/documents/op/carrel/attachments?op=copy", Icon: "dup", Title: "Copy 3 items", Head: []w.SheetHeadPart{w.Drop{Trigger: w.Pick{Value: "sftpgo · vps", Label: "Server", Variant: w.PickT12W240}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "sftpgo · vps", On: true, Href: "/app/files/sftpgo/documents/op/carrel/attachments?op=copy&server=sftpgo", Swap: "copy-sheet"},
		w.MenuItem{Label: "nextcloud", Href: "/app/files/sftpgo/documents/op/carrel/attachments?op=copy&server=nextcloud", Swap: "copy-sheet"},
		w.MenuItem{Label: "webdav · nas", Reason: "webdav · nas did not answer.", Href: "#"}}, Width: w.MenuWidth290, Pop: w.MenuUnderEnd}, Place: w.Pushed}, w.Dim{Parts: []w.Inline{w.Text("from Documents / carrel / attachments")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Body: []w.SheetPart{
		w.Rubric{Text: "Where to put it", Items: []w.RubricPart{w.Button{Label: "New folder", Icon: "plus", Small: true}}},
		w.Tree{Rows: []w.TreeRow{
			{Icon: "file", Label: "Documents", Send: w.Send{Name: "to", Value: "sftpgo/documents", One: true}, Href: "/app/files/sftpgo/documents/op/carrel/attachments?op=copy&to=sftpgo/documents", Swap: "copy-sheet"},
			{Icon: "file", Label: "carrel", Indent: w.TreeLevel22, Send: w.Send{Name: "to", Value: "sftpgo/documents/carrel", One: true}, Href: "/app/files/sftpgo/documents/op/carrel/attachments?op=copy&to=sftpgo/documents/carrel", Swap: "copy-sheet"},
			{Icon: "file", Label: "attachments", Indent: w.TreeLevel40, Send: w.Send{Name: "to", Value: "sftpgo/documents/carrel/attachments", One: true}, Href: "/app/files/sftpgo/documents/op/carrel/attachments?op=copy&to=sftpgo/documents/carrel/attachments", Swap: "copy-sheet"},
			{Icon: "file", Label: "2026-07", Indent: w.TreeLevel58, Send: w.Send{Name: "to", Value: "sftpgo/documents/carrel/attachments/2026-07", One: true}, Href: "/app/files/sftpgo/documents/op/carrel/attachments?op=copy&to=sftpgo/documents/carrel/attachments/2026-07", Swap: "copy-sheet"},
			{Icon: "file", Label: "Public", Unavailable: true, Send: w.Send{Name: "to", Value: "sftpgo/public", One: true}, Trail: w.Badge{Label: "read-only"}}}, On: 4},
		w.Msg{Parts: []w.MsgPart{w.Text("A copy travels through Carrel even inside one server: each file is downloaded and uploaded again, one at a time.")}, Space: w.MsgFlushTop2}}, Foot: []w.SheetFootPart{w.Button{Label: "Copy", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Button{Label: "Open the folder", Name: "open", Value: "1", Submit: true, NoScript: true}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth520}, "s-files-11", "tc-sheet", 1),
	overlay("form", w.Sheet{ID: "move-sheet", Action: "/app/files/sftpgo/documents/op/carrel/attachments?op=move", Icon: "move", Title: "Move 3 items", Head: []w.SheetHeadPart{w.Drop{Trigger: w.Pick{Value: "sftpgo · vps", Label: "Server", Variant: w.PickT12W240}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "sftpgo · vps", On: true, Href: "/app/files/sftpgo/documents/op/carrel/attachments?op=move&server=sftpgo", Swap: "move-sheet"},
		w.MenuItem{Label: "nextcloud", Href: "/app/files/sftpgo/documents/op/carrel/attachments?op=move&server=nextcloud", Swap: "move-sheet"},
		w.MenuItem{Label: "webdav · nas", Reason: "webdav · nas did not answer.", Href: "#"}}, Width: w.MenuWidth290, Pop: w.MenuUnderEnd}, Place: w.Pushed}, w.Dim{Parts: []w.Inline{w.Text("from Documents / carrel / attachments")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Body: []w.SheetPart{
		w.Rubric{Text: "Where to put it", Items: []w.RubricPart{w.Button{Label: "New folder", Icon: "plus", Small: true}}},
		w.Tree{Rows: []w.TreeRow{
			{Icon: "file", Label: "Documents", Send: w.Send{Name: "to", Value: "sftpgo/documents", One: true}, Href: "/app/files/sftpgo/documents/op/carrel/attachments?op=move&to=sftpgo/documents", Swap: "move-sheet"},
			{Icon: "file", Label: "carrel", Indent: w.TreeLevel22, Send: w.Send{Name: "to", Value: "sftpgo/documents/carrel", One: true}, Href: "/app/files/sftpgo/documents/op/carrel/attachments?op=move&to=sftpgo/documents/carrel", Swap: "move-sheet"},
			{Icon: "file", Label: "attachments", Indent: w.TreeLevel40, Send: w.Send{Name: "to", Value: "sftpgo/documents/carrel/attachments", One: true}, Href: "/app/files/sftpgo/documents/op/carrel/attachments?op=move&to=sftpgo/documents/carrel/attachments", Swap: "move-sheet"},
			{Icon: "file", Label: "2026-07", Indent: w.TreeLevel58, Send: w.Send{Name: "to", Value: "sftpgo/documents/carrel/attachments/2026-07", One: true}, Href: "/app/files/sftpgo/documents/op/carrel/attachments?op=move&to=sftpgo/documents/carrel/attachments/2026-07", Swap: "move-sheet"},
			{Icon: "file", Label: "Public", Unavailable: true, Send: w.Send{Name: "to", Value: "sftpgo/public", One: true}, Trail: w.Badge{Label: "read-only"}}}, On: 4}}, Foot: []w.SheetFootPart{w.Button{Label: "Move", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Button{Label: "Open the folder", Name: "open", Value: "1", Submit: true, NoScript: true}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth520}, "s-files-4", "tc-sheet", 1),
	overlay("plain-head", w.Sheet{Title: "Sources", Head: []w.SheetHeadPart{w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.Dim{Parts: []w.Inline{w.Text("The rail’s sources, on a narrow screen.")}, Size: w.Dim11_5, Variant: w.DimWrap}}}, "components-41", "tc-sheet", 1),
	overlay("plain-accent", w.Sheet{Action: "/app/files/sftpgo/mkcol", Icon: "plus", Title: "New folder", Body: []w.SheetPart{
		w.Field{Label: "On which server", Parts: []w.FieldPart{w.Pick{Name: "on.which.server", Options: []w.PickOption{{Value: "sftpgo.vps", Label: "sftpgo · vps", On: true}}, Label: "On which server"}}, Width: w.FieldW100},
		w.Field{Label: "Name", Parts: []w.FieldPart{w.Input{Name: "name", Value: "Scans", Label: "Name", Focus: true}, w.FieldHint{Parts: []w.Inline{w.Text("A folder's name is its address: there is nothing else to ask.")}, Variant: w.HintSmall}}, Width: w.FieldW100}}, Foot: []w.SheetFootPart{w.Button{Label: "Create", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth480}, "s-profile-20", "tc-sheet", 1),
	overlay("danger-block", w.Sheet{Action: "/app/files/sftpgo/documents/delete", Icon: "trash", Tone: w.SheetIconDanger, Title: "Delete on the server", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("sftpgo · vps · Documents")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("The folder and all "), w.Strong("412"), w.Text(" files in it are deleted on sftpgo · vps. Carrel holds no copy — it reads and writes, it does not store — so there is nothing here to undo it with.")}, Tone: w.MsgAlert, Space: w.MsgFlush},
		w.SheetBlock{Title: w.Rubric{Text: "What is pointing at it now", Variant: w.RubricBottom7}, Items: []w.SheetLine{{Name: "Attachments", Note: w.Dim{Parts: []w.Inline{w.Text("— the attachment folder is in it; new attachments will have nowhere to go")}}}, {Name: "Backup job «Личное на домашний диск»", Note: w.Dim{Parts: []w.Inline{w.Text("— copies this folder; it will run one short and say so")}}}}},
		w.Msg{Parts: []w.MsgPart{w.Text("Last time it was in an archive: "), w.Mono{Text: "2026-08-14"}, w.Text(". Restoring brings the files back into a "), w.Strong("new"), w.Text(" folder — not into this one, and not to this address.")}, Space: w.MsgFlush},
		w.Line{Items: []w.LinePart{w.Button{Label: "Open in Files", Icon: "file"}, w.Dim{Parts: []w.Inline{w.Text("Download what you need from it there first.")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Gap: w.LineGap8, Block: true},
		w.Field{Label: "Type the name to confirm", Parts: []w.FieldPart{w.Input{Name: "type.the.name.to.confirm", Placeholder: "Documents", Label: "Type the name to confirm", Focus: true}}, Width: w.FieldW100}}, Foot: []w.SheetFootPart{w.Button{Label: "Delete on the server", Kind: w.ButtonDangerSolid, Submit: true}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth480}, "s-profile-19", "tc-sheet", 1),
	overlay("refusal", w.Sheet{ID: "upload-sheet", Action: "/app/files/sftpgo/documents/upload/carrel/attachments", Icon: "upload", Title: "Upload files", Head: []w.SheetHeadPart{w.Drop{Trigger: w.Pick{Value: "sftpgo · vps", Label: "Server", Variant: w.PickT12W240}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "sftpgo · vps", On: true, Href: "/app/files/sftpgo/documents/upload/carrel/attachments?server=sftpgo", Swap: "upload-sheet"},
		w.MenuItem{Label: "nextcloud", Href: "/app/files/sftpgo/documents/upload/carrel/attachments?server=nextcloud", Swap: "upload-sheet"},
		w.MenuItem{Label: "webdav · nas", Reason: "webdav · nas did not answer.", Href: "#"}}, Width: w.MenuWidth290, Pop: w.MenuUnderEnd}, Place: w.Pushed}, w.Dim{Parts: []w.Inline{w.Text("Documents / carrel / attachments")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Body: []w.SheetPart{
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Where to put it"},
			w.Line{Items: []w.LinePart{w.Input{Name: "name", Value: "2026-07", Label: "Name of the new folder", Focus: true, Variant: w.InputFill}, w.Button{Label: "Create", Kind: w.ButtonPrimary, Small: true, To: "/app/files/sftpgo/documents/mkcol/carrel/attachments", Submit: true}, w.Button{Label: "Cancel", Small: true}}, Gap: w.LineGap8, Block: true},
			w.Msg{Parts: []w.MsgPart{w.Text("There is a folder called «2026-07» in attachments already. Choose another name.")}, Tone: w.MsgAlert, Space: w.MsgFlushY8},
			w.Tree{Rows: []w.TreeRow{
				{Icon: "file", Label: "Documents", Href: "/app/files/sftpgo/documents/upload/carrel/attachments?to=sftpgo/documents", Swap: "upload-sheet"},
				{Icon: "file", Label: "carrel", Indent: w.TreeLevel22, Href: "/app/files/sftpgo/documents/upload/carrel/attachments?to=sftpgo/documents/carrel", Swap: "upload-sheet"},
				{Icon: "file", Label: "attachments", Indent: w.TreeLevel40, Href: "/app/files/sftpgo/documents/upload/carrel/attachments?to=sftpgo/documents/carrel/attachments", Swap: "upload-sheet", Trail: w.Badge{Label: "attachment folder"}},
				{Icon: "file", Label: "2026-07", Indent: w.TreeLevel58, Href: "/app/files/sftpgo/documents/upload/carrel/attachments?to=sftpgo/documents/carrel/attachments/2026-07", Swap: "upload-sheet"},
				{Icon: "file", Label: "Public", Unavailable: true, Trail: w.Badge{Label: "read-only"}}}, On: 3}}, Variant: w.CardFlush}}, Foot: []w.SheetFootPart{w.Button{Label: "Upload", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}}, "s-files-19", "tc-sheet", 1),
	overlay("issued", w.Sheet{Icon: "shield", Tone: w.SheetIconDanger, Title: "Reset a password", Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("The password of "), w.Strong("anna"), w.Text(" is reset.")}, Space: w.MsgFlush},
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Temporary password", Items: []w.RubricPart{w.Mono{Text: "visible only now", Variant: w.MonoTightDanger}}},
			w.Diag{Text: "login     anna\npassword  4c1-tide-oven", Variant: w.DiagInk},
			w.CardRow{Kind: w.CardTop8, Gap: 8, Items: []w.CardRowPart{w.Button{Label: "Copy", Icon: "dup", Kind: w.ButtonPrimary}, w.Dim{Parts: []w.Inline{w.Text("Hand it over yourself; they choose their own at the next sign-in.")}, Size: w.Dim11_5, Variant: w.DimWrap}}}}, Variant: w.CardChosenFlushGrow}}, Foot: []w.SheetFootPart{w.Button{Label: "Done", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Enter"}}, Width: w.SheetWidth480}, "s-admin-14", "tc-sheet", 1),
	overlay("tree", w.Sheet{Icon: "file", Title: "Attachment folder", Body: []w.SheetPart{
		w.SheetNote{Text: "Where large attachments go: a folder of a file collection you can write to. Carrel makes it on first use."},
		w.Tree{Rows: []w.TreeRow{
			{Icon: "file", Label: "Documents", Trail: w.Dim{Parts: []w.Inline{w.Text("sftpgo")}, Size: w.Dim11}},
			{Icon: "file", Label: "carrel", Indent: w.TreeLevel22},
			{Icon: "file", Label: "attachments", Indent: w.TreeLevel40},
			{Icon: "file", Label: "Files", Trail: w.Dim{Parts: []w.Inline{w.Text("nextcloud")}, Size: w.Dim11}},
			{Icon: "file", Label: "Public", Unavailable: true, Trail: w.Badge{Label: "read-only"}}}, On: 3}}, Foot: []w.SheetFootPart{w.Button{Label: "Choose this folder", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Enter"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth520}, "attachments-4", "tc-sheet", 1),
	overlay("danger", w.Sheet{Action: "/app/settings/escrow/withdraw", Icon: "shield", Tone: w.SheetIconDanger, Title: "Withdraw my deposited key", Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("After withdrawal nobody can recover this account, including you. The copy of your data key on this server is destroyed; your records and connections stay as they are.")}, Tone: w.MsgAlert, Space: w.MsgFlush}}, Foot: []w.SheetFootPart{w.Button{Label: "Withdraw my deposited key", Kind: w.ButtonDangerSolid, Submit: true}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth480}, "account-11", "tc-sheet", 1),
	overlay("icon", w.Sheet{ID: "picker-sheet", Icon: "file", Title: "Attach from Files", Head: []w.SheetHeadPart{w.Drop{Trigger: w.Pick{Value: "sftpgo · vps", Label: "Server", Variant: w.PickT12W240}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Label: "sftpgo · vps", On: true, Href: "/app/files/picker?for=attach&server=sftpgo", Swap: "picker-sheet"},
		w.MenuItem{Label: "nextcloud", Href: "/app/files/picker?for=attach&server=nextcloud", Swap: "picker-sheet"},
		w.MenuItem{Label: "webdav · nas", Reason: "webdav · nas did not answer.", Href: "#"}}, Width: w.MenuWidth290, Pop: w.MenuUnderEnd}, Place: w.Pushed}, w.Dim{Parts: []w.Inline{w.Text("to «Собрать чемодан»")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Body: []w.SheetPart{
		w.SheetNote{Text: "Attach a file that is in Files already. Up to 2 MB it is written into the record, as a file from this computer is; a larger one stays where it lies, and the record keeps a link to it."},
		w.Tree{Rows: []w.TreeRow{
			{Icon: "file", Label: "Documents"},
			{Icon: "file", Label: "carrel", Indent: w.TreeLevel22},
			{Icon: "file", Label: "attachments", Indent: w.TreeLevel40},
			{Icon: "file", Label: "2026-07", Indent: w.TreeLevel58},
			{Icon: "img", Label: "2026-08-04-whiteboard.jpg", Indent: w.TreeLevel58, Trail: w.Dim{Parts: []w.Inline{w.Text("2.8 MB · link")}, Size: w.Dim11}},
			{Icon: "img", Label: "2026-08-10-sync-board.png", Indent: w.TreeLevel58, Trail: w.Dim{Parts: []w.Inline{w.Text("412 KB · inside")}, Size: w.Dim11}},
			{Icon: "doc", Label: "release-checklist.md", Indent: w.TreeLevel58, Trail: w.Dim{Parts: []w.Inline{w.Text("3.1 KB · inside")}, Size: w.Dim11}},
			{Icon: "doc", Label: "смета-подрядчик.pdf", Indent: w.TreeLevel58, Trail: w.Dim{Parts: []w.Inline{w.Text("1.2 MB · inside")}, Size: w.Dim11}},
			{Icon: "file", Label: "Public"}}}}, Foot: []w.SheetFootPart{w.Button{Label: "Attach", Icon: "clip", Kind: w.ButtonPrimary, State: w.ButtonDisabled}, w.Kbd{Keys: "Enter"}, w.Dim{Parts: []w.Inline{w.Text("Choose a file in the tree.")}, Size: w.Dim11_5}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth520}, "s-files-16", "tc-sheet", 1),
	overlay("icon", w.Sheet{Action: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/rewrite", Icon: "refresh", Title: "Rewrite the address in their records?", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Анна Ковалёва · saved")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Rubric{Text: "The address"},
		w.Line{Items: []w.LinePart{w.Mono{Text: "a.kovaleva@example.org"}, w.Dim{Parts: []w.Inline{w.Text("2 records name it")}, Size: w.Dim11_5}}, Gap: w.LineGap8, Block: true},
		w.SheetNote{Text: "2 records in Shared are read-only and keep the old address. Общий · baikal could not be read, and what it holds keeps the old address. Collections that are not ticked were not looked in."}}, Foot: []w.SheetFootPart{w.Button{Label: "Rewrite 0 records", Kind: w.ButtonPrimary, Title: "Nothing here can be rewritten.", Submit: true, State: w.ButtonDisabled}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Leave as is", To: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/rewrite?op=leave", Submit: true}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth520}, "s-contacts-20", "tc-sheet", 1),
	overlay("icon", w.Sheet{Action: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/rewrite", Icon: "refresh", Title: "Rewrite the address in their records?", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Анна Ковалёва · saved")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Rubric{Text: "The address"},
		w.Line{Items: []w.LinePart{w.Mono{Text: "a.kovaleva@example.org"}, w.Dim{Parts: []w.Inline{w.Text("1 record names it")}, Size: w.Dim11_5}}, Gap: w.LineGap8, Block: true},
		w.Line{Items: []w.LinePart{w.Mono{Text: "→"}, w.Pick{Name: "to.1", Options: []w.PickOption{{Value: "a.kovaleva@slovo.example", Label: "a.kovaleva@slovo.example", On: true}, {Value: "", Label: "Do not rewrite"}}, Label: "Rewrite a.kovaleva@example.org to", Variant: w.PickT12W240}}, Gap: w.LineGap8, Block: true},
		w.Rubric{Text: "What rewriting can set off"},
		w.Line{Items: []w.LinePart{w.Check{On: true, Name: "group", Value: "scheduling", Parts: []w.CheckPart{w.Text("On servers that schedule")}}, w.Drop{Trigger: w.Button{Label: "1 record", Menu: true, Small: true}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "10 Aug", Variant: w.Mono11}, Label: "Синк-ап по релизу", Href: "/app/calendar/nextcloud/rabota/aa01"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}}, Gap: w.LineGap8, Block: true},
		w.Dim{Parts: []w.Inline{w.Text("The server may send a cancellation to the old address, an invitation to the new one and an update to the other attendees.")}, Size: w.Dim11_5, Variant: w.DimWrap},
		w.SheetNote{Text: "Общий · baikal could not be read, and what it holds keeps the old address. Collections that are not ticked were not looked in."},
		w.Tally{Records: []string{"to.1 group=scheduling"}}}, Foot: []w.SheetFootPart{w.Button{Label: "Rewrite 1 record", Kind: w.ButtonPrimary, Count: "Rewrite {n} records", CountOne: "Rewrite 1 record", Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Leave as is", To: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/rewrite?op=leave", Submit: true}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth520}, "s-contacts-19", "tc-sheet", 1),
	overlay("icon", w.Sheet{Action: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/move", Icon: "move", Title: "Move contact", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Анна Ковалёва · from nextcloud · Контакты")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Rubric{Text: "Write into"},
		w.Tree{Rows: []w.TreeRow{
			{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Unavailable: true, Send: w.Send{Name: "to", Value: "nextcloud/kontakty", One: true}, Trail: w.Mono{Text: "already there", Variant: w.Mono10Dusty}},
			{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "Архив", Send: w.Send{Name: "to", Value: "nextcloud/arhiv", One: true}, Trail: w.Mono{Text: "nextcloud", Variant: w.Mono10Dusty}},
			{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Unavailable: true, Send: w.Send{Name: "to", Value: "radicale/komanda", One: true}, Trail: w.Badge{Label: "read-only"}},
			{Mark: &w.Mark{Colour: "#8C6239"}, Label: "Рассылки", Send: w.Send{Name: "to", Value: "radicale/rassylki", One: true}, Trail: w.Mono{Text: "radicale · vps", Variant: w.Mono10Dusty}}}, On: 2}}, Foot: []w.SheetFootPart{w.Button{Label: "Move", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/contacts"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth520}, "s-contacts-2", "tc-sheet", 1),
	overlay("plain", w.Sheet{Title: "Publish a calendar", Head: []w.SheetHeadPart{w.Badge{Label: "local", Tone: w.BadgeLocal, Title: "Local instance: the link will not open from outside your network"}, w.Dim{Parts: []w.Inline{w.Text("nextcloud · Личный")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}, w.Close{Label: "Close", Href: "/app/calendar"}}, Body: []w.SheetPart{
		w.Field{Label: "Link", Parts: []w.FieldPart{w.Line{Items: []w.LinePart{w.Input{Value: "https://carrel.example/p/7Qm2c8XvLk3w", Label: "Link", Mono: true, ReadOnly: true, Variant: w.InputFill}, w.Button{Label: "Copy", Icon: "dup", Small: true, Copy: "https://carrel.example/p/7Qm2c8XvLk3w", Copied: "Copied"}}, Gap: w.LineGap6}, w.FieldHint{Parts: []w.Inline{w.Text("Until 2026-09-09.")}}}, Width: w.FieldW100}}, Foot: []w.SheetFootPart{w.Button{Label: "Send by letter", Kind: w.ButtonPrimary, Href: "/app/mail/link?id=7Qm2"}, w.Kbd{Keys: "Enter"}, w.Button{Label: "Open the link", Href: "/p/7Qm2c8XvLk3w"}, w.Button{Label: "All published links", Href: "/app/files/published"}}}, "s-publish-13", "tc-sheet", 1),
	overlay("close", w.Sheet{Action: "/app/notes/nextcloud/zapisi/quick", Icon: "quill", Title: "New note", Head: []w.SheetHeadPart{w.Drop{Trigger: w.Pick{Value: "nextcloud · Записи", Mark: &w.Mark{Colour: "#4A6B52"}, Variant: w.PickT12W200}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Trail: w.Mono{Text: "last used", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "nextcloud · Записи", On: true, Name: "into", Value: "nextcloud/zapisi"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "nextcloud · Дневник", Name: "into", Value: "nextcloud/dnevnik"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "nextcloud · Общие с семьёй", Reason: "Read-only: nothing can be written here.", Name: "into", Value: "nextcloud/obshchie-s-semey"}}, Width: w.MenuWidth290, Pop: w.MenuUnderEnd}, Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.Area{Text: "Позвонить в МФЦ во вторник до 12:00 — уточнить, нужен ли оригинал доверенности.", Variant: w.AreaBodySm, Focus: true, Name: "description", Label: "Text"}}, Foot: []w.SheetFootPart{w.Button{Label: "Save note", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Ctrl+Enter"}, w.Button{Label: "Open the card", Icon: "expand", To: "/app/notes/nextcloud/zapisi/quick?op=open", Submit: true}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}}, "s-notes-5", "tc-sheet", 1),
	overlay("close", w.Sheet{Action: "/app/contacts/nextcloud/kontakty/quick", Icon: "con", Title: "New contact", Head: []w.SheetHeadPart{w.Drop{Trigger: w.Pick{Value: "nextcloud · Контакты", Mark: &w.Mark{Colour: "#4A6B52"}, Variant: w.PickT12W200}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Trail: w.Mono{Text: "last used", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "nextcloud · Контакты", On: true, Name: "into", Value: "nextcloud/kontakty"},
		w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "nextcloud · Архив", Name: "into", Value: "nextcloud/arhiv"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "radicale · vps · Команда", Reason: "Read-only: nothing can be written here.", Name: "into", Value: "radicale/komanda"}}, Width: w.MenuWidth290, Pop: w.MenuUnderEnd}, Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.SheetFields{Items: []w.Field{w.Field{Label: "Display name", Parts: []w.FieldPart{w.Input{Name: "fn", Value: "Пётр Ковалёв", Label: "Display name", Focus: true}}, Place: w.Grown}, w.Field{Label: "Phone", Parts: []w.FieldPart{w.Input{Name: "tel", Placeholder: "+7 …", Label: "Phone", Mono: true}}, Width: w.FieldFill200}, w.Field{Label: "Email", Parts: []w.FieldPart{w.Input{Name: "email", Placeholder: "name@example.org", Label: "Email"}}, Width: w.FieldFill200}}},
		w.Dim{Parts: []w.Inline{w.Text("The first word is Given, the rest Family — «Пётр Ковалёв», not «Ковалёв Пётр». Everything else is filled in on the card.")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Foot: []w.SheetFootPart{w.Button{Label: "Save contact", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Button{Label: "Open the card", Icon: "expand", To: "/app/contacts/nextcloud/kontakty/quick?op=open", Submit: true}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}}, "s-notes-8", "tc-sheet", 1),
	overlay("close", w.Sheet{Icon: "quill", Title: "New note", Head: []w.SheetHeadPart{w.Close{Label: "Close", Place: w.Pushed}}, Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("There is nowhere to write: connect an account that holds notes first.")}, Space: w.MsgFlush}}, Foot: []w.SheetFootPart{w.Button{Label: "Connect account", Icon: "server", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth480}, "s-notes-6", "tc-sheet", 1),
	overlay("close", w.Sheet{Action: "/app/tasks/nextcloud/dela/quick", Icon: "task", Title: "New task", Head: []w.SheetHeadPart{w.Drop{Trigger: w.Pick{Value: "nextcloud · Дела", Mark: &w.Mark{Colour: "#4A6B52"}, Variant: w.PickT12W200}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Trail: w.Mono{Text: "last used", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "nextcloud · Дела", On: true, Name: "into", Value: "nextcloud/dela"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа", Name: "into", Value: "nextcloud/rabota"}}, Width: w.MenuWidth290, Pop: w.MenuUnderEnd}, Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.SheetFields{Items: []w.Field{w.Field{Label: "Summary", Parts: []w.FieldPart{w.Input{Name: "summary", Value: "Отправить документы в МФЦ", Label: "Summary", Focus: true}}, Place: w.Grown}, w.Field{Label: "Due date", Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "due.date", Placeholder: "YYYY-MM-DD", Label: "Due date", Mono: true}}, Width: w.FieldFix130}, w.Field{Label: "Due time", Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Name: "due.time", Placeholder: "--:--", Label: "Due time", Mono: true}}, Width: w.FieldFix96}, w.Field{Label: "Remind", Parts: []w.FieldPart{w.Pick{Name: "remind", Options: []w.PickOption{{Value: "at.the.time", Label: "At the time"}, {Value: "5.minutes.before", Label: "5 minutes before"}, {Value: "10.minutes.before", Label: "10 minutes before"}, {Value: "30.minutes.before", Label: "30 minutes before"}, {Value: "1.hour.before", Label: "1 hour before"}, {Value: "1.day.before", Label: "1 day before"}, {Value: "custom", Label: "Custom…", Shows: "remind.count remind.unit"}, {Value: "none", Label: "No reminder", On: true}}, Label: "Remind"}}, Width: w.FieldFill200}, w.Field{Hidden: true, Label: "How long", Parts: []w.FieldPart{w.Input{Name: "remind.count", Value: "1", Label: "How long", Mono: true, Numeric: true}}, Width: w.FieldXS}, w.Field{Hidden: true, Label: "Unit", Parts: []w.FieldPart{w.Pick{Name: "remind.unit", Options: []w.PickOption{{Value: "minutes", Label: "minutes before"}, {Value: "hours", Label: "hours before"}, {Value: "days", Label: "days before", On: true}, {Value: "weeks", Label: "weeks before"}}, Label: "Unit"}}, Width: w.FieldSM}}},
		w.Dim{Parts: []w.Inline{w.Text("No time means the whole day. A task with no due date is a task all the same — both may stay empty.")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Foot: []w.SheetFootPart{w.Button{Label: "Save task", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Button{Label: "Open the card", Icon: "expand", To: "/app/tasks/nextcloud/dela/quick?op=open", Submit: true}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}}, "s-notes-10", "tc-sheet", 1),
	overlay("icon", w.Sheet{Action: "/app/settings/escrow/deposit", Icon: "shield", Tone: w.SheetIconInk, Title: "Deposit my key", Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("With the escrow master password, an administrator will be able to decrypt the DAV credentials you save here, and through them your records. You can withdraw the copy later from Account.")}, Tone: w.MsgAlert, Space: w.MsgFlush},
		w.Field{Label: "Your Carrel password", Parts: []w.FieldPart{w.Input{Name: "password", Value: "••••••••••", Label: "Your Carrel password", Focus: true}}, Width: w.FieldW100}}, Foot: []w.SheetFootPart{w.Button{Label: "Deposit my key", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth480}, "account-2", "tc-sheet", 1),
	overlay("close", w.Sheet{Icon: "search", Title: "Keyboard shortcuts", Head: []w.SheetHeadPart{w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.SheetKeys{Items: []w.SheetKey{{Keys: "N", Does: "Quick note over this screen"}, {Keys: "/", Does: "Focus the search field in the header"}, {Keys: "Enter", Does: "Open the selected record full screen"}, {Keys: "Esc", Does: "Close a sheet, the details panel or focus mode"}, {Keys: "Ctrl+S", Does: "Save in the editor"}, {Keys: "Ctrl+Enter", Does: "Save or send a sheet with text: a quick note, a letter"}, {Keys: "Ctrl+V", Does: "Paste an image as an attachment"}}},
		w.SheetHint{Text: "Shortcuts do nothing while the cursor is in a field. None of them destroy anything."}}, Foot: []w.SheetFootPart{w.Button{Label: "Close"}, w.Kbd{Keys: "Esc"}, w.Kbd{Keys: "?", Place: w.Pushed}}, Width: w.SheetWidth480}, "account-6", "tc-sheet", 1),
	overlay("icon", w.Sheet{Action: "/app/settings/unlock", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "back", Value: "/app/settings/connections"}, {Name: "for", Value: "nextcloud"}}, Title: "Show password", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("nextcloud")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Field{Label: "Your Carrel password", Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "password", Label: "Your Carrel password", Focus: true, Required: true}}, Width: w.FieldW100},
		w.Dim{Parts: []w.Inline{w.Text("Unlocks stored credentials for five minutes.")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Foot: []w.SheetFootPart{w.Button{Label: "Show", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/connections"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth400}, "s-profile-3", "tc-sheet", 1),
	overlay("icon", w.Sheet{Action: "/app/collections/new", Icon: "plus", Title: "New calendar", Body: []w.SheetPart{
		w.Field{Label: "On which server", Parts: []w.FieldPart{w.Pick{Name: "account", Options: []w.PickOption{{Value: "nextcloud", Label: "nextcloud", On: true}}, Label: "On which server"}, w.FieldHint{Parts: []w.Inline{w.Mono{Text: "baikal"}, w.Text(" is not answering, so it is not listed.")}, Variant: w.HintSmall}}, Width: w.FieldW100},
		w.Field{Label: "Name", Parts: []w.FieldPart{w.Input{Name: "name", Value: "Поездки", Label: "Name", Focus: true}}, Width: w.FieldW100},
		w.Field{Label: "Colour", Parts: []w.FieldPart{w.Swatches{Colours: []string{"#4A6B52", "#3F6E8C", "#6B4E7A", "#A05A2C", "#8C6239", "#9A9280"}, On: 2, Name: "colour"}, w.FieldHint{Parts: []w.Inline{w.Text("Written to the server as "), w.Mono{Text: "calendar-color"}, w.Text(" — DAVx5 and Thunderbird will show the same one.")}, Variant: w.HintSmall}}, Width: w.FieldW100},
		w.Field{Label: "What goes in it", Parts: []w.FieldPart{w.Group{Items: []w.FieldPart{w.Check{On: true, Name: "components", Value: "VEVENT", Parts: []w.CheckPart{w.Text("Events")}}, w.Check{On: true, Name: "components", Value: "VTODO", Parts: []w.CheckPart{w.Text("Tasks")}}, w.Check{Name: "components", Value: "VJOURNAL", Parts: []w.CheckPart{w.Text("Notes")}}}}, w.FieldHint{Parts: []w.Inline{w.Text("Most servers will not change this later. A collection holding only tasks appears in "), w.Strong("Tasks"), w.Text(", not here.")}, Variant: w.HintSmall}}, Width: w.FieldW100},
		w.Field{Label: "Address", Parts: []w.FieldPart{w.Input{Name: "address", Placeholder: "poezdki", Label: "Address", Mono: true, From: "name"}, w.FieldHint{Parts: []w.Inline{w.Mono{Text: "…/remote.php/dav/calendars/mix/"}, w.SlugOf{Of: "address", Text: "poezdki"}, w.Mono{Text: "/"}, w.Text(" — taken from the name, yours to change. Other clients and restores go by the address, not the name.")}, Variant: w.HintSmall}}, Width: w.FieldW100}}, Foot: []w.SheetFootPart{w.Button{Label: "Create", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/connections"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}}, "s-profile-5", "tc-sheet", 1),
	overlay("danger", w.Sheet{Action: "/app/collections/new", Fields: []w.Hidden{{Name: "account", Value: "radicale"}, {Name: "name", Value: "Поездки"}, {Name: "colour", Value: "#6B4E7A"}, {Name: "components", Value: "VEVENT"}, {Name: "components", Value: "VTODO"}, {Name: "address", Value: "poezdki"}}, Icon: "plus", Tone: w.SheetIconDanger, Title: "New calendar", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("radicale · vps")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("The server would not make it.")}, Tone: w.MsgAlert, Space: w.MsgFlush},
		w.Diag{Text: "[create]  MKCALENDAR /mix/poezdki/ → 403\n              <D:error><D:need-privileges/></D:error>\n              no DAV:bind on the calendar home-set"},
		w.SheetNote{Text: "This account may sign in and read, but radicale · vps would not let it create here. The server does not say why — no such right for the user, or a password issued read-only; whoever runs the server can tell."}}, Foot: []w.SheetFootPart{w.Button{Label: "Back to the form", Kind: w.ButtonPrimary, To: "/app/collections/new?op=back", Submit: true}, w.Kbd{Keys: "Enter"}, w.Button{Label: "Try again", Submit: true}}}, "s-profile-6", "tc-sheet", 1),
	overlay("danger", w.Sheet{Action: "/app/collections/delete?a=nextcloud&c=lichnyj", Icon: "trash", Tone: w.SheetIconDanger, Title: "Delete on the server", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("nextcloud · Личный")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Msg{Parts: []w.MsgPart{w.Text("The calendar and all "), w.Strong("128"), w.Text(" events in it are deleted on nextcloud. Carrel holds no copy — it reads and writes, it does not store — so there is nothing here to undo it with. nextcloud may tell the attendees of its meetings that they are cancelled.")}, Tone: w.MsgAlert, Space: w.MsgFlush},
		w.SheetBlock{Title: w.Rubric{Text: "What is pointing at it now", Variant: w.RubricBottom7}, Items: []w.SheetLine{{Name: "1 published link", Note: w.Dim{Parts: []w.Inline{w.Text("— answers “not found” from that moment")}}}, {Name: "Backup job «Личное на домашний диск»", Note: w.Dim{Parts: []w.Inline{w.Text("— copies this collection; it will run one short and say so")}}}, {Name: "3 devices through davloom", Note: w.Dim{Parts: []w.Inline{w.Text("— it disappears for them at the next sync")}}}}},
		w.Msg{Parts: []w.MsgPart{w.Text("Last time it was in an archive: "), w.Mono{Text: "2026-08-14"}, w.Text(". Restoring brings the records back into a "), w.Strong("new"), w.Text(" collection — not into this one, and not to this address.")}, Space: w.MsgFlush},
		w.Line{Items: []w.LinePart{w.Button{Label: "Download", Icon: "download", Href: "/app/calendar/nextcloud/lichnyj/export", Aside: &w.Mono{Text: ".ics", Variant: w.Mono11}, After: "first"}, w.Dim{Parts: []w.Inline{w.Text("128 events, one file, straight to you")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Gap: w.LineGap8, Block: true},
		w.Field{Label: "Type the name to confirm", Parts: []w.FieldPart{w.Input{Name: "name", Placeholder: "Личный", Label: "Type the name to confirm", Focus: true}}, Width: w.FieldW100}}, Foot: []w.SheetFootPart{w.Button{Label: "Delete on the server", Kind: w.ButtonDangerSolid, Submit: true}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/connections"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}}, "s-profile-9", "tc-sheet", 1),
	overlay("icon", w.Sheet{Action: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/rewrite", Icon: "refresh", Title: "Rewrite the addresses in their records?", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Анна Ковалёва · saved")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Rubric{Text: "The addresses"},
		w.Line{Items: []w.LinePart{w.Mono{Text: "a.kovaleva@example.org"}, w.Dim{Parts: []w.Inline{w.Text("22 records name it")}, Size: w.Dim11_5}}, Gap: w.LineGap8, Block: true},
		w.Line{Items: []w.LinePart{w.Mono{Text: "→"}, w.Pick{Name: "to.1", Options: []w.PickOption{{Value: "a.kovaleva@slovo.example", Label: "a.kovaleva@slovo.example", On: true}, {Value: "anna@home.example", Label: "anna@home.example"}, {Value: "", Label: "Do not rewrite"}}, Label: "Rewrite a.kovaleva@example.org to", Variant: w.PickT12W240}}, Gap: w.LineGap8, Block: true},
		w.Line{Items: []w.LinePart{w.Mono{Text: "anna@oldfirm.example"}, w.Dim{Parts: []w.Inline{w.Text("3 records name it; they keep it")}, Size: w.Dim11_5}}, Gap: w.LineGap8, Block: true},
		w.Line{Items: []w.LinePart{w.Mono{Text: "→"}, w.Pick{Name: "to.2", Options: []w.PickOption{{Value: "a.kovaleva@slovo.example", Label: "a.kovaleva@slovo.example"}, {Value: "anna@home.example", Label: "anna@home.example"}, {Value: "", Label: "Do not rewrite", On: true}}, Label: "Rewrite anna@oldfirm.example to", Variant: w.PickT12W240}}, Gap: w.LineGap8, Block: true},
		w.Rubric{Text: "What rewriting can set off"},
		w.Line{Items: []w.LinePart{w.Check{On: true, Name: "group", Value: "scheduling", Parts: []w.CheckPart{w.Text("On servers that schedule")}}, w.Drop{Trigger: w.Button{Label: "9 records", Menu: true, Small: true}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "10 Aug", Variant: w.Mono11}, Label: "Синк-ап по релизу", Href: "/app/calendar/nextcloud/rabota/aa01"},
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "14 Aug", Variant: w.Mono11}, Label: "Ретро спринта", Href: "/app/calendar/nextcloud/rabota/aa02"},
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "2 Sep", Variant: w.Mono11}, Label: "Планирование Q4", Href: "/app/calendar/nextcloud/rabota/aa03"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}}, Gap: w.LineGap8, Block: true},
		w.Dim{Parts: []w.Inline{w.Text("The server may send a cancellation to the old address, an invitation to the new one and an update to the other attendees.")}, Size: w.Dim11_5, Variant: w.DimWrap},
		w.Line{Items: []w.LinePart{w.Check{On: true, Name: "group", Value: "unknown", Parts: []w.CheckPart{w.Text("On baikal")}}, w.Drop{Trigger: w.Button{Label: "2 records", Menu: true, Small: true}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuItem{Mark: &w.Mark{Colour: "#9A9280"}, Detail: &w.Mono{Text: "12 Aug", Variant: w.Mono11}, Label: "Общий созвон", Href: "/app/calendar/baikal/obshchij/bb01"},
			w.MenuItem{Mark: &w.Mark{Colour: "#9A9280"}, Detail: &w.Mono{Text: "21 Aug", Variant: w.Mono11}, Label: "Демо для заказчика", Href: "/app/calendar/baikal/obshchij/bb02"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}}, Gap: w.LineGap8, Block: true},
		w.Dim{Parts: []w.Inline{w.Text("The server did not say whether it schedules.")}, Size: w.Dim11_5, Variant: w.DimWrap},
		w.Line{Items: []w.LinePart{w.Check{On: true, Name: "group", Value: "others", Parts: []w.CheckPart{w.Text("Meetings organised by others")}}, w.Drop{Trigger: w.Button{Label: "3 records", Menu: true, Small: true}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "14 Aug", Variant: w.Mono11}, Label: "Ретро спринта", Href: "/app/calendar/nextcloud/rabota/aa02"},
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "18 Aug", Variant: w.Mono11}, Label: "Встреча с подрядчиком", Href: "/app/calendar/nextcloud/rabota/cc01"},
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "5 Sep", Variant: w.Mono11}, Label: "Юбилей отдела", Href: "/app/calendar/nextcloud/rabota/cc02"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}}, Gap: w.LineGap8, Block: true},
		w.Dim{Parts: []w.Inline{w.Text("The organiser’s server may refuse the change, and its next update may undo it.")}, Size: w.Dim11_5, Variant: w.DimWrap},
		w.Line{Items: []w.LinePart{w.Check{On: true, Name: "group", Value: "past", Parts: []w.CheckPart{w.Text("Past meetings")}}, w.Drop{Trigger: w.Button{Label: "5 records", Menu: true, Small: true}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "3 Aug", Variant: w.Mono11}, Label: "Синк-ап по релизу", Href: "/app/calendar/nextcloud/rabota/dd01"},
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Detail: &w.Mono{Text: "27 Jul", Variant: w.Mono11}, Label: "Синк-ап по релизу", Href: "/app/calendar/nextcloud/rabota/dd02"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}}, Gap: w.LineGap8, Block: true},
		w.Dim{Parts: []w.Inline{w.Text("They will show the new address, as if the person had been invited at it.")}, Size: w.Dim11_5, Variant: w.DimWrap},
		w.SheetNote{Text: "2 records in Shared are read-only and keep the old address. Collections that are not ticked were not looked in."},
		w.Tally{Records: []string{"to.1 group=scheduling", "to.1 group=scheduling", "to.1 group=scheduling", "to.1 group=scheduling", "to.1 group=scheduling", "to.1 group=scheduling", "to.1 group=scheduling", "to.1 group=scheduling", "to.1 group=scheduling group=others", "to.1 group=others", "to.1 group=others", "to.1 group=past", "to.1 group=past", "to.1 group=past", "to.1 group=past", "to.1 group=past", "to.1 group=unknown", "to.1 group=unknown", "to.1", "to.1", "to.2", "to.2", "to.2"}}}, Foot: []w.SheetFootPart{w.Button{Label: "Rewrite 20 records", Kind: w.ButtonPrimary, Count: "Rewrite {n} records", CountOne: "Rewrite 1 record", Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Leave as is", To: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/rewrite?op=leave", Submit: true}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth520}, "s-contacts-15", "tc-sheet", 1),
	overlay("plain", w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "print", Label: "Print"}}, Width: w.MenuWidth212}, "s-files-7", "tc-menu", 1),
	overlay("plain", w.Sheet{Action: "/app/settings/unlock", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "back", Value: "/app/settings/devices"}}, Title: "Generate profile", Body: []w.SheetPart{
		w.Field{Label: "Your Carrel password", Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "password", Label: "Your Carrel password", Focus: true, Required: true}}, Width: w.FieldW100},
		w.Dim{Parts: []w.Inline{w.Text("Unlocks stored credentials for five minutes.")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Foot: []w.SheetFootPart{w.Button{Label: "Generate", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/devices"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth400}, "s-plan-3", "tc-sheet", 1),
	overlay("plain", w.Sheet{Action: "/app/publish?c=lichnyj", Title: "Publish a calendar", Head: []w.SheetHeadPart{w.Badge{Label: "local", Tone: w.BadgeLocal, Title: "Local instance: the link will not open from outside your network"}, w.Dim{Parts: []w.Inline{w.Text("nextcloud · Личный")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Field{Label: "How much of it", Parts: []w.FieldPart{w.Seg{Options: []w.SegOption{{Label: "A date range", On: true, Value: "range"}, {Label: "The whole calendar", Value: "whole", Disables: "from to"}}, Name: "how"}}, Width: w.FieldW100},
		w.Line{Items: []w.LinePart{w.Field{Label: "From", Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true}}, Width: w.FieldSM}, w.Field{Label: "To", Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-31", Label: "To", Mono: true}}, Width: w.FieldSM}}, Gap: w.LineGap10, Block: true},
		w.Field{Label: "What the visitor sees", Parts: []w.FieldPart{w.Group{Items: []w.FieldPart{w.Check{On: true, Name: "time", Value: "1", Reason: "Always shown", Parts: []w.CheckPart{w.Text("Time"), w.Dim{Parts: []w.Inline{w.Text("always shown")}, Size: w.Dim11_5}}}, w.Check{On: true, Name: "title", Value: "1", Parts: []w.CheckPart{w.Text("Title")}}, w.Check{Name: "location", Value: "1", Parts: []w.CheckPart{w.Text("Location")}}, w.Check{Name: "description", Value: "1", Parts: []w.CheckPart{w.Text("Description")}}, w.Check{Name: "attendees", Value: "1", Parts: []w.CheckPart{w.Text("Attendees")}}, w.Check{Name: "attachments", Value: "1", Parts: []w.CheckPart{w.Text("Attachments")}}, w.Check{Name: "tags", Value: "1", Parts: []w.CheckPart{w.Text("Tags")}}}, Form: w.GroupColumnLoose}, w.FieldHint{Parts: []w.Inline{w.Text("Without the title an event shows as “Busy”.")}}}, Width: w.FieldW100},
		w.SheetChecks{Items: []w.Check{w.Check{Name: "download", Value: "1", Parts: []w.CheckPart{w.Text("Allow "), w.Mono{Text: ".ics"}, w.Text(" download and subscription")}}}},
		w.Field{Label: "Link lives for", Parts: []w.FieldPart{w.Drop{Trigger: w.Pick{Value: "30 days"}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuItem{Label: "7 days", Name: "lives", Value: "7"},
			w.MenuItem{Label: "14 days", Name: "lives", Value: "14"},
			w.MenuItem{Label: "30 days", On: true, Name: "lives", Value: "30"},
			w.MenuSep{},
			w.MenuItem{Label: "No expiry", Name: "lives", Value: "never"}}, Width: w.MenuWidth212, Pop: w.MenuUnder}}}, Width: w.FieldSM}}, Foot: []w.SheetFootPart{w.Button{Label: "Create link", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/calendar"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}}, "s-publish-2", "tc-sheet", 1),
	overlay("close", w.Sheet{ID: "letter", Action: "/app/mail/send", Icon: "mail", Title: "Letter", Head: []w.SheetHeadPart{w.Tag{Label: "from mix@carrel.example", On: true, Place: w.Pushed}, w.Close{Label: "Close", Href: "/app/contacts"}}, Body: []w.SheetPart{
		w.Field{Label: "To", Parts: []w.FieldPart{w.Drop{Trigger: w.TagInput{Parts: []w.TagInputPart{w.Tag{Label: "Анна Ковалёва", Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, Removable: true, Aside: &w.Dim{Parts: []w.Inline{w.Text("· a.kovaleva@example.org")}}, Title: "a.kovaleva@example.org", Name: "to", Value: "a.kovaleva@example.org"}, w.Tag{Label: "Артём Белов", Mark: &w.Mark{Colour: "#6B4E7A", Size: w.MarkH11}, Removable: true, Aside: &w.Dim{Parts: []w.Inline{w.Text("· a.belov@example.org")}}, Title: "a.belov@example.org", Name: "to", Value: "a.belov@example.org"}}, Name: "to", Add: &w.TagAdd{Name: "to.add", Placeholder: "Name or address", Label: "Add a recipient", Suggest: "/app/contacts/suggest", Enter: &w.TagEnter{Name: "op", Value: "add", Label: "Add", Of: "letter"}}, Focus: true}, Menu: w.Menu{Items: []w.MenuPart{}, Width: w.MenuWidth420, Pop: w.MenuUnder}}, w.FieldHint{Parts: []w.Inline{w.Text("Up to 10 recipients. Each sees the others' addresses.")}}}, Width: w.FieldW100},
		w.Field{Label: "Subject", Parts: []w.FieldPart{w.Input{Name: "subject", Value: "Контакты подрядчика и смета", Label: "Subject"}}, Width: w.FieldW100},
		w.Area{Text: "Анна, Артём — вот карточка подрядчика и смета, которую он прислал. Договор в папке, ссылка ниже.", Variant: w.AreaLetter, Name: "text", Label: "Text"},
		w.Sec{Items: []w.SecPart{
			w.Rubric{Text: "Attached", Items: []w.RubricPart{w.Num{Value: 4}, w.Line{Items: []w.LinePart{w.Drop{Trigger: w.Button{Label: "Attach", Menu: true, Icon: "clip", Small: true}, Menu: w.Menu{Items: []w.MenuPart{
				w.MenuItem{Icon: "con", Label: "Contact", To: "/app/mail/send?op=attach&what=contact"},
				w.MenuItem{Icon: "con", Label: "Address book", To: "/app/mail/send?op=attach&what=book"},
				w.MenuItem{Icon: "cal", Label: "Event", To: "/app/mail/send?op=attach&what=event"},
				w.MenuItem{Icon: "cal", Label: "Calendar", To: "/app/mail/send?op=attach&what=calendar"},
				w.MenuItem{Icon: "task", Label: "Task", To: "/app/mail/send?op=attach&what=task"},
				w.MenuItem{Icon: "note", Label: "Note", To: "/app/mail/send?op=attach&what=note"},
				w.MenuSep{},
				w.MenuItem{Icon: "doc", Label: "File", To: "/app/mail/send?op=attach&what=file"},
				w.MenuItem{Icon: "file", Trail: w.Tag{Label: "link", Place: w.Pushed}, Label: "Folder", To: "/app/mail/send?op=attach&what=folder"}}, Width: w.MenuWidth212, Pop: w.MenuUnderEnd}}}, Gap: w.LineGap6, Place: w.Pushed}}, Variant: w.RubricRow},
			w.SecFiles{Items: []w.SecFile{{Icon: "con", Name: "Пётр Смирнов · подрядчик", Size: w.Mono{Text: ".vcf · 3 KB", Variant: w.MonoDusty11}, Tag: w.Tag{Label: "attached"}, Remove: w.Button{Icon: "close", Small: true, Title: "Remove", To: "/app/mail/send?op=remove&i=0", Submit: true}}, {Icon: "cal", Name: "Встреча на объекте · 18 August 10:00", Size: w.Mono{Text: ".ics · 2 KB", Variant: w.MonoDusty11}, Tag: w.Tag{Label: "attached"}, Remove: w.Button{Icon: "close", Small: true, Title: "Remove", To: "/app/mail/send?op=remove&i=1", Submit: true}}, {Icon: "doc", Name: "смета-подрядчик.pdf", Size: w.Mono{Text: "4.2 MB", Variant: w.MonoDusty11}, Tag: w.Tag{Label: "as a link", On: true}, Terms: &w.Mono{Text: "30 days", Variant: w.MonoDusty11}, Change: &w.Button{Label: "Change…", Small: true, To: "/app/mail/send?op=change&i=2", Submit: true}, Remove: w.Button{Icon: "close", Small: true, Title: "Remove", To: "/app/mail/send?op=remove&i=2", Submit: true}}, {Icon: "file", Name: "Договор подряда · 1 folder · 6 files", Size: w.Mono{Text: "folder", Variant: w.MonoDusty11}, Tag: w.Tag{Label: "as a link", On: true}, Terms: &w.Mono{Text: "30 days", Variant: w.MonoDusty11}, Change: &w.Button{Label: "Change…", Small: true, To: "/app/mail/send?op=change&i=3", Submit: true}, Remove: w.Button{Icon: "close", Small: true, Title: "Remove", To: "/app/mail/send?op=remove&i=3", Submit: true}}}},
			w.Msg{Parts: []w.MsgPart{w.Text("The PDF (over 2 MB) and the folder go as links; revoking one in "), w.Strong("Published links"), w.Text(" closes it for this letter too.")}, Tone: w.MsgAccent, Space: w.MsgFlushTop11}}, Flush: true}}, Foot: []w.SheetFootPart{w.Button{Label: "Send", Icon: "mail", Kind: w.ButtonPrimary, Busy: "Sending…", Submit: true}, w.Kbd{Keys: "Ctrl+Enter"}, w.Button{Label: "Preview", To: "/app/mail/send?op=preview", Submit: true}, w.Dim{Parts: []w.Inline{w.Text("Signed "), w.Strong("Михаил Енученко"), w.Text(" · replies go to mix@example.org")}, Size: w.Dim11_5, Variant: w.DimWrap}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", To: "/app/mail/send?op=close", Submit: true}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth680}, "s-mail-1", "tc-sheet", 1),
	overlay("close", w.Sheet{Action: "/app/mail/send", Icon: "mail", Title: "Send this link", Head: []w.SheetHeadPart{w.SheetSubject{Mark: w.Mark{Colour: "#4A6B52"}, Text: "Личный · 10–31 August"}, w.Close{Label: "Close", Href: "/app/files/published"}}, Body: []w.SheetPart{
		w.Field{Label: "To", Parts: []w.FieldPart{w.Input{Name: "to", Value: "n.orlova@example.org", Label: "To", Focus: true}}, Width: w.FieldW100},
		w.Field{Label: "A line from you", Parts: []w.FieldPart{w.Input{Name: "line", Placeholder: "optional", Label: "A line from you"}}, Width: w.FieldW100},
		w.SheetFrom{Label: "Sent as", Lines: [][]w.Inline{[]w.Inline{w.Strong("Михаил Енученко"), w.Text(" "), w.Mono{Text: "<mix@carrel.example>", Variant: w.Mono11}}, []w.Inline{w.Text("Replies go to "), w.Mono{Text: "mix@example.org", Variant: w.Mono11}, w.Text(" · "), w.Link{Text: "change in Account", Href: "#"}}}},
		w.Sec{Items: []w.SecPart{
			w.Rubric{Text: "The letter carries"},
			w.SecList{Items: [][]w.SecItemPart{[]w.SecItemPart{w.Text("The link and your line — nothing of the calendar itself")}, []w.SecItemPart{w.Text("Revoking the link closes it for this recipient too")}, []w.SecItemPart{w.Text("It expires on 8 September, as the link does")}}, Gap: 5, Muted: true}}, Flush: true}}, Foot: []w.SheetFootPart{w.Button{Label: "Send", Icon: "mail", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Button{Label: "Copy link instead", To: "/app/mail/send?op=link", Submit: true}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/files/published"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth680}, "s-mail-4", "tc-sheet", 1),
	overlay("pick", w.Sheet{Action: "/app/mail/attach?what=note&uid=e41a", Icon: "note", Title: "Attach a note", Head: []w.SheetHeadPart{w.Filter{Query: w.Query{Value: "подряд", Label: "Search", Name: "q", Of: "attach-search"}, Quiet: true, Variant: w.FilterT12W200, Place: w.Pushed}, w.Close{Label: "Close", Href: "/app/mail/compose"}}, Body: []w.SheetPart{
		w.SheetPick{Rail: []w.SheetPickRail{w.Tree{Rows: []w.TreeRow{
			{Mark: &w.Mark{Kind: w.MarkMulti}, Label: "All notebooks", Href: "/app/mail/attach?what=note&q=подряд"},
			{Icon: "server", Label: "nextcloud"},
			{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "Записи", Indent: w.TreeLevel22, Href: "/app/mail/attach?what=note&q=подряд&c=zapisi"},
			{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "Рабочее", Indent: w.TreeLevel22, Href: "/app/mail/attach?what=note&q=подряд&c=rabochee"}}, On: 1, Variant: w.TreePadded}}, List: []w.SheetPickList{w.ListGroup{Title: w.Rubric{Text: "Found", Form: w.RubricInline}, Count: &w.Num{Value: 3}}, w.List{Rows: []w.ListRow{
			w.BriefRow{Mark: w.Mark{Colour: "#3F6E8C"}, Title: "Встреча с подрядчиком, итоги", Date: w.Mono{Text: "18 Aug"}, Where: "nextcloud · Рабочее", On: true, Href: "/app/mail/attach?what=note&q=подряд&uid=e41a"},
			w.BriefRow{Mark: w.Mark{Colour: "#3F6E8C"}, Title: "Подрядчики: кого смотрели", Date: w.Mono{Text: "02 Aug"}, Where: "nextcloud · Рабочее", Href: "/app/mail/attach?what=note&q=подряд&uid=c27b"},
			w.BriefRow{Mark: w.Mark{Colour: "#4A6B52"}, Title: "Ремонт: что осталось", Date: w.Mono{Text: "27 Jul"}, Where: "nextcloud · Записи", Href: "/app/mail/attach?what=note&q=подряд&uid=9a03"}}}}, Short: true},
		w.Sec{Items: []w.SecPart{
			w.Field{Label: "Send it as", Parts: []w.FieldPart{w.Seg{Options: []w.SegOption{{Label: "Markdown · .md", On: true, Value: "md"}, {Label: "Source · .ics", Value: "ics"}}, Variant: w.SegStart, Name: "format"}, w.FieldHint{Parts: []w.Inline{w.Text("Markdown is for reading: the text and links to attachments, without the properties other clients wrote. The source "), w.Mono{Text: ".ics"}, w.Text(" keeps everything and needs a client that reads VJOURNAL — jtx Board, Evolution.")}}}, Width: w.FieldW100}}}}, Flush: true, Foot: []w.SheetFootPart{w.Button{Label: "Attach", Icon: "clip", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Dim{Parts: []w.Inline{w.Text("Встреча с подрядчиком, итоги · "), w.Mono{Text: ".md"}, w.Text(" · 4 KB · goes as a file")}, Size: w.Dim11_5, Variant: w.DimWrap}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/mail/compose"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth900}, "s-mail-5", "tc-sheet", 1),
	overlay("pick", w.Sheet{Icon: "cal", Title: "Attach a calendar", Head: []w.SheetHeadPart{w.Close{Label: "Close", Place: w.Pushed, Href: "/app/mail/compose"}}, Body: []w.SheetPart{
		w.SheetPick{Rail: []w.SheetPickRail{w.Tree{Rows: []w.TreeRow{
			{Icon: "server", Label: "nextcloud"},
			{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "Личный", Indent: w.TreeLevel22, Href: "/app/mail/attach?what=calendar&c=lichnyj"},
			{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Indent: w.TreeLevel22, Href: "/app/mail/attach?what=calendar&c=rabota"},
			{Icon: "server", Label: "radicale · vps"},
			{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Indent: w.TreeLevel22, Href: "/app/mail/attach?what=calendar&c=proekty"}}, On: 2, Variant: w.TreePadded}}, List: []w.SheetPickList{w.Form{Items: []w.FormPart{w.Field{Label: "What goes in", Parts: []w.FieldPart{w.Seg{Options: []w.SegOption{{Label: "Whole calendar", Value: "whole", Disables: "from to"}, {Label: "Dates", On: true, Value: "dates"}}, Variant: w.SegStart, Name: "span"}}, Width: w.FieldW100}, w.Field{Label: "From", Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true}}, Width: w.FieldSM}, w.Field{Label: "To", Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-31", Label: "To", Mono: true}, w.FieldHint{Parts: []w.Inline{w.Text("Both days are in the file.")}}}, Width: w.FieldSM}}, ID: "attach-calendar", Action: "/app/mail/attach?what=calendar&c=lichnyj", Recount: "attach-line"}}, Short: true}}, Flush: true, Foot: []w.SheetFootPart{w.Button{Label: "Attach", Icon: "clip", Kind: w.ButtonPrimary, Form: "attach-calendar"}, w.Kbd{Keys: "Enter"}, w.Dim{Parts: []w.Inline{w.Text("Calendar «Личный» · 10 August 2026 – 31 August 2026 · "), w.Mono{Text: ".ics"}, w.Text(" · 5 KB · goes as a file")}, Size: w.Dim11_5, Variant: w.DimWrap, ID: "attach-line"}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/mail/compose"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth900}, "s-mail-10", "tc-sheet", 1),
	overlay("pick", w.Sheet{Action: "/app/mail/attach?what=file&path=sftpgo/documents/stroyka/podryad/smeta-podryadchik.pdf", Icon: "doc", Title: "Attach a file", Head: []w.SheetHeadPart{w.Crumbs{Path: []w.Crumb{{Label: "Documents", Href: "#"}, {Label: "Стройка", Href: "#"}}, Current: "Подряд"}, w.Close{Label: "Close", Href: "/app/mail/compose"}}, Body: []w.SheetPart{
		w.SheetPick{Rail: []w.SheetPickRail{w.Tree{Rows: []w.TreeRow{
			{Icon: "server", Label: "sftpgo · vps", Href: "/app/mail/attach?what=file&dir=sftpgo"},
			{Icon: "file", Label: "Documents", Indent: w.TreeLevel22, Href: "/app/mail/attach?what=file&dir=sftpgo/documents"},
			{Icon: "file", Label: "Стройка", Indent: w.TreeLevel36, Href: "/app/mail/attach?what=file&dir=sftpgo/documents/stroyka"},
			{Icon: "file", Label: "Договор подряда", Indent: w.TreeLevel44, Href: "/app/mail/attach?what=file&dir=sftpgo/documents/stroyka/dogovor-podryada"},
			{Icon: "file", Label: "Подряд", Indent: w.TreeLevel44, Href: "/app/mail/attach?what=file&dir=sftpgo/documents/stroyka/podryad"},
			{Icon: "file", Label: "Личное", Indent: w.TreeLevel36, Href: "/app/mail/attach?what=file&dir=sftpgo/documents/lichnoe"}}, On: 5, Variant: w.TreePadded}}, List: []w.SheetPickList{w.Table{Cols: []w.TableCol{{Label: "Name"}, {Label: "Size", Width: w.Col90}, {Label: "Changed", Width: w.Col110, Rank: w.Rank3}, {Width: w.ColType, Rank: w.Rank4}}, Rows: []w.TableRow{
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "смета-подрядчик.pdf", Href: "/app/mail/attach?what=file&path=sftpgo/documents/stroyka/podryad/smeta-podryadchik.pdf"}}}, {Parts: []w.CellPart{w.Text("4.2 MB")}, Mono: true}, {Parts: []w.CellPart{w.Text("12 Aug")}, Mono: true}, {Parts: []w.CellPart{w.Tag{Label: "over 2 MB · as a link", On: true}}}}, State: w.RowOn},
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "договор-подряда.pdf", Href: "/app/mail/attach?what=file&path=sftpgo/documents/stroyka/podryad/dogovor-podryada.pdf"}}}, {Parts: []w.CellPart{w.Text("840 KB")}, Mono: true}, {Parts: []w.CellPart{w.Text("09 Aug")}, Mono: true}, {Parts: []w.CellPart{w.Tag{Label: "as a file"}}}}},
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.TableFile{Icon: "img", Name: "объект-фасад.jpg", Href: "/app/mail/attach?what=file&path=sftpgo/documents/stroyka/podryad/obekt-fasad.jpg"}}}, {Parts: []w.CellPart{w.Text("1.6 MB")}, Mono: true}, {Parts: []w.CellPart{w.Text("09 Aug")}, Mono: true}, {Parts: []w.CellPart{w.Tag{Label: "as a file"}}}}}}}}, Short: true},
		w.Msg{Parts: []w.MsgPart{w.Text("Over the attachment limit: the letter carries a link instead, for 30 days. Revoke it in "), w.Strong("Published links"), w.Text(".")}, Tone: w.MsgAccent, Space: w.MsgSheet}}, Flush: true, Foot: []w.SheetFootPart{w.Button{Label: "Attach", Icon: "clip", Kind: w.ButtonPrimary, Submit: true}, w.Kbd{Keys: "Enter"}, w.Dim{Parts: []w.Inline{w.Text("смета-подрядчик.pdf · 4.2 MB · goes as a link")}, Size: w.Dim11_5, Variant: w.DimWrap}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/mail/compose"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed}}, Width: w.SheetWidth900}, "s-mail-6", "tc-sheet", 1),
	overlay("close", w.Sheet{Title: "Sources", Head: []w.SheetHeadPart{w.Button{Label: "Apply", Small: true, Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.RailSec{Items: []w.RailSecPart{
			w.TickRow{Ticked: w.TickedPart, Kind: w.TickRowAll, Mark: w.Mark{Kind: w.MarkMulti}, Label: "Combined notebook", Meta: "86", Current: true}}},
		w.RailSec{Items: []w.RailSecPart{
			w.Rubric{Text: "nextcloud"},
			w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "64"},
			w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "22"},
			w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "—", Hidden: true}}}}}, "phone-9", "tc-sheet", 1),

	// tc-scrim
	overlay("plain", w.Scrim{Closed: true}, "shell-1", "tc-scrim", 1),
	overlay("pad", w.Scrim{Content: w.Sheet{Icon: "move", Title: "Move contact", Head: []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("Анна Ковалёва · from nextcloud · Контакты")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Body: []w.SheetPart{
		w.Rubric{Text: "Write into"},
		w.Tree{Rows: []w.TreeRow{
			{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Unavailable: true, Trail: w.Mono{Text: "already there", Variant: w.Mono10Dusty}},
			{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "Архив", Trail: w.Mono{Text: "nextcloud", Variant: w.Mono10Dusty}},
			{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Trail: w.Mono{Text: "radicale", Variant: w.Mono10Dusty}},
			{Mark: &w.Mark{Colour: "#8C6239"}, Label: "Shared", Unavailable: true, Trail: w.Badge{Label: "read-only"}}}, On: 2},
		w.Msg{Parts: []w.MsgPart{w.Text("Inside one server this is "), w.Mono{Text: "MOVE", Form: w.MonoCode}, w.Text(": instant, no traffic. To another address book the card is copied and then deleted — Carrel says so before it starts, not after. Copy is this same sheet with "), w.Strong("Copy"), w.Text(".")}, Space: w.MsgFlushTop2}}, Foot: []w.SheetFootPart{w.Button{Label: "Move", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Enter"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}, Width: w.SheetWidth520}, Pad: w.ScrimPad26}, "s-contacts-2", "tc-scrim", 1),
	overlay("plain", w.Scrim{Content: w.Sheet{Icon: "quill", Title: "New note", Head: []w.SheetHeadPart{w.Pick{Value: "nextcloud · Записи", Variant: w.PickT12W200, Label: "Where", Place: w.Pushed}, w.Close{Label: "Close"}}, Body: []w.SheetPart{
		w.Area{Text: "Позвонить в МФЦ во вторник до 12:00 — уточнить, нужен ли оригинал доверенности.", Variant: w.AreaBodySm, Focus: true, Label: "Text"}}, Foot: []w.SheetFootPart{w.Button{Label: "Save note", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Ctrl+Enter"}, w.Button{Label: "Open the card", Icon: "expand"}, w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}}}}}, "s-notes-5", "tc-scrim", 1),
	overlay("pad", w.Scrim{Content: w.Empty{Icon: "import", Lead: "Drop to import", Parts: []w.EmptyPart{w.Aside{Face: w.FaceLarge, Parts: []w.Inline{w.Text("One"), w.Mono{Text: ".vcf"}, w.Text("file into"), w.Strong("nextcloud · Контакты")}}, w.Dim{Parts: []w.Inline{w.Text("The preview opens first: nothing is written until it has been read")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Variant: w.EmptyDrop}, Pad: w.ScrimPad38}, "s-io-1", "tc-scrim", 1),

	// tc-overlay
	overlay("plain", w.Overlay{
		Body: w.OverlayBody{
			Rail: []w.OverlayRailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Settings"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "user", Label: "Account", Href: "#"},
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 4}, Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 2,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Text:   "davloom · running · 3 devices",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last response "), w.Mono{Text: "2026-08-16 09:41"}}, Size: w.Dim11},
						},
						w.RailStatus{
							Text: "Backups are running",
							Detail: w.Dim{
								Parts: []w.Inline{
									w.Text("Last backup "),
									w.Mono{Text: "2026-08-16 03:10"},
									w.Text(" · "),
									w.Hint{Text: "verified", Title: "The archive was read back from the destination and the checksum matched"},
								},
								Size: w.Dim11,
							},
						},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Connections"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{w.Text("Signed in as "), w.Strong("mix"), w.Text(" · 4 servers · 19 collections · 2 subscriptions")},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Connect account", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Subscribe to a feed", Icon: "link"},
							},
						},
					},
					w.PostForm{
						ID:     "conn-picks",
						Action: "/app/settings/connections?op=pick",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
					w.MainResults{
						Items: []w.ResultsPart{
							w.Rubric{
								Text:    "DAV servers",
								Items:   []w.RubricPart{w.Mono{Text: "4", Variant: w.MonoTight}},
								Variant: w.RubricSplit,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "nextcloud",
										Items: []w.RubricPart{
											w.Line{
												Items: []w.LinePart{
													w.Button{Label: "New collection", Icon: "plus", Small: true},
													w.Button{Label: "Show password", Small: true},
													w.Button{Label: "Test", Small: true},
													w.Button{
														Label: "Update collection list",
														Small: true,
														Form:  "conn-picks",
														To:    "/app/settings/connections?op=rediscover&account=nextcloud",
													},
													w.Button{Label: "Edit", Small: true},
													w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true},
												},
												Gap: w.LineGap6,
											},
										},
									},
									w.Mono{
										Text:    "https://cloud.example.org/remote.php/dav/ · mix · last answer 2 min ago",
										Form:    w.MonoLine,
										Variant: w.Mono11_5Dusty,
									},
									w.Table{
										Cols: []w.TableCol{
											{Width: w.ColTickSmall},
											{Label: "Collection"},
											{Label: "What is in it", Rank: w.Rank3},
											{Label: "Objects", Width: w.ColNum, Rank: w.Rank4},
											{Label: "Changed on the server", Width: w.ColChanged},
											{Width: w.ColMenu},
										},
										Rows: []w.TableRow{
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|lichnyj"}},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Личный"}}},
													{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
													{Parts: []w.CellPart{w.Text("2")}, Mono: true},
													{Parts: []w.CellPart{w.Text("3 h ago")}, Mono: true},
													{
														Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
														Acts:  true,
													},
												},
											},
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|dela"}},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Дела"}}},
													{Parts: []w.CellPart{w.Text("tasks")}, Mono: true},
													{Parts: []w.CellPart{w.Text("4")}, Mono: true},
													{Parts: []w.CellPart{w.Text("20 min ago")}, Mono: true},
													{
														Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
														Acts:  true,
													},
												},
											},
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|rabota"}},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Работа"}}},
													{Parts: []w.CellPart{w.Text("calendar · tasks")}, Mono: true},
													{Parts: []w.CellPart{w.Text("4")}, Mono: true},
													{Parts: []w.CellPart{w.Text("yesterday")}, Mono: true},
													{
														Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
														Acts:  true,
													},
												},
											},
											{
												Cells: []w.TableCell{
													{
														Pick: ptr(w.TickedOff),
														Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|dni-rozhdeniya"},
													},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "Дни рождения"}}},
													{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
													{Parts: []w.CellPart{w.Text("1")}, Mono: true},
													{Parts: []w.CellPart{w.Text("12 August")}, Mono: true},
													{
														Parts: []w.CellPart{
															w.Drop{
																Open:    true,
																Trigger: w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
																Menu: w.Menu{
																	Items: []w.MenuPart{
																		w.MenuItem{Icon: "pencil", Label: "Rename and colour"},
																		w.MenuItem{Icon: "move", Label: "Move to another server"},
																		w.MenuItem{Icon: "dup", Label: "Copy to another server"},
																		w.MenuItem{Icon: "export", Detail: &w.Mono{Text: ".ics", Variant: w.Mono11}, Label: "Export"},
																		w.MenuItem{Icon: "link", Label: "Publish a link"},
																		w.MenuSep{},
																		w.MenuItem{Icon: "trash", Label: "Delete on the server…", Tone: w.MenuToneDanger},
																	},
																	Width: w.MenuWidth212,
																	Pop:   w.MenuUnderEnd,
																},
															},
														},
														Acts: true,
													},
												},
												State: w.RowOff,
											},
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|zapisi"}},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Записи"}}},
													{Parts: []w.CellPart{w.Text("journal")}, Mono: true},
													{Parts: []w.CellPart{w.Text("2")}, Mono: true},
													{Parts: []w.CellPart{w.Text("2 h ago")}, Mono: true},
													{
														Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
														Acts:  true,
													},
												},
											},
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|dnevnik"}},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Дневник"}}},
													{Parts: []w.CellPart{w.Text("journal")}, Mono: true},
													{Parts: []w.CellPart{w.Text("1")}, Mono: true},
													{Parts: []w.CellPart{w.Text("5 d ago")}, Mono: true},
													{
														Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
														Acts:  true,
													},
												},
											},
											{
												Cells: []w.TableCell{
													{
														Pick: ptr(w.TickedOff),
														Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|zapiski-obshchie"},
													},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Общие с семьёй"}}},
													{Parts: []w.CellPart{w.Text("journal "), w.Badge{Label: "ro"}}, Mono: true},
													{Parts: []w.CellPart{w.Text("0")}, Mono: true},
													{Parts: []w.CellPart{w.Text("10 August")}, Mono: true},
													{
														Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
														Acts:  true,
													},
												},
												State: w.RowOff,
											},
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|kontakty"}},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Контакты"}}},
													{Parts: []w.CellPart{w.Text("address book")}, Mono: true},
													{Parts: []w.CellPart{w.Text("5")}, Mono: true},
													{Parts: []w.CellPart{w.Text("4 d ago")}, Mono: true},
													{
														Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
														Acts:  true,
													},
												},
											},
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|arhiv"}},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "Архив"}}},
													{Parts: []w.CellPart{w.Text("address book")}, Mono: true},
													{Parts: []w.CellPart{w.Text("1")}, Mono: true},
													{Parts: []w.CellPart{w.Text("28 July")}, Mono: true},
													{
														Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
														Acts:  true,
													},
												},
											},
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|files"}},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Files"}}},
													{Parts: []w.CellPart{w.Text("files")}, Mono: true},
													{Parts: []w.CellPart{w.Text("1")}, Mono: true},
													{Parts: []w.CellPart{w.Text("6 d ago")}, Mono: true},
													{
														Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
														Acts:  true,
													},
												},
											},
										},
									},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "radicale · vps",
										Items: []w.RubricPart{
											w.Line{
												Items: []w.LinePart{
													w.Button{Label: "New collection", Icon: "plus", Small: true},
													w.Button{Label: "Show password", Small: true},
													w.Button{Label: "Test", Small: true},
													w.Button{
														Label: "Update collection list",
														Small: true,
														Form:  "conn-picks",
														To:    "/app/settings/connections?op=rediscover&account=radicale",
													},
													w.Button{Label: "Edit", Small: true},
													w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true},
												},
												Gap: w.LineGap6,
											},
										},
									},
									w.Mono{
										Text:    "https://dav.vps.example/radicale/ · mix · last answer 4 min ago",
										Form:    w.MonoLine,
										Variant: w.Mono11_5Dusty,
									},
									w.Table{
										Cols: []w.TableCol{
											{Width: w.ColTickSmall},
											{Label: "Collection"},
											{Label: "What is in it", Rank: w.Rank3},
											{Label: "Objects", Width: w.ColNum, Rank: w.Rank4},
											{Label: "Changed on the server", Width: w.ColChanged},
											{Width: w.ColMenu},
										},
										Rows: []w.TableRow{
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "radicale|proekty"}},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Проекты"}}},
													{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
													{Parts: []w.CellPart{w.Text("1")}, Mono: true},
													{Parts: []w.CellPart{w.Text("1 h ago")}, Mono: true},
													{
														Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
														Acts:  true,
													},
												},
											},
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "radicale|dezhurstva"}},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#8C6239"}, Name: "Дежурства"}}},
													{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
													{Parts: []w.CellPart{w.Text("1")}, Mono: true},
													{Parts: []w.CellPart{w.Text("yesterday")}, Mono: true},
													{
														Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
														Acts:  true,
													},
												},
											},
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "radicale|komanda"}},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Команда"}}},
													{Parts: []w.CellPart{w.Text("address book "), w.Badge{Label: "ro"}}, Mono: true},
													{Parts: []w.CellPart{w.Text("1")}, Mono: true},
													{Parts: []w.CellPart{w.Text("6 d ago")}, Mono: true},
													{
														Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
														Acts:  true,
													},
												},
											},
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOff), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "radicale|rassylki"}},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Рассылки"}}},
													{Parts: []w.CellPart{w.Text("address book")}, Mono: true},
													{Parts: []w.CellPart{w.Text("0")}, Mono: true},
													{Parts: []w.CellPart{w.Text("3 August")}, Mono: true},
													{
														Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
														Acts:  true,
													},
												},
												State: w.RowOff,
											},
										},
									},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "baikal",
										Items: []w.RubricPart{
											w.Line{
												Items: []w.LinePart{
													w.Button{Label: "Show password", Small: true},
													w.Button{Label: "Test again", Small: true},
													w.Button{Label: "Edit", Small: true},
													w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true},
												},
												Gap: w.LineGap6,
											},
										},
										Variant: w.RubricDanger,
									},
									w.Mono{
										Text:    "https://dav.office.example/dav.php/ · m.yenuchenko · last answer refused",
										Form:    w.MonoLine,
										Variant: w.Mono11_5DustyGap,
									},
									w.Diag{
										Text: "[well-known]  GET /.well-known/caldav → 404\n[base]        PROPFIND /dav.php/ → 401 Unauthorized\n              WWW-Authenticate: Basic realm=\"BaikalDAV\"\n[principal]   не запрашивался: предыдущий шаг не прошёл",
									},
									w.CardText{
										Kind: w.CardAlertText,
										Parts: []w.CardTextPart{
											w.Text("The server refused the password. It may have been changed on the server: the one Carrel saved no longer fits."),
										},
									},
								},
								Variant: w.CardAlert,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "sftpgo · vps",
										Items: []w.RubricPart{
											w.Line{
												Items: []w.LinePart{
													w.Button{Label: "New folder", Icon: "plus", Small: true},
													w.Button{Label: "Show password", Small: true},
													w.Button{Label: "Test", Small: true},
													w.Button{
														Label: "Update collection list",
														Small: true,
														Form:  "conn-picks",
														To:    "/app/settings/connections?op=rediscover&account=sftpgo",
													},
													w.Button{Label: "Edit", Small: true},
													w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true},
												},
												Gap: w.LineGap6,
											},
										},
									},
									w.Mono{
										Text:    "https://files.example.org/dav/ · mix · plain WebDAV, no principal · last answer 5 min ago",
										Form:    w.MonoLine,
										Variant: w.Mono11_5Dusty,
									},
									w.Table{
										Cols: []w.TableCol{
											{Width: w.ColTickSmall},
											{Label: "Folder"},
											{Label: "What is in it", Rank: w.Rank3},
											{Label: "Objects", Width: w.ColNum, Rank: w.Rank4},
											{Label: "Changed on the server", Width: w.ColChanged},
											{Width: w.ColMenu},
										},
										Rows: []w.TableRow{
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "sftpgo|documents"}},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Documents"}}},
													{Parts: []w.CellPart{w.Text("files")}, Mono: true},
													{Parts: []w.CellPart{w.Text("9")}, Mono: true},
													{Parts: []w.CellPart{w.Text("5 min ago")}, Mono: true},
													{
														Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
														Acts:  true,
													},
												},
											},
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOff), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "sftpgo|public"}},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#9A9280"}, Name: "Public"}}},
													{Parts: []w.CellPart{w.Text("files "), w.Badge{Label: "ro"}}, Mono: true},
													{Parts: []w.CellPart{w.Text("2")}, Mono: true},
													{Parts: []w.CellPart{w.Text("3 d ago")}, Mono: true},
													{
														Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
														Acts:  true,
													},
												},
												State: w.RowOff,
											},
										},
									},
								},
							},
							w.Rubric{
								Text:    "Subscriptions · external, read-only",
								Items:   []w.RubricPart{w.Mono{Text: "2", Variant: w.MonoTight}},
								Variant: w.RubricSplit,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "Праздники РФ",
										Items: []w.RubricPart{
											w.Badge{Label: "feed"},
											w.Line{
												Items: []w.LinePart{
													w.Button{Label: "Check", Small: true},
													w.Button{Label: "Edit", Small: true},
													w.Button{Label: "Unsubscribe", Kind: w.ButtonDanger, Small: true},
												},
												Gap: w.LineGap6,
											},
										},
									},
									w.Mono{
										Text:    "https://feeds.example/ru-holidays.ics · no password · read 2 min ago · 34 KB · 0.4 s",
										Form:    w.MonoLine,
										Variant: w.Mono11_5Dusty,
									},
									w.Table{
										Cols: []w.TableCol{
											{Width: w.ColTickSmall},
											{Label: "Lands in"},
											{Label: "Objects", Width: w.ColNum, Rank: w.Rank3},
											{Label: "Changed on the server", Width: w.ColChanged, Rank: w.Rank4},
										},
										Rows: []w.TableRow{
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "ru-holidays|calendar"}},
													{
														Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#8C6239", Kind: w.MarkExternal}, Name: "Calendar"}},
													},
													{Parts: []w.CellPart{w.Text("3")}, Mono: true},
													{Parts: []w.CellPart{w.Text("yesterday")}, Mono: true, Rows: 2},
												},
											},
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOff), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "ru-holidays|tasks"}},
													{
														Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#8C6239", Kind: w.MarkExternal}, Name: "Tasks"}},
													},
													{Parts: []w.CellPart{w.Text("1")}, Mono: true},
												},
												State: w.RowOff,
											},
										},
									},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "Расписание лиги",
										Items: []w.RubricPart{
											w.Badge{Label: "feed"},
											w.Line{
												Items: []w.LinePart{
													w.Button{Label: "Show password", Small: true},
													w.Button{Label: "Check again", Small: true},
													w.Button{Label: "Edit", Small: true},
													w.Button{Label: "Unsubscribe", Kind: w.ButtonDanger, Small: true},
												},
												Gap: w.LineGap6,
											},
										},
										Variant: w.RubricDanger,
									},
									w.Mono{
										Text:    "webcal://league.example/season.ics · password saved · last read yesterday",
										Form:    w.MonoLine,
										Variant: w.Mono11_5DustyGap,
									},
									w.Diag{
										Text: "[fetch]   GET /season.ics → 401 Unauthorized\n              WWW-Authenticate: Basic realm=\"league\"\n[cache]   2 events from yesterday are still shown, marked stale",
									},
									w.CardText{
										Kind: w.CardAlertText,
										Parts: []w.CardTextPart{
											w.Text("The feed asks for a password and the saved one no longer fits. A feed that stopped answering must not look like a feed with nothing in it."),
										},
									},
								},
								Variant: w.CardAlert,
							},
						},
					},
				},
			},
			Height: w.ShellHeight390,
		},
		Scrim: w.Scrim{
			Content: w.Sheet{
				Action: "/app/publish?c=lichnyj",
				Title:  "Publish a calendar",
				Head: []w.SheetHeadPart{
					w.Badge{
						Label: "local",
						Tone:  w.BadgeLocal,
						Title: "Local instance: the link will not open from outside your network",
					},
					w.Dim{
						Parts:   []w.Inline{w.Text("nextcloud · Личный")},
						Size:    w.Dim11_5,
						Variant: w.DimWrap,
						Place:   w.Pushed,
					},
				},
				Body: []w.SheetPart{
					w.Field{
						Label: "How much of it",
						Parts: []w.FieldPart{
							w.Seg{
								Options: []w.SegOption{
									{Label: "A date range", On: true, Value: "range"},
									{Label: "The whole calendar", Value: "whole", Disables: "from to"},
								},
								Name: "how",
							},
						},
						Width: w.FieldW100,
					},
					w.Line{
						Items: []w.LinePart{
							w.Field{
								Label: "From",
								Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true}},
								Width: w.FieldSM,
							},
							w.Field{
								Label: "To",
								Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-31", Label: "To", Mono: true}},
								Width: w.FieldSM,
							},
						},
						Gap:   w.LineGap10,
						Block: true,
					},
					w.Field{
						Label: "What the visitor sees",
						Parts: []w.FieldPart{
							w.Group{
								Items: []w.FieldPart{
									w.Check{
										On:     true,
										Name:   "time",
										Value:  "1",
										Reason: "Always shown",
										Parts:  []w.CheckPart{w.Text("Time"), w.Dim{Parts: []w.Inline{w.Text("always shown")}, Size: w.Dim11_5}},
									},
									w.Check{On: true, Name: "title", Value: "1", Parts: []w.CheckPart{w.Text("Title")}},
									w.Check{Name: "location", Value: "1", Parts: []w.CheckPart{w.Text("Location")}},
									w.Check{Name: "description", Value: "1", Parts: []w.CheckPart{w.Text("Description")}},
									w.Check{Name: "attendees", Value: "1", Parts: []w.CheckPart{w.Text("Attendees")}},
									w.Check{Name: "attachments", Value: "1", Parts: []w.CheckPart{w.Text("Attachments")}},
									w.Check{Name: "tags", Value: "1", Parts: []w.CheckPart{w.Text("Tags")}},
								},
								Form: w.GroupColumnLoose,
							},
							w.FieldHint{Parts: []w.Inline{w.Text("Without the title an event shows as “Busy”.")}},
						},
						Width: w.FieldW100,
					},
					w.SheetChecks{
						Items: []w.Check{
							w.Check{
								On:    true,
								Name:  "download",
								Value: "1",
								Parts: []w.CheckPart{w.Text("Allow "), w.Mono{Text: ".ics"}, w.Text(" download and subscription")},
							},
						},
					},
					w.Field{
						Label: "Link lives for",
						Parts: []w.FieldPart{
							w.Drop{
								Trigger: w.Pick{Value: "30 days"},
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuItem{Label: "7 days", Name: "lives", Value: "7"},
										w.MenuItem{Label: "14 days", Name: "lives", Value: "14"},
										w.MenuItem{Label: "30 days", On: true, Name: "lives", Value: "30"},
										w.MenuSep{},
										w.MenuItem{Label: "No expiry", Name: "lives", Value: "never"},
									},
									Width: w.MenuWidth212,
									Pop:   w.MenuUnder,
								},
							},
						},
						Width: w.FieldSM,
					},
				},
				Foot: []w.SheetFootPart{
					w.Button{Label: "Create link", Kind: w.ButtonPrimary, Submit: true},
					w.Button{Label: "Cancel", Href: "/app/calendar"},
				},
			},
			Pad: w.ScrimPad26,
		},
	}, "s-publish-2", "tc-overlay", 1),
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
	overlay("plain", w.Overlay{
		Body: w.OverlayBody{
			Rail: []w.OverlayRailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 5}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
					},
					On: 2,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedPart,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined calendar",
							Meta:    "6",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "2"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "1"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "baikal"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Error: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Subscriptions", Items: []w.RubricPart{w.Badge{Label: "external"}}},
						w.TickRow{
							Kind:  w.TickRowExternal,
							Mark:  w.Mark{Colour: "#8C6239"},
							Label: "Праздники РФ",
							Meta:  "0",
							Trail: w.Badge{Label: "ro"},
						},
						w.TickRow{
							Ticked: w.TickedOff,
							Kind:   w.TickRowExternal,
							Mark:   w.Mark{Colour: "#6B4E7A"},
							Label:  "Расписание лиги",
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "релиз"}, w.Tag{Label: "работа"}},
				},
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined calendar"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{
									w.Text("5 of 7 calendars · 3 accounts · 10 – 16 August 2026 · "),
									w.Mono{Text: "Europe/Moscow"},
									w.Text(" · 6 events"),
								},
							},
						},
						Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}}},
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "10:00",
								Title:  []w.Inline{w.Text("Синк-ап по релизу")},
								Detail: "Переговорная 2",
								Badges: []w.Badge{{Label: "4 attendees"}},
								Source: "nextcloud · Личный",
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#3F6E8C"},
								Time:   "all day",
								Title:  []w.Inline{w.Text("Дежурство: первая линия")},
								Detail: "—",
								Badges: []w.Badge{{Label: "repeats weekly"}},
								Source: "nextcloud · Работа",
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#8C6239"},
								Time:   "all day",
								Title:  []w.Inline{w.Text("Дежурство: первая линия")},
								Detail: "—",
								Badges: []w.Badge{{Label: "repeats weekly"}},
								Source: "radicale · vps · Дежурства",
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#3F6E8C"},
								Time:   "15:00",
								Title:  []w.Inline{w.Text("Созвон с подрядчиком")},
								Detail: "Jitsi",
								Badges: []w.Badge{{Label: "tentative", Tone: w.BadgeAlert}},
								Source: "nextcloud · Работа",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Tuesday · 11 August", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#6B4E7A"},
								Time:   "09:30",
								Title:  []w.Inline{w.Text("Ревью архитектуры")},
								Detail: "Переговорная 1",
								Badges: []w.Badge{{Label: "2 attendees"}},
								Source: "radicale · vps · Проекты",
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "18:00",
								Title:  []w.Inline{w.Text("Забрать документы")},
								Detail: "МФЦ, Мясницкая 13",
								Source: "nextcloud · Личный",
							},
						},
					},
				},
			},
		},
		Scrim: w.Scrim{
			Content: w.Sheet{
				Icon:  "quill",
				Title: "New note",
				Head: []w.SheetHeadPart{
					w.Pick{Value: "nextcloud · Записи", Variant: w.PickT12W200, Place: w.Pushed},
					w.Close{Label: "Close"},
				},
				Body: []w.SheetPart{
					w.Area{
						Text:    "Позвонить в МФЦ во вторник до 12:00 — уточнить, нужен ли оригинал доверенности.",
						Variant: w.AreaBodySm,
						Focus:   true,
						Label:   "Text",
					},
				},
				Foot: []w.SheetFootPart{
					w.Button{Label: "Save note", Kind: w.ButtonPrimary},
					w.Kbd{Keys: "Ctrl+Enter"},
					w.Button{Label: "Open the card", Icon: "expand"},
					w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}},
				},
			},
		},
	}, "s-notes-5", "tc-overlay", 1),
	overlay("plain", w.Overlay{
		Body: w.OverlayBody{
			Rail: []w.OverlayRailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 5}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 1,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedPart,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined address book",
							Meta:    "5",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "3"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "1", Trail: w.Badge{Label: "ro"}},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#8C6239"}, Label: "Рассылки", Meta: "—", Hidden: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "работа"}, w.Tag{Label: "издательство"}},
				},
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined address book"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 5 cards")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New contact", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter the loaded cards", Label: "Filter the loaded cards"},
								Waiting: true,
								Variant: w.FilterH26W220,
							},
							w.Seg{
								Options: []w.SegOption{
									{Label: "A–Z", Href: "/app/contacts?sort=az", On: true},
									{Label: "Recently changed", Href: "/app/contacts?sort=changed"},
								},
								Secondary: true,
							},
							w.Button{Label: "Columns", Icon: "cols"},
							w.Drop{
								Trigger: w.Button{Icon: "dots", Title: "More"},
								More:    true,
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuItem{Label: "A–Z", On: true, Href: "/app/contacts?sort=az"},
										w.MenuItem{Label: "Recently changed", Href: "/app/contacts?sort=changed"},
										w.MenuSep{},
										w.CheckForm{
											Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}},
											Action: "/app/contacts",
											Apply:  w.Button{Label: "Apply", Small: true, Submit: true},
										},
									},
									Pop:    w.MenuUnderEnd,
									Inline: true,
								},
							},
						},
						Right: &w.BarRight{
							Items: []w.BarPart{
								w.CheckForm{
									Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}},
									Action: "/app/contacts",
									Apply:  w.Button{Label: "Apply", Small: true, Submit: true},
								},
								w.Mono{Text: "5"},
							},
							Secondary: true,
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#4A6B52"},
								Avatar:   w.Avatar{Initials: "АК"},
								Title:    []w.Inline{w.Text("Анна Ковалёва")},
								Strong:   true,
								Subtitle: "+7 912 445-10-88",
								Detail:   "a.kovaleva@example.org",
								Caption:  "Издательство «Слово»",
								Source:   "nextcloud · Контакты",
								On:       true,
							},
							w.FaceRow{
								Mark:       w.Mark{Colour: "#6B4E7A"},
								Avatar:     w.Avatar{Photo: true},
								Title:      []w.Inline{w.Text("Артём Белов ")},
								TitleBadge: &w.Badge{Label: "linked ×2", Tone: w.BadgeLinked},
								Subtitle:   "+7 903 118-22-04",
								Detail:     "a.belov@example.org",
								Caption:    "Тензор",
								Source:     "radicale · vps · Команда",
							},
							w.FaceRow{
								Mark:     w.Mark{Colour: "#4A6B52"},
								Avatar:   w.Avatar{Initials: "АС"},
								Title:    []w.Inline{w.Text("Алексей Смирнов")},
								Subtitle: "+7 916 200-71-15",
								Detail:   "—",
								Caption:  "—",
								Source:   "nextcloud · Контакты",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "В", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#4A6B52"},
								Avatar:   w.Avatar{Initials: "ВГ"},
								Title:    []w.Inline{w.Text("Виктор Гончаров")},
								Subtitle: "—",
								Detail:   "v.goncharov@example.org",
								Caption:  "Городская библиотека №4",
								Source:   "nextcloud · Контакты",
							},
							w.FaceRow{
								Mark:     w.Mark{Colour: "#A05A2C"},
								Avatar:   w.Avatar{Initials: "ВЛ"},
								Title:    []w.Inline{w.Text("Вера Лебедева")},
								Subtitle: "+7 905 331-09-42",
								Detail:   "vera@example.net",
								Caption:  "—",
								Source:   "nextcloud · Архив",
							},
						},
					},
				},
			},
			Height: w.ShellHeight560,
		},
		Scrim: w.Scrim{
			Content: w.Sheet{
				Icon:  "move",
				Title: "Move contact",
				Head: []w.SheetHeadPart{
					w.Dim{
						Parts:   []w.Inline{w.Text("Анна Ковалёва · from nextcloud · Контакты")},
						Size:    w.Dim11_5,
						Variant: w.DimWrap,
						Place:   w.Pushed,
					},
				},
				Body: []w.SheetPart{
					w.Rubric{Text: "Write into"},
					w.Tree{
						Rows: []w.TreeRow{
							{
								Mark:        &w.Mark{Colour: "#4A6B52"},
								Label:       "Контакты",
								Unavailable: true,
								Trail:       w.Mono{Text: "already there", Variant: w.Mono10Dusty},
							},
							{
								Mark:  &w.Mark{Colour: "#A05A2C"},
								Label: "Архив",
								Trail: w.Mono{Text: "nextcloud", Variant: w.Mono10Dusty},
							},
							{
								Mark:  &w.Mark{Colour: "#6B4E7A"},
								Label: "Команда",
								Trail: w.Mono{Text: "radicale", Variant: w.Mono10Dusty},
							},
							{
								Mark:        &w.Mark{Colour: "#8C6239"},
								Label:       "Shared",
								Unavailable: true,
								Trail:       w.Badge{Label: "read-only"},
							},
						},
						On: 2,
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("Inside one server this is "),
							w.Mono{Text: "MOVE", Form: w.MonoCode},
							w.Text(": instant, no traffic. To another address book the card is copied and then deleted — Carrel says so before it starts, not after."),
						},
						Space: w.MsgFlushTop2,
					},
				},
				Foot: []w.SheetFootPart{
					w.Button{Label: "Move", Kind: w.ButtonPrimary},
					w.Kbd{Keys: "Enter"},
					w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}},
				},
				Width: w.SheetWidth520,
			},
			Pad: w.ScrimPad26,
		},
	}, "s-contacts-2", "tc-overlay", 1),
	overlay("top", w.Overlay{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.OverlayBody{
			Rail: []w.OverlayRailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 5}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 2,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedPart,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined calendar",
							Meta:    "6",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "2"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "1"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "baikal"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Error: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Subscriptions", Items: []w.RubricPart{w.Badge{Label: "external"}}},
						w.TickRow{
							Kind:  w.TickRowExternal,
							Mark:  w.Mark{Colour: "#8C6239"},
							Label: "Праздники РФ",
							Meta:  "0",
							Trail: w.Badge{Label: "ro"},
						},
						w.TickRow{
							Ticked: w.TickedOff,
							Kind:   w.TickRowExternal,
							Mark:   w.Mark{Colour: "#6B4E7A"},
							Label:  "Расписание лиги",
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "релиз"}, w.Tag{Label: "работа"}},
				},
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined calendar"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{
									w.Text("5 of 7 calendars · 3 accounts · 10 – 16 August 2026 · "),
									w.Mono{Text: "Europe/Moscow"},
									w.Text(" · 6 events"),
								},
							},
						},
						Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}}},
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "10:00",
								Title:  []w.Inline{w.Text("Синк-ап по релизу")},
								Detail: "Переговорная 2",
								Badges: []w.Badge{{Label: "4 attendees"}},
								Source: "nextcloud · Личный",
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#3F6E8C"},
								Time:   "all day",
								Title:  []w.Inline{w.Text("Дежурство: первая линия")},
								Detail: "—",
								Badges: []w.Badge{{Label: "repeats weekly"}},
								Source: "nextcloud · Работа",
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#8C6239"},
								Time:   "all day",
								Title:  []w.Inline{w.Text("Дежурство: первая линия")},
								Detail: "—",
								Badges: []w.Badge{{Label: "repeats weekly"}},
								Source: "radicale · vps · Дежурства",
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#3F6E8C"},
								Time:   "15:00",
								Title:  []w.Inline{w.Text("Созвон с подрядчиком")},
								Detail: "Jitsi",
								Badges: []w.Badge{{Label: "tentative", Tone: w.BadgeAlert}},
								Source: "nextcloud · Работа",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Tuesday · 11 August", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#6B4E7A"},
								Time:   "09:30",
								Title:  []w.Inline{w.Text("Ревью архитектуры")},
								Detail: "Переговорная 1",
								Badges: []w.Badge{{Label: "2 attendees"}},
								Source: "radicale · vps · Проекты",
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "18:00",
								Title:  []w.Inline{w.Text("Забрать документы")},
								Detail: "МФЦ, Мясницкая 13",
								Source: "nextcloud · Личный",
							},
						},
					},
				},
			},
			Height: w.ShellHeight340,
		},
		Scrim: w.Scrim{
			Content: w.Sheet{
				Icon:  "search",
				Title: "Keyboard shortcuts",
				Head:  []w.SheetHeadPart{w.Close{Label: "Close"}},
				Body: []w.SheetPart{
					w.SheetKeys{
						Items: []w.SheetKey{
							{Keys: "N", Does: "Quick note over this screen"},
							{Keys: "/", Does: "Focus the search field in the header"},
							{Keys: "Enter", Does: "Open the selected record full screen"},
							{Keys: "Esc", Does: "Close a sheet, the details panel or focus mode"},
							{Keys: "Ctrl+S", Does: "Save in the editor"},
							{Keys: "Ctrl+Enter", Does: "Save a quick note"},
							{Keys: "Ctrl+V", Does: "Paste an image as an attachment"},
						},
					},
					w.SheetHint{Text: "Shortcuts do nothing while the cursor is in a field. None of them destroy anything."},
				},
				Foot:  []w.SheetFootPart{w.Button{Label: "Close"}, w.Kbd{Keys: "Esc"}, w.Kbd{Keys: "?", Place: w.Pushed}},
				Width: w.SheetWidth480,
			},
		},
	}, "account-6", "tc-overlay", 1),
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

	// tc-dialog
	overlay("wide", w.Dialog{Items: []w.DialogPart{
		w.Steps{Items: []string{"1 · Archive", "2 · Match", "3 · Report"}, On: 2},
		w.Head{Title: &w.HeadTitle{Text: "Restore from 2026-08-14"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Job «Личное на домашний диск» · manifest read · 4 collections · 1 214 records · "), w.Mono{Text: "carrel-backup/2026-08-14T09-12.zip"}}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Restore 3 collections", Icon: "restore", Kind: w.ButtonPrimary, Name: "op", Value: "restore", Submit: true}}}},
		w.Field{Label: "5 records of the archive are already here and differ", Parts: []w.FieldPart{w.Seg{Options: []w.SegOption{{Label: "Skip them", Href: "/app/settings/backup/restore?existing=skip", On: true}, {Label: "Replace them", Href: "/app/settings/backup/restore?existing=replace"}, {Label: "Import as copies", Href: "/app/settings/backup/restore?existing=copy"}}}, w.FieldHint{Parts: []w.Inline{w.Text("They stay as they are here; the archive's version is not written.")}, Variant: w.HintColumn}}, OwnLine: true},
		w.Bar{Items: []w.BarPart{w.BarSep{}, w.Button{Label: "Restore into", Icon: "restore", Name: "op", Value: "into", Submit: true}, w.Drop{Trigger: w.Pick{Value: "nextcloud · Личный", Mark: &w.Mark{Colour: "#4A6B52"}}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "radicale · vps · Проекты", Name: "into", Value: "radicale|proekty"},
			w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Личный", On: true, Name: "into", Value: "nextcloud|lichnyj"},
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа", Name: "into", Value: "nextcloud|rabota"},
			w.MenuSep{},
			w.MenuItem{Label: "New collection", Name: "into", Value: "new"},
			w.MenuItem{Label: "Skip this one", Tone: w.MenuToneDusty, Name: "into", Value: "skip"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}, w.Button{Label: "Skip", Name: "op", Value: "skip", Submit: true}}, Count: &w.BarCount{Text: "3 selected"}, Right: &w.BarRight{Items: []w.BarPart{w.Button{Folds: true, Label: "Select all", Icon: "selectall", Small: true, Title: "Select all", ScriptOnly: true}, w.Button{Folds: true, Label: "Clear selection", Icon: "close", Small: true, Title: "Clear selection"}}}, Variant: w.BarSelection},
		w.Table{Cols: []w.TableCol{{Width: w.ColTick, Tick: ptr(w.TickedPart)}, {Label: "In the archive"}, {Label: "Records", Rank: w.Rank3, Secondary: true}, {Label: "Restore into", Width: w.Col26, Rank: w.Rank4, Main: true}, {Label: "What will happen"}}, Rows: []w.TableRow{
			{Cells: []w.TableCell{{Pick: ptr(w.TickedOn), Send: w.Send{Name: "collection", Value: "nextcloud/personal"}}, {Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "nextcloud / personal"}}}, {Parts: []w.CellPart{w.Text("312")}, Mono: true}, {Parts: []w.CellPart{w.Drop{Trigger: w.Pick{Value: "nextcloud · Личный", Mark: &w.Mark{Colour: "#4A6B52"}}, Menu: w.Menu{Items: []w.MenuPart{
				w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "radicale · vps · Проекты", Name: "into.1", Value: "radicale|proekty"},
				w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Личный", On: true, Name: "into.1", Value: "nextcloud|lichnyj"},
				w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа", Name: "into.1", Value: "nextcloud|rabota"},
				w.MenuSep{},
				w.MenuItem{Label: "New collection", Name: "into.1", Value: "new"},
				w.MenuItem{Label: "Skip this one", Tone: w.MenuToneDusty, Name: "into.1", Value: "skip"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}}}, {Parts: []w.CellPart{w.Text("matched by path — 4 records missing, will be created · 3 already here")}}}, State: w.RowPicked},
			{Cells: []w.TableCell{{Pick: ptr(w.TickedOn), Send: w.Send{Name: "collection", Value: "radicale/projects"}}, {Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "radicale / projects"}}}, {Parts: []w.CellPart{w.Text("86")}, Mono: true}, {Parts: []w.CellPart{w.Drop{Trigger: w.Pick{Value: "radicale · vps · Проекты", Mark: &w.Mark{Colour: "#6B4E7A"}}, Menu: w.Menu{Items: []w.MenuPart{
				w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "radicale · vps · Проекты", On: true, Name: "into.2", Value: "radicale|proekty"},
				w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Личный", Name: "into.2", Value: "nextcloud|lichnyj"},
				w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа", Name: "into.2", Value: "nextcloud|rabota"},
				w.MenuSep{},
				w.MenuItem{Label: "New collection", Name: "into.2", Value: "new"},
				w.MenuItem{Label: "Skip this one", Tone: w.MenuToneDusty, Name: "into.2", Value: "skip"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}}}, {Parts: []w.CellPart{w.Text("matched by path — no differences")}}}, State: w.RowPicked},
			{Cells: []w.TableCell{{Pick: ptr(w.TickedOn), Send: w.Send{Name: "collection", Value: "baikal/shared"}}, {Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#9A9280"}, Name: "baikal / shared"}}}, {Parts: []w.CellPart{w.Text("140")}, Mono: true}, {Parts: []w.CellPart{w.Drop{Trigger: w.Pick{Value: "nextcloud · Личный", Mark: &w.Mark{Colour: "#4A6B52"}}, Menu: w.Menu{Items: []w.MenuPart{
				w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "radicale · vps · Проекты", Name: "into.3", Value: "radicale|proekty"},
				w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Личный", On: true, Name: "into.3", Value: "nextcloud|lichnyj"},
				w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа", Name: "into.3", Value: "nextcloud|rabota"},
				w.MenuSep{},
				w.MenuItem{Label: "New collection", Name: "into.3", Value: "new"},
				w.MenuItem{Label: "Skip this one", Tone: w.MenuToneDusty, Name: "into.3", Value: "skip"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}}}, {Parts: []w.CellPart{w.Span{Face: w.FaceDanger, Text: "that server is gone — another collection chosen"}, w.Text(" · 2 already here")}}}, State: w.RowPicked},
			{Cells: []w.TableCell{{Pick: ptr(w.TickedOff), Send: w.Send{Name: "collection", Value: "nextcloud/birthdays"}}, {Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "nextcloud / birthdays"}}}, {Parts: []w.CellPart{w.Text("51")}, Mono: true}, {Parts: []w.CellPart{w.Text("—")}, Mono: true, Unset: true, Tone: w.CellDusty}, {Parts: []w.CellPart{w.Text("skip")}, Dim: true}}}}, Variant: w.TableFit},
		w.Msg{Parts: []w.MsgPart{w.Text("Restoring is an import: records made since the archive stay, nothing is deleted.")}, Tone: w.MsgAlert, Space: w.MsgFlushSides, Column: true}}, Variant: w.DialogWide, Action: "/app/settings/backup/restore/match"}, "s-backup-8", "tc-dialog", 1),
	overlay("step", w.Dialog{Items: []w.DialogPart{
		w.Steps{Items: []string{"1 · Archive", "2 · Match", "3 · Report"}, On: 1},
		w.Head{Title: &w.HeadTitle{Text: "Restore"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Nothing is written until you have seen what goes where")}}}},
		w.Field{Label: "Where the archive is", Parts: []w.FieldPart{w.Seg{Options: []w.SegOption{{Label: "From a job", Href: "/app/settings/backup/restore?from=job"}, {Label: "From a server", Href: "/app/settings/backup/restore?from=server"}, {Label: "From this computer", Href: "/app/settings/backup/restore?from=file", On: true}}, Variant: w.SegStart}, w.FieldHint{Parts: []w.Inline{w.Text("One archive, one of three. The other two stay empty until you switch.")}}}, OwnLine: true},
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "From this computer"},
			w.CardFields{Items: []w.CardFieldsPart{w.Field{Label: "Archive", Parts: []w.FieldPart{w.Line{Items: []w.LinePart{w.Button{Label: "Choose a file", File: &w.FileField{Form: "", Name: "file", Accept: ".zip", Label: "Archive"}}}, Gap: w.LineGap8}}, Width: w.FieldLG}}, Top: true},
			w.CardFields{Items: []w.CardFieldsPart{w.Field{Label: "Storage for a large archive", Parts: []w.FieldPart{w.Pick{Name: "store", Options: []w.PickOption{{Value: "", Label: "None — up to 16 MB"}, {Value: "nextcloud|files", Label: "nextcloud · Files"}, {Value: "sftpgo|documents", Label: "sftpgo · vps · Documents", On: true}}, Label: "Storage for a large archive"}}, Width: w.FieldMD}, w.Field{Label: "Folder in it", Parts: []w.FieldPart{w.Input{Name: "folder", Value: "/carrel/restore", Label: "Folder in it", Mono: true}, w.FieldHint{Parts: []w.Inline{w.Text("An archive larger than 16 MB is put into this folder and read from there: Carrel keeps no copy of it.")}}}, Width: w.FieldLG}}, Top: true}}},
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Archive password"},
			w.CardFields{Items: []w.CardFieldsPart{w.Field{Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "password", Value: "••••••••••••", Label: "Archive password", Focus: true}}, Width: w.FieldLG}, w.Span{Face: w.FaceMuted12, Text: "The one set in the job, not your sign-in password. Without it the manifest cannot be read — Carrel will not even see the list of collections."}}}}},
		w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Open the manifest", Kind: w.ButtonPrimary, Submit: true}, w.Button{Label: "Cancel"}}}}, Variant: w.DialogStep760, Action: "/app/settings/backup/restore?from=file", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}}, "s-backup-15", "tc-dialog", 1),
	overlay("width", w.Dialog{Items: []w.DialogPart{
		w.Head{Title: &w.HeadTitle{Text: "Key escrow"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("A copy of your key on this server")}}}},
		w.Msg{Parts: []w.MsgPart{w.Text("With the escrow master password, an administrator can decrypt the DAV credentials you save here, and through them your records.")}, Tone: w.MsgAlert, Space: w.MsgFlushSides, Column: true},
		w.Fields{Items: []w.FieldsItem{
			{Term: "Deposited", Value: []w.ValuePart{w.Text("2026-03-02")}, Face: w.ValueMono},
			{Term: "Last recovery", Value: []w.ValuePart{w.Text("never")}, Face: w.ValueMono}}, Variant: w.FieldsBleed},
		w.Field{Label: "Password history", Parts: []w.FieldPart{w.Line{Items: []w.LinePart{w.Input{Value: "https://carrel.example/password-history/k7q2-m9xd-41fa", Label: "Password history", Mono: true, ReadOnly: true, Variant: w.InputFill}, w.Button{Label: "Copy", Icon: "dup", Small: true, Copy: "https://carrel.example/password-history/k7q2-m9xd-41fa", Copied: "Copied"}}, Gap: w.LineGap6}, w.FieldHint{Parts: []w.Inline{w.Text("Keep it: it shows what was done to your password, even when you cannot sign in.")}, Variant: w.HintColumn}}, Width: w.FieldFull},
		w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Withdraw my deposited key", Kind: w.ButtonDangerSolid}, w.Dim{Parts: []w.Inline{w.Text("Then nobody can recover the account.")}, Size: w.Dim11_5, Variant: w.DimWrap}}, Variant: w.FormFootFlush}}, Variant: w.Dialog560}, "account-4", "tc-dialog", 1),
	overlay("width", w.Dialog{Items: []w.DialogPart{
		w.Head{Title: &w.HeadTitle{Text: "Add device"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("An app password for one device, with its own key wrapper")}}}},
		w.Form{Items: []w.FormPart{w.Field{Label: "Name", Parts: []w.FieldPart{w.Input{Name: "name", Value: "Телефон · DAVx5", Label: "Name", Focus: true}, w.FieldHint{Parts: []w.Inline{w.Text("For you: this is how you will know what you are revoking later.")}}}, Width: w.FieldW100}, w.Field{Label: "What this device may do", Parts: []w.FieldPart{w.Seg{Options: []w.SegOption{{Label: "Read only", Value: "read.only"}, {Label: "Read and write", On: true, Value: "read.and.write"}}, Variant: w.SegStart, Name: "what.this.device.may.do"}, w.FieldHint{Parts: []w.Inline{w.Text("Read only: the device syncs, and whatever it sends back is refused. Changed later in the device’s row, without a new password.")}}}, Width: w.FieldW100}, w.Field{Label: "What the device will see", LabelAside: w.Dim{Parts: []w.Inline{w.Text("— changed later too, without re-issuing the password")}}, Parts: []w.FieldPart{w.Group{Items: []w.FieldPart{w.Check{On: true, Name: "calendars", Value: "1", Parts: []w.CheckPart{w.Text("Calendars")}}, w.Check{On: true, Name: "address.books", Value: "1", Parts: []w.CheckPart{w.Text("Address books")}}, w.Check{Name: "tasks.vtodo.not.every.client.understands.them", Value: "1", Parts: []w.CheckPart{w.Text("Tasks "), w.Dim{Parts: []w.Inline{w.Text("· VTODO — not every client understands them")}}}}, w.Check{Name: "notes.vjournal.a.different.subset.of.clients", Value: "1", Parts: []w.CheckPart{w.Text("Notes "), w.Dim{Parts: []w.Inline{w.Text("· VJOURNAL — a different subset of clients")}}}}, w.Check{On: true, Name: "webdav.folders.mount.as.a.network.drive", Value: "1", Parts: []w.CheckPart{w.Text("WebDAV folders "), w.Dim{Parts: []w.Inline{w.Text("· mount as a network drive")}}}}}, Form: w.GroupColumn}}, Width: w.FieldW100}, w.Field{Label: "Collections", Parts: []w.FieldPart{w.Group{Items: []w.FieldPart{w.Check{On: true, Name: "nextcloud", Value: "1", Parts: []w.CheckPart{w.Text("nextcloud · Личный")}}, w.Check{On: true, Name: "nextcloud", Value: "1", Parts: []w.CheckPart{w.Text("nextcloud · Работа")}}, w.Check{On: true, Name: "nextcloud", Value: "1", Parts: []w.CheckPart{w.Text("nextcloud · Контакты")}}, w.Check{Name: "radicale.vps", Value: "1", Parts: []w.CheckPart{w.Text("radicale · vps · Проекты")}}, w.Check{Name: "radicale.vps", Value: "1", Parts: []w.CheckPart{w.Text("radicale · vps · Команда")}}, w.Check{On: true, Name: "sftpgo.vps.documents", Value: "1", Parts: []w.CheckPart{w.Text("sftpgo · vps · Documents")}}}, Form: w.GroupColumn}}, Width: w.FieldW100}}, ID: "device", Action: "/app/settings/davloom/devices/new", Variant: w.FormFlush},
		w.Msg{Parts: []w.MsgPart{w.Text("The password is shown once and stored nowhere. Lose it and you create a new one and revoke this — it cannot be recovered.")}, Tone: w.MsgAlert, Space: w.MsgFlushSides},
		w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Create password", Kind: w.ButtonPrimary, Form: "device"}, w.Button{Label: "Cancel"}}, Variant: w.FormFootFlush}}, Variant: w.Dialog520}, "s-davloom-2", "tc-dialog", 1),
	overlay("page", w.Dialog{Items: []w.DialogPart{
		w.Steps{Items: []string{"1 · Archive", "2 · Match", "3 · Report"}, On: 1},
		w.Head{Title: &w.HeadTitle{Text: "Restore"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Nothing is written until you have seen what goes where")}}}},
		w.Field{Label: "Where the archive is", Parts: []w.FieldPart{w.Seg{Options: []w.SegOption{{Label: "From a job", Href: "/app/settings/backup/restore?from=job", On: true}, {Label: "From a server", Href: "/app/settings/backup/restore?from=server"}, {Label: "From this computer", Href: "/app/settings/backup/restore?from=file"}}, Variant: w.SegStart}, w.FieldHint{Parts: []w.Inline{w.Text("One archive, one of three. The other two stay empty until you switch.")}}}, OwnLine: true},
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "From a job"},
			w.Table{Rows: []w.TableRow{
				{Cells: []w.TableCell{{Pick: ptr(w.TickedOn), Send: w.Send{Name: "archive", Value: "lichnoe-na-domashnij-disk", One: true}, Width: w.ColTick}, {Parts: []w.CellPart{w.Text("Личное на домашний диск")}}, {Parts: []w.CellPart{w.Text("2026-08-14 09:12")}, Mono: true, Rank: w.Rank3}, {Parts: []w.CellPart{w.Text("2.1 GB")}, Mono: true, Rank: w.Rank4}, {Parts: []w.CellPart{w.Text("4 collections")}, Mono: true}}, State: w.RowPicked},
				{Cells: []w.TableCell{{Pick: ptr(w.TickedOff), Send: w.Send{Name: "archive", Value: "rabochee-ezhednevno", One: true}}, {Parts: []w.CellPart{w.Text("Рабочее, ежедневно")}}, {Parts: []w.CellPart{w.Text("2026-08-16 03:10")}, Mono: true, Rank: w.Rank3}, {Parts: []w.CellPart{w.Text("18.4 MB")}, Mono: true, Rank: w.Rank4}, {Parts: []w.CellPart{w.Text("3 collections")}, Mono: true}}}}}}},
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Archive password"},
			w.CardFields{Items: []w.CardFieldsPart{w.Field{Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "password", Value: "••••••••••••", Label: "Archive password", Focus: true}}, Width: w.FieldLG}, w.Span{Face: w.FaceMuted12, Text: "The one set in the job, not your sign-in password. Without it the manifest cannot be read — Carrel will not even see the list of collections."}}}}},
		w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Open the manifest", Kind: w.ButtonPrimary, Submit: true}, w.Button{Label: "Cancel"}}}}, Variant: w.DialogStep760, Action: "/app/settings/backup/restore"}, "s-backup-7", "tc-dialog", 1),
	overlay("links", w.Dialog{Items: []w.DialogPart{
		w.AuthStack{Items: []w.AuthStackPart{w.Form{Items: []w.FormPart{w.Field{Label: "Login", Parts: []w.FieldPart{w.Input{Name: "login", Value: "mix", Label: "Login"}}, Width: w.FieldMD}}, ID: "signin", Action: "/login", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}, Variant: w.FormFlushAll}, w.Button{Label: "Sign in", Kind: w.ButtonPrimary, Form: "signin"}}, Links: []w.Link{{Text: "Forgotten your password?", Href: "/forgot"}, {Text: "Create an account", Href: "/register"}}}}, Variant: w.DialogPage440}, "components-43", "tc-dialog", 1),
	overlay("crop-image", w.Dialog{Items: []w.DialogPart{
		w.PhotoCrop{Colour: "#8C7B6A", Src: "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 40 40'%3E%3Crect width='40' height='40' fill='%238C7B6A'/%3E%3Ccircle cx='20' cy='16' r='7' fill='%23E8DFCB'/%3E%3Cpath d='M6 40c2-9 8-13 14-13s12 4 14 13z' fill='%23E8DFCB'/%3E%3C/svg%3E", Note: "Drag the photo to choose the frame. The box is what other clients will see.", Fields: []w.Field{}}}, Variant: w.Dialog520}, "components-40", "tc-dialog", 1),
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
		w.Head{Title: &w.HeadTitle{Text: "Import contacts"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mono{Text: "Разное (mix).vcf"}, w.Text(" · 84 KB · vCard 3.0")}}}, Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Import 117 contacts", Kind: w.ButtonPrimary}}}},
		w.Form{Items: []w.FormPart{w.Field{Label: "Into", Parts: []w.FieldPart{w.Drop{Trigger: w.Pick{Value: "nextcloud · Контакты", Mark: &w.Mark{Colour: "#4A6B52"}}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Контакты", On: true, Href: "/app/contacts/nextcloud/kontakty/import"},
			w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "nextcloud · Архив", Href: "/app/contacts/nextcloud/arhiv/import"},
			w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "radicale · vps · Команда", Reason: "Read-only: nothing can be written here."},
			w.MenuSep{},
			w.MenuItem{Label: "New address book", Href: "/app/collections/new?kind=contact&back=/app/contacts/nextcloud/kontakty/import"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}, w.FieldHint{Parts: []w.Inline{w.Text("Any address book you may write to. What is already here is counted against it.")}}}, Width: w.FieldLG}, w.Field{Label: "5 records of this file are already in Контакты", Parts: []w.FieldPart{w.Seg{Options: []w.SegOption{{Label: "Skip them", Href: "/app/contacts/nextcloud/kontakty/import?existing=skip", On: true}, {Label: "Replace them", Href: "/app/contacts/nextcloud/kontakty/import?existing=replace"}, {Label: "Import as copies", Href: "/app/contacts/nextcloud/kontakty/import?existing=copy"}}}, w.FieldHint{Parts: []w.Inline{w.Text("They stay as they are here; the file's version is not written.")}}}, OwnLine: true}}, Tight: true},
		w.Stats{Items: []w.Stat{w.Stat{Value: "117", Label: "ready"}, w.Stat{Value: "3", Label: "with errors", Error: true}, w.Stat{Value: "5", Label: "already here"}}},
		w.Table{Cols: []w.TableCol{{Label: "Source"}, {Label: "Name"}, {Label: "UID", Rank: w.Rank3}, {Label: "Result", Rank: w.Rank4}}, Rows: []w.TableRow{
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.Text("Разное (mix).vcf")}, Mono: true}, {Parts: []w.CellPart{w.Text("Анна Ковалёва")}}, {Parts: []w.CellPart{w.Text("4f2b-9c31-a0e7")}, Mono: true}, {Parts: []w.CellPart{w.Text("Ready")}}}},
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.Text("Разное (mix).vcf")}, Mono: true}, {Parts: []w.CellPart{w.Text("Вера Лебедева")}}, {Parts: []w.CellPart{w.Text("8e20-5b1f-c7d4")}, Mono: true}, {Parts: []w.CellPart{w.Text("Ready — name from N")}}}},
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.Text("Разное (mix).vcf")}, Mono: true}, {Parts: []w.CellPart{w.Text("Артём Белов")}}, {Parts: []w.CellPart{w.Text("1a77-40de-bb02")}, Mono: true}, {Parts: []w.CellPart{w.Text("Already here — skipped")}}}},
			{Cells: []w.TableCell{{Parts: []w.CellPart{w.Text("Разное (mix).vcf")}, Mono: true}, {Parts: []w.CellPart{w.Text("—")}, Dim: true}, {Parts: []w.CellPart{w.Text("—")}, Mono: true, Tone: w.CellDusty}, {Parts: []w.CellPart{w.Text("line 12: cannot decode quoted-printable")}, Tone: w.CellDanger}}}}}}, Variant: w.DialogWide}, "s-io-5", "tc-dialog", 1),
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
			w.CardText{Kind: w.CardNum, Parts: []w.CardTextPart{w.Text("117")}}}},
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Already here, skipped", Items: []w.RubricPart{w.Mono{Text: "5", Variant: w.MonoTight}}},
			w.Diag{Text: "Артём Белов · 1a77-40de-bb02\nИрина Соколова · 33b0-9ffa-1120"}}},
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "With a new UID", Items: []w.RubricPart{w.Mono{Text: "2", Variant: w.MonoTight}}},
			w.Diag{Text: "Вера Дорошенко · 2b7e-11aa → 9d04-77c1 · imported as a copy\nВиктор Панов · 51c0-3a8e → baikal-0412 · the server gave its own"}}},
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Not kept by the server", Items: []w.RubricPart{w.Mono{Text: "1", Variant: w.MonoTight}}},
			w.Diag{Text: "Галина Мещерякова · PHOTO"}}},
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Failed", Items: []w.RubricPart{w.Mono{Text: "3", Variant: w.MonoTight}}},
			w.Diag{Text: "line 12: cannot decode quoted-printable", Variant: w.DiagDanger}}},
		w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Back to list", Kind: w.ButtonPrimary}}, Variant: w.FormFootFlush}}, Variant: w.DialogStep520}, "s-io-9", "tc-dialog", 1),
	overlay("width", w.Dialog{Items: []w.DialogPart{
		w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 1},
		w.Head{Title: &w.HeadTitle{Text: "Import contacts"}},
		w.Form{Items: []w.FormPart{w.Field{Label: "Into", Parts: []w.FieldPart{w.Drop{Trigger: w.Pick{Value: "nextcloud · Контакты", Mark: &w.Mark{Colour: "#4A6B52"}}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Контакты", On: true, Href: "/app/contacts/nextcloud/kontakty/import"},
			w.MenuItem{Mark: &w.Mark{Colour: "#A05A2C"}, Label: "nextcloud · Архив", Href: "/app/contacts/nextcloud/arhiv/import"},
			w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "radicale · vps · Команда", Reason: "Read-only: nothing can be written here."},
			w.MenuSep{},
			w.MenuItem{Label: "New address book", Href: "/app/collections/new?kind=contact&back=/app/contacts/nextcloud/kontakty/import"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}, w.FieldHint{Parts: []w.Inline{w.Text("Any address book you may write to, on every connected server.")}}}, Width: w.FieldW100}}, Variant: w.FormBottom12, Tight: true},
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "From a file"},
			w.Field{Label: "vCard file", Parts: []w.FieldPart{w.Line{Items: []w.LinePart{w.Input{Placeholder: "Разное (mix).vcf", Label: "Разное (mix).vcf", Picked: true, ID: "import-file", Variant: w.InputFill}, w.Button{Label: "Browse", File: &w.FileField{Form: "", Name: "file", Accept: ".vcf", Label: "vCard file", Shows: "import-file"}}, w.Button{Label: "Preview file", Kind: w.ButtonPrimary, Submit: true}}, Gap: w.LineGap8}}, Width: w.FieldW100},
			w.CardHint{Parts: []w.CardTextPart{w.Mono{Text: ".vcf"}, w.Text(" — vCard 2.1, 3.0 and 4.0. One file holds any number of cards; the calendar and the list of tasks take "), w.Mono{Text: ".ics"}, w.Text(" here instead.")}, Line: true}}}}, Variant: w.DialogStep560, Action: "/app/contacts/nextcloud/kontakty/import"}, "s-io-3", "tc-dialog", 1),
	overlay("narrow", w.Dialog{Items: []w.DialogPart{
		w.DialogLogo{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo48}, Wide: true},
		w.DialogTitle{Text: "Check your email", Gap: w.DialogTitleGap6},
		w.Prose{Parts: []w.Inline{w.Text("The letter went to "), w.Mono{Text: "v@example.org"}, w.Text(". The account cannot be used until you open the link in it. Sending again with the same login and address, while the account is unconfirmed, simply replaces the letter and the password.")}, Variant: w.ProseMuted},
		w.FormFoot{Variant: w.FormFootFlush, Items: []w.FormFootPart{w.Button{Label: "Back to sign in"}}}}, Variant: w.DialogNarrow}, "s-auth-3", "tc-dialog", 1),
	overlay("width", w.Dialog{Items: []w.DialogPart{
		w.Head{Title: &w.HeadTitle{Text: "Conflict"}, Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Another client changed this contact while you were editing it. Nothing has been overwritten.")}}}},
		w.DiffAgree{Property: w.Mono{Text: "FN"}, Text: "agrees — kept as it is."},
		w.Diff{Items: []w.DiffItem{
			{Title: w.Rubric{Text: "", Lead: &w.Mono{Text: "TEL;work"}, Items: []w.RubricPart{w.RubricAside{Text: "only in your edit"}}}, Options: []w.Check{w.Check{Name: "choice.0", Value: "mine", Radio: true, On: true, Parts: []w.CheckPart{w.Choice{Source: "Your edit · include", Value: "+7 495 120-45-00"}}}, w.Check{Name: "choice.0", Value: "theirs", Radio: true, Parts: []w.CheckPart{w.Choice{Source: "Leave it out", Value: "—"}}}}},
			{Title: w.Rubric{Text: "", Lead: &w.Mono{Text: "EMAIL;work"}, Items: []w.RubricPart{w.RubricAside{Text: "both changed"}}}, Options: []w.Check{w.Check{Name: "choice.1", Value: "mine", Radio: true, On: true, Parts: []w.CheckPart{w.Choice{Source: "Your edit", Value: "a.kovaleva@example.org"}}}, w.Check{Name: "choice.1", Value: "theirs", Radio: true, Parts: []w.CheckPart{w.Choice{Source: "Server", Value: "anna.kovaleva@slovo.example"}}}}},
			{Title: w.Rubric{Text: "", Lead: &w.Mono{Text: "NOTE"}, Items: []w.RubricPart{w.RubricAside{Text: "only in your edit"}}}, Options: []w.Check{w.Check{Name: "choice.2", Value: "mine", Radio: true, On: true, Parts: []w.CheckPart{w.Choice{Source: "Your edit · include", Value: "Пишет по вторникам, отвечает в тот же день."}}}, w.Check{Name: "choice.2", Value: "theirs", Radio: true, Parts: []w.CheckPart{w.Choice{Source: "Leave it out", Value: "—"}}}}},
			{Title: w.Rubric{Text: "", Lead: &w.Mono{Text: "X-EVOLUTION-FILE-AS"}, Items: []w.RubricPart{w.RubricAside{Text: "only on the server"}}}, Options: []w.Check{w.Check{Name: "choice.3", Value: "mine", Radio: true, Parts: []w.CheckPart{w.Choice{Source: "Your edit", Value: "—"}}}, w.Check{Name: "choice.3", Value: "theirs", Radio: true, On: true, Parts: []w.CheckPart{w.Choice{Source: "Server · include", Value: "Ковалёва, Анна"}}}}}}, ID: "conflict", Action: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/conflict", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "etag", Value: "3-a7c1"}, {Name: "op", Value: "combine"}}},
		w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Write this combination", Kind: w.ButtonPrimary, Form: "conflict"}, w.Button{Label: "Use only my edit", Form: "conflict.mine"}, w.Button{Label: "Use only the server version", Form: "conflict.theirs"}, w.Button{Label: "Cancel", Href: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7"}}, Variant: w.FormFootWrap},
		w.PostForm{ID: "conflict.mine", Action: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/conflict?op=mine", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "etag", Value: "3-a7c1"}}},
		w.PostForm{ID: "conflict.theirs", Action: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/conflict?op=theirs", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "etag", Value: "3-a7c1"}}}}}, "s-conflict-1", "tc-dialog", 1),
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
		w.PhotoCrop{Colour: "#8C7B6A", Src: "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 40 40'%3E%3Crect width='40' height='40' fill='%238C7B6A'/%3E%3Ccircle cx='20' cy='16' r='7' fill='%23E8DFCB'/%3E%3Cpath d='M6 40c2-9 8-13 14-13s12 4 14 13z' fill='%23E8DFCB'/%3E%3C/svg%3E", Note: "Drag to choose the frame.", Fields: []w.Field{w.Field{Label: "Zoom", Parts: []w.FieldPart{w.Zoom{Value: "1.6×"}}, Width: w.FieldW100}, w.Field{Label: "Rotate", Parts: []w.FieldPart{w.Seg{Options: []w.SegOption{{Label: "0°", On: true, Value: "0"}, {Label: "90°", Value: "90"}, {Label: "180°", Value: "180"}, {Label: "270°", Value: "270"}}, Variant: w.SegStart, Name: "rotate"}}, Width: w.FieldW100}, w.Field{Label: "Saved as", Parts: []w.FieldPart{w.Input{Value: "512 × 512 · JPEG · about 90 KB", Label: "Saved as", Mono: true, ReadOnly: true}, w.FieldHint{Parts: []w.Inline{w.Text("Kept inside the card, so it is made small.")}}}, Width: w.FieldW100}}},
		w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Save photo", Kind: w.ButtonPrimary, Submit: true}, w.Button{Label: "Cancel", Href: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/edit"}, w.Push{Items: []w.PushPart{w.Button{Label: "Remove photo", Kind: w.ButtonDanger, To: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/photo-preview?op=remove", Submit: true}}}}, Variant: w.FormFootFlush}}, Variant: w.Dialog520, Action: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/photo-preview?op=crop", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "etag", Value: "7f3a-91de"}, {Name: "zoom", Value: "1.6"}, {Name: "x", Value: "0.12"}, {Name: "y", Value: "0.08"}}}, "s-contacts-5", "tc-dialog", 1),
	overlay("narrow", w.Dialog{Items: []w.DialogPart{
		w.DialogLogo{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo60}},
		w.Tagline{Lead: "One desk. Every server.", Text: "Calendars, contacts, tasks, notes and files — one place.", Centered: true},
		w.AuthStack{Items: []w.AuthStackPart{w.Form{}, w.Button{Label: "Sign in", Kind: w.ButtonPrimary}}, Links: []w.Link{{Text: "Create an account"}, {Text: "Forgotten your password?"}}}}, Variant: w.DialogNarrow}, "s-auth-1", "tc-dialog", 1),
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
		w.AuthStack{Items: []w.AuthStackPart{w.Form{}, w.Button{Label: "Sign in", Kind: w.ButtonPrimary}}, Links: []w.Link{{Text: "Create an account"}, {Text: "Forgotten your password?"}}}}, Variant: w.DialogNarrow}}, "s-auth-1", "tc-auth", 1),
)

func init() {
	Widgets = append(Widgets, w.Menu{}, w.Drop{Open: true}, w.Sheet{}, w.Scrim{}, w.Overlay{}, w.Dialog{}, w.Auth{})
	Parts = append(Parts, w.MenuItem{}, w.MenuSep{}, w.MenuHeading{}, w.MenuField{}, w.MenuColumn{}, w.MenuFoot{}, w.TopDrop{Open: true},
		w.SheetSubject{}, w.SheetFields{}, w.SheetKey{}, w.SheetKeys{}, w.SheetHint{}, w.SheetNote{}, w.SheetLine{}, w.SheetBlock{},
		w.SheetChecks{}, w.SheetFrom{}, w.SheetPick{}, w.DialogTitle{}, w.DialogSub{}, w.DialogLogo{}, w.Tagline{}, w.AuthStack{},
		w.Lede{}, w.PhotoCrop{}, w.OverlayBody{}, w.DiffAgree{})
}
