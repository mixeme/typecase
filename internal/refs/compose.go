package refs

import (
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strings"

	w "carrel.dev/typecase/widget"
)

var inlineType = reflect.TypeOf((*w.Inline)(nil)).Elem()

// Composition is one part put into one slot of one container: the free
// composition the types allow (L-7) and a screen of Carrel may build, though
// the frame that draws the pair draws it with other neighbours. The pairs gate
// says a slot takes only what the mockup draws in it somewhere; this asks
// whether every part it takes stands well in every place that slot is in.
type Composition struct {
	Owner string // the type whose field the slot is: Field, Rubric, Card…
	Slot  string // the slot's interface: FieldPart, RubricPart…
	Part  string // the type put in
	// Host is the record the part is put into: its frame and fragment are
	// where the composition is painted.
	Host Entry
	// Fixture is the host's fixture with the part in: appended to the slot's
	// list, or in place of the slot's one value.
	Fixture Renderer
	// Drawn says the host's slot held a part of this type already: no host
	// has the slot without it.
	Drawn bool
}

func (c Composition) String() string {
	return fmt.Sprintf("%s.%s+%s in %s", c.Owner, c.Slot, c.Part, c.Host.Ref)
}

// Compositions returns one composition for each container type, slot and
// part type the records reach: the part — the first value of its type the
// records hold in a slot of the same interface and its place does not hold
// already, else anywhere — goes into the first place in record order whose slot does
// not hold that type yet and holds something (a list of one more), else the
// first place at all. skip leaves hosts out (a frame that waits).
func Compositions(skip func(Entry) bool) []Composition {
	var types []reflect.Type
	for _, w := range Widgets {
		types = append(types, reflect.TypeOf(w))
	}
	for _, p := range Parts {
		types = append(types, reflect.TypeOf(p))
	}
	samples := map[reflect.Type]reflect.Value{}
	anywhere := map[reflect.Type][]reflect.Value{}
	for _, e := range Entries {
		gather(reflect.ValueOf(e.Fixture), samples)
		gatherSome(reflect.ValueOf(e.Fixture), anywhere)
	}
	// a part as it stands in a slot of this kind: a Text in a push is a word
	// or two, in a hint a sentence; the first few, so that the part put in
	// can differ from what its place holds (a second «New note» beside the
	// first read as a fault of the screen, the owner 10.11)
	inSlot := map[[2]reflect.Type][]reflect.Value{}
	type key struct {
		owner, slot reflect.Type
		part        reflect.Type
	}
	type place struct {
		entry int
		path  []int // to the owner struct: field and element indices
		field int
		score int // 2 holds others and not this type, 1 holds others, 0 empty
		drawn bool
	}
	best := map[key]place{}
	var order []key
	for ei, e := range Entries {
		if skip != nil && skip(e) {
			continue
		}
		var walk func(v reflect.Value, path []int)
		walk = func(v reflect.Value, path []int) {
			switch v.Kind() {
			case reflect.Pointer, reflect.Interface:
				if !v.IsNil() {
					walk(v.Elem(), path)
				}
			case reflect.Slice, reflect.Array:
				for i := 0; i < v.Len(); i++ {
					walk(v.Index(i), append(append([]int(nil), path...), i))
				}
			case reflect.Struct:
				t := v.Type()
				for i := 0; i < t.NumField(); i++ {
					f := t.Field(i)
					if !f.IsExported() {
						continue
					}
					fv := v.Field(i)
					slot, held := slotOf(f.Type, fv)
					if slot != nil {
						for _, h := range heldValues(fv) {
							k := [2]reflect.Type{slot, h.Type()}
							if len(inSlot[k]) < 8 && !holds(inSlot[k], h) {
								inSlot[k] = append(inSlot[k], h)
							}
						}
						for _, pt := range types {
							if !pt.Implements(slot) {
								continue
							}
							drawn := false
							for _, h := range held {
								if h == pt {
									drawn = true
								}
							}
							score := 0
							switch {
							case len(held) > 0 && !drawn:
								score = 2
							case len(held) > 0:
								score = 1
							}
							k := key{t, slot, pt}
							p, ok := best[k]
							if !ok {
								order = append(order, k)
							}
							if !ok || score > p.score {
								best[k] = place{ei, append([]int(nil), path...), i, score, drawn}
							}
						}
					}
					walk(fv, append(append([]int(nil), path...), i))
				}
			}
		}
		walk(reflect.ValueOf(e.Fixture), nil)
	}
	var out []Composition
	for _, k := range order {
		p := best[k]
		host := Entries[p.entry]
		shown := host.Fixture.Render()
		fixture := rebuild(reflect.ValueOf(host.Fixture), p.path, func(owner reflect.Value) {
			f := owner.Field(p.field)
			// the first sample its place does not hold already and the host
			// does not show — a part from a slot of this kind, else from any
			var sample reflect.Value
			here := heldValues(f)
			fresh := func(c reflect.Value) bool { return !holds(here, c) && !repeats(shown, c) }
			for _, cs := range [][]reflect.Value{inSlot[[2]reflect.Type{k.slot, k.part}], anywhere[k.part]} {
				for _, c := range cs {
					if !sample.IsValid() && fresh(c) {
						sample = c
					}
				}
			}
			if !sample.IsValid() {
				for _, cs := range [][]reflect.Value{inSlot[[2]reflect.Type{k.slot, k.part}], anywhere[k.part]} {
					for _, c := range cs {
						if !sample.IsValid() && !holds(here, c) {
							sample = c
						}
					}
				}
			}
			if !sample.IsValid() {
				if v, ok := samples[k.part]; ok {
					sample = v
				} else {
					sample = made(k.part)
				}
			}
			part := rebuild(sample, nil, nil)
			if f.Kind() == reflect.Slice {
				// a word put after words is parted from them by a space, as
				// a text that holds it would write it
				if f.Type().Elem() == inlineType && f.Len() > 0 {
					f.Set(reflect.Append(f, reflect.ValueOf(w.Text(" "))))
				}
				f.Set(reflect.Append(f, part))
			} else {
				f.Set(part)
			}
		})
		out = append(out, Composition{
			Owner:   k.owner.Name(),
			Slot:    k.slot.Name(),
			Part:    k.part.Name(),
			Host:    host,
			Fixture: fixture.Interface().(Renderer),
			Drawn:   p.drawn,
		})
	}
	sort.SliceStable(out, func(i, j int) bool {
		a, b := out[i], out[j]
		if a.Owner != b.Owner {
			return a.Owner < b.Owner
		}
		if a.Slot != b.Slot {
			return a.Slot < b.Slot
		}
		return a.Part < b.Part
	})
	return out
}

