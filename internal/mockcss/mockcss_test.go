package mockcss

import (
	"strings"
	"testing"
)

func TestGenerateKeepsOnlyOwnContext(t *testing.T) {
	css := `/* note */
.m-btn{height:26px;border:.5px solid var(--card-border)}
.m-bar > .m-btn,.m-rail-foot > .m-btn{padding:0 7px}
.mock.is-phone .m-btn{height:38px}
.mock{--pad:20px;color:var(--ink)}
@container (max-width:560px){.m-btn{height:30px}.m-bar{gap:0}}
`
	got := Generate(css, Options{
		Keep:    map[string]bool{"mock": true, "m-btn": true, "m-rail-foot": true},
		Classes: map[string]string{"mock": "tcx", "m-btn": "tcx-btn", "m-rail-foot": "tcx-rail-foot"},
		Vars:    map[string]string{"card-border": "tc-edge", "ink": "tc-ink", "pad": "tc-page-gutter"},
		Prefix:  "tcx-",
	})
	want := `.tcx-btn{height:26px;border:.5px solid var(--tc-edge)}
.tcx-rail-foot > .tcx-btn{padding:0 7px}
.tcx{--tc-page-gutter:20px;color:var(--tc-ink)}
@container (max-width:560px){
.tcx-btn{height:30px}
}
`
	if got != want {
		t.Fatalf("got\n%s\nwant\n%s", got, want)
	}
	if strings.Contains(got, "38px") || strings.Contains(got, "m-bar") {
		t.Fatal("rule of another context leaked")
	}
}
