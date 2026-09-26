"""Emit refs records for the document family."""
import os

HERE = os.path.dirname(os.path.abspath(__file__))
src = open(os.path.join(HERE, "gensheet.py"), encoding="utf-8").read()
exec(src.split("\nout, waiting = [], []")[0])
exec(open(os.path.join(HERE, "rawlib.py"), encoding="utf-8").read())


def inl2(e):
    parts = []
    for k in e.kids:
        if isinstance(k, str):
            t = re.sub(r"\s+", " ", k)
            if t:
                parts.append(("t", t))
        elif k.tag == "b":
            parts.append(("x", "w.Strong(" + q(k.text()) + ")"))
        elif k.tag == "a":
            parts.append(("x", "w.Link{Text: " + q(k.text()) + ', Href: "#"}'))
        elif k.has("tc-mono"):
            v = ", Variant: w.MonoChip" if k.has("tc-mono--chip") else ""
            parts.append(("x", mono(k)))
        elif k.has("tc-dim--accent"):
            parts.append(("x", "w.Span{Face: w.FaceAccent, Text: " + q(k.text()) + "}"))
        else:
            raise SystemExit("inline? " + str(k.cls))
    if parts and parts[0][0] == "t":
        parts[0] = ("t", parts[0][1].lstrip())
    if parts and parts[-1][0] == "t":
        parts[-1] = ("t", parts[-1][1].rstrip())
    return "[]w.Inline{" + ", ".join(("w.Text(" + q(v) + ")") if k == "t" else v for k, v in parts if v) + "}"


def rawparts(e):
    parts = []
    for k in e.kids:
        if isinstance(k, str):
            if k:
                parts.append("w.Text(" + q(k) + ")")
        elif k.tag == "b":
            parts.append("w.Strong(" + q(k.text()) + ")")
        elif k.tag == "a":
            parts.append("w.Link{Text: " + q(k.text()) + ', Href: "#"}')
        elif k.has("tc-mono"):
            parts.append(mono(k))
        else:
            raise SystemExit("prose? " + str(k.cls))
    return "[]w.Inline{" + ", ".join(parts) + "}"


def prose(e):
    v = {("tc-prose--lead", "tc-prose--top14"): "ProseLead", ("tc-prose--lead", "tc-prose--top18"): "ProseLeadAfter",
         ("tc-prose--read", "tc-prose--top14"): "ProseReadTop", ("tc-prose--muted-flush-bot10",): "ProseMutedClose",
         ("tc-prose--t13-flush-padt11",): "ProseCard13", ("tc-prose--muted",): "ProseMuted", ("tc-prose--t13",): "Prose13",
         ("tc-prose--read",): "ProseRead", ("tc-prose--top14",): "ProseTop"}
    var = ""
    for cs, x in v.items():
        if all(e.has(c) for c in cs):
            var = ", Variant: w." + x.replace("ProseCard13", "Prose13Card")
            break
    return "w.Prose{Parts: " + rawparts(e) + var + "}"


def rub(k):
    v = {"tc-rubric--bot7": "RubricBottom7", "tc-rubric--bot8": "RubricBottom8"}
    var = [x for c, x in v.items() if k.has(c)]
    items = []
    for x in k.els():
        if x.has("tc-btn"):
            items.append(button(x))
        elif x.has("tc-badge"):
            items.append(badge(x))
    return "w.Rubric{Text: " + q(k.own_text()) + (", Items: []w.RubricPart{" + ", ".join(items) + "}" if items else "") + (", Variant: w." + var[0] if var else "") + "}"


def tasks(k, phone=False):
    items = []
    for li in kid(k, "tc-doc__list").els():
        done = kid(li, "tc-tick").has("is-done")
        items.append("{Tick: w.Tick{" + ("Done: true, " if done else "") + "Variant: w.TickNudge, Label: " + q(li.text()) + "}, Text: " + q(li.text()) + "}")
    return "w.DocTasks{Title: " + rub(kid(k, "tc-rubric")) + ", Items: []w.DocTask{" + ", ".join(items) + "}" + (", Phone: true" if phone else "") + "}"


