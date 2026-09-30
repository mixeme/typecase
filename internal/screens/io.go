// Раздел макета: 7.9 Импорт и экспорт.

package screens

import w "carrel.dev/typecase/widget"

// Io1 is frame «Файл над списком · накладка перетаскивания» (s-io-1).
func Io1() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}},
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
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "248", Current: true},
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
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · 248 cards")}}},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "New contact", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
									w.Button{Label: "Import", Icon: "up"},
									w.Button{Label: "Export", Icon: "down"},
								},
							},
						},
						w.Bar{
							Items: []w.BarPart{
								w.Filter{
									Query:   w.Query{Value: "Filter the loaded cards", Label: "Filter the loaded cards"},
									Waiting: true,
									Variant: w.FilterH26W220,
								},
								w.Seg{Options: []w.SegOption{{Label: "A–Z", On: true}, {Label: "Recently changed"}}, Secondary: true},
								w.Button{Label: "Columns", Icon: "cols"},
							},
							Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "248"}}, Secondary: true},
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
							},
						},
					},
				},
			},
			Scrim: w.Scrim{
				Content: w.Empty{
					Icon: "up",
					Lead: "Drop to import",
					Parts: []w.EmptyPart{
						w.Aside{
							Face:  w.FaceLarge,
							Parts: []w.Inline{w.Text("One "), w.Mono{Text: ".vcf"}, w.Text(" file into "), w.Strong("nextcloud · Контакты")},
						},
						w.Dim{
							Parts:   []w.Inline{w.Text("The preview opens first: nothing is written until it has been read")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
					},
					Variant: w.EmptyDrop,
				},
				Pad:  w.ScrimPad38,
				Drop: "post-import",
			},
		},
	}
}

// Io2 is frame «Файл над открытой карточкой · мастер не открывается» (s-io-2).
func Io2() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}},
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
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "2"},
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
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New event", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "up"},
								w.Button{Label: "Export", Icon: "down"},
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Time:    "10:00",
								Title:   []w.Inline{w.Text("Синк-ап по релизу")},
								Strong:  true,
								Detail:  "Переговорная 2",
								Caption: &w.Badge{Label: "4 attendees"},
								Source:  "nextcloud · Личный",
								On:      true,
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
					w.ListGroup{Title: w.Rubric{Text: "Tuesday · 11 August", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#6B4E7A"},
								Time:   "09:30",
								Title:  []w.Inline{w.Text("Ревью архитектуры")},
								Detail: "Переговорная 1",
								Source: "radicale · vps · Проекты",
							},
						},
					},
				},
			},
			Side: &w.Side{
				Items: []w.SidePart{
					w.Head{
						Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH19}, Text: "Синк-ап по релизу", Size: w.HeadTitle19},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Личный")}}},
						Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "Edit", Small: true}, w.Close{Label: "Close the panel"}}},
					},
					w.Fields{
						Items: []w.FieldsItem{
							{Term: "When", Value: []w.ValuePart{w.Mono{Text: "10:00 – 11:00"}, w.Text(", Monday 10 August")}},
							{Term: "Where", Value: []w.ValuePart{w.Text("Переговорная 2")}},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Attachments", Items: []w.RubricPart{w.Button{Label: "Attach", Menu: true, Small: true}}},
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{
										w.Text("release-checklist.md "),
										w.Dim{Parts: []w.Inline{w.Text("text/markdown · 3 KB")}},
										w.Text(" "),
										w.Badge{Label: "in the event"},
									},
								},
								Gap: 6,
							},
							w.Empty{
								Lead: "Attach to this event",
								Parts: []w.EmptyPart{
									w.Dim{
										Parts: []w.Inline{
											w.Text("Over 2 MB it goes to "),
											w.Mono{Text: "Documents / carrel"},
											w.Text(" and the event keeps a link"),
										},
										Size:    w.Dim11_5,
										Variant: w.DimWrap,
									},
								},
								Variant: w.EmptyDropSide,
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
				},
				Variant: w.SideDrop,
			},
		},
	}
}

