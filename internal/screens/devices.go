package screens

import "carrel.dev/typecase/widget"

// Devices assembles mockup 7.13: Apple profile, one state.
func Devices() string {
	left := widget.Card{
		Rubric: widget.Rubric{Text: "What goes into the profile"},
		Body: []widget.Node{
			widget.List{Groups: []widget.ListGroup{{Rows: []widget.Row{
				{Variant: widget.RowTick, Mark: widget.Mark{Colour: "#4A6B52", Label: "Личный"}, Cells: widget.RowCells{Title: "Личный · CalDAV", Due: "nextcloud", Done: true}},
				{Variant: widget.RowTick, Mark: widget.Mark{Colour: "#3F6E8C", Label: "Работа"}, Cells: widget.RowCells{Title: "Работа · CalDAV", Due: "nextcloud", Done: true}},
				{Variant: widget.RowTick, Mark: widget.Mark{Colour: "#4A6B52", Label: "Контакты"}, Cells: widget.RowCells{Title: "Контакты · CardDAV", Due: "nextcloud", Done: true}},
				{Variant: widget.RowTick, Mark: widget.Mark{Colour: "#6B4E7A", Label: "Проекты"}, Cells: widget.RowCells{Title: "Проекты · CalDAV", Due: "radicale"}},
			}}}},
			widget.Msg{
				Role:  widget.MsgAlert,
				Title: "The profile carries passwords in cleartext.",
				Text:  "It is an ordinary file: it lands in Downloads, in a messenger, in the phone’s backup. Open it on the device and delete it. The link lives 10 minutes and opens once.",
			},
		},
	}
	right := widget.Card{
		Rubric: widget.Rubric{Text: "Confirmation"},
		Body: []widget.Node{
			widget.Form{
				Sets: []widget.FieldSet{{Rows: []widget.Field{{
					Name:  "password",
					Label: "Your Carrel password",
					Kind:  widget.InputPassword,
					Value: "••••••••••",
					Hint:  "Asked again: the file hands out the credentials of every ticked server.",
					Width: widget.FieldWidthFull,
				}}}},
				Foot: widget.FormFoot{Primary: widget.Action{Label: "Generate profile"}},
			},
			widget.Sec{
				Rubric: widget.Rubric{Text: "Link ready"},
				Body: []widget.Node{
					widget.Badge{Label: "local", Role: widget.BadgeLocal},
					widget.Field{
						Name:  "link",
						Label: "Profile link",
						Kind:  widget.InputURL,
						Value: "https://carrel.example/d/2f9c…/carrel.mobileconfig",
						Extra: widget.Action{Label: "Copy"},
						Width: widget.FieldWidthFull,
					},
					widget.Bar{Left: []widget.Node{
						widget.Action{Label: "Download now", Icon: "tc-i-down"},
						widget.Action{Label: "Revoke"},
					}},
				},
			},
		},
	}
	return settingsPage("/devices", widget.Main{
		Label: "New device",
		Body: []widget.Node{
			widget.Head{
				Title: "Set up a new device",
				Sub:   "Everything discovery found — as one file or one code · credentials unlocked · 4:12",
				Rank2: []widget.Action{{Label: "Lock now"}},
			},
			widget.Seg{
				Name:    "kind",
				Chosen:  "apple",
				Options: []widget.SegOption{{Label: "Apple", Value: "apple"}, {Label: "Android", Value: "android"}, {Label: "Thunderbird", Value: "tb"}, {Label: "Anything else", Value: "else"}},
			},
			deviceSetup{left: left, right: right},
		},
	}, widget.Side{}, "")
}

// deviceSetup is two cards side by side. The catalog has no such form;
// the screen assembles it from Cols (L-3.6).
type deviceSetup struct {
	left  widget.Node
	right widget.Node
}

func (d deviceSetup) Render() string {
	return widget.Cols{
		Shape: widget.ColsHalves,
		Items: []widget.Node{d.left, d.right},
	}.Render()
}

