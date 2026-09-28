"""Redraft reference records from their frames.

    python internal/framegen/cmd/record/refresh.py s-contacts-6:tc-danger:1 s-notes-2:tc-tree:1 ...

For each frame:class:n it runs `go run ./internal/framegen/cmd/record` and puts
the expression in place of the value of the record in internal/refs that
points at that fragment. Run from the module root; gofmt afterwards.
"""
import glob, re, subprocess, sys

for arg in sys.argv[1:]:
    frame, cls, n = arg.split(":")
    expr = subprocess.run(["go", "run", "./internal/framegen/cmd/record", frame, cls, n],
                          capture_output=True, text=True, encoding="utf-8")
    if expr.returncode != 0:
        sys.exit("%s: %s" % (arg, expr.stderr.strip()))
    value = expr.stdout.strip()
    tail = '"%s", "%s", %s)' % (frame, cls, n)
    hit = False
    for path in glob.glob("internal/refs/*.go"):
        s = open(path, encoding="utf-8").read()
        k = s.find(", " + tail)
        if k < 0:
            continue
        starts = [m.end() for m in re.finditer(r'\n\t[A-Za-z]+\("[^"]*", ', s[:k])]
        if not starts:
            sys.exit("%s: no record start before it in %s" % (arg, path))
        s = s[:starts[-1]] + value + s[k:]
        open(path, "w", encoding="utf-8", newline="\n").write(s)
        print("redrafted", arg, "in", path)
        hit = True
        break
    if not hit:
        sys.exit("%s: no record points at it" % arg)
