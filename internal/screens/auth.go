// Раздел макета: 7.20 Вход и первый запуск.

package screens

import w "carrel.dev/typecase/widget"

// Auth1 is frame «Вход» (s-auth-1).
func Auth1() Screen {
	return w.Page{
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
	}
}

// Auth13 is frame «Вход · неверный пароль» (s-auth-13).
func Auth13() Screen {
	return w.Page{
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
										w.Field{Label: "Login", Parts: []w.FieldPart{w.Input{Value: "mix", Label: "Login"}}, Width: w.FieldW100},
										w.Field{
											Label: "Password",
											Parts: []w.FieldPart{
												w.Input{Label: "Password", Focus: true, Error: true},
												w.FieldHint{Parts: []w.Inline{w.Text("Login or password is wrong.")}, Variant: w.HintDanger},
											},
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
	}
}

// Auth2 is frame «Регистрация и подтверждение» (s-auth-2).
func Auth2() Screen {
	return w.Page{
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
	}
}

// Auth12 is frame «Приглашение · свой логин и пароль» (s-auth-12).
func Auth12() Screen {
	return w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogTitle{Text: "Join carrel.example", Gap: w.DialogTitleGap4},
						w.DialogSub{Parts: []w.Inline{w.Text("mix invited you. Choose a login and a password.")}},
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
									Parts: []w.FieldPart{w.Input{Value: "anna", Label: "Login", Focus: true}},
									Width: w.FieldMD,
								},
								w.Field{
									Label: "Email",
									Parts: []w.FieldPart{
										w.Input{Value: "anna@example.org", Label: "Email"},
										w.FieldHint{
											Parts: []w.Inline{w.Text("The address the invitation went to. Another one is confirmed by a letter.")},
										},
									},
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
							Items:   []w.FormFootPart{w.Button{Label: "Create account", Kind: w.ButtonPrimary}},
							Variant: w.FormFootFlush,
						},
					},
					Variant: w.DialogPage440,
				},
			},
		},
	}
}

// Auth14 is frame «Приглашение · по ссылке» (s-auth-14).
func Auth14() Screen {
	return w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogTitle{Text: "Join carrel.example", Gap: w.DialogTitleGap4},
						w.DialogSub{Parts: []w.Inline{w.Text("mix invited you. Choose a login and a password.")}},
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
									Parts: []w.FieldPart{w.Input{Value: "anna", Label: "Login", Focus: true}},
									Width: w.FieldMD,
								},
								w.Field{
									Label: "Email",
									Parts: []w.FieldPart{
										w.Input{Placeholder: "you@example.org", Label: "Email"},
										w.FieldHint{Parts: []w.Inline{w.Text("Confirmed by a letter.")}},
									},
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
							Items:   []w.FormFootPart{w.Button{Label: "Create account", Kind: w.ButtonPrimary}},
							Variant: w.FormFootFlush,
						},
					},
					Variant: w.DialogPage440,
				},
			},
		},
	}
}

// Auth3 is frame «Письмо ушло» (s-auth-3).
func Auth3() Screen {
	return w.Page{
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
	}
}

// Auth4 is frame «Первый запуск» (s-auth-4).
func Auth4() Screen {
	return w.Page{
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
	}
}

// Auth5 is frame «Забыли пароль · логин или почта» (s-auth-5).
func Auth5() Screen {
	return w.Page{
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
	}
}

// Auth6 is frame «Забыли пароль · что можно сделать» (s-auth-6).
func Auth6() Screen {
	return w.Page{
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
												w.Text("This account deposited its encryption key. The administrator can recover it with the instance master password, and your data is kept."),
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
	}
}

// Auth7 is frame «Забыли пароль · подтверждённой почты нет» (s-auth-7).
func Auth7() Screen {
	return w.Page{
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
	}
}

// Auth8 is frame «Забыли пароль · письмо отправлено» (s-auth-8).
func Auth8() Screen {
	return w.Page{
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
	}
}

// Auth9 is frame «Забыли пароль · запрос администратору» (s-auth-9).
func Auth9() Screen {
	return w.Page{
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
	}
}

// Auth10 is frame «Забыли пароль · подтверждение очистки» (s-auth-10).
func Auth10() Screen {
	return w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogTitle{Text: "Erase everything and start again?", Gap: w.DialogTitleGap8},
						w.Prose{
							Parts: []w.Inline{
								w.Text("This is not a password recovery. The key to your data comes from the password you forgot, so nothing stored under it can be read again — by you, by us, by anyone.\n\nConfirming erases what Carrel keeps encrypted for "),
								w.Strong("mix"),
								w.Text(": saved connections and their passwords, subscriptions, published links, devices, external backup triggers and the deposited copy of the key. Address books, calendars, tasks and notes on the servers are not touched — you connect to them again afterwards.\n\nThe account itself stays: same login, role and confirmed address, the same first day of the week and clock. Only the password is new, and you set it on the next screen."),
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
	}
}

// Auth11 is frame «Забыли пароль · новый пароль» (s-auth-11).
func Auth11() Screen {
	return w.Page{
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
	}
}
