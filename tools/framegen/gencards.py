"""Emit refs records for tables, cards and the rest of the last family."""
import os

HERE = os.path.dirname(os.path.abspath(__file__))
src = open(os.path.join(HERE, "gendoc.py"), encoding="utf-8").read()
exec(src.split("\nout, waiting = [], []")[0])


def tparts(e, typ="[]w.CardTextPart{"):
    """Running text of a block as []w.TextPart."""
    parts = []
    for k in e.kids:
        if isinstance(k, str):
            t = re.sub(r"\s+", " ", k)
            if t:
                parts.append(("t", t))
        elif k.tag == "br":
            parts.append(("x", "w.Break{}"))
        elif k.tag == "b" and k.has("tc-card__key"):
            parts.append(("x", "w.Key(" + q(k.text()) + ")"))
        elif k.tag == "b" and k.els():
            parts.append(("x", "w.Bold{Parts: " + tparts(k, "[]w.Inline{") + "}"))
        elif k.tag == "b":
            parts.append(("x", "w.Strong(" + q(k.text()) + ")"))
        elif k.tag == "a":
            parts.append(("x", "w.Link{Text: " + q(k.text()) + ', Href: "#"}'))
        elif k.has("tc-mono") or k.tag == "code":
            parts.append(("x", mono(k)))
        elif k.has("tc-dim--accent"):
            parts.append(("x", "w.Span{Face: w.FaceAccent, Text: " + q(k.text()) + "}"))
        elif k.has("tc-dim--danger"):
            parts.append(("x", "w.Span{Face: w.FaceDanger, Text: " + q(k.text()) + "}"))
        elif k.has("tc-dim"):
            parts.append(("x", dimx(k)))
        elif k.has("tc-badge"):
            parts.append(("x", badge(k)))
        elif k.has("tc-kbd"):
            parts.append(("x", "w.Kbd{Keys: " + q(k.text()) + "}"))
        elif k.has("tc-btn"):
            parts.append(("x", button(k)))
        elif k.has("tc-mark"):
            parts.append(("x", mark(k)))
        elif k.has("tc-tag"):
            parts.append(("x", tag(k)))
        else:
            raise SystemExit("text? " + k.tag + str(k.cls))
    if parts and parts[0][0] == "t":
        parts[0] = ("t", parts[0][1].lstrip())
    if parts and parts[-1][0] == "t":
        parts[-1] = ("t", parts[-1][1].rstrip())
    return typ + ", ".join(("w.Text(" + q(v) + ")") if k == "t" else v for k, v in parts if v) + "}"


def msg(e):
    f = []
    svg = [k for k in e.els() if k.tag == "svg"]
    if svg:
        f.append("Icon: " + icon(svg[0]))
        f.append("Parts: " + tparts(kid(e, "tc-msg__aside"), "[]w.MsgPart{"))
        pu = kid(e, "tc-push")
        if pu is not None:
            f.append("End: &w.Push{Items: []w.PushPart{" + ", ".join(button(b) for b in pu.els()) + "}}")
    else:
        f.append("Parts: " + tparts(e, "[]w.MsgPart{"))
    if e.has("is-alert"):
        f.append("Tone: w.MsgAlert")
    if e.has("is-accent"):
        f.append("Tone: w.MsgAccent")
    sp = {"tc-msg--flush": "MsgFlush", "tc-msg--flush-top2": "MsgFlushTop2", "tc-msg--flush-top4": "MsgFlushTop4", "tc-msg--flush-top11": "MsgFlushTop11",
          "tc-msg--flush-bot12": "MsgFlushBottom12", "tc-msg--mr0-ml0": "MsgFlushSides", "tc-msg--flush-top9": "MsgFlushTop9", "tc-msg--flush-r0": "MsgFlushEnd",
          "tc-msg--mx20": "MsgInset20", "tc-msg--m0-14-12": "MsgSheet"}
    for c, x in sp.items():
        if e.has(c):
            f.append("Space: w." + x)
    if e.has("tc-grow"):
        f.append("Place: w.Grown")
    return "w.Msg{" + ", ".join(f) + "}"


def diag(e):
    v = {"tc-diag--top9-ink": "DiagInk", "tc-diag--danger": "DiagDanger"}
    var = [x for c, x in v.items() if e.has(c)]
    return "w.Diag{Text: " + q(e.text_raw()) + (", Variant: w." + var[0] if var else "") + "}"


def _text_raw(self):
    return "".join(k if isinstance(k, str) else k.text_raw() for k in self.kids)


El.text_raw = _text_raw


def danger(e):
    b = kid(e, "tc-btn")
    f = ["Button: " + button(b)]
    if e.own_text():
        f.append("Note: " + q(e.own_text()))
    for c, x in {"tc-danger--topauto-flush-padt11": "DangerAtEnd", "tc-danger--top0-pad11-14": "DangerTight"}.items():
        if e.has(c):
            f.append("Space: w." + x)
    return "w.Danger{" + ", ".join(f) + "}"


