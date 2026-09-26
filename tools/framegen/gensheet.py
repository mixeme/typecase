"""Emit refs records for sheets, scrims, overlays, dialogs and the sign-in page."""
import os

HERE = os.path.dirname(os.path.abspath(__file__))
src = open(os.path.join(HERE, "genmenu.py"), encoding="utf-8").read()
exec(src.split("\nout = []")[0])

AVAILABLE = {"tc-rubric", "tc-tree", "tc-field", "tc-in", "tc-line", "tc-dim", "tc-railsec", "tc-form", "tc-formfoot", "tc-head",
             "tc-bar", "tc-btn", "tc-close", "tc-kbd", "tc-check", "tc-list", "tc-crumbs", "tc-tickrow", "tc-mark", "tc-mono", "tc-tag",
             "tc-badge", "tc-push", "tc-seg", "tc-num", "tc-range"}


def child_rail(k):
    if k.has("tc-railfoot"):
        return railfoot(k)
    if k.has("tc-railsec"):
        return railsec(k)
    if k.has("tc-nav"):
        return nav(k)
    if k.has("tc-zone"):
        return zone(k)
    return anyblock(k)


def blocks_ok(e):
    """Every block inside e has a widget type already."""
    for k in e.els():
        b = [c for c in k.cls if c.startswith("tc-") and "__" not in c and "--" not in c]
        if any(c not in AVAILABLE and c not in ("tc-grow", "tc-stack", "tc-line", "tc-root") for c in b):
            if not (k.has("tc-stack") or k.has("tc-line")) or [c for c in b if c not in ("tc-stack", "tc-line")]:
                pass
        blocks_ok(k)


def field(k):
    f = []
    lab = [x for x in k.els() if x.tag == "label"]
    if lab:
        f.append("Label: " + q(lab[0].text()))
    w = {"tc-field--w100": "FieldW100", "tc-field--sm": "FieldSM", "tc-field--md": "FieldMD", "tc-field--lg": "FieldLG", "tc-field--full": "FieldFull",
         "tc-field--fill200": "FieldFill200"}
    for c, x in w.items():
        if k.has(c):
            f.append("Width: w." + x)
    if k.has("tc-grow"):
        f.append("Place: w.Grown")
    f.append('Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}')
    return "w.Field{" + ", ".join(f) + "}"


