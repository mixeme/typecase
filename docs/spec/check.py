# -*- coding: utf-8 -*-
"""Сверка спецификации Typecase.

Запуск из корня модуля: python docs/spec/check.py
(красное — ненулевой код возврата)
"""
import io, os, re, sys

ROOT = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
SPEC = os.path.join(ROOT, "docs", "spec")
SPEC_FILE = os.path.join(SPEC, "typecase.md")

def read(p):
    return io.open(p, encoding="utf-8").read()

uilib = read(SPEC_FILE)

fails = []
def check(name, bad, note=""):
    if bad:
        fails.append((name, bad, note))
        print("КРАСНОЕ  %s: %d" % (name, len(bad)))
        for b in list(bad)[:12]:
            print("         -", b)
    else:
        print("зелёное  %s" % name)

# 4. Library independence: no cross-file links or domain words
cross = re.findall(r'\((?:carrel-(?:common|client|server|desktop))\.md\)', uilib)
domain = []
for w in ["DAV", "vCard", "iCal", "контакт", "календар", "заметк", "коллекци", "аккаунт", "Carrel"]:
    for m in re.finditer(w, uilib, re.I):
        line = uilib[:m.start()].count("\n") + 1
        if line > 90:  # §L-2.2 lists them as negative examples
            domain.append("%s:%d" % (w, line))
check("4. typecase.md не ссылается на carrel-*.md", cross)
check("4b. typecase.md без доменных слов", sorted(set(domain)))

print()
print("ИТОГ: %d красных проверок" % len(fails))
sys.exit(1 if fails else 0)
