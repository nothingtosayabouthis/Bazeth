package rdap

import (
	"fmt"
	"strings"

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
	network, err := p.client.QueryIP(result.IP)
	if err != nil {
		return err
	}

	// Fill network range.
	result.CIDR = fmt.Sprintf("%s - %s", network.StartAddress, network.EndAddress)

	// Fill country.
	result.Country = network.Country

	// Fill organization.
	if len(network.Entities) > 0 {
		entity := network.Entities[0]

		if name := entity.VCard.Name(); name != "" {
			result.Organization = name
		} else if org := entity.VCard.Org(); org != "" {
			result.Organization = org
		} else {
			result.Organization = entity.Handle
		}
	}

	// Fill abuse contact.
	for _, entity := range network.Entities {
		for _, role := range entity.Roles {
			if strings.EqualFold(role, "abuse") {
				if email := entity.VCard.Email(); email != "" {
					result.AbuseEmail = email
					break
				}
			}
		}
	}

	result.Source = append(result.Source, p.Name())

	return nil
}
