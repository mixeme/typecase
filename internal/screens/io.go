// Раздел макета: 7.9 Импорт и экспорт.

package screens

import w "carrel.dev/typecase/widget"

// Io1 is frame «Файл над списком · накладка перетаскивания» (s-io-1).
func Io1() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
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
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 10}, Href: "#"},
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
								Ticked: w.TickedPart,
								Kind:   w.TickRowAll,
								Mark:   w.Mark{Kind: w.MarkMulti},
								Label:  "Combined address book",
								Href:   "/app/contacts",
								Meta:   "10",
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Label:   "Контакты",
								Href:    "/app/contacts/nextcloud/kontakty",
								Meta:    "6",
								Current: true,
							},
							w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Href: "/app/contacts/nextcloud/arkhiv", Meta: "2"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "radicale · vps"},
							w.TickRow{
								Mark:  w.Mark{Colour: "#6B4E7A"},
								Label: "Команда",
								Href:  "/app/contacts/radicale/komanda",
								Meta:  "2",
								Trail: w.Badge{Label: "ro"},
							},
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
						w.PostForm{
							ID:     "post-import",
							Action: "/app/contacts/nextcloud/kontakty/import",
							Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
							File:   "file",
							Accept: ".vcf",
						},
						w.Head{
							Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH20}, Text: "Контакты"},
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · 3 cards")}}},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Split{
										Main: w.Button{Label: "New contact", Icon: "plus", Kind: w.ButtonPrimary},
										More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
									},
									w.Button{Label: "Import", Icon: "import"},
									w.Button{Label: "Export", Icon: "export"},
									w.Button{Label: "Print", Icon: "print"},
								},
							},
						},
						w.Bar{
							Items: []w.BarPart{
								w.Filter{
									Query:   w.Query{Placeholder: "Filter the loaded cards", Short: "Filter", Label: "Filter the loaded cards"},
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
							},
							Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "3"}}, Secondary: true},
						},
						w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "АС"},
									Title:    []w.Inline{w.Text("Алексей Смирнов")},
									Subtitle: "+7 916 200-71-15",
									Detail:   "—",
									Caption:  "—",
									Single:   true,
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "АК"},
									Title:    []w.Inline{w.Text("Анна Ковалёва")},
									Subtitle: "+7 912 445-10-88",
									Detail:   "a.kovaleva@example.org",
									Caption:  "Издательство «Слово»",
									Single:   true,
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "В", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "ВГ"},
									Title:    []w.Inline{w.Text("Виктор Гончаров")},
									Subtitle: "—",
									Detail:   "v.goncharov@example.org",
									Caption:  "Городская библиотека №4",
									Single:   true,
								},
							},
						},
					},
				},
			},
			Scrim: w.Scrim{
				Content: w.Empty{
					Icon: "import",
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
					Refusals: []w.Refusal{
						w.Refusal{
							Of: ".ics",
							Parts: []w.Inline{
								w.Text("Contacts take "),
								w.Mono{Text: ".vcf"},
								w.Text(" — the calendar and the list of tasks take this one"),
							},
						},
						w.Refusal{
							Of:    ".md",
							Parts: []w.Inline{w.Text("Contacts take "), w.Mono{Text: ".vcf"}, w.Text(" — a notebook takes this one")},
						},
						w.Refusal{
							Of:    ".zip",
							Parts: []w.Inline{w.Text("Contacts take "), w.Mono{Text: ".vcf"}, w.Text(" — a notebook takes this one")},
						},
						w.Refusal{
							Of:    "*",
							Parts: []w.Inline{w.Text("Only "), w.Mono{Text: ".vcf"}, w.Text(" files are imported here")},
						},
						w.Refusal{Of: "several", Parts: []w.Inline{w.Text("One file at a time: only notes take several")}},
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
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 10}, Href: "#"},
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
							Href:    "/app/calendar",
							Meta:    "10",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Href: "/app/calendar/nextcloud/lichnyi", Meta: "5"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Href: "/app/calendar/nextcloud/rabota", Meta: "2"},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#A05A2C"},
							Label:  "Дни рождения",
							Href:   "/app/calendar/nextcloud/dni-rozhdeniya",
							Meta:   "—",
							Hidden: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Href: "/app/calendar/radicale/proekty", Meta: "2"},
						w.TickRow{
							Mark:  w.Mark{Colour: "#8C6239"},
							Label: "Дежурства",
							Href:  "/app/calendar/radicale/dezhurstva",
							Meta:  "1",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "baikal"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Href: "/app/calendar/baikal/obshchii", Error: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Subscriptions", Items: []w.RubricPart{w.Badge{Label: "external"}}},
						w.TickRow{
							Kind:  w.TickRowExternal,
							Mark:  w.Mark{Colour: "#8C6239"},
							Label: "Праздники РФ",
							Href:  "/app/calendar/prazdniki-rf/feed",
							Meta:  "3",
							Trail: w.Badge{Label: "ro"},
						},
						w.TickRow{
							Ticked: w.TickedOff,
							Kind:   w.TickRowExternal,
							Mark:   w.Mark{Colour: "#6B4E7A"},
							Label:  "Расписание лиги",
							Href:   "/app/calendar/raspisanie-ligi/feed",
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
								w.Button{Label: "New event", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "import"},
								w.Button{Label: "Export", Icon: "export"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
					w.List{
						Rows: []w.ListRow{
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
				},
			},
			Side: &w.Side{
				Items: []w.SidePart{
					w.Head{
						Title: &w.HeadTitle{Text: "Синк-ап по релизу", Size: w.HeadTitle19},
						Subs: []w.HeadSub{
							w.HeadSub{Parts: []w.HeadSubPart{w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, w.Text("nextcloud · Личный")}},
						},
						Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Edit", Small: true}, w.Close{Label: "Close the panel"}}},
					},
					w.Fields{
						Items: []w.FieldsItem{
							{Term: "When", Value: []w.ValuePart{w.Mono{Text: "10:00 – 11:00"}, w.Text(", Monday, 10 August")}},
							{Term: "Where", Value: []w.ValuePart{w.Text("Переговорная 2")}},
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
										w.Dim{Parts: []w.Inline{w.Text("text/markdown · 3 KB")}},
										w.Text(" "),
										w.Badge{Label: "in the event"},
									},
								},
								Gap: w.StackGapTight,
							},
							w.Empty{
								Lead: "Attach to this event",
								Parts: []w.EmptyPart{
									w.Dim{
										Parts: []w.Inline{
											w.Text("Over 2 MB it goes to "),
											w.Mono{Text: "sftpgo · vps / Documents / carrel"},
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
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
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
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 1},
							w.Head{Title: &w.HeadTitle{Text: "Import contacts"}},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Into",
										Parts: []w.FieldPart{
											w.Drop{
												Trigger: w.Pick{Value: "nextcloud · Контакты", Mark: &w.Mark{Colour: "#4A6B52"}},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#4A6B52"},
															Label: "nextcloud · Контакты",
															On:    true,
															Href:  "/app/contacts/nextcloud/kontakty/import",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#A05A2C"},
															Label: "nextcloud · Архив",
															Href:  "/app/contacts/nextcloud/arhiv/import",
														},
														w.MenuItem{
															Mark:   &w.Mark{Colour: "#6B4E7A"},
															Trail:  w.Badge{Label: "ro", Place: w.Pushed},
															Label:  "radicale · vps · Команда",
															Reason: "Read-only: nothing can be written here.",
														},
														w.MenuSep{},
														w.MenuItem{
															Label: "New address book",
															Href:  "/app/collections/new?kind=contact&back=/app/contacts/nextcloud/kontakty/import",
														},
													},
													Width: w.MenuWidth290,
													Pop:   w.MenuUnder,
												},
											},
											w.FieldHint{Parts: []w.Inline{w.Text("Any address book you may write to, on every connected server.")}},
										},
										Width: w.FieldW100,
									},
								},
								Variant: w.FormBottom,
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
													w.Input{
														Placeholder: "Разное (mix).vcf",
														Label:       "Разное (mix).vcf",
														Picked:      true,
														ID:          "import-file",
														Variant:     w.InputFill,
													},
													w.Button{
														Label: "Browse",
														File:  &w.FileField{Form: "", Name: "file", Accept: ".vcf", Label: "vCard file", Shows: "import-file"},
													},
													w.Button{Label: "Preview file", Kind: w.ButtonPrimary, Submit: true},
												},
												Gap: w.LineGapNormal,
											},
										},
										Width: w.FieldW100,
									},
									w.CardHint{
										Parts: []w.CardTextPart{w.Mono{Text: ".vcf"}, w.Text(" — vCard 2.1, 3.0 and 4.0, any number of cards in one file.")},
										Line:  true,
									},
								},
							},
						},
						Variant: w.DialogStep560,
						Action:  "/app/contacts/nextcloud/kontakty/import",
					},
				},
			},
		},
	}
}

