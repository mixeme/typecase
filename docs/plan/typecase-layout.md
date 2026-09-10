# Раскладка Typecase — один объект, как на столе

Инструкция исполнителю. Развилок нет: все решения в этом файле закрыты. Чекбоксы отмечают по мере коммитов.

Числа из блока `.m-*` кладутся в таблицы пакета `layout`, не в `css/*.css` руками.

**Цель.** Разработчик экрана только компонует виджеты библиотеки и никогда не изобретает местный обход. Это держится на трёх вещах, и все три должны быть в коде, а не в договорённости:

1. **Забор.** Экран не может править виджет. Уже стоит: ворота `screens-css` и `no-inline`. В этом плане — закрыть дырку с `data-tc-`.
2. **Полнота словаря.** На каждую коробку макета есть виджет. Проверяется перечислением `.m-*`, а не рукописным списком, и подтверждается сборкой оставшихся разделов.
3. **Запасной выход.** Когда нужной формы нет, экран собирает её у себя из примитивов и контейнера — не останавливается и не идёт в CSS. Сегодня выхода нет: 24 виджета каталога из 76 встречаются не больше чем на одном экране, потому что композицию класть больше некуда.

Пакет `layout` — предпосылка всех трёх: пока геометрия размазана по четырём файлам, добавить недостающую форму дорого и рискованно, а дорогая правка библиотеки и есть причина местных обходов.

Не Qt целиком: серверный HTML остаётся. Сходство — одно описание прямоугольников и композиция вместо перерисовки.

- [x] Шаг 0 — спека L-3.1.4, L-3.6, L-4, L-8, CHANGELOG
- [x] Шаг 1 — пакет `layout`, генерация CSS, ворота, паритет с макетом, инвентарь `.m-*`
- [x] Шаг 2 — `TickRow` и `RowExcerpt` на таблицах (первый потребитель)
- [x] Шаг 3 — остальные варианты `Row`
- [x] Шаг 4 — оболочка, поиск, рейл, сайд
- [x] Шаг 5 — форма, поля, лист, панель
- [x] Шаг 6 — остальные коробки (чип, empty, msg, таблица, дерево, сетка, документ, кнопки)
- [x] **Opus.** Ревью публичного API раскладки (после шага 6, до шага 7)
- [ ] Шаг 7 — `Cols`: контейнер композиции, запасной выход
- [ ] Шаг 8 — демо-экраны 7.2 / 7.3 / 7.6
- [ ] Шаг 9 — полнота: оставшиеся разделы макета
- [ ] Шаг 10 — look-probe и сверка таблиц с `.m-*`

Источник чисел — CSS классов `.m-*` в [docs/visual/carrel-ui-mockups.html](../visual/carrel-ui-mockups.html) (блок после «Классы .m-*»). Если числа в блоке нет — из inline кадра, и в этом файле оно названо явно (панель заметок **440**, цель листа **200**, значок поиска **13**). Из figcaption — никогда.

Куда класть: таблицы в `layout/`, печать слотов в `widget/*.go`. **Запрещено** писать `display`, `grid-template-columns`, ширины дорожек, `min-height` хозяина и padding рядов в `css/*.css`, если для селектора есть таблица. **Запрещено** `style=` на виджете (ворота `no-inline` остаются). **Запрещено** править геометрию `.tc-*` в `internal/screens/screens.css` — и через классы, и через селекторы по `data-tc-`.

После шагов 1–9: `gofmt`, `go test ./layout ./widget ./css ./internal/gate -count=1` из корня модуля. Look (шаг 10) — в конце, иначе каждый шаг будет ждать Chrome.

`--tc-row-height: 38px` остаётся для **обычной** `.tc-row` (как `.m-row { min-height: 38px }`). Его **нельзя** ставить на mark, tree, table td, tickrow, skel.

Нативный попап `input[type=date|time]` не заменять кастомным календарём.

Отладочный зонд сессии 5a1a05 не вставлять.

Файл `docs/plan/typecase-mockup-geometry.md` (ручное копирование px в CSS) **не исполнять и не воскрешать**. Этот план его заменяет.

---

## Что это за слой

На столе дерево виджетов и есть раскладка: контейнер считает детей, `sizeHint` принадлежит виджету, стиль красит. Семь колонок и четыре узла невыразимы.

Сейчас у Typecase четыре описания одного кадра: структура Go, HTML, CSS, макет `.m-*`. Механизм 4 («геометрия — забота виджета») закрыт **наружу** (потребитель без CSS) и не закрыт **внутри**: сетка живёт в другом файле, чем `Render`.

После плана внутри тоже один объект: таблица дорожек. `Render` печатает ровно её слоты. CSS сетки **генерируется** из той же таблицы в `css.Bundle()`. Третьего числа нет.

Не заводить Yoga, constraint solver, канвас, owner-draw. Браузер по-прежнему считает Flex/Grid. Библиотека перестаёт держать вторую копию.

Пакет `layout` потребитель **не импортирует**. Как `token`: модуль его видит, экран приложения зовёт только `widget`.

---

## Словарь: как он растёт и когда сокращается

Правило одностороннее, и оно записывается в спеку (L-3.6), а не держится в голове.

**Наверх — по третьему повторению.** Композиция, собранная экраном из `Cols` и примитивов, переезжает в библиотеку отдельным виджетом, когда встретилась в третий раз. Два раза — совпадение.

**Вниз — почти никогда.** Виджет уходит, только если не держит **ничего** из четырёх: поведения (крючок `data-tc-`), доступности (`role`, `aria-`, порядок обхода), инварианта из своего док-комментария, числа из макета. «Использован один раз» — не основание: `Crop` использован один раз и держит ровно те грабли, ради которых библиотека и заводится.

Автоматических удалений нет. Ворота `layout-orphans` (шаг 9) печатают **отчёт**, а не падают: виджет каталога на ≤1 экране, без `data-tc-`, без `role`/`aria-`, без собственного числа. Отчёт смотрит человек перед релизом.

Перечисление вариантов `Row` (`avatar`, `lead`, `tick`, `excerpt`, `meta`, `find`, `avatarcols`, `tickcols`, `timecols`, `tickmarkcols`) — **временная форма**, а не конечная. Десять имён там, где по существу один виджет с выбираемым набором колонок. Схлопывание `Variant` в именованную раскладку — отдельный план после `Cols`; в этом плане варианты не трогать и новых сверх списка не заводить.

---

## Контейнеры: одно имя `tcbox`

Узкая раскладка считается по ширине **ближайшего прямоугольника, который держит произвольное содержимое**, а не по ширине окна и не по ширине одной заранее выбранной колонки.

Правило одно: каждый такой прямоугольник объявляет контейнер с **одним и тем же именем** `tcbox`, а правила спрашивают `@container tcbox (...)`. CSS берёт ближайшего предка с этим именем — значит строка в списке считает по списку, список в панели — по панели, `Cols` внутри `Cols` — по своей колонке. Ничего выбирать не надо.

`Box` в `layout` печатает `container: tcbox / inline-size` ровно на этих селекторах:

| Селектор | Шаг |
|---|---|
| `.tc-list` | 2 |
| `.tc-shell` | 4 |
| `.tc-main` | 4 |
| `.tc-side` | 4 |
| `.tc-sheet` | 5 |
| `.tc-dialog` | 5 |
| `.tc-cols__col` | 7 |

`Grid.Container` держит **имя** контейнера (`"tcbox"`) или пусто. Селектора там больше нет.

Порядок объявления и потребления: `.tc-list` объявляет контейнер в шаге 2, вместе с первой сеточной строкой. До шага 2 ни одного `@container` в бандле нет. Безымянных `@container` в генерации нет никогда.

