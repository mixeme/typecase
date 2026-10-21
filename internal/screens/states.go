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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
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
				w.Button{Label: "Refresh", Icon: "refresh", Title: "The poll is still running.", State: w.ButtonDisabled},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
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
						w.TickRow{
							Kind:  w.TickRowExternal,
							Mark:  w.Mark{Colour: "#8C6239"},
							Label: "Праздники РФ",
							Meta:  "—",
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
								Parts: []w.HeadSubPart{w.Text("5 of 7 calendars · 3 accounts · 10 – 16 August 2026 · "), w.Mono{Text: "Europe/Moscow"}},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Button{Icon: "back", Title: "Previous week", Href: "/app/calendar?from=2026-08-03"},
							w.Button{Label: "Today", Href: "/app/calendar"},
							w.Button{Icon: "next", Title: "Next week", Href: "/app/calendar?from=2026-08-17"},
							w.SegFold{
								Seg: w.Seg{
									Options: []w.SegOption{
										{Label: "Day", Href: "/app/calendar?view=day"},
										{Label: "Week", Href: "/app/calendar?view=week", On: true},
										{Label: "Month", Href: "/app/calendar?view=month"},
										{Label: "Range", Href: "/app/calendar?view=range"},
									},
								},
								Phone: true,
							},
							w.Range{
								From:      w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true},
								To:        w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-16", Label: "To", Mono: true},
								Show:      w.Button{Label: "Show", Submit: true},
								Hidden:    []w.Hidden{{Name: "view", Value: "range"}},
								Secondary: true,
								Action:    "/app/calendar",
							},
							w.Button{Label: "Columns", Icon: "cols"},
							w.Drop{
								Trigger: w.Button{Icon: "dots", Title: "More"},
								More:    true,
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.Range{
											From:   w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true},
											To:     w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-16", Label: "To", Mono: true},
											Show:   w.Button{Label: "Show", Submit: true},
											Hidden: []w.Hidden{{Name: "view", Value: "range"}},
											Action: "/app/calendar",
											Wrap:   true,
										},
										w.MenuSep{},
										w.CheckForm{
											Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}},
											Action: "/app/calendar",
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
									Action: "/app/calendar",
									Apply:  w.Button{Label: "Apply", Small: true, Submit: true},
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
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
							Meta:   "7",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "5"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{
							Mark:    w.Mark{Colour: "#6B4E7A"},
							Label:   "Команда",
							Meta:    "1",
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
						Subs: []w.HeadSub{
							w.HeadSub{Parts: []w.HeadSubPart{w.Text("radicale · vps · 1 card · the server granted no write access")}},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{w.Button{Label: "Export", Icon: "export"}, w.Button{Label: "Print", Icon: "print"}},
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
								w.Mono{Text: "1"},
							},
							Secondary: true,
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#6B4E7A"},
								Avatar:   w.Avatar{Photo: true},
								Title:    []w.Inline{w.Text("Артём Белов")},
								Subtitle: "+7 903 118-22-04",
								Detail:   "a.belov@example.org",
								Caption:  "Тензор",
								Single:   true,
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
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
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
									w.Text("6 of 7 calendars · 3 accounts · 10 – 16 August 2026 · "),
									w.Mono{Text: "Europe/Moscow"},
									w.Text(" · 6 events"),
								},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Button{Icon: "back", Title: "Previous week", Href: "/app/calendar?from=2026-08-03"},
							w.Button{Label: "Today", Href: "/app/calendar"},
							w.Button{Icon: "next", Title: "Next week", Href: "/app/calendar?from=2026-08-17"},
							w.SegFold{
								Seg: w.Seg{
									Options: []w.SegOption{
										{Label: "Day", Href: "/app/calendar?view=day"},
										{Label: "Week", Href: "/app/calendar?view=week", On: true},
										{Label: "Month", Href: "/app/calendar?view=month"},
										{Label: "Range", Href: "/app/calendar?view=range"},
									},
								},
								Phone: true,
							},
							w.Range{
								From:      w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true},
								To:        w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-16", Label: "To", Mono: true},
								Show:      w.Button{Label: "Show", Submit: true},
								Hidden:    []w.Hidden{{Name: "view", Value: "range"}},
								Secondary: true,
								Action:    "/app/calendar",
							},
							w.Button{Label: "Columns", Icon: "cols"},
							w.Drop{
								Trigger: w.Button{Icon: "dots", Title: "More"},
								More:    true,
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.Range{
											From:   w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true},
											To:     w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-16", Label: "To", Mono: true},
											Show:   w.Button{Label: "Show", Submit: true},
											Hidden: []w.Hidden{{Name: "view", Value: "range"}},
											Action: "/app/calendar",
											Wrap:   true,
										},
										w.MenuSep{},
										w.CheckForm{
											Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}},
											Action: "/app/calendar",
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
									Action: "/app/calendar",
									Apply:  w.Button{Label: "Apply", Small: true, Submit: true},
								},
							},
							Secondary: true,
						},
						Variant: w.BarWide,
					},
					w.Runbar{
						Meter: w.Meter{Percent: 83, Alert: true},
						Parts: []w.RunbarPart{
							w.Text("Polled 5 of 6 sources · 6 events · "),
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
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
					w.List{
						Rows: []w.ListRow{
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "2", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "2"},
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
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Field{
									Parts: []w.FieldPart{
										w.Drop{
											Trigger: w.Pick{Value: "nextcloud · Личный", Mark: &w.Mark{Colour: "#4A6B52"}},
											Menu: w.Menu{
												Items: []w.MenuPart{
													w.MenuItem{
														Mark:  &w.Mark{Colour: "#4A6B52"},
														Label: "nextcloud · Личный",
														On:    true,
														Name:  "calendar",
														Value: "nextcloud/lichnyy",
														Of:    "event-edit-loss",
													},
													w.MenuItem{
														Mark:  &w.Mark{Colour: "#3F6E8C"},
														Label: "nextcloud · Работа",
														Name:  "calendar",
														Value: "nextcloud/rabota",
														Of:    "event-edit-loss",
													},
												},
												Width: w.MenuWidth290,
												Pop:   w.MenuUnder,
											},
										},
									},
									Width: w.FieldMD,
								},
								w.Button{Label: "Cancel"},
								w.Button{Label: "Save", Kind: w.ButtonPrimary},
							},
						},
					},
					w.Runbar{
						Meter:    w.Meter{Percent: 0},
						Progress: "event-edit-loss",
						Parts:    []w.RunbarPart{w.Text("Saving · "), w.Mono{Text: "0 of 0 MB"}},
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
										Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "start.date", Value: "2026-08-10", Label: "Start date", Mono: true}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Start time",
										Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Name: "start.time", Value: "12:00", Label: "Start time", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "End date",
										Parts: []w.FieldPart{
											w.Input{
												Kind:    w.InputDate,
												Name:    "end.date",
												Value:   "2026-08-10",
												Label:   "End date",
												Mono:    true,
												Follows: "start.date",
											},
										},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "End time",
										Parts: []w.FieldPart{
											w.Input{
												Kind:    w.InputTime,
												Name:    "end.time",
												Value:   "13:00",
												Label:   "End time",
												Mono:    true,
												Follows: "start.time",
												Day:     "end.date",
											},
										},
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
												Trigger: w.Button{Label: "Attach", Menu: true, Icon: "clip"},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{
															Icon:  "upload",
															File:  &w.FileField{Form: "post-attach", Name: "file", Label: "From this computer"},
															Label: "From this computer",
														},
														w.MenuItem{Icon: "file", Label: "From Files"},
														w.MenuItem{Label: "Upload", NoScript: true, Form: "post-attach"},
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

