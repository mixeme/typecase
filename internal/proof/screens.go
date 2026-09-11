package proof

// Screens transliterate mockup frames top to bottom. Only widgets and data:
// no class, no number.

func topBar() Top {
	return Top{
		Search: "Search everything…",
		Actions: Nodes{
			Btn{Label: "New note", Icon: "quill", Chev: true, Primary: true},
			Btn{Label: "2 min ago", Icon: "refresh"},
		},
		User: "mix",
	}
}

func nav(current string) []NavItem {
	items := []NavItem{
		{Icon: "con", Label: "Contacts", Href: "#contacts", Count: num(312)},
		{Icon: "cal", Label: "Calendar", Href: "#calendar", Count: num(14)},
		{Icon: "task", Label: "Tasks", Href: "#tasks", Count: num(3)},
		{Icon: "note", Label: "Notes", Href: "#notes", Count: num(86)},
		{Icon: "file", Label: "Files", Href: "#files"},
		{Icon: "dup", Label: "Duplicates", Href: "#duplicates", Count: num(3)},
	}
	for i := range items {
		items[i].On = items[i].Label == current
	}
	return items
}

// Notes is frame «Заметки · список и редактор».
func Notes() Screen {
	return Screen{
		Top:  topBar(),
		Size: SideWide,
		Rail: Rail{
			Nav: nav("Notes"),
			Sections: []Section{
				{Rows: []Source{{Label: "Combined notebook", Mark: Mark{Multi: true}, Meta: "86", Box: Box{Part: true}, All: true, On: true}}},
				{Rubric: "nextcloud · home", Rows: []Source{
					{Label: "Записи", Mark: Mark{Colour: "#4A6B52"}, Meta: "64"},
					{Label: "Дневник", Mark: Mark{Colour: "#6B4E7A"}, Meta: "22"},
					{Label: "Архив", Mark: Mark{Colour: "#A05A2C"}, Meta: "—", Box: Box{Off: true}, Off: true},
				}},
				{Rubric: "Tags", Tags: []Tag{{Label: "all", On: true}, {Label: "встречи"}, {Label: "идеи"}, {Label: "книги"}, {Label: "рецепты"}}},
			},
			Foot: []Btn{
				{Icon: "plus", Title: "New notebook"},
				{Icon: "server", Title: "Connect account"},
				{Icon: "link", Title: "Subscribe to a feed"},
			},
			About: "Carrel", Build: "0.0.1 · 786949a",
		},
		Main: Nodes{
			Head{
				Mark:  &Mark{Multi: true},
				Title: "Combined notebook",
				Sub:   Nodes{Text("2 of 3 notebooks · 86 notes · VJOURNAL")},
				Acts: Nodes{
					Btn{Label: "New note", Icon: "plus", Chev: true, Primary: true},
					Btn{Label: "Import Markdown"},
					Btn{Label: "Export"},
				},
			},
			Bar{Items: Nodes{
				Filter{Placeholder: "Filter by word", Width: Filter200},
				Seg{Label: "Sort", Items: []SegItem{{Label: "Newest", On: true}, {Label: "Oldest"}, {Label: "Title"}}},
				Btn{Label: "Columns", Icon: "cols"},
				Right{Is2nd: true, Items: Nodes{Inline{Mono, "86"}}},
			}},
			Group{Rubric: "August 2026", Count: 9},
			List{Rows: Nodes{
				NoteRow{Mark: Mark{Colour: "#4A6B52"}, Title: "Заметка со встречи 10.08", On: true,
					Excerpt: "Договорились заморозить ветку в среду; письмо пользователям пишет Анна…",
					Tags:    []Tag{{Label: "встречи"}}, Date: "10 Aug", Source: "nextcloud · Записи"},
				NoteRow{Mark: Mark{Colour: "#4A6B52"}, Title: "Что почитать про CRDT",
					Excerpt: "Кляйн, глава 5; статья Шапиро 2011; демка automerge…",
					Tags:    []Tag{{Label: "книги"}}, Date: "8 Aug", Source: "nextcloud · Записи"},
				NoteRow{Mark: Mark{Colour: "#6B4E7A"}, Title: "Про дорогу до дачи",
					Excerpt: "Съезд после заправки, потом грунтовка два километра…",
					Badge:   &Badge{Label: "linked", Linked: true}, Date: "3 Aug", Source: "nextcloud · Дневник"},
			}},
		},
		Side: &Side{
			Head: Head{
				Title: "Заметка со встречи 10.08",
				Size:  TitlePanel,
				Sub:   Nodes{Text("nextcloud · Записи · "), Inline{Mono, "10 Aug 2026 11:05"}},
				Acts: Nodes{
					Btn{Label: "Save", Small: true, Primary: true},
					Btn{Icon: "expand", Small: true, Title: "Open full screen"},
					Btn{Icon: "close", Close: true, Title: "Close the panel"},
					Btn{Icon: "dots", Small: true, Title: "More"},
				},
			},
			Body: Nodes{
				Form{Tight: true, Fields: []Field{
					{Label: "Title", Width: WFull, In: In{ID: "n-title", Kind: InText, Value: "Заметка со встречи 10.08"}},
					{Label: "Text", Width: WFull, In: In{ID: "n-text", Kind: InArea, Tall: true, Value: noteText}},
					{Label: "Date", Width: WSM, In: In{ID: "n-date", Kind: InDate, Mono: true, Value: "2026-08-10"}},
					{Label: "Time", Width: WXS, In: In{ID: "n-time", Kind: InTime, Mono: true, Value: "11:05"}},
					{Label: "Tags", Width: WMD, In: In{ID: "n-tags", Kind: InText, Value: "встречи, релиз"}},
				}},
				Sec{Rubric: "Attachments", Action: &Btn{Label: "Attach", Small: true}, Body: Nodes{
					Attachments{
						{Name: "2026-08-10-sync-board.png", Meta: "image/png · 412 KB", Where: "in the note"},
						{Name: "запись-встречи.m4a", Meta: "audio/mp4 · 24 MB", Where: "Documents / carrel"},
					},
					Hint{Spaced: true, Text: "You can also paste a picture from the clipboard or drop one onto this card. Over 2 MB the file goes to the attachments folder and the note keeps a link."},
				}},
				Sec{Rubric: "Related", Body: Nodes{RelatedList{{Title: "Синк-ап по релизу", Meta: "· event · 10 Aug"}}}},
			},
		},
	}
}

