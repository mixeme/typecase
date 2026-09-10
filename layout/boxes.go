package layout

func line(sel, token string) Box {
	return Box{Selector: sel, Props: map[string]string{"border-color": "var(" + token + ")"}}
}

// Chip and Tag share .m-tag geometry.
var Chip = Box{
	Selector: ".tc-chip, .tc-tag",
	Props: map[string]string{
		"border-radius": "2px",
		"color":         "var(--tc-muted)",
		"font-size":     "10.5px",
		"min-height":    "16px",
		"padding":       "0 5px",
	},
}

var chipOn = Box{
	Selector: ".tc-chip.is-on, .tc-tag.is-on",
	Props: map[string]string{
		"border-color": "var(--tc-accent)",
		"color":        "var(--tc-accent)",
	},
}

// Empty is .tc-empty ↔ .m-empty.
var Empty = Box{
	Selector: ".tc-empty",
	Props: map[string]string{
		"align-items":    "flex-start",
		"display":        "flex",
		"flex-direction": "column",
		"gap":            "10px",
		"padding":        "26px var(--tc-page-gutter)",
		"text-align":     "start",
	},
}

// Msg is .tc-msg ↔ .m-msg.
var Msg = Box{
	Selector: ".tc-msg",
	Props: map[string]string{
		"background":        "var(--tc-card)",
		"border-left-color": "var(--tc-edge)",
		"border-left-style": "solid",
		"border-left-width": "2px",
		"border-radius":     "0 var(--tc-radius) var(--tc-radius) 0",
		"margin":            "12px var(--tc-page-gutter)",
		"padding":           "9px 12px",
	},
}

var msgAccent = Box{
	Selector: ".tc-msg--accent",
	Props: map[string]string{
		"border-left-color": "var(--tc-accent)",
	},
}

var msgAlert = Box{
	Selector: ".tc-msg--alert",
	Props: map[string]string{
		"background":        "transparent",
		"border":            "var(--tc-stroke) solid var(--tc-danger)",
		"border-left-width": "2px",
		"color":             "var(--tc-danger)",
	},
}

// Tiles is .tc-grid ↔ .m-grid (Box: auto-fill, no slots).
var Tiles = Box{
	Selector: ".tc-grid",
	Props: map[string]string{
		"display":               "grid",
		"gap":                   "14px",
		"grid-template-columns": "repeat(auto-fill,minmax(132px,1fr))",
	},
}

// Tree is .tc-tree ↔ .m-tree (Box: no slots).
var Tree = Box{
	Selector: ".tc-tree",
	Props: map[string]string{
		"display": "grid",
		"gap":     "1px",
	},
}

var TreeRow = Box{
	Selector: ".tc-tree__row",
	Props: map[string]string{
		"align-items":    "center",
		"display":        "grid",
		"gap":            "7px",
		"grid-auto-flow": "column",
		"height":         "26px",
		"min-height":     "26px",
	},
}

// Queue is .tc-queue ↔ .m-transfers.
var Queue = Box{
	Selector: ".tc-queue",
	Props: map[string]string{
		"bottom":   "16px",
		"position": "absolute",
		"right":    "var(--tc-page-gutter)",
		"width":    "330px",
		"z-index":  "4",
	},
}

var listRubric = Box{
	Selector: ".tc-list > .tc-rubric",
	Props: map[string]string{
		"align-items":         "baseline",
		"background":          "var(--tc-bg)",
		"border-bottom-style": "solid",
		"border-bottom-width": "var(--tc-stroke)",
		"border-color":        "var(--tc-rule)",
		"display":             "flex",
		"gap":                 "12px",
		"justify-content":     "space-between",
		"padding":             "15px var(--tc-page-gutter) 5px",
	},
}

var restBoxes = []Box{
	Chip,
	chipOn,
	Empty,
	Msg,
	msgAccent,
	msgAlert,
	Tiles,
	Tree,
	TreeRow,
	Queue,
	DocMeta,
	ReadingTitle,
	listRubric,
	{Selector: ".tc-skel__line", Props: map[string]string{
		"background": "var(--tc-hair)",
		"height":     "9px",
	}},
	{Selector: ".tc-meter", Props: map[string]string{
		"flex":   "none",
		"height": "3px",
		"width":  "64px",
	}},
	{Selector: ".tc-thumb", Props: map[string]string{"aspect-ratio": "4/3"}},
	{Selector: ".tc-table__th", Props: map[string]string{
		"font-size":      "10px",
		"letter-spacing": "0.12em",
		"padding":        "9px 12px",
	}},
	{Selector: ".tc-table__td", Props: map[string]string{"padding": "7px 12px"}},
	{Selector: ".tc-table__td .tc-btn", Props: map[string]string{
		"height":     "22px",
		"min-height": "22px",
	}},
	{Selector: ".tc-tile.is-selected", Props: map[string]string{
		"border-color": "var(--tc-accent)",
	}},
	{Selector: ".tc-menu__sep, .tc-usermenu__sep", Props: map[string]string{
		"background":   "var(--tc-hair)",
		"border-color": "var(--tc-hair)",
	}},
	line(".tc-btn", "--tc-edge"),
	line(".tc-in", "--tc-edge"),
	line(".tc-seg", "--tc-edge"),
	line(".tc-search__field", "--tc-edge"),
	line(".tc-pick__box", "--tc-edge"),
	line(".tc-tick", "--tc-edge"),
	line(".tc-tickrow__tick", "--tc-edge"),
	line(".tc-check__box", "--tc-edge"),
	line(".tc-kbd__key", "--tc-edge"),
	line(".tc-menu__control", "--tc-edge"),
	line(".tc-menu__popover", "--tc-edge"),
	line(".tc-usermenu__control", "--tc-edge"),
	line(".tc-usermenu__popover", "--tc-edge"),
	line(".tc-sheet", "--tc-edge"),
	line(".tc-dialog", "--tc-edge"),
	line(".tc-queue", "--tc-edge"),
	line(".tc-table__td", "--tc-hair"),
	line(".tc-card", "--tc-hair"),
	line(".tc-tile", "--tc-hair"),
	line(".tc-diff__option", "--tc-hair"),
	line(".tc-chip", "--tc-hair"),
	line(".tc-tag", "--tc-hair"),
	line(".tc-fset", "--tc-hair"),
	line(".tc-formfoot", "--tc-hair"),
	line(".tc-sheet__head", "--tc-hair"),
	line(".tc-sheet__foot", "--tc-hair"),
	line(".tc-railfoot", "--tc-hair"),
	line(".tc-top", "--tc-rule"),
	line(".tc-rail", "--tc-rule"),
	line(".tc-side", "--tc-rule"),
	line(".tc-bar", "--tc-rule"),
	line(".tc-table__th", "--tc-rule"),
	line(".tc-badge", "--tc-rule"),
	line(".tc-docmeta", "--tc-rule"),
}
