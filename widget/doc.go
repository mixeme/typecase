// Package widget prints the mockup's fragments from typed input (L-3).
//
// Every widget prints the fragment of its frame in docs/visual — the same
// tags, classes and text — and the library's CSS is the mockup's CSS cut
// verbatim, so the same tree under the same CSS gives the same picture. A
// container declares what it takes by the types of its fields: a slot is an
// interface with an unexported method, implemented by exactly the widgets the
// mockup puts there (L-7). Each widget's comment follows L-3.2; its «Эталон»
// names the frame and fragment it prints.
package widget
