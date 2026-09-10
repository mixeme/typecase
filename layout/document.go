package layout

// Doc is .tc-doc__layout ↔ .m-doc.
var Doc = Grid{
	Selector:  ".tc-doc__layout",
	MinHeight: "0",
	Tracks: []Track{
		{Size: "minmax(0,1fr)"},
		{Size: "296px"},
	},
}

// DocNoMeta is the same layout when the side column is not printed.
var DocNoMeta = Grid{
	Selector:  ".tc-doc__layout:not(:has(> .tc-docmeta))",
	MinHeight: "0",
	Tracks: []Track{
		{Size: "minmax(0,1fr)"},
	},
}

// DocMeta is .tc-docmeta ↔ .m-metacol width.
var DocMeta = Box{
	Selector: ".tc-docmeta",
	Props: map[string]string{
		"width": "296px",
	},
}

// ReadingTitle is the reading heading ↔ .m-title.
var ReadingTitle = Box{
	Selector: ".tc-prose__h1",
	Props: map[string]string{
		"font-size":   "27px",
		"line-height": "1.15",
	},
}
