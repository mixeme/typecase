// Раздел макета: 7.16 Настройки · публикация по ссылке.

package screens

import w "carrel.dev/typecase/widget"

// Publish1 is frame «Публикация · выданные ссылки» (s-publish-1).
func Publish1() Screen {
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
						Title: &w.HeadTitle{
							Text:  "Published links",
							Badge: &w.Badge{Label: "local", Tone: w.BadgeLocal, Title: "Local instance: may be unreachable from outside"},
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
	}
}

// Publish2 is frame «Календарь · создать ссылку» (s-publish-2).
func Publish2() Screen {
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
						On: 2,
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.TickRow{
								Ticked: w.TickedPart,
								Kind:   w.TickRowAll,
								Mark:   w.Mark{Kind: w.MarkMulti},
								Label:  "Combined calendar",
								Meta:   "14",
							},
						},
					},
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
						w.Head{
							Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH20}, Text: "Личный"},
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · 10 – 16 August 2026")}}},
							Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Icon: "plus", Kind: w.ButtonPrimary}}},
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
								},
								w.TimeRow{
									Mark:   w.Mark{Colour: "#4A6B52"},
									Time:   "18:00",
									Title:  []w.Inline{w.Text("Забрать документы")},
									Detail: "МФЦ на Ленина",
									Source: "nextcloud · Личный",
								},
							},
						},
					},
				},
				Height: w.ShellHeight390,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Title: "Publish a calendar",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("nextcloud · Личный")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
					},
					Body: []w.SheetPart{
						w.Field{
							Label: "How much of it",
							Parts: []w.FieldPart{w.Seg{Options: []w.SegOption{{Label: "A date range", On: true}, {Label: "The whole calendar"}}}},
							Width: w.FieldW100,
						},
						w.Line{
							Items: []w.LinePart{
								w.Field{
									Label: "From",
									Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Value: "2026-08-10", Label: "From", Mono: true}},
									Width: w.FieldSM,
								},
								w.Field{
									Label: "To",
									Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Value: "2026-08-31", Label: "To", Mono: true}},
									Width: w.FieldSM,
								},
							},
							Gap:   w.LineGap10,
							Block: true,
						},
						w.Field{Label: "Link lives for", Parts: []w.FieldPart{w.Pick{Value: "30 days"}}, Width: w.FieldSM},
						w.SheetChecks{
							Items: []w.Check{
								w.Check{
									On:    true,
									Parts: []w.CheckPart{w.Text("Allow "), w.Mono{Text: ".ics"}, w.Text(" download and subscription")},
								},
								w.Check{Parts: []w.CheckPart{w.Text("Show event descriptions")}},
							},
						},
					},
					Foot: []w.SheetFootPart{w.Button{Label: "Create link", Kind: w.ButtonPrimary}, w.Button{Label: "Cancel"}},
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Publish13 is frame «Календарь · ссылка создана» (s-publish-13).
func Publish13() Screen {
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
						On: 2,
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.TickRow{
								Ticked: w.TickedPart,
								Kind:   w.TickRowAll,
								Mark:   w.Mark{Kind: w.MarkMulti},
								Label:  "Combined calendar",
								Meta:   "14",
							},
						},
					},
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
						w.Head{
							Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH20}, Text: "Личный"},
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · 10 – 16 August 2026")}}},
							Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Icon: "plus", Kind: w.ButtonPrimary}}},
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
								},
								w.TimeRow{
									Mark:   w.Mark{Colour: "#4A6B52"},
									Time:   "18:00",
									Title:  []w.Inline{w.Text("Забрать документы")},
									Detail: "МФЦ на Ленина",
									Source: "nextcloud · Личный",
								},
							},
						},
					},
				},
				Height: w.ShellHeight340,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Title: "Publish a calendar",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("nextcloud · Личный")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
					},
					Body: []w.SheetPart{
						w.Field{
							Label: "Link",
							Parts: []w.FieldPart{
								w.Input{Value: "https://carrel.example/p/7Qm2c8Xv…", Label: "Link", Mono: true, ReadOnly: true},
								w.FieldHint{Parts: []w.Inline{w.Text("Until 2026-09-09.")}},
							},
							Width: w.FieldW100,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Send by letter", Kind: w.ButtonPrimary},
						w.Button{Label: "Open the link"},
						w.Button{Label: "All published links"},
					},
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Publish3 is frame «Контакт и адресная книга» (s-publish-3).
func Publish3() Screen {
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
								Ticked:  w.TickedPart,
								Kind:    w.TickRowAll,
								Mark:    w.Mark{Kind: w.MarkMulti},
								Label:   "Combined address book",
								Meta:    "312",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "248"},
							w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "51"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "radicale · vps"},
							w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "13", Trail: w.Badge{Label: "ro"}},
						},
					},
				},
				Main: w.Main{
					Items: []w.MainPart{
						w.Head{
							Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined address book"},
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 address books · 312 cards")}}},
							Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "New contact", Icon: "plus", Kind: w.ButtonPrimary}}},
						},
						w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
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
									Source:   "nextcloud · Контакты",
									On:       true,
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#6B4E7A"},
									Avatar:   w.Avatar{Photo: true},
									Title:    []w.Inline{w.Text("Артём Белов")},
									Subtitle: "+7 903 118-22-04",
									Detail:   "a.belov@example.org",
									Caption:  "Тензор",
									Source:   "radicale · vps · Команда",
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
					},
				},
				Height: w.ShellHeight340,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Title: "Publish a contact",
					Head: []w.SheetHeadPart{
						w.Dim{Parts: []w.Inline{w.Text("Анна Ковалёва")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed},
					},
					Body: []w.SheetPart{
						w.Field{
							Label: "What the visitor sees",
							Parts: []w.FieldPart{
								w.Group{
									Items: []w.FieldPart{
										w.Check{On: true, Parts: []w.CheckPart{w.Text("Name and organisation")}},
										w.Check{Parts: []w.CheckPart{w.Text("Phone numbers")}},
										w.Check{Parts: []w.CheckPart{w.Text("Email addresses")}},
										w.Check{Parts: []w.CheckPart{w.Text("Photo")}},
										w.Check{Parts: []w.CheckPart{w.Text("Note")}},
									},
									Form: w.GroupColumnLoose,
								},
							},
							Width: w.FieldW100,
						},
						w.Check{Parts: []w.CheckPart{w.Text("Allow "), w.Mono{Text: ".vcf"}, w.Text(" download and subscription")}},
						w.Field{Label: "Link lives for", Parts: []w.FieldPart{w.Pick{Value: "7 days"}}, Width: w.FieldSM},
					},
					Foot: []w.SheetFootPart{w.Button{Label: "Create link", Kind: w.ButtonPrimary}, w.Button{Label: "Cancel"}},
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Publish4 is frame «Заметка» (s-publish-4).
func Publish4() Screen {
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
						On: 4,
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.TickRow{
								Ticked:  w.TickedPart,
								Kind:    w.TickRowAll,
								Mark:    w.Mark{Kind: w.MarkMulti},
								Label:   "Combined notebook",
								Meta:    "86",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "64"},
							w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "22"},
						},
					},
				},
				Main: w.Main{
					Items: []w.MainPart{
						w.Head{
							Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined notebook"},
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 of 3 notebooks · 86 notes")}}},
							Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary}}},
						},
						w.ListGroup{Title: w.Rubric{Text: "August 2026", Form: w.RubricInline}, Count: &w.Num{Value: 9}},
						w.List{
							Rows: []w.ListRow{
								w.ExcerptRow{
									Mark:    w.Mark{Colour: "#4A6B52"},
									Title:   []w.Inline{w.Text("Заметка со встречи 10.08")},
									Strong:  true,
									Excerpt: "Договорились заморозить ветку в среду; письмо пользователям пишет Анна…",
									Tags:    []w.RowTag{w.Tag{Label: "встречи"}},
									Date:    "10 Aug",
									Source:  "nextcloud · Записи",
									On:      true,
								},
								w.ExcerptRow{
									Mark:    w.Mark{Colour: "#4A6B52"},
									Title:   []w.Inline{w.Text("Что почитать про CRDT")},
									Excerpt: "Кляйн, глава 5; статья Шапиро 2011; демка automerge…",
									Tags:    []w.RowTag{w.Tag{Label: "книги"}},
									Date:    "8 Aug",
									Source:  "nextcloud · Записи",
								},
							},
						},
					},
				},
				Height: w.ShellHeight340,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Title: "Publish a note",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("Заметка со встречи 10.08")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
					},
					Body: []w.SheetPart{
						w.Field{
							Label:      "Attachments",
							LabelAside: w.Plain{Text: "2 files · 3.2 MB"},
							Parts: []w.FieldPart{
								w.Group{
									Items: []w.FieldPart{
										w.Check{Parts: []w.CheckPart{w.Text("Publish them too — each gets its own link")}},
										w.FieldHint{
											Parts: []w.Inline{
												w.Text("The links are created together with this one, expire with it and are revoked with it. Without this the page says how many files there are, but not their names: a file name gives away more than it looks."),
											},
										},
									},
									Form: w.GroupColumnLoose,
								},
							},
							Width: w.FieldW100,
						},
						w.Check{On: true, Parts: []w.CheckPart{w.Text("Allow "), w.Mono{Text: ".md"}, w.Text(" download")}},
						w.Field{Label: "Link lives for", Parts: []w.FieldPart{w.Pick{Value: "30 days"}}, Width: w.FieldSM},
					},
					Foot: []w.SheetFootPart{w.Button{Label: "Create link", Kind: w.ButtonPrimary}, w.Button{Label: "Cancel"}},
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Publish5 is frame «Файл и папка» (s-publish-5).
func Publish5() Screen {
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
							Title: &w.HeadTitle{Text: "attachments"},
							Crumbs: &w.Crumbs{
								Path:    []w.Crumb{{Label: "Documents", Href: "#"}, {Label: "carrel", Href: "#"}},
								Current: "attachments",
							},
							Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("sftpgo · vps · "), w.Badge{Label: "attachment folder"}}}},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "Upload", Menu: true, Icon: "up", Kind: w.ButtonPrimary},
									w.Button{Label: "New folder", Icon: "plus"},
								},
							},
						},
						w.Table{
							Cols: []w.TableCol{
								{Width: w.ColTick, Tick: ptr(w.TickedOff)},
								{Label: "Name"},
								{Label: "Size", Width: w.ColSize, Rank: w.Rank3},
								{Label: "Type", Width: w.ColType, Rank: w.Rank4},
								{Label: "Changed", Width: w.ColTime},
							},
							Rows: []w.TableRow{
								{
									Cells: []w.TableCell{
										{Pick: ptr(w.TickedOff)},
										{Parts: []w.CellPart{w.TableFile{Icon: "file", Name: "2026-07"}}},
										{Parts: []w.CellPart{w.Text("—")}, Mono: true},
										{Parts: []w.CellPart{w.Text("folder · 34 items")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-07-01 09:12")}, Mono: true},
									},
									State: w.RowFolder,
								},
								{
									Cells: []w.TableCell{
										{Pick: ptr(w.TickedOn)},
										{Parts: []w.CellPart{w.TableFile{Icon: "img", Name: "2026-08-10-sync-board.png"}}},
										{Parts: []w.CellPart{w.Text("412 KB")}, Mono: true},
										{Parts: []w.CellPart{w.Text("image/png")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-10 11:07")}, Mono: true},
									},
									State: w.RowPicked,
								},
								{
									Cells: []w.TableCell{
										{Pick: ptr(w.TickedOff)},
										{Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "release-checklist.md"}}},
										{Parts: []w.CellPart{w.Text("3 KB")}, Mono: true},
										{Parts: []w.CellPart{w.Text("text/markdown")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-09 18:44")}, Mono: true},
									},
								},
							},
						},
					},
				},
				Height: w.ShellHeight340,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Title: "Publish a folder",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("Материалы к релизу · 12 files")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
					},
					Body: []w.SheetPart{
						w.Field{Label: "Link lives for", Parts: []w.FieldPart{w.Pick{Value: "14 days"}}, Width: w.FieldSM},
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("The visitor gets the folder with everything inside it, read-only, and the same address mounts as a drive in Windows, macOS and Linux."),
							},
							Space: w.MsgFlush,
						},
					},
					Foot: []w.SheetFootPart{w.Button{Label: "Create link", Kind: w.ButtonPrimary}, w.Button{Label: "Cancel"}},
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Publish6 is frame «Публичная страница · календарь» (s-publish-6).
func Publish6() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Banner{
						Logo:  w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"},
						Badge: &w.Badge{Label: "read-only"},
						End: w.Line{
							Gap:   w.LineGap6,
							Place: w.Pushed,
							Items: []w.LinePart{w.Button{Label: "Download .ics", Icon: "down"}},
						},
						Wide: true,
					},
					w.Strip{
						Label: "Subscribe in your own client",
						Items: []w.StripPart{
							w.Input{Value: "webcal://carrel.example/p/7f3ac9/personal.ics", Mono: true, Variant: w.InputLineFill},
							w.Button{Label: "Copy", Small: true},
							w.Button{Label: "Open in calendar app", Small: true},
							w.Dim{Parts: []w.Inline{w.Text("read-only, updates on its own")}, Size: w.Dim11},
						},
					},
					w.Head{
						Title: &w.HeadTitle{Text: "Личный"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{w.Text("10 – 31 August 2026 · the page is read from the server on every visit")},
							},
						},
						Pad: w.HeadPadBottom8,
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "10:00",
								Title:  []w.Inline{w.Text("Синк-ап по релизу")},
								Detail: "Переговорная 2",
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "18:00",
								Title:  []w.Inline{w.Text("Забрать документы")},
								Detail: "МФЦ, Мясницкая 13",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Tuesday · 11 August", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "09:30",
								Title:  []w.Inline{w.Text("Ревью архитектуры")},
								Detail: "Переговорная 1",
							},
						},
					},
					w.MainFoot{Text: "This link works until 8 September 2026. The owner can revoke it at any moment.", Roomy: true},
				},
			},
		},
	}
}

