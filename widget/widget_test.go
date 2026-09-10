package widget_test

import (
	"regexp"
	"strings"
	"testing"

	"carrel.dev/typecase/css"
	"carrel.dev/typecase/internal/catalog"
	"carrel.dev/typecase/token"
	"carrel.dev/typecase/widget"
)

func intPtr(n int) *int { return &n }

func TestRender_smoke(t *testing.T) {
	cases := []struct {
		name string
		html string
	}{
		{"btn", widget.Button{Label: "Go"}.Render()},
		{"input", widget.Input{Name: "x", Kind: widget.InputText}.Render()},
		{"check", widget.Check{Name: "c", Label: "Check"}.Render()},
		{"kbd", widget.Kbd{Keys: []string{"Esc"}}.Render()},
		{"close", widget.Close{For: "x", Label: "Close"}.Render()},
		{"brand", widget.Brand{Mark: "T", Word: "Test"}.Render()},
		{"mark", widget.Mark{Colour: "#000", Label: "A"}.Render()},
		{"num", widget.Num{Value: intPtr(1)}.Render()},
		{"badge", widget.Badge{Label: "New"}.Render()},
		{"avatar", widget.Avatar{Initials: "AB"}.Render()},
		{"tick", widget.Tick{Label: "Done"}.Render()},
		{"tag", widget.Tag{Label: "Tag"}.Render()},
		{"msg", widget.Msg{Title: "Hi", Text: "There"}.Render()},
		{"empty", widget.Empty{Title: "Empty", Text: "None", Action: widget.Action{Label: "Add"}}.Render()},
		{"skel", widget.Skel{Lines: 2}.Render()},
		{"meter", widget.Meter{Value: 1, Max: 10, Label: "1/10"}.Render()},
		{"form", widget.Form{Action: "/f", Sets: []widget.FieldSet{{Legend: "L", Rows: []widget.Field{{Name: "n", Label: "N", Kind: widget.InputText}}}}}.Render()},
		{"menu", widget.Menu{Label: "Menu", Items: []widget.MenuItem{{Label: "One", Href: "#"}}}.Render()},
		{"bar", widget.Bar{Left: []widget.Node{widget.Action{Label: "A"}}}.Render()},
		{"list", widget.List{Groups: []widget.ListGroup{{Rows: []widget.Row{{Cells: widget.RowCells{Title: "T"}}}}}}.Render()},
		{"table", widget.Table{Columns: []widget.Column{{Label: "C"}}, Rows: []widget.TableRow{{Cells: []string{"v"}}}}.Render()},
		{"shell", widget.Shell{Main: widget.Main{Label: "M", Body: []widget.Node{widget.Paragraph("b")}}}.Render()},
		{"doc", widget.Doc{Body: widget.Paragraph("x")}.Render()},
		{"docedit", widget.DocEdit{Title: "T", Text: "b"}.Render()},
		{"pfield", widget.PField{Name: "t", Label: "Title", Kind: widget.InputText}.Render()},
		{"pick", widget.Pick{Name: "r", Label: "Related"}.Render()},
		{"search", widget.Search{Action: "/s"}.Render()},
		{"stat", widget.Stat{Label: "Users", Value: "1,234"}.Render()},
		{"stats", widget.Stats{Items: []widget.Stat{{Label: "Users", Value: "1,234"}}}.Render()},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.html == "" {
				t.Fatal("empty render")
			}
			if strings.Contains(tc.html, "<script") {
				t.Fatal("inline script in output")
			}
		})
	}
}

func TestRowAvatar_fallbackNoImageSrc(t *testing.T) {
	html := widget.Row{
		Variant: widget.RowAvatar,
		Cells:   widget.RowCells{Title: "Alex", Subtitle: "alex@example.com"},
	}.Render()
	if strings.Contains(html, "<img") {
		t.Fatalf("avatar fallback must not emit img: %s", html)
	}
	if !strings.Contains(html, "tc-avatar__initials") {
		t.Fatalf("avatar fallback should use initials: %s", html)
	}
}

func TestButton_dangerSolid(t *testing.T) {
	outline := widget.Button{Label: "Detach", Kind: widget.ButtonDanger}.Render()
	if !strings.Contains(outline, "tc-btn--danger") || strings.Contains(outline, "tc-btn--danger-solid") {
		t.Fatalf("outline danger: %s", outline)
	}
	fill := widget.Button{Label: "Delete", Kind: widget.ButtonDangerSolid}.Render()
	if !strings.Contains(fill, "tc-btn--danger-solid") {
		t.Fatalf("solid danger: %s", fill)
	}
}

