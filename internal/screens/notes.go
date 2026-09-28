// Раздел макета: 7.4 Заметки.

package screens

import w "carrel.dev/typecase/widget"

// Notes1 is frame «Заметки · список и заметка в панели» (s-notes-1).
func Notes1() Screen {
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
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{
						w.Tag{Label: "all", On: true},
						w.Tag{Label: "встречи"},
						w.Tag{Label: "идеи"},
						w.Tag{Label: "книги"},
						w.Tag{Label: "рецепты"},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New notebook"},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined notebook"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 of 3 notebooks · 86 notes · VJOURNAL")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New note", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "up"},
								w.Button{Label: "Export", Icon: "down"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Input{Kind: w.InputSearch, Placeholder: "Filter by word", Label: "Filter by word", Variant: w.InputH26W200},
							w.Seg{Options: []w.SegOption{{Label: "Newest", On: true}, {Label: "Oldest"}, {Label: "Title"}}},
							w.Button{Label: "Columns", Icon: "cols"},
						},
						Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "86"}}, Secondary: true},
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
								Mark:       w.Mark{Colour: "#4A6B52"},
								Title:      []w.Inline{w.Text("Что почитать про CRDT ")},
								TitleBadge: &w.Badge{Label: "linked ×2", Tone: w.BadgeLinked},
								Excerpt:    "Кляйн, глава 5; статья Шапиро 2011; демка automerge…",
								Tags:       []w.RowTag{w.Tag{Label: "книги"}},
								Date:       "8 Aug",
								Source:     "2 sources",
								SourceHint: "nextcloud · Записи, radicale · vps · Дневник",
							},
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#6B4E7A"},
								Title:   []w.Inline{w.Text("Про дорогу до дачи")},
								Excerpt: "Съезд после заправки, потом грунтовка два километра…",
								Tags:    []w.RowTag{w.Badge{Label: "linked", Tone: w.BadgeLinked, Icon: "link"}},
								Date:    "3 Aug",
								Source:  "nextcloud · Дневник",
							},
						},
					},
				},
			},
			Side: &w.Side{
				Items: []w.SidePart{
					w.Head{
						Title: &w.HeadTitle{Text: "Заметка со встречи 10.08", Size: w.HeadTitle19},
						Tools: []w.Button{
							w.Button{Label: "Edit", Small: true},
							w.Button{Icon: "expand", Small: true, Title: "Open the note page"},
							w.Button{Icon: "dots", Small: true, Title: "More"},
						},
						Subs: []w.HeadSub{
							w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Записи · "), w.Mono{Text: "10 August 2026, 11:05"}}},
						},
						Acts:   &w.Acts{Items: []w.ActsPart{w.Close{Label: "Close the panel"}}},
						Middle: true,
					},
					w.Sec{
						Items: []w.SecPart{
							w.Prose{
								Parts: []w.Inline{
									w.Text("Договорились заморозить ветку в среду.\n\nПисьмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт."),
								},
								Variant: w.ProseRead,
							},
							w.SecTasks{
								Items: []w.DocTask{
									{Tick: w.Tick{Variant: w.TickNudge, Label: "проверить changelog"}, Text: "проверить changelog"},
									{
										Tick: w.Tick{Variant: w.TickNudge, Label: "попросить у Артёма скриншоты"},
										Text: "попросить у Артёма скриншоты",
									},
								},
							},
						},
					},
					w.Fields{
						Items: []w.FieldsItem{
							{Term: "Tags", Value: []w.ValuePart{w.Tag{Label: "встречи"}, w.Text(" "), w.Tag{Label: "релиз"}}}},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Attachments", Items: []w.RubricPart{w.Button{Label: "Attach", Menu: true, Small: true}}},
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{
										w.Text("2026-08-10-sync-board.png "),
										w.Dim{Parts: []w.Inline{w.Text("image/png · 412 KB")}},
										w.Text(" "),
										w.Badge{Label: "in the note"},
									},
									[]w.SecItemPart{
										w.Text("запись-встречи.m4a "),
										w.Dim{Parts: []w.Inline{w.Text("audio/mp4 · 24 MB")}},
										w.Text(" "),
										w.Badge{Label: "Documents / carrel"},
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
										w.Span{Face: w.FaceAccent, Text: "Синк-ап по релизу"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· event · 10 Aug")}},
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

// Notes2 is frame «Заметка · чтение на весь экран» (s-notes-2).
func Notes2() Screen {
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
					{Icon: "back", Label: "Back to list", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "64", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "22"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Neighbours"},
						w.Tree{
							Rows: []w.TreeRow{
								{Icon: "note", Label: "Планы на осень"},
								{Icon: "note", Label: "Заметка со встречи"},
								{Icon: "note", Label: "Что почитать про CRDT"},
							},
							On:      2,
							Variant: w.TreePadded,
						},
					},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Meta: &w.HeadSub{
							Parts: []w.HeadSubPart{
								w.Mark{Colour: "#4A6B52", Size: w.MarkH12},
								w.Text("nextcloud · Записи · "),
								w.Mono{Text: "10 August 2026, 11:05"},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Edit", Icon: "pencil", Kind: w.ButtonPrimary},
								w.Button{Label: "Export Markdown"},
								w.Button{Label: "Print", Icon: "print"},
								w.Button{Label: "Source"},
								w.Seg{Options: []w.SegOption{{Label: "Full width", On: true}, {Label: "Reading width"}}},
								w.Button{Icon: "panel", Title: "Hide the side column"},
								w.Button{Icon: "dots", Title: "More"},
							},
						},
						Pad: w.HeadRule,
					},
					w.DocLayout{
						Items: []w.DocLayoutPart{
							w.Doc{
								Items: []w.DocPart{
									w.DocTitle{Text: "Заметка со встречи 10.08"},
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
									w.Prose{
										Parts: []w.Inline{
											w.Text("Перенос старых вложений в новую папку — Виктор посмотрит, сколько их вообще, и скажет в четверг."),
										},
										Variant: w.ProseLeadAfter,
									},
								},
							},
							w.Docmeta{
								Sections: []w.DocmetaSec{
									{
										Title: w.Rubric{Text: "Tags"},
										Items: []w.DocmetaPart{
											w.Line{Block: true, Gap: w.LineGap4, Items: []w.LinePart{w.Tag{Label: "встречи"}, w.Tag{Label: "релиз"}}},
										},
									},
									{
										Title: w.Rubric{Text: "Related"},
										Items: []w.DocmetaPart{
											w.DocRefs{
												Items: []w.DocRef{
													{Name: "Синк-ап по релизу", What: w.Dim{Parts: []w.Inline{w.Text("· event · 10 Aug")}}},
													{Name: "Дописать раздел про кэш", What: w.Dim{Parts: []w.Inline{w.Text("· task")}}},
												},
											},
										},
									},
									{
										Title: w.Rubric{Text: "Attachments", Items: []w.RubricPart{w.Button{Label: "Attach", Menu: true, Small: true}}},
										Items: []w.DocmetaPart{
											w.DocAttach{
												Name:  "2026-08-10-sync-board.png",
												Size:  w.Dim{Parts: []w.Inline{w.Text("· 412 KB")}},
												Where: w.Badge{Label: "in the note"},
											},
											w.DocAttach{
												Name:  "запись-встречи.m4a",
												Size:  w.Dim{Parts: []w.Inline{w.Text("· 24 MB")}},
												Where: w.Badge{Label: "Documents / carrel"},
											},
										},
									},
									{
										Title: w.Rubric{
											Text:  "The same note in other notebooks",
											Items: []w.RubricPart{w.Badge{Label: "1", Tone: w.BadgeLinked}},
										},
										Items: []w.DocmetaPart{
											w.DocRefs{
												Items: []w.DocRef{
													{
														Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11Inline},
														Name: "nextcloud · Дневник",
														Href: "#",
														What: w.Dim{Parts: []w.Inline{w.Text("· Text differ")}},
													},
												},
											},
										},
									},
									{
										Title: w.Rubric{Text: "Source"},
										Items: []w.DocmetaPart{w.Diag{Text: "UID   9e31-40b7-c012\nRead  2 min ago"}},
									},
								},
								Danger: &w.Danger{Button: w.Button{Label: "Delete note", Kind: w.ButtonDangerSolid}, Space: w.DangerAtEnd},
							},
						},
					},
				},
			},
		},
	}
}

