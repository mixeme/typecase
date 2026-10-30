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
	primitive("plain", w.Button{Label: "How to mount", Menu: true, Small: true, Focus: true}, "s-publish-32", "tc-btn", 2),
	primitive("plain", w.Button{Label: "Download", Icon: "download", Kind: w.ButtonPrimary, Submit: true, Aside: &w.Mono{Text: ".ics", Variant: w.Mono11}}, "s-agenda-30", "tc-btn", 1),
	primitive("plain", w.Button{Icon: "dots", Variant: w.ButtonPadX3, Title: "More"}, "s-contacts-33", "tc-btn", 2),
	primitive("plain", w.Button{Label: "Merge on the server", Kind: w.ButtonDanger, State: w.ButtonDisabled}, "s-dup-1", "tc-btn", 12),
	primitive("danger-off", w.Button{Label: "End", Kind: w.ButtonDanger, Small: true, Title: "This is the session you are using: sign out instead.", State: w.ButtonDisabled}, "account-1", "tc-btn", 6),
	primitive("picks-small", w.Button{Label: "Keep both", Small: true, File: &w.FileField{Form: "taken-release", Name: "keep", Label: "Keep both"}}, "components-26", "tc-btn", 2),
	primitive("disabled", w.Button{Folds: true, Label: "Rename", Icon: "pencil", Title: "Public is read-only.", Form: "files-sel", Name: "op", Value: "rename", State: w.ButtonDisabled}, "states-34", "tc-btn", 11),
	primitive("disabled", w.Button{Folds: true, Label: "Delete", Icon: "trash", Kind: w.ButtonDangerSolid, Title: "Public is read-only.", Form: "files-sel", Name: "op", Value: "delete", State: w.ButtonDisabled}, "states-34", "tc-btn", 14),
	primitive("with-content", w.Button{Label: "Upload", Icon: "upload", Kind: w.ButtonPrimary, File: &w.FileField{Form: "post-upload", Name: "file", Label: "Files", Size: "size", Many: true}}, "states-32", "tc-btn", 11),
	primitive("wait", w.Button{Label: "Write this combination", Kind: w.ButtonPrimary, Reason: "Choose a side for every property.", State: w.ButtonWaits, Form: "conflict"}, "s-conflict-6", "tc-btn", 3),
	primitive("link", w.Button{Label: "Try again", Small: true, Href: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/conflict"}, "s-conflict-10", "tc-btn", 3),
	primitive("primary", w.Button{Label: "Done", Kind: w.ButtonPrimary, Href: "/admin/users"}, "s-admin-14", "tc-btn", 2),
	primitive("disabled", w.Button{Label: "Attach", Icon: "clip", Kind: w.ButtonPrimary, State: w.ButtonDisabled}, "s-files-16", "tc-btn", 1),
	primitive("menu", w.Button{Label: "9 records", Menu: true, Small: true}, "s-contacts-29", "tc-btn", 1),
	primitive("tertiary", w.Button{Label: "Print", Icon: "print", Tertiary: true}, "s-contacts-6", "tc-btn", 5),
	primitive("link", w.Button{Label: "Edit this copy", Kind: w.ButtonPrimary, Submit: true}, "s-contacts-32", "tc-btn", 1),
	primitive("disabled-plain", w.Button{Label: "Send invitation", Title: "SMTP is not configured: no letter can go.", Submit: true, State: w.ButtonDisabled}, "s-admin-18", "tc-btn", 2),
	primitive("copies", w.Button{Label: "Copy", Small: true, Copy: "https://carrel.example/p/9c1f4e", Copied: "Copied"}, "s-publish-1", "tc-btn", 3),
	primitive("copies", w.Button{Label: "Copy the script", Icon: "dup", Kind: w.ButtonPrimary, CopyFrom: "trigger-script", Copied: "Copied"}, "s-backup-3", "tc-btn", 1),
	primitive("more", w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}, "s-publish-1", "tc-btn", 6),
	primitive("picker", w.Button{Label: "Replace", File: &w.FileField{Form: "post-photo", Name: "photo", Accept: "image/*", Label: "Photo"}}, "s-contacts-3", "tc-btn", 5),
	primitive("link-icon", w.Button{Label: "Sources", Icon: "cal", Href: "#sources-demo"}, "components-41", "tc-btn", 1),
	primitive("post", w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Form: "post-disconnect"}, "components-37", "tc-btn", 5),
	primitive("link-off", w.Button{Label: "Export", Title: "Nothing to export: the list is empty.", State: w.ButtonDisabled, Href: "#"}, "components-37", "tc-btn", 4),
	primitive("link-primary", w.Button{Label: "New contact", Icon: "plus", Kind: w.ButtonPrimary, Href: "/app/contacts/nextcloud/kontakty/new"}, "components-37", "tc-btn", 3),
	primitive("link", w.Button{Label: "Import", Href: "/app/contacts/nextcloud/kontakty/import"}, "components-37", "tc-btn", 2),
	primitive("disabled", w.Button{Label: "Rewrite 0 records", Kind: w.ButtonPrimary, State: w.ButtonDisabled}, "s-contacts-20", "tc-btn", 1),
	primitive("disabled", w.Button{Icon: "plus", Title: "Open a folder to make one in it.", State: w.ButtonDisabled}, "s-files-6", "tc-btn", 3),
	primitive("disabled", w.Button{Label: "New folder", Icon: "plus", State: w.ButtonDisabled}, "s-files-15", "tc-btn", 2),
	primitive("disabled", w.Button{Label: "Upload", Menu: true, Icon: "upload", Kind: w.ButtonPrimary, State: w.ButtonDisabled}, "s-files-15", "tc-btn", 1),
	primitive("disabled", w.Button{Label: "Back to defaults", Small: true, Title: "Every weight is the default already.", State: w.ButtonDisabled}, "s-dup-7", "tc-btn", 3),
	primitive("icon", w.Button{Label: "2 min ago", Icon: "refresh"}, "shell-1", "tc-btn", 2),
	primitive("small", w.Button{Label: "Edit", Small: true}, "components-14", "tc-btn", 1),
	primitive("danger", w.Button{Label: "1 failed", Small: true, Menu: true, Kind: w.ButtonDanger}, "s-contacts-7", "tc-btn", 10),
	primitive("normal", w.Button{Label: "Import"}, "components-9", "tc-btn", 2),
	primitive("menu", w.Button{Label: "New note", Icon: "quill", Menu: true, Kind: w.ButtonPrimary}, "shell-1", "tc-btn", 1),
	primitive("icon", w.Button{Icon: "dots", Title: "More"}, "components-16", "tc-btn", 6),
	primitive("primary", w.Button{Label: "Save", Kind: w.ButtonPrimary}, "components-9", "tc-btn", 1),
	primitive("primary", w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary}, "components-9", "tc-btn", 7),
	primitive("variant", w.Button{Icon: "dots", Small: true, Variant: w.ButtonPadX3, Title: "More"}, "s-profile-1", "tc-btn", 11),
	primitive("folds", w.Button{Folds: true, Label: "Download", Icon: "download", Title: "Download"}, "components-16", "tc-btn", 1),
	primitive("small", w.Button{Icon: "dots", Small: true, Title: "More"}, "s-contacts-1", "tc-btn", 16),
	primitive("danger-solid", w.Button{Label: "Delete contact", Kind: w.ButtonDangerSolid}, "components-9", "tc-btn", 5),
	primitive("danger", w.Button{Label: "Disconnect", Kind: w.ButtonDanger}, "components-9", "tc-btn", 6),
	primitive("small", w.Button{Label: "Add phone", Icon: "plus", Small: true}, "s-contacts-3", "tc-btn", 10),
	primitive("danger", w.Button{Label: "End", Small: true, Kind: w.ButtonDanger}, "account-1", "tc-btn", 7),
	primitive("focus", w.Button{Label: "New note", Icon: "quill", Menu: true, Kind: w.ButtonPrimary, Focus: true}, "s-notes-5", "tc-btn", 1),
	primitive("folds", w.Button{Label: "Delete", Icon: "trash", Folds: true, Kind: w.ButtonDangerSolid, Title: "Delete"}, "components-16", "tc-btn", 3),
	primitive("danger-solid", w.Button{Label: "Delete contact", Icon: "trash", Kind: w.ButtonDangerSolid}, "s-contacts-6", "tc-btn", 9),
	primitive("menu", w.Button{Icon: "plus", Title: "New address book"}, "s-mail-1", "tc-btn", 3),
	// The summary of a closed menu is its button (Drop prints it so).
	primitive("menu", w.Button{Label: "Attach", Icon: "clip", Menu: true, Small: true}, "s-agenda-1", "tc-btn", 23),
	primitive("link", w.Button{Icon: "back", Title: "Previous week", Href: "/app/calendar?from=2026-08-03"}, "s-agenda-1", "tc-btn", 10),
	primitive("menu", w.Button{Label: "Attach", Menu: true, Icon: "clip"}, "s-agenda-3", "tc-btn", 6),
	primitive("icon", w.Button{Icon: "dots", Title: "More"}, "s-contacts-6", "tc-btn", 6),
	primitive("small", w.Button{Icon: "dots", Title: "More", Small: true}, "s-contacts-7", "tc-btn", 14),
	primitive("variant", w.Button{Icon: "menu", IconBox: w.IconBox19, Variant: w.ButtonPadX4, Title: "Sections"}, "phone-1", "tc-btn", 1),
	primitive("with-content", w.Button{Label: "B", Bold: true}, "components-29", "tc-btn", 3),
	primitive("variant", w.Button{Icon: "quill", IconBox: w.IconBox16, Variant: w.ButtonPadX8, Kind: w.ButtonPrimary, Title: "New note"}, "phone-1", "tc-btn", 3),
	primitive("variant", w.Button{Icon: "dots", IconBox: w.IconBox16, Variant: w.ButtonPadX4, Title: "More"}, "phone-6", "tc-btn", 7),
	primitive("variant", w.Button{Label: "I", Variant: w.ButtonItalic}, "s-notes-20", "tc-btn", 6),
	primitive("folds", w.Button{Folds: true, Label: "Attach", Menu: true, Icon: "clip", Title: "Attach"}, "s-notes-20", "tc-btn", 11),
	primitive("off", w.Button{Label: "Rename", Icon: "pencil", Folds: true, State: w.ButtonOff, Title: "Rename"}, "s-files-2", "tc-btn", 12),
	primitive("focus", w.Button{Label: "Weights", Small: true, Focus: true}, "s-dup-3", "tc-btn", 1),
	primitive("focus", w.Button{Icon: "plus", Focus: true, Title: "New calendar"}, "s-profile-5", "tc-btn", 3),
	primitive("variant", w.Button{Icon: "search", IconBox: w.IconBox17, Variant: w.ButtonPadX6, Title: "Search", Href: "/app/search"}, "phone-1", "tc-btn", 2),
	primitive("primary", w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "Where to write"}, "components-9", "tc-btn", 8),
	primitive("with-content", w.Button{Label: "Columns", Icon: "cols", Variant: w.ButtonOnCard, Aside: &w.Mono{Text: "5", Variant: w.MonoDusty11}}, "components-21", "tc-btn", 1),
	primitive("with-content", w.Button{Label: "Mark done", Tick: &w.Tick{Glyph: true}}, "s-tasks-2", "tc-btn", 4),
	primitive("variant", w.Button{Label: "Focus", Icon: "focus", Folds: true, Variant: w.ButtonOnCard, Title: "Focus"}, "s-notes-4", "tc-btn", 13),
	primitive("variant", w.Button{Icon: "panel", Variant: w.ButtonOnCard, Title: "Details"}, "s-files-1", "tc-btn", 9),
	primitive("disabled", w.Button{Label: "Keep these values", Kind: w.ButtonPrimary, Name: "do", Value: "choose", Submit: true}, "s-dup-1", "tc-btn", 16),
	primitive("primary", w.Button{Label: "Save", Small: true, Kind: w.ButtonPrimary}, "s-dup-3", "tc-btn", 2),
	primitive("with-content", w.Button{Label: "Download", Icon: "download", Href: "/app/collections/nextcloud/dni-rozhdeniya/export", Aside: &w.Mono{Text: ".ics", Variant: w.Mono11}, After: "first"}, "s-profile-9", "tc-btn", 9),
	primitive("with-content", w.Button{Label: "Save as", Icon: "download", SaveAs: "carrel-personal.sh", SaveFrom: "trigger-script", Aside: &w.Mono{Text: "carrel-personal.sh", Variant: w.Mono11}}, "s-backup-3", "tc-btn", 2),
	primitive("tertiary", w.Button{Label: "H2", Tertiary: true}, "s-notes-20", "tc-btn", 7),
	primitive("tertiary", w.Button{Folds: true, Label: "task", Icon: "task", Title: "Task", Tertiary: true}, "s-notes-20", "tc-btn", 9),
	primitive("tertiary", w.Button{Icon: "link", Title: "Link", Tertiary: true}, "s-notes-20", "tc-btn", 10),
	primitive("variant", w.Button{Label: "Copy", Small: true, Variant: w.ButtonSelfStart}, "s-backup-4", "tc-btn", 5),
	primitive("disabled", w.Button{Label: "Save", Kind: w.ButtonPrimary, Form: "admin-settings"}, "s-admin-3", "tc-btn", 3),
	primitive("variant", w.Button{Label: "Write a note", Variant: w.ButtonCentredFill}, "phone-3", "tc-btn", 3),
	primitive("variant", w.Button{Label: "Delete", Kind: w.ButtonDangerSolid, Variant: w.ButtonCentred}, "phone-3", "tc-btn", 4),
	primitive("danger-solid", w.Button{Icon: "trash", Kind: w.ButtonDangerSolid, Title: "Delete"}, "phone-6", "tc-btn", 5),

	// tc-close
	primitive("link", w.Close{Label: "Close", Href: "#"}, "components-41", "tc-close", 1),
	primitive("default", w.Close{Label: "Close the panel"}, "components-33", "tc-close", 1),

	// tc-mark
	primitive("plain", w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, "components-14", "tc-mark", 1),
	primitive("one", w.Mark{Colour: "#4A6B52"}, "components-11", "tc-mark", 2),
	primitive("multi", w.Mark{Kind: w.MarkMulti}, "components-11", "tc-mark", 1),
	primitive("multi", w.Mark{Kind: w.MarkMulti, Size: w.MarkH20}, "s-contacts-1", "tc-mark", 6),
	primitive("sized", w.Mark{Colour: "#A05A2C", Size: w.MarkH11Inline}, "s-contacts-6", "tc-mark", 4),
	primitive("external", w.Mark{Colour: "#8C6239", Kind: w.MarkExternal}, "s-profile-1", "tc-mark", 17),
	primitive("sized", w.Mark{Colour: "#4A6B52", Size: w.MarkH20}, "s-agenda-17", "tc-mark", 10),
	primitive("sized", w.Mark{Colour: "#4A6B52"}, "s-mail-1", "tc-mark", 16),
	primitive("sized", w.Mark{Colour: "#4A6B52", Size: w.MarkH32}, "phone-1", "tc-mark", 2),
	primitive("sized", w.Mark{Colour: "#4A6B52", Size: w.MarkH34}, "phone-4", "tc-mark", 1),
	primitive("sized", w.Mark{Size: w.MarkOnHair}, "states-3", "tc-mark", 11),
	primitive("sized", w.Mark{Colour: "#4A6B52", Size: w.MarkH12}, "s-notes-2", "tc-mark", 4),
	primitive("sized", w.Mark{Colour: "#4A6B52", Size: w.MarkH14}, "desktop-8", "tc-mark", 1),

	// tc-tag
	primitive("unloaded", w.Tag{Label: "7c1e-40aa-b913", Removable: true}, "s-tasks-11", "tc-tag", 3),
	primitive("plain", w.Tag{Label: "работа"}, "s-contacts-1", "tc-tag", 2),
	primitive("on", w.Tag{Label: "all", On: true}, "s-contacts-1", "tc-tag", 1),
	primitive("icon", w.Tag{Label: "smirnov.vcf", Icon: "con"}, "s-mail-7", "tc-tag", 1),
	primitive("removable", w.Tag{Label: "Забрать документы", Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, Removable: true}, "s-tasks-3", "tc-tag", 1),
	// Carrel's task 77: the attendees — a known person by the colour of their book and name, an unknown one by the address.
	primitive("person", w.Tag{Label: "Мария Лебедева", Mark: &w.Mark{Colour: "#6B4E7A", Size: w.MarkH11}, Fixed: true, Aside: &w.Dim{Parts: []w.Inline{w.Text("· organiser")}}, Title: "m.lebedeva@example.org", Value: "mailto:m.lebedeva@example.org"}, "s-agenda-3", "tc-tag", 1),
	primitive("person", w.Tag{Label: "Анна Ковалёва", Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, Removable: true, Aside: &w.Dim{Parts: []w.Inline{w.Text("· accepted")}}, Title: "a.kovaleva@example.org", Value: "mailto:a.kovaleva@example.org"}, "s-agenda-3", "tc-tag", 2),
	primitive("person", w.Tag{Address: &w.Mono{Text: "v@example.org"}, Removable: true, Aside: &w.Dim{Parts: []w.Inline{w.Text("· no reply")}}, Title: "v@example.org", Value: "mailto:v@example.org"}, "s-agenda-3", "tc-tag", 4),
	primitive("person", w.Tag{Label: "Вера Лебедева", Mark: &w.Mark{Colour: "#A05A2C", Size: w.MarkH11}, Removable: true, Title: "vera@example.net", Name: "attendee", Value: "mailto:vera@example.net"}, "s-agenda-12", "tc-tag", 5),
	primitive("person", w.Tag{Label: "Мария Лебедева", Mark: &w.Mark{Colour: "#6B4E7A", Size: w.MarkH11}, Fixed: true, Aside: &w.Dim{Parts: []w.Inline{w.Text("· organiser")}}, Title: "m.lebedeva@example.org"}, "s-agenda-14", "tc-tag", 1),
	primitive("person", w.Tag{Address: &w.Mono{Text: "v@example.org"}, Fixed: true, Aside: &w.Dim{Parts: []w.Inline{w.Text("· no reply")}}, Title: "v@example.org"}, "s-agenda-14", "tc-tag", 4),
	primitive("mark", w.Tag{Label: "работа"}, "s-mail-1", "tc-tag", 2),

	// tc-badge
	primitive("neutral", w.Badge{Label: "external"}, "components-11", "tc-badge", 1),
	primitive("linked", w.Badge{Label: "linked ×2", Tone: w.BadgeLinked}, "components-12", "tc-badge", 1),
	primitive("alert", w.Badge{Label: "tentative", Tone: w.BadgeAlert}, "s-agenda-1", "tc-badge", 8),
	primitive("accent", w.Badge{Label: "credentials unlocked · 4:12", Accent: true}, "s-profile-4", "tc-badge", 1),
	primitive("local", w.Badge{Label: "local", Tone: w.BadgeLocal, Title: "Local instance: may be unreachable from outside"}, "s-plan-1", "tc-badge", 2),

	// tc-tick
	primitive("sign", w.Tick{Done: true, Sign: true}, "s-publish-28", "tc-tick", 4),
	primitive("done", w.Tick{Done: true, Label: "Done"}, "s-tasks-1", "tc-tick", 5),
	primitive("open", w.Tick{Label: "Done"}, "s-tasks-1", "tc-tick", 1),
	primitive("sign", w.Tick{Variant: w.TickNudge, Sign: true}, "s-notes-1", "tc-tick", 1),
	primitive("sign", w.Tick{Done: true, Variant: w.TickNudge, Sign: true}, "s-notes-2", "tc-tick", 2),
	primitive("variant", w.Tick{Variant: w.TickEdgeDanger, Label: "Check"}, "s-admin-4", "tc-tick", 6),
	primitive("glyph", w.Tick{Glyph: true}, "s-tasks-2", "tc-tick", 1),
	primitive("variant", w.Tick{Variant: w.TickEdgeDusty, Label: "Check"}, "s-admin-4", "tc-tick", 8),

	// tc-kbd, tc-num, tc-plain
	primitive("default", w.Kbd{Keys: "Ctrl+S"}, "components-29", "tc-kbd", 1),
	primitive("normal", w.Num{Value: 3}, "components-12", "tc-num", 1),
	primitive("alert", w.Num{Value: 2, Alert: true}, "s-tasks-1", "tc-num", 6),
	primitive("default", w.Plain{Text: "2 active"}, "account-1", "tc-plain", 2),

	// tc-avatar
	primitive("image-44", w.Avatar{Src: "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 40 40'%3E%3Crect width='40' height='40' fill='%238C7B6A'/%3E%3Ccircle cx='20' cy='16' r='7' fill='%23E8DFCB'/%3E%3Cpath d='M6 40c2-9 8-13 14-13s12 4 14 13z' fill='%23E8DFCB'/%3E%3C/svg%3E", Size: w.Avatar44}, "components-40", "tc-avatar", 2),
	primitive("image", w.Avatar{Src: "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 40 40'%3E%3Crect width='40' height='40' fill='%238C7B6A'/%3E%3Ccircle cx='20' cy='16' r='7' fill='%23E8DFCB'/%3E%3Cpath d='M6 40c2-9 8-13 14-13s12 4 14 13z' fill='%23E8DFCB'/%3E%3C/svg%3E"}, "components-40", "tc-avatar", 1),
	primitive("initials", w.Avatar{Initials: "АК"}, "components-12", "tc-avatar", 1),
	primitive("photo", w.Avatar{Photo: true}, "components-12", "tc-avatar", 2),
	primitive("sized", w.Avatar{Initials: "АК", Size: w.Avatar34}, "phone-4", "tc-avatar", 1),
	primitive("sized", w.Avatar{Initials: "АК", Size: w.Avatar44}, "s-contacts-7", "tc-avatar", 1),
	primitive("photo", w.Avatar{Photo: true, Size: w.Avatar34}, "phone-4", "tc-avatar", 2),

	// tc-dim
	primitive("mark", w.Dim{Mark: &w.Mark{Colour: "#4A6B52", Size: w.MarkH11Inline}, Parts: []w.Inline{w.Text("Контакты")}}, "s-contacts-1", "tc-dim", 1),
	primitive("plain", w.Dim{Parts: []w.Inline{w.Text("issued 2 August · last run 16 August 03:10 from "), w.Mono{Text: "10.0.0.5"}, w.Text(" · 14 runs")}, Size: w.Dim11_5}, "s-backup-12", "tc-dim", 13),
	primitive("plain", w.Dim{Parts: []w.Inline{w.Text("· in 7 months")}}, "s-contacts-6", "tc-dim", 1),
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
	primitive("top", dim(w.Dim11, w.DimTop, w.Text("Arguments")), "s-backup-4", "tc-dim", 5),
	primitive("mixed", dim(w.Dim11_5, w.DimWrap, w.Text("Signed"), w.Strong("Михаил Енученко"), w.Text("· replies go to mix@example.org")), "s-mail-1", "tc-dim", 3),
	primitive("fixed", dim(w.DimText, w.DimFixed, w.Text("Sent as")), "s-mail-4", "tc-dim", 3),
	primitive("mixed", dim(w.Dim11_5, w.DimWrap, w.Dot{Tone: w.DotAccent}, w.Text("server stopped")), "desktop-6", "tc-dim", 1),

	// tc-mono
	// words in two lengths (Р-65): the run of a search, its answer's time
	primitive("plain", w.Say{Full: "Done · 13 of 14 sources answered", Short: "13 of 14"}, "s-find-1", "tc-say", 1),
	primitive("variant", w.Mono{Text: "80 points · same phone, same name", Variant: w.MonoTightDusty}, "s-dup-1", "tc-mono", 3),
	primitive("code", w.Mono{Text: ".well-known", Form: w.MonoCode}, "s-profile-2", "tc-mono", 4),
	primitive("plain", w.Mono{Text: "7 found · 2 min ago", Variant: w.Mono11DustyTight}, "s-find-6", "tc-mono", 1),
	primitive("danger", w.Mono{Text: "visible only now", Variant: w.MonoTightDanger}, "s-backup-13", "tc-mono", 7),
	primitive("span", w.Mono{Text: "14 items · 38.2 MB"}, "components-16", "tc-mono", 2),
	primitive("variant", w.Mono{Text: ".vcf", Variant: w.Mono11}, "components-19", "tc-mono", 1),
	primitive("variant", w.Mono{Text: "12", Variant: w.MonoTight}, "components-11", "tc-mono", 1),
	primitive("variant", w.Mono{Text: "last used", Variant: w.Mono10Dusty}, "components-6", "tc-mono", 1),
	primitive("line", w.Mono{Text: "https://cloud.example.org/remote.php/dav/ · mix · last answer 2 min ago", Form: w.MonoLine, Variant: w.Mono11_5Dusty}, "s-profile-1", "tc-mono", 5),
	primitive("variant", w.Mono{Text: "412 KB", Variant: w.MonoMuted}, "components-16", "tc-mono", 1),
	primitive("variant", w.Mono{Text: "5", Variant: w.MonoDusty11}, "components-21", "tc-mono", 1),
	primitive("variant", w.Mono{Text: "1.6×", Variant: w.Mono11Dusty}, "s-contacts-5", "tc-mono", 2),
	primitive("variant", w.Mono{Text: "12.4 MB", Variant: w.MonoDusty}, "components-26", "tc-mono", 1),
	primitive("line", w.Mono{Text: "https://dav.office.example/dav.php/ · m.yenuchenko · last answer refused", Form: w.MonoLine, Variant: w.Mono11_5DustyGap}, "s-profile-1", "tc-mono", 8),
	primitive("variant", w.Mono{Text: "10:00", Variant: w.MonoFace}, "print-1", "tc-mono", 1),
	primitive("line", w.Mono{Text: "264 px · sections + sources", Form: w.MonoLine, Variant: w.Mono11Top}, "shell-1", "tc-mono", 1),
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
		w.Avatar{}, w.Dim{}, w.Mono{}, w.Plain{}, w.Say{},
	)
}
