package gate

import (
	"fmt"
	"strings"
)

// invariantBindings maps each widget to executable checks for its invariant clauses
// (semicolon-separated in L-3.2 «Инварианты»). Order matches clause order.
var invariantBindings = map[string][]string{
	"Head":         {"gate:widget-fragments", "gate:widget-fragments", "gate:enumerations", "gate:widget-fragments"},
	"Bar":          {"browser:bar-one-row", "gate:widget-fragments", "gate:widget-fragments"},
	"Main":         {"gate:css-from-mockup", "gate:widget-fragments", "gate:widget-fragments"},
	"Side":         {"gate:css-from-mockup", "browser:esc-layers", "gate:widget-fragments"},
	"Shell":        {"gate:widget-fragments", "gate:css-from-mockup", "gate:widget-fragments"},
	"Brand":        {"gate:a11y", "gate:widget-fragments"},
	"Search":       {"gate:a11y", "gate:enumerations", "gate:widget-fragments"},
	"UserMenu":     {"gate:widget-fragments", "gate:widget-fragments"},
	"Emblem":       {"gate:widget-fragments", "gate:a11y", "gate:widget-fragments"},
	"Split":        {"gate:container-pairs", "gate:widget-fragments"},
	"Zone":         {"gate:container-pairs", "gate:widget-fragments"},
	"Crumbs":       {"gate:widget-fragments", "gate:widget-fragments"},
	"Acts":         {"gate:widget-fragments", "gate:widget-fragments"},
	"Nav":          {"gate:typed-input", "gate:a11y", "gate:widget-fragments"},
	"RailSec":      {"gate:widget-fragments", "gate:widget-fragments", "gate:widget-fragments"},
	"RailStatus":   {"gate:widget-fragments", "gate:widget-fragments", "gate:widget-fragments"},
	"RailFoot":     {"gate:widget-fragments", "gate:widget-fragments", "gate:widget-fragments"},
	"Top":          {"gate:widget-fragments", "gate:widget-fragments", "gate:widget-fragments"},
	"TickRow":      {"gate:a11y", "gate:enumerations", "gate:widget-fragments"},
	"Skel":         {"gate:enumerations", "gate:widget-fragments"},
	"Tree":         {"gate:enumerations", "gate:typed-input", "gate:widget-fragments"},
	"List":         {"gate:container-pairs", "gate:widget-fragments"},
	"TimeRow":      {"gate:container-pairs", "gate:widget-fragments", "gate:widget-fragments"},
	"SkelRow":      {"gate:container-pairs", "gate:widget-fragments", "gate:widget-fragments"},
	"FaceRow":      {"gate:container-pairs", "gate:widget-fragments", "gate:widget-fragments"},
	"FindRow":      {"gate:container-pairs", "gate:widget-fragments", "gate:widget-fragments"},
	"LinkedRow":    {"gate:container-pairs", "gate:widget-fragments"},
	"ExcerptRow":   {"gate:container-pairs", "gate:widget-fragments", "gate:widget-fragments"},
	"BriefRow":     {"gate:container-pairs", "gate:widget-fragments"},
	"DueRow":       {"gate:widget-fragments", "gate:container-pairs", "gate:widget-fragments"},
	"SizeRow":      {"gate:a11y", "gate:container-pairs", "gate:widget-fragments"},
	"SlimRow":      {"gate:a11y", "gate:widget-fragments"},
	"PhoneRow":     {"gate:widget-fragments", "gate:a11y", "gate:widget-fragments"},
	"tc-push":      {"gate:widget-fragments", "gate:widget-fragments"},
	"tc-grow":      {"gate:widget-fragments", "gate:widget-fragments"},
	"tc-stack":     {"gate:enumerations", "gate:widget-fragments"},
	"tc-line":      {"gate:enumerations", "gate:widget-fragments", "gate:widget-fragments"},
	"tc-q":         {"gate:widget-fragments", "gate:a11y", "gate:widget-fragments"},
	"Area":         {"gate:widget-fragments", "gate:widget-fragments", "gate:widget-fragments"},
	"TagInput":     {"gate:container-pairs", "gate:widget-fragments"},
	"Filter":       {"gate:widget-fragments", "gate:widget-fragments"},
	"Pick":         {"gate:widget-fragments", "gate:widget-fragments", "gate:widget-fragments"},
	"Input":        {"gate:widget-fragments", "gate:enumerations", "gate:a11y", "gate:widget-fragments"},
	"tc-plain":     {"gate:widget-fragments", "gate:widget-fragments"},
	"tc-mono":      {"gate:enumerations", "gate:widget-fragments"},
	"tc-dim":       {"gate:widget-fragments", "gate:enumerations", "gate:widget-fragments"},
	"tc-avatar":    {"gate:widget-fragments", "gate:enumerations", "gate:widget-fragments"},
	"tc-badge":     {"gate:inv-badge-has-text", "gate:inv-badge-not-link", "gate:widget-fragments"},
	"tc-btn":       {"gate:enumerations", "gate:a11y", "gate:widget-fragments"},
	"tc-card":      {"gate:inv-card-has-body"},
	"tc-check":     {"gate:enumerations", "gate:a11y", "gate:widget-fragments"},
	"tc-chip":      {"gate:inv-chip-state-class"},
	"tc-cols":      {"gate:layout-slots"},
	"tc-close":     {"gate:a11y", "gate:widget-fragments"},
	"tc-crop":      {"gate:inv-crop-has-frame", "gate:inv-crop-has-zoom"},
	"tc-danger":    {"gate:a11y", "gate:a11y"},
	"tc-diag":      {"gate:inv-diag-collapsed-default", "gate:inv-diag-mono"},
	"tc-dialog":    {"browser:esc-layers", "gate:a11y"},
	"tc-diff":      {"gate:inv-diff-options-column", "gate:inv-diff-one-chosen"},
	"tc-doc":       {"gate:inv-doc-mode-switch", "gate:inv-doc-no-silent-save"},
	"tc-docedit":   {"gate:a11y"},
	"tc-docbar":    {"gate:inv-docbar-actions", "gate:inv-docbar-dirty-flag"},
	"tc-docmeta":   {"gate:inv-docmeta-sections"},
	"tc-empty":     {"gate:a11y"},
	"tc-field":     {"gate:enumerations", "gate:a11y", "gate:widget-fragments"},
	"tc-fields":    {"browser:ellipsis", "gate:inv-fields-dl-structure"},
	"tc-pfield":    {"gate:a11y", "gate:a11y"},
	"tc-pfieldrow": {"gate:a11y"},
	"tc-pick":      {"gate:a11y", "gate:a11y"},
	"tc-form":      {"gate:widget-fragments", "gate:enumerations", "gate:widget-fragments"},
	"tc-formfoot":  {"gate:widget-fragments", "gate:widget-fragments"},
	"tc-fset":      {"gate:container-pairs", "gate:widget-fragments", "gate:widget-fragments"},
	"tc-grid":      {"gate:inv-grid-tiles", "gate:inv-grid-uniform"},
	"tc-in":        {"gate:a11y", "gate:a11y"},
	"tc-kbd":       {"gate:inv-kbd-not-button", "gate:widget-fragments"},
	"tc-mark":      {"gate:no-inline", "gate:enumerations", "gate:widget-fragments"},
	"tc-menu":      {"browser:menu-label-arrow", "browser:popover-anchor", "browser:focus-return"},
	"tc-meter":     {"gate:inv-meter-has-value"},
	"tc-msg":       {"gate:a11y", "gate:a11y"},
	"tc-num":       {"gate:typed-input", "gate:widget-fragments"},
	"tc-overlay":   {"gate:inv-overlay-kind-class", "browser:popover-anchor"},
	"tc-prose":     {"gate:typed-input"},
	"tc-queue":     {"gate:inv-queue-question", "gate:inv-queue-meter"},
	"tc-rail":      {"gate:inv-rail-foot", "gate:inv-rail-drawer-class"},
	"tc-range":     {"gate:widget-fragments", "gate:widget-fragments"},
	"tc-rubric":    {"gate:widget-fragments", "gate:widget-fragments", "gate:widget-fragments"},
	"tc-runbar":    {"gate:inv-runbar-meter-hidden-idle", "gate:inv-runbar-states"},
	"tc-scrim":     {"gate:inv-scrim-for-attr", "browser:esc-layers"},
	"tc-sec":       {"gate:inv-sec-like-card"},
	"tc-seg":       {"gate:widget-fragments", "gate:a11y", "gate:widget-fragments"},
	"tc-sheet":     {"browser:esc-layers", "gate:a11y", "gate:a11y"},
	"tc-stat":      {"gate:inv-stat-string-value"},
	"tc-stats":     {"gate:inv-stats-no-wrap"},
	"tc-steps":     {"gate:inv-steps-one-current", "gate:inv-steps-count"},
	"tc-table":     {"gate:inv-table-column-kind", "gate:table-number-columns", "gate:inv-table-scrollable"},
	"tc-tag":       {"gate:inv-tag-not-editable", "gate:widget-fragments"},
	"tc-thumb":     {"gate:inv-thumb-placeholder", "gate:inv-thumb-icon-fallback"},
	"tc-tick":      {"gate:a11y", "gate:enumerations", "gate:widget-fragments"},
	"tc-tile":      {"gate:inv-tile-has-name"},
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
