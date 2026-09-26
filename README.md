# Typecase

Server-rendered widget library with typed inputs, bundled styles, and behavior scripts.

## Install

```bash
go get carrel.dev/typecase
```

## Public API

- `CSS()` — concatenated stylesheet
- `JS()` — library behavior script (`tc.js`)
- `ThemeJS()` — theme boot script (load in `<head>` before first paint)
- `Fonts()` — embedded font files
- `Icons()` / `Sprite()` — neutral icon sprite

## Catalog

```bash
go run ./cmd/catalog   # http://localhost:8080/
```

## Demo screens

```bash
go run ./cmd/screens   # http://localhost:8085/
```

Every frame of the mockup that is a route (154) is a screen at `/f/<data-frame>`,
built from widgets with the frame's data; `go test ./internal/screens` holds
each one to its frame. Frames that no widget builds yet are listed in
[docs/visual/gaps.md](docs/visual/gaps.md).

The screens are generated from the mockup (`go generate ./internal/screens`).
To start a screen of your own from a frame:

```bash
go run ./cmd/frame2go s-find-2   # prints the frame as a screen function
```

How to build a screen — [docs/guide/screen.md](docs/guide/screen.md).

## Specification

See [docs/spec/typecase.md](docs/spec/typecase.md).

The look comes from the mockup,
[docs/visual/carrel-ui-mockups.html](docs/visual/carrel-ui-mockups.html):
its style block is cut into `css/typecase.css` (`go generate ./css`), each
widget prints the fragment of its frame, and a screen is a frame built of
widgets. Something the mockup does not draw is drawn there first. The
mockup's rules and checks are in [docs/visual/mockup.md](docs/visual/mockup.md);
how to build a screen — [docs/guide/screen.md](docs/guide/screen.md); the
gates the library keeps — [L-8 of the spec](docs/spec/typecase.md#l-8-ворота).

## License

MIT
