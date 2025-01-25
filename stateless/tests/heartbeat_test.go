package tests

import (
	"testing"

	"github.com/LonecastSystems/betfair-go/stateless"
)

func TestHeartbeat(t *testing.T) {
	c := NewTestStatelessClient(t)

	params := stateless.HeartbeatParams{PreferredTimeoutSeconds: 10}

	_, err := c.Heartbeat(params)
	if err != nil {
		t.Fatal(err)
	}
}
