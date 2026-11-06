// Раздел макета: 11. Печать.

package screens

import w "carrel.dev/typecase/widget"

// Print1 is frame «Агенда на бумаге» (print-1).
func Print1() Screen {
	return w.Page{
		Items: []w.PagePart{
			w.PrintTitle{Text: "Agenda · 10 – 16 August 2026"},
			w.PrintSources{Text: "nextcloud · Личный, Работа · radicale · vps · Проекты, Дежурства"},
			w.PrintRubric{Text: "Monday · 10 August"},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "10:00", Variant: w.MonoFace},
					w.PrintCell{Name: "Синк-ап по релизу", Detail: "Переговорная 2 · nextcloud · Личный"},
				},
			},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "all day", Variant: w.MonoFace},
					w.PrintCell{Name: "Дежурство: первая линия", Detail: "radicale · vps · Дежурства"},
				},
			},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "15:00", Variant: w.MonoFace},
					w.PrintCell{Name: "Созвон с подрядчиком", Detail: "Jitsi · nextcloud · Работа"},
				},
			},
			w.PrintRubric{Text: "Tuesday · 11 August"},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "09:30", Variant: w.MonoFace},
					w.PrintCell{Name: "Ревью архитектуры", Detail: "Переговорная 1 · radicale · vps · Проекты"},
				},
			},
			w.PrintFoot{Text: "Carrel · agenda · 2026-08-10 – 2026-08-16 · captured 2026-08-15 09:41 · 5 of 7 sources"},
		},
		Variant: w.PagePrint,
	}
}

// Print2 is frame «Список контактов на бумаге» (print-2).
func Print2() Screen {
	return w.Page{
		Items: []w.PagePart{
			w.PrintTitle{Text: "Contacts · Combined address book"},
			w.PrintSources{Text: "nextcloud · Контакты, Архив · radicale · vps · Команда"},
			w.PrintRow{
				Items:   []w.PrintRowPart{w.PrintCell{Name: "Name"}, w.PrintNum{Text: "Phone"}, w.PrintNum{Text: "Email"}},
				Columns: true,
				Head:    true,
			},
			w.PrintRubric{Text: "А"},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.PrintCell{Name: "Алексей Смирнов", Detail: "nextcloud · Контакты"},
					w.PrintNum{Text: "+7 916 200-71-15"},
					w.PrintNum{Text: "—"},
				},
				Columns: true,
			},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.PrintCell{Name: "Анна Ковалёва", Detail: "Издательство «Слово» · in 3 books"},
					w.PrintNum{Text: "+7 912 445-10-88"},
					w.PrintNum{Text: "a.kovaleva@example.org"},
				},
				Columns: true,
			},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.PrintCell{Name: "Артём Белов", Detail: "Тензор · in 2 books"},
					w.PrintNum{Text: "+7 903 118-22-04"},
					w.PrintNum{Text: "a.belov@example.org"},
				},
				Columns: true,
			},
			w.PrintRubric{Text: "В"},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.PrintCell{Name: "Вера Лебедева", Detail: "nextcloud · Архив"},
					w.PrintNum{Text: "+7 905 331-09-42"},
					w.PrintNum{Text: "vera@example.net"},
				},
				Columns: true,
			},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.PrintCell{Name: "Виктор Гончаров", Detail: "Городская библиотека №4 · nextcloud · Контакты"},
					w.PrintNum{Text: "—"},
					w.PrintNum{Text: "v.goncharov@example.org"},
				},
				Columns: true,
			},
			w.PrintRubric{Text: "Л"},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.PrintCell{Name: "Лебедева Вера", Detail: "Библиотека №4 · nextcloud · Контакты"},
					w.PrintNum{Text: "+7 905 331-09-42"},
					w.PrintNum{Text: "v.lebedeva@example.org"},
				},
				Columns: true,
			},
			w.PrintRubric{Text: "М"},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.PrintCell{Name: "Михаил Енученко", Detail: "nextcloud · Контакты"},
					w.PrintNum{Text: "—"},
					w.PrintNum{Text: "mix@example.org"},
				},
				Columns: true,
			},
			w.PrintFoot{Text: "Carrel · contacts · 7 cards · captured 2026-08-15 09:41 · 3 of 3 books"},
		},
		Variant: w.PagePrint,
	}
}

