// Package fragment compares a widget's output with the fragment of the mockup
// frame it prints. The comparison is static: the same CSS on
// the same tree gives the same picture, so what is compared is what changes
// the picture — tags, classes, text, the values of fields and the data the
// script turns into style — and not accessibility attributes, which a widget
// adds on top of the frame.
package fragment

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Source is the mockup copy, relative to the module root.
const Source = "docs/visual/carrel-ui-mockups.html"

// Ref addresses one block in the mockup: the frame's data-frame, the block's
// class and which occurrence of that block in the frame it is, from 1, in
// document order. An element counts when the class is its block (see Block),
// not when it merely carries the class as a mixin.
type Ref struct {
	Frame string
	Class string
	N     int
}

func (r Ref) String() string { return fmt.Sprintf("%s .%s #%d", r.Frame, r.Class, r.N) }

// Mockup is the parsed mockup: its frames by data-frame, in document order.
type Mockup struct {
	Frames []*Frame
	byID   map[string]*Frame
	pre    map[string]bool // classes whose text keeps its white space
}

// Frame is one .frame: its id and the root of the drawing (the frame without
// its label, which is page chrome).
type Frame struct {
	ID   string
	Node *html.Node
	// Page says the frame is a page without a shell (a route of Р-16 whose
	// drawing's root is a plain tc-root): there, and only there, the root is
	// the block tc-root — the page widget's. Elsewhere tc-root is the box a
	// frame draws in, as a catalog cell is, and never a block.
	Page bool
}

// catalogOnly are frames without a mark that still go to the catalog, not to
// a route (Р-16, docs/development.md): the desktop app's user menu and the tray menu.
var catalogOnly = map[string]bool{"desktop-2": true, "desktop-8": true}

// Route says whether a frame is a route of the demo screens (Р-16): its
// label carries no mark (span.frag: компонент, фрагмент, список, меню,
// кандидат, отложена) and it is not one of the menus kept for the catalog.
func (f *Frame) Route() bool {
	if catalogOnly[f.ID] {
		return false
	}
	for c := f.Node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && hasClass(c, "frame-label") {
			for s := c.FirstChild; s != nil; s = s.NextSibling {
				if s.Type == html.ElementNode && hasClass(s, "frag") {
					return false
				}
			}
		}
	}
	return true
}

// blockOf is Block, with a page's root read as the block tc-root.
func (f *Frame) blockOf(n *html.Node) string {
	if f.Page {
		if r, err := f.Root(); err == nil && r == n {
			return "tc-root"
		}
	}
	return Block(n)
}

var (
	loadOnce sync.Once
	loaded   *Mockup
	loadErr  error
)

// Load parses the mockup under the module root once per process.
func Load(root string) (*Mockup, error) {
	loadOnce.Do(func() {
		data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(Source)))
		if err != nil {
			loadErr = err
			return
		}
		loaded, loadErr = Parse(string(data))
	})
	return loaded, loadErr
}

// FindRoot walks up from dir to the directory holding go.mod.
func FindRoot(dir string) (string, error) {
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("fragment: module root not found")
		}
		dir = parent
	}
}

var preRule = regexp.MustCompile(`([^{}]+)\{[^}]*white-space:\s*pre(?:-wrap|-line)?\b`)
var singleClass = regexp.MustCompile(`\.([a-zA-Z0-9_-]+)\s*$`)

