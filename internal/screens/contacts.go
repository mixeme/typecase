// Раздел макета: 7.1 Контакты и человек целиком.

package screens

import w "carrel.dev/typecase/widget"

// Contacts1 is frame «Контакты · список и карточка» (s-contacts-1).
func Contacts1() Screen {
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
					On: 1,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedPart,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined address book",
							Meta:    "7",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "5"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "1", Trail: w.Badge{Label: "ro"}},
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 6 cards")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New contact", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter the loaded cards", Label: "Filter the loaded cards"},
								Waiting: true,
								Variant: w.FilterH26W220,
							},
							w.Seg{
								Options: []w.SegOption{
									{Label: "A–Z", Href: "/app/contacts?sort=az", On: true},
									{Label: "Recently changed", Href: "/app/contacts?sort=changed"},
								},
								Secondary: true,
							},
							w.Button{Label: "Columns", Icon: "cols"},
							w.Drop{
								Trigger: w.Button{Icon: "dots", Title: "More"},
								More:    true,
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuItem{Label: "A–Z", On: true, Href: "/app/contacts?sort=az"},
										w.MenuItem{Label: "Recently changed", Href: "/app/contacts?sort=changed"},
										w.MenuSep{},
										w.CheckForm{
											Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}},
											Action: "/app/contacts",
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
									Action: "/app/contacts",
									Apply:  w.Button{Label: "Apply", Small: true, Submit: true},
								},
								w.Mono{Text: "6"},
							},
							Secondary: true,
						},
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
								Mark:       w.Mark{Colour: "#6B4E7A"},
								Avatar:     w.Avatar{Photo: true},
								Title:      []w.Inline{w.Text("Артём Белов ")},
								TitleBadge: &w.Badge{Label: "linked ×2", Tone: w.BadgeLinked},
								Subtitle:   "+7 903 118-22-04",
								Detail:     "a.belov@example.org",
								Caption:    "Тензор",
								Source:     "radicale · vps · Команда",
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
							w.FaceRow{
								Mark:     w.Mark{Colour: "#A05A2C"},
								Avatar:   w.Avatar{Initials: "ВЛ"},
								Title:    []w.Inline{w.Text("Вера Лебедева")},
								Subtitle: "+7 905 331-09-42",
								Detail:   "vera@example.net",
								Caption:  "—",
								Source:   "nextcloud · Архив",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Л", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
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
							},
						},
					},
				},
			},
			Side: &w.Side{
				Items: []w.SidePart{
					w.Head{
						Photo: &w.HeadPhoto{Initials: "АК"},
						Title: &w.HeadTitle{Text: "Анна Ковалёва", Size: w.HeadTitle19},
						Tools: []w.HeadTool{
							w.Button{Label: "Edit", Small: true},
							w.Button{Label: "All records", Small: true},
							w.Button{Icon: "dots", Small: true, Title: "More"},
						},
						Subs:   []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}}},
						Acts:   &w.Acts{Items: []w.ActsPart{w.Close{Label: "Close the panel"}}},
						Middle: true,
					},
					w.Fields{
						Items: []w.FieldsItem{
							{Term: "Organisation", Value: []w.ValuePart{w.Text("Издательство «Слово»")}},
							{Term: "Title", Value: []w.ValuePart{w.Text("Выпускающий редактор")}},
							{Term: "Phone · mobile", Value: []w.ValuePart{w.Text("+7 912 445-10-88")}, Face: w.ValueMono},
							{Term: "Phone · work", Value: []w.ValuePart{w.Text("+7 495 120-45-00")}, Face: w.ValueMono},
							{Term: "Email · work", Value: []w.ValuePart{w.Text("a.kovaleva@example.org")}},
							{Term: "Birthday", Value: []w.ValuePart{w.Text("1987-03-14")}, Face: w.ValueMono},
							{
								Term:  "Tags",
								Value: []w.ValuePart{w.Tag{Label: "работа"}, w.Text(" "), w.Tag{Label: "издательство"}},
							},
							{Term: "Note", Value: []w.ValuePart{w.Text("Пишет по вторникам, отвечает в тот же день.")}},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Kept as the server sent them", Items: []w.RubricPart{w.Button{Label: "Show 3", Small: true}}},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Source"},
							w.Diag{Text: "UID   4f2b-9c31-a0e7\nRead  2 min ago"}},
					},
					w.Danger{Button: w.Button{Label: "Delete contact", Kind: w.ButtonDangerSolid}},
				},
			},
		},
	}
}

// Contacts2 is frame «Контакт · переместить в другую книгу» (s-contacts-2).
func Contacts2() Screen {
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
						On: 1,
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.TickRow{
								Ticked:  w.TickedPart,
								Kind:    w.TickRowAll,
								Mark:    w.Mark{Kind: w.MarkMulti},
								Label:   "Combined address book",
								Meta:    "7",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "5"},
							w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "radicale · vps"},
							w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "1", Trail: w.Badge{Label: "ro"}},
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
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 6 cards")}}},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "New contact", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
									w.Button{Label: "Import", Icon: "import"},
									w.Button{Label: "Export", Icon: "export"},
									w.Button{Label: "Print", Icon: "print"},
								},
							},
						},
						w.Bar{
							Items: []w.BarPart{
								w.Filter{
									Query:   w.Query{Placeholder: "Filter the loaded cards", Label: "Filter the loaded cards"},
									Waiting: true,
									Variant: w.FilterH26W220,
								},
								w.Seg{
									Options: []w.SegOption{
										{Label: "A–Z", Href: "/app/contacts?sort=az", On: true},
										{Label: "Recently changed", Href: "/app/contacts?sort=changed"},
									},
									Secondary: true,
								},
								w.Button{Label: "Columns", Icon: "cols"},
								w.Drop{
									Trigger: w.Button{Icon: "dots", Title: "More"},
									More:    true,
									Menu: w.Menu{
										Items: []w.MenuPart{
											w.MenuItem{Label: "A–Z", On: true, Href: "/app/contacts?sort=az"},
											w.MenuItem{Label: "Recently changed", Href: "/app/contacts?sort=changed"},
											w.MenuSep{},
											w.CheckForm{
												Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}},
												Action: "/app/contacts",
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
										Action: "/app/contacts",
										Apply:  w.Button{Label: "Apply", Small: true, Submit: true},
									},
									w.Mono{Text: "6"},
								},
								Secondary: true,
							},
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
									Mark:       w.Mark{Colour: "#6B4E7A"},
									Avatar:     w.Avatar{Photo: true},
									Title:      []w.Inline{w.Text("Артём Белов ")},
									TitleBadge: &w.Badge{Label: "linked ×2", Tone: w.BadgeLinked},
									Subtitle:   "+7 903 118-22-04",
									Detail:     "a.belov@example.org",
									Caption:    "Тензор",
									Source:     "radicale · vps · Команда",
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
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ВЛ"},
									Title:    []w.Inline{w.Text("Вера Лебедева")},
									Subtitle: "+7 905 331-09-42",
									Detail:   "vera@example.net",
									Caption:  "—",
									Source:   "nextcloud · Архив",
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Л", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
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
								},
							},
						},
					},
				},
				Height: w.ShellHeight560,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/move",
					Icon:   "move",
					Title:  "Move contact",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("Анна Ковалёва · from nextcloud · Контакты")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
					},
					Body: []w.SheetPart{
						w.Rubric{Text: "Write into"},
						w.Tree{
							Rows: []w.TreeRow{
								{
									Mark:        &w.Mark{Colour: "#4A6B52"},
									Label:       "Контакты",
									Unavailable: true,
									Send:        w.Send{Name: "to", Value: "nextcloud/kontakty", One: true},
									Trail:       w.Mono{Text: "already there", Variant: w.Mono10Dusty},
								},
								{
									Mark:  &w.Mark{Colour: "#A05A2C"},
									Label: "Архив",
									Send:  w.Send{Name: "to", Value: "nextcloud/arhiv", One: true},
									Trail: w.Mono{Text: "nextcloud", Variant: w.Mono10Dusty},
								},
								{
									Mark:        &w.Mark{Colour: "#6B4E7A"},
									Label:       "Команда",
									Unavailable: true,
									Send:        w.Send{Name: "to", Value: "radicale/komanda", One: true},
									Trail:       w.Badge{Label: "read-only"},
								},
								{
									Mark:  &w.Mark{Colour: "#8C6239"},
									Label: "Рассылки",
									Send:  w.Send{Name: "to", Value: "radicale/rassylki", One: true},
									Trail: w.Mono{Text: "radicale · vps", Variant: w.Mono10Dusty},
								},
							},
							On: 2,
						},
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("Inside one server this is "),
								w.Mono{Text: "MOVE", Form: w.MonoCode},
								w.Text(": instant, no traffic. To another address book the card is copied and then deleted — Carrel says so before it starts, not after."),
							},
							Space: w.MsgFlushTop2,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Move", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/contacts"}, w.Kbd{Keys: "Esc"}},
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

