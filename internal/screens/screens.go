package screens

// Screen is one assembled mockup page in the screens demo.
type Screen struct {
	ID    string
	Title string
	Body  func() string
}

// All returns every screen route in stable order.
func All() []Screen {
	return []Screen{
		{ID: "notes", Title: "7.4 Notes", Body: Notes},
		{ID: "notes-read", Title: "7.4 Notes · read", Body: NotesRead},
		{ID: "notes-edit", Title: "7.4 Notes · edit", Body: NotesEdit},
		{ID: "notes-quick", Title: "7.4 Notes · quick", Body: NotesQuick},
		{ID: "files", Title: "7.5 Files", Body: Files},
		{ID: "contacts", Title: "7.1 Contacts", Body: Contacts},
		{ID: "agenda", Title: "7.2 Calendar", Body: Agenda},
		{ID: "tasks", Title: "7.3 Tasks", Body: Tasks},
		{ID: "find", Title: "7.6 Find", Body: Find},
		{ID: "duplicates", Title: "7.7 Duplicates", Body: Duplicates},
		{ID: "conflict", Title: "7.8 Conflict", Body: Conflict},
		{ID: "io", Title: "7.9 Import and export", Body: ImportExport},
		{ID: "account", Title: "7.10 Account", Body: Account},
		{ID: "connections", Title: "7.11 Connections", Body: Connections},
		{ID: "attachments", Title: "7.12 Attachments", Body: Attachments},
		{ID: "devices", Title: "7.13 New device", Body: Devices},
		{ID: "davloom", Title: "7.14 davloom", Body: Davloom},
		{ID: "backup", Title: "7.15 Backup", Body: Backup},
		{ID: "published", Title: "7.16 Published links", Body: Published},
		{ID: "appearance", Title: "7.17 Appearance", Body: Appearance},
		{ID: "mail", Title: "7.18 Mail", Body: Mail},
		{ID: "admin", Title: "7.19 Administration", Body: Admin},
		{ID: "login", Title: "7.20 Sign in", Body: Login},
		{ID: "states", Title: "8. States", Body: States},
	}
}

// ByID finds one screen or nil.
func ByID(id string) *Screen {
	for _, s := range All() {
		if s.ID == id {
			cp := s
			return &cp
		}
	}
	return nil
}
