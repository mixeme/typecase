package screens

import "carrel.dev/typecase/widget"

func settingsBackNav() widget.Nav {
	return widget.Nav{
		Items: []widget.NavItem{{Label: "Back to app", Icon: "tc-i-back", Href: "/notes"}},
	}
}

func settingsRail(current string) widget.Rail {
	n3, n2 := 3, 2
	return widget.Rail{
		Sections: []widget.RailSection{{
			Title: "Settings",
			Body: []widget.Node{widget.Nav{
				Current: current,
				Items: []widget.NavItem{
					{Label: "Account", Icon: "tc-i-user", Href: "/account"},
					{Label: "Connections", Icon: "tc-i-server", Href: "/connections", Count: intPtr(n3)},
					{Label: "Attachments", Icon: "tc-i-file", Href: "/attachments"},
					{Label: "New device", Icon: "tc-i-device", Href: "/devices"},
					{Label: "davloom", Icon: "tc-i-all", Href: "/davloom", Count: intPtr(n3)},
					{Label: "Backup", Icon: "tc-i-archive", Href: "/backup", Count: intPtr(n3)},
					{Label: "Published links", Icon: "tc-i-link", Href: "/published", Count: intPtr(n2)},
					{Label: "Appearance", Icon: "tc-i-theme", Href: "/appearance"},
				},
			}},
		}},
		Foot: widget.RailFoot{Build: demoBuild},
	}
}

func settingsPage(current string, main widget.Main, side widget.Side, tail string) string {
	return shellFrame(settingsBackNav(), settingsRail(current), main, side, tail)
}

// Account assembles mockup 7.10: account settings, one state.
func Account() string {
	n2 := 2
	return settingsPage("/account", widget.Main{
		Label: "Account",
		Body: []widget.Node{
			widget.Head{Title: "Account", Sub: "mix · created 11 February 2026"},
			widget.Form{Sets: []widget.FieldSet{
				{
					Legend: "Email",
					Rows: []widget.Field{
						{
							Name:  "email",
							Label: "Address",
							Kind:  widget.InputEmail,
							Value: "mix@example.org",
							Hint:  "Used for invitations and service notices only. Pending: new@example.org",
							Width: widget.FieldWidthLG,
							Extra: widget.Action{Label: "Change"},
						},
					},
				},
			}},
			widget.Sec{
				Rubric: widget.Rubric{Text: "Password"},
				Body: []widget.Node{
					widget.Paragraph("Changing your password re-wraps your data key only; saved connections stay readable."),
					widget.Action{Label: "Change password"},
				},
			},
			widget.Sec{
				Rubric: widget.Rubric{Text: "Key escrow"},
				Body: []widget.Node{
					widget.Paragraph("No copy of your data key is deposited. Nobody but you can decrypt your data — and if you lose your password, it cannot be recovered by anyone."),
					widget.Action{Label: "Deposit my key"},
				},
			},
			widget.Sec{
				Rubric: widget.Rubric{Text: "Sessions", Count: intPtr(n2)},
				Body: []widget.Node{
					widget.Table{
						Columns: []widget.Column{{Label: "Where"}, {Label: "Started", Kind: widget.ColTime}, {Label: "Last seen", Kind: widget.ColTime}},
						Rows: []widget.TableRow{
							{Cells: []string{"Firefox · Windows · this one", "2026-08-16 09:02", "just now"}},
							{Cells: []string{"Carrel desktop · Windows", "2026-08-14 18:40", "2 hours ago"}, Actions: []widget.Action{{Label: "End", Kind: widget.ButtonDanger}}},
						},
					},
					widget.Action{Label: "End every other session", Kind: widget.ButtonDanger},
				},
			},
			widget.Danger{Title: "Danger zone", Rows: []widget.DangerRow{{
				Title:  "Delete account",
				Text:   "Deletes the account and the saved connections. Data on the DAV servers stays: Carrel never touches it.",
				Action: widget.Action{Label: "Delete account", Kind: widget.ButtonDangerSolid},
			}}},
		},
	}, widget.Side{}, "")
}

