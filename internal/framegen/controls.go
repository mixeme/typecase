package framegen

import (
	"regexp"
	"strings"
)

// kv is one entry of an ordered table: a class and what it gives.
type kv struct{ k, v string }

// firstOf is the value of the first entry whose class e has.
func firstOf(e *El, table []kv) string {
	for _, x := range table {
		if e.Has(x.k) {
			return x.v
		}
	}
	return ""
}

// allOf is the values of every entry whose class e has, in table order.
func allOf(e *El, table []kv) []string {
	var out []string
	for _, x := range table {
		if e.Has(x.k) {
			out = append(out, x.v)
		}
	}
	return out
}

func place(e *El, f []string) []string {
	if e.Has("tc-push") {
		f = append(f, "Place: w.Pushed")
	}
	if e.Has("tc-grow") {
		f = append(f, "Place: w.Grown")
	}
	if e.Has("is-2nd") {
		f = append(f, "Secondary: true")
	}
	return f
}

var (
	iconBox    = map[string]string{"tc-ico--sq16": "IconBox16", "tc-ico--sq17": "IconBox17", "tc-ico--sq19": "IconBox19"}
	buttonKind = []kv{{"is-primary", "ButtonPrimary"}, {"is-danger", "ButtonDanger"}, {"is-danger-solid", "ButtonDangerSolid"}}
	buttonPad  = []kv{{"tc-btn--padx3", "ButtonPadX3"}, {"tc-btn--padx4", "ButtonPadX4"}, {"tc-btn--padx6", "ButtonPadX6"},
		{"tc-btn--padx8", "ButtonPadX8"}, {"tc-btn--on-card", "ButtonOnCard"}}
	iconNames = map[string]string{"down": "Download", "close": "Close", "back": "Back", "quill": "New note", "menu": "Menu",
		"search": "Search", "link": "Link", "up": "Upload", "refresh": "Refresh", "move": "Move"}
	iconRe = regexp.MustCompile(`Icon: "([a-z-]+)"`)
)

// button0 is a button as its label, icons and classes give it.
func (g *gen) button0(e *El) string {
	var f []string
	ss := svgs(e)
	label := e.OwnText()
	if lab := e.Kid("tc-btn__label"); lab != nil {
		label = lab.Text()
		f = append(f, "Folds: true")
	}
	if label != "" {
		f = append(f, "Label: "+q(label))
	}
	var icons []string
	for _, s := range ss {
		icons = append(icons, icon(s))
	}
	if n := len(icons); n > 0 && icons[n-1] == `"chev"` && (label != "" || n > 1) {
		f = append(f, "Menu: true")
		icons = icons[:n-1]
	}
	if len(icons) > 0 {
		f = append(f, "Icon: "+icons[0])
		for _, c := range ss[0].Cls() {
			if b, ok := iconBox[c]; ok {
				f = append(f, "IconBox: w."+b)
			}
		}
	}
	for _, v := range allOf(e, buttonKind) {
		f = append(f, "Kind: w."+v)
	}
	if e.Has("is-sm") {
		f = append(f, "Small: true")
	}
	if e.Has("is-focus") {
		f = append(f, "Focus: true")
	}
	for _, v := range allOf(e, buttonPad) {
		f = append(f, "Variant: w."+v)
	}
	t := e.Attr("title")
	if t == "" && label == "" {
		t = "More"
	}
	if t != "" {
		f = append(f, "Title: "+q(t))
	}
	f = place(e, f)
	return "w.Button{" + join(f, ", ") + "}"
}

func dropMore(s string) string {
	s = strings.Replace(s, `Title: "More", `, "", -1)
	return strings.Replace(s, `, Title: "More"`, "", -1)
}