def meter(e):
    bar = kid(e, "tc-meter__bar")
    f = ["Percent: " + bar.attrs.get("data-tc-width", "0%").rstrip("%")]
    if e.attrs.get("data-tc-width"):
        f.append("Track: " + e.attrs["data-tc-width"].rstrip("%"))
    if bar.has("tc-meter__bar--alert"):
        f.append("Alert: true")
    for c, x in {"tc-meter--w60": "Meter60", "tc-meter--max260-fill": "MeterFill"}.items():
        if e.has(c):
            f.append("Width: w." + x)
    return "w.Meter{" + ", ".join(f) + "}"


def runbar(e):
    f = ["Meter: " + meter(kid(e, "tc-meter"))]
    ps = []
    for k in e.kids:
        if isinstance(k, str):
            t = re.sub(r"\s+", " ", k)
            if t.strip():
                ps.append("w.Text(" + q(t) + ")")
        elif k.has("tc-meter"):
            continue
        elif k.has("tc-dim--danger"):
            ps.append("w.Span{Face: w.FaceDanger, Text: " + q(k.text()) + "}")
        elif k.has("tc-dim--accent"):
            ps.append("w.Span{Face: w.FaceAccent, Text: " + q(k.text()) + "}")
        elif k.has("tc-mono"):
            ps.append(mono(k))
        elif k.has("tc-line"):
            ps.append("w.Line{Gap: w.LineGap6, Place: w.Pushed, Items: []w.LinePart{" + ", ".join(button(b) for b in k.els()) + "}}")
        elif k.has("tc-push"):
            ps.append("w.Push{Items: []w.PushPart{" + ", ".join(button(b) for b in k.els()) + "}}")
        else:
            raise SystemExit("runbar? " + str(k.cls))
    f.append("Parts: []w.RunbarPart{" + ", ".join(ps) + "}")
    if e.has("tc-runbar--danger"):
        f.append("Failed: true")
    return "w.Runbar{" + ", ".join(f) + "}"


def queue(e):
    h = kid(e, "tc-queue__head")
    f = ["Title: " + q(h.els()[0].text()), "Total: w.Mono{Text: " + q(kid(h, "tc-mono").text()) + ", Variant: w.MonoDusty, Place: w.Pushed}",
         "Toggle: " + button(kid(h, "tc-btn"))]
    lst_ = kid(e, "tc-queue__list")
    if lst_ is not None:
        its = []
        for li in lst_.els():
            row = kid(li, "tc-queue__row")
            failed = kid(row, "tc-dim--danger") is not None
            g = ["Name: " + q(row.els()[0].text()), "Meta: " + q(row.els()[1].text())]
            m = kid(li, "tc-meter")
            if m is not None:
                g.append("Meter: &" + meter(m))
            if failed:
                g.append("Failed: true")
            a = kid(li, "tc-queue__acts")
            if a is not None:
                g.append("Acts: []w.Button{" + ", ".join(button(b) for b in a.els()) + "}")
            its.append("{" + ", ".join(g) + "}")
        f.append("Items: []w.QueueItem{" + ", ".join(its) + "}")
    if e.has("is-min"):
        f.append("Min: true")
    for c, x in {"tc-queue--flush-full": "QueueFlushFull", "tc-queue--flush-wauto": "QueueFlushAuto"}.items():
        if e.has(c):
            f.append("Variant: w." + x)
    return "w.Queue{" + ", ".join(f) + "}"


def stat(e):
    return "w.Stat{Value: " + q(e.els()[0].text()) + ", Label: " + q(e.own_text()) + (", Error: true" if e.has("is-error") else "") + "}"


def stats(e):
    return "w.Stats{Items: []w.Stat{" + ", ".join(stat(k) for k in e.els()) + "}}"


def steps(e):
    items = [k for k in e.els() if k.has("tc-steps__item")]
    on = [i + 1 for i, k in enumerate(items) if k.has("is-on")]
    return "w.Steps{Items: []string{" + ", ".join(q(k.text()) for k in items) + "}" + (", On: %d" % on[0] if on else "") + "}"


def diff(e):
    its = []
    for it in e.els():
        r = kid(it, "tc-rubric")
        title = "w.Rubric{Text: \"\", Lead: &w.Mono{Text: " + q(kid(r, "tc-mono").text()) + "}, Items: []w.RubricPart{w.RubricAside{Text: " + q(kid(r, "tc-rubric__aside").text()) + "}}}"
        opts = []
        for c in kid(it, "tc-diff__options").els():
            tick = kid(c, "tc-tickrow__tick")
            v = kid(c, "tc-diff__value")
            opts.append("w.Check{" + ("On: true, " if not tick.has("is-off") else "") + "Parts: []w.CheckPart{w.Choice{Source: " + q(kid(v, "tc-diff__source").text()) + ", Value: " + q(v.own_text()) + "}}}")
        its.append("{Title: " + title + ", Options: []w.Check{" + ", ".join(opts) + "}}")
    return "w.Diff{Items: []w.DiffItem{\n\t\t" + ",\n\t\t".join(its) + "}}"