// Io15 is frame «Импорт событий · шаг 1 из 3, источник» (s-io-15).
func Io15() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "5", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "2"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
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
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 1},
							w.Head{Title: &w.HeadTitle{Text: "Import events"}},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Into",
										Parts: []w.FieldPart{
											w.Drop{
												Trigger: w.Pick{Value: "nextcloud · Личный", Mark: &w.Mark{Colour: "#4A6B52"}},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#4A6B52"},
															Label: "nextcloud · Личный",
															On:    true,
															Href:  "/app/calendar/nextcloud/lichnyj/import",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#3F6E8C"},
															Label: "nextcloud · Работа",
															Href:  "/app/calendar/nextcloud/rabota/import",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#A05A2C"},
															Label: "nextcloud · Дни рождения",
															Href:  "/app/calendar/nextcloud/dni-rozhdeniya/import",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#6B4E7A"},
															Label: "radicale · vps · Проекты",
															Href:  "/app/calendar/radicale/proekty/import",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#8C6239"},
															Label: "radicale · vps · Дежурства",
															Href:  "/app/calendar/radicale/dezhurstva/import",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#9A9280"},
															Label: "baikal · Общий",
															Href:  "/app/calendar/baikal/obshchij/import",
														},
														w.MenuSep{},
														w.MenuItem{
															Label: "New calendar",
															Href:  "/app/collections/new?kind=event&back=/app/calendar/nextcloud/lichnyj/import",
														},
													},
													Width: w.MenuWidth290,
													Pop:   w.MenuUnder,
												},
											},
											w.FieldHint{Parts: []w.Inline{w.Text("Any calendar you may write to, on every connected server.")}},
										},
										Width: w.FieldW100,
									},
								},
								Variant: w.FormBottom,
								Tight:   true,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "From a file"},
									w.Field{
										Label: "iCalendar file",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Input{
														Placeholder: "Личный 2026.ics",
														Label:       "Личный 2026.ics",
														Picked:      true,
														ID:          "import-file",
														Variant:     w.InputFill,
													},
													w.Button{
														Label: "Browse",
														File:  &w.FileField{Form: "", Name: "file", Accept: ".ics", Label: "iCalendar file", Shows: "import-file"},
													},
													w.Button{Label: "Preview file", Kind: w.ButtonPrimary, Submit: true},
												},
												Gap: w.LineGapNormal,
											},
										},
										Width: w.FieldW100,
									},
									w.CardHint{
										Parts: []w.CardTextPart{
											w.Mono{Text: ".ics"},
											w.Text(" — iCalendar, any number of events in one file; tasks and journal entries in it are skipped and named in the report."),
										},
										Line: true,
									},
								},
							},
						},
						Variant: w.DialogStep560,
						Action:  "/app/calendar/nextcloud/lichnyj/import",
					},
				},
			},
		},
	}
}

