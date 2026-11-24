// Раздел макета: 7.15 Настройки · резервная копия и восстановление.

package screens

import w "carrel.dev/typecase/widget"

// Backup1 is frame «Резервная копия · задания» (s-backup-1).
func Backup1() Screen {
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
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 4}, Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 6,
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
						Title: &w.HeadTitle{Text: "Backup"},
						Subs: []w.HeadSub{
							w.HeadSub{Parts: []w.HeadSubPart{w.Text("Three jobs · their own sources, destinations and triggers")}},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New job", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Restore", Icon: "restore"},
							},
						},
					},
					w.Table{
						Cols: []w.TableCol{
							{Width: w.ColTick},
							{Label: "Job"},
							{Label: "What it copies", Rank: w.Rank3, Secondary: true},
							{Label: "Where to", Rank: w.Rank4, Secondary: true},
							{Label: "When"},
							{Label: "Keep", Secondary: true},
							{Label: "Last backup", Main: true},
							{Width: w.ColTime},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{Pick: ptr(w.TickedOn)},
									{Parts: []w.CellPart{w.TableFile{Icon: "archive", Name: "Рабочее, ежедневно"}}},
									{
										Parts: []w.CellPart{
											w.Line{
												Items: []w.LinePart{
													w.Mark{Colour: "#3F6E8C"},
													w.Mark{Colour: "#6B4E7A"},
													w.Mark{Colour: "#4A6B52"},
													w.Mono{Text: "3 collections", Variant: w.Mono11Indent},
												},
												Gap: w.LineGapTight,
											},
										},
									},
									{
										Parts: []w.CellPart{w.Text("sftpgo · vps · Documents"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("/backup/work")}}},
										Mono:  true,
									},
									{Parts: []w.CellPart{w.Badge{Label: "trigger"}}},
									{Parts: []w.CellPart{w.Text("14")}, Mono: true},
									{
										Parts: []w.CellPart{w.Text("2026-08-16 03:10"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("18.4 MB · verified")}}},
										Mono:  true,
									},
									{
										Parts: []w.CellPart{
											w.Button{Label: "Run", Small: true, Secondary: true},
											w.Button{Label: "Edit", Small: true, Secondary: true},
										},
										Drop: &w.Drop{
											Trigger: w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
											Row:     true,
											Menu: w.Menu{
												Items: []w.MenuPart{
													w.MenuItem{Label: "Run"},
													w.MenuItem{Label: "Edit"},
													w.MenuItem{Label: "External trigger", Href: "/app/settings/backup/trigger?job=work"},
												},
												Width: w.MenuWidth212,
												Pop:   w.MenuUnderEnd,
											},
										},
										Acts: true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Pick: ptr(w.TickedOn)},
									{Parts: []w.CellPart{w.TableFile{Icon: "archive", Name: "Личное на домашний диск"}}},
									{
										Parts: []w.CellPart{
											w.Line{
												Items: []w.LinePart{
													w.Mark{Colour: "#4A6B52"},
													w.Mark{Colour: "#4A6B52"},
													w.Mark{Colour: "#A05A2C"},
													w.Mark{Colour: "#4A6B52"},
													w.Mono{Text: "4 collections + files", Variant: w.Mono11Indent},
												},
												Gap: w.LineGapTight,
											},
										},
									},
									{
										Parts: []w.CellPart{w.Text("nextcloud · Files"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("/carrel-backup")}}},
										Mono:  true,
									},
									{
										Parts: []w.CellPart{w.Text("on sign-in"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("if older than 7 days")}}},
										Mono:  true,
									},
									{Parts: []w.CellPart{w.Text("8")}, Mono: true},
									{
										Parts: []w.CellPart{w.Text("2026-08-14 09:12"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("2.1 GB · verified")}}},
										Mono:  true,
									},
									{
										Parts: []w.CellPart{
											w.Button{Label: "Run", Small: true, Secondary: true},
											w.Button{Label: "Edit", Small: true, Secondary: true},
										},
										Drop: &w.Drop{
											Trigger: w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
											Row:     true,
											Menu: w.Menu{
												Items: []w.MenuPart{
													w.MenuItem{Label: "Run"},
													w.MenuItem{Label: "Edit"},
													w.MenuItem{Label: "External trigger", Href: "/app/settings/backup/trigger?job=personal"},
												},
												Width: w.MenuWidth212,
												Pop:   w.MenuUnderEnd,
											},
										},
										Acts: true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Pick: ptr(w.TickedOff)},
									{Parts: []w.CellPart{w.TableFile{Icon: "archive", Name: "Contacts archive, by hand"}}},
									{
										Parts: []w.CellPart{
											w.Line{
												Items: []w.LinePart{w.Mark{Colour: "#A05A2C"}, w.Mono{Text: "1 book", Variant: w.Mono11Indent}},
												Gap:   w.LineGapTight,
											},
										},
									},
									{
										Parts: []w.CellPart{w.Text("download in browser"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("nothing is written")}}},
										Mono:  true,
									},
									{Parts: []w.CellPart{w.Text("—")}, Mono: true},
									{Parts: []w.CellPart{w.Text("—")}, Mono: true},
									{Parts: []w.CellPart{w.Text("never")}, Mono: true, Tone: w.CellDusty},
									{
										Parts: []w.CellPart{
											w.Button{Label: "Run", Small: true, Secondary: true},
											w.Button{Label: "Edit", Small: true, Secondary: true},
										},
										Drop: &w.Drop{
											Trigger: w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
											Row:     true,
											Menu: w.Menu{
												Items: []w.MenuPart{
													w.MenuItem{Label: "Run"},
													w.MenuItem{Label: "Edit"},
													w.MenuItem{Label: "External trigger", Href: "/app/settings/backup/trigger?job=contacts"},
												},
												Width: w.MenuWidth212,
												Pop:   w.MenuUnderEnd,
											},
										},
										Acts: true,
									},
								},
								State: w.RowFaded,
							},
						},
						Variant: w.TableWide,
					},
					w.MainHint{
						Text: "verified — the archive was read back from where it was written and the checksum matched; otherwise it is marked red.",
					},
					w.ListGroup{Title: w.Rubric{Text: "History, all jobs", Form: w.RubricInline}, Count: &w.Num{Value: 13}},
					w.Table{
						Cols: []w.TableCol{
							{Label: "When"},
							{Label: "Job"},
							{Label: "Size", Rank: w.Rank3, Secondary: true},
							{Label: "Collections", Rank: w.Rank4},
							{Label: "Verified", Main: true},
							{Label: "Trigger", Secondary: true},
							{Width: w.ColTime},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("2026-08-16 03:10")}, Mono: true},
									{Parts: []w.CellPart{w.Text("Рабочее, ежедневно")}},
									{Parts: []w.CellPart{w.Text("18.4 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("3 / 3")}, Mono: true},
									{Parts: []w.CellPart{w.Text("read back")}},
									{
										Parts: []w.CellPart{w.Text("external"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("cbk_…lamp · 10.0.0.5")}}},
										Mono:  true,
									},
									{
										Parts: []w.CellPart{
											w.Button{Label: "Restore", Icon: "restore", Small: true},
											w.Button{Icon: "download", Small: true, Variant: w.ButtonPadX3, Title: "Download"},
										},
										Acts: true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("2026-08-15 21:40")}, Mono: true},
									{Parts: []w.CellPart{w.Text("Рабочее, ежедневно")}},
									{Parts: []w.CellPart{w.Text("—")}, Mono: true},
									{Parts: []w.CellPart{w.Text("—")}, Mono: true},
									{Parts: []w.CellPart{w.Span{Face: w.FaceDanger, Text: "refused: not this job’s password"}}},
									{
										Parts: []w.CellPart{w.Text("external"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("cbk_…91de · 203.0.113.7")}}},
										Mono:  true,
									},
									{Acts: true},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("2026-08-15 03:10")}, Mono: true},
									{Parts: []w.CellPart{w.Text("Рабочее, ежедневно")}},
									{Parts: []w.CellPart{w.Text("18.2 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2 / 3")}, Mono: true, Tone: w.CellDanger},
									{
										Parts: []w.CellPart{
											w.Text("read back · "),
											w.Span{
												Face:  w.FaceDanger,
												Text:  "1 not copied",
												Title: "baikal · Общий: the server answered 401, so this collection is not in the archive. The other two were read back.",
											},
										},
									},
									{
										Parts: []w.CellPart{w.Text("external"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("cbk_…lamp · 10.0.0.5")}}},
										Mono:  true,
									},
									{
										Parts: []w.CellPart{
											w.Button{Label: "Restore", Icon: "restore", Small: true},
											w.Button{Icon: "download", Small: true, Variant: w.ButtonPadX3, Title: "Download"},
										},
										Acts: true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("2026-08-14 09:12")}, Mono: true},
									{Parts: []w.CellPart{w.Text("Личное на домашний диск")}},
									{Parts: []w.CellPart{w.Text("2.1 GB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("4 / 4 + 1 240 files")}, Mono: true},
									{Parts: []w.CellPart{w.Text("read back")}},
									{Parts: []w.CellPart{w.Text("on sign-in")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Label: "Restore", Icon: "restore", Small: true},
											w.Button{Icon: "download", Small: true, Variant: w.ButtonPadX3, Title: "Download"},
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

// Backup2 is frame «Резервная копия · задание, что и куда» (s-backup-2).
func Backup2() Screen {
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
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 4}, Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 6,
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
						Title: &w.HeadTitle{Text: "Личное на домашний диск"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Job · last copy 2 days ago")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "backup-job"}},
						},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Name",
								Parts: []w.FieldPart{w.Input{Name: "name", Value: "Личное на домашний диск", Label: "Name", Focus: true}},
								Width: w.FieldLG,
							},
							w.Field{
								Parts: []w.FieldPart{w.Check{On: true, Name: "on", Value: "1", Parts: []w.CheckPart{w.Text("Job is on")}}},
							},
							w.FieldSet{
								Rubric: w.Rubric{
									Text:  "What to copy",
									Items: []w.RubricPart{w.Plain{Text: "4 collections and 1 folder chosen · 2.1 GB last time"}},
								},
								Items: []w.FieldSetPart{
									w.FieldGrid{
										Items: []w.FieldGridPart{
											w.FieldGridColumn{
												Items: []w.TickRow{
													w.TickRow{
														Ticked: w.TickedPart,
														Form:   w.TickRowGroup,
														Mark:   w.Mark{Kind: w.MarkMulti},
														Label:  "nextcloud",
														Meta:   "4 of 6",
													},
													w.TickRow{
														Send:  w.Send{Name: "src", Value: "nextcloud|lichnyj"},
														Form:  w.TickRowNested,
														Mark:  w.Mark{Colour: "#4A6B52"},
														Label: "Личный",
														Meta:  "cal.",
													},
													w.TickRow{
														Ticked: w.TickedOff,
														Send:   w.Send{Name: "src", Value: "nextcloud|rabota"},
														Form:   w.TickRowNested,
														Mark:   w.Mark{Colour: "#3F6E8C"},
														Label:  "Работа",
														Meta:   "cal.",
													},
													w.TickRow{
														Ticked: w.TickedOff,
														Send:   w.Send{Name: "src", Value: "nextcloud|dela"},
														Form:   w.TickRowNested,
														Mark:   w.Mark{Colour: "#4A6B52"},
														Label:  "Дела",
														Meta:   "tasks",
													},
													w.TickRow{
														Send:  w.Send{Name: "src", Value: "nextcloud|kontakty"},
														Form:  w.TickRowNested,
														Mark:  w.Mark{Colour: "#4A6B52"},
														Label: "Контакты",
														Meta:  "book",
													},
													w.TickRow{
														Send:  w.Send{Name: "src", Value: "nextcloud|arhiv"},
														Form:  w.TickRowNested,
														Mark:  w.Mark{Colour: "#A05A2C"},
														Label: "Архив",
														Meta:  "book",
													},
													w.TickRow{
														Send:  w.Send{Name: "src", Value: "nextcloud|zapisi"},
														Form:  w.TickRowNested,
														Mark:  w.Mark{Colour: "#4A6B52"},
														Label: "Записи",
														Meta:  "notes",
													},
												},
											},
											w.FieldGridColumn{
												Items: []w.TickRow{
													w.TickRow{
														Ticked: w.TickedOff,
														Form:   w.TickRowGroup,
														Mark:   w.Mark{Kind: w.MarkMulti},
														Label:  "radicale · vps",
														Meta:   "0 of 2",
													},
													w.TickRow{
														Ticked: w.TickedOff,
														Send:   w.Send{Name: "src", Value: "radicale|proekty"},
														Form:   w.TickRowNested,
														Mark:   w.Mark{Colour: "#6B4E7A"},
														Label:  "Проекты",
														Meta:   "cal.",
													},
													w.TickRow{
														Ticked: w.TickedOff,
														Send:   w.Send{Name: "src", Value: "radicale|dezhurstva"},
														Form:   w.TickRowNested,
														Mark:   w.Mark{Colour: "#8C6239"},
														Label:  "Дежурства",
														Meta:   "cal.",
													},
												},
											},
											w.FieldGridColumn{
												Items: []w.TickRow{
													w.TickRow{
														Ticked: w.TickedOff,
														Form:   w.TickRowGroup,
														Mark:   w.Mark{Kind: w.MarkMulti},
														Label:  "baikal",
														Meta:   "0 of 1",
													},
													w.TickRow{
														Ticked: w.TickedOff,
														Send:   w.Send{Name: "src", Value: "baikal|obshchij"},
														Form:   w.TickRowNested,
														Mark:   w.Mark{Colour: "#9A9280"},
														Label:  "Общий",
														Meta:   "cal.",
													},
												},
											},
											w.FieldGridColumn{
												Items: []w.TickRow{
													w.TickRow{
														Ticked: w.TickedPart,
														Form:   w.TickRowGroup,
														Mark:   w.Mark{Kind: w.MarkMulti},
														Label:  "sftpgo · vps",
														Meta:   "1 of 2",
													},
													w.TickRow{
														Send:  w.Send{Name: "src", Value: "sftpgo|documents"},
														Form:  w.TickRowNested,
														Mark:  w.Mark{Colour: "#4A6B52"},
														Label: "Documents",
														Meta:  "files · 2.0 GB",
													},
													w.TickRow{
														Ticked: w.TickedOff,
														Send:   w.Send{Name: "src", Value: "sftpgo|public"},
														Form:   w.TickRowNested,
														Mark:   w.Mark{Colour: "#9A9280"},
														Label:  "Public",
														Meta:   "files · ro",
													},
												},
											},
										},
									},
									w.Msg{
										Parts: []w.MsgPart{
											w.Text("Address books go as "),
											w.Mono{Text: ".vcf"},
											w.Text("; calendars, tasks and notes as "),
											w.Mono{Text: ".ics"},
											w.Text("; files as files."),
										},
										Space:  w.MsgFlushAboveTight,
										Place:  w.Grown,
										Column: true,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Where to"},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Seg{
												Options: []w.SegOption{
													{Label: "To WebDAV", On: true, Value: "webdav"},
													{Label: "Download", Value: "download"},
													{Label: "Both", Value: "both"},
												},
												Variant: w.SegStart,
												Name:    "mode",
											},
										},
										OwnLine: true,
										Place:   w.Grown,
									},
									w.Field{
										Label: "Collection",
										Parts: []w.FieldPart{
											w.Drop{
												Trigger: w.Pick{Value: "nextcloud · Files", Mark: &w.Mark{Colour: "#4A6B52"}},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#4A6B52"},
															Label: "nextcloud · Files",
															On:    true,
															Name:  "dest",
															Value: "nextcloud|files",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#3F6E8C"},
															Label: "sftpgo · vps · Documents",
															Name:  "dest",
															Value: "sftpgo|documents",
														},
														w.MenuItem{
															Mark:   &w.Mark{Colour: "#9A9280"},
															Trail:  w.Badge{Label: "ro", Place: w.Pushed},
															Label:  "sftpgo · vps · Public files",
															Reason: "Read-only: this account cannot write here",
															Name:   "dest",
															Value:  "sftpgo|public",
														},
													},
													Width: w.MenuWidth290,
													Pop:   w.MenuUnder,
												},
											},
										},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Folder in it",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Input{Name: "folder", Value: "carrel-backup", Label: "Folder in it", Mono: true, Variant: w.InputFill},
													w.Button{Label: "Browse", To: "/app/settings/backup?op=browse&job=personal", Submit: true},
												},
												Gap: w.LineGapNormal,
											},
										},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Keep copies",
										Parts: []w.FieldPart{w.Input{Name: "keep", Value: "8", Label: "Keep copies", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts:   []w.Inline{w.Text("Older copies are deleted only "), w.Strong("after"), w.Text(" the new one is checked.")},
												Variant: w.HintColumn,
											},
										},
										Place: w.Grown,
									},
									w.Msg{
										Parts: []w.MsgPart{
											w.Strong("Download"),
											w.Text(" writes nothing to the server. On a schedule it runs only from an external trigger: someone has to take the file. "),
											w.Strong("Both"),
											w.Text(" does the two, and its trigger has "),
											w.Strong("two steps"),
											w.Text("."),
										},
										Space:  w.MsgFlushAboveTight,
										Place:  w.Grown,
										Column: true,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "How the archive is closed"},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Check{On: true, Name: "encrypted", Value: "1", Parts: []w.CheckPart{w.Text("Encrypt the archive")}},
											w.FieldHint{
												Parts:   []w.Inline{w.Text("Turn it off only where nobody but you can reach the copy.")},
												Variant: w.HintColumn,
											},
										},
										OwnLine: true,
										Place:   w.Grown,
									},
									w.Field{
										Label: "Archive password",
										Parts: []w.FieldPart{
											w.Input{Kind: w.InputPassword, Name: "password", Placeholder: "••••••••••••", Label: "Archive password"},
											w.FieldHint{Parts: []w.Inline{w.Text("Separate from the sign-in password.")}, Variant: w.HintColumn},
										},
										Width: w.FieldLG,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "When"},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Check{
														On:    true,
														Name:  "signin.on",
														Value: "on",
														Parts: []w.CheckPart{w.Text("On sign-in, if the last copy is older than")},
													},
													w.Input{Name: "signin", Value: "7", Label: "Days", Mono: true, Variant: w.InputChip},
													w.Faced{Face: w.FaceLarge, Items: []w.LinePart{w.Text("days")}},
												},
												Gap:   w.LineGapNormal,
												Block: true,
											},
										},
										Gapped: true,
										Place:  w.Grown,
									},
									w.Msg{
										Parts: []w.MsgPart{
											w.Strong("Also started by"),
											w.Break{},
											w.Text("• Run now — in the job list"),
											w.Break{},
											w.Text("• Your own clock — "),
											w.Mono{Text: "cron"},
											w.Text(", GoSentry, a job in Gitea, "),
											w.Mono{Text: "curl"},
											w.Text(": External trigger in the job’s menu"),
										},
										Space:  w.MsgFlushAboveTight,
										Place:  w.Grown,
										Column: true,
									},
								},
							},
						},
						ID:     "backup-job",
						Action: "/app/settings/backup?op=save&job=personal",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
					w.Danger{
						Button: w.Button{Label: "Delete job", Kind: w.ButtonDangerSolid},
						Note:   "Archives already made stay where they are: the schedule is deleted, not the copies.",
					},
				},
			},
		},
	}
}

