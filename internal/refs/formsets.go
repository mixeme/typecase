package refs

import w "carrel.dev/typecase/widget"

var (
	rub   = w.Rubric{Text: "x"}
	fld   = w.Field{}
	grown = w.Field{Place: w.Grown}
	fs    = w.FieldSet{}
)

func fsItems(items ...w.FieldSetPart) []w.FieldSetPart { return items }
func formItems(items ...w.FormPart) []w.FormPart       { return items }
func footItems(items ...w.FormFootPart) []w.FormFootPart {
	return items
}

// Forms whose children come from later families (message, card, table,
// fields, tick row, diagnostics) get their records with those families.
var _ = add(
	// tc-fset
	form("upload", w.PostForm{ID: "post-import", Action: "/app/contacts/nextcloud/kontakty/import", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}, File: "file"}, "components-37", "tc-post", 2),
	form("default", w.PostForm{ID: "post-disconnect", Action: "/app/settings/connections/nextcloud/disconnect", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}}, "components-37", "tc-post", 1),
	form("template", w.FieldSet{Rubric: w.Rubric{Text: "Phone numbers"}, Items: []w.FieldSetPart{w.Field{Label: "Number", Parts: []w.FieldPart{w.Input{Name: "tel.0", Value: "+7 912 445-10-88", Label: "Number", Mono: true}}, Width: w.FieldMD}, w.Template{ID: "tpl-tel", Items: []w.FieldSetPart{w.Field{Label: "Number", Parts: []w.FieldPart{w.Input{Name: "tel.{n}", Label: "Number", Mono: true}}, Width: w.FieldMD}}}, w.Grow{}, w.Button{Label: "Add phone", Icon: "plus", Small: true, Name: "add", Value: "tel", Adds: "tpl-tel"}}}, "components-38", "tc-fset", 1),
	form("plain", w.FieldSet{Rubric: rub, Items: fsItems(fld, fld, bt, w.Grow{}, fld, fld, bt, w.Grow{}, bt)}, "s-contacts-3", "tc-fset", 1),
	form("plain", w.FieldSet{Rubric: rub, Items: fsItems(w.Line{Place: w.Grown})}, "s-agenda-3", "tc-fset", 3),
	form("lead", w.FieldSet{Rubric: rub, Variant: w.FieldSetLead, Items: fsItems(fld)}, "account-1", "tc-fset", 1),
	form("plain", w.FieldSet{Rubric: rub, Items: fsItems(grown, fld)}, "account-1", "tc-fset", 3),
	form("lead", w.FieldSet{Rubric: rub, Variant: w.FieldSetLead, Items: fsItems(grown, fld)}, "account-2", "tc-fset", 1),

	// tc-form
	form("action", w.Form{Items: []w.FormPart{w.Field{Label: "Display name", Parts: []w.FieldPart{w.Input{Name: "display", Value: "Mix", Label: "Display name"}}, Width: w.FieldMD}, w.Field{Label: "Week starts on", Parts: []w.FieldPart{w.Pick{Name: "week", Options: []w.PickOption{{Value: "mon", Label: "Monday", On: true}, {Value: "sun", Label: "Sunday"}, {Value: "sat", Label: "Saturday", Unavailable: true}}, Label: "Week starts on"}}, Width: w.FieldSM}, w.Field{Label: "Theme", Parts: []w.FieldPart{w.Seg{Options: []w.SegOption{{Label: "Auto", On: true, Value: "auto"}, {Label: "Light", Value: "light"}, {Label: "Dark", Value: "dark"}}, Name: "theme", Theme: true}}}, w.Field{Parts: []w.FieldPart{w.Check{On: true, Name: "compact", Value: "1", Parts: []w.CheckPart{w.Text("Compact rows")}}, w.Check{Name: "mono", Value: "1", Parts: []w.CheckPart{w.Text("Monospaced dates")}}}, Gap: 7, OwnLine: true, Place: w.Grown}, w.Field{Label: "Password", Parts: []w.FieldPart{w.Input{Kind: w.InputPassword, Name: "password", Label: "Password", Disabled: true}}, Width: w.FieldMD}, w.Field{Parts: []w.FieldPart{w.Check{On: true, Name: "nopass", Value: "1", Disables: "password", Parts: []w.CheckPart{w.Text("The server asks no password")}}}, OwnLine: true, Place: w.Grown}, w.FieldSet{Rubric: w.Rubric{Text: "Phone numbers"}, Items: []w.FieldSetPart{w.Field{Label: "Number", Parts: []w.FieldPart{w.Input{Name: "tel.0", Value: "+7 912 445-10-88", Label: "Number", Mono: true}}, Width: w.FieldMD}, w.Template{ID: "tpl-tel", Items: []w.FieldSetPart{w.Field{Label: "Number", Parts: []w.FieldPart{w.Input{Name: "tel.{n}", Label: "Number", Mono: true}}, Width: w.FieldMD}}}, w.Grow{}, w.Button{Label: "Add phone", Icon: "plus", Small: true, Name: "add", Value: "tel", Adds: "tpl-tel"}}}, w.Field{Parts: []w.FieldPart{w.Button{Label: "Save", Kind: w.ButtonPrimary, Name: "op", Value: "save"}}}}, Action: "/app/settings/account", Fields: []w.Hidden{{Name: "csrf", Value: "7f3a91de"}}, Variant: w.FormFlushAll}, "components-38", "tc-form", 1),
	form("plain", w.Form{Items: formItems(fld, fld, fld, fld, fld, fs, fs, fs, fs, fs)}, "s-agenda-3", "tc-form", 1),
	form("tight", w.Form{Tight: true, Items: formItems(fld)}, "s-io-5", "tc-form", 1),
	form("flush", w.Form{Items: []w.FormPart{w.Field{Label: "Label", Parts: []w.FieldPart{w.Input{Placeholder: "optional", Label: "Label"}}, Width: w.FieldMD}, w.Field{Label: "Server URL", Parts: []w.FieldPart{w.Input{Value: "https://dav.office.example/dav.php/", Label: "Server URL", Mono: true, Focus: true}, w.FieldHint{Parts: []w.Inline{w.Mono{Text: ".well-known", Form: w.MonoCode}, w.Text(" is tried first, then the URL exactly as typed — which is the reliable path for Baikal.")}}}, Width: w.FieldFull}, w.Field{Label: "Username", Parts: []w.FieldPart{w.Input{Value: "m.yenuchenko", Label: "Username"}}, Width: w.FieldMD}, w.Field{Label: "Password", Parts: []w.FieldPart{w.Input{Value: "••••••••••", Label: "Password"}}, Width: w.FieldMD}, w.Field{Label: "Certificate fingerprint", LabelAside: w.Dim{Parts: []w.Inline{w.Text("SHA-256 · optional")}}, Parts: []w.FieldPart{w.Input{Placeholder: "only for a certificate the system does not trust", Label: "Certificate fingerprint", Mono: true}}, Width: w.FieldFull}}, Variant: w.FormFlush}, "s-profile-2", "tc-form", 1),
	form("flush", w.Form{Variant: w.FormFlushAll, Items: formItems(fld, fld, fld, fld)}, "s-auth-2", "tc-form", 1),
	form("flush", w.Form{Variant: w.FormGap12, Items: formItems(grown, grown, grown)}, "m-desktop-4", "tc-form", 1),
	form("columns", w.Form{Items: []w.FormPart{w.Field{Label: "Photo", LabelSpan: true, Parts: []w.FieldPart{w.Photo{Side: []w.FieldPart{w.Line{Items: []w.LinePart{w.Button{Label: "Replace"}, w.Button{Label: "Crop"}}, Gap: w.LineGap6}, w.Button{Label: "Remove", Kind: w.ButtonDanger}, w.FieldHint{Parts: []w.Inline{w.Text("EXIF orientation applied, metadata stripped.")}, Variant: w.HintNarrow}}}}, Width: w.FieldFixAuto}, w.FormCols{Items: []w.Field{w.Field{Label: "Display name", Parts: []w.FieldPart{w.Input{Value: "Анна Ковалёва", Label: "Display name", Focus: true}}, Width: w.FieldXL}, w.Field{Label: "Given", Parts: []w.FieldPart{w.Input{Value: "Анна", Label: "Given"}}, Width: w.FieldSM}, w.Field{Label: "Family", Parts: []w.FieldPart{w.Input{Value: "Ковалёва", Label: "Family"}}, Width: w.FieldSM}, w.Field{Label: "Nickname", Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Nickname"}}, Width: w.FieldSM}, w.Field{Label: "Organisation", Parts: []w.FieldPart{w.Input{Value: "Издательство «Слово»", Label: "Organisation"}}, Width: w.FieldMD}, w.Field{Label: "Title", Parts: []w.FieldPart{w.Input{Value: "Выпускающий редактор", Label: "Title"}}, Width: w.FieldMD}, w.Field{Label: "Birthday", Parts: []w.FieldPart{w.Input{Value: "14.03.1987", Label: "Birthday", Mono: true, Numeric: true}, w.FieldHint{Parts: []w.Inline{w.Text("Type it: 14.03.1987, or 14.03 without the year.")}}}, Width: w.FieldSM}}}, w.FieldSet{Rubric: w.Rubric{Text: "Phone numbers"}, Items: []w.FieldSetPart{w.Field{Label: "Label", Parts: []w.FieldPart{w.Pick{Value: "mobile"}}, Width: w.FieldXS}, w.Field{Label: "Number", Parts: []w.FieldPart{w.Input{Value: "+7 912 445-10-88", Label: "Number", Mono: true}}, Width: w.FieldMD}, w.Button{Label: "Remove", Kind: w.ButtonDanger}, w.Grow{}, w.Field{Parts: []w.FieldPart{w.Input{Value: "work", Label: "Label"}}, Width: w.FieldXS}, w.Field{Parts: []w.FieldPart{w.Input{Value: "+7 495 120-45-00", Label: "Number", Mono: true}}, Width: w.FieldMD}, w.Button{Label: "Remove", Kind: w.ButtonDanger}, w.Grow{}, w.Button{Label: "Add phone", Icon: "plus", Small: true}}}, w.FieldSet{Rubric: w.Rubric{Text: "Email addresses"}, Items: []w.FieldSetPart{w.Field{Label: "Label", Parts: []w.FieldPart{w.Pick{Value: "work"}}, Width: w.FieldXS}, w.Field{Label: "Address", Parts: []w.FieldPart{w.Input{Value: "a.kovaleva@example.org", Label: "Address"}}, Width: w.FieldLG}, w.Button{Label: "Remove", Kind: w.ButtonDanger}, w.Grow{}, w.Button{Label: "Add email", Icon: "plus", Small: true}}}, w.FieldSet{Rubric: w.Rubric{Text: "Addresses"}, Items: []w.FieldSetPart{w.Field{Label: "Label", Parts: []w.FieldPart{w.Pick{Value: "home"}}, Width: w.FieldXS}, w.Field{Label: "Street", Parts: []w.FieldPart{w.Input{Value: "ул. Садовая, 12, кв. 5", Label: "Street"}}, Width: w.FieldLG}, w.Field{Label: "City", Parts: []w.FieldPart{w.Input{Value: "Москва", Label: "City"}}, Width: w.FieldSM}, w.Field{Label: "Region", Parts: []w.FieldPart{w.Input{Placeholder: "—", Label: "Region"}}, Width: w.FieldSM}, w.Field{Label: "Postcode", Parts: []w.FieldPart{w.Input{Value: "123001", Label: "Postcode", Mono: true}}, Width: w.FieldXS}, w.Field{Label: "Country", Parts: []w.FieldPart{w.Input{Value: "Россия", Label: "Country"}}, Width: w.FieldSM}, w.Button{Label: "Remove", Kind: w.ButtonDanger}, w.Grow{}, w.Button{Label: "Add address", Icon: "plus", Small: true}}}, w.FieldSet{Rubric: w.Rubric{Text: "Other"}, Items: []w.FieldSetPart{w.Field{Label: "Tags", Parts: []w.FieldPart{w.Input{Value: "работа, издательство", Label: "Tags"}}, Width: w.FieldLG}, w.Field{Label: "Note", Parts: []w.FieldPart{w.Area{Text: "Пишет по вторникам, отвечает в тот же день.", Label: "Text"}}, Width: w.FieldFull}}}, w.Field{Parts: []w.FieldPart{w.Button{Label: "Show 6 empty fields"}}}}}, "s-contacts-3", "tc-form", 1),
	form("tight", w.Form{Tight: true, Variant: w.FormBottom12, Items: formItems(fld)}, "s-io-3", "tc-form", 1),
	form("column", w.Form{Variant: w.FormColumnFlush, Items: formItems(fld, fld)}, "s-auth-1", "tc-form", 1),
	form("flush", w.Form{Variant: w.FormFlush, Items: formItems(fld, fld, fld, grown)}, "account-3", "tc-form", 1),
	form("plain", w.Form{Variant: w.FormBottom12, Items: formItems(fld)}, "s-davloom-1", "tc-form", 1),
	form("flush", w.Form{Variant: w.FormGap12, Items: formItems(grown, fld, fld, grown)}, "m-desktop-6", "tc-form", 1),
	form("column", w.Form{Variant: w.FormColumnGap10Mid, Items: formItems(w.Rubric{Text: "x"}, fld, chk, chk, fld, fld)}, "m-desktop-11", "tc-form", 1),

	// tc-formfoot
	form("flush", w.FormFoot{Variant: w.FormFootFlush, Items: footItems(bt, bt, w.Push{})}, "s-contacts-5", "tc-formfoot", 1),
	form("plain", w.FormFoot{Items: footItems(w.Dim{})}, "s-contacts-3", "tc-formfoot", 1),
	form("no-rule", w.FormFoot{Variant: w.FormFootNoRule, Items: footItems(bt, bt, w.Button{Label: "x", Place: w.Pushed})}, "s-dup-2", "tc-formfoot", 1),
	form("wrap", w.FormFoot{Variant: w.FormFootWrap, Items: footItems(bt, bt, bt, bt)}, "s-conflict-1", "tc-formfoot", 1),
	form("no-rule", w.FormFoot{Variant: w.FormFootNoRule, Items: footItems(bt, bt)}, "s-dup-5", "tc-formfoot", 1),
	form("flush", w.FormFoot{Variant: w.FormFootFlush, Items: footItems(bt, w.Span{Face: w.FaceAccent12, Text: "Already have an account? Sign in"})}, "s-auth-2", "tc-formfoot", 1),
)

func init() {
	Widgets = append(Widgets, w.FieldSet{}, w.Form{}, w.FormFoot{}, w.PostForm{})
}

func init() {
	Parts = append(Parts, w.Text(""), w.Strong(""), w.Hint{}, w.Dot{}, w.Span{}, w.Faced{}, w.Choice{}, w.Hidden{}, w.PickOption{}, w.Template{},
		w.FieldHint{}, w.Swatches{}, w.Zoom{}, w.FieldRow{}, w.Group{}, w.Photo{}, w.FieldGrid{},
		w.FormCols{}, w.FormColumn{}, w.RubricAside{}, w.RubricCol{})
}

var _ = add(
	form("bare", w.Field{Width: w.FieldW100, Parts: []w.FieldPart{w.Diag{Text: "[fetch] GET /ru-holidays.ics → 200 · 0.4 s · 34 KB [type] text/calendar; charset=utf-8 [parse] VCALENDAR · 118 VEVENT · 6 VTODO · 2 VJOURNAL [name] X-WR-CALNAME: Праздники России [access] the server asked for no password"}}}, "s-profile-10", "tc-field", 3),
	form("plain", w.Form{Items: []w.FormPart{
		w.FieldSet{Rubric: w.Rubric{Text: "Where to"}, Items: []w.FieldSetPart{
			w.Field{Place: w.Grown, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}, OwnLine: true},
			w.Field{Label: "Collection", Width: w.FieldLG, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
			w.Field{Label: "Folder", Width: w.FieldMD, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
			w.Msg{Parts: []w.MsgPart{w.Strong("Both"), w.Text(" writes to WebDAV and still has to be taken — so the trigger below is a script, not a line. This is the same job editor as above; only this control changed.")}, Space: w.MsgFlushTop4, Place: w.Grown}}, Variant: w.FieldSetLead},
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "The script to paste", Items: []w.RubricPart{w.Mono{Text: "visible only now", Variant: w.MonoTightPlain}}},
			w.CardText{Kind: w.CardAbove, Parts: []w.CardTextPart{w.Text("Destination: "), w.Key("WebDAV and download"), w.Text(" — two steps, so this is a file, not a line")}},
			w.Seg{Options: []w.SegOption{{Label: "crontab", On: true}, {Label: "GoSentry"}, {Label: "Gitea"}, {Label: "curl"}}},
			w.Diag{Text: "#!/bin/sh set -eu h='Authorization: Bearer cbk_7f3a-91de-4402-lamp' b=https://carrel.example/api/backup out=/var/backups id=$(curl -fsS -X POST -H \"$h\" \"$b/run?job=personal\") until [ \"$(curl -fsS -H \"$h\" \"$b/status?id=$id\")\" != running ]; do sleep 20; done curl -fsS -H \"$h\" \"$b/$id/download\" -o \"$out/carrel-personal-$(date +%F).zip\"", Variant: w.DiagInk},
			w.CardRow{Kind: w.CardTop8, Gap: 8, Items: []w.CardRowPart{w.Button{Label: "Copy the script", Icon: "dup", Kind: w.ButtonPrimary}, w.Button{Label: "Save as", Icon: "down"}, w.Button{Label: "Regenerate"}}},
			w.CardText{Kind: w.CardUsage, Parts: []w.CardTextPart{w.Text("Put it where the password is to live, make it executable, and give "), w.Mono{Text: "crontab -e"}, w.Text(" one line: "), w.Mono{Text: "0 3 * * * /usr/local/bin/carrel-personal.sh"}, w.Text(" — the schedule stays a line, the work moves into the file. The date has to live in the script and not in the crontab line: "), w.Mono{Text: "cron"}, w.Text(" reads "), w.Mono{Text: "%"}, w.Text(" as a newline, so "), w.Mono{Text: "date +%F"}, w.Text(" written up there would break the entry.")}},
			w.Msg{Parts: []w.MsgPart{w.Text("A copy that fails needs no separate check: the download answers with an error, "), w.Mono{Text: "set -eu"}, w.Text(" stops the script on it, and cron sends the mail. The wait is open-ended on purpose — a copy takes what it takes, and "), w.Mono{Text: "409"}, w.Text(" until then means «not yet», not «broken».")}, Space: w.MsgFlushTop9}}, Variant: w.CardChosenFlush, Place: w.Grown}}}, "s-backup-3", "tc-form", 1),
	form("column", w.Form{Items: []w.FormPart{
		w.Card{Items: []w.CardPart{
			w.CardText{Kind: w.CardTerm, Parts: []w.CardTextPart{w.Text("Wipe the account")}},
			w.Prose{Parts: []w.Inline{w.Text("We send a letter to the confirmed address. The link in it is not a reset: after you confirm, everything stored here is erased — connections become unreadable and are deleted — and the same link then lets you set a new password. Address books, calendars, tasks and notes on the servers stay where they are.")}, Variant: w.ProseMutedClose},
			w.Button{Label: "Send a confirmation letter", Kind: w.ButtonPrimary}}, Variant: w.CardSnug},
		w.Card{Items: []w.CardPart{
			w.CardText{Kind: w.CardTerm, Parts: []w.CardTextPart{w.Text("Ask the administrator")}},
			w.Prose{Parts: []w.Inline{w.Text("This account deposited its encryption key. The administrator can recover it with the instance master password, and your data is kept. Shown only when deposit is on for this account.")}, Variant: w.ProseMutedClose},
			w.Button{Label: "Ask the administrator"}}, Variant: w.CardSnug}}, Variant: w.FormColumnGap10}, "s-auth-6", "tc-form", 1),
	form("plain", w.Form{Items: []w.FormPart{
		w.Card{Items: []w.CardPart{
			w.Rubric{Text: "The fields to paste", Items: []w.RubricPart{w.Mono{Text: "visible only now", Variant: w.MonoTightPlain}}},
			w.Seg{Options: []w.SegOption{{Label: "crontab"}, {Label: "GoSentry", On: true}, {Label: "Gitea"}, {Label: "curl"}}},
			w.CardGrid{Items: []w.CardGridPart{w.Dim{Parts: []w.Inline{w.Text("Name")}, Size: w.Dim11}, w.Input{Value: "Carrel · personal", Label: "x"}, w.Button{Label: "Copy", Small: true}, w.Dim{Parts: []w.Inline{w.Text("Folder")}, Size: w.Dim11}, w.Input{Value: "Backup", Label: "x"}, w.Button{Label: "Copy", Small: true}, w.Dim{Parts: []w.Inline{w.Text("Schedule")}, Size: w.Dim11}, w.Input{Value: "0 3 * * *", Label: "x"}, w.Button{Label: "Copy", Small: true}, w.Dim{Parts: []w.Inline{w.Text("Command")}, Size: w.Dim11}, w.Input{Value: "curl", Label: "x"}, w.Button{Label: "Copy", Small: true}, w.Dim{Parts: []w.Inline{w.Text("Arguments")}, Size: w.Dim11, Variant: w.DimTop}, w.Area{Text: pasteArgs, Variant: w.AreaPaste, Mono: true, Label: "Arguments"}, w.Button{Label: "Copy", Small: true}}},
			w.CardText{Kind: w.CardUsage, Parts: []w.CardTextPart{w.Text("GoSentry is a scheduler of commands, not an HTTP checker: its job dialog asks for "), w.Strong("Name"), w.Text(", "), w.Strong("Folder"), w.Text(", "), w.Strong("Schedule"), w.Text(", "), w.Strong("Command"), w.Text(" and "), w.Strong("Arguments"), w.Text(" (one argument per line, no quotes). This tab copies those fields, not a URL and a header. Leave "), w.Strong("Start only, do not wait for exit"), w.Text(" off — "), w.Mono{Text: "curl"}, w.Text(" exits when the "), w.Mono{Text: "202"}, w.Text(" comes back.")}},
			w.Msg{Parts: []w.MsgPart{w.Text("A second GoSentry job, same form: Command "), w.Mono{Text: "curl"}, w.Text(", Arguments a "), w.Mono{Text: "GET /api/backup/status?job=personal"}, w.Text(", failing when the last "), w.Strong("successful"), w.Text(" copy is older than 26 hours. The first job says the copy was asked for; only the second says it happened.")}, Space: w.MsgFlushTop9}}, Variant: w.CardChosenFlush, Place: w.Grown}}}, "s-backup-4", "tc-form", 1),
	form("columns", w.Form{Items: []w.FormPart{
		w.FormColumn{Wide: true, Items: []w.FormColumnPart{w.Card{Items: []w.CardPart{
			w.Rubric{Text: "What goes into the profile"},
			w.List{Rows: []w.ListRow{
				w.SlimRow{Mark: w.Mark{Kind: w.MarkMulti}, Title: "nextcloud · mix", Inside: "· Личный, Работа, Контакты", Aside: "CalDAV, CardDAV"},
				w.SlimRow{Ticked: w.TickedOff, Mark: w.Mark{Kind: w.MarkMulti}, Title: "radicale · vps · mix", Inside: "· Проекты", Aside: "CalDAV"}}, Inset: true}}}, w.Msg{Icon: "shield", Parts: []w.MsgPart{w.Strong("The profile carries passwords in cleartext."), w.Text(" It is an ordinary file: it lands in Downloads, in a messenger, in the phone’s backup. Open it on the device and delete it. The link opens once and lives 5 minutes, for you signed in to Carrel; nothing is written to the server’s disk, and the issue is recorded in the audit log.")}, Tone: w.MsgAlert, Space: w.MsgFlush}}},
		w.FormColumn{Items: []w.FormColumnPart{w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Profile"},
			w.CardHint{Parts: []w.CardTextPart{w.Text("Issued inside the unlock window: nothing more is asked.")}},
			w.CardRow{Kind: w.CardActs, Items: []w.CardRowPart{w.Button{Label: "Generate profile", Kind: w.ButtonDanger}}}}}, w.Card{Items: []w.CardPart{
			w.Rubric{Text: "Link ready until 09:46"},
			w.CardHint{Parts: []w.CardTextPart{w.Text("It opens once, for you signed in to Carrel. Open it on the device, install the profile, then delete the file.")}},
			w.CardRow{Kind: w.CardTop9, Gap: 6, Items: []w.CardRowPart{w.Button{Label: "Download profile", Icon: "down"}, w.Button{Label: "Revoke", Kind: w.ButtonDanger}}}}}}}}, Variant: w.FormFlushBottom}, "s-plan-1", "tc-form", 1),
	form("plain", w.FieldSet{Rubric: w.Rubric{Text: "Account"}, Items: []w.FieldSetPart{
		w.Field{Label: "Login", Width: w.FieldMD, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Field{Label: "Password", Width: w.FieldMD, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Msg{Parts: []w.MsgPart{w.Text("The password is the only key to the saved connections. Running locally changes nothing about that: nobody can recover it, including the app on your own computer.")}, Tone: w.MsgAlert, Space: w.MsgFlushTop2, Place: w.Grown}}, Variant: w.FieldSetRule}, "m-desktop-6", "tc-fset", 1),
	form("lead", w.FieldSet{Rubric: w.Rubric{Text: "Where to"}, Items: []w.FieldSetPart{
		w.Field{Place: w.Grown, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}, OwnLine: true},
		w.Field{Label: "Collection", Width: w.FieldLG, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Field{Label: "Folder", Width: w.FieldMD, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Msg{Parts: []w.MsgPart{w.Strong("Both"), w.Text(" writes to WebDAV and still has to be taken — so the trigger below is a script, not a line. This is the same job editor as above; only this control changed.")}, Space: w.MsgFlushTop4, Place: w.Grown}}, Variant: w.FieldSetLead}, "s-backup-3", "tc-fset", 1),
	form("plain", w.FieldSet{Rubric: w.Rubric{Text: "Where to"}, Items: []w.FieldSetPart{
		w.Field{Place: w.Grown, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}, OwnLine: true},
		w.Field{Label: "Collection", Width: w.FieldLG, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Field{Label: "Folder", Width: w.FieldMD, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Field{Label: "Keep copies", Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Field{Place: w.Grown, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}},
		w.Msg{Parts: []w.MsgPart{w.Strong("Download"), w.Text(" streams the archive out and writes nothing to the server: no folder, no retention — the copy is yours to keep. It can still run at three in the morning, but only on an external trigger, because a stream needs somebody at the other end: a browser is not there at that hour, a machine with a disk is. "), w.Strong("Both"), w.Text(" does the two at once — and then the trigger has "), w.Strong("two steps"), w.Text(", because the file still has to be taken.")}, Space: w.MsgFlushTop4, Place: w.Grown}}}, "s-backup-2", "tc-fset", 2),
	form("lead", w.FieldSet{Rubric: w.Rubric{Text: "Small files stay in the record"}, Items: []w.FieldSetPart{
		w.Fields{Items: []w.FieldsItem{
			{Term: "Up to 2 MB", Value: []w.ValuePart{w.Text("written into the event, the task or the note itself, and travels with it")}},
			{Term: "Larger", Value: []w.ValuePart{w.Text("put in the folder below; the record keeps a link to it")}}}, Variant: w.FieldsBleed130, Place: w.Grown},
		w.Field{Place: w.Grown, Parts: []w.FieldPart{w.Input{Value: "x", Label: "x"}}}}, Variant: w.FieldSetLead}, "attachments-1", "tc-fset", 1),
	form("gap", w.Line{Gap: w.LineGap9, Items: []w.LinePart{w.Meter{Percent: 100, Width: w.MeterFill}, w.Mono{Text: "0.8.0 · 14.2 MB · signature checked", Variant: w.Mono11Dusty}}}, "m-desktop-14", "tc-line", 1),
	form("default", w.Stack{Gap: w.StackGap10, Items: []w.StackPart{w.Card{Items: []w.CardPart{
		w.CardPick{Title: "Local — run it here", Note: []w.CardTextPart{w.Text("Carrel runs inside the app, on this computer. Nothing to deploy.")}, Aside: "Suits one person and one computer. A phone cannot open this Carrel until access over the network is allowed."}}, Variant: w.CardChosenFlush}, w.Card{Items: []w.CardPart{
		w.CardPick{Ticked: w.TickedOff, Title: "Remote — connect to a server", Note: []w.CardTextPart{w.Text("Carrel is already running — yours or a shared one. The app becomes its window.")}, Aside: "One instance for all your devices, with the same settings everywhere."}}, Variant: w.CardFlush}}}, "m-desktop-5", "tc-stack", 1),
	form("plain", w.FieldSet{Rubric: w.Rubric{Text: "What to copy"}, Items: []w.FieldSetPart{
		w.FieldGrid{Items: []w.FieldGridPart{w.TickRow{Ticked: w.TickedPart, Form: w.TickRowGroup, Mark: w.Mark{Kind: w.MarkMulti}, Label: "nextcloud", Meta: "4 of 6"}, w.TickRow{Ticked: w.TickedOff, Form: w.TickRowGroup, Mark: w.Mark{Kind: w.MarkMulti}, Label: "radicale · vps", Meta: "0 of 2"}, w.TickRow{Form: w.TickRowNested, Mark: w.Mark{Colour: "#4A6B52"}, Label: "Личный", Meta: "cal."}, w.TickRow{Ticked: w.TickedOff, Form: w.TickRowNested, Mark: w.Mark{Colour: "#6B4E7A"}, Label: "Проекты", Meta: "cal."}, w.TickRow{Ticked: w.TickedOff, Form: w.TickRowNested, Mark: w.Mark{Colour: "#3F6E8C"}, Label: "Работа", Meta: "cal."}, w.TickRow{Ticked: w.TickedOff, Form: w.TickRowNested, Mark: w.Mark{Colour: "#8C6239"}, Label: "Дежурства", Meta: "cal."}, w.TickRow{Form: w.TickRowNested, Mark: w.Mark{Colour: "#A05A2C"}, Label: "Архив", Meta: "book"}, w.TickRow{Ticked: w.TickedPart, Form: w.TickRowGroup, Mark: w.Mark{Kind: w.MarkMulti}, Label: "sftpgo · vps", Meta: "1 of 2"}, w.TickRow{Form: w.TickRowNested, Mark: w.Mark{Colour: "#4A6B52"}, Label: "Записи", Meta: "notes"}, w.TickRow{Form: w.TickRowNested, Mark: w.Mark{Colour: "#4A6B52"}, Label: "Documents", Meta: "files · 2.0 GB"}, w.TickRow{Ticked: w.TickedOff, Form: w.TickRowNested, Mark: w.Mark{Colour: "#9A9280"}, Label: "Public", Meta: "files · ro"}}},
		w.Msg{Parts: []w.MsgPart{w.Text("A file collection is copied as files. Tasks and notes live inside calendar collections and are copied with them — they cannot be picked separately. The size beside each collection is what it took last time.")}, Space: w.MsgFlushTop4, Place: w.Grown}}}, "s-backup-2", "tc-fset", 1),
)
