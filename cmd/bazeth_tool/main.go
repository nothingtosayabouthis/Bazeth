package main

import (
	"fmt"
	"os"

	"bazeth/internal/output"
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

	output.Print(result)
}
