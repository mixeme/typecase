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

// Phone2 is frame «Рейл открыт: источники на телефоне» (phone-2).
func Phone2() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 10}, Href: "#"},
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
							Href:    "/app/calendar",
							Meta:    "10",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Href: "/app/calendar/nextcloud/lichnyi", Meta: "5"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Href: "/app/calendar/nextcloud/rabota", Meta: "2"},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#A05A2C"},
							Label:  "Дни рождения",
							Href:   "/app/calendar/nextcloud/dni-rozhdeniya",
							Meta:   "—",
							Hidden: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Href: "/app/calendar/radicale/proekty", Meta: "2"},
						w.TickRow{
							Mark:  w.Mark{Colour: "#8C6239"},
							Label: "Дежурства",
							Href:  "/app/calendar/radicale/dezhurstva",
							Meta:  "1",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "baikal"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Href: "/app/calendar/baikal/obshchii", Error: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Subscriptions", Items: []w.RubricPart{w.Badge{Label: "external"}}},
						w.TickRow{
							Kind:  w.TickRowExternal,
							Mark:  w.Mark{Colour: "#8C6239"},
							Label: "Праздники РФ",
							Href:  "/app/calendar/prazdniki-rf/feed",
							Meta:  "3",
							Trail: w.Badge{Label: "ro"},
						},
						w.TickRow{
							Ticked: w.TickedOff,
							Kind:   w.TickRowExternal,
							Mark:   w.Mark{Colour: "#6B4E7A"},
							Label:  "Расписание лиги",
							Href:   "/app/calendar/raspisanie-ligi/feed",
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
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Split{
									Main: w.Button{Label: "New event", Icon: "plus", Kind: w.ButtonPrimary},
									More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
								},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter by word", Short: "Filter", Label: "Filter by word", Name: "filter"},
								Waiting: true,
								Action:  "/app/calendar",
								Swap:    "list-main",
								Variant: w.FilterH26W200,
							},
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
							w.Mono{Text: "from cache · 2 min ago", Place: w.Pushed},
						},
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
								Strong: true,
								Detail: "Переговорная 2",
								Badges: []w.Badge{w.Badge{Label: "4 attendees"}, w.Badge{Label: "repeats weekly"}},
								Source: "nextcloud · Личный",
								On:     true,
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
					w.Empty{
						Parts: []w.EmptyPart{
							w.Text("What arrived in time is shown. Records from the unavailable source are missing — not “there are none”, but “we have not seen them”."),
						},
						Variant: w.EmptySmall,
					},
				},
				ID: "list-main",
			},
			Side: &w.Side{
				Items: []w.SidePart{
					w.Head{
						Title: &w.HeadTitle{Text: "Синк-ап по релизу", Size: w.HeadTitle19},
						Subs: []w.HeadSub{
							w.HeadSub{Parts: []w.HeadSubPart{w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, w.Text("nextcloud · Личный")}},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Edit", Small: true},
								w.Button{Icon: "dots", Small: true, Title: "More"},
								w.Close{Label: "Close the panel"},
							},
						},
					},
					w.Fields{
						Items: []w.FieldsItem{
							{
								Term: "When",
								Value: []w.ValuePart{
									w.Mono{Text: "10:00 – 11:00"},
									w.Text(", Monday, 10 August "),
									w.Badge{Label: "created 12:00 Asia/Yekaterinburg", Title: "Created in another time zone"},
								},
							},
							{Term: "Repeats", Value: []w.ValuePart{w.Text("Weekly on Mon "), w.Badge{Label: "until 28 December"}}},
							{Term: "Reminders", Value: []w.ValuePart{w.Text("10 minutes before · 2 days before")}},
							{Term: "Where", Value: []w.ValuePart{w.Text("Переговорная 2")}},
							{Term: "Status", Value: []w.ValuePart{w.Text("Confirmed")}},
							{Term: "Tags", Value: []w.ValuePart{w.Tag{Label: "релиз"}, w.Text(" "), w.Tag{Label: "работа"}}},
						},
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
							w.Rubric{Text: "Attendees", Items: []w.RubricPart{w.Mono{Text: "4", Variant: w.MonoTight}}},
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{
										w.Mark{Colour: "#4A6B52", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "Михаил Енученко", Href: "#"},
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
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{
								Text: "Attachments",
								Items: []w.RubricPart{
									w.Drop{
										Trigger: w.Button{Label: "Attach", Menu: true, Icon: "clip", Small: true},
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
								},
							},
							w.PostForm{
								ID:        "post-attach",
								Action:    "/app/calendar/nextcloud/lichnyj/8f1c/edit?op=attach",
								Fields:    []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
								Multipart: true,
							},
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{
										w.Link{Text: "release-checklist.md", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("text/markdown · 3 KB")}},
										w.Text(" "),
										w.Badge{Label: "in the event"},
									},
								},
								Gap: w.StackGapTight,
							},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Related"},
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{
										w.Span{Face: w.FaceAccent, Text: "Заметка со встречи 10.08"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· note")}},
									},
								},
							},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Source"},
							w.Diag{
								Text: "Path  /remote.php/dav/calendars/mix/personal/8f1c.ics\nRead  2 min ago · changed on server 3 days ago",
							},
						},
					},
					w.Danger{
						Button: w.Button{Label: "Delete event", Kind: w.ButtonDangerSolid},
						Note:   "Series is edited and deleted whole.",
					},
				},
			},
			RailOpen: true,
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
										w.Mark{Colour: "#4A6B52", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "Михаил Енученко", Href: "#"},
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
						Action: "/app/files/sftpgo/files/op/documents/carrel/attachments/2026-08",
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
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 10}, Href: "#"},
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
							Href:    "/app/notes",
							Meta:    "19",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Href: "/app/notes/nextcloud/zapisi", Meta: "3"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Href: "/app/notes/nextcloud/dnevnik", Meta: "3"},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#3F6E8C"},
							Label:  "Общие с семьёй",
							Href:   "/app/notes/nextcloud/obshchie-s-semei",
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#A05A2C"},
							Label:  "Архив",
							Href:   "/app/notes/nextcloud/arkhiv",
							Meta:   "—",
							Hidden: true,
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 of 3 notebooks · 17 notes · VJOURNAL")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Split{
									Main: w.Button{Label: "New note", Icon: "plus", Kind: w.ButtonPrimary},
									More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
								},
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
						Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "17"}}, Secondary: true},
					},
					w.ListGroup{Title: w.Rubric{Text: "August 2026", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
					w.List{
						Rows: []w.ListRow{
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Title:   []w.Inline{w.Text("Заметка со встречи 10.08")},
								Strong:  true,
								Excerpt: "Договорились заморозить ветку в среду.",
								Date:    "10 Aug",
								Source:  "nextcloud · Записи",
								On:      true,
							},
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Title:   []w.Inline{w.Text("Что почитать про CRDT")},
								Excerpt: "Кляйн, глава 5; статья Шапиро 2011; демка automerge — посмотреть, как они решают конфликт двух правок одного поля.",
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
					w.ListGroup{Title: w.Rubric{Text: "July 2026", Form: w.RubricInline}, Count: &w.Num{Value: 14}},
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
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 10}, Href: "#"},
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
							Href:    "/app/notes",
							Meta:    "19",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Href: "/app/notes/nextcloud/zapisi", Meta: "3"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Href: "/app/notes/nextcloud/dnevnik", Meta: "3"},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#3F6E8C"},
							Label:  "Общие с семьёй",
							Href:   "/app/notes/nextcloud/obshchie-s-semei",
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#A05A2C"},
							Label:  "Архив",
							Href:   "/app/notes/nextcloud/arkhiv",
							Meta:   "—",
							Hidden: true,
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 of 3 notebooks · 17 notes · VJOURNAL")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Split{
									Main: w.Button{Label: "New note", Icon: "plus", Kind: w.ButtonPrimary},
									More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
								},
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
								Excerpt: "Договорились заморозить ветку в среду.",
								Date:    "10 Aug",
								Source:  "nextcloud · Записи",
								On:      true,
							},
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Title:   []w.Inline{w.Text("Что почитать про CRDT")},
								Excerpt: "Кляйн, глава 5; статья Шапиро 2011; демка automerge — посмотреть, как они решают конфликт двух правок одного поля.",
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
					w.ListGroup{Title: w.Rubric{Text: "July 2026", Form: w.RubricInline}, Count: &w.Num{Value: 14}},
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
			RailOpen: true,
		},
	}
}