`container-type: inline-size` даёт `contain: layout style`: контейнер становится контекстом наложения. Всплывающее меню строки проверяется в look (шаг 10.2).

**Сеточная строка живёт только внутри `List`.** Вне списка — flex-варианты (`RowAvatar`, `RowLead`, `RowTick`, `RowMeta`), у них узкой раскладки нет. Проверяют ворота `layout-slots`: `.tc-row--<сеточный>`, у которого прямой родитель не `.tc-list`, — падение с маршрутом и классом. Так `Row` в `Sec` внутри `Side` (сегодня — вложения заметки в [screens/notes.go](../../internal/screens/notes.go)) остаётся законным и не пытается разложиться на семь колонок в панели 440.

---

## Закрытые развилки

| Вопрос | Решение |
|---|---|
| Где таблицы | `layout/` — стандартная библиотека только. `widget` и `css` его импортируют. `layout` не импортирует `widget` |
| Где в бандле | Сентинел `"generated:layout"` в `css.Order` **после `01-base.css`, перед `02-primitives.css`**. Комментарий-маркер `/* generated: layout */`. Готового CSS в git нет (L-11.1) |
| Почему не в хвосте | В хвосте генерация перебивает уцелевшие ручные `@media` при равной специфичности и молча убивает их — зону галки 44 px (`@media (max-width: 720px)` в [06-lists.css](../../css/06-lists.css)) и `.tc-docmeta { width: auto }` в [09-document.css](../../css/09-document.css). Раньше файлов конфликта нет: свойства, которыми владеет таблица, из ручных файлов удалены воротами, а ручные `@media` остаются уточнением и выигрывают как более поздние |
| Порядок свойств в генерации | Ключи `Box.Props` и `Flow.Props` печатаются **отсортированными**. Иначе бандл отличается от запуска к запуску |
| Inline style | Нет. Ворота `no-inline` |
| Узкая раскладка | `@container tcbox (max-width: N)`. Пороги как в макете: 1190, 800, 700, 620, 460. Не `@media` по окну |
| Узкая оболочка | Поведение (ящик рейла, нижняя навигация, возврат фокуса) остаётся за `tc.js` и классом `.tc-shell--narrow` — не трогать. CSS берёт только то, чего у `tc.js` нет: сайд прячется на 1190 правилом `@container tcbox (max-width:1190px) { .tc-shell:not(.tc-shell--narrow) .tc-shell__side { display: none } }`. Двух источников порога нет: 1190 — только CSS, 640 — только JS, области не пересекаются |
| Рейл 56 px на 560 | **Не переносить.** У Typecase на узкой раскладке рейл — выдвижной ящик, он уже собран и проверен. Расхождение с макетом осознанное, записано в спеку |
| Пустой слот | `EmptyHold` — заглушка узла (дорожка живёт). `EmptyOmit` в `Grid` **запрещён**: пропущенный узел сдвигает все следующие слоты на дорожку влево, а статическая таблица дорожку убрать не может. `EmptyOmit` законен только в `Flow`. Ворота проверяют |
| Обёртка `.tc-row__cells` | У **сеточных** вариантов нет. Дети строки = слоты таблицы |
| Меню и бейджи в сеточной строке | Бейджи печатаются внутрь слота `tags`, если он в таблице есть. `Menu` на сеточном варианте — **падение ворот** с именем варианта, не молчаливая пустота. У вариантов без слота `tags` заполненный `Badges` — тоже падение |
| Ссылка TickRow | Сетка на `.tc-tickrow`. При `Href` слоты внутри `<a class="tc-tickrow__link">`. `Grid.Contents = ".tc-tickrow__link"` → `display: contents`. Нажимаемая область равна объединению слотов: зазоры 9 px не нажимаются, метка `minmax(3.5em,1fr)` закрывает остаток строки. Других обёрток-ссылок у сеток нет |
| Ссылка сеточной `Row` | Класс грида на самом `<a class="tc-row tc-row--…">`. `display:contents` не нужен |
| Сетка без слотов (`repeat(auto-fill,…)`) | Это `Box` с `display:grid` и `grid-template-columns` в `Props`, не `Grid`. Соответствия слотов у неё нет, ворота счёта дорожек к ней не применяются. Так живут `.tc-grid` (плитки) и `.tc-tree` |
| `All()` | Нарастает по шагам. Шаг 1 кладёт только TickRow + RowExcerpt и их Box. Остальное — в том шаге, где виджет переводят. Не заполнять весь план в шаге 1 |
| Типы | Три и только три: `Grid`, `Box`, `Flow`. `layout.Form` — переменная типа `Flow`, не четвёртый вид. Новых полей у типов после шага 1 не заводить |
| Имена вариантов строки | По форме, не по разделу приложения: `avatar`, `lead`, `tick`, `excerpt`, `meta`, `find`, `avatarcols`, `tickcols`, `timecols`, `tickmarkcols`. Не `contact` / `file` / `agenda` / `task` |
| `.m-src.is-root` (дорожки `12px 13px …`) | Не в этом плане. `TickRow` всегда `12px 3px minmax(3.5em,1fr)` |
| Зона галки 44 px при `@media (max-width: 720px)` | Не колонки. Остаётся ручным `@media` в `06-lists.css` и продолжает работать, потому что генерация стоит раньше |
| Месячная сетка дней | Кадра с `.m-*` нет, только фраза в подписи 7.2. Виджет не заводить |
| `m-win` / `m-tabbar` | Не виджеты Typecase. Окно — D, нижняя панель телефона — `NavLayoutBottom` (уже). Геометрию bottom-nav привести к `.m-tabbar` в шаге 4 |
| Печать (§11) | Не в этом плане |
| Yoga / отдельный движок | Нет |
| Подсказки поиска в `utility.go` | Сейчас `RowFind`. После шага 3 это сетка 7.6. Подсказки перевести на `RowLead` (title + caption). `RowFind` только выдача `/find` |
| Spacer формы `:has()` | Не `Flow.Props`. Функция `formSpacerCSS()` в `layout/form.go`, строки ниже в шаге 5 |
| `Hide` на пороге | Как в макете: в блоке порога прячется **только новый** слот. Предыдущие `@container` продолжают действовать |
| Цвет / рамка | `CSS()` сам не красит. Цвет только если ключ есть в `Box.Props` (шаг 6) |

---

## Соответствия макет → таблица

Пары макет → виджет. Куда класть числа — колонка «Таблица», не правка CSS. Узел в средней колонке — тот, на который ложится правило; он не всегда совпадает с корнем виджета.

