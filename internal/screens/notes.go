// Раздел макета: 7.4 Заметки.

package screens

import w "carrel.dev/typecase/widget"

// Notes1 is frame «Заметки · список и заметка в панели» (s-notes-1).
func Notes1() Screen {
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
					On: 4,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedPart,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined notebook",
							Meta:    "3",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "1"},
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
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 of 3 notebooks · 3 notes · VJOURNAL")}}},
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
							w.Filter{
								Query:   w.Query{Placeholder: "Filter by word", Label: "Filter by word"},
								Waiting: true,
								Variant: w.FilterH26W200,
							},
							w.Seg{Options: []w.SegOption{{Label: "Newest", On: true}, {Label: "Oldest"}, {Label: "Title"}}},
							w.Button{Label: "Columns", Icon: "cols"},
						},
						Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "3"}}, Secondary: true},
					},
					w.ListGroup{Title: w.Rubric{Text: "August 2026", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
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
						Tools: []w.HeadTool{
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
									{Tick: w.Tick{Variant: w.TickNudge, Sign: true}, Text: "проверить changelog"},
									{Tick: w.Tick{Done: true, Variant: w.TickNudge, Sign: true}, Text: "попросить у Артёма скриншоты"},
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
										w.Link{Text: "2026-08-10-sync-board.png", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("image/png · 412 KB")}},
										w.Text(" "),
										w.Badge{Label: "in the note"},
									},
									[]w.SecItemPart{
										w.Link{Text: "запись-встречи.m4a", Href: "#"},
										w.Text(" "),
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
									[]w.SecItemPart{
										w.Span{Face: w.FaceAccent, Text: "Дописать раздел про кэш"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· task")}},
									},
								},
								Gap: 6,
							},
						},
					},
				},
			},
		},
	}
}

