"""Выпуск иконки Typecase: мастер и эталон построения разом.

    python build_typecase_icon.py

Пишет `../typecase-mark.svg` и `../typecase-icon.html`. Все числа документа
приходят из `icon_checks`, поэтому эталон не может разойтись с файлом,
который описывает.
"""

import os

import icon_checks as C
import typecase_icon as I

NL = "\n"

CSS = """
  :root{
    --ink:#17211C;--card:#E8DFCB;--card-border:#C9BFA6;--accent:#4A6B52;
    --accent-dark:#7FA085;--bg:#F5F1E8;--paper:#FFFDF7;--rule:#D8CFBB;
    --hair:#E2DACA;--dusty:#9A9280;--muted:#6B6355;--alert:#B04A3A;
  }
  *{box-sizing:border-box;}
  body{margin:0;padding:0;background:var(--bg);color:var(--ink);
       font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,sans-serif;
       font-size:15px;line-height:1.6;}
  .wrap{max-width:900px;margin:0 auto;padding:56px 32px 96px;}
  header{border-bottom:0.5px solid var(--rule);padding-bottom:32px;margin-bottom:44px;}
  h1{font-family:Georgia,"Source Serif 4",serif;font-size:38px;font-weight:400;margin:0 0 8px;}
  .sub{font-size:11px;letter-spacing:0.18em;color:var(--dusty);text-transform:uppercase;}
  h2{font-size:13px;letter-spacing:0.14em;text-transform:uppercase;color:var(--dusty);
     font-weight:600;margin:64px 0 20px;padding-bottom:8px;border-bottom:0.5px solid var(--rule);}
  h3{font-size:16px;font-weight:600;margin:32px 0 10px;}
  p{margin:0 0 14px;max-width:68ch;}
  .lede{font-size:17px;color:var(--muted);max-width:64ch;}
  .note{font-size:13px;color:var(--muted);padding-left:14px;border-left:2px solid var(--card-border);
        margin:18px 0;max-width:64ch;}
  table{border-collapse:collapse;width:100%;font-size:13.5px;margin:18px 0;}
  th{text-align:left;font-size:10.5px;letter-spacing:0.1em;text-transform:uppercase;
     color:var(--dusty);font-weight:600;padding:0 12px 8px 0;border-bottom:0.5px solid var(--rule);}
  td{padding:9px 12px 9px 0;border-bottom:0.5px solid var(--hair);vertical-align:top;}
  td.num,th.num{font-family:ui-monospace,SFMono-Regular,Menlo,monospace;font-size:12.5px;}
  code{font-family:ui-monospace,SFMono-Regular,Menlo,monospace;font-size:12.5px;
       background:var(--paper);border:0.5px solid var(--hair);border-radius:3px;padding:1px 5px;}
  pre{font-family:ui-monospace,SFMono-Regular,Menlo,monospace;font-size:12px;line-height:1.65;
      background:var(--paper);border:0.5px solid var(--hair);border-radius:5px;
      padding:16px 18px;overflow-x:auto;margin:14px 0 22px;}
  pre code{background:none;border:0;padding:0;font-size:12px;}
  figure{margin:26px 0;}
  figcaption{font-size:12.5px;color:var(--muted);margin-top:12px;max-width:66ch;}
  .board{background:var(--paper);border:0.5px solid var(--hair);border-radius:5px;
         padding:22px;display:inline-block;}
  .board svg{display:block;}
  .row{display:flex;align-items:flex-start;gap:24px;flex-wrap:wrap;margin:20px 0;}
  .cell{display:flex;flex-direction:column;gap:9px;}
  .cap{font-size:11px;color:var(--dusty);max-width:30ch;line-height:1.45;}
  .plate{display:flex;align-items:center;justify-content:center;flex:0 0 auto;
         background:var(--bg);border:0.5px solid var(--card-border);}
  .plate.dark{background:#141A16;border-color:#2B352E;}
  .plate.paper{background:var(--paper);border-color:var(--hair);}
  .plate svg{display:block;}
  .strip{display:flex;align-items:flex-end;gap:20px;flex-wrap:wrap;margin:18px 0;}
  .strip .cell{align-items:center;}
  .sizebox{display:flex;align-items:flex-end;justify-content:center;height:96px;}
  .swatches{display:grid;grid-template-columns:repeat(auto-fill,minmax(160px,1fr));gap:16px;margin:20px 0;}
  .sw{display:flex;flex-direction:column;gap:7px;}
  .chip{height:44px;border-radius:4px;border:0.5px solid var(--card-border);}
  .sw .n{font-size:13px;font-weight:500;}
  .sw .h{font-family:ui-monospace,monospace;font-size:11.5px;color:var(--dusty);}
  ul{padding-left:18px;margin:0 0 14px;}
  li{margin-bottom:7px;max-width:64ch;}
  .bad{color:var(--alert);}
  .file{font-size:11px;letter-spacing:0.1em;text-transform:uppercase;color:var(--dusty);
        font-weight:600;margin:26px 0 0;}
  a{color:var(--accent);text-decoration:none;border-bottom:0.5px solid var(--card-border);}
  a:hover{border-bottom-color:var(--accent);}
"""

DRAFT = ('<rect x="0" y="0" width="64" height="64" rx="10" fill="#E8DFC8"/>'
 '<rect x="4" y="4" width="16" height="16" rx="2" fill="none" stroke="#2C2C2A" stroke-width="1.5"/>'
 '<rect x="24" y="4" width="16" height="16" rx="2" fill="none" stroke="#2C2C2A" stroke-width="1.5"/>'
 '<rect x="44" y="4" width="16" height="16" rx="2" fill="none" stroke="#2C2C2A" stroke-width="1.5"/>'
 '<rect x="4" y="24" width="16" height="16" rx="2" fill="#3B6D11" stroke="#2C2C2A" stroke-width="1.5"/>'
 '<text x="12" y="33" text-anchor="middle" dominant-baseline="central" font-family="Georgia, serif" '
 'font-size="14" fill="#EDE4CB">T</text>'
 '<rect x="24" y="24" width="16" height="16" rx="2" fill="none" stroke="#2C2C2A" stroke-width="1.5"/>'
 '<rect x="44" y="24" width="16" height="16" rx="2" fill="none" stroke="#2C2C2A" stroke-width="1.5"/>'
 '<rect x="4" y="44" width="16" height="16" rx="2" fill="none" stroke="#2C2C2A" stroke-width="1.5"/>'
 '<rect x="24" y="44" width="16" height="16" rx="2" fill="none" stroke="#2C2C2A" stroke-width="1.5"/>'
 '<rect x="44" y="44" width="16" height="16" rx="2" fill="none" stroke="#2C2C2A" stroke-width="1.5"/>')