| Макет | Typecase | Таблица |
|---|---|---|
| `.m-src` | `.tc-tickrow` | `layout.TickRow` |
| `.m-bar3` | `.tc-mark` | хозяин (рейл 14, строка 22, заголовок 20) |
| `.m-row` | `.tc-row` | min-height 38, gap 12, padding gutter; на 620 padding 12 |
| `.m-row--note` | `.tc-row--excerpt` | `layout.RowExcerpt` |
| `.m-row--contact` | `.tc-row--avatarcols` | `layout.RowAvatarCols` |
| `.m-row--file` | `.tc-row--tickcols` | `layout.RowTickCols` |
| `.m-row--find` | `.tc-row--find` | `layout.RowFind` (сетка, не колонка «find») |
| `.m-row--agenda` | `.tc-row--timecols` | `layout.RowTimeCols` |
| `.m-row--task` | `.tc-row--tickmarkcols` | `layout.RowTickMarkCols` |
| `.m-top` / `.m-nav` | `.tc-top` `.tc-nav` | `layout.Top`, `Nav` |
| `.m-search` | ширина → `.tc-search`, коробка → `.tc-search__field` | `layout.Search` (ширина), `layout.SearchField` (высота, gap, padding) |
| `.m-shell` | **`.tc-shell__body`** (не `.tc-shell`: тот держит ещё и `.tc-top`) | `layout.Shell` (рейл 264, сайд 400 / wide 440) |
| `.m-rail` | `.tc-shell__rail` | `layout.Rail` |
| `.m-form` / `.m-f` / `.w-*` | `.tc-form` `.tc-field` `.tc-pfield--*` | `layout.Form`, `FieldWidth` |
| `.m-fields` | `.tc-fields` | 104 px + 1fr, не «по самой длинной» |
| `.m-sheet` | `.tc-sheet` | ширина 560, цель 200×26 |
| `.m-dialog` в main | **не** `.tc-dialog` | `/io` остаётся `Main` |
| overlay confirm | `.tc-dialog` | как `.m-dialog.is-narrow`: **420** |
| `.m-doc` | **`.tc-doc__layout`** (не `.tc-doc`) | `layout.Doc` (296) |
| `.m-table` `.m-grid` `.m-thumb` `.m-transfers` `.m-tree` | `.tc-table` `.tc-grid` `.tc-thumb` `.tc-queue` `.tc-tree` | одноимённые таблицы; `.tc-grid` и `.tc-tree` — `Box` |
| `.m-empty` `.m-msg` `.m-skel` `.m-poll` | `.tc-empty` `.tc-msg` `.tc-skel` `.tc-runbar` | коробки |
| `.m-title` `.m-metacol` | заголовок чтения, `.tc-docmeta` | 27 / 1.15 и 296 |
| `.m-btn.is-danger` / `.is-danger-solid` | `--danger` контур; `--danger-solid` | шаг 6 |
| `.m-h1` | `.tc-head__title` `.tc-side__title` | 22 / 1.15, serif |
| `.m-group` | `.tc-rubric` **в списке** | padding 15 / gutter / 5, низ `--tc-rule` (не 0.35em) |
| `.m-box` в рейле | `.tc-tickrow__tick` | 11×11 |
| `.m-tick` в строке задачи | `.tc-tick` в `tickmarkcols` | 13, колонка 14 |
| `.m-box` в строке файла | `.tc-tick` в `tickcols` | **11×11**, колонка 16 (не 13) |

`tc-field` становится столбцом как `.m-f`. Форма контактов остаётся `Field`/`Form`. `tc-field--bare` для `/notes-quick`. `PField` — тот же столбец; не плодить третий вид поля.

---

## Классы слотов

Закрытый список для строк. Дорожки оболочки, `.tc-fields`, `.tc-doc__layout` и `.tc-cols` слотов не имеют: у них `Track.Slot` пуст, `Empty` не читается, а ворота счёта дорожек сверяют число дорожек с числом детей без сопоставления имён.

| Slot | Класс узла |
|---|---|
| `mark` | `.tc-mark` (Hold: `<span class="tc-mark"></span>`) |
| `tick` | `.tc-tickrow__tick` в TickRow; `.tc-tick` в строке |
| `label` | `.tc-tickrow__label` |
| `title` | `.tc-row__title` |
| `excerpt` | `.tc-row__excerpt` |
| `tags` | `.tc-row__tags` |
| `date` | `.tc-row__date` |
| `source` | `.tc-row__source` |
| `avatar` | корень `Avatar.Render()` (`.tc-avatar`) |
| `subtitle` | `.tc-row__subtitle` |
| `detail` | `.tc-row__detail` |
| `caption` | `.tc-row__caption` |
| `name` | `.tc-row__name` — один узел, внутри значок (`tc-i-*` как у `RowMeta`) и текст |
| `size` | `.tc-row__size` |
| `time` | `.tc-row__time` |
| `lead` | `.tc-row__lead` |
| `due` | `.tc-row__due` |
| `progress` | `.tc-row__progress` — `<progress class="tc-row__progress" max="100" value="N">` если `Cells.Progress != nil`, иначе Hold-заглушка `<span class="tc-row__progress"></span>`. Без `style=` |

`CSS()` для `Hide`: селектор `.Selector .классСлота`, не `:nth-child`.

---

## Шаг 0 — спека

Файл [docs/spec/typecase.md](../spec/typecase.md).

**L-3.1.4** дописать вторым абзацем: геометрия виджета — таблица в пакете `layout`; разметка печатает её слоты; CSS раскладки генерируется из неё. Ручная сетка в `css/*.css` для такого виджета — нарушение ворот. У потребителя по-прежнему нет CSS.

**L-3.6 — новый пункт «Словарь и запасной выход».** Три абзаца:

- Экран компонует и не изобретает. Забор: ни `.tc-` селекторов, ни `data-tc-` селекторов, ни `style=` в коде экрана.
- Когда нужной формы в каталоге нет, экран собирает её у себя из примитивов и `Cols` — именованным типом в своём пакете, не набором переопределений. Библиотеку при этом не правят и не ждут.
- Наверх — по третьему повторению. Вниз — только если виджет не держит ни поведения, ни доступности, ни инварианта, ни числа; решает человек по отчёту `layout-orphans`, автоматических удалений нет.

В таблицу L-8 две строки:

| **Раскладка** | Слотов в HTML не столько, сколько дорожек; `display`, `grid-template-columns` (и перечисленные в таблице свойства) в `css/*.css` у селектора, у которого есть таблица |
| **Размещение** | Сеточная строка вне `List`; `Menu` или `Badges` на сеточном варианте; `EmptyOmit` в `Grid`; безымянный `@container` в генерации |

**L-4 правки геометрии и состава** (не копировать CSS в спеку прозой без таблицы — отослать к тому, что таблица задаёт; здесь — факты, которых в таблице не прочитать):

- `tc-field`: вычеркнуть «подпись слева» и `10em + 1fr`. Подпись сверху 11 px `--tc-muted`; ширины через `FieldWidth` как `.w-*` (96 / 140 / 210 / 300 / 420 / full max 760). Вычеркнуть «ширины по роду, снаружи не задаются». Пустая Label → `--bare`. Инвариант «поля без подписи не существует» заменить: подпись есть или явное bare.
- `tc-form`: выравнивание по строке подписи — свойство раскладки формы (spacer 16 px у check/btn/seg, если у соседа есть label), не трюк экрана.
- `tc-pfield`: те же шесть ширин. Убрать «нет сетки 10em» как отличие от field.
- `tc-fields`: колонка подписей **104 px**, не «по самой длинной».
- `tc-mark`: 14 рейл, 22 строка, 20 заголовок.
- `tc-tickrow`: таблица `layout.TickRow`; ссылка коробки не заводит; заглушка метки всегда при Hold; нажимаемая область ссылки равна объединению слотов.
- `tc-row`: список вариантов как в развилках выше, с пометкой «форма временная». `avatarcols` — первая колонка метка; `tickcols` — метки нет, первая колонка галка 11 px; `timecols` — метка + время; `tickmarkcols` — галка 13 + метка. Сеточная строка живёт только внутри `List`.
- `tc-shell`: узкая раскладка — два непересекающихся механизма. 1190 (сайд) — CSS-контейнер; 640 (ящик рейла, нижняя навигация) — `tc.js`. Рейл 56 px из макета не перенесён, у Typecase на его месте ящик.
- `tc-msg`: подложка `--tc-card`, левая полоса 2 px; тревожный — рамка `--tc-danger` целиком, левая полоса 2 px.
- `ButtonDanger`: контур; заливка — `ButtonDangerSolid`.
- `tc-rubric` в списке: как `.m-group`. В рейле рубрика секции — как `.m-rubric` (10 px, 0.18em).
- `tc-cols`: три формы, числа не задаются снаружи (шаг 7).