def doc(e):
    items = []
    for k in e.els():
        if k.has("tc-doc__title"):
            items.append("w.DocTitle{Text: " + q(k.text()) + "}")
        elif k.has("tc-head__sub--top6"):
            ps = []
            for x in k.kids:
                if isinstance(x, str):
                    if x.strip():
                        ps.append("w.Text(" + q(re.sub(r"\s+", " ", x)) + ")")
                elif x.has("tc-mono"):
                    ps.append(mono(x))
                elif x.has("tc-tag"):
                    ps.append(tag(x))
            items.append("w.DocSub{Parts: []w.DocSubPart{" + ", ".join(ps) + "}}")
        elif k.has("tc-prose"):
            items.append(prose(k))
        elif k.has("tc-doc__sec"):
            items.append(tasks(k))
        elif k.has("tc-doc__acts"):
            items.append(tasks(k, True))
        elif k.has("tc-doc__sec--rule"):
            fs = []
            for it in kid(k, "tc-doc__files").els():
                fs.append("{Name: " + q(it.els()[0].text()) + ", Size: w.Dim{Parts: []w.Inline{w.Text(" + q(it.els()[1].text()) + ")}}}")
            items.append("w.DocFiles{Title: " + rub(kid(k, "tc-rubric")) + ", Items: []w.DocFile{" + ", ".join(fs) + "}}")
        elif k.has("tc-doc__sec--rule16"):
            l = kid(k, "tc-doc__link")
            items.append("w.DocRelated{Title: " + rub(kid(k, "tc-rubric")) + ", Name: " + q(l.own_text()) + ", What: w.Dim{Parts: []w.Inline{w.Text(" + q(kid(l, "tc-dim").text()) + ")}}}")
        elif k.tag == "textarea":
            if k.has("tc-in--title"):
                items.append(areaw(k))
            else:
                items.append(areaw(k))
        elif k.has("tc-doc__fields"):
            items.append("w.DocFields{Items: []w.Field{" + ", ".join(field(y) for y in k.els()) + "}}")
        elif k.has("tc-doc__note"):
            items.append("w.DocNote{Parts: " + inl2(k) + "}")
        elif k.has("tc-doc__heading"):
            items.append("w.DocHeading{Text: " + q(k.text()) + "}")
        elif k.has("tc-doc__bullets"):
            items.append("w.DocBullets{Items: [][]w.Inline{" + ", ".join(inl2(li) for li in k.els()) + "}}")
        elif k.has("tc-doc__body"):
            items.append("w.DocBody{Text: " + q(k.text()) + "}")
        else:
            raise SystemExit("doc? " + str(k.cls))
    f = ["Items: []w.DocPart{\n\t\t" + ",\n\t\t".join(items) + "}"]
    if e.has("tc-stack"):
        f.append("Text: true")
    return "w.Doc{" + ", ".join(f) + "}"


def docbar(e):
    items, f, end = [], [], []
    for k in e.els():
        if k.has("tc-bar__sep"):
            items.append("w.BarSep{}")
        elif k.has("tc-bar__more"):
            f.append("More: &" + button(k))
        elif k.has("tc-bar__right"):
            for x in k.els():
                if x.has("tc-check"):
                    end.append(baritem(x))
                elif x.has("tc-kbd"):
                    end.append("w.Kbd{Keys: " + q(x.text()) + (", Secondary: true" if x.has("is-2nd") else "") + "}")
                elif x.has("tc-seg"):
                    end.append(baritem(x))
                elif x.has("tc-btn"):
                    end.append(button(x))
                else:
                    end.append("w.DraftNote{Text: " + q(x.own_text()) + ", Time: w.Mono{Text: " + q(kid(x, "tc-mono").text()) + "}" + (", Secondary: true" if x.has("is-2nd") else "") + "}")
        elif k.has("tc-btn"):
            b = button(k)
            if kid(k, "b") is not None:
                b = 'w.Button{Label: "B", Bold: true}'
            if k.has("tc-btn--fontstyle-italic"):
                b = 'w.Button{Label: "I", Variant: w.ButtonItalic}'
            items.append(b)
    out = ["Items: []w.DocbarPart{" + ", ".join(items) + "}"] + f
    if end:
        out.append("End: []w.DocbarEnd{" + ", ".join(end) + "}")
    if e.has("tc-docbar--rule"):
        out.append("Rule: true")
    return "w.Docbar{" + ", ".join(out) + "}"


