package layout

// EmptyPolicy is what Render does when a slot has no content.
type EmptyPolicy uint8

const (
	EmptyHold EmptyPolicy = iota // узел-заглушка, дорожка живёт
	EmptyOmit                    // только в Flow; в Grid — падение ворот
)

// Track is one grid column.
type Track struct {
	Slot  string // ключ слота; пусто у бесслотовых сеток
	Size  string // фрагмент grid-template без пробела после запятой: "minmax(140px,1.2fr)"
	Empty EmptyPolicy
}

// Breakpoint is a named-container restatement of tracks.
type Breakpoint struct {
	MaxPx    int
	Tracks   []Track
	Hide     []string // Slot, только новые на этом пороге
	PadLeft  string
	PadRight string
}

// Grid is a CSS grid with named tracks.
type Grid struct {
	Selector    string
	Gap         string
	MinHeight   string
	Align       string // "" или "center"
	PadLeft     string
	PadRight    string
	Tracks      []Track
	AutoFlow    string // "" или "column"
	AutoColumns string // "minmax(0,auto)" если AutoFlow задан, иначе ""
	Contents    string // селектор ребёнка с display:contents; "" если нет
	Narrow      []Breakpoint
	Container   string // "tcbox" или ""
}

// Box is an arbitrary property bag on one selector.
type Box struct {
	Selector string
	Props    map[string]string // ровно эти свойства, печатаются по отсортированным ключам
}

// Flow is flex (or other non-grid) layout on one selector.
type Flow struct {
	Selector string
	Props    map[string]string // display/flex-wrap/gap/padding/align-items/flex — не grid-template-*
}

// Spec is the registry of layout tables.
type Spec struct {
	Grids []Grid
	Boxes []Box
	Flows []Flow
}
