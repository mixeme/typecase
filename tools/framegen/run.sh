#!/bin/sh
# Regenerate internal/screens from the mockup and list what still differs
# from the frames (plan 3, step 1). Run from anywhere: sh tools/framegen/run.sh
# MAXD — differences listed per frame (40), TOP — categories shown (40).
HERE=$(cd "$(dirname "$0")" && pwd)
REPO=$(cd "$HERE/../.." && pwd)
OUT=${FRAMEGEN_OUT:-$HERE/out}
mkdir -p "$OUT"
export FRAMEGEN_OUT="$OUT" PYTHONIOENCODING=utf-8
python "$HERE/genscreens.py" > "$OUT/gen.log" 2>&1 || { tail -20 "$OUT/gen.log"; exit 1; }
grep -v "^GAP" "$OUT/gen.log" | tail -3
python "$HERE/emitscreens.py" || exit 1
cd "$REPO" || exit 1
gofmt -w internal/screens || exit 1
go vet ./internal/screens 2>&1 | head -20
TC_FRAME_DIFFS=${MAXD:-40} go test ./internal/screens -run TestFramesDiffs -v -count=1 2>&1 |
	grep "diffs_test" | sed 's/.*diffs_test.go:[0-9]*: //' > "$OUT/diffs.txt"
go test ./internal/screens -run TestFramesStatic -count=1 2>&1 | grep -c "static_test.go" | sed 's/^/failing frames: /'
wc -l < "$OUT/diffs.txt" | sed 's/^/differences: /'
python "$HERE/cats.py" "$OUT/diffs.txt" | head -"${TOP:-40}"
