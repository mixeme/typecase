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
}

func liveParityPairs(spec layout.Spec) []parityPair {
	have := map[string]bool{}
	for _, g := range spec.Grids {
		have[g.Selector] = true
	}
	for _, b := range spec.Boxes {
		have[b.Selector] = true
	}
	for _, f := range spec.Flows {
		have[f.Selector] = true
	}
	all := []parityPair{
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
		{Mockup: ".m-shell", Selector: ".tc-shell__body", Props: []string{"grid-template-columns"}},
		{Mockup: ".m-top", Selector: ".tc-top", Props: []string{"height", "gap", "padding"}},
		{Mockup: ".m-search", Selector: ".tc-search", Props: []string{"width"}},
		{Mockup: ".m-search", Selector: ".tc-search__field", Props: []string{"height", "gap", "padding"}},
		{Mockup: ".m-nav", Selector: ".tc-nav", Props: []string{"gap", "padding"}},
	}
	var out []parityPair
	for _, p := range all {
		if have[p.Selector] {
			out = append(out, p)
		}
	}
	return out
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
	baseRules := indexRules(parseCSSRules(base))
	narrow := indexNarrow(atBlocks)

	var errs []string
	for _, pair := range liveParityPairs(spec) {
		mock := baseRules[pair.Mockup]
		if mock == nil {
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
			if strings.Contains(want, "var(") {
				want = resolveMockVars(css, want)
			}
			got := layoutVals[prop]
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
			mcols := narrow[key]
			got := layout.TrackSizes(bp.Tracks)
			if !parityEqual("grid-template-columns", got, mcols) {
				errs = append(errs, fmt.Sprintf("%s @%d mock %q layout %q", pair.Mockup, max, mcols, got))
			}
		}
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

func indexRules(rules []cssRule) map[string]map[string]string {
	out := make(map[string]map[string]string)
	for _, r := range rules {
		for _, sel := range splitSelectors(r.Selector) {
			if !strings.HasPrefix(sel, ".m-") {
				continue
			}
			if strings.ContainsAny(sel, " >+~") {
				continue
			}
			if strings.Count(sel, ".") != 1 {
				continue
			}
			props := cssProps(r.Body)
			prev := out[sel]
			if prev == nil {
				out[sel] = props
				continue
			}
			for k, v := range props {
				prev[k] = v
			}
		}
	}
	return out
}

var maxWidthRe = regexp.MustCompile(`max-width:\s*(\d+)px`)

func indexNarrow(atBlocks []string) map[string]string {
	out := make(map[string]string)
	for _, block := range atBlocks {
		m := maxWidthRe.FindStringSubmatch(block)
		if m == nil {
			continue
		}
		inner := innerAtBlock(block)
		for _, r := range parseCSSRules(inner) {
			for _, sel := range splitSelectors(r.Selector) {
				if !strings.HasPrefix(sel, ".m-") || strings.ContainsAny(sel, " >+~") {
					continue
				}
				cols := cssProps(r.Body)["grid-template-columns"]
				if cols == "" {
					continue
				}
				out[sel+"|"+m[1]] = cols
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
