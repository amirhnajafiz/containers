//go:build (linux && 386) || (darwin && !cgo)
// +build linux,386 darwin,!cgo

package cgroup

import (
	"os"
	"path/filepath"
	"syscall"

	"github.com/amirhnajafiz-learning/containers/pkg/enums"
)

func PivotRoot(newRoot string) error {
	putold := filepath.Join(newRoot, "/.pivot_root")

	// bind mount newroot to putold to make putold a valid mount point
	if err := syscall.Mount(newRoot, newRoot, "", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return err
	}

	// create putold directory
	if err := os.MkdirAll(putold, enums.PermGroupAll); err != nil {
		return err
	}

	// call pivot root
	if err := syscall.PivotRoot(newRoot, putold); err != nil {
		return err
	}

	// change the current working directory to the new root
	if err := os.Chdir("/"); err != nil {
		return err
	}

	// unmount putold, which now lives at /.pivot_root
	if err := syscall.Unmount("./pivot_root", syscall.MNT_DETACH); err != nil {
		return err
	}

	return nil
}
