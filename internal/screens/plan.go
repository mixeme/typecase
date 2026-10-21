// Раздел макета: 7.13 Настройки · новое устройство.

package screens

import w "carrel.dev/typecase/widget"

// Plan1 is frame «Новое устройство · профиль Apple» (s-plan-1).
func Plan1() Screen {
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
														Inside: "· 9 collections",
														Aside:  "CalDAV, CardDAV",
													},
													w.SlimRow{
														Ticked: w.TickedOff,
														Mark:   w.Mark{Kind: w.MarkMulti},
														Title:  "radicale · vps · mix",
														Inside: "· 4 collections",
														Aside:  "CalDAV, CardDAV",
													},
													w.SlimRow{
														Ticked: w.TickedOff,
														Reason: "The server refused the saved password.",
														Mark:   w.Mark{Kind: w.MarkMulti},
														Title:  "baikal · m.yenuchenko",
														Inside: "· 1 collection",
														Aside:  "password refused",
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
													w.Badge{
														Label: "local",
														Tone:  w.BadgeLocal,
														Title: "Local instance: the link will not open from outside your network",
													},
												},
											},
											w.CardHint{
												Parts: []w.CardTextPart{
													w.Text("It opens once, for you signed in to Carrel. Open it on the device, install the profile, then delete the file."),
												},
											},
											w.CardRow{
												Kind: w.CardTop9,
												Gap:  6,
												Items: []w.CardRowPart{
													w.Button{Label: "Download profile", Icon: "download"},
													w.Button{Label: "Revoke", Kind: w.ButtonDanger},
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

// Plan6 is frame «Новое устройство · выдача закрыта» (s-plan-6).
func Plan6() Screen {
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
														Inside: "· 9 collections",
														Aside:  "CalDAV, CardDAV",
													},
													w.SlimRow{
														Ticked: w.TickedOff,
														Mark:   w.Mark{Kind: w.MarkMulti},
														Title:  "radicale · vps · mix",
														Inside: "· 4 collections",
														Aside:  "CalDAV, CardDAV",
													},
													w.SlimRow{
														Ticked: w.TickedOff,
														Reason: "The server refused the saved password.",
														Mark:   w.Mark{Kind: w.MarkMulti},
														Title:  "baikal · m.yenuchenko",
														Inside: "· 1 collection",
														Aside:  "password refused",
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
											w.CardHint{Parts: []w.CardTextPart{w.Text("Asks for your Carrel password first.")}},
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
		},
	}
}