// Print3 is frame «Карточка контакта на бумаге» (print-3).
func Print3() Screen {
	return w.Page{
		Items: []w.PagePart{
			w.PrintTitle{Text: "Анна Ковалёва"},
			w.PrintSources{Text: "nextcloud · Контакты"},
			w.PrintRow{
				Items: []w.PrintRowPart{w.PrintCell{Name: "Organisation"}, w.PrintCell{Name: "Издательство «Слово»"}},
				Field: true,
			},
			w.PrintRow{
				Items: []w.PrintRowPart{w.PrintCell{Name: "Title"}, w.PrintCell{Name: "Выпускающий редактор"}},
				Field: true,
			},
			w.PrintRow{
				Items: []w.PrintRowPart{w.PrintCell{Name: "Phone · mobile"}, w.PrintCell{Name: "+7 912 445-10-88"}},
				Field: true,
			},
			w.PrintRow{
				Items: []w.PrintRowPart{w.PrintCell{Name: "Phone · work"}, w.PrintCell{Name: "+7 495 120-45-00"}},
				Field: true,
			},
			w.PrintRow{
				Items: []w.PrintRowPart{w.PrintCell{Name: "Email · work"}, w.PrintCell{Name: "a.kovaleva@example.org"}},
				Field: true,
			},
			w.PrintRow{
				Items: []w.PrintRowPart{w.PrintCell{Name: "Email · second"}, w.PrintCell{Name: "anna.kovaleva@slovo.example"}},
				Field: true,
			},
			w.PrintRow{
				Items: []w.PrintRowPart{w.PrintCell{Name: "Address · work"}, w.PrintCell{Name: "Москва, Тверская 12, стр. 3"}},
				Field: true,
			},
			w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "Birthday"}, w.PrintCell{Name: "1987-03-14"}}, Field: true},
			w.PrintRow{
				Items: []w.PrintRowPart{w.PrintCell{Name: "Tags"}, w.PrintCell{Name: "работа, издательство"}},
				Field: true,
			},
			w.PrintRow{
				Items: []w.PrintRowPart{w.PrintCell{Name: "Note"}, w.PrintCell{Name: "Пишет по вторникам, отвечает в тот же день."}},
				Field: true,
			},
			w.PrintRubric{Text: "Kept as the server sent them"},
			w.PrintRow{
				Items: []w.PrintRowPart{w.PrintCell{Name: "X-EVOLUTION-FILE-AS"}, w.PrintCell{Name: "Ковалёва, Анна"}},
				Field: true,
			},
			w.PrintRow{
				Items: []w.PrintRowPart{w.PrintCell{Name: "X-ABLABEL"}, w.PrintCell{Name: "издательство"}},
				Field: true,
			},
			w.PrintRow{
				Items: []w.PrintRowPart{w.PrintCell{Name: "X-CARREL-UNKNOWN-FLAG"}, w.PrintCell{Name: "1"}},
				Field: true,
			},
			w.PrintFoot{Text: "Carrel · contact · captured 2026-08-15 09:41"},
		},
		Variant: w.PagePrint,
	}
}

// Print4 is frame «Задачи на бумаге» (print-4).
func Print4() Screen {
	return w.Page{
		Items: []w.PagePart{
			w.PrintTitle{Text: "Tasks · Open · Combined list"},
			w.PrintSources{Text: "nextcloud · Дела, Работа"},
			w.PrintRow{
				Items: []w.PrintRowPart{w.Mono{Text: "Due", Variant: w.MonoFace}, w.PrintCell{Name: "Task"}},
				Head:  true,
			},
			w.PrintRubric{Text: "Overdue"},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "7 Aug", Variant: w.MonoFace},
					w.PrintCell{
						Name:   "Отправить документы в МФЦ",
						Detail: "nextcloud · Дела · overdue by 3 days",
						Tick:   w.PrintTickOpen,
					},
				},
			},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "8 Aug", Variant: w.MonoFace},
					w.PrintCell{Name: "Ответить подрядчику по смете", Detail: "nextcloud · Работа · 40%", Tick: w.PrintTickOpen},
				},
			},
			w.PrintRubric{Text: "This week"},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "12 Aug", Variant: w.MonoFace},
					w.PrintCell{Name: "Собрать чемодан", Detail: "18:00 · nextcloud · Дела, Работа", Tick: w.PrintTickOpen},
				},
			},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "14 Aug", Variant: w.MonoFace},
					w.PrintCell{Name: "Дописать раздел про кэш", Detail: "nextcloud · Работа · 70%", Tick: w.PrintTickOpen},
				},
			},
			w.PrintFoot{Text: "Carrel · tasks · Open · 4 of 7 · captured 2026-08-15 09:41 · 2 of 2 lists"},
		},
		Variant: w.PagePrint,
	}
}

