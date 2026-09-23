package gate

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// MockupLintRule is one check against the mockup copy. The rules are listed
// in docs/visual/mockup.md.
type MockupLintRule string

const (
	// MockupFrameIDs requires every .frame to have a unique data-frame.
	MockupFrameIDs MockupLintRule = "frame-ids"
	// MockupStyleBlocks requires product CSS in <style data-tc> and page
	// chrome in a second <style>, with no selectors crossing the cut.
	MockupStyleBlocks MockupLintRule = "style-blocks"
	// MockupFrameClasses requires classes inside .frame (except the frame
	// chrome and its label) to be tc-*, is-*, or mk-* (mk-* only in
	// m-desktop frames). .mock is forbidden.
	MockupFrameClasses MockupLintRule = "frame-classes"
	// MockupFrameElements requires real controls in frames: button, kbd.
	MockupFrameElements MockupLintRule = "frame-elements"
	// MockupFrameNoStyle forbids style= inside frames and on .frame itself
	// (frame width is a page class); .frame-label page chrome is exempt.
	MockupFrameNoStyle MockupLintRule = "frame-no-style"
	// MockupFrameNamedClass requires a class on every node in a frame except
	// the step 3 list: inline text tags, svg guts, and semantic table/list
	// / field tags.
	MockupFrameNamedClass MockupLintRule = "frame-named-class"
	// MockupDataAttrs requires every data-tc-colour and data-tc-width value
	// in frames to have a matching rule in the page <style>.
	MockupDataAttrs MockupLintRule = "data-attrs"
	// MockupSelfContained forbids <script>, <link>, and url() outside the
	// mockup directory (fonts/ next to the file).
	MockupSelfContained MockupLintRule = "self-contained"
	// MockupClassesDefined requires every class in a frame to be defined in
	// data-tc, or in the page <style> when it is a page class.
	MockupClassesDefined MockupLintRule = "classes-defined"
	// MockupDeadRules requires every data-tc rule to match at least one
	// frame (dead CSS report empty).
	MockupDeadRules MockupLintRule = "dead-rules"
	// MockupNamedContainer requires every @container to be named (tcbox or
	// a style query) and container: to be tcbox / inline-size.
	MockupNamedContainer MockupLintRule = "named-container"
	// MockupTokenVars forbids var(--x) in data-tc outside the token list
	// and the data properties --tc-colour / --tc-width.
	MockupTokenVars MockupLintRule = "token-vars"
	// MockupEmptyRules forbids a rule with no declarations, and requires
	// every name listed in a "names-without-rule" comment to be used.
	MockupEmptyRules MockupLintRule = "empty-rules"
	// MockupNoNumberedClasses forbids a class named by a number (--rN), in
	// the frames and in data-tc alike.
	MockupNoNumberedClasses MockupLintRule = "no-numbered-classes"
	// MockupNesting requires every closing tag to close the element opened
	// last, and every element to be closed, across the whole file.
	MockupNesting MockupLintRule = "nesting"
)

// MockupLintStep0 is the first rule set: frame ids only.
var MockupLintStep0 = []MockupLintRule{MockupFrameIDs}

// MockupLintLive is every mockup-lint rule; the mockup is locked on all of them.
var MockupLintLive = []MockupLintRule{
	MockupFrameIDs, MockupStyleBlocks, MockupFrameClasses, MockupFrameElements, MockupFrameNoStyle,
	MockupFrameNamedClass, MockupDataAttrs, MockupSelfContained, MockupClassesDefined,
	MockupDeadRules, MockupNamedContainer, MockupTokenVars, MockupEmptyRules,
	MockupNoNumberedClasses, MockupNesting,
}

const mockupRel = "docs/visual/carrel-ui-mockups.html"

// MockupFile is the Typecase copy of the Carrel mockup.
func MockupFile(root string) string {
	return filepath.Join(root, filepath.FromSlash(mockupRel))
}

