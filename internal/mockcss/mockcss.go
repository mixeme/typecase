// Package mockcss takes stylesheet text from the mockup's own <style> block.
//
// Proof of the "no translation" path: rules are picked by class set and
// renamed; no number is retyped. A selector part is kept only when every class
// it names is in the allowed set, so rules for other contexts (.mock.is-phone,
// .m-bar > .m-btn, :has(> .m-f)) fall away by themselves.
package mockcss

import (
	"regexp"
	"sort"
	"strings"
)

// Options says which rules to take and how to rename them.
type Options struct {
	Keep    map[string]bool   // classes allowed in a selector, without the dot
	Classes map[string]string // mockup class → library class; unlisted keep their name
	Vars    map[string]string // mockup custom property → library property (without "--")
	Prefix  string            // prefix for mockup properties not in Vars ("tcx-")
	// Elements rewrites a selector where the widget prints another element than
	// the frame (".m-seg span" → ".m-seg :is(span,button)"). Applied before Keep.
	Elements map[string]string
	// DropAt lists @media/@container preludes not carried over, with a reason
	// kept by the caller.
	DropAt []string
	// DropRules lists selectors (as written in the mockup) not carried over.
	DropRules []string
	// Selectors rewrites a selector fragment that means something else in the
	// library (".is-dark " → "[data-tc-theme=dark] "). Applied before Keep.
	Selectors map[string]string
}

// Frame returns the markup of the .mock block of the frame labelled exactly
// label, or "".
func Frame(html, label string) string {
	at := strings.Index(html, `<div class="frame-label"><span>`+label+`</span>`)
	if at < 0 {
		return ""
	}
	start := strings.Index(html[at:], `<div class="mock`)
	if start < 0 {
		return ""
	}
	start += at
	depth := 0
	for i := start; i < len(html); {
		switch {
		case strings.HasPrefix(html[i:], "<div"):
			depth++
			i += 4
		case strings.HasPrefix(html[i:], "</div>"):
			depth--
			i += 6
			if depth == 0 {
				return html[start:i]
			}
		default:
			i++
		}
	}
	return ""
}

var classAttrRe = regexp.MustCompile(`class="([^"]*)"`)

// ClassesIn returns every class named in class="" attributes of markup.
func ClassesIn(markup string) map[string]bool {
	out := map[string]bool{}
	for _, m := range classAttrRe.FindAllStringSubmatch(markup, -1) {
		for _, c := range strings.Fields(m[1]) {
			out[c] = true
		}
	}
	return out
}

// Style returns the text of the first <style> element in html.
func Style(html string) string {
	low := strings.ToLower(html)
	start := strings.Index(low, "<style")
	if start < 0 {
		return ""
	}
	gt := strings.Index(html[start:], ">")
	if gt < 0 {
		return ""
	}
	body := start + gt + 1
	end := strings.Index(low[body:], "</style>")
	if end < 0 {
		return ""
	}
	return html[body : body+end]
}

// Generate returns the rules of css that concern only o.Keep, renamed.
func Generate(css string, o Options) string {
	var b strings.Builder
	for _, r := range parse(stripComments(css)) {
		writeRule(&b, r, o)
	}
	return b.String()
}

type rule struct {
	prelude string
	body    string // declarations, or nested rules for @media/@container
	nested  []rule
}

