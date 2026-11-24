// Раздел макета: 7.24 Почта ящика — утверждена, отложена.

package screens

import w "carrel.dev/typecase/widget"

// Mailbox1 is frame «Почта · входящие» (s-mailbox-1).
func Mailbox1() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 10}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "mail", Label: "Mail", Count: &w.Num{Value: 12}, Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 6,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{
							Text: "",
							Items: []w.RubricPart{
								w.Line{
									Items: []w.LinePart{w.Mark{Colour: "#A05A2C", Size: w.MarkH11}, w.Text("home")},
									Gap:   w.LineGapNormal,
								},
							},
						},
						w.Tree{
							Rows: []w.TreeRow{
								{Icon: "inbox", Label: "Inbox", Count: &w.Num{Value: 12}},
								{Icon: "pencil", Label: "Drafts", Count: &w.Num{Value: 1}},
								{Icon: "send", Label: "Sent"},
								{Icon: "archive", Label: "Archive"},
								{Icon: "spam", Label: "Spam", Count: &w.Num{Value: 3}},
								{Icon: "trash", Label: "Trash"},
								{Icon: "file", Label: "Стройка", Count: &w.Num{Value: 2}},
								{Icon: "file", Label: "Рассылки"},
							},
							On:      1,
							Variant: w.TreePadded,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{
							Text: "",
							Items: []w.RubricPart{
								w.Line{
									Items: []w.LinePart{w.Mark{Colour: "#3F6E8C", Size: w.MarkH11}, w.Text("office")},
									Gap:   w.LineGapNormal,
								},
							},
						},
					},
					Failed: "no reply · 8 s",
					Why:    "mail.office.example:993 did not answer in 8 s. Other mailboxes and the rest of Carrel work.",
					Retry:  &w.Button{Label: "Retry", Href: "?retry=office"},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "mail", Title: "Connect a mailbox"},
						w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Inbox"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{
									w.Text("home · "),
									w.Mono{Text: "mix@home.example"},
									w.Text(" · 1 284 letters · 12 unread · newest first"),
								},
							},
						},
						Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Write", Icon: "pencil", Kind: w.ButtonPrimary}}},
					},
					w.List{
						Rows: []w.ListRow{
							w.MailRow{
								Ticked:    w.TickedOff,
								Important: true,
								Unread:    true,
								From:      "Анна Ковалёва",
								Topic:     "Вёрстка сборника: правки и встреча в четверг",
								Start:     "Михаил, добрый день! Прислала правки к третьей вёрстке, они в PDF и отмечены жёлтым…",
								Attached:  3,
								Bytes:     "2.6 MB",
								Date:      "2026-09-24 10:12",
							},
							w.MailRow{
								Ticked:   w.TickedOff,
								Unread:   true,
								From:     "Пётр Смирнов",
								Topic:    "Смета по второму этапу",
								Start:    "Во вложении смета, кровлю пересчитал по новым ценам, фасад без изменений…",
								Attached: 1,
								Bytes:    "1.1 MB",
								Date:     "2026-09-24 08:47",
							},
							w.MailRow{
								Ticked: w.TickedOff,
								Unread: true,
								From:   "Книжная ярмарка",
								Topic:  "Сентябрьская рассылка: сорок новинок осени",
								Start:  "Не видите картинки? Откройте письмо в браузере…",
								Bytes:  "84 KB",
								Date:   "2026-09-23 19:30",
							},
							w.MailRow{
								Ticked: w.TickedOff,
								From:   "Виктор Гончаров",
								Topic:  "Re: Миграция на копии базы",
								Start:  "Прогнал на копии, всё сошлось, кроме вложений старше 2019 года…",
								Bytes:  "12 KB",
								Date:   "2026-09-23 16:05",
							},
							w.MailRow{
								Ticked: w.TickedOff,
								From:   "Nextcloud",
								Topic:  "Your storage is almost full",
								Start:  "You are using 91% of your storage quota on cloud.example.org…",
								Bytes:  "9 KB",
								Date:   "2026-09-23 03:00",
							},
							w.MailRow{
								Ticked:    w.TickedOff,
								Important: true,
								From:      "Артём Белов",
								Topic:     "Скриншоты к релизу",
								Start:     "Лови, шесть штук, как договаривались. Тёмная тема — отдельно…",
								Attached:  6,
								Bytes:     "5.8 MB",
								Date:      "2026-09-22 14:21",
							},
							w.MailRow{
								Ticked: w.TickedOff,
								From:   "Вера Лебедева",
								Topic:  "Дача: ключи и счётчики",
								Start:  "Ключи оставила у соседки, показания сфотографировала…",
								Bytes:  "6 KB",
								Date:   "2026-09-21 11:40",
							},
							w.MailRow{
								Ticked:   w.TickedOff,
								From:     "Банк «Пример»",
								Topic:    "Выписка за август",
								Start:    "Выписка по счёту во вложении, пароль к файлу — последние четыре цифры…",
								Attached: 1,
								Bytes:    "240 KB",
								Date:     "2026-09-20 09:00",
							},
						},
					},
					w.Empty{
						Parts:   []w.EmptyPart{w.Span{Face: w.FaceAccent, Text: "Show 50 more", Href: "/app/mail/home/inbox?shown=100"}},
						Variant: w.EmptyDusty,
					},
				},
			},
		},
	}
}

// Mailbox2 is frame «Почта · ящика ещё нет» (s-mailbox-2).
func Mailbox2() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 10}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "mail", Label: "Mail", Count: &w.Num{Value: 12}, Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 6,
				},
				w.RailSec{Items: []w.RailSecPart{
					w.Rubric{Text: "Mailboxes"}}, Note: "No mailbox connected yet"},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "mail", Title: "Connect a mailbox"},
						w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Mail"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("No mailbox connected yet")}}},
					},
					w.Empty{
						Parts: []w.EmptyPart{
							w.Text("Carrel opens your own mailbox over IMAP and sends from it over SMTP. It keeps no copy of your mail: letters, folders and flags stay on your mail server, and Carrel only shows them. You need the address and a password — for Gmail, Yandex, Mail.ru and iCloud, an app password made in the account's security settings. "),
							w.Button{Label: "Connect a mailbox", Icon: "plus", Kind: w.ButtonPrimary},
						},
					},
				},
			},
		},
	}
}

