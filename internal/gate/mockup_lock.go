package gate

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

func checkMockupFrameNamedClass(html string) error {
	var errs []string
	_ = forEachMockupFrame(html, func(id, outer string) error {
		walkMockupFrameNodes(outer, func(n *htmlNode) {
			if len(n.class) > 0 || mockupClasslessOK(n) {
				return
			}
			errs = append(errs, fmt.Sprintf("%s: <%s> has no class", id, n.tag))
		})
		return nil
	})
	return mockupLintErrs("frame-named-class", errs, 8)
}

func mockupClasslessOK(n *htmlNode) bool {
	if n.tag == "svg" || ancestorHasTag(n, "svg") {
		return true
	}
	switch n.tag {
	case "b", "strong", "em", "i", "br",
		"thead", "tbody", "tfoot", "tr", "th", "td",
		"dt", "dd", "li", "option", "label":
		return true
	case "template":
		// The row a button «Add …» adds: not drawn.
		return true
	case "input":
		// A hidden field of a form is data, not a drawn element: CSRF, a
		// version; a file field is the form's that a dropped file goes with.
		return n.attrs["type"] == "hidden" || n.attrs["type"] == "file"
	}
	return false
}

func checkMockupDataAttrs(html string) error {
	product, page, err := mockupSplitCSS(html)
	if err != nil {
		return err
	}
	_ = product
	haveColour := map[string]bool{}
	haveWidth := map[string]bool{}
	for _, m := range mockupDataAttrRuleRe.FindAllStringSubmatch(page, -1) {
		switch m[1] {
		case "colour":
			haveColour[m[2]] = true
		case "width":
			haveWidth[m[2]] = true
		}
	}
	var errs []string
	seenC, seenW := map[string]bool{}, map[string]bool{}
	_ = forEachMockupFrame(html, func(id, outer string) error {
		walkMockupFrameNodes(outer, func(n *htmlNode) {
			if n.attrs == nil {
				return
			}
			if v := n.attrs["data-tc-colour"]; v != "" && !seenC[v] {
				seenC[v] = true
				if !haveColour[v] {
					errs = append(errs, "data-tc-colour="+v+" has no page rule")
				}
			}
			if v := n.attrs["data-tc-width"]; v != "" && !seenW[v] {
				seenW[v] = true
				if !haveWidth[v] {
					errs = append(errs, "data-tc-width="+v+" has no page rule")
				}
			}
		})
		return nil
	})
	return mockupLintErrs("data-attrs", errs, 8)
}

func checkMockupSelfContained(html string) error {
	var errs []string
	forEachHTMLStartTag(html, func(name string, attrs map[string]string) {
		switch name {
		case "script":
			errs = append(errs, "<script>")
		case "link":
			errs = append(errs, "<link>")
		}
	})
	css := ""
	sheets, err := mockupStyleSheets(html)
	if err != nil {
		return err
	}
	for _, s := range sheets {
		css += s.css
	}
	for _, m := range mockupCSSURLRe.FindAllStringSubmatch(css, -1) {
		u := m[1]
		if u == "" {
			u = m[2]
		}
		if u == "" {
			u = m[3]
		}
		u = strings.TrimSpace(u)
		if !mockupLocalURL(u) {
			errs = append(errs, "url("+u+") outside mockup dir")
		}
	}
	return mockupLintErrs("self-contained", errs, 8)
}

func mockupLocalURL(u string) bool {
	if u == "" || strings.Contains(u, "://") || strings.HasPrefix(u, "/") || strings.HasPrefix(u, "data:") {
		return false
	}
	if strings.Contains(u, "..") {
		return false
	}
	u = strings.TrimPrefix(u, "./")
	return strings.HasPrefix(u, "fonts/")
}

