package main

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

var voidHTML = map[string]bool{
	"area": true, "base": true, "br": true, "col": true, "embed": true,
	"hr": true, "img": true, "input": true, "link": true, "meta": true,
	"param": true, "source": true, "track": true, "wbr": true,
}

type node struct {
	tag                  string
	class                []string
	openStart, openEnd   int
	innerStart, innerEnd int
	closeStart, closeEnd int
	parent               *node
	children             []*node
	void                 bool
}

func (n *node) outerEnd() int {
	if n.void || n.closeEnd < 0 {
		return n.openEnd + 1
	}
	return n.closeEnd + 1
}

func hasClass(n *node, c string) bool {
	if n == nil {
		return false
	}
	for _, t := range n.class {
		if t == c {
			return true
		}
	}
	return false
}

func hasAnyClass(n *node, names ...string) bool {
	for _, name := range names {
		if hasClass(n, name) {
			return true
		}
	}
	return false
}

func parseHTML(s string) *node {
	root := &node{tag: "#root", innerStart: 0, innerEnd: len(s), closeEnd: len(s) - 1}
	stack := []*node{root}
	i := 0
	for i < len(s) {
		if strings.HasPrefix(s[i:], "<!--") {
			j := strings.Index(s[i+4:], "-->")
			if j < 0 {
				break
			}
			i += 4 + j + 3
			continue
		}
		if s[i] != '<' {
			i++
			continue
		}
		if i+1 < len(s) && s[i+1] == '/' {
			end := scanTagEnd(s, i)
			if end < 0 {
				break
			}
			name := strings.ToLower(strings.TrimSpace(s[i+2 : end]))
			if sp := strings.IndexFunc(name, unicode.IsSpace); sp >= 0 {
				name = name[:sp]
			}
			for len(stack) > 1 && stack[len(stack)-1].tag != name {
				top := stack[len(stack)-1]
				top.void = true
				top.innerEnd = top.openEnd + 1
				top.closeStart, top.closeEnd = -1, -1
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 1 && stack[len(stack)-1].tag == name {
				top := stack[len(stack)-1]
				top.innerEnd = i
				top.closeStart = i
				top.closeEnd = end
				stack = stack[:len(stack)-1]
			}
			i = end + 1
			continue
		}
		if i+1 < len(s) && (s[i+1] == '!' || s[i+1] == '?') {
			end := strings.IndexByte(s[i:], '>')
			if end < 0 {
				break
			}
			i += end + 1
			continue
		}
		end := scanTagEnd(s, i)
		if end < 0 {
			break
		}
		inner := strings.TrimSpace(s[i+1 : end])
		selfClose := strings.HasSuffix(inner, "/")
		name, class := parseStartTag(s, i, end)
		name = strings.ToLower(name)
		parent := stack[len(stack)-1]
		n := &node{
			tag: name, class: class,
			openStart: i, openEnd: end,
			innerStart: end + 1,
			parent:     parent,
			closeStart: -1, closeEnd: -1,
		}
		parent.children = append(parent.children, n)
		if name == "style" || name == "script" {
			close := strings.Index(strings.ToLower(s[end+1:]), "</"+name+">")
			if close < 0 {
				break
			}
			n.innerStart = end + 1
			n.innerEnd = end + 1 + close
			n.closeStart = end + 1 + close
			n.closeEnd = n.closeStart + len(name) + 2
			i = n.closeEnd + 1
			continue
		}
		i = end + 1
		if selfClose || voidHTML[name] {
			n.void = true
			n.innerEnd = end + 1
			continue
		}
		stack = append(stack, n)
	}
	for len(stack) > 1 {
		top := stack[len(stack)-1]
		top.void = true
		top.innerEnd = top.openEnd + 1
		stack = stack[:len(stack)-1]
	}
	return root
}

func parseStartTag(html string, start, end int) (name string, class []string) {
	raw := html[start+1 : end]
	if strings.HasSuffix(strings.TrimSpace(raw), "/") {
		raw = strings.TrimSpace(strings.TrimSpace(raw)[:len(strings.TrimSpace(raw))-1])
	}
	i := 0
	for i < len(raw) {
		r, n := utf8.DecodeRuneInString(raw[i:])
		if unicode.IsSpace(r) {
			break
		}
		i += n
	}
	name = raw[:i]
	k := i
	for k < len(raw) {
		for k < len(raw) && unicode.IsSpace(rune(raw[k])) {
			k++
		}
		if k >= len(raw) || raw[k] == '/' {
			break
		}
		keyStart := k
		for k < len(raw) && raw[k] != '=' && !unicode.IsSpace(rune(raw[k])) && raw[k] != '/' {
			k++
		}
		key := strings.ToLower(raw[keyStart:k])
		for k < len(raw) && unicode.IsSpace(rune(raw[k])) {
			k++
		}
		if k >= len(raw) || raw[k] != '=' {
			continue
		}
		k++
		for k < len(raw) && unicode.IsSpace(rune(raw[k])) {
			k++
		}
		if k >= len(raw) {
			break
		}
		var val string
		if raw[k] == '"' || raw[k] == '\'' {
			q := raw[k]
			k++
			p := strings.IndexByte(raw[k:], q)
			if p < 0 {
				break
			}
			val = raw[k : k+p]
			k += p + 1
		} else {
			p := k
			for p < len(raw) && !unicode.IsSpace(rune(raw[p])) && raw[p] != '/' {
				p++
			}
			val = raw[k:p]
			k = p
		}
		if key == "class" {
			class = strings.Fields(val)
		}
	}
	return name, class
}

func scanTagEnd(html string, start int) int {
	var quote byte
	for i := start + 1; i < len(html); i++ {
		c := html[i]
		if quote != 0 {
			if c == quote {
				quote = 0
			}
			continue
		}
		if c == '"' || c == '\'' {
			quote = c
			continue
		}
		if c == '>' {
			return i
		}
	}
	return -1
}

func attrValue(html string, n *node, key string) (string, bool) {
	raw := html[n.openStart+1 : n.openEnd]
	if strings.HasSuffix(strings.TrimSpace(raw), "/") {
		raw = strings.TrimSpace(raw[:len(strings.TrimSpace(raw))-1])
	}
	k := 0
	for k < len(raw) && !unicode.IsSpace(rune(raw[k])) {
		k++
	}
	for k < len(raw) {
		for k < len(raw) && unicode.IsSpace(rune(raw[k])) {
			k++
		}
		if k >= len(raw) || raw[k] == '/' {
			break
		}
		keyStart := k
		for k < len(raw) && raw[k] != '=' && !unicode.IsSpace(rune(raw[k])) && raw[k] != '/' {
			k++
		}
		got := strings.ToLower(raw[keyStart:k])
		for k < len(raw) && unicode.IsSpace(rune(raw[k])) {
			k++
		}
		if k >= len(raw) || raw[k] != '=' {
			if got == key {
				return "", true
			}
			continue
		}
		k++
		for k < len(raw) && unicode.IsSpace(rune(raw[k])) {
			k++
		}
		if k >= len(raw) {
			break
		}
		var val string
		if raw[k] == '"' || raw[k] == '\'' {
			q := raw[k]
			k++
			p := strings.IndexByte(raw[k:], q)
			if p < 0 {
				break
			}
			val = raw[k : k+p]
			k += p + 1
		} else {
			p := k
			for p < len(raw) && !unicode.IsSpace(rune(raw[p])) && raw[p] != '/' {
				p++
			}
			val = raw[k:p]
			k = p
		}
		if got == key {
			return val, true
		}
	}
	return "", false
}

func openTag(html string, n *node) string {
	return html[n.openStart : n.openEnd+1]
}

func closeTag(html string, n *node) string {
	if n.void || n.closeStart < 0 {
		return ""
	}
	return html[n.closeStart : n.closeEnd+1]
}
