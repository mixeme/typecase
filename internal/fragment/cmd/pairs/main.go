// Command pairs prints every parent → child block pair of the mockup with
// its count and first frame. go run ./internal/fragment/cmd/pairs [parent]
package main

import (
	"fmt"
	"os"

	"carrel.dev/typecase/internal/fragment"
)

func main() {
	wd, _ := os.Getwd()
	root, err := fragment.FindRoot(wd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	m, err := fragment.Load(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, p := range m.Pairs() {
		if len(os.Args) > 1 && p.Parent != os.Args[1] {
			continue
		}
		fmt.Printf("%-16s → %-16s ×%-4d %s\n", p.Parent, p.Child, p.Count, p.First)
	}
}
