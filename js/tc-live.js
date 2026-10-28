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

	// The search line's suggestions (Carrel's task 227): the page names
	// their address (data-tc-search-suggest on the root, Document.Suggest);
	// the focus and each pause in typing ask it ?q=…, and the answer — the
	// menu, its items saying when they show — stands under the line, where
	// the core filters it, walks it with the arrows and hides it.
	var searchTimer = 0;

	function searchSuggest(line) {
		var url = document.documentElement.getAttribute("data-tc-search-suggest");
		var search = url && line.closest(".tc-search");
		if (!search) return;
		var wrap = search.parentNode;
		if (!wrap.classList.contains("tc-top__search")) {
			wrap = document.createElement("span");
			wrap.className = "tc-line tc-top__search";
			search.parentNode.insertBefore(wrap, search);
			wrap.appendChild(search);
			line.focus();
		}
		var box = wrap.querySelector(".tc-top__suggest");
		if (!box) {
			box = document.createElement("span");
			box.className = "tc-top__suggest";
			box.hidden = true;
			box.style.display = "none";
			wrap.appendChild(box);
		}
		clearTimeout(searchTimer);
		var q = line.value.trim();
		searchTimer = setTimeout(function () {
			fetch(url + (url.indexOf("?") < 0 ? "?" : "&") + "q=" + encodeURIComponent(q), { credentials: "same-origin", headers: { Accept: "text/html" } })
				.then(function (r) {
					return r.ok ? r.text() : "";
				})
				.then(function (html) {
					if (line.value.trim() !== q) return;
					box.innerHTML = html.trim();
					tc.suggest(line);
				})
				.catch(function () {});
		}, 120);
	}

	document.addEventListener("focusin", function (e) {
		if (e.target.matches && e.target.matches(".tc-top .tc-search .tc-q")) searchSuggest(e.target);
	});
	document.addEventListener("input", function (e) {
		if (e.target.matches && e.target.matches(".tc-top .tc-search .tc-q")) searchSuggest(e.target);
	});

	// Where the server suggests, the line's own list (a page without a
	// script, Carrel's task 96) would cover the menu: it goes.
	document.addEventListener("focusin", function (e) {
		var t = e.target;
		if (t.classList && t.classList.contains("tc-in__add") && t.hasAttribute("list") && t.closest("[data-tc-suggest]")) t.removeAttribute("list");
	});

	// A form that sends files goes from the page when a line shows how far
	// it went or its answer is a part of the page. The line naming the form
	// (data-tc-progress) shows: a run's line moves its meter and its Mono
	// says «3.2 of 24 MB» (an edit's Save, the owner's decision of 02.10); a
	// queue gets a row for every file from its template, and the rows' meters
	// and words move (Carrel's task 88). A form whose target is a zone asks as
	// any form of a zone does, and its answer is put into the zones; any
	// other answer is the page the server sends, as a plain send would have it.
	function queued(queue, files) {
		var tpl = queue.querySelector(":scope > template[data-tc-row]"), list = queue.querySelector(":scope > .tc-queue__list");
		var head = queue.querySelector(":scope > .tc-queue__head"), total = 0, rows = [];
		files.forEach(function (f) {
			total += f.size;
			if (!tpl || !list || !tpl.content.firstElementChild) return;
			var row = tpl.content.firstElementChild.cloneNode(true);
			list.appendChild(row);
			rows.push({ el: row, size: f.size, name: f.webkitRelativePath || f.name });
		});
		return function (part) {
			var sent = part * total, before = 0, done = 0;
			rows.forEach(function (r) {
				var p = r.size ? Math.max(0, Math.min(1, (sent - before) / r.size)) : 1;
				before += r.size;
				if (p >= 1) done++;
				var pc = Math.round(p * 100);
				tc.fillWords(r.el, { "{name}": r.name, "{size}": tc.fileSize(r.size), "{percent}": pc });
				var bar = r.el.querySelector(".tc-meter__bar");
				if (bar) {
					bar.style.setProperty("--tc-width", pc + "%");
					bar.setAttribute("aria-valuenow", pc);
				}
			});
			if (head) tc.fillWords(head, { "{count}": files.length, "{done}": done, "{sent}": tc.fileSize(Math.round(sent)), "{total}": tc.fileSize(total), "{percent}": Math.round(part * 100) });
		};
	}

	// A file field with data-tc-ranges sends only the parts of its file the
	// server asks for (Carrel's task 213): a large archive is read where it
	// lies and never copied whole. The form's other fields with the file's
	// name and size («<name>.name», «<name>.size») go to that address first;
	// it answers with lines «offset length»; the form then goes to its
	// action with those parts under the field's name, each called
	// «offset-length», instead of the file. The page keeps the file under the
	// field's name, for a later step: a form with data-tc-ranges and
	// data-tc-ranges-of (that name) sends parts of it the same way, asking
	// with its own fields — the records of the collections chosen on the
	// next step (Carrel's task 220). When the form names a zone, its answer
	// goes there and the page stays — the file with it; else the answer is
	// the page, as a plain send would have it. Without the layer, or when the
	// asking fails, the form goes the plain way: the whole file.
	var kept = {};
	document.addEventListener("submit", function (e) {
		var form = e.target, by = e.submitter || null;
		if (e.defaultPrevented || sentPlain || !window.FormData || !window.fetch) return;
		var input = Array.prototype.filter.call(form.querySelectorAll('input[type="file"][data-tc-ranges]'), function (i) {
			return !i.disabled && i.files && i.files.length;
		})[0];
		var file, name, asked;
		if (input) {
			file = input.files[0];
			name = input.name;
			asked = input.getAttribute("data-tc-ranges");
			kept[name] = file;
		} else if (form.hasAttribute("data-tc-ranges") && kept[form.getAttribute("data-tc-ranges-of")]) {
			name = form.getAttribute("data-tc-ranges-of");
			file = kept[name];
			asked = form.getAttribute("data-tc-ranges");
		} else return;
		e.preventDefault();
		var data = new FormData(form, by || undefined);
		data.delete(name);
		data.set(name + ".name", file.name);
		data.set(name + ".size", String(file.size));
		var action = new URL((by && by.getAttribute("formaction")) || form.getAttribute("action") || location.href, location.href).href;
		var method = ((by && by.getAttribute("formmethod")) || form.getAttribute("method") || "post").toUpperCase();
		var target = targetOf(by) || targetOf(form);
		if (target && !zoneOf(target)) target = "";
		var plain = function () {
			sentPlain = true;
			if (form.requestSubmit && by) form.requestSubmit(by);
			else form.submit();
			sentPlain = false;
		};
		fetch(new URL(asked, location.href).href, { method: "POST", body: data, credentials: "same-origin" })
			.then(function (r) {
				if (!r.ok) throw new Error("ranges " + r.status);
				return r.text();
			})
			.then(function (text) {
				text.split(/\r?\n/).forEach(function (line) {
					var p = line.trim().split(/\s+/), at = Number(p[0]), length = Number(p[1]);
					if (p.length !== 2 || !(at >= 0) || !(length > 0)) return;
					data.append(name, file.slice(at, at + length), at + "-" + length);
				});
				if (target) {
					ask(action, method, data, target, plain, function () {
						tc.unbusy(form);
					});
					return;
				}
				return fetch(action, { method: method, body: data, credentials: "same-origin" }).then(function (r) {
					if (r.redirected) {
						location.assign(r.url);
						return;
					}
					return r.text().then(function (html) {
						document.open();
						document.write(html);
						document.close();
					});
				});
			})
			.catch(plain);
	});

	// A file field with data-tc-taken asks, before its files go, which of
	// their names are taken where they go (Carrel's task 214: the server
	// keeps nothing of a file while it asks). The form's hidden fields and a
	// «name» a file go to that address; it answers the taken names, a line
	// each. A row of the queue's template[data-tc-taken] asks about each —
	// its buttons' values are keep, replace and skip — and when every one is
	// answered, the skipped files leave the field, the answers go with the
	// form as «taken.<name>», and the form is sent the way it would have
	// been. Nothing taken, no queue to ask in, or the asking fails: the form
	// is sent at once.
	function askTaken(form, by, input, line) {
		var files = Array.prototype.slice.call(input.files), nameOf = function (f) {
			return f.webkitRelativePath || f.name;
		};
		// the answers of an earlier asking are not this one's
		Array.prototype.filter.call(form.elements, function (x) {
			return x.hasAttribute("data-tc-answer");
		}).forEach(function (x) {
			x.remove();
		});
		var names = new FormData();
		Array.prototype.forEach.call(form.elements, function (x) {
			if (x.type === "hidden" && x.name) names.append(x.name, x.value);
		});
		files.forEach(function (f) {
			names.append("name", nameOf(f));
		});
		var go = function () {
			form.__tcAsked = true;
			if (form.requestSubmit) form.requestSubmit(by || undefined);
			else form.submit();
		};
		fetch(new URL(input.getAttribute("data-tc-taken"), location.href).href, { method: "POST", body: names, credentials: "same-origin" })
			.then(function (r) {
				if (!r.ok) throw new Error("taken " + r.status);
				return r.text();
			})
			.then(function (text) {
				var taken = text.split(/\r?\n/).map(function (s) {
					return s.trim();
				});
				var asked = files.filter(function (f) {
					return taken.indexOf(nameOf(f)) >= 0;
				});
				var tpl = line && line.querySelector(":scope > template[data-tc-taken]"), list = line && line.querySelector(":scope > .tc-queue__list");
				if (!asked.length || !tpl || !tpl.content.firstElementChild || !list) return go();
				line.hidden = false;
				line.style.display = "";
				var left = asked.length, skipped = [];
				asked.forEach(function (f) {
					var row = tpl.content.firstElementChild.cloneNode(true);
					tc.fillWords(row, { "{name}": nameOf(f) });
					list.appendChild(row);
					row.addEventListener("click", function (ev) {
						var b = ev.target.closest && ev.target.closest("button[value]");
						if (!b) return;
						if (b.value === "skip") skipped.push(f);
						else {
							var h = document.createElement("input");
							h.type = "hidden";
							h.name = "taken." + nameOf(f);
							h.value = b.value;
							h.setAttribute("data-tc-answer", "");
							// before the files, as their sizes: a server that
							// reads the body as a stream knows what to do with a
							// file when it comes (Carrel's task 219)
							if (input.hasAttribute("form")) h.setAttribute("form", input.getAttribute("form"));
							input.parentNode.insertBefore(h, input);
						}
						row.remove();
						if (--left) return;
						if (skipped.length) {
							var keep = new DataTransfer();
							files.forEach(function (x) {
								if (skipped.indexOf(x) < 0) keep.items.add(x);
							});
							input.files = keep.files;
						}
						if (input.files.length) go();
						else if (!list.children.length) line.hidden = true;
					});
				});
			})
			.catch(go);
	}

	document.addEventListener("submit", function (e) {
		var form = e.target, by = e.submitter || null;
		if (e.defaultPrevented || sentPlain || !window.FormData || !window.XMLHttpRequest) return;
		var line = form.id ? document.querySelector('[data-tc-progress="' + form.id + '"]') : null;
		var action = (by && by.getAttribute("formaction")) || form.action;
		var target = targetOf(by) || targetOf(form), zone = target ? zoneOf(target) : null;
		if (zone && (!shows(zone) || !sameSite(action))) zone = null;
		if (!line && !zone) return;
		var asker = Array.prototype.filter.call(form.elements, function (x) {
			return x.type === "file" && x.hasAttribute("data-tc-taken") && x.files && x.files.length;
		})[0];
		if (asker && !form.__tcAsked && window.fetch) {
			e.preventDefault();
			askTaken(form, by, asker, line);
			return;
		}
		form.__tcAsked = false;
		var data = new FormData(form, by || undefined);
		var files = [];
		data.forEach(function (v) {
			if (v && typeof v === "object" && v.size) files.push(v);
		});
		if (!files.length) return;
		e.preventDefault();
		var url = new URL(action, location.href).href;
		var queue = line && line.classList.contains("tc-queue") ? queued(line, files) : null;
		var bar = line && !queue && line.querySelector(".tc-meter__bar"), words = line && !queue && line.querySelector(".tc-mono");
		if (line) {
			line.hidden = false;
			line.style.display = "";
		}
		if (queue) queue(0);
		var buttons = Array.prototype.filter.call(document.querySelectorAll('[form="' + form.id + '"], [type="submit"]'), function (b) {
			return b.form === form && !b.disabled;
		});
		buttons.forEach(function (b) {
			b.disabled = true;
		});
		var mb = function (b) {
			return (b / 1024 / 1024).toFixed(1);
		};
		var x = new XMLHttpRequest();
		x.open((form.getAttribute("method") || "get").toUpperCase(), url);
		if (zone) {
			var heads = headers(target, data);
			Object.keys(heads).forEach(function (k) {
				x.setRequestHeader(k, heads[k]);
			});
		}
		x.upload.onprogress = function (p) {
			if (!p.lengthComputable) return;
			if (queue) queue(p.loaded / p.total);
			if (bar) bar.style.width = Math.round((p.loaded / p.total) * 100) + "%";
			if (words) words.textContent = mb(p.loaded) + " of " + mb(p.total) + " MB";
		};
		x.onload = function () {
			if (x.responseURL && x.responseURL !== url) {
				location.assign(x.responseURL);
				return;
			}
			var type = x.getResponseHeader("Content-Type") || "";
			if (zone && x.status >= 200 && x.status < 300 && type.indexOf("text/html") === 0) {
				// the files are sent: the form is as it was, for the next ones
				if (queue) queue(1);
				buttons.forEach(function (b) {
					b.disabled = false;
				});
				tc.unbusy(form);
				Array.prototype.forEach.call(form.elements, function (f) {
					if (f.type === "file") f.value = "";
				});
				Array.prototype.forEach.call(form.querySelectorAll("[data-tc-answer]"), function (h) {
					h.remove();
				});
				if (line && !queue) line.hidden = true;
				if (put(target, x.responseText)) return;
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
	// A form asked again as it changes (data-tc-recount, Form.Recount;
	// Carrel's task 194): a field that changes asks the form's address with
	// its fields, by GET, and puts the elements it names in place — the line
	// at Attach says what the chosen days weigh before Attach. A field that
	// is off is not asked with, as it is not sent. When the answer fails the
	// line stays, and the server says it after the form is sent.
	document.addEventListener("change", function (e) {
		var f = e.target.form;
		if (!f || !f.hasAttribute("data-tc-recount") || !window.DOMParser || !window.FormData || !window.URL) return;
		var ids = f.getAttribute("data-tc-recount").split(/\s+/).filter(Boolean);
		if (!ids.length) return;
		var url = new URL(f.getAttribute("action") || location.href, location.href);
		var data = new FormData(f);
		data.forEach(function (v, k) {
			url.searchParams.delete(k);
		});
		data.forEach(function (v, k) {
			if (typeof v === "string") url.searchParams.append(k, v);
		});
		swap(url.href, ids);
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

	// The page's forms an answer carries after the widget of its zone
	// (Carrel's task 226): a button outside a form names its form by ID, and
	// the page prints such forms after its root. An answer carries them in a
	// template[data-tc-forms], so a zone's content that is itself a form stays
	// the zone's, and an answer of forms only empties its zone. Each takes the
	// place of the page's form of the same ID, or joins the end of the page.
	function forms(content) {
		var out = [], boxes = [];
		Array.prototype.forEach.call(content.children, function (t) {
			if (t.tagName === "TEMPLATE" && t.hasAttribute("data-tc-forms")) boxes.push(t);
		});
		boxes.forEach(function (t) {
			content.removeChild(t);
			Array.prototype.forEach.call(t.content.children, function (f) {
				if (f.tagName === "FORM" && f.id) out.push(f);
			});
		});
		return out;
	}

	function place(list) {
		list.forEach(function (f) {
			var old = document.getElementById(f.id);
			if (old && old.tagName === "FORM") old.parentNode.replaceChild(f, old);
			else document.body.appendChild(f);
		});
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
		var posts = forms(box.content);
		fill(zone, box.content);
		extra.forEach(function (t) {
			var z = zoneOf(t.getAttribute("data-tc-zone"));
			if (!z) return;
			posts = posts.concat(forms(t.content));
			fill(z, t.content);
		});
		place(posts);
		if (tc.layoutUpdates) tc.layoutUpdates();
		follow();
		return true;
	}

	// What a request for a part of the page says: that a part is wanted,
	// which zone, and the form's token in the header the document names.
	function headers(target, body) {
		var h = { Accept: "text/html" };
		var view = doc.getAttribute("data-tc-view-header"), zoneHeader = doc.getAttribute("data-tc-target-header");
		if (view) h[view] = doc.getAttribute("data-tc-view-value") || "1";
		if (zoneHeader) h[zoneHeader] = target;
		var field = doc.getAttribute("data-tc-token-field"), tokenHeader = doc.getAttribute("data-tc-token-header");
		if (body && field && tokenHeader && body.get(field)) h[tokenHeader] = body.get(field);
		return h;
	}

	// A request that is not GET is sent once (Carrel's note 58): when its
	// answer has come, whatever it is, the form is not sent again the plain
	// way — that would write twice. An address it was sent on to is gone to,
	// any other answer is shown whole, as the plain way shows it. Only a
	// request that had no answer at all goes the plain way.
	function ask(url, method, body, target, plain, then) {
		var once = method !== "GET", answered = false;
		var whole = function (r, text) {
			if (r.redirected) location.assign(r.url);
			else {
				document.open();
				document.write(text);
				document.close();
			}
		};
		fetch(url, { method: method, body: body || undefined, headers: headers(target, body), credentials: "same-origin" })
			.then(function (r) {
				answered = true;
				var type = r.headers.get("Content-Type") || "";
				if (!r.ok || r.redirected || type.indexOf("text/html") !== 0) {
					if (!once) throw new Error("not a part");
					return r.text().then(function (text) {
						whole(r, text);
						return null;
					});
				}
				return r.text().then(function (html) {
					if (once && !put(target, html)) {
						whole(r, html);
						return null;
					}
					return html;
				});
			})
			.then(function (html) {
				if (html === null) return;
				if (!once && !put(target, html)) throw new Error("no zone " + target);
				if (method === "GET" && doc.getAttribute("data-tc-address-zone") === target) {
					history.pushState({ tcSwap: true }, "", url);
					swapped = true;
				}
				if (then) then();
			})
			.catch(function () {
				if (!(once && answered)) plain();
			});
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

	// A run's line that is a zone follows its run (Carrel's task 85). It
	// names a stream of events (data-tc-stream) and the events' names: each
	// event carries a fragment, put as an answer is; on the event that ends
	// the run the stream is closed. With no stream, or one that refused or
	// broke, the line's poll address (data-tc-poll) is asked every
	// data-tc-every milliseconds for as long as the line that comes back
	// names one and is shown. One run is followed once, whatever comes to
	// replace its line. A hidden line — empty, waiting for a run, or emptied
	// by the event that ended it — is not followed (Carrel's task 226).
	var runs = {};

	function poll(name) {
		var bar = zoneOf(name);
		var url = bar && !bar.hidden && bar.getAttribute("data-tc-poll");
		if (!url || runs[name]) return;
		var every = parseFloat(bar.getAttribute("data-tc-every")) || 5000;
		runs[name] = {
			timer: setTimeout(function () {
				var again = function () {
					delete runs[name];
					poll(name);
				};
				ask(url, "GET", null, name, again, again);
			}, every)
		};
	}

	function follow() {
		Array.prototype.forEach.call(document.querySelectorAll(".tc-runbar[data-tc-zone]"), function (bar) {
			var name = bar.getAttribute("data-tc-zone"), stream = bar.getAttribute("data-tc-stream");
			if (runs[name] || bar.hidden) return;
			if (!stream || !window.EventSource) {
				poll(name);
				return;
			}
			var source = new EventSource(stream);
			var over = function () {
				source.close();
				delete runs[name];
			};
			runs[name] = { source: source };
			source.addEventListener(bar.getAttribute("data-tc-changed") || "message", function (e) {
				put(name, e.data);
			});
			if (bar.getAttribute("data-tc-ended")) {
				source.addEventListener(bar.getAttribute("data-tc-ended"), function (e) {
					over();
					put(name, e.data);
				});
			}
			source.onerror = function () {
				over();
				poll(name);
			};
		});
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
		if (files) return; // a form that sends files is sent above, or loads its answer whole
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
		}, function () {
			tc.unbusy(form);
		});
	});

	if (document.readyState === "loading") document.addEventListener("DOMContentLoaded", follow);
	else follow();

	document.addEventListener("input", function (e) {
		var t = e.target;
		if (t.classList && t.classList.contains("tc-in__add") && t.closest("[data-tc-suggest]")) suggest(t);
	});
})();
