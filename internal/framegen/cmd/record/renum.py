"""Renumber records after frames changed: a record's element is found by its
markup in the HEAD mockup and looked up among the same class in the new one.

    python internal/framegen/cmd/record/renum.py s-notes-4 s-notes-20

Run from the module root, before the change is committed (it compares with
HEAD); check its guesses with the tests — a frame that lost or gained a
look-alike element near the record can be renumbered wrong.
"""
import re, subprocess, sys, glob
sys.stdout.reconfigure(encoding="utf-8")
ROOT = "./"
old = subprocess.run(["git", "-C", ROOT, "show", "HEAD:docs/visual/carrel-ui-mockups.html"], capture_output=True).stdout.decode("utf-8")
new = open(ROOT + "docs/visual/carrel-ui-mockups.html", encoding="utf-8").read()
frames = set(sys.argv[1:])


def frame(s, fid):
    i = s.index('data-frame="%s"' % fid)
    return s[i:s.index('</figure>', i)]


def elems(f, cls):
    out = []
    for m in re.finditer(r'<(\w+)\b[^>]*\bclass="([^"]*)"[^>]*>', f):
        if cls in m.group(2).split():
            # the element's own markup up to its end, roughly: 400 chars
            out.append(f[m.start():m.start() + 400])
    return out


changed = 0
for path in glob.glob(ROOT + "internal/refs/*.go"):
    t = open(path, encoding="utf-8").read()

    def fix(m):
        global changed
        fid, cls, n = m.group(1), m.group(2), int(m.group(3))
        if frames and fid not in frames:
            return m.group(0)
        try:
            o = elems(frame(old, fid), cls)
            w = elems(frame(new, fid), cls)
        except ValueError:
            return m.group(0)
        if n > len(o):
            return m.group(0)
        sig = o[n - 1]
        if n <= len(w) and w[n - 1] == sig:
            return m.group(0)
        hits = [k + 1 for k, x in enumerate(w) if x == sig]
        if not hits:
            # a shorter signature: the opening tag and 120 chars
            short = sig[:160]
            hits = [k + 1 for k, x in enumerate(w) if x[:160] == short]
        if len(hits) >= 1:
            best = min(hits, key=lambda h: abs(h - n))
            if best != n:
                changed += 1
                print("%s %s %s: %d -> %d" % (path.split("/")[-1], fid, cls, n, best))
                return '"%s", "%s", %d)' % (fid, cls, best)
        return m.group(0)

    t2 = re.sub(r'"([a-z0-9-]+)", "([a-z0-9_-]+)", (\d+)\)', fix, t)
    if t2 != t:
        open(path, "w", encoding="utf-8", newline="").write(t2)
print("changed", changed)