// CheckMockupLintFile reads the mockup copy and runs the live rule set.
func CheckMockupLintFile(root string) error {
	data, err := os.ReadFile(MockupFile(root))
	if err != nil {
		return fmt.Errorf("mockup-lint: %w", err)
	}
	return CheckMockupLint(string(data), MockupLintLive...)
}

// CheckMockupLint runs the given rules against mockup HTML.
func CheckMockupLint(html string, rules ...MockupLintRule) error {
	if len(rules) == 0 {
		rules = MockupLintLive
	}
	var errs []string
	for _, r := range rules {
		switch r {
		case MockupFrameIDs:
			if err := checkMockupFrameIDs(html); err != nil {
				errs = append(errs, err.Error())
			}
		case MockupStyleBlocks:
			if err := checkMockupStyleBlocks(html); err != nil {
				errs = append(errs, err.Error())
			}
		case MockupFrameClasses:
			if err := checkMockupFrameClasses(html); err != nil {
				errs = append(errs, err.Error())
			}
		case MockupFrameElements:
			if err := checkMockupFrameElements(html); err != nil {
				errs = append(errs, err.Error())
			}
		case MockupFrameNoStyle:
			if err := checkMockupFrameNoStyle(html); err != nil {
				errs = append(errs, err.Error())
			}
		case MockupFrameNamedClass:
			if err := checkMockupFrameNamedClass(html); err != nil {
				errs = append(errs, err.Error())
			}
		case MockupDataAttrs:
			if err := checkMockupDataAttrs(html); err != nil {
				errs = append(errs, err.Error())
			}
		case MockupSelfContained:
			if err := checkMockupSelfContained(html); err != nil {
				errs = append(errs, err.Error())
			}
		case MockupEmptyRules:
			if err := checkMockupEmptyRules(html); err != nil {
				errs = append(errs, err.Error())
			}
		case MockupNoNumberedClasses:
			if err := checkMockupNoNumberedClasses(html); err != nil {
				errs = append(errs, err.Error())
			}
		case MockupNesting:
			if err := checkMockupNesting(html); err != nil {
				errs = append(errs, err.Error())
			}
		case MockupClassesDefined:
			if err := checkMockupClassesDefined(html); err != nil {
				errs = append(errs, err.Error())
			}
		case MockupDeadRules:
			if err := checkMockupDeadRules(html); err != nil {
				errs = append(errs, err.Error())
			}
		case MockupNamedContainer:
			if err := checkMockupNamedContainer(html); err != nil {
				errs = append(errs, err.Error())
			}
		case MockupTokenVars:
			if err := checkMockupTokenVars(html); err != nil {
				errs = append(errs, err.Error())
			}
		default:
			return fmt.Errorf("mockup-lint: unknown rule %q", r)
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("mockup-lint: %s", strings.Join(errs, "; "))
	}
	return nil
}

// ListMockupFrames returns data-frame values in document order.
func ListMockupFrames(html string) []string {
	ids, _ := collectMockupFrames(html)
	out := make([]string, 0, len(ids))
	for _, id := range ids {
		if id != "" {
			out = append(out, id)
		}
	}
	return out
}

func checkMockupFrameIDs(html string) error {
	ids, n := collectMockupFrames(html)
	if n == 0 {
		return fmt.Errorf("no .frame in mockup")
	}
	seen := make(map[string]int, n)
	var missing, dups []string
	for i, id := range ids {
		if id == "" {
			missing = append(missing, fmt.Sprintf("#%d", i+1))
			continue
		}
		if seen[id] > 0 {
			dups = append(dups, id)
		}
		seen[id]++
	}
	var parts []string
	if len(missing) > 0 {
		parts = append(parts, fmt.Sprintf("missing data-frame on %d frames (%s)", len(missing), strings.Join(missing, ", ")))
	}
	if len(dups) > 0 {
		parts = append(parts, "duplicate data-frame "+strings.Join(dups, ", "))
	}
	if len(parts) > 0 {
		return fmt.Errorf("%s", strings.Join(parts, "; "))
	}
	return nil
}

func collectMockupFrames(html string) (ids []string, n int) {
	forEachHTMLStartTag(html, func(name string, attrs map[string]string) {
		if !hasClassToken(attrs["class"], "frame") {
			return
		}
		n++
		ids = append(ids, attrs["data-frame"])
	})
	return ids, n
}

var mockupCSSClassRe = regexp.MustCompile(`\.([A-Za-z_][\w-]*)`)

var mockupPageClass = map[string]bool{
	"wrap": true, "doc-head": true, "lamp": true, "sub": true,
	"lede": true, "note": true, "toc": true, "t-h": true,
	"frame": true, "frame-label": true, "frames": true,
	"swatches": true, "sw": true, "demo": true, "cols": true,
	"plan": true, "frag": true, "chip": true, "doc": true,
	"ladder": true, "ladder__label": true,
}

var mockupPageBare = map[string]bool{
	"*": true, "body": true, "html": true,
	"h1": true, "h2": true, "h3": true, "h4": true,
	"p": true, "a": true, "code": true,
	"figure": true, "figcaption": true, "figcaption b": true,
}

func mockupProductClass(c string) bool {
	// mk-* is page chrome around the product (desktop window). Check it
	// before m-*: "mk-win" would otherwise match the m- prefix.
	if strings.HasPrefix(c, "mk-") {
		return false
	}
	if strings.HasPrefix(c, "tc-") {
		return true
	}
	if c == "mock" || strings.HasPrefix(c, "m-") || strings.HasPrefix(c, "auth-") {
		return true
	}
	switch c {
	case "w-xs", "w-sm", "w-md", "w-lg", "w-xl", "w-full", "has-side", "no-rail":
		return true
	}
	return false
}

type mockupSheet struct {
	dataTC bool
	css    string
}

func checkMockupStyleBlocks(html string) error {
	sheets, err := mockupStyleSheets(html)
	if err != nil {
		return err
	}
	var product, page []string
	for _, s := range sheets {
		if s.dataTC {
			product = append(product, s.css)
		} else {
			page = append(page, s.css)
		}
	}
	if len(product) != 1 {
		return fmt.Errorf("style-blocks: want 1 <style data-tc>, got %d", len(product))
	}
	if len(page) != 1 {
		return fmt.Errorf("style-blocks: want 1 page <style>, got %d", len(page))
	}
	var errs []string
	for _, sel := range mockupCSSSelectors(product[0]) {
		if mockupSelectorHasPage(sel) {
			errs = append(errs, "data-tc has page selector "+sel)
		}
	}
	for _, sel := range mockupCSSSelectors(page[0]) {
		if mockupSelectorHasProduct(sel) {
			errs = append(errs, "page style has product class in "+sel)
		}
	}
	if len(errs) > 6 {
		errs = append(errs[:6], fmt.Sprintf("… %d more", len(errs)-6))
	}
	if len(errs) > 0 {
		return fmt.Errorf("style-blocks: %s", strings.Join(errs, "; "))
	}
	return nil
}

func checkMockupFrameClasses(html string) error {
	var errs []string
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
		id := attrs["data-frame"]
		desktop := strings.HasPrefix(id, "m-desktop")
		walkHTML(parseHTML(outer), func(n *htmlNode) {
			if n.tag == "#root" || n.tag == "script" || n.tag == "style" {
				return
			}
			if nodeHasClass(n, "frame") || nodeHasClass(n, "frame-label") || ancestorHasClass(n, "frame-label") {
				return
			}
			for _, c := range n.class {
				if mockupFrameClassOK(c, desktop) {
					continue
				}
				label := id
				if label == "" {
					label = "(no data-frame)"
				}
				if strings.HasPrefix(c, "mk-") {
					errs = append(errs, fmt.Sprintf("%s: mk-* class %s outside m-desktop", label, c))
					continue
				}
				errs = append(errs, fmt.Sprintf("%s: class %s", label, c))
			}
		})
		i = next
	}
	if len(errs) > 8 {
		errs = append(errs[:8], fmt.Sprintf("… %d more", len(errs)-8))
	}
	if len(errs) > 0 {
		return fmt.Errorf("frame-classes: %s", strings.Join(errs, "; "))
	}
	return nil
}

