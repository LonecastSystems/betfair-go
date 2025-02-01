package tests

import (
	"testing"

	"github.com/LonecastSystems/betfair-go/stateless"
)

var sessionKey, appKey = "", ""
var rest = false

func NewTestStatelessClient(t *testing.T) *stateless.StatelessClient {
	if sessionKey == "" || appKey == "" {
		t.Skip("Invalid credentials")
	}

	c := stateless.NewStatelessClient(appKey, rest)
	if _, err := c.Client.ResumeSession(sessionKey); err != nil {
		t.Fatal(err)
	}

	return c
}
