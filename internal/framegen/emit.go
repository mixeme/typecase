package framegen

import (
	"fmt"
	"go/format"
	"regexp"
	"sort"
	"strings"
)

// splitTop splits s at top-level commas (outside braces, parens and
// strings).
func splitTop(s []rune) [][]rune {
	var parts [][]rune
	var cur []rune
	depth, instr := 0, false
	flush := func() {
		t := strings.TrimSpace(string(cur))
		parts = append(parts, []rune(t))
		cur = nil
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case instr:
			cur = append(cur, c)
			if c == '\\' {
				cur = append(cur, s[i+1])
				i++
				continue
			}
			if c == '"' {
				instr = false
			}
		case c == '"':
			instr = true
			cur = append(cur, c)
		case strings.ContainsRune("{([", c):
			depth++
			cur = append(cur, c)
		case strings.ContainsRune("})]", c):
			depth--
			cur = append(cur, c)
		case c == ',' && depth == 0:
			flush()
		default:
			cur = append(cur, c)
		}
	}
	if strings.TrimSpace(string(cur)) != "" {
		flush()
	}
	return parts
}

// matchClose is the index of the brace closing the one at s[i].
func matchClose(s []rune, i int) int {
	depth, instr := 0, false
	for j := i; j < len(s); j++ {
		c := s[j]
		if instr {
			if c == '\\' {
				j++
				continue
			}
			if c == '"' {
				instr = false
			}
			continue
		}
		switch c {
		case '"':
			instr = true
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return j
			}
		}
	}
	panic("unbalanced")
}

// breakLong breaks composite literals longer than width: one element a line.
// gofmt then indents them.
func breakLong(s []rune, width int) string {
	var out strings.Builder
	instr := false
	for i := 0; i < len(s); {
		c := s[i]
		if instr {
			out.WriteRune(c)
			if c == '\\' {
				out.WriteRune(s[i+1])
				i += 2
				continue
			}
			if c == '"' {
				instr = false
			}
			i++
			continue
		}
		if c == '"' {
			instr = true
			out.WriteRune(c)
			i++
			continue
		}
		if c == '{' {
			j := matchClose(s, i)
			body := s[i+1 : j]
			if len(body) > width {
				out.WriteString("{\n")
				for _, e := range splitTop(body) {
					out.WriteString(breakLong(e, width) + ",\n")
				}
				out.WriteString("}")
			} else {
				out.WriteString("{" + string(body) + "}")
			}
			i = j + 1
			continue
		}
		out.WriteRune(c)
		i++
	}
	return out.String()
}

const (
	nbsp       = string(rune(0xa0))
	nbspEscape = string(rune(92)) + "u00a0" // the escape of a no-break space in Go source
)

var sectionPrefix = regexp.MustCompile(`^(s|m)-`)

// fileOf is the file of a section's screens.
func fileOf(id string) string {
	n := sectionPrefix.ReplaceAllString(id, "")
	if n == "" {
		n = "misc"
	}
	return strings.ReplaceAll(n, "-", "_") + ".go"
}

var noteBodyConst = regexp.MustCompile("(?s)const noteBody = `[^`]*`")

// funcText is a screen's function, before gofmt.
func funcText(sc Screen) string {
	return fmt.Sprintf("// %s is frame «%s» (%s).\nfunc %s() Screen {\n\treturn %s\n}", sc.Name, sc.Caption, sc.Frame, sc.Name, breakLong([]rune(sc.Expr), 100))
}

// Func is a screen's function as Go source, gofmt'ed, for a screen author to
// start from.
func Func(sc Screen) (string, error) {
	src := "package screens\n\n" + funcText(sc) + "\n"
	b, err := format.Source([]byte(strings.ReplaceAll(src, nbsp, nbspEscape)))
	if err != nil {
		return "", err
	}
	return strings.TrimPrefix(string(b), "package screens\n\n"), nil
}

