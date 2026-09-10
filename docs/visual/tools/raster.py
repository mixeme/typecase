"""Растеризатор для проверок знака.

Считает долю покрытия пикселя фигурой по правилу «центр пробы внутри
контура» и чётности, как fill-rule evenodd. Готовые растеризаторы здесь не
годятся: PIL заливает полигон вместе с обеими границами, и кромка, лежащая
ровно на пикселе, получает лишнюю пробу — проверка «полутонов ноль» тогда
никогда не сходится.
"""

import typecase_logotype as T

SS = 8          # проб на пиксель по каждой оси


def edges(d, steps=64):
    """Рёбра контуров пути: (y0, y1, x0, наклон)."""
    out = []
    for poly in T.path_polylines(d, steps=steps):
        for (x1, y1), (x2, y2) in zip(poly, poly[1:]):
            if y1 == y2:
                continue
            out.append((min(y1, y2), max(y1, y2), x1, y1, (x2 - x1) / (y2 - y1)))
    return out


def coverage(d, field, size, ss=SS):
    """Доля покрытия каждого пикселя фигурой: список size × size из 0…1."""
    segs = edges(d)
    cov = [0] * (size * size)
    if not segs:
        return [0.0] * (size * size)
    step = field / (size * ss)                      # единиц знака на пробу
    for row in range(size * ss):
        yc = (row + 0.5) * step
        xs = []
        for ylo, yhi, x1, y1, slope in segs:
            if ylo <= yc < yhi:
                xs.append(x1 + (yc - y1) * slope)
        if not xs:
            continue
        xs.sort()
        out_row = (row // ss) * size
        for k in range(0, len(xs) - 1, 2):
            a, b = xs[k], xs[k + 1]
            # пробы, чьи центры попали внутрь интервала
            i0 = int(a / step - 0.5) + 1
            i1 = int(b / step - 0.5)
            if i1 < i0:
                continue
            i0 = max(i0, 0)
            i1 = min(i1, size * ss - 1)
            for i in range(i0, i1 + 1):
                cov[out_row + i // ss] += 1
    total = ss * ss
    return [c / total for c in cov]


def _rgb(h):
    return tuple(int(h.lstrip("#")[i:i + 2], 16) for i in (0, 2, 4))


def render(shapes, field, size):
    """Знак в растре: shapes — [(путь, цвет)], первый элемент — фон (путь None)."""
    base = _rgb(shapes[0][1])
    img = [list(base) for _ in range(size * size)]
    for d, color in shapes[1:]:
        c = _rgb(color)
        for i, a in enumerate(coverage(d, field, size)):
            if a <= 0:
                continue
            px = img[i]
            for k in range(3):
                px[k] = px[k] * (1 - a) + c[k] * a
    return [tuple(round(v) for v in px) for px in img]


def halftones(pixels, palette, tol=2):
    """Пиксели, не совпавшие ни с одним цветом палитры, — размытые кромки."""
    pal = [_rgb(c) for c in palette]
    return sum(1 for px in pixels
               if not any(all(abs(px[k] - q[k]) <= tol for k in range(3)) for q in pal))
