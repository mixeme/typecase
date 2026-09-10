package gate

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"carrel.dev/typecase/internal/screens"
	"carrel.dev/typecase/layout"
)

// CheckLayoutSlots enforces layout tables against generated CSS, hand CSS, and markup.
func CheckLayoutSlots(p ModulePaths, catalog []CatalogEntry) error {
	spec := layout.All()
	generated := layout.CSS()
	var errs []string
	var warns []string

	for _, g := range spec.Grids {
		if err := checkGeneratedGrid(generated, g); err != "" {
			errs = append(errs, err)
		}
		for _, t := range g.Tracks {
			if t.Empty == layout.EmptyOmit {
				errs = append(errs, fmt.Sprintf("EmptyOmit in Grid %s slot %s", g.Selector, t.Slot))
			}
		}
	}

	manual, atBlocks, err := readManualCSS(p.CSSDir)
	if err != nil {
		return err
	}
	owned := ownedProperties(spec)
	for sel, props := range owned {
		for _, r := range parseCSSRules(manual) {
			if !ruleConcerns(r.Selector, sel) {
				continue
			}
			got := cssProps(r.Body)
			for prop := range props {
				if _, ok := got[prop]; ok {
					errs = append(errs, fmt.Sprintf("%s sets %s in hand CSS", sel, prop))
				}
			}
		}
		for _, block := range atBlocks {
			for _, r := range parseCSSRules(innerAtBlock(block)) {
				if !ruleConcerns(r.Selector, sel) {
					continue
				}
				got := cssProps(r.Body)
				for prop := range props {
					if _, ok := got[prop]; ok {
						warns = append(warns, fmt.Sprintf("%s %s in hand @media/@container", sel, prop))
					}
				}
			}
		}
	}

	for _, e := range catalog {
		errs = append(errs, checkMarkupGrid(spec, e.HTML, e.Widget+"/"+e.State, false)...)
	}
	for _, sc := range screens.All() {
		html := sc.Body()
		errs = append(errs, checkMarkupGrid(spec, html, sc.ID, true)...)
		errs = append(errs, checkGridRowParent(spec, html, sc.ID)...)
	}
	errs = append(errs, checkGridRowFields(p)...)

	if len(warns) > 0 {
		fmt.Fprintf(os.Stderr, "layout-slots warning: %s\n", strings.Join(warns, "; "))
	}
	if len(errs) > 0 {
		if len(warns) > 0 {
			errs = append(errs, "warnings: "+strings.Join(warns, "; "))
		}
		return fmt.Errorf("layout-slots: %s", strings.Join(errs, "; "))
	}
	return nil
}

func innerAtBlock(block string) string {
	brace := strings.Index(block, "{")
	if brace < 0 || !strings.HasSuffix(block, "}") {
		return block
	}
	return block[brace+1 : len(block)-1]
}

func checkGeneratedGrid(css string, g layout.Grid) string {
	if !strings.Contains(css, g.Selector) {
		return fmt.Sprintf("generated CSS missing %s", g.Selector)
	}
	cols := generatedBaseColumns(css, g.Selector)
	if cols == "" {
		return fmt.Sprintf("generated CSS missing grid-template-columns for %s", g.Selector)
	}
	got := layout.SplitTracks(cols)
	if len(got) != len(g.Tracks) {
		return fmt.Sprintf("%s generated tracks %d want %d", g.Selector, len(got), len(g.Tracks))
	}
	return ""
}

func generatedBaseColumns(css, sel string) string {
	for _, r := range parseCSSRules(css) {
		if compactSel(r.Selector) != compactSel(sel) {
			continue
		}
		if cols := cssProps(r.Body)["grid-template-columns"]; cols != "" {
			return cols
		}
	}
	return ""
}

func readManualCSS(dir string) (body string, at []string, err error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", nil, fmt.Errorf("layout-slots: read css: %w", err)
	}
	var raw strings.Builder
	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), ".css") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			return "", nil, err
		}
		raw.Write(data)
		raw.WriteByte('\n')
	}
	stripped := stripCSSComments(raw.String())
	body, at = cutAtRules(stripped)
	return body, at, nil
}

func ownedProperties(spec layout.Spec) map[string]map[string]bool {
	out := make(map[string]map[string]bool)
	add := func(sel, prop string) {
		if sel == "" || prop == "" {
			return
		}
		m := out[sel]
		if m == nil {
			m = map[string]bool{}
			out[sel] = m
		}
		m[prop] = true
	}
	for _, g := range spec.Grids {
		add(g.Selector, "display")
		add(g.Selector, "grid-template-columns")
		add(g.Selector, "grid-template-rows")
		if g.Gap != "" {
			add(g.Selector, "gap")
		}
		if g.MinHeight != "" {
			add(g.Selector, "min-height")
		}
		if g.PadLeft != "" {
			add(g.Selector, "padding-left")
		}
		if g.PadRight != "" {
			add(g.Selector, "padding-right")
		}
		if g.Align != "" {
			add(g.Selector, "align-items")
		}
		if g.AutoFlow != "" {
			add(g.Selector, "grid-auto-flow")
		}
		if g.AutoColumns != "" {
			add(g.Selector, "grid-auto-columns")
		}
	}
	for _, box := range spec.Boxes {
		for k := range box.Props {
			add(box.Selector, k)
		}
	}
	for _, flow := range spec.Flows {
		for k := range flow.Props {
			add(flow.Selector, k)
		}
	}
	return out
}

