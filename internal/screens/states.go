// Раздел макета: 8. Состояния.

package screens

import w "carrel.dev/typecase/widget"

// States1 is frame «Ничего не подключено» (states-1).
func States1() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Href: "#"},
						{Icon: "cal", Label: "Calendar", Href: "#"},
						{Icon: "task", Label: "Tasks", Href: "#"},
						{Icon: "note", Label: "Notes", Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Href: "#"},
					},
					On: 2,
				},
				w.RailSec{Items: []w.RailSecPart{
					w.Rubric{Text: "Sources"}}, Note: "Nothing connected yet"},
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
						Title: &w.HeadTitle{Text: "Calendar"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("No server connected yet")}}},
					},
					w.Empty{
						Parts: []w.EmptyPart{
							w.Text("Carrel stores nothing of its own: calendars, contacts and notes come from your DAV servers. Connect the first one — you will need an address, a user name and a password. "),
							w.Button{Label: "Connect account", Icon: "plus", Kind: w.ButtonPrimary},
						},
					},
				},
			},
		},
	}
}

// States2 is frame «Файлы есть в навигации always» (states-2).
func States2() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"},
					},
					On: 5,
				},
				w.RailSec{Items: []w.RailSecPart{
					w.Rubric{Text: "Sources"}}, Note: "No folder on any connected server"},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New folder"},
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
						Title: &w.HeadTitle{Text: "Files"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("None of the connected servers offers a plain folder")}}},
					},
					w.Empty{
						Parts: []w.EmptyPart{
							w.Text("Carrel shows this section when discovery finds a WebDAV collection that is neither a calendar nor an address book. There is no setting to turn it on: Baikal serves calendars and contacts only — a folder usually means a second account, on Nextcloud or on a plain WebDAV server. "),
							w.Button{Label: "Connect account", Icon: "plus", Kind: w.ButtonPrimary},
						},
					},
				},
			},
		},
	}
}

// States3 is frame «Загрузка» (states-3).
func States3() Screen {
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
						{Icon: "con", Label: "Contacts", Href: "#"},
						{Icon: "cal", Label: "Calendar", Href: "#"},
						{Icon: "task", Label: "Tasks", Href: "#"},
						{Icon: "note", Label: "Notes", Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Href: "#"},
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
							Current: true,
							Trail:   w.Skel{Width: w.Skel16},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#A05A2C"},
							Label:  "Дни рождения",
							Hidden: true,
							Trail:  w.Skel{Width: w.Skel16},
						},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Trail: w.Skel{Width: w.Skel16}},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Trail: w.Skel{Width: w.Skel16}},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Trail: w.Skel{Width: w.Skel16}},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Trail: w.Skel{Width: w.Skel16}},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "baikal"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Trail: w.Skel{Width: w.Skel16}},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Subscriptions", Items: []w.RubricPart{w.Badge{Label: "external"}}},
						w.TickRow{Kind: w.TickRowExternal, Mark: w.Mark{Colour: "#8C6239"}, Label: "Праздники РФ", Meta: "ro"},
						w.TickRow{
							Ticked: w.TickedOff,
							Kind:   w.TickRowExternal,
							Mark:   w.Mark{Colour: "#6B4E7A"},
							Label:  "Расписание лиги",
							Meta:   "ro",
							Hidden: true,
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("5 of 7 calendars · 10 – 16 August 2026")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "up"},
								w.Button{Label: "Export", Icon: "down"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Button{Icon: "back", Title: "Previous week", Href: "/app/calendar?from=2026-08-03"},
							w.Button{Label: "Today", Href: "/app/calendar"},
							w.Button{Icon: "next", Title: "Next week", Href: "/app/calendar?from=2026-08-17"},
							w.Seg{
								Options: []w.SegOption{{Label: "Day"}, {Label: "Week", On: true}, {Label: "Month"}, {Label: "Range"}},
							},
							w.Range{
								From:      w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true},
								To:        w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-16", Label: "To", Mono: true},
								Show:      w.Button{Label: "Show"},
								Hidden:    []w.Hidden{{Name: "view", Value: "range"}},
								Secondary: true,
								Action:    "/app/calendar",
							},
							w.BarSep{Secondary: true},
							w.CheckForm{
								Checks: []w.Check{
									w.Check{On: true, Name: "show", Value: "events", Parts: []w.CheckPart{w.Text("Events")}},
									w.Check{Name: "show", Value: "tasks", Parts: []w.CheckPart{w.Text("Tasks due")}},
									w.Check{Name: "show", Value: "notes", Parts: []w.CheckPart{w.Text("Notes")}},
								},
								Action:   "/app/calendar",
								Tertiary: true,
								Apply:    w.Button{Label: "Apply", Small: true},
							},
							w.Drop{
								Trigger: w.Button{Icon: "dots", Title: "More"},
								More:    true,
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.CheckForm{
											Checks: []w.Check{
												w.Check{On: true, Name: "show", Value: "events", Parts: []w.CheckPart{w.Text("Events")}},
												w.Check{Name: "show", Value: "tasks", Parts: []w.CheckPart{w.Text("Tasks due")}},
												w.Check{Name: "show", Value: "notes", Parts: []w.CheckPart{w.Text("Notes")}},
											},
											Action:   "/app/calendar",
											Tertiary: true,
											Apply:    w.Button{Label: "Apply", Small: true},
										},
										w.MenuSep{Tertiary: true},
										w.Range{
											From:   w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true},
											To:     w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-16", Label: "To", Mono: true},
											Show:   w.Button{Label: "Show"},
											Hidden: []w.Hidden{{Name: "view", Value: "range"}},
											Action: "/app/calendar",
											Wrap:   true,
										},
										w.MenuSep{},
										w.CheckForm{
											Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}},
											Action: "/app/calendar",
											Apply:  w.Button{Label: "Apply", Small: true},
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
									Action: "/app/calendar",
									Apply:  w.Button{Label: "Apply", Small: true},
								},
							},
							Secondary: true,
						},
						Variant: w.BarWide,
					},
					w.Runbar{
						Meter: w.Meter{Percent: 22},
						Parts: []w.RunbarPart{w.Text("Polling 6 sources…"), w.Mono{Text: "0.4 s", Place: w.Pushed}},
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}},
					w.List{
						Rows: []w.ListRow{
							w.SkelRow{Title: w.Skel{Width: w.Skel60}},
							w.SkelRow{Title: w.Skel{Width: w.Skel45}},
							w.SkelRow{Title: w.Skel{Width: w.Skel52}},
						},
					},
				},
			},
		},
	}
}

