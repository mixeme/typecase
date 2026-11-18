// Раздел макета: 10. Узкий экран.

package screens

import w "carrel.dev/typecase/widget"

// Phone1 is frame «Агенда» (phone-1).
func Phone1() Screen {
	return w.Shell{
		Top: &w.Top{
			Lead:  &w.Button{Icon: "menu", IconBox: w.IconBox19, Variant: w.ButtonPadX4, Title: "Menu"},
			Brand: w.Brand{Title: "Combined calendar"},
			Actions: []w.TopAction{
				w.Button{Icon: "search", IconBox: w.IconBox17, Variant: w.ButtonPadX6, Title: "Search", Href: "/app/search"},
				w.Button{
					Icon:    "quill",
					IconBox: w.IconBox16,
					Kind:    w.ButtonPrimary,
					Variant: w.ButtonPadX8,
					Title:   "New note",
				},
			},
			Tight: true,
		},
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Runbar{
						Result: w.RunFailed,
						Parts: []w.RunbarPart{
							w.Say{Full: "Polled · 5 of 6 answered · 7 events ·", Short: "5 of 6 · 7 events ·"},
							w.Drop{
								Trigger: w.Button{Label: "1 failed", Menu: true, Kind: w.ButtonDanger, Small: true},
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuRefusal{
											Mark:   w.Mark{Colour: "#9A9280", Size: w.MarkH11},
											Source: "baikal · Общий",
											Answer: "PROPFIND /dav.php/calendars/mix/ → 401 Unauthorized: the saved password was refused. Open Connections to update it.",
											Retry:  w.Button{Label: "Retry", Small: true},
										},
									},
									Width: w.MenuWidth420,
									Pop:   w.MenuUnderEnd,
								},
							},
							w.Mono{Text: "2 min ago", Place: w.Pushed},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Mon · 10 Aug", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
					w.List{
						Rows: []w.ListRow{
							w.PhoneRow{
								Mark:  &w.Mark{Colour: "#4A6B52", Size: w.MarkH32},
								Title: "Синк-ап по релизу",
								Sub:   []w.Inline{w.Mono{Text: "10:00 – 11:00"}, w.Text(" · Переговорная 2")},
								Aside: "Личный",
							},
							w.PhoneRow{
								Mark:  &w.Mark{Colour: "#3F6E8C", Size: w.MarkH32},
								Title: "Дежурство: первая линия",
								Sub:   []w.Inline{w.Mono{Text: "all day"}, w.Text(" · repeats")},
								Aside: "Работа",
							},
							w.PhoneRow{
								Mark:  &w.Mark{Colour: "#8C6239", Size: w.MarkH32},
								Title: "Дежурство: первая линия",
								Sub:   []w.Inline{w.Mono{Text: "all day"}, w.Text(" · repeats")},
								Aside: "Дежурства",
							},
							w.PhoneRow{
								Mark:  &w.Mark{Colour: "#3F6E8C", Size: w.MarkH32},
								Title: "Созвон с подрядчиком",
								Sub:   []w.Inline{w.Mono{Text: "15:00"}, w.Text(" · Jitsi")},
								Aside: "Работа",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Tue · 11 Aug", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.PhoneRow{
								Mark:  &w.Mark{Colour: "#6B4E7A", Size: w.MarkH32},
								Title: "Ревью архитектуры",
								Sub:   []w.Inline{w.Mono{Text: "09:30"}, w.Text(" · Переговорная 1")},
								Aside: "Проекты",
							},
							w.PhoneRow{
								Mark:  &w.Mark{Colour: "#4A6B52", Size: w.MarkH32},
								Title: "Забрать документы",
								Sub:   []w.Inline{w.Mono{Text: "18:00"}, w.Text(" · МФЦ, Мясницкая 13")},
								Aside: "Личный",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Wed · 12 Aug", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.PhoneRow{
								Mark:  &w.Mark{Colour: "#4A6B52", Size: w.MarkH32},
								Title: "Поездка в Казань",
								Sub:   []w.Inline{w.Mono{Text: "all day"}, w.Text(" · day 1 of 3")},
								Aside: "Личный",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Thu · 13 Aug", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.PhoneRow{
								Mark:  &w.Mark{Colour: "#4A6B52", Size: w.MarkH32},
								Title: "Поездка в Казань",
								Sub:   []w.Inline{w.Mono{Text: "all day"}, w.Text(" · day 2 of 3")},
								Aside: "Личный",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Fri · 14 Aug", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.PhoneRow{
								Mark:  &w.Mark{Colour: "#4A6B52", Size: w.MarkH32},
								Title: "Поездка в Казань",
								Sub:   []w.Inline{w.Mono{Text: "all day"}, w.Text(" · day 3 of 3")},
								Aside: "Личный",
							},
						},
					},
					w.BottomNav{
						Items: []w.BottomItem{
							{Icon: "con", Label: "Contacts", Href: "#"},
							{Icon: "cal", Label: "Calendar", Href: "#"},
							{Icon: "task", Label: "Tasks", Href: "#"},
							{Icon: "note", Label: "Notes", Href: "#"},
							{Icon: "file", Label: "Files", Href: "#"},
						},
						On: 2,
					},
				},
			},
		},
	}
}

