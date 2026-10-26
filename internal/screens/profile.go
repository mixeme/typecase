// Раздел макета: 7.11 Настройки · подключения и подписки.

package screens

import w "carrel.dev/typecase/widget"

// Profile1 is frame «Настройки · подключения и подписки» (s-profile-1).
func Profile1() Screen {
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
											{Label: "Changed on the server", Width: w.ColChanged},
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
														Parts: []w.CellPart{
															w.Drop{
																Open:    true,
																Trigger: w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
																Menu: w.Menu{
																	Items: []w.MenuPart{
																		w.MenuItem{Icon: "pencil", Label: "Rename and colour"},
																		w.MenuItem{Icon: "move", Label: "Move to another server"},
																		w.MenuItem{Icon: "dup", Label: "Copy to another server"},
																		w.MenuItem{Icon: "export", Detail: &w.Mono{Text: ".ics", Variant: w.Mono11}, Label: "Export"},
																		w.MenuItem{Icon: "link", Label: "Publish a link"},
																		w.MenuSep{},
																		w.MenuItem{Icon: "trash", Label: "Delete on the server…", Tone: w.MenuToneDanger},
																	},
																	Width: w.MenuWidth212,
																	Pop:   w.MenuUnderEnd,
																},
															},
														},
														Acts: true,
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
											{Label: "Changed on the server", Width: w.ColChanged},
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
											{Label: "Changed on the server", Width: w.ColChanged},
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
		},
	}
}

// Profile27 is frame «Подключения · после Update collection list» (s-profile-27).
func Profile27() Screen {
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
					w.Msg{
						Parts: []w.MsgPart{w.Strong("nextcloud:"), w.Text(" 1 new collection — «Поездки». Tick it to see it in the rail.")},
						Space: w.MsgInset20,
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
											{Label: "Changed on the server", Width: w.ColChanged},
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
													{Pick: ptr(w.TickedOff), Send: w.Send{Form: "conn-picks", Name: "pick", Value: "nextcloud|poezdki"}},
													{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Поездки"}}},
													{Parts: []w.CellPart{w.Text("calendar "), w.Badge{Label: "new"}}, Mono: true},
													{Parts: []w.CellPart{w.Text("0")}, Mono: true},
													{Parts: []w.CellPart{w.Text("just now")}, Mono: true},
													{
														Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
														Acts:  true,
													},
												},
												State: w.RowOff,
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
											{Label: "Changed on the server", Width: w.ColChanged},
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
											{Label: "Changed on the server", Width: w.ColChanged},
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
		},
	}
}

// Profile25 is frame «Подключения · ни одного» (s-profile-25).
func Profile25() Screen {
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
								{Icon: "server", Label: "Connections", Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Href: "#"},
								{Icon: "archive", Label: "Backup", Href: "#"},
								{Icon: "link", Label: "Published links", Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 2,
						},
					},
				},
				w.RailFoot{About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"}},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Connections"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Signed in as "), w.Strong("mix")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Connect account", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Subscribe to a feed", Icon: "link"},
							},
						},
					},
					w.Empty{Parts: []w.EmptyPart{w.Strong("Nothing is connected yet.")}},
				},
			},
		},
	}
}

// Profile26 is frame «Отключить аккаунт · лист поверх Connections» (s-profile-26).
func Profile26() Screen {
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
												{Label: "Changed on the server", Width: w.ColChanged},
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
												{Label: "Changed on the server", Width: w.ColChanged},
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
												{Label: "Changed on the server", Width: w.ColChanged},
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
				Height: w.ShellHeight560,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/settings/connections/disconnect?account=nextcloud",
					Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					Icon:   "server",
					Tone:   w.SheetIconDanger,
					Title:  "Disconnect",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("nextcloud · cloud.example.org")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
					},
					Body: []w.SheetPart{
						w.Msg{
							Parts: []w.MsgPart{w.Text("Nothing is deleted on cloud.example.org. Carrel removes:")},
							Tone:  w.MsgAlert,
							Space: w.MsgFlush,
						},
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("• 10 collections"),
								w.Break{},
								w.Text("• 1 published link"),
								w.Break{},
								w.Text("• 1 source of a backup job"),
								w.Break{},
								w.Text("• 1 davloom device’s sync of them"),
							},
							Space: w.MsgFlush,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Disconnect", Kind: w.ButtonDangerSolid, Submit: true},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/connections"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGap7,
							Place: w.Pushed,
						},
					},
					Width: w.SheetWidth480,
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Profile2 is frame «Подключение аккаунта» (s-profile-2).
func Profile2() Screen {
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
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{
								Title: &w.HeadTitle{Text: "Connect account"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("One connection gives calendars and address books alike.")}}},
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Label",
										Parts: []w.FieldPart{w.Input{Name: "label", Placeholder: "optional", Label: "Label"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Server URL",
										Parts: []w.FieldPart{
											w.Input{
												Name:  "server",
												Value: "https://dav.office.example/dav.php/",
												Label: "Server URL",
												Mono:  true,
												Focus: true,
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Mono{Text: ".well-known", Form: w.MonoCode},
													w.Text(" is tried first, then the URL exactly as typed."),
												},
											},
										},
										Width: w.FieldFull,
									},
									w.Field{
										Label: "Username",
										Parts: []w.FieldPart{w.Input{Name: "login", Value: "m.yenuchenko", Label: "Username"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Password",
										Parts: []w.FieldPart{
											w.Input{Kind: w.InputPassword, Name: "password", Value: "xxxxxxxxxx", Label: "Password"},
											w.Check{
												Name:     "noPassword",
												Value:    "1",
												Disables: "password",
												Parts:    []w.CheckPart{w.Text("The server asks no password")},
											},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label:      "Certificate fingerprint",
										LabelAside: w.Dim{Parts: []w.Inline{w.Text("SHA-256 · optional")}},
										Parts: []w.FieldPart{
											w.Input{
												Name:        "certFingerprint",
												Placeholder: "only for a certificate the system does not trust",
												Label:       "Certificate fingerprint",
												Mono:        true,
											},
										},
										Width: w.FieldFull,
									},
								},
								ID:      "connect",
								Action:  "/app/settings/connections/new",
								Fields:  []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
								Variant: w.FormFlush,
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Connect", Kind: w.ButtonPrimary, Form: "connect"},
									w.Button{Label: "Test only", To: "/app/settings/connections/new/test", Submit: true},
									w.Button{Label: "Cancel", Href: "/app/settings/connections"},
								},
								Variant: w.FormFootFlush,
							},
						},
						Variant: w.Dialog540,
					},
				},
			},
		},
	}
}

