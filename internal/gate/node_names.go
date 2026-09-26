package gate

import (
	"fmt"
	"sort"
	"strings"

	"carrel.dev/typecase/internal/refs"
)

// CheckNodeNames requires that every element a widget prints has a class of
// the library (tc-* or is-*), except the tags the dictionary leaves bare
// (docs/visual/mockup.md, «Словарь»): inline text tags, an svg and its
// insides, and the tags a block fixes by meaning (table rows and cells, list
// items, the terms of a list of fields, a field's label, an option). It reads
// the output of every reference record, not a catalog heuristic.
func CheckNodeNames(_ []CatalogEntry) error {
	return checkNodeNamesIn(func(yield func(name, html string)) {
		for _, e := range refs.Entries {
			yield(fmt.Sprintf("%T %s", e.Fixture, e.Ref), e.Fixture.Render())
		}
	})
}

func checkNodeNamesIn(each func(yield func(name, html string))) error {
	bad := map[string]bool{}
	each(func(name, html string) {
		walkMockupFrameNodes(html, func(n *htmlNode) {
			if mockupClasslessOK(n) {
				return
			}
			for _, c := range n.class {
				if strings.HasPrefix(c, "tc-") || strings.HasPrefix(c, "is-") {
					return
				}
			}
			bad[fmt.Sprintf("%s: <%s> has no class of the library", name, n.tag)] = true
		})
	})
	if len(bad) == 0 {
		return nil
	}
	var out []string
	for k := range bad {
		out = append(out, k)
	}
	sort.Strings(out)
	if len(out) > 8 {
		out = append(out[:8], fmt.Sprintf("… %d in all", len(bad)))
	}
	return fmt.Errorf("node-names: %s", strings.Join(out, "; "))
}

// CheckNodeNamesHTML is the bite helper: one fragment.
func CheckNodeNamesHTML(html string) error {
	return checkNodeNamesIn(func(yield func(name, html string)) { yield("fragment", html) })
}
