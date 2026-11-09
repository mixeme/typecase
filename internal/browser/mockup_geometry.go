package browser

// MockupGeometryProbeJS measures one painted frame the way an eye checks a
// layout (the owner, 07.11: the doubled left indent was found by eye; a
// check should find its kind). It answers two questions:
//
//   - rows: where the first ink of each row of a table or a list starts,
//     against the first ink of the heading the rows stand under — the
//     nearest heading before them in the same block. A row that starts
//     elsewhere is reported with the difference in pixels.
//   - out: an in-flow box that sticks out of its parent's border box
//     sideways — a bleed wider than the room it bleeds into.
//   - squeezed: a part of a bar (a tool bar, a document bar, a row of
//     actions, a sheet's foot) narrower than it is when it may not shrink —
//     a pick squeezed to its mark stays inside the bar, and the eye still
//     sees it: a bar holds by ranks and «⋯», not by squeezing.
//
// The first ink is the left of the first text, mark, icon, tick or field
// that is seen; what a closed details or a hidden box holds is not seen.
const MockupGeometryProbeJS = `(() => {
  const root = document.querySelector(".tc-root");
  if (!root) return {rows: [], out: [], squeezed: []};
  const seen = (el) => {
    if (!el.getClientRects().length) return false;
    if (el.checkVisibility && !el.checkVisibility({contentVisibilityAuto: true, visibilityProperty: true})) return false;
    return true;
  };
  const marks = "svg, img, input:not([type=hidden]), select, textarea, button, .tc-mark, .tc-tickrow__tick, .tc-tick, .tc-avatar";
  const ink = (el) => {
    let best = Infinity;
    const w = document.createTreeWalker(el, NodeFilter.SHOW_TEXT);
    let n;
    while ((n = w.nextNode())) {
      if (!n.textContent.trim()) continue;
      const p = n.parentElement;
      if (!p || !seen(p)) continue;
      const r = document.createRange();
      r.selectNodeContents(n);
      for (const q of r.getClientRects()) if (q.width > 0.5) best = Math.min(best, q.left);
    }
    for (const e of el.querySelectorAll(marks)) {
      if (!seen(e)) continue;
      const q = e.getBoundingClientRect();
      if (q.width > 0.5) best = Math.min(best, q.left);
    }
    return best;
  };
  const name = (el) => {
    const c = (el.getAttribute && el.getAttribute("class")) || el.tagName.toLowerCase();
    return String(c).split(/\s+/).slice(0, 2).join(".");
  };
  const text = (el) => (el.textContent || "").trim().replace(/\s+/g, " ").slice(0, 24);
  const blocks = ".tc-sheet__pick-list, .tc-main, .tc-side, .tc-sheet, .tc-dialog, .tc-card, .tc-fset, .tc-sec, .tc-shell__rail, .tc-menu, .tc-root";
  const heads = [...root.querySelectorAll(".tc-rubric, .tc-head__title, .tc-dialog__title, .tc-sheet__title")].filter(seen);
  const rows = [];
  // a head row of a table is left out: its first cell is empty over a column of ticks
  for (const r of root.querySelectorAll(".tc-table > tbody > tr, .tc-list > .tc-row")) {
    if (!seen(r)) continue;
    const block = r.closest(blocks);
    let h = null;
    for (const x of heads) {
      if (!(x.compareDocumentPosition(r) & Node.DOCUMENT_POSITION_FOLLOWING)) continue;
      if (x.closest(blocks) !== block && !block.contains(x)) continue;
      h = x;
    }
    if (!h) continue;
    // a row of a table whose first cell is empty (the folder above has no
    // tick) starts where its next cell does, by design
    if (r.tagName === "TR" && r.cells.length && ink(r.cells[0]) === Infinity) continue;
    const d = Math.round(ink(r) - ink(h));
    if (Number.isFinite(d) && d !== 0) rows.push(d + " " + name(r) + " «" + text(r) + "» under " + name(h) + " «" + text(h) + "»");
  }
  // A box stays inside the nearest box that is seen as one: a page, a
  // panel, a sheet, a dialog, a card, a form, a menu. A bleed into the
  // padding of such a box is by design; past its edge it is not.
  const boxes = ".tc-main, .tc-side, .tc-sheet, .tc-dialog, .tc-card, .tc-form, .tc-shell__rail, .tc-menu, .tc-scrim, .tc-root";
  const out = [];
  for (const el of root.querySelectorAll("*")) {
    const p = el.parentElement && el.parentElement.closest(boxes);
    if (!p) continue;
    const cs = getComputedStyle(el);
    if (cs.position === "absolute" || cs.position === "fixed" || cs.display === "inline" || cs.display === "contents") continue;
    if (!seen(el)) continue;
    // what scrolls sideways shows the rest by scrolling
    let scrolls = false;
    for (let a = el.parentElement; a && a !== p; a = a.parentElement) {
      const ox = getComputedStyle(a).overflowX;
      if (ox !== "visible") { scrolls = true; break; }
    }
    if (scrolls) continue;
    const a = el.getBoundingClientRect(), b = p.getBoundingClientRect();
    if (a.width < 1 || a.height < 1) continue;
    const d = Math.max(a.right - b.right, b.left - a.left);
    if (d > 1) out.push(Math.round(d) + " " + name(el) + " out of " + name(p));
  }
  const squeezed = [];
  for (const b of root.querySelectorAll(".tc-bar, .tc-docbar, .tc-acts, .tc-sheet__foot")) {
    if (!seen(b)) continue;
    const cs = getComputedStyle(b);
    if (!cs.display.includes("flex") || cs.flexDirection.startsWith("column")) continue;
    const kids = [...b.children].filter((k) => {
      const kc = getComputedStyle(k);
      // a field one types into gives way by design: its text scrolls inside
      if (k.matches("input, textarea") || k.querySelector("input[type=search], input[type=text], input:not([type]), textarea")) return false;
      return kc.display !== "none" && kc.position !== "absolute" && kc.position !== "fixed" && parseFloat(kc.flexShrink) !== 0 && seen(k);
    });
    // squeezed is narrower than what it holds and no taller for it: text
    // that wraps grows down and is not squeezed; a field that fills the
    // rest of the line is as wide as it likes
    const was = kids.map((k) => k.getBoundingClientRect());
    const style = kids.map((k) => k.getAttribute("style"));
    for (const k of kids) {
      k.style.setProperty("flex", "none", "important");
      k.style.setProperty("width", "max-content", "important");
    }
    const free = kids.map((k) => k.getBoundingClientRect());
    kids.forEach((k, i) => (style[i] === null ? k.removeAttribute("style") : k.setAttribute("style", style[i])));
    kids.forEach((k, i) => {
      const d = free[i].width - was[i].width;
      if (free[i].width >= 1 && d > 4 && free[i].height >= was[i].height - 1) squeezed.push(Math.round(d) + " " + name(k) + " «" + text(k) + "» in " + name(b));
    });
  }
  return {rows: [...new Set(rows)], out: [...new Set(out)], squeezed: [...new Set(squeezed)]};
})()`
