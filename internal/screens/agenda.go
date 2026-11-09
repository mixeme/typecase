// Раздел макета: 7.2 Календарь.

package screens

import w "carrel.dev/typecase/widget"

// Agenda1 is frame «Календарь · агенда со списком и карточкой» (s-agenda-1).
func Agenda1() Screen {
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
							w.Filter{
								Query:   w.Query{Placeholder: "Filter by word", Label: "Filter by word", Name: "filter"},
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
						Meter: w.Meter{Percent: 83, Alert: true},
						Parts: []w.RunbarPart{
							w.Text("Polled 5 of 6 sources · 7 events · "),
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
										w.Dim{Parts: []w.Inline{w.Text("text/markdown · 3.1 KB")}},
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
		},
	}
}

// Agenda27 is frame «Событие · вопрос перед сохранением серии, поверх правки» (s-agenda-27).
func Agenda27() Screen {
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
		Body: w.Overlay{
			Body: w.OverlayBody{
				Rail: []w.OverlayRailPart{
					w.Nav{Items: []w.NavItem{
						{Icon: "back", Label: "Back to the card", Href: "#"}}},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "3", Current: true},
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
															Value: "nextcloud/lichnyj",
															Of:    "event-edit",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#3F6E8C"},
															Label: "nextcloud · Работа",
															Name:  "calendar",
															Value: "nextcloud/rabota",
															Of:    "event-edit",
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
									w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "event-edit"},
								},
							},
						},
						w.Runbar{
							Meter:    w.Meter{Percent: 0},
							Progress: "event-edit",
							Parts:    []w.RunbarPart{w.Text("Saving · "), w.Mono{Text: "0 of 0 MB"}},
						},
						w.Form{
							Items: []w.FormPart{
								w.Field{
									Label: "Summary",
									Parts: []w.FieldPart{w.Input{Name: "summary", Value: "Синк-ап по релизу", Label: "Summary", Focus: true}},
									Width: w.FieldXL,
								},
								w.Field{
									Label: "Location",
									Parts: []w.FieldPart{w.Input{Name: "location", Value: "Переговорная 2", Label: "Location"}},
									Width: w.FieldLG,
								},
								w.Field{
									Label: "Status",
									Parts: []w.FieldPart{
										w.Pick{
											Name:    "status",
											Options: []w.PickOption{{Value: "confirmed", Label: "Confirmed", On: true}},
											Label:   "Status",
										},
									},
									Width: w.FieldSM,
								},
								w.Field{
									Label: "Tags",
									Parts: []w.FieldPart{w.Input{Name: "tags", Value: "release, team", Label: "Tags"}},
									Width: w.FieldMD,
								},
								w.Field{
									Label: "Description",
									Parts: []w.FieldPart{
										w.Area{
											Text:  "Обсудить оставшиеся блокеры, дату заморозки и кто пишет письмо пользователям.",
											Name:  "description",
											Label: "Text",
										},
									},
									Width: w.FieldFull,
								},
								w.FieldSet{
									Rubric: w.Rubric{Text: "When"},
									Items: []w.FieldSetPart{
										w.Field{Parts: []w.FieldPart{w.Check{Name: "all.day", Value: "1", Parts: []w.CheckPart{w.Text("All day")}}}},
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
															w.Input{Kind: w.InputSearch, Name: "time.zone", Placeholder: "Find a zone", Label: "Find a zone"},
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
										w.Field{
											Label: "First",
											Parts: []w.FieldPart{
												w.Pick{
													Name:    "first",
													Options: []w.PickOption{{Value: "10.minutes.before", Label: "10 minutes before", On: true}},
													Label:   "First",
												},
											},
											Width: w.FieldMD,
										},
										w.Field{
											Label: "Reminder",
											Parts: []w.FieldPart{
												w.Pick{
													Name:    "reminder",
													Options: []w.PickOption{{Value: "custom", Label: "Custom…", On: true}},
													Label:   "Reminder",
												},
											},
											Width: w.FieldMD,
										},
										w.Field{
											Label: "How long",
											Parts: []w.FieldPart{w.Input{Name: "reminder.count", Value: "2", Label: "How long", Mono: true, Numeric: true}},
											Width: w.FieldXS,
										},
										w.Field{
											Label: "Unit",
											Parts: []w.FieldPart{
												w.Pick{
													Name:    "reminder.unit",
													Options: []w.PickOption{{Value: "days", Label: "days before", On: true}},
													Label:   "Unit",
												},
											},
											Width: w.FieldSM,
										},
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
													Parts: []w.Inline{w.Text("A name or an address. On save, the server sends the invitations, not Carrel.")},
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
											Label: "Link",
											Parts: []w.FieldPart{w.Input{Name: "link", Placeholder: "optional", Label: "Link"}},
											Width: w.FieldMD,
										},
										w.Field{
											Label: "Related to",
											Parts: []w.FieldPart{
												w.TagInput{
													Parts: []w.TagInputPart{
														w.Tag{
															Label:     "Заметка со встречи 10.08",
															Mark:      &w.Mark{Colour: "#4A6B52", Size: w.MarkH11},
															Removable: true,
														},
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
																File: &w.FileField{Form: "", Name: "file", Label: "From this computer", List: "event-attachments"},
																Row: w.Line{
																	Items: []w.LinePart{
																		w.Faced{
																			Face:  w.FaceLarge,
																			Items: []w.LinePart{w.Text("{name} "), w.Dim{Parts: []w.Inline{w.Text("{size}")}}},
																		},
																		w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Removes: true},
																	},
																	Gap: w.LineGapNormal,
																},
																Label: "From this computer",
															},
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
														w.Dim{Parts: []w.Inline{w.Text("text/markdown · 3.1 KB")}},
														w.Text(" "),
														w.Badge{Label: "in the event"},
													},
												},
												w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true},
											},
											Gap:   w.LineGapNormal,
											Block: true,
											ID:    "event-attachments",
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
												w.Pick{
													Name:    "frequency",
													Options: []w.PickOption{{Value: "weekly", Label: "Weekly", On: true}},
													Label:   "Frequency",
												},
											},
											Width: w.FieldSM,
										},
										w.Field{
											Label: "Every",
											Parts: []w.FieldPart{w.Input{Name: "every", Value: "1", Label: "Every", Mono: true}},
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
											Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "until", Value: "2027-01-18", Label: "Until", Mono: true}},
											Width: w.FieldSM,
										},
										w.Field{
											Label: "Count",
											Parts: []w.FieldPart{w.Input{Name: "count", Placeholder: "—", Label: "Count"}},
											Width: w.FieldXS,
										},
										w.Field{
											Label: "Advanced RRULE",
											Parts: []w.FieldPart{
												w.Input{
													Name:        "advanced.rrule",
													Placeholder: "FREQ=WEEKLY;BYDAY=MO;UNTIL=20270118",
													Label:       "Advanced RRULE",
													Mono:        true,
												},
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
								w.FieldSet{
									Rubric: w.Rubric{Text: "Source"},
									Items: []w.FieldSetPart{
										w.Fields{
											Items: []w.FieldsItem{
												{
													Term:  "Path",
													Value: []w.ValuePart{w.Text("/remote.php/dav/calendars/mix/personal/8f1c.ics")},
													Face:  w.ValueMono,
												},
												{Term: "Read", Value: []w.ValuePart{w.Text("2 min ago · not changed on server")}},
											},
											Place: w.Grown,
										},
									},
								},
							},
							ID:     "event-edit",
							Action: "/app/calendar/nextcloud/lichnyj/8f1c/edit",
						},
						w.Danger{
							Button: w.Button{Label: "Delete event", Icon: "trash", Kind: w.ButtonDangerSolid},
							Note:   "Series is edited and deleted whole.",
						},
					},
				},
				Height: w.ShellHeight560,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/calendar/nextcloud/lichnyj/8f1c/edit",
					Icon:   "refresh",
					Title:  "Keep the meetings that were moved by hand?",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("Синк-ап по релизу · not saved yet")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
					},
					Body: []w.SheetPart{
						w.SheetBlock{
							Title: w.Rubric{Text: "What you changed", Variant: w.RubricBottom7},
							Items: []w.SheetLine{
								{
									Name: "The series ends on 18 January",
									Note: w.Dim{Parts: []w.Inline{w.Text("— it ran weekly until 28 December")}},
								},
							},
						},
						w.SheetBlock{
							Title: w.Rubric{Text: "Left without their day · 3", Variant: w.RubricBottom7},
							Items: []w.SheetLine{
								{
									Name: "25 January, 14:00 — Синк-ап по релизу",
									Note: w.Dim{Parts: []w.Inline{w.Text("— it was the meeting of 25 January, 12:00")}},
								},
								{
									Name: "1 February, 12:00 — Синк-ап: ретро",
									Note: w.Dim{Parts: []w.Inline{w.Text("— its title was changed")}},
								},
								{
									Name: "15 February, 16:30 — Синк-ап по релизу",
									Note: w.Dim{Parts: []w.Inline{w.Text("— it was the meeting of 15 February, 12:00")}},
								},
							},
						},
						w.SheetNote{
							Text: "Kept, each becomes an event of its own, with what was changed in it. Removed, they go with the rest of the series after 18 January.",
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Keep as events", Kind: w.ButtonPrimary, Name: "left", Value: "keep", Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Button{Label: "Remove", Name: "left", Value: "remove", Submit: true},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Name: "left", Value: "cancel", Submit: true}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGapNormal,
							Place: w.Pushed,
						},
					},
					Width: w.SheetWidth520,
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Agenda28 is frame «Событие · серия записана не вся, поверх карточки» (s-agenda-28).
func Agenda28() Screen {
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
		Body: w.Overlay{
			Body: w.OverlayBody{
				Rail: []w.OverlayRailPart{
					w.Nav{Items: []w.NavItem{
						{Icon: "back", Label: "Back to list", Short: "Back", Href: "#"}}},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "3", Current: true},
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
							Subs: []w.HeadSub{
								w.HeadSub{Parts: []w.HeadSubPart{w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, w.Text("nextcloud · Личный")}},
							},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "Edit", Icon: "pencil", Kind: w.ButtonPrimary},
									w.Button{Label: "Send", Icon: "mail"},
									w.Button{Label: "Print", Icon: "print", Tertiary: true},
									w.Button{Icon: "dots", Title: "More"},
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
								{
									Term:  "Repeats",
									Value: []w.ValuePart{w.Text("Weekly on Mon "), w.Badge{Label: "until 18 January 2027"}},
								},
								{Term: "Reminders", Value: []w.ValuePart{w.Text("10 minutes before · 2 days before")}},
								{
									Term:  "Where",
									Value: []w.ValuePart{w.Text("Переговорная 2 "), w.Badge{Label: "linked", Tone: w.BadgeLinked}},
								},
								{Term: "Status", Value: []w.ValuePart{w.Text("Confirmed")}},
								{Term: "Tags", Value: []w.ValuePart{w.Tag{Label: "релиз"}, w.Text(" "), w.Tag{Label: "работа"}}},
							},
						},
						w.Sec{
							Items: []w.SecPart{
								w.Rubric{Text: "Description"},
								w.Prose{
									Parts: []w.Inline{w.Text("Обсудить оставшиеся блокеры, дату заморозки и кто пишет письмо пользователям.")},
								},
							},
						},
						w.Sec{
							Items: []w.SecPart{
								w.Rubric{Text: "Attendees", Items: []w.RubricPart{w.Mono{Text: "4", Variant: w.MonoTight}}},
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
											w.Dim{Parts: []w.Inline{w.Text("text/markdown · 3.1 KB")}},
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
								w.Rubric{
									Text:  "The same event in other calendars",
									Items: []w.RubricPart{w.Badge{Label: "2", Tone: w.BadgeLinked}},
								},
								w.SecList{
									Items: [][]w.SecItemPart{
										[]w.SecItemPart{
											w.Mark{Colour: "#6B4E7A", Size: w.MarkH11Inline},
											w.Text(" "),
											w.Link{Text: "radicale · vps · Проекты", Href: "#"},
											w.Text(" "),
											w.Dim{Parts: []w.Inline{w.Text("· Holds the same")}},
										},
										[]w.SecItemPart{
											w.Mark{Colour: "#9A9280", Size: w.MarkH11Inline},
											w.Text(" "),
											w.Link{Text: "baikal · Общий", Href: "#"},
											w.Text(" "),
											w.Dim{Parts: []w.Inline{w.Text("· Where differ")}},
										},
									},
									Gap: w.StackGapTight,
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
							Button: w.Button{Label: "Delete event", Icon: "trash", Kind: w.ButtonDangerSolid},
							Note:   "Series is edited and deleted whole.",
						},
					},
				},
				Height: w.ShellHeight560,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/calendar/nextcloud/lichnyj/8f1c/edit",
					Icon:   "refresh",
					Title:  "Keep the meetings that were moved by hand?",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("Синк-ап по релизу · saved in part")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
					},
					Body: []w.SheetPart{
						w.SheetBlock{
							Title: w.Rubric{Text: "Written", Variant: w.RubricBottom7},
							Items: []w.SheetLine{
								{Name: "The series ends on 18 January", Note: w.Dim{Parts: []w.Inline{w.Text("— saved")}}},
								{
									Name: "2 separate events",
									Note: w.Dim{Parts: []w.Inline{w.Text("— 25 January, 14:00 and 15 February, 16:30")}},
								},
							},
						},
						w.SheetBlock{
							Title: w.Rubric{Text: "Not written · 1", Variant: w.RubricBottom7},
							Items: []w.SheetLine{
								{
									Name: "1 February, 12:00 — Синк-ап: ретро",
									Note: w.Dim{Parts: []w.Inline{w.Text("— the server answered 507 Insufficient Storage")}},
								},
							},
						},
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("The meeting of 1 February was not made a separate event: it stays in the series as it was before you saved."),
							},
							Tone:  w.MsgAlert,
							Space: w.MsgFlush,
						},
						w.SheetNote{Text: "Nothing is lost: what could not be written apart is still in the series."},
					},
					Foot: []w.SheetFootPart{
						w.Button{
							Label:  "Try again",
							Kind:   w.ButtonPrimary,
							To:     "/app/calendar/nextcloud/lichnyj/8f1c/edit?op=apart&day=2027-02-01",
							Submit: true,
						},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{
								w.Button{Label: "Leave it in the series", Href: "/app/calendar/nextcloud/lichnyj/8f1c"},
								w.Kbd{Keys: "Esc"},
							},
							Gap:   w.LineGapNormal,
							Place: w.Pushed,
						},
					},
					Width: w.SheetWidth520,
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Agenda17 is frame «Календарь · один календарь» (s-agenda-17).
func Agenda17() Screen {
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
					On: 2,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked: w.TickedPart,
							Kind:   w.TickRowAll,
							Mark:   w.Mark{Kind: w.MarkMulti},
							Label:  "Combined calendar",
							Meta:   "7",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "3", Current: true},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH20}, Text: "Личный"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{w.Text("nextcloud · 10 – 16 August 2026 · "), w.Mono{Text: "Europe/Moscow"}, w.Text(" · 3 events")},
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
							w.Filter{
								Query:   w.Query{Placeholder: "Filter by word", Label: "Filter by word", Name: "filter"},
								Waiting: true,
								Action:  "/app/calendar/nextcloud/lichnyj",
								Swap:    "list-main",
								Variant: w.FilterH26W200,
							},
							w.Button{Icon: "back", Title: "Previous week", Href: "/app/calendar/nextcloud/lichnyj?from=2026-08-03"},
							w.Button{Label: "Today", Href: "/app/calendar/nextcloud/lichnyj"},
							w.Button{Icon: "next", Title: "Next week", Href: "/app/calendar/nextcloud/lichnyj?from=2026-08-17"},
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
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "10:00",
								Title:  []w.Inline{w.Text("Синк-ап по релизу")},
								Detail: "Переговорная 2",
								Badges: []w.Badge{w.Badge{Label: "4 attendees"}, w.Badge{Label: "repeats weekly"}},
								Single: true,
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
								Single: true,
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
								Single: true,
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
								Single: true,
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
								Single: true,
							},
						},
					},
				},
				ID: "list-main",
			},
		},
	}
}

