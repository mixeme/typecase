// Раздел макета: 7.13 Настройки · новое устройство.

package screens

import w "carrel.dev/typecase/widget"

// Plan1 is frame «Новое устройство · профиль Apple» (s-plan-1).
func Plan1() Screen {
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
							On: 4,
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
						Title: &w.HeadTitle{Text: "Set up a new device"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{
									w.Text("Everything discovery found — as one file or one code "),
									w.Badge{Label: "credentials unlocked · 4:12", Accent: true},
									w.Button{Label: "Lock now", Small: true},
								},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Seg{
									Options: []w.SegOption{{Label: "Apple", On: true}, {Label: "Android"}, {Label: "Thunderbird"}, {Label: "Anything else"}},
								},
							},
						},
					},
					w.Form{
						Items: []w.FormPart{
							w.FormColumn{
								Wide: true,
								Items: []w.FormColumnPart{
									w.Card{
										Items: []w.CardPart{
											w.Rubric{Text: "What goes into the profile"},
											w.List{
												Rows: []w.ListRow{
													w.SlimRow{
														Mark:   w.Mark{Kind: w.MarkMulti},
														Title:  "nextcloud · mix",
														Inside: "· Личный, Работа, Контакты",
														Aside:  "CalDAV, CardDAV",
													},
													w.SlimRow{
														Ticked: w.TickedOff,
														Mark:   w.Mark{Kind: w.MarkMulti},
														Title:  "radicale · vps · mix",
														Inside: "· Проекты",
														Aside:  "CalDAV",
													},
												},
												Inset: true,
											},
										},
									},
									w.Msg{
										Icon: "shield",
										Parts: []w.MsgPart{
											w.Strong("The profile carries passwords in cleartext."),
											w.Text(" It is an ordinary file: it lands in Downloads, in a messenger, in the phone’s backup. Open it on the device and delete it. The link opens once and lives 5 minutes, for you signed in to Carrel; nothing is written to the server’s disk, and the issue is recorded in the audit log."),
										},
										Tone:  w.MsgAlert,
										Space: w.MsgFlush,
									},
								},
							},
							w.FormColumn{
								Items: []w.FormColumnPart{
									w.Card{
										Items: []w.CardPart{
											w.Rubric{Text: "Profile"},
											w.CardHint{Parts: []w.CardTextPart{w.Text("Issued inside the unlock window: nothing more is asked.")}},
											w.CardRow{Kind: w.CardActs, Items: []w.CardRowPart{w.Button{Label: "Generate profile", Kind: w.ButtonDanger}}},
										},
									},
									w.Card{
										Items: []w.CardPart{
											w.Rubric{
												Text: "Link ready until 09:46",
												Items: []w.RubricPart{
													w.Badge{Label: "local", Tone: w.BadgeLocal, Title: "Local instance: may be unreachable from outside"},
												},
											},
											w.CardHint{
												Parts: []w.CardTextPart{
													w.Text("It opens once, for you signed in to Carrel. Open it on the device, install the profile, then delete the file."),
												},
											},
											w.CardRow{
												Kind:  w.CardTop9,
												Gap:   6,
												Items: []w.CardRowPart{w.Button{Label: "Download profile", Icon: "down"}, w.Button{Label: "Revoke", Kind: w.ButtonDanger}},
											},
										},
									},
								},
							},
						},
						Variant: w.FormFlushBottom,
					},
				},
			},
		},
	}
}

