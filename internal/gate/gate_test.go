package gate_test

import (
	"testing"

	"carrel.dev/typecase/internal/catalog"
	"carrel.dev/typecase/internal/gate"
)

func catalogEntries() []gate.CatalogEntry {
	var out []gate.CatalogEntry
	for _, e := range catalog.Entries() {
		out = append(out, gate.CatalogEntry{
			Widget: e.Widget,
			State:  e.State,
			HTML:   e.HTML,
		})
	}
	return out
}

// untilStep lists gates skipped on the plan 2 branch while the old widgets
// meet the mockup's CSS; each names the step that brings it back or removes it.
var untilStep = map[gate.Name]string{
	gate.Completeness: "library-from-mockup: шаг 4 — классы старых виджетов против CSS макета",
	gate.Geometry:     "library-from-mockup: шаг 4 — пункт «Геометрия» заменён «Эталоном», ворота уходят вместе с ним",
}

func TestL8_gates(t *testing.T) {
	root, err := gate.FindModuleRoot()
	if err != nil {
		t.Fatal(err)
	}
	p := gate.Paths(root)
	cat := catalogEntries()

	checks := []gate.Name{
		gate.Gofmt,
		gate.Vet,
		gate.Completeness,
		gate.Ownership,
		gate.NodeNames,
		gate.TypedInput,
		gate.Enumerations,
		gate.DomainNaming,
		gate.DocFormat,
		gate.Catalog,
		gate.SpecL4,
		gate.Verifiable,
		gate.NoInline,
		gate.A11y,
		gate.Geometry,
		gate.ScreensCSS,
		gate.ScreensCoverage,
		gate.ScreensCompose,
		gate.LayoutSlots,
		gate.MockupParity,
		gate.LayoutInventory,
		gate.LayoutOrphans,
		gate.CSSFromMockup,
	}
	for _, name := range checks {
		t.Run(string(name), func(t *testing.T) {
			if why, ok := untilStep[name]; ok {
				t.Skip(why)
			}
			if err := gate.Run(name, p, cat); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestL8_verifiable(t *testing.T) {
	root, err := gate.FindModuleRoot()
	if err != nil {
		t.Fatal(err)
	}
	p := gate.Paths(root)
	if err := gate.Run(gate.Verifiable, p, catalogEntries()); err != nil {
		t.Fatal(err)
	}
	if len(gate.AllGates) != 23 {
		t.Fatalf("expected 23 gates, got %d", len(gate.AllGates))
	}
}

func TestLayoutOrphans_report(t *testing.T) {
	report := gate.LayoutOrphansReport(catalogEntries())
	t.Log(report)
	if err := gate.CheckLayoutOrphans(catalogEntries()); err != nil {
		t.Fatal(err)
	}
}
