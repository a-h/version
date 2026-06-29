package main

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type PushCommand struct {
	DefaultArgs `embed:""`
	Prefix      *string `help:"The prefix to use for the version, e.g. v. If unset, the prefix is inferred from existing tags, defaulting to v when no tags exist."`
	Force       bool    `help:"Push the tag even if its prefix differs from existing tags." default:"false"`
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

	existing, hasTags, err := getLatestTagPrefix()
	if err != nil {
		return err
	}
	prefix, err := resolvePrefix(c.Prefix, c.Force, existing, hasTags)
	if err != nil {
		return err
	}
	if prefix != "" && !strings.HasPrefix(updated, prefix) {
		updated = prefix + updated
	}
	return pushTag(updated)
}

// resolvePrefix determines the tag prefix to use, preferring an explicit
// --prefix flag, then the prefix of existing tags, and finally defaulting to v
// when no tags exist. It returns an error if an explicit flag differs from the
// prefix used by existing tags, unless force is set.
func resolvePrefix(flag *string, force bool, existing string, hasTags bool) (prefix string, err error) {
	if flag != nil {
		prefix = *flag
		if hasTags && prefix != existing && !force {
			return "", fmt.Errorf("prefix %q does not match the prefix %q used by existing tags, use --force to push anyway", prefix, existing)
		}
		return prefix, nil
	}

	if hasTags {
		return existing, nil
	}
	return "v", nil
}

// getLatestTagPrefix returns the non-numeric leading characters of the
// highest-versioned existing tag. hasTags is false when the repository has no
// tags.
func getLatestTagPrefix() (prefix string, hasTags bool, err error) {
	gitPath, err := exec.LookPath("git")
	if err != nil {
		return "", false, fmt.Errorf("failed to find git on path: %w", err)
	}
	cmd := exec.Command(gitPath, "tag", "--sort=-v:refname")
	output, err := cmd.Output()
	if err != nil {
		return "", false, fmt.Errorf("failed to run git: %w", err)
	}
	tags := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(tags) == 0 || tags[0] == "" {
		return "", false, nil
	}
	return leadingPrefix(tags[0]), true, nil
}

// leadingPrefix returns the leading characters of tag that precede the first
// digit, e.g. "v" for "v0.0.2" and "" for "0.0.2". It returns the whole tag if
// it contains no digit.
func leadingPrefix(tag string) (prefix string) {
	i := strings.IndexAny(tag, "0123456789")
	if i < 0 {
		return tag
	}
	return tag[:i]
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
