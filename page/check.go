package page

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Problem is a name one part of a page gives another that does not hold: a
// button that submits a form the page has not, a file dropped on a panel
// whose form takes no file, a link that asks a zone no part carries, an id
// given twice. Such a page renders and looks right; what breaks is what is
// pressed — the script finds nothing and does nothing, or the plain way goes.
type Problem struct {
	Attr string // the attribute that names: "form", "data-tc-drop", "id"
	Name string // the name it gives
	By   string // the element that names it: its tag and first class, «button.tc-btn»
	Why  string // what is wrong: «names nothing», «names a div, not a form»
}

func (p Problem) String() string {
	return fmt.Sprintf("%s %s=%q: %s", p.By, p.Attr, p.Name, p.Why)
}

// Check prints the document and reads it as the browser and the library's
// script will: every name one part gives another — an id, a form's id, a
// zone's name, a field's name — holds, and names what it has to. A page of
// parts that each render right can still be put together wrong; Check is
// what says so before a person presses it. The consumer runs it in its
// tests over the pages it prints.
func (d Document) Check() ([]Problem, error) {
	s, err := d.Render()
	if err != nil {
		return nil, err
	}
	return Check(s)
}

// Check reads a whole page printed by the consumer — a Document's, or one it
// put together itself — as Document.Check does.
func Check(page string) ([]Problem, error) {
	root, err := html.Parse(strings.NewReader(page))
	if err != nil {
		return nil, err
	}
	c := checker{ids: map[string]*html.Node{}, twice: map[string]bool{}, zones: map[string]bool{}}
	c.walk(root, false)
	c.look(root, false)
	return c.out, nil
}

type checker struct {
	ids   map[string]*html.Node
	twice map[string]bool
	zones map[string]bool
	out   []Problem
}

func attrOf(n *html.Node, name string) (string, bool) {
	for _, a := range n.Attr {
		if a.Key == name {
			return a.Val, true
		}
	}
	return "", false
}

// children goes into a template's content as well: html.Parse puts it
// under the template element.
func children(n *html.Node) []*html.Node {
	var out []*html.Node
	for k := n.FirstChild; k != nil; k = k.NextSibling {
		out = append(out, k)
	}
	return out
}

// walk takes the ids and the zones of the page. What lies in a template is
// not on the page until the script copies it: its ids and zones are not
// counted, its names are checked against the page (look).
func (c *checker) walk(n *html.Node, inTemplate bool) {
	if n.Type == html.ElementNode && !inTemplate {
		if id, ok := attrOf(n, "id"); ok && id != "" {
			if c.ids[id] != nil && !c.twice[id] {
				c.twice[id] = true
				c.add(n, "id", id, "the id is given twice: the script finds the first only")
			}
			if c.ids[id] == nil {
				c.ids[id] = n
			}
		}
		if z, ok := attrOf(n, "data-tc-zone"); ok && z != "" && n.DataAtom != atom.Template {
			c.zones[z] = true
		}
	}
	for _, k := range children(n) {
		c.walk(k, inTemplate || (n.Type == html.ElementNode && n.DataAtom == atom.Template))
	}
}

func by(n *html.Node) string {
	s := n.Data
	if cl, ok := attrOf(n, "class"); ok {
		if f := strings.Fields(cl); len(f) > 0 {
			s += "." + f[0]
		}
	}
	return s
}

func (c *checker) add(n *html.Node, attr, name, why string) {
	c.out = append(c.out, Problem{Attr: attr, Name: name, By: by(n), Why: why})
}

// what an id names has to be
var idNames = map[string]func(*checker, *html.Node) string{
	"form":              tag("form"),
	"list":              tag("datalist"),
	"for":               labelable,
	"data-tc-progress":  tag("form"),
	"data-tc-drop":      fileForm,
	"data-tc-paste":     fileForm,
	"data-tc-add":       tag("template"),
	"data-tc-open":      anything,
	"data-tc-save-from": anything,
	"data-tc-copy-from": anything,
	"data-tc-preview":   anything,
	"data-tc-list":      anything,
	"aria-controls":     anything,
	"aria-labelledby":   anything,
	"aria-describedby":  anything,
	"data-tc-swap":      anything,
	"data-tc-recount":   anything,
}

func tag(want string) func(*checker, *html.Node) string {
	return func(_ *checker, n *html.Node) string {
		if n.Data != want {
			return "names a " + n.Data + ", not a " + want
		}
		return ""
	}
}

func anything(*checker, *html.Node) string { return "" }

func labelable(_ *checker, n *html.Node) string {
	switch n.Data {
	case "input", "select", "textarea", "button", "output", "meter", "progress":
		return ""
	}
	return "names a " + n.Data + ", which a label does not name"
}