var mockupButtonClass = map[string]bool{
	"tc-btn": true, "tc-close": true, "tc-check": true,
	"tc-seg__opt": true, "tc-menu__item": true, "tc-usermenu": true,
}

func checkMockupFrameElements(html string) error {
	var errs []string
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
		id := attrs["data-frame"]
		walkHTML(parseHTML(outer), func(n *htmlNode) {
			if n.tag == "#root" || nodeHasClass(n, "frame") || nodeHasClass(n, "frame-label") || ancestorHasClass(n, "frame-label") {
				return
			}
			for _, c := range n.class {
				if mockupButtonClass[c] && n.tag != "button" {
					if c == "tc-menu__item" && n.tag == "div" && htmlNodeHasDescendantClass(n, "tc-seg") {
						continue
					}
					errs = append(errs, fmt.Sprintf("%s: %s is <%s>, want <button>", id, c, n.tag))
				}
				if c == "tc-kbd" && n.tag != "kbd" {
					errs = append(errs, fmt.Sprintf("%s: tc-kbd is <%s>, want <kbd>", id, n.tag))
				}
			}
		})
		i = next
	}
	if len(errs) > 8 {
		errs = append(errs[:8], fmt.Sprintf("… %d more", len(errs)-8))
	}
	if len(errs) > 0 {
		return fmt.Errorf("frame-elements: %s", strings.Join(errs, "; "))
	}
	return nil
}