def sheet(e):
    blocks_ok(e)
    f = []
    head = kid(e, "tc-sheet__head")
    svg = [x for x in head.els() if x.tag == "svg"]
    if svg:
        f.append("Icon: " + icon(svg[0]))
        if svg[0].has("tc-ico--danger"):
            f.append("Tone: w.SheetIconDanger")
        elif not svg[0].has("tc-ico--accent"):
            f.append("Tone: w.SheetIconInk")
    f.append("Title: " + q(kid(head, "tc-sheet__title").text()))
    hp = []
    for x in head.els():
        if x.tag == "svg" or x.has("tc-sheet__title"):
            continue
        if x.has("tc-dim"):
            hp.append(dimx(x))
        elif x.has("tc-close"):
            hp.append('w.Close{Label: "Close"' + (", Place: w.Pushed" if x.has("tc-push") else "") + "}")
        elif x.has("tc-tag"):
            hp.append(tag(x))
        elif x.has("tc-btn"):
            hp.append(button(x))
        elif x.has("tc-in") and x.has("is-pick"):
            hp.append(pickw(x))
        elif x.has("tc-in"):
            qq = kid(x, "tc-q")
            hp.append(filterw(x))
        elif x.has("tc-crumbs"):
            hp.append(crumbs(x))
        elif x.has("tc-table__file"):
            m = kid(x, "tc-mark")
            hp.append("w.SheetSubject{Mark: " + mark(m) + ", Text: " + q(x.els()[1].text()) + "}")
        else:
            raise SystemExit("sheet head? " + str(x.cls))
    if hp:
        f.append("Head: []w.SheetHeadPart{" + ", ".join(hp) + "}")
    body = kid(e, "tc-sheet__body")
    bp = []
    for x in body.els():
        if x.has("tc-rubric"):
            bp.append(rubric_full(x))
        elif x.has("tc-tree"):
            bp.append(tree(x))
        elif x.has("tc-railsec"):
            bp.append(railsec(x))
        elif x.has("tc-sheet__fields"):
            bp.append("w.SheetFields{Items: []w.Field{" + ", ".join(field(y) for y in x.els()) + "}}")
        elif x.has("tc-field"):
            bp.append(field(x))
        elif x.tag == "textarea":
            bp.append(areaw(x))
        elif x.has("tc-sheet__keys"):
            ks = x.els()
            pairs = ["{Keys: " + q(ks[i].text()) + ", Does: " + q(ks[i + 1].text()) + "}" for i in range(0, len(ks), 2)]
            bp.append("w.SheetKeys{Items: []w.SheetKey{" + ", ".join(pairs) + "}}")
        elif x.has("tc-field__hint--top6"):
            bp.append("w.SheetHint{Text: " + q(x.text()) + "}")
        elif x.has("tc-dim--lh1-55"):
            bp.append("w.SheetNote{Text: " + q(x.text()) + "}")
        elif x.has("tc-dim"):
            bp.append(dimx(x))
        elif x.has("tc-check"):
            bp.append(checkw(x))
        elif x.has("tc-sheet__checks"):
            bp.append("w.SheetChecks{Items: []w.Check{" + ", ".join(baritem(y) for y in x.els()) + "}}")
        elif x.has("tc-sheet__from"):
            item = kid(x, "tc-line__item")
            lines_, cur = [], []
            for y in item.kids:
                if not isinstance(y, str) and y.tag == "br":
                    lines_.append(cur)
                    cur = []
                else:
                    cur.append(y)
            lines_.append(cur)
            outl = []
            for ln in lines_:
                tmp = El("x", {})
                tmp.kids = ln
                outl.append(tparts(tmp, "[]w.Inline{"))
            bp.append("w.SheetFrom{Label: " + q(kid(x, "tc-dim").text()) + ", Lines: [][]w.Inline{" + ", ".join(outl) + "}}")
        elif x.has("tc-line") and not x.has("tc-sheet__pick"):
            bp.append(linew2(x))
        elif x.has("tc-sheet__pick"):
            rail = kid(x, "tc-sheet__pick-rail")
            lst_ = kid(x, "tc-sheet__pick-list")
            rp = [railsec(y) if y.has("tc-railsec") else tree(y) for y in rail.els()]
            lp = []
            for y in lst_.els():
                if y.has("tc-list__group"):
                    lp.append(listgroup(y))
                elif y.has("tc-list"):
                    lp.append(lst(y))
                elif y.has("tc-crumbs"):
                    lp.append(crumbs(y))
                else:
                    lp.append(anyblock(y))
            bp.append("w.SheetPick{Rail: []w.SheetPickRail{" + ", ".join(rp) + "}, List: []w.SheetPickList{" + ", ".join(lp) + "}" +
                      (", Short: true" if x.has("tc-sheet__pick--196") else "") + "}")
        else:
            bp.append(anyblock(x))
    f.append("Body: []w.SheetPart{\n\t\t" + ",\n\t\t".join(bp) + "}")
    if body.has("tc-sheet__body--pad0"):
        f.append("Flush: true")
    foot = kid(e, "tc-sheet__foot")
    if foot is not None:
        fp = []
        for x in foot.els():
            if x.has("tc-btn"):
                fp.append(button(x))
            elif x.has("tc-kbd"):
                fp.append("w.Kbd{Keys: " + q(x.text()) + (", Place: w.Pushed" if x.has("tc-push") else "") + "}")
            elif x.has("tc-line"):
                fp.append("w.Line{Gap: w.LineGap7, Place: w.Pushed, Items: []w.LinePart{" + ", ".join(button(y) if y.has("tc-btn") else "w.Kbd{Keys: " + q(y.text()) + "}" for y in x.els()) + "}}")
            elif x.has("tc-dim"):
                fp.append(dimx(x))
            elif x.has("tc-push"):
                fp.append("w.Push{Items: []w.PushPart{" + ", ".join(button(y) for y in x.els()) + "}}")
            else:
                raise SystemExit("foot? " + str(x.cls))
        f.append("Foot: []w.SheetFootPart{" + ", ".join(fp) + "}")
    for c, x in {"tc-sheet--w400": "SheetWidth400", "tc-sheet--w480": "SheetWidth480", "tc-sheet--w520": "SheetWidth520",
                 "tc-sheet--w680": "SheetWidth680", "tc-sheet--max100pc-full": "SheetWidthFull"}.items():
        if e.has(c):
            f.append("Width: w." + x)
    return "w.Sheet{" + ", ".join(f) + "}"


