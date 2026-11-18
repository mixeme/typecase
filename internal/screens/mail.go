// Раздел макета: 7.18 Отправка почтой.

package screens

import w "carrel.dev/typecase/widget"

// Mail1 is frame «Письмо поверх текущего экрана» (s-mail-1).
func Mail1() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary, Focus: true},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.Overlay{
			Body: w.OverlayBody{
				Rail: []w.OverlayRailPart{
					w.Nav{
						Items: []w.NavItem{
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 7}, Href: "#"},
							{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
							{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
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
								Meta:    "10",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6"},
							w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "radicale · vps"},
							w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "2", Trail: w.Badge{Label: "ro"}},
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
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 2 selected")}}},
						},
						w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "АГ"},
									Title:    []w.Inline{w.Text("Алексей Гринёв")},
									Strong:   true,
									Subtitle: "+7 921 330-07-19",
									Detail:   "a.grinev@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
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
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "АЛ"},
									Title:    []w.Inline{w.Text("Антонина Лебедева")},
									Strong:   true,
									Subtitle: "+7 916 204-55-31",
									Detail:   "a.lebedeva@example.org",
									Caption:  "Проектное бюро",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Артём Белов")},
									Strong:   true,
									Subtitle: "+7 903 118-22-04",
									Detail:   "a.belov@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
									On:       true,
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Б", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Initials: "БЯ"},
									Title:    []w.Inline{w.Text("Белла Янковская")},
									Strong:   true,
									Subtitle: "+7 926 887-11-40",
									Detail:   "b.yankovskaya@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Богдан Ремизов")},
									Strong:   true,
									Subtitle: "+7 999 018-63-77",
									Detail:   "b.remizov@example.org",
									Caption:  "СМУ-14",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "БК"},
									Title:    []w.Inline{w.Text("Борис Кузьмин")},
									Strong:   true,
									Subtitle: "+7 905 771-42-08",
									Detail:   "b.kuzmin@example.org",
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
									Avatar:   w.Avatar{Initials: "ВД"},
									Title:    []w.Inline{w.Text("Вера Дорошенко")},
									Strong:   true,
									Subtitle: "+7 903 552-19-06",
									Detail:   "v.doroshenko@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ВП"},
									Title:    []w.Inline{w.Text("Виктор Панов")},
									Strong:   true,
									Subtitle: "+7 910 443-28-15",
									Detail:   "v.panov@example.org",
									Caption:  "Смета-Сервис",
									Source:   "nextcloud · Архив",
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Г", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "ГМ"},
									Title:    []w.Inline{w.Text("Галина Мещерякова")},
									Strong:   true,
									Subtitle: "+7 962 114-90-23",
									Detail:   "g.mescher@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ГС"},
									Title:    []w.Inline{w.Text("Геннадий Соловьёв")},
									Strong:   true,
									Subtitle: "+7 911 605-37-84",
									Detail:   "g.solovyov@example.org",
									Caption:  "Проектное бюро",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Глеб Тарасов")},
									Strong:   true,
									Subtitle: "+7 985 229-46-12",
									Detail:   "g.tarasov@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ГФ"},
									Title:    []w.Inline{w.Text("Григорий Фомин")},
									Strong:   true,
									Subtitle: "+7 917 783-51-60",
									Detail:   "g.fomin@example.org",
									Caption:  "СМУ-14",
									Source:   "nextcloud · Архив",
								},
							},
						},
					},
				},
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					ID:     "letter",
					Action: "/app/mail/send",
					Icon:   "mail",
					Title:  "Letter",
					Head: []w.SheetHeadPart{
						w.Tag{Label: "from mix@carrel.example", On: true, Place: w.Pushed},
						w.Close{Label: "Close", Href: "/app/contacts"},
					},
					Body: []w.SheetPart{
						w.Field{
							Label: "To",
							Parts: []w.FieldPart{
								w.Drop{
									Trigger: w.TagInput{
										Parts: []w.TagInputPart{
											w.Tag{
												Label:     "Анна Ковалёва",
												Mark:      &w.Mark{Colour: "#4A6B52", Size: w.MarkH11},
												Removable: true,
												Aside:     &w.Dim{Parts: []w.Inline{w.Text("· a.kovaleva@example.org")}},
												Title:     "a.kovaleva@example.org",
												Name:      "to",
												Value:     "a.kovaleva@example.org",
											},
											w.Tag{
												Label:     "Артём Белов",
												Mark:      &w.Mark{Colour: "#6B4E7A", Size: w.MarkH11},
												Removable: true,
												Aside:     &w.Dim{Parts: []w.Inline{w.Text("· a.belov@example.org")}},
												Title:     "a.belov@example.org",
												Name:      "to",
												Value:     "a.belov@example.org",
											},
										},
										Name: "to",
										Add: &w.TagAdd{
											Name:        "to.add",
											Placeholder: "Name or address",
											Label:       "Add a recipient",
											Suggest:     "/app/contacts/suggest",
											Enter:       &w.TagEnter{Name: "op", Value: "add", Label: "Add", Of: "letter"},
										},
										Focus: true,
									},
									Menu: w.Menu{Items: []w.MenuPart{}, Width: w.MenuWidth420, Pop: w.MenuUnder},
								},
								w.FieldHint{Parts: []w.Inline{w.Text("Up to 10 recipients. Each sees the others' addresses.")}},
							},
							Width: w.FieldW100,
						},
						w.Field{
							Label: "Subject",
							Parts: []w.FieldPart{w.Input{Name: "subject", Value: "Контакты подрядчика и смета", Label: "Subject"}},
							Width: w.FieldW100,
						},
						w.Area{
							Text:    "Анна, Артём — вот карточка подрядчика и смета, которую он прислал. Договор в папке, ссылка ниже.",
							Variant: w.AreaLetter,
							Name:    "text",
							Label:   "Text",
						},
						w.Sec{
							Items: []w.SecPart{
								w.Rubric{
									Text: "Attached",
									Items: []w.RubricPart{
										w.Num{Value: 4},
										w.Line{
											Items: []w.LinePart{
												w.Drop{
													Trigger: w.Button{Label: "Attach", Menu: true, Icon: "clip", Small: true},
													Menu: w.Menu{
														Items: []w.MenuPart{
															w.MenuItem{Icon: "con", Label: "Contact", To: "/app/mail/send?op=attach&what=contact"},
															w.MenuItem{Icon: "con", Label: "Address book", To: "/app/mail/send?op=attach&what=book"},
															w.MenuItem{Icon: "cal", Label: "Event", To: "/app/mail/send?op=attach&what=event"},
															w.MenuItem{Icon: "cal", Label: "Calendar", To: "/app/mail/send?op=attach&what=calendar"},
															w.MenuItem{Icon: "task", Label: "Task", To: "/app/mail/send?op=attach&what=task"},
															w.MenuItem{Icon: "note", Label: "Note", To: "/app/mail/send?op=attach&what=note"},
															w.MenuSep{},
															w.MenuItem{Icon: "doc", Label: "File", To: "/app/mail/send?op=attach&what=file"},
															w.MenuItem{
																Icon:  "file",
																Trail: w.Tag{Label: "link", Place: w.Pushed},
																Label: "Folder",
																To:    "/app/mail/send?op=attach&what=folder",
															},
														},
														Width: w.MenuWidth212,
														Pop:   w.MenuUnderEnd,
													},
												},
											},
											Gap:   w.LineGapNormal,
											Place: w.Pushed,
										},
									},
									Variant: w.RubricRow,
								},
								w.SecFiles{
									Items: []w.SecFile{
										{
											Icon:   "con",
											Name:   "Пётр Смирнов · подрядчик",
											Size:   w.Mono{Text: ".vcf · 3 KB", Variant: w.MonoDusty11},
											Tag:    w.Tag{Label: "attached"},
											Remove: w.Button{Icon: "close", Small: true, Title: "Remove", To: "/app/mail/send?op=remove&i=0", Submit: true},
										},
										{
											Icon:   "cal",
											Name:   "Встреча на объекте · 18 August 10:00",
											Size:   w.Mono{Text: ".ics · 2 KB", Variant: w.MonoDusty11},
											Tag:    w.Tag{Label: "attached"},
											Remove: w.Button{Icon: "close", Small: true, Title: "Remove", To: "/app/mail/send?op=remove&i=1", Submit: true},
										},
										{
											Icon:   "doc",
											Name:   "смета-подрядчик.pdf",
											Size:   w.Mono{Text: "4.2 MB", Variant: w.MonoDusty11},
											Tag:    w.Tag{Label: "as a link", On: true},
											Terms:  &w.Mono{Text: "30 days", Variant: w.MonoDusty11},
											Change: &w.Button{Label: "Change…", Small: true, To: "/app/mail/send?op=change&i=2", Submit: true},
											Remove: w.Button{Icon: "close", Small: true, Title: "Remove", To: "/app/mail/send?op=remove&i=2", Submit: true},
										},
										{
											Icon:   "file",
											Name:   "Договор подряда · 1 folder · 6 files",
											Size:   w.Mono{Text: "folder", Variant: w.MonoDusty11},
											Tag:    w.Tag{Label: "as a link", On: true},
											Terms:  &w.Mono{Text: "30 days", Variant: w.MonoDusty11},
											Change: &w.Button{Label: "Change…", Small: true, To: "/app/mail/send?op=change&i=3", Submit: true},
											Remove: w.Button{Icon: "close", Small: true, Title: "Remove", To: "/app/mail/send?op=remove&i=3", Submit: true},
										},
									},
								},
								w.Msg{
									Parts: []w.MsgPart{
										w.Text("The PDF (over 2 MB) and the folder go as links; revoking one in "),
										w.Strong("Published links"),
										w.Text(" closes it for this letter too."),
									},
									Tone:  w.MsgAccent,
									Space: w.MsgFlushAboveWide,
								},
							},
							Flush: true,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Send", Icon: "mail", Kind: w.ButtonPrimary, Busy: "Sending…", Submit: true},
						w.Kbd{Keys: "Ctrl+Enter"},
						w.Button{Label: "Preview", To: "/app/mail/send?op=preview", Submit: true},
						w.Dim{
							Parts:   []w.Inline{w.Text("Signed "), w.Strong("Михаил Енученко"), w.Text(" · replies go to mix@example.org")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", To: "/app/mail/send?op=close", Submit: true}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGapNormal,
							Place: w.Pushed,
						},
					},
					Width: w.SheetWidth680,
				},
			},
		},
	}
}