def empty(e):
    f = []
    svg = [k for k in e.els() if k.tag == "svg"]
    if svg:
        f.append("Icon: " + icon(svg[0]))
    lead = kid(e, "tc-empty__lead") or kid(e, "tc-empty__lead-sm")
    if lead is not None:
        f.append("Lead: " + q(lead.text()))
    ps = []
    for k in e.kids:
        if isinstance(k, str):
            t = re.sub(r"\s+", " ", k)
            if t.strip():
                ps.append("w.Text(" + q(t.strip() if not ps else t) + ")")
        elif k.tag == "svg" or k.has("tc-empty__lead") or k.has("tc-empty__lead-sm"):
            continue
        elif k.has("tc-btn"):
            ps.append(button(k))
        elif k.has("tc-dim--t12-5") or k.has("tc-dim--t12"):
            face = "FaceLarge" if k.has("tc-dim--t12-5") else "FaceSmall"
            its = []
            for y in k.kids:
                if isinstance(y, str):
                    if y.strip():
                        its.append("w.Text(" + q(" ".join(y.split())) + ")")
                elif y.has("tc-mono"):
                    its.append(mono(y))
                elif y.tag == "b":
                    its.append("w.Strong(" + q(y.text()) + ")")
            ps.append("w.Aside{Face: w." + face + ", Parts: []w.Inline{" + ", ".join(its) + "}}")
        elif k.has("tc-dim"):
            ps.append("w.Dim{Parts: []w.Inline{" + ", ".join(("w.Text(" + q(" ".join(y.split())) + ")") if isinstance(y, str) else mono(y) for y in k.kids if not isinstance(y, str) or y.strip()) + "}, Size: w.Dim11_5, Variant: w.DimWrap}")
        elif k.tag == "b" and k.has("tc-dim--accent"):
            ps.append("w.Emphasis(" + q(k.text()) + ")")
        elif k.has("tc-mono"):
            ps.append(mono(k))
        elif k.has("tc-kbd"):
            ps.append("w.Kbd{Keys: " + q(k.text()) + "}")
        else:
            raise SystemExit("empty? " + str(k.cls))
    if ps:
        f.append("Parts: []w.EmptyPart{" + ", ".join(ps) + "}")
    v = {"tc-empty--t12-pad14-20": "EmptySmall", "tc-empty--dusty": "EmptyDusty", "tc-empty--t12-pad12-20": "EmptySmallTop", "tc-empty--dashed": "EmptyDashed",
         "tc-empty--drop": "EmptyDrop", "tc-empty--drop-side": "EmptyDropSide"}
    for c, x in v.items():
        if e.has(c):
            f.append("Variant: w." + x)
            break
    return "w.Empty{" + ", ".join(f) + "}"


def thumb(e):
    forms = {"tc-thumb--on-card": "ThumbOnCard", "tc-thumb--preview": "ThumbPreview", "tc-thumb--wide": "ThumbWide", "tc-thumb--tile": "ThumbSquare"}
    f = ["Icon: " + icon([k for k in e.els() if k.tag == "svg"][0])]
    for c, x in forms.items():
        if e.has(c):
            f.append("Form: w." + x)
    return "w.Thumb{" + ", ".join(f) + "}"


def tile(e):
    return "w.Tile{Thumb: " + thumb(kid(e, "tc-thumb")) + ", Name: " + q(kid(e, "tc-tile__name").text()) + ", Meta: " + q(kid(e, "tc-tile__meta").text()) + (", Selected: true" if e.has("is-sel") else "") + "}"


def grid(e):
    return "w.Grid{Tiles: []w.Tile{" + ", ".join(tile(k) for k in e.els()) + "}" + (", Flush: true" if e.has("tc-grid--flush-fill-pad8") else "") + "}"


def pad(e):
    return "w.Pad{Items: []w.PadPart{" + ", ".join(thumb(k) if k.has("tc-thumb") else card(k) for k in e.els()) + "}}"


def fields(e):
    its = []
    ch = e.els()
    for i in range(0, len(ch), 2):
        dt, dd = ch[i], ch[i + 1]
        face = ""
        if dd.has("tc-mono--accent"):
            face = "ValueMonoAccent"
        elif dd.has("tc-mono--t11"):
            face = "ValueMonoSmall"
        elif dd.has("tc-mono"):
            face = "ValueMono"
        elif dd.has("tc-dim--accent"):
            face = "ValueAccent"
        vp = []
        for k in dd.kids:
            if isinstance(k, str):
                t = re.sub(r"\s+", " ", k)
                if t.strip():
                    vp.append("w.Text(" + q(t) + ")")
            elif k.tag == "br":
                vp.append("w.Break{}")
            elif k.has("tc-mark"):
                vp.append(mark(k))
            elif k.has("tc-badge"):
                vp.append(badge(k))
            elif k.has("tc-tag"):
                vp.append(tag(k))
            elif k.has("tc-mono"):
                vp.append(mono(k))
            elif k.has("tc-dim--danger"):
                vp.append("w.Span{Face: w.FaceDanger, Text: " + q(k.text()) + "}")
            elif k.has("tc-dim"):
                vp.append(dimx(k))
            else:
                raise SystemExit("dd? " + str(k.cls))
        if vp and vp[0].startswith("w.Text("):
            vp[0] = "w.Text(" + q(eval(vp[0][7:-1]).lstrip()) + ")"
        if vp and vp[-1].startswith("w.Text("):
            vp[-1] = "w.Text(" + q(eval(vp[-1][7:-1]).rstrip()) + ")"
        its.append("{Term: " + q(dt.text()) + ", Value: []w.ValuePart{" + ", ".join(vp) + "}" + (", Face: w." + face if face else "") + "}")
    f = ["Items: []w.FieldsItem{\n\t\t" + ",\n\t\t".join(its) + "}"]
    v = [("tc-fields--cols88", "tc-fields--pad4-14-12", "Fields88Pad"), ("tc-fields--cols88", None, "Fields88"), ("tc-fields--bleed", "tc-fields--cols130", "FieldsBleed130"),
         ("tc-fields--bleed", None, "FieldsBleed"), ("tc-fields--cols130", "tc-fields--flush-padt14", "Fields130Flush")]
    for a, b, x in v:
        if e.has(a) and (b is None or e.has(b)):
            f.append("Variant: w." + x)
            break
    return "w.Fields{" + ", ".join(f) + "}"