// Notes3 is frame «Заметка · чтение, Reading width» (s-notes-3).
func Notes3() Screen {
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
					{Icon: "back", Label: "Back to list", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "64", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "22"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Neighbours"},
						w.Tree{
							Rows: []w.TreeRow{
								{Icon: "note", Label: "Планы на осень"},
								{Icon: "note", Label: "Заметка со встречи"},
								{Icon: "note", Label: "Что почитать про CRDT"},
							},
							On:      2,
							Variant: w.TreePadded,
						},
					},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Meta: &w.HeadSub{
							Parts: []w.HeadSubPart{
								w.Mark{Colour: "#4A6B52", Size: w.MarkH12},
								w.Text("nextcloud · Записи · "),
								w.Mono{Text: "10 August 2026, 11:05"},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Edit", Icon: "pencil", Kind: w.ButtonPrimary},
								w.Button{Label: "Export Markdown"},
								w.Button{Label: "Print", Icon: "print"},
								w.Button{Label: "Source"},
								w.Seg{Options: []w.SegOption{{Label: "Full width"}, {Label: "Reading width", On: true}}},
								w.Button{Icon: "panel", Title: "Hide the side column"},
								w.Button{Icon: "dots", Title: "More"},
							},
						},
						Pad: w.HeadRule,
					},
					w.DocLayout{
						Items: []w.DocLayoutPart{
							w.Doc{
								Items: []w.DocPart{
									w.DocTitle{Text: "Заметка со встречи 10.08"},
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
									w.Prose{
										Parts: []w.Inline{
											w.Text("Перенос старых вложений в новую папку — Виктор посмотрит, сколько их вообще, и скажет в четверг."),
										},
										Variant: w.ProseLeadAfter,
									},
								},
							},
							w.Docmeta{
								Sections: []w.DocmetaSec{
									{
										Title: w.Rubric{Text: "Tags"},
										Items: []w.DocmetaPart{
											w.Line{Block: true, Gap: w.LineGap4, Items: []w.LinePart{w.Tag{Label: "встречи"}, w.Tag{Label: "релиз"}}},
										},
									},
									{
										Title: w.Rubric{Text: "Related"},
										Items: []w.DocmetaPart{
											w.DocRefs{
												Items: []w.DocRef{
													{Name: "Синк-ап по релизу", What: w.Dim{Parts: []w.Inline{w.Text("· event · 10 Aug")}}},
													{Name: "Дописать раздел про кэш", What: w.Dim{Parts: []w.Inline{w.Text("· task")}}},
												},
											},
										},
									},
									{
										Title: w.Rubric{Text: "Attachments", Items: []w.RubricPart{w.Button{Label: "Attach", Menu: true, Small: true}}},
										Items: []w.DocmetaPart{
											w.DocAttach{
												Name:  "2026-08-10-sync-board.png",
												Size:  w.Dim{Parts: []w.Inline{w.Text("· 412 KB")}},
												Where: w.Badge{Label: "in the note"},
											},
											w.DocAttach{
												Name:  "запись-встречи.m4a",
												Size:  w.Dim{Parts: []w.Inline{w.Text("· 24 MB")}},
												Where: w.Badge{Label: "Documents / carrel"},
											},
										},
									},
									{
										Title: w.Rubric{Text: "Source"},
										Items: []w.DocmetaPart{w.Diag{Text: "UID   9e31-40b7-c012\nRead  2 min ago"}},
									},
								},
								Danger: &w.Danger{Button: w.Button{Label: "Delete note", Kind: w.ButtonDangerSolid}, Space: w.DangerAtEnd},
							},
						},
						View: w.DocViewMeasure,
					},
				},
			},
		},
	}
}

