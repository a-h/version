package main

import (
	"context"
	_ "embed"
	"fmt"
)

type VersionCommand struct {
}

//go:embed .version
var version string

func (c VersionCommand) Run(ctx context.Context) (err error) {
	fmt.Println(version)
	return nil
}
