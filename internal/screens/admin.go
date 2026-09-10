package screens

import "carrel.dev/typecase/widget"

func adminBackNav() widget.Nav {
	return widget.Nav{
		Items: []widget.NavItem{{Label: "Back to app", Icon: "tc-i-back", Href: "/notes"}},
	}
}

func adminRail(current string) widget.Rail {
	n2, n6 := 2, 6
	return widget.Rail{
		Sections: []widget.RailSection{{
			Title: "Administration",
			Body: []widget.Node{widget.Nav{
				Current: current,
				Items: []widget.NavItem{
					{Label: "Add users", Icon: "tc-i-plus", Href: "/admin", Count: intPtr(n2)},
					{Label: "Users", Icon: "tc-i-con", Href: "#users", Count: intPtr(n6)},
					{Label: "Settings", Icon: "tc-i-gear", Href: "#settings"},
					{Label: "Installation", Icon: "tc-i-shield", Href: "#install"},
					{Label: "DAV validator", Icon: "tc-i-search", Href: "#validator"},
					{Label: "Key escrow", Icon: "tc-i-dup", Href: "#escrow"},
					{Label: "Audit log", Icon: "tc-i-clock", Href: "#audit"},
				},
			}},
		}},
		Foot: widget.RailFoot{Build: demoBuild},
	}
}

// Admin assembles mockup 7.19: three ways to add users.
func Admin() string {
	n2 := 2
	invite := widget.Card{
		Rubric: widget.Rubric{Text: "Invitation"},
		Body: []widget.Node{
			widget.Paragraph("Chooses their own password."),
			widget.Form{
				Sets: []widget.FieldSet{{Rows: []widget.Field{{
					Name:    "role",
					Label:   "Role",
					Kind:    widget.InputSelect,
					Width:   widget.FieldWidthSM,
					Options: []widget.SelectOption{{Label: "User", Value: "user", Selected: true}},
				}}}},
				Foot: widget.FormFoot{Primary: widget.Action{Label: "Create link"}},
			},
			widget.Field{
				Name:  "link",
				Label: "Invitation link",
				Kind:  widget.InputURL,
				Value: "https://carrel.example/invite/7f3a…",
				Extra: widget.Action{Label: "Copy"},
				Width: widget.FieldWidthFull,
			},
			widget.Field{
				Name:  "send",
				Label: "Or send to",
				Kind:  widget.InputEmail,
				Width: widget.FieldWidthMD,
				Extra: widget.Action{Label: "Send invitation"},
			},
		},
	}
	temp := widget.Card{
		Rubric: widget.Rubric{Text: "Temporary password"},
		Body: []widget.Node{
			widget.Badge{Label: "hand to hand", Role: widget.BadgeAlert},
			widget.Msg{Role: widget.MsgAlert, Text: "Until the user changes this password, you can sign in as them."},
			widget.Form{
				Sets: []widget.FieldSet{{Rows: []widget.Field{
					{Name: "t-role", Label: "Role", Kind: widget.InputSelect, Width: widget.FieldWidthXS, Options: []widget.SelectOption{{Label: "User", Value: "user", Selected: true}}},
					{Name: "t-login", Label: "Login", Kind: widget.InputText, Value: "viktor", Width: widget.FieldWidthSM},
					{Name: "t-mail", Label: "Email", Kind: widget.InputEmail, Width: widget.FieldWidthMD},
					{Name: "t-pass", Label: "Temporary password", Kind: widget.InputText, Value: "9f2-brick-lamp", Width: widget.FieldWidthMD},
				}}},
				Foot: widget.FormFoot{Primary: widget.Action{Label: "Create user"}},
			},
		},
	}
	self := widget.Card{
		Rubric: widget.Rubric{Text: "Self-registration"},
		Body: []widget.Node{
			widget.Badge{Label: "open"},
			widget.Paragraph("A visitor creates an account and confirms the address by email. Cannot be turned on until mail works."),
			widget.Field{
				Name:  "publish",
				Label: "Publish the registration form",
				Check: &widget.Check{Name: "publish", Label: "Publish the registration form", Checked: true},
			},
			widget.Field{
				Name:  "reg",
				Label: "Registration address",
				Kind:  widget.InputURL,
				Value: "https://carrel.example/register",
				Extra: widget.Action{Label: "Copy"},
				Width: widget.FieldWidthFull,
			},
			widget.Action{Label: "Save"},
		},
	}
	return shellFrame(adminBackNav(), adminRail("/admin"), widget.Main{
		Label: "Add users",
		Body: []widget.Node{
			widget.Head{
				Title: "Add users",
				Sub:   "Three ways, for three situations. There is no global mode: any of them works.",
			},
			adminWays{invite: invite, temp: temp, self: self},
			widget.Sec{
				Rubric: widget.Rubric{Text: "Pending invitations", Count: intPtr(n2)},
				Body: []widget.Node{widget.Table{
					Columns: []widget.Column{
						{Label: "Login"},
						{Label: "Delivery"},
						{Label: "Status"},
						{Label: "Expires", Kind: widget.ColTime},
					},
					Rows: []widget.TableRow{
						{Cells: []string{"anna", "email", "link unused", "2026-08-23"}, Actions: []widget.Action{{Label: "Revoke", Kind: widget.ButtonDanger}}},
						{Cells: []string{"—", "link copied", "waiting", "2026-08-20"}, Actions: []widget.Action{{Label: "Revoke", Kind: widget.ButtonDanger}}},
					},
				}},
			},
		},
	}, widget.Side{}, "")
}

// adminWays is three equal cards. The catalog has no such form;
// the screen assembles it from Cols (L-3.6).
type adminWays struct {
	invite widget.Card
	temp   widget.Card
	self   widget.Card
}

func (a adminWays) Render() string {
	return widget.Cols{
		Shape: widget.ColsThirds,
		Items: []widget.Node{a.invite, a.temp, a.self},
	}.Render()
}
