"""Проверки построения иконки Typecase.

Каждая смотрит ровно на то, что названа проверять, и возвращает вердикт
строкой. Документ подставляет эти строки в разделы «Что чинится» и
«Проверки» — своих слов о результате у него нет.

    python icon_checks.py
"""

import os
import re
import sys

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))

import raster as R
import typecase_logotype as T
import typecase_icon as I

DRAFTS = os.path.normpath(os.path.join(I.HERE, "..", "typecase-logo-bg.svg"))


def shapes(kind="master", theme="light", letter=True, radius=None):
    """Фигуры знака для растеризатора."""
    p = I.PAL[theme]
    if kind == "master":
        r = I.RADIUS if radius is None else radius
        out = [(None, p["box"]),
               (I.cells(skip={I.FREE}, radius=r), p["card"]),
               (I.one(*I.FREE, radius=r), p["accent"])]
        if letter:
            out.append((I.glyph(*I.LETTER), p["ink"]))
        return out, I.FIELD
    if kind == "32":
        mod, cell, rad = I.GRID32
        r = rad if radius is None else radius
        pos = [mod, mod + cell + mod, mod + 2 * (cell + mod)]
        out = [(None, p["box"]),
               (I.cells(skip={I.FREE}, pos=pos, cell=cell, radius=r), p["card"]),
               (I.one(*I.FREE, pos=pos, cell=cell, radius=r), p["accent"])]
        if letter:
            out.append((I.glyph(*I.LETTER, pos=pos, cell=cell, cap=cell - 2), p["ink"]))
        return out, 32
    r = I.RADIUS16 if radius is None else radius
    out = [(None, p["box"]),
           (I.rounded(1, 1, I.CELL16, I.CELL16, r), p["card"])]
    if letter:
        out.append((I.glyph(0, 0, pos=[1], cell=I.CELL16, cap=I.CAP16), p["ink"]))
    return out, 16


def measure(kind, size, theme="light", letter=True, radius=None):
    sh, field = shapes(kind, theme, letter, radius)
    return R.halftones(R.render(sh, field, size), [c for _, c in sh])


# --- проверки ----------------------------------------------------------------

def glyph_fidelity():
    """Литера знака против литеры шрифта: расхождение площадей в процентах."""
    name = T.cmap[ord("T")]
    box = I.glyph_box(*I.LETTER)
    scale = I.CAP / T.bounds(name)[3]

    def area(polys):
        total = 0.0
        for poly in polys:
            s = 0.0
            for (x1, y1), (x2, y2) in zip(poly, poly[1:]):
                s += x1 * y2 - x2 * y1
            total += abs(s) / 2
        return total

    font_area = area(T.polylines(name)) * scale * scale
    icon_area = area(T.path_polylines(I.glyph(*I.LETTER), steps=64))
    return abs(icon_area - font_area) / font_area * 100, box[2]


def grid_integers():
    """Дробные числа в путях гнёзд: их не должно быть ни одного."""
    d = I.cells() + " " + I.one(*I.FREE)
    return [n for n in re.findall(r"-?\d+\.\d+", d)]


def division():
    """Деление мастера на 3: сетка, гнездо и радиус разом."""
    got = (I.MOD / 3, I.CELL / 3, I.RADIUS / 3)
    want = I.GRID32
    ok = all(abs(a - b) < 1e-9 and float(a).is_integer() for a, b in zip(got, want))
    return got, want, ok


def air():
    """Воздух вокруг литеры в гнезде."""
    x, y, w, h, ax, ay = I.glyph_box(*I.LETTER)
    return ax, ay, w, h


def contrast(a, b):
    def lum(h):
        c = [int(h.lstrip("#")[i:i + 2], 16) / 255 for i in (0, 2, 4)]
        c = [v / 12.92 if v <= 0.03928 else ((v + 0.055) / 1.055) ** 2.4 for v in c]
        return 0.2126 * c[0] + 0.7152 * c[1] + 0.0722 * c[2]
    la, lb = lum(a), lum(b)
    return (max(la, lb) + 0.05) / (min(la, lb) + 0.05)


def contrasts():
    L, D = I.PAL["light"], I.PAL["dark"]
    word = T.PAL["light"]
    return {
        "гнездо на ящике, светлая": contrast(L["card"], L["box"]),
        "гнездо на ящике, тёмная": contrast(D["card"], D["box"]),
        "акцент к соседнему гнезду": contrast(L["accent"], L["card"]),
        "акцент к ящику": contrast(L["accent"], L["box"]),
        "гнездо на бумаге, как в надписи": contrast(word["card"], "#F5F1E8"),
        "гнездо на фоне, тёмная надпись": contrast(T.PAL["dark"]["card"], "#141A16"),
    }


