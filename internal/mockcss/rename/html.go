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
	tag           string
	class         []string
	classValStart int // inside quotes; -1 if no class attr
	classValEnd   int
	tagStart      int // '<' of start tag
	tagEnd        int // '>' of start tag
	parent        *node
	children      []*node
	nth           int // 1-based among element siblings
}

type patch struct {
	start, end int
	text       string
}

func rewriteMarkup(html string, rules []Rule) string {
	root := parseMarkup(html)
	simple := map[string][]string{}
	desc := []descRule{}
	child := []childRule{}
	nth := []nthRule{}
	tableNth := []tableNthRule{}
	tagAdd := []tagRule{}
	for _, r := range rules {
		if isSimpleClass(r.From) {
			if names := classNames(r.To); len(names) > 0 {
				simple[r.From[1:]] = names
				continue
			}
		}
		if nr, ok := parseNth(r); ok {
			nth = append(nth, nr)
			continue
		}
		if tr, ok := parseTableNth(r); ok {
			tableNth = append(tableNth, tr)
			continue
		}
		if tr, ok := parseTagRule(r); ok {
			tagAdd = append(tagAdd, tr)
			continue
		}
		if cr, ok := parseChildClass(r); ok {
			child = append(child, cr)
			continue
		}
		if dr, ok := parseDescClass(r); ok {
			desc = append(desc, dr)
		}
	}
	var patches []patch
	walkNode(root, func(n *node) {
		if n.tag == "#root" || n.tag == "style" || n.tag == "script" {
			return
		}
		old := append([]string(nil), n.class...)
		next := rewriteNodeClasses(n, simple, desc, child, nth, tableNth, tagAdd)
		if sameTokens(old, next) {
			return
		}
		val := strings.Join(next, " ")
		if n.classValStart >= 0 {
			patches = append(patches, patch{n.classValStart, n.classValEnd, val})
			return
		}
		if val == "" {
			return
		}
		insert := ` class="` + val + `"`
		patches = append(patches, patch{n.tagEnd, n.tagEnd, insert})
	})
	return applyPatches(html, patches)
}

type descRule struct {
	anc, from string
	to        []string
}
type childRule struct {
	parent, from string
	to           []string
}
type nthRule struct {
	parent string
	n      int
	add    []string
}
type tableNthRule struct {
	parent, tag string
	n           int
	add         []string
}

func parseTableNth(r Rule) (tableNthRule, bool) {
	from := strings.Join(strings.Fields(r.From), " ")
	i := strings.Index(from, ":nth-child(")
	if i < 0 {
		return tableNthRule{}, false
	}
	pre := strings.TrimSpace(from[:i])
	parts := strings.Fields(pre)
	if len(parts) != 2 || !strings.HasPrefix(parts[0], ".") {
		return tableNthRule{}, false
	}
	tag := parts[1]
	if strings.HasPrefix(tag, ".") {
		return tableNthRule{}, false
	}
	num, ok := cutInt(from[i+len(":nth-child("):])
	if !ok {
		return tableNthRule{}, false
	}
	add := lastClass(r.To)
	if len(add) == 0 {
		return tableNthRule{}, false
	}
	return tableNthRule{parent: parts[0][1:], tag: tag, n: num, add: add}, true
}

type tagRule struct {
	parent, tag string
	direct      bool
	to          []string
}

func parseNth(r Rule) (nthRule, bool) {
	from := strings.Join(strings.Fields(r.From), " ")
	i := strings.Index(from, ":nth-child(")
	if i < 0 {
		return nthRule{}, false
	}
	pre := strings.TrimSpace(from[:i])
	pre = strings.TrimSuffix(pre, ">")
	pre = strings.TrimSpace(pre)
	if !strings.HasPrefix(pre, ".") || strings.ContainsAny(pre[1:], ". >+~:[") {
		return nthRule{}, false
	}
	rest := from[i+len(":nth-child("):]
	num, ok := cutInt(rest)
	if !ok {
		return nthRule{}, false
	}
	add := lastClass(r.To)
	if len(add) == 0 {
		return nthRule{}, false
	}
	return nthRule{parent: pre[1:], n: num, add: add}, true
}

func parseTagRule(r Rule) (tagRule, bool) {
	from := strings.Join(strings.Fields(r.From), " ")
	if strings.Contains(from, ":") || strings.Contains(from, ",") {
		return tagRule{}, false
	}
	direct := strings.Contains(from, ">")
	parts := strings.Fields(strings.ReplaceAll(from, ">", " "))
	if len(parts) != 2 || !strings.HasPrefix(parts[0], ".") {
		return tagRule{}, false
	}
	tag := parts[1]
	if strings.HasPrefix(tag, ".") || strings.ContainsAny(tag, ".#:[") {
		return tagRule{}, false
	}
	to := htmlTargetClasses(r.To)
	if len(to) == 0 {
		return tagRule{}, false
	}
	return tagRule{parent: parts[0][1:], tag: tag, direct: direct, to: to}, true
}

