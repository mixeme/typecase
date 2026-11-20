// Package refs is the table of references: each record is a widget value
// filled with the data of its frame (the fixture) and the Ref of the mockup
// fragment it must print. TestWidgetFragments compares every
// record; the families fill the table on step 4.
package refs

import "carrel.dev/typecase/internal/fragment"

// Renderer is what every widget is: something that prints markup.
type Renderer interface {
	Render() string
}

// Entry is one reference record: one form of one block. Family and State
// name its catalog cell; the gate catalog wants every state a widget's
// comment lists («Состояния») to have one.
type Entry struct {
	Family  string
	State   string
	Fixture Renderer
	Ref     fragment.Ref
	Pos     Pos
}

// Pos is the file and line where a record is written, for the record tool
// (internal/framegen/cmd/record): a record has no number to be found by.
type Pos struct {
	File string
	Line int
}

// Entries is the table, family by family (one file per family).
var Entries []Entry

// add appends records; family files call it from a package-level var.
func add(es ...Entry) bool {
	Entries = append(Entries, es...)
	return true
}
