// Раздел макета: 12. Настольное приложение.

package screens

import w "carrel.dev/typecase/widget"

// Desktop1 is frame «Окно приложения» (desktop-1).
func Desktop1() Screen {
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
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
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
								Strong: true,
								Detail: "Переговорная 2",
								Badges: []w.Badge{w.Badge{Label: "4 attendees"}, w.Badge{Label: "repeats weekly"}},
								Source: "nextcloud · Личный",
								On:     true,
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
			Side: &w.Side{
				Items: []w.SidePart{
					w.Head{
						Title: &w.HeadTitle{Text: "Синк-ап по релизу", Size: w.HeadTitle17},
						Subs: []w.HeadSub{
							w.HeadSub{Parts: []w.HeadSubPart{w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, w.Text("nextcloud · Личный")}},
						},
						Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Edit", Small: true}, w.Close{Label: "Close the panel"}}},
						Pad:  w.HeadPadBottom8,
					},
					w.Fields{
						Items: []w.FieldsItem{
							{Term: "When", Value: []w.ValuePart{w.Mono{Text: "10:00 – 11:00"}}},
							{Term: "Where", Value: []w.ValuePart{w.Text("Переговорная 2")}},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Attachments"},
							w.SecOpen{
								Items: []w.SecOpenPart{
									w.SecOpened{Name: "release-checklist.md", Size: w.Dim{Parts: []w.Inline{w.Text("· 3.1 KB")}}},
									w.Line{
										Items: []w.LinePart{w.Button{Label: "Open in editor", Small: true}, w.Button{Label: "Show in folder", Small: true}},
										Gap:   w.LineGap6,
									},
									w.Dim{
										Parts: []w.Inline{
											w.Text("Edits go to the server when the program closes the file. The temporary copy lives until then and is deleted."),
										},
										Size: w.Dim11,
									},
								},
							},
						},
					},
				},
			},
			Height: w.ShellHeight390,
		},
	}
}

// Desktop3 is frame «Установка · кому эта копия» (desktop-3).
func Desktop3() Screen {
	return w.Page{
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
	}
}

// Desktop4 is frame «Portable · чем эта копия отличается» (desktop-4).
func Desktop4() Screen {
	return w.Page{
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
	}
}

// Desktop5 is frame «Первый запуск · как работать» (desktop-5).
func Desktop5() Screen {
	return w.Page{
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
	}
}