func parseChildClass(r Rule) (childRule, bool) {
	from := strings.Join(strings.Fields(r.From), " ")
	if !strings.Contains(from, ">") {
		return childRule{}, false
	}
	parts := strings.SplitN(from, ">", 2)
	if len(parts) != 2 {
		return childRule{}, false
	}
	p, c := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
	if !isSimpleClass(p) || !isSimpleClass(c) {
		return childRule{}, false
	}
	to := classNames(strings.ReplaceAll(r.To, " ", ""))
	if len(to) == 0 {
		to = lastClass(r.To)
	}
	if len(to) == 0 {
		return childRule{}, false
	}
	return childRule{parent: p[1:], from: c[1:], to: to}, true
}

func parseDescClass(r Rule) (descRule, bool) {
	from := strings.Join(strings.Fields(r.From), " ")
	if strings.ContainsAny(from, ">+~:[,") {
		return descRule{}, false
	}
	parts := strings.Fields(from)
	if len(parts) != 2 || !isSimpleClass(parts[0]) || !isSimpleClass(parts[1]) {
		return descRule{}, false
	}
	to := htmlTargetClasses(r.To)
	if len(to) == 0 {
		return descRule{}, false
	}
	return descRule{anc: parts[0][1:], from: parts[1][1:], to: to}, true
}

func htmlTargetClasses(to string) []string {
	to = strings.TrimSpace(to)
	if i := strings.Index(to, ","); i >= 0 {
		to = strings.TrimSpace(to[:i])
	}
	if strings.ContainsAny(to, " >+~:[") {
		return lastClass(to)
	}
	if names := classNames(to); len(names) > 0 {
		return names
	}
	return lastClass(to)
}

func lastClass(sel string) []string {
	sel = strings.TrimSpace(sel)
	if i := strings.LastIndex(sel, "."); i >= 0 {
		name := sel[i+1:]
		if j := strings.IndexAny(name, " >+~:[,"); j >= 0 {
			name = name[:j]
		}
		if name != "" {
			return []string{name}
		}
	}
	return nil
}

func cutInt(s string) (int, bool) {
	n := 0
	i := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		n = n*10 + int(s[i]-'0')
		i++
	}
	if i == 0 || i >= len(s) || s[i] != ')' {
		return 0, false
	}
	return n, true
}

func rewriteNodeClasses(n *node, simple map[string][]string, desc []descRule, child []childRule, nth []nthRule, tableNth []tableNthRule, tagAdd []tagRule) []string {
	tokens := append([]string(nil), n.class...)

	for _, rule := range nth {
		if n.nth == rule.n && n.parent != nil && hasClass(n.parent, rule.parent) {
			tokens = addClasses(tokens, rule.add)
		}
	}
	for _, rule := range tableNth {
		if n.tag == rule.tag && n.nth == rule.n && hasAncestor(n, rule.parent) {
			tokens = addClasses(tokens, rule.add)
		}
	}
	for _, rule := range child {
		if n.parent != nil && hasClass(n.parent, rule.parent) && hasToken(tokens, rule.from) {
			tokens = replaceToken(tokens, rule.from, rule.to)
		}
	}
	for _, rule := range desc {
		if hasToken(tokens, rule.from) && hasAncestor(n, rule.anc) {
			tokens = replaceToken(tokens, rule.from, rule.to)
		}
	}
	for _, rule := range tagAdd {
		if n.tag != rule.tag {
			continue
		}
		ok := false
		if rule.direct {
			ok = n.parent != nil && hasClass(n.parent, rule.parent)
		} else {
			ok = hasAncestor(n, rule.parent)
		}
		if ok {
			tokens = addClasses(tokens, rule.to)
		}
	}
	var out []string
	for _, t := range tokens {
		if next, ok := simple[t]; ok {
			out = addClasses(out, next)
			continue
		}
		out = addClasses(out, []string{t})
	}
	if hasToken(out, "tc-root") && n != nil {
		if descendantHas(n, "tc-shell__body") || descendantHas(n, "m-shell") {
			out = addClasses(out, []string{"tc-shell"})
		}
	}
	if hasToken(out, "is-print") && hasToken(out, "tc-root") {
		out = addClasses(out, []string{"tc-root--print"})
		out = removeToken(out, "is-print")
	}
	if n.tag == "div" && len(n.class) == 0 && n.parent != nil && hasClass(n.parent, "m-head") {
		out = addClasses(out, []string{"tc-head__titles"})
	}
	if n.tag == "div" && len(n.class) == 0 && n.parent != nil && hasClass(n.parent, "tc-head") {
		out = addClasses(out, []string{"tc-head__titles"})
	}
	return out
}

