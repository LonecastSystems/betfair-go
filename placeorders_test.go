package betfair

import (
	"testing"
)

// Use with caution
func TestPlaceOrders(t *testing.T) {
	c := NewTestClient(t)

	params := PlaceOrdersParams{}

	report, err := c.PlaceOrders(params)
	if err != nil {
		t.Fatal(err)
	}

	if report.Status != ERS_SUCCESS {
		t.Fatal(report.Status)
	}
}