// Backup19 is frame «Резервная копия · новое задание» (s-backup-19).
func Backup19() Screen {
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
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 4}, Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 6,
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
						Title: &w.HeadTitle{Text: "New job"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Nothing is copied until it is created")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Create", Kind: w.ButtonPrimary, Form: "backup-job"}},
						},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Name",
								Parts: []w.FieldPart{w.Input{Name: "name", Label: "Name", Focus: true, Required: true}},
								Width: w.FieldLG,
							},
							w.Field{
								Parts: []w.FieldPart{w.Check{On: true, Name: "on", Value: "1", Parts: []w.CheckPart{w.Text("Job is on")}}},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "What to copy", Items: []w.RubricPart{w.Plain{Text: "nothing chosen"}}},
								Items: []w.FieldSetPart{
									w.FieldGrid{
										Items: []w.FieldGridPart{
											w.FieldGridColumn{
												Items: []w.TickRow{
													w.TickRow{
														Ticked: w.TickedOff,
														Form:   w.TickRowGroup,
														Mark:   w.Mark{Kind: w.MarkMulti},
														Label:  "nextcloud",
														Meta:   "0 of 6",
													},
													w.TickRow{
														Ticked: w.TickedOff,
														Send:   w.Send{Name: "src", Value: "nextcloud|lichnyj"},
														Form:   w.TickRowNested,
														Mark:   w.Mark{Colour: "#4A6B52"},
														Label:  "Личный",
														Meta:   "cal.",
													},
													w.TickRow{
														Ticked: w.TickedOff,
														Send:   w.Send{Name: "src", Value: "nextcloud|rabota"},
														Form:   w.TickRowNested,
														Mark:   w.Mark{Colour: "#3F6E8C"},
														Label:  "Работа",
														Meta:   "cal.",
													},
													w.TickRow{
														Ticked: w.TickedOff,
														Send:   w.Send{Name: "src", Value: "nextcloud|dela"},
														Form:   w.TickRowNested,
														Mark:   w.Mark{Colour: "#4A6B52"},
														Label:  "Дела",
														Meta:   "tasks",
													},
													w.TickRow{
														Ticked: w.TickedOff,
														Send:   w.Send{Name: "src", Value: "nextcloud|kontakty"},
														Form:   w.TickRowNested,
														Mark:   w.Mark{Colour: "#4A6B52"},
														Label:  "Контакты",
														Meta:   "book",
													},
													w.TickRow{
														Ticked: w.TickedOff,
														Send:   w.Send{Name: "src", Value: "nextcloud|arhiv"},
														Form:   w.TickRowNested,
														Mark:   w.Mark{Colour: "#A05A2C"},
														Label:  "Архив",
														Meta:   "book",
													},
													w.TickRow{
														Ticked: w.TickedOff,
														Send:   w.Send{Name: "src", Value: "nextcloud|zapisi"},
														Form:   w.TickRowNested,
														Mark:   w.Mark{Colour: "#4A6B52"},
														Label:  "Записи",
														Meta:   "notes",
													},
												},
											},
											w.FieldGridColumn{
												Items: []w.TickRow{
													w.TickRow{
														Ticked: w.TickedOff,
														Form:   w.TickRowGroup,
														Mark:   w.Mark{Kind: w.MarkMulti},
														Label:  "radicale · vps",
														Meta:   "0 of 2",
													},
													w.TickRow{
														Ticked: w.TickedOff,
														Send:   w.Send{Name: "src", Value: "radicale|proekty"},
														Form:   w.TickRowNested,
														Mark:   w.Mark{Colour: "#6B4E7A"},
														Label:  "Проекты",
														Meta:   "cal.",
													},
													w.TickRow{
														Ticked: w.TickedOff,
														Send:   w.Send{Name: "src", Value: "radicale|dezhurstva"},
														Form:   w.TickRowNested,
														Mark:   w.Mark{Colour: "#8C6239"},
														Label:  "Дежурства",
														Meta:   "cal.",
													},
												},
											},
											w.FieldGridColumn{
												Items: []w.TickRow{
													w.TickRow{
														Ticked: w.TickedOff,
														Form:   w.TickRowGroup,
														Mark:   w.Mark{Kind: w.MarkMulti},
														Label:  "baikal",
														Meta:   "0 of 1",
													},
													w.TickRow{
														Ticked: w.TickedOff,
														Send:   w.Send{Name: "src", Value: "baikal|obshchij"},
														Form:   w.TickRowNested,
														Mark:   w.Mark{Colour: "#9A9280"},
														Label:  "Общий",
														Meta:   "cal.",
													},
												},
											},
											w.FieldGridColumn{
												Items: []w.TickRow{
													w.TickRow{
														Ticked: w.TickedOff,
														Form:   w.TickRowGroup,
														Mark:   w.Mark{Kind: w.MarkMulti},
														Label:  "sftpgo · vps",
														Meta:   "0 of 2",
													},
													w.TickRow{
														Ticked: w.TickedOff,
														Send:   w.Send{Name: "src", Value: "sftpgo|documents"},
														Form:   w.TickRowNested,
														Mark:   w.Mark{Colour: "#4A6B52"},
														Label:  "Documents",
														Meta:   "files · 2.0 GB",
													},
													w.TickRow{
														Ticked: w.TickedOff,
														Send:   w.Send{Name: "src", Value: "sftpgo|public"},
														Form:   w.TickRowNested,
														Mark:   w.Mark{Colour: "#9A9280"},
														Label:  "Public",
														Meta:   "files · ro",
													},
												},
											},
										},
									},
									w.Msg{
										Parts: []w.MsgPart{
											w.Text("Address books go as "),
											w.Mono{Text: ".vcf"},
											w.Text("; calendars, tasks and notes as "),
											w.Mono{Text: ".ics"},
											w.Text("; files as files."),
										},
										Space:  w.MsgFlushAboveTight,
										Place:  w.Grown,
										Column: true,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Where to"},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Seg{
												Options: []w.SegOption{
													{Label: "To WebDAV", On: true, Value: "webdav"},
													{Label: "Download", Value: "download"},
													{Label: "Both", Value: "both"},
												},
												Variant: w.SegStart,
												Name:    "mode",
											},
										},
										OwnLine: true,
										Place:   w.Grown,
									},
									w.Field{
										Label: "Collection",
										Parts: []w.FieldPart{
											w.Drop{
												Trigger: w.Pick{Value: "nextcloud · Files", Mark: &w.Mark{Colour: "#4A6B52"}},
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#4A6B52"},
															Label: "nextcloud · Files",
															On:    true,
															Name:  "dest",
															Value: "nextcloud|files",
														},
														w.MenuItem{
															Mark:  &w.Mark{Colour: "#3F6E8C"},
															Label: "sftpgo · vps · Documents",
															Name:  "dest",
															Value: "sftpgo|documents",
														},
														w.MenuItem{
															Mark:   &w.Mark{Colour: "#9A9280"},
															Trail:  w.Badge{Label: "ro", Place: w.Pushed},
															Label:  "sftpgo · vps · Public files",
															Reason: "Read-only: this account cannot write here",
															Name:   "dest",
															Value:  "sftpgo|public",
														},
													},
													Width: w.MenuWidth290,
													Pop:   w.MenuUnder,
												},
											},
										},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Folder in it",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Input{Name: "folder", Value: "carrel-backup", Label: "Folder in it", Mono: true, Variant: w.InputFill},
													w.Button{Label: "Browse", To: "/app/settings/backup?op=browse&job=personal", Submit: true},
												},
												Gap: w.LineGapNormal,
											},
										},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Keep copies",
										Parts: []w.FieldPart{w.Input{Name: "keep", Value: "8", Label: "Keep copies", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts:   []w.Inline{w.Text("Older copies are deleted only "), w.Strong("after"), w.Text(" the new one is checked.")},
												Variant: w.HintColumn,
											},
										},
										Place: w.Grown,
									},
									w.Msg{
										Parts: []w.MsgPart{
											w.Strong("Download"),
											w.Text(" writes nothing to the server. On a schedule it runs only from an external trigger: someone has to take the file. "),
											w.Strong("Both"),
											w.Text(" does the two, and its trigger has "),
											w.Strong("two steps"),
											w.Text("."),
										},
										Space:  w.MsgFlushAboveTight,
										Place:  w.Grown,
										Column: true,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "How the archive is closed"},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Check{On: true, Name: "encrypted", Value: "1", Parts: []w.CheckPart{w.Text("Encrypt the archive")}},
											w.FieldHint{
												Parts:   []w.Inline{w.Text("Turn it off only where nobody but you can reach the copy.")},
												Variant: w.HintColumn,
											},
										},
										OwnLine: true,
										Place:   w.Grown,
									},
									w.Field{
										Label: "Archive password",
										Parts: []w.FieldPart{
											w.Input{Kind: w.InputPassword, Name: "password", Label: "Archive password", Required: true},
											w.FieldHint{Parts: []w.Inline{w.Text("Separate from the sign-in password.")}, Variant: w.HintColumn},
										},
										Width: w.FieldLG,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "When"},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Check{
														On:    true,
														Name:  "signin.on",
														Value: "on",
														Parts: []w.CheckPart{w.Text("On sign-in, if the last copy is older than")},
													},
													w.Input{Name: "signin", Value: "7", Label: "Days", Mono: true, Variant: w.InputChip},
													w.Faced{Face: w.FaceLarge, Items: []w.LinePart{w.Text("days")}},
												},
												Gap:   w.LineGapNormal,
												Block: true,
											},
										},
										Gapped: true,
										Place:  w.Grown,
									},
									w.Msg{
										Parts: []w.MsgPart{
											w.Strong("Also started by"),
											w.Break{},
											w.Text("• Run now — in the job list"),
											w.Break{},
											w.Text("• Your own clock — "),
											w.Mono{Text: "cron"},
											w.Text(", GoSentry, a job in Gitea, "),
											w.Mono{Text: "curl"},
											w.Text(": External trigger in the job’s menu"),
										},
										Space:  w.MsgFlushAboveTight,
										Place:  w.Grown,
										Column: true,
									},
								},
							},
						},
						ID:     "backup-job",
						Action: "/app/settings/backup?op=create",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
				},
			},
		},
	}
}

