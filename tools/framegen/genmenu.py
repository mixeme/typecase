"""Emit refs records for menus and drops, and the forms that waited for them."""
import os

HERE = os.path.dirname(os.path.abspath(__file__))
src = open(os.path.join(HERE, "genmain.py"), encoding="utf-8").read()
exec(src.split("\nout, waiting = [], []")[0])


def trail(k):
    if k.has("tc-mono"):
        v = {"tc-mono--t10-dusty": "Mono10Dusty", "tc-mono--t11-dusty-ls0": "Mono11DustyTight", "tc-mono--t11": "Mono11"}
        var = [x for c, x in v.items() if k.has(c)]
        return mono(k)
    if k.has("tc-dim"):
        size = ""
        for c, x in {"tc-dim--t10-5": "Dim10_5", "tc-dim--t11": "Dim11"}.items():
            if k.has(c):
                size = ", Size: w." + x
        var = ", Variant: w.DimTight" if k.has("tc-dim--ls0") else ""
        return dimx(k)
    if k.has("tc-badge"):
        b = badge(k)
        return b[:-1] + (", Place: w.Pushed}" if k.has("tc-push") else "}")
    if k.has("tc-kbd"):
        return "w.Kbd{Keys: " + q(k.text()) + (", Place: w.Pushed" if k.has("tc-push") else "") + "}"
    if k.has("tc-push"):
        return "w.Push{Items: []w.PushPart{w.Text(" + q(k.text()) + ")}}"
    raise SystemExit("trail? " + str(k.cls))


def menuitem(k):
    f = []
    kids = k.els()
    for x in kids:
        if x.tag == "svg":
            f.append("Icon: " + icon(x))
            if x.has("tc-ico--op-35"):
                f.append("Faint: true")
        elif x.has("tc-mark"):
            f.append("Mark: &" + mark(x))
        elif x.has("tc-tickrow__tick"):
            f.append("Tick: &w.TickBox{" + ("Ticked: w.TickedOff" if x.has("is-off") else "") + "}")
        elif x.has("tc-menu__label"):
            f.append("Found: " + inl(x))
        elif x.has("tc-dim--danger"):
            f.append("Note: &w.Span{Face: w.FaceDanger, Text: " + q(x.text()) + "}")
        elif x.has("tc-mono") and not x.has("tc-push"):
            f.append("Detail: &" + mono(x))
        elif x.has("tc-tag"):
            f.append("Trail: w.Tag{Label: " + q(x.text()) + ", Place: w.Pushed}")
        else:
            f.append("Trail: " + trail(x))
    if not any(p.startswith("Found:") for p in f):
        f.append("Label: " + q(k.own_text()))
    if k.has("is-on"):
        f.append("On: true")
    for c, x in {"tc-menu__item--danger": "MenuToneDanger", "tc-menu__item--dusty": "MenuToneDusty", "tc-menu__item--muted": "MenuToneMuted",
                 "tc-menu__item--accent": "MenuToneAccent", "tc-menu__item--note": "MenuToneAside"}.items():
        if k.has(c):
            f.append("Tone: w." + x)
    if k.has("is-3rd"):
        f.append("Tertiary: true")
    return "w.MenuItem{" + ", ".join(f) + "}"


def rubric_full(k, form=""):
    f = ["Text: " + q(k.own_text())]
    items = []
    for x in k.els():
        if x.has("tc-mono"):
            items.append(mono(x))
        elif x.has("tc-push"):
            items.append("w.Push{Items: []w.PushPart{w.Text(" + q(x.text()) + ")}}")
        elif x.has("tc-rubric__col"):
            items.append("w.RubricCol{Text: " + q(x.text()) + "}")
    if items:
        f.append("Items: []w.RubricPart{" + ", ".join(items) + "}")
    if form:
        f.append("Form: w." + form)
    if k.has("tc-line"):
        f.append("Variant: w.RubricItemLine")
    return "w.Rubric{" + ", ".join(f) + "}"


def menu(e):
    items = []
    for k in e.els():
        if k.has("tc-menu__sep"):
            items.append("w.MenuSep{" + ("Tertiary: true" if k.has("is-3rd") else "") + "}")
        elif k.has("tc-rubric") and k.has("tc-menu__item"):
            form = "MenuHeadingRoomy" if k.has("tc-menu__item--pad5-8-7") else ("MenuHeadingSplit" if k.has("tc-menu__item--split") else "")
            items.append("w.MenuHeading{Rubric: " + rubric_full(k, "RubricItem") + (", Form: w." + form if form else "") + "}")
        elif k.has("tc-rubric"):
            items.append(rubric_full(k, "RubricInline" if k.tag == "span" else ""))
        elif k.has("tc-menu__item--col"):
            seg = kid(k, "tc-seg")
            opts = ["{Label: " + q(o.text()) + (", On: true" if o.has("is-on") else "") + "}" for o in seg.els()]
            items.append("w.MenuColumn{Label: " + q(k.els()[0].text()) + ", Control: w.Seg{Options: []w.SegOption{" + ", ".join(opts) + "}, Variant: w.SegStartSmall}}")
        elif k.tag == "div" and k.has("tc-menu__item"):
            f = ["Label: " + q(k.own_text())]
            b = kid(k, "tc-badge")
            if b is not None:
                f.append("Badge: &" + badge(b))
            f.append("Default: " + trail(kid(k, "tc-dim")))
            i = kid(k, "tc-in")
            f.append("Input: w.Input{Value: " + q(i.attrs.get("value", "")) + ", Mono: true, Variant: w.InputChip, Label: " + q(i.attrs.get("aria-label", "")) + "}")
            items.append("w.MenuField{" + ", ".join(f) + "}")
        elif k.has("tc-menu__foot"):
            items.append("w.MenuFoot{Items: []w.Button{" + ", ".join(button(b) for b in k.els()) + "}}")
        elif k.has("tc-menu__item"):
            items.append(menuitem(k))
        elif k.has("tc-in"):
            items.append(inputw(k))
        elif k.has("tc-range"):
            items.append(baritem(k)[:-1] + ", Wrap: true}")
        else:
            raise SystemExit("menu? " + str(k.cls))
    f = ["Items: []w.MenuPart{\n\t\t" + ",\n\t\t".join(items) + "}"]
    for c, x in {"tc-menu--w212": "MenuWidth212", "tc-menu--w290": "MenuWidth290", "tc-menu--w420": "MenuWidth420", "tc-menu--full": "MenuWidthFull"}.items():
        if e.has(c):
            f.append("Width: w." + x)
    if e.has("is-pop"):
        f.append("Pop: w." + ("MenuUnderEnd" if e.has("is-right") else "MenuUnder"))
    if e.has("tc-menu--fields"):
        f.append("Fields: true")
    if e.tag == "span":
        f.append("Inline: true")
    return "w.Menu{" + ", ".join(f) + "}"


