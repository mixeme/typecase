package layout

func colsNarrow() []Breakpoint {
	return []Breakpoint{{
		MaxPx:  620,
		Tracks: []Track{{Size: "minmax(0,1fr)"}},
	}}
}

// ColsHalves is .tc-cols--halves: two equal tracks.
var ColsHalves = Grid{
	Selector:  ".tc-cols--halves",
	Gap:       "14px",
	Container: "tcbox",
	Tracks: []Track{
		{Size: "minmax(0,1fr)"},
		{Size: "minmax(0,1fr)"},
	},
	Narrow: colsNarrow(),
}

// ColsThirds is .tc-cols--thirds: three equal tracks.
var ColsThirds = Grid{
	Selector:  ".tc-cols--thirds",
	Gap:       "14px",
	Container: "tcbox",
	Tracks: []Track{
		{Size: "minmax(0,1fr)"},
		{Size: "minmax(0,1fr)"},
		{Size: "minmax(0,1fr)"},
	},
	Narrow: colsNarrow(),
}

// ColsSidebar is .tc-cols--sidebar: fluid + metaColumn, as .m-doc.
var ColsSidebar = Grid{
	Selector:  ".tc-cols--sidebar",
	Gap:       "14px",
	Container: "tcbox",
	Tracks: []Track{
		{Size: "minmax(0,1fr)"},
		{Size: metaColumn},
	},
	Narrow: colsNarrow(),
}

// ColsCol is the column box: named container for nested queries.
var ColsCol = Box{
	Selector: ".tc-cols__col",
	Props: map[string]string{
		"container": "tcbox / inline-size",
		"min-width": "0",
	},
}

var colsBoxes = []Box{ColsCol}
