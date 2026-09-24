# Свёртка разовых классов — что посмотреть глазами

Свёртка по Р-13 (план `1-mockup-tails.md`, Х-1) заменила 67 разовых классов
ближайшим общим и убрала 42 к базовому виду блока. Ниже — каждый кадр, где
сдвинулся рисунок. **Ширины** — ступени переключателя «ширина» наверху макета, на
которых сдвинулся хоть один элемент кадра (замер до и после свёртки на всех десяти
ступенях). **Что смотреть** — какой элемент и что у него поменялось: «было» — разовое
значение, «стало» — общий класс или базовый вид блока.

Если стало хуже — отметьте кадр и строку, верну значение или подберу другой класс.
Когда всё просмотрено, база сверки стилей (`allowed-diffs.md`) переезжает на этот
коммит, а файл удаляется.

## 6. Компоненты

- [ ] **«Источник в рейле»** (`components-11`) · ширины: все ширины
  - секция рейла «Subscriptions external Праздники РФ ro»: было `margin-top:12px` → стало `margin-top:14px`
  - рубрика «Subscriptions external»: было `padding:0 16px 4px` → стало `padding:0 14px 6px`

- [ ] **««New …» в едином виде»** (`components-6`) · ширины: все ширины
  - меню «Write into Контактыlast used Архив»: было `width:250px` → базовый вид (`width: 236px`)

- [ ] **««Upload» в файлах»** (`components-7`) · ширины: все ширины
  - меню «Files Folder»: было `width:216px` → стало `width:212px`

- [ ] **«Шапка экрана · ⋯»** (`components-8`) · ширины: все ширины
  - меню «Import Export Published links»: было `width:238px` → базовый вид (`width: 236px`)

- [ ] **«Кнопка рядом с полем»** (`components-20`) · ширины: 1351, 1191, 1190, 1008, 1007, 820, 641, 640
  - поле формы «Address for the client https://carrel.example/dav/Copy»: было `flex:1 1 260px` → стало `flex:1 1 300px`

- [ ] **«Выбор столбцов»** (`components-21`) · ширины: все ширины
  - меню «Columns Namealways Phone Email»: было `width:258px` → базовый вид (`width: 236px`)
  - кадр «Columns5 Columns Namealways»: было `gap:30px` → базовый вид (`gap: по умолчанию`)

- [ ] **«Дерево папок»** (`components-28`) · ширины: все ширины
  - кадр «attachments 2026 смета»: было `padding:8px` → стало `padding:10px 14px`

- [ ] **«Документ»** (`components-29`) · ширины: все ширины
  - кадр «Save Cancel B Ctrl+S Заметка со встречи 10.08»: было `padding:16px 20px` → стало `padding:18px 20px`

- [ ] **«Шаги мастера и диагностика»** (`components-30`) · ширины: все ширины
  - кадр «1 · Source 2 · Preview 3 · Report 122 ready»: было `gap:12px` → базовый вид (`gap: по умолчанию`)
  - кадр «1 · Source 2 · Preview 3 · Report 122 ready»: было `padding:14px 20px` → стало `padding:18px 20px`

- [ ] **«Крестик панели»** (`components-33`) · ширины: все ширины
  - кадр «Panel title»: было `padding:12px 14px` → стало `padding:14px`

## 7.1 Контакты и человек целиком

- [ ] **«Контакт · переместить в другую книгу»** (`s-contacts-2`) · ширины: все ширины
  - диалог «Move contact Анна Ковалёва · from nextcloud · Контакты»: было `max-width:480px` → стало `max-width:520px`

- [ ] **«Фотография контакта · кадрирование»** (`s-contacts-5`) · ширины: все ширины
  - стопка «Drag the photo to choose the frame. The box is what other cl…»: было `gap:11px` → стало `gap:10px`

## 7.2 Календарь

- [ ] **«Календарь · агенда со списком и карточкой»** (`s-agenda-1`) · ширины: все ширины
  - пустое состояние «Nothing this day.»: было `padding:14px 20px` → стало `font-size:12px; padding:14px 20px`
  - меню «Didn’t fit 2026-08-10→ 2026-08-16»: было `width:252px` → базовый вид (`width: 236px`)
  - tc-range «2026-08-10→ 2026-08-16Show»: было `padding:2px 8px 7px` → базовый вид (`padding: по умолчанию`)