// Backup12 is frame «Внешний запуск · включён» (s-backup-12).
func Backup12() Screen {
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
					w.Nav{Items: []w.NavItem{
						{Icon: "back", Label: "Back to app", Href: "#"}}, Sections: true},
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
								On: 6,
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
							Title: &w.HeadTitle{Text: "Backup"},
							Subs: []w.HeadSub{
								w.HeadSub{Parts: []w.HeadSubPart{w.Text("Three jobs · their own sources, destinations and triggers")}},
							},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "New job", Icon: "plus", Kind: w.ButtonPrimary},
									w.Button{Label: "Restore", Icon: "restore"},
								},
							},
						},
						w.Table{
							Cols: []w.TableCol{
								{Width: w.ColTick},
								{Label: "Job"},
								{Label: "What it copies", Rank: w.Rank3, Secondary: true},
								{Label: "Where to", Rank: w.Rank4, Secondary: true},
								{Label: "When"},
								{Label: "Keep", Secondary: true},
								{Label: "Last backup", Main: true},
								{Width: w.ColTime},
							},
							Rows: []w.TableRow{
								{
									Cells: []w.TableCell{
										{Pick: ptr(w.TickedOn)},
										{Parts: []w.CellPart{w.TableFile{Icon: "archive", Name: "Рабочее, ежедневно"}}},
										{
											Parts: []w.CellPart{
												w.Line{
													Items: []w.LinePart{
														w.Mark{Colour: "#3F6E8C"},
														w.Mark{Colour: "#6B4E7A"},
														w.Mark{Colour: "#4A6B52"},
														w.Mono{Text: "3 collections", Variant: w.Mono11Indent},
													},
													Gap: w.LineGapTight,
												},
											},
										},
										{
											Parts: []w.CellPart{w.Text("sftpgo · vps · Documents"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("/backup/work")}}},
											Mono:  true,
										},
										{Parts: []w.CellPart{w.Badge{Label: "trigger"}}},
										{Parts: []w.CellPart{w.Text("14")}, Mono: true},
										{
											Parts: []w.CellPart{w.Text("2026-08-16 03:10"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("18.4 MB · verified")}}},
											Mono:  true,
										},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Run", Small: true, Secondary: true},
												w.Button{Label: "Edit", Small: true, Secondary: true},
											},
											Drop: &w.Drop{
												Trigger: w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
												Row:     true,
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{Label: "Run"},
														w.MenuItem{Label: "Edit"},
														w.MenuItem{Label: "External trigger", Href: "/app/settings/backup/trigger?job=work"},
													},
													Width: w.MenuWidth212,
													Pop:   w.MenuUnderEnd,
												},
											},
											Acts: true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Pick: ptr(w.TickedOn)},
										{Parts: []w.CellPart{w.TableFile{Icon: "archive", Name: "Личное на домашний диск"}}},
										{
											Parts: []w.CellPart{
												w.Line{
													Items: []w.LinePart{
														w.Mark{Colour: "#4A6B52"},
														w.Mark{Colour: "#4A6B52"},
														w.Mark{Colour: "#A05A2C"},
														w.Mark{Colour: "#4A6B52"},
														w.Mono{Text: "4 collections + files", Variant: w.Mono11Indent},
													},
													Gap: w.LineGapTight,
												},
											},
										},
										{
											Parts: []w.CellPart{w.Text("nextcloud · Files"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("/carrel-backup")}}},
											Mono:  true,
										},
										{
											Parts: []w.CellPart{w.Text("on sign-in"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("if older than 7 days")}}},
											Mono:  true,
										},
										{Parts: []w.CellPart{w.Text("8")}, Mono: true},
										{
											Parts: []w.CellPart{w.Text("2026-08-14 09:12"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("2.1 GB · verified")}}},
											Mono:  true,
										},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Run", Small: true, Secondary: true},
												w.Button{Label: "Edit", Small: true, Secondary: true},
											},
											Drop: &w.Drop{
												Trigger: w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
												Row:     true,
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{Label: "Run"},
														w.MenuItem{Label: "Edit"},
														w.MenuItem{Label: "External trigger", Href: "/app/settings/backup/trigger?job=personal"},
													},
													Width: w.MenuWidth212,
													Pop:   w.MenuUnderEnd,
												},
											},
											Acts: true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Pick: ptr(w.TickedOff)},
										{Parts: []w.CellPart{w.TableFile{Icon: "archive", Name: "Contacts archive, by hand"}}},
										{
											Parts: []w.CellPart{
												w.Line{
													Items: []w.LinePart{w.Mark{Colour: "#A05A2C"}, w.Mono{Text: "1 book", Variant: w.Mono11Indent}},
													Gap:   w.LineGapTight,
												},
											},
										},
										{
											Parts: []w.CellPart{w.Text("download in browser"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("nothing is written")}}},
											Mono:  true,
										},
										{Parts: []w.CellPart{w.Text("—")}, Mono: true},
										{Parts: []w.CellPart{w.Text("—")}, Mono: true},
										{Parts: []w.CellPart{w.Text("never")}, Mono: true, Tone: w.CellDusty},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Run", Small: true, Secondary: true},
												w.Button{Label: "Edit", Small: true, Secondary: true},
											},
											Drop: &w.Drop{
												Trigger: w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
												Row:     true,
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{Label: "Run"},
														w.MenuItem{Label: "Edit"},
														w.MenuItem{Label: "External trigger", Href: "/app/settings/backup/trigger?job=contacts"},
													},
													Width: w.MenuWidth212,
													Pop:   w.MenuUnderEnd,
												},
											},
											Acts: true,
										},
									},
									State: w.RowFaded,
								},
							},
							Variant: w.TableWide,
						},
						w.MainHint{
							Text: "verified — the archive was read back from where it was written and the checksum matched; otherwise it is marked red.",
						},
						w.ListGroup{Title: w.Rubric{Text: "History, all jobs", Form: w.RubricInline}, Count: &w.Num{Value: 13}},
						w.Table{
							Cols: []w.TableCol{
								{Label: "When"},
								{Label: "Job"},
								{Label: "Size", Rank: w.Rank3, Secondary: true},
								{Label: "Collections", Rank: w.Rank4},
								{Label: "Verified", Main: true},
								{Label: "Trigger", Secondary: true},
								{Width: w.ColTime},
							},
							Rows: []w.TableRow{
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("2026-08-16 03:10")}, Mono: true},
										{Parts: []w.CellPart{w.Text("Рабочее, ежедневно")}},
										{Parts: []w.CellPart{w.Text("18.4 MB")}, Mono: true},
										{Parts: []w.CellPart{w.Text("3 / 3")}, Mono: true},
										{Parts: []w.CellPart{w.Text("read back")}},
										{
											Parts: []w.CellPart{w.Text("external"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("cbk_…lamp · 10.0.0.5")}}},
											Mono:  true,
										},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Restore", Icon: "restore", Small: true},
												w.Button{Icon: "download", Small: true, Variant: w.ButtonPadX3, Title: "Download"},
											},
											Acts: true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("2026-08-15 21:40")}, Mono: true},
										{Parts: []w.CellPart{w.Text("Рабочее, ежедневно")}},
										{Parts: []w.CellPart{w.Text("—")}, Mono: true},
										{Parts: []w.CellPart{w.Text("—")}, Mono: true},
										{Parts: []w.CellPart{w.Span{Face: w.FaceDanger, Text: "refused: not this job’s password"}}},
										{
											Parts: []w.CellPart{w.Text("external"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("cbk_…91de · 203.0.113.7")}}},
											Mono:  true,
										},
										{Acts: true},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("2026-08-15 03:10")}, Mono: true},
										{Parts: []w.CellPart{w.Text("Рабочее, ежедневно")}},
										{Parts: []w.CellPart{w.Text("18.2 MB")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2 / 3")}, Mono: true, Tone: w.CellDanger},
										{
											Parts: []w.CellPart{
												w.Text("read back · "),
												w.Span{
													Face:  w.FaceDanger,
													Text:  "1 not copied",
													Title: "baikal · Общий: the server answered 401, so this collection is not in the archive. The other two were read back.",
												},
											},
										},
										{
											Parts: []w.CellPart{w.Text("external"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("cbk_…lamp · 10.0.0.5")}}},
											Mono:  true,
										},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Restore", Icon: "restore", Small: true},
												w.Button{Icon: "download", Small: true, Variant: w.ButtonPadX3, Title: "Download"},
											},
											Acts: true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("2026-08-14 09:12")}, Mono: true},
										{Parts: []w.CellPart{w.Text("Личное на домашний диск")}},
										{Parts: []w.CellPart{w.Text("2.1 GB")}, Mono: true},
										{Parts: []w.CellPart{w.Text("4 / 4 + 1 240 files")}, Mono: true},
										{Parts: []w.CellPart{w.Text("read back")}},
										{Parts: []w.CellPart{w.Text("on sign-in")}, Mono: true},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Restore", Icon: "restore", Small: true},
												w.Button{Icon: "download", Small: true, Variant: w.ButtonPadX3, Title: "Download"},
											},
											Acts: true,
										},
									},
								},
							},
						},
					},
				},
				Height: w.ShellHeight560,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/settings/backup/trigger?job=work",
					Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					Title:  "External trigger",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("Рабочее, ежедневно")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
					},
					Body: []w.SheetPart{
						w.Card{
							Items: []w.CardPart{
								w.Rubric{Text: "Password"},
								w.Line{
									Items: []w.LinePart{
										w.Mono{Text: "cbk_…lamp"},
										w.Dim{
											Parts: []w.Inline{
												w.Text("issued 2 August · last run 16 August 03:10 from "),
												w.Mono{Text: "10.0.0.5"},
												w.Text(" · 14 runs"),
											},
											Size: w.Dim11_5,
										},
										w.Line{
											Items: []w.LinePart{
												w.Button{Label: "New password", To: "/app/settings/backup/trigger?job=work&op=new", Submit: true},
												w.Button{
													Label:  "Revoke",
													Kind:   w.ButtonDanger,
													To:     "/app/settings/backup/trigger?job=work&op=revoke",
													Submit: true,
												},
											},
											Gap:   w.LineGapNormal,
											Place: w.Pushed,
										},
									},
									Gap:   w.LineGapNormal,
									Block: true,
								},
							},
							Variant: w.CardFlush,
						},
						w.Card{
							Items: []w.CardPart{
								w.Rubric{Text: "The line to paste", Items: []w.RubricPart{w.Badge{Label: "local", Tone: w.BadgeLocal}}},
								w.Seg{
									Options: []w.SegOption{
										{Label: "crontab", Href: "/app/settings/backup/trigger?job=work&tab=crontab", On: true},
										{Label: "GoSentry", Href: "/app/settings/backup/trigger?job=work&tab=gosentry"},
										{Label: "Gitea", Href: "/app/settings/backup/trigger?job=work&tab=gitea"},
										{Label: "curl", Href: "/app/settings/backup/trigger?job=work&tab=curl"},
									},
									Variant: w.SegStart,
								},
								w.Diag{
									Text:    "0 3 * * * curl -fsS -m 30 -X POST -H 'Authorization: Bearer cbk_••••' 'https://carrel.example/api/backup/run?job=work'",
									Variant: w.DiagInk,
								},
								w.CardText{
									Kind: w.CardUsage,
									Parts: []w.CardTextPart{
										w.Text("The password was shown once and is hidden here. "),
										w.Strong("New password"),
										w.Text(" shows the line whole; the old one stops working."),
									},
								},
							},
							Variant: w.CardFlush,
						},
					},
					Foot: []w.SheetFootPart{
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Close", Href: "/app/settings/backup"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGapNormal,
							Place: w.Pushed,
						},
					},
					Width: w.SheetWidth680,
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Backup13 is frame «Внешний запуск · пароль выдан» (s-backup-13).
func Backup13() Screen {
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
					w.Nav{Items: []w.NavItem{
						{Icon: "back", Label: "Back to app", Href: "#"}}, Sections: true},
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
								On: 6,
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
							Title: &w.HeadTitle{Text: "Backup"},
							Subs: []w.HeadSub{
								w.HeadSub{Parts: []w.HeadSubPart{w.Text("Three jobs · their own sources, destinations and triggers")}},
							},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "New job", Icon: "plus", Kind: w.ButtonPrimary},
									w.Button{Label: "Restore", Icon: "restore"},
								},
							},
						},
						w.Table{
							Cols: []w.TableCol{
								{Width: w.ColTick},
								{Label: "Job"},
								{Label: "What it copies", Rank: w.Rank3, Secondary: true},
								{Label: "Where to", Rank: w.Rank4, Secondary: true},
								{Label: "When"},
								{Label: "Keep", Secondary: true},
								{Label: "Last backup", Main: true},
								{Width: w.ColTime},
							},
							Rows: []w.TableRow{
								{
									Cells: []w.TableCell{
										{Pick: ptr(w.TickedOn)},
										{Parts: []w.CellPart{w.TableFile{Icon: "archive", Name: "Рабочее, ежедневно"}}},
										{
											Parts: []w.CellPart{
												w.Line{
													Items: []w.LinePart{
														w.Mark{Colour: "#3F6E8C"},
														w.Mark{Colour: "#6B4E7A"},
														w.Mark{Colour: "#4A6B52"},
														w.Mono{Text: "3 collections", Variant: w.Mono11Indent},
													},
													Gap: w.LineGapTight,
												},
											},
										},
										{
											Parts: []w.CellPart{w.Text("sftpgo · vps · Documents"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("/backup/work")}}},
											Mono:  true,
										},
										{Parts: []w.CellPart{w.Badge{Label: "trigger"}}},
										{Parts: []w.CellPart{w.Text("14")}, Mono: true},
										{
											Parts: []w.CellPart{w.Text("2026-08-16 03:10"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("18.4 MB · verified")}}},
											Mono:  true,
										},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Run", Small: true, Secondary: true},
												w.Button{Label: "Edit", Small: true, Secondary: true},
											},
											Drop: &w.Drop{
												Trigger: w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
												Row:     true,
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{Label: "Run"},
														w.MenuItem{Label: "Edit"},
														w.MenuItem{Label: "External trigger", Href: "/app/settings/backup/trigger?job=work"},
													},
													Width: w.MenuWidth212,
													Pop:   w.MenuUnderEnd,
												},
											},
											Acts: true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Pick: ptr(w.TickedOn)},
										{Parts: []w.CellPart{w.TableFile{Icon: "archive", Name: "Личное на домашний диск"}}},
										{
											Parts: []w.CellPart{
												w.Line{
													Items: []w.LinePart{
														w.Mark{Colour: "#4A6B52"},
														w.Mark{Colour: "#4A6B52"},
														w.Mark{Colour: "#A05A2C"},
														w.Mark{Colour: "#4A6B52"},
														w.Mono{Text: "4 collections + files", Variant: w.Mono11Indent},
													},
													Gap: w.LineGapTight,
												},
											},
										},
										{
											Parts: []w.CellPart{w.Text("nextcloud · Files"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("/carrel-backup")}}},
											Mono:  true,
										},
										{
											Parts: []w.CellPart{w.Text("on sign-in"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("if older than 7 days")}}},
											Mono:  true,
										},
										{Parts: []w.CellPart{w.Text("8")}, Mono: true},
										{
											Parts: []w.CellPart{w.Text("2026-08-14 09:12"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("2.1 GB · verified")}}},
											Mono:  true,
										},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Run", Small: true, Secondary: true},
												w.Button{Label: "Edit", Small: true, Secondary: true},
											},
											Drop: &w.Drop{
												Trigger: w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
												Row:     true,
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{Label: "Run"},
														w.MenuItem{Label: "Edit"},
														w.MenuItem{Label: "External trigger", Href: "/app/settings/backup/trigger?job=personal"},
													},
													Width: w.MenuWidth212,
													Pop:   w.MenuUnderEnd,
												},
											},
											Acts: true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Pick: ptr(w.TickedOff)},
										{Parts: []w.CellPart{w.TableFile{Icon: "archive", Name: "Contacts archive, by hand"}}},
										{
											Parts: []w.CellPart{
												w.Line{
													Items: []w.LinePart{w.Mark{Colour: "#A05A2C"}, w.Mono{Text: "1 book", Variant: w.Mono11Indent}},
													Gap:   w.LineGapTight,
												},
											},
										},
										{
											Parts: []w.CellPart{w.Text("download in browser"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("nothing is written")}}},
											Mono:  true,
										},
										{Parts: []w.CellPart{w.Text("—")}, Mono: true},
										{Parts: []w.CellPart{w.Text("—")}, Mono: true},
										{Parts: []w.CellPart{w.Text("never")}, Mono: true, Tone: w.CellDusty},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Run", Small: true, Secondary: true},
												w.Button{Label: "Edit", Small: true, Secondary: true},
											},
											Drop: &w.Drop{
												Trigger: w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
												Row:     true,
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{Label: "Run"},
														w.MenuItem{Label: "Edit"},
														w.MenuItem{Label: "External trigger", Href: "/app/settings/backup/trigger?job=contacts"},
													},
													Width: w.MenuWidth212,
													Pop:   w.MenuUnderEnd,
												},
											},
											Acts: true,
										},
									},
									State: w.RowFaded,
								},
							},
							Variant: w.TableWide,
						},
						w.MainHint{
							Text: "verified — the archive was read back from where it was written and the checksum matched; otherwise it is marked red.",
						},
						w.ListGroup{Title: w.Rubric{Text: "History, all jobs", Form: w.RubricInline}, Count: &w.Num{Value: 13}},
						w.Table{
							Cols: []w.TableCol{
								{Label: "When"},
								{Label: "Job"},
								{Label: "Size", Rank: w.Rank3, Secondary: true},
								{Label: "Collections", Rank: w.Rank4},
								{Label: "Verified", Main: true},
								{Label: "Trigger", Secondary: true},
								{Width: w.ColTime},
							},
							Rows: []w.TableRow{
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("2026-08-16 03:10")}, Mono: true},
										{Parts: []w.CellPart{w.Text("Рабочее, ежедневно")}},
										{Parts: []w.CellPart{w.Text("18.4 MB")}, Mono: true},
										{Parts: []w.CellPart{w.Text("3 / 3")}, Mono: true},
										{Parts: []w.CellPart{w.Text("read back")}},
										{
											Parts: []w.CellPart{w.Text("external"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("cbk_…lamp · 10.0.0.5")}}},
											Mono:  true,
										},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Restore", Icon: "restore", Small: true},
												w.Button{Icon: "download", Small: true, Variant: w.ButtonPadX3, Title: "Download"},
											},
											Acts: true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("2026-08-15 21:40")}, Mono: true},
										{Parts: []w.CellPart{w.Text("Рабочее, ежедневно")}},
										{Parts: []w.CellPart{w.Text("—")}, Mono: true},
										{Parts: []w.CellPart{w.Text("—")}, Mono: true},
										{Parts: []w.CellPart{w.Span{Face: w.FaceDanger, Text: "refused: not this job’s password"}}},
										{
											Parts: []w.CellPart{w.Text("external"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("cbk_…91de · 203.0.113.7")}}},
											Mono:  true,
										},
										{Acts: true},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("2026-08-15 03:10")}, Mono: true},
										{Parts: []w.CellPart{w.Text("Рабочее, ежедневно")}},
										{Parts: []w.CellPart{w.Text("18.2 MB")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2 / 3")}, Mono: true, Tone: w.CellDanger},
										{
											Parts: []w.CellPart{
												w.Text("read back · "),
												w.Span{
													Face:  w.FaceDanger,
													Text:  "1 not copied",
													Title: "baikal · Общий: the server answered 401, so this collection is not in the archive. The other two were read back.",
												},
											},
										},
										{
											Parts: []w.CellPart{w.Text("external"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("cbk_…lamp · 10.0.0.5")}}},
											Mono:  true,
										},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Restore", Icon: "restore", Small: true},
												w.Button{Icon: "download", Small: true, Variant: w.ButtonPadX3, Title: "Download"},
											},
											Acts: true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("2026-08-14 09:12")}, Mono: true},
										{Parts: []w.CellPart{w.Text("Личное на домашний диск")}},
										{Parts: []w.CellPart{w.Text("2.1 GB")}, Mono: true},
										{Parts: []w.CellPart{w.Text("4 / 4 + 1 240 files")}, Mono: true},
										{Parts: []w.CellPart{w.Text("read back")}},
										{Parts: []w.CellPart{w.Text("on sign-in")}, Mono: true},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Restore", Icon: "restore", Small: true},
												w.Button{Icon: "download", Small: true, Variant: w.ButtonPadX3, Title: "Download"},
											},
											Acts: true,
										},
									},
								},
							},
						},
					},
				},
				Height: w.ShellHeight560,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/settings/backup/trigger?job=work",
					Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					Title:  "External trigger",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("Рабочее, ежедневно")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
					},
					Body: []w.SheetPart{
						w.Card{
							Items: []w.CardPart{
								w.Rubric{
									Text:  "Password",
									Items: []w.RubricPart{w.Mono{Text: "visible only now", Variant: w.MonoTightDanger}},
								},
								w.Line{
									Items: []w.LinePart{
										w.Input{
											Value:    "cbk_2c9e-07ab-5d31-fern",
											Label:    "Password",
											Mono:     true,
											ReadOnly: true,
											ID:       "trigger-pw",
											Variant:  w.InputKeySmFill,
										},
										w.Button{Label: "Copy", CopyFrom: "trigger-pw", Copied: "Copied"},
									},
									Gap:   w.LineGapNormal,
									Block: true,
								},
							},
							Variant: w.CardFlush,
						},
						w.Card{
							Items: []w.CardPart{
								w.Rubric{Text: "The line to paste", Items: []w.RubricPart{w.Badge{Label: "local", Tone: w.BadgeLocal}}},
								w.Seg{
									Options: []w.SegOption{
										{Label: "crontab", Href: "/app/settings/backup/trigger?job=work&tab=crontab", On: true},
										{Label: "GoSentry", Href: "/app/settings/backup/trigger?job=work&tab=gosentry"},
										{Label: "Gitea", Href: "/app/settings/backup/trigger?job=work&tab=gitea"},
										{Label: "curl", Href: "/app/settings/backup/trigger?job=work&tab=curl"},
									},
									Variant: w.SegStart,
								},
								w.Diag{
									Text:    "0 3 * * * curl -fsS -m 30 -X POST -H 'Authorization: Bearer cbk_2c9e-07ab-5d31-fern' 'https://carrel.example/api/backup/run?job=work'",
									Variant: w.DiagInk,
									ID:      "trigger-line",
								},
								w.CardRow{
									Kind: w.CardTop8,
									Gap:  w.LineGapNormal,
									Items: []w.CardRowPart{
										w.Button{
											Label:    "Copy the line",
											Icon:     "dup",
											Kind:     w.ButtonPrimary,
											CopyFrom: "trigger-line",
											Copied:   "Copied",
										},
									},
								},
								w.CardText{
									Kind: w.CardUsage,
									Parts: []w.CardTextPart{
										w.Text("Every use is logged, naming this password. The old one stopped working when this one was issued."),
									},
								},
							},
							Variant: w.CardFlush,
						},
					},
					Foot: []w.SheetFootPart{
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Close", Href: "/app/settings/backup"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGapNormal,
							Place: w.Pushed,
						},
					},
					Width: w.SheetWidth680,
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Backup14 is frame «Внешний запуск · выключен» (s-backup-14).
func Backup14() Screen {
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
					w.Nav{Items: []w.NavItem{
						{Icon: "back", Label: "Back to app", Href: "#"}}, Sections: true},
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
								On: 6,
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
							Title: &w.HeadTitle{Text: "Backup"},
							Subs: []w.HeadSub{
								w.HeadSub{Parts: []w.HeadSubPart{w.Text("Three jobs · their own sources, destinations and triggers")}},
							},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "New job", Icon: "plus", Kind: w.ButtonPrimary},
									w.Button{Label: "Restore", Icon: "restore"},
								},
							},
						},
						w.Table{
							Cols: []w.TableCol{
								{Width: w.ColTick},
								{Label: "Job"},
								{Label: "What it copies", Rank: w.Rank3, Secondary: true},
								{Label: "Where to", Rank: w.Rank4, Secondary: true},
								{Label: "When"},
								{Label: "Keep", Secondary: true},
								{Label: "Last backup", Main: true},
								{Width: w.ColTime},
							},
							Rows: []w.TableRow{
								{
									Cells: []w.TableCell{
										{Pick: ptr(w.TickedOn)},
										{Parts: []w.CellPart{w.TableFile{Icon: "archive", Name: "Рабочее, ежедневно"}}},
										{
											Parts: []w.CellPart{
												w.Line{
													Items: []w.LinePart{
														w.Mark{Colour: "#3F6E8C"},
														w.Mark{Colour: "#6B4E7A"},
														w.Mark{Colour: "#4A6B52"},
														w.Mono{Text: "3 collections", Variant: w.Mono11Indent},
													},
													Gap: w.LineGapTight,
												},
											},
										},
										{
											Parts: []w.CellPart{w.Text("sftpgo · vps · Documents"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("/backup/work")}}},
											Mono:  true,
										},
										{Parts: []w.CellPart{w.Badge{Label: "trigger"}}},
										{Parts: []w.CellPart{w.Text("14")}, Mono: true},
										{
											Parts: []w.CellPart{w.Text("2026-08-16 03:10"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("18.4 MB · verified")}}},
											Mono:  true,
										},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Run", Small: true, Secondary: true},
												w.Button{Label: "Edit", Small: true, Secondary: true},
											},
											Drop: &w.Drop{
												Trigger: w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
												Row:     true,
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{Label: "Run"},
														w.MenuItem{Label: "Edit"},
														w.MenuItem{Label: "External trigger", Href: "/app/settings/backup/trigger?job=work"},
													},
													Width: w.MenuWidth212,
													Pop:   w.MenuUnderEnd,
												},
											},
											Acts: true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Pick: ptr(w.TickedOn)},
										{Parts: []w.CellPart{w.TableFile{Icon: "archive", Name: "Личное на домашний диск"}}},
										{
											Parts: []w.CellPart{
												w.Line{
													Items: []w.LinePart{
														w.Mark{Colour: "#4A6B52"},
														w.Mark{Colour: "#4A6B52"},
														w.Mark{Colour: "#A05A2C"},
														w.Mark{Colour: "#4A6B52"},
														w.Mono{Text: "4 collections + files", Variant: w.Mono11Indent},
													},
													Gap: w.LineGapTight,
												},
											},
										},
										{
											Parts: []w.CellPart{w.Text("nextcloud · Files"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("/carrel-backup")}}},
											Mono:  true,
										},
										{
											Parts: []w.CellPart{w.Text("on sign-in"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("if older than 7 days")}}},
											Mono:  true,
										},
										{Parts: []w.CellPart{w.Text("8")}, Mono: true},
										{
											Parts: []w.CellPart{w.Text("2026-08-14 09:12"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("2.1 GB · verified")}}},
											Mono:  true,
										},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Run", Small: true, Secondary: true},
												w.Button{Label: "Edit", Small: true, Secondary: true},
											},
											Drop: &w.Drop{
												Trigger: w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
												Row:     true,
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{Label: "Run"},
														w.MenuItem{Label: "Edit"},
														w.MenuItem{Label: "External trigger", Href: "/app/settings/backup/trigger?job=personal"},
													},
													Width: w.MenuWidth212,
													Pop:   w.MenuUnderEnd,
												},
											},
											Acts: true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Pick: ptr(w.TickedOff)},
										{Parts: []w.CellPart{w.TableFile{Icon: "archive", Name: "Contacts archive, by hand"}}},
										{
											Parts: []w.CellPart{
												w.Line{
													Items: []w.LinePart{w.Mark{Colour: "#A05A2C"}, w.Mono{Text: "1 book", Variant: w.Mono11Indent}},
													Gap:   w.LineGapTight,
												},
											},
										},
										{
											Parts: []w.CellPart{w.Text("download in browser"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("nothing is written")}}},
											Mono:  true,
										},
										{Parts: []w.CellPart{w.Text("—")}, Mono: true},
										{Parts: []w.CellPart{w.Text("—")}, Mono: true},
										{Parts: []w.CellPart{w.Text("never")}, Mono: true, Tone: w.CellDusty},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Run", Small: true, Secondary: true},
												w.Button{Label: "Edit", Small: true, Secondary: true},
											},
											Drop: &w.Drop{
												Trigger: w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
												Row:     true,
												Menu: w.Menu{
													Items: []w.MenuPart{
														w.MenuItem{Label: "Run"},
														w.MenuItem{Label: "Edit"},
														w.MenuItem{Label: "External trigger", Href: "/app/settings/backup/trigger?job=contacts"},
													},
													Width: w.MenuWidth212,
													Pop:   w.MenuUnderEnd,
												},
											},
											Acts: true,
										},
									},
									State: w.RowFaded,
								},
							},
							Variant: w.TableWide,
						},
						w.MainHint{
							Text: "verified — the archive was read back from where it was written and the checksum matched; otherwise it is marked red.",
						},
						w.ListGroup{Title: w.Rubric{Text: "History, all jobs", Form: w.RubricInline}, Count: &w.Num{Value: 13}},
						w.Table{
							Cols: []w.TableCol{
								{Label: "When"},
								{Label: "Job"},
								{Label: "Size", Rank: w.Rank3, Secondary: true},
								{Label: "Collections", Rank: w.Rank4},
								{Label: "Verified", Main: true},
								{Label: "Trigger", Secondary: true},
								{Width: w.ColTime},
							},
							Rows: []w.TableRow{
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("2026-08-16 03:10")}, Mono: true},
										{Parts: []w.CellPart{w.Text("Рабочее, ежедневно")}},
										{Parts: []w.CellPart{w.Text("18.4 MB")}, Mono: true},
										{Parts: []w.CellPart{w.Text("3 / 3")}, Mono: true},
										{Parts: []w.CellPart{w.Text("read back")}},
										{
											Parts: []w.CellPart{w.Text("external"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("cbk_…lamp · 10.0.0.5")}}},
											Mono:  true,
										},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Restore", Icon: "restore", Small: true},
												w.Button{Icon: "download", Small: true, Variant: w.ButtonPadX3, Title: "Download"},
											},
											Acts: true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("2026-08-15 21:40")}, Mono: true},
										{Parts: []w.CellPart{w.Text("Рабочее, ежедневно")}},
										{Parts: []w.CellPart{w.Text("—")}, Mono: true},
										{Parts: []w.CellPart{w.Text("—")}, Mono: true},
										{Parts: []w.CellPart{w.Span{Face: w.FaceDanger, Text: "refused: not this job’s password"}}},
										{
											Parts: []w.CellPart{w.Text("external"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("cbk_…91de · 203.0.113.7")}}},
											Mono:  true,
										},
										{Acts: true},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("2026-08-15 03:10")}, Mono: true},
										{Parts: []w.CellPart{w.Text("Рабочее, ежедневно")}},
										{Parts: []w.CellPart{w.Text("18.2 MB")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2 / 3")}, Mono: true, Tone: w.CellDanger},
										{
											Parts: []w.CellPart{
												w.Text("read back · "),
												w.Span{
													Face:  w.FaceDanger,
													Text:  "1 not copied",
													Title: "baikal · Общий: the server answered 401, so this collection is not in the archive. The other two were read back.",
												},
											},
										},
										{
											Parts: []w.CellPart{w.Text("external"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("cbk_…lamp · 10.0.0.5")}}},
											Mono:  true,
										},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Restore", Icon: "restore", Small: true},
												w.Button{Icon: "download", Small: true, Variant: w.ButtonPadX3, Title: "Download"},
											},
											Acts: true,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Parts: []w.CellPart{w.Text("2026-08-14 09:12")}, Mono: true},
										{Parts: []w.CellPart{w.Text("Личное на домашний диск")}},
										{Parts: []w.CellPart{w.Text("2.1 GB")}, Mono: true},
										{Parts: []w.CellPart{w.Text("4 / 4 + 1 240 files")}, Mono: true},
										{Parts: []w.CellPart{w.Text("read back")}},
										{Parts: []w.CellPart{w.Text("on sign-in")}, Mono: true},
										{
											Parts: []w.CellPart{
												w.Button{Label: "Restore", Icon: "restore", Small: true},
												w.Button{Icon: "download", Small: true, Variant: w.ButtonPadX3, Title: "Download"},
											},
											Acts: true,
										},
									},
								},
							},
						},
					},
				},
				Height: w.ShellHeight560,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Action: "/app/settings/backup/trigger?job=personal",
					Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					Title:  "External trigger",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("Личное на домашний диск")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
					},
					Body: []w.SheetPart{
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("Turning it on stores this job’s data key under the server’s key — less protection, for convenience: the job can then read without you signed in."),
							},
							Tone:  w.MsgAlert,
							Space: w.MsgFlush,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{
							Label:  "Turn on",
							Kind:   w.ButtonPrimary,
							To:     "/app/settings/backup/trigger?job=personal&op=on",
							Submit: true,
						},
						w.Kbd{Keys: "Enter"},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel", Href: "/app/settings/backup"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGapNormal,
							Place: w.Pushed,
						},
					},
					Width: w.SheetWidth680,
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Backup6 is frame «Backup на пустом экземпляре» (s-backup-6).
func Backup6() Screen {
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
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 4}, Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "archive", Label: "Backup", Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 6,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.RailStatus{
							Text:   "davloom · running · 3 devices",
							Detail: w.Dim{Parts: []w.Inline{w.Text("Last response "), w.Mono{Text: "2026-08-16 09:41"}}, Size: w.Dim11},
						},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Backup"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("No jobs yet")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New job", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Restore", Icon: "restore"},
							},
						},
					},
					w.Empty{
						Parts: []w.EmptyPart{
							w.Text("Carrel copies raw "),
							w.Mono{Text: ".vcf"},
							w.Text(", "),
							w.Mono{Text: ".ics"},
							w.Text(" and files to storage of your own, with a manifest and a read-back check. Nothing is kept here. "),
							w.Aside{
								Face: w.FaceLarge,
								Parts: []w.Inline{
									w.Text("Already have an archive from another instance? "),
									w.Strong("Restore"),
									w.Text(" takes it from a file — a job is not needed for that."),
								},
							},
						},
					},
				},
			},
		},
	}
}