// Phone10 is frame «Разделы настроек на телефоне · Sections открыт» (phone-10).
func Phone10() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}, Sections: true},
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
								Parts: []w.HeadSubPart{w.Text("Signed in as "), w.Strong("mix"), w.Text(" · 4 servers · 16 collections · 2 subscriptions")},
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
												Gap: w.LineGapNormal,
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
											{Label: "Changed on the server", Width: w.ColChanged, Main: true},
											{Width: w.ColMenu},
										},
										Rows: []w.TableRow{
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|lichnyj"}},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Личный"}}},
													{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
													{Parts: []w.CellPart{w.Text("5")}, Mono: true},
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
													{Parts: []w.CellPart{w.Text("6")}, Mono: true},
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
													{Parts: []w.CellPart{w.Text("16")}, Mono: true},
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
													{Parts: []w.CellPart{w.Text("3")}, Mono: true},
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
													{Parts: []w.CellPart{w.Text("2")}, Mono: true},
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
													{Parts: []w.CellPart{w.Text("6")}, Mono: true},
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
													{Parts: []w.CellPart{w.Text("2")}, Mono: true},
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
												Gap: w.LineGapNormal,
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
											{Label: "Changed on the server", Width: w.ColChanged, Main: true},
											{Width: w.ColMenu},
										},
										Rows: []w.TableRow{
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "radicale|proekty"}},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Проекты"}}},
													{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
													{Parts: []w.CellPart{w.Text("2")}, Mono: true},
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
													{Parts: []w.CellPart{w.Text("2")}, Mono: true},
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
												Gap: w.LineGapNormal,
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
												Gap: w.LineGapNormal,
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
											{Label: "Collection"},
											{Label: "What is in it", Rank: w.Rank3},
											{Label: "Objects", Width: w.ColNum, Rank: w.Rank4},
											{Label: "Changed on the server", Width: w.ColChanged, Main: true},
											{Width: w.ColMenu},
										},
										Rows: []w.TableRow{
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "sftpgo|files"}},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Files"}}},
													{Parts: []w.CellPart{w.Text("files")}, Mono: true},
													{Parts: []w.CellPart{w.Text("12")}, Mono: true},
													{Parts: []w.CellPart{w.Text("5 min ago")}, Mono: true},
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
												Gap: w.LineGapNormal,
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
											{Label: "Objects", Width: w.ColNum, Rank: w.Rank3, Main: true},
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
													{Parts: []w.CellPart{w.Text("2")}, Mono: true},
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
												Gap: w.LineGapNormal,
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
			RailOpen: true,
		},
	}
}
