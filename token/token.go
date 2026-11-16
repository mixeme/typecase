package token

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Palette names the twelve overridable palette slots (L-5.1). The dark theme
// redefines exactly these.
var Palette = []string{
	"--tc-ink",
	"--tc-surface",
	"--tc-bg",
	"--tc-card",
	"--tc-rule",
	"--tc-muted",
	"--tc-accent",
	"--tc-danger",
	"--tc-progress", // work under way: a count still running (the owner, 16.11)
	"--tc-edge",
	"--tc-hair",
	"--tc-dusty",
}

// Typography names the three font families (L-5.1).
var Typography = []string{
	"--tc-font-serif",
	"--tc-font-sans",
	"--tc-font-mono",
}

// Spacing names the page-level horizontal padding (L-5.1).
var Spacing = []string{
	"--tc-page-gutter",
	"--tc-page-gutter-narrow",
}

// Layout names the layout tokens. They are geometry and not overridden
// (L-5.2), but the stylesheet must define them.
var Layout = []string{
	"--tc-rail",
	"--tc-side",
	"--tc-top",
}

// Data names the properties that carry data, not design: the script sets
// them from data-tc-colour and data-tc-width (Р-18), so the stylesheet uses
// them without defining them.
var Data = []string{
	"--tc-colour",
	"--tc-width",
}

// Required is every custom property the stylesheet must define.
var Required = concat(Palette, Typography, Spacing, Layout)

var (
	declRe = regexp.MustCompile(`(--tc-[a-z0-9-]+)\s*:`)
	useRe  = regexp.MustCompile(`var\(\s*(--tc-[a-z0-9-]+)`)
)

// Defined returns every --tc-* name assigned a value in css.
func Defined(css string) map[string]struct{} {
	seen := make(map[string]struct{})
	for _, m := range declRe.FindAllStringSubmatch(css, -1) {
		seen[m[1]] = struct{}{}
	}
	return seen
}

// Used returns every --tc-* name read with var() in css.
func Used(css string) map[string]struct{} {
	seen := make(map[string]struct{})
	for _, m := range useRe.FindAllStringSubmatch(css, -1) {
		seen[m[1]] = struct{}{}
	}
	return seen
}

// Missing reports required names absent from css.
func Missing(css string) []string {
	defined := Defined(css)
	var missing []string
	for _, name := range Required {
		if _, ok := defined[name]; !ok {
			missing = append(missing, name)
		}
	}
	return missing
}

// Undefined reports names css reads but never defines, other than Data.
func Undefined(css string) []string {
	defined := Defined(css)
	data := make(map[string]bool, len(Data))
	for _, name := range Data {
		data[name] = true
	}
	var out []string
	for name := range Used(css) {
		if _, ok := defined[name]; !ok && !data[name] {
			out = append(out, name)
		}
	}
	sort.Strings(out)
	return out
}

// Validate reports an error when css lacks a required token or reads one it
// never defines: a silently missing value is a difference found on screen.
func Validate(css string) error {
	if missing := Missing(css); len(missing) > 0 {
		return fmt.Errorf("typecase: missing required tokens: %s", strings.Join(missing, ", "))
	}
	if undef := Undefined(css); len(undef) > 0 {
		return fmt.Errorf("typecase: tokens read but never defined: %s", strings.Join(undef, ", "))
	}
	return nil
}

func concat(parts ...[]string) []string {
	var out []string
	for _, p := range parts {
		out = append(out, p...)
	}
	return out
}
