package dns

import (
	mdns "github.com/miekg/dns"
)

func (c *Client) Resolve(domain string, qtype uint16) ([]string, error) {
	msg := new(mdns.Msg)
	msg.SetQuestion(mdns.Fqdn(domain), qtype)

	resp, err := c.Exchange(msg)
	if err != nil {
		return nil, err
	}

	var records []string

	for _, answer := range resp.Answer {
		switch record := answer.(type) {

		case *mdns.A:
			records = append(records, record.A.String())

		case *mdns.AAAA:
			records = append(records, record.AAAA.String())

		case *mdns.MX:
			records = append(records, record.Mx)

		case *mdns.NS:
			records = append(records, record.Ns)

		case *mdns.CNAME:
			records = append(records, record.Target)

		case *mdns.TXT:
			records = append(records, record.Txt...)
		}
	}

	return records, nil
}
