package gate

import (
	"strings"
	"testing"
)

// Bite: one changed byte in the library file is red, with where it is.
func TestBite_cssFromMockup(t *testing.T) {
	want := ":root{--tc-ink:#17211C}\n.tc-btn{height:26px}\n"
	if err := compareCut(want, want); err != nil {
		t.Fatalf("same text is red: %v", err)
	}
	got := strings.Replace(want, "26px", "28px", 1)
	err := compareCut(want, got)
	if err == nil || !strings.Contains(err.Error(), "line 2") {
		t.Fatalf("changed byte not caught with its line: %v", err)
	}
	if err := compareCut(want, want+"\n"); err == nil {
		t.Fatal("extra byte at the end not caught")
	}
}
