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
// states a record sets (is-on, is-done…) — are a run: the screen's run may be
// longer than the frame's, not shorter, each of its elements must have the
// shape of one of the frame's, and each of the frame's must be on the screen
// but for its states (three rows in the frame, one with a badge; twelve rows
// on the screen, one of them at least with a badge, any of them chosen). What
// the frame draws, the screen has (Carrel's task 76): a button of the head, a
// section of the card, the time in a field's value — words are not compared,
// so of two buttons of one shape a missing one is told by their number. A
// repeating row of rows (a month and its records, a term and its value, a day
// and its records) may repeat any number of times; repeated as
// often as the frame draws it or more, it has each part the frame draws in
// it. What is not compared: text, attributes, a template's contents,
// the options of a select, a datalist, the hidden fields of a form, an
// element with hidden but for the progress bar, a menu and the scrim (a
// script opens those, and the frame draws them hidden), the zone of
// messages (its messages stand in its place), a picture in an
// avatar or a thumbnail (Src: data, the frame draws the place), the menu
// of a button with its menu (the frames draw it beside the screen). Nor is the element's
// tag when it has classes: its look is its classes, and the meaning of HTML
// (docs/development.md, Р-25, Р-27) changes the tag, not the look — a button
// or words that are a link, a check or a segment that is a label around its
// checkbox or radio button, a tick box that is a checkbox, a pick that is a
// select, a form, a search or a filter that is a form element. The tag of an
// element without classes (b, strong, td…) is compared.
//
// The rules one by one (the review of 14.10), each with its test:
//
//   - classes are compared, the tag only without classes; words and
//     attributes are not (TestShapeNotWords);
//   - siblings of one kind but for states are a run, the screen's no shorter
//     than the frame's; each of its elements is like one of the frame's, and
//     each of the frame's is on the screen but for its states
//     (TestShapeNotWords, TestMissingIsRed);
//   - a block of runs the frame repeats repeats any number of times; with as
//     many repeats as the frame or more, each element the frame draws at a
//     place is on the screen there (TestRepeatsAnyLength, TestMissingIsRed).
//     With fewer, nothing is required of the repeats: even what every
//     repeat of the frame has is data — a day without a meeting with
//     attendees has no badge;
//   - not drawn: a hidden field of a form, a tick's box, a picture in an
//     avatar or a thumbnail, an option, a template, a datalist, an element
//     with hidden but for the progress bar, a menu and the scrim
//     (TestDroppedAndOpaque, TestAvatarPhotoIsData, TestTemplatesAndEmptyDays,
//     TestDatalistIsNotDrawn, TestHiddenIsNotDrawn);
//   - a select, and a pick that holds the browser's list, are opaque
//     (TestDroppedAndOpaque);
//   - a closed menu is its button (TestClosedMenuIsItsButton);
//   - a value's link stands for its words (TestMeaningIsNoDifference);
//   - the zone of messages stands for its messages, empty it is nothing
//     (TestNoticesAreTheirMessages);
//   - is-photo of an avatar is data (TestAvatarPhotoIsData);
//   - a screen is one element (TestDroppedAndOpaque).
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
	c := &cmp{memo: map[pair]string{}}
	if d := c.diff(want, got, want.label()); d != "" {
		return fmt.Errorf("%s: %s", frameID, d)
	}
	return nil
}

// node is an element's shape: its classes — or, without them, its tag — and
// the shapes of its element children, siblings of one kind in runs.
type node struct {
	tag     string
	name    string // the element's own tag, kept when its classes stand for it
	classes string
	runs    [][]*node        // the element children, siblings of one kind together
	opaque  bool             // a select: its insides are the browser's options, not the frame's drawing
	with    map[string]*node // the node with another element's states (as), made once
}

