// Command shapes prints the forms of a mockup block: every occurrence of the
// class in every frame, grouped by tree without text, with the first frame of
// each. go run ./internal/fragment/cmd/shapes tc-tickrow
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"

	"carrel.dev/typecase/internal/fragment"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: shapes <class> [-all|-html]")
		os.Exit(2)
	}
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
	all := len(os.Args) > 2 && os.Args[2] == "-all"
	withHTML := len(os.Args) > 2 && (os.Args[2] == "-html" || os.Args[2] == "-raw")
	raw := len(os.Args) > 2 && os.Args[2] == "-raw"
	own := len(os.Args) > 2 && os.Args[2] == "-own"
	shapes := m.Shapes(os.Args[1])
	fmt.Printf("%s: %d shapes\n", os.Args[1], len(shapes))
	for i, s := range shapes {
		fmt.Printf("\n%d. ×%d  first %s\n   %s\n", i+1, s.Count, s.First, s.Key)
		if withHTML {
			if n, err := m.Block(s.First); err == nil {
				var b strings.Builder
				_ = html.Render(&b, n)
				if raw {
					fmt.Printf("   %s\n", b.String())
				} else {
					fmt.Printf("   %s\n", strings.Join(strings.Fields(b.String()), " "))
				}
			}
		}
		if own {
			if n, err := m.Block(s.First); err == nil {
				fmt.Printf("   own: %s\n", fragment.Skeleton(m.Own(n, false)))
			}
		}
		if all {
			for _, r := range s.Refs {
				fmt.Printf("     %s\n", r)
			}
		}
	}
}