// Draft is a screen as a Go file of its own that builds outside Typecase: the
// package draft, the widget import, the frame's function returning what
// renders, and the helpers it uses (a pointer maker, the notes' text). A
// client's screen starts from it; refsInputs is internal/refs/inputs.go.
func Draft(sc Screen, refsInputs string) (string, error) {
	fn := funcText(sc)
	fn = strings.Replace(fn, "() Screen {", "() interface{ Render() string } {", 1)
	fn = strings.ReplaceAll(fn, "faceOf(", "ptr(")
	buf := []string{
		"// Drafted by frame2go from frame «" + sc.Caption + "» (" + sc.Frame + ") of the Typecase mockup: the",
		"// frame's widgets and its data, top to bottom. The data is the frame's; it",
		"// is to become the application's (docs/guide/screen.md).",
		"package draft", "", `import w "carrel.dev/typecase/widget"`, "", fn, ""}
	if strings.Contains(fn, "ptr(") {
		buf = append(buf, "// ptr is a pointer to a copy of v.", "func ptr[T any](v T) *T { return &v }", "")
	}
	src := strings.Join(buf, "\n")
	if strings.Contains(src, "noteBody") {
		c := noteBodyConst.FindString(refsInputs)
		if c == "" {
			return "", fmt.Errorf("no noteBody in internal/refs/inputs.go")
		}
		src += "\n// noteBody is the text of the note being edited in the notes frames.\n" + c + "\n"
	}
	b, err := format.Source([]byte(strings.ReplaceAll(src, nbsp, nbspEscape)))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Files is the source of internal/screens, gofmt'ed: a file per mockup
// section, a function per frame, and routes_list.go with the routes in
// mockup order. refsInputs is internal/refs/inputs.go, whose noteBody the
// notes frames share.
func Files(secs []Section, refsInputs string) (map[string][]byte, error) {
	type item struct {
		title string
		sc    Screen
	}
	var order []string
	files := map[string][]item{}
	var routes []Screen
	for _, s := range secs {
		fn := fileOf(s.ID)
		if _, ok := files[fn]; !ok {
			order = append(order, fn)
		}
		for _, sc := range s.Screens {
			files[fn] = append(files[fn], item{s.Title, sc})
			routes = append(routes, sc)
		}
	}
	out := map[string][]byte{}
	for _, fn := range order {
		items := files[fn]
		seen := map[string]bool{}
		var titles []string
		for _, it := range items {
			if !seen[it.title] {
				seen[it.title] = true
				titles = append(titles, it.title)
			}
		}
		sort.Strings(titles)
		buf := []string{"// Раздел макета: " + strings.Join(titles, "; ") + ".\n", "package screens", "", `import w "carrel.dev/typecase/widget"`, ""}
		for _, it := range items {
			buf = append(buf, funcText(it.sc), "")
		}
		src := strings.Join(buf, "\n")
		if strings.Contains(src, "noteBody") {
			c := noteBodyConst.FindString(refsInputs)
			if c == "" {
				return nil, fmt.Errorf("no noteBody in internal/refs/inputs.go")
			}
			src += "\n// noteBody is the text of the note being edited in the notes frames.\n" + c + "\n"
		}
		// A no-break space is invisible in source: write it as its escape.
		b, err := format.Source([]byte(strings.ReplaceAll(src, nbsp, nbspEscape)))
		if err != nil {
			return nil, fmt.Errorf("%s: %v", fn, err)
		}
		out[fn] = b
	}
	rb := []string{"package screens", "", "// Routes lists the frames of the mockup built as screens, in the mockup's", "// order: /f/<Frame> (Р-16).", "var Routes = []Route{"}
	for _, r := range routes {
		th := ""
		if r.Theme != "" {
			th = ", Theme: " + q(r.Theme)
		}
		rb = append(rb, fmt.Sprintf("\t{Frame: %s, Caption: %s%s, Screen: %s},", q(r.Frame), q(r.Caption), th, r.Name))
	}
	rb = append(rb, "}")
	b, err := format.Source([]byte(strings.Join(rb, "\n") + "\n"))
	if err != nil {
		return nil, fmt.Errorf("routes_list.go: %v", err)
	}
	out["routes_list.go"] = b
	return out, nil
}
