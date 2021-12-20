package main

import (
	"context"
	"lit"
)

type DialCommand struct{}

func (c *DialCommand) Run(ctx context.Context, args []string) error {
	err := lit.ErrorF(lit.ENOTIMPLEMENTED, "Dial command is not implemented, yet")
	return err
}
