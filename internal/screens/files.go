// Раздел макета: 7.5 Файлы.

package screens

import w "carrel.dev/typecase/widget"

// Files1 is frame «Файлы · обзор папки, свойства выбранного» (s-files-1).
func Files1() Screen {
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
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 10}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 5,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked: w.TickedPart,
							Kind:   w.TickRowRoot,
							Icon:   "file",
							Label:  "Selected collections",
							Meta:   "2 of 3",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "sftpgo · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Documents", Meta: "rw", Current: true},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#9A9280"},
							Label:  "Public",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Files", Meta: "rw"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Shortcuts"},
						w.Tree{
							Rows: []w.TreeRow{
								{Icon: "file", Label: "Attachments", Trail: w.TreeNote{Text: "Documents / carrel"}},
								{Icon: "link", Label: "Published", Trail: w.TreeNote{Text: "2"}},
							},
							On:      1,
							Variant: w.TreePadded,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New folder"},
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
						Title: &w.HeadTitle{Text: "2026-08"},
						Crumbs: &w.Crumbs{
							Path:    []w.Crumb{{Label: "Documents", Href: "#"}, {Label: "carrel", Href: "#"}, {Label: "attachments", Href: "#"}},
							Current: "2026-08",
						},
						Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("sftpgo · vps · 4 items")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Upload", Menu: true, Icon: "upload", Kind: w.ButtonPrimary},
								w.Button{Label: "New folder", Icon: "plus"},
								w.Button{Icon: "refresh", Title: "Refresh"},
								w.Button{Icon: "panel", Variant: w.ButtonOnCard, Title: "Details"},
								w.Button{Icon: "dots", Title: "More"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Mono{Text: "412 KB", Variant: w.MonoMuted},
							w.BarSep{},
							w.Button{
								Folds: true,
								Label: "Download",
								Icon:  "download",
								Title: "Download",
								Form:  "files-sel",
								Name:  "op",
								Value: "download",
							},
							w.Button{
								Folds: true,
								Label: "Rename",
								Icon:  "pencil",
								Title: "Rename",
								Form:  "files-sel",
								Name:  "op",
								Value: "rename",
							},
							w.Button{
								Folds: true,
								Label: "Move",
								Icon:  "move",
								Title: "Move",
								Form:  "files-sel",
								Name:  "op",
								Value: "move",
							},
							w.Button{Folds: true, Label: "Copy", Icon: "dup", Title: "Copy", Form: "files-sel", Name: "op", Value: "copy"},
							w.Button{
								Folds: true,
								Label: "Delete",
								Icon:  "trash",
								Kind:  w.ButtonDangerSolid,
								Title: "Delete",
								Form:  "files-sel",
								Name:  "op",
								Value: "delete",
							},
						},
						Count: &w.BarCount{Text: "1 selected"},
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
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter this folder", Short: "Filter", Label: "Filter this folder"},
								Waiting: true,
								Variant: w.FilterH26W190,
							},
							w.Seg{
								Options: []w.SegOption{
									{Label: "Name", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08?sort=name", On: true},
									{Label: "Size", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08?sort=size"},
									{Label: "Changed", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08?sort=changed"},
								},
								Secondary: true,
							},
							w.Seg{
								Options: []w.SegOption{
									{
										Icon:  "menu",
										Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-08?view=list",
										On:    true,
										Title: "List",
									},
									{
										Icon:  "grid",
										Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-08?view=tiles",
										Title: "Tiles",
									},
								},
							},
							w.Button{Label: "Columns", Icon: "cols"},
							w.Drop{
								Trigger: w.Button{Icon: "dots", Title: "More"},
								More:    true,
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuItem{Label: "Name", On: true, Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08?sort=name"},
										w.MenuItem{Label: "Size", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08?sort=size"},
										w.MenuItem{Label: "Changed", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08?sort=changed"},
									},
									Pop:    w.MenuUnderEnd,
									Inline: true,
								},
							},
						},
						Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "4 items · 4.4 MB"}}, Secondary: true},
					},
					w.Table{
						Cols: []w.TableCol{
							{Width: w.ColTick, Tick: ptr(w.TickedPart)},
							{Label: "Name"},
							{Label: "Size", Width: w.ColSize, Rank: w.Rank3},
							{Label: "Type", Width: w.ColType, Rank: w.Rank4, Secondary: true},
							{Label: "Changed", Width: w.ColTime, Main: true},
							{Width: w.ColNum},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{},
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon:   "back",
												Name:   "up one folder",
												Accent: true,
												Href:   "/app/files/sftpgo/documents/carrel/attachments",
											},
										},
										Span: 5,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{
										Pick: ptr(w.TickedOff),
										Send: w.Send{Of: "files-sel", Name: "item", Value: "2026-08-04-whiteboard.jpg"},
									},
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon: "img",
												Name: "2026-08-04-whiteboard.jpg",
												Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08/2026-08-04-whiteboard.jpg",
											},
										},
									},
									{Parts: []w.CellPart{w.Text("2.8 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/jpeg")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-04 17:03")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Icon: "download", Small: true, Title: "Download"},
											w.Button{Icon: "dots", Small: true, Title: "More"},
										},
										Acts: true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{
										Pick: ptr(w.TickedOn),
										Send: w.Send{Of: "files-sel", Name: "item", Value: "2026-08-10-sync-board.png"},
									},
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon: "img",
												Name: "2026-08-10-sync-board.png",
												Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08/2026-08-10-sync-board.png",
											},
										},
									},
									{Parts: []w.CellPart{w.Text("412 KB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/png")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-10 11:07")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Icon: "download", Small: true, Title: "Download"},
											w.Button{Icon: "dots", Small: true, Title: "More"},
										},
										Acts: true,
									},
								},
								State: w.RowPicked,
							},
							{
								Cells: []w.TableCell{
									{Pick: ptr(w.TickedOff), Send: w.Send{Of: "files-sel", Name: "item", Value: "release-checklist.md"}},
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon: "doc",
												Name: "release-checklist.md",
												Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08/release-checklist.md",
											},
										},
									},
									{Parts: []w.CellPart{w.Text("3 KB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("text/markdown")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-09 18:44")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Icon: "download", Small: true, Title: "Download"},
											w.Button{Icon: "dots", Small: true, Title: "More"},
										},
										Acts: true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Pick: ptr(w.TickedOff), Send: w.Send{Of: "files-sel", Name: "item", Value: "смета-подрядчик.pdf"}},
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon: "doc",
												Name: "смета-подрядчик.pdf",
												Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08/смета-подрядчик.pdf",
											},
										},
									},
									{Parts: []w.CellPart{w.Text("1.2 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("application/pdf")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-05 14:20")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Icon: "download", Small: true, Title: "Download"},
											w.Button{Icon: "dots", Small: true, Title: "More"},
										},
										Acts: true,
									},
								},
							},
						},
					},
					w.PostForm{
						ID:     "files-sel",
						Action: "/app/files/sftpgo/documents/op/carrel/attachments/2026-08",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
					w.PostForm{
						ID:        "post-upload",
						Action:    "/app/files/sftpgo/documents/upload/carrel/attachments/2026-08",
						Fields:    []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
						Multipart: true,
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("Drop files onto the window to upload them here. A name already taken is refused, not overwritten."),
						},
						ScriptOnly: true,
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("Without scripts there is no drop and no queue: "),
							w.Strong("Upload"),
							w.Text(" sends the files. A name already taken is refused, not overwritten."),
						},
						NoScript: true,
					},
					w.Queue{
						Title:    "Uploading {done} of {count}",
						Total:    w.Mono{Text: "{sent}", Variant: w.MonoDusty, Place: w.Pushed},
						Toggle:   w.Button{Label: "Hide", Small: true},
						Progress: "post-upload",
						Row:      &w.QueueItem{Name: "{name}", Meta: "{size} · {percent}%", Meter: &w.Meter{Percent: 0, Track: 100}},
						Taken: &w.QueueItem{
							Name:   "{name}",
							Meta:   "name taken",
							Failed: true,
							Acts: []w.Button{
								w.Button{Label: "Keep both", Small: true, Name: "answer", Value: "keep"},
								w.Button{Label: "Replace", Small: true, Name: "answer", Value: "replace"},
								w.Button{Label: "Skip", Small: true, Name: "answer", Value: "skip"},
							},
						},
						Fold: &w.QueueFold{Hide: "Hide", Show: "Show", Total: "{percent}%"},
					},
				},
			},
			Side: &w.Side{
				Items: []w.SidePart{
					w.Head{
						Title: &w.HeadTitle{Text: "2026-08-10-sync-board.png", Size: w.HeadTitle17},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("sftpgo · vps · Documents")}}},
						Acts:  &w.Acts{Items: []w.ActsPart{w.Close{Label: "Close the panel"}}},
						Pad:   w.HeadSnug,
					},
					w.Pad{Items: []w.PadPart{w.Thumb{Icon: "img", Form: w.ThumbPreview}}},
					w.SideActs{
						Items: []w.Button{
							w.Button{Label: "Download", Icon: "download", Kind: w.ButtonPrimary},
							w.Button{Label: "Open"},
							w.Button{Label: "Rename", Icon: "pencil"},
							w.Button{Label: "Move", Icon: "move"},
						},
					},
					w.Fields{
						Items: []w.FieldsItem{
							{Term: "Size", Value: []w.ValuePart{w.Text("412 KB")}, Face: w.ValueMono},
							{Term: "Type", Value: []w.ValuePart{w.Text("image/png")}, Face: w.ValueMono},
							{Term: "Dimensions", Value: []w.ValuePart{w.Text("1920×1080")}, Face: w.ValueMono},
							{Term: "Changed", Value: []w.ValuePart{w.Text("2026-08-10 11:07")}, Face: w.ValueMono},
							{
								Term:  "Path",
								Value: []w.ValuePart{w.Text("/dav/Documents/carrel/attachments/2026-08/2026-08-10-sync-board.png")},
								Face:  w.ValueMonoSmall,
							},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Links"},
							w.SecNote{Text: "Notes or events may hold it as an attachment: Carrel does not look through the records."},
						},
					},
					w.Danger{Button: w.Button{Label: "Delete", Icon: "trash", Kind: w.ButtonDangerSolid}},
				},
			},
		},
	}
}

