"""Remove reference records whose element is gone from its frame.

    python internal/framegen/cmd/record/remove.py s-backup-3:tc-msg:1 'account-16:tc-table:1:Recovered' ...

Each argument is frame:class:n, and an optional text the record must
contain when two records end the same way. A record is the top-level call
(a line starting with one tab, a name and «(») whose last line ends with
«"frame", "class", n),». Run from the module root; gofmt afterwards. After
a frame lost an element, remove its record, then let TestShapesCovered say
which shape needs a record again (add.py).
"""
import glob, re, sys
for arg in sys.argv[1:]:
    parts = arg.split(":", 3)
    frame, cls, n = parts[:3]
    must = parts[3] if len(parts) > 3 else None
    tail = '"%s", "%s", %s),' % (frame, cls, n)
    done = False
    for path in glob.glob("internal/refs/*.go"):
        lines = open(path, encoding="utf-8", newline="").read().split("\n")
        for i, l in enumerate(lines):
            if l.rstrip("\r").endswith(tail):
                j = i
                while not re.match(r"\t[A-Za-z_]+\(", lines[j]):
                    j -= 1
                rec = "\n".join(lines[j:i + 1])
                if must and must not in rec:
                    continue
                del lines[j:i + 1]
                open(path, "w", encoding="utf-8", newline="").write("\n".join(lines))
                print("removed", arg, "from", path)
                done = True
                break
        if done:
            break
    if not done:
        print("NOT FOUND", arg)