// Phone2 is frame «Выдвижной рейл источников» (phone-2).
func Phone2() Screen {
	return w.Shell{
		Top: &w.Top{
			Lead:    &w.Button{Icon: "back", IconBox: w.IconBox19, Variant: w.ButtonPadX4, Title: "Back"},
			Brand:   w.Brand{Title: "Sources"},
			Actions: []w.TopAction{w.Button{Label: "Apply", Small: true}},
		},
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.RailSec{
						Items: []w.RailSecPart{
							w.TickRow{
								Ticked:  w.TickedPart,
								Kind:    w.TickRowAll,
								Form:    w.TickRowTouch,
								Mark:    w.Mark{Kind: w.MarkMulti},
								Label:   "Combined calendar",
								Meta:    "7",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud", Variant: w.RubricRail},
							w.TickRow{Form: w.TickRowTouch, Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "3", Current: true},
							w.TickRow{Form: w.TickRowTouch, Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "2"},
							w.TickRow{
								Ticked: w.TickedOff,
								Form:   w.TickRowTouch,
								Mark:   w.Mark{Colour: "#A05A2C"},
								Label:  "Дни рождения",
								Meta:   "—",
								Hidden: true,
							},
						},
						Apart: true,
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "radicale · vps", Variant: w.RubricRail},
							w.TickRow{Form: w.TickRowTouch, Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "1"},
							w.TickRow{Form: w.TickRowTouch, Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "1"},
						},
						Apart: true,
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "baikal", Variant: w.RubricRail},
							w.TickRow{Form: w.TickRowTouch, Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Error: true},
						},
						Apart: true,
					},
					w.Empty{
						Parts:   []w.EmptyPart{w.Text("Unticked collections are not polled at all, which saves both time and mobile traffic.")},
						Variant: w.EmptySmall,
					},
				},
			},
		},
	}
}

// Phone3 is frame «Карточка события» (phone-3).
func Phone3() Screen {
	return w.Shell{
		Top: &w.Top{
			Lead:    &w.Button{Icon: "back", IconBox: w.IconBox19, Variant: w.ButtonPadX4, Title: "Back"},
			Brand:   w.Brand{Title: "Event"},
			Actions: []w.TopAction{w.Button{Label: "Edit", Small: true}},
		},
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Синк-ап по релизу", Size: w.HeadTitle19},
						Subs: []w.HeadSub{
							w.HeadSub{Parts: []w.HeadSubPart{w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, w.Text("nextcloud · Личный")}},
						},
						Pad: w.HeadSnug,
					},
					w.Fields{
						Items: []w.FieldsItem{
							{
								Term:  "When",
								Value: []w.ValuePart{w.Mono{Text: "10:00 – 11:00"}, w.Break{}, w.Text("Monday, 10 August")},
							},
							{Term: "Repeats", Value: []w.ValuePart{w.Text("Weekly on Mon")}},
							{Term: "Where", Value: []w.ValuePart{w.Text("Переговорная 2")}},
							{Term: "Status", Value: []w.ValuePart{w.Text("Confirmed")}},
						},
						Variant: w.Fields88,
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Description"},
							w.Prose{
								Parts:   []w.Inline{w.Text("Обсудить оставшиеся блокеры, дату заморозки и кто пишет письмо пользователям.")},
								Variant: w.Prose13,
							},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Attendees 4"},
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{
										w.Mark{Colour: "#6B4E7A", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "Мария Лебедева", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· organiser")}},
									},
									[]w.SecItemPart{
										w.Mark{Colour: "#4A6B52", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "Анна Ковалёва", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· accepted")}},
									},
									[]w.SecItemPart{
										w.Mark{Colour: "#6B4E7A", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "Артём Белов", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· tentative")}},
									},
									[]w.SecItemPart{w.Mono{Text: "v@example.org"}, w.Text(" "), w.Dim{Parts: []w.Inline{w.Text("· no reply")}}},
								},
								Columns: true,
							},
						},
					},
					w.FootActs{
						Items: []w.Button{
							w.Button{Label: "Write a note", Variant: w.ButtonCentredFill},
							w.Button{Label: "Delete", Kind: w.ButtonDangerSolid, Variant: w.ButtonCentred},
						},
					},
				},
			},
		},
	}
}

