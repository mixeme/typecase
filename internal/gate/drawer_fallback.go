package gate

import (
	"fmt"
	"strings"
)

// The drawer's look of descendants is written once, in
// @container style(--tc-shell-mode: drawer). A browser without style
// queries (Firefox before them: 132 on a phone, Carrel's task 299) skips that
// block whole, and a phone shows the tablet's look with the drawer's
// layout. The same rules stand again right after the block, under the
// width of the window and inside the shell: a window 640 px or narrower
// holds a root as narrow, so every shell in it is in the drawer and the
// copy only repeats what the block already does where style queries work.

const (
	drawerQuery    = "@container style(--tc-shell-mode: drawer)"
	drawerFallback = "@media (max-width:640px)"
	// drawerScope is where the drawer's look reaches: the shell and the
	// overlay that stands for it. :where keeps the copy's selectors as
	// specific as the block's.
	drawerScope = ":where(.tc-shell,.tc-overlay)"
)

// DrawerFallbackRules returns the copy the block asks for: each of its
// rules with every selector under drawerScope, bodies unchanged.
func DrawerFallbackRules(block string) []cssRule {
	var out []cssRule
	for _, r := range parseCSSRules(stripCSSComments(block)) {
		parts := splitSelectors(r.Selector)
		for i, p := range parts {
			parts[i] = drawerScope + " " + p
		}
		out = append(out, cssRule{Selector: strings.Join(parts, ","), Body: r.Body})
	}
	return out
}

// topAtBlocks returns the top-level rules in order: an at-rule's prelude and
// body; a plain rule stands as an empty prelude, so «followed by» means next.
func topAtBlocks(css string) (preludes, bodies []string) {
	css = stripCSSComments(css)
	i := 0
	for i < len(css) {
		at := strings.IndexByte(css[i:], '@')
		brace := strings.IndexByte(css[i:], '{')
		if brace < 0 {
			break
		}
		if at < 0 || at > brace {
			end := matchBrace(css, i+brace)
			if end < 0 {
				break
			}
			preludes = append(preludes, "")
			bodies = append(bodies, "")
			i = end + 1
			continue
		}
		end := matchBrace(css, i+brace)
		if end < 0 {
			break
		}
		preludes = append(preludes, compactSel(css[i+at:i+brace]))
		bodies = append(bodies, css[i+brace+1:end])
		i = end + 1
	}
	return preludes, bodies
}

func normRules(rs []cssRule) []string {
	out := make([]string, len(rs))
	for i, r := range rs {
		var decls []string
		for _, d := range strings.Split(r.Body, ";") {
			if d = compactSel(d); d != "" {
				decls = append(decls, d)
			}
		}
		out[i] = strings.Join(splitSelectors(r.Selector), ",") + "{" + strings.Join(decls, ";") + "}"
	}
	return out
}

// checkMockupDrawerFallback: the drawer's block is followed by its copy
// under the window's width, rule for rule.
func checkMockupDrawerFallback(html string) error {
	product, _, err := mockupSplitCSS(html)
	if err != nil {
		return err
	}
	preludes, bodies := topAtBlocks(product)
	var errs []string
	found := 0
	for i, p := range preludes {
		if p != drawerQuery {
			continue
		}
		found++
		if i+1 >= len(preludes) || preludes[i+1] != drawerFallback {
			errs = append(errs, fmt.Sprintf("%s is not followed by %s with its rules under %s", drawerQuery, drawerFallback, drawerScope))
			continue
		}
		want, got := normRules(DrawerFallbackRules(bodies[i])), normRules(parseCSSRules(stripCSSComments(bodies[i+1])))
		for k := 0; k < len(want) || k < len(got); k++ {
			switch {
			case k >= len(got):
				errs = append(errs, "missing "+want[k])
			case k >= len(want):
				errs = append(errs, "extra "+got[k])
			case want[k] != got[k]:
				errs = append(errs, fmt.Sprintf("rule %d is %s, want %s", k+1, got[k], want[k]))
			}
		}
	}
	if found != 1 {
		errs = append(errs, fmt.Sprintf("%d blocks %s, want 1", found, drawerQuery))
	}
	return mockupLintErrs("drawer-fallback", errs, 6)
}
