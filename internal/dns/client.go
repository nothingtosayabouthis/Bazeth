package dns

import (
	"time"

	mdns "github.com/miekg/dns"
)

type Client struct {
	Server  string
	Timeout time.Duration
}

func New() *Client {
	return &Client{
		Server:  "1.1.1.1:53",
		Timeout: 5 * time.Second,
	}
}

func (c *Client) Exchange(msg *mdns.Msg) (*mdns.Msg, error) {
	client := &mdns.Client{
		Timeout: c.Timeout,
	}

	resp, _, err := client.Exchange(msg, c.Server)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
