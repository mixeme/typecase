// Command record drafts and keeps the reference records of internal/refs
// from their frames. Run it from the module root.
//
//	go run ./internal/framegen/cmd/record s-contacts-6 tc-danger 1
//
// prints the Go expression of that element of the frame — the value of a
// record, drafted (add.py puts a new one in).
//
//	go run ./internal/framegen/cmd/record -refresh [-force] [frame ...]
//
// redrafts every record that no occurrence of its frame prints as (of the
// frames named, or of all) from the nearest occurrence, in its file. A record
// has no number (fragment.Ref): it stands for the occurrence its widget
// prints as, so a frame's edit that only put a block in leaves it be, and
// one that changed the record's own block is redrafted from where it stood.
// A redraft that would change what the record stands for — the root's own
// classes, its variant, tone or state — is refused (the review of 14.10);
// -force redrafts it when the frame's element changed so on purpose.
//
//	go run ./internal/framegen/cmd/record -remove frame:class[:text] ...
//
// removes the records of that frame and block that no occurrence prints as
// any more (and, with text, only those whose source holds it): after a frame
// lost an element. Then TestShapesCovered says which shape needs a record
// again (add.py).
package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"

	"carrel.dev/typecase/internal/fragment"
	"carrel.dev/typecase/internal/framegen"
	"carrel.dev/typecase/internal/refs"
)

func main() {
	refresh := flag.Bool("refresh", false, "redraft the records no occurrence prints as")
	force := flag.Bool("force", false, "with -refresh: redraft even when the record's state would change")
	remove := flag.Bool("remove", false, "remove the records of frame:class[:text] no occurrence prints as")
	flag.Parse()
	args := flag.Args()
	var err error
	switch {
	case *refresh:
		err = redraft(args, *force)
	case *remove:
		err = drop(args)
	case len(args) == 3:
		var n int
		if n, err = strconv.Atoi(args[2]); err == nil {
			var expr string
			if expr, err = framegen.Record(".", args[0], args[1], n); err == nil {
				fmt.Println(expr)
			}
		}
	default:
		fmt.Fprintln(os.Stderr, "usage: record <data-frame> <block class> <n> | -refresh [-force] [frame ...] | -remove frame:class[:text] ...")
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "record:", err)
		os.Exit(1)
	}
}

// red is a record no occurrence prints as, and its nearest one (At.N is 0
// when the frame has no such block at all).
type red struct {
	e  refs.Entry
	at fragment.Ref
}

func reds(m *fragment.Mockup, keep func(refs.Entry) bool) ([]red, error) {
	var out []red
	for _, e := range refs.Entries {
		if !keep(e) {
			continue
		}
		_, err := refs.Place(m, e)
		if err == nil {
			continue
		}
		var nf *fragment.NotFound
		if errors.As(err, &nf) {
			out = append(out, red{e, nf.At})
			continue
		}
		if _, ok := m.Frame(e.Ref.Frame); ok && strings.Contains(err.Error(), "the frame has no .") {
			out = append(out, red{e: e})
			continue
		}
		return nil, err
	}
	return out, nil
}

func redraft(frames []string, force bool) error {
	m, err := fragment.Load(".")
	if err != nil {
		return err
	}
	rs, err := reds(m, func(e refs.Entry) bool { return len(frames) == 0 || slices.Contains(frames, e.Ref.Frame) })
	if err != nil {
		return err
	}
	ed := editor{}
	var refused []string
	for _, r := range rs {
		if r.at.N == 0 {
			refused = append(refused, fmt.Sprintf("%s (%s): the frame has no such block — remove it (-remove)", r.e.Ref, where(r.e)))
			continue
		}
		if !force {
			want, err := m.RootState(r.at)
			if err != nil {
				return err
			}
			got, err := fragment.RenderedState(r.at.Class, r.e.Fixture.Render())
			if err != nil {
				return err
			}
			slices.Sort(want)
			slices.Sort(got)
			if !slices.Equal(want, got) {
				refused = append(refused, fmt.Sprintf("%s (%s): the record %q is [%s], the nearest, #%d, is [%s] — redrafting would change what it stands for; add a record for it (add.py) or pass -force",
					r.e.Ref, where(r.e), r.e.State, strings.Join(got, " "), r.at.N, strings.Join(want, " ")))
				continue
			}
		}
		expr, err := framegen.Record(".", r.at.Frame, r.at.Class, r.at.N)
		if err != nil {
			refused = append(refused, fmt.Sprintf("%s (%s): the generator cannot draft #%d: %v", r.e.Ref, where(r.e), r.at.N, err))
			continue
		}
		if err := ed.fixture(r.e, expr); err != nil {
			return err
		}
		fmt.Printf("redrafted %s from #%d (%s)\n", r.e.Ref, r.at.N, where(r.e))
	}
	if err := ed.write(); err != nil {
		return err
	}
	for _, s := range refused {
		fmt.Println("refused", s)
	}
	if len(refused) > 0 {
		return fmt.Errorf("%d refused", len(refused))
	}
	return nil
}

