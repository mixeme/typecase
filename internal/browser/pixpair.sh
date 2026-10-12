#!/bin/bash
# Shots of frames before and after, compared by pixels (docs/development.md, Р-52, Carrel's task 149).
#
#   bash internal/browser/pixpair.sh <frame>...
#
# Each frame is shot at 1351, 820 and 460 px from HEAD and from the working
# tree, one frame per run of the browser; pixdiff.py then names the shots
# that differ. Nothing printed between «shots» and «pixdiff done» means the
# frames look the same. Run from the module root, with no other browser test
# beside it. The shots of HEAD are kept in $TEMP/pix_b and reused until it is
# removed: remove it after a commit.
BD="$TEMP/pix_b"; AD="$TEMP/pix_a"; mkdir -p "$BD" "$AD"; rm -f "$AD"/*
B="$(cygpath -w "$BD")"; A="$(cygpath -w "$AD")"
for f in "$@"; do for w in 1351 820 460; do
  sh="$f@$w"
  if [ ! -f "$BD/$sh.png" ]; then for try in 1 2 3; do TC_FRAME_SHOTS=$sh TC_FRAME_SHOTS_REV=HEAD TC_FRAME_SHOTS_DIR="$B" timeout 200 go test ./internal/browser -run '^TestMockupFrameShots$' -timeout 190s 2>&1 | tail -1 | grep -q '^ok' && break; done; fi
  for try in 1 2 3; do TC_FRAME_SHOTS=$sh TC_FRAME_SHOTS_DIR="$A" timeout 200 go test ./internal/browser -run '^TestMockupFrameShots$' -timeout 190s 2>&1 | tail -1 | grep -q '^ok' && break; done
  [ -f "$AD/$sh.png" ] || echo "NO SHOT $sh"
done; done
echo "shots: $(ls "$AD" | wc -l)"
python internal/browser/pixdiff.py "$B" "$A"
echo "pixdiff done"
