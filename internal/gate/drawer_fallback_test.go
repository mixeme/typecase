package gate

import (
	"strings"
	"testing"
)

func drawerPage(css string) string {
	return "<html><head><style data-tc>" + css + "</style><style>.frame{}</style></head><body></body></html>"
}

// TestDrawerFallback: the drawer's style() block passes only with its copy
// right after it, rule for rule, under the window's width and the shell.
func TestDrawerFallback(t *testing.T) {
	block := "@container style(--tc-shell-mode: drawer){\n  /* a note */\n  .tc-row{min-height:52px;gap:10px}\n  .tc-btn,.tc-in:has(> .x, .y){height:38px}\n}\n"
	copyOK := "@media (max-width:640px){\n  :where(.tc-shell,.tc-overlay) .tc-row{min-height:52px; gap:10px}\n" +
		"  :where(.tc-shell,.tc-overlay) .tc-btn,:where(.tc-shell,.tc-overlay) .tc-in:has(> .x, .y){height:38px}\n}\n"
	if err := checkMockupDrawerFallback(drawerPage(".a{color:red}\n" + block + copyOK + ".b{color:red}")); err != nil {
		t.Fatalf("a true copy is refused: %v", err)
	}
	for name, css := range map[string]string{
		"no copy":         block,
		"copy apart":      block + ".b{color:red}\n" + copyOK,
		"a rule changed":  block + strings.Replace(copyOK, "52px", "50px", 1),
		"a rule missing":  block + strings.Replace(copyOK, "  :where(.tc-shell,.tc-overlay) .tc-row{min-height:52px; gap:10px}\n", "", 1),
		"no shell scope":  block + strings.ReplaceAll(copyOK, ":where(.tc-shell,.tc-overlay) ", ""),
		"no drawer block": ".a{color:red}",
	} {
		if err := checkMockupDrawerFallback(drawerPage(css)); err == nil {
			t.Errorf("%s: passed", name)
		}
	}
}
