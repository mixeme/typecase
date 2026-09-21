package proof

import (
	"fmt"
	"os/exec"
	"strings"
)

// MockupRev is the last commit whose mockup still speaks the m-* dictionary
// the proof translates. After it the mockup itself is written in tc-* with
// real controls and no style="", so there is nothing left to translate: the
// proof stays evidence about that revision, not about the working tree.
const MockupRev = "894fa046b259b22b072975fa9cb5b3ef05f2b26d"

const mockupRel = "docs/visual/carrel-ui-mockups.html"

// LoadMockup returns the mockup at MockupRev from the repository at root.
func LoadMockup(root string) (string, error) {
	out, err := exec.Command("git", "-C", root, "show", MockupRev+":"+mockupRel).Output()
	if err != nil {
		stderr := ""
		if ee, ok := err.(*exec.ExitError); ok {
			stderr = strings.TrimSpace(string(ee.Stderr))
		}
		return "", fmt.Errorf("proof: git show %s:%s: %v (%s)", MockupRev[:7], mockupRel, err, stderr)
	}
	return string(out), nil
}