// Backup7 is frame «Восстановление · шаг 1, какой архив» (s-backup-7).
func Backup7() Screen {
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
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 4}, Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 6,
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
							w.Steps{Items: []string{"1 · Archive", "2 · Match", "3 · Report"}, On: 1},
							w.Head{
								Title: &w.HeadTitle{Text: "Restore"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Nothing is written until you have seen what goes where")}}},
							},
							w.Field{
								Label: "Where the archive is",
								Parts: []w.FieldPart{
									w.Seg{
										Options: []w.SegOption{
											{Label: "From a job", Href: "/app/settings/backup/restore?from=job", On: true},
											{Label: "From a server", Href: "/app/settings/backup/restore?from=server"},
											{Label: "From this computer", Href: "/app/settings/backup/restore?from=file"},
										},
										Variant: w.SegStart,
									},
								},
								OwnLine: true,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "From a job"},
									w.Table{
										Rows: []w.TableRow{
											{
												Cells: []w.TableCell{
													{
														Pick:  ptr(w.TickedOn),
														Send:  w.Send{Name: "archive", Value: "lichnoe-na-domashnij-disk", One: true},
														Width: w.ColTick,
													},
													{Parts: []w.CellPart{w.Text("Личное на домашний диск")}},
													{Parts: []w.CellPart{w.Text("2026-08-14 09:12")}, Mono: true},
													{Parts: []w.CellPart{w.Text("2.1 GB")}, Mono: true, Rank: w.Rank4},
													{Parts: []w.CellPart{w.Text("4 collections")}, Mono: true, Rank: w.Rank3},
												},
												State: w.RowPicked,
											},
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOff), Send: w.Send{Name: "archive", Value: "rabochee-ezhednevno", One: true}},
													{Parts: []w.CellPart{w.Text("Рабочее, ежедневно")}},
													{Parts: []w.CellPart{w.Text("2026-08-16 03:10")}, Mono: true},
													{Parts: []w.CellPart{w.Text("18.4 MB")}, Mono: true, Rank: w.Rank4},
													{Parts: []w.CellPart{w.Text("3 collections")}, Mono: true, Rank: w.Rank3},
												},
											},
										},
									},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Archive password"},
									w.CardFields{
										Items: []w.CardFieldsPart{
											w.Field{
												Parts: []w.FieldPart{
													w.Input{
														Kind:  w.InputPassword,
														Name:  "password",
														Value: "••••••••••••",
														Label: "Archive password",
														Focus: true,
													},
													w.FieldHint{Parts: []w.Inline{w.Text("The one set in the job, not your sign-in password.")}},
												},
												Width: w.FieldLG,
											},
										},
									},
								},
							},
							w.FormFoot{
								Items: []w.FormFootPart{w.Button{Label: "Open the manifest", Kind: w.ButtonPrimary, Submit: true}, w.Button{Label: "Cancel"}},
							},
						},
						Variant: w.DialogStep760,
						Action:  "/app/settings/backup/restore",
					},
				},
			},
		},
	}
}

