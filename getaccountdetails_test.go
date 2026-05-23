package betfair

import (
	"context"
	"testing"
)

func TestGetAccountDetails(t *testing.T) {
	c := newTestClient(t)

	resp, err := c.GetAccountDetails(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if resp.CurrencyCode != "GBP" {
		t.Fatal("Details wrong")
	}
}