// States36 is frame «Событие · сервер не сохранил дни серии» (states-36).
func States36() Screen {
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "2", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "2"},
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
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Field{
									Parts: []w.FieldPart{
										w.Drop{
											Trigger: w.Pick{Value: "nextcloud · Личный", Mark: &w.Mark{Colour: "#4A6B52"}},
											Menu: w.Menu{
												Items: []w.MenuPart{
													w.MenuItem{
														Mark:  &w.Mark{Colour: "#4A6B52"},
														Label: "nextcloud · Личный",
														On:    true,
														Name:  "calendar",
														Value: "nextcloud/lichnyy",
														Of:    "event-edit-loss-days",
													},
													w.MenuItem{
														Mark:  &w.Mark{Colour: "#3F6E8C"},
														Label: "nextcloud · Работа",
														Name:  "calendar",
														Value: "nextcloud/rabota",
														Of:    "event-edit-loss-days",
													},
												},
												Width: w.MenuWidth290,
												Pop:   w.MenuUnder,
											},
										},
									},
									Width: w.FieldMD,
								},
								w.Button{Label: "Cancel"},
								w.Button{Label: "Save", Kind: w.ButtonPrimary},
							},
						},
					},
					w.Runbar{
						Meter:    w.Meter{Percent: 0},
						Progress: "event-edit-loss-days",
						Parts:    []w.RunbarPart{w.Text("Saving · "), w.Mono{Text: "0 of 0 MB"}},
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Strong("The change was saved, but the server did not keep two days of the series."),
							w.Text(" Other clients will not see them either: the meeting of 24 August, which was skipped, is back; the meeting of 31 August, which was moved to 1 September, 15:00, is back on its day at 12:00."),
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
										Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "start.date", Value: "2026-08-10", Label: "Start date", Mono: true}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Start time",
										Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Name: "start.time", Value: "12:00", Label: "Start time", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "End date",
										Parts: []w.FieldPart{
											w.Input{
												Kind:    w.InputDate,
												Name:    "end.date",
												Value:   "2026-08-10",
												Label:   "End date",
												Mono:    true,
												Follows: "start.date",
											},
										},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "End time",
										Parts: []w.FieldPart{
											w.Input{
												Kind:    w.InputTime,
												Name:    "end.time",
												Value:   "13:00",
												Label:   "End time",
												Mono:    true,
												Follows: "start.time",
												Day:     "end.date",
											},
										},
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
												Trigger: w.Button{Label: "Attach", Menu: true, Icon: "clip"},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{
															Icon:  "upload",
															File:  &w.FileField{Form: "post-attach", Name: "file", Label: "From this computer"},
															Label: "From this computer",
														},
														w.MenuItem{Icon: "file", Label: "From Files"},
														w.MenuItem{Label: "Upload", NoScript: true, Form: "post-attach"},
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
				w.Button{Label: "Refresh", Icon: "refresh", Title: "Nothing is ticked to ask.", State: w.ButtonDisabled},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 1,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedOff,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined address book",
							Meta:    "—",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "5", Hidden: true},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1", Hidden: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#6B4E7A"},
							Label:  "Команда",
							Meta:   "1",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#8C6239"}, Label: "Рассылки", Meta: "—", Hidden: true},
					},
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
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
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
				w.Button{Label: "Refresh", Icon: "refresh", Title: "The poll is still running.", State: w.ButtonDisabled},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
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
							Meta:    "—",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "5"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "—", Trail: w.Badge{Label: "ro"}},
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
								w.Mono{Text: "6"},
							},
							Secondary: true,
						},
					},
					w.Runbar{
						Meter: w.Meter{Percent: 62},
						Parts: []w.RunbarPart{w.Text("Polling · 2 of 3 answered · 6 cards"), w.Mono{Text: "0.8 s", Place: w.Pushed}},
					},
					w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
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
								Avatar:   w.Avatar{Initials: "АБ"},
								Title:    []w.Inline{w.Text("Артём Белов")},
								Subtitle: "+7 903 118-22-04",
								Detail:   "a.belov@example.org",
								Caption:  "—",
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
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
							Meta:    "6",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "5"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 6 cards")}}},
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
								w.Mono{Text: "6"},
							},
							Secondary: true,
						},
					},
					w.Runbar{
						Meter: w.Meter{Percent: 62, Alert: true},
						Parts: []w.RunbarPart{
							w.Text("Polled 2 of 3 sources · 6 cards · "),
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
							w.Mono{Text: "2 min ago", Place: w.Pushed},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
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
								Avatar:   w.Avatar{Initials: "АБ"},
								Title:    []w.Inline{w.Text("Артём Белов")},
								Subtitle: "+7 903 118-22-04",
								Detail:   "a.belov@example.org",
								Caption:  "—",
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
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
							Meta:   "7",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "5"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "1", Trail: w.Badge{Label: "ro"}},
						w.TickRow{
							Ticked:  w.TickedOff,
							Mark:    w.Mark{Colour: "#8C6239"},
							Label:   "Рассылки",
							Meta:    "—",
							Current: true,
							Hidden:  true,
						},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#8C6239", Size: w.MarkH20}, Text: "Рассылки"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("radicale · vps · 0 cards")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New contact", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
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
							w.Button{Label: "Import", Icon: "import"},
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
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
							Meta:   "2",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Current: true, Error: true},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH20}, Text: "Контакты"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New contact", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "5", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
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
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Field{
									Parts: []w.FieldPart{
										w.Drop{
											Trigger: w.Pick{Value: "nextcloud · Контакты", Mark: &w.Mark{Colour: "#4A6B52"}},
											Menu: w.Menu{
												Items: []w.MenuPart{
													w.MenuItem{
														Mark:  &w.Mark{Colour: "#4A6B52"},
														Label: "nextcloud · Контакты",
														On:    true,
														Name:  "book",
														Value: "nextcloud/kontakty",
														Of:    "contact-edit-loss",
													},
													w.MenuItem{
														Mark:  &w.Mark{Colour: "#A05A2C"},
														Label: "nextcloud · Архив",
														Name:  "book",
														Value: "nextcloud/arhiv",
														Of:    "contact-edit-loss",
													},
												},
												Width: w.MenuWidth290,
												Pop:   w.MenuUnder,
											},
										},
									},
									Width: w.FieldMD,
								},
								w.Button{Label: "Cancel"},
								w.Button{Label: "Save", Kind: w.ButtonPrimary},
							},
						},
					},
					w.Runbar{
						Meter:    w.Meter{Percent: 0},
						Progress: "contact-edit-loss",
						Parts:    []w.RunbarPart{w.Text("Saving · "), w.Mono{Text: "0 of 0 MB"}},
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
										ID:       "photo",
										Side: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Button{
														Label: "Add",
														File:  &w.FileField{Form: "", Name: "photo", Accept: "image/*", Label: "Photo", Preview: "photo"},
													},
												},
												Gap: w.LineGap6,
											},
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
							w.Field{Parts: []w.FieldPart{w.Button{Label: "Show 7 empty fields"}}},
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
				w.Button{Label: "Refresh", Icon: "refresh", Title: "The poll is still running.", State: w.ButtonDisabled},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
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
							Meta:    "—",
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
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "—"},
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
								Parts: []w.HeadSubPart{w.Text("6 of 7 calendars · 3 accounts · 10 – 16 August 2026 · "), w.Mono{Text: "Europe/Moscow"}},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Button{Icon: "back", Title: "Previous week", Href: "/app/calendar?from=2026-08-03"},
							w.Button{Label: "Today", Href: "/app/calendar"},
							w.Button{Icon: "next", Title: "Next week", Href: "/app/calendar?from=2026-08-17"},
							w.SegFold{
								Seg: w.Seg{
									Options: []w.SegOption{
										{Label: "Day", Href: "/app/calendar?view=day"},
										{Label: "Week", Href: "/app/calendar?view=week", On: true},
										{Label: "Month", Href: "/app/calendar?view=month"},
										{Label: "Range", Href: "/app/calendar?view=range"},
									},
								},
								Phone: true,
							},
							w.Range{
								From:      w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true},
								To:        w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-16", Label: "To", Mono: true},
								Show:      w.Button{Label: "Show", Submit: true},
								Hidden:    []w.Hidden{{Name: "view", Value: "range"}},
								Secondary: true,
								Action:    "/app/calendar",
							},
							w.Button{Label: "Columns", Icon: "cols"},
							w.Drop{
								Trigger: w.Button{Icon: "dots", Title: "More"},
								More:    true,
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.Range{
											From:   w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true},
											To:     w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-16", Label: "To", Mono: true},
											Show:   w.Button{Label: "Show", Submit: true},
											Hidden: []w.Hidden{{Name: "view", Value: "range"}},
											Action: "/app/calendar",
											Wrap:   true,
										},
										w.MenuSep{},
										w.CheckForm{
											Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}},
											Action: "/app/calendar",
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
									Action: "/app/calendar",
									Apply:  w.Button{Label: "Apply", Small: true, Submit: true},
								},
							},
							Secondary: true,
						},
						Variant: w.BarWide,
					},
					w.Runbar{
						Meter: w.Meter{Percent: 62},
						Parts: []w.RunbarPart{
							w.Text("Polling · 4 of 6 answered · 5 events · "),
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
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
					w.List{
						Rows: []w.ListRow{
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
								Mark:   w.Mark{Colour: "#3F6E8C"},
								Time:   "15:00",
								Title:  []w.Inline{w.Text("Созвон с подрядчиком")},
								Detail: "Jitsi",
								Badges: []w.Badge{w.Badge{Label: "tentative", Tone: w.BadgeAlert}},
								Source: "nextcloud · Работа",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Tuesday · 11 August", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
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
				w.Button{Label: "Refresh", Icon: "refresh", Title: "Nothing is ticked to ask.", State: w.ButtonDisabled},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
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
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("0 of 7 calendars · 0 accounts")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH20}, Text: "Личный"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{w.Text("nextcloud · 2 – 8 November 2026 · "), w.Mono{Text: "Europe/Moscow"}, w.Text(" · 0 events")},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New event", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Button{Icon: "back", Title: "Previous week", Href: "/app/calendar/nextcloud/lichny?from=2026-10-26"},
							w.Button{Label: "Today", Href: "/app/calendar/nextcloud/lichny"},
							w.Button{Icon: "next", Title: "Next week", Href: "/app/calendar/nextcloud/lichny?from=2026-11-09"},
							w.SegFold{
								Seg: w.Seg{
									Options: []w.SegOption{
										{Label: "Day", Href: "/app/calendar?view=day"},
										{Label: "Week", Href: "/app/calendar?view=week", On: true},
										{Label: "Month", Href: "/app/calendar?view=month"},
										{Label: "Range", Href: "/app/calendar?view=range"},
									},
								},
								Phone: true,
							},
							w.Range{
								From:      w.Input{Kind: w.InputDate, Name: "from", Value: "2026-11-02", Label: "From", Mono: true},
								To:        w.Input{Kind: w.InputDate, Name: "to", Value: "2026-11-08", Label: "To", Mono: true},
								Show:      w.Button{Label: "Show", Submit: true},
								Hidden:    []w.Hidden{{Name: "view", Value: "range"}},
								Secondary: true,
								Action:    "/app/calendar",
							},
							w.Button{Label: "Columns", Icon: "cols"},
							w.Drop{
								Trigger: w.Button{Icon: "dots", Title: "More"},
								More:    true,
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.Range{
											From:   w.Input{Kind: w.InputDate, Name: "from", Value: "2026-11-02", Label: "From", Mono: true},
											To:     w.Input{Kind: w.InputDate, Name: "to", Value: "2026-11-08", Label: "To", Mono: true},
											Show:   w.Button{Label: "Show", Submit: true},
											Hidden: []w.Hidden{{Name: "view", Value: "range"}},
											Action: "/app/calendar",
											Wrap:   true,
										},
										w.MenuSep{},
										w.CheckForm{
											Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}},
											Action: "/app/calendar",
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
									Action: "/app/calendar",
									Apply:  w.Button{Label: "Apply", Small: true, Submit: true},
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

// States18 is frame «Задачи · ни один список не отмечен» (states-18).
func States18() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.Button{Label: "Refresh", Icon: "refresh", Title: "Nothing is ticked to ask.", State: w.ButtonDisabled},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 3,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedOff,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined list",
							Meta:    "—",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "4", Hidden: true},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "3", Hidden: true},
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
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New list"},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined list"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("0 of 3 lists · 0 accounts")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New task", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Empty{
						Parts: []w.EmptyPart{
							w.Strong("No lists are ticked."),
							w.Text(" Tick the lists to show in the rail, or all of them at once. "),
							w.Button{Label: "Tick every list", Kind: w.ButtonPrimary},
						},
					},
				},
			},
		},
	}
}