// Mail16 is frame «Письмо ушло» (s-mail-16).
func Mail16() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary, Focus: true},
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
					On: 1,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedPart,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined address book",
							Meta:    "10",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "2", Trail: w.Badge{Label: "ro"}},
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 2 selected")}}},
					},
					w.Msg{Parts: []w.MsgPart{w.Text("The letter went to 2 recipients.")}, Tone: w.MsgAccent},
					w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#4A6B52"},
								Avatar:   w.Avatar{Initials: "АГ"},
								Title:    []w.Inline{w.Text("Алексей Гринёв")},
								Strong:   true,
								Subtitle: "+7 921 330-07-19",
								Detail:   "a.grinev@example.org",
								Caption:  "—",
								Source:   "nextcloud · Контакты",
							},
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
								Mark:     w.Mark{Colour: "#A05A2C"},
								Avatar:   w.Avatar{Initials: "АЛ"},
								Title:    []w.Inline{w.Text("Антонина Лебедева")},
								Strong:   true,
								Subtitle: "+7 916 204-55-31",
								Detail:   "a.lebedeva@example.org",
								Caption:  "Проектное бюро",
								Source:   "nextcloud · Архив",
							},
							w.FaceRow{
								Mark:     w.Mark{Colour: "#6B4E7A"},
								Avatar:   w.Avatar{Photo: true},
								Title:    []w.Inline{w.Text("Артём Белов")},
								Strong:   true,
								Subtitle: "+7 903 118-22-04",
								Detail:   "a.belov@example.org",
								Caption:  "—",
								Source:   "radicale · vps · Команда",
								On:       true,
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Б", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#6B4E7A"},
								Avatar:   w.Avatar{Initials: "БЯ"},
								Title:    []w.Inline{w.Text("Белла Янковская")},
								Strong:   true,
								Subtitle: "+7 926 887-11-40",
								Detail:   "b.yankovskaya@example.org",
								Caption:  "—",
								Source:   "radicale · vps · Команда",
							},
							w.FaceRow{
								Mark:     w.Mark{Colour: "#A05A2C"},
								Avatar:   w.Avatar{Photo: true},
								Title:    []w.Inline{w.Text("Богдан Ремизов")},
								Strong:   true,
								Subtitle: "+7 999 018-63-77",
								Detail:   "b.remizov@example.org",
								Caption:  "СМУ-14",
								Source:   "nextcloud · Архив",
							},
							w.FaceRow{
								Mark:     w.Mark{Colour: "#4A6B52"},
								Avatar:   w.Avatar{Initials: "БК"},
								Title:    []w.Inline{w.Text("Борис Кузьмин")},
								Strong:   true,
								Subtitle: "+7 905 771-42-08",
								Detail:   "b.kuzmin@example.org",
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
								Avatar:   w.Avatar{Initials: "ВД"},
								Title:    []w.Inline{w.Text("Вера Дорошенко")},
								Strong:   true,
								Subtitle: "+7 903 552-19-06",
								Detail:   "v.doroshenko@example.org",
								Caption:  "—",
								Source:   "nextcloud · Контакты",
							},
							w.FaceRow{
								Mark:     w.Mark{Colour: "#A05A2C"},
								Avatar:   w.Avatar{Initials: "ВП"},
								Title:    []w.Inline{w.Text("Виктор Панов")},
								Strong:   true,
								Subtitle: "+7 910 443-28-15",
								Detail:   "v.panov@example.org",
								Caption:  "Смета-Сервис",
								Source:   "nextcloud · Архив",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Г", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#4A6B52"},
								Avatar:   w.Avatar{Initials: "ГМ"},
								Title:    []w.Inline{w.Text("Галина Мещерякова")},
								Strong:   true,
								Subtitle: "+7 962 114-90-23",
								Detail:   "g.mescher@example.org",
								Caption:  "—",
								Source:   "nextcloud · Контакты",
							},
							w.FaceRow{
								Mark:     w.Mark{Colour: "#A05A2C"},
								Avatar:   w.Avatar{Initials: "ГС"},
								Title:    []w.Inline{w.Text("Геннадий Соловьёв")},
								Strong:   true,
								Subtitle: "+7 911 605-37-84",
								Detail:   "g.solovyov@example.org",
								Caption:  "Проектное бюро",
								Source:   "nextcloud · Архив",
							},
							w.FaceRow{
								Mark:     w.Mark{Colour: "#6B4E7A"},
								Avatar:   w.Avatar{Photo: true},
								Title:    []w.Inline{w.Text("Глеб Тарасов")},
								Strong:   true,
								Subtitle: "+7 985 229-46-12",
								Detail:   "g.tarasov@example.org",
								Caption:  "—",
								Source:   "radicale · vps · Команда",
							},
							w.FaceRow{
								Mark:     w.Mark{Colour: "#A05A2C"},
								Avatar:   w.Avatar{Initials: "ГФ"},
								Title:    []w.Inline{w.Text("Григорий Фомин")},
								Strong:   true,
								Subtitle: "+7 917 783-51-60",
								Detail:   "g.fomin@example.org",
								Caption:  "СМУ-14",
								Source:   "nextcloud · Архив",
							},
						},
					},
				},
			},
		},
	}
}

