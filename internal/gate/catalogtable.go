package gate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	tableRe     = regexp.MustCompile(`(?s)<table\b[^>]*class="tc-table[" ][^>]*>(.*?)</table>`)
	tableHeadRe = regexp.MustCompile(`(?s)<thead><tr>(.*?)</tr></thead>`)
	tableRowRe  = regexp.MustCompile(`(?is)<tr\b[^>]*>(.*?)</tr>`)
	thRe        = regexp.MustCompile(`<th\b`)
	tdRe        = regexp.MustCompile(`<td\b[^>]*>`)
	colspanRe   = regexp.MustCompile(`colspan="(\d+)"`)
)

// CheckTableColumns requires that no row of a tc-table is wider than its
// head: a cell spanning columns counts as that many. The frames draw tables
// without a head (a list of values) and rows that stop short (a row that no
// longer applies), so neither is refused.
func CheckTableColumns(catalog []CatalogEntry) error {
	var bad []string
	for _, e := range catalog {
		for ti, t := range tableRe.FindAllStringSubmatch(e.HTML, -1) {
			head := tableHeadRe.FindStringSubmatch(t[1])
			if head == nil {
				continue
			}
			cols := len(thRe.FindAllString(head[1], -1))
			body := t[1][strings.Index(t[1], "<tbody>"):]
			for i, row := range tableRowRe.FindAllStringSubmatch(body, -1) {
				cells := 0
				for _, td := range tdRe.FindAllString(row[1], -1) {
					n := 1
					if m := colspanRe.FindStringSubmatch(td); m != nil {
						n, _ = strconv.Atoi(m[1])
					}
					cells += n
				}
				if cells > cols {
					bad = append(bad, fmt.Sprintf("%s/%s: table %d, row %d has %d cells, %d headers", e.Widget, e.State, ti+1, i+1, cells, cols))
				}
			}
		}
	}
	if len(bad) > 0 {
		return fmt.Errorf("table-columns: %s", strings.Join(bad, "; "))
	}
	return nil
}
