package main

import (
	"fmt"
	"os"

	"bazeth/internal/active/fingerprint"
	"bazeth/internal/dns"
	"bazeth/internal/output"
	"bazeth/internal/session"
	"bazeth/internal/tld"
	"bazeth/internal/whois"
)

const Version = "0.1.0-alpha"

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printUsage()
		return
	}

	switch args[0] {

	case "ip":
		if len(args) != 2 {
			fmt.Println("Usage: bazeth ip <address>")
			os.Exit(1)
		}

		result, err := session.Lookup(args[1])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		output.Print(result)

	case "dns":
		if len(args) != 2 {
			fmt.Println("Usage: bazeth dns <domain>")
			os.Exit(1)
		}

		client := dns.New()

		fmt.Printf("[*] Target        : %s\n", args[1])

		for _, record := range dns.Supported {
			dns.PrintSection(
				record.Name,
				must(client.Resolve(args[1], record.Type)),
			)
		}

	case "fingerprint":
		if len(args) != 2 {
			fmt.Println("Usage: bazeth fingerprint <address>")
			os.Exit(1)
		}

		info, err := fingerprint.HTTP(args[1])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Printf("[*] Target        : %s\n\n", args[1])
		fmt.Println("[HTTP]")

		fmt.Printf("[*] Status        : %d\n", info.Status)

		if info.Server != "" {
			fmt.Printf("[*] Server        : %s\n", info.Server)
		}

		if info.PoweredBy != "" {
			fmt.Printf("[*] Powered By    : %s\n", info.PoweredBy)
		}

		if info.Location != "" {
			fmt.Printf("[*] Redirect      : %s\n", info.Location)
		}

		if info.TLSCommonName != "" {
			fmt.Printf("[*] TLS CN        : %s\n", info.TLSCommonName)
		}

	case "whois":
		if len(args) != 2 {
			fmt.Println("Usage: bazeth whois <domain>")
			os.Exit(1)
		}

		client := whois.New()

		result, err := client.Lookup(args[1])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		whois.Print(result)

	case "tld":
		if len(args) != 2 {
			fmt.Println("Usage: bazeth tld <extension>")
			os.Exit(1)
		}

		result, err := tld.Lookup(args[1])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		tld.Print(result)

	case "version":
		fmt.Printf("Bazeth %s\n", Version)

	default:
		printUsage()
		os.Exit(1)
	}
}

func must(records []string, err error) []string {
	if err != nil {
		return nil
	}
	return records
}

func printUsage() {
	fmt.Printf(`Bazeth %s

Usage:
  bazeth ip <address>
  bazeth dns <domain>
  bazeth whois <domain>
  bazeth tld <extension>
  bazeth fingerprint <address>
  bazeth version
`, Version)
}
