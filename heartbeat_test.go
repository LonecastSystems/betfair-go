package betfair

import (
	"testing"
)

func TestHeartbeat(t *testing.T) {
	c := NewTestClient(t)

	params := HeartbeatParams{PreferredTimeoutSeconds: 10}

	_, err := c.Heartbeat(params)
	if err != nil {
		t.Fatal(err)
	}
}
