// Command record prints the Go expression of an element of a frame — the
// value of a reference record in internal/refs, drafted from the frame:
//
//	go run ./internal/framegen/cmd/record s-contacts-6 tc-danger 1
package main

import (
	"fmt"
	"os"
	"strconv"

	"carrel.dev/typecase/internal/framegen"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "usage: record <data-frame> <block class> <n>")
		os.Exit(2)
	}
	n, err := strconv.Atoi(os.Args[3])
	if err == nil {
		var expr string
		expr, err = framegen.Record(".", os.Args[1], os.Args[2], n)
		if err == nil {
			fmt.Println(expr)
			return
		}
	}
	fmt.Fprintln(os.Stderr, "record:", err)
	os.Exit(1)
}
