package tld

import "strings"

func classify(tld string) string {
	tld = strings.TrimPrefix(strings.ToLower(tld), ".")

	switch len(tld) {

	case 2:
		return "Country Code"

	default:
		return "Generic"
	}
}