// Publish7 is frame «Публичная страница · адресная книга» (s-publish-7).
func Publish7() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Banner{
						Logo:  w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"},
						Badge: &w.Badge{Label: "read-only"},
						End: w.Line{
							Gap:   w.LineGap6,
							Place: w.Pushed,
							Items: []w.LinePart{w.Button{Label: "Download .vcf", Icon: "down"}},
						},
						Wide: true,
					},
					w.Strip{
						Label: "Address of this book",
						Items: []w.StripPart{
							w.Input{Value: "https://carrel.example/p/2b91f4/team.vcf", Mono: true, Variant: w.InputLineFill},
							w.Button{Label: "Copy", Small: true},
							w.Dim{
								Parts: []w.Inline{w.Text("another Carrel can subscribe to it; other clients import it")},
								Size:  w.Dim11,
							},
						},
					},
					w.Head{
						Title: &w.HeadTitle{Text: "Команда"},
						Subs: []w.HeadSub{
							w.HeadSub{Parts: []w.HeadSubPart{w.Text("13 contacts · the list is read from the server on every visit")}},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Filter{
									Query:   w.Query{Value: "Search this list", Label: "Search this list"},
									Waiting: true,
									Variant: w.FilterH26W190,
								},
							},
						},
						Pad: w.HeadPadBottom8,
					},
					w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#6B4E7A"},
								Avatar:   w.Avatar{Initials: "АБ"},
								Title:    []w.Inline{w.Text("Артём Белов")},
								Strong:   true,
								Subtitle: "+7 903 118-22-04",
								Detail:   "a.belov@example.org",
								Caption:  "Тензор",
							},
							w.FaceRow{
								Mark:     w.Mark{Colour: "#6B4E7A"},
								Avatar:   w.Avatar{Initials: "АК"},
								Title:    []w.Inline{w.Text("Анна Ковалёва")},
								Subtitle: "+7 912 445-10-88",
								Detail:   "a.kovaleva@example.org",
								Caption:  "Издательство «Слово»",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "В", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#6B4E7A"},
								Avatar:   w.Avatar{Initials: "ВГ"},
								Title:    []w.Inline{w.Text("Виктор Гончаров")},
								Subtitle: "—",
								Detail:   "v.goncharov@example.org",
								Caption:  "Городская библиотека №4",
							},
						},
					},
					w.MainFoot{Text: "This link works until 9 September 2026. The owner can revoke it at any moment.", Roomy: true},
				},
			},
		},
	}
}

