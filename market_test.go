package betfair

import (
	"testing"
)

func TestSubscribeToMarkets(t *testing.T) {
	c := newTestStreamingClient(t)
	config := c.Config
	config.HeartbeatMs = 524

	marketFilter := StreamMarketFilter{MarketIDs: []string{"1.240636817"}}

	fields := []MarketDataFilterField{MDFF_EX_BEST_OFFERS_DISP, MDFF_EX_BEST_OFFERS, MDFF_EX_ALL_OFFERS, MDFF_EX_TRADED, MDFF_EX_TRADED_VOL, MDFF_EX_LTP, MDFF_EX_MARKET_DEF, MDFF_SP_TRADED, MDFF_SP_PROJECTED}
	marketDataFilter := MarketDataFilter{Fields: fields, LadderLevels: 2}

	marketChanges, err := c.SubscribeToMarkets(marketFilter, marketDataFilter)
	if err != nil {
		t.Fatal(err)
	}

	for x := range marketChanges {
		if x.HeartbeatMs != config.HeartbeatMs {
			t.Fatal("Heartbeat does not match")
		} else {
			c.Close()
			return
		}
	}
}
