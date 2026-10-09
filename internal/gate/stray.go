package gate

import (
	"regexp"
	"sort"
	"strings"
)

// StrayPlan sorts the frames of a changed mockup (docs/development.md, Р-52).
// A frame whose own HTML changed was meant to change: a person looks at its
// picture. A frame whose HTML is the same can only have moved through the
// CSS or the sprite — it strays, and nobody is looking at it; those in reach
// of what changed are the ones to measure.
type StrayPlan struct {
	Meant []string // the frame's own HTML changed
	New   []string // not in the base
	Gone  []string // only in the base
	Check []string // the same HTML on both sides, in reach of what changed
	Same  int      // the same HTML on both sides
	Why   string   // what puts the frames of Check in reach; empty when nothing does
}

// PlanStrays compares the mockup before and after a change.
func PlanStrays(before, after MockupParts) StrayPlan {
	var p StrayPlan
	was := map[string]string{}
	for _, f := range before.Frames {
		was[f.ID] = f.HTML
	}
	var same []MockupFrame
	now := map[string]bool{}
	for _, f := range after.Frames {
		now[f.ID] = true
		html, ok := was[f.ID]
		switch {
		case !ok:
			p.New = append(p.New, f.ID)
		case html != f.HTML:
			p.Meant = append(p.Meant, f.ID)
		default:
			same = append(same, f)
		}
	}
	for _, f := range before.Frames {
		if !now[f.ID] {
			p.Gone = append(p.Gone, f.ID)
		}
	}
	p.Same = len(same)

	all := func(why string) StrayPlan {
		p.Why = why
		for _, f := range same {
			p.Check = append(p.Check, f.ID)
		}
		return p
	}
	if before.Sprite != after.Sprite {
		return all("the sprite changed")
	}
	reach, wide := changedCSS(styleCSS(before.Styles), styleCSS(after.Styles))
	if wide != "" {
		return all(wide)
	}
	if len(reach) == 0 {
		return p
	}
	for _, f := range same {
		for _, need := range reach {
			if hasAll(f.HTML, need) {
				p.Check = append(p.Check, f.ID)
				break
			}
		}
	}
	var names []string
	for _, need := range reach {
		names = append(names, strings.Join(need, " + "))
	}
	p.Why = "rules changed for " + strings.Join(names, ", ")
	return p
}

func hasAll(html string, need []string) bool {
	for _, tok := range need {
		if !strings.Contains(html, tok) {
			return false
		}
	}
	return true
}

// styleCSS is the CSS of the mockup's <style> elements, one after another.
func styleCSS(styles string) string {
	var b strings.Builder
	low := strings.ToLower(styles)
	i := 0
	for {
		s := strings.Index(low[i:], "<style")
		if s < 0 {
			break
		}
		s += i
		open := strings.IndexByte(low[s:], '>')
		if open < 0 {
			break
		}
		end := strings.Index(low[s+open:], "</style")
		if end < 0 {
			break
		}
		b.WriteString(styles[s+open+1 : s+open+end])
		b.WriteByte('\n')
		i = s + open + end
	}
	return b.String()
}

var (
	strayClass = regexp.MustCompile(`\.(-?[A-Za-z_][\w-]*)`)
	strayAttr  = regexp.MustCompile(`\[\s*([A-Za-z_][\w-]*)`)
	// an attribute's value may hold a dot or a bracket of its own
	strayAttrValue = regexp.MustCompile(`\[\s*([A-Za-z_][\w-]*)[^\]]*\]`)
	straySpace     = regexp.MustCompile(`\s+`)
)

