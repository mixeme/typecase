"""Emit refs records for the rows and lists family from the shapes tool."""
import re, subprocess, os
from html.parser import HTMLParser

SHAPES = os.path.join(os.environ["TMP"], "shapes.exe")
VOID = {"input", "br", "img", "use", "meta", "link", "hr"}


class El:
    def __init__(s, tag, attrs):
        s.tag, s.attrs, s.kids = tag, dict(attrs), []

    @property
    def cls(s):
        return (s.attrs.get("class") or "").split()

    def has(s, c):
        return c in s.cls

    def els(s):
        return [k for k in s.kids if isinstance(k, El)]

    def text(s):
        return norm("".join(k if isinstance(k, str) else k.text() for k in s.kids))

    def own_text(s):
        return norm(" ".join(k for k in s.kids if isinstance(k, str)))


class P(HTMLParser):
    def __init__(s):
        super().__init__(convert_charrefs=True)
        s.root = El("root", {})
        s.st = [s.root]

    def handle_starttag(s, t, a):
        e = El(t, a)
        s.st[-1].kids.append(e)
        if t not in VOID:
            s.st.append(e)

    def handle_startendtag(s, t, a):
        s.st[-1].kids.append(El(t, a))

    def handle_endtag(s, t):
        if t in VOID:
            return
        while s.st and s.st[-1].tag != t:
            s.st.pop()
        if s.st:
            s.st.pop()

    def handle_data(s, d):
        s.st[-1].kids.append(d)


def parse(h):
    p = P()
    p.feed(h)
    return p.root.els()[0]


def shapes(cls):
    out = subprocess.run([SHAPES, cls, "-html"], capture_output=True, encoding="utf-8").stdout
    res = []
    lines = out.split("\n")
    for i, ln in enumerate(lines):
        m = re.match(r"^(\d+)\. ×(\d+)\s+first (\S+) \.(\S+) #(\d+)", ln)
        if m:
            res.append((m.group(3), int(m.group(5)), parse(lines[i + 2].strip())))
    return res


def q(s):
    return '"' + s.replace("\\", "\\\\").replace('"', '\\"').replace("\n", "\\n").replace("\t", "\\t") + '"'


def norm(s):
    return " ".join(s.split())


def mark(e):
    f = []
    c = e.attrs.get("data-tc-colour")
    if c:
        f.append("Colour: " + q(c))
    if e.has("is-multi"):
        f.append("Kind: w.MarkMulti")
    if e.has("is-ext"):
        f.append("Kind: w.MarkExternal")
    for k, v in {"tc-mark--h32": "MarkH32", "tc-mark--h34": "MarkH34", "tc-mark--on-hair": "MarkOnHair", "tc-mark--h11": "MarkH11"}.items():
        if e.has(k):
            f.append("Size: w." + v)
    return "w.Mark{" + ", ".join(f) + "}"


def icon(e):
    u = [k for k in e.els() if k.tag == "use"][0]
    return q(u.attrs["href"][3:])


def badge(e):
    f = ["Label: " + q(e.text())]
    for k, v in {"is-linked": "BadgeLinked", "is-alert": "BadgeAlert", "is-local": "BadgeLocal"}.items():
        if e.has(k):
            f.append("Tone: w." + v)
    svg = [k for k in e.els() if k.tag == "svg"]
    if svg:
        f.append("Icon: " + icon(svg[0]))
    if e.attrs.get("title"):
        f.append("Title: " + q(e.attrs["title"]))
    return "w.Badge{" + ", ".join(f) + "}"


def tag(e):
    f = ["Label: " + q(e.text())]
    if e.has("is-on"):
        f.append("On: true")
    return "w.Tag{" + ", ".join(f) + "}"


