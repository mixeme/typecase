# Overrides of the record generators for whole screens (plan 3, step 1):
# every variant and every child as drawn. exec'd into genscreens' globals
# after the chain, so the chain's calls reach these definitions.
exec(open(os.path.join(HERE, "enums.py"), encoding="utf-8").read().split('\nif __name__ == "__main__":')[0])


def ev(e, t, field):
    """", Field: w.Const" for the enum t matching e's classes, or ""."""
    c = enum_for(e, t)
    return ", " + field + ": " + c if c else ""


def mark(e):
    f = []
    c = e.attrs.get("data-tc-colour")
    if c:
        f.append("Colour: " + q(c))
    k = enum_for(e, "MarkKind")
    if k:
        f.append("Kind: " + k)
    s = enum_for(e, "MarkSize")
    if s:
        f.append("Size: " + s)
    if e.has("is-focus"):
        f.append("Focus: true")
    return "w.Mark{" + ", ".join(f) + "}"


_place = {"tc-push": "Pushed", "tc-grow": "Grown"}


def placef(e):
    for c, v in _place.items():
        if e.has(c):
            return ", Place: w." + v
    return ""


def mono(e):
    f = "Text: " + q(e.text())
    if e.tag == "code":
        f += ", Form: w.MonoCode"
    elif e.tag == "div":
        f += ", Form: w.MonoLine"
    return "w.Mono{" + f + ev(e, "MonoVariant", "Variant") + placef(e) + "}"


def badge(e):
    f = ["Label: " + q(e.text())]
    t = enum_for(e, "BadgeTone")
    if t:
        f.append("Tone: " + t)
    svg = [k for k in e.els() if k.tag == "svg"]
    if svg:
        f.append("Icon: " + icon(svg[0]))
    if e.has("tc-badge--edge-accent-accent"):
        f.append("Accent: true")
    if e.attrs.get("title"):
        f.append("Title: " + q(e.attrs["title"]))
    return "w.Badge{" + ", ".join(f) + "}" if not placef(e) else "w.Badge{" + ", ".join(f) + placef(e) + "}"


_button0 = button


def button(e):
    if e.has("tc-close"):
        lab = e.attrs.get("aria-label") or e.attrs.get("title") or "Close"
        return "w.Close{Label: " + q(lab) + placef(e) + "}"
    s = _button0(e)
    extra = []
    v = enum_for(e, "ButtonVariant")
    if v and "Variant:" not in s:
        extra.append("Variant: " + v)
    if e.has("is-off"):
        extra.append("State: w.ButtonOff")
    if e.attrs.get("disabled") is not None:
        extra.append("State: w.ButtonDisabled")
    if e.has("is-3rd"):
        extra.append("Tertiary: true")
    b = [k for k in e.els() if k.tag == "b"]
    if b and not e.own_text():
        s = s.replace("w.Button{", "w.Button{Label: " + q(b[0].text()) + ", Bold: true, ", 1).replace('Title: "More", ', "").replace(', Title: "More"', "")
    if extra:
        s = s[:-1] + ", " + ", ".join(extra) + "}"
    return s


def rubric_item(x):
    if x.has("tc-mono"):
        return mono(x)
    if x.has("tc-badge"):
        return badge(x)
    if x.has("tc-num"):
        return "w.Num{Value: " + x.text() + (", Alert: true" if x.has("is-alert") else "") + "}"
    if x.has("tc-push"):
        return "w.Push{Items: []w.PushPart{" + ", ".join(push_item(y) for y in x.els()) + ("w.Text(" + q(x.own_text()) + ")" if x.own_text() and not x.els() else "") + "}}"
    if x.has("tc-rubric__aside"):
        its = [badge(y) if y.has("tc-badge") else mono(y) for y in x.els()]
        return "w.RubricAside{Text: " + q(x.own_text()) + (", Items: []w.RubricAsidePart{" + ", ".join(its) + "}" if its else "") + "}"
    if x.has("tc-rubric__col"):
        return "w.RubricCol{Text: " + q(x.text()) + "}"
    if x.has("tc-btn"):
        return button(x)
    if x.has("tc-plain"):
        return "w.Plain{Text: " + q(x.text()) + "}"
    if x.has("tc-line"):
        return linew2(x)
    raise SystemExit("rubric item? " + str(x.cls))


def push_item(y):
    if y.has("tc-btn"):
        return button(y)
    return child(y)


def rubric(e, form=None):
    f = []
    kids = [k for k in e.kids if not (isinstance(k, str) and not k.strip())]
    lead = None
    if kids and not isinstance(kids[0], str) and kids[0].has("tc-mono") and any(isinstance(k, str) for k in kids):
        lead = kids[0]
    f.append("Text: " + q(e.own_text()))
    if lead is not None:
        f.append("Lead: &" + mono(lead))
    items = [rubric_item(x) for x in e.els() if x is not lead]
    if items:
        f.append("Items: []w.RubricPart{" + ", ".join(items) + "}")
    if form is None:
        form = {"span": "RubricInline", "button": "RubricItem"}.get(e.tag)
    if form:
        f.append("Form: w." + form)
    v = enum_for(e, "RubricVariant")
    if v:
        f.append("Variant: " + v)
    return "w.Rubric{" + ", ".join(f) + "}"


