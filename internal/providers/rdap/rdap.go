package rdap

import (
	"fmt"
	"net"

	"bazeth/internal/ip"

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
	return "rdap"
}

func (p *Provider) Enrich(result *ip.Result) error {
	parsedIP := net.ParseIP(result.IP)
	if parsedIP == nil {
		return fmt.Errorf("invalid IP address")
	}

	response, err := p.client.QueryIP(parsedIP.String())
	if err != nil {
		return err
	}

	// Fill organization name.
	if len(response.Entities) > 0 {
		result.Organization = response.Entities[0].Handle
	}

	// Fill network range.
	if response.StartAddress != "" && response.EndAddress != "" {
		result.CIDR = fmt.Sprintf("%s - %s", response.StartAddress, response.EndAddress)
	}

	// Fill country when available.
	result.Country = response.Country

	// Record the provider used.
	result.Source = append(result.Source, p.Name())

	return nil
}