// Mail17 is frame «Письмо не ушло» (s-mail-17).
func Mail17() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary, Focus: true},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.Overlay{
			Body: w.OverlayBody{
				Rail: []w.OverlayRailPart{
					w.Nav{
						Items: []w.NavItem{
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 7}, Href: "#"},
							{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
							{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
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
								Meta:    "10",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6"},
							w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "radicale · vps"},
							w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "2", Trail: w.Badge{Label: "ro"}},
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
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 2 selected")}}},
						},
						w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "АГ"},
									Title:    []w.Inline{w.Text("Алексей Гринёв")},
									Strong:   true,
									Subtitle: "+7 921 330-07-19",
									Detail:   "a.grinev@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
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
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "АЛ"},
									Title:    []w.Inline{w.Text("Антонина Лебедева")},
									Strong:   true,
									Subtitle: "+7 916 204-55-31",
									Detail:   "a.lebedeva@example.org",
									Caption:  "Проектное бюро",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Артём Белов")},
									Strong:   true,
									Subtitle: "+7 903 118-22-04",
									Detail:   "a.belov@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
									On:       true,
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Б", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Initials: "БЯ"},
									Title:    []w.Inline{w.Text("Белла Янковская")},
									Strong:   true,
									Subtitle: "+7 926 887-11-40",
									Detail:   "b.yankovskaya@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Богдан Ремизов")},
									Strong:   true,
									Subtitle: "+7 999 018-63-77",
									Detail:   "b.remizov@example.org",
									Caption:  "СМУ-14",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "БК"},
									Title:    []w.Inline{w.Text("Борис Кузьмин")},
									Strong:   true,
									Subtitle: "+7 905 771-42-08",
									Detail:   "b.kuzmin@example.org",
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
									Avatar:   w.Avatar{Initials: "ВД"},
									Title:    []w.Inline{w.Text("Вера Дорошенко")},
									Strong:   true,
									Subtitle: "+7 903 552-19-06",
									Detail:   "v.doroshenko@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ВП"},
									Title:    []w.Inline{w.Text("Виктор Панов")},
									Strong:   true,
									Subtitle: "+7 910 443-28-15",
									Detail:   "v.panov@example.org",
									Caption:  "Смета-Сервис",
									Source:   "nextcloud · Архив",
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Г", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "ГМ"},
									Title:    []w.Inline{w.Text("Галина Мещерякова")},
									Strong:   true,
									Subtitle: "+7 962 114-90-23",
									Detail:   "g.mescher@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ГС"},
									Title:    []w.Inline{w.Text("Геннадий Соловьёв")},
									Strong:   true,
									Subtitle: "+7 911 605-37-84",
									Detail:   "g.solovyov@example.org",
									Caption:  "Проектное бюро",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Глеб Тарасов")},
									Strong:   true,
									Subtitle: "+7 985 229-46-12",
									Detail:   "g.tarasov@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ГФ"},
									Title:    []w.Inline{w.Text("Григорий Фомин")},
									Strong:   true,
									Subtitle: "+7 917 783-51-60",
									Detail:   "g.fomin@example.org",
									Caption:  "СМУ-14",
									Source:   "nextcloud · Архив",
								},
							},
						},
					},
				},
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					ID:     "letter",
					Action: "/app/mail/send",
					Icon:   "mail",
					Title:  "Letter",
					Head: []w.SheetHeadPart{
						w.Tag{Label: "from mix@carrel.example", On: true, Place: w.Pushed},
						w.Close{Label: "Close", Href: "/app/contacts"},
					},
					Body: []w.SheetPart{
						w.Field{
							Label: "To",
							Parts: []w.FieldPart{
								w.Drop{
									Trigger: w.TagInput{
										Parts: []w.TagInputPart{
											w.Tag{
												Label:     "Анна Ковалёва",
												Mark:      &w.Mark{Colour: "#4A6B52", Size: w.MarkH11},
												Removable: true,
												Aside:     &w.Dim{Parts: []w.Inline{w.Text("· a.kovaleva@example.org")}},
												Title:     "a.kovaleva@example.org",
												Name:      "to",
												Value:     "a.kovaleva@example.org",
											},
											w.Tag{
												Label:     "Артём Белов",
												Mark:      &w.Mark{Colour: "#6B4E7A", Size: w.MarkH11},
												Removable: true,
												Aside:     &w.Dim{Parts: []w.Inline{w.Text("· a.belov@example.org")}},
												Title:     "a.belov@example.org",
												Name:      "to",
												Value:     "a.belov@example.org",
											},
										},
										Name: "to",
										Add: &w.TagAdd{
											Name:        "to.add",
											Placeholder: "Name or address",
											Label:       "Add a recipient",
											Suggest:     "/app/contacts/suggest",
											Enter:       &w.TagEnter{Name: "op", Value: "add", Label: "Add", Of: "letter"},
										},
										Focus: true,
									},
									Menu: w.Menu{Items: []w.MenuPart{}, Width: w.MenuWidth420, Pop: w.MenuUnder},
								},
								w.FieldHint{
									Parts: []w.Inline{
										w.Span{Face: w.FaceDanger, Text: "The letter did not go: the relay answered «451 4.7.0 Try again later»."},
										w.Text(" Up to 10 recipients. Each sees the others' addresses."),
									},
								},
							},
							Width: w.FieldW100,
						},
						w.Field{
							Label: "Subject",
							Parts: []w.FieldPart{w.Input{Name: "subject", Value: "Контакты подрядчика и смета", Label: "Subject"}},
							Width: w.FieldW100,
						},
						w.Area{
							Text:    "Анна, Артём — вот карточка подрядчика и смета, которую он прислал. Договор в папке, ссылка ниже.",
							Variant: w.AreaLetter,
							Name:    "text",
							Label:   "Text",
						},
						w.Sec{
							Items: []w.SecPart{
								w.Rubric{
									Text: "Attached",
									Items: []w.RubricPart{
										w.Num{Value: 4},
										w.Line{
											Items: []w.LinePart{
												w.Drop{
													Trigger: w.Button{Label: "Attach", Menu: true, Icon: "clip", Small: true},
													Menu: w.Menu{
														Items: []w.MenuPart{
															w.MenuItem{Icon: "con", Label: "Contact", To: "/app/mail/send?op=attach&what=contact"},
															w.MenuItem{Icon: "con", Label: "Address book", To: "/app/mail/send?op=attach&what=book"},
															w.MenuItem{Icon: "cal", Label: "Event", To: "/app/mail/send?op=attach&what=event"},
															w.MenuItem{Icon: "cal", Label: "Calendar", To: "/app/mail/send?op=attach&what=calendar"},
															w.MenuItem{Icon: "task", Label: "Task", To: "/app/mail/send?op=attach&what=task"},
															w.MenuItem{Icon: "note", Label: "Note", To: "/app/mail/send?op=attach&what=note"},
															w.MenuSep{},
															w.MenuItem{Icon: "doc", Label: "File", To: "/app/mail/send?op=attach&what=file"},
															w.MenuItem{
																Icon:  "file",
																Trail: w.Tag{Label: "link", Place: w.Pushed},
																Label: "Folder",
																To:    "/app/mail/send?op=attach&what=folder",
															},
														},
														Width: w.MenuWidth212,
														Pop:   w.MenuUnderEnd,
													},
												},
											},
											Gap:   w.LineGapNormal,
											Place: w.Pushed,
										},
									},
									Variant: w.RubricRow,
								},
								w.SecFiles{
									Items: []w.SecFile{
										{
											Icon:   "con",
											Name:   "Пётр Смирнов · подрядчик",
											Size:   w.Mono{Text: ".vcf · 3 KB", Variant: w.MonoDusty11},
											Tag:    w.Tag{Label: "attached"},
											Remove: w.Button{Icon: "close", Small: true, Title: "Remove", To: "/app/mail/send?op=remove&i=0", Submit: true},
										},
										{
											Icon:   "cal",
											Name:   "Встреча на объекте · 18 August 10:00",
											Size:   w.Mono{Text: ".ics · 2 KB", Variant: w.MonoDusty11},
											Tag:    w.Tag{Label: "attached"},
											Remove: w.Button{Icon: "close", Small: true, Title: "Remove", To: "/app/mail/send?op=remove&i=1", Submit: true},
										},
										{
											Icon:   "doc",
											Name:   "смета-подрядчик.pdf",
											Size:   w.Mono{Text: "4.2 MB", Variant: w.MonoDusty11},
											Tag:    w.Tag{Label: "as a link", On: true},
											Terms:  &w.Mono{Text: "30 days", Variant: w.MonoDusty11},
											Change: &w.Button{Label: "Change…", Small: true, To: "/app/mail/send?op=change&i=2", Submit: true},
											Remove: w.Button{Icon: "close", Small: true, Title: "Remove", To: "/app/mail/send?op=remove&i=2", Submit: true},
										},
										{
											Icon:   "file",
											Name:   "Договор подряда · 1 folder · 6 files",
											Size:   w.Mono{Text: "folder", Variant: w.MonoDusty11},
											Tag:    w.Tag{Label: "as a link", On: true},
											Terms:  &w.Mono{Text: "30 days", Variant: w.MonoDusty11},
											Change: &w.Button{Label: "Change…", Small: true, To: "/app/mail/send?op=change&i=3", Submit: true},
											Remove: w.Button{Icon: "close", Small: true, Title: "Remove", To: "/app/mail/send?op=remove&i=3", Submit: true},
										},
									},
								},
								w.Msg{
									Parts: []w.MsgPart{
										w.Text("The PDF (over 2 MB) and the folder go as links; revoking one in "),
										w.Strong("Published links"),
										w.Text(" closes it for this letter too."),
									},
									Tone:  w.MsgAccent,
									Space: w.MsgFlushAboveWide,
								},
							},
							Flush: true,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Send", Icon: "mail", Kind: w.ButtonPrimary, Busy: "Sending…", Submit: true},
						w.Kbd{Keys: "Ctrl+Enter"},
						w.Button{Label: "Preview", To: "/app/mail/send?op=preview", Submit: true},
						w.Dim{
							Parts:   []w.Inline{w.Text("Signed "), w.Strong("Михаил Енученко"), w.Text(" · replies go to mix@example.org")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", To: "/app/mail/send?op=close", Submit: true}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGapNormal,
							Place: w.Pushed,
						},
					},
					Width: w.SheetWidth680,
				},
			},
		},
	}
}

