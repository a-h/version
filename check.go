package main

import (
	"context"
	"fmt"
)

type CheckCommand struct {
	DefaultArgs `embed:""`
}

func (c CheckCommand) Run(ctx context.Context) (err error) {
	current, updated, err := c.DefaultArgs.GetVersion()
	if err != nil {
		return err
	}
	if updated != current {
		return fmt.Errorf("version file %q contains %q, but current version is %q", c.DefaultArgs.Filename, current, updated)
	}
	return nil
}