// Plan2 is frame «Android · DAVx5» (s-plan-2).
func Plan2() Screen {
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
														Inside: "· 9 collections",
														Aside:  "CalDAV, CardDAV",
													},
													w.SlimRow{
														Mark:   w.Mark{Kind: w.MarkMulti},
														Title:  "radicale · vps · mix",
														Inside: "· 4 collections",
														Aside:  "CalDAV, CardDAV",
													},
													w.SlimRow{
														Ticked: w.TickedOff,
														Reason: "The server refused the saved password.",
														Mark:   w.Mark{Kind: w.MarkMulti},
														Title:  "baikal · m.yenuchenko",
														Inside: "· 1 collection",
														Aside:  "password refused",
													},
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
												Src:    "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGwAAABsAQMAAAC1qMWOAAAABlBMVEX//fcAAADqtqjxAAAAwUlEQVR42qWTsW1FMQwDz0F6eoO3/1h/A3qCS5EgwK8COeoE+QCJNOEfFTWticrH+/Cvdq1N4LDWnrKRc0AiUxaAZw8ev7fhtW/Ys2DxvDhryH4ibOBBWCP2xyjofGfAxJgSp2xsLASmbKIqWjJlTROwOL83UYp2zhbaNNKxVpHgt9RDtmBL6cW9TYAyzu8vTzv3V8XWZurvB2vtnLA58+xH2PuBnE4zCCCbdN2w58DZJ+P/DEktam90Tmxy6e9VfQFbsnKSXbiN/QAAAABJRU5ErkJggg==",
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
												Src:    "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGwAAABsAQMAAAC1qMWOAAAABlBMVEX//fcAAADqtqjxAAAAwUlEQVR42qWTsW1FMQwDz0F6eoO3/1h/A3qCS5EgwK8COeoE+QCJNOEfFTWticrH+/Cvdq1N4LDWnrKRc0AiUxaAZw8ev7fhtW/Ys2DxvDhryH4ibOBBWCP2xyjofGfAxJgSp2xsLASmbKIqWjJlTROwOL83UYp2zhbaNNKxVpHgt9RDtmBL6cW9TYAyzu8vTzv3V8XWZurvB2vtnLA58+xH2PuBnE4zCCCbdN2w58DZJ+P/DEktam90Tmxy6e9VfQFbsnKSXbiN/QAAAABJRU5ErkJggg==",
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
					Action: "/app/settings/unlock",
					Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "back", Value: "/app/settings/devices"}},
					Title:  "Generate profile",
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
						w.Button{Label: "Generate", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/devices"}, w.Kbd{Keys: "Esc"}},
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

// Davloom5 is frame «davloom · пароль перед выдачей» (s-davloom-5).
func Davloom5() Screen {
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
												w.FieldHint{Parts: []w.Inline{w.Text("For you: this is how you will know what you are revoking later.")}},
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
									Parts: []w.MsgPart{
										w.Text("The password is shown once and stored nowhere. Lose it and you create a new one and revoke this — it cannot be recovered."),
									},
									Tone:  w.MsgAlert,
									Space: w.MsgFlushSides,
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
				Height: w.ShellHeight340,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/settings/unlock",
					Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "back", Value: "/app/settings/davloom/devices/new"}},
					Title:  "Add device",
					Head: []w.SheetHeadPart{
						w.Dim{Parts: []w.Inline{w.Text("Телефон · DAVx5")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed},
					},
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
						w.Button{Label: "Create password", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/davloom/devices/new"}, w.Kbd{Keys: "Esc"}},
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

// Account15 is frame «Учётная запись · удаление» (account-15).
func Account15() Screen {
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
								On: 1,
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
							Title: &w.HeadTitle{Text: "Account"},
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("mix · created 11 February 2026")}}},
						},
						w.Form{
							Items: []w.FormPart{
								w.FieldSet{
									Rubric: w.Rubric{Text: "Name"},
									Items: []w.FieldSetPart{
										w.Field{
											Label: "Display name",
											Parts: []w.FieldPart{
												w.Line{
													Items: []w.LinePart{
														w.Input{Name: "display.name", Value: "Михаил Енученко", Label: "Display name", Variant: w.InputFill},
														w.Button{Label: "Save", Submit: true},
													},
													Gap: w.LineGap8,
												},
												w.FieldHint{
													Parts: []w.Inline{w.Text("Shown as the sender and used as the signature in letters you send. Blank means your login.")},
												},
											},
											Width: w.FieldLG,
										},
									},
									Variant: w.FieldSetLead,
									Action:  "/app/settings/account/name",
								},
								w.FieldSet{
									Rubric: w.Rubric{Text: "Email", Items: []w.RubricPart{w.Badge{Label: "pending"}, w.Plain{Text: "new@example.org"}}},
									Items: []w.FieldSetPart{
										w.Field{
											Label: "Address",
											Parts: []w.FieldPart{
												w.Line{
													Items: []w.LinePart{
														w.Input{Name: "address", Value: "mix@example.org", Label: "Address", Variant: w.InputFill},
														w.Button{Label: "Change", Submit: true},
													},
													Gap: w.LineGap8,
												},
												w.FieldHint{
													Parts: []w.Inline{
														w.Text("Used for service notices, and as "),
														w.Mono{Text: "Reply-To"},
														w.Text(" on letters you send."),
													},
												},
											},
											Width: w.FieldLG,
										},
									},
									Action: "/app/settings/account/email",
								},
								w.FieldSet{
									Rubric: w.Rubric{Text: "Password"},
									Items: []w.FieldSetPart{
										w.Field{
											Parts: []w.FieldPart{
												w.FieldHint{
													Parts: []w.Inline{w.Text("Changing your password re-wraps your data key only; saved connections stay readable.")},
												},
											},
											Place: w.Grown,
										},
										w.Field{Parts: []w.FieldPart{w.Button{Label: "Change password"}}},
									},
								},
								w.FieldSet{
									Rubric: w.Rubric{Text: "Sessions", Items: []w.RubricPart{w.Plain{Text: "2 active"}}},
									Items: []w.FieldSetPart{
										w.Table{
											Cols: []w.TableCol{
												{Label: "Where"},
												{Label: "Started"},
												{Label: "Last seen", Rank: w.Rank3},
												{Width: w.Col120, Rank: w.Rank4},
											},
											Rows: []w.TableRow{
												{
													Cells: []w.TableCell{
														{Parts: []w.CellPart{w.Text("Firefox · Windows "), w.Badge{Label: "this one"}}},
														{Parts: []w.CellPart{w.Text("2026-08-16 09:02")}, Mono: true},
														{Parts: []w.CellPart{w.Text("just now")}, Mono: true},
														{
															Parts: []w.CellPart{
																w.Button{
																	Label: "End",
																	Kind:  w.ButtonDanger,
																	Small: true,
																	Title: "This is the session you are using: sign out instead.",
																	State: w.ButtonDisabled,
																},
															},
															Acts: true,
														},
													},
												},
												{
													Cells: []w.TableCell{
														{Parts: []w.CellPart{w.Text("Carrel desktop · Windows")}},
														{Parts: []w.CellPart{w.Text("2026-08-14 18:40")}, Mono: true},
														{Parts: []w.CellPart{w.Text("2 hours ago")}, Mono: true},
														{Parts: []w.CellPart{w.Button{Label: "End", Kind: w.ButtonDanger, Small: true}}, Acts: true},
													},
												},
											},
											Variant: w.TableInset,
											Place:   w.Grown,
										},
										w.Field{Parts: []w.FieldPart{w.Button{Label: "End every other session", Kind: w.ButtonDanger}}},
									},
								},
								w.FieldSet{
									Rubric: w.Rubric{Text: "Key escrow"},
									Items: []w.FieldSetPart{
										w.Field{
											Parts: []w.FieldPart{
												w.FieldHint{
													Parts: []w.Inline{
														w.Text("No copy of your data key is deposited. Nobody but you can decrypt your data — and if you lose your password, it cannot be recovered by anyone. Depositing asks for your Carrel password first: it is the same unlock that shows a saved password, and for the same reason — the key is being handed over, not just looked at."),
													},
												},
											},
											Place: w.Grown,
										},
										w.Field{Parts: []w.FieldPart{w.Button{Label: "Deposit my key"}}},
									},
								},
							},
						},
						w.Danger{
							Button: w.Button{Label: "Delete account", Kind: w.ButtonDangerSolid},
							Note:   "Deletes the account and the saved connections. Data on the DAV servers stays.",
						},
					},
				},
				Height: w.ShellHeight340,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/settings/unlock",
					Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "back", Value: "/app/settings/account"}},
					Icon:   "trash",
					Tone:   w.SheetIconDanger,
					Title:  "Delete account",
					Body: []w.SheetPart{
						w.Msg{
							Parts: []w.MsgPart{w.Text("Deletes the account and the saved connections. Data on the DAV servers stays.")},
							Tone:  w.MsgAlert,
							Space: w.MsgFlush,
						},
						w.Field{
							Label: "Your Carrel password",
							Parts: []w.FieldPart{
								w.Input{Kind: w.InputPassword, Name: "password", Label: "Your Carrel password", Focus: true, Required: true},
							},
							Width: w.FieldW100,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Delete account", Kind: w.ButtonDangerSolid, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/account"}, w.Kbd{Keys: "Esc"}},
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