func checkMockupFrameNoStyle(html string) error {
	var errs []string
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
		id := attrs["data-frame"]
		if st := attrs["style"]; st != "" {
			errs = append(errs, fmt.Sprintf("%s: style= on .frame — width is a page class", id))
		}
		walkHTML(parseHTML(outer), func(n *htmlNode) {
			if n.tag == "#root" || nodeHasClass(n, "frame") || nodeHasClass(n, "frame-label") || ancestorHasClass(n, "frame-label") {
				return
			}
			if n.attrs != nil {
				if st := n.attrs["style"]; st != "" {
					errs = append(errs, fmt.Sprintf("%s: <%s> style=", id, n.tag))
				}
			}
		})
		i = next
	}
	if len(errs) > 8 {
		errs = append(errs[:8], fmt.Sprintf("… %d more", len(errs)-8))
	}
	if len(errs) > 0 {
		return fmt.Errorf("frame-no-style: %s", strings.Join(errs, "; "))
	}
	return nil
}

func mockupFrameClassOK(c string, desktop bool) bool {
	if strings.HasPrefix(c, "tc-") || strings.HasPrefix(c, "is-") {
		return true
	}
	return desktop && strings.HasPrefix(c, "mk-")
}

func mockupStyleSheets(html string) ([]mockupSheet, error) {
	var out []mockupSheet
	i := 0
	for i < len(html) {
		start, end, name, attrs, ok := nextStartTag(html, i)
		if !ok {
			break
		}
		if name != "style" {
			i = end + 1
			continue
		}
		outer, next, err := elementOuter(html, start, name)
		if err != nil {
			return nil, err
		}
		gt := strings.Index(outer, ">")
		closeAt := strings.LastIndex(strings.ToLower(outer), "</style>")
		css := ""
		if gt >= 0 && closeAt > gt {
			css = outer[gt+1 : closeAt]
		}
		_, dataTC := attrs["data-tc"]
		out = append(out, mockupSheet{dataTC: dataTC, css: css})
		i = next
	}
	return out, nil
}