// Files2 is frame «Файлы · выделено несколько, идёт загрузка» (s-files-2).
func Files2() Screen {
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
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 10}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 5,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked: w.TickedPart,
							Kind:   w.TickRowRoot,
							Icon:   "file",
							Label:  "Selected collections",
							Meta:   "2 of 3",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "sftpgo · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Documents", Meta: "rw", Current: true},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#9A9280"},
							Label:  "Public",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Files", Meta: "rw"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Shortcuts"},
						w.Tree{
							Rows: []w.TreeRow{
								{Icon: "file", Label: "Attachments", Trail: w.TreeNote{Text: "Documents / carrel"}},
								{Icon: "link", Label: "Published", Trail: w.TreeNote{Text: "2"}},
							},
							On:      1,
							Variant: w.TreePadded,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New folder"},
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
						Title: &w.HeadTitle{Text: "2026-08"},
						Crumbs: &w.Crumbs{
							Path:    []w.Crumb{{Label: "Documents", Href: "#"}, {Label: "carrel", Href: "#"}, {Label: "attachments", Href: "#"}},
							Current: "2026-08",
						},
						Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("sftpgo · vps · 4 items")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Upload", Menu: true, Icon: "upload", Kind: w.ButtonPrimary},
								w.Button{Label: "New folder", Icon: "plus"},
								w.Button{Icon: "refresh", Title: "Refresh"},
								w.Button{Icon: "panel", Title: "Choose one file to see its details.", State: w.ButtonDisabled},
								w.Button{Icon: "dots", Title: "More"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Mono{Text: "3.2 MB", Variant: w.MonoMuted},
							w.BarSep{},
							w.Button{
								Folds: true,
								Label: "Download",
								Icon:  "download",
								Title: "Download",
								Form:  "files-sel",
								Name:  "op",
								Value: "download",
							},
							w.Button{
								Folds: true,
								Label: "Rename",
								Icon:  "pencil",
								Title: "Rename",
								State: w.ButtonOff,
								Form:  "files-sel",
								Name:  "op",
								Value: "rename",
							},
							w.Button{
								Folds: true,
								Label: "Move",
								Icon:  "move",
								Title: "Move",
								Form:  "files-sel",
								Name:  "op",
								Value: "move",
							},
							w.Button{Folds: true, Label: "Copy", Icon: "dup", Title: "Copy", Form: "files-sel", Name: "op", Value: "copy"},
							w.Button{
								Folds: true,
								Label: "Delete",
								Icon:  "trash",
								Kind:  w.ButtonDangerSolid,
								Title: "Delete",
								Form:  "files-sel",
								Name:  "op",
								Value: "delete",
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
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter this folder", Short: "Filter", Label: "Filter this folder"},
								Waiting: true,
								Variant: w.FilterH26W190,
							},
							w.Seg{
								Options: []w.SegOption{
									{Label: "Name", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08?sort=name", On: true},
									{Label: "Size", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08?sort=size"},
									{Label: "Changed", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08?sort=changed"},
								},
								Secondary: true,
							},
							w.Seg{
								Options: []w.SegOption{
									{
										Icon:  "menu",
										Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-08?view=list",
										On:    true,
										Title: "List",
									},
									{
										Icon:  "grid",
										Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-08?view=tiles",
										Title: "Tiles",
									},
								},
							},
							w.Button{Label: "Columns", Icon: "cols"},
							w.Drop{
								Trigger: w.Button{Icon: "dots", Title: "More"},
								More:    true,
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuItem{Label: "Name", On: true, Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08?sort=name"},
										w.MenuItem{Label: "Size", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08?sort=size"},
										w.MenuItem{Label: "Changed", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08?sort=changed"},
									},
									Pop:    w.MenuUnderEnd,
									Inline: true,
								},
							},
						},
						Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "4 items · 4.4 MB"}}, Secondary: true},
					},
					w.Table{
						Cols: []w.TableCol{
							{Width: w.ColTick, Tick: ptr(w.TickedPart)},
							{Label: "Name"},
							{Label: "Size", Width: w.ColSize, Rank: w.Rank3},
							{Label: "Type", Width: w.ColType, Rank: w.Rank4, Secondary: true},
							{Label: "Changed", Width: w.ColTime, Main: true},
							{Width: w.ColNum},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{},
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon:   "back",
												Name:   "up one folder",
												Accent: true,
												Href:   "/app/files/sftpgo/documents/carrel/attachments",
											},
										},
										Span: 5,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{
										Pick: ptr(w.TickedOn),
										Send: w.Send{Of: "files-sel", Name: "item", Value: "2026-08-04-whiteboard.jpg"},
									},
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon: "img",
												Name: "2026-08-04-whiteboard.jpg",
												Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08/2026-08-04-whiteboard.jpg",
											},
										},
									},
									{Parts: []w.CellPart{w.Text("2.8 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/jpeg")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-04 17:03")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Icon: "download", Small: true, Title: "Download"},
											w.Button{Icon: "dots", Small: true, Title: "More"},
										},
										Acts: true,
									},
								},
								State: w.RowPicked,
							},
							{
								Cells: []w.TableCell{
									{
										Pick: ptr(w.TickedOn),
										Send: w.Send{Of: "files-sel", Name: "item", Value: "2026-08-10-sync-board.png"},
									},
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon: "img",
												Name: "2026-08-10-sync-board.png",
												Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08/2026-08-10-sync-board.png",
											},
										},
									},
									{Parts: []w.CellPart{w.Text("412 KB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/png")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-10 11:07")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Icon: "download", Small: true, Title: "Download"},
											w.Button{Icon: "dots", Small: true, Title: "More"},
										},
										Acts: true,
									},
								},
								State: w.RowPicked,
							},
							{
								Cells: []w.TableCell{
									{Pick: ptr(w.TickedOn), Send: w.Send{Of: "files-sel", Name: "item", Value: "release-checklist.md"}},
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon: "doc",
												Name: "release-checklist.md",
												Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08/release-checklist.md",
											},
										},
									},
									{Parts: []w.CellPart{w.Text("3 KB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("text/markdown")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-09 18:44")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Icon: "download", Small: true, Title: "Download"},
											w.Button{Icon: "dots", Small: true, Title: "More"},
										},
										Acts: true,
									},
								},
								State: w.RowPicked,
							},
							{
								Cells: []w.TableCell{
									{Pick: ptr(w.TickedOff), Send: w.Send{Of: "files-sel", Name: "item", Value: "смета-подрядчик.pdf"}},
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon: "doc",
												Name: "смета-подрядчик.pdf",
												Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08/смета-подрядчик.pdf",
											},
										},
									},
									{Parts: []w.CellPart{w.Text("1.2 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("application/pdf")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-05 14:20")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Icon: "download", Small: true, Title: "Download"},
											w.Button{Icon: "dots", Small: true, Title: "More"},
										},
										Acts: true,
									},
								},
							},
						},
					},
					w.PostForm{
						ID:     "files-sel",
						Action: "/app/files/sftpgo/documents/op/carrel/attachments/2026-08",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
					w.PostForm{
						ID:        "post-upload",
						Action:    "/app/files/sftpgo/documents/upload/carrel/attachments/2026-08",
						Fields:    []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
						Multipart: true,
					},
					w.PostForm{
						ID:        "taken-release",
						Action:    "/app/files/sftpgo/documents/upload/carrel/attachments/2026-08?taken=release-checklist.md",
						Fields:    []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
						Multipart: true,
					},
					w.Queue{
						Title:  "Uploading 2 of 4",
						Total:  w.Mono{Text: "12.4 MB", Variant: w.MonoDusty, Place: w.Pushed},
						Toggle: w.Button{Label: "Hide", Small: true},
						Items: []w.QueueItem{
							{Name: "отчёт-июль.pdf", Meta: "4.1 MB · done", Meter: &w.Meter{Percent: 100, Track: 100}},
							{Name: "плакат.png", Meta: "8.3 MB · 62%", Meter: &w.Meter{Percent: 62, Track: 100}},
							{
								Name:   "release-checklist.md",
								Meta:   "name taken",
								Failed: true,
								Acts: []w.Button{
									w.Button{
										Label: "Keep both",
										Small: true,
										File:  &w.FileField{Form: "taken-release", Name: "keep", Label: "Keep both"},
									},
									w.Button{
										Label: "Replace",
										Small: true,
										File:  &w.FileField{Form: "taken-release", Name: "replace", Label: "Replace"},
									},
									w.Button{Label: "Send", Small: true, Form: "taken-release", NoScript: true},
									w.Button{Label: "Skip", Small: true, Form: "taken-release", Name: "answer", Value: "skip"},
								},
							},
							{Name: "архив-2025.zip", Meta: "server requires the file size", Failed: true},
						},
						Progress: "post-upload",
						Row:      &w.QueueItem{Name: "{name}", Meta: "{size} · {percent}%", Meter: &w.Meter{Percent: 0, Track: 100}},
						Taken: &w.QueueItem{
							Name:   "{name}",
							Meta:   "name taken",
							Failed: true,
							Acts: []w.Button{
								w.Button{Label: "Keep both", Small: true, Name: "answer", Value: "keep"},
								w.Button{Label: "Replace", Small: true, Name: "answer", Value: "replace"},
								w.Button{Label: "Skip", Small: true, Name: "answer", Value: "skip"},
							},
						},
						Fold: &w.QueueFold{Hide: "Hide", Show: "Show", Total: "{percent}%"},
					},
				},
				Variant: w.MainRoomy,
			},
		},
	}
}

