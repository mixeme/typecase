package fragment

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// Splice puts a widget's output into a frame in place of the block r points
// at, and returns the frame's markup. The block's mixins — the classes its
// container prints on it (tc-push, tc-grow, is-2nd…) — go over to the new
// root, so the widget stands where the old one stood. frame is one .frame,
// as the mockup holds it; r.Frame must be its data-frame.
func Splice(frame string, r Ref, rendered string) (string, error) {
	m, err := Parse(frame)
	if err != nil {
		return "", err
	}
	old, err := m.Block(r)
	if err != nil {
		return "", err
	}
	nodes, err := Fragment(rendered)
	if err != nil {
		return "", err
	}
	// A widget may print a sibling beside its block (a dialog keeps the form
	// a PostForm in it posts beside itself): the block takes the mixins.
	var n *html.Node
	for _, c := range nodes {
		if Block(c) == r.Class || n == nil && hasClass(c, r.Class) {
			n = c
		}
	}
	if n == nil {
		return "", fmt.Errorf("%s: the widget printed no .%s", r, r.Class)
	}
	have := map[string]bool{}
	for _, c := range strings.Fields(attr(n, "class")) {
		have[c] = true
	}
	var add []string
	for _, c := range strings.Fields(attr(old, "class")) {
		if !OwnClass(r.Class, c) && !have[c] {
			add = append(add, c)
		}
	}
	if len(add) > 0 {
		setAttr(n, "class", strings.TrimSpace(attr(n, "class")+" "+strings.Join(add, " ")))
	}
	for _, c := range nodes {
		old.Parent.InsertBefore(c, old)
	}
	old.Parent.RemoveChild(old)
	f, _ := m.Frame(r.Frame)
	var b strings.Builder
	if err := html.Render(&b, f.Node); err != nil {
		return "", err
	}
	return b.String(), nil
}

func setAttr(n *html.Node, name, val string) {
	for i, a := range n.Attr {
		if a.Key == name {
			n.Attr[i].Val = val
			return
		}
	}
	n.Attr = append(n.Attr, html.Attribute{Key: name, Val: val})
}
