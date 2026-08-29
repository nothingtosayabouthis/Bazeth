package abuseipdb

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"bazeth/internal/ip"
	"bazeth/internal/providers"
)

const endpoint = "https://api.abuseipdb.com/api/v2/check"

type Provider struct {
	client *http.Client
	apiKey string
}

type response struct {
	Data struct {
		AbuseConfidenceScore int    `json:"abuseConfidenceScore"`
		UsageType            string `json:"usageType"`
	} `json:"data"`
}

func New() *Provider {
	return &Provider{
		client: &http.Client{Timeout: 8 * time.Second},
		apiKey: os.Getenv("BAZETH_ABUSEIPDB_KEY"),
	}
}

func (p *Provider) Name() string {
	return "abuseipdb"
}

func (p *Provider) Enrich(result *ip.Result) error {
	// Skip silently if no API key is configured.
	if p.apiKey == "" {
		return nil
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("ipAddress", result.IP)
	q.Add("maxAgeInDays", "90")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Key", p.apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := p.client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	var body response
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return err
	}

	result.ThreatScore = body.Data.AbuseConfidenceScore
	result.UsageType = body.Data.UsageType
	result.IsMalicious = result.ThreatScore >= 50

	result.Source = append(result.Source, p.Name())

	return nil
}

func init() {
	providers.Register(New())
}
