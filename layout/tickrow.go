package layout

// TickRow is .tc-tickrow ↔ .m-src.
var TickRow = Grid{
	Selector:    ".tc-tickrow",
	Align:       "center",
	Gap:         "9px",
	MinHeight:   "26px",
	PadLeft:     "12px",
	PadRight:    "14px",
	AutoFlow:    "column",
	AutoColumns: "minmax(0,auto)",
	Tracks: []Track{
		{Slot: "tick", Size: "12px", Empty: EmptyHold},
		{Slot: "mark", Size: "3px", Empty: EmptyHold},
		{Slot: "label", Size: "minmax(3.5em,1fr)", Empty: EmptyHold},
	},
}

var tickRowBoxes = []Box{
	{Selector: ".tc-tickrow", Props: map[string]string{"height": "26px", "position": "relative"}},
	{Selector: ".tc-tickrow__tick", Props: map[string]string{"height": "11px", "width": "11px"}},
	{Selector: ".tc-tickrow .tc-mark", Props: map[string]string{"height": "14px", "min-height": "14px", "width": "3px"}},
	{Selector: ".tc-row .tc-mark", Props: map[string]string{"height": "22px", "min-height": "22px"}},
	{Selector: ".tc-head .tc-mark, .tc-side__title .tc-mark", Props: map[string]string{"height": "20px", "min-height": "20px"}},
	{Selector: ".tc-list", Props: map[string]string{"container": "tcbox / inline-size"}},
	// Row link: the main text is the link, its ::after covers the row. Tick and
	// tags sit above it; a menu is positioned already and comes later in order.
	{Selector: ".tc-row__link::after, .tc-tickrow__link::after", Props: map[string]string{
		"content":  `""`,
		"inset":    "0",
		"position": "absolute",
	}},
	{Selector: ".tc-row .tc-tick, .tc-tickrow__tick", Props: map[string]string{"z-index": "1"}},
	{Selector: ".tc-row__tags", Props: map[string]string{"position": "relative"}},
}
