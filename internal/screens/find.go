// Раздел макета: 7.6 Поиск по всем источникам.

package screens

import w "carrel.dev/typecase/widget"

// Find1 is frame «Поиск по всем источникам» (s-find-1).
func Find1() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "ковал", Label: "Search everything…"}, Clear: true},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.Button{
					Label: "2 min ago",
					Icon:  "refresh",
					Title: "Search again",
					Href:  "/app/search?q=%D0%BA%D0%BE%D0%B2%D0%B0%D0%BB",
				},
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
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Sources", Items: []w.RubricPart{w.Mono{Text: "14 / 14", Variant: w.MonoTight}}},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Error: true},
						w.TickRow{Kind: w.TickRowExternal, Mark: w.Mark{Colour: "#8C6239"}, Label: "Праздники РФ", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Files", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Documents", Meta: "2"},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "«ковал»"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("7 records")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Seg{
									Options: []w.SegOption{
										{Label: "All", On: true},
										{Label: "Contacts"},
										{Label: "Events"},
										{Label: "Tasks"},
										{Label: "Notes"},
										{Label: "Files"},
									},
								},
							},
						},
					},
					w.Runbar{
						Result: w.RunFailed,
						Parts: []w.RunbarPart{
							w.Say{Full: "Searched · 13 of 14 answered · 7 records ·", Short: "13 of 14 · 7 records ·"},
							w.Drop{
								Trigger: w.Button{Label: "1 failed", Menu: true, Kind: w.ButtonDanger, Small: true},
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuRefusal{
											Mark:   w.Mark{Colour: "#9A9280", Size: w.MarkH11},
											Source: "baikal · Общий",
											Answer: "REPORT /dav.php/calendars/mix/obshchij/ → no answer in 10 s; the search went on without it.",
											Retry:  w.Button{Label: "Retry", Small: true},
										},
									},
									Width: w.MenuWidth420,
									Pop:   w.MenuUnderEnd,
								},
							},
							w.Mono{Text: "just now", Place: w.Pushed},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Contacts", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#4A6B52"},
								Avatar:   w.Avatar{Initials: "АК"},
								Title:    []w.Inline{w.Text("Анна "), w.Strong("Ковал"), w.Text("ёва")},
								Subtitle: "+7 912 445-10-88",
								Detail:   "a.kovaleva@example.org",
								Caption:  "Издательство «Слово»",
								Source:   "nextcloud · Контакты",
							},
							w.FaceRow{
								Mark:     w.Mark{Colour: "#6B4E7A"},
								Avatar:   w.Avatar{Initials: "ПК"},
								Title:    []w.Inline{w.Text("Пётр "), w.Strong("Ковал"), w.Text("ёв")},
								Subtitle: "—",
								Detail:   "p.kovalev@example.org",
								Caption:  "—",
								Source:   "radicale · vps · Команда",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Events", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "10 Aug",
								Title:  []w.Inline{w.Text("Звонок с "), w.Strong("Ковал"), w.Text("ёвой")},
								Detail: "15:00 · Zoom",
								Source: "nextcloud · Личный",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Tasks", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.DueRow{
								Tick:     w.Tick{Sign: true},
								Mark:     w.Mark{Colour: "#3F6E8C"},
								Title:    []w.Inline{w.Text("Ответить "), w.Strong("Ковал"), w.Text("ёвой по смете")},
								Due:      "12 Aug",
								Progress: "40%",
								Source:   "nextcloud · Работа",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Notes", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Title:   []w.Inline{w.Text("Звонок с "), w.Strong("Ковал"), w.Text("ёвой — тезисы")},
								Excerpt: "Договорились заморозить ветку в среду; письмо пишет Анна…",
								Date:    "10 Aug",
								Source:  "nextcloud · Записи",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Files", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.SizeRow{
								Mark:   &w.Mark{Colour: "#4A6B52"},
								Icon:   "doc",
								Name:   []w.Inline{w.Text("договор-"), w.Strong("ковал"), w.Text("ёва.pdf")},
								Bytes:  "240 KB",
								Time:   "2026-08-04 11:20",
								Access: "rw",
								Source: "sftpgo · vps · Documents",
							},
							w.SizeRow{
								Mark:   &w.Mark{Colour: "#4A6B52"},
								Icon:   "doc",
								Name:   []w.Inline{w.Strong("Ковал"), w.Text("ёва — правки по обложке.md")},
								Bytes:  "6 KB",
								Time:   "2026-07-21 16:02",
								Access: "rw",
								Source: "sftpgo · vps · Documents",
							},
						},
					},
					w.Msg{Parts: []w.MsgPart{w.Text("Files are found by their names; what is inside them is not searched.")}},
				},
			},
		},
	}
}