// Attachments assembles mockup 7.12: where attached files are put.
func Attachments() string {
	return settingsPage("/attachments", widget.Main{
		Label: "Attachments",
		Body: []widget.Node{
			widget.Head{
				Title:   "Attachments",
				Sub:     "Where a file attached to a note, an event or a task is put",
				Actions: []widget.Action{{Label: "Open in Files", Icon: "tc-i-file"}},
			},
			widget.Sec{
				Rubric: widget.Rubric{Text: "Small files stay in the record"},
				Body: []widget.Node{
					widget.Fields{Rows: []widget.FieldRow{
						{Label: "Up to 2 MB", Value: "written into the event, the task or the note itself, and travels with it"},
						{Label: "Larger", Value: "put in the folder below; the record keeps a link to it"},
					}},
					widget.Paragraph("A collection that says it takes less (max-resource-size) lowers the boundary for itself. Nothing here raises it."),
				},
			},
			widget.Form{Sets: []widget.FieldSet{
				{
					Legend: "Folder",
					Rows: []widget.Field{
						{
							Name:    "collection",
							Label:   "Collection",
							Kind:    widget.InputSelect,
							Width:   widget.FieldWidthMD,
							Hint:    "Any file collection you can write to.",
							Options: []widget.SelectOption{{Label: "Documents · sftpgo · vps", Value: "documents", Selected: true}},
						},
						{
							Name:  "folder",
							Label: "Folder in it",
							Kind:  widget.InputText,
							Value: "/carrel/attachments",
							Width: widget.FieldWidthLG,
							Hint:  "Created on first use. It stays an ordinary folder.",
							Extra: widget.Action{Label: "Browse"},
						},
						{
							Name:    "monthly",
							Label:   "A subfolder per month",
							OwnLine: true,
							Hint:    "2026-07, 2026-08. Off means one flat folder.",
							Check:   &widget.Check{Name: "monthly", Label: "A subfolder per month", Checked: true},
						},
					},
				},
				{
					Legend: "File name",
					Rows: []widget.Field{{
						Name:    "written",
						Label:   "Written as",
						Kind:    widget.InputSelect,
						Width:   widget.FieldWidthLG,
						Hint:    "2026-08-10-sync-board.png. A name already taken is never overwritten.",
						Options: []widget.SelectOption{{Label: "Date, then the original name", Value: "date-name", Selected: true}},
					}},
				},
			}},
			widget.Fields{Rows: []widget.FieldRow{
				{Label: "Files", Value: "128 · 1.4 GB"},
				{Label: "Oldest", Value: "2026-02-19"},
				{Label: "Path", Value: "https://files.example.org/dav/Documents/carrel/attachments/", Path: true},
			}},
			widget.Msg{Text: "Changing the folder moves nothing. Records already written carry the old path and keep working; only new attachments go to the new place."},
		},
	}, widget.Side{}, "")
}

// Appearance assembles mockup 7.17: theme, density, lists, dates.
func Appearance() string {
	return settingsPage("/appearance", widget.Main{
		Label: "Appearance",
		Body: []widget.Node{
			widget.Head{Title: "Appearance", Sub: "Theme in this browser · week start and clock in the account"},
			widget.Sec{
				Rubric: widget.Rubric{Text: "Theme"},
				Body: []widget.Node{
					widget.Seg{
						Name:    "theme",
						Chosen:  "auto",
						Options: []widget.SegOption{{Label: "Auto", Value: "auto"}, {Label: "Light", Value: "light"}, {Label: "Dark", Value: "dark"}},
					},
					widget.Paragraph("“Auto” follows the system. Kept in this browser — another device does not see it."),
				},
			},
			widget.Sec{
				Rubric: widget.Rubric{Text: "Density"},
				Body: []widget.Node{
					widget.Seg{
						Name:    "density",
						Chosen:  "comfortable",
						Options: []widget.SegOption{{Label: "Comfortable", Value: "comfortable"}, {Label: "Compact", Value: "compact"}},
					},
					widget.Paragraph("List row height: 38 or 32 px. On a phone always 52. This browser."),
				},
			},
			widget.Form{Sets: []widget.FieldSet{
				{
					Legend: "Lists",
					Note:   "Column sets are remembered per section, in this browser — the button below resets them.",
					Rows: []widget.Field{
						{Name: "panel", Label: "Keep the details panel open", OwnLine: true, Check: &widget.Check{Name: "panel", Label: "Keep the details panel open", Checked: true}},
						{Name: "source", Label: "Show the source column", OwnLine: true, Check: &widget.Check{Name: "source", Label: "Show the source column", Checked: true}},
						{Name: "notesfull", Label: "Open notes in full width", OwnLine: true, Check: &widget.Check{Name: "notesfull", Label: "Open notes in full width"}},
					},
				},
				{
					Legend: "Dates and times",
					Note:   "Week start and 12/24-hour clock live in the account. Time zone is one per instance, set by the administrator.",
					Rows: []widget.Field{
						{Name: "week", Label: "Week starts on", Kind: widget.InputSelect, Width: widget.FieldWidthMD, Options: []widget.SelectOption{{Label: "Monday", Value: "mon", Selected: true}, {Label: "Sunday", Value: "sun"}}},
						{Name: "clock", Label: "Time", Kind: widget.InputSelect, Width: widget.FieldWidthMD, Options: []widget.SelectOption{{Label: "24-hour", Value: "24", Selected: true}, {Label: "12-hour", Value: "12"}}},
						{Name: "tz", Label: "Time zone", Kind: widget.InputSelect, Width: widget.FieldWidthLG, Options: []widget.SelectOption{{Label: "Europe/Moscow", Value: "msk", Selected: true}}},
					},
				},
			}},
			widget.Action{Label: "Reset remembered view settings"},
		},
	}, widget.Side{}, "")
}
