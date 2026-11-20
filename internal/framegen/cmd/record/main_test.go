package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"carrel.dev/typecase/internal/refs"
)

// TestRemoveTakesOnlyTheRecord: -remove takes a record out with its lines,
// and refuses one that shares a line with something else — the call of
// add(…) or another record — instead of cutting it with them (the review of
// 21.11).
func TestRemoveTakesOnlyTheRecord(t *testing.T) {
	dir := t.TempDir()
	write := func(name, src string) string {
		p := filepath.Join(dir, name)
		if err := os.WriteFile(p, []byte(src), 0o644); err != nil {
			t.Fatal(err)
		}
		return p
	}
	alone := write("alone.go", "package refs\n\nvar _ = add(\n\tform(\"a\", x, \"f-1\", \"tc-btn\"),\n\tform(\"b\", y, \"f-1\", \"tc-btn\"),\n)\n")
	shared := write("shared.go", "package refs\n\nvar _ = add(form(\"a\", x, \"f-1\", \"tc-btn\"),\n\tform(\"b\", y, \"f-1\", \"tc-btn\"))\n")
	ed := editor{}
	if ok, err := ed.remove(refs.Entry{Pos: refs.Pos{File: alone, Line: 4}}, nil); !ok || err != nil {
		t.Fatalf("a record on lines of its own: %v %v", ok, err)
	}
	if err := ed.write(); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(alone)
	if got := string(b); strings.Contains(got, `"a"`) || !strings.Contains(got, `"b"`) || !strings.Contains(got, "add(") {
		t.Errorf("after removing the first record:\n%s", got)
	}
	ed = editor{}
	if _, err := ed.remove(refs.Entry{Pos: refs.Pos{File: shared, Line: 3}}, nil); err == nil {
		t.Error("a record on the line of add(…): removed; want a refusal")
	}
}
