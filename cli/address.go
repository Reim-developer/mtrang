package cli

import (
	"mtrang/core"
	"mtrang/utils"

	"github.com/spf13/cobra"
)

func AddressCommand(cli *cobra.Command) {
	var websiteURL string

	addressCmd := &cobra.Command{
		Use: "addr",
		Run: func(cmd *cobra.Command, _ []string) {
			if websiteURL == "" {
				utils.Log("Usage: mtrang addr --url <WEBSITE_URL>")
				return
			}

			debug, _ := cmd.Flags().GetBool("debug")
			if debug {
				core.AddrLookup(websiteURL, false)
				return
			}

			core.AddrLookup(websiteURL, true)

		},
	}

	addressCmd.Flags().StringVarP(&websiteURL, "url", "u", "", "Website you want get lookup address")
	addressCmd.Flags().BoolP("debug", "", false, "Enable debug mode")
	cli.AddCommand(addressCmd)
}