// Mailbox3 is frame «Входящие · выбраны три письма» (s-mailbox-3).
func Mailbox3() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 10}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "mail", Label: "Mail", Count: &w.Num{Value: 12}, Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 6,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{
							Text: "",
							Items: []w.RubricPart{
								w.Line{
									Items: []w.LinePart{w.Mark{Colour: "#A05A2C", Size: w.MarkH11}, w.Text("home")},
									Gap:   w.LineGapNormal,
								},
							},
						},
						w.Tree{
							Rows: []w.TreeRow{
								{Icon: "inbox", Label: "Inbox", Count: &w.Num{Value: 12}},
								{Icon: "pencil", Label: "Drafts", Count: &w.Num{Value: 1}},
								{Icon: "send", Label: "Sent"},
								{Icon: "archive", Label: "Archive"},
								{Icon: "spam", Label: "Spam", Count: &w.Num{Value: 3}},
								{Icon: "trash", Label: "Trash"},
								{Icon: "file", Label: "Стройка", Count: &w.Num{Value: 2}},
								{Icon: "file", Label: "Рассылки"},
							},
							On:      1,
							Variant: w.TreePadded,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{
							Text: "",
							Items: []w.RubricPart{
								w.Line{
									Items: []w.LinePart{w.Mark{Colour: "#3F6E8C", Size: w.MarkH11}, w.Text("office")},
									Gap:   w.LineGapNormal,
								},
							},
						},
					},
					Failed: "no reply · 8 s",
					Why:    "mail.office.example:993 did not answer in 8 s. Other mailboxes and the rest of Carrel work.",
					Retry:  &w.Button{Label: "Retry", Href: "?retry=office"},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "mail", Title: "Connect a mailbox"},
						w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Inbox"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{
									w.Text("home · "),
									w.Mono{Text: "mix@home.example"},
									w.Text(" · 1 284 letters · 12 unread · newest first"),
								},
							},
						},
						Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Write", Icon: "pencil", Kind: w.ButtonPrimary}}},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Mono{Text: "6.8 MB", Variant: w.MonoMuted},
							w.BarSep{},
							w.Button{Folds: true, Label: "Read", Icon: "mailopen", Title: "Mark as read", Secondary: true},
							w.Button{Folds: true, Label: "Unread", Icon: "mail", Title: "Mark as unread", Secondary: true},
							w.Button{Folds: true, Label: "Important", Icon: "flag", Title: "Mark as important", Secondary: true},
							w.Button{Folds: true, Label: "Move to", Menu: true, Icon: "move", Title: "Move to"},
							w.Button{Folds: true, Label: "Archive", Icon: "archive", Title: "Archive"},
							w.Button{Folds: true, Label: "Spam", Icon: "spam", Title: "Spam", Secondary: true},
							w.Button{Folds: true, Label: "Trash", Icon: "trash", Title: "Move to Trash"},
							w.Drop{
								Open:    true,
								Trigger: w.Button{Icon: "dots", Title: "More"},
								More:    true,
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuItem{Icon: "mailopen", Label: "Mark as read"},
										w.MenuItem{Icon: "mail", Label: "Mark as unread"},
										w.MenuItem{Icon: "flag", Label: "Mark as important"},
										w.MenuItem{Icon: "spam", Label: "Spam"},
									},
									Width: w.MenuWidth212,
									Pop:   w.MenuUnderEnd,
								},
							},
						},
						Count: &w.BarCount{Text: "3 selected"},
						Right: &w.BarRight{
							Items: []w.BarPart{
								w.Button{
									Folds:      true,
									Label:      "Select all",
									Icon:       "selectall",
									Small:      true,
									Title:      "Select all",
									ScriptOnly: true,
								},
								w.Button{Folds: true, Label: "Clear selection", Icon: "close", Small: true, Title: "Clear selection"},
							},
						},
						Variant: w.BarSelection,
					},
					w.List{
						Rows: []w.ListRow{
							w.MailRow{
								Important: true,
								Unread:    true,
								On:        true,
								From:      "Анна Ковалёва",
								Topic:     "Вёрстка сборника: правки и встреча в четверг",
								Start:     "Михаил, добрый день! Прислала правки к третьей вёрстке…",
								Attached:  3,
								Bytes:     "2.6 MB",
								Date:      "2026-09-24 10:12",
							},
							w.MailRow{
								Unread:   true,
								On:       true,
								From:     "Пётр Смирнов",
								Topic:    "Смета по второму этапу",
								Start:    "Во вложении смета, кровлю пересчитал по новым ценам…",
								Attached: 1,
								Bytes:    "1.1 MB",
								Date:     "2026-09-24 08:47",
							},
							w.MailRow{
								Ticked: w.TickedOff,
								Unread: true,
								From:   "Книжная ярмарка",
								Topic:  "Сентябрьская рассылка: сорок новинок осени",
								Start:  "Не видите картинки? Откройте письмо в браузере…",
								Bytes:  "84 KB",
								Date:   "2026-09-23 19:30",
							},
							w.MailRow{
								Ticked: w.TickedOff,
								From:   "Виктор Гончаров",
								Topic:  "Re: Миграция на копии базы",
								Start:  "Прогнал на копии, всё сошлось, кроме вложений старше 2019 года…",
								Bytes:  "12 KB",
								Date:   "2026-09-23 16:05",
							},
							w.MailRow{
								Important: true,
								On:        true,
								From:      "Артём Белов",
								Topic:     "Скриншоты к релизу",
								Start:     "Лови, шесть штук, как договаривались…",
								Attached:  6,
								Bytes:     "5.8 MB",
								Date:      "2026-09-22 14:21",
							},
						},
					},
				},
			},
		},
	}
}

