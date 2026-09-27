package gate

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

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
	if err := compareCut(cut.SpriteTarget, wantSprite, string(gotSprite)); err != nil {
		return err
	}
	return CheckNoForeignCSS(p.Root)
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

// selectorSubject is the last compound of a selector: what the rule styles.
var (
	cssCommentRe = regexp.MustCompile(`(?s)/\*.*?\*/`)
	cssRuleRe    = regexp.MustCompile(`([^{}@;]+)\{`)
	parenRe      = regexp.MustCompile(`\([^()]*\)`)
	tcSubjectRe  = regexp.MustCompile(`\.tc-[a-z0-9_-]+`)
)

// foreignUICSS returns the rules of a stylesheet whose subject is a class of
// the library: CSS of the interface written outside css/.
func foreignUICSS(css string) []string {
	css = cssCommentRe.ReplaceAllString(css, "")
	var bad []string
	for _, m := range cssRuleRe.FindAllStringSubmatch(css, -1) {
		list := m[1]
		for parenRe.MatchString(list) {
			list = parenRe.ReplaceAllString(list, "")
		}
		for _, sel := range strings.Split(list, ",") {
			sel = strings.TrimSpace(sel)
			parts := strings.Fields(strings.NewReplacer(">", " ", "+", " ", "~", " ").Replace(sel))
			if len(parts) == 0 {
				continue
			}
			if tcSubjectRe.MatchString(parts[len(parts)-1]) {
				bad = append(bad, strings.TrimSpace(m[1]))
				break
			}
		}
	}
	return bad
}

// CheckNoForeignCSS fails when a stylesheet of the module outside css/ styles
// a widget of the library: the interface's CSS is the mockup's cut and only
// that. Pages that show widgets — the catalog, the demo
// screens — may lay out their own boxes around them, not the widgets.
func CheckNoForeignCSS(root string) error {
	var bad []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(root, path)
		rel = filepath.ToSlash(rel)
		if d.IsDir() {
			switch {
			case rel == ".git", rel == "css", rel == "docs", strings.HasSuffix(rel, "/testdata"), rel == "node_modules":
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(rel, ".css") {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		for _, r := range foreignUICSS(string(data)) {
			bad = append(bad, rel+": "+r)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("css-from-mockup: %w", err)
	}
	if len(bad) > 0 {
		return fmt.Errorf("css-from-mockup: interface CSS outside css/: %s", strings.Join(bad, "; "))
	}
	return nil
}