// States19 is frame «Задачи · опрос идёт» (states-19).
func States19() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.Button{Label: "Refresh", Icon: "refresh", Title: "The poll is still running.", State: w.ButtonDisabled},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 3,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedPart,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined list",
							Meta:    "—",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "4"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "—"},
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
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "дом"}, w.Tag{Label: "релиз"}, w.Tag{Label: "read"}},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New list"},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined list"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 of 3 lists · nextcloud")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New task", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter by word", Label: "Filter by word"},
								Waiting: true,
								Variant: w.FilterH26W200,
							},
							w.Seg{Options: []w.SegOption{{Label: "Open 3", On: true}, {Label: "Done 1"}, {Label: "All 4"}}},
							w.Seg{
								Options:   []w.SegOption{{Label: "By due", On: true}, {Label: "By priority"}, {Label: "Recently changed"}},
								Secondary: true,
							},
							w.Button{Label: "Columns", Icon: "cols"},
						},
					},
					w.Runbar{
						Meter: w.Meter{Percent: 62},
						Parts: []w.RunbarPart{w.Text("Polling · 1 of 2 answered · 4 tasks"), w.Mono{Text: "0.8 s", Place: w.Pushed}},
					},
					w.ListGroup{Title: w.Rubric{Text: "Overdue", Form: w.RubricInline}, Count: &w.Num{Value: 1, Alert: true}},
					w.List{
						Rows: []w.ListRow{
							w.DueRow{
								Tick:     w.Tick{Label: "Done"},
								Mark:     w.Mark{Colour: "#4A6B52"},
								Title:    []w.Inline{w.Text("Отправить документы в МФЦ")},
								Due:      "7 Aug",
								Progress: "—",
								Source:   "nextcloud · Дела",
								Overdue:  true,
								Href:     "#",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "This week", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.DueRow{
								Tick:     w.Tick{Label: "Done"},
								Mark:     w.Mark{Colour: "#4A6B52"},
								Title:    []w.Inline{w.Text("Собрать чемодан")},
								Due:      "12 Aug 18:00",
								Progress: "—",
								Source:   "nextcloud · Дела",
								Href:     "#",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "No due date", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.DueRow{
								Tick:     w.Tick{Done: true, Label: "Done"},
								Mark:     w.Mark{Colour: "#4A6B52"},
								Title:    []w.Inline{w.Text("Продлить страховку")},
								Due:      "—",
								Progress: "100%",
								Source:   "nextcloud · Дела",
								Href:     "#",
							},
						},
					},
				},
			},
		},
	}
}

// States20 is frame «Задачи · источник недоступен» (states-20).
func States20() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 3,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedPart,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined list",
							Meta:    "4",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "4"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Error: true},
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
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "дом"}, w.Tag{Label: "релиз"}, w.Tag{Label: "read"}},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New list"},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined list"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 of 3 lists · nextcloud · 4 tasks")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New task", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter by word", Label: "Filter by word"},
								Waiting: true,
								Variant: w.FilterH26W200,
							},
							w.Seg{Options: []w.SegOption{{Label: "Open 3", On: true}, {Label: "Done 1"}, {Label: "All 4"}}},
							w.Seg{
								Options:   []w.SegOption{{Label: "By due", On: true}, {Label: "By priority"}, {Label: "Recently changed"}},
								Secondary: true,
							},
							w.Button{Label: "Columns", Icon: "cols"},
						},
					},
					w.Runbar{
						Meter: w.Meter{Percent: 62, Alert: true},
						Parts: []w.RunbarPart{
							w.Text("Polled 1 of 2 sources · 4 tasks · "),
							w.Drop{
								Trigger: w.Button{Label: "1 failed", Menu: true, Kind: w.ButtonDanger, Small: true},
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuRefusal{
											Mark:   w.Mark{Colour: "#3F6E8C", Size: w.MarkH11},
											Source: "nextcloud · Работа",
											Answer: "REPORT /remote.php/dav/calendars/mix/rabota/ → 503 Service Unavailable",
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
					w.ListGroup{Title: w.Rubric{Text: "Overdue", Form: w.RubricInline}, Count: &w.Num{Value: 1, Alert: true}},
					w.List{
						Rows: []w.ListRow{
							w.DueRow{
								Tick:     w.Tick{Label: "Done"},
								Mark:     w.Mark{Colour: "#4A6B52"},
								Title:    []w.Inline{w.Text("Отправить документы в МФЦ")},
								Due:      "7 Aug",
								Progress: "—",
								Source:   "nextcloud · Дела",
								Overdue:  true,
								Href:     "#",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "This week", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.DueRow{
								Tick:     w.Tick{Label: "Done"},
								Mark:     w.Mark{Colour: "#4A6B52"},
								Title:    []w.Inline{w.Text("Собрать чемодан")},
								Due:      "12 Aug 18:00",
								Progress: "—",
								Source:   "nextcloud · Дела",
								Href:     "#",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "No due date", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.DueRow{
								Tick:     w.Tick{Done: true, Label: "Done"},
								Mark:     w.Mark{Colour: "#4A6B52"},
								Title:    []w.Inline{w.Text("Продлить страховку")},
								Due:      "—",
								Progress: "100%",
								Source:   "nextcloud · Дела",
								Href:     "#",
							},
						},
					},
					w.Empty{
						Parts: []w.EmptyPart{
							w.Text("What arrived in time is shown. Tasks from the unavailable list are missing — not “there are none”, but “we have not seen them”."),
						},
						Variant: w.EmptySmall,
					},
				},
			},
		},
	}
}

// States21 is frame «Список задач пуст» (states-21).
func States21() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 3,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked: w.TickedPart,
							Kind:   w.TickRowAll,
							Mark:   w.Mark{Kind: w.MarkMulti},
							Label:  "Combined list",
							Meta:   "4",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "4"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "0", Current: true},
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
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "дом"}, w.Tag{Label: "релиз"}, w.Tag{Label: "read"}},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New list"},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#3F6E8C", Size: w.MarkH20}, Text: "Работа"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · 0 tasks")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New task", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Empty{
						Parts: []w.EmptyPart{
							w.Strong("This list is empty."),
							w.Text(" Add a task, or import a file of them. "),
							w.Button{Label: "New task", Icon: "plus", Kind: w.ButtonPrimary},
							w.Text(" "),
							w.Button{Label: "Import", Icon: "import"},
						},
					},
				},
			},
		},
	}
}