// Profile24 is frame «Подключение · сервер отказал во входе» (s-profile-24).
func Profile24() Screen {
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
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{
								Title: &w.HeadTitle{Text: "Connect account"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("One connection gives calendars and address books alike.")}}},
							},
							w.Msg{
								Parts: []w.MsgPart{
									w.Strong("Could not connect."),
									w.Text(" The server refused the login and password. Nothing was saved."),
								},
								Tone:  w.MsgAlert,
								Space: w.MsgFlushSides,
							},
							w.Diag{
								Text: "Discovery trace\nwell-known · ok · GET /.well-known/caldav → 301 → /dav.php/\nprincipal · bad · PROPFIND /dav.php/ → 401 Unauthorized",
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Label",
										Parts: []w.FieldPart{w.Input{Name: "label", Placeholder: "optional", Label: "Label"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Server URL",
										Parts: []w.FieldPart{
											w.Input{Name: "server", Value: "https://dav.office.example/dav.php/", Label: "Server URL", Mono: true},
											w.FieldHint{
												Parts: []w.Inline{
													w.Mono{Text: ".well-known", Form: w.MonoCode},
													w.Text(" is tried first, then the URL exactly as typed."),
												},
											},
										},
										Width: w.FieldFull,
									},
									w.Field{
										Label: "Username",
										Parts: []w.FieldPart{w.Input{Name: "login", Value: "m.yenuchenko", Label: "Username"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Password",
										Parts: []w.FieldPart{
											w.Input{Kind: w.InputPassword, Name: "password", Label: "Password", Error: true},
											w.Check{
												Name:     "noPassword",
												Value:    "1",
												Disables: "password",
												Parts:    []w.CheckPart{w.Text("The server asks no password")},
											},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label:      "Certificate fingerprint",
										LabelAside: w.Dim{Parts: []w.Inline{w.Text("SHA-256 · optional")}},
										Parts: []w.FieldPart{
											w.Input{
												Name:        "certFingerprint",
												Placeholder: "only for a certificate the system does not trust",
												Label:       "Certificate fingerprint",
												Mono:        true,
											},
										},
										Width: w.FieldFull,
									},
								},
								ID:      "connect",
								Action:  "/app/settings/connections/new",
								Fields:  []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
								Variant: w.FormFlush,
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Connect", Kind: w.ButtonPrimary, Form: "connect"},
									w.Button{Label: "Test only", To: "/app/settings/connections/new/test", Submit: true},
									w.Button{Label: "Cancel", Href: "/app/settings/connections"},
								},
								Variant: w.FormFootFlush,
							},
						},
						Variant: w.Dialog540,
					},
				},
			},
		},
	}
}

// Profile14 is frame «Подключение аккаунта · сервер без пароля» (s-profile-14).
func Profile14() Screen {
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
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{
								Title: &w.HeadTitle{Text: "Connect account"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("One connection gives calendars and address books alike.")}}},
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Label",
										Parts: []w.FieldPart{w.Input{Name: "label", Placeholder: "optional", Label: "Label"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Server URL",
										Parts: []w.FieldPart{
											w.Input{Name: "server", Value: "https://xandikos.home.example/", Label: "Server URL", Mono: true},
											w.FieldHint{
												Parts: []w.Inline{
													w.Mono{Text: ".well-known", Form: w.MonoCode},
													w.Text(" is tried first, then the URL exactly as typed."),
												},
											},
										},
										Width: w.FieldFull,
									},
									w.Field{
										Label: "Username",
										Parts: []w.FieldPart{w.Input{Name: "login", Value: "mix", Label: "Username"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Password",
										Parts: []w.FieldPart{
											w.Input{Kind: w.InputPassword, Name: "password", Label: "Password", Disabled: true},
											w.Check{
												On:       true,
												Name:     "noPassword",
												Value:    "1",
												Disables: "password",
												Parts:    []w.CheckPart{w.Text("The server asks no password")},
											},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label:      "Certificate fingerprint",
										LabelAside: w.Dim{Parts: []w.Inline{w.Text("SHA-256 · optional")}},
										Parts: []w.FieldPart{
											w.Input{
												Name:        "certFingerprint",
												Placeholder: "only for a certificate the system does not trust",
												Label:       "Certificate fingerprint",
												Mono:        true,
											},
										},
										Width: w.FieldFull,
									},
								},
								ID:      "connect",
								Action:  "/app/settings/connections/new",
								Fields:  []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
								Variant: w.FormFlush,
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Connect", Kind: w.ButtonPrimary, Form: "connect"},
									w.Button{Label: "Test only", To: "/app/settings/connections/new/test", Submit: true},
									w.Button{Label: "Cancel"},
								},
								Variant: w.FormFootFlush,
							},
						},
						Variant: w.Dialog540,
					},
				},
			},
		},
	}
}