// Plan2 is frame «Android · DAVx5» (s-plan-2).
func Plan2() Screen {
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
							On: 4,
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
						Title: &w.HeadTitle{Text: "Set up a new device"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{
									w.Text("DAVx5 · collections, login and password "),
									w.Badge{Label: "credentials unlocked · 4:12", Accent: true},
									w.Button{Label: "Lock now", Small: true},
								},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Seg{
									Options: []w.SegOption{{Label: "Apple"}, {Label: "Android", On: true}, {Label: "Thunderbird"}, {Label: "Anything else"}},
								},
							},
						},
					},
					w.Form{
						Items: []w.FormPart{
							w.FormColumn{
								Wide: true,
								Items: []w.FormColumnPart{
									w.Card{
										Items: []w.CardPart{
											w.Rubric{Text: "Collections DAVx5 will see"},
											w.List{
												Rows: []w.ListRow{
													w.SlimRow{
														Mark:   w.Mark{Kind: w.MarkMulti},
														Title:  "nextcloud · mix",
														Inside: "· Личный, Работа, Контакты",
														Aside:  "CalDAV, CardDAV",
													},
													w.SlimRow{Mark: w.Mark{Kind: w.MarkMulti}, Title: "radicale · vps · mix", Inside: "· Проекты", Aside: "CalDAV"},
												},
												Inset: true,
											},
										},
									},
									w.Msg{
										Icon: "shield",
										Parts: []w.MsgPart{
											w.Text("One code for each ticked account. Scan it with the phone: DAVx5 opens with the address and the login filled in, and asks for the password shown under the code."),
										},
										Space: w.MsgFlush,
									},
								},
							},
							w.FormColumn{
								Items: []w.FormColumnPart{
									w.Card{
										Items: []w.CardPart{
											w.Rubric{Text: "nextcloud · mix"},
											w.CardQR{
												Colour: "#FFFDF7",
												Symbol: "qr-demo",
												Side: []w.CardQRPart{
													w.Dim{Parts: []w.Inline{w.Text("Address and login are in the code.")}, Size: w.Dim11_5, Variant: w.DimWrap},
													w.Button{Label: "Open in DAVx5", Small: true},
												},
											},
											w.SpacedField{
												Field: w.Field{
													Label: "Password",
													Parts: []w.FieldPart{
														w.Line{
															Items: []w.LinePart{
																w.Input{Value: "z3-brick-lamp-9f2", Label: "Password", Mono: true, ReadOnly: true, Variant: w.InputFill},
																w.Button{Label: "Copy", Small: true},
															},
															Gap: w.LineGap6,
														},
													},
													Width: w.FieldW100,
												},
											},
										},
									},
									w.Card{
										Items: []w.CardPart{
											w.Rubric{Text: "radicale · vps · mix"},
											w.CardQR{
												Colour: "#FFFDF7",
												Symbol: "qr-demo",
												Side: []w.CardQRPart{
													w.Dim{Parts: []w.Inline{w.Text("Address and login are in the code.")}, Size: w.Dim11_5, Variant: w.DimWrap},
													w.Button{Label: "Open in DAVx5", Small: true},
												},
											},
											w.SpacedField{
												Field: w.Field{
													Label: "Password",
													Parts: []w.FieldPart{
														w.Line{
															Items: []w.LinePart{
																w.Input{Value: "q7-otter-fern-41c", Label: "Password", Mono: true, ReadOnly: true, Variant: w.InputFill},
																w.Button{Label: "Copy", Small: true},
															},
															Gap: w.LineGap6,
														},
													},
													Width: w.FieldW100,
												},
											},
										},
									},
								},
							},
						},
						Variant: w.FormFlushBottom,
					},
				},
			},
		},
	}
}