// Io3 is frame «Импорт · шаг 1 из 3, источник» (s-io-3).
func Io3() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "248", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "51"},
					},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Dialog{
						Items: []w.DialogPart{
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 1},
							w.Head{Title: &w.HeadTitle{Text: "Import contacts"}},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Into",
										Parts: []w.FieldPart{
											w.Pick{Value: "nextcloud · Контакты"},
											w.FieldHint{Parts: []w.Inline{w.Text("Any address book you may write to, on every connected server.")}},
										},
										Width: w.FieldW100,
									},
								},
								Variant: w.FormBottom12,
								Tight:   true,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "From a file"},
									w.Field{
										Label: "vCard file",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Input{Placeholder: "Разное (mix).vcf", Label: "Разное (mix).vcf", Variant: w.InputFill},
													w.Button{Label: "Browse"},
													w.Button{Label: "Preview file", Kind: w.ButtonPrimary},
												},
												Gap: w.LineGap8,
											},
										},
										Width: w.FieldW100,
									},
									w.CardHint{
										Parts: []w.CardTextPart{
											w.Text("vCard 2.1, 3.0 and 4.0. One file holds any number of cards; the calendar and the list of tasks take "),
											w.Mono{Text: ".ics"},
											w.Text(" here instead."),
										},
										Line: true,
									},
								},
							},
							w.Msg{
								Parts: []w.MsgPart{
									w.Text("Nothing is written until the preview has been read: the next step counts what was parsed and names what failed."),
								},
							},
						},
						Variant: w.Dialog560,
					},
				},
			},
		},
	}
}

// Io4 is frame «Импорт заметок · шаг 1 из 3, источник» (s-io-4).
func Io4() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "64", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "22"},
					},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Dialog{
						Items: []w.DialogPart{
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 1},
							w.Head{Title: &w.HeadTitle{Text: "Import notes"}},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Into",
										Parts: []w.FieldPart{
											w.Pick{Value: "nextcloud · Записи"},
											w.FieldHint{Parts: []w.Inline{w.Text("The notebook the notes are created in. Every source below lands there.")}},
										},
										Width: w.FieldW100,
									},
								},
								Variant: w.FormBottom12,
								Tight:   true,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "From a file"},
									w.Field{
										Label: "Markdown or a zip of them",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Input{Placeholder: "заметки-2026.zip", Label: "заметки-2026.zip", Variant: w.InputFill},
													w.Button{Label: "Browse"},
													w.Button{Label: "Preview file", Kind: w.ButtonPrimary},
												},
												Gap: w.LineGap8,
											},
										},
										Width: w.FieldW100,
									},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "From a WebDAV folder"},
									w.Tree{
										Rows: []w.TreeRow{
											{Icon: "file", Label: "Documents", Trail: w.Dim{Parts: []w.Inline{w.Text("sftpgo")}, Size: w.Dim11}},
											{Icon: "file", Label: "notes", Indent: w.TreeLevel22},
											{Icon: "file", Label: "carrel", Indent: w.TreeLevel22},
											{Icon: "file", Label: "Files", Trail: w.Dim{Parts: []w.Inline{w.Text("nextcloud")}, Size: w.Dim11}},
										},
										On: 2,
									},
									w.CardRow{
										Kind: w.CardTop8,
										Gap:  10,
										Items: []w.CardRowPart{
											w.FillHint{
												Parts: []w.CardTextPart{
													w.Text("Every "),
													w.Mono{Text: ".md"},
													w.Text(" directly in that folder. Subfolders are not walked, nothing on the server is moved or deleted — the files stay where they were."),
												},
											},
											w.Button{Label: "Preview folder", Kind: w.ButtonPrimary},
										},
									},
								},
								Variant: w.CardChosen,
							},
							w.Msg{
								Parts: []w.MsgPart{
									w.Text("Front matter is read when present: title, date, tags, UID and "),
									w.Mono{Text: "X-"},
									w.Text("properties exported earlier. A file without it becomes a note dated by its own timestamp."),
								},
							},
						},
						Variant: w.Dialog560,
					},
				},
			},
		},
	}
}

