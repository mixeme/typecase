// Раздел макета: 7.17 Настройки · вид.

package screens

import w "carrel.dev/typecase/widget"

// Appearance1 is frame «Настройки · вид» (appearance-1).
func Appearance1() Screen {
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
							On: 8,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Text:   "davloom · running · 3 devices",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last response"), w.Mono{Text: "2026-08-16 09:41"}}, Size: w.Dim11},
						},
						w.RailStatus{
							Text: "Backups are running",
							Detail: w.Dim{
								Parts: []w.Inline{
									w.Text("Last backup"),
									w.Mono{Text: "2026-08-16 03:10"},
									w.Text("·"),
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
						Title: &w.HeadTitle{Text: "Appearance"},
						Subs: []w.HeadSub{
							w.HeadSub{Parts: []w.HeadSubPart{w.Text("Theme in this browser · week start and clock in the account")}},
						},
					},
					w.Form{
						Items: []w.FormPart{
							w.FieldSet{
								Rubric: w.Rubric{Text: "Theme"},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Seg{
												Options: []w.SegOption{{Label: "Auto", On: true}, {Label: "Light"}, {Label: "Dark"}},
												Variant: w.SegStart,
											},
											w.FieldHint{
												Parts: []w.Inline{w.Text("“Auto” follows the system. Kept in this browser — another device does not see it.")},
											},
										},
									},
								},
								Variant: w.FieldSetLead,
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Density"},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Seg{Options: []w.SegOption{{Label: "Comfortable", On: true}, {Label: "Compact"}}, Variant: w.SegStart},
											w.FieldHint{Parts: []w.Inline{w.Text("List row height: 38 or 32 px. On a phone always 52. This browser.")}},
										},
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Lists"},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Check{On: true, Parts: []w.CheckPart{w.Text("Keep the details panel open")}},
											w.Check{On: true, Parts: []w.CheckPart{w.Text("Show the source column")}},
											w.Check{Parts: []w.CheckPart{w.Text("Open notes in full width")}},
										},
										Gap:   7,
										Place: w.Grown,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{w.Text("Column sets are remembered per section, in this browser — the button below resets them.")},
											},
										},
										Place: w.Grown,
									},
									w.Field{Parts: []w.FieldPart{w.Button{Label: "Reset remembered view settings"}}},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Dates and times"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Week starts on", Parts: []w.FieldPart{w.Pick{Value: "Monday"}}, Width: w.FieldMD},
									w.Field{Label: "Time", Parts: []w.FieldPart{w.Pick{Value: "24-hour"}}, Width: w.FieldMD},
									w.Field{
										Label: "Time zone",
										Parts: []w.FieldPart{w.Input{Value: "Europe/Moscow · this instance", Label: "Time zone"}},
										Width: w.FieldLG,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Week start and 12/24-hour clock live in the account: time is assembled on the server, so another browser sees the same. Time zone is one per instance, set by the administrator — an event shown in two zones to two people is a source of mistakes, not a convenience."),
												},
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
