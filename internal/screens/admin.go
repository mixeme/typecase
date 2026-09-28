// Раздел макета: 7.19 Администрирование.

package screens

import w "carrel.dev/typecase/widget"

// Admin1 is frame «Администрирование · как заводят пользователей» (s-admin-1).
func Admin1() Screen {
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
											w.Field{Label: "Role", Parts: []w.FieldPart{w.Pick{Value: "User"}}, Width: w.FieldSM},
											w.Button{Label: "Create link", Kind: w.ButtonPrimary},
										},
										Wide: true,
									},
									w.CardRow{
										Kind: w.CardTop10,
										Gap:  6,
										Items: []w.CardRowPart{
											w.Input{Value: "https://carrel.example/invite/7f3a…", Mono: true, Variant: w.Input11Fill},
											w.Button{Label: "Copy", Small: true},
										},
									},
									w.CardFields{
										Items: []w.CardFieldsPart{
											w.Field{
												Label: "Or send to",
												Parts: []w.FieldPart{w.Input{Placeholder: "anna@example.org", Label: "Or send to"}},
												Width: w.FieldMD,
											},
											w.Button{Label: "Send invitation"},
										},
										Wide: true,
										Rule: true,
									},
								},
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
											w.Field{Label: "Role", Parts: []w.FieldPart{w.Pick{Value: "User"}}, Width: w.FieldXS},
											w.Field{Label: "Login", Parts: []w.FieldPart{w.Input{Value: "viktor", Label: "Login"}}, Width: w.FieldSM},
											w.Field{
												Label:      "Email",
												LabelAside: w.Dim{Parts: []w.Inline{w.Text("(optional)")}},
												Parts:      []w.FieldPart{w.Input{Placeholder: "—", Label: "Email"}},
												Width:      w.FieldMD,
											},
											w.Field{
												Label: "Temporary password",
												Parts: []w.FieldPart{w.Input{Value: "9f2-brick-lamp", Label: "Temporary password", Mono: true}},
												Width: w.FieldMD,
											},
											w.Button{Label: "Create user", Kind: w.ButtonPrimary},
										},
										Wide: true,
									},
								},
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
									w.Check{On: true, Parts: []w.CheckPart{w.Text("Publish the registration form")}},
									w.CardRow{
										Kind: w.CardTop10,
										Gap:  6,
										Items: []w.CardRowPart{
											w.Input{Value: "https://carrel.example/register", Mono: true, Variant: w.Input11Fill},
											w.Button{Label: "Copy", Small: true},
										},
									},
									w.CardRow{Kind: w.CardTop10, Items: []w.CardRowPart{w.Button{Label: "Save"}}},
								},
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
											w.Button{Label: "Copy link", Small: true},
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
											w.Button{Label: "Copy link", Small: true},
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
							w.Input{
								Kind:        w.InputSearch,
								Placeholder: "Filter by login",
								Label:       "Filter by login",
								Variant:     w.InputH26W200,
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
						},
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("Deleting an account is confirmed in a dialogue of its own. Reset a password, in the head, sets a temporary password and names what a reset destroys before it is pressed."),
						},
					},
				},
			},
		},
	}
}

