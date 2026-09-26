"""Generate internal/screens: a function per route frame, built by the record
generators of plan 2 applied to the whole frame."""
import os, re, sys, traceback

os.environ["GENCARDS_MAIN"] = "0"
HERE = os.path.dirname(os.path.abspath(__file__))
REPO = os.path.abspath(os.path.join(HERE, "..", ".."))
OUT_DIR = os.environ.get("FRAMEGEN_OUT", os.path.join(HERE, "out"))
os.makedirs(OUT_DIR, exist_ok=True)
src = open(os.path.join(HERE, "genfinal.py"), encoding="utf-8").read().split("\n\ndef run(")[0]
exec(src)
exec(open(os.path.join(HERE, "gsfix.py"), encoding="utf-8").read())
exec(open(os.path.join(HERE, "gsfix2.py"), encoding="utf-8").read())

ROOT = REPO
MOCK = open(os.path.join(ROOT, "docs", "visual", "carrel-ui-mockups.html"), encoding="utf-8").read()
CATALOG_ONLY = {"m-desktop-2", "m-desktop-8"}


def frames():
    it = list(re.finditer(r'<div class="frame([^"]*)" data-frame="([^"]+)"', MOCK))
    heads = [(m.start(), m.group(1), re.sub("<[^>]+>", "", m.group(2))) for m in re.finditer(r'<h2 id="([^"]*)"[^>]*>(.*?)</h2>', MOCK)]
    for k, m in enumerate(it):
        j = it[k + 1].start() if k + 1 < len(it) else len(MOCK)
        blk = MOCK[m.start():j]
        lab = re.search(r'<div class="frame-label"><span>(.*?)</span>(.*?)</div>', blk, re.S)
        if re.search(r'class="frag"', lab.group(2)) or m.group(2) in CATALOG_ONLY:
            continue
        sec = ("", "")
        for p, hid, t in heads:
            if p < m.start():
                sec = (hid, t)
        yield m.group(2), sec, re.sub("<[^>]+>", "", lab.group(1)).strip(), inner(blk)[inner(blk).index(lab.group(0)) + len(lab.group(0)):]


def inner(blk):
    """The frame div's content, by balancing div tags."""
    depth = 0
    for t in re.finditer(r"<(/?)div\b[^>]*>", blk):
        depth += -1 if t.group(1) else 1
        if depth == 0:
            start = blk.index(">") + 1
            return blk[start:t.start()]
    raise ValueError("unbalanced frame")


def unwrap(e):
    """The desktop drawing's chrome is not product: the title bar and status
    line go, other mk-* elements give their children in their place."""
    kids = []
    for k in e.kids:
        if isinstance(k, El):
            if k.has("mk-titlebar") or k.has("mk-status"):
                continue
            if any(c.startswith("mk-") for c in k.cls):
                unwrap(k)
                kids.extend(k.kids)
                continue
            unwrap(k)
        kids.append(k)
    e.kids = kids


def fname(fid):
    s = re.sub(r"^(s|m)-", "", fid)
    return "".join(p[:1].upper() + p[1:] for p in re.split(r"[-_]", s))


class Gap(Exception):
    pass


def screen(root):
    if root.has("tc-shell"):
        return shellw2(root)
    raise Gap("root " + " ".join(root.cls) + " > " + ",".join(" ".join(k.cls) for k in root.els()[:2]))


def gen_screens():
    out = {}
    gaps = []
    names = set()
    for fid, (hid, htitle), cap, html in frames():
        root = parse(html.strip())
        unwrap(root)
        name = fname(fid)
        assert name not in names, name
        names.add(name)
        try:
            expr = screen(root)
            if expr is None:
                raise Gap("generator gave nothing")
        except Wait as e:
            gaps.append((fid, cap, "wait: " + str(e)))
            continue
        except Gap as e:
            gaps.append((fid, cap, str(e)))
            continue
        except SystemExit as e:
            tb = [f.name for f in traceback.extract_tb(e.__traceback__)][-6:]
            gaps.append((fid, cap, "unknown: " + str(e)[:160] + " @ " + ">".join(tb)))
            continue
        except Exception as e:
            tb = [f.name for f in traceback.extract_tb(e.__traceback__)][-6:]
            gaps.append((fid, cap, "error: " + type(e).__name__ + ": " + str(e)[:160] + " @ " + ">".join(tb)))
            continue
        out.setdefault((hid, htitle), []).append((fid, name, cap, expr))
    n = sum(len(v) for v in out.values())
    print("routes", n, "gaps", len(gaps))
    for g in gaps:
        print("GAP", g)
    import json
    json.dump({"sections": [[list(k), v] for k, v in out.items()], "gaps": gaps},
              open(os.path.join(OUT_DIR, "screens_out.json"), "w", encoding="utf-8"), ensure_ascii=False, indent=0)


gen_screens()
