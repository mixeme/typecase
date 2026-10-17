(function () {
	"use strict";

	var storageKey = "tc-theme";

	function readStored() {
		try {
			return localStorage.getItem(storageKey);
		} catch (e) {
			return null;
		}
	}

	function systemDark() {
		return (
			window.matchMedia &&
			window.matchMedia("(prefers-color-scheme: dark)").matches
		);
	}

	function resolveTheme() {
		try {
			var params = new URLSearchParams(window.location.search);
			var q = params.get("theme");
			if (q === "light" || q === "dark") {
				return q;
			}
		} catch (e) {}
		var stored = readStored();
		if (stored === "light" || stored === "dark") {
			return stored;
		}
		return systemDark() ? "dark" : "light";
	}

	// The server chose (page.Document.Theme): light or dark is printed
	// already; auto is the system's. What this browser remembers is not read.
	var root = document.documentElement;
	var set = root.getAttribute("data-tc-theme-set");
	if (set === "light" || set === "dark") return;
	if (set === "auto") {
		root.setAttribute("data-tc-theme", systemDark() ? "dark" : "light");
		return;
	}
	root.setAttribute("data-tc-theme", resolveTheme());
})();
