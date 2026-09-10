package layout

// RowExcerpt is .tc-row--excerpt ↔ .m-row--note.
var RowExcerpt = Grid{
	Selector:  ".tc-row--excerpt",
	Gap:       "12px",
	Align:     "center",
	Container: "tcbox",
	Tracks: []Track{
		{Slot: "mark", Size: "3px", Empty: EmptyHold},
		{Slot: "title", Size: "minmax(140px,1.2fr)"},
		{Slot: "excerpt", Size: "minmax(170px,2fr)"},
		{Slot: "tags", Size: "minmax(70px,auto)"},
		{Slot: "date", Size: "80px"},
		{Slot: "source", Size: "118px"},
	},
	Narrow: []Breakpoint{
		{
			MaxPx: 700,
			Hide:  []string{"excerpt"},
			Tracks: []Track{
				{Size: "3px"},
				{Size: "minmax(140px,1.4fr)"},
				{Size: "minmax(70px,auto)"},
				{Size: "80px"},
				{Size: "116px"},
			},
		},
		{
			MaxPx:    620,
			Hide:     []string{"tags"},
			PadLeft:  "12px",
			PadRight: "12px",
			Tracks: []Track{
				{Size: "3px"},
				{Size: "minmax(100px,1.4fr)"},
				{Size: "74px"},
				{Size: "96px"},
			},
		},
		{
			MaxPx: 460,
			Hide:  []string{"source"},
			Tracks: []Track{
				{Size: "3px"},
				{Size: "minmax(0,1fr)"},
				{Size: "70px"},
			},
		},
	},
}
