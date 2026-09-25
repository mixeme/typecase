package browser_test

import (
	"encoding/json"
	"math"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/chromedp/chromedp"

	"carrel.dev/typecase/internal/browser"
	"carrel.dev/typecase/internal/screens"
	"carrel.dev/typecase/layout"
)

type lookProbe struct {
	NavHeight       float64       `json:"navHeight"`
	NavCurrentBg    string        `json:"navCurrentBg"`
	CardBg          string        `json:"cardBg"`
	TickRowHeight   float64       `json:"tickRowHeight"`
	ListBorder      float64       `json:"listBorder"`
	BrandHasSVG     bool          `json:"brandHasSVG"`
	SideWidth       float64       `json:"sideWidth"`
	SideDisplay     string        `json:"sideDisplay"`
	MainResolvedBg  string        `json:"mainResolvedBg"`
	BgToken         string        `json:"bgToken"`
	TitleSize       float64       `json:"titleSize"`
	TitleLineHeight float64       `json:"titleLineHeight"`
	SymbolMissing   []string      `json:"symbolMissing"`
	NavEmoji        []string      `json:"navEmoji"`
	Rank2Buttons    int           `json:"rank2Buttons"`
	SheetBare       bool          `json:"sheetBare"`
	SheetDxHint     float64       `json:"sheetDxHint"`
	SheetTaMin      float64       `json:"sheetTaMin"`
	SheetTitleX     float64       `json:"sheetTitleX"`
	SheetInputX     float64       `json:"sheetInputX"`
	MarkHeight      float64       `json:"markHeight"`
	MarkOverlap     bool          `json:"markOverlap"`
	RailWidth       float64       `json:"railWidth"`
	TopHeight       float64       `json:"topHeight"`
	SearchWidth     float64       `json:"searchWidth"`
	SearchFieldH    float64       `json:"searchFieldH"`
	Drawer          bool          `json:"drawer"`
	BodyHScroll     float64       `json:"bodyHScroll"`
	ListWidth       float64       `json:"listWidth"`
	GridRows        []lookGridRow `json:"gridRows"`
	ChipHeight      float64       `json:"chipHeight"`
	QueueWidth      float64       `json:"queueWidth"`
	TdBtnHeight     float64       `json:"tdBtnHeight"`
	IntoX           float64       `json:"intoX"`
	HeadTitleX      float64       `json:"headTitleX"`
	EmptyAlign      string        `json:"emptyAlign"`
	TickLabelX      float64       `json:"tickLabelX"`
	TickBoxX        float64       `json:"tickBoxX"`
	HasTickLink     bool          `json:"hasTickLink"`
}

type lookGridRow struct {
	Class    string   `json:"className"`
	Cols     []string `json:"cols"`
	Slots    int      `json:"slots"`
	ColCount int      `json:"colCount"`
}

type lookResult struct {
	Frames []lookProbe `json:"frames"`
}

