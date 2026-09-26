// Раздел макета: 7.8 Конфликт.

package screens

import w "carrel.dev/typecase/widget"

// Conflict1 is frame «Конфликт версий · сборка из двух правок» (s-conflict-1).
func Conflict1() Screen {
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
					{Icon: "back", Label: "Back to list", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "248", Current: true},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "51"},
					},
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
											w.Check{On: true, Parts: []w.CheckPart{w.Choice{Source: "Your edit · include", Value: "+7 495 120-45-00"}}},
											w.Check{Parts: []w.CheckPart{w.Choice{Source: "Leave it out", Value: "—"}}},
										},
									},
									{
										Title: w.Rubric{
											Text:  "",
											Lead:  &w.Mono{Text: "EMAIL;work"},
											Items: []w.RubricPart{w.RubricAside{Text: "both changed"}},
										},
										Options: []w.Check{
											w.Check{On: true, Parts: []w.CheckPart{w.Choice{Source: "Your edit", Value: "a.kovaleva@example.org"}}},
											w.Check{Parts: []w.CheckPart{w.Choice{Source: "Server", Value: "anna.kovaleva@slovo.example"}}},
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
												On:    true,
												Parts: []w.CheckPart{w.Choice{Source: "Your edit · include", Value: "Пишет по вторникам, отвечает в тот же день."}},
											},
											w.Check{Parts: []w.CheckPart{w.Choice{Source: "Leave it out", Value: "—"}}},
										},
									},
									{
										Title: w.Rubric{
											Text:  "",
											Lead:  &w.Mono{Text: "X-EVOLUTION-FILE-AS"},
											Items: []w.RubricPart{w.RubricAside{Text: "only on the server"}},
										},
										Options: []w.Check{
											w.Check{Parts: []w.CheckPart{w.Choice{Source: "Your edit", Value: "—"}}},
											w.Check{On: true, Parts: []w.CheckPart{w.Choice{Source: "Server · include", Value: "Ковалёва, Анна"}}},
										},
									},
								},
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Write this combination", Kind: w.ButtonPrimary},
									w.Button{Label: "Use only my edit"},
									w.Button{Label: "Use only the server version"},
									w.Button{Label: "Cancel"},
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
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}, Variant: w.SearchFill},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary},
				w.UserMenu{Name: "mix", NoArrow: true},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{w.Nav{Items: []w.NavItem{
				{Icon: "back", Href: "#", Title: "Back"}}}},
			RailForm: w.RailIcons,
			Main: w.Main{
				Items: []w.MainPart{
					w.Dialog{
						Items: []w.DialogPart{
							w.Head{
								Title: &w.HeadTitle{Text: "Conflict"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Nothing has been overwritten.")}}},
							},
							w.DiffAgree{Property: w.Mono{Text: "FN"}, Text: "agrees — kept as it is."},
							w.Diff{
								Items: []w.DiffItem{
									{
										Title: w.Rubric{
											Text:  "",
											Lead:  &w.Mono{Text: "EMAIL;work"},
											Items: []w.RubricPart{w.RubricAside{Text: "both changed"}},
										},
										Options: []w.Check{
											w.Check{On: true, Parts: []w.CheckPart{w.Choice{Source: "Your edit", Value: "a.kovaleva@example.org"}}},
											w.Check{Parts: []w.CheckPart{w.Choice{Source: "Server", Value: "anna.kovaleva@slovo.example"}}},
										},
									},
									{
										Title: w.Rubric{
											Text:  "",
											Lead:  &w.Mono{Text: "TEL;work"},
											Items: []w.RubricPart{w.RubricAside{Text: "only in your edit"}},
										},
										Options: []w.Check{
											w.Check{On: true, Parts: []w.CheckPart{w.Choice{Source: "Your edit · include", Value: "+7 495 120-45-00"}}},
											w.Check{Parts: []w.CheckPart{w.Choice{Source: "Leave it out", Value: "—"}}},
										},
									},
								},
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Write this combination", Kind: w.ButtonPrimary},
									w.Button{Label: "Use only my edit"},
									w.Button{Label: "Cancel"},
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
