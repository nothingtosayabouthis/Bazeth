package dns

import mdns "github.com/miekg/dns"

type RecordType struct {
	Name string
	Type uint16
}

var Supported = []RecordType{
	{"A", mdns.TypeA},
	{"AAAA", mdns.TypeAAAA},
	{"CNAME", mdns.TypeCNAME},
	{"DNAME", mdns.TypeDNAME},
	{"MX", mdns.TypeMX},
	{"NS", mdns.TypeNS},
	{"SOA", mdns.TypeSOA},
	{"TXT", mdns.TypeTXT},
	{"SPF", mdns.TypeSPF},
	{"CAA", mdns.TypeCAA},
	{"SRV", mdns.TypeSRV},
	{"NAPTR", mdns.TypeNAPTR},
	{"SVCB", mdns.TypeSVCB},
	{"HTTPS", mdns.TypeHTTPS},
	{"TLSA", mdns.TypeTLSA},
	{"SSHFP", mdns.TypeSSHFP},
	{"SMIMEA", mdns.TypeSMIMEA},
	{"OPENPGPKEY", mdns.TypeOPENPGPKEY},
	{"DNSKEY", mdns.TypeDNSKEY},
	{"DS", mdns.TypeDS},
	{"RRSIG", mdns.TypeRRSIG},
	{"NSEC", mdns.TypeNSEC},
	{"NSEC3", mdns.TypeNSEC3},
	{"NSEC3PARAM", mdns.TypeNSEC3PARAM},
	{"URI", mdns.TypeURI},
	{"LOC", mdns.TypeLOC},
	{"HINFO", mdns.TypeHINFO},
	{"RP", mdns.TypeRP},
	{"CERT", mdns.TypeCERT},
}
