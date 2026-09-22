package layout

// Shell is .tc-shell__body ↔ .m-shell (rail + main).
var Shell = Grid{
	Selector: ".tc-shell__body",
	Tracks: []Track{
		{Size: "264px"},
		{Size: "minmax(0,1fr)"},
	},
}

// ShellHasSide is .m-shell.has-side. The side track is auto: width lives on
// .tc-side (400 / wide 440), so a closed panel and the 1190 hide collapse it.
var ShellHasSide = Grid{
	Selector: ".tc-shell__body:has(> .tc-shell__side)",
	Tracks: []Track{
		{Size: "264px"},
		{Size: "minmax(0,1fr)"},
		{Size: "auto"},
	},
}

// ShellNoRail is .m-shell.no-rail.
var ShellNoRail = Grid{
	Selector: ".tc-shell__body:not(:has(> .tc-shell__rail))",
	Tracks: []Track{
		{Size: "minmax(0,1fr)"},
	},
}

// Top is .tc-top ↔ .m-top.
var Top = Flow{
	Selector: ".tc-top",
	Props: map[string]string{
		"align-items": "center",
		"display":     "flex",
		"gap":         "14px",
		"height":      "52px",
		"min-height":  "52px",
		"padding":     "0 16px",
	},
}

// Search is width on .tc-search ↔ .m-search.
var Search = Box{
	Selector: ".tc-search",
	Props: map[string]string{
		"flex":      "0 1 560px",
		"max-width": "100%",
		"width":     "560px",
	},
}

// SearchField is the search box ↔ the rest of .m-search.
var SearchField = Flow{
	Selector: ".tc-search__field",
	Props: map[string]string{
		"align-items": "center",
		"display":     "flex",
		"gap":         "7px",
		"height":      "28px",
		"padding":     "0 9px",
	},
}

// Nav is .tc-nav ↔ .m-nav.
var Nav = Flow{
	Selector: ".tc-nav",
	Props: map[string]string{
		"display":        "flex",
		"flex-direction": "column",
		"gap":            "1px",
		"padding":        "0 8px",
	},
}

// NavItem is .tc-nav__item ↔ .m-nav a.
var NavItem = Flow{
	Selector: ".tc-nav__item",
	Props: map[string]string{
		"align-items": "center",
		"display":     "flex",
		"gap":         "9px",
		"height":      "29px",
		"min-height":  "29px",
		"padding":     "0 8px",
	},
}

// NavBottom is NavLayoutBottom ↔ .m-tabbar.
var NavBottom = Flow{
	Selector: ".tc-nav--bottom",
	Props: map[string]string{
		"flex-direction": "row",
	},
}

// NavBottomItem is .m-tabbar a.
var NavBottomItem = Flow{
	Selector: ".tc-nav--bottom .tc-nav__item",
	Props: map[string]string{
		"align-items":     "center",
		"flex":            "1",
		"flex-direction":  "column",
		"font-size":       "9.5px",
		"gap":             "3px",
		"height":          "auto",
		"justify-content": "center",
		"padding":         "7px 0 9px",
	},
}

// Head is .tc-head ↔ .m-head.
var Head = Flow{
	Selector: ".tc-head",
	Props: map[string]string{
		"align-items":     "flex-start",
		"display":         "flex",
		"flex-wrap":       "wrap",
		"gap":             "16px",
		"justify-content": "space-between",
		"padding":         "15px var(--tc-page-gutter) 11px",
	},
}

// Rail is the rail column ↔ .m-rail.
var Rail = Box{
	Selector: ".tc-shell__rail",
	Props: map[string]string{
		"background": "var(--tc-bg)",
		"width":      "264px",
	},
}

// RailSec is .tc-railsec ↔ .m-rail-sec.
var RailSec = Flow{
	Selector: ".tc-railsec",
	Props: map[string]string{
		"display":        "flex",
		"flex-direction": "column",
		"gap":            "3px",
	},
}

var shellBoxes = []Box{
	Search,
	Rail,
	{Selector: ".tc-search__icon", Props: map[string]string{"height": "13px", "width": "13px"}},
	{Selector: ".tc-top .tc-brand", Props: map[string]string{"flex": "1 1 0"}},
	{Selector: ".tc-top__actions", Props: map[string]string{"flex": "1 1 0"}},
	{Selector: ".tc-shell", Props: map[string]string{"container": "tcbox / inline-size"}},
	{Selector: ".tc-main", Props: map[string]string{"container": "tcbox / inline-size"}},
	{Selector: ".tc-side", Props: map[string]string{
		"container": "tcbox / inline-size",
		"width":     "400px",
	}},
	{Selector: ".tc-side--wide", Props: map[string]string{"width": "440px"}},
	{Selector: ".tc-rail", Props: map[string]string{"background": "transparent"}},
	{Selector: ".tc-nav__icon-svg", Props: map[string]string{"height": "15px", "width": "15px"}},
	{Selector: ".tc-nav--bottom .tc-nav__icon-svg", Props: map[string]string{"height": "17px", "width": "17px"}},
	{Selector: ".tc-head__title, .tc-side__title", Props: map[string]string{
		"font-family": "var(--tc-font-serif)",
		"font-size":   "22px",
		"line-height": "1.15",
	}},
	{Selector: ".tc-railsec__title", Props: map[string]string{
		"font-size":      "10px",
		"letter-spacing": "0.18em",
		"padding":        "0 16px",
	}},
	{Selector: ".tc-railsec__rows", Props: map[string]string{
		"display":        "flex",
		"flex-direction": "column",
		"gap":            "3px",
	}},
	{Selector: ".tc-side__head", Props: map[string]string{
		"flex-wrap":     "nowrap",
		"padding":       "15px var(--tc-page-gutter) 11px",
		"padding-right": "calc(var(--tc-page-gutter) + 20px)",
	}},
}

var shellFlows = []Flow{Top, SearchField, Nav, NavItem, NavBottom, NavBottomItem, Head, RailSec}

func shellHideSideCSS() string {
	return "@container tcbox (max-width:1190px){\n.tc-shell:not(.tc-shell--narrow) .tc-shell__side{display:none}\n}\n"
}