// States4 is frame «Только для чтения» (states-4).
func States4() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"},
					},
					On: 1,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked: w.TickedPart,
							Kind:   w.TickRowAll,
							Mark:   w.Mark{Kind: w.MarkMulti},
							Label:  "Combined address book",
							Meta:   "312",
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
						w.TickRow{
							Mark:    w.Mark{Colour: "#6B4E7A"},
							Label:   "Команда",
							Meta:    "13",
							Current: true,
							Trail:   w.Badge{Label: "ro"},
						},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#8C6239"}, Label: "Рассылки", Meta: "—", Hidden: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "работа"}, w.Tag{Label: "издательство"}},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New address book"},
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
						Title: &w.HeadTitle{Text: "Команда", Badge: &w.Badge{Label: "read-only"}},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("radicale · vps · the server granted no write access")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{w.Button{Label: "Export", Icon: "down"}, w.Button{Label: "Print", Icon: "print"}},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter the loaded cards", Label: "Filter the loaded cards"},
								Waiting: true,
								Variant: w.FilterH26W220,
							},
							w.Seg{Options: []w.SegOption{{Label: "A–Z", On: true}, {Label: "Recently changed"}}, Secondary: true},
							w.Button{Label: "Columns", Icon: "cols"},
						},
						More: &w.Button{Icon: "dots", Title: "More"},
						Right: &w.BarRight{
							Items: []w.BarPart{
								w.CheckForm{
									Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}},
									Action: "/app/contacts",
									Apply:  w.Button{Label: "Apply", Small: true},
								},
								w.Mono{Text: "13"},
							},
							Secondary: true,
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "И", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#6B4E7A"},
								Avatar:   w.Avatar{Initials: "ИФ"},
								Title:    []w.Inline{w.Text("Ирина Фомина")},
								Subtitle: "+7 900 000-00-00",
								Detail:   "—",
								Caption:  "—",
								Source:   "radicale · vps · Команда",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "К", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#6B4E7A"},
								Avatar:   w.Avatar{Initials: "ПК"},
								Title:    []w.Inline{w.Text("Пётр Ковалёв")},
								Subtitle: "—",
								Detail:   "p.kovalev@example.org",
								Caption:  "—",
								Source:   "radicale · vps · Команда",
							},
						},
					},
				},
			},
		},
	}
}

// States5 is frame «Нет сети» (states-5).
func States5() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Href: "#"},
						{Icon: "cal", Label: "Calendar", Href: "#"},
						{Icon: "task", Label: "Tasks", Href: "#"},
						{Icon: "note", Label: "Notes", Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Href: "#"},
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
							Meta:    "—",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Error: true},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Error: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Error: true},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Runbar{
						Meter: w.Meter{Percent: 0},
						Parts: []w.RunbarPart{
							w.Text("No connection to carrel.example · last successful attempt 4 minutes ago "),
							w.Push{Items: []w.PushPart{w.Button{Label: "Retry", Small: true}}},
						},
						Failed: true,
					},
					w.Head{
						Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined calendar"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("shown from this session’s memory")}}},
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
								Past:   true,
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#3F6E8C"},
								Time:   "15:00",
								Title:  []w.Inline{w.Text("Созвон с подрядчиком")},
								Detail: "Jitsi",
								Source: "nextcloud · Работа",
								Past:   true,
							},
						},
					},
					w.Empty{
						Parts: []w.EmptyPart{
							w.Text("Creating and editing are not possible now: writing goes to the server, and the server is not there. The buttons are not hidden — they say the same if pressed."),
						},
						Variant: w.EmptySmall,
					},
				},
			},
		},
	}
}

