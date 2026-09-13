package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

type ctx struct {
	inFrame bool
	inLabel bool
	inDemo  bool
}

func (c ctx) product() bool {
	if c.inLabel {
		return false
	}
	return c.inFrame || c.inDemo
}

func apply(html string) (string, error) {
	out := injectCSS(html)
	root := parseHTML(out)
	got := rewriteNode(out, root, ctx{})
	if got == "" {
		return "", fmt.Errorf("elements: empty rewrite")
	}
	return got, nil
}

func rewriteNode(html string, n *node, c ctx) string {
	if n.tag == "#root" {
		return rewriteInner(html, n, c)
	}
	if n.tag == "style" || n.tag == "script" {
		return html[n.openStart:n.outerEnd()]
	}
	next := c
	if hasClass(n, "frame") {
		next.inFrame = true
	}
	if hasClass(n, "frame-label") {
		next.inLabel = true
	}
	if hasClass(n, "demo") {
		next.inDemo = true
	}
	inner := rewriteInner(html, n, next)
	if n.tag == "kbd" || hasClass(n, "tc-kbd") {
		return rewriteKbd(html, n, inner)
	}
	if !next.product() {
		return openTag(html, n) + inner + closeTag(html, n)
	}
	return rewriteProduct(html, n, inner)
}

func rewriteInner(html string, n *node, c ctx) string {
	if n.void {
		return ""
	}
	if len(n.children) == 0 {
		if n.innerStart < n.innerEnd {
			return html[n.innerStart:n.innerEnd]
		}
		return ""
	}
	var b strings.Builder
	pos := n.innerStart
	product := c.product()
	for i := 0; i < len(n.children); i++ {
		ch := n.children[i]
		if ch.openStart > pos {
			b.WriteString(html[pos:ch.openStart])
		}
		if product && !hasClass(n, "tc-drop") && i+1 < len(n.children) && wrapAsDrop(ch, n.children[i+1]) {
			second := n.children[i+1]
			firstHTML := rewriteNode(html, ch, c)
			secondHTML := rewriteNode(html, second, c)
			b.WriteString(`<span class="tc-drop">` + firstHTML + secondHTML + `</span>`)
			pos = second.outerEnd()
			i++
			continue
		}
		b.WriteString(rewriteNode(html, ch, c))
		pos = ch.outerEnd()
	}
	if pos < n.innerEnd {
		b.WriteString(html[pos:n.innerEnd])
	}
	return b.String()
}

func wrapAsDrop(a, b *node) bool {
	if hasClass(b, "tc-menu") && hasClass(b, "is-pop") {
		if hasClass(a, "tc-in") || hasClass(a, "tc-btn") {
			return true
		}
	}
	return false
}

func rewriteProduct(html string, n *node, inner string) string {
	switch {
	case hasClass(n, "tc-drop"):
		return openTag(html, n) + inner + closeTag(html, n)
	case hasClass(n, "tc-search"):
		return rewriteSearch(html, n, inner)
	case hasClass(n, "tc-btn"):
		return rewriteBtn(html, n, inner)
	case hasClass(n, "tc-in"):
		return rewriteIn(html, n, inner)
	case hasClass(n, "tc-seg__opt"):
		return retag(html, n, inner, "button", map[string]string{"type": "button"})
	case hasClass(n, "tc-menu__item"):
		return retag(html, n, inner, "button", map[string]string{"type": "button", "role": "menuitem"})
	case hasClass(n, "tc-check"):
		return retag(html, n, inner, "button", map[string]string{"type": "button"})
	case hasClass(n, "tc-tick") || hasClass(n, "tc-tickrow__tick"):
		if insideButtonHost(n) {
			return asSpan(html, n, inner)
		}
		return retag(html, n, inner, "button", map[string]string{"type": "button"})
	case hasClass(n, "tc-usermenu"):
		return retag(html, n, inner, "button", map[string]string{"type": "button"})
	case hasClass(n, "tc-close"):
		return retag(html, n, inner, "button", map[string]string{"type": "button"})
	case n.tag == "a" && (hasClass(n, "tc-nav__item") || hasClass(n, "tc-crumbs__link")):
		return ensureHref(html, n, inner)
	}
	return openTag(html, n) + inner + closeTag(html, n)
}

