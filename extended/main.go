package main

import (
	"fmt"
	"os"

	"github.com/cilium/cilium/pkg/hive"
)

func main() {
	Hive = hive.New(CLIExt)

	if err := Hive.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
