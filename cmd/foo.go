package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cilium/cilium/pkg/hive/cell"
)

const fooParam = "foo-param"

var Foo = cell.Module(
	"foo",
	"foo",

	cell.Provide(func(p fooParams) CLICommandOut {
		var param *bool

		fooCmd := &cobra.Command{
			Use:   "foo",
			Short: "foo",
			Run: func(cmd *cobra.Command, args []string) {
				p.Logger.Infof("foo requested with param %v", *param)
			},
		}

		param = fooCmd.Flags().Bool(fooParam, false, "foo parameter")

		for _, cmd := range p.SubCommands {
			fooCmd.AddCommand(cmd)
		}

		return CLICommandOut{Cmd: fooCmd}
	}),
)

type fooParams struct {
	cell.In

	Logger logrus.FieldLogger

	SubCommands []*cobra.Command `group:"foo-command"`
}
