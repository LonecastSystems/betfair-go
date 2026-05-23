package betfair

import (
	"context"
	"testing"
)

func TestMarketBook(t *testing.T) {
	c := newTestClient(t)

	params := ListMarketBookParams{
		MarketIDs: []string{getRandomMarketID(t, c)},
		PriceProjection: PriceProjection{
			PriceData: []PriceData{PD_EX_ALL_OFFERS}}}

	marketBooks, err := c.ListMarketBook(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(marketBooks)
	if len == 0 {
		t.Fatal("No market books")
	}
}
