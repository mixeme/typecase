// Раздел макета: 7.16 Настройки · публикация по ссылке.

package screens

import w "carrel.dev/typecase/widget"

// Publish1 is frame «Публикация · выданные ссылки» (s-publish-1).
func Publish1() Screen {
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
												Href: "/app/files/sftpgo/documents/carrel/attachments/смета-подрядчик.pdf",
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
	}
}

// Publish20 is frame «Опубликованное · почта не настроена» (s-publish-20).
func Publish20() Screen {
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
												Href: "/app/files/sftpgo/documents/carrel/attachments/смета-подрядчик.pdf",
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
	}
}

// Publish18 is frame «Публикация · ссылок нет» (s-publish-18).
func Publish18() Screen {
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
								{Icon: "link", Label: "Published links", Href: "#"},
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
					w.Empty{
						Parts: []w.EmptyPart{
							w.Strong("No links are published."),
							w.Text(" A link is made where the thing lies: «Publish a link» in the ⋯ menu of a row, a card or a collection."),
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
								On: 2,
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
							Title: &w.HeadTitle{Text: "Connections"},
							Subs: []w.HeadSub{
								w.HeadSub{
									Parts: []w.HeadSubPart{w.Text("Signed in as "), w.Strong("mix"), w.Text(" · 4 servers · 19 collections · 2 subscriptions")},
								},
							},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "Connect account", Icon: "plus", Kind: w.ButtonPrimary},
									w.Button{Label: "Subscribe to a feed", Icon: "link"},
								},
							},
						},
						w.PostForm{
							ID:     "conn-picks",
							Action: "/app/settings/connections?op=pick",
							Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
						},
						w.MainResults{
							Items: []w.ResultsPart{
								w.Rubric{
									Text:    "DAV servers",
									Items:   []w.RubricPart{w.Mono{Text: "4", Variant: w.MonoTight}},
									Variant: w.RubricSplit,
								},
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "nextcloud",
											Items: []w.RubricPart{
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "New collection", Icon: "plus", Small: true},
														w.Button{Label: "Show password", Small: true},
														w.Button{Label: "Test", Small: true},
														w.Button{
															Label: "Update collection list",
															Small: true,
															Form:  "conn-picks",
															To:    "/app/settings/connections?op=rediscover&account=nextcloud",
														},
														w.Button{Label: "Edit", Small: true},
														w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true},
													},
													Gap: w.LineGap6,
												},
											},
										},
										w.Mono{
											Text:    "https://cloud.example.org/remote.php/dav/ · mix · last answer 2 min ago",
											Form:    w.MonoLine,
											Variant: w.Mono11_5Dusty,
										},
										w.Table{
											Cols: []w.TableCol{
												{Width: w.ColTickSmall},
												{Label: "Collection"},
												{Label: "What is in it", Rank: w.Rank3},
												{Label: "Objects", Width: w.ColNum, Rank: w.Rank4},
												{Label: "Changed on the server", Width: w.ColChanged, Main: true},
												{Width: w.ColMenu},
											},
											Rows: []w.TableRow{
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|lichnyj"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Личный"}}},
														{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
														{Parts: []w.CellPart{w.Text("2")}, Mono: true},
														{Parts: []w.CellPart{w.Text("3 h ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|dela"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Дела"}}},
														{Parts: []w.CellPart{w.Text("tasks")}, Mono: true},
														{Parts: []w.CellPart{w.Text("4")}, Mono: true},
														{Parts: []w.CellPart{w.Text("20 min ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|rabota"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Работа"}}},
														{Parts: []w.CellPart{w.Text("calendar · tasks")}, Mono: true},
														{Parts: []w.CellPart{w.Text("4")}, Mono: true},
														{Parts: []w.CellPart{w.Text("yesterday")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{
															Pick: ptr(w.TickedOff),
															Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|dni-rozhdeniya"},
														},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "Дни рождения"}}},
														{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("12 August")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
													State: w.RowOff,
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|zapisi"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Записи"}}},
														{Parts: []w.CellPart{w.Text("journal")}, Mono: true},
														{Parts: []w.CellPart{w.Text("2")}, Mono: true},
														{Parts: []w.CellPart{w.Text("2 h ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|dnevnik"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Дневник"}}},
														{Parts: []w.CellPart{w.Text("journal")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("5 d ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{
															Pick: ptr(w.TickedOff),
															Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|zapiski-obshchie"},
														},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Общие с семьёй"}}},
														{Parts: []w.CellPart{w.Text("journal "), w.Badge{Label: "ro"}}, Mono: true},
														{Parts: []w.CellPart{w.Text("0")}, Mono: true},
														{Parts: []w.CellPart{w.Text("10 August")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
													State: w.RowOff,
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|kontakty"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Контакты"}}},
														{Parts: []w.CellPart{w.Text("address book")}, Mono: true},
														{Parts: []w.CellPart{w.Text("5")}, Mono: true},
														{Parts: []w.CellPart{w.Text("4 d ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|arhiv"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "Архив"}}},
														{Parts: []w.CellPart{w.Text("address book")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("28 July")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|files"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Files"}}},
														{Parts: []w.CellPart{w.Text("files")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("6 d ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
											},
										},
									},
								},
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "radicale · vps",
											Items: []w.RubricPart{
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "New collection", Icon: "plus", Small: true},
														w.Button{Label: "Show password", Small: true},
														w.Button{Label: "Test", Small: true},
														w.Button{
															Label: "Update collection list",
															Small: true,
															Form:  "conn-picks",
															To:    "/app/settings/connections?op=rediscover&account=radicale",
														},
														w.Button{Label: "Edit", Small: true},
														w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true},
													},
													Gap: w.LineGap6,
												},
											},
										},
										w.Mono{
											Text:    "https://dav.vps.example/radicale/ · mix · last answer 4 min ago",
											Form:    w.MonoLine,
											Variant: w.Mono11_5Dusty,
										},
										w.Table{
											Cols: []w.TableCol{
												{Width: w.ColTickSmall},
												{Label: "Collection"},
												{Label: "What is in it", Rank: w.Rank3},
												{Label: "Objects", Width: w.ColNum, Rank: w.Rank4},
												{Label: "Changed on the server", Width: w.ColChanged, Main: true},
												{Width: w.ColMenu},
											},
											Rows: []w.TableRow{
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "radicale|proekty"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Проекты"}}},
														{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1 h ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "radicale|dezhurstva"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#8C6239"}, Name: "Дежурства"}}},
														{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("yesterday")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "radicale|komanda"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Команда"}}},
														{Parts: []w.CellPart{w.Text("address book "), w.Badge{Label: "ro"}}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("6 d ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOff), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "radicale|rassylki"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Рассылки"}}},
														{Parts: []w.CellPart{w.Text("address book")}, Mono: true},
														{Parts: []w.CellPart{w.Text("0")}, Mono: true},
														{Parts: []w.CellPart{w.Text("3 August")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
													State: w.RowOff,
												},
											},
										},
									},
								},
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "baikal",
											Items: []w.RubricPart{
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "Show password", Small: true},
														w.Button{Label: "Test again", Small: true},
														w.Button{Label: "Edit", Small: true},
														w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true},
													},
													Gap: w.LineGap6,
												},
											},
											Variant: w.RubricDanger,
										},
										w.Mono{
											Text:    "https://dav.office.example/dav.php/ · m.yenuchenko · last answer refused",
											Form:    w.MonoLine,
											Variant: w.Mono11_5DustyGap,
										},
										w.Diag{
											Text: "[well-known]  GET /.well-known/caldav → 404\n[base]        PROPFIND /dav.php/ → 401 Unauthorized\n              WWW-Authenticate: Basic realm=\"BaikalDAV\"\n[principal]   не запрашивался: предыдущий шаг не прошёл",
										},
										w.CardText{
											Kind: w.CardAlertText,
											Parts: []w.CardTextPart{
												w.Text("The server refused the password. It may have been changed on the server: the one Carrel saved no longer fits."),
											},
										},
									},
									Variant: w.CardAlert,
								},
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "sftpgo · vps",
											Items: []w.RubricPart{
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "New folder", Icon: "plus", Small: true},
														w.Button{Label: "Show password", Small: true},
														w.Button{Label: "Test", Small: true},
														w.Button{
															Label: "Update collection list",
															Small: true,
															Form:  "conn-picks",
															To:    "/app/settings/connections?op=rediscover&account=sftpgo",
														},
														w.Button{Label: "Edit", Small: true},
														w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true},
													},
													Gap: w.LineGap6,
												},
											},
										},
										w.Mono{
											Text:    "https://files.example.org/dav/ · mix · plain WebDAV, no principal · last answer 5 min ago",
											Form:    w.MonoLine,
											Variant: w.Mono11_5Dusty,
										},
										w.Table{
											Cols: []w.TableCol{
												{Width: w.ColTickSmall},
												{Label: "Folder"},
												{Label: "What is in it", Rank: w.Rank3},
												{Label: "Objects", Width: w.ColNum, Rank: w.Rank4},
												{Label: "Changed on the server", Width: w.ColChanged, Main: true},
												{Width: w.ColMenu},
											},
											Rows: []w.TableRow{
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "sftpgo|documents"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Documents"}}},
														{Parts: []w.CellPart{w.Text("files")}, Mono: true},
														{Parts: []w.CellPart{w.Text("9")}, Mono: true},
														{Parts: []w.CellPart{w.Text("5 min ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOff), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "sftpgo|public"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#9A9280"}, Name: "Public"}}},
														{Parts: []w.CellPart{w.Text("files "), w.Badge{Label: "ro"}}, Mono: true},
														{Parts: []w.CellPart{w.Text("2")}, Mono: true},
														{Parts: []w.CellPart{w.Text("3 d ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
													State: w.RowOff,
												},
											},
										},
									},
								},
								w.Rubric{
									Text:    "Subscriptions · external, read-only",
									Items:   []w.RubricPart{w.Mono{Text: "2", Variant: w.MonoTight}},
									Variant: w.RubricSplit,
								},
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "Праздники РФ",
											Items: []w.RubricPart{
												w.Badge{Label: "feed"},
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "Check", Small: true},
														w.Button{Label: "Edit", Small: true},
														w.Button{Label: "Unsubscribe", Kind: w.ButtonDanger, Small: true},
													},
													Gap: w.LineGap6,
												},
											},
										},
										w.Mono{
											Text:    "https://feeds.example/ru-holidays.ics · no password · read 2 min ago · 34 KB · 0.4 s",
											Form:    w.MonoLine,
											Variant: w.Mono11_5Dusty,
										},
										w.Table{
											Cols: []w.TableCol{
												{Width: w.ColTickSmall},
												{Label: "Lands in"},
												{Label: "Objects", Width: w.ColNum, Rank: w.Rank3},
												{Label: "Changed on the server", Width: w.ColChanged, Rank: w.Rank4},
											},
											Rows: []w.TableRow{
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "ru-holidays|calendar"}},
														{
															Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#8C6239", Kind: w.MarkExternal}, Name: "Calendar"}},
														},
														{Parts: []w.CellPart{w.Text("3")}, Mono: true},
														{Parts: []w.CellPart{w.Text("yesterday")}, Mono: true, Rows: 2},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOff), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "ru-holidays|tasks"}},
														{
															Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#8C6239", Kind: w.MarkExternal}, Name: "Tasks"}},
														},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
													},
													State: w.RowOff,
												},
											},
										},
									},
								},
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "Расписание лиги",
											Items: []w.RubricPart{
												w.Badge{Label: "feed"},
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "Show password", Small: true},
														w.Button{Label: "Check again", Small: true},
														w.Button{Label: "Edit", Small: true},
														w.Button{Label: "Unsubscribe", Kind: w.ButtonDanger, Small: true},
													},
													Gap: w.LineGap6,
												},
											},
											Variant: w.RubricDanger,
										},
										w.Mono{
											Text:    "webcal://league.example/season.ics · password saved · last read yesterday",
											Form:    w.MonoLine,
											Variant: w.Mono11_5DustyGap,
										},
										w.Diag{
											Text: "[fetch]   GET /season.ics → 401 Unauthorized\n              WWW-Authenticate: Basic realm=\"league\"\n[cache]   2 events from yesterday are still shown, marked stale",
										},
										w.CardText{
											Kind: w.CardAlertText,
											Parts: []w.CardTextPart{
												w.Text("The feed asks for a password and the saved one no longer fits. A feed that stopped answering must not look like a feed with nothing in it."),
											},
										},
									},
									Variant: w.CardAlert,
								},
							},
						},
					},
				},
				Height: w.ShellHeight390,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/publish?c=lichnyj",
					Title:  "Publish a calendar",
					Head: []w.SheetHeadPart{
						w.Badge{
							Label: "local",
							Tone:  w.BadgeLocal,
							Title: "Local instance: the link will not open from outside your network",
						},
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
							Parts: []w.FieldPart{
								w.Seg{
									Options: []w.SegOption{
										{Label: "A date range", On: true, Value: "range"},
										{Label: "The whole calendar", Value: "whole", Disables: "from to"},
									},
									Name: "how",
								},
							},
							Width: w.FieldW100,
						},
						w.Line{
							Items: []w.LinePart{
								w.Field{
									Label: "From",
									Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true}},
									Width: w.FieldSM,
								},
								w.Field{
									Label: "To",
									Parts: []w.FieldPart{w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-31", Label: "To", Mono: true}},
									Width: w.FieldSM,
								},
							},
							Gap:   w.LineGap10,
							Block: true,
						},
						w.Field{
							Label: "What the visitor sees",
							Parts: []w.FieldPart{
								w.Group{
									Items: []w.FieldPart{
										w.Check{
											On:     true,
											Name:   "time",
											Value:  "1",
											Reason: "Always shown",
											Parts:  []w.CheckPart{w.Text("Time"), w.Dim{Parts: []w.Inline{w.Text("always shown")}, Size: w.Dim11_5}},
										},
										w.Check{On: true, Name: "title", Value: "1", Parts: []w.CheckPart{w.Text("Title")}},
										w.Check{Name: "location", Value: "1", Parts: []w.CheckPart{w.Text("Location")}},
										w.Check{Name: "description", Value: "1", Parts: []w.CheckPart{w.Text("Description")}},
										w.Check{Name: "attendees", Value: "1", Parts: []w.CheckPart{w.Text("Attendees")}},
										w.Check{Name: "attachments", Value: "1", Parts: []w.CheckPart{w.Text("Attachments")}},
										w.Check{Name: "tags", Value: "1", Parts: []w.CheckPart{w.Text("Tags")}},
									},
									Form: w.GroupColumnLoose,
								},
								w.FieldHint{Parts: []w.Inline{w.Text("Without the title an event shows as “Busy”.")}},
							},
							Width: w.FieldW100,
						},
						w.SheetChecks{
							Items: []w.Check{
								w.Check{
									Name:  "download",
									Value: "1",
									Parts: []w.CheckPart{w.Text("Allow "), w.Mono{Text: ".ics"}, w.Text(" download and subscription")},
								},
							},
						},
						w.Field{
							Label: "Link lives for",
							Parts: []w.FieldPart{
								w.Drop{
									Trigger: w.Pick{Value: "30 days"},
									Menu: w.Menu{
										Items: []w.MenuPart{
											w.MenuItem{Label: "7 days", Name: "lives", Value: "7"},
											w.MenuItem{Label: "14 days", Name: "lives", Value: "14"},
											w.MenuItem{Label: "30 days", On: true, Name: "lives", Value: "30"},
											w.MenuSep{},
											w.MenuItem{Label: "No expiry", Name: "lives", Value: "never"},
										},
										Width: w.MenuWidth212,
										Pop:   w.MenuUnder,
									},
								},
							},
							Width: w.FieldSM,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Create link", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/calendar"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGap7,
							Place: w.Pushed,
						},
					},
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
								On: 2,
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
							Title: &w.HeadTitle{Text: "Connections"},
							Subs: []w.HeadSub{
								w.HeadSub{
									Parts: []w.HeadSubPart{w.Text("Signed in as "), w.Strong("mix"), w.Text(" · 4 servers · 19 collections · 2 subscriptions")},
								},
							},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "Connect account", Icon: "plus", Kind: w.ButtonPrimary},
									w.Button{Label: "Subscribe to a feed", Icon: "link"},
								},
							},
						},
						w.PostForm{
							ID:     "conn-picks",
							Action: "/app/settings/connections?op=pick",
							Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
						},
						w.MainResults{
							Items: []w.ResultsPart{
								w.Rubric{
									Text:    "DAV servers",
									Items:   []w.RubricPart{w.Mono{Text: "4", Variant: w.MonoTight}},
									Variant: w.RubricSplit,
								},
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "nextcloud",
											Items: []w.RubricPart{
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "New collection", Icon: "plus", Small: true},
														w.Button{Label: "Show password", Small: true},
														w.Button{Label: "Test", Small: true},
														w.Button{
															Label: "Update collection list",
															Small: true,
															Form:  "conn-picks",
															To:    "/app/settings/connections?op=rediscover&account=nextcloud",
														},
														w.Button{Label: "Edit", Small: true},
														w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true},
													},
													Gap: w.LineGap6,
												},
											},
										},
										w.Mono{
											Text:    "https://cloud.example.org/remote.php/dav/ · mix · last answer 2 min ago",
											Form:    w.MonoLine,
											Variant: w.Mono11_5Dusty,
										},
										w.Table{
											Cols: []w.TableCol{
												{Width: w.ColTickSmall},
												{Label: "Collection"},
												{Label: "What is in it", Rank: w.Rank3},
												{Label: "Objects", Width: w.ColNum, Rank: w.Rank4},
												{Label: "Changed on the server", Width: w.ColChanged, Main: true},
												{Width: w.ColMenu},
											},
											Rows: []w.TableRow{
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|lichnyj"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Личный"}}},
														{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
														{Parts: []w.CellPart{w.Text("2")}, Mono: true},
														{Parts: []w.CellPart{w.Text("3 h ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|dela"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Дела"}}},
														{Parts: []w.CellPart{w.Text("tasks")}, Mono: true},
														{Parts: []w.CellPart{w.Text("4")}, Mono: true},
														{Parts: []w.CellPart{w.Text("20 min ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|rabota"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Работа"}}},
														{Parts: []w.CellPart{w.Text("calendar · tasks")}, Mono: true},
														{Parts: []w.CellPart{w.Text("4")}, Mono: true},
														{Parts: []w.CellPart{w.Text("yesterday")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{
															Pick: ptr(w.TickedOff),
															Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|dni-rozhdeniya"},
														},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "Дни рождения"}}},
														{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("12 August")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
													State: w.RowOff,
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|zapisi"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Записи"}}},
														{Parts: []w.CellPart{w.Text("journal")}, Mono: true},
														{Parts: []w.CellPart{w.Text("2")}, Mono: true},
														{Parts: []w.CellPart{w.Text("2 h ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|dnevnik"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Дневник"}}},
														{Parts: []w.CellPart{w.Text("journal")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("5 d ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{
															Pick: ptr(w.TickedOff),
															Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|zapiski-obshchie"},
														},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Общие с семьёй"}}},
														{Parts: []w.CellPart{w.Text("journal "), w.Badge{Label: "ro"}}, Mono: true},
														{Parts: []w.CellPart{w.Text("0")}, Mono: true},
														{Parts: []w.CellPart{w.Text("10 August")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
													State: w.RowOff,
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|kontakty"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Контакты"}}},
														{Parts: []w.CellPart{w.Text("address book")}, Mono: true},
														{Parts: []w.CellPart{w.Text("5")}, Mono: true},
														{Parts: []w.CellPart{w.Text("4 d ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|arhiv"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "Архив"}}},
														{Parts: []w.CellPart{w.Text("address book")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("28 July")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|files"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Files"}}},
														{Parts: []w.CellPart{w.Text("files")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("6 d ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
											},
										},
									},
								},
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "radicale · vps",
											Items: []w.RubricPart{
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "New collection", Icon: "plus", Small: true},
														w.Button{Label: "Show password", Small: true},
														w.Button{Label: "Test", Small: true},
														w.Button{
															Label: "Update collection list",
															Small: true,
															Form:  "conn-picks",
															To:    "/app/settings/connections?op=rediscover&account=radicale",
														},
														w.Button{Label: "Edit", Small: true},
														w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true},
													},
													Gap: w.LineGap6,
												},
											},
										},
										w.Mono{
											Text:    "https://dav.vps.example/radicale/ · mix · last answer 4 min ago",
											Form:    w.MonoLine,
											Variant: w.Mono11_5Dusty,
										},
										w.Table{
											Cols: []w.TableCol{
												{Width: w.ColTickSmall},
												{Label: "Collection"},
												{Label: "What is in it", Rank: w.Rank3},
												{Label: "Objects", Width: w.ColNum, Rank: w.Rank4},
												{Label: "Changed on the server", Width: w.ColChanged, Main: true},
												{Width: w.ColMenu},
											},
											Rows: []w.TableRow{
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "radicale|proekty"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Проекты"}}},
														{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1 h ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "radicale|dezhurstva"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#8C6239"}, Name: "Дежурства"}}},
														{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("yesterday")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "radicale|komanda"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Команда"}}},
														{Parts: []w.CellPart{w.Text("address book "), w.Badge{Label: "ro"}}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("6 d ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOff), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "radicale|rassylki"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Рассылки"}}},
														{Parts: []w.CellPart{w.Text("address book")}, Mono: true},
														{Parts: []w.CellPart{w.Text("0")}, Mono: true},
														{Parts: []w.CellPart{w.Text("3 August")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
													State: w.RowOff,
												},
											},
										},
									},
								},
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "baikal",
											Items: []w.RubricPart{
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "Show password", Small: true},
														w.Button{Label: "Test again", Small: true},
														w.Button{Label: "Edit", Small: true},
														w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true},
													},
													Gap: w.LineGap6,
												},
											},
											Variant: w.RubricDanger,
										},
										w.Mono{
											Text:    "https://dav.office.example/dav.php/ · m.yenuchenko · last answer refused",
											Form:    w.MonoLine,
											Variant: w.Mono11_5DustyGap,
										},
										w.Diag{
											Text: "[well-known]  GET /.well-known/caldav → 404\n[base]        PROPFIND /dav.php/ → 401 Unauthorized\n              WWW-Authenticate: Basic realm=\"BaikalDAV\"\n[principal]   не запрашивался: предыдущий шаг не прошёл",
										},
										w.CardText{
											Kind: w.CardAlertText,
											Parts: []w.CardTextPart{
												w.Text("The server refused the password. It may have been changed on the server: the one Carrel saved no longer fits."),
											},
										},
									},
									Variant: w.CardAlert,
								},
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "sftpgo · vps",
											Items: []w.RubricPart{
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "New folder", Icon: "plus", Small: true},
														w.Button{Label: "Show password", Small: true},
														w.Button{Label: "Test", Small: true},
														w.Button{
															Label: "Update collection list",
															Small: true,
															Form:  "conn-picks",
															To:    "/app/settings/connections?op=rediscover&account=sftpgo",
														},
														w.Button{Label: "Edit", Small: true},
														w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true},
													},
													Gap: w.LineGap6,
												},
											},
										},
										w.Mono{
											Text:    "https://files.example.org/dav/ · mix · plain WebDAV, no principal · last answer 5 min ago",
											Form:    w.MonoLine,
											Variant: w.Mono11_5Dusty,
										},
										w.Table{
											Cols: []w.TableCol{
												{Width: w.ColTickSmall},
												{Label: "Folder"},
												{Label: "What is in it", Rank: w.Rank3},
												{Label: "Objects", Width: w.ColNum, Rank: w.Rank4},
												{Label: "Changed on the server", Width: w.ColChanged, Main: true},
												{Width: w.ColMenu},
											},
											Rows: []w.TableRow{
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "sftpgo|documents"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Documents"}}},
														{Parts: []w.CellPart{w.Text("files")}, Mono: true},
														{Parts: []w.CellPart{w.Text("9")}, Mono: true},
														{Parts: []w.CellPart{w.Text("5 min ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOff), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "sftpgo|public"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#9A9280"}, Name: "Public"}}},
														{Parts: []w.CellPart{w.Text("files "), w.Badge{Label: "ro"}}, Mono: true},
														{Parts: []w.CellPart{w.Text("2")}, Mono: true},
														{Parts: []w.CellPart{w.Text("3 d ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
													State: w.RowOff,
												},
											},
										},
									},
								},
								w.Rubric{
									Text:    "Subscriptions · external, read-only",
									Items:   []w.RubricPart{w.Mono{Text: "2", Variant: w.MonoTight}},
									Variant: w.RubricSplit,
								},
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "Праздники РФ",
											Items: []w.RubricPart{
												w.Badge{Label: "feed"},
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "Check", Small: true},
														w.Button{Label: "Edit", Small: true},
														w.Button{Label: "Unsubscribe", Kind: w.ButtonDanger, Small: true},
													},
													Gap: w.LineGap6,
												},
											},
										},
										w.Mono{
											Text:    "https://feeds.example/ru-holidays.ics · no password · read 2 min ago · 34 KB · 0.4 s",
											Form:    w.MonoLine,
											Variant: w.Mono11_5Dusty,
										},
										w.Table{
											Cols: []w.TableCol{
												{Width: w.ColTickSmall},
												{Label: "Lands in"},
												{Label: "Objects", Width: w.ColNum, Rank: w.Rank3},
												{Label: "Changed on the server", Width: w.ColChanged, Rank: w.Rank4},
											},
											Rows: []w.TableRow{
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "ru-holidays|calendar"}},
														{
															Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#8C6239", Kind: w.MarkExternal}, Name: "Calendar"}},
														},
														{Parts: []w.CellPart{w.Text("3")}, Mono: true},
														{Parts: []w.CellPart{w.Text("yesterday")}, Mono: true, Rows: 2},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOff), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "ru-holidays|tasks"}},
														{
															Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#8C6239", Kind: w.MarkExternal}, Name: "Tasks"}},
														},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
													},
													State: w.RowOff,
												},
											},
										},
									},
								},
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "Расписание лиги",
											Items: []w.RubricPart{
												w.Badge{Label: "feed"},
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "Show password", Small: true},
														w.Button{Label: "Check again", Small: true},
														w.Button{Label: "Edit", Small: true},
														w.Button{Label: "Unsubscribe", Kind: w.ButtonDanger, Small: true},
													},
													Gap: w.LineGap6,
												},
											},
											Variant: w.RubricDanger,
										},
										w.Mono{
											Text:    "webcal://league.example/season.ics · password saved · last read yesterday",
											Form:    w.MonoLine,
											Variant: w.Mono11_5DustyGap,
										},
										w.Diag{
											Text: "[fetch]   GET /season.ics → 401 Unauthorized\n              WWW-Authenticate: Basic realm=\"league\"\n[cache]   2 events from yesterday are still shown, marked stale",
										},
										w.CardText{
											Kind: w.CardAlertText,
											Parts: []w.CardTextPart{
												w.Text("The feed asks for a password and the saved one no longer fits. A feed that stopped answering must not look like a feed with nothing in it."),
											},
										},
									},
									Variant: w.CardAlert,
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
						w.Badge{
							Label: "local",
							Tone:  w.BadgeLocal,
							Title: "Local instance: the link will not open from outside your network",
						},
						w.Dim{
							Parts:   []w.Inline{w.Text("nextcloud · Личный")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
						w.Close{Label: "Close", Href: "/app/calendar"},
					},
					Body: []w.SheetPart{
						w.Field{
							Label: "Link",
							Parts: []w.FieldPart{
								w.Line{
									Items: []w.LinePart{
										w.Input{
											Value:    "https://carrel.example/p/7Qm2c8XvLk3w",
											Label:    "Link",
											Mono:     true,
											ReadOnly: true,
											Variant:  w.InputFill,
										},
										w.Button{
											Label:  "Copy",
											Icon:   "dup",
											Small:  true,
											Copy:   "https://carrel.example/p/7Qm2c8XvLk3w",
											Copied: "Copied",
										},
									},
									Gap: w.LineGap6,
								},
								w.FieldHint{Parts: []w.Inline{w.Text("Until 2026-09-09.")}},
							},
							Width: w.FieldW100,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Send by letter", Kind: w.ButtonPrimary, Href: "/app/mail/link?id=7Qm2"},
						w.Kbd{Keys: "Enter"},
						w.Button{Label: "Open the link", Href: "/p/7Qm2c8XvLk3w"},
						w.Button{Label: "All published links", Href: "/app/files/published"},
					},
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Publish22 is frame «Календарь · ссылка без срока создана» (s-publish-22).
func Publish22() Screen {
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
								On: 2,
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
							Title: &w.HeadTitle{Text: "Connections"},
							Subs: []w.HeadSub{
								w.HeadSub{
									Parts: []w.HeadSubPart{w.Text("Signed in as "), w.Strong("mix"), w.Text(" · 4 servers · 19 collections · 2 subscriptions")},
								},
							},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "Connect account", Icon: "plus", Kind: w.ButtonPrimary},
									w.Button{Label: "Subscribe to a feed", Icon: "link"},
								},
							},
						},
						w.PostForm{
							ID:     "conn-picks",
							Action: "/app/settings/connections?op=pick",
							Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
						},
						w.MainResults{
							Items: []w.ResultsPart{
								w.Rubric{
									Text:    "DAV servers",
									Items:   []w.RubricPart{w.Mono{Text: "4", Variant: w.MonoTight}},
									Variant: w.RubricSplit,
								},
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "nextcloud",
											Items: []w.RubricPart{
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "New collection", Icon: "plus", Small: true},
														w.Button{Label: "Show password", Small: true},
														w.Button{Label: "Test", Small: true},
														w.Button{
															Label: "Update collection list",
															Small: true,
															Form:  "conn-picks",
															To:    "/app/settings/connections?op=rediscover&account=nextcloud",
														},
														w.Button{Label: "Edit", Small: true},
														w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true},
													},
													Gap: w.LineGap6,
												},
											},
										},
										w.Mono{
											Text:    "https://cloud.example.org/remote.php/dav/ · mix · last answer 2 min ago",
											Form:    w.MonoLine,
											Variant: w.Mono11_5Dusty,
										},
										w.Table{
											Cols: []w.TableCol{
												{Width: w.ColTickSmall},
												{Label: "Collection"},
												{Label: "What is in it", Rank: w.Rank3},
												{Label: "Objects", Width: w.ColNum, Rank: w.Rank4},
												{Label: "Changed on the server", Width: w.ColChanged, Main: true},
												{Width: w.ColMenu},
											},
											Rows: []w.TableRow{
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|lichnyj"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Личный"}}},
														{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
														{Parts: []w.CellPart{w.Text("2")}, Mono: true},
														{Parts: []w.CellPart{w.Text("3 h ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|dela"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Дела"}}},
														{Parts: []w.CellPart{w.Text("tasks")}, Mono: true},
														{Parts: []w.CellPart{w.Text("4")}, Mono: true},
														{Parts: []w.CellPart{w.Text("20 min ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|rabota"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Работа"}}},
														{Parts: []w.CellPart{w.Text("calendar · tasks")}, Mono: true},
														{Parts: []w.CellPart{w.Text("4")}, Mono: true},
														{Parts: []w.CellPart{w.Text("yesterday")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{
															Pick: ptr(w.TickedOff),
															Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|dni-rozhdeniya"},
														},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "Дни рождения"}}},
														{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("12 August")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
													State: w.RowOff,
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|zapisi"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Записи"}}},
														{Parts: []w.CellPart{w.Text("journal")}, Mono: true},
														{Parts: []w.CellPart{w.Text("2")}, Mono: true},
														{Parts: []w.CellPart{w.Text("2 h ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|dnevnik"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Дневник"}}},
														{Parts: []w.CellPart{w.Text("journal")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("5 d ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{
															Pick: ptr(w.TickedOff),
															Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|zapiski-obshchie"},
														},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Общие с семьёй"}}},
														{Parts: []w.CellPart{w.Text("journal "), w.Badge{Label: "ro"}}, Mono: true},
														{Parts: []w.CellPart{w.Text("0")}, Mono: true},
														{Parts: []w.CellPart{w.Text("10 August")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
													State: w.RowOff,
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|kontakty"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Контакты"}}},
														{Parts: []w.CellPart{w.Text("address book")}, Mono: true},
														{Parts: []w.CellPart{w.Text("5")}, Mono: true},
														{Parts: []w.CellPart{w.Text("4 d ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|arhiv"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "Архив"}}},
														{Parts: []w.CellPart{w.Text("address book")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("28 July")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|files"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Files"}}},
														{Parts: []w.CellPart{w.Text("files")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("6 d ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
											},
										},
									},
								},
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "radicale · vps",
											Items: []w.RubricPart{
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "New collection", Icon: "plus", Small: true},
														w.Button{Label: "Show password", Small: true},
														w.Button{Label: "Test", Small: true},
														w.Button{
															Label: "Update collection list",
															Small: true,
															Form:  "conn-picks",
															To:    "/app/settings/connections?op=rediscover&account=radicale",
														},
														w.Button{Label: "Edit", Small: true},
														w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true},
													},
													Gap: w.LineGap6,
												},
											},
										},
										w.Mono{
											Text:    "https://dav.vps.example/radicale/ · mix · last answer 4 min ago",
											Form:    w.MonoLine,
											Variant: w.Mono11_5Dusty,
										},
										w.Table{
											Cols: []w.TableCol{
												{Width: w.ColTickSmall},
												{Label: "Collection"},
												{Label: "What is in it", Rank: w.Rank3},
												{Label: "Objects", Width: w.ColNum, Rank: w.Rank4},
												{Label: "Changed on the server", Width: w.ColChanged, Main: true},
												{Width: w.ColMenu},
											},
											Rows: []w.TableRow{
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "radicale|proekty"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Проекты"}}},
														{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1 h ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "radicale|dezhurstva"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#8C6239"}, Name: "Дежурства"}}},
														{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("yesterday")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "radicale|komanda"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Команда"}}},
														{Parts: []w.CellPart{w.Text("address book "), w.Badge{Label: "ro"}}, Mono: true},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
														{Parts: []w.CellPart{w.Text("6 d ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOff), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "radicale|rassylki"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Рассылки"}}},
														{Parts: []w.CellPart{w.Text("address book")}, Mono: true},
														{Parts: []w.CellPart{w.Text("0")}, Mono: true},
														{Parts: []w.CellPart{w.Text("3 August")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
													State: w.RowOff,
												},
											},
										},
									},
								},
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "baikal",
											Items: []w.RubricPart{
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "Show password", Small: true},
														w.Button{Label: "Test again", Small: true},
														w.Button{Label: "Edit", Small: true},
														w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true},
													},
													Gap: w.LineGap6,
												},
											},
											Variant: w.RubricDanger,
										},
										w.Mono{
											Text:    "https://dav.office.example/dav.php/ · m.yenuchenko · last answer refused",
											Form:    w.MonoLine,
											Variant: w.Mono11_5DustyGap,
										},
										w.Diag{
											Text: "[well-known]  GET /.well-known/caldav → 404\n[base]        PROPFIND /dav.php/ → 401 Unauthorized\n              WWW-Authenticate: Basic realm=\"BaikalDAV\"\n[principal]   не запрашивался: предыдущий шаг не прошёл",
										},
										w.CardText{
											Kind: w.CardAlertText,
											Parts: []w.CardTextPart{
												w.Text("The server refused the password. It may have been changed on the server: the one Carrel saved no longer fits."),
											},
										},
									},
									Variant: w.CardAlert,
								},
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "sftpgo · vps",
											Items: []w.RubricPart{
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "New folder", Icon: "plus", Small: true},
														w.Button{Label: "Show password", Small: true},
														w.Button{Label: "Test", Small: true},
														w.Button{
															Label: "Update collection list",
															Small: true,
															Form:  "conn-picks",
															To:    "/app/settings/connections?op=rediscover&account=sftpgo",
														},
														w.Button{Label: "Edit", Small: true},
														w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true},
													},
													Gap: w.LineGap6,
												},
											},
										},
										w.Mono{
											Text:    "https://files.example.org/dav/ · mix · plain WebDAV, no principal · last answer 5 min ago",
											Form:    w.MonoLine,
											Variant: w.Mono11_5Dusty,
										},
										w.Table{
											Cols: []w.TableCol{
												{Width: w.ColTickSmall},
												{Label: "Folder"},
												{Label: "What is in it", Rank: w.Rank3},
												{Label: "Objects", Width: w.ColNum, Rank: w.Rank4},
												{Label: "Changed on the server", Width: w.ColChanged, Main: true},
												{Width: w.ColMenu},
											},
											Rows: []w.TableRow{
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "sftpgo|documents"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Documents"}}},
														{Parts: []w.CellPart{w.Text("files")}, Mono: true},
														{Parts: []w.CellPart{w.Text("9")}, Mono: true},
														{Parts: []w.CellPart{w.Text("5 min ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOff), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "sftpgo|public"}},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#9A9280"}, Name: "Public"}}},
														{Parts: []w.CellPart{w.Text("files "), w.Badge{Label: "ro"}}, Mono: true},
														{Parts: []w.CellPart{w.Text("2")}, Mono: true},
														{Parts: []w.CellPart{w.Text("3 d ago")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
													State: w.RowOff,
												},
											},
										},
									},
								},
								w.Rubric{
									Text:    "Subscriptions · external, read-only",
									Items:   []w.RubricPart{w.Mono{Text: "2", Variant: w.MonoTight}},
									Variant: w.RubricSplit,
								},
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "Праздники РФ",
											Items: []w.RubricPart{
												w.Badge{Label: "feed"},
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "Check", Small: true},
														w.Button{Label: "Edit", Small: true},
														w.Button{Label: "Unsubscribe", Kind: w.ButtonDanger, Small: true},
													},
													Gap: w.LineGap6,
												},
											},
										},
										w.Mono{
											Text:    "https://feeds.example/ru-holidays.ics · no password · read 2 min ago · 34 KB · 0.4 s",
											Form:    w.MonoLine,
											Variant: w.Mono11_5Dusty,
										},
										w.Table{
											Cols: []w.TableCol{
												{Width: w.ColTickSmall},
												{Label: "Lands in"},
												{Label: "Objects", Width: w.ColNum, Rank: w.Rank3},
												{Label: "Changed on the server", Width: w.ColChanged, Rank: w.Rank4},
											},
											Rows: []w.TableRow{
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "ru-holidays|calendar"}},
														{
															Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#8C6239", Kind: w.MarkExternal}, Name: "Calendar"}},
														},
														{Parts: []w.CellPart{w.Text("3")}, Mono: true},
														{Parts: []w.CellPart{w.Text("yesterday")}, Mono: true, Rows: 2},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOff), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "ru-holidays|tasks"}},
														{
															Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#8C6239", Kind: w.MarkExternal}, Name: "Tasks"}},
														},
														{Parts: []w.CellPart{w.Text("1")}, Mono: true},
													},
													State: w.RowOff,
												},
											},
										},
									},
								},
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "Расписание лиги",
											Items: []w.RubricPart{
												w.Badge{Label: "feed"},
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "Show password", Small: true},
														w.Button{Label: "Check again", Small: true},
														w.Button{Label: "Edit", Small: true},
														w.Button{Label: "Unsubscribe", Kind: w.ButtonDanger, Small: true},
													},
													Gap: w.LineGap6,
												},
											},
											Variant: w.RubricDanger,
										},
										w.Mono{
											Text:    "webcal://league.example/season.ics · password saved · last read yesterday",
											Form:    w.MonoLine,
											Variant: w.Mono11_5DustyGap,
										},
										w.Diag{
											Text: "[fetch]   GET /season.ics → 401 Unauthorized\n              WWW-Authenticate: Basic realm=\"league\"\n[cache]   2 events from yesterday are still shown, marked stale",
										},
										w.CardText{
											Kind: w.CardAlertText,
											Parts: []w.CardTextPart{
												w.Text("The feed asks for a password and the saved one no longer fits. A feed that stopped answering must not look like a feed with nothing in it."),
											},
										},
									},
									Variant: w.CardAlert,
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
						w.Badge{
							Label: "local",
							Tone:  w.BadgeLocal,
							Title: "Local instance: the link will not open from outside your network",
						},
						w.Dim{
							Parts:   []w.Inline{w.Text("nextcloud · Личный")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
						w.Close{Label: "Close", Href: "/app/calendar"},
					},
					Body: []w.SheetPart{
						w.Field{
							Label: "Link",
							Parts: []w.FieldPart{
								w.Line{
									Items: []w.LinePart{
										w.Input{
											Value:    "https://carrel.example/p/7Qm2c8XvLk3w",
											Label:    "Link",
											Mono:     true,
											ReadOnly: true,
											Variant:  w.InputFill,
										},
										w.Button{
											Label:  "Copy",
											Icon:   "dup",
											Small:  true,
											Copy:   "https://carrel.example/p/7Qm2c8XvLk3w",
											Copied: "Copied",
										},
									},
									Gap: w.LineGap6,
								},
								w.FieldHint{Parts: []w.Inline{w.Text("No expiry. Revoke it when it is no longer needed.")}},
							},
							Width: w.FieldW100,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Open the link", Href: "/p/7Qm2c8XvLk3w"},
						w.Button{Label: "All published links", Href: "/app/files/published"},
					},
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Publish19 is frame «Событие и задача · создать ссылку» (s-publish-19).
func Publish19() Screen {
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
						{Icon: "back", Label: "Back to list", Short: "Back", Href: "#"}}},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "2", Current: true},
							w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "2"},
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
							Title: &w.HeadTitle{Text: "Синк-ап по релизу"},
							Subs: []w.HeadSub{
								w.HeadSub{Parts: []w.HeadSubPart{w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, w.Text("nextcloud · Личный")}},
							},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "Edit", Icon: "pencil", Kind: w.ButtonPrimary},
									w.Button{Label: "Send", Icon: "mail"},
									w.Button{Label: "Print", Icon: "print", Tertiary: true},
									w.Button{Icon: "dots", Title: "More"},
								},
							},
						},
						w.Fields{
							Items: []w.FieldsItem{
								{
									Term: "When",
									Value: []w.ValuePart{
										w.Mono{Text: "10:00 – 11:00"},
										w.Text(", Monday, 10 August "),
										w.Badge{Label: "created 12:00 Asia/Yekaterinburg", Title: "Created in another time zone"},
									},
								},
								{Term: "Repeats", Value: []w.ValuePart{w.Text("Weekly on Mon "), w.Badge{Label: "until 28 December"}}},
								{Term: "Reminders", Value: []w.ValuePart{w.Text("10 minutes before · 1 day before")}},
								{
									Term:  "Where",
									Value: []w.ValuePart{w.Text("Переговорная 2 "), w.Badge{Label: "linked", Tone: w.BadgeLinked}},
								},
								{Term: "Status", Value: []w.ValuePart{w.Text("Confirmed")}},
								{Term: "Tags", Value: []w.ValuePart{w.Tag{Label: "релиз"}, w.Text(" "), w.Tag{Label: "работа"}}},
							},
						},
						w.Sec{
							Items: []w.SecPart{
								w.Rubric{Text: "Description"},
								w.Prose{
									Parts: []w.Inline{w.Text("Обсудить оставшиеся блокеры, дату заморозки и кто пишет письмо пользователям.")},
								},
							},
						},
						w.Sec{
							Items: []w.SecPart{
								w.Rubric{Text: "Attendees", Items: []w.RubricPart{w.Mono{Text: "4", Variant: w.MonoTight}}},
								w.SecList{
									Items: [][]w.SecItemPart{
										[]w.SecItemPart{
											w.Mark{Colour: "#6B4E7A", Size: w.MarkH11Inline},
											w.Text(" "),
											w.Link{Text: "Мария Лебедева", Href: "#"},
											w.Text(" "),
											w.Dim{Parts: []w.Inline{w.Text("· organiser")}},
										},
										[]w.SecItemPart{
											w.Mark{Colour: "#4A6B52", Size: w.MarkH11Inline},
											w.Text(" "),
											w.Link{Text: "Анна Ковалёва", Href: "#"},
											w.Text(" "),
											w.Dim{Parts: []w.Inline{w.Text("· accepted")}},
										},
										[]w.SecItemPart{
											w.Mark{Colour: "#6B4E7A", Size: w.MarkH11Inline},
											w.Text(" "),
											w.Link{Text: "Артём Белов", Href: "#"},
											w.Text(" "),
											w.Dim{Parts: []w.Inline{w.Text("· tentative")}},
										},
										[]w.SecItemPart{w.Mono{Text: "v@example.org"}, w.Text(" "), w.Dim{Parts: []w.Inline{w.Text("· no reply")}}},
									},
									Columns: true,
								},
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
											w.Dim{Parts: []w.Inline{w.Text("text/markdown · 3.1 KB")}},
											w.Text(" "),
											w.Badge{Label: "in the event"},
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
											w.Span{Face: w.FaceAccent, Text: "Заметка со встречи 10.08"},
											w.Text(" "),
											w.Dim{Parts: []w.Inline{w.Text("· note")}},
										},
									},
								},
							},
						},
						w.Sec{
							Items: []w.SecPart{
								w.Rubric{
									Text:  "The same event in other calendars",
									Items: []w.RubricPart{w.Badge{Label: "2", Tone: w.BadgeLinked}},
								},
								w.SecList{
									Items: [][]w.SecItemPart{
										[]w.SecItemPart{
											w.Mark{Colour: "#6B4E7A", Size: w.MarkH11Inline},
											w.Text(" "),
											w.Link{Text: "radicale · vps · Проекты", Href: "#"},
											w.Text(" "),
											w.Dim{Parts: []w.Inline{w.Text("· Holds the same")}},
										},
										[]w.SecItemPart{
											w.Mark{Colour: "#9A9280", Size: w.MarkH11Inline},
											w.Text(" "),
											w.Link{Text: "baikal · Общий", Href: "#"},
											w.Text(" "),
											w.Dim{Parts: []w.Inline{w.Text("· When, where differ")}},
										},
									},
									Gap: 5,
								},
							},
						},
						w.Sec{
							Items: []w.SecPart{
								w.Rubric{Text: "Source"},
								w.Diag{
									Text: "Path  /remote.php/dav/calendars/mix/personal/8f1c.ics\nRead  2 min ago · changed on server 3 days ago",
								},
							},
						},
						w.Danger{
							Button: w.Button{Label: "Delete event", Icon: "trash", Kind: w.ButtonDangerSolid},
							Note:   "Series is edited and deleted whole.",
						},
					},
				},
				Height: w.ShellHeight390,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/publish?c=rabota&uid=aa01",
					Title:  "Publish an event",
					Head: []w.SheetHeadPart{
						w.Dim{Parts: []w.Inline{w.Text("Синк-ап по релизу")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed},
					},
					Body: []w.SheetPart{
						w.Field{
							Label: "What the visitor sees",
							Parts: []w.FieldPart{
								w.Group{
									Items: []w.FieldPart{
										w.Check{
											On:     true,
											Name:   "time",
											Value:  "1",
											Reason: "Always shown",
											Parts:  []w.CheckPart{w.Text("Time"), w.Dim{Parts: []w.Inline{w.Text("always shown")}, Size: w.Dim11_5}},
										},
										w.Check{On: true, Name: "title", Value: "1", Parts: []w.CheckPart{w.Text("Title")}},
										w.Check{Name: "location", Value: "1", Parts: []w.CheckPart{w.Text("Location")}},
										w.Check{Name: "description", Value: "1", Parts: []w.CheckPart{w.Text("Description")}},
										w.Check{Name: "attendees", Value: "1", Parts: []w.CheckPart{w.Text("Attendees")}},
										w.Check{Name: "attachments", Value: "1", Parts: []w.CheckPart{w.Text("Attachments")}},
										w.Check{Name: "tags", Value: "1", Parts: []w.CheckPart{w.Text("Tags")}},
									},
									Form: w.GroupColumnLoose,
								},
								w.FieldHint{Parts: []w.Inline{w.Text("Without the title the event shows as “Busy”.")}},
							},
							Width: w.FieldW100,
						},
						w.SheetChecks{
							Items: []w.Check{
								w.Check{
									Name:  "download",
									Value: "1",
									Parts: []w.CheckPart{w.Text("Allow "), w.Mono{Text: ".ics"}, w.Text(" download and subscription")},
								},
							},
						},
						w.Field{
							Label: "Link lives for",
							Parts: []w.FieldPart{
								w.Drop{
									Trigger: w.Pick{Value: "30 days"},
									Menu: w.Menu{
										Items: []w.MenuPart{
											w.MenuItem{Label: "7 days", Name: "lives", Value: "7"},
											w.MenuItem{Label: "14 days", Name: "lives", Value: "14"},
											w.MenuItem{Label: "30 days", On: true, Name: "lives", Value: "30"},
											w.MenuSep{},
											w.MenuItem{Label: "No expiry", Name: "lives", Value: "never"},
										},
										Width: w.MenuWidth212,
										Pop:   w.MenuUnder,
									},
								},
							},
							Width: w.FieldSM,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Create link", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/calendar"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGap7,
							Place: w.Pushed,
						},
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
						w.Head{
							Photo: &w.HeadPhoto{},
							Title: &w.HeadTitle{Text: "Анна Ковалёва"},
							Subs: []w.HeadSub{
								w.HeadSub{Parts: []w.HeadSubPart{w.Text("Издательство «Слово» · Выпускающий редактор")}},
								w.HeadSub{Parts: []w.HeadSubPart{w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, w.Text("nextcloud · Контакты")}},
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
											w.Mark{Colour: "#A05A2C", Size: w.MarkH11Inline},
											w.Text(" "),
											w.Link{Text: "nextcloud · Архив", Href: "#"},
											w.Text(" "),
											w.Dim{Parts: []w.Inline{w.Text("· Phone, email differ")}},
										},
										[]w.SecItemPart{
											w.Mark{Colour: "#6B4E7A", Size: w.MarkH11Inline},
											w.Text(" "),
											w.Link{Text: "radicale · vps · Команда", Href: "#"},
											w.Text(" "),
											w.Dim{Parts: []w.Inline{w.Text("· Phone, email differ")}},
										},
									},
									Gap: 5,
								},
							},
						},
						w.Sec{
							Items: []w.SecPart{
								w.Rubric{Text: "Timeline", Items: []w.RubricPart{w.Mono{Text: "10", Variant: w.MonoTight}}},
								w.List{
									Rows: []w.ListRow{
										w.LinkedRow{
											Mark:   w.Mark{Colour: "#4A6B52"},
											Icon:   "cal",
											Title:  "Синк-ап по релизу",
											Detail: "10 Aug · attendee · accepted",
											Href:   "/app/calendar/nextcloud/lichnyj/aa01",
										},
										w.LinkedRow{
											Mark:   w.Mark{Colour: "#3F6E8C"},
											Icon:   "task",
											Title:  "Согласовать обложку",
											Detail: "04 Aug · organiser · done",
											Href:   "/app/tasks/nextcloud/rabota/b9d3",
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
				Height: w.ShellHeight340,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/publish?c=kontakty&uid=4f2b-9c31-a0e7",
					Title:  "Publish a contact",
					Head: []w.SheetHeadPart{
						w.Dim{Parts: []w.Inline{w.Text("Анна Ковалёва")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed},
					},
					Body: []w.SheetPart{
						w.Field{
							Label: "What the visitor sees",
							Parts: []w.FieldPart{
								w.Group{
									Items: []w.FieldPart{
										w.Check{On: true, Name: "name", Value: "1", Parts: []w.CheckPart{w.Text("Name and organisation")}},
										w.Check{Name: "phones", Value: "1", Parts: []w.CheckPart{w.Text("Phone numbers")}},
										w.Check{Name: "emails", Value: "1", Parts: []w.CheckPart{w.Text("Email addresses")}},
										w.Check{Name: "photo", Value: "1", Parts: []w.CheckPart{w.Text("Photo")}},
										w.Check{Name: "note", Value: "1", Parts: []w.CheckPart{w.Text("Note")}},
									},
									Form: w.GroupColumnLoose,
								},
							},
							Width: w.FieldW100,
						},
						w.Check{
							Name:  "download",
							Value: "1",
							Parts: []w.CheckPart{w.Text("Allow "), w.Mono{Text: ".vcf"}, w.Text(" download and subscription")},
						},
						w.Field{
							Label: "Link lives for",
							Parts: []w.FieldPart{
								w.Drop{
									Trigger: w.Pick{Value: "7 days"},
									Menu: w.Menu{
										Items: []w.MenuPart{
											w.MenuItem{Label: "7 days", On: true, Name: "lives", Value: "7"},
											w.MenuItem{Label: "14 days", Name: "lives", Value: "14"},
											w.MenuItem{Label: "30 days", Name: "lives", Value: "30"},
											w.MenuSep{},
											w.MenuItem{Label: "No expiry", Name: "lives", Value: "never"},
										},
										Width: w.MenuWidth212,
										Pop:   w.MenuUnder,
									},
								},
							},
							Width: w.FieldSM,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Create link", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/contacts"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGap7,
							Place: w.Pushed,
						},
					},
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
						{Icon: "back", Label: "Back to list", Short: "Back", Href: "#"}}},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2", Current: true},
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
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "Neighbours"},
							w.Tree{
								Rows: []w.TreeRow{
									{Icon: "note", Label: "Заметка со встречи"},
									{Icon: "note", Label: "Что почитать про CRDT"}},
								On:      1,
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
									w.Button{Label: ".md", Icon: "export", Title: "Export Markdown"},
									w.Button{Label: "Send", Icon: "mail"},
									w.Button{Label: "Print", Icon: "print", Tertiary: true},
									w.Button{
										Icon:  "code",
										Title: "Source",
										Href:  "/app/notes/nextcloud/zapisi/9e31-40b7-c012?view=source",
										Swap:  "note-main",
									},
									w.Seg{
										Options: []w.SegOption{
											{Label: "Full width", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012?width=full", On: true},
											{Label: "Reading width", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012?width=read"},
										},
										Swap:     "note-main",
										Tertiary: true,
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
											Title: "Что делаем",
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
											Title: w.Rubric{Text: "Tags", Items: []w.RubricPart{w.Badge{Label: "linked", Tone: w.BadgeLinked}}},
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
											Items: []w.DocmetaPart{
												w.PostForm{
													ID:        "post-attach",
													Action:    "/app/notes/nextcloud/zapisi/9e31-40b7-c012/edit?op=attach",
													Fields:    []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
													Multipart: true,
												},
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
															Mark: &w.Mark{Colour: "#6B4E7A", Size: w.MarkH11Inline},
															Name: "nextcloud · Дневник",
															Href: "#",
															What: w.Dim{Parts: []w.Inline{w.Text("· Text, tags differ")}},
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
				Height: w.ShellHeight340,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/publish?c=zapisi&uid=9e31",
					Title:  "Publish a note",
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
							LabelAside: w.Plain{Text: "2 files · 24.4 MB"},
							Parts: []w.FieldPart{
								w.Group{
									Items: []w.FieldPart{
										w.Check{
											Name:  "attachments",
											Value: "1",
											Parts: []w.CheckPart{w.Text("Publish them too — each gets its own link")},
										},
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
						w.Check{
							Name:  "download",
							Value: "1",
							Parts: []w.CheckPart{w.Text("Allow "), w.Mono{Text: ".md"}, w.Text(" download")},
						},
						w.Field{
							Label: "Link lives for",
							Parts: []w.FieldPart{
								w.Drop{
									Trigger: w.Pick{Value: "30 days"},
									Menu: w.Menu{
										Items: []w.MenuPart{
											w.MenuItem{Label: "7 days", Name: "lives", Value: "7"},
											w.MenuItem{Label: "14 days", Name: "lives", Value: "14"},
											w.MenuItem{Label: "30 days", On: true, Name: "lives", Value: "30"},
											w.MenuSep{},
											w.MenuItem{Label: "No expiry", Name: "lives", Value: "never"},
										},
										Width: w.MenuWidth212,
										Pop:   w.MenuUnder,
									},
								},
							},
							Width: w.FieldSM,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Create link", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/notes"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGap7,
							Place: w.Pushed,
						},
					},
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
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
							{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
							{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
							{Icon: "file", Label: "Files", Href: "#"},
							{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
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
							Subs: []w.HeadSub{
								w.HeadSub{
									Parts: []w.HeadSubPart{w.Text("sftpgo · vps · "), w.Badge{Label: "attachment folder"}, w.Text(" · 3 items")},
								},
							},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "Upload", Menu: true, Icon: "upload", Kind: w.ButtonPrimary},
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
								{Label: "Changed", Width: w.ColTime, Main: true},
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
					Action: "/app/publish?c=documents&f=materialy-k-relizu",
					Title:  "Publish a folder",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("Материалы к релизу · 12 files")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
					},
					Body: []w.SheetPart{
						w.Field{
							Label: "Link lives for",
							Parts: []w.FieldPart{
								w.Drop{
									Trigger: w.Pick{Value: "14 days"},
									Menu: w.Menu{
										Items: []w.MenuPart{
											w.MenuItem{Label: "7 days", Name: "lives", Value: "7"},
											w.MenuItem{Label: "14 days", On: true, Name: "lives", Value: "14"},
											w.MenuItem{Label: "30 days", Name: "lives", Value: "30"},
											w.MenuSep{},
											w.MenuItem{Label: "No expiry", Name: "lives", Value: "never"},
										},
										Width: w.MenuWidth212,
										Pop:   w.MenuUnder,
									},
								},
							},
							Width: w.FieldSM,
						},
						w.Msg{
							Parts: []w.MsgPart{w.Text("The same address mounts as a read-only drive in Windows, macOS and Linux.")},
							Space: w.MsgFlush,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Create link", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/files"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGap7,
							Place: w.Pushed,
						},
					},
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
							Items: []w.LinePart{w.Button{Label: "Download .ics", Icon: "download", Href: "/p/7f3ac9/download"}},
						},
						Wide: true,
					},
					w.Strip{
						Label: "Subscribe in your own client",
						Items: []w.StripPart{
							w.Input{
								Value:    "webcal://carrel.example/p/7f3ac9/personal.ics",
								Label:    "Subscribe in your own client",
								Mono:     true,
								ReadOnly: true,
								Variant:  w.InputLineFill,
							},
							w.Button{Label: "Copy", Small: true, Copy: "webcal://carrel.example/p/7f3ac9/personal.ics", Copied: "Copied"},
							w.Button{Label: "Open in calendar app", Small: true, Href: "webcal://carrel.example/p/7f3ac9/personal.ics"},
							w.Dim{Parts: []w.Inline{w.Text("read-only, updates on its own")}, Size: w.Dim11},
						},
					},
					w.Head{
						Title: &w.HeadTitle{Text: "Личный"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("10 – 31 August 2026")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Filter{
									Query:   w.Query{Placeholder: "Filter by word", Label: "Filter by word", Name: "filter"},
									Waiting: true,
									Action:  "/p/7f3ac9",
									Variant: w.FilterH26W190,
								},
							},
						},
						Pad: w.HeadPadBottom8,
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "10:00",
								Title:  []w.Inline{w.Text("Синк-ап по релизу")},
								Detail: "Переговорная 2",
								Badges: []w.Badge{w.Badge{Label: "4 attendees"}, w.Badge{Label: "repeats weekly"}},
								Href:   "/p/7f3ac9?uid=aa01",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Tuesday · 11 August", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "18:00",
								Title:  []w.Inline{w.Text("Забрать документы")},
								Detail: "МФЦ, Мясницкая 13",
								Href:   "/p/7f3ac9?uid=ab02",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 17 August", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "10:00",
								Title:  []w.Inline{w.Text("Синк-ап по релизу")},
								Detail: "Переговорная 2",
								Badges: []w.Badge{w.Badge{Label: "4 attendees"}, w.Badge{Label: "repeats weekly"}},
								Href:   "/p/7f3ac9?uid=aa01",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 24 August", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "10:00",
								Title:  []w.Inline{w.Text("Синк-ап по релизу")},
								Detail: "Переговорная 2",
								Badges: []w.Badge{w.Badge{Label: "4 attendees"}, w.Badge{Label: "repeats weekly"}},
								Href:   "/p/7f3ac9?uid=aa01",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 31 August", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "10:00",
								Title:  []w.Inline{w.Text("Синк-ап по релизу")},
								Detail: "Переговорная 2",
								Badges: []w.Badge{w.Badge{Label: "4 attendees"}, w.Badge{Label: "repeats weekly"}},
								Href:   "/p/7f3ac9?uid=aa01",
							},
						},
					},
					w.MainFoot{Text: "This link works until 8 September 2026."},
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
							Items: []w.LinePart{w.Button{Label: "Download .vcf", Icon: "download", Href: "/p/2b91f4/download"}},
						},
						Wide: true,
					},
					w.Strip{
						Label: "Address of this book",
						Items: []w.StripPart{
							w.Input{
								Value:    "https://carrel.example/p/2b91f4/team.vcf",
								Label:    "Address of this book",
								Mono:     true,
								ReadOnly: true,
								Variant:  w.InputLineFill,
							},
							w.Button{Label: "Copy", Small: true, Copy: "https://carrel.example/p/2b91f4/team.vcf", Copied: "Copied"},
							w.Dim{
								Parts: []w.Inline{w.Text("another Carrel can subscribe to it; other clients import it")},
								Size:  w.Dim11,
							},
						},
					},
					w.Head{
						Title: &w.HeadTitle{Text: "Команда"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 contacts")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Filter{
									Query:   w.Query{Placeholder: "Filter by word", Label: "Filter by word", Name: "filter"},
									Waiting: true,
									Action:  "/p/2b91f4",
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
								Avatar:   w.Avatar{Initials: "АК"},
								Title:    []w.Inline{w.Text("Анна Ковалёва")},
								Subtitle: "+7 912 445-10-88",
								Detail:   "a.kovaleva@example.org",
								Caption:  "Издательство «Слово»",
								Href:     "/p/2b91f4?uid=c102",
							},
							w.FaceRow{
								Mark:     w.Mark{Colour: "#6B4E7A"},
								Avatar:   w.Avatar{Initials: "АБ"},
								Title:    []w.Inline{w.Text("Артём Белов")},
								Subtitle: "+7 903 118-22-04",
								Detail:   "a.belov@example.org",
								Caption:  "Тензор",
								Href:     "/p/2b91f4?uid=c101",
							},
						},
					},
					w.MainFoot{Text: "This link works until 9 September 2026."},
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
								w.Seg{
									Options: []w.SegOption{
										{Label: "Full width", Href: "/p/9e31a4", On: true},
										{Label: "Reading width", Href: "/p/9e31a4?width=reading"},
									},
								},
								w.Button{Label: "Download .md", Icon: "download", Href: "/p/9e31a4/download"},
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
										Title: "Что делаем",
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
									w.DocFiles{
										Title: w.Rubric{Text: "Attachments", Variant: w.RubricBottom7},
										Items: []w.DocFile{
											{
												Name: "2026-08-10-sync-board.png",
												Size: w.Dim{Parts: []w.Inline{w.Text("· 412 KB")}},
												Href: "/p/Bn4q7Zs1Wd2e",
											},
											{
												Name: "запись-встречи.m4a",
												Size: w.Dim{Parts: []w.Inline{w.Text("· 24 MB")}},
												Href: "/p/Kd9x3Pq6Tr1m",
											},
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

// Publish23 is frame «Публичная страница · заметка, вложения не отданы» (s-publish-23).
func Publish23() Screen {
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
								w.Seg{
									Options: []w.SegOption{
										{Label: "Full width", Href: "/p/9e31a4", On: true},
										{Label: "Reading width", Href: "/p/9e31a4?width=reading"},
									},
								},
								w.Button{Label: "Download .md", Icon: "download", Href: "/p/9e31a4/download"},
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
										Title: "Что делаем",
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
									w.DocNote{Parts: []w.Inline{w.Text("2 attached files are not published.")}},
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
						End: w.Line{
							Gap:   w.LineGap6,
							Place: w.Pushed,
							Items: []w.LinePart{
								w.Button{Label: "Download", Icon: "download", Kind: w.ButtonPrimary, Href: "/p/Bn4q7Zs1Wd2e/download"},
							},
						},
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
								w.Seg{
									Options: []w.SegOption{
										{Label: "Full width", Href: "/p/b72c19", On: true},
										{Label: "Reading width", Href: "/p/b72c19?width=reading"},
									},
								},
								w.Button{Label: "Download", Icon: "download", Href: "/p/b72c19/download"},
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
					w.MainFoot{Text: "This link works until 19 August 2026."},
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
						Download: w.Button{Label: "Download 84.2 MB", Icon: "download", Kind: w.ButtonPrimary, Href: "/p/d0a6f3/download"},
						Note:     "A browser cannot show this file.",
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
							Items: []w.LinePart{w.Button{Label: "Download folder", Icon: "download", Href: "/p/5c40e2/download"}},
						},
					},
					w.Head{
						Title:  &w.HeadTitle{Text: "Материалы к релизу", Size: w.HeadTitle19},
						Crumbs: &w.Crumbs{Current: "Материалы к релизу", Aside: &w.Dim{Parts: []w.Inline{w.Text("· 4 items · 6.6 MB")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Filter{
									Query:   w.Query{Placeholder: "Filter this folder", Label: "Filter this folder", Name: "filter"},
									Waiting: true,
									Action:  "/p/5c40e2",
									Variant: w.FilterH26W190,
								},
							},
						},
						Pad: w.HeadPadBottom8,
					},
					w.Strip{
						Label: "Mount as a drive",
						Items: []w.StripPart{
							w.Input{
								Value:    "https://carrel.example/p/5c40e2/dav/",
								Label:    "Mount as a drive",
								Mono:     true,
								ReadOnly: true,
								Variant:  w.InputLineFill,
							},
							w.Button{Label: "Copy", Small: true, Copy: "https://carrel.example/p/5c40e2/dav/", Copied: "Copied"},
							w.Drop{
								Trigger: w.Button{Label: "How to mount", Menu: true, Small: true},
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuText{Title: "Windows", Text: "File Explorer › This PC › Map network drive, then paste the address."},
										w.MenuText{Title: "macOS", Text: "Finder › Go › Connect to Server, then paste the address."},
										w.MenuText{
											Title: "Linux",
											Text:  "Files › Other Locations › Connect to Server, then type the address with davs:// in place of https://.",
										},
									},
									Width: w.MenuWidth290,
									Pop:   w.MenuUnder,
								},
							},
							w.Dim{Parts: []w.Inline{w.Text("read-only WebDAV, no password")}, Size: w.Dim11},
						},
					},
					w.Table{
						Cols: []w.TableCol{
							{Label: "Name"},
							{Label: "Size", Width: w.ColSize},
							{Label: "Type", Width: w.ColChanged, Rank: w.Rank3},
							{Label: "Changed", Width: w.ColTime, Rank: w.Rank4, Main: true},
							{Width: w.Col90},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Icon: "file", Name: "скриншоты", Href: "/p/5c40e2?in=скриншоты"}}},
									{Parts: []w.CellPart{w.Text("—")}, Mono: true},
									{Parts: []w.CellPart{w.Text("folder · 6 items")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-10 11:07")}, Mono: true},
									{
										Parts: []w.CellPart{w.Button{Label: "Download", Small: true, Href: "/p/5c40e2/download?in=скриншоты"}},
										Acts:  true,
									},
								},
								State: w.RowFolder,
							},
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "release-checklist.md", Href: "/p/5c40e2?file=release-checklist.md"}},
									},
									{Parts: []w.CellPart{w.Text("3.1 KB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("text/markdown")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-09 18:44")}, Mono: true},
									{
										Parts: []w.CellPart{w.Button{Label: "Download", Small: true, Href: "/p/5c40e2/file/release-checklist.md"}},
										Acts:  true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{
											w.TableFile{Icon: "img", Name: "2026-08-04-whiteboard.jpg", Href: "/p/5c40e2?file=2026-08-04-whiteboard.jpg"},
										},
									},
									{Parts: []w.CellPart{w.Text("2.8 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/jpeg")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-04 17:03")}, Mono: true},
									{
										Parts: []w.CellPart{w.Button{Label: "Download", Small: true, Href: "/p/5c40e2/file/2026-08-04-whiteboard.jpg"}},
										Acts:  true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{w.TableFile{Icon: "doc", Name: "смета-подрядчик.pdf", Href: "/p/5c40e2?file=смета-подрядчик.pdf"}},
									},
									{Parts: []w.CellPart{w.Text("1.2 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("application/pdf")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-05 14:20")}, Mono: true},
									{
										Parts: []w.CellPart{w.Button{Label: "Download", Small: true, Href: "/p/5c40e2/file/смета-подрядчик.pdf"}},
										Acts:  true,
									},
								},
							},
						},
					},
					w.MainFoot{Text: "This link works until 1 September 2026."},
				},
			},
		},
	}
}