func (g *gen) button1(e *El) string {
	if e.Has("tc-close") {
		lab := e.Attr("aria-label")
		if lab == "" {
			lab = e.Attr("title")
		}
		if lab == "" {
			lab = "Close"
		}
		return "w.Close{Label: " + q(lab) + placef(e) + "}"
	}
	s := g.button0(e)
	var extra []string
	if v := g.en.of(e, "ButtonVariant"); v != "" && !strings.Contains(s, "Variant:") {
		extra = append(extra, "Variant: "+v)
	}
	if e.Has("is-off") {
		extra = append(extra, "State: w.ButtonOff")
	}
	if e.Valued("disabled") {
		extra = append(extra, "State: w.ButtonDisabled")
	}
	if e.Has("is-3rd") {
		extra = append(extra, "Tertiary: true")
	}
	var bs []*El
	for _, k := range e.Els() {
		if k.Tag == "b" {
			bs = append(bs, k)
		}
	}
	if len(bs) > 0 && e.OwnText() == "" {
		s = dropMore(strings.Replace(s, "w.Button{", "w.Button{Label: "+q(bs[0].Text())+", Bold: true, ", 1))
	}
	if len(extra) > 0 {
		s = s[:len(s)-1] + ", " + join(extra, ", ") + "}"
	}
	return s
}

func index(kids []any, x *El) int {
	for i, k := range kids {
		if k == any(x) {
			return i
		}
	}
	return -1
}

func strs(kids []any) string {
	var b strings.Builder
	for _, k := range kids {
		if s, ok := k.(string); ok {
			b.WriteString(s)
		}
	}
	return b.String()
}

func (g *gen) button2(e *El) string {
	if e.Has("tc-close") {
		return g.button1(e)
	}
	s := g.button1(e)
	var extra []string
	if e.HasAttr("disabled") && !strings.Contains(s, "State:") {
		extra = append(extra, "State: w.ButtonDisabled")
	}
	tick := e.Kid("tc-tick")
	if tick != nil {
		d := ""
		if tick.Has("is-done") {
			d = "Done: true, "
		}
		extra = append(extra, "Tick: &w.Tick{"+d+"Glyph: true"+g.ev(tick, "TickVariant", "Variant")+"}")
	}
	m := e.Kid("tc-mono")
	if m != nil {
		extra = append(extra, "Aside: &"+g.mono(m))
		i := index(e.Kids, m)
		if after := norm(strs(e.Kids[i+1:])); after != "" {
			extra = append(extra, "After: "+q(after))
		}
		if before := norm(strs(e.Kids[:i])); before != "" {
			if loc := labelRe.FindStringIndex(s); loc != nil {
				s = s[:loc[0]] + "Label: " + q(before) + s[loc[1]:]
			}
		}
	}
	if tick != nil || m != nil {
		s = dropMore(s)
	}
	if len(extra) > 0 {
		s = s[:len(s)-1] + ", " + join(extra, ", ") + "}"
	}
	return s
}

// button is a button; an icon button the frame leaves unnamed is named by
// its icon.
func (g *gen) button(e *El) string {
	s := g.button2(e)
	m := iconRe.FindStringSubmatch(s)
	if strings.Contains(s, `Title: "More"`) && m != nil && m[1] != "dots" {
		name, ok := iconNames[m[1]]
		if !ok {
			name = strings.ToUpper(m[1][:1]) + m[1][1:]
		}
		s = strings.Replace(s, `Title: "More"`, "Title: "+q(name), -1)
	}
	return s
}

func (g *gen) buttons(es []*El) string {
	var out []string
	for _, b := range es {
		out = append(out, g.button(b))
	}
	return join(out, ", ")
}

var inputKind = map[string]string{"date": "InputDate", "time": "InputTime", "search": "InputSearch"}