def inl(e):
    """Running text of an element as []w.Inline, badges and cells left out."""
    parts = []
    for k in e.kids:
        if isinstance(k, str):
            t = re.sub(r"\s+", " ", k)
            if t == "":
                continue
            parts.append(("t", t))
        elif k.tag == "b":
            parts.append(("x", "w.Strong(" + q(k.text()) + ")"))
        elif k.has("tc-dim--accent"):
            parts.append(("x", "w.Span{Face: w.FaceAccent, Text: " + q(k.text()) + "}"))
        elif k.has("tc-mono"):
            parts.append(("x", mono(k)))
        elif k.has("tc-badge") or k.has("tc-row__cell"):
            continue
        else:
            raise SystemExit("inline? " + k.tag + str(k.cls))
    if parts and parts[0][0] == "t":
        parts[0] = ("t", parts[0][1].lstrip())
    if parts and parts[-1][0] == "t":
        parts[-1] = ("t", parts[-1][1].rstrip())
    out = []
    for kind, v in parts:
        if kind == "t":
            if v:
                out.append("w.Text(" + q(v) + ")")
        else:
            out.append(v)
    return "[]w.Inline{" + ", ".join(out) + "}"


def kid(e, c):
    for k in e.els():
        if k.has(c):
            return k
    return None


def pushed(e):
    return e.has("tc-push")


def tickrow(e):
    f = []
    tick = e.els()[0]
    if tick.has("is-part"):
        f.append("Ticked: w.TickedPart")
    if tick.has("is-off"):
        f.append("Ticked: w.TickedOff")
    for k, v in {"is-all": "TickRowAll", "is-root": "TickRowRoot", "is-ext": "TickRowExternal"}.items():
        if e.has(k):
            f.append("Kind: w." + v)
    for k, v in {"tc-tickrow--padl0": "TickRowGroup", "tc-tickrow--padl14": "TickRowNested", "tc-tickrow--h44": "TickRowTouch"}.items():
        if e.has(k):
            f.append("Form: w." + v)
    lead = e.els()[1]
    if lead.tag == "svg":
        f.append("Icon: " + icon(lead))
    else:
        f.append("Mark: " + mark(lead))
    f.append("Label: " + q(kid(e, "tc-tickrow__label").text()))
    meta = kid(e, "tc-tickrow__meta")
    if meta is not None:
        f.append("Meta: " + q(meta.text()))
        if meta.attrs.get("title"):
            f.append("MetaTitle: " + q(meta.attrs["title"]))
    if e.has("is-on"):
        f.append("Current: true")
    if e.has("is-off"):
        f.append("Hidden: true")
    if e.has("is-error"):
        f.append("Error: true")
    last = e.els()[-1]
    if last.has("tc-badge"):
        f.append("Trail: " + badge(last))
    if last.has("tc-btn"):
        f.append("Trail: w.Button{Label: " + q(last.text()) + ", Small: true}")
    if last.has("tc-skel"):
        f.append("Trail: w.Skel{Width: w.Skel16}")
    return "w.TickRow{" + ", ".join(f) + "}"


def tree(e):
    f, rows, on = [], [], 0
    for k, v in {"tc-tree--padx8": "TreePadded", "tc-tree--max520": "TreeNarrow"}.items():
        if e.has(k):
            f.append("Variant: w." + v)
    for i, r in enumerate(e.els()):
        g = []
        if r.has("is-on"):
            on = i + 1
        els = r.els()
        lead = els[0]
        if lead.tag == "svg":
            g.append("Icon: " + icon(lead))
        elif lead.has("tc-mark"):
            g.append("Mark: &" + mark(lead))
        mono = [k for k in els if k.has("tc-mono") and not pushed(k)]
        if mono:
            g.append("Label: " + q(mono[0].text()))
            g.append("Code: true")
        else:
            g.append("Label: " + q(r.own_text()))
        for k, v in {"tc-tree__row--padl1-25rem": "TreeLevel20", "tc-tree__row--padl22": "TreeLevel22", "tc-tree__row--padl36": "TreeLevel36",
                     "tc-tree__row--padl40": "TreeLevel40", "tc-tree__row--padl44": "TreeLevel44"}.items():
            if r.has(k):
                g.append("Indent: w." + v)
        if r.has("tc-tree__row--op-55"):
            g.append("Unavailable: true")
        tr = [k for k in els if pushed(k)]
        if tr:
            t = tr[0]
            if t.has("tc-tickrow__meta"):
                g.append("Trail: w.TreeNote{Text: " + q(t.text()) + "}")
            elif t.has("tc-dim"):
                g.append("Trail: w.Dim{Parts: []w.Inline{w.Text(" + q(t.text()) + ")}, Size: w.Dim11}")
            elif t.has("tc-mono"):
                g.append("Trail: w.Mono{Text: " + q(t.text()) + ", Variant: w.Mono10Dusty}")
            elif t.has("tc-badge"):
                g.append("Trail: " + badge(t))
        rows.append("{" + ", ".join(g) + "}")
    out = ["Rows: []w.TreeRow{\n\t\t" + ",\n\t\t".join(rows) + "}"]
    if on:
        out.append("On: %d" % on)
    return "w.Tree{" + ", ".join(out + f) + "}"