func replaceToken(tokens []string, from string, to []string) []string {
	var out []string
	for _, t := range tokens {
		if t == from {
			out = addClasses(out, to)
			continue
		}
		out = append(out, t)
	}
	return out
}

func addClasses(tokens, add []string) []string {
	for _, a := range add {
		if a == "" || hasToken(tokens, a) {
			continue
		}
		tokens = append(tokens, a)
	}
	return tokens
}

func removeToken(tokens []string, drop string) []string {
	var out []string
	for _, t := range tokens {
		if t != drop {
			out = append(out, t)
		}
	}
	return out
}

func hasToken(tokens []string, c string) bool {
	for _, t := range tokens {
		if t == c {
			return true
		}
	}
	return false
}

func hasClass(n *node, c string) bool {
	if n == nil {
		return false
	}
	return hasToken(n.class, c)
}

func hasAncestor(n *node, c string) bool {
	for p := n.parent; p != nil; p = p.parent {
		if hasClass(p, c) {
			return true
		}
	}
	return false
}

func descendantHas(n *node, c string) bool {
	for _, ch := range n.children {
		if hasClass(ch, c) || descendantHas(ch, c) {
			return true
		}
	}
	return false
}

func sameTokens(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func walkNode(n *node, fn func(*node)) {
	fn(n)
	for _, c := range n.children {
		walkNode(c, fn)
	}
}

func applyPatches(s string, patches []patch) string {
	if len(patches) == 0 {
		return s
	}
	for i := 0; i < len(patches); i++ {
		for j := i + 1; j < len(patches); j++ {
			if patches[j].start > patches[i].start {
				patches[i], patches[j] = patches[j], patches[i]
			}
		}
	}
	for _, p := range patches {
		s = s[:p.start] + p.text + s[p.end:]
	}
	return s
}

func parseMarkup(s string) *node {
	root := &node{tag: "#root"}
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
			end := strings.IndexByte(s[i:], '>')
			if end < 0 {
				break
			}
			name := strings.ToLower(strings.TrimSpace(s[i+2 : i+end]))
			if sp := strings.IndexFunc(name, unicode.IsSpace); sp >= 0 {
				name = name[:sp]
			}
			for len(stack) > 1 && stack[len(stack)-1].tag != name {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 1 && stack[len(stack)-1].tag == name {
				stack = stack[:len(stack)-1]
			}
			i += end + 1
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
		name, class, cs, ce := parseStartTag(s, i, end)
		name = strings.ToLower(name)
		if name == "style" || name == "script" {
			close := strings.Index(strings.ToLower(s[end+1:]), "</"+name+">")
			if close < 0 {
				break
			}
			i = end + 1 + close + len(name) + 3
			continue
		}
		parent := stack[len(stack)-1]
		n := &node{
			tag: name, class: class, classValStart: cs, classValEnd: ce,
			tagStart: i, tagEnd: end, parent: parent, nth: len(parent.children) + 1,
		}
		parent.children = append(parent.children, n)
		i = end + 1
		if selfClose || voidHTML[name] {
			continue
		}
		stack = append(stack, n)
	}
	return root
}

func parseStartTag(html string, start, end int) (name string, class []string, cs, ce int) {
	cs, ce = -1, -1
	inner := html[start+1 : end]
	if strings.HasSuffix(strings.TrimSpace(inner), "/") {
		inner = strings.TrimSpace(strings.TrimSpace(inner)[:len(strings.TrimSpace(inner))-1])
	}
	i := 0
	for i < len(inner) {
		r, n := utf8.DecodeRuneInString(inner[i:])
		if unicode.IsSpace(r) {
			break
		}
		i += n
	}
	name = inner[:i]
	rest := inner
	pos := start + 1
	// Walk attributes in the original tag slice html[start+1:end].
	raw := html[start+1 : end]
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
			vs := start + 1 + k
			p := strings.IndexByte(raw[k:], q)
			if p < 0 {
				break
			}
			val = raw[k : k+p]
			if key == "class" {
				cs, ce = vs, vs+p
				class = strings.Fields(val)
			}
			k += p + 1
			continue
		}
		vs := start + 1 + k
		p := k
		for p < len(raw) && !unicode.IsSpace(rune(raw[p])) && raw[p] != '/' {
			p++
		}
		val = raw[k:p]
		if key == "class" {
			cs, ce = vs, start+1+p
			class = strings.Fields(val)
		}
		k = p
	}
	_ = rest
	_ = pos
	return name, class, cs, ce
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
