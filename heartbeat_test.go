package betfairgo

import (
	"testing"
)

func TestHeartbeat(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := HeartbeatParams{PreferredTimeoutSeconds: 10}

	_, err := c.Heartbeat(params)
	if err != nil {
		t.Fatal(err)
	}
}