// Phone4 is frame «Контакты» (phone-4).
func Phone4() Screen {
	return w.Shell{
		Top: &w.Top{
			Lead:  &w.Button{Icon: "menu", IconBox: w.IconBox19, Variant: w.ButtonPadX4, Title: "Menu"},
			Brand: w.Brand{Title: "Combined address book"},
			Actions: []w.TopAction{
				w.Button{Icon: "search", IconBox: w.IconBox17, Variant: w.ButtonPadX6, Title: "Search", Href: "/app/search"},
				w.Button{
					Icon:    "quill",
					IconBox: w.IconBox16,
					Kind:    w.ButtonPrimary,
					Variant: w.ButtonPadX8,
					Title:   "New note",
				},
			},
			Tight: true,
		},
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
					w.List{
						Rows: []w.ListRow{
							w.PhoneRow{
								Mark:    &w.Mark{Colour: "#4A6B52", Size: w.MarkH34},
								Avatar:  &w.Avatar{Initials: "АС", Size: w.Avatar34},
								Title:   "Алексей Смирнов",
								Sub:     []w.Inline{w.Text("+7 916 200-71-15")},
								SubMono: true,
								Aside:   "Контакты",
							},
							w.PhoneRow{
								Mark:       &w.Mark{Colour: "#4A6B52", Size: w.MarkH34},
								Avatar:     &w.Avatar{Initials: "АК", Size: w.Avatar34},
								Title:      "Анна Ковалёва",
								TitleBadge: &w.Badge{Label: "×3", Tone: w.BadgeLinked},
								Sub:        []w.Inline{w.Text("+7 912 445-10-88")},
								SubMono:    true,
								Aside:      "Контакты",
							},
							w.PhoneRow{
								Mark:       &w.Mark{Colour: "#6B4E7A", Size: w.MarkH34},
								Avatar:     &w.Avatar{Photo: true, Size: w.Avatar34},
								Title:      "Артём Белов",
								TitleBadge: &w.Badge{Label: "×2", Tone: w.BadgeLinked},
								Sub:        []w.Inline{w.Text("+7 903 118-22-04")},
								SubMono:    true,
								Aside:      "Команда",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "В", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.PhoneRow{
								Mark:   &w.Mark{Colour: "#A05A2C", Size: w.MarkH34},
								Avatar: &w.Avatar{Initials: "ВЛ", Size: w.Avatar34},
								Title:  "Вера Лебедева",
								Sub:    []w.Inline{w.Text("vera@example.net")},
								Aside:  "Архив",
							},
							w.PhoneRow{
								Mark:   &w.Mark{Colour: "#4A6B52", Size: w.MarkH34},
								Avatar: &w.Avatar{Initials: "ВГ", Size: w.Avatar34},
								Title:  "Виктор Гончаров",
								Sub:    []w.Inline{w.Text("v.goncharov@example.org")},
								Aside:  "Контакты",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Л", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.PhoneRow{
								Mark:   &w.Mark{Colour: "#4A6B52", Size: w.MarkH34},
								Avatar: &w.Avatar{Photo: true, Size: w.Avatar34},
								Title:  "Лебедева Вера",
								Sub:    []w.Inline{w.Text("v.lebedeva@example.org")},
								Aside:  "Контакты",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "М", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.PhoneRow{
								Mark:   &w.Mark{Colour: "#4A6B52", Size: w.MarkH34},
								Avatar: &w.Avatar{Initials: "МЕ", Size: w.Avatar34},
								Title:  "Михаил Енученко",
								Sub:    []w.Inline{w.Text("mix@example.org")},
								Aside:  "Контакты",
							},
						},
					},
					w.BottomNav{
						Items: []w.BottomItem{
							{Icon: "con", Label: "Contacts", Href: "#"},
							{Icon: "cal", Label: "Calendar", Href: "#"},
							{Icon: "task", Label: "Tasks", Href: "#"},
							{Icon: "note", Label: "Notes", Href: "#"},
							{Icon: "file", Label: "Files", Href: "#"},
						},
						On: 1,
					},
				},
			},
		},
	}
}

