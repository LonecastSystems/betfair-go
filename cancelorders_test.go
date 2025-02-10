package betfair

import (
	"fmt"
	"testing"
)

func TestCancelOrders(t *testing.T) {
	c := NewTestClient(t)

	params := CancelOrdersParams{}

	report, err := c.CancelOrders(params)
	if err != nil {
		fmt.Printf(report.CustomerRef)
		t.Fatal(err)
	}
}