// Io5 is frame «Импорт контактов · шаг 2 из 3, предпросмотр» (s-io-5).
func Io5() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "248", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "51"},
					},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Dialog{
						Items: []w.DialogPart{
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 2},
							w.Head{
								Title: &w.HeadTitle{Text: "Import contacts"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mono{Text: "Разное (mix).vcf"}, w.Text(" · 84 KB · vCard 3.0")}}},
								Acts: &w.Acts{
									Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Import 122 contacts", Kind: w.ButtonPrimary}},
								},
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Into",
										Parts: []w.FieldPart{
											w.Pick{Value: "nextcloud · Контакты"},
											w.FieldHint{
												Parts: []w.Inline{w.Text("Any address book you may write to. The UID collisions below are counted against it.")},
											},
										},
										Width: w.FieldLG,
									},
								},
								Tight: true,
							},
							w.Stats{
								Items: []w.Stat{
									w.Stat{Value: "122", Label: "ready"},
									w.Stat{Value: "3", Label: "with errors", Error: true},
									w.Stat{Value: "5", Label: "UID collisions"},
								},
							},
							w.Table{
								Cols: []w.TableCol{{Label: "Source"}, {Label: "Name"}, {Label: "UID", Rank: w.Rank3}, {Label: "Result", Rank: w.Rank4}},
								Rows: []w.TableRow{
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Разное (mix).vcf")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Анна Ковалёва")}},
											{Parts: []w.CellPart{w.Text("4f2b-9c31-a0e7")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Ready")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Разное (mix).vcf")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Артём Белов")}},
											{Parts: []w.CellPart{w.Text("1a77-40de-bb02")}, Mono: true},
											{Parts: []w.CellPart{w.Text("New UID will be assigned")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Разное (mix).vcf")}, Mono: true},
											{Parts: []w.CellPart{w.Text("— (no FN)")}, Dim: true},
											{Parts: []w.CellPart{w.Text("—")}, Mono: true},
											{Parts: []w.CellPart{w.Text("vCard 2.1 quoted-printable: cannot decode line 12")}, Tone: w.CellDanger},
										},
									},
								},
							},
							w.Msg{
								Parts: []w.MsgPart{
									w.Text("Import always creates. A UID that already exists gets a new one and is listed in the report."),
								},
							},
						},
						Variant: w.DialogWide,
					},
				},
			},
		},
	}
}

// Io6 is frame «Импорт событий · шаг 2 из 3, предпросмотр» (s-io-6).
func Io6() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "7", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "4"},
					},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Dialog{
						Items: []w.DialogPart{
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 2},
							w.Head{
								Title: &w.HeadTitle{Text: "Import events"},
								Subs: []w.HeadSub{
									w.HeadSub{Parts: []w.HeadSubPart{w.Mono{Text: "Личный 2026.ics"}, w.Text(" · 210 KB · iCalendar 2.0")}},
								},
								Acts: &w.Acts{
									Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Import 84 events", Kind: w.ButtonPrimary}},
								},
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Into",
										Parts: []w.FieldPart{
											w.Pick{Value: "nextcloud · Личный"},
											w.FieldHint{
												Parts: []w.Inline{w.Text("Calendars only. A list of tasks is a different collection and takes different components.")},
											},
										},
										Width: w.FieldLG,
									},
								},
								Tight: true,
							},
							w.Stats{
								Items: []w.Stat{
									w.Stat{Value: "84", Label: "ready"},
									w.Stat{Value: "2", Label: "with errors", Error: true},
									w.Stat{Value: "3", Label: "UID collisions"},
									w.Stat{Value: "6", Label: "skipped"},
								},
							},
							w.Table{
								Cols: []w.TableCol{
									{Label: "Source"},
									{Label: "Summary"},
									{Label: "When", Rank: w.Rank3},
									{Label: "Result", Rank: w.Rank4},
								},
								Rows: []w.TableRow{
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Личный 2026.ics")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Дежурство: первая линия")}},
											{Parts: []w.CellPart{w.Text("Mon 09:00 · weekly until 2026-12-28")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Ready — the series as one record")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Личный 2026.ics")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Отпуск")}},
											{Parts: []w.CellPart{w.Text("2026-07-13 → 2026-07-27 · all day")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Ready")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Личный 2026.ics")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Созвон с подрядчиком")}},
											{Parts: []w.CellPart{w.Text("2026-08-17 10:00 Asia/Yekaterinburg")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Ready — time zone kept as sent")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Личный 2026.ics")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Купить билеты "), w.Badge{Label: "VTODO"}}},
											{Parts: []w.CellPart{w.Text("—")}, Mono: true},
											{
												Parts: []w.CellPart{w.Text("Skipped: a calendar takes events — import it into a list of tasks")},
												Tone:  w.CellDusty,
											},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Личный 2026.ics")}, Mono: true},
											{Parts: []w.CellPart{w.Text("— (no SUMMARY)")}, Dim: true},
											{Parts: []w.CellPart{w.Text("2026-05-04")}, Mono: true},
											{
												Parts: []w.CellPart{w.Text("DTSTART without DTEND or DURATION: cannot tell how long it lasts")},
												Tone:  w.CellDanger,
											},
										},
									},
								},
							},
							w.Msg{
								Parts: []w.MsgPart{
									w.Text("Import always creates. A repeating event keeps its rule: the series is one record, not fifty-two."),
								},
							},
						},
						Variant: w.DialogWide,
					},
				},
			},
		},
	}
}

