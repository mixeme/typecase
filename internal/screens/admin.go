// Раздел макета: 7.19 Администрирование.

package screens

import w "carrel.dev/typecase/widget"

// Admin1 is frame «Администрирование · как заводят пользователей» (s-admin-1).
func Admin1() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Administration"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "adm", Label: "Settings", Href: "#"},
								{Icon: "shield", Label: "Installation", Href: "#"},
								{Icon: "search", Label: "DAV validator", Href: "#"},
								{Icon: "dup", Label: "Key escrow", Href: "#"},
								{Icon: "clock", Label: "Audit log", Href: "#"},
							},
							On: 1,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Text:   "Mail is going out",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last sent "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
						w.RailStatus{
							Text:   "Checks passed",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last check "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Add users"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{w.Text("Three ways, for three situations. There is no global mode: any of them works.")},
							},
						},
					},
					w.MainCards{
						Cards: []w.Card{
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Invitation", Items: []w.RubricPart{w.Plain{Text: "chooses their own password"}}},
									w.CardFields{
										Items: []w.CardFieldsPart{
											w.Field{
												Label: "Role",
												Parts: []w.FieldPart{
													w.Pick{Name: "role", Options: []w.PickOption{{Value: "user", Label: "User", On: true}}, Label: "Role"},
												},
												Width: w.FieldSM,
											},
											w.Button{Label: "Create link", Kind: w.ButtonPrimary, Submit: true},
										},
										Wide: true,
									},
									w.CardRow{
										Kind: w.CardTop10,
										Gap:  6,
										Items: []w.CardRowPart{
											w.Input{
												Value:    "https://carrel.example/invite/7f3a…",
												Label:    "Invitation link",
												Mono:     true,
												ReadOnly: true,
												Variant:  w.Input11Fill,
											},
											w.Button{Label: "Copy", Small: true, Copy: "https://carrel.example/invite/7f3a…", Copied: "Copied"},
										},
									},
									w.CardFields{
										Items: []w.CardFieldsPart{
											w.Field{
												Label: "Or send to",
												Parts: []w.FieldPart{w.Input{Name: "email", Placeholder: "anna@example.org", Label: "Or send to"}},
												Width: w.FieldMD,
											},
											w.Button{Label: "Send invitation", To: "/admin/invites?action=send", Submit: true},
										},
										Wide: true,
										Rule: true,
									},
								},
								Action: "/admin/invites?action=link",
								Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text:  "Temporary password",
										Items: []w.RubricPart{w.Badge{Label: "hand to hand", Tone: w.BadgeAlert}},
									},
									w.Msg{
										Parts: []w.MsgPart{w.Text("Until the user changes this password, you can sign in as them.")},
										Tone:  w.MsgAlert,
										Space: w.MsgFlushBottom12,
									},
									w.CardFields{
										Items: []w.CardFieldsPart{
											w.Field{
												Label: "Role",
												Parts: []w.FieldPart{
													w.Pick{Name: "role", Options: []w.PickOption{{Value: "user", Label: "User", On: true}}, Label: "Role"},
												},
												Width: w.FieldXS,
											},
											w.Field{
												Label: "Login",
												Parts: []w.FieldPart{w.Input{Name: "login", Value: "viktor", Label: "Login"}},
												Width: w.FieldSM,
											},
											w.Field{
												Label:      "Email",
												LabelAside: w.Dim{Parts: []w.Inline{w.Text("(optional)")}},
												Parts:      []w.FieldPart{w.Input{Name: "email", Label: "Email"}},
												Width:      w.FieldMD,
											},
											w.Field{
												Label: "Temporary password",
												Parts: []w.FieldPart{w.Input{Name: "password", Value: "9f2-brick-lamp", Label: "Temporary password", Mono: true}},
												Width: w.FieldMD,
											},
											w.Button{Label: "Create user", Kind: w.ButtonPrimary, Submit: true},
										},
										Wide: true,
									},
								},
								Action: "/admin/invites",
								Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "action", Value: "user"}},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Self-registration", Items: []w.RubricPart{w.Badge{Label: "open"}}},
									w.CardText{
										Kind: w.CardLede,
										Parts: []w.CardTextPart{
											w.Text("A visitor creates an account and confirms the address by email. Cannot be turned on until mail works."),
										},
									},
									w.Check{
										On:    true,
										Name:  "selfReg",
										Value: "on",
										Parts: []w.CheckPart{w.Text("Publish the registration form")},
									},
									w.CardRow{
										Kind: w.CardTop10,
										Gap:  6,
										Items: []w.CardRowPart{
											w.Input{
												Value:    "https://carrel.example/register",
												Label:    "Registration address",
												Mono:     true,
												ReadOnly: true,
												Variant:  w.Input11Fill,
											},
											w.Button{Label: "Copy", Small: true, Copy: "https://carrel.example/register", Copied: "Copied"},
										},
									},
									w.CardRow{Kind: w.CardTop10, Items: []w.CardRowPart{w.Button{Label: "Save", Submit: true}}},
								},
								Action: "/admin/invites",
								Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "action", Value: "selfreg"}},
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Pending invitations", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.Table{
						Cols: []w.TableCol{
							{Label: "Address"},
							{Label: "Delivery"},
							{Label: "Status", Rank: w.Rank3},
							{Label: "Mail", Rank: w.Rank4},
							{Label: "Expires"},
							{},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("—")}},
									{Parts: []w.CellPart{w.Text("link")}, Mono: true},
									{Parts: []w.CellPart{w.Text("pending")}},
									{Parts: []w.CellPart{w.Text("—")}},
									{Parts: []w.CellPart{w.Text("2026-08-18 12:00")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Label: "Copy link", Small: true, Copy: "https://carrel.example/invite/2c9d", Copied: "Copied"},
											w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
										},
										Acts: true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("anna@example.org")}, Mono: true},
									{Parts: []w.CellPart{w.Text("email")}, Mono: true},
									{Parts: []w.CellPart{w.Text("pending")}},
									{Parts: []w.CellPart{w.Text("failed — 550 mailbox unavailable")}, Mono: true, Tone: w.CellDanger},
									{Parts: []w.CellPart{w.Text("2026-08-17 09:30")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Label: "Resend", Small: true},
											w.Button{Label: "Copy link", Small: true, Copy: "https://carrel.example/invite/a41e", Copied: "Copied"},
											w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
										},
										Acts: true,
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