// Profile16 is frame «Подключение · правка» (s-profile-16).
func Profile16() Screen {
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
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{
								Title: &w.HeadTitle{Text: "Edit connection"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("One connection gives calendars and address books alike.")}}},
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Label",
										Parts: []w.FieldPart{w.Input{Name: "label", Value: "nextcloud", Label: "Label"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Server URL",
										Parts: []w.FieldPart{
											w.Input{Name: "server", Value: "https://cloud.example.org/remote.php/dav/", Label: "Server URL", Mono: true},
											w.FieldHint{
												Parts: []w.Inline{
													w.Mono{Text: ".well-known", Form: w.MonoCode},
													w.Text(" is tried first, then the URL exactly as typed."),
												},
											},
										},
										Width: w.FieldFull,
									},
									w.Field{
										Label: "Username",
										Parts: []w.FieldPart{w.Input{Name: "login", Value: "mix", Label: "Username"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Password",
										Parts: []w.FieldPart{
											w.Input{Kind: w.InputPassword, Name: "password", Label: "Password"},
											w.FieldHint{Parts: []w.Inline{w.Text("Leave it empty to keep the stored one.")}},
											w.Check{
												Name:     "noPassword",
												Value:    "1",
												Disables: "password",
												Parts:    []w.CheckPart{w.Text("The server asks no password")},
											},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label:      "Certificate fingerprint",
										LabelAside: w.Dim{Parts: []w.Inline{w.Text("SHA-256 · optional")}},
										Parts: []w.FieldPart{
											w.Input{
												Name:        "certFingerprint",
												Placeholder: "only for a certificate the system does not trust",
												Label:       "Certificate fingerprint",
												Mono:        true,
											},
										},
										Width: w.FieldFull,
									},
								},
								ID:      "connection",
								Action:  "/app/settings/connections/nextcloud/edit",
								Fields:  []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
								Variant: w.FormFlush,
							},
							w.FormFoot{
								Items:   []w.FormFootPart{w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "connection"}, w.Button{Label: "Cancel"}},
								Variant: w.FormFootFlush,
							},
						},
						Variant: w.Dialog540,
					},
				},
			},
		},
	}
}

// Profile22 is frame «Подключение · правка, другой адрес и опубликованные ссылки» (s-profile-22).
func Profile22() Screen {
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
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{
								Title: &w.HeadTitle{Text: "Edit connection"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("One connection gives calendars and address books alike.")}}},
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Label",
										Parts: []w.FieldPart{w.Input{Name: "label", Value: "nextcloud", Label: "Label"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Server URL",
										Parts: []w.FieldPart{
											w.Input{Name: "server", Value: "https://cloud.home.example/remote.php/dav/", Label: "Server URL", Mono: true},
											w.FieldHint{
												Parts: []w.Inline{
													w.Mono{Text: ".well-known", Form: w.MonoCode},
													w.Text(" is tried first, then the URL exactly as typed."),
												},
											},
										},
										Width: w.FieldFull,
									},
									w.Field{
										Label: "Username",
										Parts: []w.FieldPart{w.Input{Name: "login", Value: "mix", Label: "Username"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Password",
										Parts: []w.FieldPart{
											w.Input{Kind: w.InputPassword, Name: "password", Label: "Password"},
											w.FieldHint{Parts: []w.Inline{w.Text("Leave it empty to keep the stored one.")}},
											w.Check{
												Name:     "noPassword",
												Value:    "1",
												Disables: "password",
												Parts:    []w.CheckPart{w.Text("The server asks no password")},
											},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label:      "Certificate fingerprint",
										LabelAside: w.Dim{Parts: []w.Inline{w.Text("SHA-256 · optional")}},
										Parts: []w.FieldPart{
											w.Input{
												Name:        "certFingerprint",
												Placeholder: "only for a certificate the system does not trust",
												Label:       "Certificate fingerprint",
												Mono:        true,
											},
										},
										Width: w.FieldFull,
									},
									w.Field{
										Changed: []string{"server.url", "username"},
										Label:   "Published links",
										Parts: []w.FieldPart{
											w.Check{
												Name:  "published.links",
												Value: "keep",
												Radio: true,
												Parts: []w.CheckPart{w.Text("Keep them: they open on what the connection points at now")},
											},
											w.Check{Name: "published.links", Value: "revoke", Radio: true, Parts: []w.CheckPart{w.Text("Revoke them")}},
											w.FieldHint{Parts: []w.Inline{w.Text("1 published link uses this connection.")}},
										},
										Width: w.FieldFull,
									},
								},
								ID:      "connection-links",
								Action:  "/app/settings/connections/nextcloud/edit",
								Fields:  []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
								Variant: w.FormFlush,
							},
							w.FormFoot{
								Items:   []w.FormFootPart{w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "connection-links"}, w.Button{Label: "Cancel"}},
								Variant: w.FormFootFlush,
							},
						},
						Variant: w.Dialog540,
					},
				},
			},
		},
	}
}

