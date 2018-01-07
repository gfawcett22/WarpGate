package main

import (
	"github.com/spf13/cobra"
)

func RunVersion(cmd *cobra.Command, args []string) {
	cmd.Printf("WarpGate v%s", version)
}
