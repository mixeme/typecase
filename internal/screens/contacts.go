// Раздел макета: 7.1 Контакты и человек целиком.

package screens

import w "carrel.dev/typecase/widget"

// Contacts1 is frame «Контакты · список и карточка» (s-contacts-1).
func Contacts1() Screen {
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
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"},
					},
					On: 1,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedPart,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined address book",
							Meta:    "312",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "248"},
						w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "51"},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "radicale · vps"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Команда", Meta: "13", Trail: w.Badge{Label: "ro"}},
						w.TickRow{Ticked: w.TickedOff, Mark: w.Mark{Colour: "#8C6239"}, Label: "Рассылки", Meta: "—", Hidden: true},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Tags"}},
					Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "работа"}, w.Tag{Label: "издательство"}},
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
						Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined address book"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 312 cards")}}},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "New contact", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
								w.Button{Label: "Import", Icon: "up"},
								w.Button{Label: "Export", Icon: "down"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
					},
					w.Bar{
						Items: []w.BarPart{
							w.Filter{
								Query:   w.Query{Value: "Filter the loaded cards", Label: "Filter the loaded cards"},
								Waiting: true,
								Variant: w.FilterH26W220,
							},
							w.Seg{Options: []w.SegOption{{Label: "A–Z", On: true}, {Label: "Recently changed"}}, Secondary: true},
							w.Button{Label: "Columns", Icon: "cols"},
						},
						More: &w.Button{Icon: "dots", Title: "More"},
						Right: &w.BarRight{
							Items:     []w.BarPart{w.Check{Parts: []w.CheckPart{w.Text("Compact rows")}}, w.Mono{Text: "312"}},
							Secondary: true,
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
					w.List{
						Rows: []w.ListRow{
							w.FaceRow{
								Mark:     w.Mark{Colour: "#4A6B52"},
								Avatar:   w.Avatar{Initials: "АК"},
								Title:    []w.Inline{w.Text("Анна Ковалёва")},
								Strong:   true,
								Subtitle: "+7 912 445-10-88",
								Detail:   "a.kovaleva@example.org",
								Caption:  "Издательство «Слово»",
								Source:   "nextcloud · Контакты",
								On:       true,
							},
							w.FaceRow{
								Mark:       w.Mark{Colour: "#6B4E7A"},
								Avatar:     w.Avatar{Photo: true},
								Title:      []w.Inline{w.Text("Артём Белов")},
								TitleBadge: &w.Badge{Label: "linked ×2", Tone: w.BadgeLinked},
								Subtitle:   "+7 903 118-22-04",
								Detail:     "a.belov@example.org",
								Caption:    "Тензор",
								Source:     "radicale · vps · Команда",
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
					w.ListGroup{Title: w.Rubric{Text: "В", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
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
							w.FaceRow{
								Mark:     w.Mark{Colour: "#A05A2C"},
								Avatar:   w.Avatar{Initials: "ВЛ"},
								Title:    []w.Inline{w.Text("Вера Лебедева")},
								Subtitle: "+7 905 331-09-42",
								Detail:   "vera@example.net",
								Caption:  "—",
								Source:   "nextcloud · Архив",
							},
						},
					},
					w.Empty{Parts: []w.EmptyPart{w.Text("Loading more…")}, Variant: w.EmptyDusty},
				},
			},
			Side: &w.Side{
				Items: []w.SidePart{
					w.Head{
						Photo: &w.HeadPhoto{Initials: "АК"},
						Title: &w.HeadTitle{Text: "Анна Ковалёва", Size: w.HeadTitle19},
						Tools: []w.Button{
							w.Button{Label: "Edit", Small: true},
							w.Button{Label: "All records", Small: true},
							w.Button{Icon: "dots", Small: true, Title: "More"},
						},
						Subs:   []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}}},
						Acts:   &w.Acts{Items: []w.ActsPart{w.Close{Label: "Close the panel"}}},
						Middle: true,
					},
					w.Fields{
						Items: []w.FieldsItem{
							{Term: "Organisation", Value: []w.ValuePart{w.Text("Издательство «Слово»")}},
							{Term: "Title", Value: []w.ValuePart{w.Text("Выпускающий редактор")}},
							{Term: "Phone · mobile", Value: []w.ValuePart{w.Text("+7 912 445-10-88")}, Face: w.ValueMono},
							{Term: "Phone · work", Value: []w.ValuePart{w.Text("+7 495 120-45-00")}, Face: w.ValueMono},
							{Term: "Email · work", Value: []w.ValuePart{w.Text("a.kovaleva@example.org")}},
							{Term: "Birthday", Value: []w.ValuePart{w.Text("1987-03-14")}, Face: w.ValueMono},
							{Term: "Categories", Value: []w.ValuePart{w.Tag{Label: "работа"}, w.Tag{Label: "издательство"}}},
							{Term: "Note", Value: []w.ValuePart{w.Text("Пишет по вторникам, отвечает в тот же день.")}},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Kept as the server sent them", Items: []w.RubricPart{w.Button{Label: "Show 3", Small: true}}},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Source"},
							w.Diag{Text: "UID   4f2b-9c31-a0e7\nRead  2 min ago"}},
					},
					w.Danger{Button: w.Button{Label: "Delete contact", Kind: w.ButtonDangerSolid}},
				},
			},
		},
	}
}