// Plan4 is frame «Thunderbird» (s-plan-4).
func Plan4() Screen {
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
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/calendars/mix/lichnyj/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Дела"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/calendars/mix/dela/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Работа"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/calendars/mix/rabota/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "Дни рождения"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{
										Parts: []w.CellPart{w.Text("…/remote.php/dav/calendars/mix/dni-rozhdeniya/")},
										Mono:  true,
										Small: true,
									},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Записи"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/calendars/mix/zapisi/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Дневник"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/calendars/mix/dnevnik/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Общие с семьёй"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{
										Parts: []w.CellPart{w.Text("…/remote.php/dav/calendars/mix/zapiski-obshchie/")},
										Mono:  true,
										Small: true,
									},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Контакты"}}},
									{Parts: []w.CellPart{w.Text("CardDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/addressbooks/mix/kontakty/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "Архив"}}},
									{Parts: []w.CellPart{w.Text("CardDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/addressbooks/mix/arhiv/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{
											w.Line{
												Items: []w.LinePart{
													w.Strong("radicale · vps"),
													w.Dim{Parts: []w.Inline{w.Text("login")}},
													w.Mono{Text: "mix"},
													w.Dim{Parts: []w.Inline{w.Text("password")}},
													w.Mono{Text: "q7-otter-fern-41c"},
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
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Проекты"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/radicale/mix/proekty/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#8C6239"}, Name: "Дежурства"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/radicale/mix/dezhurstva/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Команда"}}},
									{Parts: []w.CellPart{w.Text("CardDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/radicale/mix/komanda/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Рассылки"}}},
									{Parts: []w.CellPart{w.Text("CardDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/radicale/mix/rassylki/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{
											w.Line{
												Items: []w.LinePart{
													w.Strong("baikal"),
													w.Dim{Parts: []w.Inline{w.Text("login")}},
													w.Mono{Text: "m.yenuchenko"},
													w.Faced{Face: w.FaceDanger, Items: []w.LinePart{w.Text("the server refused the saved password")}},
												},
												Gap: w.LineGap8,
											},
										},
										Span: 4,
									},
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
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/calendars/mix/lichnyj/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Дела"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/calendars/mix/dela/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Работа"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/calendars/mix/rabota/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "Дни рождения"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{
										Parts: []w.CellPart{w.Text("…/remote.php/dav/calendars/mix/dni-rozhdeniya/")},
										Mono:  true,
										Small: true,
									},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Записи"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/calendars/mix/zapisi/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Дневник"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/calendars/mix/dnevnik/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Общие с семьёй"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{
										Parts: []w.CellPart{w.Text("…/remote.php/dav/calendars/mix/zapiski-obshchie/")},
										Mono:  true,
										Small: true,
									},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "Контакты"}}},
									{Parts: []w.CellPart{w.Text("CardDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/addressbooks/mix/kontakty/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "Архив"}}},
									{Parts: []w.CellPart{w.Text("CardDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/addressbooks/mix/arhiv/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#3F6E8C"}, Name: "Files"}}},
									{Parts: []w.CellPart{w.Text("WebDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/remote.php/dav/files/mix/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{
											w.Line{
												Items: []w.LinePart{
													w.Strong("radicale · vps"),
													w.Dim{Parts: []w.Inline{w.Text("login")}},
													w.Mono{Text: "mix"},
													w.Dim{Parts: []w.Inline{w.Text("password")}},
													w.Mono{Text: "q7-otter-fern-41c"},
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
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Проекты"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/radicale/mix/proekty/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#8C6239"}, Name: "Дежурства"}}},
									{Parts: []w.CellPart{w.Text("CalDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/radicale/mix/dezhurstva/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Команда"}}},
									{Parts: []w.CellPart{w.Text("CardDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/radicale/mix/komanda/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "Рассылки"}}},
									{Parts: []w.CellPart{w.Text("CardDAV")}, Mono: true},
									{Parts: []w.CellPart{w.Text("…/radicale/mix/rassylki/")}, Mono: true, Small: true},
									{Parts: []w.CellPart{w.Button{Label: "Copy", Small: true}}, Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{
											w.Line{
												Items: []w.LinePart{
													w.Strong("baikal"),
													w.Dim{Parts: []w.Inline{w.Text("login")}},
													w.Mono{Text: "m.yenuchenko"},
													w.Faced{Face: w.FaceDanger, Items: []w.LinePart{w.Text("the server refused the saved password")}},
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
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#9A9280"}, Name: "Public"}}},
									{Parts: []w.CellPart{w.Text("WebDAV "), w.Badge{Label: "ro"}}, Mono: true},
									{Parts: []w.CellPart{w.Text("https://files.example.org/dav/Public/")}, Mono: true, Small: true},
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