// Publish8 is frame «Публичная страница · заметка» (s-publish-8).
func Publish8() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Banner{
						Logo:  w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"},
						Badge: &w.Badge{Label: "read-only"},
						End: w.Line{
							Gap:   w.LineGap6,
							Place: w.Pushed,
							Items: []w.LinePart{
								w.Seg{Options: []w.SegOption{{Label: "Full width", On: true}, {Label: "Reading width"}}},
								w.Button{Label: "Download .md", Icon: "down"},
							},
						},
					},
					w.DocLayout{
						Items: []w.DocLayoutPart{
							w.Doc{
								Items: []w.DocPart{
									w.DocTitle{Text: "Заметка со встречи 10.08"},
									w.DocSub{Parts: []w.DocSubPart{w.Mono{Text: "10 August 2026, 11:05"}}},
									w.Prose{
										Parts: []w.Inline{
											w.Text("Договорились заморозить ветку в среду.\n\nПисьмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт."),
										},
										Variant: w.ProseLead,
									},
									w.DocTasks{
										Title: w.Rubric{Text: "Что делаем", Variant: w.RubricBottom8},
										Items: []w.DocTask{
											{Tick: w.Tick{Variant: w.TickNudge, Label: "проверить changelog"}, Text: "проверить changelog"},
											{
												Tick: w.Tick{Done: true, Variant: w.TickNudge, Label: "попросить у Артёма скриншоты"},
												Text: "попросить у Артёма скриншоты",
											},
										},
									},
									w.DocFiles{
										Title: w.Rubric{Text: "Attachments", Variant: w.RubricBottom7},
										Items: []w.DocFile{
											{Name: "2026-08-10-sync-board.png", Size: w.Dim{Parts: []w.Inline{w.Text("· 412 KB")}}},
											{Name: "release-checklist.md", Size: w.Dim{Parts: []w.Inline{w.Text("· 3 KB")}}},
										},
									},
								},
							},
						},
						View: w.DocViewFocus,
					},
					w.MainFoot{Text: "This link works until 19 August 2026."},
				},
			},
		},
	}
}

