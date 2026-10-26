// Раздел макета: 4. Каркас и сетка.

package screens

import w "carrel.dev/typecase/widget"

// Shell1 is frame «Каркас: четыре зоны» (shell-1).
func Shell1() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Placeholder: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Zone{
					Title: w.Rubric{Text: "Rail"},
					Note:  w.Mono{Text: "264 px · sections + sources", Form: w.MonoLine, Variant: w.Mono11Top},
				},
				w.RailFoot{
					Items: []w.RailFootPart{w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			RailForm: w.RailZone,
			Main: w.Main{
				Items: []w.MainPart{
					w.Notices{Items: []w.NoticesPart{}},
					w.Runbar{Closed: true},
					w.Zone{
						Title: w.Rubric{Text: "Content"},
						Note:  w.Mono{Text: "fluid · min 0", Form: w.MonoLine, Variant: w.Mono11Top},
					},
				},
				Variant: w.MainZone,
			},
			Side: &w.Side{
				Items: []w.SidePart{
					w.Zone{
						Title: w.Rubric{Text: "Details"},
						Note:  w.Mono{Text: "400 px · closes by ✕", Form: w.MonoLine, Variant: w.Mono11Top},
					},
				},
				Variant: w.SideZone,
			},
			Height: w.ShellHeight230,
		},
		Sheet:   &w.Scrim{Closed: true},
		Dialogs: []w.Scrim{w.Scrim{ID: "keys", Closed: true}},
	}
}
