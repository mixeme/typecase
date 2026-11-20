package refs

import w "carrel.dev/typecase/widget"

// Pages without a shell: sign-in, «about», a desktop
// window's inside, a page on paper. A record per page frame: the root's own
// tree is the page's whole content, so every frame is a form of its own.

func page(state string, fixture Renderer, frame, class string) Entry {
	return e("pages", state, fixture, ref(frame, class))
}

var _ = add(
	page("print", w.Page{Items: []w.PagePart{w.PrintTitle{Text: "Notes · Combined notebook"}, w.PrintSources{Text: "nextcloud · Записи, Дневник"}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "Date", Variant: w.MonoFace}, w.PrintCell{Name: "Note"}}, Head: true}, w.PrintRubric{Text: "August 2026"}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "10 Aug", Variant: w.MonoFace}, w.PrintCell{Name: "Заметка со встречи 10.08", Detail: "Договорились заморозить ветку в среду; письмо пользователям пишет Анна… · in 2 notebooks"}}}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "8 Aug", Variant: w.MonoFace}, w.PrintCell{Name: "Что почитать про CRDT", Detail: "Кляйн, глава 5; статья Шапиро 2011; демка automerge… · in 2 notebooks"}}}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "3 Aug", Variant: w.MonoFace}, w.PrintCell{Name: "Про дорогу до дачи", Detail: "Съезд после заправки, потом грунтовка два километра… · nextcloud · Дневник"}}}, w.PrintRubric{Text: "July 2026"}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "21 Jul", Variant: w.MonoFace}, w.PrintCell{Name: "Про сроки по сборнику", Detail: "Анна просит две недели на вёрстку после сдачи текста. · nextcloud · Записи"}}}, w.PrintFoot{Text: "Carrel · notes · 4 notes · captured 2026-08-15 09:41 · 2 of 2 notebooks"}}, Variant: w.PagePrint}, "print-8", "tc-root"),
	page("print", w.Page{Items: []w.PagePart{w.PrintTitle{Text: "Files · Documents / carrel / attachments / 2026-07"}, w.PrintSources{Text: "sftpgo · vps · Documents"}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Name"}, w.PrintNum{Text: "Size"}, w.PrintNum{Text: "Changed"}}, Columns: true, Head: true}, w.PrintRubric{Text: "Folders"}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "архив"}, w.PrintNum{Text: "—"}, w.PrintNum{Text: "2026-07-01 09:20"}}, Columns: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "черновики"}, w.PrintNum{Text: "—"}, w.PrintNum{Text: "2026-07-29 16:40"}}, Columns: true}, w.PrintRubric{Text: "Files"}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "whiteboard.jpg"}, w.PrintNum{Text: "2.8 MB"}, w.PrintNum{Text: "2026-07-14 17:05"}}, Columns: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "панорама.tif"}, w.PrintNum{Text: "84 MB"}, w.PrintNum{Text: "2026-07-21 12:30"}}, Columns: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "смета.pdf"}, w.PrintNum{Text: "1.2 MB"}, w.PrintNum{Text: "2026-07-08 10:15"}}, Columns: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "схема-1.png"}, w.PrintNum{Text: "640 KB"}, w.PrintNum{Text: "2026-07-16 11:02"}}, Columns: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "схема-2.png"}, w.PrintNum{Text: "712 KB"}, w.PrintNum{Text: "2026-07-16 11:04"}}, Columns: true}, w.PrintFoot{Text: "Carrel · files · 7 items · 89.4 MB · captured 2026-08-15 09:41"}}, Variant: w.PagePrint}, "print-6", "tc-root"),
	page("print", w.Page{Items: []w.PagePart{w.PrintTitle{Text: "Личный · 10 – 16 August 2026"}, w.PrintSources{Text: "nextcloud · Личный"}, w.PrintRubric{Text: "Monday · 10 August"}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "10:00", Variant: w.MonoFace}, w.PrintCell{Name: "Синк-ап по релизу", Detail: "Переговорная 2 · 4 attendees · repeats weekly"}}}, w.PrintRubric{Text: "Tuesday · 11 August"}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "18:00", Variant: w.MonoFace}, w.PrintCell{Name: "Забрать документы", Detail: "МФЦ, Мясницкая 13"}}}, w.PrintRubric{Text: "Wednesday · 12 August"}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "all day", Variant: w.MonoFace}, w.PrintCell{Name: "Поездка в Казань", Detail: "day 1 of 3"}}}, w.PrintRubric{Text: "Thursday · 13 August"}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "all day", Variant: w.MonoFace}, w.PrintCell{Name: "Поездка в Казань", Detail: "day 2 of 3"}}}, w.PrintRubric{Text: "Friday · 14 August"}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "all day", Variant: w.MonoFace}, w.PrintCell{Name: "Поездка в Казань", Detail: "day 3 of 3"}}}, w.PrintFoot{Text: "Carrel · agenda · Личный · 2026-08-10 – 2026-08-16 · captured 2026-08-15 09:41"}}, Variant: w.PagePrint}, "print-7", "tc-root"),
	page("print", w.Page{Items: []w.PagePart{w.PrintTitle{Text: "Заметка со встречи 10.08"}, w.PrintSources{Text: "nextcloud · Записи · 10 August 2026, 11:05"}, w.PrintText{Text: "Договорились заморозить ветку в среду."}, w.PrintText{Text: "Письмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт."}, w.PrintText{Text: "Что делаем", Heading: true}, w.PrintText{Text: "проверить changelog", Tick: w.PrintTickOpen}, w.PrintText{Text: "попросить у Артёма скриншоты", Tick: w.PrintTickDone}, w.PrintText{Text: "Перенос старых вложений в новую папку — Виктор посмотрит, сколько их вообще, и скажет в четверг."}, w.PrintRubric{Text: "Attachments"}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Name"}, w.PrintNum{Text: "Size"}, w.PrintNum{Text: "Where"}}, Columns: true, Head: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "2026-08-10-sync-board.png"}, w.PrintNum{Text: "412 KB"}, w.PrintNum{Text: "in the note"}}, Columns: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "запись-встречи.m4a"}, w.PrintNum{Text: "24 MB"}, w.PrintNum{Text: "Documents / carrel"}}, Columns: true}, w.PrintRubric{Text: "Kept as the server sent them"}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "DESCRIPTION"}, w.PrintCell{Name: "Договорились заморозить ветку в четверг."}}, Field: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "X-JTX-COLOR"}, w.PrintCell{Name: "-16711936"}}, Field: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "X-JTX-PINNED"}, w.PrintCell{Name: "1"}}, Field: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "X-OTHER"}, w.PrintCell{Name: "kept"}}, Field: true}, w.PrintFoot{Text: "Carrel · note · captured 2026-08-15 09:41"}}, Variant: w.PagePrint}, "print-5", "tc-root"),
	page("print", w.Page{Items: []w.PagePart{w.PrintTitle{Text: "Tasks · Open · Combined list"}, w.PrintSources{Text: "nextcloud · Дела, Работа"}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "Due", Variant: w.MonoFace}, w.PrintCell{Name: "Task"}}, Head: true}, w.PrintRubric{Text: "Overdue"}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "7 Aug", Variant: w.MonoFace}, w.PrintCell{Name: "Отправить документы в МФЦ", Detail: "nextcloud · Дела · overdue by 3 days", Tick: w.PrintTickOpen}}}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "8 Aug", Variant: w.MonoFace}, w.PrintCell{Name: "Ответить подрядчику по смете", Detail: "nextcloud · Работа · 40%", Tick: w.PrintTickOpen}}}, w.PrintRubric{Text: "This week"}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "12 Aug", Variant: w.MonoFace}, w.PrintCell{Name: "Собрать чемодан", Detail: "18:00 · nextcloud · Дела, Работа", Tick: w.PrintTickOpen}}}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "14 Aug", Variant: w.MonoFace}, w.PrintCell{Name: "Дописать раздел про кэш", Detail: "nextcloud · Работа · 70%", Tick: w.PrintTickOpen}}}, w.PrintFoot{Text: "Carrel · tasks · Open · 4 of 7 · captured 2026-08-15 09:41 · 2 of 2 lists"}}, Variant: w.PagePrint}, "print-4", "tc-root"),
	page("print", w.Page{Items: []w.PagePart{w.PrintTitle{Text: "Анна Ковалёва"}, w.PrintSources{Text: "nextcloud · Контакты"}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Organisation"}, w.PrintCell{Name: "Издательство «Слово»"}}, Field: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Title"}, w.PrintCell{Name: "Выпускающий редактор"}}, Field: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Phone · mobile"}, w.PrintCell{Name: "+7 912 445-10-88"}}, Field: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Phone · work"}, w.PrintCell{Name: "+7 495 120-45-00"}}, Field: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Email · work"}, w.PrintCell{Name: "a.kovaleva@example.org"}}, Field: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Email · second"}, w.PrintCell{Name: "anna.kovaleva@slovo.example"}}, Field: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Address · work"}, w.PrintCell{Name: "Москва, Тверская 12, стр. 3"}}, Field: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Birthday"}, w.PrintCell{Name: "1987-03-14"}}, Field: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Tags"}, w.PrintCell{Name: "работа, издательство"}}, Field: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Note"}, w.PrintCell{Name: "Пишет по вторникам, отвечает в тот же день."}}, Field: true}, w.PrintRubric{Text: "Kept as the server sent them"}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "X-EVOLUTION-FILE-AS"}, w.PrintCell{Name: "Ковалёва, Анна"}}, Field: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "X-ABLABEL"}, w.PrintCell{Name: "издательство"}}, Field: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "X-CARREL-UNKNOWN-FLAG"}, w.PrintCell{Name: "1"}}, Field: true}, w.PrintFoot{Text: "Carrel · contact · captured 2026-08-15 09:41"}}, Variant: w.PagePrint}, "print-3", "tc-root"),
	page("print", w.Page{Items: []w.PagePart{w.PrintTitle{Text: "Contacts · Combined address book"}, w.PrintSources{Text: "nextcloud · Контакты, Архив · radicale · vps · Команда"}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Name"}, w.PrintNum{Text: "Phone"}, w.PrintNum{Text: "Email"}}, Columns: true, Head: true}, w.PrintRubric{Text: "А"}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Алексей Смирнов", Detail: "nextcloud · Контакты"}, w.PrintNum{Text: "+7 916 200-71-15"}, w.PrintNum{Text: "—"}}, Columns: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Анна Ковалёва", Detail: "Издательство «Слово» · in 3 books"}, w.PrintNum{Text: "+7 912 445-10-88"}, w.PrintNum{Text: "a.kovaleva@example.org"}}, Columns: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Артём Белов", Detail: "Тензор · in 2 books"}, w.PrintNum{Text: "+7 903 118-22-04"}, w.PrintNum{Text: "a.belov@example.org"}}, Columns: true}, w.PrintRubric{Text: "В"}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Вера Лебедева", Detail: "nextcloud · Архив"}, w.PrintNum{Text: "+7 905 331-09-42"}, w.PrintNum{Text: "vera@example.net"}}, Columns: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Виктор Гончаров", Detail: "Городская библиотека №4 · nextcloud · Контакты"}, w.PrintNum{Text: "—"}, w.PrintNum{Text: "v.goncharov@example.org"}}, Columns: true}, w.PrintRubric{Text: "Л"}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Лебедева Вера", Detail: "Библиотека №4 · nextcloud · Контакты"}, w.PrintNum{Text: "+7 905 331-09-42"}, w.PrintNum{Text: "v.lebedeva@example.org"}}, Columns: true}, w.PrintRubric{Text: "М"}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Михаил Енученко", Detail: "nextcloud · Контакты"}, w.PrintNum{Text: "—"}, w.PrintNum{Text: "mix@example.org"}}, Columns: true}, w.PrintFoot{Text: "Carrel · contacts · 7 cards · captured 2026-08-15 09:41 · 3 of 3 books"}}, Variant: w.PagePrint}, "print-2", "tc-root"),
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
	}, "s-auth-1", "tc-root"),
	page("plain", w.Page{Items: []w.PagePart{w.Auth{Dialog: w.Dialog{Items: []w.DialogPart{
		w.DialogTitle{Text: "Create an account", Gap: w.DialogTitleGapTight},
		w.DialogSub{Parts: []w.Inline{w.Text("A confirmation link will be sent to the address you give; the account cannot be used until you open it.")}},
		w.Msg{Parts: []w.MsgPart{w.Text("Nobody can recover this password, this server included.")}, Tone: w.MsgAlert, Space: w.MsgFlushBelowWide},
		w.Form{Items: []w.FormPart{w.Field{Label: "Login", Parts: []w.FieldPart{w.Input{Value: "viktor", Label: "Login", Focus: true}}, Width: w.FieldMD}, w.Field{Label: "Email", Parts: []w.FieldPart{w.Input{Value: "v@example.org", Label: "Email"}}, Width: w.FieldMD}, w.Field{Label: "Password", Parts: []w.FieldPart{w.Input{Value: "••••••••••••", Label: "Password"}}, Width: w.FieldMD}, w.Field{Label: "Repeat password", Parts: []w.FieldPart{w.Input{Value: "••••••••••••", Label: "Repeat password"}}, Width: w.FieldMD}}, Variant: w.FormFlushAll},
		w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Create account", Kind: w.ButtonPrimary}, w.Span{Face: w.FaceAccent12, Text: "Already have an account? Sign in"}}, Variant: w.FormFootFlush}}, Variant: w.DialogPage440}}}}, "s-auth-2", "tc-root"),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogLogo{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo48}, Wide: true},
						w.DialogTitle{Text: "Check your email", Gap: w.DialogTitleGapNormal},
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
	}, "s-auth-3", "tc-root"),
	page("plain", w.Page{Items: []w.PagePart{w.Auth{Dialog: w.Dialog{Items: []w.DialogPart{
		w.DialogTitle{Text: "Set up Carrel", Gap: w.DialogTitleGapTight},
		w.DialogSub{Parts: []w.Inline{w.Text("The account you create here is the first one, and it can manage everyone else.")}},
		w.Msg{Parts: []w.MsgPart{w.Text("Nobody can recover this password, this server included.")}, Tone: w.MsgAlert, Space: w.MsgFlushBelowWide},
		w.Form{Items: []w.FormPart{w.Field{Label: "Login", Parts: []w.FieldPart{w.Input{Placeholder: "a-z 0-9 . _ -", Label: "Login", Mono: true, Focus: true}}, Width: w.FieldMD}, w.Field{Label: "Email", LabelAside: w.Dim{Parts: []w.Inline{w.Text("(optional)")}}, Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Email"}}, Width: w.FieldMD}, w.Field{Label: "Password", Parts: []w.FieldPart{w.Input{Value: "••••••••••••", Label: "Password"}}, Width: w.FieldMD}, w.Field{Label: "Repeat password", Parts: []w.FieldPart{w.Input{Value: "••••••••••••", Label: "Repeat password"}}, Width: w.FieldMD}}, Variant: w.FormFlushAll},
		w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Create administrator", Kind: w.ButtonPrimary}}, Variant: w.FormFootFlush}}, Variant: w.DialogPage440}}}}, "s-auth-4", "tc-root"),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogTitle{Text: "Forgotten your password?", Gap: w.DialogTitleGapNormal},
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
	}, "s-auth-5", "tc-root"),
	page("plain", w.Page{Items: []w.PagePart{w.Auth{Dialog: w.Dialog{Items: []w.DialogPart{
		w.DialogTitle{Text: "What we can do", Gap: w.DialogTitleGapNormal},
		w.Prose{Parts: []w.Inline{w.Text("Account "), w.Strong("mix"), w.Text(" has a confirmed email. A silent password reset does not exist: a new password would be a new key, and everything stored under the old one would be lost anyway.")}, Variant: w.ProseMuted},
		w.Form{Items: []w.FormPart{w.Card{Items: []w.CardPart{
			w.CardText{Kind: w.CardTerm, Parts: []w.CardTextPart{w.Text("Wipe the account")}},
			w.Prose{Parts: []w.Inline{w.Text("We send a letter to the confirmed address. The link in it is not a reset: after you confirm, everything stored here is erased — connections become unreadable and are deleted — and the same link then lets you set a new password. Address books, calendars, tasks and notes on the servers stay where they are.")}, Variant: w.ProseMutedClose},
			w.Button{Label: "Send a confirmation letter", Kind: w.ButtonPrimary}}, Variant: w.CardSnug}, w.Card{Items: []w.CardPart{
			w.CardText{Kind: w.CardTerm, Parts: []w.CardTextPart{w.Text("Ask the administrator")}},
			w.Prose{Parts: []w.Inline{w.Text("This account deposited its encryption key. The administrator can recover it with the instance master password, and your data is kept.")}, Variant: w.ProseMutedClose},
			w.Button{Label: "Ask the administrator"}}, Variant: w.CardSnug}}, Variant: w.FormColumnWide},
		w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Back to sign in"}}, Variant: w.FormFootFlush}}, Variant: w.DialogPage520}}}}, "s-auth-6", "tc-root"),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogTitle{Text: "We cannot help from here", Gap: w.DialogTitleGapNormal},
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
	}, "s-auth-7", "tc-root"),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogTitle{Text: "Check your inbox", Gap: w.DialogTitleGapNormal},
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
	}, "s-auth-8", "tc-root"),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogTitle{Text: "The administrator has been asked", Gap: w.DialogTitleGapNormal},
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
	}, "s-auth-9", "tc-root"),
	page("plain", w.Page{Items: []w.PagePart{w.Auth{Dialog: w.Dialog{Items: []w.DialogPart{
		w.DialogTitle{Text: "Erase everything and start again?", Gap: w.DialogTitleGapNormal},
		w.Prose{Parts: []w.Inline{w.Text("This is not a password recovery. The key to your data comes from the password you forgot, so nothing stored under it can be read again — by you, by us, by anyone.\n\nConfirming erases what Carrel keeps encrypted for "), w.Strong("mix"), w.Text(": saved connections and their passwords, subscriptions, published links, devices, external backup triggers and the deposited copy of the key. Address books, calendars, tasks and notes on the servers are not touched — you connect to them again afterwards.\n\nThe account itself stays: same login, role and confirmed address, the same first day of the week and clock. Only the password is new, and you set it on the next screen.")}, Variant: w.ProseMuted},
		w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Erase and set a new password", Kind: w.ButtonDangerSolid}, w.Button{Label: "Cancel"}}, Variant: w.FormFootFlush}}, Variant: w.DialogPage520}}}}, "s-auth-10", "tc-root"),
	page("plain", w.Page{Items: []w.PagePart{w.Auth{Dialog: w.Dialog{Items: []w.DialogPart{
		w.DialogTitle{Text: "Set a new password", Gap: w.DialogTitleGapTight},
		w.DialogSub{Parts: []w.Inline{w.Text("The account "), w.Mono{Text: "mix"}, w.Text(" is empty. Nothing was kept.")}},
		w.Msg{Parts: []w.MsgPart{w.Text("Nobody can recover this password, this server included.")}, Tone: w.MsgAlert, Space: w.MsgFlushBelowWide},
		w.Form{Items: []w.FormPart{w.Field{Label: "Password", Parts: []w.FieldPart{w.Input{Value: "••••••••••••", Label: "Password", Focus: true}}, Width: w.FieldMD}, w.Field{Label: "Repeat password", Parts: []w.FieldPart{w.Input{Value: "••••••••••••", Label: "Repeat password"}}, Width: w.FieldMD}}, Variant: w.FormFlushAll},
		w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Set the password and sign in", Kind: w.ButtonPrimary}}, Variant: w.FormFootFlush}}, Variant: w.DialogPage440}}}}, "s-auth-11", "tc-root"),
	page("plain", w.Page{Items: []w.PagePart{w.Auth{Dialog: w.Dialog{Items: []w.DialogPart{
		w.DialogLogo{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo48}, Wide: true},
		w.DialogTitle{Text: "This link does not work", Gap: w.DialogTitleGapNormal},
		w.Prose{Parts: []w.Inline{w.Text("It may have expired, been used already, or been withdrawn. Ask whoever sent it for a new one.")}, Variant: w.ProseMuted},
		w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Sign in", Kind: w.ButtonPrimary}}, Variant: w.FormFootFlush}}, Variant: w.DialogNarrow}}}}, "rest-1", "tc-root"),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogLogo{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo48}, Wide: true},
						w.DialogTitle{Text: "Email address", Gap: w.DialogTitleGapNormal},
						w.Prose{Parts: []w.Inline{w.Text("Address confirmed. You can sign in now.")}, Variant: w.ProseMuted},
						w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Sign in", Kind: w.ButtonPrimary}}, Variant: w.FormFootFlush},
					},
					Variant: w.DialogNarrow,
				},
			},
		},
	}, "rest-2", "tc-root"),
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
	}, "rest-3", "tc-root"),
	page("print", w.Page{Items: []w.PagePart{w.PrintTitle{Text: "Agenda · 10 – 16 August 2026"}, w.PrintSources{Text: "nextcloud · Личный, Работа · radicale · vps · Проекты, Дежурства"}, w.PrintRubric{Text: "Monday · 10 August"}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "10:00", Variant: w.MonoFace}, w.PrintCell{Name: "Синк-ап по релизу", Detail: "Переговорная 2 · 4 attendees · repeats weekly · nextcloud · Личный"}}}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "all day", Variant: w.MonoFace}, w.PrintCell{Name: "Дежурство: первая линия", Detail: "repeats weekly · nextcloud · Работа"}}}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "all day", Variant: w.MonoFace}, w.PrintCell{Name: "Дежурство: первая линия", Detail: "repeats weekly · radicale · vps · Дежурства"}}}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "15:00", Variant: w.MonoFace}, w.PrintCell{Name: "Созвон с подрядчиком", Detail: "Jitsi · tentative · nextcloud · Работа"}}}, w.PrintRubric{Text: "Tuesday · 11 August"}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "09:30", Variant: w.MonoFace}, w.PrintCell{Name: "Ревью архитектуры", Detail: "Переговорная 1 · 2 attendees · radicale · vps · Проекты"}}}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "18:00", Variant: w.MonoFace}, w.PrintCell{Name: "Забрать документы", Detail: "МФЦ, Мясницкая 13 · nextcloud · Личный"}}}, w.PrintRubric{Text: "Wednesday · 12 August"}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "all day", Variant: w.MonoFace}, w.PrintCell{Name: "Поездка в Казань", Detail: "day 1 of 3 · nextcloud · Личный"}}}, w.PrintRubric{Text: "Thursday · 13 August"}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "all day", Variant: w.MonoFace}, w.PrintCell{Name: "Поездка в Казань", Detail: "day 2 of 3 · nextcloud · Личный"}}}, w.PrintRubric{Text: "Friday · 14 August"}, w.PrintRow{Items: []w.PrintRowPart{w.Mono{Text: "all day", Variant: w.MonoFace}, w.PrintCell{Name: "Поездка в Казань", Detail: "day 3 of 3 · nextcloud · Личный"}}}, w.PrintFoot{Text: "Carrel · agenda · 2026-08-10 – 2026-08-16 · captured 2026-08-15 09:41 · 5 of 7 sources"}}, Variant: w.PagePrint}, "print-1", "tc-root"),
	page("plain", w.Page{Items: []w.PagePart{w.Window{Items: []w.WindowPart{w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo42}, w.Stack{Gap: w.StackGapNormal, Items: []w.StackPart{w.Card{Items: []w.CardPart{
		w.CardPick{Title: "Just for me", Note: []w.CardTextPart{w.Text("Installs into your own profile. No administrator needed.")}}}, Variant: w.CardChosenFlush}, w.Card{Items: []w.CardPart{
		w.CardPick{Ticked: w.TickedOff, Title: "For everyone on this computer", Note: []w.CardTextPart{w.Text("Installs into "), w.Mono{Text: "Program Files"}, w.Text(". Asks for administrator once; each account still keeps its own data and password.")}}}, Variant: w.CardFlush}, w.Card{Items: []w.CardPart{
		w.CardPick{Ticked: w.TickedOff, Title: "Portable — do not install", Note: []w.CardTextPart{w.Text("Unpacks into a folder you choose and keeps everything inside it: the app, the server binary and your encrypted state.")}, Field: &w.Field{Parts: []w.FieldPart{w.Line{Items: []w.LinePart{w.Input{Placeholder: "E:\\Carrel", Label: "E:\\Carrel", Mono: true, Variant: w.Input11Fill}, w.Button{Label: "Choose"}}, Gap: w.LineGapNormal}, w.FieldHint{Parts: []w.Inline{w.Text("Nothing is written outside this folder; move it, and Carrel moves with it.")}}}, Width: w.FieldW100}}}, Variant: w.CardFlush}}}, w.Msg{Icon: "shield", Parts: []w.MsgPart{w.Strong("Lose the stick, and the finder can guess your password as long as they like."), w.Text(" The encryption holds only as well as the password: choose a long one.")}, Tone: w.MsgAlert, Space: w.MsgFlush}, w.Line{Items: []w.LinePart{w.Button{Label: "Continue", Kind: w.ButtonPrimary}, w.Button{Label: "Cancel"}, w.Dim{Parts: []w.Inline{w.Text("64 MB · app, server, WebView2 check")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Gap: w.LineGapNormal, Block: true}}, Gap: w.WindowGapWide}}}, "desktop-3", "tc-root"),
	page("plain", w.Page{Items: []w.PagePart{w.Window{Items: []w.WindowPart{w.Form{Items: []w.FormPart{w.Field{Label: "Everything lives here", Parts: []w.FieldPart{w.Input{Value: "E:\\Carrel\\ — app, carrel.exe, data\\state.enc", Label: "Everything lives here", Mono: true, Variant: w.Input11_5}, w.FieldHint{Parts: []w.Inline{w.Text("The same folder on any computer: plug the stick in, run, sign in. In this mode the data folder always sits next to the app.")}}}, Place: w.Grown}, w.Field{Parts: []w.FieldPart{w.Check{Parts: []w.CheckPart{w.Dim{Parts: []w.Inline{w.Text("Start at sign-in — not available in a portable copy")}}}}, w.Check{On: true, Parts: []w.CheckPart{w.Text("Closing the window hides to the tray")}}}, Gapped: true, Place: w.Grown}, w.Field{Label: "Updates", Parts: []w.FieldPart{w.Line{Items: []w.LinePart{w.Input{Placeholder: "Replaces the files in this folder", Label: "Replaces the files in this folder", Variant: w.InputFill}, w.Button{Label: "Check now"}}, Gap: w.LineGapNormal}, w.FieldHint{Parts: []w.Inline{w.Text("Works while the folder is writable. On a read-only medium the app says so and offers the download instead of installing it.")}}}, Place: w.Grown}}, Variant: w.FormGapFields}, w.Msg{Parts: []w.MsgPart{w.Text("A portable copy is a full Carrel, local or remote as usual.")}, Space: w.MsgFlush}}, Gap: w.WindowGapWide}}}, "desktop-4", "tc-root"),
	page("plain", w.Page{
		Items: []w.PagePart{
			w.Window{
				Items: []w.WindowPart{
					w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo42},
					w.Stack{
						Gap: w.StackGapWide,
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
					w.Line{Items: []w.LinePart{w.Button{Label: "Next", Kind: w.ButtonPrimary}}, Gap: w.LineGapNormal, Block: true},
				},
				Gap: w.WindowGapWide,
			},
		},
	}, "desktop-5", "tc-root"),
	page("plain", w.Page{Items: []w.PagePart{w.Window{Items: []w.WindowPart{w.Form{Items: []w.FormPart{w.Field{Label: "Data folder", Parts: []w.FieldPart{w.Input{Value: "C:\\Users\\mix\\AppData\\Local\\Carrel", Label: "Data folder", Mono: true, Variant: w.Input11_5}, w.FieldHint{Parts: []w.Inline{w.Text("Your connections and settings, encrypted. Records stay on the servers.")}}}, Place: w.Grown}, w.Field{Label: "Port", Parts: []w.FieldPart{w.Input{Value: "chosen at every start · now 51734", Label: "Port", Mono: true}, w.FieldHint{Parts: []w.Inline{w.Text("A free one at every start; its number is in the status bar.")}}}, Width: w.FieldFill300}, w.Field{Parts: []w.FieldPart{w.Check{On: true, Parts: []w.CheckPart{w.Text("Stop the server on exit")}}}}, w.Field{Parts: []w.FieldPart{w.Check{Parts: []w.CheckPart{w.Text("Allow other devices on this network")}}, w.FieldHint{Parts: []w.Inline{w.Text("Turning it on "), w.Strong("pins the port"), w.Text(": a phone and DAVx5 reach "), w.Mono{Text: "http://192.168.1.14:8765"}, w.Text(" with your password.")}}}, Gapped: true, OwnLine: true, Place: w.Grown}}, Variant: w.FormGapFields}, w.FieldSet{Rubric: w.Rubric{Text: "Account"}, Items: []w.FieldSetPart{w.Field{Label: "Login", Parts: []w.FieldPart{w.Input{Value: "mix", Label: "Login"}}, Width: w.FieldMD}, w.Field{Label: "Password", Parts: []w.FieldPart{w.Input{Value: "••••••••••••", Label: "Password"}}, Width: w.FieldMD}, w.Msg{Parts: []w.MsgPart{w.Text("Nobody can recover the password, not even this app; without it the saved connections are lost.")}, Tone: w.MsgAlert, Space: w.MsgFlushTop2, Place: w.Grown}}, Variant: w.FieldSetRule}, w.Msg{Icon: "shield", Parts: []w.MsgPart{w.Strong("Not reachable from the internet."), w.Text(" "), w.Strong("Published links"), w.Text(", "), w.Strong("devices over davloom"), w.Text(", "), w.Strong("external backup triggers"), w.Text(" and "), w.Strong("invitation links"), w.Text(" work on this machine only, until you set up port forwarding, a tunnel or your own domain.")}, Tone: w.MsgAlert, Space: w.MsgFlush}, w.Line{Items: []w.LinePart{w.Button{Label: "Run", Kind: w.ButtonPrimary}, w.Button{Label: "Back"}, w.Dim{Parts: []w.Inline{w.Dot{Tone: w.DotAccent}, w.Text("server stopped")}, Size: w.Dim11_5, Variant: w.DimWrap, Place: w.Pushed}}, Gap: w.LineGapNormal, Block: true}}, Gap: w.WindowGapWide}}}, "desktop-6", "tc-root"),
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
						Variant: w.FormColumnWide,
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
						Gap:   w.LineGapNormal,
						Block: true,
					},
				},
				Gap: w.WindowGapWide,
			},
		},
	}, "desktop-7", "tc-root"),
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
						Gap:   w.LineGapNormal,
						Block: true,
					},
					w.Dim{
						Parts: []w.Inline{w.Text("The window closes itself. Carrel need not be open — the note goes to the server at once.")},
						Size:  w.Dim11,
					},
				},
				Gap:   w.WindowGapNormal,
				Small: true,
			},
		},
		Variant: w.PageOnBackground,
	}, "desktop-9", "tc-root"),
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
						Gap:   w.LineGapNormal,
						Block: true,
					},
					w.Dim{
						Parts: []w.Inline{
							w.Text("The same sheet as in the window, in a window of its own. Carrel need not be open — the event goes to the server at once."),
						},
						Size: w.Dim11,
					},
				},
				Gap:   w.WindowGapNormal,
				Small: true,
			},
		},
		Variant: w.PageOnBackground,
	}, "desktop-10", "tc-root"),
	page("on-background-column", w.Page{Items: []w.PagePart{w.Form{Items: []w.FormPart{w.Rubric{Text: "Reminders"}, w.Field{Label: "Remind me", Parts: []w.FieldPart{w.Pick{Value: "10 minutes · and as set in the event"}}, Width: w.FieldW100}, w.Check{On: true, Parts: []w.CheckPart{w.Text("Task due dates — on the morning they are due")}}, w.Check{Parts: []w.CheckPart{w.Text("Show the title on the lock screen")}}, w.Field{Label: "Quiet hours", Parts: []w.FieldPart{w.Input{Value: "22:00 – 08:00", Label: "Quiet hours", Mono: true}}, Width: w.FieldW100}, w.Field{Label: "From which collections", Parts: []w.FieldPart{w.Input{Value: "3 of 7 ticked", Label: "From which collections"}}, Width: w.FieldW100}}, Variant: w.FormColumnWideMid}}, Variant: w.PageOnBackgroundColumn}, "desktop-11", "tc-root"),
	page("on-background", w.Page{Items: []w.PagePart{w.Window{Items: []w.WindowPart{w.Msg{Parts: []w.MsgPart{w.Bold{Parts: []w.Inline{w.Text("Sign out of "), w.Mono{Text: "carrel.example"}, w.Text("?")}}, w.Text(" "), w.Strong("Your other devices stay signed in"), w.Text(": end those sessions in Settings → Account.")}, Space: w.MsgFlush}, w.Msg{Icon: "shield", Parts: []w.MsgPart{w.Strong("Delete what this app keeps on this computer?"), w.Text(" The folder "), w.Mono{Text: "C:\\Users\\mix\\AppData\\Roaming\\Carrel"}, w.Text(", the instance address and the session token. The saved connections cannot be brought back; records on your DAV servers are not touched.")}, Tone: w.MsgAlert, Space: w.MsgFlush}, w.Line{Items: []w.LinePart{w.Button{Label: "Keep and sign out", Kind: w.ButtonPrimary}, w.Kbd{Keys: "Enter"}, w.Button{Label: "Delete and sign out", Kind: w.ButtonDangerSolid}, w.Line{Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}, Gap: w.LineGapNormal, Place: w.Pushed}}, Gap: w.LineGapNormal, Block: true}}, Gap: w.WindowGapWide, Small: true}}, Variant: w.PageOnBackground}, "desktop-12", "tc-root"),
	page("plain", w.Page{Items: []w.PagePart{w.Msg{Icon: "download", Parts: []w.MsgPart{w.Strong("Carrel 0.8.0 is ready to install."), w.Text(" It is applied when the app closes, or now: the window reopens where it was.")}, End: &w.Push{Items: []w.PushPart{w.Button{Label: "Install and restart", Kind: w.ButtonPrimary}, w.Button{Label: "What’s new"}, w.Button{Label: "Later"}}}, Tone: w.MsgAccent, Space: w.MsgFlushEnd}, w.Window{Items: []w.WindowPart{w.Form{Items: []w.FormPart{w.Field{Label: "Installed", Parts: []w.FieldPart{w.Input{Value: "app 0.8.0 · server 0.8.0 — both run inside this app", Label: "Installed", Mono: true, Variant: w.Input11_5}, w.FieldHint{Parts: []w.Inline{w.Text("In local mode one update covers the app and the server. Connections and settings stay as they are.")}}}, Place: w.Grown}, w.Field{Label: "Downloaded", Parts: []w.FieldPart{w.Line{Items: []w.LinePart{w.Meter{Percent: 100, Width: w.MeterFill}, w.Mono{Text: "0.8.0 · 14.2 MB · signature checked", Variant: w.Mono11Dusty}}, Gap: w.LineGapNormal}}, Place: w.Grown}}, Variant: w.FormGapFields}, w.Msg{Parts: []w.MsgPart{w.Text("You see this because automatic checking is on; out of the box the app never goes looking.")}, Space: w.MsgFlush}}, Gap: w.WindowGapWide}}}, "desktop-14", "tc-root"),
	page("plain", w.Page{Items: []w.PagePart{w.Window{Items: []w.WindowPart{w.Msg{Icon: "shield", Parts: []w.MsgPart{w.Strong("The app is newer than the server."), w.Text(" This app is 0.8.0, "), w.Mono{Text: "carrel.example"}, w.Text(" runs 0.6.4. What both know works; what is newer than the server is not there.")}, Tone: w.MsgAlert, Space: w.MsgFlush}, w.Form{Items: []w.FormPart{w.Field{Label: "What to do", Parts: []w.FieldPart{w.Input{Placeholder: "Update the server where it runs — this app cannot do it", Label: "What to do"}, w.FieldHint{Parts: []w.Inline{w.Text("The app does not update a server from outside: that would be a way in. The server’s administrator has to update it.")}}}, Place: w.Grown}}, Variant: w.FormGapFields}, w.Line{Items: []w.LinePart{w.Button{Label: "How to update a server"}, w.Button{Label: "Do not warn about this server"}}, Gap: w.LineGapNormal, Block: true}}, Gap: w.WindowGapWide}}}, "desktop-15", "tc-root"),
	page("print", w.Page{Items: []w.PagePart{w.PrintTitle{Text: "Documents / carrel / attachments / 2026-07"}, w.PrintSources{Text: "sftpgo · vps · 7 items · 89.4 MB"}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Name"}, w.PrintNum{Text: "Size"}, w.PrintNum{Text: "Changed"}}, Columns: true, Head: true}, w.PrintRubric{Text: "Folders"}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "архив"}, w.PrintNum{Text: "—"}, w.PrintNum{Text: "2026-07-01"}}, Columns: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "черновики"}, w.PrintNum{Text: "0 items"}, w.PrintNum{Text: "2026-07-29"}}, Columns: true}, w.PrintRubric{Text: "Files"}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "панорама.tif"}, w.PrintNum{Text: "84 MB"}, w.PrintNum{Text: "2026-07-21"}}, Columns: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "схема-2.png"}, w.PrintNum{Text: "712 KB"}, w.PrintNum{Text: "2026-07-16"}}, Columns: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "схема-1.png"}, w.PrintNum{Text: "640 KB"}, w.PrintNum{Text: "2026-07-16"}}, Columns: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "whiteboard.jpg"}, w.PrintNum{Text: "2.8 MB"}, w.PrintNum{Text: "2026-07-14"}}, Columns: true}, w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "смета.pdf"}, w.PrintNum{Text: "1.2 MB"}, w.PrintNum{Text: "2026-07-08"}}, Columns: true}, w.PrintFoot{Text: "Carrel · folder · sftpgo / Documents / carrel / attachments / 2026-07 · captured 2026-08-16 20:41 · 7 of 7 shown"}}, Variant: w.PagePrint}, "desktop-16", "tc-root"),
)

func init() {
	Widgets = append(Widgets, w.Page{})
	Parts = append(Parts, w.Window{}, w.WindowNote{}, w.PrintTitle{}, w.PrintSources{}, w.PrintRubric{}, w.PrintRow{},
		w.PrintCell{}, w.PrintNum{}, w.PrintFoot{})
}