// Parse reads the mockup markup.
func Parse(markup string) (*Mockup, error) {
	doc, err := html.Parse(strings.NewReader(markup))
	if err != nil {
		return nil, err
	}
	m := &Mockup{byID: map[string]*Frame{}, pre: map[string]bool{}}
	var walk func(n *html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.DataAtom == atom.Style && hasAttr(n, "data-tc") {
				m.readPre(text(n))
			}
			if id := attr(n, "data-frame"); id != "" && hasClass(n, "frame") {
				f := &Frame{ID: id, Node: n}
				if _, dup := m.byID[id]; dup {
					err = fmt.Errorf("fragment: duplicate data-frame %q", id)
					return
				}
				m.byID[id] = f
				m.Frames = append(m.Frames, f)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)
	if err != nil {
		return nil, err
	}
	if len(m.Frames) == 0 {
		return nil, errors.New("fragment: no frames")
	}
	for _, f := range m.Frames {
		if r, err := f.Root(); err == nil && f.Route() && hasClass(r, "tc-root") && Block(r) == "" {
			f.Page = true
		}
	}
	return m, nil
}

func (m *Mockup) readPre(css string) {
	for _, g := range preRule.FindAllStringSubmatch(css, -1) {
		for _, sel := range strings.Split(g[1], ",") {
			if c := singleClass.FindStringSubmatch(strings.TrimSpace(sel)); c != nil {
				m.pre[c[1]] = true
			}
		}
	}
}

// Frame returns the frame with this data-frame.
func (m *Mockup) Frame(id string) (*Frame, bool) {
	f, ok := m.byID[id]
	return f, ok
}

// Block returns the element Ref points at.
func (m *Mockup) Block(r Ref) (*html.Node, error) {
	f, ok := m.byID[r.Frame]
	if !ok {
		return nil, fmt.Errorf("fragment: no frame %q", r.Frame)
	}
	if r.N < 1 {
		return nil, fmt.Errorf("fragment: %s: occurrence counts from 1", r)
	}
	found := f.all(r.Class)
	if r.N > len(found) {
		return nil, fmt.Errorf("fragment: %s: the frame has %d .%s", r, len(found), r.Class)
	}
	return found[r.N-1], nil
}

// all returns the elements with class cls inside the drawing, in document order.
func (f *Frame) all(cls string) []*html.Node {
	var out []*html.Node
	var walk func(n *html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if hasClass(n, "frame-label") || chrome(n) {
				return
			}
			if hasClass(n, cls) && f.blockOf(n) == cls {
				out = append(out, n)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	for c := f.Node.FirstChild; c != nil; c = c.NextSibling {
		walk(c)
	}
	return out
}

// Node is a normalized element or text.
type Node struct {
	Tag      string // "" for text
	Classes  []string
	Attrs    []string // "name=value", sorted; only attributes that change the picture
	Text     string
	Children []*Node
}

// compared lists the attributes kept per tag; data-tc-* data is kept on any tag.
var compared = map[string][]string{
	"input":    {"type", "value", "placeholder", "disabled", "checked"},
	"textarea": {"placeholder", "disabled"},
	"button":   {"disabled"},
	"td":       {"colspan", "rowspan"},
	"th":       {"colspan", "rowspan"},
	"svg":      {"viewBox", "width", "height", "stroke", "stroke-width", "fill"},
	"use":      {"href"},
	"img":      {"src", "width", "height"},
}

var dataAttrs = []string{"data-tc-colour", "data-tc-width"}

// Normalize turns an element into the compared form.
func (m *Mockup) Normalize(n *html.Node) *Node {
	return m.norm(n, false)
}

func (m *Mockup) norm(n *html.Node, pre bool) *Node {
	out := &Node{Tag: n.Data}
	for _, c := range strings.Fields(attr(n, "class")) {
		if strings.HasPrefix(c, "mk-") {
			continue // the desktop window's frame is page chrome, not product
		}
		out.Classes = append(out.Classes, c)
		if m.pre[c] {
			pre = true
		}
	}
	sort.Strings(out.Classes)
	if n.DataAtom == atom.Textarea || n.DataAtom == atom.Pre {
		pre = true
	}
	for _, name := range compared[n.Data] {
		if v, ok := attrOK(n, name); ok {
			out.Attrs = append(out.Attrs, name+"="+v)
		}
	}
	for _, name := range dataAttrs {
		if v, ok := attrOK(n, name); ok {
			out.Attrs = append(out.Attrs, name+"="+v)
		}
	}
	sort.Strings(out.Attrs)
	for _, c := range children(n) {
		switch c.Type {
		case html.ElementNode:
			out.Children = append(out.Children, m.norm(c, pre))
		case html.TextNode:
			t := c.Data
			if !pre {
				t = strings.Join(strings.Fields(t), " ")
			}
			if t == "" || !pre && strings.TrimSpace(t) == "" {
				continue
			}
			out.Children = append(out.Children, &Node{Text: t})
		}
	}
	return out
}

// Own normalizes a block down to its own elements. A nested element that is
// another block becomes a place: its block name (or "*" when anonymous) and
// the mixins this container puts on it, without its variants, states,
// attributes or insides — those are that block's own forms.
func (m *Mockup) Own(n *html.Node, anonymous bool) *Node {
	return m.own(n, false, true, anonymous)
}

func (m *Mockup) own(n *html.Node, pre, root, anonymous bool) *Node {
	if b := Block(n); b != "" && !root {
		place := &Node{Tag: b}
		if anonymous {
			place.Tag = "*"
		}
		for _, cls := range strings.Fields(attr(n, "class")) {
			if !OwnClass(b, cls) {
				place.Classes = append(place.Classes, cls)
			}
		}
		sort.Strings(place.Classes)
		return place
	}
	out := m.norm(n, pre)
	out.Children = nil
	for _, c := range strings.Fields(attr(n, "class")) {
		if m.pre[c] {
			pre = true
		}
	}
	if n.DataAtom == atom.Textarea || n.DataAtom == atom.Pre {
		pre = true
	}
	for _, c := range children(n) {
		switch c.Type {
		case html.ElementNode:
			out.Children = append(out.Children, m.own(c, pre, false, anonymous))
		case html.TextNode:
			t := c.Data
			if !pre {
				t = strings.Join(strings.Fields(t), " ")
			}
			if t == "" || !pre && strings.TrimSpace(t) == "" {
				continue
			}
			out.Children = append(out.Children, &Node{Text: t})
		}
	}
	return out
}

// OwnClass says whether cls belongs to block itself on the block's root: the
// block, its variants --…, and its states is-*, except is-2nd, the rank a bar
// gives its items. Anything else on the root — another block (tc-push,
// tc-grow), another block's element (tc-bar__more, tc-sec__acts), is-2nd —
// is a mixin the container prints.
func OwnClass(block, cls string) bool {
	switch {
	case cls == block, strings.HasPrefix(cls, block+"--"), cls == Base[block]:
		return true
	case block != "tc-line" && block != "tc-stack" &&
		(cls == "tc-line" || cls == "tc-stack" || strings.HasPrefix(cls, "tc-line--") || strings.HasPrefix(cls, "tc-stack--")):
		return true // a named block laid out as a row or a column: its own layout
	case cls == "is-2nd":
		return false
	case strings.HasPrefix(cls, "is-"):
		return true
	}
	return false
}

// WithoutMixins drops the root's mixins: the form the block prints by itself.
func WithoutMixins(block string, n *Node) *Node {
	c := *n
	c.Classes = nil
	for _, cls := range n.Classes {
		if OwnClass(block, cls) {
			c.Classes = append(c.Classes, cls)
		}
	}
	return &c
}

// Fragment parses a widget's output: one element, or several siblings.
func Fragment(markup string) ([]*html.Node, error) {
	ctx := &html.Node{Type: html.ElementNode, Data: "div", DataAtom: atom.Div}
	nodes, err := html.ParseFragment(strings.NewReader(markup), ctx)
	if err != nil {
		return nil, err
	}
	var out []*html.Node
	for _, n := range nodes {
		if n.Type == html.ElementNode {
			out = append(out, n)
		} else if n.Type == html.TextNode && strings.TrimSpace(n.Data) != "" {
			return nil, fmt.Errorf("fragment: text %q outside an element", strings.TrimSpace(n.Data))
		}
	}
	return out, nil
}

// Compare checks a widget's rendered markup against the block Ref points at.
// It compares the block's own elements: a nested block is a place with its
// name and the mixins this block puts on it, and its insides are checked by
// its own records. The mixins on the frame's root come from the frame's
// container and are not the widget's to print. The error names the first
// path where the trees part.
func (m *Mockup) Compare(r Ref, rendered string) error {
	block, err := m.Block(r)
	if err != nil {
		return err
	}
	nodes, err := Fragment(rendered)
	if err != nil {
		return err
	}
	if len(nodes) != 1 {
		return fmt.Errorf("%s: the widget printed %d top elements, want 1", r, len(nodes))
	}
	cls := r.Class
	want := WithoutMixins(cls, m.Own(block, false))
	got := m.Own(nodes[0], false)
	if b := Block(nodes[0]); b != cls && !(cls == "tc-root" && b == "" && hasClass(nodes[0], "tc-root")) {
		return fmt.Errorf("%s: the widget prints block %q, want %q", r, b, cls)
	}
	if d := Diff(want, got); d != "" {
		return fmt.Errorf("%s: %s", r, d)
	}
	return nil
}

// Diff returns "" when the trees are equal, else the first difference with
// its path, written as the frame's side (want) against the widget's (got).
func Diff(want, got *Node) string {
	return diff(want, got, label(want))
}

func diff(want, got *Node, path string) string {
	if want.Tag == "" || got.Tag == "" {
		if want.Tag != got.Tag {
			return fmt.Sprintf("%s: frame has %s, widget has %s", path, describe(want), describe(got))
		}
		if want.Text != got.Text {
			return fmt.Sprintf("%s: text %q, widget %q", path, want.Text, got.Text)
		}
		return ""
	}
	if want.Tag != got.Tag {
		return fmt.Sprintf("%s: tag <%s>, widget <%s>", path, want.Tag, got.Tag)
	}
	if a, b := strings.Join(want.Classes, " "), strings.Join(got.Classes, " "); a != b {
		return fmt.Sprintf("%s: class %q, widget %q", path, a, b)
	}
	if a, b := strings.Join(want.Attrs, " "), strings.Join(got.Attrs, " "); a != b {
		return fmt.Sprintf("%s: attributes [%s], widget [%s]", path, a, b)
	}
	n := len(want.Children)
	if len(got.Children) < n {
		n = len(got.Children)
	}
	for i := 0; i < n; i++ {
		p := path + " > " + label(want.Children[i])
		if i > 0 {
			p = fmt.Sprintf("%s > %s[%d]", path, label(want.Children[i]), i+1)
		}
		if d := diff(want.Children[i], got.Children[i], p); d != "" {
			return d
		}
	}
	if len(want.Children) != len(got.Children) {
		if len(want.Children) > n {
			return fmt.Sprintf("%s: frame has %d children, widget %d; first missing: %s", path, len(want.Children), len(got.Children), describe(want.Children[n]))
		}
		return fmt.Sprintf("%s: frame has %d children, widget %d; first extra: %s", path, len(want.Children), len(got.Children), describe(got.Children[n]))
	}
	return ""
}

func label(n *Node) string {
	if n.Tag == "" {
		return "#text"
	}
	if len(n.Classes) == 0 {
		return n.Tag
	}
	return n.Tag + "." + strings.Join(n.Classes, ".")
}

func describe(n *Node) string {
	if n.Tag == "" {
		return fmt.Sprintf("text %q", n.Text)
	}
	return "<" + label(n) + ">"
}

func attr(n *html.Node, name string) string {
	v, _ := attrOK(n, name)
	return v
}

func attrOK(n *html.Node, name string) (string, bool) {
	for _, a := range n.Attr {
		if a.Key == name {
			return a.Val, true
		}
	}
	return "", false
}

func hasAttr(n *html.Node, name string) bool {
	_, ok := attrOK(n, name)
	return ok
}

func hasClass(n *html.Node, cls string) bool {
	for _, c := range strings.Fields(attr(n, "class")) {
		if c == cls {
			return true
		}
	}
	return false
}

// chrome says whether n is the desktop drawing's own: the window's title bar
// or status line, or a notification of the operating system. It is page
// chrome, not product — the system draws it, not the library — and is not
// seen with its contents.
func chrome(n *html.Node) bool {
	if n.Type != html.ElementNode {
		return false
	}
	for _, c := range strings.Fields(attr(n, "class")) {
		if c == "mk-titlebar" || c == "mk-status" || c == "mk-toast" {
			return true
		}
	}
	return false
}

// window says whether n is another element of the desktop drawing (mk-*):
// the window, a notification, the tray menu. It is transparent — its children
// stand in its place — because what it holds is the product's.
func window(n *html.Node) bool {
	if n.Type != html.ElementNode || chrome(n) {
		return false
	}
	for _, c := range strings.Fields(attr(n, "class")) {
		if strings.HasPrefix(c, "mk-") {
			return true
		}
	}
	return false
}

// children lists n's child nodes with the desktop drawing unwrapped — the
// window's children stand in its place — and the window's frame left out.
func children(n *html.Node) []*html.Node {
	var out []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		switch {
		case window(c):
			out = append(out, children(c)...)
		case chrome(c):
		default:
			out = append(out, c)
		}
	}
	return out
}

func text(n *html.Node) string {
	var b strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			b.WriteString(c.Data)
		}
	}
	return b.String()
}

