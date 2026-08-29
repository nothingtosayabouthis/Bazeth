package fingerprint

import (
	"crypto/tls"
	"net/http"
	"time"
)

type HTTPResult struct {
	Status        int
	Server        string
	PoweredBy     string
	Location      string
	TLSCommonName string
}

func HTTP(ip string) (*HTTPResult, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Head("https://" + ip)
	if err != nil {
		resp, err = client.Head("http://" + ip)
		if err != nil {
			return nil, err
		}
	}

	result := &HTTPResult{
		Status:    resp.StatusCode,
		Server:    resp.Header.Get("Server"),
		PoweredBy: resp.Header.Get("X-Powered-By"),
		Location:  resp.Header.Get("Location"),
	}

	if resp.TLS != nil && len(resp.TLS.PeerCertificates) > 0 {
		result.TLSCommonName = resp.TLS.PeerCertificates[0].Subject.CommonName
	}

	return result, nil
}