// Print5 is frame «Заметка на бумаге» (print-5).
func Print5() Screen {
	return w.Page{
		Items: []w.PagePart{
			w.PrintTitle{Text: "Заметка со встречи 10.08"},
			w.PrintSources{Text: "nextcloud · Записи · 10 August 2026, 11:05"},
			w.PrintText{Text: "Договорились заморозить ветку в среду."},
			w.PrintText{
				Text: "Письмо пользователям пишет Анна, черновик к вечеру вторника. Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.",
			},
			w.PrintText{Text: "Что делаем", Heading: true},
			w.PrintText{Text: "проверить changelog", Tick: w.PrintTickOpen},
			w.PrintText{Text: "попросить у Артёма скриншоты", Tick: w.PrintTickDone},
			w.PrintText{
				Text: "Перенос старых вложений в новую папку — Виктор посмотрит, сколько их вообще, и скажет в четверг.",
			},
			w.PrintRubric{Text: "Attachments"},
			w.PrintRow{
				Items:   []w.PrintRowPart{w.PrintCell{Name: "Name"}, w.PrintNum{Text: "Size"}, w.PrintNum{Text: "Where"}},
				Columns: true,
				Head:    true,
			},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.PrintCell{Name: "2026-08-10-sync-board.png"},
					w.PrintNum{Text: "412 KB"},
					w.PrintNum{Text: "in the note"},
				},
				Columns: true,
			},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.PrintCell{Name: "запись-встречи.m4a"},
					w.PrintNum{Text: "24 MB"},
					w.PrintNum{Text: "Documents / carrel"},
				},
				Columns: true,
			},
			w.PrintRubric{Text: "Kept as the server sent them"},
			w.PrintRow{
				Items: []w.PrintRowPart{w.PrintCell{Name: "DESCRIPTION"}, w.PrintCell{Name: "Договорились заморозить ветку в четверг."}},
				Field: true,
			},
			w.PrintRow{
				Items: []w.PrintRowPart{w.PrintCell{Name: "X-JTX-COLOR"}, w.PrintCell{Name: "-16711936"}},
				Field: true,
			},
			w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "X-JTX-PINNED"}, w.PrintCell{Name: "1"}}, Field: true},
			w.PrintRow{Items: []w.PrintRowPart{w.PrintCell{Name: "X-OTHER"}, w.PrintCell{Name: "kept"}}, Field: true},
			w.PrintFoot{Text: "Carrel · note · captured 2026-08-15 09:41"},
		},
		Variant: w.PagePrint,
	}
}

// Print8 is frame «Список заметок на бумаге» (print-8).
func Print8() Screen {
	return w.Page{
		Items: []w.PagePart{
			w.PrintTitle{Text: "Notes · Combined notebook"},
			w.PrintSources{Text: "nextcloud · Записи, Дневник"},
			w.PrintRow{
				Items: []w.PrintRowPart{w.Mono{Text: "Date", Variant: w.MonoFace}, w.PrintCell{Name: "Note"}},
				Head:  true,
			},
			w.PrintRubric{Text: "August 2026"},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "10 Aug", Variant: w.MonoFace},
					w.PrintCell{
						Name:   "Заметка со встречи 10.08",
						Detail: "Договорились заморозить ветку в среду; письмо пользователям пишет Анна… · in 2 notebooks",
					},
				},
			},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "8 Aug", Variant: w.MonoFace},
					w.PrintCell{
						Name:   "Что почитать про CRDT",
						Detail: "Кляйн, глава 5; статья Шапиро 2011; демка automerge… · in 2 notebooks",
					},
				},
			},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "3 Aug", Variant: w.MonoFace},
					w.PrintCell{
						Name:   "Про дорогу до дачи",
						Detail: "Съезд после заправки, потом грунтовка два километра… · nextcloud · Дневник",
					},
				},
			},
			w.PrintRubric{Text: "July 2026"},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "21 Jul", Variant: w.MonoFace},
					w.PrintCell{
						Name:   "Про сроки по сборнику",
						Detail: "Анна просит две недели на вёрстку после сдачи текста. · nextcloud · Записи",
					},
				},
			},
			w.PrintFoot{Text: "Carrel · notes · 4 notes · captured 2026-08-15 09:41 · 2 of 2 notebooks"},
		},
		Variant: w.PagePrint,
	}
}