func (g *gen) input(x *El) string {
	var f []string
	t := "text"
	if x.HasAttr("type") {
		t = x.Attr("type")
	}
	if k, ok := inputKind[t]; ok {
		f = append(f, "Kind: w."+k)
	}
	for _, a := range []kv{{"name", "Name"}, {"value", "Value"}, {"placeholder", "Placeholder"}, {"aria-label", "Label"}} {
		if v := x.Attr(a.k); v != "" {
			f = append(f, a.v+": "+q(v))
		}
	}
	for _, c := range []kv{{"is-mono", "Mono"}, {"is-focus", "Focus"}, {"is-error", "Error"}} {
		if x.Has(c.k) {
			f = append(f, c.v+": true")
		}
	}
	if x.Attr("inputmode") == "numeric" {
		f = append(f, "Numeric: true")
	}
	return "w.Input{" + join(f, ", ") + g.ev(x, "InputVariant", "Variant") + placef(x) + "}"
}

func (g *gen) pick(x *El) string {
	f := []string{"Value: " + q(x.OwnText())}
	if d := x.Kid("tc-dim"); d != nil {
		f = append(f, "Aside: &"+g.dimx(d))
	}
	if x.Has("is-ph") {
		f = append(f, "Unset: true")
	}
	if v := x.Attr("aria-label"); v != "" {
		f = append(f, "Label: "+q(v))
	}
	if x.Has("is-focus") {
		f = append(f, "Focus: true")
	}
	return "w.Pick{" + join(f, ", ") + g.ev(x, "PickVariant", "Variant") + placef(x) + "}"
}

func queryOf(qq *El) string {
	return "w.Query{Value: " + q(qq.Attr("value")) + ", Label: " + q(qq.Attr("aria-label")) + "}"
}

func (g *gen) filter(x *El) string {
	f := []string{"Query: " + queryOf(x.Kid("tc-q"))}
	if x.Has("is-ph") {
		f = append(f, "Waiting: true")
	}
	if s := svgs(x); len(s) > 0 && s[0].Has("tc-ico--dusty") {
		f = append(f, "Quiet: true")
	}
	return "w.Filter{" + join(f, ", ") + g.ev(x, "FilterVariant", "Variant") + placef(x) + "}"
}

func (g *gen) tagInput(x *El) string {
	var parts []string
	for _, k := range x.Els() {
		switch {
		case k.Has("tc-tag"):
			parts = append(parts, g.tag(k))
		case k.Has("tc-dim"):
			parts = append(parts, g.dimx(k))
		default:
			parts = append(parts, g.inlineEl(k))
		}
	}
	f := []string{"Parts: []w.TagInputPart{" + join(parts, ", ") + "}"}
	if x.Has("tc-in--wrap-gap5") {
		f = append(f, "Wrap: true")
	}
	if x.Has("is-focus") {
		f = append(f, "Focus: true")
	}
	return "w.TagInput{" + join(f, ", ") + "}"
}

// preText is a text area's text as written, without its leading and
// trailing line breaks.
func preText(x *El) string {
	if x.Tag != "textarea" {
		return ""
	}
	return strings.Trim(strs(x.Kids), "\n")
}

func (g *gen) area(x *El) string {
	t := preText(x)
	if t == "" {
		t = x.Text()
	}
	f := []string{"Text: " + q(t)}
	if v := g.en.of(x, "AreaVariant"); v != "" && v != "w.AreaPlain" {
		f = append(f, "Variant: "+v)
	}
	for _, c := range []kv{{"is-mono", "Mono"}, {"is-focus", "Focus"}} {
		if x.Has(c.k) {
			f = append(f, c.v+": true")
		}
	}
	for _, a := range []kv{{"name", "Name"}, {"aria-label", "Label"}} {
		if v := x.Attr(a.k); v != "" {
			f = append(f, a.v+": "+q(v))
		}
	}
	if x.Attr("aria-label") == "" {
		// A text field the frame leaves unnamed still needs its name (a11y).
		f = append(f, `Label: "Text"`)
	}
	return "w.Area{" + join(f, ", ") + "}"
}