// Publish24 is frame «Публичная страница · весь календарь» (s-publish-24).
func Publish24() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Banner{
						Logo:  w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"},
						Badge: &w.Badge{Label: "read-only"},
					},
					w.Head{
						Title: &w.HeadTitle{Text: "Личный"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("10 – 16 August 2026")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Filter{
									Query:   w.Query{Placeholder: "Filter by word", Label: "Filter by word", Name: "filter"},
									Waiting: true,
									Action:  "/p/a41d07",
									Variant: w.FilterH26W190,
								},
							},
						},
						Pad: w.HeadPadBottom8,
					},
					w.Bar{
						Items: []w.BarPart{
							w.Button{Icon: "back", Title: "Previous week", Href: "/p/a41d07?from=2026-08-03"},
							w.Button{Label: "Today", Href: "/p/a41d07"},
							w.Button{Icon: "next", Title: "Next week", Href: "/p/a41d07?from=2026-08-17"},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "10:00",
								Title:  []w.Inline{w.Text("Синк-ап по релизу")},
								Detail: "",
								Badges: []w.Badge{w.Badge{Label: "repeats weekly"}},
								Href:   "/p/a41d07?uid=aa01",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Tuesday · 11 August", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "18:00",
								Title:  []w.Inline{w.Text("Забрать документы")},
								Detail: "",
								Href:   "/p/a41d07?uid=ab02",
							},
						},
					},
					w.MainFoot{Text: "This link does not expire."},
				},
			},
		},
	}
}

