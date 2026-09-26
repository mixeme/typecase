_RAW = None


def _frames():
    global _RAW
    if _RAW is None:
        p = P()
        p.feed(open("docs/visual/carrel-ui-mockups.html", encoding="utf-8").read())
        _RAW = {}

        def walk(e):
            for k in e.els():
                if k.attrs.get("data-frame"):
                    _RAW[k.attrs["data-frame"]] = k
                else:
                    walk(k)
        walk(p.root)
    return _RAW


def raw_block(frame, cls, n):
    found = []

    def walk(e):
        for k in e.els():
            if k.has("frame-label"):
                continue
            if k.has(cls):
                found.append(k)
            walk(k)
    walk(_frames()[frame])
    return found[n - 1]
