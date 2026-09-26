package fragment

import (
	"regexp"
	"sort"
	"strings"

	"golang.org/x/net/html"
)

// blockClass is a block name of the dictionary (Р-6): tc-block, without
// __element or --variant.
var blockClass = regexp.MustCompile(`^tc-[a-z0-9]+(?:-[a-z0-9]+)*$`)

// BlockClasses returns the block classes an element carries, in its order.
func BlockClasses(n *html.Node) []string {
	var out []string
	for _, c := range strings.Fields(attr(n, "class")) {
		if blockClass.MatchString(c) && !strings.Contains(c, "__") {
			out = append(out, c)
		}
	}
	return out
}

// Block returns the element's block. When an element carries several block
// classes, the block is the most specific one and the rest are mixins the
// context adds: tc-push and tc-grow place an element in its row, tc-line and
// tc-stack lay out a block that has its own name (tc-rubric tc-line), and a
// block built on another names its base as well (tc-btn tc-close is a close
// button). tc-root is the page's root box — base type, container tcbox — and
// never a block: in tc-root tc-shell the block is the shell, and a component
// frame's tc-root tc-root--on-bg is the frame's box, as a catalog cell is.
func Block(n *html.Node) string {
	best, rank := "", -1
	for _, c := range BlockClasses(n) {
		if c == "tc-root" {
			continue
		}
		if r := blockRank(c); r > rank {
			best, rank = c, r
		}
	}
	return best
}

func blockRank(c string) int {
	switch c {
	case "tc-push", "tc-grow":
		return 0
	case "tc-line", "tc-stack":
		return 1
	case "tc-btn":
		return 2
	}
	return 3
}

// Base names the block a block is built on: its class is the block's own.
var Base = map[string]string{"tc-close": "tc-btn"}

// Pair is a parent block and a child block: the child's nearest ancestor that
// carries a block class is the parent.
type Pair struct {
	Parent, Child string
}

// PairUse is a pair with how often and where it first occurs.
type PairUse struct {
	Pair
	Count int
	First Ref // the child, as a Ref
}

// Pairs returns every parent → child block pair in every frame.
func (m *Mockup) Pairs() []PairUse {
	byPair := map[Pair]*PairUse{}
	var order []Pair
	for _, f := range m.Frames {
		counts := map[string]int{}
		var walk func(n *html.Node, parent string)
		walk = func(n *html.Node, parent string) {
			if n.Type == html.ElementNode {
				if hasClass(n, "frame-label") {
					return
				}
				for _, b := range BlockClasses(n) {
					counts[b]++
				}
				if child := Block(n); child != "" {
					if parent != "" {
						p := Pair{parent, child}
						u, ok := byPair[p]
						if !ok {
							u = &PairUse{Pair: p, First: Ref{Frame: f.ID, Class: child, N: counts[child]}}
							byPair[p] = u
							order = append(order, p)
						}
						u.Count++
					}
					parent = child
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				walk(c, parent)
			}
		}
		for c := f.Node.FirstChild; c != nil; c = c.NextSibling {
			walk(c, "")
		}
	}
	out := make([]PairUse, 0, len(order))
	for _, p := range order {
		out = append(out, *byPair[p])
	}
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Parent != out[j].Parent {
			return out[i].Parent < out[j].Parent
		}
		return out[i].Child < out[j].Child
	})
	return out
}

// Blocks returns every block class that is some element's block in a frame.
func (m *Mockup) Blocks() []string {
	seen := map[string]bool{}
	var out []string
	for _, f := range m.Frames {
		var walk func(n *html.Node)
		walk = func(n *html.Node) {
			if n.Type == html.ElementNode {
				if hasClass(n, "frame-label") {
					return
				}
				if b := Block(n); b != "" && !seen[b] {
					seen[b] = true
					out = append(out, b)
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				walk(c)
			}
		}
		walk(f.Node)
	}
	return out
}