[CHANGELOG.md](../../CHANGELOG.md) Unreleased: слой раскладки для потребителя не новый API (пакет `layout` не звать). Новое для потребителя — `Cols`. Ломающее: `ButtonDanger` контур; `Field` подпись сверху и `Width`; новые варианты строки; сеточная строка обязана лежать в `List`; `Menu` на сеточном варианте больше не молчит, а падает на воротах. Кто собирал контакты через `RowAvatar` (две строки текста) и хочет колонки макета — `RowAvatarCols` и заполненные ячейки.

---

## Шаг 1 — пакет и ворота

C-16: сначала проверки красные на текущем коде, потом пакет.

### 1.1 Типы — `layout/spec.go`

Файлы пакета: `spec.go` (типы ниже), `css.go` (`CSS()`), `all.go` (`All()`), `inventory.go` (шаг 1.5). Именованные таблицы — отдельные файлы с шага, который их вводит (`tickrow.go`, `row.go`, …).

```go
type EmptyPolicy uint8
const (
	EmptyHold EmptyPolicy = iota // узел-заглушка, дорожка живёт
	EmptyOmit                    // только в Flow; в Grid — падение ворот
)

type Track struct {
	Slot  string // ключ слота, см. «Классы слотов»; пусто у бесслотовых сеток
	Size  string // фрагмент grid-template без пробела после запятой: "minmax(140px,1.2fr)"
	Empty EmptyPolicy
}

type Breakpoint struct {
	MaxPx    int
	Tracks   []Track
	Hide     []string // Slot, только новые на этом пороге
	PadLeft  string   // блок 620 у строк: padding 12
	PadRight string
}

type Grid struct {
	Selector    string
	Gap         string
	MinHeight   string
	Align       string // "" или "center"
	PadLeft     string
	PadRight    string
	Tracks      []Track
	AutoFlow    string // "" или "column"
	AutoColumns string // "minmax(0,auto)" если AutoFlow задан, иначе ""
	Contents    string // селектор ребёнка с display:contents; "" если нет
	Narrow      []Breakpoint
	Container   string // "tcbox" или ""
}

type Box struct {
	Selector string
	Props    map[string]string // ровно эти свойства, печатаются по отсортированным ключам
}

type Flow struct {
	Selector string
	Props    map[string]string // display/flex-wrap/gap/padding/align-items/flex — не grid-template-*
}

type Spec struct {
	Grids []Grid
	Boxes []Box
	Flows []Flow
}
```

`layout.All() Spec` — то, что сейчас в реестре. `CSS()` печатает:

- `Grid`: `display:grid` + заполненные поля (`gap`, `min-height`, `align-items`, `padding-left`/`padding-right`, `grid-template-columns` из `Tracks`, auto-flow/columns, `Contents { display: contents }`, `@container tcbox (max-width: Npx)` по `Narrow` с `Hide { display:none }` по селектору слота). У `minmax()` запятая без пробела. `minmax(110px,.9fr)` копировать с ведущей точкой, как в макете.
- `Box` / `Flow`: ровно ключи `Props`, **в порядке отсортированных ключей**. Цвет и border — только если ключ есть в `Props`. Grid сам цвет не печатает.
- Если `Grid.Narrow` не пуст, а `Container` пуст — паника при сборке: безымянных `@container` в бандле не бывает.

`All()` в шаге 1 — **две** сетки и их Box (шаг 2). Не весь план.

Тест `layout/css_test.go`: фикстура из трёх дорожек — в CSS ровно они; `@container` не содержит базовых чисел; `minmax(140px,1.2fr)` без пробела после запятой; две пересобранные строки бандла байт в байт равны.

Парсер ворот для `grid-template-columns`: резать по пробелу **снаружи** скобок (`minmax(...)` — один токен). Тот же разбор — в 1.4. Look (10.2) смотрит **вычисленные** колонки (список `Npx`) — там split по пробелу можно, и в вычисленное значение Chrome включает неявные дорожки, поэтому хвост `grid-auto-flow: column` в счёте участвует.

### 1.2 Бандл — `css/bundle.go`

В `Order` появляется сентинел `"generated:layout"` между `01-base.css` и `02-primitives.css`. Встретив его, `Bundle()` пишет `"\n/* generated: layout */\n"` и `layout.CSS()`. `token.Validate` как сейчас.

### 1.3 Ворота — `internal/gate/`

Имя `layout-slots` (в `AllGates`).

1. Для каждого `layout.All()`: в `layout.CSS()` есть `Selector` и столько значений в `grid-template-columns` базового правила, сколько `len(Tracks)` (парсить сгенерированную строку, не руками).
2. В **ручных** `css/*.css` у этих `Selector` нет свойств: `display`, `grid-template-columns`, `grid-template-rows`, `width`/`min-height`/`height`/`padding-left`/`gap`, если они заданы полями `Grid` или ключами `Props`. Иначе — падение с селектором.
3. Парсер ручного CSS: вырезать `/* … */` и блоки `@media`/`@container` целиком — иначе узкое правило выдаст чужое число.
4. Отдельный проход по вырезанным блокам: если ручной `@media`/`@container` задаёт свойство, которым владеет таблица, — **список в тексте ошибки как предупреждение**, не падение. Это уточнения, которые генерация не перебивает (она раньше), но человек должен видеть их перечень.
5. `EmptyOmit` внутри `Grid.Tracks` — падение.
6. Разметка: для каждого `Grid` со слотами взять фикстуры каталога и `screens.All()`; число прямых детей (с раскрытием `Contents`) равно `len(Tracks)`; классы детей совпадают с классами слотов по порядку. Падение с именем виджета, состоянием и обоими числами. Проверка статическая, Chrome не нужен — именно она делает одиннадцатый вариант рабочим с первого раза.
7. Сеточная `.tc-row--*`, у которой прямой родитель не `.tc-list`, — падение с маршрутом.
8. `Row` сеточного варианта с непустым `Menu` — падение. С непустым `Badges` и без слота `tags` в таблице — падение.

Ворота `screens-css` расширить: кроме `.tc-` отвергать `data-tc-` в `internal/screens/screens.css`. Существующее правило `.screens-stage [data-tc-shell] { min-height: 100%; height: 100% }` — снять; высоту сцены держать своим классом на обёртке демо.

В шаге 1 в `All()` уже стоят `TickRow` и `RowExcerpt` (числа шага 2). Ворота красные: excerpt/tickrow ещё в `06-lists.css`. Этот красный — доказательство шага 1; зелёный после удаления дубля в шаге 2.

### 1.4 Паритет с макетом — сразу, не в конце

`internal/gate/mockup_parity.go` заводится **здесь**, а не на шаге 10: Chrome ему не нужен, это разбор файла, и он должен ловить опечатку в числах в том же коммите, где число появилось.

Читать [docs/visual/carrel-ui-mockups.html](../visual/carrel-ui-mockups.html) (нет файла — `t.Skip`). Сверять только те пары, которые уже есть в `All()`; полный список пар — таблица в 10.1. `mockup-parity` в `AllGates` с шага 1.

### 1.5 Инвентарь `.m-*`

`layout/inventory.go`: карта `.m-*` → одно из трёх — селектор Typecase, `NotAWidget` (с причиной: `m-win`, `m-tabbar`, служебные классы страницы макета), `Deferred` (с причиной и ссылкой на план).

Ворота `layout-inventory`: вынуть **все** классы `.m-*` из макета, вычесть карту; остаток — падение со списком. Так «полнота инвентаря» перестаёт быть рукописным утверждением: новый класс в макете обязан быть объявлен, прежде чем план пойдёт дальше.

