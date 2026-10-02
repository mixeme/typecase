package framegen

import "strings"

func (g *gen) subparts(d *El) string {
	var parts []string
	for _, k := range d.Kids {
		switch v := k.(type) {
		case string:
			if t := textpart(d, v); t != "" {
				parts = append(parts, t)
			}
		case *El:
			switch {
			case v.Tag == "b":
				parts = append(parts, "w.Strong("+q(v.Text())+")")
			case v.Has("tc-dim--danger"):
				parts = append(parts, "w.Span{Face: w.FaceDanger, Text: "+q(v.Text())+"}")
			case v.Has("tc-mono"):
				parts = append(parts, g.mono(v))
			case v.Has("tc-mark"):
				parts = append(parts, g.mark(v))
			case v.Has("tc-badge"):
				parts = append(parts, g.badge(v))
			case v.Has("tc-btn"):
				parts = append(parts, g.button(v))
			default:
				stop("sub? " + join(v.Cls(), " "))
			}
		}
	}
	return "w.HeadSub{Parts: []w.HeadSubPart{" + join(trimEnds(parts), ", ") + "}}"
}

func (g *gen) htitle(t *El) string {
	var f []string
	if m := t.Kid("tc-mark"); m != nil {
		f = append(f, "Mark: &"+g.mark(m))
	}
	f = append(f, "Text: "+q(t.OwnText()))
	if b := t.Kid("tc-badge"); b != nil {
		f = append(f, "Badge: &"+g.badge(b))
	}
	for _, v := range allOf(t, []kv{{"tc-head__title--t19", "HeadTitle19"}, {"tc-head__title--t18", "HeadTitle18"}, {"tc-head__title--t17", "HeadTitle17"}}) {
		f = append(f, "Size: w."+v)
	}
	return "&w.HeadTitle{" + join(f, ", ") + "}"
}

func photo(p *El) string {
	var f []string
	if t := p.Text(); t != "" {
		f = append(f, "Initials: "+q(t))
	}
	if p.Has("tc-avatar--photo--sq76-t20") {
		f = append(f, "Large: true")
	}
	return "&w.HeadPhoto{" + join(f, ", ") + "}"
}

func (g *gen) titlesInto(box *El, f []string) []string {
	var subs []string
	for _, k := range box.Els() {
		switch {
		case k.Has("tc-head__title"):
			f = append(f, "Title: "+g.htitle(k))
		case k.Has("tc-head__sub"):
			subs = append(subs, g.subparts(k))
		case k.Has("tc-crumbs"):
			f = append(f, "Crumbs: &"+g.crumbs(k))
		case k.Has("tc-rubric"):
			v := ""
			if k.Has("tc-rubric--bot2") {
				v = ", Variant: w.RubricBottom2"
			}
			f = append(f, "Kicker: &w.Rubric{Text: "+q(k.OwnText())+v+"}")
		case k.Has("tc-sec__acts"):
			f = append(f, "Tools: []w.HeadTool{"+g.buttons(k.Els())+"}")
		case k.Has("tc-head__text"):
			f = g.titlesInto(k, f)
		case k.Has("tc-avatar--photo"):
			f = append(f, "Photo: "+photo(k))
		case k.Has("tc-avatar"):
			f = append(f, "Avatar: &w.Avatar{Initials: "+q(k.Text())+", Size: w.Avatar44}")
		case k.Has("tc-head__titles"):
			f = g.titlesInto(k, f)
		default:
			stop("titles? " + join(k.Cls(), " "))
		}
	}
	if len(subs) > 0 {
		f = append(f, "Subs: []w.HeadSub{"+join(subs, ", ")+"}")
	}
	return f
}

func (g *gen) head(e *El) string {
	var f []string
	for _, k := range e.Els() {
		switch {
		case k.Has("tc-head__titles"):
			f = g.titlesInto(k, f)
		case k.Has("tc-head__group"):
			f = append(f, "Grouped: true")
			f = g.titlesInto(k, f)
		case k.Has("tc-head__sub"):
			f = append(f, "Meta: &"+g.subparts(k))
		case k.Has("tc-acts"):
			f = append(f, "Acts: &"+g.acts(k))
		default:
			stop("head? " + join(k.Cls(), " "))
		}
	}
	if e.Has("tc-head--mid") {
		f = append(f, "Middle: true")
	}
	for _, v := range allOf(e, []kv{{"tc-head--padt11", "HeadPadTop11"}, {"tc-head--padb6", "HeadPadBottom6"}, {"tc-head--snug", "HeadPadBottom8"},
		{"tc-head--padb10", "HeadPadBottom10"}, {"tc-head--rule", "HeadRule"}}) {
		f = append(f, "Pad: w."+v)
	}
	return "w.Head{" + join(f, ", ") + "}"
}

func (g *gen) barCount(k *El) string {
	s := "Count: &w.BarCount{Text: " + q(k.Text())
	if k.Has("tc-bar__count--t12-5") {
		s += ", Small: true"
	}
	return s + "}"
}

func (g *gen) barRight(k *El) string {
	s := "Right: &w.BarRight{Items: []w.BarPart{" + g.each(k.Els(), g.barItem) + "}"
	if k.Has("is-2nd") {
		s += ", Secondary: true"
	}
	return s + "}"
}

func (g *gen) bar(e *El) string {
	var f, items []string
	for _, k := range e.Els() {
		switch {
		case k.Has("tc-drop"):
			items = append(items, g.drop(k))
		case k.Has("tc-bar__count"):
			f = append(f, g.barCount(k))
		case k.Has("tc-bar__more"):
			f = append(f, "More: &"+g.button(k))
		case k.Has("tc-bar__right"):
			f = append(f, g.barRight(k))
		default:
			items = append(items, g.barItem(k))
		}
	}
	f = append([]string{"Items: []w.BarPart{" + join(items, ", ") + "}"}, f...)
	for _, v := range allOf(e, []kv{{"tc-bar--sel", "BarSelection"}, {"tc-bar--wide", "BarWide"}}) {
		f = append(f, "Variant: w."+v)
	}
	return "w.Bar{" + join(f, ", ") + "}"
}

func (g *gen) blockChild(k *El) string {
	switch {
	case k.Has("tc-form"):
		return g.form(k)
	case k.Has("tc-formfoot"):
		return g.formFoot(k)
	case k.Has("tc-dialog"):
		return g.dialog(k)
	case k.Has("tc-scrim"):
		return g.scrim(k)
	case k.Has("tc-head"):
		return g.head(k)
	case k.Has("tc-bar"):
		return g.bar(k)
	case k.Has("tc-list"):
		return g.list(k)
	case k.Has("tc-railsec"):
		return g.railsec(k)
	case k.Has("tc-zone"):
		return g.zone(k)
	}
	return g.anyblock(k)
}

var mainVariant = []kv{{"tc-main--zone", "MainZone"}, {"tc-main--padb180", "MainRoomy"}, {"tc-main--minh88", "MainShort"}}

