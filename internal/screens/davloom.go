// Раздел макета: 7.14 Настройки · davloom, режим прокси.

package screens

import w "carrel.dev/typecase/widget"

// Davloom1 is frame «davloom · точка подключения и устройства» (s-davloom-1).
func Davloom1() Screen {
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
							On: 5,
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
							Text:  "davloom",
							Badge: &w.Badge{Label: "local", Tone: w.BadgeLocal, Title: "Local instance: may be unreachable from outside"},
						},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{
									w.Text("One address for every device instead of an account for every server\u00a0· "),
									w.Span{Face: w.FaceDanger, Text: "a phone on mobile data will not reach a local address"},
								},
							},
						},
						Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Add device", Icon: "plus", Kind: w.ButtonPrimary}}},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Address for the client",
								Parts: []w.FieldPart{
									w.Line{
										Items: []w.LinePart{
											w.Input{
												Value:   "https://carrel.example/dav/",
												Label:   "Address for the client",
												Mono:    true,
												Variant: w.Input11_5Fill,
											},
											w.Button{Label: "Copy", Small: true},
											w.Button{Label: "What the client sees", Small: true, Href: "/app/settings/davloom/tree"},
										},
										Gap: w.LineGapNormal,
									},
									w.FieldHint{
										Parts: []w.Inline{
											w.Text("In DAVx5 choose “Login with URL and user name”: the user name is your Carrel one, the password is the device’s, not yours."),
										},
										Variant: w.HintColumn,
									},
								},
								Width: w.FieldMax640Fill460,
							},
						},
						Variant: w.FormBottom,
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("Changing your Carrel password does "),
							w.Strong("not"),
							w.Text(" disconnect devices; only revoking does."),
						},
						Tone:   w.MsgAccent,
						Column: true,
					},
					w.ListGroup{Title: w.Rubric{Text: "Devices", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
					w.Table{
						Cols: []w.TableCol{
							{Label: "Device"},
							{Label: "What it can see and do", Main: true},
							{Label: "Created", Rank: w.Rank4, Secondary: true},
							{Label: "Last request", Secondary: true},
							{Width: w.Col90},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Icon: "device", Name: "Телефон · DAVx5"}}},
									{
										Parts: []w.CellPart{
											w.CellWrap{
												Items: []w.CellPart{
													w.Line{
														Items: []w.LinePart{
															w.Mark{Colour: "#4A6B52"},
															w.Mark{Colour: "#3F6E8C"},
															w.Mark{Colour: "#6B4E7A"},
															w.Mark{Colour: "#A05A2C"},
														},
														Gap: w.LineGapTight,
													},
													w.Mono{Text: "Calendars, Address books", Variant: w.Mono11},
													w.Button{
														Label: "5 collections · Change",
														Small: true,
														Place: w.Pushed,
														Href:  "/app/settings/davloom/scope?device=phone",
													},
													w.Seg{
														Options: []w.SegOption{
															{Label: "read only", Form: "davloom-acts", To: "/app/settings/davloom?op=read&device=phone"},
															{
																Label: "read and write",
																Form:  "davloom-acts",
																To:    "/app/settings/davloom?op=write&device=phone",
																On:    true,
															},
														},
														Variant: w.SegCompact,
													},
												},
											},
										},
									},
									{Parts: []w.CellPart{w.Text("2026-06-02")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-15 09:31")}, Mono: true},
									{Parts: []w.CellPart{w.Button{Label: "Revoke", Kind: w.ButtonDanger, Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Icon: "device", Name: "Ноутбук · Thunderbird"}}},
									{
										Parts: []w.CellPart{
											w.CellWrap{
												Items: []w.CellPart{
													w.Line{Items: []w.LinePart{w.Mark{Kind: w.MarkMulti}}, Gap: w.LineGapTight},
													w.Mono{Text: "all kinds", Variant: w.Mono11},
													w.Button{
														Label: "12 collections · Change",
														Small: true,
														Place: w.Pushed,
														Href:  "/app/settings/davloom/scope?device=laptop",
													},
													w.Seg{
														Options: []w.SegOption{
															{
																Label: "read only",
																Form:  "davloom-acts",
																To:    "/app/settings/davloom?op=read&device=laptop",
																On:    true,
															},
															{Label: "read and write", Form: "davloom-acts", To: "/app/settings/davloom?op=write&device=laptop"},
														},
														Variant: w.SegCompact,
													},
												},
											},
										},
									},
									{Parts: []w.CellPart{w.Text("2026-06-02")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-15 08:02")}, Mono: true},
									{Parts: []w.CellPart{w.Button{Label: "Revoke", Kind: w.ButtonDanger, Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Icon: "device", Name: "Планшет · Календарь iOS"}}},
									{
										Parts: []w.CellPart{
											w.CellWrap{
												Items: []w.CellPart{
													w.Line{
														Items: []w.LinePart{w.Mark{Colour: "#4A6B52"}, w.Mark{Colour: "#3F6E8C"}, w.Mark{Colour: "#8C6239"}},
														Gap:   w.LineGapTight,
													},
													w.Mono{Text: "Calendars", Variant: w.Mono11},
													w.Button{
														Label: "3 collections · Change",
														Small: true,
														Place: w.Pushed,
														Href:  "/app/settings/davloom/scope?device=tablet",
													},
													w.Seg{
														Options: []w.SegOption{
															{
																Label: "read only",
																Form:  "davloom-acts",
																To:    "/app/settings/davloom?op=read&device=tablet",
																On:    true,
															},
															{Label: "read and write", Form: "davloom-acts", To: "/app/settings/davloom?op=write&device=tablet"},
														},
														Variant: w.SegCompact,
													},
												},
											},
										},
									},
									{Parts: []w.CellPart{w.Text("2026-07-19")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-14 21:15")}, Mono: true},
									{Parts: []w.CellPart{w.Button{Label: "Revoke", Kind: w.ButtonDanger, Small: true}}, Acts: true},
								},
							},
						},
						Variant: w.TableFit,
					},
					w.PostForm{
						ID:     "davloom-acts",
						Action: "/app/settings/davloom",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
				},
			},
		},
	}
}

