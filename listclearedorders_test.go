package betfair

import (
	"context"
	"testing"
)

func TestListClearedOrders(t *testing.T) {
	c := NewTestClient(t)

	params := ClearedOrdersParams{BetStatus: BS_SETTLED}

	report, err := c.ListClearedOrders(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}

	if len(report.ClearedOrders) == 0 {
		t.Fatal("No cleared orders")
	}
}
