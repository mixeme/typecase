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
		"search": "Search", "link": "Link", "up": "Upload", "refresh": "Refresh", "move": "Move",
		"download": "Download", "upload": "Upload", "import": "Import", "export": "Export"}
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
		href := ""
		if h := e.Attr("href"); h != "" {
			href = ", Href: " + q(h)
		}
		return "w.Close{Label: " + q(lab) + placef(e) + href + "}"
	}
	s := g.button0(e)
	var extra []string
	if v := g.en.of(e, "ButtonVariant"); v != "" && !strings.Contains(s, "Variant:") {
		extra = append(extra, "Variant: "+v)
	}
	if e.Has("is-off") {
		extra = append(extra, "State: w.ButtonOff")
	}
	if e.Has("is-wait") {
		// waits for its form: its title is why (Carrel's task 198)
		extra = append(extra, "State: w.ButtonWaits")
		s = strings.Replace(s, "Title: ", "Reason: ", 1)
	}
	if e.Valued("disabled") || e.Attr("aria-disabled") == "true" {
		extra = append(extra, "State: w.ButtonDisabled")
	}
	if e.Has("is-3rd") {
		extra = append(extra, "Tertiary: true")
	}
	for _, a := range []kv{{"href", "Href"}, {"form", "Form"}, {"formaction", "To"}, {"name", "Name"}, {"value", "Value"}, {"data-tc-add", "Adds"}, {"data-tc-save", "SaveAs"}, {"data-tc-save-from", "SaveFrom"}, {"data-tc-copy", "Copy"}, {"data-tc-copy-from", "CopyFrom"}, {"data-tc-copied", "Copied"}, {"data-tc-count", "Count"}, {"data-tc-count-1", "CountOne"}, {"data-tc-swap", "Swap"}, {"data-tc-busy", "Busy"}, {"data-tc-target", "Target"}} {
		if v := e.Attr(a.k); v != "" {
			extra = append(extra, a.v+": "+q(v))
		}
	}
	if e.Tag == "a" && e.Attr("aria-disabled") == "true" {
		extra = append(extra, `Href: "#"`)
	}
	if e.Tag == "button" && e.Attr("type") == "submit" && e.Attr("form") == "" {
		extra = append(extra, "Submit: true")
	}
	if e.HasAttr("data-tc-remove") {
		extra = append(extra, "Removes: true")
	}
	extra = append(extra, editBarAttrs(e)...)
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
// fileField is the file input inside a label that picks a file (Carrel's
// task 109).
func (g *gen) fileField(f *El) string {
	s := "&w.FileField{Form: " + q(f.Attr("form")) + ", Name: " + q(f.Attr("name"))
	if a := f.Attr("accept"); a != "" {
		s += ", Accept: " + q(a)
	}
	s += ", Label: " + q(f.Attr("aria-label"))
	for _, a := range []kv{{"data-tc-to", "To"}, {"data-tc-list", "List"}, {"data-tc-preview", "Preview"}, {"data-tc-shows", "Shows"}, {"data-tc-size", "Size"}, {"data-tc-ranges", "Ranges"}, {"data-tc-taken", "Taken"}} {
		if v := f.Attr(a.k); v != "" {
			s += ", " + a.v + ": " + q(v)
		}
	}
	if f.HasAttr("webkitdirectory") {
		s += ", Folder: true"
	} else if f.HasAttr("multiple") {
		s += ", Many: true"
	}
	return s + "}"
}

// pickedRow is the row a file picked by a menu item becomes.
func (g *gen) pickedRow(tpl *El) string {
	r := tpl.Els()[0]
	if r.Has("tc-line__text") {
		return g.docAttach(r)
	}
	return g.line(r)
}

// rowTemplate is the template of a picked file's row beside a file input.
func rowTemplate(e *El) *El {
	for _, k := range e.Els() {
		if k.Tag == "template" && k.HasAttr("data-tc-row") {
			return k
		}
	}
	return nil
}

