package browser_test

import "testing"

// The catalog and the demo screens measure markup that plan 3 rebuilds: the
// catalog from the reference records, the screens from the new widgets. Until
// then their tests are skipped, not deleted. The behaviour checks of tc.js and
// the geometry checks of L-9.2 run on the new widgets since plan 2, step 5.
const untilPlan3 = "library-from-mockup: каталог и демо-экраны пересобираются в плане 3"

func skipUntilPlan3(t *testing.T) {
	t.Helper()
	t.Skip(untilPlan3)
}
