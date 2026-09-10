package screens

import "carrel.dev/typecase/widget"

// Connections assembles mockup 7.11: servers, collections, one refused source.
func Connections() string {
	n3 := 3
	return settingsPage("/connections", widget.Main{
		Label: "Connections",
		Body: []widget.Node{
			widget.Head{
				Title: "Connections",
				Sub:   "Signed in as mix · 3 servers · 12 collections · 2 subscriptions",
				Actions: []widget.Action{
					{Label: "Connect account", Kind: widget.ButtonPrimary, Icon: "tc-i-plus"},
				},
				Rank2: []widget.Action{{Label: "Subscribe to a feed", Icon: "tc-i-link"}},
			},
			widget.Rubric{Text: "DAV servers", Count: intPtr(n3)},
			widget.Card{
				Rubric: widget.Rubric{Text: "nextcloud · home"},
				Body: []widget.Node{
					widget.Paragraph("https://cloud.example.org/remote.php/dav/ · mix · last answer 2 min ago"),
					widget.Bar{Left: []widget.Node{
						widget.Action{Label: "New collection", Icon: "tc-i-plus", Size: widget.ButtonSizeTable},
						widget.Action{Label: "Show password", Size: widget.ButtonSizeTable},
						widget.Action{Label: "Test", Size: widget.ButtonSizeTable},
						widget.Action{Label: "Edit", Size: widget.ButtonSizeTable},
						widget.Action{Label: "Disconnect", Kind: widget.ButtonDanger, Size: widget.ButtonSizeTable},
					}},
					widget.Table{
						Columns: []widget.Column{
							{Label: "Collection"},
							{Label: "What is in it"},
							{Label: "Objects", Kind: widget.ColNumber},
							{Label: "Changed on the server", Kind: widget.ColTime},
						},
						Rows: []widget.TableRow{
							{Cells: []string{"Личный", "calendar", "128", "3 h ago"}, Actions: []widget.Action{{Label: "More", Icon: "tc-i-more"}}},
							{Cells: []string{"Работа", "calendar", "74", "yesterday"}, Actions: []widget.Action{{Label: "More", Icon: "tc-i-more"}}},
							{Cells: []string{"Контакты", "address book", "248", "4 d ago"}, Actions: []widget.Action{{Label: "More", Icon: "tc-i-more"}}},
							{Cells: []string{"Записи", "journal", "86", "2 h ago"}, Actions: []widget.Action{{Label: "More", Icon: "tc-i-more"}}},
							{Cells: []string{"Дела", "tasks", "30", "20 min ago"}, Actions: []widget.Action{{Label: "More", Icon: "tc-i-more"}}},
						},
					},
				},
			},
			widget.Card{
				Rubric: widget.Rubric{Text: "sftpgo · vps"},
				Body: []widget.Node{
					widget.Paragraph("https://files.example.org/dav/ · mix · plain WebDAV, no principal · last answer 5 min ago"),
					widget.Table{
						Columns: []widget.Column{
							{Label: "Folder"},
							{Label: "What is in it"},
							{Label: "Objects", Kind: widget.ColNumber},
							{Label: "Changed on the server", Kind: widget.ColTime},
						},
						Rows: []widget.TableRow{
							{Cells: []string{"Documents", "files", "412", "5 min ago"}, Actions: []widget.Action{{Label: "More", Icon: "tc-i-more"}}},
							{Cells: []string{"Public", "files · read-only", "96", "3 d ago"}, Actions: []widget.Action{{Label: "More", Icon: "tc-i-more"}}},
						},
					},
				},
			},
			widget.Card{
				Rubric: widget.Rubric{Text: "baikal · office"},
				Body: []widget.Node{
					widget.Msg{
						Role: widget.MsgAlert,
						Text: "https://dav.office.example/dav.php/ · m.yenuchenko · last answer refused",
						Diag: widget.Diag{
							Summary:  "Discovery",
							Lines:    []string{"[well-known]  GET /.well-known/caldav → 404", "[base]        PROPFIND /dav.php/ → 401 Unauthorized"},
							Expanded: true,
						},
					},
					widget.Bar{Left: []widget.Node{
						widget.Action{Label: "Show password", Size: widget.ButtonSizeTable},
						widget.Action{Label: "Test again", Size: widget.ButtonSizeTable},
						widget.Action{Label: "Edit", Size: widget.ButtonSizeTable},
						widget.Action{Label: "Disconnect", Kind: widget.ButtonDanger, Size: widget.ButtonSizeTable},
					}},
				},
			},
		},
	}, widget.Side{}, "")
}
