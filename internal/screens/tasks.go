// Раздел макета: 7.3 Задачи.

package screens

import w "carrel.dev/typecase/widget"

// Tasks1 is frame «Задачи · фильтры, просрочка» (s-tasks-1).
func Tasks1() Screen {
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
							Meta:    "7",
							Current: true,
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 of 3 lists · nextcloud · 6 tasks")}}},
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
							w.Seg{Options: []w.SegOption{{Label: "Open 5", On: true}, {Label: "Done 1"}, {Label: "All 6"}}},
							w.Seg{
								Options:   []w.SegOption{{Label: "By due", On: true}, {Label: "By priority"}, {Label: "Recently changed"}},
								Secondary: true,
							},
							w.Button{Label: "Columns", Icon: "cols"},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Overdue", Form: w.RubricInline}, Count: &w.Num{Value: 2, Alert: true}},
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
							w.DueRow{
								Tick:     w.Tick{Label: "Done"},
								Mark:     w.Mark{Colour: "#3F6E8C"},
								Title:    []w.Inline{w.Text("Ответить подрядчику по смете")},
								Due:      "8 Aug",
								Progress: "40%",
								Source:   "nextcloud · Работа",
								Overdue:  true,
								Href:     "#",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "This week", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.DueRow{
								Tick:       w.Tick{Label: "Done"},
								Mark:       w.Mark{Colour: "#4A6B52"},
								Title:      []w.Inline{w.Text("Собрать чемодан")},
								Strong:     true,
								TitleBadge: &w.Badge{Label: "linked ×2", Tone: w.BadgeLinked},
								Due:        "12 Aug 18:00",
								Progress:   "—",
								Source:     "2 sources",
								SourceHint: "nextcloud · Дела, nextcloud · Работа",
								On:         true,
								Href:       "#",
							},
							w.DueRow{
								Tick:     w.Tick{Label: "Done"},
								Mark:     w.Mark{Colour: "#3F6E8C"},
								Title:    []w.Inline{w.Text("Дописать раздел про кэш")},
								Due:      "14 Aug",
								Progress: "70%",
								Source:   "nextcloud · Работа",
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
			Side: &w.Side{
				Items: []w.SidePart{
					w.Head{
						Title: &w.HeadTitle{Text: "Собрать чемодан", Size: w.HeadTitle19},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("in 2 lists")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Edit", Small: true, Href: "/app/tasks/linked/c41a/copies"},
								w.Button{Icon: "dots", Small: true, Title: "More"},
								w.Close{Label: "Close the panel"},
							},
						},
					},
					w.Fields{
						Items: []w.FieldsItem{
							{Term: "Status", Value: []w.ValuePart{w.Text("Needs action")}},
							{
								Term: "Due",
								Value: []w.ValuePart{
									w.Text("12 August 2026, 18:00 "),
									w.Dim{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11Inline}, Parts: []w.Inline{w.Text("Дела")}},
								},
								Face: w.ValueMono,
							},
							{Term: "Complete", Value: []w.ValuePart{w.Text("0%")}, Face: w.ValueMono},
							{
								Term: "Priority",
								Value: []w.ValuePart{
									w.Text("3 "),
									w.Dim{Mark: &w.Mark{Colour: "#3F6E8C", Size: w.MarkH11Inline}, Parts: []w.Inline{w.Text("Работа")}},
								},
								Face: w.ValueMono,
							},
							{Term: "Reminders", Value: []w.ValuePart{w.Text("1 day before")}},
							{Term: "Tags", Value: []w.ValuePart{w.Tag{Label: "дом"}}},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Description"},
							w.Prose{Parts: []w.Inline{w.Text("Зарядки, переходник, распечатать посадочный.")}, Variant: w.Prose13},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Attachments"},
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{
										w.Link{Text: "посадочный-казань.pdf", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("application/pdf · 128 KB")}},
										w.Text(" "),
										w.Badge{Label: "in the task"},
										w.Text(" "),
										w.Dim{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11Inline}, Parts: []w.Inline{w.Text("Дела")}},
									},
								},
								Gap: 6,
							},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Related"},
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{
										w.Span{Face: w.FaceAccent, Text: "Поездка в Казань"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· event")}},
										w.Text(" "),
										w.Dim{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11Inline}, Parts: []w.Inline{w.Text("Дела")}},
									},
								},
								Gap: 6,
							},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Where this task is kept", Items: []w.RubricPart{w.Badge{Label: "2", Tone: w.BadgeLinked}}},
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{
										w.Mark{Colour: "#4A6B52", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "nextcloud · Дела", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· Priority differs")}},
									},
									[]w.SecItemPart{
										w.Mark{Colour: "#3F6E8C", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "nextcloud · Работа", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· Due differs")}},
									},
								},
								Gap: 5,
							},
						},
					},
				},
			},
		},
	}
}

