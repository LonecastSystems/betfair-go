package betfair

import (
	"context"
	"testing"
)

func TestHeartbeat(t *testing.T) {
	c := newTestClient(t)

	params := HeartbeatParams{PreferredTimeoutSeconds: 10}

	_, err := c.Heartbeat(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}
}
