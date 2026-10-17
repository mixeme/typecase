package widget

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// slugCyrillic is the closed transliteration table of Slug. It runs after
// NFD, so й, ё, ї and ў arrive as a base letter and a combining mark, and the
// mark falls out with everything else outside [a-z0-9-]. js/tc.js holds the
// same table (slugOf); TestSlugAgrees keeps the two the same.
var slugCyrillic = map[rune]string{
	'а': "a", 'б': "b", 'в': "v", 'г': "g", 'д': "d", 'е': "e", 'ж': "zh", 'з': "z",
	'и': "i", 'к': "k", 'л': "l", 'м': "m", 'н': "n", 'о': "o", 'п': "p", 'р': "r",
	'с': "s", 'т': "t", 'у': "u", 'ф': "f", 'х': "kh", 'ц': "ts", 'ч': "ch", 'ш': "sh",
	'щ': "shch", 'ъ': "", 'ы': "y", 'ь': "", 'э': "e", 'ю': "yu", 'я': "ya",
	'і': "i", 'є': "ye", 'ґ': "g",
}

// SlugEmpty is the slug of a name that leaves nothing behind.
const SlugEmpty = "col"

// Slug derives an address segment from a name — a new collection's address
// from its name (Carrel's task 200, its architecture.md): NFD, the
// transliteration table, lower case, a run of spaces or hyphens to one
// hyphen, then [a-z0-9-] only; empty becomes SlugEmpty. The library's script
// derives the same as the name is typed (Input.From), so a server and a page
// never spell it two ways.
func Slug(name string) string {
	var b strings.Builder
	pending := false
	for _, r := range norm.NFD.String(name) {
		r = unicode.ToLower(r)
		if unicode.IsSpace(r) || r == '-' {
			pending = b.Len() > 0
			continue
		}
		out, ok := slugCyrillic[r]
		if !ok {
			out = string(r)
		}
		for _, c := range out {
			if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
				if pending {
					b.WriteByte('-')
					pending = false
				}
				b.WriteRune(c)
			}
		}
	}
	if b.Len() == 0 {
		return SlugEmpty
	}
	return b.String()
}

// SlugOf is a piece of a hint that says the address a field derives
// (Input.From): what is typed in the field Of, else its placeholder, else
// «…»; the library's script keeps it so at every keystroke. Text is what the
// server prints.
type SlugOf struct {
	Of   string
	Text string
}

func (s SlugOf) html(mix) string {
	return `<span class="tc-mono"` + attr("data-tc-slug-of", s.Of) + `>` + esc(s.Text) + `</span>`
}

func (SlugOf) inline() {}