// Io16 is frame «Импорт задач · шаг 1 из 3, источник» (s-io-16).
func Io16() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
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
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "4"},
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
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 1},
							w.Head{Title: &w.HeadTitle{Text: "Import tasks"}},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Into",
										Parts: []w.FieldPart{
											w.Drop{
												Trigger: w.Pick{Value: "nextcloud · Дела", Mark: &w.Mark{Colour: "#4A6B52"}},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#4A6B52"},
															Label: "nextcloud · Дела",
															On:    true,
															Href:  "/app/tasks/nextcloud/dela/import",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#3F6E8C"},
															Label: "nextcloud · Работа",
															Href:  "/app/tasks/nextcloud/rabota/import",
														},
														w.MenuSep{},
														w.MenuItem{Label: "New task list", Href: "/app/collections/new?kind=task&back=/app/tasks/nextcloud/dela/import"},
													},
													Width: w.MenuWidth290,
													Pop:   w.MenuUnder,
												},
											},
											w.FieldHint{Parts: []w.Inline{w.Text("Any list of tasks you may write to, on every connected server.")}},
										},
										Width: w.FieldW100,
									},
								},
								Variant: w.FormBottom,
								Tight:   true,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "From a file"},
									w.Field{
										Label: "iCalendar file",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Input{Placeholder: "Дела.ics", Label: "Дела.ics", Picked: true, ID: "import-file", Variant: w.InputFill},
													w.Button{
														Label: "Browse",
														File:  &w.FileField{Form: "", Name: "file", Accept: ".ics", Label: "iCalendar file", Shows: "import-file"},
													},
													w.Button{Label: "Preview file", Kind: w.ButtonPrimary, Submit: true},
												},
												Gap: w.LineGapNormal,
											},
										},
										Width: w.FieldW100,
									},
									w.CardHint{
										Parts: []w.CardTextPart{
											w.Mono{Text: ".ics"},
											w.Text(" — iCalendar, any number of tasks in one file; events and journal entries in it are skipped and named in the report."),
										},
										Line: true,
									},
								},
							},
						},
						Variant: w.DialogStep560,
						Action:  "/app/tasks/nextcloud/dela/import",
					},
				},
			},
		},
	}
}

// Io17 is frame «Импорт · «Into» не выбрано» (s-io-17).
func Io17() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
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
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 1},
							w.Head{Title: &w.HeadTitle{Text: "Import contacts"}},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Into",
										Parts: []w.FieldPart{
											w.Drop{
												Trigger: w.Pick{Value: "Choose an address book", Unset: true},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#4A6B52"},
															Label: "nextcloud · Контакты",
															Href:  "/app/contacts/nextcloud/kontakty/import",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#A05A2C"},
															Label: "nextcloud · Архив",
															Href:  "/app/contacts/nextcloud/arhiv/import",
														},
														w.MenuItem{
															Mark:   &w.Mark{Colour: "#6B4E7A"},
															Trail:  w.Badge{Label: "ro", Place: w.Pushed},
															Label:  "radicale · vps · Команда",
															Reason: "Read-only: nothing can be written here.",
														},
														w.MenuSep{},
														w.MenuItem{Label: "New address book", Href: "/app/collections/new?kind=contact&back=/app/contacts/import"},
													},
													Width: w.MenuWidth290,
													Pop:   w.MenuUnder,
												},
											},
											w.FieldHint{Parts: []w.Inline{w.Text("Any address book you may write to, on every connected server.")}},
										},
										Width: w.FieldW100,
									},
								},
								Variant: w.FormBottom,
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
													w.Input{
														Value:   "Разное (mix).vcf",
														Label:   "File",
														Mono:    true,
														Picked:  true,
														ID:      "import-file",
														Variant: w.InputFill,
													},
													w.Button{
														Label: "Browse",
														File:  &w.FileField{Form: "", Name: "file", Accept: ".vcf", Label: "vCard file", Shows: "import-file"},
													},
													w.Button{
														Label:  "Preview file",
														Kind:   w.ButtonPrimary,
														Title:  "Choose under «Into» where the records go.",
														Submit: true,
														State:  w.ButtonDisabled,
													},
												},
												Gap: w.LineGapNormal,
											},
										},
										Width: w.FieldW100,
									},
									w.CardHint{
										Parts: []w.CardTextPart{w.Mono{Text: ".vcf"}, w.Text(" — vCard 2.1, 3.0 and 4.0, any number of cards in one file.")},
										Line:  true,
									},
								},
							},
						},
						Variant: w.DialogStep560,
						Action:  "/app/contacts/import",
					},
				},
			},
		},
	}
}

