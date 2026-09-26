package browser

// FramesPairJS puts a frame of the mockup and its screen on one page, one
// above the other, both roots at the same width (plan 3, step 3). Called with
// the frame's markup, the screen's markup and the width; returns false when
// the frame has no root.
const FramesPairJS = `(function(frameHTML, screenHTML, w){
  var host = document.getElementById("host");
  host.innerHTML = frameHTML;
  var f = host.querySelector(".frame");
  if (!f) return false;
  var label = f.querySelector(":scope > .frame-label");
  if (label) label.remove();
  var s = document.getElementById("screenhost");
  if (!s) { s = document.createElement("div"); s.id = "screenhost"; host.after(s); }
  var sf = f.cloneNode(false);
  sf.innerHTML = screenHTML;
  // A frame drawn in a theme carries it on its root; on a page it is the
  // page's (theme.js puts it on <html>), so the screen gets it from above.
  var fr = f.querySelector(":scope > .tc-root"), th = fr && fr.getAttribute("data-tc-theme");
  if (th) sf.setAttribute("data-tc-theme", th);
  // A frame drawn inside the desktop window's drawing (mk-*): the drawing
  // is page chrome, not the screen's, so the screen is put inside a copy of
  // it — the window, its title bar and status line — and both sides lay out
  // alike; the walk leaves the drawing out on both.
  var fwin = fr && fr.querySelector(":scope > .mk-win"), sroot = sf.querySelector(":scope > .tc-root");
  if (fwin && sroot) {
    var swin = fwin.cloneNode(false), moved = false;
    for (var c = fwin.firstElementChild; c; c = c.nextElementSibling) {
      if (c.classList.contains("mk-titlebar") || c.classList.contains("mk-status")) { swin.appendChild(c.cloneNode(true)); continue; }
      if (!moved) { while (sroot.firstChild) swin.appendChild(sroot.firstChild); moved = true; }
    }
    sroot.appendChild(swin);
  }
  s.innerHTML = "";
  s.appendChild(sf);
  [f, sf].forEach(function(fr){
    fr.style.width = "max-content"; fr.style.maxWidth = "none"; fr.style.margin = "0";
    var r = fr.querySelector(":scope > .tc-root");
    if (r) { r.style.width = w + "px"; r.style.maxWidth = w + "px"; r.style.flex = "none"; }
  });
  return !!f.querySelector(":scope > .tc-root") && !!sf.querySelector(":scope > .tc-root");
})`

