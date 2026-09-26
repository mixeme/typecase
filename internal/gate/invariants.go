package gate

import (
	"fmt"
	"strings"
)

// invariantBindings maps each widget to executable checks for its invariant clauses
// (semicolon-separated in L-3.2 «Инварианты»). Order matches clause order.
var invariantBindings = map[string][]string{
	"Msg":         {"gate:enumerations", "gate:widget-fragments"},
	"Diag":        {"gate:widget-fragments", "gate:widget-fragments"},
	"Danger":      {"gate:widget-fragments", "gate:widget-fragments", "gate:widget-fragments"},
	"Meter":       {"gate:typed-input", "gate:widget-fragments"},
	"Runbar":      {"gate:widget-fragments", "gate:widget-fragments", "gate:widget-fragments"},
	"Queue":       {"gate:widget-fragments", "gate:widget-fragments", "gate:widget-fragments"},
	"Stat":        {"gate:typed-input", "gate:widget-fragments"},
	"Stats":       {"gate:css-from-mockup", "gate:widget-fragments"},
	"Steps":       {"gate:typed-input", "gate:widget-fragments", "gate:widget-fragments"},
	"Diff":        {"gate:container-pairs", "gate:widget-fragments", "gate:widget-fragments"},
	"Empty":       {"gate:widget-fragments", "gate:enumerations", "gate:widget-fragments"},
	"Thumb":       {"gate:enumerations", "gate:widget-fragments"},
	"Tile":        {"gate:css-from-mockup", "gate:widget-fragments"},
	"Grid":        {"gate:container-pairs", "gate:widget-fragments"},
	"Pad":         {"gate:css-from-mockup", "gate:widget-fragments"},
	"Fields":      {"gate:widget-fragments", "gate:enumerations", "gate:widget-fragments"},
	"Sec":         {"gate:widget-fragments", "gate:css-from-mockup", "gate:widget-fragments"},
	"Card":        {"gate:enumerations", "gate:container-pairs", "gate:widget-fragments"},
	"Table":       {"gate:enumerations", "gate:css-from-mockup", "gate:widget-fragments"},
	"Docmeta":     {"gate:widget-fragments", "gate:widget-fragments", "gate:widget-fragments"},
	"Prose":       {"gate:css-from-mockup", "gate:enumerations", "gate:widget-fragments"},
	"Doc":         {"gate:widget-fragments", "gate:container-pairs", "gate:widget-fragments"},
	"Docbar":      {"browser:bar-one-row", "gate:widget-fragments", "gate:widget-fragments"},
	"Menu":        {"gate:a11y", "gate:enumerations", "gate:widget-fragments"},
	"Drop":        {"gate:widget-fragments", "gate:widget-fragments", "gate:widget-fragments"},
	"Sheet":       {"gate:widget-fragments", "gate:widget-fragments", "gate:enumerations", "gate:widget-fragments"},
	"Scrim":       {"gate:container-pairs", "gate:enumerations", "gate:widget-fragments"},
	"Overlay":     {"gate:container-pairs", "gate:widget-fragments"},
	"Dialog":      {"gate:enumerations", "gate:widget-fragments", "gate:widget-fragments"},
	"Auth":        {"gate:container-pairs", "gate:widget-fragments"},
	"Page":        {"gate:container-pairs", "gate:enumerations", "gate:widget-fragments"},
	"Head":        {"gate:widget-fragments", "gate:widget-fragments", "gate:enumerations", "gate:widget-fragments"},
	"Bar":         {"browser:bar-one-row", "gate:widget-fragments", "gate:widget-fragments"},
	"Main":        {"gate:css-from-mockup", "gate:widget-fragments", "gate:widget-fragments"},
	"Side":        {"gate:css-from-mockup", "browser:esc-layers", "gate:widget-fragments"},
	"Shell":       {"gate:widget-fragments", "gate:css-from-mockup", "gate:widget-fragments"},
	"Brand":       {"gate:a11y", "gate:widget-fragments"},
	"Search":      {"gate:a11y", "gate:enumerations", "gate:widget-fragments"},
	"UserMenu":    {"gate:widget-fragments", "gate:widget-fragments"},
	"Emblem":      {"gate:widget-fragments", "gate:a11y", "gate:widget-fragments"},
	"Split":       {"gate:container-pairs", "gate:widget-fragments"},
	"Zone":        {"gate:container-pairs", "gate:widget-fragments"},
	"Crumbs":      {"gate:widget-fragments", "gate:widget-fragments"},
	"Acts":        {"gate:widget-fragments", "gate:widget-fragments"},
	"Nav":         {"gate:typed-input", "gate:a11y", "gate:widget-fragments"},
	"RailSec":     {"gate:widget-fragments", "gate:widget-fragments", "gate:widget-fragments"},
	"RailStatus":  {"gate:widget-fragments", "gate:widget-fragments", "gate:widget-fragments"},
	"RailFoot":    {"gate:widget-fragments", "gate:widget-fragments", "gate:widget-fragments"},
	"Top":         {"gate:widget-fragments", "gate:widget-fragments", "gate:widget-fragments"},
	"TickRow":     {"gate:a11y", "gate:enumerations", "gate:widget-fragments"},
	"Skel":        {"gate:enumerations", "gate:widget-fragments"},
	"Tree":        {"gate:enumerations", "gate:typed-input", "gate:widget-fragments"},
	"List":        {"gate:container-pairs", "gate:widget-fragments"},
	"TimeRow":     {"gate:container-pairs", "gate:widget-fragments", "gate:widget-fragments"},
	"SkelRow":     {"gate:container-pairs", "gate:widget-fragments", "gate:widget-fragments"},
	"FaceRow":     {"gate:container-pairs", "gate:widget-fragments", "gate:widget-fragments"},
	"FindRow":     {"gate:container-pairs", "gate:widget-fragments", "gate:widget-fragments"},
	"LinkedRow":   {"gate:container-pairs", "gate:widget-fragments"},
	"ExcerptRow":  {"gate:container-pairs", "gate:widget-fragments", "gate:widget-fragments"},
	"BriefRow":    {"gate:container-pairs", "gate:widget-fragments"},
	"DueRow":      {"gate:widget-fragments", "gate:container-pairs", "gate:widget-fragments"},
	"SizeRow":     {"gate:a11y", "gate:container-pairs", "gate:widget-fragments"},
	"SlimRow":     {"gate:a11y", "gate:widget-fragments"},
	"PhoneRow":    {"gate:widget-fragments", "gate:a11y", "gate:widget-fragments"},
	"tc-push":     {"gate:widget-fragments", "gate:widget-fragments"},
	"tc-grow":     {"gate:widget-fragments", "gate:widget-fragments"},
	"tc-stack":    {"gate:enumerations", "gate:widget-fragments"},
	"tc-line":     {"gate:enumerations", "gate:widget-fragments", "gate:widget-fragments"},
	"tc-q":        {"gate:widget-fragments", "gate:a11y", "gate:widget-fragments"},
	"Area":        {"gate:widget-fragments", "gate:widget-fragments", "gate:widget-fragments"},
	"TagInput":    {"gate:container-pairs", "gate:widget-fragments"},
	"Filter":      {"gate:widget-fragments", "gate:widget-fragments"},
	"Pick":        {"gate:widget-fragments", "gate:widget-fragments", "gate:widget-fragments"},
	"Input":       {"gate:widget-fragments", "gate:enumerations", "gate:a11y", "gate:widget-fragments"},
	"tc-plain":    {"gate:widget-fragments", "gate:widget-fragments"},
	"tc-mono":     {"gate:enumerations", "gate:widget-fragments"},
	"tc-dim":      {"gate:widget-fragments", "gate:enumerations", "gate:widget-fragments"},
	"tc-avatar":   {"gate:widget-fragments", "gate:enumerations", "gate:widget-fragments"},
	"tc-badge":    {"gate:inv-badge-has-text", "gate:inv-badge-not-link", "gate:widget-fragments"},
	"tc-btn":      {"gate:enumerations", "gate:a11y", "gate:widget-fragments"},
	"tc-check":    {"gate:enumerations", "gate:a11y", "gate:widget-fragments"},
	"tc-close":    {"gate:a11y", "gate:widget-fragments"},
	"tc-field":    {"gate:enumerations", "gate:a11y", "gate:widget-fragments"},
	"tc-form":     {"gate:widget-fragments", "gate:enumerations", "gate:widget-fragments"},
	"tc-formfoot": {"gate:widget-fragments", "gate:widget-fragments"},
	"tc-fset":     {"gate:container-pairs", "gate:widget-fragments", "gate:widget-fragments"},
	"tc-in":       {"gate:a11y", "gate:a11y"},
	"tc-kbd":      {"gate:inv-kbd-not-button", "gate:widget-fragments"},
	"tc-mark":     {"gate:no-inline", "gate:enumerations", "gate:widget-fragments"},
	"tc-num":      {"gate:typed-input", "gate:widget-fragments"},
	"tc-range":    {"gate:widget-fragments", "gate:widget-fragments"},
	"tc-rubric":   {"gate:widget-fragments", "gate:widget-fragments", "gate:widget-fragments"},
	"tc-seg":      {"gate:widget-fragments", "gate:a11y", "gate:widget-fragments"},
	"tc-tag":      {"gate:inv-tag-not-editable", "gate:widget-fragments"},
	"tc-tick":     {"gate:a11y", "gate:enumerations", "gate:widget-fragments"},
}