func (g *gen) main(e *El) string {
	var items []string
	for _, k := range e.Els() {
		switch {
		case k.Has("tc-list__group"):
			items = append(items, g.listGroup(k))
		case k.Has("tc-main__tree"):
			items = append(items, "w.MainTree{Tree: "+g.tree(k.Els()[0])+"}")
		case k.Has("tc-field__hint--gutter"):
			items = append(items, "w.MainHint{Text: "+q(k.Text())+"}")
		case k.Has("tc-nav--bottom"):
			var its []string
			on := 0
			for i, a := range k.Els() {
				if a.Has("is-on") {
					on = i + 1
				}
				its = append(its, "{Icon: "+icon(svgs(a)[0])+", Label: "+q(a.OwnText())+`, Href: "#"}`)
			}
			items = append(items, "w.BottomNav{Items: []w.BottomItem{"+join(its, ", ")+"}, On: "+itoa(on)+"}")
		case k.Has("tc-main__banner") || k.Has("tc-main__banner--12"):
			f := []string{"Logo: " + logo(k.Kid("tc-brand__logo"))}
			if b := k.Kid("tc-badge"); b != nil {
				f = append(f, "Badge: &"+g.badge(b))
			}
			if end := k.Kid("tc-push"); end != nil {
				inner := g.each(end.Els(), func(x *El) string {
					if x.Has("tc-btn") {
						return g.button(x)
					}
					return g.child(x)
				})
				if end.Has("tc-line") {
					f = append(f, "End: w.Line{Gap: w.LineGap6, Place: w.Pushed, Items: []w.LinePart{"+inner+"}}")
				} else {
					f = append(f, "End: w.Push{Items: []w.PushPart{"+inner+"}}")
				}
			}
			if k.Has("tc-main__banner") {
				f = append(f, "Wide: true")
			}
			items = append(items, "w.Banner{"+join(f, ", ")+"}")
		case k.Has("tc-main__strip"):
			parts := g.each(k.Els()[1:], func(x *El) string {
				switch {
				case x.Has("tc-btn"):
					return g.button(x)
				case x.Has("tc-dim"):
					return g.dimx(x)
				}
				return g.input(x)
			})
			items = append(items, "w.Strip{Label: "+q(k.Els()[0].Text())+", Items: []w.StripPart{"+parts+"}}")
		case k.Has("tc-main__foot") || k.Has("tc-main__foot--16"):
			s := "w.MainFoot{Text: " + q(k.Text())
			if k.Has("tc-main__foot--16") {
				s += ", Roomy: true"
			}
			items = append(items, s+"}")
		case k.Has("tc-main__results"):
			items = append(items, "w.MainResults{Items: []w.ResultsPart{"+g.each(k.Els(), g.anyblock)+"}}")
		case k.Has("tc-main__cards"):
			items = append(items, "w.MainCards{Cards: []w.Card{"+g.each(k.Els(), g.card)+"}}")
		case k.Has("tc-main__diag"):
			items = append(items, "w.MainDiag{Diag: "+g.diag(k.Els()[0])+"}")
		case k.Has("tc-main__thumbs"):
			items = append(items, "w.MainThumbs{Thumb: "+g.thumb(k.Els()[0])+"}")
		case k.Has("tc-main__stage"):
			it := k.Kid("tc-stack__item")
			items = append(items, "w.Stage{Thumb: "+g.thumb(k.Kid("tc-thumb"))+", Title: "+q(it.Kid("tc-head__title").Text())+", Meta: w.Mono{Text: "+
				q(it.Kid("tc-head__sub").Text())+"}, Download: "+g.button(k.Kid("tc-btn"))+", Note: "+q(k.Kid("tc-main__stage-note").Text())+"}")
		case k.Has("tc-main__person"):
			ph := k.Kid("tc-avatar--photo")
			side := k.Kid("tc-main__person-side")
			sub := side.Kid("tc-head__sub")
			s := "w.Person{Initials: " + q(ph.Text()) + ", Name: " + q(side.Kid("tc-main__person-name").Text()) + ", Sub: " + q(sub.OwnText())
			if b := sub.Kid("tc-badge"); b != nil {
				s += ", Badge: &" + g.badge(b)
			}
			if nxt := e.Kid("tc-main__person-acts"); nxt != nil {
				s += ", Acts: []w.Button{" + g.buttons(nxt.Els()) + "}"
			}
			items = append(items, s+"}")
		case k.Has("tc-main__person-acts"):
		case k.Has("tc-main__hero"):
			side := k.Kid("tc-main__hero-side")
			s := `w.Hero{Symbol: "davloom-tile", Label: "davloom", Title: ` + q(side.Kid("tc-main__hero-title").Text()) + ", Sub: " + q(side.Kid("tc-head__sub").Text()) +
				", Text: " + g.prose(side.Kid("tc-prose"))
			if fl := side.Kid("tc-fields"); fl != nil {
				s += ", Fields: &" + g.fields(fl)
			}
			items = append(items, s+", Acts: []w.Button{"+g.buttons(side.Kid("tc-doc__acts").Els())+"}}")
		case k.Has("tc-doc__layout"):
			view := ""
			if k.Has("is-measure") {
				view = ", View: w.DocViewMeasure"
			} else if k.Has("is-focus") {
				view = ", View: w.DocViewFocus"
			}
			its := g.each(k.Els(), func(y *El) string {
				switch {
				case y.Has("tc-docmeta"):
					return g.docmeta(y)
				case y.Has("tc-docbar"):
					return g.docbar(y)
				}
				return g.doc(y)
			})
			items = append(items, "w.DocLayout{Items: []w.DocLayoutPart{"+its+"}"+view+"}")
		case k.Has("tc-main__foot-acts"):
			items = append(items, "w.FootActs{Items: []w.Button{"+g.buttons(k.Els())+"}}")
		default:
			items = append(items, g.blockChild(k))
		}
	}
	f := []string{"Items: []w.MainPart{\n\t\t" + join(items, ",\n\t\t") + "}"}
	for _, v := range allOf(e, mainVariant) {
		f = append(f, "Variant: w."+v)
	}
	return "w.Main{" + join(f, ", ") + "}"
}

func (g *gen) side(e *El) string {
	var items []string
	for _, k := range e.Els() {
		if k.Has("tc-side__acts") {
			items = append(items, "w.SideActs{Items: []w.Button{"+g.buttons(k.Els())+"}}")
		} else {
			items = append(items, g.blockChild(k))
		}
	}
	f := []string{"Items: []w.SidePart{" + join(items, ", ") + "}"}
	for _, v := range allOf(e, []kv{{"tc-side--zone", "SideZone"}, {"tc-side--drop", "SideDrop"}}) {
		f = append(f, "Variant: w."+v)
	}
	return "w.Side{" + join(f, ", ") + "}"
}

var shellHeight = []kv{{"tc-shell__body--minh230", "ShellHeight230"}, {"tc-shell__body--minh280", "ShellHeight280"}, {"tc-shell__body--minh340", "ShellHeight340"},
	{"tc-shell__body--minh390", "ShellHeight390"}, {"tc-shell__body--minh560", "ShellHeight560"}}

func (g *gen) body(b *El) string {
	var f []string
	if rail := b.Kid("tc-shell__rail"); rail != nil {
		f = append(f, "Rail: []w.RailPart{"+g.each(rail.Els(), g.childRail)+"}")
		if rail.Has("tc-shell__rail--icons") {
			f = append(f, "RailForm: w.RailIcons")
		}
		if rail.Has("tc-shell__rail--zone") {
			f = append(f, "RailForm: w.RailZone")
		}
	}
	f = append(f, "Main: "+g.main(b.Kid("tc-main")))
	if s := b.Kid("tc-side"); s != nil {
		f = append(f, "Side: &"+g.side(s))
	}
	for _, v := range allOf(b, shellHeight) {
		f = append(f, "Height: w."+v)
	}
	return "w.ShellBody{" + join(f, ", ") + "}"
}

func (g *gen) trail(k *El) string {
	switch {
	case k.Has("tc-mono"):
		return g.mono(k)
	case k.Has("tc-dim"):
		return g.dimx(k)
	case k.Has("tc-badge"):
		b := g.badge(k)
		if k.Has("tc-push") {
			return b[:len(b)-1] + ", Place: w.Pushed}"
		}
		return b
	case k.Has("tc-kbd"):
		s := "w.Kbd{Keys: " + q(k.Text())
		if k.Has("tc-push") {
			s += ", Place: w.Pushed"
		}
		return s + "}"
	case k.Has("tc-push"):
		return "w.Push{Items: []w.PushPart{w.Text(" + q(k.Text()) + ")}}"
	}
	stop("trail? " + join(k.Cls(), " "))
	return ""
}

var menuTone = []kv{{"tc-menu__item--danger", "MenuToneDanger"}, {"tc-menu__item--dusty", "MenuToneDusty"}, {"tc-menu__item--muted", "MenuToneMuted"},
	{"tc-menu__item--accent", "MenuToneAccent"}, {"tc-menu__item--note", "MenuToneAside"}}

func (g *gen) menuItem(k *El) string {
	var f []string
	found := false
	box := k.Kid("tc-check__box")
	radio := k.Kid("tc-seg__box")
	for _, x := range k.Els() {
		switch {
		case x == box || x == radio:
			continue
		case box != nil && x.Has("tc-tickrow__tick"):
			t := ""
			if !box.HasAttr("checked") {
				t = "Ticked: w.TickedOff"
			}
			f = append(f, "Tick: &w.TickBox{"+t+"}")
			continue
		case x.Tag == "svg":
			f = append(f, "Icon: "+icon(x))
			if x.Has("tc-ico--op-35") {
				f = append(f, "Faint: true")
			}
		case x.Has("tc-mark"):
			f = append(f, "Mark: &"+g.mark(x))
		case x.Has("tc-tickrow__tick"):
			t := ""
			if x.Has("is-off") {
				t = "Ticked: w.TickedOff"
			}
			f = append(f, "Tick: &w.TickBox{"+t+"}")
		case x.Has("tc-menu__label"):
			f = append(f, "Found: "+g.inl(x))
			found = true
		case x.Tag == "template":
			f = append(f, "Adds: &"+g.tag(x.Els()[0]))
		case x.Has("tc-dim--danger"):
			f = append(f, "Note: &w.Span{Face: w.FaceDanger, Text: "+q(x.Text())+"}")
		case x.Has("tc-mono") && !x.Has("tc-push"):
			f = append(f, "Detail: &"+g.mono(x))
		case x.Has("tc-tag"):
			f = append(f, "Trail: w.Tag{Label: "+q(x.Text())+", Place: w.Pushed}")
		default:
			f = append(f, "Trail: "+g.trail(x))
		}
	}
	if !found {
		f = append(f, "Label: "+q(k.OwnText()))
	}
	if k.Has("is-on") {
		f = append(f, "On: true")
	}
	for _, v := range allOf(k, menuTone) {
		f = append(f, "Tone: w."+v)
	}
	if k.Has("is-3rd") {
		f = append(f, "Tertiary: true")
	}
	if r := k.Attr("title"); r != "" {
		f = append(f, "Reason: "+q(r))
	}
	if h := k.Attr("href"); h != "" {
		f = append(f, "Href: "+q(h))
	}
	if fm := k.Attr("form"); fm != "" && box == nil {
		f = append(f, "Form: "+q(fm))
		if v := k.Attr("name"); v != "" {
			f = append(f, "Name: "+q(v), "Value: "+q(k.Attr("value")))
		}
	}
	if radio != nil {
		f = append(f, "Name: "+q(radio.Attr("name")), "Value: "+q(radio.Attr("value")))
	}
	if box != nil {
		f = append(f, "Name: "+q(box.Attr("name")), "Value: "+q(box.Attr("value")))
		if box.HasAttr("disabled") {
			f = append(f, "Fixed: true")
		}
		if k.Has("is-drag") {
			f = append(f, "Drag: true")
		}
	}
	return "w.MenuItem{" + join(f, ", ") + "}"
}

