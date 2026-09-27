// Package css serves the library stylesheet: the fonts, then the mockup's
// product CSS cut out verbatim (L-11.1).
package css

//go:generate go run ../internal/mockcss/cut/cmd/cutcss -root ..

import (
	"embed"
	"fmt"
	"io/fs"
	"strings"

	"carrel.dev/typecase/token"
)

//go:embed *.css
var files embed.FS

// Order is the concatenation order: @font-face first, then typecase.css,
// which is the mockup's <style data-tc> block byte for byte.
var Order = []string{
	"fonts.css",
	"typecase.css",
}

// Bundle reads fragments in Order, joins them, and validates the tokens.
func Bundle() (string, error) {
	var b strings.Builder
	for _, name := range Order {
		data, err := files.ReadFile(name)
		if err != nil {
			return "", fmt.Errorf("css: read %s: %w", name, err)
		}
		b.Write(data)
		if !strings.HasSuffix(string(data), "\n") {
			b.WriteByte('\n')
		}
	}
	out := b.String()
	if err := token.Validate(out); err != nil {
		return "", err
	}
	return out, nil
}

// MustBundle is Bundle that panics on error; for init-time wiring only.
func MustBundle() string {
	s, err := Bundle()
	if err != nil {
		panic(err)
	}
	return s
}

// Files returns the embedded CSS filesystem for static serving.
func Files() fs.FS {
	return files
}
