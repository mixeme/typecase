package main

import (
	"fmt"
	"strings"
)

func apply(html string, rules []Rule) (string, error) {
	out, err := rewriteStyleBlocks(html, rules)
	if err != nil {
		return "", err
	}
	out = rewriteMarkup(out, rules)
	out, err = moveMkRules(out)
	if err != nil {
		return "", err
	}
	out = rewriteScripts(out, rules)
	return out, nil
}

func rewriteStyleBlocks(html string, rules []Rule) (string, error) {
	var b strings.Builder
	i := 0
	low := strings.ToLower(html)
	for i < len(html) {
		start := strings.Index(low[i:], "<style")
		if start < 0 {
			b.WriteString(html[i:])
			return b.String(), nil
		}
		start += i
		gt := strings.Index(html[start:], ">")
		if gt < 0 {
			return "", fmt.Errorf("unclosed <style>")
		}
		openEnd := start + gt + 1
		endRel := strings.Index(low[openEnd:], "</style>")
		if endRel < 0 {
			return "", fmt.Errorf("unclosed </style>")
		}
		close := openEnd + endRel
		b.WriteString(html[i:openEnd])
		b.WriteString(rewriteSelectors(html[openEnd:close], rules))
		b.WriteString(html[close : close+len("</style>")])
		i = close + len("</style>")
	}
	return b.String(), nil
}

func rewriteScripts(html string, rules []Rule) string {
	var b strings.Builder
	i := 0
	low := strings.ToLower(html)
	for i < len(html) {
		start := strings.Index(low[i:], "<script")
		if start < 0 {
			b.WriteString(html[i:])
			return b.String()
		}
		start += i
		gt := strings.Index(html[start:], ">")
		if gt < 0 {
			b.WriteString(html[i:])
			return b.String()
		}
		openEnd := start + gt + 1
		endRel := strings.Index(low[openEnd:], "</script>")
		if endRel < 0 {
			b.WriteString(html[i:])
			return b.String()
		}
		close := openEnd + endRel
		b.WriteString(html[i:openEnd])
		body := rewriteSelectors(html[openEnd:close], rules)
		for _, r := range rules {
			if isSimpleClass(r.From) {
				names := classNames(r.To)
				if len(names) == 1 {
					body = strings.ReplaceAll(body, "'"+r.From[1:]+"'", "'"+names[0]+"'")
					body = strings.ReplaceAll(body, `"`+r.From[1:]+`"`, `"`+names[0]+`"`)
				}
			}
		}
		b.WriteString(body)
		b.WriteString(html[close : close+len("</script>")])
		i = close + len("</script>")
	}
	return b.String()
}

func moveMkRules(html string) (string, error) {
	pageStart, pageOpen, pageClose, ok := findStyle(html, false)
	if !ok {
		return html, nil
	}
	prodStart, prodOpen, prodClose, ok := findStyle(html, true)
	if !ok {
		return html, nil
	}
	pageCSS := html[pageOpen:pageClose]
	prodCSS := html[prodOpen:prodClose]
	keep, move := splitMkCSS(prodCSS)
	if move == "" {
		return html, nil
	}
	pageCSS = strings.TrimRight(pageCSS, "\n") + "\n\n/* native desktop chrome around the product */\n" + move
	if pageStart < prodStart {
		html = html[:pageOpen] + pageCSS + html[pageClose:prodOpen] + keep + html[prodClose:]
	} else {
		html = html[:prodOpen] + keep + html[prodClose:pageOpen] + pageCSS + html[pageClose:]
	}
	return html, nil
}

func findStyle(html string, dataTC bool) (start, open, close int, ok bool) {
	low := strings.ToLower(html)
	i := 0
	for i < len(html) {
		s := strings.Index(low[i:], "<style")
		if s < 0 {
			return 0, 0, 0, false
		}
		s += i
		gt := strings.Index(html[s:], ">")
		if gt < 0 {
			return 0, 0, 0, false
		}
		openTag := strings.ToLower(html[s : s+gt+1])
		body := s + gt + 1
		end := strings.Index(low[body:], "</style>")
		if end < 0 {
			return 0, 0, 0, false
		}
		has := strings.Contains(openTag, "data-tc")
		if has == dataTC {
			return s, body, body + end, true
		}
		i = body + end + len("</style>")
	}
	return 0, 0, 0, false
}

func splitMkCSS(css string) (keep, move string) {
	var k, m strings.Builder
	rest := css
	for {
		rest = strings.TrimLeft(rest, "\n")
		if rest == "" {
			break
		}
		// leading comments stay with the next rule
		lead := ""
		for strings.HasPrefix(strings.TrimSpace(rest), "/*") {
			j := strings.Index(rest, "*/")
			if j < 0 {
				break
			}
			lead += rest[:j+2]
			rest = rest[j+2:]
			if strings.HasPrefix(rest, "\n") {
				lead += "\n"
				rest = rest[1:]
			}
		}
		open := strings.IndexByte(rest, '{')
		if open < 0 {
			k.WriteString(lead)
			k.WriteString(rest)
			break
		}
		end := matchBrace(rest, open)
		if end < 0 {
			k.WriteString(lead)
			k.WriteString(rest)
			break
		}
		rule := rest[:end+1]
		if strings.HasPrefix(rest[end+1:], "\n") {
			rule += "\n"
			rest = rest[end+2:]
		} else {
			rest = rest[end+1:]
		}
		sel := rule[:open]
		if mkOnlySelector(sel) {
			m.WriteString(lead)
			m.WriteString(rule)
			if !strings.HasSuffix(rule, "\n") {
				m.WriteByte('\n')
			}
		} else {
			k.WriteString(lead)
			k.WriteString(rule)
			if !strings.HasSuffix(rule, "\n") {
				k.WriteByte('\n')
			}
		}
	}
	return k.String(), m.String()
}

func mkOnlySelector(sel string) bool {
	hasMK, hasTC := false, false
	for i := 0; i < len(sel); i++ {
		if sel[i] != '.' {
			continue
		}
		j := i + 1
		for j < len(sel) && classContinues(sel, j) {
			j++
		}
		name := sel[i+1 : j]
		if strings.HasPrefix(name, "mk-") {
			hasMK = true
		}
		if strings.HasPrefix(name, "tc-") || name == "tc-root" {
			hasTC = true
		}
	}
	return hasMK && !hasTC
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