func (g *gen) seg(x *El) string {
	var opts []string
	for _, o := range x.Els() {
		var f []string
		if t := o.OwnText(); t != "" {
			f = append(f, "Label: "+q(t))
		}
		if s := svgs(o); len(s) > 0 {
			f = append(f, "Icon: "+icon(s[0]))
		}
		if o.Has("is-on") {
			f = append(f, "On: true")
		}
		if v := o.Attr("aria-label"); v != "" {
			f = append(f, "Title: "+q(v))
		}
		opts = append(opts, "{"+join(f, ", ")+"}")
	}
	return "w.Seg{Options: []w.SegOption{" + join(opts, ", ") + "}" + g.ev(x, "SegVariant", "Variant") + placesec(x) + "}"
}

func (g *gen) checkPart(k *El) string {
	switch {
	case k.Has("tc-dim"):
		return g.dimx(k)
	case k.Has("tc-mono"):
		return g.mono(k)
	case k.Has("tc-badge"):
		return g.badge(k)
	case k.Has("tc-diff__value"):
		return "w.Choice{Source: " + q(k.Kid("tc-diff__source").Text()) + ", Value: " + q(k.OwnText()) + "}"
	}
	stop("check part? " + join(k.Cls(), " "))
	return ""
}

func (g *gen) check(x *El) string {
	tick := x.Kid("tc-tickrow__tick")
	var parts []string
	for _, k := range x.Kids {
		switch v := k.(type) {
		case string:
			if t := textpart(x, v); t != "" {
				parts = append(parts, t)
			}
		case *El:
			if v != tick {
				parts = append(parts, g.checkPart(v))
			}
		}
	}
	var f []string
	if tick != nil && !tick.Has("is-off") {
		f = append(f, "On: true")
	}
	f = append(f, "Parts: []w.CheckPart{"+join(parts, ", ")+"}")
	if v := x.Attr("aria-label"); v != "" {
		f = append(f, "Label: "+q(v))
	}
	if x.Has("is-3rd") {
		f = append(f, "Tertiary: true")
	}
	return "w.Check{" + join(f, ", ") + "}"
}

// control is any control or field element: the one dispatcher fields,
// lines and forms share; "" when x is none.
func (g *gen) control(x *El) string {
	switch {
	case x.Tag == "input" && x.Has("tc-in"):
		return g.input(x)
	case x.Tag == "button" && x.Has("tc-in") && x.Has("is-pick"):
		return g.pick(x)
	case x.Tag == "label" && x.Has("tc-in"):
		return g.filter(x)
	case x.Tag == "span" && x.Has("tc-in"):
		return g.tagInput(x)
	case x.Tag == "textarea":
		return g.area(x)
	case x.Has("tc-seg"):
		return g.seg(x)
	case x.Has("tc-check"):
		return g.check(x)
	case x.Has("tc-btn"):
		return g.button(x)
	}
	return ""
}

func (g *gen) hintw(x *El) string {
	return "w.FieldHint{Parts: []w.Inline{" + join(g.inlparts(x, g.inlineEl), ", ") + "}" + g.ev(x, "HintVariant", "Variant") + "}"
}

func (g *gen) group(x *El) string {
	form := ", Form: w.GroupColumnLoose"
	if x.Has("tc-line") {
		form = ""
	} else if x.Has("tc-stack--gap-2") {
		form = ", Form: w.GroupColumn"
	}
	return "w.Group{Items: []w.FieldPart{" + g.each(x.Els(), g.fieldPart) + "}" + form + "}"
}

// each joins what fn makes of every element.
func (g *gen) each(es []*El, fn func(*El) string) string {
	var out []string
	for _, e := range es {
		out = append(out, fn(e))
	}
	return join(out, ", ")
}

