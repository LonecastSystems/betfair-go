package betfair

import "testing"

func TestListMarketProfitAndLoss(t *testing.T) {
	c := NewTestClient(t)

	params := MarketProfitAndLossParams{
		MarketIDs: []string{GetRandomMarketID(t, c)}}

	pnl, err := c.ListMarketProfitAndLoss(params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(pnl)
	if len == 0 {
		t.Fatal("No PnL")
	}
}
