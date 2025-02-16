package betfair

import (
	"fmt"
	"testing"
)

func TestUpdateOrders(t *testing.T) {
	c := NewTestClient(t)

	params := UpdateOrdersParams{}

	report, err := c.UpdateOrders(params)
	if err != nil {
		fmt.Printf(report.CustomerRef)
		t.Fatal(err)
	}
}
