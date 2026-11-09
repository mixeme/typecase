// Раздел макета: 7.10 Настройки · учётная запись.

package screens

import w "carrel.dev/typecase/widget"

// Account1 is frame «Настройки · учётная запись» (account-1).
func Account1() Screen {
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
												Gap: w.LineGapNormal,
											},
											w.FieldHint{
												Parts:   []w.Inline{w.Text("Shown as the sender and used as the signature in letters you send. Blank means your login.")},
												Variant: w.HintColumn,
											},
										},
										Width: w.FieldW580,
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
												Gap: w.LineGapNormal,
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Used for service notices, and as "),
													w.Mono{Text: "Reply-To"},
													w.Text(" on letters you send."),
												},
												Variant: w.HintColumn,
											},
										},
										Width: w.FieldW580,
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
												Parts:   []w.Inline{w.Text("Your saved connections stay readable after the change.")},
												Variant: w.HintColumn,
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
											{Label: "Last seen", Rank: w.Rank3, Main: true},
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
													w.Text("No copy of your data key is deposited: if you lose your password, nobody can recover the account."),
												},
												Variant: w.HintColumn,
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
		},
	}
}

// Account2 is frame «Депонирование ключа · подтверждение» (account-2).
func Account2() Screen {
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
									Rubric: w.Rubric{Text: "Key escrow"},
									Items: []w.FieldSetPart{
										w.Field{
											Parts: []w.FieldPart{
												w.FieldHint{Parts: []w.Inline{w.Text("No copy of your data key is deposited.")}, Variant: w.HintColumn},
											},
											Place: w.Grown,
										},
										w.Field{Parts: []w.FieldPart{w.Button{Label: "Deposit my key"}}},
									},
									Variant: w.FieldSetLead,
								},
							},
						},
					},
				},
				Height: w.ShellHeight340,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/settings/escrow/deposit",
					Icon:   "shield",
					Tone:   w.SheetIconInk,
					Title:  "Deposit my key",
					Body: []w.SheetPart{
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("With the escrow master password, an administrator will be able to decrypt the DAV credentials you save here, and through them your records. You can withdraw the copy later from Account."),
							},
							Tone:  w.MsgAlert,
							Space: w.MsgFlush,
						},
						w.Field{
							Label: "Your Carrel password",
							Parts: []w.FieldPart{w.Input{Name: "password", Value: "••••••••••", Label: "Your Carrel password", Focus: true}},
							Width: w.FieldW100,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Deposit my key", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGapNormal,
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

// Mail9 is frame «Учётная запись · адрес не подтверждён, письмо не отправлено» (s-mail-9).
func Mail9() Screen {
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
												Gap: w.LineGapNormal,
											},
											w.FieldHint{
												Parts:   []w.Inline{w.Text("Shown as the sender and used as the signature in letters you send. Blank means your login.")},
												Variant: w.HintColumn,
											},
										},
										Width: w.FieldW580,
									},
								},
								Variant: w.FieldSetLead,
								Action:  "/app/settings/account/name",
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Email", Items: []w.RubricPart{w.Badge{Label: "not confirmed", Tone: w.BadgeAlert}}},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Address",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Input{Name: "address", Value: "mix@example.org", Label: "Address", Variant: w.InputFill},
													w.Button{Label: "Change", Submit: true},
													w.Button{Label: "Send confirmation", To: "/app/settings/account/email/confirm", Submit: true},
												},
												Gap: w.LineGapNormal,
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Used for service notices, and as "),
													w.Mono{Text: "Reply-To"},
													w.Text(" on letters you send."),
												},
												Variant: w.HintColumn,
											},
										},
										Width: w.FieldXL,
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
												Parts:   []w.Inline{w.Text("Your saved connections stay readable after the change.")},
												Variant: w.HintColumn,
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
											{Label: "Last seen", Rank: w.Rank3, Main: true},
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
													w.Text("No copy of your data key is deposited: if you lose your password, nobody can recover the account."),
												},
												Variant: w.HintColumn,
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
		},
	}
}

// Account3 is frame «Смена пароля · обычная» (account-3).
func Account3() Screen {
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
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{Title: &w.HeadTitle{Text: "Change password"}},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Current password",
										Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "current", Value: "xxxxxxxxxx", Label: "Current password"}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "New password",
										Parts: []w.FieldPart{
											w.Input{Kind: w.InputPassword, Name: "password", Value: "xxxxxxxxxxxx", Label: "New password", Focus: true},
										},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Repeat new password",
										Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "repeat", Value: "xxxxxxxxxxxx", Label: "Repeat new password"}},
										Width: w.FieldLG,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Your data stays encrypted under the same key; only the wrapping changes. Saved connections keep working."),
												},
												Variant: w.HintColumn,
											},
										},
										Place: w.Grown,
									},
								},
								ID:      "password",
								Action:  "/app/settings/account/password",
								Fields:  []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
								Variant: w.FormFlush,
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Change password", Kind: w.ButtonPrimary, Form: "password"},
									w.Button{Label: "Back to settings", Href: "/app/settings/account"},
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