// dropped are the elements of the meaning that the frame does not draw.
func dropped(n *fragment.Node) bool {
	// An element with hidden is not drawn, and a frame need not draw it — a
	// field «Custom…» hidden until its choice (Carrel's task 161). Kept are
	// the places a script opens or fills, which a frame draws hidden so that
	// the screen has them: the progress bar, a menu, the scrim.
	if n.Hidden && !opened(n) {
		return true
	}
	if n.Tag == "input" {
		for _, a := range n.Attrs {
			if a == "type=hidden" {
				return true
			}
		}
	}
	for _, c := range n.Classes {
		// tc-in__list: the browser's list inside a pick — the frame draws the
		// pick's words, the screen may send them (Carrel's task 149).
		if c == "tc-check__box" || c == "tc-seg__box" || c == "tc-tile__box" || c == "tc-tree__box" || c == "tc-post" || c == "tc-in__list" {
			return true
		}
		// A picture in its place is data, not drawing: the frame draws the
		// place (initials on a colour, a tile), the screen fills it (Src).
		if c == "tc-avatar__img" || c == "tc-thumb__img" {
			return true
		}
	}
	// A template is what a script copies in (another phone, another
	// reminder), not what the page shows (Carrel's task 75). Nor does a
	// datalist show anything: it is the choices of a line as one types
	// (TagAdd.Options, Carrel's task 146).
	return n.Tag == "option" || n.Tag == "template" || n.Tag == "datalist"
}

// opened says whether a hidden element is a place a script opens or fills.
func opened(n *fragment.Node) bool {
	for _, c := range n.Classes {
		if c == "tc-runbar" || c == "tc-menu" || c == "tc-scrim" {
			return true
		}
	}
	return false
}

// holdsList says whether an element is a pick that holds the browser's list
// (tc-in__list): what stands in it is the list's, not the frame's drawing —
// a frame draws the chosen words, a mark, an aside; the screen its options.
func holdsList(n *fragment.Node) bool {
	for _, k := range n.Children {
		for _, c := range k.Classes {
			if c == "tc-in__list" {
				return true
			}
		}
	}
	return false
}

// dataClass says whether a class of an element is its data, not its kind:
// an avatar with a photo or with initials is the same avatar.
func dataClass(classes []string, c string) bool {
	for _, x := range classes {
		if (c == "is-photo" && x == "tc-avatar") || (c == "is-ph" && x == "tc-avatar--photo") {
			return true
		}
	}
	return false
}

// kids are an element's children, with a value's link (tc-fields__link)
// opened: the link is the value's meaning, its words are the value (task 52).
// So is the zone of messages (tc-notices): it has no look, its messages
// stand where it stands, as a frame draws them (Carrel's task 238).
func kids(n *fragment.Node) []*fragment.Node {
	var out []*fragment.Node
	for _, k := range n.Children {
		link := false
		for _, c := range k.Classes {
			if c == "tc-fields__link" || c == "tc-notices" {
				link = true
			}
		}
		if link {
			out = append(out, kids(k)...)
			continue
		}
		out = append(out, k)
	}
	return out
}