// Io7 is frame «Импорт задач · шаг 2 из 3, предпросмотр» (s-io-7).
func Io7() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}},
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
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Dialog{
						Items: []w.DialogPart{
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 2},
							w.Head{
								Title: &w.HeadTitle{Text: "Import tasks"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mono{Text: "Дела.ics"}, w.Text(" · 12 KB · iCalendar 2.0")}}},
								Acts: &w.Acts{
									Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Import 29 tasks", Kind: w.ButtonPrimary}},
								},
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Into",
										Parts: []w.FieldPart{
											w.Pick{Value: "nextcloud · Дела"},
											w.FieldHint{
												Parts: []w.Inline{w.Text("Lists of tasks only. A calendar is a different collection: events go there, not here.")},
											},
										},
										Width: w.FieldLG,
									},
								},
								Tight: true,
							},
							w.Stats{
								Items: []w.Stat{
									w.Stat{Value: "29", Label: "ready"},
									w.Stat{Value: "1", Label: "with errors", Error: true},
									w.Stat{Value: "2", Label: "UID collisions"},
									w.Stat{Value: "5", Label: "skipped"},
								},
							},
							w.Table{
								Cols: []w.TableCol{
									{Label: "Source"},
									{Label: "Summary"},
									{Label: "Due", Rank: w.Rank3},
									{Label: "Result", Rank: w.Rank4},
								},
								Rows: []w.TableRow{
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Дела.ics")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Отправить документы в МФЦ")}},
											{Parts: []w.CellPart{w.Text("2026-08-28")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Ready — overdue, and stays overdue")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Дела.ics")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Ответить подрядчику по смете")}},
											{Parts: []w.CellPart{w.Text("—")}, Mono: true, Tone: w.CellDusty},
											{Parts: []w.CellPart{w.Text("Ready — a task without a date is a task")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Дела.ics")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Дописать раздел про кэш")}},
											{Parts: []w.CellPart{w.Text("2026-09-04")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Ready — 40 % done, priority 1")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Дела.ics")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Созвон с подрядчиком "), w.Badge{Label: "VEVENT"}}},
											{Parts: []w.CellPart{w.Text("2026-08-17")}, Mono: true},
											{
												Parts: []w.CellPart{w.Text("Skipped: a list of tasks takes tasks — import it into a calendar")},
												Tone:  w.CellDusty,
											},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Дела.ics")}, Mono: true},
											{Parts: []w.CellPart{w.Text("— (no SUMMARY)")}, Dim: true},
											{Parts: []w.CellPart{w.Text("—")}, Mono: true},
											{
												Parts: []w.CellPart{w.Text("Neither SUMMARY nor DESCRIPTION: nothing to name it by")},
												Tone:  w.CellDanger,
											},
										},
									},
								},
							},
							w.Msg{
								Parts: []w.MsgPart{
									w.Text("STATUS, PERCENT-COMPLETE and PRIORITY come across as they are: a task that was half done arrives half done."),
								},
							},
						},
						Variant: w.DialogWide,
					},
				},
			},
		},
	}
}