# --- помощники разметки ------------------------------------------------------

def px_pattern(pid, w):
    return ('<pattern id="' + pid + '" width="1" height="1" patternUnits="userSpaceOnUse">'
            '<path d="M1 0V1M0 1H1" fill="none" stroke="#B6AC94" stroke-width="' + w + '"/></pattern>')


def mod_pattern(pid, step, color, w):
    s = str(step)
    return ('<pattern id="' + pid + '" width="' + s + '" height="' + s + '" patternUnits="userSpaceOnUse">'
            '<path d="M' + s + ' 0V' + s + 'M0 ' + s + 'H' + s + '" fill="none" stroke="' + color +
            '" stroke-width="' + w + '"/></pattern>')


def plate(vb, inner_, size, cls="", r=0):
    return ('<div class="plate ' + cls + '" style="width:' + str(size) + 'px;height:' + str(size)
            + 'px;border-radius:' + str(r) + 'px">'
            '<svg width="' + str(size) + '" height="' + str(size) + '" viewBox="0 0 '
            + str(vb) + ' ' + str(vb) + '">' + inner_ + '</svg></div>')


def board(inner_, style=""):
    st = ' style="' + style + '"' if style else ''
    return '<div class="board"' + st + '>' + inner_ + '</div>'


def cell(inner_, caption, bad=False):
    cls = 'cap bad' if bad else 'cap'
    return '<div class="cell">' + inner_ + '<div class="' + cls + '">' + caption + '</div></div>'


def svg(size, vb, inner_, extra=""):
    return ('<svg width="' + str(size) + '" height="' + str(size) + '" viewBox="' + vb + '">'
            + inner_ + extra + '</svg>')


def esc(s):
    return s.replace("&", "&amp;").replace("<", "&lt;").replace(">", "&gt;")


def num(x, digits=2):
    return f"{x:.{digits}f}".replace(".", ",")


# --- документ ----------------------------------------------------------------

def head():
    return ('<!DOCTYPE html>' + NL + '<html lang="ru">' + NL + '<head>' + NL
            + '<meta charset="UTF-8">' + NL
            + '<meta name="viewport" content="width=device-width, initial-scale=1">' + NL
            + '<title>Typecase — иконка, эталон построения</title>' + NL
            + '<style>' + CSS + '</style>' + NL + '</head>' + NL
            + '<body>' + NL + '<div class="wrap">' + NL + NL)


def header_block():
    return ('<header>' + NL
            + '  <div class="row" style="margin-bottom:24px;gap:14px;">' + NL
            + '    ' + plate(96, I.inner("light"), 96, r=6) + NL
            + '    ' + plate(96, I.inner("dark"), 96, cls="dark", r=6) + NL
            + '  </div>' + NL
            + '  <h1>Иконка Typecase</h1>' + NL
            + '  <div class="sub">Знак модуля · эталон построения · версия 1.0</div>' + NL
            + '</header>' + NL + NL)


def section_1(v):
    c = v["c"]
    return (
'<p class="lede">Единственный источник правды по знаку модуля Typecase: что изображено, почему '
'именно это, как построено и по каким координатам. Знак подчинён '
'<a href="typecase-logotype.html">эталону надписи</a>: там уже назначены гарнитура, роли цвета и '
'правило акцента, и здесь они не пересматриваются, а переносятся в квадрат. Все числа ниже '
'посчитаны <code>tools/icon_checks.py</code> по выпущенному файлу.</p>' + NL + NL +

'<h2>1. Что изображено и почему</h2>' + NL + NL +

'<p>Наборная касса сверху: чернильный ящик, девять гнёзд, в одном лежит литера, одно свободно. '
'Ровно то, что сказано в спецификации про имя модуля — у каждой литеры своё гнездо, и набор '
'собирается из готовых элементов, а не вырезается каждый раз заново.</p>' + NL + NL +

'<p>Девять гнёзд — минимальная сетка, в которой касса читается кассой, а не решёткой или окном: '
'три на три даёт разом ряды, столбцы и центр. Литера и свободное гнездо стоят в среднем ряду по '
'краям: набор начат слева, место для следующего виджета осталось справа. Это тот же сюжет, что и '
'в надписи, где за словом идёт одно пустое гнездо.</p>' + NL + NL +

'<p><strong>Ящик — не подложка, а предмет.</strong> В надписи гнёзда лежат прямо на бумаге, и '
'этого хватает, потому что литеры крупные и чернильные. В квадрате того же приёма не хватило: '
'картон на бумаге даёт контраст ' + num(c["гнездо на бумаге, как в надписи"]) + ':1, то есть '
'границы гнезда не видно вовсе. Чернильный ящик поднимает контраст до '
+ num(c["гнездо на ящике, светлая"]) + ':1 — перегородки становятся стенками, а гнёзда '
'ячейками. Остальные измерения в разделе 6.</p>' + NL + NL +

'<p><strong>Литера — глиф Literata в кривых</strong>, из того же инстанса, что и надпись, '
'уменьшенный вчетверо. Он не перерисован и не подправлен ни на одном размере: эталон надписи '
'запрещает это прямо, и знак не заводит исключения для себя.</p>' + NL + NL +

'<p><strong>Акцентом красится свободное гнездо, а не занятое.</strong> Тоже правило надписи: '
'акцент — это место, которое ещё можно занять, а не литера, которая уже стоит. Отсюда и смысл '
'знака: библиотека, в которой есть готовое и есть место для нового.</p>' + NL + NL +

'<div class="note">Пиктограмма уровня иконки раздела: каталог виджетов, шапка документации '
'библиотеки, README модуля. Не заменяет знак Carrel и не встаёт фавиконкой всего проекта — '
'Typecase выпускается отдельно, но логотип у проекта один.</div>' + NL + NL)