// Publish39 is frame «Публичная страница · календарь, фильтр» (s-publish-39).
func Publish39() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Banner{
						Logo:  w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"},
						Badge: &w.Badge{Label: "read-only"},
					},
					w.Head{
						Title: &w.HeadTitle{Text: "Личный"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 events with «синк»")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Filter{
									Query:   w.Query{Value: "синк", Label: "Filter by word", Name: "filter"},
									Clear:   "/p/a41d07",
									Action:  "/p/a41d07",
									Variant: w.FilterH26W190,
								},
							},
						},
						Pad: w.HeadPadBottom8,
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 27 July", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "11:00",
								Title:  []w.Inline{w.Text("Синк-ап с подрядчиком")},
								Detail: "",
								Past:   true,
								Href:   "/p/a41d07?uid=a907",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 17 August", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "10:00",
								Title:  []w.Inline{w.Text("Синк-ап по релизу")},
								Detail: "",
								Badges: []w.Badge{w.Badge{Label: "repeats weekly"}},
								Href:   "/p/a41d07?uid=aa01",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Thursday · 3 September", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "15:00",
								Title:  []w.Inline{w.Text("Синк-ап: итоги квартала")},
								Detail: "",
								Href:   "/p/a41d07?uid=ae05",
							},
						},
					},
					w.MainFoot{Text: "This link does not expire."},
				},
			},
		},
	}
}

