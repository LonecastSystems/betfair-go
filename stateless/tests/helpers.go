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

	return stateless.NewStatelessClient(sessionKey, appKey, rest)
}