// Davloom2 is frame «davloom · добавление устройства, шаг 1» (s-davloom-2).
func Davloom2() Screen {
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
							On: 5,
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
								Title: &w.HeadTitle{Text: "Add device"},
								Subs: []w.HeadSub{
									w.HeadSub{Parts: []w.HeadSubPart{w.Text("An app password for one device, with its own key wrapper")}},
								},
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Name",
										Parts: []w.FieldPart{
											w.Input{Name: "name", Value: "Телефон · DAVx5", Label: "Name", Focus: true},
											w.FieldHint{
												Parts:   []w.Inline{w.Text("For you: this is how you will know what you are revoking later.")},
												Variant: w.HintColumn,
											},
										},
										Width: w.FieldW100,
									},
									w.Field{
										Label: "What this device may do",
										Parts: []w.FieldPart{
											w.Seg{
												Options: []w.SegOption{
													{Label: "Read only", Value: "read.only"},
													{Label: "Read and write", On: true, Value: "read.and.write"},
												},
												Variant: w.SegStart,
												Name:    "what.this.device.may.do",
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Read only: the device syncs, and whatever it sends back is refused. Changed later in the device’s row, without a new password."),
												},
												Variant: w.HintColumn,
											},
										},
										Width: w.FieldW100,
									},
									w.Field{
										Label:      "What the device will see",
										LabelAside: w.Dim{Parts: []w.Inline{w.Text("— changed later too, without re-issuing the password")}},
										Parts: []w.FieldPart{
											w.Group{
												Items: []w.FieldPart{
													w.Check{On: true, Name: "calendars", Value: "1", Parts: []w.CheckPart{w.Text("Calendars")}},
													w.Check{On: true, Name: "address.books", Value: "1", Parts: []w.CheckPart{w.Text("Address books")}},
													w.Check{
														Name:  "tasks.vtodo.not.every.client.understands.them",
														Value: "1",
														Parts: []w.CheckPart{w.Text("Tasks "), w.Dim{Parts: []w.Inline{w.Text("· VTODO — not every client understands them")}}},
													},
													w.Check{
														Name:  "notes.vjournal.a.different.subset.of.clients",
														Value: "1",
														Parts: []w.CheckPart{w.Text("Notes "), w.Dim{Parts: []w.Inline{w.Text("· VJOURNAL — a different subset of clients")}}},
													},
													w.Check{
														On:    true,
														Name:  "webdav.folders.mount.as.a.network.drive",
														Value: "1",
														Parts: []w.CheckPart{w.Text("WebDAV folders "), w.Dim{Parts: []w.Inline{w.Text("· mount as a network drive")}}},
													},
												},
												Form: w.GroupColumn,
											},
										},
										Width: w.FieldW100,
									},
									w.Field{
										Label: "Collections",
										Parts: []w.FieldPart{
											w.Group{
												Items: []w.FieldPart{
													w.Check{On: true, Name: "nextcloud", Value: "1", Parts: []w.CheckPart{w.Text("nextcloud · Личный")}},
													w.Check{On: true, Name: "nextcloud", Value: "1", Parts: []w.CheckPart{w.Text("nextcloud · Работа")}},
													w.Check{On: true, Name: "nextcloud", Value: "1", Parts: []w.CheckPart{w.Text("nextcloud · Контакты")}},
													w.Check{Name: "radicale.vps", Value: "1", Parts: []w.CheckPart{w.Text("radicale · vps · Проекты")}},
													w.Check{Name: "radicale.vps", Value: "1", Parts: []w.CheckPart{w.Text("radicale · vps · Команда")}},
													w.Check{
														On:    true,
														Name:  "sftpgo.vps.documents",
														Value: "1",
														Parts: []w.CheckPart{w.Text("sftpgo · vps · Documents")},
													},
												},
												Form: w.GroupColumn,
											},
										},
										Width: w.FieldW100,
									},
								},
								ID:      "device",
								Action:  "/app/settings/davloom/devices/new",
								Variant: w.FormFlush,
							},
							w.Msg{
								Parts:  []w.MsgPart{w.Text("Shown once and stored nowhere: lost, it cannot be recovered.")},
								Tone:   w.MsgAlert,
								Space:  w.MsgFlushSides,
								Column: true,
							},
							w.FormFoot{
								Items:   []w.FormFootPart{w.Button{Label: "Create password", Kind: w.ButtonPrimary, Form: "device"}, w.Button{Label: "Cancel"}},
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

// Davloom3 is frame «davloom · добавление устройства, шаг 2» (s-davloom-3).
func Davloom3() Screen {
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
							On: 5,
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
								Title: &w.HeadTitle{Text: "Телефон · DAVx5"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Password created · read and write · 8 collections")}}},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Device password"},
									w.CardRow{
										Kind: w.CardTop2,
										Gap:  w.LineGapNormal,
										Items: []w.CardRowPart{
											w.Input{Value: "brick-lamp-9f2-carrel", Mono: true, Variant: w.InputKeyFill},
											w.Button{Label: "Copy"},
										},
									},
									w.CardText{
										Kind: w.CardWarn,
										Parts: []w.CardTextPart{
											w.Text("Shown once and stored nowhere. Close this page and you will not see it again; lose it and you revoke this device and make another."),
										},
									},
								},
								Variant: w.CardChosen,
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Address",
										Parts: []w.FieldPart{
											w.Input{
												Name:    "address",
												Value:   "https://carrel.example/dav/",
												Label:   "Address",
												Mono:    true,
												Variant: w.Input11,
											},
										},
										Width: w.FieldW100,
									},
									w.Field{
										Label: "User name",
										Parts: []w.FieldPart{
											w.Input{Name: "user.name", Value: "mix", Label: "User name", Mono: true},
											w.FieldHint{
												Parts:   []w.Inline{w.Text("Your Carrel login; the password is the device’s, not yours.")},
												Variant: w.HintColumn,
											},
										},
										Width: w.FieldW100,
									},
								},
								ID:      "device",
								Action:  "/app/settings/davloom/devices/new",
								Variant: w.FormFlush,
							},
							w.FormFoot{
								Items:   []w.FormFootPart{w.Button{Label: "Done", Kind: w.ButtonPrimary, Form: "device"}},
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

// Davloom4 is frame «Что видно по адресу прокси в браузере» (s-davloom-4).
func Davloom4() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Hero{
						Symbol: "davloom-tile",
						Label:  "davloom",
						Title:  "davloom",
						Sub:    "DAV proxy for Carrel · one address instead of every configured account",
						Text: w.Prose{
							Parts: []w.Inline{
								w.Text("This is a connection point for clients, not a page to read. Enter the address in anything that speaks CalDAV, CardDAV or WebDAV — DAVx5, Thunderbird, Calendar and Contacts on Apple, a file manager mounting a network drive. The user name is your Carrel one; the password is the device password issued in Carrel."),
							},
							Variant: w.ProseTop,
						},
						Fields: &w.Fields{
							Items: []w.FieldsItem{
								{Term: "Address", Value: []w.ValuePart{w.Text("https://carrel.example/dav/")}, Face: w.ValueMono},
								{Term: "Serves", Value: []w.ValuePart{w.Text("CalDAV · CardDAV · WebDAV")}, Face: w.ValueMono},
								{Term: "Version", Value: []w.ValuePart{w.Text("0.0.1")}, Face: w.ValueMono},
								{Term: "Build", Value: []w.ValuePart{w.Text("786949a")}, Face: w.ValueMono},
								{Term: "Licence", Value: []w.ValuePart{w.Text("AGPL-3.0-or-later")}},
							},
							Variant: w.Fields130Flush,
						},
						Acts: []w.Button{w.Button{Label: "Source code", Kind: w.ButtonPrimary}, w.Button{Label: "Open Carrel"}},
					},
				},
			},
		},
	}
}

