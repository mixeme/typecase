package proof

// Four more frames: a dark one, a sheet over a list, a confirmation over
// settings, settings with an open row menu. Widgets and data only.

func topWith(focus, refresh bool) Top {
	acts := Nodes{Btn{Label: "New note", Icon: "quill", Chev: true, Primary: true, Focus: focus}}
	if refresh {
		acts = append(acts, Btn{Label: "2 min ago", Icon: "refresh"})
	}
	return Top{Search: "Search everything…", Actions: acts, User: "mix"}
}

func calNav() []NavItem {
	n := nav("Calendar")
	return n[:5] // these frames have no Duplicates
}

func calHead(mark19 bool, acts ...Node) Head {
	return Head{
		Mark:   &Mark{Multi: true},
		Mark19: mark19,
		Title:  "Combined calendar",
		Sub:    Nodes{Text("5 of 7 calendars · 10 – 16 August 2026")},
		Acts:   acts,
	}
}

var newEvent = Btn{Label: "New event", Icon: "plus", Chev: true, Primary: true}

// Dark is frame «Календарь · тёмная тема».
func Dark() Screen {
	return Screen{
		Dark: true,
		Top:  topWith(false, false),
		Rail: Rail{
			Nav: calNav(),
			Sections: []Section{
				{Rows: []Source{{Label: "Combined calendar", Mark: Mark{Multi: true}, Meta: "14", Box: Box{Part: true}, All: true, On: true}}},
				{Rubric: "nextcloud · home", Rows: []Source{
					{Label: "Личный", Mark: Mark{Colour: "#4A6B52"}, Meta: "7"},
					{Label: "Работа", Mark: Mark{Colour: "#3F6E8C"}, Meta: "4"},
					{Label: "Дни рождения", Mark: Mark{Colour: "#A05A2C"}, Meta: "—", Box: Box{Off: true}, Off: true},
				}},
				{Rubric: "baikal · office", Rows: []Source{
					{Label: "Общий", Mark: Mark{Colour: "#8C6239"}, Meta: "password · 401", Error: true,
						MetaTitle: "401 Unauthorized — the saved password was refused. Open Connections to update it.",
						Action:    &Btn{Label: "Retry", Small: true}},
				}},
			},
		},
		Main: Nodes{
			calHead(true, newEvent, Btn{Label: "Export"}),
			Poll{Percent: 83, Body: Nodes{Text("Polled 5 of 6 sources · 14 events · "), Inline{Alert, "1 unavailable"}}},
			Group{Rubric: "Monday · 10 August", Count: 3},
			List{Rows: Nodes{
				AgendaRow{Mark: Mark{Colour: "#4A6B52"}, Time: "10:00", Title: "Синк-ап по релизу", Place: "Переговорная 2", Badge: "4 attendees", Source: "nextcloud · Личный", On: true},
				AgendaRow{Mark: Mark{Colour: "#8C6239"}, Time: "all day", Title: "Дежурство: первая линия", Place: "—", Badge: "repeats", Source: "radicale · Дежурства"},
				AgendaRow{Mark: Mark{Colour: "#3F6E8C"}, Time: "15:00", Title: "Созвон с подрядчиком", Place: "Jitsi", Source: "nextcloud · Работа"},
			}},
		},
		Side: &Side{
			Head: Head{
				Title: "Синк-ап по релизу",
				Size:  TitlePanel,
				Sub:   Nodes{Text("nextcloud · Личный")},
				Acts:  Nodes{Btn{Label: "Edit", Small: true}, Btn{Icon: "close", Close: true, Title: "Close the panel"}},
			},
			Body: Nodes{
				Fields{
					{Label: "When", Rich: Nodes{Inline{Mono, "10:00 – 11:00"}}},
					{Label: "Repeats", Value: "Weekly on Mon"},
					{Label: "Where", Value: "Переговорная 2"},
					{Label: "Status", Value: "Confirmed"},
				},
				Sec{Rubric: "Source", Body: Nodes{Diag("Read  2 min ago")}},
			},
		},
	}
}

