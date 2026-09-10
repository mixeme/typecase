package layout

// metaColumn is the reading side column (.m-doc, .m-metacol). One number for
// the document layout, its meta block and, from step 7, ColsSidebar.
const metaColumn = "296px"

// Doc is .tc-doc__layout ↔ .m-doc.
var Doc = Grid{
	Selector:  ".tc-doc__layout",
	MinHeight: "0",
	Tracks: []Track{
		{Size: "minmax(0,1fr)"},
		{Size: metaColumn},
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
		"width": metaColumn,
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