// States6 is frame «Источник не ответил» (states-6).
func States6() Screen {
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
						{Icon: "con", Label: "Contacts", Href: "#"},
						{Icon: "cal", Label: "Calendar", Href: "#"},
						{Icon: "task", Label: "Tasks", Href: "#"},
						{Icon: "note", Label: "Notes", Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Href: "#"},
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
							Meta:    "2",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "1"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "0"},
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
						w.TickRow{Kind: w.TickRowExternal, Mark: w.Mark{Colour: "#8C6239"}, Label: "Праздники РФ", Meta: "ro"},
						w.TickRow{
							Ticked: w.TickedOff,
							Kind:   w.TickRowExternal,
							Mark:   w.Mark{Colour: "#6B4E7A"},
							Label:  "Расписание лиги",
							Meta:   "ro",
							Hidden: true,
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("6 of 7 calendars · one source failed")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "up"},
								w.Button{Label: "Export", Icon: "down"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Button{Icon: "back", Title: "Previous week", Href: "/app/calendar?from=2026-08-03"},
							w.Button{Label: "Today", Href: "/app/calendar"},
							w.Button{Icon: "next", Title: "Next week", Href: "/app/calendar?from=2026-08-17"},
							w.Seg{
								Options: []w.SegOption{{Label: "Day"}, {Label: "Week", On: true}, {Label: "Month"}, {Label: "Range"}},
							},
							w.Range{
								From:      w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true},
								To:        w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-16", Label: "To", Mono: true},
								Show:      w.Button{Label: "Show"},
								Hidden:    []w.Hidden{{Name: "view", Value: "range"}},
								Secondary: true,
								Action:    "/app/calendar",
							},
							w.BarSep{Secondary: true},
							w.CheckForm{
								Checks: []w.Check{
									w.Check{On: true, Name: "show", Value: "events", Parts: []w.CheckPart{w.Text("Events")}},
									w.Check{Name: "show", Value: "tasks", Parts: []w.CheckPart{w.Text("Tasks due")}},
									w.Check{Name: "show", Value: "notes", Parts: []w.CheckPart{w.Text("Notes")}},
								},
								Action:   "/app/calendar",
								Tertiary: true,
								Apply:    w.Button{Label: "Apply", Small: true},
							},
							w.Drop{
								Trigger: w.Button{Icon: "dots", Title: "More"},
								More:    true,
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.CheckForm{
											Checks: []w.Check{
												w.Check{On: true, Name: "show", Value: "events", Parts: []w.CheckPart{w.Text("Events")}},
												w.Check{Name: "show", Value: "tasks", Parts: []w.CheckPart{w.Text("Tasks due")}},
												w.Check{Name: "show", Value: "notes", Parts: []w.CheckPart{w.Text("Notes")}},
											},
											Action:   "/app/calendar",
											Tertiary: true,
											Apply:    w.Button{Label: "Apply", Small: true},
										},
										w.MenuSep{Tertiary: true},
										w.Range{
											From:   w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true},
											To:     w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-16", Label: "To", Mono: true},
											Show:   w.Button{Label: "Show"},
											Hidden: []w.Hidden{{Name: "view", Value: "range"}},
											Action: "/app/calendar",
											Wrap:   true,
										},
										w.MenuSep{},
										w.CheckForm{
											Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}},
											Action: "/app/calendar",
											Apply:  w.Button{Label: "Apply", Small: true},
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
									Action: "/app/calendar",
									Apply:  w.Button{Label: "Apply", Small: true},
								},
							},
							Secondary: true,
						},
						Variant: w.BarWide,
					},
					w.Runbar{
						Meter: w.Meter{Percent: 83, Alert: true},
						Parts: []w.RunbarPart{
							w.Text("Polled 5 of 6 sources · 2 events · "),
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
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
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
								Time:   "15:00",
								Title:  []w.Inline{w.Text("Созвон с подрядчиком")},
								Detail: "Jitsi",
								Source: "nextcloud · Работа",
							},
						},
					},
					w.Empty{
						Parts: []w.EmptyPart{
							w.Text("What arrived in time is shown. Records from the unavailable source are missing — not “there are none”, but “we have not seen them”."),
						},
						Variant: w.EmptySmall,
					},
				},
			},
		},
	}
}

