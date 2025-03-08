package cli

import (
	"mtrang/utils"

	"github.com/spf13/cobra"
)

const VERSION = "1.0.0"

func VersionCommand(cli *cobra.Command) {
	version_command := &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			utils.Log(VERSION)
		},
	}

	cli.AddCommand(version_command)
}
