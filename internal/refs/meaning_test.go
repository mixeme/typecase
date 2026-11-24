package refs_test

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"carrel.dev/typecase/internal/refs"
)

// meaningFields are the fields that give an element its meaning in HTML
// (Р-25): where it goes, what it sends.
var meaningFields = map[string]bool{"Href": true, "Form": true, "Name": true, "Action": true}

// clone is a deep copy of v: filling it must not touch the records the other
// tests compare with their frames.
func clone(v reflect.Value) reflect.Value {
	switch v.Kind() {
	case reflect.Pointer:
		if v.IsNil() {
			return v
		}
		c := reflect.New(v.Elem().Type())
		c.Elem().Set(clone(v.Elem()))
		return c
	case reflect.Interface:
		if v.IsNil() {
			return v
		}
		c := reflect.New(v.Type()).Elem()
		c.Set(clone(v.Elem()))
		return c
	case reflect.Slice:
		if v.IsNil() {
			return v
		}
		c := reflect.MakeSlice(v.Type(), v.Len(), v.Len())
		for i := 0; i < v.Len(); i++ {
			c.Index(i).Set(clone(v.Index(i)))
		}
		return c
	case reflect.Struct:
		c := reflect.New(v.Type()).Elem()
		c.Set(v)
		for i := 0; i < v.NumField(); i++ {
			if c.Field(i).CanSet() {
				c.Field(i).Set(clone(v.Field(i)))
			}
		}
		return c
	}
	return v
}

// fill sets every meaning field reachable from v to a marker, so what the
// widget can carry, it prints.
func fill(v reflect.Value) {
	switch v.Kind() {
	case reflect.Pointer, reflect.Interface:
		if !v.IsNil() {
			e := v.Elem()
			if v.Kind() == reflect.Interface {
				c := reflect.New(e.Type()).Elem()
				c.Set(e)
				fill(c)
				v.Set(c)
				return
			}
			fill(e)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			if !v.Type().Field(i).IsExported() || !f.CanSet() {
				continue
			}
			if meaningFields[v.Type().Field(i).Name] && f.Kind() == reflect.String && f.String() == "" {
				f.SetString("m")
				continue
			}
			fill(f)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			fill(v.Index(i))
		}
	}
}

func classOf(n *html.Node) string {
	for _, a := range n.Attr {
		if a.Key == "class" {
			return a.Val
		}
	}
	return ""
}

func attr(n *html.Node, k string) (string, bool) {
	for _, a := range n.Attr {
		if a.Key == k {
			return a.Val, true
		}
	}
	return "", false
}

// mute is what needs no meaning of its own: inside a menu or a sheet that is
// a drawing of an open state, a control that only shows a state.
func silent(n *html.Node) bool {
	c := " " + classOf(n) + " "
	for _, s := range []string{" tc-rubric ", " tc-menu__sep ", " tc-tickrow__tick "} {
		if strings.Contains(c, s) && n.Data != "button" {
			return true
		}
	}
	return false
}

// gaps lists the elements of a record that look pressable or take input and
// still carry no meaning once every meaning field is set.
func gaps(markup string) []string {
	nodes, err := html.ParseFragment(strings.NewReader(markup), &html.Node{Type: html.ElementNode, Data: "body", DataAtom: atom.Body})
	if err != nil {
		return []string{err.Error()}
	}
	var out []string
	var walk func(n *html.Node, inForm, inLinks bool)
	walk = func(n *html.Node, inForm, inLinks bool) {
		if n.Type == html.ElementNode {
			c := classOf(n)
			label := n.Data
			if c != "" {
				label += "." + strings.Join(strings.Fields(c), ".")
			}
			if strings.Contains(" "+c+" ", " tc-doc__files ") || strings.Contains(" "+c+" ", " tc-docmeta__sec ") {
				inLinks = true
			}
			switch n.Data {
			case "form":
				inForm = true
			case "a":
				_, href := attr(n, "href")
				if v, _ := attr(n, "aria-disabled"); !href && v != "true" {
					out = append(out, label+" without href")
				}
			case "button":
				t, _ := attr(n, "type")
				_, form := attr(n, "form")
				_, adds := attr(n, "data-tc-add")
				_, opens := attr(n, "data-tc-open")
				_, toggles := attr(n, "data-tc-toggle") // the rail as a panel (Carrel's task 274)
				_, off := attr(n, "disabled")
				if !(t == "submit" && (inForm || form)) && !adds && !opens && !toggles && !off {
					out = append(out, label+" is a button with nothing to send")
				}
			case "input", "select", "textarea":
				t, _ := attr(n, "type")
				_, name := attr(n, "name")
				_, off := attr(n, "disabled")
				if !name && t != "hidden" && !off {
					out = append(out, label+" without name")
				}
			default:
				if (linkLike(c) || inLinks && strings.Contains(" "+c+" ", " tc-dim--accent ")) && !silent(n) {
					out = append(out, label+" looks like a link and is none")
				}
			}
		}
		for k := n.FirstChild; k != nil; k = k.NextSibling {
			walk(k, inForm, inLinks)
		}
	}
	for _, n := range nodes {
		walk(n, false, false)
	}
	return out
}

// TestEveryControlHasMeaning (task 39 of Carrel's server): every element of
// every record that is pressed, followed or typed into can carry its meaning.
func TestEveryControlHasMeaning(t *testing.T) {
	byGap := map[string][]string{}
	for _, e := range refs.Entries {
		v := reflect.New(reflect.TypeOf(e.Fixture)).Elem()
		v.Set(clone(reflect.ValueOf(e.Fixture)))
		fill(v)
		r, ok := v.Interface().(interface{ Render() string })
		if !ok {
			continue
		}
		for _, g := range gaps(r.Render()) {
			byGap[fmt.Sprintf("%T: %s", e.Fixture, g)] = append(byGap[fmt.Sprintf("%T: %s", e.Fixture, g)], e.Ref.Frame)
		}
	}
	var keys []string
	for k := range byGap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fr := byGap[k]
		t.Errorf("%s (%d, first %s)", k, len(fr), fr[0])
	}
}

// linkLike says whether a class list draws words as a link: a part named a
// link. Words in the accent are a link only where a list of files or of
// related records stands (the accent is also a good outcome).
func linkLike(c string) bool {
	for _, t := range strings.Fields(c) {
		if strings.HasSuffix(t, "__link") {
			return true
		}
	}
	return false
}
