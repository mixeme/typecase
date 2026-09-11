package gate

import (
	"os"
	"strings"
	"testing"
)

func TestCheckMockupLint_live(t *testing.T) {
	root, err := FindModuleRoot()
	if err != nil {
		t.Fatal(err)
	}
	if err := CheckMockupLintFile(root); err != nil {
		t.Fatal(err)
	}
}

func TestCheckMockupLint_frameIDs(t *testing.T) {
	ok := `<h2 id="s-notes"></h2><div class="frame" data-frame="s-notes-1"></div>` +
		`<div class="frame" data-frame="s-notes-2"></div>`
	if err := CheckMockupLint(ok); err != nil {
		t.Fatal(err)
	}
	if err := CheckMockupLint(`<div class="frame"></div>`); err == nil {
		t.Fatal("expected missing data-frame to fail")
	}
	dup := `<div class="frame" data-frame="s-notes-1"></div>` +
		`<div class="frame" data-frame="s-notes-1"></div>`
	if err := CheckMockupLint(dup); err == nil {
		t.Fatal("expected duplicate data-frame to fail")
	}
	if err := CheckMockupLint(`<div class="frames"></div>`); err == nil {
		t.Fatal("expected no .frame to fail")
	}
	if got := ListMockupFrames(ok); len(got) != 2 || got[0] != "s-notes-1" || got[1] != "s-notes-2" {
		t.Fatalf("ListMockupFrames: %v", got)
	}
}

func TestParseMockupParts_live(t *testing.T) {
	root, err := FindModuleRoot()
	if err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(MockupFile(root))
	if err != nil {
		t.Fatal(err)
	}
	p, err := ParseMockupParts(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if len(p.Frames) != 163 {
		t.Fatalf("frames: got %d, want 163", len(p.Frames))
	}
	if p.Frames[0].ID != "shell-1" || !strings.Contains(p.Frames[0].HTML, "data-frame=\"shell-1\"") {
		t.Fatalf("first frame: %+v", p.Frames[0].ID)
	}
}
