package tld

type Result struct {
	TLD     string
	Type    string
	Manager string
	RDAP    string
	WHOIS   string
	DNSSEC  bool
}
