package framegen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"carrel.dev/typecase/internal/fragment"
	"golang.org/x/net/html"
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
	root2, err := frameRoot(string(mock), frameID)
	if err != nil {
		return "", err
	}
	found := occurrences(root2, class)
	if n < 1 || n > len(found) {
		return "", fmt.Errorf("%s: %d elements of .%s, not %d", frameID, len(found), class, n)
	}
	e := found[n-1]
	// The record is the block as it prints by itself: the comparison takes
	// the root's mixins (tc-grow, tc-push, tc-js…) off the frame's side, so
	// the draft leaves them out too.
	var own []string
	for _, c := range e.Cls() {
		if fragment.OwnClass(class, c) {
			own = append(own, c)
		}
	}
	e.attrs["class"] = strings.Join(own, " ")
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
	case e.Has("tc-split"):
		return g.split(e), ""
	case e.Has("tc-overlay"):
		return g.overlay(e), ""
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
	case e.Has("tc-checkform"):
		return g.checkForm(e), ""
	case e.Has("tc-barform"):
		return g.barItem(e), ""
	case e.Has("tc-drop"):
		return g.drop(e), ""
	case e.Has("tc-range"):
		return g.barItem(e), ""
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
	case e.Has("tc-tile"):
		return g.tile(e), ""
	}
	return g.blockChild(e), ""
}

// blockIs says whether class is e's block, by fragment's rule: the n-th of a
// class is counted as the comparison counts it, not every element that
// carries the class as a mixin or as another block's layout (tc-line,
// tc-stack, tc-push). Counted otherwise, a redraft took another element.
func blockIs(e *El, class string) bool {
	n := &html.Node{Type: html.ElementNode, Data: e.Tag, Attr: []html.Attribute{{Key: "class", Val: e.attrs["class"]}}}
	return fragment.Block(n) == class
}

// frameRoot is a frame's drawing, parsed, without its label.
func frameRoot(mock, frameID string) (*El, error) {
	markup, err := frameMarkup(mock, frameID)
	if err != nil {
		return nil, err
	}
	root := Parse(strings.TrimSpace(markup))
	unwrap(root)
	return root, nil
}

// occurrences are the elements of block class in a drawing, in document
// order, counted as fragment counts them (fragment.Ref's N).
func occurrences(root *El, class string) []*El {
	var found []*El
	var walk func(e *El)
	walk = func(e *El) {
		if e.Has(class) && (class == "tc-root" || blockIs(e, class)) {
			found = append(found, e)
		}
		if e.Has("mk-titlebar") || e.Has("mk-status") || e.Has("mk-toast") {
			return // the desktop's chrome: fragment does not count in it
		}
		for _, k := range e.Els() {
			walk(k)
		}
	}
	walk(root)
	return found
}
