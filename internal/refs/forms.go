package refs

import w "carrel.dev/typecase/widget"

func seg(on int, labels ...string) []w.SegOption {
	var out []w.SegOption
	for i, l := range labels {
		out = append(out, w.SegOption{Label: l, On: i == on})
	}
	return out
}

// A widget prints the pieces of running text one after another, as the
// frame writes them; these helpers put the space between words that the
// frames have, so a record reads as a sentence.
func check(on bool, parts ...w.CheckPart) w.Check {
	return w.Check{On: on, Parts: spacedParts(parts, w.CheckPart(w.Text(" ")))}
}

func hint(parts ...w.Inline) w.FieldHint {
	return w.FieldHint{Parts: spacedParts(parts, w.Inline(w.Text(" ")))}
}

func spacedParts[T any](parts []T, space T) []T {
	var out []T
	for i, p := range parts {
		if i > 0 {
			out = append(out, space)
		}
		out = append(out, p)
	}
	return out
}

var btn = w.Button{Label: "x"}

var _ = add(
	// tc-seg
	form("start-buttons", w.Seg{Options: []w.SegOption{{Label: "From a job", Href: "/app/settings/backup/restore?from=job", On: true}, {Label: "From a server", Href: "/app/settings/backup/restore?from=server"}, {Label: "From this computer", Href: "/app/settings/backup/restore?from=file"}}, Variant: w.SegStart}, "s-backup-7", "tc-seg", 1),
	form("radio-four", w.Seg{Options: []w.SegOption{{Label: "crontab", On: true, Value: "crontab"}, {Label: "GoSentry", Value: "gosentry"}, {Label: "Gitea", Value: "gitea"}, {Label: "curl", Value: "curl"}}, Variant: w.SegStart, Name: "choice"}, "s-backup-2", "tc-seg", 2),
	form("radio-start", w.Seg{Options: []w.SegOption{{Label: "Markdown · .md", On: true, Value: "markdown.md"}, {Label: "Source · .ics", Value: "source.ics"}}, Variant: w.SegStart, Name: "send.it.as"}, "s-mail-5", "tc-seg", 1),
	form("icons", w.Seg{Options: []w.SegOption{{Icon: "menu", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-07?view=list", Title: "List"}, {Icon: "grid", Href: "/app/files/sftpgo/documents/carrel/attachments/2026-07?view=tiles", On: true, Title: "Tiles"}}}, "s-files-5", "tc-seg", 2),
	form("order", w.Seg{Options: []w.SegOption{{Label: "Name", Href: "/app/files/sftpgo/documents?sort=name", On: true}, {Label: "Size", Href: "/app/files/sftpgo/documents?sort=size"}, {Label: "Changed", Href: "/app/files/sftpgo/documents?sort=changed"}}}, "s-files-1", "tc-seg", 1),
	form("links", w.Seg{Options: []w.SegOption{{Label: "A–Z", Href: "/app/contacts?sort=az", On: true}, {Label: "Recently changed", Href: "/app/contacts?sort=changed"}}}, "s-contacts-1", "tc-seg", 1),
	form("buttons", w.Seg{Options: []w.SegOption{{Label: "Apple"}, {Label: "Android", On: true}, {Label: "Thunderbird"}, {Label: "Anything else"}}}, "s-plan-2", "tc-seg", 1),
	form("start", w.Seg{Options: []w.SegOption{{Label: "crontab"}, {Label: "GoSentry"}, {Label: "Gitea"}, {Label: "curl", On: true}}, Variant: w.SegStart}, "s-backup-11", "tc-seg", 1),
	form("start", w.Seg{Options: []w.SegOption{{Label: "crontab"}, {Label: "GoSentry"}, {Label: "Gitea", On: true}, {Label: "curl"}}, Variant: w.SegStart}, "s-backup-10", "tc-seg", 1),
	form("default", w.Seg{Options: []w.SegOption{{Label: "All 2", On: true}, {Label: "Contacts 1"}, {Label: "Events 1"}, {Label: "Tasks 0"}, {Label: "Notes 0"}}}, "s-dup-1", "tc-seg", 1),
	form("views", w.SegFold{Seg: w.Seg{Options: []w.SegOption{{Label: "Day", Href: "/app/calendar?view=day"}, {Label: "Week", Href: "/app/calendar?view=week", On: true}, {Label: "Month", Href: "/app/calendar?view=month"}, {Label: "Range", Href: "/app/calendar?view=range"}}}, Phone: true}, "s-agenda-1", "tc-segfold", 1),
	form("default", w.SegFold{Seg: w.Seg{Options: []w.SegOption{{Label: "All", Count: "24", Href: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/timeline", On: true}, {Label: "Events", Count: "14", Href: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/timeline?kind=event"}, {Label: "Tasks", Count: "2", Href: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/timeline?kind=task"}, {Label: "Notes", Count: "4", Href: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/timeline?kind=note"}, {Label: "Files", Count: "4", Href: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/timeline?kind=file"}}}}, "s-contacts-7", "tc-segfold", 1),
	form("sort-changed", w.Seg{Options: []w.SegOption{{Label: "By due"}, {Label: "By priority"}, {Label: "Recently changed", On: true}}}, "s-tasks-13", "tc-seg", 2),
	form("radio", w.Seg{Options: []w.SegOption{{Label: "Auto", On: true, Value: "auto"}, {Label: "Light", Value: "light"}, {Label: "Dark", Value: "dark"}}, Name: "theme", Theme: true}, "components-38", "tc-seg", 1),
	form("table-line", w.Seg{Options: []w.SegOption{{Label: "Wrap"}, {Label: "One line", On: true}}, Variant: w.SegCompactSmall}, "s-notes-22", "tc-seg", 1),
	form("table-wrap", w.Seg{Options: []w.SegOption{{Label: "Wrap", On: true}, {Label: "One line"}}, Variant: w.SegCompactSmall}, "s-notes-21", "tc-seg", 1),
	form("default", w.Seg{Options: seg(0, "Remote", "Local")}, "components-13", "tc-seg", 1),
	form("default", w.Seg{Options: seg(0, "Open 12", "Done 48", "All")}, "components-9", "tc-seg", 1),
	form("start", w.Seg{Options: seg(0, "0°", "90°", "180°", "270°"), Variant: w.SegStart}, "s-contacts-5", "tc-seg", 1),
	form("default", w.CheckForm{Checks: []w.Check{w.Check{Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}}, Action: "/app/contacts", Apply: w.Button{Label: "Apply", Small: true}}, "s-contacts-1", "tc-checkform", 2),
	form("links", w.Seg{Options: []w.SegOption{{Label: "All 24", Href: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/timeline", On: true}, {Label: "Events 14", Href: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/timeline?kind=event"}, {Label: "Tasks 2", Href: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/timeline?kind=task"}, {Label: "Notes 4", Href: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/timeline?kind=note"}, {Label: "Files 4", Href: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/timeline?kind=file"}}}, "s-contacts-7", "tc-seg", 1),
	form("start", w.Seg{Options: []w.SegOption{{Label: "To WebDAV", On: true, Value: "to.webdav"}, {Label: "Download", Value: "download"}, {Label: "Both", Value: "both"}}, Variant: w.SegStart, Name: "choice"}, "s-backup-2", "tc-seg", 1),
	form("start", w.Seg{Options: []w.SegOption{{Label: "Comfortable", On: true, Value: "comfortable"}, {Label: "Compact", Value: "compact"}}, Variant: w.SegStart, Name: "density"}, "appearance-1", "tc-seg", 2),
	form("default", w.Seg{Options: []w.SegOption{{Label: "Day", Href: "/app/calendar?view=day"}, {Label: "Week", Href: "/app/calendar?view=week", On: true}, {Label: "Month", Href: "/app/calendar?view=month"}, {Label: "Range", Href: "/app/calendar?view=range"}}}, "s-agenda-1", "tc-seg", 1),
	form("compact", w.Seg{Options: seg(0, "Auto", "Light", "Dark"), Variant: w.SegStartSmall}, "account-5", "tc-seg", 1),
	form("default", w.Seg{Options: seg(0, "Apple", "Android", "Thunderbird", "Anything else")}, "s-plan-1", "tc-seg", 1),
	form("start", w.Seg{Options: []w.SegOption{{Label: "Read only", Value: "read.only"}, {Label: "Read and write", On: true, Value: "read.and.write"}}, Variant: w.SegStart, Name: "what.this.device.may.do"}, "s-davloom-2", "tc-seg", 1),
	form("several", w.Seg{Options: []w.SegOption{{Label: "MO", On: true, Value: "MO"}, {Label: "TU", Value: "TU"}, {Label: "WE", Value: "WE"}, {Label: "TH", Value: "TH"}, {Label: "FR", Value: "FR"}, {Label: "SA", Value: "SA"}, {Label: "SU", Value: "SU"}}, Name: "byday", Multiple: true}, "s-agenda-3", "tc-seg", 1),
	form("several", w.Seg{Options: []w.SegOption{{Label: "MO", Value: "MO"}, {Label: "TU", Value: "TU"}, {Label: "WE", Value: "WE"}, {Label: "TH", Value: "TH"}, {Label: "FR", Value: "FR"}, {Label: "SA", Value: "SA"}, {Label: "SU", Value: "SU"}}, Name: "byday", Multiple: true}, "s-agenda-15", "tc-seg", 1),
	form("default", w.Seg{Options: []w.SegOption{{Label: "Full width", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012?width=full"}, {Label: "Reading width", Href: "/app/notes/nextcloud/zapisi/9e31-40b7-c012?width=read", On: true}}, Swap: "note-main"}, "s-notes-3", "tc-seg", 1),
	form("icons", w.Seg{Options: []w.SegOption{{Icon: "menu", Href: "/app/files/sftpgo/documents?view=list", On: true, Title: "List"}, {Icon: "grid", Href: "/app/files/sftpgo/documents?view=tiles", Title: "Tiles"}}}, "s-files-1", "tc-seg", 2),
	form("default", w.Seg{Options: seg(0, "All", "Contacts", "Events", "Tasks", "Notes", "Files")}, "s-find-1", "tc-seg", 1),
	form("default", w.Seg{Options: seg(2, "Apple", "Android", "Thunderbird", "Anything else")}, "s-plan-4", "tc-seg", 1),
	form("default", w.Seg{Options: seg(3, "Apple", "Android", "Thunderbird", "Anything else")}, "s-plan-5", "tc-seg", 1),
	form("compact", w.Seg{Options: seg(1, "read only", "read and write"), Variant: w.SegCompact}, "s-davloom-1", "tc-seg", 1),
	form("compact", w.Seg{Options: seg(0, "read only", "read and write"), Variant: w.SegCompact}, "s-davloom-1", "tc-seg", 2),
	form("start", w.Seg{Options: []w.SegOption{{Label: "To WebDAV", Value: "to.webdav"}, {Label: "Download", Value: "download"}, {Label: "Both", On: true, Value: "both"}}, Variant: w.SegStart, Name: "where"}, "s-backup-3", "tc-seg", 1),
	form("start", w.Seg{Options: seg(1, "crontab", "GoSentry", "Gitea", "curl"), Variant: w.SegStart}, "s-backup-4", "tc-seg", 1),

	// tc-check
	form("off", check(false, w.Text("Print with photos")), "components-10", "tc-check", 2),
	form("on", check(true, w.Text("All day")), "components-10", "tc-check", 1),
	form("radio", w.Check{Name: "choice.FN", Value: "2", Radio: true, Parts: []w.CheckPart{w.Text("Лебедева Вера")}}, "s-dup-1", "tc-check", 2),
	form("radio-on", w.Check{On: true, Name: "choice.FN", Value: "1", Radio: true, Parts: []w.CheckPart{w.Text("Вера Лебедева")}}, "s-dup-1", "tc-check", 1),
	form("rich", w.Check{On: true, Name: "choice.1", Value: "mine", Radio: true, Parts: []w.CheckPart{w.Choice{Source: "Your edit · include", Value: "+7 495 120-45-00"}}}, "s-conflict-1", "tc-check", 1),
	form("rich", check(true, w.Text("Allow"), w.Mono{Text: ".ics"}, w.Text("download and subscription")), "s-publish-2", "tc-check", 1),
	form("box-off", w.Check{Name: "mono", Value: "1", Parts: []w.CheckPart{w.Text("Monospaced dates")}}, "components-38", "tc-check", 2),
	form("box", w.Check{On: true, Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}, "components-38", "tc-check", 1),
	form("tertiary", w.Check{Parts: []w.CheckPart{w.Text("Auto-save")}, Tertiary: true}, "s-notes-20", "tc-check", 1),
	form("rich", w.Check{Name: "choice.1", Value: "none", Radio: true, Parts: []w.CheckPart{w.Choice{Source: "Leave it out", Value: "—"}}}, "s-conflict-1", "tc-check", 2),
	// s-backup-2 .tc-check #4 holds an input inside the button: nested controls
	// (a11y). Its record waits for the owner's decision on the mockup.
	form("rich", w.Check{Name: "tasks.vtodo.not.every.client.understands.them", Value: "1", Parts: []w.CheckPart{w.Text("Tasks "), w.Dim{Parts: []w.Inline{w.Text("· VTODO — not every client understands them")}}}}, "s-davloom-2", "tc-check", 3),
	form("rich", w.Check{On: true, Name: "webdav.folders.mount.as.a.network.drive", Value: "1", Parts: []w.CheckPart{w.Text("WebDAV folders "), w.Dim{Parts: []w.Inline{w.Text("· mount as a network drive")}}}}, "s-davloom-2", "tc-check", 5),
	form("rich", w.Check{On: true, Name: "on.a.clock.of.your.own.cron.gosentry.a.job.in.gitea.local", Value: "1", Parts: []w.CheckPart{w.Text("On a clock of your own — "), w.Mono{Text: "cron"}, w.Text(", GoSentry, a job in Gitea "), w.Badge{Label: "local", Tone: w.BadgeLocal, Title: "Local instance: the address in the line answers only from inside your own network"}}}, "s-backup-2", "tc-check", 5),
	form("rich", check(false, w.Text("Allow"), w.Mono{Text: ".vcf"}, w.Text("download and subscription")), "s-publish-3", "tc-check", 6),
	form("rich", check(false, w.Dim{Parts: []w.Inline{w.Text("Start at sign-in — not available in a portable copy")}}), "desktop-4", "tc-check", 1),

	// tc-range
	form("form", w.Range{From: w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Label: "From", Mono: true}, To: w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-16", Label: "To", Mono: true}, Show: w.Button{Label: "Show"}, Hidden: []w.Hidden{{Name: "view", Value: "range"}}, Action: "/app/calendar"}, "s-agenda-1", "tc-range", 1),
	form("form", w.Range{From: w.Input{Kind: w.InputDate, Name: "from", Value: "2026-01-01", Label: "From", Mono: true}, To: w.Input{Kind: w.InputDate, Name: "to", Value: "2026-11-10", Label: "To", Mono: true}, Show: w.Button{Label: "Show"}, Action: "/app/contacts/nextcloud/kontakty/4f2b-9c31-a0e7/timeline"}, "s-contacts-7", "tc-range", 1),
	form("default", w.Range{From: w.Input{Kind: w.InputDate, Value: "2026-01-01", Mono: true, Label: "From"}, To: w.Input{Value: "today", Mono: true, Label: "To"}, Show: w.Button{Label: "Show"}}, "components-18", "tc-range", 1),
	form("wrap", w.Range{From: w.Input{Kind: w.InputDate, Name: "from", Value: "2026-08-10", Mono: true, Label: "From"}, To: w.Input{Kind: w.InputDate, Name: "to", Value: "2026-08-16", Mono: true, Label: "To"}, Show: w.Button{Label: "Show"}, Wrap: true, Action: "/app/calendar", Hidden: []w.Hidden{{Name: "view", Value: "range"}}}, "s-agenda-1", "tc-range", 2),

	// tc-line, tc-stack, tc-grow, tc-push
	form("plain", w.Line{Items: []w.LinePart{w.Mark{Colour: "#4A6B52", Size: w.MarkH11}, w.Check{On: true, Name: "copy", Value: "nextcloud/dela", Parts: []w.CheckPart{w.Text("nextcloud · Дела")}}, w.Dim{Parts: []w.Inline{w.Text("due 12 Aug 18:00")}, Size: w.Dim11_5}}, Gap: w.LineGap8, Block: true}, "s-tasks-16", "tc-line", 1),
	form("span", w.Line{Items: []w.LinePart{w.Faced{Face: w.FaceLarge, Items: []w.LinePart{w.Text("{name} "), w.Dim{Parts: []w.Inline{w.Text("{size}")}}}}, w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Place: w.Pushed, Removes: true}}, Gap: w.LineGap8}, "s-notes-11", "tc-line", 1),
	form("block", w.Line{Items: []w.LinePart{w.Line{Items: []w.LinePart{w.Faced{Face: w.FaceLarge, Items: []w.LinePart{w.Text("договор-аренды.pdf "), w.Dim{Parts: []w.Inline{w.Text("412 KB")}}}}, w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Place: w.Pushed, Removes: true}}, Gap: w.LineGap8, Place: w.Grown}, w.Line{Items: []w.LinePart{w.Faced{Face: w.FaceLarge, Items: []w.LinePart{w.Text("план-этажа.png "), w.Dim{Parts: []w.Inline{w.Text("2.1 MB")}}}}, w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Place: w.Pushed, Removes: true}}, Gap: w.LineGap8, Place: w.Grown}, w.Line{Items: []w.LinePart{w.Faced{Face: w.FaceLarge, Items: []w.LinePart{w.Text("сканы/ "), w.Dim{Parts: []w.Inline{w.Text("14 files, 38 MB")}}}}, w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Place: w.Pushed, Removes: true}}, Gap: w.LineGap8, Place: w.Grown}}, Gap: w.LineGap8, Block: true}, "s-notes-11", "tc-line", 2),
	form("picked", w.Line{Items: []w.LinePart{w.Faced{Face: w.FaceLarge, Items: []w.LinePart{w.Text("{name} "), w.Dim{Parts: []w.Inline{w.Text("{size}")}}}}, w.Button{Label: "Remove", Kind: w.ButtonDanger, Small: true, Removes: true}}, Gap: w.LineGap8}, "s-agenda-3", "tc-line", 2),
	form("picker", w.Line{Items: []w.LinePart{w.Button{Label: "Add", File: &w.FileField{Form: "", Name: "photo", Accept: "image/*", Label: "Photo", Preview: "photo"}}}, Gap: w.LineGap6}, "s-contacts-4", "tc-line", 1),
	form("span", w.Line{Items: []w.LinePart{w.Button{Label: "New collection", Icon: "plus", Small: true}, w.Button{Label: "Show password", Small: true}, w.Button{Label: "Test", Small: true}, w.Button{Label: "Edit", Small: true}, w.Button{Label: "Disconnect", Kind: w.ButtonDanger, Small: true}}, Gap: w.LineGap6}, "s-profile-1", "tc-line", 1),
	form("span", w.Line{Items: []w.LinePart{w.Button{Label: "Replace", File: &w.FileField{Form: "", Name: "photo", Accept: "image/*", Label: "Photo", Preview: "photo"}}, w.Button{Label: "Crop"}}, Gap: w.LineGap6}, "s-contacts-3", "tc-line", 1),
	form("span", w.Line{Gap: w.LineGap8, Items: []w.LinePart{w.Input{Label: "x"}, btn}}, "components-20", "tc-line", 1),
	form("gap", w.Line{Gap: w.LineGap7, Items: []w.LinePart{btn, w.Kbd{Keys: "Esc"}}}, "s-contacts-2", "tc-line", 1),
	form("block", w.Line{Gap: w.LineGap8, Block: true, Items: []w.LinePart{w.Drop{Trigger: w.Button{Label: "Attach", Icon: "up", Menu: true}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "up", Label: "From this computer"},
		w.MenuItem{Icon: "file", Label: "From Files"}}, Width: w.MenuWidth212, Pop: w.MenuUnder}}, w.Dim{}}}, "s-tasks-3", "tc-line", 1),
	form("block", w.Line{Gap: w.LineGap12, Block: true, Items: []w.LinePart{chk, chk}}, "s-dup-1", "tc-line", 1),
	form("gap", w.Line{Gap: w.LineGap3, Items: []w.LinePart{w.Mark{}, w.Mark{}, w.Mark{}, w.Mark{}}}, "s-davloom-1", "tc-line", 2),
	form("gap", w.Line{Gap: w.LineGap8, Items: []w.LinePart{w.Strong("nextcloud"), w.Dim{}, w.Mono{}, w.Dim{}, w.Mono{}, btn}}, "s-plan-4", "tc-line", 1),
	form("block", w.Line{Gap: w.LineGap8, Block: true, Items: []w.LinePart{btn, btn, w.Dim{Place: w.Pushed}}}, "desktop-3", "tc-line", 2),
	form("block", w.Line{Gap: w.LineGap4, Block: true, Items: []w.LinePart{w.Tag{}, w.Tag{}}}, "s-notes-2", "tc-line", 1),
	form("span", w.Line{Items: []w.LinePart{w.Badge{Label: "badge"}}}, "account-1", "tc-line", 3),
	form("gap", w.Line{Gap: w.LineGap8, Items: []w.LinePart{w.Mark{}, w.Mono{}, w.Text("Синк-ап по релизу")}}, "desktop-8", "tc-line", 1),
	form("block", w.Line{Gap: w.LineGap8, Block: true, Items: []w.LinePart{w.Drop{Trigger: w.Button{Label: "Attach", Icon: "up", Menu: true}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Icon: "up", Label: "From this computer"},
		w.MenuItem{Icon: "file", Label: "From Files"}}, Width: w.MenuWidth212, Pop: w.MenuUnder}},
		w.Faced{Face: w.FaceLarge, Items: []w.LinePart{w.Text("release-checklist.md"), w.Dim{}, w.Badge{Label: "badge"}}}, w.Button{Label: "Remove", Small: true, Kind: w.ButtonDanger}}}, "s-agenda-3", "tc-line", 1),
	form("block", w.Line{Gap: w.LineGap10, Block: true, Items: []w.LinePart{w.Field{}, w.Field{}}}, "s-publish-2", "tc-line", 1),
	form("block", w.Line{Block: true, Gap: w.LineGap7, Items: []w.LinePart{check(true, w.Text("On sign-in, if the last copy is older than")), w.Input{Value: "7", Mono: true, Variant: w.InputChip, Label: "Days"}, w.Faced{Face: w.FaceLarge, Items: []w.LinePart{w.Text("days")}}}}, "s-backup-2", "tc-line", 1),
	form("block", w.Line{Gap: w.LineGap8, Block: true, Items: []w.LinePart{btn, btn, w.Push{Face: faceOf(w.FaceAccent11_5), Items: []w.PushPart{w.Text("Add a second instance")}}}}, "desktop-7", "tc-line", 1),
	form("default", w.Stack{Gap: w.StackGap9, Items: []w.StackPart{btn}}, "components-21", "tc-stack", 1),
	form("default", w.Grow{}, "s-contacts-3", "tc-grow", 1),
	form("default", w.Push{Items: []w.PushPart{btn}}, "components-22", "tc-push", 1),
	form("default", w.Push{Items: []w.PushPart{w.Text("Default")}}, "s-dup-3", "tc-push", 1),

	// tc-rubric
	form("block", w.Rubric{Text: "Rail"}, "shell-1", "tc-rubric", 1),
	form("inline", w.Rubric{Text: "А", Form: w.RubricInline}, "components-12", "tc-rubric", 1),
	form("block", w.Rubric{Text: "nextcloud", Items: []w.RubricPart{w.Mono{Text: "12", Variant: w.MonoTight}}}, "components-11", "tc-rubric", 1),
	form("aside", w.Rubric{Text: "Вера Лебедева", Items: []w.RubricPart{w.RubricAside{Text: "", Items: []w.RubricAsidePart{w.Badge{Label: "candidate ×2"}, w.Mono{Text: "60 points · same phone", Variant: w.MonoTightDusty}}}}}, "s-dup-1", "tc-rubric", 2),
	form("aside", w.Rubric{Lead: &w.Mono{Text: "TEL;work"}, Items: []w.RubricPart{w.RubricAside{Text: "only in your edit"}}}, "s-conflict-1", "tc-rubric", 2),
	form("item", w.Rubric{Text: "Write into", Form: w.RubricItem}, "components-6", "tc-rubric", 1),
	form("variant", w.Rubric{Text: "What is pointing at it now", Variant: w.RubricBottom7}, "s-profile-9", "tc-rubric", 3),
	form("variant", w.Rubric{Text: "DAV servers", Variant: w.RubricSplit, Items: []w.RubricPart{w.Mono{Text: "3", Variant: w.MonoTight}}}, "s-profile-1", "tc-rubric", 2),
	form("variant", w.Rubric{Text: "baikal", Variant: w.RubricDanger, Items: []w.RubricPart{w.Line{}}}, "s-profile-1", "tc-rubric", 6),
	form("variant", w.Rubric{Text: "Что делаем", Variant: w.RubricBottom8}, "s-notes-2", "tc-rubric", 3),
	form("variant", w.Rubric{Text: "nextcloud", Variant: w.RubricRail}, "phone-2", "tc-rubric", 1),
	form("variant", w.Rubric{Text: "Timeline", Variant: w.RubricBottom2}, "s-contacts-7", "tc-rubric", 2),
	form("item", w.Rubric{Text: "mix", Form: w.RubricItem, Variant: w.RubricItemLine, Items: []w.RubricPart{w.Mono{}}}, "account-5", "tc-rubric", 1),
	form("variant", w.Rubric{Text: "Subscriptions", Variant: w.RubricRail, Items: []w.RubricPart{w.Badge{Label: "badge"}}}, "components-11", "tc-rubric", 2),
	form("item", w.Rubric{Text: "Points per signal", Form: w.RubricItem, Items: []w.RubricPart{w.Push{}, w.RubricCol{Text: "Yours"}}}, "s-dup-3", "tc-rubric", 1),
	form("variant", w.Rubric{Text: "What it opens", Variant: w.RubricGap13}, "s-backup-5", "tc-rubric", 2),
	form("row", w.Rubric{Text: "Attached", Variant: w.RubricRow, Items: []w.RubricPart{w.Num{}, w.Line{Place: w.Pushed}}}, "s-mail-1", "tc-rubric", 6),
)

func faceOf(f w.Face) *w.Face { return &f }

func init() {
	Widgets = append(Widgets, w.Seg{}, w.SegFold{}, w.CheckForm{}, w.Check{}, w.Range{}, w.Line{}, w.Stack{}, w.Grow{}, w.Push{}, w.Rubric{})
}