const lookProbeJS = `(() => {
  const rgbOf = (value) => {
    const d = document.createElement('div');
    d.style.background = value;
    document.body.appendChild(d);
    const c = getComputedStyle(d).backgroundColor;
    d.remove();
    return c;
  };
  const resolvedBg = (el) => {
    let n = el;
    while (n && n !== document.documentElement) {
      const bg = getComputedStyle(n).backgroundColor;
      if (bg && bg !== 'transparent' && bg !== 'rgba(0, 0, 0, 0)') return bg;
      n = n.parentElement;
    }
    return '';
  };
  function items(row) {
    const walk = (el) => {
      const d = getComputedStyle(el).display;
      if (d === "none") return [];
      if (d === "contents") return [...el.children].flatMap(walk);
      return [el];
    };
    return [...row.children].flatMap(walk);
  }
  const probeRoot = (root) => {
    const railNav = root.querySelector('.tc-nav:not(.tc-nav--bottom) .tc-nav__item.is-current') || root.querySelector('.tc-nav:not(.tc-nav--bottom) .tc-nav__item');
    const nav = railNav;
    const tick = root.querySelector('.tc-tickrow');
    const list = root.querySelector('.tc-list');
    const brandSVG = root.querySelector('.tc-brand svg');
    const side = root.querySelector('.tc-side');
    const main = root.querySelector('.tc-main');
    const title = root.querySelector('.tc-head__title');
    const missing = [];
    root.querySelectorAll('use[href^="#"]').forEach((u) => {
      const id = (u.getAttribute('href') || '').slice(1);
      if (id && !document.getElementById(id)) missing.push(id);
    });
    const navEmoji = [];
    root.querySelectorAll('.tc-nav__icon-svg').forEach((el) => {
      const t = el.textContent.trim();
      if (t) navEmoji.push(t);
    });
    let titleLH = 0, titleSize = 0;
    if (title) {
      const cs = getComputedStyle(title);
      titleLH = parseFloat(cs.lineHeight) || 0;
      titleSize = parseFloat(cs.fontSize) || 0;
    }
    const current = nav && nav.classList.contains('is-current');
    const sheet = root.querySelector('.tc-sheet');
    const sheetField = sheet && sheet.querySelector('.tc-field');
    const sheetIn = sheet && sheet.querySelector('.tc-in');
    const sheetHint = sheetField && sheetField.querySelector('.tc-field__hint');
    const sheetTitle = sheet && sheet.querySelector('.tc-sheet__title');
    let sheetDxHint = 0, sheetTaMin = 0, sheetTitleX = 0, sheetInputX = 0;
    if (sheetIn && sheetHint) {
      sheetDxHint = Math.round(sheetIn.getBoundingClientRect().x - sheetHint.getBoundingClientRect().x);
    }
    if (sheetIn) {
      sheetTaMin = parseFloat(getComputedStyle(sheetIn).minHeight) || 0;
      sheetInputX = sheetIn.getBoundingClientRect().x;
    }
    if (sheetTitle) sheetTitleX = sheetTitle.getBoundingClientRect().x;
    const marks = root.querySelectorAll('.tc-tickrow .tc-mark');
    let markHeight = 0, markOverlap = false;
    if (marks[0]) {
      const mr = marks[0].getBoundingClientRect();
      markHeight = mr.height;
      root.querySelectorAll('.tc-railsec__title').forEach((el) => {
        const tr = el.getBoundingClientRect();
        if (mr.top < tr.bottom && mr.bottom > tr.top) markOverlap = true;
      });
    }
    const rail = root.querySelector('.tc-shell__rail');
    const top = root.querySelector('.tc-top');
    const search = root.querySelector('.tc-search');
    const searchField = root.querySelector('.tc-search__field');
    const shell = root.querySelector('.tc-shell');
    const gridRows = [];
    root.querySelectorAll('.tc-row, .tc-tickrow').forEach((el) => {
      const d = getComputedStyle(el).display;
      if (d !== 'grid') return;
      const cols = getComputedStyle(el).gridTemplateColumns;
      if (!cols || cols === 'none') return;
      const colList = cols.trim().split(/\s+/).filter(Boolean);
      gridRows.push({
        className: el.className,
        cols: colList,
        slots: items(el).length,
        colCount: colList.length,
      });
    });
    const chip = root.querySelector('.tc-chip');
    const queue = root.querySelector('.tc-queue');
    const tdBtn = root.querySelector('.tc-table__td .tc-btn');
    const into = root.querySelector('input[name="into"], .tc-form .tc-in');
    const empty = root.querySelector('.tc-empty');
    const tickLink = root.querySelector('.tc-tickrow__link');
    const tickBox = tickLink && tickLink.closest('.tc-tickrow') && tickLink.closest('.tc-tickrow').querySelector('.tc-tickrow__tick');
    return {
      navHeight: nav ? nav.getBoundingClientRect().height : 0,
      navCurrentBg: current ? getComputedStyle(nav).backgroundColor : '',
      cardBg: rgbOf('var(--tc-card)'),
      tickRowHeight: tick ? tick.getBoundingClientRect().height : 0,
      listBorder: list ? parseFloat(getComputedStyle(list).borderTopWidth) || 0 : -1,
      brandHasSVG: !!brandSVG,
      sideWidth: side ? side.getBoundingClientRect().width : 0,
      sideDisplay: side ? getComputedStyle(side).display : '',
      mainResolvedBg: main ? resolvedBg(main) : '',
      bgToken: rgbOf('var(--tc-bg)'),
      titleSize,
      titleLineHeight: titleLH,
      symbolMissing: missing,
      navEmoji,
      rank2Buttons: root.querySelectorAll('[data-tc-bar-rank2] .tc-btn').length,
      sheetBare: !!(sheetField && sheetField.classList.contains('tc-field--bare')),
      sheetDxHint,
      sheetTaMin,
      sheetTitleX,
      sheetInputX,
      markHeight,
      markOverlap,
      railWidth: rail ? rail.getBoundingClientRect().width : 0,
      topHeight: top ? top.getBoundingClientRect().height : 0,
      searchWidth: search ? search.getBoundingClientRect().width : 0,
      searchFieldH: searchField ? searchField.getBoundingClientRect().height : 0,
      drawer: !!(shell && shell.classList.contains('tc-shell--narrow')),
      bodyHScroll: Math.max(document.body.scrollWidth - document.body.clientWidth, 0),
      listWidth: list ? list.clientWidth : 0,
      gridRows,
      chipHeight: chip ? chip.getBoundingClientRect().height : 0,
      queueWidth: queue ? queue.getBoundingClientRect().width : 0,
      tdBtnHeight: tdBtn ? tdBtn.getBoundingClientRect().height : 0,
      intoX: into ? into.getBoundingClientRect().x : 0,
      headTitleX: title ? title.getBoundingClientRect().x : 0,
      emptyAlign: empty ? getComputedStyle(empty).textAlign : '',
      tickLabelX: tickLink ? tickLink.getBoundingClientRect().x : 0,
      tickBoxX: tickBox ? tickBox.getBoundingClientRect().x : 0,
      hasTickLink: !!tickLink,
    };
  };
  const states = document.querySelectorAll('.screens-frame__mock');
  if (states.length) {
    return { frames: Array.from(states).map(probeRoot) };
  }
  return { frames: [probeRoot(document)] };
})()`

