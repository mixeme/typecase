package refs

import w "carrel.dev/typecase/widget"

func shell(state string, fixture Renderer, frame, class string, n int) Entry {
	return e("shell", state, fixture, ref(frame, class, n))
}

var _ = add(
	// tc-brand
	shell("logo", w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, "shell-1", "tc-brand", 1),
	shell("title", w.Brand{Title: "Combined calendar"}, "phone-1", "tc-brand", 1),

	// tc-search
	shell("form", w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…", Name: "q"}, Action: "/app/find"}, "components-43", "tc-search", 1),
	shell("plain", w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}}, "shell-1", "tc-search", 1),
	shell("focus", w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}, Focus: true, Variant: w.SearchFull}, "s-find-2", "tc-search", 1),
	shell("clear", w.Search{Query: w.Query{Value: "ковал", Label: "Search everything…"}, Clear: true}, "s-find-1", "tc-search", 1),
	shell("variant", w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}, Variant: w.SearchFill}, "s-conflict-2", "tc-search", 1),

	// tc-usermenu
	shell("link", w.UserMenu{Name: "mix", Href: "/app/settings/account"}, "components-43", "tc-usermenu", 1),
	shell("plain", w.UserMenu{Name: "mix"}, "shell-1", "tc-usermenu", 1),
	shell("no-arrow", w.UserMenu{Name: "mix", NoArrow: true}, "s-conflict-2", "tc-usermenu", 1),
	shell("focus", w.UserMenu{Name: "mix", Focus: true}, "account-5", "tc-usermenu", 1),

	// tc-emblem
	shell("plain", w.Emblem{Symbol: "davloom"}, "account-1", "tc-emblem", 1),
	shell("titled", w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}, "components-36", "tc-emblem", 1),
	shell("alert", w.Emblem{Symbol: "davloom", Title: "davloom is not responding · last response 2026-08-16 09:12", Alert: true}, "components-36", "tc-emblem", 2),

	// tc-split
	shell("plain", w.Split{Main: w.Button{Label: "New note", Icon: "quill", Kind: w.ButtonPrimary}, More: w.Button{Icon: "chev", Kind: w.ButtonPrimary, Title: "More"}}, "components-9", "tc-split", 1),

	// tc-zone
	shell("plain", w.Zone{Title: w.Rubric{Text: "Rail"}, Note: w.Mono{Text: "264 px · sections + sources", Variant: w.Mono11Top}}, "shell-1", "tc-zone", 1),

	// tc-crumbs
	shell("path", w.Crumbs{Path: []w.Crumb{{Label: "Documents", Href: "#"}, {Label: "carrel", Href: "#"}}, Current: "attachments"}, "components-15", "tc-crumbs", 1),
	shell("aside", w.Crumbs{Current: "Материалы к релизу", Aside: &w.Dim{Parts: []w.Inline{w.Text("· 4 files · 6.6 MB")}}}, "s-publish-12", "tc-crumbs", 1),

	// tc-acts
	shell("row", w.Acts{Items: []w.ActsPart{w.Button{Label: "Edit", Small: true}, w.Button{Label: "All records", Small: true}}}, "components-14", "tc-acts", 1),
	shell("stack", w.Acts{Items: []w.ActsPart{w.Seg{Options: []w.SegOption{{Label: "All 4", On: true}, {Label: "Contacts 1"}, {Label: "Events 1"}, {Label: "Tasks 1"}, {Label: "Notes 1"}}}, w.Field{Label: "Match threshold", Width: w.FieldInline, Parts: []w.FieldPart{w.Input{Value: "60", Mono: true, Variant: w.InputW52, Label: "Match threshold"}, w.Button{Label: "Weights", Small: true}}}}, Stack: true}, "s-dup-1", "tc-acts", 1),

	// tc-nav
	shell("back", w.Nav{Items: []w.NavItem{
		{Icon: "back", Label: "Back to the card", Href: "#"}}}, "s-contacts-3", "tc-nav", 1),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
		{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
		{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
		{Icon: "file", Label: "Files", Href: "#"},
		{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"}}, On: 2}, "s-agenda-1", "tc-nav", 1),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
		{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
		{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
		{Icon: "file", Label: "Files", Href: "#"},
		{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"}}, On: 1}, "s-contacts-1", "tc-nav", 1),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "user", Label: "Account", Href: "#"},
		{Icon: "server", Label: "Connections", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "file", Label: "Attachments", Href: "#"},
		{Icon: "device", Label: "New device", Href: "#"},
		{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
		{Icon: "theme", Label: "Appearance", Href: "#"}}, On: 2}, "s-profile-1", "tc-nav", 2),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
		{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
		{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
		{Icon: "file", Label: "Files", Href: "#"},
		{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"}}, On: 5}, "s-files-1", "tc-nav", 1),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "user", Label: "Account", Href: "#"},
		{Icon: "server", Label: "Connections", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "file", Label: "Attachments", Href: "#"},
		{Icon: "device", Label: "New device", Href: "#"},
		{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
		{Icon: "theme", Label: "Appearance", Href: "#"}}, On: 6}, "s-backup-1", "tc-nav", 2),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
		{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
		{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
		{Icon: "file", Label: "Files", Href: "#"},
		{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"}}, On: 4}, "s-notes-1", "tc-nav", 1),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "con", Label: "Contacts", Href: "#"},
		{Icon: "cal", Label: "Calendar", Href: "#"},
		{Icon: "task", Label: "Tasks", Href: "#"},
		{Icon: "note", Label: "Notes", Href: "#"},
		{Icon: "file", Label: "Files", Href: "#"},
		{Icon: "dup", Label: "Duplicates", Href: "#"}}, On: 2}, "s-notes-6", "tc-nav", 1),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "user", Label: "Account", Href: "#"},
		{Icon: "server", Label: "Connections", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "file", Label: "Attachments", Href: "#"},
		{Icon: "device", Label: "New device", Href: "#"},
		{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
		{Icon: "theme", Label: "Appearance", Href: "#"}}, On: 1}, "account-1", "tc-nav", 2),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "user", Label: "Account", Href: "#"},
		{Icon: "server", Label: "Connections", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "file", Label: "Attachments", Href: "#"},
		{Icon: "device", Label: "New device", Href: "#"},
		{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
		{Icon: "theme", Label: "Appearance", Href: "#"}}, On: 4}, "s-plan-1", "tc-nav", 2),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "con", Label: "Contacts", Href: "#"},
		{Icon: "cal", Label: "Calendar", Href: "#"},
		{Icon: "task", Label: "Tasks", Href: "#"},
		{Icon: "note", Label: "Notes", Href: "#"},
		{Icon: "file", Label: "Files", Href: "#"},
		{Icon: "dup", Label: "Duplicates", Href: "#"}}, On: 6}, "s-dup-2", "tc-nav", 1),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
		{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
		{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
		{Icon: "file", Label: "Files", Href: "#"}}, On: 2}, "s-notes-5", "tc-nav", 1),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "user", Label: "Account", Href: "#"},
		{Icon: "server", Label: "Connections", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "file", Label: "Attachments", Href: "#"},
		{Icon: "device", Label: "New device", Href: "#"},
		{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
		{Icon: "theme", Label: "Appearance", Href: "#"}}, On: 5}, "s-davloom-1", "tc-nav", 2),
	shell("apart", w.Nav{Items: []w.NavItem{
		{Icon: "back", Label: "Back to list", Href: "#"}}, Apart: true}, "s-contacts-7", "tc-nav", 2),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "user", Label: "Account", Href: "#"},
		{Icon: "server", Label: "Connections", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "file", Label: "Attachments", Href: "#"},
		{Icon: "device", Label: "New device", Href: "#"},
		{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
		{Icon: "theme", Label: "Appearance", Href: "#"}}, On: 7}, "s-publish-1", "tc-nav", 2),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
		{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
		{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
		{Icon: "file", Label: "Files", Href: "#"},
		{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"}}, On: 3}, "s-tasks-1", "tc-nav", 1),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "con", Label: "Contacts", Href: "#"},
		{Icon: "cal", Label: "Calendar", Href: "#"},
		{Icon: "task", Label: "Tasks", Href: "#"},
		{Icon: "note", Label: "Notes", Href: "#"},
		{Icon: "file", Label: "Files", Href: "#"},
		{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"}}}, "s-find-1", "tc-nav", 1),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "con", Label: "Contacts", Href: "#"},
		{Icon: "cal", Label: "Calendar", Href: "#"},
		{Icon: "task", Label: "Tasks", Href: "#"},
		{Icon: "note", Label: "Notes", Href: "#"},
		{Icon: "file", Label: "Files", Href: "#"},
		{Icon: "dup", Label: "Duplicates", Count: &w.Num{Value: 3}, Href: "#"}}, On: 6}, "s-dup-1", "tc-nav", 1),
	shell("icons", w.Nav{Items: []w.NavItem{
		{Icon: "back", Href: "#", Title: "Back"}}}, "s-conflict-2", "tc-nav", 1),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "user", Label: "Account", Href: "#"},
		{Icon: "server", Label: "Connections", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "file", Label: "Attachments", Href: "#"},
		{Icon: "device", Label: "New device", Href: "#"},
		{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
		{Icon: "theme", Label: "Appearance", Href: "#"}}, On: 3}, "attachments-1", "tc-nav", 2),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "user", Label: "Account", Href: "#"},
		{Icon: "server", Label: "Connections", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "file", Label: "Attachments", Href: "#"},
		{Icon: "device", Label: "New device", Href: "#"},
		{Emblem: &w.Emblem{Symbol: "davloom"}, Label: "davloom", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "archive", Label: "Backup", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "link", Label: "Published links", Count: &w.Num{Value: 2}, Href: "#"},
		{Icon: "theme", Label: "Appearance", Href: "#"}}, On: 8}, "appearance-1", "tc-nav", 2),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "con", Label: "Contacts", Count: &w.Num{Value: 312}, Href: "#"},
		{Icon: "cal", Label: "Calendar", Count: &w.Num{Value: 14}, Href: "#"},
		{Icon: "task", Label: "Tasks", Count: &w.Num{Value: 3}, Href: "#"},
		{Icon: "note", Label: "Notes", Count: &w.Num{Value: 86}, Href: "#"},
		{Icon: "file", Label: "Files", Href: "#"}}, On: 1}, "s-mail-1", "tc-nav", 1),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
		{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
		{Icon: "adm", Label: "Settings", Href: "#"},
		{Icon: "shield", Label: "Installation", Href: "#"},
		{Icon: "search", Label: "DAV validator", Href: "#"},
		{Icon: "dup", Label: "Key escrow", Href: "#"},
		{Icon: "clock", Label: "Audit log", Href: "#"}}, On: 1}, "s-admin-1", "tc-nav", 2),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
		{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
		{Icon: "adm", Label: "Settings", Href: "#"},
		{Icon: "shield", Label: "Installation", Href: "#"},
		{Icon: "search", Label: "DAV validator", Href: "#"},
		{Icon: "dup", Label: "Key escrow", Href: "#"},
		{Icon: "clock", Label: "Audit log", Href: "#"}}, On: 2}, "s-admin-2", "tc-nav", 2),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
		{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
		{Icon: "adm", Label: "Settings", Href: "#"},
		{Icon: "shield", Label: "Installation", Href: "#"},
		{Icon: "search", Label: "DAV validator", Href: "#"},
		{Icon: "dup", Label: "Key escrow", Href: "#"},
		{Icon: "clock", Label: "Audit log", Href: "#"}}, On: 3}, "s-admin-3", "tc-nav", 2),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
		{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
		{Icon: "adm", Label: "Settings", Href: "#"},
		{Icon: "shield", Label: "Installation", Href: "#"},
		{Icon: "search", Label: "DAV validator", Href: "#"},
		{Icon: "shield", Label: "Key escrow", Href: "#"},
		{Icon: "clock", Label: "Audit log", Href: "#"}}, On: 4}, "s-admin-4", "tc-nav", 2),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
		{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
		{Icon: "adm", Label: "Settings", Href: "#"},
		{Icon: "shield", Label: "Installation", Href: "#"},
		{Icon: "search", Label: "DAV validator", Href: "#"},
		{Icon: "dup", Label: "Key escrow", Href: "#"},
		{Icon: "clock", Label: "Audit log", Href: "#"}}, On: 5}, "s-admin-5", "tc-nav", 2),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
		{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
		{Icon: "adm", Label: "Settings", Href: "#"},
		{Icon: "shield", Label: "Installation", Href: "#"},
		{Icon: "search", Label: "DAV validator", Href: "#"},
		{Icon: "shield", Label: "Key escrow", Href: "#"},
		{Icon: "clock", Label: "Audit log", Href: "#"}}, On: 6}, "s-admin-6", "tc-nav", 2),
	shell("plain", w.Nav{Items: []w.NavItem{
		{Icon: "plus", Label: "Add users", Count: &w.Num{Value: 2}, Href: "#"},
		{Icon: "con", Label: "Users", Count: &w.Num{Value: 6}, Href: "#"},
		{Icon: "adm", Label: "Settings", Href: "#"},
		{Icon: "shield", Label: "Installation", Href: "#"},
		{Icon: "search", Label: "DAV validator", Href: "#"},
		{Icon: "dup", Label: "Key escrow", Href: "#"},
		{Icon: "clock", Label: "Audit log", Href: "#"}}, On: 7}, "s-admin-7", "tc-nav", 2),

	// tc-railsec
	shell("plain", w.RailSec{Items: []w.RailSecPart{
		w.TickRow{Ticked: w.TickedPart, Kind: w.TickRowAll, Mark: w.Mark{Kind: w.MarkMulti}, Label: "Combined calendar", Meta: "169", Current: true}}}, "components-11", "tc-railsec", 1),
	shell("apart", w.RailSec{Items: []w.RailSecPart{
		w.Rubric{Text: "Subscriptions"},
		w.TickRow{Kind: w.TickRowExternal, Mark: w.Mark{Colour: "#8C6239"}, Label: "Праздники РФ", Meta: "ro"}}, Apart: true}, "components-11", "tc-railsec", 3),
	shell("tags", w.RailSec{Items: []w.RailSecPart{
		w.Rubric{Text: "Tags"}}, Tags: []w.Tag{w.Tag{Label: "all", On: true}, w.Tag{Label: "работа"}, w.Tag{Label: "издательство"}}}, "s-contacts-1", "tc-railsec", 4),
	shell("empty", w.RailSec{Items: []w.RailSecPart{
		w.Rubric{Text: "Sources"}}, Note: "Nothing connected yet"}, "states-1", "tc-railsec", 1),

	// tc-railstatus
	shell("running", w.RailStatus{Text: "davloom · running · 3 devices", Detail: w.Dim{Parts: []w.Inline{w.Text("Last response "), w.Mono{Text: "2026-08-16 09:41"}}, Size: w.Dim11}}, "components-35", "tc-railstatus", 1),
	shell("alert", w.RailStatus{Tone: w.RailStatusAlert, Text: "davloom is not responding", Detail: w.Dim{Parts: []w.Inline{w.Text("Last response "), w.Mono{Text: "2026-08-16 09:12"}}, Size: w.Dim11}}, "components-35", "tc-railstatus", 3),
	shell("off", w.RailStatus{Tone: w.RailStatusOff, Text: "davloom is off", Detail: w.Dim{Parts: []w.Inline{w.Text("No device connects until it is started.")}, Size: w.Dim11}}, "components-35", "tc-railstatus", 2),

	// tc-railfoot
	shell("plain", w.RailFoot{Items: []w.RailFootPart{w.Button{Icon: "plus", Title: "New address book"}, w.Button{Icon: "server", Title: "Connect account"}, w.Button{Icon: "link", Title: "Subscribe to a feed"}, w.Emblem{Symbol: "davloom", Title: "davloom · running · 3 devices"}}, About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"}}, "s-contacts-1", "tc-railfoot", 1),
	shell("strip", w.RailFoot{Items: []w.RailFootPart{w.RailStatus{Text: "davloom · running · 3 devices", Detail: w.Dim{Parts: []w.Inline{w.Text("Last response "), w.Mono{Text: "2026-08-16 09:41"}}, Size: w.Dim11}}}, Strip: true}, "components-35", "tc-railfoot", 1),
	shell("about", w.RailFoot{About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"}}, "s-dup-1", "tc-railfoot", 1),
	shell("strip", w.RailFoot{Items: []w.RailFootPart{w.Button{Icon: "plus", Title: "New notebook"}, w.Button{Icon: "server", Title: "Connect account"}, w.Button{Icon: "link", Title: "Subscribe to a feed"}}, About: &w.About{Name: "Carrel", Build: w.Mono{Text: "0.0.1 · 786949a"}, Href: "#rest"}, Strip: true}, "components-36", "tc-railfoot", 1),

	// tc-top
	shell("plain", w.Top{Brand: w.Brand{Logo: &w.Logo{Label: "Carrel", ViewBox: "4.79 -31.36 291.08 116.63"}}, Search: &w.Search{Query: w.Query{Value: "Search everything…", Label: "Search everything…"}}, Actions: []w.TopAction{w.Button{Label: "New note", Menu: true, Icon: "quill", Kind: w.ButtonPrimary}, w.Button{Label: "2 min ago", Icon: "refresh"}, w.UserMenu{Name: "mix"}}}, "shell-1", "tc-top", 1),
	shell("phone", w.Top{Lead: &w.Button{Icon: "menu", IconBox: w.IconBox19, Variant: w.ButtonPadX4, Title: "More"}, Brand: w.Brand{Title: "Combined calendar"}, Actions: []w.TopAction{w.Button{Icon: "search", IconBox: w.IconBox17, Variant: w.ButtonPadX6, Title: "More"}, w.Button{Icon: "quill", IconBox: w.IconBox16, Kind: w.ButtonPrimary, Variant: w.ButtonPadX8, Title: "More"}}, Tight: true}, "phone-1", "tc-top", 1),
)

func init() {
	Widgets = append(Widgets, w.Brand{}, w.Search{}, w.UserMenu{}, w.Emblem{}, w.Split{}, w.Zone{}, w.Crumbs{}, w.Acts{},
		w.Nav{}, w.RailSec{}, w.RailStatus{}, w.RailFoot{}, w.Top{})
	Parts = append(Parts, w.Logo{}, w.Crumb{}, w.NavItem{}, w.About{})
}