// Publish9 is frame «Публичная страница · файл» (s-publish-9).
func Publish9() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Banner{
						Logo:  w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"},
						Badge: &w.Badge{Label: "read-only"},
						End:   w.Push{Items: []w.PushPart{w.Button{Label: "Download", Icon: "down", Kind: w.ButtonPrimary}}},
					},
					w.Head{
						Title: &w.HeadTitle{Text: "2026-08-10-sync-board.png", Size: w.HeadTitle19},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mono{Text: "image/png · 412 KB · changed 2026-08-10"}}}},
						Pad:   w.HeadPadBottom6,
					},
					w.MainThumbs{Thumb: w.Thumb{Icon: "img", Form: w.ThumbWide}},
					w.MainFoot{Text: "This link works until 19 August 2026."},
				},
			},
		},
	}
}

// Publish10 is frame «Публичная страница · текст и Markdown» (s-publish-10).
func Publish10() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Banner{
						Logo:  w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"},
						Badge: &w.Badge{Label: "read-only"},
						End: w.Line{
							Gap:   w.LineGap6,
							Place: w.Pushed,
							Items: []w.LinePart{
								w.Seg{Options: []w.SegOption{{Label: "Full width", On: true}, {Label: "Reading width"}}},
								w.Button{Label: "Download", Icon: "down"},
							},
						},
					},
					w.Head{
						Title: &w.HeadTitle{Text: "release-checklist.md", Size: w.HeadTitle19},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mono{Text: "text/markdown · 3.1 KB · changed 2026-08-09"}}}},
						Pad:   w.HeadPadBottom6,
					},
					w.Doc{
						Items: []w.DocPart{
							w.DocHeading{Text: "Перед заморозкой"},
							w.DocBullets{
								Items: [][]w.Inline{
									[]w.Inline{w.Text("прогнать "), w.Mono{Text: "go test -race ./...", Variant: w.MonoChip}},
									[]w.Inline{w.Text("проверить changelog")},
									[]w.Inline{w.Text("письмо пользователям — черновик у Анны")},
								},
							},
							w.DocHeading{Text: "После"},
							w.DocBody{Text: "Тег ставит Виктор. Образ собирается вручную до появления CI."},
						},
						Text: true,
					},
					w.MainFoot{Text: "This link works until 19 August 2026 · the file is read from the server on every visit"},
				},
			},
		},
	}
}

