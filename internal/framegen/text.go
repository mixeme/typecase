package framegen

import (
	"regexp"
	"strings"
	"unicode"
)

// gen translates frames: it knows the widget package's enums.
type gen struct {
	en    enums
	faces map[string]string // Face constant by its classes
}

// ev is ", Field: w.Const" for the enum t matching e's classes, or "".
func (g *gen) ev(e *El, t, field string) string {
	if c := g.en.of(e, t); c != "" {
		return ", " + field + ": " + c
	}
	return ""
}

var places = [][2]string{{"tc-push", "Pushed"}, {"tc-grow", "Grown"}}

func placef(e *El) string {
	for _, p := range places {
		if e.Has(p[0]) {
			return ", Place: w." + p[1]
		}
	}
	return ""
}

// kbd is a key: its words, where it stands and whether it is secondary —
// the same in every container (the review of 14.10: six copies kept four
// different parts of it).
func (g *gen) kbd(e *El) string {
	return "w.Kbd{Keys: " + q(e.Text()) + placesec(e) + "}"
}

// span is words in a face: the face its classes name with their size, or
// the colour alone when no face has exactly them (the review of 14.10: the
// copies dropped the size).
func (g *gen) span(e *El, colour string) string {
	face := colour
	if f, ok := g.faces[join(e.Cls(), " ")]; ok {
		face = f
	}
	return "w.Span{Face: w." + face + ", Text: " + q(e.Text()) + "}"
}

func placesec(e *El) string {
	s := placef(e)
	if e.Has("is-2nd") {
		s += ", Secondary: true"
	}
	return s
}

func icon(e *El) string {
	for _, k := range e.Els() {
		if k.Tag == "use" {
			return q(k.Attr("href")[3:])
		}
	}
	panic(fail{"icon without use"})
}

func svgs(e *El) []*El {
	var out []*El
	for _, k := range e.Els() {
		if k.Tag == "svg" {
			out = append(out, k)
		}
	}
	return out
}

// spaced is a text piece of e as drawn: runs of white space are one space,
// and a space on a side where an element stands next to it is kept — the
// browser draws it; on the element's own edge it collapses and is dropped.
// A piece that is only a space is kept between two elements. "" and false
// when nothing is left. The piece is found as the first equal text of e,
// as the Python generator found it.
func spaced(e *El, k string) (string, bool) {
	i := 0
	for j, x := range e.Kids {
		if s, ok := x.(string); ok && s == k {
			i = j
			break
		}
	}
	before, after := false, false
	for _, x := range e.Kids[:i] {
		if _, ok := x.(*El); ok {
			before = true
		}
	}
	for _, x := range e.Kids[i+1:] {
		if _, ok := x.(*El); ok {
			after = true
		}
	}
	t := collapse(k)
	if !before {
		t = strings.TrimLeft(t, " ")
	}
	if !after {
		t = strings.TrimRight(t, " ")
	}
	if strings.TrimSpace(t) == "" && !(before && after) {
		return "", false
	}
	return t, t != ""
}

func textpart(e *El, k string) string {
	if t, ok := spaced(e, k); ok {
		return "w.Text(" + q(t) + ")"
	}
	return ""
}

// say is words in two lengths (Р-65).
func (g *gen) say(e *El) string {
	f := "Full: " + q(e.Kid("tc-say__full").Text()) + ", Short: " + q(e.Kid("tc-say__short").Text())
	if e.Has("tc-say--mono") {
		f += ", Mono: true"
	}
	return "w.Say{" + f + "}"
}

func (g *gen) mono(e *El) string {
	f := "Text: " + q(e.Text())
	switch e.Tag {
	case "code":
		f += ", Form: w.MonoCode"
	case "div":
		f += ", Form: w.MonoLine"
	}
	return "w.Mono{" + f + g.ev(e, "MonoVariant", "Variant") + placef(e) + "}"
}

func (g *gen) mark(e *El) string {
	var f []string
	if c := e.Attr("data-tc-colour"); c != "" {
		f = append(f, "Colour: "+q(c))
	}
	if k := g.en.of(e, "MarkKind"); k != "" {
		f = append(f, "Kind: "+k)
	}
	if s := g.en.of(e, "MarkSize"); s != "" {
		f = append(f, "Size: "+s)
	}
	return "w.Mark{" + join(f, ", ") + "}"
}

