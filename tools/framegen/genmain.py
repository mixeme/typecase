"""Emit refs records for main, side and shell."""
import os

HERE = os.path.dirname(os.path.abspath(__file__))
src = open(os.path.join(HERE, "genhead.py"), encoding="utf-8").read()
exec(src.split("\nout, waiting = [], []")[0])


class Wait(Exception):
    pass


def block_child(k):
    if k.has("tc-head"):
        return head(k)
    if k.has("tc-bar"):
        return bar2(k)
    if k.has("tc-list"):
        return lst(k)
    if k.has("tc-form"):
        return formw(k)
    if k.has("tc-formfoot"):
        return formfootw(k)
    if k.has("tc-railsec"):
        return railsec(k)
    if k.has("tc-zone"):
        return zone(k)
    return anyblock(k)


def listgroup(k):
    r = kid(k, "tc-rubric")
    n = kid(k, "tc-num")
    f = ["Title: w.Rubric{Text: " + q(r.text()) + "}"]
    if n is not None:
        f.append("Count: &w.Num{Value: " + n.text() + (", Alert: true" if n.has("is-alert") else "") + "}")
    return "w.ListGroup{" + ", ".join(f) + "}"


def main(e):
    items = []
    for k in e.els():
        if k.has("tc-list__group"):
            items.append(listgroup(k))
        elif k.has("tc-main__tree"):
            items.append("w.MainTree{Tree: " + tree(k.els()[0]) + "}")
        elif k.has("tc-field__hint--gutter"):
            items.append("w.MainHint{Text: " + q(k.text()) + "}")
        elif k.has("tc-nav--bottom"):
            its, on = [], 0
            for i, a in enumerate(k.els()):
                if a.has("is-on"):
                    on = i + 1
                its.append("{Icon: " + icon([x for x in a.els() if x.tag == "svg"][0]) + ", Label: " + q(a.own_text()) + ", Href: \"#\"}")
            items.append("w.BottomNav{Items: []w.BottomItem{" + ", ".join(its) + "}, On: %d}" % on)
        elif k.has("tc-main__banner") or k.has("tc-main__banner--12"):
            f2 = ["Logo: " + logo(kid(k, "tc-brand__logo"))]
            b = kid(k, "tc-badge")
            if b is not None:
                f2.append("Badge: &" + badge(b))
            end = kid(k, "tc-push")
            if end is not None:
                inner = ", ".join(button(x) if x.has("tc-btn") else child(x) for x in end.els())
                if end.has("tc-line"):
                    f2.append("End: w.Line{Gap: w.LineGap6, Place: w.Pushed, Items: []w.LinePart{" + inner + "}}")
                else:
                    f2.append("End: w.Push{Items: []w.PushPart{" + inner + "}}")
            if k.has("tc-main__banner"):
                f2.append("Wide: true")
            items.append("w.Banner{" + ", ".join(f2) + "}")
        elif k.has("tc-main__strip"):
            parts = []
            for x in k.els()[1:]:
                if x.has("tc-btn"):
                    parts.append(button(x))
                elif x.has("tc-dim"):
                    parts.append(dimx(x))
                else:
                    parts.append(inputw(x))
            items.append("w.Strip{Label: " + q(k.els()[0].text()) + ", Items: []w.StripPart{" + ", ".join(parts) + "}}")
        elif k.has("tc-main__foot") or k.has("tc-main__foot--16"):
            items.append("w.MainFoot{Text: " + q(k.text()) + (", Roomy: true" if k.has("tc-main__foot--16") else "") + "}")
        elif k.has("tc-main__results"):
            items.append("w.MainResults{Items: []w.ResultsPart{" + ", ".join(anyblock(y) for y in k.els()) + "}}")
        elif k.has("tc-main__cards"):
            items.append("w.MainCards{Cards: []w.Card{" + ", ".join(card(y) for y in k.els()) + "}}")
        elif k.has("tc-main__diag"):
            items.append("w.MainDiag{Diag: " + diag(k.els()[0]) + "}")
        elif k.has("tc-main__thumbs"):
            items.append("w.MainThumbs{Thumb: " + thumb(k.els()[0]) + "}")
        elif k.has("tc-main__stage"):
            it = kid(k, "tc-stack__item")
            items.append("w.Stage{Thumb: " + thumb(kid(k, "tc-thumb")) + ", Title: " + q(kid(it, "tc-head__title").text()) + ", Meta: w.Mono{Text: " +
                         q(kid(it, "tc-head__sub").text()) + "}, Download: " + button(kid(k, "tc-btn")) + ", Note: " + q(kid(k, "tc-main__stage-note").text()) + "}")
        elif k.has("tc-main__person"):
            ph = kid(k, "tc-avatar--photo")
            side_ = kid(k, "tc-main__person-side")
            sub = kid(side_, "tc-head__sub")
            b = kid(sub, "tc-badge")
            nxt = [y for y in e.els() if y.has("tc-main__person-acts")]
            acts_ = ", Acts: []w.Button{" + ", ".join(button(y) for y in nxt[0].els()) + "}" if nxt else ""
            items.append("w.Person{Initials: " + q(ph.text()) + ", Name: " + q(kid(side_, "tc-main__person-name").text()) + ", Sub: " + q(sub.own_text()) +
                         (", Badge: &" + badge(b) if b is not None else "") + acts_ + "}")
        elif k.has("tc-main__person-acts"):
            pass
        elif k.has("tc-main__hero"):
            side_ = kid(k, "tc-main__hero-side")
            fl = kid(side_, "tc-fields")
            items.append('w.Hero{Symbol: "davloom-tile", Label: "davloom", Title: ' + q(kid(side_, "tc-main__hero-title").text()) + ", Sub: " + q(kid(side_, "tc-head__sub").text()) +
                         ", Text: " + prose(kid(side_, "tc-prose")) + (", Fields: &" + fields(fl) if fl is not None else "") + ", Acts: []w.Button{" +
                         ", ".join(button(y) for y in kid(side_, "tc-doc__acts").els()) + "}}")
        elif k.has("tc-doc__layout"):
            view = ", View: w.DocViewMeasure" if k.has("is-measure") else (", View: w.DocViewFocus" if k.has("is-focus") else "")
            its = []
            for y in k.els():
                if y.has("tc-docmeta"):
                    its.append(docmeta2(y))
                elif y.has("tc-docbar"):
                    its.append(docbar(y))
                else:
                    its.append(doc(y))
            items.append("w.DocLayout{Items: []w.DocLayoutPart{" + ", ".join(its) + "}" + view + "}")
        elif k.has("tc-main__foot-acts"):
            items.append("w.FootActs{Items: []w.Button{" + ", ".join(button(b) for b in k.els()) + "}}")
        else:
            items.append(block_child(k))
    f = ["Items: []w.MainPart{\n\t\t" + ",\n\t\t".join(items) + "}"]
    for k, v in {"tc-main--zone": "MainZone", "tc-main--padb180": "MainRoomy", "tc-main--minh88": "MainShort"}.items():
        if e.has(k):
            f.append("Variant: w." + v)
    return "w.Main{" + ", ".join(f) + "}"


