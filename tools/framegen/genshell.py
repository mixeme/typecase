"""Emit refs records for the shell family (first part)."""
import os
import sys

HERE = os.path.dirname(os.path.abspath(__file__))
exec(open(os.path.join(HERE, "genrows.py"), encoding="utf-8").read().split("FAM = ")[0])


def place(e, f):
    if e.has("tc-push"):
        f.append("Place: w.Pushed")
    if e.has("tc-grow"):
        f.append("Place: w.Grown")
    if e.has("is-2nd"):
        f.append("Secondary: true")


def button(e):
    f = []
    svgs = [k for k in e.els() if k.tag == "svg"]
    label = e.own_text()
    lab = kid(e, "tc-btn__label")
    if lab is not None:
        label = lab.text()
        f.append("Folds: true")
    if label:
        f.append("Label: " + q(label))
    icons = [icon(s) for s in svgs]
    if icons and icons[-1] == '"chev"' and (label or len(icons) > 1):
        f.append("Menu: true")
        icons = icons[:-1]
    if icons:
        f.append("Icon: " + icons[0])
        box = {"tc-ico--sq16": "IconBox16", "tc-ico--sq17": "IconBox17", "tc-ico--sq19": "IconBox19"}
        for c in svgs[0].cls:
            if c in box:
                f.append("IconBox: w." + box[c])
    for k, v in {"is-primary": "ButtonPrimary", "is-danger": "ButtonDanger", "is-danger-solid": "ButtonDangerSolid"}.items():
        if e.has(k):
            f.append("Kind: w." + v)
    if e.has("is-sm"):
        f.append("Small: true")
    if e.has("is-focus"):
        f.append("Focus: true")
    for k, v in {"tc-btn--padx3": "ButtonPadX3", "tc-btn--padx4": "ButtonPadX4", "tc-btn--padx6": "ButtonPadX6",
                 "tc-btn--padx8": "ButtonPadX8", "tc-btn--on-card": "ButtonOnCard"}.items():
        if e.has(k):
            f.append("Variant: w." + v)
    t = e.attrs.get("title") or ("" if label else "More")
    if t:
        f.append("Title: " + q(t))
    place(e, f)
    return "w.Button{" + ", ".join(f) + "}"


def mono_v(e):
    v = {"tc-mono--t11-top4": "Mono11Top", "tc-mono--t10-dusty": "Mono10Dusty", "tc-mono--t11": "Mono11"}
    for k, x in v.items():
        if e.has(k):
            return ", Variant: w." + x
    return ""


def mono(e):
    return mono(e)


def rubric(e):
    return "w.Rubric{Text: " + q(e.own_text()) + "}"


def emblem(e):
    f = ['Symbol: "davloom"']
    t = [k for k in e.els() if k.tag == "title"]
    if t:
        f.append("Title: " + q(t[0].text()))
    if e.has("is-alert"):
        f.append("Alert: true")
    return "w.Emblem{" + ", ".join(f) + "}"


def dimx(e):
    parts = []
    for k in e.kids:
        if isinstance(k, str):
            if k.strip():
                parts.append("w.Text(" + q(norm(k) if not k.endswith(" ") else norm(k) + " ") + ")")
        elif k.has("tc-mono"):
            parts.append(mono(k))
    size = ", Size: w.Dim11" if e.has("tc-dim--t11") else ""
    return "w.Dim{Parts: []w.Inline{" + ", ".join(parts) + "}" + size + "}"


def railnote(e):
    f = []
    if e.has("is-alert"):
        f.append("Tone: w.RailStatusAlert")
    if e.has("is-off"):
        f.append("Tone: w.RailStatusOff")
    f.append("Text: " + q(e.own_text()))
    f.append("Detail: " + dimx(kid(e, "tc-dim")))
    return "w.RailStatus{" + ", ".join(f) + "}"


def navitem(a):
    f = []
    em = kid(a, "tc-emblem")
    if em is not None:
        f.append("Emblem: &" + emblem(em))
    else:
        f.append("Icon: " + icon([k for k in a.els() if k.tag == "svg"][0]))
    if a.own_text():
        f.append("Label: " + q(a.own_text()))
    n = kid(a, "tc-num")
    if n is not None:
        f.append("Count: &w.Num{Value: " + n.text() + ("" if not n.has("is-alert") else ", Alert: true") + "}")
    f.append('Href: "#"')
    if a.attrs.get("title"):
        f.append("Title: " + q(a.attrs["title"]))
    return "{" + ", ".join(f) + "}"


