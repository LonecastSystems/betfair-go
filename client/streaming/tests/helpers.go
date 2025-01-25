package tests

import (
	"testing"

	"github.com/LonecastSystems/betfair-go/client/streaming"
	"github.com/LonecastSystems/betfair-go/helpers"
)

var sessionKey, appKey = "", ""
var certificate_crt_path = ""
var certificate_key_path = ""

func NewStreamingClient(t *testing.T) *streaming.StreamingClient {
	if sessionKey == "" || appKey == "" {
		t.Skip("Invalid credentials")
	}

	c := streaming.NewStreamingClient(sessionKey, appKey)

	tlsConfig, err := helpers.GetTLSConfig(certificate_crt_path, certificate_key_path)
	if err != nil {
		t.Fatal(err)
	}

	if err := c.Login(tlsConfig); err != nil {
		t.Fatal(err)
	}

	return c
}
