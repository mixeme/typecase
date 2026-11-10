package browser

import (
	"fmt"
	"strings"

	typecase "carrel.dev/typecase"
)

// FitJS is the part of tc.js that fits by measure (between its tc-fit:start
// and tc-fit:end marks): it folds every bar, turns every table its box
// cannot hold into cards and an open menu to the side where it has room — what a page with the script shows, on a frame
// painted without it. The frames draw their drops open and the rest of tc.js
// would close them; this takes the fit alone. It returns the number of bars
// and tables.
func FitJS() (string, error) {
	js, err := typecase.JS()
	if err != nil {
		return "", err
	}
	src := string(js)
	_, rest, ok := strings.Cut(src, "// tc-fit:start")
	if !ok {
		return "", fmt.Errorf("tc.js: no tc-fit:start mark")
	}
	body, _, ok := strings.Cut(rest, "// tc-fit:end")
	if !ok {
		return "", fmt.Errorf("tc.js: no tc-fit:end mark")
	}
	return `(() => {` + body + `
	var bars = document.querySelectorAll(".tc-bar, .tc-docbar");
	bars.forEach(function (b) { fitBar(b); });
	document.querySelectorAll(".tc-top").forEach(function (t) { fitBar(t, topSteps); });
	var tables = document.querySelectorAll("table.tc-table");
	tables.forEach(fitTable);
	document.querySelectorAll(".tc-menu.is-pop, .tc-drop__panel, .tc-drop__panel--right").forEach(fitMenu);
	return bars.length + tables.length;
})()`, nil
}
