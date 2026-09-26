// Package proof renders mockup frames the "no translation" way: each widget
// prints the frame's markup with renamed classes, and the CSS is taken from the
// mockup by internal/ Not part of the library API.
package proof

import (
	"strconv"
	"strings"
)

// Rail is .m-rail: nav, source sections, the foot.
type Rail struct {
	Nav      []NavItem
	Sections []Section
	Foot     []Btn
	About    string
	Build    string
}

// NavItem is one .m-nav link.
type NavItem struct {
	Icon  string
	Label string
	Href  string
	Count *int
	On    bool
}

// Section is .m-rail-sec: an optional rubric, then links, rows, tags or shortcuts.
type Section struct {
	Rubric string
	Nav    []NavItem
	Rows   []Source
	Tags   []Tag
	Tree   []Shortcut
}

// Source is .m-src: tick, colour bar (or icon at the root), name, meta.
// Error is a source that did not answer; Action follows its meta.
type Source struct {
	Label     string
	Mark      Mark
	Root      string // icon of the root row instead of the colour bar
	Meta      string
	MetaTitle string
	Badge     string // instead of Meta
	Box       Box
	All       bool
	On        bool
	Off       bool
	Error     bool
	Action    *Btn
}

// Shortcut is one row of the rail tree.
type Shortcut struct {
	Icon  string
	Label string
	Meta  string
	On    bool
}

func (r Rail) Render() string {
	var b strings.Builder
	b.WriteString(`<nav class="tcx-rail" aria-label="Sections">` + navLinks(r.Nav))
	for _, s := range r.Sections {
		b.WriteString(`<div class="tcx-rail-sec">`)
		if s.Rubric != "" {
			b.WriteString(`<div class="tcx-rubric">` + esc(s.Rubric) + `</div>`)
		}
		if len(s.Nav) > 0 {
			b.WriteString(navLinks(s.Nav))
		}
		for _, row := range s.Rows {
			b.WriteString(row.Render())
		}
		if len(s.Tags) > 0 {
			b.WriteString(`<div class="tcx-rail-tags">`)
			for _, t := range s.Tags {
				b.WriteString(t.Render())
			}
			b.WriteString(`</div>`)
		}
		if len(s.Tree) > 0 {
			b.WriteString(`<div class="tcx-tree tcx-rail-tree">`)
			for _, t := range s.Tree {
				b.WriteString(`<div` + onAttr(t.On) + `>` + Icon{Name: t.Icon}.Render() + esc(t.Label) +
					`<span class="tcx-meta tcx-aside">` + esc(t.Meta) + `</span></div>`)
			}
			b.WriteString(`</div>`)
		}
		b.WriteString(`</div>`)
	}
	if len(r.Foot) > 0 || r.About != "" {
		b.WriteString(`<div class="tcx-rail-foot">` + btns(r.Foot) +
			`<a class="tcx-rail-about" href="#about">` + esc(r.About) + ` <span class="mono">` + esc(r.Build) + `</span></a></div>`)
	}
	b.WriteString(`</nav>`)
	return b.String()
}

func navLinks(items []NavItem) string {
	var b strings.Builder
	b.WriteString(`<div class="tcx-nav">`)
	for _, n := range items {
		b.WriteString(`<a href="` + esc(n.Href) + `"` + onAttr(n.On) + `>` + Icon{Name: n.Icon}.Render() + esc(n.Label))
		if n.Count != nil {
			b.WriteString(`<span class="tcx-num">` + strconv.Itoa(*n.Count) + `</span>`)
		}
		b.WriteString(`</a>`)
	}
	b.WriteString(`</div>`)
	return b.String()
}

func (s Source) Render() string {
	c := "tcx-src"
	for _, m := range []struct {
		on  bool
		cls string
	}{{s.Root != "", "is-root"}, {s.All, "is-all"}, {s.On, "is-on"}, {s.Off, "is-off"}, {s.Error, "is-error"}} {
		if m.on {
			c += " " + m.cls
		}
	}
	lead := s.Mark.Render()
	if s.Root != "" {
		lead = Icon{Name: s.Root, Box: true}.Render()
	}
	title := ""
	if s.MetaTitle != "" {
		title = ` title="` + esc(s.MetaTitle) + `"`
	}
	tail := `<span class="tcx-meta"` + title + `>` + esc(s.Meta) + `</span>`
	if s.Badge != "" {
		tail = Badge{Label: s.Badge}.Render()
	}
	if s.Action != nil {
		tail += s.Action.Render()
	}
	box := s.Box
	box.Label = "Toggle " + s.Label
	return `<div class="` + c + `">` + box.Render() + lead + `<span class="t">` + esc(s.Label) + `</span>` + tail + `</div>`
}

func btns(bs []Btn) string {
	var b strings.Builder
	for _, x := range bs {
		b.WriteString(x.Render())
	}
	return b.String()
}

func num(v int) *int { return &v }