// Mail18 is frame «Письмо ушло не всем» (s-mail-18).
func Mail18() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary, Focus: true},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.Overlay{
			Body: w.OverlayBody{
				Rail: []w.OverlayRailPart{
					w.Nav{
						Items: []w.NavItem{
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 7}, Href: "#"},
							{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
							{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
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
								Meta:    "10",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6"},
							w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "radicale · vps"},
							w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "2", Trail: w.Badge{Label: "ro"}},
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
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 2 selected")}}},
						},
						w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "АГ"},
									Title:    []w.Inline{w.Text("Алексей Гринёв")},
									Strong:   true,
									Subtitle: "+7 921 330-07-19",
									Detail:   "a.grinev@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
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
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "АЛ"},
									Title:    []w.Inline{w.Text("Антонина Лебедева")},
									Strong:   true,
									Subtitle: "+7 916 204-55-31",
									Detail:   "a.lebedeva@example.org",
									Caption:  "Проектное бюро",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Артём Белов")},
									Strong:   true,
									Subtitle: "+7 903 118-22-04",
									Detail:   "a.belov@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
									On:       true,
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Б", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Initials: "БЯ"},
									Title:    []w.Inline{w.Text("Белла Янковская")},
									Strong:   true,
									Subtitle: "+7 926 887-11-40",
									Detail:   "b.yankovskaya@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Богдан Ремизов")},
									Strong:   true,
									Subtitle: "+7 999 018-63-77",
									Detail:   "b.remizov@example.org",
									Caption:  "СМУ-14",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "БК"},
									Title:    []w.Inline{w.Text("Борис Кузьмин")},
									Strong:   true,
									Subtitle: "+7 905 771-42-08",
									Detail:   "b.kuzmin@example.org",
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
									Avatar:   w.Avatar{Initials: "ВД"},
									Title:    []w.Inline{w.Text("Вера Дорошенко")},
									Strong:   true,
									Subtitle: "+7 903 552-19-06",
									Detail:   "v.doroshenko@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ВП"},
									Title:    []w.Inline{w.Text("Виктор Панов")},
									Strong:   true,
									Subtitle: "+7 910 443-28-15",
									Detail:   "v.panov@example.org",
									Caption:  "Смета-Сервис",
									Source:   "nextcloud · Архив",
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Г", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "ГМ"},
									Title:    []w.Inline{w.Text("Галина Мещерякова")},
									Strong:   true,
									Subtitle: "+7 962 114-90-23",
									Detail:   "g.mescher@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ГС"},
									Title:    []w.Inline{w.Text("Геннадий Соловьёв")},
									Strong:   true,
									Subtitle: "+7 911 605-37-84",
									Detail:   "g.solovyov@example.org",
									Caption:  "Проектное бюро",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Глеб Тарасов")},
									Strong:   true,
									Subtitle: "+7 985 229-46-12",
									Detail:   "g.tarasov@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ГФ"},
									Title:    []w.Inline{w.Text("Григорий Фомин")},
									Strong:   true,
									Subtitle: "+7 917 783-51-60",
									Detail:   "g.fomin@example.org",
									Caption:  "СМУ-14",
									Source:   "nextcloud · Архив",
								},
							},
						},
					},
				},
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Icon:  "mail",
					Title: "Letter",
					Head: []w.SheetHeadPart{
						w.Tag{Label: "from mix@carrel.example", On: true, Place: w.Pushed},
						w.Close{Label: "Close", Href: "/app/contacts"},
					},
					Body: []w.SheetPart{
						w.Msg{
							Parts: []w.MsgPart{w.Text("The letter went, but not to everyone. Sending again goes only to the refused.")},
							Tone:  w.MsgAlert,
							Space: w.MsgFlushSides,
						},
						w.Line{
							Items: []w.LinePart{
								w.Mark{Colour: "#4A6B52", Size: w.MarkH11},
								w.Text("Анна Ковалёва "),
								w.Dim{Parts: []w.Inline{w.Text("· a.kovaleva@example.org")}},
								w.Text(" — "),
								w.Strong("sent"),
							},
							Gap:   w.LineGapNormal,
							Block: true,
						},
						w.Line{
							Items: []w.LinePart{
								w.Mark{Colour: "#6B4E7A", Size: w.MarkH11},
								w.Text("Артём Белов "),
								w.Dim{Parts: []w.Inline{w.Text("· a.belov@example.org")}},
								w.Text(" — "),
								w.Faced{Face: w.FaceDanger, Items: []w.LinePart{w.Text("refused: «550 5.1.1 Mailbox unavailable»")}},
							},
							Gap:   w.LineGapNormal,
							Block: true,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Send again to the refused", Kind: w.ButtonPrimary, Href: "/app/mail/compose?to=refused"},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Close", Href: "/app/contacts"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGapNormal,
							Place: w.Pushed,
						},
					},
					Width: w.SheetWidth680,
				},
			},
		},
	}
}

