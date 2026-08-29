package rdap

import "strings"

func normalizeRegistry(server string) string {
	switch {
	case strings.Contains(server, "arin"):
		return "ARIN"
	case strings.Contains(server, "apnic"):
		return "APNIC"
	case strings.Contains(server, "ripe"):
		return "RIPE"
	case strings.Contains(server, "lacnic"):
		return "LACNIC"
	case strings.Contains(server, "afrinic"):
		return "AFRINIC"
	default:
		return server
	}
}
