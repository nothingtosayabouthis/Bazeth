package dns

import (
	"strings"

	mdns "github.com/miekg/dns"
)

func (c *Client) PTR(ip string) (string, error) {
	ptr, err := mdns.ReverseAddr(ip)
	if err != nil {
		return "", err
	}

	msg := new(mdns.Msg)
	msg.SetQuestion(ptr, mdns.TypePTR)

	resp, err := c.Exchange(msg)
	if err != nil {
		return "", err
	}

	for _, answer := range resp.Answer {
		if record, ok := answer.(*mdns.PTR); ok {
			return strings.TrimSuffix(record.Ptr, "."), nil
		}
	}

	return "", nil
}