// Io18 is frame «Импорт · шаг 1, файл не принят» (s-io-18).
func Io18() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
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
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 1},
							w.Head{Title: &w.HeadTitle{Text: "Import contacts"}},
							w.Msg{
								Parts: []w.MsgPart{
									w.Strong("Nothing was read."),
									w.Text(" «Разное (mix).vcf» is not a vCard file: its first line is not "),
									w.Mono{Text: "BEGIN:VCARD"},
									w.Text(". Choose the file again."),
								},
								Tone:  w.MsgAlert,
								Space: w.MsgFlushSides,
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Into",
										Parts: []w.FieldPart{
											w.Drop{
												Trigger: w.Pick{Value: "nextcloud · Контакты", Mark: &w.Mark{Colour: "#4A6B52"}},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#4A6B52"},
															Label: "nextcloud · Контакты",
															On:    true,
															Href:  "/app/contacts/nextcloud/kontakty/import",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#A05A2C"},
															Label: "nextcloud · Архив",
															Href:  "/app/contacts/nextcloud/arhiv/import",
														},
														w.MenuItem{
															Mark:   &w.Mark{Colour: "#6B4E7A"},
															Trail:  w.Badge{Label: "ro", Place: w.Pushed},
															Label:  "radicale · vps · Команда",
															Reason: "Read-only: nothing can be written here.",
														},
														w.MenuSep{},
														w.MenuItem{
															Label: "New address book",
															Href:  "/app/collections/new?kind=contact&back=/app/contacts/nextcloud/kontakty/import",
														},
													},
													Width: w.MenuWidth290,
													Pop:   w.MenuUnder,
												},
											},
											w.FieldHint{Parts: []w.Inline{w.Text("Any address book you may write to, on every connected server.")}},
										},
										Width: w.FieldW100,
									},
								},
								Variant: w.FormBottom,
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
													w.Input{
														Placeholder: "Разное (mix).vcf",
														Label:       "Разное (mix).vcf",
														Picked:      true,
														ID:          "import-file",
														Variant:     w.InputFill,
													},
													w.Button{
														Label: "Browse",
														File:  &w.FileField{Form: "", Name: "file", Accept: ".vcf", Label: "vCard file", Shows: "import-file"},
													},
													w.Button{Label: "Preview file", Kind: w.ButtonPrimary, Submit: true},
												},
												Gap: w.LineGapNormal,
											},
										},
										Width: w.FieldW100,
									},
									w.CardHint{
										Parts: []w.CardTextPart{w.Mono{Text: ".vcf"}, w.Text(" — vCard 2.1, 3.0 and 4.0, any number of cards in one file.")},
										Line:  true,
									},
								},
							},
						},
						Variant: w.DialogStep560,
						Action:  "/app/contacts/nextcloud/kontakty/import",
					},
				},
			},
		},
	}
}

