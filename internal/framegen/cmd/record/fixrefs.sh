#!/bin/bash
# Redraft the reference records the tests name red, and print what is left
# (a form without a record needs add.py). Run from the module root.
# A record whose fragment only moved is renumbered, not redrafted; a redraft
# that would change a record's variant or state is refused (refresh.py).
for round in 1 2 3; do
  go test ./internal/refs > "$TEMP/refs.txt" 2>&1
  PYTHONIOENCODING=utf-8 python internal/framegen/cmd/record/moved.py < "$TEMP/refs.txt"
  go test ./internal/refs 2>&1 | grep -o 'TestWidgetFragments/[^ ]*' | sort -u > "$TEMP/red.txt"
  [ -s "$TEMP/red.txt" ] || break
  args=$(sed -E 's#.*/([a-z0-9-]+)_\.([a-z_-]+)_\#([0-9]+)#\1:\2:\3#' "$TEMP/red.txt" | tr '\n' ' ')
  PYTHONIOENCODING=utf-8 python internal/framegen/cmd/record/refresh.py $args 2>&1 | grep -v "^redrafted"
  gofmt -w internal/refs
done
go test ./internal/refs 2>&1 | grep -v '^ok' | cut -c1-500 | head -12
