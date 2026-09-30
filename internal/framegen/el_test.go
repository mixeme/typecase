package framegen

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestParseKeepsWhatTheDrawingSays(t *testing.T) {
	e := Parse(`<p class="a b"><b>x</b> <i>y</i><button disabled></button><button disabled="">z</button></p>`)
	if e.Tag != "p" || !e.Has("b") || e.Has("c") {
		t.Fatalf("root %q %v", e.Tag, e.Cls())
	}
	if len(e.Kids) != 5 || e.Kids[1] != " " {
		t.Errorf("the space between two inline elements is a child: %q", e.Kids)
	}
	bs := e.Els()
	if !bs[2].HasAttr("disabled") || bs[2].Valued("disabled") {
		t.Error("a bare attribute is present without a value")
	}
	if !bs[3].Valued("disabled") {
		t.Error(`disabled="" has a value`)
	}
}

func TestSpaced(t *testing.T) {
	e := Parse("<p>\n  a  <b>x</b>  b \n</p>")
	for _, c := range []struct {
		kid  int
		want string
	}{{0, "a "}, {2, " b"}} {
		got, _ := spaced(e, e.Kids[c.kid].(string))
		if got != c.want {
			t.Errorf("kid %d: %q, want %q", c.kid, got, c.want)
		}
	}
	e = Parse("<p><b>x</b> <b>y</b></p>")
	if got, ok := spaced(e, " "); !ok || got != " " {
		t.Errorf("a space between two elements stays: %q", got)
	}
}

func TestFuncName(t *testing.T) {
	for id, want := range map[string]string{"s-find-2": "Find2", "desktop-10": "Desktop10", "rest-1": "Rest1", "s-auth_x-3": "AuthX3"} {
		if got := funcName(id); got != want {
			t.Errorf("%s: %s, want %s", id, got, want)
		}
	}
}

// TestUnknownIsAGap: an element the generator does not know stops the frame
// with its name; nothing is guessed.
func TestUnknownIsAGap(t *testing.T) {
	en, err := loadEnums(filepath.Join("..", "..", "widget"))
	if err != nil {
		t.Fatal(err)
	}
	g := &gen{en: en, faces: en.faces()}
	if _, why := g.translate(`<div class="tc-shell"><div class="tc-shell__body"><main class="tc-main"><div class="tc-nothing"></div></main></div></div>`); why == "" {
		t.Error("an unknown block made a screen")
	}
	expr, why := g.translate(`<div class="tc-root"><div class="tc-msg">Hello <b>there</b></div></div>`)
	if why != "" || !strings.Contains(expr, `w.Msg{Parts: []w.MsgPart{w.Text("Hello "), w.Strong("there")}}`) {
		t.Errorf("page with a message: %s %s", expr, why)
	}
}
