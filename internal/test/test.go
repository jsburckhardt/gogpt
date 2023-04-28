// Package test provides utilities for testing Cobra commands.
package test

import (
	"bytes"

	"github.com/spf13/cobra"
)

// ExecuteCommand used for running cobra commands
func ExecuteCommand(cmd *cobra.Command, args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)
	err = cmd.Execute()
	return buf.String(), err
}
