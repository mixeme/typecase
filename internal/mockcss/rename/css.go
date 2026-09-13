package main

import "strings"

func rewriteSelectors(css string, rules []Rule) string {
	for _, r := range sortRules(rules) {
		css = replaceSel(css, r.From, r.To)
	}
	return css
}

func sortRules(rules []Rule) []Rule {
	out := append([]Rule(nil), rules...)
	for i := 1; i < len(out); i++ {
		for j := i; j > 0; j-- {
			a, b := out[j-1], out[j]
			if len(b.From) > len(a.From) || (len(b.From) == len(a.From) && b.From < a.From) {
				out[j-1], out[j] = b, a
				continue
			}
			break
		}
	}
	return out
}

func replaceSel(s, from, to string) string {
	if from == "" || from == to {
		return s
	}
	var b strings.Builder
	i := 0
	for {
		j := strings.Index(s[i:], from)
		if j < 0 {
			b.WriteString(s[i:])
			return b.String()
		}
		j += i
		end := j + len(from)
		if classContinues(s, end) {
			b.WriteString(s[i:end])
			i = end
			continue
		}
		b.WriteString(s[i:j])
		b.WriteString(to)
		i = end
	}
}

func classContinues(s string, end int) bool {
	if end >= len(s) {
		return false
	}
	c := s[end]
	return c == '-' || c == '_' ||
		(c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') ||
		(c >= '0' && c <= '9')
}