// Contacts2 is frame «Контакт · переместить в другую книгу» (s-contacts-2).
func Contacts2() Screen {
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
		Body: w.Overlay{
			Body: w.OverlayBody{
				Rail: []w.OverlayRailPart{
					w.Nav{
						Items: []w.NavItem{
							{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
							{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
							{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
							{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
							{Icon: "file", Label: "Files", Href: "#"},
							{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"},
						},
						On: 1,
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.TickRow{
								Ticked:  w.TickedPart,
								Kind:    w.TickRowAll,
								Mark:    w.Mark{Kind: w.MarkMulti},
								Label:   "Combined address book",
								Meta:    "312",
								Current: true,
							},
						},
					},
					w.RailSec{
						Items: []w.RailSecPart{
							w.Rubric{Text: "nextcloud"},
							w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "248"},
							w.TickRow{Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "51"},
						},
					},
				},
				Main: w.Main{
					Items: []w.MainPart{
						w.Head{
							Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined address book"},
							Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("3 of 4 books · 2 accounts · 312 cards")}}},
							Acts: &w.Acts{
								Items: []w.ActsPart{
									w.Button{Label: "New contact", Menu: true, Icon: "plus", Kind: w.ButtonPrimary},
									w.Button{Label: "Import", Icon: "up"},
									w.Button{Label: "Export", Icon: "down"},
									w.Button{Label: "Print", Icon: "print"},
								},
							},
						},
						w.Bar{
							Items: []w.BarPart{
								w.Filter{
									Query:   w.Query{Value: "Filter the loaded cards", Label: "Filter the loaded cards"},
									Waiting: true,
									Variant: w.FilterH26W220,
								},
								w.Seg{Options: []w.SegOption{{Label: "A–Z", On: true}, {Label: "Recently changed"}}, Secondary: true},
								w.Button{Label: "Columns", Icon: "cols"},
							},
							More: &w.Button{Icon: "dots", Title: "More"},
							Right: &w.BarRight{
								Items:     []w.BarPart{w.Check{Parts: []w.CheckPart{w.Text("Compact rows")}}, w.Mono{Text: "312"}},
								Secondary: true,
							},
						},
						w.ListGroup{Title: w.Rubric{Text: "А", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
						w.List{
							Rows: []w.ListRow{
								w.FaceRow{
									Mark:     w.Mark{Colour: "#4A6B52"},
									Avatar:   w.Avatar{Initials: "АК"},
									Title:    []w.Inline{w.Text("Анна Ковалёва")},
									Strong:   true,
									Subtitle: "+7 912 445-10-88",
									Detail:   "a.kovaleva@example.org",
									Caption:  "Издательство «Слово»",
									Source:   "nextcloud · Контакты",
									On:       true,
								},
								w.FaceRow{
									Mark:       w.Mark{Colour: "#6B4E7A"},
									Avatar:     w.Avatar{Photo: true},
									Title:      []w.Inline{w.Text("Артём Белов")},
									TitleBadge: &w.Badge{Label: "linked ×2", Tone: w.BadgeLinked},
									Subtitle:   "+7 903 118-22-04",
									Detail:     "a.belov@example.org",
									Caption:    "Тензор",
									Source:     "radicale · vps · Команда",
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
						w.ListGroup{Title: w.Rubric{Text: "В", Form: w.RubricInline}, Count: &w.Num{Value: 2}},
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
								w.FaceRow{
									Mark:     w.Mark{Colour: "#A05A2C"},
									Avatar:   w.Avatar{Initials: "ВЛ"},
									Title:    []w.Inline{w.Text("Вера Лебедева")},
									Subtitle: "+7 905 331-09-42",
									Detail:   "vera@example.net",
									Caption:  "—",
									Source:   "nextcloud · Архив",
								},
							},
						},
						w.Empty{Parts: []w.EmptyPart{w.Text("Loading more…")}, Variant: w.EmptyDusty},
					},
				},
				Height: w.ShellHeight560,
			},
			Scrim: w.Scrim{
				Content: w.Sheet{
					Icon:  "move",
					Title: "Move contact",
					Head: []w.SheetHeadPart{
						w.Dim{
							Parts:   []w.Inline{w.Text("Анна Ковалёва · from nextcloud · Контакты")},
							Size:    w.Dim11_5,
							Variant: w.DimWrap,
							Place:   w.Pushed,
						},
					},
					Body: []w.SheetPart{
						w.Rubric{Text: "Write into"},
						w.Tree{
							Rows: []w.TreeRow{
								{
									Mark:        &w.Mark{Colour: "#4A6B52"},
									Label:       "Контакты",
									Unavailable: true,
									Trail:       w.Mono{Text: "already there", Variant: w.Mono10Dusty},
								},
								{
									Mark:  &w.Mark{Colour: "#A05A2C"},
									Label: "Архив",
									Trail: w.Mono{Text: "nextcloud", Variant: w.Mono10Dusty},
								},
								{
									Mark:  &w.Mark{Colour: "#6B4E7A"},
									Label: "Команда",
									Trail: w.Mono{Text: "radicale", Variant: w.Mono10Dusty},
								},
								{
									Mark:        &w.Mark{Colour: "#8C6239"},
									Label:       "Shared",
									Unavailable: true,
									Trail:       w.Badge{Label: "read-only"},
								},
							},
							On: 2,
						},
						w.Msg{
							Parts: []w.MsgPart{
								w.Text("Inside one server this is "),
								w.Mono{Text: "MOVE", Form: w.MonoCode},
								w.Text(": instant, no traffic. To another address book the card is copied and then deleted — Carrel says so before it starts, not after. Copy is this same sheet with "),
								w.Strong("Copy here"),
								w.Text("."),
							},
							Space: w.MsgFlushTop2,
						},
					},
					Foot: []w.SheetFootPart{
						w.Button{Label: "Move here", Kind: w.ButtonPrimary},
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

// Contacts3 is frame «Контакт · правка, фотография» (s-contacts-3).
func Contacts3() Screen {
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
					{Icon: "back", Label: "Back to the card", Href: "#"}}},
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
					w.Head{
						Title: &w.HeadTitle{Text: "Анна Ковалёва"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}}},
						Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Save", Kind: w.ButtonPrimary}}},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label:     "Photo",
								LabelSpan: true,
								Parts: []w.FieldPart{
									w.Photo{
										Side: []w.FieldPart{
											w.Line{Items: []w.LinePart{w.Button{Label: "Replace"}, w.Button{Label: "Crop"}}, Gap: w.LineGap6},
											w.Button{Label: "Remove", Kind: w.ButtonDanger},
											w.FieldHint{Parts: []w.Inline{w.Text("EXIF orientation applied, metadata stripped.")}, Variant: w.HintNarrow},
										},
									},
								},
								Width: w.FieldFixAuto,
							},
							w.FormCols{
								Items: []w.Field{
									w.Field{
										Label: "Display name",
										Parts: []w.FieldPart{w.Input{Value: "Анна Ковалёва", Label: "Display name", Focus: true}},
										Width: w.FieldXL,
									},
									w.Field{Label: "Given", Parts: []w.FieldPart{w.Input{Value: "Анна", Label: "Given"}}, Width: w.FieldSM},
									w.Field{Label: "Family", Parts: []w.FieldPart{w.Input{Value: "Ковалёва", Label: "Family"}}, Width: w.FieldSM},
									w.Field{
										Label: "Nickname",
										Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Nickname"}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Organisation",
										Parts: []w.FieldPart{w.Input{Value: "Издательство «Слово»", Label: "Organisation"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Title",
										Parts: []w.FieldPart{w.Input{Value: "Выпускающий редактор", Label: "Title"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Birthday",
										Parts: []w.FieldPart{
											w.Input{Value: "14.03.1987", Label: "Birthday", Mono: true, Numeric: true},
											w.FieldHint{Parts: []w.Inline{w.Text("Type it: 14.03.1987, or 14.03 without the year.")}},
										},
										Width: w.FieldSM,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Phone numbers"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Label", Parts: []w.FieldPart{w.Pick{Value: "mobile"}}, Width: w.FieldXS},
									w.Field{
										Label: "Number",
										Parts: []w.FieldPart{w.Input{Value: "+7 912 445-10-88", Label: "Number", Mono: true}},
										Width: w.FieldMD,
									},
									w.Button{Label: "Remove", Kind: w.ButtonDanger},
									w.Grow{},
									w.Field{Parts: []w.FieldPart{w.Input{Value: "work"}}, Width: w.FieldXS},
									w.Field{Parts: []w.FieldPart{w.Input{Value: "+7 495 120-45-00", Mono: true}}, Width: w.FieldMD},
									w.Button{Label: "Remove", Kind: w.ButtonDanger},
									w.Grow{},
									w.Button{Label: "Add phone", Icon: "plus", Small: true},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Email addresses"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Label", Parts: []w.FieldPart{w.Pick{Value: "work"}}, Width: w.FieldXS},
									w.Field{
										Label: "Address",
										Parts: []w.FieldPart{w.Input{Value: "a.kovaleva@example.org", Label: "Address"}},
										Width: w.FieldLG,
									},
									w.Button{Label: "Remove", Kind: w.ButtonDanger},
									w.Grow{},
									w.Button{Label: "Add email", Icon: "plus", Small: true},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Other"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Categories",
										Parts: []w.FieldPart{w.Input{Value: "работа, издательство", Label: "Categories"}},
										Width: w.FieldLG,
									},
									w.Field{
										Label: "Note",
										Parts: []w.FieldPart{w.Area{Text: "Пишет по вторникам, отвечает в тот же день."}},
										Width: w.FieldFull,
									},
									w.Field{Parts: []w.FieldPart{w.Button{Label: "Show 6 empty fields"}}},
								},
							},
						},
					},
					w.FormFoot{
						Items: []w.FormFootPart{
							w.Dim{Parts: []w.Inline{w.Text("Properties Carrel does not know are kept unchanged.")}, Size: w.Dim11_5},
						},
					},
				},
			},
		},
	}
}