// Io19 is frame «Импорт · предпросмотр, коллекция не ответила» (s-io-19).
func Io19() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
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
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 2},
							w.Head{
								Title: &w.HeadTitle{Text: "Import contacts"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mono{Text: "Разное (mix).vcf"}, w.Text(" · 84 KB · vCard 3.0")}}},
								Acts: &w.Acts{
									Items: []w.ActsPart{
										w.Button{Label: "Cancel"},
										w.Button{
											Label: "Import 117 contacts",
											Kind:  w.ButtonPrimary,
											Title: "nextcloud · Контакты did not answer",
											State: w.ButtonDisabled,
										},
									},
								},
							},
							w.Msg{
								Parts: []w.MsgPart{
									w.Strong("nextcloud · Контакты did not answer."),
									w.Text(" Nothing is written; try again, or choose another address book under «Into». "),
									w.Button{Label: "Try again", Small: true, Href: "/app/contacts/nextcloud/kontakty/import"},
								},
								Tone:  w.MsgAlert,
								Space: w.MsgFlushSides,
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Into",
										Parts: []w.FieldPart{
											w.Drop{
												Trigger: w.Pick{Value: "nextcloud · Контакты", Mark: &w.Mark{Colour: "#4A6B52"}},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#4A6B52"},
															Label: "nextcloud · Контакты",
															On:    true,
															Href:  "/app/contacts/nextcloud/kontakty/import",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#A05A2C"},
															Label: "nextcloud · Архив",
															Href:  "/app/contacts/nextcloud/arhiv/import",
														},
														w.MenuItem{
															Mark:   &w.Mark{Colour: "#6B4E7A"},
															Trail:  w.Badge{Label: "ro", Place: w.Pushed},
															Label:  "radicale · vps · Команда",
															Reason: "Read-only: nothing can be written here.",
														},
														w.MenuSep{},
														w.MenuItem{
															Label: "New address book",
															Href:  "/app/collections/new?kind=contact&back=/app/contacts/nextcloud/kontakty/import",
														},
													},
													Width: w.MenuWidth290,
													Pop:   w.MenuUnder,
												},
											},
											w.FieldHint{
												Parts: []w.Inline{w.Text("Any address book you may write to. What is already here is counted against it.")},
											},
										},
										Width: w.FieldLG,
									},
								},
								Tight: true,
							},
							w.Stats{
								Items: []w.Stat{w.Stat{Value: "117", Label: "ready"}, w.Stat{Value: "3", Label: "with errors", Error: true}},
							},
							w.Table{
								Cols: []w.TableCol{{Label: "Name"}, {Label: "Email", Rank: w.Rank3}, {Label: "Result", Rank: w.Rank4, Main: true}},
								Rows: []w.TableRow{
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Анна Ковалёва")}},
											{Parts: []w.CellPart{w.Text("a.kovaleva@example.org")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Ready")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Артём Белов")}},
											{Parts: []w.CellPart{w.Text("a.belov@example.org")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Ready")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("—")}, Dim: true},
											{Parts: []w.CellPart{w.Text("—")}, Mono: true, Tone: w.CellDusty},
											{Parts: []w.CellPart{w.Text("line 12: cannot decode quoted-printable")}, Tone: w.CellDanger},
										},
									},
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

// Io4 is frame «Импорт заметок · шаг 1 из 3, источник» (s-io-4).
func Io4() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "16", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "3"},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#3F6E8C"},
							Label:  "Общие с семьёй",
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
					w.Dialog{
						Items: []w.DialogPart{
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 1},
							w.Head{Title: &w.HeadTitle{Text: "Import notes"}},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Into",
										Parts: []w.FieldPart{
											w.Drop{
												Trigger: w.Pick{Value: "nextcloud · Записи", Mark: &w.Mark{Colour: "#4A6B52"}},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#4A6B52"},
															Label: "nextcloud · Записи",
															On:    true,
															Href:  "/app/notes/nextcloud/zapisi/import",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#6B4E7A"},
															Label: "nextcloud · Дневник",
															Href:  "/app/notes/nextcloud/dnevnik/import",
														},
														w.MenuItem{
															Mark:   &w.Mark{Colour: "#3F6E8C"},
															Trail:  w.Badge{Label: "ro", Place: w.Pushed},
															Label:  "nextcloud · Общие с семьёй",
															Reason: "Read-only: nothing can be written here.",
														},
														w.MenuSep{},
														w.MenuItem{
															Label: "New notebook",
															Href:  "/app/collections/new?kind=note&back=/app/notes/nextcloud/zapisi/import",
														},
													},
													Width: w.MenuWidth290,
													Pop:   w.MenuUnder,
												},
											},
											w.FieldHint{Parts: []w.Inline{w.Text("The notebook the notes are created in. Every source below lands there.")}},
										},
										Width: w.FieldW100,
									},
								},
								Variant: w.FormBottom,
								Tight:   true,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "From a file"},
									w.Field{
										Label: "Markdown, a zip of them, or .ics",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Input{
														Placeholder: "заметки-2026.zip",
														Label:       "заметки-2026.zip",
														Picked:      true,
														ID:          "import-file",
														Variant:     w.InputFill,
													},
													w.Button{
														Label: "Browse",
														File: &w.FileField{
															Form:   "",
															Name:   "file",
															Accept: ".md,.zip,.ics",
															Label:  "Markdown, a zip of them, or .ics",
															Shows:  "import-file",
															Many:   true,
														},
													},
													w.Button{Label: "Preview file", Kind: w.ButtonPrimary, Submit: true},
												},
												Gap: w.LineGapNormal,
											},
										},
										Width: w.FieldW100,
									},
									w.CardHint{
										Parts: []w.CardTextPart{
											w.Mono{Text: ".md"},
											w.Text(" — Markdown, "),
											w.Mono{Text: ".zip"},
											w.Text(" — any number of them, "),
											w.Mono{Text: ".ics"},
											w.Text(" — iCalendar journal entries. Several files at once; events and tasks in an "),
											w.Mono{Text: ".ics"},
											w.Text(" are skipped and named in the report."),
										},
										Line: true,
									},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "From a WebDAV folder"},
									w.Tree{
										Rows: []w.TreeRow{
											{Icon: "file", Label: "Files", Trail: w.Dim{Parts: []w.Inline{w.Text("sftpgo")}, Size: w.Dim11}},
											{Icon: "file", Label: "Documents", Indent: w.TreeLevel22},
											{Icon: "file", Label: "notes", Indent: w.TreeLevel36},
											{Icon: "file", Label: "carrel", Indent: w.TreeLevel36},
											{Icon: "file", Label: "Files", Trail: w.Dim{Parts: []w.Inline{w.Text("nextcloud")}, Size: w.Dim11}},
										},
										On: 3,
									},
									w.CardRow{
										Kind: w.CardTop8,
										Gap:  w.LineGapNormal,
										Items: []w.CardRowPart{
											w.FillHint{
												Parts: []w.CardTextPart{
													w.Text("Every "),
													w.Mono{Text: ".md"},
													w.Text(" directly in that folder, without subfolders. Nothing on the server is moved or deleted."),
												},
											},
											w.Button{Label: "Preview folder", Kind: w.ButtonPrimary},
										},
									},
								},
								Variant: w.CardChosen,
							},
							w.Msg{
								Parts: []w.MsgPart{w.Text("A file without front matter is dated by its own timestamp.")},
								Space: w.MsgFlushSides,
							},
						},
						Variant: w.DialogStep560,
						Action:  "/app/notes/nextcloud/zapisi/import",
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
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
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
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 2},
							w.Head{
								Title: &w.HeadTitle{Text: "Import contacts"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mono{Text: "Разное (mix).vcf"}, w.Text(" · 84 KB · vCard 3.0")}}},
								Acts: &w.Acts{
									Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Import 117 contacts", Kind: w.ButtonPrimary}},
								},
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Into",
										Parts: []w.FieldPart{
											w.Drop{
												Trigger: w.Pick{Value: "nextcloud · Контакты", Mark: &w.Mark{Colour: "#4A6B52"}},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#4A6B52"},
															Label: "nextcloud · Контакты",
															On:    true,
															Href:  "/app/contacts/nextcloud/kontakty/import",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#A05A2C"},
															Label: "nextcloud · Архив",
															Href:  "/app/contacts/nextcloud/arhiv/import",
														},
														w.MenuItem{
															Mark:   &w.Mark{Colour: "#6B4E7A"},
															Trail:  w.Badge{Label: "ro", Place: w.Pushed},
															Label:  "radicale · vps · Команда",
															Reason: "Read-only: nothing can be written here.",
														},
														w.MenuSep{},
														w.MenuItem{
															Label: "New address book",
															Href:  "/app/collections/new?kind=contact&back=/app/contacts/nextcloud/kontakty/import",
														},
													},
													Width: w.MenuWidth290,
													Pop:   w.MenuUnder,
												},
											},
											w.FieldHint{
												Parts: []w.Inline{w.Text("Any address book you may write to. What is already here is counted against it.")},
											},
										},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "5 records of this file are already in Контакты",
										Parts: []w.FieldPart{
											w.Seg{
												Options: []w.SegOption{
													{Label: "Skip them", Href: "/app/contacts/nextcloud/kontakty/import?existing=skip", On: true},
													{Label: "Replace them", Href: "/app/contacts/nextcloud/kontakty/import?existing=replace"},
													{Label: "Import as copies", Href: "/app/contacts/nextcloud/kontakty/import?existing=copy"},
												},
											},
											w.FieldHint{Parts: []w.Inline{w.Text("They stay as they are here; the file's version is not written.")}},
										},
										OwnLine: true,
									},
								},
								Tight: true,
							},
							w.Stats{
								Items: []w.Stat{
									w.Stat{Value: "117", Label: "ready"},
									w.Stat{Value: "3", Label: "with errors", Error: true},
									w.Stat{Value: "5", Label: "already here"},
								},
							},
							w.Table{
								Cols: []w.TableCol{{Label: "Name"}, {Label: "Email", Rank: w.Rank3}, {Label: "Result", Rank: w.Rank4, Main: true}},
								Rows: []w.TableRow{
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Анна Ковалёва")}},
											{Parts: []w.CellPart{w.Text("a.kovaleva@example.org")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Ready")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Вера Лебедева")}},
											{Parts: []w.CellPart{w.Text("vera@example.net")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Ready — name from N")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Артём Белов")}},
											{Parts: []w.CellPart{w.Text("a.belov@example.org")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Already here — skipped")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("—")}, Dim: true},
											{Parts: []w.CellPart{w.Text("—")}, Mono: true, Tone: w.CellDusty},
											{Parts: []w.CellPart{w.Text("line 12: cannot decode quoted-printable")}, Tone: w.CellDanger},
										},
									},
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
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "5", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "2"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
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
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 2},
							w.Head{
								Title: &w.HeadTitle{Text: "Import events"},
								Subs: []w.HeadSub{
									w.HeadSub{Parts: []w.HeadSubPart{w.Mono{Text: "Личный 2026.ics"}, w.Text(" · 210 KB · iCalendar 2.0")}},
								},
								Acts: &w.Acts{
									Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Import 82 events", Kind: w.ButtonPrimary}},
								},
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Into",
										Parts: []w.FieldPart{
											w.Drop{
												Trigger: w.Pick{Value: "nextcloud · Личный", Mark: &w.Mark{Colour: "#4A6B52"}},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#4A6B52"},
															Label: "nextcloud · Личный",
															On:    true,
															Href:  "/app/calendar/nextcloud/lichnyj/import",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#3F6E8C"},
															Label: "nextcloud · Работа",
															Href:  "/app/calendar/nextcloud/rabota/import",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#A05A2C"},
															Label: "nextcloud · Дни рождения",
															Href:  "/app/calendar/nextcloud/dni-rozhdeniya/import",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#6B4E7A"},
															Label: "radicale · vps · Проекты",
															Href:  "/app/calendar/radicale/proekty/import",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#8C6239"},
															Label: "radicale · vps · Дежурства",
															Href:  "/app/calendar/radicale/dezhurstva/import",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#9A9280"},
															Label: "baikal · Общий",
															Href:  "/app/calendar/baikal/obshchij/import",
														},
														w.MenuSep{},
														w.MenuItem{
															Label: "New calendar",
															Href:  "/app/collections/new?kind=event&back=/app/calendar/nextcloud/lichnyj/import",
														},
													},
													Width: w.MenuWidth290,
													Pop:   w.MenuUnder,
												},
											},
											w.FieldHint{Parts: []w.Inline{w.Text("Calendars only: a list of tasks is a different collection.")}},
										},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "3 events of this file are already in Личный",
										Parts: []w.FieldPart{
											w.Seg{
												Options: []w.SegOption{
													{Label: "Skip them", Href: "/app/calendar/nextcloud/lichnyj/import?existing=skip", On: true},
													{Label: "Replace them", Href: "/app/calendar/nextcloud/lichnyj/import?existing=replace"},
													{Label: "Import as copies", Href: "/app/calendar/nextcloud/lichnyj/import?existing=copy"},
												},
											},
											w.FieldHint{Parts: []w.Inline{w.Text("They stay as they are here; the file's version is not written.")}},
										},
										OwnLine: true,
									},
								},
								Tight: true,
							},
							w.Stats{
								Items: []w.Stat{
									w.Stat{Value: "82", Label: "ready"},
									w.Stat{Value: "1", Label: "with errors", Error: true},
									w.Stat{Value: "3", Label: "already here"},
									w.Stat{Value: "6", Label: "skipped"},
								},
							},
							w.Table{
								Cols: []w.TableCol{{Label: "Summary"}, {Label: "When", Rank: w.Rank3}, {Label: "Result", Rank: w.Rank4, Main: true}},
								Rows: []w.TableRow{
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Дежурство: первая линия")}},
											{Parts: []w.CellPart{w.Text("Monday, 10 August 2026, 09:00 · weekly until 28 December")}},
											{Parts: []w.CellPart{w.Text("Ready — the series as one record")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Отпуск")}},
											{Parts: []w.CellPart{w.Text("13 – 26 July 2026 · all day")}},
											{Parts: []w.CellPart{w.Text("Ready")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Созвон с подрядчиком")}},
											{Parts: []w.CellPart{w.Text("17 August 2026, 08:00")}},
											{Parts: []w.CellPart{w.Text("Ready — kept as sent: 10:00 Asia/Yekaterinburg")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Купить билеты "), w.Badge{Label: "VTODO"}}},
											{Parts: []w.CellPart{w.Text("—")}, Mono: true, Tone: w.CellDusty},
											{
												Parts: []w.CellPart{w.Text("Skipped: a calendar takes events — import it into a list of tasks")},
												Tone:  w.CellDusty,
											},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Untitled event")}, Dim: true},
											{Parts: []w.CellPart{w.Text("4 May 2026 · all day")}},
											{Parts: []w.CellPart{w.Text("Ready — a whole day (no end), no title")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("—")}, Dim: true},
											{Parts: []w.CellPart{w.Text("—")}, Mono: true, Tone: w.CellDusty},
											{Parts: []w.CellPart{w.Text("line 23: cannot read DTSTART «2026-02-30»")}, Tone: w.CellDanger},
										},
									},
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
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
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
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "4"},
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
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 2},
							w.Head{
								Title: &w.HeadTitle{Text: "Import tasks"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mono{Text: "Дела.ics"}, w.Text(" · 12 KB · iCalendar 2.0")}}},
								Acts: &w.Acts{
									Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Import 28 tasks", Kind: w.ButtonPrimary}},
								},
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Into",
										Parts: []w.FieldPart{
											w.Drop{
												Trigger: w.Pick{Value: "nextcloud · Дела", Mark: &w.Mark{Colour: "#4A6B52"}},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#4A6B52"},
															Label: "nextcloud · Дела",
															On:    true,
															Href:  "/app/tasks/nextcloud/dela/import",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#3F6E8C"},
															Label: "nextcloud · Работа",
															Href:  "/app/tasks/nextcloud/rabota/import",
														},
														w.MenuSep{},
														w.MenuItem{Label: "New task list", Href: "/app/collections/new?kind=task&back=/app/tasks/nextcloud/dela/import"},
													},
													Width: w.MenuWidth290,
													Pop:   w.MenuUnder,
												},
											},
											w.FieldHint{Parts: []w.Inline{w.Text("Lists of tasks only: a calendar is a different collection.")}},
										},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "2 tasks of this file are already in Дела",
										Parts: []w.FieldPart{
											w.Seg{
												Options: []w.SegOption{
													{Label: "Skip them", Href: "/app/tasks/nextcloud/dela/import?existing=skip", On: true},
													{Label: "Replace them", Href: "/app/tasks/nextcloud/dela/import?existing=replace"},
													{Label: "Import as copies", Href: "/app/tasks/nextcloud/dela/import?existing=copy"},
												},
											},
											w.FieldHint{Parts: []w.Inline{w.Text("They stay as they are here; the file's version is not written.")}},
										},
										OwnLine: true,
									},
								},
								Tight: true,
							},
							w.Stats{
								Items: []w.Stat{
									w.Stat{Value: "28", Label: "ready"},
									w.Stat{Value: "1", Label: "with errors", Error: true},
									w.Stat{Value: "2", Label: "already here"},
									w.Stat{Value: "5", Label: "skipped"},
								},
							},
							w.Table{
								Cols: []w.TableCol{{Label: "Summary"}, {Label: "Due", Rank: w.Rank3}, {Label: "Result", Rank: w.Rank4, Main: true}},
								Rows: []w.TableRow{
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Отправить документы в МФЦ")}},
											{Parts: []w.CellPart{w.Text("7 August 2026")}},
											{Parts: []w.CellPart{w.Text("Ready — overdue, and stays overdue")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Ответить подрядчику по смете")}},
											{Parts: []w.CellPart{w.Text("—")}, Mono: true, Tone: w.CellDusty},
											{Parts: []w.CellPart{w.Text("Ready — a task without a date is a task")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Дописать раздел про кэш")}},
											{Parts: []w.CellPart{w.Text("4 September 2026")}},
											{Parts: []w.CellPart{w.Text("Ready — 40 % done, priority 1")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Созвон с подрядчиком "), w.Badge{Label: "VEVENT"}}},
											{Parts: []w.CellPart{w.Text("17 August 2026")}},
											{
												Parts: []w.CellPart{w.Text("Skipped: a list of tasks takes tasks — import it into a calendar")},
												Tone:  w.CellDusty,
											},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Untitled task")}, Dim: true},
											{Parts: []w.CellPart{w.Text("10 September 2026")}},
											{Parts: []w.CellPart{w.Text("Ready — no title")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("Сдать отчёт")}},
											{Parts: []w.CellPart{w.Text("—")}, Mono: true, Tone: w.CellDusty},
											{Parts: []w.CellPart{w.Text("line 88: cannot read DUE «2026-13-01»")}, Tone: w.CellDanger},
										},
									},
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
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "16", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "3"},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#3F6E8C"},
							Label:  "Общие с семьёй",
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
					w.Dialog{
						Items: []w.DialogPart{
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 2},
							w.Head{
								Title: &w.HeadTitle{Text: "Import notes"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mono{Text: "заметки-2026.zip"}, w.Text(" · 1.2 MB · 64 files")}}},
								Acts: &w.Acts{
									Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Import 60 notes", Kind: w.ButtonPrimary}},
								},
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Into",
										Parts: []w.FieldPart{
											w.Drop{
												Trigger: w.Pick{Value: "nextcloud · Записи", Mark: &w.Mark{Colour: "#4A6B52"}},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#4A6B52"},
															Label: "nextcloud · Записи",
															On:    true,
															Href:  "/app/notes/nextcloud/zapisi/import",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#6B4E7A"},
															Label: "nextcloud · Дневник",
															Href:  "/app/notes/nextcloud/dnevnik/import",
														},
														w.MenuItem{
															Mark:   &w.Mark{Colour: "#3F6E8C"},
															Trail:  w.Badge{Label: "ro", Place: w.Pushed},
															Label:  "nextcloud · Общие с семьёй",
															Reason: "Read-only: nothing can be written here.",
														},
														w.MenuSep{},
														w.MenuItem{
															Label: "New notebook",
															Href:  "/app/collections/new?kind=note&back=/app/notes/nextcloud/zapisi/import",
														},
													},
													Width: w.MenuWidth290,
													Pop:   w.MenuUnder,
												},
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Any notebook you may write to. A WebDAV folder from the step before lands in this same table."),
												},
											},
										},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "1 note of this file is already in Записи",
										Parts: []w.FieldPart{
											w.Seg{
												Options: []w.SegOption{
													{Label: "Skip it", Href: "/app/notes/nextcloud/zapisi/import?existing=skip", On: true},
													{Label: "Replace it", Href: "/app/notes/nextcloud/zapisi/import?existing=replace"},
													{Label: "Import as a copy", Href: "/app/notes/nextcloud/zapisi/import?existing=copy"},
												},
											},
											w.FieldHint{Parts: []w.Inline{w.Text("It stays as it is here; the file's version is not written.")}},
										},
										OwnLine: true,
									},
								},
								Tight: true,
							},
							w.Stats{
								Items: []w.Stat{
									w.Stat{Value: "60", Label: "ready"},
									w.Stat{Value: "2", Label: "with errors", Error: true},
									w.Stat{Value: "1", Label: "already here"},
									w.Stat{Value: "1", Label: "skipped"},
								},
							},
							w.Table{
								Cols: []w.TableCol{
									{Label: "File"},
									{Label: "Title"},
									{Label: "Date", Rank: w.Rank3},
									{Label: "Result", Rank: w.Rank4, Main: true},
								},
								Rows: []w.TableRow{
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("2026-03-11-podryadchik.md")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Встреча с подрядчиком")}},
											{Parts: []w.CellPart{w.Text("11 March 2026")}},
											{Parts: []w.CellPart{w.Text("Ready — title, date and tags from front matter")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("idei/crdt.md")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Что почитать про CRDT")}},
											{Parts: []w.CellPart{w.Text("4 August 2026")}},
											{
												Parts: []w.CellPart{w.Text("Ready — no front matter: title from the first heading, date from the file")},
											},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("doroga.md")}, Mono: true},
											{Parts: []w.CellPart{w.Text("Про дорогу до дачи")}},
											{Parts: []w.CellPart{w.Text("30 June 2026")}},
											{Parts: []w.CellPart{w.Text("Already here — skipped")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("skan.pdf")}, Mono: true},
											{Parts: []w.CellPart{w.Text("—")}, Dim: true},
											{Parts: []w.CellPart{w.Text("—")}, Mono: true, Tone: w.CellDusty},
											{Parts: []w.CellPart{w.Text("Skipped: not Markdown")}, Tone: w.CellDusty},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.Text("pustoj.md")}, Mono: true},
											{Parts: []w.CellPart{w.Text("—")}, Dim: true},
											{Parts: []w.CellPart{w.Text("—")}, Mono: true, Tone: w.CellDusty},
											{
												Parts: []w.CellPart{w.Text("Empty file: nothing to store and nothing to name it by")},
												Tone:  w.CellDanger,
											},
										},
									},
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
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "6", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "2"},
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
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 3},
							w.Head{
								Title: &w.HeadTitle{Text: "Import report"},
								Subs: []w.HeadSub{
									w.HeadSub{Parts: []w.HeadSubPart{w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, w.Text("nextcloud · Контакты")}},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Created"},
									w.CardText{Kind: w.CardNum, Parts: []w.CardTextPart{w.Text("117")}}},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Already here, skipped", Items: []w.RubricPart{w.Mono{Text: "5", Variant: w.MonoTight}}},
									w.Diag{Text: "Артём Белов · 1a77-40de-bb02\nИрина Соколова · 33b0-9ffa-1120"},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "With a new UID", Items: []w.RubricPart{w.Mono{Text: "2", Variant: w.MonoTight}}},
									w.Diag{
										Text: "Вера Дорошенко · 2b7e-11aa → 9d04-77c1 · the file holds this UID twice\nВиктор Панов · 51c0-3a8e → c3e1-0412 · the server gave its own",
									},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Not kept by the server", Items: []w.RubricPart{w.Mono{Text: "1", Variant: w.MonoTight}}},
									w.Diag{Text: "Галина Мещерякова · PHOTO"},
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
						Variant: w.DialogStep520,
					},
				},
			},
		},
	}
}