func (g *gen) fieldPart(x *El) string {
	if c := g.control(x); c != "" {
		return c
	}
	switch {
	case x.Has("tc-field__hint"):
		return g.hintw(x)
	case x.Has("tc-field__swatches"):
		var cols []string
		on := -1
		for i, y := range x.Els() {
			cols = append(cols, q(y.Attr("data-tc-colour")))
			if on < 0 && y.Has("tc-field__swatch--on") {
				on = i
			}
		}
		s := "w.Swatches{Colours: []string{" + join(cols, ", ") + "}"
		if on > 0 {
			s += ", On: " + itoa(on)
		}
		return s + "}"
	case x.Has("tc-field__zoom"):
		return "w.Zoom{Value: " + q(x.Kid("tc-mono").Text()) + "}"
	case x.Has("tc-field__row"):
		return "w.FieldRow{Items: []w.FieldPart{" + g.each(x.Els(), g.fieldPart) + "}}"
	case x.Has("tc-card__top2"):
		return g.group(x)
	case x.Has("tc-field__photo-row"):
		ph := x.Kid("tc-avatar--photo")
		side := x.Kid("tc-field__photo-side")
		ini := ""
		if t := ph.Text(); t != "" {
			ini = "Initials: " + q(t) + ", "
		}
		return "w.Photo{" + ini + "Side: []w.FieldPart{" + g.each(side.Els(), g.fieldPart) + "}}"
	case x.Has("tc-line"):
		return g.line(x)
	case x.Has("tc-mark"):
		return g.mark(x)
	case x.Has("tc-dim"):
		return g.dimx(x)
	case x.Has("tc-drop"):
		return g.drop(x)
	case x.Has("tc-menu"):
		return g.menu(x)
	case x.Has("tc-msg"):
		return g.msg(x)
	case x.Has("tc-diag"):
		return g.diag(x)
	}
	stop("field part? " + x.Tag + " " + join(x.Cls(), " "))
	return ""
}

var fieldGap = regexp.MustCompile(`^tc-field--gap(\d+)$`)

func (g *gen) field(k *El) string {
	var f []string
	var skip *El
	var lab *El
	for _, x := range k.Els() {
		if x.Tag == "label" && !x.Has("tc-in") {
			lab = x
			break
		}
	}
	if ls := k.Kid("tc-field__label"); ls != nil {
		f = append(f, "Label: "+q(ls.Text())+", LabelSpan: true")
		skip = ls
	} else if lab != nil {
		skip = lab
		raw := strs(lab.Kids)
		if strings.Contains(raw, nbsp) && strings.TrimSpace(strings.Map(func(r rune) rune {
			if isSpace(r) {
				return ' '
			}
			return r
		}, raw)) == "" {
			// An empty label that keeps its line is a no-break space in the frame.
			f = append(f, "Label: \""+nbsp+"\"")
		} else {
			t := lab.OwnText()
			if t == "" {
				t = " "
			}
			f = append(f, "Label: "+q(t))
		}
		for _, a := range lab.Els() {
			switch {
			case a.Has("tc-dim"):
				f = append(f, "LabelAside: "+g.dimx(a))
			case a.Has("tc-plain"):
				f = append(f, "LabelAside: w.Plain{Text: "+q(a.Text())+"}")
			default:
				stop("label aside? " + join(a.Cls(), " "))
			}
		}
	}
	var parts []string
	for _, x := range k.Els() {
		if x != skip {
			parts = append(parts, g.fieldPart(x))
		}
	}
	if len(parts) > 0 {
		f = append(f, "Parts: []w.FieldPart{"+join(parts, ", ")+"}")
	}
	if v := g.en.of(k, "FieldWidth"); v != "" {
		f = append(f, "Width: "+v)
	}
	for _, c := range k.Cls() {
		if m := fieldGap.FindStringSubmatch(c); m != nil {
			f = append(f, "Gap: "+m[1])
		}
	}
	if k.Has("is-own-line") {
		f = append(f, "OwnLine: true")
	}
	return "w.Field{" + join(f, ", ") + placef(k) + "}"
}