// Contacts4 is frame «Контакт · правка, нет фотографии» (s-contacts-4).
func Contacts4() Screen {
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
					{Icon: "back", Label: "Back to the card", Href: "#"}}},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Контакты", Meta: "248", Current: true},
					},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Анна Ковалёва"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}}},
						Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "Cancel"}, w.Button{Label: "Save", Kind: w.ButtonPrimary}}},
					},
					w.Form{
						Items: []w.FormPart{
							w.Field{
								Label:     "Photo",
								LabelSpan: true,
								Parts: []w.FieldPart{
									w.Photo{
										Initials: "АК",
										Side: []w.FieldPart{
											w.Button{Label: "Add"},
											w.FieldHint{Parts: []w.Inline{w.Text("JPEG, PNG, GIF or WebP.")}, Variant: w.HintNarrow},
										},
									},
								},
								Width: w.FieldFixAuto,
							},
							w.FormCols{
								Items: []w.Field{
									w.Field{
										Label: "Display name",
										Parts: []w.FieldPart{w.Input{Value: "Анна Ковалёва", Label: "Display name", Focus: true}},
										Width: w.FieldXL,
									},
									w.Field{Label: "Given", Parts: []w.FieldPart{w.Input{Value: "Анна", Label: "Given"}}, Width: w.FieldSM},
									w.Field{Label: "Family", Parts: []w.FieldPart{w.Input{Value: "Ковалёва", Label: "Family"}}, Width: w.FieldSM},
									w.Field{
										Label: "Nickname",
										Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Nickname"}},
										Width: w.FieldSM,
									},
									w.Field{
										Label: "Organisation",
										Parts: []w.FieldPart{w.Input{Value: "Издательство «Слово»", Label: "Organisation"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Title",
										Parts: []w.FieldPart{w.Input{Value: "Выпускающий редактор", Label: "Title"}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Birthday",
										Parts: []w.FieldPart{
											w.Input{Placeholder: "DD.MM.YYYY", Label: "Birthday", Mono: true, Numeric: true},
											w.FieldHint{Parts: []w.Inline{w.Text("Type it: 14.03.1987, or 14.03 without the year.")}},
										},
										Width: w.FieldSM,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Phone numbers"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Label", Parts: []w.FieldPart{w.Pick{Value: "mobile"}}, Width: w.FieldXS},
									w.Field{
										Label: "Number",
										Parts: []w.FieldPart{w.Input{Value: "+7 912 445-10-88", Label: "Number", Mono: true}},
										Width: w.FieldMD,
									},
									w.Button{Label: "Remove", Kind: w.ButtonDanger},
									w.Grow{},
									w.Button{Label: "Add phone", Icon: "plus", Small: true},
								},
							},
						},
					},
					w.FormFoot{
						Items: []w.FormFootPart{
							w.Dim{Parts: []w.Inline{w.Text("Properties Carrel does not know are kept unchanged.")}, Size: w.Dim11_5},
						},
					},
				},
			},
		},
	}
}

