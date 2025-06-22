package betfair

import (
	"context"
	"testing"
)

// Use with caution
func TestUpdateOrders(t *testing.T) {
	c := NewTestClient(t)

	params := UpdateOrdersParams{}

	report, err := c.UpdateOrders(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}

	if report.Status != ERS_SUCCESS {
		t.Fatal(report.Status)
	}
}