- [ ] **«Событие · правка, повторение»** (`s-agenda-3`) · ширины: 1351, 1191, 1190, 1007, 820, 641, 640, 460, 390
  - диагностика «X-MOZ-GENERATION 3 X-APPLE-TRAVEL-ADVISORY-BEHAVIOR AUTOMATI…»: было `max-width:760px` → базовый вид (`max-width: по умолчанию`)

## 7.3 Задачи

- [ ] **«Задача · правка, срок, выполнение»** (`s-tasks-3`) · ширины: все ширины
  - поле формы «Path /remote.php/dav/calendars/mix/dela/aa10.ics»: было `padding:0` → базовый вид (`padding: по умолчанию`)

## 7.4 Заметки

- [ ] **«Заметка · правка на весь экран, режим сосредоточения»** (`s-notes-4`) · ширины: все ширины
  - полоса: было `height:18px` → базовый вид (`height: 20px`)
  - документ «Заметка со встречи 10.08»: было `padding-top:20px` → базовый вид (`padding-top: 20px var(--tc-page-gutter) 30px`)
  - поле формы «Time 11:05»: было `flex:0 0 110px` → стало `flex:0 0 130px`

- [ ] **«Быстрая заметка · писать некуда»** (`s-notes-6`) · ширины: все ширины
  - пустое состояние «No account connected.»: было `padding:16px` → базовый вид (`padding: 26px var(--tc-page-gutter)`)

- [ ] **«Пункт «Upload file» — куда именно»** (`s-notes-11`) · ширины: все ширины
  - tc-scrim «Upload files»: было `padding-top:20px` → стало `padding-top:26px`
  - лист «Upload files»: было `width:620px` → базовый вид (`width: 560px`)

## 7.5 Файлы

- [ ] **«Файлы · обзор папки, свойства выбранного»** (`s-files-1`) · ширины: все ширины
  - столбец таблицы: было `width:74px` → стало `width:70px`
  - заголовок «2026-08-10-sync-board.png»: было `font-size:16px` → стало `font-size:17px`
  - тело экрана «Contacts312 Calendar14»: было `--tc-side:340px` → стало `--tc-side:360px`

- [ ] **«Файлы · корень «Selected collections»»** (`s-files-6`) · ширины: все ширины
  - tc-table__n4 «Purpose»: было `width:170px` → стало `width:150px`

## 7.8 Конфликт

- [ ] **«Конфликт · планшет, 768 px»** (`s-conflict-2`) · ширины: 1351, 1191, 1190, 1008
  - tc-nav: было `padding:0` → базовый вид (`padding: 0 8px`)

## 7.10 Настройки · учётная запись

- [ ] **«Депонирование ключа · подтверждение»** (`account-2`) · ширины: все ширины
  - поле формы «Your Carrel password •••••••••• Asked because the key is bei…»: было `margin-top:12px; width:100%` → стало `width:100%`

- [ ] **«Настройки · учётная запись · депонирование ключа»** (`account-4`) · ширины: все ширины
  - диалог «Key escrow A copy of your key on this server»: было `max-width:640px` → стало `max-width:560px`

- [ ] **«Клавиши · лист поверх текущего экрана»** (`account-6`) · ширины: все ширины
  - подсказка поля «Shortcuts do nothing while the cursor is in a field. None of…»: было `margin-top:4px` → стало `margin-top:6px`

## 7.11 Настройки · подключения и подписки

- [ ] **«Настройки · подключения и подписки»** (`s-profile-1`) · ширины: все ширины
  - меню «Rename and colour»: было `width:222px` → стало `width:212px`

- [ ] **«Новая коллекция»** (`s-profile-5`) · ширины: все ширины
  - строка «Events Tasks»: было `gap:16px` → стало `gap:14px`
  - тело экрана «Contacts312»: было `min-height:600px` → стало `min-height:560px`

- [ ] **«Переименовать и сменить цвет»** (`s-profile-7`) · ширины: все ширины
  - тело экрана «Back to app»: было `min-height:430px` → базовый вид (`min-height: 480px`)

- [ ] **«Удаление коллекции»** (`s-profile-9`) · ширины: все ширины
  - рубрика «What is pointing at it now»: было `margin-bottom:6px` → стало `margin-bottom:7px`
  - тело экрана «Back to app»: было `min-height:520px` → базовый вид (`min-height: 480px`)