var menuWidth = []kv{{"tc-menu--w212", "MenuWidth212"}, {"tc-menu--w290", "MenuWidth290"}, {"tc-menu--w420", "MenuWidth420"}, {"tc-menu--full", "MenuWidthFull"}}

func (g *gen) menu(e *El) string {
	var items []string
	var hidden []*El
	for _, k := range e.Els() {
		switch {
		case k.Tag == "input" && k.Attr("type") == "hidden":
			hidden = append(hidden, k)
		case k.Has("tc-menu__sep"):
			t := ""
			if k.Has("is-3rd") {
				t = "Tertiary: true"
			}
			items = append(items, "w.MenuSep{"+t+"}")
		case k.Has("tc-rubric") && k.Has("tc-menu__item"):
			form := ""
			if k.Has("tc-menu__item--pad5-8-7") {
				form = ", Form: w.MenuHeadingRoomy"
			} else if k.Has("tc-menu__item--split") {
				form = ", Form: w.MenuHeadingSplit"
			}
			items = append(items, "w.MenuHeading{Rubric: "+g.rubric(k, "RubricItem")+form+"}")
		case k.Has("tc-rubric"):
			form := ""
			if k.Tag == "span" {
				form = "RubricInline"
			}
			items = append(items, g.rubric(k, form))
		case k.Has("tc-menu__item--col"):
			seg := k.Kid("tc-seg")
			opts := g.each(seg.Els(), func(o *El) string {
				s := "{Label: " + q(o.Text())
				if o.Has("is-on") {
					s += ", On: true"
				}
				return s + "}"
			})
			items = append(items, "w.MenuColumn{Label: "+q(k.Els()[0].Text())+", Control: w.Seg{Options: []w.SegOption{"+opts+"}, Variant: w.SegStartSmall}}")
		case k.Tag == "div" && k.Has("tc-menu__item"):
			f := []string{"Label: " + q(k.OwnText())}
			if b := k.Kid("tc-badge"); b != nil {
				f = append(f, "Badge: &"+g.badge(b))
			}
			f = append(f, "Default: "+g.trail(k.Kid("tc-dim")))
			i := k.Kid("tc-in")
			f = append(f, "Input: w.Input{Value: "+q(i.Attr("value"))+", Mono: true, Variant: w.InputChip, Label: "+q(i.Attr("aria-label"))+"}")
			items = append(items, "w.MenuField{"+join(f, ", ")+"}")
		case k.Has("tc-menu__refusal"):
			r := k.Kid("tc-menu__what")
			items = append(items, "w.MenuRefusal{Mark: "+g.mark(k.Kid("tc-mark"))+", Source: "+q(r.Kid("tc-menu__source").Text())+
				", Answer: "+q(r.Kid("tc-menu__answer").Text())+", Retry: "+g.button(k.Kid("tc-btn"))+"}")
		case k.Has("tc-menu__foot"):
			items = append(items, "w.MenuFoot{Items: []w.Button{"+g.buttons(k.Els())+"}}")
		case k.Has("tc-menu__item"):
			items = append(items, g.menuItem(k))
		case k.Has("tc-in"):
			items = append(items, g.input(k))
		case k.Has("tc-range"):
			b := g.barItem(k)
			items = append(items, b[:len(b)-1]+", Wrap: true}")
		case k.Has("tc-checkform"):
			items = append(items, g.checkForm(k))
		default:
			stop("menu? " + join(k.Cls(), " "))
		}
	}
	f := []string{"Items: []w.MenuPart{\n\t\t" + join(items, ",\n\t\t") + "}"}
	for _, v := range allOf(e, menuWidth) {
		f = append(f, "Width: w."+v)
	}
	if e.Has("is-pop") {
		if e.Has("is-right") {
			f = append(f, "Pop: w.MenuUnderEnd")
		} else {
			f = append(f, "Pop: w.MenuUnder")
		}
	}
	if e.Has("tc-menu--fields") {
		f = append(f, "Fields: true")
	}
	if e.Tag == "form" {
		f = append(f, "Action: "+q(e.Attr("action")))
		if len(hidden) > 0 {
			f = append(f, "Hidden: "+hiddens(hidden))
		}
	}
	if e.Tag == "span" {
		f = append(f, "Inline: true")
	}
	return "w.Menu{" + join(f, ", ") + "}"
}

// top2 is a top bar with a dropdown open or the search's suggestions.
func (g *gen) top2(e *El) string {
	f := []string{"Brand: " + g.brand(e.Kid("tc-brand"))}
	if line := e.Kid("tc-top__search"); line != nil {
		f = append(f, "Search: &"+g.search(line.Kid("tc-search")))
		f = append(f, "Suggest: &"+g.menu(line.Kid("tc-top__suggest").Kid("tc-menu")))
	} else {
		f = append(f, "Search: &"+g.search(e.Kid("tc-search")))
	}
	var acts []string
	for _, k := range e.Kid("tc-top__actions").Els() {
		if k.Has("tc-drop__anchor") {
			panel := k.Els()[1]
			s := "w.TopDrop{Open: true, Trigger: " + g.child(k.Els()[0]) + ", Menu: " + g.menu(panel.Els()[0])
			if panel.Has("tc-drop__panel--right") {
				s += ", End: true"
			}
			acts = append(acts, s+"}")
		} else {
			acts = append(acts, g.child(k))
		}
	}
	f = append(f, "Actions: []w.TopAction{"+join(acts, ", ")+"}")
	return "w.Top{" + join(f, ", ") + "}"
}

// top is a plain top bar; "" when it has something open (top2).
func (g *gen) top(e *El) string {
	var f []string
	if e.Els()[0].Has("tc-btn") {
		f = append(f, "Lead: &"+g.button(e.Els()[0]))
	}
	f = append(f, "Brand: "+g.brand(e.Kid("tc-brand")))
	if s := e.Kid("tc-search"); s != nil {
		f = append(f, "Search: &"+g.search(s))
	}
	a := e.Kid("tc-top__actions")
	if a.Kid("tc-drop__anchor") != nil || e.Kid("tc-top__search") != nil {
		return ""
	}
	f = append(f, "Actions: []w.TopAction{"+g.each(a.Els(), g.child)+"}")
	if a.Has("tc-top__actions--gap2") {
		f = append(f, "Tight: true")
	}
	return "w.Top{" + join(f, ", ") + "}"
}

func (g *gen) topAny(t *El) string {
	if s := g.top(t); s != "" {
		return s
	}
	return g.top2(t)
}

func (g *gen) tickrow(e *El) string {
	var f []string
	els := e.Els()
	tick := els[0]
	if t := tickedOf(tick); t != "w.TickedOn" {
		f = append(f, "Ticked: "+t)
	}
	if s := sendOf(tick); s != "" {
		f = append(f, "Send: "+s)
	}
	for _, v := range allOf(e, []kv{{"is-all", "TickRowAll"}, {"is-root", "TickRowRoot"}, {"is-ext", "TickRowExternal"}}) {
		f = append(f, "Kind: w."+v)
	}
	for _, v := range allOf(e, []kv{{"tc-tickrow--padl0", "TickRowGroup"}, {"tc-tickrow--padl14", "TickRowNested"}, {"tc-tickrow--h44", "TickRowTouch"}}) {
		f = append(f, "Form: w."+v)
	}
	if lead := els[1]; lead.Tag == "svg" {
		f = append(f, "Icon: "+icon(lead))
	} else {
		f = append(f, "Mark: "+g.mark(lead))
	}
	f = append(f, "Label: "+q(e.Kid("tc-tickrow__label").Text()))
	if meta := e.Kid("tc-tickrow__meta"); meta != nil {
		f = append(f, "Meta: "+q(meta.Text()))
		if t := meta.Attr("title"); t != "" {
			f = append(f, "MetaTitle: "+q(t))
		}
	}
	if e.Has("is-on") {
		f = append(f, "Current: true")
	}
	if e.Has("is-off") {
		f = append(f, "Hidden: true")
	}
	if e.Has("is-error") {
		f = append(f, "Error: true")
	}
	last := els[len(els)-1]
	if last.Has("tc-badge") {
		f = append(f, "Trail: "+strings.Replace(g.badge(last), ", Place: w.Pushed", "", -1))
	}
	if last.Has("tc-btn") {
		f = append(f, "Trail: w.Button{Label: "+q(last.Text())+", Small: true}")
	}
	if last.Has("tc-skel") {
		f = append(f, "Trail: w.Skel{Width: w.Skel16}")
	}
	return "w.TickRow{" + join(f, ", ") + "}"
}

