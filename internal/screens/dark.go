// Раздел макета: 9. Тёмная тема.

package screens

import w "carrel.dev/typecase/widget"

// Dark1 is frame «Календарь · тёмная тема» (dark-1).
func Dark1() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
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
					},
					On: 2,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedPart,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined calendar",
							Meta:    "6",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Дни рождения", Meta: "—", Hidden: true},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "2"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "2"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "1"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "1"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "baikal"},
						w.TickRow{Mark: w.Mark{Colour: "#9A9280"}, Label: "Общий", Error: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Subscriptions", Items: []w.RubricPart{w.Badge{Label: "external"}}},
						w.TickRow{
							Kind:  w.TickRowExternal,
							Mark:  w.Mark{Colour: "#8C6239"},
							Label: "Праздники РФ",
							Meta:  "3",
							Trail: w.Badge{Label: "ro"},
						},
						w.TickRow{
							Ticked: w.TickedOff,
							Kind:   w.TickRowExternal,
							Mark:   w.Mark{Colour: "#6B4E7A"},
							Label:  "Расписание лиги",
							Meta:   "—",
							Hidden: true,
							Trail:  w.Badge{Label: "ro"},
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "релиз"}, w.Tag{Label: "работа"}},
				},
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined calendar"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{
									w.Text("5 of 7 calendars · 3 accounts · 10 – 16 August 2026 · "),
									w.Mono{Text: "Europe/Moscow"},
									w.Text(" · 6 events"),
								},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Export"},
							},
						},
					},
					w.Runbar{
						Meter: w.Meter{Percent: 83, Alert: true},
						Parts: []w.RunbarPart{
							w.Text("Polled 5 of 6 sources · 6 events · "),
							w.Drop{
								Trigger: w.Button{Label: "1 failed", Menu: true, Kind: w.ButtonDanger, Small: true},
								Menu: w.Menu{
									Items: []w.MenuPart{
										w.MenuRefusal{
											Mark:   w.Mark{Colour: "#9A9280", Size: w.MarkH11},
											Source: "baikal · Общий",
											Answer: "PROPFIND /dav.php/calendars/mix/ → 401 Unauthorized: the saved password was refused. Open Connections to update it.",
											Retry:  w.Button{Label: "Retry", Small: true},
										},
									},
									Width: w.MenuWidth420,
									Pop:   w.MenuUnderEnd,
								},
							},
							w.Mono{Text: "from cache · 2 min ago", Place: w.Pushed},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 4}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "10:00",
								Title:  []w.Inline{w.Text("Синк-ап по релизу")},
								Strong: true,
								Detail: "Переговорная 2",
								Badges: []w.Badge{w.Badge{Label: "4 attendees"}, w.Badge{Label: "repeats weekly"}},
								Source: "nextcloud · Личный",
								On:     true,
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#3F6E8C"},
								Time:   "all day",
								Title:  []w.Inline{w.Text("Дежурство: первая линия")},
								Detail: "—",
								Badges: []w.Badge{w.Badge{Label: "repeats weekly"}},
								Source: "nextcloud · Работа",
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#8C6239"},
								Time:   "all day",
								Title:  []w.Inline{w.Text("Дежурство: первая линия")},
								Detail: "—",
								Badges: []w.Badge{w.Badge{Label: "repeats weekly"}},
								Source: "radicale · vps · Дежурства",
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#3F6E8C"},
								Time:   "15:00",
								Title:  []w.Inline{w.Text("Созвон с подрядчиком")},
								Detail: "Jitsi",
								Badges: []w.Badge{w.Badge{Label: "tentative", Tone: w.BadgeAlert}},
								Source: "nextcloud · Работа",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "Tuesday · 11 August", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:   w.Mark{Colour: "#6B4E7A"},
								Time:   "09:30",
								Title:  []w.Inline{w.Text("Ревью архитектуры")},
								Detail: "Переговорная 1",
								Badges: []w.Badge{w.Badge{Label: "2 attendees"}},
								Source: "radicale · vps · Проекты",
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Time:   "18:00",
								Title:  []w.Inline{w.Text("Забрать документы")},
								Detail: "МФЦ, Мясницкая 13",
								Source: "nextcloud · Личный",
							},
						},
					},
				},
			},
			Side: &w.Side{
				Items: []w.SidePart{
					w.Head{
						Title: &w.HeadTitle{Text: "Синк-ап по релизу", Size: w.HeadTitle19},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Личный")}}},
						Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "Edit", Small: true}, w.Close{Label: "Close the panel"}}},
					},
					w.Fields{
						Items: []w.FieldsItem{
							{Term: "When", Value: []w.ValuePart{w.Mono{Text: "10:00 – 11:00"}}},
							{Term: "Repeats", Value: []w.ValuePart{w.Text("Weekly on Mon")}},
							{Term: "Where", Value: []w.ValuePart{w.Text("Переговорная 2")}},
							{Term: "Status", Value: []w.ValuePart{w.Text("Confirmed")}},
						},
					},
					w.Sec{Items: []w.SecPart{
						w.Rubric{Text: "Source"},
						w.Diag{Text: "Read  2 min ago"}}},
				},
			},
		},
	}
}