def anyblock(k):
    """A value of the widget a block element is, for a place."""
    table_ = [("tc-msg", msg), ("tc-diag", diag), ("tc-danger", danger), ("tc-meter", meter), ("tc-runbar", runbar), ("tc-queue", queue),
              ("tc-stats", stats), ("tc-steps", steps), ("tc-diff", diff), ("tc-empty", empty), ("tc-grid", grid), ("tc-thumb", thumb),
              ("tc-fields", fields), ("tc-sec", lambda x: sec(x)), ("tc-card", lambda x: card(x)), ("tc-table", lambda x: table(x)),
              ("tc-prose", prose), ("tc-list", lst), ("tc-tree", tree), ("tc-head", head), ("tc-bar", bar2), ("tc-rubric", rubric_full),
              ("tc-tickrow", tickrow), ("tc-railsec", railsec), ("tc-docmeta", lambda x: docmeta2(x)), ("tc-pad", pad)]
    for c, fn in table_:
        if k.has(c):
            return fn(k)
    if k.has("tc-docbar"):
        return docbar(k)
    if k.has("tc-doc"):
        return doc(k)
    if k.has("tc-sheet__block"):
        lines_ = []
        for it in kid(k, "tc-line__text").els():
            lines_.append("{Name: " + q(it.els()[0].text()) + ", Note: w.Dim{Parts: []w.Inline{w.Text(" + q(kid(it, "tc-dim").text()) + ")}}}")
        return "w.SheetBlock{Title: " + rubric_full(kid(k, "tc-rubric")) + ", Items: []w.SheetLine{" + ", ".join(lines_) + "}}"
    if k.has("tc-diff__agree"):
        return "w.DiffAgree{Property: w.Mono{Text: " + q(kid(k, "tc-mono").text()) + "}, Text: " + q(k.own_text()) + "}"
    if k.has("tc-seg"):
        return baritem(k)
    if k.has("tc-check"):
        return "w.Check{Parts: []w.CheckPart{w.Text(" + q(k.text()) + ")}}"
    if k.has("tc-field"):
        return field(k)
    if k.has("tc-form"):
        return formw(k)
    if k.has("tc-formfoot"):
        return formfootw(k)
    if k.has("tc-btn"):
        return button(k)
    if k.has("tc-line"):
        return linew2(k)
    if k.has("tc-mono"):
        return mono(k)
    if k.has("tc-in"):
        return "w.Input{Value: " + q(k.attrs.get("value", "")) + ', Label: "x"}'
    raise Wait()


def sec(e):
    items = []
    for k in e.els():
        if k.has("tc-sec__list"):
            gap = 0
            for g in (5, 6):
                if k.has("tc-stack--gap-%d" % g):
                    gap = g
            items.append("w.SecList{Items: [][]w.SecItemPart{" + ", ".join(tparts(li, "[]w.SecItemPart{") for li in k.els()) + "}" + (", Gap: %d" % gap if gap else "") +
                         (", Muted: true" if k.has("tc-sec__list--muted") else "") + "}")
        elif k.has("tc-field__hint"):
            items.append("w.SecHint{Text: " + q(k.text()) + (", Apart: true" if k.has("tc-field__hint--top6") else "") + "}")
        elif k.has("tc-sec__note"):
            items.append("w.SecNote{Text: " + q(k.text()) + "}")
        elif k.has("tc-sec__acts"):
            items.append("w.SecActs{Items: []w.Button{" + ", ".join(button(b) for b in k.els()) + "}}")
        elif k.has("tc-doc__list"):
            its = ["{Tick: w.Tick{" + ("Done: true, " if kid(li, "tc-tick").has("is-done") else "") + "Variant: w.TickNudge, Label: " + q(li.text()) + "}, Text: " + q(li.text()) + "}" for li in k.els()]
            items.append("w.SecTasks{Items: []w.DocTask{" + ", ".join(its) + "}}")
        elif k.has("tc-card__top8"):
            fs = []
            for ln in k.els():
                fl = kid(ln, "tc-table__file")
                fs.append("{Icon: " + icon(fl.els()[0]) + ", Name: " + q(fl.els()[1].text()) + ", Size: w.Mono{Text: " + q(kid(ln, "tc-mono").text()) + "}, Tag: " +
                          tag(kid(ln, "tc-tag")) + ", Remove: " + button(kid(ln, "tc-btn")) + "}")
            items.append("w.SecFiles{Items: []w.SecFile{" + ", ".join(fs) + "}}")
        elif k.has("tc-line__text") and k.has("tc-stack"):
            ps = []
            for x in k.els():
                if x.has("tc-stack__item"):
                    ps.append("w.SecOpened{Name: " + q(x.own_text()) + ", Size: w.Dim{Parts: []w.Inline{w.Text(" + q(kid(x, "tc-dim").text()) + ")}}}")
                elif x.has("tc-line"):
                    ps.append("w.Line{Gap: w.LineGap6, Items: []w.LinePart{" + ", ".join(button(b) for b in x.els()) + "}}")
                else:
                    ps.append(dimx(x))
            items.append("w.SecOpen{Items: []w.SecOpenPart{" + ", ".join(ps) + "}}")
        elif k.has("tc-rubric"):
            items.append(rubric_full(k))
        else:
            items.append(anyblock(k))
    return "w.Sec{Items: []w.SecPart{\n\t\t" + ",\n\t\t".join(items) + "}" + (", Flush: true" if e.has("tc-sec--flush-flush-padt11") else "") + "}"


