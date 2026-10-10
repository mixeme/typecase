// Раздел макета: 7.8 Конфликт.

package screens

import w "carrel.dev/typecase/widget"

// Conflict1 is frame «Конфликт версий · контакт» (s-conflict-1).
func Conflict1() Screen {
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
					{Icon: "back", Label: "Back to list", Short: "Back", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "5", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{
								Title: &w.HeadTitle{Text: "Conflict"},
								Subs: []w.HeadSub{
									w.HeadSub{
										Parts: []w.HeadSubPart{
											w.Text("Another client changed this contact while you were editing it. Nothing has been overwritten."),
										},
									},
								},
							},
							w.DiffAgree{Property: w.Mono{Text: "FN"}, Text: "agrees — kept as it is."},
							w.Diff{
								Items: []w.DiffItem{
									{
										Title: w.Rubric{
											Text:  "",
											Lead:  &w.Mono{Text: "TEL;work"},
											Items: []w.RubricPart{w.RubricAside{Text: "only in your edit"}},
										},
										Options: []w.Check{
											w.Check{
												Name:  "choice.1",
												Value: "mine",
												Radio: true,
												On:    true,
												Parts: []w.CheckPart{w.Choice{Source: "Your edit · include", Value: "+7 495 120-45-00"}},
											},
											w.Check{
												Name:  "choice.1",
												Value: "none",
												Radio: true,
												Parts: []w.CheckPart{w.Choice{Source: "Leave it out", Value: "—"}},
											},
										},
									},
									{
										Title: w.Rubric{
											Text:  "",
											Lead:  &w.Mono{Text: "EMAIL;work"},
											Items: []w.RubricPart{w.RubricAside{Text: "both changed"}},
										},
										Options: []w.Check{
											w.Check{
												Name:  "choice.2",
												Value: "mine",
												Radio: true,
												On:    true,
												Parts: []w.CheckPart{w.Choice{Source: "Your edit", Value: "a.kovaleva@example.org"}},
											},
											w.Check{
												Name:  "choice.2",
												Value: "server",
												Radio: true,
												Parts: []w.CheckPart{w.Choice{Source: "Server", Value: "anna.kovaleva@slovo.example"}},
											},
										},
									},
									{
										Title: w.Rubric{
											Text:  "",
											Lead:  &w.Mono{Text: "NOTE"},
											Items: []w.RubricPart{w.RubricAside{Text: "only in your edit"}},
										},
										Options: []w.Check{
											w.Check{
												Name:  "choice.3",
												Value: "mine",
												Radio: true,
												On:    true,
												Parts: []w.CheckPart{w.Choice{Source: "Your edit · include", Value: "Пишет по вторникам, отвечает в тот же день."}},
											},
											w.Check{
												Name:  "choice.3",
												Value: "none",
												Radio: true,
												Parts: []w.CheckPart{w.Choice{Source: "Leave it out", Value: "—"}},
											},
										},
									},
									{
										Title: w.Rubric{
											Text:  "",
											Lead:  &w.Mono{Text: "X-EVOLUTION-FILE-AS"},
											Items: []w.RubricPart{w.RubricAside{Text: "only on the server"}},
										},
										Options: []w.Check{
											w.Check{
												Name:  "choice.4",
												Value: "mine",
												Radio: true,
												Parts: []w.CheckPart{w.Choice{Source: "Your edit", Value: "—"}},
											},
											w.Check{
												Name:  "choice.4",
												Value: "server",
												Radio: true,
												On:    true,
												Parts: []w.CheckPart{w.Choice{Source: "Server · include", Value: "Ковалёва, Анна"}},
											},
										},
									},
								},
								ID:     "conflict",
								Action: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/conflict",
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Write this combination", Kind: w.ButtonPrimary, Form: "conflict"},
									w.Button{Label: "Use only my edit", Form: "conflict", Name: "all", Value: "mine"},
									w.Button{Label: "Use only the server version", Form: "conflict", Name: "all", Value: "server"},
									w.Button{Label: "Cancel", Href: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7"},
								},
								Variant: w.FormFootWrap,
							},
						},
					},
				},
			},
		},
	}
}