// States22 is frame «Список задач · сервер не ответил» (states-22).
func States22() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 3,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked: w.TickedPart,
							Kind:   w.TickRowAll,
							Mark:   w.Mark{Kind: w.MarkMulti},
							Label:  "Combined list",
							Meta:   "3",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Current: true, Error: true},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "3"},
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
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "дом"}, w.Tag{Label: "релиз"}, w.Tag{Label: "read"}},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New list"},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH20}, Text: "Дела"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New task", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Strong("Дела is unavailable."),
							w.Text(" The server did not answer as expected, so nothing on this screen comes from it. "),
							w.Button{Label: "Try again", Small: true},
						},
						Tone: w.MsgAlert,
					},
					w.MainDiag{
						Diag: w.Diag{Text: "REPORT /remote.php/dav/calendars/mix/dela/\nnextcloud\ndav: HTTP 503: the server is busy"},
					},
				},
			},
		},
	}
}

// States23 is frame «Список задач только для чтения» (states-23).
func States23() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 3,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked: w.TickedPart,
							Kind:   w.TickRowAll,
							Mark:   w.Mark{Kind: w.MarkMulti},
							Label:  "Combined list",
							Meta:   "7",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "4"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "3"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Subscriptions", Items: []w.RubricPart{w.Badge{Label: "external"}}},
						w.TickRow{
							Kind:    w.TickRowExternal,
							Mark:    w.Mark{Colour: "#8C6239"},
							Label:   "Праздники РФ",
							Meta:    "2",
							Current: true,
							Trail:   w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "дом"}, w.Tag{Label: "релиз"}, w.Tag{Label: "read"}},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New list"},
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
						Title: &w.HeadTitle{
							Mark:  &w.Mark{Colour: "#8C6239", Size: w.MarkH20},
							Text:  "Праздники РФ",
							Badge: &w.Badge{Label: "read-only"},
						},
						Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Subscriptions · 2 tasks · a feed can only be read")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{w.Button{Label: "Export", Icon: "export"}, w.Button{Label: "Print", Icon: "print"}},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter by word", Label: "Filter by word"},
								Waiting: true,
								Variant: w.FilterH26W200,
							},
							w.Seg{Options: []w.SegOption{{Label: "Open 2", On: true}, {Label: "Done 0"}, {Label: "All 2"}}},
							w.Seg{
								Options:   []w.SegOption{{Label: "By due", On: true}, {Label: "By priority"}, {Label: "Recently changed"}},
								Secondary: true,
							},
							w.Button{Label: "Columns", Icon: "cols"},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "This week", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.DueRow{
								Tick:     w.Tick{Sign: true},
								Mark:     w.Mark{Colour: "#8C6239"},
								Title:    []w.Inline{w.Text("День Воздушного флота — поздравить отца")},
								Due:      "12 Aug",
								Progress: "—",
								Single:   true,
								Href:     "#",
							},
							w.DueRow{
								Tick:     w.Tick{Sign: true},
								Mark:     w.Mark{Colour: "#8C6239"},
								Title:    []w.Inline{w.Text("День археолога — открытка Вере")},
								Due:      "15 Aug",
								Progress: "—",
								Single:   true,
								Href:     "#",
							},
						},
					},
				},
			},
		},
	}
}

