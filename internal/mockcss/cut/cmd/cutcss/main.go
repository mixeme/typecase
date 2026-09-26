// Command cutcss writes the mockup's <style data-tc> block to css/typecase.css
// and its sprite symbols to icons/sprite.svg. It runs from go generate ./css.
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
	if err == nil {
		err = os.WriteFile(filepath.Join(*root, filepath.FromSlash(cut.Target)), []byte(css), 0o644)
	}
	if err == nil {
		var sprite string
		sprite, err = cut.SpriteFromRoot(*root)
		if err == nil {
			err = os.WriteFile(filepath.Join(*root, filepath.FromSlash(cut.SpriteTarget)), []byte(sprite), 0o644)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
