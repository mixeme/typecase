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
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 3}, Href: "#"},
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
													w.Input{Value: "Михаил Енученко", Label: "Display name", Variant: w.InputFill},
													w.Button{Label: "Save"},
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
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Email"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Address",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{w.Input{Value: "mix@example.org", Variant: w.InputFill}, w.Button{Label: "Change"}},
												Gap:   w.LineGap8,
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Used for service notices, and as "),
													w.Mono{Text: "Reply-To"},
													w.Text(" on letters you send."),
												},
											},
											w.Line{Items: []w.LinePart{w.Badge{Label: "pending: new@example.org"}}},
										},
										Width: w.FieldLG,
									},
								},
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
													{},
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
						Note:   "Deletes the account and the saved connections. Data on the DAV servers stays: Carrel never touches it.",
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
									{Icon: "server", Label: "Connections", Count: &w.Num{Value: 3}, Href: "#"},
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
											Parts: []w.FieldPart{w.FieldHint{Parts: []w.Inline{w.Text("No copy of your data key is deposited.")}}},
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
					Icon:  "shield",
					Tone:  w.SheetIconInk,
					Title: "Deposit my key",
					Body: []w.SheetPart{
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("An administrator who holds the escrow master password will be able to decrypt the DAV credentials you save here, and through them your contacts, calendars, tasks, notes and files. Ending sessions does not undo this. You can withdraw the copy later from Account."),
							},
							Tone:  w.MsgAlert,
							Space: w.MsgFlush,
						},
						w.Field{
							Label: "Your Carrel password",
							Parts: []w.FieldPart{
								w.Input{Value: "••••••••••", Label: "Your Carrel password", Focus: true},
								w.FieldHint{
									Parts: []w.Inline{
										w.Text("Asked because the key is being handed over, not just looked at. Same unlock as showing a saved password."),
									},
								},
							},
							Width: w.FieldW100,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Deposit my key", Kind: w.ButtonPrimary},
						w.Button{Label: "Cancel"},
						w.Kbd{Keys: "Esc", Place: w.Pushed},
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
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 3}, Href: "#"},
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
											w.Input{Value: "Михаил Енученко", Label: "Display name"},
											w.FieldHint{
												Parts: []w.Inline{w.Text("Shown as the sender and used as the signature in letters you send. Blank means your login.")},
											},
										},
										Width: w.FieldLG,
									},
								},
								Variant: w.FieldSetLead,
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Email"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Address",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{w.Input{Value: "mix@example.org", Variant: w.InputFill}, w.Button{Label: "Change"}},
												Gap:   w.LineGap8,
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Used for service notices, and as "),
													w.Mono{Text: "Reply-To"},
													w.Text(" on letters you send."),
												},
											},
											w.Msg{
												Parts: []w.MsgPart{
													w.Text("Nobody can reply to letters you send: until this address is confirmed, they go out without a reply address and say so at the foot. "),
													w.Button{Label: "Send confirmation", Small: true},
												},
												Tone:  w.MsgAlert,
												Space: w.MsgFlushTop4,
											},
										},
										Width: w.FieldLG,
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

// Account3 is frame «Смена пароля · обычная и вынужденная» (account-3).
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
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 3}, Href: "#"},
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
							w.Msg{
								Parts: []w.MsgPart{w.Text("You are signed in with a temporary password. Choose a new one before continuing.")},
								Tone:  w.MsgAlert,
								Space: w.MsgFlushSides,
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Current password",
										Parts: []w.FieldPart{w.Input{Value: "••••••••••", Label: "Current password"}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "New password",
										Parts: []w.FieldPart{w.Input{Value: "••••••••••••", Label: "New password", Focus: true}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Repeat new password",
										Parts: []w.FieldPart{w.Input{Value: "••••••••••••", Label: "Repeat new password"}},
										Width: w.FieldLG,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Your data stays encrypted under the same key; only the wrapping changes. Saved connections keep working."),
												},
											},
										},
										Place: w.Grown,
									},
								},
								Variant: w.FormFlush,
							},
							w.FormFoot{
								Items:   []w.FormFootPart{w.Button{Label: "Change password", Kind: w.ButtonPrimary}, w.Button{Label: "Back to settings"}},
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
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 3}, Href: "#"},
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
									w.Text("A copy of your data key is deposited on this server. An administrator who holds the escrow master password can decrypt the DAV credentials you save here, and through them your contacts and calendars."),
								},
								Tone:  w.MsgAlert,
								Space: w.MsgFlushSides,
							},
							w.Fields{
								Items: []w.FieldsItem{
									{Term: "Deposited", Value: []w.ValuePart{w.Text("2026-03-02")}, Face: w.ValueMono},
									{Term: "Last recovery", Value: []w.ValuePart{w.Text("never")}, Face: w.ValueMono},
								},
								Variant: w.FieldsBleed,
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Withdraw my deposited key", Kind: w.ButtonDangerSolid},
									w.Dim{
										Parts:   []w.Inline{w.Text("After withdrawal nobody can recover this account, including you.")},
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
							w.MenuItem{Icon: "search", Trail: w.Kbd{Keys: "?"}, Label: "Keyboard shortcuts"},
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
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
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
							{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
							{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
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
						Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
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
								{Keys: "Ctrl+Enter", Does: "Save a quick note"},
								{Keys: "Ctrl+V", Does: "Paste an image as an attachment"},
							},
						},
						w.SheetHint{Text: "Shortcuts do nothing while the cursor is in a field. None of them destroy anything."},
					},
					Foot:  []w.SheetFootPart{w.Button{Label: "Close"}, w.Kbd{Keys: "Esc"}, w.Kbd{Keys: "?", Place: w.Pushed}},
					Width: w.SheetWidth480,
				},
			},
		},
	}
}
