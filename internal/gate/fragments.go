package gate

import (
	"fmt"
	"strings"

	"carrel.dev/typecase/internal/fragment"
	"carrel.dev/typecase/internal/refs"
)

// CheckWidgetFragments fails when a reference record does not print the
// fragment of its frame (plan 2, step 2): every widget prints its frame.
func CheckWidgetFragments(p ModulePaths) error {
	m, err := fragment.Load(p.Root)
	if err != nil {
		return fmt.Errorf("widget-fragments: %w", err)
	}
	var bad []string
	for _, e := range refs.Entries {
		if err := m.Compare(e.Ref, e.Fixture.Render()); err != nil {
			bad = append(bad, fmt.Sprintf("%T: %v", e.Fixture, err))
		}
	}
	if len(bad) > 0 {
		return fmt.Errorf("widget-fragments:\n  %s", strings.Join(bad, "\n  "))
	}
	return nil
}

// CheckContainerPairs fails when what a container takes and what the mockup
// puts in it differ, either way (plan 2, step 3).
func CheckContainerPairs(p ModulePaths) error {
	m, err := fragment.Load(p.Root)
	if err != nil {
		return fmt.Errorf("container-pairs: %w", err)
	}
	return refs.CheckPairs(m.Pairs(), refs.Entries, refs.Widgets)
}
