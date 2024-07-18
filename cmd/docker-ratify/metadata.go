package main

import (
	"encoding/json"
	"os"

	"github.com/ratify-project/docker-ratify/internal/docker"
	"github.com/spf13/cobra"
)

var pluginMetadata = docker.PluginMetadata{
	SchemaVersion:    "0.1.0",
	Vendor:           "Ratify Project",
	Version:          "0.1.0",
	ShortDescription: "Artifact Ratification Framework",
	URL:              "https://github.com/ratify-project/docker-ratify",
	Experimental:     true,
}

func metadataCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   docker.PluginMetadataCommandName,
		Short: "Docker CLI plugin metadata",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			writer := json.NewEncoder(os.Stdout)
			writer.Encode(pluginMetadata)
		},
		Hidden: true,
	}
	return cmd
}