// Admin2 is frame «Администрирование · пользователи» (s-admin-2).
func Admin2() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Administration"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "adm", Label: "Settings", Href: "#"},
								{Icon: "shield", Label: "Installation", Href: "#"},
								{Icon: "search", Label: "DAV validator", Href: "#"},
								{Icon: "dup", Label: "Key escrow", Href: "#"},
								{Icon: "clock", Label: "Audit log", Href: "#"},
							},
							On: 2,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Tone:   w.RailStatusAlert,
							Text:   "SMTP not configured",
							Detail: w.Dim{Parts: []w.Inline{w.Text("No mail is going out.")}, Size: w.Dim11},
						},
						w.RailStatus{
							Text:   "Checks passed",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last check "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Users"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("6 accounts · 2 administrators · 4 covered by escrow")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Add user", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Reset a password"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter by login", Label: "Filter by login"},
								Waiting: true,
								Variant: w.FilterH26W200,
							},
							w.Seg{Options: []w.SegOption{{Label: "All", On: true}, {Label: "Administrators"}, {Label: "Disabled"}}},
							w.Button{Folds: true, Label: "Columns", Icon: "cols", Title: "Columns"},
						},
					},
					w.Table{
						Cols: []w.TableCol{
							{Label: "Login"},
							{Label: "Role"},
							{Label: "Created", Rank: w.Rank3, Secondary: true},
							{Label: "Last login", Rank: w.Rank4},
							{Label: "DAV"},
							{Label: "Sessions"},
							{Label: "Escrow"},
							{},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("mix")}, Name: true},
									{Parts: []w.CellPart{w.Text("Administrator")}},
									{Parts: []w.CellPart{w.Text("2026-02-11")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-15 09:20")}, Mono: true},
									{Parts: []w.CellPart{w.Text("3")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2")}, Mono: true},
									{Parts: []w.CellPart{w.Text("—")}},
									{
										Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
										Acts:  true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("anna")}},
									{Parts: []w.CellPart{w.Text("User")}},
									{Parts: []w.CellPart{w.Text("2026-03-02")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-14 18:02")}, Mono: true},
									{Parts: []w.CellPart{w.Text("1")}, Mono: true},
									{Parts: []w.CellPart{w.Text("1")}, Mono: true},
									{Parts: []w.CellPart{w.Badge{Label: "deposited"}}},
									{
										Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
										Acts:  true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("viktor "), w.Badge{Label: "disabled"}}},
									{Parts: []w.CellPart{w.Text("User")}},
									{Parts: []w.CellPart{w.Text("2026-04-19")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-06-30 11:41")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2")}, Mono: true},
									{Parts: []w.CellPart{w.Text("0")}, Mono: true},
									{Parts: []w.CellPart{w.Badge{Label: "deposited"}}},
									{
										Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
										Acts:  true,
									},
								},
								State: w.RowGone,
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("olga")}},
									{Parts: []w.CellPart{w.Text("Administrator")}},
									{Parts: []w.CellPart{w.Text("2026-05-11")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-16 08:55")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2")}, Mono: true},
									{Parts: []w.CellPart{w.Text("1")}, Mono: true},
									{Parts: []w.CellPart{w.Badge{Label: "deposited"}}},
									{
										Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
										Acts:  true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("petr")}},
									{Parts: []w.CellPart{w.Text("User")}},
									{Parts: []w.CellPart{w.Text("2026-06-02")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-12 21:14")}, Mono: true},
									{Parts: []w.CellPart{w.Text("1")}, Mono: true},
									{Parts: []w.CellPart{w.Text("1")}, Mono: true},
									{Parts: []w.CellPart{w.Badge{Label: "deposited"}}},
									{
										Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
										Acts:  true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("ivan")}},
									{Parts: []w.CellPart{w.Text("User")}},
									{Parts: []w.CellPart{w.Text("2026-07-20")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-15 10:30")}, Mono: true},
									{Parts: []w.CellPart{w.Text("0")}, Mono: true},
									{Parts: []w.CellPart{w.Text("0")}, Mono: true},
									{Parts: []w.CellPart{w.Text("—")}},
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
	}
}

// Admin24 is frame «Сброс пароля · лист поверх Users» (s-admin-24).
func Admin24() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
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
							w.Rubric{Text: "Administration"},
							w.Nav{
								Items: []w.NavItem{
									{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
									{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
									{Icon: "adm", Label: "Settings", Href: "#"},
									{Icon: "shield", Label: "Installation", Href: "#"},
									{Icon: "search", Label: "DAV validator", Href: "#"},
									{Icon: "dup", Label: "Key escrow", Href: "#"},
									{Icon: "clock", Label: "Audit log", Href: "#"},
								},
								On: 2,
							},
						},
					},
					w.RailFoot{
						Items: []w.RailFootPart{
							w.RailStatus{
								Tone:   w.RailStatusAlert,
								Text:   "SMTP not configured",
								Detail: w.Dim{Parts: []w.Inline{w.Text("No mail is going out.")}, Size: w.Dim11},
							},
							w.RailStatus{
								Text:   "Checks passed",
								Detail: w.Dim{Parts: []w.Inline{w.Text("Last check "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
							},
						},
						About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
					},
				},
				Main: w.Main{
					Items: []w.MainPart{
						w.Head{
							Title: &w.HeadTitle{Text: "Users"},
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("6 accounts · 2 administrators · 4 covered by escrow")}}},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "Add user", Icon: "plus", Kind: w.ButtonPrimary},
									w.Button{Label: "Reset a password"},
								},
							},
						},
						w.Bar{
							Items: []w.BarPart{
								w.Filter{
									Query:   w.Query{Placeholder: "Filter by login", Label: "Filter by login"},
									Waiting: true,
									Variant: w.FilterH26W200,
								},
								w.Seg{Options: []w.SegOption{{Label: "All", On: true}, {Label: "Administrators"}, {Label: "Disabled"}}},
								w.Button{Folds: true, Label: "Columns", Icon: "cols", Title: "Columns"},
							},
						},
						w.Table{
							Cols: []w.TableCol{
								{Label: "Login"},
								{Label: "Role"},
								{Label: "Created", Rank: w.Rank3, Secondary: true},
								{Label: "Last login", Rank: w.Rank4},
								{Label: "DAV"},
								{Label: "Sessions"},
								{Label: "Escrow"},
								{},
							},
							Rows: []w.TableRow{
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("mix")}, Name: true},
										{Parts: []w.CellPart{w.Text("Administrator")}},
										{Parts: []w.CellPart{w.Text("2026-02-11")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-15 09:20")}, Mono: true},
										{Parts: []w.CellPart{w.Text("3")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2")}, Mono: true},
										{Parts: []w.CellPart{w.Text("—")}},
										{
											Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
											Acts:  true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("anna")}},
										{Parts: []w.CellPart{w.Text("User")}},
										{Parts: []w.CellPart{w.Text("2026-03-02")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-14 18:02")}, Mono: true},
										{Parts: []w.CellPart{w.Text("1")}, Mono: true},
										{Parts: []w.CellPart{w.Text("1")}, Mono: true},
										{Parts: []w.CellPart{w.Badge{Label: "deposited"}}},
										{
											Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
											Acts:  true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("viktor "), w.Badge{Label: "disabled"}}},
										{Parts: []w.CellPart{w.Text("User")}},
										{Parts: []w.CellPart{w.Text("2026-04-19")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-06-30 11:41")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2")}, Mono: true},
										{Parts: []w.CellPart{w.Text("0")}, Mono: true},
										{Parts: []w.CellPart{w.Badge{Label: "deposited"}}},
										{
											Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
											Acts:  true,
										},
									},
									State: w.RowGone,
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("olga")}},
										{Parts: []w.CellPart{w.Text("Administrator")}},
										{Parts: []w.CellPart{w.Text("2026-05-11")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-16 08:55")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2")}, Mono: true},
										{Parts: []w.CellPart{w.Text("1")}, Mono: true},
										{Parts: []w.CellPart{w.Badge{Label: "deposited"}}},
										{
											Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
											Acts:  true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("petr")}},
										{Parts: []w.CellPart{w.Text("User")}},
										{Parts: []w.CellPart{w.Text("2026-06-02")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-12 21:14")}, Mono: true},
										{Parts: []w.CellPart{w.Text("1")}, Mono: true},
										{Parts: []w.CellPart{w.Text("1")}, Mono: true},
										{Parts: []w.CellPart{w.Badge{Label: "deposited"}}},
										{
											Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
											Acts:  true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("ivan")}},
										{Parts: []w.CellPart{w.Text("User")}},
										{Parts: []w.CellPart{w.Text("2026-07-20")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-15 10:30")}, Mono: true},
										{Parts: []w.CellPart{w.Text("0")}, Mono: true},
										{Parts: []w.CellPart{w.Text("0")}, Mono: true},
										{Parts: []w.CellPart{w.Text("—")}},
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
				Height: w.ShellHeight560,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/admin/users/reset",
					Icon:   "shield",
					Tone:   w.SheetIconDanger,
					Title:  "Reset a password",
					Body: []w.SheetPart{
						w.Field{
							Label: "Whose",
							Parts: []w.FieldPart{
								w.Pick{
									Name:    "login",
									Options: []w.PickOption{{Value: "anna.user", Label: "anna · User", On: true}},
									Label:   "Whose",
								},
							},
							Width: w.FieldW100,
						},
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("The passwords this person saved for DAV servers are sealed with their own password, and a reset destroys them: every account they connected has to be connected again, and their published links, devices and external backup triggers stop. Their records on the servers stay as they are."),
							},
							Tone:  w.MsgAlert,
							Space: w.MsgFlush,
						},
						w.Field{
							Label: "Temporary password",
							Parts: []w.FieldPart{
								w.Input{Name: "password", Value: "4c1-tide-oven", Label: "Temporary password", Mono: true},
								w.FieldHint{
									Parts: []w.Inline{
										w.Text("Hand it over yourself; they choose their own at the next sign-in. If their key is deposited, recover the account in Key escrow instead: nothing is destroyed then."),
									},
								},
							},
							Width: w.FieldW100,
						},
						w.Field{
							Label: "Type the login to confirm",
							Parts: []w.FieldPart{w.Input{Name: "confirm", Placeholder: "anna", Label: "Type the login to confirm", Focus: true}},
							Width: w.FieldW100,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Reset password", Kind: w.ButtonDangerSolid, Submit: true},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/admin/"}, w.Kbd{Keys: "Esc"}},
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

// Admin34 is frame «Удаление учётной записи · лист поверх Users» (s-admin-34).
func Admin34() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
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
							w.Rubric{Text: "Administration"},
							w.Nav{
								Items: []w.NavItem{
									{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
									{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
									{Icon: "adm", Label: "Settings", Href: "#"},
									{Icon: "shield", Label: "Installation", Href: "#"},
									{Icon: "search", Label: "DAV validator", Href: "#"},
									{Icon: "dup", Label: "Key escrow", Href: "#"},
									{Icon: "clock", Label: "Audit log", Href: "#"},
								},
								On: 2,
							},
						},
					},
					w.RailFoot{
						Items: []w.RailFootPart{
							w.RailStatus{
								Tone:   w.RailStatusAlert,
								Text:   "SMTP not configured",
								Detail: w.Dim{Parts: []w.Inline{w.Text("No mail is going out.")}, Size: w.Dim11},
							},
							w.RailStatus{
								Text:   "Checks passed",
								Detail: w.Dim{Parts: []w.Inline{w.Text("Last check "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
							},
						},
						About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
					},
				},
				Main: w.Main{
					Items: []w.MainPart{
						w.Head{
							Title: &w.HeadTitle{Text: "Users"},
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("6 accounts · 2 administrators · 4 covered by escrow")}}},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "Add user", Icon: "plus", Kind: w.ButtonPrimary},
									w.Button{Label: "Reset a password"},
								},
							},
						},
						w.Bar{
							Items: []w.BarPart{
								w.Filter{
									Query:   w.Query{Placeholder: "Filter by login", Label: "Filter by login"},
									Waiting: true,
									Variant: w.FilterH26W200,
								},
								w.Seg{Options: []w.SegOption{{Label: "All", On: true}, {Label: "Administrators"}, {Label: "Disabled"}}},
								w.Button{Folds: true, Label: "Columns", Icon: "cols", Title: "Columns"},
							},
						},
						w.Table{
							Cols: []w.TableCol{
								{Label: "Login"},
								{Label: "Role"},
								{Label: "Created", Rank: w.Rank3, Secondary: true},
								{Label: "Last login", Rank: w.Rank4},
								{Label: "DAV"},
								{Label: "Sessions"},
								{Label: "Escrow"},
								{},
							},
							Rows: []w.TableRow{
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("mix")}, Name: true},
										{Parts: []w.CellPart{w.Text("Administrator")}},
										{Parts: []w.CellPart{w.Text("2026-02-11")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-15 09:20")}, Mono: true},
										{Parts: []w.CellPart{w.Text("3")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2")}, Mono: true},
										{Parts: []w.CellPart{w.Text("—")}},
										{
											Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
											Acts:  true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("anna")}},
										{Parts: []w.CellPart{w.Text("User")}},
										{Parts: []w.CellPart{w.Text("2026-03-02")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-14 18:02")}, Mono: true},
										{Parts: []w.CellPart{w.Text("1")}, Mono: true},
										{Parts: []w.CellPart{w.Text("1")}, Mono: true},
										{Parts: []w.CellPart{w.Badge{Label: "deposited"}}},
										{
											Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
											Acts:  true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("viktor "), w.Badge{Label: "disabled"}}},
										{Parts: []w.CellPart{w.Text("User")}},
										{Parts: []w.CellPart{w.Text("2026-04-19")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-06-30 11:41")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2")}, Mono: true},
										{Parts: []w.CellPart{w.Text("0")}, Mono: true},
										{Parts: []w.CellPart{w.Badge{Label: "deposited"}}},
										{
											Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
											Acts:  true,
										},
									},
									State: w.RowGone,
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("olga")}},
										{Parts: []w.CellPart{w.Text("Administrator")}},
										{Parts: []w.CellPart{w.Text("2026-05-11")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-16 08:55")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2")}, Mono: true},
										{Parts: []w.CellPart{w.Text("1")}, Mono: true},
										{Parts: []w.CellPart{w.Badge{Label: "deposited"}}},
										{
											Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
											Acts:  true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("petr")}},
										{Parts: []w.CellPart{w.Text("User")}},
										{Parts: []w.CellPart{w.Text("2026-06-02")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-12 21:14")}, Mono: true},
										{Parts: []w.CellPart{w.Text("1")}, Mono: true},
										{Parts: []w.CellPart{w.Text("1")}, Mono: true},
										{Parts: []w.CellPart{w.Badge{Label: "deposited"}}},
										{
											Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}},
											Acts:  true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("ivan")}},
										{Parts: []w.CellPart{w.Text("User")}},
										{Parts: []w.CellPart{w.Text("2026-07-20")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-15 10:30")}, Mono: true},
										{Parts: []w.CellPart{w.Text("0")}, Mono: true},
										{Parts: []w.CellPart{w.Text("0")}, Mono: true},
										{Parts: []w.CellPart{w.Text("—")}},
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
				Height: w.ShellHeight560,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/admin/",
					Fields: []w.Hidden{
						{Name: "csrf", Value: "7f3a91de"},
						{Name: "action", Value: "delete"},
						{Name: "target", Value: "u-anna"},
					},
					Icon:  "trash",
					Tone:  w.SheetIconDanger,
					Title: "Delete anna",
					Body: []w.SheetPart{
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("Carrel forgets this account: its connections and saved passwords, published links, devices, backups and the history of its password. Its records on the servers stay as they are."),
							},
							Tone:  w.MsgAlert,
							Space: w.MsgFlush,
						},
						w.Field{
							Label: "Type the login to confirm",
							Parts: []w.FieldPart{
								w.Input{Name: "login", Placeholder: "anna", Label: "Type the login to confirm", Focus: true, Required: true},
							},
							Width: w.FieldW100,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Delete account", Kind: w.ButtonDangerSolid, Submit: true},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/admin/"}, w.Kbd{Keys: "Esc"}},
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

