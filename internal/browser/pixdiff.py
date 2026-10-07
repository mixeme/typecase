"""Which frame shots differ, and by how many pixels.

    python internal/browser/pixdiff.py before_dir after_dir

The shots are TestMockupFrameShots' (TC_FRAME_SHOTS, TC_FRAME_SHOTS_DIR;
TC_FRAME_SHOTS_REV=HEAD for «before»). Needs Pillow.
"""
import os, sys
from PIL import Image, ImageChops
b, a = sys.argv[1], sys.argv[2]
for n in sorted(os.listdir(a)):
    pa, pb = os.path.join(a, n), os.path.join(b, n)
    if not os.path.exists(pb):
        print("new", n); continue
    ia, ib = Image.open(pa).convert("RGB"), Image.open(pb).convert("RGB")
    if ia.size != ib.size:
        print("size", n, ib.size, "->", ia.size); continue
    d = ImageChops.difference(ia, ib)
    box = d.getbbox()
    if box:
        px = sum(1 for p in d.get_flattened_data() if p != (0, 0, 0))
        print("diff", n, px, box)