func TestCatalog_entries(t *testing.T) {
	entries := catalog.Entries()
	if len(entries) < 50 {
		t.Fatalf("expected comprehensive catalog, got %d entries", len(entries))
	}
	seen := make(map[string]struct{})
	for _, e := range entries {
		key := e.Widget + "/" + e.State
		if _, ok := seen[key]; ok {
			t.Fatalf("duplicate catalog entry %s", key)
		}
		seen[key] = struct{}{}
		if e.HTML == "" {
			t.Fatalf("empty HTML for %s", key)
		}
		if !strings.Contains(e.HTML, "tc-") {
			t.Fatalf("missing tc- class for %s", key)
		}
	}
}

func TestCSS_widgetTokenValidation(t *testing.T) {
	for _, name := range css.Order {
		if name == css.LayoutSentinel || name == "00-tokens.css" || name == "01-base.css" {
			continue
		}
		data, err := css.Files().Open(name)
		if err != nil {
			t.Fatal(err)
		}
		buf := make([]byte, 1<<20)
		n, _ := data.Read(buf)
		_ = data.Close()
		if err := token.ValidateWidget(string(buf[:n])); err != nil {
			t.Fatalf("%s: %v", name, err)
		}
	}
}

func TestEscape_html(t *testing.T) {
	html := widget.Button{Label: `<script>"&`}.Render()
	if strings.Contains(html, "<script>") {
		t.Fatal("unescaped HTML in label")
	}
}

// Пустой бейдж — ломаная фикстура ворот inv-badge-has-text. tc-usermenu отдавал
// её всякий раз, когда метка не задана (В-19).
func TestUserMenu_noEmptyBadge(t *testing.T) {
	html := widget.UserMenu{
		Name:  "Alex",
		Note:  "Signed in",
		Items: []widget.MenuItem{{Label: "Sign out", Href: "/logout"}},
	}.Render()
	if strings.Contains(html, `class="tc-badge`) {
		t.Fatalf("user menu without Mark must not render a badge: %s", html)
	}
	if strings.Contains(html, `<ul class="tc-usermenu__popover`) {
		t.Fatal("popover head must not sit inside the menu list")
	}
}

func TestBadge_emptyLabelRendersNothing(t *testing.T) {
	if got := (widget.Badge{}).Render(); got != "" {
		t.Fatalf("empty badge must render nothing, got %q", got)
	}
}

func TestRailSection_emptyTitleOmitsHeading(t *testing.T) {
	html := widget.RailSection{Rows: []widget.TickRow{{Label: "Combined notebook"}}}.Render()
	if strings.Contains(html, "tc-railsec__title") {
		t.Fatalf("empty title must omit heading: %s", html)
	}
	if !strings.Contains(html, "tc-railsec__rows") {
		t.Fatalf("rows must still render: %s", html)
	}
}

func TestRailSection_titledKeepsHeading(t *testing.T) {
	html := widget.RailSection{Title: "Views", Rows: []widget.TickRow{{Label: "All"}}}.Render()
	if !strings.Contains(html, `<h3 class="tc-railsec__title">Views</h3>`) {
		t.Fatalf("titled section must keep heading: %s", html)
	}
}

func TestRailSection_bodyWhenNoRows(t *testing.T) {
	html := widget.RailSection{
		Title: "Neighbours",
		Body:  []widget.Node{widget.Tree{Nodes: []widget.TreeNode{{ID: "a", Label: "One"}}}},
	}.Render()
	if strings.Contains(html, "tc-railsec__rows") {
		t.Fatalf("body section must not render rows container: %s", html)
	}
	if !strings.Contains(html, "tc-tree") {
		t.Fatalf("body section must render tree: %s", html)
	}
}

func TestRailSection_rowsTakePrecedenceOverBody(t *testing.T) {
	html := widget.RailSection{
		Rows: []widget.TickRow{{Label: "All"}},
		Body: []widget.Node{widget.Paragraph("ignored")},
	}.Render()
	if strings.Contains(html, "tc-prose") {
		t.Fatalf("rows must take precedence over body: %s", html)
	}
}

func TestSheet_bodyNodes(t *testing.T) {
	html := widget.Sheet{
		Title:  "Quick",
		Target: "Notes",
		Open:   true,
		Body:   []widget.Node{widget.Field{Name: "quick", Kind: widget.InputTextarea, Value: "draft"}},
	}.Render()
	if !strings.Contains(html, `class="tc-in`) {
		t.Fatalf("sheet body must render field nodes: %s", html)
	}
}

func TestTickRow_holdMark(t *testing.T) {
	html := widget.TickRow{Label: "Attachments"}.Render()
	if strings.Count(html, `class="tc-mark"`) != 1 {
		t.Fatalf("TickRow must print exactly one .tc-mark: %s", html)
	}
}