// Admin3 is frame «Администрирование · настройки» (s-admin-3).
func Admin3() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Administration"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "adm", Label: "Settings", Href: "#"},
								{Icon: "shield", Label: "Installation", Href: "#"},
								{Icon: "search", Label: "DAV validator", Href: "#"},
								{Icon: "dup", Label: "Key escrow", Href: "#"},
								{Icon: "clock", Label: "Audit log", Href: "#"},
							},
							On: 3,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Tone:   w.RailStatusAlert,
							Text:   "SMTP not configured",
							Detail: w.Dim{Parts: []w.Inline{w.Text("No mail is going out.")}, Size: w.Dim11},
						},
						w.RailStatus{
							Text:   "Checks passed",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last check "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Settings"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Global settings and mail")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Discard", Href: "/admin/settings"},
								w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "admin-settings"},
							},
						},
					},
					w.Form{
						Items: []w.FormPart{
							w.FieldSet{
								Rubric: w.Rubric{Text: "Global"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Invite lifetime, h",
										Parts: []w.FieldPart{w.Input{Name: "inviteTTL", Value: "72", Label: "Invite lifetime, h", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Idle timeout, h",
										Parts: []w.FieldPart{w.Input{Name: "idle", Value: "12", Label: "Idle timeout, h", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Session life, d",
										Parts: []w.FieldPart{w.Input{Name: "session", Value: "30", Label: "Session life, d", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Instance name",
										Parts: []w.FieldPart{w.Input{Name: "name", Value: "Carrel", Label: "Instance name"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Public address",
										Parts: []w.FieldPart{
											w.Input{
												Name:     "publicURL",
												Value:    "https://carrel.example",
												Label:    "Public address",
												Mono:     true,
												ReadOnly: true,
											},
											w.FieldHint{Parts: []w.Inline{w.Text("Links in letters start with it. Set with CARREL_PUBLIC_URL.")}},
										},
										Width: w.FieldLG,
									},
								},
								Variant: w.FieldSetLead,
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Mail", Items: []w.RubricPart{w.Badge{Label: "not configured", Tone: w.BadgeAlert}}},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Host",
										Parts: []w.FieldPart{w.Input{Name: "host", Placeholder: "smtp.example.org", Label: "Host", Mono: true}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Port",
										Parts: []w.FieldPart{w.Input{Name: "port", Value: "587", Label: "Port", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Encryption",
										Parts: []w.FieldPart{
											w.Pick{
												Name:    "security",
												Options: []w.PickOption{{Value: "starttls", Label: "STARTTLS", On: true}},
												Label:   "Encryption",
											},
										},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Username",
										Parts: []w.FieldPart{w.Input{Name: "username", Placeholder: "—", Label: "Username"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Password",
										Parts: []w.FieldPart{
											w.Input{
												Kind:        w.InputPassword,
												Name:        "password",
												Placeholder: "leave blank to keep the current one",
												Label:       "Password",
											},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "From address",
										Parts: []w.FieldPart{w.Input{Name: "from", Placeholder: "carrel@example.org", Label: "From address"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "From name",
										Parts: []w.FieldPart{w.Input{Name: "fromName", Placeholder: "Carrel", Label: "From name"}},
										Width: w.FieldMD,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Letters from users", Items: []w.RubricPart{w.Badge{Label: "off"}}},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Check{
												Name:   "people",
												Value:  "on",
												Reason: "Mail is not configured: there is nothing to send through yet.",
												Parts:  []w.CheckPart{w.Text("Let users send letters and links from Carrel")},
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Everyone signed in may send mail through this relay, from this domain. With self-registration on, that is an open sending form."),
												},
											},
										},
										OwnLine: true,
										Place:   w.Grown,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Test"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Send test message to",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Input{Name: "to", Placeholder: "you@example.org", Label: "Send test message to", Variant: w.InputFill},
													w.Button{Label: "Send test", To: "/admin/settings/mail-test", Submit: true},
												},
												Gap: w.LineGap8,
											},
										},
										Width: w.FieldLG,
									},
								},
							},
						},
						ID:     "admin-settings",
						Action: "/admin/settings",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
				},
			},
		},
	}
}

// Admin32 is frame «Настройки · почта настроена, проба ушла» (s-admin-32).
func Admin32() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Administration"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "adm", Label: "Settings", Href: "#"},
								{Icon: "shield", Label: "Installation", Href: "#"},
								{Icon: "search", Label: "DAV validator", Href: "#"},
								{Icon: "dup", Label: "Key escrow", Href: "#"},
								{Icon: "clock", Label: "Audit log", Href: "#"},
							},
							On: 3,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Text:   "Mail is going out",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last sent "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
						w.RailStatus{
							Text:   "Checks passed",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last check "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Settings"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Global settings and mail")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Discard", Href: "/admin/settings"},
								w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "admin-settings"},
							},
						},
					},
					w.Form{
						Items: []w.FormPart{
							w.FieldSet{
								Rubric: w.Rubric{Text: "Global"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Invite lifetime, h",
										Parts: []w.FieldPart{w.Input{Name: "inviteTTL", Value: "72", Label: "Invite lifetime, h", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Idle timeout, h",
										Parts: []w.FieldPart{w.Input{Name: "idle", Value: "12", Label: "Idle timeout, h", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Session life, d",
										Parts: []w.FieldPart{w.Input{Name: "session", Value: "30", Label: "Session life, d", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Instance name",
										Parts: []w.FieldPart{w.Input{Name: "name", Value: "Carrel", Label: "Instance name"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Public address",
										Parts: []w.FieldPart{
											w.Input{
												Name:     "publicURL",
												Value:    "https://carrel.example",
												Label:    "Public address",
												Mono:     true,
												ReadOnly: true,
											},
											w.FieldHint{Parts: []w.Inline{w.Text("Links in letters start with it. Set with CARREL_PUBLIC_URL.")}},
										},
										Width: w.FieldLG,
									},
								},
								Variant: w.FieldSetLead,
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Mail"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Host",
										Parts: []w.FieldPart{w.Input{Name: "host", Value: "smtp.slovo.example", Label: "Host", Mono: true}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Port",
										Parts: []w.FieldPart{w.Input{Name: "port", Value: "587", Label: "Port", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Encryption",
										Parts: []w.FieldPart{
											w.Pick{
												Name:    "security",
												Options: []w.PickOption{{Value: "starttls", Label: "STARTTLS", On: true}},
												Label:   "Encryption",
											},
										},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Username",
										Parts: []w.FieldPart{w.Input{Name: "username", Value: "carrel", Label: "Username"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Password",
										Parts: []w.FieldPart{
											w.Input{
												Kind:        w.InputPassword,
												Name:        "password",
												Placeholder: "leave blank to keep the current one",
												Label:       "Password",
											},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "From address",
										Parts: []w.FieldPart{w.Input{Name: "from", Value: "carrel@slovo.example", Label: "From address"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "From name",
										Parts: []w.FieldPart{w.Input{Name: "fromName", Placeholder: "Carrel", Label: "From name"}},
										Width: w.FieldMD,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Letters from users", Items: []w.RubricPart{w.Badge{Label: "on"}}},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Check{
												On:    true,
												Name:  "people",
												Value: "on",
												Parts: []w.CheckPart{w.Text("Let users send letters and links from Carrel")},
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Everyone signed in may send mail through this relay, from this domain. With self-registration on, that is an open sending form."),
												},
											},
										},
										OwnLine: true,
										Place:   w.Grown,
									},
									w.Field{
										Label: "Sender address",
										Parts: []w.FieldPart{
											w.Seg{
												Options: []w.SegOption{
													{Label: "carrel@slovo.example", On: true, Value: "noreply"},
													{Label: "login@slovo.example", Value: "login"},
												},
												Variant: w.SegStart,
												Name:    "sender",
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Your relay decides which of the two it will let out. The envelope sender stays "),
													w.Mono{Text: "carrel@slovo.example"},
													w.Text(" either way, so SPF is unaffected; DKIM and DMARC for this domain are set up on the relay, not here."),
												},
											},
										},
										Place: w.Grown,
									},
									w.Field{
										Label: "Attachment, MB",
										Parts: []w.FieldPart{w.Input{Name: "attach", Value: "2", Label: "Attachment, MB", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Letter, MB",
										Parts: []w.FieldPart{w.Input{Name: "letter", Value: "8", Label: "Letter, MB", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Recipients",
										Parts: []w.FieldPart{w.Input{Name: "recipients", Value: "10", Label: "Recipients", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Letters a day",
										Parts: []w.FieldPart{w.Input{Name: "perDay", Value: "50", Label: "Letters a day", Mono: true}},
										Width: w.FieldXS,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Test"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Send test message to",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Input{Name: "to", Value: "mix@slovo.example", Label: "Send test message to", Variant: w.InputFill},
													w.Button{Label: "Send test", To: "/admin/settings/mail-test", Submit: true},
												},
												Gap: w.LineGap8,
											},
										},
										Width: w.FieldLG,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.Msg{
												Parts: []w.MsgPart{
													w.Strong("Sent."),
													w.Text(" The relay answered "),
													w.Mono{Text: "250 2.0.0 Ok: queued as 4FB1C2"},
													w.Text("."),
												},
												Space: w.MsgFlush,
											},
										},
										Width: w.FieldFull,
									},
								},
							},
						},
						ID:     "admin-settings",
						Action: "/admin/settings",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
				},
			},
		},
	}
}

// Admin33 is frame «Настройки · отказ поля» (s-admin-33).
func Admin33() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Administration"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "adm", Label: "Settings", Href: "#"},
								{Icon: "shield", Label: "Installation", Href: "#"},
								{Icon: "search", Label: "DAV validator", Href: "#"},
								{Icon: "dup", Label: "Key escrow", Href: "#"},
								{Icon: "clock", Label: "Audit log", Href: "#"},
							},
							On: 3,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Text:   "Mail is going out",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last sent "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
						w.RailStatus{
							Text:   "Checks passed",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last check "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Settings"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Global settings and mail")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Discard", Href: "/admin/settings"},
								w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "admin-settings"},
							},
						},
					},
					w.Form{
						Items: []w.FormPart{
							w.FieldSet{
								Rubric: w.Rubric{Text: "Global"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Invite lifetime, h",
										Parts: []w.FieldPart{w.Input{Name: "inviteTTL", Value: "72", Label: "Invite lifetime, h", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Idle timeout, h",
										Parts: []w.FieldPart{w.Input{Name: "idle", Value: "12", Label: "Idle timeout, h", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Session life, d",
										Parts: []w.FieldPart{w.Input{Name: "session", Value: "30", Label: "Session life, d", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Instance name",
										Parts: []w.FieldPart{w.Input{Name: "name", Value: "Carrel", Label: "Instance name"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Public address",
										Parts: []w.FieldPart{
											w.Input{
												Name:     "publicURL",
												Value:    "https://carrel.example",
												Label:    "Public address",
												Mono:     true,
												ReadOnly: true,
											},
											w.FieldHint{Parts: []w.Inline{w.Text("Links in letters start with it. Set with CARREL_PUBLIC_URL.")}},
										},
										Width: w.FieldLG,
									},
								},
								Variant: w.FieldSetLead,
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Mail"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Host",
										Parts: []w.FieldPart{w.Input{Name: "host", Value: "smtp.slovo.example", Label: "Host", Mono: true}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Port",
										Parts: []w.FieldPart{
											w.Input{Name: "port", Value: "70000", Label: "Port", Mono: true, Error: true},
											w.FieldHint{Parts: []w.Inline{w.Text("A port from 1 to 65535.")}, Variant: w.HintDanger},
										},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Encryption",
										Parts: []w.FieldPart{
											w.Pick{
												Name:    "security",
												Options: []w.PickOption{{Value: "starttls", Label: "STARTTLS", On: true}},
												Label:   "Encryption",
											},
										},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Username",
										Parts: []w.FieldPart{w.Input{Name: "username", Value: "carrel", Label: "Username"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Password",
										Parts: []w.FieldPart{
											w.Input{
												Kind:        w.InputPassword,
												Name:        "password",
												Placeholder: "leave blank to keep the current one",
												Label:       "Password",
											},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "From address",
										Parts: []w.FieldPart{w.Input{Name: "from", Value: "carrel@slovo.example", Label: "From address"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "From name",
										Parts: []w.FieldPart{w.Input{Name: "fromName", Placeholder: "Carrel", Label: "From name"}},
										Width: w.FieldMD,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Letters from users", Items: []w.RubricPart{w.Badge{Label: "off"}}},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Check{
												Name:  "people",
												Value: "on",
												Parts: []w.CheckPart{w.Text("Let users send letters and links from Carrel")},
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Everyone signed in may send mail through this relay, from this domain. With self-registration on, that is an open sending form."),
												},
											},
										},
										OwnLine: true,
										Place:   w.Grown,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Test"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Send test message to",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Input{Name: "to", Value: "mix@slovo.example", Label: "Send test message to", Variant: w.InputFill},
													w.Button{Label: "Send test", To: "/admin/settings/mail-test", Submit: true},
												},
												Gap: w.LineGap8,
											},
										},
										Width: w.FieldLG,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.Msg{
												Parts: []w.MsgPart{w.Text("The relay as typed cannot be used: see the marked fields.")},
												Tone:  w.MsgAlert,
												Space: w.MsgFlush,
											},
										},
										Width: w.FieldFull,
									},
								},
							},
						},
						ID:     "admin-settings",
						Action: "/admin/settings",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
				},
			},
		},
	}
}

// Admin4 is frame «Администрирование · проверка установки» (s-admin-4).
func Admin4() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Administration"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "adm", Label: "Settings", Href: "#"},
								{Icon: "shield", Label: "Installation", Href: "#"},
								{Icon: "search", Label: "DAV validator", Href: "#"},
								{Icon: "shield", Label: "Key escrow", Href: "#"},
								{Icon: "clock", Label: "Audit log", Href: "#"},
							},
							On: 4,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Text:   "Mail is going out",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last sent "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
						w.RailStatus{
							Tone:   w.RailStatusAlert,
							Text:   "2 checks failed",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last check "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Installation check"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{w.Text("Carrel asks itself from outside, at the address you declared as the base one")},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Field{
									Parts: []w.FieldPart{
										w.Input{
											Name:    "url",
											Value:   "https://carrel.example",
											Label:   "Base address",
											Mono:    true,
											Of:      "install",
											Variant: w.Input11_5,
										},
									},
									Width: w.FieldW280,
								},
								w.Button{Label: "Run the check", Kind: w.ButtonPrimary, Form: "install"},
							},
						},
					},
					w.Runbar{
						Meter: w.Meter{Percent: 100},
						Parts: []w.RunbarPart{
							w.Text("11 checks · "),
							w.Span{Face: w.FaceAccent, Text: "8 fine"},
							w.Text(" · "),
							w.Span{Face: w.FaceDanger, Text: "2 broken"},
							w.Text(" · 1 not measured"),
							w.Mono{Text: "4.2 s", Place: w.Pushed},
						},
					},
					w.Table{
						Cols: []w.TableCol{
							{Width: w.ColMenu},
							{Label: "Check", Width: w.Col26},
							{Label: "What happened", Rank: w.Rank3},
							{Label: "What to do", Width: w.Col26, Rank: w.Rank4},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Tick{Done: true, Label: "Done"}}},
									{Parts: []w.CellPart{w.Text("Reachable from outside")}},
									{Parts: []w.CellPart{w.Text("GET https://carrel.example/healthz → 200 · 41 ms")}, Mono: true},
									{},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Tick{Done: true, Label: "Done"}}},
									{Parts: []w.CellPart{w.Text("Certificate")}},
									{
										Parts: []w.CellPart{w.Text("Let’s Encrypt R11 · until 2026-11-02 · chain complete · name matches")},
										Mono:  true,
									},
									{},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Tick{Done: true, Label: "Done"}}},
									{Parts: []w.CellPart{w.Text("Base address")}},
									{Parts: []w.CellPart{w.Text("Host and X-Forwarded-Host match base_url")}, Mono: true},
									{},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Tick{Done: true, Label: "Done"}}},
									{Parts: []w.CellPart{w.Text("Scheme behind the proxy")}},
									{Parts: []w.CellPart{w.Text("X-Forwarded-Proto: https · links are built as https")}, Mono: true},
									{},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Tick{Done: true, Label: "Done"}}},
									{Parts: []w.CellPart{w.Text("Client address")}},
									{
										Parts: []w.CellPart{w.Text("X-Forwarded-For: 203.0.113.7 · limits count this, not the proxy")},
										Mono:  true,
									},
									{},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Tick{Label: "Done", Variant: w.TickEdgeDanger}}},
									{Parts: []w.CellPart{w.Text("Event stream (SSE)")}, Tone: w.CellDanger},
									{
										Parts: []w.CellPart{w.Text("first event after 30.0 s instead of 0.2 s — the response is buffered")},
										Mono:  true,
									},
									{
										Parts: []w.CellPart{w.Text("nginx: proxy_buffering off for /app/find/*/stream")},
										Mono:  true,
										Small: true,
									},
								},
								State: w.RowPicked,
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Tick{Label: "Done", Variant: w.TickEdgeDanger}}},
									{Parts: []w.CellPart{w.Text("Upload size")}, Tone: w.CellDanger},
									{
										Parts: []w.CellPart{w.Text("20 MB probe → 413 from the proxy; Carrel streams and has no size ceiling of its own")},
										Mono:  true,
									},
									{
										Parts: []w.CellPart{w.Text("client_max_body_size — the proxy is cutting the body")},
										Mono:  true,
										Small: true,
									},
								},
								State: w.RowPicked,
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Tick{Label: "Done", Variant: w.TickEdgeDusty}}},
									{Parts: []w.CellPart{w.Text("Long request timeout")}, Tone: w.CellMuted},
									{
										Parts: []w.CellPart{w.Text("seven sources polled in 21 s; the proxy ceiling cannot be measured from outside")},
										Mono:  true,
									},
									{Parts: []w.CellPart{w.Text("proxy_read_timeout ≥ 60s")}, Mono: true, Small: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Tick{Done: true, Label: "Done"}}},
									{Parts: []w.CellPart{w.Text("Cookie")}},
									{Parts: []w.CellPart{w.Text("Secure, HttpOnly, SameSite=Lax · came back")}, Mono: true},
									{},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Tick{Done: true, Label: "Done"}}},
									{Parts: []w.CellPart{w.Text("Clock")}},
									{Parts: []w.CellPart{w.Text("0.4 s from NTP · TZ Europe/Moscow")}, Mono: true},
									{},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Tick{Done: true, Label: "Done"}}},
									{Parts: []w.CellPart{w.Text("Data folder")}},
									{Parts: []w.CellPart{w.Text("/var/lib/carrel · writable · 41 GB free")}, Mono: true},
									{},
								},
							},
						},
					},
					w.PostForm{ID: "install", Action: "/admin/install", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("The check runs from outside inwards: Carrel calls itself at the base address, so it sees exactly what a browser or DAVx5 will. Nothing is changed and nothing is stored, apart from the audit line saying the check was run."),
						},
					},
				},
			},
		},
	}
}

