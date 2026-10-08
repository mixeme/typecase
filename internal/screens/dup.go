// Раздел макета: 7.7 Дубликаты и слияние.

package screens

import w "carrel.dev/typecase/widget"

// Dup1 is frame «Дубликаты · группы и решения» (s-dup-1).
func Dup1() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 5}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 6,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Sources", Items: []w.RubricPart{w.Mono{Text: "12 / 12", Variant: w.MonoTight}}},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "3"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "4"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Meta: "1"},
						w.TickRow{Kind: w.TickRowExternal, Mark: w.Mark{Colour: "#8C6239"}, Label: "Праздники РФ", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "4"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "1"},
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
						Title: &w.HeadTitle{Text: "Duplicates"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("21 records compared · 2 to look at")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Seg{
									Options: []w.SegOption{
										{Label: "All 2", On: true},
										{Label: "Contacts 1"},
										{Label: "Events 1"},
										{Label: "Tasks 0"},
										{Label: "Notes 0"},
									},
								},
								w.Field{
									Label: "Match threshold",
									Parts: []w.FieldPart{
										w.Input{Name: "threshold", Value: "60", Label: "Match threshold", Mono: true, Variant: w.InputW52},
										w.Button{Label: "Weights", Small: true},
									},
									Width:  w.FieldInline,
									Action: "/app/duplicates",
								},
							},
							Stack: true,
						},
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("Linking changes nothing on any server: the records stay where they are and are shown as one. Merging on the server is a separate, destructive step you are asked to confirm."),
						},
						Tone: w.MsgAccent,
					},
					w.MainResults{
						Items: []w.ResultsPart{
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "Артём Белов",
										Items: []w.RubricPart{
											w.RubricAside{
												Text: "",
												Items: []w.RubricAsidePart{
													w.Badge{Label: "candidate ×2"},
													w.Mono{Text: "80 points · same phone, same name", Variant: w.MonoTightDusty},
												},
											},
										},
									},
									w.List{
										Rows: []w.ListRow{
											w.FaceRow{
												Mark:     w.Mark{Colour: "#4A6B52"},
												Avatar:   w.Avatar{Initials: "АБ"},
												Title:    []w.Inline{w.Text("Артём Белов")},
												Subtitle: "+7 903 118-22-04",
												Detail:   "a.belov@example.org",
												Caption:  "Тензор",
												Source:   "nextcloud · Контакты",
												Inset:    true,
											},
											w.FaceRow{
												Mark:     w.Mark{Colour: "#A05A2C"},
												Avatar:   w.Avatar{Photo: true},
												Title:    []w.Inline{w.Text("Белов Артём")},
												Subtitle: "+7 903 118-22-04",
												Detail:   "belov@team.example",
												Caption:  "—",
												Source:   "nextcloud · Архив",
												Inset:    true,
											},
										},
										Inset: true,
									},
									w.CardRow{
										Kind: w.CardTop11,
										Gap:  22,
										Items: []w.CardRowPart{
											w.Field{
												Label: "Display name",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Check{On: true, Name: "choice.FN", Value: "1", Radio: true, Parts: []w.CheckPart{w.Text("Артём Белов")}},
															w.Check{Name: "choice.FN", Value: "2", Radio: true, Parts: []w.CheckPart{w.Text("Белов Артём")}},
														},
														Gap:   w.LineGap12,
														Block: true,
													},
												},
											},
											w.Field{
												Label: "Organisation",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Check{On: true, Name: "choice.ORG", Value: "1", Radio: true, Parts: []w.CheckPart{w.Text("Тензор")}},
															w.Check{Name: "choice.ORG", Value: "2", Radio: true, Parts: []w.CheckPart{w.Text("—")}},
														},
														Gap:   w.LineGap12,
														Block: true,
													},
												},
											},
											w.Field{
												Label: "Photo",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Check{Name: "choice.PHOTO", Value: "1", Radio: true, Parts: []w.CheckPart{w.Text("nextcloud · Контакты")}},
															w.Check{
																On:    true,
																Name:  "choice.PHOTO",
																Value: "2",
																Radio: true,
																Parts: []w.CheckPart{w.Text("nextcloud · Архив")},
															},
														},
														Gap:   w.LineGap12,
														Block: true,
													},
												},
											},
										},
									},
									w.CardRow{
										Kind: w.CardActs,
										Gap:  8,
										Items: []w.CardRowPart{
											w.Button{Label: "Link them", Kind: w.ButtonPrimary, Name: "op", Value: "link", Submit: true},
											w.Button{Label: "Not duplicates", Name: "op", Value: "not", Submit: true},
											w.Button{
												Label:  "Merge on the server",
												Kind:   w.ButtonDanger,
												Place:  w.Pushed,
												Name:   "op",
												Value:  "merge",
												Submit: true,
											},
										},
									},
								},
								Action: "/app/duplicates/g1",
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "Дежурство: первая линия",
										Items: []w.RubricPart{
											w.RubricAside{
												Text: "",
												Items: []w.RubricAsidePart{
													w.Badge{Label: "candidate ×2"},
													w.Mono{Text: "70 points · same title, same start", Variant: w.MonoTightDusty},
												},
											},
										},
									},
									w.List{
										Rows: []w.ListRow{
											w.TimeRow{
												Mark:    w.Mark{Colour: "#3F6E8C"},
												Time:    "all day",
												Title:   []w.Inline{w.Text("Дежурство: первая линия")},
												Detail:  "—",
												Caption: &w.Badge{Label: "repeats weekly"},
												Source:  "nextcloud · Работа",
												Inset:   true,
											},
											w.TimeRow{
												Mark:    w.Mark{Colour: "#8C6239"},
												Time:    "all day",
												Title:   []w.Inline{w.Text("Дежурство: первая линия")},
												Detail:  "—",
												Caption: &w.Badge{Label: "repeats weekly"},
												Source:  "radicale · vps · Дежурства",
												Inset:   true,
											},
										},
										Inset: true,
									},
									w.CardHint{
										Parts: []w.CardTextPart{
											w.Text("The same duty in two calendars, different UIDs. Agenda shows it twice; here you decide whether to keep both or write one and delete the other."),
										},
										Line: true,
									},
									w.CardRow{
										Kind: w.CardActs,
										Gap:  8,
										Items: []w.CardRowPart{
											w.Button{Label: "Link them", Kind: w.ButtonPrimary, Name: "op", Value: "link", Submit: true},
											w.Button{Label: "Not duplicates", Name: "op", Value: "not", Submit: true},
											w.Button{
												Label:  "Merge on the server",
												Kind:   w.ButtonDanger,
												Place:  w.Pushed,
												Name:   "op",
												Value:  "merge",
												Submit: true,
											},
										},
									},
								},
								Action: "/app/duplicates/g2",
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "Собрать чемодан",
										Items: []w.RubricPart{
											w.RubricAside{
												Text: "",
												Items: []w.RubricAsidePart{
													w.Badge{Label: "linked ×2", Tone: w.BadgeLinked},
													w.Mono{Text: "70 points · same title, same due date", Variant: w.MonoTightDusty},
												},
											},
										},
									},
									w.List{
										Rows: []w.ListRow{
											w.DueRow{
												Tick:     w.Tick{Label: "Done"},
												Mark:     w.Mark{Colour: "#4A6B52"},
												Title:    []w.Inline{w.Text("Собрать чемодан")},
												Tags:     []w.RowTag{w.Tag{Label: "дом"}},
												Due:      "12 Aug 18:00",
												Progress: "—",
												Source:   "nextcloud · Дела",
												Inset:    true,
											},
											w.DueRow{
												Tick:     w.Tick{Label: "Done"},
												Mark:     w.Mark{Colour: "#6B4E7A"},
												Title:    []w.Inline{w.Text("Собрать чемодан")},
												Due:      "12 Aug",
												Progress: "—",
												Source:   "radicale · vps · Личное",
												Inset:    true,
											},
										},
										Inset: true,
									},
									w.CardRow{
										Kind: w.CardActs,
										Gap:  8,
										Items: []w.CardRowPart{
											w.Button{Label: "Unlink", Name: "op", Value: "unlink", Submit: true},
											w.Button{
												Label:  "Merge on the server",
												Kind:   w.ButtonDanger,
												Place:  w.Pushed,
												Name:   "op",
												Value:  "merge",
												Submit: true,
											},
										},
									},
								},
								Action: "/app/duplicates/g3",
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "Что почитать про CRDT",
										Items: []w.RubricPart{
											w.RubricAside{
												Text: "",
												Items: []w.RubricAsidePart{
													w.Badge{Label: "linked ×2", Tone: w.BadgeLinked},
													w.Mono{Text: "70 points · same title, same date", Variant: w.MonoTightDusty},
												},
											},
										},
									},
									w.List{
										Rows: []w.ListRow{
											w.ExcerptRow{
												Mark:    w.Mark{Colour: "#4A6B52"},
												Title:   []w.Inline{w.Text("Что почитать про CRDT")},
												Excerpt: "Кляйн, глава 5; статья Шапиро 2011…",
												Tags:    []w.RowTag{w.Tag{Label: "книги"}},
												Date:    "8 Aug",
												Source:  "nextcloud · Записи",
												Inset:   true,
											},
											w.ExcerptRow{
												Mark:    w.Mark{Colour: "#6B4E7A"},
												Title:   []w.Inline{w.Text("Что почитать про CRDT")},
												Excerpt: "Кляйн, глава 5; статья Шапиро 2011…",
												Date:    "8 Aug",
												Source:  "radicale · vps · Дневник",
												Inset:   true,
											},
										},
										Inset: true,
									},
									w.CardHint{
										Parts: []w.CardTextPart{
											w.Text("Same title and the same date: a note copied into a second notebook. Linked, the two are shown as one row; merging writes one and deletes the other."),
										},
										Line: true,
									},
									w.CardRow{
										Kind: w.CardActs,
										Gap:  8,
										Items: []w.CardRowPart{
											w.Button{Label: "Unlink", Name: "op", Value: "unlink", Submit: true},
											w.Button{
												Label:  "Merge on the server",
												Kind:   w.ButtonDanger,
												Place:  w.Pushed,
												Name:   "op",
												Value:  "merge",
												Submit: true,
											},
										},
									},
								},
								Action: "/app/duplicates/g4",
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "Анна Ковалёва",
										Items: []w.RubricPart{
											w.RubricAside{
												Text: "",
												Items: []w.RubricAsidePart{
													w.Badge{Label: "linked ×3", Tone: w.BadgeLinked},
													w.Mono{Text: "80 points · same email, same name", Variant: w.MonoTightDusty},
												},
											},
										},
									},
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
												Inset:    true,
											},
											w.FaceRow{
												Mark:     w.Mark{Colour: "#6B4E7A"},
												Avatar:   w.Avatar{Photo: true},
												Title:    []w.Inline{w.Text("Анна Ковалёва")},
												Subtitle: "+7 912 445-10-88",
												Detail:   "a.kovaleva@example.org",
												Caption:  "—",
												Source:   "radicale · vps · Команда",
												Inset:    true,
											},
											w.FaceRow{
												Mark:        w.Mark{Colour: "#9A9280"},
												Avatar:      w.Avatar{Initials: "АК"},
												Title:       []w.Inline{w.Text("А. Ковалёва")},
												Subtitle:    "—",
												Detail:      "a.kovaleva@example.org",
												Caption:     "Слово",
												Source:      "baikal · Общий",
												SourceBadge: &w.Badge{Label: "ro"},
												Inset:       true,
											},
										},
										Inset: true,
									},
									w.CardRow{
										Kind: w.CardTop11,
										Gap:  22,
										Items: []w.CardRowPart{
											w.Field{
												Label: "Photo",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Check{Name: "choice.PHOTO", Value: "1", Radio: true, Parts: []w.CheckPart{w.Text("nextcloud · Контакты")}},
															w.Check{
																On:    true,
																Name:  "choice.PHOTO",
																Value: "2",
																Radio: true,
																Parts: []w.CheckPart{w.Text("radicale · vps · Команда")},
															},
														},
														Gap:   w.LineGap12,
														Block: true,
													},
												},
											},
										},
									},
									w.CardRow{
										Kind: w.CardActs,
										Gap:  8,
										Items: []w.CardRowPart{
											w.Button{Label: "Keep these values", Kind: w.ButtonPrimary, Name: "op", Value: "keep", Submit: true},
											w.Button{Label: "Unlink", Name: "op", Value: "unlink", Submit: true},
											w.Button{Label: "Merge on the server", Kind: w.ButtonDanger, Place: w.Pushed, State: w.ButtonDisabled},
										},
									},
									w.CardText{
										Kind:  w.CardWhy,
										Parts: []w.CardTextPart{w.Text("Общий is read-only: its record could be neither rewritten nor deleted.")},
									},
								},
								Action: "/app/duplicates/g5",
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "Сергей Орлов",
										Items: []w.RubricPart{
											w.RubricAside{
												Text:  "",
												Items: []w.RubricAsidePart{w.Badge{Label: "not duplicates"}, w.Mono{Text: "60 points · same phone", Variant: w.MonoTightDusty}},
											},
										},
									},
									w.List{
										Rows: []w.ListRow{
											w.FaceRow{
												Mark:     w.Mark{Colour: "#4A6B52"},
												Avatar:   w.Avatar{Initials: "СО"},
												Title:    []w.Inline{w.Text("Сергей Орлов")},
												Subtitle: "+7 495 120-45-00",
												Detail:   "s.orlov@slovo.example",
												Caption:  "Издательство «Слово»",
												Source:   "nextcloud · Контакты",
												Inset:    true,
											},
											w.FaceRow{
												Mark:     w.Mark{Colour: "#6B4E7A"},
												Avatar:   w.Avatar{Initials: "МО"},
												Title:    []w.Inline{w.Text("Мария Орлова")},
												Subtitle: "+7 495 120-45-00",
												Detail:   "m.orlova@slovo.example",
												Caption:  "Издательство «Слово»",
												Source:   "radicale · vps · Команда",
												Inset:    true,
											},
										},
										Inset: true,
									},
									w.CardRow{
										Kind:  w.CardActs,
										Gap:   8,
										Items: []w.CardRowPart{w.Button{Label: "Offer again", Name: "op", Value: "offer", Submit: true}},
									},
								},
								Action: "/app/duplicates/g6",
							},
						},
					},
				},
			},
		},
	}
}

