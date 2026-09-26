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

## Specification

See [docs/spec/typecase.md](docs/spec/typecase.md).

The look comes from the mockup; its rules and checks are in
[docs/visual/mockup.md](docs/visual/mockup.md).

## License

MIT
