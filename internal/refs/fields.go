package refs

import w "carrel.dev/typecase/widget"

func field(label string, width w.FieldWidth, parts ...w.FieldPart) w.Field {
	return w.Field{Label: label, Width: width, Parts: parts}
}

func t(s string) w.Text       { return w.Text(s) }
func mono(s string) w.Mono    { return w.Mono{Text: s} }
func b(s string) w.Strong     { return w.Strong(s) }
func dimText(s string) *w.Dim { return &w.Dim{Parts: []w.Inline{w.Text(s)}} }

var (
	in    = w.Input{Label: "x"}
	pk    = w.Pick{Label: "x"}
	chk   = w.Check{Parts: []w.CheckPart{w.Text("Option")}}
	sg    = w.Seg{}
	ln    = w.Line{}
	bt    = w.Button{Label: "x"}
	area  = w.Area{Label: "x"}
	tags  = w.TagInput{}
	dimBt = w.Dim{}
)

func hintOf(v w.HintVariant, parts ...w.Inline) w.FieldHint {
	return w.FieldHint{Variant: v, Parts: parts}
}

var swatchColours = []string{"#4A6B52", "#3F6E8C", "#6B4E7A", "#A05A2C", "#8C6239", "#9A9280"}

var _ = add(
	form("labelled", field("Display name", w.FieldMD, in), "components-10", "tc-field", 1),
	form("labelled", field("Status", w.FieldSM, pk), "components-10", "tc-field", 5),
	form("labelled", field("Rotate", w.FieldW100, sg), "s-contacts-5", "tc-field", 2),
	form("labelled", field("Start time", w.FieldXS, in), "components-10", "tc-field", 3),
	form("bare", field("", w.FieldAuto, chk), "components-10", "tc-field", 7),
	form("labelled", field("Email", w.FieldLG, in), "components-10", "tc-field", 4),
	form("hint", field("", w.FieldAuto, hint(t("Changing your password re-wraps your data key only; saved connections stay readable."))), "account-1", "tc-field", 3),
	form("hint", field("Saved as", w.FieldW100, in, hint(t("The photo lives inside the card itself, so its size is the card’s size: a square of 512 px at quality 85 keeps a vCard near 200 KB, which servers and phones still accept. Lists are served a smaller copy of the same crop — the original file is not kept anywhere."))), "s-contacts-5", "tc-field", 3),
	form("photo-empty", w.Field{Label: "Photo", LabelSpan: true, Parts: []w.FieldPart{w.Photo{Side: []w.FieldPart{w.Line{Items: []w.LinePart{w.Button{Label: "Add", File: &w.FileField{Form: "post-photo", Name: "photo", Accept: "image/*", Label: "Photo"}}, w.Button{Label: "Upload", Form: "post-photo", NoScript: true}}, Gap: w.LineGap6}, w.FieldHint{Parts: []w.Inline{w.Text("JPEG, PNG, GIF or WebP.")}, Variant: w.HintNarrow}}}}, Width: w.FieldFixAuto}, "s-contacts-27", "tc-field", 1),
	form("hint-only", w.Field{Parts: []w.FieldPart{w.Pick{Value: "urn:uuid:9c1e-07f2-b4a1"}, w.FieldHint{Parts: []w.Inline{w.Text("Not in the books you look in.")}}}, Width: w.FieldLG}, "mentions-7", "tc-field", 4),
	form("checks-plain", w.Field{Parts: []w.FieldPart{w.Check{On: true, Parts: []w.CheckPart{w.Text("Start at sign-in")}}, w.Check{On: true, Parts: []w.CheckPart{w.Text("Closing the window hides to the tray")}}}, Gap: 8}, "desktop-13", "tc-field", 2),
	form("photo-image", w.Field{Label: "Photo", LabelSpan: true, Parts: []w.FieldPart{w.Photo{Src: "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 40 40'%3E%3Crect width='40' height='40' fill='%238C7B6A'/%3E%3Ccircle cx='20' cy='16' r='7' fill='%23E8DFCB'/%3E%3Cpath d='M6 40c2-9 8-13 14-13s12 4 14 13z' fill='%23E8DFCB'/%3E%3C/svg%3E", Side: []w.FieldPart{w.Button{Label: "Replace"}}}}, Width: w.FieldFixAuto}, "components-40", "tc-field", 1),
	form("checks", w.Field{Parts: []w.FieldPart{w.Check{On: true, Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}, w.Check{Name: "mono", Value: "1", Parts: []w.CheckPart{w.Text("Monospaced dates")}}}, Gap: 7, OwnLine: true}, "components-38", "tc-field", 4),
	form("full", w.Field{Label: "Note", Parts: []w.FieldPart{w.Area{Text: "Пишет по вторникам, отвечает в тот же день.", Label: "Text"}}, Width: w.FieldFull}, "s-contacts-3", "tc-field", 22),
	form("labelled-control", w.Field{Label: " ", Parts: []w.FieldPart{w.Button{Label: "Add reminder"}}}, "s-agenda-3", "tc-field", 14),
	form("labelled-plain", w.Field{Label: "On", Parts: []w.FieldPart{w.Seg{Options: []w.SegOption{{Label: "MO", On: true}, {Label: "TU"}, {Label: "WE"}, {Label: "TH"}, {Label: "FR"}, {Label: "SA"}, {Label: "SU"}}}}}, "s-agenda-3", "tc-field", 18),
	form("hint-dusty", w.Field{Label: "Name", Parts: []w.FieldPart{w.Input{Value: "Scans", Label: "Name", Focus: true}, w.FieldHint{Parts: []w.Inline{w.Text("A folder's name is its address: there is nothing else to ask.")}, Variant: w.HintSmall}}, Width: w.FieldW100}, "s-profile-20", "tc-field", 2),
	form("hint-xl", w.Field{Label: "Advanced RRULE", Parts: []w.FieldPart{w.Input{Placeholder: "FREQ=WEEKLY;BYDAY=MO;UNTIL=20261228", Label: "Advanced RRULE", Mono: true}, w.FieldHint{Parts: []w.Inline{w.Text("The advanced value takes precedence when set.")}}}, Width: w.FieldXL}, "s-agenda-3", "tc-field", 21),
	form("error-aside", w.Field{Label: "Email", LabelAside: w.Dim{Parts: []w.Inline{w.Text("(optional)")}}, Parts: []w.FieldPart{w.Input{Value: "viktor@", Label: "Email", Error: true}, w.FieldHint{Parts: []w.Inline{w.Text("Enter one address, like name@example.org.")}, Variant: w.HintDanger}}, Width: w.FieldMD}, "s-admin-19", "tc-field", 2),
	form("labelled", w.Field{Label: "Every", Parts: []w.FieldPart{w.Input{Value: "1", Label: "Every", Mono: true}}, Width: w.FieldXS}, "s-agenda-3", "tc-field", 17),
	form("hint", field("Into", w.FieldLG, pk, hint(t("Any address book you may write to. The UID collisions below are counted against it."))), "s-io-5", "tc-field", 1),
	form("labelled", w.Field{Label: "Street", Parts: []w.FieldPart{w.Input{Value: "ул. Садовая, 12, кв. 5", Label: "Street"}}, Width: w.FieldLG}, "s-contacts-3", "tc-field", 16),
	form("labelled", field("Time", w.FieldFix130, in), "s-notes-20", "tc-field", 3),
	form("labelled", field("Start", w.FieldFix96, in), "s-notes-9", "tc-field", 3),
	form("hint", field("Everything lives here", w.FieldAuto, in, hint(t("The same folder on any computer: plug the stick in, run, sign in with your password. The data folder cannot be moved away from the app in this mode — that is what makes it portable."))), "desktop-4", "tc-field", 1),
	form("hint", field("Birthday", w.FieldSM, in, hint(t("Type it: 14.03.1987, or 14.03 without the year."))), "components-10", "tc-field", 2),
	form("labelled", field("Display name", w.FieldXL, in), "s-contacts-3", "tc-field", 2),
	form("labelled", field("Phone", w.FieldFill200, in), "s-notes-8", "tc-field", 2),
	form("hint", field("Address", w.FieldLG, ln, hint(t("Used for service notices, and as"), mono("Reply-To"), t("on letters you send.")), ln), "account-1", "tc-field", 2),
	form("hint", field("", w.FieldAuto, chk, hint(t("Untimed events sit on the day, not on the clock."))), "components-10", "tc-field", 6),
	form("labelled", field("On a connected server", w.FieldFill300, ln), "components-20", "tc-field", 1),
	form("labelled", field("Date", w.FieldFix150, in), "s-notes-20", "tc-field", 2),
	form("hint", field("Collection", w.FieldMD, pk, hint(t("Any file collection you can write to."))), "attachments-1", "tc-field", 2),
	form("error", w.Field{Label: "Temporary password", Parts: []w.FieldPart{w.Input{Value: "4c1", Label: "Temporary password", Mono: true, Focus: true, Error: true}, w.FieldHint{Parts: []w.Inline{w.Text("Password must be at least 8 characters.")}, Variant: w.HintDanger}}, Width: w.FieldW100}, "s-admin-16", "tc-field", 1),
	form("labelled", w.Field{Label: "Certificate fingerprint", LabelAside: w.Dim{Parts: []w.Inline{w.Text("SHA-256 · optional")}}, Parts: []w.FieldPart{w.Input{Placeholder: "only for a certificate the system does not trust", Label: "Certificate fingerprint", Mono: true}}, Width: w.FieldFull}, "s-profile-2", "tc-field", 5),
	form("hint", w.Field{Label: "Birthday", Parts: []w.FieldPart{w.Input{Value: "31.02.1987", Label: "Birthday", Mono: true, Focus: true, Error: true, Numeric: true}, w.FieldHint{Parts: []w.Inline{w.Text("Type the date as 14.03.1987, or 14.03 without the year.")}, Variant: w.HintDanger}}, Width: w.FieldSM}, "s-contacts-18", "tc-field", 1),
	form("bare", w.Field{Gap: 7, Parts: []w.FieldPart{chk, chk, chk}}, "appearance-1", "tc-field", 3),
	form("labelled", w.Field{Label: "Notebook", Parts: []w.FieldPart{w.Drop{Trigger: w.Pick{Value: "nextcloud · Записи", Mark: &w.Mark{Colour: "#4A6B52"}}, Menu: w.Menu{Items: []w.MenuPart{
		w.MenuItem{Mark: &w.Mark{Colour: "#4A6B52"}, Label: "nextcloud · Записи", On: true, Name: "notebook", Value: "nextcloud/zapisi"},
		w.MenuItem{Mark: &w.Mark{Colour: "#6B4E7A"}, Label: "nextcloud · Дневник", Name: "notebook", Value: "nextcloud/dnevnik"},
		w.MenuItem{Mark: &w.Mark{Colour: "#3F6E8C"}, Trail: w.Badge{Label: "ro", Place: w.Pushed}, Label: "nextcloud · Общие с семьёй", Reason: "Read-only: this account cannot write here", Name: "notebook", Value: "nextcloud/semya"}}, Width: w.MenuWidth290, Pop: w.MenuUnder}}}, Width: w.FieldFlex2_280}, "s-notes-20", "tc-field", 1),
	form("labelled", field("Tags", w.FieldFlex3_320, in), "s-notes-20", "tc-field", 4),
	form("labelled", field("Related to", w.FieldFlex3_340, tags), "s-notes-20", "tc-field", 5),
	form("labelled", field(" ", w.FieldFill120, chk), "s-notes-9", "tc-field", 5),
	form("labelled", w.Field{Label: "Match threshold", Parts: []w.FieldPart{w.Input{Name: "threshold", Value: "60", Label: "Match threshold", Mono: true, Variant: w.InputW52}, w.Button{Label: "Weights", Small: true}}, Width: w.FieldInline, Action: "/app/duplicates"}, "s-dup-1", "tc-field", 1),
	form("own-line", w.Field{OwnLine: true, Parts: []w.FieldPart{sg}}, "s-backup-2", "tc-field", 3),
	form("own-line", w.Field{OwnLine: true, Parts: []w.FieldPart{chk, hint(t("On by default. Turning it off makes sense in one case only: the copy goes where nobody but you gets in anyway — and then any archiver opens it."))}}, "s-backup-2", "tc-field", 8),
	form("bare", w.Field{Parts: []w.FieldPart{w.Check{On: true, Parts: []w.CheckPart{w.Text("On demand — the button in the job list")}}, w.Line{Items: []w.LinePart{w.Check{On: true, Parts: []w.CheckPart{w.Text("On sign-in, if the last copy is older than")}}, w.Input{Value: "7", Label: "Days", Mono: true, Variant: w.InputChip}, w.Faced{Face: w.FaceLarge, Items: []w.LinePart{w.Text("days")}}}, Gap: w.LineGap7, Block: true}, w.Check{On: true, Parts: []w.CheckPart{w.Text("On a clock of your own — "), w.Mono{Text: "cron"}, w.Text(", GoSentry, a job in Gitea "), w.Badge{Label: "local", Tone: w.BadgeLocal, Title: "Local instance: the address in the line answers only from inside your own network"}}}, w.FieldHint{Parts: []w.Inline{w.Text("Turning it on stores this job's data key under the server's key — less protection, for convenience: the job can then read without you signed in.")}}}, Gap: 8}, "s-backup-2", "tc-field", 10),
	form("labelled", w.Field{Label: "Email", LabelAside: dimText("(optional)"), Width: w.FieldMD, Parts: []w.FieldPart{in}}, "s-admin-1", "tc-field", 5),
	form("hint", field("Sender address", w.FieldAuto, sg, hint(t("Your relay decides which of the two it will let out. The envelope sender stays"), mono("noreply@carrel.example"), t("either way, so SPF is unaffected; DKIM and DMARC for this domain are set up on the relay, not here."))), "s-admin-3", "tc-field", 14),
	form("hint", field("Port", w.FieldFill300, in, hint(t("Nothing to set here. The app asks the system for a free port on"), mono("127.0.0.1"), t(", so a busy 8080, a second copy of Carrel or someone else’s server cannot collide with it. The current number is always in the status bar — it is needed only to open the address in a browser."))), "desktop-6", "tc-field", 2),
	form("photo", w.Field{Label: "Photo", LabelSpan: true, Width: w.FieldFixAuto, Parts: []w.FieldPart{w.Photo{Side: []w.FieldPart{ln, bt, hintOf(w.HintNarrow, t("EXIF orientation applied, metadata stripped."))}}}}, "s-contacts-3", "tc-field", 1),
	form("bare", field("", w.FieldXS, in), "s-contacts-3", "tc-field", 11),
	form("bare", field("", w.FieldMD, in), "s-contacts-3", "tc-field", 12),
	form("photo", w.Field{Label: "Photo", LabelSpan: true, Width: w.FieldFixAuto, Parts: []w.FieldPart{w.Photo{Initials: "АК", Side: []w.FieldPart{w.Line{Items: []w.LinePart{w.Button{Label: "Add", File: &w.FileField{Form: "post-photo", Name: "photo", Accept: "image/*", Label: "Photo"}}, w.Button{Label: "Upload", Form: "post-photo", NoScript: true}}, Gap: w.LineGap6}, hintOf(w.HintNarrow, t("JPEG, PNG, GIF or WebP."))}}}}, "s-contacts-4", "tc-field", 1),
	form("zoom", field("Zoom", w.FieldW100, w.Zoom{Value: "1.6×"}), "s-contacts-5", "tc-field", 1),
	form("labelled", w.Field{Label: "Reminder", Parts: []w.FieldPart{w.Pick{Value: "1 day before"}}, Width: w.FieldMD}, "s-agenda-3", "tc-field", 13),
	form("hint", w.Field{Label: "Count", Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Count"}}, Width: w.FieldXS}, "s-agenda-3", "tc-field", 20),
	form("hint", w.Field{Parts: []w.FieldPart{w.Drop{Trigger: w.TagInput{}, Menu: w.Menu{Width: w.MenuWidth420, Pop: w.MenuUnder}}, w.FieldHint{Parts: []w.Inline{w.Text("A name or an address: the ticked address books are searched. On save, nextcloud sends the invitations and their changes — Carrel sends no mail itself.")}}}, Width: w.FieldFull}, "s-agenda-3", "tc-field", 15),
	form("bare", w.Field{Parts: []w.FieldPart{w.Drop{Open: true, Trigger: w.TagInput{}, Menu: w.Menu{Width: w.MenuWidth420}}}, Width: w.FieldFull}, "s-agenda-12", "tc-field", 1),
	form("hint", field("Server URL", w.FieldFull, in, hint(w.Mono{Text: ".well-known", Form: w.MonoCode}, t("is tried first, then the URL exactly as typed — which is the reliable path for Baikal."))), "s-profile-2", "tc-field", 2),
	form("hint", field("Label", w.FieldMD, in, hintOf(w.HintDanger, t("There is already a connection called “nextcloud”. Choose another label: it is how your sources are told apart."))), "s-profile-11", "tc-field", 1),
	form("hint", field("On which server", w.FieldW100, pk, hintOf(w.HintSmall, t("Two of three are listed:"), mono("baikal"), t("is not answering, and until it does there is nothing to create there."))), "s-profile-5", "tc-field", 1),
	form("swatches", w.Field{Label: "Colour", Parts: []w.FieldPart{w.Swatches{Colours: []string{"#4A6B52", "#3F6E8C", "#6B4E7A", "#A05A2C", "#8C6239", "#9A9280"}, On: 2, Name: "colour"}, w.FieldHint{Parts: []w.Inline{w.Text("Written to the server as "), w.Mono{Text: "calendar-color"}, w.Text(" — DAVx5 and Thunderbird will show the same one.")}, Variant: w.HintSmall}}, Width: w.FieldW100}, "s-profile-5", "tc-field", 3),
	form("group", field("What goes in it", w.FieldW100, w.Group{Items: []w.FieldPart{chk, chk, chk}},
		hintOf(w.HintSmall, t("Most servers settle this when the collection is made and will not change it afterwards. A collection holding only tasks appears in"), b("Tasks"), t(", not here — the rail follows what is inside, not where it was made."))), "s-profile-5", "tc-field", 4),
	form("hint", field("Address", w.FieldW100, in, hintOf(w.HintSmall, mono("…/remote.php/dav/calendars/mix/poezdki/"), t("— taken from the name, yours to change. Other clients and restore-by-path go by the address; the name they only show."))), "s-profile-5", "tc-field", 5),
	form("swatches", w.Field{Label: "Colour", Parts: []w.FieldPart{w.Swatches{Colours: []string{"#4A6B52", "#3F6E8C", "#6B4E7A", "#A05A2C", "#8C6239", "#9A9280"}, On: 3, Name: "colour"}}, Width: w.FieldW100}, "s-profile-7", "tc-field", 2),
	form("labelled", w.Field{Label: "Address", LabelAside: dimText("— stays as it is"), Width: w.FieldW100, Parts: []w.FieldPart{in}}, "s-profile-7", "tc-field", 3),
	form("hint", field("Feed address", w.FieldW100, in, hint(mono("https://"), t(","), mono("webcal://"), t("and a secret link from another Carrel are all understood."))), "s-profile-10", "tc-field", 1),
	form("own-line", w.Field{OwnLine: true, Width: w.FieldW100, Parts: []w.FieldPart{bt}}, "s-profile-10", "tc-field", 2),
	form("row", field("Label colour", w.FieldW100, w.FieldRow{Items: []w.FieldPart{w.Mark{}, w.Mark{}, w.Mark{}, w.Mark{}, w.Mark{}, w.Dim{}}},
		hint(t("A feed sends no colour of its own: Carrel derives one, and hatches the strip so a feed is recognisable among your own collections."))), "s-profile-10", "tc-field", 5),
	form("own-line", w.Field{OwnLine: true, Width: w.FieldW100, Parts: []w.FieldPart{chk, hint(t("The login and password are kept the way DAV credentials are: encrypted under your key. A secret link needs no password — the secret is already in the address."))}}, "s-profile-10", "tc-field", 6),
	form("hint", field("", w.FieldAuto, hint(t("A collection that says it takes less ("), mono("max-resource-size"), t(") lowers the boundary for itself:"), mono("nextcloud"), t("2 MB,"), mono("radicale · vps"), t("1 MB. Nothing here raises it."))), "attachments-1", "tc-field", 1),
	form("own-line", w.Field{OwnLine: true, Gap: 7, Parts: []w.FieldPart{chk, hint(mono("2026-07"), t(","), mono("2026-08"), t(". Off means one flat folder — fine for a hundred files, not for ten thousand."))}}, "attachments-1", "tc-field", 4),
	form("hint", w.Field{Label: "Written as", Parts: []w.FieldPart{w.Pick{Value: "Date, then the original name"}, w.FieldHint{Parts: []w.Inline{w.Mono{Text: "2026-08-10 sync-board.png"}, w.Text(". A name already taken is never overwritten: a counter is added.")}}}, Width: w.FieldLG}, "attachments-1", "tc-field", 5),
	form("hint", field("Address for the client", w.FieldMax640Fill460, ln, hint(t("In DAVx5 choose “Login with URL and user name”: the user name is your Carrel one, the password is the device’s, not yours."))), "s-davloom-1", "tc-field", 1),
	form("group", w.Field{Label: "What the device will see", LabelAside: dimText("— changed later too, without re-issuing the password"), Width: w.FieldW100,
		Parts: []w.FieldPart{w.Group{Form: w.GroupColumn, Items: []w.FieldPart{chk, chk, chk, chk, chk}}}}, "s-davloom-2", "tc-field", 3),
	form("group", field("Collections", w.FieldW100, w.Group{Form: w.GroupColumn, Items: []w.FieldPart{chk, chk, chk, chk, chk, chk}}), "s-davloom-2", "tc-field", 4),
	form("hint", field("", w.FieldAuto, hint(t("Copies past the eighth are deleted"), b("after"), t("the new one has been read back and checked, not before. The file is dated and never overwritten."))), "s-backup-2", "tc-field", 7),
	form("own-line", w.Field{Label: "Where the archive is", OwnLine: true, Parts: []w.FieldPart{sg, hint(t("One archive, one of three. The other two stay empty until you switch."))}}, "s-backup-7", "tc-field", 1),
	form("bare", field("", w.FieldLG, in), "s-backup-7", "tc-field", 2),
	form("group", field("What the visitor sees", w.FieldW100, w.Group{Form: w.GroupColumnLoose, Items: []w.FieldPart{chk, chk, chk, chk, chk}}), "s-publish-3", "tc-field", 1),
	form("group", w.Field{Label: "Attachments", LabelAside: w.Plain{Text: "2 files · 3.2 MB"}, Width: w.FieldW100, Parts: []w.FieldPart{w.Group{Form: w.GroupColumnLoose, Items: []w.FieldPart{chk,
		hint(t("The links are created together with this one, expire with it and are revoked with it. Without this the page says how many files there are, but not their names: a file name gives away more than it looks."))}}}}, "s-publish-4", "tc-field", 1),
	form("hint", field("Send it as", w.FieldW100, sg, hint(t("Markdown is what a person opens and reads; it carries the text and links to attachments, not the foreign properties another client wrote. The source"), mono(".ics"), t("is the faithful copy and needs a client that reads VJOURNAL — jtx Board, Evolution."))), "s-mail-5", "tc-field", 1),
	form("bare", field("", w.FieldW280, in), "s-admin-4", "tc-field", 1),
	form("hint", field("", w.FieldW100, ln, hint(t("Nothing is written outside this folder: no registry, no start-up entry, no traces in the profile. Moving the folder moves Carrel with it."))), "desktop-3", "tc-field", 1),
	form("own-line", w.Field{OwnLine: true, Gap: 7, Parts: []w.FieldPart{chk, hint(t("While this is off, Carrel listens on"), mono("127.0.0.1"), t("only — from outside it does not exist. Turning it on"), b("pins the port"), t("and gives"), mono("http://192.168.1.14:8765"), t("for a phone and for DAVx5, and asks for a password like any other instance: an address that other devices remember cannot change at every start."))}}, "desktop-6", "tc-field", 4),
	form("hint", w.Field{Gap: 7, Parts: []w.FieldPart{chk, chk, chk, hint(t("Both are off out of the box, and off means the app does not reach the network at all — not “checks and stays quiet”. The button above works either way. With checking on, an update is downloaded from the project’s releases and verified by signature before anything is replaced; with the second tick off it then waits until the button is pressed, and nothing restarts under your hands mid-work."))}}, "desktop-13", "tc-field", 10),
)

func init() {
	Widgets = append(Widgets, w.Field{})
}
