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
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Sources", Items: []w.RubricPart{w.Mono{Text: "14 / 14", Variant: w.MonoTight}}},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "0"},
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("14 of 14 collections · 7 records · 1 source timed out")}}},
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
						Meter: w.Meter{Percent: 100, Alert: true},
						Parts: []w.RunbarPart{
							w.Text("Done · 13 of 14 sources answered · "),
							w.Drop{
								Trigger: w.Button{Label: "1 failed", Menu: true, Kind: w.ButtonDanger, Small: true},
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuRefusal{
											Mark:   w.Mark{Colour: "#9A9280", Size: w.MarkH11},
											Source: "baikal · Общий",
											Answer: "REPORT /dav.php/calendars/mix/obshchiy/ → no answer in 10 s; the search went on without it.",
											Retry:  w.Button{Label: "Retry", Small: true},
										},
									},
									Width: w.MenuWidth420,
									Pop:   w.MenuUnderEnd,
								},
							},
							w.Mono{Text: "1.8 s"},
							w.Line{
								Items: []w.LinePart{w.Mono{Text: "answered 11:42"}, w.Button{Label: "Search again", Icon: "refresh", Small: true}},
								Gap:   w.LineGap8,
								Place: w.Pushed,
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Contacts", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#4A6B52"},
								Avatar:   w.Avatar{Initials: "АК"},
								Title:    []w.Inline{w.Strong("Ковал"), w.Text("ёва Анна")},
								Subtitle: "+7 912 445-10-88",
								Detail:   "a.kovaleva@example.org",
								Caption:  "Издательство «Слово»",
								Source:   "nextcloud · Контакты",
							},
							w.FaceRow{
								Mark:     w.Mark{Colour: "#6B4E7A"},
								Avatar:   w.Avatar{Initials: "ПК"},
								Title:    []w.Inline{w.Strong("Ковал"), w.Text("ёв Пётр")},
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
								Time:   "15:00",
								Title:  []w.Inline{w.Text("Звонок с "), w.Strong("Ковал"), w.Text("ёвой")},
								Detail: "Zoom",
								Source: "nextcloud · Личный",
								Short:  true,
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Tasks", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.DueRow{
								Tick:     w.Tick{Label: "Done"},
								Mark:     w.Mark{Colour: "#3F6E8C"},
								Title:    []w.Inline{w.Text("Ответить "), w.Strong("Ковал"), w.Text("ёвой по смете")},
								Tags:     []w.RowTag{w.Tag{Label: "релиз"}},
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
								Tags:    []w.RowTag{w.Tag{Label: "встречи"}},
								Date:    "10 Aug",
								Source:  "nextcloud · Записи",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Files", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.SizeRow{
								Ticked: w.TickedOff,
								Icon:   "doc",
								Name:   []w.Inline{w.Text("договор-"), w.Strong("ковал"), w.Text("ёва.pdf")},
								Bytes:  "240 KB",
								Time:   "2026-08-04 11:20",
								Access: "rw",
								Source: "sftpgo · vps · Documents",
							},
							w.SizeRow{
								Ticked: w.TickedOff,
								Icon:   "doc",
								Name:   []w.Inline{w.Strong("Ковал"), w.Text("ёва — правки по обложке.md")},
								Bytes:  "6 KB",
								Time:   "2026-07-21 16:02",
								Access: "rw",
								Source: "sftpgo · vps · Documents",
							},
						},
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("Files are searched by name: "),
							w.Mono{Text: "PROPFIND"},
							w.Text(" returns names anyway, and walking a folder costs one request per folder. Carrel does not look inside files — and says so here, so that “nothing found” is not read as “not in the files”."),
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
						Trail: w.Mono{Text: "8 found · 2 min ago", Variant: w.Mono11DustyTight, Place: w.Pushed},
						Label: "ковал",
					},
					w.MenuItem{
						Icon:  "clock",
						Trail: w.Mono{Text: "3 found · yesterday", Variant: w.Mono11DustyTight, Place: w.Pushed},
						Label: "смета",
					},
					w.MenuItem{
						Icon:  "clock",
						Trail: w.Mono{Text: "nothing found · Monday", Variant: w.Mono11DustyTight, Place: w.Pushed},
						Label: "дача",
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
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
							Meta:    "312",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "248"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "51"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "13", Trail: w.Badge{Label: "ro"}},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined address book"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 address books · 312 cards")}}},
						Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "New contact", Icon: "plus", Kind: w.ButtonPrimary}}},
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
								Mark:     w.Mark{Colour: "#6B4E7A"},
								Avatar:   w.Avatar{Photo: true},
								Title:    []w.Inline{w.Text("Артём Белов")},
								Subtitle: "+7 903 118-22-04",
								Detail:   "a.belov@example.org",
								Caption:  "Тензор",
								Source:   "radicale · vps · Команда",
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
						Found: []w.Inline{w.Text("Звонок с "), w.Strong("Ковал"), w.Text("ёвой")},
						Trail: w.Dim{Parts: []w.Inline{w.Text("10 Aug · Личный")}, Size: w.Dim11, Variant: w.DimTight, Place: w.Pushed},
					},
					w.MenuItem{
						Icon:  "file",
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
						Trail: w.Dim{Parts: []w.Inline{w.Text("asks every server")}, Size: w.Dim11, Variant: w.DimTight, Place: w.Pushed},
						Label: "Timeline · Анна Ковалёва",
					},
					w.MenuSep{},
					w.MenuHeading{Rubric: w.Rubric{Text: "Recent searches", Form: w.RubricItem}},
					w.MenuItem{
						Icon:  "clock",
						Trail: w.Mono{Text: "3 found · yesterday", Variant: w.Mono11DustyTight, Place: w.Pushed},
						Label: "смета",
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
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
							Meta:    "14",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "7"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "4"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
					},
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("5 of 7 calendars · 10 – 16 August 2026")}}},
						Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}}},
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "10:00",
								Title:  []w.Inline{w.Text("Синк-ап по релизу")},
								Detail: "Переговорная 2",
								Source: "nextcloud · Личный",
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#3F6E8C"},
								Time:   "13:30",
								Title:  []w.Inline{w.Text("Звонок с Ковалёвой")},
								Detail: "Zoom",
								Source: "nextcloud · Работа",
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#6B4E7A"},
								Time:   "18:00",
								Title:  []w.Inline{w.Text("Забрать документы")},
								Detail: "МФЦ на Ленина",
								Source: "radicale · vps · Проекты",
							},
						},
					},
				},
			},
			Height: w.ShellHeight340,
		},
	}
}
