package framegen

import (
	"strings"
)

// wait stops a block the generator cannot build yet; some places print a
// stand-in instead (tryWait).
type wait struct{}

// tryWait is fn's result, or alt when fn waits.
func tryWait(fn func() string, alt string) (out string) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(wait); !ok {
				panic(r)
			}
			out = alt
		}
	}()
	return fn()
}

var (
	cardText = []kv{{"tc-card__line", "CardLine"}, {"tc-card__num", "CardNum"}, {"tc-card__term", "CardTerm"}, {"tc-card__note", "CardUsage"},
		{"tc-card__note-dusty", "CardSmallPrint"}, {"tc-card__why", "CardWhy"}, {"tc-card__warn", "CardWarn"}, {"tc-card__alert-note", "CardAlertText"},
		{"tc-card__lede", "CardLede"}, {"tc-card__caption", "CardCaption"}, {"tc-card__above", "CardAbove"}}
	cardRow = []kv{{"tc-card__acts", "CardActs"}, {"tc-card__top2", "CardTop2"}, {"tc-card__top8", "CardTop8"}, {"tc-card__top9", "CardTop9"},
		{"tc-card__top10", "CardTop10"}, {"tc-card__top11", "CardTop11"}, {"tc-card__files", "CardAttached"}, {"tc-card__foot", "CardFoot"}}
	colWidth = []kv{{"tc-table__col--tick", "ColTick"}, {"tc-table__col--tick-sm", "ColTickSmall"}, {"tc-table__col--size", "ColSize"},
		{"tc-table__col--type", "ColType"}, {"tc-table__col--time", "ColTime"}, {"tc-table__col--num", "ColNum"}, {"tc-table__col--menu", "ColMenu"},
		{"tc-table__col--changed", "ColChanged"}, {"tc-table__col--w26", "Col26"}, {"tc-table__col--w80", "Col80"}, {"tc-table__col--w90", "Col90"},
		{"tc-table__col--w110", "Col110"}, {"tc-table__col--w120", "Col120"}}
	msgSpace = []kv{{"tc-msg--flush", "MsgFlush"}, {"tc-msg--flush-top2", "MsgFlushTop2"}, {"tc-msg--flush-top4", "MsgFlushTop4"},
		{"tc-msg--flush-top11", "MsgFlushTop11"}, {"tc-msg--flush-bot12", "MsgFlushBottom12"}, {"tc-msg--mr0-ml0", "MsgFlushSides"},
		{"tc-msg--flush-top9", "MsgFlushTop9"}, {"tc-msg--flush-r0", "MsgFlushEnd"}, {"tc-msg--mx20", "MsgInset20"}, {"tc-msg--m0-14-12", "MsgSheet"}}
)

func (g *gen) msg(e *El) string {
	var f []string
	if s := svgs(e); len(s) > 0 {
		f = append(f, "Icon: "+icon(s[0]))
		f = append(f, "Parts: "+g.tparts(e.Kid("tc-msg__aside"), "[]w.MsgPart{"))
		if pu := e.Kid("tc-push"); pu != nil {
			f = append(f, "End: &w.Push{Items: []w.PushPart{"+g.buttons(pu.Els())+"}}")
		}
	} else {
		f = append(f, "Parts: "+g.tparts(e, "[]w.MsgPart{"))
	}
	if e.Has("is-alert") {
		f = append(f, "Tone: w.MsgAlert")
	}
	if e.Has("is-accent") {
		f = append(f, "Tone: w.MsgAccent")
	}
	for _, v := range allOf(e, msgSpace) {
		f = append(f, "Space: w."+v)
	}
	if e.Has("tc-grow") {
		f = append(f, "Place: w.Grown")
	}
	f = append(f, scripted(e)...)
	return "w.Msg{" + join(f, ", ") + "}"
}

func (g *gen) diag(e *El) string {
	s := "w.Diag{Text: " + q(e.TextRaw())
	if v := firstOf(e, []kv{{"tc-diag--top9-ink", "DiagInk"}, {"tc-diag--danger", "DiagDanger"}}); v != "" {
		s += ", Variant: w." + v
	}
	if id := e.Attr("id"); id != "" {
		s += ", ID: " + q(id)
	}
	return s + placef(e) + "}"
}

func (g *gen) danger(e *El) string {
	f := []string{"Button: " + g.button(e.Kid("tc-btn"))}
	if t := e.OwnText(); t != "" {
		f = append(f, "Note: "+q(t))
	}
	for _, v := range allOf(e, []kv{{"tc-danger--topauto-flush-padt11", "DangerAtEnd"}, {"tc-danger--top0-pad11-14", "DangerTight"}}) {
		f = append(f, "Space: w."+v)
	}
	return "w.Danger{" + join(f, ", ") + "}"
}

func (g *gen) meter(e *El) string {
	bar := e.Kid("tc-meter__bar")
	w := "0%"
	if bar.HasAttr("data-tc-width") {
		w = bar.Attr("data-tc-width")
	}
	f := []string{"Percent: " + strings.TrimRight(w, "%")}
	if t := e.Attr("data-tc-width"); t != "" {
		f = append(f, "Track: "+strings.TrimRight(t, "%"))
	}
	if bar.Has("tc-meter__bar--alert") {
		f = append(f, "Alert: true")
	}
	for _, v := range allOf(e, []kv{{"tc-meter--w60", "Meter60"}, {"tc-meter--max260-fill", "MeterFill"}}) {
		f = append(f, "Width: w."+v)
	}
	return "w.Meter{" + join(f, ", ") + "}"
}

func (g *gen) runbar(e *El) string {
	f := []string{"Meter: " + g.meter(e.Kid("tc-meter"))}
	if p := e.Attr("data-tc-progress"); p != "" {
		f = append(f, "Progress: "+q(p))
		if !e.HasAttr("hidden") {
			f = append(f, "Saving: true")
		}
	}
	var ps []string
	for _, k := range e.Kids {
		switch v := k.(type) {
		case string:
			if t := textpart(e, v); t != "" {
				ps = append(ps, t)
			}
		case *El:
			switch {
			case v.Has("tc-meter"):
			case v.Has("tc-dim--danger"):
				ps = append(ps, g.span(v, "FaceDanger"))
			case v.Has("tc-dim--accent"):
				ps = append(ps, g.span(v, "FaceAccent"))
			case v.Has("tc-mono"):
				ps = append(ps, g.mono(v))
			case v.Has("tc-line"):
				ps = append(ps, g.line(v))
			case v.Has("tc-drop"):
				ps = append(ps, g.drop(v))
			case v.Has("tc-push"):
				ps = append(ps, g.push(v))
			default:
				stop("runbar? " + join(v.Cls(), " "))
			}
		}
	}
	f = append(f, "Parts: []w.RunbarPart{"+join(ps, ", ")+"}")
	if e.Has("tc-runbar--danger") {
		f = append(f, "Failed: true")
	}
	return "w.Runbar{" + join(f, ", ") + "}"
}

