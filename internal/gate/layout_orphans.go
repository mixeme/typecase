package gate

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"carrel.dev/typecase/internal/screens"
	"carrel.dev/typecase/layout"
)

var (
	orphansDataHook = regexp.MustCompile(`\bdata-tc-`)
	orphansA11y     = regexp.MustCompile(`\b(?:role|aria-[\w-]+)\s*=`)
)

// LayoutOrphansReport lists catalog widgets that appear on at most one assembled
// screen and carry none of: a data-tc- hook, role/aria-, a number in layout.
// The list is for a person; the gate does not fail.
func LayoutOrphansReport(catalog []CatalogEntry) string {
	names := catalogWidgetNames(catalog)
	htmlByWidget := map[string]string{}
	for _, e := range catalog {
		htmlByWidget[e.Widget] += e.HTML
	}
	selectors := layoutSelectors()
	var lines []string
	for _, name := range names {
		routes := screensWithWidget(name)
		if len(routes) > 1 {
			continue
		}
		html := htmlByWidget[name]
		if orphansDataHook.MatchString(html) {
			continue
		}
		if orphansA11y.MatchString(html) {
			continue
		}
		if layoutOwns(name, selectors) {
			continue
		}
		where := "(none)"
		if len(routes) == 1 {
			where = "/" + routes[0]
		}
		lines = append(lines, fmt.Sprintf("  %s  %d  %s", name, len(routes), where))
	}
	var b strings.Builder
	b.WriteString("layout-orphans: catalog widgets on ≤1 screen without data-tc-, role/aria-, or a layout number\n")
	if len(lines) == 0 {
		b.WriteString("  (none)\n")
		return b.String()
	}
	sort.Strings(lines)
	for _, line := range lines {
		b.WriteString(line)
		b.WriteByte('\n')
	}
	return b.String()
}

// CheckLayoutOrphans prints the orphans report and always succeeds.
func CheckLayoutOrphans(catalog []CatalogEntry) error {
	fmt.Print(LayoutOrphansReport(catalog))
	return nil
}

func screensWithWidget(name string) []string {
	var out []string
	for _, sc := range screens.All() {
		if screens.HasWidget(sc.Body(), name) {
			out = append(out, sc.ID)
		}
	}
	sort.Strings(out)
	return out
}

func layoutSelectors() []string {
	spec := layout.All()
	out := make([]string, 0, len(spec.Grids)+len(spec.Boxes)+len(spec.Flows))
	for _, g := range spec.Grids {
		out = append(out, g.Selector)
	}
	for _, b := range spec.Boxes {
		out = append(out, b.Selector)
	}
	for _, f := range spec.Flows {
		out = append(out, f.Selector)
	}
	return out
}

func layoutOwns(name string, selectors []string) bool {
	re := regexp.MustCompile(`\.` + regexp.QuoteMeta(name) + `(?:--|[\s,{:>+~.\[]|$)`)
	for _, sel := range selectors {
		if re.MatchString(sel) {
			return true
		}
	}
	return false
}