CARD_TEXT = {"tc-card__line": "CardLine", "tc-card__num": "CardNum", "tc-card__term": "CardTerm", "tc-card__note": "CardUsage", "tc-card__note-dusty": "CardSmallPrint",
             "tc-card__why": "CardWhy", "tc-card__warn": "CardWarn", "tc-card__alert-note": "CardAlertText", "tc-card__lede": "CardLede",
             "tc-card__caption": "CardCaption", "tc-card__above": "CardAbove"}
CARD_ROW = {"tc-card__acts": "CardActs", "tc-card__top2": "CardTop2", "tc-card__top8": "CardTop8", "tc-card__top9": "CardTop9", "tc-card__top10": "CardTop10",
            "tc-card__top11": "CardTop11", "tc-card__files": "CardAttached", "tc-card__foot": "CardFoot"}


def rowpart(x):
    if x.has("tc-btn"):
        return button(x)
    if x.has("tc-field"):
        return field(x)
    if x.has("tc-tag"):
        return tag(x)
    if x.tag == "svg" and x.has("tc-ico--turn"):
        return "w.Turn{}"
    if x.has("tc-field__hint--flush-fill"):
        return "w.FillHint{Parts: " + tparts(x) + "}"
    if x.has("tc-dim--t12-5") and not x.has("tc-dim"):
        return "w.Span{Face: w.FaceLarge, Text: " + q(x.text()) + "}"
    if x.has("tc-dim"):
        return dimx(x)
    if x.has("tc-in"):
        return "w.Input{Value: " + q(x.attrs.get("value", "")) + ', Mono: true, Variant: w.InputKeyFill, Label: "Key"}'
    raise SystemExit("card row? " + str(x.cls))


