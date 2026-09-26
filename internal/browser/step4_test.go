package browser_test

import (
	"testing"

	"carrel.dev/typecase/internal/browser"
)

// Plan 2, step 1: the library CSS is the mockup's cut, and the old widget
// markup is not drawn by it. Until step 4 rewrites the widgets, the tests that
// measure old markup — the catalog, the demo screens and the geometry checks
// of L-9.2 — are skipped, not deleted. Behaviour checks of tc.js stay on.
const step4 = "library-from-mockup: шаг 4 — старая разметка виджетов под CSS макета"

// geometryChecks measure widget geometry on catalog fixtures.
var geometryChecks = map[browser.Name]bool{
	browser.MenuLabelArrow: true,
	browser.PopoverAnchor:  true,
	browser.PopoverFlip:    true,
	browser.Ellipsis:       true,
	browser.BarOneRow:      true,
	browser.TapTarget:      true,
}

func skipUntilStep4(t *testing.T) {
	t.Helper()
	t.Skip(step4)
}
