package betfair

import (
	"fmt"
	"testing"
)

func TestPlaceOrders(t *testing.T) {
	c := NewTestClient(t)

	params := PlaceOrdersParams{}

	report, err := c.PlaceOrders(params)
	if err != nil {
		fmt.Printf(report.CustomerRef)
		t.Fatal(err)
	}
}