// Notes20 is frame «Заметка · правка» (s-notes-20).
func Notes20() Screen {
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
					{Icon: "back", Label: "Back to list", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "64", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "22"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Neighbours"},
						w.Tree{
							Rows: []w.TreeRow{
								{Icon: "note", Label: "Планы на осень"},
								{Icon: "note", Label: "Заметка со встречи"},
								{Icon: "note", Label: "Что почитать про CRDT"},
							},
							On:      2,
							Variant: w.TreePadded,
						},
					},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Docbar{
						Items: []w.DocbarPart{
							w.Button{Label: "Save", Kind: w.ButtonPrimary},
							w.Button{Label: "Cancel"},
							w.BarSep{},
							w.Button{Label: "B", Bold: true},
							w.Button{Label: "I", Variant: w.ButtonItalic},
							w.Button{Label: "H2", Tertiary: true},
							w.Button{Label: "• list", Tertiary: true},
							w.Button{Folds: true, Label: "task", Icon: "task", Title: "Task", Tertiary: true},
							w.Button{Icon: "link", Title: "Link", Tertiary: true},
							w.Button{Folds: true, Label: "Attach", Menu: true, Icon: "up", Title: "Attach"},
						},
						More: &w.Button{Icon: "dots", Title: "More"},
						End: []w.DocbarEnd{
							w.Check{Parts: []w.CheckPart{w.Text("Auto-save")}, Tertiary: true},
							w.DraftNote{Text: "Draft in this window ·", Time: w.Mono{Text: "11:42"}, Secondary: true},
							w.Kbd{Keys: "Ctrl+S", Secondary: true},
							w.Seg{Options: []w.SegOption{{Label: "Full width", On: true}, {Label: "Reading width"}}, Secondary: true},
							w.Button{Folds: true, Label: "Focus", Icon: "focus", Title: "Focus"},
						},
					},
					w.DocLayout{
						Items: []w.DocLayoutPart{
							w.Doc{
								Items: []w.DocPart{
									w.Area{Text: "Заметка со встречи 10.08", Variant: w.AreaTitle, Label: "Text"},
									w.DocFields{
										Items: []w.Field{
											w.Field{
												Label: "Notebook",
												Parts: []w.FieldPart{w.Pick{Value: "nextcloud · Записи"}},
												Width: w.FieldFlex2_280,
											},
											w.Field{
												Label: "Date",
												Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Value: "2026-08-10", Label: "Date", Mono: true}},
												Width: w.FieldFix150,
											},
											w.Field{
												Label: "Time",
												Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Value: "11:05", Label: "Time", Mono: true}},
												Width: w.FieldFix130,
											},
											w.Field{
												Label: "Tags",
												Parts: []w.FieldPart{w.Input{Value: "встречи, релиз", Label: "Tags"}},
												Width: w.FieldFlex3_320,
											},
											w.Field{
												Label: "Related to",
												Parts: []w.FieldPart{w.TagInput{Parts: []w.TagInputPart{w.Tag{Label: "Синк-ап по релизу", Removable: true}}}},
												Width: w.FieldFlex3_340,
											},
										},
									},
									w.Area{
										Text:    "Договорились заморозить ветку в среду.\n\nПисьмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.\n\n## Что делаем\n\n- [ ] проверить changelog\n- [x] попросить у Артёма скриншоты\n\n## Отложено до следующего спринта\n\nПеренос старых вложений в новую папку — Виктор посмотрит, сколько их вообще.",
										Variant: w.AreaBody,
										Focus:   true,
										Label:   "Text",
									},
									w.DocNote{
										Parts: []w.Inline{
											w.Text("Properties Carrel does not know — X-JTX-COLOR and two more — are kept unchanged and not shown here."),
										},
									},
								},
							},
							w.Docmeta{
								Sections: []w.DocmetaSec{
									{
										Title: w.Rubric{Text: "Attachments", Items: []w.RubricPart{w.Button{Label: "Attach", Menu: true, Small: true}}},
										Items: []w.DocmetaPart{
											w.DocAttach{
												Name:  "2026-08-10-sync-board.png",
												Size:  w.Dim{Parts: []w.Inline{w.Text("· 412 KB")}},
												Where: w.Badge{Label: "in the note"},
											},
											w.DocAttach{
												Name:  "запись-встречи.m4a",
												Size:  w.Dim{Parts: []w.Inline{w.Text("· 24 MB")}},
												Where: w.Badge{Label: "Documents / carrel"},
											},
										},
									},
									{
										Title: w.Rubric{
											Text:  "The same note in other notebooks",
											Items: []w.RubricPart{w.Badge{Label: "1", Tone: w.BadgeLinked}},
										},
										Items: []w.DocmetaPart{
											w.DocRefs{
												Items: []w.DocRef{
													{
														Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11Inline},
														Name: "nextcloud · Дневник",
														Href: "#",
														What: w.Dim{Parts: []w.Inline{w.Text("· Text differ")}},
													},
												},
											},
										},
									},
									{
										Title: w.Rubric{Text: "Source"},
										Items: []w.DocmetaPart{w.Diag{Text: "UID   9e31-40b7-c012\nRead  2 min ago"}},
									},
								},
								Danger: &w.Danger{Button: w.Button{Label: "Delete note", Kind: w.ButtonDangerSolid}, Space: w.DangerAtEnd},
							},
						},
					},
				},
			},
		},
	}
}

