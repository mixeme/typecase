# Fields, forms and bar items, faithful to the frame (plan 3, step 1).


def hintw(x):
    return "w.FieldHint{Parts: []w.Inline{" + ", ".join(inlparts(x)) + "}" + ev(x, "HintVariant", "Variant") + "}"


def groupw(x):
    if x.has("tc-line"):
        form = ""
    elif x.has("tc-stack--gap-2"):
        form = ", Form: w.GroupColumn"
    else:
        form = ", Form: w.GroupColumnLoose"
    return "w.Group{Items: []w.FieldPart{" + ", ".join(fieldpart(y) for y in x.els()) + "}" + form + "}"


def fieldpart(x):
    c = control(x)
    if c is not None:
        return c
    if x.has("tc-field__hint"):
        return hintw(x)
    if x.has("tc-field__swatches"):
        cols = [y.attrs.get("data-tc-colour", "") for y in x.els()]
        on = [i for i, y in enumerate(x.els()) if y.has("tc-field__swatch--on")]
        return "w.Swatches{Colours: []string{" + ", ".join(q(c) for c in cols) + "}" + (", On: %d" % on[0] if on and on[0] else "") + "}"
    if x.has("tc-field__zoom"):
        return "w.Zoom{Value: " + q(kid(x, "tc-mono").text()) + "}"
    if x.has("tc-field__row"):
        return "w.FieldRow{Items: []w.FieldPart{" + ", ".join(fieldpart(y) for y in x.els()) + "}}"
    if x.has("tc-card__top2"):
        return groupw(x)
    if x.has("tc-field__photo-row"):
        ph = kid(x, "tc-avatar--photo")
        side_ = kid(x, "tc-field__photo-side")
        ini = "Initials: " + q(ph.text()) + ", " if ph.text() else ""
        return "w.Photo{" + ini + "Side: []w.FieldPart{" + ", ".join(fieldpart(y) for y in side_.els()) + "}}"
    if x.has("tc-line"):
        return linew2(x)
    if x.has("tc-mark"):
        return mark(x)
    if x.has("tc-dim"):
        return dimx(x)
    if x.has("tc-drop"):
        return drop(x)
    if x.has("tc-menu"):
        return menu(x)
    if x.has("tc-msg"):
        return msg(x)
    if x.has("tc-diag"):
        return diag(x)
    raise SystemExit("field part? " + x.tag + " " + str(x.cls))


def field(k):
    f = []
    lab = [x for x in k.els() if x.tag == "label" and not x.has("tc-in")]
    labspan = kid(k, "tc-field__label")
    skip = set()
    if labspan is not None:
        f.append("Label: " + q(labspan.text()) + ", LabelSpan: true")
        skip.add(id(labspan))
    elif lab:
        l = lab[0]
        skip.add(id(l))
        f.append("Label: " + q(l.own_text() or " "))
        for a in l.els():
            if a.has("tc-dim"):
                f.append("LabelAside: " + dimx(a))
            elif a.has("tc-plain"):
                f.append("LabelAside: w.Plain{Text: " + q(a.text()) + "}")
            else:
                raise SystemExit("label aside? " + str(a.cls))
    parts = [fieldpart(x) for x in k.els() if id(x) not in skip]
    if parts:
        f.append("Parts: []w.FieldPart{" + ", ".join(parts) + "}")
    v = enum_for(k, "FieldWidth")
    if v:
        f.append("Width: " + v)
    for c in k.cls:
        m = re.match(r"tc-field--gap(\d+)$", c)
        if m:
            f.append("Gap: " + m.group(1))
    if k.has("is-own-line"):
        f.append("OwnLine: true")
    return "w.Field{" + ", ".join(f) + placef(k) + "}"


def linepart(x):
    c = control(x)
    if c is not None:
        return c
    if x.has("tc-field"):
        return field(x)
    if x.has("tc-kbd"):
        return "w.Kbd{Keys: " + q(x.text()) + placesec(x) + "}"
    if x.has("tc-dim"):
        return dimx(x)
    if x.has("tc-mark"):
        return mark(x)
    if x.has("tc-mono"):
        return mono(x)
    if x.has("tc-tag"):
        return tag(x)
    if x.has("tc-badge"):
        return badge(x)
    if x.has("tc-meter"):
        return meter(x)
    if x.has("tc-push"):
        return "w.Push{Items: []w.PushPart{" + ", ".join(push_item(y) for y in x.els()) + "}}"
    if x.has("tc-line"):
        return linew2(x)
    if x.tag == "b":
        return "w.Strong(" + q(x.text()) + ")"
    cl = " ".join(x.cls)
    if x.tag == "span" and cl in _faces:
        return "w.Faced{Face: w." + _faces[cl] + ", Items: []w.LinePart{" + ", ".join(linetext(x)) + "}}"
    raise SystemExit("line part? " + x.tag + " " + cl)


