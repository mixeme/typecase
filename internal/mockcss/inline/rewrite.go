package main

import (
	"fmt"
	"sort"
	"strings"
)

type stats struct {
	classN    map[string]int
	gapOK     map[int]bool
	leftoverN map[string]map[string]int // class → leftover style → frame-count-ish
}

type usedCSS struct {
	glue      map[string]bool
	gapsLine  map[int]bool
	gapsStack map[int]bool
	colours   map[string]bool
	widths    map[string]bool
	cols      map[string]string
	tints     map[string]bool
	variants  map[string]string
	folded    map[string]string
	named     map[string]string
	seq       map[string]int
}

func newUsed() usedCSS {
	return usedCSS{
		glue: map[string]bool{}, gapsLine: map[int]bool{}, gapsStack: map[int]bool{},
		colours: map[string]bool{}, widths: map[string]bool{},
		cols: map[string]string{}, tints: map[string]bool{},
		variants: map[string]string{}, folded: map[string]string{},
		named: map[string]string{}, seq: map[string]int{},
	}
}

type auditRow struct {
	Style, Class, Frames, Action string
	N                            int
}

func buildStats(hits []hit, classN map[string]int, gapFrames map[int]map[string]bool) stats {
	st := stats{
		classN:    classN,
		gapOK:     map[int]bool{},
		leftoverN: map[string]map[string]int{},
	}
	for n, frames := range gapFrames {
		if len(frames) >= 2 {
			st.gapOK[n] = true
		}
	}
	return st
}

var tintStroke = map[string]string{
	"var(--tc-accent)": "tc-ico--accent",
	"var(--tc-danger)": "tc-ico--danger",
	"var(--tc-dusty)":  "tc-ico--dusty",
	"var(--tc-muted)":  "tc-ico--muted",
	"var(--alert)":     "tc-ico--danger",
	"var(--accent)":    "tc-ico--accent",
}

var colByWidth = map[string]string{
	"28px":  "tc-table__col--tick",
	"20px":  "tc-table__col--tick",
	"92px":  "tc-table__col--size",
	"132px": "tc-table__col--type",
	"150px": "tc-table__col--time",
	"74px":  "tc-table__col--acts",
	"70px":  "tc-table__col--num",
	"26px":  "tc-table__col--acts",
	"140px": "tc-table__col--time",
}

var colWidthCSS = map[string]string{
	"tc-table__col--tick": "width:28px",
	"tc-table__col--size": "width:92px",
	"tc-table__col--type": "width:132px",
	"tc-table__col--time": "width:150px",
	"tc-table__col--acts": "width:74px",
	"tc-table__col--num":  "width:70px",
}

var namedVariant = []struct {
	needClass string
	style     string
	add       string
	css       string
}{
	{"tc-head", "padding-bottom:8px", "tc-head--snug", "padding-bottom:8px"},
	{"tc-head", "align-items:center", "tc-head--mid", "align-items:center"},
	{"tc-card", "border-color:var(--tc-danger)", "tc-card--alert", "border-color:var(--tc-danger)"},
	{"tc-sheet", "width:560px", "", ""}, // already in .tc-sheet
	{"tc-dim", "color:var(--tc-dusty)", "", "color:var(--tc-dusty)"},
	{"tc-meter__bar", "background:var(--tc-danger)", "tc-meter__bar--alert", "background:var(--tc-danger)"},
}

func skipTag(t tagOpen) bool {
	cs := t.class()
	if hasClassName(cs, "frame") || hasClassName(cs, "frame-label") {
		return true
	}
	if t.name == "svg" || t.name == "th" || t.name == "td" || t.name == "i" {
		return false
	}
	if len(cs) == 0 {
		return false
	}
	for _, c := range cs {
		if strings.HasPrefix(c, "tc-") || strings.HasPrefix(c, "is-") || strings.HasPrefix(c, "mk-") {
			return false
		}
	}
	return true
}

