package gate

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"carrel.dev/typecase/layout"
)

type parityPair struct {
	Mockup   string
	Selector string
	Props    []string
	Narrow   []int
	SideAuto bool // last shell track is auto; width lives on .tc-side
	Pad620   bool // .m-row 620 padding vs every row grid's 620 PadLeft/PadRight
}

func parityPairs() []parityPair {
	return []parityPair{
		{Mockup: ".m-src", Selector: ".tc-tickrow", Props: []string{
			"grid-template-columns", "gap", "height", "padding-left", "padding-right",
			"grid-auto-flow", "grid-auto-columns",
		}},
		{Mockup: ".m-row--note", Selector: ".tc-row--excerpt", Props: []string{"grid-template-columns"}, Narrow: []int{700, 620, 460}},
		{Mockup: ".m-row--contact", Selector: ".tc-row--avatarcols", Props: []string{"grid-template-columns"}, Narrow: []int{800, 620, 460}},
		{Mockup: ".m-row--file", Selector: ".tc-row--tickcols", Props: []string{"grid-template-columns"}, Narrow: []int{800, 620, 460}},
		{Mockup: ".m-row--find", Selector: ".tc-row--find", Props: []string{"grid-template-columns"}, Narrow: []int{800, 620, 460}},
		{Mockup: ".m-row--agenda", Selector: ".tc-row--timecols", Props: []string{"grid-template-columns"}, Narrow: []int{800, 620, 460}},
		{Mockup: ".m-row--task", Selector: ".tc-row--tickmarkcols", Props: []string{"grid-template-columns"}, Narrow: []int{800, 620, 460}},
		{Mockup: ".m-row", Selector: ".tc-row", Props: []string{"min-height", "gap"}, Pad620: true},
		{Mockup: ".m-top", Selector: ".tc-top", Props: []string{"height", "gap", "padding"}},
		{Mockup: ".m-search", Selector: ".tc-search", Props: []string{"width"}},
		{Mockup: ".m-search", Selector: ".tc-search__field", Props: []string{"height", "gap", "padding"}},
		{Mockup: ".m-nav a", Selector: ".tc-nav__item", Props: []string{"height", "gap"}},
		{Mockup: ".m-shell", Selector: ".tc-shell__body", Props: []string{"grid-template-columns"}},
		{Mockup: ".m-shell.has-side", Selector: ".tc-shell__body:has(> .tc-shell__side)", Props: []string{"grid-template-columns"}, SideAuto: true},
		{Mockup: ".m-shell.no-rail", Selector: ".tc-shell__body:not(:has(> .tc-shell__rail))", Props: []string{"grid-template-columns"}},
		{Mockup: ".m-shell.no-rail.has-side", Selector: ".tc-shell__body:not(:has(> .tc-shell__rail)):has(> .tc-shell__side)", Props: []string{"grid-template-columns"}, SideAuto: true},
		{Mockup: ".m-form", Selector: ".tc-form", Props: []string{"gap", "padding"}},
		{Mockup: ".mock .m-fields", Selector: ".tc-fields__row", Props: []string{"grid-template-columns"}},
		{Mockup: ".m-sheet", Selector: ".tc-sheet", Props: []string{"width"}},
		{Mockup: ".m-dialog.is-narrow", Selector: ".tc-dialog", Props: []string{"max-width"}},
		{Mockup: ".m-grid", Selector: ".tc-grid", Props: []string{"grid-template-columns", "gap"}},
		{Mockup: ".m-doc", Selector: ".tc-doc__layout", Props: []string{"grid-template-columns"}},
		{Mockup: ".m-title", Selector: ".tc-prose__h1", Props: []string{"font-size", "line-height"}},
	}
}

func specHasSelector(spec layout.Spec, sel string) bool {
	for _, g := range spec.Grids {
		if g.Selector == sel {
			return true
		}
	}
	for _, b := range spec.Boxes {
		if b.Selector == sel {
			return true
		}
	}
	for _, f := range spec.Flows {
		if f.Selector == sel {
			return true
		}
	}
	return false
}

// CheckMockupParity compares layout tables in All() to .m-* rules in the mockup.
func CheckMockupParity(p ModulePaths) error {
	path := filepath.Join(p.Root, "docs", "visual", "carrel-ui-mockups.html")
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("mockup-parity: read mockup: %w", err)
	}
	return CheckMockupParitySpec(layout.All(), string(data))
}