// Notes4 is frame «Заметка · правка на весь экран, режим сосредоточения» (s-notes-4).
func Notes4() Screen {
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
			Main: w.Main{
				Items: []w.MainPart{
					w.Docbar{
						Items: []w.DocbarPart{
							w.Button{Label: "Save", Kind: w.ButtonPrimary},
							w.Button{Label: "Cancel"},
							w.BarSep{},
							w.Button{Label: "B", Bold: true},
							w.Button{Label: "I", Variant: w.ButtonItalic},
							w.Button{Label: "H2"},
							w.Button{Label: "• list"},
							w.Button{Folds: true, Label: "task", Icon: "task", Title: "Task"},
							w.Button{Icon: "link", Title: "Link"},
							w.Button{Folds: true, Label: "Attach", Menu: true, Icon: "up", Title: "Attach"},
						},
						More: &w.Button{Icon: "dots", Title: "More"},
						End: []w.DocbarEnd{
							w.Check{Parts: []w.CheckPart{w.Text("Auto-save")}},
							w.DraftNote{Text: "Draft in this window ·", Time: w.Mono{Text: "11:42"}, Secondary: true},
							w.Kbd{Keys: "Ctrl+S", Secondary: true},
							w.Seg{Options: []w.SegOption{{Label: "Full width", On: true}, {Label: "Reading width"}}, Secondary: true},
							w.Button{Folds: true, Label: "Focus", Icon: "focus", Variant: w.ButtonOnCard, Title: "Focus"},
						},
					},
					w.DocLayout{
						Items: []w.DocLayoutPart{
							w.Doc{
								Items: []w.DocPart{
									w.Area{Text: "Заметка со встречи 10.08", Variant: w.AreaTitle, Focus: true, Label: "Text"},
									w.DocFields{
										Items: []w.Field{
											w.Field{
												Label: "Notebook",
												Parts: []w.FieldPart{w.Pick{Value: "nextcloud · Записи"}},
												Width: w.FieldFlex2_280,
											},
											w.Field{
												Label: "Date",
												Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Value: "2026-08-10", Label: "Date", Mono: true}},
												Width: w.FieldFix150,
											},
											w.Field{
												Label: "Time",
												Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Value: "11:05", Label: "Time", Mono: true}},
												Width: w.FieldFix130,
											},
											w.Field{
												Label: "Tags",
												Parts: []w.FieldPart{w.Input{Value: "встречи, релиз", Label: "Tags"}},
												Width: w.FieldFlex3_320,
											},
											w.Field{
												Label: "Related to",
												Parts: []w.FieldPart{
													w.Drop{
														Open: true,
														Trigger: w.TagInput{
															Parts: []w.TagInputPart{w.Tag{Label: "Синк-ап по релизу", Removable: true}, w.Dim{Parts: []w.Inline{w.Text("ковал")}}},
															Focus: true,
														},
														Menu: w.Menu{
															Items: []w.MenuPart{
																w.MenuItem{
																	Icon:  "cal",
																	Trail: w.Mono{Text: "10 Aug · Личный", Variant: w.Mono10Dusty, Place: w.Pushed},
																	Label: "Звонок с Ковалёвой",
																	On:    true,
																},
																w.MenuItem{
																	Icon:  "task",
																	Trail: w.Mono{Text: "task · Работа", Variant: w.Mono10Dusty, Place: w.Pushed},
																	Label: "Ответить Анне по смете",
																},
																w.MenuItem{
																	Icon:  "con",
																	Trail: w.Mono{Text: "contact", Variant: w.Mono10Dusty, Place: w.Pushed},
																	Label: "Анна Ковалёва",
																},
																w.MenuSep{},
																w.MenuItem{Label: "Enter a UID by hand", Tone: w.MenuToneDusty},
															},
															Width: w.MenuWidthFull,
															Pop:   w.MenuUnder,
														},
													},
												},
												Width: w.FieldFlex3_340,
											},
										},
									},
									w.Area{
										Text:    "Договорились заморозить ветку в среду.\n\nПисьмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.\n\n## Что делаем\n\n- [ ] проверить changelog\n- [x] попросить у Артёма скриншоты\n\n## Отложено до следующего спринта\n\nПеренос старых вложений в новую папку — Виктор посмотрит, сколько их вообще.",
										Variant: w.AreaBody,
										Label:   "Text",
									},
									w.DocNote{
										Parts: []w.Inline{
											w.Text("Properties Carrel does not know — "),
											w.Mono{Text: "X-JTX-COLOR"},
											w.Text(" and two more — are kept unchanged and not shown here."),
										},
									},
								},
							},
						},
						View: w.DocViewFocus,
					},
				},
			},
		},
	}
}