def scrim(e):
    s = kid(e, "tc-sheet")
    if s is None:
        pad_ = ", Pad: w.ScrimPad38" if e.has("tc-scrim--padt38") else (", Pad: w.ScrimPad26" if e.has("tc-scrim--padt26") else "")
        return "w.Scrim{Content: " + anyblock(e.els()[0]) + pad_ + "}"
    pad = ""
    for c, x in {"tc-scrim--padt26": "ScrimPad26", "tc-scrim--padt38": "ScrimPad38"}.items():
        if e.has(c):
            pad = ", Pad: w." + x
    try:
        content = sheet(s)
    except Wait:
        content = 'w.Sheet{Title: "Sheet"}'
    return "w.Scrim{Content: " + content + pad + "}"


def overlayw(e):
    f = []
    t = kid(e, "tc-top")
    if t is not None:
        tt = top(t)
        f.append("Top: &" + (tt if tt is not None else top2(t)))
    b = body(kid(e, "tc-shell__body")).replace("w.ShellBody{", "w.OverlayBody{").replace("Rail: []w.RailPart{", "Rail: []w.OverlayRailPart{")
    f.append("Body: " + b)
    sc = kid(e, "tc-scrim")
    try:
        f.append("Scrim: " + scrim(sc))
    except Wait:
        f.append('Scrim: w.Scrim{Content: w.Sheet{Title: "Sheet"}}')
    return "w.Overlay{" + ", ".join(f) + "}"


def dialog(e):
    blocks_ok(e)
    items = []
    for x in e.els():
        if x.has("tc-head"):
            items.append(head(x))
        elif x.has("tc-form"):
            items.append(formw(x))
        elif x.has("tc-formfoot"):
            items.append(formfootw(x))
        elif x.has("tc-bar"):
            items.append(bar2(x))
        elif x.has("tc-head__title"):
            gap = {"tc-head__title--bot4": "DialogTitleGap4", "tc-head__title--bot6": "DialogTitleGap6", "tc-head__title--bot8": "DialogTitleGap8"}
            g = [v for c, v in gap.items() if x.has(c)]
            items.append("w.DialogTitle{Text: " + q(x.text()) + (", Gap: w." + g[0] if g else "") + "}")
        elif x.has("tc-head__sub"):
            items.append("w.DialogSub{Parts: " + inl(x) + "}")
        elif x.has("tc-dialog__logo") or x.has("tc-dialog__logo--wide"):
            items.append("w.DialogLogo{Logo: " + logo(kid(x, "tc-brand__logo")) + (", Wide: true" if x.has("tc-dialog__logo--wide") else "") + "}")
        elif x.has("tc-auth__tagline"):
            items.append("w.Tagline{Lead: " + q(x.els()[0].text()) + ", Text: " + q(x.own_text()) + (", Centered: true" if x.has("tc-auth__tagline--mid") else "") + "}")
        elif x.has("tc-auth__stack"):
            its = []
            links = []
            for y in x.els():
                if y.has("tc-auth__links"):
                    links = [q(z.text()) for z in y.els()]
                elif y.has("tc-form"):
                    its.append(formw(y))
                else:
                    its.append(button(y))
            items.append("w.AuthStack{Items: []w.AuthStackPart{" + ", ".join(its) + "}, Links: []string{" + ", ".join(links) + "}}")
        elif x.has("tc-dialog__lede"):
            items.append("w.Lede{Text: " + q(x.text()) + "}")
        elif x.has("tc-dialog__split"):
            ph = kid(x, "tc-dialog__photo")
            side = kid(x, "tc-dialog__side")
            flds = [y for y in side.els()[1:] if y.has("tc-field")]
            aft = [y for y in side.els()[1:] if y.has("tc-dim")]
            items.append("w.PhotoCrop{Colour: " + q(ph.attrs.get("data-tc-colour", "")) + ", Note: " + q(side.els()[0].text()) + ", Fields: []w.Field{" + ", ".join(field(y) for y in flds) + "}" +
                         (", After: &w.Dim{Parts: []w.Inline{w.Text(" + q(aft[0].text()) + ")}, Size: w.Dim11}" if aft else "") + "}")
        else:
            items.append(anyblock(x))
    f = ["Items: []w.DialogPart{\n\t\t" + ",\n\t\t".join(items) + "}"]
    v = {"is-wide": "DialogWide", "is-narrow": "DialogNarrow", "tc-dialog--max520": "Dialog520", "tc-dialog--max540": "Dialog540",
         "tc-dialog--max560": "Dialog560", "tc-dialog--max440-padt0-padb0": "DialogPage440", "tc-dialog--max520-padt0-padb0": "DialogPage520",
         "tc-dialog--max420-flush": "DialogCanvas420"}
    for c, x in v.items():
        if e.has(c):
            f.append("Variant: w." + x)
    return "w.Dialog{" + ", ".join(f) + "}"


