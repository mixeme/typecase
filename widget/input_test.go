package widget

import (
	"strings"
	"testing"
)

// TestInputOf: a field outside its form names the form (Carrel's task 195);
// without Of it belongs to the form it stands in.
func TestInputOf(t *testing.T) {
	if got := (Input{Name: "url", Of: "install"}).Render(); !strings.Contains(got, ` form="install" name="url"`) {
		t.Errorf("the field does not name its form: %s", got)
	}
	if got := (Input{Name: "url"}).Render(); strings.Contains(got, "form=") {
		t.Errorf("a field without Of names a form: %s", got)
	}
}
