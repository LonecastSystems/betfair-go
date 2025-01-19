package tests

import (
	"testing"

	"github.com/LonecastSystems/betfair-go/client/streaming"
)

func TestSubscribeToMarkets(t *testing.T) {
	c := NewStreamingClient(t)
	c.HeartbeatMs = 524

	fields := []string{"EX_BEST_OFFERS_DISP", "EX_BEST_OFFERS", "EX_ALL_OFFERS", "EX_TRADED", "EX_TRADED_VOL", "EX_LTP", "EX_MARKET_DEF", "SP_TRADED", "SP_PROJECTED"}

	marketChanges, err := c.SubscribeToMarkets(streaming.MarketFilter{MarketIds: []string{"1.237874661"}}, streaming.MarketDataFilter{Fields: fields, LadderLevels: 2})
	if err != nil {
		t.Fatal(err)
	}

	for x := range marketChanges {
		if x.HeartbeatMs != c.HeartbeatMs {
			t.Fatal("Heartbeat does not match")
		}
		close(marketChanges)
	}
}
