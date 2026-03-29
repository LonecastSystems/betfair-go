package betfair

import (
	"context"
	"testing"
)

func TestListClearedOrders(t *testing.T) {
	c := NewTestClient(t)

	params := ClearedOrdersParams{BetStatus: BS_SETTLED}

	_, err := c.ListClearedOrders(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}
}