def rubric_full(k, form=""):
    return rubric(k, form or None)


def listgroup(k):
    r = kid(k, "tc-rubric")
    n = kid(k, "tc-num")
    f = ["Title: " + rubric(r)]
    if n is not None:
        f.append("Count: &w.Num{Value: " + n.text() + (", Alert: true" if n.has("is-alert") else "") + "}")
    return "w.ListGroup{" + ", ".join(f) + "}"


def spaced(e, k):
    """A text piece of e as drawn: runs of white space are one space, and a
    space on a side where an element stands next to it is kept — the browser
    draws it; on the element's own edge it collapses and is dropped. A piece
    that is only a space is kept between two elements. None when nothing is
    left."""
    i = e.kids.index(k)
    before = any(not isinstance(x, str) for x in e.kids[:i])
    after = any(not isinstance(x, str) for x in e.kids[i + 1:])
    t = re.sub(r"\s+", " ", k)
    if not before:
        t = t.lstrip()
    if not after:
        t = t.rstrip()
    if t.strip() == "" and not (before and after):
        return None
    return t or None


def textpart(e, k):
    t = spaced(e, k)
    return "w.Text(" + q(t) + ")" if t is not None else None


def dimparts(e):
    parts = []
    for k in e.kids:
        if isinstance(k, str):
            if textpart(e, k):
                parts.append(textpart(e, k))
        elif k.has("tc-mono"):
            parts.append(mono(k))
        elif k.has("tc-dim__hint"):
            parts.append("w.Hint{Text: " + q(k.text()) + (", Title: " + q(k.attrs["title"]) if k.attrs.get("title") else "") + "}")
        elif k.tag == "b":
            parts.append("w.Strong(" + q(k.text()) + ")")
        else:
            raise SystemExit("dim part? " + str(k.cls))
    return parts


def dimx(e):
    return "w.Dim{Parts: []w.Inline{" + ", ".join(dimparts(e)) + "}" + ev(e, "DimSize", "Size") + ev(e, "DimVariant", "Variant") + placef(e) + "}"


def linew2(e):
    raise SystemExit("line? " + str(e.cls))


# ---- running text -------------------------------------------------------

_faces = {c: n for n, c in ENUMS.get("Face", [])}


def inline_el(k):
    if k.tag == "b" and k.has("tc-dim--accent"):
        return "w.Emphasis(" + q(k.text()) + ")"
    if k.tag == "b":
        if k.els():
            return "w.Bold{Parts: []w.Inline{" + ", ".join(inlparts(k)) + "}}"
        return "w.Strong(" + q(k.text()) + ")"
    if k.tag == "a" and k.has("tc-dim--accent"):
        return "w.Link{Text: " + q(k.text()) + ", Href: " + q(k.attrs.get("href", "#")) + "}"
    if k.has("tc-mono") or k.tag == "code":
        return mono(k)
    if k.has("tc-dim__hint"):
        return "w.Hint{Text: " + q(k.text()) + (", Title: " + q(k.attrs["title"]) if k.attrs.get("title") else "") + "}"
    for c in k.cls:
        if c.startswith("tc-dot--"):
            return "w.Dot{Tone: w.Dot" + {"accent": "Accent", "danger": "Danger", "dusty": "Dusty"}[c[8:]] + "}"
    cl = " ".join(k.cls)
    if k.tag == "span" and cl in _faces:
        return "w.Span{Face: w." + _faces[cl] + ", Text: " + q(k.text()) + "}"
    raise SystemExit("inline? " + k.tag + " " + cl)


def inlparts(e, el=inline_el):
    out = []
    for k in e.kids:
        if isinstance(k, str):
            if textpart(e, k):
                out.append(textpart(e, k))
        else:
            out.append(el(k))
    return out


def inl(e):
    """Running text as []w.Inline; badges and cells belong to the row or the
    card that holds the text, not to it."""
    def el(k):
        return None if k.has("tc-badge") or k.has("tc-row__cell") else inline_el(k)
    return "[]w.Inline{" + ", ".join(p for p in inlparts(e, el) if p) + "}"


# ---- controls -----------------------------------------------------------

def placesec(e):
    return placef(e) + (", Secondary: true" if e.has("is-2nd") else "")


def inputw(x):
    f = []
    kind = {"date": "InputDate", "time": "InputTime", "search": "InputSearch"}.get(x.attrs.get("type", "text"))
    if kind:
        f.append("Kind: w." + kind)
    for a, n in (("name", "Name"), ("value", "Value"), ("placeholder", "Placeholder"), ("aria-label", "Label")):
        if x.attrs.get(a):
            f.append(n + ": " + q(x.attrs[a]))
    for c, n in (("is-mono", "Mono"), ("is-focus", "Focus"), ("is-error", "Error")):
        if x.has(c):
            f.append(n + ": true")
    if x.attrs.get("inputmode") == "numeric":
        f.append("Numeric: true")
    return "w.Input{" + ", ".join(f) + ev(x, "InputVariant", "Variant") + placef(x) + "}"


