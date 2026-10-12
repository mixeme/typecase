"""Make a container of a frame a form (Carrel's task 149).

    MSYS_NO_PATHCONV=1 python internal/framegen/cmd/record/toform.py <frame> <class>[#n] <action> [key=value ...]

<class> is the container's class — tc-sheet, tc-form, tc-card, tc-fset —
and n which of them in the frame (1 by default; convert the last first: a
converted container is no <div> any more and the numbers shift). The
container becomes a <form> with the address; its fields get names (of their
labels), its checks become checkboxes, its picks the browser's lists, its
segments radio buttons, and its confirming button submits.

Keys: a field's label in lower case = the name it gets; button=<words> — the
button that submits, by its words (otherwise the first primary one — look at
which it took: «Copy the script» is primary too); id=<id> — the form's id,
for a confirming button outside it (then tie the button by hand if the
frame has another primary one); submit=0 — no button is made a submit;
keeppick=1 — picks stay buttons (a server's pick in a sheet's head opens a
menu of links); wrap=1 — the container is left as it is and a
<form class="tc-form tc-form--pad0"> is put around it.

Two segments in one form need names of their own: both come out as «choice»
when they stand under no label — rename them by hand. In Git Bash the address
needs MSYS_NO_PATHCONV=1, or it turns into a Windows path. Afterwards:
go generate ./internal/screens, fixrefs.sh, the tests, pixpair.sh.
"""
import re, sys

p = "docs/visual/carrel-ui-mockups.html"
s = open(p, encoding="utf-8").read()
fid, what, action = sys.argv[1], sys.argv[2], sys.argv[3]
names = dict(a.split("=", 1) for a in sys.argv[4:])
cls, _, nth = what.partition("#")
nth = int(nth or 1)
i = s.index('data-frame="%s"' % fid)
j = s.find('data-frame="', i + 20)
b = s[i:j]
cut = b.index("<figcaption") if "<figcaption" in b else len(b)
n0 = len(b)


def text(t):
    return re.sub(r"<[^>]+>", "", t).replace("&nbsp;", " ").strip()


def slug(t):
    t = text(t).lower()
    return names.get(t) or re.sub(r"[^a-z0-9]+", ".", t).strip(".") or "field"


ms = [m for m in re.finditer(r'<div class="((?:[^"]* )?%s(?: [^"]*)?)"((?: [a-z-]+="[^"]*")*)>' % re.escape(cls), b[:cut])]
assert len(ms) >= nth, "%s: %d containers %s" % (fid, len(ms), cls)
m = ms[nth - 1]
depth, k = 0, None
for t in re.finditer(r"<(/?)div\b", b[m.start():]):
    depth += -1 if t.group(1) else 1
    if depth == 0:
        k = m.start() + t.start()
        break
body = b[m.end():k]
assert "<form" not in body, fid + ": a form inside"
count = {"named": 0, "check": 0, "pick": 0, "seg": 0}


def name_tag(tag, nm):
    count["named"] += 1
    if ' type="' in tag:
        q = tag.index('"', tag.index(' type="') + 7) + 1
        return tag[:q] + ' name="%s"' % nm + tag[q:]
    return tag[:-1] + ' name="%s">' % nm


