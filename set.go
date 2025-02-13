package main

import (
	"context"
	"fmt"
	"os"
)

type SetCommand struct {
	DefaultArgs `embed:""`
}

func (c SetCommand) Run(ctx context.Context) (err error) {
	current, updated, err := c.DefaultArgs.GetVersion()
	if err != nil {
		return err
	}
	if updated == current && current != "" {
		fmt.Printf("No change, current version is %q\n", current)
		return nil
	}
	if err := os.WriteFile(c.DefaultArgs.Filename, []byte(updated), 0644); err != nil {
		return err
	}
	fmt.Printf("Updated %q from %q to %q\n", c.DefaultArgs.Filename, current, updated)
	return nil
}
