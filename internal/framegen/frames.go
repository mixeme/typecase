package framegen

import (
	"fmt"
	"regexp"
	"strings"
)

// Screen is a route frame translated: its section, caption, theme and the
// Go expression of its screen.
type Screen struct {
	Frame   string // data-frame
	Name    string // the Go function's name
	Caption string
	Theme   string // data-tc-theme of the frame's root, or ""
	Expr    string
}

// Section is a section of the mockup and its screens, in mockup order.
type Section struct {
	ID, Title string
	Screens   []Screen
}

// Gap is a frame the generator could not translate, and why.
type Gap struct{ Frame, Caption, Why string }

// catalogOnly are the frames Р-16 leaves out of the routes.
var catalogOnly = map[string]bool{"m-desktop-2": true, "m-desktop-8": true}

var (
	frameOpen = regexp.MustCompile(`<div class="frame([^"]*)" data-frame="([^"]+)"`)
	sectionH2 = regexp.MustCompile(`<h2 id="([^"]*)"[^>]*>(.*?)</h2>`)
	frameLab  = regexp.MustCompile(`(?s)<div class="frame-label"><span>(.*?)</span>(.*?)</div>`)
	tags      = regexp.MustCompile(`<[^>]+>`)
	divTag    = regexp.MustCompile(`<(/?)div\b[^>]*>`)
	fragMark  = regexp.MustCompile(`class="frag"`)
)

type frame struct {
	id, secID, secTitle, caption, markup string
}

// frames is the route frames of the mockup: every frame but the fragments
// and the catalogue-only ones.
func frames(mock string) ([]frame, error) {
	its := frameOpen.FindAllStringSubmatchIndex(mock, -1)
	type head struct {
		at        int
		id, title string
	}
	var heads []head
	for _, m := range sectionH2.FindAllStringSubmatchIndex(mock, -1) {
		heads = append(heads, head{m[0], mock[m[2]:m[3]], tags.ReplaceAllString(mock[m[4]:m[5]], "")})
	}
	var out []frame
	for k, m := range its {
		j := len(mock)
		if k+1 < len(its) {
			j = its[k+1][0]
		}
		blk := mock[m[0]:j]
		id := mock[m[4]:m[5]]
		lab := frameLab.FindStringSubmatch(blk)
		if lab == nil {
			return nil, fmt.Errorf("frame %s has no label", id)
		}
		if fragMark.MatchString(lab[2]) || catalogOnly[id] {
			continue
		}
		var sec head
		for _, h := range heads {
			if h.at < m[0] {
				sec = h
			}
		}
		in, err := inner(blk)
		if err != nil {
			return nil, fmt.Errorf("frame %s: %v", id, err)
		}
		i := strings.Index(in, lab[0])
		out = append(out, frame{id, sec.id, sec.title, strings.TrimFunc(tags.ReplaceAllString(lab[1], ""), isSpace), in[i+len(lab[0]):]})
	}
	return out, nil
}

// inner is the frame div's content, by balancing div tags.
func inner(blk string) (string, error) {
	depth := 0
	for _, t := range divTag.FindAllStringSubmatchIndex(blk, -1) {
		if t[3] > t[2] {
			depth--
		} else {
			depth++
		}
		if depth == 0 {
			start := strings.Index(blk, ">") + 1
			return blk[start:t[0]], nil
		}
	}
	return "", fmt.Errorf("unbalanced frame")
}

// unwrap drops the desktop drawing's chrome, which is not product: the
// title bar, status line and notification go, other mk-* elements give
// their children in their place.
func unwrap(e *El) {
	var kids []any
	for _, k := range e.Kids {
		if x, ok := k.(*El); ok {
			if x.Has("mk-titlebar") || x.Has("mk-status") || x.Has("mk-toast") {
				continue
			}
			mk := false
			for _, c := range x.Cls() {
				if strings.HasPrefix(c, "mk-") {
					mk = true
				}
			}
			unwrap(x)
			if mk {
				kids = append(kids, x.Kids...)
				continue
			}
		}
		kids = append(kids, k)
	}
	e.Kids = kids
}

var nameSplit = regexp.MustCompile(`[-_]`)

// funcName is the Go name of a frame's screen: s-find-2 is Find2.
func funcName(id string) string {
	s := id
	if strings.HasPrefix(id, "s-") || strings.HasPrefix(id, "m-") {
		s = id[2:]
	}
	var b strings.Builder
	for _, p := range nameSplit.Split(s, -1) {
		if p != "" {
			b.WriteString(strings.ToUpper(p[:1]) + p[1:])
		}
	}
	return b.String()
}

// screen is the expression of a frame's root.
func (g *gen) screen(root *El) string {
	if root.Has("tc-shell") {
		return g.shell(root)
	}
	if root.Has("tc-root") {
		return g.page(root)
	}
	var kids []string
	for i, k := range root.Els() {
		if i == 2 {
			break
		}
		kids = append(kids, join(k.Cls(), " "))
	}
	stop("root " + join(root.Cls(), " ") + " > " + join(kids, ","))
	return ""
}

// translate is a frame's screen, or why it has none.
func (g *gen) translate(markup string) (expr string, why string) {
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case fail:
				why = "unknown: " + v.why
			case wait:
				why = "wait"
			default:
				why = fmt.Sprintf("error: %v", r)
			}
			expr = ""
		}
	}()
	root := Parse(strings.TrimSpace(markup))
	unwrap(root)
	return g.screen(root), ""
}

// Generate translates the route frames of the mockup, reading the widget
// package's enums from widgetDir.
func Generate(mock, widgetDir string) ([]Section, []Gap, error) {
	en, err := loadEnums(widgetDir)
	if err != nil {
		return nil, nil, err
	}
	g := &gen{en: en, faces: en.faces()}
	fs, err := frames(mock)
	if err != nil {
		return nil, nil, err
	}
	var secs []Section
	var gaps []Gap
	names := map[string]bool{}
	for _, f := range fs {
		name := funcName(f.id)
		if names[name] {
			return nil, nil, fmt.Errorf("two frames make %s", name)
		}
		names[name] = true
		expr, why := g.translate(f.markup)
		if why != "" {
			gaps = append(gaps, Gap{f.id, f.caption, why})
			continue
		}
		root := Parse(strings.TrimSpace(f.markup))
		sc := Screen{Frame: f.id, Name: name, Caption: f.caption, Theme: root.Attr("data-tc-theme"), Expr: expr}
		i := len(secs) - 1
		for ; i >= 0; i-- {
			if secs[i].ID == f.secID && secs[i].Title == f.secTitle {
				break
			}
		}
		if i < 0 {
			secs = append(secs, Section{ID: f.secID, Title: f.secTitle})
			i = len(secs) - 1
		}
		secs[i].Screens = append(secs[i].Screens, sc)
	}
	return secs, gaps, nil
}
