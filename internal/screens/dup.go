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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 6,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Sources", Items: []w.RubricPart{w.Mono{Text: "12 / 12", Variant: w.MonoTight}}},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "6"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Meta: "1"},
						w.TickRow{Kind: w.TickRowExternal, Mark: w.Mark{Colour: "#8C6239"}, Label: "Праздники РФ", Meta: "3"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "4"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "3"},
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("32 records compared · 2 to look at")}}},
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
										w.Input{Name: "value", Value: "60", Label: "Match threshold", Mono: true, Variant: w.InputW52},
										w.Button{Label: "Weights", Small: true},
									},
									Width:  w.FieldInline,
									Action: "/app/duplicates/decide",
									Post:   true,
									Sends:  []w.Hidden{w.Hidden{Name: "csrf", Value: "7f3a91de"}, w.Hidden{Name: "do", Value: "threshold"}},
								},
							},
							Stack: true,
						},
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("Linking changes nothing on any server: the records are shown as one. Merging on the server is a separate, destructive step you confirm."),
						},
						Tone: w.MsgAccent,
					},
					w.MainResults{
						Items: []w.ResultsPart{
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "Вера Лебедева",
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
												Avatar:   w.Avatar{Photo: true},
												Title:    []w.Inline{w.Text("Лебедева Вера")},
												Subtitle: "+7 905 331-09-42",
												Detail:   "v.lebedeva@example.org",
												Caption:  "Библиотека №4",
												Source:   "nextcloud · Контакты",
												Inset:    true,
											},
											w.FaceRow{
												Mark:     w.Mark{Colour: "#A05A2C"},
												Avatar:   w.Avatar{Initials: "ВЛ"},
												Title:    []w.Inline{w.Text("Вера Лебедева")},
												Subtitle: "+7 905 331-09-42",
												Detail:   "vera@example.net",
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
															w.Check{Name: "choice.FN", Value: "1", Radio: true, Parts: []w.CheckPart{w.Text("Лебедева Вера · Контакты")}},
															w.Check{
																On:    true,
																Name:  "choice.FN",
																Value: "2",
																Radio: true,
																Parts: []w.CheckPart{w.Text("Вера Лебедева · Архив")},
															},
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
															w.Check{
																On:    true,
																Name:  "choice.ORG",
																Value: "1",
																Radio: true,
																Parts: []w.CheckPart{w.Text("Библиотека №4 · Контакты")},
															},
															w.Check{Name: "choice.ORG", Value: "2", Radio: true, Parts: []w.CheckPart{w.Text("— · Архив")}},
														},
														Gap:   w.LineGap12,
														Block: true,
													},
												},
											},
											w.Field{
												Label: "Email",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Check{
																On:    true,
																Name:  "choice.EMAIL",
																Value: "1",
																Parts: []w.CheckPart{w.Text("v.lebedeva@example.org · Контакты")},
															},
															w.Check{On: true, Name: "choice.EMAIL", Value: "2", Parts: []w.CheckPart{w.Text("vera@example.net · Архив")}},
															w.Check{
																On:    true,
																Name:  "choice.EMAIL",
																Value: "3",
																Parts: []w.CheckPart{w.Text("vera.l@work.example · Архив")},
															},
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
															w.Check{
																On:    true,
																Name:  "choice.PHOTO",
																Value: "1",
																Radio: true,
																Parts: []w.CheckPart{w.Text("nextcloud · Контакты")},
															},
															w.Check{Name: "choice.PHOTO", Value: "2", Radio: true, Parts: []w.CheckPart{w.Text("nextcloud · Архив")}},
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
											w.Button{Label: "Link them", Kind: w.ButtonPrimary, Name: "do", Value: "link", Submit: true},
											w.Button{Label: "Not duplicates", Name: "do", Value: "ignore", Submit: true},
											w.Button{
												Label:  "Merge on the server",
												Kind:   w.ButtonDanger,
												Place:  w.Pushed,
												Name:   "do",
												Value:  "merge",
												Submit: true,
											},
										},
									},
								},
								Action: "/app/duplicates/decide",
								Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "group", Value: "g1"}},
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
												Mark:   w.Mark{Colour: "#3F6E8C"},
												Time:   "all day",
												Title:  []w.Inline{w.Text("Дежурство: первая линия")},
												Detail: "—",
												Badges: []w.Badge{w.Badge{Label: "repeats weekly"}},
												Source: "nextcloud · Работа",
												Inset:  true,
											},
											w.TimeRow{
												Mark:   w.Mark{Colour: "#8C6239"},
												Time:   "all day",
												Title:  []w.Inline{w.Text("Дежурство: первая линия")},
												Detail: "—",
												Badges: []w.Badge{w.Badge{Label: "repeats weekly"}},
												Source: "radicale · vps · Дежурства",
												Inset:  true,
											},
										},
										Inset: true,
									},
									w.CardHint{
										Parts: []w.CardTextPart{w.Text("The same event in two calendars, different UIDs: Agenda shows it twice.")},
										Line:  true,
									},
									w.CardRow{
										Kind: w.CardActs,
										Gap:  8,
										Items: []w.CardRowPart{
											w.Button{Label: "Link them", Kind: w.ButtonPrimary, Name: "do", Value: "link", Submit: true},
											w.Button{Label: "Not duplicates", Name: "do", Value: "ignore", Submit: true},
											w.Button{
												Label:  "Merge on the server",
												Kind:   w.ButtonDanger,
												Place:  w.Pushed,
												Name:   "do",
												Value:  "merge",
												Submit: true,
											},
										},
									},
								},
								Action: "/app/duplicates/decide",
								Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "group", Value: "g2"}},
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
													w.Mono{Text: "220 points · same UID, same phone, same email, same name", Variant: w.MonoTightDusty},
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
												Mark:     w.Mark{Colour: "#A05A2C"},
												Avatar:   w.Avatar{Initials: "АК"},
												Title:    []w.Inline{w.Text("Анна Ковалёва")},
												Subtitle: "+7 912 445-10-88",
												Detail:   "a.kovaleva@example.org",
												Caption:  "Издательство «Слово»",
												Source:   "nextcloud · Архив",
												Inset:    true,
											},
											w.FaceRow{
												Mark:     w.Mark{Colour: "#6B4E7A"},
												Avatar:   w.Avatar{Initials: "АК"},
												Title:    []w.Inline{w.Text("Анна Ковалёва")},
												Subtitle: "+7 912 445-10-88",
												Detail:   "a.kovaleva@example.org",
												Caption:  "Издательство «Слово»",
												Source:   "radicale · vps · Команда",
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
												Label: "Phone · work",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Check{Name: "choice.TEL", Value: "1", Parts: []w.CheckPart{w.Text("+7 495 120-45-00 · Контакты")}},
															w.Check{On: true, Name: "choice.TEL", Value: "2", Parts: []w.CheckPart{w.Text("+7 495 120-45-77 · Команда")}},
														},
														Gap:   w.LineGap12,
														Block: true,
													},
												},
											},
											w.Field{
												Label: "Email",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Check{
																On:    true,
																Name:  "choice.EMAIL",
																Value: "1",
																Parts: []w.CheckPart{w.Text("anna.kovaleva@slovo.example · Контакты")},
															},
														},
														Gap:   w.LineGap12,
														Block: true,
													},
												},
											},
											w.Field{
												Label: "Title",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Check{
																On:    true,
																Name:  "choice.TITLE",
																Value: "1",
																Radio: true,
																Parts: []w.CheckPart{w.Text("Выпускающий редактор · Контакты")},
															},
															w.Check{Name: "choice.TITLE", Value: "2", Radio: true, Parts: []w.CheckPart{w.Text("— · Архив, Команда")}},
														},
														Gap:   w.LineGap12,
														Block: true,
													},
												},
											},
											w.Field{
												Label: "Address · work",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Check{
																On:    true,
																Name:  "choice.ADR",
																Value: "1",
																Parts: []w.CheckPart{w.Text("Москва, Тверская 12, стр. 3 · Контакты")},
															},
														},
														Gap:   w.LineGap12,
														Block: true,
													},
												},
											},
											w.Field{
												Label: "Birthday",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Check{
																On:    true,
																Name:  "choice.BDAY",
																Value: "1",
																Radio: true,
																Parts: []w.CheckPart{w.Text("1987-03-14 · Контакты")},
															},
															w.Check{Name: "choice.BDAY", Value: "2", Radio: true, Parts: []w.CheckPart{w.Text("— · Архив, Команда")}},
														},
														Gap:   w.LineGap12,
														Block: true,
													},
												},
											},
											w.Field{
												Label: "Note",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Check{
																On:    true,
																Name:  "choice.NOTE",
																Value: "1",
																Radio: true,
																Parts: []w.CheckPart{w.Text("Пишет по вторникам… · Контакты")},
															},
															w.Check{Name: "choice.NOTE", Value: "2", Radio: true, Parts: []w.CheckPart{w.Text("— · Архив, Команда")}},
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
											w.Button{Label: "Keep these values", Kind: w.ButtonPrimary, Name: "do", Value: "choose", Submit: true},
											w.Button{Label: "Unlink", Name: "do", Value: "unlink", Submit: true},
											w.Button{Label: "Merge on the server", Kind: w.ButtonDanger, Place: w.Pushed, State: w.ButtonDisabled},
										},
									},
									w.CardText{
										Kind:  w.CardWhy,
										Parts: []w.CardTextPart{w.Text("Команда is read-only: its record could be neither rewritten nor deleted.")},
									},
								},
								Action: "/app/duplicates/decide",
								Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "group", Value: "g3"}},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "Артём Белов",
										Items: []w.RubricPart{
											w.RubricAside{
												Text: "",
												Items: []w.RubricAsidePart{
													w.Badge{Label: "linked ×2", Tone: w.BadgeLinked},
													w.Mono{Text: "140 points · same phone, same email, same name", Variant: w.MonoTightDusty},
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
												Mark:     w.Mark{Colour: "#6B4E7A"},
												Avatar:   w.Avatar{Photo: true},
												Title:    []w.Inline{w.Text("Артём Белов")},
												Subtitle: "+7 903 118-22-04",
												Detail:   "a.belov@example.org",
												Caption:  "Тензор",
												Source:   "radicale · vps · Команда",
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
												Label: "Title",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Check{Name: "choice.TITLE", Value: "1", Radio: true, Parts: []w.CheckPart{w.Text("— · Контакты")}},
															w.Check{
																On:    true,
																Name:  "choice.TITLE",
																Value: "2",
																Radio: true,
																Parts: []w.CheckPart{w.Text("Ведущий инженер · Команда")},
															},
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
											w.Button{Label: "Keep these values", Kind: w.ButtonPrimary, Name: "do", Value: "choose", Submit: true},
											w.Button{Label: "Unlink", Name: "do", Value: "unlink", Submit: true},
											w.Button{Label: "Merge on the server", Kind: w.ButtonDanger, Place: w.Pushed, State: w.ButtonDisabled},
										},
									},
									w.CardText{
										Kind:  w.CardWhy,
										Parts: []w.CardTextPart{w.Text("Команда is read-only: its record could be neither rewritten nor deleted.")},
									},
								},
								Action: "/app/duplicates/decide",
								Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "group", Value: "g4"}},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "Синк-ап по релизу",
										Items: []w.RubricPart{
											w.RubricAside{
												Text: "",
												Items: []w.RubricAsidePart{
													w.Badge{Label: "linked ×3", Tone: w.BadgeLinked},
													w.Mono{Text: "70 points · same title, same start", Variant: w.MonoTightDusty},
												},
											},
										},
									},
									w.List{
										Rows: []w.ListRow{
											w.TimeRow{
												Mark:   w.Mark{Colour: "#4A6B52"},
												Time:   "10:00",
												Title:  []w.Inline{w.Text("Синк-ап по релизу")},
												Detail: "Переговорная 2",
												Badges: []w.Badge{w.Badge{Label: "repeats weekly"}},
												Source: "nextcloud · Личный",
												Inset:  true,
											},
											w.TimeRow{
												Mark:   w.Mark{Colour: "#6B4E7A"},
												Time:   "10:00",
												Title:  []w.Inline{w.Text("Синк-ап по релизу")},
												Detail: "Переговорная 2",
												Badges: []w.Badge{w.Badge{Label: "repeats weekly"}},
												Source: "radicale · vps · Проекты",
												Inset:  true,
											},
											w.TimeRow{
												Mark:   w.Mark{Colour: "#9A9280"},
												Time:   "10:30",
												Title:  []w.Inline{w.Text("Синк-ап по релизу")},
												Detail: "Переговорная 3",
												Badges: []w.Badge{w.Badge{Label: "repeats weekly"}},
												Source: "baikal · Общий",
												Inset:  true,
											},
										},
										Inset: true,
									},
									w.CardRow{
										Kind: w.CardTop11,
										Gap:  22,
										Items: []w.CardRowPart{
											w.Field{
												Label: "When",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Check{
																On:    true,
																Name:  "choice.WHEN",
																Value: "1",
																Radio: true,
																Parts: []w.CheckPart{w.Text("Mon 10:00 – 11:00, weekly · Личный, Проекты")},
															},
															w.Check{
																Name:  "choice.WHEN",
																Value: "2",
																Radio: true,
																Parts: []w.CheckPart{w.Text("Mon 10:30 – 11:30, weekly · Общий")},
															},
														},
														Gap:   w.LineGap12,
														Block: true,
													},
												},
											},
											w.Field{
												Label: "Where",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Check{
																On:    true,
																Name:  "choice.LOCATION",
																Value: "1",
																Radio: true,
																Parts: []w.CheckPart{w.Text("Переговорная 2 · Личный, Проекты")},
															},
															w.Check{
																Name:  "choice.LOCATION",
																Value: "2",
																Radio: true,
																Parts: []w.CheckPart{w.Text("Переговорная 3 · Общий")},
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
											w.Button{Label: "Keep these values", Kind: w.ButtonPrimary, Name: "do", Value: "choose", Submit: true},
											w.Button{Label: "Unlink", Name: "do", Value: "unlink", Submit: true},
											w.Button{
												Label:  "Merge on the server",
												Kind:   w.ButtonDanger,
												Place:  w.Pushed,
												Name:   "do",
												Value:  "merge",
												Submit: true,
											},
										},
									},
								},
								Action: "/app/duplicates/decide",
								Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "group", Value: "g5"}},
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
												Tick:     w.Tick{Sign: true},
												Mark:     w.Mark{Colour: "#4A6B52"},
												Title:    []w.Inline{w.Text("Собрать чемодан")},
												Due:      "12 Aug 18:00",
												Progress: "—",
												Source:   "nextcloud · Дела",
												Inset:    true,
											},
											w.DueRow{
												Tick:     w.Tick{Sign: true},
												Mark:     w.Mark{Colour: "#3F6E8C"},
												Title:    []w.Inline{w.Text("Собрать чемодан")},
												Due:      "12 Aug",
												Progress: "—",
												Source:   "nextcloud · Работа",
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
												Label: "Due",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Check{
																On:    true,
																Name:  "choice.DUE",
																Value: "1",
																Radio: true,
																Parts: []w.CheckPart{w.Text("12 Aug 18:00 · Дела")},
															},
															w.Check{Name: "choice.DUE", Value: "2", Radio: true, Parts: []w.CheckPart{w.Text("12 Aug · Работа")}},
														},
														Gap:   w.LineGap12,
														Block: true,
													},
												},
											},
											w.Field{
												Label: "Priority",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Check{Name: "choice.PRIORITY", Value: "1", Radio: true, Parts: []w.CheckPart{w.Text("5 · Дела")}},
															w.Check{
																On:    true,
																Name:  "choice.PRIORITY",
																Value: "2",
																Radio: true,
																Parts: []w.CheckPart{w.Text("3 · Работа")},
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
											w.Button{Label: "Keep these values", Kind: w.ButtonPrimary, Name: "do", Value: "choose", Submit: true},
											w.Button{Label: "Unlink", Name: "do", Value: "unlink", Submit: true},
											w.Button{
												Label:  "Merge on the server",
												Kind:   w.ButtonDanger,
												Place:  w.Pushed,
												Name:   "do",
												Value:  "merge",
												Submit: true,
											},
										},
									},
								},
								Action: "/app/duplicates/decide",
								Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "group", Value: "g6"}},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "Заметка со встречи 10.08",
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
												Title:   []w.Inline{w.Text("Заметка со встречи 10.08")},
												Excerpt: "Договорились заморозить ветку в среду; письмо пользователям пишет Анна…",
												Date:    "10 Aug",
												Source:  "nextcloud · Записи",
												Inset:   true,
											},
											w.ExcerptRow{
												Mark:    w.Mark{Colour: "#6B4E7A"},
												Title:   []w.Inline{w.Text("Заметка со встречи 10.08")},
												Excerpt: "Договорились заморозить ветку в среду…",
												Date:    "10 Aug",
												Source:  "nextcloud · Дневник",
												Inset:   true,
											},
										},
										Inset: true,
									},
									w.CardRow{
										Kind: w.CardTop11,
										Gap:  22,
										Items: []w.CardRowPart{
											w.Field{
												Label: "Text",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Check{
																On:    true,
																Name:  "choice.TEXT",
																Value: "1",
																Radio: true,
																Parts: []w.CheckPart{w.Text("nextcloud · Записи")},
															},
															w.Check{Name: "choice.TEXT", Value: "2", Radio: true, Parts: []w.CheckPart{w.Text("nextcloud · Дневник")}},
														},
														Gap:   w.LineGap12,
														Block: true,
													},
												},
											},
											w.Field{
												Label: "Tags",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Check{On: true, Name: "choice.CATEGORIES", Value: "1", Parts: []w.CheckPart{w.Text("релиз · Записи")}},
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
											w.Button{Label: "Keep these values", Kind: w.ButtonPrimary, Name: "do", Value: "choose", Submit: true},
											w.Button{Label: "Unlink", Name: "do", Value: "unlink", Submit: true},
											w.Button{
												Label:  "Merge on the server",
												Kind:   w.ButtonDanger,
												Place:  w.Pushed,
												Name:   "do",
												Value:  "merge",
												Submit: true,
											},
										},
									},
								},
								Action: "/app/duplicates/decide",
								Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "group", Value: "g7"}},
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
												Excerpt: "Кляйн, глава 5; статья Шапиро 2011; демка automerge…",
												Date:    "8 Aug",
												Source:  "nextcloud · Записи",
												Inset:   true,
											},
											w.ExcerptRow{
												Mark:    w.Mark{Colour: "#6B4E7A"},
												Title:   []w.Inline{w.Text("Что почитать про CRDT")},
												Excerpt: "Кляйн, глава 5; статья Шапиро 2011; демка automerge…",
												Date:    "8 Aug",
												Source:  "nextcloud · Дневник",
												Inset:   true,
											},
										},
										Inset: true,
									},
									w.CardHint{
										Parts: []w.CardTextPart{w.Text("Same title and the same date: a note copied into a second notebook.")},
										Line:  true,
									},
									w.CardRow{
										Kind: w.CardActs,
										Gap:  8,
										Items: []w.CardRowPart{
											w.Button{Label: "Unlink", Name: "do", Value: "unlink", Submit: true},
											w.Button{
												Label:  "Merge on the server",
												Kind:   w.ButtonDanger,
												Place:  w.Pushed,
												Name:   "do",
												Value:  "merge",
												Submit: true,
											},
										},
									},
								},
								Action: "/app/duplicates/decide",
								Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "group", Value: "g8"}},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "Анна Ковалёва",
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
												Avatar:   w.Avatar{Initials: "АК"},
												Title:    []w.Inline{w.Text("Анна Ковалёва")},
												Subtitle: "+7 495 120-45-00",
												Detail:   "a.kovaleva@example.org",
												Caption:  "Издательство «Слово»",
												Source:   "nextcloud · Контакты",
												Inset:    true,
											},
											w.FaceRow{
												Mark:     w.Mark{Colour: "#4A6B52"},
												Avatar:   w.Avatar{Initials: "АС"},
												Title:    []w.Inline{w.Text("Алексей Смирнов")},
												Subtitle: "+7 495 120-45-00",
												Detail:   "—",
												Caption:  "—",
												Source:   "nextcloud · Контакты",
												Inset:    true,
											},
										},
										Inset: true,
									},
									w.CardRow{
										Kind:  w.CardActs,
										Gap:   8,
										Items: []w.CardRowPart{w.Button{Label: "Offer again", Name: "do", Value: "again", Submit: true}},
									},
								},
								Action: "/app/duplicates/decide",
								Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "group", Value: "g9"}},
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 6,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Sources", Items: []w.RubricPart{w.Mono{Text: "12 / 12", Variant: w.MonoTight}}},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "6"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Meta: "1"},
						w.TickRow{Kind: w.TickRowExternal, Mark: w.Mark{Colour: "#8C6239"}, Label: "Праздники РФ", Meta: "3"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "4"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "3"},
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
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Вера Лебедева · 2 records")}}},
							},
							w.Msg{
								Parts: []w.MsgPart{w.Text("The other record is deleted. This cannot be undone.")},
								Space: w.MsgFlushSides,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Where the merged record stays"},
									w.Seg{
										Options: []w.SegOption{
											{Label: "nextcloud · Контакты", Href: "/app/duplicates/merge?group=g1&keep=1"},
											{Label: "nextcloud · Архив", Href: "/app/duplicates/merge?group=g1&keep=2", On: true},
										},
									},
									w.CardText{
										Kind:  w.CardLine,
										Parts: []w.CardTextPart{w.Dim{Parts: []w.Inline{w.Text("Вера Лебедева there is rewritten with the values you chose.")}}},
									},
								},
								Variant: w.CardApart,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Deleted afterwards"},
									w.CardText{
										Kind:  w.CardLine,
										Parts: []w.CardTextPart{w.Text("Лебедева Вера "), w.Dim{Parts: []w.Inline{w.Text("· nextcloud · Контакты")}}},
									},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "The record that stays gains"},
									w.CardList{
										Items: []string{"ORG: Библиотека №4", "EMAIL;TYPE=work: v.lebedeva@example.org", "PHOTO (from nextcloud · Контакты)"},
									},
								},
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Write the merge and delete the rest", Kind: w.ButtonDangerSolid},
									w.Button{Label: "Cancel"},
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 6,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Sources", Items: []w.RubricPart{w.Mono{Text: "12 / 12", Variant: w.MonoTight}}},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "6"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Meta: "1"},
						w.TickRow{Kind: w.TickRowExternal, Mark: w.Mark{Colour: "#8C6239"}, Label: "Праздники РФ", Meta: "3"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "4"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "3"},
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
								Parts: []w.MsgPart{w.Text("The other record is deleted. This cannot be undone.")},
								Space: w.MsgFlushSides,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Where the merged record stays"},
									w.Seg{
										Options: []w.SegOption{
											{Label: "nextcloud · Работа", Href: "/app/duplicates/merge?group=g2&keep=1", On: true},
											{Label: "radicale · vps · Дежурства", Href: "/app/duplicates/merge?group=g2&keep=2"},
										},
									},
									w.CardText{
										Kind: w.CardLine,
										Parts: []w.CardTextPart{
											w.Dim{Parts: []w.Inline{w.Text("Дежурство: первая линия there is rewritten with the values you chose.")}},
										},
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
									w.Rubric{Text: "The record that stays gains"},
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 6,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Sources", Items: []w.RubricPart{w.Mono{Text: "12 / 12", Variant: w.MonoTight}}},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "6"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Meta: "1"},
						w.TickRow{Kind: w.TickRowExternal, Mark: w.Mark{Colour: "#8C6239"}, Label: "Праздники РФ", Meta: "3"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Дела", Meta: "4"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "3"},
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
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Вера Лебедева · 2 records")}}},
							},
							w.Msg{
								Parts: []w.MsgPart{
									w.Strong("Merged, but nothing was deleted:"),
									w.Text(" Лебедева Вера refused, the rest were not tried."),
								},
								Tone:  w.MsgAlert,
								Space: w.MsgFlushSides,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "What happened"},
									w.CardList{
										Items: []string{
											"Kept and rewritten: Вера Лебедева · nextcloud · Архив",
											"Still there: Лебедева Вера · nextcloud · Контакты",
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
				w.Button{
					Label: "Refresh",
					Icon:  "refresh",
					Title: "Nothing is loaded yet: open a list first.",
					State: w.ButtonDisabled,
				},
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
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "—"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "—"},
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
										w.Input{Name: "value", Value: "60", Label: "Match threshold", Mono: true, Variant: w.InputW52},
										w.Button{Label: "Weights", Small: true},
									},
									Width:  w.FieldInline,
									Action: "/app/duplicates/decide",
									Post:   true,
									Sends:  []w.Hidden{w.Hidden{Name: "csrf", Value: "7f3a91de"}, w.Hidden{Name: "do", Value: "threshold"}},
								},
							},
						},
					},
					w.Notices{
						Items: []w.NoticesPart{
							w.Msg{
								Parts: []w.MsgPart{
									w.Text("Linking changes nothing on any server: the records are shown as one. Merging on the server is a separate, destructive step you confirm."),
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
					w.MainResults{
						Items: []w.ResultsPart{
							w.Rubric{Text: "Decided earlier", Items: []w.RubricPart{w.Mono{Text: "3", Variant: w.MonoTight}}},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "Анна Ковалёва",
										Items: []w.RubricPart{
											w.RubricAside{
												Text: "",
												Items: []w.RubricAsidePart{
													w.Badge{Label: "linked ×3", Tone: w.BadgeLinked},
													w.Mono{Text: "Not loaded in this session", Variant: w.MonoTightDusty},
												},
											},
										},
									},
									w.CardText{
										Kind: w.CardLine,
										Parts: []w.CardTextPart{
											w.Text("nextcloud · Контакты, nextcloud · Архив, radicale · vps · Команда "),
											w.Dim{Parts: []w.Inline{w.Text("· UID 4f2b-9c31-a0e7")}},
										},
									},
									w.CardRow{
										Kind:  w.CardActs,
										Gap:   8,
										Items: []w.CardRowPart{w.Button{Label: "Unlink", Name: "do", Value: "unlink", Submit: true}},
									},
								},
								Action: "/app/duplicates/decide",
								Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "group", Value: "g3"}},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "Артём Белов",
										Items: []w.RubricPart{
											w.RubricAside{
												Text: "",
												Items: []w.RubricAsidePart{
													w.Badge{Label: "linked ×2", Tone: w.BadgeLinked},
													w.Mono{Text: "Not loaded in this session", Variant: w.MonoTightDusty},
												},
											},
										},
									},
									w.CardText{
										Kind: w.CardLine,
										Parts: []w.CardTextPart{
											w.Text("nextcloud · Контакты, radicale · vps · Команда "),
											w.Dim{Parts: []w.Inline{w.Text("· UID c3f7-8a21-1b09, d41a-6f30-92be")}},
										},
									},
									w.CardRow{
										Kind:  w.CardActs,
										Gap:   8,
										Items: []w.CardRowPart{w.Button{Label: "Unlink", Name: "do", Value: "unlink", Submit: true}},
									},
								},
								Action: "/app/duplicates/decide",
								Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "group", Value: "g4"}},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "Анна Ковалёва · Алексей Смирнов",
										Items: []w.RubricPart{
											w.RubricAside{
												Text: "",
												Items: []w.RubricAsidePart{
													w.Badge{Label: "not duplicates"},
													w.Mono{Text: "Not loaded in this session", Variant: w.MonoTightDusty},
												},
											},
										},
									},
									w.CardText{
										Kind: w.CardLine,
										Parts: []w.CardTextPart{
											w.Text("nextcloud · Контакты "),
											w.Dim{Parts: []w.Inline{w.Text("· UID 4f2b-9c31-a0e7, 9a1c-03de-77f2")}},
										},
									},
									w.CardRow{
										Kind:  w.CardActs,
										Gap:   8,
										Items: []w.CardRowPart{w.Button{Label: "Offer again", Name: "do", Value: "again", Submit: true}},
									},
								},
								Action: "/app/duplicates/decide",
								Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "group", Value: "g9"}},
							},
						},
					},
				},
			},
		},
	}
}