// Find4 is frame «Поиск идёт» (s-find-4).
func Find4() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "ковал", Label: "Search everything…"}, Clear: true},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.Button{Label: "Refresh", Icon: "refresh", Title: "The search is still running.", State: w.ButtonDisabled},
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
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Sources", Items: []w.RubricPart{w.Mono{Text: "14 / 14", Variant: w.MonoTight}}},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий"},
						w.TickRow{Kind: w.TickRowExternal, Mark: w.Mark{Colour: "#8C6239"}, Label: "Праздники РФ"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Files"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Documents"},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "«ковал»"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("5 records so far")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Seg{
									Options: []w.SegOption{
										{Label: "All", On: true},
										{Label: "Contacts"},
										{Label: "Events"},
										{Label: "Tasks"},
										{Label: "Notes"},
										{Label: "Files"},
									},
								},
							},
						},
					},
					w.Runbar{
						Meter: w.Meter{Percent: 62},
						Parts: []w.RunbarPart{
							w.Say{Full: "Searching · 9 of 14 answered · 5 records", Short: "9 of 14 · 5 records"},
							w.Push{Items: []w.PushPart{w.Button{Label: "Stop", Small: true, Form: "search-stop"}}},
						},
					},
					w.PostForm{
						ID:     "search-stop",
						Action: "/app/search/cancel?q=%D0%BA%D0%BE%D0%B2%D0%B0%D0%BB",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
					w.ListGroup{Title: w.Rubric{Text: "Contacts", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#4A6B52"},
								Avatar:   w.Avatar{Initials: "АК"},
								Title:    []w.Inline{w.Text("Анна "), w.Strong("Ковал"), w.Text("ёва")},
								Subtitle: "+7 912 445-10-88",
								Detail:   "a.kovaleva@example.org",
								Caption:  "Издательство «Слово»",
								Source:   "nextcloud · Контакты",
							},
							w.FaceRow{
								Mark:     w.Mark{Colour: "#6B4E7A"},
								Avatar:   w.Avatar{Initials: "ПК"},
								Title:    []w.Inline{w.Text("Пётр "), w.Strong("Ковал"), w.Text("ёв")},
								Subtitle: "—",
								Detail:   "p.kovalev@example.org",
								Caption:  "—",
								Source:   "radicale · vps · Команда",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Events", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "10 Aug",
								Title:  []w.Inline{w.Text("Звонок с "), w.Strong("Ковал"), w.Text("ёвой")},
								Detail: "15:00 · Zoom",
								Source: "nextcloud · Личный",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Tasks", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.DueRow{
								Tick:     w.Tick{Sign: true},
								Mark:     w.Mark{Colour: "#3F6E8C"},
								Title:    []w.Inline{w.Text("Ответить "), w.Strong("Ковал"), w.Text("ёвой по смете")},
								Due:      "12 Aug",
								Progress: "40%",
								Source:   "nextcloud · Работа",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Notes", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Title:   []w.Inline{w.Text("Звонок с "), w.Strong("Ковал"), w.Text("ёвой — тезисы")},
								Excerpt: "Договорились заморозить ветку в среду; письмо пишет Анна…",
								Date:    "10 Aug",
								Source:  "nextcloud · Записи",
							},
						},
					},
				},
			},
		},
	}
}