// CheckMockupParitySpec is the bite helper for mockup numbers.
func CheckMockupParitySpec(spec layout.Spec, mockupHTML string) error {
	css := mockupStyle(mockupHTML)
	if css == "" {
		return fmt.Errorf("mockup-parity: no <style> in mockup")
	}
	css = stripCSSComments(css)
	base, atBlocks := cutAtRules(css)
	baseRules := parseCSSRules(base)
	narrow := indexNarrowProps(atBlocks)

	var errs []string
	for _, pair := range parityPairs() {
		if !specHasSelector(spec, pair.Selector) {
			errs = append(errs, pair.Selector+" missing from layout.All()")
			continue
		}
		mock := propsForSelector(baseRules, pair.Mockup)
		if len(mock) == 0 {
			errs = append(errs, pair.Mockup+" missing in mockup")
			continue
		}
		layoutVals := layoutPropValues(spec, pair.Selector)
		for _, prop := range pair.Props {
			want := mock[prop]
			if prop == "padding-left" || prop == "padding-right" {
				if want == "" {
					want = paddingFromShorthand(mock["padding"], prop)
				}
			}
			want = resolveMockVars(css, want)
			got := resolveLayoutVars(layoutVals[prop])
			if pair.SideAuto && prop == "grid-template-columns" {
				got = sideAutoAsMock(got, want)
			}
			if !parityEqual(prop, got, want) {
				errs = append(errs, fmt.Sprintf("%s %s mock %q layout %q", pair.Mockup, prop, want, got))
			}
		}
		g := gridBySelector(spec, pair.Selector)
		for _, max := range pair.Narrow {
			bp := breakpoint(g, max)
			if bp.MaxPx == 0 {
				errs = append(errs, fmt.Sprintf("%s missing layout Narrow %d", pair.Selector, max))
				continue
			}
			key := fmt.Sprintf("%s|%d", pair.Mockup, max)
			mcols := narrow[key]["grid-template-columns"]
			got := layout.TrackSizes(bp.Tracks)
			if !parityEqual("grid-template-columns", got, mcols) {
				errs = append(errs, fmt.Sprintf("%s @%d mock %q layout %q", pair.Mockup, max, mcols, got))
			}
		}
		if pair.Pad620 {
			key := ".m-row|620"
			mp := narrow[key]
			wantL := mp["padding-left"]
			wantR := mp["padding-right"]
			if wantL == "" {
				wantL = paddingFromShorthand(mp["padding"], "padding-left")
			}
			if wantR == "" {
				wantR = paddingFromShorthand(mp["padding"], "padding-right")
			}
			for _, row := range spec.Grids {
				if !strings.HasPrefix(row.Selector, ".tc-row--") {
					continue
				}
				bp := breakpoint(row, 620)
				if bp.MaxPx == 0 {
					continue
				}
				if bp.PadLeft != wantL || bp.PadRight != wantR {
					errs = append(errs, fmt.Sprintf("%s @620 padding mock %s/%s layout %s/%s", row.Selector, wantL, wantR, bp.PadLeft, bp.PadRight))
				}
			}
		}
	}
	if !strings.Contains(layout.CSS(), "@container tcbox (max-width:1190px)") ||
		!strings.Contains(layout.CSS(), ".tc-shell:not(.tc-shell--narrow) .tc-shell__side{display:none}") {
		errs = append(errs, "1190 side hide missing from layout CSS")
	}
	if len(errs) > 0 {
		return fmt.Errorf("mockup-parity: %s", strings.Join(errs, "; "))
	}
	return nil
}

func mockupStyle(html string) string {
	low := strings.ToLower(html)
	start := strings.Index(low, "<style")
	if start < 0 {
		return ""
	}
	gt := strings.Index(html[start:], ">")
	if gt < 0 {
		return ""
	}
	end := strings.Index(low[start+gt+1:], "</style>")
	if end < 0 {
		return ""
	}
	return html[start+gt+1 : start+gt+1+end]
}

func propsForSelector(rules []cssRule, want string) map[string]string {
	want = compactSel(want)
	out := map[string]string{}
	for _, r := range rules {
		for _, sel := range splitSelectors(r.Selector) {
			if compactSel(sel) != want {
				continue
			}
			for k, v := range cssProps(r.Body) {
				out[k] = v
			}
		}
	}
	return out
}