// Account17 is frame «Смена пароля · вынужденная» (account-17).
func Account17() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:   w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Actions: []w.TopAction{w.UserMenu{Name: "mix"}},
		},
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{Title: &w.HeadTitle{Text: "Change password"}},
							w.Msg{
								Parts:  []w.MsgPart{w.Text("You are signed in with a temporary password. Choose a new one before continuing.")},
								Tone:   w.MsgAlert,
								Space:  w.MsgFlushSides,
								Column: true,
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Current password",
										Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "current", Value: "xxxxxxxxxx", Label: "Current password"}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "New password",
										Parts: []w.FieldPart{
											w.Input{Kind: w.InputPassword, Name: "password", Value: "xxxxxxxxxxxx", Label: "New password", Focus: true},
										},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Repeat new password",
										Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "repeat", Value: "xxxxxxxxxxxx", Label: "Repeat new password"}},
										Width: w.FieldLG,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Your data stays encrypted under the same key; only the wrapping changes. Saved connections keep working."),
												},
												Variant: w.HintColumn,
											},
										},
										Place: w.Grown,
									},
								},
								ID:      "password",
								Action:  "/app/settings/account/password",
								Fields:  []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
								Variant: w.FormFlush,
							},
							w.FormFoot{
								Items:   []w.FormFootPart{w.Button{Label: "Change password", Kind: w.ButtonPrimary, Form: "password"}},
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

// Account4 is frame «Настройки · учётная запись · депонирование ключа» (account-4).
func Account4() Screen {
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
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{
								Title: &w.HeadTitle{Text: "Key escrow"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("A copy of your key on this server")}}},
							},
							w.Msg{
								Parts: []w.MsgPart{
									w.Text("With the escrow master password, an administrator can decrypt the DAV credentials you save here, and through them your records."),
								},
								Tone:   w.MsgAlert,
								Space:  w.MsgFlushSides,
								Column: true,
							},
							w.Fields{
								Items: []w.FieldsItem{
									{Term: "Deposited", Value: []w.ValuePart{w.Text("2026-03-02")}, Face: w.ValueMono},
									{Term: "Last recovery", Value: []w.ValuePart{w.Text("never")}, Face: w.ValueMono},
								},
								Variant: w.FieldsBleed,
							},
							w.Field{
								Label: "Password history",
								Parts: []w.FieldPart{
									w.Line{
										Items: []w.LinePart{
											w.Input{
												Value:    "https://carrel.example/password-history/k7q2-m9xd-41fa",
												Label:    "Password history",
												Mono:     true,
												ReadOnly: true,
												Variant:  w.InputFill,
											},
											w.Button{
												Label:  "Copy",
												Icon:   "dup",
												Small:  true,
												Copy:   "https://carrel.example/password-history/k7q2-m9xd-41fa",
												Copied: "Copied",
											},
										},
										Gap: w.LineGapNormal,
									},
									w.FieldHint{
										Parts:   []w.Inline{w.Text("Keep it: it shows what was done to your password, even when you cannot sign in.")},
										Variant: w.HintColumn,
									},
								},
								Width: w.FieldFull,
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Withdraw my deposited key", Kind: w.ButtonDangerSolid},
									w.Dim{
										Parts:   []w.Inline{w.Text("Then nobody can recover the account.")},
										Size:    w.Dim11_5,
										Variant: w.DimWrap,
									},
								},
								Variant: w.FormFootFlush,
							},
						},
						Variant: w.Dialog560,
					},
				},
			},
		},
	}
}