// Admin25 is frame «Проверка установки · ещё не запускалась» (s-admin-25).
func Admin25() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Administration"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "adm", Label: "Settings", Href: "#"},
								{Icon: "shield", Label: "Installation", Href: "#"},
								{Icon: "search", Label: "DAV validator", Href: "#"},
								{Icon: "shield", Label: "Key escrow", Href: "#"},
								{Icon: "clock", Label: "Audit log", Href: "#"},
							},
							On: 4,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Text:   "Mail is going out",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last sent "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
						w.RailStatus{
							Tone:   w.RailStatusOff,
							Text:   "Not checked yet",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Run the check once Carrel is behind its proxy.")}, Size: w.Dim11},
						},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Installation check"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{w.Text("Carrel asks itself from outside, at the address you declared as the base one")},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Field{
									Parts: []w.FieldPart{
										w.Input{
											Name:    "url",
											Value:   "https://carrel.example",
											Label:   "Base address",
											Mono:    true,
											Of:      "install",
											Variant: w.Input11_5,
										},
									},
									Width: w.FieldW280,
								},
								w.Button{Label: "Run the check", Kind: w.ButtonPrimary, Form: "install"},
							},
						},
					},
					w.PostForm{ID: "install", Action: "/admin/install", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("The check runs from outside inwards: Carrel calls itself at the base address, so it sees exactly what a browser or DAVx5 will. Nothing is changed and nothing is stored, apart from the audit line saying the check was run."),
						},
					},
				},
			},
		},
	}
}

