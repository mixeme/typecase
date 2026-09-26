package framegen

import (
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// enum is one constant of a widget enum and the classes it prints.
type enum struct {
	name    string
	classes []string
}

// enums maps an enum type of the widget package to its constants, read from
// the source: the iota block that declares them and the class() method's
// string array that prints them, in the same order.
type enums map[string][]enum

var (
	constBlock = regexp.MustCompile(`(?s)const \(\n(.*?)\n\)`)
	iotaLine   = regexp.MustCompile(`^(\w+)\s+(\w+)\s*=\s*iota`)
	nameLine   = regexp.MustCompile(`^(\w+)$`)
	classArray = regexp.MustCompile(`(?s)func \(\w+ (\w+)\) class\(\) string \{\s*return \[\.\.\.\]string\{(.*?)\}\[\w+\]`)
	quoted     = regexp.MustCompile(`"([^"]*)"`)
)

// loadEnums reads the enums of the widget sources in dir.
func loadEnums(dir string) (enums, error) {
	paths, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		return nil, err
	}
	sort.Strings(paths)
	consts := map[string][]string{}
	arrays := map[string][]string{}
	for _, p := range paths {
		if strings.HasSuffix(p, "_test.go") {
			continue
		}
		b, err := os.ReadFile(p)
		if err != nil {
			return nil, err
		}
		src := strings.ReplaceAll(string(b), "\r\n", "\n")
		for _, m := range constBlock.FindAllStringSubmatch(src, -1) {
			cur := ""
			var names []string
			for _, ln := range strings.Split(m[1], "\n") {
				ln = strings.TrimSpace(strings.SplitN(ln, "//", 2)[0])
				if ln == "" {
					continue
				}
				if mm := iotaLine.FindStringSubmatch(ln); mm != nil {
					if cur != "" && len(names) > 0 {
						if _, ok := consts[cur]; !ok {
							consts[cur] = names
						}
					}
					cur, names = mm[2], []string{mm[1]}
					continue
				}
				if mm := nameLine.FindStringSubmatch(ln); mm != nil && cur != "" {
					names = append(names, mm[1])
				}
			}
			if cur != "" && len(names) > 0 {
				if _, ok := consts[cur]; !ok {
					consts[cur] = names
				}
			}
		}
		for _, m := range classArray.FindAllStringSubmatch(src, -1) {
			var strs []string
			for _, s := range quoted.FindAllStringSubmatch(m[2], -1) {
				strs = append(strs, s[1])
			}
			arrays[m[1]] = strs
		}
	}
	out := enums{}
	for t, strs := range arrays {
		names := consts[t]
		if len(names) == 0 || len(names) != len(strs) {
			continue
		}
		for i, n := range names {
			out[t] = append(out[t], enum{name: n, classes: strings.Fields(strs[i])})
		}
	}
	return out, nil
}

// of is the constant of enum t whose classes are all on e, the one with the
// most classes; "" when none is.
func (en enums) of(e *El, t string) string {
	best, bestn := "", 0
	cls := map[string]bool{}
	for _, c := range e.Cls() {
		cls[c] = true
	}
	for _, c := range en[t] {
		if len(c.classes) == 0 || len(c.classes) <= bestn {
			continue
		}
		all := true
		for _, x := range c.classes {
			if !cls[x] {
				all = false
				break
			}
		}
		if all {
			best, bestn = "w."+c.name, len(c.classes)
		}
	}
	return best
}

// faces maps the classes of a Face ("tc-dim--accent tc-dim--t12") to its
// constant.
func (en enums) faces() map[string]string {
	out := map[string]string{}
	for _, c := range en["Face"] {
		out[strings.Join(c.classes, " ")] = c.name
	}
	return out
}