// Profile23 is frame «Подключение · Save без ответа о ссылках» (s-profile-23).
func Profile23() Screen {
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
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{
								Title: &w.HeadTitle{Text: "Edit connection"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("One connection gives calendars and address books alike.")}}},
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Label",
										Parts: []w.FieldPart{w.Input{Name: "label", Value: "nextcloud", Label: "Label"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Server URL",
										Parts: []w.FieldPart{
											w.Input{Name: "server", Value: "https://cloud.home.example/remote.php/dav/", Label: "Server URL", Mono: true},
											w.FieldHint{
												Parts: []w.Inline{
													w.Mono{Text: ".well-known", Form: w.MonoCode},
													w.Text(" is tried first, then the URL exactly as typed."),
												},
											},
										},
										Width: w.FieldFull,
									},
									w.Field{
										Label: "Username",
										Parts: []w.FieldPart{w.Input{Name: "login", Value: "mix", Label: "Username"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Password",
										Parts: []w.FieldPart{
											w.Input{Kind: w.InputPassword, Name: "password", Label: "Password"},
											w.FieldHint{Parts: []w.Inline{w.Text("Leave it empty to keep the stored one.")}},
											w.Check{
												Name:     "noPassword",
												Value:    "1",
												Disables: "password",
												Parts:    []w.CheckPart{w.Text("The server asks no password")},
											},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label:      "Certificate fingerprint",
										LabelAside: w.Dim{Parts: []w.Inline{w.Text("SHA-256 · optional")}},
										Parts: []w.FieldPart{
											w.Input{
												Name:        "certFingerprint",
												Placeholder: "only for a certificate the system does not trust",
												Label:       "Certificate fingerprint",
												Mono:        true,
											},
										},
										Width: w.FieldFull,
									},
									w.Field{
										Changed: []string{"server.url", "username"},
										Label:   "Published links",
										Parts: []w.FieldPart{
											w.Check{
												Name:  "published.links",
												Value: "keep",
												Radio: true,
												Parts: []w.CheckPart{w.Text("Keep them: they open on what the connection points at now")},
											},
											w.Check{Name: "published.links", Value: "revoke", Radio: true, Parts: []w.CheckPart{w.Text("Revoke them")}},
											w.FieldHint{Parts: []w.Inline{w.Text("1 published link uses this connection.")}},
											w.FieldHint{Parts: []w.Inline{w.Text("Choose what becomes of them before saving.")}, Variant: w.HintDanger},
										},
										Width: w.FieldFull,
									},
								},
								ID:      "connection-links",
								Action:  "/app/settings/connections/nextcloud/edit",
								Fields:  []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
								Variant: w.FormFlush,
							},
							w.FormFoot{
								Items:   []w.FormFootPart{w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "connection-links"}, w.Button{Label: "Cancel"}},
								Variant: w.FormFootFlush,
							},
						},
						Variant: w.Dialog540,
					},
				},
			},
		},
	}
}

// Profile11 is frame «Подключение аккаунта · подпись занята» (s-profile-11).
func Profile11() Screen {
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
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{
								Title: &w.HeadTitle{Text: "Connect account"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("One connection gives calendars and address books alike.")}}},
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Label",
										Parts: []w.FieldPart{
											w.Input{Name: "label", Value: "nextcloud", Label: "Label", Focus: true, Error: true},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("There is already a connection called “nextcloud”. Choose another label: it is how your sources are told apart."),
												},
												Variant: w.HintDanger,
											},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Server URL",
										Parts: []w.FieldPart{
											w.Input{Name: "server", Value: "https://cloud.example.org/", Label: "Server URL", Mono: true},
											w.FieldHint{
												Parts: []w.Inline{
													w.Mono{Text: ".well-known", Form: w.MonoCode},
													w.Text(" is tried first, then the URL exactly as typed."),
												},
											},
										},
										Width: w.FieldFull,
									},
									w.Field{
										Label: "Username",
										Parts: []w.FieldPart{w.Input{Name: "login", Value: "mix.work", Label: "Username"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Password",
										Parts: []w.FieldPart{
											w.Input{Kind: w.InputPassword, Name: "password", Value: "xxxxxxxxxx", Label: "Password"},
											w.Check{
												Name:     "noPassword",
												Value:    "1",
												Disables: "password",
												Parts:    []w.CheckPart{w.Text("The server asks no password")},
											},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label:      "Certificate fingerprint",
										LabelAside: w.Dim{Parts: []w.Inline{w.Text("SHA-256 · optional")}},
										Parts: []w.FieldPart{
											w.Input{
												Name:        "certFingerprint",
												Placeholder: "only for a certificate the system does not trust",
												Label:       "Certificate fingerprint",
												Mono:        true,
											},
										},
										Width: w.FieldFull,
									},
								},
								ID:      "connect",
								Action:  "/app/settings/connections/new",
								Fields:  []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
								Variant: w.FormFlush,
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Connect", Kind: w.ButtonPrimary, Form: "connect"},
									w.Button{Label: "Test only", To: "/app/settings/connections/new/test", Submit: true},
									w.Button{Label: "Cancel"},
								},
								Variant: w.FormFootFlush,
							},
						},
						Variant: w.Dialog540,
					},
				},
			},
		},
	}
}