// Admin26 is frame «Проверка установки · негодный адрес» (s-admin-26).
func Admin26() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Administration"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "adm", Label: "Settings", Href: "#"},
								{Icon: "shield", Label: "Installation", Href: "#"},
								{Icon: "search", Label: "DAV validator", Href: "#"},
								{Icon: "shield", Label: "Key escrow", Href: "#"},
								{Icon: "clock", Label: "Audit log", Href: "#"},
							},
							On: 4,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Text:   "Mail is going out",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last sent "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
						w.RailStatus{
							Tone:   w.RailStatusAlert,
							Text:   "2 checks failed",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last check "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Installation check"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{w.Text("Carrel asks itself from outside, at the address you declared as the base one")},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Field{
									Parts: []w.FieldPart{
										w.Input{
											Name:    "url",
											Value:   "carrel.example",
											Label:   "Base address",
											Mono:    true,
											Error:   true,
											Of:      "install",
											Variant: w.Input11_5,
										},
									},
									Width: w.FieldW280,
								},
								w.Button{Label: "Run the check", Kind: w.ButtonPrimary, Form: "install"},
							},
						},
					},
					w.Msg{
						Parts: []w.MsgPart{w.Text("Give the address people open Carrel at, like https://carrel.example.org.")},
						Tone:  w.MsgAlert,
					},
					w.PostForm{ID: "install", Action: "/admin/install", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("The check runs from outside inwards: Carrel calls itself at the base address, so it sees exactly what a browser or DAVx5 will. Nothing is changed and nothing is stored, apart from the audit line saying the check was run."),
						},
					},
				},
			},
		},
	}
}

