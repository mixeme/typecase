package gate

import "strings"

func stripCSSComments(s string) string {
	var b strings.Builder
	i := 0
	for i < len(s) {
		if i+1 < len(s) && s[i] == '/' && s[i+1] == '*' {
			j := strings.Index(s[i+2:], "*/")
			if j < 0 {
				break
			}
			i = i + 2 + j + 2
			continue
		}
		b.WriteByte(s[i])
		i++
	}
	return b.String()
}

func matchBrace(s string, open int) int {
	depth := 0
	for i := open; i < len(s); i++ {
		switch s[i] {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return i
			}
		}
	}
	return -1
}

func cutAtRules(css string) (body string, blocks []string) {
	var b strings.Builder
	i := 0
	for i < len(css) {
		if atRuleStart(css, i) {
			brace := strings.Index(css[i:], "{")
			if brace < 0 {
				b.WriteString(css[i:])
				break
			}
			end := matchBrace(css, i+brace)
			if end < 0 {
				b.WriteString(css[i:])
				break
			}
			blocks = append(blocks, css[i:end+1])
			i = end + 1
			continue
		}
		b.WriteByte(css[i])
		i++
	}
	return b.String(), blocks
}

func atRuleStart(css string, i int) bool {
	if i > 0 {
		c := css[i-1]
		if c != ' ' && c != '\n' && c != '\t' && c != '}' && c != '{' && c != ';' {
			return false
		}
	}
	rest := css[i:]
	return strings.HasPrefix(rest, "@media") || strings.HasPrefix(rest, "@container")
}

type cssRule struct {
	Selector string
	Body     string
}

func parseCSSRules(css string) []cssRule {
	var rules []cssRule
	i := 0
	for i < len(css) {
		for i < len(css) && (css[i] == ' ' || css[i] == '\n' || css[i] == '\t' || css[i] == '\r') {
			i++
		}
		if i >= len(css) {
			break
		}
		if css[i] == '@' {
			brace := strings.Index(css[i:], "{")
			if brace < 0 {
				break
			}
			end := matchBrace(css, i+brace)
			if end < 0 {
				break
			}
			i = end + 1
			continue
		}
		brace := strings.Index(css[i:], "{")
		if brace < 0 {
			break
		}
		sel := strings.TrimSpace(css[i : i+brace])
		end := matchBrace(css, i+brace)
		if end < 0 {
			break
		}
		body := css[i+brace+1 : end]
		if sel != "" {
			rules = append(rules, cssRule{Selector: sel, Body: body})
		}
		i = end + 1
	}
	return rules
}

func cssProps(body string) map[string]string {
	out := make(map[string]string)
	for _, part := range strings.Split(body, ";") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		k, v, ok := strings.Cut(part, ":")
		if !ok {
			continue
		}
		out[strings.ToLower(strings.TrimSpace(k))] = strings.TrimSpace(v)
	}
	return out
}

func compactSel(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func splitSelectors(sel string) []string {
	var parts []string
	depth := 0
	start := 0
	for i := 0; i < len(sel); i++ {
		switch sel[i] {
		case '(':
			depth++
		case ')':
			depth--
		case ',':
			if depth == 0 {
				parts = append(parts, compactSel(sel[start:i]))
				start = i + 1
			}
		}
	}
	if start < len(sel) {
		parts = append(parts, compactSel(sel[start:]))
	}
	return parts
}

func sameElement(ruleSel, owned string) bool {
	rs := compactSel(ruleSel)
	o := compactSel(owned)
	if rs == o {
		return true
	}
	if !strings.HasPrefix(rs, o) {
		return false
	}
	rest := rs[len(o):]
	if rest == "" {
		return true
	}
	switch rest[0] {
	case ' ', '>', '+', '~':
		return false
	}
	if strings.HasPrefix(rest, "__") {
		return false
	}
	return rest[0] == '.' || rest[0] == ':' || rest[0] == '['
}

func ruleConcerns(ruleSel, owned string) bool {
	for _, rs := range splitSelectors(ruleSel) {
		for _, o := range splitSelectors(owned) {
			if sameElement(rs, o) {
				return true
			}
		}
	}
	return false
}