// FramesWalkJS walks the frame and the screen in parallel and lists what
// differs: the box of every element and text run, relative to its parent's
// box (so one shift is reported where it starts, not at every descendant),
// with a tolerance of half a pixel, and the computed style that draws it.
// The desktop drawing's frame (mk-titlebar, mk-status) is left out of the
// frame and its other mk-* elements are transparent, as the static
// comparison reads them. An input's inside is the browser's (date, time):
// its box is compared, not its parts.
const FramesWalkJS = `(function(max){
  var PROPS = ["display","position","fontFamily","fontSize","fontWeight","fontStyle","lineHeight","letterSpacing",
    "textTransform","whiteSpace","color","backgroundColor","borderTopWidth","borderRightWidth","borderBottomWidth",
    "borderLeftWidth","borderTopColor","borderTopStyle","borderTopLeftRadius","opacity","textDecorationLine",
    "paddingTop","paddingRight","paddingBottom","paddingLeft","overflowX","textOverflow","visibility"];
  var TOL = 0.5;
  var out = [];
  function chrome(n){ return n.nodeType === 1 && (n.classList.contains("mk-titlebar") || n.classList.contains("mk-status")); }
  function win(n){ if (n.nodeType !== 1 || chrome(n)) return false; for (var i = 0; i < n.classList.length; i++) if (n.classList[i].indexOf("mk-") === 0) return true; return false; }
  // items lists a node's compared children, each with the element its
  // position is measured from: its parent, or the window it stands in when
  // the window is the desktop drawing's (the screen has no window).
  function items(el, from){
    var o = [];
    if (el.tagName === "INPUT" || el.tagName === "TEXTAREA" || el.tagName === "svg" || el.tagName === "SVG") return o;
    for (var n = el.firstChild; n; n = n.nextSibling) {
      if (n.nodeType === 1) {
        if (chrome(n)) continue;
        if (win(n)) { o = o.concat(items(n, n)); continue; }
        o.push([n, from]);
      } else if (n.nodeType === 3 && n.textContent.trim()) o.push([n, from]);
    }
    return o;
  }
  // inner is a window's inside: its box less its border and the title bar
  // the drawing puts first.
  function inner(el){
    var r = el.getBoundingClientRect(), cs = getComputedStyle(el), top = r.top + parseFloat(cs.borderTopWidth);
    for (var c = el.firstElementChild; c; c = c.nextElementSibling) if (c.classList.contains("mk-titlebar")) top = Math.max(top, c.getBoundingClientRect().bottom);
    return {left: r.left + parseFloat(cs.borderLeftWidth), top: top};
  }
  function box(n){ if (n.nodeType === 3) { var r = document.createRange(); r.selectNodeContents(n); return r.getBoundingClientRect(); } return n.getBoundingClientRect(); }
  function label(n){ if (n.nodeType === 3) return "#text"; var c = (n.getAttribute("class") || "").trim(); return n.tagName.toLowerCase() + (c ? "." + c.split(/\s+/).join(".") : ""); }
  function add(path, k, a, b){ if (out.length < max) out.push(path + ": " + k + " " + a + ", screen " + b); }
  function walk(a, b, pa, pb, path){
    if (out.length >= max) return;
    if (a.nodeType !== b.nodeType || (a.nodeType === 1 && a.tagName !== b.tagName)) { add(path, "node", label(a), label(b)); return; }
    var x = box(a), y = box(b);
    if (x.width || x.height || y.width || y.height) {
      var dims = [["x", x.left - pa.left, y.left - pb.left], ["y", x.top - pa.top, y.top - pb.top], ["w", x.width, y.width], ["h", x.height, y.height]];
      for (var i = 0; i < dims.length; i++) if (Math.abs(dims[i][1] - dims[i][2]) > TOL) add(path, dims[i][0], dims[i][1].toFixed(1), dims[i][2].toFixed(1));
    }
    if (a.nodeType === 3) {
      if (a.textContent.replace(/\s+/g, " ").trim() !== b.textContent.replace(/\s+/g, " ").trim()) add(path, "text", JSON.stringify(a.textContent.trim()), JSON.stringify(b.textContent.trim()));
      return;
    }
    var ca = getComputedStyle(a), cb = getComputedStyle(b);
    for (var j = 0; j < PROPS.length; j++) if (ca[PROPS[j]] !== cb[PROPS[j]]) add(path, PROPS[j], ca[PROPS[j]], cb[PROPS[j]]);
    var ia = items(a, null), ib = items(b, null);
    if (ia.length !== ib.length) { add(path, "children", ia.length, ib.length); return; }
    for (var k = 0; k < ia.length; k++) {
      var fa = ia[k][1] ? inner(ia[k][1]) : x, fb = ib[k][1] ? inner(ib[k][1]) : y;
      walk(ia[k][0], ib[k][0], fa, fb, path + " > " + label(ia[k][0]) + "[" + (k + 1) + "]");
    }
  }
  var fr = document.querySelector("#host > .frame > .tc-root"), sr = document.querySelector("#screenhost > .frame > .tc-root");
  if (!fr || !sr) return ["no root"];
  walk(fr, sr, fr.getBoundingClientRect(), sr.getBoundingClientRect(), label(fr));
  return out;
})`

