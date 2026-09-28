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
	ok := `<style data-tc>:root{--tc-ink:#000}.tc-tag,.tc-btn{color:red}</style>` +
		`<style>body{margin:0}.frame{border:0}.wrap{padding:0}.mk-win{border:0}</style>` +
		`<div class="frame" data-frame="s-notes-1"></div>`
	if err := CheckMockupLint(ok, MockupStyleBlocks, MockupFrameIDs); err != nil {
		t.Fatal(err)
	}
	tcInPage := `<style data-tc>.tc-btn{color:red}</style>` +
		`<style>.tc-btn{height:26px}</style>` +
		`<div class="frame" data-frame="s-notes-1"></div>`
	if err := CheckMockupLint(tcInPage); err == nil {
		t.Fatal("expected product class .tc-btn in page style to fail")
	}
	mkInProduct := `<style data-tc>.mk-win{border:0}</style>` +
		`<style>body{margin:0}</style>` +
		`<div class="frame" data-frame="s-notes-1"></div>`
	if err := CheckMockupLint(mkInProduct); err == nil {
		t.Fatal("expected page class .mk-win in data-tc to fail")
	}
	pageInProduct := `<style data-tc>.frame{border:0}.tc-tag{color:red}</style>` +
		`<style>body{margin:0}</style>` +
		`<div class="frame" data-frame="s-notes-1"></div>`
	if err := CheckMockupLint(pageInProduct); err == nil {
		t.Fatal("expected page selector in data-tc to fail")
	}
	productInPage := `<style data-tc>.tc-tag{color:red}</style>` +
		`<style>.tc-tag{height:26px}</style>` +
		`<div class="frame" data-frame="s-notes-1"></div>`
	if err := CheckMockupLint(productInPage); err == nil {
		t.Fatal("expected product class in page style to fail")
	}
	oneStyle := `<style>.tc-tag{color:red}</style>` +
		`<div class="frame" data-frame="s-notes-1"></div>`
	if err := CheckMockupLint(oneStyle); err == nil {
		t.Fatal("expected a single <style> to fail")
	}
	dataVals := `<style data-tc>[data-tc-colour]{background:var(--tc-colour)}</style>` +
		`<style>body{margin:0}[data-tc-colour="#4A6B52"]{--tc-colour:#4A6B52}</style>` +
		`<div class="frame" data-frame="s-notes-1"></div>`
	if err := CheckMockupLint(dataVals, MockupStyleBlocks, MockupFrameIDs); err != nil {
		t.Fatal(err)
	}
	ladder := `<style data-tc>.tc-root{color:red}</style>` +
		`<style>body{margin:0}.ladder{display:flex}body:has(#w820:checked) .tc-root{width:820px}</style>` +
		`<div class="frame" data-frame="s-notes-1"></div>`
	if err := CheckMockupLint(ladder, MockupStyleBlocks, MockupFrameIDs); err != nil {
		t.Fatal(err)
	}
}

func TestCheckMockupLint_frameClasses(t *testing.T) {
	ok := `<style data-tc>.tc-btn{color:red}</style><style>body{}</style>` +
		`<div class="frame" data-frame="s-notes-1"><div class="tc-root"><span class="tc-btn is-primary">Go</span></div></div>`
	if err := CheckMockupLint(ok, MockupFrameClasses); err != nil {
		t.Fatal(err)
	}
	desktop := `<div class="frame" data-frame="m-desktop-1"><div class="mk-win"><div class="tc-root"></div></div></div>`
	if err := CheckMockupLint(desktop, MockupFrameClasses); err != nil {
		t.Fatal(err)
	}
	if err := CheckMockupLint(
		`<div class="frame" data-frame="s-notes-1"><div class="mock tc-root"></div></div>`,
		MockupFrameClasses,
	); err == nil {
		t.Fatal("expected .mock inside a frame to fail")
	}
	if err := CheckMockupLint(
		`<div class="frame" data-frame="s-notes-1"><span class="m-btn">Go</span></div>`,
		MockupFrameClasses,
	); err == nil {
		t.Fatal("expected leftover m-* class to fail")
	}
	if err := CheckMockupLint(
		`<div class="frame" data-frame="s-notes-1"><div class="mk-win"></div></div>`,
		MockupFrameClasses,
	); err == nil {
		t.Fatal("expected mk-* outside m-desktop to fail")
	}
	labelOK := `<div class="frame" data-frame="s-notes-1"><div class="frame-label"><span class="frag">компонент</span></div><div class="tc-root"></div></div>`
	if err := CheckMockupLint(labelOK, MockupFrameClasses); err != nil {
		t.Fatal(err)
	}
}