// Profile3 is frame «Show password · вопрос» (s-profile-3).
func Profile3() Screen {
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
							Title: &w.HeadTitle{Text: "Connections", Size: w.HeadTitle19},
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("4 servers · 15 collections · 2 subscriptions")}}},
							Pad:   w.HeadPadBottom8,
						},
						w.Pad{
							Items: []w.PadPart{
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "nextcloud",
											Items: []w.RubricPart{
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "Show password", Small: true, Focus: true},
														w.Button{Label: "Test", Small: true},
														w.Button{Label: "Edit", Small: true},
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
									},
								},
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "sftpgo · vps",
											Items: []w.RubricPart{
												w.Line{
													Items: []w.LinePart{w.Button{Label: "Show password", Small: true}, w.Button{Label: "Test", Small: true}},
													Gap:   w.LineGap6,
												},
											},
										},
										w.Mono{
											Text:    "https://files.example.org/dav/ · mix · last answer 5 min ago",
											Form:    w.MonoLine,
											Variant: w.Mono11_5Dusty,
										},
									},
								},
							},
						},
					},
				},
				Height: w.ShellHeight280,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/settings/unlock",
					Fields: []w.Hidden{
						{Name: "csrf", Value: "7f3a91de"},
						{Name: "back", Value: "/app/settings/connections"},
						{Name: "for", Value: "nextcloud"},
					},
					Title: "Show password",
					Head:  []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("nextcloud")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}},
					Body: []w.SheetPart{
						w.Field{
							Label: "Your Carrel password",
							Parts: []w.FieldPart{
								w.Input{Kind: w.InputPassword, Name: "password", Label: "Your Carrel password", Focus: true, Required: true},
							},
							Width: w.FieldW100,
						},
						w.Dim{
							Parts:   []w.Inline{w.Text("Unlocks stored credentials for five minutes.")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Show", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/connections"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGap7,
							Place: w.Pushed,
						},
					},
					Width: w.SheetWidth400,
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Profile4 is frame «Show password · показано» (s-profile-4).
func Profile4() Screen {
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
								Parts: []w.HeadSubPart{
									w.Text("Signed in as "),
									w.Strong("mix"),
									w.Text(" · 4 servers · 19 collections · 2 subscriptions "),
									w.Badge{Label: "credentials unlocked · 4:12", Accent: true},
									w.Button{Label: "Lock now", Small: true},
								},
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
													w.Button{Label: "Hide password", Small: true, Href: "/app/settings/connections"},
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
									w.CardRow{
										Kind: w.CardTop8,
										Gap:  8,
										Items: []w.CardRowPart{
											w.Input{Value: "z3-brick-lamp-9f2", Label: "Password", Mono: true, ReadOnly: true, Variant: w.InputKeySmFill},
											w.Button{Label: "Copy", Small: true, CopyFrom: "pw-nextcloud", Copied: "Copied"},
										},
									},
									w.Table{
										Cols: []w.TableCol{
											{Width: w.ColTickSmall},
											{Label: "Collection"},
											{Label: "What is in it", Rank: w.Rank3},
											{Label: "Objects", Width: w.ColNum, Rank: w.Rank4},
											{Label: "Changed on the server", Width: w.ColChanged},
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
								Variant: w.CardChosen,
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
											{Label: "Changed on the server", Width: w.ColChanged},
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
											{Label: "Changed on the server", Width: w.ColChanged},
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
		},
	}
}

// Profile5 is frame «Новая коллекция» (s-profile-5).
func Profile5() Screen {
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
								Meta:  "3",
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
							w.Button{Icon: "plus", Focus: true, Title: "New calendar"},
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
									Mark:   w.Mark{Colour: "#4A6B52"},
									Time:   "10:00",
									Title:  []w.Inline{w.Text("Синк-ап по релизу")},
									Detail: "Переговорная 2",
									Badges: []w.Badge{w.Badge{Label: "4 attendees"}, w.Badge{Label: "repeats weekly"}},
									Source: "nextcloud · Личный",
								},
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
					},
				},
				Height: w.ShellHeight560,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/collections/new",
					Icon:   "plus",
					Title:  "New calendar",
					Body: []w.SheetPart{
						w.Field{
							Label: "On which server",
							Parts: []w.FieldPart{
								w.Pick{
									Name:    "account",
									Options: []w.PickOption{{Value: "nextcloud", Label: "nextcloud", On: true}},
									Label:   "On which server",
								},
								w.FieldHint{
									Parts: []w.Inline{
										w.Text("Two of three are listed: "),
										w.Mono{Text: "baikal"},
										w.Text(" is not answering, and until it does there is nothing to create there."),
									},
									Variant: w.HintSmall,
								},
							},
							Width: w.FieldW100,
						},
						w.Field{
							Label: "Name",
							Parts: []w.FieldPart{w.Input{Name: "name", Value: "Поездки", Label: "Name", Focus: true}},
							Width: w.FieldW100,
						},
						w.Field{
							Label: "Colour",
							Parts: []w.FieldPart{
								w.Swatches{
									Colours: []string{"#4A6B52", "#3F6E8C", "#6B4E7A", "#A05A2C", "#8C6239", "#9A9280"},
									On:      2,
									Name:    "colour",
								},
								w.FieldHint{
									Parts: []w.Inline{
										w.Text("Written to the server as "),
										w.Mono{Text: "calendar-color"},
										w.Text(" — DAVx5 and Thunderbird will show the same one."),
									},
									Variant: w.HintSmall,
								},
							},
							Width: w.FieldW100,
						},
						w.Field{
							Label: "What goes in it",
							Parts: []w.FieldPart{
								w.Group{
									Items: []w.FieldPart{
										w.Check{On: true, Name: "components", Value: "VEVENT", Parts: []w.CheckPart{w.Text("Events")}},
										w.Check{On: true, Name: "components", Value: "VTODO", Parts: []w.CheckPart{w.Text("Tasks")}},
										w.Check{Name: "components", Value: "VJOURNAL", Parts: []w.CheckPart{w.Text("Notes")}},
									},
								},
								w.FieldHint{
									Parts: []w.Inline{
										w.Text("Most servers settle this when the collection is made and will not change it afterwards. A collection holding only tasks appears in "),
										w.Strong("Tasks"),
										w.Text(", not here — the rail follows what is inside, not where it was made."),
									},
									Variant: w.HintSmall,
								},
							},
							Width: w.FieldW100,
						},
						w.Field{
							Label: "Address",
							Parts: []w.FieldPart{
								w.Input{Name: "address", Placeholder: "poezdki", Label: "Address", Mono: true, From: "name"},
								w.FieldHint{
									Parts: []w.Inline{
										w.Mono{Text: "…/remote.php/dav/calendars/mix/"},
										w.SlugOf{Of: "address", Text: "poezdki"},
										w.Mono{Text: "/"},
										w.Text(" — taken from the name, yours to change. Other clients and restore-by-path go by the address; the name they only show."),
									},
									Variant: w.HintSmall,
								},
							},
							Width: w.FieldW100,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Create", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/connections"}, w.Kbd{Keys: "Esc"}},
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

