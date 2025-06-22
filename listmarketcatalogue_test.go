package betfair

import (
	"context"
	"testing"
)

func TestMarketCatalogue(t *testing.T) {
	c := NewTestClient(t)

	params := MarketCatalogueParams{
		Filter:     MarketFilter{},
		MaxResults: 10}

	markets, err := c.ListMarketCatalogue(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(markets)
	if len == 0 {
		t.Fatal("No markets")
	}
}