func mockupCSSSelectors(css string) []string {
	css = stripCSSComments(css)
	var out []string
	var walk func(string)
	walk = func(s string) {
		for {
			s = strings.TrimSpace(s)
			if s == "" {
				return
			}
			open := strings.IndexByte(s, '{')
			if open < 0 {
				return
			}
			prelude := compactSel(s[:open])
			end := matchBrace(s, open)
			if end < 0 {
				return
			}
			body := s[open+1 : end]
			if strings.HasPrefix(prelude, "@font-face") {
				out = append(out, "@font-face")
			} else if strings.HasPrefix(prelude, "@") {
				walk(body)
			} else {
				out = append(out, splitSelectors(prelude)...)
			}
			s = s[end+1:]
		}
	}
	walk(css)
	return out
}

func mockupSelectorHasPage(sel string) bool {
	sel = compactSel(sel)
	if sel == "@font-face" {
		return true
	}
	for _, m := range mockupCSSClassRe.FindAllStringSubmatch(sel, -1) {
		if mockupPageClass[m[1]] || strings.HasPrefix(m[1], "mk-") {
			return true
		}
	}
	return mockupPageBare[sel]
}

func mockupSelectorHasProduct(sel string) bool {
	sel = compactSel(sel)
	if strings.HasPrefix(sel, "body:has(") {
		// Page ladder: radios set .tc-root width. Not product CSS.
		return false
	}
	if strings.Contains(sel, "data-tc-colour") || strings.Contains(sel, "data-tc-width") {
		return false
	}
	if strings.Contains(sel, "data-tc") {
		return true
	}
	if strings.HasPrefix(sel, "@container") || strings.HasPrefix(sel, "@media") {
		return true
	}
	for _, m := range mockupCSSClassRe.FindAllStringSubmatch(sel, -1) {
		if mockupProductClass(m[1]) {
			return true
		}
	}
	return false
}

// MockupFrame is one .frame element extracted from the mockup.
type MockupFrame struct {
	ID   string
	HTML string
}

// MockupParts is the mockup split into styles, icon sprite and frames.
type MockupParts struct {
	Styles string
	Sprite string
	Frames []MockupFrame
}

// ParseMockupParts extracts <style> blocks, the zero-size sprite SVG, and each .frame.
func ParseMockupParts(html string) (MockupParts, error) {
	var p MockupParts
	i := 0
	for i < len(html) {
		start, end, name, attrs, ok := nextStartTag(html, i)
		if !ok {
			break
		}
		switch name {
		case "style":
			outer, next, err := elementOuter(html, start, name)
			if err != nil {
				return MockupParts{}, err
			}
			p.Styles += outer
			i = next
			continue
		case "svg":
			if attrs["width"] == "0" && attrs["height"] == "0" {
				outer, next, err := elementOuter(html, start, name)
				if err != nil {
					return MockupParts{}, err
				}
				p.Sprite = outer
				i = next
				continue
			}
		}
		if hasClassToken(attrs["class"], "frame") {
			outer, next, err := elementOuter(html, start, name)
			if err != nil {
				return MockupParts{}, err
			}
			p.Frames = append(p.Frames, MockupFrame{ID: attrs["data-frame"], HTML: outer})
			i = next
			continue
		}
		i = end + 1
	}
	if p.Styles == "" {
		return MockupParts{}, fmt.Errorf("mockup: no <style>")
	}
	if p.Sprite == "" {
		return MockupParts{}, fmt.Errorf("mockup: no sprite svg")
	}
	if len(p.Frames) == 0 {
		return MockupParts{}, fmt.Errorf("mockup: no .frame")
	}
	return p, nil
}