// States17 is frame «Событие · сервер не сохранил свойства» (states-17).
func States17() Screen {
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
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to the card", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "7", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "4"},
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
						Title: &w.HeadTitle{Text: "Синк-ап по релизу"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Личный")}}},
						Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Save", Kind: w.ButtonPrimary}}},
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Strong("The change was saved, but the server did not keep these properties."),
							w.Text(" Other clients will not see them either: "),
							w.Mono{Text: "X-CALENDARSERVER-PRIVATE-COMMENT"},
							w.Text(", "),
							w.Mono{Text: "X-MICROSOFT-CDO-BUSYSTATUS"},
							w.Text("."),
						},
						Space: w.MsgFlushTop2,
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Summary",
								Parts: []w.FieldPart{w.Input{Value: "Синк-ап по релизу", Label: "Summary", Focus: true}},
								Width: w.FieldXL,
							},
							w.Field{
								Label: "Location",
								Parts: []w.FieldPart{w.Input{Value: "Переговорная 2", Label: "Location"}},
								Width: w.FieldLG,
							},
							w.Field{Label: "Status", Parts: []w.FieldPart{w.Pick{Value: "Confirmed"}}, Width: w.FieldSM},
							w.Field{
								Label: "Tags",
								Parts: []w.FieldPart{w.Input{Value: "release, team", Label: "Tags"}},
								Width: w.FieldMD,
							},
							w.Field{
								Label: "Description",
								Parts: []w.FieldPart{
									w.Area{Text: "Обсудить оставшиеся блокеры, дату заморозки и кто пишет письмо пользователям.", Label: "Text"},
								},
								Width: w.FieldFull,
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "When"},
								Items: []w.FieldSetPart{
									w.Field{Parts: []w.FieldPart{w.Check{Parts: []w.CheckPart{w.Text("All day")}}}},
									w.Field{
										Label: "Start date",
										Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Value: "2026-08-10", Label: "Start date", Mono: true}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Start time",
										Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Value: "12:00", Label: "Start time", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "End date",
										Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Value: "2026-08-10", Label: "End date", Mono: true}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "End time",
										Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Value: "13:00", Label: "End time", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Time zone",
										Parts: []w.FieldPart{
											w.Drop{
												Trigger: w.Pick{
													Value: "Asia/Yekaterinburg",
													Aside: &w.Dim{Parts: []w.Inline{w.Text("as sent")}, Variant: w.DimIndent},
												},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.Input{Kind: w.InputSearch, Placeholder: "Find a zone", Label: "Find a zone"},
														w.Rubric{Text: "As sent", Form: w.RubricInline},
														w.MenuItem{Label: "Asia/Yekaterinburg", On: true, Name: "tz", Value: "Asia/Yekaterinburg"},
														w.MenuSep{},
														w.MenuItem{Label: "Europe/Moscow", Name: "tz", Value: "Europe/Moscow"},
														w.MenuItem{Label: "UTC", Name: "tz", Value: "UTC"},
														w.MenuSep{},
														w.MenuItem{Label: "Floating — the same clock time anywhere", Name: "tz", Value: "floating"},
													},
													Width: w.MenuWidth290,
													Pop:   w.MenuUnder,
												},
											},
										},
										Width: w.FieldMD,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Reminders"},
								Items: []w.FieldSetPart{
									w.Field{Label: "First", Parts: []w.FieldPart{w.Pick{Value: "10 minutes before"}}, Width: w.FieldMD},
									w.Field{Label: "Reminder", Parts: []w.FieldPart{w.Pick{Value: "1 day before"}}, Width: w.FieldMD},
									w.Field{Label: "\u00a0", Parts: []w.FieldPart{w.Button{Label: "Add reminder"}}},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Attendees"},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Drop{
												Trigger: w.TagInput{
													Parts: []w.TagInputPart{
														w.Tag{
															Label: "Мария Лебедева",
															Mark:  &w.Mark{Colour: "#6B4E7A", Size: w.MarkH11},
															Fixed: true,
															Aside: &w.Dim{Parts: []w.Inline{w.Text("· organiser")}},
															Title: "m.lebedeva@example.org",
															Name:  "attendee",
															Value: "mailto:m.lebedeva@example.org",
														},
														w.Tag{
															Label:     "Анна Ковалёва",
															Mark:      &w.Mark{Colour: "#4A6B52", Size: w.MarkH11},
															Removable: true,
															Aside:     &w.Dim{Parts: []w.Inline{w.Text("· accepted")}},
															Title:     "a.kovaleva@example.org",
															Name:      "attendee",
															Value:     "mailto:a.kovaleva@example.org",
														},
														w.Tag{
															Label:     "Артём Белов",
															Mark:      &w.Mark{Colour: "#6B4E7A", Size: w.MarkH11},
															Removable: true,
															Aside:     &w.Dim{Parts: []w.Inline{w.Text("· tentative")}},
															Title:     "a.belov@example.org",
															Name:      "attendee",
															Value:     "mailto:a.belov@example.org",
														},
														w.Tag{
															Address:   &w.Mono{Text: "v@example.org"},
															Removable: true,
															Aside:     &w.Dim{Parts: []w.Inline{w.Text("· no reply")}},
															Title:     "v@example.org",
															Name:      "attendee",
															Value:     "mailto:v@example.org",
														},
													},
													Columns: true,
													Name:    "attendee",
													Add: &w.TagAdd{
														Name:        "attendee.add",
														Placeholder: "Name or address",
														Label:       "Add attendee",
														Suggest:     "/app/contacts/suggest",
													},
												},
												Menu: w.Menu{Items: []w.MenuPart{}, Width: w.MenuWidth420, Pop: w.MenuUnder},
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("A name or an address: the ticked address books are searched. On save, nextcloud sends the invitations and their changes — Carrel sends no mail itself."),
												},
											},
										},
										Width: w.FieldFull,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Attachments"},
								Items: []w.FieldSetPart{
									w.Line{
										Items: []w.LinePart{
											w.Drop{
												Trigger: w.Button{Label: "Attach", Menu: true, Icon: "up"},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{Icon: "up", Label: "From this computer"},
														w.MenuItem{Icon: "file", Label: "From Files"},
													},
													Width: w.MenuWidth212,
													Pop:   w.MenuUnder,
												},
											},
											w.Faced{
												Face: w.FaceLarge,
												Items: []w.LinePart{
													w.Text("release-checklist.md "),
													w.Dim{Parts: []w.Inline{w.Text("text/markdown · 3 KB")}},
													w.Text(" "),
													w.Badge{Label: "in the event"},
												},
											},
											w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true},
										},
										Gap:   w.LineGap8,
										Block: true,
										Place: w.Grown,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Repeats"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Frequency", Parts: []w.FieldPart{w.Pick{Value: "Weekly"}}, Width: w.FieldSM},
									w.Field{
										Label: "Every",
										Parts: []w.FieldPart{w.Input{Value: "1", Label: "Every", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "On",
										Parts: []w.FieldPart{
											w.Seg{
												Options: []w.SegOption{
													{Label: "MO", On: true, Value: "MO"},
													{Label: "TU", Value: "TU"},
													{Label: "WE", Value: "WE"},
													{Label: "TH", Value: "TH"},
													{Label: "FR", Value: "FR"},
													{Label: "SA", Value: "SA"},
													{Label: "SU", Value: "SU"},
												},
												Name:     "byday",
												Multiple: true,
											},
										},
									},
									w.Field{
										Label: "Until",
										Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Value: "2026-12-28", Label: "Until", Mono: true}},
										Width: w.FieldSM,
									},
									w.Field{Label: "Count", Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Count"}}, Width: w.FieldXS},
									w.Field{
										Label: "Advanced RRULE",
										Parts: []w.FieldPart{
											w.Input{Placeholder: "FREQ=WEEKLY;BYDAY=MO;UNTIL=20261228", Label: "Advanced RRULE", Mono: true},
											w.FieldHint{Parts: []w.Inline{w.Text("The advanced value takes precedence when set.")}},
										},
										Width: w.FieldXL,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Kept as the server sent them"},
								Items: []w.FieldSetPart{
									w.Diag{
										Text:  "X-MOZ-GENERATION  3\nX-APPLE-TRAVEL-ADVISORY-BEHAVIOR  AUTOMATIC\nSEQUENCE  2",
										Place: w.Grown,
									},
								},
							},
						},
					},
					w.Danger{
						Button: w.Button{Label: "Delete event", Icon: "trash", Kind: w.ButtonDangerSolid},
						Note:   "Series is edited and deleted whole.",
					},
				},
			},
		},
	}
}