func TestRowExcerpt_directSlotChildren(t *testing.T) {
	html := widget.Row{
		Variant: widget.RowExcerpt,
		Cells:   widget.RowCells{Title: "Note", Excerpt: "Body", Date: "10 Aug", Source: "src"},
	}.Render()
	if strings.Contains(html, "tc-row__cells") {
		t.Fatalf("excerpt must not wrap slots in .tc-row__cells: %s", html)
	}
}

func TestRowExcerpt_dateAndSource(t *testing.T) {
	html := widget.Row{
		Variant: widget.RowExcerpt,
		Cells: widget.RowCells{
			Title:   "Note",
			Excerpt: "Body text",
			Date:    "10 Aug",
			Source:  "nextcloud · Notes",
		},
	}.Render()
	if !strings.Contains(html, `<span class="tc-row__date">10 Aug</span>`) {
		t.Fatalf("excerpt row must render date: %s", html)
	}
	if !strings.Contains(html, `<span class="tc-row__source">nextcloud · Notes</span>`) {
		t.Fatalf("excerpt row must render source: %s", html)
	}
}

func TestRowExcerpt_badgeLivesInTags(t *testing.T) {
	html := widget.Row{
		Variant: widget.RowExcerpt,
		Cells:   widget.RowCells{Title: "Note", Date: "10 Aug", Source: "src"},
		Badges:  []widget.Badge{{Label: "linked", Role: widget.BadgeLinked}},
	}.Render()
	tagsAt := strings.Index(html, `class="tc-row__tags"`)
	badgeAt := strings.Index(html, "tc-badge")
	if tagsAt < 0 || badgeAt < 0 || badgeAt < tagsAt {
		t.Fatalf("linked badge must sit in the tags cell: %s", html)
	}
	dateAt := strings.Index(html, `class="tc-row__date"`)
	if dateAt > 0 && badgeAt > dateAt {
		t.Fatalf("linked badge must come before the date column: %s", html)
	}
}

func TestRowAvatarCols_directSlotChildren(t *testing.T) {
	html := widget.Row{
		Variant: widget.RowAvatarCols,
		Cells: widget.RowCells{
			Avatar:   widget.Avatar{Initials: "AK"},
			Title:    "Alex",
			Subtitle: "+7",
			Detail:   "a@example.org",
			Caption:  "Press",
			Source:   "book",
		},
	}.Render()
	if strings.Contains(html, "tc-row__cells") {
		t.Fatalf("avatarcols must not wrap slots in .tc-row__cells: %s", html)
	}
	if !strings.Contains(html, `class="tc-avatar`) || !strings.Contains(html, `class="tc-row__subtitle"`) {
		t.Fatalf("avatarcols must print avatar and subtitle slots: %s", html)
	}
}

func TestRowTickCols_holdTick(t *testing.T) {
	html := widget.Row{
		Variant: widget.RowTickCols,
		Cells:   widget.RowCells{Name: "readme.txt", Size: "2 KB", Time: "10:30", Source: "sftp"},
	}.Render()
	if !strings.Contains(html, "tc-tick") {
		t.Fatalf("tickcols must print a tick stub: %s", html)
	}
	if strings.Contains(html, "tc-mark") {
		t.Fatalf("tickcols has no mark slot: %s", html)
	}
}

func TestSearchHints_useRowLead(t *testing.T) {
	html := widget.Search{
		Action:     "/s",
		Suggesting: true,
		History:    []widget.SearchRow{{Query: "q", Href: "/q"}},
		Jump:       []widget.JumpRow{{Label: "Alex", Hint: "Contact", Href: "/c"}},
	}.Render()
	if strings.Contains(html, "tc-row--find") {
		t.Fatalf("search hints must not use the find grid: %s", html)
	}
	if !strings.Contains(html, "tc-row--lead") {
		t.Fatalf("search hints must use RowLead: %s", html)
	}
}

func TestSearch_iconFirstInField(t *testing.T) {
	html := widget.Search{Action: "/s"}.Render()
	i := strings.Index(html, `class="tc-search__field"`)
	if i < 0 {
		t.Fatal("missing search field")
	}
	rest := html[i:]
	end := strings.Index(rest, "</div>")
	if end < 0 {
		t.Fatal("unclosed search field")
	}
	inner := rest[:end]
	iconAt := strings.Index(inner, "tc-search__icon")
	inputAt := strings.Index(inner, "tc-search__input")
	if iconAt < 0 || !strings.Contains(inner, "#tc-i-search") {
		t.Fatalf("search field must include the search icon: %s", inner)
	}
	if inputAt < 0 || iconAt > inputAt {
		t.Fatalf("icon must be the first node of the field: %s", inner)
	}
}