// Plan3 is frame «Показать пароль» (s-plan-3).
func Plan3() Screen {
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
								On: 4,
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
							Title: &w.HeadTitle{Text: "Set up a new device"},
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Everything discovery found — as one file or one code")}}},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Seg{
										Options: []w.SegOption{{Label: "Apple", On: true}, {Label: "Android"}, {Label: "Thunderbird"}, {Label: "Anything else"}},
									},
								},
							},
						},
						w.Form{
							Items: []w.FormPart{
								w.FormColumn{
									Wide: true,
									Items: []w.FormColumnPart{
										w.Card{
											Items: []w.CardPart{
												w.Rubric{Text: "What goes into the profile"},
												w.List{
													Rows: []w.ListRow{
														w.SlimRow{
															Mark:   w.Mark{Kind: w.MarkMulti},
															Title:  "nextcloud · mix",
															Inside: "· Личный, Работа, Контакты",
															Aside:  "CalDAV, CardDAV",
														},
													},
													Inset: true,
												},
											},
										},
									},
								},
								w.FormColumn{
									Items: []w.FormColumnPart{
										w.Card{
											Items: []w.CardPart{
												w.Rubric{Text: "Profile"},
												w.CardHint{
													Parts: []w.CardTextPart{
														w.Text("Confirm your password first: the profile hands out the credentials of every included account."),
													},
												},
												w.CardRow{Kind: w.CardActs, Items: []w.CardRowPart{w.Button{Label: "Generate profile", Kind: w.ButtonDanger}}},
											},
										},
									},
								},
							},
							Variant: w.FormFlushBottom,
						},
					},
				},
				Height: w.ShellHeight340,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Title: "Show password",
					Head:  []w.SheetHeadPart{w.Dim{Parts: []w.Inline{w.Text("nextcloud")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}},
					Body: []w.SheetPart{
						w.Field{
							Label: "Your Carrel password",
							Parts: []w.FieldPart{w.Input{Value: "••••••••••", Label: "Your Carrel password", Focus: true}},
							Width: w.FieldW100,
						},
						w.Dim{
							Parts: []w.Inline{
								w.Text("Unlocks credentials everywhere for five minutes: passwords, device profiles, anything else handed out."),
							},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Show", Kind: w.ButtonPrimary},
						w.Button{Label: "Cancel"},
						w.Kbd{Keys: "Esc", Place: w.Pushed},
					},
					Width: w.SheetWidth400,
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Plan4 is frame «Thunderbird» (s-plan-4).
func Plan4() Screen {
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
							On: 4,
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
						Title: &w.HeadTitle{Text: "Set up a new device"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{
									w.Text("Thunderbird · Calendar and Address Book URLs "),
									w.Badge{Label: "credentials unlocked · 4:12", Accent: true},
									w.Button{Label: "Lock now", Small: true},
								},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Seg{
									Options: []w.SegOption{{Label: "Apple"}, {Label: "Android"}, {Label: "Thunderbird", On: true}, {Label: "Anything else"}},
								},
							},
						},
					},
					w.Table{
						Cols: []w.TableCol{
							{Label: "Collection"},
							{Label: "Type"},
							{Label: "URL", Rank: w.Rank3},
							{Width: w.ColNum, Rank: w.Rank4},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{
											w.Line{
												Items: []w.LinePart{
													w.Strong("nextcloud"),
													w.Dim{Parts: []w.Inline{w.Text("login")}},
													w.Mono{Text: "mix"},
													w.Dim{Parts: []w.Inline{w.Text("password")}},
													w.Mono{Text: "z3-brick-lamp-9f2"},
													w.Button{Label: "Copy password", Small: true},
												},
												Gap: w.LineGap8,
											},
										},
										Span: 4,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Личный"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/calendars/mix/personal/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Работа"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/calendars/mix/work/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Контакты"}}},
									{Parts: []w.CellPart{w.Text("CardDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/addressbooks/mix/contacts/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
						},
					},
					w.Msg{
						Icon: "shield",
						Parts: []w.MsgPart{
							w.Text("Paste a CalDAV URL into Thunderbird Calendar, a CardDAV URL into Address Book, with the login and the password from the account row. There is no import file and none is needed."),
						},
					},
				},
			},
		},
	}
}

// Plan5 is frame «Anything else» (s-plan-5).
func Plan5() Screen {
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
							On: 4,
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
						Title: &w.HeadTitle{Text: "Set up a new device"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{
									w.Text("Exact strings for any client "),
									w.Badge{Label: "credentials unlocked · 4:12", Accent: true},
									w.Button{Label: "Lock now", Small: true},
								},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Seg{
									Options: []w.SegOption{{Label: "Apple"}, {Label: "Android"}, {Label: "Thunderbird"}, {Label: "Anything else", On: true}},
								},
							},
						},
					},
					w.Table{
						Cols: []w.TableCol{
							{Label: "Collection"},
							{Label: "Type"},
							{Label: "URL", Rank: w.Rank3},
							{Width: w.ColNum, Rank: w.Rank4},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{
											w.Line{
												Items: []w.LinePart{
													w.Strong("nextcloud"),
													w.Dim{Parts: []w.Inline{w.Text("login")}},
													w.Mono{Text: "mix"},
													w.Dim{Parts: []w.Inline{w.Text("password")}},
													w.Mono{Text: "z3-brick-lamp-9f2"},
													w.Button{Label: "Copy password", Small: true},
												},
												Gap: w.LineGap8,
											},
										},
										Span: 4,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Личный"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/calendars/mix/personal/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Работа"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/calendars/mix/work/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Контакты"}}},
									{Parts: []w.CellPart{w.Text("CardDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/addressbooks/mix/contacts/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{
											w.Line{
												Items: []w.LinePart{
													w.Strong("sftpgo · vps"),
													w.Dim{Parts: []w.Inline{w.Text("login")}},
													w.Mono{Text: "mix"},
													w.Dim{Parts: []w.Inline{w.Text("password")}},
													w.Mono{Text: "m2-kettle-pine-08d"},
													w.Button{Label: "Copy password", Small: true},
												},
												Gap: w.LineGap8,
											},
										},
										Span: 4,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Documents"}}},
									{Parts: []w.CellPart{w.Text("WebDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("https://files.example.org/dav/Documents/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
						},
					},
					w.Msg{
						Icon: "shield",
						Parts: []w.MsgPart{
							w.Text("Credentials are unlocked for five minutes — here and everywhere else they are handed out. Everything shown closes when the time is up or when you press "),
							w.Strong("Lock now"),
							w.Text("; every reveal is written to the audit log."),
						},
					},
				},
			},
		},
	}
}