func (g *gen) queue(e *El) string {
	h := e.Kid("tc-queue__head")
	f := []string{"Title: " + q(h.Els()[0].Text()), "Total: w.Mono{Text: " + q(h.Kid("tc-mono").Text()) + ", Variant: w.MonoDusty, Place: w.Pushed}",
		"Toggle: " + g.button(h.Kid("tc-btn"))}
	if l := e.Kid("tc-queue__list"); l != nil {
		var its []string
		for _, li := range l.Els() {
			row := li.Kid("tc-queue__row")
			failed := row.Kid("tc-dim--danger") != nil
			x := []string{"Name: " + q(row.Els()[0].Text()), "Meta: " + q(row.Els()[1].Text())}
			if m := li.Kid("tc-meter"); m != nil {
				x = append(x, "Meter: &"+g.meter(m))
			}
			if failed {
				x = append(x, "Failed: true")
			}
			if a := li.Kid("tc-queue__acts"); a != nil {
				x = append(x, "Acts: []w.Button{"+g.buttons(a.Els())+"}")
			}
			its = append(its, "{"+join(x, ", ")+"}")
		}
		f = append(f, "Items: []w.QueueItem{"+join(its, ", ")+"}")
	}
	if e.Has("is-min") {
		f = append(f, "Min: true")
	}
	for _, v := range allOf(e, []kv{{"tc-queue--flush-full", "QueueFlushFull"}, {"tc-queue--flush-wauto", "QueueFlushAuto"}}) {
		f = append(f, "Variant: w."+v)
	}
	if h := e.Attr("data-tc-fold-hide"); h != "" {
		f = append(f, "Fold: &w.QueueFold{Hide: "+q(h)+", Show: "+q(e.Attr("data-tc-fold-show"))+", Total: "+q(e.Attr("data-tc-fold-total"))+"}")
	}
	return "w.Queue{" + join(f, ", ") + "}"
}

func (g *gen) stat(e *El) string {
	s := "w.Stat{Value: " + q(e.Els()[0].Text()) + ", Label: " + q(e.OwnText())
	if e.Has("is-error") {
		s += ", Error: true"
	}
	return s + "}"
}

func (g *gen) stats(e *El) string {
	return "w.Stats{Items: []w.Stat{" + g.each(e.Els(), g.stat) + "}}"
}

func (g *gen) steps(e *El) string {
	var items []string
	on := 0
	n := 0
	for _, k := range e.Els() {
		if !k.Has("tc-steps__item") {
			continue
		}
		n++
		if on == 0 && k.Has("is-on") {
			on = n
		}
		items = append(items, q(k.Text()))
	}
	s := "w.Steps{Items: []string{" + join(items, ", ") + "}"
	if on > 0 {
		s += ", On: " + itoa(on)
	}
	return s + "}"
}

func (g *gen) diff(e *El) string {
	var its []string
	for _, it := range e.Els() {
		r := it.Kid("tc-rubric")
		title := `w.Rubric{Text: "", Lead: &w.Mono{Text: ` + q(r.Kid("tc-mono").Text()) + "}, Items: []w.RubricPart{w.RubricAside{Text: " + q(r.Kid("tc-rubric__aside").Text()) + "}}}"
		var opts []string
		for _, c := range it.Kid("tc-diff__options").Els() {
			tick := c.Kid("tc-tickrow__tick")
			v := c.Kid("tc-diff__value")
			on := ""
			if b := c.Kid("tc-check__box"); b != nil {
				// a version of a property is a radio of its form (Carrel's task 132)
				on = "Name: " + q(b.Attr("name")) + ", Value: " + q(b.Attr("value")) + ", Radio: true, "
				if b.HasAttr("checked") {
					on += "On: true, "
				}
			} else if !tick.Has("is-off") {
				on = "On: true, "
			}
			opts = append(opts, "w.Check{"+on+"Parts: []w.CheckPart{w.Choice{Source: "+q(v.Kid("tc-diff__source").Text())+", Value: "+q(v.OwnText())+"}}}")
		}
		its = append(its, "{Title: "+title+", Options: []w.Check{"+join(opts, ", ")+"}}")
	}
	s := "w.Diff{Items: []w.DiffItem{\n\t\t" + join(its, ",\n\t\t") + "}"
	if e.Tag == "form" {
		s += ", ID: " + q(e.Attr("id")) + ", Action: " + q(e.Attr("action"))
	}
	return s + "}"
}

var emptyVariant = []kv{{"tc-empty--t12-pad14-20", "EmptySmall"}, {"tc-empty--dusty", "EmptyDusty"}, {"tc-empty--t12-pad12-20", "EmptySmallTop"},
	{"tc-empty--dashed", "EmptyDashed"}, {"tc-empty--drop", "EmptyDrop"}, {"tc-empty--drop-side", "EmptyDropSide"}}

func (g *gen) empty(e *El) string {
	var f []string
	if s := svgs(e); len(s) > 0 {
		f = append(f, "Icon: "+icon(s[0]))
	}
	lead := e.Kid("tc-empty__lead")
	if lead == nil {
		lead = e.Kid("tc-empty__lead-sm")
	}
	if lead != nil {
		f = append(f, "Lead: "+q(lead.Text()))
	}
	var ps []string
	for _, k := range e.Kids {
		switch v := k.(type) {
		case string:
			if t := textpart(e, v); t != "" {
				ps = append(ps, t)
			}
		case *El:
			switch {
			case v.Tag == "svg" || v.Has("tc-empty__lead") || v.Has("tc-empty__lead-sm"):
			case v.Has("tc-btn"):
				ps = append(ps, g.button(v))
			case v.Has("tc-line"):
				// the list of the picked files
				ps = append(ps, g.line(v))
			case v.Has("tc-dim--t12-5") || v.Has("tc-dim--t12"):
				face := "FaceSmall"
				if v.Has("tc-dim--t12-5") {
					face = "FaceLarge"
				}
				var its []string
				for _, y := range v.Kids {
					switch z := y.(type) {
					case string:
						if strings.TrimFunc(z, isSpace) != "" {
							its = append(its, textpart(v, z))
						}
					case *El:
						if z.Has("tc-mono") {
							its = append(its, g.mono(z))
						} else if z.Tag == "b" {
							its = append(its, "w.Strong("+q(z.Text())+")")
						}
					}
				}
				ps = append(ps, "w.Aside{Face: w."+face+", Parts: []w.Inline{"+join(its, ", ")+"}}")
			case v.Has("tc-dim"):
				ps = append(ps, g.dimx(v))
			case v.Has("tc-empty__pick"):
				// words that pick a file: the label of its input (Carrel's task 119)
				pick := "w.FilePick{Label: " + q(v.Kid("tc-dim--accent").Text()) + ", File: " + g.fileField(v.Kid("tc-file"))[1:]
				if tpl := rowTemplate(v); tpl != nil {
					pick += ", Row: &" + g.line(tpl.Els()[0])
				}
				ps = append(ps, pick+"}")
			case v.Tag == "b" && v.Has("tc-dim--accent"):
				ps = append(ps, "w.Emphasis("+q(v.Text())+")")
			case v.Tag == "b" && len(v.Cls()) == 0:
				ps = append(ps, "w.Strong("+q(v.Text())+")")
			case v.Has("tc-mono"):
				ps = append(ps, g.mono(v))
			case v.Has("tc-kbd"):
				ps = append(ps, g.kbd(v))
			case v.Tag == "a" && v.Has("tc-dim--accent"):
				ps = append(ps, "w.Span{Face: w.FaceAccent, Text: "+q(v.Text())+", Href: "+q(v.Attr("href"))+"}")
			default:
				stop("empty? " + join(v.Cls(), " "))
			}
		}
	}
	if len(ps) > 0 {
		f = append(f, "Parts: []w.EmptyPart{"+join(ps, ", ")+"}")
	}
	if v := firstOf(e, emptyVariant); v != "" {
		f = append(f, "Variant: w."+v)
	}
	return "w.Empty{" + join(f, ", ") + "}"
}

