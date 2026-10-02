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

	document.addEventListener("input", function (e) {
		var t = e.target;
		if (t.classList && t.classList.contains("tc-in__add") && t.closest("[data-tc-suggest]")) suggest(t);
	});
})();
