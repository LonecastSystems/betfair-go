package betfair

import (
	"context"
	"testing"
)

const sessionToken, appKey = "", ""
const certificate_crt_path = ""
const certificate_key_path = ""
const rest = false

func NewTestClient(t *testing.T) *Client {
	if sessionToken == "" || appKey == "" {
		t.Skip("Invalid credentials")
	}

	tlsConfig, err := GetTLSConfig(certificate_crt_path, certificate_key_path)
	if err != nil {
		t.Fatal(err)
	}

	config := &ClientConfig{Tls: tlsConfig, ApplicationKey: appKey, Rest: rest}
	c := NewClient(config)

	if _, err := c.Resume(context.Background(), sessionToken); err != nil {
		t.Fatal(err)
	}

	return c
}

func NewTestStreamingClient(t *testing.T) *StreamingClient {
	if sessionToken == "" || appKey == "" {
		t.Skip("Invalid credentials")
	}

	tlsConfig, err := GetTLSConfig(certificate_crt_path, certificate_key_path)
	if err != nil {
		t.Fatal(err)
	}

	config := &ClientConfig{Tls: tlsConfig, ApplicationKey: appKey, Rest: rest}
	c := NewClient(config)
	if _, err := c.Resume(context.Background(), sessionToken); err != nil {
		t.Fatal(err)
	}

	if sc, err := c.GetStream(&StreamingClientConfig{}); err != nil {
		t.Fatal(err)
		return nil
	} else {
		return sc
	}
}

func GetRandomMarketID(t *testing.T, c *Client) string {
	mcParams := MarketCatalogueParams{
		Filter:     MarketFilter{},
		MaxResults: 1}

	markets, err := c.ListMarketCatalogue(context.Background(), mcParams)
	if err != nil {
		t.Fatal(err)
	}

	if len(markets) == 0 {
		t.Fatal("No markets")
	}

	return markets[0].MarketID
}
