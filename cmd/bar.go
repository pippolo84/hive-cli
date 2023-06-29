package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cilium/cilium/pkg/hive/cell"
)

const barParam = "bar-param"

var Bar = cell.Module(
	"bar",
	"bar",

	cell.Provide(func(p barParams) CLICommandOut {
		var param *bool

		barCmd := &cobra.Command{
			Use:   "bar",
			Short: "bar",
			Run: func(cmd *cobra.Command, args []string) {
				p.Logger.Infof("bar requested with param %v", *param)
			},
		}

		param = barCmd.Flags().Bool(barParam, false, "bar parameter")

		for _, cmd := range p.SubCommands {
			barCmd.AddCommand(cmd)
		}

		return CLICommandOut{Cmd: barCmd}
	}),
)

type barParams struct {
	cell.In

	Logger logrus.FieldLogger

	SubCommands []*cobra.Command `group:"bar-command"`
}

type BarCommandOut struct {
	cell.Out

	Cmd *cobra.Command `group:"bar-command"`
}