const menuOnTopJS = `(() => {
  const pop = document.querySelector('.tc-tickrow .tc-menu .tc-menu__popover');
  if (!pop) return false;
  const cs = getComputedStyle(pop);
  if (cs.display === 'none' || cs.visibility === 'hidden') return false;
  const r = pop.getBoundingClientRect();
  if (r.width < 2 || r.height < 2) return false;
  const hit = document.elementFromPoint(r.x + r.width / 2, r.y + r.height / 2);
  return !!(hit && (pop === hit || pop.contains(hit)));
})()`

func TestScreensLook(t *testing.T) {
	if !browser.ChromeAvailable() {
		t.Skip("no Chromium-based browser found (set CHROME_PATH)")
	}
	if testing.Short() {
		t.Skip("screen look probe needs Chrome")
	}
	srv := screensServer(t)
	h := browser.NewHarness(t)
	for _, width := range []int{1600, 1280, 900, 480} {
		width := width
		for _, sc := range screens.All() {
			sc := sc
			t.Run(sc.ID+"/"+strconv.Itoa(width), func(t *testing.T) {
				var raw json.RawMessage
				var menuTop bool
				url := screenURL(srv.URL, sc.ID, "")
				actions := []chromedp.Action{layoutSettle(), chromedp.Evaluate(lookProbeJS, &raw)}
				if sc.ID == "files" && width == 1280 {
					actions = append(actions,
						chromedp.Click(".tc-tickrow .tc-menu__control", chromedp.ByQuery),
						chromedp.Sleep(80*time.Millisecond),
						chromedp.Evaluate(menuOnTopJS, &menuTop),
					)
				}
				if err := h.Run(url, width, 900, actions...); err != nil {
					t.Fatal(err)
				}
				var res lookResult
				if err := json.Unmarshal(raw, &res); err != nil {
					t.Fatalf("decode: %v", err)
				}
				if len(res.Frames) == 0 {
					t.Fatal("look probe returned no frames")
				}
				if sc.ID == "states" && len(res.Frames) < 5 {
					t.Fatalf("states frames: got %d want >= 5", len(res.Frames))
				}
				for i, p := range res.Frames {
					checkLookProbe(t, sc.ID, width, i, p)
				}
				if sc.ID == "files" && width == 1280 && !menuTop {
					t.Fatal("files: open row menu is not the top element at its centre")
				}
			})
		}
	}
}