def pickw(x):
    f = ["Value: " + q(x.own_text())]
    d = kid(x, "tc-dim")
    if d is not None:
        f.append("Aside: &" + dimx(d))
    if x.has("is-ph"):
        f.append("Unset: true")
    if x.attrs.get("aria-label"):
        f.append("Label: " + q(x.attrs["aria-label"]))
    if x.has("is-focus"):
        f.append("Focus: true")
    return "w.Pick{" + ", ".join(f) + ev(x, "PickVariant", "Variant") + placef(x) + "}"


def filterw(x):
    qq = kid(x, "tc-q")
    f = ["Query: w.Query{Value: " + q(qq.attrs.get("value", "")) + ", Label: " + q(qq.attrs.get("aria-label", "")) + "}"]
    if x.has("is-ph"):
        f.append("Waiting: true")
    svg = [k for k in x.els() if k.tag == "svg"]
    if svg and svg[0].has("tc-ico--dusty"):
        f.append("Quiet: true")
    return "w.Filter{" + ", ".join(f) + ev(x, "FilterVariant", "Variant") + placef(x) + "}"


def taginputw(x):
    parts = []
    for k in x.els():
        if k.has("tc-tag"):
            parts.append(tag(k))
        elif k.has("tc-dim"):
            parts.append(dimx(k))
        else:
            parts.append(inline_el(k))
    f = ["Parts: []w.TagInputPart{" + ", ".join(parts) + "}"]
    if x.has("tc-in--wrap-gap5"):
        f.append("Wrap: true")
    if x.has("is-focus"):
        f.append("Focus: true")
    return "w.TagInput{" + ", ".join(f) + "}"


def areaw(x):
    f = ["Text: " + q(x.text() if not pre_text(x) else pre_text(x))]
    v = enum_for(x, "AreaVariant")
    if v and v != "w.AreaPlain":
        f.append("Variant: " + v)
    for c, n in (("is-mono", "Mono"), ("is-focus", "Focus")):
        if x.has(c):
            f.append(n + ": true")
    for a, n in (("name", "Name"), ("aria-label", "Label")):
        if x.attrs.get(a):
            f.append(n + ": " + q(x.attrs[a]))
    return "w.Area{" + ", ".join(f) + "}"


def pre_text(x):
    return "".join(k if isinstance(k, str) else "" for k in x.kids).strip("\n") if x.tag == "textarea" else ""


def segw(x):
    opts = []
    for o in x.els():
        f = []
        if o.own_text():
            f.append("Label: " + q(o.own_text()))
        svg = [k for k in o.els() if k.tag == "svg"]
        if svg:
            f.append("Icon: " + icon(svg[0]))
        if o.has("is-on"):
            f.append("On: true")
        if o.attrs.get("aria-label"):
            f.append("Title: " + q(o.attrs["aria-label"]))
        opts.append("{" + ", ".join(f) + "}")
    return "w.Seg{Options: []w.SegOption{" + ", ".join(opts) + "}" + ev(x, "SegVariant", "Variant") + placesec(x) + "}"


def checkpart(k):
    if k.has("tc-dim"):
        return dimx(k)
    if k.has("tc-mono"):
        return mono(k)
    if k.has("tc-badge"):
        return badge(k)
    if k.has("tc-diff__value"):
        return "w.Choice{Source: " + q(kid(k, "tc-diff__source").text()) + ", Value: " + q(k.own_text()) + "}"
    raise SystemExit("check part? " + str(k.cls))


def checkw(x):
    tick = kid(x, "tc-tickrow__tick")
    parts = []
    for k in x.kids:
        if isinstance(k, str):
            if textpart(x, k):
                parts.append(textpart(x, k))
        elif k is not tick:
            parts.append(checkpart(k))
    f = []
    if tick is not None and not tick.has("is-off"):
        f.append("On: true")
    f.append("Parts: []w.CheckPart{" + ", ".join(parts) + "}")
    if x.attrs.get("aria-label"):
        f.append("Label: " + q(x.attrs["aria-label"]))
    if x.has("is-3rd"):
        f.append("Tertiary: true")
    return "w.Check{" + ", ".join(f) + "}"


def control(x):
    """Any control or field element: the one dispatcher fields, lines and
    forms share."""
    if x.tag == "input" and x.has("tc-in"):
        return inputw(x)
    if x.tag == "button" and x.has("tc-in") and x.has("is-pick"):
        return pickw(x)
    if x.tag == "label" and x.has("tc-in"):
        return filterw(x)
    if x.tag == "span" and x.has("tc-in"):
        return taginputw(x)
    if x.tag == "textarea":
        return areaw(x)
    if x.has("tc-seg"):
        return segw(x)
    if x.has("tc-check"):
        return checkw(x)
    if x.has("tc-btn"):
        return button(x)
    return None
