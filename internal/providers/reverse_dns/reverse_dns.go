package reverse_dns

import (
	"net"

	"bazeth/internal/ip"
)

type Provider struct{}

func New() *Provider {
	return &Provider{}
}

func (p *Provider) Name() string {
	return "reverse_dns"
}

func (p *Provider) Enrich(result *ip.Result) error {
	names, err := net.LookupAddr(result.IP)
	if err != nil || len(names) == 0 {
		return nil
	}

	// Store the first PTR record and remove the trailing dot.
	result.ReverseDNS = names[0][:len(names[0])-1]
	result.Source = append(result.Source, p.Name())

	return nil
}