// FramesInvariantsJS checks the screen at one step of the ladder, whatever
// the frame does: no sideways scroll anywhere in it; siblings in the flow do
// not overlap; an open sheet and an open menu are on top (the point at their
// centre is theirs); text stays in its box unless it ends in an ellipsis;
// in the drawer state every control's zone (its box or the ::after that
// widens it, Р-21) is at least 44 px each way.
const FramesInvariantsJS = `(function(max){
  var root = document.querySelector("#screenhost > .frame > .tc-root");
  var out = [];
  function add(s){ if (out.length < max) out.push(s); }
  function label(n){ var c = (n.getAttribute("class") || "").trim(); return n.tagName.toLowerCase() + (c ? "." + c.split(/\s+/).join(".") : ""); }
  function shown(el){ var r = el.getBoundingClientRect(); return r.width > 0 && r.height > 0 && getComputedStyle(el).visibility !== "hidden"; }
  var all = [root].concat(Array.prototype.slice.call(root.querySelectorAll("*")));
  all.forEach(function(el){
    if (el.namespaceURI !== "http://www.w3.org/1999/xhtml") return;
    var cs = getComputedStyle(el);
    if ((cs.overflowX === "auto" || cs.overflowX === "scroll") && el.scrollWidth > el.clientWidth + 1)
      add("scrolls sideways: " + label(el) + " " + el.scrollWidth + " > " + el.clientWidth);
    if (el.children.length > 1 && cs.display !== "contents") {
      var kids = Array.prototype.filter.call(el.children, function(k){
        if (k.namespaceURI !== "http://www.w3.org/1999/xhtml" || !shown(k)) return false;
        var ks = getComputedStyle(k);
        // Inline boxes that wrap span two lines and their rectangles cross
        // though nothing is covered: only boxes of their own are compared.
        if (ks.display === "inline") return false;
        return ks.position === "static" || ks.position === "relative" || ks.position === "sticky";
      });
      for (var i = 0; i < kids.length; i++) for (var j = i + 1; j < kids.length; j++) {
        var a = kids[i].getBoundingClientRect(), b = kids[j].getBoundingClientRect();
        var ox = Math.min(a.right, b.right) - Math.max(a.left, b.left), oy = Math.min(a.bottom, b.bottom) - Math.max(a.top, b.top);
        if (ox > 1 && oy > 1) add("overlap: " + label(kids[i]) + " and " + label(kids[j]) + " in " + label(el));
      }
    }
    if (shown(el) && cs.display !== "inline" && cs.textOverflow !== "ellipsis") {
      var r0 = el.getBoundingClientRect();
      var left = r0.left + parseFloat(cs.borderLeftWidth) + parseFloat(cs.paddingLeft) - 1;
      var right = r0.right - parseFloat(cs.borderRightWidth) - parseFloat(cs.paddingRight) + 1;
      for (var n = el.firstChild; n; n = n.nextSibling) {
        if (n.nodeType !== 3 || !n.textContent.trim()) continue;
        // The text itself, without the white space at its ends: in pre-wrap
        // a trailing space or line break hangs past the box and draws nothing,
        // and so does the space where a line wraps (a rectangle one space wide).
        var tx = n.textContent, i0 = tx.search(/\S/), i1 = tx.length - tx.split("").reverse().join("").search(/\S/);
        var rg = document.createRange(); rg.setStart(n, i0); rg.setEnd(n, i1);
        // pre-wrap keeps runs of spaces, and a run that hangs at a wrap is
        // wider than a space: there the box's own scroll width says it.
        var over = /^(pre-wrap|break-spaces)$/.test(cs.whiteSpace)
          ? el.scrollWidth > el.clientWidth + 1
          : Array.prototype.some.call(rg.getClientRects(), function(tr){ return tr.width > 5 && (tr.right > right || tr.left < left); });
        if (over) { add("text out of its box: " + label(el) + " " + JSON.stringify(tx.trim().slice(0, 30))); break; }
      }
    }
  });
  root.querySelectorAll(".tc-scrim .tc-sheet, .tc-menu").forEach(function(el){
    if (!shown(el)) return;
    el.scrollIntoView({block: "start"});
    var r = el.getBoundingClientRect(), hit = document.elementFromPoint(r.left + r.width / 2, r.top + Math.min(r.height / 2, 20));
    if (hit && !(hit === el || el.contains(hit))) add("under something: " + label(el) + " (the point is " + label(hit) + ")");
  });
  var body = root.querySelector(".tc-shell__body");
  var drawer = body && getComputedStyle(body).getPropertyValue("--tc-shell-mode").trim() === "drawer";
  if (drawer) root.querySelectorAll("button, a[href], input, .tc-check, .tc-tick, .tc-tickrow__tick").forEach(function(el){
    if (!shown(el)) return;
    var r = el.getBoundingClientRect(), w = r.width, h = r.height;
    // The zone is the box or a pseudo-element that widens it: ::after, or
    // ::before where ::after draws something (a pick's arrow).
    ["::after", "::before"].forEach(function(p){
      var a = getComputedStyle(el, p);
      if (a.content !== "none" && a.position === "absolute" && a.pointerEvents !== "none") { w = Math.max(w, parseFloat(a.width) || 0); h = Math.max(h, parseFloat(a.height) || 0); }
    });
    var wrap = el.parentElement ? getComputedStyle(el.parentElement, "::after") : null;
    if ((w < 43.5 || h < 43.5) && wrap && wrap.content !== "none" && wrap.position === "absolute") { w = Math.max(w, parseFloat(wrap.width) || 0); h = Math.max(h, parseFloat(wrap.height) || 0); }
    if (w < 43.5 || h < 43.5) add("tap zone under 44: " + label(el) + " " + w.toFixed(0) + "x" + h.toFixed(0));
  });
  return out;
})`