// Find2 is frame «Подсказки: история запросов» (s-find-2).
func Find2() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{
				Query:   w.Query{Placeholder: "Search everything…", Label: "Search everything…"},
				Focus:   true,
				Variant: w.SearchFull,
			},
			Suggest: &w.Menu{
				Items: []w.MenuPart{
					w.MenuHeading{Rubric: w.Rubric{Text: "Recent searches", Form: w.RubricItem}},
					w.MenuItem{
						Icon:  "clock",
						Trail: w.Mono{Text: "7 found · 2 min ago", Variant: w.Mono11DustyTight, Place: w.Pushed},
						Label: "ковал",
						Href:  "/app/search?q=%D0%BA%D0%BE%D0%B2%D0%B0%D0%BB",
					},
					w.MenuItem{
						Icon:  "clock",
						Trail: w.Mono{Text: "3 found · yesterday", Variant: w.Mono11DustyTight, Place: w.Pushed},
						Label: "смета ковалёвой",
						Href:  "/app/search?q=%D1%81%D0%BC%D0%B5%D1%82%D0%B0+%D0%BA%D0%BE%D0%B2%D0%B0%D0%BB%D1%91%D0%B2%D0%BE%D0%B9",
					},
					w.MenuItem{
						Icon:  "clock",
						Trail: w.Mono{Text: "nothing found · Monday", Variant: w.Mono11DustyTight, Place: w.Pushed},
						Label: "дача",
						Href:  "/app/search?q=%D0%B4%D0%B0%D1%87%D0%B0",
					},
					w.MenuItem{
						Icon:  "clock",
						Trail: w.Mono{Text: "no source answered · Sunday", Variant: w.Mono11DustyTight, Place: w.Pushed},
						Label: "договор",
						Href:  "/app/search?q=%D0%B4%D0%BE%D0%B3%D0%BE%D0%B2%D0%BE%D1%80",
					},
					w.MenuSep{},
					w.MenuItem{Trail: w.Kbd{Keys: "Enter"}, Label: "Type to search everything", Tone: w.MenuToneDusty},
				},
				Width: w.MenuWidthFull,
			},
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
					On: 1,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedPart,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined address book",
							Meta:    "10",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "2", Trail: w.Badge{Label: "ro"}},
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 7 cards")}}},
						Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "New contact", Icon: "plus", Kind: w.ButtonPrimary}}},
					},
					w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#4A6B52"},
								Avatar:   w.Avatar{Initials: "АС"},
								Title:    []w.Inline{w.Text("Алексей Смирнов")},
								Subtitle: "+7 916 200-71-15",
								Detail:   "—",
								Caption:  "—",
								Source:   "nextcloud · Контакты",
							},
							w.FaceRow{
								Mark:       w.Mark{Colour: "#4A6B52"},
								Avatar:     w.Avatar{Initials: "АК"},
								Title:      []w.Inline{w.Text("Анна Ковалёва")},
								Strong:     true,
								TitleBadge: &w.Badge{Label: "linked ×3", Tone: w.BadgeLinked},
								Subtitle:   "+7 912 445-10-88",
								Detail:     "a.kovaleva@example.org",
								Caption:    "Издательство «Слово»",
								Source:     "3 sources",
								SourceHint: "nextcloud · Контакты, nextcloud · Архив, radicale · vps · Команда",
								On:         true,
							},
							w.FaceRow{
								Mark:       w.Mark{Colour: "#6B4E7A"},
								Avatar:     w.Avatar{Photo: true},
								Title:      []w.Inline{w.Text("Артём Белов")},
								TitleBadge: &w.Badge{Label: "linked ×2", Tone: w.BadgeLinked},
								Subtitle:   "+7 903 118-22-04",
								Detail:     "a.belov@example.org",
								Caption:    "Тензор",
								Source:     "2 sources",
								SourceHint: "nextcloud · Контакты, radicale · vps · Команда",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "В", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#A05A2C"},
								Avatar:   w.Avatar{Initials: "ВЛ"},
								Title:    []w.Inline{w.Text("Вера Лебедева")},
								Subtitle: "+7 905 331-09-42",
								Detail:   "vera@example.net",
								Caption:  "—",
								Source:   "nextcloud · Архив",
							},
							w.FaceRow{
								Mark:     w.Mark{Colour: "#4A6B52"},
								Avatar:   w.Avatar{Initials: "ВГ"},
								Title:    []w.Inline{w.Text("Виктор Гончаров")},
								Subtitle: "—",
								Detail:   "v.goncharov@example.org",
								Caption:  "Городская библиотека №4",
								Source:   "nextcloud · Контакты",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Л", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#4A6B52"},
								Avatar:   w.Avatar{Photo: true},
								Title:    []w.Inline{w.Text("Лебедева Вера")},
								Subtitle: "+7 905 331-09-42",
								Detail:   "v.lebedeva@example.org",
								Caption:  "Библиотека №4",
								Source:   "nextcloud · Контакты",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "М", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#4A6B52"},
								Avatar:   w.Avatar{Initials: "МЕ"},
								Title:    []w.Inline{w.Text("Михаил Енученко")},
								Subtitle: "—",
								Detail:   "mix@example.org",
								Caption:  "—",
								Source:   "nextcloud · Контакты",
							},
						},
					},
				},
			},
			Height: w.ShellHeight340,
		},
	}
}

