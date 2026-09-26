// Command mixins prints, for every block, the mixins its root carries in the
// frames and how often. go run ./internal/fragment/cmd/mixins
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

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
	for _, b := range m.Blocks() {
		counts := map[string]int{}
		for _, s := range m.Shapes(b) {
			for _, r := range s.Refs {
				n, err := m.Block(r)
				if err != nil {
					continue
				}
				for _, c := range m.Normalize(n).Classes {
					if !fragment.OwnClass(b, c) {
						counts[c]++
					}
				}
			}
		}
		if len(counts) == 0 {
			continue
		}
		var parts []string
		for c, n := range counts {
			parts = append(parts, fmt.Sprintf("%s×%d", c, n))
		}
		sort.Strings(parts)
		fmt.Printf("%-14s %s\n", b, strings.Join(parts, " "))
	}
}
