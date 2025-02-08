package betfair

import "testing"

func TestListClearedOrders(t *testing.T) {
	c := NewTestClient(t)

	params := ClearedOrdersParams{BetStatus: "SETTLED"}

	report, err := c.ListClearedOrders(params)
	if err != nil {
		t.Fatal(err)
	}

	if len(report.ClearedOrders) == 0 {
		t.Fatal("No cleared orders")
	}
}