def auth(e):
    try:
        d = dialog(kid(e, "tc-dialog"))
    except Wait:
        d = "w.Dialog{Variant: w.DialogNarrow}"
    return "w.Auth{Dialog: " + d + "}"


def ostate(cls, expr):
    if cls == "tc-sheet":
        for k, v in [("SheetPick", "pick"), ("SheetIconDanger", "danger"), ("w.Close", "close"), ("Width: w.SheetWidthFull", "wide"), ("Icon:", "icon")]:
            if k in expr:
                return v
        return "plain"
    if cls == "tc-scrim":
        return "pad" if "Pad:" in expr else "plain"
    if cls == "tc-overlay":
        return "top" if "Top:" in expr else "plain"
    if cls == "tc-dialog":
        for k, v in [("DialogWide", "wide"), ("DialogNarrow", "narrow"), ("DialogPage", "page"), ("DialogCanvas", "canvas"), ("Variant:", "width")]:
            if k in expr:
                return v
        return "plain"
    return "plain"


out, waiting = [], []
for cls, fn in [("tc-sheet", sheet), ("tc-scrim", scrim), ("tc-overlay", overlayw), ("tc-dialog", dialog), ("tc-auth", auth)]:
    out.append("\n\t// " + cls)
    for frame, n, e in shapes(cls):
        try:
            expr = fn(e)
        except Wait:
            waiting.append("ref(%s, %s, %d)" % (q(frame), q(cls), n))
            continue
        out.append("\toverlay(%s, %s, %s, %s, %d)," % (q(ostate(cls, expr)), expr, q(frame), q(cls), n))
print("\n".join(out))
print("\n// WAIT " + ", ".join(waiting))


_block_child = block_child


def block_child(k):
    if k.has("tc-dialog"):
        return dialog(k)
    if k.has("tc-scrim"):
        return scrim(k)
    return _block_child(k)


def shellw2(e):
    f = []
    t = kid(e, "tc-top")
    if t is not None:
        tt = top(t)
        f.append("Top: &" + (tt if tt is not None else top2(t)))
    o = kid(e, "tc-overlay")
    if o is not None:
        f.append("Body: " + overlayw(o))
        return "w.Shell{" + ", ".join(f) + "}"
    return shellw(e)


out2, waiting2 = [], []
for cls, fn in [("tc-main", main), ("tc-side", side), ("tc-shell", shellw2)]:
    out2.append("\n\t// " + cls)
    for frame, n, e in shapes(cls):
        try:
            expr = fn(e)
        except Wait:
            expr = None
        if expr is None:
            waiting2.append("ref(%s, %s, %d)" % (q(frame), q(cls), n))
            continue
        st = mstate(cls, e, expr)
        if cls == "tc-shell" and "w.Overlay" in expr:
            st = "overlay"
        out2.append("\tshell(%s, %s, %s, %s, %d)," % (q(st), expr, q(frame), q(cls), n))
open(os.path.join(HERE, "main2_out.txt"), "w", encoding="utf-8").write("\n".join(out2) + "\n\n// WAIT " + ", ".join(waiting2) + "\n")
