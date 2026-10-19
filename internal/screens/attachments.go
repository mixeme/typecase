// Раздел макета: 7.12 Настройки · вложения.

package screens

import w "carrel.dev/typecase/widget"

// Attachments1 is frame «Настройки · вложения» (attachments-1).
func Attachments1() Screen {
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
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Settings"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "user", Label: "Account", Href: "#"},
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 4}, Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 3,
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
						Title: &w.HeadTitle{Text: "Attachments"},
						Subs: []w.HeadSub{
							w.HeadSub{Parts: []w.HeadSubPart{w.Text("Where a file attached to a note, an event or a task is put")}},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Open in Files", Icon: "file"},
								w.Button{Label: "Discard", Href: "/app/settings/attachments"},
								w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "attachments"},
							},
						},
					},
					w.Form{
						Items: []w.FormPart{
							w.FieldSet{
								Rubric: w.Rubric{Text: "Small files stay in the record"},
								Items: []w.FieldSetPart{
									w.Fields{
										Items: []w.FieldsItem{
											{
												Term:  "Up to 2 MB",
												Value: []w.ValuePart{w.Text("written into the event, the task or the note itself, and travels with it")},
											},
											{
												Term:  "Larger",
												Value: []w.ValuePart{w.Text("put in the folder below; the record keeps a link to it")},
											},
										},
										Variant: w.FieldsBleed130,
										Place:   w.Grown,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("A collection that says it takes less ("),
													w.Mono{Text: "max-resource-size"},
													w.Text(") lowers the boundary for itself: "),
													w.Mono{Text: "nextcloud"},
													w.Text(" 2 MB, "),
													w.Mono{Text: "radicale · vps"},
													w.Text(" 1 MB. Nothing here raises it."),
												},
											},
										},
										Place: w.Grown,
									},
								},
								Variant: w.FieldSetLead,
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Folder"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Collection",
										Parts: []w.FieldPart{
											w.Pick{
												Name:    "collection",
												Options: []w.PickOption{{Value: "sftpgo.vps.documents", Label: "sftpgo · vps · Documents", On: true}},
												Label:   "Collection",
											},
											w.FieldHint{Parts: []w.Inline{w.Text("Any file collection you can write to.")}},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Folder in it",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Input{Name: "folder.in.it", Value: "/carrel/attachments", Mono: true, Variant: w.InputFill},
													w.Button{Label: "Browse"},
												},
												Gap: w.LineGap8,
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Created on first use. It stays an ordinary folder: openable in Files, mountable as a network drive, readable without Carrel."),
												},
											},
										},
										Width: w.FieldLG,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.Check{
												On:    true,
												Name:  "a.subfolder.per.month",
												Value: "1",
												Parts: []w.CheckPart{w.Text("A subfolder per month")},
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Mono{Text: "2026-07"},
													w.Text(", "),
													w.Mono{Text: "2026-08"},
													w.Text(". Off means one flat folder — fine for a hundred files, not for ten thousand."),
												},
											},
										},
										Gap:     7,
										OwnLine: true,
										Place:   w.Grown,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "File name"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Written as",
										Parts: []w.FieldPart{
											w.Pick{
												Name:    "written.as",
												Options: []w.PickOption{{Value: "date.then.the.original.name", Label: "Date, then the original name", On: true}},
												Label:   "Written as",
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Mono{Text: "2026-08-10 sync-board.png"},
													w.Text(". A name already taken is never overwritten: a counter is added."),
												},
											},
										},
										Width: w.FieldLG,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "In this folder now"},
								Items: []w.FieldSetPart{
									w.Fields{
										Items: []w.FieldsItem{
											{Term: "Files", Value: []w.ValuePart{w.Text("128 · 1.4 GB")}, Face: w.ValueMono},
											{Term: "Oldest", Value: []w.ValuePart{w.Text("2026-02-19")}, Face: w.ValueMono},
											{
												Term:  "Path",
												Value: []w.ValuePart{w.Text("https://files.example.org/dav/Documents/carrel/attachments/")},
												Face:  w.ValueMonoSmall,
											},
										},
										Variant: w.FieldsBleed130,
										Place:   w.Grown,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Carrel keeps no index of what points at these files. Nothing here deletes them: an attachment is removed from the record it hangs on, in that record."),
												},
											},
										},
										Place: w.Grown,
									},
								},
							},
						},
						ID:     "attachments",
						Action: "/app/settings/attachments",
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("Changing the folder moves nothing. Records already written carry the old path and keep working; only new attachments go to the new place."),
						},
					},
				},
			},
		},
	}
}