// Profile6 is frame «Новая коллекция · сервер отказал» (s-profile-6).
func Profile6() Screen {
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
								Meta:  "3",
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
							w.Button{Icon: "plus", Focus: true, Title: "New calendar"},
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
									Mark:   w.Mark{Colour: "#4A6B52"},
									Time:   "10:00",
									Title:  []w.Inline{w.Text("Синк-ап по релизу")},
									Detail: "Переговорная 2",
									Badges: []w.Badge{w.Badge{Label: "4 attendees"}, w.Badge{Label: "repeats weekly"}},
									Source: "nextcloud · Личный",
								},
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
					},
				},
				Height: w.ShellHeight340,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/collections/new",
					Fields: []w.Hidden{
						{Name: "account", Value: "radicale"},
						{Name: "name", Value: "Поездки"},
						{Name: "colour", Value: "#6B4E7A"},
						{Name: "components", Value: "VEVENT"},
						{Name: "components", Value: "VTODO"},
						{Name: "address", Value: "poezdki"},
					},
					Icon:  "plus",
					Tone:  w.SheetIconDanger,
					Title: "New calendar",
					Head: []w.SheetHeadPart{
						w.Dim{Parts: []w.Inline{w.Text("radicale · vps")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed},
					},
					Body: []w.SheetPart{
						w.Msg{Parts: []w.MsgPart{w.Text("The server would not make it.")}, Tone: w.MsgAlert, Space: w.MsgFlush},
						w.Diag{
							Text: "[create]  MKCALENDAR /mix/poezdki/ → 403\n              <D:error><D:need-privileges/></D:error>\n              no DAV:bind on the calendar home-set",
						},
						w.SheetNote{
							Text: "This account may sign in and read, but radicale · vps would not let it create here. The server does not say why — no such right for the user, or a password issued read-only; whoever runs the server can tell.",
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Back to the form", Kind: w.ButtonPrimary, To: "/app/collections/new?op=back", Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Button{Label: "Try again", Submit: true},
					},
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Profile7 is frame «Переименовать и сменить цвет» (s-profile-7).
func Profile7() Screen {
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
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("4 servers · 15 collections · 2 subscriptions")}}},
							Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "Connect account", Icon: "plus", Kind: w.ButtonPrimary}}},
							Pad:   w.HeadPadBottom8,
						},
						w.Pad{
							Items: []w.PadPart{
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "nextcloud",
											Items: []w.RubricPart{
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "New collection", Icon: "plus", Small: true},
														w.Button{Label: "Show password", Small: true},
														w.Button{Label: "Edit", Small: true},
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
												{Width: w.ColMenu},
											},
											Rows: []w.TableRow{
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn)},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Личный"}}},
														{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
														{Parts: []w.CellPart{w.Text("128")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn)},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "Дни рождения"}}},
														{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
														{Parts: []w.CellPart{w.Text("51")}, Mono: true},
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
							},
						},
					},
				},
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/collections/rename?a=nextcloud&c=dni-rozhdeniya",
					Icon:   "pencil",
					Title:  "Rename and colour",
					Head:   []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("nextcloud")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}},
					Body: []w.SheetPart{
						w.Field{
							Label: "Name",
							Parts: []w.FieldPart{w.Input{Name: "name", Value: "Дни рождения", Label: "Name", Focus: true}},
							Width: w.FieldW100,
						},
						w.Field{
							Label: "Colour",
							Parts: []w.FieldPart{
								w.Swatches{
									Colours: []string{"#4A6B52", "#3F6E8C", "#6B4E7A", "#A05A2C", "#8C6239", "#9A9280"},
									On:      3,
									Name:    "colour",
								},
							},
							Width: w.FieldW100,
						},
						w.Field{
							Label:      "Address",
							LabelAside: w.Dim{Parts: []w.Inline{w.Text("— stays as it is")}},
							Parts: []w.FieldPart{
								w.Input{
									Value:    "…/remote.php/dav/calendars/mix/birthdays/",
									Label:    "Address",
									Mono:     true,
									ReadOnly: true,
									Variant:  w.InputDusty,
								},
							},
							Width: w.FieldW100,
						},
						w.Msg{
							Parts: []w.MsgPart{w.Text("The address stays: links, devices and backups keep working.")},
							Space: w.MsgFlushTop2,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Save", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/connections"}, w.Kbd{Keys: "Esc"}},
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

