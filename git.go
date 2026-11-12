package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// runGit runs git with args and returns its standard output. The returned
// error includes the command that ran and, if git exited with an error, its
// standard error output.
func runGit(args ...string) (output []byte, err error) {
	gitPath, err := exec.LookPath("git")
	if err != nil {
		return nil, fmt.Errorf("failed to find git on path: %w", err)
	}
	cmd := exec.Command(gitPath, args...)
	output, err = cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) && len(exitErr.Stderr) > 0 {
			return nil, fmt.Errorf("failed to run %q: %w: %s", cmd.String(), err, strings.TrimSpace(string(exitErr.Stderr)))
		}
		return nil, fmt.Errorf("failed to run %q: %w", cmd.String(), err)
	}
	return output, nil
}