// Conflict3 is frame «Конфликт версий · событие» (s-conflict-3).
func Conflict3() Screen {
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
					{Icon: "back", Label: "Back to list", Short: "Back", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "2", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "2"},
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
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{
								Title: &w.HeadTitle{Text: "Conflict"},
								Subs: []w.HeadSub{
									w.HeadSub{
										Parts: []w.HeadSubPart{w.Text("Another client changed this event while you were editing it. Nothing has been overwritten.")},
									},
								},
							},
							w.DiffAgree{Property: w.Mono{Text: "DTSTART"}, Text: "agrees — kept as it is."},
							w.Diff{
								Items: []w.DiffItem{
									{
										Title: w.Rubric{Text: "", Lead: &w.Mono{Text: "SUMMARY"}, Items: []w.RubricPart{w.RubricAside{Text: "both changed"}}},
										Options: []w.Check{
											w.Check{
												Name:  "choice.1",
												Value: "mine",
												Radio: true,
												On:    true,
												Parts: []w.CheckPart{w.Choice{Source: "Your edit", Value: "Синк-ап по релизу и демо"}},
											},
											w.Check{
												Name:  "choice.1",
												Value: "server",
												Radio: true,
												Parts: []w.CheckPart{w.Choice{Source: "Server", Value: "Синк-ап перенесён"}},
											},
										},
									},
									{
										Title: w.Rubric{
											Text:  "",
											Lead:  &w.Mono{Text: "LOCATION"},
											Items: []w.RubricPart{w.RubricAside{Text: "only in your edit"}},
										},
										Options: []w.Check{
											w.Check{
												Name:  "choice.2",
												Value: "mine",
												Radio: true,
												On:    true,
												Parts: []w.CheckPart{w.Choice{Source: "Your edit · include", Value: "Переговорная 2"}},
											},
											w.Check{
												Name:  "choice.2",
												Value: "none",
												Radio: true,
												Parts: []w.CheckPart{w.Choice{Source: "Leave it out", Value: "—"}},
											},
										},
									},
								},
								ID:     "conflict",
								Action: "/app/calendar/nextcloud/lichnyy/5d1a-77c0/conflict",
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Write this combination", Kind: w.ButtonPrimary, Form: "conflict"},
									w.Button{Label: "Use only my edit", Form: "conflict", Name: "all", Value: "mine"},
									w.Button{Label: "Use only the server version", Form: "conflict", Name: "all", Value: "server"},
									w.Button{Label: "Cancel", Href: "/app/calendar/nextcloud/lichnyy/5d1a-77c0"},
								},
								Variant: w.FormFootWrap,
							},
						},
					},
				},
			},
		},
	}
}

// Conflict4 is frame «Конфликт версий · задача» (s-conflict-4).
func Conflict4() Screen {
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
					{Icon: "back", Label: "Back to list", Short: "Back", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Дела", Meta: "4", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "3"},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{
								Title: &w.HeadTitle{Text: "Conflict"},
								Subs: []w.HeadSub{
									w.HeadSub{
										Parts: []w.HeadSubPart{w.Text("Another client changed this task while you were editing it. Nothing has been overwritten.")},
									},
								},
							},
							w.DiffAgree{Property: w.Mono{Text: "STATUS"}, Text: "agrees — kept as it is."},
							w.Diff{
								Items: []w.DiffItem{
									{
										Title: w.Rubric{Text: "", Lead: &w.Mono{Text: "SUMMARY"}, Items: []w.RubricPart{w.RubricAside{Text: "both changed"}}},
										Options: []w.Check{
											w.Check{
												Name:  "choice.1",
												Value: "mine",
												Radio: true,
												On:    true,
												Parts: []w.CheckPart{w.Choice{Source: "Your edit", Value: "Отправить документы в МФЦ до пятницы"}},
											},
											w.Check{
												Name:  "choice.1",
												Value: "server",
												Radio: true,
												Parts: []w.CheckPart{w.Choice{Source: "Server", Value: "Документы в МФЦ"}},
											},
										},
									},
									{
										Title: w.Rubric{
											Text:  "",
											Lead:  &w.Mono{Text: "DUE"},
											Items: []w.RubricPart{w.RubricAside{Text: "only on the server"}},
										},
										Options: []w.Check{
											w.Check{
												Name:  "choice.2",
												Value: "mine",
												Radio: true,
												Parts: []w.CheckPart{w.Choice{Source: "Your edit", Value: "—"}},
											},
											w.Check{
												Name:  "choice.2",
												Value: "server",
												Radio: true,
												On:    true,
												Parts: []w.CheckPart{w.Choice{Source: "Server · include", Value: "2026-08-14 18:00"}},
											},
										},
									},
								},
								ID:     "conflict",
								Action: "/app/tasks/nextcloud/dela/4c1d-8e20-71fa/conflict",
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Write this combination", Kind: w.ButtonPrimary, Form: "conflict"},
									w.Button{Label: "Use only my edit", Form: "conflict", Name: "all", Value: "mine"},
									w.Button{Label: "Use only the server version", Form: "conflict", Name: "all", Value: "server"},
									w.Button{Label: "Cancel", Href: "/app/tasks/nextcloud/dela/4c1d-8e20-71fa"},
								},
								Variant: w.FormFootWrap,
							},
						},
					},
				},
			},
		},
	}
}

