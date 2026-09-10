package layout

import "strings"

// SlotClass is the element class for a named track slot.
func SlotClass(slot, gridSelector string) string {
	switch slot {
	case "mark":
		return "tc-mark"
	case "tick":
		if strings.Contains(gridSelector, "tickrow") {
			return "tc-tickrow__tick"
		}
		return "tc-tick"
	case "label":
		return "tc-tickrow__label"
	case "title":
		return "tc-row__title"
	case "excerpt":
		return "tc-row__excerpt"
	case "tags":
		return "tc-row__tags"
	case "date":
		return "tc-row__date"
	case "source":
		return "tc-row__source"
	case "avatar":
		return "tc-avatar"
	case "subtitle":
		return "tc-row__subtitle"
	case "detail":
		return "tc-row__detail"
	case "caption":
		return "tc-row__caption"
	case "name":
		return "tc-row__name"
	case "size":
		return "tc-row__size"
	case "time":
		return "tc-row__time"
	case "lead":
		return "tc-row__lead"
	case "due":
		return "tc-row__due"
	case "progress":
		return "tc-row__progress"
	default:
		return ""
	}
}

// HasSlot reports whether g names a slotted track.
func HasSlot(g Grid, slot string) bool {
	for _, t := range g.Tracks {
		if t.Slot == slot {
			return true
		}
	}
	return false
}

// Slotted reports whether g has named tracks.
func Slotted(g Grid) bool {
	for _, t := range g.Tracks {
		if t.Slot != "" {
			return true
		}
	}
	return false
}