// mockupNamesWithoutRule reads the names a sheet declares as named but not
// styled, from "/* names-without-rule: a b c */" comments.
//
// A widget prints a class for the structure it builds whether or not the
// mockup gives that class a look. Writing ".tc-line__item{}" to satisfy
// classes-defined puts a rule with no declarations into the product CSS, and
// the library carries it on; saying the name out loud instead keeps the
// stylesheet free of rules that do nothing.
func mockupNamesWithoutRule(css string) map[string]bool {
	out := map[string]bool{}
	for _, m := range mockupNamesRe.FindAllStringSubmatch(css, -1) {
		for _, f := range strings.Fields(strings.NewReplacer(",", " ", ".", " ").Replace(m[1])) {
			if f == "" || strings.HasPrefix(f, "—") {
				break
			}
			out[f] = true
		}
	}
	return out
}

var mockupNamesRe = regexp.MustCompile(`(?s)/\*\s*names-without-rule:(.*?)\*/`)

func checkMockupClassesDefined(html string) error {
	product, page, err := mockupSplitCSS(html)
	if err != nil {
		return err
	}
	inProduct := cssClassNames(product)
	for n := range mockupNamesWithoutRule(product) {
		inProduct[n] = true
	}
	inPage := cssClassNames(page)
	for n := range mockupNamesWithoutRule(page) {
		inPage[n] = true
	}
	var errs []string
	seen := map[string]bool{}
	_ = forEachMockupFrame(html, func(id, outer string) error {
		walkMockupFrameNodes(outer, func(n *htmlNode) {
			for _, c := range n.class {
				if seen[c] {
					continue
				}
				seen[c] = true
				switch {
				case strings.HasPrefix(c, "tc-"):
					if !inProduct[c] {
						errs = append(errs, c+" used in frames, not in data-tc")
					}
				case strings.HasPrefix(c, "is-"):
					if !inProduct[c] && !inPage[c] {
						errs = append(errs, c+" used in frames, not defined")
					}
				case strings.HasPrefix(c, "mk-") || mockupPageClass[c]:
					if !inPage[c] {
						errs = append(errs, c+" used in frames, not in page style")
					}
				default:
					errs = append(errs, c+" used in frames, not defined")
				}
			}
		})
		return nil
	})
	return mockupLintErrs("classes-defined", errs, 8)
}

// checkMockupEmptyRules keeps rules with no declarations out of both sheets and
// keeps the names-without-rule list honest.
//
// Х-4: rules such as ".tc-line__item{}" were written to satisfy
// classes-defined. They say nothing, they pass dead-rules because the selector
// matches, and the library would carry them into its own CSS. A name that has
// no look is declared in a comment instead, and every declared name has to be
// on a node — otherwise the list outlives the markup.
func checkMockupEmptyRules(html string) error {
	product, page, err := mockupSplitCSS(html)
	if err != nil {
		return err
	}
	var errs []string
	for _, sheet := range []struct {
		name string
		css  string
	}{{"data-tc", product}, {"page", page}} {
		for _, r := range parseCSSRulesDeep(sheet.css) {
			if strings.TrimSpace(r.Body) == "" {
				errs = append(errs, fmt.Sprintf("%s: %s has no declarations", sheet.name, compactSel(r.Selector)))
			}
		}
	}
	used := map[string]bool{}
	_ = forEachMockupProduct(html, func(id, outer string) error {
		walkMockupFrameNodes(outer, func(n *htmlNode) {
			for _, c := range n.class {
				used[c] = true
			}
		})
		return nil
	})
	for _, css := range []string{product, page} {
		for n := range mockupNamesWithoutRule(css) {
			if !used[n] {
				errs = append(errs, n+" is listed as a name without a rule but stands on no node")
			}
		}
	}
	sort.Strings(errs)
	return mockupLintErrs("empty-rules", errs, 8)
}

