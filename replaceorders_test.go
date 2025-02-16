package betfair

import (
	"testing"
)

// Use with caution
func TestReplaceOrders(t *testing.T) {
	c := NewTestClient(t)

	params := ReplaceOrdersParams{}

	report, err := c.ReplaceOrders(params)
	if err != nil {
		t.Fatal(err)
	}

	if report.Status != ERS_SUCCESS {
		t.Fatal(report.Status)
	}
}