- [ ] **«Настройки · подключения · панель подписки открыта»** (`s-profile-10`) · ширины: все ширины
  - форма «Feed address https://feeds.example/ru-holidays.ics»: было `padding-top:2px` → базовый вид (`padding-top: 14px var(--tc-page-gutter)`)
  - тело экрана «Back to app Settings»: было `--tc-side:400px` → базовый вид (`--tc-side: по умолчанию`)

## 7.13 Настройки · новое устройство

- [ ] **«Android · DAVx5»** (`s-plan-2`) · ширины: 1351, 1191, 1190, 1007, 820, 641, 640, 460, 390
  - стопка «nextcloud · mix»: было `flex:1 1 280px` → стало `flex:1 1 300px`

## 7.14 Настройки · davloom, режим прокси

- [ ] **«davloom · точка подключения и устройства»** (`s-davloom-1`) · ширины: все ширины
  - форма «Address for the client»: было `padding-bottom:6px` → стало `padding-bottom:12px`

- [ ] **«Что видно по адресу прокси в браузере»** (`s-davloom-4`) · ширины: все ширины
  - поле формы «Address https://carrel.example/dav/»: было `grid-template-columns:120px minmax(0,1fr)` → стало `grid-template-columns:130px minmax(0,1fr)`
  - подзаголовок «DAV proxy for Carrel · one address instead of every configur…»: было `margin-top:5px` → базовый вид (`margin-top: 4px`)
  - текст «This is a connection point for clients, not a page to read. …»: было `margin-top:14px; font-size:13.5px` → стало `margin-top:14px`

## 7.15 Настройки · резервная копия и восстановление

- [ ] **«Резервная копия · задание, что и куда»** (`s-backup-2`) · ширины: все ширины
  - карточка «The line to paste visible only now»: было `margin:10px 0 0` → стало `margin:0`
  - сообщение «A file collection is copied as files. Tasks and notes live i…»: было `margin:6px 0 0` → стало `margin:4px 0 0`

- [ ] **«Резервная копия · задание, что и куда»** (`s-backup-3`) · ширины: все ширины
  - tc-fset «Where to To WebDAV»: было `margin-top:0` → базовый вид (`margin-top: 2px`)

- [ ] **«Восстановление · шаг 1, какой архив»** (`s-backup-7`) · ширины: все ширины
  - карточка «•••••••••••• The one set in the job, not your sign-in passwo…»: было `gap:12px` → базовый вид (`gap: по умолчанию`)
  - поле формы «Where the archive is From a job From a server»: было `margin-bottom:12px` → базовый вид (`margin-bottom: по умолчанию`)

- [ ] **«Восстановление · что и куда ляжет»** (`s-backup-8`) · ширины: все ширины
  - полоса «3 selected Restore into»: было `margin:0 0 2px` → базовый вид (`margin: по умолчанию`)
  - поле ввода «radicale · vps · Проекты»: было `height:24px; font-size:12px` → базовый вид (`height: 28px; font-size: 12.5px`)
  - меню «radicale · vps · Проекты»: было `width:230px` → базовый вид (`width: 236px`)

## 7.16 Настройки · публикация по ссылке

- [ ] **«Публикация · выданные ссылки и диалог»** (`s-publish-1`) · ширины: все ширины
  - поле формы «Allow .ics download and subscription»: было `gap:7px; width:100%` → стало `width:100%`
  - форма «How much of it A date range The whole calendar»: было `padding-top:4px` → базовый вид (`padding-top: 14px var(--tc-page-gutter)`)
  - сообщение «A subscribed feed can be handed on by a secret link of its o…»: было `margin-bottom:12px` → базовый вид (`margin-bottom: 12px var(--tc-page-gutter)`)

- [ ] **«Публичная страница · заметка»** (`s-publish-8`) · ширины: все ширины
  - текст «Договорились заморозить ветку в среду. Письмо пользователям …»: было `margin-top:16px` → стало `margin-top:14px`

- [ ] **«Публичная страница · текст и Markdown»** (`s-publish-10`) · ширины: все ширины
  - документ «Перед заморозкой прогнать»: было `padding-top:10px` → базовый вид (`padding-top: 20px var(--tc-page-gutter) 30px`)
  - шапка «release-checklist.md text/markdown · 3.1 KB · changed 2026-0…»: было `padding-bottom:4px` → стало `padding-bottom:6px`

- [ ] **«Публичная страница · файл без предпросмотра»** (`s-publish-11`) · ширины: все ширины
  - кнопка «Download 84.2 MB»: было `height:30px` → базовый вид (`height: 26px`)