def nav(e):
    items, on = [], 0
    for i, a in enumerate(e.els()):
        if a.has("is-on"):
            on = i + 1
        items.append(navitem(a))
    f = ["Items: []w.NavItem{\n\t\t" + ",\n\t\t".join(items) + "}"]
    if on:
        f.append("On: %d" % on)
    if e.has("tc-nav--top2"):
        f.append("Apart: true")
    return "w.Nav{" + ", ".join(f) + "}"


def usermenu(e):
    f = ["Name: " + q(e.own_text())]
    if e.has("is-focus"):
        f.append("Focus: true")
    if not [k for k in e.els() if k.tag == "svg"]:
        f.append("NoArrow: true")
    return "w.UserMenu{" + ", ".join(f) + "}"


def query(e):
    return "w.Query{Value: " + q(e.attrs.get("value", "")) + ", Label: " + q(e.attrs.get("aria-label", "")) + "}"


def search(e):
    f = ["Query: " + query(kid(e, "tc-q"))]
    if e.has("is-focus"):
        f.append("Focus: true")
    if len([k for k in e.els() if k.tag == "svg"]) > 1:
        f.append("Clear: true")
    for k, v in {"tc-search--full": "SearchFull", "tc-search--fill180-wauto": "SearchFill"}.items():
        if e.has(k):
            f.append("Variant: w." + v)
    return "w.Search{" + ", ".join(f) + "}"


LOGO = '"4.79 -31.36 291.08 116.63"'


def logo(svg):
    f = ['Label: "Carrel"', "ViewBox: " + LOGO]
    for k, v in {"tc-brand__logo--h17-w44": "Logo17", "tc-brand__logo--h42-w104": "Logo42", "tc-brand__logo--h60-w150": "Logo60", "tc-brand__logo--h48-w118": "Logo48"}.items():
        if svg.has(k):
            f.append("Size: w." + v)
    return "w.Logo{" + ", ".join(f) + "}"


def brand(e):
    svg = kid(e, "tc-brand__logo")
    if svg is not None:
        return "w.Brand{Logo: &" + logo(svg) + "}"
    return "w.Brand{Title: " + q(e.text()) + "}"


def child(e):
    for c, fn in [("tc-btn", button), ("tc-badge", badge), ("tc-tag", tag), ("tc-rubric", rubric), ("tc-tickrow", tickrow),
                  ("tc-tree", tree), ("tc-nav", nav), ("tc-mono", mono), ("tc-emblem", emblem), ("tc-railstatus", railnote),
                  ("tc-usermenu", usermenu), ("tc-search", search), ("tc-brand", brand)]:
        if e.has(c):
            return fn(e)
    if e.has("tc-close"):
        return button(e)
    if e.has("tc-seg"):
        opts = ["{Label: " + q(o.text()) + (", On: true" if o.has("is-on") else "") + "}" for o in e.els()]
        return segw(e)
    if e.has("tc-field"):
        return field(e)
    if e.has("tc-in") and e.has("is-ph"):
        qq = kid(e, "tc-q")
        return "w.Filter{Query: w.Query{Value: " + q(qq.attrs.get("value", "")) + ", Label: " + q(qq.attrs.get("aria-label", "")) + "}, Waiting: true, Variant: w.FilterH26W190}"
    raise SystemExit("child? " + str(e.cls))


def split(e):
    b = e.els()
    return "w.Split{Main: " + button(b[0]) + ", More: " + button(b[1]) + "}"


def zone(e):
    return "w.Zone{Title: " + rubric(kid(e, "tc-rubric")) + ", Note: " + mono(kid(e, "tc-mono")) + "}"


def crumbs(e):
    path = ["{Label: " + q(a.text()) + ', Href: "#"}' for a in e.els() if a.has("tc-crumbs__link")]
    f = []
    if path:
        f.append("Path: []w.Crumb{" + ", ".join(path) + "}")
    f.append("Current: " + q([k for k in e.els() if k.tag == "b"][0].text()))
    d = kid(e, "tc-dim")
    if d is not None:
        f.append("Aside: &w.Dim{Parts: []w.Inline{w.Text(" + q(d.text()) + ")}}")
    return "w.Crumbs{" + ", ".join(f) + "}"


def acts(e):
    f = ["Items: []w.ActsPart{" + ", ".join(child(k) for k in e.els()) + "}"]
    if e.has("tc-acts--stack"):
        f.append("Stack: true")
    return "w.Acts{" + ", ".join(f) + "}"


def railsec(e):
    items, f = [], []
    for k in e.els():
        if k.has("tc-railsec__tags"):
            f.append("Tags: []w.Tag{" + ", ".join(tag(t) for t in k.els()) + "}")
        elif k.has("tc-railsec__note"):
            f.append("Note: " + q(k.text()))
        else:
            items.append(child(k))
    out = []
    if items:
        out.append("Items: []w.RailSecPart{\n\t\t" + ",\n\t\t".join(items) + "}")
    out += f
    if e.has("tc-railsec--top14"):
        out.append("Apart: true")
    return "w.RailSec{" + ", ".join(out) + "}"