func (g *gen) thumb(e *El) string {
	var f []string
	f = append(f, "Icon: "+icon(svgs(e)[0]))
	if img := e.Kid("tc-thumb__img"); img != nil {
		f = append(f, "Src: "+q(img.Attr("src")), "Alt: "+q(img.Attr("alt")))
	}
	for _, v := range allOf(e, []kv{{"tc-thumb--on-card", "ThumbOnCard"}, {"tc-thumb--preview", "ThumbPreview"}, {"tc-thumb--wide", "ThumbWide"}, {"tc-thumb--tile", "ThumbSquare"}, {"tc-thumb--folder", "ThumbFolder"}}) {
		f = append(f, "Form: w."+v)
	}
	return "w.Thumb{" + join(f, ", ") + "}"
}

func (g *gen) tile(e *El) string {
	s := "w.Tile{Thumb: " + g.thumb(e.Kid("tc-thumb")) + ", Name: " + q(e.Kid("tc-tile__name").Text()) + ", Meta: " + q(e.Kid("tc-tile__meta").Text())
	box := e.Kid("tc-tile__box")
	if e.Has("is-sel") {
		s += ", Selected: true"
	}
	if h := e.Kid("tc-tile__name").Attr("href"); h != "" {
		s += ", Href: " + q(h)
	}
	if box != nil {
		s += ", Send: " + sendOf(box)
	}
	return s + "}"
}

func (g *gen) grid(e *El) string {
	s := "w.Grid{Tiles: []w.Tile{" + g.each(e.Els(), g.tile) + "}"
	if e.Has("tc-grid--flush-fill-pad8") {
		s += ", Flush: true"
	}
	return s + "}"
}

func (g *gen) pad(e *El) string {
	return "w.Pad{Items: []w.PadPart{" + g.each(e.Els(), func(k *El) string {
		if k.Has("tc-thumb") {
			return g.thumb(k)
		}
		return g.card(k)
	}) + "}}"
}

type fieldsVariant struct{ a, b, v string }

var fieldsVariants = []fieldsVariant{{"tc-fields--cols88", "tc-fields--pad4-14-12", "Fields88Pad"}, {"tc-fields--cols88", "", "Fields88"},
	{"tc-fields--bleed", "tc-fields--cols130", "FieldsBleed130"}, {"tc-fields--bleed", "", "FieldsBleed"}, {"tc-fields--cols130", "tc-fields--flush-padt14", "Fields130Flush"}}

func (g *gen) fields(e *El) string {
	var its []string
	ch := e.Els()
	for i := 0; i+1 < len(ch); i += 2 {
		dt, dd := ch[i], ch[i+1]
		face := ""
		switch {
		case dd.Has("tc-mono--accent"):
			face = "ValueMonoAccent"
		case dd.Has("tc-mono--t11"):
			face = "ValueMonoSmall"
		case dd.Has("tc-mono"):
			face = "ValueMono"
		case dd.Has("tc-dim--accent"):
			face = "ValueAccent"
		}
		var vp []string
		href := ""
		var after []string
		if a := dd.Kid("tc-fields__link"); a != nil {
			href = a.Attr("href")
			// what stands after the link, outside it: the collection a value came from
			past := false
			for _, k := range dd.Els() {
				switch {
				case k == a:
					past = true
				case past && k.Has("tc-mark"):
					after = append(after, g.mark(k))
				case past && k.Has("tc-dim"):
					after = append(after, g.dimx(k))
				case past && k.Has("tc-badge"):
					after = append(after, g.badge(k))
				case past:
					stop("dd after its link? " + join(k.Cls(), " "))
				}
			}
			dd = a
		}
		for _, k := range dd.Kids {
			switch v := k.(type) {
			case string:
				if t := textpart(dd, v); t != "" {
					vp = append(vp, t)
				}
			case *El:
				switch {
				case v.Tag == "br":
					vp = append(vp, "w.Break{}")
				case v.Has("tc-mark"):
					vp = append(vp, g.mark(v))
				case v.Has("tc-badge"):
					vp = append(vp, g.badge(v))
				case v.Has("tc-tag"):
					vp = append(vp, g.tag(v))
				case v.Has("tc-mono"):
					vp = append(vp, g.mono(v))
				case v.Has("tc-dim--danger"):
					vp = append(vp, g.span(v, "FaceDanger"))
				case v.Has("tc-dim"):
					vp = append(vp, g.dimx(v))
				default:
					stop("dd? " + join(v.Cls(), " "))
				}
			}
		}
		vp = trimEnds(vp)
		s := "{Term: " + q(dt.Text()) + ", Value: []w.ValuePart{" + join(vp, ", ") + "}"
		if face != "" {
			s += ", Face: w." + face
		}
		if href != "" {
			s += ", Href: " + q(href)
		}
		if len(after) > 0 {
			s += ", After: []w.ValuePart{" + join(after, ", ") + "}"
		}
		its = append(its, s+"}")
	}
	f := []string{"Items: []w.FieldsItem{\n\t\t" + join(its, ",\n\t\t") + "}"}
	for _, x := range fieldsVariants {
		if e.Has(x.a) && (x.b == "" || e.Has(x.b)) {
			f = append(f, "Variant: w."+x.v)
			break
		}
	}
	s := "w.Fields{" + join(f, ", ") + "}"
	if p := placef(e); p != "" && !strings.Contains(s, "Place:") {
		s = s[:len(s)-1] + p + "}"
	}
	return s
}

