package providers

import "bazeth/internal/ip"

type Provider interface {
	Name() string
	Enrich(result *ip.Result) error
}

var registry []Provider

// Register adds a provider to the engine registry.
func Register(provider Provider) {
	registry = append(registry, provider)
}

// All returns every registered provider.
func All() []Provider {
	return registry
}