// Mailbox6 is frame «Письмо · чтение» (s-mailbox-6).
func Mailbox6() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 10}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "mail", Label: "Mail", Count: &w.Num{Value: 12}, Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 6,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{
							Text: "",
							Items: []w.RubricPart{
								w.Line{
									Items: []w.LinePart{w.Mark{Colour: "#A05A2C", Size: w.MarkH11}, w.Text("home")},
									Gap:   w.LineGapNormal,
								},
							},
						},
						w.Tree{
							Rows: []w.TreeRow{
								{Icon: "inbox", Label: "Inbox", Count: &w.Num{Value: 12}},
								{Icon: "pencil", Label: "Drafts", Count: &w.Num{Value: 1}},
								{Icon: "send", Label: "Sent"},
								{Icon: "archive", Label: "Archive"},
								{Icon: "spam", Label: "Spam", Count: &w.Num{Value: 3}},
								{Icon: "trash", Label: "Trash"},
								{Icon: "file", Label: "Стройка", Count: &w.Num{Value: 2}},
								{Icon: "file", Label: "Рассылки"},
							},
							On:      1,
							Variant: w.TreePadded,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{
							Text: "",
							Items: []w.RubricPart{
								w.Line{
									Items: []w.LinePart{w.Mark{Colour: "#3F6E8C", Size: w.MarkH11}, w.Text("office")},
									Gap:   w.LineGapNormal,
								},
							},
						},
					},
					Failed: "no reply · 8 s",
					Why:    "mail.office.example:993 did not answer in 8 s. Other mailboxes and the rest of Carrel work.",
					Retry:  &w.Button{Label: "Retry", Href: "?retry=office"},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "mail", Title: "Connect a mailbox"},
						w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Back: &w.Button{Label: "Inbox", Icon: "back", Href: "/app/mail/home/inbox#l48213"},
						Meta: &w.HeadSub{
							Parts: []w.HeadSubPart{
								w.Mark{Colour: "#A05A2C", Size: w.MarkH12},
								w.Text("home · Inbox · "),
								w.Mono{Text: "2026-09-24 10:12"},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Reply", Icon: "reply", Kind: w.ButtonPrimary},
								w.Button{Label: "Reply all", Icon: "replyall"},
								w.Button{Label: "Forward", Icon: "forward"},
								w.Button{Label: "Archive", Icon: "archive", Title: "Archive"},
								w.Button{Icon: "trash", Title: "Move to Trash"},
								w.Button{Icon: "dots", Title: "More"},
								w.Button{Icon: "up", Title: "Newer letter", Href: "/app/mail/home/inbox/48214"},
								w.Button{Icon: "down", Title: "Older letter", Href: "/app/mail/home/inbox/48190"},
							},
						},
						Pad: w.HeadRule,
					},
					w.DocLayout{
						Items: []w.DocLayoutPart{
							w.Doc{
								Items: []w.DocPart{
									w.DocTitle{Text: "Вёрстка сборника: правки и встреча в четверг"},
									w.Fields{
										Items: []w.FieldsItem{
											{
												Term:  "From",
												Value: []w.ValuePart{w.Text("Анна Ковалёва "), w.Dim{Parts: []w.Inline{w.Text("a.kovaleva@example.org")}}},
											},
											{
												Term:  "To",
												Value: []w.ValuePart{w.Text("Михаил Енученко "), w.Dim{Parts: []w.Inline{w.Text("mix@home.example")}}},
											},
											{
												Term:  "Cc",
												Value: []w.ValuePart{w.Text("Пётр Смирнов "), w.Dim{Parts: []w.Inline{w.Text("p.smirnov@example.org")}}},
											},
											{Term: "Date", Value: []w.ValuePart{w.Text("2026-09-24 10:12")}, Face: w.ValueMono},
										},
										Variant: w.FieldsLetter,
									},
									w.DocRuled{
										Items: []w.DocPart{
											w.Prose{
												Parts: []w.Inline{
													w.Text("Михаил, добрый день!\n\nПрислала правки к третьей вёрстке — они в PDF, отмечены жёлтым. Главное:"),
												},
												Variant: w.ProseRead,
											},
											w.DocBullets{
												Items: [][]w.DocItemPart{
													[]w.DocItemPart{w.Text("на с. 14 и 37 съехали подписи к иллюстрациям;")},
													[]w.DocItemPart{w.Text("в оглавлении не тот номер у главы «Переплёт»;")},
													[]w.DocItemPart{w.Text("обложку оставляем, как на прошлой неделе.")},
												},
											},
											w.Prose{
												Parts: []w.Inline{
													w.Text("В четверг в 11:00 жду вас в издательстве, приглашение во вложении. Схема проезда — на сайте: "),
													w.Link{Text: "Как нас найти", Href: "https://slovo.example/contacts", Address: "slovo.example/contacts"},
													w.Text(". Карточку Петра прикладываю — он возьмёт печать тиража.\n\n—\nАнна Ковалёва\nИздательство «Слово»"),
												},
												Variant: w.ProseRead,
											},
											w.DocQuote{
												Parts: []w.Inline{
													w.Text("2026-09-22 18:40, Михаил Енученко: Анна, третья вёрстка у вас. Посмотрите, пожалуйста, подписи — у меня на двух разворотах они уехали."),
												},
											},
										},
									},
									w.DocAttached{
										Title: w.Rubric{Text: "Attachments", Items: []w.RubricPart{w.Num{Value: 3}}, Variant: w.RubricBelow},
										Items: []w.DocAttachment{
											{
												Icon: "cal",
												Name: "Встреча в издательстве · 2026-09-26 11:00",
												Size: w.Mono{Text: ".ics · 2 KB", Variant: w.MonoDusty11},
												Acts: []w.Button{
													w.Button{Label: "Add to calendar", Icon: "cal", Small: true},
													w.Button{Icon: "download", Small: true, Title: "Download"},
												},
											},
											{
												Icon: "con",
												Name: "Пётр Смирнов",
												Size: w.Mono{Text: ".vcf · 3 KB", Variant: w.MonoDusty11},
												Acts: []w.Button{
													w.Button{Label: "Add to address book", Icon: "con", Small: true},
													w.Button{Icon: "download", Small: true, Title: "Download"},
												},
											},
											{
												Icon: "doc",
												Name: "правки-сборник-v3.pdf",
												Size: w.Mono{Text: "2.4 MB", Variant: w.MonoDusty11},
												Acts: []w.Button{
													w.Button{Label: "Save to files", Icon: "file", Small: true},
													w.Button{Icon: "download", Small: true, Title: "Download"},
												},
											},
										},
									},
								},
							},
							w.Docmeta{
								Sections: []w.DocmetaSec{
									{
										Title: w.Rubric{Text: "Sender"},
										Items: []w.DocmetaPart{
											w.DocRefs{
												Items: []w.DocRef{
													{
														Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11Inline},
														Name: "Анна Ковалёва",
														Href: "#",
														What: w.Dim{Parts: []w.Inline{w.Text("· nextcloud · Контакты")}},
													},
													{Name: "Timeline", Href: "#", What: w.Dim{Parts: []w.Inline{w.Text("· everything with her")}}},
												},
											},
										},
									},
									{
										Title: w.Rubric{Text: "Shown as"},
										Items: []w.DocmetaPart{
											w.DocAttach{Name: "The letter's text part", Size: w.Dim{Parts: []w.Inline{w.Text("· it also has HTML")}}},
											w.SecActs{Items: []w.Button{w.Button{Label: "Show original", Small: true}}},
										},
									},
									{
										Title: w.Rubric{Text: "Source"},
										Items: []w.DocmetaPart{w.Diag{Text: "UID     48213\nFolder  INBOX\nSize    2.6 MB\nRead    2 min ago"}},
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

// Mailbox9 is frame «Добавить в календарь из письма» (s-mailbox-9).
func Mailbox9() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.Overlay{
			Body: w.OverlayBody{
				Rail: []w.OverlayRailPart{
					w.Nav{
						Items: []w.NavItem{
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 10}, Href: "#"},
							{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
							{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
							{Icon: "file", Label: "Files", Href: "#"},
							{Icon: "mail", Label: "Mail", Count: &w.Num{Value: 12}, Href: "#"},
							{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
						},
						On: 6,
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{
								Text: "",
								Items: []w.RubricPart{
									w.Line{
										Items: []w.LinePart{w.Mark{Colour: "#A05A2C", Size: w.MarkH11}, w.Text("home")},
										Gap:   w.LineGapNormal,
									},
								},
							},
							w.Tree{
								Rows: []w.TreeRow{
									{Icon: "inbox", Label: "Inbox", Count: &w.Num{Value: 12}},
									{Icon: "pencil", Label: "Drafts", Count: &w.Num{Value: 1}},
									{Icon: "send", Label: "Sent"},
									{Icon: "archive", Label: "Archive"},
									{Icon: "spam", Label: "Spam", Count: &w.Num{Value: 3}},
									{Icon: "trash", Label: "Trash"},
									{Icon: "file", Label: "Стройка", Count: &w.Num{Value: 2}},
									{Icon: "file", Label: "Рассылки"},
								},
								On:      1,
								Variant: w.TreePadded,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{
								Text: "",
								Items: []w.RubricPart{
									w.Line{
										Items: []w.LinePart{w.Mark{Colour: "#3F6E8C", Size: w.MarkH11}, w.Text("office")},
										Gap:   w.LineGapNormal,
									},
								},
							},
						},
						Failed: "no reply · 8 s",
						Why:    "mail.office.example:993 did not answer in 8 s. Other mailboxes and the rest of Carrel work.",
						Retry:  &w.Button{Label: "Retry", Href: "?retry=office"},
					},
					w.RailFoot{
						Items: []w.RailFootPart{
							w.Button{Icon: "mail", Title: "Connect a mailbox"},
							w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"},
						},
						About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
					},
				},
				Main: w.Main{
					Items: []w.MainPart{
						w.Head{
							Meta: &w.HeadSub{
								Parts: []w.HeadSubPart{
									w.Mark{Colour: "#A05A2C", Size: w.MarkH12},
									w.Text("home · Inbox · "),
									w.Mono{Text: "2026-09-24 10:12"},
								},
							},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "Reply", Icon: "reply", Kind: w.ButtonPrimary},
									w.Button{Label: "Reply all", Icon: "replyall"},
									w.Button{Label: "Forward", Icon: "forward"},
									w.Button{Label: "Archive", Icon: "archive", Title: "Archive"},
									w.Button{Icon: "trash", Title: "Move to Trash"},
									w.Button{Icon: "dots", Title: "More"},
								},
							},
							Pad: w.HeadRule,
						},
						w.DocLayout{
							Items: []w.DocLayoutPart{
								w.Doc{
									Items: []w.DocPart{
										w.DocTitle{Text: "Вёрстка сборника: правки и встреча в четверг"},
										w.Fields{
											Items: []w.FieldsItem{
												{
													Term:  "From",
													Value: []w.ValuePart{w.Text("Анна Ковалёва "), w.Dim{Parts: []w.Inline{w.Text("a.kovaleva@example.org")}}},
												},
												{
													Term:  "To",
													Value: []w.ValuePart{w.Text("Михаил Енученко "), w.Dim{Parts: []w.Inline{w.Text("mix@home.example")}}},
												},
												{
													Term:  "Cc",
													Value: []w.ValuePart{w.Text("Пётр Смирнов "), w.Dim{Parts: []w.Inline{w.Text("p.smirnov@example.org")}}},
												},
												{Term: "Date", Value: []w.ValuePart{w.Text("2026-09-24 10:12")}, Face: w.ValueMono},
											},
											Variant: w.FieldsLetter,
										},
										w.DocRuled{
											Items: []w.DocPart{
												w.Prose{
													Parts: []w.Inline{
														w.Text("Михаил, добрый день!\n\nПрислала правки к третьей вёрстке — они в PDF, отмечены жёлтым. Главное:"),
													},
													Variant: w.ProseRead,
												},
												w.DocBullets{
													Items: [][]w.DocItemPart{
														[]w.DocItemPart{w.Text("на с. 14 и 37 съехали подписи к иллюстрациям;")},
														[]w.DocItemPart{w.Text("в оглавлении не тот номер у главы «Переплёт»;")},
														[]w.DocItemPart{w.Text("обложку оставляем, как на прошлой неделе.")},
													},
												},
												w.Prose{
													Parts: []w.Inline{
														w.Text("В четверг в 11:00 жду вас в издательстве, приглашение во вложении. Схема проезда — на сайте: "),
														w.Link{Text: "Как нас найти", Href: "https://slovo.example/contacts", Address: "slovo.example/contacts"},
														w.Text(". Карточку Петра прикладываю — он возьмёт печать тиража.\n\n—\nАнна Ковалёва\nИздательство «Слово»"),
													},
													Variant: w.ProseRead,
												},
												w.DocQuote{
													Parts: []w.Inline{
														w.Text("2026-09-22 18:40, Михаил Енученко: Анна, третья вёрстка у вас. Посмотрите, пожалуйста, подписи — у меня на двух разворотах они уехали."),
													},
												},
											},
										},
										w.DocAttached{
											Title: w.Rubric{Text: "Attachments", Items: []w.RubricPart{w.Num{Value: 3}}, Variant: w.RubricBelow},
											Items: []w.DocAttachment{
												{
													Icon: "cal",
													Name: "Встреча в издательстве · 2026-09-26 11:00",
													Size: w.Mono{Text: ".ics · 2 KB", Variant: w.MonoDusty11},
													Acts: []w.Button{
														w.Button{Label: "Add to calendar", Icon: "cal", Small: true},
														w.Button{Icon: "download", Small: true, Title: "Download"},
													},
												},
												{
													Icon: "con",
													Name: "Пётр Смирнов",
													Size: w.Mono{Text: ".vcf · 3 KB", Variant: w.MonoDusty11},
													Acts: []w.Button{
														w.Button{Label: "Add to address book", Icon: "con", Small: true},
														w.Button{Icon: "download", Small: true, Title: "Download"},
													},
												},
												{
													Icon: "doc",
													Name: "правки-сборник-v3.pdf",
													Size: w.Mono{Text: "2.4 MB", Variant: w.MonoDusty11},
													Acts: []w.Button{
														w.Button{Label: "Save to files", Icon: "file", Small: true},
														w.Button{Icon: "download", Small: true, Title: "Download"},
													},
												},
											},
										},
									},
								},
								w.Docmeta{
									Sections: []w.DocmetaSec{
										{
											Title: w.Rubric{Text: "Sender"},
											Items: []w.DocmetaPart{
												w.DocRefs{
													Items: []w.DocRef{
														{
															Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11Inline},
															Name: "Анна Ковалёва",
															Href: "#",
															What: w.Dim{Parts: []w.Inline{w.Text("· nextcloud · Контакты")}},
														},
														{Name: "Timeline", Href: "#", What: w.Dim{Parts: []w.Inline{w.Text("· everything with her")}}},
													},
												},
											},
										},
										{
											Title: w.Rubric{Text: "Shown as"},
											Items: []w.DocmetaPart{
												w.DocAttach{Name: "The letter's text part", Size: w.Dim{Parts: []w.Inline{w.Text("· it also has HTML")}}},
												w.SecActs{Items: []w.Button{w.Button{Label: "Show original", Small: true}}},
											},
										},
										{
											Title: w.Rubric{Text: "Source"},
											Items: []w.DocmetaPart{w.Diag{Text: "UID     48213\nFolder  INBOX\nSize    2.6 MB\nRead    2 min ago"}},
										},
									},
								},
							},
						},
					},
				},
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/mail/home/inbox/48213/attachments/1/calendar",
					Icon:   "cal",
					Title:  "Add to calendar",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("from the letter · .ics · 1 event")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
						w.Close{Label: "Close"},
					},
					Body: []w.SheetPart{
						w.Fields{
							Items: []w.FieldsItem{
								{Term: "Event", Value: []w.ValuePart{w.Text("Встреча в издательстве")}},
								{Term: "When", Value: []w.ValuePart{w.Text("2026-09-26 11:00 — 12:00")}, Face: w.ValueMono},
								{Term: "Where", Value: []w.ValuePart{w.Text("Тверская 12, стр. 3, Москва, 123001, Россия")}},
								{
									Term:  "Organiser",
									Value: []w.ValuePart{w.Text("Анна Ковалёва "), w.Dim{Parts: []w.Inline{w.Text("a.kovaleva@example.org")}}},
								},
							},
							Variant: w.FieldsLetter,
						},
						w.Rubric{Text: "Into"},
						w.Tree{
							Rows: []w.TreeRow{
								{Icon: "server", Label: "nextcloud"},
								{
									Mark:   &w.Mark{Colour: "#4A6B52"},
									Label:  "Личный",
									Indent: w.TreeLevel22,
									Send:   w.Send{Name: "into", Value: "nextcloud/lichnyj", One: true},
								},
								{
									Mark:   &w.Mark{Colour: "#3F6E8C"},
									Label:  "Работа",
									Indent: w.TreeLevel22,
									Send:   w.Send{Name: "into", Value: "nextcloud/rabota", One: true},
								},
								{Icon: "server", Label: "radicale · vps"},
								{
									Mark:   &w.Mark{Colour: "#6B4E7A"},
									Label:  "Проекты",
									Indent: w.TreeLevel22,
									Send:   w.Send{Name: "into", Value: "radicale/proekty", One: true},
								},
							},
							On:      3,
							Variant: w.TreeNarrow,
						},
						w.Msg{
							Parts: []w.MsgPart{w.Text("Анна gets no answer from Carrel; if she is waiting for one, reply to her letter.")},
							Tone:  w.MsgAccent,
							Space: w.MsgFlush,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Add to Работа", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGapNormal,
							Place: w.Pushed,
						},
					},
					Width: w.SheetWidth520,
				},
			},
		},
	}
}

