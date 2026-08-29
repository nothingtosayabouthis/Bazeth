package tld

import "fmt"

func Print(result *Result) {
	fmt.Printf("[*] TLD           : %s\n", result.TLD)
	fmt.Printf("[*] Type          : %s\n", result.Type)

	if result.Manager != "" {
		fmt.Printf("[*] Manager       : %s\n", result.Manager)
	}

	if result.RDAP != "" {
		fmt.Printf("[*] RDAP          : %s\n", result.RDAP)
	}

	if result.WHOIS != "" {
		fmt.Printf("[*] WHOIS         : %s\n", result.WHOIS)
	}
}
