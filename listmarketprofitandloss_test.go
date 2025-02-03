package betfairgo

import "testing"

func TestListMarketProfitAndLoss(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := MarketProfitAndLossParams{
		MarketIDs: []string{"1.233455113"}}

	pnl, err := c.ListMarketProfitAndLoss(params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(pnl)
	if len == 0 {
		t.Fatal("No PnL")
	}
}
