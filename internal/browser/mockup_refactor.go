package browser

import (
	"bytes"
	"fmt"
	"image/png"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	mockupRel       = "docs/visual/carrel-ui-mockups.html"
	allowedDiffsRel = "docs/visual/allowed-diffs.md"

	// MockupChannelTol is the per-channel 8-bit allowance for antialiasing.
	MockupChannelTol = 2
	// MockupAAPixelCap is how many PNG pixels may still differ after the
	// per-channel check: at MockupShotScale 0.25 a one-pixel input-vs-span
	// ClearType fringe becomes 2–4 pixels. A missing 11 px tick is ~8.
	MockupAAPixelCap = 4
	// MockupShotViewport is the Chrome viewport width while shooting frames.
	MockupShotViewport = 1600
	// MockupShotMaxHeight is the visible slice of a frame (viewport height).
	MockupShotMaxHeight = 900
	// MockupShotScale shrinks the PNG so 326 frames finish in one test run.
	MockupShotScale = 0.25
)

// MockupRefactorWidths are the two before/after screenshot widths.
var MockupRefactorWidths = []int{1351, 820}

// AllowedDiff skips one (frame, width) pair in the mockup refactor comparison.
type AllowedDiff struct {
	Frame    string
	Width    int
	Decision string
	Note     string
}

// AllowedDiffs is the parsed docs/visual/allowed-diffs.md file.
type AllowedDiffs struct {
	Base string
	Rows []AllowedDiff
}

// Skip reports whether this frame at this width is listed as an allowed change.
func (a AllowedDiffs) Skip(frame string, width int) bool {
	for _, r := range a.Rows {
		if r.Frame == frame && r.Width == width {
			return true
		}
	}
	return false
}

// ParseAllowedDiffs reads the allowed-diffs markdown.
func ParseAllowedDiffs(md string) (AllowedDiffs, error) {
	var out AllowedDiffs
	for _, line := range strings.Split(md, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "База:") {
			out.Base = strings.TrimSpace(strings.TrimPrefix(line, "База:"))
			if out.Base == "_" {
				out.Base = ""
			}
			continue
		}
		if !strings.HasPrefix(line, "|") {
			continue
		}
		cells := splitMDRow(line)
		if len(cells) < 2 {
			continue
		}
		if cells[0] == "Кадр" || strings.HasPrefix(cells[0], "-") {
			continue
		}
		w, err := strconv.Atoi(cells[1])
		if err != nil {
			return AllowedDiffs{}, fmt.Errorf("allowed-diffs: width %q", cells[1])
		}
		row := AllowedDiff{Frame: cells[0], Width: w}
		if len(cells) > 2 {
			row.Decision = cells[2]
		}
		if len(cells) > 3 {
			row.Note = cells[3]
		}
		out.Rows = append(out.Rows, row)
	}
	return out, nil
}

func splitMDRow(line string) []string {
	line = strings.Trim(line, "|")
	parts := strings.Split(line, "|")
	for i, p := range parts {
		parts[i] = strings.TrimSpace(p)
	}
	return parts
}

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
		if stderr != "" {
			return nil, fmt.Errorf("git show %s:%s: %v (%s)", rev, rel, err, stderr)
		}
		return nil, fmt.Errorf("git show %s:%s: %w", rev, rel, err)
	}
	return out, nil
}

// LoadAllowedDiffs reads docs/visual/allowed-diffs.md from module root.
func LoadAllowedDiffs(root string) (AllowedDiffs, error) {
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(allowedDiffsRel)))
	if err != nil {
		if os.IsNotExist(err) {
			return AllowedDiffs{}, nil
		}
		return AllowedDiffs{}, err
	}
	return ParseAllowedDiffs(string(data))
}

