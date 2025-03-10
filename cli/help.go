package cli

import (
	"mtrang/utils"
	"strings"

	"github.com/spf13/cobra"
)

// defaultHelpMsg : Raw string for help message
const defaultHelpMsg = `Usage: mtrang help <Command Name> for more infomation
➤ Available Commands:
	[+] addr : Lookup website IP address.
	[+] scan : Scan port opening in website.
	[+] version : Display CLI version.
	[+] help : Display this help message.
`

const addrHelp = `Usage: mtrang addr url <WEBSITE_URL> debug
➤ Description:
	[+] url: Website you want lookup address.
	[+] debug:  Debug mode, optional.
`
const scanHelp = `Usage: mtrang scan target <WEBSITE_TARGET> workers <

`

// HelpCommand : Display help
func HelpCommand(cli *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "help",
		Short: "Help all about of Command Line",
		Run: func(_ *cobra.Command, args []string) {
			switch strings.ToLower(args[0]) {
			case "addr":
				utils.Log("%s", addrHelp)
				break
			default:
				utils.Log("%s", defaultHelpMsg)
			}
		},
	}

	cli.AddCommand(cmd)
}