func (g *gen) button(e *El) string {
	if f := e.Kid("tc-file"); f != nil {
		// A label that picks a file: the button without its input, and the input as File.
		c := *e
		c.Tag = "button"
		var kids []any
		row := rowTemplate(e)
		for _, k := range e.Kids {
			if el, ok := k.(*El); ok && (el == f || el == row) {
				continue
			}
			kids = append(kids, k)
		}
		c.Kids = kids
		s := g.button(&c)
		return s[:len(s)-1] + ", File: " + g.fileField(f) + "}"
	}
	s := g.button2(e)
	if e.Has("tc-nojs") {
		s = s[:len(s)-1] + ", NoScript: true}"
	}
	if e.Has("tc-js") {
		s = s[:len(s)-1] + ", ScriptOnly: true}"
	}
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
		if b.Has("tc-drop") {
			out = append(out, g.drop(b))
			continue
		}
		out = append(out, g.button(b))
	}
	return join(out, ", ")
}

var inputKind = map[string]string{"date": "InputDate", "time": "InputTime", "search": "InputSearch", "password": "InputPassword"}

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
	if x.HasAttr("disabled") {
		f = append(f, "Disabled: true")
	}
	if x.HasAttr("data-tc-filename") {
		// the line of a picked file's name: read-only by itself (Carrel's task 133)
		f = append(f, "Picked: true")
	} else if x.HasAttr("readonly") {
		f = append(f, "ReadOnly: true")
	}
	if v := x.Attr("id"); v != "" {
		// what a button beside it names: a file field's line, a Copy
		f = append(f, "ID: "+q(v))
	}
	if x.HasAttr("required") {
		f = append(f, "Required: true")
	}
	if v := x.Attr("data-tc-follows"); v != "" {
		f = append(f, "Follows: "+q(v))
	}
	if v := x.Attr("data-tc-day"); v != "" {
		f = append(f, "Day: "+q(v))
	}
	if v := x.Attr("data-tc-from"); v != "" {
		f = append(f, "From: "+q(v))
	}
	if v := x.Attr("form"); v != "" {
		// a field of a form it does not stand in (Carrel's task 195)
		f = append(f, "Of: "+q(v))
	}
	return "w.Input{" + join(f, ", ") + g.ev(x, "InputVariant", "Variant") + placef(x) + "}"
}

