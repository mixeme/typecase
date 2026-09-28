(function () {
	"use strict";

	// tc.js gives the mockup's markup the behaviour the frames cannot draw.
	// It reads the markup the widgets print and never adds a CSS rule: what a
	// state looks like is the mockup's CSS; what is shown or hidden is the
	// hidden attribute, backed by an inline display where a rule of the
	// mockup would show the element anyway (a scrim is a flex box).

	var focusableSel =
		'a[href], button:not([disabled]), input:not([disabled]), textarea:not([disabled]), select:not([disabled]), [tabindex]:not([tabindex="-1"])';

	function closest(el, sel) {
		return el && el.nodeType === 1 ? el.closest(sel) : null;
	}

	function hide(el) {
		el.hidden = true;
		el.style.display = "none";
	}

	function show(el) {
		el.hidden = false;
		el.style.display = "";
	}

	function visible(el) {
		return !el.hidden && el.style.display !== "none" && el.getClientRects().length > 0;
	}

	// Data: a mark's colour and a bar's share come as data (Р-18); markup
	// has no inline style, so the script sets the property the mockup's CSS
	// reads: [data-tc-colour]{background:var(--tc-colour)}.
	function initDynamic(root) {
		root.querySelectorAll("[data-tc-colour]").forEach(function (el) {
			var c = el.getAttribute("data-tc-colour");
			if (c) el.style.setProperty("--tc-colour", c);
		});
		root.querySelectorAll("[data-tc-width]").forEach(function (el) {
			var w = el.getAttribute("data-tc-width");
			if (w) el.style.setProperty("--tc-width", w);
		});
	}

	// Menus. A trigger and the popup it opens, as the widgets print them:
	// Drop — span.tc-drop > trigger + .tc-menu; the top bar's TopDrop —
	// span.tc-drop__anchor > trigger + panel > .tc-menu; a bar's «⋯» followed
	// by its menu.
	var triggerSel = ".tc-drop > :first-child, .tc-drop__anchor > :first-child, .tc-bar__more";

	function popupOf(trigger) {
		var p = trigger.parentElement;
		if (!p) return null;
		if (p.classList.contains("tc-drop") && p.firstElementChild === trigger) {
			return p.querySelector(":scope > .tc-menu");
		}
		if (p.classList.contains("tc-drop__anchor") && p.firstElementChild === trigger) {
			return trigger.nextElementSibling;
		}
		var n = trigger.nextElementSibling;
		if (trigger.classList.contains("tc-bar__more") && n && n.classList.contains("tc-menu")) return n;
		return null;
	}

	function menuIn(popup) {
		return popup.classList.contains("tc-menu") ? popup : popup.querySelector(".tc-menu");
	}

	var openTrigger = null;

	function initMenus(root) {
		root.querySelectorAll(triggerSel).forEach(function (t) {
			var pop = popupOf(t);
			if (!pop) return;
			t.setAttribute("aria-haspopup", "menu");
			var open = !pop.hidden;
			t.setAttribute("aria-expanded", open ? "true" : "false");
			if (!open) hide(pop);
		});
	}

	function placeMenu(trigger, pop) {
		var menu = menuIn(pop);
		if (!menu) return;
		menu.style.top = "";
		menu.style.bottom = "";
		var r = menu.getBoundingClientRect();
		var t = trigger.getBoundingClientRect();
		if (r.bottom > window.innerHeight && t.top - r.height > 0) {
			// No room below: the menu opens above its button.
			menu.style.top = "auto";
			menu.style.bottom = "100%";
		}
	}

	function openMenu(trigger) {
		var pop = popupOf(trigger);
		if (!pop) return;
		closeMenu(false);
		show(pop);
		trigger.setAttribute("aria-expanded", "true");
		openTrigger = trigger;
		placeMenu(trigger, pop);
		var first = pop.querySelector('[role="menuitem"]');
		if (first) first.focus();
	}

	function closeMenu(restoreFocus) {
		var t = openTrigger;
		if (!t) return false;
		var pop = popupOf(t);
		if (pop) hide(pop);
		t.setAttribute("aria-expanded", "false");
		openTrigger = null;
		if (restoreFocus) t.focus();
		return true;
	}

	function menuItems() {
		if (!openTrigger) return [];
		var pop = popupOf(openTrigger);
		return pop ? Array.prototype.slice.call(pop.querySelectorAll('[role="menuitem"]')) : [];
	}

	function onMenuKey(e) {
		var items = menuItems();
		if (!items.length) return false;
		var i = items.indexOf(document.activeElement);
		var next = -1;
		if (e.key === "ArrowDown") next = i < 0 ? 0 : Math.min(i + 1, items.length - 1);
		else if (e.key === "ArrowUp") next = i < 0 ? items.length - 1 : Math.max(i - 1, 0);
		else if (e.key === "Home") next = 0;
		else if (e.key === "End") next = items.length - 1;
		else if (e.key.length === 1 && !e.ctrlKey && !e.metaKey && !e.altKey) {
			var ch = e.key.toLowerCase();
			for (var n = 1; n <= items.length; n++) {
				var j = (i + n + items.length) % items.length;
				if (items[j].textContent.trim().toLowerCase().charAt(0) === ch) {
					next = j;
					break;
				}
			}
		}
		if (next < 0) return false;
		e.preventDefault();
		items[next].focus();
		return true;
	}

	// Layers: a sheet on its scrim, a side panel. What opens one is a control
	// with data-tc-open naming it; what closes one is its close button, a
	// click on the scrim around the sheet, or Esc — the topmost layer first.
	var openers = new WeakMap();

	function trap(container) {
		function onKey(e) {
			if (e.key !== "Tab") return;
			var nodes = Array.prototype.filter.call(container.querySelectorAll(focusableSel), visible);
			if (!nodes.length) return;
			var first = nodes[0];
			var last = nodes[nodes.length - 1];
			if (e.shiftKey && document.activeElement === first) {
				e.preventDefault();
				last.focus();
			} else if (!e.shiftKey && document.activeElement === last) {
				e.preventDefault();
				first.focus();
			}
		}
		container.addEventListener("keydown", onKey);
	}

	function openLayer(layer, opener) {
		show(layer);
		if (opener) openers.set(layer, opener);
		var f = Array.prototype.filter.call(layer.querySelectorAll(focusableSel), visible)[0];
		if (f) f.focus();
	}

	function closeLayer(layer) {
		hide(layer);
		var body = layer.parentElement;
		if (layer.classList.contains("tc-side") && body && body.classList.contains("tc-shell__body")) {
			body.classList.remove("tc-shell__body--side");
		}
		var opener = openers.get(layer);
		if (opener && document.contains(opener)) opener.focus();
	}

	function openLayers() {
		var list = Array.prototype.slice.call(document.querySelectorAll(".tc-scrim, .tc-side")).filter(visible);
		// Scrims above panels, later ones above earlier.
		list.sort(function (a, b) {
			var sa = a.classList.contains("tc-scrim") ? 1 : 0;
			var sb = b.classList.contains("tc-scrim") ? 1 : 0;
			return sa - sb;
		});
		return list;
	}

	function layerOf(el) {
		return closest(el, ".tc-scrim, .tc-side");
	}

	function initLayers(root) {
		root.querySelectorAll(".tc-scrim, .tc-side").forEach(function (layer) {
			if (layer.hidden) hide(layer);
			if (layer.classList.contains("tc-scrim")) trap(layer);
		});
	}

	// The shell's state (Р-3) is a property of the mockup's CSS: full, rail or
	// drawer on tc-shell__body, set by container queries. The script reads it
	// and writes it on the shell's root for what needs to know; it keeps no
	// threshold of its own.
	function shellMode(shell) {
		var body = shell.querySelector(".tc-shell__body");
		if (!body) return "";
		return getComputedStyle(body).getPropertyValue("--tc-shell-mode").trim() || "full";
	}

	function syncShell(shell) {
		var mode = shellMode(shell);
		if (mode && shell.getAttribute("data-tc-mode") !== mode) shell.setAttribute("data-tc-mode", mode);
	}

	// The bars (Р-4): what goes under «⋯» and when is the mockup's CSS. The
	// script only catches a bar the CSS did not fit: it marks it and lets it
	// scroll sideways, so nothing is cut off out of reach.
	function checkBar(bar) {
		bar.style.overflowX = "";
		var over = bar.scrollWidth > bar.clientWidth + 1;
		if (over) {
			bar.setAttribute("data-tc-overflow", "");
			bar.style.overflowX = "auto";
		} else {
			bar.removeAttribute("data-tc-overflow");
		}
		return over;
	}

	function layoutUpdates() {
		document.querySelectorAll(".tc-shell").forEach(syncShell);
		var n = 0;
		document.querySelectorAll(".tc-bar, .tc-docbar, .tc-head").forEach(function (b) {
			if (checkBar(b)) n++;
		});
		return n;
	}

	function onClick(e) {
		var t = closest(e.target, triggerSel);
		if (t && popupOf(t)) {
			e.preventDefault();
			if (openTrigger === t) closeMenu(false);
			else openMenu(t);
			return;
		}
		if (closest(e.target, '[role="menuitem"]') && openTrigger) {
			closeMenu(true);
			return;
		}
		if (openTrigger && !closest(e.target, ".tc-menu")) closeMenu(false);

		var opt = closest(e.target, ".tc-seg__opt");
		if (opt) {
			chooseSeg(opt);
			return;
		}

		var add = closest(e.target, "[data-tc-add]");
		if (add) {
			var tpl = document.getElementById(add.getAttribute("data-tc-add"));
			if (tpl && tpl.content) {
				e.preventDefault();
				addRow(tpl);
			}
			return;
		}
		// A link to a sheet by its address (#id) opens it in place; without the
		// script the address opens it (.tc-scrim:target).
		var hash = closest(e.target, 'a[href^="#"]');
		var sheet = hash && hash.getAttribute("href").length > 1 ? document.getElementById(hash.getAttribute("href").slice(1)) : null;
		if (sheet && sheet.classList.contains("tc-scrim")) {
			e.preventDefault();
			openLayer(sheet, hash);
			return;
		}
		var opener = closest(e.target, "[data-tc-open]");
		// The current section of the navigation opens its sources only in the
		// rail of icons, where they are not in the rail; otherwise it is a link.
		var shell = opener && opener.classList.contains("tc-nav__item") ? closest(opener, ".tc-shell") : null;
		if (shell && shellMode(shell) !== "rail") opener = null;
		if (opener) {
			var target = document.getElementById(opener.getAttribute("data-tc-open"));
			if (target) {
				e.preventDefault();
				openLayer(target, opener);
			}
			return;
		}
		var close = closest(e.target, ".tc-close");
		if (close) {
			var layer = layerOf(close);
			if (layer) {
				e.preventDefault();
				closeLayer(layer);
			}
			return;
		}
		if (e.target.classList && e.target.classList.contains("tc-scrim")) closeLayer(e.target);
	}

	// A segmented switch is a radio group: pressing an option chooses it, the
	// arrows move the choice along the switch (L-4 tc-seg).
	function chooseSeg(opt) {
		var seg = closest(opt, ".tc-seg");
		if (!seg) return;
		seg.querySelectorAll(".tc-seg__opt").forEach(function (o) {
			o.classList.toggle("is-on", o === opt);
			if (o === opt) o.setAttribute("aria-checked", "true");
			else o.removeAttribute("aria-checked");
		});
		// Over a table in a text the switch is Wrap / One line (tc-doc__table).
		var table = closest(seg, ".tc-doc__table");
		if (table) table.classList.toggle("is-line", opt !== seg.querySelector(".tc-seg__opt"));
		if (seg.hasAttribute("data-tc-theme-switch")) {
			var box = opt.querySelector("input");
			setTheme((box ? box.value : opt.textContent).trim().toLowerCase());
		}
	}

	// The theme's switch keeps the choice in this browser, where theme.js
	// reads it before the first paint: light or dark, or auto — the system's.
	function setTheme(v) {
		try {
			if (v === "light" || v === "dark") localStorage.setItem("tc-theme", v);
			else localStorage.removeItem("tc-theme");
		} catch (e) {}
		if (v !== "light" && v !== "dark") v = window.matchMedia && window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light";
		document.documentElement.setAttribute("data-tc-theme", v);
	}

	// «Add …» adds its template's row above the template: «{n}» in the names
	// becomes the smallest number the form does not use for that field yet
	// (FormRow.Template). Without the script the button is a submit, and the
	// server answers with one more row.
	function addRow(tpl) {
		var scope = closest(tpl, "form") || document;
		var row = tpl.content.cloneNode(true);
		var named = row.querySelectorAll("[name*='{n}']");
		var n = 0;
		var taken = function (k) {
			return Array.prototype.some.call(named, function (f) {
				return scope.querySelector('[name="' + f.getAttribute("name").split("{n}").join(String(k)) + '"]');
			});
		};
		while (taken(n)) n++;
		named.forEach(function (f) {
			f.setAttribute("name", f.getAttribute("name").split("{n}").join(String(n)));
		});
		var first = row.querySelector("input, select, textarea");
		tpl.parentNode.insertBefore(row, tpl);
		if (first) first.focus();
	}

	// A file dragged over the page shows the drop sheet (data-tc-drop); the
	// drop sends the named form with the file. Import stays a button too.
	function dropSheet() {
		return document.querySelector(".tc-scrim[data-tc-drop]");
	}

	function carriesFiles(e) {
		return e.dataTransfer && Array.prototype.indexOf.call(e.dataTransfer.types || [], "Files") >= 0;
	}

	function onDrag(e) {
		var sheet = dropSheet();
		if (!sheet || !carriesFiles(e)) return;
		e.preventDefault();
		if (e.type === "dragleave") {
			if (!e.relatedTarget) hide(sheet);
			return;
		}
		if (e.type === "dragenter" || e.type === "dragover") {
			show(sheet);
			return;
		}
		hide(sheet);
		var form = document.getElementById(sheet.getAttribute("data-tc-drop"));
		var input = form && form.querySelector('input[type="file"]');
		if (!input || !e.dataTransfer.files.length) return;
		input.files = e.dataTransfer.files;
		form.submit();
	}

	// A check that switches a field of its form off while it is ticked
	// (data-tc-disables, FormRow.Disables): the field is off and not sent.
	function syncDisables(box) {
		var scope = box.form || document;
		var name = box.getAttribute("data-tc-disables");
		scope.querySelectorAll('[name="' + name + '"]').forEach(function (f) {
			f.disabled = box.checked;
		});
	}

	function initDisables(root) {
		root.querySelectorAll("input[data-tc-disables]").forEach(syncDisables);
	}

	function onChange(e) {
		if (e.target.matches && e.target.matches("input[data-tc-disables]")) syncDisables(e.target);
	}

	function onSegKey(e) {
		var step = { ArrowRight: 1, ArrowDown: 1, ArrowLeft: -1, ArrowUp: -1 }[e.key];
		var opt = step && closest(e.target, ".tc-seg__opt");
		if (!opt) return false;
		var opts = Array.prototype.slice.call(closest(opt, ".tc-seg").querySelectorAll(".tc-seg__opt"));
		var next = opts[(opts.indexOf(opt) + step + opts.length) % opts.length];
		e.preventDefault();
		chooseSeg(next);
		next.focus();
		return true;
	}

	function onKeydown(e) {
		if (openTrigger && onMenuKey(e)) return;
		if (onSegKey(e)) return;
		if (e.key !== "Escape") return;
		if (closeMenu(true)) {
			e.preventDefault();
			return;
		}
		if (closest(e.target, "input, textarea, select")) return;
		var layers = openLayers();
		if (!layers.length) return;
		e.preventDefault();
		closeLayer(layers[layers.length - 1]);
	}

	function init(root) {
		initDynamic(root);
		initMenus(root);
		initLayers(root);
		initDisables(root);
		layoutUpdates();
	}

	document.addEventListener("click", onClick);
	document.addEventListener("change", onChange);
	["dragenter", "dragover", "dragleave", "drop"].forEach(function (t) {
		document.addEventListener(t, onDrag);
	});
	document.addEventListener("keydown", onKeydown);
	if (typeof ResizeObserver !== "undefined") {
		var ro = new ResizeObserver(function () {
			layoutUpdates();
		});
		var observe = function () {
			document.querySelectorAll(".tc-shell, .tc-bar, .tc-docbar, .tc-head").forEach(function (el) {
				ro.observe(el);
			});
		};
		if (document.readyState === "loading") document.addEventListener("DOMContentLoaded", observe);
		else observe();
	}
	window.addEventListener("resize", layoutUpdates);

	window.__tc = { init: init, layoutUpdates: layoutUpdates, shellMode: shellMode };
	if (document.readyState === "loading") {
		document.addEventListener("DOMContentLoaded", function () {
			init(document);
		});
	} else {
		init(document);
	}
})();
