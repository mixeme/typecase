package css_test

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"testing"
)

// typoMetrics reads a font file's line metrics: OS/2's typo ascender,
// descender and line gap, in per cent of the em (head's unitsPerEm).
func typoMetrics(t *testing.T, path string) [3]float64 {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	tables := map[string][]byte{}
	n := int(binary.BigEndian.Uint16(b[4:]))
	for i := 0; i < n; i++ {
		r := b[12+16*i:]
		off, size := binary.BigEndian.Uint32(r[8:]), binary.BigEndian.Uint32(r[12:])
		tables[string(r[:4])] = b[off : off+size]
	}
	upm := float64(binary.BigEndian.Uint16(tables["head"][18:]))
	os2 := tables["OS/2"]
	pc := func(at int) float64 { return math.Abs(float64(int16(binary.BigEndian.Uint16(os2[at:])))) * 100 / upm }
	return [3]float64{pc(68), pc(70), pc(72)}
}

// TestFontMetricsPinned (Carrel's task 276 p. 5): every @font-face of the
// library and of the mockup pins the line's metrics to the font's own typo
// metrics — what Chromium lays text by — so a browser that took the win
// metrics (Firefox on Android) puts a button's text where every other
// browser does, beside its icon, not a pixel and a half under it.
func TestFontMetricsPinned(t *testing.T) {
	face := regexp.MustCompile(`(?s)@font-face\s*\{(.*?)\}`)
	prop := func(block, name string) (float64, bool) {
		m := regexp.MustCompile(name + `:\s*([0-9.]+)%`).FindStringSubmatch(block)
		if m == nil {
			return 0, false
		}
		v, _ := strconv.ParseFloat(m[1], 64)
		return v, true
	}
	file := regexp.MustCompile(`url\("[^"]*?([A-Za-z]+\.ttf)"\)`)
	for _, src := range []string{"fonts.css", filepath.Join("..", "docs", "visual", "carrel-ui-mockups.html")} {
		b, err := os.ReadFile(src)
		if err != nil {
			t.Fatal(err)
		}
		blocks := face.FindAllStringSubmatch(string(b), -1)
		if len(blocks) < 3 {
			t.Fatalf("%s: %d @font-face, want the three fonts", src, len(blocks))
		}
		for _, bl := range blocks {
			f := file.FindStringSubmatch(bl[1])
			if f == nil {
				t.Errorf("%s: an @font-face with no .ttf", src)
				continue
			}
			want := typoMetrics(t, filepath.Join("..", "fonts", f[1]))
			for i, name := range []string{"ascent-override", "descent-override", "line-gap-override"} {
				got, ok := prop(bl[1], name)
				if !ok || math.Abs(got-want[i]) > 0.01 {
					t.Errorf("%s, %s: %s %s; want %s", src, f[1], name, fmt.Sprint(got, ok), strconv.FormatFloat(want[i], 'f', -1, 64)+"%")
				}
			}
		}
	}
}
