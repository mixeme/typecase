"""Add rows to docs/visual/allowed-diffs.md for frames the owner accepted.

    python internal/framegen/cmd/record/allow.py 2026-09-28 "why" s-tasks-2 s-contacts-6 ...

A row per frame and width (1351, 820), with the frame's section and caption
read from the mockup. A frame that already has rows gets the reason added.
"""
import re, sys

date, why, frames = sys.argv[1], sys.argv[2], sys.argv[3:]
m = open("docs/visual/carrel-ui-mockups.html", encoding="utf-8").read()
heads = [(x.start(), re.sub("<[^>]+>", "", x.group(1)).split()[0]) for x in re.finditer(r'<h2 id="[^"]*"[^>]*>(.*?)</h2>', m)]
p = "docs/visual/allowed-diffs.md"
s = open(p, encoding="utf-8").read()
for fid in frames:
    i = m.find('data-frame="%s"' % fid)
    if i < 0:
        sys.exit("no frame " + fid)
    sec = [t for pos, t in heads if pos < i][-1]
    lab = re.search(r'<div class="frame-label"><span>(.*?)</span>', m[i:i + 2000], re.S)
    cap = re.sub(r"<[^>]+>", "", lab.group(1)).strip()
    have = [ln for ln in s.split("\n") if ln.startswith("| %s |" % fid)]
    if have:
        for ln in have:
            cells = ln.split(" | ")
            cells[2] = "человек %s" % date
            cells[3] = cells[3] + "; " + why
            s = s.replace(ln, " | ".join(cells))
        continue
    for w in ("1351", "820"):
        s = s.rstrip("\n") + "\n| %s | %s | человек %s | %s | %s | %s |\n" % (fid, w, date, why, sec, cap)
open(p, "w", encoding="utf-8", newline="\n").write(s)
print("ok", len(frames))
