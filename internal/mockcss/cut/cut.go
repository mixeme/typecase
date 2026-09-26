// Package cut takes the product CSS out of the mockup: the content of its one
// <style data-tc> block, verbatim. The library's css/typecase.css is that cut
// and nothing else (plan 2, step 1); the gate css-from-mockup cuts again and
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

// Target is the library stylesheet the cut is written to, relative to the module root.
const Target = "css/typecase.css"

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

// FromRoot reads the mockup under the module root and returns the cut.
func FromRoot(root string) (string, error) {
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(Source)))
	if err != nil {
		return "", err
	}
	return Cut(string(data))
}
