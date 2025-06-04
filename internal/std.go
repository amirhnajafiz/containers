package internal

import "os/exec"

// setContainerSTD sets the standard input, output, and error streams of the command to nil.
func setContainerSTD(cmd *exec.Cmd) *exec.Cmd {
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil

	return cmd
}
