package screens

import (
	"bytes"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"carrel.dev/typecase/internal/framegen"
)

// TestGenerated wants internal/screens to be what cmd/frame2go makes of the
// mockup now: a frame changed in the mockup is a screen regenerated
// (go generate ./internal/screens), never a screen edited by hand.
func TestGenerated(t *testing.T) {
	files, err := framegen.Build(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	var names []string
	for n := range files {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		old, err := os.ReadFile(n)
		if err != nil {
			t.Errorf("%s: %v", n, err)
			continue
		}
		if !bytes.Equal(old, files[n]) {
			t.Errorf("%s is not what the mockup gives: go generate ./internal/screens", n)
		}
	}
	// Every generated file of the package is one the generator makes.
	have, _ := filepath.Glob("*.go")
	for _, p := range have {
		if _, ok := files[p]; ok || p == "routes.go" || strings.HasSuffix(p, "_test.go") {
			continue
		}
		t.Errorf("%s is in internal/screens but the generator does not make it", p)
	}
}