// States7 is frame «Контакты · ни одна книга не отмечена» (states-7).
func States7() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"},
					},
					On: 1,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked: w.TickedOff,
							Kind:   w.TickRowAll,
							Mark:   w.Mark{Kind: w.MarkMulti},
							Label:  "Combined address book",
							Meta:   "312",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "248", Hidden: true},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "51", Hidden: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#6B4E7A"},
							Label:  "Команда",
							Meta:   "13",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#8C6239"}, Label: "Рассылки", Meta: "—", Hidden: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "работа"}, w.Tag{Label: "издательство"}},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New address book"},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined address book"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("0 of 4 books · 0 accounts")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New contact", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "up"},
								w.Button{Label: "Export", Icon: "down"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Empty{
						Parts: []w.EmptyPart{
							w.Strong("No books are ticked."),
							w.Text(" Tick the books to show in the rail, or all of them at once. "),
							w.Button{Label: "Tick every book", Kind: w.ButtonPrimary},
						},
					},
				},
			},
		},
	}
}

// States8 is frame «Контакты · опрос идёт» (states-8).
func States8() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"},
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
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#8C6239"}, Label: "Рассылки", Meta: "—", Hidden: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "работа"}, w.Tag{Label: "издательство"}},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New address book"},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined address book"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New contact", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "up"},
								w.Button{Label: "Export", Icon: "down"},
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
							w.Seg{Options: []w.SegOption{{Label: "A–Z", On: true}, {Label: "Recently changed"}}, Secondary: true},
							w.Button{Label: "Columns", Icon: "cols"},
						},
						More: &w.Button{Icon: "dots", Title: "More"},
						Right: &w.BarRight{
							Items: []w.BarPart{
								w.CheckForm{
									Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}},
									Action: "/app/contacts",
									Apply:  w.Button{Label: "Apply", Small: true},
								},
								w.Mono{Text: "5"},
							},
							Secondary: true,
						},
					},
					w.Runbar{Meter: w.Meter{Percent: 62}, Parts: []w.RunbarPart{w.Text("Polled 2 of 3 sources · 5 cards")}},
					w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#4A6B52"},
								Avatar:   w.Avatar{Initials: "АК"},
								Title:    []w.Inline{w.Text("Анна Ковалёва")},
								Subtitle: "+7 912 445-10-88",
								Detail:   "a.kovaleva@example.org",
								Caption:  "Издательство «Слово»",
								Source:   "nextcloud · Контакты",
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
					w.ListGroup{Title: w.Rubric{Text: "В", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
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
							w.FaceRow{
								Mark:     w.Mark{Colour: "#A05A2C"},
								Avatar:   w.Avatar{Initials: "ВП"},
								Title:    []w.Inline{w.Text("Вадим Петров")},
								Subtitle: "—",
								Detail:   "v.petrov@example.net",
								Caption:  "—",
								Source:   "nextcloud · Архив",
							},
						},
					},
				},
			},
		},
	}
}

// States9 is frame «Контакты · источник недоступен» (states-9).
func States9() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"},
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
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Error: true},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#8C6239"}, Label: "Рассылки", Meta: "—", Hidden: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "работа"}, w.Tag{Label: "издательство"}},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New address book"},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined address book"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 5 cards")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New contact", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "up"},
								w.Button{Label: "Export", Icon: "down"},
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
							w.Seg{Options: []w.SegOption{{Label: "A–Z", On: true}, {Label: "Recently changed"}}, Secondary: true},
							w.Button{Label: "Columns", Icon: "cols"},
						},
						More: &w.Button{Icon: "dots", Title: "More"},
						Right: &w.BarRight{
							Items: []w.BarPart{
								w.CheckForm{
									Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}},
									Action: "/app/contacts",
									Apply:  w.Button{Label: "Apply", Small: true},
								},
								w.Mono{Text: "5"},
							},
							Secondary: true,
						},
					},
					w.Runbar{
						Meter: w.Meter{Percent: 62, Alert: true},
						Parts: []w.RunbarPart{
							w.Text("Polled 2 of 3 sources · 5 cards · "),
							w.Drop{
								Trigger: w.Button{Label: "1 failed", Menu: true, Kind: w.ButtonDanger, Small: true},
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuRefusal{
											Mark:   w.Mark{Colour: "#6B4E7A", Size: w.MarkH11},
											Source: "radicale · vps · Команда",
											Answer: "REPORT /dav/addressbooks/team/ → 503 Service Unavailable",
											Retry:  w.Button{Label: "Retry", Small: true},
										},
									},
									Width: w.MenuWidth420,
									Pop:   w.MenuUnderEnd,
								},
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#4A6B52"},
								Avatar:   w.Avatar{Initials: "АК"},
								Title:    []w.Inline{w.Text("Анна Ковалёва")},
								Subtitle: "+7 912 445-10-88",
								Detail:   "a.kovaleva@example.org",
								Caption:  "Издательство «Слово»",
								Source:   "nextcloud · Контакты",
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
					w.ListGroup{Title: w.Rubric{Text: "В", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
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
							w.FaceRow{
								Mark:     w.Mark{Colour: "#A05A2C"},
								Avatar:   w.Avatar{Initials: "ВП"},
								Title:    []w.Inline{w.Text("Вадим Петров")},
								Subtitle: "—",
								Detail:   "v.petrov@example.net",
								Caption:  "—",
								Source:   "nextcloud · Архив",
							},
						},
					},
					w.Empty{
						Parts: []w.EmptyPart{
							w.Text("What arrived in time is shown. Cards from the unavailable book are missing — not “there are none”, but “we have not seen them”."),
						},
						Variant: w.EmptySmall,
					},
				},
			},
		},
	}
}