func insideButtonHost(n *node) bool {
	for p := n.parent; p != nil; p = p.parent {
		if hasAnyClass(p, "tc-btn", "tc-check", "tc-menu__item", "tc-close", "tc-usermenu", "tc-in") {
			return true
		}
	}
	return false
}

func asSpan(html string, n *node, inner string) string {
	if n.tag == "span" {
		return openTag(html, n) + inner + closeTag(html, n)
	}
	open := `<span class="` + strings.Join(n.class, " ") + `"`
	for _, key := range []string{"style", "title"} {
		if v, ok := attrValue(html, n, key); ok {
			open += ` ` + key + `="` + v + `"`
		}
	}
	return open + `>` + inner + `</span>`
}

func rewriteKbd(html string, n *node, inner string) string {
	if n.tag == "kbd" {
		return openTag(html, n) + inner + closeTag(html, n)
	}
	if hasClass(n, "tc-kbd") {
		return retag(html, n, inner, "kbd", nil)
	}
	return openTag(html, n) + inner + closeTag(html, n)
}

func rewriteBtn(html string, n *node, inner string) string {
	menus, rest := splitMenus(html, n)
	if menus == "" {
		return retag(html, n, inner, "button", map[string]string{"type": "button"})
	}
	var restInner strings.Builder
	pos := n.innerStart
	for _, ch := range n.children {
		if hasClass(ch, "tc-menu") {
			if ch.openStart > pos {
				restInner.WriteString(html[pos:ch.openStart])
			}
			pos = ch.outerEnd()
			continue
		}
		if ch.openStart > pos {
			restInner.WriteString(html[pos:ch.openStart])
		}
		restInner.WriteString(rewriteNode(html, ch, ctx{inFrame: true}))
		pos = ch.outerEnd()
	}
	if pos < n.innerEnd {
		restInner.WriteString(html[pos:n.innerEnd])
	}
	_ = rest
	btn := retag(html, n, restInner.String(), "button", map[string]string{"type": "button"})
	return `<span class="tc-drop">` + btn + menus + `</span>`
}

func splitMenus(html string, n *node) (menus, rest string) {
	var mb strings.Builder
	found := false
	for _, ch := range n.children {
		if hasClass(ch, "tc-menu") {
			found = true
			mb.WriteString(rewriteNode(html, ch, ctx{inFrame: true}))
		}
	}
	if !found {
		return "", ""
	}
	return mb.String(), ""
}

func rewriteSearch(html string, n *node, inner string) string {
	inner = normalizeQ(inner)
	if strings.Contains(inner, `class="tc-q"`) || strings.Contains(inner, "class='tc-q'") {
		return ensureRoleSearch(html, n, inner)
	}
	text := strings.TrimSpace(textOutsideElements(inner))
	svgs := svgChunks(inner)
	if text == "" && len(svgs) == 0 {
		return openTag(html, n) + inner + closeTag(html, n)
	}
	var b strings.Builder
	b.WriteString(withAttr(openTag(html, n), "role", "search"))
	for _, s := range svgs {
		b.WriteString(s)
	}
	ph := attrQuote(text)
	b.WriteString(`<input class="tc-q" type="search" value="` + ph + `" aria-label="` + ph + `">`)
	b.WriteString(closeTag(html, n))
	return b.String()
}

func ensureRoleSearch(html string, n *node, inner string) string {
	open := openTag(html, n)
	if _, ok := attrValue(html, n, "role"); !ok {
		open = withAttr(open, "role", "search")
	}
	return open + inner + closeTag(html, n)
}

var qInputRe = regexp.MustCompile(`<input class="tc-q" type="(?:search|text)" (?:placeholder|value)="([^"]*)" aria-label="([^"]*)">`)

func normalizeQ(s string) string {
	return qInputRe.ReplaceAllString(s, `<input class="tc-q" type="search" value="$1" aria-label="$2">`)
}