def title_fields(t, name="Title"):
    f = [name + ": " + inl(t)]
    if t.has("is-strong"):
        f.append("Strong: true")
    b = kid(t, "tc-badge")
    if b is not None:
        f.append(name + "Badge: &" + badge(b))
    return f


def src(e, f):
    s = kid(e, "tc-row__source")
    if s is None:
        return
    t = s.own_text()
    if t:
        f.append("Source: " + q(t))
    if s.attrs.get("title"):
        f.append("SourceHint: " + q(s.attrs["title"]))
    b = kid(s, "tc-badge")
    if b is not None:
        f.append("SourceBadge: &" + badge(b))


def tagsf(e, f):
    t = kid(e, "tc-row__tags")
    items = [tag(k) if k.has("tc-tag") else badge(k) for k in t.els()]
    if items:
        f.append("Tags: []w.RowTag{" + ", ".join(items) + "}")


def row(e):
    f = []
    on = e.has("is-on")
    els = e.els()
    if e.has("tc-row--timecols") and kid(e, "tc-skel") is not None:
        r = {"tc-row__detail--w40pc": 1, "tc-row__detail--w30pc": 2, "tc-row__detail--w25pc": 3}
        n = [r[c] for c in kid(e, "tc-row__detail").cls if c in r][0]
        return "w.SkelRow{Title: w.Skel{Width: w.Skel%s}}" % {1: "60", 2: "45", 3: "52"}[n]
    if e.has("tc-row--timecols"):
        f.append("Mark: " + mark(els[0]))
        f.append("Time: " + q(kid(e, "tc-row__lead").text()))
        f += title_fields(kid(e, "tc-row__title"))
        f.append("Detail: " + q(kid(e, "tc-row__detail").text()))
        c = kid(e, "tc-row__caption")
        if c.els():
            f.append("Caption: &" + badge(c.els()[0]))
        src(e, f)
        if on:
            f.append("On: true")
        if e.has("tc-row--timecols--op-7"):
            f.append("Past: true")
        if e.has("tc-row--timecols--padl14-padr14"):
            f.append("Inset: true")
        return "w.TimeRow{" + ", ".join(f) + "}"
    if e.has("tc-row--avatarcols"):
        f.append("Mark: " + mark(els[0]))
        a = kid(e, "tc-avatar")
        f.append("Avatar: " + ("w.Avatar{Photo: true}" if a.has("is-photo") else "w.Avatar{Initials: " + q(a.text()) + "}"))
        f += title_fields(kid(e, "tc-row__title"))
        f.append("Subtitle: " + q(kid(e, "tc-row__subtitle").text()))
        f.append("Detail: " + q(kid(e, "tc-row__detail").text()))
        f.append("Caption: " + q(kid(e, "tc-row__caption").text()))
        src(e, f)
        if on:
            f.append("On: true")
        if e.has("tc-row--avatarcols--padl14-padr14"):
            f.append("Inset: true")
        return "w.FaceRow{" + ", ".join(f) + "}"
    if e.has("tc-row--find"):
        lead = kid(e, "tc-row__lead")
        f.append("Mark: " + mark(els[0]))
        if lead.tag == "svg":
            f.append("Icon: " + icon(lead))
            f.append("Title: " + q(kid(e, "tc-row__title").text()))
            f.append("Detail: " + q(kid(e, "tc-row__detail").text()))
            return "w.LinkedRow{" + ", ".join(f) + "}"
        f.append("Date: " + q(lead.text()))
        f.append("Title: " + inl(kid(e, "tc-row__title")))
        f.append("Detail: " + inl(kid(e, "tc-row__detail")))
        f.append("What: " + q(kid(e, "tc-row__caption").text()))
        src(e, f)
        return "w.FindRow{" + ", ".join(f) + "}"
    if e.has("tc-row--excerpt--cols3-62-148"):
        f.append("Mark: " + mark(els[0]))
        f.append("Title: " + q(kid(e, "tc-row__title").text()))
        f.append("Date: w.Mono{Text: " + q(kid(e, "tc-mono").text()) + "}")
        f.append("Where: " + q(kid(e, "tc-row__tags").text()))
        if on:
            f.append("On: true")
        return "w.BriefRow{" + ", ".join(f) + "}"
    if e.has("tc-row--excerpt"):
        f.append("Mark: " + mark(els[0]))
        f += title_fields(kid(e, "tc-row__title"))
        f.append("Excerpt: " + q(kid(e, "tc-row__excerpt").text()))
        tagsf(e, f)
        f.append("Date: " + q(kid(e, "tc-row__date").text()))
        src(e, f)
        if on:
            f.append("On: true")
        if e.has("tc-row--excerpt--padl14-padr14"):
            f.append("Inset: true")
        return "w.ExcerptRow{" + ", ".join(f) + "}"
    if e.has("tc-row--tickmarkcols"):
        tk = els[0]
        f.append("Tick: w.Tick{" + ("Done: true, " if tk.has("is-done") else "") + 'Label: "Done"}')
        f.append("Mark: " + mark(els[1]))
        f += title_fields(kid(e, "tc-row__title"))
        tagsf(e, f)
        f.append("Due: " + q(kid(e, "tc-row__due").text()))
        f.append("Progress: " + q(kid(e, "tc-row__progress").text()))
        src(e, f)
        if on:
            f.append("On: true")
        if e.has("is-overdue"):
            f.append("Overdue: true")
        if e.has("tc-row--tickmarkcols--padl14-padr14"):
            f.append("Inset: true")
        return "w.DueRow{" + ", ".join(f) + "}"
    if e.has("tc-row--tickcols"):
        if els[0].has("is-off"):
            f.append("Ticked: w.TickedOff")
        name = kid(e, "tc-row__name")
        f.append("Icon: " + icon(name.els()[0]))
        f.append("Name: " + inl(kid(name, "tc-row__file")))
        f.append("Bytes: " + q(kid(e, "tc-row__size").text()))
        f.append("Time: " + q(kid(e, "tc-row__time").text()))
        f.append("Access: " + q(kid(e, "tc-row__detail").text()))
        src(e, f)
        return "w.SizeRow{" + ", ".join(f) + "}"
    if e.has("tc-row--slim"):
        if els[0].has("is-off"):
            f.append("Ticked: w.TickedOff")
        f.append("Mark: " + mark(els[1]))
        t = kid(e, "tc-row__title")
        f.append("Title: " + q(t.own_text()))
        f.append("Inside: " + q(kid(t, "tc-row__cell").text()))
        f.append("Aside: " + q(els[-1].text()))
        return "w.SlimRow{" + ", ".join(f) + "}"
    if e.has("tc-row--phone"):
        first = els[0]
        if first.has("tc-tickrow__tick"):
            f.append("Pick: &w.TickBox{" + ("Ticked: w.TickedOff" if first.has("is-off") else "") + "}")
        else:
            f.append("Mark: &" + mark(first))
        outer = els[1]
        text = outer
        if outer.has("tc-row__text--row"):
            lead = outer.els()[0]
            if lead.tag == "svg":
                f.append("Icon: " + icon(lead))
            else:
                a = "w.Avatar{Photo: true, Size: w.Avatar34}" if lead.has("is-photo") else "w.Avatar{Initials: " + q(lead.text()) + ", Size: w.Avatar34}"
                f.append("Avatar: &" + a)
            text = outer.els()[1]
        t = kid(text, "tc-row__title")
        f.append("Title: " + q(t.own_text()))
        b = kid(t, "tc-badge")
        if b is not None:
            f.append("TitleBadge: &" + badge(b))
        sub = kid(text, "tc-row__cell")
        f.append("Sub: " + inl(sub))
        if sub.has("tc-mono"):
            f.append("SubMono: true")
        last = els[-1]
        if last.has("tc-btn"):
            f.append('More: &w.Button{Icon: "dots", IconBox: w.IconBox16, Variant: w.ButtonPadX4, Title: "More"}')
        else:
            f.append("Aside: " + q(last.text()))
        if on:
            f.append("On: true")
        return "w.PhoneRow{" + ", ".join(f) + "}"
    raise SystemExit("row? " + str(e.cls))


