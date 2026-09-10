package layout

import (
	"fmt"
	"sort"
	"strings"
)

// CSS prints All() as stylesheet text.
func CSS() string {
	return Format(All())
}

// Format prints spec as stylesheet text. Box/Flow keys are sorted.
func Format(spec Spec) string {
	var b strings.Builder
	for _, g := range spec.Grids {
		if len(g.Narrow) > 0 && g.Container == "" {
			panic("layout: unnamed @container for " + g.Selector)
		}
		writeGrid(&b, g)
	}
	for _, box := range spec.Boxes {
		writeProps(&b, box.Selector, box.Props)
	}
	for _, flow := range spec.Flows {
		writeProps(&b, flow.Selector, flow.Props)
	}
	return b.String()
}

func writeGrid(b *strings.Builder, g Grid) {
	var props []string
	props = append(props, "display:grid")
	if cols := TrackSizes(g.Tracks); cols != "" {
		props = append(props, "grid-template-columns:"+cols)
	}
	if g.AutoFlow != "" {
		props = append(props, "grid-auto-flow:"+g.AutoFlow)
	}
	if g.AutoColumns != "" {
		props = append(props, "grid-auto-columns:"+g.AutoColumns)
	}
	if g.Gap != "" {
		props = append(props, "gap:"+g.Gap)
	}
	if g.MinHeight != "" {
		props = append(props, "min-height:"+g.MinHeight)
	}
	if g.Align != "" {
		props = append(props, "align-items:"+g.Align)
	}
	if g.PadLeft != "" {
		props = append(props, "padding-left:"+g.PadLeft)
	}
	if g.PadRight != "" {
		props = append(props, "padding-right:"+g.PadRight)
	}
	b.WriteString(g.Selector)
	b.WriteByte('{')
	b.WriteString(strings.Join(props, ";"))
	b.WriteString("}\n")
	if g.Contents != "" {
		b.WriteString(g.Contents)
		b.WriteString("{display:contents}\n")
	}
	for _, bp := range g.Narrow {
		fmt.Fprintf(b, "@container %s (max-width:%dpx){\n", g.Container, bp.MaxPx)
		var nprops []string
		if cols := TrackSizes(bp.Tracks); cols != "" {
			nprops = append(nprops, "grid-template-columns:"+cols)
		}
		if bp.PadLeft != "" {
			nprops = append(nprops, "padding-left:"+bp.PadLeft)
		}
		if bp.PadRight != "" {
			nprops = append(nprops, "padding-right:"+bp.PadRight)
		}
		if len(nprops) > 0 {
			b.WriteString(g.Selector)
			b.WriteByte('{')
			b.WriteString(strings.Join(nprops, ";"))
			b.WriteString("}\n")
		}
		for _, slot := range bp.Hide {
			cls := SlotClass(slot, g.Selector)
			if cls == "" {
				panic("layout: Hide unknown slot " + slot)
			}
			fmt.Fprintf(b, "%s .%s{display:none}\n", g.Selector, cls)
		}
		b.WriteString("}\n")
	}
}

func writeProps(b *strings.Builder, sel string, props map[string]string) {
	keys := make([]string, 0, len(props))
	for k := range props {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	b.WriteString(sel)
	b.WriteByte('{')
	for i, k := range keys {
		if i > 0 {
			b.WriteByte(';')
		}
		b.WriteString(k)
		b.WriteByte(':')
		b.WriteString(props[k])
	}
	b.WriteString("}\n")
}

// TrackSizes joins track Size fragments with spaces.
func TrackSizes(tracks []Track) string {
	parts := make([]string, 0, len(tracks))
	for _, t := range tracks {
		if t.Size != "" {
			parts = append(parts, t.Size)
		}
	}
	return strings.Join(parts, " ")
}

// SplitTracks splits a grid-template-columns value on spaces outside parentheses.
func SplitTracks(value string) []string {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil
	}
	var parts []string
	depth := 0
	start := 0
	for i := 0; i < len(value); i++ {
		switch value[i] {
		case '(':
			depth++
		case ')':
			depth--
		case ' ', '\t', '\n':
			if depth != 0 {
				continue
			}
			if i > start {
				parts = append(parts, value[start:i])
			}
			for i+1 < len(value) && (value[i+1] == ' ' || value[i+1] == '\t' || value[i+1] == '\n') {
				i++
			}
			start = i + 1
		}
	}
	if start < len(value) {
		parts = append(parts, value[start:])
	}
	return parts
}