// Notes31 is frame «Заметки · одна тетрадь» (s-notes-31).
func Notes31() Screen {
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
					On: 4,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked: w.TickedPart,
							Kind:   w.TickRowAll,
							Mark:   w.Mark{Kind: w.MarkMulti},
							Label:  "Combined notebook",
							Meta:   "3",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "1"},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH20}, Text: "Записи"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · 2 notes · VJOURNAL")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New note", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "up"},
								w.Button{Label: "Export", Icon: "down"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter by word", Label: "Filter by word"},
								Waiting: true,
								Variant: w.FilterH26W200,
							},
							w.Seg{Options: []w.SegOption{{Label: "Newest", On: true}, {Label: "Oldest"}, {Label: "Title"}}},
							w.Button{Label: "Columns", Icon: "cols"},
						},
						Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "2"}}, Secondary: true},
					},
					w.ListGroup{Title: w.Rubric{Text: "August 2026", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Title:   []w.Inline{w.Text("Заметка со встречи 10.08")},
								Excerpt: "Договорились заморозить ветку в среду; письмо пользователям пишет Анна…",
								Tags:    []w.RowTag{w.Tag{Label: "встречи"}},
								Date:    "10 Aug",
								Single:  true,
							},
							w.ExcerptRow{
								Mark:       w.Mark{Colour: "#4A6B52"},
								Title:      []w.Inline{w.Text("Что почитать про CRDT ")},
								TitleBadge: &w.Badge{Label: "linked ×2", Tone: w.BadgeLinked, Title: "nextcloud · Записи, nextcloud · Дневник"},
								Excerpt:    "Кляйн, глава 5; статья Шапиро 2011; демка automerge…",
								Tags:       []w.RowTag{w.Tag{Label: "книги"}},
								Date:       "8 Aug",
								Single:     true,
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "1"},
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
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
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
								w.Button{Label: "Source", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012?view=source", Swap: "note-main"},
								w.Seg{
									Options: []w.SegOption{
										{Label: "Full width", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012?width=full", On: true},
										{Label: "Reading width", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012?width=read"},
									},
									Swap: "note-main",
								},
								w.Button{
									Icon:  "panel",
									Title: "Hide the side column",
									Href:  "/app/notes/nextcloud/zapisi/9e31-40b7-c012?side=0",
									Swap:  "note-main",
								},
								w.Drop{
									Trigger: w.Button{Icon: "dots", Title: "More"},
									Menu: w.Menu{
										Items: []w.MenuPart{
											w.MenuItem{Icon: "move", Label: "Move", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012/move"},
											w.MenuItem{Icon: "dup", Label: "Copy", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012/copy"},
											w.MenuItem{Icon: "link", Label: "Publish a link", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012/publish"},
										},
										Width: w.MenuWidth212,
										Pop:   w.MenuUnderEnd,
									},
								},
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
											{Tick: w.Tick{Variant: w.TickNudge, Sign: true}, Text: "проверить changelog"},
											{Tick: w.Tick{Done: true, Variant: w.TickNudge, Sign: true}, Text: "попросить у Артёма скриншоты"},
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
										Title: w.Rubric{
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
										Items: []w.DocmetaPart{
											w.DocAttach{
												Name:  "2026-08-10-sync-board.png",
												Size:  w.Dim{Parts: []w.Inline{w.Text("· 412 KB")}},
												Where: w.Badge{Label: "in the note"},
												Href:  "#",
											},
											w.DocAttach{
												Name:  "запись-встречи.m4a",
												Size:  w.Dim{Parts: []w.Inline{w.Text("· 24 MB")}},
												Where: w.Badge{Label: "Documents / carrel"},
												Href:  "#",
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
				ID: "note-main",
			},
		},
	}
}

// Notes27 is frame «Заметка · исходник» (s-notes-27).
func Notes27() Screen {
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "1"},
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
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
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
								w.Button{Label: "Rendered", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012", Swap: "note-main"},
								w.Seg{
									Options: []w.SegOption{
										{Label: "Full width", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012?width=full", On: true},
										{Label: "Reading width", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012?width=read"},
									},
									Swap: "note-main",
								},
								w.Button{
									Icon:  "panel",
									Title: "Hide the side column",
									Href:  "/app/notes/nextcloud/zapisi/9e31-40b7-c012?side=0",
									Swap:  "note-main",
								},
								w.Drop{
									Trigger: w.Button{Icon: "dots", Title: "More"},
									Menu: w.Menu{
										Items: []w.MenuPart{
											w.MenuItem{Icon: "move", Label: "Move", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012/move"},
											w.MenuItem{Icon: "dup", Label: "Copy", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012/copy"},
											w.MenuItem{Icon: "link", Label: "Publish a link", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012/publish"},
										},
										Width: w.MenuWidth212,
										Pop:   w.MenuUnderEnd,
									},
								},
							},
						},
						Pad: w.HeadRule,
					},
					w.DocLayout{
						Items: []w.DocLayoutPart{
							w.Doc{
								Items: []w.DocPart{
									w.DocTitle{Text: "Заметка со встречи 10.08"},
									w.DocCode{
										Text: "Договорились заморозить ветку в среду.\n\nПисьмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.\n\n## Что делаем\n\n- [ ] проверить changelog\n- [x] попросить у Артёма скриншоты\n\nПеренос старых вложений в новую папку — Виктор посмотрит, сколько их вообще, и скажет в четверг.",
										Wrap: true,
									},
								},
								Text: true,
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
										Title: w.Rubric{
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
										Items: []w.DocmetaPart{
											w.DocAttach{
												Name:  "2026-08-10-sync-board.png",
												Size:  w.Dim{Parts: []w.Inline{w.Text("· 412 KB")}},
												Where: w.Badge{Label: "in the note"},
												Href:  "#",
											},
											w.DocAttach{
												Name:  "запись-встречи.m4a",
												Size:  w.Dim{Parts: []w.Inline{w.Text("· 24 MB")}},
												Where: w.Badge{Label: "Documents / carrel"},
												Href:  "#",
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
				ID: "note-main",
			},
		},
	}
}

// Notes3 is frame «Заметка · чтение, Reading width» (s-notes-3).
func Notes3() Screen {
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "1"},
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
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
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
								w.Button{Label: "Source", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012?view=source", Swap: "note-main"},
								w.Seg{
									Options: []w.SegOption{
										{Label: "Full width", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012?width=full"},
										{Label: "Reading width", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012?width=read", On: true},
									},
									Swap: "note-main",
								},
								w.Button{
									Icon:  "panel",
									Title: "Hide the side column",
									Href:  "/app/notes/nextcloud/zapisi/9e31-40b7-c012?side=0",
									Swap:  "note-main",
								},
								w.Drop{
									Trigger: w.Button{Icon: "dots", Title: "More"},
									Menu: w.Menu{
										Items: []w.MenuPart{
											w.MenuItem{Icon: "move", Label: "Move", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012/move"},
											w.MenuItem{Icon: "dup", Label: "Copy", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012/copy"},
											w.MenuItem{Icon: "link", Label: "Publish a link", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012/publish"},
										},
										Width: w.MenuWidth212,
										Pop:   w.MenuUnderEnd,
									},
								},
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
											{Tick: w.Tick{Variant: w.TickNudge, Sign: true}, Text: "проверить changelog"},
											{Tick: w.Tick{Done: true, Variant: w.TickNudge, Sign: true}, Text: "попросить у Артёма скриншоты"},
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
										Title: w.Rubric{
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
										Items: []w.DocmetaPart{
											w.DocAttach{
												Name:  "2026-08-10-sync-board.png",
												Size:  w.Dim{Parts: []w.Inline{w.Text("· 412 KB")}},
												Where: w.Badge{Label: "in the note"},
												Href:  "#",
											},
											w.DocAttach{
												Name:  "запись-встречи.m4a",
												Size:  w.Dim{Parts: []w.Inline{w.Text("· 24 MB")}},
												Where: w.Badge{Label: "Documents / carrel"},
												Href:  "#",
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
						View: w.DocViewMeasure,
					},
				},
				ID: "note-main",
			},
		},
	}
}

// Notes20 is frame «Заметка · правка» (s-notes-20).
func Notes20() Screen {
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "1"},
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
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Docbar{
						Items: []w.DocbarPart{
							w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "note-edit"},
							w.Button{Label: "Cancel", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012"},
							w.BarSep{},
							w.Button{Label: "B", Bold: true},
							w.Button{Label: "I", Variant: w.ButtonItalic},
							w.Button{Label: "H2", Tertiary: true},
							w.Button{Label: "• list", Tertiary: true},
							w.Button{Folds: true, Label: "task", Icon: "task", Title: "Task", Tertiary: true},
							w.Button{Icon: "link", Title: "Link", Tertiary: true},
							w.Drop{
								Trigger: w.Button{Folds: true, Label: "Attach", Menu: true, Icon: "up", Title: "Attach"},
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuItem{
											Icon: "up",
											File: &w.FileField{Form: "note-edit", Name: "file", Label: "From this computer", List: "note-attachments"},
											Row: w.DocAttach{
												Name:   "{name}",
												Size:   w.Dim{Parts: []w.Inline{w.Text("· {size}")}},
												Remove: &w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Removes: true},
											},
											Label: "From this computer",
										},
										w.MenuItem{Icon: "file", Label: "From Files"},
									},
									Width: w.MenuWidth212,
									Pop:   w.MenuUnder,
								},
							},
						},
						More: &w.Button{Icon: "dots", Title: "More"},
						End: []w.DocbarEnd{
							w.Check{Parts: []w.CheckPart{w.Text("Auto-save")}, Tertiary: true},
							w.DraftNote{Text: "Draft in this window ·", Time: w.Mono{Text: "11:42"}, Secondary: true},
							w.Kbd{Keys: "Ctrl+S", Secondary: true},
							w.Seg{Options: []w.SegOption{{Label: "Full width", On: true}, {Label: "Reading width"}}, Secondary: true},
							w.Button{Folds: true, Label: "Focus", Icon: "focus", Title: "Focus"},
						},
						Edit: true,
					},
					w.Runbar{
						Meter:    w.Meter{Percent: 0},
						Progress: "note-edit",
						Parts:    []w.RunbarPart{w.Text("Saving · "), w.Mono{Text: "0 of 0 MB"}},
					},
					w.DocLayout{
						Items: []w.DocLayoutPart{
							w.Doc{
								Items: []w.DocPart{
									w.Area{Text: "Заметка со встречи 10.08", Variant: w.AreaTitle, Name: "summary", Label: "Text"},
									w.DocFields{
										Items: []w.Field{
											w.Field{
												Label: "Notebook",
												Parts: []w.FieldPart{
													w.Drop{
														Trigger: w.Pick{Value: "nextcloud · Записи", Mark: &w.Mark{Colour: "#4A6B52"}},
														Menu: w.Menu{
															Items: []w.MenuPart{
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#4A6B52"},
																	Label: "nextcloud · Записи",
																	On:    true,
																	Name:  "notebook",
																	Value: "nextcloud/zapisi",
																},
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#6B4E7A"},
																	Label: "nextcloud · Дневник",
																	Name:  "notebook",
																	Value: "nextcloud/dnevnik",
																},
																w.MenuItem{
																	Mark:   &w.Mark{Colour: "#3F6E8C"},
																	Trail:  w.Badge{Label: "ro", Place: w.Pushed},
																	Label:  "nextcloud · Общие с семьёй",
																	Reason: "Read-only: this account cannot write here",
																	Name:   "notebook",
																	Value:  "nextcloud/semya",
																},
															},
															Width: w.MenuWidth290,
															Pop:   w.MenuUnder,
														},
													},
												},
												Width: w.FieldFlex2_280,
											},
											w.Field{
												Label: "Date",
												Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "dtstart.date", Value: "2026-08-10", Label: "Date", Mono: true}},
												Width: w.FieldFix150,
											},
											w.Field{
												Label: "Time",
												Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Name: "dtstart.time", Value: "11:05", Label: "Time", Mono: true}},
												Width: w.FieldFix130,
											},
											w.Field{
												Label: "Tags",
												Parts: []w.FieldPart{w.Input{Name: "categories", Value: "встречи, релиз", Label: "Tags"}},
												Width: w.FieldFlex3_320,
											},
											w.Field{
												Label: "Related to",
												Parts: []w.FieldPart{
													w.TagInput{
														Parts: []w.TagInputPart{w.Tag{Label: "Синк-ап по релизу", Removable: true}},
														Add:   &w.TagAdd{Name: "related.add", Placeholder: "Title or UID", Label: "Add a link"},
													},
												},
												Width: w.FieldFlex3_340,
											},
										},
									},
									w.Area{
										Text:    "Договорились заморозить ветку в среду.\n\nПисьмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.\n\n## Что делаем\n\n- [ ] проверить changelog\n- [x] попросить у Артёма скриншоты\n\nПеренос старых вложений в новую папку — Виктор посмотрит, сколько их вообще, и скажет в четверг.",
										Variant: w.AreaBody,
										Focus:   true,
										Name:    "description",
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
								ID:     "note-edit",
								Action: "/app/notes/nextcloud/zapisi/9e31-40b7-c012",
							},
							w.Docmeta{
								Sections: []w.DocmetaSec{
									{
										Title: w.Rubric{
											Text: "Attachments",
											Items: []w.RubricPart{
												w.Drop{
													Trigger: w.Button{Label: "Attach", Menu: true, Small: true},
													Menu: w.Menu{
														Items: []w.MenuPart{
															w.MenuItem{
																Icon: "up",
																File: &w.FileField{Form: "note-edit", Name: "file", Label: "From this computer", List: "note-attachments"},
																Row: w.DocAttach{
																	Name:   "{name}",
																	Size:   w.Dim{Parts: []w.Inline{w.Text("· {size}")}},
																	Remove: &w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Removes: true},
																},
																Label: "From this computer",
															},
															w.MenuItem{Icon: "file", Label: "From Files"},
														},
														Width: w.MenuWidth212,
														Pop:   w.MenuUnder,
													},
												},
											},
										},
										Items: []w.DocmetaPart{
											w.DocAttach{
												Name:   "2026-08-10-sync-board.png",
												Size:   w.Dim{Parts: []w.Inline{w.Text("· 412 KB")}},
												Where:  w.Badge{Label: "in the note"},
												Remove: &w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true},
												Href:   "#",
											},
											w.DocAttach{
												Name:   "запись-встречи.m4a",
												Size:   w.Dim{Parts: []w.Inline{w.Text("· 24 MB")}},
												Where:  w.Badge{Label: "Documents / carrel"},
												Remove: &w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true},
												Href:   "#",
											},
										},
										ID: "note-attachments",
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

// Notes30 is frame «Заметка · новая» (s-notes-30).
func Notes30() Screen {
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
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "1"},
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
					w.Docbar{
						Items: []w.DocbarPart{
							w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "note-new"},
							w.Button{Label: "Cancel", Href: "/app/notes/nextcloud/zapisi"},
							w.BarSep{},
							w.Button{Label: "B", Bold: true},
							w.Button{Label: "I", Variant: w.ButtonItalic},
							w.Button{Label: "H2", Tertiary: true},
							w.Button{Label: "• list", Tertiary: true},
							w.Button{Folds: true, Label: "task", Icon: "task", Title: "Task", Tertiary: true},
							w.Button{Icon: "link", Title: "Link", Tertiary: true},
							w.Drop{
								Trigger: w.Button{Folds: true, Label: "Attach", Menu: true, Icon: "up", Title: "Attach"},
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuItem{
											Icon: "up",
											File: &w.FileField{Form: "note-new", Name: "file", Label: "From this computer", List: "note-new-attachments"},
											Row: w.DocAttach{
												Name:   "{name}",
												Size:   w.Dim{Parts: []w.Inline{w.Text("· {size}")}},
												Remove: &w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Removes: true},
											},
											Label: "From this computer",
										},
										w.MenuItem{Icon: "file", Label: "From Files"},
									},
									Width: w.MenuWidth212,
									Pop:   w.MenuUnder,
								},
							},
						},
						More: &w.Button{Icon: "dots", Title: "More"},
						End: []w.DocbarEnd{
							w.Check{Parts: []w.CheckPart{w.Text("Auto-save")}, Tertiary: true},
							w.Kbd{Keys: "Ctrl+S", Secondary: true},
							w.Seg{Options: []w.SegOption{{Label: "Full width", On: true}, {Label: "Reading width"}}, Secondary: true},
							w.Button{Folds: true, Label: "Focus", Icon: "focus", Title: "Focus"},
						},
						Edit: true,
					},
					w.Runbar{
						Meter:    w.Meter{Percent: 0},
						Progress: "note-new",
						Parts:    []w.RunbarPart{w.Text("Saving · "), w.Mono{Text: "0 of 0 MB"}},
					},
					w.DocLayout{
						Items: []w.DocLayoutPart{
							w.Doc{
								Items: []w.DocPart{
									w.Area{Text: "", Variant: w.AreaTitle, Focus: true, Name: "summary", Label: "Text"},
									w.DocFields{
										Items: []w.Field{
											w.Field{
												Label: "Notebook",
												Parts: []w.FieldPart{
													w.Drop{
														Trigger: w.Pick{Value: "nextcloud · Записи", Mark: &w.Mark{Colour: "#4A6B52"}},
														Menu: w.Menu{
															Items: []w.MenuPart{
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#4A6B52"},
																	Label: "nextcloud · Записи",
																	On:    true,
																	Name:  "notebook",
																	Value: "nextcloud/zapisi",
																},
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#6B4E7A"},
																	Label: "nextcloud · Дневник",
																	Name:  "notebook",
																	Value: "nextcloud/dnevnik",
																},
																w.MenuItem{
																	Mark:   &w.Mark{Colour: "#3F6E8C"},
																	Trail:  w.Badge{Label: "ro", Place: w.Pushed},
																	Label:  "nextcloud · Общие с семьёй",
																	Reason: "Read-only: this account cannot write here",
																	Name:   "notebook",
																	Value:  "nextcloud/semya",
																},
															},
															Width: w.MenuWidth290,
															Pop:   w.MenuUnder,
														},
													},
												},
												Width: w.FieldFlex2_280,
											},
											w.Field{
												Label: "Date",
												Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "dtstart.date", Value: "2026-08-10", Label: "Date", Mono: true}},
												Width: w.FieldFix150,
											},
											w.Field{
												Label: "Time",
												Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Name: "dtstart.time", Value: "11:05", Label: "Time", Mono: true}},
												Width: w.FieldFix130,
											},
											w.Field{
												Label: "Tags",
												Parts: []w.FieldPart{w.Input{Name: "categories", Placeholder: "—", Label: "Tags"}},
												Width: w.FieldFlex3_320,
											},
											w.Field{
												Label: "Related to",
												Parts: []w.FieldPart{
													w.TagInput{
														Parts: []w.TagInputPart{},
														Add:   &w.TagAdd{Name: "related.add", Placeholder: "Title or UID", Label: "Add a link"},
													},
												},
												Width: w.FieldFlex3_340,
											},
										},
									},
									w.Area{Text: "", Variant: w.AreaBody, Name: "description", Label: "Text"},
								},
								ID:     "note-new",
								Action: "/app/notes/nextcloud/zapisi/new",
							},
							w.Docmeta{
								Sections: []w.DocmetaSec{
									{
										Title: w.Rubric{
											Text: "Attachments",
											Items: []w.RubricPart{
												w.Drop{
													Trigger: w.Button{Label: "Attach", Menu: true, Small: true},
													Menu: w.Menu{
														Items: []w.MenuPart{
															w.MenuItem{
																Icon: "up",
																File: &w.FileField{Form: "note-new", Name: "file", Label: "From this computer", List: "note-new-attachments"},
																Row: w.DocAttach{
																	Name:   "{name}",
																	Size:   w.Dim{Parts: []w.Inline{w.Text("· {size}")}},
																	Remove: &w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Removes: true},
																},
																Label: "From this computer",
															},
															w.MenuItem{Icon: "file", Label: "From Files"},
														},
														Width: w.MenuWidth212,
														Pop:   w.MenuUnder,
													},
												},
											},
										},
										Items: []w.DocmetaPart{},
										ID:    "note-new-attachments",
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

// Notes4 is frame «Заметка · правка на весь экран, режим сосредоточения» (s-notes-4).
func Notes4() Screen {
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
			Main: w.Main{
				Items: []w.MainPart{
					w.Docbar{
						Items: []w.DocbarPart{
							w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "note-edit"},
							w.Button{Label: "Cancel", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012"},
							w.BarSep{},
							w.Button{Label: "B", Bold: true},
							w.Button{Label: "I", Variant: w.ButtonItalic},
							w.Button{Label: "H2"},
							w.Button{Label: "• list"},
							w.Button{Folds: true, Label: "task", Icon: "task", Title: "Task"},
							w.Button{Icon: "link", Title: "Link"},
							w.Drop{
								Trigger: w.Button{Folds: true, Label: "Attach", Menu: true, Icon: "up", Title: "Attach"},
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuItem{
											Icon: "up",
											File: &w.FileField{Form: "note-edit", Name: "file", Label: "From this computer", List: "note-attachments"},
											Row: w.DocAttach{
												Name:   "{name}",
												Size:   w.Dim{Parts: []w.Inline{w.Text("· {size}")}},
												Remove: &w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Removes: true},
											},
											Label: "From this computer",
										},
										w.MenuItem{Icon: "file", Label: "From Files"},
									},
									Width: w.MenuWidth212,
									Pop:   w.MenuUnder,
								},
							},
						},
						More: &w.Button{Icon: "dots", Title: "More"},
						End: []w.DocbarEnd{
							w.Check{Parts: []w.CheckPart{w.Text("Auto-save")}},
							w.DraftNote{Text: "Draft in this window ·", Time: w.Mono{Text: "11:42"}, Secondary: true},
							w.Kbd{Keys: "Ctrl+S", Secondary: true},
							w.Seg{Options: []w.SegOption{{Label: "Full width", On: true}, {Label: "Reading width"}}, Secondary: true},
							w.Button{Folds: true, Label: "Focus", Icon: "focus", Variant: w.ButtonOnCard, Title: "Focus"},
						},
						Edit: true,
					},
					w.DocLayout{
						Items: []w.DocLayoutPart{
							w.Doc{
								Items: []w.DocPart{
									w.Area{Text: "Заметка со встречи 10.08", Variant: w.AreaTitle, Focus: true, Name: "summary", Label: "Text"},
									w.DocFields{
										Items: []w.Field{
											w.Field{
												Label: "Notebook",
												Parts: []w.FieldPart{
													w.Drop{
														Trigger: w.Pick{Value: "nextcloud · Записи", Mark: &w.Mark{Colour: "#4A6B52"}},
														Menu: w.Menu{
															Items: []w.MenuPart{
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#4A6B52"},
																	Label: "nextcloud · Записи",
																	On:    true,
																	Name:  "notebook",
																	Value: "nextcloud/zapisi",
																},
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#6B4E7A"},
																	Label: "nextcloud · Дневник",
																	Name:  "notebook",
																	Value: "nextcloud/dnevnik",
																},
																w.MenuItem{
																	Mark:   &w.Mark{Colour: "#3F6E8C"},
																	Trail:  w.Badge{Label: "ro", Place: w.Pushed},
																	Label:  "nextcloud · Общие с семьёй",
																	Reason: "Read-only: this account cannot write here",
																	Name:   "notebook",
																	Value:  "nextcloud/semya",
																},
															},
															Width: w.MenuWidth290,
															Pop:   w.MenuUnder,
														},
													},
												},
												Width: w.FieldFlex2_280,
											},
											w.Field{
												Label: "Date",
												Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "dtstart.date", Value: "2026-08-10", Label: "Date", Mono: true}},
												Width: w.FieldFix150,
											},
											w.Field{
												Label: "Time",
												Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Name: "dtstart.time", Value: "11:05", Label: "Time", Mono: true}},
												Width: w.FieldFix130,
											},
											w.Field{
												Label: "Tags",
												Parts: []w.FieldPart{w.Input{Name: "categories", Value: "встречи, релиз", Label: "Tags"}},
												Width: w.FieldFlex3_320,
											},
											w.Field{
												Label: "Related to",
												Parts: []w.FieldPart{
													w.Drop{
														Open: true,
														Trigger: w.TagInput{
															Parts: []w.TagInputPart{w.Tag{Label: "Синк-ап по релизу", Removable: true}},
															Add:   &w.TagAdd{Name: "related.add", Value: "ковал", Placeholder: "Title or UID", Label: "Add a link"},
															Focus: true,
														},
														Menu: w.Menu{
															Items: []w.MenuPart{
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#4A6B52"},
																	Trail: w.Dim{Parts: []w.Inline{w.Text("event · 10 Aug")}, Size: w.Dim11, Place: w.Pushed},
																	Label: "Звонок с Ковалёвой",
																	On:    true,
																},
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#3F6E8C"},
																	Trail: w.Dim{Parts: []w.Inline{w.Text("task · 14 Aug")}, Size: w.Dim11, Place: w.Pushed},
																	Label: "Ответить Анне по смете",
																},
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#4A6B52"},
																	Trail: w.Dim{Parts: []w.Inline{w.Text("contact")}, Size: w.Dim11, Place: w.Pushed},
																	Label: "Анна Ковалёва",
																},
																w.MenuSep{},
																w.MenuItem{Label: "By UID…", Tone: w.MenuToneDusty},
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
										Text:    "Договорились заморозить ветку в среду.\n\nПисьмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.\n\n## Что делаем\n\n- [ ] проверить changelog\n- [x] попросить у Артёма скриншоты\n\nПеренос старых вложений в новую папку — Виктор посмотрит, сколько их вообще, и скажет в четверг.",
										Variant: w.AreaBody,
										Name:    "description",
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
								ID:     "note-edit",
								Action: "/app/notes/nextcloud/zapisi/9e31-40b7-c012",
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
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 5}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
							{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
							{Icon: "note", Label: "Notes", Count: &w.Num{Value: 3}, Href: "#"},
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
								Meta:    "6",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "2"},
							w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "2"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "radicale · vps"},
							w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "1"},
							w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "1"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "baikal"},
							w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Error: true},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "Subscriptions", Items: []w.RubricPart{w.Badge{Label: "external"}}},
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
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "Tags"}},
						Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "релиз"}, w.Tag{Label: "работа"}},
					},
					w.RailFoot{
						Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
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
										w.Text(" · 6 events"),
									},
								},
							},
							Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}}},
						},
						w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.TimeRow{
									Mark:    w.Mark{Colour: "#4A6B52"},
									Time:    "10:00",
									Title:   []w.Inline{w.Text("Синк-ап по релизу")},
									Detail:  "Переговорная 2",
									Caption: &w.Badge{Label: "4 attendees"},
									Source:  "nextcloud · Личный",
								},
								w.TimeRow{
									Mark:    w.Mark{Colour: "#3F6E8C"},
									Time:    "all day",
									Title:   []w.Inline{w.Text("Дежурство: первая линия")},
									Detail:  "—",
									Caption: &w.Badge{Label: "repeats weekly"},
									Source:  "nextcloud · Работа",
								},
								w.TimeRow{
									Mark:    w.Mark{Colour: "#8C6239"},
									Time:    "all day",
									Title:   []w.Inline{w.Text("Дежурство: первая линия")},
									Detail:  "—",
									Caption: &w.Badge{Label: "repeats weekly"},
									Source:  "radicale · vps · Дежурства",
								},
								w.TimeRow{
									Mark:    w.Mark{Colour: "#3F6E8C"},
									Time:    "15:00",
									Title:   []w.Inline{w.Text("Созвон с подрядчиком")},
									Detail:  "Jitsi",
									Caption: &w.Badge{Label: "tentative", Tone: w.BadgeAlert},
									Source:  "nextcloud · Работа",
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Tuesday · 11 August", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
						w.List{
							Rows: []w.ListRow{
								w.TimeRow{
									Mark:    w.Mark{Colour: "#6B4E7A"},
									Time:    "09:30",
									Title:   []w.Inline{w.Text("Ревью архитектуры")},
									Detail:  "Переговорная 1",
									Caption: &w.Badge{Label: "2 attendees"},
									Source:  "radicale · vps · Проекты",
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
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 5}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 3}, Href: "#"},
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
							Meta:    "6",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "2"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "1"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "baikal"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Error: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Subscriptions", Items: []w.RubricPart{w.Badge{Label: "external"}}},
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
									w.Text(" · 6 events"),
								},
							},
						},
						Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}}},
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Time:    "10:00",
								Title:   []w.Inline{w.Text("Синк-ап по релизу")},
								Detail:  "Переговорная 2",
								Caption: &w.Badge{Label: "4 attendees"},
								Source:  "nextcloud · Личный",
							},
							w.TimeRow{
								Mark:    w.Mark{Colour: "#3F6E8C"},
								Time:    "all day",
								Title:   []w.Inline{w.Text("Дежурство: первая линия")},
								Detail:  "—",
								Caption: &w.Badge{Label: "repeats weekly"},
								Source:  "nextcloud · Работа",
							},
							w.TimeRow{
								Mark:    w.Mark{Colour: "#8C6239"},
								Time:    "all day",
								Title:   []w.Inline{w.Text("Дежурство: первая линия")},
								Detail:  "—",
								Caption: &w.Badge{Label: "repeats weekly"},
								Source:  "radicale · vps · Дежурства",
							},
							w.TimeRow{
								Mark:    w.Mark{Colour: "#3F6E8C"},
								Time:    "15:00",
								Title:   []w.Inline{w.Text("Созвон с подрядчиком")},
								Detail:  "Jitsi",
								Caption: &w.Badge{Label: "tentative", Tone: w.BadgeAlert},
								Source:  "nextcloud · Работа",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Tuesday · 11 August", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:    w.Mark{Colour: "#6B4E7A"},
								Time:    "09:30",
								Title:   []w.Inline{w.Text("Ревью архитектуры")},
								Detail:  "Переговорная 1",
								Caption: &w.Badge{Label: "2 attendees"},
								Source:  "radicale · vps · Проекты",
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
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 5}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
							{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
							{Icon: "note", Label: "Notes", Count: &w.Num{Value: 3}, Href: "#"},
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
								Meta:    "6",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "2"},
							w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "2"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "radicale · vps"},
							w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "1"},
							w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "1"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "baikal"},
							w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Error: true},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "Subscriptions", Items: []w.RubricPart{w.Badge{Label: "external"}}},
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
										w.Text(" · 6 events"),
									},
								},
							},
							Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}}},
						},
						w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.TimeRow{
									Mark:    w.Mark{Colour: "#4A6B52"},
									Time:    "10:00",
									Title:   []w.Inline{w.Text("Синк-ап по релизу")},
									Detail:  "Переговорная 2",
									Caption: &w.Badge{Label: "4 attendees"},
									Source:  "nextcloud · Личный",
								},
								w.TimeRow{
									Mark:    w.Mark{Colour: "#3F6E8C"},
									Time:    "all day",
									Title:   []w.Inline{w.Text("Дежурство: первая линия")},
									Detail:  "—",
									Caption: &w.Badge{Label: "repeats weekly"},
									Source:  "nextcloud · Работа",
								},
								w.TimeRow{
									Mark:    w.Mark{Colour: "#8C6239"},
									Time:    "all day",
									Title:   []w.Inline{w.Text("Дежурство: первая линия")},
									Detail:  "—",
									Caption: &w.Badge{Label: "repeats weekly"},
									Source:  "radicale · vps · Дежурства",
								},
								w.TimeRow{
									Mark:    w.Mark{Colour: "#3F6E8C"},
									Time:    "15:00",
									Title:   []w.Inline{w.Text("Созвон с подрядчиком")},
									Detail:  "Jitsi",
									Caption: &w.Badge{Label: "tentative", Tone: w.BadgeAlert},
									Source:  "nextcloud · Работа",
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Tuesday · 11 August", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
						w.List{
							Rows: []w.ListRow{
								w.TimeRow{
									Mark:    w.Mark{Colour: "#6B4E7A"},
									Time:    "09:30",
									Title:   []w.Inline{w.Text("Ревью архитектуры")},
									Detail:  "Переговорная 1",
									Caption: &w.Badge{Label: "2 attendees"},
									Source:  "radicale · vps · Проекты",
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
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 5}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
							{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
							{Icon: "note", Label: "Notes", Count: &w.Num{Value: 3}, Href: "#"},
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
								Meta:    "5",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "3"},
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
						Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
						About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
					},
				},
				Main: w.Main{
					Items: []w.MainPart{
						w.Head{
							Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined address book"},
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 5 cards")}}},
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
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 5}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
							{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
							{Icon: "note", Label: "Notes", Count: &w.Num{Value: 3}, Href: "#"},
							{Icon: "file", Label: "Files", Href: "#"},
							{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
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
								Meta:    "3",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2"},
							w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "1"},
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
						Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
						About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
					},
				},
				Main: w.Main{
					Items: []w.MainPart{
						w.Head{
							Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined notebook"},
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 of 3 notebooks · 3 notes · VJOURNAL")}}},
							Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary}}},
						},
						w.ListGroup{Title: w.Rubric{Text: "August 2026", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
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
									Date:    "3 Aug",
									Source:  "nextcloud · Дневник",
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
									Parts: []w.FieldPart{w.Input{Kind: w.InputTime, Placeholder: "--:--", Label: "Due time", Mono: true}},
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
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 5}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
							{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
							{Icon: "note", Label: "Notes", Count: &w.Num{Value: 3}, Href: "#"},
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
								Meta:    "6",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "2"},
							w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "2"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "radicale · vps"},
							w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "1"},
							w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "1"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "baikal"},
							w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Error: true},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "Subscriptions", Items: []w.RubricPart{w.Badge{Label: "external"}}},
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
										w.Text(" · 6 events"),
									},
								},
							},
							Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}}},
						},
						w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
						w.List{
							Rows: []w.ListRow{
								w.TimeRow{
									Mark:    w.Mark{Colour: "#4A6B52"},
									Time:    "10:00",
									Title:   []w.Inline{w.Text("Синк-ап по релизу")},
									Detail:  "Переговорная 2",
									Caption: &w.Badge{Label: "4 attendees"},
									Source:  "nextcloud · Личный",
								},
								w.TimeRow{
									Mark:    w.Mark{Colour: "#3F6E8C"},
									Time:    "all day",
									Title:   []w.Inline{w.Text("Дежурство: первая линия")},
									Detail:  "—",
									Caption: &w.Badge{Label: "repeats weekly"},
									Source:  "nextcloud · Работа",
								},
								w.TimeRow{
									Mark:    w.Mark{Colour: "#8C6239"},
									Time:    "all day",
									Title:   []w.Inline{w.Text("Дежурство: первая линия")},
									Detail:  "—",
									Caption: &w.Badge{Label: "repeats weekly"},
									Source:  "radicale · vps · Дежурства",
								},
								w.TimeRow{
									Mark:    w.Mark{Colour: "#3F6E8C"},
									Time:    "15:00",
									Title:   []w.Inline{w.Text("Созвон с подрядчиком")},
									Detail:  "Jitsi",
									Caption: &w.Badge{Label: "tentative", Tone: w.BadgeAlert},
									Source:  "nextcloud · Работа",
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "Tuesday · 11 August", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
						w.List{
							Rows: []w.ListRow{
								w.TimeRow{
									Mark:    w.Mark{Colour: "#6B4E7A"},
									Time:    "09:30",
									Title:   []w.Inline{w.Text("Ревью архитектуры")},
									Detail:  "Переговорная 1",
									Caption: &w.Badge{Label: "2 attendees"},
									Source:  "radicale · vps · Проекты",
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
					},
				},
				Height: w.ShellHeight560,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Icon:  "up",
					Title: "Upload files",
					Head: []w.SheetHeadPart{
						w.Pick{Value: "sftpgo · vps", Label: "Server", Variant: w.PickT12W240, Place: w.Pushed},
						w.Dim{
							Parts:   []w.Inline{w.Text("Documents / carrel / attachments / 2026-08")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
						w.Close{Label: "Close"},
					},
					Body: []w.SheetPart{
						w.Empty{
							Parts: []w.EmptyPart{
								w.Text("Drop files here, or "),
								w.FilePick{
									Label: "choose on the computer",
									File:  w.FileField{Form: "", Name: "file", Label: "Files", Size: "size", Many: true},
								},
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
										{Icon: "file", Label: "Documents"},
										{Icon: "file", Label: "carrel", Indent: w.TreeLevel22},
										{
											Icon:   "file",
											Label:  "attachments",
											Indent: w.TreeLevel40,
											Trail:  w.Badge{Label: "attachment folder"},
										},
										{Icon: "file", Label: "2026-08", Indent: w.TreeLevel58, Trail: w.Badge{Label: "just created"}},
										{Icon: "file", Label: "2026-07", Indent: w.TreeLevel58},
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