// Find3 is frame «Подсказки по мере набора» (s-find-3).
func Find3() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "ковал", Label: "Search everything…"}, Focus: true, Variant: w.SearchFull},
			Suggest: &w.Menu{
				Items: []w.MenuPart{
					w.MenuHeading{Rubric: w.Rubric{Text: "Jump to", Form: w.RubricItem}},
					w.MenuItem{
						Icon:  "con",
						Mark:  &w.Mark{Colour: "#4A6B52"},
						Found: []w.Inline{w.Text("Анна "), w.Strong("Ковал"), w.Text("ёва")},
						Trail: w.Dim{
							Parts:   []w.Inline{w.Text("nextcloud · Контакты")},
							Size:    w.Dim11,
							Variant: w.DimTight,
							Place:   w.Pushed,
						},
					},
					w.MenuItem{
						Icon:  "cal",
						Mark:  &w.Mark{Colour: "#4A6B52"},
						Found: []w.Inline{w.Text("Звонок с "), w.Strong("Ковал"), w.Text("ёвой")},
						Trail: w.Dim{
							Parts:   []w.Inline{w.Text("10 Aug · nextcloud · Личный")},
							Size:    w.Dim11,
							Variant: w.DimTight,
							Place:   w.Pushed,
						},
					},
					w.MenuItem{
						Icon:  "file",
						Mark:  &w.Mark{Colour: "#4A6B52"},
						Found: []w.Inline{w.Text("договор-"), w.Strong("ковал"), w.Text("ёва.pdf")},
						Trail: w.Dim{
							Parts:   []w.Inline{w.Text("sftpgo · vps · Documents")},
							Size:    w.Dim11,
							Variant: w.DimTight,
							Place:   w.Pushed,
						},
					},
					w.MenuSep{},
					w.MenuItem{Icon: "search", Trail: w.Kbd{Keys: "Enter"}, Label: "Show all results for «ковал»", On: true},
					w.MenuItem{
						Icon:  "all",
						Found: []w.Inline{w.Text("Timeline · Анна "), w.Strong("Ковал"), w.Text("ёва")},
						Trail: w.Dim{Parts: []w.Inline{w.Text("asks every server")}, Size: w.Dim11, Variant: w.DimTight, Place: w.Pushed},
					},
					w.MenuSep{},
					w.MenuHeading{Rubric: w.Rubric{Text: "Recent searches", Form: w.RubricItem}},
					w.MenuItem{
						Icon:  "clock",
						Found: []w.Inline{w.Text("смета "), w.Strong("ковал"), w.Text("ёвой")},
						Trail: w.Mono{Text: "3 found · yesterday", Variant: w.Mono11DustyTight, Place: w.Pushed},
					},
				},
				Width: w.MenuWidthFull,
			},
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
					On: 2,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedPart,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined calendar",
							Meta:    "7",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "3"},
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
							Meta:  "3",
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
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New calendar"},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined calendar"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{
									w.Text("5 of 7 calendars · 3 accounts · 10 – 16 August 2026 · "),
									w.Mono{Text: "Europe/Moscow"},
									w.Text(" · 7 events"),
								},
							},
						},
						Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}}},
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#3F6E8C"},
								Time:   "all day",
								Title:  []w.Inline{w.Text("Дежурство: первая линия")},
								Detail: "—",
								Badges: []w.Badge{w.Badge{Label: "repeats weekly"}},
								Source: "nextcloud · Работа",
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#8C6239"},
								Time:   "all day",
								Title:  []w.Inline{w.Text("Дежурство: первая линия")},
								Detail: "—",
								Badges: []w.Badge{w.Badge{Label: "repeats weekly"}},
								Source: "radicale · vps · Дежурства",
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "10:00",
								Title:  []w.Inline{w.Text("Синк-ап по релизу")},
								Detail: "Переговорная 2",
								Badges: []w.Badge{w.Badge{Label: "4 attendees"}, w.Badge{Label: "repeats weekly"}},
								Source: "nextcloud · Личный",
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#3F6E8C"},
								Time:   "15:00",
								Title:  []w.Inline{w.Text("Созвон с подрядчиком")},
								Detail: "Jitsi",
								Badges: []w.Badge{w.Badge{Label: "tentative", Tone: w.BadgeAlert}},
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
								Badges: []w.Badge{w.Badge{Label: "2 attendees"}},
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
					w.ListGroup{Title: w.Rubric{Text: "Wednesday · 12 August", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "all day",
								Title:  []w.Inline{w.Text("Поездка в Казань")},
								Detail: "—",
								Badges: []w.Badge{w.Badge{Label: "day 1 of 3"}},
								Source: "nextcloud · Личный",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Thursday · 13 August", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "all day",
								Title:  []w.Inline{w.Text("Поездка в Казань")},
								Detail: "—",
								Badges: []w.Badge{w.Badge{Label: "day 2 of 3"}},
								Source: "nextcloud · Личный",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Friday · 14 August", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "all day",
								Title:  []w.Inline{w.Text("Поездка в Казань")},
								Detail: "—",
								Badges: []w.Badge{w.Badge{Label: "day 3 of 3"}},
								Source: "nextcloud · Личный",
							},
						},
					},
				},
			},
			Height: w.ShellHeight340,
		},
	}
}