// Mail4 is frame «Отправить ссылку письмом» (s-mail-4).
func Mail4() Screen {
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
						{Icon: "back", Label: "Back to app", Href: "#"}}},
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
								On: 7,
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
							Title: &w.HeadTitle{
								Text: "Published links",
								Badge: &w.Badge{
									Label: "local",
									Tone:  w.BadgeLocal,
									Title: "Local instance: the link will not open from outside your network",
								},
							},
							Subs: []w.HeadSub{
								w.HeadSub{
									Parts: []w.HeadSubPart{
										w.Text("Read-only, one object or one collection per link · "),
										w.Span{Face: w.FaceDanger, Text: "this instance is local: the link opens only from your network"},
									},
								},
							},
						},
						w.Table{
							Cols: []w.TableCol{
								{Label: "What is published"},
								{Label: "Kind", Rank: w.Rank1},
								{Label: "Created", Rank: w.Rank3, Secondary: true},
								{Label: "Expires", Rank: w.Rank4, Main: true},
								{Label: "Last opened", Secondary: true},
								{Label: "Opens"},
								{Width: w.Col120},
							},
							Rows: []w.TableRow{
								{
									Cells: []w.TableCell{
										{
											Parts: []w.CellPart{
												w.TableFile{
													Icon: "doc",
													Name: "смета-подрядчик.pdf",
													Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08/смета-подрядчик.pdf",
												},
											},
										},
										{Parts: []w.CellPart{w.Text("file")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-10")}, Mono: true},
										{Parts: []w.CellPart{w.Text("No expiry")}, Mono: true},
										{Parts: []w.CellPart{w.Text("never opened")}, Mono: true, Tone: w.CellDusty},
										{Parts: []w.CellPart{w.Text("0")}, Mono: true},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Copy", Small: true, Copy: "https://carrel.example/p/9c1f4e", Copied: "Copied"},
												w.Button{Label: "Send", Small: true, Secondary: true, Href: "/app/mail/link?id=9c1f4e"},
												w.Button{
													Label:     "Revoke",
													Kind:      w.ButtonDanger,
													Small:     true,
													Secondary: true,
													Form:      "post-revoke",
													Name:      "id",
													Value:     "9c1f4e",
												},
											},
											Drop: &w.Drop{
												Trigger: w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
												Row:     true,
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{Icon: "mail", Label: "Send", Href: "/app/mail/link?id=9c1f4e"},
														w.MenuSep{},
														w.MenuItem{Label: "Revoke", Tone: w.MenuToneDanger, Form: "post-revoke", Name: "id", Value: "9c1f4e"},
													},
													Width: w.MenuWidth212,
													Pop:   w.MenuUnderEnd,
												},
											},
											Acts: true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{
											Parts: []w.CellPart{
												w.TableFile{
													Mark: &w.Mark{Colour: "#4A6B52"},
													Name: "Личный · 10–31 August",
													Href: "/app/agenda/nextcloud/lichnyj?view=range&from=2026-08-10&to=2026-08-31",
												},
											},
										},
										{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-09")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-09-08")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-10 08:41")}, Mono: true},
										{Parts: []w.CellPart{w.Text("12")}, Mono: true},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Copy", Small: true, Copy: "https://carrel.example/p/7f3ac9/personal.ics", Copied: "Copied"},
												w.Button{Label: "Send", Small: true, Secondary: true, Href: "/app/mail/link?id=7f3ac9"},
												w.Button{
													Label:     "Revoke",
													Kind:      w.ButtonDanger,
													Small:     true,
													Secondary: true,
													Form:      "post-revoke",
													Name:      "id",
													Value:     "7f3ac9",
												},
											},
											Drop: &w.Drop{
												Trigger: w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
												Row:     true,
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{Icon: "mail", Label: "Send", Href: "/app/mail/link?id=7f3ac9"},
														w.MenuSep{},
														w.MenuItem{Label: "Revoke", Tone: w.MenuToneDanger, Form: "post-revoke", Name: "id", Value: "7f3ac9"},
													},
													Width: w.MenuWidth212,
													Pop:   w.MenuUnderEnd,
												},
											},
											Acts: true,
										},
									},
								},
							},
							Variant: w.TablePacked,
						},
						w.PostForm{
							ID:     "post-revoke",
							Action: "/app/publish?op=revoke",
							Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
						},
					},
				},
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/mail/send",
					Icon:   "mail",
					Title:  "Send this link",
					Head: []w.SheetHeadPart{
						w.SheetSubject{Mark: w.Mark{Colour: "#4A6B52"}, Text: "Личный · 10–31 August"},
						w.Close{Label: "Close", Href: "/app/files/published"},
					},
					Body: []w.SheetPart{
						w.Field{
							Label: "To",
							Parts: []w.FieldPart{w.Input{Name: "to", Value: "n.orlova@example.org", Label: "To", Focus: true}},
							Width: w.FieldW100,
						},
						w.Field{
							Label: "A line from you",
							Parts: []w.FieldPart{w.Input{Name: "line", Placeholder: "optional", Label: "A line from you"}},
							Width: w.FieldW100,
						},
						w.SheetFrom{
							Label: "Sent as",
							Lines: [][]w.Inline{
								[]w.Inline{w.Strong("Михаил Енученко"), w.Text(" "), w.Mono{Text: "<mix@carrel.example>", Variant: w.Mono11}},
								[]w.Inline{
									w.Text("Replies go to "),
									w.Mono{Text: "mix@example.org", Variant: w.Mono11},
									w.Text(" · "),
									w.Link{Text: "change in Account", Href: "#"},
								},
							},
						},
						w.Sec{
							Items: []w.SecPart{
								w.Rubric{Text: "The letter carries"},
								w.SecList{
									Items: [][]w.SecItemPart{
										[]w.SecItemPart{w.Text("The link and your line — nothing of the calendar itself")},
										[]w.SecItemPart{w.Text("Revoking the link closes it for this recipient too")},
										[]w.SecItemPart{w.Text("It expires on 8 September, as the link does")},
									},
									Gap:   w.StackGapTight,
									Muted: true,
								},
							},
							Flush: true,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Send", Icon: "mail", Kind: w.ButtonPrimary, Busy: "Sending…", Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Button{Label: "Copy link instead", To: "/app/mail/send?op=link", Submit: true},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/files/published"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGapNormal,
							Place: w.Pushed,
						},
					},
					Width: w.SheetWidth680,
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Mail5 is frame «Выбрать заметку» (s-mail-5).
func Mail5() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary, Focus: true},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.Overlay{
			Body: w.OverlayBody{
				Rail: []w.OverlayRailPart{
					w.Nav{
						Items: []w.NavItem{
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 7}, Href: "#"},
							{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
							{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
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
								Meta:    "10",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6"},
							w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "radicale · vps"},
							w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "2", Trail: w.Badge{Label: "ro"}},
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
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 2 selected")}}},
						},
						w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "АГ"},
									Title:    []w.Inline{w.Text("Алексей Гринёв")},
									Strong:   true,
									Subtitle: "+7 921 330-07-19",
									Detail:   "a.grinev@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
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
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "АЛ"},
									Title:    []w.Inline{w.Text("Антонина Лебедева")},
									Strong:   true,
									Subtitle: "+7 916 204-55-31",
									Detail:   "a.lebedeva@example.org",
									Caption:  "Проектное бюро",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Артём Белов")},
									Strong:   true,
									Subtitle: "+7 903 118-22-04",
									Detail:   "a.belov@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
									On:       true,
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Б", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Initials: "БЯ"},
									Title:    []w.Inline{w.Text("Белла Янковская")},
									Strong:   true,
									Subtitle: "+7 926 887-11-40",
									Detail:   "b.yankovskaya@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Богдан Ремизов")},
									Strong:   true,
									Subtitle: "+7 999 018-63-77",
									Detail:   "b.remizov@example.org",
									Caption:  "СМУ-14",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "БК"},
									Title:    []w.Inline{w.Text("Борис Кузьмин")},
									Strong:   true,
									Subtitle: "+7 905 771-42-08",
									Detail:   "b.kuzmin@example.org",
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
									Avatar:   w.Avatar{Initials: "ВД"},
									Title:    []w.Inline{w.Text("Вера Дорошенко")},
									Strong:   true,
									Subtitle: "+7 903 552-19-06",
									Detail:   "v.doroshenko@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ВП"},
									Title:    []w.Inline{w.Text("Виктор Панов")},
									Strong:   true,
									Subtitle: "+7 910 443-28-15",
									Detail:   "v.panov@example.org",
									Caption:  "Смета-Сервис",
									Source:   "nextcloud · Архив",
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Г", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "ГМ"},
									Title:    []w.Inline{w.Text("Галина Мещерякова")},
									Strong:   true,
									Subtitle: "+7 962 114-90-23",
									Detail:   "g.mescher@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ГС"},
									Title:    []w.Inline{w.Text("Геннадий Соловьёв")},
									Strong:   true,
									Subtitle: "+7 911 605-37-84",
									Detail:   "g.solovyov@example.org",
									Caption:  "Проектное бюро",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Глеб Тарасов")},
									Strong:   true,
									Subtitle: "+7 985 229-46-12",
									Detail:   "g.tarasov@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ГФ"},
									Title:    []w.Inline{w.Text("Григорий Фомин")},
									Strong:   true,
									Subtitle: "+7 917 783-51-60",
									Detail:   "g.fomin@example.org",
									Caption:  "СМУ-14",
									Source:   "nextcloud · Архив",
								},
							},
						},
						w.PostForm{
							ID:     "attach-search",
							Action: "/app/mail/attach",
							Fields: []w.Hidden{{Name: "what", Value: "note"}},
							Method: "get",
						},
					},
				},
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/mail/attach?what=note&uid=e41a",
					Icon:   "note",
					Title:  "Attach a note",
					Head: []w.SheetHeadPart{
						w.Filter{
							Query:   w.Query{Value: "подряд", Label: "Search", Name: "q", Of: "attach-search"},
							Quiet:   true,
							Variant: w.FilterT12W200,
							Place:   w.Pushed,
						},
						w.Close{Label: "Close", Href: "/app/mail/compose"},
					},
					Body: []w.SheetPart{
						w.SheetPick{
							Rail: []w.SheetPickRail{
								w.Tree{
									Rows: []w.TreeRow{
										{
											Mark:  &w.Mark{Kind: w.MarkMulti},
											Label: "All notebooks",
											Href:  "/app/mail/attach?what=note&q=подряд",
										},
										{Icon: "server", Label: "nextcloud"},
										{
											Mark:   &w.Mark{Colour: "#4A6B52"},
											Label:  "Записи",
											Indent: w.TreeLevel22,
											Href:   "/app/mail/attach?what=note&q=подряд&c=zapisi",
										},
										{
											Mark:   &w.Mark{Colour: "#3F6E8C"},
											Label:  "Рабочее",
											Indent: w.TreeLevel22,
											Href:   "/app/mail/attach?what=note&q=подряд&c=rabochee",
										},
									},
									On:      1,
									Variant: w.TreePadded,
								},
							},
							List: []w.SheetPickList{
								w.ListGroup{Title: w.Rubric{Text: "Found", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
								w.List{
									Rows: []w.ListRow{
										w.BriefRow{
											Mark:  w.Mark{Colour: "#3F6E8C"},
											Title: "Встреча с подрядчиком, итоги",
											Date:  w.Mono{Text: "18 Aug"},
											Where: "nextcloud · Рабочее",
											On:    true,
											Href:  "/app/mail/attach?what=note&q=подряд&uid=e41a",
										},
										w.BriefRow{
											Mark:  w.Mark{Colour: "#3F6E8C"},
											Title: "Подрядчики: кого смотрели",
											Date:  w.Mono{Text: "02 Aug"},
											Where: "nextcloud · Рабочее",
											Href:  "/app/mail/attach?what=note&q=подряд&uid=c27b",
										},
										w.BriefRow{
											Mark:  w.Mark{Colour: "#4A6B52"},
											Title: "Ремонт: что осталось",
											Date:  w.Mono{Text: "27 Jul"},
											Where: "nextcloud · Записи",
											Href:  "/app/mail/attach?what=note&q=подряд&uid=9a03",
										},
									},
								},
							},
							Short: true,
						},
						w.Sec{
							Items: []w.SecPart{
								w.Field{
									Label: "Send it as",
									Parts: []w.FieldPart{
										w.Seg{
											Options: []w.SegOption{{Label: "Markdown · .md", On: true, Value: "md"}, {Label: "Source · .ics", Value: "ics"}},
											Variant: w.SegStart,
											Name:    "format",
										},
										w.FieldHint{
											Parts: []w.Inline{
												w.Text("Markdown is for reading: the text and links to attachments, without the properties other clients wrote. The source "),
												w.Mono{Text: ".ics"},
												w.Text(" keeps everything and needs a client that reads VJOURNAL — jtx Board, Evolution."),
											},
										},
									},
									Width: w.FieldW100,
								},
							},
						},
					},
					Flush: true,
					Foot: []w.SheetFootPart{
						w.Button{Label: "Attach", Icon: "clip", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Dim{
							Parts:   []w.Inline{w.Text("Встреча с подрядчиком, итоги · "), w.Mono{Text: ".md"}, w.Text(" · 4 KB · goes as a file")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/mail/compose"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGapNormal,
							Place: w.Pushed,
						},
					},
					Width: w.SheetWidth900,
				},
			},
		},
	}
}