// Conflict5 is frame «Конфликт версий · заметка» (s-conflict-5).
func Conflict5() Screen {
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
					{Icon: "back", Label: "Back to list", Short: "Back", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "2", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Дневник", Meta: "2"},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#3F6E8C"},
							Label:  "Общие с семьёй",
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{
								Title: &w.HeadTitle{Text: "Conflict"},
								Subs: []w.HeadSub{
									w.HeadSub{
										Parts: []w.HeadSubPart{w.Text("Another client changed this note while you were editing it. Nothing has been overwritten.")},
									},
								},
							},
							w.DiffAgree{Property: w.Mono{Text: "CATEGORIES"}, Text: "agrees — kept as it is."},
							w.Diff{
								Items: []w.DiffItem{
									{
										Title: w.Rubric{Text: "", Lead: &w.Mono{Text: "SUMMARY"}, Items: []w.RubricPart{w.RubricAside{Text: "both changed"}}},
										Options: []w.Check{
											w.Check{
												Name:  "choice.1",
												Value: "mine",
												Radio: true,
												On:    true,
												Parts: []w.CheckPart{w.Choice{Source: "Your edit", Value: "Заметка со встречи 10.08 — дополнено"}},
											},
											w.Check{
												Name:  "choice.1",
												Value: "server",
												Radio: true,
												Parts: []w.CheckPart{w.Choice{Source: "Server", Value: "Встреча 10.08"}},
											},
										},
									},
									{
										Title: w.Rubric{
											Text:  "",
											Lead:  &w.Mono{Text: "DESCRIPTION"},
											Items: []w.RubricPart{w.RubricAside{Text: "both changed"}},
										},
										Options: []w.Check{
											w.Check{
												Name:  "choice.2",
												Value: "mine",
												Radio: true,
												On:    true,
												Parts: []w.CheckPart{w.Choice{Source: "Your edit", Value: "Договорились: демо в пятницу, релиз в понедельник."}},
											},
											w.Check{
												Name:  "choice.2",
												Value: "server",
												Radio: true,
												Parts: []w.CheckPart{w.Choice{Source: "Server", Value: "Договорились: демо в пятницу."}},
											},
										},
									},
								},
								ID:     "conflict",
								Action: "/app/notes/nextcloud/zapisi/9e31-40b7-c012/conflict",
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Write this combination", Kind: w.ButtonPrimary, Form: "conflict"},
									w.Button{Label: "Use only my edit", Form: "conflict", Name: "all", Value: "mine"},
									w.Button{Label: "Use only the server version", Form: "conflict", Name: "all", Value: "server"},
									w.Button{Label: "Cancel", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012"},
								},
								Variant: w.FormFootWrap,
							},
						},
					},
				},
			},
		},
	}
}