// Phone5 is frame «Человек целиком» (phone-5).
func Phone5() Screen {
	return w.Shell{
		Top: &w.Top{
			Lead:    &w.Button{Icon: "back", IconBox: w.IconBox19, Variant: w.ButtonPadX4, Title: "Back"},
			Brand:   w.Brand{Title: "Анна Ковалёва"},
			Actions: []w.TopAction{w.Button{Label: "Edit", Small: true}},
		},
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Person{
						Initials: "АК",
						Name:     "Анна Ковалёва",
						Sub:      "Издательство «Слово»",
						Badge:    &w.Badge{Label: "in 3 books", Tone: w.BadgeLinked},
						Acts:     []w.Button{w.Button{Label: "Call"}, w.Button{Label: "Email"}, w.Button{Icon: "quill", Title: "New note"}},
					},
					w.Bar{
						Items: []w.BarPart{w.Seg{Options: []w.SegOption{{Label: "Card", On: true}, {Label: "Connections 27"}}}},
					},
					w.Fields{
						Items: []w.FieldsItem{
							{Term: "Phone", Value: []w.ValuePart{w.Text("+7 912 445-10-88")}, Face: w.ValueMonoAccent},
							{Term: "Email", Value: []w.ValuePart{w.Text("a.kovaleva@example.org")}, Face: w.ValueAccent},
							{Term: "Birthday", Value: []w.ValuePart{w.Text("1987-03-14")}, Face: w.ValueMono},
						},
						Variant: w.Fields88Pad,
					},
					w.BottomNav{
						Items: []w.BottomItem{
							{Icon: "con", Label: "Contacts", Href: "#"},
							{Icon: "cal", Label: "Calendar", Href: "#"},
							{Icon: "task", Label: "Tasks", Href: "#"},
							{Icon: "note", Label: "Notes", Href: "#"},
							{Icon: "file", Label: "Files", Href: "#"},
						},
						On: 1,
					},
				},
			},
		},
	}
}

// Phone6 is frame «Файлы · выделение» (phone-6).
func Phone6() Screen {
	return w.Shell{
		Top: &w.Top{
			Lead:  &w.Button{Icon: "menu", IconBox: w.IconBox19, Variant: w.ButtonPadX4, Title: "Menu"},
			Brand: w.Brand{Title: "2026-08"},
			Actions: []w.TopAction{
				w.Button{Icon: "upload", IconBox: w.IconBox16, Kind: w.ButtonPrimary, Variant: w.ButtonPadX8, Title: "Upload"},
			},
		},
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Bar{
						Items: []w.BarPart{
							w.Button{Icon: "download", Title: "Download", Form: "files-sel", Name: "op", Value: "download"},
							w.Button{Icon: "move", Title: "Move", Form: "files-sel", Name: "op", Value: "move"},
							w.Button{
								Icon:  "trash",
								Kind:  w.ButtonDangerSolid,
								Title: "Delete",
								Form:  "files-sel",
								Name:  "op",
								Value: "delete",
							},
							w.Button{Icon: "selectall", Title: "Select all", Place: w.Pushed, ScriptOnly: true},
							w.Button{Icon: "close", Title: "Clear selection", ScriptOnly: true},
						},
						Count:   &w.BarCount{Text: "2 selected", Small: true},
						Variant: w.BarSelection,
					},
					w.List{
						Rows: []w.ListRow{
							w.PhoneRow{
								Pick:    &w.TickBox{Send: w.Send{Of: "files-sel", Name: "item", Value: "2026-08-10-sync-board.png"}},
								Icon:    "img",
								Title:   "2026-08-10-sync-board.png",
								Sub:     []w.Inline{w.Text("412 KB · image/png")},
								SubMono: true,
								More:    &w.Button{Icon: "dots", IconBox: w.IconBox16, Variant: w.ButtonPadX4, Title: "More"},
								On:      true,
							},
							w.PhoneRow{
								Pick:    &w.TickBox{Send: w.Send{Of: "files-sel", Name: "item", Value: "release-checklist.md"}},
								Icon:    "doc",
								Title:   "release-checklist.md",
								Sub:     []w.Inline{w.Text("3 KB · text/markdown")},
								SubMono: true,
								More:    &w.Button{Icon: "dots", IconBox: w.IconBox16, Variant: w.ButtonPadX4, Title: "More"},
								On:      true,
							},
							w.PhoneRow{
								Pick:    &w.TickBox{Ticked: w.TickedOff, Send: w.Send{Of: "files-sel", Name: "item", Value: "смета-подрядчик.pdf"}},
								Icon:    "doc",
								Title:   "смета-подрядчик.pdf",
								Sub:     []w.Inline{w.Text("1.2 MB · application/pdf")},
								SubMono: true,
								More:    &w.Button{Icon: "dots", IconBox: w.IconBox16, Variant: w.ButtonPadX4, Title: "More"},
							},
						},
					},
					w.PostForm{
						ID:     "files-sel",
						Action: "/app/files/sftpgo/documents/op/carrel/attachments/2026-08",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
					w.BottomNav{
						Items: []w.BottomItem{
							{Icon: "con", Label: "Contacts", Href: "#"},
							{Icon: "cal", Label: "Calendar", Href: "#"},
							{Icon: "task", Label: "Tasks", Href: "#"},
							{Icon: "note", Label: "Notes", Href: "#"},
							{Icon: "file", Label: "Files", Href: "#"},
						},
						On: 5,
					},
				},
			},
		},
	}
}