// Admin3 is frame «Администрирование · настройки» (s-admin-3).
func Admin3() Screen {
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
						Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "Save", Kind: w.ButtonPrimary}}},
					},
					w.Form{
						Items: []w.FormPart{
							w.FieldSet{
								Rubric: w.Rubric{Text: "Global"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Invite lifetime, h",
										Parts: []w.FieldPart{w.Input{Value: "72", Label: "Invite lifetime, h", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Idle timeout, h",
										Parts: []w.FieldPart{w.Input{Value: "12", Label: "Idle timeout, h", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Session life, d",
										Parts: []w.FieldPart{w.Input{Value: "30", Label: "Session life, d", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Instance name",
										Parts: []w.FieldPart{w.Input{Value: "Carrel", Label: "Instance name"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Public address",
										Parts: []w.FieldPart{
											w.Input{Value: "https://carrel.example", Label: "Public address", Mono: true},
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
										Parts: []w.FieldPart{w.Input{Placeholder: "smtp.example.org", Label: "Host", Mono: true}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Port",
										Parts: []w.FieldPart{w.Input{Value: "587", Label: "Port", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{Label: "Encryption", Parts: []w.FieldPart{w.Pick{Value: "STARTTLS"}}, Width: w.FieldSM},
									w.Field{
										Label: "Username",
										Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Username"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Password",
										Parts: []w.FieldPart{w.Input{Placeholder: "leave blank to keep the current one", Label: "Password"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "From address",
										Parts: []w.FieldPart{w.Input{Placeholder: "carrel@example.org", Label: "From address"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "From name",
										Parts: []w.FieldPart{w.Input{Placeholder: "Carrel", Label: "From name"}},
										Width: w.FieldMD,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Letters from users", Items: []w.RubricPart{w.Badge{Label: "off"}}},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Check{Parts: []w.CheckPart{w.Text("Let users send letters and links from Carrel")}},
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
												Options: []w.SegOption{{Label: "noreply@carrel.example", On: true}, {Label: "login@carrel.example"}},
												Variant: w.SegStart,
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Your relay decides which of the two it will let out. The envelope sender stays "),
													w.Mono{Text: "noreply@carrel.example"},
													w.Text(" either way, so SPF is unaffected; DKIM and DMARC for this domain are set up on the relay, not here."),
												},
											},
										},
										Place: w.Grown,
									},
									w.Field{
										Label: "Attachment, MB",
										Parts: []w.FieldPart{w.Input{Value: "2", Label: "Attachment, MB", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Letter, MB",
										Parts: []w.FieldPart{w.Input{Value: "8", Label: "Letter, MB", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Recipients",
										Parts: []w.FieldPart{w.Input{Value: "10", Label: "Recipients", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Label: "Letters a day",
										Parts: []w.FieldPart{w.Input{Value: "50", Label: "Letters a day", Mono: true}},
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
													w.Input{Placeholder: "you@example.org", Label: "Send test message to", Variant: w.InputFill},
													w.Button{Label: "Send test", State: w.ButtonDisabled},
												},
												Gap: w.LineGap8,
											},
											w.FieldHint{Parts: []w.Inline{w.Text("Mail is not configured: there is nothing to send through yet.")}},
										},
										Width: w.FieldMD,
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

// Admin4 is frame «Администрирование · проверка установки» (s-admin-4).
func Admin4() Screen {
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
						Title: &w.HeadTitle{Text: "Installation check"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{w.Text("Carrel asks itself from outside, at the address you declared as the base one")},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Field{
									Parts: []w.FieldPart{w.Input{Value: "https://carrel.example", Mono: true, Variant: w.Input11_5}},
									Width: w.FieldW280,
								},
								w.Button{Label: "Run the check", Kind: w.ButtonPrimary},
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
							w.Text(" · 1 warning"),
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
								Parts: []w.FieldPart{w.Input{Value: "https://dav.office.example/dav.php/", Label: "Server URL", Mono: true}},
								Width: w.FieldXL,
							},
							w.Field{
								Label: "Username",
								Parts: []w.FieldPart{w.Input{Value: "m.yenuchenko", Label: "Username"}},
								Width: w.FieldMD,
							},
							w.Field{
								Label: "Password",
								Parts: []w.FieldPart{w.Input{Value: "••••••••", Label: "Password"}},
								Width: w.FieldMD,
							},
							w.Field{Parts: []w.FieldPart{w.Button{Label: "Test discovery", Kind: w.ButtonPrimary}}},
						},
					},
					w.MainDiag{
						Diag: w.Diag{
							Text: "[well-known]  GET /.well-known/caldav → 301 → /dav.php/\n[principal]   PROPFIND /dav.php/ → 207, principal /dav.php/principals/m.yenuchenko/\n[home-set]    calendar-home-set → /dav.php/calendars/m.yenuchenko/\n[home-set]    addressbook-home-set → /dav.php/addressbooks/m.yenuchenko/\n[collections] 4 calendars, 1 address book, 0 plain folders",
						},
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("The second, mutating stage — create test objects, exercise the operations, read them back, delete them — never runs silently: it is agreed to explicitly and written to the audit log."),
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
							w.Text("Key escrow is not a key backup. It is a way in: an administrator holding the master password can decrypt a covered user’s saved DAV credentials, and through them their contacts and calendars. Every covered user is told so in their profile, and every recovery is logged and mailed to them."),
						},
						Tone: w.MsgAlert,
					},
					w.Form{
						Items: []w.FormPart{
							w.FieldSet{
								Rubric: w.Rubric{Text: "Policy"},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{w.Check{On: true, Parts: []w.CheckPart{w.Text("New accounts are covered")}}},
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
									w.Field{Parts: []w.FieldPart{w.Button{Label: "Save"}}},
								},
								Variant: w.FieldSetLead,
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Master password"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Current",
										Parts: []w.FieldPart{w.Input{Value: "••••••••••••••••", Label: "Current"}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "New",
										Parts: []w.FieldPart{w.Input{Value: "••••••••••••••••", Label: "New"}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Repeat",
										Parts: []w.FieldPart{w.Input{Value: "••••••••••••••••", Label: "Repeat"}},
										Width: w.FieldLG,
									},
									w.Field{Parts: []w.FieldPart{w.Button{Label: "Change"}}},
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
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Recover an account"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Account", Parts: []w.FieldPart{w.Input{Value: "viktor", Label: "Account"}}, Width: w.FieldMD},
									w.Field{
										Label: "Master password",
										Parts: []w.FieldPart{w.Input{Value: "••••••••••••••••", Label: "Master password"}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Temporary password for the user",
										Parts: []w.FieldPart{w.Input{Value: "9f2-brick-lamp", Label: "Temporary password for the user", Mono: true}},
										Width: w.FieldMD,
									},
									w.Field{Parts: []w.FieldPart{w.Button{Label: "Recover", Kind: w.ButtonDanger}}},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Connections survive a recovery: the deposited key is opened with the master password and re-wrapped under a temporary one. The user is mailed about it — always, not at the administrator’s discretion."),
												},
											},
										},
										Place: w.Grown,
									},
								},
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

// Admin7 is frame «Администрирование · журнал» (s-admin-7).
func Admin7() Screen {
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
							w.Pick{Value: "All actions", Unset: true, Variant: w.PickH26W190},
							w.Input{Kind: w.InputDate, Value: "2026-08-01", Mono: true, Variant: w.InputH26W120},
							w.Seg{Options: []w.SegOption{{Label: "All", On: true}, {Label: "Security"}, {Label: "Failures"}}},
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
									{Parts: []w.CellPart{w.Text("ip 10.0.0.4")}, Mono: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("2026-08-14 22:07")}, Mono: true},
									{Parts: []w.CellPart{w.Text("login.fail")}, Mono: true, Tone: w.CellDanger},
									{Parts: []w.CellPart{w.Text("—")}},
									{Parts: []w.CellPart{w.Text("anna")}},
									{Parts: []w.CellPart{w.Text("3rd attempt · rate-limited")}, Mono: true},
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
