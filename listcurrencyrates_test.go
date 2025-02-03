package betfairgo

import "testing"

func TestListCurrencyRates(t *testing.T) {
	c := NewTestBetfairClient(t)

	resp, err := c.ListCurrencyRates(CurrencyRateParams{FromCurrency: "GBP"})
	if err != nil {
		t.Fatal(err)
	}

	if len(resp) == 0 {
		t.Fatal("No rates")
	}
}
