package dns

import (
	"fmt"

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

		case *mdns.CNAME:
			records = append(records, record.Target)

		case *mdns.DNAME:
			records = append(records, record.Target)

		case *mdns.MX:
			records = append(records, record.Mx)

		case *mdns.NS:
			records = append(records, record.Ns)

		case *mdns.SOA:
			records = append(records,
				fmt.Sprintf("%s %s", record.Ns, record.Mbox))

		case *mdns.TXT:
			records = append(records, record.Txt...)

		case *mdns.SPF:
			records = append(records, record.Txt...)

		case *mdns.CAA:
			records = append(records,
				fmt.Sprintf("%s %s", record.Tag, record.Value))

		case *mdns.SRV:
			records = append(records,
				fmt.Sprintf("%d %d %d %s",
					record.Priority,
					record.Weight,
					record.Port,
					record.Target))

		case *mdns.NAPTR:
			records = append(records, record.Replacement)

		case *mdns.SVCB:
			records = append(records, record.String())

		case *mdns.HTTPS:
			records = append(records, record.String())

		case *mdns.TLSA:
			records = append(records, record.String())

		case *mdns.SSHFP:
			records = append(records, record.String())

		case *mdns.SMIMEA:
			records = append(records, record.String())

		case *mdns.OPENPGPKEY:
			records = append(records, record.String())

		case *mdns.DNSKEY:
			records = append(records, record.String())

		case *mdns.DS:
			records = append(records, record.String())

		case *mdns.RRSIG:
			records = append(records, record.String())

		case *mdns.NSEC:
			records = append(records, record.String())

		case *mdns.NSEC3:
			records = append(records, record.String())

		case *mdns.NSEC3PARAM:
			records = append(records, record.String())

		case *mdns.URI:
			records = append(records, record.String())

		case *mdns.LOC:
			records = append(records, record.String())

		case *mdns.HINFO:
			records = append(records,
				fmt.Sprintf("%s %s", record.Cpu, record.Os))

		case *mdns.RP:
			records = append(records, record.Mbox)

		case *mdns.CERT:
			records = append(records, record.String())
		}
	}

	return records, nil
}
