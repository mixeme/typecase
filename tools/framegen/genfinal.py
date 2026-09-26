"""Regenerate the container records that waited for the last family."""
import os

os.environ["GENCARDS_MAIN"] = "0"
HERE = os.path.dirname(os.path.abspath(__file__))
exec(open(os.path.join(HERE, "gencards.py"), encoding="utf-8").read())


def shellw2(e):
    f = []
    t = kid(e, "tc-top")
    if t is not None:
        tt = top(t)
        f.append("Top: &" + (tt if tt is not None else top2(t)))
    o = kid(e, "tc-overlay")
    if o is not None:
        f.append("Body: " + overlayw(o))
        return "w.Shell{" + ", ".join(f) + "}"
    return shellw(e)


def run(fams, fn_name, statef):
    out, waiting = [], []
    for cls, fn in fams:
        out.append("\n\t// " + cls)
        for frame, n, e in shapes(cls):
            try:
                expr = fn(e)
            except Wait:
                waiting.append("ref(%s, %s, %d)" % (q(frame), q(cls), n))
                continue
            if expr is None:
                waiting.append("ref(%s, %s, %d)" % (q(frame), q(cls), n))
                continue
            out.append("\t%s(%s, %s, %s, %s, %d)," % (fn_name, q(statef(cls, e, expr)), expr, q(frame), q(cls), n))
    return out, waiting


def mstate2(cls, e, expr):
    if cls == "tc-shell" and "w.Overlay" in expr:
        return "overlay"
    if cls == "tc-side":
        for k, v in [("SideZone", "zone"), ("SideDrop", "drop"), ("SideActs", "acts")]:
            if k in expr:
                return v
        return "plain"
    return mstate(cls, e, expr)


o1, w1 = run([("tc-main", main), ("tc-side", side), ("tc-shell", shellw2)], "shell", mstate2)
open(os.path.join(HERE, "main2_out.txt"), "w", encoding="utf-8").write("\n".join(o1) + "\n\n// WAIT " + ", ".join(w1) + "\n")
o2, w2 = run([("tc-sheet", sheet), ("tc-scrim", scrim), ("tc-overlay", overlayw), ("tc-dialog", dialog), ("tc-auth", auth)], "overlay",
             lambda cls, e, expr: ostate(cls, expr))
open(os.path.join(HERE, "sheet_out.txt"), "w", encoding="utf-8").write("\n".join(o2) + "\n\n// WAIT " + ", ".join(w2) + "\n")
print("waits:", len(w1), len(w2))
