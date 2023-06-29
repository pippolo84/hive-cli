package main

import (
	"cli/cmd"
	"fmt"
	"os"

	"github.com/cilium/cilium/pkg/hive"
)

func main() {
	cmd.Hive = hive.New(cmd.CLI)

	if err := cmd.Hive.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