// Dark2 is frame «Контакты и накладка перетаскивания · тёмная тема» (dark-2).
func Dark2() Screen {
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
						On: 1,
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.TickRow{
								Ticked: w.TickedPart,
								Kind:   w.TickRowAll,
								Mark:   w.Mark{Kind: w.MarkMulti},
								Label:  "Combined address book",
								Meta:   "7",
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "5", Current: true},
							w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "1"},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "radicale · vps"},
							w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "1", Trail: w.Badge{Label: "ro"}},
						},
					},
					w.RailFoot{
						Items: []w.RailFootPart{
							w.Button{Icon: "plus", Title: "New address book"},
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
							Title: &w.HeadTitle{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH20}, Text: "Контакты"},
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · 3 cards")}}},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "New contact", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
									w.Button{Label: "Import", Icon: "up"},
									w.Button{Label: "Export", Icon: "down"},
								},
							},
						},
						w.Bar{
							Items: []w.BarPart{
								w.Filter{
									Query:   w.Query{Placeholder: "Filter the loaded cards", Label: "Filter the loaded cards"},
									Waiting: true,
									Variant: w.FilterH26W220,
								},
								w.Seg{
									Options: []w.SegOption{
										{Label: "A–Z", Href: "/app/contacts?sort=az", On: true},
										{Label: "Recently changed", Href: "/app/contacts?sort=changed"},
									},
									Secondary: true,
								},
								w.Button{Label: "Columns", Icon: "cols"},
							},
							Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "3"}}, Secondary: true},
						},
						w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "АК"},
									Title:    []w.Inline{w.Text("Анна Ковалёва")},
									Subtitle: "+7 912 445-10-88",
									Detail:   "a.kovaleva@example.org",
									Caption:  "Издательство «Слово»",
									Source:   "nextcloud · Контакты",
								},
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "АС"},
									Title:    []w.Inline{w.Text("Алексей Смирнов")},
									Subtitle: "+7 916 200-71-15",
									Detail:   "—",
									Caption:  "—",
									Source:   "nextcloud · Контакты",
								},
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "В", Form: w.RubricInline}, Count: &w.Num{Value: 1}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "ВГ"},
									Title:    []w.Inline{w.Text("Виктор Гончаров")},
									Subtitle: "—",
									Detail:   "v.goncharov@example.org",
									Caption:  "Городская библиотека №4",
									Source:   "nextcloud · Контакты",
								},
							},
						},
					},
				},
			},
			Scrim: w.Scrim{
				Content: w.Empty{
					Icon: "up",
					Lead: "Drop to import",
					Parts: []w.EmptyPart{
						w.Aside{
							Face:  w.FaceLarge,
							Parts: []w.Inline{w.Text("One "), w.Mono{Text: ".vcf"}, w.Text(" file into "), w.Strong("nextcloud · Контакты")},
						},
						w.Dim{
							Parts:   []w.Inline{w.Text("The preview opens first: nothing is written until it has been read")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
						},
					},
					Variant: w.EmptyDrop,
				},
				Pad: w.ScrimPad38,
			},
		},
	}
}

// Dark3 is frame «Файлы плитками · тёмная тема» (dark-3).
func Dark3() Screen {
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
						Subs: []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("sftpgo · vps · 6 items")}}},
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
							w.Button{
								Folds: true,
								Label: "Download",
								Icon:  "down",
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
						Right: &w.BarRight{Items: []w.BarPart{w.Mono{Text: "6 items · 89.4 MB"}}, Secondary: true},
					},
					w.Grid{
						Tiles: []w.Tile{
							w.Tile{
								Thumb: w.Thumb{Icon: "file", Form: w.ThumbFolder},
								Name:  "архив",
								Meta:  "12 items",
								Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-07/архив",
								Send:  w.Send{Of: "files-sel", Name: "item", Value: "архив"},
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
								Meta:  "640 KB · 1920×1080",
								Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-07/схема-1.png",
								Send:  w.Send{Of: "files-sel", Name: "item", Value: "схема-1.png"},
							},
							w.Tile{
								Thumb: w.Thumb{Icon: "img"},
								Name:  "схема-2.png",
								Meta:  "712 KB · 1920×1080",
								Href:  "/app/files/sftpgo/documents/carrel/attachments/2026-07/схема-2.png",
								Send:  w.Send{Of: "files-sel", Name: "item", Value: "схема-2.png"},
							},
						},
					},
					w.PostForm{
						ID:     "files-sel",
						Action: "/app/files/sftpgo/documents/op/carrel/attachments",
						Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}},
					},
				},
			},
		},
	}
}