// Mail10 is frame «Календарь · за даты» (s-mail-10).
func Mail10() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary, Focus: true},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.Overlay{
			Body: w.OverlayBody{
				Rail: []w.OverlayRailPart{
					w.Nav{
						Items: []w.NavItem{
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 7}, Href: "#"},
							{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
							{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
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
								Meta:    "10",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6"},
							w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "radicale · vps"},
							w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "2", Trail: w.Badge{Label: "ro"}},
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
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 2 selected")}}},
						},
						w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "АГ"},
									Title:    []w.Inline{w.Text("Алексей Гринёв")},
									Strong:   true,
									Subtitle: "+7 921 330-07-19",
									Detail:   "a.grinev@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
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
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "АЛ"},
									Title:    []w.Inline{w.Text("Антонина Лебедева")},
									Strong:   true,
									Subtitle: "+7 916 204-55-31",
									Detail:   "a.lebedeva@example.org",
									Caption:  "Проектное бюро",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Артём Белов")},
									Strong:   true,
									Subtitle: "+7 903 118-22-04",
									Detail:   "a.belov@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
									On:       true,
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Б", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Initials: "БЯ"},
									Title:    []w.Inline{w.Text("Белла Янковская")},
									Strong:   true,
									Subtitle: "+7 926 887-11-40",
									Detail:   "b.yankovskaya@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Богдан Ремизов")},
									Strong:   true,
									Subtitle: "+7 999 018-63-77",
									Detail:   "b.remizov@example.org",
									Caption:  "СМУ-14",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "БК"},
									Title:    []w.Inline{w.Text("Борис Кузьмин")},
									Strong:   true,
									Subtitle: "+7 905 771-42-08",
									Detail:   "b.kuzmin@example.org",
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
									Avatar:   w.Avatar{Initials: "ВД"},
									Title:    []w.Inline{w.Text("Вера Дорошенко")},
									Strong:   true,
									Subtitle: "+7 903 552-19-06",
									Detail:   "v.doroshenko@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ВП"},
									Title:    []w.Inline{w.Text("Виктор Панов")},
									Strong:   true,
									Subtitle: "+7 910 443-28-15",
									Detail:   "v.panov@example.org",
									Caption:  "Смета-Сервис",
									Source:   "nextcloud · Архив",
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Г", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "ГМ"},
									Title:    []w.Inline{w.Text("Галина Мещерякова")},
									Strong:   true,
									Subtitle: "+7 962 114-90-23",
									Detail:   "g.mescher@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ГС"},
									Title:    []w.Inline{w.Text("Геннадий Соловьёв")},
									Strong:   true,
									Subtitle: "+7 911 605-37-84",
									Detail:   "g.solovyov@example.org",
									Caption:  "Проектное бюро",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Глеб Тарасов")},
									Strong:   true,
									Subtitle: "+7 985 229-46-12",
									Detail:   "g.tarasov@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ГФ"},
									Title:    []w.Inline{w.Text("Григорий Фомин")},
									Strong:   true,
									Subtitle: "+7 917 783-51-60",
									Detail:   "g.fomin@example.org",
									Caption:  "СМУ-14",
									Source:   "nextcloud · Архив",
								},
							},
						},
					},
				},
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Icon:  "cal",
					Title: "Attach a calendar",
					Head:  []w.SheetHeadPart{w.Close{Label: "Close", Place: w.Pushed, Href: "/app/mail/compose"}},
					Body: []w.SheetPart{
						w.SheetPick{
							Rail: []w.SheetPickRail{
								w.Tree{
									Rows: []w.TreeRow{
										{Icon: "server", Label: "nextcloud"},
										{
											Mark:   &w.Mark{Colour: "#4A6B52"},
											Label:  "Личный",
											Indent: w.TreeLevel22,
											Href:   "/app/mail/attach?what=calendar&c=lichnyj",
										},
										{
											Mark:   &w.Mark{Colour: "#3F6E8C"},
											Label:  "Работа",
											Indent: w.TreeLevel22,
											Href:   "/app/mail/attach?what=calendar&c=rabota",
										},
										{Icon: "server", Label: "radicale · vps"},
										{
											Mark:   &w.Mark{Colour: "#6B4E7A"},
											Label:  "Проекты",
											Indent: w.TreeLevel22,
											Href:   "/app/mail/attach?what=calendar&c=proekty",
										},
									},
									On:      2,
									Variant: w.TreePadded,
								},
							},
							List: []w.SheetPickList{
								w.Form{
									Items: []w.FormPart{
										w.Field{
											Label: "What goes in",
											Parts: []w.FieldPart{
												w.Seg{
													Options: []w.SegOption{
														{Label: "Whole calendar", Value: "whole", Disables: "from to"},
														{Label: "Dates", On: true, Value: "dates"},
													},
													Variant: w.SegStart,
													Name:    "span",
												},
											},
											Width: w.FieldW100,
										},
										w.Field{
											Label: "From",
											Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true}},
											Width: w.FieldSM,
										},
										w.Field{
											Label: "To",
											Parts: []w.FieldPart{
												w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-31", Label: "To", Mono: true},
												w.FieldHint{Parts: []w.Inline{w.Text("Both days are in the file.")}},
											},
											Width: w.FieldSM,
										},
									},
									ID:      "attach-calendar",
									Action:  "/app/mail/attach?what=calendar&c=lichnyj",
									Recount: "attach-line",
								},
							},
							Short: true,
						},
					},
					Flush: true,
					Foot: []w.SheetFootPart{
						w.Button{Label: "Attach", Icon: "clip", Kind: w.ButtonPrimary, Form: "attach-calendar"},
						w.Kbd{Keys: "Enter"},
						w.Dim{
							Parts: []w.Inline{
								w.Text("Calendar «Личный» · 10 August 2026 – 31 August 2026 · "),
								w.Mono{Text: ".ics"},
								w.Text(" · 5 KB · goes as a file"),
							},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							ID:      "attach-line",
						},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/mail/compose"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGapNormal,
							Place: w.Pushed,
						},
					},
					Width: w.SheetWidth900,
				},
			},
		},
	}
}

// Mail6 is frame «Выбрать файл» (s-mail-6).
func Mail6() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary, Focus: true},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.Overlay{
			Body: w.OverlayBody{
				Rail: []w.OverlayRailPart{
					w.Nav{
						Items: []w.NavItem{
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 7}, Href: "#"},
							{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
							{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
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
								Meta:    "10",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6"},
							w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "radicale · vps"},
							w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "2", Trail: w.Badge{Label: "ro"}},
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
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 2 selected")}}},
						},
						w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "АГ"},
									Title:    []w.Inline{w.Text("Алексей Гринёв")},
									Strong:   true,
									Subtitle: "+7 921 330-07-19",
									Detail:   "a.grinev@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
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
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "АЛ"},
									Title:    []w.Inline{w.Text("Антонина Лебедева")},
									Strong:   true,
									Subtitle: "+7 916 204-55-31",
									Detail:   "a.lebedeva@example.org",
									Caption:  "Проектное бюро",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Артём Белов")},
									Strong:   true,
									Subtitle: "+7 903 118-22-04",
									Detail:   "a.belov@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
									On:       true,
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Б", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Initials: "БЯ"},
									Title:    []w.Inline{w.Text("Белла Янковская")},
									Strong:   true,
									Subtitle: "+7 926 887-11-40",
									Detail:   "b.yankovskaya@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Богдан Ремизов")},
									Strong:   true,
									Subtitle: "+7 999 018-63-77",
									Detail:   "b.remizov@example.org",
									Caption:  "СМУ-14",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "БК"},
									Title:    []w.Inline{w.Text("Борис Кузьмин")},
									Strong:   true,
									Subtitle: "+7 905 771-42-08",
									Detail:   "b.kuzmin@example.org",
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
									Avatar:   w.Avatar{Initials: "ВД"},
									Title:    []w.Inline{w.Text("Вера Дорошенко")},
									Strong:   true,
									Subtitle: "+7 903 552-19-06",
									Detail:   "v.doroshenko@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ВП"},
									Title:    []w.Inline{w.Text("Виктор Панов")},
									Strong:   true,
									Subtitle: "+7 910 443-28-15",
									Detail:   "v.panov@example.org",
									Caption:  "Смета-Сервис",
									Source:   "nextcloud · Архив",
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Г", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "ГМ"},
									Title:    []w.Inline{w.Text("Галина Мещерякова")},
									Strong:   true,
									Subtitle: "+7 962 114-90-23",
									Detail:   "g.mescher@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ГС"},
									Title:    []w.Inline{w.Text("Геннадий Соловьёв")},
									Strong:   true,
									Subtitle: "+7 911 605-37-84",
									Detail:   "g.solovyov@example.org",
									Caption:  "Проектное бюро",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Глеб Тарасов")},
									Strong:   true,
									Subtitle: "+7 985 229-46-12",
									Detail:   "g.tarasov@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ГФ"},
									Title:    []w.Inline{w.Text("Григорий Фомин")},
									Strong:   true,
									Subtitle: "+7 917 783-51-60",
									Detail:   "g.fomin@example.org",
									Caption:  "СМУ-14",
									Source:   "nextcloud · Архив",
								},
							},
						},
					},
				},
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/mail/attach?what=file&path=sftpgo/documents/stroyka/podryad/smeta-podryadchik.pdf",
					Icon:   "doc",
					Title:  "Attach a file",
					Head: []w.SheetHeadPart{
						w.Crumbs{Path: []w.Crumb{{Label: "Documents", Href: "#"}, {Label: "Стройка", Href: "#"}}, Current: "Подряд"},
						w.Close{Label: "Close", Href: "/app/mail/compose"},
					},
					Body: []w.SheetPart{
						w.SheetPick{
							Rail: []w.SheetPickRail{
								w.Tree{
									Rows: []w.TreeRow{
										{Icon: "server", Label: "sftpgo · vps", Href: "/app/mail/attach?what=file&dir=sftpgo"},
										{
											Icon:   "file",
											Label:  "Documents",
											Indent: w.TreeLevel22,
											Href:   "/app/mail/attach?what=file&dir=sftpgo/documents",
										},
										{
											Icon:   "file",
											Label:  "Стройка",
											Indent: w.TreeLevel36,
											Href:   "/app/mail/attach?what=file&dir=sftpgo/documents/stroyka",
										},
										{
											Icon:   "file",
											Label:  "Договор подряда",
											Indent: w.TreeLevel44,
											Href:   "/app/mail/attach?what=file&dir=sftpgo/documents/stroyka/dogovor-podryada",
										},
										{
											Icon:   "file",
											Label:  "Подряд",
											Indent: w.TreeLevel44,
											Href:   "/app/mail/attach?what=file&dir=sftpgo/documents/stroyka/podryad",
										},
										{
											Icon:   "file",
											Label:  "Личное",
											Indent: w.TreeLevel36,
											Href:   "/app/mail/attach?what=file&dir=sftpgo/documents/lichnoe",
										},
									},
									On:      5,
									Variant: w.TreePadded,
								},
							},
							List: []w.SheetPickList{
								w.Table{
									Cols: []w.TableCol{
										{Label: "Name"},
										{Label: "Size", Width: w.Col90},
										{Label: "Changed", Width: w.Col110, Rank: w.Rank3, Main: true},
										{Width: w.ColType, Rank: w.Rank4},
									},
									Rows: []w.TableRow{
										{
											Cells: []w.TableCell{
												{
													Parts: []w.CellPart{
														w.TableFile{
															Icon: "doc",
															Name: "смета-подрядчик.pdf",
															Href: "/app/mail/attach?what=file&path=sftpgo/documents/stroyka/podryad/smeta-podryadchik.pdf",
														},
													},
												},
												{Parts: []w.CellPart{w.Text("4.2 MB")}, Mono: true},
												{Parts: []w.CellPart{w.Text("12 Aug")}, Mono: true},
												{Parts: []w.CellPart{w.Tag{Label: "over 2 MB · as a link", On: true}}},
											},
											State: w.RowOn,
										},
										{
											Cells: []w.TableCell{
												{
													Parts: []w.CellPart{
														w.TableFile{
															Icon: "doc",
															Name: "договор-подряда.pdf",
															Href: "/app/mail/attach?what=file&path=sftpgo/documents/stroyka/podryad/dogovor-podryada.pdf",
														},
													},
												},
												{Parts: []w.CellPart{w.Text("840 KB")}, Mono: true},
												{Parts: []w.CellPart{w.Text("09 Aug")}, Mono: true},
												{Parts: []w.CellPart{w.Tag{Label: "as a file"}}},
											},
										},
										{
											Cells: []w.TableCell{
												{
													Parts: []w.CellPart{
														w.TableFile{
															Icon: "img",
															Name: "объект-фасад.jpg",
															Href: "/app/mail/attach?what=file&path=sftpgo/documents/stroyka/podryad/obekt-fasad.jpg",
														},
													},
												},
												{Parts: []w.CellPart{w.Text("1.6 MB")}, Mono: true},
												{Parts: []w.CellPart{w.Text("09 Aug")}, Mono: true},
												{Parts: []w.CellPart{w.Tag{Label: "as a file"}}},
											},
										},
									},
								},
							},
							Short: true,
						},
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("Over the attachment limit: the letter carries a link instead, for 30 days. Revoke it in "),
								w.Strong("Published links"),
								w.Text("."),
							},
							Tone:  w.MsgAccent,
							Space: w.MsgSheet,
						},
					},
					Flush: true,
					Foot: []w.SheetFootPart{
						w.Button{Label: "Attach", Icon: "clip", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Dim{
							Parts:   []w.Inline{w.Text("смета-подрядчик.pdf · 4.2 MB · goes as a link")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/mail/compose"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGapNormal,
							Place: w.Pushed,
						},
					},
					Width: w.SheetWidth900,
				},
			},
		},
	}
}

