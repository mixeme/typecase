package refs

import (
	"fmt"
	"sort"
	"strings"

	"carrel.dev/typecase/internal/fragment"
)

// Pending is a form of a block whose widget exists but whose record waits:
// its children belong to a family not written yet, or the frame has a defect
// the owner decides on. Step 7 empties the list.
type Pending struct {
	Ref    fragment.Ref
	Reason string
}

// Waiting lists the pending forms, family by family.
var Waiting []Pending

func wait(reason string, refs ...fragment.Ref) bool {
	for _, r := range refs {
		Waiting = append(Waiting, Pending{Ref: r, Reason: reason})
	}
	return true
}

// Uncovered returns, for every block that has a widget type, the forms that
// have neither a record nor a pending entry.
func Uncovered(m *fragment.Mockup, entries []Entry, waiting []Pending) []string {
	key := func(r fragment.Ref) string {
		n, err := m.Block(r)
		if err != nil {
			return "!" + err.Error()
		}
		return fragment.Skeleton(fragment.WithoutMixins(r.Class, m.Own(n, true)))
	}
	covered := map[string]map[string]bool{}
	mark := func(r fragment.Ref) {
		if covered[r.Class] == nil {
			covered[r.Class] = map[string]bool{}
		}
		covered[r.Class][key(r)] = true
	}
	for _, e := range entries {
		mark(e.Ref)
	}
	for _, p := range waiting {
		mark(p.Ref)
	}
	var out []string
	for class := range covered {
		for _, s := range m.Shapes(class) {
			if !covered[class][s.Key] {
				out = append(out, fmt.Sprintf("%s ×%d first %s: %s", class, s.Count, s.First, s.Key))
			}
		}
	}
	sort.Strings(out)
	return out
}

// WaitingSummary counts the pending forms by reason.
func WaitingSummary(waiting []Pending) string {
	counts := map[string]int{}
	for _, p := range waiting {
		counts[p.Reason]++
	}
	var parts []string
	for r, n := range counts {
		parts = append(parts, fmt.Sprintf("%d: %s", n, r))
	}
	sort.Strings(parts)
	return strings.Join(parts, "; ")
}

// A bar's check form under «⋯» is printed by Menu: its checks as menu items
// with tick boxes, in the same form (Carrel's task 82).
var checkFormInMenu = wait("printed by Menu: a bar's check form as its twin under «⋯»",
	fragment.Ref{Frame: "s-contacts-1", Class: "tc-checkform", N: 1},
	fragment.Ref{Frame: "s-agenda-1", Class: "tc-checkform", N: 2})
