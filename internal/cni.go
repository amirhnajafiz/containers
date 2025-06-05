//go:build linux
// +build linux

package internal

import (
	"fmt"
	"os/exec"
	"strconv"
)

// setupContainerNetwork sets up a network namespace and veth pair for the container.
func setupContainerNetwork(containerID string, pid int, containerIP string) error {
	hostVeth := "veth" + containerID[:5]
	contVeth := "eth0"

	// create veth pair
	if err := exec.Command("ip", "link", "add", hostVeth, "type", "veth", "peer", "name", contVeth).Run(); err != nil {
		return fmt.Errorf("failed to create veth pair: %w", err)
	}

	// set host veth up
	if err := exec.Command("ip", "link", "set", hostVeth, "up").Run(); err != nil {
		return fmt.Errorf("failed to set host veth up: %w", err)
	}

	// move container veth to container's netns
	if err := exec.Command("ip", "link", "set", contVeth, "netns", strconv.Itoa(pid)).Run(); err != nil {
		return fmt.Errorf("failed to move veth to netns: %w", err)
	}

	// set up container veth and assign IP inside the container netns
	if err := exec.Command("ip", "netns", "exec", strconv.Itoa(pid), "ip", "link", "set", contVeth, "up").Run(); err != nil {
		return fmt.Errorf("failed to set container veth up: %w", err)
	}
	if err := exec.Command("ip", "netns", "exec", strconv.Itoa(pid), "ip", "addr", "add", containerIP+"/24", "dev", contVeth).Run(); err != nil {
		return fmt.Errorf("failed to assign IP: %w", err)
	}

	// set default route (optional)
	if err := exec.Command("ip", "netns", "exec", strconv.Itoa(pid), "ip", "route", "add", "default", "dev", contVeth).Run(); err != nil {
		return fmt.Errorf("failed to set default route: %w", err)
	}

	return nil
}

// removeContainerNetwork removes the network setup for the container, including the veth pair.
func removeContainerNetwork(containerID string) error {
	hostVeth := "veth" + containerID[:5]

	// delete veth pair
	if err := exec.Command("ip", "link", "delete", hostVeth).Run(); err != nil {
		return fmt.Errorf("failed to delete veth pair: %w", err)
	}

	return nil
}
