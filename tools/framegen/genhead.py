"""Emit refs records for heads and bars."""
import os

HERE = os.path.dirname(os.path.abspath(__file__))
src = open(os.path.join(HERE, "genshell.py"), encoding="utf-8").read()
exec(src.split("\nFAM = [")[0])


def subparts(d):
    parts = []
    for k in d.kids:
        if isinstance(k, str):
            t = re.sub(r"\s+", " ", k)
            if t.strip():
                parts.append("w.Text(" + q(t) + ")")
        elif k.tag == "b":
            parts.append("w.Strong(" + q(k.text()) + ")")
        elif k.has("tc-dim--danger"):
            parts.append("w.Span{Face: w.FaceDanger, Text: " + q(k.text()) + "}")
        elif k.has("tc-mono"):
            parts.append(mono(k))
        elif k.has("tc-mark"):
            parts.append(mark(k))
        elif k.has("tc-badge"):
            parts.append(badge(k))
        elif k.has("tc-btn"):
            parts.append(button(k))
        else:
            raise SystemExit("sub? " + str(k.cls))
    if parts and parts[0].startswith("w.Text("):
        parts[0] = "w.Text(" + q(eval(parts[0][7:-1]).lstrip()) + ")"
    if parts and parts[-1].startswith("w.Text("):
        parts[-1] = "w.Text(" + q(eval(parts[-1][7:-1]).rstrip()) + ")"
    return "w.HeadSub{Parts: []w.HeadSubPart{" + ", ".join(parts) + "}}"


def htitle(t):
    f = []
    m = kid(t, "tc-mark")
    if m is not None:
        f.append("Mark: &" + mark(m))
    f.append("Text: " + q(t.own_text()))
    b = kid(t, "tc-badge")
    if b is not None:
        f.append("Badge: &" + badge(b))
    for k, v in {"tc-head__title--t19": "HeadTitle19", "tc-head__title--t18": "HeadTitle18", "tc-head__title--t17": "HeadTitle17"}.items():
        if t.has(k):
            f.append("Size: w." + v)
    return "&w.HeadTitle{" + ", ".join(f) + "}"


def photo(p):
    f = []
    if p.text():
        f.append("Initials: " + q(p.text()))
    if p.has("tc-avatar--photo--sq76-t20"):
        f.append("Large: true")
    return "&w.HeadPhoto{" + ", ".join(f) + "}"


def titles_into(box, f):
    subs = []
    for k in box.els():
        if k.has("tc-head__title"):
            f.append("Title: " + htitle(k))
        elif k.has("tc-head__sub"):
            subs.append(subparts(k))
        elif k.has("tc-crumbs"):
            f.append("Crumbs: &" + crumbs(k))
        elif k.has("tc-rubric"):
            v = ", Variant: w.RubricBottom2" if k.has("tc-rubric--bot2") else ""
            f.append("Kicker: &w.Rubric{Text: " + q(k.own_text()) + v + "}")
        elif k.has("tc-sec__acts"):
            f.append("Tools: []w.Button{" + ", ".join(button(b) for b in k.els()) + "}")
        elif k.has("tc-head__text"):
            titles_into(k, f)
        elif k.has("tc-avatar--photo"):
            f.append("Photo: " + photo(k))
        elif k.has("tc-avatar"):
            f.append("Avatar: &w.Avatar{Initials: " + q(k.text()) + ", Size: w.Avatar44}")
        elif k.has("tc-head__titles"):
            titles_into(k, f)
        else:
            raise SystemExit("titles? " + str(k.cls))
    if subs:
        f.append("Subs: []w.HeadSub{" + ", ".join(subs) + "}")