func checkLookProbe(t *testing.T, id string, width, frame int, p lookProbe) {
	t.Helper()
	prefix := id + "/" + strconv.Itoa(width)
	if frame > 0 || id == "states" {
		prefix = prefix + " frame " + strconv.Itoa(frame)
	}
	wantNav, _ := layout.Px(layout.NavItem.Props["height"])
	wantTick, _ := layout.Px(layout.TickRow.MinHeight)
	wantTitle := layoutSelPx(".tc-head__title, .tc-side__title", "font-size")
	wantRail, _ := layout.Px(layout.Rail.Props["width"])
	wantTop, _ := layout.Px(layout.Top.Props["height"])
	wantSearch, _ := layout.Px(layout.Search.Props["width"])
	wantField, _ := layout.Px(layout.SearchField.Props["height"])
	wantMark := layoutSelPx(".tc-tickrow .tc-mark", "height")
	wantSide := layoutSelPx(".tc-side", "width")
	wantQueue, _ := layout.Px(layout.Queue.Props["width"])

	if width >= 900 && p.NavHeight > 0 && !closeEnough(p.NavHeight, wantNav) {
		t.Fatalf("%s .tc-nav__item height %.1f want %.0f", prefix, p.NavHeight, wantNav)
	}
	if p.NavCurrentBg != "" && p.NavCurrentBg != p.CardBg {
		t.Fatalf("%s .is-current background %s want %s (--tc-card)", prefix, p.NavCurrentBg, p.CardBg)
	}
	if width >= 900 && p.TickRowHeight > 0 && !closeEnough(p.TickRowHeight, wantTick) {
		t.Fatalf("%s .tc-tickrow height %.1f want %.0f", prefix, p.TickRowHeight, wantTick)
	}
	if p.ListBorder > 0.5 {
		t.Fatalf("%s .tc-list border %.1f want 0", prefix, p.ListBorder)
	}
	if !p.BrandHasSVG {
		t.Fatalf("%s .tc-brand svg missing", prefix)
	}
	if len(p.SymbolMissing) > 0 {
		t.Fatalf("%s missing sprite symbols: %v", prefix, p.SymbolMissing)
	}
	if len(p.NavEmoji) > 0 {
		t.Fatalf("%s emoji in nav icons: %v", prefix, p.NavEmoji)
	}
	if p.MainResolvedBg != "" && p.MainResolvedBg != p.BgToken {
		t.Fatalf("%s .tc-main background %s want %s (--tc-bg)", prefix, p.MainResolvedBg, p.BgToken)
	}
	if p.TitleSize > 0 && !closeEnough(p.TitleSize, wantTitle) {
		t.Fatalf("%s .tc-head__title font-size %.1f want %.0f", prefix, p.TitleSize, wantTitle)
	}
	if p.TitleSize > 0 && !closeEnough(p.TitleLineHeight, p.TitleSize*1.15) {
		t.Fatalf("%s .tc-head__title line-height %.1f want %.1f (1.15)", prefix, p.TitleLineHeight, p.TitleSize*1.15)
	}
	if p.RailWidth > 0 && !p.Drawer && !closeEnough(p.RailWidth, wantRail) {
		t.Fatalf("%s .tc-shell__rail width %.1f want %.0f", prefix, p.RailWidth, wantRail)
	}
	if p.TopHeight > 0 && !closeEnough(p.TopHeight, wantTop) {
		t.Fatalf("%s .tc-top height %.1f want %.0f", prefix, p.TopHeight, wantTop)
	}
	if p.SearchFieldH > 0 && !closeEnough(p.SearchFieldH, wantField) {
		t.Fatalf("%s .tc-search__field height %.1f want %.0f", prefix, p.SearchFieldH, wantField)
	}
	if width >= 900 && p.SearchWidth > 0 && !closeEnough(p.SearchWidth, wantSearch) {
		t.Fatalf("%s .tc-search width %.1f want %.0f", prefix, p.SearchWidth, wantSearch)
	}
	if p.MarkHeight > 0 && !closeEnough(p.MarkHeight, wantMark) {
		t.Fatalf("%s .tc-tickrow .tc-mark height %.1f want %.0f", prefix, p.MarkHeight, wantMark)
	}
	if p.MarkOverlap {
		t.Fatalf("%s .tc-tickrow .tc-mark overlaps .tc-railsec__title", prefix)
	}
	for _, row := range p.GridRows {
		if row.Slots != row.ColCount {
			t.Fatalf("%s %s slots %d tracks %d", prefix, row.Class, row.Slots, row.ColCount)
		}
		checkRowTracks(t, prefix, p.ListWidth, row)
	}
	if p.HasTickLink && p.TickLabelX <= p.TickBoxX+12 {
		t.Fatalf("%s TickRow link x %.1f want > tick.x+12 (%.1f)", prefix, p.TickLabelX, p.TickBoxX+12)
	}

	if width == 1280 {
		if id == "notes" && p.SideWidth > 0 && !closeEnough(p.SideWidth, wantSide) {
			t.Fatalf("%s .tc-side width %.1f want %.0f", prefix, p.SideWidth, wantSide)
		}
		if id == "notes" && p.Rank2Buttons < 2 {
			t.Fatalf("%s Rank2 buttons %d want >= 2 at 1280", prefix, p.Rank2Buttons)
		}
		if id == "notes-quick" {
			if !p.SheetBare {
				t.Fatalf("%s quick sheet field must be tc-field--bare", prefix)
			}
			if math.Abs(p.SheetDxHint) > 2 {
				t.Fatalf("%s sheet textarea/hint left delta %.0f want 0", prefix, p.SheetDxHint)
			}
			if p.SheetTaMin < 132 {
				t.Fatalf("%s sheet textarea min-height %.0f want >= 132", prefix, p.SheetTaMin)
			}
			if math.Abs(p.SheetTitleX-p.SheetInputX) > 2 {
				t.Fatalf("%s sheet title.x %.1f input.x %.1f want ±2", prefix, p.SheetTitleX, p.SheetInputX)
			}
		}
		if id == "files" {
			if p.ChipHeight > 20 {
				t.Fatalf("%s chip height %.1f want ≤ 20", prefix, p.ChipHeight)
			}
			if p.QueueWidth > 0 && !closeEnough(p.QueueWidth, wantQueue) {
				t.Fatalf("%s .tc-queue width %.1f want %.0f", prefix, p.QueueWidth, wantQueue)
			}
			if p.TdBtnHeight > 24 {
				t.Fatalf("%s td button height %.1f want ≤ 24", prefix, p.TdBtnHeight)
			}
		}
		if id == "io" && p.IntoX > 0 && p.HeadTitleX > 0 && math.Abs(p.IntoX-p.HeadTitleX) > 1 {
			t.Fatalf("%s Into x %.1f heading x %.1f want ±1", prefix, p.IntoX, p.HeadTitleX)
		}
		if id == "states" && p.EmptyAlign != "" && strings.Contains(p.EmptyAlign, "center") {
			t.Fatalf("%s .tc-empty text-align %s want not center", prefix, p.EmptyAlign)
		}
	}
	if id == "notes" && width == 1600 && (p.SideDisplay == "none" || p.SideWidth < 1) {
		t.Fatalf("%s side hidden at 1600", prefix)
	}
	if id == "notes" && width == 900 {
		if p.SideDisplay != "none" && p.SideWidth >= 1 {
			t.Fatalf("%s side visible at 900 (display %s width %.1f)", prefix, p.SideDisplay, p.SideWidth)
		}
		if p.Drawer {
			t.Fatalf("%s rail drawer at 900; want rail in place", prefix)
		}
		if p.RailWidth > 0 && !closeEnough(p.RailWidth, wantRail) {
			t.Fatalf("%s rail missing at 900 (%.1f)", prefix, p.RailWidth)
		}
	}
	if width == 480 && p.BodyHScroll > 1 {
		t.Fatalf("%s body horizontal scroll %.1f", prefix, p.BodyHScroll)
	}
}