func rewriteExistingInput(html string, n *node) string {
	return openTag(html, n)
}

func rewriteIn(html string, n *node, inner string) string {
	if n.tag == "input" {
		return rewriteExistingInput(html, n)
	}
	if n.tag == "textarea" || n.tag == "button" {
		return openTag(html, n) + inner + closeTag(html, n)
	}
	if n.tag == "label" {
		return openTag(html, n) + normalizeQ(inner) + closeTag(html, n)
	}
	if hasClass(n, "is-pick") {
		return retag(html, n, inner, "button", map[string]string{"type": "button"})
	}
	if hasComplexInner(n) {
		return openTag(html, n) + inner + closeTag(html, n)
	}
	svgs := svgChildrenHTML(html, n)
	text := strings.TrimSpace(directText(html, n))
	if len(svgs) > 0 {
		return wrapIconField(html, n, svgs, text)
	}
	if hasClass(n, "is-area") || looksMultiline(html, n, inner) {
		return toTextarea(html, n, inner)
	}
	return toInput(html, n, text)
}

func hasComplexInner(n *node) bool {
	for _, ch := range n.children {
		if ch.tag == "svg" {
			continue
		}
		return true
	}
	return false
}

func looksMultiline(html string, n *node, inner string) bool {
	if strings.Contains(inner, "\n") {
		return true
	}
	style, ok := attrValue(html, n, "style")
	if ok && strings.Contains(style, "height:auto") {
		return true
	}
	return false
}

func wrapIconField(html string, n *node, svgs []string, text string) string {
	open := retagOpen(html, n, "label", nil)
	ph := attrQuote(text)
	var b strings.Builder
	b.WriteString(open)
	for _, s := range svgs {
		b.WriteString(s)
	}
	b.WriteString(`<input class="tc-q" type="search" value="` + ph + `" aria-label="` + ph + `">`)
	b.WriteString("</label>")
	return b.String()
}

func toTextarea(html string, n *node, inner string) string {
	open := retagOpen(html, n, "textarea", nil)
	return open + inner + "</textarea>"
}

var (
	dateRe = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	timeRe = regexp.MustCompile(`^\d{1,2}:\d{2}$`)
)

func toInput(html string, n *node, text string) string {
	kind := inputKind(html, n, text)
	attrs := map[string]string{"type": kind}
	if (kind == "date" || kind == "time") && hasClass(n, "is-ph") && !dateRe.MatchString(text) && !timeRe.MatchString(text) {
		attrs["placeholder"] = text
	} else {
		attrs["value"] = text
	}
	if lab := siblingLabel(html, n); lab != "" {
		if _, ok := attrValue(html, n, "aria-label"); !ok {
			attrs["aria-label"] = lab
		}
	} else if hasClass(n, "is-ph") && text != "" {
		attrs["aria-label"] = text
	}
	open := retagOpen(html, n, "input", attrs)
	if !strings.HasSuffix(open, ">") {
		open += ">"
	}
	return open
}

func inputKind(html string, n *node, text string) string {
	if text == "YYYY-MM-DD" || dateRe.MatchString(text) {
		return "date"
	}
	if timeRe.MatchString(text) {
		return "time"
	}
	lab := strings.ToLower(siblingLabel(html, n))
	switch {
	case strings.Contains(lab, "birthday"), strings.Contains(lab, "date"),
		lab == "from", lab == "to", lab == "until":
		if text == "" || text == "—" || text == "YYYY-MM-DD" {
			return "date"
		}
	case strings.Contains(lab, "time"), lab == "start", lab == "end":
		if timeRe.MatchString(text) {
			return "time"
		}
	}
	if n.parent != nil && hasClass(n.parent, "tc-range") {
		if dateRe.MatchString(text) {
			return "date"
		}
	}
	if n.parent != nil && hasClass(n.parent, "tc-bar") {
		return "search"
	}
	return "text"
}

func siblingLabel(html string, n *node) string {
	if n.parent == nil {
		return ""
	}
	for _, ch := range n.parent.children {
		if ch.tag == "label" {
			return strings.TrimSpace(directText(html, ch))
		}
	}
	return ""
}