Заполнить карту целиком в этом шаге — включая то, что переносится позже (`Deferred`). Это дешёвое перечисление, и оно сразу показывает объём.

---

## Шаг 2 — первый потребитель: TickRow и excerpt

Красные тесты `widget` до правки Render:

- `TickRow{Label:"Attachments"}` — ровно один `.tc-mark`.
- HTML excerpt: прямые дети — слоты; **нет** `.tc-row__cells`.

`var TickRow Grid` (`tickrow.go`):

- Selector `.tc-tickrow`
- Align `center`
- Tracks: `{tick, "12px", Hold}`, `{mark, "3px", Hold}`, `{label, "minmax(3.5em,1fr)", Hold}`
- AutoFlow `column`; AutoColumns `minmax(0,auto)`
- Gap `9px`; MinHeight `26px`; PadLeft `12px`; PadRight `14px`
- Contents `.tc-tickrow__link`
- Container пуст: узкой раскладки у tickrow нет

Box там же:

- `.tc-tickrow`: `height: 26px`
- `.tc-tickrow__tick`: `width: 11px`, `height: 11px`
- `.tc-tickrow .tc-mark`: `width: 3px`, `height: 14px`, `min-height: 14px`
- `.tc-row .tc-mark`: `height: 22px`, `min-height: 22px`
- `.tc-head .tc-mark, .tc-side__title .tc-mark`: `height: 20px`, `min-height: 20px`
- `.tc-list`: `container: tcbox / inline-size`

`var RowExcerpt Grid` (`row.go`):

- Selector `.tc-row--excerpt`; Gap `12px`; Align `center`; Container `tcbox`
- Tracks (порядок = дети): mark `3px` Hold, title `minmax(140px,1.2fr)`, excerpt `minmax(170px,2fr)`, tags `minmax(70px,auto)`, date `80px`, source `118px`
- Narrow:
  - 700: Hide `excerpt`; tracks `3px minmax(140px,1.4fr) minmax(70px,auto) 80px 116px`
  - 620: Hide `tags`; tracks `3px minmax(100px,1.4fr) 74px 96px`; плюс `padding-left`/`padding-right` `12px`
  - 460: Hide `source`; tracks `3px minmax(0,1fr) 70px`

`Row.Render` сетки: цикл по `Tracks`, функция `rowSlot(slot, r)` → HTML того класса. Бейджи — внутрь `tags`. Меню сеточная строка не печатает (и не принимает, см. ворота 1.3.8).

После шага: из `06-lists.css` убрать сетку `--excerpt`, её `@media` 700/620/460 и геометрию tickrow (display/height/pad/gap/tick size/mark). `@media (max-width: 720px)` зоны галки **оставить** — генерация стоит раньше, правило продолжает работать. Ворота зелёные для этих двух.

---

## Шаг 3 — остальные строки

В `types.go`: `RowAvatarCols`, `RowTickCols`, `RowTimeCols`, `RowTickMarkCols`. `RowFind` то же имя, раскладка — сетка.

Все Grid: Gap `12px`, Align `center`, Container `tcbox`, порог 620 добавляет `padding-left`/`padding-right` `12px` (как `.m-row` в блоке 620). Строки `Size` ниже — полные `grid-template-columns` без пробела после запятой внутри `minmax`.

**RowAvatarCols** `.tc-row--avatarcols`

- Tracks: mark `3px` Hold, avatar `28px`, title `minmax(150px,1.4fr)`, subtitle `142px`, detail `minmax(150px,1.1fr)`, caption `minmax(110px,.9fr)`, source `120px`
- Narrow: 800 Hide `caption` → `3px 28px minmax(130px,1.4fr) 130px minmax(120px,1fr) 110px`; 620 Hide `detail` → `3px 28px minmax(100px,1.6fr) 118px 96px`; 460 Hide `subtitle` → `3px 28px minmax(0,1fr) 74px`
- Box: `.tc-row--avatarcols .tc-row__subtitle { font-family: var(--tc-font-mono) }`

**RowTickCols** `.tc-row--tickcols`

- Tracks: tick `16px` **Hold** (узла `mark` в этом варианте нет; заглушка галки печатается всегда), name `minmax(220px,2.4fr)`, size `90px`, time `130px`, detail `60px`, source `128px`
- Box: `.tc-row--tickcols .tc-tick { width: 11px; height: 11px }`
- Narrow: 800 Hide `detail` → `16px minmax(200px,2.4fr) 84px 124px 116px`; 620 Hide `time` → `16px minmax(120px,2fr) 78px 96px`; 460 Hide `size` → `16px minmax(0,1fr) 74px`

**RowTimeCols** `.tc-row--timecols`

- Tracks: mark `3px` Hold, lead `62px`, title `minmax(180px,2fr)`, detail `minmax(110px,1fr)`, caption `96px`, source `128px`
- Narrow: 800 Hide `detail` → `3px 60px minmax(140px,2fr) 88px 116px`; 620 Hide `caption` → `3px 54px minmax(110px,1fr) 96px`; 460 Hide `source` → `3px 50px minmax(0,1fr)`

**RowTickMarkCols** `.tc-row--tickmarkcols`

- Tracks: tick `14px` Hold, mark `3px` Hold, title `minmax(190px,2fr)`, tags `minmax(90px,auto)`, due `84px`, progress `44px`, source `128px`
- Box: `.tc-row--tickmarkcols .tc-tick { width: 13px; height: 13px }`
- Narrow: 800 Hide `progress` → `14px 3px minmax(140px,2fr) minmax(60px,auto) 74px 112px`; 620 Hide `tags` → `14px 3px minmax(100px,2fr) 70px 92px`; 460 Hide `source` → `14px 3px minmax(0,1fr) 66px`

**RowFind** `.tc-row--find`

- Tracks: mark `3px` Hold, lead `76px`, title `minmax(220px,2.4fr)`, detail `minmax(120px,1fr)`, caption `74px`, source `128px`
- Narrow: 800 Hide `detail` → `3px 70px minmax(160px,2.4fr) 66px 116px`; 620 Hide `caption` → `3px 62px minmax(110px,2fr) 96px`; 460 Hide `source` → `3px 58px minmax(0,1fr)`
- Слот `lead` на find — род (76 px), не время. Класс всё равно `.tc-row__lead`.
- Подсказки в `utility.go` — `RowLead`, не `RowFind`. Удалить `.tc-row--find .tc-row__find` (колонка-стек) из `06-lists.css`.

Box `.tc-row` (общий): `min-height: 38px`, `gap: 12px`, `padding-left`/`padding-right` `var(--tc-page-gutter)`. `display: flex` — `Flow` для несеточных вариантов.

Демо: `contacts.go` — `RowAvatarCols` с заполненными ячейками; `files.go` список — `RowTickCols`. find/timecols/tickmarkcols — шаг 8, не вхолостую.

`RowAvatar` / `RowMeta` / `RowTick` / `RowLead` не удалять. Их flex — `Flow` с текущими правилами из `06-lists.css` (не грид макета). В `All()` с этого шага. Именно они остаются законными вне `List`.

Каталог: новые варианты во всех нужных состояниях.

---

## Шаг 4 — оболочка

Таблицы/боксы из `.m-top`, `.m-search`, `.m-nav`, `.m-shell`, `.m-rail`, `.m-rail-sec`, `.m-h1`, `.m-head`, `.m-side`.

`layout.Shell` — сетка на **`.tc-shell__body`**, четырьмя записями `Grid` под четыре состояния колонок (`.tc-shell__body`; `…:has(> .tc-shell__side)`; `…:not(:has(> .tc-shell__rail))`; и их сочетание), дорожки как у `.m-shell` / `.has-side` / `.no-rail` / `.no-rail.has-side`. Слотов у них нет.

