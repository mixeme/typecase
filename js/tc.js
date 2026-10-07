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

	function onClick(e) {
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

	// A button that has copied says so for a moment (data-tc-copied). Its
	// right edge stays where it is: the other word takes or gives room at
	// the left, by a margin, so nothing after the button moves.
	function said(btn) {
		var word = btn.getAttribute("data-tc-copied");
		if (!word || btn.hasAttribute("data-tc-said")) return;
		var label = btn.querySelector(".tc-btn__label");
		if (!label) {
			for (var n = btn.lastChild; n; n = n.previousSibling) {
				if (n.nodeType === 3 && n.nodeValue.trim()) {
					label = n;
					break;
				}
			}
		}
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
		var form = document.getElementById(sheet.getAttribute("data-tc-drop"));
		var input = form && form.querySelector('input[type="file"]');
		if (!input || !e.dataTransfer.files.length) return;
		input.files = e.dataTransfer.files;
		form.submit();
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

	function initDisables(root) {
		root.querySelectorAll("input[data-tc-disables]").forEach(syncDisables);
	}

	function onChange(e) {
		if (e.target.matches && e.target.matches("input[data-tc-disables]")) syncDisables(e.target);
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
		var names = [], bytes = 0;
		Array.prototype.forEach.call(input.files, function (f) {
			names.push(f.name);
			bytes += f.size;
		});
		var row = tpl.content.firstElementChild.cloneNode(true);
		fillWords(row, { "{name}": names.join(", "), "{size}": fileSize(bytes) });
		row.setAttribute("data-tc-picked", "");
		var fresh = input.cloneNode(false);
		input.parentNode.insertBefore(fresh, input);
		row.appendChild(input);
		list.appendChild(row);
		emptyWords(list);
		if (!closeMenu(true)) {
			var d = closest(fresh, "details");
			if (d) d.open = false;
		}
	}

	// «Nothing attached» stands while the list has nothing picked.
	function emptyWords(list) {
		var any = !!list.querySelector(":scope > [data-tc-picked]");
		Array.prototype.forEach.call(list.querySelectorAll(":scope > [data-tc-empty]"), function (w) {
			w.hidden = any;
		});
	}

	function fillWords(el, words) {
		var walk = document.createTreeWalker(el, NodeFilter.SHOW_TEXT);
		for (var n = walk.nextNode(); n; n = walk.nextNode()) {
			Object.keys(words).forEach(function (k) {
				if (n.nodeValue.indexOf(k) >= 0) n.nodeValue = n.nodeValue.split(k).join(words[k]);
			});
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
		if (!q.matches || !q.matches('.tc-menu > input[type="search"]')) return;
		var text = q.value.trim().toLowerCase();
		var groups = [[]];
		Array.prototype.forEach.call(q.parentElement.children, function (el) {
			if (el === q) return;
			if (el.classList.contains("tc-menu__sep")) groups.push(el, []);
			else groups[groups.length - 1].push(el);
		});
		var shown = [];
		groups.forEach(function (g, i) {
			if (!Array.isArray(g)) return;
			var any = false;
			g.forEach(function (el) {
				if (!el.classList.contains("tc-menu__item")) return;
				var ok = !text || el.textContent.toLowerCase().indexOf(text) >= 0;
				if (ok) show(el);
				else hide(el);
				any = any || ok;
			});
			g.forEach(function (el) {
				if (el.classList.contains("tc-menu__item")) return;
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

	function onKeydown(e) {
		if (onTagKey(e)) return;
		if (onMoveKey(e)) return;
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
	document.addEventListener("input", onInput);
	document.addEventListener("pointerdown", onDragStart);
	document.addEventListener("pointermove", onDragMove);
	document.addEventListener("pointerup", onDragEnd);
	document.addEventListener("pointercancel", onDragEnd);
	["dragenter", "dragover", "dragleave", "drop"].forEach(function (t) {
		document.addEventListener(t, onDrag);
	});
	document.addEventListener("keydown", onKeydown);
	// Enter in a tag field's typing line chooses, it never sends the edit —
	// on keypress too, where implicit submission happens.
	document.addEventListener("keypress", function (e) {
		if (e.key === "Enter" && e.target.classList && e.target.classList.contains("tc-in__add")) e.preventDefault();
	});
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
