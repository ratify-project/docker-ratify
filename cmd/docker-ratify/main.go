package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	cmd := &cobra.Command{
		Use:          "docker-ratify",
		Short:        "A docker plugin wrapper for ratify",
		SilenceUsage: true,
	}
	cmd.AddCommand(
		metadataCommand(),
		ratifyCommand(),
	)
	if err := cmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