// scriptFolded says every selector of the list names a state the script's
// fit sets: a bar's fold step or measure, a table's cards.
func scriptFolded(list string) bool {
	parts := splitSelectors(list)
	for _, p := range parts {
		if !strings.Contains(p, "[data-tc-fold") && !strings.Contains(p, "[data-tc-measure]") && !strings.Contains(p, "[data-tc-cards]") && !strings.Contains(p, "[data-tc-side=") {
			return false
		}
	}
	return len(parts) > 0
}

func checkMockupDeadRules(html string) error {
	product, _, err := mockupSplitCSS(html)
	if err != nil {
		return err
	}
	var frames []mockupFrameIndex
	_ = forEachMockupFrame(html, func(id, outer string) error {
		frames = append(frames, indexMockupFrame(id, outer))
		return nil
	})
	var errs []string
	for _, r := range parseCSSRulesDeep(product) {
		if mockupRuleLive(r.Selector, frames) {
			continue
		}
		// The body of a page of the library (package page) and what stands
		// right in it — the shell as high as the window (Carrel's task 272):
		// a frame holds no body, so these rules are seen on no frame by
		// design; TestShellFillsTheWindow sees them at work.
		if sel := strings.TrimSpace(r.Selector); pageRule(sel) || strings.HasPrefix(sel, "[data-tc-noscript] ") {
			continue
		}
		// A bar as tc.js folds it by measure (data-tc-fold, data-tc-measure),
		// a table it turns into cards (data-tc-cards), a menu it turns to
		// the side with room (data-tc-side): the frames are drawn
		// without the script, so no frame holds them; the browser's checks of
		// the script see these rules at work.
		if scriptFolded(r.Selector) {
			continue
		}
		errs = append(errs, compactSel(r.Selector))
	}
	if len(errs) > 8 {
		errs = append(errs[:8], fmt.Sprintf("… %d more", len(errs)-8))
	}
	if len(errs) > 0 {
		return fmt.Errorf("dead-rules: %s", strings.Join(errs, "; "))
	}
	return nil
}

func checkMockupNamedContainer(html string) error {
	product, _, err := mockupSplitCSS(html)
	if err != nil {
		return err
	}
	var errs []string
	walkCSSAtRules(product, func(prelude string) {
		p := compactSel(prelude)
		if !strings.HasPrefix(p, "@container") {
			return
		}
		rest := strings.TrimSpace(strings.TrimPrefix(p, "@container"))
		if rest == "" || strings.HasPrefix(rest, "(") {
			errs = append(errs, "unnamed "+p)
			return
		}
		if strings.HasPrefix(rest, "style(") || strings.HasPrefix(rest, "tcbox ") || rest == "tcbox" || strings.HasPrefix(rest, "tcbox(") {
			return
		}
		errs = append(errs, "unnamed "+p)
	})
	for _, r := range parseCSSRulesDeep(product) {
		for prop, val := range cssProps(r.Body) {
			switch prop {
			case "container-type":
				errs = append(errs, "container-type without name in "+compactSel(r.Selector))
			case "container-name":
				if compactSel(val) != "tcbox" {
					errs = append(errs, "container-name "+val)
				}
			case "container":
				v := compactSel(val)
				if !strings.HasPrefix(v, "tcbox") {
					errs = append(errs, "container:"+val+" in "+compactSel(r.Selector))
				}
			}
		}
	}
	return mockupLintErrs("named-container", errs, 8)
}

func checkMockupTokenVars(html string) error {
	product, _, err := mockupSplitCSS(html)
	if err != nil {
		return err
	}
	var errs []string
	seen := map[string]bool{}
	for _, m := range mockupVarRe.FindAllStringSubmatch(product, -1) {
		name := m[1]
		if mockupAllowedVar[name] || seen[name] {
			continue
		}
		seen[name] = true
		errs = append(errs, name)
	}
	return mockupLintErrs("token-vars", errs, 8)
}