// Profile8 is frame «Перенести коллекцию на другой сервер» (s-profile-8).
func Profile8() Screen {
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
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("4 servers · 15 collections · 2 subscriptions")}}},
							Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "Connect account", Icon: "plus", Kind: w.ButtonPrimary}}},
							Pad:   w.HeadPadBottom8,
						},
						w.Pad{
							Items: []w.PadPart{
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "nextcloud",
											Items: []w.RubricPart{
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "New collection", Icon: "plus", Small: true},
														w.Button{Label: "Show password", Small: true},
														w.Button{Label: "Edit", Small: true},
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
												{Width: w.ColMenu},
											},
											Rows: []w.TableRow{
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn)},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Личный"}}},
														{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
														{Parts: []w.CellPart{w.Text("128")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn)},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "Дни рождения"}}},
														{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
														{Parts: []w.CellPart{w.Text("51")}, Mono: true},
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
							},
						},
					},
				},
				Height: w.ShellHeight560,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/collections/move?a=nextcloud&c=lichnyj",
					Icon:   "move",
					Title:  "Move calendar",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("nextcloud · Личный")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
					},
					Body: []w.SheetPart{
						w.Rubric{Text: "Write onto"},
						w.Tree{
							Rows: []w.TreeRow{
								{
									Icon:        "server",
									Label:       "nextcloud",
									Unavailable: true,
									Send:        w.Send{Name: "to", Value: "nextcloud", One: true},
									Trail:       w.Mono{Text: "already there", Variant: w.Mono10Dusty},
								},
								{
									Icon:  "server",
									Label: "radicale · vps",
									Send:  w.Send{Name: "to", Value: "radicale", One: true},
									Trail: w.Mono{Text: "calendars", Variant: w.Mono10Dusty},
								},
								{
									Icon:        "server",
									Label:       "baikal",
									Unavailable: true,
									Send:        w.Send{Name: "to", Value: "baikal", One: true},
									Trail:       w.Mono{Text: "password · 401", Variant: w.Mono10Dusty},
								},
								{
									Icon:        "server",
									Label:       "sftpgo · vps",
									Unavailable: true,
									Send:        w.Send{Name: "to", Value: "sftpgo", One: true},
									Trail:       w.Mono{Text: "files only", Variant: w.Mono10Dusty},
								},
							},
							On: 2,
						},
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("Links, devices and backups lose the old address. nextcloud may send attendees a cancellation."),
							},
							Space: w.MsgFlushTop2,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Move", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/connections"}, w.Kbd{Keys: "Esc"}},
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

// Profile9 is frame «Удаление коллекции» (s-profile-9).
func Profile9() Screen {
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
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("4 servers · 15 collections · 2 subscriptions")}}},
							Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "Connect account", Icon: "plus", Kind: w.ButtonPrimary}}},
							Pad:   w.HeadPadBottom8,
						},
						w.Pad{
							Items: []w.PadPart{
								w.Card{
									Items: []w.CardPart{
										w.Rubric{
											Text: "nextcloud",
											Items: []w.RubricPart{
												w.Line{
													Items: []w.LinePart{
														w.Button{Label: "New collection", Icon: "plus", Small: true},
														w.Button{Label: "Show password", Small: true},
														w.Button{Label: "Edit", Small: true},
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
												{Width: w.ColMenu},
											},
											Rows: []w.TableRow{
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn)},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Личный"}}},
														{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
														{Parts: []w.CellPart{w.Text("128")}, Mono: true},
														{
															Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
															Acts:  true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Pick: ptr(w.TickedOn)},
														{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "Дни рождения"}}},
														{Parts: []w.CellPart{w.Text("calendar")}, Mono: true},
														{Parts: []w.CellPart{w.Text("51")}, Mono: true},
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
							},
						},
					},
				},
				Height: w.ShellHeight560,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/collections/delete?a=nextcloud&c=lichnyj",
					Icon:   "trash",
					Tone:   w.SheetIconDanger,
					Title:  "Delete on the server",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("nextcloud · Личный")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
					},
					Body: []w.SheetPart{
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("Carrel keeps no copy: this cannot be undone. nextcloud may send the attendees of its meetings a cancellation."),
							},
							Tone:  w.MsgAlert,
							Space: w.MsgFlush,
						},
						w.SheetBlock{
							Title: w.Rubric{Text: "What is pointing at it now", Variant: w.RubricBottom7},
							Items: []w.SheetLine{
								{
									Name: "1 published link",
									Note: w.Dim{Parts: []w.Inline{w.Text("— answers “not found” from that moment")}},
								},
								{
									Name: "Backup job «Личное на домашний диск»",
									Note: w.Dim{Parts: []w.Inline{w.Text("— copies this collection; it will run one short and say so")}},
								},
								{
									Name: "3 devices through davloom",
									Note: w.Dim{Parts: []w.Inline{w.Text("— it disappears for them at the next sync")}},
								},
							},
						},
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("Last archived "),
								w.Mono{Text: "2026-08-14"},
								w.Text("; a restore makes a "),
								w.Strong("new"),
								w.Text(" collection at a new address."),
							},
							Space: w.MsgFlush,
						},
						w.Line{
							Items: []w.LinePart{
								w.Button{
									Label: "Download",
									Icon:  "download",
									Href:  "/app/calendar/nextcloud/lichnyj/export",
									Aside: &w.Mono{Text: ".ics", Variant: w.Mono11},
									After: "first",
								},
								w.Dim{
									Parts:   []w.Inline{w.Text("128 events, one file, straight to you")},
									Size:    w.Dim11_5,
									Variant: w.DimWrap,
								},
							},
							Gap:   w.LineGap8,
							Block: true,
						},
						w.Field{
							Label: "Type the name to confirm",
							Parts: []w.FieldPart{w.Input{Name: "name", Placeholder: "Личный", Label: "Type the name to confirm", Focus: true}},
							Width: w.FieldW100,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Delete on the server", Kind: w.ButtonDangerSolid, Submit: true},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/connections"}, w.Kbd{Keys: "Esc"}},
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

