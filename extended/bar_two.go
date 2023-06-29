package main

import (
	"cli/cmd"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cilium/cilium/pkg/hive/cell"
)

const barTwoParam = "bar-two-param"

var BarTwo = cell.Module(
	"bar-two",
	"bar-two",

	cell.Provide(func(p barTwoParams) cmd.BarCommandOut {
		var param *bool

		barTwoCmd := &cobra.Command{
			Use:   "two",
			Short: "two",
			Run: func(cmd *cobra.Command, args []string) {
				p.Logger.Infof("bar two requested with param %v", *param)
			},
		}

		param = barTwoCmd.Flags().Bool(barTwoParam, false, "bar two parameter")

		return cmd.BarCommandOut{Cmd: barTwoCmd}
	}),
)

type barTwoParams struct {
	cell.In

	Logger logrus.FieldLogger
}
