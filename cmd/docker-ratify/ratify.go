package main

import "github.com/spf13/cobra"

func ratifyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "ratify",
	}
	cmd.AddCommand(
		pullCommand(nil),
	)
	return cmd
}
