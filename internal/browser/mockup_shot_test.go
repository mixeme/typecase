package browser_test

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"math"
	"strings"
	"testing"

	"carrel.dev/typecase/internal/browser"
)

func TestMockupShotShellFlattensStyles(t *testing.T) {
	got := browser.MockupShotShell(`<style>a{color:red}</style><style data-tc>b{color:blue}</style>`, "")
	if n := strings.Count(strings.ToLower(got), "<style"); n != 1 {
		t.Fatalf("want one <style>, got %d in %s", n, got)
	}
	if !strings.Contains(got, "a{color:red}") || !strings.Contains(got, "b{color:blue}") {
		t.Fatalf("missing concatenated rules: %s", got)
	}
}

// cropPNG cuts a frame's box out of a screenshot of the page.
func cropPNG(src []byte, x, y, w, h, scale float64) ([]byte, error) {
	im, err := png.Decode(bytes.NewReader(src))
	if err != nil {
		return nil, err
	}
	b := im.Bounds()
	r := image.Rect(
		int(math.Round(x*scale)),
		int(math.Round(y*scale)),
		int(math.Round((x+w)*scale)),
		int(math.Round((y+h)*scale)),
	).Intersect(b)
	if r.Empty() {
		return nil, fmt.Errorf("crop empty")
	}
	dst := image.NewNRGBA(image.Rect(0, 0, r.Dx(), r.Dy()))
	draw.Draw(dst, dst.Bounds(), im, r.Min, draw.Src)
	var buf bytes.Buffer
	if err := png.Encode(&buf, dst); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

const frameBoxJS = `(function(){
  var f = document.querySelector(".frame");
  if (!f) return null;
  var r = f.getBoundingClientRect();
  return {x:r.x,y:r.y,width:r.width,height:r.height,top:r.top,left:r.left,bottom:r.bottom,right:r.right};
})()`