Числа, которые уже почти стоят, всё равно **переносятся в `layout`**, ручные дубли в `08-shell.css` / `10-utility.css` удаляются:

- топ: min-height 52, padding 16 (не gutter 20), gap 14; крылья `flex: 1 1 0` на brand и `tc-top__actions`
- поиск: `layout.Search` — width 560 на `.tc-search`; `layout.SearchField` — коробка 28 / gap 7 / padding 9 на `.tc-search__field`; значок 13×13 первый узел в `utility.go`; убрать `flex: 1` у `.tc-search` в utility
- нав: item 29, icon 15, gap 9, padding 8
- рейл-колонка: width 264, фон `--tc-bg` (не surface); `.tc-rail` transparent
- секция: gap 3 у рядов; рубрика 10 px / 0.18em / padding 0 16
- сайд: 400, wide 440; title 22 / 1.15; head padding как `.m-side > .m-head`
- `NavLayoutBottom`: высота/зазор как `.m-tabbar` (padding 7 0 9, icon 17, кегль 9.5)

Рейл `background`: цвет — `Box`, не грид.

Контейнеры (`Box`, `container: tcbox / inline-size`): `.tc-shell`, `.tc-main`, `.tc-side`.

Единственное узкое правило оболочки в CSS:

```css
@container tcbox (max-width: 1190px) {
  .tc-shell:not(.tc-shell--narrow) .tc-shell__side { display: none }
}
```

Порог 560 из макета не переносится; `tc.js` и `.tc-shell--narrow` не трогать.

---

## Шаг 5 — форма, панель, лист

`FieldWidth` рядом с `PFieldWidth`, те же шесть значений:

| Значение | Класс | CSS |
|---|---|---|
| XS | `.tc-field--xs` / `.tc-pfield--xs` | `width: 96px` |
| SM | `--sm` | `140px` |
| MD | `--md` | `210px` |
| LG | `--lg` | `300px` |
| XL | `--xl` | `420px` |
| Full | `--full` | `flex: 1 1 100%; max-width: 760px` |

Классы — `layout.Box`.

`var Form Flow` (не новый тип):

```
Selector: .tc-form
Props: display flex; flex-wrap wrap; gap 13px 16px; padding 14px var(--tc-page-gutter); align-items flex-start
```

Box вложенного padding: `.tc-sheet .tc-form, .tc-dialog .tc-form, .tc-side .tc-form { padding-left: 0; padding-right: 0 }`.

`var Field Flow`: `.tc-field { display: flex; flex-direction: column; gap: 4px; min-width: 0 }`.

`formSpacerCSS()` печатает **ровно**:

```css
*:has(> .tc-field > .tc-field__label) > .tc-field:has(> .tc-check:only-child)::before,
*:has(> .tc-field > .tc-field__label) > .tc-field:has(> .tc-btn:only-child)::before,
*:has(> .tc-field > .tc-field__label) > .tc-field:has(> .tc-seg:only-child)::before{content:"";height:16px;flex:none}
.tc-field.is-own-line::before{display:none}
```

Класс `is-own-line` на `Field`, если в Go поле `OwnLine bool` (нужно для одинокой кнопки без spacer). Не писать эти правила в `03-forms.css` руками сверх генерации.

В `io.go` поле Into обернуть в `widget.Form` (сейчас голое в Main — вплотную к краю).

Лист: `.tc-sheet` 560; target как `.m-in.is-pick` псевдоэлементами, 200×26; диалог overlay `.tc-dialog` **420**. Стрелку общего примитива не делать. В спеке у `tc-sheet__target` — почему псевдоэлементы.

Контейнеры (`Box`): `.tc-sheet`, `.tc-dialog` — `container: tcbox / inline-size`.

`notes.go` ряд Date/Time/Tags: SM / XS / MD. Если Tags переносится — Width LG, не CSS ряда.

`.tc-in`: min-height 28, padding 0 9. В баре/шапке 26 — отдельные `Box`. Textarea вне листа 76; в листе 132.

Удалить `.tc-in--date, --time, --number { width: 10em }`.

---

## Шаг 6 — остальные коробки

Перенести в `layout.Box` / `Grid` и вычистить ручной дубль:

- chip и tag = `.m-tag` (10.5 px, padding 0 5, min-height 16, оба `is-on`)
- empty, msg, skel 9 px, runbar / meter 64×3
- table: без рамки обёртки; th 10 px / 0.12em; td padding 7 12, не `--tc-row-height`; кнопка в td 22
- tree row 26, gap 7 — `Box` с `display:grid`, слотов нет
- grid tiles `repeat(auto-fill,minmax(132px,1fr))`, gap 14 — `Box`, не `Grid`; thumb 4/3
- queue 330, absolute в `.tc-main` (контейнер уже даёт `contain: layout`, отдельный `position: relative` не нужен)
- `.tc-doc__layout` — `Grid` без слотов, `minmax(0,1fr) 296px`; docmeta 296; заголовок чтения 27 / 1.15. Ручное `.tc-docmeta { width: 16em }` удалить, `@media (max-width: 720px) { .tc-docmeta { width: auto } }` оставить — генерация раньше, правило работает
- rubric списка = `.m-group`
- линии (цвет, не размер; после правки — контраст в `internal/browser`):

| Токен | Селекторы |
|---|---|
| `--tc-edge` | `.tc-btn`, `.tc-in`, `.tc-seg`, `.tc-search__field`, `.tc-pick__box`, `.tc-tick`, `.tc-tickrow__tick`, `.tc-check__box`, `.tc-kbd__key`, `.tc-menu__control`, `.tc-menu__popover`, `.tc-usermenu__control`, `.tc-usermenu__popover`, `.tc-sheet`, `.tc-dialog`, `.tc-queue` |
| `--tc-hair` | `.tc-table__td`, `.tc-card`, `.tc-tile`, `.tc-diff__option`, `.tc-chip`, `.tc-tag`, `.tc-fset`, `.tc-formfoot`, `.tc-sheet__head`, `.tc-sheet__foot`, `.tc-menu__sep`, `.tc-railfoot` |
| `--tc-rule` (оставить) | `.tc-top`, `.tc-rail`, `.tc-side`, `.tc-bar`, `.tc-table__th`, `.tc-badge`, `.tc-docmeta` |

Каждый селектор линии — `Box{Selector, Props: {"border-color": "var(--tc-edge)"}}` (или `hair`/`rule`). Не писать цвет в комментарии `CSS()` «вообще никогда».

- `ButtonDangerSolid`

Покрытие: существующие `/files` `/states` `/notes-read` `/notes-edit`.

**Opus** после этого шага: вход, перечисления вариантов, инварианты слотов, пакет `layout` не торчит в экраны, форма `Cols` до того, как её напишут. Находки чинит рецензент. Composer не идёт на шаг 7, пока чекбокс Opus не отмечен.

---

## Ревью Opus (после шага 6)

Смотрели: вход виджетов, перечисление вариантов `Row`, инварианты слотов, изоляцию `layout` от экранов, форму `Cols` до кода.

**Исправлено рецензентом.**

