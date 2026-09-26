package screens

import "carrel.dev/typecase/internal/legacy/widget"

// Tasks assembles mockup 7.3: combined list, overdue group, task card.
func Tasks() string {
	return shellFrame(shellNav("/tasks"), tasksRail(), tasksMain(), tasksSide(), "")
}

func tasksRail() widget.Rail {
	n30, n21, n9 := 30, 21, 9
	return widget.Rail{
		Sections: []widget.RailSection{
			{
				Rows: []widget.TickRow{{
					Label:   "Combined list",
					Mark:    widget.Mark{Colours: []string{"#4A6B52", "#3F6E8C"}, Label: "Combined"},
					Count:   intPtr(n30),
					Tick:    widget.TickPartial,
					Current: true,
				}},
			},
			{
				Title: "nextcloud · home",
				Rows: []widget.TickRow{
					{Label: "Дела", Mark: widget.Mark{Colour: "#4A6B52", Label: "Дела"}, Count: intPtr(n21), Tick: widget.TickOn},
					{Label: "Работа", Mark: widget.Mark{Colour: "#3F6E8C", Label: "Работа"}, Count: intPtr(n9), Tick: widget.TickOn},
				},
			},
			{
				Title: "Subscriptions",
				Rows: []widget.TickRow{
					{Label: "Праздники РФ", Mark: widget.Mark{Colour: "#8C6239", Label: "Праздники РФ"}, State: widget.TickRowReadOnly, Tick: widget.TickOff},
				},
			},
		},
		Chips: []widget.Chip{
			{Label: "all", On: true},
			{Label: "дом"},
			{Label: "релиз"},
			{Label: "read"},
		},
		Foot: widget.RailFoot{
			Actions: []widget.Action{
				{Label: "New list", Icon: "tc-i-plus"},
				{Label: "Connect account", Icon: "tc-i-server"},
				{Label: "Subscribe to a feed", Icon: "tc-i-link"},
			},
			Build: demoBuild,
		},
	}
}

func tasksMain() widget.Main {
	n3, n5, n4 := 3, 5, 4
	p40, p70, p100 := 40, 70, 100
	return widget.Main{
		Label: "Tasks",
		Body: []widget.Node{
			widget.Head{
				Mark:  widget.Mark{Colours: []string{"#4A6B52", "#3F6E8C"}, Label: "Combined"},
				Title: "Combined list",
				Sub:   "2 of 3 lists · nextcloud · home · 30 tasks",
				Actions: []widget.Action{
					{Label: "New task", Kind: widget.ButtonPrimary, Icon: "tc-i-plus"},
				},
				Rank2: []widget.Action{
					{Label: "Print", Icon: "tc-i-print"},
				},
			},
			widget.Bar{
				Left: []widget.Node{
					widget.Seg{
						Name:    "status",
						Chosen:  "open",
						Options: []widget.SegOption{{Label: "Open 12", Value: "open"}, {Label: "Done 48", Value: "done"}, {Label: "All 60", Value: "all"}},
					},
					widget.Seg{
						Name:    "sort",
						Chosen:  "due",
						Options: []widget.SegOption{{Label: "By due", Value: "due"}, {Label: "By priority", Value: "priority"}, {Label: "Recently changed", Value: "changed"}},
					},
					widget.Button{Label: "Columns", Icon: "tc-i-cols", Size: widget.ButtonSizeBar},
				},
				Right: []widget.Node{
					widget.Badge{Label: "3 overdue", Role: widget.BadgeAlert},
				},
			},
			widget.List{Groups: []widget.ListGroup{{
				Rubric: widget.Rubric{Text: "Overdue", Count: intPtr(n3)},
				Rows: []widget.Row{
					{
						Variant: widget.RowTickMarkCols,
						State:   widget.RowAlert,
						Mark:    widget.Mark{Colour: "#4A6B52", Label: "Дела"},
						Cells: widget.RowCells{
							Title:  "Отправить документы в МФЦ",
							Tags:   []widget.Tag{{Label: "дом"}},
							Due:    "7 Aug",
							Source: "nextcloud · Дела",
						},
					},
					{
						Variant: widget.RowTickMarkCols,
						State:   widget.RowAlert,
						Mark:    widget.Mark{Colour: "#3F6E8C", Label: "Работа"},
						Cells: widget.RowCells{
							Title:    "Ответить подрядчику по смете",
							Tags:     []widget.Tag{{Label: "релиз"}},
							Due:      "8 Aug",
							Progress: &p40,
							Source:   "nextcloud · Работа",
						},
					},
				},
			}}},
			widget.List{Groups: []widget.ListGroup{{
				Rubric: widget.Rubric{Text: "This week", Count: intPtr(n5)},
				Rows: []widget.Row{
					{
						Variant:  widget.RowTickMarkCols,
						Selected: true,
						Mark:     widget.Mark{Colour: "#4A6B52", Label: "Дела"},
						Cells: widget.RowCells{
							Title:  "Собрать чемодан",
							Tags:   []widget.Tag{{Label: "дом"}},
							Due:    "12 Aug 18:00",
							Source: "nextcloud · Дела",
						},
					},
					{
						Variant: widget.RowTickMarkCols,
						Mark:    widget.Mark{Colour: "#3F6E8C", Label: "Работа"},
						Cells: widget.RowCells{
							Title:    "Дописать раздел про кэш",
							Tags:     []widget.Tag{{Label: "релиз"}},
							Due:      "14 Aug",
							Progress: &p70,
							Source:   "nextcloud · Работа",
						},
					},
				},
			}}},
			widget.List{Groups: []widget.ListGroup{{
				Rubric: widget.Rubric{Text: "No due date", Count: intPtr(n4)},
				Rows: []widget.Row{
					{
						Variant: widget.RowTickMarkCols,
						Mark:    widget.Mark{Colour: "#4A6B52", Label: "Дела"},
						Cells: widget.RowCells{
							Title:    "Продлить страховку",
							Due:      "—",
							Done:     true,
							Progress: &p100,
							Source:   "nextcloud · Дела",
						},
					},
				},
			}}},
		},
	}
}

func tasksSide() widget.Side {
	fields := widget.Fields{Rows: []widget.FieldRow{
		{Label: "Status", Value: "Needs action"},
		{Label: "Due", Value: "2026-08-12 18:00"},
		{Label: "Priority", Value: "5"},
		{Label: "Complete", Value: "0%"},
		{Label: "Tags", Value: "дом"},
	}}
	return widget.Side{
		Label: "Task details",
		Title: "Собрать чемодан",
		Sub:   "nextcloud · Дела",
		Open:  true,
		Actions: []widget.Action{
			{Label: "Edit"},
		},
		Blocks: []widget.SideBlock{
			{Fields: &fields},
			{Sec: &widget.Sec{
				Rubric: widget.Rubric{Text: "Description"},
				Body:   []widget.Node{widget.Paragraph("Зарядки, переходник, распечатать посадочный.")},
			}},
			{Sec: &widget.Sec{
				Rubric: widget.Rubric{Text: "Related"},
				Body:   []widget.Node{widget.Paragraph("Поездка в Казань · event")},
			}},
			{Sec: &widget.Sec{Body: []widget.Node{widget.Danger{Title: "Danger zone", Rows: []widget.DangerRow{{
				Title:  "Delete task",
				Text:   "Removes the task from the server.",
				Action: widget.Action{Label: "Delete task", Kind: widget.ButtonDangerSolid},
			}}}}}},
		},
	}
}