// Account16 is frame «Депонирование ключа · ключ уже восстанавливали» (account-16).
func Account16() Screen {
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
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{
								Title: &w.HeadTitle{Text: "Key escrow"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("A copy of your key on this server")}}},
							},
							w.Msg{
								Parts: []w.MsgPart{
									w.Text("With the escrow master password, an administrator can decrypt the DAV credentials you save here, and through them your records."),
								},
								Tone:   w.MsgAlert,
								Space:  w.MsgFlushSides,
								Column: true,
							},
							w.Fields{
								Items: []w.FieldsItem{
									{Term: "Deposited", Value: []w.ValuePart{w.Text("2026-03-02")}, Face: w.ValueMono},
									{
										Term:  "Recovered",
										Value: []w.ValuePart{w.Mono{Text: "2026-07-11 14:02"}, w.Text(" · olga, administrator: the password was reset")},
									},
								},
								Variant: w.FieldsBleed,
							},
							w.Field{
								Label: "Password history",
								Parts: []w.FieldPart{
									w.Line{
										Items: []w.LinePart{
											w.Input{
												Value:    "https://carrel.example/password-history/k7q2-m9xd-41fa",
												Label:    "Password history",
												Mono:     true,
												ReadOnly: true,
												Variant:  w.InputFill,
											},
											w.Button{
												Label:  "Copy",
												Icon:   "dup",
												Small:  true,
												Copy:   "https://carrel.example/password-history/k7q2-m9xd-41fa",
												Copied: "Copied",
											},
										},
										Gap: w.LineGapNormal,
									},
									w.FieldHint{
										Parts:   []w.Inline{w.Text("Keep it: it shows what was done to your password, even when you cannot sign in.")},
										Variant: w.HintColumn,
									},
								},
								Width: w.FieldFull,
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Withdraw my deposited key", Kind: w.ButtonDangerSolid},
									w.Dim{
										Parts:   []w.Inline{w.Text("Then nobody can recover the account.")},
										Size:    w.Dim11_5,
										Variant: w.DimWrap,
									},
								},
								Variant: w.FormFootFlush,
							},
						},
						Variant: w.Dialog560,
					},
				},
			},
		},
	}
}

// Account5 is frame «Меню пользователя» (account-5).
func Account5() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.TopDrop{
					Open:    true,
					Trigger: w.UserMenu{Name: "mix", Focus: true},
					Menu: w.Menu{
						Items: []w.MenuPart{
							w.MenuHeading{
								Rubric: w.Rubric{
									Text:    "mix",
									Items:   []w.RubricPart{w.Mono{Text: "carrel.example", Variant: w.MonoTightPlain}},
									Form:    w.RubricItem,
									Variant: w.RubricItemLine,
								},
								Form: w.MenuHeadingSplit,
							},
							w.MenuItem{Icon: "gear", Label: "Settings"},
							w.MenuItem{Icon: "adm", Label: "Administration"},
							w.MenuSep{},
							w.MenuColumn{
								Label: "Theme",
								Control: w.Seg{
									Options: []w.SegOption{{Label: "Auto", On: true}, {Label: "Light"}, {Label: "Dark"}},
									Variant: w.SegStartSmall,
								},
							},
							w.MenuItem{
								Icon:       "search",
								Trail:      w.Kbd{Keys: "?"},
								Label:      "Keyboard shortcuts",
								ScriptOnly: true,
								Href:       "#keys",
								Key:        "?",
							},
							w.MenuSep{},
							w.MenuItem{Label: "Sign out", Form: "post-signout"},
						},
						Width: w.MenuWidth290,
					},
					End: true,
				},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 7}, Href: "#"},
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
							Meta:    "7",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "3"},
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
					w.PostForm{ID: "post-signout", Action: "/logout", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}},
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
			Height: w.ShellHeight340,
		},
	}
}

// Account6 is frame «Клавиши · лист поверх текущего экрана» (account-6).
func Account6() Screen {
	return w.Shell{
		Body: w.Overlay{
			Top: &w.Top{
				Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
				Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
				Actions: []w.TopAction{
					w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
					w.Button{Label: "2 min ago", Icon: "refresh"},
					w.UserMenu{Name: "mix"},
				},
			},
			Body: w.OverlayBody{
				Rail: []w.OverlayRailPart{
					w.Nav{
						Items: []w.NavItem{
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 7}, Href: "#"},
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
								Meta:    "7",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "3"},
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
										w.Text(" · 7 events"),
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
				Height: w.ShellHeight340,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Icon:  "search",
					Title: "Keyboard shortcuts",
					Head:  []w.SheetHeadPart{w.Close{Label: "Close"}},
					Body: []w.SheetPart{
						w.SheetKeys{
							Items: []w.SheetKey{
								{Keys: "N", Does: "Quick note over this screen"},
								{Keys: "/", Does: "Focus the search field in the header"},
								{Keys: "Enter", Does: "Open the selected record full screen"},
								{Keys: "Esc", Does: "Close a sheet, the details panel or focus mode"},
								{Keys: "Ctrl+S", Does: "Save in the editor"},
								{Keys: "Ctrl+Enter", Does: "Save or send a sheet with text: a quick note, a letter"},
								{Keys: "Ctrl+V", Does: "Paste an image as an attachment"},
							},
						},
						w.SheetHint{Text: "Shortcuts do nothing while the cursor is in a field. None of them destroy anything."},
					},
					Foot:  []w.SheetFootPart{w.Button{Label: "Close"}, w.Kbd{Keys: "Esc"}, w.Kbd{Keys: "?", Place: w.Pushed}},
					Width: w.SheetWidth480,
				},
				ID: "keys",
			},
		},
	}
}