func retag(html string, n *node, inner, newTag string, extra map[string]string) string {
	if n.tag == newTag {
		open := openTag(html, n)
		for k, v := range extra {
			if _, ok := attrValue(html, n, k); !ok {
				open = withAttr(open, k, v)
			}
		}
		return open + inner + closeTag(html, n)
	}
	open := retagOpen(html, n, newTag, extra)
	return open + inner + "</" + newTag + ">"
}

func retagOpen(html string, n *node, newTag string, extra map[string]string) string {
	raw := html[n.openStart : n.openEnd+1]
	body := raw[1 : len(raw)-1]
	body = strings.TrimSpace(body)
	if strings.HasSuffix(body, "/") {
		body = strings.TrimSpace(body[:len(body)-1])
	}
	i := 0
	for i < len(body) && !unicode.IsSpace(rune(body[i])) {
		i++
	}
	rest := ""
	if i < len(body) {
		rest = body[i:]
	}
	open := "<" + newTag + rest
	open = addAttrs(open, extra)
	if !strings.HasSuffix(open, ">") {
		open += ">"
	}
	return open
}

var attrOrder = []string{"type", "role", "placeholder", "value", "aria-label", "href"}

func addAttrs(open string, extra map[string]string) string {
	if len(extra) == 0 {
		return open
	}
	seen := map[string]bool{}
	add := func(k, v string) {
		if v == "" && k != "value" {
			return
		}
		low := strings.ToLower(open)
		if strings.Contains(low, k+`="`) || strings.Contains(low, k+`='`) {
			return
		}
		open = withAttr(strings.TrimSuffix(open, ">")+">", k, v)
		open = strings.TrimSuffix(open, ">")
	}
	for _, k := range attrOrder {
		if v, ok := extra[k]; ok {
			add(k, v)
			seen[k] = true
		}
	}
	for k, v := range extra {
		if !seen[k] {
			add(k, v)
		}
	}
	return open
}

func withAttr(open, key, val string) string {
	if strings.HasSuffix(open, ">") {
		open = open[:len(open)-1]
	}
	open += ` ` + key + `="` + attrQuote(val) + `">`
	return open
}

func ensureHref(html string, n *node, inner string) string {
	open := openTag(html, n)
	if _, ok := attrValue(html, n, "href"); !ok {
		open = withAttr(open, "href", "#")
	}
	return open + inner + closeTag(html, n)
}

func attrQuote(s string) string {
	s = strings.ReplaceAll(s, `"`, "&quot;")
	return s
}

func directText(html string, n *node) string {
	if n.void {
		return ""
	}
	var b strings.Builder
	pos := n.innerStart
	for _, ch := range n.children {
		if ch.openStart > pos {
			b.WriteString(html[pos:ch.openStart])
		}
		pos = ch.outerEnd()
	}
	if pos < n.innerEnd {
		b.WriteString(html[pos:n.innerEnd])
	}
	return b.String()
}

func svgChildrenHTML(html string, n *node) []string {
	var out []string
	for _, ch := range n.children {
		if ch.tag == "svg" {
			out = append(out, html[ch.openStart:ch.outerEnd()])
		}
	}
	return out
}

func svgChunks(inner string) []string {
	var out []string
	low := strings.ToLower(inner)
	i := 0
	for {
		start := strings.Index(low[i:], "<svg")
		if start < 0 {
			break
		}
		start += i
		endRel := strings.Index(low[start:], "</svg>")
		if endRel < 0 {
			break
		}
		end := start + endRel + len("</svg>")
		out = append(out, inner[start:end])
		i = end
	}
	return out
}

func textOutsideElements(inner string) string {
	var b strings.Builder
	i := 0
	for i < len(inner) {
		if inner[i] == '<' {
			j := strings.IndexByte(inner[i:], '>')
			if j < 0 {
				break
			}
			i += j + 1
			continue
		}
		j := strings.IndexByte(inner[i:], '<')
		if j < 0 {
			b.WriteString(inner[i:])
			break
		}
		b.WriteString(inner[i : i+j])
		i += j
	}
	return b.String()
}
