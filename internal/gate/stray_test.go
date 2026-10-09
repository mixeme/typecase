package gate

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// TestPlanStrays: a frame whose HTML changed is meant; a frame that is the
// same is measured only when a changed rule can reach it, and all of them are
// when the change cannot be told from a frame's HTML.
func TestPlanStrays(t *testing.T) {
	frames := []MockupFrame{
		{ID: "a", HTML: `<div class="frame"><button class="tc-btn">Move here</button></div>`},
		{ID: "b", HTML: `<div class="frame"><kbd class="tc-kbd">Esc</kbd></div>`},
		{ID: "c", HTML: `<div class="frame" data-tc-theme="dark"><span class="tc-num">3</span></div>`},
	}
	css := `<style>:root{--tc-ink:#111}
.tc-btn{padding:4px 8px}
.tc-kbd{font-size:11px}
@container tcroot (max-width:640px){.tc-num{display:none}}
[data-tc-theme="dark"] .tc-avatar{color:#fff}</style>`
	base := MockupParts{Styles: css, Sprite: "<svg></svg>", Frames: frames}
	with := func(edit func(p *MockupParts)) MockupParts {
		p := MockupParts{Styles: base.Styles, Sprite: base.Sprite, Frames: append([]MockupFrame(nil), frames...)}
		edit(&p)
		return p
	}
	swap := func(s, old, new string) string {
		out := ""
		for i := 0; i+len(old) <= len(s); i++ {
			if s[i:i+len(old)] == old {
				return s[:i] + new + s[i+len(old):]
			}
		}
		t.Fatalf("no %q in the css", old)
		return out
	}

	for _, c := range []struct {
		name  string
		after MockupParts
		meant []string
		check []string
	}{
		{"nothing changed", base, nil, nil},
		{"a frame's words", with(func(p *MockupParts) {
			p.Frames[0].HTML = `<div class="frame"><button class="tc-btn">Move</button></div>`
		}), []string{"a"}, nil},
		{"a rule of one class", with(func(p *MockupParts) {
			p.Styles = swap(p.Styles, "font-size:11px", "font-size:12px")
		}), nil, []string{"b"}},
		{"a rule under a condition", with(func(p *MockupParts) {
			p.Styles = swap(p.Styles, "max-width:640px", "max-width:700px")
		}), nil, []string{"c"}},
		{"a rule the frame has half of", with(func(p *MockupParts) {
			p.Styles = swap(p.Styles, "color:#fff", "color:#eee")
		}), nil, nil},
		{"a rule by an attribute", with(func(p *MockupParts) {
			p.Styles = swap(p.Styles, "</style>", `[data-tc-theme="dark"]{color:red}</style>`)
		}), nil, []string{"c"}},
		{"a new rule", with(func(p *MockupParts) {
			p.Styles = swap(p.Styles, "</style>", ".tc-btn:hover{color:red}</style>")
		}), nil, []string{"a"}},
		{"a token", with(func(p *MockupParts) {
			p.Styles = swap(p.Styles, "#111", "#222")
		}), nil, []string{"a", "b", "c"}},
		{"the order of rules", with(func(p *MockupParts) {
			p.Styles = swap(p.Styles, ".tc-btn{padding:4px 8px}\n.tc-kbd{font-size:11px}", ".tc-kbd{font-size:11px}\n.tc-btn{padding:4px 8px}")
		}), nil, []string{"a", "b", "c"}},
		{"only spaces", with(func(p *MockupParts) {
			p.Styles = swap(p.Styles, ".tc-btn{padding:4px 8px}", ".tc-btn {\n  padding:4px  8px\n}")
		}), nil, nil},
		{"the sprite", with(func(p *MockupParts) { p.Sprite = "<svg><symbol/></svg>" }), nil, []string{"a", "b", "c"}},
		{"a rule that needs two classes", with(func(p *MockupParts) {
			p.Styles = swap(p.Styles, "</style>", ".tc-sheet .tc-kbd{color:red}</style>")
		}), nil, nil},
		{"a rule with an exception", with(func(p *MockupParts) {
			p.Styles = swap(p.Styles, "</style>", ".tc-kbd:not(.tc-sheet){color:red}</style>")
		}), nil, []string{"b"}},
		{"a rule with a value", with(func(p *MockupParts) {
			p.Styles = swap(p.Styles, "</style>", `[data-tc-theme="a.tc-x"] .tc-num{color:red}</style>`)
		}), nil, []string{"c"}},
		{"a frame and a rule", with(func(p *MockupParts) {
			p.Frames[1].HTML = `<div class="frame"><kbd class="tc-kbd">Enter</kbd></div>`
			p.Styles = swap(p.Styles, "font-size:11px", "font-size:12px")
		}), []string{"b"}, nil},
	} {
		p := PlanStrays(base, c.after)
		if !reflect.DeepEqual(p.Meant, c.meant) {
			t.Errorf("%s: meant %v, want %v", c.name, p.Meant, c.meant)
		}
		if !reflect.DeepEqual(p.Check, c.check) {
			t.Errorf("%s: to check %v, want %v (%s)", c.name, p.Check, c.check, p.Why)
		}
		if len(p.Check) > 0 && p.Why == "" {
			t.Errorf("%s: frames to check without a reason", c.name)
		}
	}

	p := PlanStrays(base, with(func(p *MockupParts) {
		p.Frames = append(p.Frames[:1:1], MockupFrame{ID: "d", HTML: `<div class="frame"></div>`})
	}))
	if !reflect.DeepEqual(p.New, []string{"d"}) || !reflect.DeepEqual(p.Gone, []string{"b", "c"}) {
		t.Errorf("new %v, gone %v", p.New, p.Gone)
	}
}

// TestPlanStraysOnTheMockup: the mockup against itself has nothing to check,
// and its rules are all read — a rule lost by the reader would be a change
// nobody sees.
func TestPlanStraysOnTheMockup(t *testing.T) {
	root, err := FindModuleRoot()
	if err != nil {
		t.Fatal(err)
	}
	html, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(mockupRel)))
	if err != nil {
		t.Fatal(err)
	}
	parts, err := ParseMockupParts(string(html))
	if err != nil {
		t.Fatal(err)
	}
	if p := PlanStrays(parts, parts); len(p.Check) != 0 || len(p.Meant) != 0 || p.Same != len(parts.Frames) {
		t.Errorf("the mockup against itself: %d meant, %d to check, %d the same of %d", len(p.Meant), len(p.Check), p.Same, len(parts.Frames))
	}
	css := styleCSS(parts.Styles)
	rules := flatRules(css)
	if len(rules) < 500 {
		t.Fatalf("%d rules read from the mockup", len(rules))
	}
	n := 0
	for _, r := range rules {
		n += len(r.sel) + len(r.body)
	}
	if clean := len(straySpace.ReplaceAllString(stripCSSComments(css), "")); n < clean*8/10 {
		t.Errorf("the rules read hold %d characters of %d", n, clean)
	}
}