var mockupAllowedVar = map[string]bool{
	"--tc-ink": true, "--tc-surface": true, "--tc-bg": true, "--tc-card": true,
	"--tc-rule": true, "--tc-muted": true, "--tc-accent": true, "--tc-danger": true, "--tc-progress": true,
	"--tc-edge": true, "--tc-hair": true, "--tc-dusty": true,
	"--tc-font-serif": true, "--tc-font-sans": true, "--tc-font-mono": true,
	"--tc-font-size": true, "--tc-row-height": true, "--tc-row-height-touch": true,
	"--tc-page-gutter": true, "--tc-page-gutter-narrow": true,
	"--tc-gap-tight": true, "--tc-gap": true, "--tc-gap-wide": true, "--tc-gap-loose": true, "--tc-fields-gap": true,
	"--tc-radius": true, "--tc-radius-sm": true, "--tc-stroke": true,
	"--tc-rail": true, "--tc-side": true, "--tc-top": true,
	"--tc-colour": true, "--tc-width": true,
}

var (
	mockupDataAttrRuleRe = regexp.MustCompile(`\[data-tc-(colour|width)=["']([^"']+)["']\]`)
	mockupCSSURLRe       = regexp.MustCompile(`url\(\s*(?:'([^']+)'|"([^"]+)"|([^'")\s]+))\s*\)`)
	mockupVarRe          = regexp.MustCompile(`var\(\s*(--[A-Za-z0-9-]+)`)
)

func mockupSplitCSS(html string) (product, page string, err error) {
	sheets, err := mockupStyleSheets(html)
	if err != nil {
		return "", "", err
	}
	var nProduct, nPage int
	for _, s := range sheets {
		if s.dataTC {
			product += s.css
			nProduct++
		} else {
			page += s.css
			nPage++
		}
	}
	if nProduct != 1 || nPage != 1 {
		return "", "", fmt.Errorf("style-blocks: want 1 <style data-tc> and 1 page <style>, got %d/%d", nProduct, nPage)
	}
	return product, page, nil
}

func cssClassNames(css string) map[string]bool {
	out := map[string]bool{}
	for _, sel := range mockupCSSSelectors(css) {
		for _, m := range mockupCSSClassRe.FindAllStringSubmatch(sel, -1) {
			out[m[1]] = true
		}
	}
	return out
}

func forEachMockupProduct(html string, fn func(id, outer string) error) error {
	if err := forEachMockupFrame(html, fn); err != nil {
		return err
	}
	n := 0
	i := 0
	for i < len(html) {
		start, end, name, attrs, ok := nextStartTag(html, i)
		if !ok {
			break
		}
		if !hasClassToken(attrs["class"], "demo") {
			i = end + 1
			continue
		}
		outer, next, err := elementOuter(html, start, name)
		if err != nil {
			return err
		}
		n++
		if err := fn(fmt.Sprintf("demo-%d", n), outer); err != nil {
			return err
		}
		i = next
	}
	return nil
}

func forEachMockupFrame(html string, fn func(id, outer string) error) error {
	i := 0
	for i < len(html) {
		start, end, name, attrs, ok := nextStartTag(html, i)
		if !ok {
			break
		}
		if !hasClassToken(attrs["class"], "frame") {
			i = end + 1
			continue
		}
		outer, next, err := elementOuter(html, start, name)
		if err != nil {
			return err
		}
		if err := fn(attrs["data-frame"], outer); err != nil {
			return err
		}
		i = next
	}
	return nil
}

func walkMockupFrameNodes(outer string, fn func(*htmlNode)) {
	walkHTML(parseHTML(outer), func(n *htmlNode) {
		if n.tag == "#root" || n.tag == "script" || n.tag == "style" {
			return
		}
		if nodeHasClass(n, "frame") || nodeHasClass(n, "frame-label") || ancestorHasClass(n, "frame-label") {
			return
		}
		fn(n)
	})
}

type mockupFrameEl struct {
	tag   string
	class map[string]bool
	attrs map[string]string
}

