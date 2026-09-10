(function () {
	"use strict";

	// One behavior file for the whole library (L-6.1). Widgets mark nodes with
	// data-tc-*; handlers attach here as widgets are added.
	document.addEventListener("DOMContentLoaded", function () {
		var roots = document.querySelectorAll("[data-tc-root]");
		for (var i = 0; i < roots.length; i++) {
			roots[i].classList.add("tc-root");
		}
	});
})();
