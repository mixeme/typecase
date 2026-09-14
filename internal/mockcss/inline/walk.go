package main

import (
	"strings"
	"unicode"
)

type kv struct {
	name, orig, value string
	quote             byte // ", ', or 0 for boolean
}

type tagOpen struct {
	name  string
	attrs []kv
	slash bool
	raw   string
}

func parseOpen(raw string) tagOpen {
	inner := strings.TrimSpace(raw[1 : len(raw)-1])
	slash := strings.HasSuffix(inner, "/")
	if slash {
		inner = strings.TrimSpace(inner[:len(inner)-1])
	}
	name, rest := splitName(inner)
	return tagOpen{name: strings.ToLower(name), attrs: parseAttrs(rest), slash: slash, raw: raw}
}

func splitName(inner string) (name, rest string) {
	i := 0
	for i < len(inner) {
		r := rune(inner[i])
		if unicode.IsSpace(r) {
			break
		}
		i++
	}
	return inner[:i], strings.TrimSpace(inner[i:])
}

func parseAttrs(rest string) []kv {
	var out []kv
	for rest != "" {
		j := 0
		for j < len(rest) {
			r, n := decodeRune(rest[j:])
			if r == '=' || unicode.IsSpace(r) {
				break
			}
			j += n
		}
		if j == 0 {
			rest = strings.TrimSpace(rest[1:])
			continue
		}
		keyOrig := rest[:j]
		key := strings.ToLower(keyOrig)
		rest = strings.TrimSpace(rest[j:])
		if rest == "" || rest[0] != '=' {
			out = append(out, kv{name: key, orig: keyOrig})
			continue
		}
		rest = strings.TrimSpace(rest[1:])
		if rest == "" {
			out = append(out, kv{name: key, orig: keyOrig})
			break
		}
		if rest[0] == '"' || rest[0] == '\'' {
			q := rest[0]
			k := strings.IndexByte(rest[1:], q)
			if k < 0 {
				out = append(out, kv{name: key, orig: keyOrig, value: rest[1:], quote: q})
				break
			}
			out = append(out, kv{name: key, orig: keyOrig, value: rest[1 : 1+k], quote: q})
			rest = strings.TrimSpace(rest[1+k+1:])
			continue
		}
		k := 0
		for k < len(rest) && !unicode.IsSpace(rune(rest[k])) {
			k++
		}
		out = append(out, kv{name: key, orig: keyOrig, value: rest[:k], quote: 0})
		rest = strings.TrimSpace(rest[k:])
	}
	return out
}

func decodeRune(s string) (rune, int) {
	if s == "" {
		return 0, 0
	}
	r := []rune(s)
	return r[0], len(string(r[0]))
}

func (t tagOpen) get(name string) (string, bool) {
	for _, a := range t.attrs {
		if a.name == name {
			return a.value, true
		}
	}
	return "", false
}

func (t *tagOpen) set(name, value string) {
	for i := range t.attrs {
		if t.attrs[i].name == name {
			t.attrs[i].value = value
			if t.attrs[i].quote == 0 {
				t.attrs[i].quote = '"'
			}
			return
		}
	}
	t.attrs = append(t.attrs, kv{name: name, orig: name, value: value, quote: '"'})
}

func (t *tagOpen) del(name string) {
	var out []kv
	for _, a := range t.attrs {
		if a.name != name {
			out = append(out, a)
		}
	}
	t.attrs = out
}

func (t tagOpen) class() []string {
	v, _ := t.get("class")
	return classList(v)
}

func (t *tagOpen) setClass(cs []string) {
	if len(cs) == 0 {
		t.del("class")
		return
	}
	t.set("class", strings.Join(cs, " "))
}

func (t tagOpen) String() string {
	var b strings.Builder
	b.WriteByte('<')
	b.WriteString(t.name)
	for _, a := range t.attrs {
		b.WriteByte(' ')
		if a.orig != "" {
			b.WriteString(a.orig)
		} else {
			b.WriteString(a.name)
		}
		if a.quote == 0 && a.value == "" {
			continue
		}
		q := a.quote
		if q == 0 {
			q = '"'
		}
		b.WriteByte('=')
		b.WriteByte(q)
		b.WriteString(a.value)
		b.WriteByte(q)
	}
	if t.slash {
		b.WriteString(" /")
	}
	b.WriteByte('>')
	return b.String()
}