// anyblock is the widget a block element is, for a place that takes many.
func (g *gen) anyblock(k *El) string {
	table := []struct {
		c  string
		fn func(*El) string
	}{{"tc-msg", g.msg}, {"tc-diag", g.diag}, {"tc-danger", g.danger}, {"tc-meter", g.meter}, {"tc-runbar", g.runbar}, {"tc-queue", g.queue},
		{"tc-stats", g.stats}, {"tc-steps", g.steps}, {"tc-diff", g.diff}, {"tc-empty", g.empty}, {"tc-grid", g.grid}, {"tc-thumb", g.thumb},
		{"tc-fields", g.fields}, {"tc-sec", g.sec}, {"tc-card", g.card}, {"tc-table", g.table},
		{"tc-prose", g.prose}, {"tc-list", g.list}, {"tc-tree", g.tree}, {"tc-head", g.head}, {"tc-bar", g.bar}, {"tc-rubric", func(x *El) string { return g.rubric(x, "") }},
		{"tc-tickrow", g.tickrow}, {"tc-railsec", g.railsec}, {"tc-docmeta", g.docmeta}, {"tc-pad", g.pad}}
	for _, t := range table {
		if k.Has(t.c) {
			return t.fn(k)
		}
	}
	switch {
	case k.Has("tc-docbar"):
		return g.docbar(k)
	case k.Has("tc-doc"):
		return g.doc(k)
	case k.Has("tc-sheet__block"):
		var lines []string
		for _, it := range k.Kid("tc-line__text").Els() {
			lines = append(lines, "{Name: "+q(it.Els()[0].Text())+", Note: w.Dim{Parts: []w.Inline{w.Text("+q(it.Kid("tc-dim").Text())+")}}}")
		}
		return "w.SheetBlock{Title: " + g.rubric(k.Kid("tc-rubric"), "") + ", Items: []w.SheetLine{" + join(lines, ", ") + "}}"
	case k.Has("tc-diff__agree"):
		return "w.DiffAgree{Property: w.Mono{Text: " + q(k.Kid("tc-mono").Text()) + "}, Text: " + q(k.OwnText()) + "}"
	case k.Has("tc-seg"):
		return g.barItem(k)
	case k.Has("tc-check"):
		return g.check(k)
	case k.Has("tc-field"):
		return g.field(k)
	case k.Has("tc-form"):
		return g.form(k)
	case k.Has("tc-formfoot"):
		return g.formFoot(k)
	case k.Has("tc-btn"):
		return g.button(k)
	case k.Has("tc-line"):
		return g.line(k)
	case k.Has("tc-mono"):
		return g.mono(k)
	case k.Has("tc-in"):
		return g.control(k)
	}
	panic(wait{})
}

func (g *gen) sec(e *El) string {
	var items []string
	for _, k := range e.Els() {
		switch {
		case k.Has("tc-sec__list"):
			gap := 0
			for _, n := range []int{5, 6} {
				if k.Has("tc-stack--gap-" + itoa(n)) {
					gap = n
				}
			}
			s := "w.SecList{Items: [][]w.SecItemPart{" + g.each(k.Els(), func(li *El) string { return g.tparts(li, "[]w.SecItemPart{") }) + "}"
			if gap > 0 {
				s += ", Gap: " + itoa(gap)
			}
			if k.Has("tc-sec__list--muted") {
				s += ", Muted: true"
			}
			if k.Has("tc-sec__list--cols") {
				s += ", Columns: true"
			}
			items = append(items, s+"}")
		case k.Has("tc-field__hint"):
			s := "w.SecHint{Text: " + q(k.Text())
			if k.Has("tc-field__hint--top6") {
				s += ", Apart: true"
			}
			items = append(items, s+"}")
		case k.Has("tc-sec__note"):
			items = append(items, "w.SecNote{Text: "+q(k.Text())+"}")
		case k.Has("tc-sec__acts"):
			items = append(items, "w.SecActs{Items: []w.Button{"+g.buttons(k.Els())+"}}")
		case k.Has("tc-doc__list"):
			items = append(items, "w.SecTasks{Items: []w.DocTask{"+g.each(k.Els(), docTask)+"}}")
		case k.Has("tc-card__top8"):
			var fs []string
			for _, ln := range k.Els() {
				fl := ln.Kid("tc-table__file")
				fs = append(fs, "{Icon: "+icon(fl.Els()[0])+", Name: "+q(fl.Els()[1].Text())+", Size: "+g.mono(ln.Kid("tc-mono"))+", Tag: "+
					g.tag(ln.Kid("tc-tag"))+", Remove: "+g.button(ln.Kid("tc-btn"))+"}")
			}
			items = append(items, "w.SecFiles{Items: []w.SecFile{"+join(fs, ", ")+"}}")
		case k.Has("tc-line__text") && k.Has("tc-stack"):
			var ps []string
			for _, x := range k.Els() {
				switch {
				case x.Has("tc-stack__item"):
					ps = append(ps, "w.SecOpened{Name: "+q(x.OwnText())+", Size: w.Dim{Parts: []w.Inline{w.Text("+q(x.Kid("tc-dim").Text())+")}}}")
				case x.Has("tc-line"):
					ps = append(ps, g.line(x))
				default:
					ps = append(ps, g.dimx(x))
				}
			}
			items = append(items, "w.SecOpen{Items: []w.SecOpenPart{"+join(ps, ", ")+"}}")
		case k.Has("tc-rubric"):
			items = append(items, g.rubric(k, ""))
		case g.docText(k) != "":
			items = append(items, g.docText(k)) // a note's text in its panel (Carrel's task 125)
		default:
			items = append(items, g.anyblock(k))
		}
	}
	s := "w.Sec{Items: []w.SecPart{\n\t\t" + join(items, ",\n\t\t") + "}"
	if e.Has("tc-sec--flush-flush-padt11") {
		s += ", Flush: true"
	}
	return s + "}"
}

func docTask(li *El) string {
	d := ""
	if li.Kid("tc-tick").Has("is-done") {
		d = "Done: true, "
	}
	if li.Kid("tc-tick").Tag == "span" {
		return "{Tick: w.Tick{" + d + "Variant: w.TickNudge, Sign: true}, Text: " + q(li.Text()) + "}"
	}
	return "{Tick: w.Tick{" + d + "Variant: w.TickNudge, Label: " + q(li.Text()) + "}, Text: " + q(li.Text()) + "}"
}

func (g *gen) rowPart(x *El) string {
	switch {
	case x.Has("tc-btn"):
		return g.button(x)
	case x.Has("tc-field"):
		return g.field(x)
	case x.Has("tc-tag"):
		return g.tag(x)
	case x.Tag == "svg" && x.Has("tc-ico--turn"):
		return "w.Turn{}"
	case x.Has("tc-field__hint--flush-fill"):
		return "w.FillHint{Parts: " + g.tparts(x, "[]w.CardTextPart{") + "}"
	case x.Has("tc-dim--t12-5") && !x.Has("tc-dim"):
		return "w.Span{Face: w.FaceLarge, Text: " + q(x.Text()) + "}"
	case x.Has("tc-dim"):
		return g.dimx(x)
	case x.Has("tc-in"):
		// by its tag, as anywhere (the review of 14.10: it took every
		// tc-in for a line to type in)
		if c := g.control(x); c != "" {
			return c
		}
	}
	stop("card row? " + join(x.Cls(), " "))
	return ""
}

type cardVariant = fieldsVariant