// Mailbox12 is frame «Оригинал письма · картинки не загружены» (s-mailbox-12).
func Mailbox12() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 8}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 10}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "mail", Label: "Mail", Count: &w.Num{Value: 12}, Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 6,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{
							Text: "",
							Items: []w.RubricPart{
								w.Line{
									Items: []w.LinePart{w.Mark{Colour: "#A05A2C", Size: w.MarkH11}, w.Text("home")},
									Gap:   w.LineGapNormal,
								},
							},
						},
						w.Tree{
							Rows: []w.TreeRow{
								{Icon: "inbox", Label: "Inbox", Count: &w.Num{Value: 12}},
								{Icon: "pencil", Label: "Drafts", Count: &w.Num{Value: 1}},
								{Icon: "send", Label: "Sent"},
								{Icon: "archive", Label: "Archive"},
								{Icon: "spam", Label: "Spam", Count: &w.Num{Value: 3}},
								{Icon: "trash", Label: "Trash"},
								{Icon: "file", Label: "Стройка", Count: &w.Num{Value: 2}},
								{Icon: "file", Label: "Рассылки"},
							},
							On:      1,
							Variant: w.TreePadded,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{
							Text: "",
							Items: []w.RubricPart{
								w.Line{
									Items: []w.LinePart{w.Mark{Colour: "#3F6E8C", Size: w.MarkH11}, w.Text("office")},
									Gap:   w.LineGapNormal,
								},
							},
						},
					},
					Failed: "no reply · 8 s",
					Why:    "mail.office.example:993 did not answer in 8 s. Other mailboxes and the rest of Carrel work.",
					Retry:  &w.Button{Label: "Retry", Href: "?retry=office"},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "mail", Title: "Connect a mailbox"},
						w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Meta: &w.HeadSub{
							Parts: []w.HeadSubPart{
								w.Mark{Colour: "#A05A2C", Size: w.MarkH12},
								w.Text("home · Inbox · "),
								w.Mono{Text: "2026-09-23 19:30"},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Reply", Icon: "reply", Kind: w.ButtonPrimary},
								w.Button{Label: "Forward", Icon: "forward"},
								w.Button{Label: "Archive", Icon: "archive", Title: "Archive"},
								w.Button{Icon: "trash", Title: "Move to Trash"},
								w.Button{Icon: "dots", Title: "More"},
							},
						},
						Pad: w.HeadRule,
					},
					w.Doc{
						Items: []w.DocPart{
							w.DocTitle{Text: "Сентябрьская рассылка: сорок новинок осени"},
							w.Fields{
								Items: []w.FieldsItem{
									{
										Term:  "From",
										Value: []w.ValuePart{w.Text("Книжная ярмарка "), w.Dim{Parts: []w.Inline{w.Text("news@bookfair.example")}}},
									},
									{Term: "Date", Value: []w.ValuePart{w.Text("2026-09-23 19:30")}, Face: w.ValueMono},
								},
								Variant: w.FieldsLetter,
							},
							w.Line{
								Items: []w.LinePart{
									w.Seg{
										Options: []w.SegOption{
											{Label: "Text", Href: "/app/mail/home/inbox/48190"},
											{Label: "Original", Href: "/app/mail/home/inbox/48190/original", On: true},
										},
									},
									w.Mono{Text: "text/html · 84 KB", Variant: w.MonoDusty11},
								},
								Gap:   w.LineGapNormal,
								Block: true,
							},
							w.Msg{
								Icon: "eye",
								Parts: []w.MsgPart{
									w.Strong("14 pictures from 3 addresses are not loaded."),
									w.Text(" Loading them tells the sender that you opened this letter, when, and your IP address. The choice is for this letter only. "),
									w.Button{Label: "Show pictures", Small: true},
								},
								Space: w.MsgFlushAboveWide,
							},
							w.Original{
								Title: "Original of the letter",
								Doc:   "<body style='margin:0;background:#efe9dc;font-family:Georgia,serif;color:#222'><div style='max-width:560px;margin:0 auto;background:#fff;padding-bottom:18px'><div style='height:92px;margin:0 0 18px;border-bottom:1px dashed #bbb;display:flex;align-items:center;justify-content:center;font:12px Arial,sans-serif;color:#999'>Книжная ярмарка</div><h1 style='font-size:26px;font-weight:normal;margin:0 24px 8px'>Сорок новинок осени</h1><p style='margin:0 24px 18px;font:15px/1.5 Arial,sans-serif;color:#444'>Мы собрали книги, которые выходят в сентябре и октябре: прозу, нон-фикшн и детское. Предзаказ открыт до 5 октября.</p><table style='border-collapse:collapse;margin:0 16px;width:calc(100% - 32px)'><tr><td style='width:50%;padding:8px;vertical-align:top'><div style='height:150px;border:1px dashed #bbb;display:flex;align-items:center;justify-content:center;font:11px Arial,sans-serif;color:#999'>обложка</div><p style='font:bold 14px Arial,sans-serif;margin:8px 0 2px'>Переплётчик</p><p style='font:12px Arial,sans-serif;color:#777;margin:0'>роман, 2026</p></td><td style='width:50%;padding:8px;vertical-align:top'><div style='height:150px;border:1px dashed #bbb;display:flex;align-items:center;justify-content:center;font:11px Arial,sans-serif;color:#999'>обложка</div><p style='font:bold 14px Arial,sans-serif;margin:8px 0 2px'>Как устроена типография</p><p style='font:12px Arial,sans-serif;color:#777;margin:0'>нон-фикшн</p></td></tr></table><a style='display:inline-block;background:#b0392b;color:#fff;padding:10px 18px;margin:16px 24px;font:bold 14px Arial,sans-serif;text-decoration:none'>Смотреть все новинки</a><p style='font:11px Arial,sans-serif;color:#999;margin:0 24px'>Вы получили это письмо, потому что подписались на рассылку. Отписаться</p></div></body>",
							},
						},
					},
				},
			},
		},
	}
}

