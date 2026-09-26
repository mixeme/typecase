// Раздел макета: 7.15 Настройки · резервная копия и восстановление.

package screens

import w "carrel.dev/typecase/widget"

// Backup1 is frame «Резервная копия · задания» (s-backup-1).
func Backup1() Screen {
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
							On: 6,
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
						Title: &w.HeadTitle{Text: "Backup"},
						Subs: []w.HeadSub{
							w.HeadSub{Parts: []w.HeadSubPart{w.Text("Three jobs · their own sources, destinations and triggers")}},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New job", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Restore", Icon: "back"},
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
							{Label: "Last backup"},
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
												Gap: w.LineGap3,
											},
										},
									},
									{
										Parts: []w.CellPart{w.Text("sftpgo · vps · Documents"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("/backup/work")}}},
										Mono:  true,
									},
									{
										Parts: []w.CellPart{w.Text("cron · 03:00"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("external trigger")}}},
										Mono:  true,
									},
									{Parts: []w.CellPart{w.Text("14")}, Mono: true},
									{
										Parts: []w.CellPart{w.Text("2026-08-16 03:10"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("18.4 MB · verified")}}},
										Mono:  true,
									},
									{
										Parts: []w.CellPart{
											w.Button{Label: "Run", Small: true},
											w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
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
													w.Mark{Colour: "#A05A2C"},
													w.Mono{Text: "2 collections + files", Variant: w.Mono11Indent},
												},
												Gap: w.LineGap3,
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
											w.Button{Label: "Run", Small: true},
											w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
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
												Gap:   w.LineGap3,
											},
										},
									},
									{
										Parts: []w.CellPart{w.Text("download in browser"), w.Break{}, w.Dim{Parts: []w.Inline{w.Text("nothing is written")}}},
										Mono:  true,
									},
									{Parts: []w.CellPart{w.Text("on demand")}, Mono: true},
									{Parts: []w.CellPart{w.Text("—")}, Mono: true},
									{Parts: []w.CellPart{w.Text("never")}, Mono: true, Tone: w.CellDusty},
									{
										Parts: []w.CellPart{
											w.Button{Label: "Run", Small: true},
											w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"},
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
						Text: "verified — the archive was read back from where it was written and the checksum matched. Written and not read back is not verified, and is marked red.",
					},
					w.ListGroup{Title: w.Rubric{Text: "History, all jobs", Form: w.RubricInline}, Count: &w.Num{Value: 12}},
					w.Table{
						Cols: []w.TableCol{
							{Label: "When"},
							{Label: "Job"},
							{Label: "Size", Rank: w.Rank3, Secondary: true},
							{Label: "Collections", Rank: w.Rank4},
							{Label: "Verified"},
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
									{Parts: []w.CellPart{w.Text("cron")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Label: "Restore", Small: true},
											w.Button{Icon: "down", Small: true, Variant: w.ButtonPadX3, Title: "Download"},
										},
										Acts: true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.Text("2026-08-15 03:10")}, Mono: true},
									{Parts: []w.CellPart{w.Text("Рабочее, ежедневно")}},
									{Parts: []w.CellPart{w.Text("18.2 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2 / 3")}, Mono: true, Tone: w.CellDanger},
									{Parts: []w.CellPart{w.Text("baikal · Общий: 401")}, Tone: w.CellDanger},
									{Parts: []w.CellPart{w.Text("cron")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Label: "Restore", Small: true},
											w.Button{Icon: "down", Small: true, Variant: w.ButtonPadX3, Title: "Download"},
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
									{Parts: []w.CellPart{w.Text("2 / 2 + 1 240 files")}, Mono: true},
									{Parts: []w.CellPart{w.Text("read back")}},
									{Parts: []w.CellPart{w.Text("on sign-in")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Label: "Restore", Small: true},
											w.Button{Icon: "down", Small: true, Variant: w.ButtonPadX3, Title: "Download"},
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
							On: 6,
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
						Title: &w.HeadTitle{Text: "Личное на домашний диск"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Job · last copy 2 days ago")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Cancel"},
								w.Button{Label: "Run now"},
								w.Button{Label: "Save", Kind: w.ButtonPrimary},
							},
						},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label: "Name",
								Parts: []w.FieldPart{w.Input{Value: "Личное на домашний диск", Label: "Name", Focus: true}},
								Width: w.FieldLG,
							},
							w.Field{Parts: []w.FieldPart{w.Check{On: true, Parts: []w.CheckPart{w.Text("Job is on")}}}},
							w.FieldSet{
								Rubric: w.Rubric{
									Text:  "What to copy",
									Items: []w.RubricPart{w.Plain{Text: "2 collections and 1 folder chosen · 2.1 GB last time"}},
								},
								Items: []w.FieldSetPart{
									w.FieldGrid{
										Items: []w.FieldGridPart{
											w.TickRow{
												Ticked: w.TickedPart,
												Form:   w.TickRowGroup,
												Mark:   w.Mark{Kind: w.MarkMulti},
												Label:  "nextcloud",
												Meta:   "4 of 6",
											},
											w.TickRow{
												Ticked: w.TickedOff,
												Form:   w.TickRowGroup,
												Mark:   w.Mark{Kind: w.MarkMulti},
												Label:  "radicale · vps",
												Meta:   "0 of 2",
											},
											w.TickRow{Form: w.TickRowNested, Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "cal."},
											w.TickRow{
												Ticked: w.TickedOff,
												Form:   w.TickRowNested,
												Mark:   w.Mark{Colour: "#6B4E7A"},
												Label:  "Проекты",
												Meta:   "cal.",
											},
											w.TickRow{
												Ticked: w.TickedOff,
												Form:   w.TickRowNested,
												Mark:   w.Mark{Colour: "#3F6E8C"},
												Label:  "Работа",
												Meta:   "cal.",
											},
											w.TickRow{
												Ticked: w.TickedOff,
												Form:   w.TickRowNested,
												Mark:   w.Mark{Colour: "#8C6239"},
												Label:  "Дежурства",
												Meta:   "cal.",
											},
											w.TickRow{Form: w.TickRowNested, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "book"},
											w.TickRow{
												Ticked: w.TickedPart,
												Form:   w.TickRowGroup,
												Mark:   w.Mark{Kind: w.MarkMulti},
												Label:  "sftpgo · vps",
												Meta:   "1 of 2",
											},
											w.TickRow{Form: w.TickRowNested, Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "notes"},
											w.TickRow{Form: w.TickRowNested, Mark: w.Mark{Colour: "#4A6B52"}, Label: "Documents", Meta: "files · 2.0 GB"},
											w.TickRow{
												Ticked: w.TickedOff,
												Form:   w.TickRowNested,
												Mark:   w.Mark{Colour: "#9A9280"},
												Label:  "Public",
												Meta:   "files · ro",
											},
										},
									},
									w.Msg{
										Parts: []w.MsgPart{
											w.Text("A file collection is copied as files. Tasks and notes live inside calendar collections and are copied with them — they cannot be picked separately. The size beside each collection is what it took last time."),
										},
										Space: w.MsgFlushTop4,
										Place: w.Grown,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Where to"},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Seg{
												Options: []w.SegOption{{Label: "To WebDAV", On: true}, {Label: "Download"}, {Label: "Both"}},
												Variant: w.SegStart,
											},
										},
										OwnLine: true,
										Place:   w.Grown,
									},
									w.Field{Label: "Collection", Parts: []w.FieldPart{w.Pick{Value: "nextcloud · Files"}}, Width: w.FieldLG},
									w.Field{
										Label: "Folder",
										Parts: []w.FieldPart{w.Input{Value: "carrel-backup", Label: "Folder", Mono: true}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Keep copies",
										Parts: []w.FieldPart{w.Input{Value: "8", Label: "Keep copies", Mono: true}},
										Width: w.FieldXS,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Copies past the eighth are deleted"),
													w.Strong("after"),
													w.Text("the new one has been read back and checked, not before. The file is dated and never overwritten."),
												},
											},
										},
										Place: w.Grown,
									},
									w.Msg{
										Parts: []w.MsgPart{
											w.Strong("Download"),
											w.Text(" streams the archive out and writes nothing to the server: no folder, no retention — the copy is yours to keep. It can still run at three in the morning, but only on an external trigger, because a stream needs somebody at the other end: a browser is not there at that hour, a machine with a disk is. "),
											w.Strong("Both"),
											w.Text(" does the two at once — and then the trigger has "),
											w.Strong("two steps"),
											w.Text(", because the file still has to be taken."),
										},
										Space: w.MsgFlushTop4,
										Place: w.Grown,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "How the archive is closed"},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Check{On: true, Parts: []w.CheckPart{w.Text("Encrypt the archive")}},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("On by default. Turning it off makes sense in one case only: the copy goes where nobody but you gets in anyway — and then any archiver opens it."),
												},
											},
										},
										OwnLine: true,
										Place:   w.Grown,
									},
									w.Field{
										Label: "Archive password",
										Parts: []w.FieldPart{w.Input{Value: "••••••••••••", Label: "Archive password"}},
										Width: w.FieldLG,
									},
									w.Msg{
										Parts: []w.MsgPart{
											w.Text("Separate from the sign-in password: the destination may be somebody else’s hosting. Forget it and this job’s archives open for nobody, including you."),
										},
										Tone:  w.MsgAlert,
										Space: w.MsgFlushTop2,
										Place: w.Grown,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "When"},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Check{On: true, Parts: []w.CheckPart{w.Text("On demand — the button in the job list")}},
											w.Line{
												Items: []w.LinePart{
													w.Check{On: true, Parts: []w.CheckPart{w.Text("On sign-in, if the last copy is older than")}},
													w.Input{Value: "7", Label: "Days", Mono: true, Variant: w.InputChip},
													w.Faced{Face: w.FaceLarge, Items: []w.LinePart{w.Text("days")}},
												},
												Gap:   w.LineGap7,
												Block: true,
											},
											w.Check{
												On: true,
												Parts: []w.CheckPart{
													w.Text("On a clock of your own —"),
													w.Mono{Text: "cron"},
													w.Text(", GoSentry, a job in Gitea"),
													w.Badge{
														Label: "local",
														Tone:  w.BadgeLocal,
														Title: "Local instance: the address in the line answers only from inside your own network",
													},
												},
											},
										},
										Gap:   8,
										Place: w.Grown,
									},
									w.Msg{
										Parts: []w.MsgPart{
											w.Text("No time picker here, and that is deliberate: Carrel keeps no clock of its own, so whoever holds the schedule holds the password. What it hands you instead is the line to paste where the clock already is."),
										},
										Space: w.MsgFlushTop4,
										Place: w.Grown,
									},
									w.Card{
										Items: []w.CardPart{
											w.Rubric{
												Text:  "The line to paste",
												Items: []w.RubricPart{w.Mono{Text: "visible only now", Variant: w.MonoTightDanger}},
											},
											w.Seg{
												Options: []w.SegOption{{Label: "crontab", On: true}, {Label: "GoSentry"}, {Label: "Gitea"}, {Label: "curl"}},
												Variant: w.SegStart,
											},
											w.Diag{
												Text:    "0 3 * * * curl -fsS -m 30 -X POST -H 'Authorization: Bearer cbk_7f3a-91de-4402-lamp' 'https://carrel.example/api/backup/run?job=personal'",
												Variant: w.DiagInk,
											},
											w.CardRow{
												Kind: w.CardTop8,
												Gap:  8,
												Items: []w.CardRowPart{
													w.Button{Label: "Copy the line", Icon: "dup", Kind: w.ButtonPrimary},
													w.Button{Label: "Regenerate"},
													w.Dim{
														Parts:   []w.Inline{w.Text("Runs at 03:00 · the time sits inside the line and is changed there")},
														Size:    w.Dim11_5,
														Variant: w.DimWrap,
													},
												},
											},
											w.CardText{
												Kind: w.CardUsage,
												Parts: []w.CardTextPart{
													w.Text("Paste into "),
													w.Mono{Text: "crontab -e"},
													w.Text(" on the machine that is to hold the password. It is one line and stays one line — "),
													w.Mono{Text: "cron"},
													w.Text(" understands no continuations; the wrap you see is this screen’s doing. "),
													w.Mono{Text: "-fsS"},
													w.Text(" keeps it quiet while all is well and makes cron mail you when the request itself fails."),
												},
											},
											w.Msg{
												Parts: []w.MsgPart{
													w.Text("This line "),
													w.Strong("starts"),
													w.Text(" a copy; it does not wait for one. The answer comes back at once and says the job was accepted, nothing more. Whether the copy went through is on this page and in the history — and can be watched from outside, which is what the "),
													w.Strong("GoSentry"),
													w.Text(" tab is for."),
													w.Break{},
													w.Break{},
													w.Text("Starting it is all there is to do "),
													w.Strong("because this job writes to WebDAV"),
													w.Text(". Put "),
													w.Strong("Download"),
													w.Text(" in the destination and what is handed over here becomes a short script instead: the archive is a stream, and somebody has to be at the other end to take it."),
												},
												Space: w.MsgFlushTop9,
											},
											w.CardRow{
												Kind: w.CardFoot,
												Gap:  7,
												Items: []w.CardRowPart{
													w.Turn{},
													w.Span{Face: w.FaceLarge, Text: "The password on its own, and the other endpoints"},
													w.Dim{
														Parts:   []w.Inline{w.Text("— for a request written by hand; the line above already carries it")},
														Size:    w.Dim11_5,
														Variant: w.DimWrap,
													},
												},
											},
										},
										Variant: w.CardChosenFlushGrow,
										Place:   w.Grown,
									},
								},
							},
						},
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