func (g *gen) linePart(x *El) string {
	if c := g.control(x); c != "" {
		return c
	}
	switch {
	case x.Has("tc-field"):
		return g.field(x)
	case x.Has("tc-kbd"):
		return "w.Kbd{Keys: " + q(x.Text()) + placesec(x) + "}"
	case x.Has("tc-dim"):
		return g.dimx(x)
	case x.Has("tc-mark"):
		return g.mark(x)
	case x.Has("tc-mono"):
		return g.mono(x)
	case x.Has("tc-tag"):
		return g.tag(x)
	case x.Has("tc-badge"):
		return g.badge(x)
	case x.Has("tc-meter"):
		return g.meter(x)
	case x.Has("tc-line"):
		return g.line(x)
	case x.Has("tc-push"):
		return g.push(x)
	case x.Tag == "b":
		return "w.Strong(" + q(x.Text()) + ")"
	}
	cl := join(x.Cls(), " ")
	if f, ok := g.faces[cl]; ok && x.Tag == "span" {
		return "w.Faced{Face: w." + f + ", Items: []w.LinePart{" + join(g.lineText(x), ", ") + "}}"
	}
	stop("line part? " + x.Tag + " " + cl)
	return ""
}

func (g *gen) lineText(e *El) []string {
	var out []string
	for _, k := range e.Kids {
		switch v := k.(type) {
		case string:
			if t := textpart(e, v); t != "" {
				out = append(out, t)
			}
		case *El:
			out = append(out, g.linePart(v))
		}
	}
	return out
}

func (g *gen) line(e *El) string {
	for _, c := range e.Cls() {
		if c != "tc-line" && !strings.HasPrefix(c, "tc-line--gap-") && c != "tc-push" && c != "tc-grow" {
			stop("line with " + c)
		}
	}
	f := []string{"Items: []w.LinePart{" + join(g.lineText(e), ", ") + "}"}
	if v := g.en.of(e, "LineGap"); v != "" && v != "w.LineGap0" {
		f = append(f, "Gap: "+v)
	}
	if e.Tag == "div" {
		f = append(f, "Block: true")
	}
	return "w.Line{" + join(f, ", ") + placef(e) + "}"
}

func (g *gen) fsetPart(k *El) string {
	switch {
	case k.Has("tc-fset__grid"):
		return "w.FieldGrid{Items: []w.FieldGridPart{" + g.each(k.Els(), g.tickrow) + "}}"
	case k.Has("tc-field"):
		return g.field(k)
	case k.Has("tc-btn"):
		return g.button(k)
	case k.Has("tc-line"):
		return g.line(k)
	case k.Has("tc-grow") && len(k.Els()) == 0 && k.Text() == "":
		return "w.Grow{}"
	}
	return g.anyblock(k)
}

func (g *gen) fset(e *El) string {
	kids := e.Els()
	var r *El
	if len(kids) > 0 && kids[0].Has("tc-rubric") {
		r = kids[0]
	}
	var items []string
	for _, k := range kids {
		if k != r {
			items = append(items, g.fsetPart(k))
		}
	}
	var f []string
	if r != nil {
		f = append(f, "Rubric: "+g.rubric(r, ""))
	}
	f = append(f, "Items: []w.FieldSetPart{"+join(items, ", ")+"}")
	if v := g.en.of(e, "FieldSetVariant"); v != "" {
		f = append(f, "Variant: "+v)
	}
	return "w.FieldSet{" + join(f, ", ") + "}"
}