func TestCheckMockupLint_frameElements(t *testing.T) {
	ok := `<style data-tc>.tc-btn{color:red}</style><style>body{}</style>` +
		`<div class="frame" data-frame="s-notes-1"><div class="tc-root"><button class="tc-btn">Go</button><kbd class="tc-kbd">Esc</kbd></div></div>`
	if err := CheckMockupLint(ok, MockupFrameElements); err != nil {
		t.Fatal(err)
	}
	if err := CheckMockupLint(
		`<div class="frame" data-frame="s-notes-1"><span class="tc-btn">Go</span></div>`,
		MockupFrameElements,
	); err == nil {
		t.Fatal("expected span.tc-btn to fail")
	}
	if err := CheckMockupLint(
		`<div class="frame" data-frame="s-notes-1"><span class="tc-kbd">Esc</span></div>`,
		MockupFrameElements,
	); err == nil {
		t.Fatal("expected span.tc-kbd to fail")
	}
	labelOK := `<div class="frame" data-frame="s-notes-1"><div class="frame-label"><span class="tc-btn">not a control</span></div></div>`
	if err := CheckMockupLint(labelOK, MockupFrameElements); err != nil {
		t.Fatal(err)
	}
	themeOK := `<style data-tc>.tc-btn{color:red}</style><style>body{}</style>` +
		`<div class="frame" data-frame="account-5"><div class="tc-root">` +
		`<div class="tc-menu__item tc-menu__item--r6"><span>Theme</span>` +
		`<span class="tc-seg"><button class="tc-seg__opt" type="button">Auto</button></span></div>` +
		`</div></div>`
	if err := CheckMockupLint(themeOK, MockupFrameElements); err != nil {
		t.Fatal(err)
	}
	fieldOK := `<div class="frame" data-frame="s-dup-3"><div class="tc-root">` +
		`<div class="tc-menu__item">Same UID<input class="tc-in" type="text" value="80"></div>` +
		`</div></div>`
	if err := CheckMockupLint(fieldOK, MockupFrameElements); err != nil {
		t.Fatalf("a menu row holding a field is a div: %v", err)
	}
	bare := `<div class="frame" data-frame="s-dup-3"><div class="tc-root"><div class="tc-menu__item">Save</div></div></div>`
	if err := CheckMockupLint(bare, MockupFrameElements); err == nil {
		t.Fatal("a menu row with no control of its own must be a button")
	}
}

func TestCheckMockupLint_frameNoStyle(t *testing.T) {
	ok := `<style data-tc>.tc-btn{color:red}</style><style>body{}</style>` +
		`<div class="frame w392" data-frame="s-notes-1"><div class="frame-label"><span style="color:red">x</span></div><div class="tc-root"><span class="tc-mark"></span></div></div>`
	if err := CheckMockupLint(ok, MockupFrameNoStyle); err != nil {
		t.Fatal(err)
	}
	if err := CheckMockupLint(
		`<div class="frame" data-frame="s-notes-1"><span class="tc-mark" style="background:#4A6B52"></span></div>`,
		MockupFrameNoStyle,
	); err == nil {
		t.Fatal("expected style= inside a frame to fail")
	}
	if err := CheckMockupLint(
		`<div class="frame" data-frame="s-notes-1" style="max-width:900px"><div class="tc-root"></div></div>`,
		MockupFrameNoStyle,
	); err == nil {
		t.Fatal("expected style= on .frame itself to fail")
	}
}

