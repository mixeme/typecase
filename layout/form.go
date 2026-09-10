package layout

// Fields is .tc-fields__row ↔ .mock .m-fields.
var Fields = Grid{
	Selector: ".tc-fields__row",
	Gap:      "7px 14px",
	Tracks: []Track{
		{Size: "104px"},
		{Size: "minmax(0,1fr)"},
	},
}

// Form is .tc-form ↔ .m-form.
var Form = Flow{
	Selector: ".tc-form",
	Props: map[string]string{
		"align-items": "flex-start",
		"display":     "flex",
		"flex-wrap":   "wrap",
		"gap":         "13px 16px",
		"padding":     "14px var(--tc-page-gutter)",
	},
}

// Field is .tc-field ↔ .m-f.
var Field = Flow{
	Selector: ".tc-field",
	Props: map[string]string{
		"display":        "flex",
		"flex-direction": "column",
		"gap":            "4px",
		"min-width":      "0",
	},
}

// FieldSet is .tc-fset ↔ .m-fset.
var FieldSet = Flow{
	Selector: ".tc-fset",
	Props: map[string]string{
		"align-items": "flex-start",
		"display":     "flex",
		"flex":        "1 1 100%",
		"flex-wrap":   "wrap",
		"gap":         "13px 16px",
		"padding":     "12px 0 0",
	},
}

// Sheet is .tc-sheet ↔ .m-sheet.
var Sheet = Box{
	Selector: ".tc-sheet",
	Props: map[string]string{
		"container": "tcbox / inline-size",
		"max-width": "92%",
		"width":     "560px",
	},
}

// Dialog is overlay confirm ↔ .m-dialog.is-narrow.
var Dialog = Box{
	Selector: ".tc-dialog",
	Props: map[string]string{
		"container": "tcbox / inline-size",
		"max-width": "420px",
		"width":     "100%",
	},
}

// In is .tc-in ↔ .m-in.
var In = Box{
	Selector: ".tc-in",
	Props: map[string]string{
		"min-height": "28px",
		"padding":    "0 9px",
	},
}

var formBoxes = []Box{
	Sheet,
	Dialog,
	In,
	{Selector: ".tc-field--xs, .tc-pfield--xs", Props: map[string]string{"width": "96px"}},
	{Selector: ".tc-field--sm, .tc-pfield--sm", Props: map[string]string{"width": "140px"}},
	{Selector: ".tc-field--md, .tc-pfield--md", Props: map[string]string{"width": "210px"}},
	{Selector: ".tc-field--lg, .tc-pfield--lg", Props: map[string]string{"width": "300px"}},
	{Selector: ".tc-field--xl, .tc-pfield--xl", Props: map[string]string{"width": "420px"}},
	{Selector: ".tc-field--full, .tc-pfield--full", Props: map[string]string{
		"flex":      "1 1 100%",
		"max-width": "760px",
	}},
	{Selector: ".tc-sheet .tc-form, .tc-dialog .tc-form, .tc-side .tc-form", Props: map[string]string{
		"padding-left":  "0",
		"padding-right": "0",
	}},
	{Selector: ".tc-formfoot", Props: map[string]string{"flex": "1 1 100%"}},
	{Selector: ".tc-fset > .tc-fset__legend", Props: map[string]string{"flex": "1 1 100%"}},
	{Selector: ".tc-in--textarea", Props: map[string]string{
		"min-height": "76px",
		"padding":    "7px 9px",
	}},
	{Selector: ".tc-sheet .tc-in--textarea", Props: map[string]string{"min-height": "132px"}},
	{Selector: ".tc-bar .tc-in, .tc-docbar .tc-in, .tc-head .tc-in, .tc-sheet__head .tc-in", Props: map[string]string{
		"height":     "26px",
		"min-height": "26px",
	}},
	{Selector: ".tc-sheet__target", Props: map[string]string{
		"flex":       "none",
		"height":     "26px",
		"min-height": "26px",
		"overflow":   "hidden",
		"padding":    "0 26px 0 9px",
		"position":   "relative",
		"width":      "200px",
	}},
	// .m-in.is-pick: the ground that hides overflowing text, then the chevron.
	{Selector: ".tc-sheet__target::before", Props: map[string]string{
		"background":     "var(--tc-surface)",
		"border-radius":  "0 var(--tc-radius) var(--tc-radius) 0",
		"bottom":         "0",
		"content":        `""`,
		"pointer-events": "none",
		"position":       "absolute",
		"right":          "0",
		"top":            "0",
		"width":          "24px",
		"z-index":        "1",
	}},
	{Selector: ".tc-sheet__target::after", Props: map[string]string{
		"border-bottom":  "1.3px solid currentColor",
		"border-right":   "1.3px solid currentColor",
		"content":        `""`,
		"height":         "6px",
		"margin-top":     "-4px",
		"opacity":        ".5",
		"pointer-events": "none",
		"position":       "absolute",
		"right":          "9px",
		"top":            "50%",
		"transform":      "rotate(45deg)",
		"width":          "6px",
		"z-index":        "2",
	}},
}

var formFlows = []Flow{Form, Field, FieldSet}

func formSpacerCSS() string {
	return `*:has(> .tc-field > .tc-field__label) > .tc-field:has(> .tc-check:only-child)::before,
*:has(> .tc-field > .tc-field__label) > .tc-field:has(> .tc-btn:only-child)::before,
*:has(> .tc-field > .tc-field__label) > .tc-field:has(> .tc-seg:only-child)::before{content:"";height:16px;flex:none}
.tc-field.is-own-line::before{display:none}
`
}