// Publish25 is frame «Публичная страница · календарь, пустая неделя» (s-publish-25).
func Publish25() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Banner{
						Logo:  w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"},
						Badge: &w.Badge{Label: "read-only"},
					},
					w.Head{
						Title: &w.HeadTitle{Text: "Личный"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("17 – 23 August 2026")}}},
						Pad:   w.HeadPadBottom8,
					},
					w.Bar{
						Items: []w.BarPart{
							w.Button{Icon: "back", Title: "Previous week", Href: "/p/a41d07?from=2026-08-10"},
							w.Button{Label: "Today", Href: "/p/a41d07"},
							w.Button{Icon: "next", Title: "Next week", Href: "/p/a41d07?from=2026-08-24"},
						},
					},
					w.Empty{Parts: []w.EmptyPart{w.Strong("No events this week.")}},
					w.MainFoot{Text: "This link does not expire."},
				},
			},
		},
	}
}

// Publish16 is frame «Публичная страница · событие» (s-publish-16).
func Publish16() Screen {
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
							Items: []w.LinePart{w.Button{Label: "Download .ics", Icon: "download", Href: "/p/c3f1a8/download"}},
						},
					},
					w.Head{
						Title: &w.HeadTitle{Text: "Синк-ап по релизу"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Monday, 10 August 2026, 10:00 – 11:00 · Переговорная 2")}}},
						Pad:   w.HeadPadBottom8,
					},
					w.Fields{
						Items: []w.FieldsItem{
							{Term: "Repeats", Value: []w.ValuePart{w.Text("Weekly on Mon "), w.Badge{Label: "until 28 December"}}},
							{Term: "Status", Value: []w.ValuePart{w.Text("Confirmed")}},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Description"},
							w.Prose{
								Parts: []w.Inline{w.Text("Обсудить оставшиеся блокеры, дату заморозки и кто пишет письмо пользователям.")},
							},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Attendees", Items: []w.RubricPart{w.Mono{Text: "4", Variant: w.MonoTight}}},
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{w.Text("Мария Лебедева "), w.Dim{Parts: []w.Inline{w.Text("· organiser")}}},
									[]w.SecItemPart{w.Text("Анна Ковалёва "), w.Dim{Parts: []w.Inline{w.Text("· accepted")}}},
									[]w.SecItemPart{w.Text("Артём Белов "), w.Dim{Parts: []w.Inline{w.Text("· tentative")}}},
									[]w.SecItemPart{w.Text("Виктор Гончаров "), w.Dim{Parts: []w.Inline{w.Text("· no reply")}}},
								},
								Columns: true,
							},
						},
					},
					w.MainFoot{Text: "This link works until 8 September 2026."},
				},
			},
		},
	}
}