// Contacts3 is frame «Контакт · правка, фотография» (s-contacts-3).
func Contacts3() Screen {
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "5", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
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
						Title: &w.HeadTitle{Text: "Анна Ковалёва"},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Field{
									Parts: []w.FieldPart{
										w.Drop{
											Trigger: w.Pick{Value: "nextcloud · Контакты", Mark: &w.Mark{Colour: "#4A6B52"}},
											Menu: w.Menu{
												Items: []w.MenuPart{
													w.MenuItem{
														Mark:  &w.Mark{Colour: "#4A6B52"},
														Label: "nextcloud · Контакты",
														On:    true,
														Name:  "book",
														Value: "nextcloud/kontakty",
														Of:    "contact-edit",
													},
													w.MenuItem{
														Mark:  &w.Mark{Colour: "#A05A2C"},
														Label: "nextcloud · Архив",
														Name:  "book",
														Value: "nextcloud/arhiv",
														Of:    "contact-edit",
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
						Progress: "contact-edit",
						Parts:    []w.RunbarPart{w.Text("Saving · "), w.Mono{Text: "0 of 0 MB"}},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label:     "Photo",
								LabelSpan: true,
								Parts: []w.FieldPart{
									w.Photo{
										ID: "photo",
										Side: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Button{
														Label: "Replace",
														File:  &w.FileField{Form: "", Name: "photo", Accept: "image/*", Label: "Photo", Preview: "photo"},
													},
													w.Button{Label: "Crop"},
												},
												Gap: w.LineGap6,
											},
											w.Button{Label: "Remove", Kind: w.ButtonDanger},
											w.FieldHint{Parts: []w.Inline{w.Text("EXIF orientation applied, metadata stripped.")}, Variant: w.HintNarrow},
										},
									},
								},
								Width: w.FieldFixAuto,
							},
							w.FormCols{
								Items: []w.Field{
									w.Field{
										Label: "Display name",
										Parts: []w.FieldPart{w.Input{Value: "Анна Ковалёва", Label: "Display name", Focus: true}},
										Width: w.FieldXL,
									},
									w.Field{Label: "Given", Parts: []w.FieldPart{w.Input{Value: "Анна", Label: "Given"}}, Width: w.FieldSM},
									w.Field{Label: "Family", Parts: []w.FieldPart{w.Input{Value: "Ковалёва", Label: "Family"}}, Width: w.FieldSM},
									w.Field{
										Label: "Nickname",
										Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Nickname"}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Organisation",
										Parts: []w.FieldPart{w.Input{Value: "Издательство «Слово»", Label: "Organisation"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Title",
										Parts: []w.FieldPart{w.Input{Value: "Выпускающий редактор", Label: "Title"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Birthday",
										Parts: []w.FieldPart{
											w.Input{Value: "14.03.1987", Label: "Birthday", Mono: true, Numeric: true},
											w.FieldHint{Parts: []w.Inline{w.Text("Type it: 14.03.1987, or 14.03 without the year.")}},
										},
										Width: w.FieldSM,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Phone numbers"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Label", Parts: []w.FieldPart{w.Pick{Value: "mobile"}}, Width: w.FieldXS},
									w.Field{
										Label: "Number",
										Parts: []w.FieldPart{w.Input{Value: "+7 912 445-10-88", Label: "Number", Mono: true}},
										Width: w.FieldMD,
									},
									w.Button{Label: "Remove", Kind: w.ButtonDanger},
									w.Grow{},
									w.Field{Parts: []w.FieldPart{w.Pick{Value: "work", Label: "Label"}}, Width: w.FieldXS},
									w.Field{
										Parts: []w.FieldPart{w.Input{Value: "+7 495 120-45-00", Label: "Number", Mono: true}},
										Width: w.FieldMD,
									},
									w.Button{Label: "Remove", Kind: w.ButtonDanger},
									w.Grow{},
									w.Button{Label: "Add phone", Icon: "plus", Small: true},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Email addresses"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Label", Parts: []w.FieldPart{w.Pick{Value: "work"}}, Width: w.FieldXS},
									w.Field{
										Label: "Address",
										Parts: []w.FieldPart{w.Input{Value: "a.kovaleva@example.org", Label: "Address"}},
										Width: w.FieldLG,
									},
									w.Button{Label: "Remove", Kind: w.ButtonDanger},
									w.Grow{},
									w.Button{Label: "Add email", Icon: "plus", Small: true},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Addresses"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Label", Parts: []w.FieldPart{w.Pick{Value: "home"}}, Width: w.FieldXS},
									w.Field{
										Label: "Street",
										Parts: []w.FieldPart{w.Input{Value: "ул. Садовая, 12, кв. 5", Label: "Street"}},
										Width: w.FieldLG,
									},
									w.Field{Label: "City", Parts: []w.FieldPart{w.Input{Value: "Москва", Label: "City"}}, Width: w.FieldSM},
									w.Field{Label: "Region", Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Region"}}, Width: w.FieldSM},
									w.Field{
										Label: "Postcode",
										Parts: []w.FieldPart{w.Input{Value: "123001", Label: "Postcode", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{Label: "Country", Parts: []w.FieldPart{w.Input{Value: "Россия", Label: "Country"}}, Width: w.FieldSM},
									w.Button{Label: "Remove", Kind: w.ButtonDanger},
									w.Grow{},
									w.Button{Label: "Add address", Icon: "plus", Small: true},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Other"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Tags",
										Parts: []w.FieldPart{w.Input{Value: "работа, издательство", Label: "Tags"}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Note",
										Parts: []w.FieldPart{w.Area{Text: "Пишет по вторникам, отвечает в тот же день.", Label: "Text"}},
										Width: w.FieldFull,
									},
								},
							},
							w.Field{Parts: []w.FieldPart{w.Button{Label: "Show 7 empty fields"}}},
						},
					},
					w.FormFoot{
						Items: []w.FormFootPart{
							w.Dim{Parts: []w.Inline{w.Text("Properties Carrel does not know are kept unchanged.")}, Size: w.Dim11_5},
						},
					},
				},
			},
		},
	}
}

// Contacts4 is frame «Контакт · правка, нет фотографии» (s-contacts-4).
func Contacts4() Screen {
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "5", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
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
						Title: &w.HeadTitle{Text: "Анна Ковалёва"},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Field{
									Parts: []w.FieldPart{
										w.Drop{
											Trigger: w.Pick{Value: "nextcloud · Контакты", Mark: &w.Mark{Colour: "#4A6B52"}},
											Menu: w.Menu{
												Items: []w.MenuPart{
													w.MenuItem{
														Mark:  &w.Mark{Colour: "#4A6B52"},
														Label: "nextcloud · Контакты",
														On:    true,
														Name:  "book",
														Value: "nextcloud/kontakty",
														Of:    "contact-edit",
													},
													w.MenuItem{
														Mark:  &w.Mark{Colour: "#A05A2C"},
														Label: "nextcloud · Архив",
														Name:  "book",
														Value: "nextcloud/arhiv",
														Of:    "contact-edit",
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
						Progress: "contact-edit",
						Parts:    []w.RunbarPart{w.Text("Saving · "), w.Mono{Text: "0 of 0 MB"}},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label:     "Photo",
								LabelSpan: true,
								Parts: []w.FieldPart{
									w.Photo{
										Initials: "АК",
										ID:       "photo",
										Side: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Button{
														Label: "Add",
														File:  &w.FileField{Form: "", Name: "photo", Accept: "image/*", Label: "Photo", Preview: "photo"},
													},
												},
												Gap: w.LineGap6,
											},
											w.FieldHint{Parts: []w.Inline{w.Text("JPEG, PNG, GIF or WebP.")}, Variant: w.HintNarrow},
										},
									},
								},
								Width: w.FieldFixAuto,
							},
							w.FormCols{
								Items: []w.Field{
									w.Field{
										Label: "Display name",
										Parts: []w.FieldPart{w.Input{Value: "Анна Ковалёва", Label: "Display name", Focus: true}},
										Width: w.FieldXL,
									},
									w.Field{Label: "Given", Parts: []w.FieldPart{w.Input{Value: "Анна", Label: "Given"}}, Width: w.FieldSM},
									w.Field{Label: "Family", Parts: []w.FieldPart{w.Input{Value: "Ковалёва", Label: "Family"}}, Width: w.FieldSM},
									w.Field{
										Label: "Nickname",
										Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Nickname"}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Organisation",
										Parts: []w.FieldPart{w.Input{Value: "Издательство «Слово»", Label: "Organisation"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Title",
										Parts: []w.FieldPart{w.Input{Value: "Выпускающий редактор", Label: "Title"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Birthday",
										Parts: []w.FieldPart{
											w.Input{Value: "14.03.1987", Label: "Birthday", Mono: true, Numeric: true},
											w.FieldHint{Parts: []w.Inline{w.Text("Type it: 14.03.1987, or 14.03 without the year.")}},
										},
										Width: w.FieldSM,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Phone numbers"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Label", Parts: []w.FieldPart{w.Pick{Value: "mobile"}}, Width: w.FieldXS},
									w.Field{
										Label: "Number",
										Parts: []w.FieldPart{w.Input{Value: "+7 912 445-10-88", Label: "Number", Mono: true}},
										Width: w.FieldMD,
									},
									w.Button{Label: "Remove", Kind: w.ButtonDanger},
									w.Grow{},
									w.Field{Parts: []w.FieldPart{w.Pick{Value: "work", Label: "Label"}}, Width: w.FieldXS},
									w.Field{
										Parts: []w.FieldPart{w.Input{Value: "+7 495 120-45-00", Label: "Number", Mono: true}},
										Width: w.FieldMD,
									},
									w.Button{Label: "Remove", Kind: w.ButtonDanger},
									w.Grow{},
									w.Button{Label: "Add phone", Icon: "plus", Small: true},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Email addresses"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Label", Parts: []w.FieldPart{w.Pick{Value: "work"}}, Width: w.FieldXS},
									w.Field{
										Label: "Address",
										Parts: []w.FieldPart{w.Input{Value: "a.kovaleva@example.org", Label: "Address"}},
										Width: w.FieldLG,
									},
									w.Button{Label: "Remove", Kind: w.ButtonDanger},
									w.Grow{},
									w.Button{Label: "Add email", Icon: "plus", Small: true},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Addresses"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Label", Parts: []w.FieldPart{w.Pick{Value: "home"}}, Width: w.FieldXS},
									w.Field{
										Label: "Street",
										Parts: []w.FieldPart{w.Input{Value: "ул. Садовая, 12, кв. 5", Label: "Street"}},
										Width: w.FieldLG,
									},
									w.Field{Label: "City", Parts: []w.FieldPart{w.Input{Value: "Москва", Label: "City"}}, Width: w.FieldSM},
									w.Field{Label: "Region", Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Region"}}, Width: w.FieldSM},
									w.Field{
										Label: "Postcode",
										Parts: []w.FieldPart{w.Input{Value: "123001", Label: "Postcode", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{Label: "Country", Parts: []w.FieldPart{w.Input{Value: "Россия", Label: "Country"}}, Width: w.FieldSM},
									w.Button{Label: "Remove", Kind: w.ButtonDanger},
									w.Grow{},
									w.Button{Label: "Add address", Icon: "plus", Small: true},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Other"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Tags",
										Parts: []w.FieldPart{w.Input{Value: "работа, издательство", Label: "Tags"}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Note",
										Parts: []w.FieldPart{w.Area{Text: "Пишет по вторникам, отвечает в тот же день.", Label: "Text"}},
										Width: w.FieldFull,
									},
								},
							},
							w.Field{Parts: []w.FieldPart{w.Button{Label: "Show 7 empty fields"}}},
						},
					},
					w.FormFoot{
						Items: []w.FormFootPart{
							w.Dim{Parts: []w.Inline{w.Text("Properties Carrel does not know are kept unchanged.")}, Size: w.Dim11_5},
						},
					},
				},
			},
		},
	}
}

// Contacts27 is frame «Контакт · новый, полная форма» (s-contacts-27).
func Contacts27() Screen {
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "5", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
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
						Title: &w.HeadTitle{Text: "New contact"},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Field{
									Parts: []w.FieldPart{
										w.Drop{
											Trigger: w.Pick{Value: "nextcloud · Контакты", Mark: &w.Mark{Colour: "#4A6B52"}},
											Menu: w.Menu{
												Items: []w.MenuPart{
													w.MenuItem{
														Mark:  &w.Mark{Colour: "#4A6B52"},
														Label: "nextcloud · Контакты",
														On:    true,
														Name:  "book",
														Value: "nextcloud/kontakty",
														Of:    "contact-new",
													},
													w.MenuItem{
														Mark:  &w.Mark{Colour: "#A05A2C"},
														Label: "nextcloud · Архив",
														Name:  "book",
														Value: "nextcloud/arhiv",
														Of:    "contact-new",
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
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label:     "Photo",
								LabelSpan: true,
								Parts: []w.FieldPart{
									w.Photo{
										ID: "photo",
										Side: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Button{
														Label: "Add",
														File:  &w.FileField{Form: "", Name: "photo", Accept: "image/*", Label: "Photo", Preview: "photo"},
													},
												},
												Gap: w.LineGap6,
											},
											w.FieldHint{Parts: []w.Inline{w.Text("JPEG, PNG, GIF or WebP.")}, Variant: w.HintNarrow},
										},
									},
								},
								Width: w.FieldFixAuto,
							},
							w.FormCols{
								Items: []w.Field{
									w.Field{
										Label: "Display name",
										Parts: []w.FieldPart{w.Input{Label: "Display name", Focus: true}},
										Width: w.FieldXL,
									},
									w.Field{Label: "Given", Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Given"}}, Width: w.FieldSM},
									w.Field{Label: "Family", Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Family"}}, Width: w.FieldSM},
									w.Field{
										Label: "Nickname",
										Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Nickname"}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Organisation",
										Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Organisation"}},
										Width: w.FieldMD,
									},
									w.Field{Label: "Title", Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Title"}}, Width: w.FieldMD},
									w.Field{
										Label: "Birthday",
										Parts: []w.FieldPart{
											w.Input{Placeholder: "DD.MM.YYYY", Label: "Birthday", Mono: true, Numeric: true},
											w.FieldHint{Parts: []w.Inline{w.Text("Type it: 14.03.1987, or 14.03 without the year.")}},
										},
										Width: w.FieldSM,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Phone numbers"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Label", Parts: []w.FieldPart{w.Pick{Value: "mobile"}}, Width: w.FieldXS},
									w.Field{
										Label: "Number",
										Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Number", Mono: true}},
										Width: w.FieldMD,
									},
									w.Button{Label: "Remove", Kind: w.ButtonDanger},
									w.Grow{},
									w.Button{Label: "Add phone", Icon: "plus", Small: true},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Email addresses"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Label", Parts: []w.FieldPart{w.Pick{Value: "work"}}, Width: w.FieldXS},
									w.Field{
										Label: "Address",
										Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Address"}},
										Width: w.FieldLG,
									},
									w.Button{Label: "Remove", Kind: w.ButtonDanger},
									w.Grow{},
									w.Button{Label: "Add email", Icon: "plus", Small: true},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Addresses"},
								Items:  []w.FieldSetPart{w.Button{Label: "Add address", Icon: "plus", Small: true}},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Websites"},
								Items:  []w.FieldSetPart{w.Button{Label: "Add website", Icon: "plus", Small: true}},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Dates"},
								Items:  []w.FieldSetPart{w.Button{Label: "Add date", Icon: "plus", Small: true}},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Other"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Tags", Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Tags"}}, Width: w.FieldLG},
									w.Field{Label: "Note", Parts: []w.FieldPart{w.Area{Text: "", Label: "Text"}}, Width: w.FieldFull},
								},
							},
							w.Field{Parts: []w.FieldPart{w.Button{Label: "Show 5 empty fields"}}},
						},
					},
				},
			},
		},
	}
}

// Contacts5 is frame «Фотография контакта · кадрирование» (s-contacts-5).
func Contacts5() Screen {
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "5", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
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
								Title: &w.HeadTitle{Text: "Photo"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Анна Ковалёва · metadata will be stripped")}}},
							},
							w.PhotoCrop{
								Colour: "#8C7B6A",
								Note:   "Drag the photo to choose the frame. The box is what other clients will see.",
								Fields: []w.Field{
									w.Field{Label: "Zoom", Parts: []w.FieldPart{w.Zoom{Value: "1.6×"}}, Width: w.FieldW100},
									w.Field{
										Label: "Rotate",
										Parts: []w.FieldPart{
											w.Seg{
												Options: []w.SegOption{{Label: "0°", On: true}, {Label: "90°"}, {Label: "180°"}, {Label: "270°"}},
												Variant: w.SegStart,
											},
										},
										Width: w.FieldW100,
									},
									w.Field{
										Label: "Saved as",
										Parts: []w.FieldPart{
											w.Input{Value: "512 × 512 · JPEG · about 90 KB", Label: "Saved as", Mono: true},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("The photo lives inside the card itself, so its size is the card’s size: a square of 512 px at quality 85 keeps a vCard near 200 KB, which servers and phones still accept. Lists are served a smaller copy of the same crop — the original file is not kept anywhere."),
												},
											},
										},
										Width: w.FieldW100,
									},
								},
								After: &w.Dim{
									Parts: []w.Inline{w.Text("EXIF orientation is already applied, metadata is stripped.")},
									Size:  w.Dim11,
								},
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Save photo", Kind: w.ButtonPrimary},
									w.Button{Label: "Cancel"},
									w.Push{Items: []w.PushPart{w.Button{Label: "Remove photo", Kind: w.ButtonDanger}}},
								},
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

// Contacts6 is frame «Контакт · карточка, чтение» (s-contacts-6).
func Contacts6() Screen {
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "5", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
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
						Photo: &w.HeadPhoto{},
						Title: &w.HeadTitle{Text: "Анна Ковалёва"},
						Subs: []w.HeadSub{
							w.HeadSub{Parts: []w.HeadSubPart{w.Text("Издательство «Слово» · Выпускающий редактор")}},
							w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Edit", Icon: "pencil", Kind: w.ButtonPrimary},
								w.Button{Label: "Send", Icon: "mail"},
								w.Button{Label: "Print", Icon: "print", Tertiary: true},
								w.Drop{
									Trigger: w.Button{Icon: "dots", Title: "More"},
									Menu: w.Menu{
										Items: []w.MenuPart{
											w.MenuItem{Icon: "print", Label: "Print", Tertiary: true},
											w.MenuItem{Icon: "move", Label: "Move"},
											w.MenuItem{Icon: "dup", Label: "Copy"},
											w.MenuItem{Icon: "export", Detail: &w.Mono{Text: ".vcf", Variant: w.Mono11}, Label: "Export"},
											w.MenuItem{Icon: "link", Label: "Publish a link"},
										},
										Width: w.MenuWidth212,
										Pop:   w.MenuUnderEnd,
									},
								},
							},
						},
					},
					w.Fields{
						Items: []w.FieldsItem{
							{
								Term:  "Phone · mobile",
								Value: []w.ValuePart{w.Text("+7 912 445-10-88")},
								Face:  w.ValueMono,
								Href:  "tel:+79124451088",
							},
							{
								Term:  "Phone · work",
								Value: []w.ValuePart{w.Text("+7 495 120-45-00")},
								Face:  w.ValueMono,
								Href:  "tel:+74951204500",
								After: []w.ValuePart{w.Badge{Label: "linked", Tone: w.BadgeLinked}},
							},
							{
								Term:  "Email · work",
								Value: []w.ValuePart{w.Text("a.kovaleva@example.org")},
								Href:  "mailto:a.kovaleva@example.org",
							},
							{
								Term:  "Email · second",
								Value: []w.ValuePart{w.Text("anna.kovaleva@slovo.example")},
								Href:  "mailto:anna.kovaleva@slovo.example",
								After: []w.ValuePart{w.Badge{Label: "linked", Tone: w.BadgeLinked}},
							},
							{Term: "Address · work", Value: []w.ValuePart{w.Text("Москва, Тверская 12, стр. 3")}},
							{
								Term:  "Birthday",
								Value: []w.ValuePart{w.Text("1987-03-14 "), w.Dim{Parts: []w.Inline{w.Text("· in 7 months")}}},
								Face:  w.ValueMono,
							},
							{
								Term:  "Tags",
								Value: []w.ValuePart{w.Tag{Label: "работа"}, w.Text(" "), w.Tag{Label: "издательство"}},
							},
							{Term: "Note", Value: []w.ValuePart{w.Text("Пишет по вторникам, отвечает в тот же день.")}},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Kept as the server sent them", Items: []w.RubricPart{w.Button{Label: "Show 3", Small: true}}},
							w.SecHint{Text: "Properties Carrel does not edit are stored byte for byte and go back untouched."},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{
								Text:  "The same person in other books",
								Items: []w.RubricPart{w.Badge{Label: "2", Tone: w.BadgeLinked}},
							},
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{
										w.Mark{Colour: "#6B4E7A", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "radicale · vps · Команда", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· Phone, email differ")}},
									},
									[]w.SecItemPart{
										w.Mark{Colour: "#9A9280", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "baikal · Общий", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· The server did not answer; the record stays linked")}},
									},
								},
								Gap: 5,
							},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Timeline", Items: []w.RubricPart{w.Mono{Text: "11", Variant: w.MonoTight}}},
							w.List{
								Rows: []w.ListRow{
									w.LinkedRow{
										Mark:   w.Mark{Colour: "#4A6B52"},
										Icon:   "cal",
										Title:  "Синк-ап по релизу",
										Detail: "10 Aug · attendee · accepted",
									},
									w.LinkedRow{
										Mark:   w.Mark{Colour: "#4A6B52"},
										Icon:   "file",
										Title:  "2026-08-04-обложка.pdf",
										Detail: "04 Aug · attached to «Согласовать обложку»",
									},
									w.LinkedRow{
										Mark:   w.Mark{Colour: "#3F6E8C"},
										Icon:   "task",
										Title:  "Согласовать обложку",
										Detail: "04 Aug · organiser · done",
									},
								},
							},
							w.SecActs{Items: []w.Button{w.Button{Label: "Open the timeline", Icon: "all"}}},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Source"},
							w.Diag{Text: "Path  /dav.php/addressbooks/mix/contacts/2b7e.vcf\nRead  2 min ago · not changed on server"},
						},
					},
					w.Danger{Button: w.Button{Label: "Delete contact", Icon: "trash", Kind: w.ButtonDangerSolid}},
				},
			},
		},
	}
}

