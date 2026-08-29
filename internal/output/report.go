package output

import (
	"fmt"

	"bazeth/internal/ip"
)

// Print renders the Bazeth IP report.
func Print(result *ip.Result) {
	fmt.Printf("[*] Target        : %s\n", result.IP)
	fmt.Printf("[*] Version       : %s\n", result.Version)

	if result.ASN != "" {
		fmt.Printf("[*] ASN           : %s\n", result.ASN)
	}

	if result.Organization != "" {
		fmt.Printf("[*] Organization  : %s\n", result.Organization)
	}

	if result.Registry != "" {
		fmt.Printf("[*] Registry      : %s\n", result.Registry)
	}

	if result.NetworkName != "" {
		fmt.Printf("[*] Network       : %s\n", result.NetworkName)
	}

	if result.Country != "" {
		fmt.Printf("[*] Country       : %s\n", result.Country)
	}

	if result.CIDR != "" {
		fmt.Printf("[*] CIDR          : %s\n", result.CIDR)
	}

	if result.ReverseDNS != "" {
		fmt.Printf("[*] PTR           : %s\n", result.ReverseDNS)
	}

	if result.HTTPStatus != 0 || result.HTTPServer != "" || result.TLSCommonName != "" {
		fmt.Println()
		fmt.Println("[HTTP]")

		if result.HTTPStatus != 0 {
			fmt.Printf("[*] Status        : %d\n", result.HTTPStatus)
		}

		if result.HTTPServer != "" {
			fmt.Printf("[*] Server        : %s\n", result.HTTPServer)
		}

		if result.PoweredBy != "" {
			fmt.Printf("[*] Powered By    : %s\n", result.PoweredBy)
		}

		if result.Redirect != "" {
			fmt.Printf("[*] Redirect      : %s\n", result.Redirect)
		}

		if result.TLSCommonName != "" {
			fmt.Printf("[*] TLS CN        : %s\n", result.TLSCommonName)
		}
	}

	if result.AbuseEmail != "" {
		fmt.Println()
		fmt.Println("[Contacts]")
		fmt.Printf("[*] Abuse         : %s\n", result.AbuseEmail)
	}

	if len(result.Source) > 0 {
		fmt.Println()
		fmt.Println("[Sources]")

		for _, source := range result.Source {
			fmt.Printf("[*] %s\n", source)
		}
	}
}