// repeats says the host already shows what the part would: its words, or —
// a part without words — its icons.
func repeats(shown string, v reflect.Value) bool {
	if w := words(v); w != "" {
		return strings.Contains(strings.Join(strings.Fields(markup.ReplaceAllString(shown, " ")), " "), w)
	}
	r, ok := v.Interface().(Renderer)
	if !ok {
		return false
	}
	var out string
	func() {
		defer func() { _ = recover() }()
		out = r.Render()
	}()
	icons := icon.FindAllString(out, -1)
	if len(icons) == 0 {
		return false
	}
	for _, i := range icons {
		if !strings.Contains(shown, i) {
			return false
		}
	}
	return true
}

var icon = regexp.MustCompile(`#i-[a-z0-9-]+`)

// gatherSome keeps the first few distinct values of every named type the
// value holds.
func gatherSome(v reflect.Value, into map[reflect.Type][]reflect.Value) {
	switch v.Kind() {
	case reflect.Pointer, reflect.Interface:
		if !v.IsNil() {
			gatherSome(v.Elem(), into)
		}
		return
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			gatherSome(v.Index(i), into)
		}
		return
	}
	if t := v.Type(); t.PkgPath() != "" && len(into[t]) < 32 && !holds(into[t], v) {
		into[t] = append(into[t], v)
	}
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			if v.Type().Field(i).IsExported() {
				gatherSome(v.Field(i), into)
			}
		}
	}
}

// holds says whether one of vs equals v, or says the same words: a drop
// whose button reads «New note» beside a «New note» button is the same
// action twice.
func holds(vs []reflect.Value, v reflect.Value) bool {
	w := words(v)
	for _, x := range vs {
		if x.Type() == v.Type() && reflect.DeepEqual(x.Interface(), v.Interface()) {
			return true
		}
		if w != "" && words(x) == w {
			return true
		}
	}
	return false
}

var markup = regexp.MustCompile(`<[^>]*>`)

// words is what a part prints to read, its first words: "" for one that
// prints none or does not print by itself.
func words(v reflect.Value) string {
	r, ok := v.Interface().(Renderer)
	if !ok {
		return ""
	}
	var out string
	func() {
		defer func() { _ = recover() }()
		out = r.Render()
	}()
	f := strings.Fields(markup.ReplaceAllString(out, " "))
	if len(f) > 3 {
		f = f[:3]
	}
	return strings.Join(f, " ")
}

