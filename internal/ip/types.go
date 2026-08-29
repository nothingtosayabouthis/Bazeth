package ip

type Result struct {
	IP      string
	Version string

	Organization string
	Country      string

	CIDR         string
	StartAddress string
	EndAddress   string

	ASN string

	AbuseEmail string
	ReverseDNS string

	Source []string
}