var cardVariants = []cardVariant{{"tc-card--edge-accent", "tc-card--flush", "CardChosenFlushGrow"}, {"tc-card--edge-accent-flush", "", "CardChosenFlush"},
	{"tc-card--edge-accent", "", "CardChosen"}, {"tc-card--alert", "", "CardAlert"}, {"tc-card--top14", "", "CardApart"},
	{"tc-card--pad12-14", "", "CardSnug"}, {"tc-card--flush", "", "CardFlush"}}

func (g *gen) card(e *El) string {
	var items []string
	for _, k := range e.Els() {
		kind := firstOf(k, cardText)
		row := firstOf(k, cardRow)
		switch {
		case kind != "":
			items = append(items, "w.CardText{Kind: w."+kind+", Parts: "+g.tparts(k, "[]w.CardTextPart{")+"}")
		case k.Has("tc-field") && k.Has("tc-card__top9"):
			items = append(items, "w.SpacedField{Field: "+g.field(k)+"}")
		case row != "":
			gap := 0
			for _, n := range []int{6, 7, 8, 10, 22} {
				if k.Has("tc-line--gap-" + itoa(n)) {
					gap = n
				}
			}
			s := "w.CardRow{Kind: w." + row
			if gap > 0 {
				s += ", Gap: " + itoa(gap)
			}
			items = append(items, s+", Items: []w.CardRowPart{"+g.each(k.Els(), g.rowPart)+"}}")
		case k.Has("tc-field__hint"):
			s := "w.CardHint{Parts: " + g.tparts(k, "[]w.CardTextPart{")
			if k.Tag == "div" {
				s += ", Line: true"
			}
			items = append(items, s+"}")
		case k.Has("tc-card__list"):
			items = append(items, "w.CardList{Items: []string{"+g.each(k.Els(), func(li *El) string { return q(li.Text()) })+"}}")
		case k.Has("tc-card__fields"):
			s := "w.CardFields{Items: []w.CardFieldsPart{" + g.each(k.Els(), func(x *El) string {
				switch {
				case x.Has("tc-field"):
					return g.field(x)
				case x.Has("tc-btn"):
					return g.button(x)
				}
				return "w.Span{Face: w.FaceMuted12, Text: " + q(x.Text()) + "}"
			}) + "}"
			if k.Has("tc-card__fields--gap10-12") {
				s += ", Wide: true"
			}
			if k.Has("tc-card__fields--rule") {
				s += ", Rule: true"
			}
			items = append(items, s+"}")
		case k.Has("tc-card__split"):
			qr := k.Kid("tc-card__qr")
			side := k.Kid("tc-card__qr-side")
			sp := g.each(side.Els(), func(x *El) string {
				if x.Has("tc-btn") {
					return g.button(x)
				}
				return g.dimx(x)
			})
			if img := qr.Kid("tc-card__qr-img"); img != nil {
				items = append(items, "w.CardQR{Colour: "+q(qr.Attr("data-tc-colour"))+", Src: "+q(img.Attr("src"))+", Side: []w.CardQRPart{"+sp+"}}")
				break
			}
			var use *El
			for _, x := range qr.Els()[0].Els() {
				if x.Tag == "use" {
					use = x
					break
				}
			}
			items = append(items, "w.CardQR{Colour: "+q(qr.Attr("data-tc-colour"))+", Symbol: "+q(use.Attr("href")[1:])+", Side: []w.CardQRPart{"+sp+"}}")
		case k.Has("tc-card__pick"):
			tick, item := k.Els()[0], k.Els()[1]
			var x []string
			if tick.Has("is-off") {
				x = append(x, "Ticked: w.TickedOff")
			}
			x = append(x, "Title: "+q(item.Kid("tc-stack__item-title").Text()))
			x = append(x, "Note: "+g.tparts(item.Kid("tc-stack__item-note"), "[]w.CardTextPart{"))
			if a := item.Kid("tc-stack__item-aside"); a != nil {
				x = append(x, "Aside: "+q(a.Text()))
			}
			if fl := item.Kid("tc-field"); fl != nil {
				x = append(x, "Field: &"+g.field(fl))
			}
			items = append(items, "w.CardPick{"+join(x, ", ")+"}")
		case k.Has("tc-card__grid"):
			items = append(items, "w.CardGrid{Items: []w.CardGridPart{"+g.each(k.Els(), func(x *El) string {
				switch {
				case x.Has("tc-dim"):
					return g.dimx(x)
				case x.Tag == "textarea":
					return `w.Area{Text: pasteArgs, Variant: w.AreaPaste, Mono: true, Label: "Arguments"}`
				case x.Has("tc-in"):
					return g.input(x)
				}
				return g.button(x)
			})+"}}")
		case k.Has("tc-card__rows"):
			items = append(items, "w.CardRows{Items: []w.CardRowsItem{"+g.each(k.Els(), func(it *El) string {
				return "{Head: w.Mono{Text: " + q(it.Kid("tc-mono").Text()) + "}, Note: w.Dim{Parts: " + g.inl(it.Kid("tc-dim")) + "}}"
			})+"}}")
		case k.Has("tc-card__head"):
			items = append(items, "w.CardHead{Lines: []w.CardHeadLine{"+g.each(k.Els(), func(it *El) string {
				slot := it.Kid("tc-dim--slot")
				rest := &El{Tag: "x"}
				for _, x := range it.Kids {
					if x != any(slot) {
						rest.Kids = append(rest.Kids, x)
					}
				}
				return "{Slot: " + q(slot.Text()) + ", Parts: " + g.tparts(rest, "[]w.CardTextPart{") + "}"
			})+"}}")
		case k.Has("tc-line") && !hasElementClass(k):
			items = append(items, g.line(k))
		case k.Has("tc-mono") && k.Tag == "div":
			items = append(items, g.mono(k))
		default:
			items = append(items, g.anyblock(k))
		}
	}
	f := []string{"Items: []w.CardPart{\n\t\t" + join(items, ",\n\t\t") + "}"}
	for _, x := range cardVariants {
		if e.Has(x.a) && (x.b == "" || e.Has(x.b)) {
			f = append(f, "Variant: w."+x.v)
			break
		}
	}
	if e.Has("tc-grow") {
		f = append(f, "Place: w.Grown")
	}
	if e.Tag == "form" {
		// a group of duplicates sends its choices (Carrel's task 131)
		if id := e.Attr("id"); id != "" {
			f = append(f, "ID: "+q(id))
		}
		f = append(f, "Action: "+q(e.Attr("action")))
	}
	return "w.Card{" + join(f, ", ") + "}"
}

// hasElementClass says whether e has a class of a block's element ("__").
func hasElementClass(e *El) bool {
	for _, c := range e.Cls() {
		if strings.Contains(c, "__") {
			return true
		}
	}
	return false
}

