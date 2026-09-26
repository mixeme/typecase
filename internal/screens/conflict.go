package screens

import "carrel.dev/typecase/internal/legacy/widget"

// Conflict assembles mockup 7.8: property cards, one combination to write.
func Conflict() string {
	n248, n51 := 248, 51
	return shellFrame(widget.Nav{
		Items: []widget.NavItem{{Label: "Back to list", Icon: "tc-i-back", Href: "/contacts"}},
	}, widget.Rail{
		Sections: []widget.RailSection{{
			Title: "nextcloud · home",
			Rows: []widget.TickRow{
				{Label: "Контакты", Mark: widget.Mark{Colour: "#4A6B52", Label: "Контакты"}, Count: intPtr(n248), Tick: widget.TickOn, Current: true},
				{Label: "Архив", Mark: widget.Mark{Colour: "#A05A2C", Label: "Архив"}, Count: intPtr(n51)},
			},
		}},
		Foot: widget.RailFoot{Build: demoBuild},
	}, widget.Main{
		Label: "Conflict",
		Body: []widget.Node{
			widget.Head{
				Title: "Conflict",
				Sub:   "Another client changed this contact while you were editing it. Nothing has been overwritten.",
			},
			widget.Diff{
				Property: "FN",
				Agreed:   true,
				Options:  []widget.DiffOption{{Label: "Анна Ковалёва", Value: "a"}},
				Chosen:   "a",
			},
			widget.Diff{
				Property: "TEL;work",
				Options: []widget.DiffOption{
					{Label: "Your edit · include · +7 495 120-45-00", Value: "mine"},
					{Label: "Leave it out · —", Value: "out"},
				},
				Chosen: "mine",
			},
			widget.Diff{
				Property: "EMAIL;work",
				Options: []widget.DiffOption{
					{Label: "Your edit · a.kovaleva@example.org", Value: "mine"},
					{Label: "Server · anna.kovaleva@slovo.example", Value: "server"},
				},
				Chosen: "mine",
			},
			widget.Diff{
				Property: "NOTE",
				Options: []widget.DiffOption{
					{Label: "Your edit · include · Пишет по вторникам, отвечает в тот же день.", Value: "mine"},
					{Label: "Leave it out · —", Value: "out"},
				},
				Chosen: "mine",
			},
			widget.Diff{
				Property: "X-EVOLUTION-FILE-AS",
				Options: []widget.DiffOption{
					{Label: "Your edit · —", Value: "mine"},
					{Label: "Server · include · Ковалёва, Анна", Value: "server"},
				},
				Chosen: "server",
			},
			widget.FormFoot{
				Primary: widget.Action{Label: "Write this combination"},
				Rest: []widget.Action{
					{Label: "Use only my edit"},
					{Label: "Use only the server version"},
					{Label: "Cancel"},
				},
			},
		},
	}, widget.Side{}, "")
}
