"""Read the widget package's enums: for each type with a class() method made
of a string array, which constant prints which classes."""
import glob, os, re

WIDGET = os.path.join(os.path.dirname(os.path.abspath(__file__)), "..", "..", "widget")


def load():
    consts = {}  # type -> [const names in iota order]
    arrays = {}  # type -> [class strings]
    for path in glob.glob(os.path.join(WIDGET, "*.go")):
        src = open(path, encoding="utf-8").read()
        for m in re.finditer(r"const \(\n(.*?)\n\)", src, re.S):
            cur, names = None, []
            for ln in m.group(1).split("\n"):
                ln = ln.split("//")[0].strip()
                if not ln:
                    continue
                mm = re.match(r"(\w+)\s+(\w+)\s*=\s*iota", ln)
                if mm:
                    if cur and names:
                        consts.setdefault(cur, names)
                    cur, names = mm.group(2), [mm.group(1)]
                    continue
                mm = re.match(r"(\w+)$", ln)
                if mm and cur:
                    names.append(mm.group(1))
            if cur and names:
                consts.setdefault(cur, names)
        for m in re.finditer(r"func \(\w+ (\w+)\) class\(\) string \{\s*return \[\.\.\.\]string\{(.*?)\}\[\w+\]", src, re.S):
            arrays[m.group(1)] = re.findall(r'"([^"]*)"', m.group(2))
    out = {}
    for t, strs in arrays.items():
        names = consts.get(t)
        if not names or len(names) != len(strs):
            continue
        out[t] = list(zip(names, strs))
    return out


ENUMS = load()


def enum_for(e, t, default=None):
    """The constant of enum t whose classes are all on element e, the one with
    the most classes; default when none is."""
    best, bestn = default, 0
    cls = set(e.cls)
    for name, s in ENUMS.get(t, []):
        want = set(s.split())
        if want and want <= cls and len(want) > bestn:
            best, bestn = "w." + name, len(want)
    return best


def enum_classes(t):
    return {c for _, s in ENUMS.get(t, []) for c in s.split()}


if __name__ == "__main__":
    print(len(ENUMS), "enums")
    for t in ("MarkSize", "RubricVariant", "DimSize", "ButtonVariant", "InputVariant", "FieldWidth"):
        print(t, ENUMS.get(t, "—")[:6] if ENUMS.get(t) else "—")