// Phone7 is frame «Заметка · чтение и правка» (phone-7).
func Phone7() Screen {
	return w.Shell{
		Top: &w.Top{
			Lead:    &w.Button{Icon: "back", IconBox: w.IconBox19, Variant: w.ButtonPadX4, Title: "Back"},
			Brand:   w.Brand{Title: "Записи"},
			Actions: []w.TopAction{w.Button{Label: "Edit", Icon: "pencil", Small: true}},
		},
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Doc{
						Items: []w.DocPart{
							w.DocTitle{Text: "Заметка со встречи 10.08"},
							w.DocSub{Parts: []w.DocSubPart{w.Mono{Text: "10 August, 11:05"}, w.Text(" · "), w.Tag{Label: "встречи"}}},
							w.Prose{
								Parts: []w.Inline{
									w.Text("Договорились заморозить ветку в среду.\n\nПисьмо пользователям пишет Анна, черновик к вечеру вторника."),
								},
								Variant: w.ProseReadTop,
							},
							w.DocTasks{
								Title: "Что делаем",
								Items: []w.DocTask{
									{Tick: w.Tick{Variant: w.TickNudge, Sign: true}, Text: "проверить changelog"},
									{Tick: w.Tick{Done: true, Variant: w.TickNudge, Sign: true}, Text: "попросить скриншоты"},
								},
								Phone: true,
							},
							w.DocRelated{
								Title: w.Rubric{Text: "Related to", Variant: w.RubricBelow},
								Name:  "Синк-ап по релизу",
								What:  w.Dim{Parts: []w.Inline{w.Text("· event")}},
							},
						},
					},
					w.BottomNav{
						Items: []w.BottomItem{
							{Icon: "con", Label: "Contacts", Href: "#"},
							{Icon: "cal", Label: "Calendar", Href: "#"},
							{Icon: "task", Label: "Tasks", Href: "#"},
							{Icon: "note", Label: "Notes", Href: "#"},
							{Icon: "file", Label: "Files", Href: "#"},
						},
						On: 4,
					},
				},
			},
		},
	}
}

