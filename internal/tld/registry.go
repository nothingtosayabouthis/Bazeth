package tld

import "strings"

var known = map[string]struct {
	Manager string
	WHOIS   string
}{
	".com": {"Verisign Global Registry Services", "whois.verisign-grs.com"},
	".net": {"Verisign Global Registry Services", "whois.verisign-grs.com"},
	".org": {"Public Interest Registry", "whois.pir.org"},
	".io":  {"Identity Digital", "whois.nic.io"},
	".dev": {"Google Registry", "whois.nic.google"},
	".app": {"Google Registry", "whois.nic.google"},
	".br":  {"NIC.br", "whois.registro.br"},
}

func fillRegistryInfo(result *Result) {
	key := strings.ToLower(result.TLD)

	if info, ok := known[key]; ok {
		result.Manager = info.Manager
		result.WHOIS = info.WHOIS
	}
}
