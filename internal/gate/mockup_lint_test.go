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
	if err := CheckMockupLint(ok, MockupFrameIDs); err != nil {
		t.Fatal(err)
	}
	if err := CheckMockupLint(`<div class="frame"></div>`, MockupFrameIDs); err == nil {
		t.Fatal("expected missing data-frame to fail")
	}
	dup := `<div class="frame" data-frame="s-notes-1"></div>` +
		`<div class="frame" data-frame="s-notes-1"></div>`
	if err := CheckMockupLint(dup, MockupFrameIDs); err == nil {
		t.Fatal("expected duplicate data-frame to fail")
	}
	if err := CheckMockupLint(`<div class="frames"></div>`, MockupFrameIDs); err == nil {
		t.Fatal("expected no .frame to fail")
	}
	if got := ListMockupFrames(ok); len(got) != 2 || got[0] != "s-notes-1" || got[1] != "s-notes-2" {
		t.Fatalf("ListMockupFrames: %v", got)
	}
}

func TestCheckMockupLint_styleBlocks(t *testing.T) {
	ok := `<style data-tc>:root{--tc-ink:#000}.mock,.m-btn{color:red}</style>` +
		`<style>body{margin:0}.frame{border:0}.wrap{padding:0}</style>` +
		`<div class="frame" data-frame="s-notes-1"></div>`
	if err := CheckMockupLint(ok); err != nil {
		t.Fatal(err)
	}
	pageInProduct := `<style data-tc>.frame{border:0}.m-btn{color:red}</style>` +
		`<style>body{margin:0}</style>` +
		`<div class="frame" data-frame="s-notes-1"></div>`
	if err := CheckMockupLint(pageInProduct); err == nil {
		t.Fatal("expected page selector in data-tc to fail")
	}
	productInPage := `<style data-tc>.m-btn{color:red}</style>` +
		`<style>.m-btn{height:26px}</style>` +
		`<div class="frame" data-frame="s-notes-1"></div>`
	if err := CheckMockupLint(productInPage); err == nil {
		t.Fatal("expected product class in page style to fail")
	}
	oneStyle := `<style>.m-btn{color:red}</style>` +
		`<div class="frame" data-frame="s-notes-1"></div>`
	if err := CheckMockupLint(oneStyle); err == nil {
		t.Fatal("expected a single <style> to fail")
	}
}

func TestSplitTag_booleanAttr(t *testing.T) {
	name, attrs := splitTag(`style data-tc`)
	if name != "style" {
		t.Fatalf("name: %q", name)
	}
	if _, ok := attrs["data-tc"]; !ok {
		t.Fatalf("attrs: %#v", attrs)
	}
	name, attrs = splitTag(`path d="M18 42A30 30 0 0 1 78 42" fill="#4A6B52"`)
	if name != "path" || attrs["d"] == "" || attrs["fill"] != "#4A6B52" {
		t.Fatalf("path: %q %#v", name, attrs)
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
	if !strings.Contains(p.Styles, "data-tc") || !strings.Contains(p.Styles, "@font-face") {
		t.Fatal("ParseMockupParts.Styles must include both <style data-tc> and page @font-face")
	}
	if !strings.Contains(p.Styles, "@font-face") || !strings.Contains(p.Styles, "<style data-tc>") {
		t.Fatal("Styles must include page @font-face and <style data-tc>")
	}
}
