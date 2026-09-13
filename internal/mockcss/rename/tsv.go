package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

// Rule is one TSV row: old selector fragment → new.
type Rule struct {
	From string
	To   string
}

func parseTSV(text string) ([]Rule, error) {
	var out []Rule
	sc := bufio.NewScanner(strings.NewReader(text))
	line := 0
	for sc.Scan() {
		line++
		s := strings.TrimSpace(sc.Text())
		if s == "" || strings.HasPrefix(s, "#") {
			continue
		}
		from, to, ok := strings.Cut(s, "\t")
		if !ok {
			return nil, fmt.Errorf("rename.tsv:%d: want two tab-separated columns", line)
		}
		from, to = strings.TrimSpace(from), strings.TrimSpace(to)
		if from == "" || to == "" {
			return nil, fmt.Errorf("rename.tsv:%d: empty column", line)
		}
		out = append(out, Rule{From: from, To: to})
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("rename.tsv: no rules")
	}
	return out, nil
}

func isSimpleClass(sel string) bool {
	if !strings.HasPrefix(sel, ".") {
		return false
	}
	name := sel[1:]
	if name == "" {
		return false
	}
	for _, r := range name {
		if r != '-' && r != '_' && !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func classNames(sel string) []string {
	sel = strings.TrimSpace(sel)
	if strings.Contains(sel, ",") {
		return nil
	}
	var out []string
	for _, p := range strings.Split(sel, ".") {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		if strings.ContainsAny(p, " >+~:[") {
			return nil
		}
		out = append(out, p)
	}
	return out
}
