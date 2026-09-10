package screens

import "carrel.dev/typecase/widget"

// Backup assembles mockup 7.15: backup jobs, one state.
func Backup() string {
	rail := settingsRail("/backup")
	rail.Note = "Last successful copy · 2026-08-16 03:10 · verified"
	return shellFrame(settingsBackNav(), rail, widget.Main{
		Label: "Backup",
		Body: []widget.Node{
			widget.Head{
				Title:   "Backup",
				Sub:     "Three jobs · their own sources, destinations and triggers",
				Actions: []widget.Action{{Label: "New job", Kind: widget.ButtonPrimary, Icon: "tc-i-plus"}},
				Rank2:   []widget.Action{{Label: "Restore", Icon: "tc-i-back"}},
			},
			widget.Table{
				Columns: []widget.Column{
					{Label: "Job"},
					{Label: "What it copies"},
					{Label: "Where to"},
					{Label: "When"},
					{Label: "Last copy", Kind: widget.ColTime},
				},
				Rows: []widget.TableRow{
					{
						Cells:   []string{"Рабочее, ежедневно", "3 collections", "sftpgo · Documents /backup/work", "cron · 03:00", "2026-08-16 03:10 · 18.4 MB · verified"},
						Actions: []widget.Action{{Label: "Run"}, {Label: "Edit"}, {Label: "More", Icon: "tc-i-more"}},
					},
					{
						Cells:   []string{"Личное на домашний диск", "2 collections + files", "nextcloud · Files /carrel-backup", "on sign-in if older than 7 days", "2026-08-14 09:12 · 2.1 GB · verified"},
						Actions: []widget.Action{{Label: "Run"}, {Label: "Edit"}, {Label: "More", Icon: "tc-i-more"}},
					},
					{
						Cells:   []string{"Contacts archive, by hand", "1 book", "download in browser", "on demand", "never"},
						Actions: []widget.Action{{Label: "Run"}, {Label: "Edit"}, {Label: "More", Icon: "tc-i-more"}},
					},
				},
			},
			widget.Paragraph("verified — the archive was read back from where it was written and the checksum matched."),
		},
	}, widget.Side{}, "")
}

// Published assembles mockup 7.16: issued links and a publish side panel.
func Published() string {
	n2 := 2
	side := widget.Side{
		Label: "Publish",
		Title: "Publish",
		Sub:   "Личный · nextcloud",
		Open:  true,
		Blocks: []widget.SideBlock{
			{Sec: &widget.Sec{
				Rubric: widget.Rubric{Text: "How much of it"},
				Body: []widget.Node{
					widget.Seg{
						Name:    "span",
						Chosen:  "range",
						Options: []widget.SegOption{{Label: "A date range", Value: "range"}, {Label: "The whole calendar", Value: "whole"}},
					},
				},
			}},
			{PFieldRow: &widget.PFieldRow{Fields: []widget.PField{
				{Name: "from", Label: "From", Kind: widget.InputDate, Value: "2026-08-10", Width: widget.PFieldWidthSM},
				{Name: "to", Label: "To", Kind: widget.InputDate, Value: "2026-08-31", Width: widget.PFieldWidthSM},
			}}},
			{PField: &widget.PField{
				Name:    "lives",
				Label:   "Link lives for",
				Kind:    widget.InputSelect,
				Width:   widget.PFieldWidthFull,
				Options: []widget.SelectOption{{Label: "30 days", Value: "30", Selected: true}},
			}},
			{Field: &widget.Field{
				Name:  "ics",
				Label: "Allow .ics download and subscription",
				Check: &widget.Check{Name: "ics", Label: "Allow .ics download and subscription", Checked: true},
			}},
			{Sec: &widget.Sec{
				Rubric: widget.Rubric{Text: "What the visitor gets"},
				Body: []widget.Node{widget.Prose{Blocks: []widget.ProseBlock{{
					Kind:  widget.ProseUL,
					Items: []string{"A list of events for the range, without signing in", "Nothing but this collection: the link gives no access to the account"},
				}}}},
			}},
		},
	}
	return settingsPage("/published", widget.Main{
		Label: "Published links",
		Body: []widget.Node{
			widget.Head{
				Title:   "Published links",
				Sub:     "Read-only, one object or one collection per link · this instance is local: the link opens only from your network",
				Actions: []widget.Action{{Label: "Publish", Kind: widget.ButtonPrimary, Icon: "tc-i-plus"}},
			},
			widget.Badge{Label: "local", Role: widget.BadgeLocal},
			widget.Table{
				Columns: []widget.Column{
					{Label: "What is published"},
					{Label: "Kind"},
					{Label: "Created", Kind: widget.ColTime},
					{Label: "Expires", Kind: widget.ColTime},
					{Label: "Opens", Kind: widget.ColNumber},
				},
				Rows: []widget.TableRow{
					{Cells: []string{"Личный · 10–31 August", "calendar", "2026-08-09", "2026-09-09", "12"}, Actions: []widget.Action{{Label: "Copy"}, {Label: "Revoke", Kind: widget.ButtonDanger}}},
					{Cells: []string{"смета-подрядчик.pdf", "file", "2026-08-12", "2026-08-19", "0"}, Actions: []widget.Action{{Label: "Copy"}, {Label: "Revoke", Kind: widget.ButtonDanger}}},
				},
			},
			widget.Sec{
				Rubric: widget.Rubric{Text: "External feeds", Count: intPtr(n2)},
				Body: []widget.Node{
					widget.Msg{Role: widget.MsgAccent, Text: "A subscribed feed can be handed on by a secret link of its own — the same mechanism, pointed outward."},
					widget.Table{
						Columns: []widget.Column{{Label: "Feed"}, {Label: "Kind"}, {Label: "Source"}},
						Rows: []widget.TableRow{
							{Cells: []string{"Праздники РФ", "ics · external", "https://feeds.example/ru-holidays.ics"}, Actions: []widget.Action{{Label: "Copy"}, {Label: "Revoke", Kind: widget.ButtonDanger}}},
						},
					},
					widget.Msg{Text: "The feed is fetched on every access. Carrel keeps no copy of it."},
				},
			},
		},
	}, side, "")
}