func splitInvariantClauses(text string) []string {
	text = strings.TrimSpace(text)
	low := strings.ToLower(text)
	if text == "" || low == "нет" || strings.HasPrefix(low, "нет ") {
		return nil
	}
	parts := strings.Split(text, ";")
	var out []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func checkRefValid(ref string) bool {
	switch {
	case strings.HasPrefix(ref, "gate:"):
		name := strings.TrimPrefix(ref, "gate:")
		if isRunnableInvGate(name) {
			return true
		}
		return gateImplemented(Name(name))
	case strings.HasPrefix(ref, "browser:"):
		name := strings.TrimPrefix(ref, "browser:")
		for _, c := range browserChecks {
			if c == name {
				return true
			}
		}
		return false
	case strings.HasPrefix(ref, "todo:"):
		return false
	case ref == "widget:smoke":
		return false
	default:
		return false
	}
}

// CheckInvariantCoverage ensures every invariant clause has a named executable check.
func CheckInvariantCoverage(widgetDir string, catalog []CatalogEntry) error {
	widgets, err := LoadWidgets(widgetDir)
	if err != nil {
		return err
	}
	var problems []string
	todos := 0
	for goType, spec := range widgets {
		name := spec.Name
		clauses := splitInvariantClauses(spec.Sections["Инварианты"])
		if len(clauses) == 0 {
			continue
		}
		// A binding names the Go type when a block has several (tc-in).
		checks, ok := invariantBindings[goType]
		if !ok {
			checks, ok = invariantBindings[name]
		}
		if !ok {
			problems = append(problems, name+": no invariant bindings")
			continue
		}
		if len(checks) != len(clauses) {
			problems = append(problems, fmt.Sprintf("%s: %d clauses, %d checks", name, len(clauses), len(checks)))
			continue
		}
		for i, ref := range checks {
			if strings.HasPrefix(ref, "todo:") {
				todos++
				continue
			}
			if !checkRefValid(ref) {
				problems = append(problems, fmt.Sprintf("%s[%d]: unknown check %q", name, i, ref))
				continue
			}
			if strings.HasPrefix(ref, "gate:") {
				gateName := strings.TrimPrefix(ref, "gate:")
				if isRunnableInvGate(gateName) {
					if err := checkInvariantGate(ref, name, catalog, widgetDir); err != nil {
						problems = append(problems, err.Error())
					}
				}
			}
		}
	}
	if todos > maxInvariantTodos {
		problems = append(problems, fmt.Sprintf("%d todos exceed ceiling %d", todos, maxInvariantTodos))
	}
	if len(problems) > 0 {
		return fmt.Errorf("invariant-coverage: %s", strings.Join(problems, "; "))
	}
	return nil
}
