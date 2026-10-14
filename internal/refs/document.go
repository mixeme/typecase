package refs

import w "carrel.dev/typecase/widget"

func document(state string, fixture Renderer, frame, class string, n int) Entry {
	return e("document", state, fixture, ref(frame, class, n))
}

var _ = add(
	// tc-prose
	document("inline", w.Prose{Parts: []w.Inline{w.Text("Итоги созвона в "), w.Em("четверг"), w.Text(": релиз держат две вещи, обе — до пятницы. Миграцию гоняем только с "), w.Mono{Text: "--dry-run", Variant: w.MonoChip}, w.Text(", подробности — в "), w.Link{Text: "заметке о схеме", Href: "#"}, w.Text(".")}, Variant: w.ProseRead}, "s-notes-25", "tc-prose", 1),
	document("muted", w.Prose{Parts: []w.Inline{w.Text("This login has no confirmed email — an address that was entered but never confirmed does not count. There is no letter we can send, and there is no password reset that would keep your data.\n\nIf you deposited the encryption key, ask the administrator in person. Otherwise this account cannot be recovered from the sign-in screen.")}, Variant: w.ProseMuted}, "s-auth-7", "tc-prose", 1),
	document("plain", w.Prose{Parts: []w.Inline{w.Text("Календарь на осень: три события в «Личный», одно в «Работа». Связанные — Анна и смета.")}}, "components-29", "tc-prose", 1),
	document("size", w.Prose{Parts: []w.Inline{w.Text("Обсудить оставшиеся блокеры, дату заморозки и кто пишет письмо пользователям.")}, Variant: w.Prose13}, "s-agenda-1", "tc-prose", 1),
	document("lead", w.Prose{Parts: []w.Inline{w.Text("Договорились заморозить ветку в среду.\n\nПисьмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.")}, Variant: w.ProseLead}, "s-notes-2", "tc-prose", 1),
	document("muted", w.Prose{Parts: []w.Inline{w.Text("Enter the login or the confirmed email of the account. There is no way to set a new password and keep the data stored here: the encryption key is derived from the password, and the server never has that key.")}, Variant: w.ProseMutedClose}, "s-auth-5", "tc-prose", 1),
	document("lead", w.Prose{Parts: []w.Inline{w.Text("Перенос старых вложений в новую папку — Виктор посмотрит, сколько их вообще, и скажет в четверг.")}, Variant: w.ProseLeadAfter}, "s-notes-2", "tc-prose", 2),
	document("size", w.Prose{Parts: []w.Inline{w.Text("Анна, Артём — вот карточка подрядчика и смета, которую он прислал. Договор в папке, ссылка ниже.\n\nМихаил Енученко")}, Variant: w.Prose13Card}, "s-mail-7", "tc-prose", 1),
	document("rich", w.Prose{Parts: []w.Inline{w.Text("Account "), w.Strong("mix"), w.Text(" has a confirmed email. A silent password reset does not exist: a new password would be a new key, and everything stored under the old one would be lost anyway.")}, Variant: w.ProseMuted}, "s-auth-6", "tc-prose", 1),
	document("read", w.Prose{Parts: []w.Inline{w.Text("Договорились заморозить ветку в среду.\n\nПисьмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.")}, Variant: w.ProseRead}, "s-notes-1", "tc-prose", 1),
	document("top", w.Prose{Parts: []w.Inline{w.Text("This is a connection point for clients, not a page to read. Enter the address in anything that speaks CalDAV, CardDAV or WebDAV — DAVx5, Thunderbird, Calendar and Contacts on Apple, a file manager mounting a network drive. The user name is your Carrel one; the password is the device password issued in Carrel.")}, Variant: w.ProseTop}, "s-davloom-4", "tc-prose", 1),
	document("rich", w.Prose{Parts: []w.Inline{w.Text("The letter went to "), w.Mono{Text: "v@example.org"}, w.Text(". The account cannot be used until you open the link in it.\n\nSending again with the same login and address, while the account is unconfirmed, simply replaces the letter and the password.")}, Variant: w.ProseMuted}, "s-auth-3", "tc-prose", 1),
	document("rich", w.Prose{Parts: []w.Inline{w.Text("Договорились заморозить ветку в среду. Письмо пользователям пишет "), w.Link{Text: "@Анна Ковалёва", Href: "#"}, w.Text(", черновик к вечеру вторника; сверку данных делает "), w.Mention{Name: "@Пётр Ковалёв"}, w.Text(".")}}, "mentions-4", "tc-prose", 1),
	document("read", w.Prose{Parts: []w.Inline{w.Text("Договорились заморозить ветку в среду.\n\nПисьмо пользователям пишет Анна, черновик к вечеру вторника.")}, Variant: w.ProseReadTop}, "phone-7", "tc-prose", 1),

	// tc-doc
	document("read", w.Doc{Items: []w.DocPart{
		w.DocTitle{Text: "Заметка со встречи 10.08"},
		w.Prose{Parts: []w.Inline{w.Text("Договорились заморозить ветку в среду.\n\nПисьмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.")}, Variant: w.ProseLead}}}, "s-notes-32", "tc-doc", 1),
	document("edit", w.Doc{Items: []w.DocPart{
		w.Area{Text: "", Variant: w.AreaTitle, Focus: true, Name: "summary", Label: "Text"},
		w.DocFields{Items: []w.Field{w.Field{Label: "Notebook", Parts: []w.FieldPart{w.Drop{Trigger: w.Pick{Value: "nextcloud · Записи", Mark: &w.Mark{Colour: "#4A6B52"}}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Записи", On: true, Name: "notebook", Value: "nextcloud/zapisi"},
			w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "nextcloud · Дневник", Name: "notebook", Value: "nextcloud/dnevnik"},
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "nextcloud · Общие с семьёй", Reason: "Read-only: this account cannot write here", Name: "notebook", Value: "nextcloud/semya"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}}, Width: w.FieldFlex2_280}, w.Field{Label: "Date", Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "dtstart.date", Value: "2026-08-10", Label: "Date", Mono: true}}, Width: w.FieldFix150}, w.Field{Label: "Time", Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Name: "dtstart.time", Value: "11:05", Label: "Time", Mono: true}}, Width: w.FieldFix130}, w.Field{Label: "Tags", Parts: []w.FieldPart{w.Input{Name: "categories", Placeholder: "—", Label: "Tags"}}, Width: w.FieldFlex3_320}, w.Field{Label: "Related to", Parts: []w.FieldPart{w.TagInput{Parts: []w.TagInputPart{}}}, Width: w.FieldFlex3_340}}},
		w.Area{Text: "", Variant: w.AreaBody, Name: "description", Label: "Text"}}, ID: "note-new", Action: "/app/notes/nextcloud/zapisi/new"}, "s-notes-30", "tc-doc", 1),
	document("source", w.Doc{Items: []w.DocPart{
		w.DocTitle{Text: "Заметка со встречи 10.08"},
		w.DocCode{Text: "Договорились заморозить ветку в среду.\n\nПисьмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.\n\n## Что делаем\n\n- [ ] проверить changelog\n- [x] попросить у Артёма скриншоты\n\nПеренос старых вложений в новую папку — Виктор посмотрит, сколько их вообще, и скажет в четверг.", Wrap: true}}, Text: true}, "s-notes-27", "tc-doc", 1),
	document("parts", w.Doc{Items: []w.DocPart{
		w.DocTitle{Text: "Разбор релиза 0.3"},
		w.Prose{Parts: []w.Inline{w.Text("Итоги созвона в "), w.Em("четверг"), w.Text(": релиз держат две вещи, обе — до пятницы. Миграцию гоняем только с "), w.Mono{Text: "--dry-run", Variant: w.MonoChip}, w.Text(", подробности — в "), w.Link{Text: "заметке о схеме", Href: "#"}, w.Text(".")}, Variant: w.ProseRead},
		w.DocHeading{Text: "Что держит"},
		w.DocBullets{Items: [][]w.Inline{[]w.Inline{w.Text("миграция схемы — "), w.Link{Text: "задача у Виктора", Href: "#"}}, []w.Inline{w.Text("тесты"), w.DocBullets{Items: [][]w.Inline{[]w.Inline{w.Mono{Text: "go test -race ./...", Variant: w.MonoChip}, w.Text(" — зелёный")}, []w.Inline{w.Text("нагрузочный — после 20.08")}}}}}},
		w.DocQuote{Parts: []w.Inline{w.Text("Не выпускаем, пока не прогоним миграцию на копии базы.")}},
		w.DocHeading{Text: "Как проверить"},
		w.DocCode{Text: "make db-copy\n./carrel migrate --dry-run --from 0.2 --to 0.3 --database ./copy/carrel.db"}}, Text: true}, "s-notes-25", "tc-doc", 1),
	document("table", w.Doc{Items: []w.DocPart{
		w.DocTitle{Text: "Сборы в Казань"},
		w.Prose{Parts: []w.Inline{w.Text("Что купить до четверга и кто за что отвечает.")}, Variant: w.ProseLead},
		w.DocTable{Head: [][]w.Inline{[]w.Inline{w.Text("Что")}, []w.Inline{w.Text("Кто")}, []w.Inline{w.Text("Когда")}, []w.Inline{w.Text("Сколько")}, []w.Inline{w.Text("Где")}, []w.Inline{w.Text("Статус")}, []w.Inline{w.Text("Заметка")}}, Rows: [][][]w.Inline{{[]w.Inline{w.Text("Чемодан на колёсах, большой")}, []w.Inline{w.Text("Анна")}, []w.Inline{w.Text("12.08")}, []w.Inline{w.Text("4 500")}, []w.Inline{w.Link{Text: "ozon.ru/t/chemodan", Href: "#"}}, []w.Inline{w.Text("заказан")}, []w.Inline{w.Text("синий, не чёрный")}}, {[]w.Inline{w.Text("Билеты до Казани, туда и обратно")}, []w.Inline{w.Strong("Виктор")}, []w.Inline{w.Text("до 10.08")}, []w.Inline{w.Text("11 200")}, []w.Inline{}, []w.Inline{w.Text("ждём")}, []w.Inline{w.Text("поезд, не самолёт")}}, {[]w.Inline{w.Text("Страховка")}, []w.Inline{w.Text("Анна")}, []w.Inline{}, []w.Inline{w.Text("1 350")}, []w.Inline{w.Mono{Text: "alfastrah"}}, []w.Inline{w.Text("куплена")}, []w.Inline{}}}, Align: []w.DocAlign{w.DocAlignStart, w.DocAlignStart, w.DocAlignCenter, w.DocAlignEnd, w.DocAlignStart, w.DocAlignStart, w.DocAlignStart}, Switch: w.Seg{Options: []w.SegOption{{Label: "Wrap"}, {Label: "One line", On: true}}, Variant: w.SegCompactSmall}},
		w.Prose{Parts: []w.Inline{w.Text("Остальное — на месте.")}, Variant: w.ProseLeadAfter}}}, "s-notes-22", "tc-doc", 1),
	document("table", w.Doc{Items: []w.DocPart{
		w.DocTitle{Text: "Сборы в Казань"},
		w.Prose{Parts: []w.Inline{w.Text("Что купить до четверга и кто за что отвечает.")}, Variant: w.ProseLead},
		w.DocTable{Head: [][]w.Inline{[]w.Inline{w.Text("Что")}, []w.Inline{w.Text("Кто")}, []w.Inline{w.Text("Когда")}, []w.Inline{w.Text("Сколько")}, []w.Inline{w.Text("Где")}, []w.Inline{w.Text("Статус")}, []w.Inline{w.Text("Заметка")}}, Rows: [][][]w.Inline{{[]w.Inline{w.Text("Чемодан на колёсах, большой")}, []w.Inline{w.Text("Анна")}, []w.Inline{w.Text("12.08")}, []w.Inline{w.Text("4 500")}, []w.Inline{w.Link{Text: "ozon.ru/t/chemodan", Href: "#"}}, []w.Inline{w.Text("заказан")}, []w.Inline{w.Text("синий, не чёрный")}}, {[]w.Inline{w.Text("Билеты до Казани, туда и обратно")}, []w.Inline{w.Strong("Виктор")}, []w.Inline{w.Text("до 10.08")}, []w.Inline{w.Text("11 200")}, []w.Inline{}, []w.Inline{w.Text("ждём")}, []w.Inline{w.Text("поезд, не самолёт")}}, {[]w.Inline{w.Text("Страховка")}, []w.Inline{w.Text("Анна")}, []w.Inline{}, []w.Inline{w.Text("1 350")}, []w.Inline{w.Mono{Text: "alfastrah"}}, []w.Inline{w.Text("куплена")}, []w.Inline{}}}, Align: []w.DocAlign{w.DocAlignStart, w.DocAlignStart, w.DocAlignCenter, w.DocAlignEnd, w.DocAlignStart, w.DocAlignStart, w.DocAlignStart}, Switch: w.Seg{Options: []w.SegOption{{Label: "Wrap", On: true}, {Label: "One line"}}, Variant: w.SegCompactSmall}},
		w.Prose{Parts: []w.Inline{w.Text("Остальное — на месте.")}, Variant: w.ProseLeadAfter}}}, "s-notes-21", "tc-doc", 1),
	document("foreign", w.Doc{Items: []w.DocPart{
		w.Area{Text: "Заметка со встречи 10.08", Variant: w.AreaTitle, Focus: true, Name: "summary", Label: "Text"},
		w.DocFields{Items: []w.Field{w.Field{Label: "Notebook", Parts: []w.FieldPart{w.Drop{Trigger: w.Pick{Value: "nextcloud · Записи", Mark: &w.Mark{Colour: "#4A6B52"}}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Записи", On: true, Name: "notebook", Value: "nextcloud/zapisi"},
			w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "nextcloud · Дневник", Name: "notebook", Value: "nextcloud/dnevnik"},
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "nextcloud · Общие с семьёй", Reason: "Read-only: this account cannot write here", Name: "notebook", Value: "nextcloud/semya"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}}, Width: w.FieldFlex2_280}, w.Field{Label: "Date", Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "dtstart.date", Value: "2026-08-10", Label: "Date", Mono: true}}, Width: w.FieldFix150}, w.Field{Label: "Time", Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Name: "dtstart.time", Value: "11:05", Label: "Time", Mono: true}}, Width: w.FieldFix130}, w.Field{Label: "Tags", Parts: []w.FieldPart{w.Input{Name: "categories", Value: "встречи, релиз", Label: "Tags"}}, Width: w.FieldFlex3_320}, w.Field{Label: "Related to", Parts: []w.FieldPart{w.Drop{Open: true, Trigger: w.TagInput{Parts: []w.TagInputPart{w.Tag{Label: "Синк-ап по релизу", Removable: true}, w.Dim{Parts: []w.Inline{w.Text("ковал")}}}, Focus: true}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuItem{Icon: "cal", Trail: w.Mono{Text: "10 Aug · Личный", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Звонок с Ковалёвой", On: true},
			w.MenuItem{Icon: "task", Trail: w.Mono{Text: "task · Работа", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Ответить Анне по смете"},
			w.MenuItem{Icon: "con", Trail: w.Mono{Text: "contact", Variant: w.Mono10Dusty, Place: w.Pushed}, Label: "Анна Ковалёва"},
			w.MenuSep{},
			w.MenuItem{Label: "Enter a UID by hand", Tone: w.MenuToneDusty}}, Width: w.MenuWidthFull, Pop: w.MenuUnder}}}, Width: w.FieldFlex3_340}}},
		w.Area{Text: "Договорились заморозить ветку в среду.\n\nПисьмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.\n\n## Что делаем\n\n- [ ] проверить changelog\n- [x] попросить у Артёма скриншоты\n\n## Отложено до следующего спринта\n\nПеренос старых вложений в новую папку — Виктор посмотрит, сколько их вообще.", Variant: w.AreaBody, Name: "description", Label: "Text"},
		w.DocNote{Parts: []w.Inline{w.Text("Properties Carrel does not know — "), w.Mono{Text: "X-JTX-COLOR"}, w.Text(" and two more — are kept unchanged and not shown here.")}}}, ID: "note-edit", Action: "/app/notes/nextcloud/zapisi/9e31-40b7-c012"}, "s-notes-4", "tc-doc", 1),
	document("read", w.Doc{Items: []w.DocPart{
		w.DocTitle{Text: "Заметка со встречи 10.08"},
		w.Prose{Parts: []w.Inline{w.Text("Договорились заморозить ветку в среду. Письмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.")}, Variant: w.ProseLead},
		w.DocTasks{Title: w.Rubric{Text: "Что делаем", Variant: w.RubricBottom8}, Items: []w.DocTask{{Tick: w.Tick{Variant: w.TickNudge, Label: "проверить changelog"}, Text: "проверить changelog"}, {Tick: w.Tick{Done: true, Variant: w.TickNudge, Label: "попросить у Артёма скриншоты"}, Text: "попросить у Артёма скриншоты"}}},
		w.Prose{Parts: []w.Inline{w.Text("Перенос старых вложений в новую папку — Виктор посмотрит, сколько их вообще, и скажет в четверг.")}, Variant: w.ProseLeadAfter}}}, "s-notes-2", "tc-doc", 1),
	document("edit", w.Doc{Items: []w.DocPart{
		w.Area{Text: "Заметка со встречи 10.08", Variant: w.AreaTitle, Name: "summary", Label: "Text"},
		w.DocFields{Items: []w.Field{w.Field{Label: "Notebook", Parts: []w.FieldPart{w.Drop{Trigger: w.Pick{Value: "nextcloud · Записи", Mark: &w.Mark{Colour: "#4A6B52"}}, Menu: w.Menu{Items: []w.MenuPart{
			w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Записи", On: true, Name: "notebook", Value: "nextcloud/zapisi"},
			w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "nextcloud · Дневник", Name: "notebook", Value: "nextcloud/dnevnik"},
			w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "nextcloud · Общие с семьёй", Reason: "Read-only: this account cannot write here", Name: "notebook", Value: "nextcloud/semya"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}}, Width: w.FieldFlex2_280}, w.Field{Label: "Date", Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "dtstart.date", Value: "2026-08-10", Label: "Date", Mono: true}}, Width: w.FieldFix150}, w.Field{Label: "Time", Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Name: "dtstart.time", Value: "11:05", Label: "Time", Mono: true}}, Width: w.FieldFix130}, w.Field{Label: "Tags", Parts: []w.FieldPart{w.Input{Name: "categories", Value: "встречи, релиз", Label: "Tags"}}, Width: w.FieldFlex3_320}, w.Field{Label: "Related to", Parts: []w.FieldPart{w.TagInput{Parts: []w.TagInputPart{w.Tag{Label: "Синк-ап по релизу", Removable: true}}}}, Width: w.FieldFlex3_340}}},
		w.Area{Text: "Договорились заморозить ветку в среду.\n\nПисьмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.\n\n## Что делаем\n\n- [ ] проверить changelog\n- [x] попросить у Артёма скриншоты\n\nПеренос старых вложений в новую папку — Виктор посмотрит, сколько их вообще, и скажет в четверг.", Variant: w.AreaBody, Focus: true, Name: "description", Label: "Text"},
		w.DocNote{Parts: []w.Inline{w.Text("Properties Carrel does not know — "), w.Mono{Text: "X-JTX-COLOR"}, w.Text(" and two more — are kept unchanged and not shown here.")}}}, ID: "note-edit", Action: "/app/notes/nextcloud/zapisi/9e31-40b7-c012"}, "s-notes-20", "tc-doc", 1),
	document("public", w.Doc{Items: []w.DocPart{
		w.DocTitle{Text: "Заметка со встречи 10.08"},
		w.DocSub{Parts: []w.DocSubPart{w.Mono{Text: "10 August 2026, 11:05"}}},
		w.Prose{Parts: []w.Inline{w.Text("Договорились заморозить ветку в среду. Письмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.")}, Variant: w.ProseLead},
		w.DocTasks{Title: w.Rubric{Text: "Что делаем", Variant: w.RubricBottom8}, Items: []w.DocTask{{Tick: w.Tick{Variant: w.TickNudge, Label: "проверить changelog"}, Text: "проверить changelog"}, {Tick: w.Tick{Done: true, Variant: w.TickNudge, Label: "попросить у Артёма скриншоты"}, Text: "попросить у Артёма скриншоты"}}},
		w.DocFiles{Title: w.Rubric{Text: "Attachments", Variant: w.RubricBottom7}, Items: []w.DocFile{{Name: "2026-08-10-sync-board.png", Size: w.Dim{Parts: []w.Inline{w.Text("· 412 KB")}}}, {Name: "release-checklist.md", Size: w.Dim{Parts: []w.Inline{w.Text("· 3 KB")}}}}}}}, "s-publish-8", "tc-doc", 1),
	document("text", w.Doc{Items: []w.DocPart{
		w.DocHeading{Text: "Перед заморозкой"},
		w.DocBullets{Items: [][]w.Inline{[]w.Inline{w.Text("прогнать "), w.Mono{Text: "go test -race ./...", Variant: w.MonoChip}}, []w.Inline{w.Text("проверить changelog")}, []w.Inline{w.Text("письмо пользователям — черновик у Анны")}}},
		w.DocHeading{Text: "После"},
		w.DocBody{Text: "Тег ставит Виктор. Образ собирается вручную до появления CI."}}, Text: true}, "s-publish-10", "tc-doc", 1),
	document("phone", w.Doc{Items: []w.DocPart{
		w.DocTitle{Text: "Заметка со встречи 10.08"},
		w.DocSub{Parts: []w.DocSubPart{w.Mono{Text: "10 August, 11:05"}, w.Text(" · "), w.Tag{Label: "встречи"}}},
		w.Prose{Parts: []w.Inline{w.Text("Договорились заморозить ветку в среду. Письмо пользователям пишет Анна, черновик к вечеру вторника.")}, Variant: w.ProseReadTop},
		w.DocTasks{Title: w.Rubric{Text: "Что делаем", Variant: w.RubricBottom7}, Items: []w.DocTask{{Tick: w.Tick{Variant: w.TickNudge, Label: "проверить changelog"}, Text: "проверить changelog"}, {Tick: w.Tick{Done: true, Variant: w.TickNudge, Label: "попросить скриншоты"}, Text: "попросить скриншоты"}}, Phone: true},
		w.DocRelated{Title: w.Rubric{Text: "Related to", Variant: w.RubricBottom7}, Name: "Синк-ап по релизу", What: w.Dim{Parts: []w.Inline{w.Text("· event")}}}}}, "phone-7", "tc-doc", 1),

	// tc-docbar
	document("edit", w.Docbar{Items: []w.DocbarPart{w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "note-new"}, w.Button{Label: "Cancel", Href: "/app/notes/nextcloud/zapisi"}, w.BarSep{}, w.Button{Label: "B", Bold: true}, w.Button{Label: "I", Variant: w.ButtonItalic}, w.Button{Label: "H2", Tertiary: true}, w.Button{Label: "• list", Tertiary: true}, w.Button{Folds: true, Label: "task", Icon: "task", Title: "Task", Tertiary: true}, w.Button{Icon: "link", Title: "Link", Tertiary: true}, w.Drop{Trigger: w.Button{Folds: true, Label: "Attach", Menu: true, Icon: "up", Title: "Attach"}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "up", File: &w.FileField{Form: "note-new", Name: "file", Label: "From this computer", List: "note-new-attachments"}, Row: w.DocAttach{Name: "{name}", Size: w.Dim{Parts: []w.Inline{w.Text("· {size}")}}, Remove: &w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Removes: true}}, Label: "From this computer"},
		w.MenuItem{Icon: "file", Label: "From Files"}}, Width: w.MenuWidth212, Pop: w.MenuUnder}}}, More: &w.Button{Icon: "dots", Title: "More"}, End: []w.DocbarEnd{w.Check{Parts: []w.CheckPart{w.Text("Auto-save")}, Tertiary: true}, w.Kbd{Keys: "Ctrl+S", Secondary: true}, w.Seg{Options: []w.SegOption{{Label: "Full width", On: true}, {Label: "Reading width"}}, Secondary: true}, w.Button{Folds: true, Label: "Focus", Icon: "focus", Title: "Focus"}}, Edit: true}, "s-notes-30", "tc-docbar", 1),
	document("plain", w.Docbar{Items: []w.DocbarPart{w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "note-edit"}, w.Button{Label: "Cancel", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012"}, w.BarSep{}, w.Button{Label: "B", Bold: true}, w.Button{Label: "I", Variant: w.ButtonItalic}, w.Button{Label: "H2", Tertiary: true}, w.Button{Label: "• list", Tertiary: true}, w.Button{Folds: true, Label: "task", Icon: "task", Title: "Task", Tertiary: true}, w.Button{Icon: "link", Title: "Link", Tertiary: true}, w.Drop{Trigger: w.Button{Folds: true, Label: "Attach", Menu: true, Icon: "up", Title: "Attach"}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "up", File: &w.FileField{Form: "note-edit", Name: "file", Label: "From this computer", List: "note-attachments"}, Row: w.DocAttach{Name: "{name}", Size: w.Dim{Parts: []w.Inline{w.Text("· {size}")}}, Remove: &w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Removes: true}}, Label: "From this computer"},
		w.MenuItem{Icon: "file", Label: "From Files"}}, Width: w.MenuWidth212, Pop: w.MenuUnder}}}, More: &w.Button{Icon: "dots", Title: "More"}, End: []w.DocbarEnd{w.Check{Parts: []w.CheckPart{w.Text("Auto-save")}, Tertiary: true}, w.DraftNote{Text: "Draft in this window ·", Time: w.Mono{Text: "11:42"}, Secondary: true}, w.Kbd{Keys: "Ctrl+S", Secondary: true}, w.Seg{Options: []w.SegOption{{Label: "Full width", On: true}, {Label: "Reading width"}}, Secondary: true}, w.Button{Folds: true, Label: "Focus", Icon: "focus", Title: "Focus"}}, Edit: true}, "s-notes-20", "tc-docbar", 1),
	document("rule", w.Docbar{Items: []w.DocbarPart{w.Button{Label: "Save", Kind: w.ButtonPrimary}, w.Button{Label: "Cancel"}, w.BarSep{}, w.Button{Title: "More"}}, End: []w.DocbarEnd{w.Kbd{Keys: "Ctrl+S"}}, Rule: true}, "components-29", "tc-docbar", 1),

	// tc-docmeta

)

func init() {
	Widgets = append(Widgets, w.Prose{}, w.Doc{}, w.Docbar{})
	Parts = append(Parts, w.Link{}, w.DocTitle{}, w.DocSub{}, w.DocTask{}, w.DocTasks{}, w.DocFile{}, w.DocFiles{}, w.DocRelated{},
		w.DocFields{}, w.DocNote{}, w.DocHeading{}, w.DocBullets{}, w.DocBody{}, w.DocTable{}, w.DraftNote{}, w.DocLayout{}, w.Hero{})
}