func (g *gen) cellPart(x *El) string {
	switch {
	case x.Has("tc-table__file"):
		m := x.Kid("tc-mark")
		var name *El
		for _, y := range x.Els() {
			if y.Has("tc-table__filename") && !y.Has("tc-mark") {
				name = y
				break
			}
		}
		var f []string
		if m != nil {
			f = append(f, "Mark: &"+g.mark(m))
		} else {
			f = append(f, "Icon: "+icon(x.Els()[0]))
		}
		f = append(f, "Name: "+q(name.Text()))
		if name.Has("tc-table__filename--accent") {
			f = append(f, "Accent: true")
		}
		if h := name.Attr("href"); h != "" {
			f = append(f, "Href: "+q(h))
		}
		return "w.TableFile{" + join(f, ", ") + "}"
	case x.Tag == "br":
		return "w.Break{}"
	case x.Has("tc-tick"):
		v := ""
		if x.Has("tc-tick--edge-danger") {
			v = ", Variant: w.TickEdgeDanger"
		} else if x.Has("tc-tick--edge-dusty") {
			v = ", Variant: w.TickEdgeDusty"
		}
		d := ""
		if x.Has("is-done") {
			d = "Done: true, "
		}
		lab := x.Attr("aria-label")
		if lab == "" {
			lab = "Done"
		}
		if s := sendOf(x); s != "" {
			v += ", Send: " + s
		}
		return "w.Tick{" + d + "Label: " + q(lab) + v + "}"
	case x.Has("tc-btn"):
		return g.button(x)
	case x.Has("tc-badge"):
		return g.badge(x)
	case x.Has("tc-tag"):
		return g.tag(x)
	case x.Has("tc-drop"):
		return g.drop(x)
	case x.Has("tc-table__wrap"):
		return "w.CellWrap{Items: []w.CellPart{" + g.each(x.Els(), g.cellPart) + "}}"
	case x.Has("tc-line"):
		return g.line(x)
	case x.Has("tc-mono"):
		return g.mono(x)
	case x.Has("tc-segfold"):
		return g.segfold(x)
	case x.Has("tc-seg"):
		return g.seg(x)
	case x.Tag == "svg" && x.Has("tc-ico--inline"):
		return "w.Chevron{}"
	case x.Has("tc-dim"):
		return g.dimx(x)
	}
	stop("cell? " + x.Tag + " " + join(x.Cls(), " "))
	return ""
}

var rowState = []kv{{"is-sel", "RowPicked"}, {"is-folder", "RowFolder"}, {"is-off", "RowOff"}, {"is-on", "RowOn"},
	{"tc-table__row--dim", "RowFaded"}, {"tc-table__row--off", "RowGone"}}

func anySuffix(e *El, suf string) bool {
	for _, c := range e.Cls() {
		if strings.HasSuffix(c, suf) {
			return true
		}
	}
	return false
}

func (g *gen) table(e *El) string {
	var cols []string
	var head *El
	for _, x := range e.Els() {
		if x.Tag == "thead" {
			head = x
			break
		}
	}
	if head != nil {
		for _, th := range head.Els()[0].Els() {
			var f []string
			if t := th.OwnText(); t != "" {
				f = append(f, "Label: "+q(t))
			}
			for _, v := range allOf(th, colWidth) {
				f = append(f, "Width: w."+v)
			}
			if th.Has("tc-table__n3") {
				f = append(f, "Rank: w.Rank3")
			}
			if th.Has("tc-table__n4") {
				f = append(f, "Rank: w.Rank4")
			}
			if th.Has("tc-table__n1") {
				f = append(f, "Rank: w.Rank1")
			}
			if th.Has("is-2nd") {
				f = append(f, "Secondary: true")
			}
			if t := th.Kid("tc-tickrow__tick"); t != nil {
				f = append(f, "Tick: ptr("+tickedOf(t)+")")
				if s := sendOf(t); s != "" {
					f = append(f, "Send: "+s)
				}
			}
			cols = append(cols, "{"+join(f, ", ")+"}")
		}
	}
	body := e
	for _, x := range e.Els() {
		if x.Tag == "tbody" {
			body = x
			break
		}
	}
	var rows []string
	for _, tr := range body.Els() {
		if tr.Tag != "tr" {
			continue
		}
		var cells []string
		for _, td := range tr.Els() {
			var f, parts []string
			more := ""
			for _, k := range td.Kids {
				switch v := k.(type) {
				case string:
					if t := textpart(td, v); t != "" {
						parts = append(parts, t)
					}
				case *El:
					if v.Has("tc-acts__more") {
						more = "More: &" + g.button(v)
					} else if v.Has("tc-drop") && v.Kid("tc-acts__more") != nil {
						// «⋯» with its menu (Carrel's task 117)
						more = "Drop: &" + g.drop(v)
					} else if v.Has("tc-tickrow__tick") {
						f = append(f, "Pick: ptr("+tickedOf(v)+")")
						if s := sendOf(v); s != "" {
							f = append(f, "Send: "+s)
						}
					} else {
						parts = append(parts, g.cellPart(v))
					}
				}
			}
			if len(parts) > 0 {
				f = append([]string{"Parts: []w.CellPart{" + join(parts, ", ") + "}"}, f...)
			}
			if more != "" {
				f = append(f, more)
			}
			if td.Has("tc-mono") {
				f = append(f, "Mono: true")
			}
			if td.Has("tc-dim") {
				f = append(f, "Dim: true")
			}
			if td.Has("tc-acts") {
				f = append(f, "Acts: true")
			}
			if td.Has("tc-table__name") {
				f = append(f, "Name: true")
			}
			if td.Has("is-ph") {
				f = append(f, "Unset: true")
			}
			if anySuffix(td, "--dusty") {
				f = append(f, "Tone: w.CellDusty")
			}
			if anySuffix(td, "--danger") {
				f = append(f, "Tone: w.CellDanger")
			}
			if td.Has("tc-dim--muted") {
				f = append(f, "Tone: w.CellMuted")
			}
			if anySuffix(td, "--t11") {
				f = append(f, "Small: true")
			}
			if v := td.Attr("colspan"); v != "" {
				f = append(f, "Span: "+v)
			}
			if v := td.Attr("rowspan"); v != "" {
				f = append(f, "Rows: "+v)
			}
			if head == nil {
				if td.Has("tc-table__n3") {
					f = append(f, "Rank: w.Rank3")
				}
				if td.Has("tc-table__n4") {
					f = append(f, "Rank: w.Rank4")
				}
				if td.Has("tc-table__n1") {
					f = append(f, "Rank: w.Rank1")
				}
				for _, v := range allOf(td, colWidth) {
					f = append(f, "Width: w."+v)
				}
			}
			cells = append(cells, "{"+join(f, ", ")+"}")
		}
		s := "{Cells: []w.TableCell{" + join(cells, ", ") + "}"
		if st := firstOf(tr, rowState); st != "" {
			s += ", State: w." + st
		}
		rows = append(rows, s+"}")
	}
	var f []string
	if len(cols) > 0 {
		f = append(f, "Cols: []w.TableCol{"+join(cols, ", ")+"}")
	}
	f = append(f, "Rows: []w.TableRow{\n\t\t"+join(rows, ",\n\t\t")+"}")
	for _, v := range allOf(e, []kv{{"tc-table--fit", "TableFit"}, {"tc-table--wide", "TableWide"}, {"tc-table--mx-20", "TableInset"}, {"tc-table--packed", "TablePacked"}}) {
		f = append(f, "Variant: w."+v)
	}
	return "w.Table{" + join(f, ", ") + placef(e) + "}"
}