func TestDocMeta_emptyRendersNothing(t *testing.T) {
	if got := (widget.DocMeta{}).Render(); got != "" {
		t.Fatalf("empty DocMeta must render nothing, got %q", got)
	}
}

func TestSideBlock_exactlyOne(t *testing.T) {
	sec := widget.Sec{Rubric: widget.Rubric{Text: "Info"}, Body: []widget.Node{widget.Paragraph("Details")}}
	field := widget.Field{Name: "title", Label: "Title", Kind: widget.InputText, Value: "Note"}
	fields := widget.Fields{Rows: []widget.FieldRow{{Label: "UID", Value: "42"}}}

	if got := (widget.SideBlock{Sec: &sec}).Render(); !strings.Contains(got, "tc-sec") {
		t.Fatalf("sec block: %q", got)
	}
	if got := (widget.SideBlock{Field: &field}).Render(); !strings.Contains(got, "tc-field") {
		t.Fatalf("field block: %q", got)
	}
	pfield := widget.PField{Name: "title", Label: "Title", Kind: widget.InputText, Width: widget.PFieldWidthFull}
	pfieldRow := widget.PFieldRow{Fields: []widget.PField{{Name: "date", Label: "Date", Kind: widget.InputDate, Width: widget.PFieldWidthSM}}}
	if got := (widget.SideBlock{PField: &pfield}).Render(); !strings.Contains(got, "tc-pfield") {
		t.Fatalf("pfield block: %q", got)
	}
	if got := (widget.SideBlock{PFieldRow: &pfieldRow}).Render(); !strings.Contains(got, "tc-pfieldrow") {
		t.Fatalf("pfieldrow block: %q", got)
	}
	if got := (widget.SideBlock{Fields: &fields}).Render(); !strings.Contains(got, "tc-fields") {
		t.Fatalf("fields block: %q", got)
	}
	if got := (widget.SideBlock{}).Render(); got != "" {
		t.Fatalf("empty block must render nothing, got %q", got)
	}
	both := widget.SideBlock{Sec: &sec, Field: &field}
	if got := both.Render(); got != "" {
		t.Fatalf("two kinds must render nothing, got %q", got)
	}
}

func TestField_widthAndOwnLine(t *testing.T) {
	html := widget.Field{Name: "title", Label: "Title", Kind: widget.InputText, Width: widget.FieldWidthXL, OwnLine: true}.Render()
	if !strings.Contains(html, "tc-field--xl") {
		t.Fatalf("Width must print a field width class: %s", html)
	}
	if !strings.Contains(html, "is-own-line") {
		t.Fatalf("OwnLine must print is-own-line: %s", html)
	}
}

func TestFormFoot_emptyRendersNothing(t *testing.T) {
	if got := (widget.FormFoot{}).Render(); got != "" {
		t.Fatalf("empty FormFoot must render nothing, got %q", got)
	}
}

func TestField_emptyLabelIsBare(t *testing.T) {
	html := widget.Field{Name: "quick", Kind: widget.InputTextarea, Value: "note"}.Render()
	if strings.Contains(html, "tc-field__label") {
		t.Fatalf("unlabeled field must omit the label column: %s", html)
	}
	if !strings.Contains(html, "tc-field--bare") {
		t.Fatalf("unlabeled field must be bare: %s", html)
	}
}

func TestField_bareKeepsAccessibleName(t *testing.T) {
	html := widget.Field{Name: "quick", Label: "Note", Bare: true, Kind: widget.InputTextarea}.Render()
	if strings.Contains(html, "tc-field__label") {
		t.Fatalf("bare field must not print the label: %s", html)
	}
	if !strings.Contains(html, "tc-field--bare") {
		t.Fatalf("Bare must print tc-field--bare: %s", html)
	}
	if !strings.Contains(html, `aria-label="Note"`) {
		t.Fatalf("bare field must keep Label as the accessible name: %s", html)
	}
}

func TestField_labelIsInputName(t *testing.T) {
	html := widget.Field{Name: "display", Label: "Display name", Kind: widget.InputText}.Render()
	if strings.Contains(html, `aria-label="display"`) {
		t.Fatalf("machine Name must not override the visible label: %s", html)
	}
}

