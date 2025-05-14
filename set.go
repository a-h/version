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
	if updated != current {
		if err := os.WriteFile(c.DefaultArgs.Filename, []byte(updated), 0644); err != nil {
			return err
		}
	}
	fmt.Print(updated)
	return nil
}
