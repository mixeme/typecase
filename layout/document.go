package layout

// metaColumn is the reading side column (.m-doc, .m-metacol): the details
// panel, as wide as .tc-side. In a work area of 880 and less (400 of panel
// plus 480 of text) it narrows to narrowMetaColumn (docNarrowCSS).
const metaColumn = "400px"

// narrowMetaColumn is the reading side column in a narrow work area;
// ColsSidebar keeps it at every width — a page column, not the details panel.
const narrowMetaColumn = "296px"

func docNarrowCSS() string {
	return "@container tcbox (max-width:880px){\n" +
		".tc-doc__layout{grid-template-columns:minmax(0,1fr) " + narrowMetaColumn + "}\n" +
		".tc-docmeta{width:auto}\n}\n" // the track sizes it
}

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