// Backup3 is frame «Резервная копия · задание, что и куда» (s-backup-3).
func Backup3() Screen {
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
							On: 6,
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
						Title: &w.HeadTitle{Text: "Личное на домашний диск"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{w.Text("Job · destination "), w.Strong("WebDAV and download"), w.Text(" · last copy 2 days ago")},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Cancel"},
								w.Button{Label: "Run now"},
								w.Button{Label: "Save", Kind: w.ButtonPrimary},
							},
						},
					},
					w.Form{
						Items: []w.FormPart{
							w.FieldSet{
								Rubric: w.Rubric{Text: "Where to"},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Seg{
												Options: []w.SegOption{{Label: "To WebDAV"}, {Label: "Download"}, {Label: "Both", On: true}},
												Variant: w.SegStart,
											},
										},
										OwnLine: true,
										Place:   w.Grown,
									},
									w.Field{Label: "Collection", Parts: []w.FieldPart{w.Pick{Value: "nextcloud · Files"}}, Width: w.FieldLG},
									w.Field{
										Label: "Folder",
										Parts: []w.FieldPart{w.Input{Value: "carrel-backup", Label: "Folder", Mono: true}},
										Width: w.FieldMD,
									},
									w.Msg{
										Parts: []w.MsgPart{
											w.Strong("Both"),
											w.Text(" writes to WebDAV and still has to be taken — so the trigger below is a script, not a line. This is the same job editor as above; only this control changed."),
										},
										Space: w.MsgFlushTop4,
										Place: w.Grown,
									},
								},
								Variant: w.FieldSetLead,
							},
							w.Card{
								Items: []w.CardPart{
									w.Rubric{
										Text:  "The script to paste",
										Items: []w.RubricPart{w.Mono{Text: "visible only now", Variant: w.MonoTightDanger}},
									},
									w.CardText{
										Kind: w.CardAbove,
										Parts: []w.CardTextPart{
											w.Text("Destination: "),
											w.Key("WebDAV and download"),
											w.Text(" — two steps, so this is a file, not a line"),
										},
									},
									w.Seg{
										Options: []w.SegOption{{Label: "crontab", On: true}, {Label: "GoSentry"}, {Label: "Gitea"}, {Label: "curl"}},
										Variant: w.SegStart,
									},
									w.Diag{
										Text:    "#!/bin/sh\nset -eu\nh='Authorization: Bearer cbk_7f3a-91de-4402-lamp'\nb=https://carrel.example/api/backup\nout=/var/backups\n\nid=$(curl -fsS -X POST -H \"$h\" \"$b/run?job=personal\")\nuntil [ \"$(curl -fsS -H \"$h\" \"$b/status?id=$id\")\" != running ]; do sleep 20; done\ncurl -fsS -H \"$h\" \"$b/$id/download\" -o \"$out/carrel-personal-$(date +%F).zip\"",
										Variant: w.DiagInk,
									},
									w.CardRow{
										Kind: w.CardTop8,
										Gap:  8,
										Items: []w.CardRowPart{
											w.Button{Label: "Copy the script", Icon: "dup", Kind: w.ButtonPrimary},
											w.Button{Label: "Save as", Icon: "down", Aside: &w.Mono{Text: "carrel-personal.sh", Variant: w.Mono11}},
											w.Button{Label: "Regenerate"},
										},
									},
									w.CardText{
										Kind: w.CardUsage,
										Parts: []w.CardTextPart{
											w.Text("Put it where the password is to live, make it executable, and give "),
											w.Mono{Text: "crontab -e"},
											w.Text(" one line: "),
											w.Mono{Text: "0 3 * * * /usr/local/bin/carrel-personal.sh"},
											w.Text(" — the schedule stays a line, the work moves into the file. The date has to live in the script and not in the crontab line: "),
											w.Mono{Text: "cron"},
											w.Text(" reads "),
											w.Mono{Text: "%"},
											w.Text(" as a newline, so "),
											w.Mono{Text: "date +%F"},
											w.Text(" written up there would break the entry."),
										},
									},
									w.Msg{
										Parts: []w.MsgPart{
											w.Text("A copy that fails needs no separate check: the download answers with an error, "),
											w.Mono{Text: "set -eu"},
											w.Text(" stops the script on it, and cron sends the mail. The wait is open-ended on purpose — a copy takes what it takes, and "),
											w.Mono{Text: "409"},
											w.Text(" until then means «not yet», not «broken»."),
										},
										Space: w.MsgFlushTop9,
									},
								},
								Variant: w.CardChosenFlush,
								Place:   w.Grown,
							},
						},
					},
				},
			},
		},
	}
}

