package screens

import (
	"testing"

	"carrel.dev/typecase/page"
)

// TestScreensPassCheck: a demo screen is a page, not a picture — every name
// it prints holds, as page.Check reads it (a form a button sends, the Save
// line of an edit, the file a field previews); a frame that names a form it
// does not draw is mended in the mockup. The review of 14.10 found 22 such
// screens, the one of 28.10 102 names on 52.
func TestScreensPassCheck(t *testing.T) {
	for _, r := range Routes {
		got, err := page.Check(`<!doctype html><html><body>` + r.Screen().Render() + `</body></html>`)
		if err != nil {
			t.Fatal(err)
		}
		for _, p := range got {
			t.Errorf("%s: %s", r.Frame, p)
		}
	}
}
