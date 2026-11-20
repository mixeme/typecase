// Package framegen translates a frame of the mockup into the Go expression
// of its screen: the widgets and the frame's data, top to bottom.
// cmd/frame2go prints it; internal/screens is its output, held to
// the frames by the static comparison (TestFramesStatic) and to the
// generator by TestGenerated.
//
// A translator per widget family: text.go (running text, marks, badges),
// controls.go (buttons, fields, forms, lines, rubrics), blocks.go (cards,
// tables, messages, documents), shell.go (the shell, top bar, rail, menus,
// sheets, dialogs, rows, pages). A variant is found by the element's classes
// in the widget package's own class() tables (enums.go), not by a list of
// the generator's. An element it does not know stops the frame with its
// name: a gap, never a guess.
//
// It parses with its own tokenizer (el.go), not internal/fragment: it needs
// the text runs as written, spaces between inline elements included, and
// whether an attribute was written bare — a tree builder gives neither.
//
// What it does not know: where the frame's data ends and markup begins. A
// list's rows are rows to it, not a loop over a model; that is the screen
// author's part (docs/guide/screen.md).
package framegen

import (
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

// El is an element of a frame: its tag, attributes and children in order —
// text runs as strings, elements as *El.
type El struct {
	Tag   string
	attrs map[string]string
	has   map[string]bool // attributes present, with or without a value
	bare  map[string]bool // attributes written without a value
	Kids  []any
}

var void = map[string]bool{"input": true, "br": true, "img": true, "use": true, "meta": true, "link": true, "hr": true}

// Parse reads a frame's drawing into a tree and returns its first element.
// It keeps what a browser would drop — whitespace-only text — because the
// spaces between inline elements are part of the drawing.
func Parse(markup string) *El {
	root := &El{Tag: "root"}
	stack := []*El{root}
	z := html.NewTokenizer(strings.NewReader(markup))
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			for _, k := range root.Kids {
				if e, ok := k.(*El); ok {
					return e
				}
			}
			return nil
		case html.TextToken:
			top := stack[len(stack)-1]
			t := string(z.Text())
			if n := len(top.Kids); n > 0 {
				if s, ok := top.Kids[n-1].(string); ok {
					top.Kids[n-1] = s + t
					continue
				}
			}
			top.Kids = append(top.Kids, t)
		case html.StartTagToken, html.SelfClosingTagToken:
			raw := string(z.Raw())
			name, hasAttr := z.TagName()
			e := &El{Tag: string(name), attrs: map[string]string{}, has: map[string]bool{}, bare: map[string]bool{}}
			for hasAttr {
				var k, v []byte
				k, v, hasAttr = z.TagAttr()
				e.attrs[string(k)] = string(v)
				e.has[string(k)] = true
				if v == nil || len(v) == 0 {
					e.bare[string(k)] = !valuedAttr(raw, string(k))
				}
			}
			top := stack[len(stack)-1]
			top.Kids = append(top.Kids, e)
			if tt == html.StartTagToken && !void[e.Tag] {
				stack = append(stack, e)
			}
		case html.EndTagToken:
			name, _ := z.TagName()
			t := string(name)
			if void[t] {
				continue
			}
			for len(stack) > 1 && stack[len(stack)-1].Tag != t {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
			}
		}
	}
}

// valuedAttr says whether the raw start tag gives attribute k a value
// ("k=…"), empty or not.
func valuedAttr(raw, k string) bool {
	low := strings.ToLower(raw)
	for i := 0; ; {
		j := strings.Index(low[i:], k)
		if j < 0 {
			return false
		}
		j += i
		i = j + len(k)
		if j == 0 || !isAttrSep(low[j-1]) {
			continue
		}
		rest := strings.TrimLeft(low[i:], " \t\n\r\f")
		if strings.HasPrefix(rest, "=") {
			return true
		}
	}
}

func isAttrSep(c byte) bool { return strings.IndexByte(" \t\n\r\f\"'", c) >= 0 }

// Valued says whether the attribute is present with a value, as a parser
// that gives a bare attribute no value sees it.
func (e *El) Valued(name string) bool { return e.has[name] && !e.bare[name] }

// Attr is an attribute's value, "" when absent or empty.
func (e *El) Attr(name string) string { return e.attrs[name] }

// HasAttr says whether the attribute is present.
func (e *El) HasAttr(name string) bool { return e.has[name] }

// Cls is the element's classes.
func (e *El) Cls() []string { return strings.Fields(e.attrs["class"]) }

// Has says whether the element has class c.
func (e *El) Has(c string) bool {
	for _, x := range e.Cls() {
		if x == c {
			return true
		}
	}
	return false
}

// Els is the element's child elements.
func (e *El) Els() []*El {
	var out []*El
	for _, k := range e.Kids {
		if x, ok := k.(*El); ok {
			out = append(out, x)
		}
	}
	return out
}

// Kid is the first child element with class c, or nil.
func (e *El) Kid(c string) *El {
	for _, k := range e.Els() {
		if k.Has(c) {
			return k
		}
	}
	return nil
}

// Text is all the element's text, its white space collapsed.
func (e *El) Text() string {
	var b strings.Builder
	var walk func(e *El)
	walk = func(e *El) {
		for _, k := range e.Kids {
			switch v := k.(type) {
			case string:
				b.WriteString(v)
			case *El:
				walk(v)
			}
		}
	}
	walk(e)
	return norm(b.String())
}

// OwnText is the element's own text runs, joined by a space and collapsed.
func (e *El) OwnText() string {
	var parts []string
	for _, k := range e.Kids {
		if s, ok := k.(string); ok {
			parts = append(parts, s)
		}
	}
	return norm(strings.Join(parts, " "))
}

// TextRaw is all the element's text as written.
func (e *El) TextRaw() string {
	var b strings.Builder
	for _, k := range e.Kids {
		switch v := k.(type) {
		case string:
			b.WriteString(v)
		case *El:
			b.WriteString(v.TextRaw())
		}
	}
	return b.String()
}

// norm collapses white space to one space and trims it. White space is
// HTML's (space, tab, line feeds): a no-break space is a letter of the text,
// as it is to the browser — it holds a word to the next (s-davloom-1's
// «server&nbsp;·»), and losing it moved the line break on a phone.
func norm(s string) string { return strings.Join(strings.FieldsFunc(s, htmlSpace), " ") }

// htmlSpace is the white space HTML collapses; a no-break space is not.
func htmlSpace(r rune) bool {
	return r == 0x20 || r == 0x09 || r == 0x0a || r == 0x0c || r == 0x0d
}

// collapse turns every run of white space into one space, keeping the ends.
func collapse(s string) string {
	var b strings.Builder
	space := false
	for _, r := range s {
		if htmlSpace(r) {
			if !space {
				b.WriteByte(' ')
			}
			space = true
			continue
		}
		space = false
		b.WriteRune(r)
	}
	return b.String()
}

// q is a Go string literal of s.
func q(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `"`, `\"`)
	s = strings.ReplaceAll(s, "\n", `\n`)
	s = strings.ReplaceAll(s, "\t", `\t`)
	return `"` + s + `"`
}

func itoa(n int) string { return strconv.Itoa(n) }

func join(xs []string, sep string) string { return strings.Join(xs, sep) }

// fail stops the translation of a frame: the frame has something the
// generator does not know. The frame is then a gap.
type fail struct{ why string }

func stop(why string) { panic(fail{why}) }