// States10 is frame «Книга пуста» (states-10).
func States10() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"},
					},
					On: 1,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked: w.TickedPart,
							Kind:   w.TickRowAll,
							Mark:   w.Mark{Kind: w.MarkMulti},
							Label:  "Combined address book",
							Meta:   "312",
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
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Рассылки", Meta: "0", Current: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "работа"}, w.Tag{Label: "издательство"}},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New address book"},
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
						Title: &w.HeadTitle{Text: "Рассылки"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("radicale · vps · 0 cards")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New contact", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "up"},
								w.Button{Label: "Export", Icon: "down"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Empty{
						Parts: []w.EmptyPart{
							w.Strong("This book is empty."),
							w.Text(" Add a contact, or import a file of them. "),
							w.Button{Label: "New contact", Icon: "plus", Kind: w.ButtonPrimary},
							w.Text(" "),
							w.Button{Label: "Import", Icon: "up"},
						},
					},
				},
			},
		},
	}
}

// States11 is frame «Книга · сервер не ответил» (states-11).
func States11() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"},
					},
					On: 1,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked: w.TickedPart,
							Kind:   w.TickRowAll,
							Mark:   w.Mark{Kind: w.MarkMulti},
							Label:  "Combined address book",
							Meta:   "312",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "—", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "51"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "13", Trail: w.Badge{Label: "ro"}},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#8C6239"}, Label: "Рассылки", Meta: "—", Hidden: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "работа"}, w.Tag{Label: "издательство"}},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New address book"},
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
						Title: &w.HeadTitle{Text: "Контакты"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New contact", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "up"},
								w.Button{Label: "Export", Icon: "down"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Strong("Контакты is unavailable."),
							w.Text(" The server did not answer as expected, so nothing on this screen comes from it. "),
							w.Button{Label: "Try again", Small: true},
						},
						Tone: w.MsgAlert,
					},
					w.MainDiag{Diag: w.Diag{Text: "REPORT /alice/contacts/kontakty/\nnextcloud\ndav: HTTP 503: the server is busy"}},
				},
			},
		},
	}
}

// States12 is frame «Правка · сервер не сохранил свойства» (states-12).
func States12() Screen {
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
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to the card", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "248", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "51"},
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
						Title: &w.HeadTitle{Text: "Анна Ковалёва"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}}},
						Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Save", Kind: w.ButtonPrimary}}},
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Strong("The change was saved, but the server did not keep these properties."),
							w.Text(" Other clients will not see them either: "),
							w.Mono{Text: "X-ABLABEL"},
							w.Text(", "),
							w.Mono{Text: "X-CARREL-UNKNOWN-FLAG"},
							w.Text(", "),
							w.Mono{Text: "X-EVOLUTION-FILE-AS"},
							w.Text("."),
						},
						Space: w.MsgFlushTop2,
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label:     "Photo",
								LabelSpan: true,
								Parts: []w.FieldPart{
									w.Photo{
										Initials: "АК",
										Side: []w.FieldPart{
											w.Button{Label: "Add"},
											w.FieldHint{Parts: []w.Inline{w.Text("JPEG, PNG, GIF or WebP.")}, Variant: w.HintNarrow},
										},
									},
								},
								Width: w.FieldFixAuto,
							},
							w.FormCols{
								Items: []w.Field{
									w.Field{
										Label: "Display name",
										Parts: []w.FieldPart{w.Input{Value: "Анна Ковалёва", Label: "Display name", Focus: true}},
										Width: w.FieldXL,
									},
									w.Field{Label: "Given", Parts: []w.FieldPart{w.Input{Value: "Анна", Label: "Given"}}, Width: w.FieldSM},
									w.Field{Label: "Family", Parts: []w.FieldPart{w.Input{Value: "Ковалёва", Label: "Family"}}, Width: w.FieldSM},
									w.Field{
										Label: "Nickname",
										Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Nickname"}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Organisation",
										Parts: []w.FieldPart{w.Input{Value: "Издательство «Слово»", Label: "Organisation"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Title",
										Parts: []w.FieldPart{w.Input{Value: "Выпускающий редактор", Label: "Title"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Birthday",
										Parts: []w.FieldPart{
											w.Input{Value: "14.03.1987", Label: "Birthday", Mono: true, Numeric: true},
											w.FieldHint{Parts: []w.Inline{w.Text("Type it: 14.03.1987, or 14.03 without the year.")}},
										},
										Width: w.FieldSM,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Phone numbers"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Label", Parts: []w.FieldPart{w.Pick{Value: "mobile"}}, Width: w.FieldXS},
									w.Field{
										Label: "Number",
										Parts: []w.FieldPart{w.Input{Value: "+7 912 445-10-88", Label: "Number", Mono: true}},
										Width: w.FieldMD,
									},
									w.Button{Label: "Remove", Kind: w.ButtonDanger},
									w.Grow{},
									w.Field{Parts: []w.FieldPart{w.Pick{Value: "work", Label: "Label"}}, Width: w.FieldXS},
									w.Field{
										Parts: []w.FieldPart{w.Input{Value: "+7 495 120-45-00", Label: "Number", Mono: true}},
										Width: w.FieldMD,
									},
									w.Button{Label: "Remove", Kind: w.ButtonDanger},
									w.Grow{},
									w.Button{Label: "Add phone", Icon: "plus", Small: true},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Email addresses"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Label", Parts: []w.FieldPart{w.Pick{Value: "work"}}, Width: w.FieldXS},
									w.Field{
										Label: "Address",
										Parts: []w.FieldPart{w.Input{Value: "a.kovaleva@example.org", Label: "Address"}},
										Width: w.FieldLG,
									},
									w.Button{Label: "Remove", Kind: w.ButtonDanger},
									w.Grow{},
									w.Button{Label: "Add email", Icon: "plus", Small: true},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Addresses"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Label", Parts: []w.FieldPart{w.Pick{Value: "home"}}, Width: w.FieldXS},
									w.Field{
										Label: "Street",
										Parts: []w.FieldPart{w.Input{Value: "ул. Садовая, 12, кв. 5", Label: "Street"}},
										Width: w.FieldLG,
									},
									w.Field{Label: "City", Parts: []w.FieldPart{w.Input{Value: "Москва", Label: "City"}}, Width: w.FieldSM},
									w.Field{Label: "Region", Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Region"}}, Width: w.FieldSM},
									w.Field{
										Label: "Postcode",
										Parts: []w.FieldPart{w.Input{Value: "123001", Label: "Postcode", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{Label: "Country", Parts: []w.FieldPart{w.Input{Value: "Россия", Label: "Country"}}, Width: w.FieldSM},
									w.Button{Label: "Remove", Kind: w.ButtonDanger},
									w.Grow{},
									w.Button{Label: "Add address", Icon: "plus", Small: true},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Other"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Tags",
										Parts: []w.FieldPart{w.Input{Value: "работа, издательство", Label: "Tags"}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Note",
										Parts: []w.FieldPart{w.Area{Text: "Пишет по вторникам, отвечает в тот же день.", Label: "Text"}},
										Width: w.FieldFull,
									},
								},
							},
							w.Field{Parts: []w.FieldPart{w.Button{Label: "Show 6 empty fields"}}},
						},
					},
					w.FormFoot{
						Items: []w.FormFootPart{
							w.Dim{Parts: []w.Inline{w.Text("Properties Carrel does not know are kept unchanged.")}, Size: w.Dim11_5},
						},
					},
				},
			},
		},
	}
}

