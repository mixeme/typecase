package gate

import "testing"

func TestCheckRefValid_rejectsSmoke(t *testing.T) {
	if checkRefValid("widget:smoke") {
		t.Fatal("widget:smoke must not be a valid invariant binding")
	}
	if checkRefValid("todo:pending") {
		t.Fatal("todo: must not be accepted")
	}
	if !checkRefValid("gate:inv-avatar-not-empty") {
		t.Fatal("inv gate checks must be accepted")
	}
}

func TestBadgeNotLinkLooksAtTheBadge(t *testing.T) {
	ok := `<div class="tc-nav"><a class="tc-nav__item" href="#">Files</a></div><span class="tc-badge">ro</span>`
	if err := checkInvBadgeNotLink("tc-badge", []string{ok}, ""); err != nil {
		t.Errorf("a badge beside a link is refused: %v", err)
	}
	bad := `<a class="tc-badge is-linked" href="#">linked</a>`
	if err := checkInvBadgeNotLink("tc-badge", []string{bad}, ""); err == nil {
		t.Error("a badge that is a link passes")
	}
}
