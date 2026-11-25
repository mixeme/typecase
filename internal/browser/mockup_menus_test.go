package browser_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"carrel.dev/typecase/internal/browser"
	"carrel.dev/typecase/internal/gate"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

// menusProbeJS opens every menu of the frame in #host as a person would — a
// press on its button, the library's script placing it — and says what is
// wrong with each: not under its button, not whole within the frame, an item
// on two lines or overflowing, or the place's letters (capitals, spacing).
const menusProbeJS = `(async function(){
  var host = document.getElementById("host"), root = host.querySelector(".tc-root");
  if (!root) return [];
  var out = [], rr = root.getBoundingClientRect();
  var triggers = [].slice.call(root.querySelectorAll("details.tc-drop > summary, details.tc-drop__anchor > summary, span.tc-drop > .tc-btn:first-child, span.tc-drop__anchor > .tc-btn:first-child"));
  function menuOf(t) {
    var holder = t.parentElement;
    for (var k = holder.firstElementChild; k; k = k.nextElementSibling) {
      if (k === t) continue;
      if (k.classList.contains("tc-menu")) return k;
      var m = k.querySelector && k.querySelector(":scope > .tc-menu");
      if (k.classList.contains("tc-drop__panel--right") && m) return m;
    }
    return null;
  }
  function frame() { return new Promise(function (r) { setTimeout(r, 30); }); }
  function label(t) { return (t.textContent || t.getAttribute("title") || t.getAttribute("aria-label") || "").trim().slice(0, 24) || "⋯"; }
  for (var i = 0; i < triggers.length; i++) {
    var t = triggers[i], m = menuOf(t);
    if (!m || !t.getClientRects().length) continue; // a bar's «⋯» stands only while something is folded
    var holder = t.parentElement, wasOpen = holder.tagName === "DETAILS" ? holder.open : !m.hidden && (!m.parentElement.hidden);
    // A menu drawn open is opened again by its button, as a person opens
    // it: the script places it then.
    if (wasOpen) {
      if (holder.tagName === "DETAILS") holder.open = false;
      else if (m.parentElement.classList.contains("tc-drop__panel--right")) m.parentElement.hidden = true;
      else m.hidden = true;
      await frame();
    }
    t.click();
    await frame();
    // A drop that stands static (a strip's on a phone, Carrel's task 276)
    // hangs its menu from the block it stands in, not from its button.
    var hung = m.parentElement.classList.contains("tc-drop__panel--right") ? m.parentElement : m;
    var anchor = getComputedStyle(holder).position === "static" && hung.offsetParent ? hung.offsetParent : t;
    var mr = m.getBoundingClientRect(), br = anchor.getBoundingClientRect(), name = label(t), bad = [];
    if (anchor !== t) {
      // the block's content box: its padding is not where a menu starts
      var ac = getComputedStyle(anchor);
      br = { top: br.top, bottom: br.bottom, left: br.left + parseFloat(ac.paddingLeft), right: br.right - parseFloat(ac.paddingRight) };
    }
    if (mr.width === 0 || mr.height === 0) bad.push("not shown");
    else {
      var under = mr.top >= br.bottom - 2 && mr.top <= br.bottom + 14, over = mr.bottom <= br.top + 2 && mr.bottom >= br.top - 14;
      if (!under && !over) bad.push("not at its button (" + Math.round(mr.top - br.bottom) + " px)");
      var start = Math.abs(mr.left - br.left) <= 2, end = Math.abs(mr.right - br.right) <= 2;
      var moved = mr.left <= rr.left + 13 || mr.right >= rr.right - 13;
      if (!start && !end && !moved) bad.push("not at its button's start or end");
      if (mr.left < rr.left - 1 || mr.right > rr.right + 1) bad.push("out of the frame by " + Math.round(Math.max(rr.left - mr.left, mr.right - rr.right)) + " px");
      var cs = getComputedStyle(m);
      if (cs.textTransform !== "none") bad.push("letters " + cs.textTransform);
      if (cs.letterSpacing !== "normal" && cs.letterSpacing !== "0px") bad.push("spacing " + cs.letterSpacing);
      [].forEach.call(m.querySelectorAll(".tc-menu__item"), function (it) {
        if (!it.getClientRects().length) return;
        if (it.scrollWidth > it.clientWidth + 1) bad.push("item «" + it.textContent.trim().slice(0, 20) + "» overflows");
        else if (it.getBoundingClientRect().height > 40) bad.push("item «" + it.textContent.trim().slice(0, 20) + "» on two lines");
      });
    }
    if (bad.length) out.push(name + ": " + bad.join(", "));
    {
      if (holder.tagName === "DETAILS") holder.open = false;
      else t.click();
      document.dispatchEvent(new KeyboardEvent("keydown", { key: "Escape", bubbles: true }));
      await frame();
    }
  }
  return out;
})()`