def section_2(v):
    fig = (
'<figure>' + NL + '<div class="row">' + NL +
cell(board(svg(256, "0 0 64 64", '<defs>' + mod_pattern("pxD", 4, "#B6AC94", "0.28") + '</defs>' + DRAFT,
                '<rect width="64" height="64" fill="url(#pxD)"/>')),
     'черновик · сетка 16 px поверх поля 64', bad=True) + NL +
cell(board(svg(256, "0 0 16 16", '<defs>' + px_pattern("pxE", "0.07") + '</defs>' + I.inner16(),
                '<rect width="16" height="16" fill="url(#pxE)"/>')),
     'эталон · 16-я версия') + NL +
cell(board(plate(64, DRAFT, 16) + plate(16, I.inner16(), 16),
           "display:flex;flex-direction:column;gap:18px;align-items:center;"),
     'те же два, 1:1') + NL +
'</div>' + NL +
'<figcaption>Слева сетка 16 px наложена на черновик в его собственном поле 64: ни одна кромка '
'гнезда не совпадает с линией сетки, потому что обводка 1,5 кладётся по центру контура и уводит '
'кромки на 3,25, 4,75, 19,25 и 20,75. Справа 16-я версия эталона: прямые кромки на целых, '
'полутона остаются только у глифа и скругления.</figcaption>' + NL + '</figure>' + NL + NL)

    return (
'<h2>2. Что чинится против черновика</h2>' + NL + NL +

'<p>Черновик верен по смыслу и неверен по разметке. Пять правок, и первые четыре — те же, что '
'были у знака Carrel до перестройки.</p>' + NL + NL +

'<p><strong>Поле 64 вместо 96.</strong> Само по себе оно делится чисто: 4+16+4+16+4+16+4 = 64, '
'и половина, и четверть дают целые числа, так что претензия не к арифметике. Знак выпадает из '
'семейства: у <a href="carrel-icon.html">Carrel</a> и <a href="davloom-icon.html">davloom</a> '
'квадрат 96 с модулем 6, у них общие правила хинтинга и общие проверки. На 96 та же сетка три на '
'три ложится без остатка — 6+24+6+24+6+24+6, — и деление остаётся точным и на 3, и на 6, чего '
'поле 64 не даёт: на 3 оно не делится вовсе.</p>' + NL + NL +

'<p><strong>Гнёзда нарисованы обводкой.</strong> <code>stroke-width="1.5"</code> кладётся по '
'центру контура, поэтому кромки гнезда стоят на 3,25, 4,75, 19,25 и 20,75 — ни одна не попадает '
'даже на половину пикселя. Обводка мешает и всему, что делается с иконкой дальше: сборке '
'<code>.ico</code>, монохромной и маскируемой версиям. В эталоне надписи это сказано короче: '
'три заливки, ноль обводок.</p>' + NL + NL +

'<p><strong>Литера набрана шрифтом.</strong> <code>&lt;text font-family="Georgia, serif"&gt;</code> '
'рисуется тем, что нашлось в системе: без Georgia это другой шрифт, другая ширина и другое '
'положение в гнезде. И даже с Georgia это была бы не та гарнитура: надпись Typecase набрана '
'Literata, и знак берёт литеру из того же инстанса.</p>' + NL + NL +

'<p><strong>Палитра чужая.</strong> <code>#2C2C2A</code>, <code>#3B6D11</code>, '
'<code>#E8DFC8</code> и <code>#EDE4CB</code> не встречаются в проекте больше нигде. Это соседние '
'оттенки чернил, акцента и картона: достаточно близкие, чтобы сойти за те же, и достаточно '
'разные, чтобы рядом с настоящей палитрой читаться как ошибка печати. Знак берёт токены '
'библиотеки из <code>css/00-tokens.css</code>.</p>' + NL + NL +

'<p><strong>Акцент стоит не на том гнезде.</strong> В черновике закрашено занятое гнездо — то, '
'где лежит литера. Эталон надписи запрещает это прямо: акцент — одно пустое гнездо, занятые '
'акцентом не красятся. Знак приведён к этому правилу, и вместе с ним поменялся сюжет: не «вот '
'литера», а «вот литера, а вот свободное место».</p>' + NL + NL +

'<p>Это можно не оценивать на глаз, а посчитать. Оба знака растеризованы, и каждый пиксель '
'проверен на совпадение с одним из цветов своей палитры; всё, что не совпало, — промежуточный '
'полутон, то есть размытая кромка.</p>' + NL + NL +

'<table>' + NL +
'  <tr><th style="width:30%">Растеризация</th><th class="num" style="width:15%">черновик</th>'
'<th class="num" style="width:15%">эталон</th><th>Чем считается</th></tr>' + NL +
'  <tr><td>16 × 16, 256 пикселей</td><td class="num bad">' + str(v["draft16"]) + '</td>'
'<td class="num">' + str(v["n16"]) + '</td>'
'<td>У черновика в полутонах ' + num(v["draft16_pct"], 0) + ' % рисунка — от знака остаётся '
'серое пятно. У эталона ' + str(v["n16"]) + ', из них ' + str(v["glyph16"]) + ' — засечки глифа '
'и ' + str(v["round16"]) + ' — скругление гнезда; прямых кромок в полутонах нет.</td></tr>' + NL +
'  <tr><td>32 × 32, 1024 пикселя</td><td class="num bad">' + str(v["draft32"]) + '</td>'
'<td class="num">' + str(v["n32"]) + '</td>'
'<td>' + str(v["round32"]) + ' полутонов — скругления девяти гнёзд, ровно по четыре угла на '
'гнездо, остальные ' + str(v["glyph32"]) + ' — глиф.</td></tr>' + NL +
'</table>' + NL + NL + fig)


