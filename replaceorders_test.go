package betfair

import (
	"context"
	"testing"
)

// Use with caution
func TestReplaceOrders(t *testing.T) {
	t.Skip("This may place orders, use with caution")

	c := newTestClient(t)

	params := ReplaceOrdersParams{}

	report, err := c.ReplaceOrders(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}

	if report.Status != ERS_SUCCESS {
		t.Fatal(report.Status)
	}
}