def head(e):
    f = []
    for k in e.els():
        if k.has("tc-head__titles"):
            titles_into(k, f)
        elif k.has("tc-head__group"):
            f.append("Grouped: true")
            titles_into(k, f)
        elif k.has("tc-head__sub"):
            f.append("Meta: &" + subparts(k))
        elif k.has("tc-acts"):
            f.append("Acts: &" + acts(k))
        else:
            raise SystemExit("head? " + str(k.cls))
    if e.has("tc-head--mid"):
        f.append("Middle: true")
    for k, v in {"tc-head--padt11": "HeadPadTop11", "tc-head--padb6": "HeadPadBottom6", "tc-head--snug": "HeadPadBottom8",
                 "tc-head--padb10": "HeadPadBottom10", "tc-head--rule": "HeadRule"}.items():
        if e.has(k):
            f.append("Pad: w." + v)
    return "w.Head{" + ", ".join(f) + "}"


def baritem(k):
    if k.has("tc-bar__sep"):
        return "w.BarSep{" + ("Secondary: true" if k.has("is-2nd") else "") + "}"
    if k.has("tc-seg"):
        opts = ["{Label: " + q(o.text()) + (", On: true" if o.has("is-on") else "") + "}" for o in k.els()]
        return "w.Seg{Options: []w.SegOption{" + ", ".join(opts) + "}" + (", Secondary: true" if k.has("is-2nd") else "") + "}"
    if k.has("tc-range"):
        ins = [i for i in k.els() if i.tag == "input"]
        def inp(i):
            kind = "Kind: w.InputDate, " if i.attrs.get("type") == "date" else ""
            return "w.Input{" + kind + "Value: " + q(i.attrs.get("value", "")) + ', Mono: true, Label: "Date"}'
        return "w.Range{From: " + inp(ins[0]) + ", To: " + inp(ins[1]) + ', Show: w.Button{Label: "Show"}' + (", Secondary: true" if k.has("is-2nd") else "") + "}"
    if k.has("tc-check"):
        tick = kid(k, "tc-tickrow__tick")
        return "w.Check{" + ("On: true, " if not tick.has("is-off") else "") + "Parts: []w.CheckPart{w.Text(" + q(k.own_text()) + ")}}"
    if k.tag == "input" and k.has("tc-in"):
        v = {"tc-in--w190": "InputW190", "tc-in--h26-w200": "InputH26W200"}
        var = [x for c, x in v.items() if k.has(c)]
        return "w.Input{Kind: w.InputSearch, Placeholder: " + q(k.attrs.get("placeholder", "")) + (", Variant: w." + var[0] if var else "") + ", Label: " + q(k.attrs.get("aria-label", "")) + "}"
    if k.has("tc-in") and k.tag == "label":
        qq = kid(k, "tc-q")
        return "w.Filter{Query: w.Query{Value: " + q(qq.attrs.get("value", "")) + ", Label: " + q(qq.attrs.get("aria-label", "")) + "}, Waiting: true}"
    if k.has("tc-mono"):
        v = ", Variant: w.MonoMuted" if k.has("tc-mono--muted") else ""
        return mono(k)
    if k.has("tc-btn"):
        return button(k)
    raise SystemExit("bar? " + str(k.cls))


def bar(e):
    f, items = [], []
    for k in e.els():
        if k.has("tc-drop"):
            return None
        if k.has("tc-bar__count"):
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


def hstate(cls, e, expr):
    if cls == "tc-bar":
        for k, v in [("Variant: w.BarSelection", "selection"), ("Variant: w.BarWide", "wide"), ("More:", "more"), ("Right:", "right")]:
            if k in expr:
                return v
        return "plain"
    for k, v in [("Meta:", "meta"), ("Photo:", "photo"), ("Avatar:", "avatar"), ("Crumbs:", "crumbs"), ("Middle:", "middle"),
                 ("Pad:", "padded"), ("Acts:", "acts")]:
        if k in expr:
            return v
    return "plain"


out, waiting = [], []
for cls, fn in [("tc-head", head), ("tc-bar", bar)]:
    out.append("\n\t// " + cls)
    for frame, n, e in shapes(cls):
        expr = fn(e)
        if expr is None:
            waiting.append("ref(%s, %s, %d)" % (q(frame), q(cls), n))
            continue
        out.append("\tshell(%s, %s, %s, %s, %d)," % (q(hstate(cls, e, expr)), expr, q(frame), q(cls), n))
print("\n".join(out))
print("\n// WAIT " + ", ".join(waiting))
