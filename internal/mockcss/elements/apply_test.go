package main

import (
	"strings"
	"testing"
)

func TestApplyButtonsAndFields(t *testing.T) {
	in := `<style data-tc>
.tc-root svg{display:block}
.tc-btn > .tc-menu.is-pop{margin-top:6px;font-weight:400;text-align:left}
/* элементы управления */
.tc-btn{height:26px}
</style><style>body{}</style>
<div class="frame" data-frame="s-notes-1">
  <div class="frame-label"><span class="tc-kbd">Esc</span></div>
  <div class="tc-root">
    <span class="tc-btn is-primary">Save</span>
    <div class="tc-field tc-field--sm"><label>Birthday</label><span class="tc-in is-ph">YYYY-MM-DD</span></div>
    <div class="tc-field tc-field--xs"><label>Start time</label><span class="tc-in is-mono">10:00</span></div>
    <div class="tc-field"><label>Name</label><span class="tc-in">Анна</span></div>
    <span class="tc-seg"><span class="is-on tc-seg__opt">A</span><span class="tc-seg__opt">B</span></span>
    <a class="tc-nav__item">Notes</a>
    <div class="tc-menu"><div class="tc-menu__item">Open</div></div>
    <span class="tc-kbd">Ctrl+S</span>
    <span class="tc-check"><span class="tc-tickrow__tick"></span>All day</span>
  </div>
</div>`
	got, err := apply(in)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		controlMarker,
		`<button class="tc-btn is-primary" type="button">Save</button>`,
		`type="date"`,
		`placeholder="YYYY-MM-DD"`,
		`type="time"`,
		`value="10:00"`,
		`value="Анна"`,
		`<button class="is-on tc-seg__opt" type="button">A</button>`,
		`href="#"`,
		`role="menuitem"`,
		`<kbd class="tc-kbd">Ctrl+S</kbd>`,
		`<button class="tc-check" type="button">`,
		`.tc-drop:not(.is-bar) > .tc-menu.is-pop`,
	}
	for _, w := range want {
		if !strings.Contains(got, w) {
			t.Errorf("missing %q in:\n%s", w, got)
		}
	}
	if strings.Contains(got, `<span class="tc-btn`) {
		t.Errorf("leftover span.tc-btn")
	}
	// frame-label kbd is page chrome of the frame, still converted (kbd everywhere)
	if !strings.Contains(got, `frame-label`) {
		t.Fatal("lost frame-label")
	}
	again, err := apply(got)
	if err != nil {
		t.Fatal(err)
	}
	if again != got {
		t.Fatalf("second apply changed file")
	}
}

func TestApplyDropExtractsMenu(t *testing.T) {
	in := `<style data-tc>/* элементы управления */</style><style>body{}</style>
<div class="frame" data-frame="s-cal-1"><div class="tc-root">
<span class="tc-btn tc-bar__more" title="More"><svg></svg>
  <div class="tc-menu is-pop is-right"><div class="tc-menu__item">Hide</div></div>
</span>
</div></div>`
	got, err := apply(in)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, `class="tc-drop"`) {
		t.Fatalf("no tc-drop:\n%s", got)
	}
	if strings.Contains(got, `<button type="button" class="tc-btn tc-bar__more"`) &&
		strings.Contains(got, `<div class="tc-menu`) {
		if strings.Contains(got, `<button type="button" class="tc-btn tc-bar__more" title="More"><svg></svg>
  <div class="tc-menu`) || strings.Contains(got, `tc-bar__more" title="More"><svg></svg>`) && strings.Contains(got, `</button>`) && !strings.Contains(got, `</button><`) {
			// menu must not remain inside the button
		}
	}
	btnEnd := strings.Index(got, `class="tc-btn tc-bar__more"`)
	if btnEnd < 0 {
		t.Fatalf("button missing:\n%s", got)
	}
	openBtn := strings.Index(got[btnEnd:], ">")
	closeBtn := strings.Index(got[btnEnd:], "</button>")
	menuAt := strings.Index(got, `class="tc-menu`)
	if menuAt < 0 || closeBtn < 0 {
		t.Fatalf("structure:\n%s", got)
	}
	menuAbs := strings.Index(got, `class="tc-menu`)
	btnCloseAbs := btnEnd + closeBtn
	if menuAbs > btnEnd+openBtn && menuAbs < btnCloseAbs {
		t.Fatalf("menu still inside button:\n%s", got)
	}
}

