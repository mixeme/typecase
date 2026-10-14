// Command record prints the Go expression of an element of a frame — the
// value of a reference record in internal/refs, drafted from the frame:
//
//	go run ./internal/framegen/cmd/record s-contacts-6 tc-danger 1
//
// With -check it prints nothing and fails when the record that points at
// the element would change its state by being redrafted: the element's own
// classes on its root (variants, is-…) differ from what the record's widget
// prints. refresh.py asks it before it writes (the review of 14.10).
package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"carrel.dev/typecase/internal/fragment"
	"carrel.dev/typecase/internal/framegen"
	"carrel.dev/typecase/internal/refs"
)

func main() {
	args := os.Args[1:]
	check := len(args) > 0 && args[0] == "-check"
	if check {
		args = args[1:]
	}
	if len(args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: record [-check] <data-frame> <block class> <n>")
		os.Exit(2)
	}
	n, err := strconv.Atoi(args[2])
	if err == nil && check {
		err = guard(fragment.Ref{Frame: args[0], Class: args[1], N: n})
	} else if err == nil {
		var expr string
		expr, err = framegen.Record(".", args[0], args[1], n)
		if err == nil {
			fmt.Println(expr)
			return
		}
	}
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, "record:", err)
	os.Exit(1)
}

// guard fails when the frame's element at r and the record's widget differ
// in the root's own classes.
func guard(r fragment.Ref) error {
	m, err := fragment.Load(".")
	if err != nil {
		return err
	}
	want, err := m.RootState(r)
	if err != nil {
		return err
	}
	for _, e := range refs.Entries {
		if e.Ref != r {
			continue
		}
		got, err := fragment.RenderedState(r.Class, e.Fixture.Render())
		if err != nil {
			return err
		}
		slices.Sort(want)
		slices.Sort(got)
		if !slices.Equal(want, got) {
			return fmt.Errorf("%s: the record %q is [%s], the frame's element is [%s] — redrafting would change what it stands for; renumber it or redraft by hand (--force)",
				r, e.State, strings.Join(got, " "), strings.Join(want, " "))
		}
		return nil
	}
	return nil // no record yet: add.py's work
}