// Notes5 is frame «Быстрая заметка поверх текущего экрана» (s-notes-5).
func Notes5() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}},
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
				},
				Main: w.Main{
					Items: []w.MainPart{
						w.Head{
							Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined calendar"},
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("5 of 7 calendars · 10 – 16 August 2026")}}},
							Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}}},
						},
						w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
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
									Mark:   w.Mark{Colour: "#8C6239"},
									Time:   "all day",
									Title:  []w.Inline{w.Text("Дежурство: первая линия")},
									Detail: "—",
									Source: "2 sources",
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
						w.ListGroup{Title: w.Rubric{Text: "Tuesday · 11 August", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
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
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Icon:  "quill",
					Title: "New note",
					Head: []w.SheetHeadPart{
						w.Pick{Value: "nextcloud · Записи", Variant: w.PickT12W200, Place: w.Pushed},
						w.Close{Label: "Close"},
					},
					Body: []w.SheetPart{
						w.Area{
							Text:    "Позвонить в МФЦ во вторник до 12:00 — уточнить, нужен ли оригинал доверенности.",
							Variant: w.AreaBodySm,
							Focus:   true,
							Label:   "Text",
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Save note", Kind: w.ButtonPrimary},
						w.Kbd{Keys: "Ctrl+Enter"},
						w.Button{Label: "Open the card", Icon: "expand"},
						w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}},
					},
				},
			},
		},
	}
}