// Davloom6 is frame «davloom · устройств нет» (s-davloom-6).
func Davloom6() Screen {
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
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Href: "#"},
								{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 5,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Text:   "davloom · running · no devices",
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
							Text:  "davloom",
							Badge: &w.Badge{Label: "local", Tone: w.BadgeLocal, Title: "Local instance: may be unreachable from outside"},
						},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{
									w.Text("One address for every device instead of an account for every server\u00a0· "),
									w.Span{Face: w.FaceDanger, Text: "a phone on mobile data will not reach a local address"},
								},
							},
						},
						Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Add device", Icon: "plus", Kind: w.ButtonPrimary}}},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Address for the client",
								Parts: []w.FieldPart{
									w.Line{
										Items: []w.LinePart{
											w.Input{
												Value:   "https://carrel.example/dav/",
												Label:   "Address for the client",
												Mono:    true,
												Variant: w.Input11_5Fill,
											},
											w.Button{Label: "Copy", Small: true},
											w.Button{Label: "What the client sees", Small: true, Href: "/app/settings/davloom/tree"},
										},
										Gap: w.LineGapNormal,
									},
									w.FieldHint{
										Parts: []w.Inline{
											w.Text("In DAVx5 choose “Login with URL and user name”: the user name is your Carrel one, the password is the device’s, not yours."),
										},
										Variant: w.HintColumn,
									},
								},
								Width: w.FieldMax640Fill460,
							},
						},
						Variant: w.FormBottom,
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("Changing your Carrel password does "),
							w.Strong("not"),
							w.Text(" disconnect devices; only revoking does."),
						},
						Tone:   w.MsgAccent,
						Column: true,
					},
					w.ListGroup{Title: w.Rubric{Text: "Devices", Form: w.RubricInline}},
					w.Empty{Parts: []w.EmptyPart{w.Strong("No devices yet.")}},
				},
			},
		},
	}
}