var treeIndent = []kv{{"tc-tree__row--padl1-25rem", "TreeLevel20"}, {"tc-tree__row--padl22", "TreeLevel22"}, {"tc-tree__row--padl36", "TreeLevel36"},
	{"tc-tree__row--padl40", "TreeLevel40"}, {"tc-tree__row--padl44", "TreeLevel44"}}

func (g *gen) tree(e *El) string {
	var f, rows []string
	on := 0
	for _, v := range allOf(e, []kv{{"tc-tree--padx8", "TreePadded"}, {"tc-tree--max520", "TreeNarrow"}}) {
		f = append(f, "Variant: w."+v)
	}
	for i, r := range e.Els() {
		var x []string
		if r.Has("is-on") {
			on = i + 1
		}
		els := r.Els()
		if lead := els[0]; lead.Tag == "svg" {
			x = append(x, "Icon: "+icon(lead))
		} else if lead.Has("tc-mark") {
			x = append(x, "Mark: &"+g.mark(lead))
		}
		var mono, tr *El
		for _, k := range els {
			if mono == nil && k.Has("tc-mono") && !k.Has("tc-push") {
				mono = k
			}
			if tr == nil && k.Has("tc-push") {
				tr = k
			}
		}
		if mono != nil {
			x = append(x, "Label: "+q(mono.Text()), "Code: true")
		} else {
			x = append(x, "Label: "+q(r.OwnText()))
		}
		for _, v := range allOf(r, treeIndent) {
			x = append(x, "Indent: w."+v)
		}
		if r.Has("tc-tree__row--op-55") {
			x = append(x, "Unavailable: true")
		}
		if h := r.Attr("href"); h != "" {
			x = append(x, "Href: "+q(h))
		}
		if t := tr; t != nil {
			switch {
			case t.Has("tc-tickrow__meta"):
				x = append(x, "Trail: w.TreeNote{Text: "+q(t.Text())+"}")
			case t.Has("tc-dim"):
				x = append(x, "Trail: w.Dim{Parts: []w.Inline{w.Text("+q(t.Text())+")}, Size: w.Dim11}")
			case t.Has("tc-mono"):
				x = append(x, "Trail: w.Mono{Text: "+q(t.Text())+", Variant: w.Mono10Dusty}")
			case t.Has("tc-badge"):
				x = append(x, "Trail: "+strings.Replace(g.badge(t), ", Place: w.Pushed", "", -1))
			}
		}
		rows = append(rows, "{"+join(x, ", ")+"}")
	}
	out := []string{"Rows: []w.TreeRow{\n\t\t" + join(rows, ",\n\t\t") + "}"}
	if on > 0 {
		out = append(out, "On: "+itoa(on))
	}
	return "w.Tree{" + join(append(out, f...), ", ") + "}"
}

func (g *gen) titleFields(t *El, name string) []string {
	f := []string{name + ": " + g.inl(t)}
	if t.Has("is-strong") {
		f = append(f, "Strong: true")
	}
	if b := t.Kid("tc-badge"); b != nil {
		f = append(f, name+"Badge: &"+g.badge(b))
	}
	return f
}

func (g *gen) src(e *El, f []string) []string {
	s := e.Kid("tc-row__source")
	if s == nil {
		return f
	}
	if t := s.OwnText(); t != "" {
		f = append(f, "Source: "+q(t))
	}
	if t := s.Attr("title"); t != "" {
		f = append(f, "SourceHint: "+q(t))
	}
	if b := s.Kid("tc-badge"); b != nil {
		f = append(f, "SourceBadge: &"+g.badge(b))
	}
	return f
}

func (g *gen) tagsf(e *El, f []string) []string {
	items := g.each(e.Kid("tc-row__tags").Els(), func(k *El) string {
		if k.Has("tc-tag") {
			return g.tag(k)
		}
		return g.badge(k)
	})
	if items != "" {
		f = append(f, "Tags: []w.RowTag{"+items+"}")
	}
	return f
}

// row is a row of a list; a row that links to its record has its address.
func (g *gen) row(e *El) string {
	s := g.row0(e)
	if h := e.Attr("href"); h != "" {
		s = s[:len(s)-1] + ", Href: " + q(h) + "}"
	}
	return s
}

