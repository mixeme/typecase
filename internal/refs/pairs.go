package refs

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"carrel.dev/typecase/internal/fragment"
)

// A container declares what it takes by the types of its fields (L-7):
//
//   - a field whose type is a widget is a fixed child (Head in a sheet);
//   - a field whose type is a slot — an interface with an unexported method —
//     takes every widget that implements the slot;
//   - slices, arrays and pointers of either count as their element;
//   - a plain struct field (a group of fields) is looked into.
//
// An interface field without an unexported method would take any widget;
// that is the "[]Node" slot the spec removed, and CheckPairs refuses it.

// Allowed returns, for container type t, every widget type it can hold and
// the field path that allows it. A slot implemented by a part — a piece of
// the container that is not a widget of its own, such as a field's hint — is
// looked into: what the part holds, the container holds.
func Allowed(t reflect.Type, widgets []reflect.Type, parts ...reflect.Type) (map[reflect.Type]string, error) {
	out := map[reflect.Type]string{}
	isWidget := map[reflect.Type]bool{}
	for _, w := range widgets {
		isWidget[w] = true
	}
	var errs []string
	seenPart := map[reflect.Type]bool{}
	var walk func(ft reflect.Type, path string, depth int)
	walk = func(ft reflect.Type, path string, depth int) {
		for ft.Kind() == reflect.Slice || ft.Kind() == reflect.Array || ft.Kind() == reflect.Pointer {
			ft = ft.Elem()
		}
		switch {
		case ft.Kind() == reflect.Interface:
			if !isSlot(ft) {
				errs = append(errs, fmt.Sprintf("%s.%s: %s takes any widget; a slot needs an unexported method", t.Name(), path, ft))
				return
			}
			for _, w := range widgets {
				if w.Implements(ft) {
					if _, ok := out[w]; !ok {
						out[w] = path
					}
				}
			}
			for _, p := range parts {
				if p.Implements(ft) && depth < 4 && !seenPart[p] {
					seenPart[p] = true
					walk(p, path+"("+p.Name()+")", depth+1)
				}
			}
		case isWidget[ft] && depth > 0:
			if _, ok := out[ft]; !ok {
				out[ft] = path
			}
		case ft.Kind() == reflect.Struct && depth < 5:
			for i := 0; i < ft.NumField(); i++ {
				f := ft.Field(i)
				if !f.IsExported() {
					continue
				}
				p := f.Name
				if path != "" {
					p = path + "." + f.Name
				}
				walk(f.Type, p, depth+1)
			}
		}
	}
	walk(t, "", 0)
	if len(errs) > 0 {
		return nil, fmt.Errorf("%s", strings.Join(errs, "; "))
	}
	return out, nil
}

func isSlot(it reflect.Type) bool {
	for i := 0; i < it.NumMethod(); i++ {
		if !it.Method(i).IsExported() {
			return true
		}
	}
	return false
}

// CheckPairs compares the mockup's parent → child block pairs with what the
// widget types allow, both ways: a pair the mockup draws must compile, and a
// nesting the types allow must be drawn somewhere. Classes without a widget
// type yet are left out; step 7 makes sure none stays so.
func CheckPairs(pairs []fragment.PairUse, entries []Entry, widgets []Renderer, parts ...any) error {
	classTypes := map[string][]reflect.Type{}
	typeClass := map[reflect.Type]string{}
	var errs []string
	for _, e := range entries {
		t := reflect.TypeOf(e.Fixture)
		if prev, ok := typeClass[t]; ok {
			if prev != e.Ref.Class {
				errs = append(errs, fmt.Sprintf("%s prints two blocks, .%s and .%s", t, prev, e.Ref.Class))
			}
			continue
		}
		typeClass[t] = e.Ref.Class
		classTypes[e.Ref.Class] = append(classTypes[e.Ref.Class], t)
	}
	var types, partTypes []reflect.Type
	for _, w := range widgets {
		types = append(types, reflect.TypeOf(w))
	}
	for _, p := range parts {
		partTypes = append(partTypes, reflect.TypeOf(p))
	}
	drawn := map[fragment.Pair]fragment.PairUse{}
	for _, p := range pairs {
		drawn[p.Pair] = p
	}
	allowed := map[reflect.Type]map[reflect.Type]string{}
	for t := range typeClass {
		a, err := Allowed(t, types, partTypes...)
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		allowed[t] = a
	}
	// A drawn pair compiles when some type of the parent block takes some
	// type of the child block: one block can be several types (tc-in is an
	// input, a pick, a text area), and the frame does not say which.
	for _, p := range pairs {
		pts, ok1 := classTypes[p.Parent]
		cts, ok2 := classTypes[p.Child]
		if !ok1 || !ok2 {
			continue
		}
		ok := false
		for _, pt := range pts {
			for _, ct := range cts {
				if _, yes := allowed[pt][ct]; yes {
					ok = true
				}
			}
		}
		if !ok {
			errs = append(errs, fmt.Sprintf("the mockup puts .%s in .%s (%s, ×%d) but no type of .%s has a field that takes a type of .%s",
				p.Child, p.Parent, p.First, p.Count, p.Parent, p.Child))
		}
	}
	for pt, children := range allowed {
		for ct, path := range children {
			cc, ok := typeClass[ct]
			if !ok {
				continue
			}
			if _, ok := drawn[fragment.Pair{Parent: typeClass[pt], Child: cc}]; !ok {
				errs = append(errs, fmt.Sprintf("%s.%s takes %s, but no frame puts .%s in .%s",
					pt.Name(), path, ct.Name(), cc, typeClass[pt]))
			}
		}
	}
	if len(errs) > 0 {
		sort.Strings(errs)
		return fmt.Errorf("container-pairs:\n  %s", strings.Join(errs, "\n  "))
	}
	return nil
}
