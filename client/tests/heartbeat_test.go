package tests

import (
	"testing"

	"github.com/LonecastSystems/betfair-go/client"
)

func TestHeartbeat(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := client.HeartbeatParams{PreferredTimeoutSeconds: 10}

	_, err := c.Heartbeat(params)
	if err != nil {
		t.Fatal(err)
	}
}