func rewriteTag(t tagOpen, st stats, u *usedCSS) tagOpen {
	cs := t.class()
	style, _ := t.get("style")
	ds := parseDecls(style)

	if c, ok := hexFromDecls(ds); ok {
		t.set("data-tc-colour", c)
		u.colours[c] = true
		_, ds, _ = takeProp(ds, "background")
		_, ds, _ = takeProp(ds, "background-color")
	}
	if w, rest, ok := takePctWidth(ds, cs); ok {
		t.set("data-tc-width", w)
		u.widths[w] = true
		ds = rest
	}

	ds = peelGlue(&t, ds, cs, st, u)
	cs = t.class()

	if t.name == "th" || t.name == "td" {
		ds = peelColumn(&t, ds, u)
		cs = t.class()
	}

	if stroke, ok := t.get("stroke"); ok {
		if cls, ok := tintStroke[stroke]; ok {
			cs = addClass(cs, cls)
			u.tints[cls] = true
			t.del("stroke")
		}
	}
	if fill, ok := t.get("fill"); ok {
		if cls, ok := tintStroke[fill]; ok {
			cs = addClass(cs, cls+"-fill")
			u.tints[cls+"-fill"] = true
			t.del("fill")
		}
	}
	t.setClass(cs)

	ds = peelNamed(&t, ds, cs, u)
	cs = t.class()
	ds = peelLeftover(&t, ds, cs, st, u)
	cs = t.class()

	t.setClass(cs)
	if len(ds) == 0 {
		t.del("style")
	} else {
		t.set("style", joinDecls(ds))
	}
	return t
}

func hexFromDecls(ds []decl) (string, bool) {
	for _, d := range ds {
		if d.prop == "background" || d.prop == "background-color" {
			if c, ok := hexColour(d.val); ok {
				return c, true
			}
		}
	}
	return "", false
}

func takePctWidth(ds []decl, cs []string) (string, []decl, bool) {
	if !hasClassName(cs, "tc-meter") && !hasClassName(cs, "tc-meter__bar") {
		return "", ds, false
	}
	w, rest, ok := takeProp(ds, "width")
	if !ok {
		return "", ds, false
	}
	pct, ok := widthPct(w)
	if !ok {
		return "", ds, false
	}
	return pct, rest, true
}

func peelGlue(t *tagOpen, ds []decl, cs []string, st stats, u *usedCSS) []decl {
	got, rest := takeProps(ds,
		"display", "flex-direction", "align-items", "flex-wrap",
		"gap", "margin-left", "padding", "flex", "width",
		"grid-template-columns",
	)
	if got["margin-left"] == "auto" {
		cs = addClass(cs, "tc-push")
		u.glue["tc-push"] = true
		delete(got, "margin-left")
	}
	if got["padding"] == "0 20px" || got["padding"] == "0 var(--tc-page-gutter)" {
		cs = addClass(cs, "tc-pad")
		u.glue["tc-pad"] = true
		delete(got, "padding")
	}
	if got["flex"] == "1 1 100%" || got["flex"] == "1" && got["width"] == "100%" {
		cs = addClass(cs, "tc-grow")
		u.glue["tc-grow"] = true
		delete(got, "flex")
		delete(got, "width")
	} else if got["flex"] == "1 1 100%" {
		cs = addClass(cs, "tc-grow")
		u.glue["tc-grow"] = true
		delete(got, "flex")
	}

	disp := got["display"]
	dir := got["flex-direction"]
	if disp == "flex" && dir == "column" {
		cs = addClass(cs, "tc-stack")
		u.glue["tc-stack"] = true
		delete(got, "display")
		delete(got, "flex-direction")
		if n, ok := gapPx(got["gap"]); ok && st.gapOK[n] {
			mod := fmt.Sprintf("tc-stack--gap-%d", n)
			cs = addClass(cs, mod)
			u.gapsStack[n] = true
			delete(got, "gap")
		}
	} else if disp == "flex" && (dir == "" || dir == "row") {
		align := got["align-items"]
		if align == "" || align == "center" {
			cs = addClass(cs, "tc-line")
			u.glue["tc-line"] = true
			delete(got, "display")
			delete(got, "flex-direction")
			if align == "center" {
				delete(got, "align-items")
			}
			if got["flex-wrap"] == "wrap" {
				delete(got, "flex-wrap")
			}
			if n, ok := gapPx(got["gap"]); ok && st.gapOK[n] {
				mod := fmt.Sprintf("tc-line--gap-%d", n)
				cs = addClass(cs, mod)
				u.gapsLine[n] = true
				delete(got, "gap")
			}
		}
	} else if disp == "grid" {
		cols := got["grid-template-columns"]
		switch cols {
		case "minmax(0,1fr) minmax(0,1fr)", "1fr 1fr":
			cs = addClass(cs, "tc-cols--halves")
			u.glue["tc-cols--halves"] = true
			delete(got, "display")
			delete(got, "grid-template-columns")
		case "minmax(0,1fr) minmax(0,1fr) minmax(0,1fr)", "1fr 1fr 1fr":
			cs = addClass(cs, "tc-cols--thirds")
			u.glue["tc-cols--thirds"] = true
			delete(got, "display")
			delete(got, "grid-template-columns")
		case "minmax(0,1fr) 296px", "1fr 296px":
			cs = addClass(cs, "tc-cols--sidebar")
			u.glue["tc-cols--sidebar"] = true
			delete(got, "display")
			delete(got, "grid-template-columns")
		}
	}

	t.setClass(cs)
	return appendMap(rest, got)
}

