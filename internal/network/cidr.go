package network

import (
	"fmt"
	"math/bits"
	"net"
)

// FromRange converts a start/end range into the shortest CIDR prefix.
func FromRange(start, end string) string {
	startIP := net.ParseIP(start)
	endIP := net.ParseIP(end)

	if startIP == nil || endIP == nil {
		return fmt.Sprintf("%s - %s", start, end)
	}

	// IPv4.
	if s4 := startIP.To4(); s4 != nil {
		e4 := endIP.To4()

		var xor uint32

		for i := 0; i < 4; i++ {
			xor = (xor << 8) | uint32(s4[i]^e4[i])
		}

		prefix := 32 - bits.Len32(xor)

		return fmt.Sprintf("%s/%d", s4.String(), prefix)
	}

	// IPv6.
	s16 := startIP.To16()
	e16 := endIP.To16()

	prefix := 0

	for i := 0; i < 16; i++ {

		diff := s16[i] ^ e16[i]

		if diff == 0 {
			prefix += 8
			continue
		}

		prefix += bits.LeadingZeros8(diff)
		break
	}

	return fmt.Sprintf("%s/%d", startIP.String(), prefix)
}
