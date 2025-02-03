package betfairgo

import "testing"

func TestListCurrentOrders(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := CurrentOrdersParams{
		FromRecord:  0,
		RecordCount: 1}

	report, err := c.ListCurrentOrders(params)
	if err != nil {
		t.Fatal(err)
	}

	if len(report.CurrentOrders) == 0 {
		t.Fatal("No orders")
	}
}