// Notes6 is frame «Быстрая заметка · писать некуда» (s-notes-6).
func Notes6() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}},
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
							{Icon: "con", Label: "Contacts", Href: "#"},
							{Icon: "cal", Label: "Calendar", Href: "#"},
							{Icon: "task", Label: "Tasks", Href: "#"},
							{Icon: "note", Label: "Notes", Href: "#"},
							{Icon: "file", Label: "Files", Href: "#"},
							{Icon: "dup", Label: "Duplicates", Href: "#"},
						},
						On: 2,
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.TickRow{
								Ticked: w.TickedOff,
								Kind:   w.TickRowAll,
								Mark:   w.Mark{Kind: w.MarkMulti},
								Label:  "Combined calendar",
								Meta:   "—",
							},
						},
					},
					w.Empty{Parts: []w.EmptyPart{w.Text("No account connected.")}},
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
						w.Empty{Parts: []w.EmptyPart{w.Text("Connect an account to see calendars, address books and notebooks.")}},
					},
				},
				Height: w.ShellHeight340,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Icon:  "quill",
					Title: "New note",
					Head:  []w.SheetHeadPart{w.Close{Label: "Close", Place: w.Pushed}},
					Body: []w.SheetPart{
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("There is nowhere to write. Connect an account that has a journal, or create a notebook after connecting — the note is not kept in this window."),
							},
							Space: w.MsgFlush,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Connect account", Icon: "server", Kind: w.ButtonPrimary},
						w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}},
					},
					Width: w.SheetWidth480,
				},
			},
		},
	}
}

// Notes7 is frame «Стрелка у кнопки» (s-notes-7).
func Notes7() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.TopDrop{
					Open:    true,
					Trigger: w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary, Focus: true},
					Menu: w.Menu{
						Items: []w.MenuPart{
							w.MenuItem{Icon: "quill", Trail: w.Kbd{Keys: "N"}, Label: "Note", On: true},
							w.MenuSep{},
							w.MenuItem{Icon: "con", Trail: w.Kbd{Keys: "C"}, Label: "Contact"},
							w.MenuItem{Icon: "cal", Trail: w.Kbd{Keys: "E"}, Label: "Event"},
							w.MenuItem{Icon: "task", Trail: w.Kbd{Keys: "T"}, Label: "Task"},
							w.MenuSep{},
							w.MenuItem{Icon: "up", Label: "Upload file"},
							w.MenuSep{},
							w.MenuItem{Icon: "mail", Label: "Letter"},
						},
					},
				},
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
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
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
						Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}}},
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
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
								Mark:   w.Mark{Colour: "#3F6E8C"},
								Time:   "13:30",
								Title:  []w.Inline{w.Text("Звонок с Ковалёвой")},
								Detail: "Zoom",
								Source: "nextcloud · Работа",
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#6B4E7A"},
								Time:   "18:00",
								Title:  []w.Inline{w.Text("Забрать документы")},
								Detail: "МФЦ на Ленина",
								Source: "radicale · vps · Проекты",
							},
						},
					},
				},
			},
			Height: w.ShellHeight340,
		},
	}
}

