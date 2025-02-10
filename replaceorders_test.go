package betfair

import (
	"fmt"
	"testing"
)

func TestReplaceOrders(t *testing.T) {
	c := NewTestClient(t)

	params := ReplaceOrdersParams{}

	report, err := c.ReplaceOrders(params)
	if err != nil {
		fmt.Printf(report.CustomerRef)
		t.Fatal(err)
	}
}