def side(e):
    items = []
    for k in e.els():
        if k.has("tc-side__acts"):
            items.append("w.SideActs{Items: []w.Button{" + ", ".join(button(b) for b in k.els()) + "}}")
        else:
            items.append(block_child(k))
    f = ["Items: []w.SidePart{" + ", ".join(items) + "}"]
    for k, v in {"tc-side--zone": "SideZone", "tc-side--drop": "SideDrop"}.items():
        if e.has(k):
            f.append("Variant: w." + v)
    return "w.Side{" + ", ".join(f) + "}"


def body(b):
    f = []
    rail = kid(b, "tc-shell__rail")
    if rail is not None:
        f.append("Rail: []w.RailPart{" + ", ".join(child_rail(k) for k in rail.els()) + "}")
        if rail.has("tc-shell__rail--icons"):
            f.append("RailForm: w.RailIcons")
        if rail.has("tc-shell__rail--zone"):
            f.append("RailForm: w.RailZone")
    m = kid(b, "tc-main")
    try:
        f.append("Main: " + main(m))
    except Wait:
        f.append("Main: w.Main{}")
    s = kid(b, "tc-side")
    if s is not None:
        try:
            f.append("Side: &" + side(s))
        except Wait:
            f.append("Side: &w.Side{}")
    for k, v in {"tc-shell__body--minh230": "ShellHeight230", "tc-shell__body--minh280": "ShellHeight280", "tc-shell__body--minh340": "ShellHeight340",
                 "tc-shell__body--minh390": "ShellHeight390", "tc-shell__body--minh560": "ShellHeight560"}.items():
        if b.has(k):
            f.append("Height: w." + v)
    return "w.ShellBody{" + ", ".join(f) + "}"


def child_rail(k):
    if k.has("tc-railfoot"):
        return railfoot(k)
    if k.has("tc-railsec"):
        return railsec(k)
    if k.has("tc-nav"):
        return nav(k)
    if k.has("tc-zone"):
        return zone(k)
    raise SystemExit("rail? " + str(k.cls))


def shellw(e):
    f = []
    t = kid(e, "tc-top")
    if kid(e, "tc-overlay") is not None:
        return None
    if t is not None:
        tt = top(t)
        f.append("Top: &" + (tt if tt is not None else 'w.Top{Brand: w.Brand{Title: "Carrel"}}'))
    b = kid(e, "tc-shell__body")
    if b is None:
        # the desktop window: the body stands inside the window's drawing
        for k in e.els():
            if kid(k, "tc-shell__body") is not None:
                b = kid(k, "tc-shell__body")
                t = kid(k, "tc-top")
                if t is not None and not f:
                    f.append("Top: &" + top(t))
    f.append("Body: " + body(b))
    return "w.Shell{" + ", ".join(f) + "}"


def mstate(cls, e, expr):
    if cls == "tc-main":
        for k, v in [("MainZone", "zone"), ("Variant:", "variant"), ("BottomNav", "phone"), ("ListGroup", "groups")]:
            if k in expr:
                return v
        return "plain"
    if cls == "tc-side":
        for k, v in [("SideZone", "zone"), ("SideDrop", "drop"), ("SideActs", "acts")]:
            if k in expr:
                return v
        return "plain"
    for k, v in [("RailZone", "zone"), ("RailIcons", "icons"), ("Side: &", "side"), ("Rail:", "plain")]:
        if k in expr:
            return v if k != "Rail:" or "Top:" in expr else "no-top"
    return "no-rail" if "Top:" in expr else "no-top"


out, waiting = [], []
for cls, fn in [("tc-main", main), ("tc-side", side), ("tc-shell", shellw)]:
    out.append("\n\t// " + cls)
    for frame, n, e in shapes(cls):
        try:
            expr = fn(e)
        except Wait:
            expr = None
        if expr is None:
            waiting.append("ref(%s, %s, %d)" % (q(frame), q(cls), n))
            continue
        out.append("\tshell(%s, %s, %s, %s, %d)," % (q(mstate(cls, e, expr)), expr, q(frame), q(cls), n))
print("\n".join(out))
print("\n// WAIT " + ", ".join(waiting))
