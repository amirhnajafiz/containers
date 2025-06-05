//go:build linux
// +build linux

package main

import (
	"os"

	"github.com/amirhnajafiz/containers/internal"
)

func main() {
	switch os.Args[1] {
	case "run":
		must(internal.Parent())
	case "child":
		must(internal.Child())
	default:
		panic("invalid command, expected 'run' or 'child'")
	}
}

// must is a helper function that panics if the error is not nil.
func must(err error) {
	if err != nil {
		panic(err)
	}
}