func appendMap(rest []decl, got map[string]string) []decl {
	keys := make([]string, 0, len(got))
	for k := range got {
		if got[k] != "" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for _, k := range keys {
		rest = append(rest, decl{k, got[k]})
	}
	return rest
}

func peelColumn(t *tagOpen, ds []decl, u *usedCSS) []decl {
	w, rest, ok := takeProp(ds, "width")
	if !ok {
		return ds
	}
	cls, ok := colByWidth[w]
	if !ok {
		cls = "tc-table__col--" + widthSlug(w)
		u.cols[cls] = "width:" + w
	} else {
		if w == "20px" {
			cls = "tc-table__col--tick"
			u.cols["tc-table__col--tick"] = "width:28px"
			u.cols["tc-table__col--tick-sm"] = "width:20px"
			cls = "tc-table__col--tick-sm"
		} else if w == "26px" {
			cls = "tc-table__col--menu"
			u.cols[cls] = "width:26px"
		} else if w == "140px" {
			cls = "tc-table__col--changed"
			u.cols[cls] = "width:140px"
		} else {
			u.cols[cls] = colWidthCSS[cls]
		}
	}
	cs := addClass(t.class(), cls)
	t.setClass(cs)
	return rest
}

func widthSlug(w string) string {
	w = strings.TrimSuffix(w, "%")
	w = strings.ReplaceAll(w, "px", "")
	w = strings.ReplaceAll(w, ".", "p")
	return "w" + w
}

func peelNamed(t *tagOpen, ds []decl, cs []string, u *usedCSS) []decl {
	left := joinDecls(ds)
	for _, n := range namedVariant {
		if !hasClassName(cs, n.needClass) {
			continue
		}
		if left != n.style {
			continue
		}
		if n.add != "" {
			cs = addClass(cs, n.add)
			u.variants[n.add] = n.css
		} else if n.css != "" {
			u.folded["."+n.needClass] = n.css
		}
		t.setClass(cs)
		return nil
	}
	return ds
}

var glueClass = map[string]bool{
	"tc-line": true, "tc-stack": true, "tc-push": true, "tc-pad": true, "tc-grow": true,
}

func isGlueName(c string) bool {
	if glueClass[c] {
		return true
	}
	if strings.HasPrefix(c, "tc-line--gap-") || strings.HasPrefix(c, "tc-stack--gap-") {
		return true
	}
	if strings.HasPrefix(c, "tc-cols--") || c == "tc-cols__col" {
		return true
	}
	return false
}

func peelLeftover(t *tagOpen, ds []decl, cs []string, st stats, u *usedCSS) []decl {
	if len(ds) == 0 {
		return ds
	}
	left := joinDecls(ds)
	if hasClassName(cs, "tc-dim") && left == "color:var(--tc-dusty)" {
		u.folded[".tc-dim"] = left
		t.setClass(cs)
		return nil
	}
	base := primaryClass(cs)
	if base == "" || isGlueName(base) {
		if t.name == "svg" {
			base = "tc-ico"
		} else {
			base = "tc-" + t.name
		}
	}
	mod := leftoverClass(base, left, u)
	cs = addClass(cs, mod)
	u.variants[mod] = left
	t.setClass(cs)
	return nil
}

func leftoverClass(base, left string, u *usedCSS) string {
	key := base + "\x00" + left
	if name, ok := u.named[key]; ok {
		return name
	}
	u.seq[base]++
	name := fmt.Sprintf("%s--r%d", base, u.seq[base])
	for hasVariant(u, name) {
		u.seq[base]++
		name = fmt.Sprintf("%s--r%d", base, u.seq[base])
	}
	u.named[key] = name
	return name
}

func hasVariant(u *usedCSS, name string) bool {
	_, ok := u.variants[name]
	return ok
}

func rewriteHTML(html string, st stats) (string, usedCSS) {
	u := newUsed()
	var b strings.Builder
	last := 0
	walkProductTags(html, func(start, end int, t tagOpen, c walkCtx) {
		if !product(c) || skipTag(t) {
			return
		}
		style, _ := t.get("style")
		stroke, _ := t.get("stroke")
		fill, _ := t.get("fill")
		if style == "" && !strings.HasPrefix(stroke, "var(") && !strings.HasPrefix(fill, "var(") {
			return
		}
		nt := rewriteTag(t, st, &u)
		b.WriteString(html[last:start])
		b.WriteString(nt.String())
		last = end + 1
	})
	b.WriteString(html[last:])
	return b.String(), u
}

func injectCSS(html string, u usedCSS) string {
	html = injectInto(html, "<style data-tc>", "</style>", productCSS(u), "/* layout glue */")
	html = injectInto(html, "<style>", "</style>", pageDataCSS(u), "/* frame data values */")
	return html
}

func injectInto(html, open, close, css, marker string) string {
	if strings.Contains(html, marker) {
		start := strings.Index(html, marker)
		end := strings.Index(html[start:], close)
		if end < 0 {
			return html
		}
		return html[:start] + css + html[start+end:]
	}
	low := strings.ToLower(html)
	from := 0
	if open == "<style data-tc>" {
		from = strings.Index(low, "<style data-tc")
	} else {
		from = strings.Index(low, "<style>")
		if from < 0 {
			from = strings.Index(low, "<style ")
		}
	}
	if from < 0 {
		return html
	}
	gt := strings.Index(html[from:], ">")
	endRel := strings.Index(low[from+gt+1:], close)
	if gt < 0 || endRel < 0 {
		return html
	}
	at := from + gt + 1 + endRel
	if css == "" {
		return html
	}
	return html[:at] + "\n" + css + html[at:]
}

// doubled is (0,2,0): leftover used to be an inline style and must beat
// two-class product rules (.tc-in.is-area, .tc-btn.is-primary, …).
func doubled(class string) string {
	return "." + class + "." + class
}

func productCSS(u usedCSS) string {
	var b strings.Builder
	b.WriteString("/* layout glue */\n")
	if u.glue["tc-line"] {
		fmt.Fprintf(&b, "%s{display:flex;align-items:center;flex-wrap:wrap}\n", doubled("tc-line"))
	}
	if u.glue["tc-stack"] {
		fmt.Fprintf(&b, "%s{display:flex;flex-direction:column}\n", doubled("tc-stack"))
	}
	for _, n := range sortedInts(u.gapsLine) {
		fmt.Fprintf(&b, "%s{gap:%dpx}\n", doubled(fmt.Sprintf("tc-line--gap-%d", n)), n)
	}
	for _, n := range sortedInts(u.gapsStack) {
		fmt.Fprintf(&b, "%s{gap:%dpx}\n", doubled(fmt.Sprintf("tc-stack--gap-%d", n)), n)
	}
	if u.glue["tc-push"] {
		fmt.Fprintf(&b, "%s{margin-left:auto}\n", doubled("tc-push"))
	}
	if u.glue["tc-pad"] {
		fmt.Fprintf(&b, "%s{padding:0 var(--tc-page-gutter)}\n", doubled("tc-pad"))
	}
	if u.glue["tc-grow"] {
		fmt.Fprintf(&b, "%s{flex:1 1 100%%;width:100%%;min-width:0}\n", doubled("tc-grow"))
	}
	if u.glue["tc-cols--halves"] {
		fmt.Fprintf(&b, "%s{display:grid;grid-template-columns:minmax(0,1fr) minmax(0,1fr);gap:14px;container:tcbox / inline-size}\n", doubled("tc-cols--halves"))
	}
	if u.glue["tc-cols--thirds"] {
		fmt.Fprintf(&b, "%s{display:grid;grid-template-columns:minmax(0,1fr) minmax(0,1fr) minmax(0,1fr);gap:14px;container:tcbox / inline-size}\n", doubled("tc-cols--thirds"))
	}
	if u.glue["tc-cols--sidebar"] {
		fmt.Fprintf(&b, "%s{display:grid;grid-template-columns:minmax(0,1fr) 296px;gap:14px;container:tcbox / inline-size}\n", doubled("tc-cols--sidebar"))
	}
	if u.glue["tc-cols--halves"] || u.glue["tc-cols--thirds"] || u.glue["tc-cols--sidebar"] {
		fmt.Fprintf(&b, "%s{container:tcbox / inline-size;min-width:0}\n", doubled("tc-cols__col"))
		b.WriteString("@container tcbox (max-width:620px){.tc-cols--halves,.tc-cols--thirds,.tc-cols--sidebar{grid-template-columns:minmax(0,1fr)}}\n")
	}
	if len(u.colours) > 0 {
		b.WriteString("[data-tc-colour][data-tc-colour]{background:var(--tc-colour)}\n")
	}
	if len(u.widths) > 0 {
		b.WriteString("[data-tc-width][data-tc-width]{width:var(--tc-width)}\n")
	}
	for _, cls := range sortedKeys(u.cols) {
		fmt.Fprintf(&b, "%s{%s}\n", doubled(cls), u.cols[cls])
	}
	for _, cls := range sortedKeysBool(u.tints) {
		switch cls {
		case "tc-ico--accent":
			fmt.Fprintf(&b, "%s{stroke:var(--tc-accent)}\n", doubled(cls))
		case "tc-ico--danger":
			fmt.Fprintf(&b, "%s{stroke:var(--tc-danger)}\n", doubled(cls))
		case "tc-ico--dusty":
			fmt.Fprintf(&b, "%s{stroke:var(--tc-dusty)}\n", doubled(cls))
		case "tc-ico--muted":
			fmt.Fprintf(&b, "%s{stroke:var(--tc-muted)}\n", doubled(cls))
		case "tc-ico--accent-fill":
			fmt.Fprintf(&b, "%s{fill:var(--tc-accent)}\n", doubled(cls))
		case "tc-ico--danger-fill":
			fmt.Fprintf(&b, "%s{fill:var(--tc-danger)}\n", doubled(cls))
		}
	}
	for _, cls := range sortedKeys(u.variants) {
		fmt.Fprintf(&b, "%s{%s}\n", doubled(cls), u.variants[cls])
	}
	for _, sel := range sortedKeys(u.folded) {
		fmt.Fprintf(&b, "%s{%s}\n", sel, u.folded[sel])
	}
	return b.String()
}

func pageDataCSS(u usedCSS) string {
	var b strings.Builder
	b.WriteString("/* frame data values */\n")
	for _, c := range sortedKeysBool(u.colours) {
		fmt.Fprintf(&b, "[data-tc-colour=%q]{--tc-colour:%s}\n", c, c)
	}
	for _, w := range sortedKeysBool(u.widths) {
		fmt.Fprintf(&b, "[data-tc-width=%q]{--tc-width:%s}\n", w, w)
	}
	return b.String()
}

func sortedKeys(m map[string]string) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func sortedKeysBool(m map[string]bool) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func sortedInts(m map[int]bool) []int {
	out := make([]int, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Ints(out)
	return out
}

func apply(html string) (string, []auditRow, error) {
	hits, classN, gapFrames := collectHits(html)
	st := buildStats(hits, classN, gapFrames)
	rows := auditFromHits(hits, st)
	out, u := rewriteHTML(html, st)
	if out == html && strings.Contains(html, "/* layout glue */") {
		return html, rows, nil
	}
	out = injectCSS(out, u)
	if out == "" {
		return "", nil, fmt.Errorf("inline: empty rewrite")
	}
	return out, rows, nil
}

func auditFromHits(hits []hit, st stats) []auditRow {
	type key struct{ style, class string }
	type acc struct {
		n      int
		frames map[string]bool
	}
	g := map[key]*acc{}
	for _, h := range hits {
		k := key{h.Style, h.Class}
		if g[k] == nil {
			g[k] = &acc{frames: map[string]bool{}}
		}
		g[k].n++
		if h.Frame != "" {
			g[k].frames[h.Frame] = true
		}
	}
	var rows []auditRow
	for k, a := range g {
		rows = append(rows, auditRow{
			Style:  k.style,
			Class:  k.class,
			N:      a.n,
			Frames: joinFrames(a.frames),
			Action: proposeAction(k.style, k.class, a.n, len(a.frames), st),
		})
	}
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].N != rows[j].N {
			return rows[i].N > rows[j].N
		}
		return rows[i].Style < rows[j].Style
	})
	return rows
}