// Mail12 is frame «Выбрать контакт» (s-mail-12).
func Mail12() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary, Focus: true},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.Overlay{
			Body: w.OverlayBody{
				Rail: []w.OverlayRailPart{
					w.Nav{
						Items: []w.NavItem{
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 7}, Href: "#"},
							{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
							{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
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
								Meta:    "10",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6"},
							w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "radicale · vps"},
							w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "2", Trail: w.Badge{Label: "ro"}},
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
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 2 selected")}}},
						},
						w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "АГ"},
									Title:    []w.Inline{w.Text("Алексей Гринёв")},
									Strong:   true,
									Subtitle: "+7 921 330-07-19",
									Detail:   "a.grinev@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
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
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "АЛ"},
									Title:    []w.Inline{w.Text("Антонина Лебедева")},
									Strong:   true,
									Subtitle: "+7 916 204-55-31",
									Detail:   "a.lebedeva@example.org",
									Caption:  "Проектное бюро",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Артём Белов")},
									Strong:   true,
									Subtitle: "+7 903 118-22-04",
									Detail:   "a.belov@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
									On:       true,
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Б", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Initials: "БЯ"},
									Title:    []w.Inline{w.Text("Белла Янковская")},
									Strong:   true,
									Subtitle: "+7 926 887-11-40",
									Detail:   "b.yankovskaya@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Богдан Ремизов")},
									Strong:   true,
									Subtitle: "+7 999 018-63-77",
									Detail:   "b.remizov@example.org",
									Caption:  "СМУ-14",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "БК"},
									Title:    []w.Inline{w.Text("Борис Кузьмин")},
									Strong:   true,
									Subtitle: "+7 905 771-42-08",
									Detail:   "b.kuzmin@example.org",
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
									Avatar:   w.Avatar{Initials: "ВД"},
									Title:    []w.Inline{w.Text("Вера Дорошенко")},
									Strong:   true,
									Subtitle: "+7 903 552-19-06",
									Detail:   "v.doroshenko@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ВП"},
									Title:    []w.Inline{w.Text("Виктор Панов")},
									Strong:   true,
									Subtitle: "+7 910 443-28-15",
									Detail:   "v.panov@example.org",
									Caption:  "Смета-Сервис",
									Source:   "nextcloud · Архив",
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Г", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "ГМ"},
									Title:    []w.Inline{w.Text("Галина Мещерякова")},
									Strong:   true,
									Subtitle: "+7 962 114-90-23",
									Detail:   "g.mescher@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ГС"},
									Title:    []w.Inline{w.Text("Геннадий Соловьёв")},
									Strong:   true,
									Subtitle: "+7 911 605-37-84",
									Detail:   "g.solovyov@example.org",
									Caption:  "Проектное бюро",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Глеб Тарасов")},
									Strong:   true,
									Subtitle: "+7 985 229-46-12",
									Detail:   "g.tarasov@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ГФ"},
									Title:    []w.Inline{w.Text("Григорий Фомин")},
									Strong:   true,
									Subtitle: "+7 917 783-51-60",
									Detail:   "g.fomin@example.org",
									Caption:  "СМУ-14",
									Source:   "nextcloud · Архив",
								},
							},
						},
						w.PostForm{
							ID:     "attach-search",
							Action: "/app/mail/attach",
							Fields: []w.Hidden{{Name: "what", Value: "contact"}},
							Method: "get",
						},
					},
				},
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/mail/attach?what=contact&uid=p7c2",
					Icon:   "con",
					Title:  "Attach a contact",
					Head: []w.SheetHeadPart{
						w.Filter{
							Query:   w.Query{Value: "смир", Label: "Search", Name: "q", Of: "attach-search"},
							Quiet:   true,
							Variant: w.FilterT12W200,
							Place:   w.Pushed,
						},
						w.Close{Label: "Close", Href: "/app/mail/compose"},
					},
					Body: []w.SheetPart{
						w.SheetPick{
							Rail: []w.SheetPickRail{
								w.Tree{
									Rows: []w.TreeRow{
										{
											Mark:  &w.Mark{Kind: w.MarkMulti},
											Label: "All address books",
											Href:  "/app/mail/attach?what=contact&q=смир",
										},
										{Icon: "server", Label: "nextcloud"},
										{
											Mark:   &w.Mark{Colour: "#4A6B52"},
											Label:  "Личный",
											Indent: w.TreeLevel22,
											Href:   "/app/mail/attach?what=contact&q=смир&c=lichnyj",
										},
										{
											Mark:   &w.Mark{Colour: "#3F6E8C"},
											Label:  "Работа",
											Indent: w.TreeLevel22,
											Href:   "/app/mail/attach?what=contact&q=смир&c=rabota",
										},
										{Icon: "server", Label: "radicale · vps"},
										{
											Mark:   &w.Mark{Colour: "#6B4E7A"},
											Label:  "Команда",
											Indent: w.TreeLevel22,
											Href:   "/app/mail/attach?what=contact&q=смир&c=komanda",
										},
									},
									On:      1,
									Variant: w.TreePadded,
								},
							},
							List: []w.SheetPickList{
								w.ListGroup{Title: w.Rubric{Text: "Found", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
								w.List{
									Rows: []w.ListRow{
										w.FaceRow{
											Mark:     w.Mark{Colour: "#3F6E8C"},
											Avatar:   w.Avatar{Initials: "ПС"},
											Title:    []w.Inline{w.Text("Пётр Смирнов")},
											Strong:   true,
											Subtitle: "+7 916 330-41-27",
											Detail:   "p.smirnov@example.org",
											Caption:  "подрядчик",
											Source:   "nextcloud · Работа",
											On:       true,
											Href:     "/app/mail/attach?what=contact&q=смир&uid=p7c2",
										},
										w.FaceRow{
											Mark:     w.Mark{Colour: "#4A6B52"},
											Avatar:   w.Avatar{Initials: "ИС"},
											Title:    []w.Inline{w.Text("Ирина Смирнова")},
											Strong:   true,
											Subtitle: "+7 903 221-08-64",
											Detail:   "i.smirnova@example.org",
											Caption:  "—",
											Source:   "nextcloud · Личный",
											Href:     "/app/mail/attach?what=contact&q=смир&uid=s1e8",
										},
									},
								},
							},
							Short: true,
						},
					},
					Flush: true,
					Foot: []w.SheetFootPart{
						w.Button{Label: "Attach", Icon: "clip", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Dim{
							Parts:   []w.Inline{w.Text("Пётр Смирнов · "), w.Mono{Text: ".vcf"}, w.Text(" · 3 KB · goes as a file")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/mail/compose"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGapNormal,
							Place: w.Pushed,
						},
					},
					Width: w.SheetWidth900,
				},
			},
		},
	}
}

