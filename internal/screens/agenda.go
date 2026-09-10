package screens

import "carrel.dev/typecase/widget"

// Agenda assembles mockup 7.2: combined calendar, week list, event card.
func Agenda() string {
	return shellFrame(shellNav("/agenda"), agendaRail(), agendaMain(), agendaSide(), "")
}

func agendaRail() widget.Rail {
	n14, n7, n4, n2, n1 := 14, 7, 4, 2, 1
	return widget.Rail{
		Sections: []widget.RailSection{
			{
				Rows: []widget.TickRow{{
					Label:   "Combined calendar",
					Mark:    widget.Mark{Colours: []string{"#4A6B52", "#3F6E8C", "#6B4E7A"}, Label: "Combined"},
					Count:   intPtr(n14),
					Tick:    widget.TickPartial,
					Current: true,
				}},
			},
			{
				Title: "nextcloud · home",
				Rows: []widget.TickRow{
					{Label: "Личный", Mark: widget.Mark{Colour: "#4A6B52", Label: "Личный"}, Count: intPtr(n7), Tick: widget.TickOn},
					{Label: "Работа", Mark: widget.Mark{Colour: "#3F6E8C", Label: "Работа"}, Count: intPtr(n4), Tick: widget.TickOn},
					{Label: "Дни рождения", Mark: widget.Mark{Colour: "#A05A2C", Label: "Дни рождения"}, Tick: widget.TickOff},
				},
			},
			{
				Title: "radicale · vps",
				Rows: []widget.TickRow{
					{Label: "Проекты", Mark: widget.Mark{Colour: "#6B4E7A", Label: "Проекты"}, Count: intPtr(n2), Tick: widget.TickOn},
					{Label: "Дежурства", Mark: widget.Mark{Colour: "#8C6239", Label: "Дежурства"}, Count: intPtr(n1), Tick: widget.TickOn},
				},
			},
			{
				Title: "baikal · office",
				State: widget.RailSecAlert,
				Rows: []widget.TickRow{{
					Label: "Общий",
					Mark:  widget.Mark{Colour: "#9A9280", Label: "Общий"},
					State: widget.TickRowUnavailable,
					Menu:  []widget.MenuItem{{Label: "Retry"}},
				}},
			},
			{
				Title: "Subscriptions",
				Rows: []widget.TickRow{
					{Label: "Праздники РФ", Mark: widget.Mark{Colour: "#8C6239", Label: "Праздники РФ"}, State: widget.TickRowReadOnly, Tick: widget.TickOn},
					{Label: "Расписание лиги", Mark: widget.Mark{Colour: "#6B4E7A", Label: "Расписание лиги"}, State: widget.TickRowReadOnly, Tick: widget.TickOff},
				},
			},
		},
		Chips: []widget.Chip{
			{Label: "all", On: true},
			{Label: "релиз"},
			{Label: "работа"},
		},
		Foot: widget.RailFoot{
			Actions: []widget.Action{
				{Label: "New calendar", Icon: "tc-i-plus"},
				{Label: "Connect account", Icon: "tc-i-server"},
				{Label: "Subscribe to a feed", Icon: "tc-i-link"},
			},
			Build: demoBuild,
		},
	}
}