// Notes8 is frame «Пункт «Contact» — лист поверх экрана» (s-notes-8).
func Notes8() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}},
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
							w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
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
							Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}}},
						},
						w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
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
									Mark:   w.Mark{Colour: "#3F6E8C"},
									Time:   "13:30",
									Title:  []w.Inline{w.Text("Звонок с Ковалёвой")},
									Detail: "Zoom",
									Source: "nextcloud · Работа",
								},
								w.TimeRow{
									Mark:   w.Mark{Colour: "#6B4E7A"},
									Time:   "18:00",
									Title:  []w.Inline{w.Text("Забрать документы")},
									Detail: "МФЦ на Ленина",
									Source: "radicale · vps · Проекты",
								},
							},
						},
					},
				},
				Height: w.ShellHeight340,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Icon:  "con",
					Title: "New contact",
					Head: []w.SheetHeadPart{
						w.Pick{Value: "nextcloud · Контакты", Variant: w.PickT12W200, Place: w.Pushed},
						w.Close{Label: "Close"},
					},
					Body: []w.SheetPart{
						w.SheetFields{
							Items: []w.Field{
								w.Field{
									Label: "Display name",
									Parts: []w.FieldPart{w.Input{Value: "Пётр Ковалёв", Label: "Display name", Focus: true}},
									Place: w.Grown,
								},
								w.Field{
									Label: "Phone",
									Parts: []w.FieldPart{w.Input{Placeholder: "+7 …", Label: "Phone", Mono: true}},
									Width: w.FieldFill200,
								},
								w.Field{
									Label: "Email",
									Parts: []w.FieldPart{w.Input{Placeholder: "name@example.org", Label: "Email"}},
									Width: w.FieldFill200,
								},
							},
						},
						w.Dim{
							Parts: []w.Inline{
								w.Text("The first word is Given, the rest Family — «Пётр Ковалёв», not «Ковалёв Пётр». Everything else is filled in on the card."),
							},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Save contact", Kind: w.ButtonPrimary},
						w.Kbd{Keys: "Enter"},
						w.Button{Label: "Open the card", Icon: "expand"},
						w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}},
					},
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Notes9 is frame «Пункт «Event»» (s-notes-9).
func Notes9() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}},
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
					Icon:  "cal",
					Title: "New event",
					Head: []w.SheetHeadPart{
						w.Pick{Value: "nextcloud · Личный", Variant: w.PickT12W200, Place: w.Pushed},
						w.Close{Label: "Close"},
					},
					Body: []w.SheetPart{
						w.SheetFields{
							Items: []w.Field{
								w.Field{
									Label: "Title",
									Parts: []w.FieldPart{w.Input{Value: "Забрать документы", Label: "Title", Focus: true}},
									Place: w.Grown,
								},
								w.Field{
									Label: "Date",
									Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Value: "2026-08-11", Label: "Date", Mono: true}},
									Width: w.FieldFix130,
								},
								w.Field{
									Label: "Start",
									Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Value: "18:00", Label: "Start", Mono: true}},
									Width: w.FieldFix96,
								},
								w.Field{
									Label: "End",
									Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Value: "18:40", Label: "End", Mono: true}},
									Width: w.FieldFix96,
								},
								w.Field{
									Label: "\u00a0",
									Parts: []w.FieldPart{w.Check{Parts: []w.CheckPart{w.Text("All day")}}},
									Width: w.FieldFill120,
								},
								w.Field{Label: "Remind", Parts: []w.FieldPart{w.Pick{Value: "10 minutes before"}}, Width: w.FieldFill200},
							},
						},
						w.Dim{
							Parts: []w.Inline{
								w.Text("Repeats, guests and more live on the card. A reminder is set here because that is why most events are created."),
							},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Save event", Kind: w.ButtonPrimary},
						w.Kbd{Keys: "Enter"},
						w.Button{Label: "Open the card", Icon: "expand"},
						w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}},
					},
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Notes10 is frame «Пункт «Task»» (s-notes-10).
func Notes10() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}},
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
					Icon:  "task",
					Title: "New task",
					Head:  []w.SheetHeadPart{w.Pick{Value: "nextcloud · Дела", Variant: w.PickT12W200, Place: w.Pushed}, w.Close{Label: "Close"}},
					Body: []w.SheetPart{
						w.SheetFields{
							Items: []w.Field{
								w.Field{
									Label: "Summary",
									Parts: []w.FieldPart{w.Input{Value: "Отправить документы в МФЦ", Label: "Summary", Focus: true}},
									Place: w.Grown,
								},
								w.Field{
									Label: "Due date",
									Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Placeholder: "YYYY-MM-DD", Label: "Due date", Mono: true}},
									Width: w.FieldFix130,
								},
								w.Field{
									Label: "Due time",
									Parts: []w.FieldPart{w.Input{Placeholder: "--:--", Label: "Due time", Mono: true}},
									Width: w.FieldFix96,
								},
							},
						},
						w.Dim{
							Parts: []w.Inline{
								w.Text("No time means the whole day. A task with no due date is a task all the same — both may stay empty."),
							},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Save task", Kind: w.ButtonPrimary},
						w.Kbd{Keys: "Enter"},
						w.Button{Label: "Open the card", Icon: "expand"},
						w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}},
					},
					Width: w.SheetWidth520,
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Notes11 is frame «Пункт «Upload file» — куда именно» (s-notes-11).
func Notes11() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}},
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
							w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
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
							Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}}},
						},
						w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
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
									Mark:   w.Mark{Colour: "#3F6E8C"},
									Time:   "13:30",
									Title:  []w.Inline{w.Text("Звонок с Ковалёвой")},
									Detail: "Zoom",
									Source: "nextcloud · Работа",
								},
								w.TimeRow{
									Mark:   w.Mark{Colour: "#6B4E7A"},
									Time:   "18:00",
									Title:  []w.Inline{w.Text("Забрать документы")},
									Detail: "МФЦ на Ленина",
									Source: "radicale · vps · Проекты",
								},
							},
						},
					},
				},
				Height: w.ShellHeight560,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Icon:  "up",
					Title: "Upload files",
					Head: []w.SheetHeadPart{
						w.Pick{Value: "Documents / carrel / 2026-08", Variant: w.PickT12W240, Place: w.Pushed},
						w.Close{Label: "Close"},
					},
					Body: []w.SheetPart{
						w.Empty{
							Parts: []w.EmptyPart{
								w.Text("Drop files here, or "),
								w.Emphasis("choose on the computer"),
								w.Text(" "),
								w.Aside{Face: w.FaceSmall, Parts: []w.Inline{w.Text("A name already taken is refused, not overwritten")}},
							},
							Variant: w.EmptyDashed,
						},
						w.Card{
							Items: []w.CardPart{
								w.Rubric{
									Text:  "Where to put it",
									Items: []w.RubricPart{w.Button{Label: "New folder", Icon: "plus", Small: true}},
								},
								w.Tree{
									Rows: []w.TreeRow{
										{Icon: "file", Label: "Documents", Trail: w.Dim{Parts: []w.Inline{w.Text("sftpgo")}, Size: w.Dim11}},
										{Icon: "file", Label: "carrel", Indent: w.TreeLevel22},
										{
											Icon:   "file",
											Label:  "attachments",
											Indent: w.TreeLevel40,
											Trail:  w.Badge{Label: "attachment folder"},
										},
										{Icon: "file", Label: "2026-08", Indent: w.TreeLevel40, Trail: w.Badge{Label: "just created"}},
										{Icon: "file", Label: "2026-07", Indent: w.TreeLevel40},
										{Icon: "file", Label: "Files", Trail: w.Dim{Parts: []w.Inline{w.Text("nextcloud")}, Size: w.Dim11}},
										{Icon: "file", Label: "Public", Unavailable: true, Trail: w.Badge{Label: "read-only"}},
									},
									On: 4,
								},
							},
							Variant: w.CardFlush,
						},
						w.Dim{
							Parts: []w.Inline{
								w.Text("The folder is the attachment folder from Settings → Attachments, named once and remembered — unless you make a new one here. Change it here and the next upload from this menu goes to the new one."),
							},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Upload 2 files", Kind: w.ButtonPrimary},
						w.Kbd{Keys: "Enter"},
						w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}},
					},
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}
