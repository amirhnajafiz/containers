package cgroup

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/amirhnajafiz-learning/containers/pkg/conditions"
	"github.com/amirhnajafiz-learning/containers/pkg/enums"
)

// cgroups := /root/mygrp
// memory := 2M
func EnableCgroup(cgroups string, memory string) {
	pids := filepath.Join(cgroups, "child")

	conditions.Must(os.WriteFile(filepath.Join(pids, "memory.max"), []byte(memory), enums.PermGroupAll))
	conditions.Must(os.WriteFile(filepath.Join(pids, "cgroup.proc"), []byte(strconv.Itoa(os.Getpid())), enums.PermGroupAll))
}
