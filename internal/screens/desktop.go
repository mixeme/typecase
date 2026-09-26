// Раздел макета: 12. Настольное приложение.

package screens

import w "carrel.dev/typecase/widget"

// Desktop1 is frame «Окно приложения» (m-desktop-1).
func Desktop1() Screen {
	return w.Shell{
		Top: &w.Top{
			Brand:  w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}},
			Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}},
			Actions: []w.TopAction{
				w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary},
				w.Button{Label: "2 min ago", Icon: "refresh"},
				w.UserMenu{Name: "mix"},
			},
		},
		Body: w.ShellBody{
			Rail: []w.RailPart{
				w.Nav{
					Items: []w.NavItem{
						{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
						{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
						{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
						{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
						{Icon: "file", Label: "Files", Href: "#"},
					},
					On: 2,
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.TickRow{
							Ticked:  w.TickedPart,
							Kind:    w.TickRowAll,
							Mark:    w.Mark{Kind: w.MarkMulti},
							Label:   "Combined calendar",
							Meta:    "14",
							Current: true,
						},
					},
				},
				w.RailSec{
					Items: []w.RailSecPart{
						w.Rubric{Text: "nextcloud"},
						w.TickRow{Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "7"},
						w.TickRow{Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "4"},
					},
				},
				w.RailFoot{
					Items: []w.RailFootPart{
						w.Button{Icon: "plus", Title: "New calendar"},
						w.Button{Icon: "server", Title: "Connect account"},
						w.Button{Icon: "link", Title: "Subscribe to a feed"},
						w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"},
					},
					About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"},
				},
			},
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Mark: &w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, Text: "Combined calendar"},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("5 of 7 calendars · 10 – 16 August 2026")}}},
						Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "New event", Menu: true, Icon: "plus", Kind: w.ButtonPrimary}}},
					},
					w.ListGroup{Title: w.Rubric{Text: "Monday · 10 August", Form: w.RubricInline}, Count: &w.Num{Value: 3}},
					w.List{
						Rows: []w.ListRow{
							w.TimeRow{
								Mark:    w.Mark{Colour: "#4A6B52"},
								Time:    "10:00",
								Title:   []w.Inline{w.Text("Синк-ап по релизу")},
								Strong:  true,
								Detail:  "Переговорная 2",
								Caption: &w.Badge{Label: "4 attendees"},
								Source:  "nextcloud · Личный",
								On:      true,
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#8C6239"},
								Time:   "all day",
								Title:  []w.Inline{w.Text("Дежурство: первая линия")},
								Detail: "—",
								Source: "2 sources",
							},
							w.TimeRow{
								Mark:   w.Mark{Colour: "#3F6E8C"},
								Time:   "15:00",
								Title:  []w.Inline{w.Text("Созвон с подрядчиком")},
								Detail: "Jitsi",
								Source: "nextcloud · Работа",
							},
						},
					},
				},
			},
			Side: &w.Side{
				Items: []w.SidePart{
					w.Head{
						Title: &w.HeadTitle{Text: "Синк-ап по релизу", Size: w.HeadTitle17},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("nextcloud · Личный")}}},
						Acts:  &w.Acts{Items: []w.ActsPart{w.Button{Label: "Edit", Small: true}, w.Close{Label: "Close the panel"}}},
						Pad:   w.HeadPadBottom8,
					},
					w.Fields{
						Items: []w.FieldsItem{
							{Term: "When", Value: []w.ValuePart{w.Mono{Text: "10:00 – 11:00"}}},
							{Term: "Where", Value: []w.ValuePart{w.Text("Переговорная 2")}},
						},
					},
					w.Sec{
						Items: []w.SecPart{
							w.Rubric{Text: "Attachments"},
							w.SecOpen{
								Items: []w.SecOpenPart{
									w.SecOpened{Name: "release-checklist.md", Size: w.Dim{Parts: []w.Inline{w.Text("· 3 KB")}}},
									w.Line{
										Items: []w.LinePart{w.Button{Label: "Open in editor", Small: true}, w.Button{Label: "Show in folder", Small: true}},
										Gap:   w.LineGap6,
									},
									w.Dim{
										Parts: []w.Inline{
											w.Text("Edits go to the server when the program closes the file. The temporary copy lives until then and is deleted."),
										},
										Size: w.Dim11,
									},
								},
							},
						},
					},
				},
			},
			Height: w.ShellHeight390,
		},
	}
}

