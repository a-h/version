package main

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type PushCommand struct {
	DefaultArgs `embed:""`
}

func (c PushCommand) Run(ctx context.Context) (err error) {
	// Check that we're on the main branch.
	branch, err := getBranch()
	if err != nil {
		return err
	}
	if branch != "main" {
		return fmt.Errorf("not on the main branch, currently on %q", branch)
	}

	// Check the version is up to date.
	current, updated, err := c.DefaultArgs.GetVersion()
	if err != nil {
		return err
	}
	if updated != current {
		return fmt.Errorf("version file %q contains %q, but current version is %q", c.DefaultArgs.Filename, current, updated)
	}

	if updated == "" {
		return fmt.Errorf("error creating version, version is empty")
	}

	// Push the tag.
	if !strings.HasPrefix(updated, "v") {
		updated = "v" + updated
	}
	return pushTag(updated)
}

func getBranch() (branch string, err error) {
	gitPath, err := exec.LookPath("git")
	if err != nil {
		return "", fmt.Errorf("failed to find git on path: %w", err)
	}
	cmd := exec.Command(gitPath, "branch", "--show-current")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to run git: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}

func pushTag(tag string) (err error) {
	gitPath, err := exec.LookPath("git")
	if err != nil {
		return fmt.Errorf("failed to find git on path: %w", err)
	}
	cmd := exec.Command(gitPath, "tag", tag)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run git: %w", err)
	}
	cmd = exec.Command(gitPath, "push", "origin", tag)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run git: %w", err)
	}
	return nil
}