// PNGsMatch reports whether two PNGs are within per-channel 8-bit tolerance.
func PNGsMatch(a, b []byte, tol int) error {
	ia, err := png.Decode(bytes.NewReader(a))
	if err != nil {
		return fmt.Errorf("decode before: %w", err)
	}
	ib, err := png.Decode(bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("decode after: %w", err)
	}
	ba, bb := ia.Bounds(), ib.Bounds()
	if ba.Dx() != bb.Dx() || ba.Dy() != bb.Dy() {
		return fmt.Errorf("size %dx%d vs %dx%d", ba.Dx(), ba.Dy(), bb.Dx(), bb.Dy())
	}
	nDiff := 0
	first := ""
	minX, minY, maxX, maxY := ba.Dx(), ba.Dy(), -1, -1
	for y := 0; y < ba.Dy(); y++ {
		for x := 0; x < ba.Dx(); x++ {
			r1, g1, b1, a1 := ia.At(ba.Min.X+x, ba.Min.Y+y).RGBA()
			r2, g2, b2, a2 := ib.At(bb.Min.X+x, bb.Min.Y+y).RGBA()
			if channelDelta(r1, r2) > tol || channelDelta(g1, g2) > tol ||
				channelDelta(b1, b2) > tol || channelDelta(a1, a2) > tol {
				nDiff++
				if first == "" {
					first = fmt.Sprintf("%d,%d #%02x%02x%02x vs #%02x%02x%02x",
						x, y, r1>>8, g1>>8, b1>>8, r2>>8, g2>>8, b2>>8)
				}
				if x < minX {
					minX = x
				}
				if y < minY {
					minY = y
				}
				if x > maxX {
					maxX = x
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}
	if nDiff > 0 {
		return fmt.Errorf("%d pixels (first %s, box %d,%d–%d,%d)", nDiff, first, minX, minY, maxX, maxY)
	}
	return nil
}

func channelDelta(a, b uint32) int {
	d := int(a>>8) - int(b>>8)
	if d < 0 {
		return -d
	}
	return d
}

// MockupShotCSS concatenates every <style> body from a mockup document.
func MockupShotCSS(styles string) string {
	return flattenStyleTags(styles)
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

// MockupShotShellPair puts current and baseline CSS in one document.
// Baseline uses media="not all" until the test flips it — one Navigate,
// so Edge does not re-rasterize cream type. @font-face stays on the
// current sheet only; the files are the same and a second parse hangs Edge.
func MockupShotShellPair(afterStyles, beforeStyles, sprite string) string {
	after := flattenStyleTags(afterStyles)
	before := StripFontFaceCSS(flattenStyleTags(beforeStyles))
	return mockupShotDoc(after, before, sprite)
}

func mockupShotDoc(afterCSS, beforeCSS, sprite string) string {
	var b strings.Builder
	b.WriteString("<!DOCTYPE html>\n<html lang=\"ru\"><head>\n<meta charset=\"utf-8\">\n")
	b.WriteString(`<style id="tc-css-after">`)
	b.WriteString("\n")
	b.WriteString(afterCSS)
	b.WriteString("\n</style>\n<style>\n")
	b.WriteString(mockupShotExtras)
	b.WriteString("</style>\n")
	if beforeCSS != "" {
		b.WriteString(`<style id="tc-css-before" media="not all">`)
		b.WriteString("\n")
		b.WriteString(beforeCSS)
		b.WriteString("\n</style>\n")
	}
	b.WriteString("</head><body>\n")
	b.WriteString(`<div aria-hidden="true" style="position:absolute;left:-9999px;top:0;font-size:16px">`)
	b.WriteString(`<span style="font-family:'Public Sans'">AgЯя</span>`)
	b.WriteString(`<span style="font-family:Literata">AgЯя</span>`)
	b.WriteString(`<span style="font-family:'Source Code Pro'">AgЯя</span></div>`)
	b.WriteString("\n<div id=\"tc-sprite\">")
	b.WriteString(sprite)
	b.WriteString("</div>\n<div id=\"host\"></div>\n</body></html>")
	return b.String()
}

// StripFontFaceCSS drops @font-face rules so a same-tab CSS toggle
// does not re-parse the mockup font files.
func StripFontFaceCSS(css string) string {
	low := strings.ToLower(css)
	var b strings.Builder
	i := 0
	for i < len(css) {
		j := strings.Index(low[i:], "@font-face")
		if j < 0 {
			b.WriteString(css[i:])
			break
		}
		b.WriteString(css[i : i+j])
		k := i + j + len("@font-face")
		for k < len(css) && (css[k] == ' ' || css[k] == '\t' || css[k] == '\n' || css[k] == '\r') {
			k++
		}
		if k >= len(css) || css[k] != '{' {
			i = i + j + len("@font-face")
			continue
		}
		depth := 0
		for k < len(css) {
			switch css[k] {
			case '{':
				depth++
			case '}':
				depth--
				if depth == 0 {
					k++
					depth = -1
				}
			}
			if depth < 0 {
				break
			}
			k++
		}
		i = k
	}
	return b.String()
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