// Dup2 is frame «Слияние на сервере · подтверждение» (s-dup-2).
func Dup2() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 5}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 6,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Sources", Items: []w.RubricPart{w.Mono{Text: "12 / 12", Variant: w.MonoTight}}},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "3"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "4"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Meta: "1"},
						w.TickRow{Kind: w.TickRowExternal, Mark: w.Mark{Colour: "#8C6239"}, Label: "Праздники РФ", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "4"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "1"},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{
								Title: &w.HeadTitle{Text: "Merge on the server"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Артём Белов · 2 records")}}},
							},
							w.Msg{
								Parts: []w.MsgPart{
									w.Strong("One record is rewritten and 1 other deleted."),
									w.Text(" The kept record is rewritten with the values you chose, and only once the server has kept it are the others deleted. This cannot be undone from here."),
								},
								Space: w.MsgFlushSides,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Kept and rewritten"},
									w.CardText{
										Kind:  w.CardLine,
										Parts: []w.CardTextPart{w.Text("Артём Белов "), w.Dim{Parts: []w.Inline{w.Text("· nextcloud · Контакты")}}},
									},
								},
								Variant: w.CardApart,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Deleted afterwards"},
									w.CardText{
										Kind:  w.CardLine,
										Parts: []w.CardTextPart{w.Text("Белов Артём "), w.Dim{Parts: []w.Inline{w.Text("· nextcloud · Архив")}}},
									},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "The kept record gains"},
									w.CardList{Items: []string{"EMAIL;TYPE=work: belov@team.example", "PHOTO (from nextcloud · Архив)"}},
								},
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Write the merge and delete the rest", Kind: w.ButtonDangerSolid},
									w.Button{Label: "Cancel"},
									w.Button{Label: "Keep Белов Артём instead", Place: w.Pushed},
								},
								Variant: w.FormFootNoRule,
							},
						},
					},
				},
			},
		},
	}
}

