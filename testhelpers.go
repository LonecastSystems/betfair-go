package betfairgo

import (
	"testing"
)

const sessionToken, appKey = "", ""
const certificate_crt_path = ""
const certificate_key_path = ""
const rest = false

func NewTestBetfairClient(t *testing.T) *BetfairClient {
	if sessionToken == "" || appKey == "" {
		t.Skip("Invalid credentials")
	}

	tlsConfig, err := GetTLSConfig(certificate_crt_path, certificate_key_path)
	if err != nil {
		t.Fatal(err)
	}

	c := NewBetfairClient(tlsConfig, appKey, "")
	c.Rest = rest

	if _, err := c.Resume(sessionToken); err != nil {
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

	c := NewBetfairClient(tlsConfig, appKey, "")
	if _, err := c.Resume(sessionToken); err != nil {
		t.Fatal(err)
	}

	sc := NewStreamingClient()

	if err := c.LoginStream(sc); err != nil {
		t.Fatal(err)
		return nil
	} else {
		return sc
	}
}
