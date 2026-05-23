package betfair

import (
	"context"
	"testing"
)

func TestListRunnerBook(t *testing.T) {
	c := newTestClient(t)

	params := ListRunnerBookParams{
		MarketID: getRandomMarketID(t, c),
		PriceProjection: PriceProjection{
			PriceData: []PriceData{PD_EX_ALL_OFFERS}}}

	runners, err := c.ListRunnerBook(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}

	if len(runners) == 0 {
		t.Fatal("No runner books")
	}
}
