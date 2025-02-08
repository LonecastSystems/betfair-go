package betfair

import "testing"

func TestMarketTypes(t *testing.T) {
	c := NewTestClient(t)

	params := MarketTypeParams{Filter: MarketFilter{
		EventTypeIDs:    []string{"1"},
		MarketCountries: []string{"GB"},
	}}

	marketTypes, err := c.ListMarketTypes(params)
	if err != nil {
		t.Fatal(err)
	}

	if len(marketTypes) == 0 {
		t.Fatal("No market types")
	}
}