// Profile10 is frame «Настройки · подключения · панель подписки открыта» (s-profile-10).
func Profile10() Screen {
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
											{Label: "Objects", Width: w.ColNum, Rank: w.Rank4, Secondary: true},
											{Label: "Changed on the server", Width: w.ColChanged},
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
											{Label: "Objects", Width: w.ColNum, Rank: w.Rank4, Secondary: true},
											{Label: "Changed on the server", Width: w.ColChanged},
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
											{Label: "Objects", Width: w.ColNum, Rank: w.Rank4, Secondary: true},
											{Label: "Changed on the server", Width: w.ColChanged},
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
											{Label: "Objects", Width: w.ColNum, Rank: w.Rank3, Secondary: true},
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
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("A feed is read again on every visit to the section, so there is nothing here to synchronise and nothing to refresh. "),
							w.Strong("Check"),
							w.Text(" asks whether the address still replies."),
						},
					},
				},
			},
			Side: &w.Side{
				Items: []w.SidePart{
					w.Head{
						Title: &w.HeadTitle{Text: "Subscribe to a feed", Size: w.HeadTitle17},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Somebody else’s calendar or address book, read-only")}}},
						Acts:  &w.Acts{Items: []w.ActsPart{w.Close{Label: "Close the panel"}}},
						Pad:   w.HeadPadBottom8,
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Feed address",
								Parts: []w.FieldPart{
									w.Input{
										Name:    "feed.address",
										Value:   "https://feeds.example/ru-holidays.ics",
										Label:   "Feed address",
										Mono:    true,
										Focus:   true,
										Variant: w.Input11_5,
									},
									w.FieldHint{
										Parts: []w.Inline{
											w.Mono{Text: "https://"},
											w.Text(", "),
											w.Mono{Text: "webcal://"},
											w.Text(" and a secret link from another Carrel are all understood."),
										},
									},
								},
								Width: w.FieldW100,
							},
							w.Field{
								Parts:   []w.FieldPart{w.Button{Label: "Check the address", To: "/app/settings/connections/subscribe/check", Submit: true}},
								Width:   w.FieldW100,
								OwnLine: true,
							},
							w.Field{
								Parts: []w.FieldPart{
									w.Diag{
										Text: "[fetch]   GET /ru-holidays.ics → 200 · 0.4 s · 34 KB\n[type]    text/calendar; charset=utf-8\n[parse]   VCALENDAR · 3 VEVENT · 1 VTODO\n[name]    X-WR-CALNAME: Праздники России\n[access]  the server asked for no password",
									},
								},
								Width: w.FieldW100,
							},
							w.Field{
								Label: "Name here",
								Parts: []w.FieldPart{
									w.Input{Name: "name.here", Value: "Праздники РФ", Label: "Name here"},
									w.FieldHint{
										Parts: []w.Inline{
											w.Text("Taken from the feed; change it if you like — other people’s feed names are often longer than the rail."),
										},
									},
								},
								Width: w.FieldW100,
							},
							w.Field{
								Label: "Label colour",
								Parts: []w.FieldPart{
									w.Swatches{
										Colours: []string{"#4A6B52", "#3F6E8C", "#6B4E7A", "#A05A2C", "#8C6239", "#9A9280"},
										On:      4,
										Name:    "colour",
									},
									w.Dim{Parts: []w.Inline{w.Text("derived from the address")}, Size: w.Dim11},
									w.FieldHint{
										Parts: []w.Inline{
											w.Text("A feed sends no colour of its own: Carrel derives one, and hatches the strip so a feed is recognisable among your own collections."),
										},
									},
								},
								Width: w.FieldW100,
							},
							w.Field{
								Parts: []w.FieldPart{
									w.Check{
										Name:  "this.feed.asks.for.a.password",
										Value: "1",
										Parts: []w.CheckPart{w.Text("This feed asks for a password")},
									},
									w.FieldHint{
										Parts: []w.Inline{
											w.Text("The login and password are kept the way DAV credentials are: encrypted under your key. A secret link needs no password — the secret is already in the address."),
										},
									},
								},
								Width:   w.FieldW100,
								OwnLine: true,
							},
							w.Field{
								Label: "Sections",
								Parts: []w.FieldPart{
									w.TagInput{Parts: []w.TagInputPart{w.Tag{Label: "Calendar · 3"}, w.Tag{Label: "Tasks · 1"}}},
									w.FieldHint{
										Parts: []w.Inline{
											w.Text("Not a choice: a feed is one file. What it carries decides where it lands, and that is visible before subscribing."),
										},
									},
								},
								Width: w.FieldW100,
							},
						},
						ID:     "subscribe",
						Action: "/app/settings/connections/subscribe",
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("The feed is fetched by the Carrel server, not by your browser. An address on your local network works only if the server can reach it; a public Carrel cannot, and it says so instead of returning an empty list."),
						},
						Tone:  w.MsgAlert,
						Space: w.MsgInset20,
					},
					w.FormFoot{
						Items: []w.FormFootPart{w.Button{Label: "Subscribe", Kind: w.ButtonPrimary, Form: "subscribe"}, w.Button{Label: "Cancel"}},
					},
				},
			},
		},
	}
}