func (g *gen) row0(e *El) string {
	var f []string
	on := e.Has("is-on")
	els := e.Els()
	text := func(c string) string { return q(e.Kid(c).Text()) }
	switch {
	case e.Has("tc-row--cols"):
		return g.columnRow(e)
	case e.Has("tc-row--timecols") && e.Kid("tc-skel") != nil:
		r := map[string]string{"tc-row__detail--w40pc": "60", "tc-row__detail--w30pc": "45", "tc-row__detail--w25pc": "52"}
		for _, c := range e.Kid("tc-row__detail").Cls() {
			if w, ok := r[c]; ok {
				return "w.SkelRow{Title: w.Skel{Width: w.Skel" + w + "}}"
			}
		}
		stop("skel row")
	case e.Has("tc-row--timecols"):
		f = append(f, "Mark: "+g.mark(els[0]), "Time: "+text("tc-row__lead"))
		f = append(f, g.titleFields(e.Kid("tc-row__title"), "Title")...)
		f = append(f, "Detail: "+text("tc-row__detail"))
		if c := e.Kid("tc-row__caption"); len(c.Els()) > 0 {
			f = append(f, "Caption: &"+g.badge(c.Els()[0]))
		}
		f = g.src(e, f)
		if on {
			f = append(f, "On: true")
		}
		if e.Has("tc-row--timecols--op-7") {
			f = append(f, "Past: true")
		}
		if e.Has("tc-row--timecols--padl14-padr14") {
			f = append(f, "Inset: true")
		}
		if e.Has("tc-row--timecols--time") {
			f = append(f, "Short: true")
		}
		return "w.TimeRow{" + join(f, ", ") + "}"
	case e.Has("tc-row--avatarcols"):
		f = append(f, "Mark: "+g.mark(els[0]))
		f = append(f, "Avatar: "+g.avatar(e.Kid("tc-avatar")))
		f = append(f, g.titleFields(e.Kid("tc-row__title"), "Title")...)
		f = append(f, "Subtitle: "+text("tc-row__subtitle"), "Detail: "+text("tc-row__detail"), "Caption: "+text("tc-row__caption"))
		f = g.src(e, f)
		if on {
			f = append(f, "On: true")
		}
		if e.Has("tc-row--avatarcols--padl14-padr14") {
			f = append(f, "Inset: true")
		}
		return "w.FaceRow{" + join(f, ", ") + "}"
	case e.Has("tc-row--find"):
		lead := e.Kid("tc-row__lead")
		f = append(f, "Mark: "+g.mark(els[0]))
		if lead.Tag == "svg" {
			f = append(f, "Icon: "+icon(lead), "Title: "+text("tc-row__title"), "Detail: "+text("tc-row__detail"))
			return "w.LinkedRow{" + join(f, ", ") + "}"
		}
		f = append(f, "Date: "+q(lead.Text()), "Title: "+g.inl(e.Kid("tc-row__title")), "Detail: "+g.inl(e.Kid("tc-row__detail")), "What: "+text("tc-row__caption"))
		f = g.src(e, f)
		return "w.FindRow{" + join(f, ", ") + "}"
	case e.Has("tc-row--excerpt--cols3-62-148"):
		f = append(f, "Mark: "+g.mark(els[0]), "Title: "+text("tc-row__title"), "Date: w.Mono{Text: "+text("tc-mono")+"}", "Where: "+text("tc-row__tags"))
		if on {
			f = append(f, "On: true")
		}
		return "w.BriefRow{" + join(f, ", ") + "}"
	case e.Has("tc-row--excerpt"):
		f = append(f, "Mark: "+g.mark(els[0]))
		f = append(f, g.titleFields(e.Kid("tc-row__title"), "Title")...)
		f = append(f, "Excerpt: "+text("tc-row__excerpt"))
		f = g.tagsf(e, f)
		f = append(f, "Date: "+text("tc-row__date"))
		f = g.src(e, f)
		if on {
			f = append(f, "On: true")
		}
		if e.Has("tc-row--excerpt--padl14-padr14") {
			f = append(f, "Inset: true")
		}
		return "w.ExcerptRow{" + join(f, ", ") + "}"
	case e.Has("tc-row--tickmarkcols"):
		d := ""
		if els[0].Has("is-done") {
			d = "Done: true, "
		}
		f = append(f, "Tick: w.Tick{"+d+`Label: "Done"}`, "Mark: "+g.mark(els[1]))
		f = append(f, g.titleFields(e.Kid("tc-row__title"), "Title")...)
		f = g.tagsf(e, f)
		f = append(f, "Due: "+text("tc-row__due"), "Progress: "+text("tc-row__progress"))
		f = g.src(e, f)
		if on {
			f = append(f, "On: true")
		}
		if e.Has("is-overdue") {
			f = append(f, "Overdue: true")
		}
		if e.Has("tc-row--tickmarkcols--padl14-padr14") {
			f = append(f, "Inset: true")
		}
		return "w.DueRow{" + join(f, ", ") + "}"
	case e.Has("tc-row--tickcols"):
		if els[0].Has("is-off") {
			f = append(f, "Ticked: w.TickedOff")
		}
		name := e.Kid("tc-row__name")
		f = append(f, "Icon: "+icon(name.Els()[0]), "Name: "+g.inl(name.Kid("tc-row__file")), "Bytes: "+text("tc-row__size"),
			"Time: "+text("tc-row__time"), "Access: "+text("tc-row__detail"))
		f = g.src(e, f)
		return "w.SizeRow{" + join(f, ", ") + "}"
	case e.Has("tc-row--slim"):
		if els[0].Has("is-off") {
			f = append(f, "Ticked: w.TickedOff")
		}
		t := e.Kid("tc-row__title")
		f = append(f, "Mark: "+g.mark(els[1]), "Title: "+q(t.OwnText()), "Inside: "+q(t.Kid("tc-row__cell").Text()), "Aside: "+q(els[len(els)-1].Text()))
		return "w.SlimRow{" + join(f, ", ") + "}"
	case e.Has("tc-row--phone"):
		first := els[0]
		if first.Has("tc-tickrow__tick") {
			t := ""
			if first.Has("is-off") {
				t = "Ticked: w.TickedOff"
			}
			f = append(f, "Pick: &w.TickBox{"+t+"}")
		} else {
			f = append(f, "Mark: &"+g.mark(first))
		}
		outer := els[1]
		txt := outer
		if outer.Has("tc-row__text--row") {
			lead := outer.Els()[0]
			if lead.Tag == "svg" {
				f = append(f, "Icon: "+icon(lead))
			} else if lead.Has("is-photo") {
				f = append(f, "Avatar: &w.Avatar{Photo: true, Size: w.Avatar34}")
			} else {
				f = append(f, "Avatar: &w.Avatar{Initials: "+q(lead.Text())+", Size: w.Avatar34}")
			}
			txt = outer.Els()[1]
		}
		t := txt.Kid("tc-row__title")
		f = append(f, "Title: "+q(t.OwnText()))
		if b := t.Kid("tc-badge"); b != nil {
			f = append(f, "TitleBadge: &"+g.badge(b))
		}
		sub := txt.Kid("tc-row__cell")
		f = append(f, "Sub: "+g.inl(sub))
		if sub.Has("tc-mono") {
			f = append(f, "SubMono: true")
		}
		if last := els[len(els)-1]; last.Has("tc-btn") {
			f = append(f, `More: &w.Button{Icon: "dots", IconBox: w.IconBox16, Variant: w.ButtonPadX4, Title: "More"}`)
		} else {
			f = append(f, "Aside: "+q(last.Text()))
		}
		if on {
			f = append(f, "On: true")
		}
		return "w.PhoneRow{" + join(f, ", ") + "}"
	}
	stop("row? " + join(e.Cls(), " "))
	return ""
}

func (g *gen) list(e *El) string {
	f := []string{"Rows: []w.ListRow{\n\t\t" + strings.Join(eachList(e.Els(), g.row), ",\n\t\t") + "}"}
	if e.Has("tc-list--mx-14") {
		f = append(f, "Inset: true")
	}
	if e.Has("is-compact") {
		f = append(f, "Compact: true")
	}
	return "w.List{" + join(f, ", ") + "}"
}

func eachList(es []*El, fn func(*El) string) []string {
	var out []string
	for _, e := range es {
		out = append(out, fn(e))
	}
	return out
}

func (g *gen) childRail(k *El) string {
	switch {
	case k.Has("tc-railfoot"):
		return g.railfoot(k)
	case k.Has("tc-railsec"):
		return g.railsec(k)
	case k.Has("tc-nav"):
		return g.nav(k)
	case k.Has("tc-zone"):
		return g.zone(k)
	}
	return g.anyblock(k)
}

var sheetWidth = []kv{{"tc-sheet--w400", "SheetWidth400"}, {"tc-sheet--w480", "SheetWidth480"}, {"tc-sheet--w520", "SheetWidth520"},
	{"tc-sheet--w680", "SheetWidth680"}, {"tc-sheet--max100pc-full", "SheetWidthFull"}}

func (g *gen) sheet(e *El) string {
	var f []string
	head := e.Kid("tc-sheet__head")
	if s := svgs(head); len(s) > 0 {
		f = append(f, "Icon: "+icon(s[0]))
		if s[0].Has("tc-ico--danger") {
			f = append(f, "Tone: w.SheetIconDanger")
		} else if !s[0].Has("tc-ico--accent") {
			f = append(f, "Tone: w.SheetIconInk")
		}
	}
	f = append(f, "Title: "+q(head.Kid("tc-sheet__title").Text()))
	var hp []string
	for _, x := range head.Els() {
		switch {
		case x.Tag == "svg" || x.Has("tc-sheet__title"):
		case x.Has("tc-dim"):
			hp = append(hp, g.dimx(x))
		case x.Has("tc-close"):
			p := ""
			if x.Has("tc-push") {
				p = ", Place: w.Pushed"
			}
			hp = append(hp, `w.Close{Label: "Close"`+p+"}")
		case x.Has("tc-badge"):
			hp = append(hp, g.badge(x))
		case x.Has("tc-tag"):
			hp = append(hp, g.tag(x))
		case x.Has("tc-btn"):
			hp = append(hp, g.button(x))
		case x.Has("tc-in") && x.Has("is-pick"):
			hp = append(hp, g.pick(x))
		case x.Has("tc-in"):
			hp = append(hp, g.filter(x))
		case x.Has("tc-crumbs"):
			hp = append(hp, g.crumbs(x))
		case x.Has("tc-table__file"):
			hp = append(hp, "w.SheetSubject{Mark: "+g.mark(x.Kid("tc-mark"))+", Text: "+q(x.Els()[1].Text())+"}")
		default:
			stop("sheet head? " + join(x.Cls(), " "))
		}
	}
	if len(hp) > 0 {
		f = append(f, "Head: []w.SheetHeadPart{"+join(hp, ", ")+"}")
	}
	body := e.Kid("tc-sheet__body")
	var bp []string
	for _, x := range body.Els() {
		switch {
		case x.Has("tc-rubric"):
			bp = append(bp, g.rubric(x, ""))
		case x.Has("tc-tree"):
			bp = append(bp, g.tree(x))
		case x.Has("tc-railsec"):
			bp = append(bp, g.railsec(x))
		case x.Has("tc-sheet__fields"):
			bp = append(bp, "w.SheetFields{Items: []w.Field{"+g.each(x.Els(), g.field)+"}}")
		case x.Has("tc-field"):
			bp = append(bp, g.field(x))
		case x.Tag == "textarea":
			bp = append(bp, g.area(x))
		case x.Has("tc-sheet__keys"):
			ks := x.Els()
			var pairs []string
			for i := 0; i+1 < len(ks); i += 2 {
				pairs = append(pairs, "{Keys: "+q(ks[i].Text())+", Does: "+q(ks[i+1].Text())+"}")
			}
			bp = append(bp, "w.SheetKeys{Items: []w.SheetKey{"+join(pairs, ", ")+"}}")
		case x.Has("tc-field__hint--top6"):
			bp = append(bp, "w.SheetHint{Text: "+q(x.Text())+"}")
		case x.Has("tc-dim--lh1-55"):
			bp = append(bp, "w.SheetNote{Text: "+q(x.Text())+"}")
		case x.Has("tc-dim"):
			bp = append(bp, g.dimx(x))
		case x.Has("tc-check"):
			bp = append(bp, g.check(x))
		case x.Has("tc-sheet__checks"):
			bp = append(bp, "w.SheetChecks{Items: []w.Check{"+g.each(x.Els(), g.barItem)+"}}")
		case x.Has("tc-sheet__from"):
			item := x.Kid("tc-line__item")
			var lines [][]any
			var cur []any
			for _, y := range item.Kids {
				if el, ok := y.(*El); ok && el.Tag == "br" {
					lines = append(lines, cur)
					cur = nil
				} else {
					cur = append(cur, y)
				}
			}
			lines = append(lines, cur)
			var outl []string
			for _, ln := range lines {
				outl = append(outl, g.tparts(&El{Tag: "x", Kids: ln}, "[]w.Inline{"))
			}
			bp = append(bp, "w.SheetFrom{Label: "+q(x.Kid("tc-dim").Text())+", Lines: [][]w.Inline{"+join(outl, ", ")+"}}")
		case x.Has("tc-line") && !x.Has("tc-sheet__pick"):
			bp = append(bp, g.line(x))
		case x.Has("tc-sheet__pick"):
			rail := x.Kid("tc-sheet__pick-rail")
			rp := g.each(rail.Els(), func(y *El) string {
				if y.Has("tc-railsec") {
					return g.railsec(y)
				}
				return g.tree(y)
			})
			lp := g.each(x.Kid("tc-sheet__pick-list").Els(), func(y *El) string {
				switch {
				case y.Has("tc-list__group"):
					return g.listGroup(y)
				case y.Has("tc-list"):
					return g.list(y)
				case y.Has("tc-crumbs"):
					return g.crumbs(y)
				}
				return g.anyblock(y)
			})
			s := "w.SheetPick{Rail: []w.SheetPickRail{" + rp + "}, List: []w.SheetPickList{" + lp + "}"
			if x.Has("tc-sheet__pick--196") {
				s += ", Short: true"
			}
			bp = append(bp, s+"}")
		default:
			bp = append(bp, g.anyblock(x))
		}
	}
	f = append(f, "Body: []w.SheetPart{\n\t\t"+join(bp, ",\n\t\t")+"}")
	if body.Has("tc-sheet__body--pad0") {
		f = append(f, "Flush: true")
	}
	if foot := e.Kid("tc-sheet__foot"); foot != nil {
		var fp []string
		for _, x := range foot.Els() {
			switch {
			case x.Has("tc-btn"):
				fp = append(fp, g.button(x))
			case x.Has("tc-kbd"):
				p := ""
				if x.Has("tc-push") {
					p = ", Place: w.Pushed"
				}
				fp = append(fp, "w.Kbd{Keys: "+q(x.Text())+p+"}")
			case x.Has("tc-line"):
				fp = append(fp, "w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{"+g.each(x.Els(), func(y *El) string {
					if y.Has("tc-btn") {
						return g.button(y)
					}
					return "w.Kbd{Keys: " + q(y.Text()) + "}"
				})+"}}")
			case x.Has("tc-dim"):
				fp = append(fp, g.dimx(x))
			case x.Has("tc-push"):
				fp = append(fp, "w.Push{Items: []w.PushPart{"+g.buttons(x.Els())+"}}")
			default:
				stop("foot? " + join(x.Cls(), " "))
			}
		}
		f = append(f, "Foot: []w.SheetFootPart{"+join(fp, ", ")+"}")
	}
	for _, v := range allOf(e, sheetWidth) {
		f = append(f, "Width: w."+v)
	}
	return "w.Sheet{" + join(f, ", ") + "}"
}