def card(e):
    items = []
    for k in e.els():
        kinds = [x for c, x in CARD_TEXT.items() if k.has(c)]
        rows = [x for c, x in CARD_ROW.items() if k.has(c)]
        if kinds:
            items.append("w.CardText{Kind: w." + kinds[0] + ", Parts: " + tparts(k) + "}")
        elif k.has("tc-field") and k.has("tc-card__top9"):
            items.append("w.SpacedField{Field: " + field(k) + "}")
        elif rows:
            gap = 0
            for g in (6, 7, 8, 10, 22):
                if k.has("tc-line--gap-%d" % g):
                    gap = g
            items.append("w.CardRow{Kind: w." + rows[0] + (", Gap: %d" % gap if gap else "") + ", Items: []w.CardRowPart{" + ", ".join(rowpart(x) for x in k.els()) + "}}")
        elif k.has("tc-field__hint"):
            items.append("w.CardHint{Parts: " + tparts(k) + (", Line: true" if k.tag == "div" else "") + "}")
        elif k.has("tc-card__list"):
            items.append("w.CardList{Items: []string{" + ", ".join(q(li.text()) for li in k.els()) + "}}")
        elif k.has("tc-card__fields"):
            fs = []
            for x in k.els():
                if x.has("tc-field"):
                    fs.append(field(x))
                elif x.has("tc-btn"):
                    fs.append(button(x))
                else:
                    fs.append("w.Span{Face: w.FaceMuted12, Text: " + q(x.text()) + "}")
            items.append("w.CardFields{Items: []w.CardFieldsPart{" + ", ".join(fs) + "}" + (", Wide: true" if k.has("tc-card__fields--gap10-12") else "") +
                         (", Rule: true" if k.has("tc-card__fields--rule") else "") + "}")
        elif k.has("tc-card__split"):
            qr = kid(k, "tc-card__qr")
            side = kid(k, "tc-card__qr-side")
            sp = [button(x) if x.has("tc-btn") else dimx(x) for x in side.els()]
            use = [x for x in qr.els()[0].els() if x.tag == "use"][0]
            items.append("w.CardQR{Colour: " + q(qr.attrs.get("data-tc-colour", "")) + ", Symbol: " + q(use.attrs["href"][1:]) + ", Side: []w.CardQRPart{" + ", ".join(sp) + "}}")
        elif k.has("tc-card__pick"):
            tick = k.els()[0]
            item = k.els()[1]
            g = []
            if tick.has("is-off"):
                g.append("Ticked: w.TickedOff")
            g.append("Title: " + q(kid(item, "tc-stack__item-title").text()))
            g.append("Note: " + tparts(kid(item, "tc-stack__item-note")))
            a = kid(item, "tc-stack__item-aside")
            if a is not None:
                g.append("Aside: " + q(a.text()))
            fl = kid(item, "tc-field")
            if fl is not None:
                g.append("Field: &" + field(fl))
            items.append("w.CardPick{" + ", ".join(g) + "}")
        elif k.has("tc-card__grid"):
            gs = []
            for x in k.els():
                if x.has("tc-dim"):
                    gs.append(dimx(x))
                elif x.tag == "textarea":
                    gs.append("w.Area{Text: pasteArgs, Variant: w.AreaPaste, Mono: true, Label: \"Arguments\"}")
                elif x.has("tc-in"):
                    gs.append("w.Input{Value: " + q(x.attrs.get("value", "")) + ", Label: \"x\"}")
                else:
                    gs.append(button(x))
            items.append("w.CardGrid{Items: []w.CardGridPart{" + ", ".join(gs) + "}}")
        elif k.has("tc-card__rows"):
            rs = []
            for it in k.els():
                m = kid(it, "tc-mono")
                d = kid(it, "tc-dim")
                rs.append("{Head: w.Mono{Text: " + q(m.text()) + "}, Note: w.Dim{Parts: " + inl(d).replace("[]w.Inline{", "[]w.Inline{", 1) + "}}")
            items.append("w.CardRows{Items: []w.CardRowsItem{" + ", ".join(rs) + "}}")
        elif k.has("tc-card__head"):
            ls = []
            for it in k.els():
                slot = kid(it, "tc-dim--slot")
                rest = El("x", {})
                rest.kids = [x for x in it.kids if x is not slot]
                ls.append("{Slot: " + q(slot.text()) + ", Parts: " + tparts(rest) + "}")
            items.append("w.CardHead{Lines: []w.CardHeadLine{" + ", ".join(ls) + "}}")
        elif k.has("tc-line") and not [c for c in k.cls if "__" in c]:
            items.append("w.Line{Block: true, Gap: w.LineGap8, Items: []w.LinePart{" + ", ".join(button(x) if x.has("tc-btn") else field(x) if x.has("tc-field") else
                         "w.Input{Value: " + q(x.attrs.get("value", "")) + ', Mono: true, Variant: w.InputKeySmFill, Label: "Key"}' for x in k.els()) + "}}")
        elif k.has("tc-mono") and k.tag == "div":
            items.append(mono(k))
        else:
            items.append(anyblock(k))
    f = ["Items: []w.CardPart{\n\t\t" + ",\n\t\t".join(items) + "}"]
    v = [("tc-card--edge-accent", "tc-card--flush", "CardChosenFlushGrow"), ("tc-card--edge-accent-flush", None, "CardChosenFlush"), ("tc-card--edge-accent", None, "CardChosen"),
         ("tc-card--alert", None, "CardAlert"), ("tc-card--top14", None, "CardApart"), ("tc-card--pad12-14", None, "CardSnug"), ("tc-card--flush", None, "CardFlush")]
    for a, b, x in v:
        if e.has(a) and (b is None or e.has(b)):
            f.append("Variant: w." + x)
            break
    if e.has("tc-grow"):
        f.append("Place: w.Grown")
    return "w.Card{" + ", ".join(f) + "}"


COLW = {"tc-table__col--tick": "ColTick", "tc-table__col--tick-sm": "ColTickSmall", "tc-table__col--size": "ColSize", "tc-table__col--type": "ColType",
        "tc-table__col--time": "ColTime", "tc-table__col--num": "ColNum", "tc-table__col--menu": "ColMenu", "tc-table__col--changed": "ColChanged",
        "tc-table__col--w26": "Col26", "tc-table__col--w80": "Col80", "tc-table__col--w90": "Col90", "tc-table__col--w110": "Col110", "tc-table__col--w120": "Col120"}


def cellpart(x):
    if x.has("tc-table__file"):
        m = kid(x, "tc-mark")
        name = [y for y in x.els() if y.has("tc-table__filename") and not y.has("tc-mark")][0]
        g = []
        if m is not None:
            g.append("Mark: &" + mark(m))
        else:
            g.append("Icon: " + icon(x.els()[0]))
        g.append("Name: " + q(name.text()))
        if name.has("tc-table__filename--accent"):
            g.append("Accent: true")
        return "w.TableFile{" + ", ".join(g) + "}"
    if x.tag == "br":
        return "w.Break{}"
    if x.has("tc-tick"):
        v = ", Variant: w.TickEdgeDanger" if x.has("tc-tick--edge-danger") else (", Variant: w.TickEdgeDusty" if x.has("tc-tick--edge-dusty") else "")
        return "w.Tick{" + ("Done: true, " if x.has("is-done") else "") + 'Label: "Done"' + v + "}"
    if x.has("tc-btn"):
        return button(x)
    if x.has("tc-badge"):
        return badge(x)
    if x.has("tc-tag"):
        return tag(x)
    if x.has("tc-drop"):
        return drop(x)
    if x.has("tc-table__wrap"):
        return "w.CellWrap{Items: []w.CellPart{" + ", ".join(cellpart(y) for y in x.els()) + "}}"
    if x.has("tc-line"):
        its = []
        for y in x.els():
            if y.has("tc-mark"):
                its.append(mark(y))
            elif y.has("tc-mono"):
                its.append(mono(y))
        return "w.Line{Gap: w.LineGap3, Items: []w.LinePart{" + ", ".join(its) + "}}"
    if x.has("tc-mono"):
        return mono(x)
    if x.has("tc-seg"):
        opts = ["{Label: " + q(o.text()) + (", On: true" if o.has("is-on") else "") + "}" for o in x.els()]
        return "w.Seg{Options: []w.SegOption{" + ", ".join(opts) + "}, Variant: w.SegCompact, Place: w.Pushed}"
    if x.tag == "svg" and x.has("tc-ico--inline"):
        return "w.Chevron{}"
    if x.has("tc-dim"):
        return dimx(x)
    raise SystemExit("cell? " + x.tag + str(x.cls))