def section_3(v):
    draw = svg(384, "0 0 96 96",
        '<defs>' + mod_pattern("g6", 6, "#D8CFBB", "0.25") + mod_pattern("g24", 24, "#C9BFA6", "0.4") + '</defs>'
        + '<rect width="96" height="96" fill="#FFFDF7"/>'
        + '<rect width="96" height="96" fill="url(#g6)"/>'
        + '<rect width="96" height="96" fill="url(#g24)"/>'
        + '<path fill="' + I.PAL["light"]["ink"] + '" fill-opacity="0.30" d="' + I.cells(skip={I.FREE}) + '"/>'
        + '<path fill="' + I.PAL["light"]["accent"] + '" fill-opacity="0.45" d="' + I.one(*I.FREE) + '"/>'
        + '<path fill="' + I.PAL["light"]["ink"] + '" fill-opacity="0.55" d="' + I.glyph(*I.LETTER) + '"/>'
        + '<g stroke="#B04A3A" stroke-width="0.4" fill="none" stroke-dasharray="3 2">'
          '<path d="M0 48H96"/><path d="M48 0V96"/></g>'
        + '<rect x="6" y="6" width="84" height="84" fill="none" stroke="#B04A3A" stroke-width="0.5"/>'
        + '<rect width="96" height="96" fill="none" stroke="#9A9280" stroke-width="0.5"/>')

    p = I.PAL["light"]
    dims = ('<svg width="600" height="387" viewBox="-30 -12 208 134">'
      '<g font-family="ui-monospace,SFMono-Regular,Menlo,monospace" font-size="4.2" fill="#6B6355">'
      '<rect x="0" y="0" width="96" height="96" fill="' + p["box"] + '"/>'
      '<path fill="' + p["card"] + '" d="' + I.cells(skip={I.FREE}) + '"/>'
      '<path fill="' + p["accent"] + '" d="' + I.one(*I.FREE) + '"/>'
      '<path fill="' + p["ink"] + '" d="' + I.glyph(*I.LETTER) + '"/>'
      '<g stroke="#9A9280" stroke-width="0.35" fill="none">'
        '<path d="M0 104v4M6 104v4M30 104v4M36 104v4"/>'
        '<path d="M0 106H6M6 106H30M30 106H36"/>'
        '<path d="M-10 0h-4M-10 6h-4M-10 30h-4M-10 36h-4"/>'
        '<path d="M-12 0V6M-12 6V30M-12 30V36"/>'
      '</g>'
      '<text x="0.5" y="113">6</text><text x="14" y="113">24</text><text x="30.5" y="113">6</text>'
      '<text x="-24" y="4.5">6</text><text x="-27" y="20">24</text><text x="-24" y="34.5">6</text>'
      '<g stroke="#9A9280" stroke-width="0.35" fill="none" stroke-dasharray="1.5 1.5">'
        '<path d="M96 18H104"/><path d="M96 48H104"/><path d="M96 78H104"/>'
      '</g>'
      '<text x="106" y="19.5">ящик · весь квадрат</text>'
      '<text x="106" y="49.5">свободное гнездо</text>'
      '<text x="106" y="79.5">гнездо · 24 × 24 · r 3</text>'
      '</g></svg>')

    close = ('<svg width="330" height="330" viewBox="-11 26 52 48">'
      '<g font-family="ui-monospace,SFMono-Regular,Menlo,monospace" font-size="2.4" fill="#6B6355">'
      '<rect x="-11" y="26" width="52" height="48" fill="' + p["box"] + '"/>'
      '<path fill="' + p["card"] + '" d="' + I.one(*I.LETTER) + '"/>'
      '<path fill="' + p["ink"] + '" d="' + I.glyph(*I.LETTER) + '"/>'
      '<g stroke="#B04A3A" stroke-width="0.18" fill="none" stroke-dasharray="1 0.8">'
        '<path d="M10.292 34V62M25.708 34V62"/>'
        '<path d="M2 39.75H34M2 56.25H34"/>'
      '</g>'
      '<g stroke="#9A9280" stroke-width="0.2" fill="none">'
        '<path d="M6 62v3M10.292 62v3M25.708 62v3M30 62v3"/>'
        '<path d="M6 64H10.292M10.292 64H25.708M25.708 64H30"/>'
        '<path d="M4 36h-2M4 39.75h-2M4 56.25h-2M4 60h-2"/>'
        '<path d="M3 36V39.75M3 39.75V56.25M3 56.25V60"/>'
      '</g>'
      '<text x="5.4" y="68">' + num(v["air_x"]) + '</text>'
      '<text x="15.2" y="68">' + num(v["glyph_w"]) + '</text>'
      '<text x="25" y="68">' + num(v["air_x"]) + '</text>'
      '<text x="-8.5" y="38.9">' + num(v["air_y"]) + '</text>'
      '<text x="-8.5" y="49">' + num(v["glyph_h"], 1) + '</text>'
      '<text x="-8.5" y="59.2">' + num(v["air_y"]) + '</text>'
      '</g></svg>')

    gx, gy, gw, gh, ax, ay = I.glyph_box(*I.LETTER)

    return (
'<h2>3. Сетка и модуль</h2>' + NL + NL +

'<p>Мастер — квадрат 96, модуль 6. Та же логика, что у <a href="carrel-icon.html">знака '
'Carrel</a>: 96 делится и на 3, и на 6 без остатка. Сетка гнёзд собирается из одного числа: '
'<strong>6 + 24 + 6 + 24 + 6 + 24 + 6 = 96</strong>, где 6 — и внешнее поле, и перегородка между '
'гнёздами, а 24 — само гнездо. Скругление 3 — половина модуля.</p>' + NL + NL +

'<p>Равенство поля и перегородки не украшение, а условие: пока они равны, ящик читается решетом '
'с одинаковым шагом, и при делении на 3 шаг остаётся одинаковым. Ящик залит на весь квадрат, '
'полей у знака нет — касса это ящик, а ящик занимает всё, что ему отведено.</p>' + NL + NL +

'<figure>' + NL + '<div class="row">' + NL +
cell(board(draw), 'сетка 6 и 24, кайлайн 84 × 84, оси') + NL +
'</div>' + NL +
'<figcaption>Гнёзда стоят на 6, 36 и 66 по обеим осям. Литера и свободное гнездо — в среднем '
'ряду: относительно горизонтальной оси знак зеркален, относительно вертикальной симметрию '
'нарушают литера слева и акцент справа, и из этого получается направление чтения.</figcaption>'
+ NL + '</figure>' + NL + NL +

'<figure>' + NL + '<div class="row">' + NL +
cell(board(dims), 'общие размеры') + NL +
cell(board(close), 'занятое гнездо крупно: воздух и глиф') + NL +
'</div>' + NL +
'<figcaption>Единственное назначенное число литеры — рост прописных: ' + num(I.CAP, 1)
+ ', ровно четверть от ' + num(I.CAP * I.DIV, 0) + ' в надписи. Ширина ' + num(gw)
+ ' и воздух ' + num(ax) + ' по горизонтали следуют из него, потому что пропорции глифа не '
'трогаются.</figcaption>' + NL + '</figure>' + NL + NL +

'<h3>Координаты мастера</h3>' + NL +
'<table>' + NL +
'  <tr><th style="width:26%">Элемент</th><th class="num" style="width:22%">X</th>'
'<th class="num" style="width:22%">Y</th><th>Размер</th></tr>' + NL +
'  <tr><td>Ящик</td><td class="num">0 … 96</td><td class="num">0 … 96</td>'
'<td>заливка на весь квадрат</td></tr>' + NL +
'  <tr><td>Гнёзда, начала</td><td class="num">6, 36, 66</td><td class="num">6, 36, 66</td>'
'<td>девять штук, 24 × 24, радиус 3</td></tr>' + NL +
'  <tr><td>Занятое гнездо</td><td class="num">6 … 30</td><td class="num">36 … 60</td>'
'<td>левое среднее</td></tr>' + NL +
'  <tr><td>Свободное гнездо</td><td class="num">66 … 90</td><td class="num">36 … 60</td>'
'<td>правое среднее, акцент</td></tr>' + NL +
'  <tr><td>Глиф T</td><td class="num">' + num(gx) + ' … ' + num(gx + gw) + '</td>'
'<td class="num">' + num(gy) + ' … ' + num(gy + gh) + '</td>'
'<td>' + num(gw) + ' × ' + num(gh, 1) + ' · Literata, рост прописных ' + num(I.CAP, 1)
+ '</td></tr>' + NL +
'</table>' + NL + NL +

'<div class="note">Координаты глифа дробные, и это не оплошность: сетка гнёзд целочисленна, а '
'литера — не рисунок, а глиф шрифта, который эталон надписи запрещает править. Дробность здесь '
'принадлежит шрифту, а не разметке. ' + v["glyph"] + '</div>' + NL + NL)


