package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/deislabs/ratify/pkg/executor/types"
	"github.com/spf13/cobra"
	"oras.land/oras-go/v2/registry/remote"
)

type pullOpts struct {
	image  string
	config string
}

func pullCommand(opts *pullOpts) *cobra.Command {
	if opts == nil {
		opts = &pullOpts{}
	}
	cmd := &cobra.Command{
		Use:   "pull IMAGE",
		Short: "Download an image from a registry with ratification",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.image = args[0]
			return runPull(cmd.Context(), opts)
		},
	}
	cmd.Flags().StringVarP(&opts.config, "config", "c", "", "Config File Path")
	cmd.MarkFlagRequired("config")
	return cmd
}

func runPull(ctx context.Context, opts *pullOpts) error {
	ref, err := resolveDigest(ctx, opts.image)
	if err != nil {
		return err
	}

	if err := runRatifyVerify(ctx, opts.config, ref); err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, "docker", "pull", ref)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func resolveDigest(ctx context.Context, image string) (string, error) {
	repo, err := remote.NewRepository(image)
	if err != nil {
		return "", err
	}
	desc, err := repo.Resolve(ctx, image)
	if err != nil {
		return "", err
	}
	ref := repo.Reference
	ref.Reference = desc.Digest.String()
	return ref.String(), nil
}

func runRatifyVerify(ctx context.Context, config, ref string) error {
	cmd := exec.CommandContext(ctx, "ratify", "verify", "-c", config, "-s", ref)
	cmd.Stderr = os.Stderr
	result, err := cmd.Output()
	if err != nil {
		return err
	}
	var report types.VerifyResult
	if err := json.Unmarshal(result, &report); err != nil {
		return err
	}
	if !report.IsSuccess {
		os.Stdout.Write(result)
		return fmt.Errorf("ratify verification failed")
	}
	if len(report.VerifierReports) == 0 {
		return fmt.Errorf("no ratifications found")
	}
	return nil
}