// Attachments5 is frame «Вложения · папка не выбрана» (attachments-5).
func Attachments5() Screen {
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
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Settings"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "user", Label: "Account", Href: "#"},
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 4}, Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 3,
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
						Title: &w.HeadTitle{Text: "Attachments"},
						Subs: []w.HeadSub{
							w.HeadSub{Parts: []w.HeadSubPart{w.Text("Where a file attached to a note, an event or a task is put")}},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Open in Files", Icon: "file", Title: "No folder is chosen yet.", State: w.ButtonDisabled},
								w.Button{Label: "Discard", Href: "/app/settings/attachments"},
								w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "attachments"},
							},
						},
					},
					w.Form{
						Items: []w.FormPart{
							w.FieldSet{
								Rubric: w.Rubric{Text: "Small files stay in the record"},
								Items: []w.FieldSetPart{
									w.Fields{
										Items: []w.FieldsItem{
											{
												Term:  "Up to 2 MB",
												Value: []w.ValuePart{w.Text("written into the event, the task or the note itself, and travels with it")},
											},
											{Term: "Larger", Value: []w.ValuePart{w.Text("wait for the folder below")}},
										},
										Variant: w.FieldsBleed130,
										Place:   w.Grown,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("A collection that says it takes less ("),
													w.Mono{Text: "max-resource-size"},
													w.Text(") lowers the boundary for itself: "),
													w.Mono{Text: "nextcloud"},
													w.Text(" 2 MB, "),
													w.Mono{Text: "radicale · vps"},
													w.Text(" 1 MB. Nothing here raises it."),
												},
											},
										},
										Place: w.Grown,
									},
								},
								Variant: w.FieldSetLead,
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Folder", Items: []w.RubricPart{w.Badge{Label: "not set", Tone: w.BadgeAlert}}},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Collection",
										Parts: []w.FieldPart{
											w.Pick{
												Name: "collection",
												Options: []w.PickOption{
													{Value: "", Label: "Choose a collection", On: true},
													{Value: "sftpgo.vps.documents", Label: "sftpgo · vps · Documents"},
												},
												Label: "Collection",
											},
											w.FieldHint{Parts: []w.Inline{w.Text("Any file collection you can write to.")}},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Folder in it",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Input{
														Name:        "folder.in.it",
														Placeholder: "/carrel/attachments",
														Label:       "Folder in it",
														Mono:        true,
														Variant:     w.InputFill,
													},
													w.Button{Label: "Browse"},
												},
												Gap: w.LineGap8,
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Created on first use. It stays an ordinary folder: openable in Files, mountable as a network drive, readable without Carrel."),
												},
											},
										},
										Width: w.FieldLG,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.Check{
												On:    true,
												Name:  "a.subfolder.per.month",
												Value: "1",
												Parts: []w.CheckPart{w.Text("A subfolder per month")},
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Mono{Text: "2026-07"},
													w.Text(", "),
													w.Mono{Text: "2026-08"},
													w.Text(". Off means one flat folder — fine for a hundred files, not for ten thousand."),
												},
											},
										},
										Gap:     7,
										OwnLine: true,
										Place:   w.Grown,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "File name"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Written as",
										Parts: []w.FieldPart{
											w.Pick{
												Name:    "written.as",
												Options: []w.PickOption{{Value: "date.then.the.original.name", Label: "Date, then the original name", On: true}},
												Label:   "Written as",
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Mono{Text: "2026-08-10 sync-board.png"},
													w.Text(". A name already taken is never overwritten: a counter is added."),
												},
											},
										},
										Width: w.FieldLG,
									},
								},
							},
						},
						ID:     "attachments",
						Action: "/app/settings/attachments",
					},
				},
			},
		},
	}
}

