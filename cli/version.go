package cli

import (
	"mtrang/utils"

	"github.com/spf13/cobra"
)

const VERSION = "1.0.0"

func VersionCommand(cli *cobra.Command) {
	varsionCommand := &cobra.Command{
		Use: "version",
		Run: func(_ *cobra.Command, _ []string) {
			utils.Log(VERSION)
		},
	}

	cli.AddCommand(varsionCommand)
}