def lst(e):
    rows = [row(r) for r in e.els()]
    f = ["Rows: []w.ListRow{\n\t\t" + ",\n\t\t".join(rows) + "}"]
    if e.has("tc-list--mx-14"):
        f.append("Inset: true")
    return "w.List{" + ", ".join(f) + "}"


def skel(e):
    for k, v in {"tc-skel--w16": "Skel16", "tc-skel--w45pc": "Skel45", "tc-skel--w52pc": "Skel52", "tc-skel--w60pc": "Skel60"}.items():
        if e.has(k):
            return "w.Skel{Width: w." + v + "}"
    return "w.Skel{}"


def state(cls, e):
    c = set(e.cls)
    if cls == "tc-tickrow":
        for k, v in [("is-error", "error"), ("is-root", "root"), ("is-all", "all"), ("is-ext", "external"),
                     ("tc-tickrow--padl0", "group"), ("tc-tickrow--padl14", "nested"), ("tc-tickrow--h44", "touch"),
                     ("is-on", "current"), ("is-off", "hidden")]:
            if k in c:
                return v
        return "loading" if kid(e, "tc-skel") is not None else "one"
    if cls == "tc-skel":
        return "sized" if any(x.startswith("tc-skel--") for x in c) else "fill"
    if cls == "tc-list":
        return "inset" if "tc-list--mx-14" in c else "plain"
    if cls == "tc-tree":
        for k, v in [("tc-tree--padx8", "padded"), ("tc-tree--max520", "narrow")]:
            if k in c:
                return v
        return "plain"
    return "plain"