var maxWidthRe = regexp.MustCompile(`max-width:\s*(\d+)px`)

func indexNarrowProps(atBlocks []string) map[string]map[string]string {
	out := make(map[string]map[string]string)
	for _, block := range atBlocks {
		m := maxWidthRe.FindStringSubmatch(block)
		if m == nil {
			continue
		}
		inner := innerAtBlock(block)
		for _, r := range parseCSSRules(inner) {
			for _, sel := range splitSelectors(r.Selector) {
				key := sel + "|" + m[1]
				props := cssProps(r.Body)
				prev := out[key]
				if prev == nil {
					out[key] = props
					continue
				}
				for k, v := range props {
					prev[k] = v
				}
			}
		}
	}
	return out
}

func gridBySelector(spec layout.Spec, sel string) layout.Grid {
	for _, g := range spec.Grids {
		if g.Selector == sel {
			return g
		}
	}
	return layout.Grid{}
}

func breakpoint(g layout.Grid, max int) layout.Breakpoint {
	for _, bp := range g.Narrow {
		if bp.MaxPx == max {
			return bp
		}
	}
	return layout.Breakpoint{}
}

func layoutPropValues(spec layout.Spec, sel string) map[string]string {
	out := map[string]string{}
	g := gridBySelector(spec, sel)
	if g.Selector != "" {
		out["grid-template-columns"] = layout.TrackSizes(g.Tracks)
		out["gap"] = g.Gap
		out["min-height"] = g.MinHeight
		out["padding-left"] = g.PadLeft
		out["padding-right"] = g.PadRight
		out["grid-auto-flow"] = g.AutoFlow
		out["grid-auto-columns"] = g.AutoColumns
		out["align-items"] = g.Align
	}
	for _, box := range spec.Boxes {
		if box.Selector != sel {
			continue
		}
		for k, v := range box.Props {
			out[k] = v
		}
	}
	for _, flow := range spec.Flows {
		if flow.Selector != sel {
			continue
		}
		for k, v := range flow.Props {
			out[k] = v
		}
	}
	return out
}

func paddingFromShorthand(pad, prop string) string {
	parts := strings.Fields(pad)
	if len(parts) == 0 {
		return ""
	}
	switch len(parts) {
	case 1:
		return parts[0]
	case 2:
		if prop == "padding-left" || prop == "padding-right" {
			return parts[1]
		}
		return parts[0]
	case 3:
		if prop == "padding-left" || prop == "padding-right" {
			return parts[1]
		}
		return parts[0]
	case 4:
		if prop == "padding-right" {
			return parts[1]
		}
		if prop == "padding-left" {
			return parts[3]
		}
	}
	return ""
}

func resolveMockVars(css, value string) string {
	re := regexp.MustCompile(`var\((--[^)]+)\)`)
	return re.ReplaceAllStringFunc(value, func(match string) string {
		name := strings.TrimSpace(match[4 : len(match)-1])
		def := regexp.MustCompile(regexp.QuoteMeta(name) + `\s*:\s*([^;}]+)`)
		m := def.FindStringSubmatch(css)
		if m == nil {
			return match
		}
		return strings.TrimSpace(m[1])
	})
}

func resolveLayoutVars(value string) string {
	value = strings.ReplaceAll(value, "var(--tc-page-gutter)", "20px")
	value = strings.ReplaceAll(value, "var(--tc-page-gutter-narrow)", "14px")
	return value
}

func sideAutoAsMock(layoutVal, mockVal string) string {
	a := layout.SplitTracks(layoutVal)
	b := layout.SplitTracks(mockVal)
	if len(a) == 0 || len(a) != len(b) {
		return layoutVal
	}
	if a[len(a)-1] == "auto" {
		a[len(a)-1] = b[len(b)-1]
	}
	return strings.Join(a, " ")
}

func parityEqual(prop, layoutVal, mockVal string) bool {
	layoutVal = strings.Join(strings.Fields(layoutVal), " ")
	mockVal = strings.Join(strings.Fields(mockVal), " ")
	if prop == "grid-template-columns" {
		a := layout.SplitTracks(layoutVal)
		b := layout.SplitTracks(mockVal)
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	return layoutVal == mockVal
}