// States24 is frame «Задача · сервер не сохранил свойства» (states-24).
func States24() Screen {
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "4", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "2"},
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
						Title: &w.HeadTitle{Text: "Отправить документы в МФЦ"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Badge{Label: "overdue by 3 days", Tone: w.BadgeAlert}}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Field{
									Parts: []w.FieldPart{
										w.Drop{
											Trigger: w.Pick{Value: "nextcloud · Дела", Mark: &w.Mark{Colour: "#4A6B52"}},
											Menu: w.Menu{
												Items: []w.MenuPart{
													w.MenuItem{
														Mark:  &w.Mark{Colour: "#4A6B52"},
														Label: "nextcloud · Дела",
														On:    true,
														Name:  "list",
														Value: "nextcloud/dela",
														Of:    "task-edit-loss",
													},
													w.MenuItem{
														Mark:  &w.Mark{Colour: "#3F6E8C"},
														Label: "nextcloud · Работа",
														Name:  "list",
														Value: "nextcloud/rabota",
														Of:    "task-edit-loss",
													},
												},
												Width: w.MenuWidth290,
												Pop:   w.MenuUnder,
											},
										},
									},
									Width: w.FieldMD,
								},
								w.Button{Label: "Cancel"},
								w.Button{Label: "Save", Kind: w.ButtonPrimary},
							},
						},
					},
					w.Runbar{
						Meter:    w.Meter{Percent: 0},
						Progress: "task-edit-loss",
						Parts:    []w.RunbarPart{w.Text("Saving · "), w.Mono{Text: "0 of 0 MB"}},
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Strong("The change was saved, but the server did not keep these properties."),
							w.Text(" Other clients will not see them either: "),
							w.Mono{Text: "X-APPLE-SORT-ORDER"},
							w.Text(", "),
							w.Mono{Text: "X-OC-HIDESUBTASKS"},
							w.Text("."),
						},
						Space: w.MsgFlushTop2,
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Summary",
								Parts: []w.FieldPart{w.Input{Value: "Отправить документы в МФЦ", Label: "Summary", Focus: true}},
								Width: w.FieldXL,
							},
							w.Field{
								Label: "Description",
								Parts: []w.FieldPart{w.Area{Text: "Опись, копия паспорта и заявление. Приём до 18:00, окно 4.", Label: "Text"}},
								Width: w.FieldFull,
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "When"},
								Items: []w.FieldSetPart{
									w.Field{Parts: []w.FieldPart{w.Check{On: true, Parts: []w.CheckPart{w.Text("All day")}}}},
									w.Field{
										Label: "Start date",
										Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Placeholder: "YYYY-MM-DD", Label: "Start date", Mono: true}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Start time",
										Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Placeholder: "--:--", Label: "Start time", Mono: true, Disabled: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Due date",
										Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Value: "2026-08-07", Label: "Due date", Mono: true}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Due time",
										Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Placeholder: "--:--", Label: "Due time", Mono: true, Disabled: true}},
										Width: w.FieldXS,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Progress"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Status", Parts: []w.FieldPart{w.Pick{Value: "Needs action"}}, Width: w.FieldSM},
									w.Field{Label: "Priority", Parts: []w.FieldPart{w.Pick{Value: "High · 1"}}, Width: w.FieldSM},
									w.Field{
										Label: "Percent complete",
										Parts: []w.FieldPart{w.Input{Value: "0", Label: "Percent complete", Mono: true}},
										Width: w.FieldSM,
									},
									w.Field{Label: "\u00a0", Parts: []w.FieldPart{w.Check{Parts: []w.CheckPart{w.Text("Done")}}}},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Reminders"},
								Items: []w.FieldSetPart{
									w.Field{Label: "First", Parts: []w.FieldPart{w.Pick{Value: "1 day before"}}, Width: w.FieldMD},
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
													Parts:   []w.TagInputPart{},
													Columns: true,
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
								Rubric: w.Rubric{Text: "Other"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Location",
										Parts: []w.FieldPart{w.Input{Placeholder: "optional", Label: "Location"}},
										Width: w.FieldMD,
									},
									w.Field{Label: "Tags", Parts: []w.FieldPart{w.Input{Value: "дом", Label: "Tags"}}, Width: w.FieldMD},
									w.Field{
										Label: "Link",
										Parts: []w.FieldPart{w.Input{Placeholder: "optional", Label: "Link"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Related to",
										Parts: []w.FieldPart{
											w.TagInput{
												Parts: []w.TagInputPart{
													w.Tag{Label: "Забрать документы", Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, Removable: true},
												},
												Add: &w.TagAdd{Name: "related.add", Placeholder: "Title or UID", Label: "Add a link"},
											},
										},
										Width: w.FieldLG,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Attachments"},
								Items: []w.FieldSetPart{
									w.Line{
										Items: []w.LinePart{
											w.Drop{
												Trigger: w.Button{Label: "Attach", Menu: true, Icon: "clip"},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{
															Icon: "upload",
															File: &w.FileField{Form: "", Name: "file", Label: "From this computer", List: "task-loss-attachments"},
															Row: w.Line{
																Items: []w.LinePart{
																	w.Faced{
																		Face:  w.FaceLarge,
																		Items: []w.LinePart{w.Text("{name} "), w.Dim{Parts: []w.Inline{w.Text("{size}")}}},
																	},
																	w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Removes: true},
																},
																Gap: w.LineGap8,
															},
															Label: "From this computer",
														},
														w.MenuItem{Icon: "file", Label: "From Files"},
													},
													Width: w.MenuWidth212,
													Pop:   w.MenuUnder,
												},
											},
											w.Dim{Parts: []w.Inline{w.Text("Nothing attached")}, Size: w.Dim11_5, Variant: w.DimWrap, Empty: true},
										},
										Gap:   w.LineGap8,
										Block: true,
										ID:    "task-loss-attachments",
										Place: w.Grown,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Repeats"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Frequency",
										Parts: []w.FieldPart{
											w.Pick{Value: "Does not repeat"},
											w.FieldHint{Parts: []w.Inline{w.Text("Repeats from the start date, or from the due date when there is none.")}},
										},
										Width: w.FieldSM,
									},
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
													{Label: "MO", Value: "MO"},
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
										Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Label: "Until", Mono: true}},
										Width: w.FieldSM,
									},
									w.Field{Label: "Count", Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Count"}}, Width: w.FieldXS},
									w.Field{
										Label: "Advanced RRULE",
										Parts: []w.FieldPart{
											w.Input{Placeholder: "—", Label: "Advanced RRULE", Mono: true},
											w.FieldHint{Parts: []w.Inline{w.Text("The advanced value takes precedence when set.")}},
										},
										Width: w.FieldXL,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Source"},
								Items: []w.FieldSetPart{
									w.Fields{
										Items: []w.FieldsItem{
											{
												Term:  "Path",
												Value: []w.ValuePart{w.Text("/remote.php/dav/calendars/mix/dela/aa10.ics")},
												Face:  w.ValueMono,
											},
											{Term: "Read", Value: []w.ValuePart{w.Text("2 min ago · not changed on server")}},
										},
										Place: w.Grown,
									},
								},
							},
						},
					},
					w.Danger{Button: w.Button{Label: "Delete task", Icon: "trash", Kind: w.ButtonDangerSolid}},
				},
			},
		},
	}
}

// States25 is frame «Заметки · ни одна тетрадь не отмечена» (states-25).
func States25() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.Button{Label: "Refresh", Icon: "refresh", Title: "Nothing is ticked to ask.", State: w.ButtonDisabled},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 4,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedOff,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined notebook",
							Meta:    "—",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2", Hidden: true},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "2", Hidden: true},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#3F6E8C"},
							Label:  "Общие с семьёй",
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("0 of 3 notebooks · 0 accounts")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New note", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
							},
						},
					},
					w.Empty{
						Parts: []w.EmptyPart{
							w.Strong("No notebooks are ticked."),
							w.Text(" Tick the notebooks to show in the rail, or all of them at once. "),
							w.Button{Label: "Tick every notebook", Kind: w.ButtonPrimary},
						},
					},
				},
			},
		},
	}
}

// States26 is frame «Заметки · опрос идёт» (states-26).
func States26() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.Button{Label: "Refresh", Icon: "refresh", Title: "The poll is still running.", State: w.ButtonDisabled},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
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
							Meta:    "—",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "—"},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#3F6E8C"},
							Label:  "Общие с семьёй",
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "встречи"}, w.Tag{Label: "книги"}, w.Tag{Label: "релиз"}},
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 of 3 notebooks · VJOURNAL")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New note", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter by word", Label: "Filter by word"},
								Waiting: true,
								Variant: w.FilterH26W200,
							},
							w.Seg{
								Options:   []w.SegOption{{Label: "Newest", On: true}, {Label: "Oldest"}, {Label: "Title"}, {Label: "Recently changed"}},
								Secondary: true,
							},
							w.Button{Label: "Columns", Icon: "cols"},
						},
						Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "2"}}, Secondary: true},
					},
					w.Runbar{
						Meter: w.Meter{Percent: 62},
						Parts: []w.RunbarPart{w.Text("Polling · 1 of 2 answered · 2 notes"), w.Mono{Text: "0.8 s", Place: w.Pushed}},
					},
					w.ListGroup{Title: w.Rubric{Text: "August 2026", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Title:   []w.Inline{w.Text("Заметка со встречи 10.08")},
								Excerpt: "Договорились заморозить ветку в среду; письмо пользователям пишет Анна…",
								Date:    "10 Aug",
								Source:  "nextcloud · Записи",
							},
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Title:   []w.Inline{w.Text("Что почитать про CRDT")},
								Excerpt: "Кляйн, глава 5; статья Шапиро 2011; демка automerge…",
								Date:    "8 Aug",
								Source:  "nextcloud · Записи",
							},
						},
					},
				},
			},
		},
	}
}

// States27 is frame «Заметки · источник недоступен» (states-27).
func States27() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
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
							Meta:    "2",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Error: true},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#3F6E8C"},
							Label:  "Общие с семьёй",
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "встречи"}, w.Tag{Label: "книги"}, w.Tag{Label: "релиз"}},
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 of 3 notebooks · 2 notes · VJOURNAL")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New note", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter by word", Label: "Filter by word"},
								Waiting: true,
								Variant: w.FilterH26W200,
							},
							w.Seg{
								Options:   []w.SegOption{{Label: "Newest", On: true}, {Label: "Oldest"}, {Label: "Title"}, {Label: "Recently changed"}},
								Secondary: true,
							},
							w.Button{Label: "Columns", Icon: "cols"},
						},
						Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "2"}}, Secondary: true},
					},
					w.Runbar{
						Meter: w.Meter{Percent: 62, Alert: true},
						Parts: []w.RunbarPart{
							w.Text("Polled 1 of 2 sources · 2 notes · "),
							w.Drop{
								Trigger: w.Button{Label: "1 failed", Menu: true, Kind: w.ButtonDanger, Small: true},
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuRefusal{
											Mark:   w.Mark{Colour: "#6B4E7A", Size: w.MarkH11},
											Source: "nextcloud · Дневник",
											Answer: "REPORT /remote.php/dav/calendars/mix/dnevnik/ → 503 Service Unavailable",
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
					w.ListGroup{Title: w.Rubric{Text: "August 2026", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Title:   []w.Inline{w.Text("Заметка со встречи 10.08")},
								Excerpt: "Договорились заморозить ветку в среду; письмо пользователям пишет Анна…",
								Date:    "10 Aug",
								Source:  "nextcloud · Записи",
							},
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Title:   []w.Inline{w.Text("Что почитать про CRDT")},
								Excerpt: "Кляйн, глава 5; статья Шапиро 2011; демка automerge…",
								Date:    "8 Aug",
								Source:  "nextcloud · Записи",
							},
						},
					},
					w.Empty{
						Parts: []w.EmptyPart{
							w.Text("What arrived in time is shown. Notes from the unavailable notebook are missing — not “there are none”, but “we have not seen them”."),
						},
						Variant: w.EmptySmall,
					},
				},
			},
		},
	}
}

