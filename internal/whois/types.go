package whois

type Result struct {
	Domain     string
	Registrar  string
	Created    string
	Updated    string
	Expires    string
	Status     []string
	NameServer []string
	DNSSEC     string
}