// menusPending are the frames whose menus wait for work named here — the
// list to go to nothing, as the ladder's.
var menusPending = map[string]string{}

// TestMockupMenus (Carrel's task 283): every menu of every frame, opened by
// its button, stands under it at its start or end (or moved in by the
// script), whole within the frame, its items on one line, in the menu's own
// letters — the closed menu shape compares only its button.
func TestMockupMenus(t *testing.T) {
	needChrome(t)
	if os.Getenv("CI") == "" && testing.Short() {
		t.Skip("needs Chrome")
	}
	root := browser.ModuleRoot(t)
	html, err := browser.LoadMockupHTML(root, "")
	if err != nil {
		t.Fatal(err)
	}
	parts, err := gate.ParseMockupParts(string(html))
	if err != nil {
		t.Fatal(err)
	}
	h := browser.NewHarness(t)
	run, cancel := h.Tab(40 * time.Minute)
	defer cancel()
	shell := strings.Replace(browser.MockupShotShell(parts.Styles, parts.Sprite), "</body>", `<script src="/tc.js"></script></body>`, 1)
	url := h.RegisterPage("menus", shell)
	if err := run(chromedp.EmulateViewport(int64(browser.MockupShotViewport), int64(browser.MockupShotMaxHeight)), chromedp.Navigate(url), chromedp.WaitReady("body")); err != nil {
		t.Fatal(err)
	}
	widths := []int{1351, 820, 390}
	if v := os.Getenv("TC_MENUS_WIDTHS"); v != "" {
		widths = nil
		for _, s := range strings.Split(v, ",") {
			n, _ := strconv.Atoi(s)
			widths = append(widths, n)
		}
	}
	only := map[string]bool{}
	for _, id := range strings.Split(os.Getenv("TC_FRAMES"), ",") {
		if id != "" {
			only[id] = true
		}
	}
	for _, fr := range parts.Frames {
		if len(only) > 0 && !only[fr.ID] {
			continue
		}
		if !strings.Contains(fr.HTML, "tc-drop") {
			continue
		}
		raw, err := json.Marshal(fr.HTML)
		if err != nil {
			t.Fatal(err)
		}
		for _, w := range widths {
			js := fmt.Sprintf(`(function(){
  var host = document.getElementById("host");
  host.innerHTML = %s;
  var f = host.querySelector(".frame");
  if (!f) return false;
  f.style.width = "max-content"; f.style.maxWidth = "none"; f.style.margin = "0";
  var root = f.querySelector(".tc-root");
  if (root) { root.style.width = %d + "px"; root.style.maxWidth = %d + "px"; root.style.flex = "none"; }
  return true;
})()`, raw, w, w)
			var found []string
			if err := run(chromedp.Evaluate(js, nil), chromedp.Sleep(20*time.Millisecond), chromedp.Evaluate(menusProbeJS, &found, func(p *runtime.EvaluateParams) *runtime.EvaluateParams { return p.WithAwaitPromise(true) })); err != nil {
				t.Fatalf("%s @%d: %v", fr.ID, w, err)
			}
			if len(found) == 0 {
				continue
			}
			if why, ok := menusPending[fr.ID]; ok {
				t.Logf("%s @%d pending: %s (%s)", fr.ID, w, strings.Join(found, "; "), why)
				continue
			}
			t.Errorf("%s @%d: %s", fr.ID, w, strings.Join(found, "; "))
		}
	}
}
