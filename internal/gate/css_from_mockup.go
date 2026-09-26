package gate

import (
	"fmt"
	"os"
	"path/filepath"

	"carrel.dev/typecase/internal/mockcss/cut"
)

// CheckCSSFromMockup fails when css/typecase.css is not the mockup's
// <style data-tc> block byte for byte (Р-20). The file is the one generated
// file in the repository; this gate is why it cannot drift unnoticed.
func CheckCSSFromMockup(p ModulePaths) error {
	want, err := cut.FromRoot(p.Root)
	if err != nil {
		return fmt.Errorf("css-from-mockup: %w", err)
	}
	got, err := os.ReadFile(filepath.Join(p.Root, filepath.FromSlash(cut.Target)))
	if err != nil {
		return fmt.Errorf("css-from-mockup: %w", err)
	}
	return compareCut(want, string(got))
}

func compareCut(want, got string) error {
	if want == got {
		return nil
	}
	n := len(want)
	if len(got) < n {
		n = len(got)
	}
	at := n
	for i := 0; i < n; i++ {
		if want[i] != got[i] {
			at = i
			break
		}
	}
	line := 1
	for i := 0; i < at && i < len(want); i++ {
		if want[i] == '\n' {
			line++
		}
	}
	return fmt.Errorf("css-from-mockup: %s differs from the mockup's <style data-tc> at byte %d (line %d); run go generate ./css",
		cut.Target, at, line)
}
