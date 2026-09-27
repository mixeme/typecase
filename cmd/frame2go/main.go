// Command frame2go translates frames of the mockup into screens: the
// frame's widgets and data, top to bottom. It reads
// docs/visual/carrel-ui-mockups.html and the widget package's enums of the
// module at -root.
//
//	go run ./cmd/frame2go s-find-2 ...   # print the screen function of each frame
//	go run ./cmd/frame2go                # rewrite internal/screens
//	go run ./cmd/frame2go -check         # say which files differ, change nothing
//
// internal/screens is its output (go generate). A Carrel screen starts from
// the printed function: then the frame's data becomes the application's.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"carrel.dev/typecase/internal/framegen"
)

func main() {
	root := flag.String("root", ".", "the module's root")
	check := flag.Bool("check", false, "compare with internal/screens instead of writing")
	flag.Parse()
	var err error
	if flag.NArg() > 0 {
		err = print(*root, flag.Args())
	} else {
		err = write(*root, *check)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "frame2go:", err)
		os.Exit(1)
	}
}

// print prints the screen function of each frame named.
func print(root string, frames []string) error {
	secs, gaps, err := framegen.Load(root)
	if err != nil {
		return err
	}
	for _, id := range frames {
		found := false
		for _, s := range secs {
			for _, sc := range s.Screens {
				if sc.Frame == id {
					src, err := framegen.Func(sc)
					if err != nil {
						return err
					}
					fmt.Println(src)
					found = true
				}
			}
		}
		for _, g := range gaps {
			if g.Frame == id {
				return fmt.Errorf("%s «%s» is not a screen yet: %s", id, g.Caption, g.Why)
			}
		}
		if !found {
			return fmt.Errorf("%s is not a route frame of the mockup (Р-16)", id)
		}
	}
	return nil
}

// write rewrites internal/screens, or with check only lists what differs.
func write(root string, check bool) error {
	files, err := framegen.Build(root)
	if err != nil {
		return err
	}
	dir := filepath.Join(root, "internal", "screens")
	var names []string
	for n := range files {
		names = append(names, n)
	}
	sort.Strings(names)
	stale := 0
	for _, n := range names {
		p := filepath.Join(dir, n)
		old, _ := os.ReadFile(p)
		if bytes.Equal(old, files[n]) {
			continue
		}
		stale++
		if check {
			fmt.Println("differs:", p)
			continue
		}
		if err := os.WriteFile(p, files[n], 0o644); err != nil {
			return err
		}
		fmt.Println("wrote", p)
	}
	if check && stale > 0 {
		return fmt.Errorf("%d files of internal/screens are not what the mockup gives: go generate ./internal/screens", stale)
	}
	return nil
}