// Admin5 is frame «Администрирование · проверка сервера DAV» (s-admin-5).
func Admin5() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Administration"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "adm", Label: "Settings", Href: "#"},
								{Icon: "shield", Label: "Installation", Href: "#"},
								{Icon: "search", Label: "DAV validator", Href: "#"},
								{Icon: "dup", Label: "Key escrow", Href: "#"},
								{Icon: "clock", Label: "Audit log", Href: "#"},
							},
							On: 5,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Tone:   w.RailStatusAlert,
							Text:   "SMTP not configured",
							Detail: w.Dim{Parts: []w.Inline{w.Text("No mail is going out.")}, Size: w.Dim11},
						},
						w.RailStatus{
							Text:   "Checks passed",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last check "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "DAV validator"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Discovery without saving credentials")}}},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Server URL",
								Parts: []w.FieldPart{
									w.Input{
										Name:     "server",
										Value:    "https://dav.office.example/dav.php/",
										Label:    "Server URL",
										Mono:     true,
										Required: true,
									},
								},
								Width: w.FieldXL,
							},
							w.Field{
								Label: "Username",
								Parts: []w.FieldPart{w.Input{Name: "login", Value: "m.yenuchenko", Label: "Username"}},
								Width: w.FieldMD,
							},
							w.Field{
								Label: "Password",
								Parts: []w.FieldPart{
									w.Input{Kind: w.InputPassword, Name: "password", Label: "Password"},
									w.FieldHint{Parts: []w.Inline{w.Text("Asked each time, never kept.")}},
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
								Parts: []w.FieldPart{w.Button{Label: "Test discovery", Kind: w.ButtonPrimary, To: "/admin/dav?op=test", Submit: true}},
							},
						},
						ID:     "dav",
						Action: "/admin/dav",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
					w.MainDiag{
						Diag: w.Diag{
							Text: "[well-known]  GET /.well-known/caldav → 301 → /dav.php/\n[principal]   PROPFIND /dav.php/ → 207, principal /dav.php/principals/m.yenuchenko/\n[home-set]    calendar-home-set → /dav.php/calendars/m.yenuchenko/\n[home-set]    addressbook-home-set → /dav.php/addressbooks/m.yenuchenko/\n[collections] 4 calendars, 1 address book, 0 plain folders",
						},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Parts: []w.FieldPart{
									w.Check{
										Name:  "agree",
										Value: "on",
										Of:    "dav",
										Parts: []w.CheckPart{w.Text("Create test objects, read them back and delete them")},
									},
									w.FieldHint{
										Parts: []w.Inline{
											w.Text("The second stage writes to this server: one event, task, contact and file, each deleted after it is read. It is written to the audit log."),
										},
									},
								},
								Width: w.FieldFull,
							},
							w.Field{
								Parts: []w.FieldPart{
									w.Button{Label: "Run the full check", Kind: w.ButtonDanger, Form: "dav", To: "/admin/dav?op=exercise"},
								},
							},
						},
					},
				},
			},
		},
	}
}

// Admin27 is frame «Проверка DAV · до проверки» (s-admin-27).
func Admin27() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Administration"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "adm", Label: "Settings", Href: "#"},
								{Icon: "shield", Label: "Installation", Href: "#"},
								{Icon: "search", Label: "DAV validator", Href: "#"},
								{Icon: "dup", Label: "Key escrow", Href: "#"},
								{Icon: "clock", Label: "Audit log", Href: "#"},
							},
							On: 5,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Tone:   w.RailStatusAlert,
							Text:   "SMTP not configured",
							Detail: w.Dim{Parts: []w.Inline{w.Text("No mail is going out.")}, Size: w.Dim11},
						},
						w.RailStatus{
							Text:   "Checks passed",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last check "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "DAV validator"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Discovery without saving credentials")}}},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Server URL",
								Parts: []w.FieldPart{
									w.Input{
										Name:        "server",
										Placeholder: "https://dav.example.org/",
										Label:       "Server URL",
										Mono:        true,
										Required:    true,
									},
								},
								Width: w.FieldXL,
							},
							w.Field{
								Label: "Username",
								Parts: []w.FieldPart{w.Input{Name: "login", Value: "m.yenuchenko", Label: "Username"}},
								Width: w.FieldMD,
							},
							w.Field{
								Label: "Password",
								Parts: []w.FieldPart{
									w.Input{Kind: w.InputPassword, Name: "password", Label: "Password"},
									w.FieldHint{Parts: []w.Inline{w.Text("Asked each time, never kept.")}},
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
								Parts: []w.FieldPart{w.Button{Label: "Test discovery", Kind: w.ButtonPrimary, To: "/admin/dav?op=test", Submit: true}},
							},
						},
						ID:     "dav",
						Action: "/admin/dav",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
				},
			},
		},
	}
}

// Admin28 is frame «Проверка DAV · сломанный шаг» (s-admin-28).
func Admin28() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Administration"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "adm", Label: "Settings", Href: "#"},
								{Icon: "shield", Label: "Installation", Href: "#"},
								{Icon: "search", Label: "DAV validator", Href: "#"},
								{Icon: "dup", Label: "Key escrow", Href: "#"},
								{Icon: "clock", Label: "Audit log", Href: "#"},
							},
							On: 5,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Tone:   w.RailStatusAlert,
							Text:   "SMTP not configured",
							Detail: w.Dim{Parts: []w.Inline{w.Text("No mail is going out.")}, Size: w.Dim11},
						},
						w.RailStatus{
							Text:   "Checks passed",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last check "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "DAV validator"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Discovery without saving credentials")}}},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Server URL",
								Parts: []w.FieldPart{
									w.Input{
										Name:     "server",
										Value:    "https://dav.office.example/dav.php/",
										Label:    "Server URL",
										Mono:     true,
										Required: true,
									},
								},
								Width: w.FieldXL,
							},
							w.Field{
								Label: "Username",
								Parts: []w.FieldPart{w.Input{Name: "login", Value: "m.yenuchenko", Label: "Username"}},
								Width: w.FieldMD,
							},
							w.Field{
								Label: "Password",
								Parts: []w.FieldPart{
									w.Input{Kind: w.InputPassword, Name: "password", Label: "Password"},
									w.FieldHint{Parts: []w.Inline{w.Text("Asked each time, never kept.")}},
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
								Parts: []w.FieldPart{w.Button{Label: "Test discovery", Kind: w.ButtonPrimary, To: "/admin/dav?op=test", Submit: true}},
							},
						},
						ID:     "dav",
						Action: "/admin/dav",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
					w.Msg{Parts: []w.MsgPart{w.Strong("Discovery stopped at the principal.")}, Tone: w.MsgAlert},
					w.MainDiag{
						Diag: w.Diag{
							Text: "[well-known]  GET /.well-known/caldav → 301 → /dav.php/\n[principal]   PROPFIND /dav.php/ → 401 Unauthorized",
						},
					},
				},
			},
		},
	}
}