// Attachments6 is frame «Вложения · папка недоступна» (attachments-6).
func Attachments6() Screen {
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
					{Icon: "back", Label: "Back to app", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Settings"},
						w.Nav{
							Items: []w.NavItem{
								{Icon: "user", Label: "Account", Href: "#"},
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 4}, Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 3,
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
						Title: &w.HeadTitle{Text: "Attachments"},
						Subs: []w.HeadSub{
							w.HeadSub{Parts: []w.HeadSubPart{w.Text("Where a file attached to a note, an event or a task is put")}},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{
									Label: "Open in Files",
									Icon:  "file",
									Title: "sftpgo · vps refused the password.",
									State: w.ButtonDisabled,
								},
								w.Button{Label: "Discard", Href: "/app/settings/attachments"},
								w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "attachments"},
							},
						},
					},
					w.Form{
						Items: []w.FormPart{
							w.FieldSet{
								Rubric: w.Rubric{Text: "Small files stay in the record"},
								Items: []w.FieldSetPart{
									w.Fields{
										Items: []w.FieldsItem{
											{
												Term:  "Up to 2 MB",
												Value: []w.ValuePart{w.Text("written into the event, the task or the note itself, and travels with it")},
											},
											{
												Term:  "Larger",
												Value: []w.ValuePart{w.Text("put in the folder below; the record keeps a link to it")},
											},
										},
										Variant: w.FieldsBleed130,
										Place:   w.Grown,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("A collection that says it takes less ("),
													w.Mono{Text: "max-resource-size"},
													w.Text(") lowers the boundary for itself: "),
													w.Mono{Text: "nextcloud"},
													w.Text(" 2 MB, "),
													w.Mono{Text: "radicale · vps"},
													w.Text(" 1 MB. Nothing here raises it."),
												},
											},
										},
										Place: w.Grown,
									},
								},
								Variant: w.FieldSetLead,
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Folder"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Collection",
										Parts: []w.FieldPart{
											w.Pick{
												Name:    "collection",
												Options: []w.PickOption{{Value: "sftpgo.vps.documents", Label: "sftpgo · vps · Documents", On: true}},
												Label:   "Collection",
											},
											w.FieldHint{Parts: []w.Inline{w.Text("Any file collection you can write to.")}},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Folder in it",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Input{Name: "folder.in.it", Value: "/carrel/attachments", Mono: true, Variant: w.InputFill},
													w.Button{Label: "Browse"},
												},
												Gap: w.LineGap8,
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Created on first use. It stays an ordinary folder: openable in Files, mountable as a network drive, readable without Carrel."),
												},
											},
										},
										Width: w.FieldLG,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.Check{
												On:    true,
												Name:  "a.subfolder.per.month",
												Value: "1",
												Parts: []w.CheckPart{w.Text("A subfolder per month")},
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Mono{Text: "2026-07"},
													w.Text(", "),
													w.Mono{Text: "2026-08"},
													w.Text(". Off means one flat folder — fine for a hundred files, not for ten thousand."),
												},
											},
										},
										Gap:     7,
										OwnLine: true,
										Place:   w.Grown,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "File name"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Written as",
										Parts: []w.FieldPart{
											w.Pick{
												Name:    "written.as",
												Options: []w.PickOption{{Value: "date.then.the.original.name", Label: "Date, then the original name", On: true}},
												Label:   "Written as",
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Mono{Text: "2026-08-10 sync-board.png"},
													w.Text(". A name already taken is never overwritten: a counter is added."),
												},
											},
										},
										Width: w.FieldLG,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "In this folder now", Items: []w.RubricPart{w.Badge{Label: "unavailable", Tone: w.BadgeAlert}}},
								Items: []w.FieldSetPart{
									w.Fields{
										Items: []w.FieldsItem{
											{
												Term:  "Path",
												Value: []w.ValuePart{w.Text("https://files.example.org/dav/Documents/carrel/attachments/")},
												Face:  w.ValueMonoSmall,
											},
										},
										Variant: w.FieldsBleed130,
										Place:   w.Grown,
									},
									w.Diag{Text: "PROPFIND /dav/Documents/carrel/attachments/ → 401 Unauthorized", Place: w.Grown},
								},
							},
						},
						ID:     "attachments",
						Action: "/app/settings/attachments",
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("Changing the folder moves nothing. Records already written carry the old path and keep working; only new attachments go to the new place."),
						},
					},
				},
			},
		},
	}
}