type mockupFrameIndex struct {
	els []mockupFrameEl
}

func indexMockupFrame(id, outer string) mockupFrameIndex {
	_ = id
	var idx mockupFrameIndex
	walkMockupFrameNodes(outer, func(n *htmlNode) {
		el := mockupFrameEl{tag: n.tag, class: map[string]bool{}, attrs: n.attrs}
		if el.attrs == nil {
			el.attrs = map[string]string{}
		}
		for _, c := range n.class {
			el.class[c] = true
		}
		idx.els = append(idx.els, el)
	})
	return idx
}

func (f mockupFrameIndex) matchCompound(c mockupSelCompound) bool {
	if c.root || c.star && len(c.classes) == 0 && c.tag == "" && len(c.attrs) == 0 {
		return true
	}
	for _, el := range f.els {
		if el.matchCompound(c) {
			return true
		}
	}
	return false
}

func (el mockupFrameEl) matchCompound(c mockupSelCompound) bool {
	if c.root {
		return true
	}
	if c.tag != "" && el.tag != c.tag {
		return false
	}
	for _, cl := range c.classes {
		if !el.class[cl] {
			return false
		}
	}
	for _, a := range c.attrs {
		v, ok := el.attrs[a.name]
		if !ok {
			return false
		}
		if a.hasVal && v != a.val {
			return false
		}
	}
	return true
}

func mockupRuleLive(sel string, frames []mockupFrameIndex) bool {
	for _, alt := range splitSelectors(sel) {
		root, compounds := parseMockupSel(alt)
		if root {
			return true
		}
		if len(compounds) == 0 {
			return true
		}
		for _, f := range frames {
			ok := true
			for _, c := range compounds {
				if !f.matchCompound(c) {
					ok = false
					break
				}
			}
			if ok {
				return true
			}
		}
	}
	return false
}

type mockupSelAttr struct {
	name   string
	hasVal bool
	val    string
}

type mockupSelCompound struct {
	tag     string
	star    bool
	root    bool
	classes []string
	attrs   []mockupSelAttr
}

func parseMockupSel(sel string) (root bool, compounds []mockupSelCompound) {
	sel = compactSel(sel)
	if sel == ":root" || strings.HasPrefix(sel, ":root ") || strings.HasPrefix(sel, ":root:") || strings.HasPrefix(sel, ":root.") {
		return true, nil
	}
	sel = stripSelFunctions(sel)
	sel = stripSelPseudos(sel)
	for _, part := range splitSelCompounds(sel) {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		compounds = append(compounds, parseMockupCompound(part))
	}
	return false, compounds
}

func stripSelFunctions(s string) string {
	for {
		i := strings.Index(s, ":has(")
		if i < 0 {
			i = strings.Index(s, ":not(")
		}
		if i < 0 {
			i = strings.Index(s, ":is(")
		}
		if i < 0 {
			i = strings.Index(s, ":where(")
		}
		if i < 0 {
			i = strings.Index(s, ":nth-child(")
		}
		if i < 0 {
			i = strings.Index(s, ":nth-of-type(")
		}
		if i < 0 {
			i = strings.Index(s, ":nth-last-child(")
		}
		if i < 0 {
			break
		}
		open := strings.IndexByte(s[i:], '(')
		if open < 0 {
			break
		}
		open += i
		depth := 0
		end := -1
		for j := open; j < len(s); j++ {
			switch s[j] {
			case '(':
				depth++
			case ')':
				depth--
				if depth == 0 {
					end = j
				}
			}
			if end >= 0 {
				break
			}
		}
		if end < 0 {
			break
		}
		s = s[:i] + s[end+1:]
	}
	return s
}