// Contacts5 is frame «Фотография контакта · кадрирование» (s-contacts-5).
func Contacts5() Screen {
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
					{Icon: "back", Label: "Back to the card", Href: "#"}}},
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
								Title: &w.HeadTitle{Text: "Photo"},
								Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Анна Ковалёва · metadata will be stripped")}}},
							},
							w.PhotoCrop{
								Colour: "#8C7B6A",
								Note:   "Drag the photo to choose the frame. The box is what other clients will see.",
								Fields: []w.Field{
									w.Field{Label: "Zoom", Parts: []w.FieldPart{w.Zoom{Value: "1.6×"}}, Width: w.FieldW100},
									w.Field{
										Label: "Rotate",
										Parts: []w.FieldPart{
											w.Seg{
												Options: []w.SegOption{{Label: "0°", On: true}, {Label: "90°"}, {Label: "180°"}, {Label: "270°"}},
												Variant: w.SegStart,
											},
										},
										Width: w.FieldW100,
									},
									w.Field{
										Label: "Saved as",
										Parts: []w.FieldPart{
											w.Input{Value: "512 × 512 · JPEG · about 90 KB", Label: "Saved as", Mono: true},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("The photo lives inside the card itself, so its size is the card’s size: a square of 512 px at quality 85 keeps a vCard near 200 KB, which servers and phones still accept. Lists are served a smaller copy of the same crop — the original file is not kept anywhere."),
												},
											},
										},
										Width: w.FieldW100,
									},
								},
								After: &w.Dim{
									Parts: []w.Inline{w.Text("EXIF orientation is already applied, metadata is stripped.")},
									Size:  w.Dim11,
								},
							},
							w.FormFoot{
								Items: []w.FormFootPart{
									w.Button{Label: "Save photo", Kind: w.ButtonPrimary},
									w.Button{Label: "Cancel"},
									w.Push{Items: []w.PushPart{w.Button{Label: "Remove photo", Kind: w.ButtonDanger}}},
								},
								Variant: w.FormFootFlush,
							},
						},
						Variant: w.Dialog520,
					},
				},
			},
		},
	}
}