func joinFrames(m map[string]bool) string {
	var s []string
	for k := range m {
		s = append(s, k)
	}
	sort.Strings(s)
	if len(s) > 8 {
		return strings.Join(s[:8], ", ") + "…"
	}
	return strings.Join(s, ", ")
}

func proposeAction(style, class string, n, frames int, st stats) string {
	ds := parseDecls(style)
	if c, ok := hexFromDecls(ds); ok {
		return "данные data-tc-colour=" + c
	}
	if w, _, ok := takePctWidth(ds, nil); ok {
		return "данные data-tc-width=" + w
	}
	got, _ := takeProps(ds, "display", "flex-direction", "margin-left", "padding", "gap")
	if got["display"] == "flex" || got["margin-left"] == "auto" || got["padding"] == "0 20px" {
		return "клей раскладки"
	}
	if strings.HasPrefix(style, "width:") && (strings.Contains(class, "tc-table") || class == "") {
		return "колонка таблицы"
	}
	if frames >= 2 {
		return "вариант блока (≥2 кадра)"
	}
	return "разовая, Р-13"
}

func formatAudit(rows []auditRow) string {
	var b strings.Builder
	b.WriteString("# Inline-стили макета\n\n")
	b.WriteString("Сгенерировано `go run ./internal/mockcss/inline`.\n\n")
	sum := 0
	for _, r := range rows {
		sum += r.N
	}
	fmt.Fprintf(&b, "Атрибутов `style=` в кадрах и демо: **%d**. Разных строк: **%d**.\n\n", sum, len(rows))
	b.WriteString("| раз | style | класс | кадры | действие |\n")
	b.WriteString("|----:|-------|-------|-------|----------|\n")
	for _, r := range rows {
		style := strings.ReplaceAll(r.Style, "|", "\\|")
		if style == "" {
			style = "*(svg var())*"
		}
		fmt.Fprintf(&b, "| %d | `%s` | `%s` | %s | %s |\n", r.N, style, r.Class, r.Frames, r.Action)
	}
	return b.String()
}
