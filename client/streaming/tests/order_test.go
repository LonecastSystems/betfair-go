package tests

import (
	"testing"

	"github.com/LonecastSystems/betfair-go/client/streaming"
)

func TestSubscribeToOrders(t *testing.T) {
	c := NewStreamingClient(t)
	c.HeartbeatMs = 524

	orderChanges, err := c.SubscribeToOrders(streaming.OrderFilter{})
	if err != nil {
		t.Fatal(err)
	}

	for x := range orderChanges {
		if x.HeartbeatMs != c.HeartbeatMs {
			t.Fatal("Heartbeat does not match")
		}
		close(orderChanges)
	}
}