// Print6 is frame «Список файлов на бумаге» (print-6).
func Print6() Screen {
	return w.Page{
		Items: []w.PagePart{
			w.PrintTitle{Text: "Files · Documents / carrel / attachments / 2026-07"},
			w.PrintSources{Text: "sftpgo · vps · Documents"},
			w.PrintRow{
				Items:   []w.PrintRowPart{w.PrintCell{Name: "Name"}, w.PrintNum{Text: "Size"}, w.PrintNum{Text: "Changed"}},
				Columns: true,
				Head:    true,
			},
			w.PrintRubric{Text: "Folders"},
			w.PrintRow{
				Items:   []w.PrintRowPart{w.PrintCell{Name: "архив"}, w.PrintNum{Text: "—"}, w.PrintNum{Text: "2026-07-01 09:20"}},
				Columns: true,
			},
			w.PrintRow{
				Items:   []w.PrintRowPart{w.PrintCell{Name: "черновики"}, w.PrintNum{Text: "—"}, w.PrintNum{Text: "2026-07-29 16:40"}},
				Columns: true,
			},
			w.PrintRubric{Text: "Files"},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.PrintCell{Name: "whiteboard.jpg"},
					w.PrintNum{Text: "2.8 MB"},
					w.PrintNum{Text: "2026-07-14 17:05"},
				},
				Columns: true,
			},
			w.PrintRow{
				Items:   []w.PrintRowPart{w.PrintCell{Name: "панорама.tif"}, w.PrintNum{Text: "84 MB"}, w.PrintNum{Text: "2026-07-21 12:30"}},
				Columns: true,
			},
			w.PrintRow{
				Items:   []w.PrintRowPart{w.PrintCell{Name: "смета.pdf"}, w.PrintNum{Text: "1.2 MB"}, w.PrintNum{Text: "2026-07-08 10:15"}},
				Columns: true,
			},
			w.PrintRow{
				Items:   []w.PrintRowPart{w.PrintCell{Name: "схема-1.png"}, w.PrintNum{Text: "640 KB"}, w.PrintNum{Text: "2026-07-16 11:02"}},
				Columns: true,
			},
			w.PrintRow{
				Items:   []w.PrintRowPart{w.PrintCell{Name: "схема-2.png"}, w.PrintNum{Text: "712 KB"}, w.PrintNum{Text: "2026-07-16 11:04"}},
				Columns: true,
			},
			w.PrintFoot{Text: "Carrel · files · 7 items · 89.4 MB · captured 2026-08-15 09:41"},
		},
		Variant: w.PagePrint,
	}
}

// Print7 is frame «Агенда одного календаря на бумаге» (print-7).
func Print7() Screen {
	return w.Page{
		Items: []w.PagePart{
			w.PrintTitle{Text: "Личный · 10 – 16 August 2026"},
			w.PrintSources{Text: "nextcloud · Личный"},
			w.PrintRubric{Text: "Monday · 10 August"},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "10:00", Variant: w.MonoFace},
					w.PrintCell{Name: "Синк-ап по релизу", Detail: "Переговорная 2 · 4 attendees · repeats weekly"},
				},
			},
			w.PrintRubric{Text: "Tuesday · 11 August"},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "18:00", Variant: w.MonoFace},
					w.PrintCell{Name: "Забрать документы", Detail: "МФЦ, Мясницкая 13"},
				},
			},
			w.PrintRubric{Text: "Wednesday · 12 August"},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "all day", Variant: w.MonoFace},
					w.PrintCell{Name: "Поездка в Казань", Detail: "day 1 of 3"},
				},
			},
			w.PrintRubric{Text: "Thursday · 13 August"},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "all day", Variant: w.MonoFace},
					w.PrintCell{Name: "Поездка в Казань", Detail: "day 2 of 3"},
				},
			},
			w.PrintRubric{Text: "Friday · 14 August"},
			w.PrintRow{
				Items: []w.PrintRowPart{
					w.Mono{Text: "all day", Variant: w.MonoFace},
					w.PrintCell{Name: "Поездка в Казань", Detail: "day 3 of 3"},
				},
			},
			w.PrintFoot{Text: "Carrel · agenda · Личный · 2026-08-10 – 2026-08-16 · captured 2026-08-15 09:41"},
		},
		Variant: w.PagePrint,
	}
}
