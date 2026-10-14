(function () {
	"use strict";

	// tc.js gives the mockup's markup the behaviour the frames cannot draw.
	// It reads the markup the widgets print and never adds a CSS rule: what a
	// state looks like is the mockup's CSS; what is shown or hidden is the
	// hidden attribute, backed by an inline display where a rule of the
	// mockup would show the element anyway (a scrim is a flex box).

	// The page says it has no script until this one runs: what only a page
	// without a script shows (a form check's Apply) goes.
	document.documentElement.removeAttribute("data-tc-noscript");

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

	// holder is the details a closed menu waits in: it opens without a script,
	// and the script opens and closes the details instead of the popup.
	function holder(pop) {
		var d = pop.parentElement;
		return d && d.tagName === "DETAILS" ? d : null;
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
			var d = holder(pop);
			var open = d ? d.open : !pop.hidden;
			t.setAttribute("aria-expanded", open ? "true" : "false");
			if (!open && !d) hide(pop);
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
		var d = holder(pop);
		if (d) d.open = true;
		else show(pop);
		trigger.setAttribute("aria-expanded", "true");
		openTrigger = trigger;
		// A menu with a search (a time zone) opens whole, on its search line.
		var q = pop.querySelector('.tc-menu > input[type="search"]');
		if (q && q.value) {
			q.value = "";
			onInput({ target: q });
		}
		placeMenu(trigger, pop);
		var first = pop.querySelector('.tc-menu > input[type="search"], [role="menuitem"], [role="menuitemradio"]');
		if (first) first.focus();
	}

	function closeMenu(restoreFocus) {
		var t = openTrigger;
		if (!t) return false;
		var pop = popupOf(t);
		if (pop) {
			var d = holder(pop);
			if (d) d.open = false;
			else hide(pop);
		}
		t.setAttribute("aria-expanded", "false");
		openTrigger = null;
		if (restoreFocus) (t.querySelector(".tc-in__add") || t).focus();
		return true;
	}

	function menuItems() {
		if (!openTrigger) return [];
		var pop = popupOf(openTrigger);
		if (!pop) return [];
		return Array.prototype.filter.call(pop.querySelectorAll('[role="menuitem"], [role="menuitemradio"]'), function (el) {
			return visible(closest(el, ".tc-menu__item") || el);
		});
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
		else if (e.key === "Enter" && document.activeElement && document.activeElement.type === "radio" && items.indexOf(document.activeElement) >= 0) {
			e.preventDefault();
			document.activeElement.click();
			return true;
		} else if (e.key.length === 1 && !e.ctrlKey && !e.metaKey && !e.altKey) {
			// Letters typed in a menu's search line or a tag field's line are theirs.
			if (closest(document.activeElement, '.tc-menu > input[type="search"], .tc-in__add')) return false;
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

	// A tag field's suggestions (attendees, Carrel's task 77) come from the
	// live layer (tc-live.js): it puts the server's items into the field's
	// menu and shows them here, the focus staying in the typing line. The core
	// chooses, takes tags away and keeps Enter from sending the edit.
	function showMenu(trigger) {
		var pop = popupOf(trigger);
		if (!pop) return;
		if (openTrigger !== trigger) closeMenu(false);
		var d = holder(pop);
		if (d) d.open = true;
		else show(pop);
		trigger.setAttribute("aria-expanded", "true");
		openTrigger = trigger;
	}

	function addTag(field, item) {
		var tpl = item.querySelector("template");
		var line = field.querySelector(".tc-in__add");
		if (tpl && line) {
			field.insertBefore(document.importNode(tpl.content, true), line);
			field.insertBefore(document.createTextNode(" "), line);
			line.value = "";
		}
		closeMenu(true);
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

	// A bar's check that is a form (Compact rows) sends itself when it changes;
	// without a script its Apply does.
	document.addEventListener("change", function (e) {
		var box = e.target;
		if (!box || !box.classList || !box.classList.contains("tc-check__box")) return;
		var f = closest(box, "form.tc-checkform");
		if (f) {
			if (f.requestSubmit) f.requestSubmit();
			else f.submit();
		}
	});

	// A queue that folds in place (data-tc-fold-…, Queue.Fold): its toggle
	// hides the list and says Show, the total says what the queue says
	// folded once the upload's words are known; pressed again, it unfolds.
	function fold(btn) {
		var q = closest(btn, ".tc-queue");
		if (!q || !q.hasAttribute("data-tc-fold-show") || !closest(btn, ".tc-queue__head")) return false;
		var min = !q.classList.contains("is-min");
		q.classList.toggle("is-min", min);
		var label = labelOf(btn);
		if (label) label.textContent = q.getAttribute(min ? "data-tc-fold-show" : "data-tc-fold-hide");
		var list = q.querySelector(":scope > .tc-queue__list");
		if (list) (min ? hide : show)(list);
		var head = q.querySelector(":scope > .tc-queue__head"), total = head && head.querySelector(":scope > .tc-mono");
		var said = head && lastWords.get(head), text = total && total.firstChild;
		if (said && text && text.nodeType === 3) {
			if (!folded.has(text)) folded.set(text, worded.get(text) || text.nodeValue);
			worded.set(text, min ? q.getAttribute("data-tc-fold-total") : folded.get(text));
			fillWords(head, said);
		}
		return true;
	}

	function onClick(e) {
		var toggle = closest(e.target, ".tc-queue__head .tc-btn");
		if (toggle && fold(toggle)) {
			e.preventDefault();
			return;
		}
		// A picked file's Remove takes its row away, the file with it: it
		// does not go with Save.
		var removes = closest(e.target, "[data-tc-remove]");
		if (removes) {
			e.preventDefault();
			var picked = closest(removes, "[data-tc-picked]");
			if (picked) {
				var list = picked.parentNode;
				picked.remove();
				emptyWords(list);
			} else {
				removeRow(removes);
			}
			return;
		}
		// A tag field: its cross takes the tag away, a press elsewhere goes to
		// its typing line; its menu is its suggestions, and an item's template
		// is the tag it adds.
		var cross = closest(e.target, "svg");
		var pill = cross && closest(cross, ".tc-in--tags > .tc-tag--pill");
		if (pill && pill.querySelector(".tc-check__box")) {
			e.preventDefault();
			var line = pill.parentElement.querySelector(".tc-in__add");
			pill.remove();
			if (line) line.focus();
			return;
		}
		var field = closest(e.target, "[data-tc-suggest]");
		if (field) {
			if (!closest(e.target, ".tc-tag")) {
				var add = field.querySelector(".tc-in__add");
				if (add && e.target !== add) add.focus();
			}
			return;
		}
		var sug = openTrigger && openTrigger.hasAttribute("data-tc-suggest") && closest(e.target, '[role="menuitem"]');
		if (sug) {
			e.preventDefault();
			addTag(openTrigger, sug);
			return;
		}
		var t = closest(e.target, triggerSel);
		if (t && popupOf(t)) {
			e.preventDefault();
			if (openTrigger === t) closeMenu(false);
			else openMenu(t);
			return;
		}
		// A form menu (Columns): its items are checkboxes, its handle drags;
		// neither closes the menu.
		if (closest(e.target, "form.tc-menu label.tc-menu__item")) {
			if (closest(e.target, "svg") && closest(e.target, "label") === closest(e.target, "label.tc-menu__item")) e.preventDefault();
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

		// Save as: the text of the element named becomes a file here, in the
		// page (Carrel's task 94: a secret shown once has no address).
		var save = closest(e.target, "[data-tc-save]");
		if (save) {
			var from = document.getElementById(save.getAttribute("data-tc-from"));
			if (from) {
				e.preventDefault();
				var a = document.createElement("a");
				a.href = URL.createObjectURL(new Blob([from.textContent], { type: "text/plain" }));
				a.download = save.getAttribute("data-tc-save");
				document.body.appendChild(a);
				a.click();
				a.remove();
				setTimeout(function () {
					URL.revokeObjectURL(a.href);
				}, 1000);
			}
			return;
		}

		// Copy: the text named, or the text of the element named, goes to the
		// clipboard (Carrel's task 118). A browser that refuses lets the press
		// go on: a link goes to the page that shows the text.
		var copy = closest(e.target, "[data-tc-copy], [data-tc-copy-from]");
		if (copy && navigator.clipboard && navigator.clipboard.writeText) {
			var src = document.getElementById(copy.getAttribute("data-tc-copy-from") || "");
			var text = src ? ("value" in src ? src.value : src.textContent) : copy.getAttribute("data-tc-copy");
			if (text) {
				e.preventDefault();
				navigator.clipboard.writeText(text);
				said(copy);
				return;
			}
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
		var close = closest(e.target, ".tc-close, [data-tc-close]");
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
	function markSeg(seg, opt) {
		seg.querySelectorAll(".tc-seg__opt").forEach(function (o) {
			o.classList.toggle("is-on", o === opt);
			if (o === opt) o.setAttribute("aria-checked", "true");
			else o.removeAttribute("aria-checked");
		});
	}

	function chooseSeg(opt) {
		var seg = closest(opt, ".tc-seg");
		if (!seg) return;
		markSeg(seg, opt);
		// Over a table in a text the switch is Wrap / One line (tc-doc__table).
		var table = closest(seg, ".tc-doc__table");
		if (table) table.classList.toggle("is-line", opt !== seg.querySelector(".tc-seg__opt"));
		if (seg.hasAttribute("data-tc-theme-switch")) setTheme(themeOf(opt));
	}

	function themeOf(opt) {
		var box = opt.querySelector("input");
		return (box ? box.value : opt.textContent).trim().toLowerCase();
	}

	// The theme's switch shows the choice kept in this browser, not the one
	// the server printed: the server does not know it (Carrel's task 89).
	function initTheme(root) {
		var kept = null;
		try {
			kept = localStorage.getItem("tc-theme");
		} catch (e) {}
		if (kept !== "light" && kept !== "dark") kept = "auto";
		root.querySelectorAll(".tc-seg[data-tc-theme-switch]").forEach(function (seg) {
			seg.querySelectorAll(".tc-seg__opt").forEach(function (opt) {
				if (themeOf(opt) !== kept) return;
				var box = opt.querySelector("input");
				if (box) box.checked = true;
				else markSeg(seg, opt);
			});
		});
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
		// The fields an option shows are named with the row's number too
		// (Carrel's task 161).
		var shows = row.querySelectorAll("[data-tc-shows*='{n}']");
		shows.forEach(function (o) {
			o.setAttribute("data-tc-shows", o.getAttribute("data-tc-shows").split("{n}").join(String(n)));
		});
		var lists = Array.prototype.filter.call(row.querySelectorAll("select"), function (l) {
			return l.querySelector("option[data-tc-shows]");
		});
		var first = row.querySelector("input, select, textarea");
		tpl.parentNode.insertBefore(row, tpl);
		lists.forEach(syncShows);
		if (first) first.focus();
	}

	// labelOf is the words of a button: its label, or its last words.
	function labelOf(btn) {
		var label = btn.querySelector(".tc-btn__label");
		if (label) return label;
		for (var n = btn.lastChild; n; n = n.previousSibling) {
			if (n.nodeType === 3 && n.nodeValue.trim()) return n;
		}
		return null;
	}

	// A button that has copied says so for a moment (data-tc-copied). Its
	// right edge stays where it is: the other word takes or gives room at
	// the left, by a margin, so nothing after the button moves.
	function said(btn) {
		var word = btn.getAttribute("data-tc-copied");
		if (!word || btn.hasAttribute("data-tc-said")) return;
		var label = labelOf(btn);
		if (!label) return;
		var was = label.textContent;
		var width = btn.getBoundingClientRect().width;
		btn.setAttribute("data-tc-said", "");
		label.textContent = word;
		btn.style.marginLeft = width - btn.getBoundingClientRect().width + "px";
		setTimeout(function () {
			label.textContent = was;
			btn.style.marginLeft = "";
			btn.removeAttribute("data-tc-said");
		}, 1600);
	}

	// Remove of a row of a repeated field (a phone, an address): the row's
	// fields are emptied and hidden where they stand — they stay in the form
	// under their names, and the server drops the empty row with Save
	// (Carrel's task 86). The labels the first row carries go to the row that
	// is first now.
	function removeRow(btn) {
		var isField = function (n) {
			return n && n.classList.contains("tc-field");
		};
		var row = [];
		for (var n = btn.previousElementSibling; isField(n); n = n.previousElementSibling) row.unshift(n);
		var gap = btn.nextElementSibling;
		var after = [];
		for (var m = gap && gap.classList.contains("tc-grow") ? gap.nextElementSibling : null; isField(m); m = m.nextElementSibling) after.push(m);
		row.forEach(function (f, i) {
			var lab = f.querySelector(":scope > label");
			var to = after[i];
			if (lab && to && !to.querySelector(":scope > label")) {
				lab.removeAttribute("for");
				to.insertBefore(lab, to.firstChild);
			}
			f.querySelectorAll("input, textarea").forEach(function (x) {
				if (x.type === "checkbox" || x.type === "radio") x.checked = false;
				else if (x.type !== "hidden") x.value = "";
			});
			f.hidden = true;
		});
		btn.hidden = true;
		if (gap && gap.classList.contains("tc-grow")) gap.hidden = true;
		var next = after[0] && after[0].querySelector("input, select, textarea, button");
		if (next) next.focus();
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
		give(document.getElementById(sheet.getAttribute("data-tc-drop")), e.dataTransfer.files);
	}

	// A file that was dropped or came from the clipboard goes into the file
	// field of the form named, as a file picked there: a field that keeps
	// its files for Save shows them in its list, any other sends its form —
	// by its submit event, so that the live layer can send it in place.
	function give(form, files) {
		var input =
			form &&
			Array.prototype.filter.call(form.elements, function (x) {
				return x.type === "file";
			})[0];
		if (!input || !files || !files.length) return false;
		try {
			input.files = files;
		} catch (err) {
			return false;
		}
		if (input.classList.contains("tc-file")) input.dispatchEvent(new Event("change", { bubbles: true }));
		else if (form.requestSubmit) form.requestSubmit();
		else form.submit();
		return true;
	}

	// A file pasted from the clipboard (Carrel's task 88). Outside a field it
	// goes where a dropped one would: to the form of the page's drop sheet.
	// In a field it is taken only by a text that names a form (data-tc-paste,
	// Area.Attach) and only when the clipboard has no words — a picture
	// becomes an attachment; anything else is the browser's paste.
	function onPaste(e) {
		var clip = e.clipboardData;
		if (!clip || !clip.files || !clip.files.length) return;
		var field = closest(e.target, "input, textarea, select, [contenteditable]");
		var id = "";
		if (field) {
			if (Array.prototype.indexOf.call(clip.types || [], "text/plain") >= 0) return;
			id = field.getAttribute("data-tc-paste");
		} else {
			var sheet = dropSheet();
			id = sheet && sheet.getAttribute("data-tc-drop");
		}
		if (id && give(document.getElementById(id), clip.files)) e.preventDefault();
	}

	// The items of a form menu are sent in their order (Columns): the handle
	// drags an item along the menu; Alt+↑ / Alt+↓ on an item moves it.
	var dragged = null;

	function formItems(item) {
		return Array.prototype.slice.call(closest(item, "form.tc-menu").querySelectorAll("label.tc-menu__item"));
	}

	function onDragStart(e) {
		var handle = closest(e.target, "form.tc-menu label.tc-menu__item > svg");
		if (!handle) return;
		e.preventDefault();
		dragged = closest(handle, "label.tc-menu__item");
		dragged.classList.add("is-drag");
	}

	function onDragMove(e) {
		if (!dragged) return;
		var over = document.elementFromPoint(e.clientX, e.clientY);
		var target = over && closest(over, "label.tc-menu__item");
		if (!target || target === dragged || formItems(target).indexOf(dragged) < 0) return;
		var r = target.getBoundingClientRect();
		target.parentNode.insertBefore(dragged, e.clientY < r.top + r.height / 2 ? target : target.nextSibling);
	}

	function onDragEnd() {
		if (!dragged) return;
		dragged.classList.remove("is-drag");
		dragged = null;
	}

	function onMoveKey(e) {
		if (!e.altKey || (e.key !== "ArrowUp" && e.key !== "ArrowDown")) return false;
		var item = closest(e.target, "form.tc-menu label.tc-menu__item");
		if (!item) return false;
		var items = formItems(item);
		var i = items.indexOf(item) + (e.key === "ArrowUp" ? -1 : 1);
		if (i < 0 || i >= items.length) return true;
		e.preventDefault();
		var other = items[i];
		other.parentNode.insertBefore(item, e.key === "ArrowUp" ? other : other.nextSibling);
		var box = item.querySelector("input");
		if (box) box.focus();
		return true;
	}

	// A check that switches a field of its form off while it is ticked
	// (data-tc-disables, FormRow.Disables): the field is off and not sent.
	function syncDisables(box) {
		var scope = box.form || document;
		// the names go separated by spaces, as classes do (Carrel's task 111)
		box.getAttribute("data-tc-disables").split(/\s+/).forEach(function (name) {
			if (!name) return;
			scope.querySelectorAll('[name="' + name + '"]').forEach(function (f) {
				f.disabled = box.checked;
			});
		});
	}

	// An option of a pick that shows fields of its form while it is chosen
	// (data-tc-shows, PickOption.Shows; Carrel's task 156): the fields the
	// list's options name are shown for the chosen option and hidden for the
	// others — a hidden field's controls are off and not sent.
	function syncShows(list) {
		var scope = list.form || document, on = {}, all = {};
		Array.prototype.forEach.call(list.options, function (o) {
			(o.getAttribute("data-tc-shows") || "").split(/\s+/).forEach(function (name) {
				if (!name) return;
				all[name] = true;
				if (o.selected) on[name] = true;
			});
		});
		Object.keys(all).forEach(function (name) {
			scope.querySelectorAll('[name="' + name + '"]').forEach(function (f) {
				var field = f.closest(".tc-field") || f;
				if (on[name]) show(field);
				else hide(field);
				f.disabled = !on[name];
			});
		});
	}

	// A button that counts (data-tc-count, Button.Count; the owner's decision
	// of 14.10): the records of its form's tally (template[data-tc-tally])
	// whose conditions hold — «name=value» a tick or a choice of that name
	// with that value, «name» a field of that name not empty, «a|b» either —
	// said again at each change of the form; none turns the button off.
	function tally(form) {
		var tpl = form.querySelector("template[data-tc-tally]");
		var btns = form.querySelectorAll("[data-tc-count]");
		if (!tpl || !btns.length) return;
		var holds = function (a) {
			var i = a.indexOf("="), name = i < 0 ? a : a.slice(0, i), value = i < 0 ? null : a.slice(i + 1);
			return Array.prototype.some.call(form.elements, function (f) {
				if (f.name !== name || f.disabled) return false;
				if (f.type === "checkbox" || f.type === "radio") return f.checked && (value === null || f.value === value);
				return value === null ? f.value !== "" : f.value === value;
			});
		};
		var n = 0;
		tpl.content.querySelectorAll("[data-tc-if]").forEach(function (r) {
			var all = r.getAttribute("data-tc-if").split(/\s+/).every(function (t) {
				return !t || t.split("|").some(holds);
			});
			if (all) n++;
		});
		btns.forEach(function (b) {
			var label = labelOf(b);
			if (!label) return;
			var one = b.getAttribute("data-tc-count-1");
			label.textContent = n === 1 && one ? one : b.getAttribute("data-tc-count").split("{n}").join(String(n));
			b.disabled = n === 0;
		});
	}

	function initTally(root) {
		root.querySelectorAll("form").forEach(function (f) {
			if (f.querySelector("template[data-tc-tally]")) tally(f);
		});
	}

	function initShows(root) {
		root.querySelectorAll("select").forEach(function (list) {
			if (list.querySelector("option[data-tc-shows]")) syncShows(list);
		});
	}

	function initDisables(root) {
		root.querySelectorAll("input[data-tc-disables]").forEach(syncDisables);
	}

	// An end that follows its start (data-tc-follows, Input.Follows; Carrel's
	// task 152): moved by as much as the start was, so the length stays. The
	// start keeps what it held before the change in data-tc-was.
	function startOf(end) {
		var scope = end.form || document;
		return scope.querySelector('[name="' + end.getAttribute("data-tc-follows") + '"]');
	}

	function initFollows(root) {
		root.querySelectorAll("[data-tc-follows]").forEach(function (end) {
			var start = startOf(end);
			if (start && !start.hasAttribute("data-tc-was")) start.setAttribute("data-tc-was", start.value);
		});
	}

	function minutesOf(v) {
		var m = /^(\d\d):(\d\d)/.exec(v);
		return m ? +m[1] * 60 + +m[2] : null;
	}

	function dayOf(v) {
		var m = /^(\d{4})-(\d\d)-(\d\d)$/.exec(v);
		return m ? Date.UTC(+m[1], +m[2] - 1, +m[3]) / 864e5 : null;
	}

	function dayText(n) {
		return new Date(n * 864e5).toISOString().slice(0, 10);
	}

	function two(n) {
		return (n < 10 ? "0" : "") + n;
	}

	function follow(start) {
		var was = start.getAttribute("data-tc-was"), now = start.value;
		start.setAttribute("data-tc-was", now);
		if (!was || !now || !start.name) return;
		var scope = start.form || document;
		scope.querySelectorAll('[data-tc-follows="' + start.name + '"]').forEach(function (end) {
			// an empty end stays empty; one switched off is not touched
			if (!end.value || end.disabled) return;
			if (start.type === "date" && end.type === "date") {
				var d = dayOf(end.value), a = dayOf(was), b = dayOf(now);
				if (d !== null && a !== null && b !== null) end.value = dayText(d + b - a);
				return;
			}
			var t = minutesOf(end.value), from = minutesOf(was), to = minutesOf(now);
			if (t === null || from === null || to === null) return;
			var total = t + to - from, carry = Math.floor(total / 1440);
			total -= carry * 1440;
			end.value = two(Math.floor(total / 60)) + ":" + two(total % 60);
			// a time and its day are one pair: over midnight the date goes along
			var day = carry && end.getAttribute("data-tc-day");
			var date = day && scope.querySelector('[name="' + day + '"]');
			if (date && date.value && !date.disabled && dayOf(date.value) !== null) date.value = dayText(dayOf(date.value) + carry);
		});
	}

	function onChange(e) {
		if (e.target.matches && e.target.matches("input[data-tc-disables]")) syncDisables(e.target);
		if (e.target.matches && e.target.matches("select") && e.target.querySelector("option[data-tc-shows]")) syncShows(e.target);
		if (e.target.hasAttribute && e.target.hasAttribute("data-tc-was")) follow(e.target);
		if (e.target.form && e.target.form.querySelector("template[data-tc-tally]")) tally(e.target.form);
		if (e.target.matches && e.target.matches(".tc-file[data-tc-size]")) sizeFields(e.target);
		// A file picked by a button or a menu item goes with its form at once
		// (Carrel's task 109); without the script «Upload» sends it.
		if (e.target.matches && e.target.matches(".tc-file[data-tc-shows]")) {
			// the line beside the field shows what was picked (Carrel's task 133)
			var line = document.getElementById(e.target.getAttribute("data-tc-shows"));
			if (line) line.value = Array.prototype.map.call(e.target.files, function (f) { return f.name; }).join(", ");
		}
		if (e.target.matches && e.target.matches(".tc-file[data-tc-list], .tc-file[data-tc-preview]") && e.target.files && e.target.files.length) pickFile(e.target);
		else if (e.target.matches && e.target.matches(".tc-file[form], .tc-file[data-tc-to]") && e.target.files && e.target.files.length && e.target.form) {
			// a file of a new record's form goes with the whole form to its
			// own address (Carrel's task 122)
			var form = e.target.form, to = e.target.getAttribute("data-tc-to"), was = form.getAttribute("action");
			if (to) form.setAttribute("action", to);
			if (form.requestSubmit) form.requestSubmit();
			else form.submit();
			if (to) {
				if (was === null) form.removeAttribute("action");
				else form.setAttribute("action", was);
			}
		}
		if (e.target.matches && e.target.matches('.tc-menu__item > input[type="radio"]')) choosePick(e.target);
	}

	// The lengths of the picked files go before the files themselves
	// (data-tc-size, FileField.Size; Carrel's task 119): a hidden field
	// «size.<name>» in bytes for each, the name being the file's — its path
	// inside the folder, where a folder was picked. A server that passes the
	// body on as a stream learns its length from them.
	function sizeFields(input) {
		var name = input.getAttribute("data-tc-size");
		if (!name) return;
		for (var old = input.previousElementSibling; old && old.hasAttribute("data-tc-size-of"); old = input.previousElementSibling) old.remove();
		Array.prototype.forEach.call(input.files, function (f) {
			var h = document.createElement("input");
			h.type = "hidden";
			h.name = name + "." + (f.webkitRelativePath || f.name);
			h.value = f.size;
			h.setAttribute("data-tc-size-of", "");
			if (input.hasAttribute("form")) h.setAttribute("form", input.getAttribute("form"));
			input.parentNode.insertBefore(h, input);
		});
	}

	// A file of an edit goes with its Save (the owner's decision of 02.10):
	// picked, it shows as its list's row at the list's end, the input inside
	// the row, and a fresh input waits for the next file; a photo shows in
	// its picture at once.
	function pickFile(input) {
		var preview = document.getElementById(input.getAttribute("data-tc-preview") || "");
		if (preview && input.files[0]) {
			var img = preview.querySelector("img");
			if (!img) {
				// initials give way to the picture
				img = document.createElement("img");
				img.className = "tc-avatar__img";
				img.alt = "";
				preview.textContent = "";
				preview.classList.remove("is-ph");
				preview.appendChild(img);
			}
			img.src = URL.createObjectURL(input.files[0]);
		}
		var list = document.getElementById(input.getAttribute("data-tc-list") || "");
		var tpl = input.parentNode.querySelector("template[data-tc-row]");
		if (!list || !tpl || !tpl.content.firstElementChild) return;
		// the lengths written at the pick are written again in the rows
		for (var old = input.previousElementSibling; old && old.hasAttribute("data-tc-size-of"); old = input.previousElementSibling) old.remove();
		var fresh = input.cloneNode(false);
		fresh.value = "";
		var picked = eachPicked(input);
		input.parentNode.replaceChild(fresh, input);
		picked.forEach(function (one) {
			var row = tpl.content.firstElementChild.cloneNode(true);
			fillWords(row, { "{name}": one.name, "{size}": one.size });
			row.setAttribute("data-tc-picked", "");
			row.appendChild(one.input);
			sizeFields(one.input);
			list.appendChild(row);
		});
		emptyWords(list);
		if (!closeMenu(true)) {
			var d = closest(fresh, "details");
			if (d) d.open = false;
		}
	}

	// What was picked, a row each (the owner's decision of 09.10: the files
	// are shown by name): every file of several gets an input of its own, so
	// that Remove takes one file back; a folder is one row — its name and
	// how much is in it — for a file keeps its path only in the input it was
	// picked with. Where a browser cannot fill an input, the pick is one row.
	function eachPicked(input) {
		var files = Array.prototype.slice.call(input.files), bytes = 0;
		files.forEach(function (f) {
			bytes += f.size;
		});
		if (files[0] && files[0].webkitRelativePath) {
			return [{ input: input, name: files[0].webkitRelativePath.split("/")[0] + "/", size: files.length + " files, " + fileSize(bytes) }];
		}
		var whole = [{ input: input, name: files.map(function (f) { return f.name; }).join(", "), size: fileSize(bytes) }];
		if (files.length < 2 || !window.DataTransfer) return whole;
		try {
			return files.map(function (f) {
				var one = input.cloneNode(false), dt = new DataTransfer();
				dt.items.add(f);
				one.files = dt.files;
				return { input: one, name: f.name, size: fileSize(f.size) };
			});
		} catch (e) {
			return whole;
		}
	}

	// «Nothing attached» stands while the list has nothing picked.
	function emptyWords(list) {
		var any = !!list.querySelector(":scope > [data-tc-picked]");
		Array.prototype.forEach.call(list.querySelectorAll(":scope > [data-tc-empty]"), function (w) {
			w.hidden = any;
		});
	}

	// The words in braces a template was printed with («{name}») become
	// what they name; the text as printed is remembered, so that words that
	// change — how far an upload went — are written again.
	var worded = new WeakMap(), folded = new WeakMap(), lastWords = new WeakMap();

	function fillWords(el, words) {
		lastWords.set(el, words);
		var walk = document.createTreeWalker(el, NodeFilter.SHOW_TEXT);
		for (var n = walk.nextNode(); n; n = walk.nextNode()) {
			var was = worded.get(n);
			if (was === undefined) {
				if (n.nodeValue.indexOf("{") < 0) continue;
				was = n.nodeValue;
				worded.set(n, was);
			}
			var now = was;
			Object.keys(words).forEach(function (k) {
				now = now.split(k).join(words[k]);
			});
			if (now !== n.nodeValue) n.nodeValue = now;
		}
	}

	function fileSize(b) {
		if (b < 1024) return b + " B";
		if (b < 1024 * 1024) return Math.round(b / 1024) + " KB";
		return (b / 1024 / 1024).toFixed(b < 10 * 1024 * 1024 ? 1 : 0) + " MB";
	}

	// A pick's menu of radio items (a time zone): the chosen item's words
	// become the pick's value — the value it came with keeps its aside («as
	// sent») — and the menu closes. Without the script the radio is chosen
	// all the same and goes with the form.
	function choosePick(radio) {
		var menu = closest(radio, ".tc-menu");
		menu.querySelectorAll('.tc-menu__item > input[name="' + radio.name + '"]').forEach(function (r) {
			closest(r, ".tc-menu__item").classList.toggle("is-on", r === radio);
		});
		var d = holder(menu);
		var pick = d && d.firstElementChild;
		if (!pick || !pick.classList.contains("is-pick")) return;
		if (!pick.hasAttribute("data-tc-first")) pick.setAttribute("data-tc-first", pick.innerHTML);
		if (radio.defaultChecked) pick.innerHTML = pick.getAttribute("data-tc-first");
		else {
			// the item's words, and its collection's mark (Carrel's task 103)
			var item = closest(radio, ".tc-menu__item");
			var mark = item.querySelector(".tc-mark");
			pick.textContent = item.textContent.trim();
			if (mark) pick.insertBefore(mark.cloneNode(true), pick.firstChild);
		}
		if (openTrigger === pick) closeMenu(true);
		else d.open = false;
	}

	// A menu's search line keeps the items that have what is typed; a heading
	// or a rule with nothing left beside it goes too.
	function onInput(e) {
		var q = e.target;
		if (q.matches && q.matches(".tc-top__search input") && suggestOf(q) && suggestOf(q).querySelector("[data-tc-when]")) {
			filterSuggest(q);
			show(suggestOf(q));
		}
		if (!q.matches || !q.matches('.tc-menu > input[type="search"]')) return;
		var text = q.value.trim().toLowerCase();
		showGroups(
			Array.prototype.filter.call(q.parentElement.children, function (el) {
				return el !== q;
			}),
			function (el) {
				return el.classList.contains("tc-menu__item");
			},
			function (el) {
				return !text || el.textContent.toLowerCase().indexOf(text) >= 0;
			}
		);
	}

	// The rows of a menu shown by a test, the rest hidden; what is not a
	// row — a heading — goes with its group, and a rule stands between groups
	// that are left.
	function showGroups(els, isItem, test) {
		var groups = [[]];
		els.forEach(function (el) {
			if (el.classList.contains("tc-menu__sep")) groups.push(el, []);
			else groups[groups.length - 1].push(el);
		});
		var shown = [];
		groups.forEach(function (g, i) {
			if (!Array.isArray(g)) return;
			var any = false;
			g.forEach(function (el) {
				if (!isItem(el)) return;
				var ok = test(el);
				if (ok) show(el);
				else hide(el);
				any = any || ok;
			});
			g.forEach(function (el) {
				if (isItem(el)) return;
				if (any) show(el);
				else hide(el);
			});
			shown[i] = any;
		});
		groups.forEach(function (g, i) {
			if (Array.isArray(g)) return;
			// A rule stands between what is left: something above it, and its
			// own group below.
			var before = false;
			for (var j = i - 1; j >= 0; j--) if (Array.isArray(groups[j]) && shown[j]) before = true;
			if (before && shown[i + 1]) show(g);
			else hide(g);
		});
	}

	// A search's suggestions (Carrel's task 92): the menu under the search
	// line is printed whole — recent searches, jumps to what the page already
	// holds, the rows that ask the servers — and shows what fits what is
	// typed (MenuItem.When): with nothing typed, the history and its row;
	// with something typed, the jumps whose words hold it, the typed set bold,
	// and the rows that send it, «{q}» in their words being what is typed.
	// Enter sends the search's form: all the results. The menu shows while
	// the line or the menu has the focus; Esc hides it, ↓ goes into it.
	function suggestOf(line) {
		var search = closest(line, ".tc-top__search");
		return search && search.querySelector(".tc-top__suggest");
	}

	var plainLabel = new WeakMap();

	function boldFound(label, text) {
		if (!plainLabel.has(label)) plainLabel.set(label, label.textContent);
		var words = plainLabel.get(label);
		var at = text ? words.toLowerCase().indexOf(text.toLowerCase()) : -1;
		label.textContent = "";
		if (at < 0) {
			label.textContent = words;
			return;
		}
		var b = document.createElement("b");
		b.textContent = words.slice(at, at + text.length);
		label.appendChild(document.createTextNode(words.slice(0, at)));
		label.appendChild(b);
		label.appendChild(document.createTextNode(words.slice(at + text.length)));
	}

	function filterSuggest(line) {
		var box = suggestOf(line);
		var menu = box && box.querySelector(".tc-menu");
		if (!menu) return;
		var text = line.value.trim(), low = text.toLowerCase();
		fillWords(menu, { "{q}": text });
		showGroups(
			Array.prototype.slice.call(menu.children),
			function (el) {
				return el.getAttribute("role") === "menuitem" || el.hasAttribute("data-tc-when");
			},
			function (el) {
				var when = el.getAttribute("data-tc-when");
				if (when === "empty") return !text;
				if (when === "typed") return !!text;
				if (when !== "found") return true;
				var label = el.querySelector(".tc-menu__label");
				var ok = !!text && (plainLabel.get(label) || el.textContent).toLowerCase().indexOf(low) >= 0;
				if (label) boldFound(label, ok ? text : "");
				return ok;
			}
		);
	}

	function initSuggest(root) {
		root.querySelectorAll(".tc-top__suggest").forEach(function (box) {
			var line = closest(box, ".tc-top__search").querySelector("input");
			if (!line || !box.querySelector("[data-tc-when]")) return;
			if (document.activeElement === line) filterSuggest(line);
			else hide(box);
		});
	}

	document.addEventListener("focusin", function (e) {
		var line = e.target.matches && e.target.matches(".tc-top__search input") ? e.target : null;
		var box = line && suggestOf(line);
		if (!box || !box.querySelector("[data-tc-when]")) return;
		filterSuggest(line);
		show(box);
	});

	document.addEventListener("focusout", function (e) {
		var search = closest(e.target, ".tc-top__search");
		var box = search && search.querySelector(".tc-top__suggest");
		if (!box || !box.querySelector("[data-tc-when]")) return;
		if (e.relatedTarget && search.contains(e.relatedTarget)) return;
		hide(box);
	});

	function onSuggestKey(e) {
		var search = closest(e.target, ".tc-top__search");
		var box = search && search.querySelector(".tc-top__suggest");
		if (!box || !box.querySelector("[data-tc-when]") || box.hidden) return false;
		var line = search.querySelector("input");
		var items = Array.prototype.filter.call(box.querySelectorAll('[role="menuitem"]'), visible);
		var i = items.indexOf(document.activeElement);
		if (e.key === "Escape") {
			e.preventDefault();
			// the line takes the focus back first: the menu it would show goes
			if (i >= 0) line.focus();
			hide(box);
			return true;
		}
		var next = e.key === "ArrowDown" ? i + 1 : e.key === "ArrowUp" ? i - 1 : null;
		if (next === null || !items.length) return false;
		e.preventDefault();
		if (next < 0) line.focus();
		else items[Math.min(next, items.length - 1)].focus();
		return true;
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

	function onTagKey(e) {
		var line = e.target;
		if (!line.classList || !line.classList.contains("tc-in__add")) return false;
		var field = closest(line, ".tc-in--tags");
		if (e.key === "Enter") {
			e.preventDefault();
			var first = openTrigger === field && menuItems()[0];
			if (first) addTag(field, first);
			return true;
		}
		if (e.key === "Backspace" && !line.value) {
			var pills = field.querySelectorAll(":scope > .tc-tag--pill:has(> .tc-check__box)");
			if (pills.length) {
				e.preventDefault();
				pills[pills.length - 1].remove();
			}
			return true;
		}
		return false;
	}

	// A list's rows: the arrows walk them, Home and End go to its ends; Enter
	// opens the row that has the focus — it is a link, or its title is.
	function rowLink(row) {
		return row.tagName === "A" ? row : row.querySelector("a.tc-row__title");
	}

	function onListKey(e) {
		if (e.ctrlKey || e.metaKey || e.altKey || e.shiftKey) return false;
		if (["ArrowDown", "ArrowUp", "Home", "End"].indexOf(e.key) < 0) return false;
		var row = closest(e.target, ".tc-list .tc-row");
		if (!row || closest(e.target, 'input:not([type="checkbox"]), textarea, select')) return false;
		var links = Array.prototype.map.call(closest(row, ".tc-list").querySelectorAll(".tc-row"), rowLink).filter(function (a) {
			return a && visible(a);
		});
		var i = links.indexOf(rowLink(row));
		var next = { ArrowDown: i + 1, ArrowUp: i - 1, Home: 0, End: links.length - 1 }[e.key];
		// At an end the arrow is the page's again: it scrolls.
		if (!links[next] || next === i) return false;
		e.preventDefault();
		links[next].focus();
		return true;
	}

	// Shortcuts (Carrel's task 87). What is pressed carries its key as data
	// (data-tc-key: «n», «/», «?», «Ctrl+S», «Ctrl+Enter»): the key presses it,
	// or puts the focus into it when it is a field. A letter or a sign does
	// nothing while the cursor is in a field or a menu is open; a chord with a
	// modifier works in a field too, for a button of that field's form only.
	// Nothing that destroys is pressed by a key, and nothing under a sheet.
	function chordOf(e) {
		var k = e.key;
		if (!k || k === "Control" || k === "Shift" || k === "Alt" || k === "Meta") return "";
		// A layout without Latin letters: the key's place says which it is.
		if (k.length === 1 && k.charCodeAt(0) > 127 && /^Key[A-Z]$/.test(e.code || "")) k = e.code.charAt(3);
		// Shift is a part of a chord with a letter or a named key; a sign («?»)
		// is typed with it or without, as the layout has it.
		var shift = e.shiftKey && (k.length > 1 || k.toLowerCase() !== k.toUpperCase());
		return (e.ctrlKey || e.metaKey ? "ctrl+" : "") + (e.altKey ? "alt+" : "") + (shift ? "shift+" : "") + k.toLowerCase();
	}

	function onShortcut(e) {
		if (e.repeat || e.isComposing) return false;
		var chord = chordOf(e);
		if (!chord) return false;
		var mod = e.ctrlKey || e.metaKey || e.altKey;
		var field = closest(e.target, "input, textarea, select, [contenteditable]");
		if (!mod && (field || openTrigger)) return false;
		// Enter and the space bar press what has the focus.
		if (!mod && (chord === "enter" || chord === " ") && closest(e.target, focusableSel)) return false;
		var layers = openLayers();
		var top = layers.length ? layers[layers.length - 1] : null;
		if (top && !top.classList.contains("tc-scrim")) top = null;
		var found = Array.prototype.filter.call(document.querySelectorAll("[data-tc-key]"), function (el) {
			if (el.getAttribute("data-tc-key").replace(/\s+/g, "").toLowerCase() !== chord) return false;
			if (el.disabled || el.getAttribute("aria-disabled") === "true") return false;
			if (el.matches(".is-danger, .is-danger-solid, .tc-menu__item--danger, .tc-nojs")) return false;
			if (top && !top.contains(el)) return false;
			if (field && (!field.form || el.form !== field.form)) return false;
			return true;
		});
		// What is on the screen first; then an item of a menu that is closed.
		var el =
			found.filter(visible)[0] ||
			found.filter(function (el) {
				var layer = layerOf(el);
				return closest(el, ".tc-menu") && (!layer || visible(layer));
			})[0];
		if (!el) return false;
		e.preventDefault();
		if (el.matches("input, textarea, select")) {
			el.focus();
			if (el.select) el.select();
			return true;
		}
		var was = document.activeElement;
		el.click();
		// A sheet opened by an item nobody sees gives the focus back to where
		// it was.
		if (!visible(el)) {
			openLayers().forEach(function (layer) {
				if (openers.get(layer) === el) openers.set(layer, was);
			});
		}
		return true;
	}

	function onKeydown(e) {
		if (onTagKey(e)) return;
		if (onSuggestKey(e)) return;
		if (onMoveKey(e)) return;
		if (openTrigger && onMenuKey(e)) return;
		if (onSegKey(e)) return;
		if (onListKey(e)) return;
		if (onShortcut(e)) return;
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

	// Pictures. A thumbnail's picture stands over the icon of its kind. One
	// that did not arrive — the address refused, the file is no picture —
	// goes, and the tile shows the icon, as a file without a thumbnail does.
	function initPictures(root) {
		root.querySelectorAll("img.tc-thumb__img").forEach(function (img) {
			if (img.complete && img.naturalWidth === 0) hide(img);
		});
	}

	// error does not bubble: it is caught on its way down.
	document.addEventListener(
		"error",
		function (e) {
			var el = e.target;
			if (el && el.classList && el.classList.contains("tc-thumb__img")) hide(el);
		},
		true
	);

	// The bars, heads and shells whose width decides their layout are
	// watched; what the live layer puts in later is watched too.
	var sizes = null;

	function watch(root) {
		if (!sizes) return;
		root.querySelectorAll(".tc-shell, .tc-bar, .tc-docbar, .tc-head").forEach(function (el) {
			sizes.observe(el);
		});
	}

	function init(root) {
		watch(root);
		initDynamic(root);
		initPictures(root);
		initTheme(root);
		initSuggest(root);
		initMenus(root);
		initLayers(root);
		initDisables(root);
		initShows(root);
		initTally(root);
		initFollows(root);
		layoutUpdates();
	}

	document.addEventListener("click", onClick);
	document.addEventListener("change", onChange);
	document.addEventListener("input", onInput);
	document.addEventListener("pointerdown", onDragStart);
	document.addEventListener("pointermove", onDragMove);
	document.addEventListener("pointerup", onDragEnd);
	document.addEventListener("pointercancel", onDragEnd);
	["dragenter", "dragover", "dragleave", "drop"].forEach(function (t) {
		document.addEventListener(t, onDrag);
	});
	document.addEventListener("keydown", onKeydown);
	document.addEventListener("paste", onPaste);
	// Enter in a tag field's typing line chooses, it never sends the edit —
	// on keypress too, where implicit submission happens.
	document.addEventListener("keypress", function (e) {
		if (e.key === "Enter" && e.target.classList && e.target.classList.contains("tc-in__add")) e.preventDefault();
	});
	if (typeof ResizeObserver !== "undefined") {
		sizes = new ResizeObserver(function () {
			layoutUpdates();
		});
		if (document.readyState !== "loading") watch(document);
	}
	window.addEventListener("resize", layoutUpdates);

	// A page installed as an application (Carrel's task 91): the document
	// names the library's worker and the page it shows offline; the worker
	// learns from its address what to keep — the library's files this page
	// loaded, under their base, and that page.
	function register() {
		var root = document.documentElement;
		var worker = root.getAttribute("data-tc-worker"), offline = root.getAttribute("data-tc-offline");
		if (!worker || !offline || !navigator.serviceWorker) return;
		var path = function (u) {
			return new URL(u, location.href).pathname;
		};
		var base = path(worker).replace(/js\/[^\/]*$/, "");
		var files = [];
		document.querySelectorAll('link[rel="stylesheet"][href], script[src]').forEach(function (el) {
			var p = path(el.getAttribute("href") || el.getAttribute("src"));
			if (p.indexOf(base) === 0) files.push(p);
		});
		var url = worker + "?base=" + encodeURIComponent(base) + "&offline=" + encodeURIComponent(path(offline)) + "&files=" + encodeURIComponent(files.join(" "));
		navigator.serviceWorker.register(url, { scope: root.getAttribute("data-tc-scope") || "/" }).catch(function () {});
	}
	register();

	// The live layer (tc-live.js) builds on these and on nothing else.
	window.__tc = {
		init: init,
		layoutUpdates: layoutUpdates,
		shellMode: shellMode,
		menuOf: function (trigger) {
			var pop = popupOf(trigger);
			return pop && menuIn(pop);
		},
		showMenu: showMenu,
		fillWords: fillWords,
		fileSize: fileSize,
		closeMenu: function (trigger) {
			if (openTrigger === trigger) closeMenu(false);
		}
	};
	if (document.readyState === "loading") {
		document.addEventListener("DOMContentLoaded", function () {
			init(document);
		});
	} else {
		init(document);
	}
})();
