package cmd

import "github.com/cilium/cilium/pkg/hive/cell"

var Commands = cell.Module(
	"cli-commands",
	"cli-commands",

	Foo,

	Bar,
	BarOne,
)
