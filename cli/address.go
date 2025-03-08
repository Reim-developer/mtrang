package cli

import (
	"mtrang/core"
	"mtrang/utils"

	"github.com/spf13/cobra"
)

func AddressCommand(cli *cobra.Command) {
	var website_url string
	address_cmd := &cobra.Command{
		Use: "addr",
		Run: func(cmd *cobra.Command, args []string) {
			if website_url == "" {
				utils.Log("Usage: mtrang addr --url <WEBSITE_URL>")
				return
			}

			debug, _ := cmd.Flags().GetBool("debug")
			if debug {
				core.AddrLookup(website_url, false)
			} else {
				core.AddrLookup(website_url, true)
			}
		},
	}

	address_cmd.Flags().StringVarP(&website_url, "url", "u", "", "Website you want get lookup address")
	address_cmd.Flags().BoolP("debug", "", false, "Enable debug mode")
	cli.AddCommand(address_cmd)
}