// Tasks15 is frame «Задачи · один список» (s-tasks-15).
func Tasks15() Screen {
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "4", Current: true},
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · 4 tasks")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New task", Icon: "plus", Kind: w.ButtonPrimary},
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
					w.ListGroup{Title: w.Rubric{Text: "Overdue", Form: w.RubricInline}, Count: &w.Num{Value: 1, Alert: true}},
					w.List{
						Rows: []w.ListRow{
							w.DueRow{
								Tick:     w.Tick{Label: "Done"},
								Mark:     w.Mark{Colour: "#4A6B52"},
								Title:    []w.Inline{w.Text("Отправить документы в МФЦ")},
								Due:      "7 Aug",
								Progress: "—",
								Single:   true,
								Overdue:  true,
								Href:     "#",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "This week", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.DueRow{
								Tick:       w.Tick{Label: "Done"},
								Mark:       w.Mark{Colour: "#4A6B52"},
								Title:      []w.Inline{w.Text("Собрать чемодан")},
								TitleBadge: &w.Badge{Label: "linked ×2", Tone: w.BadgeLinked, Title: "nextcloud · Дела, nextcloud · Работа"},
								Due:        "12 Aug 18:00",
								Progress:   "—",
								Single:     true,
								Href:       "#",
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

// Tasks2 is frame «Задача · чтение» (s-tasks-2).
func Tasks2() Screen {
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "4", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "3"},
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
						Title: &w.HeadTitle{Text: "Собрать чемодан"},
						Subs: []w.HeadSub{
							w.HeadSub{Parts: []w.HeadSubPart{w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, w.Text("nextcloud · Дела")}},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Edit", Icon: "pencil", Kind: w.ButtonPrimary},
								w.Button{Label: "Mark done", Tick: &w.Tick{Glyph: true}},
								w.Button{Label: "Send", Icon: "mail"},
								w.Button{Label: "Print", Icon: "print", Tertiary: true},
								w.Button{Icon: "dots", Title: "More"},
							},
						},
					},
					w.Fields{
						Items: []w.FieldsItem{
							{Term: "Status", Value: []w.ValuePart{w.Text("Needs action")}},
							{
								Term:  "Due",
								Value: []w.ValuePart{w.Text("12 August 2026, 18:00 "), w.Badge{Label: "linked", Tone: w.BadgeLinked}},
								Face:  w.ValueMono,
							},
							{Term: "Complete", Value: []w.ValuePart{w.Text("0%")}, Face: w.ValueMono},
							{
								Term:  "Priority",
								Value: []w.ValuePart{w.Text("5 "), w.Badge{Label: "linked", Tone: w.BadgeLinked}},
								Face:  w.ValueMono,
							},
							{Term: "Reminders", Value: []w.ValuePart{w.Text("1 day before")}},
							{Term: "Tags", Value: []w.ValuePart{w.Tag{Label: "дом"}}},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Description"},
							w.Prose{Parts: []w.Inline{w.Text("Зарядки, переходник, распечатать посадочный.")}},
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
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{
										w.Link{Text: "посадочный-казань.pdf", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("application/pdf · 128 KB")}},
										w.Text(" "),
										w.Badge{Label: "in the task"},
									},
								},
								Gap: 6,
							},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Related"},
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{
										w.Span{Face: w.FaceAccent, Text: "Поездка в Казань"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· event · 12 Aug")}},
									},
								},
							},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{
								Text:  "The same task in other lists",
								Items: []w.RubricPart{w.Badge{Label: "1", Tone: w.BadgeLinked}},
							},
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{
										w.Mark{Colour: "#6B4E7A", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "nextcloud · Работа", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· Due, priority differ")}},
									},
								},
								Gap: 5,
							},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Source"},
							w.Diag{Text: "Path  /remote.php/dav/calendars/mix/dela/c41a.ics\nRead  2 min ago · not changed on server"},
						},
					},
					w.Danger{Button: w.Button{Label: "Delete task", Icon: "trash", Kind: w.ButtonDangerSolid}},
				},
			},
		},
	}
}

