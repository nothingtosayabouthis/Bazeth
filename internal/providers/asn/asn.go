package asn

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"bazeth/internal/ip"
	"bazeth/internal/providers"
)

type Provider struct{}

func New() *Provider {
	return &Provider{}
}

func (p *Provider) Name() string {
	return "asn"
}

func (p *Provider) Enrich(result *ip.Result) error {
	conn, err := net.DialTimeout("tcp", "whois.cymru.com:43", 5*time.Second)
	if err != nil {
		return nil
	}
	defer conn.Close()

	fmt.Fprintf(conn, " -v %s\n", result.IP)

	scanner := bufio.NewScanner(conn)

	// Skip header.
	if !scanner.Scan() {
		return nil
	}

	// Read result line.
	if !scanner.Scan() {
		return nil
	}

	fields := strings.Split(scanner.Text(), "|")
	if len(fields) < 1 {
		return nil
	}

	result.ASN = "AS" + strings.TrimSpace(fields[0])
	result.Source = append(result.Source, p.Name())

	return nil
}

func init() {
	providers.Register(New())
}