def drop(e):
    t, m = e.els()[0], e.els()[1]
    f = []
    if t.has("tc-btn"):
        b = button(t)
        f.append("Trigger: " + b)
        if t.has("tc-bar__more"):
            f.append("More: true")
    else:
        f.append('Trigger: w.Pick{Value: ' + q(t.text()) + ', Label: "Where"}')
    f.append("Menu: " + menu(m))
    return "w.Drop{Open: true, " + ", ".join(f) + "}"


def menustate(e, expr):
    for k, v in [("Fields: true", "fields"), ("Found:", "found"), ("Pop:", "pop"), ("MenuHeading", "heading"), ("Tick:", "ticks"),
                 ("MenuToneDanger", "danger"), ("Mark:", "marks"), ("On: true", "on"), ("Icon:", "icons")]:
        if k in expr:
            return v
    return "plain"


def top2(e):
    f = []
    els = e.els()
    f.append("Brand: " + brand(kid(e, "tc-brand")))
    line = kid(e, "tc-top__search")
    if line is not None:
        f.append("Search: &" + search(kid(line, "tc-search")))
        f.append("Suggest: &" + menu(kid(kid(line, "tc-top__suggest"), "tc-menu")))
    else:
        f.append("Search: &" + search(kid(e, "tc-search")))
    acts = []
    for k in kid(e, "tc-top__actions").els():
        if k.has("tc-drop__anchor"):
            panel = k.els()[1]
            acts.append("w.TopDrop{Open: true, Trigger: " + child(k.els()[0]) + ", Menu: " + menu(panel.els()[0]) + (", End: true" if panel.has("tc-drop__panel--right") else "") + "}")
        else:
            acts.append(child(k))
    f.append("Actions: []w.TopAction{" + ", ".join(acts) + "}")
    return "w.Top{" + ", ".join(f) + "}"


def bar2(e):
    f, items = [], []
    for k in e.els():
        if k.has("tc-drop"):
            items.append(drop(k))
        elif k.has("tc-bar__count"):
            f.append("Count: &w.BarCount{Text: " + q(k.text()) + (", Small: true" if k.has("tc-bar__count--t12-5") else "") + "}")
        elif k.has("tc-bar__more"):
            f.append("More: &" + button(k))
        elif k.has("tc-bar__right"):
            f.append("Right: &w.BarRight{Items: []w.BarPart{" + ", ".join(baritem(x) for x in k.els()) + "}" + (", Secondary: true" if k.has("is-2nd") else "") + "}")
        else:
            items.append(baritem(k))
    f.insert(0, "Items: []w.BarPart{" + ", ".join(items) + "}")
    for k, v in {"tc-bar--sel": "BarSelection", "tc-bar--wide": "BarWide"}.items():
        if e.has(k):
            f.append("Variant: w." + v)
    return "w.Bar{" + ", ".join(f) + "}"


out = []
out.append("\n\t// tc-menu")
for frame, n, e in shapes("tc-menu"):
    expr = menu(e)
    out.append("\toverlay(%s, %s, %s, %s, %d)," % (q(menustate(e, expr)), expr, q(frame), q("tc-menu"), n))
out.append("\n\t// tc-drop")
for frame, n, e in shapes("tc-drop"):
    expr = drop(e)
    out.append("\toverlay(%s, %s, %s, %s, %d)," % (q("more" if "More: true" in expr else ("pick" if "w.Pick" in expr[:40] else "button")), expr, q(frame), q("tc-drop"), n))
out.append("\n\t// the forms that waited for menus")
for frame, n, cls in [("s-find-2", 1, "tc-top"), ("s-notes-7", 1, "tc-top"), ("account-5", 1, "tc-top"), ("components-19", 1, "tc-bar"), ("s-agenda-1", 1, "tc-bar")]:
    for fr, nn, e in shapes(cls):
        if fr == frame and nn == n:
            expr = top2(e) if cls == "tc-top" else bar2(e)
            st = ("suggest" if "Suggest" in expr else "drop") if cls == "tc-top" else "more"
            out.append("\toverlay(%s, %s, %s, %s, %d)," % (q(st), expr, q(frame), q(cls), n))
print("\n".join(out))