// changedCSS says what the rules that differ between two stylesheets can
// reach: for each of their selectors, the classes and attributes it names as
// they are written in a frame — a frame without one of them the selector
// cannot match. What stands in brackets (:not, :is, :has) is not required.
// wide is set when a frame cannot be ruled out by its HTML — a changed rule
// names no class and no attribute (:root, a tag, @font-face), or the rules
// both sides share stand in another order.
func changedCSS(before, after string) (reach [][]string, wide string) {
	b, a := flatRules(before), flatRules(after)
	count := map[string]int{}
	for _, r := range b {
		count[r.key()]++
	}
	var changed []flatRule
	var keptA []string
	for _, r := range a {
		if count[r.key()] > 0 {
			count[r.key()]--
			keptA = append(keptA, r.key())
			continue
		}
		changed = append(changed, r)
	}
	left := map[string]int{}
	for _, r := range a {
		left[r.key()]++
	}
	var keptB []string
	for _, r := range b {
		if left[r.key()] > 0 {
			left[r.key()]--
			keptB = append(keptB, r.key())
			continue
		}
		changed = append(changed, r)
	}
	if len(changed) == 0 && strings.Join(keptA, "\n") != strings.Join(keptB, "\n") {
		return nil, "the rules stand in another order"
	}
	if len(changed) > 0 && strings.Join(keptA, "\n") != strings.Join(keptB, "\n") {
		return nil, "rules changed and the rest stand in another order"
	}
	seen := map[string]bool{}
	for _, r := range changed {
		if r.at {
			return nil, "changed " + r.sel
		}
		for _, sel := range splitSelectors(r.sel) {
			need := map[string]bool{}
			bare := strayAttrValue.ReplaceAllString(dropBrackets(sel), "[$1]")
			for _, m := range strayClass.FindAllStringSubmatch(bare, -1) {
				need[m[1]] = true
			}
			for _, m := range strayAttr.FindAllStringSubmatch(bare, -1) {
				need[m[1]] = true
			}
			if len(need) == 0 {
				return nil, "changed a rule for " + strings.TrimSpace(sel)
			}
			var toks []string
			for t := range need {
				toks = append(toks, t)
			}
			sort.Strings(toks)
			if k := strings.Join(toks, " "); !seen[k] {
				seen[k] = true
				reach = append(reach, toks)
			}
		}
	}
	sort.Slice(reach, func(i, j int) bool { return strings.Join(reach[i], " ") < strings.Join(reach[j], " ") })
	return reach, ""
}

// dropBrackets takes out of a selector what stands in round brackets.
func dropBrackets(sel string) string {
	var b strings.Builder
	depth := 0
	for _, r := range sel {
		switch {
		case r == '(':
			depth++
		case r == ')':
			if depth > 0 {
				depth--
			}
		case depth == 0:
			b.WriteRune(r)
		}
	}
	return b.String()
}

// flatRule is a rule with the conditions it stands under.
type flatRule struct {
	ctx  string // the @media and @container around it
	sel  string
	body string
	at   bool // an at-rule kept whole (@font-face, @keyframes)
}

func (r flatRule) key() string { return r.ctx + "\x00" + r.sel + "\x00" + r.body }

func flatRules(css string) []flatRule {
	css = stripCSSComments(css)
	norm := func(s string) string { return strings.TrimSpace(straySpace.ReplaceAllString(s, " ")) }
	var rules []flatRule
	var walk func(s, ctx string)
	walk = func(s, ctx string) {
		i := 0
		for i < len(s) {
			for i < len(s) && (s[i] == ' ' || s[i] == '\n' || s[i] == '\t' || s[i] == '\r' || s[i] == ';') {
				i++
			}
			if i >= len(s) {
				return
			}
			brace := strings.Index(s[i:], "{")
			if brace < 0 {
				return
			}
			end := matchBrace(s, i+brace)
			if end < 0 {
				return
			}
			head := norm(s[i : i+brace])
			body := s[i+brace+1 : end]
			switch {
			case strings.HasPrefix(head, "@media"), strings.HasPrefix(head, "@container"), strings.HasPrefix(head, "@supports"), strings.HasPrefix(head, "@layer"):
				walk(body, ctx+head+" ")
			case strings.HasPrefix(head, "@"):
				rules = append(rules, flatRule{ctx: ctx, sel: head, body: norm(body), at: true})
			case head != "":
				rules = append(rules, flatRule{ctx: ctx, sel: head, body: norm(body)})
			}
			i = end + 1
		}
	}
	walk(css, "")
	return rules
}