func (g *gen) pick(x *El) string {
	if list := x.Kid("tc-in__list"); list != nil {
		var opts []string
		for _, o := range list.Els() {
			o2 := []string{"Value: " + q(o.Attr("value")), "Label: " + q(o.Text())}
			if o.HasAttr("selected") {
				o2 = append(o2, "On: true")
			}
			if o.HasAttr("disabled") {
				o2 = append(o2, "Unavailable: true")
			}
			if v := o.Attr("data-tc-shows"); v != "" {
				o2 = append(o2, "Shows: "+q(v))
			}
			if v := o.Attr("data-tc-disables"); v != "" {
				o2 = append(o2, "Disables: "+q(v))
			}
			opts = append(opts, "{"+join(o2, ", ")+"}")
		}
		f := []string{"Name: " + q(list.Attr("name")), "Options: []w.PickOption{" + join(opts, ", ") + "}"}
		if v := list.Attr("aria-label"); v != "" {
			f = append(f, "Label: "+q(v))
		}
		if m := x.Kid("tc-mark"); m != nil {
			f = append(f, "Mark: &"+g.mark(m))
		}
		if x.Has("is-ph") {
			f = append(f, "Unset: true")
		}
		if x.Has("is-focus") {
			f = append(f, "Focus: true")
		}
		return "w.Pick{" + join(f, ", ") + g.ev(x, "PickVariant", "Variant") + placef(x) + "}"
	}
	f := []string{"Value: " + q(x.OwnText())}
	if m := x.Kid("tc-mark"); m != nil {
		f = append(f, "Mark: &"+g.mark(m))
	}
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

func queryOf(qq, box *El) string {
	n := ""
	if v := qq.Attr("name"); v != "" {
		n = ", Name: " + q(v)
	}
	if v := qq.Attr("form"); v != "" {
		n += ", Of: " + q(v)
	}
	v := "Value: " + q(qq.Attr("value"))
	if ph := qq.Attr("placeholder"); ph != "" {
		v = "Placeholder: " + q(ph)
		// the short word stands in the hint over the line: tc-q__ph >
		// tc-q__hint > tc-q__short
		if sh := box.Kid("tc-q__ph"); sh != nil {
			if h := sh.Kid("tc-q__hint"); h != nil {
				if t := h.Kid("tc-q__short"); t != nil {
					v += ", Short: " + q(t.Text())
				}
			}
		}
	}
	return "w.Query{" + v + ", Label: " + q(qq.Attr("aria-label")) + n + "}"
}

func (g *gen) filter(x *El) string {
	f := []string{"Query: " + queryOf(x.Kid("tc-q"), x)}
	if x.Has("is-ph") {
		f = append(f, "Waiting: true")
	}
	if s := svgs(x); len(s) > 0 && s[0].Has("tc-ico--dusty") {
		f = append(f, "Quiet: true")
	}
	if c := x.Kid("tc-in__clear"); c != nil {
		f = append(f, "Clear: "+q(c.Attr("href")))
	}
	if x.Tag == "form" {
		f = append(f, "Action: "+q(x.Attr("action")))
		var hs []string
		for _, h := range x.Els() {
			if h.Tag == "input" && h.Attr("type") == "hidden" {
				hs = append(hs, "{Name: "+q(h.Attr("name"))+", Value: "+q(h.Attr("value"))+"}")
			}
		}
		if len(hs) > 0 {
			f = append(f, "Hidden: []w.Hidden{"+join(hs, ", ")+"}")
		}
		if sw := x.Attr("data-tc-swap"); sw != "" {
			f = append(f, "Swap: "+q(sw))
		}
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
		case k.Has("tc-in__add"):
			continue
		case k.Tag == "button" && k.HasAttr("hidden"):
			continue // the line's Enter (TagAdd.Enter)
		default:
			parts = append(parts, g.inlineEl(k))
		}
	}
	f := []string{"Parts: []w.TagInputPart{" + join(parts, ", ") + "}"}
	if x.Has("tc-in--wrap-gap") {
		f = append(f, "Wrap: true")
	}
	if x.Has("tc-in--cols") {
		f = append(f, "Columns: true")
	}
	name := ""
	for _, k := range x.Els() {
		if k.Has("tc-tag") && name == "" {
			for _, i := range k.Els() {
				if i.Tag == "input" && i.Attr("name") != "" {
					name = i.Attr("name")
				}
			}
		}
	}
	if name != "" {
		f = append(f, "Name: "+q(name))
	}
	if a := x.Kid("tc-in__add"); a != nil {
		af := []string{}
		for _, kv := range [][2]string{{"Name", "name"}, {"Value", "value"}, {"Placeholder", "placeholder"}, {"Label", "aria-label"}} {
			if v := a.Attr(kv[1]); v != "" {
				af = append(af, kv[0]+": "+q(v))
			}
		}
		if v := x.Attr("data-tc-suggest"); v != "" {
			af = append(af, "Suggest: "+q(v))
		}
		for _, k := range x.Els() {
			if k.Tag == "button" && k.HasAttr("hidden") {
				af = append(af, "Enter: &w.TagEnter{Name: "+q(k.Attr("name"))+", Value: "+q(k.Attr("value"))+", Label: "+q(k.Text())+", Of: "+q(k.Attr("form"))+"}")
			}
		}
		f = append(f, "Add: &w.TagAdd{"+join(af, ", ")+"}")
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

// segfold is a switch of links that folds on a narrow bar: its options are
// read from the fold's menu (label, count, address) and the current one.
// checkForm is a bar's check that is a form: its check, address and hidden
// fields.
func (g *gen) checkForm(x *El) string {
	var cs []string
	for _, k := range x.Els() {
		if k.Has("tc-check") {
			cs = append(cs, g.check(k))
		}
		if k.Has("tc-menu__item") {
			// The twin in a menu: the item's checkbox is the check.
			b := k.Kid("tc-check__box")
			c := "w.Check{"
			if b.HasAttr("checked") {
				c += "On: true, "
			}
			cs = append(cs, c+"Name: "+q(b.Attr("name"))+", Value: "+q(b.Attr("value"))+", Parts: []w.CheckPart{w.Text("+q(k.OwnText())+")}}")
		}
	}
	s := "w.CheckForm{Checks: []w.Check{" + join(cs, ", ") + "}, Action: " + q(x.Attr("action"))
	if x.Has("is-3rd") {
		s += ", Tertiary: true"
	}
	var hs []string
	for _, k := range x.Els() {
		if k.Tag == "input" && k.Attr("type") == "hidden" {
			hs = append(hs, "{Name: "+q(k.Attr("name"))+", Value: "+q(k.Attr("value"))+"}")
		}
	}
	if len(hs) > 0 {
		s += ", Hidden: []w.Hidden{" + join(hs, ", ") + "}"
	}
	if ap := x.Kid("tc-checkform__apply"); ap != nil {
		s += ", Apply: " + g.button(ap.Kid("tc-btn"))
	}
	return s + "}"
}

func (g *gen) segfold(x *El) string {
	var opts []string
	for _, it := range x.Kid("tc-segfold__pick").Kid("tc-segfold__menu").Els() {
		f := []string{"Label: " + q(it.OwnText())}
		if c := it.Kid("tc-segfold__count"); c != nil {
			f = append(f, "Count: "+q(c.Text()))
		}
		f = append(f, "Href: "+q(it.Attr("href")))
		if it.Has("is-on") {
			f = append(f, "On: true")
		}
		opts = append(opts, "{"+join(f, ", ")+"}")
	}
	ph := ""
	if x.Has("tc-segfold--phone") {
		ph = ", Phone: true"
	}
	return "w.SegFold{Seg: w.Seg{Options: []w.SegOption{" + join(opts, ", ") + "}}" + ph + "}"
}

func (g *gen) seg(x *El) string {
	var opts []string
	name := ""
	for _, o := range x.Els() {
		var f []string
		box := o.Kid("tc-seg__box")
		if t := o.OwnText(); t != "" {
			f = append(f, "Label: "+q(t))
		}
		if s := svgs(o); len(s) > 0 {
			f = append(f, "Icon: "+icon(s[0]))
		}
		if o.Tag == "a" {
			f = append(f, "Href: "+q(o.Attr("href")))
		}
		if o.Tag == "button" && o.Attr("type") == "submit" {
			f = append(f, "Form: "+q(o.Attr("form")))
			if v := o.Attr("formaction"); v != "" {
				f = append(f, "To: "+q(v))
			}
		}
		if o.Has("is-on") || (box != nil && box.HasAttr("checked")) {
			f = append(f, "On: true")
		}
		if v := o.Attr("aria-label"); v != "" {
			f = append(f, "Title: "+q(v))
		}
		if box != nil {
			name = box.Attr("name")
			f = append(f, "Value: "+q(box.Attr("value")))
			if v := box.Attr("aria-label"); v != "" {
				f = append(f, "Title: "+q(v))
			}
			if v := box.Attr("data-tc-disables"); v != "" {
				f = append(f, "Disables: "+q(v))
			}
		}
		opts = append(opts, "{"+join(f, ", ")+"}")
	}
	n := ""
	if name != "" {
		n = ", Name: " + q(name)
	}
	for _, o := range x.Els() {
		if b := o.Kid("tc-seg__box"); b != nil && b.Attr("type") == "checkbox" {
			n += ", Multiple: true"
			break
		}
	}
	if x.HasAttr("data-tc-theme-switch") {
		n += ", Theme: true"
	}
	if a := x.Kid("tc-seg__opt"); a != nil && a.Attr("data-tc-swap") != "" {
		n += ", Swap: " + q(a.Attr("data-tc-swap"))
	}
	if x.Has("is-3rd") {
		n += ", Tertiary: true"
	}
	return "w.Seg{Options: []w.SegOption{" + join(opts, ", ") + "}" + g.ev(x, "SegVariant", "Variant") + placesec(x) + n + "}"
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
	box := x.Kid("tc-check__box")
	var parts []string
	for _, k := range x.Kids {
		switch v := k.(type) {
		case string:
			if t := textpart(x, v); t != "" {
				parts = append(parts, t)
			}
		case *El:
			if v != tick && v != box {
				parts = append(parts, g.checkPart(v))
			}
		}
	}
	var f []string
	if box != nil {
		if box.HasAttr("checked") {
			f = append(f, "On: true")
		}
		f = append(f, "Name: "+q(box.Attr("name")), "Value: "+q(box.Attr("value")))
		if box.Attr("type") == "radio" {
			f = append(f, "Radio: true")
		}
		if v := box.Attr("aria-label"); v != "" {
			f = append(f, "Label: "+q(v))
		}
		if v := box.Attr("data-tc-disables"); v != "" {
			f = append(f, "Disables: "+q(v))
		}
		if v := box.Attr("form"); v != "" {
			f = append(f, "Of: "+q(v))
		}
		if box.HasAttr("required") {
			f = append(f, "Required: true")
		}
		if box.HasAttr("disabled") {
			f = append(f, "Reason: "+q(x.Attr("title")))
		}
	} else if tick != nil && !tick.Has("is-off") {
		f = append(f, "On: true")
	}
	f = append(f, "Parts: []w.CheckPart{"+join(parts, ", ")+"}")
	if v := x.Attr("aria-label"); v != "" {
		f = append(f, "Label: "+q(v))
	}
	if x.Has("is-3rd") {
		f = append(f, "Tertiary: true")
	}
	if v := x.Attr("data-tc-autosave"); v != "" {
		f = append(f, "AutoSave: "+q(v))
	}
	return "w.Check{" + join(f, ", ") + "}"
}

// control is any control or field element: the one dispatcher fields,
// lines and forms share; "" when x is none.
func (g *gen) control(x *El) string {
	switch {
	case x.Tag == "input" && x.Has("tc-in"):
		return g.input(x)
	case x.Has("tc-in") && x.Has("is-pick"):
		return g.pick(x)
	case (x.Tag == "label" || x.Tag == "form") && x.Has("tc-in"):
		return g.filter(x)
	case x.Tag == "span" && x.Has("tc-in"):
		return g.tagInput(x)
	case x.Tag == "textarea":
		return g.area(x)
	case x.Has("tc-segfold"):
		return g.segfold(x)
	case x.Has("tc-checkform"):
		return g.checkForm(x)
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
	form := ", Form: w.GroupColumn"
	if x.Has("tc-line") {
		form = ""
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
		on, name := -1, ""
		for i, y := range x.Els() {
			cols = append(cols, q(y.Attr("data-tc-colour")))
			box := y.Kid("tc-seg__box")
			if on < 0 && (y.Has("tc-field__swatch--on") || box != nil && box.HasAttr("checked")) {
				on = i
			}
			if box != nil {
				name = box.Attr("name")
			}
		}
		s := "w.Swatches{Colours: []string{" + join(cols, ", ") + "}"
		if on > 0 {
			s += ", On: " + itoa(on)
		}
		if name != "" {
			s += ", Name: " + q(name)
		}
		return s + "}"
	case x.Has("tc-field__zoom"):
		return "w.Zoom{Value: " + q(x.Kid("tc-mono").Text()) + "}"
	case x.Has("tc-card__top2"):
		return g.group(x)
	case x.Has("tc-field__photo-row"):
		ph := x.Kid("tc-avatar--photo")
		side := x.Kid("tc-field__photo-side")
		ini := ""
		if t := ph.Text(); t != "" {
			ini = "Initials: " + q(t) + ", "
		}
		if img := ph.Kid("tc-avatar__img"); img != nil {
			ini = "Src: " + q(img.Attr("src")) + ", "
		}
		if id := ph.Attr("id"); id != "" {
			ini += "ID: " + q(id) + ", "
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

func (g *gen) field(k *El) string {
	var f []string
	if k.HasAttr("hidden") {
		f = append(f, "Hidden: true")
	}
	if c := k.Attr("data-tc-changed"); c != "" {
		var ns []string
		for _, n := range strings.Fields(c) {
			ns = append(ns, q(n))
		}
		f = append(f, "Changed: []string{"+join(ns, ", ")+"}")
	}
	var skip *El
	var lab *El
	for _, x := range k.Els() {
		if x.Tag == "label" && !x.Has("tc-in") && !x.Has("tc-check") {
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
	var parts, sends []string
	for _, x := range k.Els() {
		switch {
		case x == skip:
		case x.Tag == "input" && x.Attr("type") == "hidden":
			sends = append(sends, "w.Hidden{Name: "+q(x.Attr("name"))+", Value: "+q(x.Attr("value"))+"}")
		default:
			parts = append(parts, g.fieldPart(x))
		}
	}
	if len(parts) > 0 {
		f = append(f, "Parts: []w.FieldPart{"+join(parts, ", ")+"}")
	}
	if v := g.en.of(k, "FieldWidth"); v != "" {
		f = append(f, "Width: "+v)
	}
	if k.Has("tc-field--gap") {
		f = append(f, "Gapped: true")
	}
	if k.Has("is-own-line") {
		f = append(f, "OwnLine: true")
	}
	if k.Tag == "form" {
		f = append(f, "Action: "+q(k.Attr("action")))
		if strings.EqualFold(k.Attr("method"), "post") {
			f = append(f, "Post: true")
		}
		if len(sends) > 0 {
			f = append(f, "Sends: []w.Hidden{"+join(sends, ", ")+"}")
		}
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
		return g.kbd(x)
	case x.Has("tc-dim"):
		return g.dimx(x)
	case x.Has("tc-mark"):
		return g.mark(x)
	case x.Has("tc-say"):
		return g.say(x)
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
	case x.Has("tc-drop"):
		return g.drop(x)
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

// lineStep are the classes of the line's gap steps.
var lineStep = map[string]bool{"tc-line--tight": true, "tc-line--normal": true, "tc-line--wide": true, "tc-line--fields": true}

func (g *gen) line(e *El) string {
	for _, c := range e.Cls() {
		if c != "tc-line" && !lineStep[c] && c != "tc-push" && c != "tc-grow" {
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
	if id := e.Attr("id"); id != "" {
		f = append(f, "ID: "+q(id))
	}
	return "w.Line{" + join(f, ", ") + placef(e) + "}"
}

func (g *gen) fsetPart(k *El) string {
	switch {
	case k.Has("tc-fset__grid"):
		return "w.FieldGrid{Items: []w.FieldGridPart{" + g.each(k.Els(), func(x *El) string {
			if x.Has("tc-fset__col") {
				// a source and its choices, one column (Carrel's task 193)
				return "w.FieldGridColumn{Items: []w.TickRow{" + g.each(x.Els(), g.tickrow) + "}}"
			}
			return g.tickrow(x)
		}) + "}}"
	case k.Has("tc-field"):
		return g.field(k)
	case k.Has("tc-btn"):
		return g.button(k)
	case k.Has("tc-line"):
		return g.line(k)
	case k.Has("tc-grow") && len(k.Els()) == 0 && k.Text() == "":
		return "w.Grow{}"
	case k.Tag == "template":
		return "w.Template{ID: " + q(k.Attr("id")) + ", Items: []w.FieldSetPart{" + g.each(k.Els(), g.fsetPart) + "}}"
	}
	return g.anyblock(k)
}

func (g *gen) fset(e *El) string {
	var kids, hidden []*El
	for _, k := range e.Els() {
		if k.Tag == "input" && k.Attr("type") == "hidden" {
			// the hidden fields of a group that is a form (Carrel's task 195)
			hidden = append(hidden, k)
			continue
		}
		kids = append(kids, k)
	}
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
	if a := e.Attr("action"); a != "" {
		// a group that is a form of its own (Carrel's task 149)
		f = append(f, "Action: "+q(a))
		if m := e.Attr("method"); m != "" && m != "post" {
			f = append(f, "Method: "+q(m))
		}
		if len(hidden) > 0 {
			f = append(f, "Fields: "+hiddens(hidden))
		}
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
	var items, hidden []*El
	for _, k := range e.Els() {
		if k.Tag == "input" && k.Attr("type") == "hidden" {
			hidden = append(hidden, k)
		} else {
			items = append(items, k)
		}
	}
	f := []string{"Items: []w.FormPart{" + g.each(items, g.formPart) + "}"}
	if a := e.Attr("action"); a != "" {
		if id := e.Attr("id"); id != "" {
			f = append(f, "ID: "+q(id))
		}
		f = append(f, "Action: "+q(a))
		if m := e.Attr("method"); m != "" && m != "post" {
			f = append(f, "Method: "+q(m))
		}
		if len(hidden) > 0 {
			f = append(f, "Fields: "+hiddens(hidden))
		}
		if v := e.Attr("data-tc-recount"); v != "" {
			f = append(f, "Recount: "+q(v))
		}
	}
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
		return g.push(k)
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
		var hs []string
		for _, i := range k.Els() {
			if i.Tag == "input" && i.Attr("type") == "hidden" {
				hs = append(hs, "{Name: "+q(i.Attr("name"))+", Value: "+q(i.Attr("value"))+"}")
			} else if i.Tag == "input" {
				ins = append(ins, i)
			}
		}
		s := "w.Range{From: " + g.input(ins[0]) + ", To: " + g.input(ins[1]) + ", Show: " + g.button(k.Kid("tc-btn"))
		if len(hs) > 0 {
			s += ", Hidden: []w.Hidden{" + join(hs, ", ") + "}"
		}
		if k.Has("is-2nd") {
			s += ", Secondary: true"
		}
		if k.Tag == "form" {
			s += ", Action: " + q(k.Attr("action"))
		}
		return s + "}"
	}
	if k.Has("tc-barform") {
		// a bar's filter that is a form (Carrel's task 195): its fields stand
		// in the bar, Show sends it without a script
		var its, hs []string
		show := ""
		for _, i := range k.Els() {
			switch {
			case i.Tag == "input" && i.Attr("type") == "hidden":
				hs = append(hs, "{Name: "+q(i.Attr("name"))+", Value: "+q(i.Attr("value"))+"}")
			case i.Has("tc-barform__show"):
				show = g.button(i.Kid("tc-btn"))
			default:
				its = append(its, g.barItem(i))
			}
		}
		s := "w.BarForm{Items: []w.BarFormPart{" + join(its, ", ") + "}, Action: " + q(k.Attr("action"))
		if len(hs) > 0 {
			s += ", Hidden: []w.Hidden{" + join(hs, ", ") + "}"
		}
		return s + ", Show: " + show + "}"
	}
	if c := g.control(k); c != "" {
		return c
	}
	switch {
	case k.Has("tc-mono"):
		return g.mono(k)
	case k.Has("tc-dim"):
		return g.dimx(k)
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
	if t.Has("tc-acts__more") {
		f = append(f, "Row: true")
	}
	f = append(f, "Menu: "+g.menu(m))
	if e.Has("tc-push") {
		// pushed to the end of a sheet's head (Carrel's task 205)
		f = append(f, "Place: w.Pushed")
	}
	if m.HasAttr("hidden") || (e.Tag == "details" && !e.HasAttr("open")) {
		return "w.Drop{" + join(f, ", ") + "}"
	}
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
		return g.push(x)
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
	case x.Has("tc-drop"):
		return g.drop(x)
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

// sendOf is what a tick box sends: the form it submits, the name and value
// of its button or checkbox; "" when it sends nothing.
func sendOf(t *El) string {
	var f []string
	for _, a := range []kv{{"form", "Form"}, {"name", "Name"}, {"value", "Value"}} {
		if a.k == "form" && t.Attr("type") == "checkbox" {
			a.v = "Of"
		}
		if v := t.Attr(a.k); v != "" {
			f = append(f, a.v+": "+q(v))
		}
	}
	if len(f) == 0 {
		return ""
	}
	if t.Attr("type") == "radio" {
		f = append(f, "One: true")
	}
	return "w.Send{" + join(f, ", ") + "}"
}

// tickedOf is a tick box's state: a checkbox by its checked, a button by its
// classes.
func tickedOf(t *El) string {
	switch {
	case t.Has("is-part"):
		return "w.TickedPart"
	case t.Tag == "input" && !t.HasAttr("checked"), t.Tag != "input" && t.Has("is-off"):
		return "w.TickedOff"
	}
	return "w.TickedOn"
}

// hiddens are a form's hidden fields.
func hiddens(es []*El) string {
	var out []string
	for _, h := range es {
		out = append(out, "{Name: "+q(h.Attr("name"))+", Value: "+q(h.Attr("value"))+"}")
	}
	return "[]w.Hidden{" + join(out, ", ") + "}"
}

// postForm is a form without visible fields that buttons submit by its ID.
func (g *gen) postForm(e *El) string {
	f := []string{"ID: " + q(e.Attr("id")), "Action: " + q(e.Attr("action"))}
	var hid []*El
	file, accept, many := "", "", false
	for _, k := range e.Els() {
		if k.Attr("type") == "file" {
			file, accept, many = k.Attr("name"), k.Attr("accept"), k.HasAttr("multiple")
		} else {
			hid = append(hid, k)
		}
	}
	if len(hid) > 0 {
		f = append(f, "Fields: "+hiddens(hid))
	}
	if file != "" {
		f = append(f, "File: "+q(file))
		if accept != "" {
			f = append(f, "Accept: "+q(accept))
		}
		if many {
			f = append(f, "Many: true")
		}
	} else if e.Attr("enctype") == "multipart/form-data" {
		f = append(f, "Multipart: true")
	}
	if m := e.Attr("method"); m != "" && m != "post" {
		f = append(f, "Method: "+q(m))
	}
	return "w.PostForm{" + join(f, ", ") + "}"
}

// editBarAttrs are what an edit bar's button does with the script (Carrel's
// task 286): the Markdown it puts on the text, or the edit's Focus.
func editBarAttrs(e *El) []string {
	var out []string
	if v := e.Attr("data-tc-marks"); v != "" {
		out = append(out, "Marks: w.Mark"+strings.ToUpper(v[:1])+v[1:])
	}
	if e.HasAttr("data-tc-focus-mode") {
		out = append(out, "Focuses: true")
	}
	return out
}