func (g *gen) badge(e *El) string {
	f := []string{"Label: " + q(e.Text())}
	if t := g.en.of(e, "BadgeTone"); t != "" {
		f = append(f, "Tone: "+t)
	}
	if e.Has("tc-badge--edge-accent-accent") {
		f = append(f, "Accent: true")
	}
	if t := e.Attr("title"); t != "" {
		f = append(f, "Title: "+q(t))
	}
	return "w.Badge{" + join(f, ", ") + placef(e) + "}"
}

func (g *gen) tag(e *El) string {
	f := []string{"Label: " + q(e.OwnText())}
	if m := e.Kid("tc-mono"); m != nil && e.Has("tc-tag--pill") {
		f = []string{"Address: &" + g.mono(m)}
	}
	if e.Has("is-on") {
		f = append(f, "On: true")
	}
	s := svgs(e)
	if len(s) > 0 && !e.Has("tc-tag--pill") || len(s) > 1 {
		f = append(f, "Icon: "+icon(s[0]))
	}
	if m := e.Kid("tc-mark"); m != nil {
		f = append(f, "Mark: &"+g.mark(m))
	}
	if e.Has("tc-tag--pill") && len(s) > 0 {
		f = append(f, "Removable: true")
	} else if e.Has("tc-tag--pill") {
		f = append(f, "Fixed: true")
	}
	if d := e.Kid("tc-dim"); d != nil && e.Has("tc-tag--pill") {
		f = append(f, "Aside: &"+g.dimx(d))
	}
	if t := e.Attr("title"); t != "" {
		f = append(f, "Title: "+q(t))
	}
	for _, k := range e.Els() {
		if k.Tag == "input" {
			if n := k.Attr("name"); n != "" {
				f = append(f, "Name: "+q(n))
			}
			f = append(f, "Value: "+q(k.Attr("value")))
		}
	}
	return "w.Tag{" + join(f, ", ") + placef(e) + "}"
}

var dotTone = map[string]string{"accent": "Accent", "danger": "Danger", "dusty": "Dusty"}

func (g *gen) hint(k *El) string {
	s := "w.Hint{Text: " + q(k.Text())
	if t := k.Attr("title"); t != "" {
		s += ", Title: " + q(t)
	}
	return s + "}"
}

func (g *gen) inlineEl(k *El) string {
	if k.Tag == "b" && k.Has("tc-dim--accent") {
		return "w.Emphasis(" + q(k.Text()) + ")"
	}
	if k.Tag == "b" {
		if len(k.Els()) > 0 {
			return "w.Bold{Parts: []w.Inline{" + join(g.inlparts(k, g.inlineEl), ", ") + "}}"
		}
		return "w.Strong(" + q(k.Text()) + ")"
	}
	if k.Tag == "a" && k.Has("tc-dim--accent") {
		href := "#"
		if k.HasAttr("href") {
			href = k.Attr("href")
		}
		return "w.Link{Text: " + q(k.Text()) + ", Href: " + q(href) + "}"
	}
	if k.Has("tc-mono") && k.HasAttr("data-tc-slug-of") {
		return "w.SlugOf{Of: " + q(k.Attr("data-tc-slug-of")) + ", Text: " + q(k.Text()) + "}"
	}
	if k.Has("tc-mono") || k.Tag == "code" {
		return g.mono(k)
	}
	if k.Has("tc-dim__hint") {
		return g.hint(k)
	}
	if k.Has("tc-mention") {
		return "w.Mention{Name: " + q(k.Text()) + "}"
	}
	for _, c := range k.Cls() {
		if strings.HasPrefix(c, "tc-dot--") {
			return "w.Dot{Tone: w.Dot" + dotTone[c[8:]] + "}"
		}
	}
	cl := join(k.Cls(), " ")
	if k.Tag == "span" {
		if f, ok := g.faces[cl]; ok {
			return "w.Span{Face: w." + f + ", Text: " + q(k.Text()) + "}"
		}
	}
	stop("inline? " + k.Tag + " " + cl)
	return ""
}