// Quick is frame «Быстрая заметка поверх текущего экрана».
func Quick() Screen {
	return Screen{
		Top: topWith(true, true),
		Rail: Rail{
			Nav: calNav(),
			Sections: []Section{
				{Rows: []Source{{Label: "Combined calendar", Mark: Mark{Multi: true}, Meta: "14", Box: Box{Part: true}, All: true, On: true}}},
				{Rubric: "nextcloud · home", Rows: []Source{
					{Label: "Личный", Mark: Mark{Colour: "#4A6B52"}, Meta: "7"},
					{Label: "Работа", Mark: Mark{Colour: "#3F6E8C"}, Meta: "4"},
				}},
			},
		},
		Main: Nodes{
			calHead(false, newEvent),
			Group{Rubric: "Monday · 10 August", Count: 3},
			List{Rows: Nodes{
				AgendaRow{Mark: Mark{Colour: "#4A6B52"}, Time: "10:00", Title: "Синк-ап по релизу", Place: "Переговорная 2", Source: "nextcloud · Личный"},
				AgendaRow{Mark: Mark{Colour: "#8C6239"}, Time: "all day", Title: "Дежурство: первая линия", Place: "—", Source: "2 sources"},
				AgendaRow{Mark: Mark{Colour: "#3F6E8C"}, Time: "15:00", Title: "Созвон с подрядчиком", Place: "Jitsi", Source: "nextcloud · Работа"},
			}},
			Group{Rubric: "Tuesday · 11 August", Count: 2},
			List{Rows: Nodes{
				AgendaRow{Mark: Mark{Colour: "#6B4E7A"}, Time: "09:30", Title: "Ревью архитектуры", Place: "Переговорная 1", Source: "radicale · Проекты"},
			}},
		},
		Overlay: &Overlay{Sheet: Sheet{
			Icon:  Icon{Name: "quill", Size: 15, Stroke: "1.4", Caps: true, Tint: TintAccent},
			Title: "New note",
			Pick:  &Pick{Label: "Записи · nextcloud"},
			Close: true,
			Body: Nodes{
				In{ID: "q-text", Kind: InArea, Note: true, Focus: true, Aria: "Note",
					Value: "Позвонить в МФЦ во вторник до 12:00 — уточнить, нужен ли оригинал доверенности."},
				Inline{SmallDim, "The title is taken from the first line. The date is now."},
			},
			Foot: Nodes{
				Btn{Label: "Save note", Primary: true},
				Kbd("Ctrl+Enter"),
				Btn{Label: "Open full screen", Icon: "expand"},
				FootEnd{Btn{Label: "Cancel"}, Kbd("Esc")},
			},
		}},
	}
}

func settingsRail() Rail {
	return Rail{
		Nav: []NavItem{{Icon: "back", Label: "Back to app", Href: "#app"}},
		Sections: []Section{{Rubric: "Settings", Nav: []NavItem{
			{Icon: "user", Label: "Account", Href: "#account"},
			{Icon: "server", Label: "Connections", Href: "#connections", Count: num(3), On: true},
			{Icon: "file", Label: "Attachments", Href: "#attachments"},
			{Icon: "device", Label: "New device", Href: "#devices"},
			{Icon: "davloom", Label: "davloom", Href: "#davloom", Count: num(3)},
			{Icon: "archive", Label: "Backup", Href: "#backup", Count: num(3)},
			{Icon: "link", Label: "Published links", Href: "#published", Count: num(2)},
			{Icon: "theme", Label: "Appearance", Href: "#appearance"},
		}}},
	}
}

// collection is one row of a server card table.
func collection(colour, name, kind string, objects, changed string, off bool, menu *Menu, kindBadge string) Tr {
	what := Nodes{Text(kind)}
	if kindBadge != "" {
		what = Nodes{Text(kind + " "), Badge{Label: kindBadge}}
	}
	cells := []Td{
		{Body: Nodes{Box{Off: off, Label: "Show " + name}}},
		{Body: Nodes{FName{Mark: &Mark{Colour: colour}, Name: name}}},
		{Mono: true, Body: what},
		{Mono: true, Body: Nodes{Text(objects)}},
	}
	if changed != "" {
		cells = append(cells, Td{Mono: true, Body: Nodes{Text(changed)}})
	}
	cells = append(cells, Td{Acts: true, Body: Nodes{Btn{Icon: "dots", Small: true, Narrow: true, Title: "More", Menu: menu}}})
	return Tr{Off: off, Cells: cells}
}

func serverCols(first string, changed bool) []Col {
	cols := []Col{{Kind: ColTickSm}, {Label: first}, {Label: "What is in it"}, {Label: "Objects", Kind: ColCount}}
	if changed {
		cols = append(cols, Col{Label: "Changed on the server", Kind: ColChanged})
	}
	return append(cols, Col{Kind: ColMenu})
}

