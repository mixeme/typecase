package refs_test

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"

	"carrel.dev/typecase/internal/fragment"
	"carrel.dev/typecase/internal/refs"
)

// TestCompositions: every composition prints without a panic, prints more
// than its host's record, and goes into the host's frame in place of the
// fragment; building them leaves the records as they were. The browser
// measures them (TestFreeComposition); this keeps the set honest without one.
func TestCompositions(t *testing.T) {
	m := mockup(t)
	before := make([]string, len(refs.Entries))
	for i, e := range refs.Entries {
		before[i] = e.Fixture.Render()
	}
	comps := refs.Compositions(nil)
	for i, e := range refs.Entries {
		if e.Fixture.Render() != before[i] {
			t.Fatalf("%s: building the compositions changed the record", e.Ref)
		}
	}
	frames := map[string]string{}
	slots := map[string]bool{}
	for _, c := range comps {
		slots[c.Owner+"."+c.Slot] = true
		got, err := render(c.Fixture)
		if err != nil {
			t.Errorf("%s: %v", c, err)
			continue
		}
		own := c.Host.Fixture.Render()
		if got == own && !c.Drawn {
			t.Errorf("%s: prints the host's record unchanged", c)
		}
		fr, ok := frames[c.Host.Ref.Frame]
		if !ok {
			f, ok := m.Frame(c.Host.Ref.Frame)
			if !ok {
				t.Fatalf("%s: no frame", c)
			}
			var b strings.Builder
			if err := html.Render(&b, f.Node); err != nil {
				t.Fatal(err)
			}
			fr = b.String()
			frames[c.Host.Ref.Frame] = fr
		}
		if _, err := fragment.Splice(fr, c.Host.Ref, got); err != nil {
			t.Errorf("%s: %v", c, err)
		}
	}
	if len(comps) < 700 || len(slots) < 90 {
		t.Errorf("%d compositions in %d slots: the walk lost its way", len(comps), len(slots))
	}
	t.Logf("%d compositions in %d slots", len(comps), len(slots))
}

func render(r refs.Renderer) (out string, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("panics: %v", p)
		}
	}()
	return r.Render(), nil
}