// Mail13 is frame «Выбрать папку» (s-mail-13).
func Mail13() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary, Focus: true},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.Overlay{
			Body: w.OverlayBody{
				Rail: []w.OverlayRailPart{
					w.Nav{
						Items: []w.NavItem{
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 7}, Href: "#"},
							{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
							{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
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
								Meta:    "10",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6"},
							w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "radicale · vps"},
							w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "2", Trail: w.Badge{Label: "ro"}},
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
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 2 selected")}}},
						},
						w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "АГ"},
									Title:    []w.Inline{w.Text("Алексей Гринёв")},
									Strong:   true,
									Subtitle: "+7 921 330-07-19",
									Detail:   "a.grinev@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
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
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "АЛ"},
									Title:    []w.Inline{w.Text("Антонина Лебедева")},
									Strong:   true,
									Subtitle: "+7 916 204-55-31",
									Detail:   "a.lebedeva@example.org",
									Caption:  "Проектное бюро",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Артём Белов")},
									Strong:   true,
									Subtitle: "+7 903 118-22-04",
									Detail:   "a.belov@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
									On:       true,
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Б", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Initials: "БЯ"},
									Title:    []w.Inline{w.Text("Белла Янковская")},
									Strong:   true,
									Subtitle: "+7 926 887-11-40",
									Detail:   "b.yankovskaya@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Богдан Ремизов")},
									Strong:   true,
									Subtitle: "+7 999 018-63-77",
									Detail:   "b.remizov@example.org",
									Caption:  "СМУ-14",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "БК"},
									Title:    []w.Inline{w.Text("Борис Кузьмин")},
									Strong:   true,
									Subtitle: "+7 905 771-42-08",
									Detail:   "b.kuzmin@example.org",
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
									Avatar:   w.Avatar{Initials: "ВД"},
									Title:    []w.Inline{w.Text("Вера Дорошенко")},
									Strong:   true,
									Subtitle: "+7 903 552-19-06",
									Detail:   "v.doroshenko@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ВП"},
									Title:    []w.Inline{w.Text("Виктор Панов")},
									Strong:   true,
									Subtitle: "+7 910 443-28-15",
									Detail:   "v.panov@example.org",
									Caption:  "Смета-Сервис",
									Source:   "nextcloud · Архив",
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Г", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "ГМ"},
									Title:    []w.Inline{w.Text("Галина Мещерякова")},
									Strong:   true,
									Subtitle: "+7 962 114-90-23",
									Detail:   "g.mescher@example.org",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ГС"},
									Title:    []w.Inline{w.Text("Геннадий Соловьёв")},
									Strong:   true,
									Subtitle: "+7 911 605-37-84",
									Detail:   "g.solovyov@example.org",
									Caption:  "Проектное бюро",
									Source:   "nextcloud · Архив",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Глеб Тарасов")},
									Strong:   true,
									Subtitle: "+7 985 229-46-12",
									Detail:   "g.tarasov@example.org",
									Caption:  "—",
									Source:   "radicale · vps · Команда",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ГФ"},
									Title:    []w.Inline{w.Text("Григорий Фомин")},
									Strong:   true,
									Subtitle: "+7 917 783-51-60",
									Detail:   "g.fomin@example.org",
									Caption:  "СМУ-14",
									Source:   "nextcloud · Архив",
								},
							},
						},
					},
				},
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/mail/attach?what=folder&dir=sftpgo/documents/stroyka/dogovor-podryada",
					Icon:   "file",
					Title:  "Attach a folder",
					Head: []w.SheetHeadPart{
						w.Crumbs{
							Path:    []w.Crumb{{Label: "Documents", Href: "#"}, {Label: "Стройка", Href: "#"}},
							Current: "Договор подряда",
						},
						w.Close{Label: "Close", Href: "/app/mail/compose"},
					},
					Body: []w.SheetPart{
						w.SheetPick{
							Rail: []w.SheetPickRail{
								w.Tree{
									Rows: []w.TreeRow{
										{Icon: "server", Label: "sftpgo · vps", Href: "/app/mail/attach?what=folder&dir=sftpgo"},
										{
											Icon:   "file",
											Label:  "Documents",
											Indent: w.TreeLevel22,
											Href:   "/app/mail/attach?what=folder&dir=sftpgo/documents",
										},
										{
											Icon:   "file",
											Label:  "Стройка",
											Indent: w.TreeLevel36,
											Href:   "/app/mail/attach?what=folder&dir=sftpgo/documents/stroyka",
										},
										{
											Icon:   "file",
											Label:  "Договор подряда",
											Indent: w.TreeLevel44,
											Href:   "/app/mail/attach?what=folder&dir=sftpgo/documents/stroyka/dogovor-podryada",
										},
										{
											Icon:   "file",
											Label:  "Сканы",
											Indent: w.TreeLevel58,
											Href:   "/app/mail/attach?what=folder&dir=sftpgo/documents/stroyka/dogovor-podryada/skany",
										},
										{
											Icon:   "file",
											Label:  "Подряд",
											Indent: w.TreeLevel44,
											Href:   "/app/mail/attach?what=folder&dir=sftpgo/documents/stroyka/podryad",
										},
										{
											Icon:   "file",
											Label:  "Личное",
											Indent: w.TreeLevel36,
											Href:   "/app/mail/attach?what=folder&dir=sftpgo/documents/lichnoe",
										},
									},
									On:      4,
									Variant: w.TreePadded,
								},
							},
							List: []w.SheetPickList{
								w.Table{
									Cols: []w.TableCol{
										{Label: "In it · 1 folder · 6 files"},
										{Label: "Size", Width: w.Col90},
										{Label: "Changed", Width: w.Col110, Rank: w.Rank3, Main: true},
									},
									Rows: []w.TableRow{
										{
											Cells: []w.TableCell{
												{
													Parts: []w.CellPart{
														w.TableFile{
															Icon: "file",
															Name: "Сканы",
															Href: "/app/mail/attach?what=folder&dir=sftpgo/documents/stroyka/dogovor-podryada/skany",
														},
													},
												},
												{Parts: []w.CellPart{w.Text("—")}, Mono: true},
												{Parts: []w.CellPart{w.Text("14 Aug")}, Mono: true},
											},
											State: w.RowFolder,
										},
										{
											Cells: []w.TableCell{
												{Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "договор-подряда.pdf"}}},
												{Parts: []w.CellPart{w.Text("840 KB")}, Mono: true},
												{Parts: []w.CellPart{w.Text("09 Aug")}, Mono: true},
											},
											State: w.RowGone,
										},
										{
											Cells: []w.TableCell{
												{Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "приложение-1-смета.pdf"}}},
												{Parts: []w.CellPart{w.Text("4.2 MB")}, Mono: true},
												{Parts: []w.CellPart{w.Text("12 Aug")}, Mono: true},
											},
											State: w.RowGone,
										},
										{
											Cells: []w.TableCell{
												{Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "приложение-2-график.pdf"}}},
												{Parts: []w.CellPart{w.Text("210 KB")}, Mono: true},
												{Parts: []w.CellPart{w.Text("09 Aug")}, Mono: true},
											},
											State: w.RowGone,
										},
										{
											Cells: []w.TableCell{
												{Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "акт-1.pdf"}}},
												{Parts: []w.CellPart{w.Text("180 KB")}, Mono: true},
												{Parts: []w.CellPart{w.Text("14 Aug")}, Mono: true},
											},
											State: w.RowGone,
										},
										{
											Cells: []w.TableCell{
												{Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "реквизиты.txt"}}},
												{Parts: []w.CellPart{w.Text("2 KB")}, Mono: true},
												{Parts: []w.CellPart{w.Text("09 Aug")}, Mono: true},
											},
											State: w.RowGone,
										},
										{
											Cells: []w.TableCell{
												{Parts: []w.CellPart{w.TableFile{Icon: "img", Name: "объект-1.jpg"}}},
												{Parts: []w.CellPart{w.Text("1.6 MB")}, Mono: true},
												{Parts: []w.CellPart{w.Text("09 Aug")}, Mono: true},
											},
											State: w.RowGone,
										},
									},
								},
							},
							Short: true,
						},
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("A folder goes as a link, for 30 days. Revoke it in "),
								w.Strong("Published links"),
								w.Text("."),
							},
							Tone:  w.MsgAccent,
							Space: w.MsgSheet,
						},
					},
					Flush: true,
					Foot: []w.SheetFootPart{
						w.Button{Label: "Attach", Icon: "clip", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Dim{
							Parts:   []w.Inline{w.Text("Договор подряда · 1 folder · 6 files · goes as a link")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/mail/compose"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGapNormal,
							Place: w.Pushed,
						},
					},
					Width: w.SheetWidth900,
				},
			},
		},
	}
}