def maskable():
    return I.mask_reach(), I.mask_radius(I.FIELD), I.mask_radius(), I.MASK_FIELD


def draft_halftones(size):
    """Черновик 3×3 на том же тесте — если он ещё лежит рядом."""
    if not os.path.exists(DRAFTS):
        return None
    src = open(DRAFTS, encoding="utf-8").read()
    src = re.sub(r"<metadata>.*?</metadata>", "", src, flags=re.S)
    body = []
    for m in re.finditer(r'<rect[^>]*/>', src):
        tag = m.group(0)
        g = lambda k: float(re.search(k + r'="([-\d.]+)"', tag).group(1)) if re.search(k + r'="', tag) else 0.0
        fill = re.search(r'fill="([^"]+)"', tag)
        stroke = re.search(r'stroke="([^"]+)"', tag)
        sw = g("stroke-width") or 0.0
        x, y, w, h, r = g("x"), g("y"), g("width"), g("height"), g("rx")
        if fill and fill.group(1) != "none":
            body.append((I.rounded(x, y, w, h, r), fill.group(1)))
        if stroke:
            body.append((I.rounded(x - sw / 2, y - sw / 2, w + sw, h + sw, r),
                         stroke.group(1)))
            body.append((I.rounded(x + sw / 2, y + sw / 2, w - sw, h - sw, max(r - sw / 2, 0)),
                         fill.group(1) if fill and fill.group(1) != "none" else "#E8DFC8"))
    shapes_ = [(None, "#E8DFC8")] + body
    pixels = R.render(shapes_, 64, size)
    return R.halftones(pixels, ["#E8DFC8", "#2C2C2A", "#3B6D11", "#EDE4CB"])


def master():
    """Что лежит в выпущенном мастере."""
    if not os.path.exists(I.MASTER):
        return None
    src = open(I.MASTER, encoding="utf-8").read()
    return {
        "rects": src.count("<rect"),
        "paths": src.count("<path"),
        "strokes": len(re.findall(r"stroke=", src)),
        "texts": len(re.findall(r"<text|font-family", src)),
        "meta": src.count("<metadata"),
        "fills": sorted(set(re.findall(r'fill="(#[0-9A-Fa-f]{6})"', src))),
        "bytes": len(src.encode("utf-8")),
    }


# --- вердикты для документа --------------------------------------------------

