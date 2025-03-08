package core

import (
	"mtrang/utils"
	"net"
	"regexp"
)

// Lookup website address with given URL
func AddrLookup(url string, debug bool) {
	reg_match := regexp.MustCompile(`^(https?://)`)
	url = reg_match.ReplaceAllString(url, "")

	ips, err := net.LookupHost(url)

	if err != nil {
		if debug {
			utils.Fatal("Failed to lookup host of: %s. Use --debug to show traceback", url)
			return
		} else {
			utils.Fatal("Failed to lookup host of: %s\nTraceback: %s", url, err)
			return
		}
	}

	utils.Log("Address of %s:\n", url)
	for _, address_str := range ips {
		address := net.ParseIP(address_str)

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
