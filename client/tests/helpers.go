package tests

import (
	"testing"

	"github.com/LonecastSystems/betfair-go/client"
)

const sessionToken, appKey = "", ""
const certificate_crt_path = ""
const certificate_key_path = ""
const rest = false

func NewTestBetfairClient(t *testing.T) *client.BetfairClient {
	if sessionToken == "" || appKey == "" {
		t.Skip("Invalid credentials")
	}

	tlsConfig, err := client.GetTLSConfig(certificate_crt_path, certificate_key_path)
	if err != nil {
		t.Fatal(err)
	}

	c := client.NewBetfairClient(tlsConfig, appKey, "")
	c.Rest = rest

	if _, err := c.Resume(sessionToken); err != nil {
		t.Fatal(err)
	}

	return c
}
