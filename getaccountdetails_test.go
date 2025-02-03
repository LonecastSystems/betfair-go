package betfairgo

import "testing"

func TestGetAccountDetails(t *testing.T) {
	c := NewTestBetfairClient(t)

	resp, err := c.GetAccountDetails()
	if err != nil {
		t.Fatal(err)
	}

	if resp.CurrencyCode != "GBP" {
		t.Fatal("Details wrong")
	}
}
