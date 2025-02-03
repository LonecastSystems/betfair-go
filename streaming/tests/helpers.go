package tests

import (
	"testing"

	"github.com/LonecastSystems/betfair-go/client"
	"github.com/LonecastSystems/betfair-go/streaming"
)

const sessionToken, appKey = "", ""
const certificate_crt_path = ""
const certificate_key_path = ""

func NewTestStreamingClient(t *testing.T) *streaming.StreamingClient {
	if sessionToken == "" || appKey == "" {
		t.Skip("Invalid credentials")
	}

	tlsConfig, err := client.GetTLSConfig(certificate_crt_path, certificate_key_path)
	if err != nil {
		t.Fatal(err)
	}

	c := client.NewBetfairClient(tlsConfig, appKey, "")
	if _, err := c.Resume(sessionToken); err != nil {
		t.Fatal(err)
	}

	sc := streaming.NewStreamingClient()

	if err := c.LoginStream(sc); err != nil {
		t.Fatal(err)
		return nil
	} else {
		return sc
	}
}
