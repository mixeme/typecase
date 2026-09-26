package gate

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"golang.org/x/net/html"

	"carrel.dev/typecase/internal/fragment"
	"carrel.dev/typecase/internal/screens"
)

// catalogOnly are frames without a mark that still go to the catalog, not to
// a route (Р-16): the desktop app's user menu and the tray menu.
var catalogOnly = map[string]bool{"m-desktop-2": true, "m-desktop-8": true}

var gapRow = regexp.MustCompile(`(?m)^\| ([a-z0-9-]+) \|`)

// CheckFrameCoverage fails unless every frame that is a route (Р-16: no
// mark «компонент», «фрагмент», «список», «меню», «кандидат», and not one of
// the app's menus) has a screen or a line in docs/visual/gaps.md — one of the
// two — and every screen and every gap is such a frame (plan 3, step 1).
func CheckFrameCoverage(p ModulePaths) error {
	m, err := fragment.Load(p.Root)
	if err != nil {
		return fmt.Errorf("frame-coverage: %w", err)
	}
	want := map[string]bool{}
	for _, f := range m.Frames {
		if routeFrame(f) {
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

// routeFrame says whether a frame is a route: its label carries no mark
// (span.frag) and it is not one of the menus kept for the catalog.
func routeFrame(f *fragment.Frame) bool {
	if catalogOnly[f.ID] {
		return false
	}
	for c := f.Node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && hasClassAttr(c, "frame-label") {
			for s := c.FirstChild; s != nil; s = s.NextSibling {
				if s.Type == html.ElementNode && hasClassAttr(s, "frag") {
					return false
				}
			}
		}
	}
	return true
}

func hasClassAttr(n *html.Node, cls string) bool {
	for _, a := range n.Attr {
		if a.Key == "class" {
			for _, c := range strings.Fields(a.Val) {
				if c == cls {
					return true
				}
			}
		}
	}
	return false
}