func sm(label string) Btn           { return Btn{Label: label, Small: true} }
func smIcon(label, icon string) Btn { return Btn{Label: label, Icon: icon, Small: true} }
func smDanger(label string) Btn     { return Btn{Label: label, Small: true, Danger: true} }

// Delete is frame «Удаление коллекции».
func Delete() Screen {
	return Screen{
		Top:  topWith(false, true),
		Tall: true,
		Rail: settingsRail(),
		Main: Nodes{
			Head{
				Title: "Connections",
				Tight: true,
				Sub:   Nodes{Text("3 servers · 12 collections · 2 subscriptions")},
				Acts:  Nodes{Btn{Label: "Connect account", Icon: "plus", Primary: true}},
			},
			Pad{Card{
				Name:    "nextcloud · home",
				Actions: []Btn{smIcon("New collection", "plus"), sm("Show password"), sm("Edit")},
				URL:     "https://cloud.example.org/remote.php/dav/ · mix · last answer 2 min ago",
				Table: &Table{Cols: serverCols("Collection", false), Rows: []Tr{
					collection("#4A6B52", "Личный", "calendar", "128", "", false, nil, ""),
					collection("#A05A2C", "Дни рождения", "calendar", "51", "", false, nil, ""),
				}},
			}},
		},
		Overlay: &Overlay{Top26: true, Sheet: Sheet{
			Width: Sheet560,
			Icon:  Icon{Name: "trash", Size: 15, Stroke: "1.6", Tint: TintAlert},
			Title: "Delete on the server",
			Meta:  "Дни рождения · nextcloud",
			Body: Nodes{
				Msg{Alert: true, Flush: true, Body: Nodes{
					Text("The calendar and all "), Inline{Bold, "51"}, Text(" events in it are deleted on "), Inline{Mono, "cloud.example.org"},
					Text(". Carrel holds no copy — it reads and writes, it does not store — so there is nothing here to undo it with."),
				}},
				Impacts{Title: "What is pointing at it now", Items: []Impact{
					{What: "1 published link", Effect: "— answers “not found” from that moment"},
					{What: "Backup job «Личное на домашний диск»", Effect: "— copies this collection; it will run one short and say so"},
					{What: "3 devices through davloom", Effect: "— it disappears for them at the next sync"},
				}},
				Msg{Flush: true, Body: Nodes{
					Text("Last time it was in an archive: "), Inline{Mono, "2026-08-14"}, Text(". Restoring brings the records back into a "),
					Inline{Bold, "new"}, Text(" collection — not into this one, and not to this address."),
				}},
				Line{
					Btn{Icon: "down", Rich: Nodes{Text("Download "), Inline{Mono11, ".ics"}, Text(" first")}},
					Inline{SmallDim, "51 events, one file, straight to you"},
				},
				Field{Label: "Type the name to confirm", Width: W100, In: In{ID: "d-name", Kind: InText, Placeholder: "Дни рождения", Focus: true}},
			},
			Foot: Nodes{
				Btn{Label: "Delete on the server", DangerSolid: true},
				FootEnd{Btn{Label: "Cancel"}, Kbd("Esc")},
			},
		}},
	}
}

var rowMenu = &Menu{Right: true, Wide: true, Items: []MenuItem{
	{Icon: "pencil", Label: Nodes{Text("Rename and colour")}},
	{Icon: "move", Label: Nodes{Text("Move to another server")}},
	{Icon: "dup", Label: Nodes{Text("Copy to another server")}},
	{Icon: "down", Label: Nodes{Text("Export "), Inline{Mono11, ".ics"}}},
	{Icon: "link", Label: Nodes{Text("Publish a link")}},
	{Sep: true},
	{Icon: "trash", Label: Nodes{Text("Delete on the server…")}, Danger: true},
}}

func feedRow(name string, objects string, off bool, changed *Td) Tr {
	cells := []Td{
		{Body: Nodes{Box{Off: off, Label: "Show " + name}}},
		{Body: Nodes{FName{Mark: &Mark{Colour: "#8C6239", Ext: true}, Name: name}}},
		{Mono: true, Body: Nodes{Text(objects)}},
	}
	if changed != nil {
		cells = append(cells, *changed)
	}
	return Tr{Off: off, Cells: cells}
}