// Contacts15 is frame «Контакт · адрес заменён — переписать участников?» (s-contacts-15).
func Contacts15() Screen {
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
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "5", Current: true},
							w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
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
							Title: &w.HeadTitle{Text: "Анна Ковалёва"},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Drop{
												Trigger: w.Pick{Value: "nextcloud · Контакты", Mark: &w.Mark{Colour: "#4A6B52"}},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#4A6B52"},
															Label: "nextcloud · Контакты",
															On:    true,
															Name:  "book",
															Value: "nextcloud/kontakty",
															Of:    "contact-edit",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#A05A2C"},
															Label: "nextcloud · Архив",
															Name:  "book",
															Value: "nextcloud/arhiv",
															Of:    "contact-edit",
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
									w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "contact-edit"},
								},
							},
						},
						w.Form{
							Items: []w.FormPart{
								w.Field{
									Label:     "Photo",
									LabelSpan: true,
									Parts: []w.FieldPart{
										w.Photo{
											ID: "photo-saved",
											Side: []w.FieldPart{
												w.Line{
													Items: []w.LinePart{
														w.Button{
															Label: "Replace",
															File:  &w.FileField{Form: "", Name: "photo", Accept: "image/*", Label: "Photo", Preview: "photo"},
														},
														w.Button{Label: "Crop"},
													},
													Gap: w.LineGap6,
												},
												w.Button{Label: "Remove", Kind: w.ButtonDanger},
												w.FieldHint{Parts: []w.Inline{w.Text("EXIF orientation applied, metadata stripped.")}, Variant: w.HintNarrow},
											},
										},
									},
									Width: w.FieldFixAuto,
								},
								w.FormCols{
									Items: []w.Field{
										w.Field{
											Label: "Display name",
											Parts: []w.FieldPart{w.Input{Name: "display.name", Value: "Анна Ковалёва", Label: "Display name", Focus: true}},
											Width: w.FieldXL,
										},
										w.Field{
											Label: "Given",
											Parts: []w.FieldPart{w.Input{Name: "given", Value: "Анна", Label: "Given"}},
											Width: w.FieldSM,
										},
										w.Field{
											Label: "Family",
											Parts: []w.FieldPart{w.Input{Name: "family", Value: "Ковалёва", Label: "Family"}},
											Width: w.FieldSM,
										},
										w.Field{
											Label: "Nickname",
											Parts: []w.FieldPart{w.Input{Name: "nickname", Placeholder: "—", Label: "Nickname"}},
											Width: w.FieldSM,
										},
										w.Field{
											Label: "Organisation",
											Parts: []w.FieldPart{w.Input{Name: "organisation", Value: "Издательство «Слово»", Label: "Organisation"}},
											Width: w.FieldMD,
										},
										w.Field{
											Label: "Title",
											Parts: []w.FieldPart{w.Input{Name: "title", Value: "Выпускающий редактор", Label: "Title"}},
											Width: w.FieldMD,
										},
										w.Field{
											Label: "Birthday",
											Parts: []w.FieldPart{
												w.Input{Name: "birthday", Value: "14.03.1987", Label: "Birthday", Mono: true, Numeric: true},
												w.FieldHint{Parts: []w.Inline{w.Text("Type it: 14.03.1987, or 14.03 without the year.")}},
											},
											Width: w.FieldSM,
										},
									},
								},
								w.FieldSet{
									Rubric: w.Rubric{Text: "Phone numbers"},
									Items: []w.FieldSetPart{
										w.Field{
											Label: "Label",
											Parts: []w.FieldPart{
												w.Pick{Name: "label", Options: []w.PickOption{{Value: "mobile", Label: "mobile", On: true}}, Label: "Label"},
											},
											Width: w.FieldXS,
										},
										w.Field{
											Label: "Number",
											Parts: []w.FieldPart{w.Input{Name: "number", Value: "+7 912 445-10-88", Label: "Number", Mono: true}},
											Width: w.FieldMD,
										},
										w.Button{Label: "Remove", Kind: w.ButtonDanger},
										w.Grow{},
										w.Field{Parts: []w.FieldPart{w.Input{Name: "label", Value: "work", Label: "Label"}}, Width: w.FieldXS},
										w.Field{
											Parts: []w.FieldPart{w.Input{Name: "number", Value: "+7 495 120-45-00", Label: "Number", Mono: true}},
											Width: w.FieldMD,
										},
										w.Button{Label: "Remove", Kind: w.ButtonDanger},
										w.Grow{},
										w.Button{Label: "Add phone", Icon: "plus", Small: true},
									},
								},
								w.FieldSet{
									Rubric: w.Rubric{Text: "Email addresses"},
									Items: []w.FieldSetPart{
										w.Field{
											Label: "Label",
											Parts: []w.FieldPart{
												w.Pick{Name: "label", Options: []w.PickOption{{Value: "work", Label: "work", On: true}}, Label: "Label"},
											},
											Width: w.FieldXS,
										},
										w.Field{
											Label: "Address",
											Parts: []w.FieldPart{w.Input{Name: "address", Value: "a.kovaleva@slovo.example", Label: "Address"}},
											Width: w.FieldLG,
										},
										w.Button{Label: "Remove", Kind: w.ButtonDanger},
										w.Grow{},
										w.Button{Label: "Add email", Icon: "plus", Small: true},
									},
								},
								w.FieldSet{
									Rubric: w.Rubric{Text: "Other"},
									Items: []w.FieldSetPart{
										w.Field{
											Label: "Tags",
											Parts: []w.FieldPart{w.Input{Name: "tags", Value: "работа, издательство", Label: "Tags"}},
											Width: w.FieldLG,
										},
										w.Field{
											Label: "Note",
											Parts: []w.FieldPart{w.Area{Text: "Пишет по вторникам, отвечает в тот же день.", Name: "note", Label: "Text"}},
											Width: w.FieldFull,
										},
										w.Field{Parts: []w.FieldPart{w.Button{Label: "Show 7 empty fields"}}},
									},
								},
							},
							ID:     "contact-edit",
							Action: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/edit",
						},
						w.FormFoot{
							Items: []w.FormFootPart{
								w.Dim{Parts: []w.Inline{w.Text("Properties Carrel does not know are kept unchanged.")}, Size: w.Dim11_5},
							},
						},
					},
				},
				Height: w.ShellHeight560,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/rewrite",
					Icon:   "refresh",
					Title:  "Rewrite the addresses in their records?",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("Анна Ковалёва · saved")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
					},
					Body: []w.SheetPart{
						w.Rubric{Text: "The addresses"},
						w.Line{
							Items: []w.LinePart{
								w.Mono{Text: "a.kovaleva@example.org"},
								w.Dim{Parts: []w.Inline{w.Text("22 records name it")}, Size: w.Dim11_5},
							},
							Gap:   w.LineGap8,
							Block: true,
						},
						w.Line{
							Items: []w.LinePart{
								w.Mono{Text: "→"},
								w.Pick{
									Name: "to.1",
									Options: []w.PickOption{
										{Value: "a.kovaleva@slovo.example", Label: "a.kovaleva@slovo.example", On: true},
										{Value: "anna@home.example", Label: "anna@home.example"},
										{Value: "", Label: "Do not rewrite"},
									},
									Label:   "Rewrite a.kovaleva@example.org to",
									Variant: w.PickT12W240,
								},
							},
							Gap:   w.LineGap8,
							Block: true,
						},
						w.Line{
							Items: []w.LinePart{
								w.Mono{Text: "anna@oldfirm.example"},
								w.Dim{Parts: []w.Inline{w.Text("3 records name it; they keep it")}, Size: w.Dim11_5},
							},
							Gap:   w.LineGap8,
							Block: true,
						},
						w.Line{
							Items: []w.LinePart{
								w.Mono{Text: "→"},
								w.Pick{
									Name: "to.2",
									Options: []w.PickOption{
										{Value: "a.kovaleva@slovo.example", Label: "a.kovaleva@slovo.example"},
										{Value: "anna@home.example", Label: "anna@home.example"},
										{Value: "", Label: "Do not rewrite", On: true},
									},
									Label:   "Rewrite anna@oldfirm.example to",
									Variant: w.PickT12W240,
								},
							},
							Gap:   w.LineGap8,
							Block: true,
						},
						w.Rubric{Text: "What rewriting can set off"},
						w.Line{
							Items: []w.LinePart{
								w.Check{
									On:    true,
									Name:  "group",
									Value: "scheduling",
									Parts: []w.CheckPart{w.Text("On servers that schedule")},
								},
								w.Drop{
									Trigger: w.Button{Label: "9 records", Menu: true, Small: true},
									Menu: w.Menu{
										Items: []w.MenuPart{
											w.MenuItem{
												Mark:   &w.Mark{Colour: "#3F6E8C"},
												Detail: &w.Mono{Text: "10 Aug", Variant: w.Mono11},
												Label:  "Синк-ап по релизу",
												Href:   "/app/calendar/nextcloud/rabota/aa01",
											},
											w.MenuItem{
												Mark:   &w.Mark{Colour: "#3F6E8C"},
												Detail: &w.Mono{Text: "14 Aug", Variant: w.Mono11},
												Label:  "Ретро спринта",
												Href:   "/app/calendar/nextcloud/rabota/aa02",
											},
											w.MenuItem{
												Mark:   &w.Mark{Colour: "#3F6E8C"},
												Detail: &w.Mono{Text: "2 Sep", Variant: w.Mono11},
												Label:  "Планирование Q4",
												Href:   "/app/calendar/nextcloud/rabota/aa03",
											},
										},
										Width: w.MenuWidth290,
										Pop:   w.MenuUnder,
									},
								},
							},
							Gap:   w.LineGap8,
							Block: true,
						},
						w.Dim{
							Parts: []w.Inline{
								w.Text("The server may send a cancellation to the old address, an invitation to the new one and an update to the other attendees."),
							},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
						w.Line{
							Items: []w.LinePart{
								w.Check{On: true, Name: "group", Value: "unknown", Parts: []w.CheckPart{w.Text("On baikal")}},
								w.Drop{
									Trigger: w.Button{Label: "2 records", Menu: true, Small: true},
									Menu: w.Menu{
										Items: []w.MenuPart{
											w.MenuItem{
												Mark:   &w.Mark{Colour: "#9A9280"},
												Detail: &w.Mono{Text: "12 Aug", Variant: w.Mono11},
												Label:  "Общий созвон",
												Href:   "/app/calendar/baikal/obshchiy/bb01",
											},
											w.MenuItem{
												Mark:   &w.Mark{Colour: "#9A9280"},
												Detail: &w.Mono{Text: "21 Aug", Variant: w.Mono11},
												Label:  "Демо для заказчика",
												Href:   "/app/calendar/baikal/obshchiy/bb02",
											},
										},
										Width: w.MenuWidth290,
										Pop:   w.MenuUnder,
									},
								},
							},
							Gap:   w.LineGap8,
							Block: true,
						},
						w.Dim{
							Parts:   []w.Inline{w.Text("The server did not say whether it schedules.")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
						w.Line{
							Items: []w.LinePart{
								w.Check{
									On:    true,
									Name:  "group",
									Value: "others",
									Parts: []w.CheckPart{w.Text("Meetings organised by others")},
								},
								w.Drop{
									Trigger: w.Button{Label: "3 records", Menu: true, Small: true},
									Menu: w.Menu{
										Items: []w.MenuPart{
											w.MenuItem{
												Mark:   &w.Mark{Colour: "#3F6E8C"},
												Detail: &w.Mono{Text: "14 Aug", Variant: w.Mono11},
												Label:  "Ретро спринта",
												Href:   "/app/calendar/nextcloud/rabota/aa02",
											},
											w.MenuItem{
												Mark:   &w.Mark{Colour: "#3F6E8C"},
												Detail: &w.Mono{Text: "18 Aug", Variant: w.Mono11},
												Label:  "Встреча с подрядчиком",
												Href:   "/app/calendar/nextcloud/rabota/cc01",
											},
											w.MenuItem{
												Mark:   &w.Mark{Colour: "#3F6E8C"},
												Detail: &w.Mono{Text: "5 Sep", Variant: w.Mono11},
												Label:  "Юбилей отдела",
												Href:   "/app/calendar/nextcloud/rabota/cc02",
											},
										},
										Width: w.MenuWidth290,
										Pop:   w.MenuUnder,
									},
								},
							},
							Gap:   w.LineGap8,
							Block: true,
						},
						w.Dim{
							Parts:   []w.Inline{w.Text("The organiser’s server may refuse the change, and its next update may undo it.")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
						w.Line{
							Items: []w.LinePart{
								w.Check{On: true, Name: "group", Value: "past", Parts: []w.CheckPart{w.Text("Past meetings")}},
								w.Drop{
									Trigger: w.Button{Label: "5 records", Menu: true, Small: true},
									Menu: w.Menu{
										Items: []w.MenuPart{
											w.MenuItem{
												Mark:   &w.Mark{Colour: "#3F6E8C"},
												Detail: &w.Mono{Text: "3 Aug", Variant: w.Mono11},
												Label:  "Синк-ап по релизу",
												Href:   "/app/calendar/nextcloud/rabota/dd01",
											},
											w.MenuItem{
												Mark:   &w.Mark{Colour: "#3F6E8C"},
												Detail: &w.Mono{Text: "27 Jul", Variant: w.Mono11},
												Label:  "Синк-ап по релизу",
												Href:   "/app/calendar/nextcloud/rabota/dd02",
											},
										},
										Width: w.MenuWidth290,
										Pop:   w.MenuUnder,
									},
								},
							},
							Gap:   w.LineGap8,
							Block: true,
						},
						w.Dim{
							Parts:   []w.Inline{w.Text("They will show the new address, as if the person had been invited at it.")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
						w.SheetNote{
							Text: "2 records in Shared are read-only and keep the old address. Collections that are not ticked were not looked in.",
						},
						w.Tally{
							Records: []string{
								"to.1 group=scheduling",
								"to.1 group=scheduling",
								"to.1 group=scheduling",
								"to.1 group=scheduling",
								"to.1 group=scheduling",
								"to.1 group=scheduling",
								"to.1 group=scheduling",
								"to.1 group=scheduling",
								"to.1 group=scheduling group=others",
								"to.1 group=others",
								"to.1 group=others",
								"to.1 group=past",
								"to.1 group=past",
								"to.1 group=past",
								"to.1 group=past",
								"to.1 group=past",
								"to.1 group=unknown",
								"to.1 group=unknown",
								"to.1",
								"to.1",
								"to.2",
								"to.2",
								"to.2",
							},
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{
							Label:    "Rewrite 20 records",
							Kind:     w.ButtonPrimary,
							Count:    "Rewrite {n} records",
							CountOne: "Rewrite 1 record",
							Submit:   true,
						},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{
								w.Button{
									Label:  "Leave as is",
									To:     "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/rewrite?op=leave",
									Submit: true,
								},
								w.Kbd{Keys: "Esc"},
							},
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

// Contacts16 is frame «Контакт · участники переписаны» (s-contacts-16).
func Contacts16() Screen {
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "5", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
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
						Photo: &w.HeadPhoto{},
						Title: &w.HeadTitle{Text: "Анна Ковалёва"},
						Subs: []w.HeadSub{
							w.HeadSub{Parts: []w.HeadSubPart{w.Text("Издательство «Слово» · Выпускающий редактор")}},
							w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Edit", Icon: "pencil", Kind: w.ButtonPrimary},
								w.Button{Label: "Send", Icon: "mail"},
								w.Button{Label: "Print", Icon: "print"},
								w.Button{Icon: "dots", Title: "More"},
							},
						},
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Strong("Rewritten in 11 of 12 records."),
							w.Text(" «Планёрка редакции» was refused: its organiser’s server does not take changes to attendees. 2 records in Shared are read-only and keep the old address."),
						},
						Tone: w.MsgAlert,
					},
					w.Fields{
						Items: []w.FieldsItem{
							{Term: "Phone · mobile", Value: []w.ValuePart{w.Text("+7 912 445-10-88")}, Face: w.ValueMono},
							{Term: "Phone · work", Value: []w.ValuePart{w.Text("+7 495 120-45-00")}, Face: w.ValueMono},
							{Term: "Email · work", Value: []w.ValuePart{w.Text("a.kovaleva@slovo.example")}},
							{Term: "Email · second", Value: []w.ValuePart{w.Text("anna.kovaleva@slovo.example")}},
							{Term: "Address · work", Value: []w.ValuePart{w.Text("Москва, Тверская 12, стр. 3")}},
							{
								Term:  "Birthday",
								Value: []w.ValuePart{w.Text("1987-03-14 "), w.Dim{Parts: []w.Inline{w.Text("· in 7 months")}}},
								Face:  w.ValueMono,
							},
							{
								Term:  "Tags",
								Value: []w.ValuePart{w.Tag{Label: "работа"}, w.Text(" "), w.Tag{Label: "издательство"}},
							},
							{Term: "Note", Value: []w.ValuePart{w.Text("Пишет по вторникам, отвечает в тот же день.")}},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Kept as the server sent them", Items: []w.RubricPart{w.Button{Label: "Show 3", Small: true}}},
							w.SecHint{Text: "Properties Carrel does not edit are stored byte for byte and go back untouched."},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{
								Text:  "The same person in other books",
								Items: []w.RubricPart{w.Badge{Label: "2", Tone: w.BadgeLinked}},
							},
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{
										w.Mark{Colour: "#6B4E7A", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "radicale · vps · Команда", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· Phone, email differ")}},
									},
									[]w.SecItemPart{
										w.Mark{Colour: "#9A9280", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "baikal · Общий", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· The server did not answer; the record stays linked")}},
									},
								},
								Gap: 5,
							},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Timeline", Items: []w.RubricPart{w.Mono{Text: "11", Variant: w.MonoTight}}},
							w.List{
								Rows: []w.ListRow{
									w.LinkedRow{
										Mark:   w.Mark{Colour: "#4A6B52"},
										Icon:   "cal",
										Title:  "Синк-ап по релизу",
										Detail: "10 Aug · attendee · accepted",
									},
									w.LinkedRow{
										Mark:   w.Mark{Colour: "#4A6B52"},
										Icon:   "file",
										Title:  "2026-08-04-обложка.pdf",
										Detail: "04 Aug · attached to «Согласовать обложку»",
									},
									w.LinkedRow{
										Mark:   w.Mark{Colour: "#3F6E8C"},
										Icon:   "task",
										Title:  "Согласовать обложку",
										Detail: "04 Aug · organiser · done",
									},
								},
							},
							w.SecActs{Items: []w.Button{w.Button{Label: "Open the timeline", Icon: "all"}}},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Source"},
							w.Diag{Text: "Path  /dav.php/addressbooks/mix/contacts/2b7e.vcf\nRead  2 min ago · not changed on server"},
						},
					},
					w.Danger{Button: w.Button{Label: "Delete contact", Icon: "trash", Kind: w.ButtonDangerSolid}},
				},
			},
		},
	}
}

// Contacts35 is frame «Связанная запись · какую копию править, лист поверх книги» (s-contacts-35).
func Contacts35() Screen {
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
						On: 1,
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.TickRow{
								Ticked:  w.TickedPart,
								Kind:    w.TickRowAll,
								Mark:    w.Mark{Kind: w.MarkMulti},
								Label:   "Combined address book",
								Meta:    "7",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "5"},
							w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "radicale · vps"},
							w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "1", Trail: w.Badge{Label: "ro"}},
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
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 6 cards")}}},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "New contact", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
									w.Button{Label: "Import", Icon: "import"},
									w.Button{Label: "Export", Icon: "export"},
									w.Button{Label: "Print", Icon: "print"},
								},
							},
						},
						w.Bar{
							Items: []w.BarPart{
								w.Filter{
									Query:   w.Query{Placeholder: "Filter the loaded cards", Label: "Filter the loaded cards"},
									Waiting: true,
									Variant: w.FilterH26W220,
								},
								w.Seg{
									Options: []w.SegOption{
										{Label: "A–Z", Href: "/app/contacts?sort=az", On: true},
										{Label: "Recently changed", Href: "/app/contacts?sort=changed"},
									},
									Secondary: true,
								},
								w.Button{Label: "Columns", Icon: "cols"},
								w.Drop{
									Trigger: w.Button{Icon: "dots", Title: "More"},
									More:    true,
									Menu: w.Menu{
										Items: []w.MenuPart{
											w.MenuItem{Label: "A–Z", On: true, Href: "/app/contacts?sort=az"},
											w.MenuItem{Label: "Recently changed", Href: "/app/contacts?sort=changed"},
											w.MenuSep{},
											w.CheckForm{
												Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}},
												Action: "/app/contacts",
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
										Action: "/app/contacts",
										Apply:  w.Button{Label: "Apply", Small: true, Submit: true},
									},
									w.Mono{Text: "6"},
								},
								Secondary: true,
							},
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
									Mark:       w.Mark{Colour: "#6B4E7A"},
									Avatar:     w.Avatar{Photo: true},
									Title:      []w.Inline{w.Text("Артём Белов ")},
									TitleBadge: &w.Badge{Label: "linked ×2", Tone: w.BadgeLinked},
									Subtitle:   "+7 903 118-22-04",
									Detail:     "a.belov@example.org",
									Caption:    "Тензор",
									Source:     "radicale · vps · Команда",
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
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ВЛ"},
									Title:    []w.Inline{w.Text("Вера Лебедева")},
									Strong:   true,
									Subtitle: "+7 905 331-09-42",
									Detail:   "vera@example.net",
									Caption:  "—",
									Source:   "nextcloud · Архив",
									On:       true,
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Л", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
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
								},
							},
						},
					},
				},
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/contacts/linked/8d21-55ce-0b17/edit",
					Method: "get",
					Icon:   "pencil",
					Title:  "Which copy to edit?",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("Вера Лебедева · 3 copies")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
					},
					Body: []w.SheetPart{
						w.SheetNote{Text: "Saving changes one copy only; the others stay as they are."},
						w.Line{
							Items: []w.LinePart{
								w.Mark{Colour: "#A05A2C", Size: w.MarkH11},
								w.Check{
									On:    true,
									Name:  "copy",
									Value: "nextcloud/arhiv",
									Radio: true,
									Parts: []w.CheckPart{w.Text("nextcloud · Архив")},
								},
								w.Dim{Parts: []w.Inline{w.Text("the copy you opened · vera@arhiv.example")}, Size: w.Dim11_5},
							},
							Gap:   w.LineGap8,
							Block: true,
						},
						w.Line{
							Items: []w.LinePart{
								w.Mark{Colour: "#4A6B52", Size: w.MarkH11},
								w.Check{
									Name:  "copy",
									Value: "nextcloud/kontakty",
									Radio: true,
									Parts: []w.CheckPart{w.Text("nextcloud · Контакты")},
								},
								w.Dim{Parts: []w.Inline{w.Text("Email differ · v.lebedeva@example.org")}, Size: w.Dim11_5},
							},
							Gap:   w.LineGap8,
							Block: true,
						},
						w.Line{
							Items: []w.LinePart{
								w.Mark{Colour: "#9A9280", Size: w.MarkH11},
								w.Dim{Parts: []w.Inline{w.Text("baikal · Общий")}},
								w.Dim{Parts: []w.Inline{w.Text("Phone differ · read-only: it cannot be edited")}, Size: w.Dim11_5},
							},
							Gap:   w.LineGap8,
							Block: true,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Edit this copy", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/contacts/linked/8d21-55ce-0b17"}, w.Kbd{Keys: "Esc"}},
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