// Agenda31 is frame «Календарь · фильтр по слову» (s-agenda-31).
func Agenda31() Screen {
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
					On: 2,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked: w.TickedPart,
							Kind:   w.TickRowAll,
							Mark:   w.Mark{Kind: w.MarkMulti},
							Label:  "Combined calendar",
							Meta:   "7",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "3", Current: true},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH20}, Text: "Личный"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{
									w.Text("nextcloud · 10 – 16 August 2026 · "),
									w.Mono{Text: "Europe/Moscow"},
									w.Text(" · 1 of 3 events"),
								},
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
							w.Filter{
								Query:   w.Query{Value: "синк", Label: "Filter by word", Name: "filter"},
								Clear:   "/app/calendar/nextcloud/lichnyj",
								Action:  "/app/calendar/nextcloud/lichnyj",
								Swap:    "list-main",
								Variant: w.FilterH26W200,
							},
							w.Button{
								Icon:  "back",
								Title: "Previous week",
								Href:  "/app/calendar/nextcloud/lichnyj?from=2026-08-03&filter=синк",
							},
							w.Button{Label: "Today", Href: "/app/calendar/nextcloud/lichnyj?filter=синк"},
							w.Button{
								Icon:  "next",
								Title: "Next week",
								Href:  "/app/calendar/nextcloud/lichnyj?from=2026-08-17&filter=синк",
							},
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
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "10:00",
								Title:  []w.Inline{w.Text("Синк-ап по релизу")},
								Detail: "Переговорная 2",
								Badges: []w.Badge{w.Badge{Label: "4 attendees"}, w.Badge{Label: "repeats weekly"}},
								Single: true,
							},
						},
					},
				},
				ID: "list-main",
			},
		},
	}
}