// Publish26 is frame «Публичная страница · контакт» (s-publish-26).
func Publish26() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Banner{
						Logo:  w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"},
						Badge: &w.Badge{Label: "read-only"},
					},
					w.Head{
						Photo: &w.HeadPhoto{Initials: "АК"},
						Title: &w.HeadTitle{Text: "Анна Ковалёва", Size: w.HeadTitle19},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Издательство «Слово» · Выпускающий редактор")}}},
						Pad:   w.HeadPadBottom8,
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
							},
							{
								Term:  "Email · work",
								Value: []w.ValuePart{w.Text("a.kovaleva@example.org")},
								Href:  "mailto:a.kovaleva@example.org",
							},
							{Term: "Note", Value: []w.ValuePart{w.Text("Пишет по вторникам, отвечает в тот же день.")}},
						},
					},
					w.MainFoot{Text: "This link works until 9 September 2026."},
				},
			},
		},
	}
}

// Publish27 is frame «Публичная страница · задача» (s-publish-27).
func Publish27() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Banner{
						Logo:  w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"},
						Badge: &w.Badge{Label: "read-only"},
					},
					w.Head{
						Title: &w.HeadTitle{Text: "Отправить документы в МФЦ"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Due Friday, 7 August 2026")}}},
						Pad:   w.HeadPadBottom8,
					},
					w.Fields{Items: []w.FieldsItem{
						{Term: "Status", Value: []w.ValuePart{w.Text("Needs action")}}}},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Description"},
							w.Prose{Parts: []w.Inline{w.Text("Опись, копия паспорта и заявление. Приём до 18:00, окно 4.")}},
						},
					},
					w.MainFoot{Text: "This link works until 8 September 2026."},
				},
			},
		},
	}
}

