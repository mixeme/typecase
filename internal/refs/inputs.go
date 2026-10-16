package refs

import w "carrel.dev/typecase/widget"

func form(state string, fixture Renderer, frame, class string, n int) Entry {
	return e("forms", state, fixture, ref(frame, class, n))
}

var _ = add(
	// tc-in as a line of input
	form("password", w.Input{Kind: w.InputPassword, Name: "password", Value: "••••••••••••", Label: "Archive password", Focus: true}, "s-backup-7", "tc-in", 1),
	form("disabled", w.Input{Kind: w.InputTime, Placeholder: "--:--", Label: "Start time", Mono: true, Disabled: true}, "s-tasks-3", "tc-in", 5),
	form("placeholder", w.Input{Kind: w.InputTime, Placeholder: "--:--", Label: "Start time", Mono: true}, "s-tasks-14", "tc-in", 5),
	form("focus-empty", w.Input{Label: "Display name", Focus: true}, "s-contacts-27", "tc-in", 2),
	form("error-empty", w.Input{Label: "Password", Focus: true, Error: true}, "s-auth-13", "tc-in", 2),
	form("template-row", w.Input{Name: "tel.{n}", Label: "Number", Mono: true}, "components-38", "tc-in", 5),
	form("error", w.Input{Value: "anna", Label: "Login", Error: true}, "s-admin-19", "tc-in", 1),
	form("value", w.Input{Value: "anna@example.org", Label: "Email"}, "s-auth-12", "tc-in", 2),
	form("error-mono", w.Input{Value: "9f2", Label: "Temporary password", Mono: true, Error: true}, "s-admin-19", "tc-in", 3),
	form("focus", w.Input{Value: "2026-07", Label: "Name of the new folder", Focus: true, Variant: w.InputFill}, "s-files-13", "tc-in", 2),
	form("error", w.Input{Value: "31.02.1987", Label: "Birthday", Mono: true, Focus: true, Error: true, Numeric: true}, "s-contacts-18", "tc-in", 1),
	form("value", w.Input{Value: "Анна Ковалёва", Label: "Display name"}, "components-10", "tc-in", 1),
	form("mono", w.Input{Value: "8765", Mono: true, Label: "Port"}, "components-13", "tc-in", 1),
	form("focus", w.Input{Value: "a.kovaleva@example.org", Focus: true, Label: "Email"}, "components-10", "tc-in", 4),
	form("mono", w.Input{Kind: w.InputDate, Value: "2026-01-01", Mono: true, Label: "From"}, "components-18", "tc-in", 1),
	form("placeholder", w.Input{Placeholder: "—", Label: "Nickname"}, "s-contacts-3", "tc-in", 5),
	form("mono", w.Input{Kind: w.InputTime, Value: "10:00", Mono: true, Label: "Start time"}, "components-10", "tc-in", 3),
	form("placeholder", w.Input{Placeholder: "DD.MM.YYYY", Mono: true, Numeric: true, Label: "Birthday"}, "components-10", "tc-in", 2),
	form("variant", w.Input{Value: "80", Mono: true, Variant: w.InputChip, Label: "Same UID"}, "s-dup-3", "tc-in", 2),
	form("variant", w.Input{Value: "z3-brick-lamp-9f2", Mono: true, Variant: w.InputFill, Label: "Password"}, "s-profile-4", "tc-in", 2),
	form("placeholder", w.Input{Placeholder: "Разное (mix).vcf", Variant: w.InputFill, Label: "File name"}, "s-io-3", "tc-in", 2),
	form("variant", w.Input{Value: "mix@example.org", Variant: w.InputFill, Label: "Address"}, "account-1", "tc-in", 2),
	form("variant", w.Input{Value: "https://carrel.example", Mono: true, Variant: w.Input11_5, Label: "Address"}, "s-admin-4", "tc-in", 1),
	form("variant", w.Input{Value: "https://carrel.example/dav/", Mono: true, Variant: w.Input11_5Fill, Label: "Address"}, "components-20", "tc-in", 2),
	form("waiting", w.Filter{Query: w.Query{Placeholder: "Filter by word", Label: "Filter by word"}, Waiting: true, Variant: w.FilterH26W200}, "s-notes-1", "tc-in", 1),
	form("variant", w.Input{Value: "https://carrel.example/dav/", Mono: true, Variant: w.Input11, Label: "Address"}, "s-davloom-3", "tc-in", 2),
	form("variant", w.Input{Value: "webcal://carrel.example/p/7f3ac9/personal.ics", Mono: true, Variant: w.InputLineFill, Label: "Link"}, "s-publish-6", "tc-in", 1),
	form("placeholder", w.Input{Kind: w.InputSearch, Placeholder: "Filter this folder", Variant: w.InputW190, Label: "Filter this folder"}, "components-16", "tc-in", 1),
	form("placeholder", w.Input{Kind: w.InputDate, Placeholder: "YYYY-MM-DD", Label: "Start date", Mono: true}, "s-tasks-3", "tc-in", 4),
	form("variant", w.Input{Value: "60", Mono: true, Variant: w.InputW52, Label: "Match threshold"}, "s-dup-1", "tc-in", 1),
	form("focus", w.Input{Value: "https://dav.office.example/dav.php/", Mono: true, Focus: true, Label: "Server URL"}, "s-profile-2", "tc-in", 2),
	form("variant", w.Input{Value: "Carrel · personal", Variant: w.Input11, Label: "Name"}, "s-backup-4", "tc-in", 1),
	form("variant", w.Input{Value: "https://carrel.example/invite/7f3a…", Mono: true, Variant: w.Input11Fill, Label: "Link"}, "s-admin-1", "tc-in", 2),
	form("placeholder", w.Input{Placeholder: "Documents · sftpgo / carrel-backup/…", Variant: w.Input11Fill, Label: "Folder"}, "components-20", "tc-in", 1),
	form("placeholder", w.Input{Kind: w.InputSearch, Placeholder: "Find a zone", Label: "Find a zone"}, "s-agenda-8", "tc-in", 1),
	form("pick", w.Pick{Name: "to.1", Options: []w.PickOption{{Value: "a.kovaleva@slovo.example", Label: "a.kovaleva@slovo.example", On: true}, {Value: "anna@home.example", Label: "anna@home.example"}, {Value: "", Label: "Do not rewrite"}}, Label: "Rewrite a.kovaleva@example.org to", Variant: w.PickT12W240}, "s-contacts-29", "tc-in", 1),
	form("pick", w.Pick{Name: "reminder.1.unit", Options: []w.PickOption{{Value: "minutes", Label: "minutes before"}, {Value: "hours", Label: "hours before"}, {Value: "days", Label: "days before", On: true}, {Value: "weeks", Label: "weeks before"}}, Label: "Unit"}, "s-agenda-19", "tc-in", 3),
	form("pick", w.Pick{Name: "reminder.1", Options: []w.PickOption{{Value: "at.the.time", Label: "At the time"}, {Value: "5.minutes.before", Label: "5 minutes before"}, {Value: "10.minutes.before", Label: "10 minutes before", On: true}, {Value: "30.minutes.before", Label: "30 minutes before"}, {Value: "1.hour.before", Label: "1 hour before"}, {Value: "1.day.before", Label: "1 day before"}, {Value: "custom", Label: "Custom…", Shows: "reminder.1.count reminder.1.unit"}, {Value: "none", Label: "No reminder"}}, Label: "First"}, "s-agenda-19", "tc-in", 1),
	form("select", w.Pick{Name: "remind", Options: []w.PickOption{{Value: "10.minutes.before", Label: "10 minutes before", On: true}}, Label: "Remind"}, "s-notes-9", "tc-in", 6),
	form("marked", w.Pick{Value: "nextcloud · Записи", Mark: &w.Mark{Colour: "#4A6B52"}}, "s-notes-20", "tc-in", 1),
	form("aside", w.Pick{Value: "Asia/Yekaterinburg", Aside: &w.Dim{Parts: []w.Inline{w.Text("as sent")}, Variant: w.DimIndent}, Label: "Time zone"}, "s-agenda-3", "tc-in", 11),
	form("error", w.Input{Value: "nextcloud", Error: true, Focus: true, Label: "Label"}, "s-profile-11", "tc-in", 1),
	form("disabled", w.Input{Kind: w.InputPassword, Label: "Password", Disabled: true}, "s-profile-14", "tc-in", 4),
	form("variant", w.Input{Value: "…/remote.php/dav/calendars/mix/birthdays/", Mono: true, Variant: w.InputDusty, Label: "Address"}, "s-profile-7", "tc-in", 2),
	form("placeholder", w.Input{Placeholder: "Дни рождения", Focus: true, Label: "Type the name to confirm"}, "s-profile-9", "tc-in", 1),
	form("focus", w.Input{Value: "https://feeds.example/ru-holidays.ics", Mono: true, Focus: true, Variant: w.Input11_5, Label: "Feed address"}, "s-profile-10", "tc-in", 1),
	form("variant", w.Input{Value: "brick-lamp-9f2-carrel", Mono: true, Variant: w.InputKeyFill, Label: "Password"}, "s-davloom-3", "tc-in", 1),
	form("variant", w.Input{Value: "cbk_7f3a-91de-4402-lamp", Mono: true, Variant: w.InputKeySmFill, Label: "Token"}, "s-backup-5", "tc-in", 1),
	form("variant", w.Input{Kind: w.InputDate, Value: "2026-08-01", Mono: true, Variant: w.InputH26W120, Label: "From"}, "s-admin-7", "tc-in", 2),
	form("placeholder", w.Input{Placeholder: "a-z 0-9 . _ -", Focus: true, Mono: true, Label: "Login"}, "s-auth-4", "tc-in", 1),
	form("placeholder", w.Input{Placeholder: `E:\Carrel`, Mono: true, Variant: w.Input11Fill, Label: "Folder"}, "desktop-3", "tc-in", 1),

	// tc-in as a pick
	form("pick-empty", w.Pick{Value: "—", Unset: true}, "mentions-7", "tc-in", 6),
	form("select", w.Pick{Name: "week", Options: []w.PickOption{{Value: "mon", Label: "Monday", On: true}, {Value: "sun", Label: "Sunday"}, {Value: "sat", Label: "Saturday", Unavailable: true}}, Label: "Week starts on"}, "components-38", "tc-in", 2),
	form("value", w.Pick{Value: "Confirmed", Label: "Status"}, "components-10", "tc-in", 5),
	form("variant", w.Pick{Value: "nextcloud · Записи", Variant: w.PickT12W200, Label: "Where to write"}, "s-notes-5", "tc-in", 1),
	form("variant", w.Pick{Value: "nextcloud · Записи", Variant: w.PickT12Fill, Label: "Where to write"}, "desktop-9", "tc-in", 2),
	form("variant", w.Pick{Value: "sftpgo · vps", Label: "Server", Variant: w.PickT12W240}, "s-notes-11", "tc-in", 1),
	form("focus", w.Pick{Value: "radicale · vps · Проекты", Focus: true, Label: "Restore into"}, "s-backup-8", "tc-in", 1),
	form("unset", w.Pick{Value: "All actions", Unset: true, Variant: w.PickH26W190, Label: "Action"}, "s-admin-7", "tc-in", 1),

	// tc-in as a filter, with the query line in it
	form("filter-form", w.Filter{Query: w.Query{Placeholder: "Filter", Label: "Filter", Name: "filter"}, Action: "/app/contacts/nextcloud/kontakty", Variant: w.FilterH26W220}, "components-43", "tc-in", 1),
	form("waiting", w.Filter{Query: w.Query{Placeholder: "Filter the loaded cards", Label: "Filter the loaded cards"}, Waiting: true}, "s-contacts-1", "tc-in", 1),
	form("waiting", w.Filter{Query: w.Query{Placeholder: "Filter this folder", Label: "Filter this folder"}, Waiting: true, Variant: w.FilterH26W190}, "s-files-1", "tc-in", 1),
	form("quiet", w.Filter{Query: w.Query{Value: "подряд", Label: "подряд"}, Quiet: true, Variant: w.FilterT12W200}, "s-mail-5", "tc-in", 1),
	form("default", w.Query{Placeholder: "Search everything…", Label: "Search everything…"}, "shell-1", "tc-q", 1),
	form("typed", w.Query{Value: "ковал", Label: "Search everything…"}, "s-find-1", "tc-q", 1),

	// tc-in as tags
	form("focus", w.TagInput{Parts: []w.TagInputPart{w.Tag{Label: "Забрать документы", Removable: true}}, Add: &w.TagAdd{Name: "related.add", Placeholder: "Title or UID", Label: "Add a link"}, Focus: true}, "s-tasks-10", "tc-in", 1),
	form("tags", w.TagInput{Parts: []w.TagInputPart{w.Tag{Label: "Calendar · 118"}, w.Tag{Label: "Tasks · 6"}, w.Tag{Label: "Notes · 2"}}}, "s-profile-10", "tc-in", 3),
	form("tags", w.TagInput{Parts: []w.TagInputPart{}, Add: &w.TagAdd{Name: "related.add", Placeholder: "Title or UID", Label: "Add a link"}}, "s-tasks-14", "tc-in", 17),
	form("tags", w.TagInput{Parts: []w.TagInputPart{w.Tag{Label: "Забрать документы", Removable: true}}, Add: &w.TagAdd{Name: "related.add", Placeholder: "Title or UID", Label: "Add a link"}}, "s-tasks-3", "tc-in", 16),
	form("focus", w.TagInput{Parts: []w.TagInputPart{w.Tag{Label: "Синк-ап по релизу", Removable: true}, w.Tag{Label: "Дописать раздел про кэш", Removable: true}}, Add: &w.TagAdd{Name: "related.add", Value: "ковал", Placeholder: "Title or UID", Label: "Add a link"}, Focus: true}, "s-notes-4", "tc-in", 7),
	form("add", w.TagInput{Parts: []w.TagInputPart{w.Tag{Label: "x", Fixed: true}, w.Tag{Label: "x", Removable: true}, w.Tag{Label: "x", Removable: true}, w.Tag{Label: "x", Removable: true}}, Columns: true, Name: "attendee", Add: &w.TagAdd{Name: "attendee.add", Placeholder: "Name or address", Label: "Add attendee", Suggest: "/app/contacts/suggest"}}, "s-agenda-3", "tc-in", 15),
	form("add", w.TagInput{Parts: []w.TagInputPart{w.Tag{Label: "x", Fixed: true}, w.Tag{Label: "x", Removable: true}, w.Tag{Label: "x", Removable: true}, w.Tag{Label: "x", Removable: true}}, Columns: true, Name: "attendee", Add: &w.TagAdd{Name: "attendee.add", Value: "ве", Placeholder: "Name or address", Label: "Add attendee", Suggest: "/app/contacts/suggest"}}, "s-agenda-12", "tc-in", 1),
	form("value", w.Pick{Value: "Europe/Moscow"}, "s-agenda-15", "tc-in", 11),
	form("add", w.TagInput{Columns: true, Add: &w.TagAdd{Name: "attendee.add", Placeholder: "Name or address", Label: "Add attendee", Suggest: "/app/contacts/suggest"}}, "s-agenda-15", "tc-in", 14),
	form("date", w.Input{Kind: w.InputDate, Mono: true, Label: "Until"}, "s-agenda-15", "tc-in", 19),
	form("columns", w.TagInput{Parts: []w.TagInputPart{w.Tag{Label: "x", Fixed: true}, w.Tag{Label: "x", Fixed: true}, w.Tag{Label: "x", Fixed: true}, w.Tag{Label: "x", Fixed: true}}, Columns: true}, "s-agenda-14", "tc-in", 1),
	form("wrap", w.TagInput{Wrap: true, Focus: true, Parts: []w.TagInputPart{
		w.Tag{Label: "Анна Ковалёва · a.kovaleva@example.org", Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11}},
		w.Tag{Label: "Артём Белов · a.belov@example.org", Mark: &w.Mark{Colour: "#6B4E7A", Size: w.MarkH11}}}}, "s-mail-1", "tc-in", 1),
	form("wrap", w.TagInput{Wrap: true, Focus: true, Parts: []w.TagInputPart{
		w.Tag{Label: "Анна Ковалёва · a.kovaleva@example.org", Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11}},
		w.Span{Face: w.FaceMuted, Text: "бел"}}}, "s-mail-3", "tc-in", 1),

	// tc-in as a text area
	form("plain", w.Input{Value: "ул. Садовая, 12, кв. 5", Label: "Street"}, "s-contacts-3", "tc-in", 16),
	form("body", w.Area{Text: "", Variant: w.AreaBody, Name: "description", Label: "Text"}, "s-notes-30", "tc-in", 8),
	form("title", w.Area{Text: "", Variant: w.AreaTitle, Focus: true, Name: "summary", Label: "Text"}, "s-notes-30", "tc-in", 2),
	form("area", w.Area{Text: "Пишет по вторникам, отвечает в тот же день.", Label: "Text"}, "s-contacts-3", "tc-in", 22),
	form("area-empty", w.Area{Label: "Text"}, "s-contacts-27", "tc-in", 14),
	form("body", w.Area{Text: "Позвонить в МФЦ во вторник до 12:00 — уточнить, нужен ли оригинал доверенности.", Variant: w.AreaBodySm, Focus: true, Label: "Text"}, "s-notes-5", "tc-in", 2),
	form("title", w.Area{Text: "Заметка со встречи 10.08", Variant: w.AreaTitle, Label: "Title"}, "s-notes-20", "tc-in", 2),
	form("body", w.Area{Text: noteBody, Variant: w.AreaBody, Focus: true, Label: "Text"}, "s-notes-20", "tc-in", 8),
	form("title", w.Area{Text: "Заметка со встречи 10.08", Variant: w.AreaTitle, Focus: true, Label: "Title"}, "s-notes-4", "tc-in", 2),
	form("body", w.Area{Text: noteBody, Variant: w.AreaBody, Label: "Text"}, "s-notes-4", "tc-in", 8),
	form("paste", w.Area{Text: pasteArgs, Variant: w.AreaPaste, Mono: true, Label: "Arguments"}, "s-backup-4", "tc-in", 5),
	form("letter", w.Area{Text: "Анна, Артём — вот карточка подрядчика и смета, которую он прислал. Договор в папке, ссылка ниже.", Variant: w.AreaLetter, Label: "Note"}, "s-mail-1", "tc-in", 3),
	form("tall", w.Area{Text: "Спросить у Виктора про миграцию — успеет ли до заморозки.", Variant: w.AreaTall, Focus: true, Label: "Text"}, "desktop-9", "tc-in", 1),
)

const noteBody = `Договорились заморозить ветку в среду.

Письмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.

## Что делаем

- [ ] проверить changelog
- [x] попросить у Артёма скриншоты

Перенос старых вложений в новую папку — Виктор посмотрит, сколько их вообще, и скажет в четверг.`

const pasteArgs = `-fsS
-m
30
-X
POST
-H
Authorization: Bearer cbk_7f3a-91de-4402-lamp
https://carrel.example/api/backup/run?job=personal`

func init() {
	Widgets = append(Widgets, w.Input{}, w.Pick{}, w.Filter{}, w.Query{}, w.TagInput{}, w.Area{})
}
