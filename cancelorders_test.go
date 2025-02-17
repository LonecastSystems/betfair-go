package betfair

import (
	"testing"
)

// Use with caution
func TestCancelOrders(t *testing.T) {
	c := NewTestClient(t)

	params := CancelOrdersParams{}

	report, err := c.CancelOrders(params)
	if err != nil {
		t.Fatal(err)
	}

	if report.Status != ERS_SUCCESS {
		t.Fatal(report.Status)
	}
}
