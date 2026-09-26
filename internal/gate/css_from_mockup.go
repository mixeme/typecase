package gate

import (
	"fmt"
	"os"
	"path/filepath"

	"carrel.dev/typecase/internal/mockcss/cut"
)

// CheckCSSFromMockup fails when css/typecase.css is not the mockup's
// <style data-tc> block byte for byte (Р-20), or icons/sprite.svg not its
// sprite. They are the generated files in the repository; this gate is why
// they cannot drift unnoticed.
func CheckCSSFromMockup(p ModulePaths) error {
	want, err := cut.FromRoot(p.Root)
	if err != nil {
		return fmt.Errorf("css-from-mockup: %w", err)
	}
	got, err := os.ReadFile(filepath.Join(p.Root, filepath.FromSlash(cut.Target)))
	if err != nil {
		return fmt.Errorf("css-from-mockup: %w", err)
	}
	if err := compareCut(cut.Target, want, string(got)); err != nil {
		return err
	}
	wantSprite, err := cut.SpriteFromRoot(p.Root)
	if err != nil {
		return fmt.Errorf("css-from-mockup: %w", err)
	}
	gotSprite, err := os.ReadFile(filepath.Join(p.Root, filepath.FromSlash(cut.SpriteTarget)))
	if err != nil {
		return fmt.Errorf("css-from-mockup: %w", err)
	}
	return compareCut(cut.SpriteTarget, wantSprite, string(gotSprite))
}

func compareCut(target, want, got string) error {
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
	return fmt.Errorf("css-from-mockup: %s differs from its cut of the mockup at byte %d (line %d); run go generate ./css",
		target, at, line)
}