func (g *gen) scrim(e *El) string {
	s := e.Kid("tc-sheet")
	if s == nil {
		pad := ""
		if e.Has("tc-scrim--padt38") {
			pad = ", Pad: w.ScrimPad38"
		} else if e.Has("tc-scrim--padt26") {
			pad = ", Pad: w.ScrimPad26"
		}
		return "w.Scrim{Content: " + g.anyblock(e.Els()[0]) + pad + scrimAttrs(e) + "}"
	}
	pad := ""
	for _, v := range allOf(e, []kv{{"tc-scrim--padt26", "ScrimPad26"}, {"tc-scrim--padt38", "ScrimPad38"}}) {
		pad = ", Pad: w." + v
	}
	content := tryWait(func() string { return g.sheet(s) }, `w.Sheet{Title: "Sheet"}`)
	return "w.Scrim{Content: " + content + pad + scrimAttrs(e) + "}"
}

// scrimAttrs are a scrim's name, its closed state and the form a drop sends.
func scrimAttrs(e *El) string {
	s := ""
	if v := e.Attr("id"); v != "" {
		s += ", ID: " + q(v)
	}
	if e.HasAttr("hidden") {
		s += ", Closed: true"
	}
	if v := e.Attr("data-tc-drop"); v != "" {
		s += ", Drop: " + q(v)
	}
	return s
}

func (g *gen) overlay(e *El) string {
	var f []string
	if t := e.Kid("tc-top"); t != nil {
		f = append(f, "Top: &"+g.topAny(t))
	}
	b := g.body(e.Kid("tc-shell__body"))
	b = strings.Replace(b, "w.ShellBody{", "w.OverlayBody{", -1)
	b = strings.Replace(b, "Rail: []w.RailPart{", "Rail: []w.OverlayRailPart{", -1)
	f = append(f, "Body: "+b)
	sc := e.Kid("tc-scrim")
	f = append(f, tryWait(func() string { return "Scrim: " + g.scrim(sc) }, `Scrim: w.Scrim{Content: w.Sheet{Title: "Sheet"}}`))
	return "w.Overlay{" + join(f, ", ") + "}"
}

var dialogVariant = []kv{{"is-wide", "DialogWide"}, {"is-narrow", "DialogNarrow"}, {"tc-dialog--max520", "Dialog520"}, {"tc-dialog--max540", "Dialog540"},
	{"tc-dialog--max560", "Dialog560"}, {"tc-dialog--max440-padt0-padb0", "DialogPage440"}, {"tc-dialog--max520-padt0-padb0", "DialogPage520"},
	{"tc-dialog--max420-flush", "DialogCanvas420"}}

func (g *gen) dialog(e *El) string {
	var items []string
	for _, x := range e.Els() {
		switch {
		case x.Has("tc-head"):
			items = append(items, g.head(x))
		case x.Has("tc-form"):
			items = append(items, g.form(x))
		case x.Has("tc-formfoot"):
			items = append(items, g.formFoot(x))
		case x.Has("tc-bar"):
			items = append(items, g.bar(x))
		case x.Has("tc-head__title"):
			s := "w.DialogTitle{Text: " + q(x.Text())
			if v := firstOf(x, []kv{{"tc-head__title--bot4", "DialogTitleGap4"}, {"tc-head__title--bot6", "DialogTitleGap6"}, {"tc-head__title--bot8", "DialogTitleGap8"}}); v != "" {
				s += ", Gap: w." + v
			}
			items = append(items, s+"}")
		case x.Has("tc-head__sub"):
			items = append(items, "w.DialogSub{Parts: "+g.inl(x)+"}")
		case x.Has("tc-dialog__logo") || x.Has("tc-dialog__logo--wide"):
			s := "w.DialogLogo{Logo: " + logo(x.Kid("tc-brand__logo"))
			if x.Has("tc-dialog__logo--wide") {
				s += ", Wide: true"
			}
			items = append(items, s+"}")
		case x.Has("tc-auth__tagline"):
			s := "w.Tagline{Lead: " + q(x.Els()[0].Text()) + ", Text: " + q(x.OwnText())
			if x.Has("tc-auth__tagline--mid") {
				s += ", Centered: true"
			}
			items = append(items, s+"}")
		case x.Has("tc-auth__stack"):
			var its, links []string
			for _, y := range x.Els() {
				switch {
				case y.Has("tc-auth__links"):
					links = nil
					for _, z := range y.Els() {
						l := "{Text: " + q(z.Text())
						if h := z.Attr("href"); h != "" {
							l += ", Href: " + q(h)
						}
						links = append(links, l+"}")
					}
				case y.Has("tc-form"):
					its = append(its, g.form(y))
				default:
					its = append(its, g.button(y))
				}
			}
			items = append(items, "w.AuthStack{Items: []w.AuthStackPart{"+join(its, ", ")+"}, Links: []w.Link{"+join(links, ", ")+"}}")
		case x.Has("tc-dialog__lede"):
			items = append(items, "w.Lede{Text: "+q(x.Text())+"}")
		case x.Has("tc-dialog__split"):
			ph := x.Kid("tc-dialog__photo")
			side := x.Kid("tc-dialog__side")
			var flds []string
			var aft *El
			for _, y := range side.Els()[1:] {
				if y.Has("tc-field") {
					flds = append(flds, g.field(y))
				}
				if aft == nil && y.Has("tc-dim") {
					aft = y
				}
			}
			src := ""
			if img := ph.Kid("tc-dialog__photo-img"); img != nil {
				src = ", Src: " + q(img.Attr("src"))
			}
			s := "w.PhotoCrop{Colour: " + q(ph.Attr("data-tc-colour")) + src + ", Note: " + q(side.Els()[0].Text()) + ", Fields: []w.Field{" + join(flds, ", ") + "}"
			if aft != nil {
				s += ", After: &w.Dim{Parts: []w.Inline{w.Text(" + q(aft.Text()) + ")}, Size: w.Dim11}"
			}
			items = append(items, s+"}")
		default:
			items = append(items, g.anyblock(x))
		}
	}
	f := []string{"Items: []w.DialogPart{\n\t\t" + join(items, ",\n\t\t") + "}"}
	for _, v := range allOf(e, dialogVariant) {
		f = append(f, "Variant: w."+v)
	}
	return "w.Dialog{" + join(f, ", ") + "}"
}

