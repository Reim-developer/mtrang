package core

import (
	"net"
	"regexp"

	"mtrang/utils"
)

// AddrLookup website address with given URL.
func AddrLookup(url string, debug bool) {
	regMatch := regexp.MustCompile(`^(https?://)`)
	url = regMatch.ReplaceAllString(url, "")

	ips, err := net.LookupHost(url)
	if err != nil {
		if debug {
			utils.Fatal("Failed to lookup host of: %s. Use --debug to show traceback", url)
		} else {
			utils.Fatal("Failed to lookup host of: %s\nTraceback: %s", url, err)

			return
		}
	}

	utils.Log("Address of %s:\n", url)

	for _, addressStr := range ips {
		address := net.ParseIP(addressStr)

		if address == nil {
			continue
		}

		if address.To4() != nil {
			utils.Log("Address V4: %s", address)
		} else {
			utils.Log("Address V6: %s", address)
		}
	}
}