// Contacts7 is frame «Человек · хронология» (s-contacts-7).
func Contacts7() Screen {
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
					On: 1,
				},
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to list", Short: "Back", Href: "#"}}, Apart: true},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Where to look"},
						w.TickRow{
							Ticked:  w.TickedPart,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined calendar",
							Meta:    "20",
							Current: true,
						},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "9"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "7"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "4"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "0"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Error: true},
						w.TickRow{
							Kind:  w.TickRowExternal,
							Mark:  w.Mark{Colour: "#8C6239"},
							Label: "Праздники РФ",
							Meta:  "0",
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
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Avatar: &w.Avatar{Initials: "АК", Size: w.Avatar44},
						Kicker: &w.Rubric{Text: "Timeline", Variant: w.RubricBottom2},
						Title:  &w.HeadTitle{Text: "Анна Ковалёва"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{
									w.Text("Издательство «Слово» · Выпускающий редактор "),
									w.Badge{Label: "in 2 books", Tone: w.BadgeLinked},
								},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Write a note", Icon: "quill", Kind: w.ButtonPrimary},
								w.Button{Label: "New event", Icon: "cal"},
								w.Button{Label: "New task", Icon: "task"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
						Pad: w.HeadPadBottom10,
					},
					w.Bar{
						Items: []w.BarPart{
							w.Range{
								From:   w.Input{Kind: w.InputDate, Name: "from", Value: "2026-01-01", Label: "From", Mono: true},
								To:     w.Input{Kind: w.InputDate, Name: "to", Value: "2026-11-10", Label: "To", Mono: true},
								Show:   w.Button{Label: "Show", Submit: true},
								Action: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/timeline",
							},
							w.SegFold{
								Seg: w.Seg{
									Options: []w.SegOption{
										{
											Label: "All",
											Count: "24",
											Href:  "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/timeline",
											On:    true,
										},
										{
											Label: "Events",
											Count: "14",
											Href:  "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/timeline?kind=event",
										},
										{
											Label: "Tasks",
											Count: "2",
											Href:  "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/timeline?kind=task",
										},
										{
											Label: "Notes",
											Count: "4",
											Href:  "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/timeline?kind=note",
										},
										{
											Label: "Files",
											Count: "4",
											Href:  "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/timeline?kind=file",
										},
									},
								},
							},
							w.Button{Label: "Columns", Icon: "cols", Secondary: true},
							w.Drop{
								Trigger: w.Button{Icon: "dots", Title: "More"},
								More:    true,
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuItem{Icon: "cols", Label: "Columns"}},
									Pop:    w.MenuUnderEnd,
									Inline: true,
								},
							},
						},
						Right:   &w.BarRight{Items: []w.BarPart{w.Mono{Text: "newest first"}}, Secondary: true},
						Variant: w.BarWide,
					},
					w.Runbar{
						Meter: w.Meter{Percent: 83, Alert: true},
						Parts: []w.RunbarPart{
							w.Text("Polled 5 of 6 calendars · matched by "),
							w.Mono{Text: "a.kovaleva@example.org"},
							w.Text(", "),
							w.Mono{Text: "anna.kovaleva@slovo.example"},
							w.Text(" and the name · "),
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
					w.ListGroup{Title: w.Rubric{Text: "August 2026", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
					w.List{
						Rows: []w.ListRow{
							w.FindRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Date:   "10 Aug",
								Title:  []w.Inline{w.Text("Синк-ап по релизу")},
								Detail: []w.Inline{w.Text("attendee · accepted")},
								What:   "event",
								Source: "nextcloud · Личный",
							},
							w.FindRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Date:   "4 Aug",
								Title:  []w.Inline{w.Text("2026-08-04-обложка.pdf")},
								Detail: []w.Inline{w.Text("attached to «Согласовать обложку»")},
								What:   "file",
								Source: "sftpgo · vps · Documents",
							},
							w.FindRow{
								Mark:   w.Mark{Colour: "#3F6E8C"},
								Date:   "4 Aug",
								Title:  []w.Inline{w.Text("Согласовать обложку")},
								Detail: []w.Inline{w.Text("organiser · done")},
								What:   "task",
								Source: "nextcloud · Работа",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "July 2026", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.FindRow{
								Mark:   w.Mark{Colour: "#6B4E7A"},
								Date:   "29 Jul",
								Title:  []w.Inline{w.Text("Ревью вёрстки сборника")},
								Detail: []w.Inline{w.Text("attendee · no reply")},
								What:   "event",
								Source: "radicale · vps · Проекты",
							},
						},
					},
					w.Empty{
						Parts:   []w.EmptyPart{w.Text("Nothing was looked for before January 2026. Widen the range if you need it.")},
						Variant: w.EmptySmallTop,
					},
				},
			},
			Side: &w.Side{
				Items: []w.SidePart{
					w.Head{
						Photo:  &w.HeadPhoto{Initials: "АК", Large: true},
						Kicker: &w.Rubric{Text: "Card"},
						Tools: []w.HeadTool{
							w.Button{Label: "Open the card", Small: true},
							w.Button{Label: "Edit", Small: true},
							w.Drop{
								Trigger: w.Button{Icon: "dots", Small: true, Title: "More"},
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuItem{Icon: "move", Label: "Move"},
										w.MenuItem{Icon: "dup", Label: "Copy"},
										w.MenuItem{Icon: "export", Detail: &w.Mono{Text: ".vcf", Variant: w.Mono11}, Label: "Export"},
										w.MenuItem{Icon: "link", Label: "Publish a link"},
									},
									Width: w.MenuWidth212,
									Pop:   w.MenuUnder,
								},
							},
						},
						Acts:   &w.Acts{Items: []w.ActsPart{w.Close{Label: "Close the panel"}}},
						Middle: true,
						Pad:    w.HeadPadBottom10,
					},
					w.Fields{
						Items: []w.FieldsItem{
							{Term: "Phone · mobile", Value: []w.ValuePart{w.Text("+7 912 445-10-88")}, Face: w.ValueMono},
							{Term: "Phone · work", Value: []w.ValuePart{w.Text("+7 495 120-45-00")}, Face: w.ValueMono},
							{Term: "Email · work", Value: []w.ValuePart{w.Text("a.kovaleva@example.org")}},
							{Term: "Email · second", Value: []w.ValuePart{w.Text("anna.kovaleva@slovo.example")}},
							{
								Term:  "Birthday",
								Value: []w.ValuePart{w.Text("1987-03-14 "), w.Dim{Parts: []w.Inline{w.Text("· in 7 months")}}},
								Face:  w.ValueMono,
							},
							{
								Term:  "Tags",
								Value: []w.ValuePart{w.Tag{Label: "работа"}, w.Text(" "), w.Tag{Label: "издательство"}},
							},
							{Term: "Note", Value: []w.ValuePart{w.Text("Пишет по вторникам, отвечает в тот же день.")}},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{
								Text:  "The same person in other books",
								Items: []w.RubricPart{w.Badge{Label: "2", Tone: w.BadgeLinked}},
							},
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{
										w.Mark{Colour: "#6B4E7A", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "radicale · vps · Команда", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· Phone, email differ")}},
									},
									[]w.SecItemPart{
										w.Mark{Colour: "#9A9280", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "baikal · Общий", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· The server did not answer; the record stays linked")}},
									},
								},
								Gap: 6,
							},
							w.SecHint{
								Text:  "Linked on the duplicates screen. The records stay on their own servers; here they are shown as one person.",
								Apart: true,
							},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Source"},
							w.Diag{Text: "UID   4f2b-9c31-a0e7\nRead  2 min ago"}},
					},
					w.Danger{Button: w.Button{Label: "Delete contact", Kind: w.ButtonDangerSolid}},
				},
			},
		},
	}
}