// Io14 is frame «Отчёт об импорте · с пропущенным» (s-io-14).
func Io14() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "5", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "2"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
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
							w.Steps{Items: []string{"1 · Source", "2 · Preview", "3 · Report"}, On: 3},
							w.Head{
								Title: &w.HeadTitle{Text: "Import report"},
								Subs: []w.HeadSub{
									w.HeadSub{Parts: []w.HeadSubPart{w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, w.Text("nextcloud · Личный")}},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Created"},
									w.CardText{Kind: w.CardNum, Parts: []w.CardTextPart{w.Text("82")}}},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Already here, skipped", Items: []w.RubricPart{w.Mono{Text: "3", Variant: w.MonoTight}}},
									w.Diag{Text: "Синк-ап по релизу · 7c1e-22ab-90f4\nПоездка в Казань · 0d5a-b8c3-41e7"},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Failed", Items: []w.RubricPart{w.Mono{Text: "1", Variant: w.MonoTight}}},
									w.Diag{Text: "line 418: cannot read RRULE «FREQ=FORTNIGHTLY»", Variant: w.DiagDanger},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Skipped", Items: []w.RubricPart{w.Mono{Text: "6", Variant: w.MonoTight}}},
									w.Diag{Text: "Купить билеты · VTODO: a calendar takes events — import it into a list of tasks"},
								},
							},
							w.FormFoot{
								Items:   []w.FormFootPart{w.Button{Label: "Back to list", Kind: w.ButtonPrimary}},
								Variant: w.FormFootFlush,
							},
						},
						Variant: w.DialogStep520,
					},
				},
			},
		},
	}
}
