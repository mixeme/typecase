package screens

import (
	"os"
	"strconv"
	"testing"

	"carrel.dev/typecase/internal/fragment"
)

// TestFramesDiffs prints every difference of every route, for fixing screens
// in bulk: TC_FRAME_DIFFS=<max per frame>.
func TestFramesDiffs(t *testing.T) {
	max, _ := strconv.Atoi(os.Getenv("TC_FRAME_DIFFS"))
	if max == 0 {
		t.Skip("set TC_FRAME_DIFFS=<max per frame>")
	}
	wd, _ := os.Getwd()
	root, err := fragment.FindRoot(wd)
	if err != nil {
		t.Fatal(err)
	}
	m, err := fragment.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range Routes {
		ds, err := m.FrameDiffs(r.Frame, r.Screen().Render(), max)
		if err != nil {
			t.Fatal(err)
		}
		for _, d := range ds {
			t.Log(r.Frame + ": " + d)
		}
	}
}