// Root returns the root of the frame's drawing: its one element besides the
// label, the tc-root.
func (f *Frame) Root() (*html.Node, error) {
	var out []*html.Node
	for c := f.Node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && !hasClass(c, "frame-label") {
			out = append(out, c)
		}
	}
	if len(out) != 1 {
		return nil, fmt.Errorf("fragment: frame %s has %d drawings, want 1", f.ID, len(out))
	}
	return out[0], nil
}

// CompareFrame compares a whole frame with a screen printed for it: the
// same tree, classes, compared attributes and text, top to bottom. It says which widget was chosen and in what order, which the check
// of one widget against its fragment does not.
func (m *Mockup) CompareFrame(id, rendered string) error {
	f, ok := m.byID[id]
	if !ok {
		return fmt.Errorf("fragment: no frame %q", id)
	}
	root, err := f.Root()
	if err != nil {
		return err
	}
	nodes, err := Fragment(rendered)
	if err != nil {
		return err
	}
	if len(nodes) != 1 {
		return fmt.Errorf("%s: the screen printed %d top elements, want 1", id, len(nodes))
	}
	if d := Diff(m.Normalize(root), m.Normalize(nodes[0])); d != "" {
		return fmt.Errorf("%s: %s", id, d)
	}
	return nil
}

