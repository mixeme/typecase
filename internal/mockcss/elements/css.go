package main

import "strings"

const controlMarker = "/* native controls: zero specificity */"

// controlCSS is the mockup copy of internal/proof/inline.go controlCSS,
// with .tcx → .tc-root. Zero specificity except .tc-q filling a drawn box.
const controlCSS = controlMarker + `
:where(.tc-root) :where(button,input,textarea){margin:0;padding:0;border:0;background:none;font:inherit;color:inherit;line-height:inherit;letter-spacing:inherit;text-align:inherit;text-transform:inherit;min-width:0;min-height:0}
:where(.tc-root) :where(button){cursor:pointer;appearance:none}
:where(.tc-root) :where(input[type=search]){appearance:none;-webkit-appearance:none}
:where(.tc-root) :where(input[type=search])::-webkit-search-decoration,
:where(.tc-root) :where(input[type=search])::-webkit-search-cancel-button,
:where(.tc-root) :where(input[type=search])::-webkit-search-results-button,
:where(.tc-root) :where(input[type=search])::-webkit-search-results-decoration{display:none;appearance:none}
:where(.tc-root) ::placeholder{color:inherit;opacity:1}
:where(.tc-root) :where(textarea){field-sizing:content;resize:none}
:where(.tc-root) :where(button[role=menuitem]){width:100%}
.tc-root .tc-q{flex:1;min-width:0;min-height:0;overflow:hidden}
.tc-acts > .tc-drop{flex:none}
`

const dropOffsetCSS = `.tc-drop:not(.is-bar) > .tc-menu.is-pop{margin-top:6px;font-weight:400;text-align:left}`

func injectCSS(html string) string {
	const needle = "/* элементы управления */"
	if i := strings.Index(html, controlMarker); i >= 0 {
		if j := strings.Index(html[i:], needle); j >= 0 {
			html = html[:i] + controlCSS + html[i+j:]
			return ensureDropOffset(html)
		}
		return ensureDropOffset(html)
	}
	i := strings.Index(html, needle)
	if i < 0 {
		return html
	}
	html = html[:i] + controlCSS + html[i:]
	return ensureDropOffset(html)
}

func ensureDropOffset(html string) string {
	if strings.Contains(html, ".tc-drop:not(.is-bar) > .tc-menu.is-pop") {
		return html
	}
	const nested = `.tc-btn > .tc-menu.is-pop{margin-top:6px;font-weight:400;text-align:left}`
	if i := strings.Index(html, nested); i >= 0 {
		at := i + len(nested)
		return html[:at] + "\n" + dropOffsetCSS + html[at:]
	}
	return html
}
