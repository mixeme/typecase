package framegen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Record is the Go expression of the n-th element of block class in a frame
// (counted from 1, the frame's label left out): the value a reference record
// of internal/refs carries, drafted from the frame as a screen is.
func Record(root, frameID, class string, n int) (string, error) {
	mock, err := os.ReadFile(filepath.Join(root, "docs", "visual", "carrel-ui-mockups.html"))
	if err != nil {
		return "", err
	}
	en, err := loadEnums(filepath.Join(root, "widget"))
	if err != nil {
		return "", err
	}
	g := &gen{en: en, faces: en.faces()}
	markup, err := frameMarkup(string(mock), frameID)
	if err != nil {
		return "", err
	}
	root2 := Parse(strings.TrimSpace(markup))
	unwrap(root2)
	var found []*El
	var walk func(e *El)
	walk = func(e *El) {
		if e.Has(class) {
			found = append(found, e)
		}
		for _, k := range e.Els() {
			walk(k)
		}
	}
	walk(root2)
	if n < 1 || n > len(found) {
		return "", fmt.Errorf("%s: %d elements of .%s, not %d", frameID, len(found), class, n)
	}
	e := found[n-1]
	expr, why := g.element(e)
	if why != "" {
		return "", fmt.Errorf("%s .%s #%d: %s", frameID, class, n, why)
	}
	return expr, nil
}

// frameMarkup is a frame's drawing without its label.
func frameMarkup(mock, id string) (string, error) {
	its := frameOpen.FindAllStringSubmatchIndex(mock, -1)
	for k, m := range its {
		if mock[m[4]:m[5]] != id {
			continue
		}
		j := len(mock)
		if k+1 < len(its) {
			j = its[k+1][0]
		}
		blk := mock[m[0]:j]
		in, err := inner(blk)
		if err != nil {
			return "", err
		}
		if lab := frameLab.FindString(blk); lab != "" {
			if i := strings.Index(in, lab); i >= 0 {
				in = in[i+len(lab):]
			}
		}
		return in, nil
	}
	return "", fmt.Errorf("no frame %s", id)
}

// element translates one element as whatever widget its block is.
func (g *gen) element(e *El) (expr string, why string) {
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case fail:
				why = "unknown: " + v.why
			case wait:
				why = "no widget for it here"
			default:
				why = fmt.Sprintf("error: %v", r)
			}
			expr = ""
		}
	}()
	switch {
	case e.Has("tc-shell"):
		return g.shell(e), ""
	case e.Has("tc-root") && len(e.Cls()) > 0 && e.Cls()[0] == "tc-root" && !e.Has("tc-shell"):
		return g.page(e), ""
	case e.Has("tc-main"):
		return g.main(e), ""
	case e.Has("tc-side"):
		return g.side(e), ""
	case e.Has("tc-top"):
		return g.topAny(e), ""
	case e.Has("tc-sheet"):
		return g.sheet(e), ""
	case e.Has("tc-segfold"):
		return g.segfold(e), ""
	case e.Has("tc-menu"):
		return g.menu(e), ""
	case e.Has("tc-row"):
		return g.row(e), ""
	case e.Has("tc-in") || e.Tag == "textarea":
		if c := g.control(e); c != "" {
			return c, ""
		}
	case e.Has("tc-nav"), e.Has("tc-railsec"), e.Has("tc-railfoot"), e.Has("tc-badge"), e.Has("tc-tag"),
		e.Has("tc-emblem"), e.Has("tc-railstatus"), e.Has("tc-usermenu"), e.Has("tc-search"), e.Has("tc-brand"):
		return g.child(e), ""
	case e.Has("tc-tick"):
		return g.cellPart(e), ""
	case e.Has("tc-fset"):
		return g.fset(e), ""
	case e.Has("tc-avatar"):
		return g.avatar(e), ""
	case e.Has("tc-post"):
		return g.postForm(e), ""
	case e.Has("tc-mark"):
		return g.mark(e), ""
	case e.Has("tc-dim"):
		return g.dimx(e), ""
	}
	return g.blockChild(e), ""
}
