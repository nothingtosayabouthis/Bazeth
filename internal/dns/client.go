package dns

import (
	mdns "github.com/miekg/dns"
)

type Client struct {
	Server  string
	Timeout int
}

func New() *Client {
	return &Client{
		Server:  "1.1.1.1:53",
		Timeout: 5,
	}
}

func (c *Client) Exchange(msg *mdns.Msg) (*mdns.Msg, error) {
	client := &mdns.Client{}

	resp, _, err := client.Exchange(msg, c.Server)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