func TestCheckMockupLint_lock(t *testing.T) {
	ok := `<style data-tc>:root{--tc-ink:#000}.tc-root{color:var(--tc-ink);container:tcbox / inline-size}.tc-btn{color:red}</style>` +
		`<style>body{margin:0}[data-tc-colour="#4A6B52"]{--tc-colour:#4A6B52}url("fonts/x.ttf")</style>` +
		`<div class="frame" data-frame="s-notes-1"><div class="tc-root"><button class="tc-btn" type="button">Go</button></div></div>`
	if err := CheckMockupLint(ok, MockupFrameNamedClass, MockupDataAttrs, MockupSelfContained, MockupClassesDefined, MockupDeadRules, MockupNamedContainer, MockupTokenVars); err != nil {
		t.Fatal(err)
	}
	bare := `<div class="frame" data-frame="s-notes-1"><div class="tc-root"><div></div></div></div>`
	if err := CheckMockupLint(bare, MockupFrameNamedClass); err == nil {
		t.Fatal("expected classless div to fail")
	}
	noRule := `<style data-tc>.tc-root{color:red}</style><style>body{}</style>` +
		`<div class="frame" data-frame="s-notes-1"><span class="tc-mark" data-tc-colour="#4A6B52"></span></div>`
	if err := CheckMockupLint(noRule, MockupDataAttrs); err == nil {
		t.Fatal("expected missing colour rule to fail")
	}
	script := `<style data-tc>.tc-root{color:red}</style><style>body{}</style>` +
		`<div class="frame" data-frame="s-notes-1"></div><script>alert(1)</script>`
	if err := CheckMockupLint(script, MockupSelfContained); err == nil {
		t.Fatal("expected script to fail")
	}
	undef := `<style data-tc>.tc-root{color:red}</style><style>body{}</style>` +
		`<div class="frame" data-frame="s-notes-1"><span class="tc-mark"></span></div>`
	if err := CheckMockupLint(undef, MockupClassesDefined); err == nil {
		t.Fatal("expected undefined tc-mark to fail")
	}
	dead := `<style data-tc>.tc-root{color:red}.tc-gone{display:none}</style><style>body{}</style>` +
		`<div class="frame" data-frame="s-notes-1"><div class="tc-root"></div></div>`
	if err := CheckMockupLint(dead, MockupDeadRules); err == nil {
		t.Fatal("expected dead .tc-gone to fail")
	}
	unnamed := `<style data-tc>.tc-root{container-type:inline-size}</style><style>body{}</style>` +
		`<div class="frame" data-frame="s-notes-1"></div>`
	if err := CheckMockupLint(unnamed, MockupNamedContainer); err == nil {
		t.Fatal("expected unnamed container to fail")
	}
	badVar := `<style data-tc>.tc-root{color:var(--ink)}</style><style>body{}</style>` +
		`<div class="frame" data-frame="s-notes-1"></div>`
	if err := CheckMockupLint(badVar, MockupTokenVars); err == nil {
		t.Fatal("expected unknown var to fail")
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
	if len(p.Frames) != 324 {
		t.Fatalf("frames: got %d, want 324", len(p.Frames))
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

func TestCheckMockupLint_emptyRules(t *testing.T) {
	frame := `<div class="frame" data-frame="s-notes-1"><div class="tc-root">` +
		`<span class="tc-line__item">x</span></div></div>`
	ok := `<style data-tc>.tc-root{color:red}` +
		`/* names-without-rule: tc-line__item */</style>` +
		`<style>body{margin:0}</style>` + frame
	if err := CheckMockupLint(ok, MockupEmptyRules, MockupClassesDefined); err != nil {
		t.Fatal(err)
	}
	empty := `<style data-tc>.tc-root{color:red}.tc-line__item{}</style>` +
		`<style>body{margin:0}</style>` + frame
	if err := CheckMockupLint(empty, MockupEmptyRules); err == nil {
		t.Fatal("expected a rule with no declarations to fail")
	}
	stale := `<style data-tc>.tc-root{color:red}` +
		`/* names-without-rule: tc-line__item tc-gone */</style>` +
		`<style>body{margin:0}</style>` + frame
	if err := CheckMockupLint(stale, MockupEmptyRules); err == nil {
		t.Fatal("expected a listed name that stands on no node to fail")
	}
	undeclared := `<style data-tc>.tc-root{color:red}</style>` +
		`<style>body{margin:0}</style>` + frame
	if err := CheckMockupLint(undeclared, MockupClassesDefined); err == nil {
		t.Fatal("expected an unlisted name with no rule to fail")
	}
}

func TestMockupNumberedClasses(t *testing.T) {
	html := `<style data-tc>.tc-root{color:red}` +
		`.tc-div--r7.tc-div--r7{margin-top:12px}` +
		`@container tcbox (max-width:640px){.tc-div--r7{margin-top:0}}` +
		`.tc-card--r1.tc-card--r1{margin:0}</style>` +
		`<style>body{margin:0}</style>` +
		`<div class="frame" data-frame="s-notes-1"><div class="tc-root">` +
		`<div class="tc-card tc-card--r1"><div class="tc-div--r7">x</div></div>` +
		`</div></div>`
	list, err := MockupNumberedClasses(html)
	if err != nil {
		t.Fatal(err)
	}
	if len(list) != 2 {
		t.Fatalf("want 2 numbered classes, got %d: %+v", len(list), list)
	}
	byName := map[string]NumberedClass{}
	for _, c := range list {
		byName[c.Name] = c
	}
	div, ok := byName["tc-div--r7"]
	if !ok {
		t.Fatal("tc-div--r7 missing")
	}
	if !div.TagBlock() {
		t.Error("tc-div--r7 is a block by tag name")
	}
	if len(div.Rules) != 2 {
		t.Errorf("tc-div--r7 rules: %v", div.Rules)
	}
	if len(div.Uses) != 1 || div.Uses[0].In != "tc-card" || div.Uses[0].Frame != "s-notes-1" {
		t.Errorf("tc-div--r7 uses: %+v", div.Uses)
	}
	card := byName["tc-card--r1"]
	if card.TagBlock() {
		t.Error("tc-card--r1 is a real block")
	}
	if len(card.Uses) != 1 || card.Uses[0].Own != "tc-card" {
		t.Errorf("tc-card--r1 uses: %+v", card.Uses)
	}
	if !IsNumberedClass("tc-div--r7") || IsNumberedClass("tc-line--gap-14") || IsNumberedClass("tc-table__col--w110") {
		t.Error("IsNumberedClass must catch --rN only")
	}
}

func TestCheckMockupLint_noNumberedClasses(t *testing.T) {
	head := `<style data-tc>.tc-root{color:red}`
	tail := `</style><style>body{margin:0}</style>`
	frame := func(cls string) string {
		return `<div class="frame" data-frame="s-notes-1"><div class="tc-root">` +
			`<div class="` + cls + `">x</div></div></div>`
	}
	ok := head + `.tc-table__col--w110.tc-table__col--w110{width:110px}` + tail +
		frame("tc-table__col--w110")
	if err := CheckMockupLint(ok, MockupNoNumberedClasses); err != nil {
		t.Fatalf("a variant named by its value is a name: %v", err)
	}
	bad := head + `.tc-div--r37.tc-div--r37{margin-top:12px}` + tail + frame("tc-div--r37")
	err := CheckMockupLint(bad, MockupNoNumberedClasses)
	if err == nil {
		t.Fatal("expected --rN to fail")
	}
	if !strings.Contains(err.Error(), "tc-div--r37") {
		t.Fatalf("the message must name the class: %v", err)
	}
}

func TestMockupNesting(t *testing.T) {
	ok := `<style data-tc>.a{b:c}</style><figure><div class="frame" data-frame="s-x-1">` +
		`<div class="tc-root"><svg viewBox="0 0 16 16"><title>t</title><use href="#a"/></svg>` +
		`<input class="tc-in" value="1"><br></div></div></figure>`
	if err := CheckMockupLint(ok, MockupNesting); err != nil {
		t.Fatalf("well nested: %v", err)
	}
	// the shape of the rail-footer mistake: one </div> too many inside the
	// rail, which the frame parser does not see
	bad := `<figure><div class="frame" data-frame="s-x-1"><div class="tc-shell__rail">` +
		`<div class="tc-railfoot"><a>v</a></div><a>v</a></div></div>` +
		`<div class="tc-main">m</div></div></figure>`
	err := CheckMockupLint(bad, MockupNesting)
	if err == nil {
		t.Fatal("expected a stray </div> to fail")
	}
	if !strings.Contains(err.Error(), "</div> while <figure>") {
		t.Fatalf("the message must say what is open: %v", err)
	}
	if err := CheckMockupLint(`<div><span></div>`, MockupNesting); err == nil {
		t.Fatal("expected an unclosed <span> to fail")
	}
}
