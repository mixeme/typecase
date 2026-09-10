package screens

import "carrel.dev/typecase/widget"

// Find assembles mockup 7.6: search hits as RowFind, not the search popover.
func Find() string {
	top := shellTop()
	top.Search.Query = "ковал"
	top.Search.Action = "/find"
	return shellFrameWith(top, shellNav(""), findRail(), findMain(), widget.Side{}, "")
}

func findRail() widget.Rail {
	n2, n1 := 2, 1
	return widget.Rail{
		Sections: []widget.RailSection{
			{
				Title: "Sources",
				Rows: []widget.TickRow{
					{Label: "Контакты", Mark: widget.Mark{Colour: "#4A6B52", Label: "Контакты"}, Count: intPtr(n2), Tick: widget.TickOn},
					{Label: "Личный", Mark: widget.Mark{Colour: "#4A6B52", Label: "Личный"}, Count: intPtr(n1), Tick: widget.TickOn},
					{Label: "Дела", Mark: widget.Mark{Colour: "#3F6E8C", Label: "Дела"}, Count: intPtr(n1), Tick: widget.TickOn},
					{Label: "Записи", Mark: widget.Mark{Colour: "#4A6B52", Label: "Записи"}, Count: intPtr(n1), Tick: widget.TickOn},
					{Label: "Архив", Mark: widget.Mark{Colour: "#A05A2C", Label: "Архив"}, Tick: widget.TickOff},
					{
						Label: "Общий",
						Mark:  widget.Mark{Colour: "#9A9280", Label: "Общий"},
						State: widget.TickRowUnavailable,
						Menu:  []widget.MenuItem{{Label: "Retry"}},
					},
					{Label: "Documents", Mark: widget.Mark{Colour: "#4A6B52", Label: "Documents"}, Count: intPtr(n2), Tick: widget.TickOn},
				},
			},
		},
		Foot: widget.RailFoot{Build: demoBuild},
	}
}

func findMain() widget.Main {
	n2, n1 := 2, 1
	return widget.Main{
		Label: "Find",
		Body: []widget.Node{
			widget.Head{
				Title: "«ковал»",
				Sub:   "6 of 7 collections · 7 records · 1 source timed out",
			},
			widget.Bar{
				Left: []widget.Node{
					widget.Seg{
						Name:   "kind",
						Chosen: "all",
						Options: []widget.SegOption{
							{Label: "All", Value: "all"},
							{Label: "Contacts", Value: "contacts"},
							{Label: "Events", Value: "events"},
							{Label: "Tasks", Value: "tasks"},
							{Label: "Notes", Value: "notes"},
							{Label: "Files", Value: "files"},
						},
					},
				},
			},
			widget.RunBar{
				Summary: "Done · 5 of 6 sources answered",
				At:      "answered 11:42",
				Actions: []widget.Action{{Label: "Search again", Icon: "tc-i-refresh"}},
			},
			widget.List{Groups: []widget.ListGroup{{
				Rubric: widget.Rubric{Text: "Contacts", Count: intPtr(n2)},
				Rows: []widget.Row{
					{
						Variant: widget.RowFind,
						Mark:    widget.Mark{Colour: "#4A6B52", Label: "Контакты"},
						Cells: widget.RowCells{
							Lead:    "contact",
							Title:   "Ковалёва Анна",
							Detail:  "+7 912 445-10-88",
							Caption: "Издательство «Слово»",
							Source:  "nextcloud · Контакты",
						},
					},
					{
						Variant: widget.RowFind,
						Mark:    widget.Mark{Colour: "#6B4E7A", Label: "Команда"},
						Cells: widget.RowCells{
							Lead:    "contact",
							Title:   "Ковалёв Пётр",
							Detail:  "p.kovalev@example.org",
							Caption: "—",
							Source:  "radicale · Команда",
						},
					},
				},
			}}},
			widget.List{Groups: []widget.ListGroup{{
				Rubric: widget.Rubric{Text: "Events", Count: intPtr(n1)},
				Rows: []widget.Row{
					{
						Variant: widget.RowFind,
						Mark:    widget.Mark{Colour: "#4A6B52", Label: "Личный"},
						Cells: widget.RowCells{
							Lead:    "event",
							Title:   "Звонок с Ковалёвой",
							Detail:  "Zoom",
							Caption: "15:00",
							Source:  "nextcloud · Личный",
						},
					},
				},
			}}},
			widget.List{Groups: []widget.ListGroup{{
				Rubric: widget.Rubric{Text: "Tasks", Count: intPtr(n1)},
				Rows: []widget.Row{
					{
						Variant: widget.RowFind,
						Mark:    widget.Mark{Colour: "#3F6E8C", Label: "Работа"},
						Cells: widget.RowCells{
							Lead:    "task",
							Title:   "Ответить Ковалёвой по смете",
							Detail:  "mentioned in the title",
							Caption: "12 Aug",
							Source:  "nextcloud · Работа",
						},
					},
				},
			}}},
			widget.List{Groups: []widget.ListGroup{{
				Rubric: widget.Rubric{Text: "Notes", Count: intPtr(n1)},
				Rows: []widget.Row{
					{
						Variant: widget.RowFind,
						Mark:    widget.Mark{Colour: "#4A6B52", Label: "Записи"},
						Cells: widget.RowCells{
							Lead:    "note",
							Title:   "Звонок с Ковалёвой — тезисы",
							Detail:  "mentioned in the text",
							Caption: "10 Aug",
							Source:  "nextcloud · Записи",
						},
					},
				},
			}}},
			widget.List{Groups: []widget.ListGroup{{
				Rubric: widget.Rubric{Text: "Files", Count: intPtr(n2)},
				Rows: []widget.Row{
					{
						Variant: widget.RowFind,
						Mark:    widget.Mark{Colour: "#4A6B52", Label: "Documents"},
						Cells: widget.RowCells{
							Lead:    "file",
							Title:   "договор-ковалёва.pdf",
							Detail:  "240 KB",
							Caption: "2026-08-04 11:20",
							Source:  "sftpgo · Documents",
						},
					},
					{
						Variant: widget.RowFind,
						Mark:    widget.Mark{Colour: "#4A6B52", Label: "Documents"},
						Cells: widget.RowCells{
							Lead:    "file",
							Title:   "Ковалёва — правки по обложке.md",
							Detail:  "6 KB",
							Caption: "2026-07-21 16:02",
							Source:  "sftpgo · Documents",
						},
					},
				},
			}}},
			widget.Msg{
				Text: "Files are searched by name. Carrel does not look inside files — and says so here, so that “nothing found” is not read as “not in the files”.",
			},
		},
	}
}