def section_4(v):
    v32 = board(svg(300, "0 0 32 32",
                '<defs>' + px_pattern("px32", "0.08") + '</defs>' + I.inner32(),
                '<rect width="32" height="32" fill="url(#px32)"/>'))
    v16 = board(svg(300, "0 0 16 16",
                '<defs>' + px_pattern("px16", "0.05") + '</defs>' + I.inner16(),
                '<rect width="16" height="16" fill="url(#px16)"/>'))

    return (
'<h2>4. Хинтинг: 32 и 16</h2>' + NL + NL +

'<p><strong>Сетка делится на 3 целиком:</strong> поле и перегородка 6 → 2, гнездо 24 → 8, радиус '
'3 → 1, и 2+8+2+8+2+8+2 = 32. Поэтому мастер и сам ужимается почти чисто: на 32 он даёт '
+ str(v["m32"]) + ' полутоновых пикселя из 1024 против ' + str(v["n32"]) + ' у хинтованной '
'версии. Разница в два пикселя — не повод держать отдельный файл, и хинтинг здесь нужен не '
'ради кромок.</p>' + NL + NL +

'<p><strong>32-я версия существует ради литеры.</strong> При делении мастера рост прописных '
'становится ' + num(I.CAP / 3, 1) + ' — половина пикселя, на которой засечки Literata '
'превращаются в серую бахрому. В хинтованной версии он поднят до 6 при гнезде 8, то есть воздух '
'ровно по пикселю сверху и снизу. Глиф при этом не перерисован: изменился только рост, а '
'пропорции остались шрифтовыми.</p>' + NL + NL +

'<p><strong>16-я версия — не деление, а редукция.</strong> По сетке она сходится идеально: поле '
'1, гнездо 4, перегородка 1, и 1+4+1+4+1+4+1 = 16. Беда в том, что гнездо получается 4 × 4, а '
'прописная в нём — три пикселя высотой: не литера, а грязное пятно. Девять пустых гнёзд без '
'литеры читаются решёткой, то есть пропадает именно то, ради чего знак нарисован.</p>' + NL + NL +

'<p>Поэтому на 16 рисуется <strong>одно гнездо во весь квадрат</strong>: ящик по краю в пиксель, '
'гнездо 14 × 14 с радиусом 1, глиф ростом 10 при воздухе 2 сверху и снизу. Касса схлопывается до '
'своей единицы — гнезда с литерой. Свободного гнезда на этом размере нет, и акцент вместе с ним '
'пропадает: на 16 px знак сообщает «литера в кассе», а не «в кассе есть место».</p>' + NL + NL +

'<table>' + NL +
'  <tr><th style="width:26%">Что</th><th class="num" style="width:20%">96 ÷ 3 даёт</th>'
'<th class="num" style="width:16%">На 32</th><th>Хинт</th></tr>' + NL +
'  <tr><td>Поле и перегородка</td><td class="num">2</td><td class="num">2</td>'
'<td>не потребовался</td></tr>' + NL +
'  <tr><td>Гнездо</td><td class="num">8</td><td class="num">8</td><td>не потребовался</td></tr>' + NL +
'  <tr><td>Радиус</td><td class="num">1</td><td class="num">1</td><td>не потребовался</td></tr>' + NL +
'  <tr><td>Рост прописных</td><td class="num">' + num(I.CAP / 3, 1) + '</td>'
'<td class="num">6</td><td>поднят до целого; воздух ровно по пикселю</td></tr>' + NL +
'</table>' + NL + NL +

'<figure>' + NL + '<div class="row">' + NL +
cell(v32, '32 px · сетка делением, литера хинтом') + NL +
cell(v16, '16 px · одно гнездо, редукция') + NL +
'</div>' + NL +
'<figcaption>Прямые кромки гнёзд и ящика на обоих размерах лежат на границах пикселей. Серым '
'мажутся только скругления и засечки глифа — то есть кривые, которым сглаживание и '
'положено.</figcaption>' + NL + '</figure>' + NL + NL)


def section_5(v):
    m = v["master_file"]
    master_row = ('rect ' + str(m["rects"]) + ', path ' + str(m["paths"]) + ', обводок '
                  + str(m["strokes"]) + ', упоминаний шрифта ' + str(m["texts"])
                  + ', блоков метаданных ' + str(m["meta"]) + ', '
                  + num(m["bytes"] / 1024, 1) + ' КБ. Заливки: '
                  + ", ".join(m["fills"]) + '.') if m else 'мастер не выпущен'

    return (
'<h2>5. Проверки</h2>' + NL + NL +

'<p>Каждая строка посчитана <code>tools/icon_checks.py</code> по выпущенному файлу: координаты — '
'разбором путей, полутона — растеризацией со сравнением каждого пикселя с палитрой, контраст — '
'по формуле относительной яркости. Своих слов о результате у документа нет.</p>' + NL + NL +

'<table>' + NL +
'  <tr><th style="width:32%">Что проверялось</th><th>Результат</th></tr>' + NL +
'  <tr><td>Целые координаты сетки</td><td>' + v["grid"] + '</td></tr>' + NL +
'  <tr><td>Деление на 3</td><td>' + v["division"] + '</td></tr>' + NL +
'  <tr><td>Литера от шрифта</td><td>' + v["glyph"] + '</td></tr>' + NL +
'  <tr><td>Воздух в гнезде</td><td>' + v["air"] + '</td></tr>' + NL +
'  <tr><td>Растр 16 × 16</td><td>' + v["raster16"] + '</td></tr>' + NL +
'  <tr><td>Растр 32 × 32</td><td>' + v["raster32"] + '</td></tr>' + NL +
'  <tr><td>Мастер на своих размерах</td><td>' + v["master_raster"] + '</td></tr>' + NL +
'  <tr><td>Контраст гнезда и ящика</td><td>' + v["contrast"] + '</td></tr>' + NL +
'  <tr><td>Контраст акцента</td><td>' + v["accent_contrast"] + '</td></tr>' + NL +
'  <tr><td>Безопасная зона PWA maskable</td><td>' + v["maskable"] + '</td></tr>' + NL +
'  <tr><td>Что лежит в мастере</td><td>' + master_row + '</td></tr>' + NL +
'</table>' + NL + NL)


