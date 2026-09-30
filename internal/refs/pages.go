package refs

import w "carrel.dev/typecase/widget"

// Pages without a shell: sign-in, «about», a desktop
// window's inside, a page on paper. A record per page frame: the root's own
// tree is the page's whole content, so every frame is a form of its own.

func page(state string, fixture Renderer, frame, class string, n int) Entry {
	return e("pages", state, fixture, ref(frame, class, n))
}

var _ = add(
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogLogo{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo60}},
						w.Tagline{
							Lead:     "One desk. Every server.",
							Text:     "Calendars, contacts, tasks, notes and files — one place.",
							Centered: true,
						},
						w.AuthStack{
							Items: []w.AuthStackPart{
								w.Form{
									Items: []w.FormPart{
										w.Field{
											Label: "Login",
											Parts: []w.FieldPart{w.Input{Value: "mix", Label: "Login", Focus: true}},
											Width: w.FieldW100,
										},
										w.Field{
											Label: "Password",
											Parts: []w.FieldPart{w.Input{Value: "••••••••••", Label: "Password"}},
											Width: w.FieldW100,
										},
									},
									Variant: w.FormColumnFlush,
								},
								w.Button{Label: "Sign in", Kind: w.ButtonPrimary},
							},
							Links: []w.Link{{Text: "Create an account"}, {Text: "Forgotten your password?"}},
						},
					},
					Variant: w.DialogNarrow,
				},
			},
		},
	}, "s-auth-1", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogTitle{Text: "Create an account", Gap: w.DialogTitleGap4},
						w.DialogSub{
							Parts: []w.Inline{
								w.Text("A confirmation link will be sent to the address you give; the account cannot be used until you open it."),
							},
						},
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("Your password is the only key to your own data. Nobody — including this server — can recover it for you."),
							},
							Tone:  w.MsgAlert,
							Space: w.MsgFlushBottom12,
						},
						w.Form{
							Items: []w.FormPart{
								w.Field{
									Label: "Login",
									Parts: []w.FieldPart{w.Input{Value: "viktor", Label: "Login", Focus: true}},
									Width: w.FieldMD,
								},
								w.Field{
									Label: "Email",
									Parts: []w.FieldPart{w.Input{Value: "v@example.org", Label: "Email"}},
									Width: w.FieldMD,
								},
								w.Field{
									Label: "Password",
									Parts: []w.FieldPart{w.Input{Value: "••••••••••••", Label: "Password"}},
									Width: w.FieldMD,
								},
								w.Field{
									Label: "Repeat password",
									Parts: []w.FieldPart{w.Input{Value: "••••••••••••", Label: "Repeat password"}},
									Width: w.FieldMD,
								},
							},
							Variant: w.FormFlushAll,
						},
						w.FormFoot{
							Items: []w.FormFootPart{
								w.Button{Label: "Create account", Kind: w.ButtonPrimary},
								w.Span{Face: w.FaceAccent12, Text: "Already have an account? Sign in"},
							},
							Variant: w.FormFootFlush,
						},
					},
					Variant: w.DialogPage440,
				},
			},
		},
	}, "s-auth-2", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogLogo{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo48}, Wide: true},
						w.DialogTitle{Text: "Check your email", Gap: w.DialogTitleGap6},
						w.Prose{
							Parts: []w.Inline{
								w.Text("The letter went to "),
								w.Mono{Text: "v@example.org"},
								w.Text(". The account cannot be used until you open the link in it.\n\nSending again with the same login and address, while the account is unconfirmed, simply replaces the letter and the password."),
							},
							Variant: w.ProseMuted,
						},
						w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Back to sign in"}}, Variant: w.FormFootFlush},
					},
					Variant: w.DialogNarrow,
				},
			},
		},
	}, "s-auth-3", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogTitle{Text: "Set up Carrel", Gap: w.DialogTitleGap4},
						w.DialogSub{
							Parts: []w.Inline{w.Text("The account you create here is the first one, and it can manage everyone else.")},
						},
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("Your password is the only key to your own data. Nobody — including this server — can recover it for you."),
							},
							Tone:  w.MsgAlert,
							Space: w.MsgFlushBottom12,
						},
						w.Form{
							Items: []w.FormPart{
								w.Field{
									Label: "Login",
									Parts: []w.FieldPart{w.Input{Placeholder: "a-z 0-9 . _ -", Label: "Login", Mono: true, Focus: true}},
									Width: w.FieldMD,
								},
								w.Field{
									Label:      "Email",
									LabelAside: w.Dim{Parts: []w.Inline{w.Text("(optional)")}},
									Parts:      []w.FieldPart{w.Input{Placeholder: "—", Label: "Email"}},
									Width:      w.FieldMD,
								},
								w.Field{
									Label: "Password",
									Parts: []w.FieldPart{w.Input{Value: "••••••••••••", Label: "Password"}},
									Width: w.FieldMD,
								},
								w.Field{
									Label: "Repeat password",
									Parts: []w.FieldPart{w.Input{Value: "••••••••••••", Label: "Repeat password"}},
									Width: w.FieldMD,
								},
							},
							Variant: w.FormFlushAll,
						},
						w.FormFoot{
							Items:   []w.FormFootPart{w.Button{Label: "Create administrator", Kind: w.ButtonPrimary}},
							Variant: w.FormFootFlush,
						},
					},
					Variant: w.DialogPage440,
				},
			},
		},
	}, "s-auth-4", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogTitle{Text: "Forgotten your password?", Gap: w.DialogTitleGap8},
						w.Prose{
							Parts: []w.Inline{
								w.Text("Enter the login or the confirmed email of the account. There is no way to set a new password and keep the data stored here: the encryption key is derived from the password, and the server never has that key."),
							},
							Variant: w.ProseMutedClose,
						},
						w.Form{
							Items: []w.FormPart{
								w.Field{
									Label: "Login or email",
									Parts: []w.FieldPart{w.Input{Value: "mix", Label: "Login or email", Focus: true}},
									Width: w.FieldW100,
								},
							},
							Variant: w.FormColumnFlush,
						},
						w.FormFoot{
							Items:   []w.FormFootPart{w.Button{Label: "Continue", Kind: w.ButtonPrimary}, w.Button{Label: "Back to sign in"}},
							Variant: w.FormFootFlush,
						},
					},
					Variant: w.DialogPage440,
				},
			},
		},
	}, "s-auth-5", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogTitle{Text: "What we can do", Gap: w.DialogTitleGap8},
						w.Prose{
							Parts: []w.Inline{
								w.Text("Account "),
								w.Strong("mix"),
								w.Text(" has a confirmed email. A silent password reset does not exist: a new password would be a new key, and everything stored under the old one would be lost anyway."),
							},
							Variant: w.ProseMuted,
						},
						w.Form{
							Items: []w.FormPart{
								w.Card{
									Items: []w.CardPart{
										w.CardText{Kind: w.CardTerm, Parts: []w.CardTextPart{w.Text("Wipe the account")}},
										w.Prose{
											Parts: []w.Inline{
												w.Text("We send a letter to the confirmed address. The link in it is not a reset: after you confirm, everything stored here is erased — connections become unreadable and are deleted — and the same link then lets you set a new password. Address books, calendars, tasks and notes on the servers stay where they are."),
											},
											Variant: w.ProseMutedClose,
										},
										w.Button{Label: "Send a confirmation letter", Kind: w.ButtonPrimary},
									},
									Variant: w.CardSnug,
								},
								w.Card{
									Items: []w.CardPart{
										w.CardText{Kind: w.CardTerm, Parts: []w.CardTextPart{w.Text("Ask the administrator")}},
										w.Prose{
											Parts: []w.Inline{
												w.Text("This account deposited its encryption key. The administrator can recover it with the instance master password, and your data is kept. Shown only when deposit is on for this account."),
											},
											Variant: w.ProseMutedClose,
										},
										w.Button{Label: "Ask the administrator"},
									},
									Variant: w.CardSnug,
								},
							},
							Variant: w.FormColumnGap10,
						},
						w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Back to sign in"}}, Variant: w.FormFootFlush},
					},
					Variant: w.DialogPage520,
				},
			},
		},
	}, "s-auth-6", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogTitle{Text: "We cannot help from here", Gap: w.DialogTitleGap8},
						w.Prose{
							Parts: []w.Inline{
								w.Text("This login has no confirmed email — an address that was entered but never confirmed does not count. There is no letter we can send, and there is no password reset that would keep your data.\n\nIf you deposited the encryption key, ask the administrator in person. Otherwise this account cannot be recovered from the sign-in screen."),
							},
							Variant: w.ProseMuted,
						},
						w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Back to sign in"}}, Variant: w.FormFootFlush},
					},
					Variant: w.DialogPage520,
				},
			},
		},
	}, "s-auth-7", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogTitle{Text: "Check your inbox", Gap: w.DialogTitleGap8},
						w.Prose{
							Parts: []w.Inline{
								w.Text("If that account has a confirmed email, a letter is on its way. The link erases everything stored here after you confirm it, and then asks for a new password. It cannot bring your data back."),
							},
							Variant: w.ProseMuted,
						},
						w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Back to sign in"}}, Variant: w.FormFootFlush},
					},
					Variant: w.DialogPage520,
				},
			},
		},
	}, "s-auth-8", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogTitle{Text: "The administrator has been asked", Gap: w.DialogTitleGap8},
						w.Prose{
							Parts: []w.Inline{
								w.Text("A recovery request is recorded for this account. The administrator still has to use the instance master password. Until then, the sign-in screen will not accept a new password."),
							},
							Variant: w.ProseMuted,
						},
						w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Back to sign in"}}, Variant: w.FormFootFlush},
					},
					Variant: w.DialogPage520,
				},
			},
		},
	}, "s-auth-9", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogTitle{Text: "Erase everything and start again?", Gap: w.DialogTitleGap8},
						w.Prose{
							Parts: []w.Inline{
								w.Text("This is not a password recovery. The key to your data comes from the password you forgot, so nothing stored under it can be read again — by you, by us, by anyone.\n\nConfirming erases what Carrel keeps for "),
								w.Strong("mix"),
								w.Text(": saved connections and their passwords, subscriptions, tags and the settings of this desk. Address books, calendars, tasks and notes on the servers are not touched — you connect to them again afterwards.\n\nThe account itself stays: same login, same confirmed address. Only the password is new, and you set it on the next screen."),
							},
							Variant: w.ProseMuted,
						},
						w.FormFoot{
							Items: []w.FormFootPart{
								w.Button{Label: "Erase and set a new password", Kind: w.ButtonDangerSolid},
								w.Button{Label: "Cancel"},
							},
							Variant: w.FormFootFlush,
						},
					},
					Variant: w.DialogPage520,
				},
			},
		},
	}, "s-auth-10", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogTitle{Text: "Set a new password", Gap: w.DialogTitleGap4},
						w.DialogSub{
							Parts: []w.Inline{w.Text("The account "), w.Mono{Text: "mix"}, w.Text(" is empty. Nothing was kept.")},
						},
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("Your password is the only key to your own data. Nobody — including this server — can recover it for you."),
							},
							Tone:  w.MsgAlert,
							Space: w.MsgFlushBottom12,
						},
						w.Form{
							Items: []w.FormPart{
								w.Field{
									Label: "Password",
									Parts: []w.FieldPart{w.Input{Value: "••••••••••••", Label: "Password", Focus: true}},
									Width: w.FieldMD,
								},
								w.Field{
									Label: "Repeat password",
									Parts: []w.FieldPart{w.Input{Value: "••••••••••••", Label: "Repeat password"}},
									Width: w.FieldMD,
								},
							},
							Variant: w.FormFlushAll,
						},
						w.FormFoot{
							Items:   []w.FormFootPart{w.Button{Label: "Set the password and sign in", Kind: w.ButtonPrimary}},
							Variant: w.FormFootFlush,
						},
					},
					Variant: w.DialogPage440,
				},
			},
		},
	}, "s-auth-11", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogLogo{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo48}, Wide: true},
						w.DialogTitle{Text: "Invitation", Gap: w.DialogTitleGap6},
						w.Prose{
							Parts: []w.Inline{
								w.Text("This invitation link is not valid. It may have expired, already been used, or been revoked."),
							},
							Variant: w.ProseMuted,
						},
						w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Sign in", Kind: w.ButtonPrimary}}, Variant: w.FormFootFlush},
					},
					Variant: w.DialogNarrow,
				},
			},
		},
	}, "rest-1", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogLogo{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo48}, Wide: true},
						w.DialogTitle{Text: "Email address", Gap: w.DialogTitleGap6},
						w.Prose{Parts: []w.Inline{w.Text("Address confirmed. You can sign in now.")}, Variant: w.ProseMuted},
						w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Sign in", Kind: w.ButtonPrimary}}, Variant: w.FormFootFlush},
					},
					Variant: w.DialogNarrow,
				},
			},
		},
	}, "rest-2", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogLogo{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo42}, Wide: true},
						w.DialogTitle{Text: "Carrel"},
						w.DialogSub{Parts: []w.Inline{w.Mono{Text: "0.0.1 · 786949a"}}},
						w.Tagline{Lead: "One desk. Every server.", Text: "Calendars, contacts, tasks, notes and files — one place."},
						w.Prose{
							Parts: []w.Inline{
								w.Text("AGPL-3.0-or-later: if you run a modified version for others, you must make the source available to them."),
							},
							Variant: w.ProseMuted,
						},
						w.FormFoot{
							Items: []w.FormFootPart{
								w.Button{Label: "Source code", Small: true},
								w.Button{Label: "Third-party licences", Small: true},
								w.Push{Items: []w.PushPart{w.Button{Label: "Back to sign in"}}},
							},
							Variant: w.FormFootFlush,
						},
					},
					Variant: w.DialogNarrow,
				},
			},
		},
	}, "rest-3", "tc-root", 1),
	page("print", w.Page{
		Items: []w.PagePart{
			w.PrintTitle{Text: "Agenda · 10 – 16 August 2026"},
			w.PrintSources{Text: "nextcloud · Личный, Работа · radicale · vps · Проекты, Дежурства"},
			w.PrintRubric{Text: "Monday · 10 August"},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "10:00", Variant: w.MonoFace},
					w.PrintCell{Name: "Синк-ап по релизу", Detail: "Переговорная 2 · nextcloud · Личный"},
				},
			},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "all day", Variant: w.MonoFace},
					w.PrintCell{Name: "Дежурство: первая линия", Detail: "radicale · vps · Дежурства"},
				},
			},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "15:00", Variant: w.MonoFace},
					w.PrintCell{Name: "Созвон с подрядчиком", Detail: "Jitsi · nextcloud · Работа"},
				},
			},
			w.PrintRubric{Text: "Tuesday · 11 August"},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "09:30", Variant: w.MonoFace},
					w.PrintCell{Name: "Ревью архитектуры", Detail: "Переговорная 1 · radicale · vps · Проекты"},
				},
			},
			w.PrintFoot{Text: "Carrel · agenda · 2026-08-10 – 2026-08-16 · captured 2026-08-15 09:41 · 5 of 7 sources"},
		},
		Variant: w.PagePrint,
	}, "print-1", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Window{
				Items: []w.WindowPart{
					w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo42},
					w.Stack{
						Gap: w.StackGap9,
						Items: []w.StackPart{
							w.Card{
								Items: []w.CardPart{
									w.CardPick{
										Title: "Just for me",
										Note:  []w.CardTextPart{w.Text("Installs into your own profile. No administrator needed.")},
									},
								},
								Variant: w.CardChosenFlush,
							},
							w.Card{
								Items: []w.CardPart{
									w.CardPick{
										Ticked: w.TickedOff,
										Title:  "For everyone on this computer",
										Note: []w.CardTextPart{
											w.Text("Installs into "),
											w.Mono{Text: "Program Files"},
											w.Text(". Asks for administrator once; each account still keeps its own data and password."),
										},
									},
								},
								Variant: w.CardFlush,
							},
							w.Card{
								Items: []w.CardPart{
									w.CardPick{
										Ticked: w.TickedOff,
										Title:  "Portable — do not install",
										Note: []w.CardTextPart{
											w.Text("Unpacks into a folder you choose and keeps everything inside it: the app, the server binary and your encrypted state."),
										},
										Field: &w.Field{
											Parts: []w.FieldPart{
												w.Line{
													Items: []w.LinePart{
														w.Input{Placeholder: "E:\\Carrel", Label: "E:\\Carrel", Mono: true, Variant: w.Input11Fill},
														w.Button{Label: "Choose"},
													},
													Gap: w.LineGap8,
												},
												w.FieldHint{
													Parts: []w.Inline{
														w.Text("Nothing is written outside this folder: no registry, no start-up entry, no traces in the profile. Moving the folder moves Carrel with it."),
													},
												},
											},
											Width: w.FieldW100,
										},
									},
								},
								Variant: w.CardFlush,
							},
						},
					},
					w.Msg{
						Icon: "shield",
						Parts: []w.MsgPart{
							w.Strong("On a stick, the folder is the secret."),
							w.Text(" Your saved DAV passwords are encrypted with your Carrel password and stay encrypted there — but whoever finds the stick holds that file and can try guessing. Choose a long password, or keep the portable copy where you would keep keys."),
						},
						Tone:  w.MsgAlert,
						Space: w.MsgFlush,
					},
					w.Line{
						Items: []w.LinePart{
							w.Button{Label: "Continue", Kind: w.ButtonPrimary},
							w.Button{Label: "Cancel"},
							w.Dim{
								Parts:   []w.Inline{w.Text("64 MB · app, server, WebView2 check")},
								Size:    w.Dim11_5,
								Variant: w.DimWrap,
								Place:   w.Pushed,
							},
						},
						Gap:   w.LineGap8,
						Block: true,
					},
				},
				Gap: w.WindowGap13,
			},
		},
	}, "desktop-3", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Window{
				Items: []w.WindowPart{
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Everything lives here",
								Parts: []w.FieldPart{
									w.Input{
										Value:   "E:\\Carrel\\ — app, carrel.exe, data\\state.enc",
										Label:   "Everything lives here",
										Mono:    true,
										Variant: w.Input11_5,
									},
									w.FieldHint{
										Parts: []w.Inline{
											w.Text("The same folder on any computer: plug the stick in, run, sign in with your password. The data folder cannot be moved away from the app in this mode — that is what makes it portable."),
										},
									},
								},
								Place: w.Grown,
							},
							w.Field{
								Parts: []w.FieldPart{
									w.Check{
										Parts: []w.CheckPart{w.Dim{Parts: []w.Inline{w.Text("Start at sign-in — not available in a portable copy")}}},
									},
									w.Check{On: true, Parts: []w.CheckPart{w.Text("Closing the window hides to the tray")}},
								},
								Gap:   7,
								Place: w.Grown,
							},
							w.Field{
								Label: "Updates",
								Parts: []w.FieldPart{
									w.Line{
										Items: []w.LinePart{
											w.Input{
												Placeholder: "Replaces the files in this folder",
												Label:       "Replaces the files in this folder",
												Variant:     w.InputFill,
											},
											w.Button{Label: "Check now"},
										},
										Gap: w.LineGap8,
									},
									w.FieldHint{
										Parts: []w.Inline{
											w.Text("Works while the folder is writable. On a read-only medium the app says so and offers the download instead of installing it."),
										},
									},
								},
								Place: w.Grown,
							},
						},
						Variant: w.FormGap12,
					},
					w.Msg{
						Parts: []w.MsgPart{w.Text("A portable copy is a full Carrel, local or remote as usual — not a cut-down build.")},
						Space: w.MsgFlush,
					},
				},
				Gap: w.WindowGap12,
			},
		},
	}, "desktop-4", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Window{
				Items: []w.WindowPart{
					w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo42},
					w.Stack{
						Gap: w.StackGap10,
						Items: []w.StackPart{
							w.Card{
								Items: []w.CardPart{
									w.CardPick{
										Title: "Local — run it here",
										Note:  []w.CardTextPart{w.Text("Carrel runs inside the app, on this computer. Nothing to deploy.")},
										Aside: "Suits one person and one computer. A phone cannot open this Carrel until access over the network is allowed.",
									},
								},
								Variant: w.CardChosenFlush,
							},
							w.Card{
								Items: []w.CardPart{
									w.CardPick{
										Ticked: w.TickedOff,
										Title:  "Remote — connect to a server",
										Note:   []w.CardTextPart{w.Text("Carrel is already running — yours or a shared one. The app becomes its window.")},
										Aside:  "One instance for all your devices, with the same settings everywhere.",
									},
								},
								Variant: w.CardFlush,
							},
						},
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("The mode changes by signing out: the app shows this screen again. The data ("),
							w.Mono{Text: "state.enc"},
							w.Text(") lives in the OS user profile either way."),
						},
						Space: w.MsgFlush,
					},
					w.Line{Items: []w.LinePart{w.Button{Label: "Next", Kind: w.ButtonPrimary}}, Gap: w.LineGap8, Block: true},
				},
				Gap: w.WindowGap14,
			},
		},
	}, "desktop-5", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Window{
				Items: []w.WindowPart{
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Data folder",
								Parts: []w.FieldPart{
									w.Input{
										Value:   "C:\\Users\\mix\\AppData\\Local\\Carrel",
										Label:   "Data folder",
										Mono:    true,
										Variant: w.Input11_5,
									},
									w.FieldHint{
										Parts: []w.Inline{
											w.Text("Encrypted state lives here: your DAV connections and settings. Not one contact, event or note — those stay on the servers."),
										},
									},
								},
								Place: w.Grown,
							},
							w.Field{
								Label: "Port",
								Parts: []w.FieldPart{
									w.Input{Value: "chosen at every start · now 51734", Label: "Port", Mono: true},
									w.FieldHint{
										Parts: []w.Inline{
											w.Text("Nothing to set here. The app asks the system for a free port on "),
											w.Mono{Text: "127.0.0.1"},
											w.Text(", so a busy 8080, a second copy of Carrel or someone else’s server cannot collide with it. The current number is always in the status bar — it is needed only to open the address in a browser."),
										},
									},
								},
								Width: w.FieldFill300,
							},
							w.Field{Parts: []w.FieldPart{w.Check{On: true, Parts: []w.CheckPart{w.Text("Stop the server on exit")}}}},
							w.Field{
								Parts: []w.FieldPart{
									w.Check{Parts: []w.CheckPart{w.Text("Allow other devices on this network")}},
									w.FieldHint{
										Parts: []w.Inline{
											w.Text("While this is off, Carrel listens on "),
											w.Mono{Text: "127.0.0.1"},
											w.Text(" only — from outside it does not exist. Turning it on "),
											w.Strong("pins the port"),
											w.Text(" and gives "),
											w.Mono{Text: "http://192.168.1.14:8765"},
											w.Text(" for a phone and for DAVx5, and asks for a password like any other instance: an address that other devices remember cannot change at every start."),
										},
									},
								},
								Gap:     7,
								OwnLine: true,
								Place:   w.Grown,
							},
						},
						Variant: w.FormGap12,
					},
					w.FieldSet{
						Rubric: w.Rubric{Text: "Account"},
						Items: []w.FieldSetPart{
							w.Field{Label: "Login", Parts: []w.FieldPart{w.Input{Value: "mix", Label: "Login"}}, Width: w.FieldMD},
							w.Field{
								Label: "Password",
								Parts: []w.FieldPart{w.Input{Value: "••••••••••••", Label: "Password"}},
								Width: w.FieldMD,
							},
							w.Msg{
								Parts: []w.MsgPart{
									w.Text("The password is the only key to the saved connections. Running locally changes nothing about that: nobody can recover it, including the app on your own computer."),
								},
								Tone:  w.MsgAlert,
								Space: w.MsgFlushTop2,
								Place: w.Grown,
							},
						},
						Variant: w.FieldSetRule,
					},
					w.Msg{
						Icon: "shield",
						Parts: []w.MsgPart{
							w.Strong("What only works from outside."),
							w.Text(" A local Carrel is not visible from the internet, so "),
							w.Strong("published links"),
							w.Text(", "),
							w.Strong("devices over davloom"),
							w.Text(", "),
							w.Strong("backup started by an external request"),
							w.Text(" and "),
							w.Strong("invitation links"),
							w.Text(" either do not work or need you to arrange the access yourself — an invitation letter will go out, but the address inside it opens on your machine only — port forwarding, a tunnel, your own domain and certificate. Carrel neither guesses nor configures any of that: it marks such places and shows what is not answering."),
						},
						Tone:  w.MsgAlert,
						Space: w.MsgFlush,
					},
					w.Line{
						Items: []w.LinePart{
							w.Button{Label: "Run", Kind: w.ButtonPrimary},
							w.Button{Label: "Back"},
							w.Dim{
								Parts:   []w.Inline{w.Dot{Tone: w.DotAccent}, w.Text("server stopped")},
								Size:    w.Dim11_5,
								Variant: w.DimWrap,
								Place:   w.Pushed,
							},
						},
						Gap:   w.LineGap8,
						Block: true,
					},
				},
				Gap: w.WindowGap13,
			},
		},
	}, "desktop-6", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Window{
				Items: []w.WindowPart{
					w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo42},
					w.WindowNote{
						Text: "The app needs the address of your Carrel instance. The data stays there: the app is a window, not a second copy.",
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Address",
								Parts: []w.FieldPart{w.Input{Value: "https://carrel.example", Label: "Address", Mono: true, Focus: true}},
								Width: w.FieldW100,
							},
							w.Field{Label: "Login", Parts: []w.FieldPart{w.Input{Value: "mix", Label: "Login"}}, Width: w.FieldW100},
							w.Field{
								Label: "Password",
								Parts: []w.FieldPart{w.Input{Value: "••••••••••", Label: "Password"}},
								Width: w.FieldW100,
							},
							w.Check{On: true, Parts: []w.CheckPart{w.Text("Start at sign-in")}},
						},
						Variant: w.FormColumnGap10,
					},
					w.Diag{
						Text: "[connect]  GET /healthz → 200\n[tls]      Let's Encrypt certificate, valid until 2026-11-02\n[version]  0.8.0 · compatible",
					},
					w.Line{
						Items: []w.LinePart{
							w.Button{Label: "Connect", Kind: w.ButtonPrimary},
							w.Button{Label: "Check"},
							w.Push{Items: []w.PushPart{w.Text("Add a second instance")}, Face: faceOf(w.FaceAccent11_5)},
						},
						Gap:   w.LineGap8,
						Block: true,
					},
				},
				Gap: w.WindowGap14,
			},
		},
	}, "desktop-7", "tc-root", 1),
	page("on-background", w.Page{
		Items: []w.PagePart{
			w.Window{
				Items: []w.WindowPart{
					w.Area{Text: "Спросить у Виктора про миграцию — успеет ли до заморозки.", Variant: w.AreaTall, Focus: true, Label: "Text"},
					w.Line{
						Items: []w.LinePart{
							w.Pick{Value: "nextcloud · Записи", Variant: w.PickT12Fill},
							w.Button{Label: "Save", Kind: w.ButtonPrimary},
							w.Kbd{Keys: "Ctrl+Enter"},
						},
						Gap:   w.LineGap8,
						Block: true,
					},
					w.Dim{
						Parts: []w.Inline{w.Text("The window closes itself. Carrel need not be open — the note goes to the server at once.")},
						Size:  w.Dim11,
					},
				},
				Gap:   w.WindowGap9,
				Small: true,
			},
		},
		Variant: w.PageOnBackground,
	}, "desktop-9", "tc-root", 1),
	page("on-background", w.Page{
		Items: []w.PagePart{
			w.Window{
				Items: []w.WindowPart{
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
					w.Line{
						Items: []w.LinePart{
							w.Pick{Value: "nextcloud · Личный", Variant: w.PickT12Fill},
							w.Button{Label: "Save", Kind: w.ButtonPrimary},
							w.Kbd{Keys: "Enter"},
						},
						Gap:   w.LineGap8,
						Block: true,
					},
					w.Dim{
						Parts: []w.Inline{
							w.Text("The same sheet as in the window, in a window of its own. Carrel need not be open — the event goes to the server at once."),
						},
						Size: w.Dim11,
					},
				},
				Gap:   w.WindowGap9,
				Small: true,
			},
		},
		Variant: w.PageOnBackground,
	}, "desktop-10", "tc-root", 1),
	page("on-background-column", w.Page{
		Items: []w.PagePart{
			w.Form{
				Items: []w.FormPart{
					w.Rubric{Text: "Reminders"},
					w.Field{
						Label: "Remind me",
						Parts: []w.FieldPart{w.Pick{Value: "10 minutes · and as set in the event"}},
						Width: w.FieldW100,
					},
					w.Check{On: true, Parts: []w.CheckPart{w.Text("Task due dates — on the morning they are due")}},
					w.Check{Parts: []w.CheckPart{w.Text("Show the title on the lock screen")}},
					w.Field{
						Label: "Quiet hours",
						Parts: []w.FieldPart{w.Input{Value: "22:00 – 08:00", Label: "Quiet hours", Mono: true}},
						Width: w.FieldW100,
					},
					w.Field{
						Label: "From which collections",
						Parts: []w.FieldPart{w.Input{Value: "3 of 7 ticked", Label: "From which collections"}},
						Width: w.FieldW100,
					},
				},
				Variant: w.FormColumnGap10Mid,
			},
		},
		Variant: w.PageOnBackgroundColumn,
	}, "desktop-11", "tc-root", 1),
	page("on-background", w.Page{
		Items: []w.PagePart{
			w.Window{
				Items: []w.WindowPart{
					w.Msg{
						Parts: []w.MsgPart{
							w.Bold{Parts: []w.Inline{w.Text("Sign out of "), w.Mono{Text: "carrel.example"}, w.Text("?")}},
							w.Text(" The app goes back to the first screen, where the mode is chosen again — a server or a local run. "),
							w.Strong("Your other devices stay signed in"),
							w.Text(" — those sessions end in Settings → Account, and they end there whether this app is running or not."),
						},
						Space: w.MsgFlush,
					},
					w.Msg{
						Icon: "shield",
						Parts: []w.MsgPart{
							w.Strong("Delete what this app keeps on this computer?"),
							w.Text(" The data folder "),
							w.Mono{Text: "C:\\Users\\mix\\AppData\\Roaming\\Carrel"},
							w.Text(", the instance address and name, the start-at-sign-in setting and the session token in the keychain. The next start would be indistinguishable from the first, and no one can bring the saved connections back — not even from here."),
						},
						Tone:  w.MsgAlert,
						Space: w.MsgFlush,
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("Your events, contacts and notes are not here and are not touched: they live on your DAV servers. What would be deleted is the key to them, not them."),
						},
						Space: w.MsgFlush,
					},
					w.Line{
						Items: []w.LinePart{
							w.Button{Label: "Keep and sign out", Kind: w.ButtonPrimary},
							w.Kbd{Keys: "Enter"},
							w.Button{Label: "Delete and sign out", Kind: w.ButtonDangerSolid},
							w.Line{Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGap7, Place: w.Pushed},
						},
						Gap:   w.LineGap8,
						Block: true,
					},
				},
				Gap:   w.WindowGap10,
				Small: true,
			},
		},
		Variant: w.PageOnBackground,
	}, "desktop-12", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Msg{
				Icon: "down",
				Parts: []w.MsgPart{
					w.Strong("Carrel 0.8.0 is ready to install."),
					w.Text(" It will be applied the next time the app closes, or now — the window reopens where it was."),
				},
				End: &w.Push{
					Items: []w.PushPart{
						w.Button{Label: "Install and restart", Kind: w.ButtonPrimary},
						w.Button{Label: "What’s new"},
						w.Button{Label: "Later"},
					},
				},
				Tone:  w.MsgAccent,
				Space: w.MsgFlushEnd,
			},
			w.Window{
				Items: []w.WindowPart{
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Installed",
								Parts: []w.FieldPart{
									w.Input{
										Value:   "app 0.8.0 · server 0.8.0 — both run inside this app",
										Label:   "Installed",
										Mono:    true,
										Variant: w.Input11_5,
									},
									w.FieldHint{
										Parts: []w.Inline{
											w.Text("In local mode the server is part of the app, so one update covers both. Your data folder is untouched by an update: connections and settings stay where they are."),
										},
									},
								},
								Place: w.Grown,
							},
							w.Field{
								Label: "Downloaded",
								Parts: []w.FieldPart{
									w.Line{
										Items: []w.LinePart{
											w.Meter{Percent: 100, Width: w.MeterFill},
											w.Mono{Text: "0.8.0 · 14.2 MB · signature checked", Variant: w.Mono11Dusty},
										},
										Gap: w.LineGap9,
									},
								},
								Place: w.Grown,
							},
						},
						Variant: w.FormGap12,
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("Nothing installs itself while you work: an update waits for the app to close, or for this button. You are seeing this banner because automatic checking was turned on — out of the box it is off, and the app never goes looking."),
						},
						Space: w.MsgFlush,
					},
				},
				Gap: w.WindowGap12,
			},
		},
	}, "desktop-14", "tc-root", 1),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Window{
				Items: []w.WindowPart{
					w.Msg{
						Icon: "shield",
						Parts: []w.MsgPart{
							w.Strong("The app is newer than the server."),
							w.Text(" This app is 0.8.0, "),
							w.Mono{Text: "carrel.example"},
							w.Text(" runs 0.6.4. Everything that both versions know still works; what is newer than the server simply is not there."),
						},
						Tone:  w.MsgAlert,
						Space: w.MsgFlush,
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "What to do",
								Parts: []w.FieldPart{
									w.Input{Placeholder: "Update the server where it runs — this app cannot do it", Label: "What to do"},
									w.FieldHint{
										Parts: []w.Inline{
											w.Text("The server is somebody’s machine: yours, a friend’s, a hosting account. An app that could update it from the outside would be a way in, and Carrel does not build such ways."),
										},
									},
								},
								Place: w.Grown,
							},
						},
						Variant: w.FormGap12,
					},
					w.Line{
						Items: []w.LinePart{w.Button{Label: "How to update a server"}, w.Button{Label: "Do not warn about this server"}},
						Gap:   w.LineGap8,
						Block: true,
					},
				},
				Gap: w.WindowGap12,
			},
		},
	}, "desktop-15", "tc-root", 1),
	page("print", w.Page{
		Items: []w.PagePart{
			w.PrintTitle{Text: "Documents / carrel / attachments"},
			w.PrintSources{Text: "sftpgo · vps · 14 items · 38.2 MB"},
			w.PrintRubric{Text: "Folders"},
			w.PrintRow{
				Items:   []w.PrintRowPart{w.PrintCell{Name: "2026-07"}, w.PrintNum{Text: "34 items"}, w.PrintNum{Text: "2026-07-01"}},
				Columns: true,
			},
			w.PrintRubric{Text: "Files"},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.PrintCell{Name: "2026-08-10-sync-board.png"},
					w.PrintNum{Text: "412 KB"},
					w.PrintNum{Text: "2026-08-10"},
				},
				Columns: true,
			},
			w.PrintRow{
				Items:   []w.PrintRowPart{w.PrintCell{Name: "release-checklist.md"}, w.PrintNum{Text: "3 KB"}, w.PrintNum{Text: "2026-08-09"}},
				Columns: true,
			},
			w.PrintRow{
				Items:   []w.PrintRowPart{w.PrintCell{Name: "смета-подрядчик.pdf"}, w.PrintNum{Text: "1.2 MB"}, w.PrintNum{Text: "2026-08-05"}},
				Columns: true,
			},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.PrintCell{Name: "2026-08-04-whiteboard.jpg"},
					w.PrintNum{Text: "2.8 MB"},
					w.PrintNum{Text: "2026-08-04"},
				},
				Columns: true,
			},
			w.PrintFoot{
				Text: "Carrel · folder · sftpgo / Documents / carrel / attachments · captured 2026-08-16 20:41 · 14 of 14 shown",
			},
		},
		Variant: w.PagePrint,
	}, "desktop-16", "tc-root", 1),
)

func init() {
	Widgets = append(Widgets, w.Page{})
	Parts = append(Parts, w.Window{}, w.WindowNote{}, w.PrintTitle{}, w.PrintSources{}, w.PrintRubric{}, w.PrintRow{},
		w.PrintCell{}, w.PrintNum{}, w.PrintFoot{})
}