// Io8 is frame «Импорт заметок · шаг 2 из 3, предпросмотр» (s-io-8).
func Io8() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "64", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "22"},
					},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Dialog{
						Items: []w.DialogPart{
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 2},
							w.Head{
								Title: &w.HeadTitle{Text: "Import notes"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mono{Text: "заметки-2026.zip"}, w.Text(" · 1.2 MB · 64 files")}}},
								Acts: &w.Acts{
									Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Import 61 notes", Kind: w.ButtonPrimary}},
								},
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Into",
										Parts: []w.FieldPart{
											w.Pick{Value: "nextcloud · Записи"},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Any notebook you may write to. A WebDAV folder from the step before lands in this same table."),
												},
											},
										},
										Width: w.FieldLG,
									},
								},
								Tight: true,
							},
							w.Stats{
								Items: []w.Stat{
									w.Stat{Value: "61", Label: "ready"},
									w.Stat{Value: "2", Label: "with errors", Error: true},
									w.Stat{Value: "1", Label: "UID collision"},
									w.Stat{Value: "1", Label: "skipped"},
								},
							},
							w.Table{
								Cols: []w.TableCol{{Label: "File"}, {Label: "Title"}, {Label: "Date", Rank: w.Rank3}, {Label: "Result", Rank: w.Rank4}},
								Rows: []w.TableRow{
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("2026-03-11-podryadchik.md")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Встреча с подрядчиком")}},
											{Parts: []w.CellPart{w.Text("2026-03-11")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Ready — title, date and tags from front matter")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("idei/crdt.md")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Что почитать про CRDT")}},
											{Parts: []w.CellPart{w.Text("2026-08-04")}, Mono: true},
											{
												Parts: []w.CellPart{w.Text("Ready — no front matter: title from the first heading, date from the file")},
											},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("doroga.md")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Про дорогу до дачи")}},
											{Parts: []w.CellPart{w.Text("2026-06-30")}, Mono: true},
											{Parts: []w.CellPart{w.Text("New UID will be assigned")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("skan.pdf")}, Mono: true},
											{Parts: []w.CellPart{w.Text("—")}, Dim: true},
											{Parts: []w.CellPart{w.Text("—")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Skipped: not Markdown")}, Tone: w.CellDusty},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("pustoj.md")}, Mono: true},
											{Parts: []w.CellPart{w.Text("—")}, Dim: true},
											{Parts: []w.CellPart{w.Text("—")}, Mono: true},
											{
												Parts: []w.CellPart{w.Text("Empty file: nothing to store and nothing to name it by")},
												Tone:  w.CellDanger,
											},
										},
									},
								},
							},
							w.Msg{
								Parts: []w.MsgPart{
									w.Text("Markdown is stored as it was written — nothing is reformatted — and front matter Carrel exported earlier comes back: UID, tags, X-properties."),
								},
							},
						},
						Variant: w.DialogWide,
					},
				},
			},
		},
	}
}

// Io9 is frame «Отчёт об импорте · шаг 3» (s-io-9).
func Io9() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "248", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "51"},
					},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Dialog{
						Items: []w.DialogPart{
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 3},
							w.Head{
								Title: &w.HeadTitle{Text: "Import report"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}}},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Created"},
									w.CardText{Kind: w.CardNum, Parts: []w.CardTextPart{w.Text("122")}}},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "UID collisions", Items: []w.RubricPart{w.Mono{Text: "5", Variant: w.MonoTight}}},
									w.Diag{Text: "1a77-40de-bb02 → 9c02-4411-7ee1\n33b0-9ffa-1120 → 55da-8e10-02cc"},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Failed", Items: []w.RubricPart{w.Mono{Text: "3", Variant: w.MonoTight}}},
									w.Diag{Text: "line 12: cannot decode quoted-printable", Variant: w.DiagDanger},
								},
							},
							w.FormFoot{
								Items:   []w.FormFootPart{w.Button{Label: "Back to list", Kind: w.ButtonPrimary}},
								Variant: w.FormFootFlush,
							},
						},
						Variant: w.Dialog520,
					},
				},
			},
		},
	}
}
