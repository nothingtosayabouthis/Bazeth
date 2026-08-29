package whois

import (
	"github.com/openrdap/rdap"
)

type Client struct {
	rdap *rdap.Client
}

func New() *Client {
	return &Client{
		rdap: &rdap.Client{},
	}
}

func (c *Client) Lookup(domain string) (*Result, error) {
	data, err := c.rdap.QueryDomain(domain)
	if err != nil {
		return nil, err
	}

	result := &Result{
		Domain: domain,
	}

	// Fill registrar information.
	if entity := bestRegistrar(data.Entities); entity != nil {
		if name := entity.VCard.Name(); name != "" {
			result.Registrar = name
		} else if org := entity.VCard.Org(); org != "" {
			result.Registrar = org
		}
	}

	// Fill domain events.
	for _, event := range data.Events {
		switch event.Action {
		case "registration":
			result.Created = event.Date

		case "expiration":
			result.Expires = event.Date

		case "last changed", "last update":
			result.Updated = event.Date
		}
	}

	// Fill domain status.
	result.Status = append(result.Status, data.Status...)

	// Fill nameservers.
	for _, ns := range data.Nameservers {
		result.NameServer = append(result.NameServer, ns.LDHName)
	}

	// Fill DNSSEC status.
	if data.SecureDNS != nil &&
		data.SecureDNS.DelegationSigned != nil &&
		*data.SecureDNS.DelegationSigned {

		result.DNSSEC = "Enabled"
	} else {
		result.DNSSEC = "Disabled"
	}

	return result, nil
}