func (g *gen) inlparts(e *El, el func(*El) string) []string {
	var out []string
	for _, k := range e.Kids {
		switch v := k.(type) {
		case string:
			if t := textpart(e, v); t != "" {
				out = append(out, t)
			}
		case *El:
			out = append(out, el(v))
		}
	}
	return out
}

// inl is running text as []w.Inline; badges and cells belong to the row or
// the card that holds the text, not to it.
func (g *gen) inl(e *El) string {
	var parts []string
	for _, p := range g.inlparts(e, func(k *El) string {
		if k.Has("tc-badge") || k.Has("tc-row__cell") {
			return ""
		}
		return g.inlineEl(k)
	}) {
		if p != "" {
			parts = append(parts, p)
		}
	}
	return "[]w.Inline{" + join(parts, ", ") + "}"
}

func (g *gen) dimparts(e *El) []string {
	var parts []string
	for _, k := range e.Kids {
		switch v := k.(type) {
		case string:
			if t := textpart(e, v); t != "" {
				parts = append(parts, t)
			}
		case *El:
			switch {
			case v.Has("tc-mono"):
				parts = append(parts, g.mono(v))
			case v.Has("tc-dim__hint"):
				parts = append(parts, g.hint(v))
			case v.Tag == "b":
				parts = append(parts, "w.Strong("+q(v.Text())+")")
			default:
				parts = append(parts, g.inlineEl(v))
			}
		}
	}
	return parts
}

func (g *gen) dimx(e *El) string {
	empty := ""
	if e.HasAttr("data-tc-empty") {
		empty = ", Empty: true"
	}
	return "w.Dim{Parts: []w.Inline{" + join(g.dimparts(e), ", ") + "}" + g.ev(e, "DimSize", "Size") + g.ev(e, "DimVariant", "Variant") + placef(e) + empty + "}"
}

// unq reverses q.
func unq(s string) string {
	s = s[1 : len(s)-1]
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' && i+1 < len(s) {
			i++
			switch s[i] {
			case 'n':
				b.WriteByte('\n')
			case 't':
				b.WriteByte('\t')
			default:
				b.WriteByte(s[i])
			}
			continue
		}
		b.WriteByte(s[i])
	}
	return b.String()
}

// trimEnds strips the outer ends of the first and last text parts of a list
// ("w.Text(…)" items), as the record generators did.
func trimEnds(parts []string) []string {
	if len(parts) > 0 && strings.HasPrefix(parts[0], "w.Text(") {
		parts[0] = "w.Text(" + q(strings.TrimLeftFunc(unq(parts[0][7:len(parts[0])-1]), isSpace)) + ")"
	}
	if n := len(parts); n > 0 && strings.HasPrefix(parts[n-1], "w.Text(") {
		parts[n-1] = "w.Text(" + q(strings.TrimRightFunc(unq(parts[n-1][7:len(parts[n-1])-1]), isSpace)) + ")"
	}
	return parts
}

