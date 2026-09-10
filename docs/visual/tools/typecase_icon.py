"""Построение иконки Typecase.

Знак — наборная касса сверху: ящик, девять гнёзд, в одном литера, одно
свободно. Сетка назначена (квадрат 96, модуль 6), литера не назначена: она
берётся у того же инстанса Literata, что и надпись, и уменьшается ровно
вчетверо. Всё, что относится к литере, приходит из `typecase_logotype`, чтобы
знак и надпись не могли разойтись.

Мастер пересобирается так:
    python build_typecase_icon.py
"""

import os

from fontTools.pens.svgPathPen import SVGPathPen
from fontTools.pens.transformPen import TransformPen
from fontTools.misc.transform import Transform

import typecase_logotype as T

HERE = os.path.dirname(os.path.abspath(__file__))
MASTER = os.path.normpath(os.path.join(HERE, "..", "typecase-mark.svg"))
REFERENCE = os.path.normpath(os.path.join(HERE, "..", "typecase-icon.html"))

# --- сетка -------------------------------------------------------------------

FIELD = 96                      # квадрат мастера
MOD = 6                         # модуль: внешнее поле и перегородка разом
CELL = 24                       # гнездо
RADIUS = MOD // 2               # скругление гнезда — половина модуля
POS = [MOD, MOD + CELL + MOD, MOD + 2 * (CELL + MOD)]   # 6, 36, 66

LETTER = (0, 1)                 # занятое гнездо: левое среднее
FREE = (2, 1)                   # свободное гнездо: правое среднее

DIV = 4.0                       # во сколько раз литера мельче, чем в надписи
CAP = T.CAP / DIV               # рост прописных в системе знака: 16.5
SCALE = T.SCALE / DIV           # из единиц шрифта в единицы знака

# хинтованные версии: (поле и перегородка, гнездо, радиус)
GRID32 = (2, 8, 1)              # мастер ÷ 3, деление точное
CELL16 = 14                     # 16-я версия: одно гнездо во весь квадрат
RADIUS16 = 1
CAP16 = 10.0                    # рост прописных на 16 px

PAL = {
    "light": dict(box=T.PAL["light"]["ink"], card=T.PAL["light"]["card"],
                  ink=T.PAL["light"]["ink"], accent=T.PAL["light"]["accent"]),
    # в тёмной теме самый тёмный токен — bg, самый светлый — ink
    "dark": dict(box="#141A16", card=T.PAL["dark"]["ink"],
                 ink="#141A16", accent=T.PAL["dark"]["accent"]),
}

# --- фигуры ------------------------------------------------------------------

def _n(v):
    return f"{v:.3f}".rstrip("0").rstrip(".") if isinstance(v, float) else str(v)


def rounded(x, y, w, h, r):
    """Замкнутый прямоугольник со скруглением, обход по часовой."""
    if r <= 0:
        return f"M{_n(x)} {_n(y)}H{_n(x + w)}V{_n(y + h)}H{_n(x)}Z"
    return (f"M{_n(x + r)} {_n(y)}H{_n(x + w - r)}A{_n(r)} {_n(r)} 0 0 1 {_n(x + w)} {_n(y + r)}"
            f"V{_n(y + h - r)}A{_n(r)} {_n(r)} 0 0 1 {_n(x + w - r)} {_n(y + h)}"
            f"H{_n(x + r)}A{_n(r)} {_n(r)} 0 0 1 {_n(x)} {_n(y + h - r)}"
            f"V{_n(y + r)}A{_n(r)} {_n(r)} 0 0 1 {_n(x + r)} {_n(y)}Z")


def cells(skip=(), pos=None, cell=CELL, radius=RADIUS):
    """Гнёзда одним контуром; skip — те, что рисуются отдельно."""
    pos = POS if pos is None else pos
    out = []
    for row, y in enumerate(pos):
        for col, x in enumerate(pos):
            if (col, row) in skip:
                continue
            out.append(rounded(x, y, cell, cell, radius))
    return " ".join(out)


def one(col, row, pos=None, cell=CELL, radius=RADIUS):
    pos = POS if pos is None else pos
    return rounded(pos[col], pos[row], cell, cell, radius)


def glyph(col, row, pos=None, cell=CELL, cap=CAP, char="T"):
    """Литера в гнезде: воздух поровну с четырёх сторон.

    Пропорции глифа не трогаются — назначается только рост прописных, а из
    него следует и ширина, и положение.
    """
    pos = POS if pos is None else pos
    name = T.cmap[ord(char)]
    scale = cap / T.bounds(name)[3]          # верх литеры в единицах шрифта
    b = T.bounds(name)
    ink_w = (b[2] - b[0]) * scale
    x0 = pos[col] + (cell - ink_w) / 2 - b[0] * scale
    baseline = pos[row] + (cell - cap) / 2 + cap
    spen = SVGPathPen(T.gs, ntos=lambda v: f"{v:.3f}".rstrip("0").rstrip("."))
    T.gs[name].draw(TransformPen(spen, Transform(scale, 0, 0, -scale, x0, baseline)))
    return spen.getCommands()


