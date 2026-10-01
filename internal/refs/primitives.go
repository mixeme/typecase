package refs

import (
	"carrel.dev/typecase/internal/fragment"
	w "carrel.dev/typecase/widget"
)

func ref(frame, class string, n int) fragment.Ref {
	return fragment.Ref{Frame: frame, Class: class, N: n}
}

// e is one record of a family.
func e(family, state string, fixture Renderer, r fragment.Ref) Entry {
	return Entry{Family: family, State: state, Fixture: fixture, Ref: r}
}

func primitive(state string, fixture Renderer, frame, class string, n int) Entry {
	return e("primitives", state, fixture, ref(frame, class, n))
}

var _ = add(
	// tc-btn — every form of the button.
	primitive("link-icon", w.Button{Label: "Sources", Icon: "cal", Href: "#sources-demo"}, "components-41", "tc-btn", 1),
	primitive("post", w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Form: "post-disconnect"}, "components-37", "tc-btn", 5),
	primitive("link-off", w.Button{Label: "Export", Title: "Nothing to export: the list is empty.", State: w.ButtonDisabled, Href: "#"}, "components-37", "tc-btn", 4),
	primitive("link-primary", w.Button{Label: "New contact", Icon: "plus", Kind: w.ButtonPrimary, Href: "/app/contacts/nextcloud/kontakty/new"}, "components-37", "tc-btn", 3),
	primitive("link", w.Button{Label: "Import", Href: "/app/contacts/nextcloud/kontakty/import"}, "components-37", "tc-btn", 2),
	primitive("disabled", w.Button{Label: "Rewrite 0 records", Kind: w.ButtonPrimary, State: w.ButtonDisabled}, "s-contacts-20", "tc-btn", 1),
	primitive("disabled", w.Button{Icon: "plus", Title: "Open a folder to make one in it.", State: w.ButtonDisabled}, "s-files-6", "tc-btn", 3),
	primitive("disabled", w.Button{Label: "New folder", Icon: "plus", State: w.ButtonDisabled}, "s-files-15", "tc-btn", 2),
	primitive("disabled", w.Button{Label: "Upload", Menu: true, Icon: "up", Kind: w.ButtonPrimary, State: w.ButtonDisabled}, "s-files-15", "tc-btn", 1),
	primitive("disabled", w.Button{Label: "Back to defaults", Small: true, Title: "Every weight is the default already.", State: w.ButtonDisabled}, "s-dup-7", "tc-btn", 3),
	primitive("icon", w.Button{Label: "2 min ago", Icon: "refresh"}, "shell-1", "tc-btn", 2),
	primitive("small", w.Button{Label: "Retry", Small: true}, "components-11", "tc-btn", 1),
	primitive("normal", w.Button{Label: "Import"}, "components-9", "tc-btn", 2),
	primitive("menu", w.Button{Label: "New note", Icon: "quill", Menu: true, Kind: w.ButtonPrimary}, "shell-1", "tc-btn", 1),
	primitive("icon", w.Button{Icon: "dots", Title: "More"}, "components-16", "tc-btn", 6),
	primitive("primary", w.Button{Label: "Save", Kind: w.ButtonPrimary}, "components-9", "tc-btn", 1),
	primitive("primary", w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary}, "components-9", "tc-btn", 7),
	primitive("variant", w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}, "s-profile-1", "tc-btn", 10),
	primitive("folds", w.Button{Label: "Download", Icon: "down", Folds: true, Title: "Download"}, "components-16", "tc-btn", 1),
	primitive("small", w.Button{Icon: "dots", Small: true, Title: "More"}, "s-contacts-1", "tc-btn", 15),
	primitive("danger-solid", w.Button{Label: "Delete contact", Kind: w.ButtonDangerSolid}, "components-9", "tc-btn", 5),
	primitive("danger", w.Button{Label: "Disconnect", Kind: w.ButtonDanger}, "components-9", "tc-btn", 6),
	primitive("small", w.Button{Label: "Add phone", Icon: "plus", Small: true}, "s-contacts-3", "tc-btn", 10),
	primitive("danger", w.Button{Label: "End", Small: true, Kind: w.ButtonDanger}, "account-1", "tc-btn", 6),
	primitive("focus", w.Button{Label: "New note", Icon: "quill", Menu: true, Kind: w.ButtonPrimary, Focus: true}, "s-notes-5", "tc-btn", 1),
	primitive("folds", w.Button{Label: "Delete", Icon: "trash", Folds: true, Kind: w.ButtonDangerSolid, Title: "Delete"}, "components-16", "tc-btn", 3),
	primitive("danger-solid", w.Button{Label: "Delete contact", Icon: "trash", Kind: w.ButtonDangerSolid}, "s-contacts-6", "tc-btn", 9),
	primitive("menu", w.Button{Label: "Attach", Menu: true, Small: true}, "s-agenda-1", "tc-btn", 17),
	primitive("variant", w.Button{Icon: "menu", IconBox: w.IconBox19, Variant: w.ButtonPadX4, Title: "Sections"}, "phone-1", "tc-btn", 1),
	primitive("with-content", w.Button{Label: "B", Bold: true}, "components-29", "tc-btn", 3),
	primitive("variant", w.Button{Icon: "quill", IconBox: w.IconBox16, Variant: w.ButtonPadX8, Kind: w.ButtonPrimary, Title: "New note"}, "phone-1", "tc-btn", 3),
	primitive("variant", w.Button{Icon: "dots", IconBox: w.IconBox16, Variant: w.ButtonPadX4, Title: "More"}, "phone-6", "tc-btn", 7),
	primitive("menu", w.Button{Label: "Attach", Icon: "up", Menu: true}, "s-agenda-3", "tc-btn", 6),
	primitive("variant", w.Button{Label: "I", Variant: w.ButtonItalic}, "s-notes-20", "tc-btn", 6),
	primitive("folds", w.Button{Label: "Attach", Icon: "up", Folds: true, Menu: true, Title: "Attach"}, "s-notes-20", "tc-btn", 11),
	primitive("off", w.Button{Label: "Rename", Icon: "pencil", Folds: true, State: w.ButtonOff, Title: "Rename"}, "s-files-2", "tc-btn", 9),
	primitive("focus", w.Button{Label: "Weights", Small: true, Focus: true}, "s-dup-3", "tc-btn", 1),
	primitive("focus", w.Button{Icon: "plus", Focus: true, Title: "New calendar"}, "s-profile-5", "tc-btn", 3),
	primitive("variant", w.Button{Icon: "search", IconBox: w.IconBox17, Variant: w.ButtonPadX6, Title: "Search"}, "phone-1", "tc-btn", 2),
	primitive("primary", w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "Where to write"}, "components-9", "tc-btn", 8),
	primitive("with-content", w.Button{Label: "Columns", Icon: "cols", Variant: w.ButtonOnCard, Aside: &w.Mono{Text: "5", Variant: w.MonoDusty11}}, "components-21", "tc-btn", 1),
	primitive("with-content", w.Button{Label: "Mark done", Tick: &w.Tick{Glyph: true}}, "s-tasks-2", "tc-btn", 4),
	primitive("variant", w.Button{Label: "Focus", Icon: "focus", Folds: true, Variant: w.ButtonOnCard, Title: "Focus"}, "s-notes-4", "tc-btn", 13),
	primitive("variant", w.Button{Icon: "panel", Variant: w.ButtonOnCard, Title: "Details"}, "s-files-1", "tc-btn", 9),
	primitive("disabled", w.Button{Label: "Merge on the server", Kind: w.ButtonDanger, State: w.ButtonDisabled, Reason: "Общий is read-only"}, "s-dup-1", "tc-btn", 18),
	primitive("primary", w.Button{Label: "Save", Small: true, Kind: w.ButtonPrimary}, "s-dup-3", "tc-btn", 2),
	primitive("with-content", w.Button{Label: "Download", Icon: "down", Aside: &w.Mono{Text: ".ics", Variant: w.Mono11}, After: "first"}, "s-profile-9", "tc-btn", 9),
	primitive("with-content", w.Button{Label: "Save as", Icon: "down", Aside: &w.Mono{Text: "carrel-personal.sh", Variant: w.Mono11}}, "s-backup-3", "tc-btn", 7),
	primitive("tertiary", w.Button{Label: "H2", Tertiary: true}, "s-notes-20", "tc-btn", 7),
	primitive("tertiary", w.Button{Folds: true, Label: "task", Icon: "task", Title: "Task", Tertiary: true}, "s-notes-20", "tc-btn", 9),
	primitive("tertiary", w.Button{Icon: "link", Title: "Link", Tertiary: true}, "s-notes-20", "tc-btn", 10),
	primitive("variant", w.Button{Label: "Copy", Small: true, Variant: w.ButtonSelfStart}, "s-backup-4", "tc-btn", 5),
	primitive("disabled", w.Button{Label: "Send test", State: w.ButtonDisabled, Reason: "Fill in the address first"}, "s-admin-3", "tc-btn", 4),
	primitive("variant", w.Button{Label: "Write a note", Variant: w.ButtonCentredFill}, "phone-3", "tc-btn", 3),
	primitive("variant", w.Button{Label: "Delete", Kind: w.ButtonDangerSolid, Variant: w.ButtonCentred}, "phone-3", "tc-btn", 4),
	primitive("danger-solid", w.Button{Icon: "trash", Kind: w.ButtonDangerSolid, Title: "Delete"}, "phone-6", "tc-btn", 5),

	// tc-close
	primitive("link", w.Close{Label: "Close", Href: "#"}, "components-41", "tc-close", 1),
	primitive("default", w.Close{Label: "Close the panel"}, "components-33", "tc-close", 1),

	// tc-mark
	primitive("one", w.Mark{Colour: "#4A6B52"}, "components-11", "tc-mark", 2),
	primitive("multi", w.Mark{Kind: w.MarkMulti}, "components-11", "tc-mark", 1),
	primitive("multi", w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, "s-contacts-1", "tc-mark", 6),
	primitive("sized", w.Mark{Colour: "#6B4E7A", Size: w.MarkH11Inline}, "s-contacts-6", "tc-mark", 3),
	primitive("external", w.Mark{Colour: "#8C6239", Kind: w.MarkExternal}, "s-profile-1", "tc-mark", 9),
	primitive("sized", w.Mark{Colour: "#4A6B52", Size: w.MarkH20}, "s-tasks-3", "tc-mark", 3),
	primitive("sized", w.Mark{Colour: "#4A6B52", Size: w.MarkH18}, "s-profile-10", "tc-mark", 13),
	primitive("sized", w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, "s-mail-1", "tc-mark", 16),
	primitive("sized", w.Mark{Colour: "#4A6B52", Size: w.MarkH32}, "phone-1", "tc-mark", 1),
	primitive("sized", w.Mark{Colour: "#4A6B52", Size: w.MarkH34}, "phone-4", "tc-mark", 1),
	primitive("sized", w.Mark{Colour: "#4A6B52", Size: w.MarkH19}, "s-agenda-1", "tc-mark", 17),
	primitive("sized", w.Mark{Size: w.MarkOnHair}, "states-3", "tc-mark", 6),
	primitive("sized", w.Mark{Colour: "#4A6B52", Size: w.MarkH12}, "s-notes-2", "tc-mark", 3),
	primitive("sized", w.Mark{Colour: "#4A6B52", Size: w.MarkH14}, "desktop-8", "tc-mark", 1),
	primitive("focus", w.Mark{Colour: "#8C6239", Size: w.MarkH18, Focus: true}, "s-profile-10", "tc-mark", 12),

	// tc-tag
	primitive("plain", w.Tag{Label: "работа"}, "s-contacts-1", "tc-tag", 2),
	primitive("on", w.Tag{Label: "all", On: true}, "s-contacts-1", "tc-tag", 1),
	primitive("icon", w.Tag{Label: "smirnov.vcf", Icon: "con"}, "s-mail-7", "tc-tag", 1),
	primitive("removable", w.Tag{Label: "Забрать документы", Removable: true}, "s-tasks-3", "tc-tag", 1),
	primitive("mark", w.Tag{Label: "Анна Ковалёва · a.kovaleva@example.org", Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11}}, "s-mail-1", "tc-tag", 2),

	// tc-badge
	primitive("neutral", w.Badge{Label: "external"}, "components-11", "tc-badge", 1),
	primitive("linked", w.Badge{Label: "linked ×2", Tone: w.BadgeLinked}, "components-12", "tc-badge", 1),
	primitive("alert", w.Badge{Label: "tentative", Tone: w.BadgeAlert}, "s-agenda-1", "tc-badge", 5),
	primitive("accent", w.Badge{Label: "credentials unlocked · 4:12", Accent: true}, "s-profile-4", "tc-badge", 1),
	primitive("local", w.Badge{Label: "local", Tone: w.BadgeLocal, Title: "Local instance: may be unreachable from outside"}, "s-plan-1", "tc-badge", 2),
	primitive("icon", w.Badge{Label: "linked", Tone: w.BadgeLinked, Icon: "link"}, "s-notes-1", "tc-badge", 2),

	// tc-tick
	primitive("done", w.Tick{Done: true, Label: "Done"}, "s-tasks-1", "tc-tick", 5),
	primitive("open", w.Tick{Label: "Done"}, "s-tasks-1", "tc-tick", 1),
	primitive("variant", w.Tick{Variant: w.TickNudge, Label: "Done"}, "s-notes-1", "tc-tick", 1),
	primitive("done", w.Tick{Done: true, Variant: w.TickNudge, Label: "Done"}, "s-notes-2", "tc-tick", 2),
	primitive("variant", w.Tick{Variant: w.TickEdgeDanger, Label: "Check"}, "s-admin-4", "tc-tick", 6),
	primitive("glyph", w.Tick{Glyph: true}, "s-tasks-2", "tc-tick", 1),
	primitive("variant", w.Tick{Variant: w.TickEdgeDusty, Label: "Check"}, "s-admin-4", "tc-tick", 8),

	// tc-kbd, tc-num, tc-plain
	primitive("default", w.Kbd{Keys: "Ctrl+S"}, "components-29", "tc-kbd", 1),
	primitive("normal", w.Num{Value: 3}, "components-12", "tc-num", 1),
	primitive("alert", w.Num{Value: 3, Alert: true}, "s-tasks-1", "tc-num", 6),
	primitive("default", w.Plain{Text: "2 active"}, "account-1", "tc-plain", 1),

	// tc-avatar
	primitive("image-44", w.Avatar{Src: "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 40 40'%3E%3Crect width='40' height='40' fill='%238C7B6A'/%3E%3Ccircle cx='20' cy='16' r='7' fill='%23E8DFCB'/%3E%3Cpath d='M6 40c2-9 8-13 14-13s12 4 14 13z' fill='%23E8DFCB'/%3E%3C/svg%3E", Size: w.Avatar44}, "components-40", "tc-avatar", 2),
	primitive("image", w.Avatar{Src: "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 40 40'%3E%3Crect width='40' height='40' fill='%238C7B6A'/%3E%3Ccircle cx='20' cy='16' r='7' fill='%23E8DFCB'/%3E%3Cpath d='M6 40c2-9 8-13 14-13s12 4 14 13z' fill='%23E8DFCB'/%3E%3C/svg%3E"}, "components-40", "tc-avatar", 1),
	primitive("initials", w.Avatar{Initials: "АК"}, "components-12", "tc-avatar", 1),
	primitive("photo", w.Avatar{Photo: true}, "components-12", "tc-avatar", 2),
	primitive("sized", w.Avatar{Initials: "АК", Size: w.Avatar34}, "phone-4", "tc-avatar", 1),
	primitive("sized", w.Avatar{Initials: "АК", Size: w.Avatar44}, "s-contacts-7", "tc-avatar", 1),
	primitive("photo", w.Avatar{Photo: true, Size: w.Avatar34}, "phone-4", "tc-avatar", 2),

	// tc-dim
	primitive("plain", dim(w.DimText, w.DimInline, w.Text("· in 7 months")), "s-contacts-6", "tc-dim", 1),
	primitive("mixed", dim(w.Dim11, w.DimInline, w.Text("Last response"), w.Mono{Text: "2026-08-16 09:41"}), "components-35", "tc-dim", 1),
	primitive("sized", dim(w.Dim11, w.DimInline, w.Text("No device connects until it is started.")), "components-35", "tc-dim", 2),
	primitive("mixed", dim(w.Dim11, w.DimInline, w.Text("Last backup"), w.Mono{Text: "2026-08-16 03:10"}, w.Text("·"),
		w.Hint{Text: "verified", Title: "The archive was read back from the destination and the checksum matched"}), "components-35", "tc-dim", 4),
	primitive("wrap", dim(w.Dim11_5, w.DimWrap, w.Text("Анна Ковалёва · from nextcloud · Контакты")), "s-contacts-2", "tc-dim", 1),
	primitive("sized", dim(w.Dim10_5, w.DimInline, w.Text("always")), "components-21", "tc-dim", 1),
	primitive("slot", dim(w.DimText, w.DimSlot, w.Text("From")), "s-mail-7", "tc-dim", 1),
	primitive("struck", dim(w.DimText, w.DimStruck, w.Text("попросить у Артёма скриншоты")), "s-notes-2", "tc-dim", 1),
	primitive("tight", dim(w.Dim11, w.DimTight, w.Text("nextcloud · Контакты")), "s-find-3", "tc-dim", 1),
	primitive("mixed", dim(w.Dim11_5, w.DimWrap, w.Text("Over 2 MB it goes to"), w.Mono{Text: "Documents / carrel"}, w.Text("and the event keeps a link")), "s-io-2", "tc-dim", 2),
	primitive("sized", dim(w.Dim11_5, w.DimInline, w.Text("Properties Carrel does not know are kept unchanged.")), "s-contacts-3", "tc-dim", 1),
	primitive("mixed", dim(w.DimText, w.DimInline, w.Text("starts a copy · answers"), w.Mono{Text: "202"}, w.Text("and an id, a second one while it runs is refused")), "s-backup-5", "tc-dim", 1),
	primitive("indent", dim(w.DimText, w.DimIndent, w.Text("as sent")), "s-agenda-3", "tc-dim", 1),
	primitive("indent", dim(w.Dim11, w.DimIndent, w.Text("derived from the address")), "s-profile-10", "tc-dim", 3),
	primitive("top", dim(w.Dim11, w.DimTop, w.Text("Arguments")), "s-backup-4", "tc-dim", 5),
	primitive("mixed", dim(w.Dim11_5, w.DimWrap, w.Text("Signed"), w.Strong("Михаил Енученко"), w.Text("· replies go to mix@example.org")), "s-mail-1", "tc-dim", 1),
	primitive("fixed", dim(w.DimText, w.DimFixed, w.Text("Sent as")), "s-mail-4", "tc-dim", 3),
	primitive("mixed", dim(w.Dim11_5, w.DimWrap, w.Dot{Tone: w.DotAccent}, w.Text("server stopped")), "desktop-6", "tc-dim", 1),

	// tc-mono
	primitive("span", w.Mono{Text: "14 items · 38.2 MB"}, "components-16", "tc-mono", 2),
	primitive("variant", w.Mono{Text: ".vcf", Variant: w.Mono11}, "components-19", "tc-mono", 1),
	primitive("variant", w.Mono{Text: "12", Variant: w.MonoTight}, "components-11", "tc-mono", 1),
	primitive("variant", w.Mono{Text: "last used", Variant: w.Mono10Dusty}, "components-6", "tc-mono", 1),
	primitive("line", w.Mono{Text: "https://cloud.example.org/remote.php/dav/ · mix · last answer 2 min ago", Form: w.MonoLine, Variant: w.Mono11_5Dusty}, "s-profile-1", "tc-mono", 5),
	primitive("variant", w.Mono{Text: "412 KB", Variant: w.MonoMuted}, "components-16", "tc-mono", 1),
	primitive("variant", w.Mono{Text: "80 points · same phone, same name", Variant: w.MonoTightDusty}, "s-dup-1", "tc-mono", 3),
	primitive("variant", w.Mono{Text: "5", Variant: w.MonoDusty11}, "components-21", "tc-mono", 1),
	primitive("variant", w.Mono{Text: "1.6×", Variant: w.Mono11Dusty}, "s-contacts-5", "tc-mono", 2),
	primitive("variant", w.Mono{Text: "12.4 MB", Variant: w.MonoDusty}, "components-26", "tc-mono", 1),
	primitive("variant", w.Mono{Text: "8 found · 2 min ago", Variant: w.Mono11DustyTight}, "s-find-2", "tc-mono", 1),
	primitive("line", w.Mono{Text: "https://dav.office.example/dav.php/ · m.yenuchenko · last answer refused", Form: w.MonoLine, Variant: w.Mono11_5DustyGap}, "s-profile-1", "tc-mono", 8),
	primitive("variant", w.Mono{Text: "visible only now", Variant: w.MonoTightDanger}, "s-backup-2", "tc-mono", 5),
	primitive("variant", w.Mono{Text: "10:00", Variant: w.MonoFace}, "print-1", "tc-mono", 1),
	primitive("line", w.Mono{Text: "264 px · sections + sources", Form: w.MonoLine, Variant: w.Mono11Top}, "shell-1", "tc-mono", 1),
	primitive("code", w.Mono{Text: "MOVE", Form: w.MonoCode}, "s-contacts-2", "tc-mono", 6),
	primitive("variant", w.Mono{Text: "3 collections", Variant: w.Mono11Indent}, "s-backup-1", "tc-mono", 4),
	primitive("variant", w.Mono{Text: "carrel.example", Variant: w.MonoTightPlain}, "account-5", "tc-mono", 1),
	primitive("variant", w.Mono{Text: "<mix@carrel.example>", Variant: w.Mono11Muted}, "s-mail-7", "tc-mono", 1),
	primitive("variant", w.Mono{Text: "go test -race ./...", Variant: w.MonoChip}, "s-publish-10", "tc-mono", 2),
)

func dim(size w.DimSize, variant w.DimVariant, parts ...w.Inline) w.Dim {
	return w.Dim{Parts: parts, Size: size, Variant: variant}
}

func init() {
	Widgets = append(Widgets,
		w.Button{}, w.Close{}, w.Mark{}, w.Tag{}, w.Badge{}, w.Tick{}, w.Kbd{}, w.Num{},
		w.Avatar{}, w.Dim{}, w.Mono{}, w.Plain{},
	)
}
