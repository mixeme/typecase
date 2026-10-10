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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
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
						Title: &w.HeadTitle{Text: "attachments"},
						Crumbs: &w.Crumbs{
							Path:    []w.Crumb{{Label: "Documents", Href: "#"}, {Label: "carrel", Href: "#"}},
							Current: "attachments",
						},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{w.Text("sftpgo · vps · "), w.Badge{Label: "attachment folder"}, w.Text(" · 5 items")},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Upload", Menu: true, Icon: "up", Kind: w.ButtonPrimary},
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
							w.Button{Folds: true, Label: "Download", Icon: "down", Title: "Download"},
							w.Button{Folds: true, Label: "Rename", Icon: "pencil", Title: "Rename"},
							w.Button{Folds: true, Label: "Move", Icon: "move", Title: "Move"},
							w.Button{Folds: true, Label: "Copy", Icon: "dup", Title: "Copy"},
							w.Button{Folds: true, Label: "Delete", Icon: "trash", Kind: w.ButtonDangerSolid, Title: "Delete"},
						},
						Count:   &w.BarCount{Text: "1 selected"},
						Right:   &w.BarRight{Items: []w.BarPart{w.Button{Label: "Clear selection", Small: true}}},
						Variant: w.BarSelection,
					},
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter this folder", Label: "Filter this folder"},
								Waiting: true,
								Variant: w.FilterH26W190,
							},
							w.Seg{
								Options: []w.SegOption{
									{Label: "Name", Href: "/app/files/sftpgo/documents/carrel/attachments?sort=name", On: true},
									{Label: "Size", Href: "/app/files/sftpgo/documents/carrel/attachments?sort=size"},
									{Label: "Changed", Href: "/app/files/sftpgo/documents/carrel/attachments?sort=changed"},
								},
								Secondary: true,
							},
							w.Seg{
								Options: []w.SegOption{
									{
										Icon:  "menu",
										Href:  "/app/files/sftpgo/documents/carrel/attachments?view=list",
										On:    true,
										Title: "List",
									},
									{Icon: "grid", Href: "/app/files/sftpgo/documents/carrel/attachments?view=tiles", Title: "Tiles"},
								},
							},
							w.Button{Label: "Columns", Icon: "cols"},
							w.Drop{
								Trigger: w.Button{Icon: "dots", Title: "More"},
								More:    true,
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuItem{Label: "Name", On: true, Href: "/app/files/sftpgo/documents/carrel/attachments?sort=name"},
										w.MenuItem{Label: "Size", Href: "/app/files/sftpgo/documents/carrel/attachments?sort=size"},
										w.MenuItem{Label: "Changed", Href: "/app/files/sftpgo/documents/carrel/attachments?sort=changed"},
									},
									Pop:    w.MenuUnderEnd,
									Inline: true,
								},
							},
						},
						Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "5 items · 4.4 MB"}}, Secondary: true},
					},
					w.Table{
						Cols: []w.TableCol{
							{Width: w.ColTick, Tick: ptr(w.TickedOff)},
							{Label: "Name"},
							{Label: "Size", Width: w.ColSize, Rank: w.Rank3},
							{Label: "Type", Width: w.ColType, Rank: w.Rank4, Secondary: true},
							{Label: "Changed", Width: w.ColTime},
							{Width: w.ColNum},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{},
									{
										Parts: []w.CellPart{
											w.TableFile{Icon: "back", Name: "up one folder", Accent: true, Href: "/app/files/sftpgo/documents/carrel"},
										},
										Span: 5,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Pick: ptr(w.TickedOff)},
									{
										Parts: []w.CellPart{
											w.TableFile{Icon: "file", Name: "2026-07", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-07"},
										},
									},
									{Parts: []w.CellPart{w.Text("—")}, Mono: true},
									{Parts: []w.CellPart{w.Text("folder")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-07-01 09:12")}, Mono: true},
									{Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Title: "More"}}, Acts: true},
								},
								State: w.RowFolder,
							},
							{
								Cells: []w.TableCell{
									{Pick: ptr(w.TickedOff)},
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon: "img",
												Name: "2026-08-04-whiteboard.jpg",
												Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08-04-whiteboard.jpg",
											},
										},
									},
									{Parts: []w.CellPart{w.Text("2.8 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/jpeg")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-04 17:03")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Icon: "down", Small: true, Title: "Download"},
											w.Button{Icon: "dots", Small: true, Title: "More"},
										},
										Acts: true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Pick: ptr(w.TickedOn)},
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon: "img",
												Name: "2026-08-10-sync-board.png",
												Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08-10-sync-board.png",
											},
										},
									},
									{Parts: []w.CellPart{w.Text("412 KB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/png")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-10 11:07")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Icon: "down", Small: true, Title: "Download"},
											w.Button{Icon: "dots", Small: true, Title: "More"},
										},
										Acts: true,
									},
								},
								State: w.RowPicked,
							},
							{
								Cells: []w.TableCell{
									{Pick: ptr(w.TickedOff)},
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon: "doc",
												Name: "release-checklist.md",
												Href: "/app/files/sftpgo/documents/carrel/attachments/release-checklist.md",
											},
										},
									},
									{Parts: []w.CellPart{w.Text("3 KB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("text/markdown")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-09 18:44")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Icon: "down", Small: true, Title: "Download"},
											w.Button{Icon: "dots", Small: true, Title: "More"},
										},
										Acts: true,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Pick: ptr(w.TickedOff)},
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon: "doc",
												Name: "смета-подрядчик.pdf",
												Href: "/app/files/sftpgo/documents/carrel/attachments/смета-подрядчик.pdf",
											},
										},
									},
									{Parts: []w.CellPart{w.Text("1.2 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("application/pdf")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-05 14:20")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Icon: "down", Small: true, Title: "Download"},
											w.Button{Icon: "dots", Small: true, Title: "More"},
										},
										Acts: true,
									},
								},
							},
						},
					},
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("Drop files onto the window, or onto a folder in the list. A name that is already taken is refused rather than overwritten, and the queue says so."),
						},
					},
				},
			},
			Side: &w.Side{
				Items: []w.SidePart{
					w.Head{
						Title: &w.HeadTitle{Text: "2026-08-10-sync-board.png", Size: w.HeadTitle17},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("sftpgo · vps · Documents")}}},
						Acts:  &w.Acts{Items: []w.ActsPart{w.Close{Label: "Close the panel"}}},
						Pad:   w.HeadPadBottom8,
					},
					w.Pad{Items: []w.PadPart{w.Thumb{Icon: "img", Form: w.ThumbPreview}}},
					w.SideActs{
						Items: []w.Button{
							w.Button{Label: "Download", Icon: "down", Kind: w.ButtonPrimary},
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
								Value: []w.ValuePart{w.Text("/dav/Documents/carrel/attachments/2026-08-10-sync-board.png")},
								Face:  w.ValueMonoSmall,
							},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Links"},
							w.SecNote{
								Text: "Carrel keeps no index of what points at this file. It may be attached to a note or an event — check before deleting.",
							},
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
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
						Title: &w.HeadTitle{Text: "attachments"},
						Crumbs: &w.Crumbs{
							Path:    []w.Crumb{{Label: "Documents", Href: "#"}, {Label: "carrel", Href: "#"}},
							Current: "attachments",
						},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{w.Text("sftpgo · vps · "), w.Badge{Label: "attachment folder"}, w.Text(" · 5 items")},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Upload", Menu: true, Icon: "up", Kind: w.ButtonPrimary},
								w.Button{Label: "New folder", Icon: "plus"},
								w.Button{Icon: "refresh", Title: "Refresh"},
								w.Button{Icon: "panel", Title: "Details"},
								w.Button{Icon: "dots", Title: "More"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Mono{Text: "4.4 MB", Variant: w.MonoMuted},
							w.BarSep{},
							w.Button{Folds: true, Label: "Download", Icon: "down", Title: "Download"},
							w.Button{Folds: true, Label: "Rename", Icon: "pencil", Title: "Rename", State: w.ButtonOff},
							w.Button{Folds: true, Label: "Move", Icon: "move", Title: "Move"},
							w.Button{Folds: true, Label: "Copy", Icon: "dup", Title: "Copy"},
							w.Button{Folds: true, Label: "Delete", Icon: "trash", Kind: w.ButtonDangerSolid, Title: "Delete"},
						},
						Count:   &w.BarCount{Text: "3 selected"},
						Right:   &w.BarRight{Items: []w.BarPart{w.Button{Label: "Clear selection", Small: true}}},
						Variant: w.BarSelection,
					},
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter this folder", Label: "Filter this folder"},
								Waiting: true,
								Variant: w.FilterH26W190,
							},
							w.Seg{
								Options: []w.SegOption{
									{Label: "Name", Href: "/app/files/sftpgo/documents/carrel/attachments?sort=name", On: true},
									{Label: "Size", Href: "/app/files/sftpgo/documents/carrel/attachments?sort=size"},
									{Label: "Changed", Href: "/app/files/sftpgo/documents/carrel/attachments?sort=changed"},
								},
								Secondary: true,
							},
							w.Seg{
								Options: []w.SegOption{
									{
										Icon:  "menu",
										Href:  "/app/files/sftpgo/documents/carrel/attachments?view=list",
										On:    true,
										Title: "List",
									},
									{Icon: "grid", Href: "/app/files/sftpgo/documents/carrel/attachments?view=tiles", Title: "Tiles"},
								},
							},
							w.Button{Label: "Columns", Icon: "cols"},
							w.Drop{
								Trigger: w.Button{Icon: "dots", Title: "More"},
								More:    true,
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuItem{Label: "Name", On: true, Href: "/app/files/sftpgo/documents/carrel/attachments?sort=name"},
										w.MenuItem{Label: "Size", Href: "/app/files/sftpgo/documents/carrel/attachments?sort=size"},
										w.MenuItem{Label: "Changed", Href: "/app/files/sftpgo/documents/carrel/attachments?sort=changed"},
									},
									Pop:    w.MenuUnderEnd,
									Inline: true,
								},
							},
						},
						Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "5 items · 4.4 MB"}}, Secondary: true},
					},
					w.Table{
						Cols: []w.TableCol{
							{Width: w.ColTick, Tick: ptr(w.TickedPart)},
							{Label: "Name"},
							{Label: "Size", Width: w.ColSize, Rank: w.Rank3},
							{Label: "Type", Width: w.ColType, Rank: w.Rank4, Secondary: true},
							{Label: "Changed", Width: w.ColTime},
							{Width: w.ColNum},
						},
						Rows: []w.TableRow{
							{
								Cells: []w.TableCell{
									{},
									{
										Parts: []w.CellPart{
											w.TableFile{Icon: "back", Name: "up one folder", Accent: true, Href: "/app/files/sftpgo/documents/carrel"},
										},
										Span: 5,
									},
								},
							},
							{
								Cells: []w.TableCell{
									{Pick: ptr(w.TickedOff)},
									{
										Parts: []w.CellPart{
											w.TableFile{Icon: "file", Name: "2026-07", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-07"},
										},
									},
									{Parts: []w.CellPart{w.Text("—")}, Mono: true},
									{Parts: []w.CellPart{w.Text("folder")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-07-01 09:12")}, Mono: true},
									{Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Title: "More"}}, Acts: true},
								},
								State: w.RowFolder,
							},
							{
								Cells: []w.TableCell{
									{Pick: ptr(w.TickedOn)},
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon: "img",
												Name: "2026-08-04-whiteboard.jpg",
												Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08-04-whiteboard.jpg",
											},
										},
									},
									{Parts: []w.CellPart{w.Text("2.8 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/jpeg")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-04 17:03")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Icon: "down", Small: true, Title: "Download"},
											w.Button{Icon: "dots", Small: true, Title: "More"},
										},
										Acts: true,
									},
								},
								State: w.RowPicked,
							},
							{
								Cells: []w.TableCell{
									{Pick: ptr(w.TickedOn)},
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon: "img",
												Name: "2026-08-10-sync-board.png",
												Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08-10-sync-board.png",
											},
										},
									},
									{Parts: []w.CellPart{w.Text("412 KB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("image/png")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-10 11:07")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Icon: "down", Small: true, Title: "Download"},
											w.Button{Icon: "dots", Small: true, Title: "More"},
										},
										Acts: true,
									},
								},
								State: w.RowPicked,
							},
							{
								Cells: []w.TableCell{
									{Pick: ptr(w.TickedOn)},
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon: "doc",
												Name: "release-checklist.md",
												Href: "/app/files/sftpgo/documents/carrel/attachments/release-checklist.md",
											},
										},
									},
									{Parts: []w.CellPart{w.Text("3 KB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("text/markdown")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-09 18:44")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Icon: "down", Small: true, Title: "Download"},
											w.Button{Icon: "dots", Small: true, Title: "More"},
										},
										Acts: true,
									},
								},
								State: w.RowPicked,
							},
							{
								Cells: []w.TableCell{
									{Pick: ptr(w.TickedOff)},
									{
										Parts: []w.CellPart{
											w.TableFile{
												Icon: "doc",
												Name: "смета-подрядчик.pdf",
												Href: "/app/files/sftpgo/documents/carrel/attachments/смета-подрядчик.pdf",
											},
										},
									},
									{Parts: []w.CellPart{w.Text("1.2 MB")}, Mono: true},
									{Parts: []w.CellPart{w.Text("application/pdf")}, Mono: true},
									{Parts: []w.CellPart{w.Text("2026-08-05 14:20")}, Mono: true},
									{
										Parts: []w.CellPart{
											w.Button{Icon: "down", Small: true, Title: "Download"},
											w.Button{Icon: "dots", Small: true, Title: "More"},
										},
										Acts: true,
									},
								},
							},
						},
					},
					w.Queue{
						Title:  "Uploading 2 of 3",
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
									w.Button{Label: "Keep both", Small: true},
									w.Button{Label: "Replace", Small: true},
									w.Button{Label: "Skip", Small: true},
								},
							},
						},
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
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
							{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
							{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
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
							Title: &w.HeadTitle{Text: "attachments"},
							Crumbs: &w.Crumbs{
								Path:    []w.Crumb{{Label: "Documents", Href: "#"}, {Label: "carrel", Href: "#"}},
								Current: "attachments",
							},
							Subs: []w.HeadSub{
								w.HeadSub{
									Parts: []w.HeadSubPart{w.Text("sftpgo · vps · "), w.Badge{Label: "attachment folder"}, w.Text(" · 5 items")},
								},
							},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "Upload", Menu: true, Icon: "up", Kind: w.ButtonPrimary},
									w.Button{Label: "New folder", Icon: "plus"},
									w.Button{Icon: "refresh", Title: "Refresh"},
									w.Button{Icon: "panel", Title: "Details"},
									w.Button{Icon: "dots", Title: "More"},
								},
							},
						},
						w.Bar{
							Items: []w.BarPart{
								w.Mono{Text: "4.4 MB", Variant: w.MonoMuted},
								w.BarSep{},
								w.Button{Folds: true, Label: "Download", Icon: "down", Title: "Download"},
								w.Button{Folds: true, Label: "Rename", Icon: "pencil", Title: "Rename", State: w.ButtonOff},
								w.Button{Folds: true, Label: "Move", Icon: "move", Title: "Move"},
								w.Button{Folds: true, Label: "Copy", Icon: "dup", Title: "Copy"},
								w.Button{Folds: true, Label: "Delete", Icon: "trash", Kind: w.ButtonDangerSolid, Title: "Delete"},
							},
							Count:   &w.BarCount{Text: "3 selected"},
							Right:   &w.BarRight{Items: []w.BarPart{w.Button{Label: "Clear selection", Small: true}}},
							Variant: w.BarSelection,
						},
						w.Bar{
							Items: []w.BarPart{
								w.Filter{
									Query:   w.Query{Placeholder: "Filter this folder", Label: "Filter this folder"},
									Waiting: true,
									Variant: w.FilterH26W190,
								},
								w.Seg{
									Options: []w.SegOption{
										{Label: "Name", Href: "/app/files/sftpgo/documents/carrel/attachments?sort=name", On: true},
										{Label: "Size", Href: "/app/files/sftpgo/documents/carrel/attachments?sort=size"},
										{Label: "Changed", Href: "/app/files/sftpgo/documents/carrel/attachments?sort=changed"},
									},
									Secondary: true,
								},
								w.Seg{
									Options: []w.SegOption{
										{
											Icon:  "menu",
											Href:  "/app/files/sftpgo/documents/carrel/attachments?view=list",
											On:    true,
											Title: "List",
										},
										{Icon: "grid", Href: "/app/files/sftpgo/documents/carrel/attachments?view=tiles", Title: "Tiles"},
									},
								},
								w.Button{Label: "Columns", Icon: "cols"},
								w.Drop{
									Trigger: w.Button{Icon: "dots", Title: "More"},
									More:    true,
									Menu: w.Menu{
										Items: []w.MenuPart{
											w.MenuItem{Label: "Name", On: true, Href: "/app/files/sftpgo/documents/carrel/attachments?sort=name"},
											w.MenuItem{Label: "Size", Href: "/app/files/sftpgo/documents/carrel/attachments?sort=size"},
											w.MenuItem{Label: "Changed", Href: "/app/files/sftpgo/documents/carrel/attachments?sort=changed"},
										},
										Pop:    w.MenuUnderEnd,
										Inline: true,
									},
								},
							},
							Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "5 items · 4.4 MB"}}, Secondary: true},
						},
						w.Table{
							Cols: []w.TableCol{
								{Width: w.ColTick, Tick: ptr(w.TickedPart)},
								{Label: "Name"},
								{Label: "Size", Width: w.ColSize, Rank: w.Rank3},
								{Label: "Type", Width: w.ColType, Rank: w.Rank4, Secondary: true},
								{Label: "Changed", Width: w.ColTime},
								{Width: w.ColNum},
							},
							Rows: []w.TableRow{
								{
									Cells: []w.TableCell{
										{},
										{
											Parts: []w.CellPart{
												w.TableFile{Icon: "back", Name: "up one folder", Accent: true, Href: "/app/files/sftpgo/documents/carrel"},
											},
											Span: 5,
										},
									},
								},
								{
									Cells: []w.TableCell{
										{Pick: ptr(w.TickedOff)},
										{
											Parts: []w.CellPart{
												w.TableFile{Icon: "file", Name: "2026-07", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-07"},
											},
										},
										{Parts: []w.CellPart{w.Text("—")}, Mono: true},
										{Parts: []w.CellPart{w.Text("folder")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-07-01 09:12")}, Mono: true},
										{Parts: []w.CellPart{w.Button{Icon: "dots", Small: true, Title: "More"}}, Acts: true},
									},
									State: w.RowFolder,
								},
								{
									Cells: []w.TableCell{
										{Pick: ptr(w.TickedOn)},
										{
											Parts: []w.CellPart{
												w.TableFile{
													Icon: "img",
													Name: "2026-08-04-whiteboard.jpg",
													Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08-04-whiteboard.jpg",
												},
											},
										},
										{Parts: []w.CellPart{w.Text("2.8 MB")}, Mono: true},
										{Parts: []w.CellPart{w.Text("image/jpeg")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-04 17:03")}, Mono: true},
										{
											Parts: []w.CellPart{
												w.Button{Icon: "down", Small: true, Title: "Download"},
												w.Button{Icon: "dots", Small: true, Title: "More"},
											},
											Acts: true,
										},
									},
									State: w.RowPicked,
								},
								{
									Cells: []w.TableCell{
										{Pick: ptr(w.TickedOn)},
										{
											Parts: []w.CellPart{
												w.TableFile{
													Icon: "img",
													Name: "2026-08-10-sync-board.png",
													Href: "/app/files/sftpgo/documents/carrel/attachments/2026-08-10-sync-board.png",
												},
											},
										},
										{Parts: []w.CellPart{w.Text("412 KB")}, Mono: true},
										{Parts: []w.CellPart{w.Text("image/png")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-10 11:07")}, Mono: true},
										{
											Parts: []w.CellPart{
												w.Button{Icon: "down", Small: true, Title: "Download"},
												w.Button{Icon: "dots", Small: true, Title: "More"},
											},
											Acts: true,
										},
									},
									State: w.RowPicked,
								},
								{
									Cells: []w.TableCell{
										{Pick: ptr(w.TickedOn)},
										{
											Parts: []w.CellPart{
												w.TableFile{
													Icon: "doc",
													Name: "release-checklist.md",
													Href: "/app/files/sftpgo/documents/carrel/attachments/release-checklist.md",
												},
											},
										},
										{Parts: []w.CellPart{w.Text("3 KB")}, Mono: true},
										{Parts: []w.CellPart{w.Text("text/markdown")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-09 18:44")}, Mono: true},
										{
											Parts: []w.CellPart{
												w.Button{Icon: "down", Small: true, Title: "Download"},
												w.Button{Icon: "dots", Small: true, Title: "More"},
											},
											Acts: true,
										},
									},
									State: w.RowPicked,
								},
								{
									Cells: []w.TableCell{
										{Pick: ptr(w.TickedOff)},
										{
											Parts: []w.CellPart{
												w.TableFile{
													Icon: "doc",
													Name: "смета-подрядчик.pdf",
													Href: "/app/files/sftpgo/documents/carrel/attachments/смета-подрядчик.pdf",
												},
											},
										},
										{Parts: []w.CellPart{w.Text("1.2 MB")}, Mono: true},
										{Parts: []w.CellPart{w.Text("application/pdf")}, Mono: true},
										{Parts: []w.CellPart{w.Text("2026-08-05 14:20")}, Mono: true},
										{
											Parts: []w.CellPart{
												w.Button{Icon: "down", Small: true, Title: "Download"},
												w.Button{Icon: "dots", Small: true, Title: "More"},
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
					Action: "/app/files/sftpgo/documents/op/carrel/attachments?op=move",
					Icon:   "move",
					Title:  "Move 3 items",
					Head: []w.SheetHeadPart{
						w.Pick{Value: "sftpgo · vps", Label: "Server", Variant: w.PickT12W240, Place: w.Pushed},
						w.Dim{
							Parts:   []w.Inline{w.Text("from Documents / carrel / attachments")},
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
								{Icon: "file", Label: "Documents", Send: w.Send{Name: "to", Value: "sftpgo/documents"}},
								{
									Icon:   "file",
									Label:  "carrel",
									Indent: w.TreeLevel22,
									Send:   w.Send{Name: "to", Value: "sftpgo/documents/carrel"},
								},
								{
									Icon:   "file",
									Label:  "attachments",
									Indent: w.TreeLevel40,
									Send:   w.Send{Name: "to", Value: "sftpgo/documents/carrel/attachments"},
								},
								{
									Icon:   "file",
									Label:  "2026-07",
									Indent: w.TreeLevel58,
									Send:   w.Send{Name: "to", Value: "sftpgo/documents/carrel/attachments/2026-07"},
								},
								{
									Icon:        "file",
									Label:       "Public",
									Unavailable: true,
									Send:        w.Send{Name: "to", Value: "sftpgo/public"},
									Trail:       w.Badge{Label: "read-only"},
								},
							},
							On: 4,
						},
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("Inside one server this is "),
								w.Mono{Text: "MOVE", Form: w.MonoCode},
								w.Text(": instant, no traffic. To another storage the file has to be downloaded and uploaded again — Carrel says so before it starts, not after."),
							},
							Space: w.MsgFlushTop2,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Move", Kind: w.ButtonPrimary, Submit: true},
						w.Kbd{Keys: "Enter"},
						w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{w.Button{Label: "Cancel"}, w.Kbd{Keys: "Esc"}}},
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
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
								w.Button{Label: "Upload", Menu: true, Icon: "up", Kind: w.ButtonPrimary},
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
							w.Button{Folds: true, Label: "Download", Icon: "down", Title: "Download"},
							w.Button{Folds: true, Label: "Rename", Icon: "pencil", Title: "Rename"},
							w.Button{Folds: true, Label: "Move", Icon: "move", Title: "Move"},
							w.Button{Folds: true, Label: "Copy", Icon: "dup", Title: "Copy"},
							w.Button{Folds: true, Label: "Delete", Icon: "trash", Kind: w.ButtonDangerSolid, Title: "Delete"},
						},
						Count:   &w.BarCount{Text: "1 selected"},
						Variant: w.BarSelection,
					},
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Placeholder: "Filter this folder", Label: "Filter this folder"},
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
								Thumb:    w.Thumb{Icon: "img"},
								Name:     "whiteboard.jpg",
								Meta:     "2.8 MB · 4032×3024",
								Selected: true,
								Href:     "/app/files/sftpgo/documents/carrel/attachments/2026-07/whiteboard.jpg",
								Send:     w.Send{Name: "item", Value: "whiteboard.jpg"},
							},
							w.Tile{
								Thumb: w.Thumb{Icon: "img"},
								Name:  "схема-1.png",
								Meta:  "640 KB",
								Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-07/схема-1.png",
								Send:  w.Send{Name: "item", Value: "схема-1.png"},
							},
							w.Tile{
								Thumb: w.Thumb{Icon: "img"},
								Name:  "схема-2.png",
								Meta:  "712 KB",
								Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-07/схема-2.png",
								Send:  w.Send{Name: "item", Value: "схема-2.png"},
							},
							w.Tile{
								Thumb: w.Thumb{Icon: "doc", Form: w.ThumbOnCard},
								Name:  "смета.pdf",
								Meta:  "1.2 MB",
								Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-07/смета.pdf",
								Send:  w.Send{Name: "item", Value: "смета.pdf"},
							},
							w.Tile{
								Thumb: w.Thumb{Icon: "file", Form: w.ThumbOnCard},
								Name:  "архив",
								Meta:  "",
								Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-07/архив",
								Send:  w.Send{Name: "item", Value: "архив"},
							},
							w.Tile{
								Thumb: w.Thumb{Icon: "file", Form: w.ThumbOnCard},
								Name:  "черновики",
								Meta:  "0 items",
								Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-07/черновики",
								Send:  w.Send{Name: "item", Value: "черновики"},
							},
							w.Tile{
								Thumb: w.Thumb{Icon: "img", Form: w.ThumbOnCard},
								Name:  "панорама.tif",
								Meta:  "84 MB · too large",
								Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-07/панорама.tif",
								Send:  w.Send{Name: "item", Value: "панорама.tif"},
							},
						},
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
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 7}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 6}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 2}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 4}, Href: "#"},
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
						Pad:   w.HeadPadBottom8,
					},
					w.Table{
						Cols: []w.TableCol{
							{Label: "Name"},
							{Label: "Server", Width: w.ColTime},
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
					w.Msg{
						Parts: []w.MsgPart{
							w.Text("This is where the section starts: the storages you ticked in the rail, as top-level folders. Open one and its own tree opens — contents of different servers are never mixed in one folder, because a file has one path and belongs to its own storage. "),
							w.Strong("Public"),
							w.Text(" is not ticked, so it is not here."),
						},
					},
				},
			},
		},
	}
}