- Забор не проверял импорт: экран мог импортировать `layout` и читать числа. Теперь `screens-compose` падает на любом файле `internal/screens`, который импортирует `carrel.dev/typecase/layout` (список закрытых пакетов — `screensForbiddenImports`).
- `no-inline` смотрел только каталог, а «Готово когда» обещает то же про экраны. Теперь он проверяет и тела `screens.All()`.
- `Field`: `Input` не получал `Label` и ставил `aria-label` из машинного `Name`. У поля с подписью имя из `aria-label` перебивало видимую `<label>`, у bare-поля имени не было вовсе (`/notes-quick` читался как «quick»). Теперь `Label` всегда уходит во ввод, а bare задаётся явно: `Field.Bare` снимает подпись с экрана, имя остаётся. Пустая `Label` по-прежнему даёт `--bare` — для совместимости, экраны так не пишут.
- `sheetTargetCSS()` был четвёртым видом описания (голая строка CSS с цветом), в обход трёх типов и ворот. Псевдоэлементы цели листа стали двумя `Box` на `.tc-sheet__target::before` / `::after`; свойства теперь принадлежат таблице, и ворота не дадут ручному CSS их задать. Вне трёх типов остаются только разрешённые планом `formSpacerCSS()` и правило 1190.
- `296px` стояло дважды (`Doc`, `DocMeta`). Теперь это константа `metaColumn`, её же возьмёт `ColsSidebar`.
- Блок типов в 1.1 не показывал `Breakpoint.PadLeft`/`PadRight`, хотя они в коде с шага 1. Дописаны.

**Форма `Cols` — закрыто до шага 7.**

| Вопрос | Решение |
|---|---|
| Имя поля | `Items []Node`, не `Body`. В библиотеке `Body` — поток содержимого, а у `Cols` каждый элемент встаёт в свою дорожку |
| Число `Items` | Равно числу дорожек формы: halves 2, thirds 3, sidebar 2. `Render` печатает то, что дали: не добивает пустыми, не отрезает. Несовпадение — падение `layout-slots` с маршрутом и обоими числами |
| Счёт детей | Сегодня `checkMarkupGrid` пропускает бесслотовые сетки (`!layout.Slotted(g)`). Шаг 7 считает детей у `.tc-cols--*`: у них простой селектор. Оболочку и `Doc` не считать — их вариант выбирает `:has`, статический обход его не выберет |
| Пустой `Cols` | `Items` пуст → `""`, как у `PFieldRow` и `FormFoot` |
| Много однотипных элементов с переносом | Не `Cols`. Это `.tc-grid` (плитки) |
| Сеточная строка внутри колонки | Всё равно внутри `List`. `Cols` — контейнер, не список; ворота 1.3.7 не ослаблять |
| Порог 620 | Меряется по ближайшему `tcbox` **над** `.tc-cols` (сама сетка не контейнер, контейнеры — её колонки). `Cols` внутри колонки `Cols` сворачивается по ширине этой колонки |
| Поля настройки | Только `Shape`. Ни `Gap`, ни ширин, ни выравнивания |

**Открыто — решение за человеком, рецензент не трогал.**

1. Интерактив внутри ссылки. У `TickRow` с `Href` кнопка галки лежит внутри `<a class="tc-tickrow__link">` (закрытая развилка «Ссылка TickRow», так было и до плана). То же у сеточных строк с галкой (`tickcols`, `tickmarkcols`) с `Href`: корень — `<a>`, в нём `<button class="tc-tick">`. HTML недопустим, нажатие на галку уходит по ссылке. Сейчас ни один экран так не собран, поэтому дефект скрыт. Выход — ссылка на слоте подписи (`label` / `name` / `title`), а не на корне. Но это меняет закрытую развилку, спеку `tc-tickrow` и проверку 10.2.
2. ~~Подсказка не связана с вводом.~~ Сделано: у ввода `Field` / `PField` (включая `Check` и `Pick`) `aria-describedby` на подсказку и ошибку; их `id` строятся от `id` ввода, а не от `Name`. Поле для связи у `Input`, `Check`, `Pick` неэкспортируемое — публичный вход не менялся.
3. Ворота «`Menu` на сеточной строке» статические, по AST литерала. Они не видят `Variant` из переменной и `Menu`, присвоенное после литерала, а `Render` такое меню молча выбрасывает. Для каталога и экранов этого хватает; когда появится код приложения — пересмотреть.

---

## Шаг 7 — `Cols`: запасной выход

Один виджет, публичный. Он есть затем, чтобы экран, которому не хватило формы, собрал её у себя, а не пошёл в CSS и не встал в очередь на правку библиотеки.

```go
type ColsShape string

const (
	ColsHalves  ColsShape = "halves"  // minmax(0,1fr) minmax(0,1fr)
	ColsThirds  ColsShape = "thirds"  // три равные
	ColsSidebar ColsShape = "sidebar" // minmax(0,1fr) 296px, как .m-doc
)

type Cols struct {
	Shape ColsShape
	Items []Node
}
```

Разметка: `<div class="tc-cols tc-cols--<shape>">`, каждый ребёнок обёрнут в `<div class="tc-cols__col">`.

`layout.Cols*` — три `Grid` без слотов (`ColsSidebar` берёт `metaColumn`, не второе `296px`), Container `tcbox`, gap `14px` (число `.m-grid`; других чисел для сетки содержимого в макете нет), Narrow один порог 620 → одна колонка `minmax(0,1fr)`, Hide пуст. `Box`: `.tc-cols__col { container: tcbox / inline-size; min-width: 0 }`.

Чисел наружу не отдаётся: потребитель выбирает форму из перечисления, ширины и зазор снаружи не задаются. Забор не двигается — `Cols` не даёт экрану ни px, ни доступа к CSS.

Инвариант: `Cols` не заменяет `Shell`, `Doc` и `Fields` — те формы именованы и живут своими таблицами. `Cols` — для того, чего в каталоге нет.

Каталог: три формы. Отдельный демо-экран не заводить — `Cols` появится на экранах шага 9 там, где понадобится.

Ворота `enum_naming` проверяют `ColsShape` как остальные перечисления: без доменных слов.

---

## Шаг 8 — экраны, которых не было

В `screens.All()` добавить (нейтральные фикстуры, без новых доменных **имён** виджетов):

- `/agenda` — 7.2: оболочка + `TickRow` + список `RowTimeCols` + сайд `Fields`. Переключатель Day/Week/Month/Range — `Seg`; месяц без сетки дней (кадра нет) — выбран Week.
- `/tasks` — 7.3: `RowTickMarkCols`.
- `/find` — 7.6: выдача `RowFind` сеткой, не подсказка поиска.

Покрытие `screens-coverage` само подтянет новые виджеты, если они только там. Каталог — в том же коммите, что виджет.

---

## Шаг 9 — полнота словаря

Это приёмка всего плана, а не витрина. Инвентарь (1.5) отвечает, есть ли виджет на каждую коробку. Здесь отвечают на второй вопрос: складывается ли из этих коробок конкретный кадр. Из списка классов это не выводится — вложение живёт в дереве кадра, а не в CSS.

Собрать по одному маршруту на раздел: 7.8 конфликт, 7.10–7.17 настройки, 7.18 почта, 7.19 админка, 7.20 вход. Одно состояние на раздел, не все.

**Правило записи.** Раздел считается собранным, если он не потребовал ни строки в `screens.css`, ни нового виджета. Каждый случай, где потребовал, дописывается в этот файл списком под заголовком «Чего не хватило»: раздел, что понадобилось, чем обошлись (`Cols` или новый виджет). Это и есть измерение цели; пустой список — цель достигнута, непустой — список того, что чинить, найденный до приложения, а не после.

Композицию, которой в каталоге нет, собирать через `Cols` в пакете `screens` именованным типом. В библиотеку такой тип переезжает по третьему повторению (L-3.6), не раньше.

Ворота `layout-orphans`: печатают отчёт (виджет каталога на ≤1 экране, без `data-tc-`, без `role`/`aria-`, без собственного числа в `layout`). **Не падают.** Отчёт идёт в вывод `go test ./internal/gate -v` и смотрится человеком перед релизом.

---

## Шаг 10 — замок на макет и look

### 10.1 Таблицы = макет

