// Раздел макета: 7.21 Остальные экраны и проверка полноты.

package screens

import w "carrel.dev/typecase/widget"

// Rest1 is frame «Ссылка-приглашение больше не работает» (rest-1).
func Rest1() Screen {
	return w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogLogo{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo48}, Wide: true},
						w.DialogTitle{Text: "Invitation", Gap: w.DialogTitleGap6},
						w.Prose{
							Parts: []w.Inline{
								w.Text("This invitation link is not valid. It may have expired, already been used, or been revoked."),
							},
							Variant: w.ProseMuted,
						},
						w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Sign in", Kind: w.ButtonPrimary}}, Variant: w.FormFootFlush},
					},
					Variant: w.DialogNarrow,
				},
			},
		},
	}
}

// Rest2 is frame «Адрес подтверждён» (rest-2).
func Rest2() Screen {
	return w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogLogo{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo48}, Wide: true},
						w.DialogTitle{Text: "Email address", Gap: w.DialogTitleGap6},
						w.Prose{Parts: []w.Inline{w.Text("Address confirmed. You can sign in now.")}, Variant: w.ProseMuted},
						w.FormFoot{Items: []w.FormFootPart{w.Button{Label: "Sign in", Kind: w.ButtonPrimary}}, Variant: w.FormFootFlush},
					},
					Variant: w.DialogNarrow,
				},
			},
		},
	}
}

// Rest3 is frame «О программе» (rest-3).
func Rest3() Screen {
	return w.Page{
		Items: []w.PagePart{
			w.Auth{
				Dialog: w.Dialog{
					Items: []w.DialogPart{
						w.DialogLogo{Logo: w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63", Size: w.Logo42}, Wide: true},
						w.DialogTitle{Text: "Carrel"},
						w.DialogSub{Parts: []w.Inline{w.Mono{Text: "0.0.1 · 786949a"}}},
						w.Tagline{Lead: "One desk. Every server.", Text: "Calendars, contacts, tasks, notes and files — one place."},
						w.Prose{
							Parts: []w.Inline{
								w.Text("AGPL-3.0-or-later: if you run a modified version for others, you must make the source available to them."),
							},
							Variant: w.ProseMuted,
						},
						w.FormFoot{
							Items: []w.FormFootPart{
								w.Button{Label: "Source code", Small: true},
								w.Button{Label: "Third-party licences", Small: true},
								w.Push{Items: []w.PushPart{w.Button{Label: "Back to sign in"}}},
							},
							Variant: w.FormFootFlush,
						},
					},
					Variant: w.DialogNarrow,
				},
			},
		},
	}
}