// Backup6 is frame «Backup на пустом экземпляре» (s-backup-6).
func Backup6() Screen {
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
							On: 6,
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
						Title: &w.HeadTitle{Text: "Backup"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("No jobs yet")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New job", Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Restore", Icon: "back"},
							},
						},
					},
					w.Empty{
						Parts: []w.EmptyPart{
							w.Text("Carrel copies raw"),
							w.Mono{Text: ".ics"},
							w.Text(", "),
							w.Mono{Text: ".vcf"},
							w.Text(" and files to storage of your own, with a manifest and a read-back check. Nothing is kept here. "),
							w.Aside{
								Face: w.FaceLarge,
								Parts: []w.Inline{
									w.Text("Already have an archive from another instance?"),
									w.Strong("Restore"),
									w.Text("takes it from a file — a job is not needed for that."),
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
							On: 6,
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
										Options: []w.SegOption{{Label: "From a job", On: true}, {Label: "From a server"}, {Label: "From this computer"}},
										Variant: w.SegStart,
									},
									w.FieldHint{Parts: []w.Inline{w.Text("One archive, one of three. The other two stay empty until you switch.")}},
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
													{Pick: ptr(w.TickedOn), Width: w.ColTick},
													{Parts: []w.CellPart{w.Text("Личное на домашний диск")}},
													{Parts: []w.CellPart{w.Text("2026-08-14 09:12")}, Mono: true, Rank: w.Rank3},
													{Parts: []w.CellPart{w.Text("2.1 GB")}, Mono: true, Rank: w.Rank4},
													{Parts: []w.CellPart{w.Text("4 collections")}, Mono: true},
												},
												State: w.RowPicked,
											},
											{
												Cells: []w.TableCell{
													{Pick: ptr(w.TickedOff)},
													{Parts: []w.CellPart{w.Text("Рабочее, ежедневно")}},
													{Parts: []w.CellPart{w.Text("2026-08-16 03:10")}, Mono: true, Rank: w.Rank3},
													{Parts: []w.CellPart{w.Text("18.4 MB")}, Mono: true, Rank: w.Rank4},
													{Parts: []w.CellPart{w.Text("3 collections")}, Mono: true},
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
											w.Field{Parts: []w.FieldPart{w.Input{Value: "••••••••••••", Focus: true}}, Width: w.FieldLG},
											w.Span{
												Face: w.FaceMuted12,
												Text: "The one set in the job, not your sign-in password. Without it the manifest cannot be read — Carrel will not even see the list of collections.",
											},
										},
									},
								},
							},
							w.FormFoot{
								Items: []w.FormFootPart{w.Button{Label: "Open the manifest", Kind: w.ButtonPrimary}, w.Button{Label: "Cancel"}},
							},
						},
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
							On: 6,
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
					w.Dialog{
						Items: []w.DialogPart{
							w.Steps{Items: []string{"1 · Archive", "2 · Match", "3 · Report"}, On: 2},
							w.Head{
								Title: &w.HeadTitle{Text: "Restore from 2026-08-14"},
								Subs: []w.HeadSub{
									w.HeadSub{
										Parts: []w.HeadSubPart{
											w.Text("Job «Личное на домашний диск» · manifest read · 4 collections · 1 214 records · "),
											w.Mono{Text: "carrel-backup/2026-08-14T09-12.zip"},
										},
									},
								},
								Acts: &w.Acts{
									Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Restore 3 collections", Kind: w.ButtonPrimary}},
								},
							},
							w.Bar{
								Items:   []w.BarPart{w.BarSep{}, w.Button{Label: "Restore into", Icon: "back"}, w.Button{Label: "Skip"}},
								Count:   &w.BarCount{Text: "3 selected"},
								Right:   &w.BarRight{Items: []w.BarPart{w.Button{Label: "Select all", Small: true}}},
								Variant: w.BarSelection,
							},
							w.Table{
								Cols: []w.TableCol{
									{Width: w.ColTick, Tick: ptr(w.TickedPart)},
									{Label: "In the archive"},
									{Label: "Records", Rank: w.Rank3, Secondary: true},
									{Label: "Restore into", Width: w.Col26, Rank: w.Rank4},
									{Label: "What will happen"},
								},
								Rows: []w.TableRow{
									{
										Cells: []w.TableCell{
											{Pick: ptr(w.TickedOn)},
											{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#4A6B52"}, Name: "nextcloud / personal"}}},
											{Parts: []w.CellPart{w.Text("312")}, Mono: true},
											{Parts: []w.CellPart{w.Text("nextcloud · Личный")}, Mono: true},
											{Parts: []w.CellPart{w.Text("matched by path — 4 records missing, will be created")}},
										},
										State: w.RowPicked,
									},
									{
										Cells: []w.TableCell{
											{Pick: ptr(w.TickedOn)},
											{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#6B4E7A"}, Name: "radicale / projects"}}},
											{Parts: []w.CellPart{w.Text("86")}, Mono: true},
											{
												Parts: []w.CellPart{
													w.Drop{
														Open:    true,
														Trigger: w.Pick{Value: "radicale · vps · Проекты", Focus: true},
														Menu: w.Menu{
															Items: []w.MenuPart{
																w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "radicale · vps · Проекты", On: true},
																w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Личный"},
																w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Label: "nextcloud · Работа"},
																w.MenuSep{},
																w.MenuItem{Icon: "plus", Label: "New collection"},
																w.MenuItem{Label: "Skip this one", Tone: w.MenuToneDusty},
															},
															Pop: w.MenuUnder,
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
											{Pick: ptr(w.TickedOn)},
											{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#9A9280"}, Name: "baikal / shared"}}},
											{Parts: []w.CellPart{w.Text("140")}, Mono: true},
											{Parts: []w.CellPart{w.Text("nextcloud · Личный "), w.Chevron{}}, Mono: true},
											{Parts: []w.CellPart{w.Text("that server is gone — another collection chosen")}, Tone: w.CellDanger},
										},
										State: w.RowPicked,
									},
									{
										Cells: []w.TableCell{
											{Pick: ptr(w.TickedOff)},
											{Parts: []w.CellPart{w.TableFile{Mark: &w.Mark{Colour: "#A05A2C"}, Name: "nextcloud / birthdays"}}},
											{Parts: []w.CellPart{w.Text("51")}, Mono: true},
											{Parts: []w.CellPart{w.Text("—")}, Mono: true, Unset: true, Tone: w.CellDusty},
											{Parts: []w.CellPart{w.Text("skip")}, Dim: true},
										},
									},
								},
								Variant: w.TableFit,
							},
							w.Msg{
								Parts: []w.MsgPart{
									w.Text("Restoring "),
									w.Strong("creates"),
									w.Text(" records, the way import does: a colliding UID gets a fresh one and a line in the report. Nothing on the server is deleted or overwritten — Carrel does not promise “exactly as it was”, because that would mean erasing somebody else’s changes."),
								},
								Tone: w.MsgAlert,
							},
						},
						Variant: w.DialogWide,
					},
				},
			},
		},
	}
}
