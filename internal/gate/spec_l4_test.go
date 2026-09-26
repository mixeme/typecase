package gate

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadSpecL4Widgets_multiNameHeaders(t *testing.T) {
	spec := filepath.Join(t.TempDir(), "spec.md")
	text := "# Spec" + nl + nl + "## L-4. Состав" + nl + nl + "### L-4.1 Семья" + nl + nl +
		"#### `tc-stat` и `tc-stats` — показатель и ряд" + nl + nl + "#### `tc-card`, `tc-sec` — карточка" + nl + nl +
		"## L-5. Токены" + nl + nl + "#### `tc-outside` — не в L-4" + nl
	if err := os.WriteFile(spec, []byte(text), 0o644); err != nil {
		t.Fatal(err)
	}
	declared, err := LoadSpecL4Widgets(spec)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"tc-stat", "tc-stats", "tc-card", "tc-sec"} {
		if !declared[name] {
			t.Fatalf("expected %s in L-4 registry", name)
		}
	}
	if declared["tc-outside"] {
		t.Fatal("a header after L-4 counted")
	}
}

const nl = "\n"

func TestLoadSpecL4Widgets_requiresPath(t *testing.T) {
	if _, err := LoadSpecL4Widgets(""); err == nil {
		t.Fatal("expected error for empty spec path")
	}
}

func TestLoadSpecL4Widgets_missingFile(t *testing.T) {
	if _, err := LoadSpecL4Widgets(filepath.Join(t.TempDir(), "nope.md")); err == nil {
		t.Fatal("expected error for missing spec")
	}
}

func TestLoadSpecL4Widgets_noL4Section(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "spec.md")
	if err := os.WriteFile(path, []byte("## L-5. Only\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadSpecL4Widgets(path); err == nil {
		t.Fatal("expected error when L-4 section missing")
	}
}