// Admin29 is frame «Проверка DAV · второй этап пройден» (s-admin-29).
func Admin29() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Administration"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "adm", Label: "Settings", Href: "#"},
								{Icon: "shield", Label: "Installation", Href: "#"},
								{Icon: "search", Label: "DAV validator", Href: "#"},
								{Icon: "dup", Label: "Key escrow", Href: "#"},
								{Icon: "clock", Label: "Audit log", Href: "#"},
							},
							On: 5,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Tone:   w.RailStatusAlert,
							Text:   "SMTP not configured",
							Detail: w.Dim{Parts: []w.Inline{w.Text("No mail is going out.")}, Size: w.Dim11},
						},
						w.RailStatus{
							Text:   "Checks passed",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last check "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "DAV validator"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Discovery without saving credentials")}}},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Server URL",
								Parts: []w.FieldPart{
									w.Input{
										Name:     "server",
										Value:    "https://dav.office.example/dav.php/",
										Label:    "Server URL",
										Mono:     true,
										Required: true,
									},
								},
								Width: w.FieldXL,
							},
							w.Field{
								Label: "Username",
								Parts: []w.FieldPart{w.Input{Name: "login", Value: "m.yenuchenko", Label: "Username"}},
								Width: w.FieldMD,
							},
							w.Field{
								Label: "Password",
								Parts: []w.FieldPart{
									w.Input{Kind: w.InputPassword, Name: "password", Label: "Password"},
									w.FieldHint{Parts: []w.Inline{w.Text("Asked each time, never kept.")}},
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
								Parts: []w.FieldPart{w.Button{Label: "Test discovery", Kind: w.ButtonPrimary, To: "/admin/dav?op=test", Submit: true}},
							},
						},
						ID:     "dav",
						Action: "/admin/dav",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
					w.MainDiag{
						Diag: w.Diag{
							Text: "[well-known]  GET /.well-known/caldav → 301 → /dav.php/\n[principal]   PROPFIND /dav.php/ → 207, principal /dav.php/principals/m.yenuchenko/\n[home-set]    calendar-home-set → /dav.php/calendars/m.yenuchenko/\n[home-set]    addressbook-home-set → /dav.php/addressbooks/m.yenuchenko/\n[collections] 4 calendars, 1 address book, 0 plain folders",
						},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Parts: []w.FieldPart{
									w.Check{
										Name:  "agree",
										Value: "on",
										Of:    "dav",
										Parts: []w.CheckPart{w.Text("Create test objects, read them back and delete them")},
									},
									w.FieldHint{
										Parts: []w.Inline{
											w.Text("The second stage writes to this server: one event, task, contact and file, each deleted after it is read. It is written to the audit log."),
										},
									},
								},
								Width: w.FieldFull,
							},
							w.Field{
								Parts: []w.FieldPart{
									w.Button{Label: "Run the full check", Kind: w.ButtonDanger, Form: "dav", To: "/admin/dav?op=exercise"},
								},
							},
						},
					},
					w.Table{
						Cols: []w.TableCol{{Width: w.ColMenu}, {Label: "Operation", Width: w.Col26}, {Label: "What happened", Rank: w.Rank3}},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Tick{Done: true, Label: "Done"}}},
									{Parts: []w.CellPart{w.Text("Create, read back, delete · event")}},
									{
										Parts: []w.CellPart{
											w.Text("PUT · GET · DELETE /dav.php/calendars/m.yenuchenko/personal/carrel-check.ics → 201 · 200 · 204"),
										},
										Mono: true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Tick{Done: true, Label: "Done"}}},
									{Parts: []w.CellPart{w.Text("Create, read back, delete · task")}},
									{
										Parts: []w.CellPart{
											w.Text("PUT · GET · DELETE /dav.php/calendars/m.yenuchenko/tasks/carrel-check.ics → 201 · 200 · 204"),
										},
										Mono: true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Tick{Label: "Done", Variant: w.TickEdgeDanger}}},
									{Parts: []w.CellPart{w.Text("Create, read back, delete · contact")}, Tone: w.CellDanger},
									{
										Parts: []w.CellPart{w.Text("PUT /dav.php/addressbooks/m.yenuchenko/contacts/carrel-check.vcf → 403 Forbidden")},
										Mono:  true,
									},
								},
								State: w.RowPicked,
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Tick{Label: "Done", Variant: w.TickEdgeDusty}}},
									{Parts: []w.CellPart{w.Text("Create, read back, delete · file")}, Tone: w.CellMuted},
									{
										Parts: []w.CellPart{w.Text("skipped: no collection of this kind that can be written to")},
										Mono:  true,
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

// Admin6 is frame «Администрирование · депонирование ключей» (s-admin-6).
func Admin6() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Administration"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "adm", Label: "Settings", Href: "#"},
								{Icon: "shield", Label: "Installation", Href: "#"},
								{Icon: "search", Label: "DAV validator", Href: "#"},
								{Icon: "shield", Label: "Key escrow", Href: "#"},
								{Icon: "clock", Label: "Audit log", Href: "#"},
							},
							On: 6,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Text:   "Mail is going out",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last sent "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
						w.RailStatus{
							Text:   "Checks passed",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last check "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Key escrow"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("4 of 6 accounts open with the master password")}}},
					},
					w.Msg{
						Icon: "shield",
						Parts: []w.MsgPart{
							w.Text("Key escrow is not a key backup. It is a way in: an administrator holding the master password can decrypt a covered user’s saved DAV credentials, and through them their contacts and calendars. Every covered user is told so in their profile, and every recovery is logged and shown to them."),
						},
						Tone: w.MsgAlert,
					},
					w.Form{
						Items: []w.FormPart{
							w.FieldSet{
								Rubric: w.Rubric{Text: "Policy"},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Check{On: true, Name: "coverNew", Value: "on", Parts: []w.CheckPart{w.Text("New accounts are covered")}},
										},
										Gap:   7,
										Place: w.Grown,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Turning coverage off leaves copies already taken alone: deleting one is its owner’s decision."),
												},
											},
										},
										Place: w.Grown,
									},
									w.Field{Parts: []w.FieldPart{w.Button{Label: "Save", Submit: true}}},
								},
								Variant: w.FieldSetLead,
								Action:  "/admin/escrow",
								Fields:  []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "action", Value: "policy"}},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Master password"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Current",
										Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "current", Label: "Current", Required: true}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "New",
										Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "master", Label: "New", Required: true}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Repeat",
										Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "master2", Label: "Repeat", Required: true}},
										Width: w.FieldLG,
									},
									w.Field{Parts: []w.FieldPart{w.Button{Label: "Change", Submit: true}}},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Only the recovery private key is re-encrypted. The deposited copies are untouched, so nobody has to opt in again."),
												},
											},
										},
										Place: w.Grown,
									},
								},
								Action: "/admin/escrow",
								Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "action", Value: "master"}},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Recover an account"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Account",
										Parts: []w.FieldPart{w.Input{Name: "login", Value: "viktor", Label: "Account", Required: true}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Master password",
										Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "master", Label: "Master password", Required: true}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Temporary password for the user",
										Parts: []w.FieldPart{
											w.Input{
												Name:     "temp",
												Value:    "9f2-brick-lamp",
												Label:    "Temporary password for the user",
												Mono:     true,
												Required: true,
											},
										},
										Width: w.FieldMD,
									},
									w.Field{Parts: []w.FieldPart{w.Button{Label: "Recover", Kind: w.ButtonDanger, Submit: true}}},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Connections survive a recovery. The user sees it by the password-history link they were given; a letter goes too when mail can reach them."),
												},
											},
										},
										Place: w.Grown,
									},
								},
								Action: "/admin/escrow",
								Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "action", Value: "recover"}},
							},
						},
					},
					w.ListGroup{
						Title: w.Rubric{Text: "Whom the master password opens", Form: w.RubricInline},
						Count: &w.Num{Value: 4},
					},
					w.Table{
						Cols: []w.TableCol{{Label: "Account"}, {Label: "Deposited"}, {Label: "Last recovery", Rank: w.Rank3}},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("anna")}},
									{Parts: []w.CellPart{w.Text("2026-03-02")}, Mono: true},
									{Parts: []w.CellPart{w.Text("never")}, Mono: true, Tone: w.CellDusty},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("viktor")}},
									{Parts: []w.CellPart{w.Text("2026-04-19")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-14 18:40")}, Mono: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("olga")}},
									{Parts: []w.CellPart{w.Text("2026-05-11")}, Mono: true},
									{Parts: []w.CellPart{w.Text("never")}, Mono: true, Tone: w.CellDusty},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("petr")}},
									{Parts: []w.CellPart{w.Text("2026-06-02")}, Mono: true},
									{Parts: []w.CellPart{w.Text("never")}, Mono: true, Tone: w.CellDusty},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("ivan")}},
									{Parts: []w.CellPart{w.Text("not deposited")}, Mono: true, Tone: w.CellDusty},
									{Parts: []w.CellPart{w.Text("—")}, Mono: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("mix")}},
									{Parts: []w.CellPart{w.Text("not deposited")}, Mono: true, Tone: w.CellDusty},
									{Parts: []w.CellPart{w.Text("—")}, Mono: true},
								},
							},
						},
					},
				},
			},
		},
	}
}

