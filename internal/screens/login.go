package screens

import "carrel.dev/typecase/internal/legacy/widget"

// Login assembles mockup 7.20: sign-in without shell chrome.
func Login() string {
	return widget.Card{Body: []widget.Node{
		widget.Brand{SVG: carrelBrandMarkup, Href: "/"},
		widget.Paragraph("One desk. Every server."),
		widget.Paragraph("Calendars, contacts, tasks, notes and files — one place."),
		widget.Form{
			Action: "/login",
			Sets: []widget.FieldSet{{Rows: []widget.Field{
				{Name: "login", Label: "Login", Kind: widget.InputText, Value: "mix", Width: widget.FieldWidthFull},
				{Name: "password", Label: "Password", Kind: widget.InputPassword, Value: "••••••••••", Width: widget.FieldWidthFull},
			}}},
			Foot: widget.FormFoot{
				Primary: widget.Action{Label: "Sign in"},
				Rest: []widget.Action{
					{Label: "Create an account", Href: "/register"},
					{Label: "Forgotten your password?", Href: "/forgot"},
				},
			},
		},
	}}.Render()
}