def docmeta(e):
    secs = []
    for k in e.els():
        if not k.has("tc-docmeta__sec"):
            raise Wait()
        its = []
        for x in k.els()[1:]:
            if x.has("tc-line") and not x.has("tc-line__text"):
                its.append("w.Line{Block: true, Gap: w.LineGap4, Items: []w.LinePart{" + ", ".join(tag(t) for t in x.els()) + "}}")
            elif x.has("tc-stack"):
                refs = []
                for it in x.els():
                    m = kid(it, "tc-mark")
                    a = [y for y in it.els() if y.tag == "a"]
                    name = a[0].text() if a else kid(it, "tc-dim--accent").text()
                    refs.append("{" + ("Mark: &" + mark(m) + ", " if m is not None else "") + "Name: " + q(name) + (', Href: "#"' if a else "") +
                                ", What: w.Dim{Parts: []w.Inline{w.Text(" + q(kid(it, "tc-dim").text()) + ")}}}")
                its.append("w.DocRefs{Items: []w.DocRef{" + ", ".join(refs) + "}}")
            elif x.has("tc-line__text"):
                its.append("w.DocAttach{Name: " + q(x.own_text()) + ", Size: w.Dim{Parts: []w.Inline{w.Text(" + q(kid(x, "tc-dim").text()) + ")}}, Where: " + badge(kid(x, "tc-badge")) + "}")
            else:
                raise Wait()
        secs.append("{Title: " + rub(k.els()[0]) + ", Items: []w.DocmetaPart{" + ", ".join(its) + "}}")
    return "w.Docmeta{Sections: []w.DocmetaSec{\n\t\t" + ",\n\t\t".join(secs) + "}}"


def dstate(cls, expr):
    if cls == "tc-prose":
        for k, v in [("Link{", "rich"), ("Strong(", "rich"), ("Mono{", "rich"), ("ProseMuted", "muted"), ("Prose13", "size"), ("ProseLead", "lead"),
                     ("ProseRead", "read"), ("ProseTop", "top")]:
            if k in expr:
                return v
        return "plain"
    if cls == "tc-doc":
        for k, v in [("Text: true", "text"), ("Phone: true", "phone"), ("w.Area", "edit"), ("DocFiles", "public")]:
            if k in expr:
                return v
        return "read"
    if cls == "tc-docbar":
        return "rule" if "Rule: true" in expr else "plain"
    return "read"


out, waiting = [], []
for cls, fn in [("tc-prose", prose), ("tc-doc", doc), ("tc-docbar", docbar), ("tc-docmeta", docmeta)]:
    out.append("\n\t// " + cls)
    for frame, n, e in shapes(cls):
        if cls == "tc-prose":
            e = raw_block(frame, cls, n)
        try:
            expr = fn(e)
        except Wait:
            waiting.append("ref(%s, %s, %d)" % (q(frame), q(cls), n))
            continue
        out.append("\tdocument(%s, %s, %s, %s, %d)," % (q(dstate(cls, expr)), expr, q(frame), q(cls), n))
print("\n".join(out))
print("\n// WAIT " + ", ".join(waiting))
