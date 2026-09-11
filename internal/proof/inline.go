package proof

import (
	"regexp"
	"strings"

	"carrel.dev/typecase/internal/mockcss"
)

// inlineRule is a style="" of a frame that became a widget variant. From is the
// attribute exactly as the frame has it; CSS is the part the variant carries,
// copied from From. The test checks both.
type inlineRule struct {
	Selector string
	CSS      string
	From     string
	Frame    string
}

// frameInline is the whole translation layer left: what frames decide in
// style="" instead of the <style> block. Selectors start at .tcx so they come
// after and above the generated rules, as an inline style would.
var frameInline = []inlineRule{
	{".tcx .tcx-rail-tags", "display:flex;flex-wrap:wrap;gap:4px;padding:0 16px", "display:flex;flex-wrap:wrap;gap:4px;padding:0 16px", "notes"},
	{".tcx .tcx-shell.is-side-440", "--side:440px", "--side:440px", "notes"},
	{".tcx .tcx-h1 > .tcx-bar3", "height:20px", "height:20px", "notes"},
	{".tcx .tcx-in.is-w200", "width:200px;height:26px", "width:200px;height:26px", "notes"},
	{".tcx .tcx-badge-ico", "display:inline;vertical-align:-1px", "display:inline;vertical-align:-1px", "notes"},
	{".tcx .tcx-h1.is-panel", "font-size:19px", "font-size:19px", "notes"},
	{".tcx .tcx-in.is-tall", "min-height:190px", "min-height:190px", "notes"},
	{".tcx .tcx-attach", "display:flex;flex-direction:column;gap:6px;font-size:12.5px", "display:flex;flex-direction:column;gap:6px;font-size:12.5px", "notes"},
	{".tcx .tcx-dim", "color:var(--dusty)", "color:var(--dusty)", "notes"},
	{".tcx .tcx-hint.is-spaced", "margin-top:6px", "margin-top:6px", "notes"},
	{".tcx .tcx-related", "font-size:12.5px", "font-size:12.5px", "notes"},
	{".tcx .tcx-ref", "color:var(--accent)", "color:var(--accent)", "notes"},
	{".tcx .tcx-shell.is-side-340", "--side:340px", "--side:340px", "files"},
	{".tcx .tcx-btn.is-pressed", "background:var(--card)", "background:var(--card)", "files"},
	{".tcx .tcx-strong", "font-weight:600", "font-weight:600", "files"},
	{".tcx .tcx-muted", "color:var(--muted)", "color:var(--muted)", "files"},
	{".tcx .tcx-in.is-w190", "width:190px;height:26px", "width:190px;height:26px", "files"},
	{".tcx .tcx-in-ico", "margin-right:6px", "margin-right:6px", "files"},
	{".tcx .tcx-col-tick", "width:28px", "width:28px", "files"},
	{".tcx .tcx-col-size", "width:92px", "width:92px", "files"},
	{".tcx .tcx-col-type", "width:132px", "width:132px", "files"},
	{".tcx .tcx-col-time", "width:150px", "width:150px", "files"},
	{".tcx .tcx-col-acts", "width:74px", "width:74px", "files"},
	{".tcx .tcx-rail-tree", "padding:0 8px", "padding:0 8px", "files"},
	{".tcx .tcx-aside", "margin-left:auto;color:var(--dusty)", "margin-left:auto;color:var(--dusty)", "files"},
	{".tcx .tcx-head.is-snug", "padding-bottom:8px", "padding-bottom:8px", "files"},
	{".tcx .tcx-h1.is-item", "font-size:16px", "font-size:16px", "files"},
	{".tcx .tcx-preview", "padding:0 20px", "padding:0 20px", "files"},
	{".tcx .tcx-thumb.is-wide", "aspect-ratio:16/10;border-radius:4px", "aspect-ratio:16/10;border-radius:4px", "files"},
	{".tcx .tcx-side-acts", "display:flex;gap:6px;padding:12px 20px 0;flex-wrap:wrap", "display:flex;gap:6px;padding:12px 20px 0;flex-wrap:wrap", "files"},
	{".tcx .tcx-path", "font-size:11px", "font-size:11px", "files"},
	{".tcx .tcx-note", "font-size:12px;color:var(--muted)", "font-size:12px;color:var(--muted)", "files"},

	{".tcx .tcx-h1.is-mark-19 > .tcx-bar3", "height:19px", "height:19px", "dark"},
	{".tcx .tcx-alert-text", "color:var(--alert)", "color:var(--alert)", "dark"},

	{".tcx .tcx-sheet-title", "font-size:13.5px", "font-size:13.5px", "quick"},
	{".tcx .tcx-small-dim", "font-size:11.5px;color:var(--dusty)", "font-size:11.5px;color:var(--dusty)", "quick"},
	{".tcx .tcx-foot-end", "margin-left:auto;display:flex;align-items:center;gap:7px", "margin-left:auto;display:flex;align-items:center;gap:7px", "quick"},
	{".tcx .tcx-btn.is-focus-ring", "outline:1.5px solid var(--accent);outline-offset:1px", "outline:1.5px solid var(--accent);outline-offset:1px", "quick"},
	{".tcx .tcx-in.is-note", "min-height:132px;font-size:14px;line-height:1.6", "min-height:132px;font-size:14px;line-height:1.6", "quick"},
	{".tcx .tcx-pick-head", "width:200px;font-size:12px;margin-left:auto", "width:200px;font-size:12px;margin-left:auto", "quick"},
	{".tcx .tcx-tint-accent", "stroke:var(--accent)", `stroke="var(--accent)"`, "quick"},

	{".tcx .tcx-tint-alert", "stroke:var(--alert)", `stroke="var(--alert)"`, "delete"},
	{".tcx .tcx-foot-end", "margin-left:auto;display:flex;gap:7px;align-items:center", "margin-left:auto;display:flex;gap:7px;align-items:center", "delete"},
	{".tcx .tcx-impacts", "display:flex;flex-direction:column;gap:6px;font-size:12.5px", "display:flex;flex-direction:column;gap:6px;font-size:12.5px", "delete"},
	{".tcx .tcx-line", "display:flex;gap:8px;align-items:center;flex-wrap:wrap", "display:flex;gap:8px;align-items:center;flex-wrap:wrap", "delete"},
	{".tcx .tcx-pad", "padding:0 20px", "padding:0 20px", "delete"},
	{".tcx .tcx-w100", "width:100%", "width:100%", "delete"},
	{".tcx .tcx-msg.is-flush", "margin:0", "margin:0", "delete"},
	{".tcx .tcx-rubric-6", "margin-bottom:6px", "margin-bottom:6px", "delete"},
	{".tcx .tcx-scrim.is-top-26", "padding-top:26px", "padding-top:26px", "delete"},
	{".tcx .tcx-sheet-560", "width:560px", "width:560px", "delete"},
	{".tcx .tcx-shell.is-tall-520", "min-height:520px", "min-height:520px", "delete"},
	{".tcx .tcx-url", "font-size:11.5px;color:var(--dusty)", "font-size:11.5px;color:var(--dusty)", "delete"},
	{".tcx .tcx-card-acts", "display:flex;gap:6px", "display:flex;gap:6px", "delete"},
	{".tcx .tcx-sheet-meta", "margin-left:auto;font-size:11.5px;color:var(--dusty)", "margin-left:auto;font-size:11.5px;color:var(--dusty)", "delete"},
	{".tcx .tcx-btn.is-narrow-pad", "padding:0 3px", "padding:0 3px", "delete"},
	{".tcx .tcx-mono-11", "font-size:11px", "font-size:11px", "delete"},
	{".tcx .tcx-col-tick-sm", "width:20px", "width:20px", "delete"},
	{".tcx .tcx-col-count", "width:70px", "width:70px", "delete"},
	{".tcx .tcx-col-menu", "width:26px", "width:26px", "delete"},

	{".tcx .tcx-col-changed", "width:140px", "width:140px", "conn"},
	{".tcx .tcx-danger-item", "color:var(--alert)", "color:var(--alert)", "conn"},
	{".tcx .tcx-card-note", "font-size:12.5px;color:var(--muted);margin-top:9px", "font-size:12.5px;color:var(--muted);margin-top:9px", "conn"},
	{".tcx .tcx-stack", "padding:14px 20px;display:flex;flex-direction:column;gap:10px", "padding:14px 20px;display:flex;flex-direction:column;gap:10px", "conn"},
	{".tcx .tcx-card-alert", "border-color:var(--alert)", "border-color:var(--alert)", "conn"},
	{".tcx .tcx-menu-222", "width:222px", "width:222px", "conn"},
	{".tcx .tcx-list-rubric", "display:flex;justify-content:space-between;align-items:baseline;margin:2px 0 -2px", "display:flex;justify-content:space-between;align-items:baseline;margin:2px 0 -2px", "conn"},
	{".tcx .tcx-url-gap", "margin-bottom:9px", "font-size:11.5px;color:var(--dusty);margin-bottom:9px", "conn"},
	{".tcx .tcx-ls0", "letter-spacing:0", "letter-spacing:0", "conn"},
}

