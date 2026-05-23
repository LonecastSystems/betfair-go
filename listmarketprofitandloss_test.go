package betfair

import (
	"context"
	"testing"
)

func TestListMarketProfitAndLoss(t *testing.T) {
	c := newTestClient(t)

	params := ListMarketProfitAndLossParams{
		MarketIDs: []string{getRandomMarketID(t, c)}}

	pnl, err := c.ListMarketProfitAndLoss(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(pnl)
	if len == 0 {
		t.Fatal("No PnL")
	}
}