// Publish28 is frame «Публичная страница · список задач» (s-publish-28).
func Publish28() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Banner{
						Logo:  w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"},
						Badge: &w.Badge{Label: "read-only"},
					},
					w.Head{
						Title: &w.HeadTitle{Text: "Дела"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("4 tasks")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Filter{
									Query:   w.Query{Placeholder: "Filter by word", Label: "Filter by word", Name: "filter"},
									Waiting: true,
									Action:  "/p/e5b210",
									Variant: w.FilterH26W190,
								},
							},
						},
						Pad: w.HeadPadBottom8,
					},
					w.ListGroup{Title: w.Rubric{Text: "Open", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
					w.List{
						Rows: []w.ListRow{
							w.DueRow{
								Tick:     w.Tick{Sign: true},
								Mark:     w.Mark{Colour: "#4A6B52"},
								Title:    []w.Inline{w.Text("Отправить документы в МФЦ")},
								Due:      "7 Aug",
								Progress: "—",
								Single:   true,
								Overdue:  true,
								Href:     "/p/e5b210?uid=t1",
							},
							w.DueRow{
								Tick:     w.Tick{Sign: true},
								Mark:     w.Mark{Colour: "#4A6B52"},
								Title:    []w.Inline{w.Text("Собрать чемодан")},
								Due:      "14 Aug",
								Progress: "40%",
								Single:   true,
								Href:     "/p/e5b210?uid=t2",
							},
							w.DueRow{
								Tick:     w.Tick{Sign: true},
								Mark:     w.Mark{Colour: "#4A6B52"},
								Title:    []w.Inline{w.Text("Купить адаптер для розеток")},
								Due:      "15 Aug",
								Progress: "—",
								Single:   true,
								Href:     "/p/e5b210?uid=t3",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Done", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
					w.List{
						Rows: []w.ListRow{
							w.DueRow{
								Tick:     w.Tick{Done: true, Sign: true},
								Mark:     w.Mark{Colour: "#4A6B52"},
								Title:    []w.Inline{w.Text("Продлить страховку")},
								Due:      "3 Aug",
								Progress: "100%",
								Single:   true,
								Href:     "/p/e5b210?uid=t4",
							},
						},
					},
					w.MainFoot{Text: "This link works until 8 September 2026."},
				},
			},
		},
	}
}