// Dup4 is frame «Слияние события» (s-dup-4).
func Dup4() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 5}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 6,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Sources", Items: []w.RubricPart{w.Mono{Text: "12 / 12", Variant: w.MonoTight}}},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "3"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "4"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Meta: "1"},
						w.TickRow{Kind: w.TickRowExternal, Mark: w.Mark{Colour: "#8C6239"}, Label: "Праздники РФ", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "4"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "1"},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{
								Title: &w.HeadTitle{Text: "Merge on the server"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Дежурство: первая линия · 2 records")}}},
							},
							w.Msg{
								Parts: []w.MsgPart{
									w.Strong("One record is rewritten and 1 other deleted."),
									w.Text(" The kept record is rewritten with the values you chose, and only once the server has kept it are the others deleted. This cannot be undone from here."),
								},
								Space: w.MsgFlushSides,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Kept and rewritten"},
									w.CardText{
										Kind:  w.CardLine,
										Parts: []w.CardTextPart{w.Text("Дежурство: первая линия "), w.Dim{Parts: []w.Inline{w.Text("· nextcloud · Работа")}}},
									},
								},
								Variant: w.CardApart,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Deleted afterwards"},
									w.CardText{
										Kind:  w.CardLine,
										Parts: []w.CardTextPart{w.Text("Дежурство: первая линия "), w.Dim{Parts: []w.Inline{w.Text("· radicale · vps · Дежурства")}}},
									},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "The kept record gains"},
									w.CardList{Items: []string{"2 reminders from the others"}},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Not carried over"},
									w.CardList{
										Items: []string{
											"1 rewritten occurrence of Дежурство: первая линия · radicale · vps · Дежурства, which belongs to its series and goes with it",
										},
									},
								},
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Write the merge and delete the rest", Kind: w.ButtonDangerSolid},
									w.Button{Label: "Cancel"},
									w.Button{Label: "Keep the radicale · vps copy instead", Place: w.Pushed},
								},
								Variant: w.FormFootNoRule,
							},
						},
					},
				},
			},
		},
	}
}