func (g *gen) formPart(k *El) string {
	switch {
	case k.Has("tc-fset"):
		return g.fset(k)
	case k.Has("tc-field"):
		return g.field(k)
	case k.Has("tc-check"):
		return g.check(k)
	case k.Has("tc-rubric"):
		return g.rubric(k, "")
	case k.Has("tc-card"):
		return g.card(k)
	case k.Has("tc-msg"):
		return g.msg(k)
	case k.Has("tc-form__cols"):
		return "w.FormCols{Items: []w.Field{" + g.each(k.Els(), g.field) + "}}"
	case k.Has("tc-form__col--460") || k.Has("tc-form__col--300"):
		its := g.each(k.Els(), func(y *El) string {
			if y.Has("tc-card") {
				return g.card(y)
			}
			return g.anyblock(y)
		})
		wide := ""
		if k.Has("tc-form__col--460") {
			wide = "Wide: true, "
		}
		return "w.FormColumn{" + wide + "Items: []w.FormColumnPart{" + its + "}}"
	}
	stop("form part? " + join(k.Cls(), " "))
	return ""
}

func (g *gen) form(e *El) string {
	f := []string{"Items: []w.FormPart{" + g.each(e.Els(), g.formPart) + "}"}
	if v := g.en.of(e, "FormVariant"); v != "" {
		f = append(f, "Variant: "+v)
	}
	if e.Has("is-tight") {
		f = append(f, "Tight: true")
	}
	return "w.Form{" + join(f, ", ") + "}"
}

func (g *gen) formFootPart(k *El) string {
	switch {
	case k.Has("tc-btn"):
		return g.button(k)
	case k.Has("tc-dim"):
		return g.dimx(k)
	case k.Has("tc-push"):
		return "w.Push{Items: []w.PushPart{" + g.each(k.Els(), g.pushItem) + "}}"
	}
	cl := join(k.Cls(), " ")
	if f, ok := g.faces[cl]; ok && k.Tag == "span" {
		return "w.Span{Face: w." + f + ", Text: " + q(k.Text()) + "}"
	}
	stop("form foot part? " + cl)
	return ""
}

func (g *gen) formFoot(e *El) string {
	return "w.FormFoot{Items: []w.FormFootPart{" + g.each(e.Els(), g.formFootPart) + "}" + g.ev(e, "FormFootVariant", "Variant") + "}"
}

func (g *gen) barItem(k *El) string {
	if k.Has("tc-bar__sep") {
		var f []string
		if k.Has("is-2nd") {
			f = append(f, "Secondary: true")
		}
		if k.Has("is-3rd") {
			f = append(f, "Tertiary: true")
		}
		return "w.BarSep{" + join(f, ", ") + "}"
	}
	if k.Has("tc-range") {
		var ins []*El
		for _, i := range k.Els() {
			if i.Tag == "input" {
				ins = append(ins, i)
			}
		}
		s := "w.Range{From: " + g.input(ins[0]) + ", To: " + g.input(ins[1]) + ", Show: " + g.button(k.Kid("tc-btn"))
		if k.Has("is-2nd") {
			s += ", Secondary: true"
		}
		return s + "}"
	}
	if c := g.control(k); c != "" {
		return c
	}
	switch {
	case k.Has("tc-mono"):
		return g.mono(k)
	case k.Has("tc-drop"):
		return g.drop(k)
	}
	stop("bar? " + join(k.Cls(), " "))
	return ""
}

func (g *gen) drop(e *El) string {
	els := e.Els()
	t, m := els[0], els[1]
	tr := g.control(t)
	if tr == "" {
		tr = g.button(t)
	}
	f := []string{"Trigger: " + tr}
	if t.Has("tc-bar__more") {
		f = append(f, "More: true")
	}
	f = append(f, "Menu: "+g.menu(m))
	return "w.Drop{Open: true, " + join(f, ", ") + "}"
}

func (g *gen) pushItem(y *El) string {
	if y.Has("tc-btn") {
		return g.button(y)
	}
	return g.child(y)
}

