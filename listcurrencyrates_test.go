package betfair

import (
	"context"
	"testing"
)

func TestListCurrencyRates(t *testing.T) {
	c := newTestClient(t)

	resp, err := c.ListCurrencyRates(context.Background(), ListCurrencyRatesParams{FromCurrency: "GBP"})
	if err != nil {
		t.Fatal(err)
	}

	if len(resp) == 0 {
		t.Fatal("No rates")
	}
}
