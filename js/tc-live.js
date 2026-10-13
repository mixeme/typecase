(function () {
	"use strict";

	// tc-live.js is the library's live layer (L-1): what changes a page
	// without a page load. It asks only addresses the markup names, of the
	// page's own site, and puts the HTML of the library's widgets the server
	// answers where the markup says; it has no routes, no data and no state
	// of its own. Without it every one of these is a plain link or form (U-8).
	// It builds on the core (tc.js) through window.__tc and on nothing else.
	var tc = window.__tc;
	if (!tc) return;

	// A tag field's suggestions (attendees, Carrel's task 77): typing asks
	// data-tc-suggest?q=… and the server answers with menu items; an item's
	// template is the tag it adds (the core adds it).
	var timer = 0;

	function suggest(line) {
		var field = line.closest("[data-tc-suggest]");
		var menu = field && tc.menuOf(field);
		if (!menu) return;
		clearTimeout(timer);
		var q = line.value.trim();
		if (!q) {
			tc.closeMenu(field);
			return;
		}
		timer = setTimeout(function () {
			var url = field.getAttribute("data-tc-suggest");
			url += (url.indexOf("?") < 0 ? "?" : "&") + "q=" + encodeURIComponent(q);
			fetch(url, { credentials: "same-origin", headers: { Accept: "text/html" } })
				.then(function (r) {
					return r.ok ? r.text() : "";
				})
				.then(function (html) {
					if (line.value.trim() !== q) return;
					menu.innerHTML = html.trim();
					if (menu.children.length) tc.showMenu(field);
					else tc.closeMenu(field);
				})
				.catch(function () {});
		}, 150);
	}

	// Where the server suggests, the line's own list (a page without a
	// script, Carrel's task 96) would cover the menu: it goes.
	document.addEventListener("focusin", function (e) {
		var t = e.target;
		if (t.classList && t.classList.contains("tc-in__add") && t.hasAttribute("list") && t.closest("[data-tc-suggest]")) t.removeAttribute("list");
	});

	// An edit's Save that sends files shows how far it went (the owner's
	// decision of 02.10): the line naming the form (data-tc-progress) shows,
	// its meter moves, its Mono says «3.2 of 24 MB»; the answer is the page the
	// server sends, as a plain Save would have it.
	document.addEventListener("submit", function (e) {
		var form = e.target;
		if (e.defaultPrevented || !form.id || !window.FormData || !window.XMLHttpRequest) return;
		var line = document.querySelector('[data-tc-progress="' + form.id + '"]');
		if (!line) return;
		var data = new FormData(form, e.submitter || undefined);
		var files = false;
		data.forEach(function (v) {
			if (v && typeof v === "object" && v.size) files = true;
		});
		if (!files) return;
		e.preventDefault();
		var bar = line.querySelector(".tc-meter__bar"), words = line.querySelector(".tc-mono");
		line.hidden = false;
		Array.prototype.forEach.call(document.querySelectorAll('[form="' + form.id + '"], #' + form.id + ' [type="submit"]'), function (b) {
			b.disabled = true;
		});
		var mb = function (b) {
			return (b / 1024 / 1024).toFixed(1);
		};
		var x = new XMLHttpRequest();
		x.open((form.getAttribute("method") || "get").toUpperCase(), (e.submitter && e.submitter.getAttribute("formaction")) || form.action);
		x.upload.onprogress = function (p) {
			if (!p.lengthComputable) return;
			if (bar) bar.style.width = Math.round((p.loaded / p.total) * 100) + "%";
			if (words) words.textContent = mb(p.loaded) + " of " + mb(p.total) + " MB";
		};
		x.onload = function () {
			if (x.responseURL && x.responseURL !== form.action) {
				location.assign(x.responseURL);
				return;
			}
			document.open();
			document.write(x.responseText);
			document.close();
		};
		x.onerror = function () {
			form.submit();
		};
		x.send(data);
	});

	// A link with data-tc-swap changes the page in place (Carrel's task 104):
	// the page at its address is asked for, and the elements it names are put
	// in place of the same ones here; the address goes to the history. Without
	// the layer — or when the answer has not them — the link simply goes.
	function swap(href, ids, done, failed) {
		fetch(href, { credentials: "same-origin", headers: { Accept: "text/html" } })
			.then(function (r) {
				if (!r.ok) throw new Error(r.status);
				return r.text();
			})
			.then(function (html) {
				var doc = new DOMParser().parseFromString(html, "text/html");
				var pairs = ids.map(function (id) {
					return [document.getElementById(id), doc.getElementById(id)];
				});
				if (pairs.some(function (p) { return !p[0] || !p[1]; })) throw new Error("no " + ids.join(" "));
				pairs.forEach(function (p) {
					var fresh = document.importNode(p[1], true);
					p[0].replaceWith(fresh);
					if (tc.init) tc.init(fresh);
				});
				if (done) done();
			})
			.catch(function () {
				if (failed) failed();
			});
	}

	document.addEventListener("click", function (e) {
		var a = e.target.closest && e.target.closest("a[data-tc-swap]");
		if (!a || e.defaultPrevented || e.button !== 0 || e.ctrlKey || e.metaKey || e.shiftKey || e.altKey) return;
		var ids = a.getAttribute("data-tc-swap").split(/\s+/).filter(Boolean);
		if (!ids.length || !window.DOMParser) return;
		e.preventDefault();
		var href = a.href;
		swap(href, ids, function () {
			history.pushState({ tcSwap: true }, "", href);
			swapped = true;
		}, function () {
			location.assign(href);
		});
	});

	// A tree's node that is a radio button and has an address (TreeRow with
	// Send and Href, Carrel's task 158): choosing it asks the page drawn with
	// this node chosen and puts the named elements — the sheet — in place.
	// The address does not go to the history: the sheet is one place. When
	// the answer fails the choice stays, and the form's own buttons act.
	document.addEventListener("change", function (e) {
		var t = e.target;
		if (!t.matches || !t.matches('input[type="radio"][data-tc-href][data-tc-swap]') || !window.DOMParser) return;
		var ids = t.getAttribute("data-tc-swap").split(/\s+/).filter(Boolean);
		if (ids.length) swap(t.getAttribute("data-tc-href"), ids);
	});
	// Back and forward over swapped pages load the page of the address.
	var swapped = false;
	window.addEventListener("popstate", function () {
		if (swapped) location.reload();
	});

	// Zones (Carrel's task 84). A link or a form inside an element that
	// names a target (data-tc-target — the name of a zone) is asked for from
	// the page: the same address as without the layer, with the headers the
	// document names (data-tc-view-header and -value say «a part of the page
	// is wanted», data-tc-target-header names the zone). The answer is HTML:
	// what goes into the target, and after it other zones, each in a
	// <template data-tc-zone>. A zone is an element that carries its name
	// (data-tc-zone); the names are the consumer's, the layer has none.
	// An empty answer empties the zone and hides it — a panel or a sheet
	// closes, the messages go. What was put in goes through the core's init.
	// An answer that is not HTML, not 2xx, or redirected: the page goes to
	// the address the plain way. The zone the document names as the one that
	// owns the address (data-tc-address-zone) puts the address into the
	// history; Back loads the page whole.
	var doc = document.documentElement;

	function zoneOf(name) {
		var all = document.querySelectorAll("[data-tc-zone]");
		for (var i = 0; i < all.length; i++) if (all[i].getAttribute("data-tc-zone") === name) return all[i];
		return null;
	}

	// The shell's body makes room for a side panel that is shown and takes
	// it back when the panel is not: the class the core takes away on close.
	function room(panel) {
		var body = panel.parentElement;
		if (!panel.classList.contains("tc-side") || !body || !body.classList.contains("tc-shell__body")) return;
		body.classList.toggle("tc-shell__body--side", !panel.hidden);
	}

	function fill(zone, nodes) {
		var only = nodes.children.length === 1 && !nodes.textContent.replace(nodes.firstElementChild.textContent, "").trim() ? nodes.firstElementChild : null;
		if (only && only.getAttribute("data-tc-zone") === zone.getAttribute("data-tc-zone")) {
			// the widget of the zone itself: it takes the zone's place whole
			zone.parentNode.replaceChild(only, zone);
			// the core's init looks inside what it is given: the widget is
			// held in a box for the call, so that it is looked at too
			if (tc.init) {
				var box = document.createElement("div");
				only.parentNode.insertBefore(box, only);
				box.appendChild(only);
				tc.init(box);
				box.parentNode.insertBefore(only, box);
				box.parentNode.removeChild(box);
			}
			room(only);
			return;
		}
		while (zone.firstChild) zone.removeChild(zone.firstChild);
		zone.appendChild(nodes);
		var empty = !zone.firstElementChild && !zone.textContent.trim();
		zone.hidden = empty;
		zone.style.display = empty ? "none" : "";
		room(zone);
		if (!empty && tc.init) tc.init(zone);
	}

	function put(target, html) {
		var zone = zoneOf(target);
		if (!zone) return false;
		var box = document.createElement("template");
		box.innerHTML = html;
		var extra = [];
		Array.prototype.forEach.call(box.content.querySelectorAll("template[data-tc-zone]"), function (t) {
			if (t.parentNode === box.content) extra.push(t);
		});
		extra.forEach(function (t) {
			box.content.removeChild(t);
		});
		fill(zone, box.content);
		extra.forEach(function (t) {
			var z = zoneOf(t.getAttribute("data-tc-zone"));
			if (z) fill(z, t.content);
		});
		if (tc.layoutUpdates) tc.layoutUpdates();
		return true;
	}

	function ask(url, method, body, target, plain) {
		var headers = { Accept: "text/html" };
		var view = doc.getAttribute("data-tc-view-header"), zoneHeader = doc.getAttribute("data-tc-target-header");
		if (view) headers[view] = doc.getAttribute("data-tc-view-value") || "1";
		if (zoneHeader) headers[zoneHeader] = target;
		var field = doc.getAttribute("data-tc-token-field"), tokenHeader = doc.getAttribute("data-tc-token-header");
		if (body && field && tokenHeader && body.get(field)) headers[tokenHeader] = body.get(field);
		fetch(url, { method: method, body: body || undefined, headers: headers, credentials: "same-origin" })
			.then(function (r) {
				var type = r.headers.get("Content-Type") || "";
				if (!r.ok || r.redirected || type.indexOf("text/html") !== 0) throw new Error("not a part");
				return r.text();
			})
			.then(function (html) {
				if (!put(target, html)) throw new Error("no zone " + target);
				if (method === "GET" && doc.getAttribute("data-tc-address-zone") === target) {
					history.pushState({ tcSwap: true }, "", url);
					swapped = true;
				}
			})
			.catch(plain);
	}

	// A zone the mockup's CSS does not draw at this width — the side panel on
	// a narrow screen — cannot take an answer: the link or form goes the plain
	// way, to the page that shows the thing whole. The zone is tried as it
	// would stand filled: shown, the shell's body with room for it.
	function shows(zone) {
		var body = zone.parentElement;
		var side = zone.classList.contains("tc-side") && body && body.classList.contains("tc-shell__body");
		var was = { hidden: zone.hidden, display: zone.style.display, room: side && body.classList.contains("tc-shell__body--side") };
		zone.hidden = false;
		zone.style.display = "";
		if (side) body.classList.add("tc-shell__body--side");
		var drawn = getComputedStyle(zone).display !== "none";
		zone.hidden = was.hidden;
		zone.style.display = was.display;
		if (side && !was.room) body.classList.remove("tc-shell__body--side");
		return drawn;
	}

	function targetOf(el) {
		var t = el && el.closest && el.closest("[data-tc-target]");
		return t ? t.getAttribute("data-tc-target") : "";
	}

	function sameSite(url) {
		try {
			return new URL(url, location.href).origin === location.origin;
		} catch (err) {
			return false;
		}
	}

	document.addEventListener("click", function (e) {
		var a = e.target.closest && e.target.closest("a[href]");
		if (!a || e.defaultPrevented || e.button !== 0 || e.ctrlKey || e.metaKey || e.shiftKey || e.altKey) return;
		if (a.hasAttribute("data-tc-swap") || a.hasAttribute("download") || (a.target && a.target !== "_self")) return;
		var target = targetOf(a);
		if (!target || !sameSite(a.href) || !zoneOf(target) || !shows(zoneOf(target))) return;
		e.preventDefault();
		var href = a.href;
		ask(href, "GET", null, target, function () {
			location.assign(href);
		});
	});

	var sentPlain = false;
	document.addEventListener("submit", function (e) {
		var form = e.target, by = e.submitter || null;
		if (e.defaultPrevented || sentPlain || !window.FormData) return;
		var target = targetOf(by) || targetOf(form);
		if (!target || !zoneOf(target) || !shows(zoneOf(target))) return;
		var action = (by && by.getAttribute("formaction")) || form.getAttribute("action") || location.href;
		if (!sameSite(action)) return;
		var method = ((by && by.getAttribute("formmethod")) || form.getAttribute("method") || "get").toUpperCase();
		var data = new FormData(form, by || undefined);
		var files = false;
		data.forEach(function (v) {
			if (v && typeof v === "object" && v.size) files = true;
		});
		if (files) return; // a form that sends files shows its progress and loads the answer whole
		e.preventDefault();
		var url = new URL(action, location.href), body = data;
		if (method === "GET") {
			url.search = new URLSearchParams(data).toString();
			body = null;
		}
		ask(url.href, method, body, target, function () {
			sentPlain = true;
			if (form.requestSubmit && by) form.requestSubmit(by);
			else form.submit();
			sentPlain = false;
		});
	});

	document.addEventListener("input", function (e) {
		var t = e.target;
		if (t.classList && t.classList.contains("tc-in__add") && t.closest("[data-tc-suggest]")) suggest(t);
	});
})();
