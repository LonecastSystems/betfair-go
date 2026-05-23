package betfair

import (
	"context"
	"testing"
)

func TestMarketTypes(t *testing.T) {
	c := newTestClient(t)

	params := ListMarketTypesParams{Filter: MarketFilter{
		EventTypeIDs:    []string{"1"},
		MarketCountries: []string{"GB"},
	}}

	marketTypes, err := c.ListMarketTypes(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}

	if len(marketTypes) == 0 {
		t.Fatal("No market types")
	}
}
