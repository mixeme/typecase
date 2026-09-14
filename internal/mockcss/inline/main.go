package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root, err := moduleRoot()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	htmlPath := filepath.Join(root, "docs", "visual", "carrel-ui-mockups.html")
	auditPath := filepath.Join(root, "docs", "visual", "inline-audit.md")
	in, err := os.ReadFile(htmlPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	out, rows, err := apply(string(in))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if len(rows) > 0 && out != string(in) {
		if err := os.WriteFile(auditPath, []byte(formatAudit(rows)), 0o644); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	if out == string(in) {
		fmt.Println("inline: no changes")
		if len(rows) > 0 {
			fmt.Printf("inline: wrote %s (%d rows)\n", auditPath, len(rows))
		}
		return
	}
	if err := os.WriteFile(htmlPath, []byte(out), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("inline: wrote %s\n", htmlPath)
	fmt.Printf("inline: wrote %s (%d rows)\n", auditPath, len(rows))
}

func moduleRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dir := wd
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found from %s", wd)
		}
		dir = parent
	}
}
