package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"carrel.dev/typecase/internal/framegen"
)

// TestDraftBuildsAlone: the -self draft of every route frame is a Go file
// that builds with the widget package alone — no helper of internal/screens,
// nothing of Typecase but its public API. Each draft is built as a package of
// its own, as a client would keep it.
func TestDraftBuildsAlone(t *testing.T) {
	if testing.Short() {
		t.Skip("builds every draft")
	}
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	secs, _, err := framegen.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	refsInputs, err := os.ReadFile(filepath.Join(root, "internal", "refs", "inputs.go"))
	if err != nil {
		t.Fatal(err)
	}
	dir, err := os.MkdirTemp(filepath.Join(root, "internal"), "drafts-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	n := 0
	for _, s := range secs {
		for _, sc := range s.Screens {
			src, err := framegen.Draft(sc, string(refsInputs))
			if err != nil {
				t.Fatalf("%s: %v", sc.Frame, err)
			}
			if strings.Contains(src, "screens.") || strings.Contains(src, "internal/") {
				t.Errorf("%s: the draft reaches into Typecase's inside", sc.Frame)
			}
			pkg := filepath.Join(dir, strings.ReplaceAll(sc.Frame, "-", "_"))
			if err := os.MkdirAll(pkg, 0o755); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(pkg, "draft.go"), []byte(src), 0o644); err != nil {
				t.Fatal(err)
			}
			n++
		}
	}
	rel, _ := filepath.Rel(root, dir)
	cmd := exec.Command("go", "build", "./"+filepath.ToSlash(rel)+"/...")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("%d drafts do not build: %v\n%s", n, err, out)
	}
	t.Logf("%d drafts build alone", n)
}

// TestDraftsOut (Carrel's task 245): -self -out writes each frame's draft
// to its own directory, the same bytes -self prints for it alone, in one
// call.
func TestDraftsOut(t *testing.T) {
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	frames := []string{"s-find-2", "s-mail-1", "s-tasks-3"}
	out := t.TempDir()
	if err := print(root, frames, true, out); err != nil {
		t.Fatal(err)
	}
	secs, _, err := framegen.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	refsInputs, err := os.ReadFile(filepath.Join(root, "internal", "refs", "inputs.go"))
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range frames {
		got, err := os.ReadFile(filepath.Join(out, id, "draft.go"))
		if err != nil {
			t.Fatalf("%s: %v", id, err)
		}
		for _, s := range secs {
			for _, sc := range s.Screens {
				if sc.Frame != id {
					continue
				}
				want, err := framegen.Draft(sc, string(refsInputs))
				if err != nil {
					t.Fatal(err)
				}
				if string(got) != want {
					t.Errorf("%s: the file differs from the printed draft", id)
				}
			}
		}
	}
}
