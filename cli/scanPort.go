package cli

import (
	"mtrang/core"
	"mtrang/utils"

	"github.com/spf13/cobra"
)

func ScanPortCommand(cli *cobra.Command) {
	var port string
	var channels int
	var timeout int
	var debug bool
	scanPortCmd := &cobra.Command{
		Use:   "scan",
		Short: "sc",
		Run: func(_ *cobra.Command, _ []string) {
			if port == "" {
				utils.Log("Usage: mtrang scan --target <IP or domain>")
				return
			}

			core.ScanPort(port, channels, timeout, debug)
		},
	}

	scanPortCmd.Flags().StringVarP(&port, "target", "t", "", "Your target, such as domain or IP address")
	scanPortCmd.Flags().IntVarP(&channels, "workers", "", 50, "Concurrentcy workers. Default as 50")
	scanPortCmd.Flags().IntVarP(&timeout, "timeout", "", 1000, "Timeout for Dial Timeout. Default as 1000 Milliseconds")
	scanPortCmd.Flags().BoolVarP(&debug, "debug", "", false, "Debug mode. Default as FALSE")

	cli.AddCommand(scanPortCmd)
}
