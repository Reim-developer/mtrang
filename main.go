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

	cmd.DisableFlagParsing = true
	cli.VersionCommand(cmd)
	cli.AddressCommand(cmd)
	cli.ScanPortCommand(cmd)
	cli.HelpCommand(cmd)

	if err := cmd.Execute(); err != nil {
		utils.Fatal("Fatal: %s", err)
	}
}
