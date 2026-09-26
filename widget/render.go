package widget

import (
	"html"
	"strconv"
	"strings"
)

// mix is what a container puts on a child's root: another block (tc-push,
// tc-grow), its own element (tc-bar__more, tc-sec__acts), the ranks is-2nd and is-3rd. A
// widget does not choose its mixins; the container that holds it prints them.
type mix []string

// part is every widget: it prints its fragment with its container's mixins.
// Render is the same with none.
type part interface {
	html(m mix) string
}

func esc(s string) string { return html.EscapeString(s) }

// attr writes name="value", or nothing when value is empty.
func attr(name, value string) string {
	if value == "" {
		return ""
	}
	return " " + name + `="` + esc(value) + `"`
}

// flag writes a boolean attribute when on.
func flag(name string, on bool) string {
	if on {
		return " " + name
	}
	return ""
}

// class writes class="…" from the widget's own classes and the mixins,
// skipping empty ones.
func class(m mix, own ...string) string {
	var parts []string
	for _, c := range own {
		if c != "" {
			parts = append(parts, c)
		}
	}
	for _, c := range m {
		if c != "" {
			parts = append(parts, c)
		}
	}
	if len(parts) == 0 {
		return ""
	}
	return ` class="` + strings.Join(parts, " ") + `"`
}

// is returns the state class when on.
func is(state string, on bool) string {
	if on {
		return "is-" + state
	}
	return ""
}

// with appends mixins.
func (m mix) with(more ...string) mix {
	out := append(mix{}, m...)
	return append(out, more...)
}

// Place is where a widget sits in its row. The screen's author chooses it,
// as the frame's author did: which button goes to the end, which field
// takes the rest of the line.
type Place int

const (
	InFlow Place = iota // where the row puts it
	Pushed              // pushed to the row's end (tc-push)
	Grown               // takes the rest of the row (tc-grow)
)

func (p Place) class() string {
	return [...]string{"", "tc-push", "tc-grow"}[p]
}

// rank is is-2nd for a secondary item: the first to go under «⋯» when the
// bar narrows.
func rank(secondary bool) string {
	return pick(secondary, "is-2nd")
}

// rank3 is is-3rd for a tertiary item: it goes under «⋯» on a phone, after
// the secondary ones; in the menu of «⋯» its twin with is-3rd shows there only.
func rank3(tertiary bool) string {
	return pick(tertiary, "is-3rd")
}

func itoa(n int) string { return strconv.Itoa(n) }