// Agenda2 is frame «Событие · чтение» (s-agenda-2).
func Agenda2() Screen {
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
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "3", Current: true},
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
						Subs: []w.HeadSub{
							w.HeadSub{Parts: []w.HeadSubPart{w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, w.Text("nextcloud · Личный")}},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Edit", Icon: "pencil", Kind: w.ButtonPrimary},
								w.Button{Label: "Send", Icon: "mail"},
								w.Button{Label: "Print", Icon: "print", Tertiary: true},
								w.Button{Icon: "dots", Title: "More"},
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
							{
								Term:  "Where",
								Value: []w.ValuePart{w.Text("Переговорная 2 "), w.Badge{Label: "linked", Tone: w.BadgeLinked}},
							},
							{Term: "Status", Value: []w.ValuePart{w.Text("Confirmed")}},
							{Term: "Tags", Value: []w.ValuePart{w.Tag{Label: "релиз"}, w.Text(" "), w.Tag{Label: "работа"}}},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Description"},
							w.Prose{
								Parts: []w.Inline{w.Text("Обсудить оставшиеся блокеры, дату заморозки и кто пишет письмо пользователям.")},
							},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Attendees", Items: []w.RubricPart{w.Mono{Text: "4", Variant: w.MonoTight}}},
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
										w.Dim{Parts: []w.Inline{w.Text("text/markdown · 3.1 KB")}},
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
							w.Rubric{
								Text:  "The same event in other calendars",
								Items: []w.RubricPart{w.Badge{Label: "2", Tone: w.BadgeLinked}},
							},
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{
										w.Mark{Colour: "#6B4E7A", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "radicale · vps · Проекты", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· Holds the same")}},
									},
									[]w.SecItemPart{
										w.Mark{Colour: "#9A9280", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "baikal · Общий", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· When, where differ")}},
									},
								},
								Gap: w.StackGapTight,
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
						Button: w.Button{Label: "Delete event", Icon: "trash", Kind: w.ButtonDangerSolid},
						Note:   "Series is edited and deleted whole.",
					},
				},
			},
		},
	}
}

