package screens

import "carrel.dev/typecase/internal/legacy/widget"

// Mail assembles mockup 7.18: letter sheet over a contact list.
func Mail() string {
	n3, n2 := 3, 2
	sheet := widget.Sheet{
		ID:     "letter",
		Title:  "Letter",
		Target: "from mix@carrel.example",
		Open:   true,
		Body: []widget.Node{
			widget.Pick{
				Name:   "to",
				Label:  "To",
				Chosen: []string{"Анна Ковалёва · a.kovaleva@example.org", "Артём Белов · a.belov@example.org"},
				Open:   false,
			},
			widget.Paragraph("2 of 10 recipients. Everyone sees the whole list."),
			widget.Field{Name: "subject", Label: "Subject", Kind: widget.InputText, Value: "Контакты подрядчика и смета", Width: widget.FieldWidthFull},
			widget.Field{
				Name:  "body",
				Label: "Letter",
				Kind:  widget.InputTextarea,
				Value: "Анна, Артём — вот карточка подрядчика и смета, которую он прислал. Договор в папке, ссылка ниже.",
				Width: widget.FieldWidthFull,
			},
			widget.Sec{
				Rubric: widget.Rubric{Text: "Attached"},
				Body: []widget.Node{widget.List{Groups: []widget.ListGroup{{Rows: []widget.Row{
					{Variant: widget.RowMeta, Cells: widget.RowCells{Icon: "tc-i-doc", Name: "смета-подрядчик.pdf", Size: "1.2 MB"}},
				}}}}},
			},
			widget.Msg{Text: "Over 2 MB the letter keeps a link, not the file."},
		},
		Actions: []widget.Action{
			{Label: "Send", Kind: widget.ButtonPrimary},
			{Label: "Cancel"},
		},
	}
	return shellFrame(shellNav("/contacts"), contactsRail(), widget.Main{
		Label: "Contacts",
		Body: []widget.Node{
			widget.Head{Title: "Контакты", Sub: "312 записей · 2 отмечено"},
			widget.List{Groups: []widget.ListGroup{{
				Rubric: widget.Rubric{Text: "А", Count: intPtr(n3)},
				Rows: []widget.Row{
					{
						Variant:  widget.RowAvatarCols,
						Selected: true,
						Mark:     widget.Mark{Colour: "#4A6B52", Label: "Контакты"},
						Cells: widget.RowCells{
							Avatar:   widget.Avatar{Initials: "АК", Colour: "#4A6B52"},
							Title:    "Анна Ковалёва",
							Subtitle: "+7 912 445-10-88",
							Detail:   "a.kovaleva@example.org",
							Caption:  "Издательство «Слово»",
							Source:   "nextcloud · Контакты",
						},
					},
					{
						Variant:  widget.RowAvatarCols,
						Selected: true,
						Mark:     widget.Mark{Colour: "#6B4E7A", Label: "Команда"},
						Cells: widget.RowCells{
							Avatar:   widget.Avatar{Image: "/img/sample.svg", Initials: "АБ"},
							Title:    "Артём Белов",
							Subtitle: "+7 903 118-22-04",
							Detail:   "a.belov@example.org",
							Caption:  "Тензор",
							Source:   "radicale · Команда",
						},
					},
					{
						Variant: widget.RowAvatarCols,
						Mark:    widget.Mark{Colour: "#4A6B52", Label: "Контакты"},
						Cells: widget.RowCells{
							Avatar:   widget.Avatar{Initials: "АГ", Colour: "#4A6B52"},
							Title:    "Алексей Гринёв",
							Subtitle: "+7 921 330-07-19",
							Detail:   "a.grinev@example.org",
							Caption:  "—",
							Source:   "nextcloud · Контакты",
						},
					},
				},
			}}},
			widget.List{Groups: []widget.ListGroup{{
				Rubric: widget.Rubric{Text: "Б", Count: intPtr(n2)},
				Rows: []widget.Row{
					{
						Variant: widget.RowAvatarCols,
						Mark:    widget.Mark{Colour: "#4A6B52", Label: "Контакты"},
						Cells: widget.RowCells{
							Avatar:   widget.Avatar{Initials: "БК", Colour: "#4A6B52"},
							Title:    "Борис Кузьмин",
							Subtitle: "+7 905 771-42-08",
							Detail:   "b.kuzmin@example.org",
							Caption:  "—",
							Source:   "nextcloud · Контакты",
						},
					},
				},
			}}},
		},
	}, widget.Side{}, sheet.Render())
}
