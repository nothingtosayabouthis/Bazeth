package network

import (
	"fmt"
	"net"
)

// FromRange converts a start/end range into CIDR notation.
func FromRange(start, end string) string {
	startIP := net.ParseIP(start)
	endIP := net.ParseIP(end)

	if startIP == nil || endIP == nil {
		return fmt.Sprintf("%s - %s", start, end)
	}

	start4 := startIP.To4()
	end4 := endIP.To4()

	if start4 == nil || end4 == nil {
		return fmt.Sprintf("%s - %s", start, end)
	}

	for prefix := 32; prefix >= 0; prefix-- {
		mask := net.CIDRMask(prefix, 32)
		network := start4.Mask(mask)

		if network.Equal(start4) {
			broadcast := make(net.IP, 4)

			for i := 0; i < 4; i++ {
				broadcast[i] = network[i] | ^mask[i]
			}

			if broadcast.Equal(end4) {
				return fmt.Sprintf("%s/%d", network.String(), prefix)
			}
		}
	}

	return fmt.Sprintf("%s - %s", start, end)
}
