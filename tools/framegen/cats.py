import re, sys
from collections import Counter, defaultdict
c=Counter(); ex=defaultdict(list)
for ln in open(sys.argv[1] if len(sys.argv) > 1 else "diffs.txt", encoding="utf-8"):
    ln=ln.strip()
    if ": " not in ln: continue
    fid,rest=ln.split(": ",1)
    m=re.match(r"(.*?)(: (?:class|tag|frame has|text|attributes).*)$",rest)
    path,msg=(m.group(1),m.group(2)) if m else (rest,"")
    last=" > ".join(re.sub(r"\[\d+\]","",p) for p in path.split(" > ")[-2:])
    kind=re.sub(r'"[^"]*"','"…"',msg) if "class" in msg or "attributes" in msg else msg
    kind=re.sub(r"\d+","N",kind)
    key=last+" :: "+kind[:150]
    c[key]+=1; ex[key].append(fid)
for k,v in c.most_common(): print(v,k,"|",ex[k][0])