def table(e):
    cols = []
    head = [x for x in e.els() if x.tag == "thead"]
    if head:
        for th in head[0].els()[0].els():
            g = []
            if th.own_text():
                g.append("Label: " + q(th.own_text()))
            for c, x in COLW.items():
                if th.has(c):
                    g.append("Width: w." + x)
            if th.has("tc-table__n3"):
                g.append("Rank: w.Rank3")
            if th.has("tc-table__n4"):
                g.append("Rank: w.Rank4")
            if th.has("is-2nd"):
                g.append("Secondary: true")
            t = kid(th, "tc-tickrow__tick")
            if t is not None:
                tk = "w.TickedPart" if t.has("is-part") else ("w.TickedOff" if t.has("is-off") else "w.TickedOn")
                g.append("Tick: ptr(" + tk + ")")
            cols.append("{" + ", ".join(g) + "}")
    bodies = [x for x in e.els() if x.tag == "tbody"]
    body = bodies[0] if bodies else e
    rows = []
    for tr in [x for x in body.els() if x.tag == "tr"]:
        cells = []
        for td in tr.els():
            g = []
            parts = []
            for k in td.kids:
                if isinstance(k, str):
                    t = re.sub(r"\s+", " ", k)
                    if t.strip():
                        parts.append("w.Text(" + q(t) + ")")
                elif k.has("tc-tickrow__tick"):
                    g.append("Pick: ptr(" + ("w.TickedOff" if k.has("is-off") else "w.TickedOn") + ")")
                else:
                    parts.append(cellpart(k))
            if parts:
                g.insert(0, "Parts: []w.CellPart{" + ", ".join(parts) + "}")
            if td.has("tc-mono"):
                g.append("Mono: true")
            if td.has("tc-dim"):
                g.append("Dim: true")
            if td.has("tc-acts"):
                g.append("Acts: true")
            if td.has("tc-table__name"):
                g.append("Name: true")
            if td.has("is-ph"):
                g.append("Unset: true")
            if any(c.endswith("--dusty") for c in td.cls):
                g.append("Tone: w.CellDusty")
            if any(c.endswith("--danger") for c in td.cls):
                g.append("Tone: w.CellDanger")
            if td.has("tc-dim--muted"):
                g.append("Tone: w.CellMuted")
            if any(c.endswith("--t11") for c in td.cls):
                g.append("Small: true")
            if td.attrs.get("colspan"):
                g.append("Span: " + td.attrs["colspan"])
            if not head:
                if td.has("tc-table__n3"):
                    g.append("Rank: w.Rank3")
                if td.has("tc-table__n4"):
                    g.append("Rank: w.Rank4")
                for c, x in COLW.items():
                    if td.has(c):
                        g.append("Width: w." + x)
            cells.append("{" + ", ".join(g) + "}")
        st = {"is-sel": "RowPicked", "is-folder": "RowFolder", "is-off": "RowOff", "is-on": "RowOn", "tc-table__row--dim": "RowFaded", "tc-table__row--off": "RowGone"}
        s = [x for c, x in st.items() if tr.has(c)]
        rows.append("{Cells: []w.TableCell{" + ", ".join(cells) + "}" + (", State: w." + s[0] if s else "") + "}")
    f = []
    if cols:
        f.append("Cols: []w.TableCol{" + ", ".join(cols) + "}")
    f.append("Rows: []w.TableRow{\n\t\t" + ",\n\t\t".join(rows) + "}")
    for c, x in {"tc-table--fit": "TableFit", "tc-table--wide": "TableWide", "tc-table--mx-20": "TableInset"}.items():
        if e.has(c):
            f.append("Variant: w." + x)
    return "w.Table{" + ", ".join(f) + "}"