- [ ] **«Публичная страница · папка»** (`s-publish-12`) · ширины: все ширины
  - столбец таблицы «Size»: было `width:96px` → стало `width:92px`

## 7.18 Отправка почтой

- [ ] **«Письмо поверх текущего экрана»** (`s-mail-1`) · ширины: все ширины
  - сообщение «Two of these go as links, not as files: the PDF is over the …»: было `margin:10px 0 0` → стало `margin:11px 0 0`

- [ ] **«Поле To подставляет из адресных книг»** (`s-mail-3`) · ширины: все ширины
  - меню «Артём Беловa.belov@example.org work»: было `margin-top:3px; width:100%` → стало `width:100%`

- [ ] **«Отправить ссылку письмом»** (`s-mail-4`) · ширины: все ширины
  - секция «The letter carries»: было `margin:0; padding:10px 0 0` → стало `margin:0; padding:11px 0 0`
  - тело экрана «Back to app»: было `min-height:460px` → базовый вид (`min-height: 480px`)
  - стопка «The link and your line — nothing of the calendar itself»: было `margin-top:6px` → базовый вид (`margin-top: по умолчанию`)

- [ ] **«Выбрать заметку»** (`s-mail-5`) · ширины: все ширины
  - поле ввода «подряд»: было `font-size:12px; width:210px` → стало `font-size:12px; width:200px`
  - секция «Send it as Markdown · .md Source · .ics»: было `margin:0` → базовый вид (`margin: по умолчанию`)

- [ ] **«Выбрать файл»** (`s-mail-6`) · ширины: все ширины
  - столбец таблицы «Size»: было `width:88px` → стало `width:90px`
  - крошки «Documents /Стройка /Подряд»: было `margin-left:10px` → базовый вид (`margin-left: по умолчанию`)
  - tc-tree__row «Подряд»: было `padding-left:50px` → стало `padding-left:44px`

- [ ] **«Служебный адрес, почты в профиле нет»** (`s-mail-8`) · ширины: все ширины
  - сообщение «Отвечать на это письмо не нужно: адрес почту не принимает. С…»: было `font-size:11.5px` → базовый вид (`font-size: 12.5px`)

- [ ] **«Настройки · учётная запись»** (`s-mail-9`) · ширины: все ширины
  - кнопка «Send confirmation»: было `margin-left:6px` → базовый вид (`margin-left: по умолчанию`)

## 7.19 Администрирование

- [ ] **«Администрирование · как заводят пользователей»** (`s-admin-1`) · ширины: все ширины
  - сообщение «Until the user changes this password, you can sign in as the…»: было `margin:0 0 10px` → стало `margin:0 0 12px`

- [ ] **«Администрирование · проверка установки»** (`s-admin-4`) · ширины: все ширины
  - tc-table__n4 «What to do»: было `width:30%` → стало `width:26%`

- [ ] **«Администрирование · журнал»** (`s-admin-7`) · ширины: все ширины
  - поле ввода «All actions»: было `height:26px; width:170px` → стало `height:26px; width:190px`

## 7.20 Вход и первый запуск

- [ ] **«Первый запуск»** (`s-auth-4`) · ширины: 1351, 1191, 1190, 1008, 1007, 820, 641, 640
  - диалог «Set up Carrel The account you create here is the first one, …»: было `max-width:460px; padding-top:0; padding-bottom:0` → стало `max-width:440px; padding-top:0; padding-bottom:0`

- [ ] **«Забыли пароль · логин или почта»** (`s-auth-5`) · ширины: все ширины
  - диалог «Forgotten your password? Enter the login or the confirmed em…»: было `max-width:420px; padding-top:0; padding-bottom:0` → стало `max-width:440px; padding-top:0; padding-bottom:0`
  - текст «Enter the login or the confirmed email of the account. There…»: было `color:var(--tc-muted); margin-bottom:14px` → стало `color:var(--tc-muted); margin:0 0 10px`

- [ ] **«Забыли пароль · что можно сделать»** (`s-auth-6`) · ширины: все ширины
  - форма «Wipe the account We send a letter to the confirmed address. …»: было `padding:16px 0 0` → стало `padding:0`

## 7.21 Остальные экраны и проверка полноты

- [ ] **«О программе»** (`rest-3`) · ширины: все ширины
  - логотип: было `height:39px; width:96px` → стало `height:42px; width:104px`
  - подзаголовок «0.0.1 · 786949a»: было `margin-bottom:10px` → стало `margin-bottom:12px`
  - заголовок «Carrel»: было `margin-bottom:2px` → базовый вид (`margin-bottom: по умолчанию`)