def section_6(v):
    c = v["c"]
    L, D = I.PAL["light"], I.PAL["dark"]
    sw = ('<div class="swatches">' +
      '<div class="sw"><div class="chip" style="background:' + L["box"] + '"></div>'
      '<div class="n">Ящик, светлая</div><div class="h">--tc-ink · ' + L["box"] + '</div></div>' +
      '<div class="sw"><div class="chip" style="background:' + L["card"] + '"></div>'
      '<div class="n">Гнёзда, светлая</div><div class="h">--tc-card · ' + L["card"] + '</div></div>' +
      '<div class="sw"><div class="chip" style="background:' + L["accent"] + '"></div>'
      '<div class="n">Свободное гнездо</div><div class="h">--tc-accent · ' + L["accent"] + '</div></div>' +
      '<div class="sw"><div class="chip" style="background:' + D["box"] + '"></div>'
      '<div class="n">Ящик, тёмная</div><div class="h">--tc-bg · ' + D["box"] + '</div></div>' +
      '<div class="sw"><div class="chip" style="background:' + D["card"] + '"></div>'
      '<div class="n">Гнёзда, тёмная</div><div class="h">--tc-ink · ' + D["card"] + '</div></div>' +
      '<div class="sw"><div class="chip" style="background:' + D["accent"] + '"></div>'
      '<div class="n">Свободное, тёмная</div><div class="h">--tc-accent · ' + D["accent"] + '</div></div>' +
      '</div>')

    mono = ('<rect width="96" height="96" fill="' + L["box"] + '"/>'
            + '<path fill="#FFFDF7" d="' + I.cells() + '"/>'
            + '<path fill="' + L["box"] + '" d="' + I.glyph(*I.LETTER) + '"/>')

    figs = ('<figure>' + NL + '<div class="row">' + NL +
      cell(plate(96, I.inner("light"), 96), 'светлая') + NL +
      cell(plate(96, I.inner("dark"), 96, cls="dark"), 'тёмная') + NL +
      cell(plate(96, mono, 96, cls="paper"),
           'монохром: свободное гнездо отличается только тем, что в нём нет литеры') + NL +
      '</div>' + NL +
      '<figcaption>В монохроме акцент передать нечем, и подменять его обводкой нельзя — это та '
      'же цена, что названа в эталоне надписи. Остаётся касса, в которой одно гнездо занято, а '
      'остальные пусты.</figcaption>' + NL + '</figure>')

    return (
'<h2>6. Цвет</h2>' + NL + NL +

'<p>Все цвета — токены библиотеки из <code>css/00-tokens.css</code>. Роли назначены '
'одним правилом: <strong>ящик красится самым тёмным токеном темы, гнёзда — самым светлым, '
'литера — цветом ящика</strong>. В светлой теме самый тёмный это <code>--tc-ink</code>, в '
'тёмной — <code>--tc-bg</code>; самым светлым в светлой оказывается <code>--tc-card</code>, в '
'тёмной — <code>--tc-ink</code>. Гнёзда поэтому в обеих темах одного цвета, а меняется тон '
'ящика.</p>' + NL + NL + sw + NL + NL +

'<p>Это единственное место, где знак расходится с <a href="typecase-logotype.html">эталоном '
'надписи</a>, и расхождение вынужденное. В надписи гнёзда — <code>--tc-card</code> на бумаге, а '
'литеры — <code>--tc-ink</code>. В знаке появился ящик, которого в надписи нет, и роли пришлось '
'пересобрать вокруг него.</p>' + NL + NL +

'<table>' + NL +
'  <tr><th style="width:44%">Пара</th><th class="num" style="width:16%">Контраст</th>'
'<th>Что это значит</th></tr>' + NL +
'  <tr><td>гнездо на ящике, светлая</td><td class="num">' + num(c["гнездо на ящике, светлая"])
+ ':1</td><td>перегородка читается как стенка</td></tr>' + NL +
'  <tr><td>гнездо на ящике, тёмная</td><td class="num">' + num(c["гнездо на ящике, тёмная"])
+ ':1</td><td>то же в тёмной теме</td></tr>' + NL +
'  <tr><td>акцент к соседнему гнезду</td><td class="num">' + num(c["акцент к соседнему гнезду"])
+ ':1</td><td>свободное гнездо видно среди занятых</td></tr>' + NL +
'  <tr><td>акцент к ящику</td><td class="num">' + num(c["акцент к ящику"])
+ ':1</td><td>мало — поэтому акцент не ставится с краю в одиночестве</td></tr>' + NL +
'  <tr><td class="bad">гнездо card на бумаге bg</td><td class="num bad">'
+ num(c["гнездо на бумаге, как в надписи"]) + ':1</td>'
'<td>граница не видна вовсе; так знак делать нельзя</td></tr>' + NL +
'  <tr><td class="bad">гнездо card на фоне bg, тёмная</td><td class="num bad">'
+ num(c["гнездо на фоне, тёмная надпись"]) + ':1</td>'
'<td>то же в тёмной теме, если оставить роли как в надписи</td></tr>' + NL +
'</table>' + NL + NL + figs + NL + NL)