// push is a pushed group: its buttons and words, in a face of its own when
// its classes name one.
func (g *gen) push(x *El) string {
	var items []string
	for _, k := range x.Kids {
		switch v := k.(type) {
		case string:
			if t := textpart(x, v); t != "" {
				items = append(items, t)
			}
		case *El:
			items = append(items, g.pushItem(v))
		}
	}
	var fc []string
	for _, c := range x.Cls() {
		if c != "tc-push" {
			fc = append(fc, c)
		}
	}
	f := ""
	if face, ok := g.faces[join(fc, " ")]; ok {
		f = ", Face: faceOf(w." + face + ")"
	}
	return "w.Push{Items: []w.PushPart{" + join(items, ", ") + "}" + f + "}"
}

func (g *gen) rubricItem(x *El) string {
	switch {
	case x.Has("tc-line"):
		return g.line(x)
	case x.Has("tc-mono"):
		return g.mono(x)
	case x.Has("tc-badge"):
		return g.badge(x)
	case x.Has("tc-num"):
		return g.num(x)
	case x.Has("tc-push"):
		s := "w.Push{Items: []w.PushPart{" + g.each(x.Els(), g.pushItem)
		if t := x.OwnText(); t != "" && len(x.Els()) == 0 {
			s += "w.Text(" + q(t) + ")"
		}
		return s + "}}"
	case x.Has("tc-rubric__aside"):
		its := g.each(x.Els(), func(y *El) string {
			if y.Has("tc-badge") {
				return g.badge(y)
			}
			return g.mono(y)
		})
		s := "w.RubricAside{Text: " + q(x.OwnText())
		if its != "" {
			s += ", Items: []w.RubricAsidePart{" + its + "}"
		}
		return s + "}"
	case x.Has("tc-rubric__col"):
		return "w.RubricCol{Text: " + q(x.Text()) + "}"
	case x.Has("tc-btn"):
		return g.button(x)
	case x.Has("tc-plain"):
		return "w.Plain{Text: " + q(x.Text()) + "}"
	}
	stop("rubric item? " + join(x.Cls(), " "))
	return ""
}

func (g *gen) num(n *El) string {
	s := "w.Num{Value: " + n.Text()
	if n.Has("is-alert") {
		s += ", Alert: true"
	}
	return s + "}"
}

// rubric is a rubric; form "" takes the form from the tag.
func (g *gen) rubric(e *El, form string) string {
	var f []string
	var kids []any
	anyText := false
	for _, k := range e.Kids {
		if s, ok := k.(string); ok {
			if strings.TrimSpace(strings.Map(spaceToBlank, s)) == "" {
				continue
			}
			anyText = true
		}
		kids = append(kids, k)
	}
	var lead *El
	if len(kids) > 0 {
		if k0, ok := kids[0].(*El); ok && k0.Has("tc-mono") && anyText {
			lead = k0
		}
	}
	f = append(f, "Text: "+q(e.OwnText()))
	if lead != nil {
		f = append(f, "Lead: &"+g.mono(lead))
	}
	var items []string
	for _, x := range e.Els() {
		if x != lead {
			items = append(items, g.rubricItem(x))
		}
	}
	if len(items) > 0 {
		f = append(f, "Items: []w.RubricPart{"+join(items, ", ")+"}")
	}
	if form == "" {
		form = map[string]string{"span": "RubricInline", "button": "RubricItem"}[e.Tag]
	}
	if form != "" {
		f = append(f, "Form: w."+form)
	}
	if v := g.en.of(e, "RubricVariant"); v != "" {
		f = append(f, "Variant: "+v)
	}
	return "w.Rubric{" + join(f, ", ") + "}"
}

func spaceToBlank(r rune) rune {
	if isSpace(r) {
		return ' '
	}
	return r
}

func (g *gen) listGroup(k *El) string {
	f := []string{"Title: " + g.rubric(k.Kid("tc-rubric"), "")}
	if n := k.Kid("tc-num"); n != nil {
		f = append(f, "Count: &"+g.num(n))
	}
	return "w.ListGroup{" + join(f, ", ") + "}"
}