// Backup15 is frame «Восстановление · с этого компьютера» (s-backup-15).
func Backup15() Screen {
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
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 4}, Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 6,
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
							w.Steps{Items: []string{"1 · Archive", "2 · Match", "3 · Report"}, On: 1},
							w.Head{
								Title: &w.HeadTitle{Text: "Restore"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Nothing is written until you have seen what goes where")}}},
							},
							w.Field{
								Label: "Where the archive is",
								Parts: []w.FieldPart{
									w.Seg{
										Options: []w.SegOption{
											{Label: "From a job", Href: "/app/settings/backup/restore?from=job"},
											{Label: "From a server", Href: "/app/settings/backup/restore?from=server"},
											{Label: "From this computer", Href: "/app/settings/backup/restore?from=file", On: true},
										},
										Variant: w.SegStart,
									},
								},
								OwnLine: true,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "From this computer"},
									w.CardFields{
										Items: []w.CardFieldsPart{
											w.Field{
												Label: "Storage for a large archive",
												Parts: []w.FieldPart{
													w.Pick{
														Name: "store",
														Options: []w.PickOption{
															{Value: "", Label: "None — up to 16 MB", Disables: "folder"},
															{Value: "nextcloud|files", Label: "nextcloud · Files"},
															{Value: "sftpgo|documents", Label: "sftpgo · vps · Documents", On: true},
														},
														Label: "Storage for a large archive",
													},
												},
												Width: w.FieldMD,
											},
											w.Field{
												Label: "Folder in it",
												Parts: []w.FieldPart{
													w.Input{Name: "folder", Value: "/carrel/restore", Label: "Folder in it", Mono: true},
													w.FieldHint{
														Parts: []w.Inline{w.Text("A larger archive is put here, read from there and deleted after the restore.")},
													},
												},
												Width: w.FieldLG,
											},
											w.Field{
												Label: "Archive",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Button{
																Label: "Choose a file",
																File: &w.FileField{
																	Form:   "",
																	Name:   "file",
																	Accept: ".zip",
																	Label:  "Archive",
																	Ranges: "/app/settings/backup/restore?op=ranges",
																},
															},
														},
														Gap: w.LineGapNormal,
													},
												},
												Width: w.FieldFull,
											},
										},
										Wide: true,
										Top:  true,
									},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Archive password"},
									w.CardFields{
										Items: []w.CardFieldsPart{
											w.Field{
												Parts: []w.FieldPart{
													w.Input{
														Kind:  w.InputPassword,
														Name:  "password",
														Value: "••••••••••••",
														Label: "Archive password",
														Focus: true,
													},
													w.FieldHint{Parts: []w.Inline{w.Text("The one set in the job, not your sign-in password.")}},
												},
												Width: w.FieldLG,
											},
										},
									},
								},
							},
							w.FormFoot{
								Items: []w.FormFootPart{w.Button{Label: "Open the manifest", Kind: w.ButtonPrimary, Submit: true}, w.Button{Label: "Cancel"}},
							},
						},
						Variant: w.DialogStep760,
						Action:  "/app/settings/backup/restore?from=file",
						Fields:  []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
				},
				Zone:   "restore",
				Target: "restore",
			},
		},
	}
}

