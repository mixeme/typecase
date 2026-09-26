package fragment

import (
	"sort"
	"strings"
)

// Shape is one form of a block: every occurrence of the class whose tree,
// without text, is the same. Each shape is one reference record (step 4).
type Shape struct {
	Key   string // the tree without text, one line
	Count int    // occurrences in all frames
	First Ref    // the first occurrence, in document order
	Refs  []Ref  // every occurrence
}

// Shapes groups every occurrence of cls in every frame by tree without text,
// most frequent first, then by first appearance.
func (m *Mockup) Shapes(cls string) []Shape {
	byKey := map[string]*Shape{}
	var order []string
	for _, f := range m.Frames {
		for i, n := range f.all(cls) {
			key := Skeleton(m.Normalize(n))
			ref := Ref{Frame: f.ID, Class: cls, N: i + 1}
			s, ok := byKey[key]
			if !ok {
				s = &Shape{Key: key, First: ref}
				byKey[key] = s
				order = append(order, key)
			}
			s.Count++
			s.Refs = append(s.Refs, ref)
		}
	}
	out := make([]Shape, 0, len(order))
	for _, k := range order {
		out = append(out, *byKey[k])
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].Count > out[j].Count })
	return out
}

// Skeleton writes a normalized tree without text and without attribute values
// that are data: tags, classes, compared attribute names, children.
func Skeleton(n *Node) string {
	var b strings.Builder
	skeleton(&b, n)
	return b.String()
}

func skeleton(b *strings.Builder, n *Node) {
	if n.Tag == "" {
		b.WriteString("#")
		return
	}
	b.WriteString(n.Tag)
	for _, c := range n.Classes {
		b.WriteString(".")
		b.WriteString(c)
	}
	for _, a := range n.Attrs {
		name := a
		if i := strings.IndexByte(a, '='); i >= 0 {
			name = a[:i]
		}
		switch name {
		case "type", "href", "disabled", "checked":
			b.WriteString("[" + a + "]") // part of the form, not data
		default:
			b.WriteString("[" + name + "]")
		}
	}
	if len(n.Children) == 0 {
		return
	}
	b.WriteString("(")
	for i, c := range n.Children {
		if i > 0 {
			b.WriteString(" ")
		}
		skeleton(b, c)
	}
	b.WriteString(")")
}
