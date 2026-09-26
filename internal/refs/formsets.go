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
	form("plain", w.FieldSet{Rubric: rub, Items: fsItems(fld, fld, bt, w.Grow{}, fld, fld, bt, w.Grow{}, bt)}, "s-contacts-3", "tc-fset", 1),
	form("plain", w.FieldSet{Rubric: rub, Items: fsItems(w.Line{Place: w.Grown})}, "s-agenda-3", "tc-fset", 3),
	form("lead", w.FieldSet{Rubric: rub, Variant: w.FieldSetLead, Items: fsItems(fld)}, "account-1", "tc-fset", 1),
	form("plain", w.FieldSet{Rubric: rub, Items: fsItems(grown, fld)}, "account-1", "tc-fset", 3),
	form("lead", w.FieldSet{Rubric: rub, Variant: w.FieldSetLead, Items: fsItems(grown, fld)}, "account-2", "tc-fset", 1),

	// tc-form
	form("plain", w.Form{Items: formItems(fld, fld, fld, fld, fld, fs, fs, fs, fs, fs)}, "s-agenda-3", "tc-form", 1),
	form("tight", w.Form{Tight: true, Items: formItems(fld)}, "s-io-5", "tc-form", 1),
	form("flush", w.Form{Variant: w.FormFlush, Items: formItems(fld, fld, fld, fld)}, "s-profile-2", "tc-form", 1),
	form("flush", w.Form{Variant: w.FormFlushAll, Items: formItems(fld, fld, fld, fld)}, "s-auth-2", "tc-form", 1),
	form("flush", w.Form{Variant: w.FormGap12, Items: formItems(grown, grown, grown)}, "m-desktop-4", "tc-form", 1),
	form("columns", w.Form{Items: formItems(fld, w.FormCols{Items: []w.Field{fld, fld, fld, fld, fld, fld, fld}}, fs, fs, fs)}, "s-contacts-3", "tc-form", 1),
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
	Widgets = append(Widgets, w.FieldSet{}, w.Form{}, w.FormFoot{})
}

func init() {
	Parts = append(Parts, w.Text(""), w.Strong(""), w.Hint{}, w.Dot{}, w.Span{}, w.Faced{}, w.Choice{},
		w.FieldHint{}, w.Swatches{}, w.Zoom{}, w.FieldRow{}, w.Group{}, w.Photo{}, w.FieldGrid{},
		w.FormCols{}, w.FormColumn{}, w.RubricAside{}, w.RubricCol{})
}

var _ = wait("children from the family of tables and cards (message, card, fields, table, diagnostics)",
	ref("s-profile-10", "tc-field", 3), ref("s-backup-3", "tc-form", 1), ref("s-auth-6", "tc-form", 1),
	ref("s-backup-4", "tc-form", 1), ref("s-plan-1", "tc-form", 1), ref("m-desktop-6", "tc-fset", 1),
	ref("s-backup-3", "tc-fset", 1), ref("s-backup-2", "tc-fset", 2), ref("attachments-1", "tc-fset", 1),
	ref("m-desktop-14", "tc-line", 1), ref("m-desktop-5", "tc-stack", 1))

var _ = wait("children from the family of rows and lists (tick row)", ref("s-backup-2", "tc-fset", 1))

var _ = wait("mockup defect for the owner: an input inside the check's button (nested controls)", ref("s-backup-2", "tc-check", 4))
