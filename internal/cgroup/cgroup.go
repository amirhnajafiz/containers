package cgroup

import "path/filepath"

func EnableCgroup(path string) {
	cgroups := path // /root/mygrp
	pids := filepath.Join(cgroups, "child")
}