// tparts is the running text of a block as a list of parts of type typ.
func (g *gen) tparts(e *El, typ string) string {
	type part struct {
		text bool
		v    string
	}
	var parts []part
	for _, k := range e.Kids {
		switch v := k.(type) {
		case string:
			if t := collapse(v); t != "" {
				parts = append(parts, part{true, t})
			}
		case *El:
			var x string
			switch {
			case v.Tag == "br":
				x = "w.Break{}"
			case v.Tag == "b" && v.Has("tc-card__key"):
				x = "w.Key(" + q(v.Text()) + ")"
			case v.Tag == "b" && len(v.Els()) > 0:
				x = "w.Bold{Parts: " + g.tparts(v, "[]w.Inline{") + "}"
			case v.Tag == "b":
				x = "w.Strong(" + q(v.Text()) + ")"
			case v.Tag == "a" && v.Has("tc-btn"):
				// a button that leads: Try again of a message
				x = g.button(v)
			case v.Tag == "a":
				x = "w.Link{Text: " + q(v.Text()) + `, Href: "#"}`
			case v.Has("tc-mono") || v.Tag == "code":
				x = g.mono(v)
			case v.Has("tc-dim--accent"):
				x = g.span(v, "FaceAccent")
			case v.Has("tc-dim--danger"):
				x = g.span(v, "FaceDanger")
			case v.Has("tc-dim"):
				x = g.dimx(v)
			case v.Has("tc-badge"):
				x = g.badge(v)
			case v.Has("tc-kbd"):
				x = g.kbd(v)
			case v.Has("tc-btn"):
				x = g.button(v)
			case v.Has("tc-mark"):
				x = g.mark(v)
			case v.Has("tc-tag"):
				x = g.tag(v)
			default:
				stop("text? " + v.Tag + " " + join(v.Cls(), " "))
			}
			parts = append(parts, part{false, x})
		}
	}
	if len(parts) > 0 && parts[0].text {
		parts[0].v = strings.TrimLeftFunc(parts[0].v, isSpace)
	}
	if n := len(parts); n > 0 && parts[n-1].text {
		parts[n-1].v = strings.TrimRightFunc(parts[n-1].v, isSpace)
	}
	var out []string
	for _, p := range parts {
		if p.v == "" {
			continue
		}
		if p.text {
			out = append(out, "w.Text("+q(p.v)+")")
		} else {
			out = append(out, p.v)
		}
	}
	return typ + join(out, ", ") + "}"
}

// inl2 is running text of a document's line: bold, links, monospaced, accent.
func (g *gen) inl2(e *El) string {
	type part struct {
		text bool
		v    string
	}
	var parts []part
	for _, k := range e.Kids {
		switch v := k.(type) {
		case string:
			if t := collapse(v); t != "" {
				parts = append(parts, part{true, t})
			}
		case *El:
			var x string
			switch {
			case v.Tag == "b":
				x = "w.Strong(" + q(v.Text()) + ")"
			case v.Tag == "em":
				x = "w.Em(" + q(v.Text()) + ")"
			case v.Has("tc-doc__bullets"):
				x = "w.DocBullets{Items: [][]w.Inline{" + g.each(v.Els(), g.inl2) + "}}"
			case v.Tag == "a":
				x = "w.Link{Text: " + q(v.Text()) + `, Href: "#"}`
			case v.Has("tc-mono"):
				x = g.mono(v)
			case v.Has("tc-dim--accent"):
				x = g.span(v, "FaceAccent")
			default:
				stop("inline? " + join(v.Cls(), " "))
			}
			parts = append(parts, part{false, x})
		}
	}
	if len(parts) > 0 && parts[0].text {
		parts[0].v = strings.TrimLeftFunc(parts[0].v, isSpace)
	}
	if n := len(parts); n > 0 && parts[n-1].text {
		parts[n-1].v = strings.TrimRightFunc(parts[n-1].v, isSpace)
	}
	var out []string
	for _, p := range parts {
		if p.v == "" {
			continue
		}
		if p.text {
			out = append(out, "w.Text("+q(p.v)+")")
		} else {
			out = append(out, p.v)
		}
	}
	return "[]w.Inline{" + join(out, ", ") + "}"
}

// rawparts is a paragraph's text as written: its breaks are the text's.
func (g *gen) rawparts(e *El) string {
	var parts []string
	for _, k := range e.Kids {
		switch v := k.(type) {
		case string:
			if v != "" {
				parts = append(parts, "w.Text("+q(v)+")")
			}
		case *El:
			switch {
			case v.Tag == "b":
				parts = append(parts, "w.Strong("+q(v.Text())+")")
			case v.Tag == "em":
				parts = append(parts, "w.Em("+q(v.Text())+")")
			case v.Tag == "a":
				parts = append(parts, "w.Link{Text: "+q(v.Text())+`, Href: "#"}`)
			case v.Has("tc-mono"):
				parts = append(parts, g.mono(v))
			case v.Has("tc-mention"):
				parts = append(parts, "w.Mention{Name: "+q(v.Text())+"}")
			default:
				stop("prose? " + join(v.Cls(), " "))
			}
		}
	}
	return "[]w.Inline{" + join(parts, ", ") + "}"
}

var labelRe = regexp.MustCompile(`Label: "[^"]*"`)

func isSpace(r rune) bool { return unicode.IsSpace(r) }
