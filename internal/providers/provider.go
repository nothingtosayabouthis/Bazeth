package providers

import "bazeth/internal/ip"

type Provider interface {
	Name() string
	Enrich(result *ip.Result) error
}
