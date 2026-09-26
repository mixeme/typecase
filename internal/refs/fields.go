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
	form("labelled", field("On", w.FieldAuto, sg), "s-agenda-3", "tc-field", 16),
	form("hint", field("Into", w.FieldLG, pk, hint(t("Any address book you may write to. The UID collisions below are counted against it."))), "s-io-5", "tc-field", 1),
	form("labelled", field("Note", w.FieldFull, area), "s-contacts-3", "tc-field", 16),
	form("labelled", field("Time", w.FieldFix130, in), "s-notes-20", "tc-field", 3),
	form("labelled", field("Start", w.FieldFix96, in), "s-notes-9", "tc-field", 3),
	form("hint", field("Everything lives here", w.FieldAuto, in, hint(t("The same folder on any computer: plug the stick in, run, sign in with your password. The data folder cannot be moved away from the app in this mode — that is what makes it portable."))), "m-desktop-4", "tc-field", 1),
	form("hint", field("Birthday", w.FieldSM, in, hint(t("Type it: 14.03.1987, or 14.03 without the year."))), "components-10", "tc-field", 2),
	form("labelled", field("Display name", w.FieldXL, in), "s-contacts-3", "tc-field", 2),
	form("labelled", field("Phone", w.FieldFill200, in), "s-notes-8", "tc-field", 2),
	form("hint", field("Address", w.FieldLG, ln, hint(t("Used for service notices, and as"), mono("Reply-To"), t("on letters you send.")), ln), "account-1", "tc-field", 2),
	form("hint", field("", w.FieldAuto, chk, hint(t("Untimed events sit on the day, not on the clock."))), "components-10", "tc-field", 6),
	form("labelled", field("On a connected server", w.FieldFill300, ln), "components-20", "tc-field", 1),
	form("labelled", field("Date", w.FieldFix150, in), "s-notes-20", "tc-field", 2),
	form("hint", field("Collection", w.FieldMD, pk, hint(t("Any file collection you can write to."))), "attachments-1", "tc-field", 2),
	form("bare", w.Field{Gap: 7, Parts: []w.FieldPart{chk, chk, chk}}, "appearance-1", "tc-field", 3),
	form("labelled", field("Notebook", w.FieldFlex2_280, pk), "s-notes-20", "tc-field", 1),
	form("labelled", field("Tags", w.FieldFlex3_320, in), "s-notes-20", "tc-field", 4),
	form("labelled", field("Related to", w.FieldFlex3_340, tags), "s-notes-20", "tc-field", 5),
	form("labelled", field(" ", w.FieldFill120, chk), "s-notes-9", "tc-field", 5),
	form("labelled", field("Match threshold", w.FieldInline, in, bt), "s-dup-1", "tc-field", 1),
	form("own-line", w.Field{OwnLine: true, Parts: []w.FieldPart{sg}}, "s-backup-2", "tc-field", 3),
	form("own-line", w.Field{OwnLine: true, Parts: []w.FieldPart{chk, hint(t("On by default. Turning it off makes sense in one case only: the copy goes where nobody but you gets in anyway — and then any archiver opens it."))}}, "s-backup-2", "tc-field", 8),
	form("bare", w.Field{Gap: 8, Parts: []w.FieldPart{chk, w.Line{Block: true, Gap: w.LineGap7, Items: []w.LinePart{w.Check{On: true, Parts: []w.CheckPart{w.Text("On sign-in, if the last copy is older than")}}, w.Input{Value: "7", Mono: true, Variant: w.InputChip, Label: "Days"}, w.Faced{Face: w.FaceLarge, Items: []w.LinePart{w.Text("days")}}}}, chk}}, "s-backup-2", "tc-field", 10),
	form("labelled", w.Field{Label: "Email", LabelAside: dimText("(optional)"), Width: w.FieldMD, Parts: []w.FieldPart{in}}, "s-admin-1", "tc-field", 5),
	form("hint", field("Sender address", w.FieldAuto, sg, hint(t("Your relay decides which of the two it will let out. The envelope sender stays"), mono("noreply@carrel.example"), t("either way, so SPF is unaffected; DKIM and DMARC for this domain are set up on the relay, not here."))), "s-admin-3", "tc-field", 14),
	form("hint", field("Port", w.FieldFill300, in, hint(t("Nothing to set here. The app asks the system for a free port on"), mono("127.0.0.1"), t(", so a busy 8080, a second copy of Carrel or someone else’s server cannot collide with it. The current number is always in the status bar — it is needed only to open the address in a browser."))), "m-desktop-6", "tc-field", 2),
	form("photo", w.Field{Label: "Photo", LabelSpan: true, Width: w.FieldFixAuto, Parts: []w.FieldPart{w.Photo{Side: []w.FieldPart{ln, bt, hintOf(w.HintNarrow, t("EXIF orientation applied, metadata stripped."))}}}}, "s-contacts-3", "tc-field", 1),
	form("bare", field("", w.FieldXS, in), "s-contacts-3", "tc-field", 11),
	form("bare", field("", w.FieldMD, in), "s-contacts-3", "tc-field", 12),
	form("photo", w.Field{Label: "Photo", LabelSpan: true, Width: w.FieldFixAuto, Parts: []w.FieldPart{w.Photo{Initials: "АК", Side: []w.FieldPart{bt, hintOf(w.HintNarrow, t("JPEG, PNG, GIF or WebP."))}}}}, "s-contacts-4", "tc-field", 1),
	form("zoom", field("Zoom", w.FieldW100, w.Zoom{Value: "1.6×"}), "s-contacts-5", "tc-field", 1),
	form("labelled", field(" ", w.FieldAuto, bt), "s-agenda-3", "tc-field", 13),
	form("hint", field("Advanced RRULE", w.FieldXL, in, hint(t("The advanced value takes precedence when set."))), "s-agenda-3", "tc-field", 19),
	form("hint", field("Server URL", w.FieldFull, in, hint(w.Mono{Text: ".well-known", Form: w.MonoCode}, t("is tried first, then the URL exactly as typed — which is the reliable path for Baikal."))), "s-profile-2", "tc-field", 2),
	form("hint", field("Label", w.FieldMD, in, hintOf(w.HintDanger, t("There is already a connection called “nextcloud”. Choose another label: it is how your sources are told apart."))), "s-profile-11", "tc-field", 1),
	form("hint", field("On which server", w.FieldW100, pk, hintOf(w.HintSmall, t("Two of three are listed:"), mono("baikal"), t("is not answering, and until it does there is nothing to create there."))), "s-profile-5", "tc-field", 1),
	form("swatches", field("Colour", w.FieldW100, w.Swatches{Colours: swatchColours, On: 2}, hintOf(w.HintSmall, t("Written to the server as"), mono("calendar-color"), t("— DAVx5 and Thunderbird will show the same one."))), "s-profile-5", "tc-field", 3),
	form("group", field("What goes in it", w.FieldW100, w.Group{Items: []w.FieldPart{chk, chk, chk}},
		hintOf(w.HintSmall, t("Most servers settle this when the collection is made and will not change it afterwards. A collection holding only tasks appears in"), b("Tasks"), t(", not here — the rail follows what is inside, not where it was made."))), "s-profile-5", "tc-field", 4),
	form("hint", field("Address", w.FieldW100, in, hintOf(w.HintSmall, mono("…/remote.php/dav/calendars/mix/poezdki/"), t("— taken from the name, yours to change. Other clients and restore-by-path go by the address; the name they only show."))), "s-profile-5", "tc-field", 5),
	form("swatches", field("Colour", w.FieldW100, w.Swatches{Colours: swatchColours, On: 3}), "s-profile-7", "tc-field", 2),
	form("labelled", w.Field{Label: "Address", LabelAside: dimText("— stays as it is"), Width: w.FieldW100, Parts: []w.FieldPart{in}}, "s-profile-7", "tc-field", 3),
	form("hint", field("Feed address", w.FieldW100, in, hint(mono("https://"), t(","), mono("webcal://"), t("and a secret link from another Carrel are all understood."))), "s-profile-10", "tc-field", 1),
	form("own-line", w.Field{OwnLine: true, Width: w.FieldW100, Parts: []w.FieldPart{bt}}, "s-profile-10", "tc-field", 2),
	form("row", field("Label colour", w.FieldW100, w.FieldRow{Items: []w.FieldPart{w.Mark{}, w.Mark{}, w.Mark{}, w.Mark{}, w.Mark{}, w.Dim{}}},
		hint(t("A feed sends no colour of its own: Carrel derives one, and hatches the strip so a feed is recognisable among your own collections."))), "s-profile-10", "tc-field", 5),
	form("own-line", w.Field{OwnLine: true, Width: w.FieldW100, Parts: []w.FieldPart{chk, hint(t("The login and password are kept the way DAV credentials are: encrypted under your key. A secret link needs no password — the secret is already in the address."))}}, "s-profile-10", "tc-field", 6),
	form("hint", field("", w.FieldAuto, hint(t("A collection that says it takes less ("), mono("max-resource-size"), t(") lowers the boundary for itself:"), mono("nextcloud"), t("2 MB,"), mono("radicale · vps"), t("1 MB. Nothing here raises it."))), "attachments-1", "tc-field", 1),
	form("own-line", w.Field{OwnLine: true, Gap: 7, Parts: []w.FieldPart{chk, hint(mono("2026-07"), t(","), mono("2026-08"), t(". Off means one flat folder — fine for a hundred files, not for ten thousand."))}}, "attachments-1", "tc-field", 4),
	form("hint", field("Written as", w.FieldLG, pk, hint(mono("2026-08-10-sync-board.png"), t(". A name already taken is never overwritten: a counter is added, and the queue says so."))), "attachments-1", "tc-field", 5),
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
	form("hint", field("", w.FieldW100, ln, hint(t("Nothing is written outside this folder: no registry, no start-up entry, no traces in the profile. Moving the folder moves Carrel with it."))), "m-desktop-3", "tc-field", 1),
	form("own-line", w.Field{OwnLine: true, Gap: 7, Parts: []w.FieldPart{chk, hint(t("While this is off, Carrel listens on"), mono("127.0.0.1"), t("only — from outside it does not exist. Turning it on"), b("pins the port"), t("and gives"), mono("http://192.168.1.14:8765"), t("for a phone and for DAVx5, and asks for a password like any other instance: an address that other devices remember cannot change at every start."))}}, "m-desktop-6", "tc-field", 4),
	form("hint", w.Field{Gap: 7, Parts: []w.FieldPart{chk, chk, chk, hint(t("Both are off out of the box, and off means the app does not reach the network at all — not “checks and stays quiet”. The button above works either way. With checking on, an update is downloaded from the project’s releases and verified by signature before anything is replaced; with the second tick off it then waits until the button is pressed, and nothing restarts under your hands mid-work."))}}, "m-desktop-13", "tc-field", 10),
)

func init() {
	Widgets = append(Widgets, w.Field{})
}