func writeRule(b *strings.Builder, r rule, o Options) {
	if strings.HasPrefix(r.prelude, "@") {
		if !strings.HasPrefix(r.prelude, "@container") && !strings.HasPrefix(r.prelude, "@media") {
			return
		}
		if contains(o.DropAt, r.prelude) {
			return
		}
		var inner strings.Builder
		for _, n := range r.nested {
			writeRule(&inner, n, o)
		}
		if inner.Len() == 0 {
			return
		}
		b.WriteString(r.prelude)
		b.WriteString("{\n")
		b.WriteString(inner.String())
		b.WriteString("}\n")
		return
	}
	var kept []string
	for _, part := range splitTop(r.prelude, ',') {
		part = strings.TrimSpace(part)
		if contains(o.DropRules, part) {
			continue
		}
		for _, from := range sortedKeys(o.Elements) {
			part = strings.ReplaceAll(part, from, o.Elements[from])
		}
		for _, from := range sortedKeys(o.Selectors) {
			part = strings.ReplaceAll(part, from, o.Selectors[from])
		}
		if keepSelector(part, o.Keep) {
			kept = append(kept, renameClasses(part, o.Classes))
		}
	}
	if len(kept) == 0 {
		return
	}
	b.WriteString(strings.Join(kept, ","))
	b.WriteByte('{')
	b.WriteString(renameVars(compact(r.body), o))
	b.WriteString("}\n")
}

var classRe = regexp.MustCompile(`\.([A-Za-z_][\w-]*)`)

func keepSelector(sel string, keep map[string]bool) bool {
	classes := classRe.FindAllStringSubmatch(sel, -1)
	if len(classes) == 0 {
		return false
	}
	for _, m := range classes {
		if !keep[m[1]] {
			return false
		}
	}
	return true
}

func renameClasses(sel string, names map[string]string) string {
	return classRe.ReplaceAllStringFunc(sel, func(m string) string {
		if to, ok := names[m[1:]]; ok {
			return "." + to
		}
		return m
	})
}

var varUseRe = regexp.MustCompile(`var\(--([\w-]+)`)
var varDeclRe = regexp.MustCompile(`(^|;)\s*--([\w-]+)\s*:`)

// RenameVars renames custom properties in declarations as Generate does.
func RenameVars(body string, o Options) string { return renameVars(body, o) }

func renameVars(body string, o Options) string {
	name := func(v string) string {
		if to, ok := o.Vars[v]; ok {
			return to
		}
		return o.Prefix + v
	}
	body = varUseRe.ReplaceAllStringFunc(body, func(m string) string {
		return "var(--" + name(m[len("var(--"):])
	})
	return varDeclRe.ReplaceAllStringFunc(body, func(m string) string {
		sub := varDeclRe.FindStringSubmatch(m)
		return sub[1] + "--" + name(sub[2]) + ":"
	})
}

func contains(list []string, s string) bool {
	for _, x := range list {
		if x == s {
			return true
		}
	}
	return false
}

func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func compact(body string) string {
	lines := strings.Fields(body)
	return strings.Join(lines, " ")
}

func stripComments(css string) string {
	var b strings.Builder
	for {
		i := strings.Index(css, "/*")
		if i < 0 {
			b.WriteString(css)
			return b.String()
		}
		b.WriteString(css[:i])
		j := strings.Index(css[i+2:], "*/")
		if j < 0 {
			return b.String()
		}
		css = css[i+2+j+2:]
	}
}

// parse reads top-level rules; @media/@container bodies are parsed again.
func parse(css string) []rule {
	var out []rule
	for {
		css = strings.TrimSpace(css)
		if css == "" {
			return out
		}
		open := strings.IndexByte(css, '{')
		if open < 0 {
			return out
		}
		prelude := strings.TrimSpace(css[:open])
		end := matchBrace(css, open)
		if end < 0 {
			return out
		}
		body := css[open+1 : end]
		r := rule{prelude: strings.Join(strings.Fields(prelude), " "), body: body}
		if strings.HasPrefix(prelude, "@") {
			r.nested = parse(body)
		}
		out = append(out, r)
		css = css[end+1:]
	}
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

// splitTop splits s on sep outside parentheses.
func splitTop(s string, sep byte) []string {
	var parts []string
	depth, start := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			depth++
		case ')':
			depth--
		case sep:
			if depth == 0 {
				parts = append(parts, s[start:i])
				start = i + 1
			}
		}
	}
	return append(parts, s[start:])
}
