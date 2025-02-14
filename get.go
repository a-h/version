package main

import (
	"context"
	"fmt"
)

type GetCommand struct {
	DefaultArgs `embed:""`
}

func (c GetCommand) Run(ctx context.Context) (err error) {
	current, updated, err := c.DefaultArgs.GetVersion()
	if err != nil {
		return err
	}
	if updated != current {
		return fmt.Errorf("version file %q contains %q, but current version is %q", c.DefaultArgs.Filename, current, updated)
	}
	fmt.Print(updated)
	return nil
}