func TestInjectCSSRefreshesControlBlock(t *testing.T) {
	in := `<style data-tc>
` + controlMarker + `
.tc-root .tc-q{flex:1;min-width:0}
/* элементы управления */
.tc-btn{height:26px}
</style>`
	got := injectCSS(in)
	if !strings.Contains(got, "::-webkit-search-cancel-button") {
		t.Fatalf("did not refresh control CSS:\n%s", got)
	}
	if strings.Count(got, controlMarker) != 1 {
		t.Fatalf("duplicate marker:\n%s", got)
	}
	if strings.Count(got, "/* элементы управления */") != 1 {
		t.Fatalf("lost controls comment:\n%s", got)
	}
}

func TestApplySearchAndFilter(t *testing.T) {
	in := `<style data-tc>/* элементы управления */</style><style>body{}</style>
<div class="demo tc-root">
  <div class="tc-search"><svg width="13" height="13"></svg>Search everything…</div>
  <span class="tc-in is-ph" style="width:190px"><svg width="13" height="13"></svg>Filter this folder</span>
  <span class="tc-in is-pick">Confirmed</span>
  <span class="tc-in is-area">Two
lines</span>
</div>`
	got, err := apply(in)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, `value="Search everything…"`) || !strings.Contains(got, `class="tc-q"`) {
		t.Fatalf("search:\n%s", got)
	}
	if !strings.Contains(got, `<label class="tc-in is-ph"`) {
		t.Fatalf("filter label:\n%s", got)
	}
	if !strings.Contains(got, `value="Filter this folder"`) {
		t.Fatalf("filter value:\n%s", got)
	}
	if !strings.Contains(got, `<button class="tc-in is-pick" type="button">Confirmed</button>`) {
		t.Fatalf("pick:\n%s", got)
	}
	if !strings.Contains(got, `<textarea class="tc-in is-area">`) {
		t.Fatalf("textarea:\n%s", got)
	}
}

func TestTickInsideMenuStaysSpan(t *testing.T) {
	in := `<style data-tc>/* элементы управления */</style><style>body{}</style>
<div class="frame" data-frame="x-1"><div class="tc-root">
<div class="tc-menu"><div class="tc-menu__item"><span class="tc-tickrow__tick"></span>Name</div></div>
<span class="tc-btn"><span class="tc-tick"></span>Mark done</span>
</div></div>`
	got, err := apply(in)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(got, `role="menuitem"><button`) {
		t.Fatalf("nested button in menu:\n%s", got)
	}
	if !strings.Contains(got, `<span class="tc-tickrow__tick"></span>Name`) {
		t.Fatalf("tick not span:\n%s", got)
	}
	if !strings.Contains(got, `<button class="tc-btn" type="button"><span class="tc-tick"></span>Mark done</button>`) {
		t.Fatalf("mark done:\n%s", got)
	}
}

func TestApplyLeavesCompositeField(t *testing.T) {
	in := `<style data-tc>/* элементы управления */</style><style>body{}</style>
<div class="frame" data-frame="x-1"><div class="tc-root">
<span class="tc-in is-focus"><span class="tc-tag">A</span> rest</span>
</div></div>`
	got, err := apply(in)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(got, `<input`) {
		t.Fatalf("composite became input:\n%s", got)
	}
	if !strings.Contains(got, `class="tc-in is-focus"`) {
		t.Fatalf("lost field:\n%s", got)
	}
}
