package framegen

import (
	"fmt"
	"os"
	"path/filepath"
)

// Load translates the route frames of the mockup of the module at root.
func Load(root string) ([]Section, []Gap, error) {
	mock, err := os.ReadFile(filepath.Join(root, "docs", "visual", "carrel-ui-mockups.html"))
	if err != nil {
		return nil, nil, err
	}
	return Generate(string(mock), filepath.Join(root, "widget"))
}

// Build translates the mockup of the module at root into the files of
// internal/screens. A frame it cannot translate is an error: every route
// frame is a screen (Р-16).
func Build(root string) (map[string][]byte, error) {
	secs, gaps, err := Load(root)
	if err != nil {
		return nil, err
	}
	if len(gaps) > 0 {
		msg := ""
		for _, g := range gaps {
			msg += fmt.Sprintf("\n\t%s «%s»: %s", g.Frame, g.Caption, g.Why)
		}
		return nil, fmt.Errorf("%d frames are not screens:%s", len(gaps), msg)
	}
	refs, err := os.ReadFile(filepath.Join(root, "internal", "refs", "inputs.go"))
	if err != nil {
		return nil, err
	}
	return Files(secs, string(refs))
}
