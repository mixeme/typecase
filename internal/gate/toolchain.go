package gate

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// CheckGofmt fails when gofmt -l reports any file under root. dev/ is left
// out: the working copy's scratch, with the Go cache and its temporary files.
func CheckGofmt(root string) error {
	entries, err := os.ReadDir(root)
	if err != nil {
		return fmt.Errorf("gofmt: %w", err)
	}
	args := []string{"-l"}
	for _, e := range entries {
		// gofmt parses a file it is named whatever its extension.
		if e.Name() == "dev" || e.Name() == ".git" || !e.IsDir() && filepath.Ext(e.Name()) != ".go" {
			continue
		}
		args = append(args, filepath.Join(root, e.Name()))
	}
	out, err := exec.Command("gofmt", args...).Output()
	if err != nil {
		return fmt.Errorf("gofmt: %w", err)
	}
	broken := string(filepath.Separator) + "testdata" + string(filepath.Separator) + "broken" + string(filepath.Separator)
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	var bad []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.Contains(line, broken) {
			continue
		}
		bad = append(bad, line)
	}
	if len(bad) > 0 {
		return fmt.Errorf("gofmt: %s", strings.Join(bad, ", "))
	}
	return nil
}

// CheckVet runs go vet ./... from the module root.
func CheckVet(root string) error {
	cmd := exec.Command("go", "vet", "./...")
	cmd.Dir = root
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		msg := strings.TrimSpace(stderr.String())
		if msg == "" {
			msg = err.Error()
		}
		return fmt.Errorf("vet: %s", msg)
	}
	return nil
}

// ListGofmtDirty returns paths gofmt -l would rewrite under root.
func ListGofmtDirty(root string) ([]string, error) {
	out, err := exec.Command("gofmt", "-l", root).Output()
	if err != nil {
		return nil, err
	}
	var paths []string
	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			if !filepath.IsAbs(line) {
				line = filepath.Join(root, line)
			}
			paths = append(paths, line)
		}
	}
	return paths, nil
}

// ConsumerDeps are the modules outside the standard library the packages a
// consumer imports may pull in (L-11.1, owner 28.10): x/text for Slug's
// NFD, x/net/html for page.Check's parse.
var ConsumerDeps = []string{"golang.org/x/text/", "golang.org/x/net/html"}

// CheckConsumerDeps fails when a package of the module outside internal and
// cmd depends on a non-standard package of another module that no prefix in
// allowed covers.
func CheckConsumerDeps(root string, allowed []string) error {
	golist := func(args ...string) ([]string, error) {
		cmd := exec.Command("go", append([]string{"list"}, args...)...)
		cmd.Dir = root
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		out, err := cmd.Output()
		if err != nil {
			return nil, fmt.Errorf("consumer deps: %v %s", err, strings.TrimSpace(stderr.String()))
		}
		return strings.Fields(string(out)), nil
	}
	mod, err := golist("-m")
	if err != nil {
		return err
	}
	all, err := golist("./...")
	if err != nil {
		return err
	}
	var consumer []string
	for _, p := range all {
		if !strings.Contains(p+"/", "/internal/") && !strings.Contains(p+"/", "/cmd/") {
			consumer = append(consumer, p)
		}
	}
	deps, err := golist(append([]string{"-deps", "-f", "{{if not .Standard}}{{.ImportPath}}{{end}}"}, consumer...)...)
	if err != nil {
		return err
	}
	var bad []string
	for _, p := range deps {
		if p == mod[0] || strings.HasPrefix(p, mod[0]+"/") {
			continue
		}
		ok := false
		for _, a := range allowed {
			ok = ok || strings.HasPrefix(p, a)
		}
		if !ok {
			bad = append(bad, p)
		}
	}
	if len(bad) > 0 {
		return fmt.Errorf("consumer deps: %s is not in ConsumerDeps (L-11.1: a new one is a decision)", strings.Join(bad, ", "))
	}
	return nil
}