func checkRowTracks(t *testing.T, prefix string, listWidth float64, row lookGridRow) {
	t.Helper()
	g, ok := gridForClass(row.Class)
	if !ok {
		return
	}
	tracks := layout.TracksAt(g, listWidth)
	if strings.Contains(row.Class, "tc-row--avatarcols") && len(row.Cols) > 1 {
		want, ok := layout.Px(tracks[1].Size)
		if ok && !closeEnough(parsePx(row.Cols[1]), want) {
			t.Fatalf("%s avatarcols col 2 %s want %.0fpx (list %.0f)", prefix, row.Cols[1], want, listWidth)
		}
	}
	if strings.Contains(row.Class, "tc-row--tickmarkcols") && len(row.Cols) > 0 {
		want, ok := layout.Px(tracks[0].Size)
		if ok && !closeEnough(parsePx(row.Cols[0]), want) {
			t.Fatalf("%s tickmarkcols col 1 %s want %.0fpx", prefix, row.Cols[0], want)
		}
	}
	if strings.Contains(row.Class, "tc-row--timecols") && len(row.Cols) > 1 {
		want, ok := layout.Px(tracks[1].Size)
		if ok && !closeEnough(parsePx(row.Cols[1]), want) {
			t.Fatalf("%s timecols col 2 %s want %.0fpx (list %.0f)", prefix, row.Cols[1], want, listWidth)
		}
	}
	if strings.Contains(row.Class, "tc-row--find") && len(row.Cols) > 1 {
		want, ok := layout.Px(tracks[1].Size)
		if ok && !closeEnough(parsePx(row.Cols[1]), want) {
			t.Fatalf("%s find col 2 %s want %.0fpx (list %.0f)", prefix, row.Cols[1], want, listWidth)
		}
	}
}

