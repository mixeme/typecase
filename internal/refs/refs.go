// Package refs is the table of references: each record is a widget value
// filled with the data of its frame (the fixture) and the Ref of the mockup
// fragment it must print (plan 2, step 2). TestWidgetFragments compares every
// record; the families fill the table on step 4.
package refs

import "carrel.dev/typecase/internal/fragment"

// Renderer is what every widget is: something that prints markup.
type Renderer interface {
	Render() string
}

// Entry is one reference record: one form of one block.
type Entry struct {
	Fixture Renderer
	Ref     fragment.Ref
}

// Entries is the table, family by family (one file per family).
var Entries []Entry

// add appends records; family files call it from a package-level var.
func add(es ...Entry) bool {
	Entries = append(Entries, es...)
	return true
}