// Connections is frame «Настройки · подключения и подписки».
func Connections() Screen {
	return Screen{
		Top:  topWith(false, true),
		Rail: settingsRail(),
		Main: Nodes{
			Head{
				Title: "Connections",
				Sub:   Nodes{Text("Signed in as "), Inline{Bold, "mix"}, Text(" · 3 servers · 12 collections · 2 subscriptions")},
				Acts:  Nodes{Btn{Label: "Connect account", Icon: "plus", Primary: true}, Btn{Label: "Subscribe to a feed", Icon: "link"}},
			},
			Stack{
				ListRubric{Text: "DAV servers", Count: 3},
				Card{
					Name:    "nextcloud · home",
					Actions: []Btn{smIcon("New collection", "plus"), sm("Show password"), sm("Test"), sm("Edit"), smDanger("Disconnect")},
					URL:     "https://cloud.example.org/remote.php/dav/ · mix · last answer 2 min ago",
					Table: &Table{Cols: serverCols("Collection", true), Rows: []Tr{
						collection("#4A6B52", "Личный", "calendar", "128", "3 h ago", false, nil, ""),
						collection("#3F6E8C", "Работа", "calendar", "74", "yesterday", false, nil, ""),
						collection("#A05A2C", "Дни рождения", "calendar", "31", "12 August", true, rowMenu, ""),
						collection("#4A6B52", "Контакты", "address book", "248", "4 d ago", false, nil, ""),
						collection("#4A6B52", "Записи", "journal", "86", "2 h ago", false, nil, ""),
						collection("#4A6B52", "Дела", "tasks", "30", "20 min ago", false, nil, ""),
					}},
				},
				Card{
					Name:    "sftpgo · vps",
					Actions: []Btn{smIcon("New folder", "plus"), sm("Show password"), sm("Test"), sm("Edit"), smDanger("Disconnect")},
					URL:     "https://files.example.org/dav/ · mix · plain WebDAV, no principal · last answer 5 min ago",
					Table: &Table{Cols: serverCols("Folder", true), Rows: []Tr{
						collection("#4A6B52", "Documents", "files", "412", "5 min ago", false, nil, ""),
						collection("#9A9280", "Public", "files", "96", "3 d ago", true, nil, "ro"),
					}},
				},
				Card{
					Name:    "baikal · office",
					Alert:   true,
					Actions: []Btn{sm("Show password"), sm("Test again"), sm("Edit"), smDanger("Disconnect")},
					URL:     "https://dav.office.example/dav.php/ · m.yenuchenko · last answer refused",
					Diag: "[well-known]  GET /.well-known/caldav → 404\n" +
						"[base]        PROPFIND /dav.php/ → 401 Unauthorized\n" +
						"              WWW-Authenticate: Basic realm=\"BaikalDAV\"\n" +
						"[principal]   не запрашивался: предыдущий шаг не прошёл",
					Note: "The server refused the password. It may have been changed on the server: the one Carrel saved no longer fits.",
				},
				ListRubric{Text: "Subscriptions · external, read-only", Count: 2},
				Card{
					Name:    "Праздники РФ",
					Badge:   "feed",
					Actions: []Btn{sm("Check"), sm("Edit"), smDanger("Unsubscribe")},
					URL:     "https://feeds.example/ru-holidays.ics · no password · read 2 min ago · 34 KB · 0.4 s",
					Table: &Table{
						Cols: []Col{{Kind: ColTickSm}, {Label: "Lands in"}, {Label: "Objects", Kind: ColCount}, {Label: "Changed on the server", Kind: ColChanged}},
						Rows: []Tr{
							feedRow("Calendar", "118", false, &Td{Mono: true, Rows: 3, Body: Nodes{Text("yesterday")}}),
							feedRow("Tasks", "6", true, nil),
							feedRow("Notes", "2", false, nil),
						},
					},
				},
				Card{
					Name:    "Расписание лиги",
					Badge:   "feed",
					Alert:   true,
					Actions: []Btn{sm("Show password"), sm("Check again"), sm("Edit"), smDanger("Unsubscribe")},
					URL:     "webcal://league.example/season.ics · password saved · last read yesterday",
					Diag: "[fetch]   GET /season.ics → 401 Unauthorized\n" +
						"              WWW-Authenticate: Basic realm=\"league\"\n" +
						"[cache]   46 events from yesterday are still shown, marked stale",
					Note: "The feed asks for a password and the saved one no longer fits. A feed that stopped answering must not look like a feed with nothing in it.",
				},
			},
		},
	}
}
