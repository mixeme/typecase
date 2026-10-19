"""Drop «, Place: w.Grown» from records drafted by add.py or refresh.py.

    python internal/framegen/cmd/record/ungrow.py s-backup-2:tc-msg:4 ...

The fragment compare does not see tc-grow (a place, not the widget's own
class), so a drafted record with Place: w.Grown fails it («class "tc-msg",
widget "tc-grow tc-msg"»). Run from the module root after add.py or
refresh.py when that failure shows.
"""
import glob, re, sys
# drop «, Place: w.Grown» from the records named frame:class:n
for arg in sys.argv[1:]:
    frame, cls, n = arg.split(":")
    tail = '"%s", "%s", %s),' % (frame, cls, n)
    for path in glob.glob("internal/refs/*.go"):
        lines = open(path, encoding="utf-8", newline="").read().split("\n")
        hit = False
        for i, l in enumerate(lines):
            if l.rstrip("\r").endswith(tail):
                j = i
                while not re.match(r"\t[A-Za-z_]+\(", lines[j]):
                    j -= 1
                rec = "\n".join(lines[j:i + 1])
                # only the record's own Place: the last «, Place: w.Grown}» before the tail
                k = rec.rfind(", Place: w.Grown")
                if k >= 0:
                    rec = rec[:k] + rec[k + len(", Place: w.Grown"):]
                    lines[j:i + 1] = rec.split("\n")
                    hit = True
                break
        if hit:
            open(path, "w", encoding="utf-8", newline="").write("\n".join(lines))
            print("ungrown", arg)
            break
