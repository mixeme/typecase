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

## Catalog and screens

```bash
go run ./cmd/catalog   # http://localhost:8080/
go run ./cmd/screens   # http://localhost:8081/
```

## Specification

See [docs/spec/typecase.md](docs/spec/typecase.md).

The look comes from the mockup; its rules and checks are in
[docs/visual/mockup.md](docs/visual/mockup.md).

## License

MIT
