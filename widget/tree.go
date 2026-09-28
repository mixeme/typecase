package widget

import "strings"

// TreeIndent is how deep a row of a tree stands.
type TreeIndent int

const (
	TreeLevel0  TreeIndent = iota
	TreeLevel20            // 1.25 rem: a folder under a folder in a narrow tree
	TreeLevel22            // one level down
	TreeLevel36            // two levels down in a picker of folders
	TreeLevel40            // two levels down
	TreeLevel44            // two levels down, under a row with an icon
)

func (i TreeIndent) class() string {
	return [...]string{"", "tc-tree__row--padl1-25rem", "tc-tree__row--padl22", "tc-tree__row--padl36",
		"tc-tree__row--padl40", "tc-tree__row--padl44"}[i]
}

// TreeTrail is what stands at the end of a row of a tree, pushed to it.
type TreeTrail interface {
	part
	treeTrail()
}

func (Dim) treeTrail()      {}
func (Mono) treeTrail()     {}
func (Badge) treeTrail()    {}
func (TreeNote) treeTrail() {}

// TreeNote is a dusty note at the end of a row of a tree: where a folder is,
// how many things are in it.
type TreeNote struct {
	Text string
}

func (n TreeNote) html(m mix) string {
	return `<span` + class(m, "tc-tickrow__meta", "tc-tickrow__meta--dusty") + `>` + esc(n.Text) + `</span>`
}

// TreeRow is a row of a tree: a folder, a source, a place to put a thing.
type TreeRow struct {
	Icon        Icon
	Mark        *Mark // a source's mark instead of the icon
	Label       string
	Code        bool // the label is a path or an address: monospaced
	Indent      TreeIndent
	Unavailable bool // cannot be chosen: read-only, already there
	Trail       TreeTrail
	Href        string // the node is a link: a folder opens at the address
}

func (r TreeRow) html(on bool) string {
	var s strings.Builder
	s.WriteString(`<` + rowTag(r.Href) + class(nil, is("on", on), "tc-tree__row", r.Indent.class(), pick(r.Unavailable, "tc-tree__row--op-55")) + attr("href", r.Href) + `>`)
	switch {
	case r.Mark != nil:
		s.WriteString(r.Mark.html(nil))
	case r.Icon != "":
		s.WriteString(`<svg>` + r.Icon.use() + `</svg>`)
	}
	if r.Code {
		s.WriteString(Mono{Text: r.Label}.html(nil))
	} else {
		s.WriteString(esc(r.Label))
	}
	if r.Trail != nil {
		s.WriteString(" " + r.Trail.html(mix{"tc-push"}))
	}
	s.WriteString(`</` + rowTag(r.Href) + `>`)
	return s.String()
}

// TreeVariant is the drawn form of a tree.
type TreeVariant int

const (
	TreePlain  TreeVariant = iota
	TreePadded             // 8 px at the sides: in the rail
	TreeNarrow             // at most 520 px: in a page
)

func (v TreeVariant) class() string {
	return [...]string{"", "tc-tree--padx8", "tc-tree--max520"}[v]
}

// Tree — tc-tree
//
// Роль. Дерево: папки, источники и их разделы, куда положить или откуда взять.
// Вход. Tree{Rows, On, Variant}.
// Состояния. plain | padded | narrow.
// Инварианты. Уровень строки и отступ дерева — перечисления; выбранная строка — одна, её номер, а не флаг у каждой строки; дерево печатает фрагмент своего кадра.
// Поведение. Нажатие на строку выбирает её; недоступную строку выбрать нельзя.
// Эталон. кадр `s-notes-11`, фрагмент `.tc-tree`.
// Содержимое. Строки: значок или метка источника, название (или путь моноширинным), прижатая к концу пометка — приглушённое, моноширинное, бейдж, примечание.
// Чего нет. Сворачивания ветвей: макет рисует дерево раскрытым.
type Tree struct {
	Rows    []TreeRow
	On      int // the chosen row, counted from 1; 0 for none
	Variant TreeVariant
}

// Render prints the tree.
func (t Tree) Render() string { return t.html(nil) }

func (t Tree) html(m mix) string {
	var s strings.Builder
	s.WriteString(`<div` + class(m, "tc-tree", t.Variant.class()) + `>`)
	for i, r := range t.Rows {
		s.WriteString(r.html(i+1 == t.On))
	}
	s.WriteString(`</div>`)
	return s.String()
}
