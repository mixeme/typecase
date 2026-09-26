package proof

import (
	"path/filepath"
	"strings"
	"testing"
)

func readMockup(t *testing.T) string {
	t.Helper()
	mockup, err := LoadMockup(filepath.Join("..", ".."))
	if err != nil {
		t.Skipf("mockup at MockupRev not reachable (shallow clone?): %v", err)
	}
	return mockup
}

// Every variant is a copy of a style="" its frame really has: the table adds
// names, never numbers.
func TestInlineRulesAreCopiedFromFrames(t *testing.T) {
	mockup := readMockup(t)
	frames := map[string]string{}
	for _, c := range Cases {
		frames[c.Name] = Frame(mockup, c.Label)
		if frames[c.Name] == "" {
			t.Fatalf("frame %q not found", c.Label)
		}
	}
	for _, r := range frameInline {
		if name, val, ok := strings.Cut(r.From, `="`); ok {
			// a presentation attribute such as stroke="var(--tc-danger)"
			if !strings.Contains(frames[r.Frame], r.From) {
				t.Errorf("%s: %s is not in frame %s", r.Selector, r.From, r.Frame)
			}
			if r.CSS != name+":"+strings.TrimSuffix(val, `"`) {
				t.Errorf("%s: %q is not copied from %s", r.Selector, r.CSS, r.From)
			}
			continue
		}
		if !strings.Contains(frames[r.Frame], `style="`+r.From+`"`) {
			t.Errorf("%s: style=%q is not in frame %s", r.Selector, r.From, r.Frame)
		}
		from := map[string]bool{}
		for _, d := range strings.Split(r.From, ";") {
			from[strings.TrimSpace(d)] = true
		}
		for _, d := range strings.Split(r.CSS, ";") {
			if !from[strings.TrimSpace(d)] {
				t.Errorf("%s: %q is not copied from %q", r.Selector, d, r.From)
			}
		}
	}
}

func TestOptionsFindFrames(t *testing.T) {
	o, err := Options(readMockup(t))
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []string{"m-row--note", "m-row--file", "m-table", "m-side", "m-seg"} {
		if !o.Keep[c] {
			t.Errorf("class %s of the frames is not kept", c)
		}
	}
}
