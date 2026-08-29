package whois

import "fmt"

func Print(result *Result) {
	fmt.Printf("[*] Domain        : %s\n", result.Domain)

	if result.Registrar != "" {
		fmt.Printf("[*] Registrar     : %s\n", result.Registrar)
	}

	if result.Created != "" {
		fmt.Printf("[*] Created       : %s\n", result.Created)
	}

	if result.Updated != "" {
		fmt.Printf("[*] Updated       : %s\n", result.Updated)
	}

	if result.Expires != "" {
		fmt.Printf("[*] Expires       : %s\n", result.Expires)
	}

	if result.DNSSEC != "" {
		fmt.Printf("[*] DNSSEC        : %s\n", result.DNSSEC)
	}

	if len(result.Status) > 0 {
		fmt.Println()
		fmt.Println("[Status]")

		for _, status := range result.Status {
			fmt.Printf("[*] %s\n", status)
		}
	}

	if len(result.NameServer) > 0 {
		fmt.Println()
		fmt.Println("[Nameservers]")

		for _, ns := range result.NameServer {
			fmt.Printf("[*] %s\n", ns)
		}
	}
}