// States13 is frame «Агенда · опрос идёт» (states-13).
func States13() Screen {
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
						{Icon: "con", Label: "Contacts", Href: "#"},
						{Icon: "cal", Label: "Calendar", Href: "#"},
						{Icon: "task", Label: "Tasks", Href: "#"},
						{Icon: "note", Label: "Notes", Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Href: "#"},
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
							Meta:    "2",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "1"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "0"},
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
						w.TickRow{Kind: w.TickRowExternal, Mark: w.Mark{Colour: "#8C6239"}, Label: "Праздники РФ", Meta: "ro"},
						w.TickRow{
							Ticked: w.TickedOff,
							Kind:   w.TickRowExternal,
							Mark:   w.Mark{Colour: "#6B4E7A"},
							Label:  "Расписание лиги",
							Meta:   "ro",
							Hidden: true,
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("6 of 7 calendars · 3 accounts · 10 – 16 August 2026")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "up"},
								w.Button{Label: "Export", Icon: "down"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Button{Icon: "back", Title: "Previous week", Href: "/app/calendar?from=2026-08-03"},
							w.Button{Label: "Today", Href: "/app/calendar"},
							w.Button{Icon: "next", Title: "Next week", Href: "/app/calendar?from=2026-08-17"},
							w.Seg{
								Options: []w.SegOption{{Label: "Day"}, {Label: "Week", On: true}, {Label: "Month"}, {Label: "Range"}},
							},
							w.Range{
								From:      w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true},
								To:        w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-16", Label: "To", Mono: true},
								Show:      w.Button{Label: "Show"},
								Hidden:    []w.Hidden{{Name: "view", Value: "range"}},
								Secondary: true,
								Action:    "/app/calendar",
							},
							w.BarSep{Secondary: true},
							w.CheckForm{
								Checks: []w.Check{
									w.Check{On: true, Name: "show", Value: "events", Parts: []w.CheckPart{w.Text("Events")}},
									w.Check{Name: "show", Value: "tasks", Parts: []w.CheckPart{w.Text("Tasks due")}},
									w.Check{Name: "show", Value: "notes", Parts: []w.CheckPart{w.Text("Notes")}},
								},
								Action:   "/app/calendar",
								Tertiary: true,
								Apply:    w.Button{Label: "Apply", Small: true},
							},
							w.Drop{
								Trigger: w.Button{Icon: "dots", Title: "More"},
								More:    true,
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.CheckForm{
											Checks: []w.Check{
												w.Check{On: true, Name: "show", Value: "events", Parts: []w.CheckPart{w.Text("Events")}},
												w.Check{Name: "show", Value: "tasks", Parts: []w.CheckPart{w.Text("Tasks due")}},
												w.Check{Name: "show", Value: "notes", Parts: []w.CheckPart{w.Text("Notes")}},
											},
											Action:   "/app/calendar",
											Tertiary: true,
											Apply:    w.Button{Label: "Apply", Small: true},
										},
										w.MenuSep{Tertiary: true},
										w.Range{
											From:   w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true},
											To:     w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-16", Label: "To", Mono: true},
											Show:   w.Button{Label: "Show"},
											Hidden: []w.Hidden{{Name: "view", Value: "range"}},
											Action: "/app/calendar",
											Wrap:   true,
										},
										w.MenuSep{},
										w.CheckForm{
											Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}},
											Action: "/app/calendar",
											Apply:  w.Button{Label: "Apply", Small: true},
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
									Action: "/app/calendar",
									Apply:  w.Button{Label: "Apply", Small: true},
								},
							},
							Secondary: true,
						},
						Variant: w.BarWide,
					},
					w.Runbar{
						Meter: w.Meter{Percent: 62},
						Parts: []w.RunbarPart{
							w.Text("Polled 4 of 6 sources · 9 events · "),
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
							w.Mono{Text: "0.8 s", Place: w.Pushed},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
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
								Time:   "15:00",
								Title:  []w.Inline{w.Text("Созвон с подрядчиком")},
								Detail: "Jitsi",
								Source: "nextcloud · Работа",
							},
						},
					},
				},
			},
		},
	}
}