// Files4 is frame «Перемещение и копирование» (s-files-4).
func Files4() Screen {
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
							{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
						},
						On: 5,
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.TickRow{
								Ticked: w.TickedPart,
								Kind:   w.TickRowRoot,
								Icon:   "file",
								Label:  "Selected collections",
								Meta:   "2 of 3",
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "sftpgo · vps"},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Documents", Meta: "rw", Current: true},
							w.TickRow{
								Ticked: w.TickedOff,
								Mark:   w.Mark{Colour: "#9A9280"},
								Label:  "Public",
								Hidden: true,
								Trail:  w.Badge{Label: "ro"},
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Files", Meta: "rw"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "Shortcuts"},
							w.Tree{
								Rows: []w.TreeRow{
									{Icon: "file", Label: "Attachments", Trail: w.TreeNote{Text: "Documents / carrel"}},
									{Icon: "link", Label: "Published", Trail: w.TreeNote{Text: "2"}},
								},
								On:      1,
								Variant: w.TreePadded,
							},
						},
					},
					w.RailFoot{
						Items: []w.RailFootPart{
							w.Button{Icon: "plus", Title: "New folder"},
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
							Title: &w.HeadTitle{Text: "2026-08"},
							Crumbs: &w.Crumbs{
								Path:    []w.Crumb{{Label: "Documents", Href: "#"}, {Label: "carrel", Href: "#"}, {Label: "attachments", Href: "#"}},
								Current: "2026-08",
							},
							Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("sftpgo · vps · 4 items")}}},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "Upload", Menu: true, Icon: "upload", Kind: w.ButtonPrimary},
									w.Button{Label: "New folder", Icon: "plus"},
									w.Button{Icon: "refresh", Title: "Refresh"},
									w.Button{Icon: "panel", Title: "Details"},
									w.Button{Icon: "dots", Title: "More"},
								},
							},
						},
						w.Bar{
							Items: []w.BarPart{
								w.Mono{Text: "3.2 MB", Variant: w.MonoMuted},
								w.BarSep{},
								w.Button{
									Folds: true,
									Label: "Download",
									Icon:  "download",
									Title: "Download",
									Form:  "files-sel",
									Name:  "op",
									Value: "download",
								},
								w.Button{
									Folds: true,
									Label: "Rename",
									Icon:  "pencil",
									Title: "Rename",
									State: w.ButtonOff,
									Form:  "files-sel",
									Name:  "op",
									Value: "rename",
								},
								w.Button{
									Folds: true,
									Label: "Move",
									Icon:  "move",
									Title: "Move",
									Form:  "files-sel",
									Name:  "op",
									Value: "move",
								},
								w.Button{Folds: true, Label: "Copy", Icon: "dup", Title: "Copy", Form: "files-sel", Name: "op", Value: "copy"},
								w.Button{
									Folds: true,
									Label: "Delete",
									Icon:  "trash",
									Kind:  w.ButtonDangerSolid,
									Title: "Delete",
									Form:  "files-sel",
									Name:  "op",
									Value: "delete",
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
						w.Bar{
							Items: []w.BarPart{
								w.Filter{
									Query:   w.Query{Placeholder: "Filter this folder", Short: "Filter", Label: "Filter this folder"},
									Waiting: true,
									Variant: w.FilterH26W190,
								},
								w.Seg{
									Options: []w.SegOption{
										{Label: "Name", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08?sort=name", On: true},
										{Label: "Size", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08?sort=size"},
										{Label: "Changed", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08?sort=changed"},
									},
									Secondary: true,
								},
								w.Seg{
									Options: []w.SegOption{
										{
											Icon:  "menu",
											Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-08?view=list",
											On:    true,
											Title: "List",
										},
										{
											Icon:  "grid",
											Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-08?view=tiles",
											Title: "Tiles",
										},
									},
								},
								w.Button{Label: "Columns", Icon: "cols"},
								w.Drop{
									Trigger: w.Button{Icon: "dots", Title: "More"},
									More:    true,
									Menu: w.Menu{
										Items: []w.MenuPart{
											w.MenuItem{Label: "Name", On: true, Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08?sort=name"},
											w.MenuItem{Label: "Size", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08?sort=size"},
											w.MenuItem{Label: "Changed", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08?sort=changed"},
										},
										Pop:    w.MenuUnderEnd,
										Inline: true,
									},
								},
							},
							Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "4 items · 4.4 MB"}}, Secondary: true},
						},
						w.Table{
							Cols: []w.TableCol{
								{Width: w.ColTick, Tick: ptr(w.TickedPart)},
								{Label: "Name"},
								{Label: "Size", Width: w.ColSize, Rank: w.Rank3},
								{Label: "Type", Width: w.ColType, Rank: w.Rank4, Secondary: true},
								{Label: "Changed", Width: w.ColTime, Main: true},
								{Width: w.ColNum},
							},
							Rows: []w.TableRow{
								{
									Cells: []w.TableCell{
										{},
										{
											Parts: []w.CellPart{
												w.TableFile{
													Icon:   "back",
													Name:   "up one folder",
													Accent: true,
													Href:   "/app/files/sftpgo/documents/carrel/attachments",
												},
											},
											Span: 5,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{
											Pick: ptr(w.TickedOn),
											Send: w.Send{Of: "files-sel", Name: "item", Value: "2026-08-04-whiteboard.jpg"},
										},
										{
											Parts: []w.CellPart{
												w.TableFile{
													Icon: "img",
													Name: "2026-08-04-whiteboard.jpg",
													Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08/2026-08-04-whiteboard.jpg",
												},
											},
										},
										{Parts: []w.CellPart{w.Text("2.8 MB")}, Mono: true},
										{Parts: []w.CellPart{w.Text("image/jpeg")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-04 17:03")}, Mono: true},
										{
											Parts: []w.CellPart{
												w.Button{Icon: "download", Small: true, Title: "Download"},
												w.Button{Icon: "dots", Small: true, Title: "More"},
											},
											Acts: true,
										},
									},
									State: w.RowPicked,
								},
								{
									Cells: []w.TableCell{
										{
											Pick: ptr(w.TickedOn),
											Send: w.Send{Of: "files-sel", Name: "item", Value: "2026-08-10-sync-board.png"},
										},
										{
											Parts: []w.CellPart{
												w.TableFile{
													Icon: "img",
													Name: "2026-08-10-sync-board.png",
													Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08/2026-08-10-sync-board.png",
												},
											},
										},
										{Parts: []w.CellPart{w.Text("412 KB")}, Mono: true},
										{Parts: []w.CellPart{w.Text("image/png")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-10 11:07")}, Mono: true},
										{
											Parts: []w.CellPart{
												w.Button{Icon: "download", Small: true, Title: "Download"},
												w.Button{Icon: "dots", Small: true, Title: "More"},
											},
											Acts: true,
										},
									},
									State: w.RowPicked,
								},
								{
									Cells: []w.TableCell{
										{Pick: ptr(w.TickedOn), Send: w.Send{Of: "files-sel", Name: "item", Value: "release-checklist.md"}},
										{
											Parts: []w.CellPart{
												w.TableFile{
													Icon: "doc",
													Name: "release-checklist.md",
													Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08/release-checklist.md",
												},
											},
										},
										{Parts: []w.CellPart{w.Text("3 KB")}, Mono: true},
										{Parts: []w.CellPart{w.Text("text/markdown")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-09 18:44")}, Mono: true},
										{
											Parts: []w.CellPart{
												w.Button{Icon: "download", Small: true, Title: "Download"},
												w.Button{Icon: "dots", Small: true, Title: "More"},
											},
											Acts: true,
										},
									},
									State: w.RowPicked,
								},
								{
									Cells: []w.TableCell{
										{Pick: ptr(w.TickedOff), Send: w.Send{Of: "files-sel", Name: "item", Value: "смета-подрядчик.pdf"}},
										{
											Parts: []w.CellPart{
												w.TableFile{
													Icon: "doc",
													Name: "смета-подрядчик.pdf",
													Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08/смета-подрядчик.pdf",
												},
											},
										},
										{Parts: []w.CellPart{w.Text("1.2 MB")}, Mono: true},
										{Parts: []w.CellPart{w.Text("application/pdf")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-05 14:20")}, Mono: true},
										{
											Parts: []w.CellPart{
												w.Button{Icon: "download", Small: true, Title: "Download"},
												w.Button{Icon: "dots", Small: true, Title: "More"},
											},
											Acts: true,
										},
									},
								},
							},
						},
						w.PostForm{
							ID:     "files-sel",
							Action: "/app/files/sftpgo/documents/op/carrel/attachments/2026-08",
							Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
						},
					},
				},
				Height: w.ShellHeight560,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					ID:     "move-sheet",
					Action: "/app/files/sftpgo/documents/op/carrel/attachments/2026-08?op=move",
					Icon:   "move",
					Title:  "Move 3 items",
					Head: []w.SheetHeadPart{
						w.Drop{
							Trigger: w.Pick{Value: "sftpgo · vps", Label: "Server", Variant: w.PickT12W240},
							Menu: w.Menu{
								Items: []w.MenuPart{
									w.MenuItem{
										Label: "sftpgo · vps",
										On:    true,
										Href:  "/app/files/sftpgo/documents/op/carrel/attachments/2026-08?op=move&server=sftpgo",
										Swap:  "move-sheet",
									},
									w.MenuItem{
										Label: "nextcloud",
										Href:  "/app/files/sftpgo/documents/op/carrel/attachments/2026-08?op=move&server=nextcloud",
										Swap:  "move-sheet",
									},
									w.MenuItem{Label: "webdav · nas", Reason: "webdav · nas did not answer.", Href: "#"},
								},
								Width: w.MenuWidth290,
								Pop:   w.MenuUnderEnd,
							},
							Place: w.Pushed,
						},
						w.Dim{
							Parts:   []w.Inline{w.Text("from Documents / carrel / attachments / 2026-08")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
					},
					Body: []w.SheetPart{
						w.Rubric{
							Text:  "Where to put it",
							Items: []w.RubricPart{w.Button{Label: "New folder", Icon: "plus", Small: true}},
						},
						w.Tree{
							Rows: []w.TreeRow{
								{
									Icon:  "file",
									Label: "Documents",
									Send:  w.Send{Name: "to", Value: "sftpgo/documents", One: true},
									Href:  "/app/files/sftpgo/documents/op/carrel/attachments/2026-08?op=move&to=sftpgo/documents",
									Swap:  "move-sheet",
								},
								{
									Icon:   "file",
									Label:  "carrel",
									Indent: w.TreeLevel22,
									Send:   w.Send{Name: "to", Value: "sftpgo/documents/carrel", One: true},
									Href:   "/app/files/sftpgo/documents/op/carrel/attachments/2026-08?op=move&to=sftpgo/documents/carrel",
									Swap:   "move-sheet",
								},
								{
									Icon:   "file",
									Label:  "attachments",
									Indent: w.TreeLevel40,
									Send:   w.Send{Name: "to", Value: "sftpgo/documents/carrel/attachments", One: true},
									Href:   "/app/files/sftpgo/documents/op/carrel/attachments/2026-08?op=move&to=sftpgo/documents/carrel/attachments",
									Swap:   "move-sheet",
								},
								{
									Icon:   "file",
									Label:  "2026-07",
									Indent: w.TreeLevel58,
									Send:   w.Send{Name: "to", Value: "sftpgo/documents/carrel/attachments/2026-07", One: true},
									Href:   "/app/files/sftpgo/documents/op/carrel/attachments/2026-08?op=move&to=sftpgo/documents/carrel/attachments/2026-07",
									Swap:   "move-sheet",
								},
								{
									Icon:   "file",
									Label:  "2026-08",
									Indent: w.TreeLevel58,
									Send:   w.Send{Name: "to", Value: "sftpgo/documents/carrel/attachments/2026-08", One: true},
									Href:   "/app/files/sftpgo/documents/op/carrel/attachments/2026-08?op=move&to=sftpgo/documents/carrel/attachments/2026-08",
									Swap:   "move-sheet",
									Trail:  w.TreeNote{Text: "the items are here"},
								},
								{
									Icon:        "file",
									Label:       "Public",
									Unavailable: true,
									Send:        w.Send{Name: "to", Value: "sftpgo/public", One: true},
									Trail:       w.Badge{Label: "read-only"},
								},
							},
							On: 4,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Move", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Button{Label: "Open the folder", Name: "open", Value: "1", Submit: true, NoScript: true},
						w.Line{
							Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}},
							Gap:   w.LineGapNormal,
							Place: w.Pushed,
						},
					},
					Width: w.SheetWidth520,
				},
				Pad: w.ScrimPad26,
			},
		},
	}
}

// Files5 is frame «Файлы · плитки» (s-files-5).
func Files5() Screen {
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
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 10}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 5,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked: w.TickedPart,
							Kind:   w.TickRowRoot,
							Icon:   "file",
							Label:  "Selected collections",
							Meta:   "2 of 3",
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "sftpgo · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Documents", Meta: "rw", Current: true},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#9A9280"},
							Label:  "Public",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Files", Meta: "rw"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Shortcuts"},
						w.Tree{
							Rows: []w.TreeRow{
								{Icon: "file", Label: "Attachments", Trail: w.TreeNote{Text: "Documents / carrel"}},
								{Icon: "link", Label: "Published", Trail: w.TreeNote{Text: "2"}},
							},
							On:      1,
							Variant: w.TreePadded,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New folder"},
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
						Title: &w.HeadTitle{Text: "2026-07"},
						Crumbs: &w.Crumbs{
							Path:    []w.Crumb{{Label: "Documents", Href: "#"}, {Label: "carrel", Href: "#"}, {Label: "attachments", Href: "#"}},
							Current: "2026-07",
						},
						Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("sftpgo · vps · 7 items")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Upload", Menu: true, Icon: "upload", Kind: w.ButtonPrimary},
								w.Button{Label: "New folder", Icon: "plus"},
								w.Button{Icon: "refresh", Title: "Refresh"},
								w.Button{Icon: "panel", Title: "Details"},
								w.Button{Icon: "dots", Title: "More"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Mono{Text: "2.8 MB", Variant: w.MonoMuted},
							w.BarSep{},
							w.Button{
								Folds: true,
								Label: "Download",
								Icon:  "download",
								Title: "Download",
								Form:  "files-sel",
								Name:  "op",
								Value: "download",
							},
							w.Button{
								Folds: true,
								Label: "Rename",
								Icon:  "pencil",
								Title: "Rename",
								Form:  "files-sel",
								Name:  "op",
								Value: "rename",
							},
							w.Button{
								Folds: true,
								Label: "Move",
								Icon:  "move",
								Title: "Move",
								Form:  "files-sel",
								Name:  "op",
								Value: "move",
							},
							w.Button{Folds: true, Label: "Copy", Icon: "dup", Title: "Copy", Form: "files-sel", Name: "op", Value: "copy"},
							w.Button{
								Folds: true,
								Label: "Delete",
								Icon:  "trash",
								Kind:  w.ButtonDangerSolid,
								Title: "Delete",
								Form:  "files-sel",
								Name:  "op",
								Value: "delete",
							},
						},
						Count: &w.BarCount{Text: "1 selected"},
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
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter this folder", Short: "Filter", Label: "Filter this folder"},
								Waiting: true,
								Variant: w.FilterH26W190,
							},
							w.Seg{
								Options: []w.SegOption{
									{Label: "Name", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-07?sort=name", On: true},
									{Label: "Size", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-07?sort=size"},
									{Label: "Changed", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-07?sort=changed"},
								},
								Secondary: true,
							},
							w.Seg{
								Options: []w.SegOption{
									{
										Icon:  "menu",
										Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-07?view=list",
										Title: "List",
									},
									{
										Icon:  "grid",
										Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-07?view=tiles",
										On:    true,
										Title: "Tiles",
									},
								},
							},
							w.Drop{
								Trigger: w.Button{Icon: "dots", Title: "More"},
								More:    true,
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuItem{Label: "Name", On: true, Href: "/app/files/sftpgo/documents/carrel/attachments/2026-07?sort=name"},
										w.MenuItem{Label: "Size", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-07?sort=size"},
										w.MenuItem{Label: "Changed", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-07?sort=changed"},
									},
									Pop:    w.MenuUnderEnd,
									Inline: true,
								},
							},
						},
						Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "7 items · 89.4 MB"}}, Secondary: true},
					},
					w.Grid{
						Tiles: []w.Tile{
							w.Tile{
								Thumb: w.Thumb{Icon: "file", Form: w.ThumbFolder},
								Name:  "архив",
								Meta:  "",
								Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-07/архив",
								Send:  w.Send{Of: "files-sel", Name: "item", Value: "архив"},
							},
							w.Tile{
								Thumb: w.Thumb{Icon: "file", Form: w.ThumbFolder},
								Name:  "черновики",
								Meta:  "0 items",
								Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-07/черновики",
								Send:  w.Send{Of: "files-sel", Name: "item", Value: "черновики"},
							},
							w.Tile{
								Thumb:    w.Thumb{Icon: "img"},
								Name:     "whiteboard.jpg",
								Meta:     "2.8 MB · 4032×3024",
								Selected: true,
								Href:     "/app/files/sftpgo/documents/carrel/attachments/2026-07/whiteboard.jpg",
								Send:     w.Send{Of: "files-sel", Name: "item", Value: "whiteboard.jpg"},
							},
							w.Tile{
								Thumb: w.Thumb{Icon: "img", Form: w.ThumbOnCard},
								Name:  "панорама.tif",
								Meta:  "84 MB · too large",
								Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-07/панорама.tif",
								Send:  w.Send{Of: "files-sel", Name: "item", Value: "панорама.tif"},
							},
							w.Tile{
								Thumb: w.Thumb{Icon: "doc", Form: w.ThumbOnCard},
								Name:  "смета.pdf",
								Meta:  "1.2 MB",
								Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-07/смета.pdf",
								Send:  w.Send{Of: "files-sel", Name: "item", Value: "смета.pdf"},
							},
							w.Tile{
								Thumb: w.Thumb{Icon: "img"},
								Name:  "схема-1.png",
								Meta:  "640 KB",
								Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-07/схема-1.png",
								Send:  w.Send{Of: "files-sel", Name: "item", Value: "схема-1.png"},
							},
							w.Tile{
								Thumb: w.Thumb{Icon: "img"},
								Name:  "схема-2.png",
								Meta:  "712 KB",
								Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-07/схема-2.png",
								Send:  w.Send{Of: "files-sel", Name: "item", Value: "схема-2.png"},
							},
						},
					},
					w.PostForm{
						ID:     "files-sel",
						Action: "/app/files/sftpgo/documents/op/carrel/attachments/2026-07",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
				},
			},
		},
	}
}

// Files6 is frame «Файлы · корень «Selected collections»» (s-files-6).
func Files6() Screen {
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
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 10}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 5}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 2}, Href: "#"},
					},
					On: 5,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedPart,
							Kind:    w.TickRowRoot,
							Icon:    "file",
							Label:   "Selected collections",
							Meta:    "2 of 3",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "sftpgo · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Documents", Meta: "rw"},
						w.TickRow{
							Ticked: w.TickedOff,
							Mark:   w.Mark{Colour: "#9A9280"},
							Label:  "Public",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Files", Meta: "rw"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Shortcuts"},
						w.Tree{
							Rows: []w.TreeRow{
								{Icon: "file", Label: "Attachments", Trail: w.TreeNote{Text: "Documents / carrel"}},
								{Icon: "link", Label: "Published", Trail: w.TreeNote{Text: "2"}},
							},
							Variant: w.TreePadded,
						},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "Open a folder to make one in it.", State: w.ButtonDisabled},
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
						Title: &w.HeadTitle{Text: "Selected collections", Size: w.HeadTitle18},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("2 of 3 storages on 2 servers")}}},
						Pad:   w.HeadSnug,
					},
					w.Table{
						Cols: []w.TableCol{
							{Label: "Name"},
							{Label: "Server", Width: w.ColTime, Main: true},
							{Label: "Access", Width: w.Col80, Rank: w.Rank3},
							{Label: "Purpose", Width: w.ColTime, Rank: w.Rank4},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{
										Parts: []w.CellPart{w.TableFile{Icon: "file", Name: "Documents", Href: "/app/files/sftpgo/documents"}},
									},
									{Parts: []w.CellPart{w.Text("sftpgo · vps")}, Mono: true},
									{Parts: []w.CellPart{w.Text("rw")}, Mono: true},
									{Parts: []w.CellPart{w.Badge{Label: "attachments"}}},
								},
								State: w.RowFolder,
							},
							{
								Cells: []w.TableCell{
									{Parts: []w.CellPart{w.TableFile{Icon: "file", Name: "Files", Href: "/app/files/nextcloud/files"}}},
									{Parts: []w.CellPart{w.Text("nextcloud")}, Mono: true},
									{Parts: []w.CellPart{w.Text("rw")}, Mono: true},
									{},
								},
								State: w.RowFolder,
							},
						},
					},
					w.Msg{Parts: []w.MsgPart{w.Text("The storages ticked in the rail, each with its own tree.")}},
				},
			},
		},
	}
}