// Backup16 is frame «Восстановление · архив больше 16 MB, папки нет» (s-backup-16).
func Backup16() Screen {
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
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 4}, Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 6,
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
							w.Steps{Items: []string{"1 · Archive", "2 · Match", "3 · Report"}, On: 1},
							w.Head{
								Title: &w.HeadTitle{Text: "Restore"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Nothing is written until you have seen what goes where")}}},
							},
							w.Field{
								Label: "Where the archive is",
								Parts: []w.FieldPart{
									w.Seg{
										Options: []w.SegOption{
											{Label: "From a job", Href: "/app/settings/backup/restore?from=job"},
											{Label: "From a server", Href: "/app/settings/backup/restore?from=server"},
											{Label: "From this computer", Href: "/app/settings/backup/restore?from=file", On: true},
										},
										Variant: w.SegStart,
									},
								},
								OwnLine: true,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "From this computer"},
									w.Msg{
										Parts: []w.MsgPart{
											w.Text("The archive is larger than 16 MB. Choose a folder of your storage for it, or put it there and choose "),
											w.Strong("From a server"),
											w.Text("."),
										},
										Tone:  w.MsgAlert,
										Space: w.MsgFlushBelowWide,
									},
									w.CardFields{
										Items: []w.CardFieldsPart{
											w.Field{
												Label: "Storage for a large archive",
												Parts: []w.FieldPart{
													w.Pick{
														Name: "store",
														Options: []w.PickOption{
															{Value: "", Label: "None — up to 16 MB", On: true, Disables: "folder"},
															{Value: "nextcloud|files", Label: "nextcloud · Files"},
															{Value: "sftpgo|documents", Label: "sftpgo · vps · Documents"},
														},
														Label: "Storage for a large archive",
													},
												},
												Width: w.FieldMD,
											},
											w.Field{
												Label: "Folder in it",
												Parts: []w.FieldPart{
													w.Input{Name: "folder", Value: "/carrel/restore", Label: "Folder in it", Mono: true, Disabled: true},
													w.FieldHint{
														Parts: []w.Inline{w.Text("A larger archive is put here, read from there and deleted after the restore.")},
													},
												},
												Width: w.FieldLG,
											},
											w.Field{
												Label: "Archive",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Button{
																Label: "Choose a file",
																File: &w.FileField{
																	Form:   "",
																	Name:   "file",
																	Accept: ".zip",
																	Label:  "Archive",
																	Ranges: "/app/settings/backup/restore?op=ranges",
																},
															},
														},
														Gap: w.LineGapNormal,
													},
												},
												Width: w.FieldFull,
											},
										},
										Wide: true,
										Top:  true,
									},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Archive password"},
									w.CardFields{
										Items: []w.CardFieldsPart{
											w.Field{
												Parts: []w.FieldPart{
													w.Input{
														Kind:  w.InputPassword,
														Name:  "password",
														Value: "••••••••••••",
														Label: "Archive password",
														Focus: true,
													},
													w.FieldHint{Parts: []w.Inline{w.Text("The one set in the job, not your sign-in password.")}},
												},
												Width: w.FieldLG,
											},
										},
									},
								},
							},
							w.FormFoot{
								Items: []w.FormFootPart{w.Button{Label: "Open the manifest", Kind: w.ButtonPrimary, Submit: true}, w.Button{Label: "Cancel"}},
							},
						},
						Variant: w.DialogStep760,
						Action:  "/app/settings/backup/restore?from=file",
						Fields:  []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
				},
				Zone:   "restore",
				Target: "restore",
			},
		},
	}
}