// Phone8 is frame «Рейл из значков» (phone-8).
func Phone8() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 4,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedPart,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined notebook",
							Meta:    "6",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "3"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "3"},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#3F6E8C"},
							Label:  "Общие с семьёй",
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "—", Hidden: true},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New notebook"},
						w.Button{Icon: "server", Title: "Connect account"},
						w.Button{Icon: "link", Title: "Subscribe to a feed"},
						w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined notebook"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 of 3 notebooks · 4 notes · VJOURNAL")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New note", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter by word", Short: "Filter", Label: "Filter by word"},
								Waiting: true,
								Variant: w.FilterH26W200,
							},
							w.Seg{
								Options:   []w.SegOption{{Label: "Newest", On: true}, {Label: "Oldest"}, {Label: "Title"}, {Label: "Recently changed"}},
								Secondary: true,
							},
							w.Button{Label: "Columns", Icon: "cols"},
						},
						Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "4"}}, Secondary: true},
					},
					w.ListGroup{Title: w.Rubric{Text: "August 2026", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
					w.List{
						Rows: []w.ListRow{
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Title:   []w.Inline{w.Text("Заметка со встречи 10.08")},
								Strong:  true,
								Excerpt: "Договорились заморозить ветку в среду; письмо пользователям пишет Анна…",
								Date:    "10 Aug",
								Source:  "nextcloud · Записи",
								On:      true,
							},
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Title:   []w.Inline{w.Text("Что почитать про CRDT")},
								Excerpt: "Кляйн, глава 5; статья Шапиро 2011; демка automerge…",
								Date:    "8 Aug",
								Source:  "nextcloud · Записи",
							},
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#6B4E7A"},
								Title:   []w.Inline{w.Text("Про дорогу до дачи")},
								Excerpt: "Съезд после заправки, потом грунтовка два километра…",
								Date:    "3 Aug",
								Source:  "nextcloud · Дневник",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "July 2026", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Title:   []w.Inline{w.Text("Про сроки по сборнику")},
								Excerpt: "Анна просит две недели на вёрстку после сдачи текста.",
								Date:    "21 Jul",
								Source:  "nextcloud · Записи",
							},
						},
					},
				},
			},
		},
	}
}

// Phone9 is frame «Рейл из значков · источники открыты» (phone-9).
func Phone9() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 4,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedPart,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined notebook",
							Meta:    "6",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "3"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "3"},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#3F6E8C"},
							Label:  "Общие с семьёй",
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "—", Hidden: true},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New notebook"},
						w.Button{Icon: "server", Title: "Connect account"},
						w.Button{Icon: "link", Title: "Subscribe to a feed"},
						w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined notebook"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 of 3 notebooks · 4 notes · VJOURNAL")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New note", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "August 2026", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
					w.List{
						Rows: []w.ListRow{
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Title:   []w.Inline{w.Text("Заметка со встречи 10.08")},
								Strong:  true,
								Excerpt: "Договорились заморозить ветку в среду; письмо пользователям пишет Анна…",
								Date:    "10 Aug",
								Source:  "nextcloud · Записи",
								On:      true,
							},
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Title:   []w.Inline{w.Text("Что почитать про CRDT")},
								Excerpt: "Кляйн, глава 5; статья Шапиро 2011; демка automerge…",
								Date:    "8 Aug",
								Source:  "nextcloud · Записи",
							},
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#6B4E7A"},
								Title:   []w.Inline{w.Text("Про дорогу до дачи")},
								Excerpt: "Съезд после заправки, потом грунтовка два километра…",
								Date:    "3 Aug",
								Source:  "nextcloud · Дневник",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "July 2026", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Title:   []w.Inline{w.Text("Про сроки по сборнику")},
								Excerpt: "Анна просит две недели на вёрстку после сдачи текста.",
								Date:    "21 Jul",
								Source:  "nextcloud · Записи",
							},
						},
					},
					w.Scrim{
						Content: w.Sheet{
							Title: "Sources",
							Head:  []w.SheetHeadPart{w.Button{Label: "Apply", Small: true, Place: w.Pushed}, w.Close{Label: "Close"}},
							Body: []w.SheetPart{
								w.RailSec{
									Items: []w.RailSecPart{
										w.TickRow{
											Ticked:  w.TickedPart,
											Kind:    w.TickRowAll,
											Mark:    w.Mark{Kind: w.MarkMulti},
											Label:   "Combined notebook",
											Meta:    "6",
											Current: true,
										},
									},
								},
								w.RailSec{
									Items: []w.RailSecPart{
										w.Rubric{Text: "nextcloud"},
										w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "3"},
										w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "3"},
										w.TickRow{
											Ticked: w.TickedOff,
											Mark:   w.Mark{Colour: "#3F6E8C"},
											Label:  "Общие с семьёй",
											Meta:   "—",
											Hidden: true,
											Trail:  w.Badge{Label: "ro"},
										},
										w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "—", Hidden: true},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