// Dup5 is frame «Слияние · записано не всё» (s-dup-5).
func Dup5() Screen {
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 5}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 6,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Sources", Items: []w.RubricPart{w.Mono{Text: "12 / 12", Variant: w.MonoTight}}},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "3"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "4"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Meta: "1"},
						w.TickRow{Kind: w.TickRowExternal, Mark: w.Mark{Colour: "#8C6239"}, Label: "Праздники РФ", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "4"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "1"},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{
								Title: &w.HeadTitle{Text: "Merge on the server"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Артём Белов · 2 records")}}},
							},
							w.Msg{
								Parts: []w.MsgPart{
									w.Strong("Written, but not every other record was deleted."),
									w.Text(" The merge was written and 0 records deleted; Белов Артём refused, and the rest were not tried. Nothing is undone."),
								},
								Tone:  w.MsgAlert,
								Space: w.MsgFlushSides,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "What happened"},
									w.CardList{
										Items: []string{
											"Kept and rewritten: Артём Белов · nextcloud · Контакты",
											"Still there: Белов Артём · nextcloud · Архив",
										},
									},
								},
								Variant: w.CardApart,
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Open the kept record", Kind: w.ButtonPrimary},
									w.Button{Label: "Back to duplicates"},
								},
								Variant: w.FormFootNoRule,
							},
						},
					},
				},
			},
		},
	}
}

