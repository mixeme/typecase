package gate

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// numberedRe matches the классы-номера the inline script left behind:
// .tc-div--r37, .tc-span--r15. Only --r<digits> counts; --gap-14 and --w110
// are names by value, not numbers.
var numberedRe = regexp.MustCompile(`--r(\d+)$`)

// IsNumberedClass reports whether c is a numbered class (block--rN).
func IsNumberedClass(c string) bool { return numberedRe.MatchString(c) }

// NumberedUse is one node carrying a numbered class.
type NumberedUse struct {
	Frame string // data-frame, or demo-N
	Tag   string
	Own   string // the node's other tc-* classes
	In    string // nearest block above it
}

// NumberedClass gathers everything a decision about one --rN class needs:
// what it declares, how often it is used, in which frames, on which element of
// which block.
type NumberedClass struct {
	Name  string
	Base  string   // the part before --rN
	Rules []string // rules naming it, in sheet order, as written
	Uses  []NumberedUse
}

// TagBlock reports whether the class is hung on a tag name rather than a block
// (tc-div, tc-span). Those are the ones Р-6 does not allow at all.
func (n NumberedClass) TagBlock() bool { return mockupTagBlock[n.Base] }

var mockupTagBlock = map[string]bool{
	"tc-div": true, "tc-span": true, "tc-ul": true, "tc-b": true,
	"tc-td": true, "tc-tr": true, "tc-p": true, "tc-li": true,
	"tc-strong": true, "tc-dd": true,
}

// MockupNumberedClasses returns every --rN class in the mockup.
func MockupNumberedClasses(html string) ([]NumberedClass, error) {
	product, _, err := mockupSplitCSS(html)
	if err != nil {
		return nil, err
	}
	byName := map[string]*NumberedClass{}
	get := func(name string) *NumberedClass {
		c, ok := byName[name]
		if !ok {
			c = &NumberedClass{Name: name, Base: name[:strings.LastIndex(name, "--r")]}
			byName[name] = c
		}
		return c
	}
	for _, r := range parseCSSRulesDeep(product) {
		text := compactSel(r.Selector) + "{" + compactSel(r.Body) + "}"
		seen := map[string]bool{}
		for _, c := range cssSelectorClasses(r.Selector) {
			if !IsNumberedClass(c) || seen[c] {
				continue
			}
			seen[c] = true
			cl := get(c)
			cl.Rules = append(cl.Rules, text)
		}
	}
	if err := forEachMockupProduct(html, func(id, outer string) error {
		walkMockupFrameNodes(outer, func(n *htmlNode) {
			var own []string
			var numbered []string
			for _, c := range n.class {
				if IsNumberedClass(c) {
					numbered = append(numbered, c)
					continue
				}
				own = append(own, c)
			}
			if len(numbered) == 0 {
				return
			}
			use := NumberedUse{Frame: id, Tag: n.tag, Own: strings.Join(own, " "), In: mockupBlockAbove(n)}
			for _, c := range numbered {
				cl := get(c)
				cl.Uses = append(cl.Uses, use)
			}
		})
		return nil
	}); err != nil {
		return nil, err
	}
	out := make([]NumberedClass, 0, len(byName))
	for _, c := range byName {
		out = append(out, *c)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Base != out[j].Base {
			return out[i].Base < out[j].Base
		}
		return numberOf(out[i].Name) < numberOf(out[j].Name)
	})
	return out, nil
}

func numberOf(name string) int {
	m := numberedRe.FindStringSubmatch(name)
	if m == nil {
		return 0
	}
	n, _ := strconv.Atoi(m[1])
	return n
}

// mockupBlockAbove names the nearest block a node sits in, skipping the glue
// components and the tag-name pseudo-blocks: those say where a node is, not
// what it belongs to.
func mockupBlockAbove(n *htmlNode) string {
	for p := n.parent; p != nil; p = p.parent {
		var blocks []string
		for _, c := range p.class {
			if !strings.HasPrefix(c, "tc-") || IsNumberedClass(c) {
				continue
			}
			base := c
			if i := strings.Index(base, "--"); i >= 0 {
				base = base[:i]
			}
			if mockupTagBlock[base] || mockupGlue[base] {
				continue
			}
			blocks = append(blocks, c)
		}
		if len(blocks) > 0 {
			return strings.Join(blocks, " ")
		}
	}
	return "—"
}

var mockupGlue = map[string]bool{
	"tc-line": true, "tc-stack": true, "tc-push": true,
	"tc-pad": true, "tc-grow": true, "tc-cols": true,
}

// cssSelectorClasses returns the class names a selector names.
func cssSelectorClasses(sel string) []string {
	var out []string
	for _, m := range cssClassRe.FindAllStringSubmatch(sel, -1) {
		out = append(out, m[1])
	}
	return out
}

var cssClassRe = regexp.MustCompile(`\.([A-Za-z_][\w-]*)`)

// checkMockupNoNumberedClasses keeps the numbered classes out for good.
//
// The inline script of the closed plan moved each distinct style= line into a
// class with a number: tc-div--r37 is neither a block, nor an element, nor a
// variant (Р-6), and a widget that prints it would be printing a line number.
// A variant named by its value — tc-table__col--w110, tc-line--gap-14 — is a
// name and passes.
func checkMockupNoNumberedClasses(html string) error {
	list, err := MockupNumberedClasses(html)
	if err != nil {
		return err
	}
	var errs []string
	for _, c := range list {
		where := "data-tc"
		if len(c.Uses) > 0 {
			where = c.Uses[0].Frame
		}
		errs = append(errs, fmt.Sprintf("%s (%s, %d uses)", c.Name, where, len(c.Uses)))
	}
	sort.Strings(errs)
	return mockupLintErrs("no-numbered-classes", errs, 8)
}
