package ip

import "fmt"

func Lookup(address string) (*Result, error) {
	version, ok := Validate(address)
	if !ok {
		return nil, fmt.Errorf("invalid IP address")
	}

	result := &Result{
		IP:      address,
		Version: version,
	}

	return result, nil
}
