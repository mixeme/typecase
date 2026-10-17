"""Apply the refs test's «renumber the record to #N» hints, all at once.

    python internal/framegen/cmd/record/hintfix.py

Runs `go test ./internal/refs -run TestWidgetFragments`, takes every hint
«<frame> .<class> #<n>: … renumber the record to #<m>» and moves that record
to #m in internal/refs, by a placeholder pass so that swaps do not collide.
Run it again until it says «0 hints»; what is left needs refresh.py or a
hand. Run from the module root, after `go generate ./internal/screens`.
"""
import re,subprocess,glob,sys
out=subprocess.run(['go','test','./internal/refs','-run','TestWidgetFragments'],capture_output=True,text=True,encoding='utf-8').stdout
hints=re.findall(r'refs_test.go:34: (\S+) \.(\S+) #(\d+):.*?renumber the record to #(\d+)',out)
print(len(hints),'hints')
files={p:open(p,encoding='utf-8').read() for p in glob.glob('internal/refs/*.go')}
for fr,cls,a,b in hints:
    tail='"%s", "%s", %s)'%(fr,cls,a)
    for p,t in files.items():
        if t.count(tail)==1:
            files[p]=t.replace(tail,'"%s", "%s", @@%s)'%(fr,cls,b)); break
    else: print('not found / ambiguous',tail)
for p,t in files.items():
    open(p,'w',encoding='utf-8',newline='\n').write(re.sub(r'@@(\d+)\)',r'\1)',t))