func agendaMain() widget.Main {
	n3, n2, n0, p83 := 3, 2, 0, 83
	return widget.Main{
		Label: "Calendar",
		Body: []widget.Node{
			widget.Head{
				Mark:  widget.Mark{Colours: []string{"#4A6B52", "#3F6E8C", "#6B4E7A"}, Label: "Combined"},
				Title: "Combined calendar",
				Sub:   "5 of 7 calendars · 3 accounts · 10 – 16 August 2026 · Europe/Moscow",
				Actions: []widget.Action{
					{Label: "New event", Kind: widget.ButtonPrimary, Icon: "tc-i-plus"},
				},
				Rank2: []widget.Action{
					{Label: "Import", Icon: "tc-i-up"},
					{Label: "Export", Icon: "tc-i-down"},
					{Label: "Print", Icon: "tc-i-print"},
				},
			},
			widget.Bar{
				Left: []widget.Node{
					widget.Seg{
						Name:    "span",
						Chosen:  "week",
						Options: []widget.SegOption{{Label: "Day", Value: "day"}, {Label: "Week", Value: "week"}, {Label: "Month", Value: "month"}, {Label: "Range", Value: "range"}},
					},
					widget.Range{From: "2026-08-10", To: "2026-08-16", Presets: []widget.Action{{Label: "Show"}}},
					widget.Check{Name: "events", Label: "Events", Checked: true},
					widget.Check{Name: "due", Label: "Tasks due"},
					widget.Check{Name: "journal", Label: "Notes"},
				},
				Right: []widget.Node{
					widget.Check{Name: "compact", Label: "Compact rows"},
				},
			},
			widget.RunBar{
				Summary: "Polled 5 of 6 sources · 14 events · baikal · office unavailable",
				At:      "from cache · 2 min ago",
				Running: true,
				Percent: &p83,
			},
			widget.List{Groups: []widget.ListGroup{{
				Rubric: widget.Rubric{Text: "Monday · 10 August", Count: intPtr(n3)},
				Rows: []widget.Row{
					{
						Variant:  widget.RowTimeCols,
						Selected: true,
						Mark:     widget.Mark{Colour: "#4A6B52", Label: "Личный"},
						Cells: widget.RowCells{
							Lead:    "10:00",
							Title:   "Синк-ап по релизу",
							Detail:  "Переговорная 2",
							Caption: "4 attendees",
							Source:  "nextcloud · Личный",
						},
					},
					{
						Variant: widget.RowTimeCols,
						Mark:    widget.Mark{Colour: "#8C6239", Label: "Дежурства"},
						Cells: widget.RowCells{
							Lead:    "all day",
							Title:   "Дежурство: первая линия",
							Detail:  "duplicate? ×2",
							Caption: "repeats weekly",
							Source:  "2 sources",
						},
					},
					{
						Variant: widget.RowTimeCols,
						Mark:    widget.Mark{Colour: "#3F6E8C", Label: "Работа"},
						Cells: widget.RowCells{
							Lead:    "15:00",
							Title:   "Созвон с подрядчиком",
							Detail:  "Jitsi",
							Caption: "tentative",
							Source:  "nextcloud · Работа",
						},
					},
				},
			}}},
			widget.List{Groups: []widget.ListGroup{{
				Rubric: widget.Rubric{Text: "Tuesday · 11 August", Count: intPtr(n2)},
				Rows: []widget.Row{
					{
						Variant: widget.RowTimeCols,
						Mark:    widget.Mark{Colour: "#6B4E7A", Label: "Проекты"},
						Cells: widget.RowCells{
							Lead:    "09:30",
							Title:   "Ревью архитектуры",
							Detail:  "Переговорная 1",
							Caption: "2 attendees",
							Source:  "radicale · Проекты",
						},
					},
					{
						Variant: widget.RowTimeCols,
						Mark:    widget.Mark{Colour: "#4A6B52", Label: "Личный"},
						Cells: widget.RowCells{
							Lead:   "18:00",
							Title:  "Забрать документы",
							Detail: "МФЦ, Мясницкая 13",
							Source: "nextcloud · Личный",
						},
					},
				},
			}}},
			widget.List{Groups: []widget.ListGroup{{
				Rubric: widget.Rubric{Text: "Wednesday · 12 August", Count: intPtr(n0)},
			}}},
			widget.Empty{Text: "Nothing this day."},
		},
	}
}

func agendaSide() widget.Side {
	fields := widget.Fields{Rows: []widget.FieldRow{
		{Label: "When", Value: "10:00 – 11:00, Monday 10 August"},
		{Label: "Reminders", Value: "10 minutes before"},
		{Label: "Repeats", Value: "Weekly on Mon · until 2026-12-28"},
		{Label: "Where", Value: "Переговорная 2"},
		{Label: "Status", Value: "Confirmed"},
		{Label: "Categories", Value: "release, team"},
	}}
	return widget.Side{
		Label: "Event details",
		Title: "Синк-ап по релизу",
		Sub:   "nextcloud · Личный",
		Open:  true,
		Actions: []widget.Action{
			{Label: "Edit"},
		},
		Blocks: []widget.SideBlock{
			{Fields: &fields},
			{Sec: &widget.Sec{
				Rubric: widget.Rubric{Text: "Description"},
				Body:   []widget.Node{widget.Paragraph("Обсудить оставшиеся блокеры, дату заморозки и кто пишет письмо пользователям.")},
			}},
			{Sec: &widget.Sec{
				Rubric: widget.Rubric{Text: "Attendees", Count: intPtr(4)},
				Body: []widget.Node{widget.Prose{Blocks: []widget.ProseBlock{{
					Kind:  widget.ProseUL,
					Items: []string{"Анна Ковалёва · accepted", "Артём Белов · tentative", "Виктор Гончаров · no reply"},
				}}}},
			}},
			{Sec: &widget.Sec{
				Rubric: widget.Rubric{Text: "Source"},
				Body:   []widget.Node{widget.Prose{Blocks: []widget.ProseBlock{{Kind: widget.ProseCode, Code: "Path  /dav.php/calendars/mix/personal/8f1c.ics\nRead  2 min ago · changed on server 3 days ago"}}}},
			}},
			{Sec: &widget.Sec{Body: []widget.Node{widget.Danger{Title: "Danger zone", Rows: []widget.DangerRow{{
				Title:  "Delete event",
				Text:   "Series is edited whole.",
				Action: widget.Action{Label: "Delete event", Kind: widget.ButtonDangerSolid},
			}}}}}},
		},
	}
}
