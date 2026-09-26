package catalog

import (
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html"

	"carrel.dev/typecase/internal/fragment"
)

// Ladder is the ladder of widths the catalog lays its cells out on: the same
// steps as the mockup page's switch and the browser checks
// (browser.MockupLadderWidths; a test holds them equal).
var Ladder = []int{1351, 1191, 1190, 1008, 1007, 820, 641, 640, 460, 390}

// Frame is a component or fragment frame of the mockup (Р-16 of plan 3): not
// a route, shown in the catalog next to the widget it is built around.
type Frame struct {
	ID      string
	Caption string
	Block   string // the block of the first element that is one: the widget it shows
	HTML    string // the frame's drawing as the mockup has it
}

// Frames reads the component and fragment frames of the mockup under root.
func Frames(root string) ([]Frame, error) {
	m, err := fragment.Load(root)
	if err != nil {
		return nil, err
	}
	var out []Frame
	for _, f := range m.Frames {
		if f.Route() || mark(f) == "кандидат" {
			continue
		}
		r, err := f.Root()
		if err != nil {
			continue
		}
		var b strings.Builder
		if err := html.Render(&b, r); err != nil {
			return nil, err
		}
		out = append(out, Frame{ID: f.ID, Caption: caption(f), Block: firstBlock(r), HTML: b.String()})
	}
	return out, nil
}

// mark is the mark of a frame's label (компонент, фрагмент, список, меню,
// кандидат), or "". A candidate is a frame of a function not taken into
// work: neither a route nor a catalog entry (Р-16).
func mark(f *fragment.Frame) string {
	for c := f.Node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && hasClass(c, "frame-label") {
			for s := c.FirstChild; s != nil; s = s.NextSibling {
				if s.Type == html.ElementNode && hasClass(s, "frag") {
					return textOf(s)
				}
			}
		}
	}
	return ""
}

func caption(f *fragment.Frame) string {
	for c := f.Node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && hasClass(c, "frame-label") {
			for s := c.FirstChild; s != nil; s = s.NextSibling {
				if s.Type == html.ElementNode && s.Data == "span" && s.FirstChild != nil {
					return textOf(s)
				}
			}
		}
	}
	return f.ID
}

func firstBlock(n *html.Node) string {
	var found string
	var walk func(n *html.Node)
	walk = func(n *html.Node) {
		for c := n.FirstChild; c != nil && found == ""; c = c.NextSibling {
			if c.Type != html.ElementNode {
				continue
			}
			if b := fragment.Block(c); b != "" {
				found = b
				return
			}
			walk(c)
		}
	}
	walk(n)
	if found == "" {
		// A frame that shows a block's element alone (a table's file name):
		// the block is the element's.
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type != html.ElementNode {
				continue
			}
			for _, a := range c.Attr {
				if a.Key != "class" {
					continue
				}
				for _, cl := range strings.Fields(a.Val) {
					if i := strings.Index(cl, "__"); strings.HasPrefix(cl, "tc-") && i > 0 {
						return cl[:i]
					}
				}
			}
		}
	}
	return found
}

func hasClass(n *html.Node, cls string) bool {
	for _, a := range n.Attr {
		if a.Key == "class" {
			for _, c := range strings.Fields(a.Val) {
				if c == cls {
					return true
				}
			}
		}
	}
	return false
}

func textOf(n *html.Node) string {
	var b strings.Builder
	var walk func(n *html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.TextNode {
			b.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(n)
	return strings.TrimSpace(b.String())
}

// ladderCSS is the page's width switch: a radio per step sets every cell's
// width, as the mockup page's switch sets its frames' (the cell is a tcbox
// container, so its width drives the container queries).
func ladderCSS() string {
	var b strings.Builder
	for _, w := range Ladder {
		n := strconv.Itoa(w)
		b.WriteString("body:has(#w" + n + ":checked) .catalog-demo{width:" + n + "px;max-width:none}\n")
	}
	return b.String()
}

// ladderHTML is the switch itself, with the theme links theme.js reads.
func ladderHTML() string {
	var b strings.Builder
	b.WriteString(`<form class="catalog-ladder"><label><input type="radio" name="w" id="wauto" checked> as wide as the page</label>`)
	for _, w := range Ladder {
		n := strconv.Itoa(w)
		b.WriteString(`<label><input type="radio" name="w" id="w` + n + `"> ` + n + `</label>`)
	}
	b.WriteString(`<span class="catalog-theme"><a href="?theme=light">light</a> <a href="?theme=dark">dark</a></span></form>`)
	return b.String()
}

// roomJS is the page's own layout script, on both pages (it is not widget
// behaviour): an open menu, the search's suggestions and a scrim hang out of
// their box, as the mockup draws them, so a cell that shows one is padded
// until it is inside.
// It runs again when the width switch or the window changes the layout.
const roomJS = `<script>
(function () {
	var sel = ".tc-menu.is-pop, [class*='tc-drop__panel'], .tc-top__suggest, .tc-scrim";
	function room() {
		document.querySelectorAll(".catalog-demo").forEach(function (cell) {
			if (!cell.querySelector(sel)) return;
			cell.style.padding = "";
			for (var pass = 0; pass < 3; pass++) {
				var c = cell.getBoundingClientRect(), cs = getComputedStyle(cell);
				var need = { bottom: 0, left: 0, right: 0 };
				cell.querySelectorAll(sel).forEach(function (m) {
					if (!m.getClientRects().length) return;
					var r = m.getBoundingClientRect();
					need.bottom = Math.max(need.bottom, r.bottom - c.bottom);
					need.left = Math.max(need.left, c.left - r.left);
					need.right = Math.max(need.right, r.right - c.right);
				});
				if (need.bottom <= 0 && need.left <= 0 && need.right <= 0) break;
				["bottom", "left", "right"].forEach(function (side) {
					if (need[side] > 0) {
						var p = "padding" + side[0].toUpperCase() + side.slice(1);
						cell.style[p] = (parseFloat(cs[p]) + need[side] + 8) + "px";
					}
				});
			}
		});
	}
	// After tc.js has laid the bars out: what went to a hidden "⋯" needs no room.
	function later() {
		requestAnimationFrame(function () {
			requestAnimationFrame(room);
		});
	}
	document.addEventListener("change", function (e) {
		if (e.target.name === "w") later();
	});
	window.addEventListener("resize", later);
	if (document.readyState === "complete") room();
	else window.addEventListener("load", room);
})();
</script>`

// moduleRoot finds the module the catalog is served from, for the mockup.
func moduleRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return fragment.FindRoot(wd)
}
