package refs_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"carrel.dev/typecase/internal/fragment"
	"carrel.dev/typecase/internal/refs"
)

func mockup(t *testing.T) *fragment.Mockup {
	t.Helper()
	wd, _ := os.Getwd()
	root, err := fragment.FindRoot(wd)
	if err != nil {
		t.Fatal(err)
	}
	m, err := fragment.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	return m
}

// TestWidgetFragments: every record prints the fragment of its frame.
func TestWidgetFragments(t *testing.T) {
	m := mockup(t)
	for _, e := range refs.Entries {
		name := fmt.Sprintf("%T/%s", e.Fixture, e.Ref)
		t.Run(name, func(t *testing.T) {
			if err := m.Compare(e.Ref, e.Fixture.Render()); err != nil {
				t.Fatal(err)
			}
		})
	}
}

// raw prints fixed markup: the bite needs a widget whose output it controls.
type raw string

func (r raw) Render() string { return string(r) }

// Bite: a widget printing the frame's fragment is green; the same markup with
// one class spoiled is red and the message names the path to the node.
func TestBite_widgetFragments(t *testing.T) {
	m := mockup(t)
	ref := fragment.Ref{Frame: "components-19", Class: "tc-menu", N: 1}
	good := raw(`<div class="tc-menu is-pop is-right tc-menu--w212">
  <button class="tc-menu__item" type="button" role="menuitem"><svg><use href="#i-pencil"/></svg>Rename</button>
  <button class="tc-menu__item" type="button" role="menuitem"><svg><use href="#i-down"/></svg>Export <span class="tc-mono tc-mono--t11">.vcf</span></button>
</div>`)
	if err := m.Compare(ref, good.Render()); err != nil {
		t.Fatalf("the frame's own fragment is red: %v", err)
	}
	spoiled := raw(strings.Replace(string(good), `<span class="tc-mono tc-mono--t11">`, `<span class="tc-mono">`, 1))
	err := m.Compare(ref, spoiled.Render())
	if err == nil {
		t.Fatal("a spoiled class passed")
	}
	if !strings.Contains(err.Error(), "button.tc-menu__item[2] > span.tc-mono.tc-mono--t11") {
		t.Fatalf("the message does not name the path: %v", err)
	}
	wrongText := raw(strings.Replace(string(good), "Rename", "Rename it", 1))
	if err := m.Compare(ref, wrongText.Render()); err == nil || !strings.Contains(err.Error(), `"Rename"`) {
		t.Fatalf("changed text not caught: %v", err)
	}
	wrongIcon := raw(strings.Replace(string(good), "#i-pencil", "#i-trash", 1))
	if err := m.Compare(ref, wrongIcon.Render()); err == nil {
		t.Fatal("changed icon not caught")
	}
	aria := raw(strings.Replace(string(good), `role="menuitem"`, `role="menuitem" aria-label="Rename the card" title="x" id="m1"`, 1))
	if err := m.Compare(ref, aria.Render()); err != nil {
		t.Fatalf("accessibility attributes must not be compared: %v", err)
	}
}

// Every listed widget really prints something with a tc- class.
func TestWidgetsList(t *testing.T) {
	if len(refs.Widgets) == 0 {
		t.Fatal("refs.Widgets is empty")
	}
	seen := map[string]bool{}
	for _, w := range refs.Widgets {
		name := fmt.Sprintf("%T", w)
		if seen[name] {
			t.Fatalf("%s listed twice", name)
		}
		seen[name] = true
	}
}