// Agenda3 is frame «Событие · правка, повторение» (s-agenda-3).
func Agenda3() Screen {
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "3", Current: true},
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
														Value: "nextcloud/lichnyj",
														Of:    "event-edit",
													},
													w.MenuItem{
														Mark:  &w.Mark{Colour: "#3F6E8C"},
														Label: "nextcloud · Работа",
														Name:  "calendar",
														Value: "nextcloud/rabota",
														Of:    "event-edit",
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
								w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "event-edit"},
							},
						},
					},
					w.Runbar{
						Meter:    w.Meter{Percent: 0},
						Progress: "event-edit",
						Parts:    []w.RunbarPart{w.Text("Saving · "), w.Mono{Text: "0 of 0 MB"}},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Summary",
								Parts: []w.FieldPart{w.Input{Name: "summary", Value: "Синк-ап по релизу", Label: "Summary", Focus: true}},
								Width: w.FieldXL,
							},
							w.Field{
								Label: "Location",
								Parts: []w.FieldPart{w.Input{Name: "location", Value: "Переговорная 2", Label: "Location"}},
								Width: w.FieldLG,
							},
							w.Field{
								Label: "Status",
								Parts: []w.FieldPart{
									w.Pick{
										Name:    "status",
										Options: []w.PickOption{{Value: "confirmed", Label: "Confirmed", On: true}},
										Label:   "Status",
									},
								},
								Width: w.FieldSM,
							},
							w.Field{
								Label: "Tags",
								Parts: []w.FieldPart{w.Input{Name: "tags", Value: "release, team", Label: "Tags"}},
								Width: w.FieldMD,
							},
							w.Field{
								Label: "Description",
								Parts: []w.FieldPart{
									w.Area{
										Text:  "Обсудить оставшиеся блокеры, дату заморозки и кто пишет письмо пользователям.",
										Name:  "description",
										Label: "Text",
									},
								},
								Width: w.FieldFull,
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "When"},
								Items: []w.FieldSetPart{
									w.Field{Parts: []w.FieldPart{w.Check{Name: "all.day", Value: "1", Parts: []w.CheckPart{w.Text("All day")}}}},
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
														w.Input{Kind: w.InputSearch, Name: "time.zone", Placeholder: "Find a zone", Label: "Find a zone"},
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
									w.Field{
										Label: "First",
										Parts: []w.FieldPart{
											w.Pick{
												Name:    "first",
												Options: []w.PickOption{{Value: "10.minutes.before", Label: "10 minutes before", On: true}},
												Label:   "First",
											},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Reminder",
										Parts: []w.FieldPart{
											w.Pick{
												Name:    "reminder",
												Options: []w.PickOption{{Value: "custom", Label: "Custom…", On: true}},
												Label:   "Reminder",
											},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "How long",
										Parts: []w.FieldPart{w.Input{Name: "reminder.count", Value: "2", Label: "How long", Mono: true, Numeric: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Unit",
										Parts: []w.FieldPart{
											w.Pick{
												Name:    "reminder.unit",
												Options: []w.PickOption{{Value: "days", Label: "days before", On: true}},
												Label:   "Unit",
											},
										},
										Width: w.FieldSM,
									},
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
												Parts: []w.Inline{w.Text("A name or an address. On save, the server sends the invitations, not Carrel.")},
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
										Label: "Link",
										Parts: []w.FieldPart{w.Input{Name: "link", Placeholder: "optional", Label: "Link"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Related to",
										Parts: []w.FieldPart{
											w.TagInput{
												Parts: []w.TagInputPart{
													w.Tag{
														Label:     "Заметка со встречи 10.08",
														Mark:      &w.Mark{Colour: "#4A6B52", Size: w.MarkH11},
														Removable: true,
													},
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
															File: &w.FileField{Form: "", Name: "file", Label: "From this computer", List: "event-attachments"},
															Row: w.Line{
																Items: []w.LinePart{
																	w.Faced{
																		Face:  w.FaceLarge,
																		Items: []w.LinePart{w.Text("{name} "), w.Dim{Parts: []w.Inline{w.Text("{size}")}}},
																	},
																	w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Removes: true},
																},
																Gap: w.LineGapNormal,
															},
															Label: "From this computer",
														},
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
													w.Dim{Parts: []w.Inline{w.Text("text/markdown · 3.1 KB")}},
													w.Text(" "),
													w.Badge{Label: "in the event"},
												},
											},
											w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true},
										},
										Gap:   w.LineGapNormal,
										Block: true,
										ID:    "event-attachments",
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
											w.Pick{
												Name:    "frequency",
												Options: []w.PickOption{{Value: "weekly", Label: "Weekly", On: true}},
												Label:   "Frequency",
											},
										},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Every",
										Parts: []w.FieldPart{w.Input{Name: "every", Value: "1", Label: "Every", Mono: true}},
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
										Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "until", Value: "2026-12-28", Label: "Until", Mono: true}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Count",
										Parts: []w.FieldPart{w.Input{Name: "count", Placeholder: "—", Label: "Count"}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Advanced RRULE",
										Parts: []w.FieldPart{
											w.Input{
												Name:        "advanced.rrule",
												Placeholder: "FREQ=WEEKLY;BYDAY=MO;UNTIL=20261228",
												Label:       "Advanced RRULE",
												Mono:        true,
											},
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
							w.FieldSet{
								Rubric: w.Rubric{Text: "Source"},
								Items: []w.FieldSetPart{
									w.Fields{
										Items: []w.FieldsItem{
											{
												Term:  "Path",
												Value: []w.ValuePart{w.Text("/remote.php/dav/calendars/mix/personal/8f1c.ics")},
												Face:  w.ValueMono,
											},
											{Term: "Read", Value: []w.ValuePart{w.Text("2 min ago · not changed on server")}},
										},
										Place: w.Grown,
									},
								},
							},
						},
						ID:     "event-edit",
						Action: "/app/calendar/nextcloud/lichnyj/8f1c/edit",
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

// Agenda15 is frame «Событие · новое» (s-agenda-15).
func Agenda15() Screen {
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
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "3", Current: true},
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
						Title: &w.HeadTitle{Text: "New event"},
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
														Value: "nextcloud/lichnyj",
														Of:    "event-new",
													},
													w.MenuItem{
														Mark:  &w.Mark{Colour: "#3F6E8C"},
														Label: "nextcloud · Работа",
														Name:  "calendar",
														Value: "nextcloud/rabota",
														Of:    "event-new",
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
								w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "event-new"},
							},
						},
					},
					w.Runbar{
						Meter:    w.Meter{Percent: 0},
						Progress: "event-new",
						Parts:    []w.RunbarPart{w.Text("Saving · "), w.Mono{Text: "0 of 0 MB"}},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Summary",
								Parts: []w.FieldPart{w.Input{Name: "summary", Placeholder: "—", Label: "Summary", Focus: true}},
								Width: w.FieldXL,
							},
							w.Field{
								Label: "Location",
								Parts: []w.FieldPart{w.Input{Name: "location", Placeholder: "—", Label: "Location"}},
								Width: w.FieldLG,
							},
							w.Field{
								Label: "Status",
								Parts: []w.FieldPart{
									w.Pick{
										Name:    "status",
										Options: []w.PickOption{{Value: "no.status", Label: "No status", On: true}},
										Label:   "Status",
									},
								},
								Width: w.FieldSM,
							},
							w.Field{
								Label: "Tags",
								Parts: []w.FieldPart{w.Input{Name: "tags", Placeholder: "—", Label: "Tags"}},
								Width: w.FieldMD,
							},
							w.Field{
								Label: "Description",
								Parts: []w.FieldPart{w.Area{Text: "", Name: "description", Label: "Text"}},
								Width: w.FieldFull,
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "When"},
								Items: []w.FieldSetPart{
									w.Field{Parts: []w.FieldPart{w.Check{Name: "all.day", Value: "1", Parts: []w.CheckPart{w.Text("All day")}}}},
									w.Field{
										Label: "Start date",
										Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "start.date", Value: "2026-08-10", Label: "Start date", Mono: true}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Start time",
										Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Name: "start.time", Value: "10:00", Label: "Start time", Mono: true}},
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
												Value:   "11:00",
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
												Trigger: w.Pick{Value: "Europe/Moscow"},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.Input{Kind: w.InputSearch, Name: "time.zone", Placeholder: "Find a zone", Label: "Find a zone"},
														w.MenuItem{Label: "Europe/Moscow", On: true, Name: "tz", Value: "Europe/Moscow"},
														w.MenuSep{},
														w.MenuItem{Label: "Asia/Yekaterinburg", Name: "tz", Value: "Asia/Yekaterinburg"},
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
									w.Field{
										Label: "First",
										Parts: []w.FieldPart{
											w.Pick{
												Name:    "first",
												Options: []w.PickOption{{Value: "10.minutes.before", Label: "10 minutes before", On: true}},
												Label:   "First",
											},
										},
										Width: w.FieldMD,
									},
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
												Parts: []w.Inline{w.Text("A name or an address. On save, the server sends the invitations, not Carrel.")},
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
										Label: "Link",
										Parts: []w.FieldPart{w.Input{Name: "link", Placeholder: "optional", Label: "Link"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Related to",
										Parts: []w.FieldPart{
											w.TagInput{
												Parts: []w.TagInputPart{},
												Add:   &w.TagAdd{Name: "related.add", Placeholder: "Title or UID", Label: "Add a link"},
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
															File: &w.FileField{Form: "", Name: "file", Label: "From this computer", List: "event-attachments"},
															Row: w.Line{
																Items: []w.LinePart{
																	w.Faced{
																		Face:  w.FaceLarge,
																		Items: []w.LinePart{w.Text("{name} "), w.Dim{Parts: []w.Inline{w.Text("{size}")}}},
																	},
																	w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Removes: true},
																},
																Gap: w.LineGapNormal,
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
										Gap:   w.LineGapNormal,
										Block: true,
										ID:    "event-attachments",
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
											w.Pick{
												Name:    "frequency",
												Options: []w.PickOption{{Value: "does.not.repeat", Label: "Does not repeat", On: true}},
												Label:   "Frequency",
											},
										},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Every",
										Parts: []w.FieldPart{w.Input{Name: "every", Value: "1", Label: "Every", Mono: true}},
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
										Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "until", Label: "Until", Mono: true}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Count",
										Parts: []w.FieldPart{w.Input{Name: "count", Placeholder: "—", Label: "Count"}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Advanced RRULE",
										Parts: []w.FieldPart{
											w.Input{Name: "advanced.rrule", Placeholder: "—", Label: "Advanced RRULE", Mono: true},
											w.FieldHint{Parts: []w.Inline{w.Text("The advanced value takes precedence when set.")}},
										},
										Width: w.FieldXL,
									},
								},
							},
						},
						ID:     "event-new",
						Action: "/app/calendar/nextcloud/lichnyj/new",
					},
				},
			},
		},
	}
}

