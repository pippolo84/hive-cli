package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cilium/cilium/pkg/hive"
	"github.com/cilium/cilium/pkg/hive/cell"
)

type CLIParams struct {
	cell.In

	Logger     logrus.FieldLogger
	Lifecycle  hive.Lifecycle
	Shutdowner hive.Shutdowner

	SubCommands []*cobra.Command `group:"cli-command"`
}

type CLICommandOut struct {
	cell.Out

	Cmd *cobra.Command `group:"cli-command"`
}

var Hive *hive.Hive

var CLI = cell.Module(
	"cli",
	"cli",

	cell.Invoke(registerRootCmd),

	Commands,
)

func registerRootCmd(p CLIParams) {
	rootCmd := &cobra.Command{
		Use:   "cli",
		Short: "cli",
		Run: func(cmd *cobra.Command, args []string) {
			p.Logger.Info("cli requested")
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