def controls(html, label):
    """Names the controls of a piece that stands under one label (or none)."""
    nm = slug(label) if label and text(label) else None

    def inp(im):
        tag = im.group(0)
        if ' name="' in tag or 'type="hidden"' in tag or "tc-check__box" in tag or "tc-seg__box" in tag or 'type="file"' in tag:
            return tag
        a = re.search(r'aria-label="([^"]*)"', tag)
        return name_tag(tag, nm or (slug(a.group(1)) if a else names.get("", "text")))
    html = re.sub(r"<(?:input|textarea)[ ][^>]*>", inp, html)

    def pick(pm):
        c, aria, inner = pm.group(1), pm.group(2) or "", pm.group(3)
        if names.get("keeppick"):
            return pm.group(0)  # a pick that opens a menu of links: left a button
        mark = "".join(re.findall(r'<span class="tc-mark[^"]*"[^>]*></span>', inner))
        words = text(inner)
        n2 = nm or (slug(re.search(r'aria-label="([^"]*)"', aria).group(1)) if aria else "pick")
        if not aria and label and text(label):
            aria = ' aria-label="%s"' % text(label)
        count["pick"] += 1
        if "<span" in inner.replace(mark, ""):
            return pm.group(0)  # an aside inside: left as it is drawn
        return '<span class="%s">%s<select class="tc-in__list" name="%s"%s><option value="%s" selected>%s</option></select></span>' % (
            c, mark, n2, aria, slug(words), words)
    html = re.sub(r'<button class="(tc-in (?:is-ph )?is-pick[^"]*)" type="button"( aria-label="[^"]*")?>((?:(?!</button>).)*)</button>', pick, html, flags=re.S)

    def check(cm):
        extra, off, words = cm.group(1), cm.group(2), cm.group(3)
        count["check"] += 1
        return ('<label class="tc-check%s"><input class="tc-check__box" type="checkbox" name="%s" value="1"%s>'
                '<span class="tc-tickrow__tick%s"></span>%s</label>') % (extra, slug(words), "" if off else " checked", off or "", words)
    html = re.sub(r'<button class="tc-check((?: is-[a-z0-9]+)*)" type="button"><span class="tc-tickrow__tick( is-off)?"></span>((?:(?!</button>).)*)</button>', check, html, flags=re.S)

    def seg(sm):
        head, opts = sm.group(1), sm.group(2)
        if "<a " in opts or "<label" in opts:
            return sm.group(0)
        n2 = nm or "choice"

        def opt(om):
            return '<label class="%s"><input class="tc-seg__box" type="radio" name="%s" value="%s"%s>%s</label>' % (
                om.group(1).replace("is-on ", ""), n2, slug(om.group(2)), " checked" if "is-on" in om.group(1) else "", om.group(2))
        opts2 = re.sub(r'<button class="((?:is-on )?tc-seg__opt)" type="button">([^<]*)</button>', opt, opts)
        if "<button" in opts2:
            return sm.group(0)
        count["seg"] += 1
        return head[:-1] + ' role="radiogroup">' + opts2 + "</span>"
    html = re.sub(r'(<span class="tc-seg[^"]*">)((?:(?!</span>).)*)</span>', seg, html, flags=re.S)
    return html


# a field's label names what stands in it; a field may hold a hint (a span) but no other field
out, pos = [], 0
for fm in re.finditer(r'<div class="tc-field[^"]*"[^>]*>', body):
    if fm.start() < pos:
        continue
    d, e = 0, None
    for t in re.finditer(r"<(/?)div\b", body[fm.start():]):
        d += -1 if t.group(1) else 1
        if d == 0:
            e = fm.start() + t.start()
            break
    inner = body[fm.end():e]
    lm = re.match(r"\s*<label>((?:(?!</label>).)*)</label>", inner, flags=re.S)
    out.append(controls(body[pos:fm.start()], None))
    out.append(fm.group(0) + controls(inner, lm.group(1) if lm else None) + "</div>")
    pos = e + 6
out.append(controls(body[pos:], None))
body = "".join(out)

wrap = names.pop("wrap", None)
fid_attr = names.get("id")
submit = names.get("submit", "1")
CONFIRM = r'(<button class="tc-btn (?:is-sm )?(?:is-primary|is-danger-solid)[^"]*") type="button"'
outer = False
button = names.get("button")
if button:
    # the confirming button is named by its words
    body, n = re.subn(r'(<button class="tc-btn[^"]*") type="button"(>(?:<svg.*?</svg>)?%s</button>)' % re.escape(button), r'\g<1> type="submit"\g<2>', body, count=1, flags=re.S)
    assert n == 1, fid + ": no button «%s» in %s" % (button, what)
elif submit != "0":
    body, n = re.subn(CONFIRM, r'\g<1> type="submit"', body, count=1)
    if n == 0:
        assert fid_attr, fid + ": no confirming button in " + what + " and no id to name the form by"
        outer = True
idattr = ' id="%s"' % fid_attr if fid_attr else ""
head, tail = b[:m.start()], b[k + 6:]
if outer:
    # the confirming button stands outside the form: it names the form it sends
    rep = r'\g<1> type="submit" form="%s"' % fid_attr
    tail2, n = re.subn(CONFIRM, rep, tail, count=1)
    if n == 0:
        head2, n = re.subn(CONFIRM, rep, head, count=1)
        assert n == 1, fid + ": no confirming button in the frame"
        head = head2
    else:
        tail = tail2
if wrap:
    b = (head + '<form class="tc-form tc-form--pad0"%s action="%s" method="post">' % (idattr, action) + m.group(0) + body + "</div></form>" + tail)
else:
    b = (head + '<form class="%s"%s%s action="%s" method="post">' % (m.group(1), m.group(2), idattr, action) + body + "</form>" + tail)
print(fid, what, count, len(b) - n0)
assert abs(len(b) - n0) < 4000
s = s[:i] + b + s[j:]
open(p, "w", encoding="utf-8", newline="\n").write(s)
