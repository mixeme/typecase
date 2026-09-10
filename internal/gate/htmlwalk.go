package gate

import (
	"strings"
	"unicode"
)

type htmlNode struct {
	tag      string
	class    []string
	parent   *htmlNode
	children []*htmlNode
}

var voidHTML = map[string]bool{
	"area": true, "base": true, "br": true, "col": true, "embed": true,
	"hr": true, "img": true, "input": true, "link": true, "meta": true,
	"param": true, "source": true, "track": true, "wbr": true,
}

func parseHTML(fragment string) *htmlNode {
	root := &htmlNode{tag: "#root"}
	stack := []*htmlNode{root}
	s := fragment
	for len(s) > 0 {
		if strings.HasPrefix(s, "<!--") {
			if i := strings.Index(s[4:], "-->"); i >= 0 {
				s = s[4+i+3:]
				continue
			}
			break
		}
		if s[0] != '<' {
			if i := strings.IndexByte(s, '<'); i >= 0 {
				s = s[i:]
				continue
			}
			break
		}
		if len(s) > 1 && s[1] == '/' {
			end := strings.IndexByte(s, '>')
			if end < 0 {
				break
			}
			name := strings.ToLower(strings.TrimSpace(s[2:end]))
			if i := strings.IndexFunc(name, func(r rune) bool { return unicode.IsSpace(r) }); i >= 0 {
				name = name[:i]
			}
			for len(stack) > 1 && stack[len(stack)-1].tag != name {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 1 && stack[len(stack)-1].tag == name {
				stack = stack[:len(stack)-1]
			}
			s = s[end+1:]
			continue
		}
		if len(s) > 1 && (s[1] == '!' || s[1] == '?') {
			end := strings.IndexByte(s, '>')
			if end < 0 {
				break
			}
			s = s[end+1:]
			continue
		}
		end := strings.IndexByte(s, '>')
		if end < 0 {
			break
		}
		inner := s[1:end]
		selfClose := strings.HasSuffix(inner, "/")
		if selfClose {
			inner = strings.TrimSpace(inner[:len(inner)-1])
		}
		name, attrs := splitTag(inner)
		name = strings.ToLower(name)
		n := &htmlNode{tag: name, class: classTokens(attrs["class"]), parent: stack[len(stack)-1]}
		stack[len(stack)-1].children = append(stack[len(stack)-1].children, n)
		s = s[end+1:]
		if selfClose || voidHTML[name] {
			continue
		}
		stack = append(stack, n)
	}
	return root
}

func splitTag(inner string) (name string, attrs map[string]string) {
	attrs = map[string]string{}
	inner = strings.TrimSpace(inner)
	i := 0
	for i < len(inner) && !unicode.IsSpace(rune(inner[i])) {
		i++
	}
	name = inner[:i]
	rest := strings.TrimSpace(inner[i:])
	for rest != "" {
		eq := strings.IndexByte(rest, '=')
		if eq < 0 {
			break
		}
		key := strings.ToLower(strings.TrimSpace(rest[:eq]))
		rest = strings.TrimSpace(rest[eq+1:])
		if rest == "" {
			break
		}
		var val string
		if rest[0] == '"' {
			if j := strings.IndexByte(rest[1:], '"'); j >= 0 {
				val = rest[1 : 1+j]
				rest = strings.TrimSpace(rest[1+j+1:])
			} else {
				break
			}
		} else if rest[0] == '\'' {
			if j := strings.IndexByte(rest[1:], '\''); j >= 0 {
				val = rest[1 : 1+j]
				rest = strings.TrimSpace(rest[1+j+1:])
			} else {
				break
			}
		} else {
			j := 0
			for j < len(rest) && !unicode.IsSpace(rune(rest[j])) {
				j++
			}
			val = rest[:j]
			rest = strings.TrimSpace(rest[j:])
		}
		attrs[key] = val
	}
	return name, attrs
}

func classTokens(v string) []string {
	return strings.Fields(v)
}

func nodeHasClass(n *htmlNode, cls string) bool {
	for _, c := range n.class {
		if c == cls {
			return true
		}
	}
	return false
}

func walkHTML(n *htmlNode, fn func(*htmlNode)) {
	fn(n)
	for _, c := range n.children {
		walkHTML(c, fn)
	}
}

func className(sel string) string {
	return strings.TrimPrefix(sel, ".")
}