var voidTags = map[string]bool{
	"area": true, "base": true, "br": true, "col": true, "embed": true,
	"hr": true, "img": true, "input": true, "link": true, "meta": true,
	"param": true, "source": true, "track": true, "wbr": true,
}

func scanTagEnd(s string, start int) int {
	var q byte
	for i := start + 1; i < len(s); i++ {
		c := s[i]
		if q != 0 {
			if c == q {
				q = 0
			}
			continue
		}
		if c == '"' || c == '\'' {
			q = c
			continue
		}
		if c == '>' {
			return i
		}
	}
	return -1
}

type hit struct {
	Frame     string
	Tag       string
	Class     string
	Parent    string
	Style     string
	SVGStroke string
}

type walkCtx struct {
	frame, label, demo bool
	frameID            string
	tag                string
	class              string
}

func product(c walkCtx) bool {
	return (c.frame || c.demo) && !c.label
}

func walkProductTags(html string, fn func(start, end int, t tagOpen, c walkCtx)) {
	var stack []walkCtx
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
		if i+1 < len(html) && (html[i+1] == '!' || html[i+1] == '?') {
			j := strings.IndexByte(html[i:], '>')
			if j < 0 {
				return
			}
			i += j + 1
			continue
		}
		if i+1 < len(html) && html[i+1] == '/' {
			end := strings.IndexByte(html[i:], '>')
			if end < 0 {
				return
			}
			name := strings.ToLower(strings.TrimSpace(html[i+2 : i+end]))
			if sp := strings.IndexFunc(name, unicode.IsSpace); sp >= 0 {
				name = name[:sp]
			}
			for len(stack) > 0 && stack[len(stack)-1].tag != name {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 && stack[len(stack)-1].tag == name {
				stack = stack[:len(stack)-1]
			}
			i += end + 1
			continue
		}
		end := scanTagEnd(html, i)
		if end < 0 {
			return
		}
		raw := html[i : end+1]
		low := strings.ToLower(raw)
		if strings.HasPrefix(low, "<style") || strings.HasPrefix(low, "<script") {
			closeName := "</style>"
			if strings.HasPrefix(low, "<script") {
				closeName = "</script>"
			}
			j := strings.Index(strings.ToLower(html[end+1:]), closeName)
			if j < 0 {
				return
			}
			i = end + 1 + j + len(closeName)
			continue
		}
		t := parseOpen(raw)
		cur := walkCtx{}
		parentClass := ""
		if len(stack) > 0 {
			cur = stack[len(stack)-1]
			parentClass = stack[len(stack)-1].class
		}
		cs := t.class()
		if hasClassName(cs, "frame") {
			cur.frame = true
			cur.frameID, _ = t.get("data-frame")
		}
		if hasClassName(cs, "frame-label") {
			cur.label = true
		}
		if hasClassName(cs, "demo") {
			cur.demo = true
		}
		cur.tag = t.name
		cur.class = strings.Join(cs, " ")
		call := cur
		call.class = parentClass
		fn(i, end, t, call)
		if !t.slash && !voidTags[t.name] {
			stack = append(stack, cur)
		}
		i = end + 1
	}
}

func collectHits(html string) (hits []hit, classN map[string]int, gapFrames map[int]map[string]bool) {
	classN = map[string]int{}
	gapFrames = map[int]map[string]bool{}
	walkProductTags(html, func(_, _ int, t tagOpen, c walkCtx) {
		if !product(c) {
			return
		}
		for _, cl := range t.class() {
			if strings.HasPrefix(cl, "tc-") {
				classN[cl]++
			}
		}
		style, _ := t.get("style")
		stroke, _ := t.get("stroke")
		fill, _ := t.get("fill")
		if style == "" && !strings.HasPrefix(stroke, "var(") && !strings.HasPrefix(fill, "var(") {
			return
		}
		h := hit{
			Frame:     c.frameID,
			Tag:       t.name,
			Class:     strings.Join(t.class(), " "),
			Parent:    c.class,
			Style:     style,
			SVGStroke: stroke,
		}
		if c.demo && h.Frame == "" {
			h.Frame = "demo"
		}
		hits = append(hits, h)
		for _, d := range parseDecls(style) {
			if d.prop == "gap" {
				if n, ok := gapPx(d.val); ok {
					if gapFrames[n] == nil {
						gapFrames[n] = map[string]bool{}
					}
					gapFrames[n][h.Frame] = true
				}
			}
		}
	})
	return hits, classN, gapFrames
}