def docmeta2(e):
    secs = []
    danger_ = None
    for k in e.els():
        if k.has("tc-danger"):
            danger_ = danger(k)
            continue
        its = []
        for x in k.els()[1:]:
            if x.has("tc-diag"):
                its.append(diag(x))
            elif x.has("tc-line") and not x.has("tc-line__text"):
                its.append("w.Line{Block: true, Gap: w.LineGap4, Items: []w.LinePart{" + ", ".join(tag(t) for t in x.els()) + "}}")
            elif x.has("tc-stack"):
                refs_ = []
                for it in x.els():
                    m = kid(it, "tc-mark")
                    a = [y for y in it.els() if y.tag == "a"]
                    name = a[0].text() if a else kid(it, "tc-dim--accent").text()
                    refs_.append("{" + ("Mark: &" + mark(m) + ", " if m is not None else "") + "Name: " + q(name) + (', Href: "#"' if a else "") +
                                 ", What: w.Dim{Parts: []w.Inline{w.Text(" + q(kid(it, "tc-dim").text()) + ")}}}")
                its.append("w.DocRefs{Items: []w.DocRef{" + ", ".join(refs_) + "}}")
            elif x.has("tc-line__text"):
                its.append("w.DocAttach{Name: " + q(x.own_text()) + ", Size: w.Dim{Parts: []w.Inline{w.Text(" + q(kid(x, "tc-dim").text()) + ")}}, Where: " + badge(kid(x, "tc-badge")) + "}")
            else:
                raise SystemExit("docmeta? " + str(x.cls))
        secs.append("{Title: " + rub(k.els()[0]) + ", Items: []w.DocmetaPart{" + ", ".join(its) + "}}")
    return "w.Docmeta{Sections: []w.DocmetaSec{\n\t\t" + ",\n\t\t".join(secs) + "}" + (", Danger: &" + danger_ if danger_ else "") + "}"


def cstate(cls, expr):
    rules = {
        "tc-msg": [("Icon:", "icon"), ("MsgAlert", "alert"), ("MsgAccent", "accent"), ("Strong(", "rich"), ("Mono{", "rich"), ("Space:", "flush")],
        "tc-diag": [("DiagInk", "ink"), ("DiagDanger", "danger")],
        "tc-danger": [("Space:", "end"), ("Note:", "aside")],
        "tc-meter": [("Alert:", "alert"), ("Width:", "width"), ("Track:", "track")],
        "tc-runbar": [("Failed:", "failed"), ("w.Push{", "stop"), ("w.Button{", "stop")],
        "tc-queue": [("Min: true", "min"), ("Variant:", "flush")],
        "tc-stat": [("Error:", "error")],
        "tc-steps": [("On: 1", "first"), ("On: 3", "last"), ("On: 2", "middle")],
        "tc-empty": [("EmptyDrop", "drop"), ("EmptyDashed", "dashed"), ("EmptyDusty", "dusty"), ("w.Button{", "action"), ("EmptySmall", "small")],
        "tc-thumb": [("ThumbOnCard", "card"), ("ThumbPreview", "preview"), ("ThumbWide", "wide"), ("ThumbSquare", "square")],
        "tc-tile": [("Selected:", "selected")],
        "tc-grid": [("Flush:", "flush")],
        "tc-fields": [("Fields88", "phone"), ("FieldsBleed", "bleed"), ("Fields130", "narrow")],
        "tc-sec": [("Flush: true", "flush"), ("SecList", "list"), ("SecActs", "acts"), ("SecHint", "hint")],
        "tc-card": [("CardPick", "pick"), ("CardAlert", "alert"), ("CardApart", "apart"), ("CardSnug", "snug"), ("CardChosen", "chosen"), ("CardFlush", "flush")],
        "tc-table": [("TableWide", "wide"), ("TableFit", "fit"), ("RowPicked", "picked"), ("RowFolder", "folder"), ("RowOff", "off"), ("Acts: true", "acts"), ("Cols:", "head")],
    }
    defaults = {"tc-steps": "middle", "tc-thumb": "tile", "tc-docmeta": "read"}
    for k, v in rules.get(cls, []):
        if k in expr:
            return v
    return defaults.get(cls, "plain")


FAM = [("tc-msg", msg), ("tc-diag", diag), ("tc-danger", danger), ("tc-meter", meter), ("tc-runbar", runbar), ("tc-queue", queue), ("tc-stat", stat),
       ("tc-stats", stats), ("tc-steps", steps), ("tc-diff", diff), ("tc-empty", empty), ("tc-thumb", thumb), ("tc-tile", tile), ("tc-grid", grid),
       ("tc-pad", pad), ("tc-fields", fields), ("tc-sec", sec), ("tc-card", card), ("tc-table", table), ("tc-docmeta", docmeta2)]
RAW = {"tc-diag", "tc-msg"}
if os.environ.get("GENCARDS_MAIN", "1") == "1":
    out, waiting = [], []
    for cls, fn in FAM:
        out.append("\n\t// " + cls)
        for frame, n, e in shapes(cls):
            if cls in RAW:
                e = raw_block(frame, cls, n)
            try:
                expr = fn(e)
            except Wait:
                waiting.append("ref(%s, %s, %d)" % (q(frame), q(cls), n))
                continue
            if expr.endswith(", Place: w.Grown}") and cls in ("tc-msg", "tc-card"):
                expr = expr[:-len(", Place: w.Grown}")] + "}"
            out.append("\tcards(%s, %s, %s, %s, %d)," % (q(cstate(cls, expr)), expr, q(frame), q(cls), n))
    print("\n".join(out))
    print("\n// WAIT " + ", ".join(waiting))
