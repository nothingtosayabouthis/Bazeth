package session

import (
	"fmt"

	"bazeth/internal/ip"
	"bazeth/internal/providers/rdap"
	"bazeth/internal/providers/reverse_dns"
)

func Lookup(address string) (*ip.Result, error) {
	version, ok := ip.Validate(address)
	if !ok {
		return nil, fmt.Errorf("invalid IP address")
	}

	result := &ip.Result{
		IP:      address,
		Version: version,
	}

	rdapProvider := rdap.New()
	if err := rdapProvider.Enrich(result); err != nil {
		return nil, err
	}

	dnsProvider := reverse_dns.New()
	_ = dnsProvider.Enrich(result)

	return result, nil
}