// fileForm: a dropped or pasted file goes into the first file field of the
// form — one inside it or one that names it.
func fileForm(c *checker, n *html.Node) string {
	if n.Data != "form" {
		return "names a " + n.Data + ", not a form"
	}
	id, _ := attrOf(n, "id")
	found := false
	var find func(*html.Node, bool)
	find = func(k *html.Node, inside bool) {
		if found {
			return
		}
		if k.Type == html.ElementNode && k.Data == "input" {
			if t, _ := attrOf(k, "type"); t == "file" {
				f, has := attrOf(k, "form")
				if (has && f == id) || (!has && inside) {
					found = true
				}
			}
		}
		for _, x := range children(k) {
			find(x, inside || k == n)
		}
	}
	root := n
	for root.Parent != nil {
		root = root.Parent
	}
	find(root, false)
	if !found {
		return "names a form with no file field: the file goes nowhere"
	}
	return ""
}

// formOf is the form whose fields the script searches for a name: the
// element's own form, or the page.
func (c *checker) formOf(n *html.Node) *html.Node {
	if f, ok := attrOf(n, "form"); ok {
		if el := c.ids[f]; el != nil && el.Data == "form" {
			return el
		}
		return nil
	}
	for p := n.Parent; p != nil; p = p.Parent {
		if p.Type == html.ElementNode && p.Data == "form" {
			return p
		}
	}
	return nil
}

// named: a field with the name under scope (the page when scope is nil),
// as the script's querySelectorAll('[name=…]') finds it.
func named(scope *html.Node, name string) bool {
	if scope == nil {
		return false
	}
	found := false
	var find func(*html.Node)
	find = func(k *html.Node) {
		if found {
			return
		}
		if k.Type == html.ElementNode && k.DataAtom != atom.Template {
			if v, ok := attrOf(k, "name"); ok && v == name {
				found = true
				return
			}
		}
		if k.Type == html.ElementNode && k.DataAtom == atom.Template {
			return
		}
		for _, x := range children(k) {
			find(x)
		}
	}
	find(scope)
	return found
}

func (c *checker) look(n *html.Node, inTemplate bool) {
	if n.Type == html.ElementNode {
		c.names(n, inTemplate)
	}
	for _, k := range children(n) {
		c.look(k, inTemplate || (n.Type == html.ElementNode && n.DataAtom == atom.Template))
	}
}

func (c *checker) names(n *html.Node, inTemplate bool) {
	for _, a := range n.Attr {
		kind, ok := idNames[a.Key]
		// a file field's shows is the id of the line beside it; an
		// option's shows are names of fields (below)
		if a.Key == "data-tc-shows" && n.Data == "input" {
			kind, ok = tag("input"), true
		}
		if !ok {
			continue
		}
		for _, id := range strings.Fields(a.Val) {
			if strings.Contains(id, "{n}") {
				continue // a template's number, given when the row is added
			}
			el := c.ids[id]
			if el == nil {
				c.add(n, a.Key, id, "names nothing on the page")
				continue
			}
			if why := kind(c, el); why != "" {
				c.add(n, a.Key, id, why)
			}
		}
	}
	// names of fields, searched in the element's form
	for _, key := range []string{"data-tc-disables", "data-tc-follows", "data-tc-shows", "data-tc-from", "data-tc-changed", "data-tc-day", "data-tc-slug-of"} {
		v, ok := attrOf(n, key)
		if !ok || (key == "data-tc-shows" && n.Data == "input") {
			continue
		}
		scope := n
		if n.Data == "option" {
			for scope = n.Parent; scope != nil && scope.Data != "select"; scope = scope.Parent {
			}
			if scope == nil {
				scope = n
			}
		}
		form := c.formOf(scope)
		if form == nil {
			form = rootOf(n)
		}
		for _, name := range strings.Fields(v) {
			if strings.Contains(name, "{n}") || inTemplate {
				continue
			}
			if !named(form, name) {
				c.add(n, key, name, "names no field of its form")
			}
		}
	}
	// data-tc-ranges-of names the file a field kept a step before, on a
	// page the zone's answer replaced: it stands on no page of its own form
	// (s-backup-20 after s-backup-16), and is not checked
	// a submit button of no form does nothing when pressed: a sheet whose
	// body holds a form of its own prints no form of its Action
	if submits(n) && !inTemplate && c.formOf(n) == nil {
		if _, named := attrOf(n, "form"); !named {
			c.add(n, "type", "submit", "submits no form: pressed, it does nothing")
		}
	}
	if z, ok := attrOf(n, "data-tc-target"); ok && z != "" && !c.zones[z] {
		c.add(n, "data-tc-target", z, "names no zone on the page: what is pressed goes the plain way")
	}
	if z, ok := attrOf(n, "data-tc-address-zone"); ok && z != "" && !c.zones[z] {
		c.add(n, "data-tc-address-zone", z, "names no zone on the page")
	}
}

func rootOf(n *html.Node) *html.Node {
	for n.Parent != nil {
		n = n.Parent
	}
	return n
}

// submits: a button that sends its form when pressed — a button whose type
// is submit or not given, an input of type submit or image.
func submits(n *html.Node) bool {
	t, _ := attrOf(n, "type")
	switch n.Data {
	case "button":
		return t == "" || strings.EqualFold(t, "submit")
	case "input":
		return strings.EqualFold(t, "submit") || strings.EqualFold(t, "image")
	}
	return false
}
