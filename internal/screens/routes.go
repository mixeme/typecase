// Package screens holds the demo screens: each is a frame of the mockup
// (docs/visual/carrel-ui-mockups.html) built from widgets, top to bottom, with
// the frame's data and nothing else. A route is /f/<data-frame>; which frames
// are routes is decision Р-16 (docs/development.md).
//
// The screens are generated from the mockup by cmd/frame2go; edit the
// mockup or the generator, not them. TestGenerated wants them current.
package screens

//go:generate go run ../../cmd/frame2go -root ../..

import (
	_ "embed"

	w "carrel.dev/typecase/widget"
)

// PageCSS is the demo page's own look: the bar above a screen and the list.
//
//go:embed screens.css
var PageCSS string

// Screen is what a route prints: a widget.
type Screen interface{ Render() string }

// Route is a frame built as a screen.
type Route struct {
	Frame   string // the frame's data-frame
	Caption string // the frame's caption in the mockup
	Theme   string // the theme the frame is drawn in ("dark"); the page sets it, as theme.js does
	Screen  func() Screen
}

// Find returns the route of a frame.
func Find(frame string) (Route, bool) {
	for _, r := range Routes {
		if r.Frame == frame {
			return r, true
		}
	}
	return Route{}, false
}

func ptr(t w.Ticked) *w.Ticked { return &t }

func faceOf(f w.Face) *w.Face { return &f }
