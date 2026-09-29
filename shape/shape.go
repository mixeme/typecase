// Package shape checks a client's screen against its frame of the mockup by
// the frame's make-up — the elements, their order and their classes, states
// included — and not by its words or data (task 29 of Carrel's server, step
// 2). A screen shows the application's records, not the frame's: a list has
// as many rows as there are records, a name is another name. What stays is
// the shape.
//
//	m, err := shape.Parse(mockupHTML) // the mockup the client keeps, byte for byte Typecase's
//	err = m.Check("s-find-2", screen.Render())
//
// What is compared: each element's class list and its place among its
// siblings. Siblings of one kind — one tag and one class list but for the
// states a record sets (is-on, is-done…) — are a run: the screen's run may be longer or shorter than the
// frame's, and each of its elements must have the shape of one of the frame's
// (three rows in the frame, one chosen and one with a badge; twelve rows on
// the screen, any of them chosen or badged). So one of two buttons of one
// shape side by side — Import next to Export — is not missed: words are not
// compared. What is not compared: text, attributes,
// the options of a select, the hidden fields of a form. Nor is the element's
// tag when it has classes: its look is its classes, and the meaning of HTML
// (docs/development.md, Р-25, Р-27) changes the tag, not the look — a button
// or words that are a link, a check or a segment that is a label around its
// checkbox or radio button, a tick box that is a checkbox, a pick that is a
// select, a form, a search or a filter that is a form element. The tag of an
// element without classes (b, strong, td…) is compared.
package shape

import (
	"fmt"
	"strings"

	"carrel.dev/typecase/internal/fragment"
)

// Mockup is the mockup read for checking screens.
type Mockup struct {
	m *fragment.Mockup
}

// Parse reads the mockup's HTML.
func Parse(mockupHTML string) (*Mockup, error) {
	m, err := fragment.Parse(mockupHTML)
	if err != nil {
		return nil, err
	}
	return &Mockup{m: m}, nil
}

// Check compares a screen printed for a frame with the frame by its shape
// and names the first place where they part.
func (m *Mockup) Check(frameID, rendered string) error {
	f, ok := m.m.Frame(frameID)
	if !ok {
		return fmt.Errorf("shape: no frame %q", frameID)
	}
	root, err := f.Root()
	if err != nil {
		return err
	}
	nodes, err := fragment.Fragment(rendered)
	if err != nil {
		return err
	}
	if len(nodes) != 1 {
		return fmt.Errorf("%s: the screen printed %d top elements, want 1", frameID, len(nodes))
	}
	want := of(m.m.Normalize(root))
	got := of(m.m.Normalize(nodes[0]))
	if d := diff(want, got, want.label()); d != "" {
		return fmt.Errorf("%s: %s", frameID, d)
	}
	return nil
}

// node is an element's shape: its classes — or, without them, its tag — and
// the shapes of its element children, siblings of one kind in runs.
type node struct {
	tag     string
	classes string
	runs    [][]*node // the element children, siblings of one kind together
	opaque  bool      // a select: its insides are the browser's options, not the frame's drawing
}

// dropped are the elements of the meaning that the frame does not draw.
func dropped(n *fragment.Node) bool {
	if n.Tag == "input" {
		for _, a := range n.Attrs {
			if a == "type=hidden" {
				return true
			}
		}
	}
	for _, c := range n.Classes {
		if c == "tc-check__box" || c == "tc-seg__box" || c == "tc-post" {
			return true
		}
	}
	return n.Tag == "option"
}

func of(n *fragment.Node) *node {
	out := &node{tag: n.Tag, classes: strings.Join(n.Classes, " "), opaque: n.Tag == "select"}
	if out.classes != "" {
		out.tag = ""
	}
	if out.opaque {
		return out
	}
	for _, k := range n.Children {
		if k.Tag == "" || dropped(k) {
			continue
		}
		c := of(k)
		if last := len(out.runs) - 1; last >= 0 && out.runs[last][0].kind() == c.kind() {
			out.runs[last] = append(out.runs[last], c)
		} else {
			out.runs = append(out.runs, []*node{c})
		}
	}
	return out
}

// states are the classes a record's data sets on an element of a run: which
// row is chosen, done, overdue, what a tick shows. The rest of a class list —
// is-primary too — is the kind of the element.
var states = map[string]bool{"is-on": true, "is-off": true, "is-done": true, "is-overdue": true, "is-part": true,
	"is-focus": true, "is-ph": true, "is-error": true, "is-strong": true, "is-linked": true, "is-ext": true, "is-all": true}

// kind is the tag and the classes without the states: what makes siblings one
// run.
func (n *node) kind() string {
	var cs []string
	for _, c := range strings.Fields(n.classes) {
		if !states[c] {
			cs = append(cs, c)
		}
	}
	return n.tag + "." + strings.Join(cs, ".")
}

func (n *node) label() string {
	if n.classes == "" {
		return n.tag
	}
	return "." + strings.ReplaceAll(n.classes, " ", ".")
}

func diff(want, got *node, path string) string {
	if want.tag != got.tag {
		return fmt.Sprintf("%s: tag <%s>, screen <%s>", path, want.tag, got.tag)
	}
	if want.classes != got.classes {
		return fmt.Sprintf("%s: class %q, screen %q", path, want.classes, got.classes)
	}
	if want.opaque || got.opaque {
		return ""
	}
	n := len(want.runs)
	if len(got.runs) < n {
		n = len(got.runs)
	}
	for i := 0; i < n; i++ {
		wr, gr := want.runs[i], got.runs[i]
		p := fmt.Sprintf("%s > %s[%d]", path, wr[0].label(), i+1)
		if wr[0].kind() != gr[0].kind() {
			return fmt.Sprintf("%s: the screen has <%s> where the frame has <%s>", p, gr[0].label(), wr[0].label())
		}
		for _, g := range gr {
			first := ""
			for _, w := range wr {
				d := diff(w, g, p)
				if d == "" {
					first = ""
					break
				}
				if first == "" {
					first = d
				}
			}
			if first != "" {
				return first
			}
		}
	}
	if len(want.runs) > n {
		return fmt.Sprintf("%s: the screen has no <%s> where the frame has it", path, want.runs[n][0].label())
	}
	if len(got.runs) > n {
		return fmt.Sprintf("%s: the screen has <%s> the frame does not draw", path, got.runs[n][0].label())
	}
	return ""
}