// Desktop13 is frame «Настройки приложения» (m-desktop-13).
func Desktop13() Screen {
	return w.Shell{
		Body: w.ShellBody{
			Main: w.Main{
				Items: []w.MainPart{
					w.Head{
						Title: &w.HeadTitle{Text: "Application", Size: w.HeadTitle19},
						Subs:  []w.HeadSub{w.HeadSub{Parts: []w.HeadSubPart{w.Text("Account and collection settings are inside, in Settings")}}},
						Pad:   w.HeadPadBottom6,
					},
					w.Form{
						Items: []w.FormPart{
							w.FieldSet{
								Rubric: w.Rubric{Text: "Connection"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Signed in as",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Input{Value: "mix · https://carrel.example", Mono: true, Variant: w.InputFill},
													w.Button{Label: "Sign out…"},
												},
												Gap: w.LineGap8,
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Signing out brings back the first screen, where the mode is chosen again — a server or a local run. Nothing is deleted by it. It asks first, which is what the ellipsis promises — in the web user menu the same words end the session at once and carry no ellipsis. In local mode the line reads"),
													w.Mono{Text: "mix · local run"},
													w.Text(": there is no address to show."),
												},
											},
										},
										Place: w.Grown,
									},
								},
								Variant: w.FieldSetLead,
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Behaviour"},
								Items: []w.FieldSetPart{
									w.Field{
										Parts: []w.FieldPart{
											w.Check{On: true, Parts: []w.CheckPart{w.Text("Start at sign-in")}},
											w.Check{On: true, Parts: []w.CheckPart{w.Text("Closing the window hides to the tray")}},
										},
										Gap:   8,
										Place: w.Grown,
									},
									w.Field{
										Label: "Global note shortcut",
										Parts: []w.FieldPart{
											w.Input{Value: "Ctrl+Shift+N", Label: "Global note shortcut", Mono: true},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("The note opens on the screen the pointer is on; the main window comes back where it was left."),
												},
											},
										},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Downloads folder",
										Parts: []w.FieldPart{w.Input{Value: "C:\\Users\\mix\\Downloads", Label: "Downloads folder", Mono: true}},
										Width: w.FieldMD,
									},
									w.Field{
										Label: "Open files with",
										Parts: []w.FieldPart{w.Pick{Value: "default application"}},
										Width: w.FieldMD,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Network"},
								Items: []w.FieldSetPart{
									w.Field{Label: "Proxy", Parts: []w.FieldPart{w.Pick{Value: "system"}}, Width: w.FieldMD},
									w.Field{
										Parts: []w.FieldPart{w.Check{Parts: []w.CheckPart{w.Text("Trust this instance’s self-signed certificate")}}},
									},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("This appears only after a certificate has been rejected, and it remembers that one fingerprint — it is not “turn checking off”."),
												},
											},
										},
										Place: w.Grown,
									},
								},
							},
							w.FieldSet{
								Rubric: w.Rubric{Text: "Updates"},
								Items: []w.FieldSetPart{
									w.Field{
										Label: "Installed",
										Parts: []w.FieldPart{
											w.Line{
												Items: []w.LinePart{
													w.Input{Value: "app 0.8.0 · server 0.8.0", Mono: true, Variant: w.Input11_5Fill},
													w.Button{Label: "Check now"},
												},
												Gap: w.LineGap8,
											},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Last checked 2 hours ago, by this button ·"),
													w.Mono{Text: "0.8.0"},
													w.Text("was the latest. Nothing is checked on its own while the tick below is clear."),
												},
											},
										},
										Width: w.FieldFill300,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.Check{Parts: []w.CheckPart{w.Text("Check for updates automatically")}},
											w.Check{Parts: []w.CheckPart{w.Text("Install when I close the app")}},
											w.Check{Parts: []w.CheckPart{w.Text("Include pre-releases")}},
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("Both are off out of the box, and off means the app does not reach the network at all — not “checks and stays quiet”. The button above works either way. With checking on, an update is downloaded from the project’s releases and verified by signature before anything is replaced; with the second tick off it then waits until the button is pressed, and nothing restarts under your hands mid-work."),
												},
											},
										},
										Gap:   7,
										Place: w.Grown,
									},
									w.Field{
										Parts: []w.FieldPart{
											w.FieldHint{
												Parts: []w.Inline{
													w.Text("The app and the server are versioned separately. In local mode both are updated here at once, because the server runs inside the app. Connected to your own server, this updates the app only — the server is updated where it runs, and the app says so when the two versions stop matching."),
												},
											},
										},
										Place: w.Grown,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