// Contacts6 is frame «Контакт · карточка, чтение» (s-contacts-6).
func Contacts6() Screen {
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
					w.Head{
						Photo: &w.HeadPhoto{},
						Title: &w.HeadTitle{Text: "Анна Ковалёва"},
						Subs: []w.HeadSub{
							w.HeadSub{Parts: []w.HeadSubPart{w.Text("Издательство «Слово» · Выпускающий редактор")}},
							w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Контакты")}},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Edit", Icon: "pencil", Kind: w.ButtonPrimary},
								w.Button{Label: "Send", Icon: "mail"},
								w.Button{Label: "Print", Icon: "print"},
								w.Button{Icon: "dots", Title: "More"},
							},
						},
					},
					w.Fields{
						Items: []w.FieldsItem{
							{Term: "Phone · mobile", Value: []w.ValuePart{w.Text("+7 912 445-10-88")}, Face: w.ValueMono},
							{Term: "Phone · work", Value: []w.ValuePart{w.Text("+7 495 120-45-00")}, Face: w.ValueMono},
							{Term: "Email · work", Value: []w.ValuePart{w.Text("a.kovaleva@example.org")}},
							{Term: "Email · second", Value: []w.ValuePart{w.Text("anna.kovaleva@slovo.example")}},
							{Term: "Address · work", Value: []w.ValuePart{w.Text("Москва, Тверская 12, стр. 3")}},
							{
								Term:  "Birthday",
								Value: []w.ValuePart{w.Text("1987-03-14 "), w.Dim{Parts: []w.Inline{w.Text("· in 7 months")}}},
								Face:  w.ValueMono,
							},
							{Term: "Categories", Value: []w.ValuePart{w.Tag{Label: "работа"}, w.Tag{Label: "издательство"}}},
							{Term: "Note", Value: []w.ValuePart{w.Text("Пишет по вторникам, отвечает в тот же день.")}},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Kept as the server sent them", Items: []w.RubricPart{w.Button{Label: "Show 3", Small: true}}},
							w.SecHint{Text: "Properties Carrel does not edit are stored byte for byte and go back untouched."},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{
								Text:  "The same person in other books",
								Items: []w.RubricPart{w.Badge{Label: "2", Tone: w.BadgeLinked}},
							},
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{
										w.Mark{Colour: "#6B4E7A", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "radicale · vps · Команда", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· Phone, email differ")}},
									},
									[]w.SecItemPart{
										w.Mark{Colour: "#9A9280", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "baikal · Общий", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· The server did not answer; the record stays linked")}},
									},
								},
								Gap: 5,
							},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Timeline", Items: []w.RubricPart{w.Mono{Text: "11", Variant: w.MonoTight}}},
							w.List{
								Rows: []w.ListRow{
									w.LinkedRow{
										Mark:   w.Mark{Colour: "#4A6B52"},
										Icon:   "cal",
										Title:  "Синк-ап по релизу",
										Detail: "10 Aug · участник · приняла",
									},
									w.LinkedRow{
										Mark:   w.Mark{Colour: "#4A6B52"},
										Icon:   "note",
										Title:  "Заметка со встречи 10.08",
										Detail: "10 Aug · упомянута в тексте",
									},
									w.LinkedRow{
										Mark:   w.Mark{Colour: "#3F6E8C"},
										Icon:   "task",
										Title:  "Отправить макет Ковалёвой",
										Detail: "08 Aug · упомянута в названии",
									},
								},
							},
							w.SecActs{Items: []w.Button{w.Button{Label: "Open the timeline", Icon: "all"}}},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Source"},
							w.Diag{Text: "Path  /dav.php/addressbooks/mix/contacts/2b7e.vcf\nRead  2 min ago · not changed on server"},
						},
					},
					w.Danger{
						Button: w.Button{Label: "Delete contact", Icon: "trash", Kind: w.ButtonDangerSolid},
						Note:   "Удаление уходит на сервер сразу и отменить его нечем.",
					},
				},
			},
		},
	}
}

