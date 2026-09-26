// Package cut takes the product's parts out of the mockup verbatim: the
// content of its one <style data-tc> block, which is the library's
// css/typecase.css, and the symbols of its sprite, which are the library's
// icons/sprite.svg (plan 2, step 1). The gate css-from-mockup cuts again and
// compares byte for byte.
package cut

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// Source is the mockup copy, relative to the module root.
const Source = "docs/visual/carrel-ui-mockups.html"

// Target is the library stylesheet the CSS cut is written to.
const Target = "css/typecase.css"

// SpriteTarget is the library sprite the symbols are written to.
const SpriteTarget = "icons/sprite.svg"

const open = "<style data-tc>"

// Cut returns the text between <style data-tc> and its </style>, unchanged.
func Cut(html string) (string, error) {
	i := strings.Index(html, open)
	if i < 0 {
		return "", errors.New("cut: no <style data-tc> in the mockup")
	}
	if strings.Contains(html[i+len(open):], open) {
		return "", errors.New("cut: more than one <style data-tc> in the mockup")
	}
	body := html[i+len(open):]
	j := strings.Index(body, "</style>")
	if j < 0 {
		return "", errors.New("cut: <style data-tc> is not closed")
	}
	return body[:j], nil
}

// spriteHead wraps the symbols. The mockup's own <svg> carries a style
// attribute; a page under a strict CSP must not, so the library's sprite is
// hidden by its size alone.
const spriteHead = `<svg xmlns="http://www.w3.org/2000/svg" width="0" height="0" aria-hidden="true"><defs>`
const spriteTail = "</defs></svg>\n"

// Sprite returns the mockup's sprite: the content of its one <defs>,
// unchanged, in the library's wrapper.
func Sprite(html string) (string, error) {
	i := strings.Index(html, "<defs>")
	if i < 0 {
		return "", errors.New("cut: no <defs> sprite in the mockup")
	}
	if strings.Contains(html[i+len("<defs>"):], "<defs>") {
		return "", errors.New("cut: more than one <defs> in the mockup")
	}
	body := html[i+len("<defs>"):]
	j := strings.Index(body, "</defs>")
	if j < 0 {
		return "", errors.New("cut: <defs> is not closed")
	}
	return spriteHead + body[:j] + spriteTail, nil
}

func read(root string) (string, error) {
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(Source)))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// FromRoot reads the mockup under the module root and returns the CSS cut.
func FromRoot(root string) (string, error) {
	html, err := read(root)
	if err != nil {
		return "", err
	}
	return Cut(html)
}

// SpriteFromRoot reads the mockup under the module root and returns the sprite.
func SpriteFromRoot(root string) (string, error) {
	html, err := read(root)
	if err != nil {
		return "", err
	}
	return Sprite(html)
}
