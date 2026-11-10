package browser

import (
	"fmt"
	"strings"

	typecase "carrel.dev/typecase"
)

// FitBarsJS is the part of tc.js that folds the bars by measure (between its
// tc-fit:start and tc-fit:end marks), run on every bar of the page: what a
// page with the script shows, on a frame painted without it. The frames draw
// their drops open and the rest of tc.js would close them; this takes the
// bars' fold alone. It returns the number of bars folded.
func FitBarsJS() (string, error) {
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
	bars.forEach(fitBar);
	return bars.length;
})()`, nil
}
