package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cilium/cilium/pkg/hive/cell"
)

const barOneParam = "bar-one-param"

var BarOne = cell.Module(
	"bar-one",
	"bar-one",

	cell.Provide(func(p barOneParams) BarCommandOut {
		var param *bool

		barOneCmd := &cobra.Command{
			Use:   "one",
			Short: "one",
			Run: func(cmd *cobra.Command, args []string) {
				p.Logger.Infof("bar one requested with param %v", *param)
			},
		}

		param = barOneCmd.Flags().Bool(barOneParam, false, "bar one parameter")

		return BarCommandOut{Cmd: barOneCmd}
	}),
)

type barOneParams struct {
	cell.In

	Logger logrus.FieldLogger
}