def railfoot(e):
    items, f = [], []
    for k in e.els():
        if k.has("tc-railfoot__about"):
            f.append('About: &w.About{Name: ' + q(k.own_text()) + ", Build: " + mono(kid(k, "tc-mono")) + ', Href: "#rest"}')
        else:
            items.append(child(k))
    out = []
    if items:
        out.append("Items: []w.RailFootPart{" + ", ".join(items) + "}")
    out += f
    if e.has("tc-railfoot--strip"):
        out.append("Strip: true")
    return "w.RailFoot{" + ", ".join(out) + "}"


def top(e):
    f = []
    els = e.els()
    if els[0].has("tc-btn"):
        f.append("Lead: &" + button(els[0]))
    f.append("Brand: " + brand(kid(e, "tc-brand")))
    s = kid(e, "tc-search")
    if s is not None:
        f.append("Search: &" + search(s))
    a = kid(e, "tc-top__actions")
    if [k for k in a.els() if k.has("tc-drop__anchor")] or kid(e, "tc-top__search") is not None:
        return None
    f.append("Actions: []w.TopAction{" + ", ".join(child(k) for k in a.els()) + "}")
    if a.has("tc-top__actions--gap2"):
        f.append("Tight: true")
    return "w.Top{" + ", ".join(f) + "}"


def gstate(cls, e, expr):
    c = set(e.cls)
    table = {
        "tc-brand": [("tc-brand--t16", "title")],
        "tc-search": [("is-focus", "focus"), ("tc-search--fill180-wauto", "variant")],
        "tc-usermenu": [("is-focus", "focus")],
        "tc-emblem": [("is-alert", "alert")],
        "tc-crumbs": [],
        "tc-acts": [("tc-acts--stack", "stack")],
        "tc-nav": [("tc-nav--top2", "apart")],
        "tc-railsec": [("tc-railsec--top14", "apart")],
        "tc-railstatus": [("is-alert", "alert"), ("is-off", "off")],
        "tc-railfoot": [("tc-railfoot--strip", "strip")],
        "tc-top": [("", "")],
    }
    for k, v in table.get(cls, []):
        if k and k in c:
            return v
    defaults = {"tc-brand": "logo", "tc-search": "plain", "tc-usermenu": "plain", "tc-emblem": "plain", "tc-split": "plain",
                "tc-zone": "plain", "tc-crumbs": "path", "tc-acts": "row", "tc-nav": "plain", "tc-railsec": "plain",
                "tc-railstatus": "running", "tc-railfoot": "plain", "tc-top": "plain"}
    st = defaults[cls]
    if cls == "tc-search" and "Clear: true" in expr:
        st = "clear"
    if cls == "tc-usermenu" and "NoArrow" in expr:
        st = "no-arrow"
    if cls == "tc-emblem" and "Title:" in expr:
        st = "titled"
    if cls == "tc-crumbs" and "Aside" in expr:
        st = "aside"
    if cls == "tc-nav" and "Label:" not in expr:
        st = "icons"
    if cls == "tc-nav" and st == "plain" and expr.count("Icon:") == 1 and '"back"' in expr:
        st = "back"
    if cls == "tc-railsec" and "Tags:" in expr:
        st = "tags"
    if cls == "tc-railsec" and "Note:" in expr:
        st = "note"
    if cls == "tc-railfoot" and st == "plain" and "Items:" not in expr:
        st = "about"
    if cls == "tc-top" and "Lead:" in expr:
        st = "phone"
    return st


FAM = [("tc-brand", brand), ("tc-search", search), ("tc-usermenu", usermenu), ("tc-emblem", emblem), ("tc-split", split),
       ("tc-zone", zone), ("tc-crumbs", crumbs), ("tc-acts", acts), ("tc-nav", nav), ("tc-railsec", railsec),
       ("tc-railstatus", railnote), ("tc-railfoot", railfoot), ("tc-top", top)]
out, waiting = [], []
for cls, fn in FAM:
    out.append("\n\t// " + cls)
    for frame, n, e in shapes(cls):
        expr = fn(e)
        if expr is None:
            waiting.append('ref(%s, %s, %d)' % (q(frame), q(cls), n))
            continue
        out.append("\tshell(%s, %s, %s, %s, %d)," % (q(gstate(cls, e, expr)), expr, q(frame), q(cls), n))
print("\n".join(out))
print("\n// WAIT " + ", ".join(waiting))
