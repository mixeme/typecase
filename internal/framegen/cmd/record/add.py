"""Add reference records for forms that have none, drafted from their frames.

    python internal/framegen/cmd/record/add.py s-dup-7:tc-btn:3:disabled s-files-16:tc-tree:1:plain ...

For each frame:class:n:state it runs `go run ./internal/framegen/cmd/record`,
finds a record of the same widget type in internal/refs, and puts the new
record before it, with the same helper. Run from the module root; gofmt
afterwards.
"""
import glob, re, subprocess, sys

for arg in sys.argv[1:]:
    frame, cls, n, state = arg.split(":")
    r = subprocess.run(["go", "run", "./internal/framegen/cmd/record", frame, cls, n],
                       capture_output=True, text=True, encoding="utf-8")
    if r.returncode != 0:
        sys.exit("%s: %s" % (arg, r.stderr.strip()))
    value = r.stdout.strip()
    typ = re.match(r"(w\.[A-Za-z]+)\{", value).group(1)
    done = False
    for path in sorted(glob.glob("internal/refs/*.go")):
        s = open(path, encoding="utf-8").read()
        m = re.search(r'\n\t([A-Za-z]+)\("[^"]*", ' + re.escape(typ) + r"\{", s)
        if not m:
            continue
        rec = '\n\t%s("%s", %s, "%s", "%s", %s),' % (m.group(1), state, value, frame, cls, n)
        s = s[:m.start()] + rec + s[m.start():]
        open(path, "w", encoding="utf-8", newline="\n").write(s)
        print("added", arg, "as", m.group(1), "in", path)
        done = True
        break
    if not done:
        sys.exit("%s: no record of %s to stand beside" % (arg, typ))
