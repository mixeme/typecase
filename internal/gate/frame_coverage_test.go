package gate

import (
	"strings"
	"testing"
)

// The frame-coverage gate bites: a route frame with neither a screen nor a
// gap, a screen of a frame that is not a route, a frame both built and a gap.
func TestFrameCoverageBites(t *testing.T) {
	want := map[string]bool{"a-1": true, "a-2": true, "a-3": true}
	if err := frameCoverage(want, []string{"a-1", "a-2"}, []string{"a-3"}); err != nil {
		t.Fatalf("a covered set: %v", err)
	}
	for name, c := range map[string]struct {
		routes, gaps []string
		says         string
	}{
		"missing": {[]string{"a-1"}, []string{"a-3"}, "a-2: a route frame with neither"},
		"stray":   {[]string{"a-1", "a-2", "b-1"}, []string{"a-3"}, "b-1: a screen, but not a route frame"},
		"twice":   {[]string{"a-1", "a-2", "a-3"}, []string{"a-3"}, "a-3: a screen and a gap"},
	} {
		err := frameCoverage(want, c.routes, c.gaps)
		if err == nil || !strings.Contains(err.Error(), c.says) {
			t.Errorf("%s: got %v, want %q", name, err, c.says)
		}
	}
	if got := gapFrames("| Кадр | Подпись |\n|------|---|\n| s-auth-1 | Вход | x |\n| m-print-1 | Агенда | y |\n"); strings.Join(got, ",") != "s-auth-1,m-print-1" {
		t.Errorf("gap rows: %v", got)
	}
}