// Desktop6 is frame «Локальный запуск» (desktop-6).
func Desktop6() Screen {
	return w.Page{
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
									w.FieldHint{Parts: []w.Inline{w.Text("Your connections and settings, encrypted. Records stay on the servers.")}},
								},
								Place: w.Grown,
							},
							w.Field{
								Label: "Port",
								Parts: []w.FieldPart{
									w.Input{Value: "chosen at every start · now 51734", Label: "Port", Mono: true},
									w.FieldHint{Parts: []w.Inline{w.Text("A free one at every start; its number is in the status bar.")}},
								},
								Width: w.FieldFill300,
							},
							w.Field{Parts: []w.FieldPart{w.Check{On: true, Parts: []w.CheckPart{w.Text("Stop the server on exit")}}}},
							w.Field{
								Parts: []w.FieldPart{
									w.Check{Parts: []w.CheckPart{w.Text("Allow other devices on this network")}},
									w.FieldHint{
										Parts: []w.Inline{
											w.Text("Turning it on "),
											w.Strong("pins the port"),
											w.Text(": a phone and DAVx5 reach "),
											w.Mono{Text: "http://192.168.1.14:8765"},
											w.Text(" with your password."),
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
									w.Text("Nobody can recover the password, not even this app; without it the saved connections are lost."),
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
							w.Strong("Not reachable from the internet."),
							w.Text(" "),
							w.Strong("Published links"),
							w.Text(", "),
							w.Strong("devices over davloom"),
							w.Text(", "),
							w.Strong("external backup triggers"),
							w.Text(" and "),
							w.Strong("invitation links"),
							w.Text(" work on this machine only, until you set up port forwarding, a tunnel or your own domain."),
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
	}
}

// Desktop7 is frame «Первый запуск · к какому экземпляру» (desktop-7).
func Desktop7() Screen {
	return w.Page{
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
	}
}

// Desktop9 is frame «Заметка из любой программы» (desktop-9).
func Desktop9() Screen {
	return w.Page{
		Items: []w.PagePart{
			w.Window{
				Items: []w.WindowPart{
					w.Area{
						Text:    "Спросить у Виктора про миграцию — успеет ли до заморозки.",
						Variant: w.AreaTall,
						Focus:   true,
						Label:   "Text",
					},
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
	}
}

// Desktop10 is frame «Встреча из трея» (desktop-10).
func Desktop10() Screen {
	return w.Page{
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
	}
}

// Desktop11 is frame «Напоминания» (desktop-11).
func Desktop11() Screen {
	return w.Page{
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
	}
}

// Desktop12 is frame «Выход из учётной записи · удалять ли данные» (desktop-12).
func Desktop12() Screen {
	return w.Page{
		Items: []w.PagePart{
			w.Window{
				Items: []w.WindowPart{
					w.Msg{
						Parts: []w.MsgPart{
							w.Bold{Parts: []w.Inline{w.Text("Sign out of "), w.Mono{Text: "carrel.example"}, w.Text("?")}},
							w.Text(" "),
							w.Strong("Your other devices stay signed in"),
							w.Text(": end those sessions in Settings → Account."),
						},
						Space: w.MsgFlush,
					},
					w.Msg{
						Icon: "shield",
						Parts: []w.MsgPart{
							w.Strong("Delete what this app keeps on this computer?"),
							w.Text(" The folder "),
							w.Mono{Text: "C:\\Users\\mix\\AppData\\Roaming\\Carrel"},
							w.Text(", the instance address and the session token. The saved connections cannot be brought back; records on your DAV servers are not touched."),
						},
						Tone:  w.MsgAlert,
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
	}
}

// Desktop13 is frame «Настройки приложения» (desktop-13).
func Desktop13() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Application", Size: w.HeadTitle19},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Account and collection settings are inside, in Settings")}}},
						Pad:   w.HeadPadBottom6,
					},
					w.Form{
						Items: []w.FormPart{
							w.FieldSet{
								Rubric: w.Rubric{Text: "Connection"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Signed in as",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Input{Value: "mix · https://carrel.example", Label: "Signed in as", Mono: true, Variant: w.InputFill},
													w.Button{Label: "Sign out…"},
												},
												Gap: w.LineGap8,
											},
											w.FieldHint{Parts: []w.Inline{w.Text("Nothing is deleted: the app returns to the first screen.")}},
										},
										Place: w.Grown,
									},
								},
								Variant: w.FieldSetLead,
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Behaviour"},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Check{On: true, Parts: []w.CheckPart{w.Text("Start at sign-in")}},
											w.Check{On: true, Parts: []w.CheckPart{w.Text("Closing the window hides to the tray")}},
										},
										Gap:   8,
										Place: w.Grown,
									},
									w.Field{
										Label: "Global note shortcut",
										Parts: []w.FieldPart{
											w.Input{Value: "Ctrl+Shift+N", Label: "Global note shortcut", Mono: true},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("The note opens on the screen the pointer is on; the main window comes back where it was left."),
												},
											},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Downloads folder",
										Parts: []w.FieldPart{w.Input{Value: "C:\\Users\\mix\\Downloads", Label: "Downloads folder", Mono: true}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Open files with",
										Parts: []w.FieldPart{w.Pick{Value: "default application"}},
										Width: w.FieldMD,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Network"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Proxy", Parts: []w.FieldPart{w.Pick{Value: "system"}}, Width: w.FieldMD},
									w.Field{
										Parts: []w.FieldPart{w.Check{Parts: []w.CheckPart{w.Text("Trust this instance’s self-signed certificate")}}},
									},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Appears only after a certificate was rejected; it trusts that one fingerprint, checking stays on."),
												},
											},
										},
										Place: w.Grown,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Updates"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Installed",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Input{Value: "app 0.8.0 · server 0.8.0", Label: "Installed", Mono: true, Variant: w.Input11_5Fill},
													w.Button{Label: "Check now"},
												},
												Gap: w.LineGap8,
											},
											w.FieldHint{
												Parts: []w.Inline{w.Text("Last checked 2 hours ago: "), w.Mono{Text: "0.8.0"}, w.Text(" was the latest.")},
											},
										},
										Width: w.FieldFill300,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.Check{Parts: []w.CheckPart{w.Text("Check for updates automatically")}},
											w.Check{Parts: []w.CheckPart{w.Text("Install when I close the app")}},
											w.Check{Parts: []w.CheckPart{w.Text("Include pre-releases")}},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("With these clear, the app does not go online. An update is verified by its signature; without the second tick it waits until you install it."),
												},
											},
										},
										Gap:   7,
										Place: w.Grown,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{w.Text("In local mode this updates the server too; with your own server, the app only.")},
											},
										},
										Place: w.Grown,
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

// Desktop14 is frame «Обновление · скачано и ждёт» (desktop-14).
func Desktop14() Screen {
	return w.Page{
		Items: []w.PagePart{
			w.Msg{
				Icon: "download",
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
	}
}

// Desktop15 is frame «Обновление · сервер отстал» (desktop-15).
func Desktop15() Screen {
	return w.Page{
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
	}
}

// Desktop16 is frame «Папка на бумаге» (desktop-16).
func Desktop16() Screen {
	return w.Page{
		Items: []w.PagePart{
			w.PrintTitle{Text: "Documents / carrel / attachments"},
			w.PrintSources{Text: "sftpgo · vps · 14 items · 38.2 MB"},
			w.PrintRow{
				Items:   []w.PrintRowPart{w.PrintCell{Name: "Name"}, w.PrintNum{Text: "Size"}, w.PrintNum{Text: "Changed"}},
				Columns: true,
				Head:    true,
			},
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
				Items: []w.PrintRowPart{
					w.PrintCell{Name: "release-checklist.md"},
					w.PrintNum{Text: "3.1 KB"},
					w.PrintNum{Text: "2026-08-09"},
				},
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
	}
}