// controlCSS is for real controls where the frame draws text: buttons, inputs,
// text areas. Not in the mockup because the mockup has none. Zero specificity
// except the one rule that lets a text input fill a drawn box.
const controlCSS = `:where(.tcx) :where(button,input,textarea){margin:0;padding:0;border:0;background:none;font:inherit;color:inherit;line-height:inherit;letter-spacing:inherit;text-align:inherit;text-transform:inherit}
:where(.tcx) :where(button){cursor:pointer}
:where(.tcx) :where(input[type=search]){appearance:none}
:where(.tcx) ::placeholder{color:inherit;opacity:1}
:where(.tcx) :where(textarea){field-sizing:content;resize:none}
:where(.tcx) :where(button[role=menuitem]){width:100%}
.tcx .tcx-q{flex:1;min-width:0}
`

func inlineCSS(o mockcss.Options) string {
	var b strings.Builder
	for _, r := range frameInline {
		b.WriteString(r.Selector + "{" + mockcss.RenameVars(r.CSS, o) + "}\n")
	}
	return b.String()
}

var inlineClassRe = regexp.MustCompile(`\.((?:tcx-|is-)[\w-]+)`)

// inlineClasses are the classes the inline table adds; the comparison ignores
// them because the frame has a style="" there instead.
func inlineClasses() []string {
	seen := map[string]bool{}
	var out []string
	for _, r := range frameInline {
		for _, m := range inlineClassRe.FindAllStringSubmatch(r.Selector, -1) {
			c := m[1]
			if strings.HasPrefix(c, "tcx-") && knownFromMockup[c] {
				continue
			}
			if !seen[c] {
				seen[c] = true
				out = append(out, c)
			}
		}
	}
	return out
}

// knownFromMockup are tcx- classes that are renamed mockup classes, not
// inline ones, where they appear in an inline selector as context.
var knownFromMockup = map[string]bool{
	"tcx-shell": true, "tcx-h1": true, "tcx-bar3": true, "tcx-in": true, "tcx-hint": true,
	"tcx-btn": true, "tcx-head": true, "tcx-thumb": true, "tcx-scrim": true, "tcx-msg": true,
}