// Publish29 is frame «Публичная страница · тетрадь» (s-publish-29).
func Publish29() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Banner{
						Logo:  w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"},
						Badge: &w.Badge{Label: "read-only"},
					},
					w.Head{
						Title: &w.HeadTitle{Text: "Записи"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 notes")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Filter{
									Query:   w.Query{Placeholder: "Filter by word", Label: "Filter by word", Name: "filter"},
									Waiting: true,
									Action:  "/p/61d4c7",
									Variant: w.FilterH26W190,
								},
							},
						},
						Pad: w.HeadPadBottom8,
					},
					w.ListGroup{Title: w.Rubric{Text: "August 2026", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Title:   []w.Inline{w.Text("Заметка со встречи 10.08")},
								Excerpt: "Договорились заморозить ветку в среду; письмо пользователям пишет Анна…",
								Date:    "10 Aug",
								Single:  true,
								Href:    "/p/61d4c7?uid=n1",
							},
							w.ExcerptRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Title:   []w.Inline{w.Text("Что почитать про CRDT")},
								Excerpt: "Кляйн, глава 5; статья Шапиро 2011; демка automerge…",
								Date:    "8 Aug",
								Single:  true,
								Href:    "/p/61d4c7?uid=n2",
							},
						},
					},
					w.MainFoot{Text: "This link works until 19 August 2026."},
				},
			},
		},
	}
}

// Publish33 is frame «Публичная страница · PDF» (s-publish-33).
func Publish33() Screen {
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
							Items: []w.LinePart{w.Button{Label: "Download", Icon: "download", Kind: w.ButtonPrimary, Href: "/p/4f60c2/download"}},
						},
					},
					w.Head{
						Title: &w.HeadTitle{Text: "смета-подрядчик.pdf", Size: w.HeadTitle19},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mono{Text: "application/pdf · 1.2 MB · changed 2026-08-05"}}}},
						Pad:   w.HeadPadBottom6,
					},
					w.MainThumbs{Thumb: w.Thumb{Icon: "doc", Form: w.ThumbWide}},
					w.MainFoot{Text: "This link works until 19 August 2026."},
				},
			},
		},
	}
}

// Publish34 is frame «Публичная страница · текст сверх потолка» (s-publish-34).
func Publish34() Screen {
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
						Title:    "server-log-august.txt",
						Meta:     w.Mono{Text: "text/plain · 14.6 MB · changed 2026-08-11"},
						Download: w.Button{Label: "Download 14.6 MB", Icon: "download", Kind: w.ButtonPrimary, Href: "/p/5a2f9e/download"},
						Note:     "Too large to show on the page.",
					},
					w.MainFoot{Text: "This link works until 19 August 2026."},
				},
			},
		},
	}
}

