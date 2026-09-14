package main

import (
	"strings"
	"testing"
)

func TestApplyColourAndMeter(t *testing.T) {
	in := `<style>body{}</style><style data-tc>.tc-mark{width:3px}.tc-meter__bar{height:3px}</style>
<div class="frame" data-frame="s-notes-1">
  <span class="tc-mark" style="background:#4A6B52"></span>
  <span class="tc-meter"><i style="width:83%" class="tc-meter__bar"></i></span>
</div>`
	got, rows, err := apply(in)
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) == 0 {
		t.Fatal("expected audit rows")
	}
	for _, w := range []string{
		`data-tc-colour="#4A6B52"`,
		`data-tc-width="83%"`,
		`[data-tc-colour][data-tc-colour]{background:var(--tc-colour)}`,
		`[data-tc-width][data-tc-width]{width:var(--tc-width)}`,
		`[data-tc-colour="#4A6B52"]{--tc-colour:#4A6B52}`,
		`[data-tc-width="83%"]{--tc-width:83%}`,
		`/* layout glue */`,
		`/* frame data values */`,
	} {
		if !strings.Contains(got, w) {
			t.Errorf("missing %q in:\n%s", w, got)
		}
	}
	if strings.Contains(got, `style="background:#4A6B52"`) {
		t.Fatal("leftover mark style")
	}
	if strings.Contains(got, `style="width:83%"`) {
		t.Fatal("leftover meter style")
	}
	again, _, err := apply(got)
	if err != nil {
		t.Fatal(err)
	}
	if again != got {
		t.Fatal("second apply changed file")
	}
}

func TestApplyGlueAndColumn(t *testing.T) {
	in := `<style>body{}</style><style data-tc>.tc-table{width:100%}</style>
<div class="frame" data-frame="s-files-1">
  <span style="margin-left:auto" class="tc-mono">2 min ago</span>
  <div style="display:flex;align-items:center;gap:8px;flex-wrap:wrap"></div>
  <div style="display:flex;flex-direction:column;gap:6px"></div>
  <table class="tc-table"><thead><tr><th style="width:28px"></th><th style="width:92px">Size</th></tr></thead></table>
</div>
<div class="frame" data-frame="s-files-2">
  <div style="display:flex;align-items:center;gap:8px;flex-wrap:wrap"></div>
  <div style="display:flex;flex-direction:column;gap:6px"></div>
</div>`
	got, _, err := apply(in)
	if err != nil {
		t.Fatal(err)
	}
	for _, w := range []string{
		`class="tc-mono tc-push"`,
		`tc-line tc-line--gap-8`,
		`tc-stack tc-stack--gap-6`,
		`tc-table__col--tick`,
		`tc-table__col--size`,
		`.tc-push.tc-push{margin-left:auto}`,
		`.tc-line.tc-line{display:flex;align-items:center;flex-wrap:wrap}`,
		`.tc-line--gap-8.tc-line--gap-8{gap:8px}`,
		`.tc-stack--gap-6.tc-stack--gap-6{gap:6px}`,
	} {
		if !strings.Contains(got, w) {
			t.Errorf("missing %q in:\n%s", w, got)
		}
	}
	if strings.Contains(got, `style="`) {
		t.Fatalf("leftover style: %s", got)
	}
}

func TestApplyTintAndSkipPage(t *testing.T) {
	in := `<style>body{}</style><style data-tc>.tc-btn{height:26px}</style>
<div class="chip" style="background:#17211C"></div>
<div class="frame" data-frame="s-cal-1">
  <svg width="15" height="15" viewBox="0 0 24 24" stroke="var(--tc-accent)"><use href="#i-cal"/></svg>
</div>`
	got, _, err := apply(in)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, `class="chip" style="background:#17211C"`) {
		t.Fatal("page chrome style must stay")
	}
	if !strings.Contains(got, `tc-ico--accent`) {
		t.Fatalf("missing tint class: %s", got)
	}
	if strings.Contains(got, `stroke="var(--tc-accent)"`) {
		t.Fatal("leftover svg stroke")
	}
	if strings.Contains(got, `viewbox=`) {
		t.Fatal("viewBox must keep its case")
	}
}

func TestApplyKeepsFrameWidth(t *testing.T) {
	in := `<style>body{}</style><style data-tc>.tc-root{color:red}</style>
<div class="frame" data-frame="m-phone-1" style="width:392px">
  <span class="sub" style="letter-spacing:.1em">phone</span>
  <div class="tc-root"><span class="tc-mark" style="background:#4A6B52"></span></div>
</div>`
	got, _, err := apply(in)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, `style="width:392px"`) {
		t.Fatalf("frame width must stay: %s", got)
	}
	if !strings.Contains(got, `class="sub" style="letter-spacing:.1em"`) {
		t.Fatal("page class sub must keep style")
	}
}

func TestApplyLeftoverDoesNotClobberGlue(t *testing.T) {
	in := `<style>body{}</style><style data-tc>.tc-in{height:26px}</style>
<div class="frame" data-frame="s-notes-1">
  <span style="margin-left:auto;color:var(--tc-dusty);font-size:11px">sftpgo</span>
  <button class="tc-in is-pick" style="width:240px;font-size:12px;margin-left:auto" type="button">Folder</button>
</div>`
	got, _, err := apply(in)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, `.tc-push.tc-push{margin-left:auto}`) {
		t.Fatal("glue tc-push missing")
	}
	if strings.Contains(got, `.tc-push.tc-push{font-size`) || strings.Contains(got, `.tc-push.tc-push{color`) {
		t.Fatalf("leftover clobbered tc-push: %s", got)
	}
	if !strings.Contains(got, `width:240px`) {
		t.Fatalf("pick width lost: %s", got)
	}
}

func TestApplyDoesNotTouchFrameLabel(t *testing.T) {
	in := `<style>body{}</style><style data-tc>.tc-root{color:red}</style>
<div class="frame" data-frame="s-notes-1">
  <div class="frame-label"><span class="frag" style="border:1px dashed red">компонент</span></div>
  <div class="tc-root"><span class="tc-mark" style="background:#4A6B52"></span></div>
</div>`
	got, _, err := apply(in)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, `style="border:1px dashed red"`) {
		t.Fatal("frame-label style must stay")
	}
}

func TestApplyLeftoverBeatsTwoClassProduct(t *testing.T) {
	in := `<style>body{}</style><style data-tc>.tc-in.is-area{min-height:76px}</style>
<div class="frame" data-frame="s-notes-1">
  <textarea class="tc-in is-area" style="min-height:132px;font-size:14px">note</textarea>
</div>`
	got, _, err := apply(in)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, `.tc-in--r1.tc-in--r1{`) {
		t.Fatalf("leftover must be two-class: %s", got)
	}
	if !strings.Contains(got, `min-height:132px`) {
		t.Fatalf("textarea min-height lost: %s", got)
	}
}
