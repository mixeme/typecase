package gate

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"carrel.dev/typecase/layout"
)

var (
	mockupCSSClass  = regexp.MustCompile(`(?:^|[^A-Za-z0-9_-])\.m-([a-z][a-z0-9_-]*)`)
	mockupHTMLClass = regexp.MustCompile(`\bclass\s*=\s*("([^"]*)"|'([^']*)')`)
	mockupName      = regexp.MustCompile(`^m-[a-z][a-z0-9_-]*$`)
)

// CheckLayoutInventory fails when a mockup .m-* class is missing from layout.Inventory.
func CheckLayoutInventory(p ModulePaths) error {
	path := filepath.Join(p.Root, "docs", "visual", "carrel-ui-mockups.html")
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("layout-inventory: read mockup: %w", err)
	}
	found := mockupMClasses(string(data))
	var missing []string
	for _, name := range found {
		if _, ok := layout.Inventory[name]; !ok {
			missing = append(missing, name)
		}
	}
	if len(missing) > 0 {
		sort.Strings(missing)
		return fmt.Errorf("layout-inventory: undeclared .m-* classes: %s", strings.Join(missing, ", "))
	}
	return nil
}

func mockupMClasses(html string) []string {
	seen := map[string]bool{}
	for _, m := range mockupCSSClass.FindAllStringSubmatch(html, -1) {
		seen["m-"+m[1]] = true
	}
	for _, m := range mockupHTMLClass.FindAllStringSubmatch(html, -1) {
		raw := m[2]
		if raw == "" {
			raw = m[3]
		}
		for _, c := range strings.Fields(raw) {
			if mockupName.MatchString(c) {
				seen[c] = true
			}
		}
	}
	out := make([]string, 0, len(seen))
	for name := range seen {
		out = append(out, name)
	}
	sort.Strings(out)
	return out
}