def glyph_box(col, row, pos=None, cell=CELL, cap=CAP, char="T"):
    """Габарит литеры в единицах знака: (x, y, ширина, высота, воздух по x, по y)."""
    pos = POS if pos is None else pos
    name = T.cmap[ord(char)]
    b = T.bounds(name)
    scale = cap / b[3]
    ink_w = (b[2] - b[0]) * scale
    x = pos[col] + (cell - ink_w) / 2
    y = pos[row] + (cell - cap) / 2
    return x, y, ink_w, cap, (cell - ink_w) / 2, (cell - cap) / 2


# --- знак --------------------------------------------------------------------

def inner(theme="light", box=True, view=None):
    """Содержимое знака без обёртки <svg>: ящик, гнёзда, свободное гнездо, литера."""
    p = PAL[theme]
    out = []
    if box:
        if view is None:
            out.append(f'<rect width="{FIELD}" height="{FIELD}" fill="{p["box"]}"/>')
        else:
            x, y, w, h = view
            out.append(f'<rect x="{_n(x)}" y="{_n(y)}" width="{_n(w)}" height="{_n(h)}" '
                       f'fill="{p["box"]}"/>')
    out.append(f'<path fill="{p["card"]}" d="{cells(skip={FREE})}"/>')
    out.append(f'<path fill="{p["accent"]}" d="{one(*FREE)}"/>')
    out.append(f'<path fill="{p["ink"]}" d="{glyph(*LETTER)}"/>')
    return "".join(out)


def inner32(theme="light", box=True):
    """32 px: мастер ÷ 3. Сетка делится точно, литера пересаживается по росту."""
    p = PAL[theme]
    mod, cell, radius = GRID32
    pos = [mod, mod + cell + mod, mod + 2 * (cell + mod)]
    cap = cell - 2                                  # воздух по единице сверху и снизу
    out = []
    if box:
        out.append(f'<rect width="32" height="32" fill="{p["box"]}"/>')
    out.append(f'<path fill="{p["card"]}" d="{cells(skip={FREE}, pos=pos, cell=cell, radius=radius)}"/>')
    out.append(f'<path fill="{p["accent"]}" d="{one(*FREE, pos=pos, cell=cell, radius=radius)}"/>')
    out.append(f'<path fill="{p["ink"]}" d="{glyph(*LETTER, pos=pos, cell=cell, cap=cap)}"/>')
    return "".join(out)


def inner16(theme="light", box=True):
    """16 px: редукция до одного гнезда. Свободного гнезда здесь нет."""
    p = PAL[theme]
    pos = [1]
    out = []
    if box:
        out.append(f'<rect width="16" height="16" fill="{p["box"]}"/>')
    out.append(f'<path fill="{p["card"]}" d="{rounded(1, 1, CELL16, CELL16, RADIUS16)}"/>')
    out.append(f'<path fill="{p["ink"]}" d="{glyph(0, 0, pos=pos, cell=CELL16, cap=CAP16)}"/>')
    return "".join(out)


def svg(theme="light", view=None):
    """Мастер целиком."""
    if view is None:
        vb = f"0 0 {FIELD} {FIELD}"
        body = inner(theme)
    else:
        x, y, w, h = view
        vb = f"{_n(x)} {_n(y)} {_n(w)} {_n(h)}"
        body = inner(theme, view=view)
    return ('<svg xmlns="http://www.w3.org/2000/svg" viewBox="' + vb + '" '
            'role="img" aria-label="Typecase">\n'
            + "\n".join("  " + line for line in body.replace("><", ">\n<").split("\n"))
            + "\n</svg>\n")


# --- маскируемая версия ------------------------------------------------------

MASK_FIELD = 156                                  # ближайшее поле, кратное модулю
MASK_VIEW = (-(MASK_FIELD - FIELD) / 2, -(MASK_FIELD - FIELD) / 2, MASK_FIELD, MASK_FIELD)


def mask_reach():
    """Самая дальняя точка знака от центра квадрата, с учётом скругления."""
    cx = cy = FIELD / 2
    corner = POS[-1] + CELL - RADIUS               # центр угловой дуги
    dx, dy = corner - cx, corner - cy
    return (dx * dx + dy * dy) ** 0.5 + RADIUS


def mask_radius(field=MASK_FIELD):
    """Безопасный радиус маски: 40 % холста."""
    return 0.4 * field


if __name__ == "__main__":
    print("сетка       ", MOD, CELL, MOD, "· гнёзда на", POS)
    print("литера      ", f"рост {CAP}, ширина {glyph_box(*LETTER)[2]:.3f}")
    print("воздух      ", f"{glyph_box(*LETTER)[4]:.3f} по x, {glyph_box(*LETTER)[5]:.3f} по y")
    print("maskable    ", f"вылет {mask_reach():.2f}, радиус {mask_radius():.1f}")
