package main

import (
	"os"

	release_log "github.com/openshift/cluster-debug-tools/pkg/cmd/release-log"

	analyze_e2e "github.com/openshift/cluster-debug-tools/pkg/cmd/analyze-e2e"
	"github.com/openshift/cluster-debug-tools/pkg/cmd/audit"
	"github.com/openshift/cluster-debug-tools/pkg/cmd/certinspection"
	"github.com/openshift/cluster-debug-tools/pkg/cmd/events"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	mustgather "github.com/openshift/cluster-debug-tools/pkg/cmd"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type DevHelpersOptions struct {
	configFlags *genericclioptions.ConfigFlags

	genericclioptions.IOStreams
}

func NewDevHelpersOptions(streams genericclioptions.IOStreams) *DevHelpersOptions {
	return &DevHelpersOptions{
		configFlags: genericclioptions.NewConfigFlags(true),
		IOStreams:   streams,
	}
}

func NewCmdDevHelpers(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewDevHelpersOptions(streams)

	cmd := &cobra.Command{
		Use:          "openshift-dev-helpers",
		Short:        "Set of helpers for OpenShift developer teams",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(c, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			if err := o.Run(); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.AddCommand(events.NewCmdEvent("openshift-dev-helpers", streams))
	cmd.AddCommand(audit.NewCmdAudit("openshift-dev-helpers", streams))
	cmd.AddCommand(mustgather.NewCmdRevisionStatus("openshift-dev-helpers", streams))
	cmd.AddCommand(certinspection.NewCmdCertInspection(streams))
	cmd.AddCommand(analyze_e2e.NewCmdAnalyze("openshift-dev-helpers", streams))
	cmd.AddCommand(release_log.NewCmdReleaseLog("openshift-dev-helpers", streams))

	return cmd
}

func (o *DevHelpersOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

func (o *DevHelpersOptions) Validate() error {
	return nil
}

func (o *DevHelpersOptions) Run() error {
	return nil
}

func main() {
	flags := pflag.NewFlagSet("dev-helpers", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := NewCmdDevHelpers(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
