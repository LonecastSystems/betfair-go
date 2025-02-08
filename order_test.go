package betfair

import (
	"testing"
)

func TestSubscribeToOrders(t *testing.T) {
	c := NewTestStreamingClient(t)
	c.HeartbeatMs = 524

	orderChanges, err := c.SubscribeToOrders(OrderFilter{})
	if err != nil {
		t.Fatal(err)
	}

	for x := range orderChanges {
		if x.HeartbeatMs != c.HeartbeatMs {
			t.Fatal("Heartbeat does not match")
		} else {
			c.Close()
			return
		}
	}
}