func stripSelPseudos(s string) string {
	var b strings.Builder
	i := 0
	for i < len(s) {
		if s[i] == ':' {
			i++
			if i < len(s) && s[i] == ':' {
				i++
			}
			for i < len(s) && (s[i] == '-' || s[i] == '_' || unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i]))) {
				i++
			}
			continue
		}
		b.WriteByte(s[i])
		i++
	}
	return b.String()
}

func splitSelCompounds(sel string) []string {
	var parts []string
	depth := 0
	start := 0
	for i := 0; i < len(sel); i++ {
		switch sel[i] {
		case '[':
			depth++
		case ']':
			if depth > 0 {
				depth--
			}
		case ' ', '>', '+', '~':
			if depth == 0 {
				if i > start {
					parts = append(parts, sel[start:i])
				}
				start = i + 1
			}
		}
	}
	if start < len(sel) {
		parts = append(parts, sel[start:])
	}
	return parts
}

func parseMockupCompound(s string) mockupSelCompound {
	s = strings.TrimSpace(s)
	var c mockupSelCompound
	if s == "*" {
		c.star = true
		return c
	}
	if s == ":root" {
		c.root = true
		return c
	}
	i := 0
	if i < len(s) && (s[i] == '*' || isSelIdentStart(s[i])) {
		if s[i] == '*' {
			c.star = true
			i++
		} else {
			j := i + 1
			for j < len(s) && isSelIdentChar(s[j]) {
				j++
			}
			c.tag = strings.ToLower(s[i:j])
			i = j
		}
	}
	for i < len(s) {
		switch s[i] {
		case '.':
			j := i + 1
			for j < len(s) && isSelIdentChar(s[j]) {
				j++
			}
			if j > i+1 {
				c.classes = append(c.classes, s[i+1:j])
			}
			i = j
		case '[':
			end := strings.IndexByte(s[i:], ']')
			if end < 0 {
				return c
			}
			c.attrs = append(c.attrs, parseMockupSelAttr(s[i+1:i+end]))
			i = i + end + 1
		case '*':
			c.star = true
			i++
		default:
			i++
		}
	}
	return c
}

func parseMockupSelAttr(s string) mockupSelAttr {
	s = strings.TrimSpace(s)
	name, val, ok := strings.Cut(s, "=")
	a := mockupSelAttr{name: strings.ToLower(strings.TrimSpace(name))}
	if !ok {
		return a
	}
	val = strings.TrimSpace(val)
	val = strings.Trim(val, `"'`)
	a.hasVal = true
	a.val = val
	return a
}

func isSelIdentStart(b byte) bool {
	return b == '_' || b == '-' || unicode.IsLetter(rune(b))
}

func isSelIdentChar(b byte) bool {
	return isSelIdentStart(b) || unicode.IsDigit(rune(b))
}

func walkCSSAtRules(css string, fn func(prelude string)) {
	css = stripCSSComments(css)
	var walk func(string)
	walk = func(s string) {
		i := 0
		for i < len(s) {
			at := strings.IndexByte(s[i:], '@')
			if at < 0 {
				return
			}
			i += at
			brace := strings.Index(s[i:], "{")
			if brace < 0 {
				return
			}
			end := matchBrace(s, i+brace)
			if end < 0 {
				return
			}
			fn(strings.TrimSpace(s[i : i+brace]))
			walk(s[i+brace+1 : end])
			i = end + 1
		}
	}
	walk(css)
}

func mockupLintErrs(rule string, errs []string, capN int) error {
	if len(errs) == 0 {
		return nil
	}
	if len(errs) > capN {
		errs = append(errs[:capN], fmt.Sprintf("… %d more", len(errs)-capN))
	}
	return fmt.Errorf("%s: %s", rule, strings.Join(errs, "; "))
}

// pageRule says every selector of a rule starts at the body of a page of
// the library.
func pageRule(sel string) bool {
	for _, one := range strings.Split(sel, ",") {
		if !strings.HasPrefix(strings.TrimSpace(one), "body[data-tc-root]") {
			return false
		}
	}
	return true
}