// Backup17 is frame «Восстановление · с сервера» (s-backup-17).
func Backup17() Screen {
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
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 4}, Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 6,
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
							w.Steps{Items: []string{"1 · Archive", "2 · Match", "3 · Report"}, On: 1},
							w.Head{
								Title: &w.HeadTitle{Text: "Restore"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Nothing is written until you have seen what goes where")}}},
							},
							w.Field{
								Label: "Where the archive is",
								Parts: []w.FieldPart{
									w.Seg{
										Options: []w.SegOption{
											{Label: "From a job", Href: "/app/settings/backup/restore?from=job"},
											{Label: "From a server", Href: "/app/settings/backup/restore?from=server", On: true},
											{Label: "From this computer", Href: "/app/settings/backup/restore?from=file"},
										},
										Variant: w.SegStart,
									},
								},
								OwnLine: true,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "From a server"},
									w.CardFields{
										Items: []w.CardFieldsPart{
											w.Field{
												Label: "Collection",
												Parts: []w.FieldPart{
													w.Pick{
														Name: "store",
														Options: []w.PickOption{
															{Value: "nextcloud|files", Label: "nextcloud · Files"},
															{Value: "sftpgo|documents", Label: "sftpgo · vps · Documents", On: true},
														},
														Label: "Collection",
													},
												},
												Width: w.FieldMD,
											},
											w.Field{
												Label: "Archive in it",
												Parts: []w.FieldPart{
													w.Line{
														Items: []w.LinePart{
															w.Input{
																Name:    "path",
																Value:   "/old-carrel/carrel-personal-2026-05-30.zip",
																Label:   "Archive in it",
																Mono:    true,
																Variant: w.InputFill,
															},
															w.Button{Label: "Browse", To: "/app/settings/backup/restore?from=server&op=browse", Submit: true},
														},
														Gap: w.LineGapNormal,
													},
												},
												Width: w.FieldXL,
											},
										},
										Wide: true,
										Top:  true,
									},
								},
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{Text: "Archive password"},
									w.CardFields{
										Items: []w.CardFieldsPart{
											w.Field{
												Parts: []w.FieldPart{
													w.Input{
														Kind:  w.InputPassword,
														Name:  "password",
														Value: "••••••••••••",
														Label: "Archive password",
														Focus: true,
													},
													w.FieldHint{Parts: []w.Inline{w.Text("The one set in the job, not your sign-in password.")}},
												},
												Width: w.FieldLG,
											},
										},
									},
								},
							},
							w.FormFoot{
								Items: []w.FormFootPart{w.Button{Label: "Open the manifest", Kind: w.ButtonPrimary, Submit: true}, w.Button{Label: "Cancel"}},
							},
						},
						Variant: w.DialogStep760,
						Action:  "/app/settings/backup/restore?from=server",
						Fields:  []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
				},
			},
		},
	}
}

// Backup8 is frame «Восстановление · что и куда ляжет» (s-backup-8).
func Backup8() Screen {
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
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 4}, Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 6,
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
							w.Steps{Items: []string{"1 · Archive", "2 · Match", "3 · Report"}, On: 2},
							w.Head{
								Title: &w.HeadTitle{Text: "Restore from 2026-08-14"},
								Subs: []w.HeadSub{
									w.HeadSub{
										Parts: []w.HeadSubPart{
											w.Text("Job «Личное на домашний диск» · manifest read · 4 collections · 147 records · "),
											w.Mono{Text: "carrel-backup/2026-08-14T09-12.zip"},
										},
									},
								},
								Acts: &w.Acts{
									Items: []w.ActsPart{
										w.Button{Label: "Cancel"},
										w.Button{
											Label:  "Restore 3 collections",
											Icon:   "restore",
											Kind:   w.ButtonPrimary,
											Name:   "op",
											Value:  "restore",
											Submit: true,
										},
									},
								},
							},
							w.Field{
								Label: "2 records of the archive are already here and differ",
								Parts: []w.FieldPart{
									w.Seg{
										Options: []w.SegOption{
											{Label: "Skip them", Href: "/app/settings/backup/restore?existing=skip", On: true},
											{Label: "Replace them", Href: "/app/settings/backup/restore?existing=replace"},
											{Label: "Import as copies", Href: "/app/settings/backup/restore?existing=copy"},
										},
									},
									w.FieldHint{
										Parts:   []w.Inline{w.Text("They stay as they are here; the archive's version is not written.")},
										Variant: w.HintColumn,
									},
								},
								OwnLine: true,
							},
							w.Bar{
								Items: []w.BarPart{
									w.BarSep{},
									w.Button{Label: "Restore into", Icon: "restore", Name: "op", Value: "into", Submit: true},
									w.Drop{
										Trigger: w.Pick{Value: "nextcloud · Личный", Mark: &w.Mark{Colour: "#4A6B52"}},
										Menu: w.Menu{
											Items: []w.MenuPart{
												w.MenuItem{
													Mark:  &w.Mark{Colour: "#6B4E7A"},
													Label: "radicale · vps · Проекты",
													Name:  "into",
													Value: "radicale|proekty",
												},
												w.MenuItem{
													Mark:  &w.Mark{Colour: "#4A6B52"},
													Label: "nextcloud · Личный",
													On:    true,
													Name:  "into",
													Value: "nextcloud|lichnyj",
												},
												w.MenuItem{
													Mark:  &w.Mark{Colour: "#3F6E8C"},
													Label: "nextcloud · Работа",
													Name:  "into",
													Value: "nextcloud|rabota",
												},
												w.MenuSep{},
												w.MenuItem{Label: "New collection", Name: "into", Value: "new"},
												w.MenuItem{Label: "Skip this one", Tone: w.MenuToneDusty, Name: "into", Value: "skip"},
											},
											Width: w.MenuWidth290,
											Pop:   w.MenuUnder,
										},
									},
									w.Button{Label: "Skip", Name: "op", Value: "skip", Submit: true},
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
							w.Table{
								Cols: []w.TableCol{
									{Width: w.ColTick, Tick: ptr(w.TickedPart)},
									{Label: "In the archive"},
									{Label: "Records", Rank: w.Rank3, Secondary: true},
									{Label: "Restore into", Width: w.Col26, Rank: w.Rank4, Main: true},
									{Label: "What will happen"},
								},
								Rows: []w.TableRow{
									{
										Cells: []w.TableCell{
											{Pick: ptr(w.TickedOn), Send: w.Send{Name: "collection", Value: "nextcloud/personal"}},
											{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "nextcloud / personal"}}},
											{Parts: []w.CellPart{w.Text("5")}, Mono: true},
											{
												Parts: []w.CellPart{
													w.Drop{
														Trigger: w.Pick{Value: "nextcloud · Личный", Mark: &w.Mark{Colour: "#4A6B52"}},
														Menu: w.Menu{
															Items: []w.MenuPart{
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#6B4E7A"},
																	Label: "radicale · vps · Проекты",
																	Name:  "into.1",
																	Value: "radicale|proekty",
																},
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#4A6B52"},
																	Label: "nextcloud · Личный",
																	On:    true,
																	Name:  "into.1",
																	Value: "nextcloud|lichnyj",
																},
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#3F6E8C"},
																	Label: "nextcloud · Работа",
																	Name:  "into.1",
																	Value: "nextcloud|rabota",
																},
																w.MenuSep{},
																w.MenuItem{Label: "New collection", Name: "into.1", Value: "new"},
																w.MenuItem{Label: "Skip this one", Tone: w.MenuToneDusty, Name: "into.1", Value: "skip"},
															},
															Width: w.MenuWidth290,
															Pop:   w.MenuUnder,
														},
													},
												},
											},
											{Parts: []w.CellPart{w.Text("matched by path — 4 records missing, will be created · 1 already here")}},
										},
										State: w.RowPicked,
									},
									{
										Cells: []w.TableCell{
											{Pick: ptr(w.TickedOn), Send: w.Send{Name: "collection", Value: "radicale/projects"}},
											{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "radicale / projects"}}},
											{Parts: []w.CellPart{w.Text("1")}, Mono: true},
											{
												Parts: []w.CellPart{
													w.Drop{
														Trigger: w.Pick{Value: "radicale · vps · Проекты", Mark: &w.Mark{Colour: "#6B4E7A"}},
														Menu: w.Menu{
															Items: []w.MenuPart{
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#6B4E7A"},
																	Label: "radicale · vps · Проекты",
																	On:    true,
																	Name:  "into.2",
																	Value: "radicale|proekty",
																},
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#4A6B52"},
																	Label: "nextcloud · Личный",
																	Name:  "into.2",
																	Value: "nextcloud|lichnyj",
																},
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#3F6E8C"},
																	Label: "nextcloud · Работа",
																	Name:  "into.2",
																	Value: "nextcloud|rabota",
																},
																w.MenuSep{},
																w.MenuItem{Label: "New collection", Name: "into.2", Value: "new"},
																w.MenuItem{Label: "Skip this one", Tone: w.MenuToneDusty, Name: "into.2", Value: "skip"},
															},
															Width: w.MenuWidth290,
															Pop:   w.MenuUnder,
														},
													},
												},
											},
											{Parts: []w.CellPart{w.Text("matched by path — no differences")}},
										},
										State: w.RowPicked,
									},
									{
										Cells: []w.TableCell{
											{Pick: ptr(w.TickedOn), Send: w.Send{Name: "collection", Value: "baikal/shared"}},
											{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#9A9280"}, Name: "baikal / shared"}}},
											{Parts: []w.CellPart{w.Text("140")}, Mono: true},
											{
												Parts: []w.CellPart{
													w.Drop{
														Trigger: w.Pick{Value: "nextcloud · Личный", Mark: &w.Mark{Colour: "#4A6B52"}},
														Menu: w.Menu{
															Items: []w.MenuPart{
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#6B4E7A"},
																	Label: "radicale · vps · Проекты",
																	Name:  "into.3",
																	Value: "radicale|proekty",
																},
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#4A6B52"},
																	Label: "nextcloud · Личный",
																	On:    true,
																	Name:  "into.3",
																	Value: "nextcloud|lichnyj",
																},
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#3F6E8C"},
																	Label: "nextcloud · Работа",
																	Name:  "into.3",
																	Value: "nextcloud|rabota",
																},
																w.MenuSep{},
																w.MenuItem{Label: "New collection", Name: "into.3", Value: "new"},
																w.MenuItem{Label: "Skip this one", Tone: w.MenuToneDusty, Name: "into.3", Value: "skip"},
															},
															Width: w.MenuWidth290,
															Pop:   w.MenuUnder,
														},
													},
												},
											},
											{
												Parts: []w.CellPart{
													w.Span{Face: w.FaceDanger, Text: "that server is gone — another collection chosen"},
													w.Text(" · 1 already here"),
												},
											},
										},
										State: w.RowPicked,
									},
									{
										Cells: []w.TableCell{
											{Pick: ptr(w.TickedOff), Send: w.Send{Name: "collection", Value: "nextcloud/birthdays"}},
											{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "nextcloud / birthdays"}}},
											{Parts: []w.CellPart{w.Text("1")}, Mono: true},
											{Parts: []w.CellPart{w.Text("—")}, Mono: true, Unset: true, Tone: w.CellDusty},
											{Parts: []w.CellPart{w.Text("skip")}, Dim: true},
										},
									},
								},
								Variant: w.TableFit,
							},
							w.Msg{
								Parts:  []w.MsgPart{w.Text("Restoring is an import: records made since the archive stay, nothing is deleted.")},
								Tone:   w.MsgAlert,
								Space:  w.MsgFlushSides,
								Column: true,
							},
						},
						Variant: w.DialogWide,
						Action:  "/app/settings/backup/restore/match",
					},
				},
			},
		},
	}
}