func nextStartTag(html string, i int) (start, end int, name string, attrs map[string]string, ok bool) {
	for i < len(html) {
		if strings.HasPrefix(html[i:], "<!--") {
			j := strings.Index(html[i+4:], "-->")
			if j < 0 {
				return 0, 0, "", nil, false
			}
			i += 4 + j + 3
			continue
		}
		if html[i] != '<' {
			i++
			continue
		}
		if i+1 < len(html) && (html[i+1] == '/' || html[i+1] == '!' || html[i+1] == '?') {
			j := strings.IndexByte(html[i:], '>')
			if j < 0 {
				return 0, 0, "", nil, false
			}
			i += j + 1
			continue
		}
		end = scanTagEnd(html, i)
		if end < 0 {
			return 0, 0, "", nil, false
		}
		inner := strings.TrimSpace(html[i+1 : end])
		if strings.HasSuffix(inner, "/") {
			inner = strings.TrimSpace(inner[:len(inner)-1])
		}
		name, attrs = splitTag(inner)
		return i, end, name, attrs, true
	}
	return 0, 0, "", nil, false
}

func elementOuter(html string, start int, name string) (string, int, error) {
	tagEnd := scanTagEnd(html, start)
	if tagEnd < 0 {
		return "", 0, fmt.Errorf("mockup: unclosed <%s>", name)
	}
	inner := strings.TrimSpace(html[start+1 : tagEnd])
	selfClose := strings.HasSuffix(inner, "/")
	if selfClose || voidHTML[name] {
		return html[start : tagEnd+1], tagEnd + 1, nil
	}
	depth := 1
	i := tagEnd + 1
	for i < len(html) {
		if strings.HasPrefix(html[i:], "<!--") {
			j := strings.Index(html[i+4:], "-->")
			if j < 0 {
				break
			}
			i += 4 + j + 3
			continue
		}
		if html[i] != '<' {
			i++
			continue
		}
		if strings.HasPrefix(html[i:], "</") {
			closeEnd := strings.IndexByte(html[i:], '>')
			if closeEnd < 0 {
				break
			}
			closeEnd += i
			cname := strings.ToLower(strings.TrimSpace(html[i+2 : closeEnd]))
			if sp := strings.IndexFunc(cname, func(r rune) bool { return r == ' ' || r == '\t' || r == '\n' }); sp >= 0 {
				cname = cname[:sp]
			}
			if cname == name {
				depth--
				if depth == 0 {
					return html[start : closeEnd+1], closeEnd + 1, nil
				}
			}
			i = closeEnd + 1
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
		openEnd := scanTagEnd(html, i)
		if openEnd < 0 {
			break
		}
		openInner := strings.TrimSpace(html[i+1 : openEnd])
		openSelf := strings.HasSuffix(openInner, "/")
		if openSelf {
			openInner = strings.TrimSpace(openInner[:len(openInner)-1])
		}
		openName, _ := splitTag(openInner)
		if openName == name && !openSelf && !voidHTML[openName] {
			depth++
		}
		i = openEnd + 1
	}
	return "", 0, fmt.Errorf("mockup: unclosed <%s>", name)
}

func hasClassToken(class, want string) bool {
	for _, c := range classTokens(class) {
		if c == want {
			return true
		}
	}
	return false
}

func forEachHTMLStartTag(html string, fn func(name string, attrs map[string]string)) {
	i := 0
	for i < len(html) {
		if strings.HasPrefix(html[i:], "<!--") {
			j := strings.Index(html[i+4:], "-->")
			if j < 0 {
				return
			}
			i += 4 + j + 3
			continue
		}
		if html[i] != '<' {
			i++
			continue
		}
		if i+1 < len(html) && (html[i+1] == '/' || html[i+1] == '!' || html[i+1] == '?') {
			j := strings.IndexByte(html[i:], '>')
			if j < 0 {
				return
			}
			i += j + 1
			continue
		}
		end := scanTagEnd(html, i)
		if end < 0 {
			return
		}
		inner := strings.TrimSpace(html[i+1 : end])
		if strings.HasSuffix(inner, "/") {
			inner = strings.TrimSpace(inner[:len(inner)-1])
		}
		name, attrs := splitTag(inner)
		fn(name, attrs)
		i = end + 1
	}
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