// Davloom7 is frame «davloom · что видит устройство» (s-davloom-7).
func Davloom7() Screen {
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
								On: 5,
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
								Text:  "davloom",
								Badge: &w.Badge{Label: "local", Tone: w.BadgeLocal, Title: "Local instance: may be unreachable from outside"},
							},
							Subs: []w.HeadSub{
								w.HeadSub{
									Parts: []w.HeadSubPart{
										w.Text("One address for every device instead of an account for every server\u00a0· "),
										w.Span{Face: w.FaceDanger, Text: "a phone on mobile data will not reach a local address"},
									},
								},
							},
							Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Add device", Icon: "plus", Kind: w.ButtonPrimary}}},
						},
						w.Form{
							Items: []w.FormPart{
								w.Field{
									Label: "Address for the client",
									Parts: []w.FieldPart{
										w.Line{
											Items: []w.LinePart{
												w.Input{
													Value:   "https://carrel.example/dav/",
													Label:   "Address for the client",
													Mono:    true,
													Variant: w.Input11_5Fill,
												},
												w.Button{Label: "Copy", Small: true},
												w.Button{Label: "What the client sees", Small: true, Href: "/app/settings/davloom/tree"},
											},
											Gap: w.LineGapNormal,
										},
										w.FieldHint{
											Parts: []w.Inline{
												w.Text("In DAVx5 choose “Login with URL and user name”: the user name is your Carrel one, the password is the device’s, not yours."),
											},
											Variant: w.HintColumn,
										},
									},
									Width: w.FieldMax640Fill460,
								},
							},
							Variant: w.FormBottom,
						},
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("Changing your Carrel password does "),
								w.Strong("not"),
								w.Text(" disconnect devices; only revoking does."),
							},
							Tone:   w.MsgAccent,
							Column: true,
						},
						w.ListGroup{Title: w.Rubric{Text: "Devices", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
						w.Table{
							Cols: []w.TableCol{
								{Label: "Device"},
								{Label: "What it can see and do", Main: true},
								{Label: "Created", Rank: w.Rank4, Secondary: true},
								{Label: "Last request", Secondary: true},
								{Width: w.Col90},
							},
							Rows: []w.TableRow{
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.TableFile{Icon: "device", Name: "Телефон · DAVx5"}}},
										{
											Parts: []w.CellPart{
												w.CellWrap{
													Items: []w.CellPart{
														w.Line{
															Items: []w.LinePart{
																w.Mark{Colour: "#4A6B52"},
																w.Mark{Colour: "#3F6E8C"},
																w.Mark{Colour: "#6B4E7A"},
																w.Mark{Colour: "#A05A2C"},
															},
															Gap: w.LineGapTight,
														},
														w.Mono{Text: "Calendars, Address books", Variant: w.Mono11},
														w.Button{
															Label: "5 collections · Change",
															Small: true,
															Place: w.Pushed,
															Href:  "/app/settings/davloom/scope?device=phone",
														},
														w.Seg{
															Options: []w.SegOption{
																{Label: "read only", Form: "davloom-acts", To: "/app/settings/davloom?op=read&device=phone"},
																{
																	Label: "read and write",
																	Form:  "davloom-acts",
																	To:    "/app/settings/davloom?op=write&device=phone",
																	On:    true,
																},
															},
															Variant: w.SegCompact,
														},
													},
												},
											},
										},
										{Parts: []w.CellPart{w.Text("2026-06-02")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-15 09:31")}, Mono: true},
										{Parts: []w.CellPart{w.Button{Label: "Revoke", Kind: w.ButtonDanger, Small: true}}, Acts: true},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.TableFile{Icon: "device", Name: "Ноутбук · Thunderbird"}}},
										{
											Parts: []w.CellPart{
												w.CellWrap{
													Items: []w.CellPart{
														w.Line{Items: []w.LinePart{w.Mark{Kind: w.MarkMulti}}, Gap: w.LineGapTight},
														w.Mono{Text: "all kinds", Variant: w.Mono11},
														w.Button{
															Label: "12 collections · Change",
															Small: true,
															Place: w.Pushed,
															Href:  "/app/settings/davloom/scope?device=laptop",
														},
														w.Seg{
															Options: []w.SegOption{
																{
																	Label: "read only",
																	Form:  "davloom-acts",
																	To:    "/app/settings/davloom?op=read&device=laptop",
																	On:    true,
																},
																{Label: "read and write", Form: "davloom-acts", To: "/app/settings/davloom?op=write&device=laptop"},
															},
															Variant: w.SegCompact,
														},
													},
												},
											},
										},
										{Parts: []w.CellPart{w.Text("2026-06-02")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-15 08:02")}, Mono: true},
										{Parts: []w.CellPart{w.Button{Label: "Revoke", Kind: w.ButtonDanger, Small: true}}, Acts: true},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.TableFile{Icon: "device", Name: "Планшет · Календарь iOS"}}},
										{
											Parts: []w.CellPart{
												w.CellWrap{
													Items: []w.CellPart{
														w.Line{
															Items: []w.LinePart{w.Mark{Colour: "#4A6B52"}, w.Mark{Colour: "#3F6E8C"}, w.Mark{Colour: "#8C6239"}},
															Gap:   w.LineGapTight,
														},
														w.Mono{Text: "Calendars", Variant: w.Mono11},
														w.Button{
															Label: "3 collections · Change",
															Small: true,
															Place: w.Pushed,
															Href:  "/app/settings/davloom/scope?device=tablet",
														},
														w.Seg{
															Options: []w.SegOption{
																{
																	Label: "read only",
																	Form:  "davloom-acts",
																	To:    "/app/settings/davloom?op=read&device=tablet",
																	On:    true,
																},
																{Label: "read and write", Form: "davloom-acts", To: "/app/settings/davloom?op=write&device=tablet"},
															},
															Variant: w.SegCompact,
														},
													},
												},
											},
										},
										{Parts: []w.CellPart{w.Text("2026-07-19")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-14 21:15")}, Mono: true},
										{Parts: []w.CellPart{w.Button{Label: "Revoke", Kind: w.ButtonDanger, Small: true}}, Acts: true},
									},
								},
							},
							Variant: w.TableFit,
						},
						w.PostForm{
							ID:     "davloom-acts",
							Action: "/app/settings/davloom",
							Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
						},
					},
				},
				Height: w.ShellHeight560,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/settings/davloom/scope?device=phone",
					Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					Title:  "What it sees",
					Head: []w.SheetHeadPart{
						w.Dim{Parts: []w.Inline{w.Text("Телефон · DAVx5")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed},
					},
					Body: []w.SheetPart{
						w.Form{
							Items: []w.FormPart{
								w.FieldSet{
									Items: []w.FieldSetPart{
										w.FieldGrid{
											Items: []w.FieldGridPart{
												w.FieldGridColumn{
													Items: []w.TickRow{
														w.TickRow{
															Ticked: w.TickedPart,
															Form:   w.TickRowGroup,
															Mark:   w.Mark{Kind: w.MarkMulti},
															Label:  "nextcloud",
															Meta:   "4 of 9",
														},
														w.TickRow{
															Send:  w.Send{Name: "col", Value: "nextcloud|lichnyj"},
															Form:  w.TickRowNested,
															Mark:  w.Mark{Colour: "#4A6B52"},
															Label: "Личный",
															Meta:  "cal.",
														},
														w.TickRow{
															Ticked: w.TickedOff,
															Send:   w.Send{Name: "col", Value: "nextcloud|dela"},
															Form:   w.TickRowNested,
															Mark:   w.Mark{Colour: "#4A6B52"},
															Label:  "Дела",
															Meta:   "tasks",
														},
														w.TickRow{
															Send:  w.Send{Name: "col", Value: "nextcloud|rabota"},
															Form:  w.TickRowNested,
															Mark:  w.Mark{Colour: "#3F6E8C"},
															Label: "Работа",
															Meta:  "cal.",
														},
														w.TickRow{
															Ticked: w.TickedOff,
															Send:   w.Send{Name: "col", Value: "nextcloud|dni-rozhdeniya"},
															Form:   w.TickRowNested,
															Mark:   w.Mark{Colour: "#A05A2C"},
															Label:  "Дни рождения",
															Meta:   "cal.",
														},
														w.TickRow{
															Ticked: w.TickedOff,
															Send:   w.Send{Name: "col", Value: "nextcloud|zapisi"},
															Form:   w.TickRowNested,
															Mark:   w.Mark{Colour: "#4A6B52"},
															Label:  "Записи",
															Meta:   "notes",
														},
														w.TickRow{
															Ticked: w.TickedOff,
															Send:   w.Send{Name: "col", Value: "nextcloud|dnevnik"},
															Form:   w.TickRowNested,
															Mark:   w.Mark{Colour: "#6B4E7A"},
															Label:  "Дневник",
															Meta:   "notes",
														},
														w.TickRow{
															Ticked: w.TickedOff,
															Send:   w.Send{Name: "col", Value: "nextcloud|zapiski-obshchie"},
															Form:   w.TickRowNested,
															Mark:   w.Mark{Colour: "#3F6E8C"},
															Label:  "Общие с семьёй",
															Meta:   "notes",
														},
														w.TickRow{
															Send:  w.Send{Name: "col", Value: "nextcloud|kontakty"},
															Form:  w.TickRowNested,
															Mark:  w.Mark{Colour: "#4A6B52"},
															Label: "Контакты",
															Meta:  "book",
														},
														w.TickRow{
															Send:  w.Send{Name: "col", Value: "nextcloud|arhiv"},
															Form:  w.TickRowNested,
															Mark:  w.Mark{Colour: "#A05A2C"},
															Label: "Архив",
															Meta:  "book",
														},
													},
												},
												w.FieldGridColumn{
													Items: []w.TickRow{
														w.TickRow{
															Ticked: w.TickedPart,
															Form:   w.TickRowGroup,
															Mark:   w.Mark{Kind: w.MarkMulti},
															Label:  "radicale · vps",
															Meta:   "1 of 4",
														},
														w.TickRow{
															Send:  w.Send{Name: "col", Value: "radicale|proekty"},
															Form:  w.TickRowNested,
															Mark:  w.Mark{Colour: "#6B4E7A"},
															Label: "Проекты",
															Meta:  "cal.",
														},
														w.TickRow{
															Ticked: w.TickedOff,
															Send:   w.Send{Name: "col", Value: "radicale|dezhurstva"},
															Form:   w.TickRowNested,
															Mark:   w.Mark{Colour: "#8C6239"},
															Label:  "Дежурства",
															Meta:   "cal.",
														},
														w.TickRow{
															Ticked: w.TickedOff,
															Send:   w.Send{Name: "col", Value: "radicale|komanda"},
															Form:   w.TickRowNested,
															Mark:   w.Mark{Colour: "#6B4E7A"},
															Label:  "Команда",
															Meta:   "book",
														},
														w.TickRow{
															Ticked: w.TickedOff,
															Send:   w.Send{Name: "col", Value: "radicale|rassylki"},
															Form:   w.TickRowNested,
															Mark:   w.Mark{Colour: "#6B4E7A"},
															Label:  "Рассылки",
															Meta:   "book",
														},
													},
												},
											},
										},
									},
									Variant: w.FieldSetLead,
								},
							},
							Variant: w.FormFlush,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Save", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/davloom"}, w.Kbd{Keys: "Esc"}},
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

// Davloom8 is frame «davloom · выключен» (s-davloom-8).
func Davloom8() Screen {
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
							On: 5,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Tone:   w.RailStatusOff,
							Text:   "davloom is off",
							Detail: w.Dim{Parts: []w.Inline{w.Text("No device connects until it is started.")}, Size: w.Dim11},
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
						Title: &w.HeadTitle{Text: "davloom", Badge: &w.Badge{Label: "off", Tone: w.BadgeAlert}},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("No device connects until it is started.")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{
									Label: "Add device",
									Icon:  "plus",
									Kind:  w.ButtonPrimary,
									Title: "davloom is off on this instance.",
									State: w.ButtonDisabled,
								},
							},
						},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Address for the client",
								Parts: []w.FieldPart{
									w.Line{
										Items: []w.LinePart{
											w.Input{
												Value:   "https://carrel.example/dav/",
												Label:   "Address for the client",
												Mono:    true,
												Variant: w.Input11_5Fill,
											},
											w.Button{Label: "Copy", Small: true},
											w.Button{Label: "What the client sees", Small: true, Href: "/app/settings/davloom/tree"},
										},
										Gap: w.LineGapNormal,
									},
									w.FieldHint{
										Parts: []w.Inline{
											w.Text("In DAVx5 choose “Login with URL and user name”: the user name is your Carrel one, the password is the device’s, not yours."),
										},
										Variant: w.HintColumn,
									},
								},
								Width: w.FieldMax640Fill460,
							},
						},
						Variant: w.FormBottom,
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("Changing your Carrel password does "),
							w.Strong("not"),
							w.Text(" disconnect devices; only revoking does."),
						},
						Tone:   w.MsgAccent,
						Column: true,
					},
					w.ListGroup{Title: w.Rubric{Text: "Devices", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
					w.Table{
						Cols: []w.TableCol{
							{Label: "Device"},
							{Label: "What it can see and do", Main: true},
							{Label: "Created", Rank: w.Rank4, Secondary: true},
							{Label: "Last request", Secondary: true},
							{Width: w.Col90},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Icon: "device", Name: "Телефон · DAVx5"}}},
									{
										Parts: []w.CellPart{
											w.CellWrap{
												Items: []w.CellPart{
													w.Line{
														Items: []w.LinePart{
															w.Mark{Colour: "#4A6B52"},
															w.Mark{Colour: "#3F6E8C"},
															w.Mark{Colour: "#6B4E7A"},
															w.Mark{Colour: "#A05A2C"},
														},
														Gap: w.LineGapTight,
													},
													w.Mono{Text: "Calendars, Address books", Variant: w.Mono11},
													w.Button{
														Label: "5 collections · Change",
														Small: true,
														Place: w.Pushed,
														Href:  "/app/settings/davloom/scope?device=phone",
													},
													w.Seg{
														Options: []w.SegOption{
															{Label: "read only", Form: "davloom-acts", To: "/app/settings/davloom?op=read&device=phone"},
															{
																Label: "read and write",
																Form:  "davloom-acts",
																To:    "/app/settings/davloom?op=write&device=phone",
																On:    true,
															},
														},
														Variant: w.SegCompact,
													},
												},
											},
										},
									},
									{Parts: []w.CellPart{w.Text("2026-06-02")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-15 09:31")}, Mono: true},
									{Parts: []w.CellPart{w.Button{Label: "Revoke", Kind: w.ButtonDanger, Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Icon: "device", Name: "Ноутбук · Thunderbird"}}},
									{
										Parts: []w.CellPart{
											w.CellWrap{
												Items: []w.CellPart{
													w.Line{Items: []w.LinePart{w.Mark{Kind: w.MarkMulti}}, Gap: w.LineGapTight},
													w.Mono{Text: "all kinds", Variant: w.Mono11},
													w.Button{
														Label: "12 collections · Change",
														Small: true,
														Place: w.Pushed,
														Href:  "/app/settings/davloom/scope?device=laptop",
													},
													w.Seg{
														Options: []w.SegOption{
															{
																Label: "read only",
																Form:  "davloom-acts",
																To:    "/app/settings/davloom?op=read&device=laptop",
																On:    true,
															},
															{Label: "read and write", Form: "davloom-acts", To: "/app/settings/davloom?op=write&device=laptop"},
														},
														Variant: w.SegCompact,
													},
												},
											},
										},
									},
									{Parts: []w.CellPart{w.Text("2026-06-02")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-15 08:02")}, Mono: true},
									{Parts: []w.CellPart{w.Button{Label: "Revoke", Kind: w.ButtonDanger, Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Icon: "device", Name: "Планшет · Календарь iOS"}}},
									{
										Parts: []w.CellPart{
											w.CellWrap{
												Items: []w.CellPart{
													w.Line{
														Items: []w.LinePart{w.Mark{Colour: "#4A6B52"}, w.Mark{Colour: "#3F6E8C"}, w.Mark{Colour: "#8C6239"}},
														Gap:   w.LineGapTight,
													},
													w.Mono{Text: "Calendars", Variant: w.Mono11},
													w.Button{
														Label: "3 collections · Change",
														Small: true,
														Place: w.Pushed,
														Href:  "/app/settings/davloom/scope?device=tablet",
													},
													w.Seg{
														Options: []w.SegOption{
															{
																Label: "read only",
																Form:  "davloom-acts",
																To:    "/app/settings/davloom?op=read&device=tablet",
																On:    true,
															},
															{Label: "read and write", Form: "davloom-acts", To: "/app/settings/davloom?op=write&device=tablet"},
														},
														Variant: w.SegCompact,
													},
												},
											},
										},
									},
									{Parts: []w.CellPart{w.Text("2026-07-19")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-14 21:15")}, Mono: true},
									{Parts: []w.CellPart{w.Button{Label: "Revoke", Kind: w.ButtonDanger, Small: true}}, Acts: true},
								},
							},
						},
						Variant: w.TableFit,
					},
					w.PostForm{
						ID:     "davloom-acts",
						Action: "/app/settings/davloom",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
				},
			},
		},
	}
}