func dimText(t string) string { return "w.Dim{Parts: []w.Inline{w.Text(" + q(t) + ")}}" }

func (g *gen) docmeta(e *El) string {
	var secs []string
	danger := ""
	for _, k := range e.Els() {
		if k.Has("tc-danger") {
			danger = g.danger(k)
			continue
		}
		var its []string
		for _, x := range k.Els()[1:] {
			switch {
			case x.Has("tc-diag"):
				its = append(its, g.diag(x))
			case x.Has("tc-line") && !x.Has("tc-line__text"):
				// tags, and where they come from when there was a choice (Р-58)
				its = append(its, "w.Line{Block: true, Gap: w.LineGap4, Items: []w.LinePart{"+g.each(x.Els(), g.linePart)+"}}")
			case x.Has("tc-stack"):
				var refs []string
				for _, it := range x.Els() {
					m := it.Kid("tc-mark")
					var a *El
					for _, y := range it.Els() {
						if y.Tag == "a" {
							a = y
							break
						}
					}
					s := "{"
					if m != nil {
						s += "Mark: &" + g.mark(m) + ", "
					}
					if a != nil {
						s += "Name: " + q(a.Text()) + `, Href: "#"`
					} else {
						s += "Name: " + q(it.Kid("tc-dim--accent").Text())
					}
					if d := it.Kid("tc-dim"); d != nil {
						s += ", What: " + dimText(d.Text())
					}
					refs = append(refs, s+"}")
				}
				its = append(its, "w.DocRefs{Items: []w.DocRef{"+join(refs, ", ")+"}}")
			case x.Has("tc-line__text"):
				its = append(its, g.docAttach(x))
			default:
				stop("docmeta? " + join(x.Cls(), " "))
			}
		}
		id := ""
		if v := k.Attr("id"); v != "" {
			id = ", ID: " + q(v)
		}
		secs = append(secs, "{Title: "+g.rub(k.Els()[0])+", Items: []w.DocmetaPart{"+join(its, ", ")+"}"+id+"}")
	}
	s := "w.Docmeta{Sections: []w.DocmetaSec{\n\t\t" + join(secs, ",\n\t\t") + "}"
	if danger != "" {
		s += ", Danger: &" + danger
	}
	return s + "}"
}

type proseVariant struct {
	cls []string
	v   string
}

var proseVariants = []proseVariant{{[]string{"tc-prose--lead", "tc-prose--top14"}, "ProseLead"}, {[]string{"tc-prose--lead", "tc-prose--top18"}, "ProseLeadAfter"},
	{[]string{"tc-prose--read", "tc-prose--top14"}, "ProseReadTop"}, {[]string{"tc-prose--muted-flush-bot10"}, "ProseMutedClose"},
	{[]string{"tc-prose--t13-flush-padt11"}, "Prose13Card"}, {[]string{"tc-prose--muted"}, "ProseMuted"}, {[]string{"tc-prose--t13"}, "Prose13"},
	{[]string{"tc-prose--read"}, "ProseRead"}, {[]string{"tc-prose--top14"}, "ProseTop"}}

func (g *gen) prose(e *El) string {
	v := ""
	for _, p := range proseVariants {
		all := true
		for _, c := range p.cls {
			if !e.Has(c) {
				all = false
			}
		}
		if all {
			v = ", Variant: w." + p.v
			break
		}
	}
	return "w.Prose{Parts: " + g.rawparts(e) + v + "}"
}

func (g *gen) rub(k *El) string {
	var items []string
	for _, x := range k.Els() {
		if x.Has("tc-btn") {
			items = append(items, g.button(x))
		} else if x.Has("tc-badge") {
			items = append(items, g.badge(x))
		} else if x.Has("tc-drop") {
			items = append(items, g.drop(x))
		}
	}
	s := "w.Rubric{Text: " + q(k.OwnText())
	if len(items) > 0 {
		s += ", Items: []w.RubricPart{" + join(items, ", ") + "}"
	}
	if v := firstOf(k, []kv{{"tc-rubric--bot7", "RubricBottom7"}, {"tc-rubric--bot8", "RubricBottom8"}}); v != "" {
		s += ", Variant: w." + v
	}
	return s + "}"
}

func (g *gen) tasks(k *El, phone bool) string {
	s := "w.DocTasks{Title: " + g.rub(k.Kid("tc-rubric")) + ", Items: []w.DocTask{" + g.each(k.Kid("tc-doc__list").Els(), docTask) + "}"
	if phone {
		s += ", Phone: true"
	}
	return s + "}"
}

func (g *gen) doc(e *El) string {
	var items []string
	var hidden []*El
	for _, k := range e.Els() {
		switch {
		case k.Tag == "input" && k.Attr("type") == "hidden":
			hidden = append(hidden, k)
		case k.Has("tc-doc__title"):
			items = append(items, "w.DocTitle{Text: "+q(k.Text())+"}")
		case k.Has("tc-head__sub--top6"):
			var ps []string
			for _, x := range k.Kids {
				switch v := x.(type) {
				case string:
					if strings.TrimFunc(v, isSpace) != "" {
						ps = append(ps, "w.Text("+q(collapse(v))+")")
					}
				case *El:
					if v.Has("tc-mono") {
						ps = append(ps, g.mono(v))
					} else if v.Has("tc-tag") {
						ps = append(ps, g.tag(v))
					}
				}
			}
			items = append(items, "w.DocSub{Parts: []w.DocSubPart{"+join(ps, ", ")+"}}")
		case k.Has("tc-prose"):
			items = append(items, g.prose(k))
		case k.Has("tc-doc__sec"):
			items = append(items, g.tasks(k, false))
		case k.Has("tc-doc__acts"):
			items = append(items, g.tasks(k, true))
		case k.Has("tc-doc__sec--rule"):
			var fs []string
			for _, it := range k.Kid("tc-doc__files").Els() {
				fs = append(fs, "{Name: "+q(it.Els()[0].Text())+", Size: "+dimText(it.Els()[1].Text())+"}")
			}
			items = append(items, "w.DocFiles{Title: "+g.rub(k.Kid("tc-rubric"))+", Items: []w.DocFile{"+join(fs, ", ")+"}}")
		case k.Has("tc-doc__sec--rule16"):
			l := k.Kid("tc-doc__link")
			items = append(items, "w.DocRelated{Title: "+g.rub(k.Kid("tc-rubric"))+", Name: "+q(l.OwnText())+", What: "+dimText(l.Kid("tc-dim").Text())+"}")
		case k.Tag == "textarea":
			items = append(items, g.area(k))
		case k.Has("tc-doc__fields"):
			items = append(items, "w.DocFields{Items: []w.Field{"+g.each(k.Els(), g.field)+"}}")
		case k.Has("tc-doc__note"):
			items = append(items, "w.DocNote{Parts: "+g.inl2(k)+"}")
		case k.Has("tc-doc__heading"):
			items = append(items, "w.DocHeading{Text: "+q(k.Text())+"}")
		case k.Has("tc-doc__bullets"):
			items = append(items, "w.DocBullets{Items: [][]w.Inline{"+g.each(k.Els(), g.inl2)+"}}")
		case k.Has("tc-doc__body"):
			items = append(items, "w.DocBody{Text: "+q(k.Text())+"}")
		case k.Has("tc-doc__quote"):
			items = append(items, "w.DocQuote{Parts: "+g.inl2(k)+"}")
		case k.Has("tc-doc__code"):
			w := ""
			if k.Has("is-wrap") {
				w = ", Wrap: true"
			}
			items = append(items, "w.DocCode{Text: "+q(k.TextRaw())+w+"}")
		case k.Has("tc-doc__table"):
			items = append(items, g.docTable(k))
		default:
			stop("doc? " + join(k.Cls(), " "))
		}
	}
	f := []string{"Items: []w.DocPart{\n\t\t" + join(items, ",\n\t\t") + "}"}
	if e.Has("tc-stack") {
		f = append(f, "Text: true")
	}
	if e.Tag == "form" {
		if id := e.Attr("id"); id != "" {
			f = append(f, "ID: "+q(id))
		}
		f = append(f, "Action: "+q(e.Attr("action")))
		if m := e.Attr("method"); m != "" && m != "post" {
			f = append(f, "Method: "+q(m))
		}
		if len(hidden) > 0 {
			f = append(f, "Fields: "+hiddens(hidden))
		}
	}
	return "w.Doc{" + join(f, ", ") + "}"
}