// Publish30 is frame «Публичная страница · папка в папке» (s-publish-30).
func Publish30() Screen {
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
							Items: []w.LinePart{w.Button{Label: "Download folder", Icon: "download", Href: "/p/5c40e2/download?in=скриншоты"}},
						},
					},
					w.Head{
						Title: &w.HeadTitle{Text: "Материалы к релизу", Size: w.HeadTitle19},
						Crumbs: &w.Crumbs{
							Path:    []w.Crumb{{Label: "Материалы к релизу", Href: "#"}},
							Current: "скриншоты",
							Aside:   &w.Dim{Parts: []w.Inline{w.Text("· 6 items · 4.2 MB")}},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Filter{
									Query:   w.Query{Placeholder: "Filter this folder", Label: "Filter this folder", Name: "filter"},
									Waiting: true,
									Action:  "/p/5c40e2",
									Hidden:  []w.Hidden{{Name: "in", Value: "скриншоты"}},
									Variant: w.FilterH26W190,
								},
							},
						},
						Pad: w.HeadPadBottom8,
					},
					w.Strip{
						Label: "Mount as a drive",
						Items: []w.StripPart{
							w.Input{
								Value:    "https://carrel.example/p/5c40e2/dav/",
								Label:    "Mount as a drive",
								Mono:     true,
								ReadOnly: true,
								Variant:  w.InputLineFill,
							},
							w.Button{Label: "Copy", Small: true, Copy: "https://carrel.example/p/5c40e2/dav/", Copied: "Copied"},
							w.Drop{
								Trigger: w.Button{Label: "How to mount", Menu: true, Small: true},
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuText{Title: "Windows", Text: "File Explorer › This PC › Map network drive, then paste the address."},
										w.MenuText{Title: "macOS", Text: "Finder › Go › Connect to Server, then paste the address."},
										w.MenuText{
											Title: "Linux",
											Text:  "Files › Other Locations › Connect to Server, then type the address with davs:// in place of https://.",
										},
									},
									Width: w.MenuWidth290,
									Pop:   w.MenuUnder,
								},
							},
							w.Dim{Parts: []w.Inline{w.Text("read-only WebDAV, no password")}, Size: w.Dim11},
						},
					},
					w.Table{
						Cols: []w.TableCol{
							{Label: "Name"},
							{Label: "Size", Width: w.ColSize},
							{Label: "Type", Width: w.ColChanged, Rank: w.Rank3},
							{Label: "Changed", Width: w.ColTime, Rank: w.Rank4, Main: true},
							{Width: w.Col90},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon: "img",
												Name: "2026-08-10-sync-board.png",
												Href: "/p/5c40e2?file=скриншоты/2026-08-10-sync-board.png",
											},
										},
									},
									{Parts: []w.CellPart{w.Text("412 KB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/png")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-10 11:07")}, Mono: true},
									{
										Parts: []w.CellPart{w.Button{Label: "Download", Small: true, Href: "/p/5c40e2/file/скриншоты/2026-08-10-sync-board.png"}},
										Acts:  true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon: "img",
												Name: "2026-08-10-timeline.png",
												Href: "/p/5c40e2?file=скриншоты/2026-08-10-timeline.png",
											},
										},
									},
									{Parts: []w.CellPart{w.Text("688 KB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/png")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-10 11:05")}, Mono: true},
									{
										Parts: []w.CellPart{w.Button{Label: "Download", Small: true, Href: "/p/5c40e2/file/скриншоты/2026-08-10-timeline.png"}},
										Acts:  true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{
											w.TableFile{Icon: "img", Name: "export-dialog.png", Href: "/p/5c40e2?file=скриншоты/export-dialog.png"},
										},
									},
									{Parts: []w.CellPart{w.Text("540 KB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/png")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-09 17:40")}, Mono: true},
									{
										Parts: []w.CellPart{w.Button{Label: "Download", Small: true, Href: "/p/5c40e2/file/скриншоты/export-dialog.png"}},
										Acts:  true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{
											w.TableFile{Icon: "img", Name: "import-wizard.png", Href: "/p/5c40e2?file=скриншоты/import-wizard.png"},
										},
									},
									{Parts: []w.CellPart{w.Text("1.1 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/png")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-09 17:32")}, Mono: true},
									{
										Parts: []w.CellPart{w.Button{Label: "Download", Small: true, Href: "/p/5c40e2/file/скриншоты/import-wizard.png"}},
										Acts:  true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{
											w.TableFile{Icon: "img", Name: "publish-sheet.png", Href: "/p/5c40e2?file=скриншоты/publish-sheet.png"},
										},
									},
									{Parts: []w.CellPart{w.Text("730 KB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/png")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-09 17:20")}, Mono: true},
									{
										Parts: []w.CellPart{w.Button{Label: "Download", Small: true, Href: "/p/5c40e2/file/скриншоты/publish-sheet.png"}},
										Acts:  true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{
											w.TableFile{Icon: "img", Name: "search-results.png", Href: "/p/5c40e2?file=скриншоты/search-results.png"},
										},
									},
									{Parts: []w.CellPart{w.Text("790 KB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/png")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-09 17:11")}, Mono: true},
									{
										Parts: []w.CellPart{w.Button{Label: "Download", Small: true, Href: "/p/5c40e2/file/скриншоты/search-results.png"}},
										Acts:  true,
									},
								},
							},
						},
					},
					w.MainFoot{Text: "This link works until 1 September 2026."},
				},
			},
		},
	}
}

// Publish31 is frame «Публичная страница · папка сверх потолка» (s-publish-31).
func Publish31() Screen {
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
							Items: []w.LinePart{w.Button{Label: "Download folder", Icon: "download", Href: "/p/e93b06/download"}},
						},
					},
					w.Head{
						Title:  &w.HeadTitle{Text: "Фото с выезда", Size: w.HeadTitle19},
						Crumbs: &w.Crumbs{Current: "Фото с выезда", Aside: &w.Dim{Parts: []w.Inline{w.Text("· 2 340 items · 7.6 GB")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Filter{
									Query:   w.Query{Placeholder: "Filter this folder", Label: "Filter this folder", Name: "filter"},
									Waiting: true,
									Action:  "/p/e93b06",
									Variant: w.FilterH26W190,
								},
							},
						},
						Pad: w.HeadPadBottom8,
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Strong("Showing 500 of 2 340."),
							w.Text(" Filter by name, download the folder or mount it to get all of them."),
						},
					},
					w.Strip{
						Label: "Mount as a drive",
						Items: []w.StripPart{
							w.Input{
								Value:    "https://carrel.example/p/e93b06/dav/",
								Label:    "Mount as a drive",
								Mono:     true,
								ReadOnly: true,
								Variant:  w.InputLineFill,
							},
							w.Button{Label: "Copy", Small: true, Copy: "https://carrel.example/p/e93b06/dav/", Copied: "Copied"},
							w.Drop{
								Trigger: w.Button{Label: "How to mount", Menu: true, Small: true},
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuText{Title: "Windows", Text: "File Explorer › This PC › Map network drive, then paste the address."},
										w.MenuText{Title: "macOS", Text: "Finder › Go › Connect to Server, then paste the address."},
										w.MenuText{
											Title: "Linux",
											Text:  "Files › Other Locations › Connect to Server, then type the address with davs:// in place of https://.",
										},
									},
									Width: w.MenuWidth290,
									Pop:   w.MenuUnder,
								},
							},
							w.Dim{Parts: []w.Inline{w.Text("read-only WebDAV, no password")}, Size: w.Dim11},
						},
					},
					w.Table{
						Cols: []w.TableCol{
							{Label: "Name"},
							{Label: "Size", Width: w.ColSize},
							{Label: "Type", Width: w.ColChanged, Rank: w.Rank3},
							{Label: "Changed", Width: w.ColTime, Rank: w.Rank4, Main: true},
							{Width: w.Col90},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{w.TableFile{Icon: "img", Name: "IMG_0001.jpg", Href: "/p/e93b06?file=IMG_0001.jpg"}},
									},
									{Parts: []w.CellPart{w.Text("3.4 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/jpeg")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-06-01 09:12")}, Mono: true},
									{
										Parts: []w.CellPart{w.Button{Label: "Download", Small: true, Href: "/p/e93b06/file/IMG_0001.jpg"}},
										Acts:  true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{w.TableFile{Icon: "img", Name: "IMG_0002.jpg", Href: "/p/e93b06?file=IMG_0002.jpg"}},
									},
									{Parts: []w.CellPart{w.Text("3.1 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/jpeg")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-06-01 09:12")}, Mono: true},
									{
										Parts: []w.CellPart{w.Button{Label: "Download", Small: true, Href: "/p/e93b06/file/IMG_0002.jpg"}},
										Acts:  true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{w.TableFile{Icon: "img", Name: "IMG_0003.jpg", Href: "/p/e93b06?file=IMG_0003.jpg"}},
									},
									{Parts: []w.CellPart{w.Text("2.9 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/jpeg")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-06-01 09:13")}, Mono: true},
									{
										Parts: []w.CellPart{w.Button{Label: "Download", Small: true, Href: "/p/e93b06/file/IMG_0003.jpg"}},
										Acts:  true,
									},
								},
							},
						},
					},
					w.MainFoot{Text: "This link works until 1 September 2026."},
				},
			},
		},
	}
}

// Publish35 is frame «Публичная страница · пустая книга» (s-publish-35).
func Publish35() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Banner{
						Logo:  w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"},
						Badge: &w.Badge{Label: "read-only"},
					},
					w.Head{
						Title: &w.HeadTitle{Text: "Рассылки"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("0 contacts")}}},
						Pad:   w.HeadPadBottom8,
					},
					w.Empty{Parts: []w.EmptyPart{w.Strong("This book is empty.")}},
					w.MainFoot{Text: "This link works until 9 September 2026."},
				},
			},
		},
	}
}

// Publish36 is frame «Публичная страница · пустая папка» (s-publish-36).
func Publish36() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Banner{
						Logo:  w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"},
						Badge: &w.Badge{Label: "read-only"},
					},
					w.Head{
						Title:  &w.HeadTitle{Text: "Черновики", Size: w.HeadTitle19},
						Crumbs: &w.Crumbs{Current: "Черновики", Aside: &w.Dim{Parts: []w.Inline{w.Text("· 0 items")}}},
						Pad:    w.HeadPadBottom8,
					},
					w.Strip{
						Label: "Mount as a drive",
						Items: []w.StripPart{
							w.Input{
								Value:    "https://carrel.example/p/7d1e88/dav/",
								Label:    "Mount as a drive",
								Mono:     true,
								ReadOnly: true,
								Variant:  w.InputLineFill,
							},
							w.Button{Label: "Copy", Small: true, Copy: "https://carrel.example/p/7d1e88/dav/", Copied: "Copied"},
							w.Drop{
								Trigger: w.Button{Label: "How to mount", Menu: true, Small: true},
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuText{Title: "Windows", Text: "File Explorer › This PC › Map network drive, then paste the address."},
										w.MenuText{Title: "macOS", Text: "Finder › Go › Connect to Server, then paste the address."},
										w.MenuText{
											Title: "Linux",
											Text:  "Files › Other Locations › Connect to Server, then type the address with davs:// in place of https://.",
										},
									},
									Width: w.MenuWidth290,
									Pop:   w.MenuUnder,
								},
							},
							w.Dim{Parts: []w.Inline{w.Text("read-only WebDAV, no password")}, Size: w.Dim11},
						},
					},
					w.Empty{Parts: []w.EmptyPart{w.Strong("This folder is empty.")}},
					w.MainFoot{Text: "This link works until 1 September 2026."},
				},
			},
		},
	}
}

// Publish37 is frame «Публичная страница · источник не ответил» (s-publish-37).
func Publish37() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Banner{
						Logo:  w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"},
						Badge: &w.Badge{Label: "read-only"},
					},
					w.Head{Title: &w.HeadTitle{Text: "Личный"}, Pad: w.HeadPadBottom8},
					w.Msg{
						Parts: []w.MsgPart{w.Strong("The source did not answer."), w.Text(" Try again later.")},
						Tone:  w.MsgAlert,
					},
					w.MainFoot{Text: "This link works until 8 September 2026."},
				},
			},
		},
	}
}

// Publish38 is frame «Публичная страница · опубликованного больше нет» (s-publish-38).
func Publish38() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Banner{
						Logo:  w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"},
						Badge: &w.Badge{Label: "read-only"},
					},
					w.Head{Title: &w.HeadTitle{Text: "отчёт-июль.pdf"}, Pad: w.HeadPadBottom8},
					w.Msg{Parts: []w.MsgPart{w.Strong("No longer there."), w.Text(" It was deleted or moved.")}},
				},
			},
		},
	}
}