func checkMarkupGrid(spec layout.Spec, html, where string, screens bool) []string {
	root := parseHTML(html)
	var errs []string
	for _, g := range spec.Grids {
		slotted := layout.Slotted(g)
		cols := isColsGrid(g)
		if !slotted && !cols {
			continue
		}
		cls := className(g.Selector)
		walkHTML(root, func(n *htmlNode) {
			if n.tag == "#root" || !nodeHasClass(n, cls) {
				return
			}
			kids := n.children
			want := len(g.Tracks)
			if g.AutoFlow != "" {
				if len(kids) < want {
					errs = append(errs, fmt.Sprintf("%s %s children %d want >= %d", where, g.Selector, len(kids), want))
					return
				}
				kids = kids[:want]
			} else if len(kids) != want {
				errs = append(errs, fmt.Sprintf("%s %s children %d want %d", where, g.Selector, len(kids), want))
				return
			}
			if !slotted {
				return
			}
			for i, t := range g.Tracks {
				slotCls := layout.SlotClass(t.Slot, g.Selector)
				if slotCls == "" {
					continue
				}
				if !nodeHasClass(kids[i], slotCls) {
					errs = append(errs, fmt.Sprintf("%s %s child %d class want %s", where, g.Selector, i, slotCls))
				}
			}
		})
		_ = screens
	}
	return errs
}

func isColsGrid(g layout.Grid) bool {
	return strings.HasPrefix(g.Selector, ".tc-cols--")
}

func checkGridRowParent(spec layout.Spec, html, route string) []string {
	root := parseHTML(html)
	var errs []string
	for _, g := range spec.Grids {
		if !strings.HasPrefix(g.Selector, ".tc-row--") {
			continue
		}
		cls := className(g.Selector)
		walkHTML(root, func(n *htmlNode) {
			if n.tag == "#root" || !nodeHasClass(n, cls) {
				return
			}
			if n.parent == nil || n.parent.tag == "#root" {
				return
			}
			if !nodeHasClass(n.parent, "tc-list") {
				errs = append(errs, fmt.Sprintf("%s %s parent is not .tc-list", route, cls))
			}
		})
	}
	return errs
}

// checkGridRowFields reads Row literals in widget, catalog and screens source.
//
// Blind spot (spec L-8.1): a Variant coming from a variable or a function, or
// Menu assigned after the literal (r.Menu = …), is not seen here, and Render
// drops such a menu without an error. Enough while every row is a literal in
// the catalog or a demo screen; revisit when application code exists.
func checkGridRowFields(p ModulePaths) []string {
	gridByVariant := map[string]layout.Grid{}
	for _, g := range layout.All().Grids {
		if strings.HasPrefix(g.Selector, ".tc-row--") {
			gridByVariant[strings.TrimPrefix(g.Selector, ".tc-row--")] = g
		}
	}
	if len(gridByVariant) == 0 {
		return nil
	}
	var errs []string
	dirs := []string{
		p.WidgetDir,
		filepath.Join(p.Root, "internal", "catalog"),
		filepath.Join(p.Root, "internal", "screens"),
	}
	for _, dir := range dirs {
		if _, err := os.Stat(dir); err != nil {
			continue
		}
		fset := token.NewFileSet()
		pkgs, err := parser.ParseDir(fset, dir, func(fi os.FileInfo) bool {
			return strings.HasSuffix(fi.Name(), ".go") && !strings.HasSuffix(fi.Name(), "_test.go")
		}, 0)
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		for _, pkg := range pkgs {
			for _, f := range pkg.Files {
				ast.Inspect(f, func(n ast.Node) bool {
					cl, ok := n.(*ast.CompositeLit)
					if !ok {
						return true
					}
					if !isRowType(cl.Type) {
						return true
					}
					variant := rowVariantValue(cl)
					g, ok := gridByVariant[variant]
					if !ok {
						return true
					}
					if fieldPresentNonEmpty(cl, "Menu") {
						errs = append(errs, fmt.Sprintf("Menu on grid row %s", variant))
					}
					if fieldPresentNonEmpty(cl, "Badges") && !layout.HasSlot(g, "tags") {
						errs = append(errs, fmt.Sprintf("Badges on grid row %s without tags slot", variant))
					}
					return true
				})
			}
		}
	}
	return errs
}

func isRowType(expr ast.Expr) bool {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name == "Row"
	case *ast.SelectorExpr:
		return t.Sel.Name == "Row"
	}
	return false
}

func rowVariantValue(cl *ast.CompositeLit) string {
	for _, el := range cl.Elts {
		kv, ok := el.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		id, ok := kv.Key.(*ast.Ident)
		if !ok || id.Name != "Variant" {
			continue
		}
		switch v := kv.Value.(type) {
		case *ast.Ident:
			return strings.TrimPrefix(strings.ToLower(v.Name), "row")
		case *ast.SelectorExpr:
			return strings.TrimPrefix(strings.ToLower(v.Sel.Name), "row")
		case *ast.BasicLit:
			return strings.Trim(v.Value, `"`)
		}
	}
	return ""
}

func fieldPresentNonEmpty(cl *ast.CompositeLit, name string) bool {
	for _, el := range cl.Elts {
		kv, ok := el.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		id, ok := kv.Key.(*ast.Ident)
		if !ok || id.Name != name {
			continue
		}
		switch v := kv.Value.(type) {
		case *ast.Ident:
			return v.Name != "nil"
		case *ast.CompositeLit:
			return len(v.Elts) > 0
		case *ast.CallExpr:
			return true
		default:
			return true
		}
	}
	return false
}