// docTable is a table in a document's text: its head, its rows, how each
// column is aligned, and whether its rows keep to one line.
func (g *gen) docTable(e *El) string {
	tbl := e.Kid("tc-doc__scroll").Els()[0]
	head := tbl.Els()[0].Els()[0].Els()
	var cells, align, rows []string
	aligned := false
	for _, th := range head {
		cells = append(cells, g.inl2(th))
		a := "w.DocAlignStart"
		if th.Has("is-center") {
			a, aligned = "w.DocAlignCenter", true
		} else if th.Has("is-end") {
			a, aligned = "w.DocAlignEnd", true
		}
		align = append(align, a)
	}
	for _, tr := range tbl.Els()[1].Els() {
		rows = append(rows, "{"+g.each(tr.Els(), g.inl2)+"}")
	}
	f := []string{"Head: [][]w.Inline{" + join(cells, ", ") + "}", "Rows: [][][]w.Inline{" + join(rows, ", ") + "}"}
	if aligned {
		f = append(f, "Align: []w.DocAlign{"+join(align, ", ")+"}")
	}
	f = append(f, "Switch: "+g.child(e.Kid("tc-doc__tablebar").Els()[0]))
	return "w.DocTable{" + join(f, ", ") + "}"
}

func (g *gen) docbar(e *El) string {
	var items, f, end []string
	for _, k := range e.Els() {
		switch {
		case k.Has("tc-bar__sep"):
			items = append(items, "w.BarSep{}")
		case k.Has("tc-bar__more"):
			f = append(f, "More: &"+g.button(k))
		case k.Has("tc-drop"):
			items = append(items, g.drop(k)) // Attach with its menu (Carrel's task 126)
		case k.Has("tc-bar__right"):
			for _, x := range k.Els() {
				sec := ""
				if x.Has("is-2nd") {
					sec = ", Secondary: true"
				}
				switch {
				case x.Has("tc-check"), x.Has("tc-seg"):
					end = append(end, g.barItem(x))
				case x.Has("tc-kbd"):
					end = append(end, g.kbd(x))
				case x.Has("tc-btn"):
					end = append(end, g.button(x))
				default:
					end = append(end, "w.DraftNote{Text: "+q(x.OwnText())+", Time: w.Mono{Text: "+q(x.Kid("tc-mono").Text())+"}"+sec+"}")
				}
			}
		case k.Has("tc-btn"):
			b := g.button(k)
			if k.Kid("b") != nil {
				b = `w.Button{Label: "B", Bold: true}`
			}
			if k.Has("tc-btn--fontstyle-italic") {
				b = `w.Button{Label: "I", Variant: w.ButtonItalic}`
			}
			items = append(items, b)
		default:
			stop("docbar? " + join(k.Cls(), " ")) // nothing is dropped unsaid (the review of 14.10)
		}
	}
	out := append([]string{"Items: []w.DocbarPart{" + join(items, ", ") + "}"}, f...)
	if len(end) > 0 {
		out = append(out, "End: []w.DocbarEnd{"+join(end, ", ")+"}")
	}
	if e.Has("tc-docbar--rule") {
		out = append(out, "Rule: true")
	}
	if e.Has("tc-docbar--edit") {
		out = append(out, "Edit: true")
	}
	return "w.Docbar{" + join(out, ", ") + "}"
}

// docAttach is an attachment in a document's column: its name (a link when
// it opens), size, where it lies, Remove on an edit.
func (g *gen) docAttach(x *El) string {
	name, href := strings.TrimSpace(x.OwnText()), ""
	if a := x.Kid("tc-dim--accent"); a != nil {
		name, href = a.Text(), ", Href: "+q(a.Attr("href"))
	}
	s := "w.DocAttach{Name: " + q(name) + ", Size: " + dimText(x.Kid("tc-dim").Text())
	if b := x.Kid("tc-badge"); b != nil {
		s += ", Where: " + g.badge(b)
	}
	if b := x.Kid("tc-btn"); b != nil {
		s += ", Remove: &" + g.button(b)
	}
	return s + href + "}"
}

// docText is a part of a document's text: a heading, a list, a quote, code,
// a table — on a page and in a panel alike.
func (g *gen) docText(k *El) string {
	switch {
	case k.Has("tc-doc__heading"):
		return "w.DocHeading{Text: " + q(k.Text()) + "}"
	case k.Has("tc-doc__bullets"):
		return "w.DocBullets{Items: [][]w.Inline{" + g.each(k.Els(), g.inl2) + "}}"
	case k.Has("tc-doc__quote"):
		return "w.DocQuote{Parts: " + g.inl2(k) + "}"
	case k.Has("tc-doc__code"):
		w := ""
		if k.Has("is-wrap") {
			w = ", Wrap: true"
		}
		return "w.DocCode{Text: " + q(k.TextRaw()) + w + "}"
	case k.Has("tc-doc__table"):
		return g.docTable(k)
	}
	return ""
}

// scripted is the fields of a widget shown only with a script or only
// without one (tc-js, tc-nojs).
func scripted(e *El) []string {
	var f []string
	if e.Has("tc-js") {
		f = append(f, "ScriptOnly: true")
	}
	if e.Has("tc-nojs") {
		f = append(f, "NoScript: true")
	}
	return f
}
