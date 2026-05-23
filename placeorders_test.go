package betfair

import (
	"context"
	"testing"
)

// Use with caution
func TestPlaceOrders(t *testing.T) {
	t.Skip("This may place orders, use with caution")

	c := newTestClient(t)

	params := PlaceOrdersParams{}

	report, err := c.PlaceOrders(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}

	if report.Status != ERS_SUCCESS {
		t.Fatal(report.Status)
	}
}
