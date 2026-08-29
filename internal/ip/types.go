package ip

type Result struct {
	IP      string
	Version string

	Organization string
	Country      string
	Registry     string
	NetworkName  string

	CIDR         string
	StartAddress string
	EndAddress   string

	ASN string

	AbuseEmail  string
	ReverseDNS  string
	ThreatScore int
	IsMalicious bool
	UsageType   string
	HTTPStatus  int

	HTTPServer    string
	PoweredBy     string
	Redirect      string
	TLSCommonName string
	Source        []string
}
