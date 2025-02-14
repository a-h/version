package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/alecthomas/kong"
)

type DefaultArgs struct {
	Template string `help:"The template to use for the version." default:"0.0.%d"`
	Filename string `help:"The name of the file to write the version to." default:".version"`
	FirstRun bool   `help:"Use to create the version file." default:"false"`
}

func isDirty() (dirty bool, err error) {
	gitPath, err := exec.LookPath("git")
	if err != nil {
		return false, fmt.Errorf("failed to find git on path: %w", err)
	}
	cmd := exec.Command(gitPath, "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("failed to run git: %w", err)
	}
	return len(strings.TrimSpace(string(output))) > 0, nil
}

func getCommitCount() (count int, err error) {
	gitPath, err := exec.LookPath("git")
	if err != nil {
		return 0, fmt.Errorf("failed to find git on path: %w", err)
	}
	cmd := exec.Command(gitPath, "rev-list", "main", "--count")
	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("failed to run git: %w", err)
	}
	count, err = strconv.Atoi(strings.TrimSpace(string(output)))
	if err != nil {
		return 0, fmt.Errorf("failed to parse git output: %w", err)
	}
	return count, nil
}

func (da DefaultArgs) GetVersion() (current, updated string, err error) {
	currentFileBytes, err := os.ReadFile(da.Filename)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return "", "", fmt.Errorf("failed to read version file: %w", err)
	}
	if errors.Is(err, os.ErrNotExist) && !da.FirstRun {
		return "", "", fmt.Errorf("version file %q does not exist, use the --first-run flag to create it", da.Filename)
	}
	current = strings.TrimSpace(string(currentFileBytes))

	count, err := getCommitCount()
	if err != nil {
		return "", "", fmt.Errorf("failed to get commit count: %w", err)
	}
	dirty, err := isDirty()
	if err != nil {
		return "", "", fmt.Errorf("failed to check dirty status: %w", err)
	}
	if dirty {
		// If there are changes to be committed, then the version number will be incremented.
		count++
	}
	updated = fmt.Sprintf(da.Template, count)

	return current, updated, nil
}

type CLI struct {
	Check   CheckCommand   `cmd:"check" help:"Check that the version has been properly updated. Returns a human readable message."`
	Get     GetCommand     `cmd:"get" help:"Get the version number, fails if the version file is not up-to-date."`
	Set     SetCommand     `cmd:"set" help:"Update the version number file if needed."`
	Push    PushCommand    `cmd:"push" help:"Push an updated tag to git."`
	Version VersionCommand `cmd:"version" help:"Print the version number."`
}

func main() {
	var cli CLI
	ctx := context.Background()
	kctx := kong.Parse(&cli, kong.UsageOnError(), kong.BindTo(ctx, (*context.Context)(nil)))
	if err := kctx.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
