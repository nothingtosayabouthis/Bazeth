package session

import (
	"fmt"

	"bazeth/internal/ip"
	"bazeth/internal/providers"

	"bazeth/internal/active/fingerprint"
	_ "bazeth/internal/providers/abuseipdb"
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
	if httpInfo, err := fingerprint.HTTP(result.IP); err == nil {
		result.HTTPStatus = httpInfo.Status
		result.HTTPServer = httpInfo.Server
		result.PoweredBy = httpInfo.PoweredBy
		result.Redirect = httpInfo.Location
		result.TLSCommonName = httpInfo.TLSCommonName
	}
	return result, nil
}
