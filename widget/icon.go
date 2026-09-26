package widget

// Icon names a symbol of the sprite without its "i-" prefix: "plus" is
// #i-plus. The library ships the mockup's sprite (icons/sprite.svg), but the
// set is the consumer's (L-5.1): the library gives the presentation — the
// wrapper each frame draws — and does not enumerate the symbols.
type Icon string

// use is the reference to the symbol.
func (i Icon) use() string {
	return `<use href="#i-` + esc(string(i)) + `"/>`
}