const noteText = `Договорились заморозить ветку в среду.

Письмо пользователям пишет Анна, черновик к вечеру вторника.
Виктор проверяет миграцию на копии базы — если не успевает, переносим на следующий спринт.

- [ ] проверить changelog
- [ ] попросить у Артёма скриншоты`

// Files is frame «Файлы · обзор папки, свойства выбранного».
func Files() Screen {
	acts := func(download bool) Td {
		b := Nodes{}
		if download {
			b = append(b, Btn{Icon: "down", Small: true, Title: "Download"})
		}
		return Td{Acts: true, Body: append(b, Btn{Icon: "dots", Small: true, Title: "More"})}
	}
	file := func(icon, name, size, typ, changed string, sel bool) Tr {
		return Tr{Sel: sel, Cells: []Td{
			{Body: Nodes{Box{Off: !sel, Label: "Select " + name}}},
			{Body: Nodes{FName{Icon: icon, Name: name}}},
			{Mono: true, Body: Nodes{Text(size)}},
			{Mono: true, Body: Nodes{Text(typ)}},
			{Mono: true, Body: Nodes{Text(changed)}},
			acts(true),
		}}
	}
	return Screen{
		Top:  topBar(),
		Size: SideNarrow,
		Rail: Rail{
			Nav: nav("Files"),
			Sections: []Section{
				{Rows: []Source{{Label: "Selected collections", Root: "file", Meta: "2 of 3", Box: Box{Part: true}}}},
				{Rubric: "sftpgo · vps", Rows: []Source{
					{Label: "Documents", Mark: Mark{Colour: "#4A6B52"}, Meta: "rw", On: true},
					{Label: "Public", Mark: Mark{Colour: "#9A9280"}, Badge: "ro", Box: Box{Off: true}, Off: true},
				}},
				{Rubric: "nextcloud · home", Rows: []Source{
					{Label: "Files", Mark: Mark{Colour: "#3F6E8C"}, Meta: "rw"},
				}},
				{Rubric: "Shortcuts", Tree: []Shortcut{
					{Icon: "file", Label: "Attachments", Meta: "Documents / carrel", On: true},
					{Icon: "link", Label: "Published", Meta: "2"},
				}},
			},
			Foot:  []Btn{{Icon: "plus", Title: "Connect account"}},
			About: "Carrel", Build: "0.0.1 · 786949a",
		},
		Main: Nodes{
			Head{
				Title:  "attachments",
				Crumbs: []Crumb{{Label: "Documents", Href: "#documents"}, {Label: "carrel", Href: "#carrel"}, {Label: "attachments"}},
				Sub:    Nodes{Text("sftpgo · vps · "), Badge{Label: "attachment folder"}},
				Acts: Nodes{
					Btn{Label: "Upload", Icon: "up", Chev: true, Primary: true},
					Btn{Label: "New folder", Icon: "plus"},
					Btn{Icon: "refresh", Title: "Refresh"},
					Btn{Icon: "panel", Title: "Details", Pressed: true},
					Btn{Icon: "dots", Title: "More"},
				},
			},
			Bar{Sel: true, Items: Nodes{
				Inline{Strong, "1 selected"}, Inline{MonoMuted, "412 KB"},
				Sep{},
				Btn{Label: "Download", Icon: "down", Title: "Download", Lbl: true},
				Btn{Label: "Rename", Icon: "pencil", Title: "Rename", Lbl: true},
				Btn{Label: "Move", Icon: "move", Title: "Move", Lbl: true},
				Btn{Label: "Copy", Icon: "dup", Title: "Copy", Lbl: true},
				Btn{Label: "Delete", Icon: "trash", Title: "Delete", Lbl: true, DangerSolid: true},
				Right{Items: Nodes{Btn{Label: "Clear selection", Small: true}}},
			}},
			Bar{Items: Nodes{
				Filter{Placeholder: "Filter this folder", Width: Filter190, Icon: true},
				Seg{Label: "Sort", Is2nd: true, Items: []SegItem{{Label: "Name", On: true}, {Label: "Size"}, {Label: "Changed"}}},
				Seg{Label: "View", Items: []SegItem{{Label: "List", Icon: "menu", On: true}, {Label: "Tiles", Icon: "grid"}}},
				Btn{Label: "Columns", Icon: "cols"},
				Btn{Icon: "dots", More: true, Title: "More"},
				Right{Is2nd: true, Items: Nodes{Inline{Mono, "14 items · 38.2 MB"}}},
			}},
			Table{
				Cols: []Col{{Kind: ColTick, Tick: true}, {Label: "Name"}, {Label: "Size", Kind: ColSize}, {Label: "Type", Kind: ColType}, {Label: "Changed", Kind: ColTime}, {Kind: ColActs}},
				Rows: []Tr{
					{Cells: []Td{{}, {Span: 5, Body: Nodes{FName{Icon: "back", Name: "up one folder", Up: true}}}}},
					{Folder: true, Cells: []Td{
						{Body: Nodes{Box{Off: true, Label: "Select 2026-07"}}},
						{Body: Nodes{FName{Icon: "file", Name: "2026-07"}}},
						{Mono: true, Body: Nodes{Text("—")}},
						{Mono: true, Body: Nodes{Text("folder · 34 items")}},
						{Mono: true, Body: Nodes{Text("2026-07-01 09:12")}},
						acts(false),
					}},
					file("img", "2026-08-10-sync-board.png", "412 KB", "image/png", "2026-08-10 11:07", true),
					file("doc", "release-checklist.md", "3 KB", "text/markdown", "2026-08-09 18:44", false),
					file("doc", "смета-подрядчик.pdf", "1.2 MB", "application/pdf", "2026-08-05 14:20", false),
					file("img", "2026-08-04-whiteboard.jpg", "2.8 MB", "image/jpeg", "2026-08-04 17:03", false),
				},
			},
			List{Rows: Nodes{
				FileRow{Icon: "img", Name: "2026-08-10-sync-board.png", Size: "412 KB", Time: "2026-08-10 11:07", Mode: "rw", Source: "sftpgo · vps", Sel: true},
				FileRow{Icon: "doc", Name: "release-checklist.md", Size: "3 KB", Time: "2026-08-09 18:44", Mode: "rw", Source: "sftpgo · vps"},
			}},
			Msg{Text: "Drop files onto the window, or onto a folder in the list. A name that is already taken is refused rather than overwritten, and the queue says so."},
		},
		Side: &Side{
			Head: Head{
				Title: "2026-08-10-sync-board.png",
				Size:  TitleItem,
				Tight: true,
				Sub:   Nodes{Text("sftpgo · vps · Documents")},
				Acts:  Nodes{Btn{Icon: "close", Close: true, Title: "Close the panel"}},
			},
			Body: Nodes{
				Preview{Icon: "img"},
				PanelActions{
					Btn{Label: "Download", Icon: "down", Primary: true},
					Btn{Label: "Open"},
					Btn{Label: "Rename", Icon: "pencil"},
					Btn{Label: "Move", Icon: "move"},
				},
				Fields{
					{Label: "Size", Value: "412 KB", Mono: true},
					{Label: "Type", Value: "image/png", Mono: true},
					{Label: "Changed", Value: "2026-08-10 11:07", Mono: true},
					{Label: "Path", Value: "/dav/Documents/carrel/attachments/2026-08-10-sync-board.png", Mono: true, Path: true},
				},
				Sec{Rubric: "Links", Body: Nodes{Note("Carrel keeps no index of what points at this file. It may be attached to a note or an event — check before deleting.")}},
				DangerZone{Btn{Label: "Delete", Icon: "trash", DangerSolid: true}},
			},
		},
	}
}
