"""Write internal/screens from screens_out.json: a file per mockup section, a
function per frame, the route list in mockup order."""
import json, os, re

HERE = os.path.dirname(os.path.abspath(__file__))
REPO = os.path.abspath(os.path.join(HERE, "..", ".."))
OUT_DIR = os.environ.get("FRAMEGEN_OUT", os.path.join(HERE, "out"))
OUT = os.path.join(REPO, "internal", "screens")
d = json.load(open(os.path.join(OUT_DIR, "screens_out.json"), encoding="utf-8"))


def split_top(s):
    """Split s at top-level commas (outside braces, parens and strings)."""
    parts, depth, cur, i, instr = [], 0, [], 0, False
    while i < len(s):
        c = s[i]
        if instr:
            cur.append(c)
            if c == "\\":
                cur.append(s[i + 1])
                i += 2
                continue
            if c == '"':
                instr = False
        elif c == '"':
            instr = True
            cur.append(c)
        elif c in "{([":
            depth += 1
            cur.append(c)
        elif c in "})]":
            depth -= 1
            cur.append(c)
        elif c == "," and depth == 0:
            parts.append("".join(cur).strip())
            cur = []
        else:
            cur.append(c)
        i += 1
    if "".join(cur).strip():
        parts.append("".join(cur).strip())
    return parts


def match_close(s, i):
    """Index of the brace closing the one at s[i]."""
    depth, instr, j = 0, False, i
    while j < len(s):
        c = s[j]
        if instr:
            if c == "\\":
                j += 2
                continue
            if c == '"':
                instr = False
        elif c == '"':
            instr = True
        elif c == "{":
            depth += 1
        elif c == "}":
            depth -= 1
            if depth == 0:
                return j
        j += 1
    raise ValueError("unbalanced")


def fmt(s, width=100):
    """Break composite literals longer than width: one element a line."""
    out, i, instr = [], 0, False
    while i < len(s):
        c = s[i]
        if instr:
            out.append(c)
            if c == "\\":
                out.append(s[i + 1])
                i += 2
                continue
            if c == '"':
                instr = False
            i += 1
            continue
        if c == '"':
            instr = True
            out.append(c)
            i += 1
            continue
        if c == "{":
            j = match_close(s, i)
            body = s[i + 1:j]
            if len(body) > width:
                elems = split_top(body)
                out.append("{\n")
                for e in elems:
                    out.append(fmt(e, width) + ",\n")
                out.append("}")
            else:
                out.append("{" + body + "}")
            i = j + 1
            continue
        out.append(c)
        i += 1
    return "".join(out)


def fname(hid):
    n = re.sub(r"^(s|m)-", "", hid) or "misc"
    return n.replace("-", "_") + ".go"


files = {}
routes = []
for (hid, htitle), items in d["sections"]:
    f = files.setdefault(fname(hid), [])
    for fid, name, cap, expr in items:
        f.append((hid, htitle, fid, name, cap, expr))
        routes.append((fid, cap, name))

for fn, items in files.items():
    buf = ["package screens", "", 'import w "carrel.dev/typecase/widget"', ""]
    titles = sorted(set(it[1] for it in items))
    buf.insert(0, "// Раздел макета: " + "; ".join(titles) + ".\n")
    for hid, htitle, fid, name, cap, expr in items:
        buf.append("// %s is frame «%s» (%s)." % (name, cap, fid))
        buf.append("func %s() Screen {" % name)
        buf.append("\treturn " + fmt(expr))
        buf.append("}")
        buf.append("")
    src = "\n".join(buf)
    if "noteBody" in src:
        ref = open(os.path.join(REPO, "internal", "refs", "inputs.go"), encoding="utf-8").read()
        i = ref.index("const noteBody = `")
        j = ref.index("`", i + len("const noteBody = `")) + 1
        src += "\n// noteBody is the text of the note being edited in the notes frames.\n" + ref[i:j] + "\n"
    open(os.path.join(OUT, fn), "w", encoding="utf-8", newline="\n").write(src)

rb = ["package screens", "", "// Routes lists the frames of the mockup built as screens, in the mockup's", "// order: /f/<Frame> (plan 3, step 1; Р-16).", "var Routes = []Route{"]
for fid, cap, name in routes:
    rb.append("\t{Frame: %s, Caption: %s, Screen: %s}," % (json.dumps(fid), json.dumps(cap, ensure_ascii=False), name))
rb.append("}")
open(os.path.join(OUT, "routes_list.go"), "w", encoding="utf-8", newline="\n").write("\n".join(rb) + "\n")
print(len(files), "files", len(routes), "routes")
