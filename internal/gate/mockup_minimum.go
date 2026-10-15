package gate

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// MockupContextMinimum forbids a rule of a context to take a control's
// minimum or to cut it (the owner's decision of 15.10, Р-65): a control —
// a field, a button, a switch, a tag, a badge, a range, a search — declares
// how narrow it may go in its own rules; a rule whose subject is such a
// block and whose selector names another block before it may not set
// min-width to zero nor cut it (overflow hidden or clip, text-overflow).
// Where it does not fit, the container has an exit of its own: it wraps,
// folds into «⋯», or scrolls. The exceptions are listed with their reason.
const MockupContextMinimum MockupLintRule = "context-minimum"

// controlBlocks are the blocks whose content has a size of its own.
var controlBlocks = map[string]bool{
	"tc-in": true, "tc-btn": true, "tc-seg": true, "tc-check": true, "tc-tag": true, "tc-badge": true,
	"tc-range": true, "tc-q": true, "tc-tick": true, "tc-segfold": true, "tc-search": true,
}

// contextMinimumAllowed are the selectors that may do it, and why.
var contextMinimumAllowed = map[string]string{
	// a phone: the search folds into its icon, the field lies over it
	// unseen (opacity 0) — nothing that is seen is cut
	".tc-top .tc-search.tc-search": "the folded search: its field is invisible over the icon",
}

var (
	minClassRe   = regexp.MustCompile(`\.(tc-[a-z0-9_-]+)`)
	cssParensRe  = regexp.MustCompile(`\([^()]*\)`)
	cssCompoundR = regexp.MustCompile(`\s*[>+~]\s*|\s+`)
)

func cssBlock(class string) string {
	if i := strings.Index(class, "__"); i >= 0 {
		class = class[:i]
	}
	if i := strings.Index(class, "--"); i >= 0 {
		class = class[:i]
	}
	return class
}

// takesMinimum names the declarations of a body that take a minimum or cut.
func takesMinimum(body string) []string {
	var out []string
	for k, v := range cssProps(body) {
		v = strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(v), "!important"))
		switch k {
		case "min-width":
			if v == "0" || v == "0px" {
				out = append(out, k+":"+v)
			}
		case "overflow", "overflow-x", "overflow-y":
			if v == "hidden" || v == "clip" {
				out = append(out, k+":"+v)
			}
		case "text-overflow":
			if v == "ellipsis" || v == "clip" {
				out = append(out, k+":"+v)
			}
		}
	}
	sort.Strings(out)
	return out
}

func checkMockupContextMinimum(html string) error {
	sheets, err := mockupStyleSheets(html)
	if err != nil {
		return err
	}
	var bad []string
	for _, sh := range sheets {
		if !sh.dataTC {
			continue
		}
		for _, r := range parseCSSRulesDeep(sh.css) {
			took := takesMinimum(r.Body)
			if len(took) == 0 {
				continue
			}
			for _, sel := range splitSelectors(r.Selector) {
				if _, ok := contextMinimumAllowed[sel]; ok {
					continue
				}
				// what is in parentheses (:has, :not) names no context
				flat := sel
				for cssParensRe.MatchString(flat) {
					flat = cssParensRe.ReplaceAllString(flat, "")
				}
				parts := cssCompoundR.Split(strings.TrimSpace(flat), -1)
				if len(parts) < 2 {
					continue
				}
				subject := map[string]bool{}
				for _, m := range minClassRe.FindAllStringSubmatch(parts[len(parts)-1], -1) {
					if !strings.Contains(m[1], "__") && controlBlocks[cssBlock(m[1])] {
						subject[cssBlock(m[1])] = true
					}
				}
				if len(subject) == 0 {
					continue
				}
				own := false
				context := false
				for _, p := range parts[:len(parts)-1] {
					for _, m := range minClassRe.FindAllStringSubmatch(p, -1) {
						b := cssBlock(m[1])
						if b == "tc-root" {
							continue
						}
						context = true
						if subject[b] {
							own = true
						}
					}
				}
				if context && !own {
					bad = append(bad, fmt.Sprintf("%s {%s}", sel, strings.Join(took, "; ")))
				}
			}
		}
	}
	if len(bad) > 0 {
		return fmt.Errorf("%s: a context takes a control's minimum or cuts it — move the minimum into the control's own rule, give the container an exit (Р-65): %s",
			MockupContextMinimum, strings.Join(bad, ", "))
	}
	return nil
}
