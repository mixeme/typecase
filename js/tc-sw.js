(function () {
	"use strict";

	// tc-sw.js is the library's worker for a page installed as an
	// application (Carrel's task 91). It keeps only the library's own files —
	// those under the base the page names — and the page the application
	// shows when the network is gone; it keeps no page of the application's
	// and no data. A page that loads goes to the network; when the network
	// does not answer, the offline page is shown. The library's files come
	// from the network too, and the copy kept is for when there is none.
	// tc.js registers it with what it needs in its address: base, offline
	// and files (the stylesheet and the scripts the page loaded).
	var params = new URL(self.location.href).searchParams;
	var base = params.get("base") || "";
	var offline = params.get("offline") || "";
	var files = (params.get("files") || "").split(" ").filter(Boolean);
	var kept = "tc-offline";

	self.addEventListener("install", function (e) {
		e.waitUntil(
			caches
				.open(kept)
				.then(function (c) {
					return c.addAll(offline ? files.concat([offline]) : files);
				})
				.then(function () {
					return self.skipWaiting();
				})
		);
	});

	// A worker of another version takes over: what the one before kept
	// outside this one's base — the library's files of a base gone, an
	// offline page renamed — goes; the files under the base stay, refreshed
	// as they are loaded.
	self.addEventListener("activate", function (e) {
		e.waitUntil(
			caches
				.open(kept)
				.then(function (c) {
					return c.keys().then(function (ks) {
						return Promise.all(
							ks.map(function (k) {
								var p = new URL(k.url).pathname;
								if (p === offline || (base && p.indexOf(base) === 0)) return null;
								return c.delete(k);
							})
						);
					});
				})
				.then(function () {
					return self.clients.claim();
				})
		);
	});

	self.addEventListener("fetch", function (e) {
		var r = e.request;
		if (r.method !== "GET") return;
		var url = new URL(r.url);
		if (url.origin !== self.location.origin) return;
		if (r.mode === "navigate") {
			if (!offline) return;
			e.respondWith(
				fetch(r).catch(function () {
					return caches.match(offline);
				})
			);
			return;
		}
		if (!base || url.pathname.indexOf(base) !== 0) return;
		e.respondWith(
			fetch(r)
				.then(function (res) {
					if (res.ok) {
						var copy = res.clone();
						caches.open(kept).then(function (c) {
							c.put(r, copy);
						});
					}
					return res;
				})
				.catch(function () {
					return caches.match(r).then(function (m) {
						return m || Response.error();
					});
				})
		);
	});
})();
