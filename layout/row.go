package layout

// RowExcerpt is .tc-row--excerpt ↔ .m-row--note.
var RowExcerpt = Grid{
	Selector:  ".tc-row--excerpt",
	Gap:       "12px",
	Align:     "center",
	Container: "tcbox",
	Tracks: []Track{
		{Slot: "mark", Size: "3px", Empty: EmptyHold},
		{Slot: "title", Size: "minmax(118px,1.2fr)"},
		{Slot: "excerpt", Size: "minmax(170px,2fr)"},
		{Slot: "tags", Size: "minmax(70px,auto)"},
		{Slot: "date", Size: "80px"},
		{Slot: "source", Size: "140px"},
	},
	Narrow: []Breakpoint{
		{
			MaxPx: 700,
			Hide:  []string{"excerpt"},
			Tracks: []Track{
				{Size: "3px"},
				{Size: "minmax(116px,1.4fr)"},
				{Size: "minmax(70px,auto)"},
				{Size: "80px"},
				{Size: "140px"},
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
				{Size: "130px"},
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

// RowAvatarCols is .tc-row--avatarcols ↔ .m-row--contact.
var RowAvatarCols = Grid{
	Selector:  ".tc-row--avatarcols",
	Gap:       "12px",
	Align:     "center",
	Container: "tcbox",
	Tracks: []Track{
		{Slot: "mark", Size: "3px", Empty: EmptyHold},
		{Slot: "avatar", Size: "28px"},
		{Slot: "title", Size: "minmax(130px,1.4fr)"},
		{Slot: "subtitle", Size: "142px"},
		{Slot: "detail", Size: "minmax(150px,1.1fr)"},
		{Slot: "caption", Size: "minmax(110px,.9fr)"},
		{Slot: "source", Size: "140px"},
	},
	Narrow: []Breakpoint{
		{
			MaxPx: 800,
			Hide:  []string{"caption"},
			Tracks: []Track{
				{Size: "3px"},
				{Size: "28px"},
				{Size: "minmax(100px,1.4fr)"},
				{Size: "130px"},
				{Size: "minmax(120px,1fr)"},
				{Size: "140px"},
			},
		},
		{
			MaxPx:    620,
			Hide:     []string{"detail"},
			PadLeft:  "12px",
			PadRight: "12px",
			Tracks: []Track{
				{Size: "3px"},
				{Size: "28px"},
				{Size: "minmax(100px,1.6fr)"},
				{Size: "118px"},
				{Size: "130px"},
			},
		},
		{
			MaxPx: 460,
			Hide:  []string{"source"},
			Tracks: []Track{
				{Size: "3px"},
				{Size: "28px"},
				{Size: "minmax(0,1fr)"},
				{Size: "118px"},
			},
		},
	},
}

// RowTickCols is .tc-row--tickcols ↔ .m-row--file.
var RowTickCols = Grid{
	Selector:  ".tc-row--tickcols",
	Gap:       "12px",
	Align:     "center",
	Container: "tcbox",
	Tracks: []Track{
		{Slot: "tick", Size: "16px", Empty: EmptyHold},
		{Slot: "name", Size: "minmax(208px,2.4fr)"},
		{Slot: "size", Size: "90px"},
		{Slot: "time", Size: "130px"},
		{Slot: "detail", Size: "60px"},
		{Slot: "source", Size: "140px"},
	},
	Narrow: []Breakpoint{
		{
			MaxPx: 800,
			Hide:  []string{"detail"},
			Tracks: []Track{
				{Size: "16px"},
				{Size: "minmax(176px,2.4fr)"},
				{Size: "84px"},
				{Size: "124px"},
				{Size: "140px"},
			},
		},
		{
			MaxPx:    620,
			Hide:     []string{"time"},
			PadLeft:  "12px",
			PadRight: "12px",
			Tracks: []Track{
				{Size: "16px"},
				{Size: "minmax(120px,2fr)"},
				{Size: "78px"},
				{Size: "130px"},
			},
		},
		{
			MaxPx: 460,
			Hide:  []string{"source"},
			Tracks: []Track{
				{Size: "16px"},
				{Size: "minmax(0,1fr)"},
				{Size: "74px"},
			},
		},
	},
}

// RowTimeCols is .tc-row--timecols ↔ .m-row--agenda.
var RowTimeCols = Grid{
	Selector:  ".tc-row--timecols",
	Gap:       "12px",
	Align:     "center",
	Container: "tcbox",
	Tracks: []Track{
		{Slot: "mark", Size: "3px", Empty: EmptyHold},
		{Slot: "lead", Size: "62px"},
		{Slot: "title", Size: "minmax(168px,2fr)"},
		{Slot: "detail", Size: "minmax(110px,1fr)"},
		{Slot: "caption", Size: "96px"},
		{Slot: "source", Size: "140px"},
	},
	Narrow: []Breakpoint{
		{
			MaxPx: 800,
			Hide:  []string{"detail"},
			Tracks: []Track{
				{Size: "3px"},
				{Size: "60px"},
				{Size: "minmax(116px,2fr)"},
				{Size: "88px"},
				{Size: "140px"},
			},
		},
		{
			MaxPx:    620,
			Hide:     []string{"caption"},
			PadLeft:  "12px",
			PadRight: "12px",
			Tracks: []Track{
				{Size: "3px"},
				{Size: "54px"},
				{Size: "minmax(110px,1fr)"},
				{Size: "130px"},
			},
		},
		{
			MaxPx: 460,
			Hide:  []string{"source"},
			Tracks: []Track{
				{Size: "3px"},
				{Size: "50px"},
				{Size: "minmax(0,1fr)"},
			},
		},
	},
}

// RowTickMarkCols is .tc-row--tickmarkcols ↔ .m-row--task.
var RowTickMarkCols = Grid{
	Selector:  ".tc-row--tickmarkcols",
	Gap:       "12px",
	Align:     "center",
	Container: "tcbox",
	Tracks: []Track{
		{Slot: "tick", Size: "14px", Empty: EmptyHold},
		{Slot: "mark", Size: "3px", Empty: EmptyHold},
		{Slot: "title", Size: "minmax(178px,2fr)"},
		{Slot: "tags", Size: "minmax(90px,auto)"},
		{Slot: "due", Size: "84px"},
		{Slot: "progress", Size: "44px"},
		{Slot: "source", Size: "140px"},
	},
	Narrow: []Breakpoint{
		{
			MaxPx: 800,
			Hide:  []string{"progress"},
			Tracks: []Track{
				{Size: "14px"},
				{Size: "3px"},
				{Size: "minmax(112px,2fr)"},
				{Size: "minmax(60px,auto)"},
				{Size: "74px"},
				{Size: "140px"},
			},
		},
		{
			MaxPx:    620,
			Hide:     []string{"tags"},
			PadLeft:  "12px",
			PadRight: "12px",
			Tracks: []Track{
				{Size: "14px"},
				{Size: "3px"},
				{Size: "minmax(100px,2fr)"},
				{Size: "70px"},
				{Size: "130px"},
			},
		},
		{
			MaxPx: 460,
			Hide:  []string{"source"},
			Tracks: []Track{
				{Size: "14px"},
				{Size: "3px"},
				{Size: "minmax(0,1fr)"},
				{Size: "66px"},
			},
		},
	},
}

// RowFind is .tc-row--find ↔ .m-row--find (search hits, not the suggestion stack).
var RowFind = Grid{
	Selector:  ".tc-row--find",
	Gap:       "12px",
	Align:     "center",
	Container: "tcbox",
	Tracks: []Track{
		{Slot: "mark", Size: "3px", Empty: EmptyHold},
		{Slot: "lead", Size: "76px"},
		{Slot: "title", Size: "minmax(208px,2.4fr)"},
		{Slot: "detail", Size: "minmax(120px,1fr)"},
		{Slot: "caption", Size: "74px"},
		{Slot: "source", Size: "140px"},
	},
	Narrow: []Breakpoint{
		{
			MaxPx: 800,
			Hide:  []string{"detail"},
			Tracks: []Track{
				{Size: "3px"},
				{Size: "70px"},
				{Size: "minmax(136px,2.4fr)"},
				{Size: "66px"},
				{Size: "140px"},
			},
		},
		{
			MaxPx:    620,
			Hide:     []string{"caption"},
			PadLeft:  "12px",
			PadRight: "12px",
			Tracks: []Track{
				{Size: "3px"},
				{Size: "62px"},
				{Size: "minmax(110px,2fr)"},
				{Size: "130px"},
			},
		},
		{
			MaxPx: 460,
			Hide:  []string{"source"},
			Tracks: []Track{
				{Size: "3px"},
				{Size: "58px"},
				{Size: "minmax(0,1fr)"},
			},
		},
	},
}

var rowBoxes = []Box{
	{Selector: ".tc-row", Props: map[string]string{
		"gap":           "12px",
		"min-height":    "38px",
		"padding-left":  "var(--tc-page-gutter)",
		"padding-right": "var(--tc-page-gutter)",
		"position":      "relative",
	}},
	{Selector: ".tc-row--avatarcols .tc-row__subtitle", Props: map[string]string{
		"font-family": "var(--tc-font-mono)",
	}},
	{Selector: ".tc-row--tickcols .tc-tick", Props: map[string]string{
		"height": "11px",
		"width":  "11px",
	}},
	{Selector: ".tc-row--tickmarkcols .tc-tick", Props: map[string]string{
		"height": "13px",
		"width":  "13px",
	}},
}

var rowFlows = []Flow{
	{Selector: ".tc-row", Props: map[string]string{"display": "flex"}},
}
