package browser

// MockupLadderWidths is the shared width ladder of the mockup, the screens' checks, and
// the catalog (docs/visual/mockup.md). Pairs sit on both sides of the Р-3 thresholds.
var MockupLadderWidths = []int{1351, 1191, 1190, 1008, 1007, 820, 641, 640, 460, 390}

// MockupLadderProbeJS checks one painted frame: no horizontal scroll of the
// frame or of any box inside it (a scrolling .tc-main hid a too-wide bar),
// in-flow siblings do not overlap, text stays in its box except ellipsis,
// and in drawer mode every tappable has a 44 px hit zone (same idea as tapHitJS).
const MockupLadderProbeJS = `(() => {
  const root = document.querySelector(".tc-root");
  const frame = document.querySelector(".frame");
  const issues = [];
  if (!root || !frame) return {issues: ["missing .tc-root or .frame"], taps: [], mode: ""};

  const cls = (el) => {
    let c = "";
    try { c = el.getAttribute("class") || ""; } catch (e) {}
    if (typeof c !== "string" || !c || c.indexOf("[object") === 0) return el.tagName || "?";
    return c;
  };
  const hScroll = Math.max(root.scrollWidth - root.clientWidth, frame.scrollWidth - frame.clientWidth);
  if (hScroll > 2) issues.push("h-scroll " + Math.round(hScroll) + "px");
  for (const el of root.querySelectorAll("*")) {
    const ox = getComputedStyle(el).overflowX;
    if (ox !== "auto" && ox !== "scroll") continue;
    // A table in a text scrolls sideways by design: what does not fit after
    // wrapping between words, or a row kept to one line (tc-doc__table).
    if (el.classList && el.classList.contains("tc-doc__scroll")) continue;
    const d = el.scrollWidth - el.clientWidth;
    if (d > 2) issues.push("inner h-scroll " + cls(el).split(" ")[0] + " " + Math.round(d) + "px");
  }
  const frameR = frame.getBoundingClientRect();
  const flowKids = (el) => {
    const out = [];
    for (const k of el.children) {
      const tag = (k.tagName || "").toUpperCase();
      if (tag === "SVG" || tag === "PATH" || tag === "USE" || tag === "IMG" || tag === "BR" ||
          tag === "B" || tag === "STRONG" || tag === "EM" || tag === "I" || tag === "CODE" ||
          tag === "KBD" || tag === "SPAN") continue;
      const cs = getComputedStyle(k);
      if (cs.display === "none" || cs.visibility === "hidden" || cs.display === "inline") continue;
      if (cs.position === "absolute" || cs.position === "fixed") continue;
      const r = k.getBoundingClientRect();
      if (r.width < 0.5 || r.height < 0.5) continue;
      out.push(k);
    }
    return out;
  };
  const walkOverlap = (el) => {
    const kids = flowKids(el);
    for (let i = 0; i < kids.length; i++) {
      const a = kids[i].getBoundingClientRect();
      for (let j = i + 1; j < kids.length; j++) {
        const b = kids[j].getBoundingClientRect();
        const ox = Math.min(a.right, b.right) - Math.max(a.left, b.left);
        const oy = Math.min(a.bottom, b.bottom) - Math.max(a.top, b.top);
        if (ox > 2 && oy > 2) {
          issues.push("overlap " + cls(kids[i]).split(" ")[0] + " / " + cls(kids[j]).split(" ")[0]);
        }
      }
      walkOverlap(kids[i]);
    }
  };
  walkOverlap(root);

  const skipText = (el) => {
    const tag = (el.tagName || "").toUpperCase();
    if (tag === "SVG" || tag === "PATH" || tag === "USE" || tag === "INPUT" ||
        tag === "TEXTAREA" || tag === "IMG" || tag === "BR" || tag === "LABEL" ||
        tag === "B" || tag === "STRONG" || tag === "KBD") return true;
    const cs = getComputedStyle(el);
    if (cs.display === "none" || cs.visibility === "hidden") return true;
    if (cs.textOverflow === "ellipsis") return true;
    if (cs.overflow === "hidden" || cs.overflowX === "hidden" || cs.overflow === "clip") return true;
    if (el.matches("button, a[href], input, textarea, select, .tc-tag, .tc-btn, .tc-seg__opt")) return true;
    if (!el.childElementCount && !(el.textContent || "").trim()) return true;
    return false;
  };
  const walkText = (el) => {
    if (!skipText(el) && !el.childElementCount) {
      if (el.scrollWidth > el.clientWidth + 1) {
        issues.push("text " + cls(el).split(" ")[0]);
      }
    }
    for (const k of el.children) walkText(k);
  };
  walkText(root);

  const body = root.querySelector(".tc-shell__body");
  const mode = body ? getComputedStyle(body).getPropertyValue("--tc-shell-mode").trim() : "";
  const taps = [];
  if (mode === "drawer") {
    const sels = root.querySelectorAll(
      "button, a[href], input, textarea, select, .tc-check, .tc-tick, .tc-tickrow__tick, .tc-close, .tc-nav__item, .tc-seg__opt"
    );
    const vw = window.innerWidth, vh = window.innerHeight;
    for (const el of sels) {
      const cs = getComputedStyle(el);
      if (cs.display === "none" || cs.visibility === "hidden") continue;
      const r = el.getBoundingClientRect();
      if (r.width < 0.5 || r.height < 0.5) continue;
      if (r.bottom < 0 || r.top > vh || r.right < 0 || r.left > vw) continue;
      const cx = r.x + r.width / 2, cy = r.y + r.height / 2;
      const mid = document.elementFromPoint(cx, cy);
      if (!mid || !(el === mid || el.contains(mid) || mid.contains(el))) continue;
      let miss = "";
      for (const [x, y] of [[cx - 21, cy], [cx + 21, cy], [cx, cy - 21], [cx, cy + 21]]) {
        if (x < 0 || y < 0 || x >= vw || y >= vh) continue;
        if (x < frameR.left || x > frameR.right || y < frameR.top || y > frameR.bottom) continue;
        const hit = document.elementFromPoint(x, y);
        if (!hit) { miss = Math.round(x - cx) + "," + Math.round(y - cy) + " hit=null"; break; }
        if (el === hit || el.contains(hit) || hit.contains(el)) continue;
        if (hit.closest && hit.closest("button, a[href], input, textarea, select, label, .tc-check, .tc-tick, .tc-tickrow__tick, .tc-close, .tc-nav__item, .tc-seg__opt, .tc-menu, .tc-menu__item, .tc-scrim, .tc-sheet, .tc-field, .tc-range, .tc-search, .tc-rubric, .tc-overlay")) continue;
        miss = Math.round(x - cx) + "," + Math.round(y - cy) + " hit=" + cls(hit).split(/\s+/).slice(0, 2).join(".");
        break;
      }
      if (miss) {
        taps.push(cls(el).split(/\s+/).slice(0, 2).join(".") + " @" + miss);
      }
    }
  }
  return {issues: issues.slice(0, 8), taps: taps.slice(0, 8), mode};
})()`
