package main

import (
	"cli/cmd"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cilium/cilium/pkg/hive"
	"github.com/cilium/cilium/pkg/hive/cell"
)

var Hive *hive.Hive

var CLIExt = cell.Module(
	"cli-ext",
	"cli-ext",

	cell.Invoke(registerRootCmd),

	cmd.Commands,

	// extended-only commands
	BarTwo,
)

type params struct {
	cell.In

	Logger     logrus.FieldLogger
	Lifecycle  hive.Lifecycle
	Shutdowner hive.Shutdowner

	SubCommands []*cobra.Command `group:"cli-command"`
}

func registerRootCmd(p params) {
	rootCmd := &cobra.Command{
		Use:   "cli-ext",
		Short: "cli extended",
		Run: func(cmd *cobra.Command, args []string) {
			p.Logger.Info("cli-ext requested")
		},
	}

	for _, cmd := range p.SubCommands {
		rootCmd.AddCommand(cmd)
	}

	rootCmd.AddCommand(Hive.Command())

	Hive.RegisterFlags(rootCmd.Flags())

	p.Lifecycle.Append(hive.Hook{
		OnStart: func(hive.HookContext) error {
			if err := rootCmd.Execute(); err != nil {
				return err
			}
			p.Shutdowner.Shutdown()
			return nil
		},
	})
}
