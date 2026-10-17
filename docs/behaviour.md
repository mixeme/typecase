# Поведение виджетов

Таблица «виджет × поведение × проверка» — пятый пункт ревью библиотеки
(решение человека 2026-10-01, сделано 2026-10-14). Поведение — пункт
«Поведение» описания виджета (L-3.2); проверка — браузерные тесты
(`internal/browser`), в которых виджет печатается, либо почему теста не
нужно. Собрана сценарием из описаний и тестов; при новом виджете или
поведении — дописать строку.

**Итог.** 98 виджетов с описанием; 54 печатаются в браузерных тестах. У 19
поведение описано, а своего теста нет: 14 работают без сценария (ссылка,
форма, `details`), 3 — общим правилом, проверенным на другом виджете, у 2
описание обещало то, чего в `tc.js` нет (`Filter` — сужение по мере набора,
`TickRow` — галка группы ставит галки под ней): описания исправлены, сами
поведения — в хвостах (`docs/roadmap.md`).

| Виджет | Блок | Поведение | Проверка |
|---|---|---|---|
| `Acts` | `tc-acts` | нет: поведение у кнопок. | тест: `TestZonePanel` |
| `Area` | `tc-in` | Ввод — браузера. | тест: `TestAddRowAndDrop`, `TestNoScriptSheetAndChecks`, `TestOpensInPlace`, `TestPasteAndQueue` |
| `Auth` | `tc-auth` | нет. | — |
| `Avatar` | `tc-avatar` | нет. | — |
| `Badge` | `tc-badge` | нет. | тест: `TestNotebookPick` |
| `Bar` | `tc-bar` | «⋯» открывает меню с тем, что не поместилось. | тест: `TestSelectionForm` |
| `Brand` | `tc-brand` | нет. | — |
| `BriefRow` | `tc-row` | Нажатие выбирает строку. | без сценария (ссылка, форма, `details`) |
| `Button` | `tc-btn` | Нажатие — действие потребителя; кнопка со стрелкой (Menu) открывает свой список через узел меню, а не сама; клавиша (Key) нажимает кнопку, тревожная клавиши не несёт. | тест: `TestAddRowAndDrop`, `TestCheckFormInMenu`, `TestCheckFormSubmits`, `TestClosedMenuWithoutScripts`, `TestColumnsMenu`, `TestConflictForm`, `TestCopy`, `TestEditFile`, `TestFilePick`, `TestFileSizes`, `TestFormsWithoutScripts`, `TestNoScriptSheetAndChecks`, `TestOpensInPlace`, `TestPasteAndQueue`, `TestRailReopensSources`, `TestRemoveRow`, `TestSaveAs`, `TestScriptOnly`, `TestSelectionForm`, `TestShortcuts`, `TestSignInOutsideItsForm`, `TestSwap`, `TestTallyCount`, `TestTemplateShows`, `TestTilePick`, `TestTreeNodeSwap`, `TestTreeSend`, `TestZoneFormOnce`, `TestZoneSheet`, `TestZones` |
| `Card` | `tc-card` | нет: поведение у того, что в ней. | — |
| `Check` | `tc-check` | Нажатие переключает; пробел — тоже. | тест: `TestCheckFormInMenu`, `TestCheckFormSubmits`, `TestConflictForm`, `TestEndFollowsStart`, `TestFormsWithoutScripts`, `TestNoScriptSheetAndChecks`, `TestSwap`, `TestTallyCount` |
| `CheckForm` | `tc-checkform` | Смена отметки отправляет форму (сценарий); без сценария — Apply рядом. | тест: `TestCheckFormInMenu`, `TestCheckFormSubmits` |
| `Close` | `tc-close` | Закрывает свой слой; Esc делает то же, фокус возвращается на вызвавший контрол. | тест: `TestShortcuts`, `TestZonePanel` |
| `ColumnRow` | `tc-row` | Нажатие выбирает строку; строка со ссылкой ведёт на запись; галка дела отмечает его сделанным. | без сценария (ссылка, форма, `details`) |
| `Crumbs` | `tc-crumbs` | Нажатие на шаг открывает ту папку. | без сценария (ссылка, форма, `details`) |
| `Danger` | `tc-danger` | Нажатие спрашивает подтверждения листом. | общим правилом: общее открытие листа (`Button.Opens`, `TestOpensInPlace`) |
| `Diag` | `tc-diag` | нет. | тест: `TestCopy`, `TestSaveAs`, `TestZonePanel` |
| `Dialog` | `tc-dialog` | Esc закрывает диалог, если он открыт поверх страницы, и возвращает фокус. | тест: `TestSignInOutsideItsForm` |
| `Diff` | `tc-diff` | Нажатие выбирает вариант свойства. | тест: `TestConflictForm` |
| `Dim` | `tc-dim` | нет. | тест: `TestEditFile`, `TestFileSizes`, `TestRailReopensSources`, `TestZonePick` |
| `Doc` | `tc-doc` | Галка дела в тексте — знак (Tick.Sign): дело отмечают в правке; над таблицей в узком контейнере переключатель «Wrap / One line» держит строки таблицы в одну линию. | тест: `TestDocTableLine` |
| `Docbar` | `tc-docbar` | Ctrl+S сохраняет; «⋯» открывает то, что не поместилось. | тест: `TestSwap` |
| `Docmeta` | `tc-docmeta` | нет: поведение у ссылок и кнопок. | — |
| `Drop` | `tc-drop` | Нажатие открывает меню под кнопкой; Esc закрывает и возвращает фокус; у выбора с меню из радио (пояс) выбранный пункт становится значением выбора. | тест: `TestAttendeesSuggest`, `TestCheckFormInMenu`, `TestClosedMenuWithoutScripts`, `TestEditFile`, `TestNotebookPick`, `TestShortcuts`, `TestZonePick` |
| `DueRow` | `tc-row` | Галка отмечает дело сделанным; нажатие на строку открывает его; в списке только для чтения галка — знак (Tick.Sign): дело не отметить. | без сценария (ссылка, форма, `details`) |
| `Emblem` | `tc-emblem` | нет. | — |
| `Empty` | `tc-empty` | Бросок файла на зону загружает его. | тест: `TestFileSizes` |
| `ExcerptRow` | `tc-row` | Нажатие открывает текст; стрелки ходят по строкам списка. | тест: `TestNoScriptSheetAndChecks`, `TestShortcuts` |
| `FaceRow` | `tc-row` | Нажатие открывает карточку; стрелки ходят по строкам списка. | тест: `TestZonePanel` |
| `Field` | `tc-field` | Поле с Changed появляется, пока названные поля формы отличаются от пришедших со страницей, и прячется, когда правку вернули; остальное — у контролов. | тест: `TestAddRowAndDrop`, `TestEditFile`, `TestEndFollowsStart`, `TestFormsWithoutScripts`, `TestNoScriptSheetAndChecks`, `TestPickShows`, `TestRemoveRow`, `TestShortcuts`, `TestSignInOutsideItsForm`, `TestTemplateShows`, `TestZoneFormOnce`, `TestZoneSheet`, `TestZones`, `TestFieldChanged` |
| `FieldSet` | `tc-fset` | Группа с адресом (Action) — своя форма: её кнопка шлёт её поля, а не поля соседних групп. | тест: `TestAddRowAndDrop`, `TestEndFollowsStart`, `TestPickShows`, `TestRemoveRow`, `TestTemplateShows` |
| `Fields` | `tc-fields` | нет. | — |
| `Filter` | `tc-in` | Сужает список по мере набора; Esc очищает. | **описание обещало несделанное — исправлено** (14.10) |
| `FindRow` | `tc-row` | Нажатие открывает запись. | без сценария (ссылка, форма, `details`) |
| `Form` | `tc-form` | Enter в поле формы отправляет форму. | тест: `TestAddRowAndDrop`, `TestEndFollowsStart`, `TestFormsWithoutScripts`, `TestNoScriptSheetAndChecks`, `TestPickShows`, `TestRemoveRow`, `TestSignInOutsideItsForm`, `TestTemplateShows`, `TestZoneFormOnce`, `TestZones` |
| `FormFoot` | `tc-formfoot` | нет: поведение у кнопок. | — |
| `Grid` | `tc-grid` | Стрелки ходят по плиткам. | тест: `TestSelectionForm`, `TestTilePick` |
| `Grow` | `tc-grow` | нет. | тест: `TestRemoveRow` |
| `Head` | `tc-head` | нет: поведение у действий. | тест: `TestRailReopensSources`, `TestZonePanel` |
| `Input` | `tc-in` | Ввод — браузера; нативный выбор даты и времени не заменяется. Конец, названный идущим за началом (Follows), сценарий сдвигает на столько же, на сколько сдвинули начало; без сценария он стоит. | тест: `TestAddRowAndDrop`, `TestEndFollowsStart`, `TestNoScriptSheetAndChecks`, `TestPickShows`, `TestRemoveRow`, `TestShortcuts`, `TestSignInOutsideItsForm`, `TestTemplateShows`, `TestZoneFormOnce`, `TestZonePick`, `TestZoneSheet`, `TestZones` |
| `Kbd` | `tc-kbd` | нет. | — |
| `Line` | `tc-line` | нет. | тест: `TestEditFile`, `TestFilePick`, `TestFileSizes` |
| `LinkedRow` | `tc-row` | Нажатие открывает запись. | без сценария (ссылка, форма, `details`) |
| `List` | `tc-list` | Стрелки ходят по строкам; Enter открывает строку. | тест: `TestNoScriptSheetAndChecks`, `TestShortcuts`, `TestZonePanel` |
| `Main` | `tc-main` | нет: поведение у того, что в ней. | тест: `TestRailReopensSources`, `TestZonePanel` |
| `Mark` | `tc-mark` | нет. | тест: `TestAttendeesSuggest`, `TestNoScriptSheetAndChecks`, `TestNotebookPick` |
| `Menu` | `tc-menu` | Стрелки ходят по пунктам, Enter выбирает, Esc закрывает меню и возвращает фокус кнопке, которая его открыла; строка поиска оставляет пункты, где есть набранное (без сценария её нет). | тест: `TestAttendeesSuggest`, `TestCheckFormInMenu`, `TestClosedMenuWithoutScripts`, `TestColumnsMenu`, `TestEditFile`, `TestFileInForm`, `TestFilePick`, `TestFileSizes`, `TestNotebookPick`, `TestSearchSuggest`, `TestShortcuts`, `TestZonePick` |
| `Meter` | `tc-meter` | нет. | тест: `TestEditFile`, `TestPasteAndQueue`, `TestRunFollows` |
| `Mono` | `tc-mono` | нет. | тест: `TestEditFile`, `TestPasteAndQueue` |
| `Msg` | `tc-msg` | нет. | тест: `TestAppWorker`, `TestPageUnderStrictPolicy`, `TestScriptOnly`, `TestZonePanel` |
| `Nav` | `tc-nav` | Нажатие открывает раздел; текущий помечен для чтения (aria-current). В рейле из значков (641–1007) повторное нажатие на текущий раздел открывает лист его источников (Opens). | тест: `TestRailReopensSources`, `TestZones` |
| `Num` | `tc-num` | нет. | — |
| `Overlay` | `tc-overlay` | Пока лист открыт, фокус не уходит под затемнение. | общим правилом: ловушка фокуса затемнения (`TestZoneSheet`) |
| `Pad` | `tc-pad` | нет. | — |
| `Page` | `tc-root` | нет. | тест: `TestAppWorker`, `TestPageUnderStrictPolicy` |
| `PhoneRow` | `tc-row` | Нажатие открывает запись; галка выбирает её. | без сценария (ссылка, форма, `details`) |
| `Pick` | `tc-in` | Нажатие открывает список под полем; выбор закрывает его, фокус возвращается на поле. | тест: `TestFormsWithoutScripts`, `TestNotebookPick`, `TestPickShows`, `TestTemplateShows`, `TestZonePick` |
| `Plain` | `tc-plain` | нет. | — |
| `PostForm` | `tc-post` | Отправку делает браузер — сценарий не нужен. | тест: `TestAddRowAndDrop`, `TestPasteAndQueue`, `TestSelectionForm`, `TestSideDrop` |
| `Prose` | `tc-prose` | нет. | — |
| `Push` | `tc-push` | нет. | — |
| `Query` | `tc-q` | Ввод — браузера; Esc очищает, если поле не пусто; клавиша (Key) ставит в строку фокус. | тест: `TestSearchSuggest`, `TestShortcuts` |
| `Queue` | `tc-queue` | «Hide» сворачивает очередь в строку, «Show» разворачивает. | тест: `TestPasteAndQueue` |
| `RailFoot` | `tc-railfoot` | нет: поведение у кнопок. | — |
| `RailSec` | `tc-railsec` | нет: поведение у строк и тегов. | — |
| `RailStatus` | `tc-railstatus` | нет. | — |
| `Range` | `tc-range` | Show применяет диапазон; Enter в поле — то же. С адресом (Action) диапазон — форма GET: From и To уходят полями, Show её отправляет, без сценария. | без сценария (ссылка, форма, `details`) |
| `Rubric` | `tc-rubric` | нет. | тест: `TestAddRowAndDrop`, `TestConflictForm`, `TestEndFollowsStart`, `TestPickShows`, `TestRemoveRow`, `TestSearchSuggest`, `TestTemplateShows`, `TestZonePick` |
| `Runbar` | `tc-runbar` | «Stop» прерывает опрос; строка-область с адресом потока или опроса следит за ходом сама. | тест: `TestEditFile`, `TestRunFollows` |
| `Scrim` | `tc-scrim` | Нажатие на затемнение закрывает лист, как Esc; кнопка с Opens открывает закрытое затемнение на месте. | тест: `TestAddRowAndDrop`, `TestNoScriptSheetAndChecks`, `TestOpensInPlace`, `TestPasteAndQueue`, `TestRailReopensSources`, `TestShortcuts`, `TestSideDrop`, `TestZoneSheet` |
| `Search` | `tc-search` | Ввод ищет; Esc очищает запрос и закрывает подсказки. | тест: `TestSearchSuggest`, `TestShortcuts` |
| `Sec` | `tc-sec` | нет. | — |
| `Seg` | `tc-seg` | Нажатие выбирает вариант; стрелки ходят по вариантам внутри переключателя. | тест: `TestFormsWithoutScripts`, `TestNoScriptSheetAndChecks`, `TestSwap`, `TestThemeShowsKept`, `TestThemeFromServer` |
| `SegFold` | `tc-segfold` | На узкой полосе нажатие на текущий вариант открывает меню остальных (details, без сценария); выбор — переход по ссылке. Короткий переключатель (`Phone`) складывается только на полосе телефона, 640&nbsp;px и уже: Day / Week / Month / Range агенды. | без сценария (ссылка, форма, `details`) |
| `Sheet` | `tc-sheet` | Esc закрывает лист и возвращает фокус тому, что его открыло; Enter делает главное. | тест: `TestAddRowAndDrop`, `TestFileSizes`, `TestNoScriptSheetAndChecks`, `TestOpensInPlace`, `TestPasteAndQueue`, `TestRailReopensSources`, `TestShortcuts`, `TestTallyCount`, `TestTreeNodeSwap`, `TestTreeSend`, `TestZoneSheet` |
| `Shell` | `tc-shell` | На узком экране рейл становится ящиком, панель — листом поверх страницы. | тест: `TestRailReopensSources`, `TestZonePanel` |
| `Side` | `tc-side` | Esc закрывает панель и возвращает фокус строке, которая её открыла. Панель с Drop: файл над ней отмечает её (drop) вместо листа броска страницы, брошенный — уходит в названную форму. | тест: `TestSideDrop`, `TestZonePanel` |
| `SizeRow` | `tc-row` | Галка выбирает вещь для действий шапки; нажатие на имя открывает её. | без сценария (ссылка, форма, `details`) |
| `Skel` | `tc-skel` | нет. | — |
| `SkelRow` | `tc-row` | нет. | — |
| `SlimRow` | `tc-row` | Галка выбирает строку. | без сценария (ссылка, форма, `details`) |
| `Split` | `tc-split` | Главная половина делает главное; стрелка открывает меню. | без сценария (ссылка, форма, `details`) |
| `Stack` | `tc-stack` | нет. | — |
| `Stat` | `tc-stat` | нет. | — |
| `Stats` | `tc-stats` | нет. | — |
| `Steps` | `tc-steps` | нет: шаги не нажимаются, переход — кнопками мастера. | — |
| `Table` | `tc-table` | Галка выбирает строку для действий бара; кнопки строки с `is-2nd` на узкой ширине уходят под «⋯» строки (`TableCell.More`), и «⋯» открывает её меню; имя со ссылкой (`TableFile.Href`) открывает строку. | тест: `TestSelectionForm` |
| `Tag` | `tc-tag` | Выбранный фильтр (On) снимается нажатием; тег с крестиком убирается крестиком. | тест: `TestAttendeesSuggest` |
| `TagInput` | `tc-in` | С Name каждый тег отправляет своё значение, снятый крестиком — нет, и без сценария тоже. С живым слоем (tc-live.js) набор открывает подсказки под полем (Add.Suggest: сервер отвечает пунктами меню); выбор добавляет тег; крестик и Backspace в пустой строке снимают тег; Enter не отправляет форму; Esc закрывает подсказки. | тест: `TestAttendeesSuggest`, `TestTagList` |
| `Thumb` | `tc-thumb` | Картинку, которая не пришла, сценарий убирает: остаётся значок. | тест: `TestSelectionForm`, `TestThumbPictureGone`, `TestTilePick` |
| `Tick` | `tc-tick` | Нажатие переключает отметку; знак (Glyph) не нажимается: нажимается кнопка, в которой он стоит; галка дела в тексте (Sign) — тоже: заметка правится в правке (задача Carrel 99). | без сценария (ссылка, форма, `details`) |
| `TickRow` | `tc-tickrow` | Нажатие на галку показывает или скрывает источник; галка группы ставит или снимает все галки под ней. | **описание обещало несделанное — исправлено** (14.10) |
| `Tile` | `tc-tile` | Нажатие выбирает плитку: с `Send` она — флажок формы во всю плитку, без своего вида, выбранную рисует рамка; нажатие на имя со ссылкой открывает файл. | тест: `TestSelectionForm`, `TestTilePick` |
| `TimeRow` | `tc-row` | Нажатие открывает запись; стрелки ходят по строкам списка. | без сценария (ссылка, форма, `details`) |
| `Top` | `tc-top` | нет: поведение у поиска, кнопок и меню учётной записи. | тест: `TestSearchSuggest` |
| `Tree` | `tc-tree` | Нажатие на строку выбирает её; недоступную строку выбрать нельзя; строка с `Send` — радиокнопка формы, в которой стоит дерево, и выбирает без сценария. | тест: `TestTreeNodeSwap`, `TestTreeSend` |
| `UserMenu` | `tc-usermenu` | Нажатие открывает меню учётной записи; Esc закрывает его и возвращает фокус кнопке. | общим правилом: общий Esc меню (тесты меню) |
| `Zone` | `tc-zone` | нет. | — |

## Против настольных библиотек

Что у виджета Qt, GTK или Cocoa есть, а у нашего — нет, и стоит ли
подтягивать, с учётом веба: страница печатается на сервере, работает без
сценария, состояние — в DOM и в отправленной форме.

- **Вид и дети** — полно: кадр, эталон, типы частей; ничего не нужно.
- **Смысл и доступность** — хорошо (ворота `a11y`, роли и подписи из
  разметки).
- **Родитель** — виджет о нём не знает; место задают вариант и CSS. Правило
  «где виджету стоять» уже держат пары (`container-pairs`) в обе стороны;
  обратного контракта сверх них не нужно.
- **Поведение** — эта таблица: слова в описании, код — общий `tc.js` по
  селекторам. Расхождение слов и кода нашлось дважды; таблица — способ его
  видеть.
- **Соседи** — подпись и поле, группа радио, порядок фокуса — отдаём
  браузеру: `label`, общее `name`, порядок в DOM. Своей раскладки и
  фокус-цепочек, как в настольных библиотеках, не нужно: браузер делает это
  сам и одинаково без сценария.
- **Сигналы и события, своё состояние** — сознательно нет: состояние —
  поля формы, событие — её отправка; сигналы настольных библиотек заменяет
  сервер.
