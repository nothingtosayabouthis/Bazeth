package asn

import (
	"bazeth/internal/ip"
	"bazeth/internal/providers"

	"github.com/openrdap/rdap"
)

type Provider struct {
	client *rdap.Client
}

func New() *Provider {
	return &Provider{
		client: &rdap.Client{},
	}
}

func (p *Provider) Name() string {
	return "asn"
}

func (p *Provider) Enrich(result *ip.Result) error {
	network, err := p.client.QueryIP(result.IP)
	if err != nil {
		return nil
	}

	if network.Handle != "" {
		result.ASN = network.Handle
	}

	result.Source = append(result.Source, p.Name())

	return nil
}

func init() {
	providers.Register(New())
}
