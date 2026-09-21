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

// FormatNumberedAudit writes the report Х-1 asks for: every --rN class with
// its declarations, how often it is used, in which frames, on what element of
// what block.
func FormatNumberedAudit(list []NumberedClass) string {
	var b strings.Builder
	tagBlocks, once, many, uses := 0, 0, 0, 0
	for _, c := range list {
		uses += len(c.Uses)
		if c.TagBlock() {
			tagBlocks++
		}
		if len(c.Uses) == 1 {
			once++
		} else {
			many++
		}
	}
	fmt.Fprintf(&b, "# Классы-номера макета\n\n")
	fmt.Fprintf(&b, "Отчёт для Х-1 плана [1-mockup-tails](../plan/1-mockup-tails.md).\n")
	fmt.Fprintf(&b, "Пересобрать: `go run ./internal/mockcss/numbered`.\n\n")
	fmt.Fprintf(&b, "Классов `--rN`: **%d**, употреблений в кадрах и витринах: **%d**.\n", len(list), uses)
	fmt.Fprintf(&b, "Разовых (одно употребление): **%d**, повторяющихся: **%d**.\n", once, many)
	fmt.Fprintf(&b, "Висят на имени тега, а не на блоке (Р-6 не выполнено): **%d**.\n\n", tagBlocks)
	b.WriteString("Колонки: класс · сколько раз · тег · остальные классы узла · блок, ")
	b.WriteString("в котором узел стоит · кадры · объявления.\n\n")
	for _, c := range list {
		fmt.Fprintf(&b, "## %s — %d\n\n", c.Name, len(c.Uses))
		if c.TagBlock() {
			b.WriteString("*Блок по имени тега.*\n\n")
		}
		for _, r := range c.Rules {
			fmt.Fprintf(&b, "    %s\n", r)
		}
		b.WriteString("\n")
		tags, owns, ins, frames := map[string]int{}, map[string]int{}, map[string]int{}, map[string]int{}
		for _, u := range c.Uses {
			tags[u.Tag]++
			owns[u.Own]++
			ins[u.In]++
			frames[u.Frame]++
		}
		fmt.Fprintf(&b, "- тег: %s\n", countList(tags))
		fmt.Fprintf(&b, "- узел: %s\n", countList(owns))
		fmt.Fprintf(&b, "- в блоке: %s\n", countList(ins))
		fmt.Fprintf(&b, "- кадры: %s\n\n", countList(frames))
	}
	return b.String()
}

func countList(m map[string]int) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if m[keys[i]] != m[keys[j]] {
			return m[keys[i]] > m[keys[j]]
		}
		return keys[i] < keys[j]
	})
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		name := k
		if name == "" {
			name = "(без класса)"
		}
		if m[k] > 1 {
			name += " ×" + strconv.Itoa(m[k])
		}
		parts = append(parts, name)
	}
	return strings.Join(parts, ", ")
}