// Diffs is Diff that goes on past a difference: every differing place, up
// to max. A place whose tag differs is reported and not entered; elsewhere
// the children are compared pairwise.
func Diffs(want, got *Node, max int) []string {
	var out []string
	var walk func(want, got *Node, path string)
	walk = func(want, got *Node, path string) {
		if len(out) >= max {
			return
		}
		if want.Tag == "" || got.Tag == "" || want.Tag != got.Tag {
			if d := diff(want, got, path); d != "" {
				out = append(out, d)
			}
			return
		}
		if a, b := strings.Join(want.Classes, " "), strings.Join(got.Classes, " "); a != b {
			out = append(out, fmt.Sprintf("%s: class %q, widget %q", path, a, b))
		} else if a, b := strings.Join(want.Attrs, " "), strings.Join(got.Attrs, " "); a != b {
			out = append(out, fmt.Sprintf("%s: attributes [%s], widget [%s]", path, a, b))
		}
		n := min(len(want.Children), len(got.Children))
		if len(want.Children) != len(got.Children) {
			d := fmt.Sprintf("%s: frame has %d children, widget %d", path, len(want.Children), len(got.Children))
			if len(want.Children) > n {
				d += "; first missing: " + describe(want.Children[n])
			} else {
				d += "; first extra: " + describe(got.Children[n])
			}
			out = append(out, d)
			return
		}
		for i := 0; i < n; i++ {
			p := fmt.Sprintf("%s > %s[%d]", path, label(want.Children[i]), i+1)
			walk(want.Children[i], got.Children[i], p)
		}
	}
	walk(want, got, label(want))
	return out
}

// FrameDiffs lists up to max differences between a frame and its screen.
func (m *Mockup) FrameDiffs(id, rendered string, max int) ([]string, error) {
	f, ok := m.byID[id]
	if !ok {
		return nil, fmt.Errorf("fragment: no frame %q", id)
	}
	root, err := f.Root()
	if err != nil {
		return nil, err
	}
	nodes, err := Fragment(rendered)
	if err != nil {
		return nil, err
	}
	if len(nodes) != 1 {
		return []string{fmt.Sprintf("the screen printed %d top elements, want 1", len(nodes))}, nil
	}
	return Diffs(m.Normalize(root), m.Normalize(nodes[0]), max), nil
}