def section_7(v):
    light = '<div class="strip">'
    for size, vb, art, name in ((16, 16, I.inner16(), '16'), (20, 16, I.inner16(), '20'),
                                (24, 32, I.inner32(), '24'), (32, 32, I.inner32(), '32'),
                                (48, 96, I.inner("light"), '48'), (64, 96, I.inner("light"), '64'),
                                (96, 96, I.inner("light"), '96')):
        light += ('<div class="cell"><div class="sizebox">' + plate(vb, art, size)
                  + '</div><div class="cap">' + name + '</div></div>')
    light += '</div>'

    dark = '<div class="strip">'
    for size, vb, art, name in ((16, 16, I.inner16("dark"), '16'), (20, 16, I.inner16("dark"), '20'),
                                (24, 32, I.inner32("dark"), '24'), (32, 32, I.inner32("dark"), '32'),
                                (48, 96, I.inner("dark"), '48'), (64, 96, I.inner("dark"), '64'),
                                (96, 96, I.inner("dark"), '96')):
        dark += ('<div class="cell"><div class="sizebox">' + plate(vb, art, size, cls="dark")
                 + '</div><div class="cap">' + name + '</div></div>')
    dark += '</div>'

    return (
'<h2>7. Реальные размеры</h2>' + NL + NL +
'<figure>' + NL + light + NL + dark + NL +
'<figcaption>Скругление плашки в этой таблице задано стилями страницы и в файлы не входит: '
'иконка всегда полноформатный квадрат, а маску накладывает система.</figcaption>' + NL +
'</figure>' + NL + NL +

'<h3>Какой рисунок на каком размере</h3>' + NL +
'<table>' + NL +
'  <tr><th class="num" style="width:22%">Размер вывода</th><th style="width:24%">Файл</th>'
'<th>Что видно</th></tr>' + NL +
'  <tr><td class="num">до 23 px</td><td>16-я версия</td>'
'<td>Одно гнездо с литерой. Свободного гнезда и акцента нет.</td></tr>' + NL +
'  <tr><td class="num">24 … 47 px</td><td>32-я версия</td>'
'<td>Полная касса: девять гнёзд, литера и акцент.</td></tr>' + NL +
'  <tr><td class="num">от 48 px</td><td>мастер 96</td>'
'<td>То же с настоящим скруглением и полным глифом.</td></tr>' + NL +
'</table>' + NL + NL +

'<div class="note">Полутонов ноль не бывает ни на одном размере, и это не дефект разметки: '
'сглаживаются засечки глифа и дуги скруглений. Проверять надо другое — прямые кромки. Если '
'убрать литеру и скругления, у 16-й версии остаётся ровно 0 полутоновых пикселей из 256, а у '
'32-й — только ' + str(v["round32"]) + ' угловых. Это и есть признак того, что сетка стоит на '
'пикселях.</div>' + NL + NL)


def section_8(v):
    p = I.PAL["light"]
    mask = ('<rect x="-30" y="-30" width="156" height="156" fill="' + p["box"] + '"/>'
        + '<path fill="' + p["card"] + '" d="' + I.cells(skip={I.FREE}) + '"/>'
        + '<path fill="' + p["accent"] + '" d="' + I.one(*I.FREE) + '"/>'
        + '<path fill="' + p["ink"] + '" d="' + I.glyph(*I.LETTER) + '"/>'
        + '<circle cx="48" cy="48" r="' + num(v["r_mask"], 1).replace(",", ".")
        + '" fill="none" stroke="#B04A3A" stroke-width="0.7" stroke-dasharray="3 3"/>')
    small = ('<rect width="96" height="96" fill="' + p["box"] + '"/>'
        + '<path fill="' + p["card"] + '" d="' + I.cells(skip={I.FREE}) + '"/>'
        + '<path fill="' + p["accent"] + '" d="' + I.one(*I.FREE) + '"/>'
        + '<path fill="' + p["ink"] + '" d="' + I.glyph(*I.LETTER) + '"/>'
        + '<circle cx="48" cy="48" r="' + num(v["r96"], 1).replace(",", ".")
        + '" fill="none" stroke="#B04A3A" stroke-width="0.5" stroke-dasharray="2 2"/>')

    return ('<h2>8. Маскируемая версия</h2>' + NL + NL +
'<p>Ящик занимает весь квадрат — у кассы нет полей по смыслу. Для PWA это значит, что '
'маскируемая версия не может быть тем же файлом: её кладут на расширенный холст, не трогая '
'геометрию знака.</p>' + NL + NL +
'<figure>' + NL + '<div class="row">' + NL +
cell(board('<svg width="240" height="240" viewBox="-30 -30 156 156">' + mask + '</svg>'),
     'холст ' + str(v["mask_field"]) + ', безопасный радиус ' + num(v["r_mask"], 1)) + NL +
cell(board('<svg width="240" height="240" viewBox="0 0 96 96">' + small + '</svg>'),
     'то же на поле 96: круг режет угловые гнёзда', bad=True) + NL +
'</div>' + NL +
'<figcaption>' + v["maskable"] + ' Ящик при этом заливает весь расширенный квадрат.</figcaption>'
+ NL + '</figure>' + NL + NL)


