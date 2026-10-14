package typecase

import (
	"embed"
	"io/fs"

	"carrel.dev/typecase/css"
)

//go:embed js/*
var jsFS embed.FS

//go:embed fonts/*
var fontFS embed.FS

//go:embed icons/*
var iconFS embed.FS

// CSS returns the concatenated stylesheet built from css/ in declared order.
func CSS() (string, error) {
	return css.Bundle()
}

// JS returns the library behavior script (tc.js).
func JS() ([]byte, error) {
	return jsFS.ReadFile("js/tc.js")
}

// LiveJS returns the library's live layer (tc-live.js, L-1): requests to
// the addresses the markup names, for the library's own areas. A page that
// wants it loads it after tc.js.
func LiveJS() ([]byte, error) {
	return jsFS.ReadFile("js/tc-live.js")
}

// ThemeJS returns the theme boot script that must run before first paint (L-5.3).
func ThemeJS() ([]byte, error) {
	return jsFS.ReadFile("js/theme.js")
}

// WorkerJS returns the library's worker for a page installed as an
// application (tc-sw.js, Carrel's task 91): it keeps the library's own files
// and the page shown offline, nothing else.
func WorkerJS() ([]byte, error) {
	return jsFS.ReadFile("js/tc-sw.js")
}

// Fonts returns the embedded font files.
func Fonts() fs.FS {
	return fontFS
}

// Icons returns the embedded icon sprite.
func Icons() fs.FS {
	return iconFS
}

// Sprite returns the neutral icon sprite SVG bytes.
func Sprite() ([]byte, error) {
	return iconFS.ReadFile("icons/sprite.svg")
}
