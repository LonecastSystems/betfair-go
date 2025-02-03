package betfairgo

import (
	"testing"
)

func TestSubscribeToMarkets(t *testing.T) {
	c := NewTestStreamingClient(t)
	c.HeartbeatMs = 524

	marketFilter := StreamMarketFilter{MarketIDs: []string{"1.238426694"}}

	fields := []string{"EX_BEST_OFFERS_DISP", "EX_BEST_OFFERS", "EX_ALL_OFFERS", "EX_TRADED", "EX_TRADED_VOL", "EX_LTP", "EX_MARKET_DEF", "SP_TRADED", "SP_PROJECTED"}
	marketDataFilter := MarketDataFilter{Fields: fields, LadderLevels: 2}

	marketChanges, err := c.SubscribeToMarkets(marketFilter, marketDataFilter)
	if err != nil {
		t.Fatal(err)
	}

	for x := range marketChanges {
		if x.HeartbeatMs != c.HeartbeatMs {
			t.Fatal("Heartbeat does not match")
		} else {
			c.Connection.Close()
			return
		}
	}
}
