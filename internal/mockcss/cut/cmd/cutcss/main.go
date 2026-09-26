// Command cutcss writes the mockup's <style data-tc> block to css/typecase.css.
// It runs from go generate ./css.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"carrel.dev/typecase/internal/mockcss/cut"
)

func main() {
	root := flag.String("root", ".", "module root")
	flag.Parse()
	css, err := cut.FromRoot(*root)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	target := filepath.Join(*root, filepath.FromSlash(cut.Target))
	if err := os.WriteFile(target, []byte(css), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
