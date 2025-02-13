package main

import (
	"context"
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
	if updated == current {
		return nil
	}
	return os.WriteFile(c.DefaultArgs.Filename, []byte(updated), 0644)
}