## 8. Состояния

- [ ] **«Загрузка»** (`m-states-3`) · ширины: все ширины
  - заглушка загрузки: было `width:18px` → стало `width:16px`

## 10. Узкий экран

- [ ] **«Агенда»** (`m-phone-1`) · ширины: все ширины
  - полоса опроса «5 of 6 · 14 events · 1 unavailable»: было `font-size:11px` → базовый вид (`font-size: 11.5px`)

- [ ] **«Выдвижной рейл источников»** (`m-phone-2`) · ширины: все ширины
  - пустое состояние «Unticked collections are not polled at all, which saves both…»: было `font-size:12px; padding:16px 14px` → стало `font-size:12px; padding:14px 20px`
  - секция рейла «Combined calendar 14»: было `padding-top:12px` → базовый вид (`padding-top: по умолчанию`)
  - секция рейла «nextcloud Личный 7 Работа»: было `margin-top:10px` → стало `margin-top:14px`

- [ ] **«Контакты»** (`m-phone-4`) · ширины: 1351, 1191, 1190, 1008, 1007, 820, 641
  - аватар: было `height:34px; width:34px` → стало `height:34px; font-size:12px; width:34px`

- [ ] **«Человек целиком»** (`m-phone-5`) · ширины: все ширины
  - поле формы «Phone +7 912 445-10-88 Email a.kovaleva@example.org»: было `grid-template-columns:84px minmax(0,1fr)` → стало `grid-template-columns:88px minmax(0,1fr)`
  - подзаголовок «Издательство «Слово» in 2 books»: было `margin-top:3px` → базовый вид (`margin-top: 4px`)

- [ ] **«Файлы · выделение»** (`m-phone-6`) · ширины: все ширины
  - полоса «2 selected»: было `gap:8px` → базовый вид (`gap: по умолчанию`)

- [ ] **«Заметка · чтение и правка»** (`m-phone-7`) · ширины: все ширины
  - документ «Заметка со встречи 10.08 10 August, 11:05»: было `padding:14px` → базовый вид (`padding: 20px var(--tc-page-gutter) 30px`)
  - заголовок документа «Заметка со встречи 10.08»: было `font-size:21px` → базовый вид (`font-size: 27px`)
  - текст «Договорились заморозить ветку в среду. Письмо пользователям …»: было `margin-top:12px` → стало `margin-top:14px`

## 12. Настольное приложение

- [ ] **«Установка · кому эта копия»** (`m-desktop-3`) · ширины: все ширины
  - логотип: было `height:40px; width:98px` → стало `height:42px; width:104px`
  - стопка: было `background:var(--tc-bg); padding:20px` → стало `background:var(--tc-bg); padding:16px 20px`
  - поле формы: было `margin-top:9px; width:100%` → стало `width:100%`

- [ ] **«Первый запуск · как работать»** (`m-desktop-5`) · ширины: все ширины
  - стопка: было `background:var(--tc-bg); padding:22px` → стало `background:var(--tc-bg); padding:16px 20px`

- [ ] **«Локальный запуск»** (`m-desktop-6`) · ширины: все ширины
  - стопка «Data folder C:\Users\mix\AppData\Local\Carrel»: было `background:var(--tc-bg); padding:18px 20px` → стало `background:var(--tc-bg); padding:16px 20px`

- [ ] **«Первый запуск · к какому экземпляру»** (`m-desktop-7`) · ширины: все ширины
  - стопка «The app needs the address of your Carrel instance. The data …»: было `background:var(--tc-bg); padding:22px 22px 20px` → стало `background:var(--tc-bg); padding:16px 20px`
  - форма «Address https://carrel.example Login»: было `gap:11px` → стало `gap:10px`

- [ ] **«Значок в трее»** (`m-desktop-8`) · ширины: все ширины
  - рубрика «Today · 16 August»: было `margin-bottom:5px` → стало `margin-bottom:7px`

- [ ] **«Напоминания»** (`m-desktop-11`) · ширины: все ширины
  - кадр «Carrel · in 10 minutes»: было `gap:16px` → стало `gap:14px`

- [ ] **«Выход из учётной записи · удалять ли данные»** (`m-desktop-12`) · ширины: все ширины
  - стопка «Sign out of carrel.example ? The app goes back to the first …»: было `background:var(--tc-surface); padding:14px` → стало `background:var(--tc-surface); padding:12px`
