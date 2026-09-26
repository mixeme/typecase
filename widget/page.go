package widget

import "strings"

// BannerEnd is what stands at the end of a public page's banner.
type BannerEnd interface {
	part
	bannerEnd()
}

func (Line) bannerEnd() {}
func (Push) bannerEnd() {}

// Banner is the head of a page opened by a public link: the logo, what the
// visitor may do with it, the downloads at the end.
type Banner struct {
	Logo  Logo
	Badge *Badge
	End   BannerEnd // pushed to the end
	Wide  bool      // 14 px between the pieces: a page of a calendar
}

func (b Banner) html(mix) string {
	s := `<div class="tc-line tc-line--gap-14 tc-main__banner">`
	if !b.Wide {
		s = `<div class="tc-line tc-line--gap-12 tc-main__banner--12">`
	}
	s += b.Logo.html(nil)
	if b.Badge != nil {
		s += b.Badge.html(nil)
	}
	if b.End != nil {
		s += b.End.html(nil)
	}
	return s + `</div>`
}

// StripPart is what a public page's address strip holds.
type StripPart interface {
	part
	stripPart()
}

func (Input) stripPart()  {}
func (Button) stripPart() {}
func (Dim) stripPart()    {}

// Strip is the line of a public page that gives its address to a client:
// what to do, the address, the buttons, what the address gives.
type Strip struct {
	Label string
	Items []StripPart
}

func (p Strip) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-line tc-line--gap-8 tc-main__strip"><span class="tc-dim--muted tc-dim--t12">` + esc(p.Label) + `</span>`)
	for _, it := range p.Items {
		s.WriteString(it.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}

// MainFoot is the last line of a public page: how long the link lives.
type MainFoot struct {
	Text  string
	Roomy bool // 16 px above: after a list
}

func (f MainFoot) html(mix) string {
	return `<div class="` + pick(f.Roomy, "tc-main__foot--16") + pick(!f.Roomy, "tc-main__foot") + `">` + esc(f.Text) + `</div>`
}

// Stage is the middle of a public page that shows one file it cannot
// preview: what the file is, the one button, why there is no preview.
type Stage struct {
	Title    string
	Meta     Mono
	Download Button
	Note     string
}

func (g Stage) html(mix) string {
	return `<div class="tc-stack tc-stack--gap-14 tc-main__stage"><div class="tc-stack__item"><div class="tc-head__title tc-head__title--t18">` + esc(g.Title) +
		`</div><div class="tc-head__sub tc-head__sub--jc-center">` + g.Meta.html(nil) + `</div></div>` + g.Download.html(nil) +
		`<div class="tc-main__stage-note">` + esc(g.Note) + `</div></div>`
}

// Person is the head of a person's page on a phone: the photo, the name,
// where the person works, and the buttons under it.
type Person struct {
	Initials string // shown while there is no photo
	Name     string
	Sub      string
	Badge    *Badge
	Acts     []Button
}

func (p Person) html(mix) string {
	var s strings.Builder
	s.WriteString(`<div class="tc-line tc-line--gap-12 tc-main__person"><span` + class(nil, "tc-avatar--photo", is("ph", p.Initials != ""), "tc-avatar--photo--sq62-t17") + `>` +
		esc(p.Initials) + `</span><div class="tc-main__person-side"><div class="tc-main__person-name">` + esc(p.Name) + `</div><div class="tc-head__sub">` + esc(p.Sub))
	if p.Badge != nil {
		s.WriteString(" " + p.Badge.html(nil))
	}
	s.WriteString(`</div></div></div><div class="tc-line tc-line--gap-6 tc-main__person-acts">`)
	for _, b := range p.Acts {
		s.WriteString(b.html(nil))
	}
	s.WriteString(`</div>`)
	return s.String()
}