def row_state(expr, e):
    if expr.startswith("w.SkelRow"):
        return "loading"
    checks = [
        (e.has("is-on"), "on"), (e.has("is-overdue"), "overdue"), (e.has("is-done"), "done"),
        ("Inset: true" in expr, "inset"), ("Past: true" in expr, "past"),
        ("w.Strong(" in expr or "w.FaceAccent" in expr, "found"), ("TitleBadge" in expr, "badge"),
        ("Caption: &" in expr, "caption"), ("Pick:" in expr, "pick"), ("Ticked: w.TickedOff" in expr, "off"),
        ("Strong: true" in expr, "strong"), ("Avatar: &" in expr, "avatar"),
        (("w.ExcerptRow" in expr or "w.DueRow" in expr) and "Tags:" not in expr, "untagged"),
    ]
    for ok, name in checks:
        if ok:
            return name
    return "plain"


FAM = [("tc-tickrow", tickrow), ("tc-tree", tree), ("tc-list", lst), ("tc-row", row), ("tc-skel", skel)]
out = []
for cls, fn in FAM:
    out.append("\n\t// " + cls)
    for frame, n, e in shapes(cls):
        expr = fn(e)
        st = row_state(expr, e) if cls == "tc-row" else state(cls, e)
        out.append("\trows(%s, %s, %s, %s, %d)," % (q(st), expr, q(frame), q(cls), n))
print("\n".join(out))
