package refs

import w "carrel.dev/typecase/widget"

func seg(on int, labels ...string) []w.SegOption {
	var out []w.SegOption
	for i, l := range labels {
		out = append(out, w.SegOption{Label: l, On: i == on})
	}
	return out
}

func check(on bool, parts ...w.CheckPart) w.Check { return w.Check{On: on, Parts: parts} }

func hint(parts ...w.Inline) w.FieldHint { return w.FieldHint{Parts: parts} }

var btn = w.Button{Label: "x"}

var _ = add(
	// tc-seg
	form("default", w.Seg{Options: seg(0, "Remote", "Local")}, "components-13", "tc-seg", 1),
	form("default", w.Seg{Options: seg(0, "Open 12", "Done 48", "All")}, "components-9", "tc-seg", 1),
	form("start", w.Seg{Options: seg(0, "0°", "90°", "180°", "270°"), Variant: w.SegStart}, "s-contacts-5", "tc-seg", 1),
	form("default", w.Seg{Options: seg(0, "All 24", "Events 14", "Tasks 2", "Notes 4", "Files 4")}, "s-contacts-7", "tc-seg", 1),
	form("start", w.Seg{Options: seg(0, "To WebDAV", "Download", "Both"), Variant: w.SegStart}, "s-backup-2", "tc-seg", 1),
	form("start", w.Seg{Options: seg(0, "Comfortable", "Compact"), Variant: w.SegStart}, "appearance-1", "tc-seg", 2),
	form("default", w.Seg{Options: seg(1, "Day", "Week", "Month", "Range")}, "s-agenda-1", "tc-seg", 1),
	form("icons", w.Seg{Options: []w.SegOption{{Label: "List", Icon: "menu"}, {Label: "Tiles", Icon: "grid", On: true}}}, "s-files-5", "tc-seg", 1),
	form("compact", w.Seg{Options: seg(0, "Auto", "Light", "Dark"), Variant: w.SegStartSmall}, "account-5", "tc-seg", 1),
	form("default", w.Seg{Options: seg(0, "Apple", "Android", "Thunderbird", "Anything else")}, "s-plan-1", "tc-seg", 1),
	form("start", w.Seg{Options: seg(1, "Read only", "Read and write"), Variant: w.SegStart}, "s-davloom-2", "tc-seg", 1),
	form("default", w.Seg{Options: seg(0, "MO", "TU", "WE", "TH", "FR", "SA", "SU")}, "s-agenda-3", "tc-seg", 1),
	form("default", w.Seg{Options: seg(1, "Full width", "Reading width")}, "s-notes-3", "tc-seg", 1),
	form("icons", w.Seg{Options: []w.SegOption{{Icon: "menu", On: true, Title: "List"}, {Icon: "grid", Title: "Tiles"}}}, "s-files-1", "tc-seg", 2),
	form("default", w.Seg{Options: seg(0, "All", "Contacts", "Events", "Tasks", "Notes", "Files")}, "s-find-1", "tc-seg", 1),
	form("default", w.Seg{Options: seg(2, "Apple", "Android", "Thunderbird", "Anything else")}, "s-plan-4", "tc-seg", 1),
	form("default", w.Seg{Options: seg(3, "Apple", "Android", "Thunderbird", "Anything else")}, "s-plan-5", "tc-seg", 1),
	form("compact", w.Seg{Options: seg(1, "read only", "read and write"), Variant: w.SegCompact}, "s-davloom-1", "tc-seg", 1),
	form("compact", w.Seg{Options: seg(0, "read only", "read and write"), Variant: w.SegCompact}, "s-davloom-1", "tc-seg", 2),
	form("start", w.Seg{Options: seg(2, "To WebDAV", "Download", "Both"), Variant: w.SegStart}, "s-backup-3", "tc-seg", 1),
	form("start", w.Seg{Options: seg(1, "crontab", "GoSentry", "Gitea", "curl"), Variant: w.SegStart}, "s-backup-4", "tc-seg", 1),

	// tc-check
	form("off", check(false, w.Text("Print with photos")), "components-10", "tc-check", 2),
	form("on", check(true, w.Text("All day")), "components-10", "tc-check", 1),
	form("rich", check(true, w.Choice{Source: "Your edit · include", Value: "+7 495 120-45-00"}), "s-conflict-1", "tc-check", 1),
	form("rich", check(true, w.Text("Allow"), w.Mono{Text: ".ics"}, w.Text("download and subscription")), "s-publish-2", "tc-check", 1),
	form("tertiary", w.Check{On: true, Parts: []w.CheckPart{w.Text("Events")}, Tertiary: true}, "s-agenda-1", "tc-check", 1),
	form("tertiary", w.Check{Parts: []w.CheckPart{w.Text("Tasks due")}, Tertiary: true}, "s-agenda-1", "tc-check", 2),
	form("rich", check(false, w.Choice{Source: "Leave it out", Value: "—"}), "s-conflict-1", "tc-check", 2),
	// s-backup-2 .tc-check #4 holds an input inside the button: nested controls
	// (a11y). Its record waits for the owner's decision on the mockup.
	form("rich", check(false, w.Text("Tasks"), w.Dim{Parts: []w.Inline{w.Text("· VTODO — not every client understands them")}}), "s-davloom-2", "tc-check", 3),
	form("rich", check(true, w.Text("WebDAV folders"), w.Dim{Parts: []w.Inline{w.Text("· mount as a network drive")}}), "s-davloom-2", "tc-check", 5),
	form("rich", check(true, w.Text("On a clock of your own —"), w.Mono{Text: "cron"}, w.Text(", GoSentry, a job in Gitea"),
		w.Badge{Label: "local", Tone: w.BadgeLocal, Title: "Local instance: the address in the line answers only from inside your own network"}), "s-backup-2", "tc-check", 5),
	form("rich", check(false, w.Text("Allow"), w.Mono{Text: ".vcf"}, w.Text("download and subscription")), "s-publish-3", "tc-check", 6),
	form("rich", check(false, w.Dim{Parts: []w.Inline{w.Text("Start at sign-in — not available in a portable copy")}}), "m-desktop-4", "tc-check", 1),

	// tc-range
	form("default", w.Range{From: w.Input{Kind: w.InputDate, Value: "2026-01-01", Mono: true, Label: "From"}, To: w.Input{Value: "today", Mono: true, Label: "To"}, Show: w.Button{Label: "Show"}}, "components-18", "tc-range", 1),
	form("wrap", w.Range{From: w.Input{Kind: w.InputDate, Value: "2026-08-10", Mono: true, Label: "From"}, To: w.Input{Kind: w.InputDate, Value: "2026-08-16", Mono: true, Label: "To"}, Show: w.Button{Label: "Show"}, Wrap: true}, "s-agenda-1", "tc-range", 2),

	// tc-line, tc-stack, tc-grow, tc-push
	form("span", w.Line{Gap: w.LineGap6, Items: []w.LinePart{btn, btn}}, "s-contacts-3", "tc-line", 1),
	form("span", w.Line{Gap: w.LineGap8, Items: []w.LinePart{w.Input{Label: "x"}, btn}}, "components-20", "tc-line", 1),
	form("gap", w.Line{Gap: w.LineGap7, Items: []w.LinePart{btn, w.Kbd{Keys: "Esc"}}}, "s-contacts-2", "tc-line", 1),
	form("block", w.Line{Gap: w.LineGap8, Block: true, Items: []w.LinePart{btn, w.Dim{}}}, "s-tasks-3", "tc-line", 1),
	form("block", w.Line{Gap: w.LineGap12, Block: true, Items: []w.LinePart{chk, chk}}, "s-dup-1", "tc-line", 1),
	form("gap", w.Line{Gap: w.LineGap3, Items: []w.LinePart{w.Mark{}, w.Mark{}, w.Mark{}, w.Mark{}}}, "s-davloom-1", "tc-line", 2),
	form("gap", w.Line{Gap: w.LineGap8, Items: []w.LinePart{w.Strong("nextcloud"), w.Dim{}, w.Mono{}, w.Dim{}, w.Mono{}, btn}}, "s-plan-4", "tc-line", 1),
	form("block", w.Line{Gap: w.LineGap8, Block: true, Items: []w.LinePart{btn, btn, w.Dim{Place: w.Pushed}}}, "m-desktop-3", "tc-line", 2),
	form("block", w.Line{Gap: w.LineGap4, Block: true, Items: []w.LinePart{w.Tag{}, w.Tag{}}}, "s-notes-2", "tc-line", 1),
	form("span", w.Line{Items: []w.LinePart{w.Badge{Label: "badge"}}}, "account-1", "tc-line", 2),
	form("gap", w.Line{Gap: w.LineGap8, Items: []w.LinePart{w.Mark{}, w.Mono{}, w.Text("Синк-ап по релизу")}}, "m-desktop-8", "tc-line", 1),
	form("block", w.Line{Gap: w.LineGap8, Block: true, Items: []w.LinePart{btn,
		w.Faced{Face: w.FaceLarge, Items: []w.LinePart{w.Text("release-checklist.md"), w.Dim{}, w.Badge{Label: "badge"}}}}}, "s-agenda-3", "tc-line", 1),
	form("block", w.Line{Gap: w.LineGap10, Block: true, Items: []w.LinePart{w.Field{}, w.Field{}}}, "s-publish-2", "tc-line", 1),
	form("block", w.Line{Block: true, Gap: w.LineGap7, Items: []w.LinePart{check(true, w.Text("On sign-in, if the last copy is older than")), w.Input{Value: "7", Mono: true, Variant: w.InputChip, Label: "Days"}, w.Faced{Face: w.FaceLarge, Items: []w.LinePart{w.Text("days")}}}}, "s-backup-2", "tc-line", 1),
	form("block", w.Line{Gap: w.LineGap8, Block: true, Items: []w.LinePart{btn, btn, w.Push{Face: faceOf(w.FaceAccent11_5), Items: []w.PushPart{w.Text("Add a second instance")}}}}, "m-desktop-7", "tc-line", 1),
	form("default", w.Stack{Gap: w.StackGap9, Items: []w.StackPart{btn}}, "components-21", "tc-stack", 1),
	form("default", w.Grow{}, "s-contacts-3", "tc-grow", 1),
	form("default", w.Push{Items: []w.PushPart{btn}}, "components-22", "tc-push", 1),
	form("default", w.Push{Items: []w.PushPart{w.Text("Default")}}, "s-dup-3", "tc-push", 1),

	// tc-rubric
	form("block", w.Rubric{Text: "Rail"}, "shell-1", "tc-rubric", 1),
	form("inline", w.Rubric{Text: "А", Form: w.RubricInline}, "components-12", "tc-rubric", 1),
	form("block", w.Rubric{Text: "nextcloud", Items: []w.RubricPart{w.Mono{Text: "12", Variant: w.MonoTight}}}, "components-11", "tc-rubric", 1),
	form("aside", w.Rubric{Text: "Артём Белов", Items: []w.RubricPart{w.RubricAside{Items: []w.RubricAsidePart{w.Badge{Label: "badge"}, w.Mono{}}}}}, "s-dup-1", "tc-rubric", 2),
	form("aside", w.Rubric{Lead: &w.Mono{Text: "TEL;work"}, Items: []w.RubricPart{w.RubricAside{Text: "only in your edit"}}}, "s-conflict-1", "tc-rubric", 2),
	form("item", w.Rubric{Text: "Write into", Form: w.RubricItem}, "components-6", "tc-rubric", 1),
	form("variant", w.Rubric{Text: "What is pointing at it now", Variant: w.RubricBottom7}, "s-profile-9", "tc-rubric", 3),
	form("variant", w.Rubric{Text: "DAV servers", Variant: w.RubricSplit, Items: []w.RubricPart{w.Mono{Text: "3", Variant: w.MonoTight}}}, "s-profile-1", "tc-rubric", 2),
	form("variant", w.Rubric{Text: "baikal", Variant: w.RubricDanger, Items: []w.RubricPart{w.Line{}}}, "s-profile-1", "tc-rubric", 5),
	form("variant", w.Rubric{Text: "Что делаем", Variant: w.RubricBottom8}, "s-notes-2", "tc-rubric", 3),
	form("variant", w.Rubric{Text: "nextcloud", Variant: w.RubricRail}, "m-phone-2", "tc-rubric", 1),
	form("variant", w.Rubric{Text: "Timeline", Variant: w.RubricBottom2}, "s-contacts-7", "tc-rubric", 2),
	form("item", w.Rubric{Text: "mix", Form: w.RubricItem, Variant: w.RubricItemLine, Items: []w.RubricPart{w.Mono{}}}, "account-5", "tc-rubric", 1),
	form("variant", w.Rubric{Text: "Subscriptions", Variant: w.RubricRail, Items: []w.RubricPart{w.Badge{Label: "badge"}}}, "components-11", "tc-rubric", 2),
	form("item", w.Rubric{Text: "Points per signal", Form: w.RubricItem, Items: []w.RubricPart{w.Push{}, w.RubricCol{Text: "Yours"}}}, "s-dup-3", "tc-rubric", 1),
	form("variant", w.Rubric{Text: "What it opens", Variant: w.RubricGap13}, "s-backup-5", "tc-rubric", 2),
	form("row", w.Rubric{Text: "Attached", Variant: w.RubricRow, Items: []w.RubricPart{w.Num{}, w.Line{Place: w.Pushed}}}, "s-mail-1", "tc-rubric", 6),
)

func faceOf(f w.Face) *w.Face { return &f }

func init() {
	Widgets = append(Widgets, w.Seg{}, w.Check{}, w.Range{}, w.Line{}, w.Stack{}, w.Grow{}, w.Push{}, w.Rubric{})
}
