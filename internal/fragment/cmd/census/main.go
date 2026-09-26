// Command census prints every block of the mockup with its number of shapes
// and occurrences. go run ./internal/fragment/cmd/census
package main

import (
	"fmt"
	"os"
	"sort"

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
	blocks := m.Blocks()
	sort.Strings(blocks)
	total := 0
	for _, b := range blocks {
		shapes := m.Shapes(b)
		n := 0
		for _, s := range shapes {
			n += s.Count
		}
		total += len(shapes)
		fmt.Printf("%-14s shapes %-4d uses %d\n", b, len(shapes), n)
	}
	fmt.Printf("\n%d blocks, %d shapes\n", len(blocks), total)
}