// Tasks3 is frame «Задача · правка, срок, выполнение» (s-tasks-3).
func Tasks3() Screen {
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
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "3"},
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
														Of:    "task-edit",
													},
													w.MenuItem{
														Mark:  &w.Mark{Colour: "#3F6E8C"},
														Label: "nextcloud · Работа",
														Name:  "list",
														Value: "nextcloud/rabota",
														Of:    "task-edit",
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
						Progress: "task-edit",
						Parts:    []w.RunbarPart{w.Text("Saving · "), w.Mono{Text: "0 of 0 MB"}},
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
															File: &w.FileField{Form: "", Name: "file", Label: "From this computer", List: "task-attachments"},
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
										ID:    "task-attachments",
										Place: w.Grown,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Repeats"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Frequency", Parts: []w.FieldPart{w.Pick{Value: "Does not repeat"}}, Width: w.FieldSM},
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
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("The advanced value takes precedence when set. A task repeats from its start date, or from the due date if it has none."),
												},
											},
										},
										Width: w.FieldXL,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Kept as the server sent them"},
								Items:  []w.FieldSetPart{w.Diag{Text: "X-APPLE-SORT-ORDER  712\nX-OC-HIDESUBTASKS  0\nSEQUENCE  1", Place: w.Grown}},
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

