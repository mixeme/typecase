package browser

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	mockupRel = "docs/visual/carrel-ui-mockups.html"

	// MockupShotViewport is the Chrome viewport width while shooting frames.
	MockupShotViewport = 1600
	// MockupShotMaxHeight is the visible slice of a frame (viewport height).
	MockupShotMaxHeight = 900
)

// LoadMockupHTML returns the mockup at git revision rev, or the working tree
// file when rev is empty.
func LoadMockupHTML(root, rev string) ([]byte, error) {
	rel := filepath.ToSlash(mockupRel)
	if rev == "" {
		return os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	}
	cmd := exec.Command("git", "-C", root, "show", rev+":"+rel)
	out, err := cmd.Output()
	if err != nil {
		stderr := ""
		if ee, ok := err.(*exec.ExitError); ok {
			stderr = strings.TrimSpace(string(ee.Stderr))
		}
		if ShallowRepo(root) {
			return nil, fmt.Errorf("git show %s:%s: shallow clone — the comparison needs the full history (fetch-depth: 0)", rev, rel)
		}
		if stderr != "" {
			return nil, fmt.Errorf("git show %s:%s: %v (%s)", rev, rel, err, stderr)
		}
		return nil, fmt.Errorf("git show %s:%s: %w", rev, rel, err)
	}
	return out, nil
}

// ShallowRepo reports whether root is a shallow git clone.
//
// The comparison reads the mockup at the base commit, which needs the full
// history (the runner asks for it with fetch-depth: 0). A shallow clone fails
// in git show with "bad object", which says nothing about the runner.
func ShallowRepo(root string) bool {
	out, err := exec.Command("git", "-C", root, "rev-parse", "--is-shallow-repository").Output()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(out)) == "true"
}

const mockupShotExtras = "body{margin:0;padding:0}\n" +
	"*{caret-color:transparent!important;animation:none!important;transition:none!important}\n" +
	"*:focus,*:focus-visible{outline:none!important}\n" +
	"input::-webkit-search-decoration,input::-webkit-search-cancel-button{display:none}\n" +
	"::-webkit-scrollbar{display:none}\n"

// MockupShotShell is a page with mockup CSS and sprite, and an empty host for frames.
// Style tags are flattened into one block so a split page/product stylesheet
// paints like the single <style> of the step 1 baseline.
func MockupShotShell(styles, sprite string) string {
	css := flattenStyleTags(styles)
	return "<!DOCTYPE html>\n<html lang=\"ru\"><head>\n<meta charset=\"utf-8\">\n" +
		"<style>\n" + css + "\n" +
		mockupShotExtras +
		"</style>\n</head><body>\n" +
		`<div aria-hidden="true" style="position:absolute;left:-9999px;top:0;font-size:16px">` +
		`<span style="font-family:'Public Sans'">AgЯя</span>` +
		`<span style="font-family:Literata">AgЯя</span>` +
		`<span style="font-family:'Source Code Pro'">AgЯя</span></div>` + "\n" +
		sprite + "\n<div id=\"host\"></div>\n</body></html>"
}

func flattenStyleTags(html string) string {
	var b strings.Builder
	for _, body := range styleTagBodies(html) {
		b.WriteString(body)
	}
	if b.Len() == 0 {
		return html
	}
	return b.String()
}

func styleTagBodies(html string) []string {
	low := strings.ToLower(html)
	var out []string
	i := 0
	for i < len(low) {
		start := strings.Index(low[i:], "<style")
		if start < 0 {
			break
		}
		start += i
		gt := strings.Index(html[start:], ">")
		if gt < 0 {
			break
		}
		body := start + gt + 1
		endRel := strings.Index(low[body:], "</style>")
		if endRel < 0 {
			break
		}
		out = append(out, html[body:body+endRel])
		i = body + endRel + len("</style>")
	}
	return out
}
