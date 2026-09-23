package gate

import (
	"fmt"
	"strings"
)

// checkMockupNesting requires every closing tag to close the element opened
// last, and every element to be closed.
//
// The frame parser counts only tags of the frame's own name, so a stray
// </div> inside a frame does not show there: the frame still comes out
// balanced, just one level short, and the browser quietly moves whatever
// followed out of its container. A rail footer rewritten with one </div> too
// many took the main area out of the shell grid on seven frames, and every
// other check passed. Read strictly, the file is well nested everywhere else,
// so this costs nothing to keep.
func checkMockupNesting(html string) error {
	type open struct {
		name string
		line int
	}
	var stack []open
	var errs []string
	line := func(i int) int { return strings.Count(html[:i], "\n") + 1 }

	i := 0
	for i < len(html) {
		if strings.HasPrefix(html[i:], "<!--") {
			j := strings.Index(html[i+4:], "-->")
			if j < 0 {
				return fmt.Errorf("nesting: unclosed comment at line %d", line(i))
			}
			i += 4 + j + 3
			continue
		}
		if html[i] != '<' {
			i++
			continue
		}
		if i+1 < len(html) && (html[i+1] == '!' || html[i+1] == '?') {
			j := strings.IndexByte(html[i:], '>')
			if j < 0 {
				break
			}
			i += j + 1
			continue
		}
		if strings.HasPrefix(html[i:], "</") {
			j := strings.IndexByte(html[i:], '>')
			if j < 0 {
				break
			}
			name := strings.ToLower(strings.TrimSpace(html[i+2 : i+j]))
			if voidHTML[name] {
				i += j + 1
				continue
			}
			switch {
			case len(stack) > 0 && stack[len(stack)-1].name == name:
				stack = stack[:len(stack)-1]
			case len(stack) > 0:
				top := stack[len(stack)-1]
				errs = append(errs, fmt.Sprintf("line %d: </%s> while <%s> from line %d is open",
					line(i), name, top.name, top.line))
				for k := len(stack) - 1; k >= 0; k-- {
					if stack[k].name == name {
						stack = stack[:k]
						break
					}
				}
			default:
				errs = append(errs, fmt.Sprintf("line %d: </%s> with nothing open", line(i), name))
			}
			i += j + 1
			continue
		}
		end := scanTagEnd(html, i)
		if end < 0 {
			break
		}
		inner := strings.TrimSpace(html[i+1 : end])
		self := strings.HasSuffix(inner, "/")
		if self {
			inner = strings.TrimSpace(inner[:len(inner)-1])
		}
		name, _ := splitTag(inner)
		i = end + 1
		if self || voidHTML[name] {
			continue
		}
		// raw text: the body is not markup
		if name == "style" || name == "script" || name == "textarea" || name == "title" {
			j := strings.Index(strings.ToLower(html[i:]), "</"+name)
			if j < 0 {
				return fmt.Errorf("nesting: unclosed <%s> at line %d", name, line(i))
			}
			i += j
			stack = append(stack, open{name, line(i)})
			continue
		}
		stack = append(stack, open{name, line(i)})
	}
	for _, o := range stack {
		errs = append(errs, fmt.Sprintf("<%s> from line %d is never closed", o.name, o.line))
	}
	if len(errs) == 0 {
		return nil
	}
	if len(errs) > 12 {
		errs = append(errs[:12], fmt.Sprintf("… %d in all", len(errs)))
	}
	return fmt.Errorf("nesting: %s", strings.Join(errs, "; "))
}