// Davloom9 is frame «davloom · что видит клиент» (s-davloom-9).
func Davloom9() Screen {
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
								On: 5,
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
								Text:  "davloom",
								Badge: &w.Badge{Label: "local", Tone: w.BadgeLocal, Title: "Local instance: may be unreachable from outside"},
							},
							Subs: []w.HeadSub{
								w.HeadSub{
									Parts: []w.HeadSubPart{
										w.Text("One address for every device instead of an account for every server\u00a0· "),
										w.Span{Face: w.FaceDanger, Text: "a phone on mobile data will not reach a local address"},
									},
								},
							},
							Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Add device", Icon: "plus", Kind: w.ButtonPrimary}}},
						},
						w.Form{
							Items: []w.FormPart{
								w.Field{
									Label: "Address for the client",
									Parts: []w.FieldPart{
										w.Line{
											Items: []w.LinePart{
												w.Input{
													Value:   "https://carrel.example/dav/",
													Label:   "Address for the client",
													Mono:    true,
													Variant: w.Input11_5Fill,
												},
												w.Button{Label: "Copy", Small: true},
												w.Button{Label: "What the client sees", Small: true, Href: "/app/settings/davloom/tree"},
											},
											Gap: w.LineGapNormal,
										},
										w.FieldHint{
											Parts: []w.Inline{
												w.Text("In DAVx5 choose “Login with URL and user name”: the user name is your Carrel one, the password is the device’s, not yours."),
											},
											Variant: w.HintColumn,
										},
									},
									Width: w.FieldMax640Fill460,
								},
							},
							Variant: w.FormBottom,
						},
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("Changing your Carrel password does "),
								w.Strong("not"),
								w.Text(" disconnect devices; only revoking does."),
							},
							Tone:   w.MsgAccent,
							Column: true,
						},
						w.ListGroup{Title: w.Rubric{Text: "Devices", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
						w.Table{
							Cols: []w.TableCol{
								{Label: "Device"},
								{Label: "What it can see and do", Main: true},
								{Label: "Created", Rank: w.Rank4, Secondary: true},
								{Label: "Last request", Secondary: true},
								{Width: w.Col90},
							},
							Rows: []w.TableRow{
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.TableFile{Icon: "device", Name: "Телефон · DAVx5"}}},
										{
											Parts: []w.CellPart{
												w.CellWrap{
													Items: []w.CellPart{
														w.Line{
															Items: []w.LinePart{
																w.Mark{Colour: "#4A6B52"},
																w.Mark{Colour: "#3F6E8C"},
																w.Mark{Colour: "#6B4E7A"},
																w.Mark{Colour: "#A05A2C"},
															},
															Gap: w.LineGapTight,
														},
														w.Mono{Text: "Calendars, Address books", Variant: w.Mono11},
														w.Button{
															Label: "5 collections · Change",
															Small: true,
															Place: w.Pushed,
															Href:  "/app/settings/davloom/scope?device=phone",
														},
														w.Seg{
															Options: []w.SegOption{
																{Label: "read only", Form: "davloom-acts", To: "/app/settings/davloom?op=read&device=phone"},
																{
																	Label: "read and write",
																	Form:  "davloom-acts",
																	To:    "/app/settings/davloom?op=write&device=phone",
																	On:    true,
																},
															},
															Variant: w.SegCompact,
														},
													},
												},
											},
										},
										{Parts: []w.CellPart{w.Text("2026-06-02")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-15 09:31")}, Mono: true},
										{Parts: []w.CellPart{w.Button{Label: "Revoke", Kind: w.ButtonDanger, Small: true}}, Acts: true},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.TableFile{Icon: "device", Name: "Ноутбук · Thunderbird"}}},
										{
											Parts: []w.CellPart{
												w.CellWrap{
													Items: []w.CellPart{
														w.Line{Items: []w.LinePart{w.Mark{Kind: w.MarkMulti}}, Gap: w.LineGapTight},
														w.Mono{Text: "all kinds", Variant: w.Mono11},
														w.Button{
															Label: "12 collections · Change",
															Small: true,
															Place: w.Pushed,
															Href:  "/app/settings/davloom/scope?device=laptop",
														},
														w.Seg{
															Options: []w.SegOption{
																{
																	Label: "read only",
																	Form:  "davloom-acts",
																	To:    "/app/settings/davloom?op=read&device=laptop",
																	On:    true,
																},
																{Label: "read and write", Form: "davloom-acts", To: "/app/settings/davloom?op=write&device=laptop"},
															},
															Variant: w.SegCompact,
														},
													},
												},
											},
										},
										{Parts: []w.CellPart{w.Text("2026-06-02")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-15 08:02")}, Mono: true},
										{Parts: []w.CellPart{w.Button{Label: "Revoke", Kind: w.ButtonDanger, Small: true}}, Acts: true},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.TableFile{Icon: "device", Name: "Планшет · Календарь iOS"}}},
										{
											Parts: []w.CellPart{
												w.CellWrap{
													Items: []w.CellPart{
														w.Line{
															Items: []w.LinePart{w.Mark{Colour: "#4A6B52"}, w.Mark{Colour: "#3F6E8C"}, w.Mark{Colour: "#8C6239"}},
															Gap:   w.LineGapTight,
														},
														w.Mono{Text: "Calendars", Variant: w.Mono11},
														w.Button{
															Label: "3 collections · Change",
															Small: true,
															Place: w.Pushed,
															Href:  "/app/settings/davloom/scope?device=tablet",
														},
														w.Seg{
															Options: []w.SegOption{
																{
																	Label: "read only",
																	Form:  "davloom-acts",
																	To:    "/app/settings/davloom?op=read&device=tablet",
																	On:    true,
																},
																{Label: "read and write", Form: "davloom-acts", To: "/app/settings/davloom?op=write&device=tablet"},
															},
															Variant: w.SegCompact,
														},
													},
												},
											},
										},
										{Parts: []w.CellPart{w.Text("2026-07-19")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-14 21:15")}, Mono: true},
										{Parts: []w.CellPart{w.Button{Label: "Revoke", Kind: w.ButtonDanger, Small: true}}, Acts: true},
									},
								},
							},
							Variant: w.TableFit,
						},
						w.PostForm{
							ID:     "davloom-acts",
							Action: "/app/settings/davloom",
							Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
						},
					},
				},
				Height: w.ShellHeight560,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Title: "What the client sees",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("https://carrel.example/dav/")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
					},
					Body: []w.SheetPart{
						w.Tree{
							Rows: []w.TreeRow{
								{
									Icon:  "file",
									Label: "/dav/",
									Code:  true,
									Trail: w.Dim{Parts: []w.Inline{w.Text("root")}, Size: w.Dim11},
								},
								{Icon: "con", Label: "principal/", Code: true, Indent: w.TreeLevel22},
								{
									Icon:   "cal",
									Label:  "cal/",
									Code:   true,
									Indent: w.TreeLevel22,
									Trail:  w.Dim{Parts: []w.Inline{w.Text("Calendars · events, tasks, notes")}, Size: w.Dim11},
								},
								{
									Mark:   &w.Mark{Colour: "#4A6B52"},
									Label:  "c/k4f2/%2F…%2Fpersonal%2F/",
									Code:   true,
									Indent: w.TreeLevel44,
									Trail:  w.Dim{Parts: []w.Inline{w.Text("Личный")}, Size: w.Dim11},
								},
								{
									Mark:   &w.Mark{Colour: "#3F6E8C"},
									Label:  "c/k4f2/%2F…%2Fwork%2F/",
									Code:   true,
									Indent: w.TreeLevel44,
									Trail:  w.Dim{Parts: []w.Inline{w.Text("Работа")}, Size: w.Dim11},
								},
								{
									Mark:   &w.Mark{Colour: "#6B4E7A"},
									Label:  "c/r9a1/%2Fmix%2Fprojects%2F/",
									Code:   true,
									Indent: w.TreeLevel44,
									Trail:  w.Dim{Parts: []w.Inline{w.Text("Проекты")}, Size: w.Dim11},
								},
								{
									Icon:   "con",
									Label:  "card/",
									Code:   true,
									Indent: w.TreeLevel22,
									Trail:  w.Dim{Parts: []w.Inline{w.Text("Address books")}, Size: w.Dim11},
								},
								{
									Mark:   &w.Mark{Colour: "#4A6B52"},
									Label:  "c/k4f2/%2F…%2Fcontacts%2F/",
									Code:   true,
									Indent: w.TreeLevel44,
									Trail:  w.Dim{Parts: []w.Inline{w.Text("Контакты")}, Size: w.Dim11},
								},
								{
									Mark:   &w.Mark{Colour: "#6B4E7A"},
									Label:  "c/r9a1/%2Fmix%2Fteam%2F/",
									Code:   true,
									Indent: w.TreeLevel44,
									Trail:  w.Dim{Parts: []w.Inline{w.Text("Команда")}, Size: w.Dim11},
								},
								{
									Icon:   "file",
									Label:  "files/",
									Code:   true,
									Indent: w.TreeLevel22,
									Trail:  w.Dim{Parts: []w.Inline{w.Text("Files · plain WebDAV")}, Size: w.Dim11},
								},
								{
									Mark:   &w.Mark{Colour: "#4A6B52"},
									Label:  "c/s3c8/%2FDocuments%2F/",
									Code:   true,
									Indent: w.TreeLevel44,
									Trail:  w.Dim{Parts: []w.Inline{w.Text("Documents")}, Size: w.Dim11},
								},
							},
							Variant: w.TreeNarrow,
						},
					},
					Foot: []w.SheetFootPart{
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Close", Href: "/app/settings/davloom"}, w.Kbd{Keys: "Esc"}},
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