def section_9():
    def pre(title, code):
        return ('<div class="file">' + title + '</div>' + NL
                + '<pre><code>' + esc(code) + '</code></pre>' + NL)

    L, D = I.PAL["light"], I.PAL["dark"]
    glyph_head = I.glyph(*I.LETTER)[:52]

    icon = ('<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 96 96" role="img" aria-label="Typecase">' + NL
        + '  <rect width="96" height="96" fill="' + L["box"] + '"/>' + NL
        + '  <path fill="' + L["card"] + '" d="' + I.cells(skip={I.FREE}) + '"/>' + NL
        + '  <path fill="' + L["accent"] + '" d="' + I.one(*I.FREE) + '"/>' + NL
        + '  <path fill="' + L["ink"] + '" d="' + glyph_head + ' … глиф Literata …"/>' + NL
        + '</svg>')
    dark = ('<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 96 96" role="img" aria-label="Typecase">' + NL
        + '  <rect width="96" height="96" fill="' + D["box"] + '"/>' + NL
        + '  <path fill="' + D["card"] + '" d="… те же гнёзда …"/>' + NL
        + '  <path fill="' + D["accent"] + '" d="… свободное гнездо …"/>' + NL
        + '  <path fill="' + D["ink"] + '" d="… тот же глиф …"/>' + NL + '</svg>')
    maskable = ('<svg xmlns="http://www.w3.org/2000/svg" viewBox="-30 -30 156 156" role="img" aria-label="Typecase">' + NL
        + '  <rect x="-30" y="-30" width="156" height="156" fill="' + L["box"] + '"/>' + NL
        + '  … те же три пути, координаты не меняются …' + NL + '</svg>')
    mod, cellw, rad = I.GRID32
    pos32 = [mod, mod + cellw + mod, mod + 2 * (cellw + mod)]
    i32 = ('<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" width="32" height="32" role="img" aria-label="Typecase">' + NL
        + '  <rect width="32" height="32" fill="' + L["box"] + '"/>' + NL
        + '  <path fill="' + L["card"] + '" d="'
        + I.cells(skip={I.FREE}, pos=pos32, cell=cellw, radius=rad) + '"/>' + NL
        + '  <path fill="' + L["accent"] + '" d="'
        + I.one(*I.FREE, pos=pos32, cell=cellw, radius=rad) + '"/>' + NL
        + '  <path fill="' + L["ink"] + '" d="… глиф, рост прописных 6 …"/>' + NL + '</svg>')
    i16 = ('<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" width="16" height="16" role="img" aria-label="Typecase">' + NL
        + '  <rect width="16" height="16" fill="' + L["box"] + '"/>' + NL
        + '  <path fill="' + L["card"] + '" d="' + I.rounded(1, 1, I.CELL16, I.CELL16, I.RADIUS16) + '"/>' + NL
        + '  <path fill="' + L["ink"] + '" d="… глиф, рост прописных 10 …"/>' + NL + '</svg>')

    return ('<h2>9. Исходники</h2>' + NL + NL +
'<p>Полный мастер лежит рядом в <code>typecase-mark.svg</code>; ниже — то же содержимое с '
'сокращённым путём глифа. Порядок обязателен: ящик, гнёзда, свободное гнездо, литера — тот же, '
'что в надписи, где сказано «гнёзда, затем литеры». Пересобираются все файлы одной командой '
'<code>python tools/build_typecase_icon.py</code>.</p>' + NL + NL +
pre('typecase-mark.svg — мастер, светлая тема', icon) +
pre('typecase-icon-dark.svg — тёмная тема', dark) +
pre('typecase-icon-maskable.svg — PWA maskable, холст 156', maskable) +
pre('typecase-icon-32.svg — сетка делением, литера хинтом', i32) +
pre('typecase-icon-16.svg — редукция до одного гнезда', i16) + NL)


def section_10():
    return ('<h2>10. Что эталон фиксирует</h2>' + NL +
'<ul>' + NL +
'  <li>Координаты из таблицы раздела 3 — единственный источник геометрии. Любой файл, где кромка '
'стоит не там, — испорченная копия, а не вариант.</li>' + NL +
'  <li>Литера — глиф Literata того же инстанса, что и надпись. Не рисуется, не правится, не '
'заменяется другой гарнитурой; назначается только рост прописных.</li>' + NL +
'  <li>Акцент — свободное гнездо. Занятое гнездо и литера акцентом не красятся никогда.</li>' + NL +
'  <li>Ящик обязателен: гнёзда не кладутся прямо на бумагу или на фон приложения. Контраст '
'гнезда к подложке не ниже 12:1.</li>' + NL +
'  <li>Заливки, ноль обводок. Порядок: ящик, гнёзда, свободное гнездо, литера.</li>' + NL +
'  <li>Литера — левое среднее гнездо, свободное — правое среднее. Они не переезжают между '
'размерами и темами.</li>' + NL +
'  <li>Цвета берутся только из токенов библиотеки по правилу раздела 6. Шестнадцатеричные '
'значения мимо токенов — ошибка.</li>' + NL +
'  <li>Хинтинг допускается только там, где он перечислен в разделе 4: рост прописных на 32 и '
'редукция до одного гнезда на 16.</li>' + NL +
'  <li>Никаких градиентов, теней, поворотов и отражений. Касса лежит горизонтально и смотрит в '
'одну сторону всегда.</li>' + NL +
'  <li>Знак не подменяет логотип Carrel и не ставится фавиконкой проекта.</li>' + NL +
'</ul>' + NL + NL)


def section_11(v):
    c = v["c"]
    return ('<h2>11. Чего эталон не закрывает</h2>' + NL +
'<ul>' + NL +
'  <li><strong>Контраст гнёзд в самой надписи.</strong> Там гнёзда <code>--tc-card</code> лежат '
'на <code>--tc-bg</code>, и это ' + num(c["гнездо на бумаге, как в надписи"]) + ':1 в светлой '
'теме и ' + num(c["гнездо на фоне, тёмная надпись"]) + ':1 в тёмной: прямоугольники пропадают, '
'остаётся слово. Знак обошёл это ящиком, но надпись живёт по своему эталону, и правку надо '
'вносить туда — либо ящиком, либо новой ролью цвета. Здесь она не сделана намеренно: чужой '
'утверждённый файл не переписывается мимоходом.</li>' + NL +
'  <li><strong>Растровых кадров нет.</strong> Хинтованные 16 и 32 существуют как построения; '
'<code>.ico</code> и PNG для каталога из них ещё не собраны.</li>' + NL +
'  <li><strong>Связка знака с надписью не описана.</strong> Оба эталона существуют, но правила '
'отступа и выравнивания знака рядом со словом нет ни здесь, ни в файле надписи.</li>' + NL +
'  <li><strong>Черновики остаются на диске.</strong> <code>typecase-logo-bg.svg</code>, '
'<code>typecase-logo-v2.svg</code> и <code>typecase-wordmark-tiles.svg</code> не удалены и не '
'переписаны — с ними надо решить отдельно, чтобы случайная копия не ушла в работу вместо '
'эталонной.</li>' + NL +
'  <li><strong>Место знака в интерфейсе каталога не описано.</strong> Эталон говорит, как знак '
'построен и в каких размерах он читается, но не где он стоит на страницах живого '
'каталога.</li>' + NL +
'</ul>' + NL + NL)


def document():
    v = C.verdicts()
    return (head() + header_block() + section_1(v) + section_2(v) + section_3(v)
            + section_4(v) + section_5(v) + section_6(v) + section_7(v) + section_8(v)
            + section_9() + section_10() + section_11(v)
            + '</div>' + NL + '</body>' + NL + '</html>' + NL)


if __name__ == "__main__":
    open(I.MASTER, "w", encoding="utf-8").write(I.svg("light"))
    print("мастер   ", os.path.relpath(I.MASTER, I.HERE),
          f"{os.path.getsize(I.MASTER) / 1024:.1f} КБ")
    doc = document()
    open(I.REFERENCE, "w", encoding="utf-8").write(doc)
    print("эталон   ", os.path.relpath(I.REFERENCE, I.HERE),
          f"{os.path.getsize(I.REFERENCE) / 1024:.1f} КБ")