// Backup20 is frame «Восстановление с этого компьютера · что и куда ляжет» (s-backup-20).
func Backup20() Screen {
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
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 4}, Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 6,
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
							w.Steps{Items: []string{"1 · Archive", "2 · Match", "3 · Report"}, On: 2},
							w.Head{
								Title: &w.HeadTitle{Text: "Restore from 2026-08-14"},
								Subs: []w.HeadSub{
									w.HeadSub{
										Parts: []w.HeadSubPart{
											w.Text("From this computer · manifest read · 4 collections · 147 records · "),
											w.Mono{Text: "carrel-backup-2026-08-14.zip"},
										},
									},
								},
								Acts: &w.Acts{
									Items: []w.ActsPart{
										w.Button{Label: "Cancel"},
										w.Button{
											Label:  "Restore 3 collections",
											Icon:   "restore",
											Kind:   w.ButtonPrimary,
											Name:   "op",
											Value:  "restore",
											Submit: true,
										},
									},
								},
							},
							w.Field{
								Label: "Records of the archive that are already here and differ",
								Parts: []w.FieldPart{
									w.Seg{
										Options: []w.SegOption{
											{Label: "Skip them", Href: "/app/settings/backup/restore?existing=skip", On: true},
											{Label: "Replace them", Href: "/app/settings/backup/restore?existing=replace"},
											{Label: "Import as copies", Href: "/app/settings/backup/restore?existing=copy"},
										},
									},
									w.FieldHint{
										Parts:   []w.Inline{w.Text("They stay as they are here; the archive's version is not written.")},
										Variant: w.HintColumn,
									},
								},
								OwnLine: true,
							},
							w.Bar{
								Items: []w.BarPart{
									w.BarSep{},
									w.Button{Label: "Restore into", Icon: "restore", Name: "op", Value: "into", Submit: true},
									w.Drop{
										Trigger: w.Pick{Value: "nextcloud · Личный", Mark: &w.Mark{Colour: "#4A6B52"}},
										Menu: w.Menu{
											Items: []w.MenuPart{
												w.MenuItem{
													Mark:  &w.Mark{Colour: "#6B4E7A"},
													Label: "radicale · vps · Проекты",
													Name:  "into",
													Value: "radicale|proekty",
												},
												w.MenuItem{
													Mark:  &w.Mark{Colour: "#4A6B52"},
													Label: "nextcloud · Личный",
													On:    true,
													Name:  "into",
													Value: "nextcloud|lichnyj",
												},
												w.MenuItem{
													Mark:  &w.Mark{Colour: "#3F6E8C"},
													Label: "nextcloud · Работа",
													Name:  "into",
													Value: "nextcloud|rabota",
												},
												w.MenuSep{},
												w.MenuItem{Label: "New collection", Name: "into", Value: "new"},
												w.MenuItem{Label: "Skip this one", Tone: w.MenuToneDusty, Name: "into", Value: "skip"},
											},
											Width: w.MenuWidth290,
											Pop:   w.MenuUnder,
										},
									},
									w.Button{Label: "Skip", Name: "op", Value: "skip", Submit: true},
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
							w.Table{
								Cols: []w.TableCol{
									{Width: w.ColTick, Tick: ptr(w.TickedPart)},
									{Label: "In the archive"},
									{Label: "Records", Rank: w.Rank3, Secondary: true},
									{Label: "Restore into", Width: w.Col26, Rank: w.Rank4, Main: true},
									{Label: "What will happen"},
								},
								Rows: []w.TableRow{
									{
										Cells: []w.TableCell{
											{Pick: ptr(w.TickedOn), Send: w.Send{Name: "collection", Value: "nextcloud/personal"}},
											{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "nextcloud / personal"}}},
											{Parts: []w.CellPart{w.Text("5")}, Mono: true},
											{
												Parts: []w.CellPart{
													w.Drop{
														Trigger: w.Pick{Value: "nextcloud · Личный", Mark: &w.Mark{Colour: "#4A6B52"}},
														Menu: w.Menu{
															Items: []w.MenuPart{
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#6B4E7A"},
																	Label: "radicale · vps · Проекты",
																	Name:  "into.1",
																	Value: "radicale|proekty",
																},
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#4A6B52"},
																	Label: "nextcloud · Личный",
																	On:    true,
																	Name:  "into.1",
																	Value: "nextcloud|lichnyj",
																},
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#3F6E8C"},
																	Label: "nextcloud · Работа",
																	Name:  "into.1",
																	Value: "nextcloud|rabota",
																},
																w.MenuSep{},
																w.MenuItem{Label: "New collection", Name: "into.1", Value: "new"},
																w.MenuItem{Label: "Skip this one", Tone: w.MenuToneDusty, Name: "into.1", Value: "skip"},
															},
															Width: w.MenuWidth290,
															Pop:   w.MenuUnder,
														},
													},
												},
											},
											{Parts: []w.CellPart{w.Text("matched by path — compared when you restore")}},
										},
										State: w.RowPicked,
									},
									{
										Cells: []w.TableCell{
											{Pick: ptr(w.TickedOn), Send: w.Send{Name: "collection", Value: "radicale/projects"}},
											{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "radicale / projects"}}},
											{Parts: []w.CellPart{w.Text("1")}, Mono: true},
											{
												Parts: []w.CellPart{
													w.Drop{
														Trigger: w.Pick{Value: "radicale · vps · Проекты", Mark: &w.Mark{Colour: "#6B4E7A"}},
														Menu: w.Menu{
															Items: []w.MenuPart{
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#6B4E7A"},
																	Label: "radicale · vps · Проекты",
																	On:    true,
																	Name:  "into.2",
																	Value: "radicale|proekty",
																},
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#4A6B52"},
																	Label: "nextcloud · Личный",
																	Name:  "into.2",
																	Value: "nextcloud|lichnyj",
																},
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#3F6E8C"},
																	Label: "nextcloud · Работа",
																	Name:  "into.2",
																	Value: "nextcloud|rabota",
																},
																w.MenuSep{},
																w.MenuItem{Label: "New collection", Name: "into.2", Value: "new"},
																w.MenuItem{Label: "Skip this one", Tone: w.MenuToneDusty, Name: "into.2", Value: "skip"},
															},
															Width: w.MenuWidth290,
															Pop:   w.MenuUnder,
														},
													},
												},
											},
											{Parts: []w.CellPart{w.Text("matched by path — compared when you restore")}},
										},
										State: w.RowPicked,
									},
									{
										Cells: []w.TableCell{
											{Pick: ptr(w.TickedOn), Send: w.Send{Name: "collection", Value: "baikal/shared"}},
											{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#9A9280"}, Name: "baikal / shared"}}},
											{Parts: []w.CellPart{w.Text("140")}, Mono: true},
											{
												Parts: []w.CellPart{
													w.Drop{
														Trigger: w.Pick{Value: "nextcloud · Личный", Mark: &w.Mark{Colour: "#4A6B52"}},
														Menu: w.Menu{
															Items: []w.MenuPart{
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#6B4E7A"},
																	Label: "radicale · vps · Проекты",
																	Name:  "into.3",
																	Value: "radicale|proekty",
																},
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#4A6B52"},
																	Label: "nextcloud · Личный",
																	On:    true,
																	Name:  "into.3",
																	Value: "nextcloud|lichnyj",
																},
																w.MenuItem{
																	Mark:  &w.Mark{Colour: "#3F6E8C"},
																	Label: "nextcloud · Работа",
																	Name:  "into.3",
																	Value: "nextcloud|rabota",
																},
																w.MenuSep{},
																w.MenuItem{Label: "New collection", Name: "into.3", Value: "new"},
																w.MenuItem{Label: "Skip this one", Tone: w.MenuToneDusty, Name: "into.3", Value: "skip"},
															},
															Width: w.MenuWidth290,
															Pop:   w.MenuUnder,
														},
													},
												},
											},
											{Parts: []w.CellPart{w.Text("that server is gone — another collection chosen")}, Tone: w.CellDanger},
										},
										State: w.RowPicked,
									},
									{
										Cells: []w.TableCell{
											{Pick: ptr(w.TickedOff), Send: w.Send{Name: "collection", Value: "nextcloud/birthdays"}},
											{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "nextcloud / birthdays"}}},
											{Parts: []w.CellPart{w.Text("1")}, Mono: true},
											{Parts: []w.CellPart{w.Text("—")}, Mono: true, Unset: true, Tone: w.CellDusty},
											{Parts: []w.CellPart{w.Text("skip")}, Dim: true},
										},
									},
								},
								Variant: w.TableFit,
							},
							w.Msg{
								Parts:  []w.MsgPart{w.Text("Restoring is an import: records made since the archive stay, nothing is deleted.")},
								Tone:   w.MsgAlert,
								Space:  w.MsgFlushSides,
								Column: true,
							},
						},
						Variant:  w.DialogWide,
						Action:   "/app/settings/backup/restore/match",
						Ranges:   "/app/settings/backup/restore?op=ranges",
						RangesOf: "file",
					},
				},
				Zone:   "restore",
				Target: "restore",
			},
		},
	}
}

// Backup18 is frame «Восстановление · шаг 3, что сделано» (s-backup-18).
func Backup18() Screen {
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
								{Icon: "server", Label: "Connections", Count: &w.Num{Value: 4}, Href: "#"},
								{Icon: "file", Label: "Attachments", Href: "#"},
								{Icon: "device", Label: "New device", Href: "#"},
								{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
								{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
								{Icon: "theme", Label: "Appearance", Href: "#"},
							},
							On: 6,
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
							w.Steps{Items: []string{"1 · Archive", "2 · Match", "3 · Report"}, On: 3},
							w.Head{
								Title: &w.HeadTitle{Text: "Restored from 2026-08-14"},
								Subs: []w.HeadSub{
									w.HeadSub{
										Parts: []w.HeadSubPart{w.Text("3 collections · 142 records created · 2 already here, skipped · 1 not restored")},
									},
								},
								Acts: &w.Acts{Items: []w.ActsPart{w.Button{Label: "Done", Kind: w.ButtonPrimary, Href: "/app/settings/backup"}}},
							},
							w.Table{
								Cols: []w.TableCol{
									{Label: "In the archive"},
									{Label: "Restored into", Width: w.Col26, Rank: w.Rank4},
									{Label: "Created", Rank: w.Rank3, Secondary: true},
									{Label: "What happened", Main: true},
								},
								Rows: []w.TableRow{
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "nextcloud / personal"}}},
											{Parts: []w.CellPart{w.Text("nextcloud · Личный")}, Mono: true},
											{Parts: []w.CellPart{w.Text("4")}, Mono: true},
											{Parts: []w.CellPart{w.Text("4 missing records created · 1 already here — skipped")}},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "radicale / projects"}}},
											{Parts: []w.CellPart{w.Text("radicale · vps · Проекты")}, Mono: true},
											{Parts: []w.CellPart{w.Text("0")}, Mono: true},
											{Parts: []w.CellPart{w.Text("no differences")}, Dim: true},
										},
									},
									{
										Cells: []w.TableCell{
											{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#9A9280"}, Name: "baikal / shared"}}},
											{Parts: []w.CellPart{w.Text("nextcloud · Личный")}, Mono: true},
											{Parts: []w.CellPart{w.Text("138")}, Mono: true},
											{
												Parts: []w.CellPart{
													w.Text("1 already here — skipped · "),
													w.Span{Face: w.FaceDanger, Text: "1 not restored: the server refused it (413, too large)"},
												},
											},
										},
									},
								},
								Variant: w.TableFit,
							},
						},
						Variant: w.DialogWide,
					},
				},
			},
		},
	}
}
