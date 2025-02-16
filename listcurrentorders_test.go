package betfair

import "testing"

func TestListCurrentOrders(t *testing.T) {
	c := NewTestClient(t)

	params := CurrentOrdersParams{
		MarketIDs:   []string{GetRandomMarketID(t, c)},
		FromRecord:  0,
		RecordCount: 1}

	_, err := c.ListCurrentOrders(params)
	if err != nil {
		t.Fatal(err)
	}
}