// Tasks14 is frame «Задача · новая» (s-tasks-14).
func Tasks14() Screen {
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "4", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "3"},
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
						Title: &w.HeadTitle{Text: "New task"},
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
														Of:    "task-new",
													},
													w.MenuItem{
														Mark:  &w.Mark{Colour: "#3F6E8C"},
														Label: "nextcloud · Работа",
														Name:  "list",
														Value: "nextcloud/rabota",
														Of:    "task-new",
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
						Progress: "task-new",
						Parts:    []w.RunbarPart{w.Text("Saving · "), w.Mono{Text: "0 of 0 MB"}},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Summary",
								Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Summary", Focus: true}},
								Width: w.FieldXL,
							},
							w.Field{Label: "Description", Parts: []w.FieldPart{w.Area{Text: "", Label: "Text"}}, Width: w.FieldFull},
							w.FieldSet{
								Rubric: w.Rubric{Text: "When"},
								Items: []w.FieldSetPart{
									w.Field{Parts: []w.FieldPart{w.Check{Parts: []w.CheckPart{w.Text("All day")}}}},
									w.Field{
										Label: "Start date",
										Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Placeholder: "YYYY-MM-DD", Label: "Start date", Mono: true}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Start time",
										Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Placeholder: "--:--", Label: "Start time", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Due date",
										Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Placeholder: "YYYY-MM-DD", Label: "Due date", Mono: true}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Due time",
										Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Placeholder: "--:--", Label: "Due time", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Time zone",
										Parts: []w.FieldPart{
											w.Drop{
												Trigger: w.Pick{Value: "Europe/Moscow"},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.Input{Kind: w.InputSearch, Placeholder: "Find a zone", Label: "Find a zone"},
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
								Rubric: w.Rubric{Text: "Progress"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Status", Parts: []w.FieldPart{w.Pick{Value: "Needs action"}}, Width: w.FieldSM},
									w.Field{Label: "Priority", Parts: []w.FieldPart{w.Pick{Value: "No priority"}}, Width: w.FieldSM},
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
								Items:  []w.FieldSetPart{w.Field{Parts: []w.FieldPart{w.Button{Label: "Add reminder"}}}},
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
									w.Field{Label: "Tags", Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Tags"}}, Width: w.FieldMD},
									w.Field{
										Label: "Link",
										Parts: []w.FieldPart{w.Input{Placeholder: "optional", Label: "Link"}},
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
															File: &w.FileField{Form: "", Name: "file", Label: "From this computer", List: "task-new-attachments"},
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
										ID:    "task-new-attachments",
										Place: w.Grown,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Repeats"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Frequency", Parts: []w.FieldPart{w.Pick{Value: "Does not repeat"}}, Width: w.FieldSM},
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
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("The advanced value takes precedence when set. A task repeats from its start date, or from the due date if it has none."),
												},
											},
										},
										Width: w.FieldXL,
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

// Tasks19 is frame «Связанная задача · какую копию отметить, лист поверх списка» (s-tasks-19).
func Tasks19() Screen {
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
								Meta:    "7",
								Current: true,
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
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 of 3 lists · nextcloud · 6 tasks")}}},
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
								w.Seg{Options: []w.SegOption{{Label: "Open 5", On: true}, {Label: "Done 1"}, {Label: "All 6"}}},
								w.Seg{
									Options:   []w.SegOption{{Label: "By due", On: true}, {Label: "By priority"}, {Label: "Recently changed"}},
									Secondary: true,
								},
								w.Button{Label: "Columns", Icon: "cols"},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Overdue", Form: w.RubricInline}, Count: &w.Num{Value: 2, Alert: true}},
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
								w.DueRow{
									Tick:     w.Tick{Label: "Done"},
									Mark:     w.Mark{Colour: "#3F6E8C"},
									Title:    []w.Inline{w.Text("Ответить подрядчику по смете")},
									Due:      "8 Aug",
									Progress: "40%",
									Source:   "nextcloud · Работа",
									Overdue:  true,
									Href:     "#",
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "This week", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
						w.List{
							Rows: []w.ListRow{
								w.DueRow{
									Tick:       w.Tick{Label: "Done"},
									Mark:       w.Mark{Colour: "#4A6B52"},
									Title:      []w.Inline{w.Text("Собрать чемодан")},
									Strong:     true,
									TitleBadge: &w.Badge{Label: "linked ×2", Tone: w.BadgeLinked},
									Due:        "12 Aug 18:00",
									Progress:   "—",
									Source:     "2 sources",
									SourceHint: "nextcloud · Дела, nextcloud · Работа",
									On:         true,
									Href:       "#",
								},
								w.DueRow{
									Tick:     w.Tick{Label: "Done"},
									Mark:     w.Mark{Colour: "#3F6E8C"},
									Title:    []w.Inline{w.Text("Дописать раздел про кэш")},
									Due:      "14 Aug",
									Progress: "70%",
									Source:   "nextcloud · Работа",
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
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/tasks/done",
					Icon:   "task",
					Title:  "Which copy to mark done?",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("Собрать чемодан · 2 copies")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
					},
					Body: []w.SheetPart{
						w.SheetNote{Text: "A copy that is not ticked stays as it is."},
						w.Line{
							Items: []w.LinePart{
								w.Mark{Colour: "#4A6B52", Size: w.MarkH11},
								w.Check{On: true, Name: "copy", Value: "nextcloud/dela", Parts: []w.CheckPart{w.Text("nextcloud · Дела")}},
								w.Dim{Parts: []w.Inline{w.Text("due 12 Aug 18:00")}, Size: w.Dim11_5},
							},
							Gap:   w.LineGap8,
							Block: true,
						},
						w.Line{
							Items: []w.LinePart{
								w.Mark{Colour: "#3F6E8C", Size: w.MarkH11},
								w.Check{Name: "copy", Value: "nextcloud/rabota", Parts: []w.CheckPart{w.Text("nextcloud · Работа")}},
								w.Dim{Parts: []w.Inline{w.Text("Due, priority differ · due 12 Aug")}, Size: w.Dim11_5},
							},
							Gap:   w.LineGap8,
							Block: true,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Mark done", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/tasks"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGap7,
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
