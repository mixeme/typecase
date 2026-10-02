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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"},
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
							Meta:    "30",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "21"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "9"},
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 of 3 lists · nextcloud · 30 tasks")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New task", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "up"},
								w.Button{Label: "Export", Icon: "down"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Seg{Options: []w.SegOption{{Label: "Open 12", On: true}, {Label: "Done 48"}, {Label: "All 60"}}},
							w.Seg{
								Options:   []w.SegOption{{Label: "By due", On: true}, {Label: "By priority"}, {Label: "Recently changed"}},
								Secondary: true,
							},
							w.Button{Label: "Columns", Icon: "cols"},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Overdue", Form: w.RubricInline}, Count: &w.Num{Value: 3, Alert: true}},
					w.List{
						Rows: []w.ListRow{
							w.DueRow{
								Tick:     w.Tick{Label: "Done"},
								Mark:     w.Mark{Colour: "#4A6B52"},
								Title:    []w.Inline{w.Text("Отправить документы в МФЦ")},
								Tags:     []w.RowTag{w.Tag{Label: "дом"}},
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
								Tags:     []w.RowTag{w.Tag{Label: "релиз"}},
								Due:      "8 Aug",
								Progress: "40%",
								Source:   "nextcloud · Работа",
								Overdue:  true,
								Href:     "#",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "This week", Form: w.RubricInline}, Count: &w.Num{Value: 5}},
					w.List{
						Rows: []w.ListRow{
							w.DueRow{
								Tick:       w.Tick{Label: "Done"},
								Mark:       w.Mark{Colour: "#4A6B52"},
								Title:      []w.Inline{w.Text("Собрать чемодан ")},
								Strong:     true,
								TitleBadge: &w.Badge{Label: "linked ×2", Tone: w.BadgeLinked},
								Tags:       []w.RowTag{w.Tag{Label: "дом"}},
								Due:        "12 Aug 18:00",
								Progress:   "—",
								Source:     "2 sources",
								SourceHint: "nextcloud · Дела, radicale · vps · Личное",
								On:         true,
								Href:       "#",
							},
							w.DueRow{
								Tick:     w.Tick{Label: "Done"},
								Mark:     w.Mark{Colour: "#3F6E8C"},
								Title:    []w.Inline{w.Text("Дописать раздел про кэш")},
								Tags:     []w.RowTag{w.Tag{Label: "релиз"}},
								Due:      "14 Aug",
								Progress: "70%",
								Source:   "nextcloud · Работа",
								Href:     "#",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "No due date", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Дела")}}},
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
							{Term: "Status", Value: []w.ValuePart{w.Text("Needs action")}},
							{Term: "Due", Value: []w.ValuePart{w.Text("2026-08-12 18:00")}, Face: w.ValueMono},
							{Term: "Priority", Value: []w.ValuePart{w.Text("5")}, Face: w.ValueMono},
							{Term: "Complete", Value: []w.ValuePart{w.Text("0%")}, Face: w.ValueMono},
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
							w.Rubric{Text: "Related"},
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{
										w.Span{Face: w.FaceAccent, Text: "Поездка в Казань"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· event")}},
									},
								},
							},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{
								Text: "Attachments",
								Items: []w.RubricPart{
									w.Drop{
										Trigger: w.Button{Label: "Attach", Menu: true, Small: true},
										Menu: w.Menu{
											Items: []w.MenuPart{
												w.MenuItem{
													Icon:  "up",
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
							w.SecHint{
								Text:  "The same block events and notes have: up to 2 MB the file is written into the task itself, above that it goes to the attachments folder and the task keeps a link. Pasting and dropping work the same way.",
								Apart: true,
							},
						},
					},
					w.Danger{Button: w.Button{Label: "Delete task", Kind: w.ButtonDangerSolid}},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "21", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "9"},
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
							w.HeadSub{
								Parts: []w.HeadSubPart{w.Text("nextcloud · Дела · "), w.Span{Face: w.FaceDanger, Text: "overdue by 3 days"}},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Edit", Icon: "pencil", Kind: w.ButtonPrimary},
								w.Button{Label: "Mark done", Tick: &w.Tick{Glyph: true}},
								w.Button{Label: "Print", Icon: "print"},
								w.Button{Icon: "dots", Title: "More"},
							},
						},
					},
					w.Fields{
						Items: []w.FieldsItem{
							{Term: "Status", Value: []w.ValuePart{w.Text("Needs action")}},
							{
								Term: "Due",
								Value: []w.ValuePart{
									w.Mono{Text: "2026-08-12 18:00"},
									w.Text(" "),
									w.Span{Face: w.FaceDanger, Text: "· overdue by 3 days"},
								},
							},
							{Term: "Complete", Value: []w.ValuePart{w.Text("0%")}, Face: w.ValueMono},
							{Term: "Priority", Value: []w.ValuePart{w.Text("5")}, Face: w.ValueMono},
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
										Trigger: w.Button{Label: "Attach", Menu: true, Small: true},
										Menu: w.Menu{
											Items: []w.MenuPart{
												w.MenuItem{
													Icon:  "up",
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
										w.Link{Text: "radicale · vps · Личное", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· Time differ")}},
									},
								},
								Gap: 5,
							},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Source"},
							w.Diag{Text: "Path  /dav.php/calendars/mix/tasks/c41a.ics\nRead  2 min ago · not changed on server"},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "21", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "9"},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH20}, Text: "Отправить документы в МФЦ"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{w.Text("nextcloud · Дела · "), w.Badge{Label: "overdue by 3 days", Tone: w.BadgeAlert}},
							},
						},
						Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Save", Kind: w.ButtonPrimary}}},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Summary",
								Parts: []w.FieldPart{w.Input{Value: "Отправить документы в МФЦ", Label: "Summary", Focus: true}},
								Width: w.FieldFull,
							},
							w.Field{
								Label: "Description",
								Parts: []w.FieldPart{w.Area{Text: "Опись, копия паспорта и заявление. Приём до 18:00, окно 4.", Label: "Text"}},
								Width: w.FieldFull,
							},
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
										Parts: []w.FieldPart{w.Input{Placeholder: "--:--", Label: "Start time", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Due date",
										Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Value: "2026-08-07", Label: "Due date", Mono: true}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Due time",
										Parts: []w.FieldPart{w.Input{Placeholder: "--:--", Label: "Due time", Mono: true}},
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
								Rubric: w.Rubric{Text: "Other"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Location",
										Parts: []w.FieldPart{w.Input{Placeholder: "optional", Label: "Location"}},
										Width: w.FieldMD,
									},
									w.Field{Label: "Tags", Parts: []w.FieldPart{w.Input{Value: "дом", Label: "Tags"}}, Width: w.FieldMD},
									w.Field{
										Label: "Related to",
										Parts: []w.FieldPart{w.TagInput{Parts: []w.TagInputPart{w.Tag{Label: "Забрать документы", Removable: true}}}},
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
												Trigger: w.Button{Label: "Attach", Menu: true, Icon: "up"},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{
															Icon:  "up",
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
											w.Dim{Parts: []w.Inline{w.Text("Ничего не вложено")}, Size: w.Dim11_5, Variant: w.DimWrap},
										},
										Gap:   w.LineGap8,
										Block: true,
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
												Value: []w.ValuePart{w.Text("/remote.php/dav/calendars/mix/dela/aa10.ics")},
												Face:  w.ValueMono,
											},
											{Term: "Read", Value: []w.ValuePart{w.Text("2 min ago · не менялась на сервере")}},
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
