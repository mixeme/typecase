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
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
							{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
							{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
							{Icon: "file", Label: "Files", Href: "#"},
						},
						On: 1,
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "128"},
							w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "41"},
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
							Title: &w.HeadTitle{Text: "Контакты"},
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("312 записей · 2 отмечено")}}},
						},
						w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "АК"},
									Title:    []w.Inline{w.Text("Анна Ковалёва")},
									Strong:   true,
									Subtitle: "+7 912 445-10-88",
									Detail:   "a.kovaleva@example.org",
									Caption:  "Издательство «Слово»",
									Source:   "nextcloud · Личный",
									On:       true,
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
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "АГ"},
									Title:    []w.Inline{w.Text("Алексей Гринёв")},
									Strong:   true,
									Subtitle: "+7 921 330-07-19",
									Detail:   "a.grinev@example.org",
									Caption:  "—",
									Source:   "nextcloud · Личный",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#3F6E8C"},
									Avatar:   w.Avatar{Initials: "АЛ"},
									Title:    []w.Inline{w.Text("Антонина Лебедева")},
									Strong:   true,
									Subtitle: "+7 916 204-55-31",
									Detail:   "a.lebedeva@example.org",
									Caption:  "Проектное бюро",
									Source:   "nextcloud · Работа",
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Б", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "БК"},
									Title:    []w.Inline{w.Text("Борис Кузьмин")},
									Strong:   true,
									Subtitle: "+7 905 771-42-08",
									Detail:   "b.kuzmin@example.org",
									Caption:  "—",
									Source:   "nextcloud · Личный",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#3F6E8C"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Богдан Ремизов")},
									Strong:   true,
									Subtitle: "+7 999 018-63-77",
									Detail:   "b.remizov@example.org",
									Caption:  "СМУ-14",
									Source:   "nextcloud · Работа",
								},
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
									Source:   "nextcloud · Личный",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#8C6239"},
									Avatar:   w.Avatar{Initials: "ВП"},
									Title:    []w.Inline{w.Text("Виктор Панов")},
									Strong:   true,
									Subtitle: "+7 910 443-28-15",
									Detail:   "v.panov@example.org",
									Caption:  "Смета-Сервис",
									Source:   "baikal · Офис",
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
									Source:   "nextcloud · Личный",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#3F6E8C"},
									Avatar:   w.Avatar{Initials: "ГС"},
									Title:    []w.Inline{w.Text("Геннадий Соловьёв")},
									Strong:   true,
									Subtitle: "+7 911 605-37-84",
									Detail:   "g.solovyov@example.org",
									Caption:  "Проектное бюро",
									Source:   "nextcloud · Работа",
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
									Mark:     w.Mark{Colour: "#8C6239"},
									Avatar:   w.Avatar{Initials: "ГФ"},
									Title:    []w.Inline{w.Text("Григорий Фомин")},
									Strong:   true,
									Subtitle: "+7 917 783-51-60",
									Detail:   "g.fomin@example.org",
									Caption:  "СМУ-14",
									Source:   "baikal · Офис",
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
					Head:  []w.SheetHeadPart{w.Tag{Label: "from mix@carrel.example", On: true, Place: w.Pushed}, w.Close{Label: "Close"}},
					Body: []w.SheetPart{
						w.Field{
							Label: "To",
							Parts: []w.FieldPart{
								w.TagInput{
									Parts: []w.TagInputPart{
										w.Tag{Label: "Анна Ковалёва · a.kovaleva@example.org", Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11}},
										w.Tag{Label: "Артём Белов · a.belov@example.org", Mark: &w.Mark{Colour: "#6B4E7A", Size: w.MarkH11}},
									},
									Wrap:  true,
									Focus: true,
								},
								w.FieldHint{Parts: []w.Inline{w.Text("2 of 10 recipients. Everyone sees the whole list.")}},
							},
							Width: w.FieldW100,
						},
						w.Field{
							Label: "Subject",
							Parts: []w.FieldPart{w.Input{Value: "Контакты подрядчика и смета", Label: "Subject"}},
							Width: w.FieldW100,
						},
						w.Area{
							Text:    "Анна, Артём — вот карточка подрядчика и смета, которую он прислал. Договор в папке, ссылка ниже.",
							Variant: w.AreaLetter,
							Label:   "Text",
						},
						w.Sec{
							Items: []w.SecPart{
								w.Rubric{
									Text: "Attached",
									Items: []w.RubricPart{
										w.Num{Value: 4},
										w.Line{
											Items: []w.LinePart{w.Button{Label: "Attach", Menu: true, Small: true}},
											Gap:   w.LineGap6,
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
											Remove: w.Button{Icon: "close", Small: true, Title: "Close"},
										},
										{
											Icon:   "cal",
											Name:   "Встреча на объекте · 18 August 10:00",
											Size:   w.Mono{Text: ".ics · 2 KB", Variant: w.MonoDusty11},
											Tag:    w.Tag{Label: "attached"},
											Remove: w.Button{Icon: "close", Small: true, Title: "Close"},
										},
										{
											Icon:   "doc",
											Name:   "смета-подрядчик.pdf",
											Size:   w.Mono{Text: "4.2 MB", Variant: w.MonoDusty11},
											Tag:    w.Tag{Label: "as a link", On: true},
											Remove: w.Button{Icon: "close", Small: true, Title: "Close"},
										},
										{
											Icon:   "file",
											Name:   "Договор подряда · 9 файлов",
											Size:   w.Mono{Text: "folder", Variant: w.MonoDusty11},
											Tag:    w.Tag{Label: "as a link", On: true},
											Remove: w.Button{Icon: "close", Small: true, Title: "Close"},
										},
									},
								},
								w.Msg{
									Parts: []w.MsgPart{
										w.Text("Two of these go as links, not as files: the PDF is over the 2 MB ceiling, and a folder always goes as a link. Both links live 30 days and can be revoked in "),
										w.Strong("Published links"),
										w.Text(" — including for the people in this letter."),
									},
									Tone:  w.MsgAccent,
									Space: w.MsgFlushTop11,
								},
							},
							Flush: true,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Send", Icon: "mail", Kind: w.ButtonPrimary},
						w.Button{Label: "Preview"},
						w.Dim{
							Parts:   []w.Inline{w.Text("Signed "), w.Strong("Михаил Енученко"), w.Text(" · replies go to mix@example.org")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
						w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}},
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
									{Icon: "server", Label: "Connections", Count: &w.Num{Value: 3}, Href: "#"},
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
							Title: &w.HeadTitle{Text: "Published links"},
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 links · 1 expires in 3 days")}}},
						},
						w.Table{
							Cols: []w.TableCol{
								{Label: "What is published"},
								{Label: "Kind"},
								{Label: "Created", Rank: w.Rank3, Secondary: true},
								{Label: "Expires", Rank: w.Rank4},
								{Label: "Last opened", Secondary: true},
								{Label: "Opens"},
								{Width: w.Col120},
							},
							Rows: []w.TableRow{
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "смета-подрядчик.pdf"}}},
										{Parts: []w.CellPart{w.Text("file")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-10")}, Mono: true},
										{Parts: []w.CellPart{w.Text("No expiry")}, Mono: true},
										{Parts: []w.CellPart{w.Text("never opened")}, Mono: true, Tone: w.CellDusty},
										{Parts: []w.CellPart{w.Text("0")}, Mono: true},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Copy", Small: true},
												w.Button{Label: "Send", Small: true, Secondary: true},
												w.Button{Label: "Revoke", Kind: w.ButtonDanger, Small: true, Secondary: true},
											},
											More: &w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
											Acts: true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Личный · 10–31 August"}}},
										{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-09")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-09-08")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-10 08:41")}, Mono: true},
										{Parts: []w.CellPart{w.Text("12")}, Mono: true},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Copy", Small: true},
												w.Button{Label: "Send", Small: true, Secondary: true},
												w.Button{Label: "Revoke", Kind: w.ButtonDanger, Small: true, Secondary: true},
											},
											More: &w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
											Acts: true,
										},
									},
								},
							},
						},
					},
				},
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Icon:  "mail",
					Title: "Send this link",
					Head: []w.SheetHeadPart{
						w.SheetSubject{Mark: w.Mark{Colour: "#4A6B52"}, Text: "Личный · 10–31 August"},
						w.Close{Label: "Close"},
					},
					Body: []w.SheetPart{
						w.Field{
							Label: "To",
							Parts: []w.FieldPart{w.Input{Value: "n.orlova@example.org", Label: "To", Focus: true}},
							Width: w.FieldW100,
						},
						w.Field{
							Label: "A line from you",
							Parts: []w.FieldPart{w.Input{Placeholder: "optional", Label: "A line from you"}},
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
									w.Span{Face: w.FaceAccent, Text: "change in Account"},
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
									Gap:   5,
									Muted: true,
								},
							},
							Flush: true,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Send", Icon: "mail", Kind: w.ButtonPrimary},
						w.Button{Label: "Copy link instead"},
						w.Push{Items: []w.PushPart{w.Button{Label: "Cancel"}}},
					},
					Width: w.SheetWidthFull,
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}
