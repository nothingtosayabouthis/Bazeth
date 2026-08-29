package main

import (
	"fmt"
	"os"

	"bazeth/internal/session"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: bazeth <ip>")
		os.Exit(1)
	}

	result, err := session.Lookup(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("IP:", result.IP)
	fmt.Println("Version:", result.Version)
	fmt.Println("ASN:", result.ASN)
	fmt.Println("Organization:", result.Organization)
	fmt.Println("Registry:", result.Registry)
	fmt.Println("Network:", result.NetworkName)
	fmt.Println("Country:", result.Country)
	fmt.Println("CIDR:", result.CIDR)
	fmt.Println("Abuse:", result.AbuseEmail)
	fmt.Println("Reverse DNS:", result.ReverseDNS)
	fmt.Println("Sources:", result.Source)
}
