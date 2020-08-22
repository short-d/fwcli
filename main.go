package main

import (
	"fmt"
	"github.com/short-d/app/fw/cli"
	"github.com/short-d/fwcli/cmd"
	"os"
)

func main() {
	factory := cli.NewCobraFactory()
	rootCmd, err := cmd.NewRoot(factory)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