func of(n *fragment.Node) *node {
	// A button with its menu (tc-drop, tc-drop__anchor) is its button: the
	// frames draw the button where it stands and the menu's contents in a
	// frame of its own beside the screen (Carrel's task 49).
	for _, c := range n.Classes {
		if c == "tc-drop" || c == "tc-drop__anchor" {
			for _, k := range n.Children {
				if k.Tag != "" && !dropped(k) {
					return of(k)
				}
			}
		}
	}
	var cls []string
	for _, c := range n.Classes {
		if !dataClass(n.Classes, c) {
			cls = append(cls, c)
		}
	}
	out := &node{tag: n.Tag, name: n.Tag, classes: strings.Join(cls, " "), opaque: n.Tag == "select" || holdsList(n)}
	if out.classes != "" {
		out.tag = ""
	}
	if out.opaque {
		return out
	}
	for _, k := range kids(n) {
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

// covered says whether each of the frame's elements is like one of the
// screen's, but for its states.
func (c *cmp) covered(wr, gr []*node, p string) string {
	for k, w := range wr {
		first := ""
		for _, g := range gr {
			d := c.diff(w.as(g), g, p)
			if d == "" {
				first = ""
				break
			}
			if first == "" && w.kind() == g.kind() {
				first = d
			}
		}
		if first != "" {
			return first // where the nearest of the screen's parts from it
		}
		found := false
		for _, g := range gr {
			found = found || c.diff(w.as(g), g, p) == ""
		}
		if !found {
			return fmt.Sprintf("%s: the screen has none like the frame's <%s> #%d", p, w.label(), k+1)
		}
	}
	return ""
}

// as is the node with the states of g: a frame's element is found on the
// screen whatever record is chosen or done there.
func (n *node) as(g *node) *node {
	if n.kind() != g.kind() || n.classes == g.classes {
		return n
	}
	if n.with == nil {
		n.with = map[string]*node{}
	}
	if c, ok := n.with[g.classes]; ok {
		return c
	}
	c := *n
	c.classes = g.classes
	c.with = nil
	n.with[g.classes] = &c
	return &c
}

func (n *node) label() string {
	if n.classes == "" {
		return n.tag
	}
	return "." + strings.ReplaceAll(n.classes, " ", ".")
}

// cmp compares a frame with a screen, each pair of elements once: a run is
// matched both ways, and the same pair comes back.
type cmp struct {
	memo map[pair]string
}

type pair struct{ want, got *node }

func (c *cmp) diff(want, got *node, path string) string {
	k := pair{want, got}
	if d, ok := c.memo[k]; ok {
		return d
	}
	d := c.walk(want, got, path)
	c.memo[k] = d
	return d
}

func (c *cmp) walk(want, got *node, path string) string {
	if want.tag != got.tag {
		return fmt.Sprintf("%s: tag <%s>, screen <%s>", path, want.tag, got.tag)
	}
	if want.classes != got.classes {
		return fmt.Sprintf("%s: class %q, screen %q", path, want.classes, got.classes)
	}
	if want.opaque || got.opaque {
		return ""
	}
	if pre, p, wn, gn, ok := repeat(want.runs, got.runs); ok {
		// A repeating row of rows (a month and its records, a term and its
		// value) is compared as a run: any number of repeats on the screen,
		// each like one of the frame's (Carrel's task 61).
		for i := 0; i < pre; i++ {
			if d := c.runDiff(want.runs[i], got.runs[i], path, i); d != "" {
				return d
			}
		}
		for u := 0; u < gn; u++ {
			for j := 0; j < p; j++ {
				gr := got.runs[pre+u*p+j]
				var pool []*node
				for v := 0; v < wn; v++ {
					pool = append(pool, want.runs[pre+v*p+j]...)
				}
				if d := c.poolDiff(pool, gr, path, pre+u*p+j); d != "" {
					return d
				}
			}
		}
		if gn >= wn {
			// As many repeats as the frame draws, or more: each element the
			// frame draws at a place of the block is on the screen at that
			// place (Carrel's task 76) — a field's time and badge as well.
			for j := 0; j < p; j++ {
				var pool, gr []*node
				for v := 0; v < wn; v++ {
					pool = append(pool, want.runs[pre+v*p+j]...)
				}
				for u := 0; u < gn; u++ {
					gr = append(gr, got.runs[pre+u*p+j]...)
				}
				if d := c.covered(pool, gr, fmt.Sprintf("%s > %s[%d]", path, pool[0].label(), pre+j+1)); d != "" {
					return d
				}
			}
		}
		for k := 0; pre+wn*p+k < len(want.runs); k++ {
			if d := c.runDiff(want.runs[pre+wn*p+k], got.runs[pre+gn*p+k], path, pre+gn*p+k); d != "" {
				return d
			}
		}
		return ""
	}
	n := len(want.runs)
	if len(got.runs) < n {
		n = len(got.runs)
	}
	for i := 0; i < n; i++ {
		if d := c.runDiff(want.runs[i], got.runs[i], path, i); d != "" {
			return d
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

// poolDiff compares a run of the screen with the frame's runs at the same
// place of a repeat: each element like one of them, whatever its kind.
func (c *cmp) poolDiff(pool, gr []*node, path string, i int) string {
	p := fmt.Sprintf("%s > %s[%d]", path, gr[0].label(), i+1)
	for _, g := range gr {
		first := ""
		for _, w := range pool {
			d := c.diff(w, g, p)
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
	return ""
}

// runDiff compares a run of the screen with the frame's: of one kind, each
// element of it like one of the frame's, and — what the frame draws, the
// screen has (Carrel's task 76) — no shorter than the frame's, each of the
// frame's elements like one of the screen's but for its states.
func (c *cmp) runDiff(wr, gr []*node, path string, i int) string {
	p := fmt.Sprintf("%s > %s[%d]", path, wr[0].label(), i+1)
	if wr[0].kind() != gr[0].kind() {
		return fmt.Sprintf("%s: the screen has <%s> where the frame has <%s>", p, gr[0].label(), wr[0].label())
	}
	if len(gr) < len(wr) {
		return fmt.Sprintf("%s: the screen has %d where the frame draws %d", p, len(gr), len(wr))
	}
	if d := c.covered(wr, gr, p); d != "" {
		return d
	}
	for _, g := range gr {
		first := ""
		for _, w := range wr {
			d := c.diff(w, g, p)
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
	return ""
}

// repeat finds, in the frame's runs, a block of p runs repeated at least
// twice, after the same first runs and before the same last ones as the
// screen's, where the screen repeats the same block any number of times.
// Runs are alike by kind, or — in a list of terms and values, where a
// value's classes vary (dd and dd.tc-mono) — by tag.
func repeat(want, got [][]*node) (pre, p, wn, gn int, ok bool) {
	for _, byTag := range []bool{false, true} {
		key := func(r []*node) string { return r[0].kind() }
		if byTag {
			key = func(r []*node) string { return r[0].name }
		}
		wk, gk := keys(want, key), keys(got, key)
		for p = 2; p <= 3; p++ {
			for pre = 0; pre+2*p <= len(wk); pre++ {
				base := wk[pre : pre+p]
				if distinct(base) < 2 {
					continue
				}
				if byTag && !termsAndValues(base) {
					continue
				}
				wn = reps(wk[pre:], base)
				if wn < 2 {
					continue
				}
				suf := wk[pre+wn*p:]
				if len(gk) < pre+len(suf)+p || !same(gk[:pre], wk[:pre]) || !same(gk[len(gk)-len(suf):], suf) {
					continue
				}
				mid := gk[pre : len(gk)-len(suf)]
				if len(mid)%p != 0 || reps(mid, base) != len(mid)/p {
					continue
				}
				if len(mid)/p == wn && same(gk, wk) && !byTag {
					return 0, 0, 0, 0, false // the same rows: the plain walk compares them
				}
				return pre, p, wn, len(mid) / p, true
			}
		}
	}
	return 0, 0, 0, 0, false
}

func keys(runs [][]*node, key func([]*node) string) []string {
	out := make([]string, len(runs))
	for i, r := range runs {
		out[i] = key(r)
	}
	return out
}

func reps(ks, base []string) int {
	n := 0
	for len(ks) >= len(base) && same(ks[:len(base)], base) {
		ks = ks[len(base):]
		n++
	}
	return n
}

func same(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// termsAndValues says whether a block is a list's terms and values (dt,
// dd): the one place a row's tag, not its classes, makes the repeat.
func termsAndValues(base []string) bool {
	for _, k := range base {
		if k != "dt" && k != "dd" {
			return false
		}
	}
	return true
}

func distinct(ks []string) int {
	m := map[string]bool{}
	for _, k := range ks {
		m[k] = true
	}
	return len(m)
}
