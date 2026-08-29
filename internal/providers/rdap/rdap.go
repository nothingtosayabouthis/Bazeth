package rdap

import (
	"fmt"
	"net"
	"strings"

	"bazeth/internal/ip"
	"bazeth/internal/providers"

	netutil "bazeth/internal/network"

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

	// Fill registry information.
	if result.NetworkName == "" && network.Handle != "" {
		result.NetworkName = network.Handle
	}
	if network.Port43 != "" {
		if result.Registry == "" && network.Port43 != "" {
			result.Registry = normalizeRegistry(network.Port43)
		}
	}

	// Fill network range.
	result.StartAddress = network.StartAddress
	result.EndAddress = network.EndAddress
	result.CIDR = netutil.FromRange(
		result.StartAddress,
		result.EndAddress,
	)

	// Fill country only when RDAP provides a value.
	if network.Country != "" {
		result.Country = network.Country
	}

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

// Register the provider when the package is loaded.
func init() {
	providers.Register(New())
}

// cidrFromRange converts a network range into CIDR notation when possible.
func cidrFromRange(start, end string) string {
	startIP := net.ParseIP(start)
	endIP := net.ParseIP(end)

	if startIP == nil || endIP == nil {
		return fmt.Sprintf("%s - %s", start, end)
	}

	ones, bits := net.IPMask(startIP.To4()).Size()

	if bits == 32 {
		return fmt.Sprintf("%s/%d", start, ones)
	}

	return fmt.Sprintf("%s - %s", start, end)
}
