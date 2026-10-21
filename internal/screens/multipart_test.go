package screens

import (
	"strings"
	"testing"

	w "carrel.dev/typecase/widget"
)

// TestRestoreSendsTheFile (Carrel's task 218): the restore from this
// computer is a dialog with a file field — it sends multipart, without a
// script too, and the archive is its card's last field, after the storage
// and the folder: the parts go in the form's order. A post form without a
// file field of its own sends multipart when asked (the answer to a taken
// name).
func TestRestoreSendsTheFile(t *testing.T) {
	page := Backup15().Render()
	open := page[strings.Index(page, `<form class="tc-dialog`):]
	open = open[:strings.Index(open, ">")]
	if !strings.Contains(open, `enctype="multipart/form-data"`) {
		t.Errorf("the restore's dialog: %s; want multipart", open)
	}
	store, folder, file := strings.Index(page, `name="store"`), strings.Index(page, `name="folder"`), strings.Index(page, `name="file"`)
	if store < 0 || folder < 0 || file < 0 || !(store < folder && folder < file) {
		t.Errorf("the order of store, folder, file: %d %d %d; want the file last", store, folder, file)
	}
	if got := (w.PostForm{ID: "taken", Action: "/x", Multipart: true}).Render(); !strings.Contains(got, `enctype="multipart/form-data"`) || strings.Contains(got, `type="file"`) {
		t.Errorf("a multipart post form: %s; want enctype and no file field", got)
	}
}
