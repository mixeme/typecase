package refs_test

import (
	"strings"
	"testing"

	"carrel.dev/typecase/internal/fragment"
	"carrel.dev/typecase/internal/refs"
)

// TestContainerPairs: every pair the mockup draws is allowed by the types,
// and every nesting the types allow is drawn (step 3).
func TestContainerPairs(t *testing.T) {
	m := mockup(t)
	if err := refs.CheckPairs(m.Pairs(), refs.Entries, refs.Widgets); err != nil {
		t.Fatal(err)
	}
}

// Toy widgets for the bite: a sheet with a head and a foot slot.
type footPart interface{ footPart() }

type toySheet struct {
	Head toyHead
	Foot []footPart
}
type toyHead struct{ Title string }
type toyButton struct{ Label string }
type toyKbd struct{ Key string }
type toyAny interface{ Render() string }
type toyLoose struct{ Body []toyAny }

func (toySheet) Render() string  { return "" }
func (toyHead) Render() string   { return "" }
func (toyButton) Render() string { return "" }
func (toyKbd) Render() string    { return "" }
func (toyLoose) Render() string  { return "" }
func (toyButton) footPart()      {}

func pairs(ps ...[2]string) []fragment.PairUse {
	var out []fragment.PairUse
	for _, p := range ps {
		out = append(out, fragment.PairUse{Pair: fragment.Pair{Parent: p[0], Child: p[1]}, Count: 1})
	}
	return out
}

var toyEntries = []refs.Entry{
	{Fixture: toySheet{}, Ref: fragment.Ref{Class: "tc-sheet"}},
	{Fixture: toyHead{}, Ref: fragment.Ref{Class: "tc-head"}},
	{Fixture: toyButton{}, Ref: fragment.Ref{Class: "tc-btn"}},
	{Fixture: toyKbd{}, Ref: fragment.Ref{Class: "tc-kbd"}},
}
var toyWidgets = []refs.Renderer{toySheet{}, toyHead{}, toyButton{}, toyKbd{}}

func TestBite_containerPairs(t *testing.T) {
	drawn := pairs([2]string{"tc-sheet", "tc-head"}, [2]string{"tc-sheet", "tc-btn"})
	if err := refs.CheckPairs(drawn, toyEntries, toyWidgets); err != nil {
		t.Fatalf("matching toys are red: %v", err)
	}
	// The mockup draws a kbd in the sheet; no field takes it.
	more := append(drawn, pairs([2]string{"tc-sheet", "tc-kbd"})...)
	if err := refs.CheckPairs(more, toyEntries, toyWidgets); err == nil || !strings.Contains(err.Error(), "no field that takes toyKbd") {
		t.Fatalf("an undrawable pair passed: %v", err)
	}
	// The types allow a button in the sheet; no frame draws it.
	fewer := pairs([2]string{"tc-sheet", "tc-head"})
	if err := refs.CheckPairs(fewer, toyEntries, toyWidgets); err == nil || !strings.Contains(err.Error(), "toySheet.Foot takes toyButton") {
		t.Fatalf("an extra permission passed: %v", err)
	}
	// A slot that takes any widget is refused.
	loose := append(toyEntries, refs.Entry{Fixture: toyLoose{}, Ref: fragment.Ref{Class: "tc-card"}})
	if err := refs.CheckPairs(drawn, loose, append(toyWidgets, toyLoose{})); err == nil || !strings.Contains(err.Error(), "takes any widget") {
		t.Fatalf("an any-widget slot passed: %v", err)
	}
	// One class printed by two types is refused.
	twice := append(toyEntries, refs.Entry{Fixture: toyKbd{}, Ref: fragment.Ref{Class: "tc-btn"}})
	if err := refs.CheckPairs(drawn, twice, toyWidgets); err == nil {
		t.Fatal("one class, two types passed")
	}
}