Ворота `mockup-parity` из 1.4 к этому шагу покрывают весь список пар. Сверять **полную строку** `grid-template-columns` (токены снаружи скобок) и эти свойства:

| Пара | Свойства |
|---|---|
| `.m-src` ↔ `layout.TickRow` | `grid-template-columns`, `gap`, `height`, `padding-left`, `padding-right`, `grid-auto-flow`, `grid-auto-columns` |
| `.m-row--note` ↔ `RowExcerpt` | `grid-template-columns` база + каждый `@container` 700/620/460 |
| `.m-row--contact` ↔ `RowAvatarCols` | то же, пороги 800/620/460 |
| `.m-row--file` ↔ `RowTickCols` | то же, 800/620/460 |
| `.m-row--find` ↔ `RowFind` | то же, 800/620/460 |
| `.m-row--agenda` ↔ `RowTimeCols` | то же, 800/620/460 |
| `.m-row--task` ↔ `RowTickMarkCols` | то же, 800/620/460 |
| `.m-row` ↔ Box `.tc-row` | `min-height`, `gap`; блок 620 — `padding-left`, `padding-right` |
| `.m-top` ↔ `layout.Top` | `height` (резолвить `--top` с `.mock` → 52), `gap`, `padding` |
| `.m-search` ↔ `layout.Search` + `layout.SearchField` | `width` у первой; `height`, `gap`, `padding` у второй |
| `.m-nav a` ↔ `layout.Nav` | `height`, `gap` |
| `.m-shell` ↔ `layout.Shell` | `grid-template-columns` всех четырёх состояний; блок 1190 — только строка сайда |
| `.m-form` ↔ `layout.Form` | `gap`, `padding` |
| `.mock .m-fields` ↔ `layout.Fields` | `grid-template-columns` |
| `.m-sheet` ↔ `layout.Sheet` | `width` |
| `.m-dialog.is-narrow` ↔ `.tc-dialog` | `max-width` 420 |
| `.m-grid` ↔ Box `.tc-grid` | `grid-template-columns`, `gap` |
| `.m-doc` ↔ `layout.Doc` | `grid-template-columns` (296) |
| `.m-title` ↔ заголовок чтения | `font-size`, `line-height` |

`.m-top height: var(--top)` — резолвить `--top` с `.mock` (52). Расхождение — оба числа в ошибке.

Авторитет у пары макет ↔ `layout`, не у ручного CSS.

Парсер: комментарии вырезать; `@media`/`@container` в сверке базовых чисел не участвуют (узкое — отдельные поля `Narrow`, их сверять с контейнерными блоками макета по порогу).

Осознанные расхождения, которые паритет **не** сверяет и которые записаны в спеку: рейл 56 px на 560 (у Typecase ящик), `.m-src.is-root`, `m-win`, `m-tabbar`.

C-16: показать красным, подменив одну дорожку в тесте-укусе.

### 10.2 Look — `screens_look_test.go`

Прогон на четырёх ширинах: **1600, 1280, 900, 480** (высота 900). `Harness.Run` уже принимает ширину — это цикл, не новая обвязка. Без узких прогонов контейнерные запросы не проверены ничем.

Ожидаемое **вычисляется из `layout`** по измеренной ширине ближайшего `tcbox`, а не пишется числом в тест. Третьей копии чисел не заводить.

Общее, если узел есть: рейл 264, топ 52, поиск 560 / поле 28, нав 29, tickrow 26, mark рейла 14, title 22 / 1.15.

По маршруту (на 1280, если не сказано иное):

- `/notes`: сайд 440; excerpt слотов = дорожек
- `/notes-quick`: bare; dx textarea/hint ≤ 2; title.x = input.x ±2
- `/files`: chip height ≤ 20; queue 330; tickcols слотов = дорожек; первая td ≤ 24; открытое меню строки — верхний элемент в точке своего центра (`elementFromPoint`): контекст наложения от `container-type` не увёл его под соседа
- `/contacts`: avatarcols, вторая вычисленная колонка `28px`
- `/agenda`: timecols; вторая колонка — то, что даёт `layout` для измеренной ширины списка (на 1280 при рейле 264 и сайде 400 список 616 px, то есть блок 620 и `54px`, а не базовые `62px`)
- `/tasks`: tickmarkcols, первая `14px`
- `/find`: find, вторая — из `layout` по измеренной ширине
- `/io`: левый край Into = заголовок ±1
- `/states`: empty не `text-align: center`
- на 1600: у `/notes` сайд виден; на 900: сайд скрыт (правило 1190), рейл на месте, ящика нет
- на 480: у каждой сеточной строки слотов = дорожек, горизонтальной прокрутки у `body` нет

Счёт слотов (ошибка 7.2 прежнего плана):

```js
function items(row) {
  const walk = (el) => {
    const d = getComputedStyle(el).display;
    if (d === "none") return [];
    if (d === "contents") return [...el.children].flatMap(walk);
    return [el];
  };
  return [...row.children].flatMap(walk);
}
```

Для каждой `.tc-row` с `display:grid` и каждого `.tc-tickrow`: длина **вычисленного** `grid-template-columns` (split по пробелу списка `Npx`, не `none`) === `items(el).length`. В вычисленное значение Chrome включает неявные дорожки, поэтому хвост TickRow (счётчик, меню, слово состояния) считается наравне с явными. Иначе маршрут, класс, оба числа. Заданный CSS с `minmax(...)` в look не разбирать.

TickRow с href: `label.getBoundingClientRect().x` больше, чем `tick.x + 12`, не в трёхпиксельной колонке.

Контраст — прогнать `internal/browser` целиком после смены edge/hair.

---

## Порядок (не параллелить)

0 → 1 → 2 → 3 → 4 → 5 → 6 → **Opus** → 7 → 8 → 9 → 10.

Шаги 3 и 4 после 2: реестр и ворота уже едят ручной CSS. Не переносить все строки, пока excerpt не зелёный — иначе откатывать два семейства сразу.

Шаг 7 после Opus: `Cols` — публичный API, его форму смотрят до того, как она застынет.

Шаг 9 после 7 и 8: без `Cols` собрать оставшиеся разделы нечем, а без 7.2/7.3/7.6 не проверены строки.

## Вне скоупа

- Месячная сетка дней, `m-win`, печать, кастомный date picker, новый `Sheet.Pick`
- `.m-src.is-root`
- Схлопывание `Row.Variant` в именованную раскладку — отдельный план после `Cols`
- Единый порог узкой оболочки (сегодня 1190 в CSS и 640 в `tc.js`) — отдельное решение, не здесь
- Перевод contacts `Form` на `PField`
- Вынос репозитория; как копия макета в `docs/` не разойдётся с оригиналом, иначе паритет и инвентарь после разделения меряют снимок
- Клиент, desktop-приложение, сервер

## Готово когда

`go test ./layout ./widget ./css ./internal/gate ./internal/browser -count=1` зелёный (Chrome для browser), и:

1. **Забор.** `screens.css` не содержит ни `.tc-`, ни `data-tc-`. Ни один экран не имеет `style=`.
2. **Инвентарь.** Каждый класс `.m-*` макета объявлен в карте: перенесён, отложен с причиной или не виджет. Рукописных утверждений о полноте не осталось.
3. **Сборка.** Список «Чего не хватило» шага 9 пуст — либо непуст, и тогда он и есть план следующего круга, а не сюрприз в приложении.
4. **Выход.** Экран, которому не хватило формы, собирает её через `Cols` у себя: без правки библиотеки, без CSS, без ожидания.

Новый экран приложения из этих виджетов не требует CSS экрана. Новый вид строки макета — новая таблица в `layout` плюс слоты в `Row.Render`, не правка `06-lists.css`.
