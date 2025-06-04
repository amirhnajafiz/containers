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
		panic("wat should I do")
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
