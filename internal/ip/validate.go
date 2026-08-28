package ip

import "net"

func Validate(address string) (string, bool) {
	ip := net.ParseIP(address)
	if ip == nil {
		return "", false
	}

	if ip.To4() != nil {
		return "IPv4", true
	}

	return "IPv6", true
}
