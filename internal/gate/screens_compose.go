package gate

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"carrel.dev/typecase/internal/screens"
)

var screensComposeBypass = []string{
	`<main class="tc-main"`,
	`<div class="tc-shell"`,
	`<div class="tc-bar"`,
	`<aside class="tc-side"`,
}

// screensForbiddenImports are library packages a screen must not reach past widget.
var screensForbiddenImports = []string{
	"carrel.dev/typecase/layout",
}

func forbiddenScreenImport(path string, src []byte) string {
	f, err := parser.ParseFile(token.NewFileSet(), path, src, parser.ImportsOnly)
	if err != nil {
		return ""
	}
	for _, imp := range f.Imports {
		p := strings.Trim(imp.Path.Value, `"`)
		for _, bad := range screensForbiddenImports {
			if p == bad {
				return p
			}
		}
	}
	return ""
}

// CheckScreensCompose rejects hand-built shell markup in demo screen assemblers.
func CheckScreensCompose(p ModulePaths) error {
	screensDir := filepath.Join(p.Root, "internal", "screens")
	entries, err := os.ReadDir(screensDir)
	if err != nil {
		return fmt.Errorf("screens-compose: read screens: %w", err)
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") {
			continue
		}
		path := filepath.Join(screensDir, e.Name())
		src, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("screens-compose: read %s: %w", e.Name(), err)
		}
		if imp := forbiddenScreenImport(path, src); imp != "" {
			return fmt.Errorf("screens-compose: %s imports %s; screens call widget only", e.Name(), imp)
		}
		switch e.Name() {
		case "fixtures.go", "serve.go", "shell_common.go", "style.go", "widgets.go":
			continue
		}
		text := string(src)
		var found []string
		for _, pat := range screensComposeBypass {
			if strings.Contains(text, pat) {
				found = append(found, pat)
			}
		}
		if len(found) > 0 {
			return fmt.Errorf("screens-compose: %s bypasses widgets: %s", e.Name(), strings.Join(found, ", "))
		}
	}
	html := screens.Notes()
	if strings.Contains(html, "tc-row--excerpt") {
		if !strings.Contains(html, "tc-row__date") || !strings.Contains(html, "tc-row__source") {
			return fmt.Errorf("screens-compose: notes excerpt rows must include date and source cells")
		}
	}
	contacts := screens.Contacts()
	if !strings.Contains(contacts, "tc-row--avatarcols") {
		return fmt.Errorf("screens-compose: contacts list must use avatarcols")
	}
	if !strings.Contains(contacts, "tc-row__subtitle") || !strings.Contains(contacts, "tc-row__detail") || !strings.Contains(contacts, "tc-row__source") {
		return fmt.Errorf("screens-compose: contacts avatarcols rows must include subtitle, detail and source cells")
	}
	files := screens.Files()
	if !strings.Contains(files, "tc-row--tickcols") {
		return fmt.Errorf("screens-compose: files list must use tickcols")
	}
	agenda := screens.Agenda()
	if !strings.Contains(agenda, "tc-row--timecols") {
		return fmt.Errorf("screens-compose: agenda list must use timecols")
	}
	if !strings.Contains(agenda, "tc-tickrow") {
		return fmt.Errorf("screens-compose: agenda rail must use tickrow")
	}
	if !strings.Contains(agenda, "tc-fields") {
		return fmt.Errorf("screens-compose: agenda side must use fields")
	}
	if !strings.Contains(agenda, `data-value="week"`) || !strings.Contains(agenda, "is-chosen") {
		return fmt.Errorf("screens-compose: agenda span must be week")
	}
	tasks := screens.Tasks()
	if !strings.Contains(tasks, "tc-row--tickmarkcols") {
		return fmt.Errorf("screens-compose: tasks list must use tickmarkcols")
	}
	find := screens.Find()
	if !strings.Contains(find, "tc-row--find") {
		return fmt.Errorf("screens-compose: find hits must use row find")
	}
	if strings.Contains(find, "tc-search__popover") {
		return fmt.Errorf("screens-compose: find must be the hit list, not the search popover")
	}
	list := screens.Notes()
	if strings.Contains(list, "tc-doc") {
		return fmt.Errorf("screens-compose: notes list must not include tc-doc preview")
	}
	read := screens.NotesRead()
	if !strings.Contains(read, "tc-tree") {
		return fmt.Errorf("screens-compose: notes-read must include tc-tree neighbours")
	}
	quick := screens.NotesQuick()
	if !strings.Contains(quick, "tc-sheet is-open") {
		return fmt.Errorf("screens-compose: notes-quick must have open sheet")
	}
	edit := screens.NotesEdit()
	if strings.Contains(edit, "tc-shell__rail") {
		return fmt.Errorf("screens-compose: notes-edit must not include rail column")
	}
	if strings.Contains(edit, "tc-side") {
		return fmt.Errorf("screens-compose: notes-edit must not include side panel")
	}
	if !strings.Contains(list, "tc-pfieldrow") || !strings.Contains(list, `class="tc-pfield`) {
		return fmt.Errorf("screens-compose: notes side must use panel fields")
	}
	if !strings.Contains(read, "tc-docbar--frame") {
		return fmt.Errorf("screens-compose: notes-read must use one docbar frame")
	}
	if !strings.Contains(read, "tc-docmeta__danger") {
		return fmt.Errorf("screens-compose: notes-read meta must include danger")
	}
	if !strings.Contains(edit, "tc-docedit") {
		return fmt.Errorf("screens-compose: notes-edit must use tc-docedit")
	}
	if !strings.Contains(edit, "tc-pick") {
		return fmt.Errorf("screens-compose: notes-edit must include related pick")
	}
	conflict := screens.Conflict()
	if !strings.Contains(conflict, "tc-diff") {
		return fmt.Errorf("screens-compose: conflict must use diff cards")
	}
	devices := screens.Devices()
	if !strings.Contains(devices, "tc-cols--halves") {
		return fmt.Errorf("screens-compose: devices must assemble from Cols halves")
	}
	admin := screens.Admin()
	if !strings.Contains(admin, "tc-cols--thirds") {
		return fmt.Errorf("screens-compose: admin must assemble from Cols thirds")
	}
	mail := screens.Mail()
	if !strings.Contains(mail, "tc-sheet is-open") {
		return fmt.Errorf("screens-compose: mail must have an open sheet")
	}
	if !strings.Contains(mail, "tc-row--avatarcols") {
		return fmt.Errorf("screens-compose: mail must keep the contact list")
	}
	login := screens.Login()
	if strings.Contains(login, "tc-shell") {
		return fmt.Errorf("screens-compose: login must not use the shell chrome")
	}
	account := screens.Account()
	if !strings.Contains(account, "tc-nav") {
		return fmt.Errorf("screens-compose: account must use settings nav")
	}
	return nil
}