// Agenda32 is frame «Событие · новое, файлы больше не хранятся» (s-agenda-32).
func Agenda32() Screen {
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
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "3", Current: true},
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
						Title: &w.HeadTitle{Text: "New event"},
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
														Value: "nextcloud/lichnyj",
														Of:    "event-new",
													},
													w.MenuItem{
														Mark:  &w.Mark{Colour: "#3F6E8C"},
														Label: "nextcloud · Работа",
														Name:  "calendar",
														Value: "nextcloud/rabota",
														Of:    "event-new",
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
								w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "event-new"},
							},
						},
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Strong("Nothing was saved:"),
							w.Text(" the files attached here are no longer kept — newer forms took their place. Attach them again, or save without them."),
						},
						Tone:  w.MsgAlert,
						Space: w.MsgFlushTop2,
					},
					w.Runbar{
						Meter:    w.Meter{Percent: 0},
						Progress: "event-new",
						Parts:    []w.RunbarPart{w.Text("Saving · "), w.Mono{Text: "0 of 0 MB"}},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Summary",
								Parts: []w.FieldPart{w.Input{Name: "summary", Value: "Планёрка по обложке", Label: "Summary"}},
								Width: w.FieldXL,
							},
							w.Field{
								Label: "Location",
								Parts: []w.FieldPart{w.Input{Name: "location", Value: "Переговорная 1", Label: "Location"}},
								Width: w.FieldLG,
							},
							w.Field{
								Label: "Status",
								Parts: []w.FieldPart{
									w.Pick{
										Name:    "status",
										Options: []w.PickOption{{Value: "no.status", Label: "No status", On: true}},
										Label:   "Status",
									},
								},
								Width: w.FieldSM,
							},
							w.Field{
								Label: "Tags",
								Parts: []w.FieldPart{w.Input{Name: "tags", Placeholder: "—", Label: "Tags"}},
								Width: w.FieldMD,
							},
							w.Field{
								Label: "Description",
								Parts: []w.FieldPart{w.Area{Text: "", Name: "description", Label: "Text"}},
								Width: w.FieldFull,
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "When"},
								Items: []w.FieldSetPart{
									w.Field{Parts: []w.FieldPart{w.Check{Name: "all.day", Value: "1", Parts: []w.CheckPart{w.Text("All day")}}}},
									w.Field{
										Label: "Start date",
										Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "start.date", Value: "2026-08-13", Label: "Start date", Mono: true}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Start time",
										Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Name: "start.time", Value: "14:00", Label: "Start time", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "End date",
										Parts: []w.FieldPart{
											w.Input{
												Kind:    w.InputDate,
												Name:    "end.date",
												Value:   "2026-08-13",
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
												Value:   "15:00",
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
												Trigger: w.Pick{Value: "Europe/Moscow"},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.Input{Kind: w.InputSearch, Name: "time.zone", Placeholder: "Find a zone", Label: "Find a zone"},
														w.MenuItem{Label: "Europe/Moscow", On: true, Name: "tz", Value: "Europe/Moscow"},
														w.MenuSep{},
														w.MenuItem{Label: "Asia/Yekaterinburg", Name: "tz", Value: "Asia/Yekaterinburg"},
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
									w.Field{
										Label: "First",
										Parts: []w.FieldPart{
											w.Pick{
												Name:    "first",
												Options: []w.PickOption{{Value: "10.minutes.before", Label: "10 minutes before", On: true}},
												Label:   "First",
											},
										},
										Width: w.FieldMD,
									},
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
												Parts: []w.Inline{w.Text("A name or an address. On save, the server sends the invitations, not Carrel.")},
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
										Label: "Link",
										Parts: []w.FieldPart{w.Input{Name: "link", Placeholder: "optional", Label: "Link"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Related to",
										Parts: []w.FieldPart{
											w.TagInput{
												Parts: []w.TagInputPart{},
												Add:   &w.TagAdd{Name: "related.add", Placeholder: "Title or UID", Label: "Add a link"},
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
															File: &w.FileField{Form: "", Name: "file", Label: "From this computer", List: "event-attachments"},
															Row: w.Line{
																Items: []w.LinePart{
																	w.Faced{
																		Face:  w.FaceLarge,
																		Items: []w.LinePart{w.Text("{name} "), w.Dim{Parts: []w.Inline{w.Text("{size}")}}},
																	},
																	w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Removes: true},
																},
																Gap: w.LineGapNormal,
															},
															Label: "From this computer",
														},
														w.MenuItem{Icon: "file", Label: "From Files"},
													},
													Width: w.MenuWidth212,
													Pop:   w.MenuUnder,
												},
											},
											w.Faced{
												Face: w.FaceLarge,
												Items: []w.LinePart{
													w.Text("макет-обложки-v3.pdf "),
													w.Dim{Parts: []w.Inline{w.Text("application/pdf · 1.8 MB")}},
													w.Text(" "),
													w.Badge{Label: "not kept", Tone: w.BadgeAlert},
												},
											},
											w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true},
											w.Faced{
												Face: w.FaceLarge,
												Items: []w.LinePart{
													w.Text("шрифты-обложки.png "),
													w.Dim{Parts: []w.Inline{w.Text("image/png · 640 KB")}},
													w.Text(" "),
													w.Badge{Label: "not kept", Tone: w.BadgeAlert},
												},
											},
											w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true},
										},
										Gap:   w.LineGapNormal,
										Block: true,
										ID:    "event-attachments",
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
											w.Pick{
												Name:    "frequency",
												Options: []w.PickOption{{Value: "does.not.repeat", Label: "Does not repeat", On: true}},
												Label:   "Frequency",
											},
										},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Every",
										Parts: []w.FieldPart{w.Input{Name: "every", Value: "1", Label: "Every", Mono: true}},
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
										Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "until", Label: "Until", Mono: true}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Count",
										Parts: []w.FieldPart{w.Input{Name: "count", Placeholder: "—", Label: "Count"}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Advanced RRULE",
										Parts: []w.FieldPart{
											w.Input{Name: "advanced.rrule", Placeholder: "—", Label: "Advanced RRULE", Mono: true},
											w.FieldHint{Parts: []w.Inline{w.Text("The advanced value takes precedence when set.")}},
										},
										Width: w.FieldXL,
									},
								},
							},
						},
						ID:     "event-new",
						Action: "/app/calendar/nextcloud/lichnyj/new",
					},
				},
			},
		},
	}
}