// slotOf says whether a field is a slot — an interface with an unexported
// method, or a list of one — and which types it holds now.
func slotOf(ft reflect.Type, v reflect.Value) (reflect.Type, []reflect.Type) {
	switch {
	case ft.Kind() == reflect.Interface && isSlot(ft):
		if v.IsNil() {
			return ft, nil
		}
		return ft, []reflect.Type{valueType(v.Elem())}
	case ft.Kind() == reflect.Slice && ft.Elem().Kind() == reflect.Interface && isSlot(ft.Elem()):
		var held []reflect.Type
		for i := 0; i < v.Len(); i++ {
			if e := v.Index(i); !e.IsNil() {
				held = append(held, valueType(e.Elem()))
			}
		}
		return ft.Elem(), held
	}
	return nil, nil
}

// valueType is the type of a held value, a pointer read as what it points at
// (a record may hold a *Dim where the type list has Dim).
func valueType(v reflect.Value) reflect.Type {
	if v.Kind() == reflect.Pointer {
		return v.Type().Elem()
	}
	return v.Type()
}

// heldValues are the values a slot holds now, pointers read through.
func heldValues(v reflect.Value) []reflect.Value {
	var out []reflect.Value
	switch v.Kind() {
	case reflect.Interface:
		if !v.IsNil() {
			out = append(out, v.Elem())
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if e := v.Index(i); !e.IsNil() {
				out = append(out, e.Elem())
			}
		}
	}
	for i, h := range out {
		if h.Kind() == reflect.Pointer && !h.IsNil() {
			out[i] = h.Elem()
		}
	}
	return out
}

// gather keeps the first value of every named type the value holds.
func gather(v reflect.Value, into map[reflect.Type]reflect.Value) {
	switch v.Kind() {
	case reflect.Pointer, reflect.Interface:
		if !v.IsNil() {
			gather(v.Elem(), into)
		}
		return
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			gather(v.Index(i), into)
		}
		return
	}
	if v.Type().PkgPath() != "" {
		if _, ok := into[v.Type()]; !ok {
			into[v.Type()] = v
		}
	}
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			if v.Type().Field(i).IsExported() {
				gather(v.Field(i), into)
			}
		}
	}
}

// made is a value for a type no record holds: words for a text, else zero.
func made(t reflect.Type) reflect.Value {
	v := reflect.New(t).Elem()
	if t.Kind() == reflect.String {
		v.SetString(t.Name())
	}
	return v
}

// rebuild copies a value deep, so that a change does not reach the record it
// came from; on the way along path (field and element indices; pointers and
// interfaces take no step) it calls at on the struct the path ends at.
func rebuild(v reflect.Value, path []int, at func(reflect.Value)) reflect.Value {
	switch v.Kind() {
	case reflect.Pointer:
		if v.IsNil() {
			return v
		}
		n := reflect.New(v.Type().Elem())
		n.Elem().Set(rebuild(v.Elem(), path, at))
		return n
	case reflect.Interface:
		if v.IsNil() {
			return v
		}
		n := reflect.New(v.Type()).Elem()
		n.Set(rebuild(v.Elem(), path, at))
		return n
	case reflect.Slice:
		if v.IsNil() {
			return v
		}
		n := reflect.MakeSlice(v.Type(), v.Len(), v.Len())
		for i := 0; i < v.Len(); i++ {
			if at != nil && len(path) > 0 && path[0] == i {
				n.Index(i).Set(rebuild(v.Index(i), path[1:], at))
			} else {
				n.Index(i).Set(rebuild(v.Index(i), nil, nil))
			}
		}
		return n
	case reflect.Array:
		n := reflect.New(v.Type()).Elem()
		for i := 0; i < v.Len(); i++ {
			if at != nil && len(path) > 0 && path[0] == i {
				n.Index(i).Set(rebuild(v.Index(i), path[1:], at))
			} else {
				n.Index(i).Set(rebuild(v.Index(i), nil, nil))
			}
		}
		return n
	case reflect.Struct:
		n := reflect.New(v.Type()).Elem()
		n.Set(v)
		for i := 0; i < v.NumField(); i++ {
			if !v.Type().Field(i).IsExported() {
				continue
			}
			if at != nil && len(path) > 0 && path[0] == i {
				n.Field(i).Set(rebuild(v.Field(i), path[1:], at))
			} else {
				n.Field(i).Set(rebuild(v.Field(i), nil, nil))
			}
		}
		if at != nil && len(path) == 0 {
			at(n)
		}
		return n
	}
	return v
}
