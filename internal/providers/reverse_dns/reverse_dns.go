package reverse_dns

import (
	"bazeth/internal/dns"
	"bazeth/internal/ip"
	"bazeth/internal/providers"
)

type Provider struct {
	client *dns.Client
}

func New() *Provider {
	return &Provider{
		client: dns.New(),
	}
}

func (p *Provider) Name() string {
	return "reverse_dns"
}

func (p *Provider) Enrich(result *ip.Result) error {
	name, err := p.client.PTR(result.IP)
	if err != nil || name == "" {
		return nil
	}

	result.ReverseDNS = name
	result.Source = append(result.Source, p.Name())

	return nil
}

func init() {
	providers.Register(New())
}
