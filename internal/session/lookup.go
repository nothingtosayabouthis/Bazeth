package session

import (
	"fmt"

	"bazeth/internal/ip"
	"bazeth/internal/providers/rdap"
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

	provider := rdap.New()

	if err := provider.Enrich(result); err != nil {
		return nil, err
	}

	return result, nil
}