func gridForClass(class string) (layout.Grid, bool) {
	switch {
	case strings.Contains(class, "tc-row--excerpt"):
		return layout.RowExcerpt, true
	case strings.Contains(class, "tc-row--avatarcols"):
		return layout.RowAvatarCols, true
	case strings.Contains(class, "tc-row--tickcols"):
		return layout.RowTickCols, true
	case strings.Contains(class, "tc-row--timecols"):
		return layout.RowTimeCols, true
	case strings.Contains(class, "tc-row--tickmarkcols"):
		return layout.RowTickMarkCols, true
	case strings.Contains(class, "tc-row--find"):
		return layout.RowFind, true
	default:
		return layout.Grid{}, false
	}
}

func layoutSelPx(sel, prop string) float64 {
	for _, b := range layout.All().Boxes {
		for _, part := range strings.Split(b.Selector, ",") {
			if strings.TrimSpace(part) == sel || b.Selector == sel {
				if n, ok := layout.Px(b.Props[prop]); ok {
					return n
				}
			}
		}
	}
	for _, f := range layout.All().Flows {
		if f.Selector == sel {
			if n, ok := layout.Px(f.Props[prop]); ok {
				return n
			}
		}
	}
	return 0
}

func parsePx(v string) float64 {
	n, ok := layout.Px(v)
	if ok {
		return n
	}
	n, _ = strconv.ParseFloat(strings.TrimSuffix(v, "px"), 64)
	return n
}

func layoutSettle() chromedp.Action {
	return chromedp.Evaluate(`new Promise(function(resolve) {
  (function tick(n) {
    if (window.__tc && window.__tc.layoutUpdates) {
      window.__tc.layoutUpdates();
      requestAnimationFrame(function() {
        window.__tc.layoutUpdates();
        requestAnimationFrame(resolve);
      });
      return;
    }
    if (n > 200) { resolve(); return; }
    setTimeout(function() { tick(n + 1); }, 25);
  })(0);
})`, nil)
}

func closeEnough(got, want float64) bool {
	return math.Abs(got-want) <= 1
}