// Mailbox22 is frame «Подключить почтовый ящик · узнанный провайдер» (s-mailbox-22).
func Mailbox22() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}, Sections: true},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Settings"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "user", Label: "Account", Href: "#"},
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 2,
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
								Title: &w.HeadTitle{Text: "Connect a mailbox"},
								Subs: []w.HeadSub{
									w.HeadSub{
										Parts: []w.HeadSubPart{w.Text("Read it and write from it over IMAP and SMTP. Letters stay on your mail server.")},
									},
								},
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Email address",
										Parts: []w.FieldPart{w.Input{Name: "address", Value: "mix.letters@yandex.ru", Label: "Email address"}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Label",
										Parts: []w.FieldPart{
											w.Input{Name: "label", Value: "yandex", Label: "Label"},
											w.FieldHint{Parts: []w.Inline{w.Text("How this mailbox is named in the rail: «yandex · Inbox».")}},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Colour",
										Parts: []w.FieldPart{
											w.Swatches{
												Colours: []string{"#4A6B52", "#3F6E8C", "#6B4E7A", "#A05A2C", "#8C6239", "#9A9280"},
												On:      2,
												Name:    "colour",
											},
										},
										Width: w.FieldMD,
									},
									w.Msg{
										Parts: []w.MsgPart{
											w.Text("Yandex lets mail programs in with an "),
											w.Strong("app password"),
											w.Text(", not with the password of your Yandex account. Make one in Yandex ID → Security → App passwords → Mail, and paste it here."),
										},
										Tone:  w.MsgAccent,
										Space: w.MsgFlush,
										Place: w.Grown,
									},
									w.Field{
										Label: "App password",
										Parts: []w.FieldPart{
											w.Input{
												Kind:  w.InputPassword,
												Name:  "password",
												Value: "xxxxxxxxxxxxxxxx",
												Label: "App password",
												Focus: true,
											},
										},
										Width: w.FieldMD,
									},
									w.FieldSet{
										Rubric: w.Rubric{Text: "Servers", Items: []w.RubricPart{w.Badge{Label: "found", Tone: w.BadgeLinked}}},
										Items: []w.FieldSetPart{
											w.Diag{
												Text:  "IMAP  imap.yandex.ru:993  TLS\nSMTP  smtp.yandex.ru:465  TLS\n      from autoconfig.yandex.ru",
												Place: w.Grown,
											},
											w.Button{Label: "Enter servers by hand", Small: true, Href: "/app/settings/mailboxes/new?servers=manual"},
										},
									},
									w.Msg{
										Parts: []w.MsgPart{
											w.Strong("Key escrow is on:"),
											w.Text(" the administrator can open this password too — and through a mailbox, reset your other passwords."),
										},
										Tone:  w.MsgAlert,
										Space: w.MsgFlush,
										Place: w.Grown,
									},
								},
								ID:      "mailbox",
								Action:  "/app/settings/mailboxes/new",
								Fields:  []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
								Variant: w.FormFlush,
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Connect", Kind: w.ButtonPrimary, Form: "mailbox"},
									w.Button{Label: "Test only", Form: "mailbox", To: "/app/settings/mailboxes/new/test"},
									w.Button{Label: "Cancel", Href: "/app/settings/connections"},
								},
								Variant: w.FormFootFlush,
							},
						},
						Variant: w.Dialog540,
					},
				},
			},
		},
	}
}