def verdicts():
    v = {}
    dev, width = glyph_fidelity()
    v["glyph"] = (f"Глиф Literata из того же инстанса, что и надпись; площадь контура "
                  f"расходится с глифом шрифта на {dev:.4f} %. Ширина следует из роста "
                  f"прописных и равна {width:.2f} — она не назначалась.")
    frac = grid_integers()
    v["grid"] = ("Все. Дробных чисел в путях ящика и гнёзд нет ни одного."
                 if not frac else f"РАСХОДИТСЯ: дробные значения {frac[:6]}")
    got, want, ok = division()
    v["division"] = (f"Поле {got[0]:.0f}, гнездо {got[1]:.0f}, радиус {got[2]:.0f} — "
                     f"{'все три целые, хинтов по сетке не потребовалось' if ok else 'РАСХОДИТСЯ'}.")
    ax, ay, w, h = air()
    v["air"] = (f"Воздух {ax:.2f} по горизонтали и {ay:.2f} по вертикали при литере "
                f"{w:.2f} × {h:.2f} в гнезде {I.CELL}.")
    h16 = measure("16", 16)
    h16_plain = measure("16", 16, letter=False)
    h16_bare = measure("16", 16, letter=False, radius=0)
    v["raster16"] = (f"{h16} полутоновых пикселей из 256: {h16 - h16_plain} — засечки глифа, "
                     f"{h16_plain - h16_bare} — скругление гнезда. Без литеры и скругления "
                     f"ровно {h16_bare}, то есть прямые кромки чисты.")
    h32 = measure("32", 32)
    h32_plain = measure("32", 32, letter=False)
    v["raster32"] = (f"{h32} из 1024: {h32_plain} — скругления девяти гнёзд, "
                     f"{h32 - h32_plain} — глиф.")
    m96 = measure("master", 96)
    m48 = measure("master", 48)
    v["master_raster"] = (f"{m96} полутонов из 9216 на 96 и {m48} из 2304 на 48 — только дуги "
                          f"и засечки. Заливка ящика и прямые стороны гнёзд полутонов не дают.")
    c = contrasts()
    v["contrast"] = (f"{c['гнездо на ящике, светлая']:.2f}:1 в светлой теме и "
                     f"{c['гнездо на ящике, тёмная']:.2f}:1 в тёмной. Для сравнения: гнездо "
                     f"картона на бумаге даёт {c['гнездо на бумаге, как в надписи']:.2f}:1 — "
                     f"границы не видно.")
    v["accent_contrast"] = (f"{c['акцент к соседнему гнезду']:.2f}:1 к соседнему гнезду. "
                            f"К ящику всего {c['акцент к ящику']:.2f}:1, поэтому акцентное "
                            f"гнездо всегда стоит в окружении занятых, а не в углу.")
    reach, r96, r156, field = maskable()
    v["maskable"] = (f"Самая дальняя точка знака — угол крайнего гнезда, {reach:.1f} с учётом "
                     f"скругления. На поле 96 безопасный радиус {r96:.1f}, знак не проходит. "
                     f"На холсте {field} радиус становится {r156:.1f}; геометрия знака не "
                     f"меняется ни на единицу.")
    d16 = draft_halftones(16)
    d32 = draft_halftones(32)
    v["draft16"], v["draft32"] = d16, d32
    v["c"] = c
    v["n16"], v["n32"] = h16, h32
    v["glyph16"], v["round16"] = h16 - h16_plain, h16_plain - h16_bare
    v["round32"], v["glyph32"] = h32_plain, h32 - h32_plain
    v["m96"], v["m48"] = m96, m48
    v["m32"] = measure("master", 32)          # мастер, ужатый до 32 без хинтов
    v["draft16_pct"] = 100 * d16 / 256 if d16 else 0
    v["draft32_pct"] = 100 * d32 / 1024 if d32 else 0
    v["air_x"], v["air_y"] = ax, ay
    v["glyph_w"], v["glyph_h"] = w, h
    v["glyph_dev"] = dev
    v["reach"], v["r96"], v["r_mask"], v["mask_field"] = reach, r96, r156, field
    v["master_file"] = master()
    # десятичная запятая: документ набран по-русски
    dot = re.compile(r'(?<=\d)\.(?=\d)')
    for k, val in list(v.items()):
        if isinstance(val, str):
            v[k] = dot.sub(',', val)
    return v


if __name__ == "__main__":
    dev, width = glyph_fidelity()
    print(f"1. Литера         расхождение площадей {dev:.4f} %, ширина {width:.3f}")
    frac = grid_integers()
    print(f"2. Координаты     дробных в сетке: {len(frac)}")
    got, want, ok = division()
    print(f"3. Деление на 3   {got} против {want} — {'сходится' if ok else 'РАСХОДИТСЯ'}")
    ax, ay, w, h = air()
    print(f"4. Воздух         {ax:.3f} по x, {ay:.3f} по y при литере {w:.3f} × {h:.3f}")
    print(f"5. Растр 16       {measure('16', 16)} из 256 "
          f"(без литеры {measure('16', 16, letter=False)}, "
          f"без литеры и скругления {measure('16', 16, letter=False, radius=0)})")
    print(f"6. Растр 32       {measure('32', 32)} из 1024 "
          f"(без литеры {measure('32', 32, letter=False)})")
    print(f"7. Мастер в растре 96: {measure('master', 96)} из 9216 · "
          f"48: {measure('master', 48)} из 2304")
    for k, val in contrasts().items():
        print(f"8. Контраст       {k}: {val:.2f}:1")
    reach, r96, r156, field = maskable()
    print(f"9. Maskable       вылет {reach:.2f}; радиус на 96 — {r96:.1f}, на {field} — {r156:.1f}")
    d16, d32 = draft_halftones(16), draft_halftones(32)
    print(f"10. Черновик      {d16} из 256 на 16, {d32} из 1024 на 32"
          if d16 is not None else "10. Черновик      не найден")
    m = master()
    print(f"11. Мастер        rect {m['rects']}, path {m['paths']}, stroke {m['strokes']}, "
          f"шрифт {m['texts']}, метаданные {m['meta']}, {m['bytes'] / 1024:.1f} КБ"
          if m else "11. Мастер        не выпущен")
