package widget

import (
	"strings"
	"testing"
)

// TestTickTitle: a repeating task's tick says what pressing does in its
// hint (Carrel's task 264); a plain tick has none.
func TestTickTitle(t *testing.T) {
	const words = "Repeating task: Done completes this occurrence and moves the task to the next one, 24 Aug."
	got := Tick{Label: "Полить цветы", Title: words}.Render()
	if !strings.Contains(got, `title="`+words+`"`) {
		t.Errorf("the tick has no hint:\n%s", got)
	}
	if plain := (Tick{Label: "Полить цветы"}).Render(); strings.Contains(plain, "title=") {
		t.Errorf("a plain tick has a hint:\n%s", plain)
	}
}