// Dup6 is frame «Дубликаты · ничего не загружено» (s-dup-6).
func Dup6() Screen {
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
					On: 6,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Sources", Items: []w.RubricPart{w.Mono{Text: "0 / 12", Variant: w.MonoTight}}},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "—"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "—"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "—"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "—"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "—"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "—"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "—"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Meta: "—"},
						w.TickRow{Kind: w.TickRowExternal, Mark: w.Mark{Colour: "#8C6239"}, Label: "Праздники РФ", Meta: "—"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "—"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "—"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "—"},
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
						Title: &w.HeadTitle{Text: "Duplicates"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("0 records compared")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Field{
									Label: "Match threshold",
									Parts: []w.FieldPart{
										w.Input{Name: "threshold", Value: "60", Label: "Match threshold", Mono: true, Variant: w.InputW52},
										w.Button{Label: "Weights", Small: true},
									},
									Width:  w.FieldInline,
									Action: "/app/duplicates",
								},
							},
						},
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("Linking changes nothing on any server: the records stay where they are and are shown as one. Merging on the server is a separate, destructive step you are asked to confirm."),
						},
						Tone: w.MsgAccent,
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Strong("Nothing is loaded yet."),
							w.Text(" Duplicates are found only among the records this session has already read, so no server was asked. Open a list or a combined view, then come back."),
						},
					},
				},
			},
		},
	}
}