// Contacts7 is frame «Человек · хронология» (s-contacts-7).
func Contacts7() Screen {
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
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
						{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"},
					},
					On: 1,
				},
				w.Nav{Items: []w.NavItem{
					{Icon: "back", Label: "Back to list", Href: "#"}}, Apart: true},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "Where to look"},
						w.TickRow{
							Ticked:  w.TickedPart,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined calendar",
							Meta:    "5",
							Current: true,
						},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "9"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "7"},
						w.TickRow{Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "4"},
						w.TickRow{
							Mark:      w.Mark{Colour: "#9A9280"},
							Label:     "Общий",
							Meta:      "password · 401",
							MetaTitle: "401 Unauthorized — the saved password was refused. Open Connections to update it.",
							Error:     true,
							Trail:     w.Button{Label: "Retry", Small: true},
						},
					},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Avatar: &w.Avatar{Initials: "АК", Size: w.Avatar44},
						Kicker: &w.Rubric{Text: "Timeline", Variant: w.RubricBottom2},
						Title:  &w.HeadTitle{Text: "Анна Ковалёва"},
						Subs: []w.HeadSub{
							w.HeadSub{
								Parts: []w.HeadSubPart{
									w.Text("Издательство «Слово» · Выпускающий редактор "),
									w.Badge{Label: "in 2 books", Tone: w.BadgeLinked},
								},
							},
						},
						Acts: &w.Acts{
							Items: []w.ActsPart{
								w.Button{Label: "Write a note", Icon: "quill", Kind: w.ButtonPrimary},
								w.Button{Label: "New event", Icon: "cal"},
								w.Button{Label: "New task", Icon: "task"},
								w.Button{Label: "Print", Icon: "print"},
							},
						},
						Pad: w.HeadPadBottom10,
					},
					w.Bar{
						Items: []w.BarPart{
							w.Range{
								From: w.Input{Kind: w.InputDate, Value: "2026-01-01", Mono: true},
								To:   w.Input{Value: "today", Mono: true},
								Show: w.Button{Label: "Show"},
							},
							w.Seg{
								Options: []w.SegOption{
									{Label: "All 24", On: true},
									{Label: "Events 14"},
									{Label: "Tasks 2"},
									{Label: "Notes 4"},
									{Label: "Files 4"},
								},
								Secondary: true,
							},
							w.Button{Label: "Columns", Icon: "cols", Secondary: true},
						},
						More:    &w.Button{Icon: "dots", Title: "More"},
						Right:   &w.BarRight{Items: []w.BarPart{w.Mono{Text: "newest first"}}, Secondary: true},
						Variant: w.BarWide,
					},
					w.Runbar{
						Meter: w.Meter{Percent: 80},
						Parts: []w.RunbarPart{
							w.Text("Polled 4 of 5 calendars · matched by "),
							w.Mono{Text: "a.kovaleva@example.org"},
							w.Text(", "),
							w.Mono{Text: "anna.kovaleva@slovo.example"},
							w.Text(" and the name · "),
							w.Span{Face: w.FaceDanger, Text: "baikal · Общий unavailable"},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "August 2026", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
					w.List{
						Rows: []w.ListRow{
							w.FindRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Date:   "10 Aug",
								Title:  []w.Inline{w.Text("Синк-ап по релизу")},
								Detail: []w.Inline{w.Text("attendee · accepted")},
								What:   "event",
								Source: "nextcloud · Личный",
							},
							w.FindRow{
								Mark:   w.Mark{Colour: "#4A6B52"},
								Date:   "4 Aug",
								Title:  []w.Inline{w.Text("2026-08-04-обложка.pdf")},
								Detail: []w.Inline{w.Text("attached to «Согласовать обложку»")},
								What:   "file",
								Source: "sftpgo · vps · Documents",
							},
							w.FindRow{
								Mark:   w.Mark{Colour: "#3F6E8C"},
								Date:   "4 Aug",
								Title:  []w.Inline{w.Text("Согласовать обложку")},
								Detail: []w.Inline{w.Text("organiser · done")},
								What:   "task",
								Source: "nextcloud · Работа",
							},
						},
					},
					w.ListGroup{Title: w.Rubric{Text: "July 2026", Form: w.RubricInline}, Count: &w.Num{Value: 5}},
					w.List{
						Rows: []w.ListRow{
							w.FindRow{
								Mark:   w.Mark{Colour: "#6B4E7A"},
								Date:   "29 Jul",
								Title:  []w.Inline{w.Text("Ревью вёрстки сборника")},
								Detail: []w.Inline{w.Text("attendee · no reply")},
								What:   "event",
								Source: "radicale · vps · Проекты",
							},
						},
					},
					w.Empty{
						Parts:   []w.EmptyPart{w.Text("Nothing was looked for before January 2026. Widen the range if you need it.")},
						Variant: w.EmptySmallTop,
					},
				},
			},
			Side: &w.Side{
				Items: []w.SidePart{
					w.Head{
						Photo:  &w.HeadPhoto{Initials: "АК", Large: true},
						Kicker: &w.Rubric{Text: "Card"},
						Tools: []w.Button{
							w.Button{Label: "Open the card", Small: true},
							w.Button{Label: "Edit", Small: true},
							w.Button{Icon: "dots", Small: true, Title: "More"},
						},
						Acts:   &w.Acts{Items: []w.ActsPart{w.Close{Label: "Close the panel"}}},
						Middle: true,
						Pad:    w.HeadPadBottom10,
					},
					w.Fields{
						Items: []w.FieldsItem{
							{Term: "Phone · mobile", Value: []w.ValuePart{w.Text("+7 912 445-10-88")}, Face: w.ValueMono},
							{Term: "Phone · work", Value: []w.ValuePart{w.Text("+7 495 120-45-00")}, Face: w.ValueMono},
							{Term: "Email · work", Value: []w.ValuePart{w.Text("a.kovaleva@example.org")}},
							{Term: "Email · second", Value: []w.ValuePart{w.Text("anna.kovaleva@slovo.example")}},
							{
								Term:  "Birthday",
								Value: []w.ValuePart{w.Text("1987-03-14 "), w.Dim{Parts: []w.Inline{w.Text("· in 7 months")}}},
								Face:  w.ValueMono,
							},
							{Term: "Categories", Value: []w.ValuePart{w.Tag{Label: "работа"}, w.Tag{Label: "издательство"}}},
							{Term: "Note", Value: []w.ValuePart{w.Text("Пишет по вторникам, отвечает в тот же день.")}},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{
								Text:  "The same person in other books",
								Items: []w.RubricPart{w.Badge{Label: "2", Tone: w.BadgeLinked}},
							},
							w.SecList{
								Items: [][]w.SecItemPart{
									[]w.SecItemPart{
										w.Mark{Colour: "#6B4E7A", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "radicale · vps · Команда", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· Phone, email differ")}},
									},
									[]w.SecItemPart{
										w.Mark{Colour: "#9A9280", Size: w.MarkH11Inline},
										w.Text(" "),
										w.Link{Text: "baikal · Общий", Href: "#"},
										w.Text(" "),
										w.Dim{Parts: []w.Inline{w.Text("· The server did not answer; the record stays linked")}},
									},
								},
								Gap: 6,
							},
							w.SecHint{
								Text:  "Linked on the duplicates screen. The records stay on their own servers; here they are shown as one person.",
								Apart: true,
							},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Source"},
							w.Diag{Text: "UID   4f2b-9c31-a0e7\nRead  2 min ago"}},
					},
					w.Danger{Button: w.Button{Label: "Delete contact", Kind: w.ButtonDangerSolid}},
				},
			},
		},
	}
}