def linetext(e):
    out = []
    for k in e.kids:
        if isinstance(k, str):
            if norm(k):
                out.append("w.Text(" + q(norm(k)) + ")")
        else:
            out.append(linepart(k))
    return out


def linew2(e):
    own = {"tc-line"} | {c for c in e.cls if c.startswith("tc-line--gap-")}
    extra = [c for c in e.cls if c not in own and c not in ("tc-push", "tc-grow")]
    if extra:
        raise SystemExit("line with " + str(extra))
    f = ["Items: []w.LinePart{" + ", ".join(linetext(e)) + "}"]
    g = enum_for(e, "LineGap")
    if g and g != "w.LineGap0":
        f.append("Gap: " + g)
    if e.tag == "div":
        f.append("Block: true")
    return "w.Line{" + ", ".join(f) + placef(e) + "}"


def fsetpart(k):
    if k.has("tc-fset__grid"):
        return "w.FieldGrid{Items: []w.FieldGridPart{" + ", ".join(tickrow(x) for x in k.els()) + "}}"
    if k.has("tc-field"):
        return field(k)
    if k.has("tc-btn"):
        return button(k)
    if k.has("tc-line"):
        return linew2(k)
    if k.has("tc-grow") and not k.els() and not k.text():
        return "w.Grow{}"
    return anyblock(k)


def fset(e):
    kids = e.els()
    r = kids[0] if kids and kids[0].has("tc-rubric") else None
    items = [fsetpart(k) for k in kids if k is not r]
    f = []
    if r is not None:
        f.append("Rubric: " + rubric(r))
    f.append("Items: []w.FieldSetPart{" + ", ".join(items) + "}")
    v = enum_for(e, "FieldSetVariant")
    if v:
        f.append("Variant: " + v)
    return "w.FieldSet{" + ", ".join(f) + "}"


def formpart(k):
    if k.has("tc-fset"):
        return fset(k)
    if k.has("tc-field"):
        return field(k)
    if k.has("tc-check"):
        return checkw(k)
    if k.has("tc-rubric"):
        return rubric(k)
    if k.has("tc-card"):
        return card(k)
    if k.has("tc-msg"):
        return msg(k)
    if k.has("tc-form__cols"):
        return "w.FormCols{Items: []w.Field{" + ", ".join(field(y) for y in k.els()) + "}}"
    if k.has("tc-form__col--460") or k.has("tc-form__col--300"):
        its = [card(y) if y.has("tc-card") else anyblock(y) for y in k.els()]
        wide = "Wide: true, " if k.has("tc-form__col--460") else ""
        return "w.FormColumn{" + wide + "Items: []w.FormColumnPart{" + ", ".join(its) + "}}"
    raise SystemExit("form part? " + str(k.cls))


def formw(e):
    f = ["Items: []w.FormPart{" + ", ".join(formpart(k) for k in e.els()) + "}"]
    v = enum_for(e, "FormVariant")
    if v:
        f.append("Variant: " + v)
    if e.has("is-tight"):
        f.append("Tight: true")
    return "w.Form{" + ", ".join(f) + "}"


def formfootpart(k):
    if k.has("tc-btn"):
        return button(k)
    if k.has("tc-dim"):
        return dimx(k)
    if k.has("tc-push"):
        return "w.Push{Items: []w.PushPart{" + ", ".join(push_item(y) for y in k.els()) + "}}"
    cl = " ".join(k.cls)
    if k.tag == "span" and cl in _faces:
        return "w.Span{Face: w." + _faces[cl] + ", Text: " + q(k.text()) + "}"
    raise SystemExit("form foot part? " + str(k.cls))


def formfootw(e):
    return "w.FormFoot{Items: []w.FormFootPart{" + ", ".join(formfootpart(k) for k in e.els()) + "}" + ev(e, "FormFootVariant", "Variant") + "}"


def baritem(k):
    if k.has("tc-bar__sep"):
        f = []
        if k.has("is-2nd"):
            f.append("Secondary: true")
        if k.has("is-3rd"):
            f.append("Tertiary: true")
        return "w.BarSep{" + ", ".join(f) + "}"
    if k.has("tc-range"):
        ins = [i for i in k.els() if i.tag == "input"]
        return ("w.Range{From: " + inputw(ins[0]) + ", To: " + inputw(ins[1]) + ", Show: " + button(kid(k, "tc-btn")) +
                (", Secondary: true" if k.has("is-2nd") else "") + "}")
    c = control(k)
    if c is not None:
        return c
    if k.has("tc-mono"):
        return mono(k)
    if k.has("tc-drop"):
        return drop(k)
    raise SystemExit("bar? " + str(k.cls))


_block_child0 = block_child


def block_child(k):
    if k.has("tc-form"):
        return formw(k)
    if k.has("tc-formfoot"):
        return formfootw(k)
    if k.has("tc-dialog"):
        return dialog(k)
    return _block_child0(k)
