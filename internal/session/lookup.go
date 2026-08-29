package session

import (
	"fmt"
	"sync"

	"bazeth/internal/active/fingerprint"
	"bazeth/internal/ip"
	"bazeth/internal/providers"

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

	var wg sync.WaitGroup
	var mu sync.Mutex

	// Execute every registered provider concurrently.
	for _, provider := range providers.All() {
		wg.Add(1)

		go func(p providers.Provider) {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()

			_ = p.Enrich(result)
		}(provider)
	}

	// Execute HTTP fingerprint concurrently.
	wg.Add(1)

	go func() {
		defer wg.Done()

		if httpInfo, err := fingerprint.HTTP(result.IP); err == nil {
			mu.Lock()
			defer mu.Unlock()

			result.HTTPStatus = httpInfo.Status
			result.HTTPServer = httpInfo.Server
			result.PoweredBy = httpInfo.PoweredBy
			result.Redirect = httpInfo.Location
			result.TLSCommonName = httpInfo.TLSCommonName
		}
	}()

	wg.Wait()

	return result, nil
}