// States28 is frame «Тетрадь пуста» (states-28).
func States28() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 4,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked: w.TickedPart,
							Kind:   w.TickRowAll,
							Mark:   w.Mark{Kind: w.MarkMulti},
							Label:  "Combined notebook",
							Meta:   "2",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "0", Current: true},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#3F6E8C"},
							Label:  "Общие с семьёй",
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "встречи"}, w.Tag{Label: "книги"}, w.Tag{Label: "релиз"}},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#6B4E7A", Size: w.MarkH20}, Text: "Дневник"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · 0 notes · VJOURNAL")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New note", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
							},
						},
					},
					w.Empty{
						Parts: []w.EmptyPart{
							w.Strong("This notebook is empty."),
							w.Text(" Write a note, or import a file of them. "),
							w.Button{Label: "New note", Icon: "plus", Kind: w.ButtonPrimary},
							w.Text(" "),
							w.Button{Label: "Import", Icon: "import"},
						},
					},
				},
			},
		},
	}
}

// States29 is frame «Тетрадь · сервер не ответил» (states-29).
func States29() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 4,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked: w.TickedPart,
							Kind:   w.TickRowAll,
							Mark:   w.Mark{Kind: w.MarkMulti},
							Label:  "Combined notebook",
							Meta:   "2",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Current: true, Error: true},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "2"},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#3F6E8C"},
							Label:  "Общие с семьёй",
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "встречи"}, w.Tag{Label: "книги"}, w.Tag{Label: "релиз"}},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH20}, Text: "Записи"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New note", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
							},
						},
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Strong("Записи is unavailable."),
							w.Text(" The server did not answer as expected, so nothing on this screen comes from it. "),
							w.Button{Label: "Try again", Small: true},
						},
						Tone: w.MsgAlert,
					},
					w.MainDiag{
						Diag: w.Diag{Text: "REPORT /remote.php/dav/calendars/mix/zapisi/\nnextcloud\ndav: HTTP 503: the server is busy"},
					},
				},
			},
		},
	}
}

// States30 is frame «Тетрадь только для чтения» (states-30).
func States30() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 4,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked: w.TickedPart,
							Kind:   w.TickRowAll,
							Mark:   w.Mark{Kind: w.MarkMulti},
							Label:  "Combined notebook",
							Meta:   "6",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "2"},
						w.TickRow{
							Mark:    w.Mark{Colour: "#3F6E8C"},
							Label:   "Общие с семьёй",
							Meta:    "2",
							Current: true,
							Trail:   w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{
						w.Tag{Label: "all", On: true},
						w.Tag{Label: "встречи"},
						w.Tag{Label: "книги"},
						w.Tag{Label: "релиз"},
						w.Tag{Label: "рецепты"},
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
						Title: &w.HeadTitle{
							Mark:  &w.Mark{Colour: "#3F6E8C", Size: w.MarkH20},
							Text:  "Общие с семьёй",
							Badge: &w.Badge{Label: "read-only"},
						},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{w.Text("nextcloud · 2 notes · VJOURNAL · the server granted no write access")},
							},
						},
						Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Export", Icon: "export"}}},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter by word", Label: "Filter by word"},
								Waiting: true,
								Variant: w.FilterH26W200,
							},
							w.Seg{
								Options:   []w.SegOption{{Label: "Newest", On: true}, {Label: "Oldest"}, {Label: "Title"}, {Label: "Recently changed"}},
								Secondary: true,
							},
							w.Button{Label: "Columns", Icon: "cols"},
						},
						Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "2"}}, Secondary: true},
					},
					w.ListGroup{Title: w.Rubric{Text: "August 2026", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#3F6E8C"},
								Title:   []w.Inline{w.Text("Список на дачу")},
								Excerpt: "Семена, шланг, две банки краски для забора…",
								Date:    "6 Aug",
								Single:  true,
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "July 2026", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#3F6E8C"},
								Title:   []w.Inline{w.Text("Рецепт бабушкиного пирога")},
								Excerpt: "Мука — три стакана, кефир — стакан, яблоки кислые…",
								Date:    "19 Jul",
								Single:  true,
							},
						},
					},
				},
			},
		},
	}
}