// States14 is frame «Агенда · ни один календарь не отмечен» (states-14).
func States14() Screen {
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
						{Icon: "con", Label: "Contacts", Href: "#"},
						{Icon: "cal", Label: "Calendar", Href: "#"},
						{Icon: "task", Label: "Tasks", Href: "#"},
						{Icon: "note", Label: "Notes", Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Href: "#"},
					},
					On: 2,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedOff,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined calendar",
							Meta:    "—",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "—", Hidden: true},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "—", Hidden: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "—", Hidden: true},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "—", Hidden: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "baikal"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Meta: "—", Hidden: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Subscriptions", Items: []w.RubricPart{w.Badge{Label: "external"}}},
						w.TickRow{
							Ticked: w.TickedOff,
							Kind:   w.TickRowExternal,
							Mark:   w.Mark{Colour: "#8C6239"},
							Label:  "Праздники РФ",
							Meta:   "ro",
							Hidden: true,
						},
						w.TickRow{
							Ticked: w.TickedOff,
							Kind:   w.TickRowExternal,
							Mark:   w.Mark{Colour: "#6B4E7A"},
							Label:  "Расписание лиги",
							Meta:   "ro",
							Hidden: true,
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("0 of 7 calendars · 0 accounts")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "up"},
								w.Button{Label: "Export", Icon: "down"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Empty{
						Parts: []w.EmptyPart{
							w.Strong("No calendars are ticked."),
							w.Text(" Tick the calendars to show in the rail, or all of them at once. "),
							w.Button{Label: "Tick every calendar", Kind: w.ButtonPrimary},
						},
					},
				},
			},
		},
	}
}

// States15 is frame «Календарь · пустая неделя» (states-15).
func States15() Screen {
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
						{Icon: "con", Label: "Contacts", Href: "#"},
						{Icon: "cal", Label: "Calendar", Href: "#"},
						{Icon: "task", Label: "Tasks", Href: "#"},
						{Icon: "note", Label: "Notes", Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Href: "#"},
					},
					On: 2,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked: w.TickedPart,
							Kind:   w.TickRowAll,
							Mark:   w.Mark{Kind: w.MarkMulti},
							Label:  "Combined calendar",
							Meta:   "0",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "0", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "0"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "0"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "baikal"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Meta: "0"}},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Subscriptions", Items: []w.RubricPart{w.Badge{Label: "external"}}},
						w.TickRow{Kind: w.TickRowExternal, Mark: w.Mark{Colour: "#8C6239"}, Label: "Праздники РФ", Meta: "ro"},
						w.TickRow{
							Ticked: w.TickedOff,
							Kind:   w.TickRowExternal,
							Mark:   w.Mark{Colour: "#6B4E7A"},
							Label:  "Расписание лиги",
							Meta:   "ro",
							Hidden: true,
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
						Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH20}, Text: "Личный"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · 2 – 8 November 2026 · 0 events")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New event", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "up"},
								w.Button{Label: "Export", Icon: "down"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Button{Icon: "back", Title: "Previous week", Href: "/app/calendar/nextcloud/lichny?from=2026-10-26"},
							w.Button{Label: "Today", Href: "/app/calendar/nextcloud/lichny"},
							w.Button{Icon: "next", Title: "Next week", Href: "/app/calendar/nextcloud/lichny?from=2026-11-09"},
							w.Seg{
								Options: []w.SegOption{{Label: "Day"}, {Label: "Week", On: true}, {Label: "Month"}, {Label: "Range"}},
							},
							w.Range{
								From:      w.Input{Kind: w.InputDate, Name: "from", Value: "2026-11-02", Label: "From", Mono: true},
								To:        w.Input{Kind: w.InputDate, Name: "to", Value: "2026-11-08", Label: "To", Mono: true},
								Show:      w.Button{Label: "Show"},
								Hidden:    []w.Hidden{{Name: "view", Value: "range"}},
								Secondary: true,
								Action:    "/app/calendar",
							},
							w.BarSep{Secondary: true},
							w.CheckForm{
								Checks: []w.Check{
									w.Check{On: true, Name: "show", Value: "events", Parts: []w.CheckPart{w.Text("Events")}},
									w.Check{Name: "show", Value: "tasks", Parts: []w.CheckPart{w.Text("Tasks due")}},
									w.Check{Name: "show", Value: "notes", Parts: []w.CheckPart{w.Text("Notes")}},
								},
								Action:   "/app/calendar",
								Tertiary: true,
								Apply:    w.Button{Label: "Apply", Small: true},
							},
							w.Drop{
								Trigger: w.Button{Icon: "dots", Title: "More"},
								More:    true,
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.CheckForm{
											Checks: []w.Check{
												w.Check{On: true, Name: "show", Value: "events", Parts: []w.CheckPart{w.Text("Events")}},
												w.Check{Name: "show", Value: "tasks", Parts: []w.CheckPart{w.Text("Tasks due")}},
												w.Check{Name: "show", Value: "notes", Parts: []w.CheckPart{w.Text("Notes")}},
											},
											Action:   "/app/calendar",
											Tertiary: true,
											Apply:    w.Button{Label: "Apply", Small: true},
										},
										w.MenuSep{Tertiary: true},
										w.Range{
											From:   w.Input{Kind: w.InputDate, Name: "from", Value: "2026-11-02", Label: "From", Mono: true},
											To:     w.Input{Kind: w.InputDate, Name: "to", Value: "2026-11-08", Label: "To", Mono: true},
											Show:   w.Button{Label: "Show"},
											Hidden: []w.Hidden{{Name: "view", Value: "range"}},
											Action: "/app/calendar",
											Wrap:   true,
										},
										w.MenuSep{},
										w.CheckForm{
											Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}},
											Action: "/app/calendar",
											Apply:  w.Button{Label: "Apply", Small: true},
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
									Action: "/app/calendar",
									Apply:  w.Button{Label: "Apply", Small: true},
								},
							},
							Secondary: true,
						},
						Variant: w.BarWide,
					},
					w.Empty{
						Parts: []w.EmptyPart{w.Strong("No events this week."), w.Text(" Nothing in Личный from 2 to 8 November.")},
					},
				},
			},
		},
	}
}
