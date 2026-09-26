package gate

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"sort"
	"strings"

	"carrel.dev/typecase/internal/fragment"
	"carrel.dev/typecase/internal/refs"
)

// CheckWidgetFragments fails when a reference record does not print the
// fragment of its frame (plan 2, step 2), and when an exported widget has no
// record (step 7): every widget prints its frame.
func CheckWidgetFragments(p ModulePaths) error {
	m, err := fragment.Load(p.Root)
	if err != nil {
		return fmt.Errorf("widget-fragments: %w", err)
	}
	var bad []string
	for _, e := range refs.Entries {
		if err := m.Compare(e.Ref, e.Fixture.Render()); err != nil {
			bad = append(bad, fmt.Sprintf("%T: %v", e.Fixture, err))
		}
	}
	if len(bad) > 0 {
		return fmt.Errorf("widget-fragments:\n  %s", strings.Join(bad, "\n  "))
	}
	var listed []string
	for _, w := range refs.Widgets {
		listed = append(listed, typeName(w))
	}
	recorded := map[string]int{}
	for _, e := range refs.Entries {
		recorded[typeName(e.Fixture)]++
	}
	return CheckEveryWidgetHasRecord(p.WidgetDir, listed, recorded)
}

// CheckContainerPairs fails when what a container takes and what the mockup
// puts in it differ, either way (plan 2, step 3).
func CheckContainerPairs(p ModulePaths) error {
	m, err := fragment.Load(p.Root)
	if err != nil {
		return fmt.Errorf("container-pairs: %w", err)
	}
	return refs.CheckPairs(m.Pairs(), refs.Entries, refs.Widgets, refs.Parts...)
}

// renderTypes lists the exported types of the widget package that have a
// method Render() string: the widgets, read from the source.
func renderTypes(widgetDir string) ([]string, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, widgetDir, func(fi os.FileInfo) bool {
		return strings.HasSuffix(fi.Name(), ".go") && !strings.HasSuffix(fi.Name(), "_test.go")
	}, 0)
	if err != nil {
		return nil, err
	}
	seen := map[string]bool{}
	for _, pkg := range pkgs {
		for _, f := range pkg.Files {
			for _, d := range f.Decls {
				fn, ok := d.(*ast.FuncDecl)
				if !ok || fn.Recv == nil || fn.Name.Name != "Render" || len(fn.Recv.List) != 1 {
					continue
				}
				if fn.Type.Params.NumFields() != 0 || fn.Type.Results.NumFields() != 1 {
					continue
				}
				if id, ok := fn.Type.Results.List[0].Type.(*ast.Ident); !ok || id.Name != "string" {
					continue
				}
				recv := fn.Recv.List[0].Type
				if star, ok := recv.(*ast.StarExpr); ok {
					recv = star.X
				}
				if id, ok := recv.(*ast.Ident); ok && ast.IsExported(id.Name) {
					seen[id.Name] = true
				}
			}
		}
	}
	var out []string
	for n := range seen {
		out = append(out, n)
	}
	sort.Strings(out)
	return out, nil
}

// CheckEveryWidgetHasRecord fails when an exported widget type has no
// reference record, or when refs.Widgets and the source disagree on what the
// widgets are (plan 2, step 7): a widget without a frame does not stay.
func CheckEveryWidgetHasRecord(widgetDir string, listed []string, recorded map[string]int) error {
	src, err := renderTypes(widgetDir)
	if err != nil {
		return err
	}
	inList := map[string]bool{}
	for _, n := range listed {
		inList[n] = true
	}
	var bad []string
	inSrc := map[string]bool{}
	for _, n := range src {
		inSrc[n] = true
		if !inList[n] {
			bad = append(bad, n+": not in refs.Widgets")
		}
		if recorded[n] == 0 {
			bad = append(bad, n+": no reference record")
		}
	}
	for _, n := range listed {
		if !inSrc[n] {
			bad = append(bad, n+": in refs.Widgets, not a widget of the source")
		}
	}
	if len(bad) > 0 {
		return fmt.Errorf("widget-fragments: %s", strings.Join(bad, "; "))
	}
	return nil
}

func typeName(v any) string {
	s := fmt.Sprintf("%T", v)
	return s[strings.LastIndex(s, ".")+1:]
}
