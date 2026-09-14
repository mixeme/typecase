package main

import (
	"regexp"
	"strconv"
	"strings"
)

type decl struct {
	prop, val string
}

func parseDecls(style string) []decl {
	style = strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(style), ";"))
	if style == "" {
		return nil
	}
	var out []decl
	for _, part := range strings.Split(style, ";") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		prop, val, ok := strings.Cut(part, ":")
		if !ok {
			continue
		}
		prop = strings.ToLower(strings.TrimSpace(prop))
		val = compactVal(strings.TrimSpace(val))
		if prop == "" || val == "" {
			continue
		}
		out = append(out, decl{prop, val})
	}
	return out
}

func compactVal(v string) string {
	v = strings.Join(strings.Fields(v), " ")
	v = strings.ReplaceAll(v, ", ", ",")
	return v
}

func joinDecls(ds []decl) string {
	if len(ds) == 0 {
		return ""
	}
	var b strings.Builder
	for i, d := range ds {
		if i > 0 {
			b.WriteByte(';')
		}
		b.WriteString(d.prop)
		b.WriteByte(':')
		b.WriteString(d.val)
	}
	return b.String()
}

func takeProp(ds []decl, prop string) (val string, rest []decl, ok bool) {
	for _, d := range ds {
		if d.prop == prop {
			ok = true
			val = d.val
			continue
		}
		rest = append(rest, d)
	}
	return val, rest, ok
}

func takeProps(ds []decl, props ...string) (got map[string]string, rest []decl) {
	want := map[string]bool{}
	for _, p := range props {
		want[p] = true
	}
	got = map[string]string{}
	for _, d := range ds {
		if want[d.prop] {
			got[d.prop] = d.val
			continue
		}
		rest = append(rest, d)
	}
	return got, rest
}

var hexBgRe = regexp.MustCompile(`(?i)^#([0-9a-f]{3}|[0-9a-f]{6}|[0-9a-f]{8})$`)
var pctRe = regexp.MustCompile(`^(\d+(?:\.\d+)?)%$`)
var pxRe = regexp.MustCompile(`^(\d+(?:\.\d+)?)px$`)

func hexColour(v string) (string, bool) {
	if !hexBgRe.MatchString(v) {
		return "", false
	}
	return strings.ToUpper(v), true
}

func gapPx(v string) (int, bool) {
	m := pxRe.FindStringSubmatch(v)
	if m == nil {
		return 0, false
	}
	n, err := strconv.Atoi(m[1])
	if err != nil || n <= 0 {
		return 0, false
	}
	return n, true
}

func widthPct(v string) (string, bool) {
	if v == "0" {
		return "0%", true
	}
	if pctRe.MatchString(v) {
		return v, true
	}
	return "", false
}

func classList(v string) []string {
	return strings.Fields(v)
}

func hasClassName(classes []string, name string) bool {
	for _, c := range classes {
		if c == name {
			return true
		}
	}
	return false
}

func addClass(classes []string, name string) []string {
	if name == "" || hasClassName(classes, name) {
		return classes
	}
	return append(classes, name)
}

func primaryClass(classes []string) string {
	var prim, other string
	for _, c := range classes {
		if !strings.HasPrefix(c, "tc-") {
			continue
		}
		if strings.Contains(c, "__") || strings.Contains(c, "--") {
			if isGlueName(c) {
				continue
			}
			return c
		}
		if isGlueName(c) {
			continue
		}
		if sharedPrimitive[c] {
			if prim == "" {
				prim = c
			}
			continue
		}
		other = c
	}
	if other != "" {
		return other
	}
	return prim
}

var sharedPrimitive = map[string]bool{
	"tc-btn": true, "tc-in": true, "tc-head": true, "tc-sheet": true,
	"tc-card": true, "tc-menu": true, "tc-mark": true, "tc-row": true,
	"tc-root": true, "tc-shell": true, "tc-main": true, "tc-side": true,
	"tc-top": true, "tc-bar": true, "tc-list": true, "tc-table": true,
	"tc-field": true, "tc-form": true, "tc-nav": true, "tc-drop": true,
	"tc-acts": true, "tc-mono": true, "tc-badge": true, "tc-tag": true,
	"tc-avatar": true, "tc-meter": true, "tc-runbar": true, "tc-msg": true,
	"tc-tickrow": true, "tc-rubric": true, "tc-q": true, "tc-check": true,
	"tc-seg": true, "tc-kbd": true, "tc-close": true, "tc-brand": true,
}

func isTableCol(tag string, classes []string) bool {
	return tag == "th" || tag == "td" || hasClassName(classes, "tc-table__n3") || hasClassName(classes, "tc-table__n4")
}