// Admin30 is frame «Депонирование · учётная запись восстановлена» (s-admin-30).
func Admin30() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Administration"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "adm", Label: "Settings", Href: "#"},
								{Icon: "shield", Label: "Installation", Href: "#"},
								{Icon: "search", Label: "DAV validator", Href: "#"},
								{Icon: "shield", Label: "Key escrow", Href: "#"},
								{Icon: "clock", Label: "Audit log", Href: "#"},
							},
							On: 6,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Text:   "Mail is going out",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last sent "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
						w.RailStatus{
							Text:   "Checks passed",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last check "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Key escrow"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("4 of 6 accounts open with the master password")}}},
					},
					w.Msg{
						Icon: "shield",
						Parts: []w.MsgPart{
							w.Text("Key escrow is not a key backup. It is a way in: an administrator holding the master password can decrypt a covered user’s saved DAV credentials, and through them their contacts and calendars. Every covered user is told so in their profile, and every recovery is logged and shown to them."),
						},
						Tone: w.MsgAlert,
					},
					w.Msg{Parts: []w.MsgPart{w.Text("The account of "), w.Strong("viktor"), w.Text(" is recovered.")}},
					w.Pad{
						Items: []w.PadPart{
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text:  "Temporary password",
										Items: []w.RubricPart{w.Mono{Text: "visible only now", Variant: w.MonoTightDanger}},
									},
									w.Diag{Text: "login     viktor\npassword  9f2-brick-lamp", Variant: w.DiagInk, ID: "s-admin-30-secret"},
									w.CardRow{
										Kind: w.CardTop8,
										Gap:  8,
										Items: []w.CardRowPart{
											w.Button{Label: "Copy", Icon: "dup", Kind: w.ButtonPrimary, CopyFrom: "s-admin-30-secret", Copied: "Copied"},
											w.Button{Label: "Done", Href: "/admin/escrow"},
											w.Dim{
												Parts:   []w.Inline{w.Text("No letter went: viktor has no confirmed address. Tell them yourself.")},
												Size:    w.Dim11_5,
												Variant: w.DimWrap,
											},
										},
									},
								},
								Variant: w.CardChosenFlushGrow,
							},
						},
					},
					w.Form{
						Items: []w.FormPart{
							w.FieldSet{
								Rubric: w.Rubric{Text: "Policy"},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Check{On: true, Name: "coverNew", Value: "on", Parts: []w.CheckPart{w.Text("New accounts are covered")}},
										},
										Gap:   7,
										Place: w.Grown,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Turning coverage off leaves copies already taken alone: deleting one is its owner’s decision."),
												},
											},
										},
										Place: w.Grown,
									},
									w.Field{Parts: []w.FieldPart{w.Button{Label: "Save", Submit: true}}},
								},
								Variant: w.FieldSetLead,
								Action:  "/admin/escrow",
								Fields:  []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "action", Value: "policy"}},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Master password"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Current",
										Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "current", Label: "Current", Required: true}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "New",
										Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "master", Label: "New", Required: true}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Repeat",
										Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "master2", Label: "Repeat", Required: true}},
										Width: w.FieldLG,
									},
									w.Field{Parts: []w.FieldPart{w.Button{Label: "Change", Submit: true}}},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Only the recovery private key is re-encrypted. The deposited copies are untouched, so nobody has to opt in again."),
												},
											},
										},
										Place: w.Grown,
									},
								},
								Action: "/admin/escrow",
								Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "action", Value: "master"}},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Recover an account"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Account",
										Parts: []w.FieldPart{w.Input{Name: "login", Value: "viktor", Label: "Account", Required: true}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Master password",
										Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "master", Label: "Master password", Required: true}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Temporary password for the user",
										Parts: []w.FieldPart{
											w.Input{
												Name:     "temp",
												Value:    "9f2-brick-lamp",
												Label:    "Temporary password for the user",
												Mono:     true,
												Required: true,
											},
										},
										Width: w.FieldMD,
									},
									w.Field{Parts: []w.FieldPart{w.Button{Label: "Recover", Kind: w.ButtonDanger, Submit: true}}},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Connections survive a recovery. The user sees it by the password-history link they were given; a letter goes too when mail can reach them."),
												},
											},
										},
										Place: w.Grown,
									},
								},
								Action: "/admin/escrow",
								Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "action", Value: "recover"}},
							},
						},
					},
					w.ListGroup{
						Title: w.Rubric{Text: "Whom the master password opens", Form: w.RubricInline},
						Count: &w.Num{Value: 4},
					},
					w.Table{
						Cols: []w.TableCol{{Label: "Account"}, {Label: "Deposited"}, {Label: "Last recovery", Rank: w.Rank3}},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("anna")}},
									{Parts: []w.CellPart{w.Text("2026-03-02")}, Mono: true},
									{Parts: []w.CellPart{w.Text("never")}, Mono: true, Tone: w.CellDusty},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("viktor")}},
									{Parts: []w.CellPart{w.Text("2026-04-19")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-14 18:40")}, Mono: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("olga")}},
									{Parts: []w.CellPart{w.Text("2026-05-11")}, Mono: true},
									{Parts: []w.CellPart{w.Text("never")}, Mono: true, Tone: w.CellDusty},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("petr")}},
									{Parts: []w.CellPart{w.Text("2026-06-02")}, Mono: true},
									{Parts: []w.CellPart{w.Text("never")}, Mono: true, Tone: w.CellDusty},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("ivan")}},
									{Parts: []w.CellPart{w.Text("not deposited")}, Mono: true, Tone: w.CellDusty},
									{Parts: []w.CellPart{w.Text("—")}, Mono: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("mix")}},
									{Parts: []w.CellPart{w.Text("not deposited")}, Mono: true, Tone: w.CellDusty},
									{Parts: []w.CellPart{w.Text("—")}, Mono: true},
								},
							},
						},
					},
				},
			},
		},
	}
}

// Admin31 is frame «Депонирование · не включено» (s-admin-31).
func Admin31() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Administration"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "adm", Label: "Settings", Href: "#"},
								{Icon: "shield", Label: "Installation", Href: "#"},
								{Icon: "search", Label: "DAV validator", Href: "#"},
								{Icon: "shield", Label: "Key escrow", Href: "#"},
								{Icon: "clock", Label: "Audit log", Href: "#"},
							},
							On: 6,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Text:   "Mail is going out",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last sent "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
						w.RailStatus{
							Text:   "Checks passed",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last check "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Key escrow"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Off")}}},
					},
					w.Msg{
						Icon: "shield",
						Parts: []w.MsgPart{
							w.Text("Key escrow is not a key backup. It is a way in: an administrator holding the master password can decrypt a covered user’s saved DAV credentials, and through them their contacts and calendars. Every covered user is told so in their profile, and every recovery is logged and shown to them."),
						},
						Tone: w.MsgAlert,
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Master password",
								Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "master", Label: "Master password", Required: true}},
								Width: w.FieldLG,
							},
							w.Field{
								Label: "Repeat",
								Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "master2", Label: "Repeat", Required: true}},
								Width: w.FieldLG,
							},
							w.Field{Parts: []w.FieldPart{w.Button{Label: "Turn on key escrow", Kind: w.ButtonPrimary, Submit: true}}},
							w.Field{
								Parts: []w.FieldPart{
									w.FieldHint{
										Parts: []w.Inline{w.Text("If you lose this password, deposited keys cannot be opened. This is not backup.")},
									},
								},
								Place: w.Grown,
							},
						},
						Action: "/admin/escrow",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}, {Name: "action", Value: "enable"}},
					},
				},
			},
		},
	}
}

// Admin7 is frame «Администрирование · журнал» (s-admin-7).
func Admin7() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Administration"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "adm", Label: "Settings", Href: "#"},
								{Icon: "shield", Label: "Installation", Href: "#"},
								{Icon: "search", Label: "DAV validator", Href: "#"},
								{Icon: "dup", Label: "Key escrow", Href: "#"},
								{Icon: "clock", Label: "Audit log", Href: "#"},
							},
							On: 7,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Tone:   w.RailStatusAlert,
							Text:   "SMTP not configured",
							Detail: w.Dim{Parts: []w.Inline{w.Text("No mail is going out.")}, Size: w.Dim11},
						},
						w.RailStatus{
							Text:   "Checks passed",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last check "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11},
						},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Audit log"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("1 240 records")}}},
						Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "Export CSV"}}},
					},
					w.Bar{
						Items: []w.BarPart{
							w.BarForm{
								Items: []w.BarFormPart{
									w.Drop{
										Trigger: w.Pick{Value: "All actions", Unset: true, Variant: w.PickH26W190},
										Menu: w.Menu{
											Items: []w.MenuPart{
												w.MenuItem{Label: "All actions", On: true, Name: "action", Value: ""},
												w.MenuSep{},
												w.MenuItem{Label: "login.ok", Name: "action", Value: "login.ok"},
												w.MenuItem{Label: "login.fail", Name: "action", Value: "login.fail"},
												w.MenuItem{Label: "escrow.recover", Name: "action", Value: "escrow.recover"},
												w.MenuItem{Label: "settings.instance", Name: "action", Value: "settings.instance"},
											},
											Width: w.MenuWidth290,
											Pop:   w.MenuUnder,
										},
									},
									w.Dim{Parts: []w.Inline{w.Text("since")}},
									w.Input{
										Kind:    w.InputDate,
										Name:    "since",
										Value:   "2026-08-01",
										Label:   "Since",
										Mono:    true,
										Variant: w.InputH26W120,
									},
								},
								Action: "/admin/audit",
								Show:   w.Button{Label: "Show", Submit: true},
							},
							w.Seg{
								Options: []w.SegOption{
									{Label: "All", Href: "/admin/audit?since=2026-08-01", On: true},
									{Label: "Security", Href: "/admin/audit?show=security&since=2026-08-01"},
									{Label: "Failures", Href: "/admin/audit?show=failures&since=2026-08-01"},
								},
							},
							w.Button{Folds: true, Label: "Columns", Icon: "cols", Title: "Columns"},
						},
					},
					w.Table{
						Cols: []w.TableCol{
							{Label: "When"},
							{Label: "Action"},
							{Label: "Actor", Rank: w.Rank3},
							{Label: "Target", Rank: w.Rank4},
							{Label: "Detail"},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("2026-08-15 09:20")}, Mono: true},
									{Parts: []w.CellPart{w.Text("login.ok")}, Mono: true},
									{Parts: []w.CellPart{w.Text("mix")}},
									{Parts: []w.CellPart{w.Text("—")}},
									{Parts: []w.CellPart{w.Text("—")}, Mono: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("2026-08-14 22:07")}, Mono: true},
									{Parts: []w.CellPart{w.Text("login.fail")}, Mono: true, Tone: w.CellDanger},
									{Parts: []w.CellPart{w.Text("—")}},
									{Parts: []w.CellPart{w.Text("anna")}},
									{Parts: []w.CellPart{w.Text("3rd attempt · ip 203.0.113.9 · rate-limited")}, Mono: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("2026-08-14 18:40")}, Mono: true},
									{Parts: []w.CellPart{w.Text("escrow.recover")}, Mono: true},
									{Parts: []w.CellPart{w.Text("mix")}},
									{Parts: []w.CellPart{w.Text("viktor")}},
									{Parts: []w.CellPart{w.Text("notice mailed")}, Mono: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("2026-08-12 10:02")}, Mono: true},
									{Parts: []w.CellPart{w.Text("settings.instance")}, Mono: true},
									{Parts: []w.CellPart{w.Text("mix")}},
									{Parts: []w.CellPart{w.Text("—")}},
									{Parts: []w.CellPart{w.Text("smtp_host, smtp_port")}, Mono: true},
								},
							},
						},
					},
				},
			},
		},
	}
}