// Davloom assembles mockup 7.14: proxy address, devices, synthetic tree.
func Davloom() string {
	n3 := 3
	rail := settingsRail("/davloom")
	rail.Note = "Proxy on · writing off"
	return shellFrame(settingsBackNav(), rail, widget.Main{
		Label: "davloom",
		Body: []widget.Node{
			widget.Head{
				Title:   "davloom",
				Sub:     "One address for every device instead of an account for every server · a phone on mobile data will not reach a local address",
				Actions: []widget.Action{{Label: "Add device", Kind: widget.ButtonPrimary, Icon: "tc-i-plus"}},
			},
			widget.Badge{Label: "local", Role: widget.BadgeLocal},
			widget.Form{Sets: []widget.FieldSet{{Rows: []widget.Field{
				{
					Name:  "address",
					Label: "Address for the client",
					Kind:  widget.InputURL,
					Value: "https://carrel.example/dav/",
					Hint:  "In DAVx5 choose “Login with URL and user name”: the user name is your Carrel one, the password is the device’s, not yours.",
					Width: widget.FieldWidthFull,
					Extra: widget.Action{Label: "Copy"},
				},
			}}}},
			widget.Sec{
				Rubric: widget.Rubric{Text: "What devices may do"},
				Body: []widget.Node{
					widget.Seg{
						Name:    "access",
						Chosen:  "read",
						Options: []widget.SegOption{{Label: "Read only", Value: "read"}, {Label: "Read and write", Value: "write"}},
					},
					widget.Paragraph("The ceiling for every device: none of them gets more than is set here."),
				},
			},
			widget.Msg{
				Role: widget.MsgAccent,
				Text: "Changing your Carrel password does not disconnect devices: each has its own key wrapper. To disconnect one, revoke its password here — that is the only thing that stops it.",
			},
			widget.Table{
				Columns: []widget.Column{
					{Label: "Device"},
					{Label: "What it can see"},
					{Label: "Access"},
					{Label: "Created", Kind: widget.ColTime},
					{Label: "Last request", Kind: widget.ColTime},
				},
				Rows: []widget.TableRow{
					{Cells: []string{"Телефон · DAVx5", "4 calendars · 1 book", "read", "2026-06-02", "2026-08-15 09:31"}, Actions: []widget.Action{{Label: "Change", Size: widget.ButtonSizeTable}, {Label: "Revoke", Kind: widget.ButtonDanger}}},
					{Cells: []string{"Ноутбук · Thunderbird", "all · 12 collections", "read", "2026-06-02", "2026-08-15 08:02"}, Actions: []widget.Action{{Label: "Change", Size: widget.ButtonSizeTable}, {Label: "Revoke", Kind: widget.ButtonDanger}}},
					{Cells: []string{"Old phone", "—", "revoked", "2026-03-14", "2026-05-30 12:10"}, Actions: []widget.Action{{Label: "Delete"}}},
				},
			},
			widget.Sec{
				Rubric: widget.Rubric{Text: "What the client sees", Count: intPtr(n3)},
				Body: []widget.Node{widget.Tree{Chosen: "root", Nodes: []widget.TreeNode{
					{ID: "root", Label: "/dav/", Open: true, Children: []widget.TreeNode{
						{ID: "prin", Label: "principals/mix/"},
						{ID: "cals", Label: "calendars/", Open: true, Children: []widget.TreeNode{
							{ID: "c1", Label: "nextcloud-personal/", Mark: widget.Mark{Colour: "#4A6B52", Label: "Личный"}},
							{ID: "c2", Label: "nextcloud-work/", Mark: widget.Mark{Colour: "#3F6E8C", Label: "Работа"}},
							{ID: "c3", Label: "radicale-projects/", Mark: widget.Mark{Colour: "#6B4E7A", Label: "Проекты"}},
						}},
						{ID: "books", Label: "addressbooks/", Open: true, Children: []widget.TreeNode{
							{ID: "b1", Label: "nextcloud-contacts/", Mark: widget.Mark{Colour: "#4A6B52", Label: "Контакты"}},
						}},
						{ID: "files", Label: "files/", Open: true, Children: []widget.TreeNode{
							{ID: "f1", Label: "sftpgo-documents/", Mark: widget.Mark{Colour: "#4A6B52", Label: "Documents"}},
						}},
					}},
				}}},
			},
		},
	}, widget.Side{}, "")
}
