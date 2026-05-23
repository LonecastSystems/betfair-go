package betfair

import (
	"context"
	"testing"
)

const (
	testSessionToken       = ""
	testAppKey             = ""
	testCertificateCRTPath = ""
	testCertificateKeyPath = ""
	testRest               = false
)

func newTestClient(t *testing.T) *Client {
	if testSessionToken == "" || testAppKey == "" {
		t.Skip("Invalid credentials")
	}

	tlsConfig, err := GetTLSConfig(testCertificateCRTPath, testCertificateKeyPath)
	if err != nil {
		t.Fatal(err)
	}

	config := &ClientConfig{Tls: tlsConfig, ApplicationKey: testAppKey, Rest: testRest}
	c := NewClient(config)

	if _, err := c.Resume(context.Background(), testSessionToken); err != nil {
		t.Fatal(err)
	}

	return c
}

func newTestStreamingClient(t *testing.T) *StreamingClient {
	if testSessionToken == "" || testAppKey == "" {
		t.Skip("Invalid credentials")
	}

	tlsConfig, err := GetTLSConfig(testCertificateCRTPath, testCertificateKeyPath)
	if err != nil {
		t.Fatal(err)
	}

	config := &ClientConfig{Tls: tlsConfig, ApplicationKey: testAppKey, Rest: testRest}
	c := NewClient(config)
	if _, err := c.Resume(context.Background(), testSessionToken); err != nil {
		t.Fatal(err)
	}

	sc, err := c.GetStream(&StreamingClientConfig{})
	if err != nil {
		t.Fatal(err)
	}

	return sc
}

func getRandomMarketID(t *testing.T, c *Client) string {
	mcParams := ListMarketCatalogueParams{
		Filter:     MarketFilter{},
		MaxResults: 1,
	}

	markets, err := c.ListMarketCatalogue(context.Background(), mcParams)
	if err != nil {
		t.Fatal(err)
	}

	if len(markets) == 0 {
		t.Fatal("No markets")
	}

	return markets[0].MarketID
}