// States31 is frame «Заметка · сервер не сохранил свойства» (states-31).
func States31() Screen {
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
					{Icon: "back", Label: "Back to list", Short: "Back", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "2"},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#3F6E8C"},
							Label:  "Общие с семьёй",
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Neighbours"},
						w.Tree{
							Rows: []w.TreeRow{
								{Icon: "note", Label: "Заметка со встречи"},
								{Icon: "note", Label: "Что почитать про CRDT"}},
							On:      1,
							Variant: w.TreePadded,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Docbar{
						Items: []w.DocbarPart{
							w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "note-loss"},
							w.Button{Label: "Cancel", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012"},
							w.Field{
								Parts: []w.FieldPart{
									w.Drop{
										Trigger: w.Pick{Value: "nextcloud · Записи", Mark: &w.Mark{Colour: "#4A6B52"}},
										Menu: w.Menu{
											Items: []w.MenuPart{
												w.MenuItem{
													Mark:  &w.Mark{Colour: "#4A6B52"},
													Label: "nextcloud · Записи",
													On:    true,
													Name:  "notebook",
													Value: "nextcloud/zapisi",
													Of:    "note-loss",
												},
												w.MenuItem{
													Mark:  &w.Mark{Colour: "#6B4E7A"},
													Label: "nextcloud · Дневник",
													Name:  "notebook",
													Value: "nextcloud/dnevnik",
													Of:    "note-loss",
												},
												w.MenuItem{
													Mark:   &w.Mark{Colour: "#3F6E8C"},
													Trail:  w.Badge{Label: "ro", Place: w.Pushed},
													Label:  "nextcloud · Общие с семьёй",
													Reason: "Read-only: this account cannot write here",
													Name:   "notebook",
													Value:  "nextcloud/semya",
													Of:     "note-loss",
												},
											},
											Width: w.MenuWidth290,
											Pop:   w.MenuUnder,
										},
									},
								},
								Width: w.FieldMD,
							},
							w.BarSep{},
							w.Button{Label: "B", Bold: true},
							w.Button{Label: "I", Variant: w.ButtonItalic},
							w.Button{Label: "H2", Tertiary: true},
							w.Button{Label: "• list", Tertiary: true},
							w.Button{Folds: true, Label: "task", Icon: "task", Title: "Task", Tertiary: true},
							w.Button{Icon: "link", Title: "Link", Tertiary: true},
							w.Drop{
								Trigger: w.Button{Folds: true, Label: "Attach", Menu: true, Icon: "clip", Title: "Attach"},
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuItem{
											Icon: "upload",
											File: &w.FileField{Form: "note-loss", Name: "file", Label: "From this computer", List: "note-loss-attachments"},
											Row: w.DocAttach{
												Name:   "{name}",
												Size:   w.Dim{Parts: []w.Inline{w.Text("· {size}")}},
												Remove: &w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Removes: true},
											},
											Label: "From this computer",
										},
										w.MenuItem{Icon: "file", Label: "From Files"},
									},
									Width: w.MenuWidth212,
									Pop:   w.MenuUnder,
								},
							},
						},
						More: &w.Button{Icon: "dots", Title: "More"},
						End: []w.DocbarEnd{
							w.Check{Parts: []w.CheckPart{w.Text("Auto-save")}, Tertiary: true},
							w.DraftNote{Text: "Draft in this window ·", Time: w.Mono{Text: "11:42"}, Secondary: true},
							w.Kbd{Keys: "Ctrl+S", Secondary: true},
							w.Seg{Options: []w.SegOption{{Label: "Full width", On: true}, {Label: "Reading width"}}, Secondary: true},
							w.Button{Folds: true, Label: "Focus", Icon: "focus", Title: "Focus"},
						},
						Edit: true,
					},
					w.Runbar{
						Meter:    w.Meter{Percent: 0},
						Progress: "note-loss",
						Parts:    []w.RunbarPart{w.Text("Saving · "), w.Mono{Text: "0 of 0 MB"}},
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Strong("The change was saved, but the server did not keep these properties."),
							w.Text(" Other clients will not see them either: "),
							w.Mono{Text: "X-EVOLUTION-MEMO-COLOR"},
							w.Text(", "),
							w.Mono{Text: "X-CARREL-UNKNOWN-FLAG"},
							w.Text("."),
						},
						Space: w.MsgFlushTop2,
					},
					w.DocLayout{
						Items: []w.DocLayoutPart{
							w.Doc{
								Items: []w.DocPart{
									w.Area{Text: "Заметка со встречи 10.08", Variant: w.AreaTitle, Name: "summary", Label: "Text"},
									w.DocFields{
										Items: []w.Field{
											w.Field{
												Label: "Date",
												Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "dtstart.date", Value: "2026-08-10", Label: "Date", Mono: true}},
												Width: w.FieldFix150,
											},
											w.Field{
												Label: "Time",
												Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Name: "dtstart.time", Value: "11:05", Label: "Time", Mono: true}},
												Width: w.FieldFix130,
											},
											w.Field{
												Label: "Tags",
												Parts: []w.FieldPart{w.Input{Name: "categories", Value: "встречи, релиз", Label: "Tags"}},
												Width: w.FieldFlex3_320,
											},
											w.Field{
												Label: "Link",
												Parts: []w.FieldPart{w.Input{Name: "url", Placeholder: "optional", Label: "Link"}},
												Width: w.FieldFlex2_280,
											},
											w.Field{
												Label: "Related to",
												Parts: []w.FieldPart{
													w.TagInput{
														Parts: []w.TagInputPart{
															w.Tag{Label: "Синк-ап по релизу", Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, Removable: true},
															w.Tag{Label: "Дописать раздел про кэш", Mark: &w.Mark{Colour: "#3F6E8C", Size: w.MarkH11}, Removable: true},
														},
														Add: &w.TagAdd{Name: "related.add", Placeholder: "Title or UID", Label: "Add a link"},
													},
												},
												Width: w.FieldFlex3_340,
											},
										},
									},
									w.Area{
										Text:    "Договорились заморозить ветку в среду.\n\nПисьмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.\n\n## Что делаем\n\n- [ ] проверить changelog\n- [x] попросить у Артёма скриншоты\n\nПеренос старых вложений в новую папку — Виктор посмотрит, сколько их вообще, и скажет в четверг.",
										Variant: w.AreaBody,
										Focus:   true,
										Name:    "description",
										Label:   "Text",
									},
									w.DocNote{
										Parts: []w.Inline{
											w.Text("Properties Carrel does not know — "),
											w.Mono{Text: "X-JTX-COLOR"},
											w.Text(" and two more — are kept unchanged and not shown here."),
										},
									},
								},
								ID:     "note-loss",
								Action: "/app/notes/nextcloud/zapisi/9e31-40b7-c012",
							},
							w.Docmeta{
								Sections: []w.DocmetaSec{
									{
										Title: w.Rubric{
											Text: "Attachments",
											Items: []w.RubricPart{
												w.Drop{
													Trigger: w.Button{Label: "Attach", Menu: true, Icon: "clip", Small: true},
													Menu: w.Menu{
														Items: []w.MenuPart{
															w.MenuItem{
																Icon: "upload",
																File: &w.FileField{Form: "note-loss", Name: "file", Label: "From this computer", List: "note-loss-attachments"},
																Row: w.DocAttach{
																	Name:   "{name}",
																	Size:   w.Dim{Parts: []w.Inline{w.Text("· {size}")}},
																	Remove: &w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Removes: true},
																},
																Label: "From this computer",
															},
															w.MenuItem{Icon: "file", Label: "From Files"},
														},
														Width: w.MenuWidth212,
														Pop:   w.MenuUnder,
													},
												},
											},
										},
										Items: []w.DocmetaPart{
											w.DocAttach{
												Name:   "2026-08-10-sync-board.png",
												Size:   w.Dim{Parts: []w.Inline{w.Text("· 412 KB")}},
												Where:  w.Badge{Label: "in the note"},
												Remove: &w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true},
												Href:   "#",
											},
											w.DocAttach{
												Name:   "запись-встречи.m4a",
												Size:   w.Dim{Parts: []w.Inline{w.Text("· 24 MB")}},
												Where:  w.Badge{Label: "Documents / carrel"},
												Remove: &w.Button{Label: "Unlink", Small: true, Title: "Remove the link; the file stays in Documents / carrel"},
												Href:   "#",
											},
										},
										ID: "note-loss-attachments",
									},
									{
										Title: w.Rubric{Text: "Source"},
										Items: []w.DocmetaPart{w.Diag{Text: "UID   9e31-40b7-c012\nRead  2 min ago"}},
									},
								},
								Danger: &w.Danger{Button: w.Button{Label: "Delete note", Kind: w.ButtonDangerSolid}, Space: w.DangerAtEnd},
							},
						},
					},
				},
			},
		},
	}
}

// States32 is frame «Папка пуста» (states-32).
func States32() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 5,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked: w.TickedPart,
							Kind:   w.TickRowRoot,
							Icon:   "file",
							Label:  "Selected collections",
							Meta:   "2 of 3",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "sftpgo · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Documents", Meta: "rw", Current: true},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#9A9280"},
							Label:  "Public",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Files", Meta: "rw"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Shortcuts"},
						w.Tree{
							Rows: []w.TreeRow{
								{Icon: "file", Label: "Attachments", Trail: w.TreeNote{Text: "Documents / carrel"}},
								{Icon: "link", Label: "Published", Trail: w.TreeNote{Text: "2"}},
							},
							On:      1,
							Variant: w.TreePadded,
						},
					},
				},
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
						Title: &w.HeadTitle{Text: "2026-08"},
						Crumbs: &w.Crumbs{
							Path:    []w.Crumb{{Label: "Documents", Href: "#"}, {Label: "carrel", Href: "#"}, {Label: "attachments", Href: "#"}},
							Current: "2026-08",
						},
						Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("sftpgo · vps · 0 items")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Upload", Menu: true, Icon: "upload", Kind: w.ButtonPrimary},
								w.Button{Label: "New folder", Icon: "plus"},
								w.Button{Icon: "refresh", Title: "Refresh"},
								w.Button{Icon: "panel", Title: "Choose one file to see its details.", State: w.ButtonDisabled},
								w.Button{Icon: "dots", Title: "More"},
							},
						},
					},
					w.Empty{
						Parts: []w.EmptyPart{
							w.Strong("This folder is empty."),
							w.Text(" Upload files into it, or drop them onto the window. "),
							w.Button{
								Label: "Upload",
								Icon:  "upload",
								Kind:  w.ButtonPrimary,
								File: &w.FileField{
									Form:  "post-upload",
									Name:  "file",
									Label: "Files",
									Size:  "size",
									Taken: "/app/files/sftpgo/documents/upload/carrel/attachments/2026-08?op=taken",
									Many:  true,
								},
							},
							w.Text(" "),
							w.Button{Label: "Send", Form: "post-upload", NoScript: true},
							w.Text(" "),
							w.Button{Label: "New folder", Icon: "plus"},
						},
					},
					w.PostForm{
						ID:        "post-upload",
						Action:    "/app/files/sftpgo/documents/upload/carrel/attachments/2026-08",
						Fields:    []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
						Multipart: true,
					},
					w.Queue{
						Title:    "Uploading {done} of {count}",
						Total:    w.Mono{Text: "{sent}", Variant: w.MonoDusty, Place: w.Pushed},
						Toggle:   w.Button{Label: "Hide", Small: true},
						Progress: "post-upload",
						Row:      &w.QueueItem{Name: "{name}", Meta: "{size} · {percent}%", Meter: &w.Meter{Percent: 0, Track: 100}},
						Taken: &w.QueueItem{
							Name:   "{name}",
							Meta:   "name taken",
							Failed: true,
							Acts: []w.Button{
								w.Button{Label: "Keep both", Small: true, Name: "answer", Value: "keep"},
								w.Button{Label: "Replace", Small: true, Name: "answer", Value: "replace"},
								w.Button{Label: "Skip", Small: true, Name: "answer", Value: "skip"},
							},
						},
						Fold: &w.QueueFold{Hide: "Hide", Show: "Show", Total: "{percent}%"},
					},
				},
			},
		},
	}
}

