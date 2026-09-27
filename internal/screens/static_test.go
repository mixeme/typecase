package screens

import (
	"os"
	"testing"

	"carrel.dev/typecase/internal/fragment"
)

// TestFramesStatic: every route prints the tree of its
// frame — the same widgets in the same order, with the frame's data.
func TestFramesStatic(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	root, err := fragment.FindRoot(wd)
	if err != nil {
		t.Fatal(err)
	}
	m, err := fragment.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range Routes {
		if err := m.CompareFrame(r.Frame, r.Screen().Render()); err != nil {
			t.Error(err)
		}
	}
}
