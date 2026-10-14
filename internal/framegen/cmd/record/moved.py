"""Renumber the reference records whose fragment moved in its frame.

    go test ./internal/refs 2>&1 | python internal/framegen/cmd/record/moved.py

Reads the test output: a record whose markup is the frame's #k now (the
comparison says «renumber the record to #k») points at #k. All of them are
renumbered at once, so two records that swapped places stay apart. Run from
the module root.
"""
import glob, re, sys

moves = {}
for line in sys.stdin:
    m = re.search(r'(\S+) \.(\S+) #(\d+): .*renumber the record to #(\d+)', line)
    if m:
        moves[(m.group(1), m.group(2), m.group(3))] = m.group(4)
for path in glob.glob("internal/refs/*.go"):
    s = open(path, encoding="utf-8").read()
    t = s
    for (frame, cls, n), k in moves.items():
        t = t.replace('"%s", "%s", %s)' % (frame, cls, n), '"%s", "%s", \x00%s)' % (frame, cls, k))
    t = t.replace("\x00", "")
    if t != s:
        open(path, "w", encoding="utf-8", newline="\n").write(t)
for (frame, cls, n), k in sorted(moves.items()):
    print("renumbered %s .%s #%s -> #%s" % (frame, cls, n, k))
