# Разрешённые отличия сверки макета

Сверка `internal/browser/mockup_refactor_test.go` снимает каждый кадр
(`data-frame`) на ширинах **1351** и **820** — видимые **900 px** высоты —
и сравнивает текущую копию макета с базовым коммитом. Пара кадр×ширина
пропускается, только если она записана в таблице.

Один Chrome на кадр. Один кадр:

```
go test ./internal/browser -timeout 8m -run TestMockupRefactor/s-notes/1
```

Раздел:

```
go test ./internal/browser -timeout 5m -run TestMockupRefactor/s-notes
```

Все кадры:

```
go test ./internal/browser -timeout 35m -run TestMockupRefactor
```

Базовый коммит — конец шага 1 плана 1 (токены и шрифты как в продукте).
Когда база задана и файлы разошлись, каждый этап сравнивает до/после.

База: 2cc387df9a4dad7fe72087d7156cb2e4356235f6

| Кадр | Ширина | Решение | Что меняется |
|------|-------:|---------|----------------|
| s-contacts-3 | 1351 | шаг 4 | нативный input type=date/time |
| s-contacts-3 | 820 | шаг 4 | нативный input type=date/time |
| s-contacts-4 | 1351 | шаг 4 | нативный input type=date/time |
| s-contacts-4 | 820 | шаг 4 | нативный input type=date/time |
| s-contacts-7 | 1351 | шаг 4 | нативный input type=date/time |
| s-contacts-7 | 820 | шаг 4 | нативный input type=date/time |
| s-agenda-1 | 1351 | шаг 4 | нативный input type=date/time |
| s-agenda-1 | 820 | шаг 4 | нативный input type=date/time |
| s-agenda-3 | 1351 | шаг 4 | нативный input type=date/time |
| s-agenda-3 | 820 | шаг 4 | нативный input type=date/time |
| s-tasks-3 | 1351 | шаг 4 | нативный input type=date/time |
| s-tasks-3 | 820 | шаг 4 | нативный input type=date/time |
| s-notes-1 | 1351 | шаг 4 | нативный input type=date/time |
| s-notes-1 | 820 | шаг 4 | нативный input type=date/time |
| s-notes-4 | 1351 | шаг 4 | нативный input type=date/time |
| s-notes-4 | 820 | шаг 4 | нативный input type=date/time |
| s-notes-9 | 1351 | шаг 4 | нативный input type=date/time |
| s-notes-9 | 820 | шаг 4 | нативный input type=date/time |
| s-notes-10 | 1351 | шаг 4 | нативный input type=date/time |
| s-notes-10 | 820 | шаг 4 | нативный input type=date/time |
| s-publish-1 | 1351 | шаг 4 | нативный input type=date/time |
| s-publish-1 | 820 | шаг 4 | нативный input type=date/time |
| s-publish-2 | 1351 | шаг 4 | нативный input type=date/time |
| s-publish-2 | 820 | шаг 4 | нативный input type=date/time |
| s-admin-7 | 1351 | шаг 4 | нативный input type=date/time |
| s-admin-7 | 820 | шаг 4 | нативный input type=date/time |
| m-desktop-10 | 1351 | шаг 4 | нативный input type=date/time |
| m-desktop-10 | 820 | шаг 4 | нативный input type=date/time |
| s-profile-1 | 1351 | шаг 4 | иконки пунктов меню после выноса в .tc-drop |
| s-profile-1 | 820 | шаг 4 | иконки пунктов меню после выноса в .tc-drop |
| s-backup-8 | 1351 | шаг 4 | иконки пунктов меню после выноса в .tc-drop |
| s-backup-8 | 820 | шаг 4 | иконки пунктов меню после выноса в .tc-drop |