// Conflict2 is frame «Конфликт · планшет, 768 px» (s-conflict-2).
func Conflict2() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{
				Query:   w.Query{Placeholder: "Search everything…", Label: "Search everything…"},
				Variant: w.SearchFill,
			},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix", NoArrow: true},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Href: "#", Title: "Back"}}},
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			RailForm: w.RailIcons,
			Main: w.Main{
				Items: []w.MainPart{
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{
								Title: &w.HeadTitle{Text: "Conflict"},
								Subs: []w.HeadSub{
									w.HeadSub{
										Parts: []w.HeadSubPart{
											w.Text("Another client changed this contact while you were editing it. Nothing has been overwritten."),
										},
									},
								},
							},
							w.DiffAgree{Property: w.Mono{Text: "FN"}, Text: "agrees — kept as it is."},
							w.Diff{
								Items: []w.DiffItem{
									{
										Title: w.Rubric{
											Text:  "",
											Lead:  &w.Mono{Text: "TEL;work"},
											Items: []w.RubricPart{w.RubricAside{Text: "only in your edit"}},
										},
										Options: []w.Check{
											w.Check{
												Name:  "choice.1",
												Value: "mine",
												Radio: true,
												On:    true,
												Parts: []w.CheckPart{w.Choice{Source: "Your edit · include", Value: "+7 495 120-45-00"}},
											},
											w.Check{
												Name:  "choice.1",
												Value: "none",
												Radio: true,
												Parts: []w.CheckPart{w.Choice{Source: "Leave it out", Value: "—"}},
											},
										},
									},
									{
										Title: w.Rubric{
											Text:  "",
											Lead:  &w.Mono{Text: "EMAIL;work"},
											Items: []w.RubricPart{w.RubricAside{Text: "both changed"}},
										},
										Options: []w.Check{
											w.Check{
												Name:  "choice.2",
												Value: "mine",
												Radio: true,
												On:    true,
												Parts: []w.CheckPart{w.Choice{Source: "Your edit", Value: "a.kovaleva@example.org"}},
											},
											w.Check{
												Name:  "choice.2",
												Value: "server",
												Radio: true,
												Parts: []w.CheckPart{w.Choice{Source: "Server", Value: "anna.kovaleva@slovo.example"}},
											},
										},
									},
									{
										Title: w.Rubric{
											Text:  "",
											Lead:  &w.Mono{Text: "NOTE"},
											Items: []w.RubricPart{w.RubricAside{Text: "only in your edit"}},
										},
										Options: []w.Check{
											w.Check{
												Name:  "choice.3",
												Value: "mine",
												Radio: true,
												On:    true,
												Parts: []w.CheckPart{w.Choice{Source: "Your edit · include", Value: "Пишет по вторникам, отвечает в тот же день."}},
											},
											w.Check{
												Name:  "choice.3",
												Value: "none",
												Radio: true,
												Parts: []w.CheckPart{w.Choice{Source: "Leave it out", Value: "—"}},
											},
										},
									},
									{
										Title: w.Rubric{
											Text:  "",
											Lead:  &w.Mono{Text: "X-EVOLUTION-FILE-AS"},
											Items: []w.RubricPart{w.RubricAside{Text: "only on the server"}},
										},
										Options: []w.Check{
											w.Check{
												Name:  "choice.4",
												Value: "mine",
												Radio: true,
												Parts: []w.CheckPart{w.Choice{Source: "Your edit", Value: "—"}},
											},
											w.Check{
												Name:  "choice.4",
												Value: "server",
												Radio: true,
												On:    true,
												Parts: []w.CheckPart{w.Choice{Source: "Server · include", Value: "Ковалёва, Анна"}},
											},
										},
									},
								},
								ID:     "conflict",
								Action: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/conflict",
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Write this combination", Kind: w.ButtonPrimary, Form: "conflict"},
									w.Button{Label: "Use only my edit", Form: "conflict", Name: "all", Value: "mine"},
									w.Button{Label: "Use only the server version", Form: "conflict", Name: "all", Value: "server"},
									w.Button{Label: "Cancel", Href: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7"},
								},
								Variant: w.FormFootWrap,
							},
						},
					},
				},
			},
		},
	}
}
