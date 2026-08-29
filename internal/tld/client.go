package tld

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const bootstrapURL = "https://data.iana.org/rdap/dns.json"

type bootstrap struct {
	Services [][]any `json:"services"`
}

func Lookup(input string) (*Result, error) {
	tld := strings.TrimPrefix(strings.ToLower(input), ".")

	result := &Result{
		TLD:  "." + tld,
		Type: classify(tld),
	}

	// First, try the official IANA RDAP bootstrap.
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(bootstrapURL)
	if err == nil {
		defer resp.Body.Close()

		var data bootstrap

		if json.NewDecoder(resp.Body).Decode(&data) == nil {
			for _, service := range data.Services {

				names, ok := service[0].([]any)
				if !ok {
					continue
				}

				urls, ok := service[1].([]any)
				if !ok {
					continue
				}

				for _, name := range names {
					if fmt.Sprint(name) == tld {
						if len(urls) > 0 {
							result.RDAP = fmt.Sprint(urls[0])
						}

						fillRegistryInfo(result)
						return result, nil
					}
				}
			}
		}
	}

	// Fallback to Bazeth's built-in registry knowledge.
	if info, ok := known[result.TLD]; ok {
		result.Manager = info.Manager
		result.WHOIS = info.WHOIS

		// Keep a sensible RDAP fallback for well-known TLDs.
		if result.RDAP == "" {
			switch result.TLD {
			case ".io":
				result.RDAP = "https://rdap.identitydigital.services/rdap/"
			}
		}

		return result, nil
	}

	return nil, fmt.Errorf("TLD not found")
}
