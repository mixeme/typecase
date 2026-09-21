// Command numbered writes docs/visual/numbered-audit.md: every --rN class in
// the mockup with its declarations, its uses and where they stand.
//
// The report is the first step of Х-1 ("Хвосты макета"): the naming decisions
// are made on it, and it empties as the names land.
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"carrel.dev/typecase/internal/gate"
)

func main() {
	root, err := gate.FindModuleRoot()
	if err != nil {
		fail(err)
	}
	html, err := os.ReadFile(gate.MockupFile(root))
	if err != nil {
		fail(err)
	}
	list, err := gate.MockupNumberedClasses(string(html))
	if err != nil {
		fail(err)
	}
	out := filepath.Join(root, "docs", "visual", "numbered-audit.md")
	if err := os.WriteFile(out, []byte(gate.FormatNumberedAudit(list)), 0o644); err != nil {
		fail(err)
	}
	uses := 0
	for _, c := range list {
		uses += len(c.Uses)
	}
	fmt.Printf("numbered: %d classes, %d uses → %s\n", len(list), uses, out)
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
