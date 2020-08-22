package cmd

import (
	"fmt"
	"github.com/short-d/app/fw/cli"
	"os"
)

func NewRoot(factory cli.CommandFactory) (cli.Command, error) {
	config := cli.CommandConfig{
		Usage: "fwcli",
		OnExecute: func(cmd cli.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
	rootCmd := factory.NewCommand(config)

	newCmd := newNew(factory)
	err := rootCmd.AddSubCommand(newCmd)
	return rootCmd, err
}