func (g *gen) auth(e *El) string {
	d := tryWait(func() string { return g.dialog(e.Kid("tc-dialog")) }, "w.Dialog{Variant: w.DialogNarrow}")
	return "w.Auth{Dialog: " + d + "}"
}

func emblem(e *El) string {
	f := []string{`Symbol: "davloom"`}
	for _, k := range e.Els() {
		if k.Tag == "title" {
			f = append(f, "Title: "+q(k.Text()))
			break
		}
	}
	if e.Has("is-alert") {
		f = append(f, "Alert: true")
	}
	return "w.Emblem{" + join(f, ", ") + "}"
}

func (g *gen) railnote(e *El) string {
	var f []string
	if e.Has("is-alert") {
		f = append(f, "Tone: w.RailStatusAlert")
	}
	if e.Has("is-off") {
		f = append(f, "Tone: w.RailStatusOff")
	}
	f = append(f, "Text: "+q(e.OwnText()), "Detail: "+g.dimx(e.Kid("tc-dim")))
	return "w.RailStatus{" + join(f, ", ") + "}"
}

func navItem(a *El) string {
	var f []string
	if em := a.Kid("tc-emblem"); em != nil {
		f = append(f, "Emblem: &"+emblem(em))
	} else {
		f = append(f, "Icon: "+icon(svgs(a)[0]))
	}
	if t := a.OwnText(); t != "" {
		f = append(f, "Label: "+q(t))
	}
	if full := a.Kid("tc-nav__full"); full != nil {
		f = append(f, "Label: "+q(full.Text()), "Short: "+q(a.Kid("tc-nav__short").Text()))
	}
	if n := a.Kid("tc-num"); n != nil {
		s := "Count: &w.Num{Value: " + n.Text()
		if n.Has("is-alert") {
			s += ", Alert: true"
		}
		f = append(f, s+"}")
	}
	f = append(f, `Href: "#"`)
	if t := a.Attr("title"); t != "" {
		f = append(f, "Title: "+q(t))
	}
	return "{" + join(f, ", ") + "}"
}

func (g *gen) nav(e *El) string {
	var items []string
	on := 0
	for i, a := range e.Els() {
		if a.Has("is-on") {
			on = i + 1
		}
		items = append(items, navItem(a))
	}
	f := []string{"Items: []w.NavItem{\n\t\t" + join(items, ",\n\t\t") + "}"}
	if on > 0 {
		f = append(f, "On: "+itoa(on))
	}
	if e.Has("tc-nav--top2") {
		f = append(f, "Apart: true")
	}
	return "w.Nav{" + join(f, ", ") + "}"
}

func usermenu(e *El) string {
	f := []string{"Name: " + q(e.OwnText())}
	if h := e.Attr("href"); h != "" {
		f = append(f, "Href: "+q(h))
	}
	if e.Has("is-focus") {
		f = append(f, "Focus: true")
	}
	if len(svgs(e)) == 0 {
		f = append(f, "NoArrow: true")
	}
	return "w.UserMenu{" + join(f, ", ") + "}"
}

func (g *gen) search(e *El) string {
	f := []string{"Query: " + queryOf(e.Kid("tc-q"))}
	if e.Has("is-focus") {
		f = append(f, "Focus: true")
	}
	if len(svgs(e)) > 1 {
		f = append(f, "Clear: true")
	}
	for _, v := range allOf(e, []kv{{"tc-search--full", "SearchFull"}, {"tc-search--fill180-wauto", "SearchFill"}}) {
		f = append(f, "Variant: w."+v)
	}
	if e.Tag == "form" {
		f = append(f, "Action: "+q(e.Attr("action")))
	}
	return "w.Search{" + join(f, ", ") + "}"
}

var logoSize = []kv{{"tc-brand__logo--h17-w44", "Logo17"}, {"tc-brand__logo--h42-w104", "Logo42"}, {"tc-brand__logo--h60-w150", "Logo60"}, {"tc-brand__logo--h48-w118", "Logo48"}}

func logo(svg *El) string {
	f := []string{`Label: "Carrel"`, `ViewBox: "4.79 -31.36 291.08 116.63"`}
	for _, v := range allOf(svg, logoSize) {
		f = append(f, "Size: w."+v)
	}
	return "w.Logo{" + join(f, ", ") + "}"
}

func (g *gen) brand(e *El) string {
	if svg := e.Kid("tc-brand__logo"); svg != nil {
		return "w.Brand{Logo: &" + logo(svg) + "}"
	}
	return "w.Brand{Title: " + q(e.Text()) + "}"
}

func (g *gen) child(e *El) string {
	table := []struct {
		c  string
		fn func(*El) string
	}{{"tc-btn", g.button}, {"tc-avatar", g.avatar}, {"tc-badge", g.badge}, {"tc-tag", g.tag}, {"tc-rubric", func(x *El) string { return g.rubric(x, "") }}, {"tc-tickrow", g.tickrow},
		{"tc-tree", g.tree}, {"tc-nav", g.nav}, {"tc-mono", g.mono}, {"tc-emblem", emblem}, {"tc-railstatus", g.railnote},
		{"tc-usermenu", usermenu}, {"tc-search", g.search}, {"tc-brand", g.brand}}
	for _, t := range table {
		if e.Has(t.c) {
			return t.fn(e)
		}
	}
	if e.Has("tc-drop") {
		return g.drop(e)
	}
	switch {
	case e.Has("tc-close"):
		return g.button(e)
	case e.Has("tc-segfold"):
		return g.segfold(e)
	case e.Has("tc-seg"):
		return g.seg(e)
	case e.Has("tc-field"):
		return g.field(e)
	case e.Has("tc-in") && e.Has("is-ph"):
		return "w.Filter{Query: " + queryOf(e.Kid("tc-q")) + ", Waiting: true, Variant: w.FilterH26W190}"
	}
	stop("child? " + join(e.Cls(), " "))
	return ""
}

func (g *gen) zone(e *El) string {
	return "w.Zone{Title: " + g.rubric(e.Kid("tc-rubric"), "") + ", Note: " + g.mono(e.Kid("tc-mono")) + "}"
}

func (g *gen) crumbs(e *El) string {
	var path []string
	var cur *El
	for _, a := range e.Els() {
		if a.Has("tc-crumbs__link") {
			path = append(path, "{Label: "+q(a.Text())+`, Href: "#"}`)
		}
		if cur == nil && a.Tag == "b" {
			cur = a
		}
	}
	var f []string
	if len(path) > 0 {
		f = append(f, "Path: []w.Crumb{"+join(path, ", ")+"}")
	}
	f = append(f, "Current: "+q(cur.Text()))
	if d := e.Kid("tc-dim"); d != nil {
		f = append(f, "Aside: &"+dimText(d.Text()))
	}
	return "w.Crumbs{" + join(f, ", ") + "}"
}

func (g *gen) acts(e *El) string {
	f := []string{"Items: []w.ActsPart{" + g.each(e.Els(), g.child) + "}"}
	if e.Has("tc-acts--stack") {
		f = append(f, "Stack: true")
	}
	return "w.Acts{" + join(f, ", ") + "}"
}

func (g *gen) railsec(e *El) string {
	var items, f []string
	for _, k := range e.Els() {
		switch {
		case k.Has("tc-railsec__tags"):
			f = append(f, "Tags: []w.Tag{"+g.each(k.Els(), g.tag)+"}")
		case k.Has("tc-railsec__note"):
			f = append(f, "Note: "+q(k.Text()))
		default:
			items = append(items, g.child(k))
		}
	}
	var out []string
	if len(items) > 0 {
		out = append(out, "Items: []w.RailSecPart{\n\t\t"+join(items, ",\n\t\t")+"}")
	}
	out = append(out, f...)
	if e.Has("tc-railsec--top14") {
		out = append(out, "Apart: true")
	}
	return "w.RailSec{" + join(out, ", ") + "}"
}

