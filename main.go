package main

import (
	"github.com/spf13/cobra"

	"mtrang/cli"
	"mtrang/utils"
)

func main() {
	cmd := &cobra.Command{
		Use:   "mtrang",
		Short: "Pentest Command Line Interface",

		Run: func(_ *cobra.Command, _ []string) {
			utils.Log("Try: mtrang help for more information")
		},
	}

	cli.VersionCommand(cmd)
	cli.AddressCommand(cmd)
	cli.ScanPortCommand(cmd)

	if err := cmd.Execute(); err != nil {
		utils.Fatal("Fatal: %s", err)
	}
}
