package gate

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"carrel.dev/typecase/internal/fragment"
	"carrel.dev/typecase/internal/screens"
)

var gapRow = regexp.MustCompile(`(?m)^\| ([a-z0-9-]+) \|`)

// CheckFrameCoverage fails unless every frame that is a route (Р-16: no
// mark «компонент», «фрагмент», «список», «меню», «кандидат», «отложена», and not one of
// the app's menus) has a screen or a line in docs/visual/gaps.md — one of the
// two — and every screen and every gap is such a frame.
func CheckFrameCoverage(p ModulePaths) error {
	m, err := fragment.Load(p.Root)
	if err != nil {
		return fmt.Errorf("frame-coverage: %w", err)
	}
	want := map[string]bool{}
	for _, f := range m.Frames {
		if f.Route() {
			want[f.ID] = true
		}
	}
	data, err := os.ReadFile(filepath.Join(p.Root, "docs", "visual", "gaps.md"))
	if err != nil {
		return fmt.Errorf("frame-coverage: %w", err)
	}
	return frameCoverage(want, screenFrames(), gapFrames(string(data)))
}

func frameCoverage(want map[string]bool, routes, gaps []string) error {
	var bad []string
	seen := map[string]string{}
	for _, list := range []struct {
		name string
		ids  []string
	}{{"screen", routes}, {"gap", gaps}} {
		for _, id := range list.ids {
			if prev, ok := seen[id]; ok {
				bad = append(bad, fmt.Sprintf("%s: a %s and a %s", id, prev, list.name))
				continue
			}
			seen[id] = list.name
			if !want[id] {
				bad = append(bad, fmt.Sprintf("%s: a %s, but not a route frame of the mockup", id, list.name))
			}
		}
	}
	for id := range want {
		if _, ok := seen[id]; !ok {
			bad = append(bad, id+": a route frame with neither a screen nor a line in gaps.md")
		}
	}
	if len(bad) > 0 {
		sort.Strings(bad)
		return fmt.Errorf("frame-coverage: %s", strings.Join(bad, "; "))
	}
	return nil
}

func screenFrames() []string {
	var out []string
	for _, r := range screens.Routes {
		out = append(out, r.Frame)
	}
	return out
}

func gapFrames(md string) []string {
	var out []string
	for _, m := range gapRow.FindAllStringSubmatch(md, -1) {
		out = append(out, m[1])
	}
	return out
}