// States33 is frame «Хранилище · сервер не ответил» (states-33).
func States33() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 5,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked: w.TickedPart,
							Kind:   w.TickRowRoot,
							Icon:   "file",
							Label:  "Selected collections",
							Meta:   "2 of 3",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "sftpgo · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Documents", Current: true, Error: true},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#9A9280"},
							Label:  "Public",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Files", Meta: "rw"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Shortcuts"},
						w.Tree{
							Rows: []w.TreeRow{
								{Icon: "file", Label: "Attachments", Trail: w.TreeNote{Text: "Documents / carrel"}},
								{Icon: "link", Label: "Published", Trail: w.TreeNote{Text: "2"}},
							},
							On:      1,
							Variant: w.TreePadded,
						},
					},
				},
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
						Title: &w.HeadTitle{Text: "Documents"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("sftpgo · vps")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Upload", Menu: true, Icon: "upload", Kind: w.ButtonPrimary},
								w.Button{Label: "New folder", Icon: "plus"},
								w.Button{Icon: "refresh", Title: "Refresh"},
								w.Button{Icon: "panel", Title: "Choose one file to see its details.", State: w.ButtonDisabled},
								w.Button{Icon: "dots", Title: "More"},
							},
						},
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Strong("Documents is unavailable."),
							w.Text(" The server did not answer as expected, so nothing on this screen comes from it. "),
							w.Button{Label: "Try again", Small: true},
						},
						Tone: w.MsgAlert,
					},
					w.MainDiag{Diag: w.Diag{Text: "PROPFIND /dav/Documents/\nsftpgo · vps\ndav: HTTP 503: the server is busy"}},
				},
			},
		},
	}
}

// States34 is frame «Хранилище только для чтения» (states-34).
func States34() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 5,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{Kind: w.TickRowRoot, Icon: "file", Label: "Selected collections", Meta: "3 of 3"}},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "sftpgo · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Documents", Meta: "rw"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Public", Current: true, Trail: w.Badge{Label: "ro"}},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Files", Meta: "rw"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Shortcuts"},
						w.Tree{
							Rows: []w.TreeRow{
								{Icon: "file", Label: "Attachments", Trail: w.TreeNote{Text: "Documents / carrel"}},
								{Icon: "link", Label: "Published", Trail: w.TreeNote{Text: "2"}},
							},
							Variant: w.TreePadded,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "This storage is read-only.", State: w.ButtonDisabled},
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
						Title: &w.HeadTitle{Text: "Public"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("sftpgo · vps · 2 items · "), w.Badge{Label: "read-only"}}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Upload", Menu: true, Icon: "upload", Kind: w.ButtonPrimary, State: w.ButtonDisabled},
								w.Button{Label: "New folder", Icon: "plus", State: w.ButtonDisabled},
								w.Button{Icon: "refresh", Title: "Refresh"},
								w.Button{Icon: "panel", Title: "Choose one file to see its details.", State: w.ButtonDisabled},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.BarSep{},
							w.Button{
								Folds: true,
								Label: "Download",
								Icon:  "download",
								Title: "Download",
								Form:  "files-sel",
								Name:  "op",
								Value: "download",
							},
							w.Button{
								Folds: true,
								Label: "Rename",
								Icon:  "pencil",
								Title: "Public is read-only.",
								Form:  "files-sel",
								Name:  "op",
								Value: "rename",
								State: w.ButtonDisabled,
							},
							w.Button{
								Folds: true,
								Label: "Move",
								Icon:  "move",
								Title: "Public is read-only.",
								Form:  "files-sel",
								Name:  "op",
								Value: "move",
								State: w.ButtonDisabled,
							},
							w.Button{Folds: true, Label: "Copy", Icon: "dup", Title: "Copy", Form: "files-sel", Name: "op", Value: "copy"},
							w.Button{
								Folds: true,
								Label: "Delete",
								Icon:  "trash",
								Kind:  w.ButtonDangerSolid,
								Title: "Public is read-only.",
								Form:  "files-sel",
								Name:  "op",
								Value: "delete",
								State: w.ButtonDisabled,
							},
						},
						Count:    &w.BarCount{Text: "With the ticked"},
						Variant:  w.BarSelection,
						NoScript: true,
					},
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter this folder", Label: "Filter this folder"},
								Waiting: true,
								Variant: w.FilterH26W190,
							},
							w.Seg{
								Options: []w.SegOption{
									{Label: "Name", Href: "/app/files/sftpgo/public?sort=name", On: true},
									{Label: "Size", Href: "/app/files/sftpgo/public?sort=size"},
									{Label: "Changed", Href: "/app/files/sftpgo/public?sort=changed"},
								},
								Secondary: true,
							},
							w.Seg{
								Options: []w.SegOption{
									{Icon: "menu", Href: "/app/files/sftpgo/public?view=list", On: true, Title: "List"},
									{Icon: "grid", Href: "/app/files/sftpgo/public?view=tiles", Title: "Tiles"},
								},
							},
							w.Button{Label: "Columns", Icon: "cols"},
							w.Drop{
								Trigger: w.Button{Icon: "dots", Title: "More"},
								More:    true,
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuItem{Label: "Name", On: true, Href: "/app/files/sftpgo/public?sort=name"},
										w.MenuItem{Label: "Size", Href: "/app/files/sftpgo/public?sort=size"},
										w.MenuItem{Label: "Changed", Href: "/app/files/sftpgo/public?sort=changed"},
									},
									Pop:    w.MenuUnderEnd,
									Inline: true,
								},
							},
						},
						Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "2 items · 5.3 MB"}}, Secondary: true},
					},
					w.Table{
						Cols: []w.TableCol{
							{Width: w.ColTick, Tick: ptr(w.TickedOff)},
							{Label: "Name"},
							{Label: "Size", Width: w.ColSize, Rank: w.Rank3},
							{Label: "Type", Width: w.ColType, Rank: w.Rank4, Secondary: true},
							{Label: "Changed", Width: w.ColTime},
							{Width: w.ColNum},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{Pick: ptr(w.TickedOff), Send: w.Send{Of: "files-sel", Name: "item", Value: "прайс-2026.pdf"}},
									{
										Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "прайс-2026.pdf", Href: "/app/files/sftpgo/public/прайс-2026.pdf"}},
									},
									{Parts: []w.CellPart{w.Text("1.1 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("application/pdf")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-07-28 10:15")}, Mono: true},
									{Parts: []w.CellPart{w.Button{Icon: "download", Small: true, Title: "Download"}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Pick: ptr(w.TickedOff), Send: w.Send{Of: "files-sel", Name: "item", Value: "схема-проезда.png"}},
									{
										Parts: []w.CellPart{
											w.TableFile{Icon: "img", Name: "схема-проезда.png", Href: "/app/files/sftpgo/public/схема-проезда.png"},
										},
									},
									{Parts: []w.CellPart{w.Text("4.2 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/png")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-06-02 16:40")}, Mono: true},
									{Parts: []w.CellPart{w.Button{Icon: "download", Small: true, Title: "Download"}}, Acts: true},
								},
							},
						},
					},
					w.PostForm{
						ID:     "files-sel",
						Action: "/app/files/sftpgo/public/op",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
				},
			},
		},
	}
}

// States35 is frame «Файлы · ни одно хранилище не отмечено» (states-35).
func States35() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 5,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedOff,
							Kind:    w.TickRowRoot,
							Icon:    "file",
							Label:   "Selected collections",
							Meta:    "0 of 3",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "sftpgo · vps"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#4A6B52"}, Label: "Documents", Meta: "rw", Hidden: true},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#9A9280"},
							Label:  "Public",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Files", Meta: "rw", Hidden: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Shortcuts"},
						w.Tree{
							Rows: []w.TreeRow{
								{Icon: "file", Label: "Attachments", Trail: w.TreeNote{Text: "Documents / carrel"}},
								{Icon: "link", Label: "Published", Trail: w.TreeNote{Text: "2"}},
							},
							Variant: w.TreePadded,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "Open a folder to make one in it.", State: w.ButtonDisabled},
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
						Title: &w.HeadTitle{Text: "Selected collections", Size: w.HeadTitle18},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("0 of 3 storages")}}},
						Pad:   w.HeadPadBottom8,
					},
					w.Empty{
						Parts: []w.EmptyPart{
							w.Strong("No storages are ticked."),
							w.Text(" Tick the storages to show in the rail, or all of them at once. "),
							w.Button{Label: "Tick every storage", Kind: w.ButtonPrimary},
						},
					},
				},
			},
		},
	}
}