// Publish11 is frame «Публичная страница · файл без предпросмотра» (s-publish-11).
func Publish11() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Banner{
						Logo:  w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"},
						Badge: &w.Badge{Label: "read-only"},
					},
					w.Stage{
						Thumb:    w.Thumb{Icon: "doc", Form: w.ThumbSquare},
						Title:    "макет-обложки.psd",
						Meta:     w.Mono{Text: "image/vnd.adobe.photoshop · 84.2 MB · changed 2026-08-04"},
						Download: w.Button{Label: "Download 84.2 MB", Icon: "down", Kind: w.ButtonPrimary},
						Note:     "This file cannot be shown in a browser, and Carrel does not try. It is streamed with its own type, transcoded nowhere and stored nowhere.",
					},
					w.MainFoot{Text: "This link works until 19 August 2026."},
				},
			},
		},
	}
}

// Publish12 is frame «Публичная страница · папка» (s-publish-12).
func Publish12() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Banner{
						Logo:  w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"},
						Badge: &w.Badge{Label: "read-only"},
						End: w.Line{
							Gap:   w.LineGap6,
							Place: w.Pushed,
							Items: []w.LinePart{w.Button{Label: "Download folder", Icon: "down"}},
						},
					},
					w.Head{
						Title:  &w.HeadTitle{Text: "Материалы к релизу", Size: w.HeadTitle19},
						Crumbs: &w.Crumbs{Current: "Материалы к релизу", Aside: &w.Dim{Parts: []w.Inline{w.Text("· 4 files · 6.6 MB")}}},
						Pad:    w.HeadPadBottom8,
					},
					w.Strip{
						Label: "Mount as a drive",
						Items: []w.StripPart{
							w.Input{Value: "https://carrel.example/p/5c40e2/dav/", Mono: true, Variant: w.InputLineFill},
							w.Button{Label: "Copy", Small: true},
							w.Button{Label: "How to mount", Small: true},
							w.Dim{Parts: []w.Inline{w.Text("read-only WebDAV, no password")}, Size: w.Dim11},
						},
					},
					w.Table{
						Cols: []w.TableCol{
							{Label: "Name"},
							{Label: "Size", Width: w.ColSize},
							{Label: "Type", Width: w.ColChanged, Rank: w.Rank3},
							{Label: "Changed", Width: w.ColTime, Rank: w.Rank4},
							{Width: w.Col90},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Icon: "file", Name: "скриншоты"}}},
									{Parts: []w.CellPart{w.Text("—")}, Mono: true},
									{Parts: []w.CellPart{w.Text("folder · 6 files")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-10 11:07")}, Mono: true},
									{Parts: []w.CellPart{w.Button{Label: "Open", Small: true}}, Acts: true},
								},
								State: w.RowFolder,
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "release-checklist.md"}}},
									{Parts: []w.CellPart{w.Text("3 KB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("text/markdown")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-09 18:44")}, Mono: true},
									{Parts: []w.CellPart{w.Button{Label: "Download", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Icon: "img", Name: "2026-08-04-whiteboard.jpg"}}},
									{Parts: []w.CellPart{w.Text("2.8 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/jpeg")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-04 17:03")}, Mono: true},
									{Parts: []w.CellPart{w.Button{Label: "Download", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "смета-подрядчик.pdf"}}},
									{Parts: []w.CellPart{w.Text("1.2 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("application/pdf")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-05 14:20")}, Mono: true},
									{Parts: []w.CellPart{w.Button{Label: "Download", Small: true}}, Acts: true},
								},
							},
						},
					},
					w.MainFoot{
						Text: "This link works until 1 September 2026 · the contents are read from the server on every visit",
					},
				},
			},
		},
	}
}