func drop(args []string) error {
	m, err := fragment.Load(".")
	if err != nil {
		return err
	}
	ed := editor{}
	for _, a := range args {
		parts := strings.SplitN(a, ":", 3)
		if len(parts) < 2 {
			return fmt.Errorf("%s: want frame:class[:text]", a)
		}
		rs, err := reds(m, func(e refs.Entry) bool { return e.Ref.Frame == parts[0] && e.Ref.Class == parts[1] })
		if err != nil {
			return err
		}
		n := 0
		for _, r := range rs {
			ok, err := ed.remove(r.e, parts[2:])
			if err != nil {
				return err
			}
			if ok {
				n++
				fmt.Printf("removed %s (%s)\n", r.e.Ref, where(r.e))
			}
		}
		if n == 0 {
			fmt.Println("no record to remove:", a)
		}
	}
	return ed.write()
}

func where(e refs.Entry) string { return fmt.Sprintf("%s:%d", filepath.Base(e.Pos.File), e.Pos.Line) }

// editor gathers the changes to the record files and writes each file once,
// from its end, so the offsets stay true.
type editor struct {
	files map[string]*source
}

type source struct {
	fset  *token.FileSet
	src   []byte
	calls []*ast.CallExpr // the records: the arguments of add(…)
	edits []edit
}

type edit struct {
	from, to int
	text     string
}

func (ed *editor) load(path string) (*source, error) {
	if ed.files == nil {
		ed.files = map[string]*source{}
	}
	if s, ok := ed.files[path]; ok {
		return s, nil
	}
	src, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, src, 0)
	if err != nil {
		return nil, err
	}
	s := &source{fset: fset, src: src}
	ast.Inspect(f, func(n ast.Node) bool {
		c, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "add" {
			for _, a := range c.Args {
				if rec, ok := a.(*ast.CallExpr); ok {
					s.calls = append(s.calls, rec)
				}
			}
			return false
		}
		return true
	})
	ed.files[path] = s
	return s, nil
}

// record is the call that writes e: the argument of add(…) whose lines hold
// the line e was made on.
func (ed *editor) record(e refs.Entry) (*source, *ast.CallExpr, error) {
	if e.Pos.File == "" {
		return nil, nil, fmt.Errorf("%s: the record does not know where it is written", e.Ref)
	}
	s, err := ed.load(e.Pos.File)
	if err != nil {
		return nil, nil, err
	}
	for _, c := range s.calls {
		if s.fset.Position(c.Pos()).Line <= e.Pos.Line && e.Pos.Line <= s.fset.Position(c.End()).Line {
			if len(c.Args) < 3 {
				break
			}
			return s, c, nil
		}
	}
	return nil, nil, fmt.Errorf("%s (%s): no record call of add(…) on that line", e.Ref, where(e))
}

// fixture puts expr in place of the record's value: the argument before its
// frame and its class.
func (ed *editor) fixture(e refs.Entry, expr string) error {
	s, c, err := ed.record(e)
	if err != nil {
		return err
	}
	v := c.Args[len(c.Args)-3]
	s.edits = append(s.edits, edit{s.fset.Position(v.Pos()).Offset, s.fset.Position(v.End()).Offset, expr})
	return nil
}

// remove takes the record out with its line, when its source holds every
// text of must.
func (ed *editor) remove(e refs.Entry, must []string) (bool, error) {
	s, c, err := ed.record(e)
	if err != nil {
		return false, err
	}
	from, to := s.fset.Position(c.Pos()).Offset, s.fset.Position(c.End()).Offset
	for _, t := range must {
		if !bytes.Contains(s.src[from:to], []byte(t)) {
			return false, nil
		}
	}
	for from > 0 && s.src[from-1] != '\n' {
		from--
	}
	for to < len(s.src) && s.src[to] != '\n' {
		to++
	}
	s.edits = append(s.edits, edit{from, min(to+1, len(s.src)), ""})
	return true, nil
}

func (ed *editor) write() error {
	for path, s := range ed.files {
		if len(s.edits) == 0 {
			continue
		}
		sort.Slice(s.edits, func(i, j int) bool { return s.edits[i].from > s.edits[j].from })
		out := s.src
		for _, x := range s.edits {
			out = append(append(append([]byte{}, out[:x.from]...), x.text...), out[x.to:]...)
		}
		b, err := format.Source(out)
		if err != nil {
			return fmt.Errorf("%s: %v", path, err)
		}
		if err := os.WriteFile(path, b, 0o644); err != nil {
			return err
		}
	}
	return nil
}
