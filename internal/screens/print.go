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
