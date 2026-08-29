package session

import (
	"fmt"

	"bazeth/internal/ip"
	"bazeth/internal/providers"

	_ "bazeth/internal/providers/asn"
	_ "bazeth/internal/providers/rdap"
	_ "bazeth/internal/providers/reverse_dns"
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

	// Execute every registered provider.
	for _, provider := range providers.All() {
		_ = provider.Enrich(result)
	}

	return result, nil
}