func (g *gen) railfoot(e *El) string {
	var items, f []string
	for _, k := range e.Els() {
		if k.Has("tc-railfoot__about") {
			f = append(f, "About: &w.About{Name: "+q(k.OwnText())+", Build: "+g.mono(k.Kid("tc-mono"))+`, Href: "#rest"}`)
		} else {
			items = append(items, g.child(k))
		}
	}
	var out []string
	if len(items) > 0 {
		out = append(out, "Items: []w.RailFootPart{"+join(items, ", ")+"}")
	}
	out = append(out, f...)
	if e.Has("tc-railfoot--strip") {
		out = append(out, "Strip: true")
	}
	return "w.RailFoot{" + join(out, ", ") + "}"
}

// shell is the whole shell: its top bar as drawn, then an overlay or the
// body — also when the body stands inside a desktop window.
func (g *gen) shell(e *El) string {
	var f []string
	if t := e.Kid("tc-top"); t != nil {
		f = append(f, "Top: &"+g.topAny(t))
	}
	if o := e.Kid("tc-overlay"); o != nil {
		f = append(f, "Body: "+g.overlay(o))
		return "w.Shell{" + join(f, ", ") + "}"
	}
	f = append(f, "Body: "+g.body(e.Kid("tc-shell__body")))
	return "w.Shell{" + join(f, ", ") + "}"
}

func (g *gen) printRow(x *El) string {
	var its []string
	for _, y := range x.Els() {
		switch {
		case y.Has("tc-mono"):
			its = append(its, g.mono(y))
		case y.Has("tc-print__num"):
			its = append(its, "w.PrintNum{Text: "+q(y.Text())+"}")
		case y.Has("tc-print__cell"):
			var b, d *El
			for _, z := range y.Els() {
				if b == nil && z.Tag == "b" {
					b = z
				}
				if d == nil && z.Has("tc-dim--print") {
					d = z
				}
			}
			if b != nil {
				s := "w.PrintCell{Name: " + q(b.Text())
				if d != nil {
					s += ", Detail: " + q(d.Text())
				}
				its = append(its, s+"}")
			} else {
				its = append(its, "w.PrintCell{Name: "+q(y.Text())+"}")
			}
		default:
			stop("print row? " + join(y.Cls(), " "))
		}
	}
	s := "w.PrintRow{Items: []w.PrintRowPart{" + join(its, ", ") + "}"
	if x.Has("tc-print__row--3col") {
		s += ", Columns: true"
	}
	return s + "}"
}

func (g *gen) windowPart(x *El) string {
	switch {
	case x.Has("tc-form"):
		return g.form(x)
	case x.Has("tc-msg"):
		return g.msg(x)
	case x.Has("tc-sheet__fields"):
		return "w.SheetFields{Items: []w.Field{" + g.each(x.Els(), g.field) + "}}"
	case x.Has("tc-line"):
		return g.line(x)
	case x.Has("tc-stack"):
		gap := "StackGap9"
		if x.Has("tc-stack--gap-10") {
			gap = "StackGap10"
		}
		return "w.Stack{Gap: w." + gap + ", Items: []w.StackPart{" + g.each(x.Els(), func(y *El) string {
			if y.Has("tc-card") {
				return g.card(y)
			}
			return g.button(y)
		}) + "}}"
	case x.Tag == "svg" && x.Has("tc-brand__logo"):
		return logo(x)
	case x.Has("tc-win__note"):
		return "w.WindowNote{Text: " + q(x.Text()) + "}"
	case x.Has("tc-diag"):
		return g.diag(x)
	case x.Has("tc-fset"):
		return g.fset(x)
	case x.Tag == "textarea":
		return g.area(x)
	case x.Has("tc-dim"):
		return g.dimx(x)
	}
	stop("window part? " + x.Tag + " " + join(x.Cls(), " "))
	return ""
}

func (g *gen) window(x *El) string {
	gap := firstOf(x, []kv{{"tc-stack--gap-9", "WindowGap9"}, {"tc-stack--gap-10", "WindowGap10"}, {"tc-stack--gap-12", "WindowGap12"},
		{"tc-stack--gap-13", "WindowGap13"}, {"tc-stack--gap-14", "WindowGap14"}})
	if gap == "" {
		stop("window without a gap")
	}
	s := "w.Window{Items: []w.WindowPart{" + g.each(x.Els(), g.windowPart) + "}, Gap: w." + gap
	if x.Has("tc-win__sheet--pad12") {
		s += ", Small: true"
	}
	return s + "}"
}

var printParts = []kv{{"tc-print__h1", "PrintTitle"}, {"tc-print__sources", "PrintSources"}, {"tc-print__rub", "PrintRubric"}, {"tc-print__foot", "PrintFoot"}}

func (g *gen) pagePart(x *El) string {
	switch {
	case x.Has("tc-auth"):
		return g.auth(x)
	case x.Has("tc-win__body--pad16-20") || x.Has("tc-win__sheet--pad12"):
		return g.window(x)
	case x.Has("tc-msg"):
		return g.msg(x)
	case x.Has("tc-form"):
		return g.form(x)
	}
	if t := firstOf(x, printParts); t != "" {
		return "w." + t + "{Text: " + q(x.Text()) + "}"
	}
	if x.Has("tc-print__row") {
		return g.printRow(x)
	}
	stop("page part? " + x.Tag + " " + join(x.Cls(), " "))
	return ""
}

var pageVariants = []kv{{"tc-stack", "PageOnBackgroundColumn"}, {"tc-root--on-bg", "PageOnBackground"}, {"tc-root--print", "PagePrint"}}

func (g *gen) page(root *El) string {
	s := "w.Page{Items: []w.PagePart{" + g.each(root.Els(), g.pagePart) + "}"
	if v := firstOf(root, pageVariants); v != "" {
		s += ", Variant: w." + v
	}
	return s + "}"
}

// avatar is a record's face: its initials, the place of a photo, or the
// photo itself (an img in the place).
func (g *gen) avatar(a *El) string {
	var f []string
	switch img := a.Kid("tc-avatar__img"); {
	case img != nil:
		f = append(f, "Src: "+q(img.Attr("src")))
	case a.Has("is-photo"):
		f = append(f, "Photo: true")
	default:
		f = append(f, "Initials: "+q(a.Text()))
	}
	for _, v := range allOf(a, []kv{{"tc-avatar--sq34-t12", "Avatar34"}, {"tc-avatar--sq44-t15-r4", "Avatar44"}}) {
		f = append(f, "Size: w."+v)
	}
	return "w.Avatar{" + join(f, ", ") + "}"
}

var columnOf = []kv{{"tc-col--name", "ColumnName"}, {"tc-col--phone", "ColumnPhone"}, {"tc-col--email", "ColumnEmail"}, {"tc-col--org", "ColumnOrg"},
	{"tc-col--birthday", "ColumnBirthday"}, {"tc-col--tags", "ColumnTags"}, {"tc-col--changed", "ColumnChanged"}, {"tc-col--source", "ColumnSource"},
	{"tc-col--due", "ColumnDue"}, {"tc-col--complete", "ColumnComplete"}, {"tc-col--priority", "ColumnPriority"}, {"tc-col--status", "ColumnStatus"},
	{"tc-col--excerpt", "ColumnExcerpt"}, {"tc-col--date", "ColumnDate"}, {"tc-col--detail", "ColumnDetail"}, {"tc-col--kind", "ColumnKind"}}

// columnRow is a row of the columns a person chose.
func (g *gen) columnRow(e *El) string {
	var f, cells []string
	els := e.Els()
	if els[0].Has("tc-tick") {
		d := ""
		if els[0].Has("is-done") {
			d = "Done: true, "
		}
		f = append(f, "Tick: &w.Tick{"+d+`Label: "Done"}`)
		els = els[1:]
	}
	f = append(f, "Mark: "+g.mark(els[0]))
	for _, k := range els[1:] {
		if k.Has("tc-avatar") {
			f = append(f, "Avatar: &"+g.avatar(k))
			continue
		}
		for _, c := range columnOf {
			if !k.Has(c.k) {
				continue
			}
			if c.v == "ColumnTags" {
				var tags []string
				for _, t := range k.Els() {
					tags = append(tags, g.tag(t))
				}
				cells = append(cells, "{Column: w.ColumnTags, Tags: []w.RowTag{"+join(tags, ", ")+"}}")
			} else {
				cells = append(cells, "{Column: w."+c.v+", Text: "+q(k.Text())+"}")
			}
			if c.v == "ColumnName" && k.Has("is-strong") {
				f = append(f, "Strong: true")
			}
		}
	}
	f = append(f, "Cells: []w.ColumnCell{"+join(cells, ", ")+"}")
	if e.Has("is-on") {
		f = append(f, "On: true")
	}
	if e.Has("is-overdue") {
		f = append(f, "Overdue: true")
	}
	return "w.ColumnRow{" + join(f, ", ") + "}"
}