func TestPField_labelAboveInput(t *testing.T) {
	html := widget.PField{Name: "date", Label: "Date", Kind: widget.InputDate, Width: widget.PFieldWidthSM}.Render()
	labelAt := strings.Index(html, `class="tc-pfield__label"`)
	inputAt := strings.Index(html, `<input`)
	if labelAt < 0 || inputAt < 0 || labelAt > inputAt {
		t.Fatalf("label must precede input: %s", html)
	}
	if strings.Contains(html, "tc-field") {
		t.Fatalf("panel field must not use form grid: %s", html)
	}
}

func TestPFieldRow_emptyAndSelectWrap(t *testing.T) {
	if got := (widget.PFieldRow{}).Render(); got != "" {
		t.Fatalf("empty row must render nothing, got %q", got)
	}
	html := widget.PFieldRow{Fields: []widget.PField{
		{Name: "book", Label: "Notebook", Kind: widget.InputSelect, Options: []widget.SelectOption{{Label: "Notes", Value: "n", Selected: true}}, Width: widget.PFieldWidthLG},
	}}.Render()
	if strings.Contains(html, "tc-in--pick") || strings.Contains(html, "tc-in tc-in--pick-wrap") {
		t.Fatalf("select wrap class must not be rewritten: %s", html)
	}
	if !strings.Contains(html, "tc-in-wrap--select") {
		t.Fatalf("select wrap missing: %s", html)
	}
}

func TestPick_queryIsInput(t *testing.T) {
	closed := widget.Pick{Name: "rel", Label: "Related to", Chosen: []string{"Sync-up"}}.Render()
	if !strings.Contains(closed, `class="tc-pick__query"`) || !strings.Contains(closed, `<input type="text"`) {
		t.Fatalf("pick query must be an input: %s", closed)
	}
	if strings.Contains(closed, "aria-controls") {
		t.Fatalf("closed pick without menu must not point at a missing listbox: %s", closed)
	}
	open := widget.Pick{Name: "rel", Label: "Related to", Query: "kov", Open: true, Options: []widget.PickItem{{Label: "Call"}}, Manual: "Enter a UID by hand"}.Render()
	if !strings.Contains(open, `aria-expanded="true"`) || !strings.Contains(open, "tc-pick__menu") {
		t.Fatalf("open pick must expose the listbox: %s", open)
	}
}

func TestInput_pickKindEmitsNothing(t *testing.T) {
	if got := (widget.Input{Name: "rel", Kind: widget.InputPick}).Render(); got != "" {
		t.Fatalf("InputPick belongs on PField, not tc-in: %s", got)
	}
}

func TestDocEdit_titleAndBody(t *testing.T) {
	html := widget.DocEdit{Title: "Note", Text: "Hello"}.Render()
	if !strings.Contains(html, "tc-pfield--title") {
		t.Fatalf("title field missing: %s", html)
	}
	if !strings.Contains(html, "Hello") || !strings.Contains(html, "tc-in--textarea") {
		t.Fatalf("body textarea missing: %s", html)
	}
}

func TestField_hintAndErrorDescribeControl(t *testing.T) {
	cases := map[string]string{
		"field":  widget.Field{Name: "mail", Label: "Mail", Kind: widget.InputEmail, Hint: "Work address", Error: "Required"}.Render(),
		"check":  widget.Field{Name: "agree", Label: "Agree", Check: &widget.Check{}, Hint: "Once", Error: "Required"}.Render(),
		"pfield": widget.PField{Name: "date", Label: "Date", Kind: widget.InputDate, Hint: "Local time", Error: "Bad date"}.Render(),
		"pick":   widget.PField{Name: "rel", Label: "Related", Kind: widget.InputPick, Pick: &widget.Pick{}, Hint: "Any record", Error: "Gone"}.Render(),
	}
	idRe := regexp.MustCompile(`class="tc-p?field__(hint|error)" id="([^"]+)"`)
	for name, html := range cases {
		ids := idRe.FindAllStringSubmatch(html, -1)
		if len(ids) != 2 {
			t.Fatalf("%s: want hint and error ids, got %v in %s", name, ids, html)
		}
		want := `aria-describedby="` + ids[0][2] + " " + ids[1][2] + `"`
		if strings.Count(html, want) != 1 {
			t.Fatalf("%s: control must carry %s: %s", name, want, html)
		}
	}
}

func TestField_hintIDsUniquePerRender(t *testing.T) {
	a := widget.Field{Name: "label", Label: "Label", Hint: "h"}.Render()
	b := widget.Field{Name: "label", Label: "Label", Hint: "h"}.Render()
	idRe := regexp.MustCompile(`id="([^"]+-hint)"`)
	if idRe.FindStringSubmatch(a)[1] == idRe.FindStringSubmatch(b)[1] {
		t.Fatalf("two fields with one Name must not share a hint id")
	}
}