// Mailbox25 is frame «Подключение · чтение работает, отправка нет» (s-mailbox-25).
func Mailbox25() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}, Sections: true},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Settings"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "user", Label: "Account", Href: "#"},
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 2,
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
								Title: &w.HeadTitle{Text: "Connect a mailbox"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("office · m.yenuchenko@office.example")}}},
							},
							w.Msg{
								Parts: []w.MsgPart{
									w.Strong("Saved. Reading works, sending does not:"),
									w.Text(" the server refused the password for sending."),
								},
								Tone:  w.MsgAccent,
								Space: w.MsgFlush,
							},
							w.Sec{
								Items: []w.SecPart{
									w.Rubric{Text: "What was tried"},
									w.Diag{
										Text: "[srv]   _imaps._tcp.office.example       → mail.office.example:993\n[srv]   _submissions._tcp.office.example → none\n[srv]   _submission._tcp.office.example  → mail.office.example:587\n[imap]  mail.office.example:993  TLS · certificate matches\n        LOGIN → OK · CAPABILITY IMAP4rev1 MOVE UIDPLUS SPECIAL-USE CONDSTORE\n[smtp]  mail.office.example:587  EHLO → 250 STARTTLS · SIZE 52428800\n        STARTTLS → 220 · certificate matches\n        AUTH PLAIN → 535 5.7.8 Error: authentication failed",
									},
									w.SecNote{
										Text: "Some servers keep a separate password for sending — ask whoever runs office.example. When it is fixed, Test on the Connections screen checks it without sending anything.",
									},
								},
								Flush: true,
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Done", Kind: w.ButtonPrimary, Href: "/app/settings/connections"},
									w.Button{Label: "Enter servers by hand", Href: "/app/settings/mailboxes/office/edit?servers=manual"},
									w.Button{Label: "Change the password", Href: "/app/settings/mailboxes/office/edit"},
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

// Mailbox26 is frame «Настройки · подключения с ящиками» (s-mailbox-26).
func Mailbox26() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}, Sections: true},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Settings"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "user", Label: "Account", Href: "#"},
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 2,
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
						Title: &w.HeadTitle{Text: "Connections"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{
									w.Text("Signed in as "),
									w.Strong("mix"),
									w.Text(" · 4 servers · 2 mailboxes · 17 collections · 2 subscriptions"),
								},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Connect account", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Connect a mailbox", Icon: "mail"},
								w.Button{Label: "Subscribe to a feed", Icon: "link"},
							},
						},
					},
					w.PostForm{
						ID:     "conn-picks",
						Action: "/app/settings/connections?op=pick",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
					w.MainResults{
						Items: []w.ResultsPart{
							w.Rubric{
								Text:    "DAV servers",
								Items:   []w.RubricPart{w.Mono{Text: "4", Variant: w.MonoTight}},
								Variant: w.RubricSplit,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "nextcloud",
										Items: []w.RubricPart{
											w.Line{
												Items: []w.LinePart{
													w.Button{Label: "New collection", Icon: "plus", Small: true},
													w.Button{Label: "Show password", Small: true},
													w.Button{Label: "Test", Small: true},
													w.Button{
														Label: "Update collection list",
														Small: true,
														Form:  "conn-picks",
														To:    "/app/settings/connections?op=rediscover&account=nextcloud",
													},
													w.Button{Label: "Edit", Small: true},
													w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true},
												},
												Gap: w.LineGapNormal,
											},
										},
									},
									w.Mono{
										Text:    "https://cloud.example.org/remote.php/dav/ · mix · last answer 2 min ago",
										Form:    w.MonoLine,
										Variant: w.Mono11_5Dusty,
									},
								},
							},
							w.Rubric{
								Text:    "Mailboxes",
								Items:   []w.RubricPart{w.Mono{Text: "2", Variant: w.MonoTight}},
								Variant: w.RubricSplit,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "",
										Items: []w.RubricPart{
											w.Line{
												Items: []w.LinePart{w.Mark{Colour: "#A05A2C", Size: w.MarkH11}, w.Text("home")},
												Gap:   w.LineGapNormal,
											},
											w.Line{
												Items: []w.LinePart{
													w.Button{Label: "Show password", Small: true},
													w.Button{Label: "Test", Small: true},
													w.Button{Label: "Edit", Small: true},
													w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true},
												},
												Gap: w.LineGapNormal,
											},
										},
									},
									w.Mono{
										Text:    "mix@home.example · IMAP mail.home.example:993 · SMTP mail.home.example:465 · up to 25 MB · last answer 1 min ago",
										Form:    w.MonoLine,
										Variant: w.Mono11_5Dusty,
									},
									w.Table{
										Cols: []w.TableCol{
											{Label: "Role"},
											{Label: "Folder on the server", Main: true},
											{Label: "How it was found", Rank: w.Rank3},
											{Width: w.ColMenu},
										},
										Rows: []w.TableRow{
											{
												Cells: []w.TableCell{
													{Parts: []w.CellPart{w.TableFile{Icon: "send", Name: "Sent"}}},
													{Parts: []w.CellPart{w.Text("Sent")}, Mono: true},
													{Parts: []w.CellPart{w.Text("announced by the server")}, Dim: true},
													{Parts: []w.CellPart{w.Button{Label: "Change", Small: true}}, Acts: true},
												},
											},
											{
												Cells: []w.TableCell{
													{Parts: []w.CellPart{w.TableFile{Icon: "pencil", Name: "Drafts"}}},
													{Parts: []w.CellPart{w.Text("Drafts")}, Mono: true},
													{Parts: []w.CellPart{w.Text("announced by the server")}, Dim: true},
													{Parts: []w.CellPart{w.Button{Label: "Change", Small: true}}, Acts: true},
												},
											},
											{
												Cells: []w.TableCell{
													{Parts: []w.CellPart{w.TableFile{Icon: "archive", Name: "Archive"}}},
													{Parts: []w.CellPart{w.Text("Архив")}, Mono: true},
													{Parts: []w.CellPart{w.Text("you chose it")}, Dim: true},
													{Parts: []w.CellPart{w.Button{Label: "Change", Small: true}}, Acts: true},
												},
											},
											{
												Cells: []w.TableCell{
													{Parts: []w.CellPart{w.TableFile{Icon: "spam", Name: "Spam"}}},
													{Parts: []w.CellPart{w.Text("Junk")}, Mono: true},
													{Parts: []w.CellPart{w.Text("you chose it")}, Dim: true},
													{Parts: []w.CellPart{w.Button{Label: "Change", Small: true}}, Acts: true},
												},
											},
											{
												Cells: []w.TableCell{
													{Parts: []w.CellPart{w.TableFile{Icon: "trash", Name: "Trash"}}},
													{Parts: []w.CellPart{w.Text("Trash")}, Mono: true},
													{Parts: []w.CellPart{w.Text("announced by the server")}, Dim: true},
													{Parts: []w.CellPart{w.Button{Label: "Change", Small: true}}, Acts: true},
												},
											},
										},
									},
									w.CardRow{
										Kind:  w.CardFoot,
										Items: []w.CardRowPart{w.Check{Parts: []w.CheckPart{w.Text("The server puts a copy of sent letters into Sent itself")}}},
									},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "",
										Items: []w.RubricPart{
											w.Line{
												Items: []w.LinePart{
													w.Mark{Colour: "#3F6E8C", Size: w.MarkH11},
													w.Text("office "),
													w.Badge{Label: "sending does not work", Tone: w.BadgeAlert},
												},
												Gap: w.LineGapNormal,
											},
											w.Line{
												Items: []w.LinePart{
													w.Button{Label: "Show password", Small: true},
													w.Button{Label: "Test again", Small: true},
													w.Button{Label: "Edit", Small: true},
													w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true},
												},
												Gap: w.LineGapNormal,
											},
										},
										Variant: w.RubricDanger,
									},
									w.Mono{
										Text:    "m.yenuchenko@office.example · IMAP mail.office.example:993 · SMTP mail.office.example:587 · reading works",
										Form:    w.MonoLine,
										Variant: w.Mono11_5DustyGap,
									},
									w.Diag{
										Text: "[smtp]  mail.office.example:587  EHLO → 250 · STARTTLS → 220\n        AUTH PLAIN → 535 5.7.8 Error: authentication failed",
									},
									w.CardText{
										Kind: w.CardAlertText,
										Parts: []w.CardTextPart{
											w.Text("Letters are read; none can be sent from this mailbox. In the letter sheet it stays in «From», switched off, until Test passes."),
										},
									},
								},
								Variant: w.CardAlert,
							},
							w.Rubric{
								Text:    "Subscriptions · external, read-only",
								Items:   []w.RubricPart{w.Mono{Text: "2", Variant: w.MonoTight}},
								Variant: w.RubricSplit,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text: "Праздники РФ",
										Items: []w.RubricPart{
											w.Badge{Label: "feed"},
											w.Line{
												Items: []w.LinePart{
													w.Button{Label: "Check", Small: true},
													w.Button{Label: "Edit", Small: true},
													w.Button{Label: "Unsubscribe", Kind: w.ButtonDanger, Small: true},
												},
												Gap: w.LineGapNormal,
											},
										},
									},
									w.Mono{
										Text:    "https://feeds.example/ru-holidays.ics · no password · read 2 min ago",
										Form:    w.MonoLine,
										Variant: w.Mono11_5Dusty,
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

// Mailbox27 is frame «Ящик · правка» (s-mailbox-27).
func Mailbox27() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}, Sections: true},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Settings"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "user", Label: "Account", Href: "#"},
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 6}, Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 2,
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
								Title: &w.HeadTitle{Text: "home"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Mono{Text: "mix@home.example"}, w.Text(" · mailbox")}}},
							},
							w.Form{
								Items: []w.FormPart{
									w.Field{
										Label: "Label",
										Parts: []w.FieldPart{w.Input{Name: "label", Value: "home", Label: "Label"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Colour",
										Parts: []w.FieldPart{
											w.Swatches{
												Colours: []string{"#4A6B52", "#3F6E8C", "#6B4E7A", "#A05A2C", "#8C6239", "#9A9280"},
												On:      3,
												Name:    "colour",
											},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Name in From",
										Parts: []w.FieldPart{
											w.Input{Name: "fromName", Value: "Михаил Енученко", Label: "Name in From"},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Recipients see «Михаил Енученко <mix@home.example>». Your profile name until you change it here."),
												},
											},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Signature",
										Parts: []w.FieldPart{
											w.Area{
												Text:  "Михаил Енученко\nвёрстка и допечатная подготовка",
												Focus: true,
												Name:  "signature",
												Label: "Signature",
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Goes under your text in every letter from home, above the quote, after the line «-- ». Plain text. Empty — no signature."),
												},
											},
										},
										Width: w.FieldFull,
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
									w.FieldSet{
										Rubric: w.Rubric{Text: "Servers"},
										Items: []w.FieldSetPart{
											w.Diag{
												Text:  "IMAP  mail.home.example:993  TLS\nSMTP  mail.home.example:465  TLS\n      found by SRV",
												Place: w.Grown,
											},
											w.Button{
												Label: "Enter servers by hand",
												Small: true,
												Href:  "/app/settings/mailboxes/home/edit?servers=manual",
											},
										},
									},
								},
								ID:      "mailbox",
								Action:  "/app/settings/mailboxes/home/edit",
								Fields:  []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
								Variant: w.FormFlush,
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "mailbox"},
									w.Button{Label: "Test only", Form: "mailbox", To: "/app/settings/mailboxes/home/edit/test"},
									w.Button{Label: "Cancel", Href: "/app/settings/connections"},
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

// Mailbox29 is frame «Администрирование · настройки, свои ящики людей» (s-mailbox-29).
func Mailbox29() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Split{
					Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
					More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More to create"},
				},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to app", Href: "#"}}, Sections: true},
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
											w.FieldHint{
												Parts:   []w.Inline{w.Text("Links in letters start with it. Set with CARREL_PUBLIC_URL.")},
												Variant: w.HintColumn,
											},
										},
										Width: w.FieldW580,
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
												Variant: w.HintColumn,
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
													w.Text("Your relay decides which of the two it lets out. The envelope sender stays "),
													w.Mono{Text: "carrel@slovo.example"},
													w.Text(", so SPF is unaffected; DKIM and DMARC are set up on the relay."),
												},
												Variant: w.HintColumn,
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
								Rubric: w.Rubric{Text: "Users’ own mailboxes", Items: []w.RubricPart{w.Badge{Label: "off"}}},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Check{
												Name:  "mailboxes",
												Value: "on",
												Parts: []w.CheckPart{w.Text("Let users connect their own mailboxes over IMAP and SMTP")},
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Carrel then keeps each person’s mailbox password, usually the most valuable one they have. Their letters go through their own mail server, not the relay above."),
												},
											},
										},
										OwnLine: true,
										Place:   w.Grown,
									},
									w.Msg{
										Parts: []w.MsgPart{
											w.Text("When off, there is no mail anywhere in Carrel. Connected mailboxes are kept and come back when it is on again."),
										},
										Space: w.MsgFlush,
										Place: w.Grown,
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
												Gap: w.LineGapNormal,
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
