package main

import (
	"mtrang/cli"
	"mtrang/utils"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "mtrang",
		Short: "Pentest Command Line Interface",

		Run: func(cmd *cobra.Command, args []string) {
			utils.Log("Try: mtrang help for more infomation")
		},
	}

	cli.VersionCommand(cmd)
	cli.AddressCommand(cmd)

	if err := cmd.Execute(); err != nil {
		utils.Fatal("Fatal: %s", err)
	}

}
