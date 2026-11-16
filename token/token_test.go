package token

import (
	"strings"
	"testing"
)

const complete = `:root{--tc-ink:#000;--tc-surface:#111;--tc-bg:#222;--tc-card:#777;--tc-rule:#333;
--tc-muted:#444;--tc-accent:#555;--tc-danger:#666;--tc-progress:#6a6;--tc-edge:#888;--tc-hair:#999;--tc-dusty:#aaa;
--tc-font-serif:a;--tc-font-sans:b;--tc-font-mono:c}
.tc-root{--tc-rail:264px;--tc-side:400px;--tc-top:52px;--tc-page-gutter:20px;--tc-page-gutter-narrow:14px}
[data-tc-colour]{background:var(--tc-colour)}`

func TestValidate_complete(t *testing.T) {
	if err := Validate(complete); err != nil {
		t.Fatal(err)
	}
}

func TestValidate_missing(t *testing.T) {
	err := Validate(":root { --tc-ink: #000; }")
	if err == nil || !strings.Contains(err.Error(), "--tc-surface") {
		t.Fatalf("want missing --tc-surface, got %v", err)
	}
}

func TestValidate_undefined(t *testing.T) {
	err := Validate(complete + ".tc-btn{border-radius:var(--tc-radius)}")
	if err == nil || !strings.Contains(err.Error(), "--tc-radius") {
		t.Fatalf("want undefined --tc-radius, got %v", err)
	}
}
